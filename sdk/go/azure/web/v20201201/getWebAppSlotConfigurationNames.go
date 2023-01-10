


package v20201201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppSlotConfigurationNames(ctx *pulumi.Context, args *LookupWebAppSlotConfigurationNamesArgs, opts ...pulumi.InvokeOption) (*LookupWebAppSlotConfigurationNamesResult, error) {
	var rv LookupWebAppSlotConfigurationNamesResult
	err := ctx.Invoke("azure-native:web/v20201201:getWebAppSlotConfigurationNames", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppSlotConfigurationNamesArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupWebAppSlotConfigurationNamesResult struct {
	AppSettingNames         []string `pulumi:"appSettingNames"`
	AzureStorageConfigNames []string `pulumi:"azureStorageConfigNames"`
	ConnectionStringNames   []string `pulumi:"connectionStringNames"`
	Id                      string   `pulumi:"id"`
	Kind                    *string  `pulumi:"kind"`
	Name                    string   `pulumi:"name"`
	Type                    string   `pulumi:"type"`
}

func LookupWebAppSlotConfigurationNamesOutput(ctx *pulumi.Context, args LookupWebAppSlotConfigurationNamesOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppSlotConfigurationNamesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppSlotConfigurationNamesResult, error) {
			args := v.(LookupWebAppSlotConfigurationNamesArgs)
			r, err := LookupWebAppSlotConfigurationNames(ctx, &args, opts...)
			var s LookupWebAppSlotConfigurationNamesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppSlotConfigurationNamesResultOutput)
}

type LookupWebAppSlotConfigurationNamesOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupWebAppSlotConfigurationNamesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppSlotConfigurationNamesArgs)(nil)).Elem()
}


type LookupWebAppSlotConfigurationNamesResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppSlotConfigurationNamesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppSlotConfigurationNamesResult)(nil)).Elem()
}

func (o LookupWebAppSlotConfigurationNamesResultOutput) ToLookupWebAppSlotConfigurationNamesResultOutput() LookupWebAppSlotConfigurationNamesResultOutput {
	return o
}

func (o LookupWebAppSlotConfigurationNamesResultOutput) ToLookupWebAppSlotConfigurationNamesResultOutputWithContext(ctx context.Context) LookupWebAppSlotConfigurationNamesResultOutput {
	return o
}

func (o LookupWebAppSlotConfigurationNamesResultOutput) AppSettingNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWebAppSlotConfigurationNamesResult) []string { return v.AppSettingNames }).(pulumi.StringArrayOutput)
}

func (o LookupWebAppSlotConfigurationNamesResultOutput) AzureStorageConfigNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWebAppSlotConfigurationNamesResult) []string { return v.AzureStorageConfigNames }).(pulumi.StringArrayOutput)
}

func (o LookupWebAppSlotConfigurationNamesResultOutput) ConnectionStringNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWebAppSlotConfigurationNamesResult) []string { return v.ConnectionStringNames }).(pulumi.StringArrayOutput)
}

func (o LookupWebAppSlotConfigurationNamesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSlotConfigurationNamesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWebAppSlotConfigurationNamesResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSlotConfigurationNamesResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSlotConfigurationNamesResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSlotConfigurationNamesResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWebAppSlotConfigurationNamesResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSlotConfigurationNamesResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppSlotConfigurationNamesResultOutput{})
}
