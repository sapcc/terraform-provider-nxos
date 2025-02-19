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
	_ datasource.DataSource              = &BGPPeerDataSource{}
	_ datasource.DataSourceWithConfigure = &BGPPeerDataSource{}
)

func NewBGPPeerDataSource() datasource.DataSource {
	return &BGPPeerDataSource{}
}

type BGPPeerDataSource struct {
	data *NxosProviderData
}

func (d *BGPPeerDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_bgp_peer"
}

func (d *BGPPeerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This data source can read the BGP peer configuration.", "bgpPeer", "Routing%20and%20Forwarding/bgp:Peer/").String,

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The distinguished name of the object.",
				Computed:            true,
			},
			"asn": schema.StringAttribute{
				MarkdownDescription: "Autonomous system number.",
				Required:            true,
			},
			"vrf": schema.StringAttribute{
				MarkdownDescription: "VRF name.",
				Required:            true,
			},
			"address": schema.StringAttribute{
				MarkdownDescription: "Peer address.",
				Required:            true,
			},
			"remote_asn": schema.StringAttribute{
				MarkdownDescription: "Peer autonomous system number.",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Peer description.",
				Computed:            true,
			},
			"peer_template": schema.StringAttribute{
				MarkdownDescription: "Peer template name.",
				Computed:            true,
			},
			"peer_type": schema.StringAttribute{
				MarkdownDescription: "Neighbor Fabric Type.",
				Computed:            true,
			},
			"source_interface": schema.StringAttribute{
				MarkdownDescription: "Source Interface. Must match first field in the output of `show intf brief`.",
				Computed:            true,
			},
		},
	}
}

func (d *BGPPeerDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.data = req.ProviderData.(*NxosProviderData)
}

func (d *BGPPeerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config, state BGPPeer

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
