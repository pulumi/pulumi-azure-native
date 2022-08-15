


package v20220101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListNetworkManagerDeploymentStatus(ctx *pulumi.Context, args *ListNetworkManagerDeploymentStatusArgs, opts ...pulumi.InvokeOption) (*ListNetworkManagerDeploymentStatusResult, error) {
	var rv ListNetworkManagerDeploymentStatusResult
	err := ctx.Invoke("azure-native:network/v20220101:listNetworkManagerDeploymentStatus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListNetworkManagerDeploymentStatusArgs struct {
	DeploymentTypes    []string `pulumi:"deploymentTypes"`
	NetworkManagerName string   `pulumi:"networkManagerName"`
	Regions            []string `pulumi:"regions"`
	ResourceGroupName  string   `pulumi:"resourceGroupName"`
	SkipToken          *string  `pulumi:"skipToken"`
}


type ListNetworkManagerDeploymentStatusResult struct {
	SkipToken *string                                  `pulumi:"skipToken"`
	Value     []NetworkManagerDeploymentStatusResponse `pulumi:"value"`
}

func ListNetworkManagerDeploymentStatusOutput(ctx *pulumi.Context, args ListNetworkManagerDeploymentStatusOutputArgs, opts ...pulumi.InvokeOption) ListNetworkManagerDeploymentStatusResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListNetworkManagerDeploymentStatusResult, error) {
			args := v.(ListNetworkManagerDeploymentStatusArgs)
			r, err := ListNetworkManagerDeploymentStatus(ctx, &args, opts...)
			var s ListNetworkManagerDeploymentStatusResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListNetworkManagerDeploymentStatusResultOutput)
}

type ListNetworkManagerDeploymentStatusOutputArgs struct {
	DeploymentTypes    pulumi.StringArrayInput `pulumi:"deploymentTypes"`
	NetworkManagerName pulumi.StringInput      `pulumi:"networkManagerName"`
	Regions            pulumi.StringArrayInput `pulumi:"regions"`
	ResourceGroupName  pulumi.StringInput      `pulumi:"resourceGroupName"`
	SkipToken          pulumi.StringPtrInput   `pulumi:"skipToken"`
}

func (ListNetworkManagerDeploymentStatusOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListNetworkManagerDeploymentStatusArgs)(nil)).Elem()
}


type ListNetworkManagerDeploymentStatusResultOutput struct{ *pulumi.OutputState }

func (ListNetworkManagerDeploymentStatusResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListNetworkManagerDeploymentStatusResult)(nil)).Elem()
}

func (o ListNetworkManagerDeploymentStatusResultOutput) ToListNetworkManagerDeploymentStatusResultOutput() ListNetworkManagerDeploymentStatusResultOutput {
	return o
}

func (o ListNetworkManagerDeploymentStatusResultOutput) ToListNetworkManagerDeploymentStatusResultOutputWithContext(ctx context.Context) ListNetworkManagerDeploymentStatusResultOutput {
	return o
}

func (o ListNetworkManagerDeploymentStatusResultOutput) SkipToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListNetworkManagerDeploymentStatusResult) *string { return v.SkipToken }).(pulumi.StringPtrOutput)
}

func (o ListNetworkManagerDeploymentStatusResultOutput) Value() NetworkManagerDeploymentStatusResponseArrayOutput {
	return o.ApplyT(func(v ListNetworkManagerDeploymentStatusResult) []NetworkManagerDeploymentStatusResponse {
		return v.Value
	}).(NetworkManagerDeploymentStatusResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListNetworkManagerDeploymentStatusResultOutput{})
}
