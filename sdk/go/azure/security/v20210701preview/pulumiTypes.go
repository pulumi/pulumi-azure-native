


package v20210701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CspmMonitorAwsOffering struct {
	NativeCloudConnection *CspmMonitorAwsOfferingNativeCloudConnection `pulumi:"nativeCloudConnection"`
	OfferingType          string                                       `pulumi:"offeringType"`
}





type CspmMonitorAwsOfferingInput interface {
	pulumi.Input

	ToCspmMonitorAwsOfferingOutput() CspmMonitorAwsOfferingOutput
	ToCspmMonitorAwsOfferingOutputWithContext(context.Context) CspmMonitorAwsOfferingOutput
}

type CspmMonitorAwsOfferingArgs struct {
	NativeCloudConnection CspmMonitorAwsOfferingNativeCloudConnectionPtrInput `pulumi:"nativeCloudConnection"`
	OfferingType          pulumi.StringInput                                  `pulumi:"offeringType"`
}

func (CspmMonitorAwsOfferingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CspmMonitorAwsOffering)(nil)).Elem()
}

func (i CspmMonitorAwsOfferingArgs) ToCspmMonitorAwsOfferingOutput() CspmMonitorAwsOfferingOutput {
	return i.ToCspmMonitorAwsOfferingOutputWithContext(context.Background())
}

func (i CspmMonitorAwsOfferingArgs) ToCspmMonitorAwsOfferingOutputWithContext(ctx context.Context) CspmMonitorAwsOfferingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CspmMonitorAwsOfferingOutput)
}

type CspmMonitorAwsOfferingOutput struct{ *pulumi.OutputState }

func (CspmMonitorAwsOfferingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CspmMonitorAwsOffering)(nil)).Elem()
}

func (o CspmMonitorAwsOfferingOutput) ToCspmMonitorAwsOfferingOutput() CspmMonitorAwsOfferingOutput {
	return o
}

func (o CspmMonitorAwsOfferingOutput) ToCspmMonitorAwsOfferingOutputWithContext(ctx context.Context) CspmMonitorAwsOfferingOutput {
	return o
}

func (o CspmMonitorAwsOfferingOutput) NativeCloudConnection() CspmMonitorAwsOfferingNativeCloudConnectionPtrOutput {
	return o.ApplyT(func(v CspmMonitorAwsOffering) *CspmMonitorAwsOfferingNativeCloudConnection {
		return v.NativeCloudConnection
	}).(CspmMonitorAwsOfferingNativeCloudConnectionPtrOutput)
}

func (o CspmMonitorAwsOfferingOutput) OfferingType() pulumi.StringOutput {
	return o.ApplyT(func(v CspmMonitorAwsOffering) string { return v.OfferingType }).(pulumi.StringOutput)
}

type CspmMonitorAwsOfferingNativeCloudConnection struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}





type CspmMonitorAwsOfferingNativeCloudConnectionInput interface {
	pulumi.Input

	ToCspmMonitorAwsOfferingNativeCloudConnectionOutput() CspmMonitorAwsOfferingNativeCloudConnectionOutput
	ToCspmMonitorAwsOfferingNativeCloudConnectionOutputWithContext(context.Context) CspmMonitorAwsOfferingNativeCloudConnectionOutput
}

type CspmMonitorAwsOfferingNativeCloudConnectionArgs struct {
	CloudRoleArn pulumi.StringPtrInput `pulumi:"cloudRoleArn"`
}

func (CspmMonitorAwsOfferingNativeCloudConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CspmMonitorAwsOfferingNativeCloudConnection)(nil)).Elem()
}

func (i CspmMonitorAwsOfferingNativeCloudConnectionArgs) ToCspmMonitorAwsOfferingNativeCloudConnectionOutput() CspmMonitorAwsOfferingNativeCloudConnectionOutput {
	return i.ToCspmMonitorAwsOfferingNativeCloudConnectionOutputWithContext(context.Background())
}

func (i CspmMonitorAwsOfferingNativeCloudConnectionArgs) ToCspmMonitorAwsOfferingNativeCloudConnectionOutputWithContext(ctx context.Context) CspmMonitorAwsOfferingNativeCloudConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CspmMonitorAwsOfferingNativeCloudConnectionOutput)
}

func (i CspmMonitorAwsOfferingNativeCloudConnectionArgs) ToCspmMonitorAwsOfferingNativeCloudConnectionPtrOutput() CspmMonitorAwsOfferingNativeCloudConnectionPtrOutput {
	return i.ToCspmMonitorAwsOfferingNativeCloudConnectionPtrOutputWithContext(context.Background())
}

func (i CspmMonitorAwsOfferingNativeCloudConnectionArgs) ToCspmMonitorAwsOfferingNativeCloudConnectionPtrOutputWithContext(ctx context.Context) CspmMonitorAwsOfferingNativeCloudConnectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CspmMonitorAwsOfferingNativeCloudConnectionOutput).ToCspmMonitorAwsOfferingNativeCloudConnectionPtrOutputWithContext(ctx)
}









type CspmMonitorAwsOfferingNativeCloudConnectionPtrInput interface {
	pulumi.Input

	ToCspmMonitorAwsOfferingNativeCloudConnectionPtrOutput() CspmMonitorAwsOfferingNativeCloudConnectionPtrOutput
	ToCspmMonitorAwsOfferingNativeCloudConnectionPtrOutputWithContext(context.Context) CspmMonitorAwsOfferingNativeCloudConnectionPtrOutput
}

type cspmMonitorAwsOfferingNativeCloudConnectionPtrType CspmMonitorAwsOfferingNativeCloudConnectionArgs

func CspmMonitorAwsOfferingNativeCloudConnectionPtr(v *CspmMonitorAwsOfferingNativeCloudConnectionArgs) CspmMonitorAwsOfferingNativeCloudConnectionPtrInput {
	return (*cspmMonitorAwsOfferingNativeCloudConnectionPtrType)(v)
}

func (*cspmMonitorAwsOfferingNativeCloudConnectionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CspmMonitorAwsOfferingNativeCloudConnection)(nil)).Elem()
}

func (i *cspmMonitorAwsOfferingNativeCloudConnectionPtrType) ToCspmMonitorAwsOfferingNativeCloudConnectionPtrOutput() CspmMonitorAwsOfferingNativeCloudConnectionPtrOutput {
	return i.ToCspmMonitorAwsOfferingNativeCloudConnectionPtrOutputWithContext(context.Background())
}

func (i *cspmMonitorAwsOfferingNativeCloudConnectionPtrType) ToCspmMonitorAwsOfferingNativeCloudConnectionPtrOutputWithContext(ctx context.Context) CspmMonitorAwsOfferingNativeCloudConnectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CspmMonitorAwsOfferingNativeCloudConnectionPtrOutput)
}

type CspmMonitorAwsOfferingNativeCloudConnectionOutput struct{ *pulumi.OutputState }

func (CspmMonitorAwsOfferingNativeCloudConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CspmMonitorAwsOfferingNativeCloudConnection)(nil)).Elem()
}

func (o CspmMonitorAwsOfferingNativeCloudConnectionOutput) ToCspmMonitorAwsOfferingNativeCloudConnectionOutput() CspmMonitorAwsOfferingNativeCloudConnectionOutput {
	return o
}

func (o CspmMonitorAwsOfferingNativeCloudConnectionOutput) ToCspmMonitorAwsOfferingNativeCloudConnectionOutputWithContext(ctx context.Context) CspmMonitorAwsOfferingNativeCloudConnectionOutput {
	return o
}

func (o CspmMonitorAwsOfferingNativeCloudConnectionOutput) ToCspmMonitorAwsOfferingNativeCloudConnectionPtrOutput() CspmMonitorAwsOfferingNativeCloudConnectionPtrOutput {
	return o.ToCspmMonitorAwsOfferingNativeCloudConnectionPtrOutputWithContext(context.Background())
}

func (o CspmMonitorAwsOfferingNativeCloudConnectionOutput) ToCspmMonitorAwsOfferingNativeCloudConnectionPtrOutputWithContext(ctx context.Context) CspmMonitorAwsOfferingNativeCloudConnectionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CspmMonitorAwsOfferingNativeCloudConnection) *CspmMonitorAwsOfferingNativeCloudConnection {
		return &v
	}).(CspmMonitorAwsOfferingNativeCloudConnectionPtrOutput)
}

func (o CspmMonitorAwsOfferingNativeCloudConnectionOutput) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CspmMonitorAwsOfferingNativeCloudConnection) *string { return v.CloudRoleArn }).(pulumi.StringPtrOutput)
}

type CspmMonitorAwsOfferingNativeCloudConnectionPtrOutput struct{ *pulumi.OutputState }

func (CspmMonitorAwsOfferingNativeCloudConnectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CspmMonitorAwsOfferingNativeCloudConnection)(nil)).Elem()
}

func (o CspmMonitorAwsOfferingNativeCloudConnectionPtrOutput) ToCspmMonitorAwsOfferingNativeCloudConnectionPtrOutput() CspmMonitorAwsOfferingNativeCloudConnectionPtrOutput {
	return o
}

func (o CspmMonitorAwsOfferingNativeCloudConnectionPtrOutput) ToCspmMonitorAwsOfferingNativeCloudConnectionPtrOutputWithContext(ctx context.Context) CspmMonitorAwsOfferingNativeCloudConnectionPtrOutput {
	return o
}

func (o CspmMonitorAwsOfferingNativeCloudConnectionPtrOutput) Elem() CspmMonitorAwsOfferingNativeCloudConnectionOutput {
	return o.ApplyT(func(v *CspmMonitorAwsOfferingNativeCloudConnection) CspmMonitorAwsOfferingNativeCloudConnection {
		if v != nil {
			return *v
		}
		var ret CspmMonitorAwsOfferingNativeCloudConnection
		return ret
	}).(CspmMonitorAwsOfferingNativeCloudConnectionOutput)
}

func (o CspmMonitorAwsOfferingNativeCloudConnectionPtrOutput) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CspmMonitorAwsOfferingNativeCloudConnection) *string {
		if v == nil {
			return nil
		}
		return v.CloudRoleArn
	}).(pulumi.StringPtrOutput)
}

type CspmMonitorAwsOfferingResponse struct {
	Description           string                                               `pulumi:"description"`
	NativeCloudConnection *CspmMonitorAwsOfferingResponseNativeCloudConnection `pulumi:"nativeCloudConnection"`
	OfferingType          string                                               `pulumi:"offeringType"`
}





type CspmMonitorAwsOfferingResponseInput interface {
	pulumi.Input

	ToCspmMonitorAwsOfferingResponseOutput() CspmMonitorAwsOfferingResponseOutput
	ToCspmMonitorAwsOfferingResponseOutputWithContext(context.Context) CspmMonitorAwsOfferingResponseOutput
}

type CspmMonitorAwsOfferingResponseArgs struct {
	Description           pulumi.StringInput                                          `pulumi:"description"`
	NativeCloudConnection CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrInput `pulumi:"nativeCloudConnection"`
	OfferingType          pulumi.StringInput                                          `pulumi:"offeringType"`
}

func (CspmMonitorAwsOfferingResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CspmMonitorAwsOfferingResponse)(nil)).Elem()
}

func (i CspmMonitorAwsOfferingResponseArgs) ToCspmMonitorAwsOfferingResponseOutput() CspmMonitorAwsOfferingResponseOutput {
	return i.ToCspmMonitorAwsOfferingResponseOutputWithContext(context.Background())
}

func (i CspmMonitorAwsOfferingResponseArgs) ToCspmMonitorAwsOfferingResponseOutputWithContext(ctx context.Context) CspmMonitorAwsOfferingResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CspmMonitorAwsOfferingResponseOutput)
}

type CspmMonitorAwsOfferingResponseOutput struct{ *pulumi.OutputState }

func (CspmMonitorAwsOfferingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CspmMonitorAwsOfferingResponse)(nil)).Elem()
}

func (o CspmMonitorAwsOfferingResponseOutput) ToCspmMonitorAwsOfferingResponseOutput() CspmMonitorAwsOfferingResponseOutput {
	return o
}

func (o CspmMonitorAwsOfferingResponseOutput) ToCspmMonitorAwsOfferingResponseOutputWithContext(ctx context.Context) CspmMonitorAwsOfferingResponseOutput {
	return o
}

func (o CspmMonitorAwsOfferingResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v CspmMonitorAwsOfferingResponse) string { return v.Description }).(pulumi.StringOutput)
}

func (o CspmMonitorAwsOfferingResponseOutput) NativeCloudConnection() CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput {
	return o.ApplyT(func(v CspmMonitorAwsOfferingResponse) *CspmMonitorAwsOfferingResponseNativeCloudConnection {
		return v.NativeCloudConnection
	}).(CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput)
}

func (o CspmMonitorAwsOfferingResponseOutput) OfferingType() pulumi.StringOutput {
	return o.ApplyT(func(v CspmMonitorAwsOfferingResponse) string { return v.OfferingType }).(pulumi.StringOutput)
}

type CspmMonitorAwsOfferingResponseNativeCloudConnection struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}





type CspmMonitorAwsOfferingResponseNativeCloudConnectionInput interface {
	pulumi.Input

	ToCspmMonitorAwsOfferingResponseNativeCloudConnectionOutput() CspmMonitorAwsOfferingResponseNativeCloudConnectionOutput
	ToCspmMonitorAwsOfferingResponseNativeCloudConnectionOutputWithContext(context.Context) CspmMonitorAwsOfferingResponseNativeCloudConnectionOutput
}

type CspmMonitorAwsOfferingResponseNativeCloudConnectionArgs struct {
	CloudRoleArn pulumi.StringPtrInput `pulumi:"cloudRoleArn"`
}

func (CspmMonitorAwsOfferingResponseNativeCloudConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CspmMonitorAwsOfferingResponseNativeCloudConnection)(nil)).Elem()
}

func (i CspmMonitorAwsOfferingResponseNativeCloudConnectionArgs) ToCspmMonitorAwsOfferingResponseNativeCloudConnectionOutput() CspmMonitorAwsOfferingResponseNativeCloudConnectionOutput {
	return i.ToCspmMonitorAwsOfferingResponseNativeCloudConnectionOutputWithContext(context.Background())
}

func (i CspmMonitorAwsOfferingResponseNativeCloudConnectionArgs) ToCspmMonitorAwsOfferingResponseNativeCloudConnectionOutputWithContext(ctx context.Context) CspmMonitorAwsOfferingResponseNativeCloudConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CspmMonitorAwsOfferingResponseNativeCloudConnectionOutput)
}

func (i CspmMonitorAwsOfferingResponseNativeCloudConnectionArgs) ToCspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput() CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput {
	return i.ToCspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutputWithContext(context.Background())
}

func (i CspmMonitorAwsOfferingResponseNativeCloudConnectionArgs) ToCspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutputWithContext(ctx context.Context) CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CspmMonitorAwsOfferingResponseNativeCloudConnectionOutput).ToCspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutputWithContext(ctx)
}









type CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrInput interface {
	pulumi.Input

	ToCspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput() CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput
	ToCspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutputWithContext(context.Context) CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput
}

type cspmMonitorAwsOfferingResponseNativeCloudConnectionPtrType CspmMonitorAwsOfferingResponseNativeCloudConnectionArgs

func CspmMonitorAwsOfferingResponseNativeCloudConnectionPtr(v *CspmMonitorAwsOfferingResponseNativeCloudConnectionArgs) CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrInput {
	return (*cspmMonitorAwsOfferingResponseNativeCloudConnectionPtrType)(v)
}

func (*cspmMonitorAwsOfferingResponseNativeCloudConnectionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CspmMonitorAwsOfferingResponseNativeCloudConnection)(nil)).Elem()
}

func (i *cspmMonitorAwsOfferingResponseNativeCloudConnectionPtrType) ToCspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput() CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput {
	return i.ToCspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutputWithContext(context.Background())
}

func (i *cspmMonitorAwsOfferingResponseNativeCloudConnectionPtrType) ToCspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutputWithContext(ctx context.Context) CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput)
}

type CspmMonitorAwsOfferingResponseNativeCloudConnectionOutput struct{ *pulumi.OutputState }

func (CspmMonitorAwsOfferingResponseNativeCloudConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CspmMonitorAwsOfferingResponseNativeCloudConnection)(nil)).Elem()
}

func (o CspmMonitorAwsOfferingResponseNativeCloudConnectionOutput) ToCspmMonitorAwsOfferingResponseNativeCloudConnectionOutput() CspmMonitorAwsOfferingResponseNativeCloudConnectionOutput {
	return o
}

func (o CspmMonitorAwsOfferingResponseNativeCloudConnectionOutput) ToCspmMonitorAwsOfferingResponseNativeCloudConnectionOutputWithContext(ctx context.Context) CspmMonitorAwsOfferingResponseNativeCloudConnectionOutput {
	return o
}

func (o CspmMonitorAwsOfferingResponseNativeCloudConnectionOutput) ToCspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput() CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput {
	return o.ToCspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutputWithContext(context.Background())
}

func (o CspmMonitorAwsOfferingResponseNativeCloudConnectionOutput) ToCspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutputWithContext(ctx context.Context) CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CspmMonitorAwsOfferingResponseNativeCloudConnection) *CspmMonitorAwsOfferingResponseNativeCloudConnection {
		return &v
	}).(CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput)
}

func (o CspmMonitorAwsOfferingResponseNativeCloudConnectionOutput) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CspmMonitorAwsOfferingResponseNativeCloudConnection) *string { return v.CloudRoleArn }).(pulumi.StringPtrOutput)
}

type CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput struct{ *pulumi.OutputState }

func (CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CspmMonitorAwsOfferingResponseNativeCloudConnection)(nil)).Elem()
}

func (o CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput) ToCspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput() CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput {
	return o
}

