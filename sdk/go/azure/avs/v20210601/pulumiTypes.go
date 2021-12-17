


package v20210601

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
	AddonType  string `pulumi:"addonType"`
	LicenseKey string `pulumi:"licenseKey"`
}





type AddonSrmPropertiesInput interface {
	pulumi.Input

	ToAddonSrmPropertiesOutput() AddonSrmPropertiesOutput
	ToAddonSrmPropertiesOutputWithContext(context.Context) AddonSrmPropertiesOutput
}

type AddonSrmPropertiesArgs struct {
	AddonType  pulumi.StringInput `pulumi:"addonType"`
	LicenseKey pulumi.StringInput `pulumi:"licenseKey"`
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

func (o AddonSrmPropertiesOutput) LicenseKey() pulumi.StringOutput {
	return o.ApplyT(func(v AddonSrmProperties) string { return v.LicenseKey }).(pulumi.StringOutput)
}

type AddonSrmPropertiesResponse struct {
	AddonType         string `pulumi:"addonType"`
	LicenseKey        string `pulumi:"licenseKey"`
	ProvisioningState string `pulumi:"provisioningState"`
}





type AddonSrmPropertiesResponseInput interface {
	pulumi.Input

	ToAddonSrmPropertiesResponseOutput() AddonSrmPropertiesResponseOutput
	ToAddonSrmPropertiesResponseOutputWithContext(context.Context) AddonSrmPropertiesResponseOutput
}

type AddonSrmPropertiesResponseArgs struct {
	AddonType         pulumi.StringInput `pulumi:"addonType"`
	LicenseKey        pulumi.StringInput `pulumi:"licenseKey"`
	ProvisioningState pulumi.StringInput `pulumi:"provisioningState"`
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

func (o AddonSrmPropertiesResponseOutput) LicenseKey() pulumi.StringOutput {
	return o.ApplyT(func(v AddonSrmPropertiesResponse) string { return v.LicenseKey }).(pulumi.StringOutput)
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


func (val *DiskPoolVolume) Defaults() *DiskPoolVolume {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MountOption) {
		mountOption_ := "MOUNT"
		tmp.MountOption = &mountOption_
	}
	return &tmp
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


func (val *DiskPoolVolumeResponse) Defaults() *DiskPoolVolumeResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MountOption) {
		mountOption_ := "MOUNT"
		tmp.MountOption = &mountOption_
	}
	return &tmp
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
	ClusterSize int `pulumi:"clusterSize"`
}





type ManagementClusterInput interface {
	pulumi.Input

	ToManagementClusterOutput() ManagementClusterOutput
	ToManagementClusterOutputWithContext(context.Context) ManagementClusterOutput
}

type ManagementClusterArgs struct {
	ClusterSize pulumi.IntInput `pulumi:"clusterSize"`
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
	pulumi.RegisterOutputType(CircuitResponseOutput{})
	pulumi.RegisterOutputType(CircuitResponsePtrOutput{})
	pulumi.RegisterOutputType(DiskPoolVolumeOutput{})
	pulumi.RegisterOutputType(DiskPoolVolumePtrOutput{})
	pulumi.RegisterOutputType(DiskPoolVolumeResponseOutput{})
	pulumi.RegisterOutputType(DiskPoolVolumeResponsePtrOutput{})
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
	pulumi.RegisterOutputType(ScriptSecureStringExecutionParameterOutput{})
	pulumi.RegisterOutputType(ScriptSecureStringExecutionParameterResponseOutput{})
	pulumi.RegisterOutputType(ScriptStringExecutionParameterOutput{})
	pulumi.RegisterOutputType(ScriptStringExecutionParameterResponseOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
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
