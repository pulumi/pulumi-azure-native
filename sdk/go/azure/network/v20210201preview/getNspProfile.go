


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNspProfile(ctx *pulumi.Context, args *LookupNspProfileArgs, opts ...pulumi.InvokeOption) (*LookupNspProfileResult, error) {
	var rv LookupNspProfileResult
	err := ctx.Invoke("azure-native:network/v20210201preview:getNspProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNspProfileArgs struct {
	NetworkSecurityPerimeterName string `pulumi:"networkSecurityPerimeterName"`
	ProfileName                  string `pulumi:"profileName"`
	ResourceGroupName            string `pulumi:"resourceGroupName"`
}


type LookupNspProfileResult struct {
	AccessRulesVersion        string            `pulumi:"accessRulesVersion"`
	DiagnosticSettingsVersion string            `pulumi:"diagnosticSettingsVersion"`
	Id                        string            `pulumi:"id"`
	Location                  *string           `pulumi:"location"`
	Name                      string            `pulumi:"name"`
	Tags                      map[string]string `pulumi:"tags"`
	Type                      string            `pulumi:"type"`
}

func LookupNspProfileOutput(ctx *pulumi.Context, args LookupNspProfileOutputArgs, opts ...pulumi.InvokeOption) LookupNspProfileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNspProfileResult, error) {
			args := v.(LookupNspProfileArgs)
			r, err := LookupNspProfile(ctx, &args, opts...)
			var s LookupNspProfileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNspProfileResultOutput)
}

type LookupNspProfileOutputArgs struct {
	NetworkSecurityPerimeterName pulumi.StringInput `pulumi:"networkSecurityPerimeterName"`
	ProfileName                  pulumi.StringInput `pulumi:"profileName"`
	ResourceGroupName            pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNspProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNspProfileArgs)(nil)).Elem()
}


type LookupNspProfileResultOutput struct{ *pulumi.OutputState }

func (LookupNspProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNspProfileResult)(nil)).Elem()
}

func (o LookupNspProfileResultOutput) ToLookupNspProfileResultOutput() LookupNspProfileResultOutput {
	return o
}

func (o LookupNspProfileResultOutput) ToLookupNspProfileResultOutputWithContext(ctx context.Context) LookupNspProfileResultOutput {
	return o
}

func (o LookupNspProfileResultOutput) AccessRulesVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNspProfileResult) string { return v.AccessRulesVersion }).(pulumi.StringOutput)
}

func (o LookupNspProfileResultOutput) DiagnosticSettingsVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNspProfileResult) string { return v.DiagnosticSettingsVersion }).(pulumi.StringOutput)
}

func (o LookupNspProfileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNspProfileResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNspProfileResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNspProfileResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupNspProfileResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNspProfileResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNspProfileResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNspProfileResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupNspProfileResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNspProfileResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNspProfileResultOutput{})
}
