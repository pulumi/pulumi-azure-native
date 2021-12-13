


package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AutoPauseProperties struct {
	DelayInMinutes *int  `pulumi:"delayInMinutes"`
	Enabled        *bool `pulumi:"enabled"`
}





type AutoPausePropertiesInput interface {
	pulumi.Input

	ToAutoPausePropertiesOutput() AutoPausePropertiesOutput
	ToAutoPausePropertiesOutputWithContext(context.Context) AutoPausePropertiesOutput
}

type AutoPausePropertiesArgs struct {
	DelayInMinutes pulumi.IntPtrInput  `pulumi:"delayInMinutes"`
	Enabled        pulumi.BoolPtrInput `pulumi:"enabled"`
}

func (AutoPausePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoPauseProperties)(nil)).Elem()
}

func (i AutoPausePropertiesArgs) ToAutoPausePropertiesOutput() AutoPausePropertiesOutput {
	return i.ToAutoPausePropertiesOutputWithContext(context.Background())
}

func (i AutoPausePropertiesArgs) ToAutoPausePropertiesOutputWithContext(ctx context.Context) AutoPausePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoPausePropertiesOutput)
}

func (i AutoPausePropertiesArgs) ToAutoPausePropertiesPtrOutput() AutoPausePropertiesPtrOutput {
	return i.ToAutoPausePropertiesPtrOutputWithContext(context.Background())
}

func (i AutoPausePropertiesArgs) ToAutoPausePropertiesPtrOutputWithContext(ctx context.Context) AutoPausePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoPausePropertiesOutput).ToAutoPausePropertiesPtrOutputWithContext(ctx)
}









type AutoPausePropertiesPtrInput interface {
	pulumi.Input

	ToAutoPausePropertiesPtrOutput() AutoPausePropertiesPtrOutput
	ToAutoPausePropertiesPtrOutputWithContext(context.Context) AutoPausePropertiesPtrOutput
}

type autoPausePropertiesPtrType AutoPausePropertiesArgs

func AutoPausePropertiesPtr(v *AutoPausePropertiesArgs) AutoPausePropertiesPtrInput {
	return (*autoPausePropertiesPtrType)(v)
}

func (*autoPausePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoPauseProperties)(nil)).Elem()
}

func (i *autoPausePropertiesPtrType) ToAutoPausePropertiesPtrOutput() AutoPausePropertiesPtrOutput {
	return i.ToAutoPausePropertiesPtrOutputWithContext(context.Background())
}

func (i *autoPausePropertiesPtrType) ToAutoPausePropertiesPtrOutputWithContext(ctx context.Context) AutoPausePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoPausePropertiesPtrOutput)
}

type AutoPausePropertiesOutput struct{ *pulumi.OutputState }

func (AutoPausePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoPauseProperties)(nil)).Elem()
}

func (o AutoPausePropertiesOutput) ToAutoPausePropertiesOutput() AutoPausePropertiesOutput {
	return o
}

func (o AutoPausePropertiesOutput) ToAutoPausePropertiesOutputWithContext(ctx context.Context) AutoPausePropertiesOutput {
	return o
}

func (o AutoPausePropertiesOutput) ToAutoPausePropertiesPtrOutput() AutoPausePropertiesPtrOutput {
	return o.ToAutoPausePropertiesPtrOutputWithContext(context.Background())
}

func (o AutoPausePropertiesOutput) ToAutoPausePropertiesPtrOutputWithContext(ctx context.Context) AutoPausePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoPauseProperties) *AutoPauseProperties {
		return &v
	}).(AutoPausePropertiesPtrOutput)
}

func (o AutoPausePropertiesOutput) DelayInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoPauseProperties) *int { return v.DelayInMinutes }).(pulumi.IntPtrOutput)
}

func (o AutoPausePropertiesOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AutoPauseProperties) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

type AutoPausePropertiesPtrOutput struct{ *pulumi.OutputState }

func (AutoPausePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoPauseProperties)(nil)).Elem()
}

func (o AutoPausePropertiesPtrOutput) ToAutoPausePropertiesPtrOutput() AutoPausePropertiesPtrOutput {
	return o
}

func (o AutoPausePropertiesPtrOutput) ToAutoPausePropertiesPtrOutputWithContext(ctx context.Context) AutoPausePropertiesPtrOutput {
	return o
}

func (o AutoPausePropertiesPtrOutput) Elem() AutoPausePropertiesOutput {
	return o.ApplyT(func(v *AutoPauseProperties) AutoPauseProperties {
		if v != nil {
			return *v
		}
		var ret AutoPauseProperties
		return ret
	}).(AutoPausePropertiesOutput)
}

func (o AutoPausePropertiesPtrOutput) DelayInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoPauseProperties) *int {
		if v == nil {
			return nil
		}
		return v.DelayInMinutes
	}).(pulumi.IntPtrOutput)
}

func (o AutoPausePropertiesPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AutoPauseProperties) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

type AutoPausePropertiesResponse struct {
	DelayInMinutes *int  `pulumi:"delayInMinutes"`
	Enabled        *bool `pulumi:"enabled"`
}

type AutoPausePropertiesResponseOutput struct{ *pulumi.OutputState }

func (AutoPausePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoPausePropertiesResponse)(nil)).Elem()
}

func (o AutoPausePropertiesResponseOutput) ToAutoPausePropertiesResponseOutput() AutoPausePropertiesResponseOutput {
	return o
}

func (o AutoPausePropertiesResponseOutput) ToAutoPausePropertiesResponseOutputWithContext(ctx context.Context) AutoPausePropertiesResponseOutput {
	return o
}

func (o AutoPausePropertiesResponseOutput) DelayInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoPausePropertiesResponse) *int { return v.DelayInMinutes }).(pulumi.IntPtrOutput)
}

func (o AutoPausePropertiesResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AutoPausePropertiesResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

type AutoPausePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (AutoPausePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoPausePropertiesResponse)(nil)).Elem()
}

func (o AutoPausePropertiesResponsePtrOutput) ToAutoPausePropertiesResponsePtrOutput() AutoPausePropertiesResponsePtrOutput {
	return o
}

func (o AutoPausePropertiesResponsePtrOutput) ToAutoPausePropertiesResponsePtrOutputWithContext(ctx context.Context) AutoPausePropertiesResponsePtrOutput {
	return o
}

func (o AutoPausePropertiesResponsePtrOutput) Elem() AutoPausePropertiesResponseOutput {
	return o.ApplyT(func(v *AutoPausePropertiesResponse) AutoPausePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret AutoPausePropertiesResponse
		return ret
	}).(AutoPausePropertiesResponseOutput)
}

func (o AutoPausePropertiesResponsePtrOutput) DelayInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoPausePropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return v.DelayInMinutes
	}).(pulumi.IntPtrOutput)
}

func (o AutoPausePropertiesResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AutoPausePropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

type AutoScaleProperties struct {
	Enabled      *bool `pulumi:"enabled"`
	MaxNodeCount *int  `pulumi:"maxNodeCount"`
	MinNodeCount *int  `pulumi:"minNodeCount"`
}





type AutoScalePropertiesInput interface {
	pulumi.Input

	ToAutoScalePropertiesOutput() AutoScalePropertiesOutput
	ToAutoScalePropertiesOutputWithContext(context.Context) AutoScalePropertiesOutput
}

type AutoScalePropertiesArgs struct {
	Enabled      pulumi.BoolPtrInput `pulumi:"enabled"`
	MaxNodeCount pulumi.IntPtrInput  `pulumi:"maxNodeCount"`
	MinNodeCount pulumi.IntPtrInput  `pulumi:"minNodeCount"`
}

func (AutoScalePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoScaleProperties)(nil)).Elem()
}

func (i AutoScalePropertiesArgs) ToAutoScalePropertiesOutput() AutoScalePropertiesOutput {
	return i.ToAutoScalePropertiesOutputWithContext(context.Background())
}

func (i AutoScalePropertiesArgs) ToAutoScalePropertiesOutputWithContext(ctx context.Context) AutoScalePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoScalePropertiesOutput)
}

func (i AutoScalePropertiesArgs) ToAutoScalePropertiesPtrOutput() AutoScalePropertiesPtrOutput {
	return i.ToAutoScalePropertiesPtrOutputWithContext(context.Background())
}

func (i AutoScalePropertiesArgs) ToAutoScalePropertiesPtrOutputWithContext(ctx context.Context) AutoScalePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoScalePropertiesOutput).ToAutoScalePropertiesPtrOutputWithContext(ctx)
}









type AutoScalePropertiesPtrInput interface {
	pulumi.Input

	ToAutoScalePropertiesPtrOutput() AutoScalePropertiesPtrOutput
	ToAutoScalePropertiesPtrOutputWithContext(context.Context) AutoScalePropertiesPtrOutput
}

type autoScalePropertiesPtrType AutoScalePropertiesArgs

func AutoScalePropertiesPtr(v *AutoScalePropertiesArgs) AutoScalePropertiesPtrInput {
	return (*autoScalePropertiesPtrType)(v)
}

func (*autoScalePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoScaleProperties)(nil)).Elem()
}

func (i *autoScalePropertiesPtrType) ToAutoScalePropertiesPtrOutput() AutoScalePropertiesPtrOutput {
	return i.ToAutoScalePropertiesPtrOutputWithContext(context.Background())
}

func (i *autoScalePropertiesPtrType) ToAutoScalePropertiesPtrOutputWithContext(ctx context.Context) AutoScalePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoScalePropertiesPtrOutput)
}

type AutoScalePropertiesOutput struct{ *pulumi.OutputState }

func (AutoScalePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoScaleProperties)(nil)).Elem()
}

func (o AutoScalePropertiesOutput) ToAutoScalePropertiesOutput() AutoScalePropertiesOutput {
	return o
}

func (o AutoScalePropertiesOutput) ToAutoScalePropertiesOutputWithContext(ctx context.Context) AutoScalePropertiesOutput {
	return o
}

func (o AutoScalePropertiesOutput) ToAutoScalePropertiesPtrOutput() AutoScalePropertiesPtrOutput {
	return o.ToAutoScalePropertiesPtrOutputWithContext(context.Background())
}

func (o AutoScalePropertiesOutput) ToAutoScalePropertiesPtrOutputWithContext(ctx context.Context) AutoScalePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoScaleProperties) *AutoScaleProperties {
		return &v
	}).(AutoScalePropertiesPtrOutput)
}

func (o AutoScalePropertiesOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AutoScaleProperties) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o AutoScalePropertiesOutput) MaxNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoScaleProperties) *int { return v.MaxNodeCount }).(pulumi.IntPtrOutput)
}

func (o AutoScalePropertiesOutput) MinNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoScaleProperties) *int { return v.MinNodeCount }).(pulumi.IntPtrOutput)
}

type AutoScalePropertiesPtrOutput struct{ *pulumi.OutputState }

func (AutoScalePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoScaleProperties)(nil)).Elem()
}

func (o AutoScalePropertiesPtrOutput) ToAutoScalePropertiesPtrOutput() AutoScalePropertiesPtrOutput {
	return o
}

func (o AutoScalePropertiesPtrOutput) ToAutoScalePropertiesPtrOutputWithContext(ctx context.Context) AutoScalePropertiesPtrOutput {
	return o
}

func (o AutoScalePropertiesPtrOutput) Elem() AutoScalePropertiesOutput {
	return o.ApplyT(func(v *AutoScaleProperties) AutoScaleProperties {
		if v != nil {
			return *v
		}
		var ret AutoScaleProperties
		return ret
	}).(AutoScalePropertiesOutput)
}

func (o AutoScalePropertiesPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AutoScaleProperties) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o AutoScalePropertiesPtrOutput) MaxNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoScaleProperties) *int {
		if v == nil {
			return nil
		}
		return v.MaxNodeCount
	}).(pulumi.IntPtrOutput)
}

func (o AutoScalePropertiesPtrOutput) MinNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoScaleProperties) *int {
		if v == nil {
			return nil
		}
		return v.MinNodeCount
	}).(pulumi.IntPtrOutput)
}

type AutoScalePropertiesResponse struct {
	Enabled      *bool `pulumi:"enabled"`
	MaxNodeCount *int  `pulumi:"maxNodeCount"`
	MinNodeCount *int  `pulumi:"minNodeCount"`
}

type AutoScalePropertiesResponseOutput struct{ *pulumi.OutputState }

func (AutoScalePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoScalePropertiesResponse)(nil)).Elem()
}

func (o AutoScalePropertiesResponseOutput) ToAutoScalePropertiesResponseOutput() AutoScalePropertiesResponseOutput {
	return o
}

func (o AutoScalePropertiesResponseOutput) ToAutoScalePropertiesResponseOutputWithContext(ctx context.Context) AutoScalePropertiesResponseOutput {
	return o
}

func (o AutoScalePropertiesResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AutoScalePropertiesResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o AutoScalePropertiesResponseOutput) MaxNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoScalePropertiesResponse) *int { return v.MaxNodeCount }).(pulumi.IntPtrOutput)
}

func (o AutoScalePropertiesResponseOutput) MinNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoScalePropertiesResponse) *int { return v.MinNodeCount }).(pulumi.IntPtrOutput)
}

type AutoScalePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (AutoScalePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoScalePropertiesResponse)(nil)).Elem()
}

func (o AutoScalePropertiesResponsePtrOutput) ToAutoScalePropertiesResponsePtrOutput() AutoScalePropertiesResponsePtrOutput {
	return o
}

func (o AutoScalePropertiesResponsePtrOutput) ToAutoScalePropertiesResponsePtrOutputWithContext(ctx context.Context) AutoScalePropertiesResponsePtrOutput {
	return o
}

func (o AutoScalePropertiesResponsePtrOutput) Elem() AutoScalePropertiesResponseOutput {
	return o.ApplyT(func(v *AutoScalePropertiesResponse) AutoScalePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret AutoScalePropertiesResponse
		return ret
	}).(AutoScalePropertiesResponseOutput)
}

func (o AutoScalePropertiesResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AutoScalePropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o AutoScalePropertiesResponsePtrOutput) MaxNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoScalePropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxNodeCount
	}).(pulumi.IntPtrOutput)
}

func (o AutoScalePropertiesResponsePtrOutput) MinNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoScalePropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return v.MinNodeCount
	}).(pulumi.IntPtrOutput)
}

type AzureSku struct {
	Capacity *int   `pulumi:"capacity"`
	Name     string `pulumi:"name"`
	Size     string `pulumi:"size"`
}





type AzureSkuInput interface {
	pulumi.Input

	ToAzureSkuOutput() AzureSkuOutput
	ToAzureSkuOutputWithContext(context.Context) AzureSkuOutput
}

type AzureSkuArgs struct {
	Capacity pulumi.IntPtrInput `pulumi:"capacity"`
	Name     pulumi.StringInput `pulumi:"name"`
	Size     pulumi.StringInput `pulumi:"size"`
}

func (AzureSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSku)(nil)).Elem()
}

func (i AzureSkuArgs) ToAzureSkuOutput() AzureSkuOutput {
	return i.ToAzureSkuOutputWithContext(context.Background())
}

func (i AzureSkuArgs) ToAzureSkuOutputWithContext(ctx context.Context) AzureSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureSkuOutput)
}

type AzureSkuOutput struct{ *pulumi.OutputState }

func (AzureSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSku)(nil)).Elem()
}

func (o AzureSkuOutput) ToAzureSkuOutput() AzureSkuOutput {
	return o
}

func (o AzureSkuOutput) ToAzureSkuOutputWithContext(ctx context.Context) AzureSkuOutput {
	return o
}

func (o AzureSkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureSku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o AzureSkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AzureSku) string { return v.Name }).(pulumi.StringOutput)
}

func (o AzureSkuOutput) Size() pulumi.StringOutput {
	return o.ApplyT(func(v AzureSku) string { return v.Size }).(pulumi.StringOutput)
}

type AzureSkuResponse struct {
	Capacity *int   `pulumi:"capacity"`
	Name     string `pulumi:"name"`
	Size     string `pulumi:"size"`
}

type AzureSkuResponseOutput struct{ *pulumi.OutputState }

func (AzureSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSkuResponse)(nil)).Elem()
}

func (o AzureSkuResponseOutput) ToAzureSkuResponseOutput() AzureSkuResponseOutput {
	return o
}

func (o AzureSkuResponseOutput) ToAzureSkuResponseOutputWithContext(ctx context.Context) AzureSkuResponseOutput {
	return o
}

func (o AzureSkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureSkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o AzureSkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AzureSkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o AzureSkuResponseOutput) Size() pulumi.StringOutput {
	return o.ApplyT(func(v AzureSkuResponse) string { return v.Size }).(pulumi.StringOutput)
}

type CmdkeySetup struct {
	Password   SecureString `pulumi:"password"`
	TargetName interface{}  `pulumi:"targetName"`
	Type       string       `pulumi:"type"`
	UserName   interface{}  `pulumi:"userName"`
}

type CmdkeySetupResponse struct {
	Password   SecureStringResponse `pulumi:"password"`
	TargetName interface{}          `pulumi:"targetName"`
	Type       string               `pulumi:"type"`
	UserName   interface{}          `pulumi:"userName"`
}

type ComponentSetup struct {
	ComponentName string        `pulumi:"componentName"`
	LicenseKey    *SecureString `pulumi:"licenseKey"`
	Type          string        `pulumi:"type"`
}

type ComponentSetupResponse struct {
	ComponentName string                `pulumi:"componentName"`
	LicenseKey    *SecureStringResponse `pulumi:"licenseKey"`
	Type          string                `pulumi:"type"`
}

type CspWorkspaceAdminProperties struct {
	InitialWorkspaceAdminObjectId *string `pulumi:"initialWorkspaceAdminObjectId"`
}





type CspWorkspaceAdminPropertiesInput interface {
	pulumi.Input

	ToCspWorkspaceAdminPropertiesOutput() CspWorkspaceAdminPropertiesOutput
	ToCspWorkspaceAdminPropertiesOutputWithContext(context.Context) CspWorkspaceAdminPropertiesOutput
}

type CspWorkspaceAdminPropertiesArgs struct {
	InitialWorkspaceAdminObjectId pulumi.StringPtrInput `pulumi:"initialWorkspaceAdminObjectId"`
}

func (CspWorkspaceAdminPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CspWorkspaceAdminProperties)(nil)).Elem()
}

func (i CspWorkspaceAdminPropertiesArgs) ToCspWorkspaceAdminPropertiesOutput() CspWorkspaceAdminPropertiesOutput {
	return i.ToCspWorkspaceAdminPropertiesOutputWithContext(context.Background())
}

func (i CspWorkspaceAdminPropertiesArgs) ToCspWorkspaceAdminPropertiesOutputWithContext(ctx context.Context) CspWorkspaceAdminPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CspWorkspaceAdminPropertiesOutput)
}

func (i CspWorkspaceAdminPropertiesArgs) ToCspWorkspaceAdminPropertiesPtrOutput() CspWorkspaceAdminPropertiesPtrOutput {
	return i.ToCspWorkspaceAdminPropertiesPtrOutputWithContext(context.Background())
}

func (i CspWorkspaceAdminPropertiesArgs) ToCspWorkspaceAdminPropertiesPtrOutputWithContext(ctx context.Context) CspWorkspaceAdminPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CspWorkspaceAdminPropertiesOutput).ToCspWorkspaceAdminPropertiesPtrOutputWithContext(ctx)
}









type CspWorkspaceAdminPropertiesPtrInput interface {
	pulumi.Input

	ToCspWorkspaceAdminPropertiesPtrOutput() CspWorkspaceAdminPropertiesPtrOutput
	ToCspWorkspaceAdminPropertiesPtrOutputWithContext(context.Context) CspWorkspaceAdminPropertiesPtrOutput
}

type cspWorkspaceAdminPropertiesPtrType CspWorkspaceAdminPropertiesArgs

func CspWorkspaceAdminPropertiesPtr(v *CspWorkspaceAdminPropertiesArgs) CspWorkspaceAdminPropertiesPtrInput {
	return (*cspWorkspaceAdminPropertiesPtrType)(v)
}

func (*cspWorkspaceAdminPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CspWorkspaceAdminProperties)(nil)).Elem()
}

func (i *cspWorkspaceAdminPropertiesPtrType) ToCspWorkspaceAdminPropertiesPtrOutput() CspWorkspaceAdminPropertiesPtrOutput {
	return i.ToCspWorkspaceAdminPropertiesPtrOutputWithContext(context.Background())
}

func (i *cspWorkspaceAdminPropertiesPtrType) ToCspWorkspaceAdminPropertiesPtrOutputWithContext(ctx context.Context) CspWorkspaceAdminPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CspWorkspaceAdminPropertiesPtrOutput)
}

type CspWorkspaceAdminPropertiesOutput struct{ *pulumi.OutputState }

func (CspWorkspaceAdminPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CspWorkspaceAdminProperties)(nil)).Elem()
}

func (o CspWorkspaceAdminPropertiesOutput) ToCspWorkspaceAdminPropertiesOutput() CspWorkspaceAdminPropertiesOutput {
	return o
}

func (o CspWorkspaceAdminPropertiesOutput) ToCspWorkspaceAdminPropertiesOutputWithContext(ctx context.Context) CspWorkspaceAdminPropertiesOutput {
	return o
}

func (o CspWorkspaceAdminPropertiesOutput) ToCspWorkspaceAdminPropertiesPtrOutput() CspWorkspaceAdminPropertiesPtrOutput {
	return o.ToCspWorkspaceAdminPropertiesPtrOutputWithContext(context.Background())
}

func (o CspWorkspaceAdminPropertiesOutput) ToCspWorkspaceAdminPropertiesPtrOutputWithContext(ctx context.Context) CspWorkspaceAdminPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CspWorkspaceAdminProperties) *CspWorkspaceAdminProperties {
		return &v
	}).(CspWorkspaceAdminPropertiesPtrOutput)
}

func (o CspWorkspaceAdminPropertiesOutput) InitialWorkspaceAdminObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CspWorkspaceAdminProperties) *string { return v.InitialWorkspaceAdminObjectId }).(pulumi.StringPtrOutput)
}

type CspWorkspaceAdminPropertiesPtrOutput struct{ *pulumi.OutputState }

func (CspWorkspaceAdminPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CspWorkspaceAdminProperties)(nil)).Elem()
}

func (o CspWorkspaceAdminPropertiesPtrOutput) ToCspWorkspaceAdminPropertiesPtrOutput() CspWorkspaceAdminPropertiesPtrOutput {
	return o
}

func (o CspWorkspaceAdminPropertiesPtrOutput) ToCspWorkspaceAdminPropertiesPtrOutputWithContext(ctx context.Context) CspWorkspaceAdminPropertiesPtrOutput {
	return o
}

func (o CspWorkspaceAdminPropertiesPtrOutput) Elem() CspWorkspaceAdminPropertiesOutput {
	return o.ApplyT(func(v *CspWorkspaceAdminProperties) CspWorkspaceAdminProperties {
		if v != nil {
			return *v
		}
		var ret CspWorkspaceAdminProperties
		return ret
	}).(CspWorkspaceAdminPropertiesOutput)
}

func (o CspWorkspaceAdminPropertiesPtrOutput) InitialWorkspaceAdminObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CspWorkspaceAdminProperties) *string {
		if v == nil {
			return nil
		}
		return v.InitialWorkspaceAdminObjectId
	}).(pulumi.StringPtrOutput)
}

type CspWorkspaceAdminPropertiesResponse struct {
	InitialWorkspaceAdminObjectId *string `pulumi:"initialWorkspaceAdminObjectId"`
}

type CspWorkspaceAdminPropertiesResponseOutput struct{ *pulumi.OutputState }

func (CspWorkspaceAdminPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CspWorkspaceAdminPropertiesResponse)(nil)).Elem()
}

func (o CspWorkspaceAdminPropertiesResponseOutput) ToCspWorkspaceAdminPropertiesResponseOutput() CspWorkspaceAdminPropertiesResponseOutput {
	return o
}

func (o CspWorkspaceAdminPropertiesResponseOutput) ToCspWorkspaceAdminPropertiesResponseOutputWithContext(ctx context.Context) CspWorkspaceAdminPropertiesResponseOutput {
	return o
}

func (o CspWorkspaceAdminPropertiesResponseOutput) InitialWorkspaceAdminObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CspWorkspaceAdminPropertiesResponse) *string { return v.InitialWorkspaceAdminObjectId }).(pulumi.StringPtrOutput)
}

type CspWorkspaceAdminPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (CspWorkspaceAdminPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CspWorkspaceAdminPropertiesResponse)(nil)).Elem()
}

func (o CspWorkspaceAdminPropertiesResponsePtrOutput) ToCspWorkspaceAdminPropertiesResponsePtrOutput() CspWorkspaceAdminPropertiesResponsePtrOutput {
	return o
}

func (o CspWorkspaceAdminPropertiesResponsePtrOutput) ToCspWorkspaceAdminPropertiesResponsePtrOutputWithContext(ctx context.Context) CspWorkspaceAdminPropertiesResponsePtrOutput {
	return o
}

func (o CspWorkspaceAdminPropertiesResponsePtrOutput) Elem() CspWorkspaceAdminPropertiesResponseOutput {
	return o.ApplyT(func(v *CspWorkspaceAdminPropertiesResponse) CspWorkspaceAdminPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret CspWorkspaceAdminPropertiesResponse
		return ret
	}).(CspWorkspaceAdminPropertiesResponseOutput)
}

func (o CspWorkspaceAdminPropertiesResponsePtrOutput) InitialWorkspaceAdminObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CspWorkspaceAdminPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.InitialWorkspaceAdminObjectId
	}).(pulumi.StringPtrOutput)
}

type CustomerManagedKeyDetails struct {
	KekIdentity *KekIdentityProperties `pulumi:"kekIdentity"`
	Key         *WorkspaceKeyDetails   `pulumi:"key"`
}





type CustomerManagedKeyDetailsInput interface {
	pulumi.Input

	ToCustomerManagedKeyDetailsOutput() CustomerManagedKeyDetailsOutput
	ToCustomerManagedKeyDetailsOutputWithContext(context.Context) CustomerManagedKeyDetailsOutput
}

type CustomerManagedKeyDetailsArgs struct {
	KekIdentity KekIdentityPropertiesPtrInput `pulumi:"kekIdentity"`
	Key         WorkspaceKeyDetailsPtrInput   `pulumi:"key"`
}

func (CustomerManagedKeyDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomerManagedKeyDetails)(nil)).Elem()
}

func (i CustomerManagedKeyDetailsArgs) ToCustomerManagedKeyDetailsOutput() CustomerManagedKeyDetailsOutput {
	return i.ToCustomerManagedKeyDetailsOutputWithContext(context.Background())
}

func (i CustomerManagedKeyDetailsArgs) ToCustomerManagedKeyDetailsOutputWithContext(ctx context.Context) CustomerManagedKeyDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerManagedKeyDetailsOutput)
}

