// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_tag

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessTagResultEnvelope struct {
	Result ZeroTrustAccessTagModel `json:"result"`
}

type ZeroTrustAccessTagModel struct {
	ID        types.String      `tfsdk:"id" json:"-,computed"`
	Name      types.String      `tfsdk:"name" json:"name"`
	AccountID types.String      `tfsdk:"account_id" path:"account_id"`
	AppCount  types.Int64       `tfsdk:"app_count" json:"app_count,computed"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed"`
	UpdatedAt timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed"`
}