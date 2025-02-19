// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type DefaultQOSPolicyMapMatchClassMapSetQOSGroup struct {
	Device          types.String `tfsdk:"device"`
	Dn              types.String `tfsdk:"id"`
	Policy_map_name types.String `tfsdk:"policy_map_name"`
	Class_map_name  types.String `tfsdk:"class_map_name"`
	Id              types.Int64  `tfsdk:"qos_group_id"`
}

func (data DefaultQOSPolicyMapMatchClassMapSetQOSGroup) getDn() string {
	return fmt.Sprintf("sys/ipqos/dflt/p/name-[%s]/cmap-[%s]/setGrp", data.Policy_map_name.ValueString(), data.Class_map_name.ValueString())
}

func (data DefaultQOSPolicyMapMatchClassMapSetQOSGroup) getClassName() string {
	return "ipqosSetQoSGrp"
}

func (data DefaultQOSPolicyMapMatchClassMapSetQOSGroup) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("id", strconv.FormatInt(data.Id.ValueInt64(), 10))
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *DefaultQOSPolicyMapMatchClassMapSetQOSGroup) fromBody(res gjson.Result) {
	data.Id = types.Int64Value(res.Get("*.attributes.id").Int())
}

func (data *DefaultQOSPolicyMapMatchClassMapSetQOSGroup) fromPlan(plan DefaultQOSPolicyMapMatchClassMapSetQOSGroup) {
	data.Device = plan.Device
	data.Dn = plan.Dn
	data.Policy_map_name = plan.Policy_map_name
	data.Class_map_name = plan.Class_map_name
}
