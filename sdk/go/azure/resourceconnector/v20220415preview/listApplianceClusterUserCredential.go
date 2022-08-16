


package v20220415preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListApplianceClusterUserCredential(ctx *pulumi.Context, args *ListApplianceClusterUserCredentialArgs, opts ...pulumi.InvokeOption) (*ListApplianceClusterUserCredentialResult, error) {
	var rv ListApplianceClusterUserCredentialResult
	err := ctx.Invoke("azure-native:resourceconnector/v20220415preview:listApplianceClusterUserCredential", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListApplianceClusterUserCredentialArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type ListApplianceClusterUserCredentialResult struct {
	HybridConnectionConfig HybridConnectionConfigResponse          `pulumi:"hybridConnectionConfig"`
	Kubeconfigs            []ApplianceCredentialKubeconfigResponse `pulumi:"kubeconfigs"`
}

func ListApplianceClusterUserCredentialOutput(ctx *pulumi.Context, args ListApplianceClusterUserCredentialOutputArgs, opts ...pulumi.InvokeOption) ListApplianceClusterUserCredentialResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListApplianceClusterUserCredentialResult, error) {
			args := v.(ListApplianceClusterUserCredentialArgs)
			r, err := ListApplianceClusterUserCredential(ctx, &args, opts...)
			var s ListApplianceClusterUserCredentialResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListApplianceClusterUserCredentialResultOutput)
}

type ListApplianceClusterUserCredentialOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (ListApplianceClusterUserCredentialOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListApplianceClusterUserCredentialArgs)(nil)).Elem()
}


type ListApplianceClusterUserCredentialResultOutput struct{ *pulumi.OutputState }

func (ListApplianceClusterUserCredentialResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListApplianceClusterUserCredentialResult)(nil)).Elem()
}

func (o ListApplianceClusterUserCredentialResultOutput) ToListApplianceClusterUserCredentialResultOutput() ListApplianceClusterUserCredentialResultOutput {
	return o
}

func (o ListApplianceClusterUserCredentialResultOutput) ToListApplianceClusterUserCredentialResultOutputWithContext(ctx context.Context) ListApplianceClusterUserCredentialResultOutput {
	return o
}

func (o ListApplianceClusterUserCredentialResultOutput) HybridConnectionConfig() HybridConnectionConfigResponseOutput {
	return o.ApplyT(func(v ListApplianceClusterUserCredentialResult) HybridConnectionConfigResponse {
		return v.HybridConnectionConfig
	}).(HybridConnectionConfigResponseOutput)
}

func (o ListApplianceClusterUserCredentialResultOutput) Kubeconfigs() ApplianceCredentialKubeconfigResponseArrayOutput {
	return o.ApplyT(func(v ListApplianceClusterUserCredentialResult) []ApplianceCredentialKubeconfigResponse {
		return v.Kubeconfigs
	}).(ApplianceCredentialKubeconfigResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListApplianceClusterUserCredentialResultOutput{})
}
