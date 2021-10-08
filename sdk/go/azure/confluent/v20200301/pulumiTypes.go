


package v20200301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type OrganizationResourcePropertiesOfferDetail struct {
	Id          *string `pulumi:"id"`
	PlanId      *string `pulumi:"planId"`
	PlanName    *string `pulumi:"planName"`
	PublisherId *string `pulumi:"publisherId"`
	TermUnit    *string `pulumi:"termUnit"`
}





type OrganizationResourcePropertiesOfferDetailInput interface {
	pulumi.Input

	ToOrganizationResourcePropertiesOfferDetailOutput() OrganizationResourcePropertiesOfferDetailOutput
	ToOrganizationResourcePropertiesOfferDetailOutputWithContext(context.Context) OrganizationResourcePropertiesOfferDetailOutput
}

type OrganizationResourcePropertiesOfferDetailArgs struct {
	Id          pulumi.StringPtrInput `pulumi:"id"`
	PlanId      pulumi.StringPtrInput `pulumi:"planId"`
	PlanName    pulumi.StringPtrInput `pulumi:"planName"`
	PublisherId pulumi.StringPtrInput `pulumi:"publisherId"`
	TermUnit    pulumi.StringPtrInput `pulumi:"termUnit"`
}

func (OrganizationResourcePropertiesOfferDetailArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationResourcePropertiesOfferDetail)(nil)).Elem()
}

func (i OrganizationResourcePropertiesOfferDetailArgs) ToOrganizationResourcePropertiesOfferDetailOutput() OrganizationResourcePropertiesOfferDetailOutput {
	return i.ToOrganizationResourcePropertiesOfferDetailOutputWithContext(context.Background())
}

func (i OrganizationResourcePropertiesOfferDetailArgs) ToOrganizationResourcePropertiesOfferDetailOutputWithContext(ctx context.Context) OrganizationResourcePropertiesOfferDetailOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationResourcePropertiesOfferDetailOutput)
}

func (i OrganizationResourcePropertiesOfferDetailArgs) ToOrganizationResourcePropertiesOfferDetailPtrOutput() OrganizationResourcePropertiesOfferDetailPtrOutput {
	return i.ToOrganizationResourcePropertiesOfferDetailPtrOutputWithContext(context.Background())
}

func (i OrganizationResourcePropertiesOfferDetailArgs) ToOrganizationResourcePropertiesOfferDetailPtrOutputWithContext(ctx context.Context) OrganizationResourcePropertiesOfferDetailPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationResourcePropertiesOfferDetailOutput).ToOrganizationResourcePropertiesOfferDetailPtrOutputWithContext(ctx)
}









type OrganizationResourcePropertiesOfferDetailPtrInput interface {
	pulumi.Input

	ToOrganizationResourcePropertiesOfferDetailPtrOutput() OrganizationResourcePropertiesOfferDetailPtrOutput
	ToOrganizationResourcePropertiesOfferDetailPtrOutputWithContext(context.Context) OrganizationResourcePropertiesOfferDetailPtrOutput
}

type organizationResourcePropertiesOfferDetailPtrType OrganizationResourcePropertiesOfferDetailArgs

func OrganizationResourcePropertiesOfferDetailPtr(v *OrganizationResourcePropertiesOfferDetailArgs) OrganizationResourcePropertiesOfferDetailPtrInput {
	return (*organizationResourcePropertiesOfferDetailPtrType)(v)
}

func (*organizationResourcePropertiesOfferDetailPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationResourcePropertiesOfferDetail)(nil)).Elem()
}

func (i *organizationResourcePropertiesOfferDetailPtrType) ToOrganizationResourcePropertiesOfferDetailPtrOutput() OrganizationResourcePropertiesOfferDetailPtrOutput {
	return i.ToOrganizationResourcePropertiesOfferDetailPtrOutputWithContext(context.Background())
}

func (i *organizationResourcePropertiesOfferDetailPtrType) ToOrganizationResourcePropertiesOfferDetailPtrOutputWithContext(ctx context.Context) OrganizationResourcePropertiesOfferDetailPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationResourcePropertiesOfferDetailPtrOutput)
}

type OrganizationResourcePropertiesOfferDetailOutput struct{ *pulumi.OutputState }

func (OrganizationResourcePropertiesOfferDetailOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationResourcePropertiesOfferDetail)(nil)).Elem()
}

func (o OrganizationResourcePropertiesOfferDetailOutput) ToOrganizationResourcePropertiesOfferDetailOutput() OrganizationResourcePropertiesOfferDetailOutput {
	return o
}

