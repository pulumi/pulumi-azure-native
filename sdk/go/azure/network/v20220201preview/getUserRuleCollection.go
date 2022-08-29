


package v20220201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupUserRuleCollection(ctx *pulumi.Context, args *LookupUserRuleCollectionArgs, opts ...pulumi.InvokeOption) (*LookupUserRuleCollectionResult, error) {
	var rv LookupUserRuleCollectionResult
	err := ctx.Invoke("azure-native:network/v20220201preview:getUserRuleCollection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUserRuleCollectionArgs struct {
	ConfigurationName  string `pulumi:"configurationName"`
	NetworkManagerName string `pulumi:"networkManagerName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	RuleCollectionName string `pulumi:"ruleCollectionName"`
}


type LookupUserRuleCollectionResult struct {
	AppliesToGroups   []NetworkManagerSecurityGroupItemResponse `pulumi:"appliesToGroups"`
	Description       *string                                   `pulumi:"description"`
	Etag              string                                    `pulumi:"etag"`
	Id                string                                    `pulumi:"id"`
	Name              string                                    `pulumi:"name"`
	ProvisioningState string                                    `pulumi:"provisioningState"`
	SystemData        SystemDataResponse                        `pulumi:"systemData"`
	Type              string                                    `pulumi:"type"`
}

func LookupUserRuleCollectionOutput(ctx *pulumi.Context, args LookupUserRuleCollectionOutputArgs, opts ...pulumi.InvokeOption) LookupUserRuleCollectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUserRuleCollectionResult, error) {
			args := v.(LookupUserRuleCollectionArgs)
			r, err := LookupUserRuleCollection(ctx, &args, opts...)
			var s LookupUserRuleCollectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUserRuleCollectionResultOutput)
}

type LookupUserRuleCollectionOutputArgs struct {
	ConfigurationName  pulumi.StringInput `pulumi:"configurationName"`
	NetworkManagerName pulumi.StringInput `pulumi:"networkManagerName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
	RuleCollectionName pulumi.StringInput `pulumi:"ruleCollectionName"`
}

func (LookupUserRuleCollectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserRuleCollectionArgs)(nil)).Elem()
}


type LookupUserRuleCollectionResultOutput struct{ *pulumi.OutputState }

func (LookupUserRuleCollectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserRuleCollectionResult)(nil)).Elem()
}

func (o LookupUserRuleCollectionResultOutput) ToLookupUserRuleCollectionResultOutput() LookupUserRuleCollectionResultOutput {
	return o
}

func (o LookupUserRuleCollectionResultOutput) ToLookupUserRuleCollectionResultOutputWithContext(ctx context.Context) LookupUserRuleCollectionResultOutput {
	return o
}

func (o LookupUserRuleCollectionResultOutput) AppliesToGroups() NetworkManagerSecurityGroupItemResponseArrayOutput {
	return o.ApplyT(func(v LookupUserRuleCollectionResult) []NetworkManagerSecurityGroupItemResponse {
		return v.AppliesToGroups
	}).(NetworkManagerSecurityGroupItemResponseArrayOutput)
}

func (o LookupUserRuleCollectionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserRuleCollectionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupUserRuleCollectionResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserRuleCollectionResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupUserRuleCollectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserRuleCollectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupUserRuleCollectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserRuleCollectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupUserRuleCollectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserRuleCollectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupUserRuleCollectionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupUserRuleCollectionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupUserRuleCollectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserRuleCollectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserRuleCollectionResultOutput{})
}
