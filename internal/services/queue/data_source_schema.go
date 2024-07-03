// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package queue

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = &QueueDataSource{}
var _ datasource.DataSourceWithValidateConfig = &QueueDataSource{}

func (r QueueDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"queue_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"consumers": schema.StringAttribute{
				Optional: true,
			},
			"consumers_total_count": schema.StringAttribute{
				Optional: true,
			},
			"created_on": schema.StringAttribute{
				Optional: true,
			},
			"modified_on": schema.StringAttribute{
				Optional: true,
			},
			"producers": schema.StringAttribute{
				Optional: true,
			},
			"producers_total_count": schema.StringAttribute{
				Optional: true,
			},
			"queue_name": schema.StringAttribute{
				Optional: true,
			},
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
				},
			},
		},
	}
}

func (r *QueueDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *QueueDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}