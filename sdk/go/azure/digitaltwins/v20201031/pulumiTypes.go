


package v20201031

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EventGrid struct {
	AccessKey1       string  `pulumi:"accessKey1"`
	AccessKey2       *string `pulumi:"accessKey2"`
	DeadLetterSecret *string `pulumi:"deadLetterSecret"`
	EndpointType     string  `pulumi:"endpointType"`
	TopicEndpoint    string  `pulumi:"topicEndpoint"`
}





type EventGridInput interface {
	pulumi.Input

	ToEventGridOutput() EventGridOutput
	ToEventGridOutputWithContext(context.Context) EventGridOutput
}

type EventGridArgs struct {
	AccessKey1       pulumi.StringInput    `pulumi:"accessKey1"`
	AccessKey2       pulumi.StringPtrInput `pulumi:"accessKey2"`
	DeadLetterSecret pulumi.StringPtrInput `pulumi:"deadLetterSecret"`
	EndpointType     pulumi.StringInput    `pulumi:"endpointType"`
	TopicEndpoint    pulumi.StringInput    `pulumi:"topicEndpoint"`
}

func (EventGridArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventGrid)(nil)).Elem()
}

func (i EventGridArgs) ToEventGridOutput() EventGridOutput {
	return i.ToEventGridOutputWithContext(context.Background())
}

func (i EventGridArgs) ToEventGridOutputWithContext(ctx context.Context) EventGridOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventGridOutput)
}

type EventGridOutput struct{ *pulumi.OutputState }

func (EventGridOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventGrid)(nil)).Elem()
}

func (o EventGridOutput) ToEventGridOutput() EventGridOutput {
	return o
}

func (o EventGridOutput) ToEventGridOutputWithContext(ctx context.Context) EventGridOutput {
	return o
}

func (o EventGridOutput) AccessKey1() pulumi.StringOutput {
	return o.ApplyT(func(v EventGrid) string { return v.AccessKey1 }).(pulumi.StringOutput)
}

func (o EventGridOutput) AccessKey2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventGrid) *string { return v.AccessKey2 }).(pulumi.StringPtrOutput)
}

func (o EventGridOutput) DeadLetterSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventGrid) *string { return v.DeadLetterSecret }).(pulumi.StringPtrOutput)
}

func (o EventGridOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v EventGrid) string { return v.EndpointType }).(pulumi.StringOutput)
}

func (o EventGridOutput) TopicEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v EventGrid) string { return v.TopicEndpoint }).(pulumi.StringOutput)
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





type EventGridResponseInput interface {
	pulumi.Input

	ToEventGridResponseOutput() EventGridResponseOutput
	ToEventGridResponseOutputWithContext(context.Context) EventGridResponseOutput
}

type EventGridResponseArgs struct {
	AccessKey1        pulumi.StringInput    `pulumi:"accessKey1"`
	AccessKey2        pulumi.StringPtrInput `pulumi:"accessKey2"`
	CreatedTime       pulumi.StringInput    `pulumi:"createdTime"`
	DeadLetterSecret  pulumi.StringPtrInput `pulumi:"deadLetterSecret"`
	EndpointType      pulumi.StringInput    `pulumi:"endpointType"`
	ProvisioningState pulumi.StringInput    `pulumi:"provisioningState"`
	TopicEndpoint     pulumi.StringInput    `pulumi:"topicEndpoint"`
}

func (EventGridResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventGridResponse)(nil)).Elem()
}

func (i EventGridResponseArgs) ToEventGridResponseOutput() EventGridResponseOutput {
	return i.ToEventGridResponseOutputWithContext(context.Background())
}

func (i EventGridResponseArgs) ToEventGridResponseOutputWithContext(ctx context.Context) EventGridResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventGridResponseOutput)
}

type EventGridResponseOutput struct{ *pulumi.OutputState }

func (EventGridResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventGridResponse)(nil)).Elem()
}

func (o EventGridResponseOutput) ToEventGridResponseOutput() EventGridResponseOutput {
	return o
}

func (o EventGridResponseOutput) ToEventGridResponseOutputWithContext(ctx context.Context) EventGridResponseOutput {
	return o
}

func (o EventGridResponseOutput) AccessKey1() pulumi.StringOutput {
	return o.ApplyT(func(v EventGridResponse) string { return v.AccessKey1 }).(pulumi.StringOutput)
}

func (o EventGridResponseOutput) AccessKey2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventGridResponse) *string { return v.AccessKey2 }).(pulumi.StringPtrOutput)
}

func (o EventGridResponseOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v EventGridResponse) string { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o EventGridResponseOutput) DeadLetterSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventGridResponse) *string { return v.DeadLetterSecret }).(pulumi.StringPtrOutput)
}

func (o EventGridResponseOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v EventGridResponse) string { return v.EndpointType }).(pulumi.StringOutput)
}

func (o EventGridResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v EventGridResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o EventGridResponseOutput) TopicEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v EventGridResponse) string { return v.TopicEndpoint }).(pulumi.StringOutput)
}

