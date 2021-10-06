


package v20191001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ExportDeliveryDestination struct {
	Container      string  `pulumi:"container"`
	ResourceId     string  `pulumi:"resourceId"`
	RootFolderPath *string `pulumi:"rootFolderPath"`
}





type ExportDeliveryDestinationInput interface {
	pulumi.Input

	ToExportDeliveryDestinationOutput() ExportDeliveryDestinationOutput
	ToExportDeliveryDestinationOutputWithContext(context.Context) ExportDeliveryDestinationOutput
}

type ExportDeliveryDestinationArgs struct {
	Container      pulumi.StringInput    `pulumi:"container"`
	ResourceId     pulumi.StringInput    `pulumi:"resourceId"`
	RootFolderPath pulumi.StringPtrInput `pulumi:"rootFolderPath"`
}

func (ExportDeliveryDestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportDeliveryDestination)(nil)).Elem()
}

func (i ExportDeliveryDestinationArgs) ToExportDeliveryDestinationOutput() ExportDeliveryDestinationOutput {
	return i.ToExportDeliveryDestinationOutputWithContext(context.Background())
}

func (i ExportDeliveryDestinationArgs) ToExportDeliveryDestinationOutputWithContext(ctx context.Context) ExportDeliveryDestinationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportDeliveryDestinationOutput)
}

func (i ExportDeliveryDestinationArgs) ToExportDeliveryDestinationPtrOutput() ExportDeliveryDestinationPtrOutput {
	return i.ToExportDeliveryDestinationPtrOutputWithContext(context.Background())
}

func (i ExportDeliveryDestinationArgs) ToExportDeliveryDestinationPtrOutputWithContext(ctx context.Context) ExportDeliveryDestinationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportDeliveryDestinationOutput).ToExportDeliveryDestinationPtrOutputWithContext(ctx)
}









type ExportDeliveryDestinationPtrInput interface {
	pulumi.Input

	ToExportDeliveryDestinationPtrOutput() ExportDeliveryDestinationPtrOutput
	ToExportDeliveryDestinationPtrOutputWithContext(context.Context) ExportDeliveryDestinationPtrOutput
}

type exportDeliveryDestinationPtrType ExportDeliveryDestinationArgs

func ExportDeliveryDestinationPtr(v *ExportDeliveryDestinationArgs) ExportDeliveryDestinationPtrInput {
	return (*exportDeliveryDestinationPtrType)(v)
}

func (*exportDeliveryDestinationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportDeliveryDestination)(nil)).Elem()
}

func (i *exportDeliveryDestinationPtrType) ToExportDeliveryDestinationPtrOutput() ExportDeliveryDestinationPtrOutput {
	return i.ToExportDeliveryDestinationPtrOutputWithContext(context.Background())
}

func (i *exportDeliveryDestinationPtrType) ToExportDeliveryDestinationPtrOutputWithContext(ctx context.Context) ExportDeliveryDestinationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportDeliveryDestinationPtrOutput)
}

type ExportDeliveryDestinationOutput struct{ *pulumi.OutputState }

func (ExportDeliveryDestinationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportDeliveryDestination)(nil)).Elem()
}

func (o ExportDeliveryDestinationOutput) ToExportDeliveryDestinationOutput() ExportDeliveryDestinationOutput {
	return o
}

func (o ExportDeliveryDestinationOutput) ToExportDeliveryDestinationOutputWithContext(ctx context.Context) ExportDeliveryDestinationOutput {
	return o
}

func (o ExportDeliveryDestinationOutput) ToExportDeliveryDestinationPtrOutput() ExportDeliveryDestinationPtrOutput {
	return o.ToExportDeliveryDestinationPtrOutputWithContext(context.Background())
}

func (o ExportDeliveryDestinationOutput) ToExportDeliveryDestinationPtrOutputWithContext(ctx context.Context) ExportDeliveryDestinationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExportDeliveryDestination) *ExportDeliveryDestination {
		return &v
	}).(ExportDeliveryDestinationPtrOutput)
}

func (o ExportDeliveryDestinationOutput) Container() pulumi.StringOutput {
	return o.ApplyT(func(v ExportDeliveryDestination) string { return v.Container }).(pulumi.StringOutput)
}

func (o ExportDeliveryDestinationOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v ExportDeliveryDestination) string { return v.ResourceId }).(pulumi.StringOutput)
}

func (o ExportDeliveryDestinationOutput) RootFolderPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportDeliveryDestination) *string { return v.RootFolderPath }).(pulumi.StringPtrOutput)
}

type ExportDeliveryDestinationPtrOutput struct{ *pulumi.OutputState }

func (ExportDeliveryDestinationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportDeliveryDestination)(nil)).Elem()
}

func (o ExportDeliveryDestinationPtrOutput) ToExportDeliveryDestinationPtrOutput() ExportDeliveryDestinationPtrOutput {
	return o
}

func (o ExportDeliveryDestinationPtrOutput) ToExportDeliveryDestinationPtrOutputWithContext(ctx context.Context) ExportDeliveryDestinationPtrOutput {
	return o
}

func (o ExportDeliveryDestinationPtrOutput) Elem() ExportDeliveryDestinationOutput {
	return o.ApplyT(func(v *ExportDeliveryDestination) ExportDeliveryDestination {
		if v != nil {
			return *v
		}
		var ret ExportDeliveryDestination
		return ret
	}).(ExportDeliveryDestinationOutput)
}

func (o ExportDeliveryDestinationPtrOutput) Container() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportDeliveryDestination) *string {
		if v == nil {
			return nil
		}
		return &v.Container
	}).(pulumi.StringPtrOutput)
}

func (o ExportDeliveryDestinationPtrOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportDeliveryDestination) *string {
		if v == nil {
			return nil
		}
		return &v.ResourceId
	}).(pulumi.StringPtrOutput)
}

func (o ExportDeliveryDestinationPtrOutput) RootFolderPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportDeliveryDestination) *string {
		if v == nil {
			return nil
		}
		return v.RootFolderPath
	}).(pulumi.StringPtrOutput)
}

type ExportDeliveryDestinationResponse struct {
	Container      string  `pulumi:"container"`
	ResourceId     string  `pulumi:"resourceId"`
	RootFolderPath *string `pulumi:"rootFolderPath"`
}





type ExportDeliveryDestinationResponseInput interface {
	pulumi.Input

	ToExportDeliveryDestinationResponseOutput() ExportDeliveryDestinationResponseOutput
	ToExportDeliveryDestinationResponseOutputWithContext(context.Context) ExportDeliveryDestinationResponseOutput
}

type ExportDeliveryDestinationResponseArgs struct {
	Container      pulumi.StringInput    `pulumi:"container"`
	ResourceId     pulumi.StringInput    `pulumi:"resourceId"`
	RootFolderPath pulumi.StringPtrInput `pulumi:"rootFolderPath"`
}

func (ExportDeliveryDestinationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportDeliveryDestinationResponse)(nil)).Elem()
}

func (i ExportDeliveryDestinationResponseArgs) ToExportDeliveryDestinationResponseOutput() ExportDeliveryDestinationResponseOutput {
	return i.ToExportDeliveryDestinationResponseOutputWithContext(context.Background())
}

func (i ExportDeliveryDestinationResponseArgs) ToExportDeliveryDestinationResponseOutputWithContext(ctx context.Context) ExportDeliveryDestinationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportDeliveryDestinationResponseOutput)
}

func (i ExportDeliveryDestinationResponseArgs) ToExportDeliveryDestinationResponsePtrOutput() ExportDeliveryDestinationResponsePtrOutput {
	return i.ToExportDeliveryDestinationResponsePtrOutputWithContext(context.Background())
}

func (i ExportDeliveryDestinationResponseArgs) ToExportDeliveryDestinationResponsePtrOutputWithContext(ctx context.Context) ExportDeliveryDestinationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportDeliveryDestinationResponseOutput).ToExportDeliveryDestinationResponsePtrOutputWithContext(ctx)
}









type ExportDeliveryDestinationResponsePtrInput interface {
	pulumi.Input

	ToExportDeliveryDestinationResponsePtrOutput() ExportDeliveryDestinationResponsePtrOutput
	ToExportDeliveryDestinationResponsePtrOutputWithContext(context.Context) ExportDeliveryDestinationResponsePtrOutput
}

type exportDeliveryDestinationResponsePtrType ExportDeliveryDestinationResponseArgs

func ExportDeliveryDestinationResponsePtr(v *ExportDeliveryDestinationResponseArgs) ExportDeliveryDestinationResponsePtrInput {
	return (*exportDeliveryDestinationResponsePtrType)(v)
}

func (*exportDeliveryDestinationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportDeliveryDestinationResponse)(nil)).Elem()
}

func (i *exportDeliveryDestinationResponsePtrType) ToExportDeliveryDestinationResponsePtrOutput() ExportDeliveryDestinationResponsePtrOutput {
	return i.ToExportDeliveryDestinationResponsePtrOutputWithContext(context.Background())
}

func (i *exportDeliveryDestinationResponsePtrType) ToExportDeliveryDestinationResponsePtrOutputWithContext(ctx context.Context) ExportDeliveryDestinationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportDeliveryDestinationResponsePtrOutput)
}

type ExportDeliveryDestinationResponseOutput struct{ *pulumi.OutputState }

func (ExportDeliveryDestinationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportDeliveryDestinationResponse)(nil)).Elem()
}

func (o ExportDeliveryDestinationResponseOutput) ToExportDeliveryDestinationResponseOutput() ExportDeliveryDestinationResponseOutput {
	return o
}

func (o ExportDeliveryDestinationResponseOutput) ToExportDeliveryDestinationResponseOutputWithContext(ctx context.Context) ExportDeliveryDestinationResponseOutput {
	return o
}

func (o ExportDeliveryDestinationResponseOutput) ToExportDeliveryDestinationResponsePtrOutput() ExportDeliveryDestinationResponsePtrOutput {
	return o.ToExportDeliveryDestinationResponsePtrOutputWithContext(context.Background())
}

func (o ExportDeliveryDestinationResponseOutput) ToExportDeliveryDestinationResponsePtrOutputWithContext(ctx context.Context) ExportDeliveryDestinationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExportDeliveryDestinationResponse) *ExportDeliveryDestinationResponse {
		return &v
	}).(ExportDeliveryDestinationResponsePtrOutput)
}

func (o ExportDeliveryDestinationResponseOutput) Container() pulumi.StringOutput {
	return o.ApplyT(func(v ExportDeliveryDestinationResponse) string { return v.Container }).(pulumi.StringOutput)
}

func (o ExportDeliveryDestinationResponseOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v ExportDeliveryDestinationResponse) string { return v.ResourceId }).(pulumi.StringOutput)
}

func (o ExportDeliveryDestinationResponseOutput) RootFolderPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportDeliveryDestinationResponse) *string { return v.RootFolderPath }).(pulumi.StringPtrOutput)
}

type ExportDeliveryDestinationResponsePtrOutput struct{ *pulumi.OutputState }

func (ExportDeliveryDestinationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportDeliveryDestinationResponse)(nil)).Elem()
}

func (o ExportDeliveryDestinationResponsePtrOutput) ToExportDeliveryDestinationResponsePtrOutput() ExportDeliveryDestinationResponsePtrOutput {
	return o
}

func (o ExportDeliveryDestinationResponsePtrOutput) ToExportDeliveryDestinationResponsePtrOutputWithContext(ctx context.Context) ExportDeliveryDestinationResponsePtrOutput {
	return o
}

func (o ExportDeliveryDestinationResponsePtrOutput) Elem() ExportDeliveryDestinationResponseOutput {
	return o.ApplyT(func(v *ExportDeliveryDestinationResponse) ExportDeliveryDestinationResponse {
		if v != nil {
			return *v
		}
		var ret ExportDeliveryDestinationResponse
		return ret
	}).(ExportDeliveryDestinationResponseOutput)
}

func (o ExportDeliveryDestinationResponsePtrOutput) Container() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportDeliveryDestinationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Container
	}).(pulumi.StringPtrOutput)
}

func (o ExportDeliveryDestinationResponsePtrOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportDeliveryDestinationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ResourceId
	}).(pulumi.StringPtrOutput)
}

func (o ExportDeliveryDestinationResponsePtrOutput) RootFolderPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportDeliveryDestinationResponse) *string {
		if v == nil {
			return nil
		}
		return v.RootFolderPath
	}).(pulumi.StringPtrOutput)
}

type ExportDeliveryInfo struct {
	Destination ExportDeliveryDestination `pulumi:"destination"`
}





type ExportDeliveryInfoInput interface {
	pulumi.Input

	ToExportDeliveryInfoOutput() ExportDeliveryInfoOutput
	ToExportDeliveryInfoOutputWithContext(context.Context) ExportDeliveryInfoOutput
}

type ExportDeliveryInfoArgs struct {
	Destination ExportDeliveryDestinationInput `pulumi:"destination"`
}

func (ExportDeliveryInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportDeliveryInfo)(nil)).Elem()
}

func (i ExportDeliveryInfoArgs) ToExportDeliveryInfoOutput() ExportDeliveryInfoOutput {
	return i.ToExportDeliveryInfoOutputWithContext(context.Background())
}

func (i ExportDeliveryInfoArgs) ToExportDeliveryInfoOutputWithContext(ctx context.Context) ExportDeliveryInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportDeliveryInfoOutput)
}

func (i ExportDeliveryInfoArgs) ToExportDeliveryInfoPtrOutput() ExportDeliveryInfoPtrOutput {
	return i.ToExportDeliveryInfoPtrOutputWithContext(context.Background())
}

func (i ExportDeliveryInfoArgs) ToExportDeliveryInfoPtrOutputWithContext(ctx context.Context) ExportDeliveryInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportDeliveryInfoOutput).ToExportDeliveryInfoPtrOutputWithContext(ctx)
}









type ExportDeliveryInfoPtrInput interface {
	pulumi.Input

	ToExportDeliveryInfoPtrOutput() ExportDeliveryInfoPtrOutput
	ToExportDeliveryInfoPtrOutputWithContext(context.Context) ExportDeliveryInfoPtrOutput
}

type exportDeliveryInfoPtrType ExportDeliveryInfoArgs

func ExportDeliveryInfoPtr(v *ExportDeliveryInfoArgs) ExportDeliveryInfoPtrInput {
	return (*exportDeliveryInfoPtrType)(v)
}

func (*exportDeliveryInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportDeliveryInfo)(nil)).Elem()
}

func (i *exportDeliveryInfoPtrType) ToExportDeliveryInfoPtrOutput() ExportDeliveryInfoPtrOutput {
	return i.ToExportDeliveryInfoPtrOutputWithContext(context.Background())
}

func (i *exportDeliveryInfoPtrType) ToExportDeliveryInfoPtrOutputWithContext(ctx context.Context) ExportDeliveryInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportDeliveryInfoPtrOutput)
}

type ExportDeliveryInfoOutput struct{ *pulumi.OutputState }

func (ExportDeliveryInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportDeliveryInfo)(nil)).Elem()
}

func (o ExportDeliveryInfoOutput) ToExportDeliveryInfoOutput() ExportDeliveryInfoOutput {
	return o
}

func (o ExportDeliveryInfoOutput) ToExportDeliveryInfoOutputWithContext(ctx context.Context) ExportDeliveryInfoOutput {
	return o
}

func (o ExportDeliveryInfoOutput) ToExportDeliveryInfoPtrOutput() ExportDeliveryInfoPtrOutput {
	return o.ToExportDeliveryInfoPtrOutputWithContext(context.Background())
}

func (o ExportDeliveryInfoOutput) ToExportDeliveryInfoPtrOutputWithContext(ctx context.Context) ExportDeliveryInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExportDeliveryInfo) *ExportDeliveryInfo {
		return &v
	}).(ExportDeliveryInfoPtrOutput)
}

func (o ExportDeliveryInfoOutput) Destination() ExportDeliveryDestinationOutput {
	return o.ApplyT(func(v ExportDeliveryInfo) ExportDeliveryDestination { return v.Destination }).(ExportDeliveryDestinationOutput)
}

type ExportDeliveryInfoPtrOutput struct{ *pulumi.OutputState }

func (ExportDeliveryInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportDeliveryInfo)(nil)).Elem()
}

func (o ExportDeliveryInfoPtrOutput) ToExportDeliveryInfoPtrOutput() ExportDeliveryInfoPtrOutput {
	return o
}

func (o ExportDeliveryInfoPtrOutput) ToExportDeliveryInfoPtrOutputWithContext(ctx context.Context) ExportDeliveryInfoPtrOutput {
	return o
}

func (o ExportDeliveryInfoPtrOutput) Elem() ExportDeliveryInfoOutput {
	return o.ApplyT(func(v *ExportDeliveryInfo) ExportDeliveryInfo {
		if v != nil {
			return *v
		}
		var ret ExportDeliveryInfo
		return ret
	}).(ExportDeliveryInfoOutput)
}

func (o ExportDeliveryInfoPtrOutput) Destination() ExportDeliveryDestinationPtrOutput {
	return o.ApplyT(func(v *ExportDeliveryInfo) *ExportDeliveryDestination {
		if v == nil {
			return nil
		}
		return &v.Destination
	}).(ExportDeliveryDestinationPtrOutput)
}

type ExportDeliveryInfoResponse struct {
	Destination ExportDeliveryDestinationResponse `pulumi:"destination"`
}





type ExportDeliveryInfoResponseInput interface {
	pulumi.Input

	ToExportDeliveryInfoResponseOutput() ExportDeliveryInfoResponseOutput
	ToExportDeliveryInfoResponseOutputWithContext(context.Context) ExportDeliveryInfoResponseOutput
}

type ExportDeliveryInfoResponseArgs struct {
	Destination ExportDeliveryDestinationResponseInput `pulumi:"destination"`
}

func (ExportDeliveryInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportDeliveryInfoResponse)(nil)).Elem()
}

