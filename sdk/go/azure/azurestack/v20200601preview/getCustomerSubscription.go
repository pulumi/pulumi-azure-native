


package v20200601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCustomerSubscription(ctx *pulumi.Context, args *LookupCustomerSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupCustomerSubscriptionResult, error) {
	var rv LookupCustomerSubscriptionResult
	err := ctx.Invoke("azure-native:azurestack/v20200601preview:getCustomerSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCustomerSubscriptionArgs struct {
	CustomerSubscriptionName string `pulumi:"customerSubscriptionName"`
	RegistrationName         string `pulumi:"registrationName"`
	ResourceGroup            string `pulumi:"resourceGroup"`
}


type LookupCustomerSubscriptionResult struct {
	Etag       *string            `pulumi:"etag"`
	Id         string             `pulumi:"id"`
	Name       string             `pulumi:"name"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	TenantId   *string            `pulumi:"tenantId"`
	Type       string             `pulumi:"type"`
}

func LookupCustomerSubscriptionOutput(ctx *pulumi.Context, args LookupCustomerSubscriptionOutputArgs, opts ...pulumi.InvokeOption) LookupCustomerSubscriptionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCustomerSubscriptionResult, error) {
			args := v.(LookupCustomerSubscriptionArgs)
			r, err := LookupCustomerSubscription(ctx, &args, opts...)
			var s LookupCustomerSubscriptionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCustomerSubscriptionResultOutput)
}

type LookupCustomerSubscriptionOutputArgs struct {
	CustomerSubscriptionName pulumi.StringInput `pulumi:"customerSubscriptionName"`
	RegistrationName         pulumi.StringInput `pulumi:"registrationName"`
	ResourceGroup            pulumi.StringInput `pulumi:"resourceGroup"`
}

func (LookupCustomerSubscriptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomerSubscriptionArgs)(nil)).Elem()
}


type LookupCustomerSubscriptionResultOutput struct{ *pulumi.OutputState }

func (LookupCustomerSubscriptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomerSubscriptionResult)(nil)).Elem()
}

func (o LookupCustomerSubscriptionResultOutput) ToLookupCustomerSubscriptionResultOutput() LookupCustomerSubscriptionResultOutput {
	return o
}

func (o LookupCustomerSubscriptionResultOutput) ToLookupCustomerSubscriptionResultOutputWithContext(ctx context.Context) LookupCustomerSubscriptionResultOutput {
	return o
}

func (o LookupCustomerSubscriptionResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomerSubscriptionResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupCustomerSubscriptionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomerSubscriptionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCustomerSubscriptionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomerSubscriptionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCustomerSubscriptionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupCustomerSubscriptionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupCustomerSubscriptionResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomerSubscriptionResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o LookupCustomerSubscriptionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomerSubscriptionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCustomerSubscriptionResultOutput{})
}
