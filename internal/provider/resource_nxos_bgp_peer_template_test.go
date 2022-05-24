// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNxosBGPPeerTemplate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosBGPPeerTemplatePrerequisitesConfig + testAccNxosBGPPeerTemplateConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_bgp_peer_template.test", "asn", "65001"),
					resource.TestCheckResourceAttr("nxos_bgp_peer_template.test", "template_name", "SPINE-PEERS"),
					resource.TestCheckResourceAttr("nxos_bgp_peer_template.test", "remote_asn", "65002"),
					resource.TestCheckResourceAttr("nxos_bgp_peer_template.test", "description", "My Description"),
					resource.TestCheckResourceAttr("nxos_bgp_peer_template.test", "peer_type", "fabric-internal"),
					resource.TestCheckResourceAttr("nxos_bgp_peer_template.test", "source_interface", "lo0"),
				),
			},
			{
				ResourceName:  "nxos_bgp_peer_template.test",
				ImportState:   true,
				ImportStateId: "sys/bgp/inst/dom-[default]/peercont-[SPINE-PEERS]",
			},
		},
	})
}

const testAccNxosBGPPeerTemplatePrerequisitesConfig = `
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

func testAccNxosBGPPeerTemplateConfig_minimum() string {
	return `
	resource "nxos_bgp_peer_template" "test" {
		asn = "65001"
		template_name = "SPINE-PEERS"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, ]
	}
	`
}

func testAccNxosBGPPeerTemplateConfig_all() string {
	return `
	resource "nxos_bgp_peer_template" "test" {
		asn = "65001"
		template_name = "SPINE-PEERS"
		remote_asn = "65002"
		description = "My Description"
		peer_type = "fabric-internal"
		source_interface = "lo0"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, ]
	}
	`
}