func (i ExportDeliveryInfoResponseArgs) ToExportDeliveryInfoResponseOutput() ExportDeliveryInfoResponseOutput {
	return i.ToExportDeliveryInfoResponseOutputWithContext(context.Background())
}

func (i ExportDeliveryInfoResponseArgs) ToExportDeliveryInfoResponseOutputWithContext(ctx context.Context) ExportDeliveryInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportDeliveryInfoResponseOutput)
}

func (i ExportDeliveryInfoResponseArgs) ToExportDeliveryInfoResponsePtrOutput() ExportDeliveryInfoResponsePtrOutput {
	return i.ToExportDeliveryInfoResponsePtrOutputWithContext(context.Background())
}

func (i ExportDeliveryInfoResponseArgs) ToExportDeliveryInfoResponsePtrOutputWithContext(ctx context.Context) ExportDeliveryInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportDeliveryInfoResponseOutput).ToExportDeliveryInfoResponsePtrOutputWithContext(ctx)
}









type ExportDeliveryInfoResponsePtrInput interface {
	pulumi.Input

	ToExportDeliveryInfoResponsePtrOutput() ExportDeliveryInfoResponsePtrOutput
	ToExportDeliveryInfoResponsePtrOutputWithContext(context.Context) ExportDeliveryInfoResponsePtrOutput
}

type exportDeliveryInfoResponsePtrType ExportDeliveryInfoResponseArgs

func ExportDeliveryInfoResponsePtr(v *ExportDeliveryInfoResponseArgs) ExportDeliveryInfoResponsePtrInput {
	return (*exportDeliveryInfoResponsePtrType)(v)
}

func (*exportDeliveryInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportDeliveryInfoResponse)(nil)).Elem()
}

func (i *exportDeliveryInfoResponsePtrType) ToExportDeliveryInfoResponsePtrOutput() ExportDeliveryInfoResponsePtrOutput {
	return i.ToExportDeliveryInfoResponsePtrOutputWithContext(context.Background())
}

func (i *exportDeliveryInfoResponsePtrType) ToExportDeliveryInfoResponsePtrOutputWithContext(ctx context.Context) ExportDeliveryInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportDeliveryInfoResponsePtrOutput)
}

type ExportDeliveryInfoResponseOutput struct{ *pulumi.OutputState }

func (ExportDeliveryInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportDeliveryInfoResponse)(nil)).Elem()
}

func (o ExportDeliveryInfoResponseOutput) ToExportDeliveryInfoResponseOutput() ExportDeliveryInfoResponseOutput {
	return o
}

func (o ExportDeliveryInfoResponseOutput) ToExportDeliveryInfoResponseOutputWithContext(ctx context.Context) ExportDeliveryInfoResponseOutput {
	return o
}

func (o ExportDeliveryInfoResponseOutput) ToExportDeliveryInfoResponsePtrOutput() ExportDeliveryInfoResponsePtrOutput {
	return o.ToExportDeliveryInfoResponsePtrOutputWithContext(context.Background())
}

func (o ExportDeliveryInfoResponseOutput) ToExportDeliveryInfoResponsePtrOutputWithContext(ctx context.Context) ExportDeliveryInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExportDeliveryInfoResponse) *ExportDeliveryInfoResponse {
		return &v
	}).(ExportDeliveryInfoResponsePtrOutput)
}

func (o ExportDeliveryInfoResponseOutput) Destination() ExportDeliveryDestinationResponseOutput {
	return o.ApplyT(func(v ExportDeliveryInfoResponse) ExportDeliveryDestinationResponse { return v.Destination }).(ExportDeliveryDestinationResponseOutput)
}

type ExportDeliveryInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (ExportDeliveryInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportDeliveryInfoResponse)(nil)).Elem()
}

func (o ExportDeliveryInfoResponsePtrOutput) ToExportDeliveryInfoResponsePtrOutput() ExportDeliveryInfoResponsePtrOutput {
	return o
}

func (o ExportDeliveryInfoResponsePtrOutput) ToExportDeliveryInfoResponsePtrOutputWithContext(ctx context.Context) ExportDeliveryInfoResponsePtrOutput {
	return o
}

func (o ExportDeliveryInfoResponsePtrOutput) Elem() ExportDeliveryInfoResponseOutput {
	return o.ApplyT(func(v *ExportDeliveryInfoResponse) ExportDeliveryInfoResponse {
		if v != nil {
			return *v
		}
		var ret ExportDeliveryInfoResponse
		return ret
	}).(ExportDeliveryInfoResponseOutput)
}

func (o ExportDeliveryInfoResponsePtrOutput) Destination() ExportDeliveryDestinationResponsePtrOutput {
	return o.ApplyT(func(v *ExportDeliveryInfoResponse) *ExportDeliveryDestinationResponse {
		if v == nil {
			return nil
		}
		return &v.Destination
	}).(ExportDeliveryDestinationResponsePtrOutput)
}

type ExportRecurrencePeriod struct {
	From string  `pulumi:"from"`
	To   *string `pulumi:"to"`
}





type ExportRecurrencePeriodInput interface {
	pulumi.Input

	ToExportRecurrencePeriodOutput() ExportRecurrencePeriodOutput
	ToExportRecurrencePeriodOutputWithContext(context.Context) ExportRecurrencePeriodOutput
}

type ExportRecurrencePeriodArgs struct {
	From pulumi.StringInput    `pulumi:"from"`
	To   pulumi.StringPtrInput `pulumi:"to"`
}

func (ExportRecurrencePeriodArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportRecurrencePeriod)(nil)).Elem()
}

func (i ExportRecurrencePeriodArgs) ToExportRecurrencePeriodOutput() ExportRecurrencePeriodOutput {
	return i.ToExportRecurrencePeriodOutputWithContext(context.Background())
}

func (i ExportRecurrencePeriodArgs) ToExportRecurrencePeriodOutputWithContext(ctx context.Context) ExportRecurrencePeriodOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportRecurrencePeriodOutput)
}

func (i ExportRecurrencePeriodArgs) ToExportRecurrencePeriodPtrOutput() ExportRecurrencePeriodPtrOutput {
	return i.ToExportRecurrencePeriodPtrOutputWithContext(context.Background())
}

func (i ExportRecurrencePeriodArgs) ToExportRecurrencePeriodPtrOutputWithContext(ctx context.Context) ExportRecurrencePeriodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportRecurrencePeriodOutput).ToExportRecurrencePeriodPtrOutputWithContext(ctx)
}









type ExportRecurrencePeriodPtrInput interface {
	pulumi.Input

	ToExportRecurrencePeriodPtrOutput() ExportRecurrencePeriodPtrOutput
	ToExportRecurrencePeriodPtrOutputWithContext(context.Context) ExportRecurrencePeriodPtrOutput
}

type exportRecurrencePeriodPtrType ExportRecurrencePeriodArgs

func ExportRecurrencePeriodPtr(v *ExportRecurrencePeriodArgs) ExportRecurrencePeriodPtrInput {
	return (*exportRecurrencePeriodPtrType)(v)
}

func (*exportRecurrencePeriodPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportRecurrencePeriod)(nil)).Elem()
}

func (i *exportRecurrencePeriodPtrType) ToExportRecurrencePeriodPtrOutput() ExportRecurrencePeriodPtrOutput {
	return i.ToExportRecurrencePeriodPtrOutputWithContext(context.Background())
}

func (i *exportRecurrencePeriodPtrType) ToExportRecurrencePeriodPtrOutputWithContext(ctx context.Context) ExportRecurrencePeriodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportRecurrencePeriodPtrOutput)
}

type ExportRecurrencePeriodOutput struct{ *pulumi.OutputState }

func (ExportRecurrencePeriodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportRecurrencePeriod)(nil)).Elem()
}

func (o ExportRecurrencePeriodOutput) ToExportRecurrencePeriodOutput() ExportRecurrencePeriodOutput {
	return o
}

func (o ExportRecurrencePeriodOutput) ToExportRecurrencePeriodOutputWithContext(ctx context.Context) ExportRecurrencePeriodOutput {
	return o
}

func (o ExportRecurrencePeriodOutput) ToExportRecurrencePeriodPtrOutput() ExportRecurrencePeriodPtrOutput {
	return o.ToExportRecurrencePeriodPtrOutputWithContext(context.Background())
}

func (o ExportRecurrencePeriodOutput) ToExportRecurrencePeriodPtrOutputWithContext(ctx context.Context) ExportRecurrencePeriodPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExportRecurrencePeriod) *ExportRecurrencePeriod {
		return &v
	}).(ExportRecurrencePeriodPtrOutput)
}

func (o ExportRecurrencePeriodOutput) From() pulumi.StringOutput {
	return o.ApplyT(func(v ExportRecurrencePeriod) string { return v.From }).(pulumi.StringOutput)
}

func (o ExportRecurrencePeriodOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportRecurrencePeriod) *string { return v.To }).(pulumi.StringPtrOutput)
}

type ExportRecurrencePeriodPtrOutput struct{ *pulumi.OutputState }

func (ExportRecurrencePeriodPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportRecurrencePeriod)(nil)).Elem()
}

func (o ExportRecurrencePeriodPtrOutput) ToExportRecurrencePeriodPtrOutput() ExportRecurrencePeriodPtrOutput {
	return o
}

func (o ExportRecurrencePeriodPtrOutput) ToExportRecurrencePeriodPtrOutputWithContext(ctx context.Context) ExportRecurrencePeriodPtrOutput {
	return o
}

func (o ExportRecurrencePeriodPtrOutput) Elem() ExportRecurrencePeriodOutput {
	return o.ApplyT(func(v *ExportRecurrencePeriod) ExportRecurrencePeriod {
		if v != nil {
			return *v
		}
		var ret ExportRecurrencePeriod
		return ret
	}).(ExportRecurrencePeriodOutput)
}

func (o ExportRecurrencePeriodPtrOutput) From() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportRecurrencePeriod) *string {
		if v == nil {
			return nil
		}
		return &v.From
	}).(pulumi.StringPtrOutput)
}

func (o ExportRecurrencePeriodPtrOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportRecurrencePeriod) *string {
		if v == nil {
			return nil
		}
		return v.To
	}).(pulumi.StringPtrOutput)
}

type ExportRecurrencePeriodResponse struct {
	From string  `pulumi:"from"`
	To   *string `pulumi:"to"`
}





type ExportRecurrencePeriodResponseInput interface {
	pulumi.Input

	ToExportRecurrencePeriodResponseOutput() ExportRecurrencePeriodResponseOutput
	ToExportRecurrencePeriodResponseOutputWithContext(context.Context) ExportRecurrencePeriodResponseOutput
}

type ExportRecurrencePeriodResponseArgs struct {
	From pulumi.StringInput    `pulumi:"from"`
	To   pulumi.StringPtrInput `pulumi:"to"`
}

func (ExportRecurrencePeriodResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportRecurrencePeriodResponse)(nil)).Elem()
}

func (i ExportRecurrencePeriodResponseArgs) ToExportRecurrencePeriodResponseOutput() ExportRecurrencePeriodResponseOutput {
	return i.ToExportRecurrencePeriodResponseOutputWithContext(context.Background())
}

func (i ExportRecurrencePeriodResponseArgs) ToExportRecurrencePeriodResponseOutputWithContext(ctx context.Context) ExportRecurrencePeriodResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportRecurrencePeriodResponseOutput)
}

func (i ExportRecurrencePeriodResponseArgs) ToExportRecurrencePeriodResponsePtrOutput() ExportRecurrencePeriodResponsePtrOutput {
	return i.ToExportRecurrencePeriodResponsePtrOutputWithContext(context.Background())
}

func (i ExportRecurrencePeriodResponseArgs) ToExportRecurrencePeriodResponsePtrOutputWithContext(ctx context.Context) ExportRecurrencePeriodResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportRecurrencePeriodResponseOutput).ToExportRecurrencePeriodResponsePtrOutputWithContext(ctx)
}









type ExportRecurrencePeriodResponsePtrInput interface {
	pulumi.Input

	ToExportRecurrencePeriodResponsePtrOutput() ExportRecurrencePeriodResponsePtrOutput
	ToExportRecurrencePeriodResponsePtrOutputWithContext(context.Context) ExportRecurrencePeriodResponsePtrOutput
}

type exportRecurrencePeriodResponsePtrType ExportRecurrencePeriodResponseArgs

func ExportRecurrencePeriodResponsePtr(v *ExportRecurrencePeriodResponseArgs) ExportRecurrencePeriodResponsePtrInput {
	return (*exportRecurrencePeriodResponsePtrType)(v)
}

func (*exportRecurrencePeriodResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportRecurrencePeriodResponse)(nil)).Elem()
}

func (i *exportRecurrencePeriodResponsePtrType) ToExportRecurrencePeriodResponsePtrOutput() ExportRecurrencePeriodResponsePtrOutput {
	return i.ToExportRecurrencePeriodResponsePtrOutputWithContext(context.Background())
}

func (i *exportRecurrencePeriodResponsePtrType) ToExportRecurrencePeriodResponsePtrOutputWithContext(ctx context.Context) ExportRecurrencePeriodResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportRecurrencePeriodResponsePtrOutput)
}

type ExportRecurrencePeriodResponseOutput struct{ *pulumi.OutputState }

func (ExportRecurrencePeriodResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportRecurrencePeriodResponse)(nil)).Elem()
}

func (o ExportRecurrencePeriodResponseOutput) ToExportRecurrencePeriodResponseOutput() ExportRecurrencePeriodResponseOutput {
	return o
}

func (o ExportRecurrencePeriodResponseOutput) ToExportRecurrencePeriodResponseOutputWithContext(ctx context.Context) ExportRecurrencePeriodResponseOutput {
	return o
}

func (o ExportRecurrencePeriodResponseOutput) ToExportRecurrencePeriodResponsePtrOutput() ExportRecurrencePeriodResponsePtrOutput {
	return o.ToExportRecurrencePeriodResponsePtrOutputWithContext(context.Background())
}

func (o ExportRecurrencePeriodResponseOutput) ToExportRecurrencePeriodResponsePtrOutputWithContext(ctx context.Context) ExportRecurrencePeriodResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExportRecurrencePeriodResponse) *ExportRecurrencePeriodResponse {
		return &v
	}).(ExportRecurrencePeriodResponsePtrOutput)
}

func (o ExportRecurrencePeriodResponseOutput) From() pulumi.StringOutput {
	return o.ApplyT(func(v ExportRecurrencePeriodResponse) string { return v.From }).(pulumi.StringOutput)
}

func (o ExportRecurrencePeriodResponseOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportRecurrencePeriodResponse) *string { return v.To }).(pulumi.StringPtrOutput)
}

type ExportRecurrencePeriodResponsePtrOutput struct{ *pulumi.OutputState }

func (ExportRecurrencePeriodResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportRecurrencePeriodResponse)(nil)).Elem()
}

func (o ExportRecurrencePeriodResponsePtrOutput) ToExportRecurrencePeriodResponsePtrOutput() ExportRecurrencePeriodResponsePtrOutput {
	return o
}

func (o ExportRecurrencePeriodResponsePtrOutput) ToExportRecurrencePeriodResponsePtrOutputWithContext(ctx context.Context) ExportRecurrencePeriodResponsePtrOutput {
	return o
}

func (o ExportRecurrencePeriodResponsePtrOutput) Elem() ExportRecurrencePeriodResponseOutput {
	return o.ApplyT(func(v *ExportRecurrencePeriodResponse) ExportRecurrencePeriodResponse {
		if v != nil {
			return *v
		}
		var ret ExportRecurrencePeriodResponse
		return ret
	}).(ExportRecurrencePeriodResponseOutput)
}

func (o ExportRecurrencePeriodResponsePtrOutput) From() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportRecurrencePeriodResponse) *string {
		if v == nil {
			return nil
		}
		return &v.From
	}).(pulumi.StringPtrOutput)
}

func (o ExportRecurrencePeriodResponsePtrOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportRecurrencePeriodResponse) *string {
		if v == nil {
			return nil
		}
		return v.To
	}).(pulumi.StringPtrOutput)
}

type ExportSchedule struct {
	Recurrence       string                  `pulumi:"recurrence"`
	RecurrencePeriod *ExportRecurrencePeriod `pulumi:"recurrencePeriod"`
	Status           *string                 `pulumi:"status"`
}





type ExportScheduleInput interface {
	pulumi.Input

	ToExportScheduleOutput() ExportScheduleOutput
	ToExportScheduleOutputWithContext(context.Context) ExportScheduleOutput
}

type ExportScheduleArgs struct {
	Recurrence       pulumi.StringInput             `pulumi:"recurrence"`
	RecurrencePeriod ExportRecurrencePeriodPtrInput `pulumi:"recurrencePeriod"`
	Status           pulumi.StringPtrInput          `pulumi:"status"`
}

func (ExportScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportSchedule)(nil)).Elem()
}

func (i ExportScheduleArgs) ToExportScheduleOutput() ExportScheduleOutput {
	return i.ToExportScheduleOutputWithContext(context.Background())
}

func (i ExportScheduleArgs) ToExportScheduleOutputWithContext(ctx context.Context) ExportScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportScheduleOutput)
}

func (i ExportScheduleArgs) ToExportSchedulePtrOutput() ExportSchedulePtrOutput {
	return i.ToExportSchedulePtrOutputWithContext(context.Background())
}

func (i ExportScheduleArgs) ToExportSchedulePtrOutputWithContext(ctx context.Context) ExportSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportScheduleOutput).ToExportSchedulePtrOutputWithContext(ctx)
}









type ExportSchedulePtrInput interface {
	pulumi.Input

	ToExportSchedulePtrOutput() ExportSchedulePtrOutput
	ToExportSchedulePtrOutputWithContext(context.Context) ExportSchedulePtrOutput
}

type exportSchedulePtrType ExportScheduleArgs

func ExportSchedulePtr(v *ExportScheduleArgs) ExportSchedulePtrInput {
	return (*exportSchedulePtrType)(v)
}

func (*exportSchedulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportSchedule)(nil)).Elem()
}

