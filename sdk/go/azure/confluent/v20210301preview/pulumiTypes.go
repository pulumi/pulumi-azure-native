


package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type OfferDetail struct {
	Id          string `pulumi:"id"`
	PlanId      string `pulumi:"planId"`
	PlanName    string `pulumi:"planName"`
	PublisherId string `pulumi:"publisherId"`
	TermUnit    string `pulumi:"termUnit"`
}





type OfferDetailInput interface {
	pulumi.Input

	ToOfferDetailOutput() OfferDetailOutput
	ToOfferDetailOutputWithContext(context.Context) OfferDetailOutput
}

type OfferDetailArgs struct {
	Id          pulumi.StringInput `pulumi:"id"`
	PlanId      pulumi.StringInput `pulumi:"planId"`
	PlanName    pulumi.StringInput `pulumi:"planName"`
	PublisherId pulumi.StringInput `pulumi:"publisherId"`
	TermUnit    pulumi.StringInput `pulumi:"termUnit"`
}

func (OfferDetailArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OfferDetail)(nil)).Elem()
}

func (i OfferDetailArgs) ToOfferDetailOutput() OfferDetailOutput {
	return i.ToOfferDetailOutputWithContext(context.Background())
}

func (i OfferDetailArgs) ToOfferDetailOutputWithContext(ctx context.Context) OfferDetailOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OfferDetailOutput)
}

type OfferDetailOutput struct{ *pulumi.OutputState }

func (OfferDetailOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OfferDetail)(nil)).Elem()
}

func (o OfferDetailOutput) ToOfferDetailOutput() OfferDetailOutput {
	return o
}

func (o OfferDetailOutput) ToOfferDetailOutputWithContext(ctx context.Context) OfferDetailOutput {
	return o
}

func (o OfferDetailOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v OfferDetail) string { return v.Id }).(pulumi.StringOutput)
}

func (o OfferDetailOutput) PlanId() pulumi.StringOutput {
	return o.ApplyT(func(v OfferDetail) string { return v.PlanId }).(pulumi.StringOutput)
}

func (o OfferDetailOutput) PlanName() pulumi.StringOutput {
	return o.ApplyT(func(v OfferDetail) string { return v.PlanName }).(pulumi.StringOutput)
}

func (o OfferDetailOutput) PublisherId() pulumi.StringOutput {
	return o.ApplyT(func(v OfferDetail) string { return v.PublisherId }).(pulumi.StringOutput)
}

func (o OfferDetailOutput) TermUnit() pulumi.StringOutput {
	return o.ApplyT(func(v OfferDetail) string { return v.TermUnit }).(pulumi.StringOutput)
}

type OfferDetailResponse struct {
	Id          string `pulumi:"id"`
	PlanId      string `pulumi:"planId"`
	PlanName    string `pulumi:"planName"`
	PublisherId string `pulumi:"publisherId"`
	Status      string `pulumi:"status"`
	TermUnit    string `pulumi:"termUnit"`
}

type OfferDetailResponseOutput struct{ *pulumi.OutputState }

func (OfferDetailResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OfferDetailResponse)(nil)).Elem()
}

func (o OfferDetailResponseOutput) ToOfferDetailResponseOutput() OfferDetailResponseOutput {
	return o
}

func (o OfferDetailResponseOutput) ToOfferDetailResponseOutputWithContext(ctx context.Context) OfferDetailResponseOutput {
	return o
}

func (o OfferDetailResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v OfferDetailResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o OfferDetailResponseOutput) PlanId() pulumi.StringOutput {
	return o.ApplyT(func(v OfferDetailResponse) string { return v.PlanId }).(pulumi.StringOutput)
}

func (o OfferDetailResponseOutput) PlanName() pulumi.StringOutput {
	return o.ApplyT(func(v OfferDetailResponse) string { return v.PlanName }).(pulumi.StringOutput)
}

func (o OfferDetailResponseOutput) PublisherId() pulumi.StringOutput {
	return o.ApplyT(func(v OfferDetailResponse) string { return v.PublisherId }).(pulumi.StringOutput)
}

func (o OfferDetailResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v OfferDetailResponse) string { return v.Status }).(pulumi.StringOutput)
}

func (o OfferDetailResponseOutput) TermUnit() pulumi.StringOutput {
	return o.ApplyT(func(v OfferDetailResponse) string { return v.TermUnit }).(pulumi.StringOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
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

type UserDetail struct {
	EmailAddress string  `pulumi:"emailAddress"`
	FirstName    *string `pulumi:"firstName"`
	LastName     *string `pulumi:"lastName"`
}





type UserDetailInput interface {
	pulumi.Input

	ToUserDetailOutput() UserDetailOutput
	ToUserDetailOutputWithContext(context.Context) UserDetailOutput
}

type UserDetailArgs struct {
	EmailAddress pulumi.StringInput    `pulumi:"emailAddress"`
	FirstName    pulumi.StringPtrInput `pulumi:"firstName"`
	LastName     pulumi.StringPtrInput `pulumi:"lastName"`
}

func (UserDetailArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserDetail)(nil)).Elem()
}

func (i UserDetailArgs) ToUserDetailOutput() UserDetailOutput {
	return i.ToUserDetailOutputWithContext(context.Background())
}

func (i UserDetailArgs) ToUserDetailOutputWithContext(ctx context.Context) UserDetailOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserDetailOutput)
}

type UserDetailOutput struct{ *pulumi.OutputState }

func (UserDetailOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserDetail)(nil)).Elem()
}

func (o UserDetailOutput) ToUserDetailOutput() UserDetailOutput {
	return o
}

func (o UserDetailOutput) ToUserDetailOutputWithContext(ctx context.Context) UserDetailOutput {
	return o
}

func (o UserDetailOutput) EmailAddress() pulumi.StringOutput {
	return o.ApplyT(func(v UserDetail) string { return v.EmailAddress }).(pulumi.StringOutput)
}

func (o UserDetailOutput) FirstName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserDetail) *string { return v.FirstName }).(pulumi.StringPtrOutput)
}

func (o UserDetailOutput) LastName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserDetail) *string { return v.LastName }).(pulumi.StringPtrOutput)
}

type UserDetailResponse struct {
	EmailAddress string  `pulumi:"emailAddress"`
	FirstName    *string `pulumi:"firstName"`
	LastName     *string `pulumi:"lastName"`
}

type UserDetailResponseOutput struct{ *pulumi.OutputState }

func (UserDetailResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserDetailResponse)(nil)).Elem()
}

func (o UserDetailResponseOutput) ToUserDetailResponseOutput() UserDetailResponseOutput {
	return o
}

func (o UserDetailResponseOutput) ToUserDetailResponseOutputWithContext(ctx context.Context) UserDetailResponseOutput {
	return o
}

func (o UserDetailResponseOutput) EmailAddress() pulumi.StringOutput {
	return o.ApplyT(func(v UserDetailResponse) string { return v.EmailAddress }).(pulumi.StringOutput)
}

func (o UserDetailResponseOutput) FirstName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserDetailResponse) *string { return v.FirstName }).(pulumi.StringPtrOutput)
}

func (o UserDetailResponseOutput) LastName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserDetailResponse) *string { return v.LastName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(OfferDetailOutput{})
	pulumi.RegisterOutputType(OfferDetailResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(UserDetailOutput{})
	pulumi.RegisterOutputType(UserDetailResponseOutput{})
}
