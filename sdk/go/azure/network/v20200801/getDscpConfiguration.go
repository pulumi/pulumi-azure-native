


package v20200801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDscpConfiguration(ctx *pulumi.Context, args *LookupDscpConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupDscpConfigurationResult, error) {
	var rv LookupDscpConfigurationResult
	err := ctx.Invoke("azure-native:network/v20200801:getDscpConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDscpConfigurationArgs struct {
	DscpConfigurationName string `pulumi:"dscpConfigurationName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupDscpConfigurationResult struct {
	AssociatedNetworkInterfaces []NetworkInterfaceResponse `pulumi:"associatedNetworkInterfaces"`
	DestinationIpRanges         []QosIpRangeResponse       `pulumi:"destinationIpRanges"`
	DestinationPortRanges       []QosPortRangeResponse     `pulumi:"destinationPortRanges"`
	Etag                        string                     `pulumi:"etag"`
	Id                          *string                    `pulumi:"id"`
	Location                    *string                    `pulumi:"location"`
	Markings                    []int                      `pulumi:"markings"`
	Name                        string                     `pulumi:"name"`
	Protocol                    *string                    `pulumi:"protocol"`
	ProvisioningState           string                     `pulumi:"provisioningState"`
	QosCollectionId             string                     `pulumi:"qosCollectionId"`
	ResourceGuid                string                     `pulumi:"resourceGuid"`
	SourceIpRanges              []QosIpRangeResponse       `pulumi:"sourceIpRanges"`
	SourcePortRanges            []QosPortRangeResponse     `pulumi:"sourcePortRanges"`
	Tags                        map[string]string          `pulumi:"tags"`
	Type                        string                     `pulumi:"type"`
}

func LookupDscpConfigurationOutput(ctx *pulumi.Context, args LookupDscpConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupDscpConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDscpConfigurationResult, error) {
			args := v.(LookupDscpConfigurationArgs)
			r, err := LookupDscpConfiguration(ctx, &args, opts...)
			var s LookupDscpConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDscpConfigurationResultOutput)
}

type LookupDscpConfigurationOutputArgs struct {
	DscpConfigurationName pulumi.StringInput `pulumi:"dscpConfigurationName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDscpConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDscpConfigurationArgs)(nil)).Elem()
}


type LookupDscpConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupDscpConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDscpConfigurationResult)(nil)).Elem()
}

func (o LookupDscpConfigurationResultOutput) ToLookupDscpConfigurationResultOutput() LookupDscpConfigurationResultOutput {
	return o
}

func (o LookupDscpConfigurationResultOutput) ToLookupDscpConfigurationResultOutputWithContext(ctx context.Context) LookupDscpConfigurationResultOutput {
	return o
}

func (o LookupDscpConfigurationResultOutput) AssociatedNetworkInterfaces() NetworkInterfaceResponseArrayOutput {
	return o.ApplyT(func(v LookupDscpConfigurationResult) []NetworkInterfaceResponse { return v.AssociatedNetworkInterfaces }).(NetworkInterfaceResponseArrayOutput)
}

func (o LookupDscpConfigurationResultOutput) DestinationIpRanges() QosIpRangeResponseArrayOutput {
	return o.ApplyT(func(v LookupDscpConfigurationResult) []QosIpRangeResponse { return v.DestinationIpRanges }).(QosIpRangeResponseArrayOutput)
}

func (o LookupDscpConfigurationResultOutput) DestinationPortRanges() QosPortRangeResponseArrayOutput {
	return o.ApplyT(func(v LookupDscpConfigurationResult) []QosPortRangeResponse { return v.DestinationPortRanges }).(QosPortRangeResponseArrayOutput)
}

func (o LookupDscpConfigurationResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDscpConfigurationResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupDscpConfigurationResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDscpConfigurationResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupDscpConfigurationResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDscpConfigurationResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupDscpConfigurationResultOutput) Markings() pulumi.IntArrayOutput {
	return o.ApplyT(func(v LookupDscpConfigurationResult) []int { return v.Markings }).(pulumi.IntArrayOutput)
}

func (o LookupDscpConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDscpConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDscpConfigurationResultOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDscpConfigurationResult) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o LookupDscpConfigurationResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDscpConfigurationResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupDscpConfigurationResultOutput) QosCollectionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDscpConfigurationResult) string { return v.QosCollectionId }).(pulumi.StringOutput)
}

func (o LookupDscpConfigurationResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDscpConfigurationResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o LookupDscpConfigurationResultOutput) SourceIpRanges() QosIpRangeResponseArrayOutput {
	return o.ApplyT(func(v LookupDscpConfigurationResult) []QosIpRangeResponse { return v.SourceIpRanges }).(QosIpRangeResponseArrayOutput)
}

func (o LookupDscpConfigurationResultOutput) SourcePortRanges() QosPortRangeResponseArrayOutput {
	return o.ApplyT(func(v LookupDscpConfigurationResult) []QosPortRangeResponse { return v.SourcePortRanges }).(QosPortRangeResponseArrayOutput)
}

func (o LookupDscpConfigurationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDscpConfigurationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDscpConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDscpConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDscpConfigurationResultOutput{})
}
