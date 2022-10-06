


package v20220901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetBuildServiceBuildResultLog(ctx *pulumi.Context, args *GetBuildServiceBuildResultLogArgs, opts ...pulumi.InvokeOption) (*GetBuildServiceBuildResultLogResult, error) {
	var rv GetBuildServiceBuildResultLogResult
	err := ctx.Invoke("azure-native:appplatform/v20220901preview:getBuildServiceBuildResultLog", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetBuildServiceBuildResultLogArgs struct {
	BuildName         string `pulumi:"buildName"`
	BuildResultName   string `pulumi:"buildResultName"`
	BuildServiceName  string `pulumi:"buildServiceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type GetBuildServiceBuildResultLogResult struct {
	BlobUrl *string `pulumi:"blobUrl"`
}

func GetBuildServiceBuildResultLogOutput(ctx *pulumi.Context, args GetBuildServiceBuildResultLogOutputArgs, opts ...pulumi.InvokeOption) GetBuildServiceBuildResultLogResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetBuildServiceBuildResultLogResult, error) {
			args := v.(GetBuildServiceBuildResultLogArgs)
			r, err := GetBuildServiceBuildResultLog(ctx, &args, opts...)
			var s GetBuildServiceBuildResultLogResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetBuildServiceBuildResultLogResultOutput)
}

type GetBuildServiceBuildResultLogOutputArgs struct {
	BuildName         pulumi.StringInput `pulumi:"buildName"`
	BuildResultName   pulumi.StringInput `pulumi:"buildResultName"`
	BuildServiceName  pulumi.StringInput `pulumi:"buildServiceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (GetBuildServiceBuildResultLogOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBuildServiceBuildResultLogArgs)(nil)).Elem()
}


type GetBuildServiceBuildResultLogResultOutput struct{ *pulumi.OutputState }

func (GetBuildServiceBuildResultLogResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBuildServiceBuildResultLogResult)(nil)).Elem()
}

func (o GetBuildServiceBuildResultLogResultOutput) ToGetBuildServiceBuildResultLogResultOutput() GetBuildServiceBuildResultLogResultOutput {
	return o
}

func (o GetBuildServiceBuildResultLogResultOutput) ToGetBuildServiceBuildResultLogResultOutputWithContext(ctx context.Context) GetBuildServiceBuildResultLogResultOutput {
	return o
}

func (o GetBuildServiceBuildResultLogResultOutput) BlobUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBuildServiceBuildResultLogResult) *string { return v.BlobUrl }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBuildServiceBuildResultLogResultOutput{})
}
