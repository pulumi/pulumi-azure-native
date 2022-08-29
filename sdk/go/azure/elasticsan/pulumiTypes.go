


package elasticsan

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IscsiTargetInfoResponse struct {
	ProvisioningState    string  `pulumi:"provisioningState"`
	Status               *string `pulumi:"status"`
	TargetIqn            string  `pulumi:"targetIqn"`
	TargetPortalHostname string  `pulumi:"targetPortalHostname"`
	TargetPortalPort     int     `pulumi:"targetPortalPort"`
}

type IscsiTargetInfoResponseOutput struct{ *pulumi.OutputState }

func (IscsiTargetInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IscsiTargetInfoResponse)(nil)).Elem()
}

func (o IscsiTargetInfoResponseOutput) ToIscsiTargetInfoResponseOutput() IscsiTargetInfoResponseOutput {
	return o
}

func (o IscsiTargetInfoResponseOutput) ToIscsiTargetInfoResponseOutputWithContext(ctx context.Context) IscsiTargetInfoResponseOutput {
	return o
}

func (o IscsiTargetInfoResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v IscsiTargetInfoResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o IscsiTargetInfoResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IscsiTargetInfoResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o IscsiTargetInfoResponseOutput) TargetIqn() pulumi.StringOutput {
	return o.ApplyT(func(v IscsiTargetInfoResponse) string { return v.TargetIqn }).(pulumi.StringOutput)
}

func (o IscsiTargetInfoResponseOutput) TargetPortalHostname() pulumi.StringOutput {
	return o.ApplyT(func(v IscsiTargetInfoResponse) string { return v.TargetPortalHostname }).(pulumi.StringOutput)
}

func (o IscsiTargetInfoResponseOutput) TargetPortalPort() pulumi.IntOutput {
	return o.ApplyT(func(v IscsiTargetInfoResponse) int { return v.TargetPortalPort }).(pulumi.IntOutput)
}

type NetworkRuleSet struct {
	VirtualNetworkRules []VirtualNetworkRule `pulumi:"virtualNetworkRules"`
}





type NetworkRuleSetInput interface {
	pulumi.Input

	ToNetworkRuleSetOutput() NetworkRuleSetOutput
	ToNetworkRuleSetOutputWithContext(context.Context) NetworkRuleSetOutput
}

type NetworkRuleSetArgs struct {
	VirtualNetworkRules VirtualNetworkRuleArrayInput `pulumi:"virtualNetworkRules"`
}

func (NetworkRuleSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSet)(nil)).Elem()
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetOutput() NetworkRuleSetOutput {
	return i.ToNetworkRuleSetOutputWithContext(context.Background())
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetOutputWithContext(ctx context.Context) NetworkRuleSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetOutput)
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return i.ToNetworkRuleSetPtrOutputWithContext(context.Background())
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetOutput).ToNetworkRuleSetPtrOutputWithContext(ctx)
}









type NetworkRuleSetPtrInput interface {
	pulumi.Input

	ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput
	ToNetworkRuleSetPtrOutputWithContext(context.Context) NetworkRuleSetPtrOutput
}

type networkRuleSetPtrType NetworkRuleSetArgs

func NetworkRuleSetPtr(v *NetworkRuleSetArgs) NetworkRuleSetPtrInput {
	return (*networkRuleSetPtrType)(v)
}

func (*networkRuleSetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSet)(nil)).Elem()
}

func (i *networkRuleSetPtrType) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return i.ToNetworkRuleSetPtrOutputWithContext(context.Background())
}

func (i *networkRuleSetPtrType) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetPtrOutput)
}

type NetworkRuleSetOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSet)(nil)).Elem()
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetOutput() NetworkRuleSetOutput {
	return o
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetOutputWithContext(ctx context.Context) NetworkRuleSetOutput {
	return o
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return o.ToNetworkRuleSetPtrOutputWithContext(context.Background())
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkRuleSet) *NetworkRuleSet {
		return &v
	}).(NetworkRuleSetPtrOutput)
}

