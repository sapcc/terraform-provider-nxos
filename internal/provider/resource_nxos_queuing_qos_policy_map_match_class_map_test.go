// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNxosQueuingQOSPolicyMapMatchClassMap(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosQueuingQOSPolicyMapMatchClassMapPrerequisitesConfig + testAccNxosQueuingQOSPolicyMapMatchClassMapConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_queuing_qos_policy_map_match_class_map.test", "policy_map_name", "PM1"),
					resource.TestCheckResourceAttr("nxos_queuing_qos_policy_map_match_class_map.test", "name", "c-out-q1"),
				),
			},
			{
				ResourceName:  "nxos_queuing_qos_policy_map_match_class_map.test",
				ImportState:   true,
				ImportStateId: "sys/ipqos/queuing/p/name-[PM1]/cmap-[c-out-q1]",
			},
		},
	})
}

const testAccNxosQueuingQOSPolicyMapMatchClassMapPrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/ipqos/queuing/p/name-[PM1]"
  class_name = "ipqosPMapInst"
  content = {
      name = "PM1"
  }
}

`

func testAccNxosQueuingQOSPolicyMapMatchClassMapConfig_minimum() string {
	return `
	resource "nxos_queuing_qos_policy_map_match_class_map" "test" {
		policy_map_name = "PM1"
		name = "c-out-q1"
  		depends_on = [nxos_rest.PreReq0, ]
	}
	`
}

func testAccNxosQueuingQOSPolicyMapMatchClassMapConfig_all() string {
	return `
	resource "nxos_queuing_qos_policy_map_match_class_map" "test" {
		policy_map_name = "PM1"
		name = "c-out-q1"
  		depends_on = [nxos_rest.PreReq0, ]
	}
	`
}
