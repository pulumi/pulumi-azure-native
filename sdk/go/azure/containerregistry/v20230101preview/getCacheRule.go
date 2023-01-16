


package v20230101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCacheRule(ctx *pulumi.Context, args *LookupCacheRuleArgs, opts ...pulumi.InvokeOption) (*LookupCacheRuleResult, error) {
	var rv LookupCacheRuleResult
	err := ctx.Invoke("azure-native:containerregistry/v20230101preview:getCacheRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCacheRuleArgs struct {
	CacheRuleName     string `pulumi:"cacheRuleName"`
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupCacheRuleResult struct {
	CreationDate            string             `pulumi:"creationDate"`
	CredentialSetResourceId *string            `pulumi:"credentialSetResourceId"`
	Id                      string             `pulumi:"id"`
	Name                    string             `pulumi:"name"`
	ProvisioningState       string             `pulumi:"provisioningState"`
	SourceRepository        *string            `pulumi:"sourceRepository"`
	SystemData              SystemDataResponse `pulumi:"systemData"`
	TargetRepository        *string            `pulumi:"targetRepository"`
	Type                    string             `pulumi:"type"`
}

func LookupCacheRuleOutput(ctx *pulumi.Context, args LookupCacheRuleOutputArgs, opts ...pulumi.InvokeOption) LookupCacheRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCacheRuleResult, error) {
			args := v.(LookupCacheRuleArgs)
			r, err := LookupCacheRule(ctx, &args, opts...)
			var s LookupCacheRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCacheRuleResultOutput)
}

type LookupCacheRuleOutputArgs struct {
	CacheRuleName     pulumi.StringInput `pulumi:"cacheRuleName"`
	RegistryName      pulumi.StringInput `pulumi:"registryName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCacheRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCacheRuleArgs)(nil)).Elem()
}


type LookupCacheRuleResultOutput struct{ *pulumi.OutputState }

func (LookupCacheRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCacheRuleResult)(nil)).Elem()
}

func (o LookupCacheRuleResultOutput) ToLookupCacheRuleResultOutput() LookupCacheRuleResultOutput {
	return o
}

func (o LookupCacheRuleResultOutput) ToLookupCacheRuleResultOutputWithContext(ctx context.Context) LookupCacheRuleResultOutput {
	return o
}

func (o LookupCacheRuleResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCacheRuleResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

func (o LookupCacheRuleResultOutput) CredentialSetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCacheRuleResult) *string { return v.CredentialSetResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupCacheRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCacheRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCacheRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCacheRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCacheRuleResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCacheRuleResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupCacheRuleResultOutput) SourceRepository() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCacheRuleResult) *string { return v.SourceRepository }).(pulumi.StringPtrOutput)
}

func (o LookupCacheRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupCacheRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupCacheRuleResultOutput) TargetRepository() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCacheRuleResult) *string { return v.TargetRepository }).(pulumi.StringPtrOutput)
}

func (o LookupCacheRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCacheRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCacheRuleResultOutput{})
}
