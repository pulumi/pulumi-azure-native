


package v20190101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSubscription(ctx *pulumi.Context, args *LookupSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupSubscriptionResult, error) {
	var rv LookupSubscriptionResult
	err := ctx.Invoke("azure-native:apimanagement/v20190101:getSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSubscriptionArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
	Sid               string `pulumi:"sid"`
}


type LookupSubscriptionResult struct {
	AllowTracing     *bool   `pulumi:"allowTracing"`
	CreatedDate      string  `pulumi:"createdDate"`
	DisplayName      *string `pulumi:"displayName"`
	EndDate          *string `pulumi:"endDate"`
	ExpirationDate   *string `pulumi:"expirationDate"`
	Id               string  `pulumi:"id"`
	Name             string  `pulumi:"name"`
	NotificationDate *string `pulumi:"notificationDate"`
	OwnerId          *string `pulumi:"ownerId"`
	PrimaryKey       string  `pulumi:"primaryKey"`
	Scope            string  `pulumi:"scope"`
	SecondaryKey     string  `pulumi:"secondaryKey"`
	StartDate        *string `pulumi:"startDate"`
	State            string  `pulumi:"state"`
	StateComment     *string `pulumi:"stateComment"`
	Type             string  `pulumi:"type"`
}

func LookupSubscriptionOutput(ctx *pulumi.Context, args LookupSubscriptionOutputArgs, opts ...pulumi.InvokeOption) LookupSubscriptionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSubscriptionResult, error) {
			args := v.(LookupSubscriptionArgs)
			r, err := LookupSubscription(ctx, &args, opts...)
			var s LookupSubscriptionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSubscriptionResultOutput)
}

type LookupSubscriptionOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
	Sid               pulumi.StringInput `pulumi:"sid"`
}

func (LookupSubscriptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubscriptionArgs)(nil)).Elem()
}


type LookupSubscriptionResultOutput struct{ *pulumi.OutputState }

func (LookupSubscriptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubscriptionResult)(nil)).Elem()
}

func (o LookupSubscriptionResultOutput) ToLookupSubscriptionResultOutput() LookupSubscriptionResultOutput {
	return o
}

func (o LookupSubscriptionResultOutput) ToLookupSubscriptionResultOutputWithContext(ctx context.Context) LookupSubscriptionResultOutput {
	return o
}

func (o LookupSubscriptionResultOutput) AllowTracing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) *bool { return v.AllowTracing }).(pulumi.BoolPtrOutput)
}

func (o LookupSubscriptionResultOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.CreatedDate }).(pulumi.StringOutput)
}

func (o LookupSubscriptionResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupSubscriptionResultOutput) EndDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) *string { return v.EndDate }).(pulumi.StringPtrOutput)
}

func (o LookupSubscriptionResultOutput) ExpirationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) *string { return v.ExpirationDate }).(pulumi.StringPtrOutput)
}

func (o LookupSubscriptionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSubscriptionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSubscriptionResultOutput) NotificationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) *string { return v.NotificationDate }).(pulumi.StringPtrOutput)
}

func (o LookupSubscriptionResultOutput) OwnerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) *string { return v.OwnerId }).(pulumi.StringPtrOutput)
}

func (o LookupSubscriptionResultOutput) PrimaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.PrimaryKey }).(pulumi.StringOutput)
}

func (o LookupSubscriptionResultOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.Scope }).(pulumi.StringOutput)
}

func (o LookupSubscriptionResultOutput) SecondaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.SecondaryKey }).(pulumi.StringOutput)
}

func (o LookupSubscriptionResultOutput) StartDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) *string { return v.StartDate }).(pulumi.StringPtrOutput)
}

func (o LookupSubscriptionResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupSubscriptionResultOutput) StateComment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) *string { return v.StateComment }).(pulumi.StringPtrOutput)
}

func (o LookupSubscriptionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSubscriptionResultOutput{})
}