func (i *exportSchedulePtrType) ToExportSchedulePtrOutput() ExportSchedulePtrOutput {
	return i.ToExportSchedulePtrOutputWithContext(context.Background())
}

func (i *exportSchedulePtrType) ToExportSchedulePtrOutputWithContext(ctx context.Context) ExportSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportSchedulePtrOutput)
}

type ExportScheduleOutput struct{ *pulumi.OutputState }

func (ExportScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportSchedule)(nil)).Elem()
}

func (o ExportScheduleOutput) ToExportScheduleOutput() ExportScheduleOutput {
	return o
}

func (o ExportScheduleOutput) ToExportScheduleOutputWithContext(ctx context.Context) ExportScheduleOutput {
	return o
}

func (o ExportScheduleOutput) ToExportSchedulePtrOutput() ExportSchedulePtrOutput {
	return o.ToExportSchedulePtrOutputWithContext(context.Background())
}

func (o ExportScheduleOutput) ToExportSchedulePtrOutputWithContext(ctx context.Context) ExportSchedulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExportSchedule) *ExportSchedule {
		return &v
	}).(ExportSchedulePtrOutput)
}

func (o ExportScheduleOutput) Recurrence() pulumi.StringOutput {
	return o.ApplyT(func(v ExportSchedule) string { return v.Recurrence }).(pulumi.StringOutput)
}

func (o ExportScheduleOutput) RecurrencePeriod() ExportRecurrencePeriodPtrOutput {
	return o.ApplyT(func(v ExportSchedule) *ExportRecurrencePeriod { return v.RecurrencePeriod }).(ExportRecurrencePeriodPtrOutput)
}

func (o ExportScheduleOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportSchedule) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type ExportSchedulePtrOutput struct{ *pulumi.OutputState }

func (ExportSchedulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportSchedule)(nil)).Elem()
}

func (o ExportSchedulePtrOutput) ToExportSchedulePtrOutput() ExportSchedulePtrOutput {
	return o
}

func (o ExportSchedulePtrOutput) ToExportSchedulePtrOutputWithContext(ctx context.Context) ExportSchedulePtrOutput {
	return o
}

func (o ExportSchedulePtrOutput) Elem() ExportScheduleOutput {
	return o.ApplyT(func(v *ExportSchedule) ExportSchedule {
		if v != nil {
			return *v
		}
		var ret ExportSchedule
		return ret
	}).(ExportScheduleOutput)
}

func (o ExportSchedulePtrOutput) Recurrence() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportSchedule) *string {
		if v == nil {
			return nil
		}
		return &v.Recurrence
	}).(pulumi.StringPtrOutput)
}

func (o ExportSchedulePtrOutput) RecurrencePeriod() ExportRecurrencePeriodPtrOutput {
	return o.ApplyT(func(v *ExportSchedule) *ExportRecurrencePeriod {
		if v == nil {
			return nil
		}
		return v.RecurrencePeriod
	}).(ExportRecurrencePeriodPtrOutput)
}

func (o ExportSchedulePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportSchedule) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type ExportScheduleResponse struct {
	Recurrence       string                          `pulumi:"recurrence"`
	RecurrencePeriod *ExportRecurrencePeriodResponse `pulumi:"recurrencePeriod"`
	Status           *string                         `pulumi:"status"`
}





type ExportScheduleResponseInput interface {
	pulumi.Input

	ToExportScheduleResponseOutput() ExportScheduleResponseOutput
	ToExportScheduleResponseOutputWithContext(context.Context) ExportScheduleResponseOutput
}

type ExportScheduleResponseArgs struct {
	Recurrence       pulumi.StringInput                     `pulumi:"recurrence"`
	RecurrencePeriod ExportRecurrencePeriodResponsePtrInput `pulumi:"recurrencePeriod"`
	Status           pulumi.StringPtrInput                  `pulumi:"status"`
}

func (ExportScheduleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportScheduleResponse)(nil)).Elem()
}

func (i ExportScheduleResponseArgs) ToExportScheduleResponseOutput() ExportScheduleResponseOutput {
	return i.ToExportScheduleResponseOutputWithContext(context.Background())
}

func (i ExportScheduleResponseArgs) ToExportScheduleResponseOutputWithContext(ctx context.Context) ExportScheduleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportScheduleResponseOutput)
}

func (i ExportScheduleResponseArgs) ToExportScheduleResponsePtrOutput() ExportScheduleResponsePtrOutput {
	return i.ToExportScheduleResponsePtrOutputWithContext(context.Background())
}

func (i ExportScheduleResponseArgs) ToExportScheduleResponsePtrOutputWithContext(ctx context.Context) ExportScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportScheduleResponseOutput).ToExportScheduleResponsePtrOutputWithContext(ctx)
}









type ExportScheduleResponsePtrInput interface {
	pulumi.Input

	ToExportScheduleResponsePtrOutput() ExportScheduleResponsePtrOutput
	ToExportScheduleResponsePtrOutputWithContext(context.Context) ExportScheduleResponsePtrOutput
}

type exportScheduleResponsePtrType ExportScheduleResponseArgs

func ExportScheduleResponsePtr(v *ExportScheduleResponseArgs) ExportScheduleResponsePtrInput {
	return (*exportScheduleResponsePtrType)(v)
}

func (*exportScheduleResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportScheduleResponse)(nil)).Elem()
}

func (i *exportScheduleResponsePtrType) ToExportScheduleResponsePtrOutput() ExportScheduleResponsePtrOutput {
	return i.ToExportScheduleResponsePtrOutputWithContext(context.Background())
}

func (i *exportScheduleResponsePtrType) ToExportScheduleResponsePtrOutputWithContext(ctx context.Context) ExportScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportScheduleResponsePtrOutput)
}

type ExportScheduleResponseOutput struct{ *pulumi.OutputState }

func (ExportScheduleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportScheduleResponse)(nil)).Elem()
}

func (o ExportScheduleResponseOutput) ToExportScheduleResponseOutput() ExportScheduleResponseOutput {
	return o
}

func (o ExportScheduleResponseOutput) ToExportScheduleResponseOutputWithContext(ctx context.Context) ExportScheduleResponseOutput {
	return o
}

func (o ExportScheduleResponseOutput) ToExportScheduleResponsePtrOutput() ExportScheduleResponsePtrOutput {
	return o.ToExportScheduleResponsePtrOutputWithContext(context.Background())
}

func (o ExportScheduleResponseOutput) ToExportScheduleResponsePtrOutputWithContext(ctx context.Context) ExportScheduleResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExportScheduleResponse) *ExportScheduleResponse {
		return &v
	}).(ExportScheduleResponsePtrOutput)
}

func (o ExportScheduleResponseOutput) Recurrence() pulumi.StringOutput {
	return o.ApplyT(func(v ExportScheduleResponse) string { return v.Recurrence }).(pulumi.StringOutput)
}

func (o ExportScheduleResponseOutput) RecurrencePeriod() ExportRecurrencePeriodResponsePtrOutput {
	return o.ApplyT(func(v ExportScheduleResponse) *ExportRecurrencePeriodResponse { return v.RecurrencePeriod }).(ExportRecurrencePeriodResponsePtrOutput)
}

func (o ExportScheduleResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportScheduleResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type ExportScheduleResponsePtrOutput struct{ *pulumi.OutputState }

func (ExportScheduleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportScheduleResponse)(nil)).Elem()
}

func (o ExportScheduleResponsePtrOutput) ToExportScheduleResponsePtrOutput() ExportScheduleResponsePtrOutput {
	return o
}

func (o ExportScheduleResponsePtrOutput) ToExportScheduleResponsePtrOutputWithContext(ctx context.Context) ExportScheduleResponsePtrOutput {
	return o
}

func (o ExportScheduleResponsePtrOutput) Elem() ExportScheduleResponseOutput {
	return o.ApplyT(func(v *ExportScheduleResponse) ExportScheduleResponse {
		if v != nil {
			return *v
		}
		var ret ExportScheduleResponse
		return ret
	}).(ExportScheduleResponseOutput)
}

func (o ExportScheduleResponsePtrOutput) Recurrence() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Recurrence
	}).(pulumi.StringPtrOutput)
}

func (o ExportScheduleResponsePtrOutput) RecurrencePeriod() ExportRecurrencePeriodResponsePtrOutput {
	return o.ApplyT(func(v *ExportScheduleResponse) *ExportRecurrencePeriodResponse {
		if v == nil {
			return nil
		}
		return v.RecurrencePeriod
	}).(ExportRecurrencePeriodResponsePtrOutput)
}

func (o ExportScheduleResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type QueryAggregation struct {
	Function string `pulumi:"function"`
	Name     string `pulumi:"name"`
}





type QueryAggregationInput interface {
	pulumi.Input

	ToQueryAggregationOutput() QueryAggregationOutput
	ToQueryAggregationOutputWithContext(context.Context) QueryAggregationOutput
}

type QueryAggregationArgs struct {
	Function pulumi.StringInput `pulumi:"function"`
	Name     pulumi.StringInput `pulumi:"name"`
}

func (QueryAggregationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryAggregation)(nil)).Elem()
}

func (i QueryAggregationArgs) ToQueryAggregationOutput() QueryAggregationOutput {
	return i.ToQueryAggregationOutputWithContext(context.Background())
}

func (i QueryAggregationArgs) ToQueryAggregationOutputWithContext(ctx context.Context) QueryAggregationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryAggregationOutput)
}





type QueryAggregationMapInput interface {
	pulumi.Input

	ToQueryAggregationMapOutput() QueryAggregationMapOutput
	ToQueryAggregationMapOutputWithContext(context.Context) QueryAggregationMapOutput
}

type QueryAggregationMap map[string]QueryAggregationInput

func (QueryAggregationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]QueryAggregation)(nil)).Elem()
}

func (i QueryAggregationMap) ToQueryAggregationMapOutput() QueryAggregationMapOutput {
	return i.ToQueryAggregationMapOutputWithContext(context.Background())
}

func (i QueryAggregationMap) ToQueryAggregationMapOutputWithContext(ctx context.Context) QueryAggregationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryAggregationMapOutput)
}

type QueryAggregationOutput struct{ *pulumi.OutputState }

func (QueryAggregationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryAggregation)(nil)).Elem()
}

func (o QueryAggregationOutput) ToQueryAggregationOutput() QueryAggregationOutput {
	return o
}

func (o QueryAggregationOutput) ToQueryAggregationOutputWithContext(ctx context.Context) QueryAggregationOutput {
	return o
}

func (o QueryAggregationOutput) Function() pulumi.StringOutput {
	return o.ApplyT(func(v QueryAggregation) string { return v.Function }).(pulumi.StringOutput)
}

func (o QueryAggregationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v QueryAggregation) string { return v.Name }).(pulumi.StringOutput)
}

type QueryAggregationMapOutput struct{ *pulumi.OutputState }

func (QueryAggregationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]QueryAggregation)(nil)).Elem()
}

func (o QueryAggregationMapOutput) ToQueryAggregationMapOutput() QueryAggregationMapOutput {
	return o
}

func (o QueryAggregationMapOutput) ToQueryAggregationMapOutputWithContext(ctx context.Context) QueryAggregationMapOutput {
	return o
}

func (o QueryAggregationMapOutput) MapIndex(k pulumi.StringInput) QueryAggregationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) QueryAggregation {
		return vs[0].(map[string]QueryAggregation)[vs[1].(string)]
	}).(QueryAggregationOutput)
}

type QueryAggregationResponse struct {
	Function string `pulumi:"function"`
	Name     string `pulumi:"name"`
}





type QueryAggregationResponseInput interface {
	pulumi.Input

	ToQueryAggregationResponseOutput() QueryAggregationResponseOutput
	ToQueryAggregationResponseOutputWithContext(context.Context) QueryAggregationResponseOutput
}

type QueryAggregationResponseArgs struct {
	Function pulumi.StringInput `pulumi:"function"`
	Name     pulumi.StringInput `pulumi:"name"`
}

func (QueryAggregationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryAggregationResponse)(nil)).Elem()
}

func (i QueryAggregationResponseArgs) ToQueryAggregationResponseOutput() QueryAggregationResponseOutput {
	return i.ToQueryAggregationResponseOutputWithContext(context.Background())
}

func (i QueryAggregationResponseArgs) ToQueryAggregationResponseOutputWithContext(ctx context.Context) QueryAggregationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryAggregationResponseOutput)
}





type QueryAggregationResponseMapInput interface {
	pulumi.Input

	ToQueryAggregationResponseMapOutput() QueryAggregationResponseMapOutput
	ToQueryAggregationResponseMapOutputWithContext(context.Context) QueryAggregationResponseMapOutput
}

type QueryAggregationResponseMap map[string]QueryAggregationResponseInput

func (QueryAggregationResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]QueryAggregationResponse)(nil)).Elem()
}

func (i QueryAggregationResponseMap) ToQueryAggregationResponseMapOutput() QueryAggregationResponseMapOutput {
	return i.ToQueryAggregationResponseMapOutputWithContext(context.Background())
}

func (i QueryAggregationResponseMap) ToQueryAggregationResponseMapOutputWithContext(ctx context.Context) QueryAggregationResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryAggregationResponseMapOutput)
}

type QueryAggregationResponseOutput struct{ *pulumi.OutputState }

func (QueryAggregationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryAggregationResponse)(nil)).Elem()
}

func (o QueryAggregationResponseOutput) ToQueryAggregationResponseOutput() QueryAggregationResponseOutput {
	return o
}

func (o QueryAggregationResponseOutput) ToQueryAggregationResponseOutputWithContext(ctx context.Context) QueryAggregationResponseOutput {
	return o
}

func (o QueryAggregationResponseOutput) Function() pulumi.StringOutput {
	return o.ApplyT(func(v QueryAggregationResponse) string { return v.Function }).(pulumi.StringOutput)
}

func (o QueryAggregationResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v QueryAggregationResponse) string { return v.Name }).(pulumi.StringOutput)
}

type QueryAggregationResponseMapOutput struct{ *pulumi.OutputState }

func (QueryAggregationResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]QueryAggregationResponse)(nil)).Elem()
}

func (o QueryAggregationResponseMapOutput) ToQueryAggregationResponseMapOutput() QueryAggregationResponseMapOutput {
	return o
}

func (o QueryAggregationResponseMapOutput) ToQueryAggregationResponseMapOutputWithContext(ctx context.Context) QueryAggregationResponseMapOutput {
	return o
}

func (o QueryAggregationResponseMapOutput) MapIndex(k pulumi.StringInput) QueryAggregationResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) QueryAggregationResponse {
		return vs[0].(map[string]QueryAggregationResponse)[vs[1].(string)]
	}).(QueryAggregationResponseOutput)
}

type QueryComparisonExpression struct {
	Name     string   `pulumi:"name"`
	Operator string   `pulumi:"operator"`
	Values   []string `pulumi:"values"`
}





type QueryComparisonExpressionInput interface {
	pulumi.Input

	ToQueryComparisonExpressionOutput() QueryComparisonExpressionOutput
	ToQueryComparisonExpressionOutputWithContext(context.Context) QueryComparisonExpressionOutput
}

type QueryComparisonExpressionArgs struct {
	Name     pulumi.StringInput      `pulumi:"name"`
	Operator pulumi.StringInput      `pulumi:"operator"`
	Values   pulumi.StringArrayInput `pulumi:"values"`
}

func (QueryComparisonExpressionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryComparisonExpression)(nil)).Elem()
}

func (i QueryComparisonExpressionArgs) ToQueryComparisonExpressionOutput() QueryComparisonExpressionOutput {
	return i.ToQueryComparisonExpressionOutputWithContext(context.Background())
}

func (i QueryComparisonExpressionArgs) ToQueryComparisonExpressionOutputWithContext(ctx context.Context) QueryComparisonExpressionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryComparisonExpressionOutput)
}

func (i QueryComparisonExpressionArgs) ToQueryComparisonExpressionPtrOutput() QueryComparisonExpressionPtrOutput {
	return i.ToQueryComparisonExpressionPtrOutputWithContext(context.Background())
}

func (i QueryComparisonExpressionArgs) ToQueryComparisonExpressionPtrOutputWithContext(ctx context.Context) QueryComparisonExpressionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryComparisonExpressionOutput).ToQueryComparisonExpressionPtrOutputWithContext(ctx)
}









type QueryComparisonExpressionPtrInput interface {
	pulumi.Input

	ToQueryComparisonExpressionPtrOutput() QueryComparisonExpressionPtrOutput
	ToQueryComparisonExpressionPtrOutputWithContext(context.Context) QueryComparisonExpressionPtrOutput
}

type queryComparisonExpressionPtrType QueryComparisonExpressionArgs

func QueryComparisonExpressionPtr(v *QueryComparisonExpressionArgs) QueryComparisonExpressionPtrInput {
	return (*queryComparisonExpressionPtrType)(v)
}

func (*queryComparisonExpressionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryComparisonExpression)(nil)).Elem()
}

func (i *queryComparisonExpressionPtrType) ToQueryComparisonExpressionPtrOutput() QueryComparisonExpressionPtrOutput {
	return i.ToQueryComparisonExpressionPtrOutputWithContext(context.Background())
}

func (i *queryComparisonExpressionPtrType) ToQueryComparisonExpressionPtrOutputWithContext(ctx context.Context) QueryComparisonExpressionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryComparisonExpressionPtrOutput)
}

type QueryComparisonExpressionOutput struct{ *pulumi.OutputState }

func (QueryComparisonExpressionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryComparisonExpression)(nil)).Elem()
}

func (o QueryComparisonExpressionOutput) ToQueryComparisonExpressionOutput() QueryComparisonExpressionOutput {
	return o
}

func (o QueryComparisonExpressionOutput) ToQueryComparisonExpressionOutputWithContext(ctx context.Context) QueryComparisonExpressionOutput {
	return o
}

