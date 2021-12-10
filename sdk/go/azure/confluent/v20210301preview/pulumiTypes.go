


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

func (i OfferDetailArgs) ToOfferDetailPtrOutput() OfferDetailPtrOutput {
	return i.ToOfferDetailPtrOutputWithContext(context.Background())
}

func (i OfferDetailArgs) ToOfferDetailPtrOutputWithContext(ctx context.Context) OfferDetailPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OfferDetailOutput).ToOfferDetailPtrOutputWithContext(ctx)
}









type OfferDetailPtrInput interface {
	pulumi.Input

	ToOfferDetailPtrOutput() OfferDetailPtrOutput
	ToOfferDetailPtrOutputWithContext(context.Context) OfferDetailPtrOutput
}

type offerDetailPtrType OfferDetailArgs

func OfferDetailPtr(v *OfferDetailArgs) OfferDetailPtrInput {
	return (*offerDetailPtrType)(v)
}

func (*offerDetailPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OfferDetail)(nil)).Elem()
}

func (i *offerDetailPtrType) ToOfferDetailPtrOutput() OfferDetailPtrOutput {
	return i.ToOfferDetailPtrOutputWithContext(context.Background())
}

func (i *offerDetailPtrType) ToOfferDetailPtrOutputWithContext(ctx context.Context) OfferDetailPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OfferDetailPtrOutput)
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

func (o OfferDetailOutput) ToOfferDetailPtrOutput() OfferDetailPtrOutput {
	return o.ToOfferDetailPtrOutputWithContext(context.Background())
}

func (o OfferDetailOutput) ToOfferDetailPtrOutputWithContext(ctx context.Context) OfferDetailPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OfferDetail) *OfferDetail {
		return &v
	}).(OfferDetailPtrOutput)
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

type OfferDetailPtrOutput struct{ *pulumi.OutputState }

func (OfferDetailPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OfferDetail)(nil)).Elem()
}

func (o OfferDetailPtrOutput) ToOfferDetailPtrOutput() OfferDetailPtrOutput {
	return o
}

func (o OfferDetailPtrOutput) ToOfferDetailPtrOutputWithContext(ctx context.Context) OfferDetailPtrOutput {
	return o
}

func (o OfferDetailPtrOutput) Elem() OfferDetailOutput {
	return o.ApplyT(func(v *OfferDetail) OfferDetail {
		if v != nil {
			return *v
		}
		var ret OfferDetail
		return ret
	}).(OfferDetailOutput)
}

func (o OfferDetailPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OfferDetail) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o OfferDetailPtrOutput) PlanId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OfferDetail) *string {
		if v == nil {
			return nil
		}
		return &v.PlanId
	}).(pulumi.StringPtrOutput)
}

func (o OfferDetailPtrOutput) PlanName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OfferDetail) *string {
		if v == nil {
			return nil
		}
		return &v.PlanName
	}).(pulumi.StringPtrOutput)
}

func (o OfferDetailPtrOutput) PublisherId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OfferDetail) *string {
		if v == nil {
			return nil
		}
		return &v.PublisherId
	}).(pulumi.StringPtrOutput)
}

func (o OfferDetailPtrOutput) TermUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OfferDetail) *string {
		if v == nil {
			return nil
		}
		return &v.TermUnit
	}).(pulumi.StringPtrOutput)
}

type OfferDetailResponse struct {
	Id          string `pulumi:"id"`
	PlanId      string `pulumi:"planId"`
	PlanName    string `pulumi:"planName"`
	PublisherId string `pulumi:"publisherId"`
	Status      string `pulumi:"status"`
	TermUnit    string `pulumi:"termUnit"`
}





type OfferDetailResponseInput interface {
	pulumi.Input

	ToOfferDetailResponseOutput() OfferDetailResponseOutput
	ToOfferDetailResponseOutputWithContext(context.Context) OfferDetailResponseOutput
}

