


package avs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPlacementPolicy(ctx *pulumi.Context, args *LookupPlacementPolicyArgs, opts ...pulumi.InvokeOption) (*LookupPlacementPolicyResult, error) {
	var rv LookupPlacementPolicyResult
	err := ctx.Invoke("azure-native:avs:getPlacementPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPlacementPolicyArgs struct {
	ClusterName         string `pulumi:"clusterName"`
	PlacementPolicyName string `pulumi:"placementPolicyName"`
	PrivateCloudName    string `pulumi:"privateCloudName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}


type LookupPlacementPolicyResult struct {
	Id         string      `pulumi:"id"`
	Name       string      `pulumi:"name"`
	Properties interface{} `pulumi:"properties"`
	Type       string      `pulumi:"type"`
}

func LookupPlacementPolicyOutput(ctx *pulumi.Context, args LookupPlacementPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupPlacementPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPlacementPolicyResult, error) {
			args := v.(LookupPlacementPolicyArgs)
			r, err := LookupPlacementPolicy(ctx, &args, opts...)
			var s LookupPlacementPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPlacementPolicyResultOutput)
}

type LookupPlacementPolicyOutputArgs struct {
	ClusterName         pulumi.StringInput `pulumi:"clusterName"`
	PlacementPolicyName pulumi.StringInput `pulumi:"placementPolicyName"`
	PrivateCloudName    pulumi.StringInput `pulumi:"privateCloudName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPlacementPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPlacementPolicyArgs)(nil)).Elem()
}


type LookupPlacementPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupPlacementPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPlacementPolicyResult)(nil)).Elem()
}

func (o LookupPlacementPolicyResultOutput) ToLookupPlacementPolicyResultOutput() LookupPlacementPolicyResultOutput {
	return o
}

func (o LookupPlacementPolicyResultOutput) ToLookupPlacementPolicyResultOutputWithContext(ctx context.Context) LookupPlacementPolicyResultOutput {
	return o
}

func (o LookupPlacementPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPlacementPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPlacementPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPlacementPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPlacementPolicyResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupPlacementPolicyResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func (o LookupPlacementPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPlacementPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPlacementPolicyResultOutput{})
}
