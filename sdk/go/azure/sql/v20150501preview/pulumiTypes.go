


package v20150501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FailoverGroupReadOnlyEndpoint struct {
	FailoverPolicy *string `pulumi:"failoverPolicy"`
}





type FailoverGroupReadOnlyEndpointInput interface {
	pulumi.Input

	ToFailoverGroupReadOnlyEndpointOutput() FailoverGroupReadOnlyEndpointOutput
	ToFailoverGroupReadOnlyEndpointOutputWithContext(context.Context) FailoverGroupReadOnlyEndpointOutput
}

type FailoverGroupReadOnlyEndpointArgs struct {
	FailoverPolicy pulumi.StringPtrInput `pulumi:"failoverPolicy"`
}

func (FailoverGroupReadOnlyEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FailoverGroupReadOnlyEndpoint)(nil)).Elem()
}

func (i FailoverGroupReadOnlyEndpointArgs) ToFailoverGroupReadOnlyEndpointOutput() FailoverGroupReadOnlyEndpointOutput {
	return i.ToFailoverGroupReadOnlyEndpointOutputWithContext(context.Background())
}

func (i FailoverGroupReadOnlyEndpointArgs) ToFailoverGroupReadOnlyEndpointOutputWithContext(ctx context.Context) FailoverGroupReadOnlyEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverGroupReadOnlyEndpointOutput)
}

func (i FailoverGroupReadOnlyEndpointArgs) ToFailoverGroupReadOnlyEndpointPtrOutput() FailoverGroupReadOnlyEndpointPtrOutput {
	return i.ToFailoverGroupReadOnlyEndpointPtrOutputWithContext(context.Background())
}

func (i FailoverGroupReadOnlyEndpointArgs) ToFailoverGroupReadOnlyEndpointPtrOutputWithContext(ctx context.Context) FailoverGroupReadOnlyEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverGroupReadOnlyEndpointOutput).ToFailoverGroupReadOnlyEndpointPtrOutputWithContext(ctx)
}









type FailoverGroupReadOnlyEndpointPtrInput interface {
	pulumi.Input

	ToFailoverGroupReadOnlyEndpointPtrOutput() FailoverGroupReadOnlyEndpointPtrOutput
	ToFailoverGroupReadOnlyEndpointPtrOutputWithContext(context.Context) FailoverGroupReadOnlyEndpointPtrOutput
}

type failoverGroupReadOnlyEndpointPtrType FailoverGroupReadOnlyEndpointArgs

func FailoverGroupReadOnlyEndpointPtr(v *FailoverGroupReadOnlyEndpointArgs) FailoverGroupReadOnlyEndpointPtrInput {
	return (*failoverGroupReadOnlyEndpointPtrType)(v)
}

func (*failoverGroupReadOnlyEndpointPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FailoverGroupReadOnlyEndpoint)(nil)).Elem()
}

func (i *failoverGroupReadOnlyEndpointPtrType) ToFailoverGroupReadOnlyEndpointPtrOutput() FailoverGroupReadOnlyEndpointPtrOutput {
	return i.ToFailoverGroupReadOnlyEndpointPtrOutputWithContext(context.Background())
}

func (i *failoverGroupReadOnlyEndpointPtrType) ToFailoverGroupReadOnlyEndpointPtrOutputWithContext(ctx context.Context) FailoverGroupReadOnlyEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverGroupReadOnlyEndpointPtrOutput)
}

type FailoverGroupReadOnlyEndpointOutput struct{ *pulumi.OutputState }

func (FailoverGroupReadOnlyEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FailoverGroupReadOnlyEndpoint)(nil)).Elem()
}

func (o FailoverGroupReadOnlyEndpointOutput) ToFailoverGroupReadOnlyEndpointOutput() FailoverGroupReadOnlyEndpointOutput {
	return o
}

func (o FailoverGroupReadOnlyEndpointOutput) ToFailoverGroupReadOnlyEndpointOutputWithContext(ctx context.Context) FailoverGroupReadOnlyEndpointOutput {
	return o
}

func (o FailoverGroupReadOnlyEndpointOutput) ToFailoverGroupReadOnlyEndpointPtrOutput() FailoverGroupReadOnlyEndpointPtrOutput {
	return o.ToFailoverGroupReadOnlyEndpointPtrOutputWithContext(context.Background())
}

func (o FailoverGroupReadOnlyEndpointOutput) ToFailoverGroupReadOnlyEndpointPtrOutputWithContext(ctx context.Context) FailoverGroupReadOnlyEndpointPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FailoverGroupReadOnlyEndpoint) *FailoverGroupReadOnlyEndpoint {
		return &v
	}).(FailoverGroupReadOnlyEndpointPtrOutput)
}

func (o FailoverGroupReadOnlyEndpointOutput) FailoverPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FailoverGroupReadOnlyEndpoint) *string { return v.FailoverPolicy }).(pulumi.StringPtrOutput)
}

type FailoverGroupReadOnlyEndpointPtrOutput struct{ *pulumi.OutputState }

func (FailoverGroupReadOnlyEndpointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FailoverGroupReadOnlyEndpoint)(nil)).Elem()
}

func (o FailoverGroupReadOnlyEndpointPtrOutput) ToFailoverGroupReadOnlyEndpointPtrOutput() FailoverGroupReadOnlyEndpointPtrOutput {
	return o
}

func (o FailoverGroupReadOnlyEndpointPtrOutput) ToFailoverGroupReadOnlyEndpointPtrOutputWithContext(ctx context.Context) FailoverGroupReadOnlyEndpointPtrOutput {
	return o
}

func (o FailoverGroupReadOnlyEndpointPtrOutput) Elem() FailoverGroupReadOnlyEndpointOutput {
	return o.ApplyT(func(v *FailoverGroupReadOnlyEndpoint) FailoverGroupReadOnlyEndpoint {
		if v != nil {
			return *v
		}
		var ret FailoverGroupReadOnlyEndpoint
		return ret
	}).(FailoverGroupReadOnlyEndpointOutput)
}

func (o FailoverGroupReadOnlyEndpointPtrOutput) FailoverPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FailoverGroupReadOnlyEndpoint) *string {
		if v == nil {
			return nil
		}
		return v.FailoverPolicy
	}).(pulumi.StringPtrOutput)
}

type FailoverGroupReadOnlyEndpointResponse struct {
	FailoverPolicy *string `pulumi:"failoverPolicy"`
}





type FailoverGroupReadOnlyEndpointResponseInput interface {
	pulumi.Input

	ToFailoverGroupReadOnlyEndpointResponseOutput() FailoverGroupReadOnlyEndpointResponseOutput
	ToFailoverGroupReadOnlyEndpointResponseOutputWithContext(context.Context) FailoverGroupReadOnlyEndpointResponseOutput
}

type FailoverGroupReadOnlyEndpointResponseArgs struct {
	FailoverPolicy pulumi.StringPtrInput `pulumi:"failoverPolicy"`
}

func (FailoverGroupReadOnlyEndpointResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FailoverGroupReadOnlyEndpointResponse)(nil)).Elem()
}

func (i FailoverGroupReadOnlyEndpointResponseArgs) ToFailoverGroupReadOnlyEndpointResponseOutput() FailoverGroupReadOnlyEndpointResponseOutput {
	return i.ToFailoverGroupReadOnlyEndpointResponseOutputWithContext(context.Background())
}

func (i FailoverGroupReadOnlyEndpointResponseArgs) ToFailoverGroupReadOnlyEndpointResponseOutputWithContext(ctx context.Context) FailoverGroupReadOnlyEndpointResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverGroupReadOnlyEndpointResponseOutput)
}

func (i FailoverGroupReadOnlyEndpointResponseArgs) ToFailoverGroupReadOnlyEndpointResponsePtrOutput() FailoverGroupReadOnlyEndpointResponsePtrOutput {
	return i.ToFailoverGroupReadOnlyEndpointResponsePtrOutputWithContext(context.Background())
}

func (i FailoverGroupReadOnlyEndpointResponseArgs) ToFailoverGroupReadOnlyEndpointResponsePtrOutputWithContext(ctx context.Context) FailoverGroupReadOnlyEndpointResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverGroupReadOnlyEndpointResponseOutput).ToFailoverGroupReadOnlyEndpointResponsePtrOutputWithContext(ctx)
}









type FailoverGroupReadOnlyEndpointResponsePtrInput interface {
	pulumi.Input

	ToFailoverGroupReadOnlyEndpointResponsePtrOutput() FailoverGroupReadOnlyEndpointResponsePtrOutput
	ToFailoverGroupReadOnlyEndpointResponsePtrOutputWithContext(context.Context) FailoverGroupReadOnlyEndpointResponsePtrOutput
}

type failoverGroupReadOnlyEndpointResponsePtrType FailoverGroupReadOnlyEndpointResponseArgs

func FailoverGroupReadOnlyEndpointResponsePtr(v *FailoverGroupReadOnlyEndpointResponseArgs) FailoverGroupReadOnlyEndpointResponsePtrInput {
	return (*failoverGroupReadOnlyEndpointResponsePtrType)(v)
}

func (*failoverGroupReadOnlyEndpointResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FailoverGroupReadOnlyEndpointResponse)(nil)).Elem()
}

func (i *failoverGroupReadOnlyEndpointResponsePtrType) ToFailoverGroupReadOnlyEndpointResponsePtrOutput() FailoverGroupReadOnlyEndpointResponsePtrOutput {
	return i.ToFailoverGroupReadOnlyEndpointResponsePtrOutputWithContext(context.Background())
}

func (i *failoverGroupReadOnlyEndpointResponsePtrType) ToFailoverGroupReadOnlyEndpointResponsePtrOutputWithContext(ctx context.Context) FailoverGroupReadOnlyEndpointResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverGroupReadOnlyEndpointResponsePtrOutput)
}

type FailoverGroupReadOnlyEndpointResponseOutput struct{ *pulumi.OutputState }

func (FailoverGroupReadOnlyEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FailoverGroupReadOnlyEndpointResponse)(nil)).Elem()
}

func (o FailoverGroupReadOnlyEndpointResponseOutput) ToFailoverGroupReadOnlyEndpointResponseOutput() FailoverGroupReadOnlyEndpointResponseOutput {
	return o
}

func (o FailoverGroupReadOnlyEndpointResponseOutput) ToFailoverGroupReadOnlyEndpointResponseOutputWithContext(ctx context.Context) FailoverGroupReadOnlyEndpointResponseOutput {
	return o
}

func (o FailoverGroupReadOnlyEndpointResponseOutput) ToFailoverGroupReadOnlyEndpointResponsePtrOutput() FailoverGroupReadOnlyEndpointResponsePtrOutput {
	return o.ToFailoverGroupReadOnlyEndpointResponsePtrOutputWithContext(context.Background())
}

func (o FailoverGroupReadOnlyEndpointResponseOutput) ToFailoverGroupReadOnlyEndpointResponsePtrOutputWithContext(ctx context.Context) FailoverGroupReadOnlyEndpointResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FailoverGroupReadOnlyEndpointResponse) *FailoverGroupReadOnlyEndpointResponse {
		return &v
	}).(FailoverGroupReadOnlyEndpointResponsePtrOutput)
}

func (o FailoverGroupReadOnlyEndpointResponseOutput) FailoverPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FailoverGroupReadOnlyEndpointResponse) *string { return v.FailoverPolicy }).(pulumi.StringPtrOutput)
}

type FailoverGroupReadOnlyEndpointResponsePtrOutput struct{ *pulumi.OutputState }

func (FailoverGroupReadOnlyEndpointResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FailoverGroupReadOnlyEndpointResponse)(nil)).Elem()
}

func (o FailoverGroupReadOnlyEndpointResponsePtrOutput) ToFailoverGroupReadOnlyEndpointResponsePtrOutput() FailoverGroupReadOnlyEndpointResponsePtrOutput {
	return o
}

func (o FailoverGroupReadOnlyEndpointResponsePtrOutput) ToFailoverGroupReadOnlyEndpointResponsePtrOutputWithContext(ctx context.Context) FailoverGroupReadOnlyEndpointResponsePtrOutput {
	return o
}

func (o FailoverGroupReadOnlyEndpointResponsePtrOutput) Elem() FailoverGroupReadOnlyEndpointResponseOutput {
	return o.ApplyT(func(v *FailoverGroupReadOnlyEndpointResponse) FailoverGroupReadOnlyEndpointResponse {
		if v != nil {
			return *v
		}
		var ret FailoverGroupReadOnlyEndpointResponse
		return ret
	}).(FailoverGroupReadOnlyEndpointResponseOutput)
}

func (o FailoverGroupReadOnlyEndpointResponsePtrOutput) FailoverPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FailoverGroupReadOnlyEndpointResponse) *string {
		if v == nil {
			return nil
		}
		return v.FailoverPolicy
	}).(pulumi.StringPtrOutput)
}

type FailoverGroupReadWriteEndpoint struct {
	FailoverPolicy                         string `pulumi:"failoverPolicy"`
	FailoverWithDataLossGracePeriodMinutes *int   `pulumi:"failoverWithDataLossGracePeriodMinutes"`
}





