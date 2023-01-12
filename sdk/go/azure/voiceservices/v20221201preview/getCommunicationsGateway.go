


package v20221201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCommunicationsGateway(ctx *pulumi.Context, args *LookupCommunicationsGatewayArgs, opts ...pulumi.InvokeOption) (*LookupCommunicationsGatewayResult, error) {
	var rv LookupCommunicationsGatewayResult
	err := ctx.Invoke("azure-native:voiceservices/v20221201preview:getCommunicationsGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCommunicationsGatewayArgs struct {
	CommunicationsGatewayName string `pulumi:"communicationsGatewayName"`
	ResourceGroupName         string `pulumi:"resourceGroupName"`
}


type LookupCommunicationsGatewayResult struct {
	ApiBridge         interface{}                       `pulumi:"apiBridge"`
	Codecs            []string                          `pulumi:"codecs"`
	Connectivity      string                            `pulumi:"connectivity"`
	E911Type          string                            `pulumi:"e911Type"`
	Id                string                            `pulumi:"id"`
	Location          string                            `pulumi:"location"`
	Name              string                            `pulumi:"name"`
	Platforms         []string                          `pulumi:"platforms"`
	ProvisioningState string                            `pulumi:"provisioningState"`
	ServiceLocations  []ServiceRegionPropertiesResponse `pulumi:"serviceLocations"`
	Status            string                            `pulumi:"status"`
	SystemData        SystemDataResponse                `pulumi:"systemData"`
	Tags              map[string]string                 `pulumi:"tags"`
	Type              string                            `pulumi:"type"`
}

func LookupCommunicationsGatewayOutput(ctx *pulumi.Context, args LookupCommunicationsGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupCommunicationsGatewayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCommunicationsGatewayResult, error) {
			args := v.(LookupCommunicationsGatewayArgs)
			r, err := LookupCommunicationsGateway(ctx, &args, opts...)
			var s LookupCommunicationsGatewayResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCommunicationsGatewayResultOutput)
}

type LookupCommunicationsGatewayOutputArgs struct {
	CommunicationsGatewayName pulumi.StringInput `pulumi:"communicationsGatewayName"`
	ResourceGroupName         pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCommunicationsGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCommunicationsGatewayArgs)(nil)).Elem()
}


type LookupCommunicationsGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupCommunicationsGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCommunicationsGatewayResult)(nil)).Elem()
}

func (o LookupCommunicationsGatewayResultOutput) ToLookupCommunicationsGatewayResultOutput() LookupCommunicationsGatewayResultOutput {
	return o
}

func (o LookupCommunicationsGatewayResultOutput) ToLookupCommunicationsGatewayResultOutputWithContext(ctx context.Context) LookupCommunicationsGatewayResultOutput {
	return o
}

func (o LookupCommunicationsGatewayResultOutput) ApiBridge() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupCommunicationsGatewayResult) interface{} { return v.ApiBridge }).(pulumi.AnyOutput)
}

func (o LookupCommunicationsGatewayResultOutput) Codecs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupCommunicationsGatewayResult) []string { return v.Codecs }).(pulumi.StringArrayOutput)
}

func (o LookupCommunicationsGatewayResultOutput) Connectivity() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommunicationsGatewayResult) string { return v.Connectivity }).(pulumi.StringOutput)
}

func (o LookupCommunicationsGatewayResultOutput) E911Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommunicationsGatewayResult) string { return v.E911Type }).(pulumi.StringOutput)
}

func (o LookupCommunicationsGatewayResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommunicationsGatewayResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCommunicationsGatewayResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommunicationsGatewayResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupCommunicationsGatewayResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommunicationsGatewayResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCommunicationsGatewayResultOutput) Platforms() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupCommunicationsGatewayResult) []string { return v.Platforms }).(pulumi.StringArrayOutput)
}

func (o LookupCommunicationsGatewayResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommunicationsGatewayResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupCommunicationsGatewayResultOutput) ServiceLocations() ServiceRegionPropertiesResponseArrayOutput {
	return o.ApplyT(func(v LookupCommunicationsGatewayResult) []ServiceRegionPropertiesResponse { return v.ServiceLocations }).(ServiceRegionPropertiesResponseArrayOutput)
}

func (o LookupCommunicationsGatewayResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommunicationsGatewayResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupCommunicationsGatewayResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupCommunicationsGatewayResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupCommunicationsGatewayResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCommunicationsGatewayResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupCommunicationsGatewayResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommunicationsGatewayResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCommunicationsGatewayResultOutput{})
}
