


package v20220615preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDataController(ctx *pulumi.Context, args *LookupDataControllerArgs, opts ...pulumi.InvokeOption) (*LookupDataControllerResult, error) {
	var rv LookupDataControllerResult
	err := ctx.Invoke("azure-native:azurearcdata/v20220615preview:getDataController", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupDataControllerArgs struct {
	DataControllerName string `pulumi:"dataControllerName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupDataControllerResult struct {
	ExtendedLocation *ExtendedLocationResponse        `pulumi:"extendedLocation"`
	Id               string                           `pulumi:"id"`
	Location         string                           `pulumi:"location"`
	Name             string                           `pulumi:"name"`
	Properties       DataControllerPropertiesResponse `pulumi:"properties"`
	SystemData       SystemDataResponse               `pulumi:"systemData"`
	Tags             map[string]string                `pulumi:"tags"`
	Type             string                           `pulumi:"type"`
}


func (val *LookupDataControllerResult) Defaults() *LookupDataControllerResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}

func LookupDataControllerOutput(ctx *pulumi.Context, args LookupDataControllerOutputArgs, opts ...pulumi.InvokeOption) LookupDataControllerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDataControllerResult, error) {
			args := v.(LookupDataControllerArgs)
			r, err := LookupDataController(ctx, &args, opts...)
			var s LookupDataControllerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDataControllerResultOutput)
}

type LookupDataControllerOutputArgs struct {
	DataControllerName pulumi.StringInput `pulumi:"dataControllerName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDataControllerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataControllerArgs)(nil)).Elem()
}


type LookupDataControllerResultOutput struct{ *pulumi.OutputState }

func (LookupDataControllerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataControllerResult)(nil)).Elem()
}

func (o LookupDataControllerResultOutput) ToLookupDataControllerResultOutput() LookupDataControllerResultOutput {
	return o
}

func (o LookupDataControllerResultOutput) ToLookupDataControllerResultOutputWithContext(ctx context.Context) LookupDataControllerResultOutput {
	return o
}

func (o LookupDataControllerResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupDataControllerResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o LookupDataControllerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataControllerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDataControllerResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataControllerResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupDataControllerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataControllerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDataControllerResultOutput) Properties() DataControllerPropertiesResponseOutput {
	return o.ApplyT(func(v LookupDataControllerResult) DataControllerPropertiesResponse { return v.Properties }).(DataControllerPropertiesResponseOutput)
}

func (o LookupDataControllerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDataControllerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupDataControllerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDataControllerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDataControllerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataControllerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataControllerResultOutput{})
}
