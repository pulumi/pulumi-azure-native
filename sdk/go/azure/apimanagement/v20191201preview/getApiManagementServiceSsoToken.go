


package v20191201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetApiManagementServiceSsoToken(ctx *pulumi.Context, args *GetApiManagementServiceSsoTokenArgs, opts ...pulumi.InvokeOption) (*GetApiManagementServiceSsoTokenResult, error) {
	var rv GetApiManagementServiceSsoTokenResult
	err := ctx.Invoke("azure-native:apimanagement/v20191201preview:getApiManagementServiceSsoToken", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetApiManagementServiceSsoTokenArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type GetApiManagementServiceSsoTokenResult struct {
	RedirectUri *string `pulumi:"redirectUri"`
}

func GetApiManagementServiceSsoTokenOutput(ctx *pulumi.Context, args GetApiManagementServiceSsoTokenOutputArgs, opts ...pulumi.InvokeOption) GetApiManagementServiceSsoTokenResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetApiManagementServiceSsoTokenResult, error) {
			args := v.(GetApiManagementServiceSsoTokenArgs)
			r, err := GetApiManagementServiceSsoToken(ctx, &args, opts...)
			var s GetApiManagementServiceSsoTokenResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetApiManagementServiceSsoTokenResultOutput)
}

type GetApiManagementServiceSsoTokenOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (GetApiManagementServiceSsoTokenOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApiManagementServiceSsoTokenArgs)(nil)).Elem()
}


type GetApiManagementServiceSsoTokenResultOutput struct{ *pulumi.OutputState }

func (GetApiManagementServiceSsoTokenResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApiManagementServiceSsoTokenResult)(nil)).Elem()
}

func (o GetApiManagementServiceSsoTokenResultOutput) ToGetApiManagementServiceSsoTokenResultOutput() GetApiManagementServiceSsoTokenResultOutput {
	return o
}

func (o GetApiManagementServiceSsoTokenResultOutput) ToGetApiManagementServiceSsoTokenResultOutputWithContext(ctx context.Context) GetApiManagementServiceSsoTokenResultOutput {
	return o
}

func (o GetApiManagementServiceSsoTokenResultOutput) RedirectUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApiManagementServiceSsoTokenResult) *string { return v.RedirectUri }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetApiManagementServiceSsoTokenResultOutput{})
}
