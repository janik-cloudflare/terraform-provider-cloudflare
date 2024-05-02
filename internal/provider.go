// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package internal

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-terraform/internal/resources/access_application"
	"github.com/cloudflare/cloudflare-terraform/internal/resources/access_identity_provider"
	"github.com/cloudflare/cloudflare-terraform/internal/resources/access_mutual_tls_certificate"
	"github.com/cloudflare/cloudflare-terraform/internal/resources/account_member"
	"github.com/cloudflare/cloudflare-terraform/internal/resources/cloudforce_one_request"
	"github.com/cloudflare/cloudflare-terraform/internal/resources/cloudforce_one_request_priority"
	"github.com/cloudflare/cloudflare-terraform/internal/resources/device_posture_rule"
	"github.com/cloudflare/cloudflare-terraform/internal/resources/email_routing_rule"
	"github.com/cloudflare/cloudflare-terraform/internal/resources/firewall_waf_override"
	"github.com/cloudflare/cloudflare-terraform/internal/resources/intel_indicator_feed"
	"github.com/cloudflare/cloudflare-terraform/internal/resources/logpush_job"
	"github.com/cloudflare/cloudflare-terraform/internal/resources/magic_network_monitoring_rule"
	"github.com/cloudflare/cloudflare-terraform/internal/resources/magic_transit_site"
	"github.com/cloudflare/cloudflare-terraform/internal/resources/magic_transit_site_acl"
	"github.com/cloudflare/cloudflare-terraform/internal/resources/notification_policy"
	"github.com/cloudflare/cloudflare-terraform/internal/resources/notification_policy_webhooks"
	"github.com/cloudflare/cloudflare-terraform/internal/resources/page_shield_policy"
	"github.com/cloudflare/cloudflare-terraform/internal/resources/record"
	"github.com/cloudflare/cloudflare-terraform/internal/resources/ruleset"
	"github.com/cloudflare/cloudflare-terraform/internal/resources/secondary_dns_acl"
	"github.com/cloudflare/cloudflare-terraform/internal/resources/secondary_dns_peer"
	"github.com/cloudflare/cloudflare-terraform/internal/resources/secondary_dns_tsig"
	"github.com/cloudflare/cloudflare-terraform/internal/resources/spectrum_application"
	"github.com/cloudflare/cloudflare-terraform/internal/resources/teams_list"
	"github.com/cloudflare/cloudflare-terraform/internal/resources/teams_location"
	"github.com/cloudflare/cloudflare-terraform/internal/resources/waiting_room"
	"github.com/cloudflare/cloudflare-terraform/internal/resources/waiting_room_event"
	"github.com/cloudflare/cloudflare-terraform/internal/resources/web3_hostname_ipfs_universal_path_content_list_entry"
	"github.com/cloudflare/cloudflare-terraform/internal/resources/zero_trust_access_bookmark"
	"github.com/cloudflare/cloudflare-terraform/internal/resources/zone_lockdown"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.Provider = &CloudflareProvider{}

// CloudflareProvider defines the provider implementation.
type CloudflareProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// CloudflareProviderModel describes the provider data model.
type CloudflareProviderModel struct {
	BaseURL        types.String `tfsdk:"base_url" json:"base_url"`
	APIToken       types.String `tfsdk:"api_token" json:"api_token"`
	APIKey         types.String `tfsdk:"api_key" json:"api_key"`
	APIEmail       types.String `tfsdk:"api_email" json:"api_email"`
	UserServiceKey types.String `tfsdk:"user_service_key" json:"user_service_key"`
}

func (p *CloudflareProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "cloudflare"
	resp.Version = p.version
}

func (p CloudflareProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"base_url": schema.StringAttribute{
				Description: "Set the base url that the provider connects to. This can be used for testing in other environments.",
				Optional:    true,
			},
			"api_token": schema.StringAttribute{
				Optional: true,
			},
			"api_key": schema.StringAttribute{
				Optional: true,
			},
			"api_email": schema.StringAttribute{
				Optional: true,
			},
			"user_service_key": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

func (p *CloudflareProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

	// TODO(terraform): apiKey := os.Getenv("API_KEY")

	var data CloudflareProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	opts := []option.RequestOption{}

	if !data.BaseURL.IsNull() {
		opts = append(opts, option.WithBaseURL(data.BaseURL.ValueString()))
	}
	if !data.APIToken.IsNull() {
		opts = append(opts, option.WithAPIToken(data.APIToken.ValueString()))
	}
	if !data.APIKey.IsNull() {
		opts = append(opts, option.WithAPIKey(data.APIKey.ValueString()))
	}
	if !data.APIEmail.IsNull() {
		opts = append(opts, option.WithAPIEmail(data.APIEmail.ValueString()))
	}
	if !data.UserServiceKey.IsNull() {
		opts = append(opts, option.WithUserServiceKey(data.UserServiceKey.ValueString()))
	}

	client := cloudflare.NewClient(
		opts...,
	)

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *CloudflareProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		account_member.NewResource,
		record.NewResource,
		email_routing_rule.NewResource,
		zone_lockdown.NewResource,
		firewall_waf_override.NewResource,
		logpush_job.NewResource,
		secondary_dns_acl.NewResource,
		secondary_dns_peer.NewResource,
		secondary_dns_tsig.NewResource,
		waiting_room.NewResource,
		waiting_room_event.NewResource,
		web3_hostname_ipfs_universal_path_content_list_entry.NewResource,
		page_shield_policy.NewResource,
		ruleset.NewResource,
		spectrum_application.NewResource,
		intel_indicator_feed.NewResource,
		magic_transit_site.NewResource,
		magic_transit_site_acl.NewResource,
		magic_network_monitoring_rule.NewResource,
		notification_policy_webhooks.NewResource,
		notification_policy.NewResource,
		device_posture_rule.NewResource,
		access_identity_provider.NewResource,
		access_application.NewResource,
		access_mutual_tls_certificate.NewResource,
		zero_trust_access_bookmark.NewResource,
		teams_list.NewResource,
		teams_location.NewResource,
		cloudforce_one_request.NewResource,
		cloudforce_one_request_priority.NewResource,
	}
}

func (p *CloudflareProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func NewProvider(version string) func() provider.Provider {
	return func() provider.Provider {
		return &CloudflareProvider{
			version: version,
		}
	}
}