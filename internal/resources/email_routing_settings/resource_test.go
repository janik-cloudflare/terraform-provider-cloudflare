package email_routing_settings_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/stainless-sdks/cloudflare-terraform/internal/acctest"
	"github.com/stainless-sdks/cloudflare-terraform/internal/consts"
	"github.com/stainless-sdks/cloudflare-terraform/internal/utils"
)

func testEmailRoutingSettingsConfig(resourceID, zoneID string, enabled bool) string {
	return fmt.Sprintf(`
		resource "cloudflare_email_routing_settings" "%[1]s" {
		  zone_id = "%[2]s"
		  enabled = "%[3]t"
		}
		`, resourceID, zoneID, enabled)
}

func TestAccTestEmailRoutingSettings(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_email_routing_settings." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testEmailRoutingSettingsConfig(rnd, zoneID, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "enabled", "true"),
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
				),
			},
		},
	})
}
