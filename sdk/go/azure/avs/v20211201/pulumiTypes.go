


package v20211201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AddonHcxProperties struct {
	AddonType string `pulumi:"addonType"`
	Offer     string `pulumi:"offer"`
}





type AddonHcxPropertiesInput interface {
	pulumi.Input

	ToAddonHcxPropertiesOutput() AddonHcxPropertiesOutput
	ToAddonHcxPropertiesOutputWithContext(context.Context) AddonHcxPropertiesOutput
}

type AddonHcxPropertiesArgs struct {
	AddonType pulumi.StringInput `pulumi:"addonType"`
	Offer     pulumi.StringInput `pulumi:"offer"`
}

func (AddonHcxPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AddonHcxProperties)(nil)).Elem()
}

func (i AddonHcxPropertiesArgs) ToAddonHcxPropertiesOutput() AddonHcxPropertiesOutput {
	return i.ToAddonHcxPropertiesOutputWithContext(context.Background())
}

func (i AddonHcxPropertiesArgs) ToAddonHcxPropertiesOutputWithContext(ctx context.Context) AddonHcxPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddonHcxPropertiesOutput)
}

type AddonHcxPropertiesOutput struct{ *pulumi.OutputState }

func (AddonHcxPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AddonHcxProperties)(nil)).Elem()
}

func (o AddonHcxPropertiesOutput) ToAddonHcxPropertiesOutput() AddonHcxPropertiesOutput {
	return o
}

func (o AddonHcxPropertiesOutput) ToAddonHcxPropertiesOutputWithContext(ctx context.Context) AddonHcxPropertiesOutput {
	return o
}

func (o AddonHcxPropertiesOutput) AddonType() pulumi.StringOutput {
	return o.ApplyT(func(v AddonHcxProperties) string { return v.AddonType }).(pulumi.StringOutput)
}

func (o AddonHcxPropertiesOutput) Offer() pulumi.StringOutput {
	return o.ApplyT(func(v AddonHcxProperties) string { return v.Offer }).(pulumi.StringOutput)
}

type AddonHcxPropertiesResponse struct {
	AddonType         string `pulumi:"addonType"`
	Offer             string `pulumi:"offer"`
	ProvisioningState string `pulumi:"provisioningState"`
}





type AddonHcxPropertiesResponseInput interface {
	pulumi.Input

	ToAddonHcxPropertiesResponseOutput() AddonHcxPropertiesResponseOutput
	ToAddonHcxPropertiesResponseOutputWithContext(context.Context) AddonHcxPropertiesResponseOutput
}

type AddonHcxPropertiesResponseArgs struct {
	AddonType         pulumi.StringInput `pulumi:"addonType"`
	Offer             pulumi.StringInput `pulumi:"offer"`
	ProvisioningState pulumi.StringInput `pulumi:"provisioningState"`
}

func (AddonHcxPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AddonHcxPropertiesResponse)(nil)).Elem()
}

func (i AddonHcxPropertiesResponseArgs) ToAddonHcxPropertiesResponseOutput() AddonHcxPropertiesResponseOutput {
	return i.ToAddonHcxPropertiesResponseOutputWithContext(context.Background())
}

func (i AddonHcxPropertiesResponseArgs) ToAddonHcxPropertiesResponseOutputWithContext(ctx context.Context) AddonHcxPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddonHcxPropertiesResponseOutput)
}

type AddonHcxPropertiesResponseOutput struct{ *pulumi.OutputState }

func (AddonHcxPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AddonHcxPropertiesResponse)(nil)).Elem()
}

func (o AddonHcxPropertiesResponseOutput) ToAddonHcxPropertiesResponseOutput() AddonHcxPropertiesResponseOutput {
	return o
}

func (o AddonHcxPropertiesResponseOutput) ToAddonHcxPropertiesResponseOutputWithContext(ctx context.Context) AddonHcxPropertiesResponseOutput {
	return o
}

func (o AddonHcxPropertiesResponseOutput) AddonType() pulumi.StringOutput {
	return o.ApplyT(func(v AddonHcxPropertiesResponse) string { return v.AddonType }).(pulumi.StringOutput)
}

func (o AddonHcxPropertiesResponseOutput) Offer() pulumi.StringOutput {
	return o.ApplyT(func(v AddonHcxPropertiesResponse) string { return v.Offer }).(pulumi.StringOutput)
}

func (o AddonHcxPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v AddonHcxPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type AddonSrmProperties struct {
	AddonType  string  `pulumi:"addonType"`
	LicenseKey *string `pulumi:"licenseKey"`
}





type AddonSrmPropertiesInput interface {
	pulumi.Input

	ToAddonSrmPropertiesOutput() AddonSrmPropertiesOutput
	ToAddonSrmPropertiesOutputWithContext(context.Context) AddonSrmPropertiesOutput
}

type AddonSrmPropertiesArgs struct {
	AddonType  pulumi.StringInput    `pulumi:"addonType"`
	LicenseKey pulumi.StringPtrInput `pulumi:"licenseKey"`
}

func (AddonSrmPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AddonSrmProperties)(nil)).Elem()
}

func (i AddonSrmPropertiesArgs) ToAddonSrmPropertiesOutput() AddonSrmPropertiesOutput {
	return i.ToAddonSrmPropertiesOutputWithContext(context.Background())
}

func (i AddonSrmPropertiesArgs) ToAddonSrmPropertiesOutputWithContext(ctx context.Context) AddonSrmPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddonSrmPropertiesOutput)
}

type AddonSrmPropertiesOutput struct{ *pulumi.OutputState }

func (AddonSrmPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AddonSrmProperties)(nil)).Elem()
}

func (o AddonSrmPropertiesOutput) ToAddonSrmPropertiesOutput() AddonSrmPropertiesOutput {
	return o
}

func (o AddonSrmPropertiesOutput) ToAddonSrmPropertiesOutputWithContext(ctx context.Context) AddonSrmPropertiesOutput {
	return o
}

func (o AddonSrmPropertiesOutput) AddonType() pulumi.StringOutput {
	return o.ApplyT(func(v AddonSrmProperties) string { return v.AddonType }).(pulumi.StringOutput)
}

func (o AddonSrmPropertiesOutput) LicenseKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AddonSrmProperties) *string { return v.LicenseKey }).(pulumi.StringPtrOutput)
}

type AddonSrmPropertiesResponse struct {
	AddonType         string  `pulumi:"addonType"`
	LicenseKey        *string `pulumi:"licenseKey"`
	ProvisioningState string  `pulumi:"provisioningState"`
}





type AddonSrmPropertiesResponseInput interface {
	pulumi.Input

	ToAddonSrmPropertiesResponseOutput() AddonSrmPropertiesResponseOutput
	ToAddonSrmPropertiesResponseOutputWithContext(context.Context) AddonSrmPropertiesResponseOutput
}

type AddonSrmPropertiesResponseArgs struct {
	AddonType         pulumi.StringInput    `pulumi:"addonType"`
	LicenseKey        pulumi.StringPtrInput `pulumi:"licenseKey"`
	ProvisioningState pulumi.StringInput    `pulumi:"provisioningState"`
}

func (AddonSrmPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AddonSrmPropertiesResponse)(nil)).Elem()
}

func (i AddonSrmPropertiesResponseArgs) ToAddonSrmPropertiesResponseOutput() AddonSrmPropertiesResponseOutput {
	return i.ToAddonSrmPropertiesResponseOutputWithContext(context.Background())
}

func (i AddonSrmPropertiesResponseArgs) ToAddonSrmPropertiesResponseOutputWithContext(ctx context.Context) AddonSrmPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddonSrmPropertiesResponseOutput)
}

type AddonSrmPropertiesResponseOutput struct{ *pulumi.OutputState }

func (AddonSrmPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AddonSrmPropertiesResponse)(nil)).Elem()
}

func (o AddonSrmPropertiesResponseOutput) ToAddonSrmPropertiesResponseOutput() AddonSrmPropertiesResponseOutput {
	return o
}

func (o AddonSrmPropertiesResponseOutput) ToAddonSrmPropertiesResponseOutputWithContext(ctx context.Context) AddonSrmPropertiesResponseOutput {
	return o
}

func (o AddonSrmPropertiesResponseOutput) AddonType() pulumi.StringOutput {
	return o.ApplyT(func(v AddonSrmPropertiesResponse) string { return v.AddonType }).(pulumi.StringOutput)
}

func (o AddonSrmPropertiesResponseOutput) LicenseKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AddonSrmPropertiesResponse) *string { return v.LicenseKey }).(pulumi.StringPtrOutput)
}

func (o AddonSrmPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v AddonSrmPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type AddonVrProperties struct {
	AddonType string `pulumi:"addonType"`
	VrsCount  int    `pulumi:"vrsCount"`
}





type AddonVrPropertiesInput interface {
	pulumi.Input

	ToAddonVrPropertiesOutput() AddonVrPropertiesOutput
	ToAddonVrPropertiesOutputWithContext(context.Context) AddonVrPropertiesOutput
}

type AddonVrPropertiesArgs struct {
	AddonType pulumi.StringInput `pulumi:"addonType"`
	VrsCount  pulumi.IntInput    `pulumi:"vrsCount"`
}

func (AddonVrPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AddonVrProperties)(nil)).Elem()
}

func (i AddonVrPropertiesArgs) ToAddonVrPropertiesOutput() AddonVrPropertiesOutput {
	return i.ToAddonVrPropertiesOutputWithContext(context.Background())
}

func (i AddonVrPropertiesArgs) ToAddonVrPropertiesOutputWithContext(ctx context.Context) AddonVrPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddonVrPropertiesOutput)
}

type AddonVrPropertiesOutput struct{ *pulumi.OutputState }

func (AddonVrPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AddonVrProperties)(nil)).Elem()
}

func (o AddonVrPropertiesOutput) ToAddonVrPropertiesOutput() AddonVrPropertiesOutput {
	return o
}

func (o AddonVrPropertiesOutput) ToAddonVrPropertiesOutputWithContext(ctx context.Context) AddonVrPropertiesOutput {
	return o
}

func (o AddonVrPropertiesOutput) AddonType() pulumi.StringOutput {
	return o.ApplyT(func(v AddonVrProperties) string { return v.AddonType }).(pulumi.StringOutput)
}

func (o AddonVrPropertiesOutput) VrsCount() pulumi.IntOutput {
	return o.ApplyT(func(v AddonVrProperties) int { return v.VrsCount }).(pulumi.IntOutput)
}

type AddonVrPropertiesResponse struct {
	AddonType         string `pulumi:"addonType"`
	ProvisioningState string `pulumi:"provisioningState"`
	VrsCount          int    `pulumi:"vrsCount"`
}





type AddonVrPropertiesResponseInput interface {
	pulumi.Input

	ToAddonVrPropertiesResponseOutput() AddonVrPropertiesResponseOutput
	ToAddonVrPropertiesResponseOutputWithContext(context.Context) AddonVrPropertiesResponseOutput
}

type AddonVrPropertiesResponseArgs struct {
	AddonType         pulumi.StringInput `pulumi:"addonType"`
	ProvisioningState pulumi.StringInput `pulumi:"provisioningState"`
	VrsCount          pulumi.IntInput    `pulumi:"vrsCount"`
}

func (AddonVrPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AddonVrPropertiesResponse)(nil)).Elem()
}

func (i AddonVrPropertiesResponseArgs) ToAddonVrPropertiesResponseOutput() AddonVrPropertiesResponseOutput {
	return i.ToAddonVrPropertiesResponseOutputWithContext(context.Background())
}

func (i AddonVrPropertiesResponseArgs) ToAddonVrPropertiesResponseOutputWithContext(ctx context.Context) AddonVrPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddonVrPropertiesResponseOutput)
}

type AddonVrPropertiesResponseOutput struct{ *pulumi.OutputState }

func (AddonVrPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AddonVrPropertiesResponse)(nil)).Elem()
}

func (o AddonVrPropertiesResponseOutput) ToAddonVrPropertiesResponseOutput() AddonVrPropertiesResponseOutput {
	return o
}

func (o AddonVrPropertiesResponseOutput) ToAddonVrPropertiesResponseOutputWithContext(ctx context.Context) AddonVrPropertiesResponseOutput {
	return o
}

func (o AddonVrPropertiesResponseOutput) AddonType() pulumi.StringOutput {
	return o.ApplyT(func(v AddonVrPropertiesResponse) string { return v.AddonType }).(pulumi.StringOutput)
}

func (o AddonVrPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v AddonVrPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o AddonVrPropertiesResponseOutput) VrsCount() pulumi.IntOutput {
	return o.ApplyT(func(v AddonVrPropertiesResponse) int { return v.VrsCount }).(pulumi.IntOutput)
}

type AvailabilityProperties struct {
	SecondaryZone *int    `pulumi:"secondaryZone"`
	Strategy      *string `pulumi:"strategy"`
	Zone          *int    `pulumi:"zone"`
}





type AvailabilityPropertiesInput interface {
	pulumi.Input

	ToAvailabilityPropertiesOutput() AvailabilityPropertiesOutput
	ToAvailabilityPropertiesOutputWithContext(context.Context) AvailabilityPropertiesOutput
}

type AvailabilityPropertiesArgs struct {
	SecondaryZone pulumi.IntPtrInput    `pulumi:"secondaryZone"`
	Strategy      pulumi.StringPtrInput `pulumi:"strategy"`
	Zone          pulumi.IntPtrInput    `pulumi:"zone"`
}

func (AvailabilityPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AvailabilityProperties)(nil)).Elem()
}

func (i AvailabilityPropertiesArgs) ToAvailabilityPropertiesOutput() AvailabilityPropertiesOutput {
	return i.ToAvailabilityPropertiesOutputWithContext(context.Background())
}

func (i AvailabilityPropertiesArgs) ToAvailabilityPropertiesOutputWithContext(ctx context.Context) AvailabilityPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AvailabilityPropertiesOutput)
}

func (i AvailabilityPropertiesArgs) ToAvailabilityPropertiesPtrOutput() AvailabilityPropertiesPtrOutput {
	return i.ToAvailabilityPropertiesPtrOutputWithContext(context.Background())
}

func (i AvailabilityPropertiesArgs) ToAvailabilityPropertiesPtrOutputWithContext(ctx context.Context) AvailabilityPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AvailabilityPropertiesOutput).ToAvailabilityPropertiesPtrOutputWithContext(ctx)
}









type AvailabilityPropertiesPtrInput interface {
	pulumi.Input

	ToAvailabilityPropertiesPtrOutput() AvailabilityPropertiesPtrOutput
	ToAvailabilityPropertiesPtrOutputWithContext(context.Context) AvailabilityPropertiesPtrOutput
}

type availabilityPropertiesPtrType AvailabilityPropertiesArgs

func AvailabilityPropertiesPtr(v *AvailabilityPropertiesArgs) AvailabilityPropertiesPtrInput {
	return (*availabilityPropertiesPtrType)(v)
}

func (*availabilityPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AvailabilityProperties)(nil)).Elem()
}

func (i *availabilityPropertiesPtrType) ToAvailabilityPropertiesPtrOutput() AvailabilityPropertiesPtrOutput {
	return i.ToAvailabilityPropertiesPtrOutputWithContext(context.Background())
}

func (i *availabilityPropertiesPtrType) ToAvailabilityPropertiesPtrOutputWithContext(ctx context.Context) AvailabilityPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AvailabilityPropertiesPtrOutput)
}

type AvailabilityPropertiesOutput struct{ *pulumi.OutputState }

func (AvailabilityPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AvailabilityProperties)(nil)).Elem()
}

func (o AvailabilityPropertiesOutput) ToAvailabilityPropertiesOutput() AvailabilityPropertiesOutput {
	return o
}

func (o AvailabilityPropertiesOutput) ToAvailabilityPropertiesOutputWithContext(ctx context.Context) AvailabilityPropertiesOutput {
	return o
}

func (o AvailabilityPropertiesOutput) ToAvailabilityPropertiesPtrOutput() AvailabilityPropertiesPtrOutput {
	return o.ToAvailabilityPropertiesPtrOutputWithContext(context.Background())
}

func (o AvailabilityPropertiesOutput) ToAvailabilityPropertiesPtrOutputWithContext(ctx context.Context) AvailabilityPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AvailabilityProperties) *AvailabilityProperties {
		return &v
	}).(AvailabilityPropertiesPtrOutput)
}

func (o AvailabilityPropertiesOutput) SecondaryZone() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AvailabilityProperties) *int { return v.SecondaryZone }).(pulumi.IntPtrOutput)
}

func (o AvailabilityPropertiesOutput) Strategy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AvailabilityProperties) *string { return v.Strategy }).(pulumi.StringPtrOutput)
}

func (o AvailabilityPropertiesOutput) Zone() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AvailabilityProperties) *int { return v.Zone }).(pulumi.IntPtrOutput)
}

type AvailabilityPropertiesPtrOutput struct{ *pulumi.OutputState }

func (AvailabilityPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AvailabilityProperties)(nil)).Elem()
}

func (o AvailabilityPropertiesPtrOutput) ToAvailabilityPropertiesPtrOutput() AvailabilityPropertiesPtrOutput {
	return o
}

func (o AvailabilityPropertiesPtrOutput) ToAvailabilityPropertiesPtrOutputWithContext(ctx context.Context) AvailabilityPropertiesPtrOutput {
	return o
}

func (o AvailabilityPropertiesPtrOutput) Elem() AvailabilityPropertiesOutput {
	return o.ApplyT(func(v *AvailabilityProperties) AvailabilityProperties {
		if v != nil {
			return *v
		}
		var ret AvailabilityProperties
		return ret
	}).(AvailabilityPropertiesOutput)
}

func (o AvailabilityPropertiesPtrOutput) SecondaryZone() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AvailabilityProperties) *int {
		if v == nil {
			return nil
		}
		return v.SecondaryZone
	}).(pulumi.IntPtrOutput)
}

func (o AvailabilityPropertiesPtrOutput) Strategy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AvailabilityProperties) *string {
		if v == nil {
			return nil
		}
		return v.Strategy
	}).(pulumi.StringPtrOutput)
}

func (o AvailabilityPropertiesPtrOutput) Zone() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AvailabilityProperties) *int {
		if v == nil {
			return nil
		}
		return v.Zone
	}).(pulumi.IntPtrOutput)
}

type AvailabilityPropertiesResponse struct {
	SecondaryZone *int    `pulumi:"secondaryZone"`
	Strategy      *string `pulumi:"strategy"`
	Zone          *int    `pulumi:"zone"`
}





type AvailabilityPropertiesResponseInput interface {
	pulumi.Input

	ToAvailabilityPropertiesResponseOutput() AvailabilityPropertiesResponseOutput
	ToAvailabilityPropertiesResponseOutputWithContext(context.Context) AvailabilityPropertiesResponseOutput
}

type AvailabilityPropertiesResponseArgs struct {
	SecondaryZone pulumi.IntPtrInput    `pulumi:"secondaryZone"`
	Strategy      pulumi.StringPtrInput `pulumi:"strategy"`
	Zone          pulumi.IntPtrInput    `pulumi:"zone"`
}

func (AvailabilityPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AvailabilityPropertiesResponse)(nil)).Elem()
}

func (i AvailabilityPropertiesResponseArgs) ToAvailabilityPropertiesResponseOutput() AvailabilityPropertiesResponseOutput {
	return i.ToAvailabilityPropertiesResponseOutputWithContext(context.Background())
}

func (i AvailabilityPropertiesResponseArgs) ToAvailabilityPropertiesResponseOutputWithContext(ctx context.Context) AvailabilityPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AvailabilityPropertiesResponseOutput)
}

func (i AvailabilityPropertiesResponseArgs) ToAvailabilityPropertiesResponsePtrOutput() AvailabilityPropertiesResponsePtrOutput {
	return i.ToAvailabilityPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i AvailabilityPropertiesResponseArgs) ToAvailabilityPropertiesResponsePtrOutputWithContext(ctx context.Context) AvailabilityPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AvailabilityPropertiesResponseOutput).ToAvailabilityPropertiesResponsePtrOutputWithContext(ctx)
}









type AvailabilityPropertiesResponsePtrInput interface {
	pulumi.Input

	ToAvailabilityPropertiesResponsePtrOutput() AvailabilityPropertiesResponsePtrOutput
	ToAvailabilityPropertiesResponsePtrOutputWithContext(context.Context) AvailabilityPropertiesResponsePtrOutput
}

type availabilityPropertiesResponsePtrType AvailabilityPropertiesResponseArgs

func AvailabilityPropertiesResponsePtr(v *AvailabilityPropertiesResponseArgs) AvailabilityPropertiesResponsePtrInput {
	return (*availabilityPropertiesResponsePtrType)(v)
}

func (*availabilityPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AvailabilityPropertiesResponse)(nil)).Elem()
}

func (i *availabilityPropertiesResponsePtrType) ToAvailabilityPropertiesResponsePtrOutput() AvailabilityPropertiesResponsePtrOutput {
	return i.ToAvailabilityPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *availabilityPropertiesResponsePtrType) ToAvailabilityPropertiesResponsePtrOutputWithContext(ctx context.Context) AvailabilityPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AvailabilityPropertiesResponsePtrOutput)
}

