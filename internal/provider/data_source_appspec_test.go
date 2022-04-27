package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceAppspec(t *testing.T) {
	// t.Skip("data source not yet implemented, remove this once you add your own code")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceAppspec,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("data.dorentest_appspec.foo", "raw_appspec", regexp.MustCompile("---")),
					// resource.TestMatchResourceAttr("data.dorentest_appspec.foo", "function_name", regexp.MustCompile("some-function-name")),
					// resource.TestMatchTypeSetElemNestedAttrs("data.dorentest_appspec.foo", "lambdas.*", map[string]*regexp.Regexp{
					// 	"function_name": regexp.MustCompile("some-function-name"),
					// }),
				),
			},
		},
	})
}

const testAccDataSourceAppspec = `
data "dorentest_appspec" "foo" {
  raw_appspec = <<EOT
---
lambdas:
  version: 1.0.0
  lambda_functions:
    - function: "some-function-name"
      image_uri: 395127396906.dkr.ecr.us-west-2.amazonaws.com/chp-test-r635542-demo-lambda-container:latest
EOT
}
`
