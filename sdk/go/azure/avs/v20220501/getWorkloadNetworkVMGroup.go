


package v20220501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkloadNetworkVMGroup(ctx *pulumi.Context, args *LookupWorkloadNetworkVMGroupArgs, opts ...pulumi.InvokeOption) (*LookupWorkloadNetworkVMGroupResult, error) {
	var rv LookupWorkloadNetworkVMGroupResult
	err := ctx.Invoke("azure-native:avs/v20220501:getWorkloadNetworkVMGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkloadNetworkVMGroupArgs struct {
	PrivateCloudName  string `pulumi:"privateCloudName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VmGroupId         string `pulumi:"vmGroupId"`
}


type LookupWorkloadNetworkVMGroupResult struct {
	DisplayName       *string  `pulumi:"displayName"`
	Id                string   `pulumi:"id"`
	Members           []string `pulumi:"members"`
	Name              string   `pulumi:"name"`
	ProvisioningState string   `pulumi:"provisioningState"`
	Revision          *float64 `pulumi:"revision"`
	Status            string   `pulumi:"status"`
	Type              string   `pulumi:"type"`
}

func LookupWorkloadNetworkVMGroupOutput(ctx *pulumi.Context, args LookupWorkloadNetworkVMGroupOutputArgs, opts ...pulumi.InvokeOption) LookupWorkloadNetworkVMGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkloadNetworkVMGroupResult, error) {
			args := v.(LookupWorkloadNetworkVMGroupArgs)
			r, err := LookupWorkloadNetworkVMGroup(ctx, &args, opts...)
			var s LookupWorkloadNetworkVMGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkloadNetworkVMGroupResultOutput)
}

type LookupWorkloadNetworkVMGroupOutputArgs struct {
	PrivateCloudName  pulumi.StringInput `pulumi:"privateCloudName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	VmGroupId         pulumi.StringInput `pulumi:"vmGroupId"`
}

func (LookupWorkloadNetworkVMGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkloadNetworkVMGroupArgs)(nil)).Elem()
}


type LookupWorkloadNetworkVMGroupResultOutput struct{ *pulumi.OutputState }

func (LookupWorkloadNetworkVMGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkloadNetworkVMGroupResult)(nil)).Elem()
}

func (o LookupWorkloadNetworkVMGroupResultOutput) ToLookupWorkloadNetworkVMGroupResultOutput() LookupWorkloadNetworkVMGroupResultOutput {
	return o
}

func (o LookupWorkloadNetworkVMGroupResultOutput) ToLookupWorkloadNetworkVMGroupResultOutputWithContext(ctx context.Context) LookupWorkloadNetworkVMGroupResultOutput {
	return o
}

func (o LookupWorkloadNetworkVMGroupResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkVMGroupResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupWorkloadNetworkVMGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkVMGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWorkloadNetworkVMGroupResultOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkVMGroupResult) []string { return v.Members }).(pulumi.StringArrayOutput)
}

func (o LookupWorkloadNetworkVMGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkVMGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWorkloadNetworkVMGroupResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkVMGroupResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupWorkloadNetworkVMGroupResultOutput) Revision() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkVMGroupResult) *float64 { return v.Revision }).(pulumi.Float64PtrOutput)
}

func (o LookupWorkloadNetworkVMGroupResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkVMGroupResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupWorkloadNetworkVMGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadNetworkVMGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkloadNetworkVMGroupResultOutput{})
}
