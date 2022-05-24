// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNxosBGPRouteControl(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosBGPRouteControlPrerequisitesConfig + testAccNxosBGPRouteControlConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_bgp_route_control.test", "asn", "65001"),
					resource.TestCheckResourceAttr("nxos_bgp_route_control.test", "vrf", "default"),
					resource.TestCheckResourceAttr("nxos_bgp_route_control.test", "enforce_first_as", "disabled"),
					resource.TestCheckResourceAttr("nxos_bgp_route_control.test", "fib_accelerate", "enabled"),
					resource.TestCheckResourceAttr("nxos_bgp_route_control.test", "log_neighbor_changes", "enabled"),
					resource.TestCheckResourceAttr("nxos_bgp_route_control.test", "suppress_routes", "disabled"),
				),
			},
			{
				ResourceName:  "nxos_bgp_route_control.test",
				ImportState:   true,
				ImportStateId: "sys/bgp/inst/dom-[default]/rtctrl",
			},
		},
	})
}

const testAccNxosBGPRouteControlPrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/fm/bgp"
  class_name = "fmBgp"
  delete = false
  content = {
      adminSt = "enabled"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/bgp"
  class_name = "bgpEntity"
  content = {
  }
  depends_on = [nxos_rest.PreReq0, ]
}

resource "nxos_rest" "PreReq2" {
  dn = "sys/bgp/inst"
  class_name = "bgpInst"
  content = {
      adminSt = "enabled"
      asn = "65001"
  }
  depends_on = [nxos_rest.PreReq1, ]
}

resource "nxos_rest" "PreReq3" {
  dn = "sys/bgp/inst/dom-[default]"
  class_name = "bgpDom"
  content = {
      name = "default"
  }
  depends_on = [nxos_rest.PreReq2, ]
}

`

func testAccNxosBGPRouteControlConfig_minimum() string {
	return `
	resource "nxos_bgp_route_control" "test" {
		vrf = "default"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, ]
	}
	`
}

func testAccNxosBGPRouteControlConfig_all() string {
	return `
	resource "nxos_bgp_route_control" "test" {
		asn = "65001"
		vrf = "default"
		enforce_first_as = "disabled"
		fib_accelerate = "enabled"
		log_neighbor_changes = "enabled"
		suppress_routes = "disabled"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, ]
	}
	`
}
