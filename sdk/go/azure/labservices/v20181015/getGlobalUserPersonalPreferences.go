


package v20181015

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetGlobalUserPersonalPreferences(ctx *pulumi.Context, args *GetGlobalUserPersonalPreferencesArgs, opts ...pulumi.InvokeOption) (*GetGlobalUserPersonalPreferencesResult, error) {
	var rv GetGlobalUserPersonalPreferencesResult
	err := ctx.Invoke("azure-native:labservices/v20181015:getGlobalUserPersonalPreferences", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetGlobalUserPersonalPreferencesArgs struct {
	AddRemove            *string `pulumi:"addRemove"`
	LabAccountResourceId *string `pulumi:"labAccountResourceId"`
	LabResourceId        *string `pulumi:"labResourceId"`
	UserName             string  `pulumi:"userName"`
}


type GetGlobalUserPersonalPreferencesResult struct {
	FavoriteLabResourceIds []string `pulumi:"favoriteLabResourceIds"`
	Id                     *string  `pulumi:"id"`
}

func GetGlobalUserPersonalPreferencesOutput(ctx *pulumi.Context, args GetGlobalUserPersonalPreferencesOutputArgs, opts ...pulumi.InvokeOption) GetGlobalUserPersonalPreferencesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetGlobalUserPersonalPreferencesResult, error) {
			args := v.(GetGlobalUserPersonalPreferencesArgs)
			r, err := GetGlobalUserPersonalPreferences(ctx, &args, opts...)
			var s GetGlobalUserPersonalPreferencesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetGlobalUserPersonalPreferencesResultOutput)
}

type GetGlobalUserPersonalPreferencesOutputArgs struct {
	AddRemove            pulumi.StringPtrInput `pulumi:"addRemove"`
	LabAccountResourceId pulumi.StringPtrInput `pulumi:"labAccountResourceId"`
	LabResourceId        pulumi.StringPtrInput `pulumi:"labResourceId"`
	UserName             pulumi.StringInput    `pulumi:"userName"`
}

func (GetGlobalUserPersonalPreferencesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGlobalUserPersonalPreferencesArgs)(nil)).Elem()
}


type GetGlobalUserPersonalPreferencesResultOutput struct{ *pulumi.OutputState }

func (GetGlobalUserPersonalPreferencesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGlobalUserPersonalPreferencesResult)(nil)).Elem()
}

func (o GetGlobalUserPersonalPreferencesResultOutput) ToGetGlobalUserPersonalPreferencesResultOutput() GetGlobalUserPersonalPreferencesResultOutput {
	return o
}

func (o GetGlobalUserPersonalPreferencesResultOutput) ToGetGlobalUserPersonalPreferencesResultOutputWithContext(ctx context.Context) GetGlobalUserPersonalPreferencesResultOutput {
	return o
}

func (o GetGlobalUserPersonalPreferencesResultOutput) FavoriteLabResourceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetGlobalUserPersonalPreferencesResult) []string { return v.FavoriteLabResourceIds }).(pulumi.StringArrayOutput)
}

func (o GetGlobalUserPersonalPreferencesResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGlobalUserPersonalPreferencesResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetGlobalUserPersonalPreferencesResultOutput{})
}
