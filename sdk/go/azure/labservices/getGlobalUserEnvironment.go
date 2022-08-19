


package labservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetGlobalUserEnvironment(ctx *pulumi.Context, args *GetGlobalUserEnvironmentArgs, opts ...pulumi.InvokeOption) (*GetGlobalUserEnvironmentResult, error) {
	var rv GetGlobalUserEnvironmentResult
	err := ctx.Invoke("azure-native:labservices:getGlobalUserEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetGlobalUserEnvironmentArgs struct {
	EnvironmentId string  `pulumi:"environmentId"`
	Expand        *string `pulumi:"expand"`
	UserName      string  `pulumi:"userName"`
}


type GetGlobalUserEnvironmentResult struct {
	Environment EnvironmentDetailsResponse `pulumi:"environment"`
}

func GetGlobalUserEnvironmentOutput(ctx *pulumi.Context, args GetGlobalUserEnvironmentOutputArgs, opts ...pulumi.InvokeOption) GetGlobalUserEnvironmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetGlobalUserEnvironmentResult, error) {
			args := v.(GetGlobalUserEnvironmentArgs)
			r, err := GetGlobalUserEnvironment(ctx, &args, opts...)
			var s GetGlobalUserEnvironmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetGlobalUserEnvironmentResultOutput)
}

type GetGlobalUserEnvironmentOutputArgs struct {
	EnvironmentId pulumi.StringInput    `pulumi:"environmentId"`
	Expand        pulumi.StringPtrInput `pulumi:"expand"`
	UserName      pulumi.StringInput    `pulumi:"userName"`
}

func (GetGlobalUserEnvironmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGlobalUserEnvironmentArgs)(nil)).Elem()
}


type GetGlobalUserEnvironmentResultOutput struct{ *pulumi.OutputState }

func (GetGlobalUserEnvironmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGlobalUserEnvironmentResult)(nil)).Elem()
}

func (o GetGlobalUserEnvironmentResultOutput) ToGetGlobalUserEnvironmentResultOutput() GetGlobalUserEnvironmentResultOutput {
	return o
}

func (o GetGlobalUserEnvironmentResultOutput) ToGetGlobalUserEnvironmentResultOutputWithContext(ctx context.Context) GetGlobalUserEnvironmentResultOutput {
	return o
}

func (o GetGlobalUserEnvironmentResultOutput) Environment() EnvironmentDetailsResponseOutput {
	return o.ApplyT(func(v GetGlobalUserEnvironmentResult) EnvironmentDetailsResponse { return v.Environment }).(EnvironmentDetailsResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(GetGlobalUserEnvironmentResultOutput{})
}
