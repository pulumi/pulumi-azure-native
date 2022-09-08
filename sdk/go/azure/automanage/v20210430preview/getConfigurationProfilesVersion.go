


package v20210430preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConfigurationProfilesVersion(ctx *pulumi.Context, args *LookupConfigurationProfilesVersionArgs, opts ...pulumi.InvokeOption) (*LookupConfigurationProfilesVersionResult, error) {
	var rv LookupConfigurationProfilesVersionResult
	err := ctx.Invoke("azure-native:automanage/v20210430preview:getConfigurationProfilesVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConfigurationProfilesVersionArgs struct {
	ConfigurationProfileName string `pulumi:"configurationProfileName"`
	ResourceGroupName        string `pulumi:"resourceGroupName"`
	VersionName              string `pulumi:"versionName"`
}


type LookupConfigurationProfilesVersionResult struct {
	Id         string                                 `pulumi:"id"`
	Location   string                                 `pulumi:"location"`
	Name       string                                 `pulumi:"name"`
	Properties ConfigurationProfilePropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse                     `pulumi:"systemData"`
	Tags       map[string]string                      `pulumi:"tags"`
	Type       string                                 `pulumi:"type"`
}

func LookupConfigurationProfilesVersionOutput(ctx *pulumi.Context, args LookupConfigurationProfilesVersionOutputArgs, opts ...pulumi.InvokeOption) LookupConfigurationProfilesVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConfigurationProfilesVersionResult, error) {
			args := v.(LookupConfigurationProfilesVersionArgs)
			r, err := LookupConfigurationProfilesVersion(ctx, &args, opts...)
			var s LookupConfigurationProfilesVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConfigurationProfilesVersionResultOutput)
}

type LookupConfigurationProfilesVersionOutputArgs struct {
	ConfigurationProfileName pulumi.StringInput `pulumi:"configurationProfileName"`
	ResourceGroupName        pulumi.StringInput `pulumi:"resourceGroupName"`
	VersionName              pulumi.StringInput `pulumi:"versionName"`
}

func (LookupConfigurationProfilesVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigurationProfilesVersionArgs)(nil)).Elem()
}


type LookupConfigurationProfilesVersionResultOutput struct{ *pulumi.OutputState }

func (LookupConfigurationProfilesVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigurationProfilesVersionResult)(nil)).Elem()
}

func (o LookupConfigurationProfilesVersionResultOutput) ToLookupConfigurationProfilesVersionResultOutput() LookupConfigurationProfilesVersionResultOutput {
	return o
}

func (o LookupConfigurationProfilesVersionResultOutput) ToLookupConfigurationProfilesVersionResultOutputWithContext(ctx context.Context) LookupConfigurationProfilesVersionResultOutput {
	return o
}

func (o LookupConfigurationProfilesVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationProfilesVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupConfigurationProfilesVersionResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationProfilesVersionResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupConfigurationProfilesVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationProfilesVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupConfigurationProfilesVersionResultOutput) Properties() ConfigurationProfilePropertiesResponseOutput {
	return o.ApplyT(func(v LookupConfigurationProfilesVersionResult) ConfigurationProfilePropertiesResponse {
		return v.Properties
	}).(ConfigurationProfilePropertiesResponseOutput)
}

func (o LookupConfigurationProfilesVersionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupConfigurationProfilesVersionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupConfigurationProfilesVersionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupConfigurationProfilesVersionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupConfigurationProfilesVersionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationProfilesVersionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConfigurationProfilesVersionResultOutput{})
}