type FailoverGroupReadWriteEndpointInput interface {
	pulumi.Input

	ToFailoverGroupReadWriteEndpointOutput() FailoverGroupReadWriteEndpointOutput
	ToFailoverGroupReadWriteEndpointOutputWithContext(context.Context) FailoverGroupReadWriteEndpointOutput
}

type FailoverGroupReadWriteEndpointArgs struct {
	FailoverPolicy                         pulumi.StringInput `pulumi:"failoverPolicy"`
	FailoverWithDataLossGracePeriodMinutes pulumi.IntPtrInput `pulumi:"failoverWithDataLossGracePeriodMinutes"`
}

func (FailoverGroupReadWriteEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FailoverGroupReadWriteEndpoint)(nil)).Elem()
}

func (i FailoverGroupReadWriteEndpointArgs) ToFailoverGroupReadWriteEndpointOutput() FailoverGroupReadWriteEndpointOutput {
	return i.ToFailoverGroupReadWriteEndpointOutputWithContext(context.Background())
}

func (i FailoverGroupReadWriteEndpointArgs) ToFailoverGroupReadWriteEndpointOutputWithContext(ctx context.Context) FailoverGroupReadWriteEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverGroupReadWriteEndpointOutput)
}

func (i FailoverGroupReadWriteEndpointArgs) ToFailoverGroupReadWriteEndpointPtrOutput() FailoverGroupReadWriteEndpointPtrOutput {
	return i.ToFailoverGroupReadWriteEndpointPtrOutputWithContext(context.Background())
}

func (i FailoverGroupReadWriteEndpointArgs) ToFailoverGroupReadWriteEndpointPtrOutputWithContext(ctx context.Context) FailoverGroupReadWriteEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverGroupReadWriteEndpointOutput).ToFailoverGroupReadWriteEndpointPtrOutputWithContext(ctx)
}









type FailoverGroupReadWriteEndpointPtrInput interface {
	pulumi.Input

	ToFailoverGroupReadWriteEndpointPtrOutput() FailoverGroupReadWriteEndpointPtrOutput
	ToFailoverGroupReadWriteEndpointPtrOutputWithContext(context.Context) FailoverGroupReadWriteEndpointPtrOutput
}

type failoverGroupReadWriteEndpointPtrType FailoverGroupReadWriteEndpointArgs

func FailoverGroupReadWriteEndpointPtr(v *FailoverGroupReadWriteEndpointArgs) FailoverGroupReadWriteEndpointPtrInput {
	return (*failoverGroupReadWriteEndpointPtrType)(v)
}

func (*failoverGroupReadWriteEndpointPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FailoverGroupReadWriteEndpoint)(nil)).Elem()
}

func (i *failoverGroupReadWriteEndpointPtrType) ToFailoverGroupReadWriteEndpointPtrOutput() FailoverGroupReadWriteEndpointPtrOutput {
	return i.ToFailoverGroupReadWriteEndpointPtrOutputWithContext(context.Background())
}

func (i *failoverGroupReadWriteEndpointPtrType) ToFailoverGroupReadWriteEndpointPtrOutputWithContext(ctx context.Context) FailoverGroupReadWriteEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverGroupReadWriteEndpointPtrOutput)
}

type FailoverGroupReadWriteEndpointOutput struct{ *pulumi.OutputState }

func (FailoverGroupReadWriteEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FailoverGroupReadWriteEndpoint)(nil)).Elem()
}

func (o FailoverGroupReadWriteEndpointOutput) ToFailoverGroupReadWriteEndpointOutput() FailoverGroupReadWriteEndpointOutput {
	return o
}

func (o FailoverGroupReadWriteEndpointOutput) ToFailoverGroupReadWriteEndpointOutputWithContext(ctx context.Context) FailoverGroupReadWriteEndpointOutput {
	return o
}

func (o FailoverGroupReadWriteEndpointOutput) ToFailoverGroupReadWriteEndpointPtrOutput() FailoverGroupReadWriteEndpointPtrOutput {
	return o.ToFailoverGroupReadWriteEndpointPtrOutputWithContext(context.Background())
}

func (o FailoverGroupReadWriteEndpointOutput) ToFailoverGroupReadWriteEndpointPtrOutputWithContext(ctx context.Context) FailoverGroupReadWriteEndpointPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FailoverGroupReadWriteEndpoint) *FailoverGroupReadWriteEndpoint {
		return &v
	}).(FailoverGroupReadWriteEndpointPtrOutput)
}

func (o FailoverGroupReadWriteEndpointOutput) FailoverPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v FailoverGroupReadWriteEndpoint) string { return v.FailoverPolicy }).(pulumi.StringOutput)
}

func (o FailoverGroupReadWriteEndpointOutput) FailoverWithDataLossGracePeriodMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FailoverGroupReadWriteEndpoint) *int { return v.FailoverWithDataLossGracePeriodMinutes }).(pulumi.IntPtrOutput)
}

type FailoverGroupReadWriteEndpointPtrOutput struct{ *pulumi.OutputState }

func (FailoverGroupReadWriteEndpointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FailoverGroupReadWriteEndpoint)(nil)).Elem()
}

func (o FailoverGroupReadWriteEndpointPtrOutput) ToFailoverGroupReadWriteEndpointPtrOutput() FailoverGroupReadWriteEndpointPtrOutput {
	return o
}

func (o FailoverGroupReadWriteEndpointPtrOutput) ToFailoverGroupReadWriteEndpointPtrOutputWithContext(ctx context.Context) FailoverGroupReadWriteEndpointPtrOutput {
	return o
}

func (o FailoverGroupReadWriteEndpointPtrOutput) Elem() FailoverGroupReadWriteEndpointOutput {
	return o.ApplyT(func(v *FailoverGroupReadWriteEndpoint) FailoverGroupReadWriteEndpoint {
		if v != nil {
			return *v
		}
		var ret FailoverGroupReadWriteEndpoint
		return ret
	}).(FailoverGroupReadWriteEndpointOutput)
}

func (o FailoverGroupReadWriteEndpointPtrOutput) FailoverPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FailoverGroupReadWriteEndpoint) *string {
		if v == nil {
			return nil
		}
		return &v.FailoverPolicy
	}).(pulumi.StringPtrOutput)
}

func (o FailoverGroupReadWriteEndpointPtrOutput) FailoverWithDataLossGracePeriodMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FailoverGroupReadWriteEndpoint) *int {
		if v == nil {
			return nil
		}
		return v.FailoverWithDataLossGracePeriodMinutes
	}).(pulumi.IntPtrOutput)
}

type FailoverGroupReadWriteEndpointResponse struct {
	FailoverPolicy                         string `pulumi:"failoverPolicy"`
	FailoverWithDataLossGracePeriodMinutes *int   `pulumi:"failoverWithDataLossGracePeriodMinutes"`
}





type FailoverGroupReadWriteEndpointResponseInput interface {
	pulumi.Input

	ToFailoverGroupReadWriteEndpointResponseOutput() FailoverGroupReadWriteEndpointResponseOutput
	ToFailoverGroupReadWriteEndpointResponseOutputWithContext(context.Context) FailoverGroupReadWriteEndpointResponseOutput
}

type FailoverGroupReadWriteEndpointResponseArgs struct {
	FailoverPolicy                         pulumi.StringInput `pulumi:"failoverPolicy"`
	FailoverWithDataLossGracePeriodMinutes pulumi.IntPtrInput `pulumi:"failoverWithDataLossGracePeriodMinutes"`
}

func (FailoverGroupReadWriteEndpointResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FailoverGroupReadWriteEndpointResponse)(nil)).Elem()
}

func (i FailoverGroupReadWriteEndpointResponseArgs) ToFailoverGroupReadWriteEndpointResponseOutput() FailoverGroupReadWriteEndpointResponseOutput {
	return i.ToFailoverGroupReadWriteEndpointResponseOutputWithContext(context.Background())
}

func (i FailoverGroupReadWriteEndpointResponseArgs) ToFailoverGroupReadWriteEndpointResponseOutputWithContext(ctx context.Context) FailoverGroupReadWriteEndpointResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverGroupReadWriteEndpointResponseOutput)
}

func (i FailoverGroupReadWriteEndpointResponseArgs) ToFailoverGroupReadWriteEndpointResponsePtrOutput() FailoverGroupReadWriteEndpointResponsePtrOutput {
	return i.ToFailoverGroupReadWriteEndpointResponsePtrOutputWithContext(context.Background())
}

func (i FailoverGroupReadWriteEndpointResponseArgs) ToFailoverGroupReadWriteEndpointResponsePtrOutputWithContext(ctx context.Context) FailoverGroupReadWriteEndpointResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverGroupReadWriteEndpointResponseOutput).ToFailoverGroupReadWriteEndpointResponsePtrOutputWithContext(ctx)
}









type FailoverGroupReadWriteEndpointResponsePtrInput interface {
	pulumi.Input

	ToFailoverGroupReadWriteEndpointResponsePtrOutput() FailoverGroupReadWriteEndpointResponsePtrOutput
	ToFailoverGroupReadWriteEndpointResponsePtrOutputWithContext(context.Context) FailoverGroupReadWriteEndpointResponsePtrOutput
}

type failoverGroupReadWriteEndpointResponsePtrType FailoverGroupReadWriteEndpointResponseArgs

func FailoverGroupReadWriteEndpointResponsePtr(v *FailoverGroupReadWriteEndpointResponseArgs) FailoverGroupReadWriteEndpointResponsePtrInput {
	return (*failoverGroupReadWriteEndpointResponsePtrType)(v)
}

func (*failoverGroupReadWriteEndpointResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FailoverGroupReadWriteEndpointResponse)(nil)).Elem()
}

func (i *failoverGroupReadWriteEndpointResponsePtrType) ToFailoverGroupReadWriteEndpointResponsePtrOutput() FailoverGroupReadWriteEndpointResponsePtrOutput {
	return i.ToFailoverGroupReadWriteEndpointResponsePtrOutputWithContext(context.Background())
}

func (i *failoverGroupReadWriteEndpointResponsePtrType) ToFailoverGroupReadWriteEndpointResponsePtrOutputWithContext(ctx context.Context) FailoverGroupReadWriteEndpointResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverGroupReadWriteEndpointResponsePtrOutput)
}

type FailoverGroupReadWriteEndpointResponseOutput struct{ *pulumi.OutputState }

func (FailoverGroupReadWriteEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FailoverGroupReadWriteEndpointResponse)(nil)).Elem()
}

func (o FailoverGroupReadWriteEndpointResponseOutput) ToFailoverGroupReadWriteEndpointResponseOutput() FailoverGroupReadWriteEndpointResponseOutput {
	return o
}

func (o FailoverGroupReadWriteEndpointResponseOutput) ToFailoverGroupReadWriteEndpointResponseOutputWithContext(ctx context.Context) FailoverGroupReadWriteEndpointResponseOutput {
	return o
}

func (o FailoverGroupReadWriteEndpointResponseOutput) ToFailoverGroupReadWriteEndpointResponsePtrOutput() FailoverGroupReadWriteEndpointResponsePtrOutput {
	return o.ToFailoverGroupReadWriteEndpointResponsePtrOutputWithContext(context.Background())
}

func (o FailoverGroupReadWriteEndpointResponseOutput) ToFailoverGroupReadWriteEndpointResponsePtrOutputWithContext(ctx context.Context) FailoverGroupReadWriteEndpointResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FailoverGroupReadWriteEndpointResponse) *FailoverGroupReadWriteEndpointResponse {
		return &v
	}).(FailoverGroupReadWriteEndpointResponsePtrOutput)
}

func (o FailoverGroupReadWriteEndpointResponseOutput) FailoverPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v FailoverGroupReadWriteEndpointResponse) string { return v.FailoverPolicy }).(pulumi.StringOutput)
}

func (o FailoverGroupReadWriteEndpointResponseOutput) FailoverWithDataLossGracePeriodMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FailoverGroupReadWriteEndpointResponse) *int { return v.FailoverWithDataLossGracePeriodMinutes }).(pulumi.IntPtrOutput)
}

type FailoverGroupReadWriteEndpointResponsePtrOutput struct{ *pulumi.OutputState }

func (FailoverGroupReadWriteEndpointResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FailoverGroupReadWriteEndpointResponse)(nil)).Elem()
}

func (o FailoverGroupReadWriteEndpointResponsePtrOutput) ToFailoverGroupReadWriteEndpointResponsePtrOutput() FailoverGroupReadWriteEndpointResponsePtrOutput {
	return o
}

func (o FailoverGroupReadWriteEndpointResponsePtrOutput) ToFailoverGroupReadWriteEndpointResponsePtrOutputWithContext(ctx context.Context) FailoverGroupReadWriteEndpointResponsePtrOutput {
	return o
}