func (i CustomerManagedKeyDetailsArgs) ToCustomerManagedKeyDetailsPtrOutput() CustomerManagedKeyDetailsPtrOutput {
	return i.ToCustomerManagedKeyDetailsPtrOutputWithContext(context.Background())
}

func (i CustomerManagedKeyDetailsArgs) ToCustomerManagedKeyDetailsPtrOutputWithContext(ctx context.Context) CustomerManagedKeyDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerManagedKeyDetailsOutput).ToCustomerManagedKeyDetailsPtrOutputWithContext(ctx)
}









type CustomerManagedKeyDetailsPtrInput interface {
	pulumi.Input

	ToCustomerManagedKeyDetailsPtrOutput() CustomerManagedKeyDetailsPtrOutput
	ToCustomerManagedKeyDetailsPtrOutputWithContext(context.Context) CustomerManagedKeyDetailsPtrOutput
}

type customerManagedKeyDetailsPtrType CustomerManagedKeyDetailsArgs

func CustomerManagedKeyDetailsPtr(v *CustomerManagedKeyDetailsArgs) CustomerManagedKeyDetailsPtrInput {
	return (*customerManagedKeyDetailsPtrType)(v)
}

func (*customerManagedKeyDetailsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomerManagedKeyDetails)(nil)).Elem()
}

func (i *customerManagedKeyDetailsPtrType) ToCustomerManagedKeyDetailsPtrOutput() CustomerManagedKeyDetailsPtrOutput {
	return i.ToCustomerManagedKeyDetailsPtrOutputWithContext(context.Background())
}

func (i *customerManagedKeyDetailsPtrType) ToCustomerManagedKeyDetailsPtrOutputWithContext(ctx context.Context) CustomerManagedKeyDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerManagedKeyDetailsPtrOutput)
}

type CustomerManagedKeyDetailsOutput struct{ *pulumi.OutputState }

func (CustomerManagedKeyDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomerManagedKeyDetails)(nil)).Elem()
}

func (o CustomerManagedKeyDetailsOutput) ToCustomerManagedKeyDetailsOutput() CustomerManagedKeyDetailsOutput {
	return o
}

func (o CustomerManagedKeyDetailsOutput) ToCustomerManagedKeyDetailsOutputWithContext(ctx context.Context) CustomerManagedKeyDetailsOutput {
	return o
}

func (o CustomerManagedKeyDetailsOutput) ToCustomerManagedKeyDetailsPtrOutput() CustomerManagedKeyDetailsPtrOutput {
	return o.ToCustomerManagedKeyDetailsPtrOutputWithContext(context.Background())
}

func (o CustomerManagedKeyDetailsOutput) ToCustomerManagedKeyDetailsPtrOutputWithContext(ctx context.Context) CustomerManagedKeyDetailsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomerManagedKeyDetails) *CustomerManagedKeyDetails {
		return &v
	}).(CustomerManagedKeyDetailsPtrOutput)
}

func (o CustomerManagedKeyDetailsOutput) KekIdentity() KekIdentityPropertiesPtrOutput {
	return o.ApplyT(func(v CustomerManagedKeyDetails) *KekIdentityProperties { return v.KekIdentity }).(KekIdentityPropertiesPtrOutput)
}

func (o CustomerManagedKeyDetailsOutput) Key() WorkspaceKeyDetailsPtrOutput {
	return o.ApplyT(func(v CustomerManagedKeyDetails) *WorkspaceKeyDetails { return v.Key }).(WorkspaceKeyDetailsPtrOutput)
}

type CustomerManagedKeyDetailsPtrOutput struct{ *pulumi.OutputState }

func (CustomerManagedKeyDetailsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomerManagedKeyDetails)(nil)).Elem()
}

func (o CustomerManagedKeyDetailsPtrOutput) ToCustomerManagedKeyDetailsPtrOutput() CustomerManagedKeyDetailsPtrOutput {
	return o
}

func (o CustomerManagedKeyDetailsPtrOutput) ToCustomerManagedKeyDetailsPtrOutputWithContext(ctx context.Context) CustomerManagedKeyDetailsPtrOutput {
	return o
}

func (o CustomerManagedKeyDetailsPtrOutput) Elem() CustomerManagedKeyDetailsOutput {
	return o.ApplyT(func(v *CustomerManagedKeyDetails) CustomerManagedKeyDetails {
		if v != nil {
			return *v
		}
		var ret CustomerManagedKeyDetails
		return ret
	}).(CustomerManagedKeyDetailsOutput)
}

func (o CustomerManagedKeyDetailsPtrOutput) KekIdentity() KekIdentityPropertiesPtrOutput {
	return o.ApplyT(func(v *CustomerManagedKeyDetails) *KekIdentityProperties {
		if v == nil {
			return nil
		}
		return v.KekIdentity
	}).(KekIdentityPropertiesPtrOutput)
}

func (o CustomerManagedKeyDetailsPtrOutput) Key() WorkspaceKeyDetailsPtrOutput {
	return o.ApplyT(func(v *CustomerManagedKeyDetails) *WorkspaceKeyDetails {
		if v == nil {
			return nil
		}
		return v.Key
	}).(WorkspaceKeyDetailsPtrOutput)
}

type CustomerManagedKeyDetailsResponse struct {
	KekIdentity *KekIdentityPropertiesResponse `pulumi:"kekIdentity"`
	Key         *WorkspaceKeyDetailsResponse   `pulumi:"key"`
	Status      string                         `pulumi:"status"`
}

type CustomerManagedKeyDetailsResponseOutput struct{ *pulumi.OutputState }

func (CustomerManagedKeyDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomerManagedKeyDetailsResponse)(nil)).Elem()
}

func (o CustomerManagedKeyDetailsResponseOutput) ToCustomerManagedKeyDetailsResponseOutput() CustomerManagedKeyDetailsResponseOutput {
	return o
}

func (o CustomerManagedKeyDetailsResponseOutput) ToCustomerManagedKeyDetailsResponseOutputWithContext(ctx context.Context) CustomerManagedKeyDetailsResponseOutput {
	return o
}

func (o CustomerManagedKeyDetailsResponseOutput) KekIdentity() KekIdentityPropertiesResponsePtrOutput {
	return o.ApplyT(func(v CustomerManagedKeyDetailsResponse) *KekIdentityPropertiesResponse { return v.KekIdentity }).(KekIdentityPropertiesResponsePtrOutput)
}

func (o CustomerManagedKeyDetailsResponseOutput) Key() WorkspaceKeyDetailsResponsePtrOutput {
	return o.ApplyT(func(v CustomerManagedKeyDetailsResponse) *WorkspaceKeyDetailsResponse { return v.Key }).(WorkspaceKeyDetailsResponsePtrOutput)
}

func (o CustomerManagedKeyDetailsResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v CustomerManagedKeyDetailsResponse) string { return v.Status }).(pulumi.StringOutput)
}

type CustomerManagedKeyDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (CustomerManagedKeyDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomerManagedKeyDetailsResponse)(nil)).Elem()
}

func (o CustomerManagedKeyDetailsResponsePtrOutput) ToCustomerManagedKeyDetailsResponsePtrOutput() CustomerManagedKeyDetailsResponsePtrOutput {
	return o
}

func (o CustomerManagedKeyDetailsResponsePtrOutput) ToCustomerManagedKeyDetailsResponsePtrOutputWithContext(ctx context.Context) CustomerManagedKeyDetailsResponsePtrOutput {
	return o
}

func (o CustomerManagedKeyDetailsResponsePtrOutput) Elem() CustomerManagedKeyDetailsResponseOutput {
	return o.ApplyT(func(v *CustomerManagedKeyDetailsResponse) CustomerManagedKeyDetailsResponse {
		if v != nil {
			return *v
		}
		var ret CustomerManagedKeyDetailsResponse
		return ret
	}).(CustomerManagedKeyDetailsResponseOutput)
}

func (o CustomerManagedKeyDetailsResponsePtrOutput) KekIdentity() KekIdentityPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *CustomerManagedKeyDetailsResponse) *KekIdentityPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.KekIdentity
	}).(KekIdentityPropertiesResponsePtrOutput)
}

func (o CustomerManagedKeyDetailsResponsePtrOutput) Key() WorkspaceKeyDetailsResponsePtrOutput {
	return o.ApplyT(func(v *CustomerManagedKeyDetailsResponse) *WorkspaceKeyDetailsResponse {
		if v == nil {
			return nil
		}
		return v.Key
	}).(WorkspaceKeyDetailsResponsePtrOutput)
}

func (o CustomerManagedKeyDetailsResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomerManagedKeyDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

type DataLakeStorageAccountDetails struct {
	AccountUrl                   *string `pulumi:"accountUrl"`
	CreateManagedPrivateEndpoint *bool   `pulumi:"createManagedPrivateEndpoint"`
	Filesystem                   *string `pulumi:"filesystem"`
	ResourceId                   *string `pulumi:"resourceId"`
}





type DataLakeStorageAccountDetailsInput interface {
	pulumi.Input

	ToDataLakeStorageAccountDetailsOutput() DataLakeStorageAccountDetailsOutput
	ToDataLakeStorageAccountDetailsOutputWithContext(context.Context) DataLakeStorageAccountDetailsOutput
}

type DataLakeStorageAccountDetailsArgs struct {
	AccountUrl                   pulumi.StringPtrInput `pulumi:"accountUrl"`
	CreateManagedPrivateEndpoint pulumi.BoolPtrInput   `pulumi:"createManagedPrivateEndpoint"`
	Filesystem                   pulumi.StringPtrInput `pulumi:"filesystem"`
	ResourceId                   pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (DataLakeStorageAccountDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataLakeStorageAccountDetails)(nil)).Elem()
}

func (i DataLakeStorageAccountDetailsArgs) ToDataLakeStorageAccountDetailsOutput() DataLakeStorageAccountDetailsOutput {
	return i.ToDataLakeStorageAccountDetailsOutputWithContext(context.Background())
}

func (i DataLakeStorageAccountDetailsArgs) ToDataLakeStorageAccountDetailsOutputWithContext(ctx context.Context) DataLakeStorageAccountDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataLakeStorageAccountDetailsOutput)
}

func (i DataLakeStorageAccountDetailsArgs) ToDataLakeStorageAccountDetailsPtrOutput() DataLakeStorageAccountDetailsPtrOutput {
	return i.ToDataLakeStorageAccountDetailsPtrOutputWithContext(context.Background())
}

func (i DataLakeStorageAccountDetailsArgs) ToDataLakeStorageAccountDetailsPtrOutputWithContext(ctx context.Context) DataLakeStorageAccountDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataLakeStorageAccountDetailsOutput).ToDataLakeStorageAccountDetailsPtrOutputWithContext(ctx)
}









type DataLakeStorageAccountDetailsPtrInput interface {
	pulumi.Input

	ToDataLakeStorageAccountDetailsPtrOutput() DataLakeStorageAccountDetailsPtrOutput
	ToDataLakeStorageAccountDetailsPtrOutputWithContext(context.Context) DataLakeStorageAccountDetailsPtrOutput
}

type dataLakeStorageAccountDetailsPtrType DataLakeStorageAccountDetailsArgs

func DataLakeStorageAccountDetailsPtr(v *DataLakeStorageAccountDetailsArgs) DataLakeStorageAccountDetailsPtrInput {
	return (*dataLakeStorageAccountDetailsPtrType)(v)
}

func (*dataLakeStorageAccountDetailsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DataLakeStorageAccountDetails)(nil)).Elem()
}

func (i *dataLakeStorageAccountDetailsPtrType) ToDataLakeStorageAccountDetailsPtrOutput() DataLakeStorageAccountDetailsPtrOutput {
	return i.ToDataLakeStorageAccountDetailsPtrOutputWithContext(context.Background())
}

func (i *dataLakeStorageAccountDetailsPtrType) ToDataLakeStorageAccountDetailsPtrOutputWithContext(ctx context.Context) DataLakeStorageAccountDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataLakeStorageAccountDetailsPtrOutput)
}

type DataLakeStorageAccountDetailsOutput struct{ *pulumi.OutputState }

func (DataLakeStorageAccountDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataLakeStorageAccountDetails)(nil)).Elem()
}

func (o DataLakeStorageAccountDetailsOutput) ToDataLakeStorageAccountDetailsOutput() DataLakeStorageAccountDetailsOutput {
	return o
}

func (o DataLakeStorageAccountDetailsOutput) ToDataLakeStorageAccountDetailsOutputWithContext(ctx context.Context) DataLakeStorageAccountDetailsOutput {
	return o
}

func (o DataLakeStorageAccountDetailsOutput) ToDataLakeStorageAccountDetailsPtrOutput() DataLakeStorageAccountDetailsPtrOutput {
	return o.ToDataLakeStorageAccountDetailsPtrOutputWithContext(context.Background())
}

func (o DataLakeStorageAccountDetailsOutput) ToDataLakeStorageAccountDetailsPtrOutputWithContext(ctx context.Context) DataLakeStorageAccountDetailsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataLakeStorageAccountDetails) *DataLakeStorageAccountDetails {
		return &v
	}).(DataLakeStorageAccountDetailsPtrOutput)
}

func (o DataLakeStorageAccountDetailsOutput) AccountUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataLakeStorageAccountDetails) *string { return v.AccountUrl }).(pulumi.StringPtrOutput)
}

func (o DataLakeStorageAccountDetailsOutput) CreateManagedPrivateEndpoint() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DataLakeStorageAccountDetails) *bool { return v.CreateManagedPrivateEndpoint }).(pulumi.BoolPtrOutput)
}

func (o DataLakeStorageAccountDetailsOutput) Filesystem() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataLakeStorageAccountDetails) *string { return v.Filesystem }).(pulumi.StringPtrOutput)
}

func (o DataLakeStorageAccountDetailsOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataLakeStorageAccountDetails) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type DataLakeStorageAccountDetailsPtrOutput struct{ *pulumi.OutputState }

func (DataLakeStorageAccountDetailsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataLakeStorageAccountDetails)(nil)).Elem()
}

func (o DataLakeStorageAccountDetailsPtrOutput) ToDataLakeStorageAccountDetailsPtrOutput() DataLakeStorageAccountDetailsPtrOutput {
	return o
}

func (o DataLakeStorageAccountDetailsPtrOutput) ToDataLakeStorageAccountDetailsPtrOutputWithContext(ctx context.Context) DataLakeStorageAccountDetailsPtrOutput {
	return o
}

func (o DataLakeStorageAccountDetailsPtrOutput) Elem() DataLakeStorageAccountDetailsOutput {
	return o.ApplyT(func(v *DataLakeStorageAccountDetails) DataLakeStorageAccountDetails {
		if v != nil {
			return *v
		}
		var ret DataLakeStorageAccountDetails
		return ret
	}).(DataLakeStorageAccountDetailsOutput)
}

func (o DataLakeStorageAccountDetailsPtrOutput) AccountUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataLakeStorageAccountDetails) *string {
		if v == nil {
			return nil
		}
		return v.AccountUrl
	}).(pulumi.StringPtrOutput)
}

func (o DataLakeStorageAccountDetailsPtrOutput) CreateManagedPrivateEndpoint() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DataLakeStorageAccountDetails) *bool {
		if v == nil {
			return nil
		}
		return v.CreateManagedPrivateEndpoint
	}).(pulumi.BoolPtrOutput)
}

func (o DataLakeStorageAccountDetailsPtrOutput) Filesystem() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataLakeStorageAccountDetails) *string {
		if v == nil {
			return nil
		}
		return v.Filesystem
	}).(pulumi.StringPtrOutput)
}

func (o DataLakeStorageAccountDetailsPtrOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataLakeStorageAccountDetails) *string {
		if v == nil {
			return nil
		}
		return v.ResourceId
	}).(pulumi.StringPtrOutput)
}

type DataLakeStorageAccountDetailsResponse struct {
	AccountUrl                   *string `pulumi:"accountUrl"`
	CreateManagedPrivateEndpoint *bool   `pulumi:"createManagedPrivateEndpoint"`
	Filesystem                   *string `pulumi:"filesystem"`
	ResourceId                   *string `pulumi:"resourceId"`
}

type DataLakeStorageAccountDetailsResponseOutput struct{ *pulumi.OutputState }

func (DataLakeStorageAccountDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataLakeStorageAccountDetailsResponse)(nil)).Elem()
}

func (o DataLakeStorageAccountDetailsResponseOutput) ToDataLakeStorageAccountDetailsResponseOutput() DataLakeStorageAccountDetailsResponseOutput {
	return o
}

func (o DataLakeStorageAccountDetailsResponseOutput) ToDataLakeStorageAccountDetailsResponseOutputWithContext(ctx context.Context) DataLakeStorageAccountDetailsResponseOutput {
	return o
}

func (o DataLakeStorageAccountDetailsResponseOutput) AccountUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataLakeStorageAccountDetailsResponse) *string { return v.AccountUrl }).(pulumi.StringPtrOutput)
}

func (o DataLakeStorageAccountDetailsResponseOutput) CreateManagedPrivateEndpoint() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DataLakeStorageAccountDetailsResponse) *bool { return v.CreateManagedPrivateEndpoint }).(pulumi.BoolPtrOutput)
}

func (o DataLakeStorageAccountDetailsResponseOutput) Filesystem() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataLakeStorageAccountDetailsResponse) *string { return v.Filesystem }).(pulumi.StringPtrOutput)
}

func (o DataLakeStorageAccountDetailsResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataLakeStorageAccountDetailsResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type DataLakeStorageAccountDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (DataLakeStorageAccountDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataLakeStorageAccountDetailsResponse)(nil)).Elem()
}

func (o DataLakeStorageAccountDetailsResponsePtrOutput) ToDataLakeStorageAccountDetailsResponsePtrOutput() DataLakeStorageAccountDetailsResponsePtrOutput {
	return o
}

func (o DataLakeStorageAccountDetailsResponsePtrOutput) ToDataLakeStorageAccountDetailsResponsePtrOutputWithContext(ctx context.Context) DataLakeStorageAccountDetailsResponsePtrOutput {
	return o
}

func (o DataLakeStorageAccountDetailsResponsePtrOutput) Elem() DataLakeStorageAccountDetailsResponseOutput {
	return o.ApplyT(func(v *DataLakeStorageAccountDetailsResponse) DataLakeStorageAccountDetailsResponse {
		if v != nil {
			return *v
		}
		var ret DataLakeStorageAccountDetailsResponse
		return ret
	}).(DataLakeStorageAccountDetailsResponseOutput)
}

func (o DataLakeStorageAccountDetailsResponsePtrOutput) AccountUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataLakeStorageAccountDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.AccountUrl
	}).(pulumi.StringPtrOutput)
}

func (o DataLakeStorageAccountDetailsResponsePtrOutput) CreateManagedPrivateEndpoint() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DataLakeStorageAccountDetailsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.CreateManagedPrivateEndpoint
	}).(pulumi.BoolPtrOutput)
}

func (o DataLakeStorageAccountDetailsResponsePtrOutput) Filesystem() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataLakeStorageAccountDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Filesystem
	}).(pulumi.StringPtrOutput)
}

func (o DataLakeStorageAccountDetailsResponsePtrOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataLakeStorageAccountDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceId
	}).(pulumi.StringPtrOutput)
}

type DatabaseStatisticsResponse struct {
	Size *float64 `pulumi:"size"`
}

type DatabaseStatisticsResponseOutput struct{ *pulumi.OutputState }

func (DatabaseStatisticsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseStatisticsResponse)(nil)).Elem()
}

func (o DatabaseStatisticsResponseOutput) ToDatabaseStatisticsResponseOutput() DatabaseStatisticsResponseOutput {
	return o
}

func (o DatabaseStatisticsResponseOutput) ToDatabaseStatisticsResponseOutputWithContext(ctx context.Context) DatabaseStatisticsResponseOutput {
	return o
}

func (o DatabaseStatisticsResponseOutput) Size() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v DatabaseStatisticsResponse) *float64 { return v.Size }).(pulumi.Float64PtrOutput)
}

type DynamicExecutorAllocation struct {
	Enabled *bool `pulumi:"enabled"`
}





type DynamicExecutorAllocationInput interface {
	pulumi.Input

	ToDynamicExecutorAllocationOutput() DynamicExecutorAllocationOutput
	ToDynamicExecutorAllocationOutputWithContext(context.Context) DynamicExecutorAllocationOutput
}

type DynamicExecutorAllocationArgs struct {
	Enabled pulumi.BoolPtrInput `pulumi:"enabled"`
}

func (DynamicExecutorAllocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DynamicExecutorAllocation)(nil)).Elem()
}

func (i DynamicExecutorAllocationArgs) ToDynamicExecutorAllocationOutput() DynamicExecutorAllocationOutput {
	return i.ToDynamicExecutorAllocationOutputWithContext(context.Background())
}

func (i DynamicExecutorAllocationArgs) ToDynamicExecutorAllocationOutputWithContext(ctx context.Context) DynamicExecutorAllocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DynamicExecutorAllocationOutput)
}

func (i DynamicExecutorAllocationArgs) ToDynamicExecutorAllocationPtrOutput() DynamicExecutorAllocationPtrOutput {
	return i.ToDynamicExecutorAllocationPtrOutputWithContext(context.Background())
}

func (i DynamicExecutorAllocationArgs) ToDynamicExecutorAllocationPtrOutputWithContext(ctx context.Context) DynamicExecutorAllocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DynamicExecutorAllocationOutput).ToDynamicExecutorAllocationPtrOutputWithContext(ctx)
}









type DynamicExecutorAllocationPtrInput interface {
	pulumi.Input

	ToDynamicExecutorAllocationPtrOutput() DynamicExecutorAllocationPtrOutput
	ToDynamicExecutorAllocationPtrOutputWithContext(context.Context) DynamicExecutorAllocationPtrOutput
}

type dynamicExecutorAllocationPtrType DynamicExecutorAllocationArgs

func DynamicExecutorAllocationPtr(v *DynamicExecutorAllocationArgs) DynamicExecutorAllocationPtrInput {
	return (*dynamicExecutorAllocationPtrType)(v)
}

func (*dynamicExecutorAllocationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DynamicExecutorAllocation)(nil)).Elem()
}

func (i *dynamicExecutorAllocationPtrType) ToDynamicExecutorAllocationPtrOutput() DynamicExecutorAllocationPtrOutput {
	return i.ToDynamicExecutorAllocationPtrOutputWithContext(context.Background())
}

func (i *dynamicExecutorAllocationPtrType) ToDynamicExecutorAllocationPtrOutputWithContext(ctx context.Context) DynamicExecutorAllocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DynamicExecutorAllocationPtrOutput)
}

type DynamicExecutorAllocationOutput struct{ *pulumi.OutputState }

func (DynamicExecutorAllocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DynamicExecutorAllocation)(nil)).Elem()
}

func (o DynamicExecutorAllocationOutput) ToDynamicExecutorAllocationOutput() DynamicExecutorAllocationOutput {
	return o
}

func (o DynamicExecutorAllocationOutput) ToDynamicExecutorAllocationOutputWithContext(ctx context.Context) DynamicExecutorAllocationOutput {
	return o
}

func (o DynamicExecutorAllocationOutput) ToDynamicExecutorAllocationPtrOutput() DynamicExecutorAllocationPtrOutput {
	return o.ToDynamicExecutorAllocationPtrOutputWithContext(context.Background())
}

func (o DynamicExecutorAllocationOutput) ToDynamicExecutorAllocationPtrOutputWithContext(ctx context.Context) DynamicExecutorAllocationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DynamicExecutorAllocation) *DynamicExecutorAllocation {
		return &v
	}).(DynamicExecutorAllocationPtrOutput)
}

func (o DynamicExecutorAllocationOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DynamicExecutorAllocation) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

type DynamicExecutorAllocationPtrOutput struct{ *pulumi.OutputState }

func (DynamicExecutorAllocationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DynamicExecutorAllocation)(nil)).Elem()
}

func (o DynamicExecutorAllocationPtrOutput) ToDynamicExecutorAllocationPtrOutput() DynamicExecutorAllocationPtrOutput {
	return o
}

func (o DynamicExecutorAllocationPtrOutput) ToDynamicExecutorAllocationPtrOutputWithContext(ctx context.Context) DynamicExecutorAllocationPtrOutput {
	return o
}

func (o DynamicExecutorAllocationPtrOutput) Elem() DynamicExecutorAllocationOutput {
	return o.ApplyT(func(v *DynamicExecutorAllocation) DynamicExecutorAllocation {
		if v != nil {
			return *v
		}
		var ret DynamicExecutorAllocation
		return ret
	}).(DynamicExecutorAllocationOutput)
}

func (o DynamicExecutorAllocationPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DynamicExecutorAllocation) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

type DynamicExecutorAllocationResponse struct {
	Enabled *bool `pulumi:"enabled"`
}

type DynamicExecutorAllocationResponseOutput struct{ *pulumi.OutputState }

func (DynamicExecutorAllocationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DynamicExecutorAllocationResponse)(nil)).Elem()
}

func (o DynamicExecutorAllocationResponseOutput) ToDynamicExecutorAllocationResponseOutput() DynamicExecutorAllocationResponseOutput {
	return o
}

func (o DynamicExecutorAllocationResponseOutput) ToDynamicExecutorAllocationResponseOutputWithContext(ctx context.Context) DynamicExecutorAllocationResponseOutput {
	return o
}

func (o DynamicExecutorAllocationResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DynamicExecutorAllocationResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

type DynamicExecutorAllocationResponsePtrOutput struct{ *pulumi.OutputState }

func (DynamicExecutorAllocationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DynamicExecutorAllocationResponse)(nil)).Elem()
}

func (o DynamicExecutorAllocationResponsePtrOutput) ToDynamicExecutorAllocationResponsePtrOutput() DynamicExecutorAllocationResponsePtrOutput {
	return o
}

func (o DynamicExecutorAllocationResponsePtrOutput) ToDynamicExecutorAllocationResponsePtrOutputWithContext(ctx context.Context) DynamicExecutorAllocationResponsePtrOutput {
	return o
}

func (o DynamicExecutorAllocationResponsePtrOutput) Elem() DynamicExecutorAllocationResponseOutput {
	return o.ApplyT(func(v *DynamicExecutorAllocationResponse) DynamicExecutorAllocationResponse {
		if v != nil {
			return *v
		}
		var ret DynamicExecutorAllocationResponse
		return ret
	}).(DynamicExecutorAllocationResponseOutput)
}

func (o DynamicExecutorAllocationResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DynamicExecutorAllocationResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

type EncryptionDetails struct {
	Cmk *CustomerManagedKeyDetails `pulumi:"cmk"`
}





type EncryptionDetailsInput interface {
	pulumi.Input

	ToEncryptionDetailsOutput() EncryptionDetailsOutput
	ToEncryptionDetailsOutputWithContext(context.Context) EncryptionDetailsOutput
}

type EncryptionDetailsArgs struct {
	Cmk CustomerManagedKeyDetailsPtrInput `pulumi:"cmk"`
}

func (EncryptionDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionDetails)(nil)).Elem()
}

