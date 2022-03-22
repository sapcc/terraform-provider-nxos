//go:build ignore
// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNxos{{camelCase .Name}}(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{{- $configs := ""}}
			{{- if len .Parents}}
			{{- $configs = printf "%s%s%s" "testAccNxos" (camelCase (index .Parents 0)) "Config_all()"}}
			{{- end}}
			{{- range $index, $item := .Parents}}
			{{- if ge $index 1}}{{$configs = printf "%s%s%s%s" $configs "+testAccNxos" (camelCase .) "Config_all()"}}{{end}}
			{
				Config: {{$configs}},
			},
			{{- end}}
			{
				Config: {{- range .Parents}}testAccNxos{{camelCase .}}Config_all()+{{end}}testAccNxos{{camelCase .Name}}Config_minimum(),
				Check: resource.ComposeTestCheckFunc(
					{{- $name := .Name }}
					{{- range  .Attributes}}
					{{- if or (eq .Id true) (eq .Mandatory true)}}
					resource.TestCheckResourceAttr("nxos_{{snakeCase $name}}.test", "{{.TfName}}", "{{.Example}}"),
					{{- end}}
					{{- end}}
				),
			},
			{
				Config: {{- range .Parents}}testAccNxos{{camelCase .}}Config_all()+{{end}}testAccNxos{{camelCase .Name}}Config_all(),
				Check: resource.ComposeTestCheckFunc(
					{{- $name := .Name }}
					{{- range  .Attributes}}
					resource.TestCheckResourceAttr("nxos_{{snakeCase $name}}.test", "{{.TfName}}", "{{.Example}}"),
					{{- end}}
				),
			},
			{
				ResourceName:  "nxos_{{snakeCase $name}}.test",
				ImportState:   true,
				ImportStateId: "{{getExampleDn .Dn .Attributes}}",
			},
		},
	})
}

func testAccNxos{{camelCase .Name}}Config_minimum() string {
	return `
	resource "nxos_{{snakeCase $name}}" "test" {
	{{- range  .Attributes}}
	{{- if or (eq .Id true) (eq .Mandatory true)}}
		{{.TfName}} = {{if eq .Type "String"}}"{{end}}{{.Example}}{{if eq .Type "String"}}"{{end}}
	{{- end}}
	{{- end}}
	}
	`
}

func testAccNxos{{camelCase .Name}}Config_all() string {
	return `
	resource "nxos_{{snakeCase $name}}" "test" {
	{{- range  .Attributes}}
		{{.TfName}} = {{if eq .Type "String"}}"{{end}}{{.Example}}{{if eq .Type "String"}}"{{end}}
	{{- end}}
	}
	`
}
