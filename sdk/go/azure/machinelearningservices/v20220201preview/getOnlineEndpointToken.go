


package v20220201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetOnlineEndpointToken(ctx *pulumi.Context, args *GetOnlineEndpointTokenArgs, opts ...pulumi.InvokeOption) (*GetOnlineEndpointTokenResult, error) {
	var rv GetOnlineEndpointTokenResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20220201preview:getOnlineEndpointToken", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type GetOnlineEndpointTokenArgs struct {
	EndpointName      string `pulumi:"endpointName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type GetOnlineEndpointTokenResult struct {
	AccessToken         *string  `pulumi:"accessToken"`
	ExpiryTimeUtc       *float64 `pulumi:"expiryTimeUtc"`
	RefreshAfterTimeUtc *float64 `pulumi:"refreshAfterTimeUtc"`
	TokenType           *string  `pulumi:"tokenType"`
}


func (val *GetOnlineEndpointTokenResult) Defaults() *GetOnlineEndpointTokenResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ExpiryTimeUtc) {
		expiryTimeUtc_ := 0.0
		tmp.ExpiryTimeUtc = &expiryTimeUtc_
	}
	if isZero(tmp.RefreshAfterTimeUtc) {
		refreshAfterTimeUtc_ := 0.0
		tmp.RefreshAfterTimeUtc = &refreshAfterTimeUtc_
	}
	return &tmp
}

func GetOnlineEndpointTokenOutput(ctx *pulumi.Context, args GetOnlineEndpointTokenOutputArgs, opts ...pulumi.InvokeOption) GetOnlineEndpointTokenResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetOnlineEndpointTokenResult, error) {
			args := v.(GetOnlineEndpointTokenArgs)
			r, err := GetOnlineEndpointToken(ctx, &args, opts...)
			var s GetOnlineEndpointTokenResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetOnlineEndpointTokenResultOutput)
}

type GetOnlineEndpointTokenOutputArgs struct {
	EndpointName      pulumi.StringInput `pulumi:"endpointName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (GetOnlineEndpointTokenOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOnlineEndpointTokenArgs)(nil)).Elem()
}


type GetOnlineEndpointTokenResultOutput struct{ *pulumi.OutputState }

func (GetOnlineEndpointTokenResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOnlineEndpointTokenResult)(nil)).Elem()
}

func (o GetOnlineEndpointTokenResultOutput) ToGetOnlineEndpointTokenResultOutput() GetOnlineEndpointTokenResultOutput {
	return o
}

func (o GetOnlineEndpointTokenResultOutput) ToGetOnlineEndpointTokenResultOutputWithContext(ctx context.Context) GetOnlineEndpointTokenResultOutput {
	return o
}

func (o GetOnlineEndpointTokenResultOutput) AccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOnlineEndpointTokenResult) *string { return v.AccessToken }).(pulumi.StringPtrOutput)
}

func (o GetOnlineEndpointTokenResultOutput) ExpiryTimeUtc() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v GetOnlineEndpointTokenResult) *float64 { return v.ExpiryTimeUtc }).(pulumi.Float64PtrOutput)
}

func (o GetOnlineEndpointTokenResultOutput) RefreshAfterTimeUtc() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v GetOnlineEndpointTokenResult) *float64 { return v.RefreshAfterTimeUtc }).(pulumi.Float64PtrOutput)
}

func (o GetOnlineEndpointTokenResultOutput) TokenType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOnlineEndpointTokenResult) *string { return v.TokenType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOnlineEndpointTokenResultOutput{})
}
