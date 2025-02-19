// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-nxos"
	"github.com/netascode/terraform-provider-nxos/internal/provider/helpers"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &PhysicalInterfaceDataSource{}
	_ datasource.DataSourceWithConfigure = &PhysicalInterfaceDataSource{}
)

func NewPhysicalInterfaceDataSource() datasource.DataSource {
	return &PhysicalInterfaceDataSource{}
}

type PhysicalInterfaceDataSource struct {
	data *NxosProviderData
}

func (d *PhysicalInterfaceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_physical_interface"
}

func (d *PhysicalInterfaceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This data source can read the configuration of a physical interface.", "l1PhysIf", "System/l1:PhysIf/").String,

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The distinguished name of the object.",
				Computed:            true,
			},
			"interface_id": schema.StringAttribute{
				MarkdownDescription: "Must match first field in the output of `show intf brief`. Example: `eth1/1`.",
				Required:            true,
			},
			"fec_mode": schema.StringAttribute{
				MarkdownDescription: "FEC mode.",
				Computed:            true,
			},
			"access_vlan": schema.StringAttribute{
				MarkdownDescription: "Access VLAN. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.",
				Computed:            true,
			},
			"admin_state": schema.StringAttribute{
				MarkdownDescription: "Administrative port state.",
				Computed:            true,
			},
			"auto_negotiation": schema.StringAttribute{
				MarkdownDescription: "Administrative port auto-negotiation.",
				Computed:            true,
			},
			"bandwidth": schema.Int64Attribute{
				MarkdownDescription: "The bandwidth parameter for a routed interface, port channel, or subinterface.",
				Computed:            true,
			},
			"delay": schema.Int64Attribute{
				MarkdownDescription: "The administrative port delay time.",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Interface description.",
				Computed:            true,
			},
			"duplex": schema.StringAttribute{
				MarkdownDescription: "Duplex.",
				Computed:            true,
			},
			"layer": schema.StringAttribute{
				MarkdownDescription: "Administrative port layer.",
				Computed:            true,
			},
			"link_logging": schema.StringAttribute{
				MarkdownDescription: "Administrative link logging.",
				Computed:            true,
			},
			"link_debounce_down": schema.Int64Attribute{
				MarkdownDescription: "Administrative port link debounce interval.",
				Computed:            true,
			},
			"link_debounce_up": schema.Int64Attribute{
				MarkdownDescription: "Link Debounce Interval - LinkUp Event.",
				Computed:            true,
			},
			"medium": schema.StringAttribute{
				MarkdownDescription: "The administrative port medium type.",
				Computed:            true,
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: "Administrative port mode.",
				Computed:            true,
			},
			"mtu": schema.Int64Attribute{
				MarkdownDescription: "Administrative port MTU.",
				Computed:            true,
			},
			"native_vlan": schema.StringAttribute{
				MarkdownDescription: "Native VLAN. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.",
				Computed:            true,
			},
			"speed": schema.StringAttribute{
				MarkdownDescription: "Administrative port speed.",
				Computed:            true,
			},
			"speed_group": schema.StringAttribute{
				MarkdownDescription: "Speed group.",
				Computed:            true,
			},
			"trunk_vlans": schema.StringAttribute{
				MarkdownDescription: "List of trunk VLANs.",
				Computed:            true,
			},
			"uni_directional_ethernet": schema.StringAttribute{
				MarkdownDescription: "UDE (Uni-Directional Ethernet).",
				Computed:            true,
			},
			"user_configured_flags": schema.StringAttribute{
				MarkdownDescription: "Port User Config Flags.",
				Computed:            true,
			},
		},
	}
}

func (d *PhysicalInterfaceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.data = req.ProviderData.(*NxosProviderData)
}

func (d *PhysicalInterfaceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config, state PhysicalInterface

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