type OfferDetailResponseArgs struct {
	Id          pulumi.StringInput `pulumi:"id"`
	PlanId      pulumi.StringInput `pulumi:"planId"`
	PlanName    pulumi.StringInput `pulumi:"planName"`
	PublisherId pulumi.StringInput `pulumi:"publisherId"`
	Status      pulumi.StringInput `pulumi:"status"`
	TermUnit    pulumi.StringInput `pulumi:"termUnit"`
}

func (OfferDetailResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OfferDetailResponse)(nil)).Elem()
}

func (i OfferDetailResponseArgs) ToOfferDetailResponseOutput() OfferDetailResponseOutput {
	return i.ToOfferDetailResponseOutputWithContext(context.Background())
}

func (i OfferDetailResponseArgs) ToOfferDetailResponseOutputWithContext(ctx context.Context) OfferDetailResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OfferDetailResponseOutput)
}

func (i OfferDetailResponseArgs) ToOfferDetailResponsePtrOutput() OfferDetailResponsePtrOutput {
	return i.ToOfferDetailResponsePtrOutputWithContext(context.Background())
}

func (i OfferDetailResponseArgs) ToOfferDetailResponsePtrOutputWithContext(ctx context.Context) OfferDetailResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OfferDetailResponseOutput).ToOfferDetailResponsePtrOutputWithContext(ctx)
}









type OfferDetailResponsePtrInput interface {
	pulumi.Input

	ToOfferDetailResponsePtrOutput() OfferDetailResponsePtrOutput
	ToOfferDetailResponsePtrOutputWithContext(context.Context) OfferDetailResponsePtrOutput
}

type offerDetailResponsePtrType OfferDetailResponseArgs

func OfferDetailResponsePtr(v *OfferDetailResponseArgs) OfferDetailResponsePtrInput {
	return (*offerDetailResponsePtrType)(v)
}

func (*offerDetailResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OfferDetailResponse)(nil)).Elem()
}

func (i *offerDetailResponsePtrType) ToOfferDetailResponsePtrOutput() OfferDetailResponsePtrOutput {
	return i.ToOfferDetailResponsePtrOutputWithContext(context.Background())
}

func (i *offerDetailResponsePtrType) ToOfferDetailResponsePtrOutputWithContext(ctx context.Context) OfferDetailResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OfferDetailResponsePtrOutput)
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

func (o OfferDetailResponseOutput) ToOfferDetailResponsePtrOutput() OfferDetailResponsePtrOutput {
	return o.ToOfferDetailResponsePtrOutputWithContext(context.Background())
}

func (o OfferDetailResponseOutput) ToOfferDetailResponsePtrOutputWithContext(ctx context.Context) OfferDetailResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OfferDetailResponse) *OfferDetailResponse {
		return &v
	}).(OfferDetailResponsePtrOutput)
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

type OfferDetailResponsePtrOutput struct{ *pulumi.OutputState }

func (OfferDetailResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OfferDetailResponse)(nil)).Elem()
}

func (o OfferDetailResponsePtrOutput) ToOfferDetailResponsePtrOutput() OfferDetailResponsePtrOutput {
	return o
}

func (o OfferDetailResponsePtrOutput) ToOfferDetailResponsePtrOutputWithContext(ctx context.Context) OfferDetailResponsePtrOutput {
	return o
}

func (o OfferDetailResponsePtrOutput) Elem() OfferDetailResponseOutput {
	return o.ApplyT(func(v *OfferDetailResponse) OfferDetailResponse {
		if v != nil {
			return *v
		}
		var ret OfferDetailResponse
		return ret
	}).(OfferDetailResponseOutput)
}

func (o OfferDetailResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OfferDetailResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o OfferDetailResponsePtrOutput) PlanId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OfferDetailResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PlanId
	}).(pulumi.StringPtrOutput)
}

func (o OfferDetailResponsePtrOutput) PlanName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OfferDetailResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PlanName
	}).(pulumi.StringPtrOutput)
}

