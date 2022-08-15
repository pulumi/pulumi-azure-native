


package v20190101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetUserSharedAccessToken(ctx *pulumi.Context, args *GetUserSharedAccessTokenArgs, opts ...pulumi.InvokeOption) (*GetUserSharedAccessTokenResult, error) {
	var rv GetUserSharedAccessTokenResult
	err := ctx.Invoke("azure-native:apimanagement/v20190101:getUserSharedAccessToken", args.Defaults(), &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetUserSharedAccessTokenArgs struct {
	Expiry            string  `pulumi:"expiry"`
	KeyType           KeyType `pulumi:"keyType"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServiceName       string  `pulumi:"serviceName"`
	UserId            string  `pulumi:"userId"`
}


func (val *GetUserSharedAccessTokenArgs) Defaults() *GetUserSharedAccessTokenArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.KeyType) {
		tmp.KeyType = KeyType("primary")
	}
	return &tmp
}


type GetUserSharedAccessTokenResult struct {
	Value *string `pulumi:"value"`
}

func GetUserSharedAccessTokenOutput(ctx *pulumi.Context, args GetUserSharedAccessTokenOutputArgs, opts ...pulumi.InvokeOption) GetUserSharedAccessTokenResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetUserSharedAccessTokenResult, error) {
			args := v.(GetUserSharedAccessTokenArgs)
			r, err := GetUserSharedAccessToken(ctx, &args, opts...)
			var s GetUserSharedAccessTokenResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetUserSharedAccessTokenResultOutput)
}

type GetUserSharedAccessTokenOutputArgs struct {
	Expiry            pulumi.StringInput `pulumi:"expiry"`
	KeyType           KeyTypeInput       `pulumi:"keyType"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
	UserId            pulumi.StringInput `pulumi:"userId"`
}

func (GetUserSharedAccessTokenOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUserSharedAccessTokenArgs)(nil)).Elem()
}


type GetUserSharedAccessTokenResultOutput struct{ *pulumi.OutputState }

func (GetUserSharedAccessTokenResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUserSharedAccessTokenResult)(nil)).Elem()
}

func (o GetUserSharedAccessTokenResultOutput) ToGetUserSharedAccessTokenResultOutput() GetUserSharedAccessTokenResultOutput {
	return o
}

func (o GetUserSharedAccessTokenResultOutput) ToGetUserSharedAccessTokenResultOutputWithContext(ctx context.Context) GetUserSharedAccessTokenResultOutput {
	return o
}

func (o GetUserSharedAccessTokenResultOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUserSharedAccessTokenResult) *string { return v.Value }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetUserSharedAccessTokenResultOutput{})
}
