


package v20200401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConfigurationProfile(ctx *pulumi.Context, args *LookupConfigurationProfileArgs, opts ...pulumi.InvokeOption) (*LookupConfigurationProfileResult, error) {
	var rv LookupConfigurationProfileResult
	err := ctx.Invoke("azure-native:changeanalysis/v20200401preview:getConfigurationProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConfigurationProfileArgs struct {
	ProfileName string `pulumi:"profileName"`
}


type LookupConfigurationProfileResult struct {
	Id         string                                         `pulumi:"id"`
	Identity   *ResourceIdentityResponse                      `pulumi:"identity"`
	Location   *string                                        `pulumi:"location"`
	Name       string                                         `pulumi:"name"`
	Properties ConfigurationProfileResourcePropertiesResponse `pulumi:"properties"`
	SystemData *SystemDataResponse                            `pulumi:"systemData"`
	Type       string                                         `pulumi:"type"`
}

func LookupConfigurationProfileOutput(ctx *pulumi.Context, args LookupConfigurationProfileOutputArgs, opts ...pulumi.InvokeOption) LookupConfigurationProfileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConfigurationProfileResult, error) {
			args := v.(LookupConfigurationProfileArgs)
			r, err := LookupConfigurationProfile(ctx, &args, opts...)
			var s LookupConfigurationProfileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConfigurationProfileResultOutput)
}

type LookupConfigurationProfileOutputArgs struct {
	ProfileName pulumi.StringInput `pulumi:"profileName"`
}

func (LookupConfigurationProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigurationProfileArgs)(nil)).Elem()
}


type LookupConfigurationProfileResultOutput struct{ *pulumi.OutputState }

func (LookupConfigurationProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigurationProfileResult)(nil)).Elem()
}

func (o LookupConfigurationProfileResultOutput) ToLookupConfigurationProfileResultOutput() LookupConfigurationProfileResultOutput {
	return o
}

func (o LookupConfigurationProfileResultOutput) ToLookupConfigurationProfileResultOutputWithContext(ctx context.Context) LookupConfigurationProfileResultOutput {
	return o
}

func (o LookupConfigurationProfileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationProfileResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupConfigurationProfileResultOutput) Identity() ResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupConfigurationProfileResult) *ResourceIdentityResponse { return v.Identity }).(ResourceIdentityResponsePtrOutput)
}

func (o LookupConfigurationProfileResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConfigurationProfileResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupConfigurationProfileResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationProfileResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupConfigurationProfileResultOutput) Properties() ConfigurationProfileResourcePropertiesResponseOutput {
	return o.ApplyT(func(v LookupConfigurationProfileResult) ConfigurationProfileResourcePropertiesResponse {
		return v.Properties
	}).(ConfigurationProfileResourcePropertiesResponseOutput)
}

func (o LookupConfigurationProfileResultOutput) SystemData() SystemDataResponsePtrOutput {
	return o.ApplyT(func(v LookupConfigurationProfileResult) *SystemDataResponse { return v.SystemData }).(SystemDataResponsePtrOutput)
}

func (o LookupConfigurationProfileResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationProfileResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConfigurationProfileResultOutput{})
}