func (i EncryptionDetailsArgs) ToEncryptionDetailsOutput() EncryptionDetailsOutput {
	return i.ToEncryptionDetailsOutputWithContext(context.Background())
}

func (i EncryptionDetailsArgs) ToEncryptionDetailsOutputWithContext(ctx context.Context) EncryptionDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionDetailsOutput)
}

func (i EncryptionDetailsArgs) ToEncryptionDetailsPtrOutput() EncryptionDetailsPtrOutput {
	return i.ToEncryptionDetailsPtrOutputWithContext(context.Background())
}

func (i EncryptionDetailsArgs) ToEncryptionDetailsPtrOutputWithContext(ctx context.Context) EncryptionDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionDetailsOutput).ToEncryptionDetailsPtrOutputWithContext(ctx)
}









type EncryptionDetailsPtrInput interface {
	pulumi.Input

	ToEncryptionDetailsPtrOutput() EncryptionDetailsPtrOutput
	ToEncryptionDetailsPtrOutputWithContext(context.Context) EncryptionDetailsPtrOutput
}

type encryptionDetailsPtrType EncryptionDetailsArgs

func EncryptionDetailsPtr(v *EncryptionDetailsArgs) EncryptionDetailsPtrInput {
	return (*encryptionDetailsPtrType)(v)
}

func (*encryptionDetailsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionDetails)(nil)).Elem()
}

func (i *encryptionDetailsPtrType) ToEncryptionDetailsPtrOutput() EncryptionDetailsPtrOutput {
	return i.ToEncryptionDetailsPtrOutputWithContext(context.Background())
}

func (i *encryptionDetailsPtrType) ToEncryptionDetailsPtrOutputWithContext(ctx context.Context) EncryptionDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionDetailsPtrOutput)
}

type EncryptionDetailsOutput struct{ *pulumi.OutputState }

func (EncryptionDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionDetails)(nil)).Elem()
}

func (o EncryptionDetailsOutput) ToEncryptionDetailsOutput() EncryptionDetailsOutput {
	return o
}

func (o EncryptionDetailsOutput) ToEncryptionDetailsOutputWithContext(ctx context.Context) EncryptionDetailsOutput {
	return o
}

func (o EncryptionDetailsOutput) ToEncryptionDetailsPtrOutput() EncryptionDetailsPtrOutput {
	return o.ToEncryptionDetailsPtrOutputWithContext(context.Background())
}

func (o EncryptionDetailsOutput) ToEncryptionDetailsPtrOutputWithContext(ctx context.Context) EncryptionDetailsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionDetails) *EncryptionDetails {
		return &v
	}).(EncryptionDetailsPtrOutput)
}

func (o EncryptionDetailsOutput) Cmk() CustomerManagedKeyDetailsPtrOutput {
	return o.ApplyT(func(v EncryptionDetails) *CustomerManagedKeyDetails { return v.Cmk }).(CustomerManagedKeyDetailsPtrOutput)
}

type EncryptionDetailsPtrOutput struct{ *pulumi.OutputState }

func (EncryptionDetailsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionDetails)(nil)).Elem()
}

func (o EncryptionDetailsPtrOutput) ToEncryptionDetailsPtrOutput() EncryptionDetailsPtrOutput {
	return o
}

func (o EncryptionDetailsPtrOutput) ToEncryptionDetailsPtrOutputWithContext(ctx context.Context) EncryptionDetailsPtrOutput {
	return o
}

func (o EncryptionDetailsPtrOutput) Elem() EncryptionDetailsOutput {
	return o.ApplyT(func(v *EncryptionDetails) EncryptionDetails {
		if v != nil {
			return *v
		}
		var ret EncryptionDetails
		return ret
	}).(EncryptionDetailsOutput)
}

func (o EncryptionDetailsPtrOutput) Cmk() CustomerManagedKeyDetailsPtrOutput {
	return o.ApplyT(func(v *EncryptionDetails) *CustomerManagedKeyDetails {
		if v == nil {
			return nil
		}
		return v.Cmk
	}).(CustomerManagedKeyDetailsPtrOutput)
}

type EncryptionDetailsResponse struct {
	Cmk                     *CustomerManagedKeyDetailsResponse `pulumi:"cmk"`
	DoubleEncryptionEnabled bool                               `pulumi:"doubleEncryptionEnabled"`
}

type EncryptionDetailsResponseOutput struct{ *pulumi.OutputState }

func (EncryptionDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionDetailsResponse)(nil)).Elem()
}

func (o EncryptionDetailsResponseOutput) ToEncryptionDetailsResponseOutput() EncryptionDetailsResponseOutput {
	return o
}

func (o EncryptionDetailsResponseOutput) ToEncryptionDetailsResponseOutputWithContext(ctx context.Context) EncryptionDetailsResponseOutput {
	return o
}

func (o EncryptionDetailsResponseOutput) Cmk() CustomerManagedKeyDetailsResponsePtrOutput {
	return o.ApplyT(func(v EncryptionDetailsResponse) *CustomerManagedKeyDetailsResponse { return v.Cmk }).(CustomerManagedKeyDetailsResponsePtrOutput)
}

func (o EncryptionDetailsResponseOutput) DoubleEncryptionEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v EncryptionDetailsResponse) bool { return v.DoubleEncryptionEnabled }).(pulumi.BoolOutput)
}

type EncryptionDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionDetailsResponse)(nil)).Elem()
}

func (o EncryptionDetailsResponsePtrOutput) ToEncryptionDetailsResponsePtrOutput() EncryptionDetailsResponsePtrOutput {
	return o
}

func (o EncryptionDetailsResponsePtrOutput) ToEncryptionDetailsResponsePtrOutputWithContext(ctx context.Context) EncryptionDetailsResponsePtrOutput {
	return o
}

func (o EncryptionDetailsResponsePtrOutput) Elem() EncryptionDetailsResponseOutput {
	return o.ApplyT(func(v *EncryptionDetailsResponse) EncryptionDetailsResponse {
		if v != nil {
			return *v
		}
		var ret EncryptionDetailsResponse
		return ret
	}).(EncryptionDetailsResponseOutput)
}

func (o EncryptionDetailsResponsePtrOutput) Cmk() CustomerManagedKeyDetailsResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionDetailsResponse) *CustomerManagedKeyDetailsResponse {
		if v == nil {
			return nil
		}
		return v.Cmk
	}).(CustomerManagedKeyDetailsResponsePtrOutput)
}

func (o EncryptionDetailsResponsePtrOutput) DoubleEncryptionEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EncryptionDetailsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.DoubleEncryptionEnabled
	}).(pulumi.BoolPtrOutput)
}

type EntityReference struct {
	ReferenceName *string `pulumi:"referenceName"`
	Type          *string `pulumi:"type"`
}

type EntityReferenceResponse struct {
	ReferenceName *string `pulumi:"referenceName"`
	Type          *string `pulumi:"type"`
}

type EnvironmentVariableSetup struct {
	Type          string `pulumi:"type"`
	VariableName  string `pulumi:"variableName"`
	VariableValue string `pulumi:"variableValue"`
}

type EnvironmentVariableSetupResponse struct {
	Type          string `pulumi:"type"`
	VariableName  string `pulumi:"variableName"`
	VariableValue string `pulumi:"variableValue"`
}

type FollowerDatabaseDefinitionResponse struct {
	AttachedDatabaseConfigurationName string `pulumi:"attachedDatabaseConfigurationName"`
	DatabaseName                      string `pulumi:"databaseName"`
	KustoPoolResourceId               string `pulumi:"kustoPoolResourceId"`
}

type IntegrationRuntimeComputeProperties struct {
	DataFlowProperties           *IntegrationRuntimeDataFlowProperties `pulumi:"dataFlowProperties"`
	Location                     *string                               `pulumi:"location"`
	MaxParallelExecutionsPerNode *int                                  `pulumi:"maxParallelExecutionsPerNode"`
	NodeSize                     *string                               `pulumi:"nodeSize"`
	NumberOfNodes                *int                                  `pulumi:"numberOfNodes"`
	VNetProperties               *IntegrationRuntimeVNetProperties     `pulumi:"vNetProperties"`
}

type IntegrationRuntimeComputePropertiesResponse struct {
	DataFlowProperties           *IntegrationRuntimeDataFlowPropertiesResponse `pulumi:"dataFlowProperties"`
	Location                     *string                                       `pulumi:"location"`
	MaxParallelExecutionsPerNode *int                                          `pulumi:"maxParallelExecutionsPerNode"`
	NodeSize                     *string                                       `pulumi:"nodeSize"`
	NumberOfNodes                *int                                          `pulumi:"numberOfNodes"`
	VNetProperties               *IntegrationRuntimeVNetPropertiesResponse     `pulumi:"vNetProperties"`
}

type IntegrationRuntimeCustomSetupScriptProperties struct {
	BlobContainerUri *string       `pulumi:"blobContainerUri"`
	SasToken         *SecureString `pulumi:"sasToken"`
}

type IntegrationRuntimeCustomSetupScriptPropertiesResponse struct {
	BlobContainerUri *string               `pulumi:"blobContainerUri"`
	SasToken         *SecureStringResponse `pulumi:"sasToken"`
}

type IntegrationRuntimeDataFlowProperties struct {
	Cleanup     *bool   `pulumi:"cleanup"`
	ComputeType *string `pulumi:"computeType"`
	CoreCount   *int    `pulumi:"coreCount"`
	TimeToLive  *int    `pulumi:"timeToLive"`
}

type IntegrationRuntimeDataFlowPropertiesResponse struct {
	Cleanup     *bool   `pulumi:"cleanup"`
	ComputeType *string `pulumi:"computeType"`
	CoreCount   *int    `pulumi:"coreCount"`
	TimeToLive  *int    `pulumi:"timeToLive"`
}

type IntegrationRuntimeDataProxyProperties struct {
	ConnectVia           *EntityReference `pulumi:"connectVia"`
	Path                 *string          `pulumi:"path"`
	StagingLinkedService *EntityReference `pulumi:"stagingLinkedService"`
}

type IntegrationRuntimeDataProxyPropertiesResponse struct {
	ConnectVia           *EntityReferenceResponse `pulumi:"connectVia"`
	Path                 *string                  `pulumi:"path"`
	StagingLinkedService *EntityReferenceResponse `pulumi:"stagingLinkedService"`
}

type IntegrationRuntimeSsisCatalogInfo struct {
	CatalogAdminPassword  *SecureString `pulumi:"catalogAdminPassword"`
	CatalogAdminUserName  *string       `pulumi:"catalogAdminUserName"`
	CatalogPricingTier    *string       `pulumi:"catalogPricingTier"`
	CatalogServerEndpoint *string       `pulumi:"catalogServerEndpoint"`
}

type IntegrationRuntimeSsisCatalogInfoResponse struct {
	CatalogAdminPassword  *SecureStringResponse `pulumi:"catalogAdminPassword"`
	CatalogAdminUserName  *string               `pulumi:"catalogAdminUserName"`
	CatalogPricingTier    *string               `pulumi:"catalogPricingTier"`
	CatalogServerEndpoint *string               `pulumi:"catalogServerEndpoint"`
}

type IntegrationRuntimeSsisProperties struct {
	CatalogInfo                  *IntegrationRuntimeSsisCatalogInfo             `pulumi:"catalogInfo"`
	CustomSetupScriptProperties  *IntegrationRuntimeCustomSetupScriptProperties `pulumi:"customSetupScriptProperties"`
	DataProxyProperties          *IntegrationRuntimeDataProxyProperties         `pulumi:"dataProxyProperties"`
	Edition                      *string                                        `pulumi:"edition"`
	ExpressCustomSetupProperties []interface{}                                  `pulumi:"expressCustomSetupProperties"`
	LicenseType                  *string                                        `pulumi:"licenseType"`
}

type IntegrationRuntimeSsisPropertiesResponse struct {
	CatalogInfo                  *IntegrationRuntimeSsisCatalogInfoResponse             `pulumi:"catalogInfo"`
	CustomSetupScriptProperties  *IntegrationRuntimeCustomSetupScriptPropertiesResponse `pulumi:"customSetupScriptProperties"`
	DataProxyProperties          *IntegrationRuntimeDataProxyPropertiesResponse         `pulumi:"dataProxyProperties"`
	Edition                      *string                                                `pulumi:"edition"`
	ExpressCustomSetupProperties []interface{}                                          `pulumi:"expressCustomSetupProperties"`
	LicenseType                  *string                                                `pulumi:"licenseType"`
}

type IntegrationRuntimeVNetProperties struct {
	PublicIPs []string `pulumi:"publicIPs"`
	Subnet    *string  `pulumi:"subnet"`
	SubnetId  *string  `pulumi:"subnetId"`
	VNetId    *string  `pulumi:"vNetId"`
}

type IntegrationRuntimeVNetPropertiesResponse struct {
	PublicIPs []string `pulumi:"publicIPs"`
	Subnet    *string  `pulumi:"subnet"`
	SubnetId  *string  `pulumi:"subnetId"`
	VNetId    *string  `pulumi:"vNetId"`
}

type KekIdentityProperties struct {
	UseSystemAssignedIdentity interface{} `pulumi:"useSystemAssignedIdentity"`
	UserAssignedIdentity      *string     `pulumi:"userAssignedIdentity"`
}





type KekIdentityPropertiesInput interface {
	pulumi.Input

	ToKekIdentityPropertiesOutput() KekIdentityPropertiesOutput
	ToKekIdentityPropertiesOutputWithContext(context.Context) KekIdentityPropertiesOutput
}

type KekIdentityPropertiesArgs struct {
	UseSystemAssignedIdentity pulumi.Input          `pulumi:"useSystemAssignedIdentity"`
	UserAssignedIdentity      pulumi.StringPtrInput `pulumi:"userAssignedIdentity"`
}

func (KekIdentityPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KekIdentityProperties)(nil)).Elem()
}

func (i KekIdentityPropertiesArgs) ToKekIdentityPropertiesOutput() KekIdentityPropertiesOutput {
	return i.ToKekIdentityPropertiesOutputWithContext(context.Background())
}

func (i KekIdentityPropertiesArgs) ToKekIdentityPropertiesOutputWithContext(ctx context.Context) KekIdentityPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KekIdentityPropertiesOutput)
}

func (i KekIdentityPropertiesArgs) ToKekIdentityPropertiesPtrOutput() KekIdentityPropertiesPtrOutput {
	return i.ToKekIdentityPropertiesPtrOutputWithContext(context.Background())
}

func (i KekIdentityPropertiesArgs) ToKekIdentityPropertiesPtrOutputWithContext(ctx context.Context) KekIdentityPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KekIdentityPropertiesOutput).ToKekIdentityPropertiesPtrOutputWithContext(ctx)
}









type KekIdentityPropertiesPtrInput interface {
	pulumi.Input

	ToKekIdentityPropertiesPtrOutput() KekIdentityPropertiesPtrOutput
	ToKekIdentityPropertiesPtrOutputWithContext(context.Context) KekIdentityPropertiesPtrOutput
}

type kekIdentityPropertiesPtrType KekIdentityPropertiesArgs

func KekIdentityPropertiesPtr(v *KekIdentityPropertiesArgs) KekIdentityPropertiesPtrInput {
	return (*kekIdentityPropertiesPtrType)(v)
}

func (*kekIdentityPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KekIdentityProperties)(nil)).Elem()
}

func (i *kekIdentityPropertiesPtrType) ToKekIdentityPropertiesPtrOutput() KekIdentityPropertiesPtrOutput {
	return i.ToKekIdentityPropertiesPtrOutputWithContext(context.Background())
}

func (i *kekIdentityPropertiesPtrType) ToKekIdentityPropertiesPtrOutputWithContext(ctx context.Context) KekIdentityPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KekIdentityPropertiesPtrOutput)
}

type KekIdentityPropertiesOutput struct{ *pulumi.OutputState }

func (KekIdentityPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KekIdentityProperties)(nil)).Elem()
}

func (o KekIdentityPropertiesOutput) ToKekIdentityPropertiesOutput() KekIdentityPropertiesOutput {
	return o
}

func (o KekIdentityPropertiesOutput) ToKekIdentityPropertiesOutputWithContext(ctx context.Context) KekIdentityPropertiesOutput {
	return o
}

func (o KekIdentityPropertiesOutput) ToKekIdentityPropertiesPtrOutput() KekIdentityPropertiesPtrOutput {
	return o.ToKekIdentityPropertiesPtrOutputWithContext(context.Background())
}

func (o KekIdentityPropertiesOutput) ToKekIdentityPropertiesPtrOutputWithContext(ctx context.Context) KekIdentityPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KekIdentityProperties) *KekIdentityProperties {
		return &v
	}).(KekIdentityPropertiesPtrOutput)
}

func (o KekIdentityPropertiesOutput) UseSystemAssignedIdentity() pulumi.AnyOutput {
	return o.ApplyT(func(v KekIdentityProperties) interface{} { return v.UseSystemAssignedIdentity }).(pulumi.AnyOutput)
}

func (o KekIdentityPropertiesOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KekIdentityProperties) *string { return v.UserAssignedIdentity }).(pulumi.StringPtrOutput)
}

type KekIdentityPropertiesPtrOutput struct{ *pulumi.OutputState }

func (KekIdentityPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KekIdentityProperties)(nil)).Elem()
}

func (o KekIdentityPropertiesPtrOutput) ToKekIdentityPropertiesPtrOutput() KekIdentityPropertiesPtrOutput {
	return o
}

func (o KekIdentityPropertiesPtrOutput) ToKekIdentityPropertiesPtrOutputWithContext(ctx context.Context) KekIdentityPropertiesPtrOutput {
	return o
}

func (o KekIdentityPropertiesPtrOutput) Elem() KekIdentityPropertiesOutput {
	return o.ApplyT(func(v *KekIdentityProperties) KekIdentityProperties {
		if v != nil {
			return *v
		}
		var ret KekIdentityProperties
		return ret
	}).(KekIdentityPropertiesOutput)
}

func (o KekIdentityPropertiesPtrOutput) UseSystemAssignedIdentity() pulumi.AnyOutput {
	return o.ApplyT(func(v *KekIdentityProperties) interface{} {
		if v == nil {
			return nil
		}
		return v.UseSystemAssignedIdentity
	}).(pulumi.AnyOutput)
}

func (o KekIdentityPropertiesPtrOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KekIdentityProperties) *string {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentity
	}).(pulumi.StringPtrOutput)
}

type KekIdentityPropertiesResponse struct {
	UseSystemAssignedIdentity interface{} `pulumi:"useSystemAssignedIdentity"`
	UserAssignedIdentity      *string     `pulumi:"userAssignedIdentity"`
}

type KekIdentityPropertiesResponseOutput struct{ *pulumi.OutputState }

func (KekIdentityPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KekIdentityPropertiesResponse)(nil)).Elem()
}

func (o KekIdentityPropertiesResponseOutput) ToKekIdentityPropertiesResponseOutput() KekIdentityPropertiesResponseOutput {
	return o
}

func (o KekIdentityPropertiesResponseOutput) ToKekIdentityPropertiesResponseOutputWithContext(ctx context.Context) KekIdentityPropertiesResponseOutput {
	return o
}

func (o KekIdentityPropertiesResponseOutput) UseSystemAssignedIdentity() pulumi.AnyOutput {
	return o.ApplyT(func(v KekIdentityPropertiesResponse) interface{} { return v.UseSystemAssignedIdentity }).(pulumi.AnyOutput)
}

func (o KekIdentityPropertiesResponseOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KekIdentityPropertiesResponse) *string { return v.UserAssignedIdentity }).(pulumi.StringPtrOutput)
}

type KekIdentityPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (KekIdentityPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KekIdentityPropertiesResponse)(nil)).Elem()
}

func (o KekIdentityPropertiesResponsePtrOutput) ToKekIdentityPropertiesResponsePtrOutput() KekIdentityPropertiesResponsePtrOutput {
	return o
}

func (o KekIdentityPropertiesResponsePtrOutput) ToKekIdentityPropertiesResponsePtrOutputWithContext(ctx context.Context) KekIdentityPropertiesResponsePtrOutput {
	return o
}

func (o KekIdentityPropertiesResponsePtrOutput) Elem() KekIdentityPropertiesResponseOutput {
	return o.ApplyT(func(v *KekIdentityPropertiesResponse) KekIdentityPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret KekIdentityPropertiesResponse
		return ret
	}).(KekIdentityPropertiesResponseOutput)
}

func (o KekIdentityPropertiesResponsePtrOutput) UseSystemAssignedIdentity() pulumi.AnyOutput {
	return o.ApplyT(func(v *KekIdentityPropertiesResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.UseSystemAssignedIdentity
	}).(pulumi.AnyOutput)
}

func (o KekIdentityPropertiesResponsePtrOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KekIdentityPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentity
	}).(pulumi.StringPtrOutput)
}

type LanguageExtensionResponse struct {
	LanguageExtensionName *string `pulumi:"languageExtensionName"`
}

type LanguageExtensionResponseOutput struct{ *pulumi.OutputState }

func (LanguageExtensionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LanguageExtensionResponse)(nil)).Elem()
}

func (o LanguageExtensionResponseOutput) ToLanguageExtensionResponseOutput() LanguageExtensionResponseOutput {
	return o
}

func (o LanguageExtensionResponseOutput) ToLanguageExtensionResponseOutputWithContext(ctx context.Context) LanguageExtensionResponseOutput {
	return o
}

func (o LanguageExtensionResponseOutput) LanguageExtensionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LanguageExtensionResponse) *string { return v.LanguageExtensionName }).(pulumi.StringPtrOutput)
}

type LanguageExtensionResponseArrayOutput struct{ *pulumi.OutputState }

func (LanguageExtensionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LanguageExtensionResponse)(nil)).Elem()
}

func (o LanguageExtensionResponseArrayOutput) ToLanguageExtensionResponseArrayOutput() LanguageExtensionResponseArrayOutput {
	return o
}

func (o LanguageExtensionResponseArrayOutput) ToLanguageExtensionResponseArrayOutputWithContext(ctx context.Context) LanguageExtensionResponseArrayOutput {
	return o
}

func (o LanguageExtensionResponseArrayOutput) Index(i pulumi.IntInput) LanguageExtensionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LanguageExtensionResponse {
		return vs[0].([]LanguageExtensionResponse)[vs[1].(int)]
	}).(LanguageExtensionResponseOutput)
}

type LanguageExtensionsListResponse struct {
	Value []LanguageExtensionResponse `pulumi:"value"`
}

type LanguageExtensionsListResponseOutput struct{ *pulumi.OutputState }

func (LanguageExtensionsListResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LanguageExtensionsListResponse)(nil)).Elem()
}

func (o LanguageExtensionsListResponseOutput) ToLanguageExtensionsListResponseOutput() LanguageExtensionsListResponseOutput {
	return o
}

func (o LanguageExtensionsListResponseOutput) ToLanguageExtensionsListResponseOutputWithContext(ctx context.Context) LanguageExtensionsListResponseOutput {
	return o
}

func (o LanguageExtensionsListResponseOutput) Value() LanguageExtensionResponseArrayOutput {
	return o.ApplyT(func(v LanguageExtensionsListResponse) []LanguageExtensionResponse { return v.Value }).(LanguageExtensionResponseArrayOutput)
}

type LibraryInfo struct {
	ContainerName     *string `pulumi:"containerName"`
	Name              *string `pulumi:"name"`
	Path              *string `pulumi:"path"`
	Type              *string `pulumi:"type"`
	UploadedTimestamp *string `pulumi:"uploadedTimestamp"`
}





type LibraryInfoInput interface {
	pulumi.Input

	ToLibraryInfoOutput() LibraryInfoOutput
	ToLibraryInfoOutputWithContext(context.Context) LibraryInfoOutput
}

type LibraryInfoArgs struct {
	ContainerName     pulumi.StringPtrInput `pulumi:"containerName"`
	Name              pulumi.StringPtrInput `pulumi:"name"`
	Path              pulumi.StringPtrInput `pulumi:"path"`
	Type              pulumi.StringPtrInput `pulumi:"type"`
	UploadedTimestamp pulumi.StringPtrInput `pulumi:"uploadedTimestamp"`
}

func (LibraryInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LibraryInfo)(nil)).Elem()
}

func (i LibraryInfoArgs) ToLibraryInfoOutput() LibraryInfoOutput {
	return i.ToLibraryInfoOutputWithContext(context.Background())
}

func (i LibraryInfoArgs) ToLibraryInfoOutputWithContext(ctx context.Context) LibraryInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LibraryInfoOutput)
}





type LibraryInfoArrayInput interface {
	pulumi.Input

	ToLibraryInfoArrayOutput() LibraryInfoArrayOutput
	ToLibraryInfoArrayOutputWithContext(context.Context) LibraryInfoArrayOutput
}

type LibraryInfoArray []LibraryInfoInput

func (LibraryInfoArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LibraryInfo)(nil)).Elem()
}

func (i LibraryInfoArray) ToLibraryInfoArrayOutput() LibraryInfoArrayOutput {
	return i.ToLibraryInfoArrayOutputWithContext(context.Background())
}

func (i LibraryInfoArray) ToLibraryInfoArrayOutputWithContext(ctx context.Context) LibraryInfoArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LibraryInfoArrayOutput)
}

type LibraryInfoOutput struct{ *pulumi.OutputState }

func (LibraryInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LibraryInfo)(nil)).Elem()
}

func (o LibraryInfoOutput) ToLibraryInfoOutput() LibraryInfoOutput {
	return o
}

func (o LibraryInfoOutput) ToLibraryInfoOutputWithContext(ctx context.Context) LibraryInfoOutput {
	return o
}

func (o LibraryInfoOutput) ContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LibraryInfo) *string { return v.ContainerName }).(pulumi.StringPtrOutput)
}

func (o LibraryInfoOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LibraryInfo) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LibraryInfoOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LibraryInfo) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o LibraryInfoOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LibraryInfo) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o LibraryInfoOutput) UploadedTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LibraryInfo) *string { return v.UploadedTimestamp }).(pulumi.StringPtrOutput)
}

type LibraryInfoArrayOutput struct{ *pulumi.OutputState }

func (LibraryInfoArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LibraryInfo)(nil)).Elem()
}

func (o LibraryInfoArrayOutput) ToLibraryInfoArrayOutput() LibraryInfoArrayOutput {
	return o
}

func (o LibraryInfoArrayOutput) ToLibraryInfoArrayOutputWithContext(ctx context.Context) LibraryInfoArrayOutput {
	return o
}

func (o LibraryInfoArrayOutput) Index(i pulumi.IntInput) LibraryInfoOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LibraryInfo {
		return vs[0].([]LibraryInfo)[vs[1].(int)]
	}).(LibraryInfoOutput)
}