func (o CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput) ToCspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutputWithContext(ctx context.Context) CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput {
	return o
}

func (o CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput) Elem() CspmMonitorAwsOfferingResponseNativeCloudConnectionOutput {
	return o.ApplyT(func(v *CspmMonitorAwsOfferingResponseNativeCloudConnection) CspmMonitorAwsOfferingResponseNativeCloudConnection {
		if v != nil {
			return *v
		}
		var ret CspmMonitorAwsOfferingResponseNativeCloudConnection
		return ret
	}).(CspmMonitorAwsOfferingResponseNativeCloudConnectionOutput)
}

func (o CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CspmMonitorAwsOfferingResponseNativeCloudConnection) *string {
		if v == nil {
			return nil
		}
		return v.CloudRoleArn
	}).(pulumi.StringPtrOutput)
}

type DefenderForContainersAwsOffering struct {
	CloudWatchToKinesis   *DefenderForContainersAwsOfferingCloudWatchToKinesis   `pulumi:"cloudWatchToKinesis"`
	KinesisToS3           *DefenderForContainersAwsOfferingKinesisToS3           `pulumi:"kinesisToS3"`
	KubernetesScubaReader *DefenderForContainersAwsOfferingKubernetesScubaReader `pulumi:"kubernetesScubaReader"`
	KubernetesService     *DefenderForContainersAwsOfferingKubernetesService     `pulumi:"kubernetesService"`
	OfferingType          string                                                 `pulumi:"offeringType"`
}





type DefenderForContainersAwsOfferingInput interface {
	pulumi.Input

	ToDefenderForContainersAwsOfferingOutput() DefenderForContainersAwsOfferingOutput
	ToDefenderForContainersAwsOfferingOutputWithContext(context.Context) DefenderForContainersAwsOfferingOutput
}

type DefenderForContainersAwsOfferingArgs struct {
	CloudWatchToKinesis   DefenderForContainersAwsOfferingCloudWatchToKinesisPtrInput   `pulumi:"cloudWatchToKinesis"`
	KinesisToS3           DefenderForContainersAwsOfferingKinesisToS3PtrInput           `pulumi:"kinesisToS3"`
	KubernetesScubaReader DefenderForContainersAwsOfferingKubernetesScubaReaderPtrInput `pulumi:"kubernetesScubaReader"`
	KubernetesService     DefenderForContainersAwsOfferingKubernetesServicePtrInput     `pulumi:"kubernetesService"`
	OfferingType          pulumi.StringInput                                            `pulumi:"offeringType"`
}

func (DefenderForContainersAwsOfferingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForContainersAwsOffering)(nil)).Elem()
}

func (i DefenderForContainersAwsOfferingArgs) ToDefenderForContainersAwsOfferingOutput() DefenderForContainersAwsOfferingOutput {
	return i.ToDefenderForContainersAwsOfferingOutputWithContext(context.Background())
}

func (i DefenderForContainersAwsOfferingArgs) ToDefenderForContainersAwsOfferingOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingOutput)
}

type DefenderForContainersAwsOfferingOutput struct{ *pulumi.OutputState }

func (DefenderForContainersAwsOfferingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForContainersAwsOffering)(nil)).Elem()
}

func (o DefenderForContainersAwsOfferingOutput) ToDefenderForContainersAwsOfferingOutput() DefenderForContainersAwsOfferingOutput {
	return o
}

func (o DefenderForContainersAwsOfferingOutput) ToDefenderForContainersAwsOfferingOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingOutput {
	return o
}

func (o DefenderForContainersAwsOfferingOutput) CloudWatchToKinesis() DefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput {
	return o.ApplyT(func(v DefenderForContainersAwsOffering) *DefenderForContainersAwsOfferingCloudWatchToKinesis {
		return v.CloudWatchToKinesis
	}).(DefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput)
}

func (o DefenderForContainersAwsOfferingOutput) KinesisToS3() DefenderForContainersAwsOfferingKinesisToS3PtrOutput {
	return o.ApplyT(func(v DefenderForContainersAwsOffering) *DefenderForContainersAwsOfferingKinesisToS3 {
		return v.KinesisToS3
	}).(DefenderForContainersAwsOfferingKinesisToS3PtrOutput)
}

func (o DefenderForContainersAwsOfferingOutput) KubernetesScubaReader() DefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput {
	return o.ApplyT(func(v DefenderForContainersAwsOffering) *DefenderForContainersAwsOfferingKubernetesScubaReader {
		return v.KubernetesScubaReader
	}).(DefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput)
}

func (o DefenderForContainersAwsOfferingOutput) KubernetesService() DefenderForContainersAwsOfferingKubernetesServicePtrOutput {
	return o.ApplyT(func(v DefenderForContainersAwsOffering) *DefenderForContainersAwsOfferingKubernetesService {
		return v.KubernetesService
	}).(DefenderForContainersAwsOfferingKubernetesServicePtrOutput)
}

func (o DefenderForContainersAwsOfferingOutput) OfferingType() pulumi.StringOutput {
	return o.ApplyT(func(v DefenderForContainersAwsOffering) string { return v.OfferingType }).(pulumi.StringOutput)
}

type DefenderForContainersAwsOfferingCloudWatchToKinesis struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}





type DefenderForContainersAwsOfferingCloudWatchToKinesisInput interface {
	pulumi.Input

	ToDefenderForContainersAwsOfferingCloudWatchToKinesisOutput() DefenderForContainersAwsOfferingCloudWatchToKinesisOutput
	ToDefenderForContainersAwsOfferingCloudWatchToKinesisOutputWithContext(context.Context) DefenderForContainersAwsOfferingCloudWatchToKinesisOutput
}

type DefenderForContainersAwsOfferingCloudWatchToKinesisArgs struct {
	CloudRoleArn pulumi.StringPtrInput `pulumi:"cloudRoleArn"`
}

func (DefenderForContainersAwsOfferingCloudWatchToKinesisArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForContainersAwsOfferingCloudWatchToKinesis)(nil)).Elem()
}

func (i DefenderForContainersAwsOfferingCloudWatchToKinesisArgs) ToDefenderForContainersAwsOfferingCloudWatchToKinesisOutput() DefenderForContainersAwsOfferingCloudWatchToKinesisOutput {
	return i.ToDefenderForContainersAwsOfferingCloudWatchToKinesisOutputWithContext(context.Background())
}

func (i DefenderForContainersAwsOfferingCloudWatchToKinesisArgs) ToDefenderForContainersAwsOfferingCloudWatchToKinesisOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingCloudWatchToKinesisOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingCloudWatchToKinesisOutput)
}

func (i DefenderForContainersAwsOfferingCloudWatchToKinesisArgs) ToDefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput() DefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput {
	return i.ToDefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutputWithContext(context.Background())
}

func (i DefenderForContainersAwsOfferingCloudWatchToKinesisArgs) ToDefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingCloudWatchToKinesisOutput).ToDefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutputWithContext(ctx)
}









type DefenderForContainersAwsOfferingCloudWatchToKinesisPtrInput interface {
	pulumi.Input

	ToDefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput() DefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput
	ToDefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutputWithContext(context.Context) DefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput
}

type defenderForContainersAwsOfferingCloudWatchToKinesisPtrType DefenderForContainersAwsOfferingCloudWatchToKinesisArgs

func DefenderForContainersAwsOfferingCloudWatchToKinesisPtr(v *DefenderForContainersAwsOfferingCloudWatchToKinesisArgs) DefenderForContainersAwsOfferingCloudWatchToKinesisPtrInput {
	return (*defenderForContainersAwsOfferingCloudWatchToKinesisPtrType)(v)
}

func (*defenderForContainersAwsOfferingCloudWatchToKinesisPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForContainersAwsOfferingCloudWatchToKinesis)(nil)).Elem()
}

func (i *defenderForContainersAwsOfferingCloudWatchToKinesisPtrType) ToDefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput() DefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput {
	return i.ToDefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutputWithContext(context.Background())
}

func (i *defenderForContainersAwsOfferingCloudWatchToKinesisPtrType) ToDefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput)
}

type DefenderForContainersAwsOfferingCloudWatchToKinesisOutput struct{ *pulumi.OutputState }

func (DefenderForContainersAwsOfferingCloudWatchToKinesisOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForContainersAwsOfferingCloudWatchToKinesis)(nil)).Elem()
}

func (o DefenderForContainersAwsOfferingCloudWatchToKinesisOutput) ToDefenderForContainersAwsOfferingCloudWatchToKinesisOutput() DefenderForContainersAwsOfferingCloudWatchToKinesisOutput {
	return o
}

func (o DefenderForContainersAwsOfferingCloudWatchToKinesisOutput) ToDefenderForContainersAwsOfferingCloudWatchToKinesisOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingCloudWatchToKinesisOutput {
	return o
}

func (o DefenderForContainersAwsOfferingCloudWatchToKinesisOutput) ToDefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput() DefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput {
	return o.ToDefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutputWithContext(context.Background())
}

func (o DefenderForContainersAwsOfferingCloudWatchToKinesisOutput) ToDefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefenderForContainersAwsOfferingCloudWatchToKinesis) *DefenderForContainersAwsOfferingCloudWatchToKinesis {
		return &v
	}).(DefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput)
}

func (o DefenderForContainersAwsOfferingCloudWatchToKinesisOutput) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DefenderForContainersAwsOfferingCloudWatchToKinesis) *string { return v.CloudRoleArn }).(pulumi.StringPtrOutput)
}

type DefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput struct{ *pulumi.OutputState }

func (DefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForContainersAwsOfferingCloudWatchToKinesis)(nil)).Elem()
}

func (o DefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput) ToDefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput() DefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput {
	return o
}

func (o DefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput) ToDefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput {
	return o
}

func (o DefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput) Elem() DefenderForContainersAwsOfferingCloudWatchToKinesisOutput {
	return o.ApplyT(func(v *DefenderForContainersAwsOfferingCloudWatchToKinesis) DefenderForContainersAwsOfferingCloudWatchToKinesis {
		if v != nil {
			return *v
		}
		var ret DefenderForContainersAwsOfferingCloudWatchToKinesis
		return ret
	}).(DefenderForContainersAwsOfferingCloudWatchToKinesisOutput)
}

func (o DefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefenderForContainersAwsOfferingCloudWatchToKinesis) *string {
		if v == nil {
			return nil
		}
		return v.CloudRoleArn
	}).(pulumi.StringPtrOutput)
}

type DefenderForContainersAwsOfferingKinesisToS3 struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}





type DefenderForContainersAwsOfferingKinesisToS3Input interface {
	pulumi.Input

	ToDefenderForContainersAwsOfferingKinesisToS3Output() DefenderForContainersAwsOfferingKinesisToS3Output
	ToDefenderForContainersAwsOfferingKinesisToS3OutputWithContext(context.Context) DefenderForContainersAwsOfferingKinesisToS3Output
}

type DefenderForContainersAwsOfferingKinesisToS3Args struct {
	CloudRoleArn pulumi.StringPtrInput `pulumi:"cloudRoleArn"`
}

func (DefenderForContainersAwsOfferingKinesisToS3Args) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForContainersAwsOfferingKinesisToS3)(nil)).Elem()
}

func (i DefenderForContainersAwsOfferingKinesisToS3Args) ToDefenderForContainersAwsOfferingKinesisToS3Output() DefenderForContainersAwsOfferingKinesisToS3Output {
	return i.ToDefenderForContainersAwsOfferingKinesisToS3OutputWithContext(context.Background())
}

func (i DefenderForContainersAwsOfferingKinesisToS3Args) ToDefenderForContainersAwsOfferingKinesisToS3OutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingKinesisToS3Output {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingKinesisToS3Output)
}

func (i DefenderForContainersAwsOfferingKinesisToS3Args) ToDefenderForContainersAwsOfferingKinesisToS3PtrOutput() DefenderForContainersAwsOfferingKinesisToS3PtrOutput {
	return i.ToDefenderForContainersAwsOfferingKinesisToS3PtrOutputWithContext(context.Background())
}

func (i DefenderForContainersAwsOfferingKinesisToS3Args) ToDefenderForContainersAwsOfferingKinesisToS3PtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingKinesisToS3PtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingKinesisToS3Output).ToDefenderForContainersAwsOfferingKinesisToS3PtrOutputWithContext(ctx)
}









type DefenderForContainersAwsOfferingKinesisToS3PtrInput interface {
	pulumi.Input

	ToDefenderForContainersAwsOfferingKinesisToS3PtrOutput() DefenderForContainersAwsOfferingKinesisToS3PtrOutput
	ToDefenderForContainersAwsOfferingKinesisToS3PtrOutputWithContext(context.Context) DefenderForContainersAwsOfferingKinesisToS3PtrOutput
}

type defenderForContainersAwsOfferingKinesisToS3PtrType DefenderForContainersAwsOfferingKinesisToS3Args

func DefenderForContainersAwsOfferingKinesisToS3Ptr(v *DefenderForContainersAwsOfferingKinesisToS3Args) DefenderForContainersAwsOfferingKinesisToS3PtrInput {
	return (*defenderForContainersAwsOfferingKinesisToS3PtrType)(v)
}

func (*defenderForContainersAwsOfferingKinesisToS3PtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForContainersAwsOfferingKinesisToS3)(nil)).Elem()
}

func (i *defenderForContainersAwsOfferingKinesisToS3PtrType) ToDefenderForContainersAwsOfferingKinesisToS3PtrOutput() DefenderForContainersAwsOfferingKinesisToS3PtrOutput {
	return i.ToDefenderForContainersAwsOfferingKinesisToS3PtrOutputWithContext(context.Background())
}

func (i *defenderForContainersAwsOfferingKinesisToS3PtrType) ToDefenderForContainersAwsOfferingKinesisToS3PtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingKinesisToS3PtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingKinesisToS3PtrOutput)
}

type DefenderForContainersAwsOfferingKinesisToS3Output struct{ *pulumi.OutputState }

func (DefenderForContainersAwsOfferingKinesisToS3Output) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForContainersAwsOfferingKinesisToS3)(nil)).Elem()
}

func (o DefenderForContainersAwsOfferingKinesisToS3Output) ToDefenderForContainersAwsOfferingKinesisToS3Output() DefenderForContainersAwsOfferingKinesisToS3Output {
	return o
}

func (o DefenderForContainersAwsOfferingKinesisToS3Output) ToDefenderForContainersAwsOfferingKinesisToS3OutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingKinesisToS3Output {
	return o
}

func (o DefenderForContainersAwsOfferingKinesisToS3Output) ToDefenderForContainersAwsOfferingKinesisToS3PtrOutput() DefenderForContainersAwsOfferingKinesisToS3PtrOutput {
	return o.ToDefenderForContainersAwsOfferingKinesisToS3PtrOutputWithContext(context.Background())
}

func (o DefenderForContainersAwsOfferingKinesisToS3Output) ToDefenderForContainersAwsOfferingKinesisToS3PtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingKinesisToS3PtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefenderForContainersAwsOfferingKinesisToS3) *DefenderForContainersAwsOfferingKinesisToS3 {
		return &v
	}).(DefenderForContainersAwsOfferingKinesisToS3PtrOutput)
}

func (o DefenderForContainersAwsOfferingKinesisToS3Output) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DefenderForContainersAwsOfferingKinesisToS3) *string { return v.CloudRoleArn }).(pulumi.StringPtrOutput)
}

type DefenderForContainersAwsOfferingKinesisToS3PtrOutput struct{ *pulumi.OutputState }

func (DefenderForContainersAwsOfferingKinesisToS3PtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForContainersAwsOfferingKinesisToS3)(nil)).Elem()
}

func (o DefenderForContainersAwsOfferingKinesisToS3PtrOutput) ToDefenderForContainersAwsOfferingKinesisToS3PtrOutput() DefenderForContainersAwsOfferingKinesisToS3PtrOutput {
	return o
}

func (o DefenderForContainersAwsOfferingKinesisToS3PtrOutput) ToDefenderForContainersAwsOfferingKinesisToS3PtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingKinesisToS3PtrOutput {
	return o
}

func (o DefenderForContainersAwsOfferingKinesisToS3PtrOutput) Elem() DefenderForContainersAwsOfferingKinesisToS3Output {
	return o.ApplyT(func(v *DefenderForContainersAwsOfferingKinesisToS3) DefenderForContainersAwsOfferingKinesisToS3 {
		if v != nil {
			return *v
		}
		var ret DefenderForContainersAwsOfferingKinesisToS3
		return ret
	}).(DefenderForContainersAwsOfferingKinesisToS3Output)
}

func (o DefenderForContainersAwsOfferingKinesisToS3PtrOutput) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefenderForContainersAwsOfferingKinesisToS3) *string {
		if v == nil {
			return nil
		}
		return v.CloudRoleArn
	}).(pulumi.StringPtrOutput)
}

type DefenderForContainersAwsOfferingKubernetesScubaReader struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}





type DefenderForContainersAwsOfferingKubernetesScubaReaderInput interface {
	pulumi.Input

	ToDefenderForContainersAwsOfferingKubernetesScubaReaderOutput() DefenderForContainersAwsOfferingKubernetesScubaReaderOutput
	ToDefenderForContainersAwsOfferingKubernetesScubaReaderOutputWithContext(context.Context) DefenderForContainersAwsOfferingKubernetesScubaReaderOutput
}

type DefenderForContainersAwsOfferingKubernetesScubaReaderArgs struct {
	CloudRoleArn pulumi.StringPtrInput `pulumi:"cloudRoleArn"`
}

func (DefenderForContainersAwsOfferingKubernetesScubaReaderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForContainersAwsOfferingKubernetesScubaReader)(nil)).Elem()
}

