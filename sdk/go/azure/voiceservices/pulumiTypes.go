


package voiceservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrimaryRegionProperties struct {
	EsrpAddresses     []string `pulumi:"esrpAddresses"`
	OperatorAddresses []string `pulumi:"operatorAddresses"`
}





type PrimaryRegionPropertiesInput interface {
	pulumi.Input

	ToPrimaryRegionPropertiesOutput() PrimaryRegionPropertiesOutput
	ToPrimaryRegionPropertiesOutputWithContext(context.Context) PrimaryRegionPropertiesOutput
}

type PrimaryRegionPropertiesArgs struct {
	EsrpAddresses     pulumi.StringArrayInput `pulumi:"esrpAddresses"`
	OperatorAddresses pulumi.StringArrayInput `pulumi:"operatorAddresses"`
}

func (PrimaryRegionPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrimaryRegionProperties)(nil)).Elem()
}

func (i PrimaryRegionPropertiesArgs) ToPrimaryRegionPropertiesOutput() PrimaryRegionPropertiesOutput {
	return i.ToPrimaryRegionPropertiesOutputWithContext(context.Background())
}

func (i PrimaryRegionPropertiesArgs) ToPrimaryRegionPropertiesOutputWithContext(ctx context.Context) PrimaryRegionPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrimaryRegionPropertiesOutput)
}

type PrimaryRegionPropertiesOutput struct{ *pulumi.OutputState }

func (PrimaryRegionPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrimaryRegionProperties)(nil)).Elem()
}

func (o PrimaryRegionPropertiesOutput) ToPrimaryRegionPropertiesOutput() PrimaryRegionPropertiesOutput {
	return o
}

func (o PrimaryRegionPropertiesOutput) ToPrimaryRegionPropertiesOutputWithContext(ctx context.Context) PrimaryRegionPropertiesOutput {
	return o
}

func (o PrimaryRegionPropertiesOutput) EsrpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrimaryRegionProperties) []string { return v.EsrpAddresses }).(pulumi.StringArrayOutput)
}

func (o PrimaryRegionPropertiesOutput) OperatorAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrimaryRegionProperties) []string { return v.OperatorAddresses }).(pulumi.StringArrayOutput)
}

type PrimaryRegionPropertiesResponse struct {
	EsrpAddresses     []string `pulumi:"esrpAddresses"`
	OperatorAddresses []string `pulumi:"operatorAddresses"`
}

type PrimaryRegionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (PrimaryRegionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrimaryRegionPropertiesResponse)(nil)).Elem()
}

func (o PrimaryRegionPropertiesResponseOutput) ToPrimaryRegionPropertiesResponseOutput() PrimaryRegionPropertiesResponseOutput {
	return o
}

func (o PrimaryRegionPropertiesResponseOutput) ToPrimaryRegionPropertiesResponseOutputWithContext(ctx context.Context) PrimaryRegionPropertiesResponseOutput {
	return o
}

func (o PrimaryRegionPropertiesResponseOutput) EsrpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrimaryRegionPropertiesResponse) []string { return v.EsrpAddresses }).(pulumi.StringArrayOutput)
}

func (o PrimaryRegionPropertiesResponseOutput) OperatorAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrimaryRegionPropertiesResponse) []string { return v.OperatorAddresses }).(pulumi.StringArrayOutput)
}

type ServiceRegionProperties struct {
	Name                    string                  `pulumi:"name"`
	PrimaryRegionProperties PrimaryRegionProperties `pulumi:"primaryRegionProperties"`
}





type ServiceRegionPropertiesInput interface {
	pulumi.Input

	ToServiceRegionPropertiesOutput() ServiceRegionPropertiesOutput
	ToServiceRegionPropertiesOutputWithContext(context.Context) ServiceRegionPropertiesOutput
}

type ServiceRegionPropertiesArgs struct {
	Name                    pulumi.StringInput           `pulumi:"name"`
	PrimaryRegionProperties PrimaryRegionPropertiesInput `pulumi:"primaryRegionProperties"`
}

func (ServiceRegionPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceRegionProperties)(nil)).Elem()
}

func (i ServiceRegionPropertiesArgs) ToServiceRegionPropertiesOutput() ServiceRegionPropertiesOutput {
	return i.ToServiceRegionPropertiesOutputWithContext(context.Background())
}

func (i ServiceRegionPropertiesArgs) ToServiceRegionPropertiesOutputWithContext(ctx context.Context) ServiceRegionPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceRegionPropertiesOutput)
}





type ServiceRegionPropertiesArrayInput interface {
	pulumi.Input

	ToServiceRegionPropertiesArrayOutput() ServiceRegionPropertiesArrayOutput
	ToServiceRegionPropertiesArrayOutputWithContext(context.Context) ServiceRegionPropertiesArrayOutput
}

