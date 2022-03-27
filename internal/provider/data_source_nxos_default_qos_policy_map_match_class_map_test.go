// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosDefaultQOSPolicyMapMatchClassMap(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosDefaultQOSPolicyMapMatchClassMapPrerequisitesConfig + testAccDataSourceNxosDefaultQOSPolicyMapMatchClassMapConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map.test", "name", "Voice"),
				),
			},
		},
	})
}

const testAccDataSourceNxosDefaultQOSPolicyMapMatchClassMapPrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/ipqos/dflt/p/name-[PM1]"
  class_name = "ipqosPMapInst"
  content = {
      name = "PM1"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/ipqos/dflt/c/name-[Voice]"
  class_name = "ipqosCMapInst"
  content = {
      name = "Voice"
  }
}

`

const testAccDataSourceNxosDefaultQOSPolicyMapMatchClassMapConfig = `

resource "nxos_default_qos_policy_map_match_class_map" "test" {
  policy_map_name = "PM1"
  name = "Voice"
  depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, ]
}

data "nxos_default_qos_policy_map_match_class_map" "test" {
  policy_map_name = "PM1"
  name = "Voice"
  depends_on = [nxos_default_qos_policy_map_match_class_map.test]
}
`