func (i DefenderForContainersAwsOfferingKubernetesScubaReaderArgs) ToDefenderForContainersAwsOfferingKubernetesScubaReaderOutput() DefenderForContainersAwsOfferingKubernetesScubaReaderOutput {
	return i.ToDefenderForContainersAwsOfferingKubernetesScubaReaderOutputWithContext(context.Background())
}

func (i DefenderForContainersAwsOfferingKubernetesScubaReaderArgs) ToDefenderForContainersAwsOfferingKubernetesScubaReaderOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingKubernetesScubaReaderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingKubernetesScubaReaderOutput)
}

func (i DefenderForContainersAwsOfferingKubernetesScubaReaderArgs) ToDefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput() DefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput {
	return i.ToDefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutputWithContext(context.Background())
}

func (i DefenderForContainersAwsOfferingKubernetesScubaReaderArgs) ToDefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingKubernetesScubaReaderOutput).ToDefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutputWithContext(ctx)
}









type DefenderForContainersAwsOfferingKubernetesScubaReaderPtrInput interface {
	pulumi.Input

	ToDefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput() DefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput
	ToDefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutputWithContext(context.Context) DefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput
}

type defenderForContainersAwsOfferingKubernetesScubaReaderPtrType DefenderForContainersAwsOfferingKubernetesScubaReaderArgs

func DefenderForContainersAwsOfferingKubernetesScubaReaderPtr(v *DefenderForContainersAwsOfferingKubernetesScubaReaderArgs) DefenderForContainersAwsOfferingKubernetesScubaReaderPtrInput {
	return (*defenderForContainersAwsOfferingKubernetesScubaReaderPtrType)(v)
}

func (*defenderForContainersAwsOfferingKubernetesScubaReaderPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForContainersAwsOfferingKubernetesScubaReader)(nil)).Elem()
}

func (i *defenderForContainersAwsOfferingKubernetesScubaReaderPtrType) ToDefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput() DefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput {
	return i.ToDefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutputWithContext(context.Background())
}

func (i *defenderForContainersAwsOfferingKubernetesScubaReaderPtrType) ToDefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput)
}

type DefenderForContainersAwsOfferingKubernetesScubaReaderOutput struct{ *pulumi.OutputState }

func (DefenderForContainersAwsOfferingKubernetesScubaReaderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForContainersAwsOfferingKubernetesScubaReader)(nil)).Elem()
}

func (o DefenderForContainersAwsOfferingKubernetesScubaReaderOutput) ToDefenderForContainersAwsOfferingKubernetesScubaReaderOutput() DefenderForContainersAwsOfferingKubernetesScubaReaderOutput {
	return o
}

func (o DefenderForContainersAwsOfferingKubernetesScubaReaderOutput) ToDefenderForContainersAwsOfferingKubernetesScubaReaderOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingKubernetesScubaReaderOutput {
	return o
}

func (o DefenderForContainersAwsOfferingKubernetesScubaReaderOutput) ToDefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput() DefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput {
	return o.ToDefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutputWithContext(context.Background())
}

func (o DefenderForContainersAwsOfferingKubernetesScubaReaderOutput) ToDefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefenderForContainersAwsOfferingKubernetesScubaReader) *DefenderForContainersAwsOfferingKubernetesScubaReader {
		return &v
	}).(DefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput)
}

func (o DefenderForContainersAwsOfferingKubernetesScubaReaderOutput) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DefenderForContainersAwsOfferingKubernetesScubaReader) *string { return v.CloudRoleArn }).(pulumi.StringPtrOutput)
}

type DefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput struct{ *pulumi.OutputState }

func (DefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForContainersAwsOfferingKubernetesScubaReader)(nil)).Elem()
}

func (o DefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput) ToDefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput() DefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput {
	return o
}

func (o DefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput) ToDefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput {
	return o
}

func (o DefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput) Elem() DefenderForContainersAwsOfferingKubernetesScubaReaderOutput {
	return o.ApplyT(func(v *DefenderForContainersAwsOfferingKubernetesScubaReader) DefenderForContainersAwsOfferingKubernetesScubaReader {
		if v != nil {
			return *v
		}
		var ret DefenderForContainersAwsOfferingKubernetesScubaReader
		return ret
	}).(DefenderForContainersAwsOfferingKubernetesScubaReaderOutput)
}

func (o DefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefenderForContainersAwsOfferingKubernetesScubaReader) *string {
		if v == nil {
			return nil
		}
		return v.CloudRoleArn
	}).(pulumi.StringPtrOutput)
}

type DefenderForContainersAwsOfferingKubernetesService struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}





type DefenderForContainersAwsOfferingKubernetesServiceInput interface {
	pulumi.Input

	ToDefenderForContainersAwsOfferingKubernetesServiceOutput() DefenderForContainersAwsOfferingKubernetesServiceOutput
	ToDefenderForContainersAwsOfferingKubernetesServiceOutputWithContext(context.Context) DefenderForContainersAwsOfferingKubernetesServiceOutput
}

type DefenderForContainersAwsOfferingKubernetesServiceArgs struct {
	CloudRoleArn pulumi.StringPtrInput `pulumi:"cloudRoleArn"`
}

func (DefenderForContainersAwsOfferingKubernetesServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForContainersAwsOfferingKubernetesService)(nil)).Elem()
}

func (i DefenderForContainersAwsOfferingKubernetesServiceArgs) ToDefenderForContainersAwsOfferingKubernetesServiceOutput() DefenderForContainersAwsOfferingKubernetesServiceOutput {
	return i.ToDefenderForContainersAwsOfferingKubernetesServiceOutputWithContext(context.Background())
}

func (i DefenderForContainersAwsOfferingKubernetesServiceArgs) ToDefenderForContainersAwsOfferingKubernetesServiceOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingKubernetesServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingKubernetesServiceOutput)
}

func (i DefenderForContainersAwsOfferingKubernetesServiceArgs) ToDefenderForContainersAwsOfferingKubernetesServicePtrOutput() DefenderForContainersAwsOfferingKubernetesServicePtrOutput {
	return i.ToDefenderForContainersAwsOfferingKubernetesServicePtrOutputWithContext(context.Background())
}

func (i DefenderForContainersAwsOfferingKubernetesServiceArgs) ToDefenderForContainersAwsOfferingKubernetesServicePtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingKubernetesServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingKubernetesServiceOutput).ToDefenderForContainersAwsOfferingKubernetesServicePtrOutputWithContext(ctx)
}









type DefenderForContainersAwsOfferingKubernetesServicePtrInput interface {
	pulumi.Input

	ToDefenderForContainersAwsOfferingKubernetesServicePtrOutput() DefenderForContainersAwsOfferingKubernetesServicePtrOutput
	ToDefenderForContainersAwsOfferingKubernetesServicePtrOutputWithContext(context.Context) DefenderForContainersAwsOfferingKubernetesServicePtrOutput
}

type defenderForContainersAwsOfferingKubernetesServicePtrType DefenderForContainersAwsOfferingKubernetesServiceArgs

func DefenderForContainersAwsOfferingKubernetesServicePtr(v *DefenderForContainersAwsOfferingKubernetesServiceArgs) DefenderForContainersAwsOfferingKubernetesServicePtrInput {
	return (*defenderForContainersAwsOfferingKubernetesServicePtrType)(v)
}

func (*defenderForContainersAwsOfferingKubernetesServicePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForContainersAwsOfferingKubernetesService)(nil)).Elem()
}

func (i *defenderForContainersAwsOfferingKubernetesServicePtrType) ToDefenderForContainersAwsOfferingKubernetesServicePtrOutput() DefenderForContainersAwsOfferingKubernetesServicePtrOutput {
	return i.ToDefenderForContainersAwsOfferingKubernetesServicePtrOutputWithContext(context.Background())
}

func (i *defenderForContainersAwsOfferingKubernetesServicePtrType) ToDefenderForContainersAwsOfferingKubernetesServicePtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingKubernetesServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingKubernetesServicePtrOutput)
}

type DefenderForContainersAwsOfferingKubernetesServiceOutput struct{ *pulumi.OutputState }

func (DefenderForContainersAwsOfferingKubernetesServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForContainersAwsOfferingKubernetesService)(nil)).Elem()
}

func (o DefenderForContainersAwsOfferingKubernetesServiceOutput) ToDefenderForContainersAwsOfferingKubernetesServiceOutput() DefenderForContainersAwsOfferingKubernetesServiceOutput {
	return o
}

func (o DefenderForContainersAwsOfferingKubernetesServiceOutput) ToDefenderForContainersAwsOfferingKubernetesServiceOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingKubernetesServiceOutput {
	return o
}

func (o DefenderForContainersAwsOfferingKubernetesServiceOutput) ToDefenderForContainersAwsOfferingKubernetesServicePtrOutput() DefenderForContainersAwsOfferingKubernetesServicePtrOutput {
	return o.ToDefenderForContainersAwsOfferingKubernetesServicePtrOutputWithContext(context.Background())
}

func (o DefenderForContainersAwsOfferingKubernetesServiceOutput) ToDefenderForContainersAwsOfferingKubernetesServicePtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingKubernetesServicePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefenderForContainersAwsOfferingKubernetesService) *DefenderForContainersAwsOfferingKubernetesService {
		return &v
	}).(DefenderForContainersAwsOfferingKubernetesServicePtrOutput)
}

func (o DefenderForContainersAwsOfferingKubernetesServiceOutput) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DefenderForContainersAwsOfferingKubernetesService) *string { return v.CloudRoleArn }).(pulumi.StringPtrOutput)
}

type DefenderForContainersAwsOfferingKubernetesServicePtrOutput struct{ *pulumi.OutputState }

func (DefenderForContainersAwsOfferingKubernetesServicePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForContainersAwsOfferingKubernetesService)(nil)).Elem()
}

func (o DefenderForContainersAwsOfferingKubernetesServicePtrOutput) ToDefenderForContainersAwsOfferingKubernetesServicePtrOutput() DefenderForContainersAwsOfferingKubernetesServicePtrOutput {
	return o
}

func (o DefenderForContainersAwsOfferingKubernetesServicePtrOutput) ToDefenderForContainersAwsOfferingKubernetesServicePtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingKubernetesServicePtrOutput {
	return o
}

func (o DefenderForContainersAwsOfferingKubernetesServicePtrOutput) Elem() DefenderForContainersAwsOfferingKubernetesServiceOutput {
	return o.ApplyT(func(v *DefenderForContainersAwsOfferingKubernetesService) DefenderForContainersAwsOfferingKubernetesService {
		if v != nil {
			return *v
		}
		var ret DefenderForContainersAwsOfferingKubernetesService
		return ret
	}).(DefenderForContainersAwsOfferingKubernetesServiceOutput)
}

func (o DefenderForContainersAwsOfferingKubernetesServicePtrOutput) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefenderForContainersAwsOfferingKubernetesService) *string {
		if v == nil {
			return nil
		}
		return v.CloudRoleArn
	}).(pulumi.StringPtrOutput)
}

type DefenderForContainersAwsOfferingResponse struct {
	CloudWatchToKinesis   *DefenderForContainersAwsOfferingResponseCloudWatchToKinesis   `pulumi:"cloudWatchToKinesis"`
	Description           string                                                         `pulumi:"description"`
	KinesisToS3           *DefenderForContainersAwsOfferingResponseKinesisToS3           `pulumi:"kinesisToS3"`
	KubernetesScubaReader *DefenderForContainersAwsOfferingResponseKubernetesScubaReader `pulumi:"kubernetesScubaReader"`
	KubernetesService     *DefenderForContainersAwsOfferingResponseKubernetesService     `pulumi:"kubernetesService"`
	OfferingType          string                                                         `pulumi:"offeringType"`
}





type DefenderForContainersAwsOfferingResponseInput interface {
	pulumi.Input

	ToDefenderForContainersAwsOfferingResponseOutput() DefenderForContainersAwsOfferingResponseOutput
	ToDefenderForContainersAwsOfferingResponseOutputWithContext(context.Context) DefenderForContainersAwsOfferingResponseOutput
}

type DefenderForContainersAwsOfferingResponseArgs struct {
	CloudWatchToKinesis   DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrInput   `pulumi:"cloudWatchToKinesis"`
	Description           pulumi.StringInput                                                    `pulumi:"description"`
	KinesisToS3           DefenderForContainersAwsOfferingResponseKinesisToS3PtrInput           `pulumi:"kinesisToS3"`
	KubernetesScubaReader DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrInput `pulumi:"kubernetesScubaReader"`
	KubernetesService     DefenderForContainersAwsOfferingResponseKubernetesServicePtrInput     `pulumi:"kubernetesService"`
	OfferingType          pulumi.StringInput                                                    `pulumi:"offeringType"`
}

func (DefenderForContainersAwsOfferingResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForContainersAwsOfferingResponse)(nil)).Elem()
}

func (i DefenderForContainersAwsOfferingResponseArgs) ToDefenderForContainersAwsOfferingResponseOutput() DefenderForContainersAwsOfferingResponseOutput {
	return i.ToDefenderForContainersAwsOfferingResponseOutputWithContext(context.Background())
}

func (i DefenderForContainersAwsOfferingResponseArgs) ToDefenderForContainersAwsOfferingResponseOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingResponseOutput)
}

type DefenderForContainersAwsOfferingResponseOutput struct{ *pulumi.OutputState }

func (DefenderForContainersAwsOfferingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForContainersAwsOfferingResponse)(nil)).Elem()
}

func (o DefenderForContainersAwsOfferingResponseOutput) ToDefenderForContainersAwsOfferingResponseOutput() DefenderForContainersAwsOfferingResponseOutput {
	return o
}

func (o DefenderForContainersAwsOfferingResponseOutput) ToDefenderForContainersAwsOfferingResponseOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseOutput {
	return o
}

func (o DefenderForContainersAwsOfferingResponseOutput) CloudWatchToKinesis() DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput {
	return o.ApplyT(func(v DefenderForContainersAwsOfferingResponse) *DefenderForContainersAwsOfferingResponseCloudWatchToKinesis {
		return v.CloudWatchToKinesis
	}).(DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput)
}

func (o DefenderForContainersAwsOfferingResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v DefenderForContainersAwsOfferingResponse) string { return v.Description }).(pulumi.StringOutput)
}

func (o DefenderForContainersAwsOfferingResponseOutput) KinesisToS3() DefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput {
	return o.ApplyT(func(v DefenderForContainersAwsOfferingResponse) *DefenderForContainersAwsOfferingResponseKinesisToS3 {
		return v.KinesisToS3
	}).(DefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput)
}

func (o DefenderForContainersAwsOfferingResponseOutput) KubernetesScubaReader() DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput {
	return o.ApplyT(func(v DefenderForContainersAwsOfferingResponse) *DefenderForContainersAwsOfferingResponseKubernetesScubaReader {
		return v.KubernetesScubaReader
	}).(DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput)
}

func (o DefenderForContainersAwsOfferingResponseOutput) KubernetesService() DefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput {
	return o.ApplyT(func(v DefenderForContainersAwsOfferingResponse) *DefenderForContainersAwsOfferingResponseKubernetesService {
		return v.KubernetesService
	}).(DefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput)
}

func (o DefenderForContainersAwsOfferingResponseOutput) OfferingType() pulumi.StringOutput {
	return o.ApplyT(func(v DefenderForContainersAwsOfferingResponse) string { return v.OfferingType }).(pulumi.StringOutput)
}

type DefenderForContainersAwsOfferingResponseCloudWatchToKinesis struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}





type DefenderForContainersAwsOfferingResponseCloudWatchToKinesisInput interface {
	pulumi.Input

	ToDefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutput() DefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutput
	ToDefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutputWithContext(context.Context) DefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutput
}

type DefenderForContainersAwsOfferingResponseCloudWatchToKinesisArgs struct {
	CloudRoleArn pulumi.StringPtrInput `pulumi:"cloudRoleArn"`
}

func (DefenderForContainersAwsOfferingResponseCloudWatchToKinesisArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForContainersAwsOfferingResponseCloudWatchToKinesis)(nil)).Elem()
}

func (i DefenderForContainersAwsOfferingResponseCloudWatchToKinesisArgs) ToDefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutput() DefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutput {
	return i.ToDefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutputWithContext(context.Background())
}

func (i DefenderForContainersAwsOfferingResponseCloudWatchToKinesisArgs) ToDefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutput)
}

func (i DefenderForContainersAwsOfferingResponseCloudWatchToKinesisArgs) ToDefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput() DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput {
	return i.ToDefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutputWithContext(context.Background())
}

func (i DefenderForContainersAwsOfferingResponseCloudWatchToKinesisArgs) ToDefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutput).ToDefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutputWithContext(ctx)
}









type DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrInput interface {
	pulumi.Input

	ToDefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput() DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput
	ToDefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutputWithContext(context.Context) DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput
}

type defenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrType DefenderForContainersAwsOfferingResponseCloudWatchToKinesisArgs

func DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtr(v *DefenderForContainersAwsOfferingResponseCloudWatchToKinesisArgs) DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrInput {
	return (*defenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrType)(v)
}

func (*defenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForContainersAwsOfferingResponseCloudWatchToKinesis)(nil)).Elem()
}

func (i *defenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrType) ToDefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput() DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput {
	return i.ToDefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutputWithContext(context.Background())
}

func (i *defenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrType) ToDefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput)
}

type DefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutput struct{ *pulumi.OutputState }

func (DefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForContainersAwsOfferingResponseCloudWatchToKinesis)(nil)).Elem()
}

func (o DefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutput) ToDefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutput() DefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutput {
	return o
}

func (o DefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutput) ToDefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutput {
	return o
}

func (o DefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutput) ToDefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput() DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput {
	return o.ToDefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutputWithContext(context.Background())
}

