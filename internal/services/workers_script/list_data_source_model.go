// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_script

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/workers"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersScriptsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[WorkersScriptsResultDataSourceModel] `json:"result,computed"`
}

type WorkersScriptsDataSourceModel struct {
	AccountID types.String                                                      `tfsdk:"account_id" path:"account_id"`
	MaxItems  types.Int64                                                       `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[WorkersScriptsResultDataSourceModel] `tfsdk:"result"`
}

func (m *WorkersScriptsDataSourceModel) toListParams() (params workers.ScriptListParams, diags diag.Diagnostics) {
	params = workers.ScriptListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type WorkersScriptsResultDataSourceModel struct {
	ID            types.String                                   `tfsdk:"id" json:"id,computed"`
	CreatedOn     timetypes.RFC3339                              `tfsdk:"created_on" json:"created_on,computed"`
	Etag          types.String                                   `tfsdk:"etag" json:"etag,computed"`
	Logpush       types.Bool                                     `tfsdk:"logpush" json:"logpush,computed_optional"`
	ModifiedOn    timetypes.RFC3339                              `tfsdk:"modified_on" json:"modified_on,computed"`
	PlacementMode types.String                                   `tfsdk:"placement_mode" json:"placement_mode,computed_optional"`
	TailConsumers *[]*WorkersScriptsTailConsumersDataSourceModel `tfsdk:"tail_consumers" json:"tail_consumers,computed_optional"`
	UsageModel    types.String                                   `tfsdk:"usage_model" json:"usage_model,computed_optional"`
}

type WorkersScriptsTailConsumersDataSourceModel struct {
	Service     types.String `tfsdk:"service" json:"service,computed"`
	Environment types.String `tfsdk:"environment" json:"environment,computed_optional"`
	Namespace   types.String `tfsdk:"namespace" json:"namespace,computed_optional"`
}