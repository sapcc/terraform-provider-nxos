// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type FeatureNVOverlay struct {
	Device  types.String `tfsdk:"device"`
	Dn      types.String `tfsdk:"id"`
	AdminSt types.String `tfsdk:"admin_state"`
}

func (data FeatureNVOverlay) getDn() string {
	return "sys/fm/nvo"
}

func (data FeatureNVOverlay) getClassName() string {
	return "fmNvo"
}

func (data FeatureNVOverlay) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("adminSt", data.AdminSt.ValueString())
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *FeatureNVOverlay) fromBody(res gjson.Result) {
	data.AdminSt = types.StringValue(res.Get("*.attributes.adminSt").String())
}

func (data *FeatureNVOverlay) fromPlan(plan FeatureNVOverlay) {
	data.Device = plan.Device
	data.Dn = plan.Dn
}