func (o OrganizationResourcePropertiesOfferDetailOutput) ToOrganizationResourcePropertiesOfferDetailOutputWithContext(ctx context.Context) OrganizationResourcePropertiesOfferDetailOutput {
	return o
}

func (o OrganizationResourcePropertiesOfferDetailOutput) ToOrganizationResourcePropertiesOfferDetailPtrOutput() OrganizationResourcePropertiesOfferDetailPtrOutput {
	return o.ToOrganizationResourcePropertiesOfferDetailPtrOutputWithContext(context.Background())
}

func (o OrganizationResourcePropertiesOfferDetailOutput) ToOrganizationResourcePropertiesOfferDetailPtrOutputWithContext(ctx context.Context) OrganizationResourcePropertiesOfferDetailPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OrganizationResourcePropertiesOfferDetail) *OrganizationResourcePropertiesOfferDetail {
		return &v
	}).(OrganizationResourcePropertiesOfferDetailPtrOutput)
}

func (o OrganizationResourcePropertiesOfferDetailOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OrganizationResourcePropertiesOfferDetail) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o OrganizationResourcePropertiesOfferDetailOutput) PlanId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OrganizationResourcePropertiesOfferDetail) *string { return v.PlanId }).(pulumi.StringPtrOutput)
}

func (o OrganizationResourcePropertiesOfferDetailOutput) PlanName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OrganizationResourcePropertiesOfferDetail) *string { return v.PlanName }).(pulumi.StringPtrOutput)
}

func (o OrganizationResourcePropertiesOfferDetailOutput) PublisherId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OrganizationResourcePropertiesOfferDetail) *string { return v.PublisherId }).(pulumi.StringPtrOutput)
}

func (o OrganizationResourcePropertiesOfferDetailOutput) TermUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OrganizationResourcePropertiesOfferDetail) *string { return v.TermUnit }).(pulumi.StringPtrOutput)
}

type OrganizationResourcePropertiesOfferDetailPtrOutput struct{ *pulumi.OutputState }

func (OrganizationResourcePropertiesOfferDetailPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationResourcePropertiesOfferDetail)(nil)).Elem()
}

func (o OrganizationResourcePropertiesOfferDetailPtrOutput) ToOrganizationResourcePropertiesOfferDetailPtrOutput() OrganizationResourcePropertiesOfferDetailPtrOutput {
	return o
}

func (o OrganizationResourcePropertiesOfferDetailPtrOutput) ToOrganizationResourcePropertiesOfferDetailPtrOutputWithContext(ctx context.Context) OrganizationResourcePropertiesOfferDetailPtrOutput {
	return o
}

func (o OrganizationResourcePropertiesOfferDetailPtrOutput) Elem() OrganizationResourcePropertiesOfferDetailOutput {
	return o.ApplyT(func(v *OrganizationResourcePropertiesOfferDetail) OrganizationResourcePropertiesOfferDetail {
		if v != nil {
			return *v
		}
		var ret OrganizationResourcePropertiesOfferDetail
		return ret
	}).(OrganizationResourcePropertiesOfferDetailOutput)
}

func (o OrganizationResourcePropertiesOfferDetailPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationResourcePropertiesOfferDetail) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o OrganizationResourcePropertiesOfferDetailPtrOutput) PlanId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationResourcePropertiesOfferDetail) *string {
		if v == nil {
			return nil
		}
		return v.PlanId
	}).(pulumi.StringPtrOutput)
}

func (o OrganizationResourcePropertiesOfferDetailPtrOutput) PlanName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationResourcePropertiesOfferDetail) *string {
		if v == nil {
			return nil
		}
		return v.PlanName
	}).(pulumi.StringPtrOutput)
}

func (o OrganizationResourcePropertiesOfferDetailPtrOutput) PublisherId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationResourcePropertiesOfferDetail) *string {
		if v == nil {
			return nil
		}
		return v.PublisherId
	}).(pulumi.StringPtrOutput)
}

func (o OrganizationResourcePropertiesOfferDetailPtrOutput) TermUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationResourcePropertiesOfferDetail) *string {
		if v == nil {
			return nil
		}
		return v.TermUnit
	}).(pulumi.StringPtrOutput)
}

type OrganizationResourcePropertiesResponseOfferDetail struct {
	Id          *string `pulumi:"id"`
	PlanId      *string `pulumi:"planId"`
	PlanName    *string `pulumi:"planName"`
	PublisherId *string `pulumi:"publisherId"`
	Status      string  `pulumi:"status"`
	TermUnit    *string `pulumi:"termUnit"`
}