func (o FailoverGroupReadWriteEndpointResponsePtrOutput) Elem() FailoverGroupReadWriteEndpointResponseOutput {
	return o.ApplyT(func(v *FailoverGroupReadWriteEndpointResponse) FailoverGroupReadWriteEndpointResponse {
		if v != nil {
			return *v
		}
		var ret FailoverGroupReadWriteEndpointResponse
		return ret
	}).(FailoverGroupReadWriteEndpointResponseOutput)
}

func (o FailoverGroupReadWriteEndpointResponsePtrOutput) FailoverPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FailoverGroupReadWriteEndpointResponse) *string {
		if v == nil {
			return nil
		}
		return &v.FailoverPolicy
	}).(pulumi.StringPtrOutput)
}

func (o FailoverGroupReadWriteEndpointResponsePtrOutput) FailoverWithDataLossGracePeriodMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FailoverGroupReadWriteEndpointResponse) *int {
		if v == nil {
			return nil
		}
		return v.FailoverWithDataLossGracePeriodMinutes
	}).(pulumi.IntPtrOutput)
}

type PartnerInfo struct {
	Id string `pulumi:"id"`
}





type PartnerInfoInput interface {
	pulumi.Input

	ToPartnerInfoOutput() PartnerInfoOutput
	ToPartnerInfoOutputWithContext(context.Context) PartnerInfoOutput
}

type PartnerInfoArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (PartnerInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PartnerInfo)(nil)).Elem()
}

func (i PartnerInfoArgs) ToPartnerInfoOutput() PartnerInfoOutput {
	return i.ToPartnerInfoOutputWithContext(context.Background())
}

func (i PartnerInfoArgs) ToPartnerInfoOutputWithContext(ctx context.Context) PartnerInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartnerInfoOutput)
}





type PartnerInfoArrayInput interface {
	pulumi.Input

	ToPartnerInfoArrayOutput() PartnerInfoArrayOutput
	ToPartnerInfoArrayOutputWithContext(context.Context) PartnerInfoArrayOutput
}

type PartnerInfoArray []PartnerInfoInput

func (PartnerInfoArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PartnerInfo)(nil)).Elem()
}

func (i PartnerInfoArray) ToPartnerInfoArrayOutput() PartnerInfoArrayOutput {
	return i.ToPartnerInfoArrayOutputWithContext(context.Background())
}

func (i PartnerInfoArray) ToPartnerInfoArrayOutputWithContext(ctx context.Context) PartnerInfoArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartnerInfoArrayOutput)
}

type PartnerInfoOutput struct{ *pulumi.OutputState }

func (PartnerInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PartnerInfo)(nil)).Elem()
}

func (o PartnerInfoOutput) ToPartnerInfoOutput() PartnerInfoOutput {
	return o
}

func (o PartnerInfoOutput) ToPartnerInfoOutputWithContext(ctx context.Context) PartnerInfoOutput {
	return o
}

func (o PartnerInfoOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PartnerInfo) string { return v.Id }).(pulumi.StringOutput)
}

type PartnerInfoArrayOutput struct{ *pulumi.OutputState }

func (PartnerInfoArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PartnerInfo)(nil)).Elem()
}

func (o PartnerInfoArrayOutput) ToPartnerInfoArrayOutput() PartnerInfoArrayOutput {
	return o
}

func (o PartnerInfoArrayOutput) ToPartnerInfoArrayOutputWithContext(ctx context.Context) PartnerInfoArrayOutput {
	return o
}

func (o PartnerInfoArrayOutput) Index(i pulumi.IntInput) PartnerInfoOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PartnerInfo {
		return vs[0].([]PartnerInfo)[vs[1].(int)]
	}).(PartnerInfoOutput)
}

type PartnerInfoResponse struct {
	Id              string `pulumi:"id"`
	Location        string `pulumi:"location"`
	ReplicationRole string `pulumi:"replicationRole"`
}





type PartnerInfoResponseInput interface {
	pulumi.Input

	ToPartnerInfoResponseOutput() PartnerInfoResponseOutput
	ToPartnerInfoResponseOutputWithContext(context.Context) PartnerInfoResponseOutput
}

type PartnerInfoResponseArgs struct {
	Id              pulumi.StringInput `pulumi:"id"`
	Location        pulumi.StringInput `pulumi:"location"`
	ReplicationRole pulumi.StringInput `pulumi:"replicationRole"`
}

func (PartnerInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PartnerInfoResponse)(nil)).Elem()
}

func (i PartnerInfoResponseArgs) ToPartnerInfoResponseOutput() PartnerInfoResponseOutput {
	return i.ToPartnerInfoResponseOutputWithContext(context.Background())
}

func (i PartnerInfoResponseArgs) ToPartnerInfoResponseOutputWithContext(ctx context.Context) PartnerInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartnerInfoResponseOutput)
}





type PartnerInfoResponseArrayInput interface {
	pulumi.Input

	ToPartnerInfoResponseArrayOutput() PartnerInfoResponseArrayOutput
	ToPartnerInfoResponseArrayOutputWithContext(context.Context) PartnerInfoResponseArrayOutput
}

type PartnerInfoResponseArray []PartnerInfoResponseInput

func (PartnerInfoResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PartnerInfoResponse)(nil)).Elem()
}

func (i PartnerInfoResponseArray) ToPartnerInfoResponseArrayOutput() PartnerInfoResponseArrayOutput {
	return i.ToPartnerInfoResponseArrayOutputWithContext(context.Background())
}

func (i PartnerInfoResponseArray) ToPartnerInfoResponseArrayOutputWithContext(ctx context.Context) PartnerInfoResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartnerInfoResponseArrayOutput)
}

type PartnerInfoResponseOutput struct{ *pulumi.OutputState }

func (PartnerInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PartnerInfoResponse)(nil)).Elem()
}

func (o PartnerInfoResponseOutput) ToPartnerInfoResponseOutput() PartnerInfoResponseOutput {
	return o
}

func (o PartnerInfoResponseOutput) ToPartnerInfoResponseOutputWithContext(ctx context.Context) PartnerInfoResponseOutput {
	return o
}

func (o PartnerInfoResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PartnerInfoResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PartnerInfoResponseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v PartnerInfoResponse) string { return v.Location }).(pulumi.StringOutput)
}

func (o PartnerInfoResponseOutput) ReplicationRole() pulumi.StringOutput {
	return o.ApplyT(func(v PartnerInfoResponse) string { return v.ReplicationRole }).(pulumi.StringOutput)
}

type PartnerInfoResponseArrayOutput struct{ *pulumi.OutputState }

func (PartnerInfoResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PartnerInfoResponse)(nil)).Elem()
}

func (o PartnerInfoResponseArrayOutput) ToPartnerInfoResponseArrayOutput() PartnerInfoResponseArrayOutput {
	return o
}

func (o PartnerInfoResponseArrayOutput) ToPartnerInfoResponseArrayOutputWithContext(ctx context.Context) PartnerInfoResponseArrayOutput {
	return o
}

func (o PartnerInfoResponseArrayOutput) Index(i pulumi.IntInput) PartnerInfoResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PartnerInfoResponse {
		return vs[0].([]PartnerInfoResponse)[vs[1].(int)]
	}).(PartnerInfoResponseOutput)
}

type RecommendedActionErrorInfoResponse struct {
	ErrorCode   string `pulumi:"errorCode"`
	IsRetryable string `pulumi:"isRetryable"`
}





type RecommendedActionErrorInfoResponseInput interface {
	pulumi.Input

	ToRecommendedActionErrorInfoResponseOutput() RecommendedActionErrorInfoResponseOutput
	ToRecommendedActionErrorInfoResponseOutputWithContext(context.Context) RecommendedActionErrorInfoResponseOutput
}

type RecommendedActionErrorInfoResponseArgs struct {
	ErrorCode   pulumi.StringInput `pulumi:"errorCode"`
	IsRetryable pulumi.StringInput `pulumi:"isRetryable"`
}

func (RecommendedActionErrorInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RecommendedActionErrorInfoResponse)(nil)).Elem()
}

func (i RecommendedActionErrorInfoResponseArgs) ToRecommendedActionErrorInfoResponseOutput() RecommendedActionErrorInfoResponseOutput {
	return i.ToRecommendedActionErrorInfoResponseOutputWithContext(context.Background())
}

func (i RecommendedActionErrorInfoResponseArgs) ToRecommendedActionErrorInfoResponseOutputWithContext(ctx context.Context) RecommendedActionErrorInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecommendedActionErrorInfoResponseOutput)
}

type RecommendedActionErrorInfoResponseOutput struct{ *pulumi.OutputState }

func (RecommendedActionErrorInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecommendedActionErrorInfoResponse)(nil)).Elem()
}

func (o RecommendedActionErrorInfoResponseOutput) ToRecommendedActionErrorInfoResponseOutput() RecommendedActionErrorInfoResponseOutput {
	return o
}

func (o RecommendedActionErrorInfoResponseOutput) ToRecommendedActionErrorInfoResponseOutputWithContext(ctx context.Context) RecommendedActionErrorInfoResponseOutput {
	return o
}

func (o RecommendedActionErrorInfoResponseOutput) ErrorCode() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionErrorInfoResponse) string { return v.ErrorCode }).(pulumi.StringOutput)
}

func (o RecommendedActionErrorInfoResponseOutput) IsRetryable() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionErrorInfoResponse) string { return v.IsRetryable }).(pulumi.StringOutput)
}

type RecommendedActionImpactRecordResponse struct {
	AbsoluteValue       float64 `pulumi:"absoluteValue"`
	ChangeValueAbsolute float64 `pulumi:"changeValueAbsolute"`
	ChangeValueRelative float64 `pulumi:"changeValueRelative"`
	DimensionName       string  `pulumi:"dimensionName"`
	Unit                string  `pulumi:"unit"`
}





type RecommendedActionImpactRecordResponseInput interface {
	pulumi.Input

	ToRecommendedActionImpactRecordResponseOutput() RecommendedActionImpactRecordResponseOutput
	ToRecommendedActionImpactRecordResponseOutputWithContext(context.Context) RecommendedActionImpactRecordResponseOutput
}

type RecommendedActionImpactRecordResponseArgs struct {
	AbsoluteValue       pulumi.Float64Input `pulumi:"absoluteValue"`
	ChangeValueAbsolute pulumi.Float64Input `pulumi:"changeValueAbsolute"`
	ChangeValueRelative pulumi.Float64Input `pulumi:"changeValueRelative"`
	DimensionName       pulumi.StringInput  `pulumi:"dimensionName"`
	Unit                pulumi.StringInput  `pulumi:"unit"`
}

func (RecommendedActionImpactRecordResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RecommendedActionImpactRecordResponse)(nil)).Elem()
}

func (i RecommendedActionImpactRecordResponseArgs) ToRecommendedActionImpactRecordResponseOutput() RecommendedActionImpactRecordResponseOutput {
	return i.ToRecommendedActionImpactRecordResponseOutputWithContext(context.Background())
}

func (i RecommendedActionImpactRecordResponseArgs) ToRecommendedActionImpactRecordResponseOutputWithContext(ctx context.Context) RecommendedActionImpactRecordResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecommendedActionImpactRecordResponseOutput)
}





type RecommendedActionImpactRecordResponseArrayInput interface {
	pulumi.Input

	ToRecommendedActionImpactRecordResponseArrayOutput() RecommendedActionImpactRecordResponseArrayOutput
	ToRecommendedActionImpactRecordResponseArrayOutputWithContext(context.Context) RecommendedActionImpactRecordResponseArrayOutput
}

type RecommendedActionImpactRecordResponseArray []RecommendedActionImpactRecordResponseInput

func (RecommendedActionImpactRecordResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RecommendedActionImpactRecordResponse)(nil)).Elem()
}

func (i RecommendedActionImpactRecordResponseArray) ToRecommendedActionImpactRecordResponseArrayOutput() RecommendedActionImpactRecordResponseArrayOutput {
	return i.ToRecommendedActionImpactRecordResponseArrayOutputWithContext(context.Background())
}

func (i RecommendedActionImpactRecordResponseArray) ToRecommendedActionImpactRecordResponseArrayOutputWithContext(ctx context.Context) RecommendedActionImpactRecordResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecommendedActionImpactRecordResponseArrayOutput)
}

type RecommendedActionImpactRecordResponseOutput struct{ *pulumi.OutputState }

func (RecommendedActionImpactRecordResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecommendedActionImpactRecordResponse)(nil)).Elem()
}

func (o RecommendedActionImpactRecordResponseOutput) ToRecommendedActionImpactRecordResponseOutput() RecommendedActionImpactRecordResponseOutput {
	return o
}

func (o RecommendedActionImpactRecordResponseOutput) ToRecommendedActionImpactRecordResponseOutputWithContext(ctx context.Context) RecommendedActionImpactRecordResponseOutput {
	return o
}

func (o RecommendedActionImpactRecordResponseOutput) AbsoluteValue() pulumi.Float64Output {
	return o.ApplyT(func(v RecommendedActionImpactRecordResponse) float64 { return v.AbsoluteValue }).(pulumi.Float64Output)
}

