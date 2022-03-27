// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosLoopbackInterface(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosLoopbackInterfaceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_loopback_interface.test", "interface_id", "lo123"),
					resource.TestCheckResourceAttr("data.nxos_loopback_interface.test", "admin_state", "down"),
					resource.TestCheckResourceAttr("data.nxos_loopback_interface.test", "description", "My Description"),
				),
			},
		},
	})
}

const testAccDataSourceNxosLoopbackInterfaceConfig = `

resource "nxos_loopback_interface" "test" {
  interface_id = "lo123"
  admin_state = "down"
  description = "My Description"
}

data "nxos_loopback_interface" "test" {
  interface_id = "lo123"
  depends_on = [nxos_loopback_interface.test]
}
`
