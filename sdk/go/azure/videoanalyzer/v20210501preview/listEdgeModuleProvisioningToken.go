


package v20210501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListEdgeModuleProvisioningToken(ctx *pulumi.Context, args *ListEdgeModuleProvisioningTokenArgs, opts ...pulumi.InvokeOption) (*ListEdgeModuleProvisioningTokenResult, error) {
	var rv ListEdgeModuleProvisioningTokenResult
	err := ctx.Invoke("azure-native:videoanalyzer/v20210501preview:listEdgeModuleProvisioningToken", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListEdgeModuleProvisioningTokenArgs struct {
	AccountName       string `pulumi:"accountName"`
	EdgeModuleName    string `pulumi:"edgeModuleName"`
	ExpirationDate    string `pulumi:"expirationDate"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListEdgeModuleProvisioningTokenResult struct {
	ExpirationDate string `pulumi:"expirationDate"`
	Token          string `pulumi:"token"`
}

func ListEdgeModuleProvisioningTokenOutput(ctx *pulumi.Context, args ListEdgeModuleProvisioningTokenOutputArgs, opts ...pulumi.InvokeOption) ListEdgeModuleProvisioningTokenResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListEdgeModuleProvisioningTokenResult, error) {
			args := v.(ListEdgeModuleProvisioningTokenArgs)
			r, err := ListEdgeModuleProvisioningToken(ctx, &args, opts...)
			var s ListEdgeModuleProvisioningTokenResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListEdgeModuleProvisioningTokenResultOutput)
}

type ListEdgeModuleProvisioningTokenOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	EdgeModuleName    pulumi.StringInput `pulumi:"edgeModuleName"`
	ExpirationDate    pulumi.StringInput `pulumi:"expirationDate"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListEdgeModuleProvisioningTokenOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListEdgeModuleProvisioningTokenArgs)(nil)).Elem()
}


type ListEdgeModuleProvisioningTokenResultOutput struct{ *pulumi.OutputState }

func (ListEdgeModuleProvisioningTokenResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListEdgeModuleProvisioningTokenResult)(nil)).Elem()
}

func (o ListEdgeModuleProvisioningTokenResultOutput) ToListEdgeModuleProvisioningTokenResultOutput() ListEdgeModuleProvisioningTokenResultOutput {
	return o
}

func (o ListEdgeModuleProvisioningTokenResultOutput) ToListEdgeModuleProvisioningTokenResultOutputWithContext(ctx context.Context) ListEdgeModuleProvisioningTokenResultOutput {
	return o
}

func (o ListEdgeModuleProvisioningTokenResultOutput) ExpirationDate() pulumi.StringOutput {
	return o.ApplyT(func(v ListEdgeModuleProvisioningTokenResult) string { return v.ExpirationDate }).(pulumi.StringOutput)
}

func (o ListEdgeModuleProvisioningTokenResultOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v ListEdgeModuleProvisioningTokenResult) string { return v.Token }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListEdgeModuleProvisioningTokenResultOutput{})
}
