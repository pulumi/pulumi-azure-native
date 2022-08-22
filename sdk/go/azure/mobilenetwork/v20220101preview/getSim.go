


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSim(ctx *pulumi.Context, args *LookupSimArgs, opts ...pulumi.InvokeOption) (*LookupSimResult, error) {
	var rv LookupSimResult
	err := ctx.Invoke("azure-native:mobilenetwork/v20220101preview:getSim", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSimArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SimName           string `pulumi:"simName"`
}


type LookupSimResult struct {
	ConfigurationState                    string                           `pulumi:"configurationState"`
	CreatedAt                             *string                          `pulumi:"createdAt"`
	CreatedBy                             *string                          `pulumi:"createdBy"`
	CreatedByType                         *string                          `pulumi:"createdByType"`
	DeviceType                            *string                          `pulumi:"deviceType"`
	Id                                    string                           `pulumi:"id"`
	IntegratedCircuitCardIdentifier       *string                          `pulumi:"integratedCircuitCardIdentifier"`
	InternationalMobileSubscriberIdentity string                           `pulumi:"internationalMobileSubscriberIdentity"`
	LastModifiedAt                        *string                          `pulumi:"lastModifiedAt"`
	LastModifiedBy                        *string                          `pulumi:"lastModifiedBy"`
	LastModifiedByType                    *string                          `pulumi:"lastModifiedByType"`
	Location                              string                           `pulumi:"location"`
	MobileNetwork                         *MobileNetworkResourceIdResponse `pulumi:"mobileNetwork"`
	Name                                  string                           `pulumi:"name"`
	ProvisioningState                     string                           `pulumi:"provisioningState"`
	SimPolicy                             *SimPolicyResourceIdResponse     `pulumi:"simPolicy"`
	StaticIpConfiguration                 []SimStaticIpPropertiesResponse  `pulumi:"staticIpConfiguration"`
	Tags                                  map[string]string                `pulumi:"tags"`
	Type                                  string                           `pulumi:"type"`
}

func LookupSimOutput(ctx *pulumi.Context, args LookupSimOutputArgs, opts ...pulumi.InvokeOption) LookupSimResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSimResult, error) {
			args := v.(LookupSimArgs)
			r, err := LookupSim(ctx, &args, opts...)
			var s LookupSimResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSimResultOutput)
}

type LookupSimOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	SimName           pulumi.StringInput `pulumi:"simName"`
}

func (LookupSimOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSimArgs)(nil)).Elem()
}


type LookupSimResultOutput struct{ *pulumi.OutputState }

func (LookupSimResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSimResult)(nil)).Elem()
}

func (o LookupSimResultOutput) ToLookupSimResultOutput() LookupSimResultOutput {
	return o
}

func (o LookupSimResultOutput) ToLookupSimResultOutputWithContext(ctx context.Context) LookupSimResultOutput {
	return o
}

func (o LookupSimResultOutput) ConfigurationState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSimResult) string { return v.ConfigurationState }).(pulumi.StringOutput)
}

func (o LookupSimResultOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSimResult) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o LookupSimResultOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSimResult) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o LookupSimResultOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSimResult) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o LookupSimResultOutput) DeviceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSimResult) *string { return v.DeviceType }).(pulumi.StringPtrOutput)
}

func (o LookupSimResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSimResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSimResultOutput) IntegratedCircuitCardIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSimResult) *string { return v.IntegratedCircuitCardIdentifier }).(pulumi.StringPtrOutput)
}

func (o LookupSimResultOutput) InternationalMobileSubscriberIdentity() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSimResult) string { return v.InternationalMobileSubscriberIdentity }).(pulumi.StringOutput)
}

func (o LookupSimResultOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSimResult) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o LookupSimResultOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSimResult) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o LookupSimResultOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSimResult) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func (o LookupSimResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSimResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSimResultOutput) MobileNetwork() MobileNetworkResourceIdResponsePtrOutput {
	return o.ApplyT(func(v LookupSimResult) *MobileNetworkResourceIdResponse { return v.MobileNetwork }).(MobileNetworkResourceIdResponsePtrOutput)
}

func (o LookupSimResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSimResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSimResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSimResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupSimResultOutput) SimPolicy() SimPolicyResourceIdResponsePtrOutput {
	return o.ApplyT(func(v LookupSimResult) *SimPolicyResourceIdResponse { return v.SimPolicy }).(SimPolicyResourceIdResponsePtrOutput)
}

func (o LookupSimResultOutput) StaticIpConfiguration() SimStaticIpPropertiesResponseArrayOutput {
	return o.ApplyT(func(v LookupSimResult) []SimStaticIpPropertiesResponse { return v.StaticIpConfiguration }).(SimStaticIpPropertiesResponseArrayOutput)
}

func (o LookupSimResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSimResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSimResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSimResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSimResultOutput{})
}
