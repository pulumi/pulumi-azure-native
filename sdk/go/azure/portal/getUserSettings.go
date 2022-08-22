


package portal

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupUserSettings(ctx *pulumi.Context, args *LookupUserSettingsArgs, opts ...pulumi.InvokeOption) (*LookupUserSettingsResult, error) {
	var rv LookupUserSettingsResult
	err := ctx.Invoke("azure-native:portal:getUserSettings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUserSettingsArgs struct {
	UserSettingsName string `pulumi:"userSettingsName"`
}


type LookupUserSettingsResult struct {
	Properties UserPropertiesResponse `pulumi:"properties"`
}

func LookupUserSettingsOutput(ctx *pulumi.Context, args LookupUserSettingsOutputArgs, opts ...pulumi.InvokeOption) LookupUserSettingsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUserSettingsResult, error) {
			args := v.(LookupUserSettingsArgs)
			r, err := LookupUserSettings(ctx, &args, opts...)
			var s LookupUserSettingsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUserSettingsResultOutput)
}

type LookupUserSettingsOutputArgs struct {
	UserSettingsName pulumi.StringInput `pulumi:"userSettingsName"`
}

func (LookupUserSettingsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserSettingsArgs)(nil)).Elem()
}


type LookupUserSettingsResultOutput struct{ *pulumi.OutputState }

func (LookupUserSettingsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserSettingsResult)(nil)).Elem()
}

func (o LookupUserSettingsResultOutput) ToLookupUserSettingsResultOutput() LookupUserSettingsResultOutput {
	return o
}

func (o LookupUserSettingsResultOutput) ToLookupUserSettingsResultOutputWithContext(ctx context.Context) LookupUserSettingsResultOutput {
	return o
}

func (o LookupUserSettingsResultOutput) Properties() UserPropertiesResponseOutput {
	return o.ApplyT(func(v LookupUserSettingsResult) UserPropertiesResponse { return v.Properties }).(UserPropertiesResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserSettingsResultOutput{})
}