func (o QueryComparisonExpressionOutput) ToQueryComparisonExpressionPtrOutput() QueryComparisonExpressionPtrOutput {
	return o.ToQueryComparisonExpressionPtrOutputWithContext(context.Background())
}

func (o QueryComparisonExpressionOutput) ToQueryComparisonExpressionPtrOutputWithContext(ctx context.Context) QueryComparisonExpressionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v QueryComparisonExpression) *QueryComparisonExpression {
		return &v
	}).(QueryComparisonExpressionPtrOutput)
}

func (o QueryComparisonExpressionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v QueryComparisonExpression) string { return v.Name }).(pulumi.StringOutput)
}

func (o QueryComparisonExpressionOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v QueryComparisonExpression) string { return v.Operator }).(pulumi.StringOutput)
}

func (o QueryComparisonExpressionOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v QueryComparisonExpression) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type QueryComparisonExpressionPtrOutput struct{ *pulumi.OutputState }

func (QueryComparisonExpressionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryComparisonExpression)(nil)).Elem()
}

func (o QueryComparisonExpressionPtrOutput) ToQueryComparisonExpressionPtrOutput() QueryComparisonExpressionPtrOutput {
	return o
}

func (o QueryComparisonExpressionPtrOutput) ToQueryComparisonExpressionPtrOutputWithContext(ctx context.Context) QueryComparisonExpressionPtrOutput {
	return o
}

func (o QueryComparisonExpressionPtrOutput) Elem() QueryComparisonExpressionOutput {
	return o.ApplyT(func(v *QueryComparisonExpression) QueryComparisonExpression {
		if v != nil {
			return *v
		}
		var ret QueryComparisonExpression
		return ret
	}).(QueryComparisonExpressionOutput)
}

func (o QueryComparisonExpressionPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueryComparisonExpression) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o QueryComparisonExpressionPtrOutput) Operator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueryComparisonExpression) *string {
		if v == nil {
			return nil
		}
		return &v.Operator
	}).(pulumi.StringPtrOutput)
}

func (o QueryComparisonExpressionPtrOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *QueryComparisonExpression) []string {
		if v == nil {
			return nil
		}
		return v.Values
	}).(pulumi.StringArrayOutput)
}

type QueryComparisonExpressionResponse struct {
	Name     string   `pulumi:"name"`
	Operator string   `pulumi:"operator"`
	Values   []string `pulumi:"values"`
}





type QueryComparisonExpressionResponseInput interface {
	pulumi.Input

	ToQueryComparisonExpressionResponseOutput() QueryComparisonExpressionResponseOutput
	ToQueryComparisonExpressionResponseOutputWithContext(context.Context) QueryComparisonExpressionResponseOutput
}

type QueryComparisonExpressionResponseArgs struct {
	Name     pulumi.StringInput      `pulumi:"name"`
	Operator pulumi.StringInput      `pulumi:"operator"`
	Values   pulumi.StringArrayInput `pulumi:"values"`
}

func (QueryComparisonExpressionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryComparisonExpressionResponse)(nil)).Elem()
}

func (i QueryComparisonExpressionResponseArgs) ToQueryComparisonExpressionResponseOutput() QueryComparisonExpressionResponseOutput {
	return i.ToQueryComparisonExpressionResponseOutputWithContext(context.Background())
}

func (i QueryComparisonExpressionResponseArgs) ToQueryComparisonExpressionResponseOutputWithContext(ctx context.Context) QueryComparisonExpressionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryComparisonExpressionResponseOutput)
}

func (i QueryComparisonExpressionResponseArgs) ToQueryComparisonExpressionResponsePtrOutput() QueryComparisonExpressionResponsePtrOutput {
	return i.ToQueryComparisonExpressionResponsePtrOutputWithContext(context.Background())
}

func (i QueryComparisonExpressionResponseArgs) ToQueryComparisonExpressionResponsePtrOutputWithContext(ctx context.Context) QueryComparisonExpressionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryComparisonExpressionResponseOutput).ToQueryComparisonExpressionResponsePtrOutputWithContext(ctx)
}









type QueryComparisonExpressionResponsePtrInput interface {
	pulumi.Input

	ToQueryComparisonExpressionResponsePtrOutput() QueryComparisonExpressionResponsePtrOutput
	ToQueryComparisonExpressionResponsePtrOutputWithContext(context.Context) QueryComparisonExpressionResponsePtrOutput
}

type queryComparisonExpressionResponsePtrType QueryComparisonExpressionResponseArgs

func QueryComparisonExpressionResponsePtr(v *QueryComparisonExpressionResponseArgs) QueryComparisonExpressionResponsePtrInput {
	return (*queryComparisonExpressionResponsePtrType)(v)
}

func (*queryComparisonExpressionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryComparisonExpressionResponse)(nil)).Elem()
}

func (i *queryComparisonExpressionResponsePtrType) ToQueryComparisonExpressionResponsePtrOutput() QueryComparisonExpressionResponsePtrOutput {
	return i.ToQueryComparisonExpressionResponsePtrOutputWithContext(context.Background())
}

func (i *queryComparisonExpressionResponsePtrType) ToQueryComparisonExpressionResponsePtrOutputWithContext(ctx context.Context) QueryComparisonExpressionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryComparisonExpressionResponsePtrOutput)
}

type QueryComparisonExpressionResponseOutput struct{ *pulumi.OutputState }

func (QueryComparisonExpressionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryComparisonExpressionResponse)(nil)).Elem()
}

func (o QueryComparisonExpressionResponseOutput) ToQueryComparisonExpressionResponseOutput() QueryComparisonExpressionResponseOutput {
	return o
}

func (o QueryComparisonExpressionResponseOutput) ToQueryComparisonExpressionResponseOutputWithContext(ctx context.Context) QueryComparisonExpressionResponseOutput {
	return o
}

func (o QueryComparisonExpressionResponseOutput) ToQueryComparisonExpressionResponsePtrOutput() QueryComparisonExpressionResponsePtrOutput {
	return o.ToQueryComparisonExpressionResponsePtrOutputWithContext(context.Background())
}

func (o QueryComparisonExpressionResponseOutput) ToQueryComparisonExpressionResponsePtrOutputWithContext(ctx context.Context) QueryComparisonExpressionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v QueryComparisonExpressionResponse) *QueryComparisonExpressionResponse {
		return &v
	}).(QueryComparisonExpressionResponsePtrOutput)
}

func (o QueryComparisonExpressionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v QueryComparisonExpressionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o QueryComparisonExpressionResponseOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v QueryComparisonExpressionResponse) string { return v.Operator }).(pulumi.StringOutput)
}

func (o QueryComparisonExpressionResponseOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v QueryComparisonExpressionResponse) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type QueryComparisonExpressionResponsePtrOutput struct{ *pulumi.OutputState }

func (QueryComparisonExpressionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryComparisonExpressionResponse)(nil)).Elem()
}

func (o QueryComparisonExpressionResponsePtrOutput) ToQueryComparisonExpressionResponsePtrOutput() QueryComparisonExpressionResponsePtrOutput {
	return o
}

func (o QueryComparisonExpressionResponsePtrOutput) ToQueryComparisonExpressionResponsePtrOutputWithContext(ctx context.Context) QueryComparisonExpressionResponsePtrOutput {
	return o
}

func (o QueryComparisonExpressionResponsePtrOutput) Elem() QueryComparisonExpressionResponseOutput {
	return o.ApplyT(func(v *QueryComparisonExpressionResponse) QueryComparisonExpressionResponse {
		if v != nil {
			return *v
		}
		var ret QueryComparisonExpressionResponse
		return ret
	}).(QueryComparisonExpressionResponseOutput)
}

func (o QueryComparisonExpressionResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueryComparisonExpressionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o QueryComparisonExpressionResponsePtrOutput) Operator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueryComparisonExpressionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Operator
	}).(pulumi.StringPtrOutput)
}

func (o QueryComparisonExpressionResponsePtrOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *QueryComparisonExpressionResponse) []string {
		if v == nil {
			return nil
		}
		return v.Values
	}).(pulumi.StringArrayOutput)
}

type QueryDataset struct {
	Aggregation   map[string]QueryAggregation `pulumi:"aggregation"`
	Configuration *QueryDatasetConfiguration  `pulumi:"configuration"`
	Filter        *QueryFilter                `pulumi:"filter"`
	Granularity   *string                     `pulumi:"granularity"`
	Grouping      []QueryGrouping             `pulumi:"grouping"`
	Sorting       []QuerySortingConfiguration `pulumi:"sorting"`
}





type QueryDatasetInput interface {
	pulumi.Input

	ToQueryDatasetOutput() QueryDatasetOutput
	ToQueryDatasetOutputWithContext(context.Context) QueryDatasetOutput
}

type QueryDatasetArgs struct {
	Aggregation   QueryAggregationMapInput            `pulumi:"aggregation"`
	Configuration QueryDatasetConfigurationPtrInput   `pulumi:"configuration"`
	Filter        QueryFilterPtrInput                 `pulumi:"filter"`
	Granularity   pulumi.StringPtrInput               `pulumi:"granularity"`
	Grouping      QueryGroupingArrayInput             `pulumi:"grouping"`
	Sorting       QuerySortingConfigurationArrayInput `pulumi:"sorting"`
}

func (QueryDatasetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryDataset)(nil)).Elem()
}

func (i QueryDatasetArgs) ToQueryDatasetOutput() QueryDatasetOutput {
	return i.ToQueryDatasetOutputWithContext(context.Background())
}

func (i QueryDatasetArgs) ToQueryDatasetOutputWithContext(ctx context.Context) QueryDatasetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryDatasetOutput)
}

func (i QueryDatasetArgs) ToQueryDatasetPtrOutput() QueryDatasetPtrOutput {
	return i.ToQueryDatasetPtrOutputWithContext(context.Background())
}

func (i QueryDatasetArgs) ToQueryDatasetPtrOutputWithContext(ctx context.Context) QueryDatasetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryDatasetOutput).ToQueryDatasetPtrOutputWithContext(ctx)
}









type QueryDatasetPtrInput interface {
	pulumi.Input

	ToQueryDatasetPtrOutput() QueryDatasetPtrOutput
	ToQueryDatasetPtrOutputWithContext(context.Context) QueryDatasetPtrOutput
}

type queryDatasetPtrType QueryDatasetArgs

func QueryDatasetPtr(v *QueryDatasetArgs) QueryDatasetPtrInput {
	return (*queryDatasetPtrType)(v)
}

func (*queryDatasetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryDataset)(nil)).Elem()
}

func (i *queryDatasetPtrType) ToQueryDatasetPtrOutput() QueryDatasetPtrOutput {
	return i.ToQueryDatasetPtrOutputWithContext(context.Background())
}

func (i *queryDatasetPtrType) ToQueryDatasetPtrOutputWithContext(ctx context.Context) QueryDatasetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryDatasetPtrOutput)
}

type QueryDatasetOutput struct{ *pulumi.OutputState }

func (QueryDatasetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryDataset)(nil)).Elem()
}

func (o QueryDatasetOutput) ToQueryDatasetOutput() QueryDatasetOutput {
	return o
}

func (o QueryDatasetOutput) ToQueryDatasetOutputWithContext(ctx context.Context) QueryDatasetOutput {
	return o
}

func (o QueryDatasetOutput) ToQueryDatasetPtrOutput() QueryDatasetPtrOutput {
	return o.ToQueryDatasetPtrOutputWithContext(context.Background())
}

func (o QueryDatasetOutput) ToQueryDatasetPtrOutputWithContext(ctx context.Context) QueryDatasetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v QueryDataset) *QueryDataset {
		return &v
	}).(QueryDatasetPtrOutput)
}

func (o QueryDatasetOutput) Aggregation() QueryAggregationMapOutput {
	return o.ApplyT(func(v QueryDataset) map[string]QueryAggregation { return v.Aggregation }).(QueryAggregationMapOutput)
}

func (o QueryDatasetOutput) Configuration() QueryDatasetConfigurationPtrOutput {
	return o.ApplyT(func(v QueryDataset) *QueryDatasetConfiguration { return v.Configuration }).(QueryDatasetConfigurationPtrOutput)
}

func (o QueryDatasetOutput) Filter() QueryFilterPtrOutput {
	return o.ApplyT(func(v QueryDataset) *QueryFilter { return v.Filter }).(QueryFilterPtrOutput)
}

func (o QueryDatasetOutput) Granularity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QueryDataset) *string { return v.Granularity }).(pulumi.StringPtrOutput)
}

func (o QueryDatasetOutput) Grouping() QueryGroupingArrayOutput {
	return o.ApplyT(func(v QueryDataset) []QueryGrouping { return v.Grouping }).(QueryGroupingArrayOutput)
}

func (o QueryDatasetOutput) Sorting() QuerySortingConfigurationArrayOutput {
	return o.ApplyT(func(v QueryDataset) []QuerySortingConfiguration { return v.Sorting }).(QuerySortingConfigurationArrayOutput)
}

type QueryDatasetPtrOutput struct{ *pulumi.OutputState }

func (QueryDatasetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryDataset)(nil)).Elem()
}

func (o QueryDatasetPtrOutput) ToQueryDatasetPtrOutput() QueryDatasetPtrOutput {
	return o
}

func (o QueryDatasetPtrOutput) ToQueryDatasetPtrOutputWithContext(ctx context.Context) QueryDatasetPtrOutput {
	return o
}

func (o QueryDatasetPtrOutput) Elem() QueryDatasetOutput {
	return o.ApplyT(func(v *QueryDataset) QueryDataset {
		if v != nil {
			return *v
		}
		var ret QueryDataset
		return ret
	}).(QueryDatasetOutput)
}

func (o QueryDatasetPtrOutput) Aggregation() QueryAggregationMapOutput {
	return o.ApplyT(func(v *QueryDataset) map[string]QueryAggregation {
		if v == nil {
			return nil
		}
		return v.Aggregation
	}).(QueryAggregationMapOutput)
}

func (o QueryDatasetPtrOutput) Configuration() QueryDatasetConfigurationPtrOutput {
	return o.ApplyT(func(v *QueryDataset) *QueryDatasetConfiguration {
		if v == nil {
			return nil
		}
		return v.Configuration
	}).(QueryDatasetConfigurationPtrOutput)
}

func (o QueryDatasetPtrOutput) Filter() QueryFilterPtrOutput {
	return o.ApplyT(func(v *QueryDataset) *QueryFilter {
		if v == nil {
			return nil
		}
		return v.Filter
	}).(QueryFilterPtrOutput)
}

func (o QueryDatasetPtrOutput) Granularity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueryDataset) *string {
		if v == nil {
			return nil
		}
		return v.Granularity
	}).(pulumi.StringPtrOutput)
}

func (o QueryDatasetPtrOutput) Grouping() QueryGroupingArrayOutput {
	return o.ApplyT(func(v *QueryDataset) []QueryGrouping {
		if v == nil {
			return nil
		}
		return v.Grouping
	}).(QueryGroupingArrayOutput)
}

func (o QueryDatasetPtrOutput) Sorting() QuerySortingConfigurationArrayOutput {
	return o.ApplyT(func(v *QueryDataset) []QuerySortingConfiguration {
		if v == nil {
			return nil
		}
		return v.Sorting
	}).(QuerySortingConfigurationArrayOutput)
}

type QueryDatasetConfiguration struct {
	Columns []string `pulumi:"columns"`
}





type QueryDatasetConfigurationInput interface {
	pulumi.Input

	ToQueryDatasetConfigurationOutput() QueryDatasetConfigurationOutput
	ToQueryDatasetConfigurationOutputWithContext(context.Context) QueryDatasetConfigurationOutput
}

type QueryDatasetConfigurationArgs struct {
	Columns pulumi.StringArrayInput `pulumi:"columns"`
}

func (QueryDatasetConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryDatasetConfiguration)(nil)).Elem()
}

func (i QueryDatasetConfigurationArgs) ToQueryDatasetConfigurationOutput() QueryDatasetConfigurationOutput {
	return i.ToQueryDatasetConfigurationOutputWithContext(context.Background())
}

func (i QueryDatasetConfigurationArgs) ToQueryDatasetConfigurationOutputWithContext(ctx context.Context) QueryDatasetConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryDatasetConfigurationOutput)
}

func (i QueryDatasetConfigurationArgs) ToQueryDatasetConfigurationPtrOutput() QueryDatasetConfigurationPtrOutput {
	return i.ToQueryDatasetConfigurationPtrOutputWithContext(context.Background())
}

func (i QueryDatasetConfigurationArgs) ToQueryDatasetConfigurationPtrOutputWithContext(ctx context.Context) QueryDatasetConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryDatasetConfigurationOutput).ToQueryDatasetConfigurationPtrOutputWithContext(ctx)
}









type QueryDatasetConfigurationPtrInput interface {
	pulumi.Input

	ToQueryDatasetConfigurationPtrOutput() QueryDatasetConfigurationPtrOutput
	ToQueryDatasetConfigurationPtrOutputWithContext(context.Context) QueryDatasetConfigurationPtrOutput
}

type queryDatasetConfigurationPtrType QueryDatasetConfigurationArgs

func QueryDatasetConfigurationPtr(v *QueryDatasetConfigurationArgs) QueryDatasetConfigurationPtrInput {
	return (*queryDatasetConfigurationPtrType)(v)
}

func (*queryDatasetConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryDatasetConfiguration)(nil)).Elem()
}

func (i *queryDatasetConfigurationPtrType) ToQueryDatasetConfigurationPtrOutput() QueryDatasetConfigurationPtrOutput {
	return i.ToQueryDatasetConfigurationPtrOutputWithContext(context.Background())
}

func (i *queryDatasetConfigurationPtrType) ToQueryDatasetConfigurationPtrOutputWithContext(ctx context.Context) QueryDatasetConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryDatasetConfigurationPtrOutput)
}

type QueryDatasetConfigurationOutput struct{ *pulumi.OutputState }

func (QueryDatasetConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryDatasetConfiguration)(nil)).Elem()
}

