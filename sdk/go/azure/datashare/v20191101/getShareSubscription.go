


package v20191101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupShareSubscription(ctx *pulumi.Context, args *LookupShareSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupShareSubscriptionResult, error) {
	var rv LookupShareSubscriptionResult
	err := ctx.Invoke("azure-native:datashare/v20191101:getShareSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupShareSubscriptionArgs struct {
	AccountName           string `pulumi:"accountName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	ShareSubscriptionName string `pulumi:"shareSubscriptionName"`
}


type LookupShareSubscriptionResult struct {
	CreatedAt               string `pulumi:"createdAt"`
	Id                      string `pulumi:"id"`
	InvitationId            string `pulumi:"invitationId"`
	Name                    string `pulumi:"name"`
	ProviderEmail           string `pulumi:"providerEmail"`
	ProviderName            string `pulumi:"providerName"`
	ProviderTenantName      string `pulumi:"providerTenantName"`
	ProvisioningState       string `pulumi:"provisioningState"`
	ShareDescription        string `pulumi:"shareDescription"`
	ShareKind               string `pulumi:"shareKind"`
	ShareName               string `pulumi:"shareName"`
	ShareSubscriptionStatus string `pulumi:"shareSubscriptionStatus"`
	ShareTerms              string `pulumi:"shareTerms"`
	SourceShareLocation     string `pulumi:"sourceShareLocation"`
	Type                    string `pulumi:"type"`
	UserEmail               string `pulumi:"userEmail"`
	UserName                string `pulumi:"userName"`
}

func LookupShareSubscriptionOutput(ctx *pulumi.Context, args LookupShareSubscriptionOutputArgs, opts ...pulumi.InvokeOption) LookupShareSubscriptionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupShareSubscriptionResult, error) {
			args := v.(LookupShareSubscriptionArgs)
			r, err := LookupShareSubscription(ctx, &args, opts...)
			var s LookupShareSubscriptionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupShareSubscriptionResultOutput)
}

type LookupShareSubscriptionOutputArgs struct {
	AccountName           pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	ShareSubscriptionName pulumi.StringInput `pulumi:"shareSubscriptionName"`
}

func (LookupShareSubscriptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupShareSubscriptionArgs)(nil)).Elem()
}


type LookupShareSubscriptionResultOutput struct{ *pulumi.OutputState }

func (LookupShareSubscriptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupShareSubscriptionResult)(nil)).Elem()
}

func (o LookupShareSubscriptionResultOutput) ToLookupShareSubscriptionResultOutput() LookupShareSubscriptionResultOutput {
	return o
}

func (o LookupShareSubscriptionResultOutput) ToLookupShareSubscriptionResultOutputWithContext(ctx context.Context) LookupShareSubscriptionResultOutput {
	return o
}

func (o LookupShareSubscriptionResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareSubscriptionResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupShareSubscriptionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareSubscriptionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupShareSubscriptionResultOutput) InvitationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareSubscriptionResult) string { return v.InvitationId }).(pulumi.StringOutput)
}

func (o LookupShareSubscriptionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareSubscriptionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupShareSubscriptionResultOutput) ProviderEmail() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareSubscriptionResult) string { return v.ProviderEmail }).(pulumi.StringOutput)
}

func (o LookupShareSubscriptionResultOutput) ProviderName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareSubscriptionResult) string { return v.ProviderName }).(pulumi.StringOutput)
}

func (o LookupShareSubscriptionResultOutput) ProviderTenantName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareSubscriptionResult) string { return v.ProviderTenantName }).(pulumi.StringOutput)
}

func (o LookupShareSubscriptionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareSubscriptionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupShareSubscriptionResultOutput) ShareDescription() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareSubscriptionResult) string { return v.ShareDescription }).(pulumi.StringOutput)
}

func (o LookupShareSubscriptionResultOutput) ShareKind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareSubscriptionResult) string { return v.ShareKind }).(pulumi.StringOutput)
}

func (o LookupShareSubscriptionResultOutput) ShareName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareSubscriptionResult) string { return v.ShareName }).(pulumi.StringOutput)
}

func (o LookupShareSubscriptionResultOutput) ShareSubscriptionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareSubscriptionResult) string { return v.ShareSubscriptionStatus }).(pulumi.StringOutput)
}

func (o LookupShareSubscriptionResultOutput) ShareTerms() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareSubscriptionResult) string { return v.ShareTerms }).(pulumi.StringOutput)
}

func (o LookupShareSubscriptionResultOutput) SourceShareLocation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareSubscriptionResult) string { return v.SourceShareLocation }).(pulumi.StringOutput)
}

func (o LookupShareSubscriptionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareSubscriptionResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupShareSubscriptionResultOutput) UserEmail() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareSubscriptionResult) string { return v.UserEmail }).(pulumi.StringOutput)
}

func (o LookupShareSubscriptionResultOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareSubscriptionResult) string { return v.UserName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupShareSubscriptionResultOutput{})
}
