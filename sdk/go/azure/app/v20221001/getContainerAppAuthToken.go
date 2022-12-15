


package v20221001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetContainerAppAuthToken(ctx *pulumi.Context, args *GetContainerAppAuthTokenArgs, opts ...pulumi.InvokeOption) (*GetContainerAppAuthTokenResult, error) {
	var rv GetContainerAppAuthTokenResult
	err := ctx.Invoke("azure-native:app/v20221001:getContainerAppAuthToken", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetContainerAppAuthTokenArgs struct {
	ContainerAppName  string `pulumi:"containerAppName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type GetContainerAppAuthTokenResult struct {
	Expires    string             `pulumi:"expires"`
	Id         string             `pulumi:"id"`
	Location   string             `pulumi:"location"`
	Name       string             `pulumi:"name"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Tags       map[string]string  `pulumi:"tags"`
	Token      string             `pulumi:"token"`
	Type       string             `pulumi:"type"`
}

func GetContainerAppAuthTokenOutput(ctx *pulumi.Context, args GetContainerAppAuthTokenOutputArgs, opts ...pulumi.InvokeOption) GetContainerAppAuthTokenResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetContainerAppAuthTokenResult, error) {
			args := v.(GetContainerAppAuthTokenArgs)
			r, err := GetContainerAppAuthToken(ctx, &args, opts...)
			var s GetContainerAppAuthTokenResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetContainerAppAuthTokenResultOutput)
}

type GetContainerAppAuthTokenOutputArgs struct {
	ContainerAppName  pulumi.StringInput `pulumi:"containerAppName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetContainerAppAuthTokenOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetContainerAppAuthTokenArgs)(nil)).Elem()
}


type GetContainerAppAuthTokenResultOutput struct{ *pulumi.OutputState }

func (GetContainerAppAuthTokenResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetContainerAppAuthTokenResult)(nil)).Elem()
}

func (o GetContainerAppAuthTokenResultOutput) ToGetContainerAppAuthTokenResultOutput() GetContainerAppAuthTokenResultOutput {
	return o
}

func (o GetContainerAppAuthTokenResultOutput) ToGetContainerAppAuthTokenResultOutputWithContext(ctx context.Context) GetContainerAppAuthTokenResultOutput {
	return o
}

func (o GetContainerAppAuthTokenResultOutput) Expires() pulumi.StringOutput {
	return o.ApplyT(func(v GetContainerAppAuthTokenResult) string { return v.Expires }).(pulumi.StringOutput)
}

func (o GetContainerAppAuthTokenResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetContainerAppAuthTokenResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetContainerAppAuthTokenResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetContainerAppAuthTokenResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o GetContainerAppAuthTokenResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetContainerAppAuthTokenResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetContainerAppAuthTokenResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v GetContainerAppAuthTokenResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o GetContainerAppAuthTokenResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetContainerAppAuthTokenResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GetContainerAppAuthTokenResultOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v GetContainerAppAuthTokenResult) string { return v.Token }).(pulumi.StringOutput)
}

func (o GetContainerAppAuthTokenResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetContainerAppAuthTokenResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetContainerAppAuthTokenResultOutput{})
}
