// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosPIMSSMRange(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosPIMConfig_all(),
			},
			{
				Config: testAccNxosPIMConfig_all() + testAccNxosPIMInstanceConfig_all(),
			},
			{
				Config: testAccNxosPIMConfig_all() + testAccNxosPIMInstanceConfig_all() + testAccNxosPIMVRFConfig_all(),
			},
			{
				Config: testAccNxosPIMConfig_all() + testAccNxosPIMInstanceConfig_all() + testAccNxosPIMVRFConfig_all() + testAccNxosPIMSSMPolicyConfig_all(),
			},
			{
				Config: testAccNxosPIMConfig_all() + testAccNxosPIMInstanceConfig_all() + testAccNxosPIMVRFConfig_all() + testAccNxosPIMSSMPolicyConfig_all() + testAccNxosPIMSSMRangeConfig_all(),
			},
			{
				Config: testAccDataSourceNxosPIMSSMRangeConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_pim_ssm_range.test", "group_list_1", "232.0.0.0/8"),
					resource.TestCheckResourceAttr("data.nxos_pim_ssm_range.test", "group_list_2", "233.0.0.0/8"),
					resource.TestCheckResourceAttr("data.nxos_pim_ssm_range.test", "group_list_3", "0.0.0.0"),
					resource.TestCheckResourceAttr("data.nxos_pim_ssm_range.test", "group_list_4", "0.0.0.0"),
					resource.TestCheckResourceAttr("data.nxos_pim_ssm_range.test", "prefix_list", ""),
					resource.TestCheckResourceAttr("data.nxos_pim_ssm_range.test", "route_map", ""),
					resource.TestCheckResourceAttr("data.nxos_pim_ssm_range.test", "ssm_none", "false"),
				),
			},
		},
	})
}

const testAccDataSourceNxosPIMSSMRangeConfig = `
data "nxos_pim_ssm_range" "test" {
  vrf_name = "default"
}
`
