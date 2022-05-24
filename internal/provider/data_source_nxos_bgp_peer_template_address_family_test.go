// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosBGPPeerTemplateAddressFamily(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosBGPPeerTemplateAddressFamilyPrerequisitesConfig + testAccDataSourceNxosBGPPeerTemplateAddressFamilyConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_bgp_peer_template_address_family.test", "address_family", "ipv4-ucast"),
					resource.TestCheckResourceAttr("data.nxos_bgp_peer_template_address_family.test", "control", "nh-self,rr-client"),
					resource.TestCheckResourceAttr("data.nxos_bgp_peer_template_address_family.test", "send_community_extended", "enabled"),
					resource.TestCheckResourceAttr("data.nxos_bgp_peer_template_address_family.test", "send_community_standard", "enabled"),
				),
			},
		},
	})
}

const testAccDataSourceNxosBGPPeerTemplateAddressFamilyPrerequisitesConfig = `
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

resource "nxos_rest" "PreReq4" {
  dn = "sys/bgp/inst/dom-[default]/peercont-[SPINE-PEERS]"
  class_name = "bgpPeerCont"
  content = {
      name = "SPINE-PEERS"
  }
  depends_on = [nxos_rest.PreReq3, ]
}

`

const testAccDataSourceNxosBGPPeerTemplateAddressFamilyConfig = `

resource "nxos_bgp_peer_template_address_family" "test" {
  asn = "65001"
  template_name = "SPINE-PEERS"
  address_family = "ipv4-ucast"
  control = "nh-self,rr-client"
  send_community_extended = "enabled"
  send_community_standard = "enabled"
  depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, nxos_rest.PreReq4, ]
}

data "nxos_bgp_peer_template_address_family" "test" {
  template_name = "SPINE-PEERS"
  address_family = "ipv4-ucast"
  depends_on = [nxos_bgp_peer_template_address_family.test]
}
`