func (o RecommendedActionImpactRecordResponseOutput) ChangeValueAbsolute() pulumi.Float64Output {
	return o.ApplyT(func(v RecommendedActionImpactRecordResponse) float64 { return v.ChangeValueAbsolute }).(pulumi.Float64Output)
}

func (o RecommendedActionImpactRecordResponseOutput) ChangeValueRelative() pulumi.Float64Output {
	return o.ApplyT(func(v RecommendedActionImpactRecordResponse) float64 { return v.ChangeValueRelative }).(pulumi.Float64Output)
}

func (o RecommendedActionImpactRecordResponseOutput) DimensionName() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionImpactRecordResponse) string { return v.DimensionName }).(pulumi.StringOutput)
}

func (o RecommendedActionImpactRecordResponseOutput) Unit() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionImpactRecordResponse) string { return v.Unit }).(pulumi.StringOutput)
}

type RecommendedActionImpactRecordResponseArrayOutput struct{ *pulumi.OutputState }

func (RecommendedActionImpactRecordResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RecommendedActionImpactRecordResponse)(nil)).Elem()
}

func (o RecommendedActionImpactRecordResponseArrayOutput) ToRecommendedActionImpactRecordResponseArrayOutput() RecommendedActionImpactRecordResponseArrayOutput {
	return o
}

func (o RecommendedActionImpactRecordResponseArrayOutput) ToRecommendedActionImpactRecordResponseArrayOutputWithContext(ctx context.Context) RecommendedActionImpactRecordResponseArrayOutput {
	return o
}

func (o RecommendedActionImpactRecordResponseArrayOutput) Index(i pulumi.IntInput) RecommendedActionImpactRecordResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RecommendedActionImpactRecordResponse {
		return vs[0].([]RecommendedActionImpactRecordResponse)[vs[1].(int)]
	}).(RecommendedActionImpactRecordResponseOutput)
}

type RecommendedActionImplementationInfoResponse struct {
	Method string `pulumi:"method"`
	Script string `pulumi:"script"`
}





type RecommendedActionImplementationInfoResponseInput interface {
	pulumi.Input

	ToRecommendedActionImplementationInfoResponseOutput() RecommendedActionImplementationInfoResponseOutput
	ToRecommendedActionImplementationInfoResponseOutputWithContext(context.Context) RecommendedActionImplementationInfoResponseOutput
}

type RecommendedActionImplementationInfoResponseArgs struct {
	Method pulumi.StringInput `pulumi:"method"`
	Script pulumi.StringInput `pulumi:"script"`
}

func (RecommendedActionImplementationInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RecommendedActionImplementationInfoResponse)(nil)).Elem()
}

func (i RecommendedActionImplementationInfoResponseArgs) ToRecommendedActionImplementationInfoResponseOutput() RecommendedActionImplementationInfoResponseOutput {
	return i.ToRecommendedActionImplementationInfoResponseOutputWithContext(context.Background())
}

func (i RecommendedActionImplementationInfoResponseArgs) ToRecommendedActionImplementationInfoResponseOutputWithContext(ctx context.Context) RecommendedActionImplementationInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecommendedActionImplementationInfoResponseOutput)
}

type RecommendedActionImplementationInfoResponseOutput struct{ *pulumi.OutputState }

func (RecommendedActionImplementationInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecommendedActionImplementationInfoResponse)(nil)).Elem()
}

func (o RecommendedActionImplementationInfoResponseOutput) ToRecommendedActionImplementationInfoResponseOutput() RecommendedActionImplementationInfoResponseOutput {
	return o
}

func (o RecommendedActionImplementationInfoResponseOutput) ToRecommendedActionImplementationInfoResponseOutputWithContext(ctx context.Context) RecommendedActionImplementationInfoResponseOutput {
	return o
}

func (o RecommendedActionImplementationInfoResponseOutput) Method() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionImplementationInfoResponse) string { return v.Method }).(pulumi.StringOutput)
}

func (o RecommendedActionImplementationInfoResponseOutput) Script() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionImplementationInfoResponse) string { return v.Script }).(pulumi.StringOutput)
}

type RecommendedActionMetricInfoResponse struct {
	MetricName string  `pulumi:"metricName"`
	StartTime  string  `pulumi:"startTime"`
	TimeGrain  string  `pulumi:"timeGrain"`
	Unit       string  `pulumi:"unit"`
	Value      float64 `pulumi:"value"`
}





type RecommendedActionMetricInfoResponseInput interface {
	pulumi.Input

	ToRecommendedActionMetricInfoResponseOutput() RecommendedActionMetricInfoResponseOutput
	ToRecommendedActionMetricInfoResponseOutputWithContext(context.Context) RecommendedActionMetricInfoResponseOutput
}

type RecommendedActionMetricInfoResponseArgs struct {
	MetricName pulumi.StringInput  `pulumi:"metricName"`
	StartTime  pulumi.StringInput  `pulumi:"startTime"`
	TimeGrain  pulumi.StringInput  `pulumi:"timeGrain"`
	Unit       pulumi.StringInput  `pulumi:"unit"`
	Value      pulumi.Float64Input `pulumi:"value"`
}

func (RecommendedActionMetricInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RecommendedActionMetricInfoResponse)(nil)).Elem()
}

func (i RecommendedActionMetricInfoResponseArgs) ToRecommendedActionMetricInfoResponseOutput() RecommendedActionMetricInfoResponseOutput {
	return i.ToRecommendedActionMetricInfoResponseOutputWithContext(context.Background())
}

func (i RecommendedActionMetricInfoResponseArgs) ToRecommendedActionMetricInfoResponseOutputWithContext(ctx context.Context) RecommendedActionMetricInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecommendedActionMetricInfoResponseOutput)
}





type RecommendedActionMetricInfoResponseArrayInput interface {
	pulumi.Input

	ToRecommendedActionMetricInfoResponseArrayOutput() RecommendedActionMetricInfoResponseArrayOutput
	ToRecommendedActionMetricInfoResponseArrayOutputWithContext(context.Context) RecommendedActionMetricInfoResponseArrayOutput
}

type RecommendedActionMetricInfoResponseArray []RecommendedActionMetricInfoResponseInput

func (RecommendedActionMetricInfoResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RecommendedActionMetricInfoResponse)(nil)).Elem()
}

func (i RecommendedActionMetricInfoResponseArray) ToRecommendedActionMetricInfoResponseArrayOutput() RecommendedActionMetricInfoResponseArrayOutput {
	return i.ToRecommendedActionMetricInfoResponseArrayOutputWithContext(context.Background())
}

func (i RecommendedActionMetricInfoResponseArray) ToRecommendedActionMetricInfoResponseArrayOutputWithContext(ctx context.Context) RecommendedActionMetricInfoResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecommendedActionMetricInfoResponseArrayOutput)
}

type RecommendedActionMetricInfoResponseOutput struct{ *pulumi.OutputState }

func (RecommendedActionMetricInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecommendedActionMetricInfoResponse)(nil)).Elem()
}

func (o RecommendedActionMetricInfoResponseOutput) ToRecommendedActionMetricInfoResponseOutput() RecommendedActionMetricInfoResponseOutput {
	return o
}

func (o RecommendedActionMetricInfoResponseOutput) ToRecommendedActionMetricInfoResponseOutputWithContext(ctx context.Context) RecommendedActionMetricInfoResponseOutput {
	return o
}

func (o RecommendedActionMetricInfoResponseOutput) MetricName() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionMetricInfoResponse) string { return v.MetricName }).(pulumi.StringOutput)
}

func (o RecommendedActionMetricInfoResponseOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionMetricInfoResponse) string { return v.StartTime }).(pulumi.StringOutput)
}

func (o RecommendedActionMetricInfoResponseOutput) TimeGrain() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionMetricInfoResponse) string { return v.TimeGrain }).(pulumi.StringOutput)
}

func (o RecommendedActionMetricInfoResponseOutput) Unit() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionMetricInfoResponse) string { return v.Unit }).(pulumi.StringOutput)
}

func (o RecommendedActionMetricInfoResponseOutput) Value() pulumi.Float64Output {
	return o.ApplyT(func(v RecommendedActionMetricInfoResponse) float64 { return v.Value }).(pulumi.Float64Output)
}

type RecommendedActionMetricInfoResponseArrayOutput struct{ *pulumi.OutputState }

func (RecommendedActionMetricInfoResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RecommendedActionMetricInfoResponse)(nil)).Elem()
}

func (o RecommendedActionMetricInfoResponseArrayOutput) ToRecommendedActionMetricInfoResponseArrayOutput() RecommendedActionMetricInfoResponseArrayOutput {
	return o
}

func (o RecommendedActionMetricInfoResponseArrayOutput) ToRecommendedActionMetricInfoResponseArrayOutputWithContext(ctx context.Context) RecommendedActionMetricInfoResponseArrayOutput {
	return o
}

func (o RecommendedActionMetricInfoResponseArrayOutput) Index(i pulumi.IntInput) RecommendedActionMetricInfoResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RecommendedActionMetricInfoResponse {
		return vs[0].([]RecommendedActionMetricInfoResponse)[vs[1].(int)]
	}).(RecommendedActionMetricInfoResponseOutput)
}

type RecommendedActionResponse struct {
	Details                    map[string]interface{}                      `pulumi:"details"`
	ErrorDetails               RecommendedActionErrorInfoResponse          `pulumi:"errorDetails"`
	EstimatedImpact            []RecommendedActionImpactRecordResponse     `pulumi:"estimatedImpact"`
	ExecuteActionDuration      string                                      `pulumi:"executeActionDuration"`
	ExecuteActionInitiatedBy   string                                      `pulumi:"executeActionInitiatedBy"`
	ExecuteActionInitiatedTime string                                      `pulumi:"executeActionInitiatedTime"`
	ExecuteActionStartTime     string                                      `pulumi:"executeActionStartTime"`
	Id                         string                                      `pulumi:"id"`
	ImplementationDetails      RecommendedActionImplementationInfoResponse `pulumi:"implementationDetails"`
	IsArchivedAction           bool                                        `pulumi:"isArchivedAction"`
	IsExecutableAction         bool                                        `pulumi:"isExecutableAction"`
	IsRevertableAction         bool                                        `pulumi:"isRevertableAction"`
	Kind                       string                                      `pulumi:"kind"`
	LastRefresh                string                                      `pulumi:"lastRefresh"`
	LinkedObjects              []string                                    `pulumi:"linkedObjects"`
	Location                   string                                      `pulumi:"location"`
	Name                       string                                      `pulumi:"name"`
	ObservedImpact             []RecommendedActionImpactRecordResponse     `pulumi:"observedImpact"`
	RecommendationReason       string                                      `pulumi:"recommendationReason"`
	RevertActionDuration       string                                      `pulumi:"revertActionDuration"`
	RevertActionInitiatedBy    string                                      `pulumi:"revertActionInitiatedBy"`
	RevertActionInitiatedTime  string                                      `pulumi:"revertActionInitiatedTime"`
	RevertActionStartTime      string                                      `pulumi:"revertActionStartTime"`
	Score                      int                                         `pulumi:"score"`
	State                      RecommendedActionStateInfoResponse          `pulumi:"state"`
	TimeSeries                 []RecommendedActionMetricInfoResponse       `pulumi:"timeSeries"`
	Type                       string                                      `pulumi:"type"`
	ValidSince                 string                                      `pulumi:"validSince"`
}





type RecommendedActionResponseInput interface {
	pulumi.Input

	ToRecommendedActionResponseOutput() RecommendedActionResponseOutput
	ToRecommendedActionResponseOutputWithContext(context.Context) RecommendedActionResponseOutput
}

type RecommendedActionResponseArgs struct {
	Details                    pulumi.MapInput                                  `pulumi:"details"`
	ErrorDetails               RecommendedActionErrorInfoResponseInput          `pulumi:"errorDetails"`
	EstimatedImpact            RecommendedActionImpactRecordResponseArrayInput  `pulumi:"estimatedImpact"`
	ExecuteActionDuration      pulumi.StringInput                               `pulumi:"executeActionDuration"`
	ExecuteActionInitiatedBy   pulumi.StringInput                               `pulumi:"executeActionInitiatedBy"`
	ExecuteActionInitiatedTime pulumi.StringInput                               `pulumi:"executeActionInitiatedTime"`
	ExecuteActionStartTime     pulumi.StringInput                               `pulumi:"executeActionStartTime"`
	Id                         pulumi.StringInput                               `pulumi:"id"`
	ImplementationDetails      RecommendedActionImplementationInfoResponseInput `pulumi:"implementationDetails"`
	IsArchivedAction           pulumi.BoolInput                                 `pulumi:"isArchivedAction"`
	IsExecutableAction         pulumi.BoolInput                                 `pulumi:"isExecutableAction"`
	IsRevertableAction         pulumi.BoolInput                                 `pulumi:"isRevertableAction"`
	Kind                       pulumi.StringInput                               `pulumi:"kind"`
	LastRefresh                pulumi.StringInput                               `pulumi:"lastRefresh"`
	LinkedObjects              pulumi.StringArrayInput                          `pulumi:"linkedObjects"`
	Location                   pulumi.StringInput                               `pulumi:"location"`
	Name                       pulumi.StringInput                               `pulumi:"name"`
	ObservedImpact             RecommendedActionImpactRecordResponseArrayInput  `pulumi:"observedImpact"`
	RecommendationReason       pulumi.StringInput                               `pulumi:"recommendationReason"`
	RevertActionDuration       pulumi.StringInput                               `pulumi:"revertActionDuration"`
	RevertActionInitiatedBy    pulumi.StringInput                               `pulumi:"revertActionInitiatedBy"`
	RevertActionInitiatedTime  pulumi.StringInput                               `pulumi:"revertActionInitiatedTime"`
	RevertActionStartTime      pulumi.StringInput                               `pulumi:"revertActionStartTime"`
	Score                      pulumi.IntInput                                  `pulumi:"score"`
	State                      RecommendedActionStateInfoResponseInput          `pulumi:"state"`
	TimeSeries                 RecommendedActionMetricInfoResponseArrayInput    `pulumi:"timeSeries"`
	Type                       pulumi.StringInput                               `pulumi:"type"`
	ValidSince                 pulumi.StringInput                               `pulumi:"validSince"`
}

