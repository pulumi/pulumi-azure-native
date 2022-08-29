


package kubernetes

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListConnectedClusterUserCredentials(ctx *pulumi.Context, args *ListConnectedClusterUserCredentialsArgs, opts ...pulumi.InvokeOption) (*ListConnectedClusterUserCredentialsResult, error) {
	var rv ListConnectedClusterUserCredentialsResult
	err := ctx.Invoke("azure-native:kubernetes:listConnectedClusterUserCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListConnectedClusterUserCredentialsArgs struct {
	AuthenticationMethod string `pulumi:"authenticationMethod"`
	ClientProxy          bool   `pulumi:"clientProxy"`
	ClusterName          string `pulumi:"clusterName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type ListConnectedClusterUserCredentialsResult struct {
	HybridConnectionConfig HybridConnectionConfigResponse `pulumi:"hybridConnectionConfig"`
	Kubeconfigs            []CredentialResultResponse     `pulumi:"kubeconfigs"`
}

func ListConnectedClusterUserCredentialsOutput(ctx *pulumi.Context, args ListConnectedClusterUserCredentialsOutputArgs, opts ...pulumi.InvokeOption) ListConnectedClusterUserCredentialsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListConnectedClusterUserCredentialsResult, error) {
			args := v.(ListConnectedClusterUserCredentialsArgs)
			r, err := ListConnectedClusterUserCredentials(ctx, &args, opts...)
			var s ListConnectedClusterUserCredentialsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListConnectedClusterUserCredentialsResultOutput)
}

type ListConnectedClusterUserCredentialsOutputArgs struct {
	AuthenticationMethod pulumi.StringInput `pulumi:"authenticationMethod"`
	ClientProxy          pulumi.BoolInput   `pulumi:"clientProxy"`
	ClusterName          pulumi.StringInput `pulumi:"clusterName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListConnectedClusterUserCredentialsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListConnectedClusterUserCredentialsArgs)(nil)).Elem()
}


type ListConnectedClusterUserCredentialsResultOutput struct{ *pulumi.OutputState }

func (ListConnectedClusterUserCredentialsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListConnectedClusterUserCredentialsResult)(nil)).Elem()
}

func (o ListConnectedClusterUserCredentialsResultOutput) ToListConnectedClusterUserCredentialsResultOutput() ListConnectedClusterUserCredentialsResultOutput {
	return o
}

func (o ListConnectedClusterUserCredentialsResultOutput) ToListConnectedClusterUserCredentialsResultOutputWithContext(ctx context.Context) ListConnectedClusterUserCredentialsResultOutput {
	return o
}

func (o ListConnectedClusterUserCredentialsResultOutput) HybridConnectionConfig() HybridConnectionConfigResponseOutput {
	return o.ApplyT(func(v ListConnectedClusterUserCredentialsResult) HybridConnectionConfigResponse {
		return v.HybridConnectionConfig
	}).(HybridConnectionConfigResponseOutput)
}

func (o ListConnectedClusterUserCredentialsResultOutput) Kubeconfigs() CredentialResultResponseArrayOutput {
	return o.ApplyT(func(v ListConnectedClusterUserCredentialsResult) []CredentialResultResponse { return v.Kubeconfigs }).(CredentialResultResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListConnectedClusterUserCredentialsResultOutput{})
}
