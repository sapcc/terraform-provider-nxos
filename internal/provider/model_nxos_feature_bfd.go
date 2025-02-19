// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type FeatureBFD struct {
	Device  types.String `tfsdk:"device"`
	Dn      types.String `tfsdk:"id"`
	AdminSt types.String `tfsdk:"admin_state"`
}

func (data FeatureBFD) getDn() string {
	return "sys/fm/bfd"
}

func (data FeatureBFD) getClassName() string {
	return "fmBfd"
}

func (data FeatureBFD) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("adminSt", data.AdminSt.ValueString())
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *FeatureBFD) fromBody(res gjson.Result) {
	data.AdminSt = types.StringValue(res.Get("*.attributes.adminSt").String())
}

func (data *FeatureBFD) fromPlan(plan FeatureBFD) {
	data.Device = plan.Device
	data.Dn = plan.Dn
}
