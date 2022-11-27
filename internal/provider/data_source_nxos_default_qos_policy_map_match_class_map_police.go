// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-nxos"
	"github.com/netascode/terraform-provider-nxos/internal/provider/helpers"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &DefaultQOSPolicyMapMatchClassMapPoliceDataSource{}
	_ datasource.DataSourceWithConfigure = &DefaultQOSPolicyMapMatchClassMapPoliceDataSource{}
)

func NewDefaultQOSPolicyMapMatchClassMapPoliceDataSource() datasource.DataSource {
	return &DefaultQOSPolicyMapMatchClassMapPoliceDataSource{}
}

type DefaultQOSPolicyMapMatchClassMapPoliceDataSource struct {
	data *NxosProviderData
}

func (d *DefaultQOSPolicyMapMatchClassMapPoliceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_default_qos_policy_map_match_class_map_police"
}

func (d *DefaultQOSPolicyMapMatchClassMapPoliceDataSource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This data source can read the default QoS policy map match class map police configuration.", "ipqosPolice", "Qos/ipqos:Police/").String,

		Attributes: map[string]tfsdk.Attribute{
			"device": {
				MarkdownDescription: "A device name from the provider configuration.",
				Type:                types.StringType,
				Optional:            true,
			},
			"id": {
				MarkdownDescription: "The distinguished name of the object.",
				Type:                types.StringType,
				Computed:            true,
			},
			"policy_map_name": {
				MarkdownDescription: "Policy map name.",
				Type:                types.StringType,
				Required:            true,
			},
			"class_map_name": {
				MarkdownDescription: "Class map name.",
				Type:                types.StringType,
				Required:            true,
			},
			"bc_rate": {
				MarkdownDescription: "CIR burst rate.",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"bc_unit": {
				MarkdownDescription: "CIR burst rate unit.",
				Type:                types.StringType,
				Computed:            true,
			},
			"be_rate": {
				MarkdownDescription: "PIR burst rate.",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"be_unit": {
				MarkdownDescription: "PIR burst rate unit.",
				Type:                types.StringType,
				Computed:            true,
			},
			"cir_rate": {
				MarkdownDescription: "CIR rate.",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"cir_unit": {
				MarkdownDescription: "CIR rate unit.",
				Type:                types.StringType,
				Computed:            true,
			},
			"conform_action": {
				MarkdownDescription: "Conform action.",
				Type:                types.StringType,
				Computed:            true,
			},
			"conform_set_cos": {
				MarkdownDescription: "Set CoS for conforming traffic.",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"conform_set_dscp": {
				MarkdownDescription: "Set DSCP for conforming traffic.",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"conform_set_precedence": {
				MarkdownDescription: "Set precedence for conforming traffic.",
				Type:                types.StringType,
				Computed:            true,
			},
			"conform_set_qos_group": {
				MarkdownDescription: "Set qos-group for conforming traffic.",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"exceed_action": {
				MarkdownDescription: "Exceed action.",
				Type:                types.StringType,
				Computed:            true,
			},
			"exceed_set_cos": {
				MarkdownDescription: "Set CoS for exceeding traffic.",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"exceed_set_dscp": {
				MarkdownDescription: "Set DSCP for exceeding traffic.",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"exceed_set_precedence": {
				MarkdownDescription: "Set precedence for exceeding traffic.",
				Type:                types.StringType,
				Computed:            true,
			},
			"exceed_set_qos_group": {
				MarkdownDescription: "Set qos-group for exceeding traffic.",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"pir_rate": {
				MarkdownDescription: "PIR rate.",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"pir_unit": {
				MarkdownDescription: "PIR rate unit.",
				Type:                types.StringType,
				Computed:            true,
			},
			"violate_action": {
				MarkdownDescription: "Violate action.",
				Type:                types.StringType,
				Computed:            true,
			},
			"violate_set_cos": {
				MarkdownDescription: "Set CoS for violating traffic.",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"violate_set_dscp": {
				MarkdownDescription: "Set DSCP for violating traffic.",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"violate_set_precedence": {
				MarkdownDescription: "Set precedence for violating traffic.",
				Type:                types.StringType,
				Computed:            true,
			},
			"violate_set_qos_group": {
				MarkdownDescription: "Set qos-group for violating traffic.",
				Type:                types.Int64Type,
				Computed:            true,
			},
		},
	}, nil
}

func (d *DefaultQOSPolicyMapMatchClassMapPoliceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.data = req.ProviderData.(*NxosProviderData)
}

func (d *DefaultQOSPolicyMapMatchClassMapPoliceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config, state DefaultQOSPolicyMapMatchClassMapPolice

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getDn()))

	res, err := d.data.client.GetDn(config.getDn(), nxos.OverrideUrl(d.data.devices[config.Device.ValueString()]))

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	state.fromBody(res)
	state.fromPlan(config)
	state.Dn = types.StringValue(config.getDn())

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.getDn()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}
