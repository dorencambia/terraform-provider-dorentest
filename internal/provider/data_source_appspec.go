package provider

import (
	"context"

	"github.com/dorencambia/terraform-provider-dorentest/internal/appspec"
	"github.com/dorencambia/terraform-provider-dorentest/internal/appspec/lambda"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAppspec() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Pass in the content of a juno_appspec.yml file and get modified data back",

		ReadContext: dataSourceAppspecRead,

		Schema: map[string]*schema.Schema{
			"raw_appspec": {
				// This description is used by the documentation generator and the language server.
				Description: "Raw text of juno_appspec.yml file",
				Type:        schema.TypeString,
				Required:    true,
			},
			"lambdas": {
				Computed: true,
				Type:     schema.TypeList,
				Elem:     lambda.LambdaResource,
				// Elem:     lambda.Schema(),
			},
		},
	}
}

func dataSourceAppspecRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	raw_appspec := d.Get("raw_appspec").(string)
	// func ParseAppspec(input string) (Appspec, error) {
	spec, err := appspec.ParseAppspec(raw_appspec)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("lambdas", lambda.Converter(spec.Lambdas.LambdaFunctions))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("appspecid")
	return nil
}

// func dataSourceAppspec() *schema.Resource {
// 	return &schema.Resource{
// 		// This description is used by the documentation generator and the language server.
// 		Description: "Pass in the content of a juno_appspec.yml file and get modified data back",

// 		ReadContext: dataSourceAppspecRead,

// 		Schema: map[string]*schema.Schema{
// 			"raw_appspec": {
// 				// This description is used by the documentation generator and the language server.
// 				Description: "Raw text of juno_appspec.yml file",
// 				Type:        schema.TypeString,
// 				Required:    true,
// 			},
// 			"appspec": {
// 				// This description is used by the documentation generator and the language server.
// 				Description: "Map of appspec",
// 				Type:        schema.TypeMap,
// 				Computed:    true,
// 			},
// 			"appspec_json_string": {
// 				// This description is used by the documentation generator and the language server.
// 				Description: "Raw text of juno_appspec.yml file",
// 				Type:        schema.TypeString,
// 				Computed:    true,
// 			},
// 			"return_map": {
// 				// This description is used by the documentation generator and the language server.
// 				Description: "Whether or not to try to return the map",
// 				Type:        schema.TypeBool,
// 				Default:     true,
// 				Optional:    true,
// 			},
// 		},
// 	}
// }

// func dataSourceAppspecRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
// 	type T struct{}
// 	appspec := T{}
// 	data := d.Get("raw_appspec").(string)
// 	err := yaml.Unmarshal([]byte(data), &appspec)
// 	if err != nil {
// 		return diag.FromErr(err)
// 	}

// 	d.SetId("some id")
// 	d.Set("raw_appspec", "This is null in the response if I do not set it here?")

// 	bytes, err := json.Marshal(appspec)
// 	if err != nil {
// 		fmt.Println("error:", err)
// 	}
// 	err = d.Set("appspec_json_string", string(bytes))
// 	if err != nil {
// 		fmt.Println("error:", err)
// 	}

// 	if !d.Get("return_map").(bool) {
// 		return nil
// 	}

// 	err = d.Set("appspec", appspec)
// 	if err != nil {
// 		return diag.FromErr(err)
// 	}
// 	return nil

// 	// use the meta value to retrieve your client from the provider configure method
// 	// client := meta.(*apiClient)

// 	// idFromAPI := "my-id"
// 	// d.SetId(idFromAPI)

// 	// return diag.Errorf("not implemented")
// }
