


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAdminRuleCollection(ctx *pulumi.Context, args *LookupAdminRuleCollectionArgs, opts ...pulumi.InvokeOption) (*LookupAdminRuleCollectionResult, error) {
	var rv LookupAdminRuleCollectionResult
	err := ctx.Invoke("azure-native:network/v20210201preview:getAdminRuleCollection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAdminRuleCollectionArgs struct {
	ConfigurationName  string `pulumi:"configurationName"`
	NetworkManagerName string `pulumi:"networkManagerName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	RuleCollectionName string `pulumi:"ruleCollectionName"`
}


type LookupAdminRuleCollectionResult struct {
	AppliesToGroups   []NetworkManagerSecurityGroupItemResponse `pulumi:"appliesToGroups"`
	Description       *string                                   `pulumi:"description"`
	DisplayName       *string                                   `pulumi:"displayName"`
	Etag              string                                    `pulumi:"etag"`
	Id                string                                    `pulumi:"id"`
	Name              string                                    `pulumi:"name"`
	ProvisioningState string                                    `pulumi:"provisioningState"`
	SystemData        SystemDataResponse                        `pulumi:"systemData"`
	Type              string                                    `pulumi:"type"`
}

func LookupAdminRuleCollectionOutput(ctx *pulumi.Context, args LookupAdminRuleCollectionOutputArgs, opts ...pulumi.InvokeOption) LookupAdminRuleCollectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAdminRuleCollectionResult, error) {
			args := v.(LookupAdminRuleCollectionArgs)
			r, err := LookupAdminRuleCollection(ctx, &args, opts...)
			var s LookupAdminRuleCollectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAdminRuleCollectionResultOutput)
}

type LookupAdminRuleCollectionOutputArgs struct {
	ConfigurationName  pulumi.StringInput `pulumi:"configurationName"`
	NetworkManagerName pulumi.StringInput `pulumi:"networkManagerName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
	RuleCollectionName pulumi.StringInput `pulumi:"ruleCollectionName"`
}

func (LookupAdminRuleCollectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAdminRuleCollectionArgs)(nil)).Elem()
}


type LookupAdminRuleCollectionResultOutput struct{ *pulumi.OutputState }

func (LookupAdminRuleCollectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAdminRuleCollectionResult)(nil)).Elem()
}

func (o LookupAdminRuleCollectionResultOutput) ToLookupAdminRuleCollectionResultOutput() LookupAdminRuleCollectionResultOutput {
	return o
}

func (o LookupAdminRuleCollectionResultOutput) ToLookupAdminRuleCollectionResultOutputWithContext(ctx context.Context) LookupAdminRuleCollectionResultOutput {
	return o
}

func (o LookupAdminRuleCollectionResultOutput) AppliesToGroups() NetworkManagerSecurityGroupItemResponseArrayOutput {
	return o.ApplyT(func(v LookupAdminRuleCollectionResult) []NetworkManagerSecurityGroupItemResponse {
		return v.AppliesToGroups
	}).(NetworkManagerSecurityGroupItemResponseArrayOutput)
}

func (o LookupAdminRuleCollectionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAdminRuleCollectionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupAdminRuleCollectionResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAdminRuleCollectionResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupAdminRuleCollectionResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdminRuleCollectionResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupAdminRuleCollectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdminRuleCollectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAdminRuleCollectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdminRuleCollectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAdminRuleCollectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdminRuleCollectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupAdminRuleCollectionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAdminRuleCollectionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupAdminRuleCollectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdminRuleCollectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAdminRuleCollectionResultOutput{})
}
