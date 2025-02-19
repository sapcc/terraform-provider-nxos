// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type Ethernet struct {
	Device         types.String `tfsdk:"device"`
	Dn             types.String `tfsdk:"id"`
	SystemJumboMtu types.Int64  `tfsdk:"mtu"`
}

func (data Ethernet) getDn() string {
	return "sys/ethpm/inst"
}

func (data Ethernet) getClassName() string {
	return "ethpmInst"
}

func (data Ethernet) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("systemJumboMtu", strconv.FormatInt(data.SystemJumboMtu.ValueInt64(), 10))
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *Ethernet) fromBody(res gjson.Result) {
	data.SystemJumboMtu = types.Int64Value(res.Get("*.attributes.systemJumboMtu").Int())
}

func (data *Ethernet) fromPlan(plan Ethernet) {
	data.Device = plan.Device
	data.Dn = plan.Dn
}