type OrganizationResourcePropertiesResponseOfferDetailInput interface {
	pulumi.Input

	ToOrganizationResourcePropertiesResponseOfferDetailOutput() OrganizationResourcePropertiesResponseOfferDetailOutput
	ToOrganizationResourcePropertiesResponseOfferDetailOutputWithContext(context.Context) OrganizationResourcePropertiesResponseOfferDetailOutput
}

type OrganizationResourcePropertiesResponseOfferDetailArgs struct {
	Id          pulumi.StringPtrInput `pulumi:"id"`
	PlanId      pulumi.StringPtrInput `pulumi:"planId"`
	PlanName    pulumi.StringPtrInput `pulumi:"planName"`
	PublisherId pulumi.StringPtrInput `pulumi:"publisherId"`
	Status      pulumi.StringInput    `pulumi:"status"`
	TermUnit    pulumi.StringPtrInput `pulumi:"termUnit"`
}

func (OrganizationResourcePropertiesResponseOfferDetailArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationResourcePropertiesResponseOfferDetail)(nil)).Elem()
}

func (i OrganizationResourcePropertiesResponseOfferDetailArgs) ToOrganizationResourcePropertiesResponseOfferDetailOutput() OrganizationResourcePropertiesResponseOfferDetailOutput {
	return i.ToOrganizationResourcePropertiesResponseOfferDetailOutputWithContext(context.Background())
}

func (i OrganizationResourcePropertiesResponseOfferDetailArgs) ToOrganizationResourcePropertiesResponseOfferDetailOutputWithContext(ctx context.Context) OrganizationResourcePropertiesResponseOfferDetailOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationResourcePropertiesResponseOfferDetailOutput)
}

func (i OrganizationResourcePropertiesResponseOfferDetailArgs) ToOrganizationResourcePropertiesResponseOfferDetailPtrOutput() OrganizationResourcePropertiesResponseOfferDetailPtrOutput {
	return i.ToOrganizationResourcePropertiesResponseOfferDetailPtrOutputWithContext(context.Background())
}

func (i OrganizationResourcePropertiesResponseOfferDetailArgs) ToOrganizationResourcePropertiesResponseOfferDetailPtrOutputWithContext(ctx context.Context) OrganizationResourcePropertiesResponseOfferDetailPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationResourcePropertiesResponseOfferDetailOutput).ToOrganizationResourcePropertiesResponseOfferDetailPtrOutputWithContext(ctx)
}









type OrganizationResourcePropertiesResponseOfferDetailPtrInput interface {
	pulumi.Input

	ToOrganizationResourcePropertiesResponseOfferDetailPtrOutput() OrganizationResourcePropertiesResponseOfferDetailPtrOutput
	ToOrganizationResourcePropertiesResponseOfferDetailPtrOutputWithContext(context.Context) OrganizationResourcePropertiesResponseOfferDetailPtrOutput
}

type organizationResourcePropertiesResponseOfferDetailPtrType OrganizationResourcePropertiesResponseOfferDetailArgs

func OrganizationResourcePropertiesResponseOfferDetailPtr(v *OrganizationResourcePropertiesResponseOfferDetailArgs) OrganizationResourcePropertiesResponseOfferDetailPtrInput {
	return (*organizationResourcePropertiesResponseOfferDetailPtrType)(v)
}

func (*organizationResourcePropertiesResponseOfferDetailPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationResourcePropertiesResponseOfferDetail)(nil)).Elem()
}

func (i *organizationResourcePropertiesResponseOfferDetailPtrType) ToOrganizationResourcePropertiesResponseOfferDetailPtrOutput() OrganizationResourcePropertiesResponseOfferDetailPtrOutput {
	return i.ToOrganizationResourcePropertiesResponseOfferDetailPtrOutputWithContext(context.Background())
}

func (i *organizationResourcePropertiesResponseOfferDetailPtrType) ToOrganizationResourcePropertiesResponseOfferDetailPtrOutputWithContext(ctx context.Context) OrganizationResourcePropertiesResponseOfferDetailPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationResourcePropertiesResponseOfferDetailPtrOutput)
}

type OrganizationResourcePropertiesResponseOfferDetailOutput struct{ *pulumi.OutputState }

func (OrganizationResourcePropertiesResponseOfferDetailOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationResourcePropertiesResponseOfferDetail)(nil)).Elem()
}

