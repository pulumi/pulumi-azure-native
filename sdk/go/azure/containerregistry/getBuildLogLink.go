


package containerregistry

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetBuildLogLink(ctx *pulumi.Context, args *GetBuildLogLinkArgs, opts ...pulumi.InvokeOption) (*GetBuildLogLinkResult, error) {
	var rv GetBuildLogLinkResult
	err := ctx.Invoke("azure-native:containerregistry:getBuildLogLink", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetBuildLogLinkArgs struct {
	BuildId           string `pulumi:"buildId"`
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type GetBuildLogLinkResult struct {
	LogLink *string `pulumi:"logLink"`
}

func GetBuildLogLinkOutput(ctx *pulumi.Context, args GetBuildLogLinkOutputArgs, opts ...pulumi.InvokeOption) GetBuildLogLinkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetBuildLogLinkResult, error) {
			args := v.(GetBuildLogLinkArgs)
			r, err := GetBuildLogLink(ctx, &args, opts...)
			var s GetBuildLogLinkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetBuildLogLinkResultOutput)
}

type GetBuildLogLinkOutputArgs struct {
	BuildId           pulumi.StringInput `pulumi:"buildId"`
	RegistryName      pulumi.StringInput `pulumi:"registryName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetBuildLogLinkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBuildLogLinkArgs)(nil)).Elem()
}


type GetBuildLogLinkResultOutput struct{ *pulumi.OutputState }

func (GetBuildLogLinkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBuildLogLinkResult)(nil)).Elem()
}

func (o GetBuildLogLinkResultOutput) ToGetBuildLogLinkResultOutput() GetBuildLogLinkResultOutput {
	return o
}

func (o GetBuildLogLinkResultOutput) ToGetBuildLogLinkResultOutputWithContext(ctx context.Context) GetBuildLogLinkResultOutput {
	return o
}

func (o GetBuildLogLinkResultOutput) LogLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBuildLogLinkResult) *string { return v.LogLink }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBuildLogLinkResultOutput{})
}
