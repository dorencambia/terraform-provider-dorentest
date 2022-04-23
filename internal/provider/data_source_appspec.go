package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gopkg.in/yaml.v2"
)

func dataSourceAppspec() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample data source in the Terraform provider scaffolding.",

		ReadContext: dataSourceAppspecRead,

		Schema: map[string]*schema.Schema{
			"raw_appspec": {
				// This description is used by the documentation generator and the language server.
				Description: "Raw text of juno_appspec.yml file",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func dataSourceAppspecRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	type T struct{}
	appspec := T{}
	data := d.Get("raw_appspec").(string)
	err := yaml.Unmarshal([]byte(data), &appspec)
	if err != nil {
		return diag.FromErr(err)
		// log.Fatalf("error: %v", err)
	}

	d.Set("appspec", appspec)
	return nil

	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	// idFromAPI := "my-id"
	// d.SetId(idFromAPI)

	// return diag.Errorf("not implemented")
}