func (RecommendedActionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RecommendedActionResponse)(nil)).Elem()
}

func (i RecommendedActionResponseArgs) ToRecommendedActionResponseOutput() RecommendedActionResponseOutput {
	return i.ToRecommendedActionResponseOutputWithContext(context.Background())
}

func (i RecommendedActionResponseArgs) ToRecommendedActionResponseOutputWithContext(ctx context.Context) RecommendedActionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecommendedActionResponseOutput)
}





type RecommendedActionResponseArrayInput interface {
	pulumi.Input

	ToRecommendedActionResponseArrayOutput() RecommendedActionResponseArrayOutput
	ToRecommendedActionResponseArrayOutputWithContext(context.Context) RecommendedActionResponseArrayOutput
}

type RecommendedActionResponseArray []RecommendedActionResponseInput

func (RecommendedActionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RecommendedActionResponse)(nil)).Elem()
}

func (i RecommendedActionResponseArray) ToRecommendedActionResponseArrayOutput() RecommendedActionResponseArrayOutput {
	return i.ToRecommendedActionResponseArrayOutputWithContext(context.Background())
}

func (i RecommendedActionResponseArray) ToRecommendedActionResponseArrayOutputWithContext(ctx context.Context) RecommendedActionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecommendedActionResponseArrayOutput)
}

type RecommendedActionResponseOutput struct{ *pulumi.OutputState }

func (RecommendedActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecommendedActionResponse)(nil)).Elem()
}

func (o RecommendedActionResponseOutput) ToRecommendedActionResponseOutput() RecommendedActionResponseOutput {
	return o
}

func (o RecommendedActionResponseOutput) ToRecommendedActionResponseOutputWithContext(ctx context.Context) RecommendedActionResponseOutput {
	return o
}

func (o RecommendedActionResponseOutput) Details() pulumi.MapOutput {
	return o.ApplyT(func(v RecommendedActionResponse) map[string]interface{} { return v.Details }).(pulumi.MapOutput)
}

func (o RecommendedActionResponseOutput) ErrorDetails() RecommendedActionErrorInfoResponseOutput {
	return o.ApplyT(func(v RecommendedActionResponse) RecommendedActionErrorInfoResponse { return v.ErrorDetails }).(RecommendedActionErrorInfoResponseOutput)
}

func (o RecommendedActionResponseOutput) EstimatedImpact() RecommendedActionImpactRecordResponseArrayOutput {
	return o.ApplyT(func(v RecommendedActionResponse) []RecommendedActionImpactRecordResponse { return v.EstimatedImpact }).(RecommendedActionImpactRecordResponseArrayOutput)
}

func (o RecommendedActionResponseOutput) ExecuteActionDuration() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionResponse) string { return v.ExecuteActionDuration }).(pulumi.StringOutput)
}

func (o RecommendedActionResponseOutput) ExecuteActionInitiatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionResponse) string { return v.ExecuteActionInitiatedBy }).(pulumi.StringOutput)
}

func (o RecommendedActionResponseOutput) ExecuteActionInitiatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionResponse) string { return v.ExecuteActionInitiatedTime }).(pulumi.StringOutput)
}

func (o RecommendedActionResponseOutput) ExecuteActionStartTime() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionResponse) string { return v.ExecuteActionStartTime }).(pulumi.StringOutput)
}

func (o RecommendedActionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o RecommendedActionResponseOutput) ImplementationDetails() RecommendedActionImplementationInfoResponseOutput {
	return o.ApplyT(func(v RecommendedActionResponse) RecommendedActionImplementationInfoResponse {
		return v.ImplementationDetails
	}).(RecommendedActionImplementationInfoResponseOutput)
}

func (o RecommendedActionResponseOutput) IsArchivedAction() pulumi.BoolOutput {
	return o.ApplyT(func(v RecommendedActionResponse) bool { return v.IsArchivedAction }).(pulumi.BoolOutput)
}

func (o RecommendedActionResponseOutput) IsExecutableAction() pulumi.BoolOutput {
	return o.ApplyT(func(v RecommendedActionResponse) bool { return v.IsExecutableAction }).(pulumi.BoolOutput)
}

func (o RecommendedActionResponseOutput) IsRevertableAction() pulumi.BoolOutput {
	return o.ApplyT(func(v RecommendedActionResponse) bool { return v.IsRevertableAction }).(pulumi.BoolOutput)
}

func (o RecommendedActionResponseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionResponse) string { return v.Kind }).(pulumi.StringOutput)
}

func (o RecommendedActionResponseOutput) LastRefresh() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionResponse) string { return v.LastRefresh }).(pulumi.StringOutput)
}

func (o RecommendedActionResponseOutput) LinkedObjects() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RecommendedActionResponse) []string { return v.LinkedObjects }).(pulumi.StringArrayOutput)
}

func (o RecommendedActionResponseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionResponse) string { return v.Location }).(pulumi.StringOutput)
}

func (o RecommendedActionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o RecommendedActionResponseOutput) ObservedImpact() RecommendedActionImpactRecordResponseArrayOutput {
	return o.ApplyT(func(v RecommendedActionResponse) []RecommendedActionImpactRecordResponse { return v.ObservedImpact }).(RecommendedActionImpactRecordResponseArrayOutput)
}

func (o RecommendedActionResponseOutput) RecommendationReason() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionResponse) string { return v.RecommendationReason }).(pulumi.StringOutput)
}

func (o RecommendedActionResponseOutput) RevertActionDuration() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionResponse) string { return v.RevertActionDuration }).(pulumi.StringOutput)
}

func (o RecommendedActionResponseOutput) RevertActionInitiatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionResponse) string { return v.RevertActionInitiatedBy }).(pulumi.StringOutput)
}

func (o RecommendedActionResponseOutput) RevertActionInitiatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionResponse) string { return v.RevertActionInitiatedTime }).(pulumi.StringOutput)
}

func (o RecommendedActionResponseOutput) RevertActionStartTime() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionResponse) string { return v.RevertActionStartTime }).(pulumi.StringOutput)
}

func (o RecommendedActionResponseOutput) Score() pulumi.IntOutput {
	return o.ApplyT(func(v RecommendedActionResponse) int { return v.Score }).(pulumi.IntOutput)
}

func (o RecommendedActionResponseOutput) State() RecommendedActionStateInfoResponseOutput {
	return o.ApplyT(func(v RecommendedActionResponse) RecommendedActionStateInfoResponse { return v.State }).(RecommendedActionStateInfoResponseOutput)
}

func (o RecommendedActionResponseOutput) TimeSeries() RecommendedActionMetricInfoResponseArrayOutput {
	return o.ApplyT(func(v RecommendedActionResponse) []RecommendedActionMetricInfoResponse { return v.TimeSeries }).(RecommendedActionMetricInfoResponseArrayOutput)
}

func (o RecommendedActionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o RecommendedActionResponseOutput) ValidSince() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionResponse) string { return v.ValidSince }).(pulumi.StringOutput)
}

type RecommendedActionResponseArrayOutput struct{ *pulumi.OutputState }

func (RecommendedActionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RecommendedActionResponse)(nil)).Elem()
}

func (o RecommendedActionResponseArrayOutput) ToRecommendedActionResponseArrayOutput() RecommendedActionResponseArrayOutput {
	return o
}

func (o RecommendedActionResponseArrayOutput) ToRecommendedActionResponseArrayOutputWithContext(ctx context.Context) RecommendedActionResponseArrayOutput {
	return o
}

func (o RecommendedActionResponseArrayOutput) Index(i pulumi.IntInput) RecommendedActionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RecommendedActionResponse {
		return vs[0].([]RecommendedActionResponse)[vs[1].(int)]
	}).(RecommendedActionResponseOutput)
}

type RecommendedActionStateInfoResponse struct {
	ActionInitiatedBy string `pulumi:"actionInitiatedBy"`
	CurrentValue      string `pulumi:"currentValue"`
	LastModified      string `pulumi:"lastModified"`
}





type RecommendedActionStateInfoResponseInput interface {
	pulumi.Input

	ToRecommendedActionStateInfoResponseOutput() RecommendedActionStateInfoResponseOutput
	ToRecommendedActionStateInfoResponseOutputWithContext(context.Context) RecommendedActionStateInfoResponseOutput
}

type RecommendedActionStateInfoResponseArgs struct {
	ActionInitiatedBy pulumi.StringInput `pulumi:"actionInitiatedBy"`
	CurrentValue      pulumi.StringInput `pulumi:"currentValue"`
	LastModified      pulumi.StringInput `pulumi:"lastModified"`
}

func (RecommendedActionStateInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RecommendedActionStateInfoResponse)(nil)).Elem()
}

func (i RecommendedActionStateInfoResponseArgs) ToRecommendedActionStateInfoResponseOutput() RecommendedActionStateInfoResponseOutput {
	return i.ToRecommendedActionStateInfoResponseOutputWithContext(context.Background())
}

func (i RecommendedActionStateInfoResponseArgs) ToRecommendedActionStateInfoResponseOutputWithContext(ctx context.Context) RecommendedActionStateInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecommendedActionStateInfoResponseOutput)
}

type RecommendedActionStateInfoResponseOutput struct{ *pulumi.OutputState }

func (RecommendedActionStateInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecommendedActionStateInfoResponse)(nil)).Elem()
}

func (o RecommendedActionStateInfoResponseOutput) ToRecommendedActionStateInfoResponseOutput() RecommendedActionStateInfoResponseOutput {
	return o
}

func (o RecommendedActionStateInfoResponseOutput) ToRecommendedActionStateInfoResponseOutputWithContext(ctx context.Context) RecommendedActionStateInfoResponseOutput {
	return o
}

func (o RecommendedActionStateInfoResponseOutput) ActionInitiatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionStateInfoResponse) string { return v.ActionInitiatedBy }).(pulumi.StringOutput)
}

func (o RecommendedActionStateInfoResponseOutput) CurrentValue() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionStateInfoResponse) string { return v.CurrentValue }).(pulumi.StringOutput)
}

func (o RecommendedActionStateInfoResponseOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendedActionStateInfoResponse) string { return v.LastModified }).(pulumi.StringOutput)
}

type ResourceIdentity struct {
	Type *string `pulumi:"type"`
}





type ResourceIdentityInput interface {
	pulumi.Input

	ToResourceIdentityOutput() ResourceIdentityOutput
	ToResourceIdentityOutputWithContext(context.Context) ResourceIdentityOutput
}

type ResourceIdentityArgs struct {
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (ResourceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentity)(nil)).Elem()
}

func (i ResourceIdentityArgs) ToResourceIdentityOutput() ResourceIdentityOutput {
	return i.ToResourceIdentityOutputWithContext(context.Background())
}

func (i ResourceIdentityArgs) ToResourceIdentityOutputWithContext(ctx context.Context) ResourceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityOutput)
}

func (i ResourceIdentityArgs) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return i.ToResourceIdentityPtrOutputWithContext(context.Background())
}

func (i ResourceIdentityArgs) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityOutput).ToResourceIdentityPtrOutputWithContext(ctx)
}









type ResourceIdentityPtrInput interface {
	pulumi.Input

	ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput
	ToResourceIdentityPtrOutputWithContext(context.Context) ResourceIdentityPtrOutput
}

type resourceIdentityPtrType ResourceIdentityArgs

func ResourceIdentityPtr(v *ResourceIdentityArgs) ResourceIdentityPtrInput {
	return (*resourceIdentityPtrType)(v)
}

func (*resourceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentity)(nil)).Elem()
}

func (i *resourceIdentityPtrType) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return i.ToResourceIdentityPtrOutputWithContext(context.Background())
}

func (i *resourceIdentityPtrType) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityPtrOutput)
}

type ResourceIdentityOutput struct{ *pulumi.OutputState }

func (ResourceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentity)(nil)).Elem()
}

