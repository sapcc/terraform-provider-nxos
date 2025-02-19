// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type PhysicalInterfaceVRF struct {
	Device types.String `tfsdk:"device"`
	Dn     types.String `tfsdk:"id"`
	Id     types.String `tfsdk:"interface_id"`
	TDn    types.String `tfsdk:"vrf_dn"`
}

func (data PhysicalInterfaceVRF) getDn() string {
	return fmt.Sprintf("sys/intf/phys-[%s]/rtvrfMbr", data.Id.ValueString())
}

func (data PhysicalInterfaceVRF) getClassName() string {
	return "nwRtVrfMbr"
}

func (data PhysicalInterfaceVRF) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("tDn", data.TDn.ValueString())
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *PhysicalInterfaceVRF) fromBody(res gjson.Result) {
	data.TDn = types.StringValue(res.Get("*.attributes.tDn").String())
}

func (data *PhysicalInterfaceVRF) fromPlan(plan PhysicalInterfaceVRF) {
	data.Device = plan.Device
	data.Dn = plan.Dn
	data.Id = plan.Id
}