func (o QueryDatasetConfigurationOutput) ToQueryDatasetConfigurationOutput() QueryDatasetConfigurationOutput {
	return o
}

func (o QueryDatasetConfigurationOutput) ToQueryDatasetConfigurationOutputWithContext(ctx context.Context) QueryDatasetConfigurationOutput {
	return o
}

func (o QueryDatasetConfigurationOutput) ToQueryDatasetConfigurationPtrOutput() QueryDatasetConfigurationPtrOutput {
	return o.ToQueryDatasetConfigurationPtrOutputWithContext(context.Background())
}

func (o QueryDatasetConfigurationOutput) ToQueryDatasetConfigurationPtrOutputWithContext(ctx context.Context) QueryDatasetConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v QueryDatasetConfiguration) *QueryDatasetConfiguration {
		return &v
	}).(QueryDatasetConfigurationPtrOutput)
}

func (o QueryDatasetConfigurationOutput) Columns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v QueryDatasetConfiguration) []string { return v.Columns }).(pulumi.StringArrayOutput)
}

type QueryDatasetConfigurationPtrOutput struct{ *pulumi.OutputState }

func (QueryDatasetConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryDatasetConfiguration)(nil)).Elem()
}

func (o QueryDatasetConfigurationPtrOutput) ToQueryDatasetConfigurationPtrOutput() QueryDatasetConfigurationPtrOutput {
	return o
}

func (o QueryDatasetConfigurationPtrOutput) ToQueryDatasetConfigurationPtrOutputWithContext(ctx context.Context) QueryDatasetConfigurationPtrOutput {
	return o
}

func (o QueryDatasetConfigurationPtrOutput) Elem() QueryDatasetConfigurationOutput {
	return o.ApplyT(func(v *QueryDatasetConfiguration) QueryDatasetConfiguration {
		if v != nil {
			return *v
		}
		var ret QueryDatasetConfiguration
		return ret
	}).(QueryDatasetConfigurationOutput)
}

func (o QueryDatasetConfigurationPtrOutput) Columns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *QueryDatasetConfiguration) []string {
		if v == nil {
			return nil
		}
		return v.Columns
	}).(pulumi.StringArrayOutput)
}

type QueryDatasetConfigurationResponse struct {
	Columns []string `pulumi:"columns"`
}





type QueryDatasetConfigurationResponseInput interface {
	pulumi.Input

	ToQueryDatasetConfigurationResponseOutput() QueryDatasetConfigurationResponseOutput
	ToQueryDatasetConfigurationResponseOutputWithContext(context.Context) QueryDatasetConfigurationResponseOutput
}

type QueryDatasetConfigurationResponseArgs struct {
	Columns pulumi.StringArrayInput `pulumi:"columns"`
}

func (QueryDatasetConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryDatasetConfigurationResponse)(nil)).Elem()
}

func (i QueryDatasetConfigurationResponseArgs) ToQueryDatasetConfigurationResponseOutput() QueryDatasetConfigurationResponseOutput {
	return i.ToQueryDatasetConfigurationResponseOutputWithContext(context.Background())
}

func (i QueryDatasetConfigurationResponseArgs) ToQueryDatasetConfigurationResponseOutputWithContext(ctx context.Context) QueryDatasetConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryDatasetConfigurationResponseOutput)
}

func (i QueryDatasetConfigurationResponseArgs) ToQueryDatasetConfigurationResponsePtrOutput() QueryDatasetConfigurationResponsePtrOutput {
	return i.ToQueryDatasetConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i QueryDatasetConfigurationResponseArgs) ToQueryDatasetConfigurationResponsePtrOutputWithContext(ctx context.Context) QueryDatasetConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryDatasetConfigurationResponseOutput).ToQueryDatasetConfigurationResponsePtrOutputWithContext(ctx)
}









type QueryDatasetConfigurationResponsePtrInput interface {
	pulumi.Input

	ToQueryDatasetConfigurationResponsePtrOutput() QueryDatasetConfigurationResponsePtrOutput
	ToQueryDatasetConfigurationResponsePtrOutputWithContext(context.Context) QueryDatasetConfigurationResponsePtrOutput
}

type queryDatasetConfigurationResponsePtrType QueryDatasetConfigurationResponseArgs

func QueryDatasetConfigurationResponsePtr(v *QueryDatasetConfigurationResponseArgs) QueryDatasetConfigurationResponsePtrInput {
	return (*queryDatasetConfigurationResponsePtrType)(v)
}

func (*queryDatasetConfigurationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryDatasetConfigurationResponse)(nil)).Elem()
}

func (i *queryDatasetConfigurationResponsePtrType) ToQueryDatasetConfigurationResponsePtrOutput() QueryDatasetConfigurationResponsePtrOutput {
	return i.ToQueryDatasetConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i *queryDatasetConfigurationResponsePtrType) ToQueryDatasetConfigurationResponsePtrOutputWithContext(ctx context.Context) QueryDatasetConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryDatasetConfigurationResponsePtrOutput)
}

type QueryDatasetConfigurationResponseOutput struct{ *pulumi.OutputState }

func (QueryDatasetConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryDatasetConfigurationResponse)(nil)).Elem()
}

func (o QueryDatasetConfigurationResponseOutput) ToQueryDatasetConfigurationResponseOutput() QueryDatasetConfigurationResponseOutput {
	return o
}

func (o QueryDatasetConfigurationResponseOutput) ToQueryDatasetConfigurationResponseOutputWithContext(ctx context.Context) QueryDatasetConfigurationResponseOutput {
	return o
}

func (o QueryDatasetConfigurationResponseOutput) ToQueryDatasetConfigurationResponsePtrOutput() QueryDatasetConfigurationResponsePtrOutput {
	return o.ToQueryDatasetConfigurationResponsePtrOutputWithContext(context.Background())
}

func (o QueryDatasetConfigurationResponseOutput) ToQueryDatasetConfigurationResponsePtrOutputWithContext(ctx context.Context) QueryDatasetConfigurationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v QueryDatasetConfigurationResponse) *QueryDatasetConfigurationResponse {
		return &v
	}).(QueryDatasetConfigurationResponsePtrOutput)
}

func (o QueryDatasetConfigurationResponseOutput) Columns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v QueryDatasetConfigurationResponse) []string { return v.Columns }).(pulumi.StringArrayOutput)
}

type QueryDatasetConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (QueryDatasetConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryDatasetConfigurationResponse)(nil)).Elem()
}

func (o QueryDatasetConfigurationResponsePtrOutput) ToQueryDatasetConfigurationResponsePtrOutput() QueryDatasetConfigurationResponsePtrOutput {
	return o
}

func (o QueryDatasetConfigurationResponsePtrOutput) ToQueryDatasetConfigurationResponsePtrOutputWithContext(ctx context.Context) QueryDatasetConfigurationResponsePtrOutput {
	return o
}

func (o QueryDatasetConfigurationResponsePtrOutput) Elem() QueryDatasetConfigurationResponseOutput {
	return o.ApplyT(func(v *QueryDatasetConfigurationResponse) QueryDatasetConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret QueryDatasetConfigurationResponse
		return ret
	}).(QueryDatasetConfigurationResponseOutput)
}

func (o QueryDatasetConfigurationResponsePtrOutput) Columns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *QueryDatasetConfigurationResponse) []string {
		if v == nil {
			return nil
		}
		return v.Columns
	}).(pulumi.StringArrayOutput)
}

type QueryDatasetResponse struct {
	Aggregation   map[string]QueryAggregationResponse `pulumi:"aggregation"`
	Configuration *QueryDatasetConfigurationResponse  `pulumi:"configuration"`
	Filter        *QueryFilterResponse                `pulumi:"filter"`
	Granularity   *string                             `pulumi:"granularity"`
	Grouping      []QueryGroupingResponse             `pulumi:"grouping"`
	Sorting       []QuerySortingConfigurationResponse `pulumi:"sorting"`
}





type QueryDatasetResponseInput interface {
	pulumi.Input

	ToQueryDatasetResponseOutput() QueryDatasetResponseOutput
	ToQueryDatasetResponseOutputWithContext(context.Context) QueryDatasetResponseOutput
}

type QueryDatasetResponseArgs struct {
	Aggregation   QueryAggregationResponseMapInput            `pulumi:"aggregation"`
	Configuration QueryDatasetConfigurationResponsePtrInput   `pulumi:"configuration"`
	Filter        QueryFilterResponsePtrInput                 `pulumi:"filter"`
	Granularity   pulumi.StringPtrInput                       `pulumi:"granularity"`
	Grouping      QueryGroupingResponseArrayInput             `pulumi:"grouping"`
	Sorting       QuerySortingConfigurationResponseArrayInput `pulumi:"sorting"`
}

func (QueryDatasetResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryDatasetResponse)(nil)).Elem()
}

func (i QueryDatasetResponseArgs) ToQueryDatasetResponseOutput() QueryDatasetResponseOutput {
	return i.ToQueryDatasetResponseOutputWithContext(context.Background())
}

func (i QueryDatasetResponseArgs) ToQueryDatasetResponseOutputWithContext(ctx context.Context) QueryDatasetResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryDatasetResponseOutput)
}

func (i QueryDatasetResponseArgs) ToQueryDatasetResponsePtrOutput() QueryDatasetResponsePtrOutput {
	return i.ToQueryDatasetResponsePtrOutputWithContext(context.Background())
}

func (i QueryDatasetResponseArgs) ToQueryDatasetResponsePtrOutputWithContext(ctx context.Context) QueryDatasetResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryDatasetResponseOutput).ToQueryDatasetResponsePtrOutputWithContext(ctx)
}









type QueryDatasetResponsePtrInput interface {
	pulumi.Input

	ToQueryDatasetResponsePtrOutput() QueryDatasetResponsePtrOutput
	ToQueryDatasetResponsePtrOutputWithContext(context.Context) QueryDatasetResponsePtrOutput
}

type queryDatasetResponsePtrType QueryDatasetResponseArgs

func QueryDatasetResponsePtr(v *QueryDatasetResponseArgs) QueryDatasetResponsePtrInput {
	return (*queryDatasetResponsePtrType)(v)
}

func (*queryDatasetResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryDatasetResponse)(nil)).Elem()
}

func (i *queryDatasetResponsePtrType) ToQueryDatasetResponsePtrOutput() QueryDatasetResponsePtrOutput {
	return i.ToQueryDatasetResponsePtrOutputWithContext(context.Background())
}

func (i *queryDatasetResponsePtrType) ToQueryDatasetResponsePtrOutputWithContext(ctx context.Context) QueryDatasetResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryDatasetResponsePtrOutput)
}

type QueryDatasetResponseOutput struct{ *pulumi.OutputState }

func (QueryDatasetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryDatasetResponse)(nil)).Elem()
}

func (o QueryDatasetResponseOutput) ToQueryDatasetResponseOutput() QueryDatasetResponseOutput {
	return o
}

func (o QueryDatasetResponseOutput) ToQueryDatasetResponseOutputWithContext(ctx context.Context) QueryDatasetResponseOutput {
	return o
}

func (o QueryDatasetResponseOutput) ToQueryDatasetResponsePtrOutput() QueryDatasetResponsePtrOutput {
	return o.ToQueryDatasetResponsePtrOutputWithContext(context.Background())
}

func (o QueryDatasetResponseOutput) ToQueryDatasetResponsePtrOutputWithContext(ctx context.Context) QueryDatasetResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v QueryDatasetResponse) *QueryDatasetResponse {
		return &v
	}).(QueryDatasetResponsePtrOutput)
}

func (o QueryDatasetResponseOutput) Aggregation() QueryAggregationResponseMapOutput {
	return o.ApplyT(func(v QueryDatasetResponse) map[string]QueryAggregationResponse { return v.Aggregation }).(QueryAggregationResponseMapOutput)
}

func (o QueryDatasetResponseOutput) Configuration() QueryDatasetConfigurationResponsePtrOutput {
	return o.ApplyT(func(v QueryDatasetResponse) *QueryDatasetConfigurationResponse { return v.Configuration }).(QueryDatasetConfigurationResponsePtrOutput)
}

func (o QueryDatasetResponseOutput) Filter() QueryFilterResponsePtrOutput {
	return o.ApplyT(func(v QueryDatasetResponse) *QueryFilterResponse { return v.Filter }).(QueryFilterResponsePtrOutput)
}

func (o QueryDatasetResponseOutput) Granularity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QueryDatasetResponse) *string { return v.Granularity }).(pulumi.StringPtrOutput)
}

func (o QueryDatasetResponseOutput) Grouping() QueryGroupingResponseArrayOutput {
	return o.ApplyT(func(v QueryDatasetResponse) []QueryGroupingResponse { return v.Grouping }).(QueryGroupingResponseArrayOutput)
}

func (o QueryDatasetResponseOutput) Sorting() QuerySortingConfigurationResponseArrayOutput {
	return o.ApplyT(func(v QueryDatasetResponse) []QuerySortingConfigurationResponse { return v.Sorting }).(QuerySortingConfigurationResponseArrayOutput)
}

type QueryDatasetResponsePtrOutput struct{ *pulumi.OutputState }

func (QueryDatasetResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryDatasetResponse)(nil)).Elem()
}

func (o QueryDatasetResponsePtrOutput) ToQueryDatasetResponsePtrOutput() QueryDatasetResponsePtrOutput {
	return o
}

func (o QueryDatasetResponsePtrOutput) ToQueryDatasetResponsePtrOutputWithContext(ctx context.Context) QueryDatasetResponsePtrOutput {
	return o
}

func (o QueryDatasetResponsePtrOutput) Elem() QueryDatasetResponseOutput {
	return o.ApplyT(func(v *QueryDatasetResponse) QueryDatasetResponse {
		if v != nil {
			return *v
		}
		var ret QueryDatasetResponse
		return ret
	}).(QueryDatasetResponseOutput)
}

func (o QueryDatasetResponsePtrOutput) Aggregation() QueryAggregationResponseMapOutput {
	return o.ApplyT(func(v *QueryDatasetResponse) map[string]QueryAggregationResponse {
		if v == nil {
			return nil
		}
		return v.Aggregation
	}).(QueryAggregationResponseMapOutput)
}

func (o QueryDatasetResponsePtrOutput) Configuration() QueryDatasetConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *QueryDatasetResponse) *QueryDatasetConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.Configuration
	}).(QueryDatasetConfigurationResponsePtrOutput)
}

func (o QueryDatasetResponsePtrOutput) Filter() QueryFilterResponsePtrOutput {
	return o.ApplyT(func(v *QueryDatasetResponse) *QueryFilterResponse {
		if v == nil {
			return nil
		}
		return v.Filter
	}).(QueryFilterResponsePtrOutput)
}

func (o QueryDatasetResponsePtrOutput) Granularity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueryDatasetResponse) *string {
		if v == nil {
			return nil
		}
		return v.Granularity
	}).(pulumi.StringPtrOutput)
}

func (o QueryDatasetResponsePtrOutput) Grouping() QueryGroupingResponseArrayOutput {
	return o.ApplyT(func(v *QueryDatasetResponse) []QueryGroupingResponse {
		if v == nil {
			return nil
		}
		return v.Grouping
	}).(QueryGroupingResponseArrayOutput)
}

func (o QueryDatasetResponsePtrOutput) Sorting() QuerySortingConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *QueryDatasetResponse) []QuerySortingConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.Sorting
	}).(QuerySortingConfigurationResponseArrayOutput)
}

type QueryDefinition struct {
	Dataset    *QueryDataset    `pulumi:"dataset"`
	TimePeriod *QueryTimePeriod `pulumi:"timePeriod"`
	Timeframe  string           `pulumi:"timeframe"`
	Type       string           `pulumi:"type"`
}





type QueryDefinitionInput interface {
	pulumi.Input

	ToQueryDefinitionOutput() QueryDefinitionOutput
	ToQueryDefinitionOutputWithContext(context.Context) QueryDefinitionOutput
}

type QueryDefinitionArgs struct {
	Dataset    QueryDatasetPtrInput    `pulumi:"dataset"`
	TimePeriod QueryTimePeriodPtrInput `pulumi:"timePeriod"`
	Timeframe  pulumi.StringInput      `pulumi:"timeframe"`
	Type       pulumi.StringInput      `pulumi:"type"`
}

func (QueryDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryDefinition)(nil)).Elem()
}

func (i QueryDefinitionArgs) ToQueryDefinitionOutput() QueryDefinitionOutput {
	return i.ToQueryDefinitionOutputWithContext(context.Background())
}

func (i QueryDefinitionArgs) ToQueryDefinitionOutputWithContext(ctx context.Context) QueryDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryDefinitionOutput)
}

func (i QueryDefinitionArgs) ToQueryDefinitionPtrOutput() QueryDefinitionPtrOutput {
	return i.ToQueryDefinitionPtrOutputWithContext(context.Background())
}

func (i QueryDefinitionArgs) ToQueryDefinitionPtrOutputWithContext(ctx context.Context) QueryDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryDefinitionOutput).ToQueryDefinitionPtrOutputWithContext(ctx)
}









type QueryDefinitionPtrInput interface {
	pulumi.Input

	ToQueryDefinitionPtrOutput() QueryDefinitionPtrOutput
	ToQueryDefinitionPtrOutputWithContext(context.Context) QueryDefinitionPtrOutput
}

type queryDefinitionPtrType QueryDefinitionArgs

func QueryDefinitionPtr(v *QueryDefinitionArgs) QueryDefinitionPtrInput {
	return (*queryDefinitionPtrType)(v)
}

func (*queryDefinitionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryDefinition)(nil)).Elem()
}

func (i *queryDefinitionPtrType) ToQueryDefinitionPtrOutput() QueryDefinitionPtrOutput {
	return i.ToQueryDefinitionPtrOutputWithContext(context.Background())
}

func (i *queryDefinitionPtrType) ToQueryDefinitionPtrOutputWithContext(ctx context.Context) QueryDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryDefinitionPtrOutput)
}

