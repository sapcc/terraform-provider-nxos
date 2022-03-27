// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosQueuingQOSPolicySystemOut(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosQueuingQOSPolicySystemOutConfig,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
	})
}

const testAccDataSourceNxosQueuingQOSPolicySystemOutConfig = `

resource "nxos_queuing_qos_policy_system_out" "test" {
}

data "nxos_queuing_qos_policy_system_out" "test" {
  depends_on = [nxos_queuing_qos_policy_system_out.test]
}
`
