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
	_ datasource.DataSource              = &PIMInterfaceDataSource{}
	_ datasource.DataSourceWithConfigure = &PIMInterfaceDataSource{}
)

func NewPIMInterfaceDataSource() datasource.DataSource {
	return &PIMInterfaceDataSource{}
}

type PIMInterfaceDataSource struct {
	data *NxosProviderData
}

func (d *PIMInterfaceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_pim_interface"
}

func (d *PIMInterfaceDataSource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This data source can read the PIM interface configuration.", "pimIf", "Layer%203/pim:If/").String,

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
			"vrf_name": {
				MarkdownDescription: "VRF name.",
				Type:                types.StringType,
				Required:            true,
			},
			"interface_id": {
				MarkdownDescription: "Must match first field in the output of `show intf brief`. Example: `eth1/1`.",
				Type:                types.StringType,
				Required:            true,
			},
			"admin_state": {
				MarkdownDescription: "Administrative state.",
				Type:                types.StringType,
				Computed:            true,
			},
			"bfd": {
				MarkdownDescription: "BFD.",
				Type:                types.StringType,
				Computed:            true,
			},
			"dr_priority": {
				MarkdownDescription: "Designated Router priority level.",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"passive": {
				MarkdownDescription: "Passive interface.",
				Type:                types.BoolType,
				Computed:            true,
			},
			"sparse_mode": {
				MarkdownDescription: "Sparse mode.",
				Type:                types.BoolType,
				Computed:            true,
			},
		},
	}, nil
}

func (d *PIMInterfaceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.data = req.ProviderData.(*NxosProviderData)
}

func (d *PIMInterfaceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config, state PIMInterface

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
