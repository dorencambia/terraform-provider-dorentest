package lambda

type Lambda struct {
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
}