func (o DefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutput) ToDefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefenderForContainersAwsOfferingResponseCloudWatchToKinesis) *DefenderForContainersAwsOfferingResponseCloudWatchToKinesis {
		return &v
	}).(DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput)
}

func (o DefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutput) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DefenderForContainersAwsOfferingResponseCloudWatchToKinesis) *string { return v.CloudRoleArn }).(pulumi.StringPtrOutput)
}

type DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput struct{ *pulumi.OutputState }

func (DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForContainersAwsOfferingResponseCloudWatchToKinesis)(nil)).Elem()
}

func (o DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput) ToDefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput() DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput {
	return o
}

func (o DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput) ToDefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput {
	return o
}

func (o DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput) Elem() DefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutput {
	return o.ApplyT(func(v *DefenderForContainersAwsOfferingResponseCloudWatchToKinesis) DefenderForContainersAwsOfferingResponseCloudWatchToKinesis {
		if v != nil {
			return *v
		}
		var ret DefenderForContainersAwsOfferingResponseCloudWatchToKinesis
		return ret
	}).(DefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutput)
}

func (o DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefenderForContainersAwsOfferingResponseCloudWatchToKinesis) *string {
		if v == nil {
			return nil
		}
		return v.CloudRoleArn
	}).(pulumi.StringPtrOutput)
}

type DefenderForContainersAwsOfferingResponseKinesisToS3 struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}





type DefenderForContainersAwsOfferingResponseKinesisToS3Input interface {
	pulumi.Input

	ToDefenderForContainersAwsOfferingResponseKinesisToS3Output() DefenderForContainersAwsOfferingResponseKinesisToS3Output
	ToDefenderForContainersAwsOfferingResponseKinesisToS3OutputWithContext(context.Context) DefenderForContainersAwsOfferingResponseKinesisToS3Output
}

type DefenderForContainersAwsOfferingResponseKinesisToS3Args struct {
	CloudRoleArn pulumi.StringPtrInput `pulumi:"cloudRoleArn"`
}

func (DefenderForContainersAwsOfferingResponseKinesisToS3Args) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForContainersAwsOfferingResponseKinesisToS3)(nil)).Elem()
}

func (i DefenderForContainersAwsOfferingResponseKinesisToS3Args) ToDefenderForContainersAwsOfferingResponseKinesisToS3Output() DefenderForContainersAwsOfferingResponseKinesisToS3Output {
	return i.ToDefenderForContainersAwsOfferingResponseKinesisToS3OutputWithContext(context.Background())
}

func (i DefenderForContainersAwsOfferingResponseKinesisToS3Args) ToDefenderForContainersAwsOfferingResponseKinesisToS3OutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseKinesisToS3Output {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingResponseKinesisToS3Output)
}

func (i DefenderForContainersAwsOfferingResponseKinesisToS3Args) ToDefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput() DefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput {
	return i.ToDefenderForContainersAwsOfferingResponseKinesisToS3PtrOutputWithContext(context.Background())
}

func (i DefenderForContainersAwsOfferingResponseKinesisToS3Args) ToDefenderForContainersAwsOfferingResponseKinesisToS3PtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingResponseKinesisToS3Output).ToDefenderForContainersAwsOfferingResponseKinesisToS3PtrOutputWithContext(ctx)
}









type DefenderForContainersAwsOfferingResponseKinesisToS3PtrInput interface {
	pulumi.Input

	ToDefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput() DefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput
	ToDefenderForContainersAwsOfferingResponseKinesisToS3PtrOutputWithContext(context.Context) DefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput
}

type defenderForContainersAwsOfferingResponseKinesisToS3PtrType DefenderForContainersAwsOfferingResponseKinesisToS3Args

func DefenderForContainersAwsOfferingResponseKinesisToS3Ptr(v *DefenderForContainersAwsOfferingResponseKinesisToS3Args) DefenderForContainersAwsOfferingResponseKinesisToS3PtrInput {
	return (*defenderForContainersAwsOfferingResponseKinesisToS3PtrType)(v)
}

func (*defenderForContainersAwsOfferingResponseKinesisToS3PtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForContainersAwsOfferingResponseKinesisToS3)(nil)).Elem()
}

func (i *defenderForContainersAwsOfferingResponseKinesisToS3PtrType) ToDefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput() DefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput {
	return i.ToDefenderForContainersAwsOfferingResponseKinesisToS3PtrOutputWithContext(context.Background())
}

func (i *defenderForContainersAwsOfferingResponseKinesisToS3PtrType) ToDefenderForContainersAwsOfferingResponseKinesisToS3PtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput)
}

type DefenderForContainersAwsOfferingResponseKinesisToS3Output struct{ *pulumi.OutputState }

func (DefenderForContainersAwsOfferingResponseKinesisToS3Output) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForContainersAwsOfferingResponseKinesisToS3)(nil)).Elem()
}

func (o DefenderForContainersAwsOfferingResponseKinesisToS3Output) ToDefenderForContainersAwsOfferingResponseKinesisToS3Output() DefenderForContainersAwsOfferingResponseKinesisToS3Output {
	return o
}

func (o DefenderForContainersAwsOfferingResponseKinesisToS3Output) ToDefenderForContainersAwsOfferingResponseKinesisToS3OutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseKinesisToS3Output {
	return o
}

func (o DefenderForContainersAwsOfferingResponseKinesisToS3Output) ToDefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput() DefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput {
	return o.ToDefenderForContainersAwsOfferingResponseKinesisToS3PtrOutputWithContext(context.Background())
}

func (o DefenderForContainersAwsOfferingResponseKinesisToS3Output) ToDefenderForContainersAwsOfferingResponseKinesisToS3PtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefenderForContainersAwsOfferingResponseKinesisToS3) *DefenderForContainersAwsOfferingResponseKinesisToS3 {
		return &v
	}).(DefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput)
}

func (o DefenderForContainersAwsOfferingResponseKinesisToS3Output) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DefenderForContainersAwsOfferingResponseKinesisToS3) *string { return v.CloudRoleArn }).(pulumi.StringPtrOutput)
}

type DefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput struct{ *pulumi.OutputState }

func (DefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForContainersAwsOfferingResponseKinesisToS3)(nil)).Elem()
}

func (o DefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput) ToDefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput() DefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput {
	return o
}

func (o DefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput) ToDefenderForContainersAwsOfferingResponseKinesisToS3PtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput {
	return o
}

func (o DefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput) Elem() DefenderForContainersAwsOfferingResponseKinesisToS3Output {
	return o.ApplyT(func(v *DefenderForContainersAwsOfferingResponseKinesisToS3) DefenderForContainersAwsOfferingResponseKinesisToS3 {
		if v != nil {
			return *v
		}
		var ret DefenderForContainersAwsOfferingResponseKinesisToS3
		return ret
	}).(DefenderForContainersAwsOfferingResponseKinesisToS3Output)
}

func (o DefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefenderForContainersAwsOfferingResponseKinesisToS3) *string {
		if v == nil {
			return nil
		}
		return v.CloudRoleArn
	}).(pulumi.StringPtrOutput)
}

type DefenderForContainersAwsOfferingResponseKubernetesScubaReader struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}





type DefenderForContainersAwsOfferingResponseKubernetesScubaReaderInput interface {
	pulumi.Input

	ToDefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutput() DefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutput
	ToDefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutputWithContext(context.Context) DefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutput
}

type DefenderForContainersAwsOfferingResponseKubernetesScubaReaderArgs struct {
	CloudRoleArn pulumi.StringPtrInput `pulumi:"cloudRoleArn"`
}

func (DefenderForContainersAwsOfferingResponseKubernetesScubaReaderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForContainersAwsOfferingResponseKubernetesScubaReader)(nil)).Elem()
}

func (i DefenderForContainersAwsOfferingResponseKubernetesScubaReaderArgs) ToDefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutput() DefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutput {
	return i.ToDefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutputWithContext(context.Background())
}

func (i DefenderForContainersAwsOfferingResponseKubernetesScubaReaderArgs) ToDefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutput)
}

func (i DefenderForContainersAwsOfferingResponseKubernetesScubaReaderArgs) ToDefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput() DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput {
	return i.ToDefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutputWithContext(context.Background())
}

func (i DefenderForContainersAwsOfferingResponseKubernetesScubaReaderArgs) ToDefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutput).ToDefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutputWithContext(ctx)
}









type DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrInput interface {
	pulumi.Input

	ToDefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput() DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput
	ToDefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutputWithContext(context.Context) DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput
}

type defenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrType DefenderForContainersAwsOfferingResponseKubernetesScubaReaderArgs

func DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtr(v *DefenderForContainersAwsOfferingResponseKubernetesScubaReaderArgs) DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrInput {
	return (*defenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrType)(v)
}

func (*defenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForContainersAwsOfferingResponseKubernetesScubaReader)(nil)).Elem()
}

func (i *defenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrType) ToDefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput() DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput {
	return i.ToDefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutputWithContext(context.Background())
}

func (i *defenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrType) ToDefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput)
}

type DefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutput struct{ *pulumi.OutputState }

func (DefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForContainersAwsOfferingResponseKubernetesScubaReader)(nil)).Elem()
}

func (o DefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutput) ToDefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutput() DefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutput {
	return o
}

func (o DefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutput) ToDefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutput {
	return o
}

func (o DefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutput) ToDefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput() DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput {
	return o.ToDefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutputWithContext(context.Background())
}

func (o DefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutput) ToDefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefenderForContainersAwsOfferingResponseKubernetesScubaReader) *DefenderForContainersAwsOfferingResponseKubernetesScubaReader {
		return &v
	}).(DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput)
}

func (o DefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutput) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DefenderForContainersAwsOfferingResponseKubernetesScubaReader) *string { return v.CloudRoleArn }).(pulumi.StringPtrOutput)
}

type DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput struct{ *pulumi.OutputState }

func (DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForContainersAwsOfferingResponseKubernetesScubaReader)(nil)).Elem()
}

func (o DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput) ToDefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput() DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput {
	return o
}

func (o DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput) ToDefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput {
	return o
}

func (o DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput) Elem() DefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutput {
	return o.ApplyT(func(v *DefenderForContainersAwsOfferingResponseKubernetesScubaReader) DefenderForContainersAwsOfferingResponseKubernetesScubaReader {
		if v != nil {
			return *v
		}
		var ret DefenderForContainersAwsOfferingResponseKubernetesScubaReader
		return ret
	}).(DefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutput)
}

func (o DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefenderForContainersAwsOfferingResponseKubernetesScubaReader) *string {
		if v == nil {
			return nil
		}
		return v.CloudRoleArn
	}).(pulumi.StringPtrOutput)
}

type DefenderForContainersAwsOfferingResponseKubernetesService struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}





type DefenderForContainersAwsOfferingResponseKubernetesServiceInput interface {
	pulumi.Input

	ToDefenderForContainersAwsOfferingResponseKubernetesServiceOutput() DefenderForContainersAwsOfferingResponseKubernetesServiceOutput
	ToDefenderForContainersAwsOfferingResponseKubernetesServiceOutputWithContext(context.Context) DefenderForContainersAwsOfferingResponseKubernetesServiceOutput
}

type DefenderForContainersAwsOfferingResponseKubernetesServiceArgs struct {
	CloudRoleArn pulumi.StringPtrInput `pulumi:"cloudRoleArn"`
}

func (DefenderForContainersAwsOfferingResponseKubernetesServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForContainersAwsOfferingResponseKubernetesService)(nil)).Elem()
}

func (i DefenderForContainersAwsOfferingResponseKubernetesServiceArgs) ToDefenderForContainersAwsOfferingResponseKubernetesServiceOutput() DefenderForContainersAwsOfferingResponseKubernetesServiceOutput {
	return i.ToDefenderForContainersAwsOfferingResponseKubernetesServiceOutputWithContext(context.Background())
}

func (i DefenderForContainersAwsOfferingResponseKubernetesServiceArgs) ToDefenderForContainersAwsOfferingResponseKubernetesServiceOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseKubernetesServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingResponseKubernetesServiceOutput)
}

func (i DefenderForContainersAwsOfferingResponseKubernetesServiceArgs) ToDefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput() DefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput {
	return i.ToDefenderForContainersAwsOfferingResponseKubernetesServicePtrOutputWithContext(context.Background())
}

func (i DefenderForContainersAwsOfferingResponseKubernetesServiceArgs) ToDefenderForContainersAwsOfferingResponseKubernetesServicePtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingResponseKubernetesServiceOutput).ToDefenderForContainersAwsOfferingResponseKubernetesServicePtrOutputWithContext(ctx)
}









type DefenderForContainersAwsOfferingResponseKubernetesServicePtrInput interface {
	pulumi.Input

	ToDefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput() DefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput
	ToDefenderForContainersAwsOfferingResponseKubernetesServicePtrOutputWithContext(context.Context) DefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput
}

type defenderForContainersAwsOfferingResponseKubernetesServicePtrType DefenderForContainersAwsOfferingResponseKubernetesServiceArgs

func DefenderForContainersAwsOfferingResponseKubernetesServicePtr(v *DefenderForContainersAwsOfferingResponseKubernetesServiceArgs) DefenderForContainersAwsOfferingResponseKubernetesServicePtrInput {
	return (*defenderForContainersAwsOfferingResponseKubernetesServicePtrType)(v)
}

func (*defenderForContainersAwsOfferingResponseKubernetesServicePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForContainersAwsOfferingResponseKubernetesService)(nil)).Elem()
}

func (i *defenderForContainersAwsOfferingResponseKubernetesServicePtrType) ToDefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput() DefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput {
	return i.ToDefenderForContainersAwsOfferingResponseKubernetesServicePtrOutputWithContext(context.Background())
}

func (i *defenderForContainersAwsOfferingResponseKubernetesServicePtrType) ToDefenderForContainersAwsOfferingResponseKubernetesServicePtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput)
}

type DefenderForContainersAwsOfferingResponseKubernetesServiceOutput struct{ *pulumi.OutputState }

func (DefenderForContainersAwsOfferingResponseKubernetesServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForContainersAwsOfferingResponseKubernetesService)(nil)).Elem()
}

func (o DefenderForContainersAwsOfferingResponseKubernetesServiceOutput) ToDefenderForContainersAwsOfferingResponseKubernetesServiceOutput() DefenderForContainersAwsOfferingResponseKubernetesServiceOutput {
	return o
}

func (o DefenderForContainersAwsOfferingResponseKubernetesServiceOutput) ToDefenderForContainersAwsOfferingResponseKubernetesServiceOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseKubernetesServiceOutput {
	return o
}

func (o DefenderForContainersAwsOfferingResponseKubernetesServiceOutput) ToDefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput() DefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput {
	return o.ToDefenderForContainersAwsOfferingResponseKubernetesServicePtrOutputWithContext(context.Background())
}

func (o DefenderForContainersAwsOfferingResponseKubernetesServiceOutput) ToDefenderForContainersAwsOfferingResponseKubernetesServicePtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefenderForContainersAwsOfferingResponseKubernetesService) *DefenderForContainersAwsOfferingResponseKubernetesService {
		return &v
	}).(DefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput)
}

func (o DefenderForContainersAwsOfferingResponseKubernetesServiceOutput) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DefenderForContainersAwsOfferingResponseKubernetesService) *string { return v.CloudRoleArn }).(pulumi.StringPtrOutput)
}

type DefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput struct{ *pulumi.OutputState }

func (DefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForContainersAwsOfferingResponseKubernetesService)(nil)).Elem()
}

func (o DefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput) ToDefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput() DefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput {
	return o
}

func (o DefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput) ToDefenderForContainersAwsOfferingResponseKubernetesServicePtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput {
	return o
}

func (o DefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput) Elem() DefenderForContainersAwsOfferingResponseKubernetesServiceOutput {
	return o.ApplyT(func(v *DefenderForContainersAwsOfferingResponseKubernetesService) DefenderForContainersAwsOfferingResponseKubernetesService {
		if v != nil {
			return *v
		}
		var ret DefenderForContainersAwsOfferingResponseKubernetesService
		return ret
	}).(DefenderForContainersAwsOfferingResponseKubernetesServiceOutput)
}

func (o DefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefenderForContainersAwsOfferingResponseKubernetesService) *string {
		if v == nil {
			return nil
		}
		return v.CloudRoleArn
	}).(pulumi.StringPtrOutput)
}

type DefenderForServersAwsOffering struct {
	ArcAutoProvisioning *DefenderForServersAwsOfferingArcAutoProvisioning `pulumi:"arcAutoProvisioning"`
	DefenderForServers  *DefenderForServersAwsOfferingDefenderForServers  `pulumi:"defenderForServers"`
	OfferingType        string                                            `pulumi:"offeringType"`
}





type DefenderForServersAwsOfferingInput interface {
	pulumi.Input

	ToDefenderForServersAwsOfferingOutput() DefenderForServersAwsOfferingOutput
	ToDefenderForServersAwsOfferingOutputWithContext(context.Context) DefenderForServersAwsOfferingOutput
}

type DefenderForServersAwsOfferingArgs struct {
	ArcAutoProvisioning DefenderForServersAwsOfferingArcAutoProvisioningPtrInput `pulumi:"arcAutoProvisioning"`
	DefenderForServers  DefenderForServersAwsOfferingDefenderForServersPtrInput  `pulumi:"defenderForServers"`
	OfferingType        pulumi.StringInput                                       `pulumi:"offeringType"`
}

func (DefenderForServersAwsOfferingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForServersAwsOffering)(nil)).Elem()
}

func (i DefenderForServersAwsOfferingArgs) ToDefenderForServersAwsOfferingOutput() DefenderForServersAwsOfferingOutput {
	return i.ToDefenderForServersAwsOfferingOutputWithContext(context.Background())
}

func (i DefenderForServersAwsOfferingArgs) ToDefenderForServersAwsOfferingOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForServersAwsOfferingOutput)
}

