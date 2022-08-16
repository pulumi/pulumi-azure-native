


package v20170401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNamespaceAuthorizationRule(ctx *pulumi.Context, args *LookupNamespaceAuthorizationRuleArgs, opts ...pulumi.InvokeOption) (*LookupNamespaceAuthorizationRuleResult, error) {
	var rv LookupNamespaceAuthorizationRuleResult
	err := ctx.Invoke("azure-native:notificationhubs/v20170401:getNamespaceAuthorizationRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNamespaceAuthorizationRuleArgs struct {
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	NamespaceName         string `pulumi:"namespaceName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupNamespaceAuthorizationRuleResult struct {
	ClaimType    string            `pulumi:"claimType"`
	ClaimValue   string            `pulumi:"claimValue"`
	CreatedTime  string            `pulumi:"createdTime"`
	Id           string            `pulumi:"id"`
	KeyName      string            `pulumi:"keyName"`
	Location     *string           `pulumi:"location"`
	ModifiedTime string            `pulumi:"modifiedTime"`
	Name         string            `pulumi:"name"`
	PrimaryKey   string            `pulumi:"primaryKey"`
	Revision     int               `pulumi:"revision"`
	Rights       []string          `pulumi:"rights"`
	SecondaryKey string            `pulumi:"secondaryKey"`
	Sku          *SkuResponse      `pulumi:"sku"`
	Tags         map[string]string `pulumi:"tags"`
	Type         string            `pulumi:"type"`
}

func LookupNamespaceAuthorizationRuleOutput(ctx *pulumi.Context, args LookupNamespaceAuthorizationRuleOutputArgs, opts ...pulumi.InvokeOption) LookupNamespaceAuthorizationRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNamespaceAuthorizationRuleResult, error) {
			args := v.(LookupNamespaceAuthorizationRuleArgs)
			r, err := LookupNamespaceAuthorizationRule(ctx, &args, opts...)
			var s LookupNamespaceAuthorizationRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNamespaceAuthorizationRuleResultOutput)
}

type LookupNamespaceAuthorizationRuleOutputArgs struct {
	AuthorizationRuleName pulumi.StringInput `pulumi:"authorizationRuleName"`
	NamespaceName         pulumi.StringInput `pulumi:"namespaceName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNamespaceAuthorizationRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNamespaceAuthorizationRuleArgs)(nil)).Elem()
}


type LookupNamespaceAuthorizationRuleResultOutput struct{ *pulumi.OutputState }

func (LookupNamespaceAuthorizationRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNamespaceAuthorizationRuleResult)(nil)).Elem()
}

func (o LookupNamespaceAuthorizationRuleResultOutput) ToLookupNamespaceAuthorizationRuleResultOutput() LookupNamespaceAuthorizationRuleResultOutput {
	return o
}

func (o LookupNamespaceAuthorizationRuleResultOutput) ToLookupNamespaceAuthorizationRuleResultOutputWithContext(ctx context.Context) LookupNamespaceAuthorizationRuleResultOutput {
	return o
}

func (o LookupNamespaceAuthorizationRuleResultOutput) ClaimType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) string { return v.ClaimType }).(pulumi.StringOutput)
}

func (o LookupNamespaceAuthorizationRuleResultOutput) ClaimValue() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) string { return v.ClaimValue }).(pulumi.StringOutput)
}

func (o LookupNamespaceAuthorizationRuleResultOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) string { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o LookupNamespaceAuthorizationRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNamespaceAuthorizationRuleResultOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) string { return v.KeyName }).(pulumi.StringOutput)
}

func (o LookupNamespaceAuthorizationRuleResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupNamespaceAuthorizationRuleResultOutput) ModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) string { return v.ModifiedTime }).(pulumi.StringOutput)
}

func (o LookupNamespaceAuthorizationRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNamespaceAuthorizationRuleResultOutput) PrimaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) string { return v.PrimaryKey }).(pulumi.StringOutput)
}

func (o LookupNamespaceAuthorizationRuleResultOutput) Revision() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) int { return v.Revision }).(pulumi.IntOutput)
}

func (o LookupNamespaceAuthorizationRuleResultOutput) Rights() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) []string { return v.Rights }).(pulumi.StringArrayOutput)
}

func (o LookupNamespaceAuthorizationRuleResultOutput) SecondaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) string { return v.SecondaryKey }).(pulumi.StringOutput)
}

func (o LookupNamespaceAuthorizationRuleResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupNamespaceAuthorizationRuleResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupNamespaceAuthorizationRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNamespaceAuthorizationRuleResultOutput{})
}
