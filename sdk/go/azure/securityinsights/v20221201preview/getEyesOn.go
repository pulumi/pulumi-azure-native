


package v20221201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEyesOn(ctx *pulumi.Context, args *LookupEyesOnArgs, opts ...pulumi.InvokeOption) (*LookupEyesOnResult, error) {
	var rv LookupEyesOnResult
	err := ctx.Invoke("azure-native:securityinsights/v20221201preview:getEyesOn", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEyesOnArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SettingsName      string `pulumi:"settingsName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupEyesOnResult struct {
	Etag       *string            `pulumi:"etag"`
	Id         string             `pulumi:"id"`
	IsEnabled  bool               `pulumi:"isEnabled"`
	Kind       string             `pulumi:"kind"`
	Name       string             `pulumi:"name"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Type       string             `pulumi:"type"`
}

func LookupEyesOnOutput(ctx *pulumi.Context, args LookupEyesOnOutputArgs, opts ...pulumi.InvokeOption) LookupEyesOnResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEyesOnResult, error) {
			args := v.(LookupEyesOnArgs)
			r, err := LookupEyesOn(ctx, &args, opts...)
			var s LookupEyesOnResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEyesOnResultOutput)
}

type LookupEyesOnOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	SettingsName      pulumi.StringInput `pulumi:"settingsName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupEyesOnOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEyesOnArgs)(nil)).Elem()
}


type LookupEyesOnResultOutput struct{ *pulumi.OutputState }

func (LookupEyesOnResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEyesOnResult)(nil)).Elem()
}

func (o LookupEyesOnResultOutput) ToLookupEyesOnResultOutput() LookupEyesOnResultOutput {
	return o
}

func (o LookupEyesOnResultOutput) ToLookupEyesOnResultOutputWithContext(ctx context.Context) LookupEyesOnResultOutput {
	return o
}

func (o LookupEyesOnResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEyesOnResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupEyesOnResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEyesOnResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupEyesOnResultOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupEyesOnResult) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o LookupEyesOnResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEyesOnResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupEyesOnResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEyesOnResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupEyesOnResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupEyesOnResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupEyesOnResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEyesOnResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEyesOnResultOutput{})
}
