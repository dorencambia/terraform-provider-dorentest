package lambda

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var LambdaResource = &schema.Resource{
	// This description is used by the documentation generator and the language server.
	Description: "List of lambdas meant to be used with juno modules",

	// ReadContext: dataSourceAppspecRead,

	Schema: LambdaSchema,
}

var LambdaSchema = map[string]*schema.Schema{

	"cloudwatch_event_rule": {
		Optional: true,
		Type:     schema.TypeMap,
		Elem:     &schema.Schema{Type: schema.TypeString},
		// Type:     schema.TypeList,
		// MaxItems: 1,
		// Elem: &schema.Resource{
		// 	Schema: map[string]*schema.Schema{
		// 		"description": {
		// 			Optional: true,
		// 			Type:     schema.TypeString,
		// 		},
		// 		"input": {
		// 			Optional: true,
		// 			Type:     schema.TypeList,
		// 			// MaxItems: 1,
		// 			Elem: &schema.Resource{
		// 				Schema: map[string]*schema.Schema{
		// 					"env": {
		// 						Optional: true,
		// 						Type:     schema.TypeString,
		// 					},
		// 					"id": {
		// 						Optional: true,
		// 						Type:     schema.TypeString,
		// 					},
		// 					"source": {
		// 						Optional: true,
		// 						Type:     schema.TypeString,
		// 					},
		// 				},
		// 			},
		// 		},
		// 		"name": {
		// 			Optional: true,
		// 			Type:     schema.TypeString,
		// 		},
		// 		"schedule_expression": {
		// 			Optional: true,
		// 			Type:     schema.TypeString,
		// 		},
		// 	},
		// },
	},
	"concurrency_limit": {
		Optional: true,
		Type:     schema.TypeInt,
	},
	"description": {
		Optional: true,
		Type:     schema.TypeString,
	},
	"egress_rules": {
		Optional: true,
		Type:     schema.TypeSet,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"cidr_ip": {
					Optional: true,
					Type:     schema.TypeString,
				},
				"ports": {
					Optional: true,
					Type:     schema.TypeSet,
					Elem:     &schema.Schema{Type: schema.TypeInt},
				},
				"proto": {
					Optional: true,
					Type:     schema.TypeString,
				},
				"rule_desc": {
					Optional: true,
					Type:     schema.TypeString,
				},
			},
		},
	},
	"environment_variables": {
		Optional: true,
		Type:     schema.TypeMap,
		Elem:     &schema.Schema{Type: schema.TypeString},
		// Type:     schema.TypeList,
		// MaxItems: 1,
		// Elem: &schema.Resource{
		// 	Schema: map[string]*schema.Schema{
		// 		"testvar1": {
		// 			Optional: true,
		// 			Type:     schema.TypeString,
		// 		},
		// 		"testvar2": {
		// 			Optional: true,
		// 			Type:     schema.TypeString,
		// 		},
		// 	},
		// },
	},
	"function_name": {
		Optional: true,
		Type:     schema.TypeString,
	},
	"function": {
		Optional: true,
		Type:     schema.TypeString,
	},
	"handler": {
		Optional: true,
		Type:     schema.TypeString,
	},
	"image_uri": {
		Optional: true,
		Type:     schema.TypeString,
	},
	"memory_size": {
		Optional: true,
		Type:     schema.TypeInt,
	},
	"sns_trigger": {
		Optional: true,
		Type:     schema.TypeMap,
		Elem:     &schema.Schema{Type: schema.TypeString},
		// Type:     schema.TypeList,
		// MaxItems: 1,
		// Elem: &schema.Resource{
		// 	Schema: map[string]*schema.Schema{
		// 		"dead_letter_queue": {
		// 			Optional: true,
		// 			Type:     schema.TypeString,
		// 		},
		// 		"sns_topic": {
		// 			Optional: true,
		// 			Type:     schema.TypeString,
		// 		},
		// 	},
		// },
	},
	"source_dir": {
		Optional: true,
		Type:     schema.TypeString,
	},
	"sqs_event_sources": {
		Optional: true,
		Type:     schema.TypeSet,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"batch_size": {
					Optional: true,
					Type:     schema.TypeInt,
				},
				"queue_name": {
					Optional: true,
					Type:     schema.TypeString,
				},
			},
		},
	},
	"vpc_attached": {
		Optional: true,
		Type:     schema.TypeBool,
	},
}
