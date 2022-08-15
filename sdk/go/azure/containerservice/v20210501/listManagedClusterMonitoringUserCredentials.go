


package v20210501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListManagedClusterMonitoringUserCredentials(ctx *pulumi.Context, args *ListManagedClusterMonitoringUserCredentialsArgs, opts ...pulumi.InvokeOption) (*ListManagedClusterMonitoringUserCredentialsResult, error) {
	var rv ListManagedClusterMonitoringUserCredentialsResult
	err := ctx.Invoke("azure-native:containerservice/v20210501:listManagedClusterMonitoringUserCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListManagedClusterMonitoringUserCredentialsArgs struct {
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ResourceName      string  `pulumi:"resourceName"`
	ServerFqdn        *string `pulumi:"serverFqdn"`
}


type ListManagedClusterMonitoringUserCredentialsResult struct {
	Kubeconfigs []CredentialResultResponse `pulumi:"kubeconfigs"`
}

func ListManagedClusterMonitoringUserCredentialsOutput(ctx *pulumi.Context, args ListManagedClusterMonitoringUserCredentialsOutputArgs, opts ...pulumi.InvokeOption) ListManagedClusterMonitoringUserCredentialsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListManagedClusterMonitoringUserCredentialsResult, error) {
			args := v.(ListManagedClusterMonitoringUserCredentialsArgs)
			r, err := ListManagedClusterMonitoringUserCredentials(ctx, &args, opts...)
			var s ListManagedClusterMonitoringUserCredentialsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListManagedClusterMonitoringUserCredentialsResultOutput)
}

type ListManagedClusterMonitoringUserCredentialsOutputArgs struct {
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput    `pulumi:"resourceName"`
	ServerFqdn        pulumi.StringPtrInput `pulumi:"serverFqdn"`
}

func (ListManagedClusterMonitoringUserCredentialsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListManagedClusterMonitoringUserCredentialsArgs)(nil)).Elem()
}


type ListManagedClusterMonitoringUserCredentialsResultOutput struct{ *pulumi.OutputState }

func (ListManagedClusterMonitoringUserCredentialsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListManagedClusterMonitoringUserCredentialsResult)(nil)).Elem()
}

func (o ListManagedClusterMonitoringUserCredentialsResultOutput) ToListManagedClusterMonitoringUserCredentialsResultOutput() ListManagedClusterMonitoringUserCredentialsResultOutput {
	return o
}

func (o ListManagedClusterMonitoringUserCredentialsResultOutput) ToListManagedClusterMonitoringUserCredentialsResultOutputWithContext(ctx context.Context) ListManagedClusterMonitoringUserCredentialsResultOutput {
	return o
}

func (o ListManagedClusterMonitoringUserCredentialsResultOutput) Kubeconfigs() CredentialResultResponseArrayOutput {
	return o.ApplyT(func(v ListManagedClusterMonitoringUserCredentialsResult) []CredentialResultResponse {
		return v.Kubeconfigs
	}).(CredentialResultResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListManagedClusterMonitoringUserCredentialsResultOutput{})
}