func (o OrganizationResourcePropertiesResponseOfferDetailOutput) ToOrganizationResourcePropertiesResponseOfferDetailOutput() OrganizationResourcePropertiesResponseOfferDetailOutput {
	return o
}

func (o OrganizationResourcePropertiesResponseOfferDetailOutput) ToOrganizationResourcePropertiesResponseOfferDetailOutputWithContext(ctx context.Context) OrganizationResourcePropertiesResponseOfferDetailOutput {
	return o
}

func (o OrganizationResourcePropertiesResponseOfferDetailOutput) ToOrganizationResourcePropertiesResponseOfferDetailPtrOutput() OrganizationResourcePropertiesResponseOfferDetailPtrOutput {
	return o.ToOrganizationResourcePropertiesResponseOfferDetailPtrOutputWithContext(context.Background())
}

func (o OrganizationResourcePropertiesResponseOfferDetailOutput) ToOrganizationResourcePropertiesResponseOfferDetailPtrOutputWithContext(ctx context.Context) OrganizationResourcePropertiesResponseOfferDetailPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OrganizationResourcePropertiesResponseOfferDetail) *OrganizationResourcePropertiesResponseOfferDetail {
		return &v
	}).(OrganizationResourcePropertiesResponseOfferDetailPtrOutput)
}

func (o OrganizationResourcePropertiesResponseOfferDetailOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OrganizationResourcePropertiesResponseOfferDetail) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o OrganizationResourcePropertiesResponseOfferDetailOutput) PlanId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OrganizationResourcePropertiesResponseOfferDetail) *string { return v.PlanId }).(pulumi.StringPtrOutput)
}

func (o OrganizationResourcePropertiesResponseOfferDetailOutput) PlanName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OrganizationResourcePropertiesResponseOfferDetail) *string { return v.PlanName }).(pulumi.StringPtrOutput)
}

func (o OrganizationResourcePropertiesResponseOfferDetailOutput) PublisherId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OrganizationResourcePropertiesResponseOfferDetail) *string { return v.PublisherId }).(pulumi.StringPtrOutput)
}

func (o OrganizationResourcePropertiesResponseOfferDetailOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v OrganizationResourcePropertiesResponseOfferDetail) string { return v.Status }).(pulumi.StringOutput)
}

func (o OrganizationResourcePropertiesResponseOfferDetailOutput) TermUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OrganizationResourcePropertiesResponseOfferDetail) *string { return v.TermUnit }).(pulumi.StringPtrOutput)
}

type OrganizationResourcePropertiesResponseOfferDetailPtrOutput struct{ *pulumi.OutputState }

func (OrganizationResourcePropertiesResponseOfferDetailPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationResourcePropertiesResponseOfferDetail)(nil)).Elem()
}

func (o OrganizationResourcePropertiesResponseOfferDetailPtrOutput) ToOrganizationResourcePropertiesResponseOfferDetailPtrOutput() OrganizationResourcePropertiesResponseOfferDetailPtrOutput {
	return o
}

func (o OrganizationResourcePropertiesResponseOfferDetailPtrOutput) ToOrganizationResourcePropertiesResponseOfferDetailPtrOutputWithContext(ctx context.Context) OrganizationResourcePropertiesResponseOfferDetailPtrOutput {
	return o
}

func (o OrganizationResourcePropertiesResponseOfferDetailPtrOutput) Elem() OrganizationResourcePropertiesResponseOfferDetailOutput {
	return o.ApplyT(func(v *OrganizationResourcePropertiesResponseOfferDetail) OrganizationResourcePropertiesResponseOfferDetail {
		if v != nil {
			return *v
		}
		var ret OrganizationResourcePropertiesResponseOfferDetail
		return ret
	}).(OrganizationResourcePropertiesResponseOfferDetailOutput)
}

func (o OrganizationResourcePropertiesResponseOfferDetailPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationResourcePropertiesResponseOfferDetail) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o OrganizationResourcePropertiesResponseOfferDetailPtrOutput) PlanId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationResourcePropertiesResponseOfferDetail) *string {
		if v == nil {
			return nil
		}
		return v.PlanId
	}).(pulumi.StringPtrOutput)
}

func (o OrganizationResourcePropertiesResponseOfferDetailPtrOutput) PlanName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationResourcePropertiesResponseOfferDetail) *string {
		if v == nil {
			return nil
		}
		return v.PlanName
	}).(pulumi.StringPtrOutput)
}

func (o OrganizationResourcePropertiesResponseOfferDetailPtrOutput) PublisherId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationResourcePropertiesResponseOfferDetail) *string {
		if v == nil {
			return nil
		}
		return v.PublisherId
	}).(pulumi.StringPtrOutput)
}

