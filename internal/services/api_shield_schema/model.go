// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_schema

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APIShieldSchemaResultEnvelope struct {
	Result APIShieldSchemaModel `json:"result,computed"`
}

type APIShieldSchemaModel struct {
	ZoneID            types.String `tfsdk:"zone_id" path:"zone_id"`
	SchemaID          types.String `tfsdk:"schema_id" path:"schema_id"`
	ValidationEnabled types.String `tfsdk:"validation_enabled" json:"validation_enabled"`
	File              types.String `tfsdk:"file" json:"file"`
	Kind              types.String `tfsdk:"kind" json:"kind"`
	Name              types.String `tfsdk:"name" json:"name"`
	CreatedAt         types.String `tfsdk:"created_at" json:"created_at,computed"`
	Source            types.String `tfsdk:"source" json:"source,computed"`
}

type APIShieldSchemaSchemaModel struct {
	CreatedAt         types.String `tfsdk:"created_at" json:"created_at,computed"`
	Kind              types.String `tfsdk:"kind" json:"kind"`
	Name              types.String `tfsdk:"name" json:"name"`
	SchemaID          types.String `tfsdk:"schema_id" json:"schema_id,computed"`
	Source            types.String `tfsdk:"source" json:"source"`
	ValidationEnabled types.Bool   `tfsdk:"validation_enabled" json:"validation_enabled"`
}

type APIShieldSchemaUploadDetailsModel struct {
	Warnings *[]*APIShieldSchemaUploadDetailsWarningsModel `tfsdk:"warnings" json:"warnings"`
}

type APIShieldSchemaUploadDetailsWarningsModel struct {
	Code      types.Int64     `tfsdk:"code" json:"code"`
	Locations *[]types.String `tfsdk:"locations" json:"locations"`
	Message   types.String    `tfsdk:"message" json:"message"`
}