type AvailabilityPropertiesResponseOutput struct{ *pulumi.OutputState }

func (AvailabilityPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AvailabilityPropertiesResponse)(nil)).Elem()
}

func (o AvailabilityPropertiesResponseOutput) ToAvailabilityPropertiesResponseOutput() AvailabilityPropertiesResponseOutput {
	return o
}

func (o AvailabilityPropertiesResponseOutput) ToAvailabilityPropertiesResponseOutputWithContext(ctx context.Context) AvailabilityPropertiesResponseOutput {
	return o
}

func (o AvailabilityPropertiesResponseOutput) ToAvailabilityPropertiesResponsePtrOutput() AvailabilityPropertiesResponsePtrOutput {
	return o.ToAvailabilityPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o AvailabilityPropertiesResponseOutput) ToAvailabilityPropertiesResponsePtrOutputWithContext(ctx context.Context) AvailabilityPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AvailabilityPropertiesResponse) *AvailabilityPropertiesResponse {
		return &v
	}).(AvailabilityPropertiesResponsePtrOutput)
}

func (o AvailabilityPropertiesResponseOutput) SecondaryZone() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AvailabilityPropertiesResponse) *int { return v.SecondaryZone }).(pulumi.IntPtrOutput)
}

func (o AvailabilityPropertiesResponseOutput) Strategy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AvailabilityPropertiesResponse) *string { return v.Strategy }).(pulumi.StringPtrOutput)
}

func (o AvailabilityPropertiesResponseOutput) Zone() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AvailabilityPropertiesResponse) *int { return v.Zone }).(pulumi.IntPtrOutput)
}

type AvailabilityPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (AvailabilityPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AvailabilityPropertiesResponse)(nil)).Elem()
}

func (o AvailabilityPropertiesResponsePtrOutput) ToAvailabilityPropertiesResponsePtrOutput() AvailabilityPropertiesResponsePtrOutput {
	return o
}

func (o AvailabilityPropertiesResponsePtrOutput) ToAvailabilityPropertiesResponsePtrOutputWithContext(ctx context.Context) AvailabilityPropertiesResponsePtrOutput {
	return o
}

func (o AvailabilityPropertiesResponsePtrOutput) Elem() AvailabilityPropertiesResponseOutput {
	return o.ApplyT(func(v *AvailabilityPropertiesResponse) AvailabilityPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret AvailabilityPropertiesResponse
		return ret
	}).(AvailabilityPropertiesResponseOutput)
}

func (o AvailabilityPropertiesResponsePtrOutput) SecondaryZone() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AvailabilityPropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return v.SecondaryZone
	}).(pulumi.IntPtrOutput)
}

func (o AvailabilityPropertiesResponsePtrOutput) Strategy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AvailabilityPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Strategy
	}).(pulumi.StringPtrOutput)
}

func (o AvailabilityPropertiesResponsePtrOutput) Zone() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AvailabilityPropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return v.Zone
	}).(pulumi.IntPtrOutput)
}

type CircuitResponse struct {
	ExpressRouteID               string `pulumi:"expressRouteID"`
	ExpressRoutePrivatePeeringID string `pulumi:"expressRoutePrivatePeeringID"`
	PrimarySubnet                string `pulumi:"primarySubnet"`
	SecondarySubnet              string `pulumi:"secondarySubnet"`
}





type CircuitResponseInput interface {
	pulumi.Input

	ToCircuitResponseOutput() CircuitResponseOutput
	ToCircuitResponseOutputWithContext(context.Context) CircuitResponseOutput
}

type CircuitResponseArgs struct {
	ExpressRouteID               pulumi.StringInput `pulumi:"expressRouteID"`
	ExpressRoutePrivatePeeringID pulumi.StringInput `pulumi:"expressRoutePrivatePeeringID"`
	PrimarySubnet                pulumi.StringInput `pulumi:"primarySubnet"`
	SecondarySubnet              pulumi.StringInput `pulumi:"secondarySubnet"`
}

func (CircuitResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CircuitResponse)(nil)).Elem()
}

func (i CircuitResponseArgs) ToCircuitResponseOutput() CircuitResponseOutput {
	return i.ToCircuitResponseOutputWithContext(context.Background())
}

func (i CircuitResponseArgs) ToCircuitResponseOutputWithContext(ctx context.Context) CircuitResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CircuitResponseOutput)
}

func (i CircuitResponseArgs) ToCircuitResponsePtrOutput() CircuitResponsePtrOutput {
	return i.ToCircuitResponsePtrOutputWithContext(context.Background())
}

func (i CircuitResponseArgs) ToCircuitResponsePtrOutputWithContext(ctx context.Context) CircuitResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CircuitResponseOutput).ToCircuitResponsePtrOutputWithContext(ctx)
}









type CircuitResponsePtrInput interface {
	pulumi.Input

	ToCircuitResponsePtrOutput() CircuitResponsePtrOutput
	ToCircuitResponsePtrOutputWithContext(context.Context) CircuitResponsePtrOutput
}

type circuitResponsePtrType CircuitResponseArgs

func CircuitResponsePtr(v *CircuitResponseArgs) CircuitResponsePtrInput {
	return (*circuitResponsePtrType)(v)
}

func (*circuitResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CircuitResponse)(nil)).Elem()
}

func (i *circuitResponsePtrType) ToCircuitResponsePtrOutput() CircuitResponsePtrOutput {
	return i.ToCircuitResponsePtrOutputWithContext(context.Background())
}

func (i *circuitResponsePtrType) ToCircuitResponsePtrOutputWithContext(ctx context.Context) CircuitResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CircuitResponsePtrOutput)
}

type CircuitResponseOutput struct{ *pulumi.OutputState }

func (CircuitResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CircuitResponse)(nil)).Elem()
}

func (o CircuitResponseOutput) ToCircuitResponseOutput() CircuitResponseOutput {
	return o
}

func (o CircuitResponseOutput) ToCircuitResponseOutputWithContext(ctx context.Context) CircuitResponseOutput {
	return o
}

func (o CircuitResponseOutput) ToCircuitResponsePtrOutput() CircuitResponsePtrOutput {
	return o.ToCircuitResponsePtrOutputWithContext(context.Background())
}

func (o CircuitResponseOutput) ToCircuitResponsePtrOutputWithContext(ctx context.Context) CircuitResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CircuitResponse) *CircuitResponse {
		return &v
	}).(CircuitResponsePtrOutput)
}

func (o CircuitResponseOutput) ExpressRouteID() pulumi.StringOutput {
	return o.ApplyT(func(v CircuitResponse) string { return v.ExpressRouteID }).(pulumi.StringOutput)
}

func (o CircuitResponseOutput) ExpressRoutePrivatePeeringID() pulumi.StringOutput {
	return o.ApplyT(func(v CircuitResponse) string { return v.ExpressRoutePrivatePeeringID }).(pulumi.StringOutput)
}

func (o CircuitResponseOutput) PrimarySubnet() pulumi.StringOutput {
	return o.ApplyT(func(v CircuitResponse) string { return v.PrimarySubnet }).(pulumi.StringOutput)
}

func (o CircuitResponseOutput) SecondarySubnet() pulumi.StringOutput {
	return o.ApplyT(func(v CircuitResponse) string { return v.SecondarySubnet }).(pulumi.StringOutput)
}

type CircuitResponsePtrOutput struct{ *pulumi.OutputState }

func (CircuitResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CircuitResponse)(nil)).Elem()
}

func (o CircuitResponsePtrOutput) ToCircuitResponsePtrOutput() CircuitResponsePtrOutput {
	return o
}

func (o CircuitResponsePtrOutput) ToCircuitResponsePtrOutputWithContext(ctx context.Context) CircuitResponsePtrOutput {
	return o
}

func (o CircuitResponsePtrOutput) Elem() CircuitResponseOutput {
	return o.ApplyT(func(v *CircuitResponse) CircuitResponse {
		if v != nil {
			return *v
		}
		var ret CircuitResponse
		return ret
	}).(CircuitResponseOutput)
}

func (o CircuitResponsePtrOutput) ExpressRouteID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CircuitResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ExpressRouteID
	}).(pulumi.StringPtrOutput)
}

func (o CircuitResponsePtrOutput) ExpressRoutePrivatePeeringID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CircuitResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ExpressRoutePrivatePeeringID
	}).(pulumi.StringPtrOutput)
}

func (o CircuitResponsePtrOutput) PrimarySubnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CircuitResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrimarySubnet
	}).(pulumi.StringPtrOutput)
}

func (o CircuitResponsePtrOutput) SecondarySubnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CircuitResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SecondarySubnet
	}).(pulumi.StringPtrOutput)
}

type DiskPoolVolume struct {
	LunName     string  `pulumi:"lunName"`
	MountOption *string `pulumi:"mountOption"`
	TargetId    string  `pulumi:"targetId"`
}





type DiskPoolVolumeInput interface {
	pulumi.Input

	ToDiskPoolVolumeOutput() DiskPoolVolumeOutput
	ToDiskPoolVolumeOutputWithContext(context.Context) DiskPoolVolumeOutput
}

type DiskPoolVolumeArgs struct {
	LunName     pulumi.StringInput    `pulumi:"lunName"`
	MountOption pulumi.StringPtrInput `pulumi:"mountOption"`
	TargetId    pulumi.StringInput    `pulumi:"targetId"`
}

func (DiskPoolVolumeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskPoolVolume)(nil)).Elem()
}

func (i DiskPoolVolumeArgs) ToDiskPoolVolumeOutput() DiskPoolVolumeOutput {
	return i.ToDiskPoolVolumeOutputWithContext(context.Background())
}

func (i DiskPoolVolumeArgs) ToDiskPoolVolumeOutputWithContext(ctx context.Context) DiskPoolVolumeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskPoolVolumeOutput)
}

func (i DiskPoolVolumeArgs) ToDiskPoolVolumePtrOutput() DiskPoolVolumePtrOutput {
	return i.ToDiskPoolVolumePtrOutputWithContext(context.Background())
}

func (i DiskPoolVolumeArgs) ToDiskPoolVolumePtrOutputWithContext(ctx context.Context) DiskPoolVolumePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskPoolVolumeOutput).ToDiskPoolVolumePtrOutputWithContext(ctx)
}









type DiskPoolVolumePtrInput interface {
	pulumi.Input

	ToDiskPoolVolumePtrOutput() DiskPoolVolumePtrOutput
	ToDiskPoolVolumePtrOutputWithContext(context.Context) DiskPoolVolumePtrOutput
}

type diskPoolVolumePtrType DiskPoolVolumeArgs

func DiskPoolVolumePtr(v *DiskPoolVolumeArgs) DiskPoolVolumePtrInput {
	return (*diskPoolVolumePtrType)(v)
}

func (*diskPoolVolumePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskPoolVolume)(nil)).Elem()
}

func (i *diskPoolVolumePtrType) ToDiskPoolVolumePtrOutput() DiskPoolVolumePtrOutput {
	return i.ToDiskPoolVolumePtrOutputWithContext(context.Background())
}

func (i *diskPoolVolumePtrType) ToDiskPoolVolumePtrOutputWithContext(ctx context.Context) DiskPoolVolumePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskPoolVolumePtrOutput)
}

type DiskPoolVolumeOutput struct{ *pulumi.OutputState }

func (DiskPoolVolumeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskPoolVolume)(nil)).Elem()
}

func (o DiskPoolVolumeOutput) ToDiskPoolVolumeOutput() DiskPoolVolumeOutput {
	return o
}

func (o DiskPoolVolumeOutput) ToDiskPoolVolumeOutputWithContext(ctx context.Context) DiskPoolVolumeOutput {
	return o
}

func (o DiskPoolVolumeOutput) ToDiskPoolVolumePtrOutput() DiskPoolVolumePtrOutput {
	return o.ToDiskPoolVolumePtrOutputWithContext(context.Background())
}

func (o DiskPoolVolumeOutput) ToDiskPoolVolumePtrOutputWithContext(ctx context.Context) DiskPoolVolumePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiskPoolVolume) *DiskPoolVolume {
		return &v
	}).(DiskPoolVolumePtrOutput)
}

func (o DiskPoolVolumeOutput) LunName() pulumi.StringOutput {
	return o.ApplyT(func(v DiskPoolVolume) string { return v.LunName }).(pulumi.StringOutput)
}

func (o DiskPoolVolumeOutput) MountOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiskPoolVolume) *string { return v.MountOption }).(pulumi.StringPtrOutput)
}

func (o DiskPoolVolumeOutput) TargetId() pulumi.StringOutput {
	return o.ApplyT(func(v DiskPoolVolume) string { return v.TargetId }).(pulumi.StringOutput)
}

type DiskPoolVolumePtrOutput struct{ *pulumi.OutputState }

func (DiskPoolVolumePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskPoolVolume)(nil)).Elem()
}

func (o DiskPoolVolumePtrOutput) ToDiskPoolVolumePtrOutput() DiskPoolVolumePtrOutput {
	return o
}

func (o DiskPoolVolumePtrOutput) ToDiskPoolVolumePtrOutputWithContext(ctx context.Context) DiskPoolVolumePtrOutput {
	return o
}

func (o DiskPoolVolumePtrOutput) Elem() DiskPoolVolumeOutput {
	return o.ApplyT(func(v *DiskPoolVolume) DiskPoolVolume {
		if v != nil {
			return *v
		}
		var ret DiskPoolVolume
		return ret
	}).(DiskPoolVolumeOutput)
}

func (o DiskPoolVolumePtrOutput) LunName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiskPoolVolume) *string {
		if v == nil {
			return nil
		}
		return &v.LunName
	}).(pulumi.StringPtrOutput)
}

func (o DiskPoolVolumePtrOutput) MountOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiskPoolVolume) *string {
		if v == nil {
			return nil
		}
		return v.MountOption
	}).(pulumi.StringPtrOutput)
}

func (o DiskPoolVolumePtrOutput) TargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiskPoolVolume) *string {
		if v == nil {
			return nil
		}
		return &v.TargetId
	}).(pulumi.StringPtrOutput)
}

type DiskPoolVolumeResponse struct {
	LunName     string  `pulumi:"lunName"`
	MountOption *string `pulumi:"mountOption"`
	Path        string  `pulumi:"path"`
	TargetId    string  `pulumi:"targetId"`
}





type DiskPoolVolumeResponseInput interface {
	pulumi.Input

	ToDiskPoolVolumeResponseOutput() DiskPoolVolumeResponseOutput
	ToDiskPoolVolumeResponseOutputWithContext(context.Context) DiskPoolVolumeResponseOutput
}

type DiskPoolVolumeResponseArgs struct {
	LunName     pulumi.StringInput    `pulumi:"lunName"`
	MountOption pulumi.StringPtrInput `pulumi:"mountOption"`
	Path        pulumi.StringInput    `pulumi:"path"`
	TargetId    pulumi.StringInput    `pulumi:"targetId"`
}

func (DiskPoolVolumeResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskPoolVolumeResponse)(nil)).Elem()
}

func (i DiskPoolVolumeResponseArgs) ToDiskPoolVolumeResponseOutput() DiskPoolVolumeResponseOutput {
	return i.ToDiskPoolVolumeResponseOutputWithContext(context.Background())
}

func (i DiskPoolVolumeResponseArgs) ToDiskPoolVolumeResponseOutputWithContext(ctx context.Context) DiskPoolVolumeResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskPoolVolumeResponseOutput)
}

func (i DiskPoolVolumeResponseArgs) ToDiskPoolVolumeResponsePtrOutput() DiskPoolVolumeResponsePtrOutput {
	return i.ToDiskPoolVolumeResponsePtrOutputWithContext(context.Background())
}

func (i DiskPoolVolumeResponseArgs) ToDiskPoolVolumeResponsePtrOutputWithContext(ctx context.Context) DiskPoolVolumeResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskPoolVolumeResponseOutput).ToDiskPoolVolumeResponsePtrOutputWithContext(ctx)
}









type DiskPoolVolumeResponsePtrInput interface {
	pulumi.Input

	ToDiskPoolVolumeResponsePtrOutput() DiskPoolVolumeResponsePtrOutput
	ToDiskPoolVolumeResponsePtrOutputWithContext(context.Context) DiskPoolVolumeResponsePtrOutput
}

type diskPoolVolumeResponsePtrType DiskPoolVolumeResponseArgs

func DiskPoolVolumeResponsePtr(v *DiskPoolVolumeResponseArgs) DiskPoolVolumeResponsePtrInput {
	return (*diskPoolVolumeResponsePtrType)(v)
}

func (*diskPoolVolumeResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskPoolVolumeResponse)(nil)).Elem()
}

func (i *diskPoolVolumeResponsePtrType) ToDiskPoolVolumeResponsePtrOutput() DiskPoolVolumeResponsePtrOutput {
	return i.ToDiskPoolVolumeResponsePtrOutputWithContext(context.Background())
}

func (i *diskPoolVolumeResponsePtrType) ToDiskPoolVolumeResponsePtrOutputWithContext(ctx context.Context) DiskPoolVolumeResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskPoolVolumeResponsePtrOutput)
}

type DiskPoolVolumeResponseOutput struct{ *pulumi.OutputState }

func (DiskPoolVolumeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskPoolVolumeResponse)(nil)).Elem()
}

func (o DiskPoolVolumeResponseOutput) ToDiskPoolVolumeResponseOutput() DiskPoolVolumeResponseOutput {
	return o
}

func (o DiskPoolVolumeResponseOutput) ToDiskPoolVolumeResponseOutputWithContext(ctx context.Context) DiskPoolVolumeResponseOutput {
	return o
}

func (o DiskPoolVolumeResponseOutput) ToDiskPoolVolumeResponsePtrOutput() DiskPoolVolumeResponsePtrOutput {
	return o.ToDiskPoolVolumeResponsePtrOutputWithContext(context.Background())
}

func (o DiskPoolVolumeResponseOutput) ToDiskPoolVolumeResponsePtrOutputWithContext(ctx context.Context) DiskPoolVolumeResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiskPoolVolumeResponse) *DiskPoolVolumeResponse {
		return &v
	}).(DiskPoolVolumeResponsePtrOutput)
}

func (o DiskPoolVolumeResponseOutput) LunName() pulumi.StringOutput {
	return o.ApplyT(func(v DiskPoolVolumeResponse) string { return v.LunName }).(pulumi.StringOutput)
}

func (o DiskPoolVolumeResponseOutput) MountOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiskPoolVolumeResponse) *string { return v.MountOption }).(pulumi.StringPtrOutput)
}

func (o DiskPoolVolumeResponseOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v DiskPoolVolumeResponse) string { return v.Path }).(pulumi.StringOutput)
}

func (o DiskPoolVolumeResponseOutput) TargetId() pulumi.StringOutput {
	return o.ApplyT(func(v DiskPoolVolumeResponse) string { return v.TargetId }).(pulumi.StringOutput)
}

type DiskPoolVolumeResponsePtrOutput struct{ *pulumi.OutputState }

func (DiskPoolVolumeResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskPoolVolumeResponse)(nil)).Elem()
}

func (o DiskPoolVolumeResponsePtrOutput) ToDiskPoolVolumeResponsePtrOutput() DiskPoolVolumeResponsePtrOutput {
	return o
}

func (o DiskPoolVolumeResponsePtrOutput) ToDiskPoolVolumeResponsePtrOutputWithContext(ctx context.Context) DiskPoolVolumeResponsePtrOutput {
	return o
}

func (o DiskPoolVolumeResponsePtrOutput) Elem() DiskPoolVolumeResponseOutput {
	return o.ApplyT(func(v *DiskPoolVolumeResponse) DiskPoolVolumeResponse {
		if v != nil {
			return *v
		}
		var ret DiskPoolVolumeResponse
		return ret
	}).(DiskPoolVolumeResponseOutput)
}

func (o DiskPoolVolumeResponsePtrOutput) LunName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiskPoolVolumeResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LunName
	}).(pulumi.StringPtrOutput)
}

func (o DiskPoolVolumeResponsePtrOutput) MountOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiskPoolVolumeResponse) *string {
		if v == nil {
			return nil
		}
		return v.MountOption
	}).(pulumi.StringPtrOutput)
}

func (o DiskPoolVolumeResponsePtrOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiskPoolVolumeResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Path
	}).(pulumi.StringPtrOutput)
}

func (o DiskPoolVolumeResponsePtrOutput) TargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiskPoolVolumeResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TargetId
	}).(pulumi.StringPtrOutput)
}

type Encryption struct {
	KeyVaultProperties *EncryptionKeyVaultProperties `pulumi:"keyVaultProperties"`
	Status             *string                       `pulumi:"status"`
}





type EncryptionInput interface {
	pulumi.Input

	ToEncryptionOutput() EncryptionOutput
	ToEncryptionOutputWithContext(context.Context) EncryptionOutput
}

type EncryptionArgs struct {
	KeyVaultProperties EncryptionKeyVaultPropertiesPtrInput `pulumi:"keyVaultProperties"`
	Status             pulumi.StringPtrInput                `pulumi:"status"`
}

func (EncryptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Encryption)(nil)).Elem()
}

