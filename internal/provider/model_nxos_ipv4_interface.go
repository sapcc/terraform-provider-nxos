// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type IPv4Interface struct {
	Device     types.String `tfsdk:"device"`
	Dn         types.String `tfsdk:"id"`
	Dom        types.String `tfsdk:"vrf"`
	Id         types.String `tfsdk:"interface_id"`
	DropGlean  types.String `tfsdk:"drop_glean"`
	Forward    types.String `tfsdk:"forward"`
	Unnumbered types.String `tfsdk:"unnumbered"`
	Urpf       types.String `tfsdk:"urpf"`
}

func (data IPv4Interface) getDn() string {
	return fmt.Sprintf("sys/ipv4/inst/dom-[%s]/if-[%s]", data.Dom.ValueString(), data.Id.ValueString())
}

func (data IPv4Interface) getClassName() string {
	return "ipv4If"
}

func (data IPv4Interface) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("id", data.Id.ValueString()).
		Set("dropGlean", data.DropGlean.ValueString()).
		Set("forward", data.Forward.ValueString()).
		Set("unnumbered", data.Unnumbered.ValueString()).
		Set("urpf", data.Urpf.ValueString())
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *IPv4Interface) fromBody(res gjson.Result) {
	data.Id = types.StringValue(res.Get("*.attributes.id").String())
	data.DropGlean = types.StringValue(res.Get("*.attributes.dropGlean").String())
	data.Forward = types.StringValue(res.Get("*.attributes.forward").String())
	data.Unnumbered = types.StringValue(res.Get("*.attributes.unnumbered").String())
	data.Urpf = types.StringValue(res.Get("*.attributes.urpf").String())
}

func (data *IPv4Interface) fromPlan(plan IPv4Interface) {
	data.Device = plan.Device
	data.Dn = plan.Dn
	data.Dom = plan.Dom
}