type EventHub struct {
	ConnectionStringPrimaryKey   string  `pulumi:"connectionStringPrimaryKey"`
	ConnectionStringSecondaryKey *string `pulumi:"connectionStringSecondaryKey"`
	DeadLetterSecret             *string `pulumi:"deadLetterSecret"`
	EndpointType                 string  `pulumi:"endpointType"`
}





type EventHubInput interface {
	pulumi.Input

	ToEventHubOutput() EventHubOutput
	ToEventHubOutputWithContext(context.Context) EventHubOutput
}

type EventHubArgs struct {
	ConnectionStringPrimaryKey   pulumi.StringInput    `pulumi:"connectionStringPrimaryKey"`
	ConnectionStringSecondaryKey pulumi.StringPtrInput `pulumi:"connectionStringSecondaryKey"`
	DeadLetterSecret             pulumi.StringPtrInput `pulumi:"deadLetterSecret"`
	EndpointType                 pulumi.StringInput    `pulumi:"endpointType"`
}

func (EventHubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHub)(nil)).Elem()
}

func (i EventHubArgs) ToEventHubOutput() EventHubOutput {
	return i.ToEventHubOutputWithContext(context.Background())
}

func (i EventHubArgs) ToEventHubOutputWithContext(ctx context.Context) EventHubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubOutput)
}

type EventHubOutput struct{ *pulumi.OutputState }

func (EventHubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHub)(nil)).Elem()
}

func (o EventHubOutput) ToEventHubOutput() EventHubOutput {
	return o
}

func (o EventHubOutput) ToEventHubOutputWithContext(ctx context.Context) EventHubOutput {
	return o
}

func (o EventHubOutput) ConnectionStringPrimaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v EventHub) string { return v.ConnectionStringPrimaryKey }).(pulumi.StringOutput)
}

func (o EventHubOutput) ConnectionStringSecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHub) *string { return v.ConnectionStringSecondaryKey }).(pulumi.StringPtrOutput)
}

func (o EventHubOutput) DeadLetterSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHub) *string { return v.DeadLetterSecret }).(pulumi.StringPtrOutput)
}

func (o EventHubOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v EventHub) string { return v.EndpointType }).(pulumi.StringOutput)
}

type EventHubResponse struct {
	ConnectionStringPrimaryKey   string  `pulumi:"connectionStringPrimaryKey"`
	ConnectionStringSecondaryKey *string `pulumi:"connectionStringSecondaryKey"`
	CreatedTime                  string  `pulumi:"createdTime"`
	DeadLetterSecret             *string `pulumi:"deadLetterSecret"`
	EndpointType                 string  `pulumi:"endpointType"`
	ProvisioningState            string  `pulumi:"provisioningState"`
}





type EventHubResponseInput interface {
	pulumi.Input

	ToEventHubResponseOutput() EventHubResponseOutput
	ToEventHubResponseOutputWithContext(context.Context) EventHubResponseOutput
}

type EventHubResponseArgs struct {
	ConnectionStringPrimaryKey   pulumi.StringInput    `pulumi:"connectionStringPrimaryKey"`
	ConnectionStringSecondaryKey pulumi.StringPtrInput `pulumi:"connectionStringSecondaryKey"`
	CreatedTime                  pulumi.StringInput    `pulumi:"createdTime"`
	DeadLetterSecret             pulumi.StringPtrInput `pulumi:"deadLetterSecret"`
	EndpointType                 pulumi.StringInput    `pulumi:"endpointType"`
	ProvisioningState            pulumi.StringInput    `pulumi:"provisioningState"`
}

func (EventHubResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubResponse)(nil)).Elem()
}

func (i EventHubResponseArgs) ToEventHubResponseOutput() EventHubResponseOutput {
	return i.ToEventHubResponseOutputWithContext(context.Background())
}

func (i EventHubResponseArgs) ToEventHubResponseOutputWithContext(ctx context.Context) EventHubResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubResponseOutput)
}

type EventHubResponseOutput struct{ *pulumi.OutputState }

func (EventHubResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubResponse)(nil)).Elem()
}

func (o EventHubResponseOutput) ToEventHubResponseOutput() EventHubResponseOutput {
	return o
}

func (o EventHubResponseOutput) ToEventHubResponseOutputWithContext(ctx context.Context) EventHubResponseOutput {
	return o
}

func (o EventHubResponseOutput) ConnectionStringPrimaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v EventHubResponse) string { return v.ConnectionStringPrimaryKey }).(pulumi.StringOutput)
}

func (o EventHubResponseOutput) ConnectionStringSecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubResponse) *string { return v.ConnectionStringSecondaryKey }).(pulumi.StringPtrOutput)
}

func (o EventHubResponseOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v EventHubResponse) string { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o EventHubResponseOutput) DeadLetterSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubResponse) *string { return v.DeadLetterSecret }).(pulumi.StringPtrOutput)
}

func (o EventHubResponseOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v EventHubResponse) string { return v.EndpointType }).(pulumi.StringOutput)
}