func (i EncryptionArgs) ToEncryptionOutput() EncryptionOutput {
	return i.ToEncryptionOutputWithContext(context.Background())
}

func (i EncryptionArgs) ToEncryptionOutputWithContext(ctx context.Context) EncryptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionOutput)
}

func (i EncryptionArgs) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return i.ToEncryptionPtrOutputWithContext(context.Background())
}

func (i EncryptionArgs) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionOutput).ToEncryptionPtrOutputWithContext(ctx)
}









type EncryptionPtrInput interface {
	pulumi.Input

	ToEncryptionPtrOutput() EncryptionPtrOutput
	ToEncryptionPtrOutputWithContext(context.Context) EncryptionPtrOutput
}

type encryptionPtrType EncryptionArgs

func EncryptionPtr(v *EncryptionArgs) EncryptionPtrInput {
	return (*encryptionPtrType)(v)
}

func (*encryptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Encryption)(nil)).Elem()
}

func (i *encryptionPtrType) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return i.ToEncryptionPtrOutputWithContext(context.Background())
}

func (i *encryptionPtrType) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPtrOutput)
}

type EncryptionOutput struct{ *pulumi.OutputState }

func (EncryptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Encryption)(nil)).Elem()
}

func (o EncryptionOutput) ToEncryptionOutput() EncryptionOutput {
	return o
}

func (o EncryptionOutput) ToEncryptionOutputWithContext(ctx context.Context) EncryptionOutput {
	return o
}

func (o EncryptionOutput) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return o.ToEncryptionPtrOutputWithContext(context.Background())
}

func (o EncryptionOutput) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Encryption) *Encryption {
		return &v
	}).(EncryptionPtrOutput)
}

func (o EncryptionOutput) KeyVaultProperties() EncryptionKeyVaultPropertiesPtrOutput {
	return o.ApplyT(func(v Encryption) *EncryptionKeyVaultProperties { return v.KeyVaultProperties }).(EncryptionKeyVaultPropertiesPtrOutput)
}

func (o EncryptionOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Encryption) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type EncryptionPtrOutput struct{ *pulumi.OutputState }

func (EncryptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Encryption)(nil)).Elem()
}

func (o EncryptionPtrOutput) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return o
}

func (o EncryptionPtrOutput) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return o
}

func (o EncryptionPtrOutput) Elem() EncryptionOutput {
	return o.ApplyT(func(v *Encryption) Encryption {
		if v != nil {
			return *v
		}
		var ret Encryption
		return ret
	}).(EncryptionOutput)
}

func (o EncryptionPtrOutput) KeyVaultProperties() EncryptionKeyVaultPropertiesPtrOutput {
	return o.ApplyT(func(v *Encryption) *EncryptionKeyVaultProperties {
		if v == nil {
			return nil
		}
		return v.KeyVaultProperties
	}).(EncryptionKeyVaultPropertiesPtrOutput)
}

func (o EncryptionPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Encryption) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type EncryptionKeyVaultProperties struct {
	KeyName     *string `pulumi:"keyName"`
	KeyVaultUrl *string `pulumi:"keyVaultUrl"`
	KeyVersion  *string `pulumi:"keyVersion"`
}





type EncryptionKeyVaultPropertiesInput interface {
	pulumi.Input

	ToEncryptionKeyVaultPropertiesOutput() EncryptionKeyVaultPropertiesOutput
	ToEncryptionKeyVaultPropertiesOutputWithContext(context.Context) EncryptionKeyVaultPropertiesOutput
}

type EncryptionKeyVaultPropertiesArgs struct {
	KeyName     pulumi.StringPtrInput `pulumi:"keyName"`
	KeyVaultUrl pulumi.StringPtrInput `pulumi:"keyVaultUrl"`
	KeyVersion  pulumi.StringPtrInput `pulumi:"keyVersion"`
}

func (EncryptionKeyVaultPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionKeyVaultProperties)(nil)).Elem()
}

func (i EncryptionKeyVaultPropertiesArgs) ToEncryptionKeyVaultPropertiesOutput() EncryptionKeyVaultPropertiesOutput {
	return i.ToEncryptionKeyVaultPropertiesOutputWithContext(context.Background())
}

func (i EncryptionKeyVaultPropertiesArgs) ToEncryptionKeyVaultPropertiesOutputWithContext(ctx context.Context) EncryptionKeyVaultPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionKeyVaultPropertiesOutput)
}

func (i EncryptionKeyVaultPropertiesArgs) ToEncryptionKeyVaultPropertiesPtrOutput() EncryptionKeyVaultPropertiesPtrOutput {
	return i.ToEncryptionKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (i EncryptionKeyVaultPropertiesArgs) ToEncryptionKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) EncryptionKeyVaultPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionKeyVaultPropertiesOutput).ToEncryptionKeyVaultPropertiesPtrOutputWithContext(ctx)
}









type EncryptionKeyVaultPropertiesPtrInput interface {
	pulumi.Input

	ToEncryptionKeyVaultPropertiesPtrOutput() EncryptionKeyVaultPropertiesPtrOutput
	ToEncryptionKeyVaultPropertiesPtrOutputWithContext(context.Context) EncryptionKeyVaultPropertiesPtrOutput
}

type encryptionKeyVaultPropertiesPtrType EncryptionKeyVaultPropertiesArgs

func EncryptionKeyVaultPropertiesPtr(v *EncryptionKeyVaultPropertiesArgs) EncryptionKeyVaultPropertiesPtrInput {
	return (*encryptionKeyVaultPropertiesPtrType)(v)
}

func (*encryptionKeyVaultPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionKeyVaultProperties)(nil)).Elem()
}

func (i *encryptionKeyVaultPropertiesPtrType) ToEncryptionKeyVaultPropertiesPtrOutput() EncryptionKeyVaultPropertiesPtrOutput {
	return i.ToEncryptionKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (i *encryptionKeyVaultPropertiesPtrType) ToEncryptionKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) EncryptionKeyVaultPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionKeyVaultPropertiesPtrOutput)
}

type EncryptionKeyVaultPropertiesOutput struct{ *pulumi.OutputState }

func (EncryptionKeyVaultPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionKeyVaultProperties)(nil)).Elem()
}

func (o EncryptionKeyVaultPropertiesOutput) ToEncryptionKeyVaultPropertiesOutput() EncryptionKeyVaultPropertiesOutput {
	return o
}

func (o EncryptionKeyVaultPropertiesOutput) ToEncryptionKeyVaultPropertiesOutputWithContext(ctx context.Context) EncryptionKeyVaultPropertiesOutput {
	return o
}

func (o EncryptionKeyVaultPropertiesOutput) ToEncryptionKeyVaultPropertiesPtrOutput() EncryptionKeyVaultPropertiesPtrOutput {
	return o.ToEncryptionKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (o EncryptionKeyVaultPropertiesOutput) ToEncryptionKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) EncryptionKeyVaultPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionKeyVaultProperties) *EncryptionKeyVaultProperties {
		return &v
	}).(EncryptionKeyVaultPropertiesPtrOutput)
}

func (o EncryptionKeyVaultPropertiesOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionKeyVaultProperties) *string { return v.KeyName }).(pulumi.StringPtrOutput)
}

func (o EncryptionKeyVaultPropertiesOutput) KeyVaultUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionKeyVaultProperties) *string { return v.KeyVaultUrl }).(pulumi.StringPtrOutput)
}

func (o EncryptionKeyVaultPropertiesOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionKeyVaultProperties) *string { return v.KeyVersion }).(pulumi.StringPtrOutput)
}

type EncryptionKeyVaultPropertiesPtrOutput struct{ *pulumi.OutputState }

func (EncryptionKeyVaultPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionKeyVaultProperties)(nil)).Elem()
}

func (o EncryptionKeyVaultPropertiesPtrOutput) ToEncryptionKeyVaultPropertiesPtrOutput() EncryptionKeyVaultPropertiesPtrOutput {
	return o
}

func (o EncryptionKeyVaultPropertiesPtrOutput) ToEncryptionKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) EncryptionKeyVaultPropertiesPtrOutput {
	return o
}

func (o EncryptionKeyVaultPropertiesPtrOutput) Elem() EncryptionKeyVaultPropertiesOutput {
	return o.ApplyT(func(v *EncryptionKeyVaultProperties) EncryptionKeyVaultProperties {
		if v != nil {
			return *v
		}
		var ret EncryptionKeyVaultProperties
		return ret
	}).(EncryptionKeyVaultPropertiesOutput)
}

func (o EncryptionKeyVaultPropertiesPtrOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionKeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.KeyName
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionKeyVaultPropertiesPtrOutput) KeyVaultUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionKeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultUrl
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionKeyVaultPropertiesPtrOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionKeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.KeyVersion
	}).(pulumi.StringPtrOutput)
}

type EncryptionKeyVaultPropertiesResponse struct {
	KeyName     *string `pulumi:"keyName"`
	KeyState    string  `pulumi:"keyState"`
	KeyVaultUrl *string `pulumi:"keyVaultUrl"`
	KeyVersion  *string `pulumi:"keyVersion"`
	VersionType string  `pulumi:"versionType"`
}





type EncryptionKeyVaultPropertiesResponseInput interface {
	pulumi.Input

	ToEncryptionKeyVaultPropertiesResponseOutput() EncryptionKeyVaultPropertiesResponseOutput
	ToEncryptionKeyVaultPropertiesResponseOutputWithContext(context.Context) EncryptionKeyVaultPropertiesResponseOutput
}

type EncryptionKeyVaultPropertiesResponseArgs struct {
	KeyName     pulumi.StringPtrInput `pulumi:"keyName"`
	KeyState    pulumi.StringInput    `pulumi:"keyState"`
	KeyVaultUrl pulumi.StringPtrInput `pulumi:"keyVaultUrl"`
	KeyVersion  pulumi.StringPtrInput `pulumi:"keyVersion"`
	VersionType pulumi.StringInput    `pulumi:"versionType"`
}

func (EncryptionKeyVaultPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionKeyVaultPropertiesResponse)(nil)).Elem()
}

func (i EncryptionKeyVaultPropertiesResponseArgs) ToEncryptionKeyVaultPropertiesResponseOutput() EncryptionKeyVaultPropertiesResponseOutput {
	return i.ToEncryptionKeyVaultPropertiesResponseOutputWithContext(context.Background())
}

func (i EncryptionKeyVaultPropertiesResponseArgs) ToEncryptionKeyVaultPropertiesResponseOutputWithContext(ctx context.Context) EncryptionKeyVaultPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionKeyVaultPropertiesResponseOutput)
}

func (i EncryptionKeyVaultPropertiesResponseArgs) ToEncryptionKeyVaultPropertiesResponsePtrOutput() EncryptionKeyVaultPropertiesResponsePtrOutput {
	return i.ToEncryptionKeyVaultPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i EncryptionKeyVaultPropertiesResponseArgs) ToEncryptionKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) EncryptionKeyVaultPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionKeyVaultPropertiesResponseOutput).ToEncryptionKeyVaultPropertiesResponsePtrOutputWithContext(ctx)
}









type EncryptionKeyVaultPropertiesResponsePtrInput interface {
	pulumi.Input

	ToEncryptionKeyVaultPropertiesResponsePtrOutput() EncryptionKeyVaultPropertiesResponsePtrOutput
	ToEncryptionKeyVaultPropertiesResponsePtrOutputWithContext(context.Context) EncryptionKeyVaultPropertiesResponsePtrOutput
}

type encryptionKeyVaultPropertiesResponsePtrType EncryptionKeyVaultPropertiesResponseArgs

func EncryptionKeyVaultPropertiesResponsePtr(v *EncryptionKeyVaultPropertiesResponseArgs) EncryptionKeyVaultPropertiesResponsePtrInput {
	return (*encryptionKeyVaultPropertiesResponsePtrType)(v)
}

func (*encryptionKeyVaultPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionKeyVaultPropertiesResponse)(nil)).Elem()
}

func (i *encryptionKeyVaultPropertiesResponsePtrType) ToEncryptionKeyVaultPropertiesResponsePtrOutput() EncryptionKeyVaultPropertiesResponsePtrOutput {
	return i.ToEncryptionKeyVaultPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *encryptionKeyVaultPropertiesResponsePtrType) ToEncryptionKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) EncryptionKeyVaultPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionKeyVaultPropertiesResponsePtrOutput)
}

type EncryptionKeyVaultPropertiesResponseOutput struct{ *pulumi.OutputState }

func (EncryptionKeyVaultPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionKeyVaultPropertiesResponse)(nil)).Elem()
}

func (o EncryptionKeyVaultPropertiesResponseOutput) ToEncryptionKeyVaultPropertiesResponseOutput() EncryptionKeyVaultPropertiesResponseOutput {
	return o
}

func (o EncryptionKeyVaultPropertiesResponseOutput) ToEncryptionKeyVaultPropertiesResponseOutputWithContext(ctx context.Context) EncryptionKeyVaultPropertiesResponseOutput {
	return o
}

func (o EncryptionKeyVaultPropertiesResponseOutput) ToEncryptionKeyVaultPropertiesResponsePtrOutput() EncryptionKeyVaultPropertiesResponsePtrOutput {
	return o.ToEncryptionKeyVaultPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o EncryptionKeyVaultPropertiesResponseOutput) ToEncryptionKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) EncryptionKeyVaultPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionKeyVaultPropertiesResponse) *EncryptionKeyVaultPropertiesResponse {
		return &v
	}).(EncryptionKeyVaultPropertiesResponsePtrOutput)
}

func (o EncryptionKeyVaultPropertiesResponseOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionKeyVaultPropertiesResponse) *string { return v.KeyName }).(pulumi.StringPtrOutput)
}

func (o EncryptionKeyVaultPropertiesResponseOutput) KeyState() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionKeyVaultPropertiesResponse) string { return v.KeyState }).(pulumi.StringOutput)
}

func (o EncryptionKeyVaultPropertiesResponseOutput) KeyVaultUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionKeyVaultPropertiesResponse) *string { return v.KeyVaultUrl }).(pulumi.StringPtrOutput)
}

func (o EncryptionKeyVaultPropertiesResponseOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionKeyVaultPropertiesResponse) *string { return v.KeyVersion }).(pulumi.StringPtrOutput)
}

func (o EncryptionKeyVaultPropertiesResponseOutput) VersionType() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionKeyVaultPropertiesResponse) string { return v.VersionType }).(pulumi.StringOutput)
}

type EncryptionKeyVaultPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionKeyVaultPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionKeyVaultPropertiesResponse)(nil)).Elem()
}

func (o EncryptionKeyVaultPropertiesResponsePtrOutput) ToEncryptionKeyVaultPropertiesResponsePtrOutput() EncryptionKeyVaultPropertiesResponsePtrOutput {
	return o
}

func (o EncryptionKeyVaultPropertiesResponsePtrOutput) ToEncryptionKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) EncryptionKeyVaultPropertiesResponsePtrOutput {
	return o
}

func (o EncryptionKeyVaultPropertiesResponsePtrOutput) Elem() EncryptionKeyVaultPropertiesResponseOutput {
	return o.ApplyT(func(v *EncryptionKeyVaultPropertiesResponse) EncryptionKeyVaultPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret EncryptionKeyVaultPropertiesResponse
		return ret
	}).(EncryptionKeyVaultPropertiesResponseOutput)
}

func (o EncryptionKeyVaultPropertiesResponsePtrOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionKeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyName
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionKeyVaultPropertiesResponsePtrOutput) KeyState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionKeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.KeyState
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionKeyVaultPropertiesResponsePtrOutput) KeyVaultUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionKeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultUrl
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionKeyVaultPropertiesResponsePtrOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionKeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyVersion
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionKeyVaultPropertiesResponsePtrOutput) VersionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionKeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.VersionType
	}).(pulumi.StringPtrOutput)
}

type EncryptionResponse struct {
	KeyVaultProperties *EncryptionKeyVaultPropertiesResponse `pulumi:"keyVaultProperties"`
	Status             *string                               `pulumi:"status"`
}





type EncryptionResponseInput interface {
	pulumi.Input

	ToEncryptionResponseOutput() EncryptionResponseOutput
	ToEncryptionResponseOutputWithContext(context.Context) EncryptionResponseOutput
}

type EncryptionResponseArgs struct {
	KeyVaultProperties EncryptionKeyVaultPropertiesResponsePtrInput `pulumi:"keyVaultProperties"`
	Status             pulumi.StringPtrInput                        `pulumi:"status"`
}

func (EncryptionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionResponse)(nil)).Elem()
}

func (i EncryptionResponseArgs) ToEncryptionResponseOutput() EncryptionResponseOutput {
	return i.ToEncryptionResponseOutputWithContext(context.Background())
}

func (i EncryptionResponseArgs) ToEncryptionResponseOutputWithContext(ctx context.Context) EncryptionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionResponseOutput)
}

func (i EncryptionResponseArgs) ToEncryptionResponsePtrOutput() EncryptionResponsePtrOutput {
	return i.ToEncryptionResponsePtrOutputWithContext(context.Background())
}

func (i EncryptionResponseArgs) ToEncryptionResponsePtrOutputWithContext(ctx context.Context) EncryptionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionResponseOutput).ToEncryptionResponsePtrOutputWithContext(ctx)
}









type EncryptionResponsePtrInput interface {
	pulumi.Input

	ToEncryptionResponsePtrOutput() EncryptionResponsePtrOutput
	ToEncryptionResponsePtrOutputWithContext(context.Context) EncryptionResponsePtrOutput
}

type encryptionResponsePtrType EncryptionResponseArgs

func EncryptionResponsePtr(v *EncryptionResponseArgs) EncryptionResponsePtrInput {
	return (*encryptionResponsePtrType)(v)
}

func (*encryptionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionResponse)(nil)).Elem()
}

func (i *encryptionResponsePtrType) ToEncryptionResponsePtrOutput() EncryptionResponsePtrOutput {
	return i.ToEncryptionResponsePtrOutputWithContext(context.Background())
}

func (i *encryptionResponsePtrType) ToEncryptionResponsePtrOutputWithContext(ctx context.Context) EncryptionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionResponsePtrOutput)
}

type EncryptionResponseOutput struct{ *pulumi.OutputState }

func (EncryptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionResponse)(nil)).Elem()
}

func (o EncryptionResponseOutput) ToEncryptionResponseOutput() EncryptionResponseOutput {
	return o
}

func (o EncryptionResponseOutput) ToEncryptionResponseOutputWithContext(ctx context.Context) EncryptionResponseOutput {
	return o
}

func (o EncryptionResponseOutput) ToEncryptionResponsePtrOutput() EncryptionResponsePtrOutput {
	return o.ToEncryptionResponsePtrOutputWithContext(context.Background())
}

func (o EncryptionResponseOutput) ToEncryptionResponsePtrOutputWithContext(ctx context.Context) EncryptionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionResponse) *EncryptionResponse {
		return &v
	}).(EncryptionResponsePtrOutput)
}

func (o EncryptionResponseOutput) KeyVaultProperties() EncryptionKeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v EncryptionResponse) *EncryptionKeyVaultPropertiesResponse { return v.KeyVaultProperties }).(EncryptionKeyVaultPropertiesResponsePtrOutput)
}

func (o EncryptionResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type EncryptionResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionResponse)(nil)).Elem()
}

func (o EncryptionResponsePtrOutput) ToEncryptionResponsePtrOutput() EncryptionResponsePtrOutput {
	return o
}

func (o EncryptionResponsePtrOutput) ToEncryptionResponsePtrOutputWithContext(ctx context.Context) EncryptionResponsePtrOutput {
	return o
}

func (o EncryptionResponsePtrOutput) Elem() EncryptionResponseOutput {
	return o.ApplyT(func(v *EncryptionResponse) EncryptionResponse {
		if v != nil {
			return *v
		}
		var ret EncryptionResponse
		return ret
	}).(EncryptionResponseOutput)
}

func (o EncryptionResponsePtrOutput) KeyVaultProperties() EncryptionKeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionResponse) *EncryptionKeyVaultPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.KeyVaultProperties
	}).(EncryptionKeyVaultPropertiesResponsePtrOutput)
}

func (o EncryptionResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type EndpointsResponse struct {
	HcxCloudManager string `pulumi:"hcxCloudManager"`
	NsxtManager     string `pulumi:"nsxtManager"`
	Vcsa            string `pulumi:"vcsa"`
}





type EndpointsResponseInput interface {
	pulumi.Input

	ToEndpointsResponseOutput() EndpointsResponseOutput
	ToEndpointsResponseOutputWithContext(context.Context) EndpointsResponseOutput
}

type EndpointsResponseArgs struct {
	HcxCloudManager pulumi.StringInput `pulumi:"hcxCloudManager"`
	NsxtManager     pulumi.StringInput `pulumi:"nsxtManager"`
	Vcsa            pulumi.StringInput `pulumi:"vcsa"`
}

func (EndpointsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointsResponse)(nil)).Elem()
}

func (i EndpointsResponseArgs) ToEndpointsResponseOutput() EndpointsResponseOutput {
	return i.ToEndpointsResponseOutputWithContext(context.Background())
}