func (o OfferDetailResponsePtrOutput) PublisherId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OfferDetailResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PublisherId
	}).(pulumi.StringPtrOutput)
}

func (o OfferDetailResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OfferDetailResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

func (o OfferDetailResponsePtrOutput) TermUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OfferDetailResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TermUnit
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

func (i UserDetailArgs) ToUserDetailPtrOutput() UserDetailPtrOutput {
	return i.ToUserDetailPtrOutputWithContext(context.Background())
}

func (i UserDetailArgs) ToUserDetailPtrOutputWithContext(ctx context.Context) UserDetailPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserDetailOutput).ToUserDetailPtrOutputWithContext(ctx)
}









type UserDetailPtrInput interface {
	pulumi.Input

	ToUserDetailPtrOutput() UserDetailPtrOutput
	ToUserDetailPtrOutputWithContext(context.Context) UserDetailPtrOutput
}

type userDetailPtrType UserDetailArgs

func UserDetailPtr(v *UserDetailArgs) UserDetailPtrInput {
	return (*userDetailPtrType)(v)
}

func (*userDetailPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserDetail)(nil)).Elem()
}

func (i *userDetailPtrType) ToUserDetailPtrOutput() UserDetailPtrOutput {
	return i.ToUserDetailPtrOutputWithContext(context.Background())
}

func (i *userDetailPtrType) ToUserDetailPtrOutputWithContext(ctx context.Context) UserDetailPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserDetailPtrOutput)
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

func (o UserDetailOutput) ToUserDetailPtrOutput() UserDetailPtrOutput {
	return o.ToUserDetailPtrOutputWithContext(context.Background())
}

func (o UserDetailOutput) ToUserDetailPtrOutputWithContext(ctx context.Context) UserDetailPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserDetail) *UserDetail {
		return &v
	}).(UserDetailPtrOutput)
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

type UserDetailPtrOutput struct{ *pulumi.OutputState }

func (UserDetailPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserDetail)(nil)).Elem()
}

func (o UserDetailPtrOutput) ToUserDetailPtrOutput() UserDetailPtrOutput {
	return o
}

func (o UserDetailPtrOutput) ToUserDetailPtrOutputWithContext(ctx context.Context) UserDetailPtrOutput {
	return o
}

func (o UserDetailPtrOutput) Elem() UserDetailOutput {
	return o.ApplyT(func(v *UserDetail) UserDetail {
		if v != nil {
			return *v
		}
		var ret UserDetail
		return ret
	}).(UserDetailOutput)
}

func (o UserDetailPtrOutput) EmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserDetail) *string {
		if v == nil {
			return nil
		}
		return &v.EmailAddress
	}).(pulumi.StringPtrOutput)
}

func (o UserDetailPtrOutput) FirstName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserDetail) *string {
		if v == nil {
			return nil
		}
		return v.FirstName
	}).(pulumi.StringPtrOutput)
}

func (o UserDetailPtrOutput) LastName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserDetail) *string {
		if v == nil {
			return nil
		}
		return v.LastName
	}).(pulumi.StringPtrOutput)
}

type UserDetailResponse struct {
	EmailAddress string  `pulumi:"emailAddress"`
	FirstName    *string `pulumi:"firstName"`
	LastName     *string `pulumi:"lastName"`
}





type UserDetailResponseInput interface {
	pulumi.Input

	ToUserDetailResponseOutput() UserDetailResponseOutput
	ToUserDetailResponseOutputWithContext(context.Context) UserDetailResponseOutput
}

type UserDetailResponseArgs struct {
	EmailAddress pulumi.StringInput    `pulumi:"emailAddress"`
	FirstName    pulumi.StringPtrInput `pulumi:"firstName"`
	LastName     pulumi.StringPtrInput `pulumi:"lastName"`
}

func (UserDetailResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserDetailResponse)(nil)).Elem()
}