func (o EventHubResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v EventHubResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type ServiceBus struct {
	DeadLetterSecret          *string `pulumi:"deadLetterSecret"`
	EndpointType              string  `pulumi:"endpointType"`
	PrimaryConnectionString   string  `pulumi:"primaryConnectionString"`
	SecondaryConnectionString *string `pulumi:"secondaryConnectionString"`
}





type ServiceBusInput interface {
	pulumi.Input

	ToServiceBusOutput() ServiceBusOutput
	ToServiceBusOutputWithContext(context.Context) ServiceBusOutput
}

type ServiceBusArgs struct {
	DeadLetterSecret          pulumi.StringPtrInput `pulumi:"deadLetterSecret"`
	EndpointType              pulumi.StringInput    `pulumi:"endpointType"`
	PrimaryConnectionString   pulumi.StringInput    `pulumi:"primaryConnectionString"`
	SecondaryConnectionString pulumi.StringPtrInput `pulumi:"secondaryConnectionString"`
}

func (ServiceBusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceBus)(nil)).Elem()
}

func (i ServiceBusArgs) ToServiceBusOutput() ServiceBusOutput {
	return i.ToServiceBusOutputWithContext(context.Background())
}

func (i ServiceBusArgs) ToServiceBusOutputWithContext(ctx context.Context) ServiceBusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceBusOutput)
}

type ServiceBusOutput struct{ *pulumi.OutputState }

func (ServiceBusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceBus)(nil)).Elem()
}

func (o ServiceBusOutput) ToServiceBusOutput() ServiceBusOutput {
	return o
}

func (o ServiceBusOutput) ToServiceBusOutputWithContext(ctx context.Context) ServiceBusOutput {
	return o
}

func (o ServiceBusOutput) DeadLetterSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBus) *string { return v.DeadLetterSecret }).(pulumi.StringPtrOutput)
}

func (o ServiceBusOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceBus) string { return v.EndpointType }).(pulumi.StringOutput)
}

func (o ServiceBusOutput) PrimaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceBus) string { return v.PrimaryConnectionString }).(pulumi.StringOutput)
}

func (o ServiceBusOutput) SecondaryConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBus) *string { return v.SecondaryConnectionString }).(pulumi.StringPtrOutput)
}

type ServiceBusResponse struct {
	CreatedTime               string  `pulumi:"createdTime"`
	DeadLetterSecret          *string `pulumi:"deadLetterSecret"`
	EndpointType              string  `pulumi:"endpointType"`
	PrimaryConnectionString   string  `pulumi:"primaryConnectionString"`
	ProvisioningState         string  `pulumi:"provisioningState"`
	SecondaryConnectionString *string `pulumi:"secondaryConnectionString"`
}





type ServiceBusResponseInput interface {
	pulumi.Input

	ToServiceBusResponseOutput() ServiceBusResponseOutput
	ToServiceBusResponseOutputWithContext(context.Context) ServiceBusResponseOutput
}

type ServiceBusResponseArgs struct {
	CreatedTime               pulumi.StringInput    `pulumi:"createdTime"`
	DeadLetterSecret          pulumi.StringPtrInput `pulumi:"deadLetterSecret"`
	EndpointType              pulumi.StringInput    `pulumi:"endpointType"`
	PrimaryConnectionString   pulumi.StringInput    `pulumi:"primaryConnectionString"`
	ProvisioningState         pulumi.StringInput    `pulumi:"provisioningState"`
	SecondaryConnectionString pulumi.StringPtrInput `pulumi:"secondaryConnectionString"`
}

func (ServiceBusResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceBusResponse)(nil)).Elem()
}

func (i ServiceBusResponseArgs) ToServiceBusResponseOutput() ServiceBusResponseOutput {
	return i.ToServiceBusResponseOutputWithContext(context.Background())
}

func (i ServiceBusResponseArgs) ToServiceBusResponseOutputWithContext(ctx context.Context) ServiceBusResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceBusResponseOutput)
}

type ServiceBusResponseOutput struct{ *pulumi.OutputState }

func (ServiceBusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceBusResponse)(nil)).Elem()
}

func (o ServiceBusResponseOutput) ToServiceBusResponseOutput() ServiceBusResponseOutput {
	return o
}

func (o ServiceBusResponseOutput) ToServiceBusResponseOutputWithContext(ctx context.Context) ServiceBusResponseOutput {
	return o
}

func (o ServiceBusResponseOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceBusResponse) string { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o ServiceBusResponseOutput) DeadLetterSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusResponse) *string { return v.DeadLetterSecret }).(pulumi.StringPtrOutput)
}

func (o ServiceBusResponseOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceBusResponse) string { return v.EndpointType }).(pulumi.StringOutput)
}

func (o ServiceBusResponseOutput) PrimaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceBusResponse) string { return v.PrimaryConnectionString }).(pulumi.StringOutput)
}

func (o ServiceBusResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceBusResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ServiceBusResponseOutput) SecondaryConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusResponse) *string { return v.SecondaryConnectionString }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(EventGridOutput{})
	pulumi.RegisterOutputType(EventGridResponseOutput{})
	pulumi.RegisterOutputType(EventHubOutput{})
	pulumi.RegisterOutputType(EventHubResponseOutput{})
	pulumi.RegisterOutputType(ServiceBusOutput{})
	pulumi.RegisterOutputType(ServiceBusResponseOutput{})
}