func (o OrganizationResourcePropertiesResponseOfferDetailPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationResourcePropertiesResponseOfferDetail) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

func (o OrganizationResourcePropertiesResponseOfferDetailPtrOutput) TermUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationResourcePropertiesResponseOfferDetail) *string {
		if v == nil {
			return nil
		}
		return v.TermUnit
	}).(pulumi.StringPtrOutput)
}

type OrganizationResourcePropertiesResponseUserDetail struct {
	EmailAddress *string `pulumi:"emailAddress"`
	FirstName    *string `pulumi:"firstName"`
	LastName     *string `pulumi:"lastName"`
}





type OrganizationResourcePropertiesResponseUserDetailInput interface {
	pulumi.Input

	ToOrganizationResourcePropertiesResponseUserDetailOutput() OrganizationResourcePropertiesResponseUserDetailOutput
	ToOrganizationResourcePropertiesResponseUserDetailOutputWithContext(context.Context) OrganizationResourcePropertiesResponseUserDetailOutput
}

type OrganizationResourcePropertiesResponseUserDetailArgs struct {
	EmailAddress pulumi.StringPtrInput `pulumi:"emailAddress"`
	FirstName    pulumi.StringPtrInput `pulumi:"firstName"`
	LastName     pulumi.StringPtrInput `pulumi:"lastName"`
}

func (OrganizationResourcePropertiesResponseUserDetailArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationResourcePropertiesResponseUserDetail)(nil)).Elem()
}

func (i OrganizationResourcePropertiesResponseUserDetailArgs) ToOrganizationResourcePropertiesResponseUserDetailOutput() OrganizationResourcePropertiesResponseUserDetailOutput {
	return i.ToOrganizationResourcePropertiesResponseUserDetailOutputWithContext(context.Background())
}

func (i OrganizationResourcePropertiesResponseUserDetailArgs) ToOrganizationResourcePropertiesResponseUserDetailOutputWithContext(ctx context.Context) OrganizationResourcePropertiesResponseUserDetailOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationResourcePropertiesResponseUserDetailOutput)
}

func (i OrganizationResourcePropertiesResponseUserDetailArgs) ToOrganizationResourcePropertiesResponseUserDetailPtrOutput() OrganizationResourcePropertiesResponseUserDetailPtrOutput {
	return i.ToOrganizationResourcePropertiesResponseUserDetailPtrOutputWithContext(context.Background())
}

func (i OrganizationResourcePropertiesResponseUserDetailArgs) ToOrganizationResourcePropertiesResponseUserDetailPtrOutputWithContext(ctx context.Context) OrganizationResourcePropertiesResponseUserDetailPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationResourcePropertiesResponseUserDetailOutput).ToOrganizationResourcePropertiesResponseUserDetailPtrOutputWithContext(ctx)
}









type OrganizationResourcePropertiesResponseUserDetailPtrInput interface {
	pulumi.Input

	ToOrganizationResourcePropertiesResponseUserDetailPtrOutput() OrganizationResourcePropertiesResponseUserDetailPtrOutput
	ToOrganizationResourcePropertiesResponseUserDetailPtrOutputWithContext(context.Context) OrganizationResourcePropertiesResponseUserDetailPtrOutput
}

type organizationResourcePropertiesResponseUserDetailPtrType OrganizationResourcePropertiesResponseUserDetailArgs

func OrganizationResourcePropertiesResponseUserDetailPtr(v *OrganizationResourcePropertiesResponseUserDetailArgs) OrganizationResourcePropertiesResponseUserDetailPtrInput {
	return (*organizationResourcePropertiesResponseUserDetailPtrType)(v)
}

func (*organizationResourcePropertiesResponseUserDetailPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationResourcePropertiesResponseUserDetail)(nil)).Elem()
}

func (i *organizationResourcePropertiesResponseUserDetailPtrType) ToOrganizationResourcePropertiesResponseUserDetailPtrOutput() OrganizationResourcePropertiesResponseUserDetailPtrOutput {
	return i.ToOrganizationResourcePropertiesResponseUserDetailPtrOutputWithContext(context.Background())
}

func (i *organizationResourcePropertiesResponseUserDetailPtrType) ToOrganizationResourcePropertiesResponseUserDetailPtrOutputWithContext(ctx context.Context) OrganizationResourcePropertiesResponseUserDetailPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationResourcePropertiesResponseUserDetailPtrOutput)
}