type LibraryInfoResponse struct {
	ContainerName      *string `pulumi:"containerName"`
	CreatorId          string  `pulumi:"creatorId"`
	Name               *string `pulumi:"name"`
	Path               *string `pulumi:"path"`
	ProvisioningStatus string  `pulumi:"provisioningStatus"`
	Type               *string `pulumi:"type"`
	UploadedTimestamp  *string `pulumi:"uploadedTimestamp"`
}

type LibraryInfoResponseOutput struct{ *pulumi.OutputState }

func (LibraryInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LibraryInfoResponse)(nil)).Elem()
}

func (o LibraryInfoResponseOutput) ToLibraryInfoResponseOutput() LibraryInfoResponseOutput {
	return o
}

func (o LibraryInfoResponseOutput) ToLibraryInfoResponseOutputWithContext(ctx context.Context) LibraryInfoResponseOutput {
	return o
}

func (o LibraryInfoResponseOutput) ContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LibraryInfoResponse) *string { return v.ContainerName }).(pulumi.StringPtrOutput)
}

func (o LibraryInfoResponseOutput) CreatorId() pulumi.StringOutput {
	return o.ApplyT(func(v LibraryInfoResponse) string { return v.CreatorId }).(pulumi.StringOutput)
}

func (o LibraryInfoResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LibraryInfoResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LibraryInfoResponseOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LibraryInfoResponse) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o LibraryInfoResponseOutput) ProvisioningStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LibraryInfoResponse) string { return v.ProvisioningStatus }).(pulumi.StringOutput)
}

func (o LibraryInfoResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LibraryInfoResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o LibraryInfoResponseOutput) UploadedTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LibraryInfoResponse) *string { return v.UploadedTimestamp }).(pulumi.StringPtrOutput)
}

type LibraryInfoResponseArrayOutput struct{ *pulumi.OutputState }

func (LibraryInfoResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LibraryInfoResponse)(nil)).Elem()
}

func (o LibraryInfoResponseArrayOutput) ToLibraryInfoResponseArrayOutput() LibraryInfoResponseArrayOutput {
	return o
}

func (o LibraryInfoResponseArrayOutput) ToLibraryInfoResponseArrayOutputWithContext(ctx context.Context) LibraryInfoResponseArrayOutput {
	return o
}

func (o LibraryInfoResponseArrayOutput) Index(i pulumi.IntInput) LibraryInfoResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LibraryInfoResponse {
		return vs[0].([]LibraryInfoResponse)[vs[1].(int)]
	}).(LibraryInfoResponseOutput)
}

type LibraryRequirements struct {
	Content  *string `pulumi:"content"`
	Filename *string `pulumi:"filename"`
}





type LibraryRequirementsInput interface {
	pulumi.Input

	ToLibraryRequirementsOutput() LibraryRequirementsOutput
	ToLibraryRequirementsOutputWithContext(context.Context) LibraryRequirementsOutput
}

type LibraryRequirementsArgs struct {
	Content  pulumi.StringPtrInput `pulumi:"content"`
	Filename pulumi.StringPtrInput `pulumi:"filename"`
}

func (LibraryRequirementsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LibraryRequirements)(nil)).Elem()
}

func (i LibraryRequirementsArgs) ToLibraryRequirementsOutput() LibraryRequirementsOutput {
	return i.ToLibraryRequirementsOutputWithContext(context.Background())
}

func (i LibraryRequirementsArgs) ToLibraryRequirementsOutputWithContext(ctx context.Context) LibraryRequirementsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LibraryRequirementsOutput)
}

func (i LibraryRequirementsArgs) ToLibraryRequirementsPtrOutput() LibraryRequirementsPtrOutput {
	return i.ToLibraryRequirementsPtrOutputWithContext(context.Background())
}

func (i LibraryRequirementsArgs) ToLibraryRequirementsPtrOutputWithContext(ctx context.Context) LibraryRequirementsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LibraryRequirementsOutput).ToLibraryRequirementsPtrOutputWithContext(ctx)
}









type LibraryRequirementsPtrInput interface {
	pulumi.Input

	ToLibraryRequirementsPtrOutput() LibraryRequirementsPtrOutput
	ToLibraryRequirementsPtrOutputWithContext(context.Context) LibraryRequirementsPtrOutput
}

type libraryRequirementsPtrType LibraryRequirementsArgs

func LibraryRequirementsPtr(v *LibraryRequirementsArgs) LibraryRequirementsPtrInput {
	return (*libraryRequirementsPtrType)(v)
}

func (*libraryRequirementsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LibraryRequirements)(nil)).Elem()
}

func (i *libraryRequirementsPtrType) ToLibraryRequirementsPtrOutput() LibraryRequirementsPtrOutput {
	return i.ToLibraryRequirementsPtrOutputWithContext(context.Background())
}

func (i *libraryRequirementsPtrType) ToLibraryRequirementsPtrOutputWithContext(ctx context.Context) LibraryRequirementsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LibraryRequirementsPtrOutput)
}

type LibraryRequirementsOutput struct{ *pulumi.OutputState }

func (LibraryRequirementsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LibraryRequirements)(nil)).Elem()
}

func (o LibraryRequirementsOutput) ToLibraryRequirementsOutput() LibraryRequirementsOutput {
	return o
}

func (o LibraryRequirementsOutput) ToLibraryRequirementsOutputWithContext(ctx context.Context) LibraryRequirementsOutput {
	return o
}

func (o LibraryRequirementsOutput) ToLibraryRequirementsPtrOutput() LibraryRequirementsPtrOutput {
	return o.ToLibraryRequirementsPtrOutputWithContext(context.Background())
}

func (o LibraryRequirementsOutput) ToLibraryRequirementsPtrOutputWithContext(ctx context.Context) LibraryRequirementsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LibraryRequirements) *LibraryRequirements {
		return &v
	}).(LibraryRequirementsPtrOutput)
}

func (o LibraryRequirementsOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LibraryRequirements) *string { return v.Content }).(pulumi.StringPtrOutput)
}

func (o LibraryRequirementsOutput) Filename() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LibraryRequirements) *string { return v.Filename }).(pulumi.StringPtrOutput)
}

type LibraryRequirementsPtrOutput struct{ *pulumi.OutputState }

func (LibraryRequirementsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LibraryRequirements)(nil)).Elem()
}

func (o LibraryRequirementsPtrOutput) ToLibraryRequirementsPtrOutput() LibraryRequirementsPtrOutput {
	return o
}

func (o LibraryRequirementsPtrOutput) ToLibraryRequirementsPtrOutputWithContext(ctx context.Context) LibraryRequirementsPtrOutput {
	return o
}

func (o LibraryRequirementsPtrOutput) Elem() LibraryRequirementsOutput {
	return o.ApplyT(func(v *LibraryRequirements) LibraryRequirements {
		if v != nil {
			return *v
		}
		var ret LibraryRequirements
		return ret
	}).(LibraryRequirementsOutput)
}

func (o LibraryRequirementsPtrOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LibraryRequirements) *string {
		if v == nil {
			return nil
		}
		return v.Content
	}).(pulumi.StringPtrOutput)
}

func (o LibraryRequirementsPtrOutput) Filename() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LibraryRequirements) *string {
		if v == nil {
			return nil
		}
		return v.Filename
	}).(pulumi.StringPtrOutput)
}

type LibraryRequirementsResponse struct {
	Content  *string `pulumi:"content"`
	Filename *string `pulumi:"filename"`
	Time     string  `pulumi:"time"`
}

type LibraryRequirementsResponseOutput struct{ *pulumi.OutputState }

func (LibraryRequirementsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LibraryRequirementsResponse)(nil)).Elem()
}

func (o LibraryRequirementsResponseOutput) ToLibraryRequirementsResponseOutput() LibraryRequirementsResponseOutput {
	return o
}

func (o LibraryRequirementsResponseOutput) ToLibraryRequirementsResponseOutputWithContext(ctx context.Context) LibraryRequirementsResponseOutput {
	return o
}

func (o LibraryRequirementsResponseOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LibraryRequirementsResponse) *string { return v.Content }).(pulumi.StringPtrOutput)
}

func (o LibraryRequirementsResponseOutput) Filename() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LibraryRequirementsResponse) *string { return v.Filename }).(pulumi.StringPtrOutput)
}

func (o LibraryRequirementsResponseOutput) Time() pulumi.StringOutput {
	return o.ApplyT(func(v LibraryRequirementsResponse) string { return v.Time }).(pulumi.StringOutput)
}

type LibraryRequirementsResponsePtrOutput struct{ *pulumi.OutputState }

func (LibraryRequirementsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LibraryRequirementsResponse)(nil)).Elem()
}

func (o LibraryRequirementsResponsePtrOutput) ToLibraryRequirementsResponsePtrOutput() LibraryRequirementsResponsePtrOutput {
	return o
}

func (o LibraryRequirementsResponsePtrOutput) ToLibraryRequirementsResponsePtrOutputWithContext(ctx context.Context) LibraryRequirementsResponsePtrOutput {
	return o
}

func (o LibraryRequirementsResponsePtrOutput) Elem() LibraryRequirementsResponseOutput {
	return o.ApplyT(func(v *LibraryRequirementsResponse) LibraryRequirementsResponse {
		if v != nil {
			return *v
		}
		var ret LibraryRequirementsResponse
		return ret
	}).(LibraryRequirementsResponseOutput)
}

func (o LibraryRequirementsResponsePtrOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LibraryRequirementsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Content
	}).(pulumi.StringPtrOutput)
}

func (o LibraryRequirementsResponsePtrOutput) Filename() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LibraryRequirementsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Filename
	}).(pulumi.StringPtrOutput)
}

func (o LibraryRequirementsResponsePtrOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LibraryRequirementsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Time
	}).(pulumi.StringPtrOutput)
}

type LinkedIntegrationRuntimeKeyAuthorization struct {
	AuthorizationType string       `pulumi:"authorizationType"`
	Key               SecureString `pulumi:"key"`
}

type LinkedIntegrationRuntimeKeyAuthorizationResponse struct {
	AuthorizationType string               `pulumi:"authorizationType"`
	Key               SecureStringResponse `pulumi:"key"`
}

type LinkedIntegrationRuntimeRbacAuthorization struct {
	AuthorizationType string `pulumi:"authorizationType"`
	ResourceId        string `pulumi:"resourceId"`
}

type LinkedIntegrationRuntimeRbacAuthorizationResponse struct {
	AuthorizationType string `pulumi:"authorizationType"`
	ResourceId        string `pulumi:"resourceId"`
}

type LinkedIntegrationRuntimeResponse struct {
	CreateTime          string `pulumi:"createTime"`
	DataFactoryLocation string `pulumi:"dataFactoryLocation"`
	DataFactoryName     string `pulumi:"dataFactoryName"`
	Name                string `pulumi:"name"`
	SubscriptionId      string `pulumi:"subscriptionId"`
}

