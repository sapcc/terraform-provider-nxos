// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNxosDefaultQOSClassMap(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosDefaultQOSClassMapConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_default_qos_class_map.test", "name", "Voice"),
					resource.TestCheckResourceAttr("nxos_default_qos_class_map.test", "match_type", "match-any"),
				),
			},
			{
				ResourceName:  "nxos_default_qos_class_map.test",
				ImportState:   true,
				ImportStateId: "sys/ipqos/dflt/c/name-[Voice]",
			},
		},
	})
}

func testAccNxosDefaultQOSClassMapConfig_minimum() string {
	return `
	resource "nxos_default_qos_class_map" "test" {
		name = "Voice"
	}
	`
}

func testAccNxosDefaultQOSClassMapConfig_all() string {
	return `
	resource "nxos_default_qos_class_map" "test" {
		name = "Voice"
		match_type = "match-any"
	}
	`
}
