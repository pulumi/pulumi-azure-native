


package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupSiteSlotConfigNames(ctx *pulumi.Context, args *LookupSiteSlotConfigNamesArgs, opts ...pulumi.InvokeOption) (*LookupSiteSlotConfigNamesResult, error) {
	var rv LookupSiteSlotConfigNamesResult
	err := ctx.Invoke("azure-native:web/v20150801:getSiteSlotConfigNames", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSiteSlotConfigNamesArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupSiteSlotConfigNamesResult struct {
	AppSettingNames       []string          `pulumi:"appSettingNames"`
	ConnectionStringNames []string          `pulumi:"connectionStringNames"`
	Id                    *string           `pulumi:"id"`
	Kind                  *string           `pulumi:"kind"`
	Location              string            `pulumi:"location"`
	Name                  *string           `pulumi:"name"`
	Tags                  map[string]string `pulumi:"tags"`
	Type                  *string           `pulumi:"type"`
}

func LookupSiteSlotConfigNamesOutput(ctx *pulumi.Context, args LookupSiteSlotConfigNamesOutputArgs, opts ...pulumi.InvokeOption) LookupSiteSlotConfigNamesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSiteSlotConfigNamesResult, error) {
			args := v.(LookupSiteSlotConfigNamesArgs)
			r, err := LookupSiteSlotConfigNames(ctx, &args, opts...)
			var s LookupSiteSlotConfigNamesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSiteSlotConfigNamesResultOutput)
}

type LookupSiteSlotConfigNamesOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupSiteSlotConfigNamesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSiteSlotConfigNamesArgs)(nil)).Elem()
}


type LookupSiteSlotConfigNamesResultOutput struct{ *pulumi.OutputState }

func (LookupSiteSlotConfigNamesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSiteSlotConfigNamesResult)(nil)).Elem()
}

func (o LookupSiteSlotConfigNamesResultOutput) ToLookupSiteSlotConfigNamesResultOutput() LookupSiteSlotConfigNamesResultOutput {
	return o
}

func (o LookupSiteSlotConfigNamesResultOutput) ToLookupSiteSlotConfigNamesResultOutputWithContext(ctx context.Context) LookupSiteSlotConfigNamesResultOutput {
	return o
}

func (o LookupSiteSlotConfigNamesResultOutput) AppSettingNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSiteSlotConfigNamesResult) []string { return v.AppSettingNames }).(pulumi.StringArrayOutput)
}

func (o LookupSiteSlotConfigNamesResultOutput) ConnectionStringNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSiteSlotConfigNamesResult) []string { return v.ConnectionStringNames }).(pulumi.StringArrayOutput)
}

func (o LookupSiteSlotConfigNamesResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteSlotConfigNamesResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupSiteSlotConfigNamesResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteSlotConfigNamesResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupSiteSlotConfigNamesResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteSlotConfigNamesResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSiteSlotConfigNamesResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteSlotConfigNamesResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupSiteSlotConfigNamesResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSiteSlotConfigNamesResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSiteSlotConfigNamesResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteSlotConfigNamesResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSiteSlotConfigNamesResultOutput{})
}