func (i EndpointsResponseArgs) ToEndpointsResponseOutputWithContext(ctx context.Context) EndpointsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointsResponseOutput)
}

func (i EndpointsResponseArgs) ToEndpointsResponsePtrOutput() EndpointsResponsePtrOutput {
	return i.ToEndpointsResponsePtrOutputWithContext(context.Background())
}

func (i EndpointsResponseArgs) ToEndpointsResponsePtrOutputWithContext(ctx context.Context) EndpointsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointsResponseOutput).ToEndpointsResponsePtrOutputWithContext(ctx)
}









type EndpointsResponsePtrInput interface {
	pulumi.Input

	ToEndpointsResponsePtrOutput() EndpointsResponsePtrOutput
	ToEndpointsResponsePtrOutputWithContext(context.Context) EndpointsResponsePtrOutput
}

type endpointsResponsePtrType EndpointsResponseArgs

func EndpointsResponsePtr(v *EndpointsResponseArgs) EndpointsResponsePtrInput {
	return (*endpointsResponsePtrType)(v)
}

func (*endpointsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointsResponse)(nil)).Elem()
}

func (i *endpointsResponsePtrType) ToEndpointsResponsePtrOutput() EndpointsResponsePtrOutput {
	return i.ToEndpointsResponsePtrOutputWithContext(context.Background())
}

func (i *endpointsResponsePtrType) ToEndpointsResponsePtrOutputWithContext(ctx context.Context) EndpointsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointsResponsePtrOutput)
}

type EndpointsResponseOutput struct{ *pulumi.OutputState }

func (EndpointsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointsResponse)(nil)).Elem()
}

func (o EndpointsResponseOutput) ToEndpointsResponseOutput() EndpointsResponseOutput {
	return o
}

func (o EndpointsResponseOutput) ToEndpointsResponseOutputWithContext(ctx context.Context) EndpointsResponseOutput {
	return o
}

func (o EndpointsResponseOutput) ToEndpointsResponsePtrOutput() EndpointsResponsePtrOutput {
	return o.ToEndpointsResponsePtrOutputWithContext(context.Background())
}

func (o EndpointsResponseOutput) ToEndpointsResponsePtrOutputWithContext(ctx context.Context) EndpointsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EndpointsResponse) *EndpointsResponse {
		return &v
	}).(EndpointsResponsePtrOutput)
}

func (o EndpointsResponseOutput) HcxCloudManager() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointsResponse) string { return v.HcxCloudManager }).(pulumi.StringOutput)
}

func (o EndpointsResponseOutput) NsxtManager() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointsResponse) string { return v.NsxtManager }).(pulumi.StringOutput)
}

func (o EndpointsResponseOutput) Vcsa() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointsResponse) string { return v.Vcsa }).(pulumi.StringOutput)
}

type EndpointsResponsePtrOutput struct{ *pulumi.OutputState }

func (EndpointsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointsResponse)(nil)).Elem()
}

func (o EndpointsResponsePtrOutput) ToEndpointsResponsePtrOutput() EndpointsResponsePtrOutput {
	return o
}

func (o EndpointsResponsePtrOutput) ToEndpointsResponsePtrOutputWithContext(ctx context.Context) EndpointsResponsePtrOutput {
	return o
}

func (o EndpointsResponsePtrOutput) Elem() EndpointsResponseOutput {
	return o.ApplyT(func(v *EndpointsResponse) EndpointsResponse {
		if v != nil {
			return *v
		}
		var ret EndpointsResponse
		return ret
	}).(EndpointsResponseOutput)
}

func (o EndpointsResponsePtrOutput) HcxCloudManager() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.HcxCloudManager
	}).(pulumi.StringPtrOutput)
}

func (o EndpointsResponsePtrOutput) NsxtManager() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.NsxtManager
	}).(pulumi.StringPtrOutput)
}

func (o EndpointsResponsePtrOutput) Vcsa() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Vcsa
	}).(pulumi.StringPtrOutput)
}

type IdentitySource struct {
	Alias           *string `pulumi:"alias"`
	BaseGroupDN     *string `pulumi:"baseGroupDN"`
	BaseUserDN      *string `pulumi:"baseUserDN"`
	Domain          *string `pulumi:"domain"`
	Name            *string `pulumi:"name"`
	Password        *string `pulumi:"password"`
	PrimaryServer   *string `pulumi:"primaryServer"`
	SecondaryServer *string `pulumi:"secondaryServer"`
	Ssl             *string `pulumi:"ssl"`
	Username        *string `pulumi:"username"`
}





type IdentitySourceInput interface {
	pulumi.Input

	ToIdentitySourceOutput() IdentitySourceOutput
	ToIdentitySourceOutputWithContext(context.Context) IdentitySourceOutput
}

type IdentitySourceArgs struct {
	Alias           pulumi.StringPtrInput `pulumi:"alias"`
	BaseGroupDN     pulumi.StringPtrInput `pulumi:"baseGroupDN"`
	BaseUserDN      pulumi.StringPtrInput `pulumi:"baseUserDN"`
	Domain          pulumi.StringPtrInput `pulumi:"domain"`
	Name            pulumi.StringPtrInput `pulumi:"name"`
	Password        pulumi.StringPtrInput `pulumi:"password"`
	PrimaryServer   pulumi.StringPtrInput `pulumi:"primaryServer"`
	SecondaryServer pulumi.StringPtrInput `pulumi:"secondaryServer"`
	Ssl             pulumi.StringPtrInput `pulumi:"ssl"`
	Username        pulumi.StringPtrInput `pulumi:"username"`
}

func (IdentitySourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentitySource)(nil)).Elem()
}

func (i IdentitySourceArgs) ToIdentitySourceOutput() IdentitySourceOutput {
	return i.ToIdentitySourceOutputWithContext(context.Background())
}

func (i IdentitySourceArgs) ToIdentitySourceOutputWithContext(ctx context.Context) IdentitySourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentitySourceOutput)
}





type IdentitySourceArrayInput interface {
	pulumi.Input

	ToIdentitySourceArrayOutput() IdentitySourceArrayOutput
	ToIdentitySourceArrayOutputWithContext(context.Context) IdentitySourceArrayOutput
}

type IdentitySourceArray []IdentitySourceInput

func (IdentitySourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IdentitySource)(nil)).Elem()
}

func (i IdentitySourceArray) ToIdentitySourceArrayOutput() IdentitySourceArrayOutput {
	return i.ToIdentitySourceArrayOutputWithContext(context.Background())
}

func (i IdentitySourceArray) ToIdentitySourceArrayOutputWithContext(ctx context.Context) IdentitySourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentitySourceArrayOutput)
}

type IdentitySourceOutput struct{ *pulumi.OutputState }

func (IdentitySourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentitySource)(nil)).Elem()
}

func (o IdentitySourceOutput) ToIdentitySourceOutput() IdentitySourceOutput {
	return o
}

func (o IdentitySourceOutput) ToIdentitySourceOutputWithContext(ctx context.Context) IdentitySourceOutput {
	return o
}

func (o IdentitySourceOutput) Alias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySource) *string { return v.Alias }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceOutput) BaseGroupDN() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySource) *string { return v.BaseGroupDN }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceOutput) BaseUserDN() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySource) *string { return v.BaseUserDN }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySource) *string { return v.Domain }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySource) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySource) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceOutput) PrimaryServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySource) *string { return v.PrimaryServer }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceOutput) SecondaryServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySource) *string { return v.SecondaryServer }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceOutput) Ssl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySource) *string { return v.Ssl }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySource) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type IdentitySourceArrayOutput struct{ *pulumi.OutputState }

func (IdentitySourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IdentitySource)(nil)).Elem()
}

func (o IdentitySourceArrayOutput) ToIdentitySourceArrayOutput() IdentitySourceArrayOutput {
	return o
}

func (o IdentitySourceArrayOutput) ToIdentitySourceArrayOutputWithContext(ctx context.Context) IdentitySourceArrayOutput {
	return o
}

func (o IdentitySourceArrayOutput) Index(i pulumi.IntInput) IdentitySourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IdentitySource {
		return vs[0].([]IdentitySource)[vs[1].(int)]
	}).(IdentitySourceOutput)
}

type IdentitySourceResponse struct {
	Alias           *string `pulumi:"alias"`
	BaseGroupDN     *string `pulumi:"baseGroupDN"`
	BaseUserDN      *string `pulumi:"baseUserDN"`
	Domain          *string `pulumi:"domain"`
	Name            *string `pulumi:"name"`
	Password        *string `pulumi:"password"`
	PrimaryServer   *string `pulumi:"primaryServer"`
	SecondaryServer *string `pulumi:"secondaryServer"`
	Ssl             *string `pulumi:"ssl"`
	Username        *string `pulumi:"username"`
}





type IdentitySourceResponseInput interface {
	pulumi.Input

	ToIdentitySourceResponseOutput() IdentitySourceResponseOutput
	ToIdentitySourceResponseOutputWithContext(context.Context) IdentitySourceResponseOutput
}

type IdentitySourceResponseArgs struct {
	Alias           pulumi.StringPtrInput `pulumi:"alias"`
	BaseGroupDN     pulumi.StringPtrInput `pulumi:"baseGroupDN"`
	BaseUserDN      pulumi.StringPtrInput `pulumi:"baseUserDN"`
	Domain          pulumi.StringPtrInput `pulumi:"domain"`
	Name            pulumi.StringPtrInput `pulumi:"name"`
	Password        pulumi.StringPtrInput `pulumi:"password"`
	PrimaryServer   pulumi.StringPtrInput `pulumi:"primaryServer"`
	SecondaryServer pulumi.StringPtrInput `pulumi:"secondaryServer"`
	Ssl             pulumi.StringPtrInput `pulumi:"ssl"`
	Username        pulumi.StringPtrInput `pulumi:"username"`
}

func (IdentitySourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentitySourceResponse)(nil)).Elem()
}

func (i IdentitySourceResponseArgs) ToIdentitySourceResponseOutput() IdentitySourceResponseOutput {
	return i.ToIdentitySourceResponseOutputWithContext(context.Background())
}

func (i IdentitySourceResponseArgs) ToIdentitySourceResponseOutputWithContext(ctx context.Context) IdentitySourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentitySourceResponseOutput)
}





type IdentitySourceResponseArrayInput interface {
	pulumi.Input

	ToIdentitySourceResponseArrayOutput() IdentitySourceResponseArrayOutput
	ToIdentitySourceResponseArrayOutputWithContext(context.Context) IdentitySourceResponseArrayOutput
}

type IdentitySourceResponseArray []IdentitySourceResponseInput

func (IdentitySourceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IdentitySourceResponse)(nil)).Elem()
}

func (i IdentitySourceResponseArray) ToIdentitySourceResponseArrayOutput() IdentitySourceResponseArrayOutput {
	return i.ToIdentitySourceResponseArrayOutputWithContext(context.Background())
}

func (i IdentitySourceResponseArray) ToIdentitySourceResponseArrayOutputWithContext(ctx context.Context) IdentitySourceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentitySourceResponseArrayOutput)
}

type IdentitySourceResponseOutput struct{ *pulumi.OutputState }

func (IdentitySourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentitySourceResponse)(nil)).Elem()
}

func (o IdentitySourceResponseOutput) ToIdentitySourceResponseOutput() IdentitySourceResponseOutput {
	return o
}

func (o IdentitySourceResponseOutput) ToIdentitySourceResponseOutputWithContext(ctx context.Context) IdentitySourceResponseOutput {
	return o
}

func (o IdentitySourceResponseOutput) Alias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySourceResponse) *string { return v.Alias }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceResponseOutput) BaseGroupDN() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySourceResponse) *string { return v.BaseGroupDN }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceResponseOutput) BaseUserDN() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySourceResponse) *string { return v.BaseUserDN }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceResponseOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySourceResponse) *string { return v.Domain }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySourceResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceResponseOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySourceResponse) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceResponseOutput) PrimaryServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySourceResponse) *string { return v.PrimaryServer }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceResponseOutput) SecondaryServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySourceResponse) *string { return v.SecondaryServer }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceResponseOutput) Ssl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySourceResponse) *string { return v.Ssl }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceResponseOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySourceResponse) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type IdentitySourceResponseArrayOutput struct{ *pulumi.OutputState }

func (IdentitySourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IdentitySourceResponse)(nil)).Elem()
}

func (o IdentitySourceResponseArrayOutput) ToIdentitySourceResponseArrayOutput() IdentitySourceResponseArrayOutput {
	return o
}

func (o IdentitySourceResponseArrayOutput) ToIdentitySourceResponseArrayOutputWithContext(ctx context.Context) IdentitySourceResponseArrayOutput {
	return o
}

func (o IdentitySourceResponseArrayOutput) Index(i pulumi.IntInput) IdentitySourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IdentitySourceResponse {
		return vs[0].([]IdentitySourceResponse)[vs[1].(int)]
	}).(IdentitySourceResponseOutput)
}

type ManagementCluster struct {
	ClusterSize int      `pulumi:"clusterSize"`
	Hosts       []string `pulumi:"hosts"`
}





type ManagementClusterInput interface {
	pulumi.Input

	ToManagementClusterOutput() ManagementClusterOutput
	ToManagementClusterOutputWithContext(context.Context) ManagementClusterOutput
}

type ManagementClusterArgs struct {
	ClusterSize pulumi.IntInput         `pulumi:"clusterSize"`
	Hosts       pulumi.StringArrayInput `pulumi:"hosts"`
}

func (ManagementClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementCluster)(nil)).Elem()
}

func (i ManagementClusterArgs) ToManagementClusterOutput() ManagementClusterOutput {
	return i.ToManagementClusterOutputWithContext(context.Background())
}

func (i ManagementClusterArgs) ToManagementClusterOutputWithContext(ctx context.Context) ManagementClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementClusterOutput)
}

func (i ManagementClusterArgs) ToManagementClusterPtrOutput() ManagementClusterPtrOutput {
	return i.ToManagementClusterPtrOutputWithContext(context.Background())
}

func (i ManagementClusterArgs) ToManagementClusterPtrOutputWithContext(ctx context.Context) ManagementClusterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementClusterOutput).ToManagementClusterPtrOutputWithContext(ctx)
}









type ManagementClusterPtrInput interface {
	pulumi.Input

	ToManagementClusterPtrOutput() ManagementClusterPtrOutput
	ToManagementClusterPtrOutputWithContext(context.Context) ManagementClusterPtrOutput
}

type managementClusterPtrType ManagementClusterArgs

func ManagementClusterPtr(v *ManagementClusterArgs) ManagementClusterPtrInput {
	return (*managementClusterPtrType)(v)
}

func (*managementClusterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementCluster)(nil)).Elem()
}

func (i *managementClusterPtrType) ToManagementClusterPtrOutput() ManagementClusterPtrOutput {
	return i.ToManagementClusterPtrOutputWithContext(context.Background())
}

func (i *managementClusterPtrType) ToManagementClusterPtrOutputWithContext(ctx context.Context) ManagementClusterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementClusterPtrOutput)
}

type ManagementClusterOutput struct{ *pulumi.OutputState }

func (ManagementClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementCluster)(nil)).Elem()
}

func (o ManagementClusterOutput) ToManagementClusterOutput() ManagementClusterOutput {
	return o
}

func (o ManagementClusterOutput) ToManagementClusterOutputWithContext(ctx context.Context) ManagementClusterOutput {
	return o
}

func (o ManagementClusterOutput) ToManagementClusterPtrOutput() ManagementClusterPtrOutput {
	return o.ToManagementClusterPtrOutputWithContext(context.Background())
}

func (o ManagementClusterOutput) ToManagementClusterPtrOutputWithContext(ctx context.Context) ManagementClusterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagementCluster) *ManagementCluster {
		return &v
	}).(ManagementClusterPtrOutput)
}

func (o ManagementClusterOutput) ClusterSize() pulumi.IntOutput {
	return o.ApplyT(func(v ManagementCluster) int { return v.ClusterSize }).(pulumi.IntOutput)
}

func (o ManagementClusterOutput) Hosts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ManagementCluster) []string { return v.Hosts }).(pulumi.StringArrayOutput)
}

type ManagementClusterPtrOutput struct{ *pulumi.OutputState }

func (ManagementClusterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementCluster)(nil)).Elem()
}

func (o ManagementClusterPtrOutput) ToManagementClusterPtrOutput() ManagementClusterPtrOutput {
	return o
}

func (o ManagementClusterPtrOutput) ToManagementClusterPtrOutputWithContext(ctx context.Context) ManagementClusterPtrOutput {
	return o
}

func (o ManagementClusterPtrOutput) Elem() ManagementClusterOutput {
	return o.ApplyT(func(v *ManagementCluster) ManagementCluster {
		if v != nil {
			return *v
		}
		var ret ManagementCluster
		return ret
	}).(ManagementClusterOutput)
}

func (o ManagementClusterPtrOutput) ClusterSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ManagementCluster) *int {
		if v == nil {
			return nil
		}
		return &v.ClusterSize
	}).(pulumi.IntPtrOutput)
}

func (o ManagementClusterPtrOutput) Hosts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ManagementCluster) []string {
		if v == nil {
			return nil
		}
		return v.Hosts
	}).(pulumi.StringArrayOutput)
}

type ManagementClusterResponse struct {
	ClusterId         int      `pulumi:"clusterId"`
	ClusterSize       int      `pulumi:"clusterSize"`
	Hosts             []string `pulumi:"hosts"`
	ProvisioningState string   `pulumi:"provisioningState"`
}





type ManagementClusterResponseInput interface {
	pulumi.Input

	ToManagementClusterResponseOutput() ManagementClusterResponseOutput
	ToManagementClusterResponseOutputWithContext(context.Context) ManagementClusterResponseOutput
}

type ManagementClusterResponseArgs struct {
	ClusterId         pulumi.IntInput         `pulumi:"clusterId"`
	ClusterSize       pulumi.IntInput         `pulumi:"clusterSize"`
	Hosts             pulumi.StringArrayInput `pulumi:"hosts"`
	ProvisioningState pulumi.StringInput      `pulumi:"provisioningState"`
}

func (ManagementClusterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementClusterResponse)(nil)).Elem()
}

func (i ManagementClusterResponseArgs) ToManagementClusterResponseOutput() ManagementClusterResponseOutput {
	return i.ToManagementClusterResponseOutputWithContext(context.Background())
}

func (i ManagementClusterResponseArgs) ToManagementClusterResponseOutputWithContext(ctx context.Context) ManagementClusterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementClusterResponseOutput)
}

func (i ManagementClusterResponseArgs) ToManagementClusterResponsePtrOutput() ManagementClusterResponsePtrOutput {
	return i.ToManagementClusterResponsePtrOutputWithContext(context.Background())
}

func (i ManagementClusterResponseArgs) ToManagementClusterResponsePtrOutputWithContext(ctx context.Context) ManagementClusterResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementClusterResponseOutput).ToManagementClusterResponsePtrOutputWithContext(ctx)
}









type ManagementClusterResponsePtrInput interface {
	pulumi.Input

	ToManagementClusterResponsePtrOutput() ManagementClusterResponsePtrOutput
	ToManagementClusterResponsePtrOutputWithContext(context.Context) ManagementClusterResponsePtrOutput
}

type managementClusterResponsePtrType ManagementClusterResponseArgs

func ManagementClusterResponsePtr(v *ManagementClusterResponseArgs) ManagementClusterResponsePtrInput {
	return (*managementClusterResponsePtrType)(v)
}

func (*managementClusterResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementClusterResponse)(nil)).Elem()
}

func (i *managementClusterResponsePtrType) ToManagementClusterResponsePtrOutput() ManagementClusterResponsePtrOutput {
	return i.ToManagementClusterResponsePtrOutputWithContext(context.Background())
}

func (i *managementClusterResponsePtrType) ToManagementClusterResponsePtrOutputWithContext(ctx context.Context) ManagementClusterResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementClusterResponsePtrOutput)
}

type ManagementClusterResponseOutput struct{ *pulumi.OutputState }

func (ManagementClusterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementClusterResponse)(nil)).Elem()
}

func (o ManagementClusterResponseOutput) ToManagementClusterResponseOutput() ManagementClusterResponseOutput {
	return o
}

func (o ManagementClusterResponseOutput) ToManagementClusterResponseOutputWithContext(ctx context.Context) ManagementClusterResponseOutput {
	return o
}

func (o ManagementClusterResponseOutput) ToManagementClusterResponsePtrOutput() ManagementClusterResponsePtrOutput {
	return o.ToManagementClusterResponsePtrOutputWithContext(context.Background())
}

func (o ManagementClusterResponseOutput) ToManagementClusterResponsePtrOutputWithContext(ctx context.Context) ManagementClusterResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagementClusterResponse) *ManagementClusterResponse {
		return &v
	}).(ManagementClusterResponsePtrOutput)
}

func (o ManagementClusterResponseOutput) ClusterId() pulumi.IntOutput {
	return o.ApplyT(func(v ManagementClusterResponse) int { return v.ClusterId }).(pulumi.IntOutput)
}