type QueryDefinitionOutput struct{ *pulumi.OutputState }

func (QueryDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryDefinition)(nil)).Elem()
}

func (o QueryDefinitionOutput) ToQueryDefinitionOutput() QueryDefinitionOutput {
	return o
}

func (o QueryDefinitionOutput) ToQueryDefinitionOutputWithContext(ctx context.Context) QueryDefinitionOutput {
	return o
}

func (o QueryDefinitionOutput) ToQueryDefinitionPtrOutput() QueryDefinitionPtrOutput {
	return o.ToQueryDefinitionPtrOutputWithContext(context.Background())
}

func (o QueryDefinitionOutput) ToQueryDefinitionPtrOutputWithContext(ctx context.Context) QueryDefinitionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v QueryDefinition) *QueryDefinition {
		return &v
	}).(QueryDefinitionPtrOutput)
}

func (o QueryDefinitionOutput) Dataset() QueryDatasetPtrOutput {
	return o.ApplyT(func(v QueryDefinition) *QueryDataset { return v.Dataset }).(QueryDatasetPtrOutput)
}

func (o QueryDefinitionOutput) TimePeriod() QueryTimePeriodPtrOutput {
	return o.ApplyT(func(v QueryDefinition) *QueryTimePeriod { return v.TimePeriod }).(QueryTimePeriodPtrOutput)
}

func (o QueryDefinitionOutput) Timeframe() pulumi.StringOutput {
	return o.ApplyT(func(v QueryDefinition) string { return v.Timeframe }).(pulumi.StringOutput)
}

func (o QueryDefinitionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v QueryDefinition) string { return v.Type }).(pulumi.StringOutput)
}

type QueryDefinitionPtrOutput struct{ *pulumi.OutputState }

func (QueryDefinitionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryDefinition)(nil)).Elem()
}

func (o QueryDefinitionPtrOutput) ToQueryDefinitionPtrOutput() QueryDefinitionPtrOutput {
	return o
}

func (o QueryDefinitionPtrOutput) ToQueryDefinitionPtrOutputWithContext(ctx context.Context) QueryDefinitionPtrOutput {
	return o
}

func (o QueryDefinitionPtrOutput) Elem() QueryDefinitionOutput {
	return o.ApplyT(func(v *QueryDefinition) QueryDefinition {
		if v != nil {
			return *v
		}
		var ret QueryDefinition
		return ret
	}).(QueryDefinitionOutput)
}

func (o QueryDefinitionPtrOutput) Dataset() QueryDatasetPtrOutput {
	return o.ApplyT(func(v *QueryDefinition) *QueryDataset {
		if v == nil {
			return nil
		}
		return v.Dataset
	}).(QueryDatasetPtrOutput)
}

func (o QueryDefinitionPtrOutput) TimePeriod() QueryTimePeriodPtrOutput {
	return o.ApplyT(func(v *QueryDefinition) *QueryTimePeriod {
		if v == nil {
			return nil
		}
		return v.TimePeriod
	}).(QueryTimePeriodPtrOutput)
}

func (o QueryDefinitionPtrOutput) Timeframe() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueryDefinition) *string {
		if v == nil {
			return nil
		}
		return &v.Timeframe
	}).(pulumi.StringPtrOutput)
}

func (o QueryDefinitionPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueryDefinition) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type QueryDefinitionResponse struct {
	Dataset    *QueryDatasetResponse    `pulumi:"dataset"`
	TimePeriod *QueryTimePeriodResponse `pulumi:"timePeriod"`
	Timeframe  string                   `pulumi:"timeframe"`
	Type       string                   `pulumi:"type"`
}





type QueryDefinitionResponseInput interface {
	pulumi.Input

	ToQueryDefinitionResponseOutput() QueryDefinitionResponseOutput
	ToQueryDefinitionResponseOutputWithContext(context.Context) QueryDefinitionResponseOutput
}

type QueryDefinitionResponseArgs struct {
	Dataset    QueryDatasetResponsePtrInput    `pulumi:"dataset"`
	TimePeriod QueryTimePeriodResponsePtrInput `pulumi:"timePeriod"`
	Timeframe  pulumi.StringInput              `pulumi:"timeframe"`
	Type       pulumi.StringInput              `pulumi:"type"`
}

func (QueryDefinitionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryDefinitionResponse)(nil)).Elem()
}

func (i QueryDefinitionResponseArgs) ToQueryDefinitionResponseOutput() QueryDefinitionResponseOutput {
	return i.ToQueryDefinitionResponseOutputWithContext(context.Background())
}

func (i QueryDefinitionResponseArgs) ToQueryDefinitionResponseOutputWithContext(ctx context.Context) QueryDefinitionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryDefinitionResponseOutput)
}

func (i QueryDefinitionResponseArgs) ToQueryDefinitionResponsePtrOutput() QueryDefinitionResponsePtrOutput {
	return i.ToQueryDefinitionResponsePtrOutputWithContext(context.Background())
}

func (i QueryDefinitionResponseArgs) ToQueryDefinitionResponsePtrOutputWithContext(ctx context.Context) QueryDefinitionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryDefinitionResponseOutput).ToQueryDefinitionResponsePtrOutputWithContext(ctx)
}









type QueryDefinitionResponsePtrInput interface {
	pulumi.Input

	ToQueryDefinitionResponsePtrOutput() QueryDefinitionResponsePtrOutput
	ToQueryDefinitionResponsePtrOutputWithContext(context.Context) QueryDefinitionResponsePtrOutput
}

type queryDefinitionResponsePtrType QueryDefinitionResponseArgs

func QueryDefinitionResponsePtr(v *QueryDefinitionResponseArgs) QueryDefinitionResponsePtrInput {
	return (*queryDefinitionResponsePtrType)(v)
}

func (*queryDefinitionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryDefinitionResponse)(nil)).Elem()
}

func (i *queryDefinitionResponsePtrType) ToQueryDefinitionResponsePtrOutput() QueryDefinitionResponsePtrOutput {
	return i.ToQueryDefinitionResponsePtrOutputWithContext(context.Background())
}

func (i *queryDefinitionResponsePtrType) ToQueryDefinitionResponsePtrOutputWithContext(ctx context.Context) QueryDefinitionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryDefinitionResponsePtrOutput)
}

type QueryDefinitionResponseOutput struct{ *pulumi.OutputState }

func (QueryDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryDefinitionResponse)(nil)).Elem()
}

func (o QueryDefinitionResponseOutput) ToQueryDefinitionResponseOutput() QueryDefinitionResponseOutput {
	return o
}

func (o QueryDefinitionResponseOutput) ToQueryDefinitionResponseOutputWithContext(ctx context.Context) QueryDefinitionResponseOutput {
	return o
}

func (o QueryDefinitionResponseOutput) ToQueryDefinitionResponsePtrOutput() QueryDefinitionResponsePtrOutput {
	return o.ToQueryDefinitionResponsePtrOutputWithContext(context.Background())
}

func (o QueryDefinitionResponseOutput) ToQueryDefinitionResponsePtrOutputWithContext(ctx context.Context) QueryDefinitionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v QueryDefinitionResponse) *QueryDefinitionResponse {
		return &v
	}).(QueryDefinitionResponsePtrOutput)
}

func (o QueryDefinitionResponseOutput) Dataset() QueryDatasetResponsePtrOutput {
	return o.ApplyT(func(v QueryDefinitionResponse) *QueryDatasetResponse { return v.Dataset }).(QueryDatasetResponsePtrOutput)
}

func (o QueryDefinitionResponseOutput) TimePeriod() QueryTimePeriodResponsePtrOutput {
	return o.ApplyT(func(v QueryDefinitionResponse) *QueryTimePeriodResponse { return v.TimePeriod }).(QueryTimePeriodResponsePtrOutput)
}

func (o QueryDefinitionResponseOutput) Timeframe() pulumi.StringOutput {
	return o.ApplyT(func(v QueryDefinitionResponse) string { return v.Timeframe }).(pulumi.StringOutput)
}

func (o QueryDefinitionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v QueryDefinitionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type QueryDefinitionResponsePtrOutput struct{ *pulumi.OutputState }

func (QueryDefinitionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryDefinitionResponse)(nil)).Elem()
}

func (o QueryDefinitionResponsePtrOutput) ToQueryDefinitionResponsePtrOutput() QueryDefinitionResponsePtrOutput {
	return o
}

func (o QueryDefinitionResponsePtrOutput) ToQueryDefinitionResponsePtrOutputWithContext(ctx context.Context) QueryDefinitionResponsePtrOutput {
	return o
}

func (o QueryDefinitionResponsePtrOutput) Elem() QueryDefinitionResponseOutput {
	return o.ApplyT(func(v *QueryDefinitionResponse) QueryDefinitionResponse {
		if v != nil {
			return *v
		}
		var ret QueryDefinitionResponse
		return ret
	}).(QueryDefinitionResponseOutput)
}

func (o QueryDefinitionResponsePtrOutput) Dataset() QueryDatasetResponsePtrOutput {
	return o.ApplyT(func(v *QueryDefinitionResponse) *QueryDatasetResponse {
		if v == nil {
			return nil
		}
		return v.Dataset
	}).(QueryDatasetResponsePtrOutput)
}

func (o QueryDefinitionResponsePtrOutput) TimePeriod() QueryTimePeriodResponsePtrOutput {
	return o.ApplyT(func(v *QueryDefinitionResponse) *QueryTimePeriodResponse {
		if v == nil {
			return nil
		}
		return v.TimePeriod
	}).(QueryTimePeriodResponsePtrOutput)
}

func (o QueryDefinitionResponsePtrOutput) Timeframe() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueryDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Timeframe
	}).(pulumi.StringPtrOutput)
}

func (o QueryDefinitionResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueryDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type QueryFilter struct {
	And       []QueryFilter              `pulumi:"and"`
	Dimension *QueryComparisonExpression `pulumi:"dimension"`
	Not       *QueryFilter               `pulumi:"not"`
	Or        []QueryFilter              `pulumi:"or"`
	Tag       *QueryComparisonExpression `pulumi:"tag"`
}





type QueryFilterInput interface {
	pulumi.Input

	ToQueryFilterOutput() QueryFilterOutput
	ToQueryFilterOutputWithContext(context.Context) QueryFilterOutput
}

type QueryFilterArgs struct {
	And       QueryFilterArrayInput             `pulumi:"and"`
	Dimension QueryComparisonExpressionPtrInput `pulumi:"dimension"`
	Not       QueryFilterPtrInput               `pulumi:"not"`
	Or        QueryFilterArrayInput             `pulumi:"or"`
	Tag       QueryComparisonExpressionPtrInput `pulumi:"tag"`
}

func (QueryFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryFilter)(nil)).Elem()
}

func (i QueryFilterArgs) ToQueryFilterOutput() QueryFilterOutput {
	return i.ToQueryFilterOutputWithContext(context.Background())
}

func (i QueryFilterArgs) ToQueryFilterOutputWithContext(ctx context.Context) QueryFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryFilterOutput)
}

func (i QueryFilterArgs) ToQueryFilterPtrOutput() QueryFilterPtrOutput {
	return i.ToQueryFilterPtrOutputWithContext(context.Background())
}

func (i QueryFilterArgs) ToQueryFilterPtrOutputWithContext(ctx context.Context) QueryFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryFilterOutput).ToQueryFilterPtrOutputWithContext(ctx)
}









type QueryFilterPtrInput interface {
	pulumi.Input

	ToQueryFilterPtrOutput() QueryFilterPtrOutput
	ToQueryFilterPtrOutputWithContext(context.Context) QueryFilterPtrOutput
}

type queryFilterPtrType QueryFilterArgs

func QueryFilterPtr(v *QueryFilterArgs) QueryFilterPtrInput {
	return (*queryFilterPtrType)(v)
}

func (*queryFilterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryFilter)(nil)).Elem()
}

func (i *queryFilterPtrType) ToQueryFilterPtrOutput() QueryFilterPtrOutput {
	return i.ToQueryFilterPtrOutputWithContext(context.Background())
}

func (i *queryFilterPtrType) ToQueryFilterPtrOutputWithContext(ctx context.Context) QueryFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryFilterPtrOutput)
}





type QueryFilterArrayInput interface {
	pulumi.Input

	ToQueryFilterArrayOutput() QueryFilterArrayOutput
	ToQueryFilterArrayOutputWithContext(context.Context) QueryFilterArrayOutput
}

type QueryFilterArray []QueryFilterInput

func (QueryFilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]QueryFilter)(nil)).Elem()
}

func (i QueryFilterArray) ToQueryFilterArrayOutput() QueryFilterArrayOutput {
	return i.ToQueryFilterArrayOutputWithContext(context.Background())
}

func (i QueryFilterArray) ToQueryFilterArrayOutputWithContext(ctx context.Context) QueryFilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryFilterArrayOutput)
}

type QueryFilterOutput struct{ *pulumi.OutputState }

func (QueryFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryFilter)(nil)).Elem()
}

func (o QueryFilterOutput) ToQueryFilterOutput() QueryFilterOutput {
	return o
}

func (o QueryFilterOutput) ToQueryFilterOutputWithContext(ctx context.Context) QueryFilterOutput {
	return o
}

func (o QueryFilterOutput) ToQueryFilterPtrOutput() QueryFilterPtrOutput {
	return o.ToQueryFilterPtrOutputWithContext(context.Background())
}

func (o QueryFilterOutput) ToQueryFilterPtrOutputWithContext(ctx context.Context) QueryFilterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v QueryFilter) *QueryFilter {
		return &v
	}).(QueryFilterPtrOutput)
}

func (o QueryFilterOutput) And() QueryFilterArrayOutput {
	return o.ApplyT(func(v QueryFilter) []QueryFilter { return v.And }).(QueryFilterArrayOutput)
}

func (o QueryFilterOutput) Dimension() QueryComparisonExpressionPtrOutput {
	return o.ApplyT(func(v QueryFilter) *QueryComparisonExpression { return v.Dimension }).(QueryComparisonExpressionPtrOutput)
}

func (o QueryFilterOutput) Not() QueryFilterPtrOutput {
	return o.ApplyT(func(v QueryFilter) *QueryFilter { return v.Not }).(QueryFilterPtrOutput)
}

func (o QueryFilterOutput) Or() QueryFilterArrayOutput {
	return o.ApplyT(func(v QueryFilter) []QueryFilter { return v.Or }).(QueryFilterArrayOutput)
}

func (o QueryFilterOutput) Tag() QueryComparisonExpressionPtrOutput {
	return o.ApplyT(func(v QueryFilter) *QueryComparisonExpression { return v.Tag }).(QueryComparisonExpressionPtrOutput)
}

type QueryFilterPtrOutput struct{ *pulumi.OutputState }

func (QueryFilterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryFilter)(nil)).Elem()
}

func (o QueryFilterPtrOutput) ToQueryFilterPtrOutput() QueryFilterPtrOutput {
	return o
}

func (o QueryFilterPtrOutput) ToQueryFilterPtrOutputWithContext(ctx context.Context) QueryFilterPtrOutput {
	return o
}

func (o QueryFilterPtrOutput) Elem() QueryFilterOutput {
	return o.ApplyT(func(v *QueryFilter) QueryFilter {
		if v != nil {
			return *v
		}
		var ret QueryFilter
		return ret
	}).(QueryFilterOutput)
}

func (o QueryFilterPtrOutput) And() QueryFilterArrayOutput {
	return o.ApplyT(func(v *QueryFilter) []QueryFilter {
		if v == nil {
			return nil
		}
		return v.And
	}).(QueryFilterArrayOutput)
}

func (o QueryFilterPtrOutput) Dimension() QueryComparisonExpressionPtrOutput {
	return o.ApplyT(func(v *QueryFilter) *QueryComparisonExpression {
		if v == nil {
			return nil
		}
		return v.Dimension
	}).(QueryComparisonExpressionPtrOutput)
}

func (o QueryFilterPtrOutput) Not() QueryFilterPtrOutput {
	return o.ApplyT(func(v *QueryFilter) *QueryFilter {
		if v == nil {
			return nil
		}
		return v.Not
	}).(QueryFilterPtrOutput)
}

func (o QueryFilterPtrOutput) Or() QueryFilterArrayOutput {
	return o.ApplyT(func(v *QueryFilter) []QueryFilter {
		if v == nil {
			return nil
		}
		return v.Or
	}).(QueryFilterArrayOutput)
}

func (o QueryFilterPtrOutput) Tag() QueryComparisonExpressionPtrOutput {
	return o.ApplyT(func(v *QueryFilter) *QueryComparisonExpression {
		if v == nil {
			return nil
		}
		return v.Tag
	}).(QueryComparisonExpressionPtrOutput)
}

type QueryFilterArrayOutput struct{ *pulumi.OutputState }

func (QueryFilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]QueryFilter)(nil)).Elem()
}

func (o QueryFilterArrayOutput) ToQueryFilterArrayOutput() QueryFilterArrayOutput {
	return o
}

func (o QueryFilterArrayOutput) ToQueryFilterArrayOutputWithContext(ctx context.Context) QueryFilterArrayOutput {
	return o
}

func (o QueryFilterArrayOutput) Index(i pulumi.IntInput) QueryFilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) QueryFilter {
		return vs[0].([]QueryFilter)[vs[1].(int)]
	}).(QueryFilterOutput)
}

type QueryFilterResponse struct {
	And       []QueryFilterResponse              `pulumi:"and"`
	Dimension *QueryComparisonExpressionResponse `pulumi:"dimension"`
	Not       *QueryFilterResponse               `pulumi:"not"`
	Or        []QueryFilterResponse              `pulumi:"or"`
	Tag       *QueryComparisonExpressionResponse `pulumi:"tag"`
}





type QueryFilterResponseInput interface {
	pulumi.Input

	ToQueryFilterResponseOutput() QueryFilterResponseOutput
	ToQueryFilterResponseOutputWithContext(context.Context) QueryFilterResponseOutput
}

