package provider

import (
	"context"
	"fmt"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceCloudflareRecord() *schema.Resource {
	return &schema.Resource{
		Description: heredoc.Doc(`
			Use this data source to lookup a [DNS Record](https://api.cloudflare.com/#dns-records-for-a-zone-properties)
		`),
		ReadContext: dataSourceCloudflareRecordRead,
		Schema: map[string]*schema.Schema{
			"zone_id": {
				Description: "The zone identifier to target for the resource.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "A",
				ValidateFunc: validation.StringInSlice([]string{"A", "AAAA", "CAA", "CNAME", "TXT", "SRV", "LOC", "MX", "NS", "SPF", "CERT", "DNSKEY", "DS", "NAPTR", "SMIMEA", "SSHFP", "TLSA", "URI", "PTR", "HTTPS"}, false),
			},
			"priority": {
				Type:             schema.TypeInt,
				Optional:         true,
				DiffSuppressFunc: suppressPriority,
			},
			"value": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"proxied": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"ttl": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"proxiable": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"locked": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"zone_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceCloudflareRecordRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	searchRecord := cloudflare.DNSRecord{
		Name: d.Get("hostname").(string),
		Type: d.Get("type").(string),
	}
	if priority, ok := d.GetOkExists("priority"); ok {
		p := uint16(priority.(int))
		searchRecord.Priority = &p
	}
	records, err := client.DNSRecords(ctx, zoneID, searchRecord)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error listing DNS records: %w", err))
	}

	if len(records) == 0 {
		return diag.Errorf("didn't get any DNS records for hostname: %s", searchRecord.Name)
	}

	if len(records) != 1 && !contains([]string{"MX", "URI"}, searchRecord.Type) {
		return diag.Errorf("only wanted 1 DNS record. Got %d records", len(records))
	} else {
		for _, record := range records {
			if record.Priority == searchRecord.Priority {
				records = []cloudflare.DNSRecord{record}
				break
			}
		}
		if len(records) != 1 {
			return diag.Errorf("unable to find single record for %s type %s", searchRecord.Name, searchRecord.Type)
		}
	}

	record := records[0]
	d.SetId(record.ID)
	d.Set("type", record.Type)
	d.Set("value", record.Content)
	d.Set("proxied", record.Proxied)
	d.Set("ttl", record.TTL)
	d.Set("proxiable", record.Proxiable)
	d.Set("locked", record.Locked)
	d.Set("zone_name", record.ZoneName)

	if record.Priority != nil {
		priority := record.Priority
		p := *priority
		d.Set("priority", int(p))
	}

	return nil
}