func (o ManagementClusterResponseOutput) ClusterSize() pulumi.IntOutput {
	return o.ApplyT(func(v ManagementClusterResponse) int { return v.ClusterSize }).(pulumi.IntOutput)
}

func (o ManagementClusterResponseOutput) Hosts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ManagementClusterResponse) []string { return v.Hosts }).(pulumi.StringArrayOutput)
}

func (o ManagementClusterResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ManagementClusterResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type ManagementClusterResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagementClusterResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementClusterResponse)(nil)).Elem()
}

func (o ManagementClusterResponsePtrOutput) ToManagementClusterResponsePtrOutput() ManagementClusterResponsePtrOutput {
	return o
}

func (o ManagementClusterResponsePtrOutput) ToManagementClusterResponsePtrOutputWithContext(ctx context.Context) ManagementClusterResponsePtrOutput {
	return o
}

func (o ManagementClusterResponsePtrOutput) Elem() ManagementClusterResponseOutput {
	return o.ApplyT(func(v *ManagementClusterResponse) ManagementClusterResponse {
		if v != nil {
			return *v
		}
		var ret ManagementClusterResponse
		return ret
	}).(ManagementClusterResponseOutput)
}

func (o ManagementClusterResponsePtrOutput) ClusterId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ManagementClusterResponse) *int {
		if v == nil {
			return nil
		}
		return &v.ClusterId
	}).(pulumi.IntPtrOutput)
}

func (o ManagementClusterResponsePtrOutput) ClusterSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ManagementClusterResponse) *int {
		if v == nil {
			return nil
		}
		return &v.ClusterSize
	}).(pulumi.IntPtrOutput)
}

func (o ManagementClusterResponsePtrOutput) Hosts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ManagementClusterResponse) []string {
		if v == nil {
			return nil
		}
		return v.Hosts
	}).(pulumi.StringArrayOutput)
}

func (o ManagementClusterResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagementClusterResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

type NetAppVolume struct {
	Id string `pulumi:"id"`
}





type NetAppVolumeInput interface {
	pulumi.Input

	ToNetAppVolumeOutput() NetAppVolumeOutput
	ToNetAppVolumeOutputWithContext(context.Context) NetAppVolumeOutput
}

type NetAppVolumeArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (NetAppVolumeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetAppVolume)(nil)).Elem()
}

func (i NetAppVolumeArgs) ToNetAppVolumeOutput() NetAppVolumeOutput {
	return i.ToNetAppVolumeOutputWithContext(context.Background())
}

func (i NetAppVolumeArgs) ToNetAppVolumeOutputWithContext(ctx context.Context) NetAppVolumeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetAppVolumeOutput)
}

func (i NetAppVolumeArgs) ToNetAppVolumePtrOutput() NetAppVolumePtrOutput {
	return i.ToNetAppVolumePtrOutputWithContext(context.Background())
}

func (i NetAppVolumeArgs) ToNetAppVolumePtrOutputWithContext(ctx context.Context) NetAppVolumePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetAppVolumeOutput).ToNetAppVolumePtrOutputWithContext(ctx)
}









type NetAppVolumePtrInput interface {
	pulumi.Input

	ToNetAppVolumePtrOutput() NetAppVolumePtrOutput
	ToNetAppVolumePtrOutputWithContext(context.Context) NetAppVolumePtrOutput
}

type netAppVolumePtrType NetAppVolumeArgs

func NetAppVolumePtr(v *NetAppVolumeArgs) NetAppVolumePtrInput {
	return (*netAppVolumePtrType)(v)
}

func (*netAppVolumePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetAppVolume)(nil)).Elem()
}

func (i *netAppVolumePtrType) ToNetAppVolumePtrOutput() NetAppVolumePtrOutput {
	return i.ToNetAppVolumePtrOutputWithContext(context.Background())
}

func (i *netAppVolumePtrType) ToNetAppVolumePtrOutputWithContext(ctx context.Context) NetAppVolumePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetAppVolumePtrOutput)
}

type NetAppVolumeOutput struct{ *pulumi.OutputState }

func (NetAppVolumeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetAppVolume)(nil)).Elem()
}

func (o NetAppVolumeOutput) ToNetAppVolumeOutput() NetAppVolumeOutput {
	return o
}

func (o NetAppVolumeOutput) ToNetAppVolumeOutputWithContext(ctx context.Context) NetAppVolumeOutput {
	return o
}

func (o NetAppVolumeOutput) ToNetAppVolumePtrOutput() NetAppVolumePtrOutput {
	return o.ToNetAppVolumePtrOutputWithContext(context.Background())
}

func (o NetAppVolumeOutput) ToNetAppVolumePtrOutputWithContext(ctx context.Context) NetAppVolumePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetAppVolume) *NetAppVolume {
		return &v
	}).(NetAppVolumePtrOutput)
}

func (o NetAppVolumeOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v NetAppVolume) string { return v.Id }).(pulumi.StringOutput)
}

type NetAppVolumePtrOutput struct{ *pulumi.OutputState }

func (NetAppVolumePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetAppVolume)(nil)).Elem()
}

func (o NetAppVolumePtrOutput) ToNetAppVolumePtrOutput() NetAppVolumePtrOutput {
	return o
}

func (o NetAppVolumePtrOutput) ToNetAppVolumePtrOutputWithContext(ctx context.Context) NetAppVolumePtrOutput {
	return o
}

func (o NetAppVolumePtrOutput) Elem() NetAppVolumeOutput {
	return o.ApplyT(func(v *NetAppVolume) NetAppVolume {
		if v != nil {
			return *v
		}
		var ret NetAppVolume
		return ret
	}).(NetAppVolumeOutput)
}

func (o NetAppVolumePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetAppVolume) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type NetAppVolumeResponse struct {
	Id string `pulumi:"id"`
}





type NetAppVolumeResponseInput interface {
	pulumi.Input

	ToNetAppVolumeResponseOutput() NetAppVolumeResponseOutput
	ToNetAppVolumeResponseOutputWithContext(context.Context) NetAppVolumeResponseOutput
}

type NetAppVolumeResponseArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (NetAppVolumeResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetAppVolumeResponse)(nil)).Elem()
}

func (i NetAppVolumeResponseArgs) ToNetAppVolumeResponseOutput() NetAppVolumeResponseOutput {
	return i.ToNetAppVolumeResponseOutputWithContext(context.Background())
}

func (i NetAppVolumeResponseArgs) ToNetAppVolumeResponseOutputWithContext(ctx context.Context) NetAppVolumeResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetAppVolumeResponseOutput)
}

func (i NetAppVolumeResponseArgs) ToNetAppVolumeResponsePtrOutput() NetAppVolumeResponsePtrOutput {
	return i.ToNetAppVolumeResponsePtrOutputWithContext(context.Background())
}

func (i NetAppVolumeResponseArgs) ToNetAppVolumeResponsePtrOutputWithContext(ctx context.Context) NetAppVolumeResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetAppVolumeResponseOutput).ToNetAppVolumeResponsePtrOutputWithContext(ctx)
}









type NetAppVolumeResponsePtrInput interface {
	pulumi.Input

	ToNetAppVolumeResponsePtrOutput() NetAppVolumeResponsePtrOutput
	ToNetAppVolumeResponsePtrOutputWithContext(context.Context) NetAppVolumeResponsePtrOutput
}

type netAppVolumeResponsePtrType NetAppVolumeResponseArgs

func NetAppVolumeResponsePtr(v *NetAppVolumeResponseArgs) NetAppVolumeResponsePtrInput {
	return (*netAppVolumeResponsePtrType)(v)
}

func (*netAppVolumeResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetAppVolumeResponse)(nil)).Elem()
}

func (i *netAppVolumeResponsePtrType) ToNetAppVolumeResponsePtrOutput() NetAppVolumeResponsePtrOutput {
	return i.ToNetAppVolumeResponsePtrOutputWithContext(context.Background())
}

func (i *netAppVolumeResponsePtrType) ToNetAppVolumeResponsePtrOutputWithContext(ctx context.Context) NetAppVolumeResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetAppVolumeResponsePtrOutput)
}

type NetAppVolumeResponseOutput struct{ *pulumi.OutputState }

func (NetAppVolumeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetAppVolumeResponse)(nil)).Elem()
}

func (o NetAppVolumeResponseOutput) ToNetAppVolumeResponseOutput() NetAppVolumeResponseOutput {
	return o
}

func (o NetAppVolumeResponseOutput) ToNetAppVolumeResponseOutputWithContext(ctx context.Context) NetAppVolumeResponseOutput {
	return o
}

func (o NetAppVolumeResponseOutput) ToNetAppVolumeResponsePtrOutput() NetAppVolumeResponsePtrOutput {
	return o.ToNetAppVolumeResponsePtrOutputWithContext(context.Background())
}

func (o NetAppVolumeResponseOutput) ToNetAppVolumeResponsePtrOutputWithContext(ctx context.Context) NetAppVolumeResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetAppVolumeResponse) *NetAppVolumeResponse {
		return &v
	}).(NetAppVolumeResponsePtrOutput)
}

func (o NetAppVolumeResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v NetAppVolumeResponse) string { return v.Id }).(pulumi.StringOutput)
}

type NetAppVolumeResponsePtrOutput struct{ *pulumi.OutputState }

func (NetAppVolumeResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetAppVolumeResponse)(nil)).Elem()
}

func (o NetAppVolumeResponsePtrOutput) ToNetAppVolumeResponsePtrOutput() NetAppVolumeResponsePtrOutput {
	return o
}

func (o NetAppVolumeResponsePtrOutput) ToNetAppVolumeResponsePtrOutputWithContext(ctx context.Context) NetAppVolumeResponsePtrOutput {
	return o
}

func (o NetAppVolumeResponsePtrOutput) Elem() NetAppVolumeResponseOutput {
	return o.ApplyT(func(v *NetAppVolumeResponse) NetAppVolumeResponse {
		if v != nil {
			return *v
		}
		var ret NetAppVolumeResponse
		return ret
	}).(NetAppVolumeResponseOutput)
}

func (o NetAppVolumeResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetAppVolumeResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type PSCredentialExecutionParameter struct {
	Name     string  `pulumi:"name"`
	Password *string `pulumi:"password"`
	Type     string  `pulumi:"type"`
	Username *string `pulumi:"username"`
}





type PSCredentialExecutionParameterInput interface {
	pulumi.Input

	ToPSCredentialExecutionParameterOutput() PSCredentialExecutionParameterOutput
	ToPSCredentialExecutionParameterOutputWithContext(context.Context) PSCredentialExecutionParameterOutput
}

type PSCredentialExecutionParameterArgs struct {
	Name     pulumi.StringInput    `pulumi:"name"`
	Password pulumi.StringPtrInput `pulumi:"password"`
	Type     pulumi.StringInput    `pulumi:"type"`
	Username pulumi.StringPtrInput `pulumi:"username"`
}

func (PSCredentialExecutionParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PSCredentialExecutionParameter)(nil)).Elem()
}

func (i PSCredentialExecutionParameterArgs) ToPSCredentialExecutionParameterOutput() PSCredentialExecutionParameterOutput {
	return i.ToPSCredentialExecutionParameterOutputWithContext(context.Background())
}

func (i PSCredentialExecutionParameterArgs) ToPSCredentialExecutionParameterOutputWithContext(ctx context.Context) PSCredentialExecutionParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PSCredentialExecutionParameterOutput)
}

type PSCredentialExecutionParameterOutput struct{ *pulumi.OutputState }

func (PSCredentialExecutionParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PSCredentialExecutionParameter)(nil)).Elem()
}

func (o PSCredentialExecutionParameterOutput) ToPSCredentialExecutionParameterOutput() PSCredentialExecutionParameterOutput {
	return o
}

func (o PSCredentialExecutionParameterOutput) ToPSCredentialExecutionParameterOutputWithContext(ctx context.Context) PSCredentialExecutionParameterOutput {
	return o
}

func (o PSCredentialExecutionParameterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PSCredentialExecutionParameter) string { return v.Name }).(pulumi.StringOutput)
}

func (o PSCredentialExecutionParameterOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PSCredentialExecutionParameter) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o PSCredentialExecutionParameterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PSCredentialExecutionParameter) string { return v.Type }).(pulumi.StringOutput)
}

func (o PSCredentialExecutionParameterOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PSCredentialExecutionParameter) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type PSCredentialExecutionParameterResponse struct {
	Name     string  `pulumi:"name"`
	Password *string `pulumi:"password"`
	Type     string  `pulumi:"type"`
	Username *string `pulumi:"username"`
}





type PSCredentialExecutionParameterResponseInput interface {
	pulumi.Input

	ToPSCredentialExecutionParameterResponseOutput() PSCredentialExecutionParameterResponseOutput
	ToPSCredentialExecutionParameterResponseOutputWithContext(context.Context) PSCredentialExecutionParameterResponseOutput
}

type PSCredentialExecutionParameterResponseArgs struct {
	Name     pulumi.StringInput    `pulumi:"name"`
	Password pulumi.StringPtrInput `pulumi:"password"`
	Type     pulumi.StringInput    `pulumi:"type"`
	Username pulumi.StringPtrInput `pulumi:"username"`
}

func (PSCredentialExecutionParameterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PSCredentialExecutionParameterResponse)(nil)).Elem()
}

func (i PSCredentialExecutionParameterResponseArgs) ToPSCredentialExecutionParameterResponseOutput() PSCredentialExecutionParameterResponseOutput {
	return i.ToPSCredentialExecutionParameterResponseOutputWithContext(context.Background())
}

func (i PSCredentialExecutionParameterResponseArgs) ToPSCredentialExecutionParameterResponseOutputWithContext(ctx context.Context) PSCredentialExecutionParameterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PSCredentialExecutionParameterResponseOutput)
}

type PSCredentialExecutionParameterResponseOutput struct{ *pulumi.OutputState }

func (PSCredentialExecutionParameterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PSCredentialExecutionParameterResponse)(nil)).Elem()
}

func (o PSCredentialExecutionParameterResponseOutput) ToPSCredentialExecutionParameterResponseOutput() PSCredentialExecutionParameterResponseOutput {
	return o
}

func (o PSCredentialExecutionParameterResponseOutput) ToPSCredentialExecutionParameterResponseOutputWithContext(ctx context.Context) PSCredentialExecutionParameterResponseOutput {
	return o
}

func (o PSCredentialExecutionParameterResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PSCredentialExecutionParameterResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PSCredentialExecutionParameterResponseOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PSCredentialExecutionParameterResponse) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o PSCredentialExecutionParameterResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PSCredentialExecutionParameterResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o PSCredentialExecutionParameterResponseOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PSCredentialExecutionParameterResponse) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type PrivateCloudIdentity struct {
	Type *string `pulumi:"type"`
}





type PrivateCloudIdentityInput interface {
	pulumi.Input

	ToPrivateCloudIdentityOutput() PrivateCloudIdentityOutput
	ToPrivateCloudIdentityOutputWithContext(context.Context) PrivateCloudIdentityOutput
}

type PrivateCloudIdentityArgs struct {
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (PrivateCloudIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateCloudIdentity)(nil)).Elem()
}

func (i PrivateCloudIdentityArgs) ToPrivateCloudIdentityOutput() PrivateCloudIdentityOutput {
	return i.ToPrivateCloudIdentityOutputWithContext(context.Background())
}

func (i PrivateCloudIdentityArgs) ToPrivateCloudIdentityOutputWithContext(ctx context.Context) PrivateCloudIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateCloudIdentityOutput)
}

func (i PrivateCloudIdentityArgs) ToPrivateCloudIdentityPtrOutput() PrivateCloudIdentityPtrOutput {
	return i.ToPrivateCloudIdentityPtrOutputWithContext(context.Background())
}

func (i PrivateCloudIdentityArgs) ToPrivateCloudIdentityPtrOutputWithContext(ctx context.Context) PrivateCloudIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateCloudIdentityOutput).ToPrivateCloudIdentityPtrOutputWithContext(ctx)
}









type PrivateCloudIdentityPtrInput interface {
	pulumi.Input

	ToPrivateCloudIdentityPtrOutput() PrivateCloudIdentityPtrOutput
	ToPrivateCloudIdentityPtrOutputWithContext(context.Context) PrivateCloudIdentityPtrOutput
}

type privateCloudIdentityPtrType PrivateCloudIdentityArgs

func PrivateCloudIdentityPtr(v *PrivateCloudIdentityArgs) PrivateCloudIdentityPtrInput {
	return (*privateCloudIdentityPtrType)(v)
}

func (*privateCloudIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateCloudIdentity)(nil)).Elem()
}

func (i *privateCloudIdentityPtrType) ToPrivateCloudIdentityPtrOutput() PrivateCloudIdentityPtrOutput {
	return i.ToPrivateCloudIdentityPtrOutputWithContext(context.Background())
}

func (i *privateCloudIdentityPtrType) ToPrivateCloudIdentityPtrOutputWithContext(ctx context.Context) PrivateCloudIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateCloudIdentityPtrOutput)
}

type PrivateCloudIdentityOutput struct{ *pulumi.OutputState }

func (PrivateCloudIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateCloudIdentity)(nil)).Elem()
}

func (o PrivateCloudIdentityOutput) ToPrivateCloudIdentityOutput() PrivateCloudIdentityOutput {
	return o
}

func (o PrivateCloudIdentityOutput) ToPrivateCloudIdentityOutputWithContext(ctx context.Context) PrivateCloudIdentityOutput {
	return o
}

func (o PrivateCloudIdentityOutput) ToPrivateCloudIdentityPtrOutput() PrivateCloudIdentityPtrOutput {
	return o.ToPrivateCloudIdentityPtrOutputWithContext(context.Background())
}

func (o PrivateCloudIdentityOutput) ToPrivateCloudIdentityPtrOutputWithContext(ctx context.Context) PrivateCloudIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateCloudIdentity) *PrivateCloudIdentity {
		return &v
	}).(PrivateCloudIdentityPtrOutput)
}

func (o PrivateCloudIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateCloudIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type PrivateCloudIdentityPtrOutput struct{ *pulumi.OutputState }

func (PrivateCloudIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateCloudIdentity)(nil)).Elem()
}

func (o PrivateCloudIdentityPtrOutput) ToPrivateCloudIdentityPtrOutput() PrivateCloudIdentityPtrOutput {
	return o
}

func (o PrivateCloudIdentityPtrOutput) ToPrivateCloudIdentityPtrOutputWithContext(ctx context.Context) PrivateCloudIdentityPtrOutput {
	return o
}

func (o PrivateCloudIdentityPtrOutput) Elem() PrivateCloudIdentityOutput {
	return o.ApplyT(func(v *PrivateCloudIdentity) PrivateCloudIdentity {
		if v != nil {
			return *v
		}
		var ret PrivateCloudIdentity
		return ret
	}).(PrivateCloudIdentityOutput)
}

func (o PrivateCloudIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateCloudIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type PrivateCloudIdentityResponse struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}





type PrivateCloudIdentityResponseInput interface {
	pulumi.Input

	ToPrivateCloudIdentityResponseOutput() PrivateCloudIdentityResponseOutput
	ToPrivateCloudIdentityResponseOutputWithContext(context.Context) PrivateCloudIdentityResponseOutput
}