func (o ResourceIdentityOutput) ToResourceIdentityOutput() ResourceIdentityOutput {
	return o
}

func (o ResourceIdentityOutput) ToResourceIdentityOutputWithContext(ctx context.Context) ResourceIdentityOutput {
	return o
}

func (o ResourceIdentityOutput) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return o.ToResourceIdentityPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityOutput) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentity) *ResourceIdentity {
		return &v
	}).(ResourceIdentityPtrOutput)
}

func (o ResourceIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ResourceIdentityPtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentity)(nil)).Elem()
}

func (o ResourceIdentityPtrOutput) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return o
}

func (o ResourceIdentityPtrOutput) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return o
}

func (o ResourceIdentityPtrOutput) Elem() ResourceIdentityOutput {
	return o.ApplyT(func(v *ResourceIdentity) ResourceIdentity {
		if v != nil {
			return *v
		}
		var ret ResourceIdentity
		return ret
	}).(ResourceIdentityOutput)
}

func (o ResourceIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ResourceIdentityResponse struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}





type ResourceIdentityResponseInput interface {
	pulumi.Input

	ToResourceIdentityResponseOutput() ResourceIdentityResponseOutput
	ToResourceIdentityResponseOutputWithContext(context.Context) ResourceIdentityResponseOutput
}

type ResourceIdentityResponseArgs struct {
	PrincipalId pulumi.StringInput    `pulumi:"principalId"`
	TenantId    pulumi.StringInput    `pulumi:"tenantId"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
}

func (ResourceIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityResponse)(nil)).Elem()
}

func (i ResourceIdentityResponseArgs) ToResourceIdentityResponseOutput() ResourceIdentityResponseOutput {
	return i.ToResourceIdentityResponseOutputWithContext(context.Background())
}

func (i ResourceIdentityResponseArgs) ToResourceIdentityResponseOutputWithContext(ctx context.Context) ResourceIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityResponseOutput)
}

func (i ResourceIdentityResponseArgs) ToResourceIdentityResponsePtrOutput() ResourceIdentityResponsePtrOutput {
	return i.ToResourceIdentityResponsePtrOutputWithContext(context.Background())
}

func (i ResourceIdentityResponseArgs) ToResourceIdentityResponsePtrOutputWithContext(ctx context.Context) ResourceIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityResponseOutput).ToResourceIdentityResponsePtrOutputWithContext(ctx)
}









type ResourceIdentityResponsePtrInput interface {
	pulumi.Input

	ToResourceIdentityResponsePtrOutput() ResourceIdentityResponsePtrOutput
	ToResourceIdentityResponsePtrOutputWithContext(context.Context) ResourceIdentityResponsePtrOutput
}

type resourceIdentityResponsePtrType ResourceIdentityResponseArgs

func ResourceIdentityResponsePtr(v *ResourceIdentityResponseArgs) ResourceIdentityResponsePtrInput {
	return (*resourceIdentityResponsePtrType)(v)
}

func (*resourceIdentityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentityResponse)(nil)).Elem()
}

func (i *resourceIdentityResponsePtrType) ToResourceIdentityResponsePtrOutput() ResourceIdentityResponsePtrOutput {
	return i.ToResourceIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *resourceIdentityResponsePtrType) ToResourceIdentityResponsePtrOutputWithContext(ctx context.Context) ResourceIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityResponsePtrOutput)
}

type ResourceIdentityResponseOutput struct{ *pulumi.OutputState }

func (ResourceIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityResponse)(nil)).Elem()
}

func (o ResourceIdentityResponseOutput) ToResourceIdentityResponseOutput() ResourceIdentityResponseOutput {
	return o
}

func (o ResourceIdentityResponseOutput) ToResourceIdentityResponseOutputWithContext(ctx context.Context) ResourceIdentityResponseOutput {
	return o
}

func (o ResourceIdentityResponseOutput) ToResourceIdentityResponsePtrOutput() ResourceIdentityResponsePtrOutput {
	return o.ToResourceIdentityResponsePtrOutputWithContext(context.Background())
}

func (o ResourceIdentityResponseOutput) ToResourceIdentityResponsePtrOutputWithContext(ctx context.Context) ResourceIdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentityResponse) *ResourceIdentityResponse {
		return &v
	}).(ResourceIdentityResponsePtrOutput)
}

func (o ResourceIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ResourceIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ResourceIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ResourceIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentityResponse)(nil)).Elem()
}

func (o ResourceIdentityResponsePtrOutput) ToResourceIdentityResponsePtrOutput() ResourceIdentityResponsePtrOutput {
	return o
}

func (o ResourceIdentityResponsePtrOutput) ToResourceIdentityResponsePtrOutputWithContext(ctx context.Context) ResourceIdentityResponsePtrOutput {
	return o
}

func (o ResourceIdentityResponsePtrOutput) Elem() ResourceIdentityResponseOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) ResourceIdentityResponse {
		if v != nil {
			return *v
		}
		var ret ResourceIdentityResponse
		return ret
	}).(ResourceIdentityResponseOutput)
}

func (o ResourceIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ResourceIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ResourceIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type Sku struct {
	Capacity *int    `pulumi:"capacity"`
	Family   *string `pulumi:"family"`
	Name     string  `pulumi:"name"`
	Size     *string `pulumi:"size"`
	Tier     *string `pulumi:"tier"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Family   pulumi.StringPtrInput `pulumi:"family"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Size     pulumi.StringPtrInput `pulumi:"size"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}









type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
}

type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
}

func (o SkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Sku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku {
		if v != nil {
			return *v
		}
		var ret Sku
		return ret
	}).(SkuOutput)
}

func (o SkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Sku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuPtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Capacity *int    `pulumi:"capacity"`
	Family   *string `pulumi:"family"`
	Name     string  `pulumi:"name"`
	Size     *string `pulumi:"size"`
	Tier     *string `pulumi:"tier"`
}





type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

type SkuResponseArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Family   pulumi.StringPtrInput `pulumi:"family"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Size     pulumi.StringPtrInput `pulumi:"size"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
}

func (SkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (i SkuResponseArgs) ToSkuResponseOutput() SkuResponseOutput {
	return i.ToSkuResponseOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput)
}

func (i SkuResponseArgs) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput).ToSkuResponsePtrOutputWithContext(ctx)
}









type SkuResponsePtrInput interface {
	pulumi.Input

	ToSkuResponsePtrOutput() SkuResponsePtrOutput
	ToSkuResponsePtrOutputWithContext(context.Context) SkuResponsePtrOutput
}

type skuResponsePtrType SkuResponseArgs

func SkuResponsePtr(v *SkuResponseArgs) SkuResponsePtrInput {
	return (*skuResponsePtrType)(v)
}

func (*skuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponsePtrOutput)
}

type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (o SkuResponseOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuResponse) *SkuResponse {
		return &v
	}).(SkuResponsePtrOutput)
}

func (o SkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuResponseOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse {
		if v != nil {
			return *v
		}
		var ret SkuResponse
		return ret
	}).(SkuResponseOutput)
}

func (o SkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuResponsePtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SyncGroupSchema struct {
	MasterSyncMemberName *string                `pulumi:"masterSyncMemberName"`
	Tables               []SyncGroupSchemaTable `pulumi:"tables"`
}





type SyncGroupSchemaInput interface {
	pulumi.Input

	ToSyncGroupSchemaOutput() SyncGroupSchemaOutput
	ToSyncGroupSchemaOutputWithContext(context.Context) SyncGroupSchemaOutput
}

type SyncGroupSchemaArgs struct {
	MasterSyncMemberName pulumi.StringPtrInput          `pulumi:"masterSyncMemberName"`
	Tables               SyncGroupSchemaTableArrayInput `pulumi:"tables"`
}

func (SyncGroupSchemaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncGroupSchema)(nil)).Elem()
}

func (i SyncGroupSchemaArgs) ToSyncGroupSchemaOutput() SyncGroupSchemaOutput {
	return i.ToSyncGroupSchemaOutputWithContext(context.Background())
}

func (i SyncGroupSchemaArgs) ToSyncGroupSchemaOutputWithContext(ctx context.Context) SyncGroupSchemaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncGroupSchemaOutput)
}

func (i SyncGroupSchemaArgs) ToSyncGroupSchemaPtrOutput() SyncGroupSchemaPtrOutput {
	return i.ToSyncGroupSchemaPtrOutputWithContext(context.Background())
}

func (i SyncGroupSchemaArgs) ToSyncGroupSchemaPtrOutputWithContext(ctx context.Context) SyncGroupSchemaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncGroupSchemaOutput).ToSyncGroupSchemaPtrOutputWithContext(ctx)
}









type SyncGroupSchemaPtrInput interface {
	pulumi.Input

	ToSyncGroupSchemaPtrOutput() SyncGroupSchemaPtrOutput
	ToSyncGroupSchemaPtrOutputWithContext(context.Context) SyncGroupSchemaPtrOutput
}

type syncGroupSchemaPtrType SyncGroupSchemaArgs

func SyncGroupSchemaPtr(v *SyncGroupSchemaArgs) SyncGroupSchemaPtrInput {
	return (*syncGroupSchemaPtrType)(v)
}

func (*syncGroupSchemaPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncGroupSchema)(nil)).Elem()
}

func (i *syncGroupSchemaPtrType) ToSyncGroupSchemaPtrOutput() SyncGroupSchemaPtrOutput {
	return i.ToSyncGroupSchemaPtrOutputWithContext(context.Background())
}

func (i *syncGroupSchemaPtrType) ToSyncGroupSchemaPtrOutputWithContext(ctx context.Context) SyncGroupSchemaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncGroupSchemaPtrOutput)
}

type SyncGroupSchemaOutput struct{ *pulumi.OutputState }

func (SyncGroupSchemaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncGroupSchema)(nil)).Elem()
}

func (o SyncGroupSchemaOutput) ToSyncGroupSchemaOutput() SyncGroupSchemaOutput {
	return o
}

func (o SyncGroupSchemaOutput) ToSyncGroupSchemaOutputWithContext(ctx context.Context) SyncGroupSchemaOutput {
	return o
}

func (o SyncGroupSchemaOutput) ToSyncGroupSchemaPtrOutput() SyncGroupSchemaPtrOutput {
	return o.ToSyncGroupSchemaPtrOutputWithContext(context.Background())
}

func (o SyncGroupSchemaOutput) ToSyncGroupSchemaPtrOutputWithContext(ctx context.Context) SyncGroupSchemaPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SyncGroupSchema) *SyncGroupSchema {
		return &v
	}).(SyncGroupSchemaPtrOutput)
}

func (o SyncGroupSchemaOutput) MasterSyncMemberName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyncGroupSchema) *string { return v.MasterSyncMemberName }).(pulumi.StringPtrOutput)
}

func (o SyncGroupSchemaOutput) Tables() SyncGroupSchemaTableArrayOutput {
	return o.ApplyT(func(v SyncGroupSchema) []SyncGroupSchemaTable { return v.Tables }).(SyncGroupSchemaTableArrayOutput)
}

type SyncGroupSchemaPtrOutput struct{ *pulumi.OutputState }

func (SyncGroupSchemaPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncGroupSchema)(nil)).Elem()
}

func (o SyncGroupSchemaPtrOutput) ToSyncGroupSchemaPtrOutput() SyncGroupSchemaPtrOutput {
	return o
}

func (o SyncGroupSchemaPtrOutput) ToSyncGroupSchemaPtrOutputWithContext(ctx context.Context) SyncGroupSchemaPtrOutput {
	return o
}

func (o SyncGroupSchemaPtrOutput) Elem() SyncGroupSchemaOutput {
	return o.ApplyT(func(v *SyncGroupSchema) SyncGroupSchema {
		if v != nil {
			return *v
		}
		var ret SyncGroupSchema
		return ret
	}).(SyncGroupSchemaOutput)
}

func (o SyncGroupSchemaPtrOutput) MasterSyncMemberName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncGroupSchema) *string {
		if v == nil {
			return nil
		}
		return v.MasterSyncMemberName
	}).(pulumi.StringPtrOutput)
}

func (o SyncGroupSchemaPtrOutput) Tables() SyncGroupSchemaTableArrayOutput {
	return o.ApplyT(func(v *SyncGroupSchema) []SyncGroupSchemaTable {
		if v == nil {
			return nil
		}
		return v.Tables
	}).(SyncGroupSchemaTableArrayOutput)
}

type SyncGroupSchemaResponse struct {
	MasterSyncMemberName *string                        `pulumi:"masterSyncMemberName"`
	Tables               []SyncGroupSchemaTableResponse `pulumi:"tables"`
}





type SyncGroupSchemaResponseInput interface {
	pulumi.Input

	ToSyncGroupSchemaResponseOutput() SyncGroupSchemaResponseOutput
	ToSyncGroupSchemaResponseOutputWithContext(context.Context) SyncGroupSchemaResponseOutput
}

type SyncGroupSchemaResponseArgs struct {
	MasterSyncMemberName pulumi.StringPtrInput                  `pulumi:"masterSyncMemberName"`
	Tables               SyncGroupSchemaTableResponseArrayInput `pulumi:"tables"`
}

