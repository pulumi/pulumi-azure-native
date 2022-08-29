


package v20180601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetFactoryDataPlaneAccess(ctx *pulumi.Context, args *GetFactoryDataPlaneAccessArgs, opts ...pulumi.InvokeOption) (*GetFactoryDataPlaneAccessResult, error) {
	var rv GetFactoryDataPlaneAccessResult
	err := ctx.Invoke("azure-native:datafactory/v20180601:getFactoryDataPlaneAccess", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetFactoryDataPlaneAccessArgs struct {
	AccessResourcePath *string `pulumi:"accessResourcePath"`
	ExpireTime         *string `pulumi:"expireTime"`
	FactoryName        string  `pulumi:"factoryName"`
	Permissions        *string `pulumi:"permissions"`
	ProfileName        *string `pulumi:"profileName"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
	StartTime          *string `pulumi:"startTime"`
}


type GetFactoryDataPlaneAccessResult struct {
	AccessToken  *string                   `pulumi:"accessToken"`
	DataPlaneUrl *string                   `pulumi:"dataPlaneUrl"`
	Policy       *UserAccessPolicyResponse `pulumi:"policy"`
}

func GetFactoryDataPlaneAccessOutput(ctx *pulumi.Context, args GetFactoryDataPlaneAccessOutputArgs, opts ...pulumi.InvokeOption) GetFactoryDataPlaneAccessResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFactoryDataPlaneAccessResult, error) {
			args := v.(GetFactoryDataPlaneAccessArgs)
			r, err := GetFactoryDataPlaneAccess(ctx, &args, opts...)
			var s GetFactoryDataPlaneAccessResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetFactoryDataPlaneAccessResultOutput)
}

type GetFactoryDataPlaneAccessOutputArgs struct {
	AccessResourcePath pulumi.StringPtrInput `pulumi:"accessResourcePath"`
	ExpireTime         pulumi.StringPtrInput `pulumi:"expireTime"`
	FactoryName        pulumi.StringInput    `pulumi:"factoryName"`
	Permissions        pulumi.StringPtrInput `pulumi:"permissions"`
	ProfileName        pulumi.StringPtrInput `pulumi:"profileName"`
	ResourceGroupName  pulumi.StringInput    `pulumi:"resourceGroupName"`
	StartTime          pulumi.StringPtrInput `pulumi:"startTime"`
}

func (GetFactoryDataPlaneAccessOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFactoryDataPlaneAccessArgs)(nil)).Elem()
}


type GetFactoryDataPlaneAccessResultOutput struct{ *pulumi.OutputState }

func (GetFactoryDataPlaneAccessResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFactoryDataPlaneAccessResult)(nil)).Elem()
}

func (o GetFactoryDataPlaneAccessResultOutput) ToGetFactoryDataPlaneAccessResultOutput() GetFactoryDataPlaneAccessResultOutput {
	return o
}

func (o GetFactoryDataPlaneAccessResultOutput) ToGetFactoryDataPlaneAccessResultOutputWithContext(ctx context.Context) GetFactoryDataPlaneAccessResultOutput {
	return o
}

func (o GetFactoryDataPlaneAccessResultOutput) AccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFactoryDataPlaneAccessResult) *string { return v.AccessToken }).(pulumi.StringPtrOutput)
}

func (o GetFactoryDataPlaneAccessResultOutput) DataPlaneUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFactoryDataPlaneAccessResult) *string { return v.DataPlaneUrl }).(pulumi.StringPtrOutput)
}

func (o GetFactoryDataPlaneAccessResultOutput) Policy() UserAccessPolicyResponsePtrOutput {
	return o.ApplyT(func(v GetFactoryDataPlaneAccessResult) *UserAccessPolicyResponse { return v.Policy }).(UserAccessPolicyResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFactoryDataPlaneAccessResultOutput{})
}