type PrivateCloudIdentityResponseArgs struct {
	PrincipalId pulumi.StringInput    `pulumi:"principalId"`
	TenantId    pulumi.StringInput    `pulumi:"tenantId"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
}

func (PrivateCloudIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateCloudIdentityResponse)(nil)).Elem()
}

func (i PrivateCloudIdentityResponseArgs) ToPrivateCloudIdentityResponseOutput() PrivateCloudIdentityResponseOutput {
	return i.ToPrivateCloudIdentityResponseOutputWithContext(context.Background())
}

func (i PrivateCloudIdentityResponseArgs) ToPrivateCloudIdentityResponseOutputWithContext(ctx context.Context) PrivateCloudIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateCloudIdentityResponseOutput)
}

func (i PrivateCloudIdentityResponseArgs) ToPrivateCloudIdentityResponsePtrOutput() PrivateCloudIdentityResponsePtrOutput {
	return i.ToPrivateCloudIdentityResponsePtrOutputWithContext(context.Background())
}

func (i PrivateCloudIdentityResponseArgs) ToPrivateCloudIdentityResponsePtrOutputWithContext(ctx context.Context) PrivateCloudIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateCloudIdentityResponseOutput).ToPrivateCloudIdentityResponsePtrOutputWithContext(ctx)
}









type PrivateCloudIdentityResponsePtrInput interface {
	pulumi.Input

	ToPrivateCloudIdentityResponsePtrOutput() PrivateCloudIdentityResponsePtrOutput
	ToPrivateCloudIdentityResponsePtrOutputWithContext(context.Context) PrivateCloudIdentityResponsePtrOutput
}

type privateCloudIdentityResponsePtrType PrivateCloudIdentityResponseArgs

func PrivateCloudIdentityResponsePtr(v *PrivateCloudIdentityResponseArgs) PrivateCloudIdentityResponsePtrInput {
	return (*privateCloudIdentityResponsePtrType)(v)
}

func (*privateCloudIdentityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateCloudIdentityResponse)(nil)).Elem()
}

func (i *privateCloudIdentityResponsePtrType) ToPrivateCloudIdentityResponsePtrOutput() PrivateCloudIdentityResponsePtrOutput {
	return i.ToPrivateCloudIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *privateCloudIdentityResponsePtrType) ToPrivateCloudIdentityResponsePtrOutputWithContext(ctx context.Context) PrivateCloudIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateCloudIdentityResponsePtrOutput)
}

type PrivateCloudIdentityResponseOutput struct{ *pulumi.OutputState }

func (PrivateCloudIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateCloudIdentityResponse)(nil)).Elem()
}

func (o PrivateCloudIdentityResponseOutput) ToPrivateCloudIdentityResponseOutput() PrivateCloudIdentityResponseOutput {
	return o
}

func (o PrivateCloudIdentityResponseOutput) ToPrivateCloudIdentityResponseOutputWithContext(ctx context.Context) PrivateCloudIdentityResponseOutput {
	return o
}

func (o PrivateCloudIdentityResponseOutput) ToPrivateCloudIdentityResponsePtrOutput() PrivateCloudIdentityResponsePtrOutput {
	return o.ToPrivateCloudIdentityResponsePtrOutputWithContext(context.Background())
}

func (o PrivateCloudIdentityResponseOutput) ToPrivateCloudIdentityResponsePtrOutputWithContext(ctx context.Context) PrivateCloudIdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateCloudIdentityResponse) *PrivateCloudIdentityResponse {
		return &v
	}).(PrivateCloudIdentityResponsePtrOutput)
}

func (o PrivateCloudIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateCloudIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o PrivateCloudIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateCloudIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o PrivateCloudIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateCloudIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type PrivateCloudIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateCloudIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateCloudIdentityResponse)(nil)).Elem()
}

func (o PrivateCloudIdentityResponsePtrOutput) ToPrivateCloudIdentityResponsePtrOutput() PrivateCloudIdentityResponsePtrOutput {
	return o
}

func (o PrivateCloudIdentityResponsePtrOutput) ToPrivateCloudIdentityResponsePtrOutputWithContext(ctx context.Context) PrivateCloudIdentityResponsePtrOutput {
	return o
}

func (o PrivateCloudIdentityResponsePtrOutput) Elem() PrivateCloudIdentityResponseOutput {
	return o.ApplyT(func(v *PrivateCloudIdentityResponse) PrivateCloudIdentityResponse {
		if v != nil {
			return *v
		}
		var ret PrivateCloudIdentityResponse
		return ret
	}).(PrivateCloudIdentityResponseOutput)
}

func (o PrivateCloudIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateCloudIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o PrivateCloudIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateCloudIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o PrivateCloudIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateCloudIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ScriptSecureStringExecutionParameter struct {
	Name        string  `pulumi:"name"`
	SecureValue *string `pulumi:"secureValue"`
	Type        string  `pulumi:"type"`
}





type ScriptSecureStringExecutionParameterInput interface {
	pulumi.Input

	ToScriptSecureStringExecutionParameterOutput() ScriptSecureStringExecutionParameterOutput
	ToScriptSecureStringExecutionParameterOutputWithContext(context.Context) ScriptSecureStringExecutionParameterOutput
}

type ScriptSecureStringExecutionParameterArgs struct {
	Name        pulumi.StringInput    `pulumi:"name"`
	SecureValue pulumi.StringPtrInput `pulumi:"secureValue"`
	Type        pulumi.StringInput    `pulumi:"type"`
}

func (ScriptSecureStringExecutionParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScriptSecureStringExecutionParameter)(nil)).Elem()
}

func (i ScriptSecureStringExecutionParameterArgs) ToScriptSecureStringExecutionParameterOutput() ScriptSecureStringExecutionParameterOutput {
	return i.ToScriptSecureStringExecutionParameterOutputWithContext(context.Background())
}

func (i ScriptSecureStringExecutionParameterArgs) ToScriptSecureStringExecutionParameterOutputWithContext(ctx context.Context) ScriptSecureStringExecutionParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScriptSecureStringExecutionParameterOutput)
}

type ScriptSecureStringExecutionParameterOutput struct{ *pulumi.OutputState }

func (ScriptSecureStringExecutionParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScriptSecureStringExecutionParameter)(nil)).Elem()
}

func (o ScriptSecureStringExecutionParameterOutput) ToScriptSecureStringExecutionParameterOutput() ScriptSecureStringExecutionParameterOutput {
	return o
}

func (o ScriptSecureStringExecutionParameterOutput) ToScriptSecureStringExecutionParameterOutputWithContext(ctx context.Context) ScriptSecureStringExecutionParameterOutput {
	return o
}

func (o ScriptSecureStringExecutionParameterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ScriptSecureStringExecutionParameter) string { return v.Name }).(pulumi.StringOutput)
}

func (o ScriptSecureStringExecutionParameterOutput) SecureValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScriptSecureStringExecutionParameter) *string { return v.SecureValue }).(pulumi.StringPtrOutput)
}

func (o ScriptSecureStringExecutionParameterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ScriptSecureStringExecutionParameter) string { return v.Type }).(pulumi.StringOutput)
}

type ScriptSecureStringExecutionParameterResponse struct {
	Name        string  `pulumi:"name"`
	SecureValue *string `pulumi:"secureValue"`
	Type        string  `pulumi:"type"`
}





type ScriptSecureStringExecutionParameterResponseInput interface {
	pulumi.Input

	ToScriptSecureStringExecutionParameterResponseOutput() ScriptSecureStringExecutionParameterResponseOutput
	ToScriptSecureStringExecutionParameterResponseOutputWithContext(context.Context) ScriptSecureStringExecutionParameterResponseOutput
}

type ScriptSecureStringExecutionParameterResponseArgs struct {
	Name        pulumi.StringInput    `pulumi:"name"`
	SecureValue pulumi.StringPtrInput `pulumi:"secureValue"`
	Type        pulumi.StringInput    `pulumi:"type"`
}

func (ScriptSecureStringExecutionParameterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScriptSecureStringExecutionParameterResponse)(nil)).Elem()
}

func (i ScriptSecureStringExecutionParameterResponseArgs) ToScriptSecureStringExecutionParameterResponseOutput() ScriptSecureStringExecutionParameterResponseOutput {
	return i.ToScriptSecureStringExecutionParameterResponseOutputWithContext(context.Background())
}

func (i ScriptSecureStringExecutionParameterResponseArgs) ToScriptSecureStringExecutionParameterResponseOutputWithContext(ctx context.Context) ScriptSecureStringExecutionParameterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScriptSecureStringExecutionParameterResponseOutput)
}

type ScriptSecureStringExecutionParameterResponseOutput struct{ *pulumi.OutputState }

func (ScriptSecureStringExecutionParameterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScriptSecureStringExecutionParameterResponse)(nil)).Elem()
}

func (o ScriptSecureStringExecutionParameterResponseOutput) ToScriptSecureStringExecutionParameterResponseOutput() ScriptSecureStringExecutionParameterResponseOutput {
	return o
}

func (o ScriptSecureStringExecutionParameterResponseOutput) ToScriptSecureStringExecutionParameterResponseOutputWithContext(ctx context.Context) ScriptSecureStringExecutionParameterResponseOutput {
	return o
}

func (o ScriptSecureStringExecutionParameterResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ScriptSecureStringExecutionParameterResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ScriptSecureStringExecutionParameterResponseOutput) SecureValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScriptSecureStringExecutionParameterResponse) *string { return v.SecureValue }).(pulumi.StringPtrOutput)
}

func (o ScriptSecureStringExecutionParameterResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ScriptSecureStringExecutionParameterResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ScriptStringExecutionParameter struct {
	Name  string  `pulumi:"name"`
	Type  string  `pulumi:"type"`
	Value *string `pulumi:"value"`
}





type ScriptStringExecutionParameterInput interface {
	pulumi.Input

	ToScriptStringExecutionParameterOutput() ScriptStringExecutionParameterOutput
	ToScriptStringExecutionParameterOutputWithContext(context.Context) ScriptStringExecutionParameterOutput
}

type ScriptStringExecutionParameterArgs struct {
	Name  pulumi.StringInput    `pulumi:"name"`
	Type  pulumi.StringInput    `pulumi:"type"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (ScriptStringExecutionParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScriptStringExecutionParameter)(nil)).Elem()
}

func (i ScriptStringExecutionParameterArgs) ToScriptStringExecutionParameterOutput() ScriptStringExecutionParameterOutput {
	return i.ToScriptStringExecutionParameterOutputWithContext(context.Background())
}

func (i ScriptStringExecutionParameterArgs) ToScriptStringExecutionParameterOutputWithContext(ctx context.Context) ScriptStringExecutionParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScriptStringExecutionParameterOutput)
}

type ScriptStringExecutionParameterOutput struct{ *pulumi.OutputState }

func (ScriptStringExecutionParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScriptStringExecutionParameter)(nil)).Elem()
}

func (o ScriptStringExecutionParameterOutput) ToScriptStringExecutionParameterOutput() ScriptStringExecutionParameterOutput {
	return o
}

func (o ScriptStringExecutionParameterOutput) ToScriptStringExecutionParameterOutputWithContext(ctx context.Context) ScriptStringExecutionParameterOutput {
	return o
}

func (o ScriptStringExecutionParameterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ScriptStringExecutionParameter) string { return v.Name }).(pulumi.StringOutput)
}

func (o ScriptStringExecutionParameterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ScriptStringExecutionParameter) string { return v.Type }).(pulumi.StringOutput)
}

func (o ScriptStringExecutionParameterOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScriptStringExecutionParameter) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type ScriptStringExecutionParameterResponse struct {
	Name  string  `pulumi:"name"`
	Type  string  `pulumi:"type"`
	Value *string `pulumi:"value"`
}





type ScriptStringExecutionParameterResponseInput interface {
	pulumi.Input

	ToScriptStringExecutionParameterResponseOutput() ScriptStringExecutionParameterResponseOutput
	ToScriptStringExecutionParameterResponseOutputWithContext(context.Context) ScriptStringExecutionParameterResponseOutput
}

type ScriptStringExecutionParameterResponseArgs struct {
	Name  pulumi.StringInput    `pulumi:"name"`
	Type  pulumi.StringInput    `pulumi:"type"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (ScriptStringExecutionParameterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScriptStringExecutionParameterResponse)(nil)).Elem()
}

func (i ScriptStringExecutionParameterResponseArgs) ToScriptStringExecutionParameterResponseOutput() ScriptStringExecutionParameterResponseOutput {
	return i.ToScriptStringExecutionParameterResponseOutputWithContext(context.Background())
}

func (i ScriptStringExecutionParameterResponseArgs) ToScriptStringExecutionParameterResponseOutputWithContext(ctx context.Context) ScriptStringExecutionParameterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScriptStringExecutionParameterResponseOutput)
}

type ScriptStringExecutionParameterResponseOutput struct{ *pulumi.OutputState }

func (ScriptStringExecutionParameterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScriptStringExecutionParameterResponse)(nil)).Elem()
}

func (o ScriptStringExecutionParameterResponseOutput) ToScriptStringExecutionParameterResponseOutput() ScriptStringExecutionParameterResponseOutput {
	return o
}

func (o ScriptStringExecutionParameterResponseOutput) ToScriptStringExecutionParameterResponseOutputWithContext(ctx context.Context) ScriptStringExecutionParameterResponseOutput {
	return o
}

func (o ScriptStringExecutionParameterResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ScriptStringExecutionParameterResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ScriptStringExecutionParameterResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ScriptStringExecutionParameterResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o ScriptStringExecutionParameterResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScriptStringExecutionParameterResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type Sku struct {
	Name string `pulumi:"name"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}









type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
}

type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku {
		if v != nil {
			return *v
		}
		var ret Sku
		return ret
	}).(SkuOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Name string `pulumi:"name"`
}





type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

type SkuResponseArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (SkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (i SkuResponseArgs) ToSkuResponseOutput() SkuResponseOutput {
	return i.ToSkuResponseOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput)
}

func (i SkuResponseArgs) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput).ToSkuResponsePtrOutputWithContext(ctx)
}









type SkuResponsePtrInput interface {
	pulumi.Input

	ToSkuResponsePtrOutput() SkuResponsePtrOutput
	ToSkuResponsePtrOutputWithContext(context.Context) SkuResponsePtrOutput
}

type skuResponsePtrType SkuResponseArgs

func SkuResponsePtr(v *SkuResponseArgs) SkuResponsePtrInput {
	return (*skuResponsePtrType)(v)
}

func (*skuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponsePtrOutput)
}

type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (o SkuResponseOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuResponse) *SkuResponse {
		return &v
	}).(SkuResponsePtrOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse {
		if v != nil {
			return *v
		}
		var ret SkuResponse
		return ret
	}).(SkuResponseOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type VmHostPlacementPolicyProperties struct {
	AffinityType string   `pulumi:"affinityType"`
	DisplayName  *string  `pulumi:"displayName"`
	HostMembers  []string `pulumi:"hostMembers"`
	State        *string  `pulumi:"state"`
	Type         string   `pulumi:"type"`
	VmMembers    []string `pulumi:"vmMembers"`
}





type VmHostPlacementPolicyPropertiesInput interface {
	pulumi.Input

	ToVmHostPlacementPolicyPropertiesOutput() VmHostPlacementPolicyPropertiesOutput
	ToVmHostPlacementPolicyPropertiesOutputWithContext(context.Context) VmHostPlacementPolicyPropertiesOutput
}

type VmHostPlacementPolicyPropertiesArgs struct {
	AffinityType pulumi.StringInput      `pulumi:"affinityType"`
	DisplayName  pulumi.StringPtrInput   `pulumi:"displayName"`
	HostMembers  pulumi.StringArrayInput `pulumi:"hostMembers"`
	State        pulumi.StringPtrInput   `pulumi:"state"`
	Type         pulumi.StringInput      `pulumi:"type"`
	VmMembers    pulumi.StringArrayInput `pulumi:"vmMembers"`
}

func (VmHostPlacementPolicyPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VmHostPlacementPolicyProperties)(nil)).Elem()
}

func (i VmHostPlacementPolicyPropertiesArgs) ToVmHostPlacementPolicyPropertiesOutput() VmHostPlacementPolicyPropertiesOutput {
	return i.ToVmHostPlacementPolicyPropertiesOutputWithContext(context.Background())
}

func (i VmHostPlacementPolicyPropertiesArgs) ToVmHostPlacementPolicyPropertiesOutputWithContext(ctx context.Context) VmHostPlacementPolicyPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmHostPlacementPolicyPropertiesOutput)
}

type VmHostPlacementPolicyPropertiesOutput struct{ *pulumi.OutputState }

func (VmHostPlacementPolicyPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VmHostPlacementPolicyProperties)(nil)).Elem()
}

func (o VmHostPlacementPolicyPropertiesOutput) ToVmHostPlacementPolicyPropertiesOutput() VmHostPlacementPolicyPropertiesOutput {
	return o
}

func (o VmHostPlacementPolicyPropertiesOutput) ToVmHostPlacementPolicyPropertiesOutputWithContext(ctx context.Context) VmHostPlacementPolicyPropertiesOutput {
	return o
}

func (o VmHostPlacementPolicyPropertiesOutput) AffinityType() pulumi.StringOutput {
	return o.ApplyT(func(v VmHostPlacementPolicyProperties) string { return v.AffinityType }).(pulumi.StringOutput)
}

func (o VmHostPlacementPolicyPropertiesOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VmHostPlacementPolicyProperties) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o VmHostPlacementPolicyPropertiesOutput) HostMembers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VmHostPlacementPolicyProperties) []string { return v.HostMembers }).(pulumi.StringArrayOutput)
}

func (o VmHostPlacementPolicyPropertiesOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VmHostPlacementPolicyProperties) *string { return v.State }).(pulumi.StringPtrOutput)
}

func (o VmHostPlacementPolicyPropertiesOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VmHostPlacementPolicyProperties) string { return v.Type }).(pulumi.StringOutput)
}

func (o VmHostPlacementPolicyPropertiesOutput) VmMembers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VmHostPlacementPolicyProperties) []string { return v.VmMembers }).(pulumi.StringArrayOutput)
}

type VmHostPlacementPolicyPropertiesResponse struct {
	AffinityType      string   `pulumi:"affinityType"`
	DisplayName       *string  `pulumi:"displayName"`
	HostMembers       []string `pulumi:"hostMembers"`
	ProvisioningState string   `pulumi:"provisioningState"`
	State             *string  `pulumi:"state"`
	Type              string   `pulumi:"type"`
	VmMembers         []string `pulumi:"vmMembers"`
}





type VmHostPlacementPolicyPropertiesResponseInput interface {
	pulumi.Input

	ToVmHostPlacementPolicyPropertiesResponseOutput() VmHostPlacementPolicyPropertiesResponseOutput
	ToVmHostPlacementPolicyPropertiesResponseOutputWithContext(context.Context) VmHostPlacementPolicyPropertiesResponseOutput
}

type VmHostPlacementPolicyPropertiesResponseArgs struct {
	AffinityType      pulumi.StringInput      `pulumi:"affinityType"`
	DisplayName       pulumi.StringPtrInput   `pulumi:"displayName"`
	HostMembers       pulumi.StringArrayInput `pulumi:"hostMembers"`
	ProvisioningState pulumi.StringInput      `pulumi:"provisioningState"`
	State             pulumi.StringPtrInput   `pulumi:"state"`
	Type              pulumi.StringInput      `pulumi:"type"`
	VmMembers         pulumi.StringArrayInput `pulumi:"vmMembers"`
}

func (VmHostPlacementPolicyPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VmHostPlacementPolicyPropertiesResponse)(nil)).Elem()
}

func (i VmHostPlacementPolicyPropertiesResponseArgs) ToVmHostPlacementPolicyPropertiesResponseOutput() VmHostPlacementPolicyPropertiesResponseOutput {
	return i.ToVmHostPlacementPolicyPropertiesResponseOutputWithContext(context.Background())
}

func (i VmHostPlacementPolicyPropertiesResponseArgs) ToVmHostPlacementPolicyPropertiesResponseOutputWithContext(ctx context.Context) VmHostPlacementPolicyPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmHostPlacementPolicyPropertiesResponseOutput)
}

type VmHostPlacementPolicyPropertiesResponseOutput struct{ *pulumi.OutputState }

func (VmHostPlacementPolicyPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VmHostPlacementPolicyPropertiesResponse)(nil)).Elem()
}

func (o VmHostPlacementPolicyPropertiesResponseOutput) ToVmHostPlacementPolicyPropertiesResponseOutput() VmHostPlacementPolicyPropertiesResponseOutput {
	return o
}

func (o VmHostPlacementPolicyPropertiesResponseOutput) ToVmHostPlacementPolicyPropertiesResponseOutputWithContext(ctx context.Context) VmHostPlacementPolicyPropertiesResponseOutput {
	return o
}

func (o VmHostPlacementPolicyPropertiesResponseOutput) AffinityType() pulumi.StringOutput {
	return o.ApplyT(func(v VmHostPlacementPolicyPropertiesResponse) string { return v.AffinityType }).(pulumi.StringOutput)
}

func (o VmHostPlacementPolicyPropertiesResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VmHostPlacementPolicyPropertiesResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o VmHostPlacementPolicyPropertiesResponseOutput) HostMembers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VmHostPlacementPolicyPropertiesResponse) []string { return v.HostMembers }).(pulumi.StringArrayOutput)
}

func (o VmHostPlacementPolicyPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VmHostPlacementPolicyPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VmHostPlacementPolicyPropertiesResponseOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VmHostPlacementPolicyPropertiesResponse) *string { return v.State }).(pulumi.StringPtrOutput)
}

func (o VmHostPlacementPolicyPropertiesResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VmHostPlacementPolicyPropertiesResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o VmHostPlacementPolicyPropertiesResponseOutput) VmMembers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VmHostPlacementPolicyPropertiesResponse) []string { return v.VmMembers }).(pulumi.StringArrayOutput)
}

type VmVmPlacementPolicyProperties struct {
	AffinityType string   `pulumi:"affinityType"`
	DisplayName  *string  `pulumi:"displayName"`
	State        *string  `pulumi:"state"`
	Type         string   `pulumi:"type"`
	VmMembers    []string `pulumi:"vmMembers"`
}





type VmVmPlacementPolicyPropertiesInput interface {
	pulumi.Input

	ToVmVmPlacementPolicyPropertiesOutput() VmVmPlacementPolicyPropertiesOutput
	ToVmVmPlacementPolicyPropertiesOutputWithContext(context.Context) VmVmPlacementPolicyPropertiesOutput
}

type VmVmPlacementPolicyPropertiesArgs struct {
	AffinityType pulumi.StringInput      `pulumi:"affinityType"`
	DisplayName  pulumi.StringPtrInput   `pulumi:"displayName"`
	State        pulumi.StringPtrInput   `pulumi:"state"`
	Type         pulumi.StringInput      `pulumi:"type"`
	VmMembers    pulumi.StringArrayInput `pulumi:"vmMembers"`
}

func (VmVmPlacementPolicyPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VmVmPlacementPolicyProperties)(nil)).Elem()
}

func (i VmVmPlacementPolicyPropertiesArgs) ToVmVmPlacementPolicyPropertiesOutput() VmVmPlacementPolicyPropertiesOutput {
	return i.ToVmVmPlacementPolicyPropertiesOutputWithContext(context.Background())
}

func (i VmVmPlacementPolicyPropertiesArgs) ToVmVmPlacementPolicyPropertiesOutputWithContext(ctx context.Context) VmVmPlacementPolicyPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmVmPlacementPolicyPropertiesOutput)
}

type VmVmPlacementPolicyPropertiesOutput struct{ *pulumi.OutputState }

func (VmVmPlacementPolicyPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VmVmPlacementPolicyProperties)(nil)).Elem()
}

func (o VmVmPlacementPolicyPropertiesOutput) ToVmVmPlacementPolicyPropertiesOutput() VmVmPlacementPolicyPropertiesOutput {
	return o
}

func (o VmVmPlacementPolicyPropertiesOutput) ToVmVmPlacementPolicyPropertiesOutputWithContext(ctx context.Context) VmVmPlacementPolicyPropertiesOutput {
	return o
}

func (o VmVmPlacementPolicyPropertiesOutput) AffinityType() pulumi.StringOutput {
	return o.ApplyT(func(v VmVmPlacementPolicyProperties) string { return v.AffinityType }).(pulumi.StringOutput)
}

func (o VmVmPlacementPolicyPropertiesOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VmVmPlacementPolicyProperties) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o VmVmPlacementPolicyPropertiesOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VmVmPlacementPolicyProperties) *string { return v.State }).(pulumi.StringPtrOutput)
}

func (o VmVmPlacementPolicyPropertiesOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VmVmPlacementPolicyProperties) string { return v.Type }).(pulumi.StringOutput)
}

func (o VmVmPlacementPolicyPropertiesOutput) VmMembers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VmVmPlacementPolicyProperties) []string { return v.VmMembers }).(pulumi.StringArrayOutput)
}

type VmVmPlacementPolicyPropertiesResponse struct {
	AffinityType      string   `pulumi:"affinityType"`
	DisplayName       *string  `pulumi:"displayName"`
	ProvisioningState string   `pulumi:"provisioningState"`
	State             *string  `pulumi:"state"`
	Type              string   `pulumi:"type"`
	VmMembers         []string `pulumi:"vmMembers"`
}





type VmVmPlacementPolicyPropertiesResponseInput interface {
	pulumi.Input

	ToVmVmPlacementPolicyPropertiesResponseOutput() VmVmPlacementPolicyPropertiesResponseOutput
	ToVmVmPlacementPolicyPropertiesResponseOutputWithContext(context.Context) VmVmPlacementPolicyPropertiesResponseOutput
}

type VmVmPlacementPolicyPropertiesResponseArgs struct {
	AffinityType      pulumi.StringInput      `pulumi:"affinityType"`
	DisplayName       pulumi.StringPtrInput   `pulumi:"displayName"`
	ProvisioningState pulumi.StringInput      `pulumi:"provisioningState"`
	State             pulumi.StringPtrInput   `pulumi:"state"`
	Type              pulumi.StringInput      `pulumi:"type"`
	VmMembers         pulumi.StringArrayInput `pulumi:"vmMembers"`
}

func (VmVmPlacementPolicyPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VmVmPlacementPolicyPropertiesResponse)(nil)).Elem()
}

func (i VmVmPlacementPolicyPropertiesResponseArgs) ToVmVmPlacementPolicyPropertiesResponseOutput() VmVmPlacementPolicyPropertiesResponseOutput {
	return i.ToVmVmPlacementPolicyPropertiesResponseOutputWithContext(context.Background())
}

func (i VmVmPlacementPolicyPropertiesResponseArgs) ToVmVmPlacementPolicyPropertiesResponseOutputWithContext(ctx context.Context) VmVmPlacementPolicyPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmVmPlacementPolicyPropertiesResponseOutput)
}

type VmVmPlacementPolicyPropertiesResponseOutput struct{ *pulumi.OutputState }

func (VmVmPlacementPolicyPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VmVmPlacementPolicyPropertiesResponse)(nil)).Elem()
}

func (o VmVmPlacementPolicyPropertiesResponseOutput) ToVmVmPlacementPolicyPropertiesResponseOutput() VmVmPlacementPolicyPropertiesResponseOutput {
	return o
}

func (o VmVmPlacementPolicyPropertiesResponseOutput) ToVmVmPlacementPolicyPropertiesResponseOutputWithContext(ctx context.Context) VmVmPlacementPolicyPropertiesResponseOutput {
	return o
}

func (o VmVmPlacementPolicyPropertiesResponseOutput) AffinityType() pulumi.StringOutput {
	return o.ApplyT(func(v VmVmPlacementPolicyPropertiesResponse) string { return v.AffinityType }).(pulumi.StringOutput)
}

func (o VmVmPlacementPolicyPropertiesResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VmVmPlacementPolicyPropertiesResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o VmVmPlacementPolicyPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VmVmPlacementPolicyPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VmVmPlacementPolicyPropertiesResponseOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VmVmPlacementPolicyPropertiesResponse) *string { return v.State }).(pulumi.StringPtrOutput)
}

func (o VmVmPlacementPolicyPropertiesResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VmVmPlacementPolicyPropertiesResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o VmVmPlacementPolicyPropertiesResponseOutput) VmMembers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VmVmPlacementPolicyPropertiesResponse) []string { return v.VmMembers }).(pulumi.StringArrayOutput)
}

type WorkloadNetworkDhcpRelay struct {
	DhcpType        string   `pulumi:"dhcpType"`
	DisplayName     *string  `pulumi:"displayName"`
	Revision        *float64 `pulumi:"revision"`
	ServerAddresses []string `pulumi:"serverAddresses"`
}





type WorkloadNetworkDhcpRelayInput interface {
	pulumi.Input

	ToWorkloadNetworkDhcpRelayOutput() WorkloadNetworkDhcpRelayOutput
	ToWorkloadNetworkDhcpRelayOutputWithContext(context.Context) WorkloadNetworkDhcpRelayOutput
}

type WorkloadNetworkDhcpRelayArgs struct {
	DhcpType        pulumi.StringInput      `pulumi:"dhcpType"`
	DisplayName     pulumi.StringPtrInput   `pulumi:"displayName"`
	Revision        pulumi.Float64PtrInput  `pulumi:"revision"`
	ServerAddresses pulumi.StringArrayInput `pulumi:"serverAddresses"`
}

func (WorkloadNetworkDhcpRelayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadNetworkDhcpRelay)(nil)).Elem()
}

func (i WorkloadNetworkDhcpRelayArgs) ToWorkloadNetworkDhcpRelayOutput() WorkloadNetworkDhcpRelayOutput {
	return i.ToWorkloadNetworkDhcpRelayOutputWithContext(context.Background())
}

func (i WorkloadNetworkDhcpRelayArgs) ToWorkloadNetworkDhcpRelayOutputWithContext(ctx context.Context) WorkloadNetworkDhcpRelayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadNetworkDhcpRelayOutput)
}

type WorkloadNetworkDhcpRelayOutput struct{ *pulumi.OutputState }

func (WorkloadNetworkDhcpRelayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadNetworkDhcpRelay)(nil)).Elem()
}

func (o WorkloadNetworkDhcpRelayOutput) ToWorkloadNetworkDhcpRelayOutput() WorkloadNetworkDhcpRelayOutput {
	return o
}

func (o WorkloadNetworkDhcpRelayOutput) ToWorkloadNetworkDhcpRelayOutputWithContext(ctx context.Context) WorkloadNetworkDhcpRelayOutput {
	return o
}

func (o WorkloadNetworkDhcpRelayOutput) DhcpType() pulumi.StringOutput {
	return o.ApplyT(func(v WorkloadNetworkDhcpRelay) string { return v.DhcpType }).(pulumi.StringOutput)
}

func (o WorkloadNetworkDhcpRelayOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadNetworkDhcpRelay) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o WorkloadNetworkDhcpRelayOutput) Revision() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v WorkloadNetworkDhcpRelay) *float64 { return v.Revision }).(pulumi.Float64PtrOutput)
}

func (o WorkloadNetworkDhcpRelayOutput) ServerAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WorkloadNetworkDhcpRelay) []string { return v.ServerAddresses }).(pulumi.StringArrayOutput)
}

type WorkloadNetworkDhcpRelayResponse struct {
	DhcpType          string   `pulumi:"dhcpType"`
	DisplayName       *string  `pulumi:"displayName"`
	ProvisioningState string   `pulumi:"provisioningState"`
	Revision          *float64 `pulumi:"revision"`
	Segments          []string `pulumi:"segments"`
	ServerAddresses   []string `pulumi:"serverAddresses"`
}





type WorkloadNetworkDhcpRelayResponseInput interface {
	pulumi.Input

	ToWorkloadNetworkDhcpRelayResponseOutput() WorkloadNetworkDhcpRelayResponseOutput
	ToWorkloadNetworkDhcpRelayResponseOutputWithContext(context.Context) WorkloadNetworkDhcpRelayResponseOutput
}

type WorkloadNetworkDhcpRelayResponseArgs struct {
	DhcpType          pulumi.StringInput      `pulumi:"dhcpType"`
	DisplayName       pulumi.StringPtrInput   `pulumi:"displayName"`
	ProvisioningState pulumi.StringInput      `pulumi:"provisioningState"`
	Revision          pulumi.Float64PtrInput  `pulumi:"revision"`
	Segments          pulumi.StringArrayInput `pulumi:"segments"`
	ServerAddresses   pulumi.StringArrayInput `pulumi:"serverAddresses"`
}

func (WorkloadNetworkDhcpRelayResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadNetworkDhcpRelayResponse)(nil)).Elem()
}

func (i WorkloadNetworkDhcpRelayResponseArgs) ToWorkloadNetworkDhcpRelayResponseOutput() WorkloadNetworkDhcpRelayResponseOutput {
	return i.ToWorkloadNetworkDhcpRelayResponseOutputWithContext(context.Background())
}

func (i WorkloadNetworkDhcpRelayResponseArgs) ToWorkloadNetworkDhcpRelayResponseOutputWithContext(ctx context.Context) WorkloadNetworkDhcpRelayResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadNetworkDhcpRelayResponseOutput)
}

type WorkloadNetworkDhcpRelayResponseOutput struct{ *pulumi.OutputState }

func (WorkloadNetworkDhcpRelayResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadNetworkDhcpRelayResponse)(nil)).Elem()
}

func (o WorkloadNetworkDhcpRelayResponseOutput) ToWorkloadNetworkDhcpRelayResponseOutput() WorkloadNetworkDhcpRelayResponseOutput {
	return o
}

func (o WorkloadNetworkDhcpRelayResponseOutput) ToWorkloadNetworkDhcpRelayResponseOutputWithContext(ctx context.Context) WorkloadNetworkDhcpRelayResponseOutput {
	return o
}

func (o WorkloadNetworkDhcpRelayResponseOutput) DhcpType() pulumi.StringOutput {
	return o.ApplyT(func(v WorkloadNetworkDhcpRelayResponse) string { return v.DhcpType }).(pulumi.StringOutput)
}

func (o WorkloadNetworkDhcpRelayResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadNetworkDhcpRelayResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o WorkloadNetworkDhcpRelayResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v WorkloadNetworkDhcpRelayResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o WorkloadNetworkDhcpRelayResponseOutput) Revision() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v WorkloadNetworkDhcpRelayResponse) *float64 { return v.Revision }).(pulumi.Float64PtrOutput)
}

func (o WorkloadNetworkDhcpRelayResponseOutput) Segments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WorkloadNetworkDhcpRelayResponse) []string { return v.Segments }).(pulumi.StringArrayOutput)
}

func (o WorkloadNetworkDhcpRelayResponseOutput) ServerAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WorkloadNetworkDhcpRelayResponse) []string { return v.ServerAddresses }).(pulumi.StringArrayOutput)
}

type WorkloadNetworkDhcpServer struct {
	DhcpType      string   `pulumi:"dhcpType"`
	DisplayName   *string  `pulumi:"displayName"`
	LeaseTime     *float64 `pulumi:"leaseTime"`
	Revision      *float64 `pulumi:"revision"`
	ServerAddress *string  `pulumi:"serverAddress"`
}





type WorkloadNetworkDhcpServerInput interface {
	pulumi.Input

	ToWorkloadNetworkDhcpServerOutput() WorkloadNetworkDhcpServerOutput
	ToWorkloadNetworkDhcpServerOutputWithContext(context.Context) WorkloadNetworkDhcpServerOutput
}

type WorkloadNetworkDhcpServerArgs struct {
	DhcpType      pulumi.StringInput     `pulumi:"dhcpType"`
	DisplayName   pulumi.StringPtrInput  `pulumi:"displayName"`
	LeaseTime     pulumi.Float64PtrInput `pulumi:"leaseTime"`
	Revision      pulumi.Float64PtrInput `pulumi:"revision"`
	ServerAddress pulumi.StringPtrInput  `pulumi:"serverAddress"`
}

func (WorkloadNetworkDhcpServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadNetworkDhcpServer)(nil)).Elem()
}

func (i WorkloadNetworkDhcpServerArgs) ToWorkloadNetworkDhcpServerOutput() WorkloadNetworkDhcpServerOutput {
	return i.ToWorkloadNetworkDhcpServerOutputWithContext(context.Background())
}

func (i WorkloadNetworkDhcpServerArgs) ToWorkloadNetworkDhcpServerOutputWithContext(ctx context.Context) WorkloadNetworkDhcpServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadNetworkDhcpServerOutput)
}

type WorkloadNetworkDhcpServerOutput struct{ *pulumi.OutputState }

func (WorkloadNetworkDhcpServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadNetworkDhcpServer)(nil)).Elem()
}

func (o WorkloadNetworkDhcpServerOutput) ToWorkloadNetworkDhcpServerOutput() WorkloadNetworkDhcpServerOutput {
	return o
}

func (o WorkloadNetworkDhcpServerOutput) ToWorkloadNetworkDhcpServerOutputWithContext(ctx context.Context) WorkloadNetworkDhcpServerOutput {
	return o
}

func (o WorkloadNetworkDhcpServerOutput) DhcpType() pulumi.StringOutput {
	return o.ApplyT(func(v WorkloadNetworkDhcpServer) string { return v.DhcpType }).(pulumi.StringOutput)
}

func (o WorkloadNetworkDhcpServerOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadNetworkDhcpServer) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o WorkloadNetworkDhcpServerOutput) LeaseTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v WorkloadNetworkDhcpServer) *float64 { return v.LeaseTime }).(pulumi.Float64PtrOutput)
}

func (o WorkloadNetworkDhcpServerOutput) Revision() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v WorkloadNetworkDhcpServer) *float64 { return v.Revision }).(pulumi.Float64PtrOutput)
}

func (o WorkloadNetworkDhcpServerOutput) ServerAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadNetworkDhcpServer) *string { return v.ServerAddress }).(pulumi.StringPtrOutput)
}

type WorkloadNetworkDhcpServerResponse struct {
	DhcpType          string   `pulumi:"dhcpType"`
	DisplayName       *string  `pulumi:"displayName"`
	LeaseTime         *float64 `pulumi:"leaseTime"`
	ProvisioningState string   `pulumi:"provisioningState"`
	Revision          *float64 `pulumi:"revision"`
	Segments          []string `pulumi:"segments"`
	ServerAddress     *string  `pulumi:"serverAddress"`
}





type WorkloadNetworkDhcpServerResponseInput interface {
	pulumi.Input

	ToWorkloadNetworkDhcpServerResponseOutput() WorkloadNetworkDhcpServerResponseOutput
	ToWorkloadNetworkDhcpServerResponseOutputWithContext(context.Context) WorkloadNetworkDhcpServerResponseOutput
}

type WorkloadNetworkDhcpServerResponseArgs struct {
	DhcpType          pulumi.StringInput      `pulumi:"dhcpType"`
	DisplayName       pulumi.StringPtrInput   `pulumi:"displayName"`
	LeaseTime         pulumi.Float64PtrInput  `pulumi:"leaseTime"`
	ProvisioningState pulumi.StringInput      `pulumi:"provisioningState"`
	Revision          pulumi.Float64PtrInput  `pulumi:"revision"`
	Segments          pulumi.StringArrayInput `pulumi:"segments"`
	ServerAddress     pulumi.StringPtrInput   `pulumi:"serverAddress"`
}

func (WorkloadNetworkDhcpServerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadNetworkDhcpServerResponse)(nil)).Elem()
}

func (i WorkloadNetworkDhcpServerResponseArgs) ToWorkloadNetworkDhcpServerResponseOutput() WorkloadNetworkDhcpServerResponseOutput {
	return i.ToWorkloadNetworkDhcpServerResponseOutputWithContext(context.Background())
}

func (i WorkloadNetworkDhcpServerResponseArgs) ToWorkloadNetworkDhcpServerResponseOutputWithContext(ctx context.Context) WorkloadNetworkDhcpServerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadNetworkDhcpServerResponseOutput)
}

type WorkloadNetworkDhcpServerResponseOutput struct{ *pulumi.OutputState }

func (WorkloadNetworkDhcpServerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadNetworkDhcpServerResponse)(nil)).Elem()
}

func (o WorkloadNetworkDhcpServerResponseOutput) ToWorkloadNetworkDhcpServerResponseOutput() WorkloadNetworkDhcpServerResponseOutput {
	return o
}

func (o WorkloadNetworkDhcpServerResponseOutput) ToWorkloadNetworkDhcpServerResponseOutputWithContext(ctx context.Context) WorkloadNetworkDhcpServerResponseOutput {
	return o
}

func (o WorkloadNetworkDhcpServerResponseOutput) DhcpType() pulumi.StringOutput {
	return o.ApplyT(func(v WorkloadNetworkDhcpServerResponse) string { return v.DhcpType }).(pulumi.StringOutput)
}

func (o WorkloadNetworkDhcpServerResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadNetworkDhcpServerResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o WorkloadNetworkDhcpServerResponseOutput) LeaseTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v WorkloadNetworkDhcpServerResponse) *float64 { return v.LeaseTime }).(pulumi.Float64PtrOutput)
}

func (o WorkloadNetworkDhcpServerResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v WorkloadNetworkDhcpServerResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o WorkloadNetworkDhcpServerResponseOutput) Revision() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v WorkloadNetworkDhcpServerResponse) *float64 { return v.Revision }).(pulumi.Float64PtrOutput)
}

func (o WorkloadNetworkDhcpServerResponseOutput) Segments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WorkloadNetworkDhcpServerResponse) []string { return v.Segments }).(pulumi.StringArrayOutput)
}

func (o WorkloadNetworkDhcpServerResponseOutput) ServerAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadNetworkDhcpServerResponse) *string { return v.ServerAddress }).(pulumi.StringPtrOutput)
}

type WorkloadNetworkSegmentPortVifResponse struct {
	PortName *string `pulumi:"portName"`
}





type WorkloadNetworkSegmentPortVifResponseInput interface {
	pulumi.Input

	ToWorkloadNetworkSegmentPortVifResponseOutput() WorkloadNetworkSegmentPortVifResponseOutput
	ToWorkloadNetworkSegmentPortVifResponseOutputWithContext(context.Context) WorkloadNetworkSegmentPortVifResponseOutput
}

type WorkloadNetworkSegmentPortVifResponseArgs struct {
	PortName pulumi.StringPtrInput `pulumi:"portName"`
}

func (WorkloadNetworkSegmentPortVifResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadNetworkSegmentPortVifResponse)(nil)).Elem()
}

func (i WorkloadNetworkSegmentPortVifResponseArgs) ToWorkloadNetworkSegmentPortVifResponseOutput() WorkloadNetworkSegmentPortVifResponseOutput {
	return i.ToWorkloadNetworkSegmentPortVifResponseOutputWithContext(context.Background())
}

func (i WorkloadNetworkSegmentPortVifResponseArgs) ToWorkloadNetworkSegmentPortVifResponseOutputWithContext(ctx context.Context) WorkloadNetworkSegmentPortVifResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadNetworkSegmentPortVifResponseOutput)
}





type WorkloadNetworkSegmentPortVifResponseArrayInput interface {
	pulumi.Input

	ToWorkloadNetworkSegmentPortVifResponseArrayOutput() WorkloadNetworkSegmentPortVifResponseArrayOutput
	ToWorkloadNetworkSegmentPortVifResponseArrayOutputWithContext(context.Context) WorkloadNetworkSegmentPortVifResponseArrayOutput
}

type WorkloadNetworkSegmentPortVifResponseArray []WorkloadNetworkSegmentPortVifResponseInput

func (WorkloadNetworkSegmentPortVifResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkloadNetworkSegmentPortVifResponse)(nil)).Elem()
}