type OrganizationResourcePropertiesResponseUserDetailOutput struct{ *pulumi.OutputState }

func (OrganizationResourcePropertiesResponseUserDetailOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationResourcePropertiesResponseUserDetail)(nil)).Elem()
}

func (o OrganizationResourcePropertiesResponseUserDetailOutput) ToOrganizationResourcePropertiesResponseUserDetailOutput() OrganizationResourcePropertiesResponseUserDetailOutput {
	return o
}

func (o OrganizationResourcePropertiesResponseUserDetailOutput) ToOrganizationResourcePropertiesResponseUserDetailOutputWithContext(ctx context.Context) OrganizationResourcePropertiesResponseUserDetailOutput {
	return o
}

func (o OrganizationResourcePropertiesResponseUserDetailOutput) ToOrganizationResourcePropertiesResponseUserDetailPtrOutput() OrganizationResourcePropertiesResponseUserDetailPtrOutput {
	return o.ToOrganizationResourcePropertiesResponseUserDetailPtrOutputWithContext(context.Background())
}

func (o OrganizationResourcePropertiesResponseUserDetailOutput) ToOrganizationResourcePropertiesResponseUserDetailPtrOutputWithContext(ctx context.Context) OrganizationResourcePropertiesResponseUserDetailPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OrganizationResourcePropertiesResponseUserDetail) *OrganizationResourcePropertiesResponseUserDetail {
		return &v
	}).(OrganizationResourcePropertiesResponseUserDetailPtrOutput)
}

func (o OrganizationResourcePropertiesResponseUserDetailOutput) EmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OrganizationResourcePropertiesResponseUserDetail) *string { return v.EmailAddress }).(pulumi.StringPtrOutput)
}

func (o OrganizationResourcePropertiesResponseUserDetailOutput) FirstName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OrganizationResourcePropertiesResponseUserDetail) *string { return v.FirstName }).(pulumi.StringPtrOutput)
}

func (o OrganizationResourcePropertiesResponseUserDetailOutput) LastName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OrganizationResourcePropertiesResponseUserDetail) *string { return v.LastName }).(pulumi.StringPtrOutput)
}

type OrganizationResourcePropertiesResponseUserDetailPtrOutput struct{ *pulumi.OutputState }

func (OrganizationResourcePropertiesResponseUserDetailPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationResourcePropertiesResponseUserDetail)(nil)).Elem()
}

func (o OrganizationResourcePropertiesResponseUserDetailPtrOutput) ToOrganizationResourcePropertiesResponseUserDetailPtrOutput() OrganizationResourcePropertiesResponseUserDetailPtrOutput {
	return o
}

func (o OrganizationResourcePropertiesResponseUserDetailPtrOutput) ToOrganizationResourcePropertiesResponseUserDetailPtrOutputWithContext(ctx context.Context) OrganizationResourcePropertiesResponseUserDetailPtrOutput {
	return o
}

func (o OrganizationResourcePropertiesResponseUserDetailPtrOutput) Elem() OrganizationResourcePropertiesResponseUserDetailOutput {
	return o.ApplyT(func(v *OrganizationResourcePropertiesResponseUserDetail) OrganizationResourcePropertiesResponseUserDetail {
		if v != nil {
			return *v
		}
		var ret OrganizationResourcePropertiesResponseUserDetail
		return ret
	}).(OrganizationResourcePropertiesResponseUserDetailOutput)
}

func (o OrganizationResourcePropertiesResponseUserDetailPtrOutput) EmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationResourcePropertiesResponseUserDetail) *string {
		if v == nil {
			return nil
		}
		return v.EmailAddress
	}).(pulumi.StringPtrOutput)
}

func (o OrganizationResourcePropertiesResponseUserDetailPtrOutput) FirstName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationResourcePropertiesResponseUserDetail) *string {
		if v == nil {
			return nil
		}
		return v.FirstName
	}).(pulumi.StringPtrOutput)
}

func (o OrganizationResourcePropertiesResponseUserDetailPtrOutput) LastName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationResourcePropertiesResponseUserDetail) *string {
		if v == nil {
			return nil
		}
		return v.LastName
	}).(pulumi.StringPtrOutput)
}

type OrganizationResourcePropertiesUserDetail struct {
	EmailAddress *string `pulumi:"emailAddress"`
	FirstName    *string `pulumi:"firstName"`
	LastName     *string `pulumi:"lastName"`
}