type DefenderForServersAwsOfferingOutput struct{ *pulumi.OutputState }

func (DefenderForServersAwsOfferingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForServersAwsOffering)(nil)).Elem()
}

func (o DefenderForServersAwsOfferingOutput) ToDefenderForServersAwsOfferingOutput() DefenderForServersAwsOfferingOutput {
	return o
}

func (o DefenderForServersAwsOfferingOutput) ToDefenderForServersAwsOfferingOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingOutput {
	return o
}

func (o DefenderForServersAwsOfferingOutput) ArcAutoProvisioning() DefenderForServersAwsOfferingArcAutoProvisioningPtrOutput {
	return o.ApplyT(func(v DefenderForServersAwsOffering) *DefenderForServersAwsOfferingArcAutoProvisioning {
		return v.ArcAutoProvisioning
	}).(DefenderForServersAwsOfferingArcAutoProvisioningPtrOutput)
}

func (o DefenderForServersAwsOfferingOutput) DefenderForServers() DefenderForServersAwsOfferingDefenderForServersPtrOutput {
	return o.ApplyT(func(v DefenderForServersAwsOffering) *DefenderForServersAwsOfferingDefenderForServers {
		return v.DefenderForServers
	}).(DefenderForServersAwsOfferingDefenderForServersPtrOutput)
}

func (o DefenderForServersAwsOfferingOutput) OfferingType() pulumi.StringOutput {
	return o.ApplyT(func(v DefenderForServersAwsOffering) string { return v.OfferingType }).(pulumi.StringOutput)
}

type DefenderForServersAwsOfferingArcAutoProvisioning struct {
	Enabled                        *bool                                                        `pulumi:"enabled"`
	ServicePrincipalSecretMetadata *DefenderForServersAwsOfferingServicePrincipalSecretMetadata `pulumi:"servicePrincipalSecretMetadata"`
}





type DefenderForServersAwsOfferingArcAutoProvisioningInput interface {
	pulumi.Input

	ToDefenderForServersAwsOfferingArcAutoProvisioningOutput() DefenderForServersAwsOfferingArcAutoProvisioningOutput
	ToDefenderForServersAwsOfferingArcAutoProvisioningOutputWithContext(context.Context) DefenderForServersAwsOfferingArcAutoProvisioningOutput
}

type DefenderForServersAwsOfferingArcAutoProvisioningArgs struct {
	Enabled                        pulumi.BoolPtrInput                                                 `pulumi:"enabled"`
	ServicePrincipalSecretMetadata DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrInput `pulumi:"servicePrincipalSecretMetadata"`
}

func (DefenderForServersAwsOfferingArcAutoProvisioningArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForServersAwsOfferingArcAutoProvisioning)(nil)).Elem()
}

func (i DefenderForServersAwsOfferingArcAutoProvisioningArgs) ToDefenderForServersAwsOfferingArcAutoProvisioningOutput() DefenderForServersAwsOfferingArcAutoProvisioningOutput {
	return i.ToDefenderForServersAwsOfferingArcAutoProvisioningOutputWithContext(context.Background())
}

func (i DefenderForServersAwsOfferingArcAutoProvisioningArgs) ToDefenderForServersAwsOfferingArcAutoProvisioningOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingArcAutoProvisioningOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForServersAwsOfferingArcAutoProvisioningOutput)
}

func (i DefenderForServersAwsOfferingArcAutoProvisioningArgs) ToDefenderForServersAwsOfferingArcAutoProvisioningPtrOutput() DefenderForServersAwsOfferingArcAutoProvisioningPtrOutput {
	return i.ToDefenderForServersAwsOfferingArcAutoProvisioningPtrOutputWithContext(context.Background())
}

func (i DefenderForServersAwsOfferingArcAutoProvisioningArgs) ToDefenderForServersAwsOfferingArcAutoProvisioningPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingArcAutoProvisioningPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForServersAwsOfferingArcAutoProvisioningOutput).ToDefenderForServersAwsOfferingArcAutoProvisioningPtrOutputWithContext(ctx)
}









type DefenderForServersAwsOfferingArcAutoProvisioningPtrInput interface {
	pulumi.Input

	ToDefenderForServersAwsOfferingArcAutoProvisioningPtrOutput() DefenderForServersAwsOfferingArcAutoProvisioningPtrOutput
	ToDefenderForServersAwsOfferingArcAutoProvisioningPtrOutputWithContext(context.Context) DefenderForServersAwsOfferingArcAutoProvisioningPtrOutput
}

type defenderForServersAwsOfferingArcAutoProvisioningPtrType DefenderForServersAwsOfferingArcAutoProvisioningArgs

func DefenderForServersAwsOfferingArcAutoProvisioningPtr(v *DefenderForServersAwsOfferingArcAutoProvisioningArgs) DefenderForServersAwsOfferingArcAutoProvisioningPtrInput {
	return (*defenderForServersAwsOfferingArcAutoProvisioningPtrType)(v)
}

func (*defenderForServersAwsOfferingArcAutoProvisioningPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForServersAwsOfferingArcAutoProvisioning)(nil)).Elem()
}

func (i *defenderForServersAwsOfferingArcAutoProvisioningPtrType) ToDefenderForServersAwsOfferingArcAutoProvisioningPtrOutput() DefenderForServersAwsOfferingArcAutoProvisioningPtrOutput {
	return i.ToDefenderForServersAwsOfferingArcAutoProvisioningPtrOutputWithContext(context.Background())
}

func (i *defenderForServersAwsOfferingArcAutoProvisioningPtrType) ToDefenderForServersAwsOfferingArcAutoProvisioningPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingArcAutoProvisioningPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForServersAwsOfferingArcAutoProvisioningPtrOutput)
}

type DefenderForServersAwsOfferingArcAutoProvisioningOutput struct{ *pulumi.OutputState }

func (DefenderForServersAwsOfferingArcAutoProvisioningOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForServersAwsOfferingArcAutoProvisioning)(nil)).Elem()
}

func (o DefenderForServersAwsOfferingArcAutoProvisioningOutput) ToDefenderForServersAwsOfferingArcAutoProvisioningOutput() DefenderForServersAwsOfferingArcAutoProvisioningOutput {
	return o
}

func (o DefenderForServersAwsOfferingArcAutoProvisioningOutput) ToDefenderForServersAwsOfferingArcAutoProvisioningOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingArcAutoProvisioningOutput {
	return o
}

func (o DefenderForServersAwsOfferingArcAutoProvisioningOutput) ToDefenderForServersAwsOfferingArcAutoProvisioningPtrOutput() DefenderForServersAwsOfferingArcAutoProvisioningPtrOutput {
	return o.ToDefenderForServersAwsOfferingArcAutoProvisioningPtrOutputWithContext(context.Background())
}

func (o DefenderForServersAwsOfferingArcAutoProvisioningOutput) ToDefenderForServersAwsOfferingArcAutoProvisioningPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingArcAutoProvisioningPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefenderForServersAwsOfferingArcAutoProvisioning) *DefenderForServersAwsOfferingArcAutoProvisioning {
		return &v
	}).(DefenderForServersAwsOfferingArcAutoProvisioningPtrOutput)
}

func (o DefenderForServersAwsOfferingArcAutoProvisioningOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DefenderForServersAwsOfferingArcAutoProvisioning) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o DefenderForServersAwsOfferingArcAutoProvisioningOutput) ServicePrincipalSecretMetadata() DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput {
	return o.ApplyT(func(v DefenderForServersAwsOfferingArcAutoProvisioning) *DefenderForServersAwsOfferingServicePrincipalSecretMetadata {
		return v.ServicePrincipalSecretMetadata
	}).(DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput)
}

type DefenderForServersAwsOfferingArcAutoProvisioningPtrOutput struct{ *pulumi.OutputState }

func (DefenderForServersAwsOfferingArcAutoProvisioningPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForServersAwsOfferingArcAutoProvisioning)(nil)).Elem()
}

func (o DefenderForServersAwsOfferingArcAutoProvisioningPtrOutput) ToDefenderForServersAwsOfferingArcAutoProvisioningPtrOutput() DefenderForServersAwsOfferingArcAutoProvisioningPtrOutput {
	return o
}

func (o DefenderForServersAwsOfferingArcAutoProvisioningPtrOutput) ToDefenderForServersAwsOfferingArcAutoProvisioningPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingArcAutoProvisioningPtrOutput {
	return o
}

func (o DefenderForServersAwsOfferingArcAutoProvisioningPtrOutput) Elem() DefenderForServersAwsOfferingArcAutoProvisioningOutput {
	return o.ApplyT(func(v *DefenderForServersAwsOfferingArcAutoProvisioning) DefenderForServersAwsOfferingArcAutoProvisioning {
		if v != nil {
			return *v
		}
		var ret DefenderForServersAwsOfferingArcAutoProvisioning
		return ret
	}).(DefenderForServersAwsOfferingArcAutoProvisioningOutput)
}

func (o DefenderForServersAwsOfferingArcAutoProvisioningPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DefenderForServersAwsOfferingArcAutoProvisioning) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o DefenderForServersAwsOfferingArcAutoProvisioningPtrOutput) ServicePrincipalSecretMetadata() DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput {
	return o.ApplyT(func(v *DefenderForServersAwsOfferingArcAutoProvisioning) *DefenderForServersAwsOfferingServicePrincipalSecretMetadata {
		if v == nil {
			return nil
		}
		return v.ServicePrincipalSecretMetadata
	}).(DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput)
}

type DefenderForServersAwsOfferingDefenderForServers struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}





type DefenderForServersAwsOfferingDefenderForServersInput interface {
	pulumi.Input

	ToDefenderForServersAwsOfferingDefenderForServersOutput() DefenderForServersAwsOfferingDefenderForServersOutput
	ToDefenderForServersAwsOfferingDefenderForServersOutputWithContext(context.Context) DefenderForServersAwsOfferingDefenderForServersOutput
}

type DefenderForServersAwsOfferingDefenderForServersArgs struct {
	CloudRoleArn pulumi.StringPtrInput `pulumi:"cloudRoleArn"`
}

func (DefenderForServersAwsOfferingDefenderForServersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForServersAwsOfferingDefenderForServers)(nil)).Elem()
}

func (i DefenderForServersAwsOfferingDefenderForServersArgs) ToDefenderForServersAwsOfferingDefenderForServersOutput() DefenderForServersAwsOfferingDefenderForServersOutput {
	return i.ToDefenderForServersAwsOfferingDefenderForServersOutputWithContext(context.Background())
}

func (i DefenderForServersAwsOfferingDefenderForServersArgs) ToDefenderForServersAwsOfferingDefenderForServersOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingDefenderForServersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForServersAwsOfferingDefenderForServersOutput)
}

func (i DefenderForServersAwsOfferingDefenderForServersArgs) ToDefenderForServersAwsOfferingDefenderForServersPtrOutput() DefenderForServersAwsOfferingDefenderForServersPtrOutput {
	return i.ToDefenderForServersAwsOfferingDefenderForServersPtrOutputWithContext(context.Background())
}

func (i DefenderForServersAwsOfferingDefenderForServersArgs) ToDefenderForServersAwsOfferingDefenderForServersPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingDefenderForServersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForServersAwsOfferingDefenderForServersOutput).ToDefenderForServersAwsOfferingDefenderForServersPtrOutputWithContext(ctx)
}









type DefenderForServersAwsOfferingDefenderForServersPtrInput interface {
	pulumi.Input

	ToDefenderForServersAwsOfferingDefenderForServersPtrOutput() DefenderForServersAwsOfferingDefenderForServersPtrOutput
	ToDefenderForServersAwsOfferingDefenderForServersPtrOutputWithContext(context.Context) DefenderForServersAwsOfferingDefenderForServersPtrOutput
}

type defenderForServersAwsOfferingDefenderForServersPtrType DefenderForServersAwsOfferingDefenderForServersArgs

func DefenderForServersAwsOfferingDefenderForServersPtr(v *DefenderForServersAwsOfferingDefenderForServersArgs) DefenderForServersAwsOfferingDefenderForServersPtrInput {
	return (*defenderForServersAwsOfferingDefenderForServersPtrType)(v)
}

func (*defenderForServersAwsOfferingDefenderForServersPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForServersAwsOfferingDefenderForServers)(nil)).Elem()
}

func (i *defenderForServersAwsOfferingDefenderForServersPtrType) ToDefenderForServersAwsOfferingDefenderForServersPtrOutput() DefenderForServersAwsOfferingDefenderForServersPtrOutput {
	return i.ToDefenderForServersAwsOfferingDefenderForServersPtrOutputWithContext(context.Background())
}

func (i *defenderForServersAwsOfferingDefenderForServersPtrType) ToDefenderForServersAwsOfferingDefenderForServersPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingDefenderForServersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForServersAwsOfferingDefenderForServersPtrOutput)
}

type DefenderForServersAwsOfferingDefenderForServersOutput struct{ *pulumi.OutputState }

func (DefenderForServersAwsOfferingDefenderForServersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForServersAwsOfferingDefenderForServers)(nil)).Elem()
}

func (o DefenderForServersAwsOfferingDefenderForServersOutput) ToDefenderForServersAwsOfferingDefenderForServersOutput() DefenderForServersAwsOfferingDefenderForServersOutput {
	return o
}

func (o DefenderForServersAwsOfferingDefenderForServersOutput) ToDefenderForServersAwsOfferingDefenderForServersOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingDefenderForServersOutput {
	return o
}

func (o DefenderForServersAwsOfferingDefenderForServersOutput) ToDefenderForServersAwsOfferingDefenderForServersPtrOutput() DefenderForServersAwsOfferingDefenderForServersPtrOutput {
	return o.ToDefenderForServersAwsOfferingDefenderForServersPtrOutputWithContext(context.Background())
}

func (o DefenderForServersAwsOfferingDefenderForServersOutput) ToDefenderForServersAwsOfferingDefenderForServersPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingDefenderForServersPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefenderForServersAwsOfferingDefenderForServers) *DefenderForServersAwsOfferingDefenderForServers {
		return &v
	}).(DefenderForServersAwsOfferingDefenderForServersPtrOutput)
}

func (o DefenderForServersAwsOfferingDefenderForServersOutput) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DefenderForServersAwsOfferingDefenderForServers) *string { return v.CloudRoleArn }).(pulumi.StringPtrOutput)
}

type DefenderForServersAwsOfferingDefenderForServersPtrOutput struct{ *pulumi.OutputState }

func (DefenderForServersAwsOfferingDefenderForServersPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForServersAwsOfferingDefenderForServers)(nil)).Elem()
}

func (o DefenderForServersAwsOfferingDefenderForServersPtrOutput) ToDefenderForServersAwsOfferingDefenderForServersPtrOutput() DefenderForServersAwsOfferingDefenderForServersPtrOutput {
	return o
}

func (o DefenderForServersAwsOfferingDefenderForServersPtrOutput) ToDefenderForServersAwsOfferingDefenderForServersPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingDefenderForServersPtrOutput {
	return o
}

func (o DefenderForServersAwsOfferingDefenderForServersPtrOutput) Elem() DefenderForServersAwsOfferingDefenderForServersOutput {
	return o.ApplyT(func(v *DefenderForServersAwsOfferingDefenderForServers) DefenderForServersAwsOfferingDefenderForServers {
		if v != nil {
			return *v
		}
		var ret DefenderForServersAwsOfferingDefenderForServers
		return ret
	}).(DefenderForServersAwsOfferingDefenderForServersOutput)
}

func (o DefenderForServersAwsOfferingDefenderForServersPtrOutput) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefenderForServersAwsOfferingDefenderForServers) *string {
		if v == nil {
			return nil
		}
		return v.CloudRoleArn
	}).(pulumi.StringPtrOutput)
}

type DefenderForServersAwsOfferingResponse struct {
	ArcAutoProvisioning *DefenderForServersAwsOfferingResponseArcAutoProvisioning `pulumi:"arcAutoProvisioning"`
	DefenderForServers  *DefenderForServersAwsOfferingResponseDefenderForServers  `pulumi:"defenderForServers"`
	Description         string                                                    `pulumi:"description"`
	OfferingType        string                                                    `pulumi:"offeringType"`
}





type DefenderForServersAwsOfferingResponseInput interface {
	pulumi.Input

	ToDefenderForServersAwsOfferingResponseOutput() DefenderForServersAwsOfferingResponseOutput
	ToDefenderForServersAwsOfferingResponseOutputWithContext(context.Context) DefenderForServersAwsOfferingResponseOutput
}

type DefenderForServersAwsOfferingResponseArgs struct {
	ArcAutoProvisioning DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrInput `pulumi:"arcAutoProvisioning"`
	DefenderForServers  DefenderForServersAwsOfferingResponseDefenderForServersPtrInput  `pulumi:"defenderForServers"`
	Description         pulumi.StringInput                                               `pulumi:"description"`
	OfferingType        pulumi.StringInput                                               `pulumi:"offeringType"`
}

func (DefenderForServersAwsOfferingResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForServersAwsOfferingResponse)(nil)).Elem()
}

func (i DefenderForServersAwsOfferingResponseArgs) ToDefenderForServersAwsOfferingResponseOutput() DefenderForServersAwsOfferingResponseOutput {
	return i.ToDefenderForServersAwsOfferingResponseOutputWithContext(context.Background())
}

func (i DefenderForServersAwsOfferingResponseArgs) ToDefenderForServersAwsOfferingResponseOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForServersAwsOfferingResponseOutput)
}

type DefenderForServersAwsOfferingResponseOutput struct{ *pulumi.OutputState }

func (DefenderForServersAwsOfferingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForServersAwsOfferingResponse)(nil)).Elem()
}

