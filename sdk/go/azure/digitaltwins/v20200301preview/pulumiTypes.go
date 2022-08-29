


package v20200301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DigitalTwinsSkuInfo struct {
	Name string `pulumi:"name"`
}





type DigitalTwinsSkuInfoInput interface {
	pulumi.Input

	ToDigitalTwinsSkuInfoOutput() DigitalTwinsSkuInfoOutput
	ToDigitalTwinsSkuInfoOutputWithContext(context.Context) DigitalTwinsSkuInfoOutput
}

type DigitalTwinsSkuInfoArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (DigitalTwinsSkuInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DigitalTwinsSkuInfo)(nil)).Elem()
}

func (i DigitalTwinsSkuInfoArgs) ToDigitalTwinsSkuInfoOutput() DigitalTwinsSkuInfoOutput {
	return i.ToDigitalTwinsSkuInfoOutputWithContext(context.Background())
}

func (i DigitalTwinsSkuInfoArgs) ToDigitalTwinsSkuInfoOutputWithContext(ctx context.Context) DigitalTwinsSkuInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DigitalTwinsSkuInfoOutput)
}

func (i DigitalTwinsSkuInfoArgs) ToDigitalTwinsSkuInfoPtrOutput() DigitalTwinsSkuInfoPtrOutput {
	return i.ToDigitalTwinsSkuInfoPtrOutputWithContext(context.Background())
}

func (i DigitalTwinsSkuInfoArgs) ToDigitalTwinsSkuInfoPtrOutputWithContext(ctx context.Context) DigitalTwinsSkuInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DigitalTwinsSkuInfoOutput).ToDigitalTwinsSkuInfoPtrOutputWithContext(ctx)
}









type DigitalTwinsSkuInfoPtrInput interface {
	pulumi.Input

	ToDigitalTwinsSkuInfoPtrOutput() DigitalTwinsSkuInfoPtrOutput
	ToDigitalTwinsSkuInfoPtrOutputWithContext(context.Context) DigitalTwinsSkuInfoPtrOutput
}

type digitalTwinsSkuInfoPtrType DigitalTwinsSkuInfoArgs

func DigitalTwinsSkuInfoPtr(v *DigitalTwinsSkuInfoArgs) DigitalTwinsSkuInfoPtrInput {
	return (*digitalTwinsSkuInfoPtrType)(v)
}

func (*digitalTwinsSkuInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DigitalTwinsSkuInfo)(nil)).Elem()
}

func (i *digitalTwinsSkuInfoPtrType) ToDigitalTwinsSkuInfoPtrOutput() DigitalTwinsSkuInfoPtrOutput {
	return i.ToDigitalTwinsSkuInfoPtrOutputWithContext(context.Background())
}

func (i *digitalTwinsSkuInfoPtrType) ToDigitalTwinsSkuInfoPtrOutputWithContext(ctx context.Context) DigitalTwinsSkuInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DigitalTwinsSkuInfoPtrOutput)
}

type DigitalTwinsSkuInfoOutput struct{ *pulumi.OutputState }

func (DigitalTwinsSkuInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DigitalTwinsSkuInfo)(nil)).Elem()
}

func (o DigitalTwinsSkuInfoOutput) ToDigitalTwinsSkuInfoOutput() DigitalTwinsSkuInfoOutput {
	return o
}

func (o DigitalTwinsSkuInfoOutput) ToDigitalTwinsSkuInfoOutputWithContext(ctx context.Context) DigitalTwinsSkuInfoOutput {
	return o
}

func (o DigitalTwinsSkuInfoOutput) ToDigitalTwinsSkuInfoPtrOutput() DigitalTwinsSkuInfoPtrOutput {
	return o.ToDigitalTwinsSkuInfoPtrOutputWithContext(context.Background())
}

func (o DigitalTwinsSkuInfoOutput) ToDigitalTwinsSkuInfoPtrOutputWithContext(ctx context.Context) DigitalTwinsSkuInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DigitalTwinsSkuInfo) *DigitalTwinsSkuInfo {
		return &v
	}).(DigitalTwinsSkuInfoPtrOutput)
}

func (o DigitalTwinsSkuInfoOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DigitalTwinsSkuInfo) string { return v.Name }).(pulumi.StringOutput)
}

type DigitalTwinsSkuInfoPtrOutput struct{ *pulumi.OutputState }

func (DigitalTwinsSkuInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DigitalTwinsSkuInfo)(nil)).Elem()
}

func (o DigitalTwinsSkuInfoPtrOutput) ToDigitalTwinsSkuInfoPtrOutput() DigitalTwinsSkuInfoPtrOutput {
	return o
}

func (o DigitalTwinsSkuInfoPtrOutput) ToDigitalTwinsSkuInfoPtrOutputWithContext(ctx context.Context) DigitalTwinsSkuInfoPtrOutput {
	return o
}

