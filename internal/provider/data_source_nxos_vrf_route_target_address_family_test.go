// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosVRFRouteTargetAddressFamily(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosVRFRouteTargetAddressFamilyPrerequisitesConfig + testAccDataSourceNxosVRFRouteTargetAddressFamilyConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_vrf_route_target_address_family.test", "route_target_address_family", "ipv4-ucast"),
				),
			},
		},
	})
}

const testAccDataSourceNxosVRFRouteTargetAddressFamilyPrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/inst-[VRF1]"
  class_name = "l3Inst"
  content = {
      name = "VRF1"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/ipv4/inst/dom-[VRF1]"
  class_name = "ipv4Dom"
  content = {
      name = "VRF1"
  }
}

resource "nxos_rest" "PreReq2" {
  dn = "sys/inst-[VRF1]/dom-[VRF1]"
  class_name = "rtctrlDom"
  content = {
  }
  depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, ]
}

resource "nxos_rest" "PreReq3" {
  dn = "sys/inst-[VRF1]/dom-[VRF1]/af-[ipv4-ucast]"
  class_name = "rtctrlDomAf"
  content = {
      type = "ipv4-ucast"
  }
  depends_on = [nxos_rest.PreReq2, ]
}

`

const testAccDataSourceNxosVRFRouteTargetAddressFamilyConfig = `

resource "nxos_vrf_route_target_address_family" "test" {
  vrf = "VRF1"
  address_family = "ipv4-ucast"
  route_target_address_family = "ipv4-ucast"
  depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, ]
}

data "nxos_vrf_route_target_address_family" "test" {
  vrf = "VRF1"
  address_family = "ipv4-ucast"
  route_target_address_family = "ipv4-ucast"
  depends_on = [nxos_vrf_route_target_address_family.test]
}
`