func (o DefenderForServersAwsOfferingResponseOutput) ToDefenderForServersAwsOfferingResponseOutput() DefenderForServersAwsOfferingResponseOutput {
	return o
}

func (o DefenderForServersAwsOfferingResponseOutput) ToDefenderForServersAwsOfferingResponseOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingResponseOutput {
	return o
}

func (o DefenderForServersAwsOfferingResponseOutput) ArcAutoProvisioning() DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput {
	return o.ApplyT(func(v DefenderForServersAwsOfferingResponse) *DefenderForServersAwsOfferingResponseArcAutoProvisioning {
		return v.ArcAutoProvisioning
	}).(DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput)
}

func (o DefenderForServersAwsOfferingResponseOutput) DefenderForServers() DefenderForServersAwsOfferingResponseDefenderForServersPtrOutput {
	return o.ApplyT(func(v DefenderForServersAwsOfferingResponse) *DefenderForServersAwsOfferingResponseDefenderForServers {
		return v.DefenderForServers
	}).(DefenderForServersAwsOfferingResponseDefenderForServersPtrOutput)
}

func (o DefenderForServersAwsOfferingResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v DefenderForServersAwsOfferingResponse) string { return v.Description }).(pulumi.StringOutput)
}

func (o DefenderForServersAwsOfferingResponseOutput) OfferingType() pulumi.StringOutput {
	return o.ApplyT(func(v DefenderForServersAwsOfferingResponse) string { return v.OfferingType }).(pulumi.StringOutput)
}

type DefenderForServersAwsOfferingResponseArcAutoProvisioning struct {
	Enabled                        *bool                                                                `pulumi:"enabled"`
	ServicePrincipalSecretMetadata *DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadata `pulumi:"servicePrincipalSecretMetadata"`
}





type DefenderForServersAwsOfferingResponseArcAutoProvisioningInput interface {
	pulumi.Input

	ToDefenderForServersAwsOfferingResponseArcAutoProvisioningOutput() DefenderForServersAwsOfferingResponseArcAutoProvisioningOutput
	ToDefenderForServersAwsOfferingResponseArcAutoProvisioningOutputWithContext(context.Context) DefenderForServersAwsOfferingResponseArcAutoProvisioningOutput
}

type DefenderForServersAwsOfferingResponseArcAutoProvisioningArgs struct {
	Enabled                        pulumi.BoolPtrInput                                                         `pulumi:"enabled"`
	ServicePrincipalSecretMetadata DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrInput `pulumi:"servicePrincipalSecretMetadata"`
}

func (DefenderForServersAwsOfferingResponseArcAutoProvisioningArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForServersAwsOfferingResponseArcAutoProvisioning)(nil)).Elem()
}

func (i DefenderForServersAwsOfferingResponseArcAutoProvisioningArgs) ToDefenderForServersAwsOfferingResponseArcAutoProvisioningOutput() DefenderForServersAwsOfferingResponseArcAutoProvisioningOutput {
	return i.ToDefenderForServersAwsOfferingResponseArcAutoProvisioningOutputWithContext(context.Background())
}

func (i DefenderForServersAwsOfferingResponseArcAutoProvisioningArgs) ToDefenderForServersAwsOfferingResponseArcAutoProvisioningOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingResponseArcAutoProvisioningOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForServersAwsOfferingResponseArcAutoProvisioningOutput)
}

func (i DefenderForServersAwsOfferingResponseArcAutoProvisioningArgs) ToDefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput() DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput {
	return i.ToDefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutputWithContext(context.Background())
}

func (i DefenderForServersAwsOfferingResponseArcAutoProvisioningArgs) ToDefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForServersAwsOfferingResponseArcAutoProvisioningOutput).ToDefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutputWithContext(ctx)
}









type DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrInput interface {
	pulumi.Input

	ToDefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput() DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput
	ToDefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutputWithContext(context.Context) DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput
}

type defenderForServersAwsOfferingResponseArcAutoProvisioningPtrType DefenderForServersAwsOfferingResponseArcAutoProvisioningArgs

func DefenderForServersAwsOfferingResponseArcAutoProvisioningPtr(v *DefenderForServersAwsOfferingResponseArcAutoProvisioningArgs) DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrInput {
	return (*defenderForServersAwsOfferingResponseArcAutoProvisioningPtrType)(v)
}

func (*defenderForServersAwsOfferingResponseArcAutoProvisioningPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForServersAwsOfferingResponseArcAutoProvisioning)(nil)).Elem()
}

func (i *defenderForServersAwsOfferingResponseArcAutoProvisioningPtrType) ToDefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput() DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput {
	return i.ToDefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutputWithContext(context.Background())
}

func (i *defenderForServersAwsOfferingResponseArcAutoProvisioningPtrType) ToDefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput)
}

type DefenderForServersAwsOfferingResponseArcAutoProvisioningOutput struct{ *pulumi.OutputState }

func (DefenderForServersAwsOfferingResponseArcAutoProvisioningOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForServersAwsOfferingResponseArcAutoProvisioning)(nil)).Elem()
}

func (o DefenderForServersAwsOfferingResponseArcAutoProvisioningOutput) ToDefenderForServersAwsOfferingResponseArcAutoProvisioningOutput() DefenderForServersAwsOfferingResponseArcAutoProvisioningOutput {
	return o
}

func (o DefenderForServersAwsOfferingResponseArcAutoProvisioningOutput) ToDefenderForServersAwsOfferingResponseArcAutoProvisioningOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingResponseArcAutoProvisioningOutput {
	return o
}

func (o DefenderForServersAwsOfferingResponseArcAutoProvisioningOutput) ToDefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput() DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput {
	return o.ToDefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutputWithContext(context.Background())
}

func (o DefenderForServersAwsOfferingResponseArcAutoProvisioningOutput) ToDefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefenderForServersAwsOfferingResponseArcAutoProvisioning) *DefenderForServersAwsOfferingResponseArcAutoProvisioning {
		return &v
	}).(DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput)
}

func (o DefenderForServersAwsOfferingResponseArcAutoProvisioningOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DefenderForServersAwsOfferingResponseArcAutoProvisioning) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o DefenderForServersAwsOfferingResponseArcAutoProvisioningOutput) ServicePrincipalSecretMetadata() DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput {
	return o.ApplyT(func(v DefenderForServersAwsOfferingResponseArcAutoProvisioning) *DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadata {
		return v.ServicePrincipalSecretMetadata
	}).(DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput)
}

type DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput struct{ *pulumi.OutputState }

func (DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForServersAwsOfferingResponseArcAutoProvisioning)(nil)).Elem()
}

func (o DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput) ToDefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput() DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput {
	return o
}

func (o DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput) ToDefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput {
	return o
}

func (o DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput) Elem() DefenderForServersAwsOfferingResponseArcAutoProvisioningOutput {
	return o.ApplyT(func(v *DefenderForServersAwsOfferingResponseArcAutoProvisioning) DefenderForServersAwsOfferingResponseArcAutoProvisioning {
		if v != nil {
			return *v
		}
		var ret DefenderForServersAwsOfferingResponseArcAutoProvisioning
		return ret
	}).(DefenderForServersAwsOfferingResponseArcAutoProvisioningOutput)
}

func (o DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DefenderForServersAwsOfferingResponseArcAutoProvisioning) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput) ServicePrincipalSecretMetadata() DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput {
	return o.ApplyT(func(v *DefenderForServersAwsOfferingResponseArcAutoProvisioning) *DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadata {
		if v == nil {
			return nil
		}
		return v.ServicePrincipalSecretMetadata
	}).(DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput)
}

type DefenderForServersAwsOfferingResponseDefenderForServers struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}





type DefenderForServersAwsOfferingResponseDefenderForServersInput interface {
	pulumi.Input

	ToDefenderForServersAwsOfferingResponseDefenderForServersOutput() DefenderForServersAwsOfferingResponseDefenderForServersOutput
	ToDefenderForServersAwsOfferingResponseDefenderForServersOutputWithContext(context.Context) DefenderForServersAwsOfferingResponseDefenderForServersOutput
}

type DefenderForServersAwsOfferingResponseDefenderForServersArgs struct {
	CloudRoleArn pulumi.StringPtrInput `pulumi:"cloudRoleArn"`
}

func (DefenderForServersAwsOfferingResponseDefenderForServersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForServersAwsOfferingResponseDefenderForServers)(nil)).Elem()
}

func (i DefenderForServersAwsOfferingResponseDefenderForServersArgs) ToDefenderForServersAwsOfferingResponseDefenderForServersOutput() DefenderForServersAwsOfferingResponseDefenderForServersOutput {
	return i.ToDefenderForServersAwsOfferingResponseDefenderForServersOutputWithContext(context.Background())
}

func (i DefenderForServersAwsOfferingResponseDefenderForServersArgs) ToDefenderForServersAwsOfferingResponseDefenderForServersOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingResponseDefenderForServersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForServersAwsOfferingResponseDefenderForServersOutput)
}

func (i DefenderForServersAwsOfferingResponseDefenderForServersArgs) ToDefenderForServersAwsOfferingResponseDefenderForServersPtrOutput() DefenderForServersAwsOfferingResponseDefenderForServersPtrOutput {
	return i.ToDefenderForServersAwsOfferingResponseDefenderForServersPtrOutputWithContext(context.Background())
}

func (i DefenderForServersAwsOfferingResponseDefenderForServersArgs) ToDefenderForServersAwsOfferingResponseDefenderForServersPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingResponseDefenderForServersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForServersAwsOfferingResponseDefenderForServersOutput).ToDefenderForServersAwsOfferingResponseDefenderForServersPtrOutputWithContext(ctx)
}









type DefenderForServersAwsOfferingResponseDefenderForServersPtrInput interface {
	pulumi.Input

	ToDefenderForServersAwsOfferingResponseDefenderForServersPtrOutput() DefenderForServersAwsOfferingResponseDefenderForServersPtrOutput
	ToDefenderForServersAwsOfferingResponseDefenderForServersPtrOutputWithContext(context.Context) DefenderForServersAwsOfferingResponseDefenderForServersPtrOutput
}

type defenderForServersAwsOfferingResponseDefenderForServersPtrType DefenderForServersAwsOfferingResponseDefenderForServersArgs

func DefenderForServersAwsOfferingResponseDefenderForServersPtr(v *DefenderForServersAwsOfferingResponseDefenderForServersArgs) DefenderForServersAwsOfferingResponseDefenderForServersPtrInput {
	return (*defenderForServersAwsOfferingResponseDefenderForServersPtrType)(v)
}

func (*defenderForServersAwsOfferingResponseDefenderForServersPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForServersAwsOfferingResponseDefenderForServers)(nil)).Elem()
}

func (i *defenderForServersAwsOfferingResponseDefenderForServersPtrType) ToDefenderForServersAwsOfferingResponseDefenderForServersPtrOutput() DefenderForServersAwsOfferingResponseDefenderForServersPtrOutput {
	return i.ToDefenderForServersAwsOfferingResponseDefenderForServersPtrOutputWithContext(context.Background())
}

func (i *defenderForServersAwsOfferingResponseDefenderForServersPtrType) ToDefenderForServersAwsOfferingResponseDefenderForServersPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingResponseDefenderForServersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForServersAwsOfferingResponseDefenderForServersPtrOutput)
}

type DefenderForServersAwsOfferingResponseDefenderForServersOutput struct{ *pulumi.OutputState }

func (DefenderForServersAwsOfferingResponseDefenderForServersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForServersAwsOfferingResponseDefenderForServers)(nil)).Elem()
}

func (o DefenderForServersAwsOfferingResponseDefenderForServersOutput) ToDefenderForServersAwsOfferingResponseDefenderForServersOutput() DefenderForServersAwsOfferingResponseDefenderForServersOutput {
	return o
}

func (o DefenderForServersAwsOfferingResponseDefenderForServersOutput) ToDefenderForServersAwsOfferingResponseDefenderForServersOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingResponseDefenderForServersOutput {
	return o
}

func (o DefenderForServersAwsOfferingResponseDefenderForServersOutput) ToDefenderForServersAwsOfferingResponseDefenderForServersPtrOutput() DefenderForServersAwsOfferingResponseDefenderForServersPtrOutput {
	return o.ToDefenderForServersAwsOfferingResponseDefenderForServersPtrOutputWithContext(context.Background())
}

func (o DefenderForServersAwsOfferingResponseDefenderForServersOutput) ToDefenderForServersAwsOfferingResponseDefenderForServersPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingResponseDefenderForServersPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefenderForServersAwsOfferingResponseDefenderForServers) *DefenderForServersAwsOfferingResponseDefenderForServers {
		return &v
	}).(DefenderForServersAwsOfferingResponseDefenderForServersPtrOutput)
}

func (o DefenderForServersAwsOfferingResponseDefenderForServersOutput) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DefenderForServersAwsOfferingResponseDefenderForServers) *string { return v.CloudRoleArn }).(pulumi.StringPtrOutput)
}

type DefenderForServersAwsOfferingResponseDefenderForServersPtrOutput struct{ *pulumi.OutputState }

func (DefenderForServersAwsOfferingResponseDefenderForServersPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForServersAwsOfferingResponseDefenderForServers)(nil)).Elem()
}

func (o DefenderForServersAwsOfferingResponseDefenderForServersPtrOutput) ToDefenderForServersAwsOfferingResponseDefenderForServersPtrOutput() DefenderForServersAwsOfferingResponseDefenderForServersPtrOutput {
	return o
}

func (o DefenderForServersAwsOfferingResponseDefenderForServersPtrOutput) ToDefenderForServersAwsOfferingResponseDefenderForServersPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingResponseDefenderForServersPtrOutput {
	return o
}

func (o DefenderForServersAwsOfferingResponseDefenderForServersPtrOutput) Elem() DefenderForServersAwsOfferingResponseDefenderForServersOutput {
	return o.ApplyT(func(v *DefenderForServersAwsOfferingResponseDefenderForServers) DefenderForServersAwsOfferingResponseDefenderForServers {
		if v != nil {
			return *v
		}
		var ret DefenderForServersAwsOfferingResponseDefenderForServers
		return ret
	}).(DefenderForServersAwsOfferingResponseDefenderForServersOutput)
}

func (o DefenderForServersAwsOfferingResponseDefenderForServersPtrOutput) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefenderForServersAwsOfferingResponseDefenderForServers) *string {
		if v == nil {
			return nil
		}
		return v.CloudRoleArn
	}).(pulumi.StringPtrOutput)
}

type DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadata struct {
	ExpiryDate           *string `pulumi:"expiryDate"`
	ParameterNameInStore *string `pulumi:"parameterNameInStore"`
	ParameterStoreRegion *string `pulumi:"parameterStoreRegion"`
}





type DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataInput interface {
	pulumi.Input

	ToDefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutput() DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutput
	ToDefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutputWithContext(context.Context) DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutput
}

type DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataArgs struct {
	ExpiryDate           pulumi.StringPtrInput `pulumi:"expiryDate"`
	ParameterNameInStore pulumi.StringPtrInput `pulumi:"parameterNameInStore"`
	ParameterStoreRegion pulumi.StringPtrInput `pulumi:"parameterStoreRegion"`
}

func (DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadata)(nil)).Elem()
}

func (i DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataArgs) ToDefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutput() DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutput {
	return i.ToDefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutputWithContext(context.Background())
}

func (i DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataArgs) ToDefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutput)
}

func (i DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataArgs) ToDefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput() DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput {
	return i.ToDefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutputWithContext(context.Background())
}

func (i DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataArgs) ToDefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutput).ToDefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutputWithContext(ctx)
}









type DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrInput interface {
	pulumi.Input

	ToDefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput() DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput
	ToDefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutputWithContext(context.Context) DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput
}

type defenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrType DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataArgs

func DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtr(v *DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataArgs) DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrInput {
	return (*defenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrType)(v)
}

func (*defenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadata)(nil)).Elem()
}

func (i *defenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrType) ToDefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput() DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput {
	return i.ToDefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutputWithContext(context.Background())
}

func (i *defenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrType) ToDefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput)
}

type DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutput struct{ *pulumi.OutputState }

func (DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadata)(nil)).Elem()
}

func (o DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutput) ToDefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutput() DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutput {
	return o
}

func (o DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutput) ToDefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutput {
	return o
}

func (o DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutput) ToDefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput() DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput {
	return o.ToDefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutputWithContext(context.Background())
}

func (o DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutput) ToDefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadata) *DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadata {
		return &v
	}).(DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput)
}

func (o DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutput) ExpiryDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadata) *string {
		return v.ExpiryDate
	}).(pulumi.StringPtrOutput)
}

func (o DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutput) ParameterNameInStore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadata) *string {
		return v.ParameterNameInStore
	}).(pulumi.StringPtrOutput)
}

func (o DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutput) ParameterStoreRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadata) *string {
		return v.ParameterStoreRegion
	}).(pulumi.StringPtrOutput)
}

type DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput struct{ *pulumi.OutputState }

func (DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadata)(nil)).Elem()
}

func (o DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput) ToDefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput() DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput {
	return o
}

func (o DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput) ToDefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput {
	return o
}

func (o DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput) Elem() DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutput {
	return o.ApplyT(func(v *DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadata) DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadata {
		if v != nil {
			return *v
		}
		var ret DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadata
		return ret
	}).(DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutput)
}

func (o DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput) ExpiryDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadata) *string {
		if v == nil {
			return nil
		}
		return v.ExpiryDate
	}).(pulumi.StringPtrOutput)
}

func (o DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput) ParameterNameInStore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadata) *string {
		if v == nil {
			return nil
		}
		return v.ParameterNameInStore
	}).(pulumi.StringPtrOutput)
}

func (o DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput) ParameterStoreRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadata) *string {
		if v == nil {
			return nil
		}
		return v.ParameterStoreRegion
	}).(pulumi.StringPtrOutput)
}

