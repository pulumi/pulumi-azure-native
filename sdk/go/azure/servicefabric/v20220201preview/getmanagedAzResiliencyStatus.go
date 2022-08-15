


package v20220201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetmanagedAzResiliencyStatus(ctx *pulumi.Context, args *GetmanagedAzResiliencyStatusArgs, opts ...pulumi.InvokeOption) (*GetmanagedAzResiliencyStatusResult, error) {
	var rv GetmanagedAzResiliencyStatusResult
	err := ctx.Invoke("azure-native:servicefabric/v20220201preview:getmanagedAzResiliencyStatus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetmanagedAzResiliencyStatusArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type GetmanagedAzResiliencyStatusResult struct {
	BaseResourceStatus     []ResourceAzStatusResponse `pulumi:"baseResourceStatus"`
	IsClusterZoneResilient bool                       `pulumi:"isClusterZoneResilient"`
}

func GetmanagedAzResiliencyStatusOutput(ctx *pulumi.Context, args GetmanagedAzResiliencyStatusOutputArgs, opts ...pulumi.InvokeOption) GetmanagedAzResiliencyStatusResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetmanagedAzResiliencyStatusResult, error) {
			args := v.(GetmanagedAzResiliencyStatusArgs)
			r, err := GetmanagedAzResiliencyStatus(ctx, &args, opts...)
			var s GetmanagedAzResiliencyStatusResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetmanagedAzResiliencyStatusResultOutput)
}

type GetmanagedAzResiliencyStatusOutputArgs struct {
	ClusterName       pulumi.StringInput `pulumi:"clusterName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetmanagedAzResiliencyStatusOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetmanagedAzResiliencyStatusArgs)(nil)).Elem()
}


type GetmanagedAzResiliencyStatusResultOutput struct{ *pulumi.OutputState }

func (GetmanagedAzResiliencyStatusResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetmanagedAzResiliencyStatusResult)(nil)).Elem()
}

func (o GetmanagedAzResiliencyStatusResultOutput) ToGetmanagedAzResiliencyStatusResultOutput() GetmanagedAzResiliencyStatusResultOutput {
	return o
}

func (o GetmanagedAzResiliencyStatusResultOutput) ToGetmanagedAzResiliencyStatusResultOutputWithContext(ctx context.Context) GetmanagedAzResiliencyStatusResultOutput {
	return o
}

func (o GetmanagedAzResiliencyStatusResultOutput) BaseResourceStatus() ResourceAzStatusResponseArrayOutput {
	return o.ApplyT(func(v GetmanagedAzResiliencyStatusResult) []ResourceAzStatusResponse { return v.BaseResourceStatus }).(ResourceAzStatusResponseArrayOutput)
}

func (o GetmanagedAzResiliencyStatusResultOutput) IsClusterZoneResilient() pulumi.BoolOutput {
	return o.ApplyT(func(v GetmanagedAzResiliencyStatusResult) bool { return v.IsClusterZoneResilient }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(GetmanagedAzResiliencyStatusResultOutput{})
}
