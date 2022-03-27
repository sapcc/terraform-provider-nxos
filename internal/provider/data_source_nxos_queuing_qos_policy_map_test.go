// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosQueuingQOSPolicyMap(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosQueuingQOSPolicyMapConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_queuing_qos_policy_map.test", "name", "PM1"),
					resource.TestCheckResourceAttr("data.nxos_queuing_qos_policy_map.test", "match_type", "match-any"),
				),
			},
		},
	})
}

const testAccDataSourceNxosQueuingQOSPolicyMapConfig = `

resource "nxos_queuing_qos_policy_map" "test" {
  name = "PM1"
  match_type = "match-any"
}

data "nxos_queuing_qos_policy_map" "test" {
  name = "PM1"
  depends_on = [nxos_queuing_qos_policy_map.test]
}
`
