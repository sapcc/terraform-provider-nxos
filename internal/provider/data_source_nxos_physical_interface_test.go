// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosPhysicalInterface(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosPhysicalInterfaceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_physical_interface.test", "interface_id", "eth1/10"),
					resource.TestCheckResourceAttr("data.nxos_physical_interface.test", "fec_mode", "auto"),
					resource.TestCheckResourceAttr("data.nxos_physical_interface.test", "access_vlan", "unknown"),
					resource.TestCheckResourceAttr("data.nxos_physical_interface.test", "admin_state", "up"),
					resource.TestCheckResourceAttr("data.nxos_physical_interface.test", "auto_negotiation", "on"),
					resource.TestCheckResourceAttr("data.nxos_physical_interface.test", "bandwidth", "1000"),
					resource.TestCheckResourceAttr("data.nxos_physical_interface.test", "delay", "10"),
					resource.TestCheckResourceAttr("data.nxos_physical_interface.test", "description", "My Description"),
					resource.TestCheckResourceAttr("data.nxos_physical_interface.test", "duplex", "auto"),
					resource.TestCheckResourceAttr("data.nxos_physical_interface.test", "layer", "Layer3"),
					resource.TestCheckResourceAttr("data.nxos_physical_interface.test", "link_logging", "enable"),
					resource.TestCheckResourceAttr("data.nxos_physical_interface.test", "medium", "broadcast"),
					resource.TestCheckResourceAttr("data.nxos_physical_interface.test", "mode", "access"),
					resource.TestCheckResourceAttr("data.nxos_physical_interface.test", "mtu", "1500"),
					resource.TestCheckResourceAttr("data.nxos_physical_interface.test", "native_vlan", "unknown"),
					resource.TestCheckResourceAttr("data.nxos_physical_interface.test", "speed", "auto"),
					resource.TestCheckResourceAttr("data.nxos_physical_interface.test", "speed_group", "auto"),
					resource.TestCheckResourceAttr("data.nxos_physical_interface.test", "trunk_vlans", "1-4094"),
					resource.TestCheckResourceAttr("data.nxos_physical_interface.test", "uni_directional_ethernet", "disable"),
				),
			},
		},
	})
}

const testAccDataSourceNxosPhysicalInterfaceConfig = `

resource "nxos_physical_interface" "test" {
  interface_id = "eth1/10"
  fec_mode = "auto"
  access_vlan = "unknown"
  admin_state = "up"
  auto_negotiation = "on"
  bandwidth = 1000
  delay = 10
  description = "My Description"
  duplex = "auto"
  layer = "Layer3"
  link_logging = "enable"
  medium = "broadcast"
  mode = "access"
  mtu = 1500
  native_vlan = "unknown"
  speed = "auto"
  speed_group = "auto"
  trunk_vlans = "1-4094"
  uni_directional_ethernet = "disable"
}

data "nxos_physical_interface" "test" {
  interface_id = "eth1/10"
  depends_on = [nxos_physical_interface.test]
}
`
