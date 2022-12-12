


package v20220504

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConfigurationProfile(ctx *pulumi.Context, args *LookupConfigurationProfileArgs, opts ...pulumi.InvokeOption) (*LookupConfigurationProfileResult, error) {
	var rv LookupConfigurationProfileResult
	err := ctx.Invoke("azure-native:automanage/v20220504:getConfigurationProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConfigurationProfileArgs struct {
	ConfigurationProfileName string `pulumi:"configurationProfileName"`
	ResourceGroupName        string `pulumi:"resourceGroupName"`
}


type LookupConfigurationProfileResult struct {
	Id         string                                 `pulumi:"id"`
	Location   string                                 `pulumi:"location"`
	Name       string                                 `pulumi:"name"`
	Properties ConfigurationProfilePropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse                     `pulumi:"systemData"`
	Tags       map[string]string                      `pulumi:"tags"`
	Type       string                                 `pulumi:"type"`
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
	ConfigurationProfileName pulumi.StringInput `pulumi:"configurationProfileName"`
	ResourceGroupName        pulumi.StringInput `pulumi:"resourceGroupName"`
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

func (o LookupConfigurationProfileResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationProfileResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupConfigurationProfileResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationProfileResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupConfigurationProfileResultOutput) Properties() ConfigurationProfilePropertiesResponseOutput {
	return o.ApplyT(func(v LookupConfigurationProfileResult) ConfigurationProfilePropertiesResponse { return v.Properties }).(ConfigurationProfilePropertiesResponseOutput)
}

func (o LookupConfigurationProfileResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupConfigurationProfileResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupConfigurationProfileResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupConfigurationProfileResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupConfigurationProfileResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationProfileResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConfigurationProfileResultOutput{})
}
