package lambda

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Schema() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Pass in the content of a juno_appspec.yml file and get modified data back",

		// ReadContext: dataSourceAppspecRead,

		Schema: map[string]*schema.Schema{
			"function_name": {
				Description: "A unique name for your Lambda Function.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"sns_trigger": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Trigger the lambda from an SNS topic",
				// Default:     "{}",
			},
			"sqs_event_sources": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Trigger the lambda from these SQS queues",
				// Default:     "[]",
			},
			"cloudwatch_event_rule": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Trigger the lambda from an SNS topic",
				// Default:     "{}",
			},
			"vpc_attached": {
				Optional: true,
				Type:     schema.TypeString,
				// Default: false,
			},
			"egress_rules": {
				Optional: true,
				Type:     schema.TypeString,
				// Default: "[]",
			},
			"image_uri": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Full path to a Docker image",
				// Default:     "",
			},
			"runtime": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The runtime environment for the Lambda function you are uploading.",
				// Default:     "",
			},
			"handler": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The function entrypoint in your code.",
				// Default:     "",
			},
			"filename": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The path to the function's deployment package within the local filesystem.",
				// Default:     "",
			},
			"description": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Description of what your Lambda Function does.",
				// Default:     "",
			},
			"environment": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Environment (e.g. env variables) configuration for the Lambda function enable you to dynamically pass settings to your function code and libraries",
				// Type:        "map(string)",
				// Default:     {},
			},
			"memory_size": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Amount of memory in MB your Lambda Function can use at runtime. Defaults to 128.",
				// Default:     128,
			},
			"publish": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Whether to publish creation/change as new Lambda Function Version. Defaults to false.",
				// Default:     false,
			},
			"reserved_concurrent_executions": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The amount of reserved concurrent executions for this lambda function. A value of 0 disables lambda from being triggered and -1 removes any concurrency limitations. Defaults to Unreserved Concurrency Limits -1.",
				// Default:     "-1",
			},
			"tags": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "A mapping of tags to assign to the Lambda function.",
				// Type:        "map(string)",
				// Default:     {},
			},
			"timeout": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The amount of time your Lambda Function has to run in seconds. Defaults to 3.",
				// Default:     3,
			},
			"vpc_config": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Provide this to allow your function to access your VPC (if both 'subnet_ids' and 'security_group_ids' are empty then vpc_config is considered to be empty or unset, see https://docs.aws.amazon.com/lambda/latest/dg/vpc.html for details).",
				// Type:        "map(list(string))",
				// Default:     {},
			},
		},
	}
}

// func Schema() map[string]*schema.Schema {
// 	return map[string]*schema.Schema{
// 		"function_name": {
// 			Description: "A unique name for your Lambda Function.",
// 			Type:        schema.TypeString,
// 		},
// 		// "sns_trigger": {
// 		// 	Description: "Trigger the lambda from an SNS topic",
// 		// 	// Default:     "{}",
// 		// },
// 		// "sqs_event_sources": {
// 		// 	Description: "Trigger the lambda from these SQS queues",
// 		// 	// Default:     "[]",
// 		// },
// 		// "cloudwatch_event_rule": {
// 		// 	Description: "Trigger the lambda from an SNS topic",
// 		// 	// Default:     "{}",
// 		// },
// 		// "vpc_attached": {
// 		// 	// Default: false,
// 		// },
// 		// "egress_rules": {
// 		// 	// Default: "[]",
// 		// },
// 		// "account": {
// 		// 	Type: schema.TypeString,
// 		// },
// 		// "function": {
// 		// 	Type: schema.TypeString,
// 		// },
// 		// "namespace": {
// 		// 	Type: schema.TypeString,
// 		// },
// 		// "region": {
// 		// 	Type: schema.TypeString,
// 		// },
// 		// "image_uri": {
// 		// 	Description: "Full path to a Docker image",
// 		// 	// Default:     "",
// 		// },
// 		// "runtime": {
// 		// 	Description: "The runtime environment for the Lambda function you are uploading.",
// 		// 	// Default:     "",
// 		// },
// 		// "handler": {
// 		// 	Description: "The function entrypoint in your code.",
// 		// 	// Default:     "",
// 		// },
// 		// "filename": {
// 		// 	Description: "The path to the function's deployment package within the local filesystem.",
// 		// 	// Default:     "",
// 		// },
// 		// "description": {
// 		// 	Description: "Description of what your Lambda Function does.",
// 		// 	// Default:     "",
// 		// },
// 		// "environment": {
// 		// 	Description: "Environment (e.g. env variables) configuration for the Lambda function enable you to dynamically pass settings to your function code and libraries",
// 		// 	// Type:        "map(string)",
// 		// 	// Default:     {},
// 		// },
// 		// "memory_size": {
// 		// 	Description: "Amount of memory in MB your Lambda Function can use at runtime. Defaults to 128.",
// 		// 	// Default:     128,
// 		// },
// 		// "publish": {
// 		// 	Description: "Whether to publish creation/change as new Lambda Function Version. Defaults to false.",
// 		// 	// Default:     false,
// 		// },
// 		// "reserved_concurrent_executions": {
// 		// 	Description: "The amount of reserved concurrent executions for this lambda function. A value of 0 disables lambda from being triggered and -1 removes any concurrency limitations. Defaults to Unreserved Concurrency Limits -1.",
// 		// 	// Default:     "-1",
// 		// },
// 		// "tags": {
// 		// 	Description: "A mapping of tags to assign to the Lambda function.",
// 		// 	// Type:        "map(string)",
// 		// 	// Default:     {},
// 		// },
// 		// "timeout": {
// 		// 	Description: "The amount of time your Lambda Function has to run in seconds. Defaults to 3.",
// 		// 	// Default:     3,
// 		// },
// 		// "vpc_config": {
// 		// 	Description: "Provide this to allow your function to access your VPC (if both 'subnet_ids' and 'security_group_ids' are empty then vpc_config is considered to be empty or unset, see https://docs.aws.amazon.com/lambda/latest/dg/vpc.html for details).",
// 		// 	// Type:        "map(list(string))",
// 		// 	// Default:     {},
// 		// },
// 	}
// 	// Schema: map[string]*schema.Schema{
// 	// 	"raw_appspec": {
// 	// 		// This description is used by the documentation generator and the language server.
// 	// 		Description: "Raw text of juno_appspec.yml file",
// 	// 		Type:        schema.TypeString,
// 	// 		Required:    true,
// 	// 	},
// 	// 	"lambdas": {
// 	// 		Computed: true,
// 	// 		Type:     schema.TypeList,
// 	// 		Elem:     appspec.LambdaSchema(),
// 	// 	},
// 	// },
// }
