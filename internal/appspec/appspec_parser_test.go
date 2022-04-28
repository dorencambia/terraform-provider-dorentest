package appspec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseAppspec(t *testing.T) {
	input := `---
lambdas:
  version: 1.0.0
  lambda_functions:
    - function: "{{ function }}-2"
      image_uri: 395127396906.dkr.ecr.us-west-2.amazonaws.com/chp-test-r635542-demo-lambda-container:latest
    - function: "{{ function }}"
      image_uri: 395127396906.dkr.ecr.us-west-2.amazonaws.com/chp-test-r635542-demo-lambda-container:latest
      # image_tag: latest
      # runtime: 'nodejs14.x'
      handler: "event_source.handler"
      description: "testing event_source lambda {{ env }}"
      source_dir: "src"
      concurrency_limit: 5
      environment_variables:
        testvar1: "testvalue1"
        testvar2: "testvalue2"
      vpc_attached: True
      egress_rules:
        - proto: tcp
          ports:
            - 443
            - 80
          cidr_ip: 10.0.0.0/8
          rule_desc: "egress from cambia networks"
      sns_trigger:
        sns_topic: "{{ function }}-sns-topic"
        dead_letter_queue: "{{ function }}-dlq"
      sqs_event_sources:
        - queue_name: "{{ function }}-event"
          batch_size: 5 # 5 queued messages per run
      cloudwatch_event_rule:
        name: "{{ namespace }}-{{ function }}-{{ env }}"
        description: "This is my cloudwatch event rule"
        schedule_expression: "cron(0 20 * * ? *)"
        input:
          source: "my.trigger"
          env: "poc"
          id: "122345"
sns_topics:
  version: 0.0.1
  topics:
    - function: "{{ function }}-sns-topic"
      state: present
      delivery_policy:
        http:
          defaultHealthyRetryPolicy:
            numRetries: 3
            minDelayTarget: 20
            maxDelayTarget: 20
            numMaxDelayRetries: 0
            numNoDelayRetries: 0
            numMinDelayRetries: 0
            backoffFunction: "linear"
          disableSubscriptionOverrides: False
          defaultThrottlePolicy:
            maxReceivesPerSecond: 10
      subscriptions:
        - endpoint: "doren.proctor@cambiahealth.com"
          protocol: "email"
      purge_subscriptions: True
      kms_alias: "aws/sns"
`
	got, err := ParseAppspec(input)
	assert.Equal(t, nil, err)
	// assert.Equal(t, "", got)
	assert.Equal(t, "0.0.1", got.SnsTopics.Version)
	assert.Len(t, got.Lambdas.LambdaFunctions, 2)
	assert.Equal(t, "{{ function }}-2", got.Lambdas.LambdaFunctions[0].Function)
}