type ManagedIdentity struct {
	Type                   *ResourceIdentityType  `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type ManagedIdentityInput interface {
	pulumi.Input

	ToManagedIdentityOutput() ManagedIdentityOutput
	ToManagedIdentityOutputWithContext(context.Context) ManagedIdentityOutput
}

type ManagedIdentityArgs struct {
	Type                   ResourceIdentityTypePtrInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput              `pulumi:"userAssignedIdentities"`
}

func (ManagedIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentity)(nil)).Elem()
}

func (i ManagedIdentityArgs) ToManagedIdentityOutput() ManagedIdentityOutput {
	return i.ToManagedIdentityOutputWithContext(context.Background())
}

func (i ManagedIdentityArgs) ToManagedIdentityOutputWithContext(ctx context.Context) ManagedIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentityOutput)
}

func (i ManagedIdentityArgs) ToManagedIdentityPtrOutput() ManagedIdentityPtrOutput {
	return i.ToManagedIdentityPtrOutputWithContext(context.Background())
}

func (i ManagedIdentityArgs) ToManagedIdentityPtrOutputWithContext(ctx context.Context) ManagedIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentityOutput).ToManagedIdentityPtrOutputWithContext(ctx)
}









type ManagedIdentityPtrInput interface {
	pulumi.Input

	ToManagedIdentityPtrOutput() ManagedIdentityPtrOutput
	ToManagedIdentityPtrOutputWithContext(context.Context) ManagedIdentityPtrOutput
}

type managedIdentityPtrType ManagedIdentityArgs

func ManagedIdentityPtr(v *ManagedIdentityArgs) ManagedIdentityPtrInput {
	return (*managedIdentityPtrType)(v)
}

func (*managedIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedIdentity)(nil)).Elem()
}

func (i *managedIdentityPtrType) ToManagedIdentityPtrOutput() ManagedIdentityPtrOutput {
	return i.ToManagedIdentityPtrOutputWithContext(context.Background())
}

func (i *managedIdentityPtrType) ToManagedIdentityPtrOutputWithContext(ctx context.Context) ManagedIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentityPtrOutput)
}

type ManagedIdentityOutput struct{ *pulumi.OutputState }

func (ManagedIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentity)(nil)).Elem()
}

func (o ManagedIdentityOutput) ToManagedIdentityOutput() ManagedIdentityOutput {
	return o
}

func (o ManagedIdentityOutput) ToManagedIdentityOutputWithContext(ctx context.Context) ManagedIdentityOutput {
	return o
}

func (o ManagedIdentityOutput) ToManagedIdentityPtrOutput() ManagedIdentityPtrOutput {
	return o.ToManagedIdentityPtrOutputWithContext(context.Background())
}

func (o ManagedIdentityOutput) ToManagedIdentityPtrOutputWithContext(ctx context.Context) ManagedIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedIdentity) *ManagedIdentity {
		return &v
	}).(ManagedIdentityPtrOutput)
}

func (o ManagedIdentityOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v ManagedIdentity) *ResourceIdentityType { return v.Type }).(ResourceIdentityTypePtrOutput)
}

func (o ManagedIdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v ManagedIdentity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type ManagedIdentityPtrOutput struct{ *pulumi.OutputState }

func (ManagedIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedIdentity)(nil)).Elem()
}

func (o ManagedIdentityPtrOutput) ToManagedIdentityPtrOutput() ManagedIdentityPtrOutput {
	return o
}

func (o ManagedIdentityPtrOutput) ToManagedIdentityPtrOutputWithContext(ctx context.Context) ManagedIdentityPtrOutput {
	return o
}

func (o ManagedIdentityPtrOutput) Elem() ManagedIdentityOutput {
	return o.ApplyT(func(v *ManagedIdentity) ManagedIdentity {
		if v != nil {
			return *v
		}
		var ret ManagedIdentity
		return ret
	}).(ManagedIdentityOutput)
}

func (o ManagedIdentityPtrOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v *ManagedIdentity) *ResourceIdentityType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(ResourceIdentityTypePtrOutput)
}

func (o ManagedIdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *ManagedIdentity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type ManagedIdentityResponse struct {
	PrincipalId            string                                         `pulumi:"principalId"`
	TenantId               string                                         `pulumi:"tenantId"`
	Type                   *string                                        `pulumi:"type"`
	UserAssignedIdentities map[string]UserAssignedManagedIdentityResponse `pulumi:"userAssignedIdentities"`
}

type ManagedIdentityResponseOutput struct{ *pulumi.OutputState }

func (ManagedIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentityResponse)(nil)).Elem()
}

func (o ManagedIdentityResponseOutput) ToManagedIdentityResponseOutput() ManagedIdentityResponseOutput {
	return o
}

func (o ManagedIdentityResponseOutput) ToManagedIdentityResponseOutputWithContext(ctx context.Context) ManagedIdentityResponseOutput {
	return o
}

func (o ManagedIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ManagedIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ManagedIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityResponseOutput) UserAssignedIdentities() UserAssignedManagedIdentityResponseMapOutput {
	return o.ApplyT(func(v ManagedIdentityResponse) map[string]UserAssignedManagedIdentityResponse {
		return v.UserAssignedIdentities
	}).(UserAssignedManagedIdentityResponseMapOutput)
}

type ManagedIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedIdentityResponse)(nil)).Elem()
}

func (o ManagedIdentityResponsePtrOutput) ToManagedIdentityResponsePtrOutput() ManagedIdentityResponsePtrOutput {
	return o
}

func (o ManagedIdentityResponsePtrOutput) ToManagedIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedIdentityResponsePtrOutput {
	return o
}

func (o ManagedIdentityResponsePtrOutput) Elem() ManagedIdentityResponseOutput {
	return o.ApplyT(func(v *ManagedIdentityResponse) ManagedIdentityResponse {
		if v != nil {
			return *v
		}
		var ret ManagedIdentityResponse
		return ret
	}).(ManagedIdentityResponseOutput)
}

func (o ManagedIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityResponsePtrOutput) UserAssignedIdentities() UserAssignedManagedIdentityResponseMapOutput {
	return o.ApplyT(func(v *ManagedIdentityResponse) map[string]UserAssignedManagedIdentityResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserAssignedManagedIdentityResponseMapOutput)
}

type ManagedIntegrationRuntime struct {
	ComputeProperties *IntegrationRuntimeComputeProperties `pulumi:"computeProperties"`
	Description       *string                              `pulumi:"description"`
	SsisProperties    *IntegrationRuntimeSsisProperties    `pulumi:"ssisProperties"`
	Type              string                               `pulumi:"type"`
}

type ManagedIntegrationRuntimeErrorResponse struct {
	Code       string   `pulumi:"code"`
	Message    string   `pulumi:"message"`
	Parameters []string `pulumi:"parameters"`
	Time       string   `pulumi:"time"`
}

type ManagedIntegrationRuntimeNodeResponse struct {
	Errors []ManagedIntegrationRuntimeErrorResponse `pulumi:"errors"`
	NodeId string                                   `pulumi:"nodeId"`
	Status string                                   `pulumi:"status"`
}

type ManagedIntegrationRuntimeOperationResultResponse struct {
	ActivityId string   `pulumi:"activityId"`
	ErrorCode  string   `pulumi:"errorCode"`
	Parameters []string `pulumi:"parameters"`
	Result     string   `pulumi:"result"`
	StartTime  string   `pulumi:"startTime"`
	Type       string   `pulumi:"type"`
}

type ManagedIntegrationRuntimeResponse struct {
	ComputeProperties *IntegrationRuntimeComputePropertiesResponse `pulumi:"computeProperties"`
	Description       *string                                      `pulumi:"description"`
	SsisProperties    *IntegrationRuntimeSsisPropertiesResponse    `pulumi:"ssisProperties"`
	State             string                                       `pulumi:"state"`
	Type              string                                       `pulumi:"type"`
}

type ManagedIntegrationRuntimeStatusResponse struct {
	CreateTime      string                                           `pulumi:"createTime"`
	DataFactoryName string                                           `pulumi:"dataFactoryName"`
	LastOperation   ManagedIntegrationRuntimeOperationResultResponse `pulumi:"lastOperation"`
	Nodes           []ManagedIntegrationRuntimeNodeResponse          `pulumi:"nodes"`
	OtherErrors     []ManagedIntegrationRuntimeErrorResponse         `pulumi:"otherErrors"`
	State           string                                           `pulumi:"state"`
	Type            string                                           `pulumi:"type"`
}

type ManagedVirtualNetworkSettings struct {
	AllowedAadTenantIdsForLinking     []string `pulumi:"allowedAadTenantIdsForLinking"`
	LinkedAccessCheckOnTargetResource *bool    `pulumi:"linkedAccessCheckOnTargetResource"`
	PreventDataExfiltration           *bool    `pulumi:"preventDataExfiltration"`
}





type ManagedVirtualNetworkSettingsInput interface {
	pulumi.Input

	ToManagedVirtualNetworkSettingsOutput() ManagedVirtualNetworkSettingsOutput
	ToManagedVirtualNetworkSettingsOutputWithContext(context.Context) ManagedVirtualNetworkSettingsOutput
}

type ManagedVirtualNetworkSettingsArgs struct {
	AllowedAadTenantIdsForLinking     pulumi.StringArrayInput `pulumi:"allowedAadTenantIdsForLinking"`
	LinkedAccessCheckOnTargetResource pulumi.BoolPtrInput     `pulumi:"linkedAccessCheckOnTargetResource"`
	PreventDataExfiltration           pulumi.BoolPtrInput     `pulumi:"preventDataExfiltration"`
}

func (ManagedVirtualNetworkSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedVirtualNetworkSettings)(nil)).Elem()
}

func (i ManagedVirtualNetworkSettingsArgs) ToManagedVirtualNetworkSettingsOutput() ManagedVirtualNetworkSettingsOutput {
	return i.ToManagedVirtualNetworkSettingsOutputWithContext(context.Background())
}

func (i ManagedVirtualNetworkSettingsArgs) ToManagedVirtualNetworkSettingsOutputWithContext(ctx context.Context) ManagedVirtualNetworkSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedVirtualNetworkSettingsOutput)
}

func (i ManagedVirtualNetworkSettingsArgs) ToManagedVirtualNetworkSettingsPtrOutput() ManagedVirtualNetworkSettingsPtrOutput {
	return i.ToManagedVirtualNetworkSettingsPtrOutputWithContext(context.Background())
}

func (i ManagedVirtualNetworkSettingsArgs) ToManagedVirtualNetworkSettingsPtrOutputWithContext(ctx context.Context) ManagedVirtualNetworkSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedVirtualNetworkSettingsOutput).ToManagedVirtualNetworkSettingsPtrOutputWithContext(ctx)
}









type ManagedVirtualNetworkSettingsPtrInput interface {
	pulumi.Input

	ToManagedVirtualNetworkSettingsPtrOutput() ManagedVirtualNetworkSettingsPtrOutput
	ToManagedVirtualNetworkSettingsPtrOutputWithContext(context.Context) ManagedVirtualNetworkSettingsPtrOutput
}

type managedVirtualNetworkSettingsPtrType ManagedVirtualNetworkSettingsArgs

func ManagedVirtualNetworkSettingsPtr(v *ManagedVirtualNetworkSettingsArgs) ManagedVirtualNetworkSettingsPtrInput {
	return (*managedVirtualNetworkSettingsPtrType)(v)
}

func (*managedVirtualNetworkSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedVirtualNetworkSettings)(nil)).Elem()
}

func (i *managedVirtualNetworkSettingsPtrType) ToManagedVirtualNetworkSettingsPtrOutput() ManagedVirtualNetworkSettingsPtrOutput {
	return i.ToManagedVirtualNetworkSettingsPtrOutputWithContext(context.Background())
}

func (i *managedVirtualNetworkSettingsPtrType) ToManagedVirtualNetworkSettingsPtrOutputWithContext(ctx context.Context) ManagedVirtualNetworkSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedVirtualNetworkSettingsPtrOutput)
}

type ManagedVirtualNetworkSettingsOutput struct{ *pulumi.OutputState }

func (ManagedVirtualNetworkSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedVirtualNetworkSettings)(nil)).Elem()
}

func (o ManagedVirtualNetworkSettingsOutput) ToManagedVirtualNetworkSettingsOutput() ManagedVirtualNetworkSettingsOutput {
	return o
}

func (o ManagedVirtualNetworkSettingsOutput) ToManagedVirtualNetworkSettingsOutputWithContext(ctx context.Context) ManagedVirtualNetworkSettingsOutput {
	return o
}

func (o ManagedVirtualNetworkSettingsOutput) ToManagedVirtualNetworkSettingsPtrOutput() ManagedVirtualNetworkSettingsPtrOutput {
	return o.ToManagedVirtualNetworkSettingsPtrOutputWithContext(context.Background())
}

func (o ManagedVirtualNetworkSettingsOutput) ToManagedVirtualNetworkSettingsPtrOutputWithContext(ctx context.Context) ManagedVirtualNetworkSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedVirtualNetworkSettings) *ManagedVirtualNetworkSettings {
		return &v
	}).(ManagedVirtualNetworkSettingsPtrOutput)
}

func (o ManagedVirtualNetworkSettingsOutput) AllowedAadTenantIdsForLinking() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ManagedVirtualNetworkSettings) []string { return v.AllowedAadTenantIdsForLinking }).(pulumi.StringArrayOutput)
}

func (o ManagedVirtualNetworkSettingsOutput) LinkedAccessCheckOnTargetResource() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ManagedVirtualNetworkSettings) *bool { return v.LinkedAccessCheckOnTargetResource }).(pulumi.BoolPtrOutput)
}

func (o ManagedVirtualNetworkSettingsOutput) PreventDataExfiltration() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ManagedVirtualNetworkSettings) *bool { return v.PreventDataExfiltration }).(pulumi.BoolPtrOutput)
}

type ManagedVirtualNetworkSettingsPtrOutput struct{ *pulumi.OutputState }

func (ManagedVirtualNetworkSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedVirtualNetworkSettings)(nil)).Elem()
}

func (o ManagedVirtualNetworkSettingsPtrOutput) ToManagedVirtualNetworkSettingsPtrOutput() ManagedVirtualNetworkSettingsPtrOutput {
	return o
}

func (o ManagedVirtualNetworkSettingsPtrOutput) ToManagedVirtualNetworkSettingsPtrOutputWithContext(ctx context.Context) ManagedVirtualNetworkSettingsPtrOutput {
	return o
}

func (o ManagedVirtualNetworkSettingsPtrOutput) Elem() ManagedVirtualNetworkSettingsOutput {
	return o.ApplyT(func(v *ManagedVirtualNetworkSettings) ManagedVirtualNetworkSettings {
		if v != nil {
			return *v
		}
		var ret ManagedVirtualNetworkSettings
		return ret
	}).(ManagedVirtualNetworkSettingsOutput)
}

func (o ManagedVirtualNetworkSettingsPtrOutput) AllowedAadTenantIdsForLinking() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ManagedVirtualNetworkSettings) []string {
		if v == nil {
			return nil
		}
		return v.AllowedAadTenantIdsForLinking
	}).(pulumi.StringArrayOutput)
}

func (o ManagedVirtualNetworkSettingsPtrOutput) LinkedAccessCheckOnTargetResource() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedVirtualNetworkSettings) *bool {
		if v == nil {
			return nil
		}
		return v.LinkedAccessCheckOnTargetResource
	}).(pulumi.BoolPtrOutput)
}

func (o ManagedVirtualNetworkSettingsPtrOutput) PreventDataExfiltration() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedVirtualNetworkSettings) *bool {
		if v == nil {
			return nil
		}
		return v.PreventDataExfiltration
	}).(pulumi.BoolPtrOutput)
}

type ManagedVirtualNetworkSettingsResponse struct {
	AllowedAadTenantIdsForLinking     []string `pulumi:"allowedAadTenantIdsForLinking"`
	LinkedAccessCheckOnTargetResource *bool    `pulumi:"linkedAccessCheckOnTargetResource"`
	PreventDataExfiltration           *bool    `pulumi:"preventDataExfiltration"`
}

type ManagedVirtualNetworkSettingsResponseOutput struct{ *pulumi.OutputState }

func (ManagedVirtualNetworkSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedVirtualNetworkSettingsResponse)(nil)).Elem()
}

func (o ManagedVirtualNetworkSettingsResponseOutput) ToManagedVirtualNetworkSettingsResponseOutput() ManagedVirtualNetworkSettingsResponseOutput {
	return o
}

func (o ManagedVirtualNetworkSettingsResponseOutput) ToManagedVirtualNetworkSettingsResponseOutputWithContext(ctx context.Context) ManagedVirtualNetworkSettingsResponseOutput {
	return o
}

func (o ManagedVirtualNetworkSettingsResponseOutput) AllowedAadTenantIdsForLinking() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ManagedVirtualNetworkSettingsResponse) []string { return v.AllowedAadTenantIdsForLinking }).(pulumi.StringArrayOutput)
}

func (o ManagedVirtualNetworkSettingsResponseOutput) LinkedAccessCheckOnTargetResource() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ManagedVirtualNetworkSettingsResponse) *bool { return v.LinkedAccessCheckOnTargetResource }).(pulumi.BoolPtrOutput)
}

func (o ManagedVirtualNetworkSettingsResponseOutput) PreventDataExfiltration() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ManagedVirtualNetworkSettingsResponse) *bool { return v.PreventDataExfiltration }).(pulumi.BoolPtrOutput)
}

type ManagedVirtualNetworkSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedVirtualNetworkSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedVirtualNetworkSettingsResponse)(nil)).Elem()
}

func (o ManagedVirtualNetworkSettingsResponsePtrOutput) ToManagedVirtualNetworkSettingsResponsePtrOutput() ManagedVirtualNetworkSettingsResponsePtrOutput {
	return o
}

func (o ManagedVirtualNetworkSettingsResponsePtrOutput) ToManagedVirtualNetworkSettingsResponsePtrOutputWithContext(ctx context.Context) ManagedVirtualNetworkSettingsResponsePtrOutput {
	return o
}

func (o ManagedVirtualNetworkSettingsResponsePtrOutput) Elem() ManagedVirtualNetworkSettingsResponseOutput {
	return o.ApplyT(func(v *ManagedVirtualNetworkSettingsResponse) ManagedVirtualNetworkSettingsResponse {
		if v != nil {
			return *v
		}
		var ret ManagedVirtualNetworkSettingsResponse
		return ret
	}).(ManagedVirtualNetworkSettingsResponseOutput)
}

func (o ManagedVirtualNetworkSettingsResponsePtrOutput) AllowedAadTenantIdsForLinking() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ManagedVirtualNetworkSettingsResponse) []string {
		if v == nil {
			return nil
		}
		return v.AllowedAadTenantIdsForLinking
	}).(pulumi.StringArrayOutput)
}

func (o ManagedVirtualNetworkSettingsResponsePtrOutput) LinkedAccessCheckOnTargetResource() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedVirtualNetworkSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.LinkedAccessCheckOnTargetResource
	}).(pulumi.BoolPtrOutput)
}

func (o ManagedVirtualNetworkSettingsResponsePtrOutput) PreventDataExfiltration() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedVirtualNetworkSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.PreventDataExfiltration
	}).(pulumi.BoolPtrOutput)
}

type OptimizedAutoscale struct {
	IsEnabled bool `pulumi:"isEnabled"`
	Maximum   int  `pulumi:"maximum"`
	Minimum   int  `pulumi:"minimum"`
	Version   int  `pulumi:"version"`
}





type OptimizedAutoscaleInput interface {
	pulumi.Input

	ToOptimizedAutoscaleOutput() OptimizedAutoscaleOutput
	ToOptimizedAutoscaleOutputWithContext(context.Context) OptimizedAutoscaleOutput
}

type OptimizedAutoscaleArgs struct {
	IsEnabled pulumi.BoolInput `pulumi:"isEnabled"`
	Maximum   pulumi.IntInput  `pulumi:"maximum"`
	Minimum   pulumi.IntInput  `pulumi:"minimum"`
	Version   pulumi.IntInput  `pulumi:"version"`
}

func (OptimizedAutoscaleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OptimizedAutoscale)(nil)).Elem()
}

func (i OptimizedAutoscaleArgs) ToOptimizedAutoscaleOutput() OptimizedAutoscaleOutput {
	return i.ToOptimizedAutoscaleOutputWithContext(context.Background())
}

func (i OptimizedAutoscaleArgs) ToOptimizedAutoscaleOutputWithContext(ctx context.Context) OptimizedAutoscaleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OptimizedAutoscaleOutput)
}

func (i OptimizedAutoscaleArgs) ToOptimizedAutoscalePtrOutput() OptimizedAutoscalePtrOutput {
	return i.ToOptimizedAutoscalePtrOutputWithContext(context.Background())
}

func (i OptimizedAutoscaleArgs) ToOptimizedAutoscalePtrOutputWithContext(ctx context.Context) OptimizedAutoscalePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OptimizedAutoscaleOutput).ToOptimizedAutoscalePtrOutputWithContext(ctx)
}









type OptimizedAutoscalePtrInput interface {
	pulumi.Input

	ToOptimizedAutoscalePtrOutput() OptimizedAutoscalePtrOutput
	ToOptimizedAutoscalePtrOutputWithContext(context.Context) OptimizedAutoscalePtrOutput
}

type optimizedAutoscalePtrType OptimizedAutoscaleArgs

func OptimizedAutoscalePtr(v *OptimizedAutoscaleArgs) OptimizedAutoscalePtrInput {
	return (*optimizedAutoscalePtrType)(v)
}

func (*optimizedAutoscalePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OptimizedAutoscale)(nil)).Elem()
}

func (i *optimizedAutoscalePtrType) ToOptimizedAutoscalePtrOutput() OptimizedAutoscalePtrOutput {
	return i.ToOptimizedAutoscalePtrOutputWithContext(context.Background())
}

func (i *optimizedAutoscalePtrType) ToOptimizedAutoscalePtrOutputWithContext(ctx context.Context) OptimizedAutoscalePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OptimizedAutoscalePtrOutput)
}

type OptimizedAutoscaleOutput struct{ *pulumi.OutputState }

func (OptimizedAutoscaleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OptimizedAutoscale)(nil)).Elem()
}

func (o OptimizedAutoscaleOutput) ToOptimizedAutoscaleOutput() OptimizedAutoscaleOutput {
	return o
}

func (o OptimizedAutoscaleOutput) ToOptimizedAutoscaleOutputWithContext(ctx context.Context) OptimizedAutoscaleOutput {
	return o
}

func (o OptimizedAutoscaleOutput) ToOptimizedAutoscalePtrOutput() OptimizedAutoscalePtrOutput {
	return o.ToOptimizedAutoscalePtrOutputWithContext(context.Background())
}

func (o OptimizedAutoscaleOutput) ToOptimizedAutoscalePtrOutputWithContext(ctx context.Context) OptimizedAutoscalePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OptimizedAutoscale) *OptimizedAutoscale {
		return &v
	}).(OptimizedAutoscalePtrOutput)
}

func (o OptimizedAutoscaleOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v OptimizedAutoscale) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o OptimizedAutoscaleOutput) Maximum() pulumi.IntOutput {
	return o.ApplyT(func(v OptimizedAutoscale) int { return v.Maximum }).(pulumi.IntOutput)
}

func (o OptimizedAutoscaleOutput) Minimum() pulumi.IntOutput {
	return o.ApplyT(func(v OptimizedAutoscale) int { return v.Minimum }).(pulumi.IntOutput)
}

func (o OptimizedAutoscaleOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v OptimizedAutoscale) int { return v.Version }).(pulumi.IntOutput)
}

type OptimizedAutoscalePtrOutput struct{ *pulumi.OutputState }

func (OptimizedAutoscalePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OptimizedAutoscale)(nil)).Elem()
}

func (o OptimizedAutoscalePtrOutput) ToOptimizedAutoscalePtrOutput() OptimizedAutoscalePtrOutput {
	return o
}

func (o OptimizedAutoscalePtrOutput) ToOptimizedAutoscalePtrOutputWithContext(ctx context.Context) OptimizedAutoscalePtrOutput {
	return o
}

func (o OptimizedAutoscalePtrOutput) Elem() OptimizedAutoscaleOutput {
	return o.ApplyT(func(v *OptimizedAutoscale) OptimizedAutoscale {
		if v != nil {
			return *v
		}
		var ret OptimizedAutoscale
		return ret
	}).(OptimizedAutoscaleOutput)
}

func (o OptimizedAutoscalePtrOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OptimizedAutoscale) *bool {
		if v == nil {
			return nil
		}
		return &v.IsEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o OptimizedAutoscalePtrOutput) Maximum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OptimizedAutoscale) *int {
		if v == nil {
			return nil
		}
		return &v.Maximum
	}).(pulumi.IntPtrOutput)
}

func (o OptimizedAutoscalePtrOutput) Minimum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OptimizedAutoscale) *int {
		if v == nil {
			return nil
		}
		return &v.Minimum
	}).(pulumi.IntPtrOutput)
}

func (o OptimizedAutoscalePtrOutput) Version() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OptimizedAutoscale) *int {
		if v == nil {
			return nil
		}
		return &v.Version
	}).(pulumi.IntPtrOutput)
}

type OptimizedAutoscaleResponse struct {
	IsEnabled bool `pulumi:"isEnabled"`
	Maximum   int  `pulumi:"maximum"`
	Minimum   int  `pulumi:"minimum"`
	Version   int  `pulumi:"version"`
}

type OptimizedAutoscaleResponseOutput struct{ *pulumi.OutputState }

func (OptimizedAutoscaleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OptimizedAutoscaleResponse)(nil)).Elem()
}

func (o OptimizedAutoscaleResponseOutput) ToOptimizedAutoscaleResponseOutput() OptimizedAutoscaleResponseOutput {
	return o
}

func (o OptimizedAutoscaleResponseOutput) ToOptimizedAutoscaleResponseOutputWithContext(ctx context.Context) OptimizedAutoscaleResponseOutput {
	return o
}

func (o OptimizedAutoscaleResponseOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v OptimizedAutoscaleResponse) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o OptimizedAutoscaleResponseOutput) Maximum() pulumi.IntOutput {
	return o.ApplyT(func(v OptimizedAutoscaleResponse) int { return v.Maximum }).(pulumi.IntOutput)
}

func (o OptimizedAutoscaleResponseOutput) Minimum() pulumi.IntOutput {
	return o.ApplyT(func(v OptimizedAutoscaleResponse) int { return v.Minimum }).(pulumi.IntOutput)
}

func (o OptimizedAutoscaleResponseOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v OptimizedAutoscaleResponse) int { return v.Version }).(pulumi.IntOutput)
}

type OptimizedAutoscaleResponsePtrOutput struct{ *pulumi.OutputState }

func (OptimizedAutoscaleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OptimizedAutoscaleResponse)(nil)).Elem()
}

func (o OptimizedAutoscaleResponsePtrOutput) ToOptimizedAutoscaleResponsePtrOutput() OptimizedAutoscaleResponsePtrOutput {
	return o
}

func (o OptimizedAutoscaleResponsePtrOutput) ToOptimizedAutoscaleResponsePtrOutputWithContext(ctx context.Context) OptimizedAutoscaleResponsePtrOutput {
	return o
}

func (o OptimizedAutoscaleResponsePtrOutput) Elem() OptimizedAutoscaleResponseOutput {
	return o.ApplyT(func(v *OptimizedAutoscaleResponse) OptimizedAutoscaleResponse {
		if v != nil {
			return *v
		}
		var ret OptimizedAutoscaleResponse
		return ret
	}).(OptimizedAutoscaleResponseOutput)
}

func (o OptimizedAutoscaleResponsePtrOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OptimizedAutoscaleResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.IsEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o OptimizedAutoscaleResponsePtrOutput) Maximum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OptimizedAutoscaleResponse) *int {
		if v == nil {
			return nil
		}
		return &v.Maximum
	}).(pulumi.IntPtrOutput)
}

func (o OptimizedAutoscaleResponsePtrOutput) Minimum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OptimizedAutoscaleResponse) *int {
		if v == nil {
			return nil
		}
		return &v.Minimum
	}).(pulumi.IntPtrOutput)
}

func (o OptimizedAutoscaleResponsePtrOutput) Version() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OptimizedAutoscaleResponse) *int {
		if v == nil {
			return nil
		}
		return &v.Version
	}).(pulumi.IntPtrOutput)
}

type PrivateEndpointConnectionType struct {
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
}





type PrivateEndpointConnectionTypeInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionTypeOutput() PrivateEndpointConnectionTypeOutput
	ToPrivateEndpointConnectionTypeOutputWithContext(context.Context) PrivateEndpointConnectionTypeOutput
}

type PrivateEndpointConnectionTypeArgs struct {
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStatePtrInput `pulumi:"privateLinkServiceConnectionState"`
}

func (PrivateEndpointConnectionTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionType)(nil)).Elem()
}

func (i PrivateEndpointConnectionTypeArgs) ToPrivateEndpointConnectionTypeOutput() PrivateEndpointConnectionTypeOutput {
	return i.ToPrivateEndpointConnectionTypeOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionTypeArgs) ToPrivateEndpointConnectionTypeOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionTypeOutput)
}





type PrivateEndpointConnectionTypeArrayInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionTypeArrayOutput() PrivateEndpointConnectionTypeArrayOutput
	ToPrivateEndpointConnectionTypeArrayOutputWithContext(context.Context) PrivateEndpointConnectionTypeArrayOutput
}

type PrivateEndpointConnectionTypeArray []PrivateEndpointConnectionTypeInput

func (PrivateEndpointConnectionTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionType)(nil)).Elem()
}

func (i PrivateEndpointConnectionTypeArray) ToPrivateEndpointConnectionTypeArrayOutput() PrivateEndpointConnectionTypeArrayOutput {
	return i.ToPrivateEndpointConnectionTypeArrayOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionTypeArray) ToPrivateEndpointConnectionTypeArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionTypeArrayOutput)
}

type PrivateEndpointConnectionTypeOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionType)(nil)).Elem()
}

func (o PrivateEndpointConnectionTypeOutput) ToPrivateEndpointConnectionTypeOutput() PrivateEndpointConnectionTypeOutput {
	return o
}

func (o PrivateEndpointConnectionTypeOutput) ToPrivateEndpointConnectionTypeOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeOutput {
	return o
}

func (o PrivateEndpointConnectionTypeOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionType) *PrivateLinkServiceConnectionState {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStatePtrOutput)
}

type PrivateEndpointConnectionTypeArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionType)(nil)).Elem()
}

func (o PrivateEndpointConnectionTypeArrayOutput) ToPrivateEndpointConnectionTypeArrayOutput() PrivateEndpointConnectionTypeArrayOutput {
	return o
}

func (o PrivateEndpointConnectionTypeArrayOutput) ToPrivateEndpointConnectionTypeArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeArrayOutput {
	return o
}

func (o PrivateEndpointConnectionTypeArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionType {
		return vs[0].([]PrivateEndpointConnectionType)[vs[1].(int)]
	}).(PrivateEndpointConnectionTypeOutput)
}

type PrivateEndpointConnectionForPrivateLinkHubBasicResponse struct {
	Id                                string                                     `pulumi:"id"`
	PrivateEndpoint                   *PrivateEndpointResponse                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                     `pulumi:"provisioningState"`
}

type PrivateEndpointConnectionForPrivateLinkHubBasicResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionForPrivateLinkHubBasicResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionForPrivateLinkHubBasicResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionForPrivateLinkHubBasicResponseOutput) ToPrivateEndpointConnectionForPrivateLinkHubBasicResponseOutput() PrivateEndpointConnectionForPrivateLinkHubBasicResponseOutput {
	return o
}

func (o PrivateEndpointConnectionForPrivateLinkHubBasicResponseOutput) ToPrivateEndpointConnectionForPrivateLinkHubBasicResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionForPrivateLinkHubBasicResponseOutput {
	return o
}

func (o PrivateEndpointConnectionForPrivateLinkHubBasicResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionForPrivateLinkHubBasicResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionForPrivateLinkHubBasicResponseOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionForPrivateLinkHubBasicResponse) *PrivateEndpointResponse {
		return v.PrivateEndpoint
	}).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointConnectionForPrivateLinkHubBasicResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionForPrivateLinkHubBasicResponse) *PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

func (o PrivateEndpointConnectionForPrivateLinkHubBasicResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionForPrivateLinkHubBasicResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type PrivateEndpointConnectionForPrivateLinkHubBasicResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionForPrivateLinkHubBasicResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionForPrivateLinkHubBasicResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionForPrivateLinkHubBasicResponseArrayOutput) ToPrivateEndpointConnectionForPrivateLinkHubBasicResponseArrayOutput() PrivateEndpointConnectionForPrivateLinkHubBasicResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionForPrivateLinkHubBasicResponseArrayOutput) ToPrivateEndpointConnectionForPrivateLinkHubBasicResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionForPrivateLinkHubBasicResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionForPrivateLinkHubBasicResponseArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionForPrivateLinkHubBasicResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionForPrivateLinkHubBasicResponse {
		return vs[0].([]PrivateEndpointConnectionForPrivateLinkHubBasicResponse)[vs[1].(int)]
	}).(PrivateEndpointConnectionForPrivateLinkHubBasicResponseOutput)
}

type PrivateEndpointConnectionResponse struct {
	Id                                string                                     `pulumi:"id"`
	Name                              string                                     `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                     `pulumi:"provisioningState"`
	Type                              string                                     `pulumi:"type"`
}

type PrivateEndpointConnectionResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type PrivateEndpointConnectionResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionResponse {
		return vs[0].([]PrivateEndpointConnectionResponse)[vs[1].(int)]
	}).(PrivateEndpointConnectionResponseOutput)
}

type PrivateEndpointResponse struct {
	Id string `pulumi:"id"`
}

type PrivateEndpointResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointResponse) string { return v.Id }).(pulumi.StringOutput)
}

type PrivateEndpointResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) Elem() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) PrivateEndpointResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointResponse
		return ret
	}).(PrivateEndpointResponseOutput)
}

func (o PrivateEndpointResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionState struct {
	Description *string `pulumi:"description"`
	Status      *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput
	ToPrivateLinkServiceConnectionStateOutputWithContext(context.Context) PrivateLinkServiceConnectionStateOutput
}

type PrivateLinkServiceConnectionStateArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Status      pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return i.ToPrivateLinkServiceConnectionStateOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput)
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput).ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStatePtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput
	ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePtrOutput
}

type privateLinkServiceConnectionStatePtrType PrivateLinkServiceConnectionStateArgs

func PrivateLinkServiceConnectionStatePtr(v *PrivateLinkServiceConnectionStateArgs) PrivateLinkServiceConnectionStatePtrInput {
	return (*privateLinkServiceConnectionStatePtrType)(v)
}

func (*privateLinkServiceConnectionStatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStatePtrType) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStatePtrType) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePtrOutput)
}

type PrivateLinkServiceConnectionStateOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return o.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionState) *PrivateLinkServiceConnectionState {
		return &v
	}).(PrivateLinkServiceConnectionStatePtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStatePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Elem() PrivateLinkServiceConnectionStateOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) PrivateLinkServiceConnectionState {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionState
		return ret
	}).(PrivateLinkServiceConnectionStateOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponse struct {
	ActionsRequired string  `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}

type PrivateLinkServiceConnectionStateResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ActionsRequired() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) string { return v.ActionsRequired }).(pulumi.StringOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Elem() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) PrivateLinkServiceConnectionStateResponse {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionStateResponse
		return ret
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type PurviewConfiguration struct {
	PurviewResourceId *string `pulumi:"purviewResourceId"`
}





type PurviewConfigurationInput interface {
	pulumi.Input

	ToPurviewConfigurationOutput() PurviewConfigurationOutput
	ToPurviewConfigurationOutputWithContext(context.Context) PurviewConfigurationOutput
}

type PurviewConfigurationArgs struct {
	PurviewResourceId pulumi.StringPtrInput `pulumi:"purviewResourceId"`
}

func (PurviewConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PurviewConfiguration)(nil)).Elem()
}

func (i PurviewConfigurationArgs) ToPurviewConfigurationOutput() PurviewConfigurationOutput {
	return i.ToPurviewConfigurationOutputWithContext(context.Background())
}

func (i PurviewConfigurationArgs) ToPurviewConfigurationOutputWithContext(ctx context.Context) PurviewConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PurviewConfigurationOutput)
}

func (i PurviewConfigurationArgs) ToPurviewConfigurationPtrOutput() PurviewConfigurationPtrOutput {
	return i.ToPurviewConfigurationPtrOutputWithContext(context.Background())
}

func (i PurviewConfigurationArgs) ToPurviewConfigurationPtrOutputWithContext(ctx context.Context) PurviewConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PurviewConfigurationOutput).ToPurviewConfigurationPtrOutputWithContext(ctx)
}









type PurviewConfigurationPtrInput interface {
	pulumi.Input

	ToPurviewConfigurationPtrOutput() PurviewConfigurationPtrOutput
	ToPurviewConfigurationPtrOutputWithContext(context.Context) PurviewConfigurationPtrOutput
}

type purviewConfigurationPtrType PurviewConfigurationArgs

func PurviewConfigurationPtr(v *PurviewConfigurationArgs) PurviewConfigurationPtrInput {
	return (*purviewConfigurationPtrType)(v)
}

func (*purviewConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PurviewConfiguration)(nil)).Elem()
}

func (i *purviewConfigurationPtrType) ToPurviewConfigurationPtrOutput() PurviewConfigurationPtrOutput {
	return i.ToPurviewConfigurationPtrOutputWithContext(context.Background())
}

func (i *purviewConfigurationPtrType) ToPurviewConfigurationPtrOutputWithContext(ctx context.Context) PurviewConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PurviewConfigurationPtrOutput)
}

type PurviewConfigurationOutput struct{ *pulumi.OutputState }

func (PurviewConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PurviewConfiguration)(nil)).Elem()
}

func (o PurviewConfigurationOutput) ToPurviewConfigurationOutput() PurviewConfigurationOutput {
	return o
}

func (o PurviewConfigurationOutput) ToPurviewConfigurationOutputWithContext(ctx context.Context) PurviewConfigurationOutput {
	return o
}

func (o PurviewConfigurationOutput) ToPurviewConfigurationPtrOutput() PurviewConfigurationPtrOutput {
	return o.ToPurviewConfigurationPtrOutputWithContext(context.Background())
}

func (o PurviewConfigurationOutput) ToPurviewConfigurationPtrOutputWithContext(ctx context.Context) PurviewConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PurviewConfiguration) *PurviewConfiguration {
		return &v
	}).(PurviewConfigurationPtrOutput)
}

func (o PurviewConfigurationOutput) PurviewResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PurviewConfiguration) *string { return v.PurviewResourceId }).(pulumi.StringPtrOutput)
}

type PurviewConfigurationPtrOutput struct{ *pulumi.OutputState }

func (PurviewConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PurviewConfiguration)(nil)).Elem()
}

func (o PurviewConfigurationPtrOutput) ToPurviewConfigurationPtrOutput() PurviewConfigurationPtrOutput {
	return o
}

func (o PurviewConfigurationPtrOutput) ToPurviewConfigurationPtrOutputWithContext(ctx context.Context) PurviewConfigurationPtrOutput {
	return o
}

func (o PurviewConfigurationPtrOutput) Elem() PurviewConfigurationOutput {
	return o.ApplyT(func(v *PurviewConfiguration) PurviewConfiguration {
		if v != nil {
			return *v
		}
		var ret PurviewConfiguration
		return ret
	}).(PurviewConfigurationOutput)
}

func (o PurviewConfigurationPtrOutput) PurviewResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PurviewConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.PurviewResourceId
	}).(pulumi.StringPtrOutput)
}

type PurviewConfigurationResponse struct {
	PurviewResourceId *string `pulumi:"purviewResourceId"`
}

type PurviewConfigurationResponseOutput struct{ *pulumi.OutputState }

func (PurviewConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PurviewConfigurationResponse)(nil)).Elem()
}

func (o PurviewConfigurationResponseOutput) ToPurviewConfigurationResponseOutput() PurviewConfigurationResponseOutput {
	return o
}

func (o PurviewConfigurationResponseOutput) ToPurviewConfigurationResponseOutputWithContext(ctx context.Context) PurviewConfigurationResponseOutput {
	return o
}

func (o PurviewConfigurationResponseOutput) PurviewResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PurviewConfigurationResponse) *string { return v.PurviewResourceId }).(pulumi.StringPtrOutput)
}

type PurviewConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (PurviewConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PurviewConfigurationResponse)(nil)).Elem()
}

func (o PurviewConfigurationResponsePtrOutput) ToPurviewConfigurationResponsePtrOutput() PurviewConfigurationResponsePtrOutput {
	return o
}

func (o PurviewConfigurationResponsePtrOutput) ToPurviewConfigurationResponsePtrOutputWithContext(ctx context.Context) PurviewConfigurationResponsePtrOutput {
	return o
}

func (o PurviewConfigurationResponsePtrOutput) Elem() PurviewConfigurationResponseOutput {
	return o.ApplyT(func(v *PurviewConfigurationResponse) PurviewConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret PurviewConfigurationResponse
		return ret
	}).(PurviewConfigurationResponseOutput)
}

func (o PurviewConfigurationResponsePtrOutput) PurviewResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PurviewConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.PurviewResourceId
	}).(pulumi.StringPtrOutput)
}

type SecureString struct {
	Type  string `pulumi:"type"`
	Value string `pulumi:"value"`
}

type SecureStringResponse struct {
	Type  string `pulumi:"type"`
	Value string `pulumi:"value"`
}

type SelfHostedIntegrationRuntime struct {
	Description *string     `pulumi:"description"`
	LinkedInfo  interface{} `pulumi:"linkedInfo"`
	Type        string      `pulumi:"type"`
}

type SelfHostedIntegrationRuntimeNodeResponse struct {
	Capabilities        map[string]string `pulumi:"capabilities"`
	ConcurrentJobsLimit int               `pulumi:"concurrentJobsLimit"`
	ExpiryTime          string            `pulumi:"expiryTime"`
	HostServiceUri      string            `pulumi:"hostServiceUri"`
	IsActiveDispatcher  bool              `pulumi:"isActiveDispatcher"`
	LastConnectTime     string            `pulumi:"lastConnectTime"`
	LastEndUpdateTime   string            `pulumi:"lastEndUpdateTime"`
	LastStartTime       string            `pulumi:"lastStartTime"`
	LastStartUpdateTime string            `pulumi:"lastStartUpdateTime"`
	LastStopTime        string            `pulumi:"lastStopTime"`
	LastUpdateResult    string            `pulumi:"lastUpdateResult"`
	MachineName         string            `pulumi:"machineName"`
	MaxConcurrentJobs   int               `pulumi:"maxConcurrentJobs"`
	NodeName            string            `pulumi:"nodeName"`
	RegisterTime        string            `pulumi:"registerTime"`
	Status              string            `pulumi:"status"`
	Version             string            `pulumi:"version"`
	VersionStatus       string            `pulumi:"versionStatus"`
}

type SelfHostedIntegrationRuntimeResponse struct {
	Description *string     `pulumi:"description"`
	LinkedInfo  interface{} `pulumi:"linkedInfo"`
	Type        string      `pulumi:"type"`
}

type SelfHostedIntegrationRuntimeStatusResponse struct {
	AutoUpdate                             string                                     `pulumi:"autoUpdate"`
	AutoUpdateETA                          string                                     `pulumi:"autoUpdateETA"`
	Capabilities                           map[string]string                          `pulumi:"capabilities"`
	CreateTime                             string                                     `pulumi:"createTime"`
	DataFactoryName                        string                                     `pulumi:"dataFactoryName"`
	InternalChannelEncryption              string                                     `pulumi:"internalChannelEncryption"`
	LatestVersion                          string                                     `pulumi:"latestVersion"`
	Links                                  []LinkedIntegrationRuntimeResponse         `pulumi:"links"`
	LocalTimeZoneOffset                    string                                     `pulumi:"localTimeZoneOffset"`
	NodeCommunicationChannelEncryptionMode string                                     `pulumi:"nodeCommunicationChannelEncryptionMode"`
	Nodes                                  []SelfHostedIntegrationRuntimeNodeResponse `pulumi:"nodes"`
	PushedVersion                          string                                     `pulumi:"pushedVersion"`
	ScheduledUpdateDate                    string                                     `pulumi:"scheduledUpdateDate"`
	ServiceUrls                            []string                                   `pulumi:"serviceUrls"`
	State                                  string                                     `pulumi:"state"`
	TaskQueueId                            string                                     `pulumi:"taskQueueId"`
	Type                                   string                                     `pulumi:"type"`
	UpdateDelayOffset                      string                                     `pulumi:"updateDelayOffset"`
	Version                                string                                     `pulumi:"version"`
	VersionStatus                          string                                     `pulumi:"versionStatus"`
}

type Sku struct {
	Capacity *int    `pulumi:"capacity"`
	Name     *string `pulumi:"name"`
	Tier     *string `pulumi:"tier"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
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

func (o SkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Sku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Tier }).(pulumi.StringPtrOutput)
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

func (o SkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Sku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Capacity *int    `pulumi:"capacity"`
	Name     *string `pulumi:"name"`
	Tier     *string `pulumi:"tier"`
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

func (o SkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
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

func (o SkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SparkConfigProperties struct {
	ConfigurationType *string `pulumi:"configurationType"`
	Content           *string `pulumi:"content"`
	Filename          *string `pulumi:"filename"`
}





type SparkConfigPropertiesInput interface {
	pulumi.Input

	ToSparkConfigPropertiesOutput() SparkConfigPropertiesOutput
	ToSparkConfigPropertiesOutputWithContext(context.Context) SparkConfigPropertiesOutput
}

type SparkConfigPropertiesArgs struct {
	ConfigurationType pulumi.StringPtrInput `pulumi:"configurationType"`
	Content           pulumi.StringPtrInput `pulumi:"content"`
	Filename          pulumi.StringPtrInput `pulumi:"filename"`
}

func (SparkConfigPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SparkConfigProperties)(nil)).Elem()
}

func (i SparkConfigPropertiesArgs) ToSparkConfigPropertiesOutput() SparkConfigPropertiesOutput {
	return i.ToSparkConfigPropertiesOutputWithContext(context.Background())
}

func (i SparkConfigPropertiesArgs) ToSparkConfigPropertiesOutputWithContext(ctx context.Context) SparkConfigPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SparkConfigPropertiesOutput)
}