func (o DigitalTwinsSkuInfoPtrOutput) Elem() DigitalTwinsSkuInfoOutput {
	return o.ApplyT(func(v *DigitalTwinsSkuInfo) DigitalTwinsSkuInfo {
		if v != nil {
			return *v
		}
		var ret DigitalTwinsSkuInfo
		return ret
	}).(DigitalTwinsSkuInfoOutput)
}

func (o DigitalTwinsSkuInfoPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DigitalTwinsSkuInfo) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type DigitalTwinsSkuInfoResponse struct {
	Name string `pulumi:"name"`
}

type DigitalTwinsSkuInfoResponseOutput struct{ *pulumi.OutputState }

func (DigitalTwinsSkuInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DigitalTwinsSkuInfoResponse)(nil)).Elem()
}

func (o DigitalTwinsSkuInfoResponseOutput) ToDigitalTwinsSkuInfoResponseOutput() DigitalTwinsSkuInfoResponseOutput {
	return o
}

func (o DigitalTwinsSkuInfoResponseOutput) ToDigitalTwinsSkuInfoResponseOutputWithContext(ctx context.Context) DigitalTwinsSkuInfoResponseOutput {
	return o
}

func (o DigitalTwinsSkuInfoResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DigitalTwinsSkuInfoResponse) string { return v.Name }).(pulumi.StringOutput)
}

type DigitalTwinsSkuInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (DigitalTwinsSkuInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DigitalTwinsSkuInfoResponse)(nil)).Elem()
}

func (o DigitalTwinsSkuInfoResponsePtrOutput) ToDigitalTwinsSkuInfoResponsePtrOutput() DigitalTwinsSkuInfoResponsePtrOutput {
	return o
}

func (o DigitalTwinsSkuInfoResponsePtrOutput) ToDigitalTwinsSkuInfoResponsePtrOutputWithContext(ctx context.Context) DigitalTwinsSkuInfoResponsePtrOutput {
	return o
}

func (o DigitalTwinsSkuInfoResponsePtrOutput) Elem() DigitalTwinsSkuInfoResponseOutput {
	return o.ApplyT(func(v *DigitalTwinsSkuInfoResponse) DigitalTwinsSkuInfoResponse {
		if v != nil {
			return *v
		}
		var ret DigitalTwinsSkuInfoResponse
		return ret
	}).(DigitalTwinsSkuInfoResponseOutput)
}

func (o DigitalTwinsSkuInfoResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DigitalTwinsSkuInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type EventGrid struct {
	AccessKey1    string            `pulumi:"accessKey1"`
	AccessKey2    string            `pulumi:"accessKey2"`
	EndpointType  string            `pulumi:"endpointType"`
	Tags          map[string]string `pulumi:"tags"`
	TopicEndpoint *string           `pulumi:"topicEndpoint"`
}

type EventGridResponse struct {
	AccessKey1        string            `pulumi:"accessKey1"`
	AccessKey2        string            `pulumi:"accessKey2"`
	CreatedTime       string            `pulumi:"createdTime"`
	EndpointType      string            `pulumi:"endpointType"`
	ProvisioningState string            `pulumi:"provisioningState"`
	Tags              map[string]string `pulumi:"tags"`
	TopicEndpoint     *string           `pulumi:"topicEndpoint"`
}

type EventHub struct {
	ConnectionStringPrimaryKey   string            `pulumi:"connectionStringPrimaryKey"`
	ConnectionStringSecondaryKey string            `pulumi:"connectionStringSecondaryKey"`
	EndpointType                 string            `pulumi:"endpointType"`
	Tags                         map[string]string `pulumi:"tags"`
}

type EventHubResponse struct {
	ConnectionStringPrimaryKey   string            `pulumi:"connectionStringPrimaryKey"`
	ConnectionStringSecondaryKey string            `pulumi:"connectionStringSecondaryKey"`
	CreatedTime                  string            `pulumi:"createdTime"`
	EndpointType                 string            `pulumi:"endpointType"`
	ProvisioningState            string            `pulumi:"provisioningState"`
	Tags                         map[string]string `pulumi:"tags"`
}

type ServiceBus struct {
	EndpointType              string            `pulumi:"endpointType"`
	PrimaryConnectionString   string            `pulumi:"primaryConnectionString"`
	SecondaryConnectionString string            `pulumi:"secondaryConnectionString"`
	Tags                      map[string]string `pulumi:"tags"`
}

type ServiceBusResponse struct {
	CreatedTime               string            `pulumi:"createdTime"`
	EndpointType              string            `pulumi:"endpointType"`
	PrimaryConnectionString   string            `pulumi:"primaryConnectionString"`
	ProvisioningState         string            `pulumi:"provisioningState"`
	SecondaryConnectionString string            `pulumi:"secondaryConnectionString"`
	Tags                      map[string]string `pulumi:"tags"`
}

func init() {
	pulumi.RegisterOutputType(DigitalTwinsSkuInfoOutput{})
	pulumi.RegisterOutputType(DigitalTwinsSkuInfoPtrOutput{})
	pulumi.RegisterOutputType(DigitalTwinsSkuInfoResponseOutput{})
	pulumi.RegisterOutputType(DigitalTwinsSkuInfoResponsePtrOutput{})
}
