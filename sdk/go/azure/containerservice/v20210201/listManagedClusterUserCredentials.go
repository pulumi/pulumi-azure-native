


package v20210201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListManagedClusterUserCredentials(ctx *pulumi.Context, args *ListManagedClusterUserCredentialsArgs, opts ...pulumi.InvokeOption) (*ListManagedClusterUserCredentialsResult, error) {
	var rv ListManagedClusterUserCredentialsResult
	err := ctx.Invoke("azure-native:containerservice/v20210201:listManagedClusterUserCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListManagedClusterUserCredentialsArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type ListManagedClusterUserCredentialsResult struct {
	Kubeconfigs []CredentialResultResponse `pulumi:"kubeconfigs"`
}

func ListManagedClusterUserCredentialsOutput(ctx *pulumi.Context, args ListManagedClusterUserCredentialsOutputArgs, opts ...pulumi.InvokeOption) ListManagedClusterUserCredentialsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListManagedClusterUserCredentialsResult, error) {
			args := v.(ListManagedClusterUserCredentialsArgs)
			r, err := ListManagedClusterUserCredentials(ctx, &args, opts...)
			var s ListManagedClusterUserCredentialsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListManagedClusterUserCredentialsResultOutput)
}

type ListManagedClusterUserCredentialsOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (ListManagedClusterUserCredentialsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListManagedClusterUserCredentialsArgs)(nil)).Elem()
}


type ListManagedClusterUserCredentialsResultOutput struct{ *pulumi.OutputState }

func (ListManagedClusterUserCredentialsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListManagedClusterUserCredentialsResult)(nil)).Elem()
}

func (o ListManagedClusterUserCredentialsResultOutput) ToListManagedClusterUserCredentialsResultOutput() ListManagedClusterUserCredentialsResultOutput {
	return o
}

func (o ListManagedClusterUserCredentialsResultOutput) ToListManagedClusterUserCredentialsResultOutputWithContext(ctx context.Context) ListManagedClusterUserCredentialsResultOutput {
	return o
}

func (o ListManagedClusterUserCredentialsResultOutput) Kubeconfigs() CredentialResultResponseArrayOutput {
	return o.ApplyT(func(v ListManagedClusterUserCredentialsResult) []CredentialResultResponse { return v.Kubeconfigs }).(CredentialResultResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListManagedClusterUserCredentialsResultOutput{})
}
