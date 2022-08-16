


package v20220602preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListFleetCredentials(ctx *pulumi.Context, args *ListFleetCredentialsArgs, opts ...pulumi.InvokeOption) (*ListFleetCredentialsResult, error) {
	var rv ListFleetCredentialsResult
	err := ctx.Invoke("azure-native:containerservice/v20220602preview:listFleetCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListFleetCredentialsArgs struct {
	FleetName         string `pulumi:"fleetName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListFleetCredentialsResult struct {
	Kubeconfigs []FleetCredentialResultResponse `pulumi:"kubeconfigs"`
}

func ListFleetCredentialsOutput(ctx *pulumi.Context, args ListFleetCredentialsOutputArgs, opts ...pulumi.InvokeOption) ListFleetCredentialsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListFleetCredentialsResult, error) {
			args := v.(ListFleetCredentialsArgs)
			r, err := ListFleetCredentials(ctx, &args, opts...)
			var s ListFleetCredentialsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListFleetCredentialsResultOutput)
}

type ListFleetCredentialsOutputArgs struct {
	FleetName         pulumi.StringInput `pulumi:"fleetName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListFleetCredentialsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListFleetCredentialsArgs)(nil)).Elem()
}


type ListFleetCredentialsResultOutput struct{ *pulumi.OutputState }

func (ListFleetCredentialsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListFleetCredentialsResult)(nil)).Elem()
}

func (o ListFleetCredentialsResultOutput) ToListFleetCredentialsResultOutput() ListFleetCredentialsResultOutput {
	return o
}

func (o ListFleetCredentialsResultOutput) ToListFleetCredentialsResultOutputWithContext(ctx context.Context) ListFleetCredentialsResultOutput {
	return o
}

func (o ListFleetCredentialsResultOutput) Kubeconfigs() FleetCredentialResultResponseArrayOutput {
	return o.ApplyT(func(v ListFleetCredentialsResult) []FleetCredentialResultResponse { return v.Kubeconfigs }).(FleetCredentialResultResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListFleetCredentialsResultOutput{})
}
