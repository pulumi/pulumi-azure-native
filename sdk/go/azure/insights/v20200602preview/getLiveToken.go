


package v20200602preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetLiveToken(ctx *pulumi.Context, args *GetLiveTokenArgs, opts ...pulumi.InvokeOption) (*GetLiveTokenResult, error) {
	var rv GetLiveTokenResult
	err := ctx.Invoke("azure-native:insights/v20200602preview:getLiveToken", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetLiveTokenArgs struct {
	ResourceUri string `pulumi:"resourceUri"`
}


type GetLiveTokenResult struct {
	LiveToken string `pulumi:"liveToken"`
}

func GetLiveTokenOutput(ctx *pulumi.Context, args GetLiveTokenOutputArgs, opts ...pulumi.InvokeOption) GetLiveTokenResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetLiveTokenResult, error) {
			args := v.(GetLiveTokenArgs)
			r, err := GetLiveToken(ctx, &args, opts...)
			var s GetLiveTokenResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetLiveTokenResultOutput)
}

type GetLiveTokenOutputArgs struct {
	ResourceUri pulumi.StringInput `pulumi:"resourceUri"`
}

func (GetLiveTokenOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLiveTokenArgs)(nil)).Elem()
}


type GetLiveTokenResultOutput struct{ *pulumi.OutputState }

func (GetLiveTokenResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLiveTokenResult)(nil)).Elem()
}

func (o GetLiveTokenResultOutput) ToGetLiveTokenResultOutput() GetLiveTokenResultOutput {
	return o
}

func (o GetLiveTokenResultOutput) ToGetLiveTokenResultOutputWithContext(ctx context.Context) GetLiveTokenResultOutput {
	return o
}

func (o GetLiveTokenResultOutput) LiveToken() pulumi.StringOutput {
	return o.ApplyT(func(v GetLiveTokenResult) string { return v.LiveToken }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLiveTokenResultOutput{})
}