func (i UserDetailResponseArgs) ToUserDetailResponseOutput() UserDetailResponseOutput {
	return i.ToUserDetailResponseOutputWithContext(context.Background())
}

func (i UserDetailResponseArgs) ToUserDetailResponseOutputWithContext(ctx context.Context) UserDetailResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserDetailResponseOutput)
}

func (i UserDetailResponseArgs) ToUserDetailResponsePtrOutput() UserDetailResponsePtrOutput {
	return i.ToUserDetailResponsePtrOutputWithContext(context.Background())
}

func (i UserDetailResponseArgs) ToUserDetailResponsePtrOutputWithContext(ctx context.Context) UserDetailResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserDetailResponseOutput).ToUserDetailResponsePtrOutputWithContext(ctx)
}









type UserDetailResponsePtrInput interface {
	pulumi.Input

	ToUserDetailResponsePtrOutput() UserDetailResponsePtrOutput
	ToUserDetailResponsePtrOutputWithContext(context.Context) UserDetailResponsePtrOutput
}

type userDetailResponsePtrType UserDetailResponseArgs

func UserDetailResponsePtr(v *UserDetailResponseArgs) UserDetailResponsePtrInput {
	return (*userDetailResponsePtrType)(v)
}

func (*userDetailResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserDetailResponse)(nil)).Elem()
}

func (i *userDetailResponsePtrType) ToUserDetailResponsePtrOutput() UserDetailResponsePtrOutput {
	return i.ToUserDetailResponsePtrOutputWithContext(context.Background())
}

func (i *userDetailResponsePtrType) ToUserDetailResponsePtrOutputWithContext(ctx context.Context) UserDetailResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserDetailResponsePtrOutput)
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

func (o UserDetailResponseOutput) ToUserDetailResponsePtrOutput() UserDetailResponsePtrOutput {
	return o.ToUserDetailResponsePtrOutputWithContext(context.Background())
}

func (o UserDetailResponseOutput) ToUserDetailResponsePtrOutputWithContext(ctx context.Context) UserDetailResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserDetailResponse) *UserDetailResponse {
		return &v
	}).(UserDetailResponsePtrOutput)
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

type UserDetailResponsePtrOutput struct{ *pulumi.OutputState }

func (UserDetailResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserDetailResponse)(nil)).Elem()
}

func (o UserDetailResponsePtrOutput) ToUserDetailResponsePtrOutput() UserDetailResponsePtrOutput {
	return o
}

func (o UserDetailResponsePtrOutput) ToUserDetailResponsePtrOutputWithContext(ctx context.Context) UserDetailResponsePtrOutput {
	return o
}

func (o UserDetailResponsePtrOutput) Elem() UserDetailResponseOutput {
	return o.ApplyT(func(v *UserDetailResponse) UserDetailResponse {
		if v != nil {
			return *v
		}
		var ret UserDetailResponse
		return ret
	}).(UserDetailResponseOutput)
}

func (o UserDetailResponsePtrOutput) EmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserDetailResponse) *string {
		if v == nil {
			return nil
		}
		return &v.EmailAddress
	}).(pulumi.StringPtrOutput)
}

func (o UserDetailResponsePtrOutput) FirstName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserDetailResponse) *string {
		if v == nil {
			return nil
		}
		return v.FirstName
	}).(pulumi.StringPtrOutput)
}

func (o UserDetailResponsePtrOutput) LastName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserDetailResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastName
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(OfferDetailOutput{})
	pulumi.RegisterOutputType(OfferDetailPtrOutput{})
	pulumi.RegisterOutputType(OfferDetailResponseOutput{})
	pulumi.RegisterOutputType(OfferDetailResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
	pulumi.RegisterOutputType(UserDetailOutput{})
	pulumi.RegisterOutputType(UserDetailPtrOutput{})
	pulumi.RegisterOutputType(UserDetailResponseOutput{})
	pulumi.RegisterOutputType(UserDetailResponsePtrOutput{})
}
