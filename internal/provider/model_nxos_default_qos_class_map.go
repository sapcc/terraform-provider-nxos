// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type DefaultQOSClassMap struct {
	Device    types.String `tfsdk:"device"`
	Dn        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"`
	MatchType types.String `tfsdk:"match_type"`
}

func (data DefaultQOSClassMap) getDn() string {
	return fmt.Sprintf("sys/ipqos/dflt/c/name-[%s]", data.Name.ValueString())
}

func (data DefaultQOSClassMap) getClassName() string {
	return "ipqosCMapInst"
}

func (data DefaultQOSClassMap) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("name", data.Name.ValueString()).
		Set("matchType", data.MatchType.ValueString())
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *DefaultQOSClassMap) fromBody(res gjson.Result) {
	data.Name = types.StringValue(res.Get("*.attributes.name").String())
	data.MatchType = types.StringValue(res.Get("*.attributes.matchType").String())
}

func (data *DefaultQOSClassMap) fromPlan(plan DefaultQOSClassMap) {
	data.Device = plan.Device
	data.Dn = plan.Dn
}
