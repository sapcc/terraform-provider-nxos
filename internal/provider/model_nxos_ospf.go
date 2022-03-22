// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type OSPF struct {
	Dn types.String `tfsdk:"id"`
	AdminSt types.String `tfsdk:"admin_state"`
}

func (data OSPF) getDn() string {
	return "sys/ospf"
}

func (data OSPF) getClassName() string {
	return "ospfEntity"
}

func (data OSPF) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("adminSt", data.AdminSt.Value)
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *OSPF) fromBody(res gjson.Result) {
	data.Dn.Value = res.Get("*.attributes.dn").String()
	data.AdminSt.Value = res.Get("*.attributes.adminSt").String()
}

func (data *OSPF) fromPlan(plan OSPF) {
}
