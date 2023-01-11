


package v20161001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupChapSetting(ctx *pulumi.Context, args *LookupChapSettingArgs, opts ...pulumi.InvokeOption) (*LookupChapSettingResult, error) {
	var rv LookupChapSettingResult
	err := ctx.Invoke("azure-native:storsimple/v20161001:getChapSetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupChapSettingArgs struct {
	ChapUserName      string `pulumi:"chapUserName"`
	DeviceName        string `pulumi:"deviceName"`
	ManagerName       string `pulumi:"managerName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupChapSettingResult struct {
	Id       string                            `pulumi:"id"`
	Name     string                            `pulumi:"name"`
	Password AsymmetricEncryptedSecretResponse `pulumi:"password"`
	Type     string                            `pulumi:"type"`
}

func LookupChapSettingOutput(ctx *pulumi.Context, args LookupChapSettingOutputArgs, opts ...pulumi.InvokeOption) LookupChapSettingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupChapSettingResult, error) {
			args := v.(LookupChapSettingArgs)
			r, err := LookupChapSetting(ctx, &args, opts...)
			var s LookupChapSettingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupChapSettingResultOutput)
}

type LookupChapSettingOutputArgs struct {
	ChapUserName      pulumi.StringInput `pulumi:"chapUserName"`
	DeviceName        pulumi.StringInput `pulumi:"deviceName"`
	ManagerName       pulumi.StringInput `pulumi:"managerName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupChapSettingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupChapSettingArgs)(nil)).Elem()
}


type LookupChapSettingResultOutput struct{ *pulumi.OutputState }

func (LookupChapSettingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupChapSettingResult)(nil)).Elem()
}

func (o LookupChapSettingResultOutput) ToLookupChapSettingResultOutput() LookupChapSettingResultOutput {
	return o
}

func (o LookupChapSettingResultOutput) ToLookupChapSettingResultOutputWithContext(ctx context.Context) LookupChapSettingResultOutput {
	return o
}

func (o LookupChapSettingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupChapSettingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupChapSettingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupChapSettingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupChapSettingResultOutput) Password() AsymmetricEncryptedSecretResponseOutput {
	return o.ApplyT(func(v LookupChapSettingResult) AsymmetricEncryptedSecretResponse { return v.Password }).(AsymmetricEncryptedSecretResponseOutput)
}

func (o LookupChapSettingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupChapSettingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupChapSettingResultOutput{})
}
