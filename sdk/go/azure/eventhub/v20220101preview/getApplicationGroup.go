


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApplicationGroup(ctx *pulumi.Context, args *LookupApplicationGroupArgs, opts ...pulumi.InvokeOption) (*LookupApplicationGroupResult, error) {
	var rv LookupApplicationGroupResult
	err := ctx.Invoke("azure-native:eventhub/v20220101preview:getApplicationGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationGroupArgs struct {
	ApplicationGroupName string `pulumi:"applicationGroupName"`
	NamespaceName        string `pulumi:"namespaceName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupApplicationGroupResult struct {
	ClientAppGroupIdentifier string                     `pulumi:"clientAppGroupIdentifier"`
	Id                       string                     `pulumi:"id"`
	IsEnabled                *bool                      `pulumi:"isEnabled"`
	Location                 string                     `pulumi:"location"`
	Name                     string                     `pulumi:"name"`
	Policies                 []ThrottlingPolicyResponse `pulumi:"policies"`
	SystemData               SystemDataResponse         `pulumi:"systemData"`
	Type                     string                     `pulumi:"type"`
}

func LookupApplicationGroupOutput(ctx *pulumi.Context, args LookupApplicationGroupOutputArgs, opts ...pulumi.InvokeOption) LookupApplicationGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApplicationGroupResult, error) {
			args := v.(LookupApplicationGroupArgs)
			r, err := LookupApplicationGroup(ctx, &args, opts...)
			var s LookupApplicationGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApplicationGroupResultOutput)
}

type LookupApplicationGroupOutputArgs struct {
	ApplicationGroupName pulumi.StringInput `pulumi:"applicationGroupName"`
	NamespaceName        pulumi.StringInput `pulumi:"namespaceName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupApplicationGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationGroupArgs)(nil)).Elem()
}


type LookupApplicationGroupResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationGroupResult)(nil)).Elem()
}

func (o LookupApplicationGroupResultOutput) ToLookupApplicationGroupResultOutput() LookupApplicationGroupResultOutput {
	return o
}

func (o LookupApplicationGroupResultOutput) ToLookupApplicationGroupResultOutputWithContext(ctx context.Context) LookupApplicationGroupResultOutput {
	return o
}

func (o LookupApplicationGroupResultOutput) ClientAppGroupIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) string { return v.ClientAppGroupIdentifier }).(pulumi.StringOutput)
}

func (o LookupApplicationGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApplicationGroupResultOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) *bool { return v.IsEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupApplicationGroupResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupApplicationGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupApplicationGroupResultOutput) Policies() ThrottlingPolicyResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) []ThrottlingPolicyResponse { return v.Policies }).(ThrottlingPolicyResponseArrayOutput)
}

func (o LookupApplicationGroupResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupApplicationGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationGroupResultOutput{})
}