type DefenderForServersAwsOfferingServicePrincipalSecretMetadata struct {
	ExpiryDate           *string `pulumi:"expiryDate"`
	ParameterNameInStore *string `pulumi:"parameterNameInStore"`
	ParameterStoreRegion *string `pulumi:"parameterStoreRegion"`
}





type DefenderForServersAwsOfferingServicePrincipalSecretMetadataInput interface {
	pulumi.Input

	ToDefenderForServersAwsOfferingServicePrincipalSecretMetadataOutput() DefenderForServersAwsOfferingServicePrincipalSecretMetadataOutput
	ToDefenderForServersAwsOfferingServicePrincipalSecretMetadataOutputWithContext(context.Context) DefenderForServersAwsOfferingServicePrincipalSecretMetadataOutput
}

type DefenderForServersAwsOfferingServicePrincipalSecretMetadataArgs struct {
	ExpiryDate           pulumi.StringPtrInput `pulumi:"expiryDate"`
	ParameterNameInStore pulumi.StringPtrInput `pulumi:"parameterNameInStore"`
	ParameterStoreRegion pulumi.StringPtrInput `pulumi:"parameterStoreRegion"`
}

func (DefenderForServersAwsOfferingServicePrincipalSecretMetadataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForServersAwsOfferingServicePrincipalSecretMetadata)(nil)).Elem()
}

func (i DefenderForServersAwsOfferingServicePrincipalSecretMetadataArgs) ToDefenderForServersAwsOfferingServicePrincipalSecretMetadataOutput() DefenderForServersAwsOfferingServicePrincipalSecretMetadataOutput {
	return i.ToDefenderForServersAwsOfferingServicePrincipalSecretMetadataOutputWithContext(context.Background())
}

func (i DefenderForServersAwsOfferingServicePrincipalSecretMetadataArgs) ToDefenderForServersAwsOfferingServicePrincipalSecretMetadataOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingServicePrincipalSecretMetadataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForServersAwsOfferingServicePrincipalSecretMetadataOutput)
}

func (i DefenderForServersAwsOfferingServicePrincipalSecretMetadataArgs) ToDefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput() DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput {
	return i.ToDefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutputWithContext(context.Background())
}

func (i DefenderForServersAwsOfferingServicePrincipalSecretMetadataArgs) ToDefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForServersAwsOfferingServicePrincipalSecretMetadataOutput).ToDefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutputWithContext(ctx)
}









type DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrInput interface {
	pulumi.Input

	ToDefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput() DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput
	ToDefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutputWithContext(context.Context) DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput
}

type defenderForServersAwsOfferingServicePrincipalSecretMetadataPtrType DefenderForServersAwsOfferingServicePrincipalSecretMetadataArgs

func DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtr(v *DefenderForServersAwsOfferingServicePrincipalSecretMetadataArgs) DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrInput {
	return (*defenderForServersAwsOfferingServicePrincipalSecretMetadataPtrType)(v)
}

func (*defenderForServersAwsOfferingServicePrincipalSecretMetadataPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForServersAwsOfferingServicePrincipalSecretMetadata)(nil)).Elem()
}

func (i *defenderForServersAwsOfferingServicePrincipalSecretMetadataPtrType) ToDefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput() DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput {
	return i.ToDefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutputWithContext(context.Background())
}

func (i *defenderForServersAwsOfferingServicePrincipalSecretMetadataPtrType) ToDefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput)
}

type DefenderForServersAwsOfferingServicePrincipalSecretMetadataOutput struct{ *pulumi.OutputState }

func (DefenderForServersAwsOfferingServicePrincipalSecretMetadataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForServersAwsOfferingServicePrincipalSecretMetadata)(nil)).Elem()
}

func (o DefenderForServersAwsOfferingServicePrincipalSecretMetadataOutput) ToDefenderForServersAwsOfferingServicePrincipalSecretMetadataOutput() DefenderForServersAwsOfferingServicePrincipalSecretMetadataOutput {
	return o
}

func (o DefenderForServersAwsOfferingServicePrincipalSecretMetadataOutput) ToDefenderForServersAwsOfferingServicePrincipalSecretMetadataOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingServicePrincipalSecretMetadataOutput {
	return o
}

func (o DefenderForServersAwsOfferingServicePrincipalSecretMetadataOutput) ToDefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput() DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput {
	return o.ToDefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutputWithContext(context.Background())
}

func (o DefenderForServersAwsOfferingServicePrincipalSecretMetadataOutput) ToDefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefenderForServersAwsOfferingServicePrincipalSecretMetadata) *DefenderForServersAwsOfferingServicePrincipalSecretMetadata {
		return &v
	}).(DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput)
}

func (o DefenderForServersAwsOfferingServicePrincipalSecretMetadataOutput) ExpiryDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DefenderForServersAwsOfferingServicePrincipalSecretMetadata) *string { return v.ExpiryDate }).(pulumi.StringPtrOutput)
}

func (o DefenderForServersAwsOfferingServicePrincipalSecretMetadataOutput) ParameterNameInStore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DefenderForServersAwsOfferingServicePrincipalSecretMetadata) *string {
		return v.ParameterNameInStore
	}).(pulumi.StringPtrOutput)
}

func (o DefenderForServersAwsOfferingServicePrincipalSecretMetadataOutput) ParameterStoreRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DefenderForServersAwsOfferingServicePrincipalSecretMetadata) *string {
		return v.ParameterStoreRegion
	}).(pulumi.StringPtrOutput)
}

type DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput struct{ *pulumi.OutputState }

func (DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForServersAwsOfferingServicePrincipalSecretMetadata)(nil)).Elem()
}

func (o DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput) ToDefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput() DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput {
	return o
}

func (o DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput) ToDefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput {
	return o
}

func (o DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput) Elem() DefenderForServersAwsOfferingServicePrincipalSecretMetadataOutput {
	return o.ApplyT(func(v *DefenderForServersAwsOfferingServicePrincipalSecretMetadata) DefenderForServersAwsOfferingServicePrincipalSecretMetadata {
		if v != nil {
			return *v
		}
		var ret DefenderForServersAwsOfferingServicePrincipalSecretMetadata
		return ret
	}).(DefenderForServersAwsOfferingServicePrincipalSecretMetadataOutput)
}

func (o DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput) ExpiryDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefenderForServersAwsOfferingServicePrincipalSecretMetadata) *string {
		if v == nil {
			return nil
		}
		return v.ExpiryDate
	}).(pulumi.StringPtrOutput)
}

func (o DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput) ParameterNameInStore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefenderForServersAwsOfferingServicePrincipalSecretMetadata) *string {
		if v == nil {
			return nil
		}
		return v.ParameterNameInStore
	}).(pulumi.StringPtrOutput)
}

func (o DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput) ParameterStoreRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefenderForServersAwsOfferingServicePrincipalSecretMetadata) *string {
		if v == nil {
			return nil
		}
		return v.ParameterStoreRegion
	}).(pulumi.StringPtrOutput)
}

type SecurityConnectorPropertiesOrganizationalData struct {
	ExcludedAccountId          []string `pulumi:"excludedAccountId"`
	OrganizationMembershipType *string  `pulumi:"organizationMembershipType"`
	ParentHierarchyId          *string  `pulumi:"parentHierarchyId"`
	StacksetName               *string  `pulumi:"stacksetName"`
}





type SecurityConnectorPropertiesOrganizationalDataInput interface {
	pulumi.Input

	ToSecurityConnectorPropertiesOrganizationalDataOutput() SecurityConnectorPropertiesOrganizationalDataOutput
	ToSecurityConnectorPropertiesOrganizationalDataOutputWithContext(context.Context) SecurityConnectorPropertiesOrganizationalDataOutput
}

type SecurityConnectorPropertiesOrganizationalDataArgs struct {
	ExcludedAccountId          pulumi.StringArrayInput `pulumi:"excludedAccountId"`
	OrganizationMembershipType pulumi.StringPtrInput   `pulumi:"organizationMembershipType"`
	ParentHierarchyId          pulumi.StringPtrInput   `pulumi:"parentHierarchyId"`
	StacksetName               pulumi.StringPtrInput   `pulumi:"stacksetName"`
}

func (SecurityConnectorPropertiesOrganizationalDataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityConnectorPropertiesOrganizationalData)(nil)).Elem()
}

func (i SecurityConnectorPropertiesOrganizationalDataArgs) ToSecurityConnectorPropertiesOrganizationalDataOutput() SecurityConnectorPropertiesOrganizationalDataOutput {
	return i.ToSecurityConnectorPropertiesOrganizationalDataOutputWithContext(context.Background())
}

func (i SecurityConnectorPropertiesOrganizationalDataArgs) ToSecurityConnectorPropertiesOrganizationalDataOutputWithContext(ctx context.Context) SecurityConnectorPropertiesOrganizationalDataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityConnectorPropertiesOrganizationalDataOutput)
}

func (i SecurityConnectorPropertiesOrganizationalDataArgs) ToSecurityConnectorPropertiesOrganizationalDataPtrOutput() SecurityConnectorPropertiesOrganizationalDataPtrOutput {
	return i.ToSecurityConnectorPropertiesOrganizationalDataPtrOutputWithContext(context.Background())
}

func (i SecurityConnectorPropertiesOrganizationalDataArgs) ToSecurityConnectorPropertiesOrganizationalDataPtrOutputWithContext(ctx context.Context) SecurityConnectorPropertiesOrganizationalDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityConnectorPropertiesOrganizationalDataOutput).ToSecurityConnectorPropertiesOrganizationalDataPtrOutputWithContext(ctx)
}









type SecurityConnectorPropertiesOrganizationalDataPtrInput interface {
	pulumi.Input

	ToSecurityConnectorPropertiesOrganizationalDataPtrOutput() SecurityConnectorPropertiesOrganizationalDataPtrOutput
	ToSecurityConnectorPropertiesOrganizationalDataPtrOutputWithContext(context.Context) SecurityConnectorPropertiesOrganizationalDataPtrOutput
}

type securityConnectorPropertiesOrganizationalDataPtrType SecurityConnectorPropertiesOrganizationalDataArgs

func SecurityConnectorPropertiesOrganizationalDataPtr(v *SecurityConnectorPropertiesOrganizationalDataArgs) SecurityConnectorPropertiesOrganizationalDataPtrInput {
	return (*securityConnectorPropertiesOrganizationalDataPtrType)(v)
}

func (*securityConnectorPropertiesOrganizationalDataPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityConnectorPropertiesOrganizationalData)(nil)).Elem()
}

func (i *securityConnectorPropertiesOrganizationalDataPtrType) ToSecurityConnectorPropertiesOrganizationalDataPtrOutput() SecurityConnectorPropertiesOrganizationalDataPtrOutput {
	return i.ToSecurityConnectorPropertiesOrganizationalDataPtrOutputWithContext(context.Background())
}

func (i *securityConnectorPropertiesOrganizationalDataPtrType) ToSecurityConnectorPropertiesOrganizationalDataPtrOutputWithContext(ctx context.Context) SecurityConnectorPropertiesOrganizationalDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityConnectorPropertiesOrganizationalDataPtrOutput)
}

type SecurityConnectorPropertiesOrganizationalDataOutput struct{ *pulumi.OutputState }

func (SecurityConnectorPropertiesOrganizationalDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityConnectorPropertiesOrganizationalData)(nil)).Elem()
}

func (o SecurityConnectorPropertiesOrganizationalDataOutput) ToSecurityConnectorPropertiesOrganizationalDataOutput() SecurityConnectorPropertiesOrganizationalDataOutput {
	return o
}

func (o SecurityConnectorPropertiesOrganizationalDataOutput) ToSecurityConnectorPropertiesOrganizationalDataOutputWithContext(ctx context.Context) SecurityConnectorPropertiesOrganizationalDataOutput {
	return o
}

func (o SecurityConnectorPropertiesOrganizationalDataOutput) ToSecurityConnectorPropertiesOrganizationalDataPtrOutput() SecurityConnectorPropertiesOrganizationalDataPtrOutput {
	return o.ToSecurityConnectorPropertiesOrganizationalDataPtrOutputWithContext(context.Background())
}

func (o SecurityConnectorPropertiesOrganizationalDataOutput) ToSecurityConnectorPropertiesOrganizationalDataPtrOutputWithContext(ctx context.Context) SecurityConnectorPropertiesOrganizationalDataPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecurityConnectorPropertiesOrganizationalData) *SecurityConnectorPropertiesOrganizationalData {
		return &v
	}).(SecurityConnectorPropertiesOrganizationalDataPtrOutput)
}

func (o SecurityConnectorPropertiesOrganizationalDataOutput) ExcludedAccountId() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityConnectorPropertiesOrganizationalData) []string { return v.ExcludedAccountId }).(pulumi.StringArrayOutput)
}

func (o SecurityConnectorPropertiesOrganizationalDataOutput) OrganizationMembershipType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityConnectorPropertiesOrganizationalData) *string { return v.OrganizationMembershipType }).(pulumi.StringPtrOutput)
}

func (o SecurityConnectorPropertiesOrganizationalDataOutput) ParentHierarchyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityConnectorPropertiesOrganizationalData) *string { return v.ParentHierarchyId }).(pulumi.StringPtrOutput)
}

func (o SecurityConnectorPropertiesOrganizationalDataOutput) StacksetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityConnectorPropertiesOrganizationalData) *string { return v.StacksetName }).(pulumi.StringPtrOutput)
}

type SecurityConnectorPropertiesOrganizationalDataPtrOutput struct{ *pulumi.OutputState }

func (SecurityConnectorPropertiesOrganizationalDataPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityConnectorPropertiesOrganizationalData)(nil)).Elem()
}

func (o SecurityConnectorPropertiesOrganizationalDataPtrOutput) ToSecurityConnectorPropertiesOrganizationalDataPtrOutput() SecurityConnectorPropertiesOrganizationalDataPtrOutput {
	return o
}

func (o SecurityConnectorPropertiesOrganizationalDataPtrOutput) ToSecurityConnectorPropertiesOrganizationalDataPtrOutputWithContext(ctx context.Context) SecurityConnectorPropertiesOrganizationalDataPtrOutput {
	return o
}

func (o SecurityConnectorPropertiesOrganizationalDataPtrOutput) Elem() SecurityConnectorPropertiesOrganizationalDataOutput {
	return o.ApplyT(func(v *SecurityConnectorPropertiesOrganizationalData) SecurityConnectorPropertiesOrganizationalData {
		if v != nil {
			return *v
		}
		var ret SecurityConnectorPropertiesOrganizationalData
		return ret
	}).(SecurityConnectorPropertiesOrganizationalDataOutput)
}

func (o SecurityConnectorPropertiesOrganizationalDataPtrOutput) ExcludedAccountId() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecurityConnectorPropertiesOrganizationalData) []string {
		if v == nil {
			return nil
		}
		return v.ExcludedAccountId
	}).(pulumi.StringArrayOutput)
}

func (o SecurityConnectorPropertiesOrganizationalDataPtrOutput) OrganizationMembershipType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityConnectorPropertiesOrganizationalData) *string {
		if v == nil {
			return nil
		}
		return v.OrganizationMembershipType
	}).(pulumi.StringPtrOutput)
}

func (o SecurityConnectorPropertiesOrganizationalDataPtrOutput) ParentHierarchyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityConnectorPropertiesOrganizationalData) *string {
		if v == nil {
			return nil
		}
		return v.ParentHierarchyId
	}).(pulumi.StringPtrOutput)
}

func (o SecurityConnectorPropertiesOrganizationalDataPtrOutput) StacksetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityConnectorPropertiesOrganizationalData) *string {
		if v == nil {
			return nil
		}
		return v.StacksetName
	}).(pulumi.StringPtrOutput)
}

type SecurityConnectorPropertiesResponseOrganizationalData struct {
	ExcludedAccountId          []string `pulumi:"excludedAccountId"`
	OrganizationMembershipType *string  `pulumi:"organizationMembershipType"`
	ParentHierarchyId          *string  `pulumi:"parentHierarchyId"`
	StacksetName               *string  `pulumi:"stacksetName"`
}





type SecurityConnectorPropertiesResponseOrganizationalDataInput interface {
	pulumi.Input

	ToSecurityConnectorPropertiesResponseOrganizationalDataOutput() SecurityConnectorPropertiesResponseOrganizationalDataOutput
	ToSecurityConnectorPropertiesResponseOrganizationalDataOutputWithContext(context.Context) SecurityConnectorPropertiesResponseOrganizationalDataOutput
}

type SecurityConnectorPropertiesResponseOrganizationalDataArgs struct {
	ExcludedAccountId          pulumi.StringArrayInput `pulumi:"excludedAccountId"`
	OrganizationMembershipType pulumi.StringPtrInput   `pulumi:"organizationMembershipType"`
	ParentHierarchyId          pulumi.StringPtrInput   `pulumi:"parentHierarchyId"`
	StacksetName               pulumi.StringPtrInput   `pulumi:"stacksetName"`
}

func (SecurityConnectorPropertiesResponseOrganizationalDataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityConnectorPropertiesResponseOrganizationalData)(nil)).Elem()
}

func (i SecurityConnectorPropertiesResponseOrganizationalDataArgs) ToSecurityConnectorPropertiesResponseOrganizationalDataOutput() SecurityConnectorPropertiesResponseOrganizationalDataOutput {
	return i.ToSecurityConnectorPropertiesResponseOrganizationalDataOutputWithContext(context.Background())
}

func (i SecurityConnectorPropertiesResponseOrganizationalDataArgs) ToSecurityConnectorPropertiesResponseOrganizationalDataOutputWithContext(ctx context.Context) SecurityConnectorPropertiesResponseOrganizationalDataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityConnectorPropertiesResponseOrganizationalDataOutput)
}

