


package migrate

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetSolutionConfig(ctx *pulumi.Context, args *GetSolutionConfigArgs, opts ...pulumi.InvokeOption) (*GetSolutionConfigResult, error) {
	var rv GetSolutionConfigResult
	err := ctx.Invoke("azure-native:migrate:getSolutionConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetSolutionConfigArgs struct {
	MigrateProjectName string `pulumi:"migrateProjectName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	SolutionName       string `pulumi:"solutionName"`
}


type GetSolutionConfigResult struct {
	PublisherSasUri *string `pulumi:"publisherSasUri"`
}

func GetSolutionConfigOutput(ctx *pulumi.Context, args GetSolutionConfigOutputArgs, opts ...pulumi.InvokeOption) GetSolutionConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSolutionConfigResult, error) {
			args := v.(GetSolutionConfigArgs)
			r, err := GetSolutionConfig(ctx, &args, opts...)
			var s GetSolutionConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSolutionConfigResultOutput)
}

type GetSolutionConfigOutputArgs struct {
	MigrateProjectName pulumi.StringInput `pulumi:"migrateProjectName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
	SolutionName       pulumi.StringInput `pulumi:"solutionName"`
}

func (GetSolutionConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSolutionConfigArgs)(nil)).Elem()
}


type GetSolutionConfigResultOutput struct{ *pulumi.OutputState }

func (GetSolutionConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSolutionConfigResult)(nil)).Elem()
}

func (o GetSolutionConfigResultOutput) ToGetSolutionConfigResultOutput() GetSolutionConfigResultOutput {
	return o
}

func (o GetSolutionConfigResultOutput) ToGetSolutionConfigResultOutputWithContext(ctx context.Context) GetSolutionConfigResultOutput {
	return o
}

func (o GetSolutionConfigResultOutput) PublisherSasUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSolutionConfigResult) *string { return v.PublisherSasUri }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSolutionConfigResultOutput{})
}
