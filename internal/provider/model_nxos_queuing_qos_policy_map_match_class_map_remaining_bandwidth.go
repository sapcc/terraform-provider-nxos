// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type QueuingQOSPolicyMapMatchClassMapRemainingBandwidth struct {
	Device          types.String `tfsdk:"device"`
	Dn              types.String `tfsdk:"id"`
	Policy_map_name types.String `tfsdk:"policy_map_name"`
	Class_map_name  types.String `tfsdk:"class_map_name"`
	Val             types.Int64  `tfsdk:"value"`
}

func (data QueuingQOSPolicyMapMatchClassMapRemainingBandwidth) getDn() string {
	return fmt.Sprintf("sys/ipqos/queuing/p/name-[%s]/cmap-[%s]/setRemBW", data.Policy_map_name.ValueString(), data.Class_map_name.ValueString())
}

func (data QueuingQOSPolicyMapMatchClassMapRemainingBandwidth) getClassName() string {
	return "ipqosSetRemBW"
}

func (data QueuingQOSPolicyMapMatchClassMapRemainingBandwidth) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("val", strconv.FormatInt(data.Val.ValueInt64(), 10))
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *QueuingQOSPolicyMapMatchClassMapRemainingBandwidth) fromBody(res gjson.Result) {
	data.Val = types.Int64Value(res.Get("*.attributes.val").Int())
}

func (data *QueuingQOSPolicyMapMatchClassMapRemainingBandwidth) fromPlan(plan QueuingQOSPolicyMapMatchClassMapRemainingBandwidth) {
	data.Device = plan.Device
	data.Dn = plan.Dn
	data.Policy_map_name = plan.Policy_map_name
	data.Class_map_name = plan.Class_map_name
}