type OrganizationResourcePropertiesUserDetailInput interface {
	pulumi.Input

	ToOrganizationResourcePropertiesUserDetailOutput() OrganizationResourcePropertiesUserDetailOutput
	ToOrganizationResourcePropertiesUserDetailOutputWithContext(context.Context) OrganizationResourcePropertiesUserDetailOutput
}

type OrganizationResourcePropertiesUserDetailArgs struct {
	EmailAddress pulumi.StringPtrInput `pulumi:"emailAddress"`
	FirstName    pulumi.StringPtrInput `pulumi:"firstName"`
	LastName     pulumi.StringPtrInput `pulumi:"lastName"`
}

func (OrganizationResourcePropertiesUserDetailArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationResourcePropertiesUserDetail)(nil)).Elem()
}

func (i OrganizationResourcePropertiesUserDetailArgs) ToOrganizationResourcePropertiesUserDetailOutput() OrganizationResourcePropertiesUserDetailOutput {
	return i.ToOrganizationResourcePropertiesUserDetailOutputWithContext(context.Background())
}

func (i OrganizationResourcePropertiesUserDetailArgs) ToOrganizationResourcePropertiesUserDetailOutputWithContext(ctx context.Context) OrganizationResourcePropertiesUserDetailOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationResourcePropertiesUserDetailOutput)
}

func (i OrganizationResourcePropertiesUserDetailArgs) ToOrganizationResourcePropertiesUserDetailPtrOutput() OrganizationResourcePropertiesUserDetailPtrOutput {
	return i.ToOrganizationResourcePropertiesUserDetailPtrOutputWithContext(context.Background())
}

func (i OrganizationResourcePropertiesUserDetailArgs) ToOrganizationResourcePropertiesUserDetailPtrOutputWithContext(ctx context.Context) OrganizationResourcePropertiesUserDetailPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationResourcePropertiesUserDetailOutput).ToOrganizationResourcePropertiesUserDetailPtrOutputWithContext(ctx)
}









type OrganizationResourcePropertiesUserDetailPtrInput interface {
	pulumi.Input

	ToOrganizationResourcePropertiesUserDetailPtrOutput() OrganizationResourcePropertiesUserDetailPtrOutput
	ToOrganizationResourcePropertiesUserDetailPtrOutputWithContext(context.Context) OrganizationResourcePropertiesUserDetailPtrOutput
}

type organizationResourcePropertiesUserDetailPtrType OrganizationResourcePropertiesUserDetailArgs

func OrganizationResourcePropertiesUserDetailPtr(v *OrganizationResourcePropertiesUserDetailArgs) OrganizationResourcePropertiesUserDetailPtrInput {
	return (*organizationResourcePropertiesUserDetailPtrType)(v)
}

func (*organizationResourcePropertiesUserDetailPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationResourcePropertiesUserDetail)(nil)).Elem()
}

func (i *organizationResourcePropertiesUserDetailPtrType) ToOrganizationResourcePropertiesUserDetailPtrOutput() OrganizationResourcePropertiesUserDetailPtrOutput {
	return i.ToOrganizationResourcePropertiesUserDetailPtrOutputWithContext(context.Background())
}

func (i *organizationResourcePropertiesUserDetailPtrType) ToOrganizationResourcePropertiesUserDetailPtrOutputWithContext(ctx context.Context) OrganizationResourcePropertiesUserDetailPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationResourcePropertiesUserDetailPtrOutput)
}

type OrganizationResourcePropertiesUserDetailOutput struct{ *pulumi.OutputState }

func (OrganizationResourcePropertiesUserDetailOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationResourcePropertiesUserDetail)(nil)).Elem()
}

func (o OrganizationResourcePropertiesUserDetailOutput) ToOrganizationResourcePropertiesUserDetailOutput() OrganizationResourcePropertiesUserDetailOutput {
	return o
}

func (o OrganizationResourcePropertiesUserDetailOutput) ToOrganizationResourcePropertiesUserDetailOutputWithContext(ctx context.Context) OrganizationResourcePropertiesUserDetailOutput {
	return o
}

func (o OrganizationResourcePropertiesUserDetailOutput) ToOrganizationResourcePropertiesUserDetailPtrOutput() OrganizationResourcePropertiesUserDetailPtrOutput {
	return o.ToOrganizationResourcePropertiesUserDetailPtrOutputWithContext(context.Background())
}

func (o OrganizationResourcePropertiesUserDetailOutput) ToOrganizationResourcePropertiesUserDetailPtrOutputWithContext(ctx context.Context) OrganizationResourcePropertiesUserDetailPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OrganizationResourcePropertiesUserDetail) *OrganizationResourcePropertiesUserDetail {
		return &v
	}).(OrganizationResourcePropertiesUserDetailPtrOutput)
}

