


package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListConnectedClusterUserCredential(ctx *pulumi.Context, args *ListConnectedClusterUserCredentialArgs, opts ...pulumi.InvokeOption) (*ListConnectedClusterUserCredentialResult, error) {
	var rv ListConnectedClusterUserCredentialResult
	err := ctx.Invoke("azure-native:kubernetes/v20220501preview:listConnectedClusterUserCredential", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListConnectedClusterUserCredentialArgs struct {
	AuthenticationMethod string `pulumi:"authenticationMethod"`
	ClientProxy          bool   `pulumi:"clientProxy"`
	ClusterName          string `pulumi:"clusterName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type ListConnectedClusterUserCredentialResult struct {
	HybridConnectionConfig HybridConnectionConfigResponse `pulumi:"hybridConnectionConfig"`
	Kubeconfigs            []CredentialResultResponse     `pulumi:"kubeconfigs"`
}

func ListConnectedClusterUserCredentialOutput(ctx *pulumi.Context, args ListConnectedClusterUserCredentialOutputArgs, opts ...pulumi.InvokeOption) ListConnectedClusterUserCredentialResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListConnectedClusterUserCredentialResult, error) {
			args := v.(ListConnectedClusterUserCredentialArgs)
			r, err := ListConnectedClusterUserCredential(ctx, &args, opts...)
			var s ListConnectedClusterUserCredentialResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListConnectedClusterUserCredentialResultOutput)
}

type ListConnectedClusterUserCredentialOutputArgs struct {
	AuthenticationMethod pulumi.StringInput `pulumi:"authenticationMethod"`
	ClientProxy          pulumi.BoolInput   `pulumi:"clientProxy"`
	ClusterName          pulumi.StringInput `pulumi:"clusterName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListConnectedClusterUserCredentialOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListConnectedClusterUserCredentialArgs)(nil)).Elem()
}


type ListConnectedClusterUserCredentialResultOutput struct{ *pulumi.OutputState }

func (ListConnectedClusterUserCredentialResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListConnectedClusterUserCredentialResult)(nil)).Elem()
}

func (o ListConnectedClusterUserCredentialResultOutput) ToListConnectedClusterUserCredentialResultOutput() ListConnectedClusterUserCredentialResultOutput {
	return o
}

func (o ListConnectedClusterUserCredentialResultOutput) ToListConnectedClusterUserCredentialResultOutputWithContext(ctx context.Context) ListConnectedClusterUserCredentialResultOutput {
	return o
}

func (o ListConnectedClusterUserCredentialResultOutput) HybridConnectionConfig() HybridConnectionConfigResponseOutput {
	return o.ApplyT(func(v ListConnectedClusterUserCredentialResult) HybridConnectionConfigResponse {
		return v.HybridConnectionConfig
	}).(HybridConnectionConfigResponseOutput)
}

func (o ListConnectedClusterUserCredentialResultOutput) Kubeconfigs() CredentialResultResponseArrayOutput {
	return o.ApplyT(func(v ListConnectedClusterUserCredentialResult) []CredentialResultResponse { return v.Kubeconfigs }).(CredentialResultResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListConnectedClusterUserCredentialResultOutput{})
}