type ServiceRegionPropertiesArray []ServiceRegionPropertiesInput

func (ServiceRegionPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceRegionProperties)(nil)).Elem()
}

func (i ServiceRegionPropertiesArray) ToServiceRegionPropertiesArrayOutput() ServiceRegionPropertiesArrayOutput {
	return i.ToServiceRegionPropertiesArrayOutputWithContext(context.Background())
}

func (i ServiceRegionPropertiesArray) ToServiceRegionPropertiesArrayOutputWithContext(ctx context.Context) ServiceRegionPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceRegionPropertiesArrayOutput)
}

type ServiceRegionPropertiesOutput struct{ *pulumi.OutputState }

func (ServiceRegionPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceRegionProperties)(nil)).Elem()
}

func (o ServiceRegionPropertiesOutput) ToServiceRegionPropertiesOutput() ServiceRegionPropertiesOutput {
	return o
}

func (o ServiceRegionPropertiesOutput) ToServiceRegionPropertiesOutputWithContext(ctx context.Context) ServiceRegionPropertiesOutput {
	return o
}

func (o ServiceRegionPropertiesOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceRegionProperties) string { return v.Name }).(pulumi.StringOutput)
}

func (o ServiceRegionPropertiesOutput) PrimaryRegionProperties() PrimaryRegionPropertiesOutput {
	return o.ApplyT(func(v ServiceRegionProperties) PrimaryRegionProperties { return v.PrimaryRegionProperties }).(PrimaryRegionPropertiesOutput)
}

type ServiceRegionPropertiesArrayOutput struct{ *pulumi.OutputState }

func (ServiceRegionPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceRegionProperties)(nil)).Elem()
}

func (o ServiceRegionPropertiesArrayOutput) ToServiceRegionPropertiesArrayOutput() ServiceRegionPropertiesArrayOutput {
	return o
}

func (o ServiceRegionPropertiesArrayOutput) ToServiceRegionPropertiesArrayOutputWithContext(ctx context.Context) ServiceRegionPropertiesArrayOutput {
	return o
}

func (o ServiceRegionPropertiesArrayOutput) Index(i pulumi.IntInput) ServiceRegionPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceRegionProperties {
		return vs[0].([]ServiceRegionProperties)[vs[1].(int)]
	}).(ServiceRegionPropertiesOutput)
}

type ServiceRegionPropertiesResponse struct {
	Name                    string                          `pulumi:"name"`
	PrimaryRegionProperties PrimaryRegionPropertiesResponse `pulumi:"primaryRegionProperties"`
}

type ServiceRegionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ServiceRegionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceRegionPropertiesResponse)(nil)).Elem()
}

func (o ServiceRegionPropertiesResponseOutput) ToServiceRegionPropertiesResponseOutput() ServiceRegionPropertiesResponseOutput {
	return o
}

func (o ServiceRegionPropertiesResponseOutput) ToServiceRegionPropertiesResponseOutputWithContext(ctx context.Context) ServiceRegionPropertiesResponseOutput {
	return o
}

func (o ServiceRegionPropertiesResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceRegionPropertiesResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ServiceRegionPropertiesResponseOutput) PrimaryRegionProperties() PrimaryRegionPropertiesResponseOutput {
	return o.ApplyT(func(v ServiceRegionPropertiesResponse) PrimaryRegionPropertiesResponse {
		return v.PrimaryRegionProperties
	}).(PrimaryRegionPropertiesResponseOutput)
}

type ServiceRegionPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (ServiceRegionPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceRegionPropertiesResponse)(nil)).Elem()
}

func (o ServiceRegionPropertiesResponseArrayOutput) ToServiceRegionPropertiesResponseArrayOutput() ServiceRegionPropertiesResponseArrayOutput {
	return o
}

func (o ServiceRegionPropertiesResponseArrayOutput) ToServiceRegionPropertiesResponseArrayOutputWithContext(ctx context.Context) ServiceRegionPropertiesResponseArrayOutput {
	return o
}

func (o ServiceRegionPropertiesResponseArrayOutput) Index(i pulumi.IntInput) ServiceRegionPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceRegionPropertiesResponse {
		return vs[0].([]ServiceRegionPropertiesResponse)[vs[1].(int)]
	}).(ServiceRegionPropertiesResponseOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(PrimaryRegionPropertiesOutput{})
	pulumi.RegisterOutputType(PrimaryRegionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ServiceRegionPropertiesOutput{})
	pulumi.RegisterOutputType(ServiceRegionPropertiesArrayOutput{})
	pulumi.RegisterOutputType(ServiceRegionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ServiceRegionPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
