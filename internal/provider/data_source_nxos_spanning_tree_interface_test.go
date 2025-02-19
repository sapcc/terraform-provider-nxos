// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosSpanningTreeInterface(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosSpanningTreeInterfaceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_spanning_tree_interface.test", "interface_id", "eth1/9"),
					resource.TestCheckResourceAttr("data.nxos_spanning_tree_interface.test", "admin_state", "enabled"),
					resource.TestCheckResourceAttr("data.nxos_spanning_tree_interface.test", "bpdu_filter", "enable"),
					resource.TestCheckResourceAttr("data.nxos_spanning_tree_interface.test", "bpdu_guard", "enable"),
					resource.TestCheckResourceAttr("data.nxos_spanning_tree_interface.test", "cost", "100"),
					resource.TestCheckResourceAttr("data.nxos_spanning_tree_interface.test", "ctrl", "bpdu-guard"),
					resource.TestCheckResourceAttr("data.nxos_spanning_tree_interface.test", "guard", "root"),
					resource.TestCheckResourceAttr("data.nxos_spanning_tree_interface.test", "link_type", "p2p"),
					resource.TestCheckResourceAttr("data.nxos_spanning_tree_interface.test", "mode", "edge"),
					resource.TestCheckResourceAttr("data.nxos_spanning_tree_interface.test", "priority", "200"),
				),
			},
		},
	})
}

const testAccDataSourceNxosSpanningTreeInterfaceConfig = `

resource "nxos_spanning_tree_interface" "test" {
  interface_id = "eth1/9"
  admin_state = "enabled"
  bpdu_filter = "enable"
  bpdu_guard = "enable"
  cost = 100
  ctrl = "bpdu-guard"
  guard = "root"
  link_type = "p2p"
  mode = "edge"
  priority = 200
}

data "nxos_spanning_tree_interface" "test" {
  interface_id = "eth1/9"
  depends_on = [nxos_spanning_tree_interface.test]
}
`
