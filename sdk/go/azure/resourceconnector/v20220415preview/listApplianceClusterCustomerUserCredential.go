


package v20220415preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListApplianceClusterCustomerUserCredential(ctx *pulumi.Context, args *ListApplianceClusterCustomerUserCredentialArgs, opts ...pulumi.InvokeOption) (*ListApplianceClusterCustomerUserCredentialResult, error) {
	var rv ListApplianceClusterCustomerUserCredentialResult
	err := ctx.Invoke("azure-native:resourceconnector/v20220415preview:listApplianceClusterCustomerUserCredential", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListApplianceClusterCustomerUserCredentialArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type ListApplianceClusterCustomerUserCredentialResult struct {
	Kubeconfigs []ApplianceCredentialKubeconfigResponse `pulumi:"kubeconfigs"`
	SshKeys     map[string]SSHKeyResponse               `pulumi:"sshKeys"`
}

func ListApplianceClusterCustomerUserCredentialOutput(ctx *pulumi.Context, args ListApplianceClusterCustomerUserCredentialOutputArgs, opts ...pulumi.InvokeOption) ListApplianceClusterCustomerUserCredentialResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListApplianceClusterCustomerUserCredentialResult, error) {
			args := v.(ListApplianceClusterCustomerUserCredentialArgs)
			r, err := ListApplianceClusterCustomerUserCredential(ctx, &args, opts...)
			var s ListApplianceClusterCustomerUserCredentialResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListApplianceClusterCustomerUserCredentialResultOutput)
}

type ListApplianceClusterCustomerUserCredentialOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (ListApplianceClusterCustomerUserCredentialOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListApplianceClusterCustomerUserCredentialArgs)(nil)).Elem()
}


type ListApplianceClusterCustomerUserCredentialResultOutput struct{ *pulumi.OutputState }

func (ListApplianceClusterCustomerUserCredentialResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListApplianceClusterCustomerUserCredentialResult)(nil)).Elem()
}

func (o ListApplianceClusterCustomerUserCredentialResultOutput) ToListApplianceClusterCustomerUserCredentialResultOutput() ListApplianceClusterCustomerUserCredentialResultOutput {
	return o
}

func (o ListApplianceClusterCustomerUserCredentialResultOutput) ToListApplianceClusterCustomerUserCredentialResultOutputWithContext(ctx context.Context) ListApplianceClusterCustomerUserCredentialResultOutput {
	return o
}

func (o ListApplianceClusterCustomerUserCredentialResultOutput) Kubeconfigs() ApplianceCredentialKubeconfigResponseArrayOutput {
	return o.ApplyT(func(v ListApplianceClusterCustomerUserCredentialResult) []ApplianceCredentialKubeconfigResponse {
		return v.Kubeconfigs
	}).(ApplianceCredentialKubeconfigResponseArrayOutput)
}

func (o ListApplianceClusterCustomerUserCredentialResultOutput) SshKeys() SSHKeyResponseMapOutput {
	return o.ApplyT(func(v ListApplianceClusterCustomerUserCredentialResult) map[string]SSHKeyResponse { return v.SshKeys }).(SSHKeyResponseMapOutput)
}

func init() {
	pulumi.RegisterOutputType(ListApplianceClusterCustomerUserCredentialResultOutput{})
}