func (SyncGroupSchemaResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncGroupSchemaResponse)(nil)).Elem()
}

func (i SyncGroupSchemaResponseArgs) ToSyncGroupSchemaResponseOutput() SyncGroupSchemaResponseOutput {
	return i.ToSyncGroupSchemaResponseOutputWithContext(context.Background())
}

func (i SyncGroupSchemaResponseArgs) ToSyncGroupSchemaResponseOutputWithContext(ctx context.Context) SyncGroupSchemaResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncGroupSchemaResponseOutput)
}

func (i SyncGroupSchemaResponseArgs) ToSyncGroupSchemaResponsePtrOutput() SyncGroupSchemaResponsePtrOutput {
	return i.ToSyncGroupSchemaResponsePtrOutputWithContext(context.Background())
}

func (i SyncGroupSchemaResponseArgs) ToSyncGroupSchemaResponsePtrOutputWithContext(ctx context.Context) SyncGroupSchemaResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncGroupSchemaResponseOutput).ToSyncGroupSchemaResponsePtrOutputWithContext(ctx)
}









type SyncGroupSchemaResponsePtrInput interface {
	pulumi.Input

	ToSyncGroupSchemaResponsePtrOutput() SyncGroupSchemaResponsePtrOutput
	ToSyncGroupSchemaResponsePtrOutputWithContext(context.Context) SyncGroupSchemaResponsePtrOutput
}

type syncGroupSchemaResponsePtrType SyncGroupSchemaResponseArgs

func SyncGroupSchemaResponsePtr(v *SyncGroupSchemaResponseArgs) SyncGroupSchemaResponsePtrInput {
	return (*syncGroupSchemaResponsePtrType)(v)
}

func (*syncGroupSchemaResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncGroupSchemaResponse)(nil)).Elem()
}

func (i *syncGroupSchemaResponsePtrType) ToSyncGroupSchemaResponsePtrOutput() SyncGroupSchemaResponsePtrOutput {
	return i.ToSyncGroupSchemaResponsePtrOutputWithContext(context.Background())
}

func (i *syncGroupSchemaResponsePtrType) ToSyncGroupSchemaResponsePtrOutputWithContext(ctx context.Context) SyncGroupSchemaResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncGroupSchemaResponsePtrOutput)
}

type SyncGroupSchemaResponseOutput struct{ *pulumi.OutputState }

func (SyncGroupSchemaResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncGroupSchemaResponse)(nil)).Elem()
}

func (o SyncGroupSchemaResponseOutput) ToSyncGroupSchemaResponseOutput() SyncGroupSchemaResponseOutput {
	return o
}

func (o SyncGroupSchemaResponseOutput) ToSyncGroupSchemaResponseOutputWithContext(ctx context.Context) SyncGroupSchemaResponseOutput {
	return o
}

func (o SyncGroupSchemaResponseOutput) ToSyncGroupSchemaResponsePtrOutput() SyncGroupSchemaResponsePtrOutput {
	return o.ToSyncGroupSchemaResponsePtrOutputWithContext(context.Background())
}

func (o SyncGroupSchemaResponseOutput) ToSyncGroupSchemaResponsePtrOutputWithContext(ctx context.Context) SyncGroupSchemaResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SyncGroupSchemaResponse) *SyncGroupSchemaResponse {
		return &v
	}).(SyncGroupSchemaResponsePtrOutput)
}

func (o SyncGroupSchemaResponseOutput) MasterSyncMemberName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyncGroupSchemaResponse) *string { return v.MasterSyncMemberName }).(pulumi.StringPtrOutput)
}

func (o SyncGroupSchemaResponseOutput) Tables() SyncGroupSchemaTableResponseArrayOutput {
	return o.ApplyT(func(v SyncGroupSchemaResponse) []SyncGroupSchemaTableResponse { return v.Tables }).(SyncGroupSchemaTableResponseArrayOutput)
}

type SyncGroupSchemaResponsePtrOutput struct{ *pulumi.OutputState }

func (SyncGroupSchemaResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncGroupSchemaResponse)(nil)).Elem()
}

func (o SyncGroupSchemaResponsePtrOutput) ToSyncGroupSchemaResponsePtrOutput() SyncGroupSchemaResponsePtrOutput {
	return o
}

func (o SyncGroupSchemaResponsePtrOutput) ToSyncGroupSchemaResponsePtrOutputWithContext(ctx context.Context) SyncGroupSchemaResponsePtrOutput {
	return o
}

func (o SyncGroupSchemaResponsePtrOutput) Elem() SyncGroupSchemaResponseOutput {
	return o.ApplyT(func(v *SyncGroupSchemaResponse) SyncGroupSchemaResponse {
		if v != nil {
			return *v
		}
		var ret SyncGroupSchemaResponse
		return ret
	}).(SyncGroupSchemaResponseOutput)
}

func (o SyncGroupSchemaResponsePtrOutput) MasterSyncMemberName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncGroupSchemaResponse) *string {
		if v == nil {
			return nil
		}
		return v.MasterSyncMemberName
	}).(pulumi.StringPtrOutput)
}

func (o SyncGroupSchemaResponsePtrOutput) Tables() SyncGroupSchemaTableResponseArrayOutput {
	return o.ApplyT(func(v *SyncGroupSchemaResponse) []SyncGroupSchemaTableResponse {
		if v == nil {
			return nil
		}
		return v.Tables
	}).(SyncGroupSchemaTableResponseArrayOutput)
}

type SyncGroupSchemaTable struct {
	Columns    []SyncGroupSchemaTableColumn `pulumi:"columns"`
	QuotedName *string                      `pulumi:"quotedName"`
}





type SyncGroupSchemaTableInput interface {
	pulumi.Input

	ToSyncGroupSchemaTableOutput() SyncGroupSchemaTableOutput
	ToSyncGroupSchemaTableOutputWithContext(context.Context) SyncGroupSchemaTableOutput
}

type SyncGroupSchemaTableArgs struct {
	Columns    SyncGroupSchemaTableColumnArrayInput `pulumi:"columns"`
	QuotedName pulumi.StringPtrInput                `pulumi:"quotedName"`
}

func (SyncGroupSchemaTableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncGroupSchemaTable)(nil)).Elem()
}

func (i SyncGroupSchemaTableArgs) ToSyncGroupSchemaTableOutput() SyncGroupSchemaTableOutput {
	return i.ToSyncGroupSchemaTableOutputWithContext(context.Background())
}

func (i SyncGroupSchemaTableArgs) ToSyncGroupSchemaTableOutputWithContext(ctx context.Context) SyncGroupSchemaTableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncGroupSchemaTableOutput)
}





type SyncGroupSchemaTableArrayInput interface {
	pulumi.Input

	ToSyncGroupSchemaTableArrayOutput() SyncGroupSchemaTableArrayOutput
	ToSyncGroupSchemaTableArrayOutputWithContext(context.Context) SyncGroupSchemaTableArrayOutput
}

type SyncGroupSchemaTableArray []SyncGroupSchemaTableInput

func (SyncGroupSchemaTableArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SyncGroupSchemaTable)(nil)).Elem()
}

func (i SyncGroupSchemaTableArray) ToSyncGroupSchemaTableArrayOutput() SyncGroupSchemaTableArrayOutput {
	return i.ToSyncGroupSchemaTableArrayOutputWithContext(context.Background())
}

func (i SyncGroupSchemaTableArray) ToSyncGroupSchemaTableArrayOutputWithContext(ctx context.Context) SyncGroupSchemaTableArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncGroupSchemaTableArrayOutput)
}

type SyncGroupSchemaTableOutput struct{ *pulumi.OutputState }

func (SyncGroupSchemaTableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncGroupSchemaTable)(nil)).Elem()
}

func (o SyncGroupSchemaTableOutput) ToSyncGroupSchemaTableOutput() SyncGroupSchemaTableOutput {
	return o
}

func (o SyncGroupSchemaTableOutput) ToSyncGroupSchemaTableOutputWithContext(ctx context.Context) SyncGroupSchemaTableOutput {
	return o
}

func (o SyncGroupSchemaTableOutput) Columns() SyncGroupSchemaTableColumnArrayOutput {
	return o.ApplyT(func(v SyncGroupSchemaTable) []SyncGroupSchemaTableColumn { return v.Columns }).(SyncGroupSchemaTableColumnArrayOutput)
}

func (o SyncGroupSchemaTableOutput) QuotedName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyncGroupSchemaTable) *string { return v.QuotedName }).(pulumi.StringPtrOutput)
}

type SyncGroupSchemaTableArrayOutput struct{ *pulumi.OutputState }

func (SyncGroupSchemaTableArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SyncGroupSchemaTable)(nil)).Elem()
}

func (o SyncGroupSchemaTableArrayOutput) ToSyncGroupSchemaTableArrayOutput() SyncGroupSchemaTableArrayOutput {
	return o
}

func (o SyncGroupSchemaTableArrayOutput) ToSyncGroupSchemaTableArrayOutputWithContext(ctx context.Context) SyncGroupSchemaTableArrayOutput {
	return o
}

func (o SyncGroupSchemaTableArrayOutput) Index(i pulumi.IntInput) SyncGroupSchemaTableOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SyncGroupSchemaTable {
		return vs[0].([]SyncGroupSchemaTable)[vs[1].(int)]
	}).(SyncGroupSchemaTableOutput)
}

type SyncGroupSchemaTableColumn struct {
	DataSize   *string `pulumi:"dataSize"`
	DataType   *string `pulumi:"dataType"`
	QuotedName *string `pulumi:"quotedName"`
}





type SyncGroupSchemaTableColumnInput interface {
	pulumi.Input

	ToSyncGroupSchemaTableColumnOutput() SyncGroupSchemaTableColumnOutput
	ToSyncGroupSchemaTableColumnOutputWithContext(context.Context) SyncGroupSchemaTableColumnOutput
}

type SyncGroupSchemaTableColumnArgs struct {
	DataSize   pulumi.StringPtrInput `pulumi:"dataSize"`
	DataType   pulumi.StringPtrInput `pulumi:"dataType"`
	QuotedName pulumi.StringPtrInput `pulumi:"quotedName"`
}

func (SyncGroupSchemaTableColumnArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncGroupSchemaTableColumn)(nil)).Elem()
}

func (i SyncGroupSchemaTableColumnArgs) ToSyncGroupSchemaTableColumnOutput() SyncGroupSchemaTableColumnOutput {
	return i.ToSyncGroupSchemaTableColumnOutputWithContext(context.Background())
}

func (i SyncGroupSchemaTableColumnArgs) ToSyncGroupSchemaTableColumnOutputWithContext(ctx context.Context) SyncGroupSchemaTableColumnOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncGroupSchemaTableColumnOutput)
}





type SyncGroupSchemaTableColumnArrayInput interface {
	pulumi.Input

	ToSyncGroupSchemaTableColumnArrayOutput() SyncGroupSchemaTableColumnArrayOutput
	ToSyncGroupSchemaTableColumnArrayOutputWithContext(context.Context) SyncGroupSchemaTableColumnArrayOutput
}

type SyncGroupSchemaTableColumnArray []SyncGroupSchemaTableColumnInput

func (SyncGroupSchemaTableColumnArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SyncGroupSchemaTableColumn)(nil)).Elem()
}

func (i SyncGroupSchemaTableColumnArray) ToSyncGroupSchemaTableColumnArrayOutput() SyncGroupSchemaTableColumnArrayOutput {
	return i.ToSyncGroupSchemaTableColumnArrayOutputWithContext(context.Background())
}

func (i SyncGroupSchemaTableColumnArray) ToSyncGroupSchemaTableColumnArrayOutputWithContext(ctx context.Context) SyncGroupSchemaTableColumnArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncGroupSchemaTableColumnArrayOutput)
}

type SyncGroupSchemaTableColumnOutput struct{ *pulumi.OutputState }

func (SyncGroupSchemaTableColumnOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncGroupSchemaTableColumn)(nil)).Elem()
}

func (o SyncGroupSchemaTableColumnOutput) ToSyncGroupSchemaTableColumnOutput() SyncGroupSchemaTableColumnOutput {
	return o
}

func (o SyncGroupSchemaTableColumnOutput) ToSyncGroupSchemaTableColumnOutputWithContext(ctx context.Context) SyncGroupSchemaTableColumnOutput {
	return o
}

func (o SyncGroupSchemaTableColumnOutput) DataSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyncGroupSchemaTableColumn) *string { return v.DataSize }).(pulumi.StringPtrOutput)
}

func (o SyncGroupSchemaTableColumnOutput) DataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyncGroupSchemaTableColumn) *string { return v.DataType }).(pulumi.StringPtrOutput)
}

func (o SyncGroupSchemaTableColumnOutput) QuotedName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyncGroupSchemaTableColumn) *string { return v.QuotedName }).(pulumi.StringPtrOutput)
}

type SyncGroupSchemaTableColumnArrayOutput struct{ *pulumi.OutputState }

func (SyncGroupSchemaTableColumnArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SyncGroupSchemaTableColumn)(nil)).Elem()
}

func (o SyncGroupSchemaTableColumnArrayOutput) ToSyncGroupSchemaTableColumnArrayOutput() SyncGroupSchemaTableColumnArrayOutput {
	return o
}

func (o SyncGroupSchemaTableColumnArrayOutput) ToSyncGroupSchemaTableColumnArrayOutputWithContext(ctx context.Context) SyncGroupSchemaTableColumnArrayOutput {
	return o
}

func (o SyncGroupSchemaTableColumnArrayOutput) Index(i pulumi.IntInput) SyncGroupSchemaTableColumnOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SyncGroupSchemaTableColumn {
		return vs[0].([]SyncGroupSchemaTableColumn)[vs[1].(int)]
	}).(SyncGroupSchemaTableColumnOutput)
}

type SyncGroupSchemaTableColumnResponse struct {
	DataSize   *string `pulumi:"dataSize"`
	DataType   *string `pulumi:"dataType"`
	QuotedName *string `pulumi:"quotedName"`
}





type SyncGroupSchemaTableColumnResponseInput interface {
	pulumi.Input

	ToSyncGroupSchemaTableColumnResponseOutput() SyncGroupSchemaTableColumnResponseOutput
	ToSyncGroupSchemaTableColumnResponseOutputWithContext(context.Context) SyncGroupSchemaTableColumnResponseOutput
}

type SyncGroupSchemaTableColumnResponseArgs struct {
	DataSize   pulumi.StringPtrInput `pulumi:"dataSize"`
	DataType   pulumi.StringPtrInput `pulumi:"dataType"`
	QuotedName pulumi.StringPtrInput `pulumi:"quotedName"`
}

func (SyncGroupSchemaTableColumnResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncGroupSchemaTableColumnResponse)(nil)).Elem()
}

func (i SyncGroupSchemaTableColumnResponseArgs) ToSyncGroupSchemaTableColumnResponseOutput() SyncGroupSchemaTableColumnResponseOutput {
	return i.ToSyncGroupSchemaTableColumnResponseOutputWithContext(context.Background())
}

func (i SyncGroupSchemaTableColumnResponseArgs) ToSyncGroupSchemaTableColumnResponseOutputWithContext(ctx context.Context) SyncGroupSchemaTableColumnResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncGroupSchemaTableColumnResponseOutput)
}





type SyncGroupSchemaTableColumnResponseArrayInput interface {
	pulumi.Input

	ToSyncGroupSchemaTableColumnResponseArrayOutput() SyncGroupSchemaTableColumnResponseArrayOutput
	ToSyncGroupSchemaTableColumnResponseArrayOutputWithContext(context.Context) SyncGroupSchemaTableColumnResponseArrayOutput
}

type SyncGroupSchemaTableColumnResponseArray []SyncGroupSchemaTableColumnResponseInput

func (SyncGroupSchemaTableColumnResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SyncGroupSchemaTableColumnResponse)(nil)).Elem()
}

func (i SyncGroupSchemaTableColumnResponseArray) ToSyncGroupSchemaTableColumnResponseArrayOutput() SyncGroupSchemaTableColumnResponseArrayOutput {
	return i.ToSyncGroupSchemaTableColumnResponseArrayOutputWithContext(context.Background())
}

func (i SyncGroupSchemaTableColumnResponseArray) ToSyncGroupSchemaTableColumnResponseArrayOutputWithContext(ctx context.Context) SyncGroupSchemaTableColumnResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncGroupSchemaTableColumnResponseArrayOutput)
}

type SyncGroupSchemaTableColumnResponseOutput struct{ *pulumi.OutputState }

func (SyncGroupSchemaTableColumnResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncGroupSchemaTableColumnResponse)(nil)).Elem()
}

func (o SyncGroupSchemaTableColumnResponseOutput) ToSyncGroupSchemaTableColumnResponseOutput() SyncGroupSchemaTableColumnResponseOutput {
	return o
}

func (o SyncGroupSchemaTableColumnResponseOutput) ToSyncGroupSchemaTableColumnResponseOutputWithContext(ctx context.Context) SyncGroupSchemaTableColumnResponseOutput {
	return o
}

func (o SyncGroupSchemaTableColumnResponseOutput) DataSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyncGroupSchemaTableColumnResponse) *string { return v.DataSize }).(pulumi.StringPtrOutput)
}

func (o SyncGroupSchemaTableColumnResponseOutput) DataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyncGroupSchemaTableColumnResponse) *string { return v.DataType }).(pulumi.StringPtrOutput)
}

func (o SyncGroupSchemaTableColumnResponseOutput) QuotedName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyncGroupSchemaTableColumnResponse) *string { return v.QuotedName }).(pulumi.StringPtrOutput)
}

type SyncGroupSchemaTableColumnResponseArrayOutput struct{ *pulumi.OutputState }

func (SyncGroupSchemaTableColumnResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SyncGroupSchemaTableColumnResponse)(nil)).Elem()
}

func (o SyncGroupSchemaTableColumnResponseArrayOutput) ToSyncGroupSchemaTableColumnResponseArrayOutput() SyncGroupSchemaTableColumnResponseArrayOutput {
	return o
}

func (o SyncGroupSchemaTableColumnResponseArrayOutput) ToSyncGroupSchemaTableColumnResponseArrayOutputWithContext(ctx context.Context) SyncGroupSchemaTableColumnResponseArrayOutput {
	return o
}

func (o SyncGroupSchemaTableColumnResponseArrayOutput) Index(i pulumi.IntInput) SyncGroupSchemaTableColumnResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SyncGroupSchemaTableColumnResponse {
		return vs[0].([]SyncGroupSchemaTableColumnResponse)[vs[1].(int)]
	}).(SyncGroupSchemaTableColumnResponseOutput)
}

type SyncGroupSchemaTableResponse struct {
	Columns    []SyncGroupSchemaTableColumnResponse `pulumi:"columns"`
	QuotedName *string                              `pulumi:"quotedName"`
}





type SyncGroupSchemaTableResponseInput interface {
	pulumi.Input

	ToSyncGroupSchemaTableResponseOutput() SyncGroupSchemaTableResponseOutput
	ToSyncGroupSchemaTableResponseOutputWithContext(context.Context) SyncGroupSchemaTableResponseOutput
}

type SyncGroupSchemaTableResponseArgs struct {
	Columns    SyncGroupSchemaTableColumnResponseArrayInput `pulumi:"columns"`
	QuotedName pulumi.StringPtrInput                        `pulumi:"quotedName"`
}

func (SyncGroupSchemaTableResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncGroupSchemaTableResponse)(nil)).Elem()
}

func (i SyncGroupSchemaTableResponseArgs) ToSyncGroupSchemaTableResponseOutput() SyncGroupSchemaTableResponseOutput {
	return i.ToSyncGroupSchemaTableResponseOutputWithContext(context.Background())
}

func (i SyncGroupSchemaTableResponseArgs) ToSyncGroupSchemaTableResponseOutputWithContext(ctx context.Context) SyncGroupSchemaTableResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncGroupSchemaTableResponseOutput)
}





type SyncGroupSchemaTableResponseArrayInput interface {
	pulumi.Input

	ToSyncGroupSchemaTableResponseArrayOutput() SyncGroupSchemaTableResponseArrayOutput
	ToSyncGroupSchemaTableResponseArrayOutputWithContext(context.Context) SyncGroupSchemaTableResponseArrayOutput
}

type SyncGroupSchemaTableResponseArray []SyncGroupSchemaTableResponseInput

func (SyncGroupSchemaTableResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SyncGroupSchemaTableResponse)(nil)).Elem()
}

func (i SyncGroupSchemaTableResponseArray) ToSyncGroupSchemaTableResponseArrayOutput() SyncGroupSchemaTableResponseArrayOutput {
	return i.ToSyncGroupSchemaTableResponseArrayOutputWithContext(context.Background())
}

func (i SyncGroupSchemaTableResponseArray) ToSyncGroupSchemaTableResponseArrayOutputWithContext(ctx context.Context) SyncGroupSchemaTableResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncGroupSchemaTableResponseArrayOutput)
}

type SyncGroupSchemaTableResponseOutput struct{ *pulumi.OutputState }

func (SyncGroupSchemaTableResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncGroupSchemaTableResponse)(nil)).Elem()
}

func (o SyncGroupSchemaTableResponseOutput) ToSyncGroupSchemaTableResponseOutput() SyncGroupSchemaTableResponseOutput {
	return o
}

func (o SyncGroupSchemaTableResponseOutput) ToSyncGroupSchemaTableResponseOutputWithContext(ctx context.Context) SyncGroupSchemaTableResponseOutput {
	return o
}

func (o SyncGroupSchemaTableResponseOutput) Columns() SyncGroupSchemaTableColumnResponseArrayOutput {
	return o.ApplyT(func(v SyncGroupSchemaTableResponse) []SyncGroupSchemaTableColumnResponse { return v.Columns }).(SyncGroupSchemaTableColumnResponseArrayOutput)
}

func (o SyncGroupSchemaTableResponseOutput) QuotedName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyncGroupSchemaTableResponse) *string { return v.QuotedName }).(pulumi.StringPtrOutput)
}

type SyncGroupSchemaTableResponseArrayOutput struct{ *pulumi.OutputState }

func (SyncGroupSchemaTableResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SyncGroupSchemaTableResponse)(nil)).Elem()
}

func (o SyncGroupSchemaTableResponseArrayOutput) ToSyncGroupSchemaTableResponseArrayOutput() SyncGroupSchemaTableResponseArrayOutput {
	return o
}

func (o SyncGroupSchemaTableResponseArrayOutput) ToSyncGroupSchemaTableResponseArrayOutputWithContext(ctx context.Context) SyncGroupSchemaTableResponseArrayOutput {
	return o
}

func (o SyncGroupSchemaTableResponseArrayOutput) Index(i pulumi.IntInput) SyncGroupSchemaTableResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SyncGroupSchemaTableResponse {
		return vs[0].([]SyncGroupSchemaTableResponse)[vs[1].(int)]
	}).(SyncGroupSchemaTableResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(FailoverGroupReadOnlyEndpointOutput{})
	pulumi.RegisterOutputType(FailoverGroupReadOnlyEndpointPtrOutput{})
	pulumi.RegisterOutputType(FailoverGroupReadOnlyEndpointResponseOutput{})
	pulumi.RegisterOutputType(FailoverGroupReadOnlyEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(FailoverGroupReadWriteEndpointOutput{})
	pulumi.RegisterOutputType(FailoverGroupReadWriteEndpointPtrOutput{})
	pulumi.RegisterOutputType(FailoverGroupReadWriteEndpointResponseOutput{})
	pulumi.RegisterOutputType(FailoverGroupReadWriteEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PartnerInfoOutput{})
	pulumi.RegisterOutputType(PartnerInfoArrayOutput{})
	pulumi.RegisterOutputType(PartnerInfoResponseOutput{})
	pulumi.RegisterOutputType(PartnerInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(RecommendedActionErrorInfoResponseOutput{})
	pulumi.RegisterOutputType(RecommendedActionImpactRecordResponseOutput{})
	pulumi.RegisterOutputType(RecommendedActionImpactRecordResponseArrayOutput{})
	pulumi.RegisterOutputType(RecommendedActionImplementationInfoResponseOutput{})
	pulumi.RegisterOutputType(RecommendedActionMetricInfoResponseOutput{})
	pulumi.RegisterOutputType(RecommendedActionMetricInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(RecommendedActionResponseOutput{})
	pulumi.RegisterOutputType(RecommendedActionResponseArrayOutput{})
	pulumi.RegisterOutputType(RecommendedActionStateInfoResponseOutput{})
	pulumi.RegisterOutputType(ResourceIdentityOutput{})
	pulumi.RegisterOutputType(ResourceIdentityPtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityResponseOutput{})
	pulumi.RegisterOutputType(ResourceIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SyncGroupSchemaOutput{})
	pulumi.RegisterOutputType(SyncGroupSchemaPtrOutput{})
	pulumi.RegisterOutputType(SyncGroupSchemaResponseOutput{})
	pulumi.RegisterOutputType(SyncGroupSchemaResponsePtrOutput{})
	pulumi.RegisterOutputType(SyncGroupSchemaTableOutput{})
	pulumi.RegisterOutputType(SyncGroupSchemaTableArrayOutput{})
	pulumi.RegisterOutputType(SyncGroupSchemaTableColumnOutput{})
	pulumi.RegisterOutputType(SyncGroupSchemaTableColumnArrayOutput{})
	pulumi.RegisterOutputType(SyncGroupSchemaTableColumnResponseOutput{})
	pulumi.RegisterOutputType(SyncGroupSchemaTableColumnResponseArrayOutput{})
	pulumi.RegisterOutputType(SyncGroupSchemaTableResponseOutput{})
	pulumi.RegisterOutputType(SyncGroupSchemaTableResponseArrayOutput{})
}
