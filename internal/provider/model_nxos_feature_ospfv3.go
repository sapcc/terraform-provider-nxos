// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type FeatureOSPFv3 struct {
	Device  types.String `tfsdk:"device"`
	Dn      types.String `tfsdk:"id"`
	AdminSt types.String `tfsdk:"admin_state"`
}

func (data FeatureOSPFv3) getDn() string {
	return "sys/fm/ospfv3"
}

func (data FeatureOSPFv3) getClassName() string {
	return "fmOspfv3"
}

func (data FeatureOSPFv3) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("adminSt", data.AdminSt.ValueString())
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *FeatureOSPFv3) fromBody(res gjson.Result) {
	data.AdminSt = types.StringValue(res.Get("*.attributes.adminSt").String())
}

func (data *FeatureOSPFv3) fromPlan(plan FeatureOSPFv3) {
	data.Device = plan.Device
	data.Dn = plan.Dn
}
