


package v20211031preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAppliance(ctx *pulumi.Context, args *LookupApplianceArgs, opts ...pulumi.InvokeOption) (*LookupApplianceResult, error) {
	var rv LookupApplianceResult
	err := ctx.Invoke("azure-native:resourceconnector/v20211031preview:getAppliance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupApplianceArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupApplianceResult struct {
	Distro               *string                                          `pulumi:"distro"`
	Id                   string                                           `pulumi:"id"`
	Identity             *IdentityResponse                                `pulumi:"identity"`
	InfrastructureConfig *AppliancePropertiesResponseInfrastructureConfig `pulumi:"infrastructureConfig"`
	Location             string                                           `pulumi:"location"`
	Name                 string                                           `pulumi:"name"`
	ProvisioningState    string                                           `pulumi:"provisioningState"`
	PublicKey            *string                                          `pulumi:"publicKey"`
	Status               string                                           `pulumi:"status"`
	SystemData           SystemDataResponse                               `pulumi:"systemData"`
	Tags                 map[string]string                                `pulumi:"tags"`
	Type                 string                                           `pulumi:"type"`
	Version              string                                           `pulumi:"version"`
}


func (val *LookupApplianceResult) Defaults() *LookupApplianceResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Distro) {
		distro_ := "AKSEdge"
		tmp.Distro = &distro_
	}
	return &tmp
}

func LookupApplianceOutput(ctx *pulumi.Context, args LookupApplianceOutputArgs, opts ...pulumi.InvokeOption) LookupApplianceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApplianceResult, error) {
			args := v.(LookupApplianceArgs)
			r, err := LookupAppliance(ctx, &args, opts...)
			var s LookupApplianceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApplianceResultOutput)
}

type LookupApplianceOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupApplianceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplianceArgs)(nil)).Elem()
}


type LookupApplianceResultOutput struct{ *pulumi.OutputState }

func (LookupApplianceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplianceResult)(nil)).Elem()
}

func (o LookupApplianceResultOutput) ToLookupApplianceResultOutput() LookupApplianceResultOutput {
	return o
}

func (o LookupApplianceResultOutput) ToLookupApplianceResultOutputWithContext(ctx context.Context) LookupApplianceResultOutput {
	return o
}

func (o LookupApplianceResultOutput) Distro() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplianceResult) *string { return v.Distro }).(pulumi.StringPtrOutput)
}

func (o LookupApplianceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplianceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApplianceResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupApplianceResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o LookupApplianceResultOutput) InfrastructureConfig() AppliancePropertiesResponseInfrastructureConfigPtrOutput {
	return o.ApplyT(func(v LookupApplianceResult) *AppliancePropertiesResponseInfrastructureConfig {
		return v.InfrastructureConfig
	}).(AppliancePropertiesResponseInfrastructureConfigPtrOutput)
}

func (o LookupApplianceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplianceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupApplianceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplianceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupApplianceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplianceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupApplianceResultOutput) PublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplianceResult) *string { return v.PublicKey }).(pulumi.StringPtrOutput)
}

func (o LookupApplianceResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplianceResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupApplianceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupApplianceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupApplianceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupApplianceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupApplianceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplianceResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupApplianceResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplianceResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplianceResultOutput{})
}