func (i SparkConfigPropertiesArgs) ToSparkConfigPropertiesPtrOutput() SparkConfigPropertiesPtrOutput {
	return i.ToSparkConfigPropertiesPtrOutputWithContext(context.Background())
}

func (i SparkConfigPropertiesArgs) ToSparkConfigPropertiesPtrOutputWithContext(ctx context.Context) SparkConfigPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SparkConfigPropertiesOutput).ToSparkConfigPropertiesPtrOutputWithContext(ctx)
}









type SparkConfigPropertiesPtrInput interface {
	pulumi.Input

	ToSparkConfigPropertiesPtrOutput() SparkConfigPropertiesPtrOutput
	ToSparkConfigPropertiesPtrOutputWithContext(context.Context) SparkConfigPropertiesPtrOutput
}

type sparkConfigPropertiesPtrType SparkConfigPropertiesArgs

func SparkConfigPropertiesPtr(v *SparkConfigPropertiesArgs) SparkConfigPropertiesPtrInput {
	return (*sparkConfigPropertiesPtrType)(v)
}

func (*sparkConfigPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SparkConfigProperties)(nil)).Elem()
}

func (i *sparkConfigPropertiesPtrType) ToSparkConfigPropertiesPtrOutput() SparkConfigPropertiesPtrOutput {
	return i.ToSparkConfigPropertiesPtrOutputWithContext(context.Background())
}

func (i *sparkConfigPropertiesPtrType) ToSparkConfigPropertiesPtrOutputWithContext(ctx context.Context) SparkConfigPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SparkConfigPropertiesPtrOutput)
}

type SparkConfigPropertiesOutput struct{ *pulumi.OutputState }

func (SparkConfigPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SparkConfigProperties)(nil)).Elem()
}

func (o SparkConfigPropertiesOutput) ToSparkConfigPropertiesOutput() SparkConfigPropertiesOutput {
	return o
}

func (o SparkConfigPropertiesOutput) ToSparkConfigPropertiesOutputWithContext(ctx context.Context) SparkConfigPropertiesOutput {
	return o
}

func (o SparkConfigPropertiesOutput) ToSparkConfigPropertiesPtrOutput() SparkConfigPropertiesPtrOutput {
	return o.ToSparkConfigPropertiesPtrOutputWithContext(context.Background())
}

func (o SparkConfigPropertiesOutput) ToSparkConfigPropertiesPtrOutputWithContext(ctx context.Context) SparkConfigPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SparkConfigProperties) *SparkConfigProperties {
		return &v
	}).(SparkConfigPropertiesPtrOutput)
}

func (o SparkConfigPropertiesOutput) ConfigurationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SparkConfigProperties) *string { return v.ConfigurationType }).(pulumi.StringPtrOutput)
}

func (o SparkConfigPropertiesOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SparkConfigProperties) *string { return v.Content }).(pulumi.StringPtrOutput)
}

func (o SparkConfigPropertiesOutput) Filename() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SparkConfigProperties) *string { return v.Filename }).(pulumi.StringPtrOutput)
}

type SparkConfigPropertiesPtrOutput struct{ *pulumi.OutputState }

func (SparkConfigPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SparkConfigProperties)(nil)).Elem()
}

func (o SparkConfigPropertiesPtrOutput) ToSparkConfigPropertiesPtrOutput() SparkConfigPropertiesPtrOutput {
	return o
}

func (o SparkConfigPropertiesPtrOutput) ToSparkConfigPropertiesPtrOutputWithContext(ctx context.Context) SparkConfigPropertiesPtrOutput {
	return o
}

func (o SparkConfigPropertiesPtrOutput) Elem() SparkConfigPropertiesOutput {
	return o.ApplyT(func(v *SparkConfigProperties) SparkConfigProperties {
		if v != nil {
			return *v
		}
		var ret SparkConfigProperties
		return ret
	}).(SparkConfigPropertiesOutput)
}

func (o SparkConfigPropertiesPtrOutput) ConfigurationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SparkConfigProperties) *string {
		if v == nil {
			return nil
		}
		return v.ConfigurationType
	}).(pulumi.StringPtrOutput)
}

func (o SparkConfigPropertiesPtrOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SparkConfigProperties) *string {
		if v == nil {
			return nil
		}
		return v.Content
	}).(pulumi.StringPtrOutput)
}

func (o SparkConfigPropertiesPtrOutput) Filename() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SparkConfigProperties) *string {
		if v == nil {
			return nil
		}
		return v.Filename
	}).(pulumi.StringPtrOutput)
}

type SparkConfigPropertiesResponse struct {
	ConfigurationType *string `pulumi:"configurationType"`
	Content           *string `pulumi:"content"`
	Filename          *string `pulumi:"filename"`
	Time              string  `pulumi:"time"`
}

type SparkConfigPropertiesResponseOutput struct{ *pulumi.OutputState }

func (SparkConfigPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SparkConfigPropertiesResponse)(nil)).Elem()
}

func (o SparkConfigPropertiesResponseOutput) ToSparkConfigPropertiesResponseOutput() SparkConfigPropertiesResponseOutput {
	return o
}

func (o SparkConfigPropertiesResponseOutput) ToSparkConfigPropertiesResponseOutputWithContext(ctx context.Context) SparkConfigPropertiesResponseOutput {
	return o
}

func (o SparkConfigPropertiesResponseOutput) ConfigurationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SparkConfigPropertiesResponse) *string { return v.ConfigurationType }).(pulumi.StringPtrOutput)
}

func (o SparkConfigPropertiesResponseOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SparkConfigPropertiesResponse) *string { return v.Content }).(pulumi.StringPtrOutput)
}

func (o SparkConfigPropertiesResponseOutput) Filename() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SparkConfigPropertiesResponse) *string { return v.Filename }).(pulumi.StringPtrOutput)
}

func (o SparkConfigPropertiesResponseOutput) Time() pulumi.StringOutput {
	return o.ApplyT(func(v SparkConfigPropertiesResponse) string { return v.Time }).(pulumi.StringOutput)
}

type SparkConfigPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (SparkConfigPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SparkConfigPropertiesResponse)(nil)).Elem()
}

func (o SparkConfigPropertiesResponsePtrOutput) ToSparkConfigPropertiesResponsePtrOutput() SparkConfigPropertiesResponsePtrOutput {
	return o
}

func (o SparkConfigPropertiesResponsePtrOutput) ToSparkConfigPropertiesResponsePtrOutputWithContext(ctx context.Context) SparkConfigPropertiesResponsePtrOutput {
	return o
}

func (o SparkConfigPropertiesResponsePtrOutput) Elem() SparkConfigPropertiesResponseOutput {
	return o.ApplyT(func(v *SparkConfigPropertiesResponse) SparkConfigPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret SparkConfigPropertiesResponse
		return ret
	}).(SparkConfigPropertiesResponseOutput)
}

func (o SparkConfigPropertiesResponsePtrOutput) ConfigurationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SparkConfigPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ConfigurationType
	}).(pulumi.StringPtrOutput)
}

func (o SparkConfigPropertiesResponsePtrOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SparkConfigPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Content
	}).(pulumi.StringPtrOutput)
}

func (o SparkConfigPropertiesResponsePtrOutput) Filename() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SparkConfigPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Filename
	}).(pulumi.StringPtrOutput)
}

func (o SparkConfigPropertiesResponsePtrOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SparkConfigPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Time
	}).(pulumi.StringPtrOutput)
}

type SqlPoolVulnerabilityAssessmentRuleBaselineItem struct {
	Result []string `pulumi:"result"`
}





type SqlPoolVulnerabilityAssessmentRuleBaselineItemInput interface {
	pulumi.Input

	ToSqlPoolVulnerabilityAssessmentRuleBaselineItemOutput() SqlPoolVulnerabilityAssessmentRuleBaselineItemOutput
	ToSqlPoolVulnerabilityAssessmentRuleBaselineItemOutputWithContext(context.Context) SqlPoolVulnerabilityAssessmentRuleBaselineItemOutput
}

type SqlPoolVulnerabilityAssessmentRuleBaselineItemArgs struct {
	Result pulumi.StringArrayInput `pulumi:"result"`
}

func (SqlPoolVulnerabilityAssessmentRuleBaselineItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlPoolVulnerabilityAssessmentRuleBaselineItem)(nil)).Elem()
}

func (i SqlPoolVulnerabilityAssessmentRuleBaselineItemArgs) ToSqlPoolVulnerabilityAssessmentRuleBaselineItemOutput() SqlPoolVulnerabilityAssessmentRuleBaselineItemOutput {
	return i.ToSqlPoolVulnerabilityAssessmentRuleBaselineItemOutputWithContext(context.Background())
}

func (i SqlPoolVulnerabilityAssessmentRuleBaselineItemArgs) ToSqlPoolVulnerabilityAssessmentRuleBaselineItemOutputWithContext(ctx context.Context) SqlPoolVulnerabilityAssessmentRuleBaselineItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlPoolVulnerabilityAssessmentRuleBaselineItemOutput)
}





type SqlPoolVulnerabilityAssessmentRuleBaselineItemArrayInput interface {
	pulumi.Input

	ToSqlPoolVulnerabilityAssessmentRuleBaselineItemArrayOutput() SqlPoolVulnerabilityAssessmentRuleBaselineItemArrayOutput
	ToSqlPoolVulnerabilityAssessmentRuleBaselineItemArrayOutputWithContext(context.Context) SqlPoolVulnerabilityAssessmentRuleBaselineItemArrayOutput
}

type SqlPoolVulnerabilityAssessmentRuleBaselineItemArray []SqlPoolVulnerabilityAssessmentRuleBaselineItemInput

func (SqlPoolVulnerabilityAssessmentRuleBaselineItemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SqlPoolVulnerabilityAssessmentRuleBaselineItem)(nil)).Elem()
}

func (i SqlPoolVulnerabilityAssessmentRuleBaselineItemArray) ToSqlPoolVulnerabilityAssessmentRuleBaselineItemArrayOutput() SqlPoolVulnerabilityAssessmentRuleBaselineItemArrayOutput {
	return i.ToSqlPoolVulnerabilityAssessmentRuleBaselineItemArrayOutputWithContext(context.Background())
}

func (i SqlPoolVulnerabilityAssessmentRuleBaselineItemArray) ToSqlPoolVulnerabilityAssessmentRuleBaselineItemArrayOutputWithContext(ctx context.Context) SqlPoolVulnerabilityAssessmentRuleBaselineItemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlPoolVulnerabilityAssessmentRuleBaselineItemArrayOutput)
}

type SqlPoolVulnerabilityAssessmentRuleBaselineItemOutput struct{ *pulumi.OutputState }

func (SqlPoolVulnerabilityAssessmentRuleBaselineItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlPoolVulnerabilityAssessmentRuleBaselineItem)(nil)).Elem()
}

func (o SqlPoolVulnerabilityAssessmentRuleBaselineItemOutput) ToSqlPoolVulnerabilityAssessmentRuleBaselineItemOutput() SqlPoolVulnerabilityAssessmentRuleBaselineItemOutput {
	return o
}

func (o SqlPoolVulnerabilityAssessmentRuleBaselineItemOutput) ToSqlPoolVulnerabilityAssessmentRuleBaselineItemOutputWithContext(ctx context.Context) SqlPoolVulnerabilityAssessmentRuleBaselineItemOutput {
	return o
}

func (o SqlPoolVulnerabilityAssessmentRuleBaselineItemOutput) Result() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SqlPoolVulnerabilityAssessmentRuleBaselineItem) []string { return v.Result }).(pulumi.StringArrayOutput)
}

type SqlPoolVulnerabilityAssessmentRuleBaselineItemArrayOutput struct{ *pulumi.OutputState }

func (SqlPoolVulnerabilityAssessmentRuleBaselineItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SqlPoolVulnerabilityAssessmentRuleBaselineItem)(nil)).Elem()
}

func (o SqlPoolVulnerabilityAssessmentRuleBaselineItemArrayOutput) ToSqlPoolVulnerabilityAssessmentRuleBaselineItemArrayOutput() SqlPoolVulnerabilityAssessmentRuleBaselineItemArrayOutput {
	return o
}

func (o SqlPoolVulnerabilityAssessmentRuleBaselineItemArrayOutput) ToSqlPoolVulnerabilityAssessmentRuleBaselineItemArrayOutputWithContext(ctx context.Context) SqlPoolVulnerabilityAssessmentRuleBaselineItemArrayOutput {
	return o
}

func (o SqlPoolVulnerabilityAssessmentRuleBaselineItemArrayOutput) Index(i pulumi.IntInput) SqlPoolVulnerabilityAssessmentRuleBaselineItemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SqlPoolVulnerabilityAssessmentRuleBaselineItem {
		return vs[0].([]SqlPoolVulnerabilityAssessmentRuleBaselineItem)[vs[1].(int)]
	}).(SqlPoolVulnerabilityAssessmentRuleBaselineItemOutput)
}

type SqlPoolVulnerabilityAssessmentRuleBaselineItemResponse struct {
	Result []string `pulumi:"result"`
}

type SqlPoolVulnerabilityAssessmentRuleBaselineItemResponseOutput struct{ *pulumi.OutputState }

func (SqlPoolVulnerabilityAssessmentRuleBaselineItemResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlPoolVulnerabilityAssessmentRuleBaselineItemResponse)(nil)).Elem()
}

func (o SqlPoolVulnerabilityAssessmentRuleBaselineItemResponseOutput) ToSqlPoolVulnerabilityAssessmentRuleBaselineItemResponseOutput() SqlPoolVulnerabilityAssessmentRuleBaselineItemResponseOutput {
	return o
}

func (o SqlPoolVulnerabilityAssessmentRuleBaselineItemResponseOutput) ToSqlPoolVulnerabilityAssessmentRuleBaselineItemResponseOutputWithContext(ctx context.Context) SqlPoolVulnerabilityAssessmentRuleBaselineItemResponseOutput {
	return o
}

func (o SqlPoolVulnerabilityAssessmentRuleBaselineItemResponseOutput) Result() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SqlPoolVulnerabilityAssessmentRuleBaselineItemResponse) []string { return v.Result }).(pulumi.StringArrayOutput)
}

type SqlPoolVulnerabilityAssessmentRuleBaselineItemResponseArrayOutput struct{ *pulumi.OutputState }

func (SqlPoolVulnerabilityAssessmentRuleBaselineItemResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SqlPoolVulnerabilityAssessmentRuleBaselineItemResponse)(nil)).Elem()
}

func (o SqlPoolVulnerabilityAssessmentRuleBaselineItemResponseArrayOutput) ToSqlPoolVulnerabilityAssessmentRuleBaselineItemResponseArrayOutput() SqlPoolVulnerabilityAssessmentRuleBaselineItemResponseArrayOutput {
	return o
}

func (o SqlPoolVulnerabilityAssessmentRuleBaselineItemResponseArrayOutput) ToSqlPoolVulnerabilityAssessmentRuleBaselineItemResponseArrayOutputWithContext(ctx context.Context) SqlPoolVulnerabilityAssessmentRuleBaselineItemResponseArrayOutput {
	return o
}

func (o SqlPoolVulnerabilityAssessmentRuleBaselineItemResponseArrayOutput) Index(i pulumi.IntInput) SqlPoolVulnerabilityAssessmentRuleBaselineItemResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SqlPoolVulnerabilityAssessmentRuleBaselineItemResponse {
		return vs[0].([]SqlPoolVulnerabilityAssessmentRuleBaselineItemResponse)[vs[1].(int)]
	}).(SqlPoolVulnerabilityAssessmentRuleBaselineItemResponseOutput)
}

type SsisEnvironmentReferenceResponse struct {
	EnvironmentFolderName *string  `pulumi:"environmentFolderName"`
	EnvironmentName       *string  `pulumi:"environmentName"`
	Id                    *float64 `pulumi:"id"`
	ReferenceType         *string  `pulumi:"referenceType"`
}

type SsisEnvironmentResponse struct {
	Description *string                `pulumi:"description"`
	FolderId    *float64               `pulumi:"folderId"`
	Id          *float64               `pulumi:"id"`
	Name        *string                `pulumi:"name"`
	Type        string                 `pulumi:"type"`
	Variables   []SsisVariableResponse `pulumi:"variables"`
}

type SsisFolderResponse struct {
	Description *string  `pulumi:"description"`
	Id          *float64 `pulumi:"id"`
	Name        *string  `pulumi:"name"`
	Type        string   `pulumi:"type"`
}

type SsisPackageResponse struct {
	Description    *string                 `pulumi:"description"`
	FolderId       *float64                `pulumi:"folderId"`
	Id             *float64                `pulumi:"id"`
	Name           *string                 `pulumi:"name"`
	Parameters     []SsisParameterResponse `pulumi:"parameters"`
	ProjectId      *float64                `pulumi:"projectId"`
	ProjectVersion *float64                `pulumi:"projectVersion"`
	Type           string                  `pulumi:"type"`
}

type SsisParameterResponse struct {
	DataType              *string  `pulumi:"dataType"`
	DefaultValue          *string  `pulumi:"defaultValue"`
	Description           *string  `pulumi:"description"`
	DesignDefaultValue    *string  `pulumi:"designDefaultValue"`
	Id                    *float64 `pulumi:"id"`
	Name                  *string  `pulumi:"name"`
	Required              *bool    `pulumi:"required"`
	Sensitive             *bool    `pulumi:"sensitive"`
	SensitiveDefaultValue *string  `pulumi:"sensitiveDefaultValue"`
	ValueSet              *bool    `pulumi:"valueSet"`
	ValueType             *string  `pulumi:"valueType"`
	Variable              *string  `pulumi:"variable"`
}

type SsisProjectResponse struct {
	Description     *string                            `pulumi:"description"`
	EnvironmentRefs []SsisEnvironmentReferenceResponse `pulumi:"environmentRefs"`
	FolderId        *float64                           `pulumi:"folderId"`
	Id              *float64                           `pulumi:"id"`
	Name            *string                            `pulumi:"name"`
	Parameters      []SsisParameterResponse            `pulumi:"parameters"`
	Type            string                             `pulumi:"type"`
	Version         *float64                           `pulumi:"version"`
}

type SsisVariableResponse struct {
	DataType       *string  `pulumi:"dataType"`
	Description    *string  `pulumi:"description"`
	Id             *float64 `pulumi:"id"`
	Name           *string  `pulumi:"name"`
	Sensitive      *bool    `pulumi:"sensitive"`
	SensitiveValue *string  `pulumi:"sensitiveValue"`
	Value          *string  `pulumi:"value"`
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

type TableLevelSharingProperties struct {
	ExternalTablesToExclude    []string `pulumi:"externalTablesToExclude"`
	ExternalTablesToInclude    []string `pulumi:"externalTablesToInclude"`
	MaterializedViewsToExclude []string `pulumi:"materializedViewsToExclude"`
	MaterializedViewsToInclude []string `pulumi:"materializedViewsToInclude"`
	TablesToExclude            []string `pulumi:"tablesToExclude"`
	TablesToInclude            []string `pulumi:"tablesToInclude"`
}





type TableLevelSharingPropertiesInput interface {
	pulumi.Input

	ToTableLevelSharingPropertiesOutput() TableLevelSharingPropertiesOutput
	ToTableLevelSharingPropertiesOutputWithContext(context.Context) TableLevelSharingPropertiesOutput
}

type TableLevelSharingPropertiesArgs struct {
	ExternalTablesToExclude    pulumi.StringArrayInput `pulumi:"externalTablesToExclude"`
	ExternalTablesToInclude    pulumi.StringArrayInput `pulumi:"externalTablesToInclude"`
	MaterializedViewsToExclude pulumi.StringArrayInput `pulumi:"materializedViewsToExclude"`
	MaterializedViewsToInclude pulumi.StringArrayInput `pulumi:"materializedViewsToInclude"`
	TablesToExclude            pulumi.StringArrayInput `pulumi:"tablesToExclude"`
	TablesToInclude            pulumi.StringArrayInput `pulumi:"tablesToInclude"`
}

func (TableLevelSharingPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TableLevelSharingProperties)(nil)).Elem()
}

func (i TableLevelSharingPropertiesArgs) ToTableLevelSharingPropertiesOutput() TableLevelSharingPropertiesOutput {
	return i.ToTableLevelSharingPropertiesOutputWithContext(context.Background())
}

func (i TableLevelSharingPropertiesArgs) ToTableLevelSharingPropertiesOutputWithContext(ctx context.Context) TableLevelSharingPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableLevelSharingPropertiesOutput)
}

func (i TableLevelSharingPropertiesArgs) ToTableLevelSharingPropertiesPtrOutput() TableLevelSharingPropertiesPtrOutput {
	return i.ToTableLevelSharingPropertiesPtrOutputWithContext(context.Background())
}

func (i TableLevelSharingPropertiesArgs) ToTableLevelSharingPropertiesPtrOutputWithContext(ctx context.Context) TableLevelSharingPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableLevelSharingPropertiesOutput).ToTableLevelSharingPropertiesPtrOutputWithContext(ctx)
}









type TableLevelSharingPropertiesPtrInput interface {
	pulumi.Input

	ToTableLevelSharingPropertiesPtrOutput() TableLevelSharingPropertiesPtrOutput
	ToTableLevelSharingPropertiesPtrOutputWithContext(context.Context) TableLevelSharingPropertiesPtrOutput
}

type tableLevelSharingPropertiesPtrType TableLevelSharingPropertiesArgs

func TableLevelSharingPropertiesPtr(v *TableLevelSharingPropertiesArgs) TableLevelSharingPropertiesPtrInput {
	return (*tableLevelSharingPropertiesPtrType)(v)
}

func (*tableLevelSharingPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TableLevelSharingProperties)(nil)).Elem()
}

func (i *tableLevelSharingPropertiesPtrType) ToTableLevelSharingPropertiesPtrOutput() TableLevelSharingPropertiesPtrOutput {
	return i.ToTableLevelSharingPropertiesPtrOutputWithContext(context.Background())
}

func (i *tableLevelSharingPropertiesPtrType) ToTableLevelSharingPropertiesPtrOutputWithContext(ctx context.Context) TableLevelSharingPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableLevelSharingPropertiesPtrOutput)
}

type TableLevelSharingPropertiesOutput struct{ *pulumi.OutputState }

func (TableLevelSharingPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TableLevelSharingProperties)(nil)).Elem()
}

func (o TableLevelSharingPropertiesOutput) ToTableLevelSharingPropertiesOutput() TableLevelSharingPropertiesOutput {
	return o
}

func (o TableLevelSharingPropertiesOutput) ToTableLevelSharingPropertiesOutputWithContext(ctx context.Context) TableLevelSharingPropertiesOutput {
	return o
}

func (o TableLevelSharingPropertiesOutput) ToTableLevelSharingPropertiesPtrOutput() TableLevelSharingPropertiesPtrOutput {
	return o.ToTableLevelSharingPropertiesPtrOutputWithContext(context.Background())
}

func (o TableLevelSharingPropertiesOutput) ToTableLevelSharingPropertiesPtrOutputWithContext(ctx context.Context) TableLevelSharingPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TableLevelSharingProperties) *TableLevelSharingProperties {
		return &v
	}).(TableLevelSharingPropertiesPtrOutput)
}

func (o TableLevelSharingPropertiesOutput) ExternalTablesToExclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TableLevelSharingProperties) []string { return v.ExternalTablesToExclude }).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesOutput) ExternalTablesToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TableLevelSharingProperties) []string { return v.ExternalTablesToInclude }).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesOutput) MaterializedViewsToExclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TableLevelSharingProperties) []string { return v.MaterializedViewsToExclude }).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesOutput) MaterializedViewsToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TableLevelSharingProperties) []string { return v.MaterializedViewsToInclude }).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesOutput) TablesToExclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TableLevelSharingProperties) []string { return v.TablesToExclude }).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesOutput) TablesToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TableLevelSharingProperties) []string { return v.TablesToInclude }).(pulumi.StringArrayOutput)
}

