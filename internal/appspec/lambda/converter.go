package lambda

// https://learn.hashicorp.com/tutorials/terraform/provider-complex-read

func Converter(lambdas []Lambda) []interface{} {
	if lambdas == nil {
		return nil
	}
	converted := make([]interface{}, len(lambdas), len(lambdas))
	for i, lambda := range lambdas {
		x := make(map[string]interface{})
		x["function_name"] = lambda.Function
		converted[i] = x
	}
	return converted
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
