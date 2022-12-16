


package hdinsight

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetClusterGatewaySettings(ctx *pulumi.Context, args *GetClusterGatewaySettingsArgs, opts ...pulumi.InvokeOption) (*GetClusterGatewaySettingsResult, error) {
	var rv GetClusterGatewaySettingsResult
	err := ctx.Invoke("azure-native:hdinsight:getClusterGatewaySettings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetClusterGatewaySettingsArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type GetClusterGatewaySettingsResult struct {
	IsCredentialEnabled string `pulumi:"isCredentialEnabled"`
	Password            string `pulumi:"password"`
	UserName            string `pulumi:"userName"`
}

func GetClusterGatewaySettingsOutput(ctx *pulumi.Context, args GetClusterGatewaySettingsOutputArgs, opts ...pulumi.InvokeOption) GetClusterGatewaySettingsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetClusterGatewaySettingsResult, error) {
			args := v.(GetClusterGatewaySettingsArgs)
			r, err := GetClusterGatewaySettings(ctx, &args, opts...)
			var s GetClusterGatewaySettingsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetClusterGatewaySettingsResultOutput)
}

type GetClusterGatewaySettingsOutputArgs struct {
	ClusterName       pulumi.StringInput `pulumi:"clusterName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetClusterGatewaySettingsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClusterGatewaySettingsArgs)(nil)).Elem()
}


type GetClusterGatewaySettingsResultOutput struct{ *pulumi.OutputState }

func (GetClusterGatewaySettingsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClusterGatewaySettingsResult)(nil)).Elem()
}

func (o GetClusterGatewaySettingsResultOutput) ToGetClusterGatewaySettingsResultOutput() GetClusterGatewaySettingsResultOutput {
	return o
}

func (o GetClusterGatewaySettingsResultOutput) ToGetClusterGatewaySettingsResultOutputWithContext(ctx context.Context) GetClusterGatewaySettingsResultOutput {
	return o
}

func (o GetClusterGatewaySettingsResultOutput) IsCredentialEnabled() pulumi.StringOutput {
	return o.ApplyT(func(v GetClusterGatewaySettingsResult) string { return v.IsCredentialEnabled }).(pulumi.StringOutput)
}

func (o GetClusterGatewaySettingsResultOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v GetClusterGatewaySettingsResult) string { return v.Password }).(pulumi.StringOutput)
}

func (o GetClusterGatewaySettingsResultOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v GetClusterGatewaySettingsResult) string { return v.UserName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetClusterGatewaySettingsResultOutput{})
}