type TableLevelSharingPropertiesPtrOutput struct{ *pulumi.OutputState }

func (TableLevelSharingPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TableLevelSharingProperties)(nil)).Elem()
}

func (o TableLevelSharingPropertiesPtrOutput) ToTableLevelSharingPropertiesPtrOutput() TableLevelSharingPropertiesPtrOutput {
	return o
}

func (o TableLevelSharingPropertiesPtrOutput) ToTableLevelSharingPropertiesPtrOutputWithContext(ctx context.Context) TableLevelSharingPropertiesPtrOutput {
	return o
}

func (o TableLevelSharingPropertiesPtrOutput) Elem() TableLevelSharingPropertiesOutput {
	return o.ApplyT(func(v *TableLevelSharingProperties) TableLevelSharingProperties {
		if v != nil {
			return *v
		}
		var ret TableLevelSharingProperties
		return ret
	}).(TableLevelSharingPropertiesOutput)
}

func (o TableLevelSharingPropertiesPtrOutput) ExternalTablesToExclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TableLevelSharingProperties) []string {
		if v == nil {
			return nil
		}
		return v.ExternalTablesToExclude
	}).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesPtrOutput) ExternalTablesToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TableLevelSharingProperties) []string {
		if v == nil {
			return nil
		}
		return v.ExternalTablesToInclude
	}).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesPtrOutput) MaterializedViewsToExclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TableLevelSharingProperties) []string {
		if v == nil {
			return nil
		}
		return v.MaterializedViewsToExclude
	}).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesPtrOutput) MaterializedViewsToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TableLevelSharingProperties) []string {
		if v == nil {
			return nil
		}
		return v.MaterializedViewsToInclude
	}).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesPtrOutput) TablesToExclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TableLevelSharingProperties) []string {
		if v == nil {
			return nil
		}
		return v.TablesToExclude
	}).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesPtrOutput) TablesToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TableLevelSharingProperties) []string {
		if v == nil {
			return nil
		}
		return v.TablesToInclude
	}).(pulumi.StringArrayOutput)
}

type TableLevelSharingPropertiesResponse struct {
	ExternalTablesToExclude    []string `pulumi:"externalTablesToExclude"`
	ExternalTablesToInclude    []string `pulumi:"externalTablesToInclude"`
	MaterializedViewsToExclude []string `pulumi:"materializedViewsToExclude"`
	MaterializedViewsToInclude []string `pulumi:"materializedViewsToInclude"`
	TablesToExclude            []string `pulumi:"tablesToExclude"`
	TablesToInclude            []string `pulumi:"tablesToInclude"`
}

type TableLevelSharingPropertiesResponseOutput struct{ *pulumi.OutputState }

func (TableLevelSharingPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TableLevelSharingPropertiesResponse)(nil)).Elem()
}

func (o TableLevelSharingPropertiesResponseOutput) ToTableLevelSharingPropertiesResponseOutput() TableLevelSharingPropertiesResponseOutput {
	return o
}

func (o TableLevelSharingPropertiesResponseOutput) ToTableLevelSharingPropertiesResponseOutputWithContext(ctx context.Context) TableLevelSharingPropertiesResponseOutput {
	return o
}

func (o TableLevelSharingPropertiesResponseOutput) ExternalTablesToExclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TableLevelSharingPropertiesResponse) []string { return v.ExternalTablesToExclude }).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesResponseOutput) ExternalTablesToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TableLevelSharingPropertiesResponse) []string { return v.ExternalTablesToInclude }).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesResponseOutput) MaterializedViewsToExclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TableLevelSharingPropertiesResponse) []string { return v.MaterializedViewsToExclude }).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesResponseOutput) MaterializedViewsToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TableLevelSharingPropertiesResponse) []string { return v.MaterializedViewsToInclude }).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesResponseOutput) TablesToExclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TableLevelSharingPropertiesResponse) []string { return v.TablesToExclude }).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesResponseOutput) TablesToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TableLevelSharingPropertiesResponse) []string { return v.TablesToInclude }).(pulumi.StringArrayOutput)
}

type TableLevelSharingPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (TableLevelSharingPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TableLevelSharingPropertiesResponse)(nil)).Elem()
}

func (o TableLevelSharingPropertiesResponsePtrOutput) ToTableLevelSharingPropertiesResponsePtrOutput() TableLevelSharingPropertiesResponsePtrOutput {
	return o
}

func (o TableLevelSharingPropertiesResponsePtrOutput) ToTableLevelSharingPropertiesResponsePtrOutputWithContext(ctx context.Context) TableLevelSharingPropertiesResponsePtrOutput {
	return o
}

func (o TableLevelSharingPropertiesResponsePtrOutput) Elem() TableLevelSharingPropertiesResponseOutput {
	return o.ApplyT(func(v *TableLevelSharingPropertiesResponse) TableLevelSharingPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret TableLevelSharingPropertiesResponse
		return ret
	}).(TableLevelSharingPropertiesResponseOutput)
}

func (o TableLevelSharingPropertiesResponsePtrOutput) ExternalTablesToExclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TableLevelSharingPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.ExternalTablesToExclude
	}).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesResponsePtrOutput) ExternalTablesToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TableLevelSharingPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.ExternalTablesToInclude
	}).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesResponsePtrOutput) MaterializedViewsToExclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TableLevelSharingPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.MaterializedViewsToExclude
	}).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesResponsePtrOutput) MaterializedViewsToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TableLevelSharingPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.MaterializedViewsToInclude
	}).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesResponsePtrOutput) TablesToExclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TableLevelSharingPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.TablesToExclude
	}).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesResponsePtrOutput) TablesToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TableLevelSharingPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.TablesToInclude
	}).(pulumi.StringArrayOutput)
}

type UserAssignedManagedIdentityResponse struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}

type UserAssignedManagedIdentityResponseOutput struct{ *pulumi.OutputState }

func (UserAssignedManagedIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedManagedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedManagedIdentityResponseOutput) ToUserAssignedManagedIdentityResponseOutput() UserAssignedManagedIdentityResponseOutput {
	return o
}

func (o UserAssignedManagedIdentityResponseOutput) ToUserAssignedManagedIdentityResponseOutputWithContext(ctx context.Context) UserAssignedManagedIdentityResponseOutput {
	return o
}

func (o UserAssignedManagedIdentityResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedManagedIdentityResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o UserAssignedManagedIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedManagedIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type UserAssignedManagedIdentityResponseMapOutput struct{ *pulumi.OutputState }

func (UserAssignedManagedIdentityResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedManagedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedManagedIdentityResponseMapOutput) ToUserAssignedManagedIdentityResponseMapOutput() UserAssignedManagedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedManagedIdentityResponseMapOutput) ToUserAssignedManagedIdentityResponseMapOutputWithContext(ctx context.Context) UserAssignedManagedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedManagedIdentityResponseMapOutput) MapIndex(k pulumi.StringInput) UserAssignedManagedIdentityResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserAssignedManagedIdentityResponse {
		return vs[0].(map[string]UserAssignedManagedIdentityResponse)[vs[1].(string)]
	}).(UserAssignedManagedIdentityResponseOutput)
}

type VirtualNetworkProfile struct {
	ComputeSubnetId *string `pulumi:"computeSubnetId"`
}





type VirtualNetworkProfileInput interface {
	pulumi.Input

	ToVirtualNetworkProfileOutput() VirtualNetworkProfileOutput
	ToVirtualNetworkProfileOutputWithContext(context.Context) VirtualNetworkProfileOutput
}

type VirtualNetworkProfileArgs struct {
	ComputeSubnetId pulumi.StringPtrInput `pulumi:"computeSubnetId"`
}

func (VirtualNetworkProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkProfile)(nil)).Elem()
}

func (i VirtualNetworkProfileArgs) ToVirtualNetworkProfileOutput() VirtualNetworkProfileOutput {
	return i.ToVirtualNetworkProfileOutputWithContext(context.Background())
}

func (i VirtualNetworkProfileArgs) ToVirtualNetworkProfileOutputWithContext(ctx context.Context) VirtualNetworkProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkProfileOutput)
}

func (i VirtualNetworkProfileArgs) ToVirtualNetworkProfilePtrOutput() VirtualNetworkProfilePtrOutput {
	return i.ToVirtualNetworkProfilePtrOutputWithContext(context.Background())
}

func (i VirtualNetworkProfileArgs) ToVirtualNetworkProfilePtrOutputWithContext(ctx context.Context) VirtualNetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkProfileOutput).ToVirtualNetworkProfilePtrOutputWithContext(ctx)
}









type VirtualNetworkProfilePtrInput interface {
	pulumi.Input

	ToVirtualNetworkProfilePtrOutput() VirtualNetworkProfilePtrOutput
	ToVirtualNetworkProfilePtrOutputWithContext(context.Context) VirtualNetworkProfilePtrOutput
}

type virtualNetworkProfilePtrType VirtualNetworkProfileArgs

func VirtualNetworkProfilePtr(v *VirtualNetworkProfileArgs) VirtualNetworkProfilePtrInput {
	return (*virtualNetworkProfilePtrType)(v)
}

func (*virtualNetworkProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkProfile)(nil)).Elem()
}

func (i *virtualNetworkProfilePtrType) ToVirtualNetworkProfilePtrOutput() VirtualNetworkProfilePtrOutput {
	return i.ToVirtualNetworkProfilePtrOutputWithContext(context.Background())
}

func (i *virtualNetworkProfilePtrType) ToVirtualNetworkProfilePtrOutputWithContext(ctx context.Context) VirtualNetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkProfilePtrOutput)
}

type VirtualNetworkProfileOutput struct{ *pulumi.OutputState }

func (VirtualNetworkProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkProfile)(nil)).Elem()
}

func (o VirtualNetworkProfileOutput) ToVirtualNetworkProfileOutput() VirtualNetworkProfileOutput {
	return o
}

func (o VirtualNetworkProfileOutput) ToVirtualNetworkProfileOutputWithContext(ctx context.Context) VirtualNetworkProfileOutput {
	return o
}

func (o VirtualNetworkProfileOutput) ToVirtualNetworkProfilePtrOutput() VirtualNetworkProfilePtrOutput {
	return o.ToVirtualNetworkProfilePtrOutputWithContext(context.Background())
}

func (o VirtualNetworkProfileOutput) ToVirtualNetworkProfilePtrOutputWithContext(ctx context.Context) VirtualNetworkProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworkProfile) *VirtualNetworkProfile {
		return &v
	}).(VirtualNetworkProfilePtrOutput)
}

func (o VirtualNetworkProfileOutput) ComputeSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkProfile) *string { return v.ComputeSubnetId }).(pulumi.StringPtrOutput)
}

type VirtualNetworkProfilePtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkProfile)(nil)).Elem()
}

func (o VirtualNetworkProfilePtrOutput) ToVirtualNetworkProfilePtrOutput() VirtualNetworkProfilePtrOutput {
	return o
}

func (o VirtualNetworkProfilePtrOutput) ToVirtualNetworkProfilePtrOutputWithContext(ctx context.Context) VirtualNetworkProfilePtrOutput {
	return o
}

func (o VirtualNetworkProfilePtrOutput) Elem() VirtualNetworkProfileOutput {
	return o.ApplyT(func(v *VirtualNetworkProfile) VirtualNetworkProfile {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkProfile
		return ret
	}).(VirtualNetworkProfileOutput)
}

func (o VirtualNetworkProfilePtrOutput) ComputeSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.ComputeSubnetId
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworkProfileResponse struct {
	ComputeSubnetId *string `pulumi:"computeSubnetId"`
}

type VirtualNetworkProfileResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkProfileResponse)(nil)).Elem()
}

func (o VirtualNetworkProfileResponseOutput) ToVirtualNetworkProfileResponseOutput() VirtualNetworkProfileResponseOutput {
	return o
}

func (o VirtualNetworkProfileResponseOutput) ToVirtualNetworkProfileResponseOutputWithContext(ctx context.Context) VirtualNetworkProfileResponseOutput {
	return o
}

func (o VirtualNetworkProfileResponseOutput) ComputeSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkProfileResponse) *string { return v.ComputeSubnetId }).(pulumi.StringPtrOutput)
}

type VirtualNetworkProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkProfileResponse)(nil)).Elem()
}

func (o VirtualNetworkProfileResponsePtrOutput) ToVirtualNetworkProfileResponsePtrOutput() VirtualNetworkProfileResponsePtrOutput {
	return o
}

func (o VirtualNetworkProfileResponsePtrOutput) ToVirtualNetworkProfileResponsePtrOutputWithContext(ctx context.Context) VirtualNetworkProfileResponsePtrOutput {
	return o
}

func (o VirtualNetworkProfileResponsePtrOutput) Elem() VirtualNetworkProfileResponseOutput {
	return o.ApplyT(func(v *VirtualNetworkProfileResponse) VirtualNetworkProfileResponse {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkProfileResponse
		return ret
	}).(VirtualNetworkProfileResponseOutput)
}

func (o VirtualNetworkProfileResponsePtrOutput) ComputeSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.ComputeSubnetId
	}).(pulumi.StringPtrOutput)
}

type VulnerabilityAssessmentRecurringScansProperties struct {
	EmailSubscriptionAdmins *bool    `pulumi:"emailSubscriptionAdmins"`
	Emails                  []string `pulumi:"emails"`
	IsEnabled               *bool    `pulumi:"isEnabled"`
}


func (val *VulnerabilityAssessmentRecurringScansProperties) Defaults() *VulnerabilityAssessmentRecurringScansProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EmailSubscriptionAdmins) {
		emailSubscriptionAdmins_ := true
		tmp.EmailSubscriptionAdmins = &emailSubscriptionAdmins_
	}
	return &tmp
}





type VulnerabilityAssessmentRecurringScansPropertiesInput interface {
	pulumi.Input

	ToVulnerabilityAssessmentRecurringScansPropertiesOutput() VulnerabilityAssessmentRecurringScansPropertiesOutput
	ToVulnerabilityAssessmentRecurringScansPropertiesOutputWithContext(context.Context) VulnerabilityAssessmentRecurringScansPropertiesOutput
}

type VulnerabilityAssessmentRecurringScansPropertiesArgs struct {
	EmailSubscriptionAdmins pulumi.BoolPtrInput     `pulumi:"emailSubscriptionAdmins"`
	Emails                  pulumi.StringArrayInput `pulumi:"emails"`
	IsEnabled               pulumi.BoolPtrInput     `pulumi:"isEnabled"`
}

func (VulnerabilityAssessmentRecurringScansPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VulnerabilityAssessmentRecurringScansProperties)(nil)).Elem()
}

func (i VulnerabilityAssessmentRecurringScansPropertiesArgs) ToVulnerabilityAssessmentRecurringScansPropertiesOutput() VulnerabilityAssessmentRecurringScansPropertiesOutput {
	return i.ToVulnerabilityAssessmentRecurringScansPropertiesOutputWithContext(context.Background())
}

func (i VulnerabilityAssessmentRecurringScansPropertiesArgs) ToVulnerabilityAssessmentRecurringScansPropertiesOutputWithContext(ctx context.Context) VulnerabilityAssessmentRecurringScansPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VulnerabilityAssessmentRecurringScansPropertiesOutput)
}

func (i VulnerabilityAssessmentRecurringScansPropertiesArgs) ToVulnerabilityAssessmentRecurringScansPropertiesPtrOutput() VulnerabilityAssessmentRecurringScansPropertiesPtrOutput {
	return i.ToVulnerabilityAssessmentRecurringScansPropertiesPtrOutputWithContext(context.Background())
}

func (i VulnerabilityAssessmentRecurringScansPropertiesArgs) ToVulnerabilityAssessmentRecurringScansPropertiesPtrOutputWithContext(ctx context.Context) VulnerabilityAssessmentRecurringScansPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VulnerabilityAssessmentRecurringScansPropertiesOutput).ToVulnerabilityAssessmentRecurringScansPropertiesPtrOutputWithContext(ctx)
}









type VulnerabilityAssessmentRecurringScansPropertiesPtrInput interface {
	pulumi.Input

	ToVulnerabilityAssessmentRecurringScansPropertiesPtrOutput() VulnerabilityAssessmentRecurringScansPropertiesPtrOutput
	ToVulnerabilityAssessmentRecurringScansPropertiesPtrOutputWithContext(context.Context) VulnerabilityAssessmentRecurringScansPropertiesPtrOutput
}

type vulnerabilityAssessmentRecurringScansPropertiesPtrType VulnerabilityAssessmentRecurringScansPropertiesArgs

func VulnerabilityAssessmentRecurringScansPropertiesPtr(v *VulnerabilityAssessmentRecurringScansPropertiesArgs) VulnerabilityAssessmentRecurringScansPropertiesPtrInput {
	return (*vulnerabilityAssessmentRecurringScansPropertiesPtrType)(v)
}

func (*vulnerabilityAssessmentRecurringScansPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VulnerabilityAssessmentRecurringScansProperties)(nil)).Elem()
}

func (i *vulnerabilityAssessmentRecurringScansPropertiesPtrType) ToVulnerabilityAssessmentRecurringScansPropertiesPtrOutput() VulnerabilityAssessmentRecurringScansPropertiesPtrOutput {
	return i.ToVulnerabilityAssessmentRecurringScansPropertiesPtrOutputWithContext(context.Background())
}

func (i *vulnerabilityAssessmentRecurringScansPropertiesPtrType) ToVulnerabilityAssessmentRecurringScansPropertiesPtrOutputWithContext(ctx context.Context) VulnerabilityAssessmentRecurringScansPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VulnerabilityAssessmentRecurringScansPropertiesPtrOutput)
}

type VulnerabilityAssessmentRecurringScansPropertiesOutput struct{ *pulumi.OutputState }

func (VulnerabilityAssessmentRecurringScansPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VulnerabilityAssessmentRecurringScansProperties)(nil)).Elem()
}

func (o VulnerabilityAssessmentRecurringScansPropertiesOutput) ToVulnerabilityAssessmentRecurringScansPropertiesOutput() VulnerabilityAssessmentRecurringScansPropertiesOutput {
	return o
}

func (o VulnerabilityAssessmentRecurringScansPropertiesOutput) ToVulnerabilityAssessmentRecurringScansPropertiesOutputWithContext(ctx context.Context) VulnerabilityAssessmentRecurringScansPropertiesOutput {
	return o
}

func (o VulnerabilityAssessmentRecurringScansPropertiesOutput) ToVulnerabilityAssessmentRecurringScansPropertiesPtrOutput() VulnerabilityAssessmentRecurringScansPropertiesPtrOutput {
	return o.ToVulnerabilityAssessmentRecurringScansPropertiesPtrOutputWithContext(context.Background())
}

func (o VulnerabilityAssessmentRecurringScansPropertiesOutput) ToVulnerabilityAssessmentRecurringScansPropertiesPtrOutputWithContext(ctx context.Context) VulnerabilityAssessmentRecurringScansPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VulnerabilityAssessmentRecurringScansProperties) *VulnerabilityAssessmentRecurringScansProperties {
		return &v
	}).(VulnerabilityAssessmentRecurringScansPropertiesPtrOutput)
}

func (o VulnerabilityAssessmentRecurringScansPropertiesOutput) EmailSubscriptionAdmins() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VulnerabilityAssessmentRecurringScansProperties) *bool { return v.EmailSubscriptionAdmins }).(pulumi.BoolPtrOutput)
}

func (o VulnerabilityAssessmentRecurringScansPropertiesOutput) Emails() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VulnerabilityAssessmentRecurringScansProperties) []string { return v.Emails }).(pulumi.StringArrayOutput)
}

func (o VulnerabilityAssessmentRecurringScansPropertiesOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VulnerabilityAssessmentRecurringScansProperties) *bool { return v.IsEnabled }).(pulumi.BoolPtrOutput)
}

type VulnerabilityAssessmentRecurringScansPropertiesPtrOutput struct{ *pulumi.OutputState }

func (VulnerabilityAssessmentRecurringScansPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VulnerabilityAssessmentRecurringScansProperties)(nil)).Elem()
}

func (o VulnerabilityAssessmentRecurringScansPropertiesPtrOutput) ToVulnerabilityAssessmentRecurringScansPropertiesPtrOutput() VulnerabilityAssessmentRecurringScansPropertiesPtrOutput {
	return o
}

func (o VulnerabilityAssessmentRecurringScansPropertiesPtrOutput) ToVulnerabilityAssessmentRecurringScansPropertiesPtrOutputWithContext(ctx context.Context) VulnerabilityAssessmentRecurringScansPropertiesPtrOutput {
	return o
}

func (o VulnerabilityAssessmentRecurringScansPropertiesPtrOutput) Elem() VulnerabilityAssessmentRecurringScansPropertiesOutput {
	return o.ApplyT(func(v *VulnerabilityAssessmentRecurringScansProperties) VulnerabilityAssessmentRecurringScansProperties {
		if v != nil {
			return *v
		}
		var ret VulnerabilityAssessmentRecurringScansProperties
		return ret
	}).(VulnerabilityAssessmentRecurringScansPropertiesOutput)
}

func (o VulnerabilityAssessmentRecurringScansPropertiesPtrOutput) EmailSubscriptionAdmins() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VulnerabilityAssessmentRecurringScansProperties) *bool {
		if v == nil {
			return nil
		}
		return v.EmailSubscriptionAdmins
	}).(pulumi.BoolPtrOutput)
}

func (o VulnerabilityAssessmentRecurringScansPropertiesPtrOutput) Emails() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VulnerabilityAssessmentRecurringScansProperties) []string {
		if v == nil {
			return nil
		}
		return v.Emails
	}).(pulumi.StringArrayOutput)
}

func (o VulnerabilityAssessmentRecurringScansPropertiesPtrOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VulnerabilityAssessmentRecurringScansProperties) *bool {
		if v == nil {
			return nil
		}
		return v.IsEnabled
	}).(pulumi.BoolPtrOutput)
}

type VulnerabilityAssessmentRecurringScansPropertiesResponse struct {
	EmailSubscriptionAdmins *bool    `pulumi:"emailSubscriptionAdmins"`
	Emails                  []string `pulumi:"emails"`
	IsEnabled               *bool    `pulumi:"isEnabled"`
}


func (val *VulnerabilityAssessmentRecurringScansPropertiesResponse) Defaults() *VulnerabilityAssessmentRecurringScansPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EmailSubscriptionAdmins) {
		emailSubscriptionAdmins_ := true
		tmp.EmailSubscriptionAdmins = &emailSubscriptionAdmins_
	}
	return &tmp
}

type VulnerabilityAssessmentRecurringScansPropertiesResponseOutput struct{ *pulumi.OutputState }

func (VulnerabilityAssessmentRecurringScansPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VulnerabilityAssessmentRecurringScansPropertiesResponse)(nil)).Elem()
}

func (o VulnerabilityAssessmentRecurringScansPropertiesResponseOutput) ToVulnerabilityAssessmentRecurringScansPropertiesResponseOutput() VulnerabilityAssessmentRecurringScansPropertiesResponseOutput {
	return o
}

func (o VulnerabilityAssessmentRecurringScansPropertiesResponseOutput) ToVulnerabilityAssessmentRecurringScansPropertiesResponseOutputWithContext(ctx context.Context) VulnerabilityAssessmentRecurringScansPropertiesResponseOutput {
	return o
}

func (o VulnerabilityAssessmentRecurringScansPropertiesResponseOutput) EmailSubscriptionAdmins() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VulnerabilityAssessmentRecurringScansPropertiesResponse) *bool {
		return v.EmailSubscriptionAdmins
	}).(pulumi.BoolPtrOutput)
}

func (o VulnerabilityAssessmentRecurringScansPropertiesResponseOutput) Emails() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VulnerabilityAssessmentRecurringScansPropertiesResponse) []string { return v.Emails }).(pulumi.StringArrayOutput)
}

func (o VulnerabilityAssessmentRecurringScansPropertiesResponseOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VulnerabilityAssessmentRecurringScansPropertiesResponse) *bool { return v.IsEnabled }).(pulumi.BoolPtrOutput)
}

type VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VulnerabilityAssessmentRecurringScansPropertiesResponse)(nil)).Elem()
}

func (o VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput) ToVulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput() VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput {
	return o
}

func (o VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput) ToVulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutputWithContext(ctx context.Context) VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput {
	return o
}

func (o VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput) Elem() VulnerabilityAssessmentRecurringScansPropertiesResponseOutput {
	return o.ApplyT(func(v *VulnerabilityAssessmentRecurringScansPropertiesResponse) VulnerabilityAssessmentRecurringScansPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret VulnerabilityAssessmentRecurringScansPropertiesResponse
		return ret
	}).(VulnerabilityAssessmentRecurringScansPropertiesResponseOutput)
}

func (o VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput) EmailSubscriptionAdmins() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VulnerabilityAssessmentRecurringScansPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EmailSubscriptionAdmins
	}).(pulumi.BoolPtrOutput)
}

func (o VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput) Emails() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VulnerabilityAssessmentRecurringScansPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.Emails
	}).(pulumi.StringArrayOutput)
}

func (o VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VulnerabilityAssessmentRecurringScansPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsEnabled
	}).(pulumi.BoolPtrOutput)
}

type WorkspaceKeyDetails struct {
	KeyVaultUrl *string `pulumi:"keyVaultUrl"`
	Name        *string `pulumi:"name"`
}





type WorkspaceKeyDetailsInput interface {
	pulumi.Input

	ToWorkspaceKeyDetailsOutput() WorkspaceKeyDetailsOutput
	ToWorkspaceKeyDetailsOutputWithContext(context.Context) WorkspaceKeyDetailsOutput
}

type WorkspaceKeyDetailsArgs struct {
	KeyVaultUrl pulumi.StringPtrInput `pulumi:"keyVaultUrl"`
	Name        pulumi.StringPtrInput `pulumi:"name"`
}

func (WorkspaceKeyDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceKeyDetails)(nil)).Elem()
}

func (i WorkspaceKeyDetailsArgs) ToWorkspaceKeyDetailsOutput() WorkspaceKeyDetailsOutput {
	return i.ToWorkspaceKeyDetailsOutputWithContext(context.Background())
}

func (i WorkspaceKeyDetailsArgs) ToWorkspaceKeyDetailsOutputWithContext(ctx context.Context) WorkspaceKeyDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceKeyDetailsOutput)
}

func (i WorkspaceKeyDetailsArgs) ToWorkspaceKeyDetailsPtrOutput() WorkspaceKeyDetailsPtrOutput {
	return i.ToWorkspaceKeyDetailsPtrOutputWithContext(context.Background())
}

