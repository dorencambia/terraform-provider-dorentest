package lambda

import "encoding/json"

// https://learn.hashicorp.com/tutorials/terraform/provider-complex-read

func Converter(lambdas []Lambda) []interface{} {
	if lambdas == nil {
		return nil
	}
	converted := make([]interface{}, len(lambdas))
	for i, lambda := range lambdas {
		x := make(map[string]interface{})
		x["function_name"] = lambda.Function
		x["function"] = lambda.Function
		x["image_uri"] = lambda.ImageURI
		x["handler"] = lambda.Handler
		x["description"] = lambda.Description
		x["source_dir"] = lambda.SourceDir
		x["concurrency_limit"] = lambda.ConcurrencyLimit
		x["memory_size"] = lambda.MemorySize
		x["environment_variables"] = StructToMap(lambda.EnvironmentVariables)
		x["vpc_attached"] = lambda.VpcAttached
		x["egress_rules"] = lambda.EgressRules
		x["sns_trigger"] = StructToMap(lambda.SnsTrigger)
		x["sqs_event_sources"] = lambda.SqsEventSources
		// x["cloudwatch_event_rule"] = StructToMap(&lambda.CloudwatchEventRule)
		converted[i] = x
	}
	return converted
}

// Converts a struct to a map while maintaining the json alias as keys
func StructToMap(obj interface{}) (newMap map[string]interface{}) {
	data, err := json.Marshal(obj) // Convert to a json string

	if err != nil {
		return nil
	}

	err = json.Unmarshal(data, &newMap) // Convert to a map
	return newMap
}

// func flattenOrderItemsData(orderItems *[]hc.OrderItem) []interface{} {
// 	if orderItems != nil {
// 		ois := make([]interface{}, len(*orderItems), len(*orderItems))

// 		for i, orderItem := range *orderItems {
// 			oi := make(map[string]interface{})

// 			oi["coffee_id"] = orderItem.Coffee.ID
// 			oi["coffee_name"] = orderItem.Coffee.Name
// 			oi["coffee_teaser"] = orderItem.Coffee.Teaser
// 			oi["coffee_description"] = orderItem.Coffee.Description
// 			oi["coffee_price"] = orderItem.Coffee.Price
// 			oi["coffee_image"] = orderItem.Coffee.Image
// 			oi["quantity"] = orderItem.Quantity

// 			ois[i] = oi
// 		}

// 		return ois
// 	}

// 	return make([]interface{}, 0)
// }

// https://github.com/hashicorp/terraform-provider-aws/blob/07b55c2a4f02262668eb3b41ba3526b0d599fc91/internal/service/lambda/flex.go#L21
// func flattenLayers(layers []*lambda.Layer) []interface{} {
// 	arns := make([]*string, len(layers))
// 	for i, layer := range layers {
// 		arns[i] = layer.Arn
// 	}
// 	return flex.FlattenStringList(arns)
// }
