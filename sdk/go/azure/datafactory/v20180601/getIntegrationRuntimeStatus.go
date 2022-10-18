


package v20180601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetIntegrationRuntimeStatus(ctx *pulumi.Context, args *GetIntegrationRuntimeStatusArgs, opts ...pulumi.InvokeOption) (*GetIntegrationRuntimeStatusResult, error) {
	var rv GetIntegrationRuntimeStatusResult
	err := ctx.Invoke("azure-native:datafactory/v20180601:getIntegrationRuntimeStatus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetIntegrationRuntimeStatusArgs struct {
	FactoryName            string `pulumi:"factoryName"`
	IntegrationRuntimeName string `pulumi:"integrationRuntimeName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
}


type GetIntegrationRuntimeStatusResult struct {
	Name       string      `pulumi:"name"`
	Properties interface{} `pulumi:"properties"`
}

func GetIntegrationRuntimeStatusOutput(ctx *pulumi.Context, args GetIntegrationRuntimeStatusOutputArgs, opts ...pulumi.InvokeOption) GetIntegrationRuntimeStatusResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetIntegrationRuntimeStatusResult, error) {
			args := v.(GetIntegrationRuntimeStatusArgs)
			r, err := GetIntegrationRuntimeStatus(ctx, &args, opts...)
			var s GetIntegrationRuntimeStatusResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetIntegrationRuntimeStatusResultOutput)
}

type GetIntegrationRuntimeStatusOutputArgs struct {
	FactoryName            pulumi.StringInput `pulumi:"factoryName"`
	IntegrationRuntimeName pulumi.StringInput `pulumi:"integrationRuntimeName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetIntegrationRuntimeStatusOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIntegrationRuntimeStatusArgs)(nil)).Elem()
}


type GetIntegrationRuntimeStatusResultOutput struct{ *pulumi.OutputState }

func (GetIntegrationRuntimeStatusResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIntegrationRuntimeStatusResult)(nil)).Elem()
}

func (o GetIntegrationRuntimeStatusResultOutput) ToGetIntegrationRuntimeStatusResultOutput() GetIntegrationRuntimeStatusResultOutput {
	return o
}

func (o GetIntegrationRuntimeStatusResultOutput) ToGetIntegrationRuntimeStatusResultOutputWithContext(ctx context.Context) GetIntegrationRuntimeStatusResultOutput {
	return o
}

func (o GetIntegrationRuntimeStatusResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetIntegrationRuntimeStatusResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetIntegrationRuntimeStatusResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v GetIntegrationRuntimeStatusResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIntegrationRuntimeStatusResultOutput{})
}
