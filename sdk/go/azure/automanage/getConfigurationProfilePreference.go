


package automanage

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConfigurationProfilePreference(ctx *pulumi.Context, args *LookupConfigurationProfilePreferenceArgs, opts ...pulumi.InvokeOption) (*LookupConfigurationProfilePreferenceResult, error) {
	var rv LookupConfigurationProfilePreferenceResult
	err := ctx.Invoke("azure-native:automanage:getConfigurationProfilePreference", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConfigurationProfilePreferenceArgs struct {
	ConfigurationProfilePreferenceName string `pulumi:"configurationProfilePreferenceName"`
	ResourceGroupName                  string `pulumi:"resourceGroupName"`
}


type LookupConfigurationProfilePreferenceResult struct {
	Id         string                                           `pulumi:"id"`
	Location   string                                           `pulumi:"location"`
	Name       string                                           `pulumi:"name"`
	Properties ConfigurationProfilePreferencePropertiesResponse `pulumi:"properties"`
	Tags       map[string]string                                `pulumi:"tags"`
	Type       string                                           `pulumi:"type"`
}

func LookupConfigurationProfilePreferenceOutput(ctx *pulumi.Context, args LookupConfigurationProfilePreferenceOutputArgs, opts ...pulumi.InvokeOption) LookupConfigurationProfilePreferenceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConfigurationProfilePreferenceResult, error) {
			args := v.(LookupConfigurationProfilePreferenceArgs)
			r, err := LookupConfigurationProfilePreference(ctx, &args, opts...)
			var s LookupConfigurationProfilePreferenceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConfigurationProfilePreferenceResultOutput)
}

type LookupConfigurationProfilePreferenceOutputArgs struct {
	ConfigurationProfilePreferenceName pulumi.StringInput `pulumi:"configurationProfilePreferenceName"`
	ResourceGroupName                  pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupConfigurationProfilePreferenceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigurationProfilePreferenceArgs)(nil)).Elem()
}


type LookupConfigurationProfilePreferenceResultOutput struct{ *pulumi.OutputState }

func (LookupConfigurationProfilePreferenceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigurationProfilePreferenceResult)(nil)).Elem()
}

func (o LookupConfigurationProfilePreferenceResultOutput) ToLookupConfigurationProfilePreferenceResultOutput() LookupConfigurationProfilePreferenceResultOutput {
	return o
}

func (o LookupConfigurationProfilePreferenceResultOutput) ToLookupConfigurationProfilePreferenceResultOutputWithContext(ctx context.Context) LookupConfigurationProfilePreferenceResultOutput {
	return o
}

func (o LookupConfigurationProfilePreferenceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationProfilePreferenceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupConfigurationProfilePreferenceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationProfilePreferenceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupConfigurationProfilePreferenceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationProfilePreferenceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupConfigurationProfilePreferenceResultOutput) Properties() ConfigurationProfilePreferencePropertiesResponseOutput {
	return o.ApplyT(func(v LookupConfigurationProfilePreferenceResult) ConfigurationProfilePreferencePropertiesResponse {
		return v.Properties
	}).(ConfigurationProfilePreferencePropertiesResponseOutput)
}

func (o LookupConfigurationProfilePreferenceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupConfigurationProfilePreferenceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupConfigurationProfilePreferenceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationProfilePreferenceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConfigurationProfilePreferenceResultOutput{})
}
