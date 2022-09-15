


package v20220404preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEnergyService(ctx *pulumi.Context, args *LookupEnergyServiceArgs, opts ...pulumi.InvokeOption) (*LookupEnergyServiceResult, error) {
	var rv LookupEnergyServiceResult
	err := ctx.Invoke("azure-native:openenergyplatform/v20220404preview:getEnergyService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEnergyServiceArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}

type LookupEnergyServiceResult struct {
	Id         string                          `pulumi:"id"`
	Location   string                          `pulumi:"location"`
	Name       string                          `pulumi:"name"`
	Properties EnergyServicePropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse              `pulumi:"systemData"`
	Tags       map[string]string               `pulumi:"tags"`
	Type       string                          `pulumi:"type"`
}

func LookupEnergyServiceOutput(ctx *pulumi.Context, args LookupEnergyServiceOutputArgs, opts ...pulumi.InvokeOption) LookupEnergyServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEnergyServiceResult, error) {
			args := v.(LookupEnergyServiceArgs)
			r, err := LookupEnergyService(ctx, &args, opts...)
			var s LookupEnergyServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEnergyServiceResultOutput)
}

type LookupEnergyServiceOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupEnergyServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnergyServiceArgs)(nil)).Elem()
}

type LookupEnergyServiceResultOutput struct{ *pulumi.OutputState }

func (LookupEnergyServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnergyServiceResult)(nil)).Elem()
}

func (o LookupEnergyServiceResultOutput) ToLookupEnergyServiceResultOutput() LookupEnergyServiceResultOutput {
	return o
}

func (o LookupEnergyServiceResultOutput) ToLookupEnergyServiceResultOutputWithContext(ctx context.Context) LookupEnergyServiceResultOutput {
	return o
}

func (o LookupEnergyServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnergyServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupEnergyServiceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnergyServiceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupEnergyServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnergyServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupEnergyServiceResultOutput) Properties() EnergyServicePropertiesResponseOutput {
	return o.ApplyT(func(v LookupEnergyServiceResult) EnergyServicePropertiesResponse { return v.Properties }).(EnergyServicePropertiesResponseOutput)
}

func (o LookupEnergyServiceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupEnergyServiceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupEnergyServiceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupEnergyServiceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupEnergyServiceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnergyServiceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEnergyServiceResultOutput{})
}
