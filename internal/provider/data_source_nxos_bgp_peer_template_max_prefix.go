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
	_ datasource.DataSource              = &BGPPeerTemplateMaxPrefixDataSource{}
	_ datasource.DataSourceWithConfigure = &BGPPeerTemplateMaxPrefixDataSource{}
)

func NewBGPPeerTemplateMaxPrefixDataSource() datasource.DataSource {
	return &BGPPeerTemplateMaxPrefixDataSource{}
}

type BGPPeerTemplateMaxPrefixDataSource struct {
	data *NxosProviderData
}

func (d *BGPPeerTemplateMaxPrefixDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_bgp_peer_template_max_prefix"
}

func (d *BGPPeerTemplateMaxPrefixDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This data source can read the BGP peer template Maximum Prefix Policy configuration.", "bgpMaxPfxP", "Routing%20and%20Forwarding/bgp:MaxPfxP/").String,

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
			"template_name": schema.StringAttribute{
				MarkdownDescription: "Peer template name.",
				Required:            true,
			},
			"address_family": schema.StringAttribute{
				MarkdownDescription: "Address Family.",
				Required:            true,
			},
			"action": schema.StringAttribute{
				MarkdownDescription: "Action to do when limit is exceeded.",
				Computed:            true,
			},
			"maximum_prefix": schema.Int64Attribute{
				MarkdownDescription: "Maximum number of prefixes allowed from the peer.",
				Computed:            true,
			},
			"restart_time": schema.Int64Attribute{
				MarkdownDescription: "The period of time in minutes before restarting the peer when the prefix limit is reached.",
				Computed:            true,
			},
			"threshold": schema.Int64Attribute{
				MarkdownDescription: "The period of time in minutes before restarting the peer when the prefix limit is reached.",
				Computed:            true,
			},
		},
	}
}

func (d *BGPPeerTemplateMaxPrefixDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.data = req.ProviderData.(*NxosProviderData)
}

func (d *BGPPeerTemplateMaxPrefixDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config, state BGPPeerTemplateMaxPrefix

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