type QueryFilterResponseArgs struct {
	And       QueryFilterResponseArrayInput             `pulumi:"and"`
	Dimension QueryComparisonExpressionResponsePtrInput `pulumi:"dimension"`
	Not       QueryFilterResponsePtrInput               `pulumi:"not"`
	Or        QueryFilterResponseArrayInput             `pulumi:"or"`
	Tag       QueryComparisonExpressionResponsePtrInput `pulumi:"tag"`
}

func (QueryFilterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryFilterResponse)(nil)).Elem()
}

func (i QueryFilterResponseArgs) ToQueryFilterResponseOutput() QueryFilterResponseOutput {
	return i.ToQueryFilterResponseOutputWithContext(context.Background())
}

func (i QueryFilterResponseArgs) ToQueryFilterResponseOutputWithContext(ctx context.Context) QueryFilterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryFilterResponseOutput)
}

func (i QueryFilterResponseArgs) ToQueryFilterResponsePtrOutput() QueryFilterResponsePtrOutput {
	return i.ToQueryFilterResponsePtrOutputWithContext(context.Background())
}

func (i QueryFilterResponseArgs) ToQueryFilterResponsePtrOutputWithContext(ctx context.Context) QueryFilterResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryFilterResponseOutput).ToQueryFilterResponsePtrOutputWithContext(ctx)
}









type QueryFilterResponsePtrInput interface {
	pulumi.Input

	ToQueryFilterResponsePtrOutput() QueryFilterResponsePtrOutput
	ToQueryFilterResponsePtrOutputWithContext(context.Context) QueryFilterResponsePtrOutput
}

type queryFilterResponsePtrType QueryFilterResponseArgs

func QueryFilterResponsePtr(v *QueryFilterResponseArgs) QueryFilterResponsePtrInput {
	return (*queryFilterResponsePtrType)(v)
}

func (*queryFilterResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryFilterResponse)(nil)).Elem()
}

func (i *queryFilterResponsePtrType) ToQueryFilterResponsePtrOutput() QueryFilterResponsePtrOutput {
	return i.ToQueryFilterResponsePtrOutputWithContext(context.Background())
}

func (i *queryFilterResponsePtrType) ToQueryFilterResponsePtrOutputWithContext(ctx context.Context) QueryFilterResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryFilterResponsePtrOutput)
}





type QueryFilterResponseArrayInput interface {
	pulumi.Input

	ToQueryFilterResponseArrayOutput() QueryFilterResponseArrayOutput
	ToQueryFilterResponseArrayOutputWithContext(context.Context) QueryFilterResponseArrayOutput
}

type QueryFilterResponseArray []QueryFilterResponseInput

func (QueryFilterResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]QueryFilterResponse)(nil)).Elem()
}

func (i QueryFilterResponseArray) ToQueryFilterResponseArrayOutput() QueryFilterResponseArrayOutput {
	return i.ToQueryFilterResponseArrayOutputWithContext(context.Background())
}

func (i QueryFilterResponseArray) ToQueryFilterResponseArrayOutputWithContext(ctx context.Context) QueryFilterResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryFilterResponseArrayOutput)
}

type QueryFilterResponseOutput struct{ *pulumi.OutputState }

func (QueryFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryFilterResponse)(nil)).Elem()
}

func (o QueryFilterResponseOutput) ToQueryFilterResponseOutput() QueryFilterResponseOutput {
	return o
}

func (o QueryFilterResponseOutput) ToQueryFilterResponseOutputWithContext(ctx context.Context) QueryFilterResponseOutput {
	return o
}

func (o QueryFilterResponseOutput) ToQueryFilterResponsePtrOutput() QueryFilterResponsePtrOutput {
	return o.ToQueryFilterResponsePtrOutputWithContext(context.Background())
}

func (o QueryFilterResponseOutput) ToQueryFilterResponsePtrOutputWithContext(ctx context.Context) QueryFilterResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v QueryFilterResponse) *QueryFilterResponse {
		return &v
	}).(QueryFilterResponsePtrOutput)
}

func (o QueryFilterResponseOutput) And() QueryFilterResponseArrayOutput {
	return o.ApplyT(func(v QueryFilterResponse) []QueryFilterResponse { return v.And }).(QueryFilterResponseArrayOutput)
}

func (o QueryFilterResponseOutput) Dimension() QueryComparisonExpressionResponsePtrOutput {
	return o.ApplyT(func(v QueryFilterResponse) *QueryComparisonExpressionResponse { return v.Dimension }).(QueryComparisonExpressionResponsePtrOutput)
}

func (o QueryFilterResponseOutput) Not() QueryFilterResponsePtrOutput {
	return o.ApplyT(func(v QueryFilterResponse) *QueryFilterResponse { return v.Not }).(QueryFilterResponsePtrOutput)
}

func (o QueryFilterResponseOutput) Or() QueryFilterResponseArrayOutput {
	return o.ApplyT(func(v QueryFilterResponse) []QueryFilterResponse { return v.Or }).(QueryFilterResponseArrayOutput)
}

func (o QueryFilterResponseOutput) Tag() QueryComparisonExpressionResponsePtrOutput {
	return o.ApplyT(func(v QueryFilterResponse) *QueryComparisonExpressionResponse { return v.Tag }).(QueryComparisonExpressionResponsePtrOutput)
}

type QueryFilterResponsePtrOutput struct{ *pulumi.OutputState }

func (QueryFilterResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryFilterResponse)(nil)).Elem()
}

func (o QueryFilterResponsePtrOutput) ToQueryFilterResponsePtrOutput() QueryFilterResponsePtrOutput {
	return o
}

func (o QueryFilterResponsePtrOutput) ToQueryFilterResponsePtrOutputWithContext(ctx context.Context) QueryFilterResponsePtrOutput {
	return o
}

func (o QueryFilterResponsePtrOutput) Elem() QueryFilterResponseOutput {
	return o.ApplyT(func(v *QueryFilterResponse) QueryFilterResponse {
		if v != nil {
			return *v
		}
		var ret QueryFilterResponse
		return ret
	}).(QueryFilterResponseOutput)
}

func (o QueryFilterResponsePtrOutput) And() QueryFilterResponseArrayOutput {
	return o.ApplyT(func(v *QueryFilterResponse) []QueryFilterResponse {
		if v == nil {
			return nil
		}
		return v.And
	}).(QueryFilterResponseArrayOutput)
}

func (o QueryFilterResponsePtrOutput) Dimension() QueryComparisonExpressionResponsePtrOutput {
	return o.ApplyT(func(v *QueryFilterResponse) *QueryComparisonExpressionResponse {
		if v == nil {
			return nil
		}
		return v.Dimension
	}).(QueryComparisonExpressionResponsePtrOutput)
}

func (o QueryFilterResponsePtrOutput) Not() QueryFilterResponsePtrOutput {
	return o.ApplyT(func(v *QueryFilterResponse) *QueryFilterResponse {
		if v == nil {
			return nil
		}
		return v.Not
	}).(QueryFilterResponsePtrOutput)
}

func (o QueryFilterResponsePtrOutput) Or() QueryFilterResponseArrayOutput {
	return o.ApplyT(func(v *QueryFilterResponse) []QueryFilterResponse {
		if v == nil {
			return nil
		}
		return v.Or
	}).(QueryFilterResponseArrayOutput)
}

func (o QueryFilterResponsePtrOutput) Tag() QueryComparisonExpressionResponsePtrOutput {
	return o.ApplyT(func(v *QueryFilterResponse) *QueryComparisonExpressionResponse {
		if v == nil {
			return nil
		}
		return v.Tag
	}).(QueryComparisonExpressionResponsePtrOutput)
}

type QueryFilterResponseArrayOutput struct{ *pulumi.OutputState }

func (QueryFilterResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]QueryFilterResponse)(nil)).Elem()
}

func (o QueryFilterResponseArrayOutput) ToQueryFilterResponseArrayOutput() QueryFilterResponseArrayOutput {
	return o
}

func (o QueryFilterResponseArrayOutput) ToQueryFilterResponseArrayOutputWithContext(ctx context.Context) QueryFilterResponseArrayOutput {
	return o
}

func (o QueryFilterResponseArrayOutput) Index(i pulumi.IntInput) QueryFilterResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) QueryFilterResponse {
		return vs[0].([]QueryFilterResponse)[vs[1].(int)]
	}).(QueryFilterResponseOutput)
}

type QueryGrouping struct {
	Name string `pulumi:"name"`
	Type string `pulumi:"type"`
}





type QueryGroupingInput interface {
	pulumi.Input

	ToQueryGroupingOutput() QueryGroupingOutput
	ToQueryGroupingOutputWithContext(context.Context) QueryGroupingOutput
}

type QueryGroupingArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
	Type pulumi.StringInput `pulumi:"type"`
}

func (QueryGroupingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryGrouping)(nil)).Elem()
}

func (i QueryGroupingArgs) ToQueryGroupingOutput() QueryGroupingOutput {
	return i.ToQueryGroupingOutputWithContext(context.Background())
}

func (i QueryGroupingArgs) ToQueryGroupingOutputWithContext(ctx context.Context) QueryGroupingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryGroupingOutput)
}





type QueryGroupingArrayInput interface {
	pulumi.Input

	ToQueryGroupingArrayOutput() QueryGroupingArrayOutput
	ToQueryGroupingArrayOutputWithContext(context.Context) QueryGroupingArrayOutput
}

type QueryGroupingArray []QueryGroupingInput

func (QueryGroupingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]QueryGrouping)(nil)).Elem()
}

func (i QueryGroupingArray) ToQueryGroupingArrayOutput() QueryGroupingArrayOutput {
	return i.ToQueryGroupingArrayOutputWithContext(context.Background())
}

func (i QueryGroupingArray) ToQueryGroupingArrayOutputWithContext(ctx context.Context) QueryGroupingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryGroupingArrayOutput)
}

type QueryGroupingOutput struct{ *pulumi.OutputState }

func (QueryGroupingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryGrouping)(nil)).Elem()
}

func (o QueryGroupingOutput) ToQueryGroupingOutput() QueryGroupingOutput {
	return o
}

func (o QueryGroupingOutput) ToQueryGroupingOutputWithContext(ctx context.Context) QueryGroupingOutput {
	return o
}

func (o QueryGroupingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v QueryGrouping) string { return v.Name }).(pulumi.StringOutput)
}

func (o QueryGroupingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v QueryGrouping) string { return v.Type }).(pulumi.StringOutput)
}

type QueryGroupingArrayOutput struct{ *pulumi.OutputState }

func (QueryGroupingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]QueryGrouping)(nil)).Elem()
}

func (o QueryGroupingArrayOutput) ToQueryGroupingArrayOutput() QueryGroupingArrayOutput {
	return o
}

func (o QueryGroupingArrayOutput) ToQueryGroupingArrayOutputWithContext(ctx context.Context) QueryGroupingArrayOutput {
	return o
}

func (o QueryGroupingArrayOutput) Index(i pulumi.IntInput) QueryGroupingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) QueryGrouping {
		return vs[0].([]QueryGrouping)[vs[1].(int)]
	}).(QueryGroupingOutput)
}

type QueryGroupingResponse struct {
	Name string `pulumi:"name"`
	Type string `pulumi:"type"`
}





type QueryGroupingResponseInput interface {
	pulumi.Input

	ToQueryGroupingResponseOutput() QueryGroupingResponseOutput
	ToQueryGroupingResponseOutputWithContext(context.Context) QueryGroupingResponseOutput
}

type QueryGroupingResponseArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
	Type pulumi.StringInput `pulumi:"type"`
}

func (QueryGroupingResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryGroupingResponse)(nil)).Elem()
}

func (i QueryGroupingResponseArgs) ToQueryGroupingResponseOutput() QueryGroupingResponseOutput {
	return i.ToQueryGroupingResponseOutputWithContext(context.Background())
}

func (i QueryGroupingResponseArgs) ToQueryGroupingResponseOutputWithContext(ctx context.Context) QueryGroupingResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryGroupingResponseOutput)
}





type QueryGroupingResponseArrayInput interface {
	pulumi.Input

	ToQueryGroupingResponseArrayOutput() QueryGroupingResponseArrayOutput
	ToQueryGroupingResponseArrayOutputWithContext(context.Context) QueryGroupingResponseArrayOutput
}

type QueryGroupingResponseArray []QueryGroupingResponseInput

func (QueryGroupingResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]QueryGroupingResponse)(nil)).Elem()
}

func (i QueryGroupingResponseArray) ToQueryGroupingResponseArrayOutput() QueryGroupingResponseArrayOutput {
	return i.ToQueryGroupingResponseArrayOutputWithContext(context.Background())
}

func (i QueryGroupingResponseArray) ToQueryGroupingResponseArrayOutputWithContext(ctx context.Context) QueryGroupingResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryGroupingResponseArrayOutput)
}

type QueryGroupingResponseOutput struct{ *pulumi.OutputState }

func (QueryGroupingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryGroupingResponse)(nil)).Elem()
}

func (o QueryGroupingResponseOutput) ToQueryGroupingResponseOutput() QueryGroupingResponseOutput {
	return o
}

func (o QueryGroupingResponseOutput) ToQueryGroupingResponseOutputWithContext(ctx context.Context) QueryGroupingResponseOutput {
	return o
}

func (o QueryGroupingResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v QueryGroupingResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o QueryGroupingResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v QueryGroupingResponse) string { return v.Type }).(pulumi.StringOutput)
}

type QueryGroupingResponseArrayOutput struct{ *pulumi.OutputState }

func (QueryGroupingResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]QueryGroupingResponse)(nil)).Elem()
}

func (o QueryGroupingResponseArrayOutput) ToQueryGroupingResponseArrayOutput() QueryGroupingResponseArrayOutput {
	return o
}

func (o QueryGroupingResponseArrayOutput) ToQueryGroupingResponseArrayOutputWithContext(ctx context.Context) QueryGroupingResponseArrayOutput {
	return o
}

func (o QueryGroupingResponseArrayOutput) Index(i pulumi.IntInput) QueryGroupingResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) QueryGroupingResponse {
		return vs[0].([]QueryGroupingResponse)[vs[1].(int)]
	}).(QueryGroupingResponseOutput)
}

type QuerySortingConfiguration struct {
	Name                  *string `pulumi:"name"`
	QuerySortingDirection *string `pulumi:"querySortingDirection"`
}





type QuerySortingConfigurationInput interface {
	pulumi.Input

	ToQuerySortingConfigurationOutput() QuerySortingConfigurationOutput
	ToQuerySortingConfigurationOutputWithContext(context.Context) QuerySortingConfigurationOutput
}

type QuerySortingConfigurationArgs struct {
	Name                  pulumi.StringPtrInput `pulumi:"name"`
	QuerySortingDirection pulumi.StringPtrInput `pulumi:"querySortingDirection"`
}

func (QuerySortingConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QuerySortingConfiguration)(nil)).Elem()
}

func (i QuerySortingConfigurationArgs) ToQuerySortingConfigurationOutput() QuerySortingConfigurationOutput {
	return i.ToQuerySortingConfigurationOutputWithContext(context.Background())
}

func (i QuerySortingConfigurationArgs) ToQuerySortingConfigurationOutputWithContext(ctx context.Context) QuerySortingConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuerySortingConfigurationOutput)
}





type QuerySortingConfigurationArrayInput interface {
	pulumi.Input

	ToQuerySortingConfigurationArrayOutput() QuerySortingConfigurationArrayOutput
	ToQuerySortingConfigurationArrayOutputWithContext(context.Context) QuerySortingConfigurationArrayOutput
}

type QuerySortingConfigurationArray []QuerySortingConfigurationInput

func (QuerySortingConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]QuerySortingConfiguration)(nil)).Elem()
}

func (i QuerySortingConfigurationArray) ToQuerySortingConfigurationArrayOutput() QuerySortingConfigurationArrayOutput {
	return i.ToQuerySortingConfigurationArrayOutputWithContext(context.Background())
}

func (i QuerySortingConfigurationArray) ToQuerySortingConfigurationArrayOutputWithContext(ctx context.Context) QuerySortingConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuerySortingConfigurationArrayOutput)
}

type QuerySortingConfigurationOutput struct{ *pulumi.OutputState }

func (QuerySortingConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QuerySortingConfiguration)(nil)).Elem()
}

func (o QuerySortingConfigurationOutput) ToQuerySortingConfigurationOutput() QuerySortingConfigurationOutput {
	return o
}

func (o QuerySortingConfigurationOutput) ToQuerySortingConfigurationOutputWithContext(ctx context.Context) QuerySortingConfigurationOutput {
	return o
}

func (o QuerySortingConfigurationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QuerySortingConfiguration) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o QuerySortingConfigurationOutput) QuerySortingDirection() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QuerySortingConfiguration) *string { return v.QuerySortingDirection }).(pulumi.StringPtrOutput)
}

type QuerySortingConfigurationArrayOutput struct{ *pulumi.OutputState }

func (QuerySortingConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]QuerySortingConfiguration)(nil)).Elem()
}

func (o QuerySortingConfigurationArrayOutput) ToQuerySortingConfigurationArrayOutput() QuerySortingConfigurationArrayOutput {
	return o
}

func (o QuerySortingConfigurationArrayOutput) ToQuerySortingConfigurationArrayOutputWithContext(ctx context.Context) QuerySortingConfigurationArrayOutput {
	return o
}

func (o QuerySortingConfigurationArrayOutput) Index(i pulumi.IntInput) QuerySortingConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) QuerySortingConfiguration {
		return vs[0].([]QuerySortingConfiguration)[vs[1].(int)]
	}).(QuerySortingConfigurationOutput)
}

type QuerySortingConfigurationResponse struct {
	Name                  *string `pulumi:"name"`
	QuerySortingDirection *string `pulumi:"querySortingDirection"`
}





