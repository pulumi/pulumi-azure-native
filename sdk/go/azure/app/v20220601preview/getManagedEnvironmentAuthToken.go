


package v20220601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetManagedEnvironmentAuthToken(ctx *pulumi.Context, args *GetManagedEnvironmentAuthTokenArgs, opts ...pulumi.InvokeOption) (*GetManagedEnvironmentAuthTokenResult, error) {
	var rv GetManagedEnvironmentAuthTokenResult
	err := ctx.Invoke("azure-native:app/v20220601preview:getManagedEnvironmentAuthToken", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetManagedEnvironmentAuthTokenArgs struct {
	EnvironmentName   string `pulumi:"environmentName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type GetManagedEnvironmentAuthTokenResult struct {
	Expires    string             `pulumi:"expires"`
	Id         string             `pulumi:"id"`
	Location   string             `pulumi:"location"`
	Name       string             `pulumi:"name"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Tags       map[string]string  `pulumi:"tags"`
	Token      string             `pulumi:"token"`
	Type       string             `pulumi:"type"`
}

func GetManagedEnvironmentAuthTokenOutput(ctx *pulumi.Context, args GetManagedEnvironmentAuthTokenOutputArgs, opts ...pulumi.InvokeOption) GetManagedEnvironmentAuthTokenResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetManagedEnvironmentAuthTokenResult, error) {
			args := v.(GetManagedEnvironmentAuthTokenArgs)
			r, err := GetManagedEnvironmentAuthToken(ctx, &args, opts...)
			var s GetManagedEnvironmentAuthTokenResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetManagedEnvironmentAuthTokenResultOutput)
}

type GetManagedEnvironmentAuthTokenOutputArgs struct {
	EnvironmentName   pulumi.StringInput `pulumi:"environmentName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetManagedEnvironmentAuthTokenOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetManagedEnvironmentAuthTokenArgs)(nil)).Elem()
}


type GetManagedEnvironmentAuthTokenResultOutput struct{ *pulumi.OutputState }

func (GetManagedEnvironmentAuthTokenResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetManagedEnvironmentAuthTokenResult)(nil)).Elem()
}

func (o GetManagedEnvironmentAuthTokenResultOutput) ToGetManagedEnvironmentAuthTokenResultOutput() GetManagedEnvironmentAuthTokenResultOutput {
	return o
}

func (o GetManagedEnvironmentAuthTokenResultOutput) ToGetManagedEnvironmentAuthTokenResultOutputWithContext(ctx context.Context) GetManagedEnvironmentAuthTokenResultOutput {
	return o
}

func (o GetManagedEnvironmentAuthTokenResultOutput) Expires() pulumi.StringOutput {
	return o.ApplyT(func(v GetManagedEnvironmentAuthTokenResult) string { return v.Expires }).(pulumi.StringOutput)
}

func (o GetManagedEnvironmentAuthTokenResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetManagedEnvironmentAuthTokenResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetManagedEnvironmentAuthTokenResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetManagedEnvironmentAuthTokenResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o GetManagedEnvironmentAuthTokenResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetManagedEnvironmentAuthTokenResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetManagedEnvironmentAuthTokenResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v GetManagedEnvironmentAuthTokenResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o GetManagedEnvironmentAuthTokenResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetManagedEnvironmentAuthTokenResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GetManagedEnvironmentAuthTokenResultOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v GetManagedEnvironmentAuthTokenResult) string { return v.Token }).(pulumi.StringOutput)
}

func (o GetManagedEnvironmentAuthTokenResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetManagedEnvironmentAuthTokenResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetManagedEnvironmentAuthTokenResultOutput{})
}