func (i WorkloadNetworkSegmentPortVifResponseArray) ToWorkloadNetworkSegmentPortVifResponseArrayOutput() WorkloadNetworkSegmentPortVifResponseArrayOutput {
	return i.ToWorkloadNetworkSegmentPortVifResponseArrayOutputWithContext(context.Background())
}

func (i WorkloadNetworkSegmentPortVifResponseArray) ToWorkloadNetworkSegmentPortVifResponseArrayOutputWithContext(ctx context.Context) WorkloadNetworkSegmentPortVifResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadNetworkSegmentPortVifResponseArrayOutput)
}

type WorkloadNetworkSegmentPortVifResponseOutput struct{ *pulumi.OutputState }

func (WorkloadNetworkSegmentPortVifResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadNetworkSegmentPortVifResponse)(nil)).Elem()
}

func (o WorkloadNetworkSegmentPortVifResponseOutput) ToWorkloadNetworkSegmentPortVifResponseOutput() WorkloadNetworkSegmentPortVifResponseOutput {
	return o
}

func (o WorkloadNetworkSegmentPortVifResponseOutput) ToWorkloadNetworkSegmentPortVifResponseOutputWithContext(ctx context.Context) WorkloadNetworkSegmentPortVifResponseOutput {
	return o
}

func (o WorkloadNetworkSegmentPortVifResponseOutput) PortName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadNetworkSegmentPortVifResponse) *string { return v.PortName }).(pulumi.StringPtrOutput)
}

type WorkloadNetworkSegmentPortVifResponseArrayOutput struct{ *pulumi.OutputState }

func (WorkloadNetworkSegmentPortVifResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkloadNetworkSegmentPortVifResponse)(nil)).Elem()
}

func (o WorkloadNetworkSegmentPortVifResponseArrayOutput) ToWorkloadNetworkSegmentPortVifResponseArrayOutput() WorkloadNetworkSegmentPortVifResponseArrayOutput {
	return o
}

func (o WorkloadNetworkSegmentPortVifResponseArrayOutput) ToWorkloadNetworkSegmentPortVifResponseArrayOutputWithContext(ctx context.Context) WorkloadNetworkSegmentPortVifResponseArrayOutput {
	return o
}

func (o WorkloadNetworkSegmentPortVifResponseArrayOutput) Index(i pulumi.IntInput) WorkloadNetworkSegmentPortVifResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WorkloadNetworkSegmentPortVifResponse {
		return vs[0].([]WorkloadNetworkSegmentPortVifResponse)[vs[1].(int)]
	}).(WorkloadNetworkSegmentPortVifResponseOutput)
}

type WorkloadNetworkSegmentSubnet struct {
	DhcpRanges     []string `pulumi:"dhcpRanges"`
	GatewayAddress *string  `pulumi:"gatewayAddress"`
}





type WorkloadNetworkSegmentSubnetInput interface {
	pulumi.Input

	ToWorkloadNetworkSegmentSubnetOutput() WorkloadNetworkSegmentSubnetOutput
	ToWorkloadNetworkSegmentSubnetOutputWithContext(context.Context) WorkloadNetworkSegmentSubnetOutput
}

type WorkloadNetworkSegmentSubnetArgs struct {
	DhcpRanges     pulumi.StringArrayInput `pulumi:"dhcpRanges"`
	GatewayAddress pulumi.StringPtrInput   `pulumi:"gatewayAddress"`
}

func (WorkloadNetworkSegmentSubnetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadNetworkSegmentSubnet)(nil)).Elem()
}

func (i WorkloadNetworkSegmentSubnetArgs) ToWorkloadNetworkSegmentSubnetOutput() WorkloadNetworkSegmentSubnetOutput {
	return i.ToWorkloadNetworkSegmentSubnetOutputWithContext(context.Background())
}

func (i WorkloadNetworkSegmentSubnetArgs) ToWorkloadNetworkSegmentSubnetOutputWithContext(ctx context.Context) WorkloadNetworkSegmentSubnetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadNetworkSegmentSubnetOutput)
}

func (i WorkloadNetworkSegmentSubnetArgs) ToWorkloadNetworkSegmentSubnetPtrOutput() WorkloadNetworkSegmentSubnetPtrOutput {
	return i.ToWorkloadNetworkSegmentSubnetPtrOutputWithContext(context.Background())
}

func (i WorkloadNetworkSegmentSubnetArgs) ToWorkloadNetworkSegmentSubnetPtrOutputWithContext(ctx context.Context) WorkloadNetworkSegmentSubnetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadNetworkSegmentSubnetOutput).ToWorkloadNetworkSegmentSubnetPtrOutputWithContext(ctx)
}









type WorkloadNetworkSegmentSubnetPtrInput interface {
	pulumi.Input

	ToWorkloadNetworkSegmentSubnetPtrOutput() WorkloadNetworkSegmentSubnetPtrOutput
	ToWorkloadNetworkSegmentSubnetPtrOutputWithContext(context.Context) WorkloadNetworkSegmentSubnetPtrOutput
}

type workloadNetworkSegmentSubnetPtrType WorkloadNetworkSegmentSubnetArgs

func WorkloadNetworkSegmentSubnetPtr(v *WorkloadNetworkSegmentSubnetArgs) WorkloadNetworkSegmentSubnetPtrInput {
	return (*workloadNetworkSegmentSubnetPtrType)(v)
}

func (*workloadNetworkSegmentSubnetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadNetworkSegmentSubnet)(nil)).Elem()
}

func (i *workloadNetworkSegmentSubnetPtrType) ToWorkloadNetworkSegmentSubnetPtrOutput() WorkloadNetworkSegmentSubnetPtrOutput {
	return i.ToWorkloadNetworkSegmentSubnetPtrOutputWithContext(context.Background())
}

func (i *workloadNetworkSegmentSubnetPtrType) ToWorkloadNetworkSegmentSubnetPtrOutputWithContext(ctx context.Context) WorkloadNetworkSegmentSubnetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadNetworkSegmentSubnetPtrOutput)
}

type WorkloadNetworkSegmentSubnetOutput struct{ *pulumi.OutputState }

func (WorkloadNetworkSegmentSubnetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadNetworkSegmentSubnet)(nil)).Elem()
}

func (o WorkloadNetworkSegmentSubnetOutput) ToWorkloadNetworkSegmentSubnetOutput() WorkloadNetworkSegmentSubnetOutput {
	return o
}

func (o WorkloadNetworkSegmentSubnetOutput) ToWorkloadNetworkSegmentSubnetOutputWithContext(ctx context.Context) WorkloadNetworkSegmentSubnetOutput {
	return o
}

func (o WorkloadNetworkSegmentSubnetOutput) ToWorkloadNetworkSegmentSubnetPtrOutput() WorkloadNetworkSegmentSubnetPtrOutput {
	return o.ToWorkloadNetworkSegmentSubnetPtrOutputWithContext(context.Background())
}

func (o WorkloadNetworkSegmentSubnetOutput) ToWorkloadNetworkSegmentSubnetPtrOutputWithContext(ctx context.Context) WorkloadNetworkSegmentSubnetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkloadNetworkSegmentSubnet) *WorkloadNetworkSegmentSubnet {
		return &v
	}).(WorkloadNetworkSegmentSubnetPtrOutput)
}

func (o WorkloadNetworkSegmentSubnetOutput) DhcpRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WorkloadNetworkSegmentSubnet) []string { return v.DhcpRanges }).(pulumi.StringArrayOutput)
}

func (o WorkloadNetworkSegmentSubnetOutput) GatewayAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadNetworkSegmentSubnet) *string { return v.GatewayAddress }).(pulumi.StringPtrOutput)
}

type WorkloadNetworkSegmentSubnetPtrOutput struct{ *pulumi.OutputState }

func (WorkloadNetworkSegmentSubnetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadNetworkSegmentSubnet)(nil)).Elem()
}

func (o WorkloadNetworkSegmentSubnetPtrOutput) ToWorkloadNetworkSegmentSubnetPtrOutput() WorkloadNetworkSegmentSubnetPtrOutput {
	return o
}

func (o WorkloadNetworkSegmentSubnetPtrOutput) ToWorkloadNetworkSegmentSubnetPtrOutputWithContext(ctx context.Context) WorkloadNetworkSegmentSubnetPtrOutput {
	return o
}

func (o WorkloadNetworkSegmentSubnetPtrOutput) Elem() WorkloadNetworkSegmentSubnetOutput {
	return o.ApplyT(func(v *WorkloadNetworkSegmentSubnet) WorkloadNetworkSegmentSubnet {
		if v != nil {
			return *v
		}
		var ret WorkloadNetworkSegmentSubnet
		return ret
	}).(WorkloadNetworkSegmentSubnetOutput)
}

func (o WorkloadNetworkSegmentSubnetPtrOutput) DhcpRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WorkloadNetworkSegmentSubnet) []string {
		if v == nil {
			return nil
		}
		return v.DhcpRanges
	}).(pulumi.StringArrayOutput)
}

func (o WorkloadNetworkSegmentSubnetPtrOutput) GatewayAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkSegmentSubnet) *string {
		if v == nil {
			return nil
		}
		return v.GatewayAddress
	}).(pulumi.StringPtrOutput)
}

type WorkloadNetworkSegmentSubnetResponse struct {
	DhcpRanges     []string `pulumi:"dhcpRanges"`
	GatewayAddress *string  `pulumi:"gatewayAddress"`
}





type WorkloadNetworkSegmentSubnetResponseInput interface {
	pulumi.Input

	ToWorkloadNetworkSegmentSubnetResponseOutput() WorkloadNetworkSegmentSubnetResponseOutput
	ToWorkloadNetworkSegmentSubnetResponseOutputWithContext(context.Context) WorkloadNetworkSegmentSubnetResponseOutput
}

type WorkloadNetworkSegmentSubnetResponseArgs struct {
	DhcpRanges     pulumi.StringArrayInput `pulumi:"dhcpRanges"`
	GatewayAddress pulumi.StringPtrInput   `pulumi:"gatewayAddress"`
}

func (WorkloadNetworkSegmentSubnetResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadNetworkSegmentSubnetResponse)(nil)).Elem()
}

func (i WorkloadNetworkSegmentSubnetResponseArgs) ToWorkloadNetworkSegmentSubnetResponseOutput() WorkloadNetworkSegmentSubnetResponseOutput {
	return i.ToWorkloadNetworkSegmentSubnetResponseOutputWithContext(context.Background())
}

func (i WorkloadNetworkSegmentSubnetResponseArgs) ToWorkloadNetworkSegmentSubnetResponseOutputWithContext(ctx context.Context) WorkloadNetworkSegmentSubnetResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadNetworkSegmentSubnetResponseOutput)
}

func (i WorkloadNetworkSegmentSubnetResponseArgs) ToWorkloadNetworkSegmentSubnetResponsePtrOutput() WorkloadNetworkSegmentSubnetResponsePtrOutput {
	return i.ToWorkloadNetworkSegmentSubnetResponsePtrOutputWithContext(context.Background())
}

func (i WorkloadNetworkSegmentSubnetResponseArgs) ToWorkloadNetworkSegmentSubnetResponsePtrOutputWithContext(ctx context.Context) WorkloadNetworkSegmentSubnetResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadNetworkSegmentSubnetResponseOutput).ToWorkloadNetworkSegmentSubnetResponsePtrOutputWithContext(ctx)
}









type WorkloadNetworkSegmentSubnetResponsePtrInput interface {
	pulumi.Input

	ToWorkloadNetworkSegmentSubnetResponsePtrOutput() WorkloadNetworkSegmentSubnetResponsePtrOutput
	ToWorkloadNetworkSegmentSubnetResponsePtrOutputWithContext(context.Context) WorkloadNetworkSegmentSubnetResponsePtrOutput
}

type workloadNetworkSegmentSubnetResponsePtrType WorkloadNetworkSegmentSubnetResponseArgs

func WorkloadNetworkSegmentSubnetResponsePtr(v *WorkloadNetworkSegmentSubnetResponseArgs) WorkloadNetworkSegmentSubnetResponsePtrInput {
	return (*workloadNetworkSegmentSubnetResponsePtrType)(v)
}

func (*workloadNetworkSegmentSubnetResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadNetworkSegmentSubnetResponse)(nil)).Elem()
}

func (i *workloadNetworkSegmentSubnetResponsePtrType) ToWorkloadNetworkSegmentSubnetResponsePtrOutput() WorkloadNetworkSegmentSubnetResponsePtrOutput {
	return i.ToWorkloadNetworkSegmentSubnetResponsePtrOutputWithContext(context.Background())
}

func (i *workloadNetworkSegmentSubnetResponsePtrType) ToWorkloadNetworkSegmentSubnetResponsePtrOutputWithContext(ctx context.Context) WorkloadNetworkSegmentSubnetResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadNetworkSegmentSubnetResponsePtrOutput)
}

type WorkloadNetworkSegmentSubnetResponseOutput struct{ *pulumi.OutputState }

func (WorkloadNetworkSegmentSubnetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadNetworkSegmentSubnetResponse)(nil)).Elem()
}

func (o WorkloadNetworkSegmentSubnetResponseOutput) ToWorkloadNetworkSegmentSubnetResponseOutput() WorkloadNetworkSegmentSubnetResponseOutput {
	return o
}

func (o WorkloadNetworkSegmentSubnetResponseOutput) ToWorkloadNetworkSegmentSubnetResponseOutputWithContext(ctx context.Context) WorkloadNetworkSegmentSubnetResponseOutput {
	return o
}

func (o WorkloadNetworkSegmentSubnetResponseOutput) ToWorkloadNetworkSegmentSubnetResponsePtrOutput() WorkloadNetworkSegmentSubnetResponsePtrOutput {
	return o.ToWorkloadNetworkSegmentSubnetResponsePtrOutputWithContext(context.Background())
}

func (o WorkloadNetworkSegmentSubnetResponseOutput) ToWorkloadNetworkSegmentSubnetResponsePtrOutputWithContext(ctx context.Context) WorkloadNetworkSegmentSubnetResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkloadNetworkSegmentSubnetResponse) *WorkloadNetworkSegmentSubnetResponse {
		return &v
	}).(WorkloadNetworkSegmentSubnetResponsePtrOutput)
}

func (o WorkloadNetworkSegmentSubnetResponseOutput) DhcpRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WorkloadNetworkSegmentSubnetResponse) []string { return v.DhcpRanges }).(pulumi.StringArrayOutput)
}

func (o WorkloadNetworkSegmentSubnetResponseOutput) GatewayAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadNetworkSegmentSubnetResponse) *string { return v.GatewayAddress }).(pulumi.StringPtrOutput)
}

type WorkloadNetworkSegmentSubnetResponsePtrOutput struct{ *pulumi.OutputState }

func (WorkloadNetworkSegmentSubnetResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadNetworkSegmentSubnetResponse)(nil)).Elem()
}

func (o WorkloadNetworkSegmentSubnetResponsePtrOutput) ToWorkloadNetworkSegmentSubnetResponsePtrOutput() WorkloadNetworkSegmentSubnetResponsePtrOutput {
	return o
}

func (o WorkloadNetworkSegmentSubnetResponsePtrOutput) ToWorkloadNetworkSegmentSubnetResponsePtrOutputWithContext(ctx context.Context) WorkloadNetworkSegmentSubnetResponsePtrOutput {
	return o
}

func (o WorkloadNetworkSegmentSubnetResponsePtrOutput) Elem() WorkloadNetworkSegmentSubnetResponseOutput {
	return o.ApplyT(func(v *WorkloadNetworkSegmentSubnetResponse) WorkloadNetworkSegmentSubnetResponse {
		if v != nil {
			return *v
		}
		var ret WorkloadNetworkSegmentSubnetResponse
		return ret
	}).(WorkloadNetworkSegmentSubnetResponseOutput)
}

func (o WorkloadNetworkSegmentSubnetResponsePtrOutput) DhcpRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WorkloadNetworkSegmentSubnetResponse) []string {
		if v == nil {
			return nil
		}
		return v.DhcpRanges
	}).(pulumi.StringArrayOutput)
}

func (o WorkloadNetworkSegmentSubnetResponsePtrOutput) GatewayAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkSegmentSubnetResponse) *string {
		if v == nil {
			return nil
		}
		return v.GatewayAddress
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AddonHcxPropertiesOutput{})
	pulumi.RegisterOutputType(AddonHcxPropertiesResponseOutput{})
	pulumi.RegisterOutputType(AddonSrmPropertiesOutput{})
	pulumi.RegisterOutputType(AddonSrmPropertiesResponseOutput{})
	pulumi.RegisterOutputType(AddonVrPropertiesOutput{})
	pulumi.RegisterOutputType(AddonVrPropertiesResponseOutput{})
	pulumi.RegisterOutputType(AvailabilityPropertiesOutput{})
	pulumi.RegisterOutputType(AvailabilityPropertiesPtrOutput{})
	pulumi.RegisterOutputType(AvailabilityPropertiesResponseOutput{})
	pulumi.RegisterOutputType(AvailabilityPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(CircuitResponseOutput{})
	pulumi.RegisterOutputType(CircuitResponsePtrOutput{})
	pulumi.RegisterOutputType(DiskPoolVolumeOutput{})
	pulumi.RegisterOutputType(DiskPoolVolumePtrOutput{})
	pulumi.RegisterOutputType(DiskPoolVolumeResponseOutput{})
	pulumi.RegisterOutputType(DiskPoolVolumeResponsePtrOutput{})
	pulumi.RegisterOutputType(EncryptionOutput{})
	pulumi.RegisterOutputType(EncryptionPtrOutput{})
	pulumi.RegisterOutputType(EncryptionKeyVaultPropertiesOutput{})
	pulumi.RegisterOutputType(EncryptionKeyVaultPropertiesPtrOutput{})
	pulumi.RegisterOutputType(EncryptionKeyVaultPropertiesResponseOutput{})
	pulumi.RegisterOutputType(EncryptionKeyVaultPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(EncryptionResponseOutput{})
	pulumi.RegisterOutputType(EncryptionResponsePtrOutput{})
	pulumi.RegisterOutputType(EndpointsResponseOutput{})
	pulumi.RegisterOutputType(EndpointsResponsePtrOutput{})
	pulumi.RegisterOutputType(IdentitySourceOutput{})
	pulumi.RegisterOutputType(IdentitySourceArrayOutput{})
	pulumi.RegisterOutputType(IdentitySourceResponseOutput{})
	pulumi.RegisterOutputType(IdentitySourceResponseArrayOutput{})
	pulumi.RegisterOutputType(ManagementClusterOutput{})
	pulumi.RegisterOutputType(ManagementClusterPtrOutput{})
	pulumi.RegisterOutputType(ManagementClusterResponseOutput{})
	pulumi.RegisterOutputType(ManagementClusterResponsePtrOutput{})
	pulumi.RegisterOutputType(NetAppVolumeOutput{})
	pulumi.RegisterOutputType(NetAppVolumePtrOutput{})
	pulumi.RegisterOutputType(NetAppVolumeResponseOutput{})
	pulumi.RegisterOutputType(NetAppVolumeResponsePtrOutput{})
	pulumi.RegisterOutputType(PSCredentialExecutionParameterOutput{})
	pulumi.RegisterOutputType(PSCredentialExecutionParameterResponseOutput{})
	pulumi.RegisterOutputType(PrivateCloudIdentityOutput{})
	pulumi.RegisterOutputType(PrivateCloudIdentityPtrOutput{})
	pulumi.RegisterOutputType(PrivateCloudIdentityResponseOutput{})
	pulumi.RegisterOutputType(PrivateCloudIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(ScriptSecureStringExecutionParameterOutput{})
	pulumi.RegisterOutputType(ScriptSecureStringExecutionParameterResponseOutput{})
	pulumi.RegisterOutputType(ScriptStringExecutionParameterOutput{})
	pulumi.RegisterOutputType(ScriptStringExecutionParameterResponseOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(VmHostPlacementPolicyPropertiesOutput{})
	pulumi.RegisterOutputType(VmHostPlacementPolicyPropertiesResponseOutput{})
	pulumi.RegisterOutputType(VmVmPlacementPolicyPropertiesOutput{})
	pulumi.RegisterOutputType(VmVmPlacementPolicyPropertiesResponseOutput{})
	pulumi.RegisterOutputType(WorkloadNetworkDhcpRelayOutput{})
	pulumi.RegisterOutputType(WorkloadNetworkDhcpRelayResponseOutput{})
	pulumi.RegisterOutputType(WorkloadNetworkDhcpServerOutput{})
	pulumi.RegisterOutputType(WorkloadNetworkDhcpServerResponseOutput{})
	pulumi.RegisterOutputType(WorkloadNetworkSegmentPortVifResponseOutput{})
	pulumi.RegisterOutputType(WorkloadNetworkSegmentPortVifResponseArrayOutput{})
	pulumi.RegisterOutputType(WorkloadNetworkSegmentSubnetOutput{})
	pulumi.RegisterOutputType(WorkloadNetworkSegmentSubnetPtrOutput{})
	pulumi.RegisterOutputType(WorkloadNetworkSegmentSubnetResponseOutput{})
	pulumi.RegisterOutputType(WorkloadNetworkSegmentSubnetResponsePtrOutput{})
}
