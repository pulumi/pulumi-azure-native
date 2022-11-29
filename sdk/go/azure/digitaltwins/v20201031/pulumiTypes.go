


package v20201031

type EventGrid struct {
	AccessKey1       string  `pulumi:"accessKey1"`
	AccessKey2       *string `pulumi:"accessKey2"`
	DeadLetterSecret *string `pulumi:"deadLetterSecret"`
	EndpointType     string  `pulumi:"endpointType"`
	TopicEndpoint    string  `pulumi:"topicEndpoint"`
}

type EventGridResponse struct {
	AccessKey1        string  `pulumi:"accessKey1"`
	AccessKey2        *string `pulumi:"accessKey2"`
	CreatedTime       string  `pulumi:"createdTime"`
	DeadLetterSecret  *string `pulumi:"deadLetterSecret"`
	EndpointType      string  `pulumi:"endpointType"`
	ProvisioningState string  `pulumi:"provisioningState"`
	TopicEndpoint     string  `pulumi:"topicEndpoint"`
}

type EventHub struct {
	ConnectionStringPrimaryKey   string  `pulumi:"connectionStringPrimaryKey"`
	ConnectionStringSecondaryKey *string `pulumi:"connectionStringSecondaryKey"`
	DeadLetterSecret             *string `pulumi:"deadLetterSecret"`
	EndpointType                 string  `pulumi:"endpointType"`
}

type EventHubResponse struct {
	ConnectionStringPrimaryKey   string  `pulumi:"connectionStringPrimaryKey"`
	ConnectionStringSecondaryKey *string `pulumi:"connectionStringSecondaryKey"`
	CreatedTime                  string  `pulumi:"createdTime"`
	DeadLetterSecret             *string `pulumi:"deadLetterSecret"`
	EndpointType                 string  `pulumi:"endpointType"`
	ProvisioningState            string  `pulumi:"provisioningState"`
}

type ServiceBus struct {
	DeadLetterSecret          *string `pulumi:"deadLetterSecret"`
	EndpointType              string  `pulumi:"endpointType"`
	PrimaryConnectionString   string  `pulumi:"primaryConnectionString"`
	SecondaryConnectionString *string `pulumi:"secondaryConnectionString"`
}

type ServiceBusResponse struct {
	CreatedTime               string  `pulumi:"createdTime"`
	DeadLetterSecret          *string `pulumi:"deadLetterSecret"`
	EndpointType              string  `pulumi:"endpointType"`
	PrimaryConnectionString   string  `pulumi:"primaryConnectionString"`
	ProvisioningState         string  `pulumi:"provisioningState"`
	SecondaryConnectionString *string `pulumi:"secondaryConnectionString"`
}

func init() {
}
