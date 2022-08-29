


package v20200301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetprivateLinkForAzureAd(ctx *pulumi.Context, args *GetprivateLinkForAzureAdArgs, opts ...pulumi.InvokeOption) (*GetprivateLinkForAzureAdResult, error) {
	var rv GetprivateLinkForAzureAdResult
	err := ctx.Invoke("azure-native:aadiam/v20200301preview:getprivateLinkForAzureAd", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetprivateLinkForAzureAdArgs struct {
	PolicyName        string `pulumi:"policyName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type GetprivateLinkForAzureAdResult struct {
	AllTenants     *bool             `pulumi:"allTenants"`
	Id             string            `pulumi:"id"`
	Name           *string           `pulumi:"name"`
	OwnerTenantId  *string           `pulumi:"ownerTenantId"`
	ResourceGroup  *string           `pulumi:"resourceGroup"`
	ResourceName   *string           `pulumi:"resourceName"`
	SubscriptionId *string           `pulumi:"subscriptionId"`
	Tags           map[string]string `pulumi:"tags"`
	Tenants        []string          `pulumi:"tenants"`
	Type           string            `pulumi:"type"`
}

func GetprivateLinkForAzureAdOutput(ctx *pulumi.Context, args GetprivateLinkForAzureAdOutputArgs, opts ...pulumi.InvokeOption) GetprivateLinkForAzureAdResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetprivateLinkForAzureAdResult, error) {
			args := v.(GetprivateLinkForAzureAdArgs)
			r, err := GetprivateLinkForAzureAd(ctx, &args, opts...)
			var s GetprivateLinkForAzureAdResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetprivateLinkForAzureAdResultOutput)
}

type GetprivateLinkForAzureAdOutputArgs struct {
	PolicyName        pulumi.StringInput `pulumi:"policyName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetprivateLinkForAzureAdOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetprivateLinkForAzureAdArgs)(nil)).Elem()
}


type GetprivateLinkForAzureAdResultOutput struct{ *pulumi.OutputState }

func (GetprivateLinkForAzureAdResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetprivateLinkForAzureAdResult)(nil)).Elem()
}

func (o GetprivateLinkForAzureAdResultOutput) ToGetprivateLinkForAzureAdResultOutput() GetprivateLinkForAzureAdResultOutput {
	return o
}

func (o GetprivateLinkForAzureAdResultOutput) ToGetprivateLinkForAzureAdResultOutputWithContext(ctx context.Context) GetprivateLinkForAzureAdResultOutput {
	return o
}

func (o GetprivateLinkForAzureAdResultOutput) AllTenants() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetprivateLinkForAzureAdResult) *bool { return v.AllTenants }).(pulumi.BoolPtrOutput)
}

func (o GetprivateLinkForAzureAdResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetprivateLinkForAzureAdResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetprivateLinkForAzureAdResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetprivateLinkForAzureAdResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetprivateLinkForAzureAdResultOutput) OwnerTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetprivateLinkForAzureAdResult) *string { return v.OwnerTenantId }).(pulumi.StringPtrOutput)
}

func (o GetprivateLinkForAzureAdResultOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetprivateLinkForAzureAdResult) *string { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

func (o GetprivateLinkForAzureAdResultOutput) ResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetprivateLinkForAzureAdResult) *string { return v.ResourceName }).(pulumi.StringPtrOutput)
}

func (o GetprivateLinkForAzureAdResultOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetprivateLinkForAzureAdResult) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

func (o GetprivateLinkForAzureAdResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetprivateLinkForAzureAdResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GetprivateLinkForAzureAdResultOutput) Tenants() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetprivateLinkForAzureAdResult) []string { return v.Tenants }).(pulumi.StringArrayOutput)
}

func (o GetprivateLinkForAzureAdResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetprivateLinkForAzureAdResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetprivateLinkForAzureAdResultOutput{})
}
