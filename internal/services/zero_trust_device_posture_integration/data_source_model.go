// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_posture_integration

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDevicePostureIntegrationResultDataSourceEnvelope struct {
	Result ZeroTrustDevicePostureIntegrationDataSourceModel `json:"result,computed"`
}

type ZeroTrustDevicePostureIntegrationResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustDevicePostureIntegrationDataSourceModel] `json:"result,computed"`
}

type ZeroTrustDevicePostureIntegrationDataSourceModel struct {
	AccountID     types.String                                               `tfsdk:"account_id" path:"account_id"`
	IntegrationID types.String                                               `tfsdk:"integration_id" path:"integration_id"`
	ID            types.String                                               `tfsdk:"id" json:"id,computed_optional"`
	Interval      types.String                                               `tfsdk:"interval" json:"interval,computed_optional"`
	Name          types.String                                               `tfsdk:"name" json:"name,computed_optional"`
	Type          types.String                                               `tfsdk:"type" json:"type,computed_optional"`
	Config        *ZeroTrustDevicePostureIntegrationConfigDataSourceModel    `tfsdk:"config" json:"config,computed_optional"`
	Filter        *ZeroTrustDevicePostureIntegrationFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *ZeroTrustDevicePostureIntegrationDataSourceModel) toReadParams() (params zero_trust.DevicePostureIntegrationGetParams, diags diag.Diagnostics) {
	params = zero_trust.DevicePostureIntegrationGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *ZeroTrustDevicePostureIntegrationDataSourceModel) toListParams() (params zero_trust.DevicePostureIntegrationListParams, diags diag.Diagnostics) {
	params = zero_trust.DevicePostureIntegrationListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	return
}

type ZeroTrustDevicePostureIntegrationConfigDataSourceModel struct {
	APIURL   types.String `tfsdk:"api_url" json:"api_url,computed"`
	AuthURL  types.String `tfsdk:"auth_url" json:"auth_url,computed"`
	ClientID types.String `tfsdk:"client_id" json:"client_id,computed"`
}

type ZeroTrustDevicePostureIntegrationFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}