


package logz

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSubAccount(ctx *pulumi.Context, args *LookupSubAccountArgs, opts ...pulumi.InvokeOption) (*LookupSubAccountResult, error) {
	var rv LookupSubAccountResult
	err := ctx.Invoke("azure-native:logz:getSubAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSubAccountArgs struct {
	MonitorName       string `pulumi:"monitorName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SubAccountName    string `pulumi:"subAccountName"`
}

type LookupSubAccountResult struct {
	Id         string                      `pulumi:"id"`
	Identity   *IdentityPropertiesResponse `pulumi:"identity"`
	Location   string                      `pulumi:"location"`
	Name       string                      `pulumi:"name"`
	Properties MonitorPropertiesResponse   `pulumi:"properties"`
	SystemData SystemDataResponse          `pulumi:"systemData"`
	Tags       map[string]string           `pulumi:"tags"`
	Type       string                      `pulumi:"type"`
}

func LookupSubAccountOutput(ctx *pulumi.Context, args LookupSubAccountOutputArgs, opts ...pulumi.InvokeOption) LookupSubAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSubAccountResult, error) {
			args := v.(LookupSubAccountArgs)
			r, err := LookupSubAccount(ctx, &args, opts...)
			var s LookupSubAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSubAccountResultOutput)
}

type LookupSubAccountOutputArgs struct {
	MonitorName       pulumi.StringInput `pulumi:"monitorName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	SubAccountName    pulumi.StringInput `pulumi:"subAccountName"`
}

func (LookupSubAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubAccountArgs)(nil)).Elem()
}

type LookupSubAccountResultOutput struct{ *pulumi.OutputState }

func (LookupSubAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubAccountResult)(nil)).Elem()
}

func (o LookupSubAccountResultOutput) ToLookupSubAccountResultOutput() LookupSubAccountResultOutput {
	return o
}

func (o LookupSubAccountResultOutput) ToLookupSubAccountResultOutputWithContext(ctx context.Context) LookupSubAccountResultOutput {
	return o
}

func (o LookupSubAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSubAccountResultOutput) Identity() IdentityPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupSubAccountResult) *IdentityPropertiesResponse { return v.Identity }).(IdentityPropertiesResponsePtrOutput)
}

func (o LookupSubAccountResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubAccountResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSubAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSubAccountResultOutput) Properties() MonitorPropertiesResponseOutput {
	return o.ApplyT(func(v LookupSubAccountResult) MonitorPropertiesResponse { return v.Properties }).(MonitorPropertiesResponseOutput)
}

func (o LookupSubAccountResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSubAccountResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSubAccountResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSubAccountResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSubAccountResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubAccountResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSubAccountResultOutput{})
}
