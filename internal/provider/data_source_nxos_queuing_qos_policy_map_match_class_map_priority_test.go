// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosQueuingQOSPolicyMapMatchClassMapPriority(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosQueuingQOSPolicyMapMatchClassMapPriorityPrerequisitesConfig + testAccDataSourceNxosQueuingQOSPolicyMapMatchClassMapPriorityConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_queuing_qos_policy_map_match_class_map_priority.test", "level", "1"),
				),
			},
		},
	})
}

const testAccDataSourceNxosQueuingQOSPolicyMapMatchClassMapPriorityPrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/ipqos/queuing/p/name-[PM1]"
  class_name = "ipqosPMapInst"
  content = {
      name = "PM1"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/ipqos/queuing/p/name-[PM1]/cmap-[c-out-q1]"
  class_name = "ipqosMatchCMap"
  content = {
      name = "c-out-q1"
  }
  depends_on = [nxos_rest.PreReq0, ]
}

`

const testAccDataSourceNxosQueuingQOSPolicyMapMatchClassMapPriorityConfig = `

resource "nxos_queuing_qos_policy_map_match_class_map_priority" "test" {
  policy_map_name = "PM1"
  class_map_name = "c-out-q1"
  level = 1
  depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, ]
}

data "nxos_queuing_qos_policy_map_match_class_map_priority" "test" {
  policy_map_name = "PM1"
  class_map_name = "c-out-q1"
  depends_on = [nxos_queuing_qos_policy_map_match_class_map_priority.test]
}
`
