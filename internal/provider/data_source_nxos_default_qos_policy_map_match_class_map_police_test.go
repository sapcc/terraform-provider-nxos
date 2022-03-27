// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosDefaultQOSPolicyMapMatchClassMapPolice(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosDefaultQOSPolicyMapMatchClassMapPolicePrerequisitesConfig + testAccDataSourceNxosDefaultQOSPolicyMapMatchClassMapPoliceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_police.test", "bc_rate", "200"),
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_police.test", "bc_unit", "mbytes"),
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_police.test", "be_rate", "200"),
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_police.test", "be_unit", "mbytes"),
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_police.test", "cir_rate", "10000"),
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_police.test", "cir_unit", "mbps"),
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_police.test", "conform_action", "transmit"),
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_police.test", "conform_set_cos", "0"),
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_police.test", "conform_set_dscp", "0"),
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_police.test", "conform_set_precedence", "routine"),
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_police.test", "conform_set_qos_group", "0"),
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_police.test", "exceed_action", "transmit"),
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_police.test", "exceed_set_cos", "0"),
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_police.test", "exceed_set_dscp", "0"),
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_police.test", "exceed_set_precedence", "routine"),
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_police.test", "exceed_set_qos_group", "0"),
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_police.test", "pir_rate", "10000"),
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_police.test", "pir_unit", "mbps"),
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_police.test", "violate_action", "drop"),
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_police.test", "violate_set_cos", "0"),
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_police.test", "violate_set_dscp", "0"),
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_police.test", "violate_set_precedence", "routine"),
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_police.test", "violate_set_qos_group", "0"),
				),
			},
		},
	})
}

const testAccDataSourceNxosDefaultQOSPolicyMapMatchClassMapPolicePrerequisitesConfig = `
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

resource "nxos_rest" "PreReq2" {
  dn = "sys/ipqos/dflt/p/name-[PM1]/cmap-[Voice]"
  class_name = "ipqosMatchCMap"
  content = {
      name = "Voice"
  }
  depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, ]
}

`

const testAccDataSourceNxosDefaultQOSPolicyMapMatchClassMapPoliceConfig = `

resource "nxos_default_qos_policy_map_match_class_map_police" "test" {
  policy_map_name = "PM1"
  class_map_name = "Voice"
  bc_rate = 200
  bc_unit = "mbytes"
  be_rate = 200
  be_unit = "mbytes"
  cir_rate = 10000
  cir_unit = "mbps"
  conform_action = "transmit"
  conform_set_cos = 0
  conform_set_dscp = 0
  conform_set_precedence = "routine"
  conform_set_qos_group = 0
  exceed_action = "transmit"
  exceed_set_cos = 0
  exceed_set_dscp = 0
  exceed_set_precedence = "routine"
  exceed_set_qos_group = 0
  pir_rate = 10000
  pir_unit = "mbps"
  violate_action = "drop"
  violate_set_cos = 0
  violate_set_dscp = 0
  violate_set_precedence = "routine"
  violate_set_qos_group = 0
  depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, ]
}

data "nxos_default_qos_policy_map_match_class_map_police" "test" {
  policy_map_name = "PM1"
  class_map_name = "Voice"
  depends_on = [nxos_default_qos_policy_map_match_class_map_police.test]
}
`