func (o NetworkRuleSetOutput) VirtualNetworkRules() VirtualNetworkRuleArrayOutput {
	return o.ApplyT(func(v NetworkRuleSet) []VirtualNetworkRule { return v.VirtualNetworkRules }).(VirtualNetworkRuleArrayOutput)
}

type NetworkRuleSetPtrOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSet)(nil)).Elem()
}

func (o NetworkRuleSetPtrOutput) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return o
}

func (o NetworkRuleSetPtrOutput) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return o
}

func (o NetworkRuleSetPtrOutput) Elem() NetworkRuleSetOutput {
	return o.ApplyT(func(v *NetworkRuleSet) NetworkRuleSet {
		if v != nil {
			return *v
		}
		var ret NetworkRuleSet
		return ret
	}).(NetworkRuleSetOutput)
}

func (o NetworkRuleSetPtrOutput) VirtualNetworkRules() VirtualNetworkRuleArrayOutput {
	return o.ApplyT(func(v *NetworkRuleSet) []VirtualNetworkRule {
		if v == nil {
			return nil
		}
		return v.VirtualNetworkRules
	}).(VirtualNetworkRuleArrayOutput)
}

type NetworkRuleSetResponse struct {
	VirtualNetworkRules []VirtualNetworkRuleResponse `pulumi:"virtualNetworkRules"`
}

type NetworkRuleSetResponseOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSetResponse)(nil)).Elem()
}

func (o NetworkRuleSetResponseOutput) ToNetworkRuleSetResponseOutput() NetworkRuleSetResponseOutput {
	return o
}

func (o NetworkRuleSetResponseOutput) ToNetworkRuleSetResponseOutputWithContext(ctx context.Context) NetworkRuleSetResponseOutput {
	return o
}

func (o NetworkRuleSetResponseOutput) VirtualNetworkRules() VirtualNetworkRuleResponseArrayOutput {
	return o.ApplyT(func(v NetworkRuleSetResponse) []VirtualNetworkRuleResponse { return v.VirtualNetworkRules }).(VirtualNetworkRuleResponseArrayOutput)
}

type NetworkRuleSetResponsePtrOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSetResponse)(nil)).Elem()
}

func (o NetworkRuleSetResponsePtrOutput) ToNetworkRuleSetResponsePtrOutput() NetworkRuleSetResponsePtrOutput {
	return o
}

func (o NetworkRuleSetResponsePtrOutput) ToNetworkRuleSetResponsePtrOutputWithContext(ctx context.Context) NetworkRuleSetResponsePtrOutput {
	return o
}

func (o NetworkRuleSetResponsePtrOutput) Elem() NetworkRuleSetResponseOutput {
	return o.ApplyT(func(v *NetworkRuleSetResponse) NetworkRuleSetResponse {
		if v != nil {
			return *v
		}
		var ret NetworkRuleSetResponse
		return ret
	}).(NetworkRuleSetResponseOutput)
}

func (o NetworkRuleSetResponsePtrOutput) VirtualNetworkRules() VirtualNetworkRuleResponseArrayOutput {
	return o.ApplyT(func(v *NetworkRuleSetResponse) []VirtualNetworkRuleResponse {
		if v == nil {
			return nil
		}
		return v.VirtualNetworkRules
	}).(VirtualNetworkRuleResponseArrayOutput)
}

