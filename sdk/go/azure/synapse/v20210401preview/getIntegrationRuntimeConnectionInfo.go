


package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetIntegrationRuntimeConnectionInfo(ctx *pulumi.Context, args *GetIntegrationRuntimeConnectionInfoArgs, opts ...pulumi.InvokeOption) (*GetIntegrationRuntimeConnectionInfoResult, error) {
	var rv GetIntegrationRuntimeConnectionInfoResult
	err := ctx.Invoke("azure-native:synapse/v20210401preview:getIntegrationRuntimeConnectionInfo", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetIntegrationRuntimeConnectionInfoArgs struct {
	IntegrationRuntimeName string `pulumi:"integrationRuntimeName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	WorkspaceName          string `pulumi:"workspaceName"`
}


type GetIntegrationRuntimeConnectionInfoResult struct {
	HostServiceUri         string `pulumi:"hostServiceUri"`
	IdentityCertThumbprint string `pulumi:"identityCertThumbprint"`
	IsIdentityCertExprired bool   `pulumi:"isIdentityCertExprired"`
	PublicKey              string `pulumi:"publicKey"`
	ServiceToken           string `pulumi:"serviceToken"`
	Version                string `pulumi:"version"`
}

func GetIntegrationRuntimeConnectionInfoOutput(ctx *pulumi.Context, args GetIntegrationRuntimeConnectionInfoOutputArgs, opts ...pulumi.InvokeOption) GetIntegrationRuntimeConnectionInfoResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetIntegrationRuntimeConnectionInfoResult, error) {
			args := v.(GetIntegrationRuntimeConnectionInfoArgs)
			r, err := GetIntegrationRuntimeConnectionInfo(ctx, &args, opts...)
			var s GetIntegrationRuntimeConnectionInfoResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetIntegrationRuntimeConnectionInfoResultOutput)
}

type GetIntegrationRuntimeConnectionInfoOutputArgs struct {
	IntegrationRuntimeName pulumi.StringInput `pulumi:"integrationRuntimeName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName          pulumi.StringInput `pulumi:"workspaceName"`
}

func (GetIntegrationRuntimeConnectionInfoOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIntegrationRuntimeConnectionInfoArgs)(nil)).Elem()
}


type GetIntegrationRuntimeConnectionInfoResultOutput struct{ *pulumi.OutputState }

func (GetIntegrationRuntimeConnectionInfoResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIntegrationRuntimeConnectionInfoResult)(nil)).Elem()
}

func (o GetIntegrationRuntimeConnectionInfoResultOutput) ToGetIntegrationRuntimeConnectionInfoResultOutput() GetIntegrationRuntimeConnectionInfoResultOutput {
	return o
}

func (o GetIntegrationRuntimeConnectionInfoResultOutput) ToGetIntegrationRuntimeConnectionInfoResultOutputWithContext(ctx context.Context) GetIntegrationRuntimeConnectionInfoResultOutput {
	return o
}

func (o GetIntegrationRuntimeConnectionInfoResultOutput) HostServiceUri() pulumi.StringOutput {
	return o.ApplyT(func(v GetIntegrationRuntimeConnectionInfoResult) string { return v.HostServiceUri }).(pulumi.StringOutput)
}

func (o GetIntegrationRuntimeConnectionInfoResultOutput) IdentityCertThumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v GetIntegrationRuntimeConnectionInfoResult) string { return v.IdentityCertThumbprint }).(pulumi.StringOutput)
}

func (o GetIntegrationRuntimeConnectionInfoResultOutput) IsIdentityCertExprired() pulumi.BoolOutput {
	return o.ApplyT(func(v GetIntegrationRuntimeConnectionInfoResult) bool { return v.IsIdentityCertExprired }).(pulumi.BoolOutput)
}

func (o GetIntegrationRuntimeConnectionInfoResultOutput) PublicKey() pulumi.StringOutput {
	return o.ApplyT(func(v GetIntegrationRuntimeConnectionInfoResult) string { return v.PublicKey }).(pulumi.StringOutput)
}

func (o GetIntegrationRuntimeConnectionInfoResultOutput) ServiceToken() pulumi.StringOutput {
	return o.ApplyT(func(v GetIntegrationRuntimeConnectionInfoResult) string { return v.ServiceToken }).(pulumi.StringOutput)
}

func (o GetIntegrationRuntimeConnectionInfoResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v GetIntegrationRuntimeConnectionInfoResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIntegrationRuntimeConnectionInfoResultOutput{})
}
