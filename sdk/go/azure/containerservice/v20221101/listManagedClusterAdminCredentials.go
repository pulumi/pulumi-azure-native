


package v20221101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListManagedClusterAdminCredentials(ctx *pulumi.Context, args *ListManagedClusterAdminCredentialsArgs, opts ...pulumi.InvokeOption) (*ListManagedClusterAdminCredentialsResult, error) {
	var rv ListManagedClusterAdminCredentialsResult
	err := ctx.Invoke("azure-native:containerservice/v20221101:listManagedClusterAdminCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListManagedClusterAdminCredentialsArgs struct {
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ResourceName      string  `pulumi:"resourceName"`
	ServerFqdn        *string `pulumi:"serverFqdn"`
}


type ListManagedClusterAdminCredentialsResult struct {
	Kubeconfigs []CredentialResultResponse `pulumi:"kubeconfigs"`
}

func ListManagedClusterAdminCredentialsOutput(ctx *pulumi.Context, args ListManagedClusterAdminCredentialsOutputArgs, opts ...pulumi.InvokeOption) ListManagedClusterAdminCredentialsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListManagedClusterAdminCredentialsResult, error) {
			args := v.(ListManagedClusterAdminCredentialsArgs)
			r, err := ListManagedClusterAdminCredentials(ctx, &args, opts...)
			var s ListManagedClusterAdminCredentialsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListManagedClusterAdminCredentialsResultOutput)
}

type ListManagedClusterAdminCredentialsOutputArgs struct {
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput    `pulumi:"resourceName"`
	ServerFqdn        pulumi.StringPtrInput `pulumi:"serverFqdn"`
}

func (ListManagedClusterAdminCredentialsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListManagedClusterAdminCredentialsArgs)(nil)).Elem()
}


type ListManagedClusterAdminCredentialsResultOutput struct{ *pulumi.OutputState }

func (ListManagedClusterAdminCredentialsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListManagedClusterAdminCredentialsResult)(nil)).Elem()
}

func (o ListManagedClusterAdminCredentialsResultOutput) ToListManagedClusterAdminCredentialsResultOutput() ListManagedClusterAdminCredentialsResultOutput {
	return o
}

func (o ListManagedClusterAdminCredentialsResultOutput) ToListManagedClusterAdminCredentialsResultOutputWithContext(ctx context.Context) ListManagedClusterAdminCredentialsResultOutput {
	return o
}

func (o ListManagedClusterAdminCredentialsResultOutput) Kubeconfigs() CredentialResultResponseArrayOutput {
	return o.ApplyT(func(v ListManagedClusterAdminCredentialsResult) []CredentialResultResponse { return v.Kubeconfigs }).(CredentialResultResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListManagedClusterAdminCredentialsResultOutput{})
}