func (o OrganizationResourcePropertiesUserDetailOutput) EmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OrganizationResourcePropertiesUserDetail) *string { return v.EmailAddress }).(pulumi.StringPtrOutput)
}

func (o OrganizationResourcePropertiesUserDetailOutput) FirstName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OrganizationResourcePropertiesUserDetail) *string { return v.FirstName }).(pulumi.StringPtrOutput)
}

func (o OrganizationResourcePropertiesUserDetailOutput) LastName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OrganizationResourcePropertiesUserDetail) *string { return v.LastName }).(pulumi.StringPtrOutput)
}

type OrganizationResourcePropertiesUserDetailPtrOutput struct{ *pulumi.OutputState }

func (OrganizationResourcePropertiesUserDetailPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationResourcePropertiesUserDetail)(nil)).Elem()
}

func (o OrganizationResourcePropertiesUserDetailPtrOutput) ToOrganizationResourcePropertiesUserDetailPtrOutput() OrganizationResourcePropertiesUserDetailPtrOutput {
	return o
}

func (o OrganizationResourcePropertiesUserDetailPtrOutput) ToOrganizationResourcePropertiesUserDetailPtrOutputWithContext(ctx context.Context) OrganizationResourcePropertiesUserDetailPtrOutput {
	return o
}

func (o OrganizationResourcePropertiesUserDetailPtrOutput) Elem() OrganizationResourcePropertiesUserDetailOutput {
	return o.ApplyT(func(v *OrganizationResourcePropertiesUserDetail) OrganizationResourcePropertiesUserDetail {
		if v != nil {
			return *v
		}
		var ret OrganizationResourcePropertiesUserDetail
		return ret
	}).(OrganizationResourcePropertiesUserDetailOutput)
}

func (o OrganizationResourcePropertiesUserDetailPtrOutput) EmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationResourcePropertiesUserDetail) *string {
		if v == nil {
			return nil
		}
		return v.EmailAddress
	}).(pulumi.StringPtrOutput)
}

func (o OrganizationResourcePropertiesUserDetailPtrOutput) FirstName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationResourcePropertiesUserDetail) *string {
		if v == nil {
			return nil
		}
		return v.FirstName
	}).(pulumi.StringPtrOutput)
}

func (o OrganizationResourcePropertiesUserDetailPtrOutput) LastName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationResourcePropertiesUserDetail) *string {
		if v == nil {
			return nil
		}
		return v.LastName
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationResourcePropertiesOfferDetailInput)(nil)).Elem(), OrganizationResourcePropertiesOfferDetailArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationResourcePropertiesOfferDetailPtrInput)(nil)).Elem(), OrganizationResourcePropertiesOfferDetailArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationResourcePropertiesResponseOfferDetailInput)(nil)).Elem(), OrganizationResourcePropertiesResponseOfferDetailArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationResourcePropertiesResponseOfferDetailPtrInput)(nil)).Elem(), OrganizationResourcePropertiesResponseOfferDetailArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationResourcePropertiesResponseUserDetailInput)(nil)).Elem(), OrganizationResourcePropertiesResponseUserDetailArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationResourcePropertiesResponseUserDetailPtrInput)(nil)).Elem(), OrganizationResourcePropertiesResponseUserDetailArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationResourcePropertiesUserDetailInput)(nil)).Elem(), OrganizationResourcePropertiesUserDetailArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationResourcePropertiesUserDetailPtrInput)(nil)).Elem(), OrganizationResourcePropertiesUserDetailArgs{})
	pulumi.RegisterOutputType(OrganizationResourcePropertiesOfferDetailOutput{})
	pulumi.RegisterOutputType(OrganizationResourcePropertiesOfferDetailPtrOutput{})
	pulumi.RegisterOutputType(OrganizationResourcePropertiesResponseOfferDetailOutput{})
	pulumi.RegisterOutputType(OrganizationResourcePropertiesResponseOfferDetailPtrOutput{})
	pulumi.RegisterOutputType(OrganizationResourcePropertiesResponseUserDetailOutput{})
	pulumi.RegisterOutputType(OrganizationResourcePropertiesResponseUserDetailPtrOutput{})
	pulumi.RegisterOutputType(OrganizationResourcePropertiesUserDetailOutput{})
	pulumi.RegisterOutputType(OrganizationResourcePropertiesUserDetailPtrOutput{})
}