func (i SecurityConnectorPropertiesResponseOrganizationalDataArgs) ToSecurityConnectorPropertiesResponseOrganizationalDataPtrOutput() SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput {
	return i.ToSecurityConnectorPropertiesResponseOrganizationalDataPtrOutputWithContext(context.Background())
}

func (i SecurityConnectorPropertiesResponseOrganizationalDataArgs) ToSecurityConnectorPropertiesResponseOrganizationalDataPtrOutputWithContext(ctx context.Context) SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityConnectorPropertiesResponseOrganizationalDataOutput).ToSecurityConnectorPropertiesResponseOrganizationalDataPtrOutputWithContext(ctx)
}









type SecurityConnectorPropertiesResponseOrganizationalDataPtrInput interface {
	pulumi.Input

	ToSecurityConnectorPropertiesResponseOrganizationalDataPtrOutput() SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput
	ToSecurityConnectorPropertiesResponseOrganizationalDataPtrOutputWithContext(context.Context) SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput
}

type securityConnectorPropertiesResponseOrganizationalDataPtrType SecurityConnectorPropertiesResponseOrganizationalDataArgs

func SecurityConnectorPropertiesResponseOrganizationalDataPtr(v *SecurityConnectorPropertiesResponseOrganizationalDataArgs) SecurityConnectorPropertiesResponseOrganizationalDataPtrInput {
	return (*securityConnectorPropertiesResponseOrganizationalDataPtrType)(v)
}

func (*securityConnectorPropertiesResponseOrganizationalDataPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityConnectorPropertiesResponseOrganizationalData)(nil)).Elem()
}

func (i *securityConnectorPropertiesResponseOrganizationalDataPtrType) ToSecurityConnectorPropertiesResponseOrganizationalDataPtrOutput() SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput {
	return i.ToSecurityConnectorPropertiesResponseOrganizationalDataPtrOutputWithContext(context.Background())
}

func (i *securityConnectorPropertiesResponseOrganizationalDataPtrType) ToSecurityConnectorPropertiesResponseOrganizationalDataPtrOutputWithContext(ctx context.Context) SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput)
}

type SecurityConnectorPropertiesResponseOrganizationalDataOutput struct{ *pulumi.OutputState }

func (SecurityConnectorPropertiesResponseOrganizationalDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityConnectorPropertiesResponseOrganizationalData)(nil)).Elem()
}

func (o SecurityConnectorPropertiesResponseOrganizationalDataOutput) ToSecurityConnectorPropertiesResponseOrganizationalDataOutput() SecurityConnectorPropertiesResponseOrganizationalDataOutput {
	return o
}

func (o SecurityConnectorPropertiesResponseOrganizationalDataOutput) ToSecurityConnectorPropertiesResponseOrganizationalDataOutputWithContext(ctx context.Context) SecurityConnectorPropertiesResponseOrganizationalDataOutput {
	return o
}

func (o SecurityConnectorPropertiesResponseOrganizationalDataOutput) ToSecurityConnectorPropertiesResponseOrganizationalDataPtrOutput() SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput {
	return o.ToSecurityConnectorPropertiesResponseOrganizationalDataPtrOutputWithContext(context.Background())
}

func (o SecurityConnectorPropertiesResponseOrganizationalDataOutput) ToSecurityConnectorPropertiesResponseOrganizationalDataPtrOutputWithContext(ctx context.Context) SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecurityConnectorPropertiesResponseOrganizationalData) *SecurityConnectorPropertiesResponseOrganizationalData {
		return &v
	}).(SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput)
}

func (o SecurityConnectorPropertiesResponseOrganizationalDataOutput) ExcludedAccountId() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityConnectorPropertiesResponseOrganizationalData) []string { return v.ExcludedAccountId }).(pulumi.StringArrayOutput)
}

func (o SecurityConnectorPropertiesResponseOrganizationalDataOutput) OrganizationMembershipType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityConnectorPropertiesResponseOrganizationalData) *string {
		return v.OrganizationMembershipType
	}).(pulumi.StringPtrOutput)
}

func (o SecurityConnectorPropertiesResponseOrganizationalDataOutput) ParentHierarchyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityConnectorPropertiesResponseOrganizationalData) *string { return v.ParentHierarchyId }).(pulumi.StringPtrOutput)
}

func (o SecurityConnectorPropertiesResponseOrganizationalDataOutput) StacksetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityConnectorPropertiesResponseOrganizationalData) *string { return v.StacksetName }).(pulumi.StringPtrOutput)
}

type SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput struct{ *pulumi.OutputState }

func (SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityConnectorPropertiesResponseOrganizationalData)(nil)).Elem()
}

func (o SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput) ToSecurityConnectorPropertiesResponseOrganizationalDataPtrOutput() SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput {
	return o
}

func (o SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput) ToSecurityConnectorPropertiesResponseOrganizationalDataPtrOutputWithContext(ctx context.Context) SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput {
	return o
}

func (o SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput) Elem() SecurityConnectorPropertiesResponseOrganizationalDataOutput {
	return o.ApplyT(func(v *SecurityConnectorPropertiesResponseOrganizationalData) SecurityConnectorPropertiesResponseOrganizationalData {
		if v != nil {
			return *v
		}
		var ret SecurityConnectorPropertiesResponseOrganizationalData
		return ret
	}).(SecurityConnectorPropertiesResponseOrganizationalDataOutput)
}

func (o SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput) ExcludedAccountId() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecurityConnectorPropertiesResponseOrganizationalData) []string {
		if v == nil {
			return nil
		}
		return v.ExcludedAccountId
	}).(pulumi.StringArrayOutput)
}

func (o SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput) OrganizationMembershipType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityConnectorPropertiesResponseOrganizationalData) *string {
		if v == nil {
			return nil
		}
		return v.OrganizationMembershipType
	}).(pulumi.StringPtrOutput)
}

func (o SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput) ParentHierarchyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityConnectorPropertiesResponseOrganizationalData) *string {
		if v == nil {
			return nil
		}
		return v.ParentHierarchyId
	}).(pulumi.StringPtrOutput)
}

func (o SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput) StacksetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityConnectorPropertiesResponseOrganizationalData) *string {
		if v == nil {
			return nil
		}
		return v.StacksetName
	}).(pulumi.StringPtrOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}





type SystemDataResponseInput interface {
	pulumi.Input

	ToSystemDataResponseOutput() SystemDataResponseOutput
	ToSystemDataResponseOutputWithContext(context.Context) SystemDataResponseOutput
}

type SystemDataResponseArgs struct {
	CreatedAt          pulumi.StringPtrInput `pulumi:"createdAt"`
	CreatedBy          pulumi.StringPtrInput `pulumi:"createdBy"`
	CreatedByType      pulumi.StringPtrInput `pulumi:"createdByType"`
	LastModifiedAt     pulumi.StringPtrInput `pulumi:"lastModifiedAt"`
	LastModifiedBy     pulumi.StringPtrInput `pulumi:"lastModifiedBy"`
	LastModifiedByType pulumi.StringPtrInput `pulumi:"lastModifiedByType"`
}

func (SystemDataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return i.ToSystemDataResponseOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput)
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput).ToSystemDataResponsePtrOutputWithContext(ctx)
}









type SystemDataResponsePtrInput interface {
	pulumi.Input

	ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput
	ToSystemDataResponsePtrOutputWithContext(context.Context) SystemDataResponsePtrOutput
}

type systemDataResponsePtrType SystemDataResponseArgs

func SystemDataResponsePtr(v *SystemDataResponseArgs) SystemDataResponsePtrInput {
	return (*systemDataResponsePtrType)(v)
}

func (*systemDataResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponsePtrOutput)
}

type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemDataResponse) *SystemDataResponse {
		return &v
	}).(SystemDataResponsePtrOutput)
}

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

type SystemDataResponsePtrOutput struct{ *pulumi.OutputState }

func (SystemDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) Elem() SystemDataResponseOutput {
	return o.ApplyT(func(v *SystemDataResponse) SystemDataResponse {
		if v != nil {
			return *v
		}
		var ret SystemDataResponse
		return ret
	}).(SystemDataResponseOutput)
}

func (o SystemDataResponsePtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedByType
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedByType
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CspmMonitorAwsOfferingInput)(nil)).Elem(), CspmMonitorAwsOfferingArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CspmMonitorAwsOfferingNativeCloudConnectionInput)(nil)).Elem(), CspmMonitorAwsOfferingNativeCloudConnectionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CspmMonitorAwsOfferingNativeCloudConnectionPtrInput)(nil)).Elem(), CspmMonitorAwsOfferingNativeCloudConnectionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CspmMonitorAwsOfferingResponseInput)(nil)).Elem(), CspmMonitorAwsOfferingResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CspmMonitorAwsOfferingResponseNativeCloudConnectionInput)(nil)).Elem(), CspmMonitorAwsOfferingResponseNativeCloudConnectionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrInput)(nil)).Elem(), CspmMonitorAwsOfferingResponseNativeCloudConnectionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefenderForContainersAwsOfferingInput)(nil)).Elem(), DefenderForContainersAwsOfferingArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefenderForContainersAwsOfferingCloudWatchToKinesisInput)(nil)).Elem(), DefenderForContainersAwsOfferingCloudWatchToKinesisArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefenderForContainersAwsOfferingCloudWatchToKinesisPtrInput)(nil)).Elem(), DefenderForContainersAwsOfferingCloudWatchToKinesisArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefenderForContainersAwsOfferingKinesisToS3Input)(nil)).Elem(), DefenderForContainersAwsOfferingKinesisToS3Args{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefenderForContainersAwsOfferingKinesisToS3PtrInput)(nil)).Elem(), DefenderForContainersAwsOfferingKinesisToS3Args{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefenderForContainersAwsOfferingKubernetesScubaReaderInput)(nil)).Elem(), DefenderForContainersAwsOfferingKubernetesScubaReaderArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefenderForContainersAwsOfferingKubernetesScubaReaderPtrInput)(nil)).Elem(), DefenderForContainersAwsOfferingKubernetesScubaReaderArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefenderForContainersAwsOfferingKubernetesServiceInput)(nil)).Elem(), DefenderForContainersAwsOfferingKubernetesServiceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefenderForContainersAwsOfferingKubernetesServicePtrInput)(nil)).Elem(), DefenderForContainersAwsOfferingKubernetesServiceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefenderForContainersAwsOfferingResponseInput)(nil)).Elem(), DefenderForContainersAwsOfferingResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefenderForContainersAwsOfferingResponseCloudWatchToKinesisInput)(nil)).Elem(), DefenderForContainersAwsOfferingResponseCloudWatchToKinesisArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrInput)(nil)).Elem(), DefenderForContainersAwsOfferingResponseCloudWatchToKinesisArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefenderForContainersAwsOfferingResponseKinesisToS3Input)(nil)).Elem(), DefenderForContainersAwsOfferingResponseKinesisToS3Args{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefenderForContainersAwsOfferingResponseKinesisToS3PtrInput)(nil)).Elem(), DefenderForContainersAwsOfferingResponseKinesisToS3Args{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefenderForContainersAwsOfferingResponseKubernetesScubaReaderInput)(nil)).Elem(), DefenderForContainersAwsOfferingResponseKubernetesScubaReaderArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrInput)(nil)).Elem(), DefenderForContainersAwsOfferingResponseKubernetesScubaReaderArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefenderForContainersAwsOfferingResponseKubernetesServiceInput)(nil)).Elem(), DefenderForContainersAwsOfferingResponseKubernetesServiceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefenderForContainersAwsOfferingResponseKubernetesServicePtrInput)(nil)).Elem(), DefenderForContainersAwsOfferingResponseKubernetesServiceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefenderForServersAwsOfferingInput)(nil)).Elem(), DefenderForServersAwsOfferingArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefenderForServersAwsOfferingArcAutoProvisioningInput)(nil)).Elem(), DefenderForServersAwsOfferingArcAutoProvisioningArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefenderForServersAwsOfferingArcAutoProvisioningPtrInput)(nil)).Elem(), DefenderForServersAwsOfferingArcAutoProvisioningArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefenderForServersAwsOfferingDefenderForServersInput)(nil)).Elem(), DefenderForServersAwsOfferingDefenderForServersArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefenderForServersAwsOfferingDefenderForServersPtrInput)(nil)).Elem(), DefenderForServersAwsOfferingDefenderForServersArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefenderForServersAwsOfferingResponseInput)(nil)).Elem(), DefenderForServersAwsOfferingResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefenderForServersAwsOfferingResponseArcAutoProvisioningInput)(nil)).Elem(), DefenderForServersAwsOfferingResponseArcAutoProvisioningArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrInput)(nil)).Elem(), DefenderForServersAwsOfferingResponseArcAutoProvisioningArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefenderForServersAwsOfferingResponseDefenderForServersInput)(nil)).Elem(), DefenderForServersAwsOfferingResponseDefenderForServersArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefenderForServersAwsOfferingResponseDefenderForServersPtrInput)(nil)).Elem(), DefenderForServersAwsOfferingResponseDefenderForServersArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataInput)(nil)).Elem(), DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrInput)(nil)).Elem(), DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefenderForServersAwsOfferingServicePrincipalSecretMetadataInput)(nil)).Elem(), DefenderForServersAwsOfferingServicePrincipalSecretMetadataArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrInput)(nil)).Elem(), DefenderForServersAwsOfferingServicePrincipalSecretMetadataArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityConnectorPropertiesOrganizationalDataInput)(nil)).Elem(), SecurityConnectorPropertiesOrganizationalDataArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityConnectorPropertiesOrganizationalDataPtrInput)(nil)).Elem(), SecurityConnectorPropertiesOrganizationalDataArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityConnectorPropertiesResponseOrganizationalDataInput)(nil)).Elem(), SecurityConnectorPropertiesResponseOrganizationalDataArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityConnectorPropertiesResponseOrganizationalDataPtrInput)(nil)).Elem(), SecurityConnectorPropertiesResponseOrganizationalDataArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponseInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponsePtrInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterOutputType(CspmMonitorAwsOfferingOutput{})
	pulumi.RegisterOutputType(CspmMonitorAwsOfferingNativeCloudConnectionOutput{})
	pulumi.RegisterOutputType(CspmMonitorAwsOfferingNativeCloudConnectionPtrOutput{})
	pulumi.RegisterOutputType(CspmMonitorAwsOfferingResponseOutput{})
	pulumi.RegisterOutputType(CspmMonitorAwsOfferingResponseNativeCloudConnectionOutput{})
	pulumi.RegisterOutputType(CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput{})
	pulumi.RegisterOutputType(DefenderForContainersAwsOfferingOutput{})
	pulumi.RegisterOutputType(DefenderForContainersAwsOfferingCloudWatchToKinesisOutput{})
	pulumi.RegisterOutputType(DefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput{})
	pulumi.RegisterOutputType(DefenderForContainersAwsOfferingKinesisToS3Output{})
	pulumi.RegisterOutputType(DefenderForContainersAwsOfferingKinesisToS3PtrOutput{})
	pulumi.RegisterOutputType(DefenderForContainersAwsOfferingKubernetesScubaReaderOutput{})
	pulumi.RegisterOutputType(DefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput{})
	pulumi.RegisterOutputType(DefenderForContainersAwsOfferingKubernetesServiceOutput{})
	pulumi.RegisterOutputType(DefenderForContainersAwsOfferingKubernetesServicePtrOutput{})
	pulumi.RegisterOutputType(DefenderForContainersAwsOfferingResponseOutput{})
	pulumi.RegisterOutputType(DefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutput{})
	pulumi.RegisterOutputType(DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput{})
	pulumi.RegisterOutputType(DefenderForContainersAwsOfferingResponseKinesisToS3Output{})
	pulumi.RegisterOutputType(DefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput{})
	pulumi.RegisterOutputType(DefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutput{})
	pulumi.RegisterOutputType(DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput{})
	pulumi.RegisterOutputType(DefenderForContainersAwsOfferingResponseKubernetesServiceOutput{})
	pulumi.RegisterOutputType(DefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput{})
	pulumi.RegisterOutputType(DefenderForServersAwsOfferingOutput{})
	pulumi.RegisterOutputType(DefenderForServersAwsOfferingArcAutoProvisioningOutput{})
	pulumi.RegisterOutputType(DefenderForServersAwsOfferingArcAutoProvisioningPtrOutput{})
	pulumi.RegisterOutputType(DefenderForServersAwsOfferingDefenderForServersOutput{})
	pulumi.RegisterOutputType(DefenderForServersAwsOfferingDefenderForServersPtrOutput{})
	pulumi.RegisterOutputType(DefenderForServersAwsOfferingResponseOutput{})
	pulumi.RegisterOutputType(DefenderForServersAwsOfferingResponseArcAutoProvisioningOutput{})
	pulumi.RegisterOutputType(DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput{})
	pulumi.RegisterOutputType(DefenderForServersAwsOfferingResponseDefenderForServersOutput{})
	pulumi.RegisterOutputType(DefenderForServersAwsOfferingResponseDefenderForServersPtrOutput{})
	pulumi.RegisterOutputType(DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutput{})
	pulumi.RegisterOutputType(DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput{})
	pulumi.RegisterOutputType(DefenderForServersAwsOfferingServicePrincipalSecretMetadataOutput{})
	pulumi.RegisterOutputType(DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput{})
	pulumi.RegisterOutputType(SecurityConnectorPropertiesOrganizationalDataOutput{})
	pulumi.RegisterOutputType(SecurityConnectorPropertiesOrganizationalDataPtrOutput{})
	pulumi.RegisterOutputType(SecurityConnectorPropertiesResponseOrganizationalDataOutput{})
	pulumi.RegisterOutputType(SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
}
