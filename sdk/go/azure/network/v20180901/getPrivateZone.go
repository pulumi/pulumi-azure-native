


package v20180901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateZone(ctx *pulumi.Context, args *LookupPrivateZoneArgs, opts ...pulumi.InvokeOption) (*LookupPrivateZoneResult, error) {
	var rv LookupPrivateZoneResult
	err := ctx.Invoke("azure-native:network/v20180901:getPrivateZone", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateZoneArgs struct {
	PrivateZoneName   string `pulumi:"privateZoneName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupPrivateZoneResult struct {
	Etag                                           *string           `pulumi:"etag"`
	Id                                             string            `pulumi:"id"`
	Location                                       *string           `pulumi:"location"`
	MaxNumberOfRecordSets                          float64           `pulumi:"maxNumberOfRecordSets"`
	MaxNumberOfVirtualNetworkLinks                 float64           `pulumi:"maxNumberOfVirtualNetworkLinks"`
	MaxNumberOfVirtualNetworkLinksWithRegistration float64           `pulumi:"maxNumberOfVirtualNetworkLinksWithRegistration"`
	Name                                           string            `pulumi:"name"`
	NumberOfRecordSets                             float64           `pulumi:"numberOfRecordSets"`
	NumberOfVirtualNetworkLinks                    float64           `pulumi:"numberOfVirtualNetworkLinks"`
	NumberOfVirtualNetworkLinksWithRegistration    float64           `pulumi:"numberOfVirtualNetworkLinksWithRegistration"`
	ProvisioningState                              string            `pulumi:"provisioningState"`
	Tags                                           map[string]string `pulumi:"tags"`
	Type                                           string            `pulumi:"type"`
}

func LookupPrivateZoneOutput(ctx *pulumi.Context, args LookupPrivateZoneOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateZoneResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateZoneResult, error) {
			args := v.(LookupPrivateZoneArgs)
			r, err := LookupPrivateZone(ctx, &args, opts...)
			var s LookupPrivateZoneResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateZoneResultOutput)
}

type LookupPrivateZoneOutputArgs struct {
	PrivateZoneName   pulumi.StringInput `pulumi:"privateZoneName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPrivateZoneOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateZoneArgs)(nil)).Elem()
}


type LookupPrivateZoneResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateZoneResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateZoneResult)(nil)).Elem()
}

func (o LookupPrivateZoneResultOutput) ToLookupPrivateZoneResultOutput() LookupPrivateZoneResultOutput {
	return o
}

func (o LookupPrivateZoneResultOutput) ToLookupPrivateZoneResultOutputWithContext(ctx context.Context) LookupPrivateZoneResultOutput {
	return o
}

func (o LookupPrivateZoneResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateZoneResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupPrivateZoneResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateZoneResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPrivateZoneResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateZoneResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupPrivateZoneResultOutput) MaxNumberOfRecordSets() pulumi.Float64Output {
	return o.ApplyT(func(v LookupPrivateZoneResult) float64 { return v.MaxNumberOfRecordSets }).(pulumi.Float64Output)
}

func (o LookupPrivateZoneResultOutput) MaxNumberOfVirtualNetworkLinks() pulumi.Float64Output {
	return o.ApplyT(func(v LookupPrivateZoneResult) float64 { return v.MaxNumberOfVirtualNetworkLinks }).(pulumi.Float64Output)
}

func (o LookupPrivateZoneResultOutput) MaxNumberOfVirtualNetworkLinksWithRegistration() pulumi.Float64Output {
	return o.ApplyT(func(v LookupPrivateZoneResult) float64 { return v.MaxNumberOfVirtualNetworkLinksWithRegistration }).(pulumi.Float64Output)
}

func (o LookupPrivateZoneResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateZoneResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPrivateZoneResultOutput) NumberOfRecordSets() pulumi.Float64Output {
	return o.ApplyT(func(v LookupPrivateZoneResult) float64 { return v.NumberOfRecordSets }).(pulumi.Float64Output)
}

func (o LookupPrivateZoneResultOutput) NumberOfVirtualNetworkLinks() pulumi.Float64Output {
	return o.ApplyT(func(v LookupPrivateZoneResult) float64 { return v.NumberOfVirtualNetworkLinks }).(pulumi.Float64Output)
}

func (o LookupPrivateZoneResultOutput) NumberOfVirtualNetworkLinksWithRegistration() pulumi.Float64Output {
	return o.ApplyT(func(v LookupPrivateZoneResult) float64 { return v.NumberOfVirtualNetworkLinksWithRegistration }).(pulumi.Float64Output)
}

func (o LookupPrivateZoneResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateZoneResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupPrivateZoneResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPrivateZoneResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupPrivateZoneResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateZoneResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateZoneResultOutput{})
}
