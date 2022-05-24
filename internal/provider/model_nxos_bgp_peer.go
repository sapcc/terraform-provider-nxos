// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type BGPPeer struct {
	Device   types.String `tfsdk:"device"`
	Dn       types.String `tfsdk:"id"`
	Bgp_asn  types.String `tfsdk:"asn"`
	Vrf_name types.String `tfsdk:"vrf"`
	Addr     types.String `tfsdk:"address"`
	Asn      types.String `tfsdk:"remote_asn"`
	Name     types.String `tfsdk:"description"`
	PeerImp  types.String `tfsdk:"peer_template"`
	PeerType types.String `tfsdk:"peer_type"`
	SrcIf    types.String `tfsdk:"source_interface"`
}

func (data BGPPeer) getDn() string {
	return fmt.Sprintf("sys/bgp/inst/dom-[%s]/peer-[%s]", data.Vrf_name.Value, data.Addr.Value)
}

func (data BGPPeer) getClassName() string {
	return "bgpPeer"
}

func (data BGPPeer) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("addr", data.Addr.Value).
		Set("asn", data.Asn.Value).
		Set("name", data.Name.Value).
		Set("peerImp", data.PeerImp.Value).
		Set("peerType", data.PeerType.Value).
		Set("srcIf", data.SrcIf.Value)
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *BGPPeer) fromBody(res gjson.Result) {
	data.Addr.Value = res.Get("*.attributes.addr").String()
	data.Asn.Value = res.Get("*.attributes.asn").String()
	data.Name.Value = res.Get("*.attributes.name").String()
	data.PeerImp.Value = res.Get("*.attributes.peerImp").String()
	data.PeerType.Value = res.Get("*.attributes.peerType").String()
	data.SrcIf.Value = res.Get("*.attributes.srcIf").String()
}

func (data *BGPPeer) fromPlan(plan BGPPeer) {
	data.Device = plan.Device
	data.Dn.Value = plan.Dn.Value
	data.Bgp_asn.Value = plan.Bgp_asn.Value
	data.Vrf_name.Value = plan.Vrf_name.Value
}
