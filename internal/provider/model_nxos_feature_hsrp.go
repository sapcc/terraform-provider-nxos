// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type FeatureHSRP struct {
	Device  types.String `tfsdk:"device"`
	Dn      types.String `tfsdk:"id"`
	AdminSt types.String `tfsdk:"admin_state"`
}

func (data FeatureHSRP) getDn() string {
	return "sys/fm/hsrp"
}

func (data FeatureHSRP) getClassName() string {
	return "fmHsrp"
}

func (data FeatureHSRP) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("adminSt", data.AdminSt.ValueString())
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *FeatureHSRP) fromBody(res gjson.Result) {
	data.AdminSt = types.StringValue(res.Get("*.attributes.adminSt").String())
}

func (data *FeatureHSRP) fromPlan(plan FeatureHSRP) {
	data.Device = plan.Device
	data.Dn = plan.Dn
}
