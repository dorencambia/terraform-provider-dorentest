package appspec

import (
	"github.com/dorencambia/terraform-provider-dorentest/internal/appspec/lambda"
	"gopkg.in/yaml.v2"
)

func ParseAppspec(input string) (Appspec, error) {
	a := Appspec{}
	// var a Appspec
	err := yaml.Unmarshal([]byte(input), &a)
	if err != nil {
		return Appspec{}, err
	}
	return a, nil
}

type Appspec struct {
	Lambdas struct {
		Version         string          `yaml:"version"`
		LambdaFunctions []lambda.Lambda `yaml:"lambda_functions"`
	} `yaml:"lambdas"`
	SnsTopics struct {
		Version string `yaml:"version"`
		Topics  []struct {
			Function       string `yaml:"function"`
			State          string `yaml:"state"`
			DeliveryPolicy struct {
				HTTP struct {
					DefaultHealthyRetryPolicy struct {
						NumRetries         int    `yaml:"numRetries"`
						MinDelayTarget     int    `yaml:"minDelayTarget"`
						MaxDelayTarget     int    `yaml:"maxDelayTarget"`
						NumMaxDelayRetries int    `yaml:"numMaxDelayRetries"`
						NumNoDelayRetries  int    `yaml:"numNoDelayRetries"`
						NumMinDelayRetries int    `yaml:"numMinDelayRetries"`
						BackoffFunction    string `yaml:"backoffFunction"`
					} `yaml:"defaultHealthyRetryPolicy"`
					DisableSubscriptionOverrides bool `yaml:"disableSubscriptionOverrides"`
					DefaultThrottlePolicy        struct {
						MaxReceivesPerSecond int `yaml:"maxReceivesPerSecond"`
					} `yaml:"defaultThrottlePolicy"`
				} `yaml:"http"`
			} `yaml:"delivery_policy"`
			Subscriptions []struct {
				Endpoint string `yaml:"endpoint"`
				Protocol string `yaml:"protocol"`
			} `yaml:"subscriptions"`
			PurgeSubscriptions bool   `yaml:"purge_subscriptions"`
			KmsAlias           string `yaml:"kms_alias"`
		} `yaml:"topics"`
	} `yaml:"sns_topics"`
}