func (i WorkspaceKeyDetailsArgs) ToWorkspaceKeyDetailsPtrOutputWithContext(ctx context.Context) WorkspaceKeyDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceKeyDetailsOutput).ToWorkspaceKeyDetailsPtrOutputWithContext(ctx)
}









type WorkspaceKeyDetailsPtrInput interface {
	pulumi.Input

	ToWorkspaceKeyDetailsPtrOutput() WorkspaceKeyDetailsPtrOutput
	ToWorkspaceKeyDetailsPtrOutputWithContext(context.Context) WorkspaceKeyDetailsPtrOutput
}

type workspaceKeyDetailsPtrType WorkspaceKeyDetailsArgs

func WorkspaceKeyDetailsPtr(v *WorkspaceKeyDetailsArgs) WorkspaceKeyDetailsPtrInput {
	return (*workspaceKeyDetailsPtrType)(v)
}

func (*workspaceKeyDetailsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceKeyDetails)(nil)).Elem()
}

func (i *workspaceKeyDetailsPtrType) ToWorkspaceKeyDetailsPtrOutput() WorkspaceKeyDetailsPtrOutput {
	return i.ToWorkspaceKeyDetailsPtrOutputWithContext(context.Background())
}

func (i *workspaceKeyDetailsPtrType) ToWorkspaceKeyDetailsPtrOutputWithContext(ctx context.Context) WorkspaceKeyDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceKeyDetailsPtrOutput)
}

type WorkspaceKeyDetailsOutput struct{ *pulumi.OutputState }

func (WorkspaceKeyDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceKeyDetails)(nil)).Elem()
}

func (o WorkspaceKeyDetailsOutput) ToWorkspaceKeyDetailsOutput() WorkspaceKeyDetailsOutput {
	return o
}

func (o WorkspaceKeyDetailsOutput) ToWorkspaceKeyDetailsOutputWithContext(ctx context.Context) WorkspaceKeyDetailsOutput {
	return o
}

func (o WorkspaceKeyDetailsOutput) ToWorkspaceKeyDetailsPtrOutput() WorkspaceKeyDetailsPtrOutput {
	return o.ToWorkspaceKeyDetailsPtrOutputWithContext(context.Background())
}

func (o WorkspaceKeyDetailsOutput) ToWorkspaceKeyDetailsPtrOutputWithContext(ctx context.Context) WorkspaceKeyDetailsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkspaceKeyDetails) *WorkspaceKeyDetails {
		return &v
	}).(WorkspaceKeyDetailsPtrOutput)
}

func (o WorkspaceKeyDetailsOutput) KeyVaultUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkspaceKeyDetails) *string { return v.KeyVaultUrl }).(pulumi.StringPtrOutput)
}

func (o WorkspaceKeyDetailsOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkspaceKeyDetails) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type WorkspaceKeyDetailsPtrOutput struct{ *pulumi.OutputState }

func (WorkspaceKeyDetailsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceKeyDetails)(nil)).Elem()
}

func (o WorkspaceKeyDetailsPtrOutput) ToWorkspaceKeyDetailsPtrOutput() WorkspaceKeyDetailsPtrOutput {
	return o
}

func (o WorkspaceKeyDetailsPtrOutput) ToWorkspaceKeyDetailsPtrOutputWithContext(ctx context.Context) WorkspaceKeyDetailsPtrOutput {
	return o
}

func (o WorkspaceKeyDetailsPtrOutput) Elem() WorkspaceKeyDetailsOutput {
	return o.ApplyT(func(v *WorkspaceKeyDetails) WorkspaceKeyDetails {
		if v != nil {
			return *v
		}
		var ret WorkspaceKeyDetails
		return ret
	}).(WorkspaceKeyDetailsOutput)
}

func (o WorkspaceKeyDetailsPtrOutput) KeyVaultUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceKeyDetails) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultUrl
	}).(pulumi.StringPtrOutput)
}

func (o WorkspaceKeyDetailsPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceKeyDetails) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type WorkspaceKeyDetailsResponse struct {
	KeyVaultUrl *string `pulumi:"keyVaultUrl"`
	Name        *string `pulumi:"name"`
}

type WorkspaceKeyDetailsResponseOutput struct{ *pulumi.OutputState }

func (WorkspaceKeyDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceKeyDetailsResponse)(nil)).Elem()
}

func (o WorkspaceKeyDetailsResponseOutput) ToWorkspaceKeyDetailsResponseOutput() WorkspaceKeyDetailsResponseOutput {
	return o
}

func (o WorkspaceKeyDetailsResponseOutput) ToWorkspaceKeyDetailsResponseOutputWithContext(ctx context.Context) WorkspaceKeyDetailsResponseOutput {
	return o
}

func (o WorkspaceKeyDetailsResponseOutput) KeyVaultUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkspaceKeyDetailsResponse) *string { return v.KeyVaultUrl }).(pulumi.StringPtrOutput)
}

func (o WorkspaceKeyDetailsResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkspaceKeyDetailsResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type WorkspaceKeyDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (WorkspaceKeyDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceKeyDetailsResponse)(nil)).Elem()
}

func (o WorkspaceKeyDetailsResponsePtrOutput) ToWorkspaceKeyDetailsResponsePtrOutput() WorkspaceKeyDetailsResponsePtrOutput {
	return o
}

func (o WorkspaceKeyDetailsResponsePtrOutput) ToWorkspaceKeyDetailsResponsePtrOutputWithContext(ctx context.Context) WorkspaceKeyDetailsResponsePtrOutput {
	return o
}

func (o WorkspaceKeyDetailsResponsePtrOutput) Elem() WorkspaceKeyDetailsResponseOutput {
	return o.ApplyT(func(v *WorkspaceKeyDetailsResponse) WorkspaceKeyDetailsResponse {
		if v != nil {
			return *v
		}
		var ret WorkspaceKeyDetailsResponse
		return ret
	}).(WorkspaceKeyDetailsResponseOutput)
}

func (o WorkspaceKeyDetailsResponsePtrOutput) KeyVaultUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceKeyDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultUrl
	}).(pulumi.StringPtrOutput)
}

func (o WorkspaceKeyDetailsResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceKeyDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type WorkspaceRepositoryConfiguration struct {
	AccountName         *string `pulumi:"accountName"`
	CollaborationBranch *string `pulumi:"collaborationBranch"`
	HostName            *string `pulumi:"hostName"`
	LastCommitId        *string `pulumi:"lastCommitId"`
	ProjectName         *string `pulumi:"projectName"`
	RepositoryName      *string `pulumi:"repositoryName"`
	RootFolder          *string `pulumi:"rootFolder"`
	TenantId            *string `pulumi:"tenantId"`
	Type                *string `pulumi:"type"`
}





type WorkspaceRepositoryConfigurationInput interface {
	pulumi.Input

	ToWorkspaceRepositoryConfigurationOutput() WorkspaceRepositoryConfigurationOutput
	ToWorkspaceRepositoryConfigurationOutputWithContext(context.Context) WorkspaceRepositoryConfigurationOutput
}

type WorkspaceRepositoryConfigurationArgs struct {
	AccountName         pulumi.StringPtrInput `pulumi:"accountName"`
	CollaborationBranch pulumi.StringPtrInput `pulumi:"collaborationBranch"`
	HostName            pulumi.StringPtrInput `pulumi:"hostName"`
	LastCommitId        pulumi.StringPtrInput `pulumi:"lastCommitId"`
	ProjectName         pulumi.StringPtrInput `pulumi:"projectName"`
	RepositoryName      pulumi.StringPtrInput `pulumi:"repositoryName"`
	RootFolder          pulumi.StringPtrInput `pulumi:"rootFolder"`
	TenantId            pulumi.StringPtrInput `pulumi:"tenantId"`
	Type                pulumi.StringPtrInput `pulumi:"type"`
}

func (WorkspaceRepositoryConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceRepositoryConfiguration)(nil)).Elem()
}

func (i WorkspaceRepositoryConfigurationArgs) ToWorkspaceRepositoryConfigurationOutput() WorkspaceRepositoryConfigurationOutput {
	return i.ToWorkspaceRepositoryConfigurationOutputWithContext(context.Background())
}

func (i WorkspaceRepositoryConfigurationArgs) ToWorkspaceRepositoryConfigurationOutputWithContext(ctx context.Context) WorkspaceRepositoryConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceRepositoryConfigurationOutput)
}

func (i WorkspaceRepositoryConfigurationArgs) ToWorkspaceRepositoryConfigurationPtrOutput() WorkspaceRepositoryConfigurationPtrOutput {
	return i.ToWorkspaceRepositoryConfigurationPtrOutputWithContext(context.Background())
}

func (i WorkspaceRepositoryConfigurationArgs) ToWorkspaceRepositoryConfigurationPtrOutputWithContext(ctx context.Context) WorkspaceRepositoryConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceRepositoryConfigurationOutput).ToWorkspaceRepositoryConfigurationPtrOutputWithContext(ctx)
}









type WorkspaceRepositoryConfigurationPtrInput interface {
	pulumi.Input

	ToWorkspaceRepositoryConfigurationPtrOutput() WorkspaceRepositoryConfigurationPtrOutput
	ToWorkspaceRepositoryConfigurationPtrOutputWithContext(context.Context) WorkspaceRepositoryConfigurationPtrOutput
}

type workspaceRepositoryConfigurationPtrType WorkspaceRepositoryConfigurationArgs

func WorkspaceRepositoryConfigurationPtr(v *WorkspaceRepositoryConfigurationArgs) WorkspaceRepositoryConfigurationPtrInput {
	return (*workspaceRepositoryConfigurationPtrType)(v)
}

func (*workspaceRepositoryConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceRepositoryConfiguration)(nil)).Elem()
}

func (i *workspaceRepositoryConfigurationPtrType) ToWorkspaceRepositoryConfigurationPtrOutput() WorkspaceRepositoryConfigurationPtrOutput {
	return i.ToWorkspaceRepositoryConfigurationPtrOutputWithContext(context.Background())
}

func (i *workspaceRepositoryConfigurationPtrType) ToWorkspaceRepositoryConfigurationPtrOutputWithContext(ctx context.Context) WorkspaceRepositoryConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceRepositoryConfigurationPtrOutput)
}

type WorkspaceRepositoryConfigurationOutput struct{ *pulumi.OutputState }

func (WorkspaceRepositoryConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceRepositoryConfiguration)(nil)).Elem()
}

func (o WorkspaceRepositoryConfigurationOutput) ToWorkspaceRepositoryConfigurationOutput() WorkspaceRepositoryConfigurationOutput {
	return o
}

func (o WorkspaceRepositoryConfigurationOutput) ToWorkspaceRepositoryConfigurationOutputWithContext(ctx context.Context) WorkspaceRepositoryConfigurationOutput {
	return o
}

func (o WorkspaceRepositoryConfigurationOutput) ToWorkspaceRepositoryConfigurationPtrOutput() WorkspaceRepositoryConfigurationPtrOutput {
	return o.ToWorkspaceRepositoryConfigurationPtrOutputWithContext(context.Background())
}

func (o WorkspaceRepositoryConfigurationOutput) ToWorkspaceRepositoryConfigurationPtrOutputWithContext(ctx context.Context) WorkspaceRepositoryConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkspaceRepositoryConfiguration) *WorkspaceRepositoryConfiguration {
		return &v
	}).(WorkspaceRepositoryConfigurationPtrOutput)
}

func (o WorkspaceRepositoryConfigurationOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkspaceRepositoryConfiguration) *string { return v.AccountName }).(pulumi.StringPtrOutput)
}

func (o WorkspaceRepositoryConfigurationOutput) CollaborationBranch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkspaceRepositoryConfiguration) *string { return v.CollaborationBranch }).(pulumi.StringPtrOutput)
}

func (o WorkspaceRepositoryConfigurationOutput) HostName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkspaceRepositoryConfiguration) *string { return v.HostName }).(pulumi.StringPtrOutput)
}

func (o WorkspaceRepositoryConfigurationOutput) LastCommitId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkspaceRepositoryConfiguration) *string { return v.LastCommitId }).(pulumi.StringPtrOutput)
}

func (o WorkspaceRepositoryConfigurationOutput) ProjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkspaceRepositoryConfiguration) *string { return v.ProjectName }).(pulumi.StringPtrOutput)
}

func (o WorkspaceRepositoryConfigurationOutput) RepositoryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkspaceRepositoryConfiguration) *string { return v.RepositoryName }).(pulumi.StringPtrOutput)
}

func (o WorkspaceRepositoryConfigurationOutput) RootFolder() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkspaceRepositoryConfiguration) *string { return v.RootFolder }).(pulumi.StringPtrOutput)
}

func (o WorkspaceRepositoryConfigurationOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkspaceRepositoryConfiguration) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o WorkspaceRepositoryConfigurationOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkspaceRepositoryConfiguration) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type WorkspaceRepositoryConfigurationPtrOutput struct{ *pulumi.OutputState }

func (WorkspaceRepositoryConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceRepositoryConfiguration)(nil)).Elem()
}

func (o WorkspaceRepositoryConfigurationPtrOutput) ToWorkspaceRepositoryConfigurationPtrOutput() WorkspaceRepositoryConfigurationPtrOutput {
	return o
}

func (o WorkspaceRepositoryConfigurationPtrOutput) ToWorkspaceRepositoryConfigurationPtrOutputWithContext(ctx context.Context) WorkspaceRepositoryConfigurationPtrOutput {
	return o
}

func (o WorkspaceRepositoryConfigurationPtrOutput) Elem() WorkspaceRepositoryConfigurationOutput {
	return o.ApplyT(func(v *WorkspaceRepositoryConfiguration) WorkspaceRepositoryConfiguration {
		if v != nil {
			return *v
		}
		var ret WorkspaceRepositoryConfiguration
		return ret
	}).(WorkspaceRepositoryConfigurationOutput)
}

func (o WorkspaceRepositoryConfigurationPtrOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceRepositoryConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.AccountName
	}).(pulumi.StringPtrOutput)
}

func (o WorkspaceRepositoryConfigurationPtrOutput) CollaborationBranch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceRepositoryConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.CollaborationBranch
	}).(pulumi.StringPtrOutput)
}

func (o WorkspaceRepositoryConfigurationPtrOutput) HostName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceRepositoryConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.HostName
	}).(pulumi.StringPtrOutput)
}

func (o WorkspaceRepositoryConfigurationPtrOutput) LastCommitId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceRepositoryConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.LastCommitId
	}).(pulumi.StringPtrOutput)
}

func (o WorkspaceRepositoryConfigurationPtrOutput) ProjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceRepositoryConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.ProjectName
	}).(pulumi.StringPtrOutput)
}

func (o WorkspaceRepositoryConfigurationPtrOutput) RepositoryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceRepositoryConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.RepositoryName
	}).(pulumi.StringPtrOutput)
}

func (o WorkspaceRepositoryConfigurationPtrOutput) RootFolder() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceRepositoryConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.RootFolder
	}).(pulumi.StringPtrOutput)
}

func (o WorkspaceRepositoryConfigurationPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceRepositoryConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o WorkspaceRepositoryConfigurationPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceRepositoryConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type WorkspaceRepositoryConfigurationResponse struct {
	AccountName         *string `pulumi:"accountName"`
	CollaborationBranch *string `pulumi:"collaborationBranch"`
	HostName            *string `pulumi:"hostName"`
	LastCommitId        *string `pulumi:"lastCommitId"`
	ProjectName         *string `pulumi:"projectName"`
	RepositoryName      *string `pulumi:"repositoryName"`
	RootFolder          *string `pulumi:"rootFolder"`
	TenantId            *string `pulumi:"tenantId"`
	Type                *string `pulumi:"type"`
}

type WorkspaceRepositoryConfigurationResponseOutput struct{ *pulumi.OutputState }

func (WorkspaceRepositoryConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceRepositoryConfigurationResponse)(nil)).Elem()
}

func (o WorkspaceRepositoryConfigurationResponseOutput) ToWorkspaceRepositoryConfigurationResponseOutput() WorkspaceRepositoryConfigurationResponseOutput {
	return o
}

func (o WorkspaceRepositoryConfigurationResponseOutput) ToWorkspaceRepositoryConfigurationResponseOutputWithContext(ctx context.Context) WorkspaceRepositoryConfigurationResponseOutput {
	return o
}

func (o WorkspaceRepositoryConfigurationResponseOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkspaceRepositoryConfigurationResponse) *string { return v.AccountName }).(pulumi.StringPtrOutput)
}

func (o WorkspaceRepositoryConfigurationResponseOutput) CollaborationBranch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkspaceRepositoryConfigurationResponse) *string { return v.CollaborationBranch }).(pulumi.StringPtrOutput)
}

func (o WorkspaceRepositoryConfigurationResponseOutput) HostName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkspaceRepositoryConfigurationResponse) *string { return v.HostName }).(pulumi.StringPtrOutput)
}

func (o WorkspaceRepositoryConfigurationResponseOutput) LastCommitId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkspaceRepositoryConfigurationResponse) *string { return v.LastCommitId }).(pulumi.StringPtrOutput)
}

func (o WorkspaceRepositoryConfigurationResponseOutput) ProjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkspaceRepositoryConfigurationResponse) *string { return v.ProjectName }).(pulumi.StringPtrOutput)
}

func (o WorkspaceRepositoryConfigurationResponseOutput) RepositoryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkspaceRepositoryConfigurationResponse) *string { return v.RepositoryName }).(pulumi.StringPtrOutput)
}

func (o WorkspaceRepositoryConfigurationResponseOutput) RootFolder() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkspaceRepositoryConfigurationResponse) *string { return v.RootFolder }).(pulumi.StringPtrOutput)
}

func (o WorkspaceRepositoryConfigurationResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkspaceRepositoryConfigurationResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o WorkspaceRepositoryConfigurationResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkspaceRepositoryConfigurationResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type WorkspaceRepositoryConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (WorkspaceRepositoryConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceRepositoryConfigurationResponse)(nil)).Elem()
}

func (o WorkspaceRepositoryConfigurationResponsePtrOutput) ToWorkspaceRepositoryConfigurationResponsePtrOutput() WorkspaceRepositoryConfigurationResponsePtrOutput {
	return o
}

func (o WorkspaceRepositoryConfigurationResponsePtrOutput) ToWorkspaceRepositoryConfigurationResponsePtrOutputWithContext(ctx context.Context) WorkspaceRepositoryConfigurationResponsePtrOutput {
	return o
}

func (o WorkspaceRepositoryConfigurationResponsePtrOutput) Elem() WorkspaceRepositoryConfigurationResponseOutput {
	return o.ApplyT(func(v *WorkspaceRepositoryConfigurationResponse) WorkspaceRepositoryConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret WorkspaceRepositoryConfigurationResponse
		return ret
	}).(WorkspaceRepositoryConfigurationResponseOutput)
}

func (o WorkspaceRepositoryConfigurationResponsePtrOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceRepositoryConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.AccountName
	}).(pulumi.StringPtrOutput)
}

func (o WorkspaceRepositoryConfigurationResponsePtrOutput) CollaborationBranch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceRepositoryConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.CollaborationBranch
	}).(pulumi.StringPtrOutput)
}

func (o WorkspaceRepositoryConfigurationResponsePtrOutput) HostName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceRepositoryConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.HostName
	}).(pulumi.StringPtrOutput)
}

func (o WorkspaceRepositoryConfigurationResponsePtrOutput) LastCommitId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceRepositoryConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastCommitId
	}).(pulumi.StringPtrOutput)
}

func (o WorkspaceRepositoryConfigurationResponsePtrOutput) ProjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceRepositoryConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProjectName
	}).(pulumi.StringPtrOutput)
}

func (o WorkspaceRepositoryConfigurationResponsePtrOutput) RepositoryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceRepositoryConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.RepositoryName
	}).(pulumi.StringPtrOutput)
}

func (o WorkspaceRepositoryConfigurationResponsePtrOutput) RootFolder() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceRepositoryConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.RootFolder
	}).(pulumi.StringPtrOutput)
}

func (o WorkspaceRepositoryConfigurationResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceRepositoryConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o WorkspaceRepositoryConfigurationResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceRepositoryConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AutoPausePropertiesOutput{})
	pulumi.RegisterOutputType(AutoPausePropertiesPtrOutput{})
	pulumi.RegisterOutputType(AutoPausePropertiesResponseOutput{})
	pulumi.RegisterOutputType(AutoPausePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(AutoScalePropertiesOutput{})
	pulumi.RegisterOutputType(AutoScalePropertiesPtrOutput{})
	pulumi.RegisterOutputType(AutoScalePropertiesResponseOutput{})
	pulumi.RegisterOutputType(AutoScalePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(AzureSkuOutput{})
	pulumi.RegisterOutputType(AzureSkuResponseOutput{})
	pulumi.RegisterOutputType(CspWorkspaceAdminPropertiesOutput{})
	pulumi.RegisterOutputType(CspWorkspaceAdminPropertiesPtrOutput{})
	pulumi.RegisterOutputType(CspWorkspaceAdminPropertiesResponseOutput{})
	pulumi.RegisterOutputType(CspWorkspaceAdminPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(CustomerManagedKeyDetailsOutput{})
	pulumi.RegisterOutputType(CustomerManagedKeyDetailsPtrOutput{})
	pulumi.RegisterOutputType(CustomerManagedKeyDetailsResponseOutput{})
	pulumi.RegisterOutputType(CustomerManagedKeyDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(DataLakeStorageAccountDetailsOutput{})
	pulumi.RegisterOutputType(DataLakeStorageAccountDetailsPtrOutput{})
	pulumi.RegisterOutputType(DataLakeStorageAccountDetailsResponseOutput{})
	pulumi.RegisterOutputType(DataLakeStorageAccountDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(DatabaseStatisticsResponseOutput{})
	pulumi.RegisterOutputType(DynamicExecutorAllocationOutput{})
	pulumi.RegisterOutputType(DynamicExecutorAllocationPtrOutput{})
	pulumi.RegisterOutputType(DynamicExecutorAllocationResponseOutput{})
	pulumi.RegisterOutputType(DynamicExecutorAllocationResponsePtrOutput{})
	pulumi.RegisterOutputType(EncryptionDetailsOutput{})
	pulumi.RegisterOutputType(EncryptionDetailsPtrOutput{})
	pulumi.RegisterOutputType(EncryptionDetailsResponseOutput{})
	pulumi.RegisterOutputType(EncryptionDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(KekIdentityPropertiesOutput{})
	pulumi.RegisterOutputType(KekIdentityPropertiesPtrOutput{})
	pulumi.RegisterOutputType(KekIdentityPropertiesResponseOutput{})
	pulumi.RegisterOutputType(KekIdentityPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(LanguageExtensionResponseOutput{})
	pulumi.RegisterOutputType(LanguageExtensionResponseArrayOutput{})
	pulumi.RegisterOutputType(LanguageExtensionsListResponseOutput{})
	pulumi.RegisterOutputType(LibraryInfoOutput{})
	pulumi.RegisterOutputType(LibraryInfoArrayOutput{})
	pulumi.RegisterOutputType(LibraryInfoResponseOutput{})
	pulumi.RegisterOutputType(LibraryInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(LibraryRequirementsOutput{})
	pulumi.RegisterOutputType(LibraryRequirementsPtrOutput{})
	pulumi.RegisterOutputType(LibraryRequirementsResponseOutput{})
	pulumi.RegisterOutputType(LibraryRequirementsResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagedIdentityOutput{})
	pulumi.RegisterOutputType(ManagedIdentityPtrOutput{})
	pulumi.RegisterOutputType(ManagedIdentityResponseOutput{})
	pulumi.RegisterOutputType(ManagedIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagedVirtualNetworkSettingsOutput{})
	pulumi.RegisterOutputType(ManagedVirtualNetworkSettingsPtrOutput{})
	pulumi.RegisterOutputType(ManagedVirtualNetworkSettingsResponseOutput{})
	pulumi.RegisterOutputType(ManagedVirtualNetworkSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(OptimizedAutoscaleOutput{})
	pulumi.RegisterOutputType(OptimizedAutoscalePtrOutput{})
	pulumi.RegisterOutputType(OptimizedAutoscaleResponseOutput{})
	pulumi.RegisterOutputType(OptimizedAutoscaleResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionTypeOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionTypeArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionForPrivateLinkHubBasicResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionForPrivateLinkHubBasicResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponsePtrOutput{})
	pulumi.RegisterOutputType(PurviewConfigurationOutput{})
	pulumi.RegisterOutputType(PurviewConfigurationPtrOutput{})
	pulumi.RegisterOutputType(PurviewConfigurationResponseOutput{})
	pulumi.RegisterOutputType(PurviewConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SparkConfigPropertiesOutput{})
	pulumi.RegisterOutputType(SparkConfigPropertiesPtrOutput{})
	pulumi.RegisterOutputType(SparkConfigPropertiesResponseOutput{})
	pulumi.RegisterOutputType(SparkConfigPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(SqlPoolVulnerabilityAssessmentRuleBaselineItemOutput{})
	pulumi.RegisterOutputType(SqlPoolVulnerabilityAssessmentRuleBaselineItemArrayOutput{})
	pulumi.RegisterOutputType(SqlPoolVulnerabilityAssessmentRuleBaselineItemResponseOutput{})
	pulumi.RegisterOutputType(SqlPoolVulnerabilityAssessmentRuleBaselineItemResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(TableLevelSharingPropertiesOutput{})
	pulumi.RegisterOutputType(TableLevelSharingPropertiesPtrOutput{})
	pulumi.RegisterOutputType(TableLevelSharingPropertiesResponseOutput{})
	pulumi.RegisterOutputType(TableLevelSharingPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(UserAssignedManagedIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedManagedIdentityResponseMapOutput{})
	pulumi.RegisterOutputType(VirtualNetworkProfileOutput{})
	pulumi.RegisterOutputType(VirtualNetworkProfilePtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkProfileResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(VulnerabilityAssessmentRecurringScansPropertiesOutput{})
	pulumi.RegisterOutputType(VulnerabilityAssessmentRecurringScansPropertiesPtrOutput{})
	pulumi.RegisterOutputType(VulnerabilityAssessmentRecurringScansPropertiesResponseOutput{})
	pulumi.RegisterOutputType(VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(WorkspaceKeyDetailsOutput{})
	pulumi.RegisterOutputType(WorkspaceKeyDetailsPtrOutput{})
	pulumi.RegisterOutputType(WorkspaceKeyDetailsResponseOutput{})
	pulumi.RegisterOutputType(WorkspaceKeyDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(WorkspaceRepositoryConfigurationOutput{})
	pulumi.RegisterOutputType(WorkspaceRepositoryConfigurationPtrOutput{})
	pulumi.RegisterOutputType(WorkspaceRepositoryConfigurationResponseOutput{})
	pulumi.RegisterOutputType(WorkspaceRepositoryConfigurationResponsePtrOutput{})
}