type QuerySortingConfigurationResponseInput interface {
	pulumi.Input

	ToQuerySortingConfigurationResponseOutput() QuerySortingConfigurationResponseOutput
	ToQuerySortingConfigurationResponseOutputWithContext(context.Context) QuerySortingConfigurationResponseOutput
}

type QuerySortingConfigurationResponseArgs struct {
	Name                  pulumi.StringPtrInput `pulumi:"name"`
	QuerySortingDirection pulumi.StringPtrInput `pulumi:"querySortingDirection"`
}

func (QuerySortingConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QuerySortingConfigurationResponse)(nil)).Elem()
}

func (i QuerySortingConfigurationResponseArgs) ToQuerySortingConfigurationResponseOutput() QuerySortingConfigurationResponseOutput {
	return i.ToQuerySortingConfigurationResponseOutputWithContext(context.Background())
}

func (i QuerySortingConfigurationResponseArgs) ToQuerySortingConfigurationResponseOutputWithContext(ctx context.Context) QuerySortingConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuerySortingConfigurationResponseOutput)
}





type QuerySortingConfigurationResponseArrayInput interface {
	pulumi.Input

	ToQuerySortingConfigurationResponseArrayOutput() QuerySortingConfigurationResponseArrayOutput
	ToQuerySortingConfigurationResponseArrayOutputWithContext(context.Context) QuerySortingConfigurationResponseArrayOutput
}

type QuerySortingConfigurationResponseArray []QuerySortingConfigurationResponseInput

func (QuerySortingConfigurationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]QuerySortingConfigurationResponse)(nil)).Elem()
}

func (i QuerySortingConfigurationResponseArray) ToQuerySortingConfigurationResponseArrayOutput() QuerySortingConfigurationResponseArrayOutput {
	return i.ToQuerySortingConfigurationResponseArrayOutputWithContext(context.Background())
}

func (i QuerySortingConfigurationResponseArray) ToQuerySortingConfigurationResponseArrayOutputWithContext(ctx context.Context) QuerySortingConfigurationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuerySortingConfigurationResponseArrayOutput)
}

type QuerySortingConfigurationResponseOutput struct{ *pulumi.OutputState }

func (QuerySortingConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QuerySortingConfigurationResponse)(nil)).Elem()
}

func (o QuerySortingConfigurationResponseOutput) ToQuerySortingConfigurationResponseOutput() QuerySortingConfigurationResponseOutput {
	return o
}

func (o QuerySortingConfigurationResponseOutput) ToQuerySortingConfigurationResponseOutputWithContext(ctx context.Context) QuerySortingConfigurationResponseOutput {
	return o
}

func (o QuerySortingConfigurationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QuerySortingConfigurationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o QuerySortingConfigurationResponseOutput) QuerySortingDirection() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QuerySortingConfigurationResponse) *string { return v.QuerySortingDirection }).(pulumi.StringPtrOutput)
}

type QuerySortingConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (QuerySortingConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]QuerySortingConfigurationResponse)(nil)).Elem()
}

func (o QuerySortingConfigurationResponseArrayOutput) ToQuerySortingConfigurationResponseArrayOutput() QuerySortingConfigurationResponseArrayOutput {
	return o
}

func (o QuerySortingConfigurationResponseArrayOutput) ToQuerySortingConfigurationResponseArrayOutputWithContext(ctx context.Context) QuerySortingConfigurationResponseArrayOutput {
	return o
}

func (o QuerySortingConfigurationResponseArrayOutput) Index(i pulumi.IntInput) QuerySortingConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) QuerySortingConfigurationResponse {
		return vs[0].([]QuerySortingConfigurationResponse)[vs[1].(int)]
	}).(QuerySortingConfigurationResponseOutput)
}

type QueryTimePeriod struct {
	From string `pulumi:"from"`
	To   string `pulumi:"to"`
}





type QueryTimePeriodInput interface {
	pulumi.Input

	ToQueryTimePeriodOutput() QueryTimePeriodOutput
	ToQueryTimePeriodOutputWithContext(context.Context) QueryTimePeriodOutput
}

type QueryTimePeriodArgs struct {
	From pulumi.StringInput `pulumi:"from"`
	To   pulumi.StringInput `pulumi:"to"`
}

func (QueryTimePeriodArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryTimePeriod)(nil)).Elem()
}

func (i QueryTimePeriodArgs) ToQueryTimePeriodOutput() QueryTimePeriodOutput {
	return i.ToQueryTimePeriodOutputWithContext(context.Background())
}

func (i QueryTimePeriodArgs) ToQueryTimePeriodOutputWithContext(ctx context.Context) QueryTimePeriodOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryTimePeriodOutput)
}

func (i QueryTimePeriodArgs) ToQueryTimePeriodPtrOutput() QueryTimePeriodPtrOutput {
	return i.ToQueryTimePeriodPtrOutputWithContext(context.Background())
}

func (i QueryTimePeriodArgs) ToQueryTimePeriodPtrOutputWithContext(ctx context.Context) QueryTimePeriodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryTimePeriodOutput).ToQueryTimePeriodPtrOutputWithContext(ctx)
}









type QueryTimePeriodPtrInput interface {
	pulumi.Input

	ToQueryTimePeriodPtrOutput() QueryTimePeriodPtrOutput
	ToQueryTimePeriodPtrOutputWithContext(context.Context) QueryTimePeriodPtrOutput
}

type queryTimePeriodPtrType QueryTimePeriodArgs

func QueryTimePeriodPtr(v *QueryTimePeriodArgs) QueryTimePeriodPtrInput {
	return (*queryTimePeriodPtrType)(v)
}

func (*queryTimePeriodPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryTimePeriod)(nil)).Elem()
}

func (i *queryTimePeriodPtrType) ToQueryTimePeriodPtrOutput() QueryTimePeriodPtrOutput {
	return i.ToQueryTimePeriodPtrOutputWithContext(context.Background())
}

func (i *queryTimePeriodPtrType) ToQueryTimePeriodPtrOutputWithContext(ctx context.Context) QueryTimePeriodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryTimePeriodPtrOutput)
}

type QueryTimePeriodOutput struct{ *pulumi.OutputState }

func (QueryTimePeriodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryTimePeriod)(nil)).Elem()
}

func (o QueryTimePeriodOutput) ToQueryTimePeriodOutput() QueryTimePeriodOutput {
	return o
}

func (o QueryTimePeriodOutput) ToQueryTimePeriodOutputWithContext(ctx context.Context) QueryTimePeriodOutput {
	return o
}

func (o QueryTimePeriodOutput) ToQueryTimePeriodPtrOutput() QueryTimePeriodPtrOutput {
	return o.ToQueryTimePeriodPtrOutputWithContext(context.Background())
}

func (o QueryTimePeriodOutput) ToQueryTimePeriodPtrOutputWithContext(ctx context.Context) QueryTimePeriodPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v QueryTimePeriod) *QueryTimePeriod {
		return &v
	}).(QueryTimePeriodPtrOutput)
}

func (o QueryTimePeriodOutput) From() pulumi.StringOutput {
	return o.ApplyT(func(v QueryTimePeriod) string { return v.From }).(pulumi.StringOutput)
}

func (o QueryTimePeriodOutput) To() pulumi.StringOutput {
	return o.ApplyT(func(v QueryTimePeriod) string { return v.To }).(pulumi.StringOutput)
}

type QueryTimePeriodPtrOutput struct{ *pulumi.OutputState }

func (QueryTimePeriodPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryTimePeriod)(nil)).Elem()
}

func (o QueryTimePeriodPtrOutput) ToQueryTimePeriodPtrOutput() QueryTimePeriodPtrOutput {
	return o
}

func (o QueryTimePeriodPtrOutput) ToQueryTimePeriodPtrOutputWithContext(ctx context.Context) QueryTimePeriodPtrOutput {
	return o
}

func (o QueryTimePeriodPtrOutput) Elem() QueryTimePeriodOutput {
	return o.ApplyT(func(v *QueryTimePeriod) QueryTimePeriod {
		if v != nil {
			return *v
		}
		var ret QueryTimePeriod
		return ret
	}).(QueryTimePeriodOutput)
}

func (o QueryTimePeriodPtrOutput) From() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueryTimePeriod) *string {
		if v == nil {
			return nil
		}
		return &v.From
	}).(pulumi.StringPtrOutput)
}

func (o QueryTimePeriodPtrOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueryTimePeriod) *string {
		if v == nil {
			return nil
		}
		return &v.To
	}).(pulumi.StringPtrOutput)
}

type QueryTimePeriodResponse struct {
	From string `pulumi:"from"`
	To   string `pulumi:"to"`
}





type QueryTimePeriodResponseInput interface {
	pulumi.Input

	ToQueryTimePeriodResponseOutput() QueryTimePeriodResponseOutput
	ToQueryTimePeriodResponseOutputWithContext(context.Context) QueryTimePeriodResponseOutput
}

type QueryTimePeriodResponseArgs struct {
	From pulumi.StringInput `pulumi:"from"`
	To   pulumi.StringInput `pulumi:"to"`
}

func (QueryTimePeriodResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryTimePeriodResponse)(nil)).Elem()
}

func (i QueryTimePeriodResponseArgs) ToQueryTimePeriodResponseOutput() QueryTimePeriodResponseOutput {
	return i.ToQueryTimePeriodResponseOutputWithContext(context.Background())
}

func (i QueryTimePeriodResponseArgs) ToQueryTimePeriodResponseOutputWithContext(ctx context.Context) QueryTimePeriodResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryTimePeriodResponseOutput)
}

func (i QueryTimePeriodResponseArgs) ToQueryTimePeriodResponsePtrOutput() QueryTimePeriodResponsePtrOutput {
	return i.ToQueryTimePeriodResponsePtrOutputWithContext(context.Background())
}

func (i QueryTimePeriodResponseArgs) ToQueryTimePeriodResponsePtrOutputWithContext(ctx context.Context) QueryTimePeriodResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryTimePeriodResponseOutput).ToQueryTimePeriodResponsePtrOutputWithContext(ctx)
}









type QueryTimePeriodResponsePtrInput interface {
	pulumi.Input

	ToQueryTimePeriodResponsePtrOutput() QueryTimePeriodResponsePtrOutput
	ToQueryTimePeriodResponsePtrOutputWithContext(context.Context) QueryTimePeriodResponsePtrOutput
}

type queryTimePeriodResponsePtrType QueryTimePeriodResponseArgs

func QueryTimePeriodResponsePtr(v *QueryTimePeriodResponseArgs) QueryTimePeriodResponsePtrInput {
	return (*queryTimePeriodResponsePtrType)(v)
}

func (*queryTimePeriodResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryTimePeriodResponse)(nil)).Elem()
}

func (i *queryTimePeriodResponsePtrType) ToQueryTimePeriodResponsePtrOutput() QueryTimePeriodResponsePtrOutput {
	return i.ToQueryTimePeriodResponsePtrOutputWithContext(context.Background())
}

func (i *queryTimePeriodResponsePtrType) ToQueryTimePeriodResponsePtrOutputWithContext(ctx context.Context) QueryTimePeriodResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryTimePeriodResponsePtrOutput)
}

type QueryTimePeriodResponseOutput struct{ *pulumi.OutputState }

func (QueryTimePeriodResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryTimePeriodResponse)(nil)).Elem()
}

func (o QueryTimePeriodResponseOutput) ToQueryTimePeriodResponseOutput() QueryTimePeriodResponseOutput {
	return o
}

func (o QueryTimePeriodResponseOutput) ToQueryTimePeriodResponseOutputWithContext(ctx context.Context) QueryTimePeriodResponseOutput {
	return o
}

func (o QueryTimePeriodResponseOutput) ToQueryTimePeriodResponsePtrOutput() QueryTimePeriodResponsePtrOutput {
	return o.ToQueryTimePeriodResponsePtrOutputWithContext(context.Background())
}

func (o QueryTimePeriodResponseOutput) ToQueryTimePeriodResponsePtrOutputWithContext(ctx context.Context) QueryTimePeriodResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v QueryTimePeriodResponse) *QueryTimePeriodResponse {
		return &v
	}).(QueryTimePeriodResponsePtrOutput)
}

func (o QueryTimePeriodResponseOutput) From() pulumi.StringOutput {
	return o.ApplyT(func(v QueryTimePeriodResponse) string { return v.From }).(pulumi.StringOutput)
}

func (o QueryTimePeriodResponseOutput) To() pulumi.StringOutput {
	return o.ApplyT(func(v QueryTimePeriodResponse) string { return v.To }).(pulumi.StringOutput)
}

type QueryTimePeriodResponsePtrOutput struct{ *pulumi.OutputState }

func (QueryTimePeriodResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryTimePeriodResponse)(nil)).Elem()
}

func (o QueryTimePeriodResponsePtrOutput) ToQueryTimePeriodResponsePtrOutput() QueryTimePeriodResponsePtrOutput {
	return o
}

func (o QueryTimePeriodResponsePtrOutput) ToQueryTimePeriodResponsePtrOutputWithContext(ctx context.Context) QueryTimePeriodResponsePtrOutput {
	return o
}

func (o QueryTimePeriodResponsePtrOutput) Elem() QueryTimePeriodResponseOutput {
	return o.ApplyT(func(v *QueryTimePeriodResponse) QueryTimePeriodResponse {
		if v != nil {
			return *v
		}
		var ret QueryTimePeriodResponse
		return ret
	}).(QueryTimePeriodResponseOutput)
}

func (o QueryTimePeriodResponsePtrOutput) From() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueryTimePeriodResponse) *string {
		if v == nil {
			return nil
		}
		return &v.From
	}).(pulumi.StringPtrOutput)
}

func (o QueryTimePeriodResponsePtrOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueryTimePeriodResponse) *string {
		if v == nil {
			return nil
		}
		return &v.To
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ExportDeliveryDestinationOutput{})
	pulumi.RegisterOutputType(ExportDeliveryDestinationPtrOutput{})
	pulumi.RegisterOutputType(ExportDeliveryDestinationResponseOutput{})
	pulumi.RegisterOutputType(ExportDeliveryDestinationResponsePtrOutput{})
	pulumi.RegisterOutputType(ExportDeliveryInfoOutput{})
	pulumi.RegisterOutputType(ExportDeliveryInfoPtrOutput{})
	pulumi.RegisterOutputType(ExportDeliveryInfoResponseOutput{})
	pulumi.RegisterOutputType(ExportDeliveryInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(ExportRecurrencePeriodOutput{})
	pulumi.RegisterOutputType(ExportRecurrencePeriodPtrOutput{})
	pulumi.RegisterOutputType(ExportRecurrencePeriodResponseOutput{})
	pulumi.RegisterOutputType(ExportRecurrencePeriodResponsePtrOutput{})
	pulumi.RegisterOutputType(ExportScheduleOutput{})
	pulumi.RegisterOutputType(ExportSchedulePtrOutput{})
	pulumi.RegisterOutputType(ExportScheduleResponseOutput{})
	pulumi.RegisterOutputType(ExportScheduleResponsePtrOutput{})
	pulumi.RegisterOutputType(QueryAggregationOutput{})
	pulumi.RegisterOutputType(QueryAggregationMapOutput{})
	pulumi.RegisterOutputType(QueryAggregationResponseOutput{})
	pulumi.RegisterOutputType(QueryAggregationResponseMapOutput{})
	pulumi.RegisterOutputType(QueryComparisonExpressionOutput{})
	pulumi.RegisterOutputType(QueryComparisonExpressionPtrOutput{})
	pulumi.RegisterOutputType(QueryComparisonExpressionResponseOutput{})
	pulumi.RegisterOutputType(QueryComparisonExpressionResponsePtrOutput{})
	pulumi.RegisterOutputType(QueryDatasetOutput{})
	pulumi.RegisterOutputType(QueryDatasetPtrOutput{})
	pulumi.RegisterOutputType(QueryDatasetConfigurationOutput{})
	pulumi.RegisterOutputType(QueryDatasetConfigurationPtrOutput{})
	pulumi.RegisterOutputType(QueryDatasetConfigurationResponseOutput{})
	pulumi.RegisterOutputType(QueryDatasetConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(QueryDatasetResponseOutput{})
	pulumi.RegisterOutputType(QueryDatasetResponsePtrOutput{})
	pulumi.RegisterOutputType(QueryDefinitionOutput{})
	pulumi.RegisterOutputType(QueryDefinitionPtrOutput{})
	pulumi.RegisterOutputType(QueryDefinitionResponseOutput{})
	pulumi.RegisterOutputType(QueryDefinitionResponsePtrOutput{})
	pulumi.RegisterOutputType(QueryFilterOutput{})
	pulumi.RegisterOutputType(QueryFilterPtrOutput{})
	pulumi.RegisterOutputType(QueryFilterArrayOutput{})
	pulumi.RegisterOutputType(QueryFilterResponseOutput{})
	pulumi.RegisterOutputType(QueryFilterResponsePtrOutput{})
	pulumi.RegisterOutputType(QueryFilterResponseArrayOutput{})
	pulumi.RegisterOutputType(QueryGroupingOutput{})
	pulumi.RegisterOutputType(QueryGroupingArrayOutput{})
	pulumi.RegisterOutputType(QueryGroupingResponseOutput{})
	pulumi.RegisterOutputType(QueryGroupingResponseArrayOutput{})
	pulumi.RegisterOutputType(QuerySortingConfigurationOutput{})
	pulumi.RegisterOutputType(QuerySortingConfigurationArrayOutput{})
	pulumi.RegisterOutputType(QuerySortingConfigurationResponseOutput{})
	pulumi.RegisterOutputType(QuerySortingConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(QueryTimePeriodOutput{})
	pulumi.RegisterOutputType(QueryTimePeriodPtrOutput{})
	pulumi.RegisterOutputType(QueryTimePeriodResponseOutput{})
	pulumi.RegisterOutputType(QueryTimePeriodResponsePtrOutput{})
}
