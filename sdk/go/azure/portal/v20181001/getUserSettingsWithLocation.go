


package v20181001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupUserSettingsWithLocation(ctx *pulumi.Context, args *LookupUserSettingsWithLocationArgs, opts ...pulumi.InvokeOption) (*LookupUserSettingsWithLocationResult, error) {
	var rv LookupUserSettingsWithLocationResult
	err := ctx.Invoke("azure-native:portal/v20181001:getUserSettingsWithLocation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUserSettingsWithLocationArgs struct {
	Location         string `pulumi:"location"`
	UserSettingsName string `pulumi:"userSettingsName"`
}


type LookupUserSettingsWithLocationResult struct {
	Properties UserPropertiesResponse `pulumi:"properties"`
}

func LookupUserSettingsWithLocationOutput(ctx *pulumi.Context, args LookupUserSettingsWithLocationOutputArgs, opts ...pulumi.InvokeOption) LookupUserSettingsWithLocationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUserSettingsWithLocationResult, error) {
			args := v.(LookupUserSettingsWithLocationArgs)
			r, err := LookupUserSettingsWithLocation(ctx, &args, opts...)
			var s LookupUserSettingsWithLocationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUserSettingsWithLocationResultOutput)
}

type LookupUserSettingsWithLocationOutputArgs struct {
	Location         pulumi.StringInput `pulumi:"location"`
	UserSettingsName pulumi.StringInput `pulumi:"userSettingsName"`
}

func (LookupUserSettingsWithLocationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserSettingsWithLocationArgs)(nil)).Elem()
}


type LookupUserSettingsWithLocationResultOutput struct{ *pulumi.OutputState }

func (LookupUserSettingsWithLocationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserSettingsWithLocationResult)(nil)).Elem()
}

func (o LookupUserSettingsWithLocationResultOutput) ToLookupUserSettingsWithLocationResultOutput() LookupUserSettingsWithLocationResultOutput {
	return o
}

func (o LookupUserSettingsWithLocationResultOutput) ToLookupUserSettingsWithLocationResultOutputWithContext(ctx context.Context) LookupUserSettingsWithLocationResultOutput {
	return o
}

func (o LookupUserSettingsWithLocationResultOutput) Properties() UserPropertiesResponseOutput {
	return o.ApplyT(func(v LookupUserSettingsWithLocationResult) UserPropertiesResponse { return v.Properties }).(UserPropertiesResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserSettingsWithLocationResultOutput{})
}