type Sku struct {
	Name string  `pulumi:"name"`
	Tier *string `pulumi:"tier"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Name pulumi.StringInput    `pulumi:"name"`
	Tier pulumi.StringPtrInput `pulumi:"tier"`
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

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Name string  `pulumi:"name"`
	Tier *string `pulumi:"tier"`
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

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SourceCreationData struct {
	CreateSource *VolumeCreateOption `pulumi:"createSource"`
	SourceUri    *string             `pulumi:"sourceUri"`
}





type SourceCreationDataInput interface {
	pulumi.Input

	ToSourceCreationDataOutput() SourceCreationDataOutput
	ToSourceCreationDataOutputWithContext(context.Context) SourceCreationDataOutput
}

type SourceCreationDataArgs struct {
	CreateSource VolumeCreateOptionPtrInput `pulumi:"createSource"`
	SourceUri    pulumi.StringPtrInput      `pulumi:"sourceUri"`
}

func (SourceCreationDataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceCreationData)(nil)).Elem()
}

func (i SourceCreationDataArgs) ToSourceCreationDataOutput() SourceCreationDataOutput {
	return i.ToSourceCreationDataOutputWithContext(context.Background())
}

func (i SourceCreationDataArgs) ToSourceCreationDataOutputWithContext(ctx context.Context) SourceCreationDataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceCreationDataOutput)
}

func (i SourceCreationDataArgs) ToSourceCreationDataPtrOutput() SourceCreationDataPtrOutput {
	return i.ToSourceCreationDataPtrOutputWithContext(context.Background())
}

func (i SourceCreationDataArgs) ToSourceCreationDataPtrOutputWithContext(ctx context.Context) SourceCreationDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceCreationDataOutput).ToSourceCreationDataPtrOutputWithContext(ctx)
}









type SourceCreationDataPtrInput interface {
	pulumi.Input

	ToSourceCreationDataPtrOutput() SourceCreationDataPtrOutput
	ToSourceCreationDataPtrOutputWithContext(context.Context) SourceCreationDataPtrOutput
}

type sourceCreationDataPtrType SourceCreationDataArgs

func SourceCreationDataPtr(v *SourceCreationDataArgs) SourceCreationDataPtrInput {
	return (*sourceCreationDataPtrType)(v)
}

func (*sourceCreationDataPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceCreationData)(nil)).Elem()
}

func (i *sourceCreationDataPtrType) ToSourceCreationDataPtrOutput() SourceCreationDataPtrOutput {
	return i.ToSourceCreationDataPtrOutputWithContext(context.Background())
}

func (i *sourceCreationDataPtrType) ToSourceCreationDataPtrOutputWithContext(ctx context.Context) SourceCreationDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceCreationDataPtrOutput)
}

type SourceCreationDataOutput struct{ *pulumi.OutputState }

func (SourceCreationDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceCreationData)(nil)).Elem()
}

func (o SourceCreationDataOutput) ToSourceCreationDataOutput() SourceCreationDataOutput {
	return o
}

func (o SourceCreationDataOutput) ToSourceCreationDataOutputWithContext(ctx context.Context) SourceCreationDataOutput {
	return o
}

func (o SourceCreationDataOutput) ToSourceCreationDataPtrOutput() SourceCreationDataPtrOutput {
	return o.ToSourceCreationDataPtrOutputWithContext(context.Background())
}

func (o SourceCreationDataOutput) ToSourceCreationDataPtrOutputWithContext(ctx context.Context) SourceCreationDataPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SourceCreationData) *SourceCreationData {
		return &v
	}).(SourceCreationDataPtrOutput)
}

func (o SourceCreationDataOutput) CreateSource() VolumeCreateOptionPtrOutput {
	return o.ApplyT(func(v SourceCreationData) *VolumeCreateOption { return v.CreateSource }).(VolumeCreateOptionPtrOutput)
}

func (o SourceCreationDataOutput) SourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SourceCreationData) *string { return v.SourceUri }).(pulumi.StringPtrOutput)
}

type SourceCreationDataPtrOutput struct{ *pulumi.OutputState }

func (SourceCreationDataPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceCreationData)(nil)).Elem()
}

func (o SourceCreationDataPtrOutput) ToSourceCreationDataPtrOutput() SourceCreationDataPtrOutput {
	return o
}

func (o SourceCreationDataPtrOutput) ToSourceCreationDataPtrOutputWithContext(ctx context.Context) SourceCreationDataPtrOutput {
	return o
}

func (o SourceCreationDataPtrOutput) Elem() SourceCreationDataOutput {
	return o.ApplyT(func(v *SourceCreationData) SourceCreationData {
		if v != nil {
			return *v
		}
		var ret SourceCreationData
		return ret
	}).(SourceCreationDataOutput)
}

func (o SourceCreationDataPtrOutput) CreateSource() VolumeCreateOptionPtrOutput {
	return o.ApplyT(func(v *SourceCreationData) *VolumeCreateOption {
		if v == nil {
			return nil
		}
		return v.CreateSource
	}).(VolumeCreateOptionPtrOutput)
}

func (o SourceCreationDataPtrOutput) SourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceCreationData) *string {
		if v == nil {
			return nil
		}
		return v.SourceUri
	}).(pulumi.StringPtrOutput)
}

type SourceCreationDataResponse struct {
	CreateSource *string `pulumi:"createSource"`
	SourceUri    *string `pulumi:"sourceUri"`
}

type SourceCreationDataResponseOutput struct{ *pulumi.OutputState }

func (SourceCreationDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceCreationDataResponse)(nil)).Elem()
}

func (o SourceCreationDataResponseOutput) ToSourceCreationDataResponseOutput() SourceCreationDataResponseOutput {
	return o
}

func (o SourceCreationDataResponseOutput) ToSourceCreationDataResponseOutputWithContext(ctx context.Context) SourceCreationDataResponseOutput {
	return o
}

func (o SourceCreationDataResponseOutput) CreateSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SourceCreationDataResponse) *string { return v.CreateSource }).(pulumi.StringPtrOutput)
}

func (o SourceCreationDataResponseOutput) SourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SourceCreationDataResponse) *string { return v.SourceUri }).(pulumi.StringPtrOutput)
}

type SourceCreationDataResponsePtrOutput struct{ *pulumi.OutputState }

func (SourceCreationDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceCreationDataResponse)(nil)).Elem()
}

func (o SourceCreationDataResponsePtrOutput) ToSourceCreationDataResponsePtrOutput() SourceCreationDataResponsePtrOutput {
	return o
}

func (o SourceCreationDataResponsePtrOutput) ToSourceCreationDataResponsePtrOutputWithContext(ctx context.Context) SourceCreationDataResponsePtrOutput {
	return o
}

func (o SourceCreationDataResponsePtrOutput) Elem() SourceCreationDataResponseOutput {
	return o.ApplyT(func(v *SourceCreationDataResponse) SourceCreationDataResponse {
		if v != nil {
			return *v
		}
		var ret SourceCreationDataResponse
		return ret
	}).(SourceCreationDataResponseOutput)
}

func (o SourceCreationDataResponsePtrOutput) CreateSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceCreationDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreateSource
	}).(pulumi.StringPtrOutput)
}

func (o SourceCreationDataResponsePtrOutput) SourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceCreationDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.SourceUri
	}).(pulumi.StringPtrOutput)
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

type VirtualNetworkRule struct {
	Action                   *Action `pulumi:"action"`
	VirtualNetworkResourceId string  `pulumi:"virtualNetworkResourceId"`
}


func (val *VirtualNetworkRule) Defaults() *VirtualNetworkRule {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Action) {
		action_ := Action("Allow")
		tmp.Action = &action_
	}
	return &tmp
}





type VirtualNetworkRuleInput interface {
	pulumi.Input

	ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput
	ToVirtualNetworkRuleOutputWithContext(context.Context) VirtualNetworkRuleOutput
}

type VirtualNetworkRuleArgs struct {
	Action                   ActionPtrInput     `pulumi:"action"`
	VirtualNetworkResourceId pulumi.StringInput `pulumi:"virtualNetworkResourceId"`
}


func (val *VirtualNetworkRuleArgs) Defaults() *VirtualNetworkRuleArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Action) {
		tmp.Action = Action("Allow")
	}
	return &tmp
}
func (VirtualNetworkRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRule)(nil)).Elem()
}

func (i VirtualNetworkRuleArgs) ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput {
	return i.ToVirtualNetworkRuleOutputWithContext(context.Background())
}

func (i VirtualNetworkRuleArgs) ToVirtualNetworkRuleOutputWithContext(ctx context.Context) VirtualNetworkRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRuleOutput)
}





type VirtualNetworkRuleArrayInput interface {
	pulumi.Input

	ToVirtualNetworkRuleArrayOutput() VirtualNetworkRuleArrayOutput
	ToVirtualNetworkRuleArrayOutputWithContext(context.Context) VirtualNetworkRuleArrayOutput
}

type VirtualNetworkRuleArray []VirtualNetworkRuleInput

func (VirtualNetworkRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkRule)(nil)).Elem()
}

func (i VirtualNetworkRuleArray) ToVirtualNetworkRuleArrayOutput() VirtualNetworkRuleArrayOutput {
	return i.ToVirtualNetworkRuleArrayOutputWithContext(context.Background())
}

func (i VirtualNetworkRuleArray) ToVirtualNetworkRuleArrayOutputWithContext(ctx context.Context) VirtualNetworkRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRuleArrayOutput)
}

type VirtualNetworkRuleOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRule)(nil)).Elem()
}

func (o VirtualNetworkRuleOutput) ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput {
	return o
}

func (o VirtualNetworkRuleOutput) ToVirtualNetworkRuleOutputWithContext(ctx context.Context) VirtualNetworkRuleOutput {
	return o
}

func (o VirtualNetworkRuleOutput) Action() ActionPtrOutput {
	return o.ApplyT(func(v VirtualNetworkRule) *Action { return v.Action }).(ActionPtrOutput)
}

func (o VirtualNetworkRuleOutput) VirtualNetworkResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkRule) string { return v.VirtualNetworkResourceId }).(pulumi.StringOutput)
}

type VirtualNetworkRuleArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkRule)(nil)).Elem()
}

func (o VirtualNetworkRuleArrayOutput) ToVirtualNetworkRuleArrayOutput() VirtualNetworkRuleArrayOutput {
	return o
}

func (o VirtualNetworkRuleArrayOutput) ToVirtualNetworkRuleArrayOutputWithContext(ctx context.Context) VirtualNetworkRuleArrayOutput {
	return o
}

func (o VirtualNetworkRuleArrayOutput) Index(i pulumi.IntInput) VirtualNetworkRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkRule {
		return vs[0].([]VirtualNetworkRule)[vs[1].(int)]
	}).(VirtualNetworkRuleOutput)
}

type VirtualNetworkRuleResponse struct {
	Action                   *string `pulumi:"action"`
	State                    string  `pulumi:"state"`
	VirtualNetworkResourceId string  `pulumi:"virtualNetworkResourceId"`
}


func (val *VirtualNetworkRuleResponse) Defaults() *VirtualNetworkRuleResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Action) {
		action_ := "Allow"
		tmp.Action = &action_
	}
	return &tmp
}

type VirtualNetworkRuleResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRuleResponse)(nil)).Elem()
}

func (o VirtualNetworkRuleResponseOutput) ToVirtualNetworkRuleResponseOutput() VirtualNetworkRuleResponseOutput {
	return o
}

func (o VirtualNetworkRuleResponseOutput) ToVirtualNetworkRuleResponseOutputWithContext(ctx context.Context) VirtualNetworkRuleResponseOutput {
	return o
}

func (o VirtualNetworkRuleResponseOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkRuleResponse) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkRuleResponseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkRuleResponse) string { return v.State }).(pulumi.StringOutput)
}

func (o VirtualNetworkRuleResponseOutput) VirtualNetworkResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkRuleResponse) string { return v.VirtualNetworkResourceId }).(pulumi.StringOutput)
}

type VirtualNetworkRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkRuleResponse)(nil)).Elem()
}

func (o VirtualNetworkRuleResponseArrayOutput) ToVirtualNetworkRuleResponseArrayOutput() VirtualNetworkRuleResponseArrayOutput {
	return o
}

func (o VirtualNetworkRuleResponseArrayOutput) ToVirtualNetworkRuleResponseArrayOutputWithContext(ctx context.Context) VirtualNetworkRuleResponseArrayOutput {
	return o
}

func (o VirtualNetworkRuleResponseArrayOutput) Index(i pulumi.IntInput) VirtualNetworkRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkRuleResponse {
		return vs[0].([]VirtualNetworkRuleResponse)[vs[1].(int)]
	}).(VirtualNetworkRuleResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(IscsiTargetInfoResponseOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetPtrOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetResponseOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SourceCreationDataOutput{})
	pulumi.RegisterOutputType(SourceCreationDataPtrOutput{})
	pulumi.RegisterOutputType(SourceCreationDataResponseOutput{})
	pulumi.RegisterOutputType(SourceCreationDataResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleResponseArrayOutput{})
}
