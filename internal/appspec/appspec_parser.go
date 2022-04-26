package appspec

import (
	"gopkg.in/yaml.v2"
)

type Appspec struct {
	Lambdas struct {
		Version         string `yaml:"version"`
		LambdaFunctions []struct {
			Function             string `yaml:"function"`
			ImageURI             string `yaml:"image_uri"`
			Handler              string `yaml:"handler,omitempty"`
			Description          string `yaml:"description,omitempty"`
			SourceDir            string `yaml:"source_dir,omitempty"`
			ConcurrencyLimit     int    `yaml:"concurrency_limit,omitempty"`
			EnvironmentVariables struct {
				Testvar1 string `yaml:"testvar1"`
				Testvar2 string `yaml:"testvar2"`
			} `yaml:"environment_variables,omitempty"`
			VpcAttached bool `yaml:"vpc_attached,omitempty"`
			EgressRules []struct {
				Proto    string `yaml:"proto"`
				Ports    []int  `yaml:"ports"`
				CidrIP   string `yaml:"cidr_ip"`
				RuleDesc string `yaml:"rule_desc"`
			} `yaml:"egress_rules,omitempty"`
			SnsTrigger struct {
				SnsTopic        string `yaml:"sns_topic"`
				DeadLetterQueue string `yaml:"dead_letter_queue"`
			} `yaml:"sns_trigger,omitempty"`
			SqsEventSources []struct {
				QueueName string `yaml:"queue_name"`
				BatchSize int    `yaml:"batch_size"`
			} `yaml:"sqs_event_sources,omitempty"`
			CloudwatchEventRule struct {
				Name               string `yaml:"name"`
				Description        string `yaml:"description"`
				ScheduleExpression string `yaml:"schedule_expression"`
				Input              struct {
					Source string `yaml:"source"`
					Env    string `yaml:"env"`
					ID     string `yaml:"id"`
				} `yaml:"input"`
			} `yaml:"cloudwatch_event_rule,omitempty"`
		} `yaml:"lambda_functions"`
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

func ParseAppspec(input string) (Appspec, error) {
	a := Appspec{}
	err := yaml.Unmarshal([]byte(input), &a)
	if err != nil {
		return Appspec{}, err
	}
	return a, nil
}
