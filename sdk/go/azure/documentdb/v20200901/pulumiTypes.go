


package v20200901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApiProperties struct {
	ServerVersion *string `pulumi:"serverVersion"`
}





type ApiPropertiesInput interface {
	pulumi.Input

	ToApiPropertiesOutput() ApiPropertiesOutput
	ToApiPropertiesOutputWithContext(context.Context) ApiPropertiesOutput
}

type ApiPropertiesArgs struct {
	ServerVersion pulumi.StringPtrInput `pulumi:"serverVersion"`
}

func (ApiPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiProperties)(nil)).Elem()
}

func (i ApiPropertiesArgs) ToApiPropertiesOutput() ApiPropertiesOutput {
	return i.ToApiPropertiesOutputWithContext(context.Background())
}

func (i ApiPropertiesArgs) ToApiPropertiesOutputWithContext(ctx context.Context) ApiPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiPropertiesOutput)
}

func (i ApiPropertiesArgs) ToApiPropertiesPtrOutput() ApiPropertiesPtrOutput {
	return i.ToApiPropertiesPtrOutputWithContext(context.Background())
}

func (i ApiPropertiesArgs) ToApiPropertiesPtrOutputWithContext(ctx context.Context) ApiPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiPropertiesOutput).ToApiPropertiesPtrOutputWithContext(ctx)
}









type ApiPropertiesPtrInput interface {
	pulumi.Input

	ToApiPropertiesPtrOutput() ApiPropertiesPtrOutput
	ToApiPropertiesPtrOutputWithContext(context.Context) ApiPropertiesPtrOutput
}

type apiPropertiesPtrType ApiPropertiesArgs

func ApiPropertiesPtr(v *ApiPropertiesArgs) ApiPropertiesPtrInput {
	return (*apiPropertiesPtrType)(v)
}

func (*apiPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiProperties)(nil)).Elem()
}

func (i *apiPropertiesPtrType) ToApiPropertiesPtrOutput() ApiPropertiesPtrOutput {
	return i.ToApiPropertiesPtrOutputWithContext(context.Background())
}

func (i *apiPropertiesPtrType) ToApiPropertiesPtrOutputWithContext(ctx context.Context) ApiPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiPropertiesPtrOutput)
}

type ApiPropertiesOutput struct{ *pulumi.OutputState }

func (ApiPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiProperties)(nil)).Elem()
}

func (o ApiPropertiesOutput) ToApiPropertiesOutput() ApiPropertiesOutput {
	return o
}

func (o ApiPropertiesOutput) ToApiPropertiesOutputWithContext(ctx context.Context) ApiPropertiesOutput {
	return o
}

func (o ApiPropertiesOutput) ToApiPropertiesPtrOutput() ApiPropertiesPtrOutput {
	return o.ToApiPropertiesPtrOutputWithContext(context.Background())
}

func (o ApiPropertiesOutput) ToApiPropertiesPtrOutputWithContext(ctx context.Context) ApiPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiProperties) *ApiProperties {
		return &v
	}).(ApiPropertiesPtrOutput)
}

func (o ApiPropertiesOutput) ServerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiProperties) *string { return v.ServerVersion }).(pulumi.StringPtrOutput)
}

type ApiPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ApiPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiProperties)(nil)).Elem()
}

func (o ApiPropertiesPtrOutput) ToApiPropertiesPtrOutput() ApiPropertiesPtrOutput {
	return o
}

func (o ApiPropertiesPtrOutput) ToApiPropertiesPtrOutputWithContext(ctx context.Context) ApiPropertiesPtrOutput {
	return o
}

func (o ApiPropertiesPtrOutput) Elem() ApiPropertiesOutput {
	return o.ApplyT(func(v *ApiProperties) ApiProperties {
		if v != nil {
			return *v
		}
		var ret ApiProperties
		return ret
	}).(ApiPropertiesOutput)
}

func (o ApiPropertiesPtrOutput) ServerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiProperties) *string {
		if v == nil {
			return nil
		}
		return v.ServerVersion
	}).(pulumi.StringPtrOutput)
}

type ApiPropertiesResponse struct {
	ServerVersion *string `pulumi:"serverVersion"`
}

type ApiPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ApiPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiPropertiesResponse)(nil)).Elem()
}

func (o ApiPropertiesResponseOutput) ToApiPropertiesResponseOutput() ApiPropertiesResponseOutput {
	return o
}

func (o ApiPropertiesResponseOutput) ToApiPropertiesResponseOutputWithContext(ctx context.Context) ApiPropertiesResponseOutput {
	return o
}

func (o ApiPropertiesResponseOutput) ServerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiPropertiesResponse) *string { return v.ServerVersion }).(pulumi.StringPtrOutput)
}

type ApiPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ApiPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiPropertiesResponse)(nil)).Elem()
}

func (o ApiPropertiesResponsePtrOutput) ToApiPropertiesResponsePtrOutput() ApiPropertiesResponsePtrOutput {
	return o
}

func (o ApiPropertiesResponsePtrOutput) ToApiPropertiesResponsePtrOutputWithContext(ctx context.Context) ApiPropertiesResponsePtrOutput {
	return o
}

func (o ApiPropertiesResponsePtrOutput) Elem() ApiPropertiesResponseOutput {
	return o.ApplyT(func(v *ApiPropertiesResponse) ApiPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ApiPropertiesResponse
		return ret
	}).(ApiPropertiesResponseOutput)
}

func (o ApiPropertiesResponsePtrOutput) ServerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ServerVersion
	}).(pulumi.StringPtrOutput)
}

type AutoscaleSettings struct {
	MaxThroughput *int `pulumi:"maxThroughput"`
}





type AutoscaleSettingsInput interface {
	pulumi.Input

	ToAutoscaleSettingsOutput() AutoscaleSettingsOutput
	ToAutoscaleSettingsOutputWithContext(context.Context) AutoscaleSettingsOutput
}

type AutoscaleSettingsArgs struct {
	MaxThroughput pulumi.IntPtrInput `pulumi:"maxThroughput"`
}

func (AutoscaleSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoscaleSettings)(nil)).Elem()
}

func (i AutoscaleSettingsArgs) ToAutoscaleSettingsOutput() AutoscaleSettingsOutput {
	return i.ToAutoscaleSettingsOutputWithContext(context.Background())
}

func (i AutoscaleSettingsArgs) ToAutoscaleSettingsOutputWithContext(ctx context.Context) AutoscaleSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscaleSettingsOutput)
}

func (i AutoscaleSettingsArgs) ToAutoscaleSettingsPtrOutput() AutoscaleSettingsPtrOutput {
	return i.ToAutoscaleSettingsPtrOutputWithContext(context.Background())
}

func (i AutoscaleSettingsArgs) ToAutoscaleSettingsPtrOutputWithContext(ctx context.Context) AutoscaleSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscaleSettingsOutput).ToAutoscaleSettingsPtrOutputWithContext(ctx)
}









type AutoscaleSettingsPtrInput interface {
	pulumi.Input

	ToAutoscaleSettingsPtrOutput() AutoscaleSettingsPtrOutput
	ToAutoscaleSettingsPtrOutputWithContext(context.Context) AutoscaleSettingsPtrOutput
}

type autoscaleSettingsPtrType AutoscaleSettingsArgs

func AutoscaleSettingsPtr(v *AutoscaleSettingsArgs) AutoscaleSettingsPtrInput {
	return (*autoscaleSettingsPtrType)(v)
}

func (*autoscaleSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoscaleSettings)(nil)).Elem()
}

func (i *autoscaleSettingsPtrType) ToAutoscaleSettingsPtrOutput() AutoscaleSettingsPtrOutput {
	return i.ToAutoscaleSettingsPtrOutputWithContext(context.Background())
}

func (i *autoscaleSettingsPtrType) ToAutoscaleSettingsPtrOutputWithContext(ctx context.Context) AutoscaleSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscaleSettingsPtrOutput)
}

type AutoscaleSettingsOutput struct{ *pulumi.OutputState }

func (AutoscaleSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoscaleSettings)(nil)).Elem()
}

func (o AutoscaleSettingsOutput) ToAutoscaleSettingsOutput() AutoscaleSettingsOutput {
	return o
}

func (o AutoscaleSettingsOutput) ToAutoscaleSettingsOutputWithContext(ctx context.Context) AutoscaleSettingsOutput {
	return o
}

func (o AutoscaleSettingsOutput) ToAutoscaleSettingsPtrOutput() AutoscaleSettingsPtrOutput {
	return o.ToAutoscaleSettingsPtrOutputWithContext(context.Background())
}

func (o AutoscaleSettingsOutput) ToAutoscaleSettingsPtrOutputWithContext(ctx context.Context) AutoscaleSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoscaleSettings) *AutoscaleSettings {
		return &v
	}).(AutoscaleSettingsPtrOutput)
}

func (o AutoscaleSettingsOutput) MaxThroughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoscaleSettings) *int { return v.MaxThroughput }).(pulumi.IntPtrOutput)
}

type AutoscaleSettingsPtrOutput struct{ *pulumi.OutputState }

func (AutoscaleSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoscaleSettings)(nil)).Elem()
}

func (o AutoscaleSettingsPtrOutput) ToAutoscaleSettingsPtrOutput() AutoscaleSettingsPtrOutput {
	return o
}

func (o AutoscaleSettingsPtrOutput) ToAutoscaleSettingsPtrOutputWithContext(ctx context.Context) AutoscaleSettingsPtrOutput {
	return o
}

func (o AutoscaleSettingsPtrOutput) Elem() AutoscaleSettingsOutput {
	return o.ApplyT(func(v *AutoscaleSettings) AutoscaleSettings {
		if v != nil {
			return *v
		}
		var ret AutoscaleSettings
		return ret
	}).(AutoscaleSettingsOutput)
}

func (o AutoscaleSettingsPtrOutput) MaxThroughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoscaleSettings) *int {
		if v == nil {
			return nil
		}
		return v.MaxThroughput
	}).(pulumi.IntPtrOutput)
}

type AutoscaleSettingsResponse struct {
	MaxThroughput *int `pulumi:"maxThroughput"`
}

type AutoscaleSettingsResponseOutput struct{ *pulumi.OutputState }

func (AutoscaleSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoscaleSettingsResponse)(nil)).Elem()
}

func (o AutoscaleSettingsResponseOutput) ToAutoscaleSettingsResponseOutput() AutoscaleSettingsResponseOutput {
	return o
}

func (o AutoscaleSettingsResponseOutput) ToAutoscaleSettingsResponseOutputWithContext(ctx context.Context) AutoscaleSettingsResponseOutput {
	return o
}

func (o AutoscaleSettingsResponseOutput) MaxThroughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoscaleSettingsResponse) *int { return v.MaxThroughput }).(pulumi.IntPtrOutput)
}

type AutoscaleSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (AutoscaleSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoscaleSettingsResponse)(nil)).Elem()
}

func (o AutoscaleSettingsResponsePtrOutput) ToAutoscaleSettingsResponsePtrOutput() AutoscaleSettingsResponsePtrOutput {
	return o
}

func (o AutoscaleSettingsResponsePtrOutput) ToAutoscaleSettingsResponsePtrOutputWithContext(ctx context.Context) AutoscaleSettingsResponsePtrOutput {
	return o
}

func (o AutoscaleSettingsResponsePtrOutput) Elem() AutoscaleSettingsResponseOutput {
	return o.ApplyT(func(v *AutoscaleSettingsResponse) AutoscaleSettingsResponse {
		if v != nil {
			return *v
		}
		var ret AutoscaleSettingsResponse
		return ret
	}).(AutoscaleSettingsResponseOutput)
}

func (o AutoscaleSettingsResponsePtrOutput) MaxThroughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoscaleSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxThroughput
	}).(pulumi.IntPtrOutput)
}

type Capability struct {
	Name *string `pulumi:"name"`
}





type CapabilityInput interface {
	pulumi.Input

	ToCapabilityOutput() CapabilityOutput
	ToCapabilityOutputWithContext(context.Context) CapabilityOutput
}

type CapabilityArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (CapabilityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Capability)(nil)).Elem()
}

func (i CapabilityArgs) ToCapabilityOutput() CapabilityOutput {
	return i.ToCapabilityOutputWithContext(context.Background())
}

func (i CapabilityArgs) ToCapabilityOutputWithContext(ctx context.Context) CapabilityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapabilityOutput)
}





type CapabilityArrayInput interface {
	pulumi.Input

	ToCapabilityArrayOutput() CapabilityArrayOutput
	ToCapabilityArrayOutputWithContext(context.Context) CapabilityArrayOutput
}

type CapabilityArray []CapabilityInput

func (CapabilityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Capability)(nil)).Elem()
}

func (i CapabilityArray) ToCapabilityArrayOutput() CapabilityArrayOutput {
	return i.ToCapabilityArrayOutputWithContext(context.Background())
}

func (i CapabilityArray) ToCapabilityArrayOutputWithContext(ctx context.Context) CapabilityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapabilityArrayOutput)
}

type CapabilityOutput struct{ *pulumi.OutputState }

func (CapabilityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Capability)(nil)).Elem()
}

func (o CapabilityOutput) ToCapabilityOutput() CapabilityOutput {
	return o
}

func (o CapabilityOutput) ToCapabilityOutputWithContext(ctx context.Context) CapabilityOutput {
	return o
}

func (o CapabilityOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Capability) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type CapabilityArrayOutput struct{ *pulumi.OutputState }

func (CapabilityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Capability)(nil)).Elem()
}

func (o CapabilityArrayOutput) ToCapabilityArrayOutput() CapabilityArrayOutput {
	return o
}

func (o CapabilityArrayOutput) ToCapabilityArrayOutputWithContext(ctx context.Context) CapabilityArrayOutput {
	return o
}

func (o CapabilityArrayOutput) Index(i pulumi.IntInput) CapabilityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Capability {
		return vs[0].([]Capability)[vs[1].(int)]
	}).(CapabilityOutput)
}

type CapabilityResponse struct {
	Name *string `pulumi:"name"`
}

type CapabilityResponseOutput struct{ *pulumi.OutputState }

func (CapabilityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CapabilityResponse)(nil)).Elem()
}

func (o CapabilityResponseOutput) ToCapabilityResponseOutput() CapabilityResponseOutput {
	return o
}

func (o CapabilityResponseOutput) ToCapabilityResponseOutputWithContext(ctx context.Context) CapabilityResponseOutput {
	return o
}

func (o CapabilityResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CapabilityResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type CapabilityResponseArrayOutput struct{ *pulumi.OutputState }

func (CapabilityResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CapabilityResponse)(nil)).Elem()
}

func (o CapabilityResponseArrayOutput) ToCapabilityResponseArrayOutput() CapabilityResponseArrayOutput {
	return o
}

func (o CapabilityResponseArrayOutput) ToCapabilityResponseArrayOutputWithContext(ctx context.Context) CapabilityResponseArrayOutput {
	return o
}

func (o CapabilityResponseArrayOutput) Index(i pulumi.IntInput) CapabilityResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CapabilityResponse {
		return vs[0].([]CapabilityResponse)[vs[1].(int)]
	}).(CapabilityResponseOutput)
}

type CassandraKeyspaceGetPropertiesResponseOptions struct {
	AutoscaleSettings *AutoscaleSettingsResponse `pulumi:"autoscaleSettings"`
	Throughput        *int                       `pulumi:"throughput"`
}

type CassandraKeyspaceGetPropertiesResponseOptionsOutput struct{ *pulumi.OutputState }

func (CassandraKeyspaceGetPropertiesResponseOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraKeyspaceGetPropertiesResponseOptions)(nil)).Elem()
}

func (o CassandraKeyspaceGetPropertiesResponseOptionsOutput) ToCassandraKeyspaceGetPropertiesResponseOptionsOutput() CassandraKeyspaceGetPropertiesResponseOptionsOutput {
	return o
}

func (o CassandraKeyspaceGetPropertiesResponseOptionsOutput) ToCassandraKeyspaceGetPropertiesResponseOptionsOutputWithContext(ctx context.Context) CassandraKeyspaceGetPropertiesResponseOptionsOutput {
	return o
}

func (o CassandraKeyspaceGetPropertiesResponseOptionsOutput) AutoscaleSettings() AutoscaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v CassandraKeyspaceGetPropertiesResponseOptions) *AutoscaleSettingsResponse {
		return v.AutoscaleSettings
	}).(AutoscaleSettingsResponsePtrOutput)
}

func (o CassandraKeyspaceGetPropertiesResponseOptionsOutput) Throughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CassandraKeyspaceGetPropertiesResponseOptions) *int { return v.Throughput }).(pulumi.IntPtrOutput)
}

type CassandraKeyspaceGetPropertiesResponseOptionsPtrOutput struct{ *pulumi.OutputState }

func (CassandraKeyspaceGetPropertiesResponseOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraKeyspaceGetPropertiesResponseOptions)(nil)).Elem()
}

func (o CassandraKeyspaceGetPropertiesResponseOptionsPtrOutput) ToCassandraKeyspaceGetPropertiesResponseOptionsPtrOutput() CassandraKeyspaceGetPropertiesResponseOptionsPtrOutput {
	return o
}

func (o CassandraKeyspaceGetPropertiesResponseOptionsPtrOutput) ToCassandraKeyspaceGetPropertiesResponseOptionsPtrOutputWithContext(ctx context.Context) CassandraKeyspaceGetPropertiesResponseOptionsPtrOutput {
	return o
}

func (o CassandraKeyspaceGetPropertiesResponseOptionsPtrOutput) Elem() CassandraKeyspaceGetPropertiesResponseOptionsOutput {
	return o.ApplyT(func(v *CassandraKeyspaceGetPropertiesResponseOptions) CassandraKeyspaceGetPropertiesResponseOptions {
		if v != nil {
			return *v
		}
		var ret CassandraKeyspaceGetPropertiesResponseOptions
		return ret
	}).(CassandraKeyspaceGetPropertiesResponseOptionsOutput)
}

func (o CassandraKeyspaceGetPropertiesResponseOptionsPtrOutput) AutoscaleSettings() AutoscaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v *CassandraKeyspaceGetPropertiesResponseOptions) *AutoscaleSettingsResponse {
		if v == nil {
			return nil
		}
		return v.AutoscaleSettings
	}).(AutoscaleSettingsResponsePtrOutput)
}

func (o CassandraKeyspaceGetPropertiesResponseOptionsPtrOutput) Throughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CassandraKeyspaceGetPropertiesResponseOptions) *int {
		if v == nil {
			return nil
		}
		return v.Throughput
	}).(pulumi.IntPtrOutput)
}

type CassandraKeyspaceGetPropertiesResponseResource struct {
	Etag string  `pulumi:"etag"`
	Id   string  `pulumi:"id"`
	Rid  string  `pulumi:"rid"`
	Ts   float64 `pulumi:"ts"`
}

type CassandraKeyspaceGetPropertiesResponseResourceOutput struct{ *pulumi.OutputState }

func (CassandraKeyspaceGetPropertiesResponseResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraKeyspaceGetPropertiesResponseResource)(nil)).Elem()
}

func (o CassandraKeyspaceGetPropertiesResponseResourceOutput) ToCassandraKeyspaceGetPropertiesResponseResourceOutput() CassandraKeyspaceGetPropertiesResponseResourceOutput {
	return o
}

func (o CassandraKeyspaceGetPropertiesResponseResourceOutput) ToCassandraKeyspaceGetPropertiesResponseResourceOutputWithContext(ctx context.Context) CassandraKeyspaceGetPropertiesResponseResourceOutput {
	return o
}

func (o CassandraKeyspaceGetPropertiesResponseResourceOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v CassandraKeyspaceGetPropertiesResponseResource) string { return v.Etag }).(pulumi.StringOutput)
}

func (o CassandraKeyspaceGetPropertiesResponseResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v CassandraKeyspaceGetPropertiesResponseResource) string { return v.Id }).(pulumi.StringOutput)
}

func (o CassandraKeyspaceGetPropertiesResponseResourceOutput) Rid() pulumi.StringOutput {
	return o.ApplyT(func(v CassandraKeyspaceGetPropertiesResponseResource) string { return v.Rid }).(pulumi.StringOutput)
}

func (o CassandraKeyspaceGetPropertiesResponseResourceOutput) Ts() pulumi.Float64Output {
	return o.ApplyT(func(v CassandraKeyspaceGetPropertiesResponseResource) float64 { return v.Ts }).(pulumi.Float64Output)
}

type CassandraKeyspaceGetPropertiesResponseResourcePtrOutput struct{ *pulumi.OutputState }

func (CassandraKeyspaceGetPropertiesResponseResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraKeyspaceGetPropertiesResponseResource)(nil)).Elem()
}

func (o CassandraKeyspaceGetPropertiesResponseResourcePtrOutput) ToCassandraKeyspaceGetPropertiesResponseResourcePtrOutput() CassandraKeyspaceGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o CassandraKeyspaceGetPropertiesResponseResourcePtrOutput) ToCassandraKeyspaceGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) CassandraKeyspaceGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o CassandraKeyspaceGetPropertiesResponseResourcePtrOutput) Elem() CassandraKeyspaceGetPropertiesResponseResourceOutput {
	return o.ApplyT(func(v *CassandraKeyspaceGetPropertiesResponseResource) CassandraKeyspaceGetPropertiesResponseResource {
		if v != nil {
			return *v
		}
		var ret CassandraKeyspaceGetPropertiesResponseResource
		return ret
	}).(CassandraKeyspaceGetPropertiesResponseResourceOutput)
}

func (o CassandraKeyspaceGetPropertiesResponseResourcePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CassandraKeyspaceGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o CassandraKeyspaceGetPropertiesResponseResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CassandraKeyspaceGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o CassandraKeyspaceGetPropertiesResponseResourcePtrOutput) Rid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CassandraKeyspaceGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Rid
	}).(pulumi.StringPtrOutput)
}

func (o CassandraKeyspaceGetPropertiesResponseResourcePtrOutput) Ts() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *CassandraKeyspaceGetPropertiesResponseResource) *float64 {
		if v == nil {
			return nil
		}
		return &v.Ts
	}).(pulumi.Float64PtrOutput)
}

type CassandraKeyspaceResource struct {
	Id string `pulumi:"id"`
}





type CassandraKeyspaceResourceInput interface {
	pulumi.Input

	ToCassandraKeyspaceResourceOutput() CassandraKeyspaceResourceOutput
	ToCassandraKeyspaceResourceOutputWithContext(context.Context) CassandraKeyspaceResourceOutput
}

type CassandraKeyspaceResourceArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (CassandraKeyspaceResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraKeyspaceResource)(nil)).Elem()
}

func (i CassandraKeyspaceResourceArgs) ToCassandraKeyspaceResourceOutput() CassandraKeyspaceResourceOutput {
	return i.ToCassandraKeyspaceResourceOutputWithContext(context.Background())
}

func (i CassandraKeyspaceResourceArgs) ToCassandraKeyspaceResourceOutputWithContext(ctx context.Context) CassandraKeyspaceResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraKeyspaceResourceOutput)
}

type CassandraKeyspaceResourceOutput struct{ *pulumi.OutputState }

func (CassandraKeyspaceResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraKeyspaceResource)(nil)).Elem()
}

func (o CassandraKeyspaceResourceOutput) ToCassandraKeyspaceResourceOutput() CassandraKeyspaceResourceOutput {
	return o
}

func (o CassandraKeyspaceResourceOutput) ToCassandraKeyspaceResourceOutputWithContext(ctx context.Context) CassandraKeyspaceResourceOutput {
	return o
}

func (o CassandraKeyspaceResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v CassandraKeyspaceResource) string { return v.Id }).(pulumi.StringOutput)
}

type CassandraPartitionKey struct {
	Name *string `pulumi:"name"`
}





type CassandraPartitionKeyInput interface {
	pulumi.Input

	ToCassandraPartitionKeyOutput() CassandraPartitionKeyOutput
	ToCassandraPartitionKeyOutputWithContext(context.Context) CassandraPartitionKeyOutput
}

type CassandraPartitionKeyArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (CassandraPartitionKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraPartitionKey)(nil)).Elem()
}

func (i CassandraPartitionKeyArgs) ToCassandraPartitionKeyOutput() CassandraPartitionKeyOutput {
	return i.ToCassandraPartitionKeyOutputWithContext(context.Background())
}

func (i CassandraPartitionKeyArgs) ToCassandraPartitionKeyOutputWithContext(ctx context.Context) CassandraPartitionKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraPartitionKeyOutput)
}





type CassandraPartitionKeyArrayInput interface {
	pulumi.Input

	ToCassandraPartitionKeyArrayOutput() CassandraPartitionKeyArrayOutput
	ToCassandraPartitionKeyArrayOutputWithContext(context.Context) CassandraPartitionKeyArrayOutput
}

type CassandraPartitionKeyArray []CassandraPartitionKeyInput

func (CassandraPartitionKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CassandraPartitionKey)(nil)).Elem()
}

func (i CassandraPartitionKeyArray) ToCassandraPartitionKeyArrayOutput() CassandraPartitionKeyArrayOutput {
	return i.ToCassandraPartitionKeyArrayOutputWithContext(context.Background())
}

func (i CassandraPartitionKeyArray) ToCassandraPartitionKeyArrayOutputWithContext(ctx context.Context) CassandraPartitionKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraPartitionKeyArrayOutput)
}

type CassandraPartitionKeyOutput struct{ *pulumi.OutputState }

func (CassandraPartitionKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraPartitionKey)(nil)).Elem()
}

func (o CassandraPartitionKeyOutput) ToCassandraPartitionKeyOutput() CassandraPartitionKeyOutput {
	return o
}

func (o CassandraPartitionKeyOutput) ToCassandraPartitionKeyOutputWithContext(ctx context.Context) CassandraPartitionKeyOutput {
	return o
}

func (o CassandraPartitionKeyOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CassandraPartitionKey) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type CassandraPartitionKeyArrayOutput struct{ *pulumi.OutputState }

func (CassandraPartitionKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CassandraPartitionKey)(nil)).Elem()
}

func (o CassandraPartitionKeyArrayOutput) ToCassandraPartitionKeyArrayOutput() CassandraPartitionKeyArrayOutput {
	return o
}

func (o CassandraPartitionKeyArrayOutput) ToCassandraPartitionKeyArrayOutputWithContext(ctx context.Context) CassandraPartitionKeyArrayOutput {
	return o
}

func (o CassandraPartitionKeyArrayOutput) Index(i pulumi.IntInput) CassandraPartitionKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CassandraPartitionKey {
		return vs[0].([]CassandraPartitionKey)[vs[1].(int)]
	}).(CassandraPartitionKeyOutput)
}

type CassandraPartitionKeyResponse struct {
	Name *string `pulumi:"name"`
}

type CassandraPartitionKeyResponseOutput struct{ *pulumi.OutputState }

func (CassandraPartitionKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraPartitionKeyResponse)(nil)).Elem()
}

func (o CassandraPartitionKeyResponseOutput) ToCassandraPartitionKeyResponseOutput() CassandraPartitionKeyResponseOutput {
	return o
}

func (o CassandraPartitionKeyResponseOutput) ToCassandraPartitionKeyResponseOutputWithContext(ctx context.Context) CassandraPartitionKeyResponseOutput {
	return o
}

func (o CassandraPartitionKeyResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CassandraPartitionKeyResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type CassandraPartitionKeyResponseArrayOutput struct{ *pulumi.OutputState }

func (CassandraPartitionKeyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CassandraPartitionKeyResponse)(nil)).Elem()
}

func (o CassandraPartitionKeyResponseArrayOutput) ToCassandraPartitionKeyResponseArrayOutput() CassandraPartitionKeyResponseArrayOutput {
	return o
}

func (o CassandraPartitionKeyResponseArrayOutput) ToCassandraPartitionKeyResponseArrayOutputWithContext(ctx context.Context) CassandraPartitionKeyResponseArrayOutput {
	return o
}

func (o CassandraPartitionKeyResponseArrayOutput) Index(i pulumi.IntInput) CassandraPartitionKeyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CassandraPartitionKeyResponse {
		return vs[0].([]CassandraPartitionKeyResponse)[vs[1].(int)]
	}).(CassandraPartitionKeyResponseOutput)
}

type CassandraSchema struct {
	ClusterKeys   []ClusterKey            `pulumi:"clusterKeys"`
	Columns       []Column                `pulumi:"columns"`
	PartitionKeys []CassandraPartitionKey `pulumi:"partitionKeys"`
}





type CassandraSchemaInput interface {
	pulumi.Input

	ToCassandraSchemaOutput() CassandraSchemaOutput
	ToCassandraSchemaOutputWithContext(context.Context) CassandraSchemaOutput
}

type CassandraSchemaArgs struct {
	ClusterKeys   ClusterKeyArrayInput            `pulumi:"clusterKeys"`
	Columns       ColumnArrayInput                `pulumi:"columns"`
	PartitionKeys CassandraPartitionKeyArrayInput `pulumi:"partitionKeys"`
}

func (CassandraSchemaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraSchema)(nil)).Elem()
}

func (i CassandraSchemaArgs) ToCassandraSchemaOutput() CassandraSchemaOutput {
	return i.ToCassandraSchemaOutputWithContext(context.Background())
}

func (i CassandraSchemaArgs) ToCassandraSchemaOutputWithContext(ctx context.Context) CassandraSchemaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraSchemaOutput)
}

func (i CassandraSchemaArgs) ToCassandraSchemaPtrOutput() CassandraSchemaPtrOutput {
	return i.ToCassandraSchemaPtrOutputWithContext(context.Background())
}

func (i CassandraSchemaArgs) ToCassandraSchemaPtrOutputWithContext(ctx context.Context) CassandraSchemaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraSchemaOutput).ToCassandraSchemaPtrOutputWithContext(ctx)
}









type CassandraSchemaPtrInput interface {
	pulumi.Input

	ToCassandraSchemaPtrOutput() CassandraSchemaPtrOutput
	ToCassandraSchemaPtrOutputWithContext(context.Context) CassandraSchemaPtrOutput
}

type cassandraSchemaPtrType CassandraSchemaArgs

func CassandraSchemaPtr(v *CassandraSchemaArgs) CassandraSchemaPtrInput {
	return (*cassandraSchemaPtrType)(v)
}

func (*cassandraSchemaPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraSchema)(nil)).Elem()
}

func (i *cassandraSchemaPtrType) ToCassandraSchemaPtrOutput() CassandraSchemaPtrOutput {
	return i.ToCassandraSchemaPtrOutputWithContext(context.Background())
}

func (i *cassandraSchemaPtrType) ToCassandraSchemaPtrOutputWithContext(ctx context.Context) CassandraSchemaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraSchemaPtrOutput)
}

type CassandraSchemaOutput struct{ *pulumi.OutputState }

func (CassandraSchemaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraSchema)(nil)).Elem()
}

func (o CassandraSchemaOutput) ToCassandraSchemaOutput() CassandraSchemaOutput {
	return o
}

func (o CassandraSchemaOutput) ToCassandraSchemaOutputWithContext(ctx context.Context) CassandraSchemaOutput {
	return o
}

func (o CassandraSchemaOutput) ToCassandraSchemaPtrOutput() CassandraSchemaPtrOutput {
	return o.ToCassandraSchemaPtrOutputWithContext(context.Background())
}

func (o CassandraSchemaOutput) ToCassandraSchemaPtrOutputWithContext(ctx context.Context) CassandraSchemaPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CassandraSchema) *CassandraSchema {
		return &v
	}).(CassandraSchemaPtrOutput)
}

func (o CassandraSchemaOutput) ClusterKeys() ClusterKeyArrayOutput {
	return o.ApplyT(func(v CassandraSchema) []ClusterKey { return v.ClusterKeys }).(ClusterKeyArrayOutput)
}

func (o CassandraSchemaOutput) Columns() ColumnArrayOutput {
	return o.ApplyT(func(v CassandraSchema) []Column { return v.Columns }).(ColumnArrayOutput)
}

func (o CassandraSchemaOutput) PartitionKeys() CassandraPartitionKeyArrayOutput {
	return o.ApplyT(func(v CassandraSchema) []CassandraPartitionKey { return v.PartitionKeys }).(CassandraPartitionKeyArrayOutput)
}

type CassandraSchemaPtrOutput struct{ *pulumi.OutputState }

func (CassandraSchemaPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraSchema)(nil)).Elem()
}

func (o CassandraSchemaPtrOutput) ToCassandraSchemaPtrOutput() CassandraSchemaPtrOutput {
	return o
}

func (o CassandraSchemaPtrOutput) ToCassandraSchemaPtrOutputWithContext(ctx context.Context) CassandraSchemaPtrOutput {
	return o
}

func (o CassandraSchemaPtrOutput) Elem() CassandraSchemaOutput {
	return o.ApplyT(func(v *CassandraSchema) CassandraSchema {
		if v != nil {
			return *v
		}
		var ret CassandraSchema
		return ret
	}).(CassandraSchemaOutput)
}

func (o CassandraSchemaPtrOutput) ClusterKeys() ClusterKeyArrayOutput {
	return o.ApplyT(func(v *CassandraSchema) []ClusterKey {
		if v == nil {
			return nil
		}
		return v.ClusterKeys
	}).(ClusterKeyArrayOutput)
}

func (o CassandraSchemaPtrOutput) Columns() ColumnArrayOutput {
	return o.ApplyT(func(v *CassandraSchema) []Column {
		if v == nil {
			return nil
		}
		return v.Columns
	}).(ColumnArrayOutput)
}

func (o CassandraSchemaPtrOutput) PartitionKeys() CassandraPartitionKeyArrayOutput {
	return o.ApplyT(func(v *CassandraSchema) []CassandraPartitionKey {
		if v == nil {
			return nil
		}
		return v.PartitionKeys
	}).(CassandraPartitionKeyArrayOutput)
}

type CassandraSchemaResponse struct {
	ClusterKeys   []ClusterKeyResponse            `pulumi:"clusterKeys"`
	Columns       []ColumnResponse                `pulumi:"columns"`
	PartitionKeys []CassandraPartitionKeyResponse `pulumi:"partitionKeys"`
}

type CassandraSchemaResponseOutput struct{ *pulumi.OutputState }

func (CassandraSchemaResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraSchemaResponse)(nil)).Elem()
}

func (o CassandraSchemaResponseOutput) ToCassandraSchemaResponseOutput() CassandraSchemaResponseOutput {
	return o
}

func (o CassandraSchemaResponseOutput) ToCassandraSchemaResponseOutputWithContext(ctx context.Context) CassandraSchemaResponseOutput {
	return o
}

func (o CassandraSchemaResponseOutput) ClusterKeys() ClusterKeyResponseArrayOutput {
	return o.ApplyT(func(v CassandraSchemaResponse) []ClusterKeyResponse { return v.ClusterKeys }).(ClusterKeyResponseArrayOutput)
}

func (o CassandraSchemaResponseOutput) Columns() ColumnResponseArrayOutput {
	return o.ApplyT(func(v CassandraSchemaResponse) []ColumnResponse { return v.Columns }).(ColumnResponseArrayOutput)
}

func (o CassandraSchemaResponseOutput) PartitionKeys() CassandraPartitionKeyResponseArrayOutput {
	return o.ApplyT(func(v CassandraSchemaResponse) []CassandraPartitionKeyResponse { return v.PartitionKeys }).(CassandraPartitionKeyResponseArrayOutput)
}

type CassandraSchemaResponsePtrOutput struct{ *pulumi.OutputState }

func (CassandraSchemaResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraSchemaResponse)(nil)).Elem()
}

func (o CassandraSchemaResponsePtrOutput) ToCassandraSchemaResponsePtrOutput() CassandraSchemaResponsePtrOutput {
	return o
}

func (o CassandraSchemaResponsePtrOutput) ToCassandraSchemaResponsePtrOutputWithContext(ctx context.Context) CassandraSchemaResponsePtrOutput {
	return o
}

func (o CassandraSchemaResponsePtrOutput) Elem() CassandraSchemaResponseOutput {
	return o.ApplyT(func(v *CassandraSchemaResponse) CassandraSchemaResponse {
		if v != nil {
			return *v
		}
		var ret CassandraSchemaResponse
		return ret
	}).(CassandraSchemaResponseOutput)
}

func (o CassandraSchemaResponsePtrOutput) ClusterKeys() ClusterKeyResponseArrayOutput {
	return o.ApplyT(func(v *CassandraSchemaResponse) []ClusterKeyResponse {
		if v == nil {
			return nil
		}
		return v.ClusterKeys
	}).(ClusterKeyResponseArrayOutput)
}

func (o CassandraSchemaResponsePtrOutput) Columns() ColumnResponseArrayOutput {
	return o.ApplyT(func(v *CassandraSchemaResponse) []ColumnResponse {
		if v == nil {
			return nil
		}
		return v.Columns
	}).(ColumnResponseArrayOutput)
}

func (o CassandraSchemaResponsePtrOutput) PartitionKeys() CassandraPartitionKeyResponseArrayOutput {
	return o.ApplyT(func(v *CassandraSchemaResponse) []CassandraPartitionKeyResponse {
		if v == nil {
			return nil
		}
		return v.PartitionKeys
	}).(CassandraPartitionKeyResponseArrayOutput)
}

type CassandraTableGetPropertiesResponseOptions struct {
	AutoscaleSettings *AutoscaleSettingsResponse `pulumi:"autoscaleSettings"`
	Throughput        *int                       `pulumi:"throughput"`
}

type CassandraTableGetPropertiesResponseOptionsOutput struct{ *pulumi.OutputState }

func (CassandraTableGetPropertiesResponseOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraTableGetPropertiesResponseOptions)(nil)).Elem()
}

func (o CassandraTableGetPropertiesResponseOptionsOutput) ToCassandraTableGetPropertiesResponseOptionsOutput() CassandraTableGetPropertiesResponseOptionsOutput {
	return o
}

func (o CassandraTableGetPropertiesResponseOptionsOutput) ToCassandraTableGetPropertiesResponseOptionsOutputWithContext(ctx context.Context) CassandraTableGetPropertiesResponseOptionsOutput {
	return o
}

func (o CassandraTableGetPropertiesResponseOptionsOutput) AutoscaleSettings() AutoscaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v CassandraTableGetPropertiesResponseOptions) *AutoscaleSettingsResponse {
		return v.AutoscaleSettings
	}).(AutoscaleSettingsResponsePtrOutput)
}

func (o CassandraTableGetPropertiesResponseOptionsOutput) Throughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CassandraTableGetPropertiesResponseOptions) *int { return v.Throughput }).(pulumi.IntPtrOutput)
}

type CassandraTableGetPropertiesResponseOptionsPtrOutput struct{ *pulumi.OutputState }

func (CassandraTableGetPropertiesResponseOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraTableGetPropertiesResponseOptions)(nil)).Elem()
}

func (o CassandraTableGetPropertiesResponseOptionsPtrOutput) ToCassandraTableGetPropertiesResponseOptionsPtrOutput() CassandraTableGetPropertiesResponseOptionsPtrOutput {
	return o
}

func (o CassandraTableGetPropertiesResponseOptionsPtrOutput) ToCassandraTableGetPropertiesResponseOptionsPtrOutputWithContext(ctx context.Context) CassandraTableGetPropertiesResponseOptionsPtrOutput {
	return o
}

func (o CassandraTableGetPropertiesResponseOptionsPtrOutput) Elem() CassandraTableGetPropertiesResponseOptionsOutput {
	return o.ApplyT(func(v *CassandraTableGetPropertiesResponseOptions) CassandraTableGetPropertiesResponseOptions {
		if v != nil {
			return *v
		}
		var ret CassandraTableGetPropertiesResponseOptions
		return ret
	}).(CassandraTableGetPropertiesResponseOptionsOutput)
}

func (o CassandraTableGetPropertiesResponseOptionsPtrOutput) AutoscaleSettings() AutoscaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v *CassandraTableGetPropertiesResponseOptions) *AutoscaleSettingsResponse {
		if v == nil {
			return nil
		}
		return v.AutoscaleSettings
	}).(AutoscaleSettingsResponsePtrOutput)
}

func (o CassandraTableGetPropertiesResponseOptionsPtrOutput) Throughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CassandraTableGetPropertiesResponseOptions) *int {
		if v == nil {
			return nil
		}
		return v.Throughput
	}).(pulumi.IntPtrOutput)
}

type CassandraTableGetPropertiesResponseResource struct {
	AnalyticalStorageTtl *int                     `pulumi:"analyticalStorageTtl"`
	DefaultTtl           *int                     `pulumi:"defaultTtl"`
	Etag                 string                   `pulumi:"etag"`
	Id                   string                   `pulumi:"id"`
	Rid                  string                   `pulumi:"rid"`
	Schema               *CassandraSchemaResponse `pulumi:"schema"`
	Ts                   float64                  `pulumi:"ts"`
}

type CassandraTableGetPropertiesResponseResourceOutput struct{ *pulumi.OutputState }

func (CassandraTableGetPropertiesResponseResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraTableGetPropertiesResponseResource)(nil)).Elem()
}

func (o CassandraTableGetPropertiesResponseResourceOutput) ToCassandraTableGetPropertiesResponseResourceOutput() CassandraTableGetPropertiesResponseResourceOutput {
	return o
}

func (o CassandraTableGetPropertiesResponseResourceOutput) ToCassandraTableGetPropertiesResponseResourceOutputWithContext(ctx context.Context) CassandraTableGetPropertiesResponseResourceOutput {
	return o
}

func (o CassandraTableGetPropertiesResponseResourceOutput) AnalyticalStorageTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CassandraTableGetPropertiesResponseResource) *int { return v.AnalyticalStorageTtl }).(pulumi.IntPtrOutput)
}

func (o CassandraTableGetPropertiesResponseResourceOutput) DefaultTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CassandraTableGetPropertiesResponseResource) *int { return v.DefaultTtl }).(pulumi.IntPtrOutput)
}

func (o CassandraTableGetPropertiesResponseResourceOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v CassandraTableGetPropertiesResponseResource) string { return v.Etag }).(pulumi.StringOutput)
}

func (o CassandraTableGetPropertiesResponseResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v CassandraTableGetPropertiesResponseResource) string { return v.Id }).(pulumi.StringOutput)
}

func (o CassandraTableGetPropertiesResponseResourceOutput) Rid() pulumi.StringOutput {
	return o.ApplyT(func(v CassandraTableGetPropertiesResponseResource) string { return v.Rid }).(pulumi.StringOutput)
}

func (o CassandraTableGetPropertiesResponseResourceOutput) Schema() CassandraSchemaResponsePtrOutput {
	return o.ApplyT(func(v CassandraTableGetPropertiesResponseResource) *CassandraSchemaResponse { return v.Schema }).(CassandraSchemaResponsePtrOutput)
}

func (o CassandraTableGetPropertiesResponseResourceOutput) Ts() pulumi.Float64Output {
	return o.ApplyT(func(v CassandraTableGetPropertiesResponseResource) float64 { return v.Ts }).(pulumi.Float64Output)
}

type CassandraTableGetPropertiesResponseResourcePtrOutput struct{ *pulumi.OutputState }

func (CassandraTableGetPropertiesResponseResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraTableGetPropertiesResponseResource)(nil)).Elem()
}

func (o CassandraTableGetPropertiesResponseResourcePtrOutput) ToCassandraTableGetPropertiesResponseResourcePtrOutput() CassandraTableGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o CassandraTableGetPropertiesResponseResourcePtrOutput) ToCassandraTableGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) CassandraTableGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o CassandraTableGetPropertiesResponseResourcePtrOutput) Elem() CassandraTableGetPropertiesResponseResourceOutput {
	return o.ApplyT(func(v *CassandraTableGetPropertiesResponseResource) CassandraTableGetPropertiesResponseResource {
		if v != nil {
			return *v
		}
		var ret CassandraTableGetPropertiesResponseResource
		return ret
	}).(CassandraTableGetPropertiesResponseResourceOutput)
}

func (o CassandraTableGetPropertiesResponseResourcePtrOutput) AnalyticalStorageTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CassandraTableGetPropertiesResponseResource) *int {
		if v == nil {
			return nil
		}
		return v.AnalyticalStorageTtl
	}).(pulumi.IntPtrOutput)
}

func (o CassandraTableGetPropertiesResponseResourcePtrOutput) DefaultTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CassandraTableGetPropertiesResponseResource) *int {
		if v == nil {
			return nil
		}
		return v.DefaultTtl
	}).(pulumi.IntPtrOutput)
}

func (o CassandraTableGetPropertiesResponseResourcePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CassandraTableGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o CassandraTableGetPropertiesResponseResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CassandraTableGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o CassandraTableGetPropertiesResponseResourcePtrOutput) Rid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CassandraTableGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Rid
	}).(pulumi.StringPtrOutput)
}

func (o CassandraTableGetPropertiesResponseResourcePtrOutput) Schema() CassandraSchemaResponsePtrOutput {
	return o.ApplyT(func(v *CassandraTableGetPropertiesResponseResource) *CassandraSchemaResponse {
		if v == nil {
			return nil
		}
		return v.Schema
	}).(CassandraSchemaResponsePtrOutput)
}

func (o CassandraTableGetPropertiesResponseResourcePtrOutput) Ts() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *CassandraTableGetPropertiesResponseResource) *float64 {
		if v == nil {
			return nil
		}
		return &v.Ts
	}).(pulumi.Float64PtrOutput)
}

type CassandraTableResource struct {
	AnalyticalStorageTtl *int             `pulumi:"analyticalStorageTtl"`
	DefaultTtl           *int             `pulumi:"defaultTtl"`
	Id                   string           `pulumi:"id"`
	Schema               *CassandraSchema `pulumi:"schema"`
}





type CassandraTableResourceInput interface {
	pulumi.Input

	ToCassandraTableResourceOutput() CassandraTableResourceOutput
	ToCassandraTableResourceOutputWithContext(context.Context) CassandraTableResourceOutput
}

type CassandraTableResourceArgs struct {
	AnalyticalStorageTtl pulumi.IntPtrInput      `pulumi:"analyticalStorageTtl"`
	DefaultTtl           pulumi.IntPtrInput      `pulumi:"defaultTtl"`
	Id                   pulumi.StringInput      `pulumi:"id"`
	Schema               CassandraSchemaPtrInput `pulumi:"schema"`
}

func (CassandraTableResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraTableResource)(nil)).Elem()
}

func (i CassandraTableResourceArgs) ToCassandraTableResourceOutput() CassandraTableResourceOutput {
	return i.ToCassandraTableResourceOutputWithContext(context.Background())
}

func (i CassandraTableResourceArgs) ToCassandraTableResourceOutputWithContext(ctx context.Context) CassandraTableResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraTableResourceOutput)
}

type CassandraTableResourceOutput struct{ *pulumi.OutputState }

func (CassandraTableResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraTableResource)(nil)).Elem()
}

func (o CassandraTableResourceOutput) ToCassandraTableResourceOutput() CassandraTableResourceOutput {
	return o
}

func (o CassandraTableResourceOutput) ToCassandraTableResourceOutputWithContext(ctx context.Context) CassandraTableResourceOutput {
	return o
}

func (o CassandraTableResourceOutput) AnalyticalStorageTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CassandraTableResource) *int { return v.AnalyticalStorageTtl }).(pulumi.IntPtrOutput)
}

func (o CassandraTableResourceOutput) DefaultTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CassandraTableResource) *int { return v.DefaultTtl }).(pulumi.IntPtrOutput)
}

func (o CassandraTableResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v CassandraTableResource) string { return v.Id }).(pulumi.StringOutput)
}

func (o CassandraTableResourceOutput) Schema() CassandraSchemaPtrOutput {
	return o.ApplyT(func(v CassandraTableResource) *CassandraSchema { return v.Schema }).(CassandraSchemaPtrOutput)
}

type ClusterKey struct {
	Name    *string `pulumi:"name"`
	OrderBy *string `pulumi:"orderBy"`
}





type ClusterKeyInput interface {
	pulumi.Input

	ToClusterKeyOutput() ClusterKeyOutput
	ToClusterKeyOutputWithContext(context.Context) ClusterKeyOutput
}

type ClusterKeyArgs struct {
	Name    pulumi.StringPtrInput `pulumi:"name"`
	OrderBy pulumi.StringPtrInput `pulumi:"orderBy"`
}

func (ClusterKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterKey)(nil)).Elem()
}

func (i ClusterKeyArgs) ToClusterKeyOutput() ClusterKeyOutput {
	return i.ToClusterKeyOutputWithContext(context.Background())
}

func (i ClusterKeyArgs) ToClusterKeyOutputWithContext(ctx context.Context) ClusterKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterKeyOutput)
}





type ClusterKeyArrayInput interface {
	pulumi.Input

	ToClusterKeyArrayOutput() ClusterKeyArrayOutput
	ToClusterKeyArrayOutputWithContext(context.Context) ClusterKeyArrayOutput
}

type ClusterKeyArray []ClusterKeyInput

func (ClusterKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterKey)(nil)).Elem()
}

func (i ClusterKeyArray) ToClusterKeyArrayOutput() ClusterKeyArrayOutput {
	return i.ToClusterKeyArrayOutputWithContext(context.Background())
}

func (i ClusterKeyArray) ToClusterKeyArrayOutputWithContext(ctx context.Context) ClusterKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterKeyArrayOutput)
}

type ClusterKeyOutput struct{ *pulumi.OutputState }

func (ClusterKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterKey)(nil)).Elem()
}

func (o ClusterKeyOutput) ToClusterKeyOutput() ClusterKeyOutput {
	return o
}

func (o ClusterKeyOutput) ToClusterKeyOutputWithContext(ctx context.Context) ClusterKeyOutput {
	return o
}

func (o ClusterKeyOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterKey) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ClusterKeyOutput) OrderBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterKey) *string { return v.OrderBy }).(pulumi.StringPtrOutput)
}

type ClusterKeyArrayOutput struct{ *pulumi.OutputState }

func (ClusterKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterKey)(nil)).Elem()
}

func (o ClusterKeyArrayOutput) ToClusterKeyArrayOutput() ClusterKeyArrayOutput {
	return o
}

func (o ClusterKeyArrayOutput) ToClusterKeyArrayOutputWithContext(ctx context.Context) ClusterKeyArrayOutput {
	return o
}

func (o ClusterKeyArrayOutput) Index(i pulumi.IntInput) ClusterKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClusterKey {
		return vs[0].([]ClusterKey)[vs[1].(int)]
	}).(ClusterKeyOutput)
}

type ClusterKeyResponse struct {
	Name    *string `pulumi:"name"`
	OrderBy *string `pulumi:"orderBy"`
}

type ClusterKeyResponseOutput struct{ *pulumi.OutputState }

func (ClusterKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterKeyResponse)(nil)).Elem()
}

func (o ClusterKeyResponseOutput) ToClusterKeyResponseOutput() ClusterKeyResponseOutput {
	return o
}

func (o ClusterKeyResponseOutput) ToClusterKeyResponseOutputWithContext(ctx context.Context) ClusterKeyResponseOutput {
	return o
}

func (o ClusterKeyResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterKeyResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ClusterKeyResponseOutput) OrderBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterKeyResponse) *string { return v.OrderBy }).(pulumi.StringPtrOutput)
}

type ClusterKeyResponseArrayOutput struct{ *pulumi.OutputState }

func (ClusterKeyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterKeyResponse)(nil)).Elem()
}

func (o ClusterKeyResponseArrayOutput) ToClusterKeyResponseArrayOutput() ClusterKeyResponseArrayOutput {
	return o
}

func (o ClusterKeyResponseArrayOutput) ToClusterKeyResponseArrayOutputWithContext(ctx context.Context) ClusterKeyResponseArrayOutput {
	return o
}

func (o ClusterKeyResponseArrayOutput) Index(i pulumi.IntInput) ClusterKeyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClusterKeyResponse {
		return vs[0].([]ClusterKeyResponse)[vs[1].(int)]
	}).(ClusterKeyResponseOutput)
}

type Column struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}





type ColumnInput interface {
	pulumi.Input

	ToColumnOutput() ColumnOutput
	ToColumnOutputWithContext(context.Context) ColumnOutput
}

type ColumnArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (ColumnArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Column)(nil)).Elem()
}

func (i ColumnArgs) ToColumnOutput() ColumnOutput {
	return i.ToColumnOutputWithContext(context.Background())
}

func (i ColumnArgs) ToColumnOutputWithContext(ctx context.Context) ColumnOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ColumnOutput)
}





type ColumnArrayInput interface {
	pulumi.Input

	ToColumnArrayOutput() ColumnArrayOutput
	ToColumnArrayOutputWithContext(context.Context) ColumnArrayOutput
}

type ColumnArray []ColumnInput

func (ColumnArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Column)(nil)).Elem()
}

func (i ColumnArray) ToColumnArrayOutput() ColumnArrayOutput {
	return i.ToColumnArrayOutputWithContext(context.Background())
}

func (i ColumnArray) ToColumnArrayOutputWithContext(ctx context.Context) ColumnArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ColumnArrayOutput)
}

type ColumnOutput struct{ *pulumi.OutputState }

func (ColumnOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Column)(nil)).Elem()
}

func (o ColumnOutput) ToColumnOutput() ColumnOutput {
	return o
}

func (o ColumnOutput) ToColumnOutputWithContext(ctx context.Context) ColumnOutput {
	return o
}

func (o ColumnOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Column) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ColumnOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Column) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ColumnArrayOutput struct{ *pulumi.OutputState }

func (ColumnArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Column)(nil)).Elem()
}

func (o ColumnArrayOutput) ToColumnArrayOutput() ColumnArrayOutput {
	return o
}

func (o ColumnArrayOutput) ToColumnArrayOutputWithContext(ctx context.Context) ColumnArrayOutput {
	return o
}

func (o ColumnArrayOutput) Index(i pulumi.IntInput) ColumnOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Column {
		return vs[0].([]Column)[vs[1].(int)]
	}).(ColumnOutput)
}

type ColumnResponse struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}

type ColumnResponseOutput struct{ *pulumi.OutputState }

func (ColumnResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ColumnResponse)(nil)).Elem()
}

func (o ColumnResponseOutput) ToColumnResponseOutput() ColumnResponseOutput {
	return o
}

func (o ColumnResponseOutput) ToColumnResponseOutputWithContext(ctx context.Context) ColumnResponseOutput {
	return o
}

func (o ColumnResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ColumnResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ColumnResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ColumnResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ColumnResponseArrayOutput struct{ *pulumi.OutputState }

func (ColumnResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ColumnResponse)(nil)).Elem()
}

func (o ColumnResponseArrayOutput) ToColumnResponseArrayOutput() ColumnResponseArrayOutput {
	return o
}

func (o ColumnResponseArrayOutput) ToColumnResponseArrayOutputWithContext(ctx context.Context) ColumnResponseArrayOutput {
	return o
}

func (o ColumnResponseArrayOutput) Index(i pulumi.IntInput) ColumnResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ColumnResponse {
		return vs[0].([]ColumnResponse)[vs[1].(int)]
	}).(ColumnResponseOutput)
}

type CompositePath struct {
	Order *string `pulumi:"order"`
	Path  *string `pulumi:"path"`
}





type CompositePathInput interface {
	pulumi.Input

	ToCompositePathOutput() CompositePathOutput
	ToCompositePathOutputWithContext(context.Context) CompositePathOutput
}

type CompositePathArgs struct {
	Order pulumi.StringPtrInput `pulumi:"order"`
	Path  pulumi.StringPtrInput `pulumi:"path"`
}

func (CompositePathArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CompositePath)(nil)).Elem()
}

func (i CompositePathArgs) ToCompositePathOutput() CompositePathOutput {
	return i.ToCompositePathOutputWithContext(context.Background())
}

func (i CompositePathArgs) ToCompositePathOutputWithContext(ctx context.Context) CompositePathOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CompositePathOutput)
}





type CompositePathArrayInput interface {
	pulumi.Input

	ToCompositePathArrayOutput() CompositePathArrayOutput
	ToCompositePathArrayOutputWithContext(context.Context) CompositePathArrayOutput
}

type CompositePathArray []CompositePathInput

func (CompositePathArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CompositePath)(nil)).Elem()
}

func (i CompositePathArray) ToCompositePathArrayOutput() CompositePathArrayOutput {
	return i.ToCompositePathArrayOutputWithContext(context.Background())
}

func (i CompositePathArray) ToCompositePathArrayOutputWithContext(ctx context.Context) CompositePathArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CompositePathArrayOutput)
}

type CompositePathOutput struct{ *pulumi.OutputState }

func (CompositePathOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CompositePath)(nil)).Elem()
}

func (o CompositePathOutput) ToCompositePathOutput() CompositePathOutput {
	return o
}

func (o CompositePathOutput) ToCompositePathOutputWithContext(ctx context.Context) CompositePathOutput {
	return o
}

func (o CompositePathOutput) Order() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CompositePath) *string { return v.Order }).(pulumi.StringPtrOutput)
}

func (o CompositePathOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CompositePath) *string { return v.Path }).(pulumi.StringPtrOutput)
}

type CompositePathArrayOutput struct{ *pulumi.OutputState }

func (CompositePathArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CompositePath)(nil)).Elem()
}

func (o CompositePathArrayOutput) ToCompositePathArrayOutput() CompositePathArrayOutput {
	return o
}

func (o CompositePathArrayOutput) ToCompositePathArrayOutputWithContext(ctx context.Context) CompositePathArrayOutput {
	return o
}

func (o CompositePathArrayOutput) Index(i pulumi.IntInput) CompositePathOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CompositePath {
		return vs[0].([]CompositePath)[vs[1].(int)]
	}).(CompositePathOutput)
}

type CompositePathResponse struct {
	Order *string `pulumi:"order"`
	Path  *string `pulumi:"path"`
}

type CompositePathResponseOutput struct{ *pulumi.OutputState }

func (CompositePathResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CompositePathResponse)(nil)).Elem()
}

func (o CompositePathResponseOutput) ToCompositePathResponseOutput() CompositePathResponseOutput {
	return o
}

func (o CompositePathResponseOutput) ToCompositePathResponseOutputWithContext(ctx context.Context) CompositePathResponseOutput {
	return o
}

func (o CompositePathResponseOutput) Order() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CompositePathResponse) *string { return v.Order }).(pulumi.StringPtrOutput)
}

func (o CompositePathResponseOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CompositePathResponse) *string { return v.Path }).(pulumi.StringPtrOutput)
}

type CompositePathResponseArrayOutput struct{ *pulumi.OutputState }

func (CompositePathResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CompositePathResponse)(nil)).Elem()
}

func (o CompositePathResponseArrayOutput) ToCompositePathResponseArrayOutput() CompositePathResponseArrayOutput {
	return o
}

func (o CompositePathResponseArrayOutput) ToCompositePathResponseArrayOutputWithContext(ctx context.Context) CompositePathResponseArrayOutput {
	return o
}

func (o CompositePathResponseArrayOutput) Index(i pulumi.IntInput) CompositePathResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CompositePathResponse {
		return vs[0].([]CompositePathResponse)[vs[1].(int)]
	}).(CompositePathResponseOutput)
}

type ConflictResolutionPolicy struct {
	ConflictResolutionPath      *string `pulumi:"conflictResolutionPath"`
	ConflictResolutionProcedure *string `pulumi:"conflictResolutionProcedure"`
	Mode                        *string `pulumi:"mode"`
}


func (val *ConflictResolutionPolicy) Defaults() *ConflictResolutionPolicy {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Mode) {
		mode_ := "LastWriterWins"
		tmp.Mode = &mode_
	}
	return &tmp
}





type ConflictResolutionPolicyInput interface {
	pulumi.Input

	ToConflictResolutionPolicyOutput() ConflictResolutionPolicyOutput
	ToConflictResolutionPolicyOutputWithContext(context.Context) ConflictResolutionPolicyOutput
}

type ConflictResolutionPolicyArgs struct {
	ConflictResolutionPath      pulumi.StringPtrInput `pulumi:"conflictResolutionPath"`
	ConflictResolutionProcedure pulumi.StringPtrInput `pulumi:"conflictResolutionProcedure"`
	Mode                        pulumi.StringPtrInput `pulumi:"mode"`
}

func (ConflictResolutionPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConflictResolutionPolicy)(nil)).Elem()
}

func (i ConflictResolutionPolicyArgs) ToConflictResolutionPolicyOutput() ConflictResolutionPolicyOutput {
	return i.ToConflictResolutionPolicyOutputWithContext(context.Background())
}

func (i ConflictResolutionPolicyArgs) ToConflictResolutionPolicyOutputWithContext(ctx context.Context) ConflictResolutionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConflictResolutionPolicyOutput)
}

func (i ConflictResolutionPolicyArgs) ToConflictResolutionPolicyPtrOutput() ConflictResolutionPolicyPtrOutput {
	return i.ToConflictResolutionPolicyPtrOutputWithContext(context.Background())
}

func (i ConflictResolutionPolicyArgs) ToConflictResolutionPolicyPtrOutputWithContext(ctx context.Context) ConflictResolutionPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConflictResolutionPolicyOutput).ToConflictResolutionPolicyPtrOutputWithContext(ctx)
}









type ConflictResolutionPolicyPtrInput interface {
	pulumi.Input

	ToConflictResolutionPolicyPtrOutput() ConflictResolutionPolicyPtrOutput
	ToConflictResolutionPolicyPtrOutputWithContext(context.Context) ConflictResolutionPolicyPtrOutput
}

type conflictResolutionPolicyPtrType ConflictResolutionPolicyArgs

func ConflictResolutionPolicyPtr(v *ConflictResolutionPolicyArgs) ConflictResolutionPolicyPtrInput {
	return (*conflictResolutionPolicyPtrType)(v)
}

func (*conflictResolutionPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConflictResolutionPolicy)(nil)).Elem()
}

func (i *conflictResolutionPolicyPtrType) ToConflictResolutionPolicyPtrOutput() ConflictResolutionPolicyPtrOutput {
	return i.ToConflictResolutionPolicyPtrOutputWithContext(context.Background())
}

func (i *conflictResolutionPolicyPtrType) ToConflictResolutionPolicyPtrOutputWithContext(ctx context.Context) ConflictResolutionPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConflictResolutionPolicyPtrOutput)
}

type ConflictResolutionPolicyOutput struct{ *pulumi.OutputState }

func (ConflictResolutionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConflictResolutionPolicy)(nil)).Elem()
}

func (o ConflictResolutionPolicyOutput) ToConflictResolutionPolicyOutput() ConflictResolutionPolicyOutput {
	return o
}

func (o ConflictResolutionPolicyOutput) ToConflictResolutionPolicyOutputWithContext(ctx context.Context) ConflictResolutionPolicyOutput {
	return o
}

func (o ConflictResolutionPolicyOutput) ToConflictResolutionPolicyPtrOutput() ConflictResolutionPolicyPtrOutput {
	return o.ToConflictResolutionPolicyPtrOutputWithContext(context.Background())
}

func (o ConflictResolutionPolicyOutput) ToConflictResolutionPolicyPtrOutputWithContext(ctx context.Context) ConflictResolutionPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConflictResolutionPolicy) *ConflictResolutionPolicy {
		return &v
	}).(ConflictResolutionPolicyPtrOutput)
}

func (o ConflictResolutionPolicyOutput) ConflictResolutionPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConflictResolutionPolicy) *string { return v.ConflictResolutionPath }).(pulumi.StringPtrOutput)
}

func (o ConflictResolutionPolicyOutput) ConflictResolutionProcedure() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConflictResolutionPolicy) *string { return v.ConflictResolutionProcedure }).(pulumi.StringPtrOutput)
}

func (o ConflictResolutionPolicyOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConflictResolutionPolicy) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

type ConflictResolutionPolicyPtrOutput struct{ *pulumi.OutputState }

func (ConflictResolutionPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConflictResolutionPolicy)(nil)).Elem()
}

func (o ConflictResolutionPolicyPtrOutput) ToConflictResolutionPolicyPtrOutput() ConflictResolutionPolicyPtrOutput {
	return o
}

func (o ConflictResolutionPolicyPtrOutput) ToConflictResolutionPolicyPtrOutputWithContext(ctx context.Context) ConflictResolutionPolicyPtrOutput {
	return o
}

func (o ConflictResolutionPolicyPtrOutput) Elem() ConflictResolutionPolicyOutput {
	return o.ApplyT(func(v *ConflictResolutionPolicy) ConflictResolutionPolicy {
		if v != nil {
			return *v
		}
		var ret ConflictResolutionPolicy
		return ret
	}).(ConflictResolutionPolicyOutput)
}

func (o ConflictResolutionPolicyPtrOutput) ConflictResolutionPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConflictResolutionPolicy) *string {
		if v == nil {
			return nil
		}
		return v.ConflictResolutionPath
	}).(pulumi.StringPtrOutput)
}

func (o ConflictResolutionPolicyPtrOutput) ConflictResolutionProcedure() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConflictResolutionPolicy) *string {
		if v == nil {
			return nil
		}
		return v.ConflictResolutionProcedure
	}).(pulumi.StringPtrOutput)
}

func (o ConflictResolutionPolicyPtrOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConflictResolutionPolicy) *string {
		if v == nil {
			return nil
		}
		return v.Mode
	}).(pulumi.StringPtrOutput)
}

type ConflictResolutionPolicyResponse struct {
	ConflictResolutionPath      *string `pulumi:"conflictResolutionPath"`
	ConflictResolutionProcedure *string `pulumi:"conflictResolutionProcedure"`
	Mode                        *string `pulumi:"mode"`
}


func (val *ConflictResolutionPolicyResponse) Defaults() *ConflictResolutionPolicyResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Mode) {
		mode_ := "LastWriterWins"
		tmp.Mode = &mode_
	}
	return &tmp
}

type ConflictResolutionPolicyResponseOutput struct{ *pulumi.OutputState }

func (ConflictResolutionPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConflictResolutionPolicyResponse)(nil)).Elem()
}

func (o ConflictResolutionPolicyResponseOutput) ToConflictResolutionPolicyResponseOutput() ConflictResolutionPolicyResponseOutput {
	return o
}

func (o ConflictResolutionPolicyResponseOutput) ToConflictResolutionPolicyResponseOutputWithContext(ctx context.Context) ConflictResolutionPolicyResponseOutput {
	return o
}

func (o ConflictResolutionPolicyResponseOutput) ConflictResolutionPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConflictResolutionPolicyResponse) *string { return v.ConflictResolutionPath }).(pulumi.StringPtrOutput)
}

func (o ConflictResolutionPolicyResponseOutput) ConflictResolutionProcedure() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConflictResolutionPolicyResponse) *string { return v.ConflictResolutionProcedure }).(pulumi.StringPtrOutput)
}

func (o ConflictResolutionPolicyResponseOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConflictResolutionPolicyResponse) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

type ConflictResolutionPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (ConflictResolutionPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConflictResolutionPolicyResponse)(nil)).Elem()
}

func (o ConflictResolutionPolicyResponsePtrOutput) ToConflictResolutionPolicyResponsePtrOutput() ConflictResolutionPolicyResponsePtrOutput {
	return o
}

func (o ConflictResolutionPolicyResponsePtrOutput) ToConflictResolutionPolicyResponsePtrOutputWithContext(ctx context.Context) ConflictResolutionPolicyResponsePtrOutput {
	return o
}

func (o ConflictResolutionPolicyResponsePtrOutput) Elem() ConflictResolutionPolicyResponseOutput {
	return o.ApplyT(func(v *ConflictResolutionPolicyResponse) ConflictResolutionPolicyResponse {
		if v != nil {
			return *v
		}
		var ret ConflictResolutionPolicyResponse
		return ret
	}).(ConflictResolutionPolicyResponseOutput)
}

func (o ConflictResolutionPolicyResponsePtrOutput) ConflictResolutionPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConflictResolutionPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.ConflictResolutionPath
	}).(pulumi.StringPtrOutput)
}

func (o ConflictResolutionPolicyResponsePtrOutput) ConflictResolutionProcedure() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConflictResolutionPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.ConflictResolutionProcedure
	}).(pulumi.StringPtrOutput)
}

func (o ConflictResolutionPolicyResponsePtrOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConflictResolutionPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Mode
	}).(pulumi.StringPtrOutput)
}

type ConsistencyPolicy struct {
	DefaultConsistencyLevel DefaultConsistencyLevel `pulumi:"defaultConsistencyLevel"`
	MaxIntervalInSeconds    *int                    `pulumi:"maxIntervalInSeconds"`
	MaxStalenessPrefix      *float64                `pulumi:"maxStalenessPrefix"`
}





type ConsistencyPolicyInput interface {
	pulumi.Input

	ToConsistencyPolicyOutput() ConsistencyPolicyOutput
	ToConsistencyPolicyOutputWithContext(context.Context) ConsistencyPolicyOutput
}

type ConsistencyPolicyArgs struct {
	DefaultConsistencyLevel DefaultConsistencyLevelInput `pulumi:"defaultConsistencyLevel"`
	MaxIntervalInSeconds    pulumi.IntPtrInput           `pulumi:"maxIntervalInSeconds"`
	MaxStalenessPrefix      pulumi.Float64PtrInput       `pulumi:"maxStalenessPrefix"`
}

func (ConsistencyPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConsistencyPolicy)(nil)).Elem()
}

func (i ConsistencyPolicyArgs) ToConsistencyPolicyOutput() ConsistencyPolicyOutput {
	return i.ToConsistencyPolicyOutputWithContext(context.Background())
}

func (i ConsistencyPolicyArgs) ToConsistencyPolicyOutputWithContext(ctx context.Context) ConsistencyPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsistencyPolicyOutput)
}

func (i ConsistencyPolicyArgs) ToConsistencyPolicyPtrOutput() ConsistencyPolicyPtrOutput {
	return i.ToConsistencyPolicyPtrOutputWithContext(context.Background())
}

func (i ConsistencyPolicyArgs) ToConsistencyPolicyPtrOutputWithContext(ctx context.Context) ConsistencyPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsistencyPolicyOutput).ToConsistencyPolicyPtrOutputWithContext(ctx)
}









type ConsistencyPolicyPtrInput interface {
	pulumi.Input

	ToConsistencyPolicyPtrOutput() ConsistencyPolicyPtrOutput
	ToConsistencyPolicyPtrOutputWithContext(context.Context) ConsistencyPolicyPtrOutput
}

type consistencyPolicyPtrType ConsistencyPolicyArgs

func ConsistencyPolicyPtr(v *ConsistencyPolicyArgs) ConsistencyPolicyPtrInput {
	return (*consistencyPolicyPtrType)(v)
}

func (*consistencyPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConsistencyPolicy)(nil)).Elem()
}

func (i *consistencyPolicyPtrType) ToConsistencyPolicyPtrOutput() ConsistencyPolicyPtrOutput {
	return i.ToConsistencyPolicyPtrOutputWithContext(context.Background())
}

func (i *consistencyPolicyPtrType) ToConsistencyPolicyPtrOutputWithContext(ctx context.Context) ConsistencyPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsistencyPolicyPtrOutput)
}

type ConsistencyPolicyOutput struct{ *pulumi.OutputState }

func (ConsistencyPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConsistencyPolicy)(nil)).Elem()
}

func (o ConsistencyPolicyOutput) ToConsistencyPolicyOutput() ConsistencyPolicyOutput {
	return o
}

func (o ConsistencyPolicyOutput) ToConsistencyPolicyOutputWithContext(ctx context.Context) ConsistencyPolicyOutput {
	return o
}

func (o ConsistencyPolicyOutput) ToConsistencyPolicyPtrOutput() ConsistencyPolicyPtrOutput {
	return o.ToConsistencyPolicyPtrOutputWithContext(context.Background())
}

func (o ConsistencyPolicyOutput) ToConsistencyPolicyPtrOutputWithContext(ctx context.Context) ConsistencyPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConsistencyPolicy) *ConsistencyPolicy {
		return &v
	}).(ConsistencyPolicyPtrOutput)
}

func (o ConsistencyPolicyOutput) DefaultConsistencyLevel() DefaultConsistencyLevelOutput {
	return o.ApplyT(func(v ConsistencyPolicy) DefaultConsistencyLevel { return v.DefaultConsistencyLevel }).(DefaultConsistencyLevelOutput)
}

func (o ConsistencyPolicyOutput) MaxIntervalInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ConsistencyPolicy) *int { return v.MaxIntervalInSeconds }).(pulumi.IntPtrOutput)
}

func (o ConsistencyPolicyOutput) MaxStalenessPrefix() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ConsistencyPolicy) *float64 { return v.MaxStalenessPrefix }).(pulumi.Float64PtrOutput)
}

type ConsistencyPolicyPtrOutput struct{ *pulumi.OutputState }

func (ConsistencyPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConsistencyPolicy)(nil)).Elem()
}

func (o ConsistencyPolicyPtrOutput) ToConsistencyPolicyPtrOutput() ConsistencyPolicyPtrOutput {
	return o
}

func (o ConsistencyPolicyPtrOutput) ToConsistencyPolicyPtrOutputWithContext(ctx context.Context) ConsistencyPolicyPtrOutput {
	return o
}

func (o ConsistencyPolicyPtrOutput) Elem() ConsistencyPolicyOutput {
	return o.ApplyT(func(v *ConsistencyPolicy) ConsistencyPolicy {
		if v != nil {
			return *v
		}
		var ret ConsistencyPolicy
		return ret
	}).(ConsistencyPolicyOutput)
}

func (o ConsistencyPolicyPtrOutput) DefaultConsistencyLevel() DefaultConsistencyLevelPtrOutput {
	return o.ApplyT(func(v *ConsistencyPolicy) *DefaultConsistencyLevel {
		if v == nil {
			return nil
		}
		return &v.DefaultConsistencyLevel
	}).(DefaultConsistencyLevelPtrOutput)
}

func (o ConsistencyPolicyPtrOutput) MaxIntervalInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ConsistencyPolicy) *int {
		if v == nil {
			return nil
		}
		return v.MaxIntervalInSeconds
	}).(pulumi.IntPtrOutput)
}

func (o ConsistencyPolicyPtrOutput) MaxStalenessPrefix() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ConsistencyPolicy) *float64 {
		if v == nil {
			return nil
		}
		return v.MaxStalenessPrefix
	}).(pulumi.Float64PtrOutput)
}

type ConsistencyPolicyResponse struct {
	DefaultConsistencyLevel string   `pulumi:"defaultConsistencyLevel"`
	MaxIntervalInSeconds    *int     `pulumi:"maxIntervalInSeconds"`
	MaxStalenessPrefix      *float64 `pulumi:"maxStalenessPrefix"`
}

type ConsistencyPolicyResponseOutput struct{ *pulumi.OutputState }

func (ConsistencyPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConsistencyPolicyResponse)(nil)).Elem()
}

func (o ConsistencyPolicyResponseOutput) ToConsistencyPolicyResponseOutput() ConsistencyPolicyResponseOutput {
	return o
}

func (o ConsistencyPolicyResponseOutput) ToConsistencyPolicyResponseOutputWithContext(ctx context.Context) ConsistencyPolicyResponseOutput {
	return o
}

func (o ConsistencyPolicyResponseOutput) DefaultConsistencyLevel() pulumi.StringOutput {
	return o.ApplyT(func(v ConsistencyPolicyResponse) string { return v.DefaultConsistencyLevel }).(pulumi.StringOutput)
}

func (o ConsistencyPolicyResponseOutput) MaxIntervalInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ConsistencyPolicyResponse) *int { return v.MaxIntervalInSeconds }).(pulumi.IntPtrOutput)
}

func (o ConsistencyPolicyResponseOutput) MaxStalenessPrefix() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ConsistencyPolicyResponse) *float64 { return v.MaxStalenessPrefix }).(pulumi.Float64PtrOutput)
}

type ConsistencyPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (ConsistencyPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConsistencyPolicyResponse)(nil)).Elem()
}

func (o ConsistencyPolicyResponsePtrOutput) ToConsistencyPolicyResponsePtrOutput() ConsistencyPolicyResponsePtrOutput {
	return o
}

func (o ConsistencyPolicyResponsePtrOutput) ToConsistencyPolicyResponsePtrOutputWithContext(ctx context.Context) ConsistencyPolicyResponsePtrOutput {
	return o
}

func (o ConsistencyPolicyResponsePtrOutput) Elem() ConsistencyPolicyResponseOutput {
	return o.ApplyT(func(v *ConsistencyPolicyResponse) ConsistencyPolicyResponse {
		if v != nil {
			return *v
		}
		var ret ConsistencyPolicyResponse
		return ret
	}).(ConsistencyPolicyResponseOutput)
}

func (o ConsistencyPolicyResponsePtrOutput) DefaultConsistencyLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConsistencyPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DefaultConsistencyLevel
	}).(pulumi.StringPtrOutput)
}

func (o ConsistencyPolicyResponsePtrOutput) MaxIntervalInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ConsistencyPolicyResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxIntervalInSeconds
	}).(pulumi.IntPtrOutput)
}

func (o ConsistencyPolicyResponsePtrOutput) MaxStalenessPrefix() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ConsistencyPolicyResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.MaxStalenessPrefix
	}).(pulumi.Float64PtrOutput)
}

type ContainerPartitionKey struct {
	Kind    *string  `pulumi:"kind"`
	Paths   []string `pulumi:"paths"`
	Version *int     `pulumi:"version"`
}


func (val *ContainerPartitionKey) Defaults() *ContainerPartitionKey {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Kind) {
		kind_ := "Hash"
		tmp.Kind = &kind_
	}
	return &tmp
}





type ContainerPartitionKeyInput interface {
	pulumi.Input

	ToContainerPartitionKeyOutput() ContainerPartitionKeyOutput
	ToContainerPartitionKeyOutputWithContext(context.Context) ContainerPartitionKeyOutput
}

type ContainerPartitionKeyArgs struct {
	Kind    pulumi.StringPtrInput   `pulumi:"kind"`
	Paths   pulumi.StringArrayInput `pulumi:"paths"`
	Version pulumi.IntPtrInput      `pulumi:"version"`
}

func (ContainerPartitionKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerPartitionKey)(nil)).Elem()
}

func (i ContainerPartitionKeyArgs) ToContainerPartitionKeyOutput() ContainerPartitionKeyOutput {
	return i.ToContainerPartitionKeyOutputWithContext(context.Background())
}

func (i ContainerPartitionKeyArgs) ToContainerPartitionKeyOutputWithContext(ctx context.Context) ContainerPartitionKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerPartitionKeyOutput)
}

func (i ContainerPartitionKeyArgs) ToContainerPartitionKeyPtrOutput() ContainerPartitionKeyPtrOutput {
	return i.ToContainerPartitionKeyPtrOutputWithContext(context.Background())
}

func (i ContainerPartitionKeyArgs) ToContainerPartitionKeyPtrOutputWithContext(ctx context.Context) ContainerPartitionKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerPartitionKeyOutput).ToContainerPartitionKeyPtrOutputWithContext(ctx)
}









type ContainerPartitionKeyPtrInput interface {
	pulumi.Input

	ToContainerPartitionKeyPtrOutput() ContainerPartitionKeyPtrOutput
	ToContainerPartitionKeyPtrOutputWithContext(context.Context) ContainerPartitionKeyPtrOutput
}

type containerPartitionKeyPtrType ContainerPartitionKeyArgs

func ContainerPartitionKeyPtr(v *ContainerPartitionKeyArgs) ContainerPartitionKeyPtrInput {
	return (*containerPartitionKeyPtrType)(v)
}

func (*containerPartitionKeyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerPartitionKey)(nil)).Elem()
}

func (i *containerPartitionKeyPtrType) ToContainerPartitionKeyPtrOutput() ContainerPartitionKeyPtrOutput {
	return i.ToContainerPartitionKeyPtrOutputWithContext(context.Background())
}

func (i *containerPartitionKeyPtrType) ToContainerPartitionKeyPtrOutputWithContext(ctx context.Context) ContainerPartitionKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerPartitionKeyPtrOutput)
}

type ContainerPartitionKeyOutput struct{ *pulumi.OutputState }

func (ContainerPartitionKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerPartitionKey)(nil)).Elem()
}

func (o ContainerPartitionKeyOutput) ToContainerPartitionKeyOutput() ContainerPartitionKeyOutput {
	return o
}

func (o ContainerPartitionKeyOutput) ToContainerPartitionKeyOutputWithContext(ctx context.Context) ContainerPartitionKeyOutput {
	return o
}

func (o ContainerPartitionKeyOutput) ToContainerPartitionKeyPtrOutput() ContainerPartitionKeyPtrOutput {
	return o.ToContainerPartitionKeyPtrOutputWithContext(context.Background())
}

func (o ContainerPartitionKeyOutput) ToContainerPartitionKeyPtrOutputWithContext(ctx context.Context) ContainerPartitionKeyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerPartitionKey) *ContainerPartitionKey {
		return &v
	}).(ContainerPartitionKeyPtrOutput)
}

func (o ContainerPartitionKeyOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerPartitionKey) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ContainerPartitionKeyOutput) Paths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ContainerPartitionKey) []string { return v.Paths }).(pulumi.StringArrayOutput)
}

func (o ContainerPartitionKeyOutput) Version() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContainerPartitionKey) *int { return v.Version }).(pulumi.IntPtrOutput)
}

type ContainerPartitionKeyPtrOutput struct{ *pulumi.OutputState }

func (ContainerPartitionKeyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerPartitionKey)(nil)).Elem()
}

func (o ContainerPartitionKeyPtrOutput) ToContainerPartitionKeyPtrOutput() ContainerPartitionKeyPtrOutput {
	return o
}

func (o ContainerPartitionKeyPtrOutput) ToContainerPartitionKeyPtrOutputWithContext(ctx context.Context) ContainerPartitionKeyPtrOutput {
	return o
}

func (o ContainerPartitionKeyPtrOutput) Elem() ContainerPartitionKeyOutput {
	return o.ApplyT(func(v *ContainerPartitionKey) ContainerPartitionKey {
		if v != nil {
			return *v
		}
		var ret ContainerPartitionKey
		return ret
	}).(ContainerPartitionKeyOutput)
}

func (o ContainerPartitionKeyPtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerPartitionKey) *string {
		if v == nil {
			return nil
		}
		return v.Kind
	}).(pulumi.StringPtrOutput)
}

func (o ContainerPartitionKeyPtrOutput) Paths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ContainerPartitionKey) []string {
		if v == nil {
			return nil
		}
		return v.Paths
	}).(pulumi.StringArrayOutput)
}

func (o ContainerPartitionKeyPtrOutput) Version() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ContainerPartitionKey) *int {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.IntPtrOutput)
}

type ContainerPartitionKeyResponse struct {
	Kind    *string  `pulumi:"kind"`
	Paths   []string `pulumi:"paths"`
	Version *int     `pulumi:"version"`
}


func (val *ContainerPartitionKeyResponse) Defaults() *ContainerPartitionKeyResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Kind) {
		kind_ := "Hash"
		tmp.Kind = &kind_
	}
	return &tmp
}

type ContainerPartitionKeyResponseOutput struct{ *pulumi.OutputState }

func (ContainerPartitionKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerPartitionKeyResponse)(nil)).Elem()
}

func (o ContainerPartitionKeyResponseOutput) ToContainerPartitionKeyResponseOutput() ContainerPartitionKeyResponseOutput {
	return o
}

func (o ContainerPartitionKeyResponseOutput) ToContainerPartitionKeyResponseOutputWithContext(ctx context.Context) ContainerPartitionKeyResponseOutput {
	return o
}

func (o ContainerPartitionKeyResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerPartitionKeyResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ContainerPartitionKeyResponseOutput) Paths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ContainerPartitionKeyResponse) []string { return v.Paths }).(pulumi.StringArrayOutput)
}

func (o ContainerPartitionKeyResponseOutput) Version() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContainerPartitionKeyResponse) *int { return v.Version }).(pulumi.IntPtrOutput)
}

type ContainerPartitionKeyResponsePtrOutput struct{ *pulumi.OutputState }

func (ContainerPartitionKeyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerPartitionKeyResponse)(nil)).Elem()
}

func (o ContainerPartitionKeyResponsePtrOutput) ToContainerPartitionKeyResponsePtrOutput() ContainerPartitionKeyResponsePtrOutput {
	return o
}

func (o ContainerPartitionKeyResponsePtrOutput) ToContainerPartitionKeyResponsePtrOutputWithContext(ctx context.Context) ContainerPartitionKeyResponsePtrOutput {
	return o
}

func (o ContainerPartitionKeyResponsePtrOutput) Elem() ContainerPartitionKeyResponseOutput {
	return o.ApplyT(func(v *ContainerPartitionKeyResponse) ContainerPartitionKeyResponse {
		if v != nil {
			return *v
		}
		var ret ContainerPartitionKeyResponse
		return ret
	}).(ContainerPartitionKeyResponseOutput)
}

func (o ContainerPartitionKeyResponsePtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerPartitionKeyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Kind
	}).(pulumi.StringPtrOutput)
}

func (o ContainerPartitionKeyResponsePtrOutput) Paths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ContainerPartitionKeyResponse) []string {
		if v == nil {
			return nil
		}
		return v.Paths
	}).(pulumi.StringArrayOutput)
}

func (o ContainerPartitionKeyResponsePtrOutput) Version() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ContainerPartitionKeyResponse) *int {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.IntPtrOutput)
}

type ContinuousModeBackupPolicy struct {
	Type string `pulumi:"type"`
}

type ContinuousModeBackupPolicyResponse struct {
	Type string `pulumi:"type"`
}

type CorsPolicy struct {
	AllowedHeaders  *string  `pulumi:"allowedHeaders"`
	AllowedMethods  *string  `pulumi:"allowedMethods"`
	AllowedOrigins  string   `pulumi:"allowedOrigins"`
	ExposedHeaders  *string  `pulumi:"exposedHeaders"`
	MaxAgeInSeconds *float64 `pulumi:"maxAgeInSeconds"`
}





type CorsPolicyInput interface {
	pulumi.Input

	ToCorsPolicyOutput() CorsPolicyOutput
	ToCorsPolicyOutputWithContext(context.Context) CorsPolicyOutput
}

type CorsPolicyArgs struct {
	AllowedHeaders  pulumi.StringPtrInput  `pulumi:"allowedHeaders"`
	AllowedMethods  pulumi.StringPtrInput  `pulumi:"allowedMethods"`
	AllowedOrigins  pulumi.StringInput     `pulumi:"allowedOrigins"`
	ExposedHeaders  pulumi.StringPtrInput  `pulumi:"exposedHeaders"`
	MaxAgeInSeconds pulumi.Float64PtrInput `pulumi:"maxAgeInSeconds"`
}

func (CorsPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CorsPolicy)(nil)).Elem()
}

func (i CorsPolicyArgs) ToCorsPolicyOutput() CorsPolicyOutput {
	return i.ToCorsPolicyOutputWithContext(context.Background())
}

func (i CorsPolicyArgs) ToCorsPolicyOutputWithContext(ctx context.Context) CorsPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorsPolicyOutput)
}





type CorsPolicyArrayInput interface {
	pulumi.Input

	ToCorsPolicyArrayOutput() CorsPolicyArrayOutput
	ToCorsPolicyArrayOutputWithContext(context.Context) CorsPolicyArrayOutput
}

type CorsPolicyArray []CorsPolicyInput

func (CorsPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CorsPolicy)(nil)).Elem()
}

func (i CorsPolicyArray) ToCorsPolicyArrayOutput() CorsPolicyArrayOutput {
	return i.ToCorsPolicyArrayOutputWithContext(context.Background())
}

func (i CorsPolicyArray) ToCorsPolicyArrayOutputWithContext(ctx context.Context) CorsPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorsPolicyArrayOutput)
}

type CorsPolicyOutput struct{ *pulumi.OutputState }

func (CorsPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CorsPolicy)(nil)).Elem()
}

func (o CorsPolicyOutput) ToCorsPolicyOutput() CorsPolicyOutput {
	return o
}

func (o CorsPolicyOutput) ToCorsPolicyOutputWithContext(ctx context.Context) CorsPolicyOutput {
	return o
}

func (o CorsPolicyOutput) AllowedHeaders() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorsPolicy) *string { return v.AllowedHeaders }).(pulumi.StringPtrOutput)
}

func (o CorsPolicyOutput) AllowedMethods() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorsPolicy) *string { return v.AllowedMethods }).(pulumi.StringPtrOutput)
}

func (o CorsPolicyOutput) AllowedOrigins() pulumi.StringOutput {
	return o.ApplyT(func(v CorsPolicy) string { return v.AllowedOrigins }).(pulumi.StringOutput)
}

func (o CorsPolicyOutput) ExposedHeaders() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorsPolicy) *string { return v.ExposedHeaders }).(pulumi.StringPtrOutput)
}

func (o CorsPolicyOutput) MaxAgeInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v CorsPolicy) *float64 { return v.MaxAgeInSeconds }).(pulumi.Float64PtrOutput)
}

type CorsPolicyArrayOutput struct{ *pulumi.OutputState }

func (CorsPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CorsPolicy)(nil)).Elem()
}

func (o CorsPolicyArrayOutput) ToCorsPolicyArrayOutput() CorsPolicyArrayOutput {
	return o
}

func (o CorsPolicyArrayOutput) ToCorsPolicyArrayOutputWithContext(ctx context.Context) CorsPolicyArrayOutput {
	return o
}

func (o CorsPolicyArrayOutput) Index(i pulumi.IntInput) CorsPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CorsPolicy {
		return vs[0].([]CorsPolicy)[vs[1].(int)]
	}).(CorsPolicyOutput)
}

type CorsPolicyResponse struct {
	AllowedHeaders  *string  `pulumi:"allowedHeaders"`
	AllowedMethods  *string  `pulumi:"allowedMethods"`
	AllowedOrigins  string   `pulumi:"allowedOrigins"`
	ExposedHeaders  *string  `pulumi:"exposedHeaders"`
	MaxAgeInSeconds *float64 `pulumi:"maxAgeInSeconds"`
}

type CorsPolicyResponseOutput struct{ *pulumi.OutputState }

func (CorsPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CorsPolicyResponse)(nil)).Elem()
}

func (o CorsPolicyResponseOutput) ToCorsPolicyResponseOutput() CorsPolicyResponseOutput {
	return o
}

func (o CorsPolicyResponseOutput) ToCorsPolicyResponseOutputWithContext(ctx context.Context) CorsPolicyResponseOutput {
	return o
}

func (o CorsPolicyResponseOutput) AllowedHeaders() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorsPolicyResponse) *string { return v.AllowedHeaders }).(pulumi.StringPtrOutput)
}

func (o CorsPolicyResponseOutput) AllowedMethods() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorsPolicyResponse) *string { return v.AllowedMethods }).(pulumi.StringPtrOutput)
}

func (o CorsPolicyResponseOutput) AllowedOrigins() pulumi.StringOutput {
	return o.ApplyT(func(v CorsPolicyResponse) string { return v.AllowedOrigins }).(pulumi.StringOutput)
}

func (o CorsPolicyResponseOutput) ExposedHeaders() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorsPolicyResponse) *string { return v.ExposedHeaders }).(pulumi.StringPtrOutput)
}

func (o CorsPolicyResponseOutput) MaxAgeInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v CorsPolicyResponse) *float64 { return v.MaxAgeInSeconds }).(pulumi.Float64PtrOutput)
}

type CorsPolicyResponseArrayOutput struct{ *pulumi.OutputState }

func (CorsPolicyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CorsPolicyResponse)(nil)).Elem()
}

func (o CorsPolicyResponseArrayOutput) ToCorsPolicyResponseArrayOutput() CorsPolicyResponseArrayOutput {
	return o
}

func (o CorsPolicyResponseArrayOutput) ToCorsPolicyResponseArrayOutputWithContext(ctx context.Context) CorsPolicyResponseArrayOutput {
	return o
}

func (o CorsPolicyResponseArrayOutput) Index(i pulumi.IntInput) CorsPolicyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CorsPolicyResponse {
		return vs[0].([]CorsPolicyResponse)[vs[1].(int)]
	}).(CorsPolicyResponseOutput)
}

type CreateUpdateOptions struct {
	AutoscaleSettings *AutoscaleSettings `pulumi:"autoscaleSettings"`
	Throughput        *int               `pulumi:"throughput"`
}





type CreateUpdateOptionsInput interface {
	pulumi.Input

	ToCreateUpdateOptionsOutput() CreateUpdateOptionsOutput
	ToCreateUpdateOptionsOutputWithContext(context.Context) CreateUpdateOptionsOutput
}

type CreateUpdateOptionsArgs struct {
	AutoscaleSettings AutoscaleSettingsPtrInput `pulumi:"autoscaleSettings"`
	Throughput        pulumi.IntPtrInput        `pulumi:"throughput"`
}

func (CreateUpdateOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateUpdateOptions)(nil)).Elem()
}

func (i CreateUpdateOptionsArgs) ToCreateUpdateOptionsOutput() CreateUpdateOptionsOutput {
	return i.ToCreateUpdateOptionsOutputWithContext(context.Background())
}

func (i CreateUpdateOptionsArgs) ToCreateUpdateOptionsOutputWithContext(ctx context.Context) CreateUpdateOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateUpdateOptionsOutput)
}

func (i CreateUpdateOptionsArgs) ToCreateUpdateOptionsPtrOutput() CreateUpdateOptionsPtrOutput {
	return i.ToCreateUpdateOptionsPtrOutputWithContext(context.Background())
}

func (i CreateUpdateOptionsArgs) ToCreateUpdateOptionsPtrOutputWithContext(ctx context.Context) CreateUpdateOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateUpdateOptionsOutput).ToCreateUpdateOptionsPtrOutputWithContext(ctx)
}









type CreateUpdateOptionsPtrInput interface {
	pulumi.Input

	ToCreateUpdateOptionsPtrOutput() CreateUpdateOptionsPtrOutput
	ToCreateUpdateOptionsPtrOutputWithContext(context.Context) CreateUpdateOptionsPtrOutput
}

type createUpdateOptionsPtrType CreateUpdateOptionsArgs

func CreateUpdateOptionsPtr(v *CreateUpdateOptionsArgs) CreateUpdateOptionsPtrInput {
	return (*createUpdateOptionsPtrType)(v)
}

func (*createUpdateOptionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CreateUpdateOptions)(nil)).Elem()
}

func (i *createUpdateOptionsPtrType) ToCreateUpdateOptionsPtrOutput() CreateUpdateOptionsPtrOutput {
	return i.ToCreateUpdateOptionsPtrOutputWithContext(context.Background())
}

func (i *createUpdateOptionsPtrType) ToCreateUpdateOptionsPtrOutputWithContext(ctx context.Context) CreateUpdateOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateUpdateOptionsPtrOutput)
}

type CreateUpdateOptionsOutput struct{ *pulumi.OutputState }

func (CreateUpdateOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateUpdateOptions)(nil)).Elem()
}

func (o CreateUpdateOptionsOutput) ToCreateUpdateOptionsOutput() CreateUpdateOptionsOutput {
	return o
}

func (o CreateUpdateOptionsOutput) ToCreateUpdateOptionsOutputWithContext(ctx context.Context) CreateUpdateOptionsOutput {
	return o
}

func (o CreateUpdateOptionsOutput) ToCreateUpdateOptionsPtrOutput() CreateUpdateOptionsPtrOutput {
	return o.ToCreateUpdateOptionsPtrOutputWithContext(context.Background())
}

func (o CreateUpdateOptionsOutput) ToCreateUpdateOptionsPtrOutputWithContext(ctx context.Context) CreateUpdateOptionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CreateUpdateOptions) *CreateUpdateOptions {
		return &v
	}).(CreateUpdateOptionsPtrOutput)
}

func (o CreateUpdateOptionsOutput) AutoscaleSettings() AutoscaleSettingsPtrOutput {
	return o.ApplyT(func(v CreateUpdateOptions) *AutoscaleSettings { return v.AutoscaleSettings }).(AutoscaleSettingsPtrOutput)
}

func (o CreateUpdateOptionsOutput) Throughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CreateUpdateOptions) *int { return v.Throughput }).(pulumi.IntPtrOutput)
}

type CreateUpdateOptionsPtrOutput struct{ *pulumi.OutputState }

func (CreateUpdateOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CreateUpdateOptions)(nil)).Elem()
}

func (o CreateUpdateOptionsPtrOutput) ToCreateUpdateOptionsPtrOutput() CreateUpdateOptionsPtrOutput {
	return o
}

func (o CreateUpdateOptionsPtrOutput) ToCreateUpdateOptionsPtrOutputWithContext(ctx context.Context) CreateUpdateOptionsPtrOutput {
	return o
}

func (o CreateUpdateOptionsPtrOutput) Elem() CreateUpdateOptionsOutput {
	return o.ApplyT(func(v *CreateUpdateOptions) CreateUpdateOptions {
		if v != nil {
			return *v
		}
		var ret CreateUpdateOptions
		return ret
	}).(CreateUpdateOptionsOutput)
}

func (o CreateUpdateOptionsPtrOutput) AutoscaleSettings() AutoscaleSettingsPtrOutput {
	return o.ApplyT(func(v *CreateUpdateOptions) *AutoscaleSettings {
		if v == nil {
			return nil
		}
		return v.AutoscaleSettings
	}).(AutoscaleSettingsPtrOutput)
}

func (o CreateUpdateOptionsPtrOutput) Throughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CreateUpdateOptions) *int {
		if v == nil {
			return nil
		}
		return v.Throughput
	}).(pulumi.IntPtrOutput)
}

type DatabaseAccountConnectionStringResponse struct {
	ConnectionString string `pulumi:"connectionString"`
	Description      string `pulumi:"description"`
}

type ExcludedPath struct {
	Path *string `pulumi:"path"`
}





type ExcludedPathInput interface {
	pulumi.Input

	ToExcludedPathOutput() ExcludedPathOutput
	ToExcludedPathOutputWithContext(context.Context) ExcludedPathOutput
}

type ExcludedPathArgs struct {
	Path pulumi.StringPtrInput `pulumi:"path"`
}

func (ExcludedPathArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExcludedPath)(nil)).Elem()
}

func (i ExcludedPathArgs) ToExcludedPathOutput() ExcludedPathOutput {
	return i.ToExcludedPathOutputWithContext(context.Background())
}

func (i ExcludedPathArgs) ToExcludedPathOutputWithContext(ctx context.Context) ExcludedPathOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExcludedPathOutput)
}





type ExcludedPathArrayInput interface {
	pulumi.Input

	ToExcludedPathArrayOutput() ExcludedPathArrayOutput
	ToExcludedPathArrayOutputWithContext(context.Context) ExcludedPathArrayOutput
}

type ExcludedPathArray []ExcludedPathInput

func (ExcludedPathArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExcludedPath)(nil)).Elem()
}

func (i ExcludedPathArray) ToExcludedPathArrayOutput() ExcludedPathArrayOutput {
	return i.ToExcludedPathArrayOutputWithContext(context.Background())
}

func (i ExcludedPathArray) ToExcludedPathArrayOutputWithContext(ctx context.Context) ExcludedPathArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExcludedPathArrayOutput)
}

type ExcludedPathOutput struct{ *pulumi.OutputState }

func (ExcludedPathOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExcludedPath)(nil)).Elem()
}

func (o ExcludedPathOutput) ToExcludedPathOutput() ExcludedPathOutput {
	return o
}

func (o ExcludedPathOutput) ToExcludedPathOutputWithContext(ctx context.Context) ExcludedPathOutput {
	return o
}

func (o ExcludedPathOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExcludedPath) *string { return v.Path }).(pulumi.StringPtrOutput)
}

type ExcludedPathArrayOutput struct{ *pulumi.OutputState }

func (ExcludedPathArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExcludedPath)(nil)).Elem()
}

func (o ExcludedPathArrayOutput) ToExcludedPathArrayOutput() ExcludedPathArrayOutput {
	return o
}

func (o ExcludedPathArrayOutput) ToExcludedPathArrayOutputWithContext(ctx context.Context) ExcludedPathArrayOutput {
	return o
}

func (o ExcludedPathArrayOutput) Index(i pulumi.IntInput) ExcludedPathOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExcludedPath {
		return vs[0].([]ExcludedPath)[vs[1].(int)]
	}).(ExcludedPathOutput)
}

type ExcludedPathResponse struct {
	Path *string `pulumi:"path"`
}

type ExcludedPathResponseOutput struct{ *pulumi.OutputState }

func (ExcludedPathResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExcludedPathResponse)(nil)).Elem()
}

func (o ExcludedPathResponseOutput) ToExcludedPathResponseOutput() ExcludedPathResponseOutput {
	return o
}

func (o ExcludedPathResponseOutput) ToExcludedPathResponseOutputWithContext(ctx context.Context) ExcludedPathResponseOutput {
	return o
}

func (o ExcludedPathResponseOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExcludedPathResponse) *string { return v.Path }).(pulumi.StringPtrOutput)
}

type ExcludedPathResponseArrayOutput struct{ *pulumi.OutputState }

func (ExcludedPathResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExcludedPathResponse)(nil)).Elem()
}

func (o ExcludedPathResponseArrayOutput) ToExcludedPathResponseArrayOutput() ExcludedPathResponseArrayOutput {
	return o
}

func (o ExcludedPathResponseArrayOutput) ToExcludedPathResponseArrayOutputWithContext(ctx context.Context) ExcludedPathResponseArrayOutput {
	return o
}

func (o ExcludedPathResponseArrayOutput) Index(i pulumi.IntInput) ExcludedPathResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExcludedPathResponse {
		return vs[0].([]ExcludedPathResponse)[vs[1].(int)]
	}).(ExcludedPathResponseOutput)
}

type FailoverPolicyResponse struct {
	FailoverPriority *int    `pulumi:"failoverPriority"`
	Id               string  `pulumi:"id"`
	LocationName     *string `pulumi:"locationName"`
}

type FailoverPolicyResponseOutput struct{ *pulumi.OutputState }

func (FailoverPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FailoverPolicyResponse)(nil)).Elem()
}

func (o FailoverPolicyResponseOutput) ToFailoverPolicyResponseOutput() FailoverPolicyResponseOutput {
	return o
}

func (o FailoverPolicyResponseOutput) ToFailoverPolicyResponseOutputWithContext(ctx context.Context) FailoverPolicyResponseOutput {
	return o
}

func (o FailoverPolicyResponseOutput) FailoverPriority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FailoverPolicyResponse) *int { return v.FailoverPriority }).(pulumi.IntPtrOutput)
}

func (o FailoverPolicyResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v FailoverPolicyResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o FailoverPolicyResponseOutput) LocationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FailoverPolicyResponse) *string { return v.LocationName }).(pulumi.StringPtrOutput)
}

type FailoverPolicyResponseArrayOutput struct{ *pulumi.OutputState }

func (FailoverPolicyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FailoverPolicyResponse)(nil)).Elem()
}

func (o FailoverPolicyResponseArrayOutput) ToFailoverPolicyResponseArrayOutput() FailoverPolicyResponseArrayOutput {
	return o
}

func (o FailoverPolicyResponseArrayOutput) ToFailoverPolicyResponseArrayOutputWithContext(ctx context.Context) FailoverPolicyResponseArrayOutput {
	return o
}

func (o FailoverPolicyResponseArrayOutput) Index(i pulumi.IntInput) FailoverPolicyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FailoverPolicyResponse {
		return vs[0].([]FailoverPolicyResponse)[vs[1].(int)]
	}).(FailoverPolicyResponseOutput)
}

type GremlinDatabaseGetPropertiesResponseOptions struct {
	AutoscaleSettings *AutoscaleSettingsResponse `pulumi:"autoscaleSettings"`
	Throughput        *int                       `pulumi:"throughput"`
}

type GremlinDatabaseGetPropertiesResponseOptionsOutput struct{ *pulumi.OutputState }

func (GremlinDatabaseGetPropertiesResponseOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GremlinDatabaseGetPropertiesResponseOptions)(nil)).Elem()
}

func (o GremlinDatabaseGetPropertiesResponseOptionsOutput) ToGremlinDatabaseGetPropertiesResponseOptionsOutput() GremlinDatabaseGetPropertiesResponseOptionsOutput {
	return o
}

func (o GremlinDatabaseGetPropertiesResponseOptionsOutput) ToGremlinDatabaseGetPropertiesResponseOptionsOutputWithContext(ctx context.Context) GremlinDatabaseGetPropertiesResponseOptionsOutput {
	return o
}

func (o GremlinDatabaseGetPropertiesResponseOptionsOutput) AutoscaleSettings() AutoscaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v GremlinDatabaseGetPropertiesResponseOptions) *AutoscaleSettingsResponse {
		return v.AutoscaleSettings
	}).(AutoscaleSettingsResponsePtrOutput)
}

func (o GremlinDatabaseGetPropertiesResponseOptionsOutput) Throughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GremlinDatabaseGetPropertiesResponseOptions) *int { return v.Throughput }).(pulumi.IntPtrOutput)
}

type GremlinDatabaseGetPropertiesResponseOptionsPtrOutput struct{ *pulumi.OutputState }

func (GremlinDatabaseGetPropertiesResponseOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GremlinDatabaseGetPropertiesResponseOptions)(nil)).Elem()
}

func (o GremlinDatabaseGetPropertiesResponseOptionsPtrOutput) ToGremlinDatabaseGetPropertiesResponseOptionsPtrOutput() GremlinDatabaseGetPropertiesResponseOptionsPtrOutput {
	return o
}

func (o GremlinDatabaseGetPropertiesResponseOptionsPtrOutput) ToGremlinDatabaseGetPropertiesResponseOptionsPtrOutputWithContext(ctx context.Context) GremlinDatabaseGetPropertiesResponseOptionsPtrOutput {
	return o
}

func (o GremlinDatabaseGetPropertiesResponseOptionsPtrOutput) Elem() GremlinDatabaseGetPropertiesResponseOptionsOutput {
	return o.ApplyT(func(v *GremlinDatabaseGetPropertiesResponseOptions) GremlinDatabaseGetPropertiesResponseOptions {
		if v != nil {
			return *v
		}
		var ret GremlinDatabaseGetPropertiesResponseOptions
		return ret
	}).(GremlinDatabaseGetPropertiesResponseOptionsOutput)
}

func (o GremlinDatabaseGetPropertiesResponseOptionsPtrOutput) AutoscaleSettings() AutoscaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v *GremlinDatabaseGetPropertiesResponseOptions) *AutoscaleSettingsResponse {
		if v == nil {
			return nil
		}
		return v.AutoscaleSettings
	}).(AutoscaleSettingsResponsePtrOutput)
}

func (o GremlinDatabaseGetPropertiesResponseOptionsPtrOutput) Throughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *GremlinDatabaseGetPropertiesResponseOptions) *int {
		if v == nil {
			return nil
		}
		return v.Throughput
	}).(pulumi.IntPtrOutput)
}

type GremlinDatabaseGetPropertiesResponseResource struct {
	Etag string  `pulumi:"etag"`
	Id   string  `pulumi:"id"`
	Rid  string  `pulumi:"rid"`
	Ts   float64 `pulumi:"ts"`
}

type GremlinDatabaseGetPropertiesResponseResourceOutput struct{ *pulumi.OutputState }

func (GremlinDatabaseGetPropertiesResponseResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GremlinDatabaseGetPropertiesResponseResource)(nil)).Elem()
}

func (o GremlinDatabaseGetPropertiesResponseResourceOutput) ToGremlinDatabaseGetPropertiesResponseResourceOutput() GremlinDatabaseGetPropertiesResponseResourceOutput {
	return o
}

func (o GremlinDatabaseGetPropertiesResponseResourceOutput) ToGremlinDatabaseGetPropertiesResponseResourceOutputWithContext(ctx context.Context) GremlinDatabaseGetPropertiesResponseResourceOutput {
	return o
}

func (o GremlinDatabaseGetPropertiesResponseResourceOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v GremlinDatabaseGetPropertiesResponseResource) string { return v.Etag }).(pulumi.StringOutput)
}

func (o GremlinDatabaseGetPropertiesResponseResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GremlinDatabaseGetPropertiesResponseResource) string { return v.Id }).(pulumi.StringOutput)
}

func (o GremlinDatabaseGetPropertiesResponseResourceOutput) Rid() pulumi.StringOutput {
	return o.ApplyT(func(v GremlinDatabaseGetPropertiesResponseResource) string { return v.Rid }).(pulumi.StringOutput)
}

func (o GremlinDatabaseGetPropertiesResponseResourceOutput) Ts() pulumi.Float64Output {
	return o.ApplyT(func(v GremlinDatabaseGetPropertiesResponseResource) float64 { return v.Ts }).(pulumi.Float64Output)
}

type GremlinDatabaseGetPropertiesResponseResourcePtrOutput struct{ *pulumi.OutputState }

func (GremlinDatabaseGetPropertiesResponseResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GremlinDatabaseGetPropertiesResponseResource)(nil)).Elem()
}

func (o GremlinDatabaseGetPropertiesResponseResourcePtrOutput) ToGremlinDatabaseGetPropertiesResponseResourcePtrOutput() GremlinDatabaseGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o GremlinDatabaseGetPropertiesResponseResourcePtrOutput) ToGremlinDatabaseGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) GremlinDatabaseGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o GremlinDatabaseGetPropertiesResponseResourcePtrOutput) Elem() GremlinDatabaseGetPropertiesResponseResourceOutput {
	return o.ApplyT(func(v *GremlinDatabaseGetPropertiesResponseResource) GremlinDatabaseGetPropertiesResponseResource {
		if v != nil {
			return *v
		}
		var ret GremlinDatabaseGetPropertiesResponseResource
		return ret
	}).(GremlinDatabaseGetPropertiesResponseResourceOutput)
}

func (o GremlinDatabaseGetPropertiesResponseResourcePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GremlinDatabaseGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o GremlinDatabaseGetPropertiesResponseResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GremlinDatabaseGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o GremlinDatabaseGetPropertiesResponseResourcePtrOutput) Rid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GremlinDatabaseGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Rid
	}).(pulumi.StringPtrOutput)
}

func (o GremlinDatabaseGetPropertiesResponseResourcePtrOutput) Ts() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *GremlinDatabaseGetPropertiesResponseResource) *float64 {
		if v == nil {
			return nil
		}
		return &v.Ts
	}).(pulumi.Float64PtrOutput)
}

type GremlinDatabaseResource struct {
	Id string `pulumi:"id"`
}





type GremlinDatabaseResourceInput interface {
	pulumi.Input

	ToGremlinDatabaseResourceOutput() GremlinDatabaseResourceOutput
	ToGremlinDatabaseResourceOutputWithContext(context.Context) GremlinDatabaseResourceOutput
}

type GremlinDatabaseResourceArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (GremlinDatabaseResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GremlinDatabaseResource)(nil)).Elem()
}

func (i GremlinDatabaseResourceArgs) ToGremlinDatabaseResourceOutput() GremlinDatabaseResourceOutput {
	return i.ToGremlinDatabaseResourceOutputWithContext(context.Background())
}

func (i GremlinDatabaseResourceArgs) ToGremlinDatabaseResourceOutputWithContext(ctx context.Context) GremlinDatabaseResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GremlinDatabaseResourceOutput)
}

type GremlinDatabaseResourceOutput struct{ *pulumi.OutputState }

func (GremlinDatabaseResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GremlinDatabaseResource)(nil)).Elem()
}

func (o GremlinDatabaseResourceOutput) ToGremlinDatabaseResourceOutput() GremlinDatabaseResourceOutput {
	return o
}

func (o GremlinDatabaseResourceOutput) ToGremlinDatabaseResourceOutputWithContext(ctx context.Context) GremlinDatabaseResourceOutput {
	return o
}

func (o GremlinDatabaseResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GremlinDatabaseResource) string { return v.Id }).(pulumi.StringOutput)
}

type GremlinGraphGetPropertiesResponseOptions struct {
	AutoscaleSettings *AutoscaleSettingsResponse `pulumi:"autoscaleSettings"`
	Throughput        *int                       `pulumi:"throughput"`
}

type GremlinGraphGetPropertiesResponseOptionsOutput struct{ *pulumi.OutputState }

func (GremlinGraphGetPropertiesResponseOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GremlinGraphGetPropertiesResponseOptions)(nil)).Elem()
}

func (o GremlinGraphGetPropertiesResponseOptionsOutput) ToGremlinGraphGetPropertiesResponseOptionsOutput() GremlinGraphGetPropertiesResponseOptionsOutput {
	return o
}

func (o GremlinGraphGetPropertiesResponseOptionsOutput) ToGremlinGraphGetPropertiesResponseOptionsOutputWithContext(ctx context.Context) GremlinGraphGetPropertiesResponseOptionsOutput {
	return o
}

func (o GremlinGraphGetPropertiesResponseOptionsOutput) AutoscaleSettings() AutoscaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v GremlinGraphGetPropertiesResponseOptions) *AutoscaleSettingsResponse {
		return v.AutoscaleSettings
	}).(AutoscaleSettingsResponsePtrOutput)
}

func (o GremlinGraphGetPropertiesResponseOptionsOutput) Throughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GremlinGraphGetPropertiesResponseOptions) *int { return v.Throughput }).(pulumi.IntPtrOutput)
}

type GremlinGraphGetPropertiesResponseOptionsPtrOutput struct{ *pulumi.OutputState }

func (GremlinGraphGetPropertiesResponseOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GremlinGraphGetPropertiesResponseOptions)(nil)).Elem()
}

func (o GremlinGraphGetPropertiesResponseOptionsPtrOutput) ToGremlinGraphGetPropertiesResponseOptionsPtrOutput() GremlinGraphGetPropertiesResponseOptionsPtrOutput {
	return o
}

func (o GremlinGraphGetPropertiesResponseOptionsPtrOutput) ToGremlinGraphGetPropertiesResponseOptionsPtrOutputWithContext(ctx context.Context) GremlinGraphGetPropertiesResponseOptionsPtrOutput {
	return o
}

func (o GremlinGraphGetPropertiesResponseOptionsPtrOutput) Elem() GremlinGraphGetPropertiesResponseOptionsOutput {
	return o.ApplyT(func(v *GremlinGraphGetPropertiesResponseOptions) GremlinGraphGetPropertiesResponseOptions {
		if v != nil {
			return *v
		}
		var ret GremlinGraphGetPropertiesResponseOptions
		return ret
	}).(GremlinGraphGetPropertiesResponseOptionsOutput)
}

func (o GremlinGraphGetPropertiesResponseOptionsPtrOutput) AutoscaleSettings() AutoscaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v *GremlinGraphGetPropertiesResponseOptions) *AutoscaleSettingsResponse {
		if v == nil {
			return nil
		}
		return v.AutoscaleSettings
	}).(AutoscaleSettingsResponsePtrOutput)
}

func (o GremlinGraphGetPropertiesResponseOptionsPtrOutput) Throughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *GremlinGraphGetPropertiesResponseOptions) *int {
		if v == nil {
			return nil
		}
		return v.Throughput
	}).(pulumi.IntPtrOutput)
}

type GremlinGraphGetPropertiesResponseResource struct {
	ConflictResolutionPolicy *ConflictResolutionPolicyResponse `pulumi:"conflictResolutionPolicy"`
	DefaultTtl               *int                              `pulumi:"defaultTtl"`
	Etag                     string                            `pulumi:"etag"`
	Id                       string                            `pulumi:"id"`
	IndexingPolicy           *IndexingPolicyResponse           `pulumi:"indexingPolicy"`
	PartitionKey             *ContainerPartitionKeyResponse    `pulumi:"partitionKey"`
	Rid                      string                            `pulumi:"rid"`
	Ts                       float64                           `pulumi:"ts"`
	UniqueKeyPolicy          *UniqueKeyPolicyResponse          `pulumi:"uniqueKeyPolicy"`
}


func (val *GremlinGraphGetPropertiesResponseResource) Defaults() *GremlinGraphGetPropertiesResponseResource {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ConflictResolutionPolicy = tmp.ConflictResolutionPolicy.Defaults()

	tmp.IndexingPolicy = tmp.IndexingPolicy.Defaults()

	tmp.PartitionKey = tmp.PartitionKey.Defaults()

	return &tmp
}

type GremlinGraphGetPropertiesResponseResourceOutput struct{ *pulumi.OutputState }

func (GremlinGraphGetPropertiesResponseResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GremlinGraphGetPropertiesResponseResource)(nil)).Elem()
}

func (o GremlinGraphGetPropertiesResponseResourceOutput) ToGremlinGraphGetPropertiesResponseResourceOutput() GremlinGraphGetPropertiesResponseResourceOutput {
	return o
}

func (o GremlinGraphGetPropertiesResponseResourceOutput) ToGremlinGraphGetPropertiesResponseResourceOutputWithContext(ctx context.Context) GremlinGraphGetPropertiesResponseResourceOutput {
	return o
}

func (o GremlinGraphGetPropertiesResponseResourceOutput) ConflictResolutionPolicy() ConflictResolutionPolicyResponsePtrOutput {
	return o.ApplyT(func(v GremlinGraphGetPropertiesResponseResource) *ConflictResolutionPolicyResponse {
		return v.ConflictResolutionPolicy
	}).(ConflictResolutionPolicyResponsePtrOutput)
}

func (o GremlinGraphGetPropertiesResponseResourceOutput) DefaultTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GremlinGraphGetPropertiesResponseResource) *int { return v.DefaultTtl }).(pulumi.IntPtrOutput)
}

func (o GremlinGraphGetPropertiesResponseResourceOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v GremlinGraphGetPropertiesResponseResource) string { return v.Etag }).(pulumi.StringOutput)
}

func (o GremlinGraphGetPropertiesResponseResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GremlinGraphGetPropertiesResponseResource) string { return v.Id }).(pulumi.StringOutput)
}

func (o GremlinGraphGetPropertiesResponseResourceOutput) IndexingPolicy() IndexingPolicyResponsePtrOutput {
	return o.ApplyT(func(v GremlinGraphGetPropertiesResponseResource) *IndexingPolicyResponse { return v.IndexingPolicy }).(IndexingPolicyResponsePtrOutput)
}

func (o GremlinGraphGetPropertiesResponseResourceOutput) PartitionKey() ContainerPartitionKeyResponsePtrOutput {
	return o.ApplyT(func(v GremlinGraphGetPropertiesResponseResource) *ContainerPartitionKeyResponse {
		return v.PartitionKey
	}).(ContainerPartitionKeyResponsePtrOutput)
}

func (o GremlinGraphGetPropertiesResponseResourceOutput) Rid() pulumi.StringOutput {
	return o.ApplyT(func(v GremlinGraphGetPropertiesResponseResource) string { return v.Rid }).(pulumi.StringOutput)
}

func (o GremlinGraphGetPropertiesResponseResourceOutput) Ts() pulumi.Float64Output {
	return o.ApplyT(func(v GremlinGraphGetPropertiesResponseResource) float64 { return v.Ts }).(pulumi.Float64Output)
}

func (o GremlinGraphGetPropertiesResponseResourceOutput) UniqueKeyPolicy() UniqueKeyPolicyResponsePtrOutput {
	return o.ApplyT(func(v GremlinGraphGetPropertiesResponseResource) *UniqueKeyPolicyResponse { return v.UniqueKeyPolicy }).(UniqueKeyPolicyResponsePtrOutput)
}

type GremlinGraphGetPropertiesResponseResourcePtrOutput struct{ *pulumi.OutputState }

func (GremlinGraphGetPropertiesResponseResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GremlinGraphGetPropertiesResponseResource)(nil)).Elem()
}

func (o GremlinGraphGetPropertiesResponseResourcePtrOutput) ToGremlinGraphGetPropertiesResponseResourcePtrOutput() GremlinGraphGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o GremlinGraphGetPropertiesResponseResourcePtrOutput) ToGremlinGraphGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) GremlinGraphGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o GremlinGraphGetPropertiesResponseResourcePtrOutput) Elem() GremlinGraphGetPropertiesResponseResourceOutput {
	return o.ApplyT(func(v *GremlinGraphGetPropertiesResponseResource) GremlinGraphGetPropertiesResponseResource {
		if v != nil {
			return *v
		}
		var ret GremlinGraphGetPropertiesResponseResource
		return ret
	}).(GremlinGraphGetPropertiesResponseResourceOutput)
}

func (o GremlinGraphGetPropertiesResponseResourcePtrOutput) ConflictResolutionPolicy() ConflictResolutionPolicyResponsePtrOutput {
	return o.ApplyT(func(v *GremlinGraphGetPropertiesResponseResource) *ConflictResolutionPolicyResponse {
		if v == nil {
			return nil
		}
		return v.ConflictResolutionPolicy
	}).(ConflictResolutionPolicyResponsePtrOutput)
}

func (o GremlinGraphGetPropertiesResponseResourcePtrOutput) DefaultTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *GremlinGraphGetPropertiesResponseResource) *int {
		if v == nil {
			return nil
		}
		return v.DefaultTtl
	}).(pulumi.IntPtrOutput)
}

func (o GremlinGraphGetPropertiesResponseResourcePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GremlinGraphGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o GremlinGraphGetPropertiesResponseResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GremlinGraphGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o GremlinGraphGetPropertiesResponseResourcePtrOutput) IndexingPolicy() IndexingPolicyResponsePtrOutput {
	return o.ApplyT(func(v *GremlinGraphGetPropertiesResponseResource) *IndexingPolicyResponse {
		if v == nil {
			return nil
		}
		return v.IndexingPolicy
	}).(IndexingPolicyResponsePtrOutput)
}

func (o GremlinGraphGetPropertiesResponseResourcePtrOutput) PartitionKey() ContainerPartitionKeyResponsePtrOutput {
	return o.ApplyT(func(v *GremlinGraphGetPropertiesResponseResource) *ContainerPartitionKeyResponse {
		if v == nil {
			return nil
		}
		return v.PartitionKey
	}).(ContainerPartitionKeyResponsePtrOutput)
}

func (o GremlinGraphGetPropertiesResponseResourcePtrOutput) Rid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GremlinGraphGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Rid
	}).(pulumi.StringPtrOutput)
}

func (o GremlinGraphGetPropertiesResponseResourcePtrOutput) Ts() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *GremlinGraphGetPropertiesResponseResource) *float64 {
		if v == nil {
			return nil
		}
		return &v.Ts
	}).(pulumi.Float64PtrOutput)
}

func (o GremlinGraphGetPropertiesResponseResourcePtrOutput) UniqueKeyPolicy() UniqueKeyPolicyResponsePtrOutput {
	return o.ApplyT(func(v *GremlinGraphGetPropertiesResponseResource) *UniqueKeyPolicyResponse {
		if v == nil {
			return nil
		}
		return v.UniqueKeyPolicy
	}).(UniqueKeyPolicyResponsePtrOutput)
}

type GremlinGraphResource struct {
	ConflictResolutionPolicy *ConflictResolutionPolicy `pulumi:"conflictResolutionPolicy"`
	DefaultTtl               *int                      `pulumi:"defaultTtl"`
	Id                       string                    `pulumi:"id"`
	IndexingPolicy           *IndexingPolicy           `pulumi:"indexingPolicy"`
	PartitionKey             *ContainerPartitionKey    `pulumi:"partitionKey"`
	UniqueKeyPolicy          *UniqueKeyPolicy          `pulumi:"uniqueKeyPolicy"`
}


func (val *GremlinGraphResource) Defaults() *GremlinGraphResource {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ConflictResolutionPolicy = tmp.ConflictResolutionPolicy.Defaults()

	tmp.IndexingPolicy = tmp.IndexingPolicy.Defaults()

	tmp.PartitionKey = tmp.PartitionKey.Defaults()

	return &tmp
}





type GremlinGraphResourceInput interface {
	pulumi.Input

	ToGremlinGraphResourceOutput() GremlinGraphResourceOutput
	ToGremlinGraphResourceOutputWithContext(context.Context) GremlinGraphResourceOutput
}

type GremlinGraphResourceArgs struct {
	ConflictResolutionPolicy ConflictResolutionPolicyPtrInput `pulumi:"conflictResolutionPolicy"`
	DefaultTtl               pulumi.IntPtrInput               `pulumi:"defaultTtl"`
	Id                       pulumi.StringInput               `pulumi:"id"`
	IndexingPolicy           IndexingPolicyPtrInput           `pulumi:"indexingPolicy"`
	PartitionKey             ContainerPartitionKeyPtrInput    `pulumi:"partitionKey"`
	UniqueKeyPolicy          UniqueKeyPolicyPtrInput          `pulumi:"uniqueKeyPolicy"`
}

func (GremlinGraphResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GremlinGraphResource)(nil)).Elem()
}

func (i GremlinGraphResourceArgs) ToGremlinGraphResourceOutput() GremlinGraphResourceOutput {
	return i.ToGremlinGraphResourceOutputWithContext(context.Background())
}

func (i GremlinGraphResourceArgs) ToGremlinGraphResourceOutputWithContext(ctx context.Context) GremlinGraphResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GremlinGraphResourceOutput)
}

type GremlinGraphResourceOutput struct{ *pulumi.OutputState }

func (GremlinGraphResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GremlinGraphResource)(nil)).Elem()
}

func (o GremlinGraphResourceOutput) ToGremlinGraphResourceOutput() GremlinGraphResourceOutput {
	return o
}

func (o GremlinGraphResourceOutput) ToGremlinGraphResourceOutputWithContext(ctx context.Context) GremlinGraphResourceOutput {
	return o
}

func (o GremlinGraphResourceOutput) ConflictResolutionPolicy() ConflictResolutionPolicyPtrOutput {
	return o.ApplyT(func(v GremlinGraphResource) *ConflictResolutionPolicy { return v.ConflictResolutionPolicy }).(ConflictResolutionPolicyPtrOutput)
}

func (o GremlinGraphResourceOutput) DefaultTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GremlinGraphResource) *int { return v.DefaultTtl }).(pulumi.IntPtrOutput)
}

func (o GremlinGraphResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GremlinGraphResource) string { return v.Id }).(pulumi.StringOutput)
}

func (o GremlinGraphResourceOutput) IndexingPolicy() IndexingPolicyPtrOutput {
	return o.ApplyT(func(v GremlinGraphResource) *IndexingPolicy { return v.IndexingPolicy }).(IndexingPolicyPtrOutput)
}

func (o GremlinGraphResourceOutput) PartitionKey() ContainerPartitionKeyPtrOutput {
	return o.ApplyT(func(v GremlinGraphResource) *ContainerPartitionKey { return v.PartitionKey }).(ContainerPartitionKeyPtrOutput)
}

func (o GremlinGraphResourceOutput) UniqueKeyPolicy() UniqueKeyPolicyPtrOutput {
	return o.ApplyT(func(v GremlinGraphResource) *UniqueKeyPolicy { return v.UniqueKeyPolicy }).(UniqueKeyPolicyPtrOutput)
}

type IncludedPath struct {
	Indexes []Indexes `pulumi:"indexes"`
	Path    *string   `pulumi:"path"`
}





type IncludedPathInput interface {
	pulumi.Input

	ToIncludedPathOutput() IncludedPathOutput
	ToIncludedPathOutputWithContext(context.Context) IncludedPathOutput
}

type IncludedPathArgs struct {
	Indexes IndexesArrayInput     `pulumi:"indexes"`
	Path    pulumi.StringPtrInput `pulumi:"path"`
}

func (IncludedPathArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IncludedPath)(nil)).Elem()
}

func (i IncludedPathArgs) ToIncludedPathOutput() IncludedPathOutput {
	return i.ToIncludedPathOutputWithContext(context.Background())
}

func (i IncludedPathArgs) ToIncludedPathOutputWithContext(ctx context.Context) IncludedPathOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IncludedPathOutput)
}





type IncludedPathArrayInput interface {
	pulumi.Input

	ToIncludedPathArrayOutput() IncludedPathArrayOutput
	ToIncludedPathArrayOutputWithContext(context.Context) IncludedPathArrayOutput
}

type IncludedPathArray []IncludedPathInput

func (IncludedPathArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IncludedPath)(nil)).Elem()
}

func (i IncludedPathArray) ToIncludedPathArrayOutput() IncludedPathArrayOutput {
	return i.ToIncludedPathArrayOutputWithContext(context.Background())
}

func (i IncludedPathArray) ToIncludedPathArrayOutputWithContext(ctx context.Context) IncludedPathArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IncludedPathArrayOutput)
}

type IncludedPathOutput struct{ *pulumi.OutputState }

func (IncludedPathOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IncludedPath)(nil)).Elem()
}

func (o IncludedPathOutput) ToIncludedPathOutput() IncludedPathOutput {
	return o
}

func (o IncludedPathOutput) ToIncludedPathOutputWithContext(ctx context.Context) IncludedPathOutput {
	return o
}

func (o IncludedPathOutput) Indexes() IndexesArrayOutput {
	return o.ApplyT(func(v IncludedPath) []Indexes { return v.Indexes }).(IndexesArrayOutput)
}

func (o IncludedPathOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IncludedPath) *string { return v.Path }).(pulumi.StringPtrOutput)
}

type IncludedPathArrayOutput struct{ *pulumi.OutputState }

func (IncludedPathArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IncludedPath)(nil)).Elem()
}

func (o IncludedPathArrayOutput) ToIncludedPathArrayOutput() IncludedPathArrayOutput {
	return o
}

func (o IncludedPathArrayOutput) ToIncludedPathArrayOutputWithContext(ctx context.Context) IncludedPathArrayOutput {
	return o
}

func (o IncludedPathArrayOutput) Index(i pulumi.IntInput) IncludedPathOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IncludedPath {
		return vs[0].([]IncludedPath)[vs[1].(int)]
	}).(IncludedPathOutput)
}

type IncludedPathResponse struct {
	Indexes []IndexesResponse `pulumi:"indexes"`
	Path    *string           `pulumi:"path"`
}

type IncludedPathResponseOutput struct{ *pulumi.OutputState }

func (IncludedPathResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IncludedPathResponse)(nil)).Elem()
}

func (o IncludedPathResponseOutput) ToIncludedPathResponseOutput() IncludedPathResponseOutput {
	return o
}

func (o IncludedPathResponseOutput) ToIncludedPathResponseOutputWithContext(ctx context.Context) IncludedPathResponseOutput {
	return o
}

func (o IncludedPathResponseOutput) Indexes() IndexesResponseArrayOutput {
	return o.ApplyT(func(v IncludedPathResponse) []IndexesResponse { return v.Indexes }).(IndexesResponseArrayOutput)
}

func (o IncludedPathResponseOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IncludedPathResponse) *string { return v.Path }).(pulumi.StringPtrOutput)
}

type IncludedPathResponseArrayOutput struct{ *pulumi.OutputState }

func (IncludedPathResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IncludedPathResponse)(nil)).Elem()
}

func (o IncludedPathResponseArrayOutput) ToIncludedPathResponseArrayOutput() IncludedPathResponseArrayOutput {
	return o
}

func (o IncludedPathResponseArrayOutput) ToIncludedPathResponseArrayOutputWithContext(ctx context.Context) IncludedPathResponseArrayOutput {
	return o
}

func (o IncludedPathResponseArrayOutput) Index(i pulumi.IntInput) IncludedPathResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IncludedPathResponse {
		return vs[0].([]IncludedPathResponse)[vs[1].(int)]
	}).(IncludedPathResponseOutput)
}

type Indexes struct {
	DataType  *string `pulumi:"dataType"`
	Kind      *string `pulumi:"kind"`
	Precision *int    `pulumi:"precision"`
}


func (val *Indexes) Defaults() *Indexes {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DataType) {
		dataType_ := "String"
		tmp.DataType = &dataType_
	}
	if isZero(tmp.Kind) {
		kind_ := "Hash"
		tmp.Kind = &kind_
	}
	return &tmp
}





type IndexesInput interface {
	pulumi.Input

	ToIndexesOutput() IndexesOutput
	ToIndexesOutputWithContext(context.Context) IndexesOutput
}

type IndexesArgs struct {
	DataType  pulumi.StringPtrInput `pulumi:"dataType"`
	Kind      pulumi.StringPtrInput `pulumi:"kind"`
	Precision pulumi.IntPtrInput    `pulumi:"precision"`
}

func (IndexesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Indexes)(nil)).Elem()
}

func (i IndexesArgs) ToIndexesOutput() IndexesOutput {
	return i.ToIndexesOutputWithContext(context.Background())
}

func (i IndexesArgs) ToIndexesOutputWithContext(ctx context.Context) IndexesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IndexesOutput)
}





type IndexesArrayInput interface {
	pulumi.Input

	ToIndexesArrayOutput() IndexesArrayOutput
	ToIndexesArrayOutputWithContext(context.Context) IndexesArrayOutput
}

type IndexesArray []IndexesInput

func (IndexesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Indexes)(nil)).Elem()
}

func (i IndexesArray) ToIndexesArrayOutput() IndexesArrayOutput {
	return i.ToIndexesArrayOutputWithContext(context.Background())
}

func (i IndexesArray) ToIndexesArrayOutputWithContext(ctx context.Context) IndexesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IndexesArrayOutput)
}

type IndexesOutput struct{ *pulumi.OutputState }

func (IndexesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Indexes)(nil)).Elem()
}

func (o IndexesOutput) ToIndexesOutput() IndexesOutput {
	return o
}

func (o IndexesOutput) ToIndexesOutputWithContext(ctx context.Context) IndexesOutput {
	return o
}

func (o IndexesOutput) DataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Indexes) *string { return v.DataType }).(pulumi.StringPtrOutput)
}

func (o IndexesOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Indexes) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o IndexesOutput) Precision() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Indexes) *int { return v.Precision }).(pulumi.IntPtrOutput)
}

type IndexesArrayOutput struct{ *pulumi.OutputState }

func (IndexesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Indexes)(nil)).Elem()
}

func (o IndexesArrayOutput) ToIndexesArrayOutput() IndexesArrayOutput {
	return o
}

func (o IndexesArrayOutput) ToIndexesArrayOutputWithContext(ctx context.Context) IndexesArrayOutput {
	return o
}

func (o IndexesArrayOutput) Index(i pulumi.IntInput) IndexesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Indexes {
		return vs[0].([]Indexes)[vs[1].(int)]
	}).(IndexesOutput)
}

type IndexesResponse struct {
	DataType  *string `pulumi:"dataType"`
	Kind      *string `pulumi:"kind"`
	Precision *int    `pulumi:"precision"`
}


func (val *IndexesResponse) Defaults() *IndexesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DataType) {
		dataType_ := "String"
		tmp.DataType = &dataType_
	}
	if isZero(tmp.Kind) {
		kind_ := "Hash"
		tmp.Kind = &kind_
	}
	return &tmp
}

type IndexesResponseOutput struct{ *pulumi.OutputState }

func (IndexesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IndexesResponse)(nil)).Elem()
}

func (o IndexesResponseOutput) ToIndexesResponseOutput() IndexesResponseOutput {
	return o
}

func (o IndexesResponseOutput) ToIndexesResponseOutputWithContext(ctx context.Context) IndexesResponseOutput {
	return o
}

func (o IndexesResponseOutput) DataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IndexesResponse) *string { return v.DataType }).(pulumi.StringPtrOutput)
}

func (o IndexesResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IndexesResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o IndexesResponseOutput) Precision() pulumi.IntPtrOutput {
	return o.ApplyT(func(v IndexesResponse) *int { return v.Precision }).(pulumi.IntPtrOutput)
}

type IndexesResponseArrayOutput struct{ *pulumi.OutputState }

func (IndexesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IndexesResponse)(nil)).Elem()
}

func (o IndexesResponseArrayOutput) ToIndexesResponseArrayOutput() IndexesResponseArrayOutput {
	return o
}

func (o IndexesResponseArrayOutput) ToIndexesResponseArrayOutputWithContext(ctx context.Context) IndexesResponseArrayOutput {
	return o
}

func (o IndexesResponseArrayOutput) Index(i pulumi.IntInput) IndexesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IndexesResponse {
		return vs[0].([]IndexesResponse)[vs[1].(int)]
	}).(IndexesResponseOutput)
}

type IndexingPolicy struct {
	Automatic        *bool             `pulumi:"automatic"`
	CompositeIndexes [][]CompositePath `pulumi:"compositeIndexes"`
	ExcludedPaths    []ExcludedPath    `pulumi:"excludedPaths"`
	IncludedPaths    []IncludedPath    `pulumi:"includedPaths"`
	IndexingMode     *string           `pulumi:"indexingMode"`
	SpatialIndexes   []SpatialSpec     `pulumi:"spatialIndexes"`
}


func (val *IndexingPolicy) Defaults() *IndexingPolicy {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IndexingMode) {
		indexingMode_ := "Consistent"
		tmp.IndexingMode = &indexingMode_
	}
	return &tmp
}





type IndexingPolicyInput interface {
	pulumi.Input

	ToIndexingPolicyOutput() IndexingPolicyOutput
	ToIndexingPolicyOutputWithContext(context.Context) IndexingPolicyOutput
}

type IndexingPolicyArgs struct {
	Automatic        pulumi.BoolPtrInput          `pulumi:"automatic"`
	CompositeIndexes CompositePathArrayArrayInput `pulumi:"compositeIndexes"`
	ExcludedPaths    ExcludedPathArrayInput       `pulumi:"excludedPaths"`
	IncludedPaths    IncludedPathArrayInput       `pulumi:"includedPaths"`
	IndexingMode     pulumi.StringPtrInput        `pulumi:"indexingMode"`
	SpatialIndexes   SpatialSpecArrayInput        `pulumi:"spatialIndexes"`
}

func (IndexingPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IndexingPolicy)(nil)).Elem()
}

func (i IndexingPolicyArgs) ToIndexingPolicyOutput() IndexingPolicyOutput {
	return i.ToIndexingPolicyOutputWithContext(context.Background())
}

func (i IndexingPolicyArgs) ToIndexingPolicyOutputWithContext(ctx context.Context) IndexingPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IndexingPolicyOutput)
}

func (i IndexingPolicyArgs) ToIndexingPolicyPtrOutput() IndexingPolicyPtrOutput {
	return i.ToIndexingPolicyPtrOutputWithContext(context.Background())
}

func (i IndexingPolicyArgs) ToIndexingPolicyPtrOutputWithContext(ctx context.Context) IndexingPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IndexingPolicyOutput).ToIndexingPolicyPtrOutputWithContext(ctx)
}









type IndexingPolicyPtrInput interface {
	pulumi.Input

	ToIndexingPolicyPtrOutput() IndexingPolicyPtrOutput
	ToIndexingPolicyPtrOutputWithContext(context.Context) IndexingPolicyPtrOutput
}

type indexingPolicyPtrType IndexingPolicyArgs

func IndexingPolicyPtr(v *IndexingPolicyArgs) IndexingPolicyPtrInput {
	return (*indexingPolicyPtrType)(v)
}

func (*indexingPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IndexingPolicy)(nil)).Elem()
}

func (i *indexingPolicyPtrType) ToIndexingPolicyPtrOutput() IndexingPolicyPtrOutput {
	return i.ToIndexingPolicyPtrOutputWithContext(context.Background())
}

func (i *indexingPolicyPtrType) ToIndexingPolicyPtrOutputWithContext(ctx context.Context) IndexingPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IndexingPolicyPtrOutput)
}

type IndexingPolicyOutput struct{ *pulumi.OutputState }

func (IndexingPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IndexingPolicy)(nil)).Elem()
}

func (o IndexingPolicyOutput) ToIndexingPolicyOutput() IndexingPolicyOutput {
	return o
}

func (o IndexingPolicyOutput) ToIndexingPolicyOutputWithContext(ctx context.Context) IndexingPolicyOutput {
	return o
}

func (o IndexingPolicyOutput) ToIndexingPolicyPtrOutput() IndexingPolicyPtrOutput {
	return o.ToIndexingPolicyPtrOutputWithContext(context.Background())
}

func (o IndexingPolicyOutput) ToIndexingPolicyPtrOutputWithContext(ctx context.Context) IndexingPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IndexingPolicy) *IndexingPolicy {
		return &v
	}).(IndexingPolicyPtrOutput)
}

func (o IndexingPolicyOutput) Automatic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IndexingPolicy) *bool { return v.Automatic }).(pulumi.BoolPtrOutput)
}

func (o IndexingPolicyOutput) CompositeIndexes() CompositePathArrayArrayOutput {
	return o.ApplyT(func(v IndexingPolicy) [][]CompositePath { return v.CompositeIndexes }).(CompositePathArrayArrayOutput)
}

func (o IndexingPolicyOutput) ExcludedPaths() ExcludedPathArrayOutput {
	return o.ApplyT(func(v IndexingPolicy) []ExcludedPath { return v.ExcludedPaths }).(ExcludedPathArrayOutput)
}

func (o IndexingPolicyOutput) IncludedPaths() IncludedPathArrayOutput {
	return o.ApplyT(func(v IndexingPolicy) []IncludedPath { return v.IncludedPaths }).(IncludedPathArrayOutput)
}

func (o IndexingPolicyOutput) IndexingMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IndexingPolicy) *string { return v.IndexingMode }).(pulumi.StringPtrOutput)
}

func (o IndexingPolicyOutput) SpatialIndexes() SpatialSpecArrayOutput {
	return o.ApplyT(func(v IndexingPolicy) []SpatialSpec { return v.SpatialIndexes }).(SpatialSpecArrayOutput)
}

type IndexingPolicyPtrOutput struct{ *pulumi.OutputState }

func (IndexingPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IndexingPolicy)(nil)).Elem()
}

func (o IndexingPolicyPtrOutput) ToIndexingPolicyPtrOutput() IndexingPolicyPtrOutput {
	return o
}

func (o IndexingPolicyPtrOutput) ToIndexingPolicyPtrOutputWithContext(ctx context.Context) IndexingPolicyPtrOutput {
	return o
}

func (o IndexingPolicyPtrOutput) Elem() IndexingPolicyOutput {
	return o.ApplyT(func(v *IndexingPolicy) IndexingPolicy {
		if v != nil {
			return *v
		}
		var ret IndexingPolicy
		return ret
	}).(IndexingPolicyOutput)
}

func (o IndexingPolicyPtrOutput) Automatic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IndexingPolicy) *bool {
		if v == nil {
			return nil
		}
		return v.Automatic
	}).(pulumi.BoolPtrOutput)
}

func (o IndexingPolicyPtrOutput) CompositeIndexes() CompositePathArrayArrayOutput {
	return o.ApplyT(func(v *IndexingPolicy) [][]CompositePath {
		if v == nil {
			return nil
		}
		return v.CompositeIndexes
	}).(CompositePathArrayArrayOutput)
}

func (o IndexingPolicyPtrOutput) ExcludedPaths() ExcludedPathArrayOutput {
	return o.ApplyT(func(v *IndexingPolicy) []ExcludedPath {
		if v == nil {
			return nil
		}
		return v.ExcludedPaths
	}).(ExcludedPathArrayOutput)
}

func (o IndexingPolicyPtrOutput) IncludedPaths() IncludedPathArrayOutput {
	return o.ApplyT(func(v *IndexingPolicy) []IncludedPath {
		if v == nil {
			return nil
		}
		return v.IncludedPaths
	}).(IncludedPathArrayOutput)
}

func (o IndexingPolicyPtrOutput) IndexingMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IndexingPolicy) *string {
		if v == nil {
			return nil
		}
		return v.IndexingMode
	}).(pulumi.StringPtrOutput)
}

func (o IndexingPolicyPtrOutput) SpatialIndexes() SpatialSpecArrayOutput {
	return o.ApplyT(func(v *IndexingPolicy) []SpatialSpec {
		if v == nil {
			return nil
		}
		return v.SpatialIndexes
	}).(SpatialSpecArrayOutput)
}

type IndexingPolicyResponse struct {
	Automatic        *bool                     `pulumi:"automatic"`
	CompositeIndexes [][]CompositePathResponse `pulumi:"compositeIndexes"`
	ExcludedPaths    []ExcludedPathResponse    `pulumi:"excludedPaths"`
	IncludedPaths    []IncludedPathResponse    `pulumi:"includedPaths"`
	IndexingMode     *string                   `pulumi:"indexingMode"`
	SpatialIndexes   []SpatialSpecResponse     `pulumi:"spatialIndexes"`
}


func (val *IndexingPolicyResponse) Defaults() *IndexingPolicyResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IndexingMode) {
		indexingMode_ := "Consistent"
		tmp.IndexingMode = &indexingMode_
	}
	return &tmp
}

type IndexingPolicyResponseOutput struct{ *pulumi.OutputState }

func (IndexingPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IndexingPolicyResponse)(nil)).Elem()
}

func (o IndexingPolicyResponseOutput) ToIndexingPolicyResponseOutput() IndexingPolicyResponseOutput {
	return o
}

func (o IndexingPolicyResponseOutput) ToIndexingPolicyResponseOutputWithContext(ctx context.Context) IndexingPolicyResponseOutput {
	return o
}

func (o IndexingPolicyResponseOutput) Automatic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IndexingPolicyResponse) *bool { return v.Automatic }).(pulumi.BoolPtrOutput)
}

func (o IndexingPolicyResponseOutput) CompositeIndexes() CompositePathResponseArrayArrayOutput {
	return o.ApplyT(func(v IndexingPolicyResponse) [][]CompositePathResponse { return v.CompositeIndexes }).(CompositePathResponseArrayArrayOutput)
}

func (o IndexingPolicyResponseOutput) ExcludedPaths() ExcludedPathResponseArrayOutput {
	return o.ApplyT(func(v IndexingPolicyResponse) []ExcludedPathResponse { return v.ExcludedPaths }).(ExcludedPathResponseArrayOutput)
}

func (o IndexingPolicyResponseOutput) IncludedPaths() IncludedPathResponseArrayOutput {
	return o.ApplyT(func(v IndexingPolicyResponse) []IncludedPathResponse { return v.IncludedPaths }).(IncludedPathResponseArrayOutput)
}

func (o IndexingPolicyResponseOutput) IndexingMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IndexingPolicyResponse) *string { return v.IndexingMode }).(pulumi.StringPtrOutput)
}

func (o IndexingPolicyResponseOutput) SpatialIndexes() SpatialSpecResponseArrayOutput {
	return o.ApplyT(func(v IndexingPolicyResponse) []SpatialSpecResponse { return v.SpatialIndexes }).(SpatialSpecResponseArrayOutput)
}

type IndexingPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (IndexingPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IndexingPolicyResponse)(nil)).Elem()
}

func (o IndexingPolicyResponsePtrOutput) ToIndexingPolicyResponsePtrOutput() IndexingPolicyResponsePtrOutput {
	return o
}

func (o IndexingPolicyResponsePtrOutput) ToIndexingPolicyResponsePtrOutputWithContext(ctx context.Context) IndexingPolicyResponsePtrOutput {
	return o
}

func (o IndexingPolicyResponsePtrOutput) Elem() IndexingPolicyResponseOutput {
	return o.ApplyT(func(v *IndexingPolicyResponse) IndexingPolicyResponse {
		if v != nil {
			return *v
		}
		var ret IndexingPolicyResponse
		return ret
	}).(IndexingPolicyResponseOutput)
}

func (o IndexingPolicyResponsePtrOutput) Automatic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IndexingPolicyResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Automatic
	}).(pulumi.BoolPtrOutput)
}

func (o IndexingPolicyResponsePtrOutput) CompositeIndexes() CompositePathResponseArrayArrayOutput {
	return o.ApplyT(func(v *IndexingPolicyResponse) [][]CompositePathResponse {
		if v == nil {
			return nil
		}
		return v.CompositeIndexes
	}).(CompositePathResponseArrayArrayOutput)
}

func (o IndexingPolicyResponsePtrOutput) ExcludedPaths() ExcludedPathResponseArrayOutput {
	return o.ApplyT(func(v *IndexingPolicyResponse) []ExcludedPathResponse {
		if v == nil {
			return nil
		}
		return v.ExcludedPaths
	}).(ExcludedPathResponseArrayOutput)
}

func (o IndexingPolicyResponsePtrOutput) IncludedPaths() IncludedPathResponseArrayOutput {
	return o.ApplyT(func(v *IndexingPolicyResponse) []IncludedPathResponse {
		if v == nil {
			return nil
		}
		return v.IncludedPaths
	}).(IncludedPathResponseArrayOutput)
}

func (o IndexingPolicyResponsePtrOutput) IndexingMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IndexingPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.IndexingMode
	}).(pulumi.StringPtrOutput)
}

func (o IndexingPolicyResponsePtrOutput) SpatialIndexes() SpatialSpecResponseArrayOutput {
	return o.ApplyT(func(v *IndexingPolicyResponse) []SpatialSpecResponse {
		if v == nil {
			return nil
		}
		return v.SpatialIndexes
	}).(SpatialSpecResponseArrayOutput)
}

type IpAddressOrRange struct {
	IpAddressOrRange *string `pulumi:"ipAddressOrRange"`
}





type IpAddressOrRangeInput interface {
	pulumi.Input

	ToIpAddressOrRangeOutput() IpAddressOrRangeOutput
	ToIpAddressOrRangeOutputWithContext(context.Context) IpAddressOrRangeOutput
}

type IpAddressOrRangeArgs struct {
	IpAddressOrRange pulumi.StringPtrInput `pulumi:"ipAddressOrRange"`
}

func (IpAddressOrRangeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IpAddressOrRange)(nil)).Elem()
}

func (i IpAddressOrRangeArgs) ToIpAddressOrRangeOutput() IpAddressOrRangeOutput {
	return i.ToIpAddressOrRangeOutputWithContext(context.Background())
}

func (i IpAddressOrRangeArgs) ToIpAddressOrRangeOutputWithContext(ctx context.Context) IpAddressOrRangeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpAddressOrRangeOutput)
}





type IpAddressOrRangeArrayInput interface {
	pulumi.Input

	ToIpAddressOrRangeArrayOutput() IpAddressOrRangeArrayOutput
	ToIpAddressOrRangeArrayOutputWithContext(context.Context) IpAddressOrRangeArrayOutput
}

type IpAddressOrRangeArray []IpAddressOrRangeInput

func (IpAddressOrRangeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpAddressOrRange)(nil)).Elem()
}

func (i IpAddressOrRangeArray) ToIpAddressOrRangeArrayOutput() IpAddressOrRangeArrayOutput {
	return i.ToIpAddressOrRangeArrayOutputWithContext(context.Background())
}

func (i IpAddressOrRangeArray) ToIpAddressOrRangeArrayOutputWithContext(ctx context.Context) IpAddressOrRangeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpAddressOrRangeArrayOutput)
}

type IpAddressOrRangeOutput struct{ *pulumi.OutputState }

func (IpAddressOrRangeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpAddressOrRange)(nil)).Elem()
}

func (o IpAddressOrRangeOutput) ToIpAddressOrRangeOutput() IpAddressOrRangeOutput {
	return o
}

func (o IpAddressOrRangeOutput) ToIpAddressOrRangeOutputWithContext(ctx context.Context) IpAddressOrRangeOutput {
	return o
}

func (o IpAddressOrRangeOutput) IpAddressOrRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpAddressOrRange) *string { return v.IpAddressOrRange }).(pulumi.StringPtrOutput)
}

type IpAddressOrRangeArrayOutput struct{ *pulumi.OutputState }

func (IpAddressOrRangeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpAddressOrRange)(nil)).Elem()
}

func (o IpAddressOrRangeArrayOutput) ToIpAddressOrRangeArrayOutput() IpAddressOrRangeArrayOutput {
	return o
}

func (o IpAddressOrRangeArrayOutput) ToIpAddressOrRangeArrayOutputWithContext(ctx context.Context) IpAddressOrRangeArrayOutput {
	return o
}

func (o IpAddressOrRangeArrayOutput) Index(i pulumi.IntInput) IpAddressOrRangeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IpAddressOrRange {
		return vs[0].([]IpAddressOrRange)[vs[1].(int)]
	}).(IpAddressOrRangeOutput)
}

type IpAddressOrRangeResponse struct {
	IpAddressOrRange *string `pulumi:"ipAddressOrRange"`
}

type IpAddressOrRangeResponseOutput struct{ *pulumi.OutputState }

func (IpAddressOrRangeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpAddressOrRangeResponse)(nil)).Elem()
}

func (o IpAddressOrRangeResponseOutput) ToIpAddressOrRangeResponseOutput() IpAddressOrRangeResponseOutput {
	return o
}

func (o IpAddressOrRangeResponseOutput) ToIpAddressOrRangeResponseOutputWithContext(ctx context.Context) IpAddressOrRangeResponseOutput {
	return o
}

func (o IpAddressOrRangeResponseOutput) IpAddressOrRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpAddressOrRangeResponse) *string { return v.IpAddressOrRange }).(pulumi.StringPtrOutput)
}

type IpAddressOrRangeResponseArrayOutput struct{ *pulumi.OutputState }

func (IpAddressOrRangeResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpAddressOrRangeResponse)(nil)).Elem()
}

func (o IpAddressOrRangeResponseArrayOutput) ToIpAddressOrRangeResponseArrayOutput() IpAddressOrRangeResponseArrayOutput {
	return o
}

func (o IpAddressOrRangeResponseArrayOutput) ToIpAddressOrRangeResponseArrayOutputWithContext(ctx context.Context) IpAddressOrRangeResponseArrayOutput {
	return o
}

func (o IpAddressOrRangeResponseArrayOutput) Index(i pulumi.IntInput) IpAddressOrRangeResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IpAddressOrRangeResponse {
		return vs[0].([]IpAddressOrRangeResponse)[vs[1].(int)]
	}).(IpAddressOrRangeResponseOutput)
}

type Location struct {
	FailoverPriority *int    `pulumi:"failoverPriority"`
	IsZoneRedundant  *bool   `pulumi:"isZoneRedundant"`
	LocationName     *string `pulumi:"locationName"`
}





type LocationInput interface {
	pulumi.Input

	ToLocationOutput() LocationOutput
	ToLocationOutputWithContext(context.Context) LocationOutput
}

type LocationArgs struct {
	FailoverPriority pulumi.IntPtrInput    `pulumi:"failoverPriority"`
	IsZoneRedundant  pulumi.BoolPtrInput   `pulumi:"isZoneRedundant"`
	LocationName     pulumi.StringPtrInput `pulumi:"locationName"`
}

func (LocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Location)(nil)).Elem()
}

func (i LocationArgs) ToLocationOutput() LocationOutput {
	return i.ToLocationOutputWithContext(context.Background())
}

func (i LocationArgs) ToLocationOutputWithContext(ctx context.Context) LocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationOutput)
}





type LocationArrayInput interface {
	pulumi.Input

	ToLocationArrayOutput() LocationArrayOutput
	ToLocationArrayOutputWithContext(context.Context) LocationArrayOutput
}

type LocationArray []LocationInput

func (LocationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Location)(nil)).Elem()
}

func (i LocationArray) ToLocationArrayOutput() LocationArrayOutput {
	return i.ToLocationArrayOutputWithContext(context.Background())
}

func (i LocationArray) ToLocationArrayOutputWithContext(ctx context.Context) LocationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationArrayOutput)
}

type LocationOutput struct{ *pulumi.OutputState }

func (LocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Location)(nil)).Elem()
}

func (o LocationOutput) ToLocationOutput() LocationOutput {
	return o
}

func (o LocationOutput) ToLocationOutputWithContext(ctx context.Context) LocationOutput {
	return o
}

func (o LocationOutput) FailoverPriority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Location) *int { return v.FailoverPriority }).(pulumi.IntPtrOutput)
}

func (o LocationOutput) IsZoneRedundant() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Location) *bool { return v.IsZoneRedundant }).(pulumi.BoolPtrOutput)
}

func (o LocationOutput) LocationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Location) *string { return v.LocationName }).(pulumi.StringPtrOutput)
}

type LocationArrayOutput struct{ *pulumi.OutputState }

func (LocationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Location)(nil)).Elem()
}

func (o LocationArrayOutput) ToLocationArrayOutput() LocationArrayOutput {
	return o
}

func (o LocationArrayOutput) ToLocationArrayOutputWithContext(ctx context.Context) LocationArrayOutput {
	return o
}

func (o LocationArrayOutput) Index(i pulumi.IntInput) LocationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Location {
		return vs[0].([]Location)[vs[1].(int)]
	}).(LocationOutput)
}

type LocationResponse struct {
	DocumentEndpoint  string  `pulumi:"documentEndpoint"`
	FailoverPriority  *int    `pulumi:"failoverPriority"`
	Id                string  `pulumi:"id"`
	IsZoneRedundant   *bool   `pulumi:"isZoneRedundant"`
	LocationName      *string `pulumi:"locationName"`
	ProvisioningState string  `pulumi:"provisioningState"`
}

type LocationResponseOutput struct{ *pulumi.OutputState }

func (LocationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LocationResponse)(nil)).Elem()
}

func (o LocationResponseOutput) ToLocationResponseOutput() LocationResponseOutput {
	return o
}

func (o LocationResponseOutput) ToLocationResponseOutputWithContext(ctx context.Context) LocationResponseOutput {
	return o
}

func (o LocationResponseOutput) DocumentEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LocationResponse) string { return v.DocumentEndpoint }).(pulumi.StringOutput)
}

func (o LocationResponseOutput) FailoverPriority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LocationResponse) *int { return v.FailoverPriority }).(pulumi.IntPtrOutput)
}

func (o LocationResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LocationResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o LocationResponseOutput) IsZoneRedundant() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LocationResponse) *bool { return v.IsZoneRedundant }).(pulumi.BoolPtrOutput)
}

func (o LocationResponseOutput) LocationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocationResponse) *string { return v.LocationName }).(pulumi.StringPtrOutput)
}

func (o LocationResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LocationResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type LocationResponseArrayOutput struct{ *pulumi.OutputState }

func (LocationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LocationResponse)(nil)).Elem()
}

func (o LocationResponseArrayOutput) ToLocationResponseArrayOutput() LocationResponseArrayOutput {
	return o
}

func (o LocationResponseArrayOutput) ToLocationResponseArrayOutputWithContext(ctx context.Context) LocationResponseArrayOutput {
	return o
}

func (o LocationResponseArrayOutput) Index(i pulumi.IntInput) LocationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LocationResponse {
		return vs[0].([]LocationResponse)[vs[1].(int)]
	}).(LocationResponseOutput)
}

type MongoDBCollectionGetPropertiesResponseOptions struct {
	AutoscaleSettings *AutoscaleSettingsResponse `pulumi:"autoscaleSettings"`
	Throughput        *int                       `pulumi:"throughput"`
}

type MongoDBCollectionGetPropertiesResponseOptionsOutput struct{ *pulumi.OutputState }

func (MongoDBCollectionGetPropertiesResponseOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoDBCollectionGetPropertiesResponseOptions)(nil)).Elem()
}

func (o MongoDBCollectionGetPropertiesResponseOptionsOutput) ToMongoDBCollectionGetPropertiesResponseOptionsOutput() MongoDBCollectionGetPropertiesResponseOptionsOutput {
	return o
}

func (o MongoDBCollectionGetPropertiesResponseOptionsOutput) ToMongoDBCollectionGetPropertiesResponseOptionsOutputWithContext(ctx context.Context) MongoDBCollectionGetPropertiesResponseOptionsOutput {
	return o
}

func (o MongoDBCollectionGetPropertiesResponseOptionsOutput) AutoscaleSettings() AutoscaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v MongoDBCollectionGetPropertiesResponseOptions) *AutoscaleSettingsResponse {
		return v.AutoscaleSettings
	}).(AutoscaleSettingsResponsePtrOutput)
}

func (o MongoDBCollectionGetPropertiesResponseOptionsOutput) Throughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MongoDBCollectionGetPropertiesResponseOptions) *int { return v.Throughput }).(pulumi.IntPtrOutput)
}

type MongoDBCollectionGetPropertiesResponseOptionsPtrOutput struct{ *pulumi.OutputState }

func (MongoDBCollectionGetPropertiesResponseOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoDBCollectionGetPropertiesResponseOptions)(nil)).Elem()
}

func (o MongoDBCollectionGetPropertiesResponseOptionsPtrOutput) ToMongoDBCollectionGetPropertiesResponseOptionsPtrOutput() MongoDBCollectionGetPropertiesResponseOptionsPtrOutput {
	return o
}

func (o MongoDBCollectionGetPropertiesResponseOptionsPtrOutput) ToMongoDBCollectionGetPropertiesResponseOptionsPtrOutputWithContext(ctx context.Context) MongoDBCollectionGetPropertiesResponseOptionsPtrOutput {
	return o
}

func (o MongoDBCollectionGetPropertiesResponseOptionsPtrOutput) Elem() MongoDBCollectionGetPropertiesResponseOptionsOutput {
	return o.ApplyT(func(v *MongoDBCollectionGetPropertiesResponseOptions) MongoDBCollectionGetPropertiesResponseOptions {
		if v != nil {
			return *v
		}
		var ret MongoDBCollectionGetPropertiesResponseOptions
		return ret
	}).(MongoDBCollectionGetPropertiesResponseOptionsOutput)
}

func (o MongoDBCollectionGetPropertiesResponseOptionsPtrOutput) AutoscaleSettings() AutoscaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v *MongoDBCollectionGetPropertiesResponseOptions) *AutoscaleSettingsResponse {
		if v == nil {
			return nil
		}
		return v.AutoscaleSettings
	}).(AutoscaleSettingsResponsePtrOutput)
}

func (o MongoDBCollectionGetPropertiesResponseOptionsPtrOutput) Throughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MongoDBCollectionGetPropertiesResponseOptions) *int {
		if v == nil {
			return nil
		}
		return v.Throughput
	}).(pulumi.IntPtrOutput)
}

type MongoDBCollectionGetPropertiesResponseResource struct {
	AnalyticalStorageTtl *int                 `pulumi:"analyticalStorageTtl"`
	Etag                 string               `pulumi:"etag"`
	Id                   string               `pulumi:"id"`
	Indexes              []MongoIndexResponse `pulumi:"indexes"`
	Rid                  string               `pulumi:"rid"`
	ShardKey             map[string]string    `pulumi:"shardKey"`
	Ts                   float64              `pulumi:"ts"`
}

type MongoDBCollectionGetPropertiesResponseResourceOutput struct{ *pulumi.OutputState }

func (MongoDBCollectionGetPropertiesResponseResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoDBCollectionGetPropertiesResponseResource)(nil)).Elem()
}

func (o MongoDBCollectionGetPropertiesResponseResourceOutput) ToMongoDBCollectionGetPropertiesResponseResourceOutput() MongoDBCollectionGetPropertiesResponseResourceOutput {
	return o
}

func (o MongoDBCollectionGetPropertiesResponseResourceOutput) ToMongoDBCollectionGetPropertiesResponseResourceOutputWithContext(ctx context.Context) MongoDBCollectionGetPropertiesResponseResourceOutput {
	return o
}

func (o MongoDBCollectionGetPropertiesResponseResourceOutput) AnalyticalStorageTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MongoDBCollectionGetPropertiesResponseResource) *int { return v.AnalyticalStorageTtl }).(pulumi.IntPtrOutput)
}

func (o MongoDBCollectionGetPropertiesResponseResourceOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v MongoDBCollectionGetPropertiesResponseResource) string { return v.Etag }).(pulumi.StringOutput)
}

func (o MongoDBCollectionGetPropertiesResponseResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v MongoDBCollectionGetPropertiesResponseResource) string { return v.Id }).(pulumi.StringOutput)
}

func (o MongoDBCollectionGetPropertiesResponseResourceOutput) Indexes() MongoIndexResponseArrayOutput {
	return o.ApplyT(func(v MongoDBCollectionGetPropertiesResponseResource) []MongoIndexResponse { return v.Indexes }).(MongoIndexResponseArrayOutput)
}

func (o MongoDBCollectionGetPropertiesResponseResourceOutput) Rid() pulumi.StringOutput {
	return o.ApplyT(func(v MongoDBCollectionGetPropertiesResponseResource) string { return v.Rid }).(pulumi.StringOutput)
}

func (o MongoDBCollectionGetPropertiesResponseResourceOutput) ShardKey() pulumi.StringMapOutput {
	return o.ApplyT(func(v MongoDBCollectionGetPropertiesResponseResource) map[string]string { return v.ShardKey }).(pulumi.StringMapOutput)
}

func (o MongoDBCollectionGetPropertiesResponseResourceOutput) Ts() pulumi.Float64Output {
	return o.ApplyT(func(v MongoDBCollectionGetPropertiesResponseResource) float64 { return v.Ts }).(pulumi.Float64Output)
}

type MongoDBCollectionGetPropertiesResponseResourcePtrOutput struct{ *pulumi.OutputState }

func (MongoDBCollectionGetPropertiesResponseResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoDBCollectionGetPropertiesResponseResource)(nil)).Elem()
}

func (o MongoDBCollectionGetPropertiesResponseResourcePtrOutput) ToMongoDBCollectionGetPropertiesResponseResourcePtrOutput() MongoDBCollectionGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o MongoDBCollectionGetPropertiesResponseResourcePtrOutput) ToMongoDBCollectionGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) MongoDBCollectionGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o MongoDBCollectionGetPropertiesResponseResourcePtrOutput) Elem() MongoDBCollectionGetPropertiesResponseResourceOutput {
	return o.ApplyT(func(v *MongoDBCollectionGetPropertiesResponseResource) MongoDBCollectionGetPropertiesResponseResource {
		if v != nil {
			return *v
		}
		var ret MongoDBCollectionGetPropertiesResponseResource
		return ret
	}).(MongoDBCollectionGetPropertiesResponseResourceOutput)
}

func (o MongoDBCollectionGetPropertiesResponseResourcePtrOutput) AnalyticalStorageTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MongoDBCollectionGetPropertiesResponseResource) *int {
		if v == nil {
			return nil
		}
		return v.AnalyticalStorageTtl
	}).(pulumi.IntPtrOutput)
}

func (o MongoDBCollectionGetPropertiesResponseResourcePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MongoDBCollectionGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o MongoDBCollectionGetPropertiesResponseResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MongoDBCollectionGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o MongoDBCollectionGetPropertiesResponseResourcePtrOutput) Indexes() MongoIndexResponseArrayOutput {
	return o.ApplyT(func(v *MongoDBCollectionGetPropertiesResponseResource) []MongoIndexResponse {
		if v == nil {
			return nil
		}
		return v.Indexes
	}).(MongoIndexResponseArrayOutput)
}

func (o MongoDBCollectionGetPropertiesResponseResourcePtrOutput) Rid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MongoDBCollectionGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Rid
	}).(pulumi.StringPtrOutput)
}

func (o MongoDBCollectionGetPropertiesResponseResourcePtrOutput) ShardKey() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MongoDBCollectionGetPropertiesResponseResource) map[string]string {
		if v == nil {
			return nil
		}
		return v.ShardKey
	}).(pulumi.StringMapOutput)
}

func (o MongoDBCollectionGetPropertiesResponseResourcePtrOutput) Ts() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *MongoDBCollectionGetPropertiesResponseResource) *float64 {
		if v == nil {
			return nil
		}
		return &v.Ts
	}).(pulumi.Float64PtrOutput)
}

type MongoDBCollectionResource struct {
	AnalyticalStorageTtl *int              `pulumi:"analyticalStorageTtl"`
	Id                   string            `pulumi:"id"`
	Indexes              []MongoIndex      `pulumi:"indexes"`
	ShardKey             map[string]string `pulumi:"shardKey"`
}





type MongoDBCollectionResourceInput interface {
	pulumi.Input

	ToMongoDBCollectionResourceOutput() MongoDBCollectionResourceOutput
	ToMongoDBCollectionResourceOutputWithContext(context.Context) MongoDBCollectionResourceOutput
}

type MongoDBCollectionResourceArgs struct {
	AnalyticalStorageTtl pulumi.IntPtrInput    `pulumi:"analyticalStorageTtl"`
	Id                   pulumi.StringInput    `pulumi:"id"`
	Indexes              MongoIndexArrayInput  `pulumi:"indexes"`
	ShardKey             pulumi.StringMapInput `pulumi:"shardKey"`
}

func (MongoDBCollectionResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoDBCollectionResource)(nil)).Elem()
}

func (i MongoDBCollectionResourceArgs) ToMongoDBCollectionResourceOutput() MongoDBCollectionResourceOutput {
	return i.ToMongoDBCollectionResourceOutputWithContext(context.Background())
}

func (i MongoDBCollectionResourceArgs) ToMongoDBCollectionResourceOutputWithContext(ctx context.Context) MongoDBCollectionResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoDBCollectionResourceOutput)
}

type MongoDBCollectionResourceOutput struct{ *pulumi.OutputState }

func (MongoDBCollectionResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoDBCollectionResource)(nil)).Elem()
}

func (o MongoDBCollectionResourceOutput) ToMongoDBCollectionResourceOutput() MongoDBCollectionResourceOutput {
	return o
}

func (o MongoDBCollectionResourceOutput) ToMongoDBCollectionResourceOutputWithContext(ctx context.Context) MongoDBCollectionResourceOutput {
	return o
}

func (o MongoDBCollectionResourceOutput) AnalyticalStorageTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MongoDBCollectionResource) *int { return v.AnalyticalStorageTtl }).(pulumi.IntPtrOutput)
}

func (o MongoDBCollectionResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v MongoDBCollectionResource) string { return v.Id }).(pulumi.StringOutput)
}

func (o MongoDBCollectionResourceOutput) Indexes() MongoIndexArrayOutput {
	return o.ApplyT(func(v MongoDBCollectionResource) []MongoIndex { return v.Indexes }).(MongoIndexArrayOutput)
}

func (o MongoDBCollectionResourceOutput) ShardKey() pulumi.StringMapOutput {
	return o.ApplyT(func(v MongoDBCollectionResource) map[string]string { return v.ShardKey }).(pulumi.StringMapOutput)
}

type MongoDBDatabaseGetPropertiesResponseOptions struct {
	AutoscaleSettings *AutoscaleSettingsResponse `pulumi:"autoscaleSettings"`
	Throughput        *int                       `pulumi:"throughput"`
}

type MongoDBDatabaseGetPropertiesResponseOptionsOutput struct{ *pulumi.OutputState }

func (MongoDBDatabaseGetPropertiesResponseOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoDBDatabaseGetPropertiesResponseOptions)(nil)).Elem()
}

func (o MongoDBDatabaseGetPropertiesResponseOptionsOutput) ToMongoDBDatabaseGetPropertiesResponseOptionsOutput() MongoDBDatabaseGetPropertiesResponseOptionsOutput {
	return o
}

func (o MongoDBDatabaseGetPropertiesResponseOptionsOutput) ToMongoDBDatabaseGetPropertiesResponseOptionsOutputWithContext(ctx context.Context) MongoDBDatabaseGetPropertiesResponseOptionsOutput {
	return o
}

func (o MongoDBDatabaseGetPropertiesResponseOptionsOutput) AutoscaleSettings() AutoscaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v MongoDBDatabaseGetPropertiesResponseOptions) *AutoscaleSettingsResponse {
		return v.AutoscaleSettings
	}).(AutoscaleSettingsResponsePtrOutput)
}

func (o MongoDBDatabaseGetPropertiesResponseOptionsOutput) Throughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MongoDBDatabaseGetPropertiesResponseOptions) *int { return v.Throughput }).(pulumi.IntPtrOutput)
}

type MongoDBDatabaseGetPropertiesResponseOptionsPtrOutput struct{ *pulumi.OutputState }

func (MongoDBDatabaseGetPropertiesResponseOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoDBDatabaseGetPropertiesResponseOptions)(nil)).Elem()
}

func (o MongoDBDatabaseGetPropertiesResponseOptionsPtrOutput) ToMongoDBDatabaseGetPropertiesResponseOptionsPtrOutput() MongoDBDatabaseGetPropertiesResponseOptionsPtrOutput {
	return o
}

func (o MongoDBDatabaseGetPropertiesResponseOptionsPtrOutput) ToMongoDBDatabaseGetPropertiesResponseOptionsPtrOutputWithContext(ctx context.Context) MongoDBDatabaseGetPropertiesResponseOptionsPtrOutput {
	return o
}

func (o MongoDBDatabaseGetPropertiesResponseOptionsPtrOutput) Elem() MongoDBDatabaseGetPropertiesResponseOptionsOutput {
	return o.ApplyT(func(v *MongoDBDatabaseGetPropertiesResponseOptions) MongoDBDatabaseGetPropertiesResponseOptions {
		if v != nil {
			return *v
		}
		var ret MongoDBDatabaseGetPropertiesResponseOptions
		return ret
	}).(MongoDBDatabaseGetPropertiesResponseOptionsOutput)
}

func (o MongoDBDatabaseGetPropertiesResponseOptionsPtrOutput) AutoscaleSettings() AutoscaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v *MongoDBDatabaseGetPropertiesResponseOptions) *AutoscaleSettingsResponse {
		if v == nil {
			return nil
		}
		return v.AutoscaleSettings
	}).(AutoscaleSettingsResponsePtrOutput)
}

func (o MongoDBDatabaseGetPropertiesResponseOptionsPtrOutput) Throughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MongoDBDatabaseGetPropertiesResponseOptions) *int {
		if v == nil {
			return nil
		}
		return v.Throughput
	}).(pulumi.IntPtrOutput)
}

type MongoDBDatabaseGetPropertiesResponseResource struct {
	Etag string  `pulumi:"etag"`
	Id   string  `pulumi:"id"`
	Rid  string  `pulumi:"rid"`
	Ts   float64 `pulumi:"ts"`
}

type MongoDBDatabaseGetPropertiesResponseResourceOutput struct{ *pulumi.OutputState }

func (MongoDBDatabaseGetPropertiesResponseResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoDBDatabaseGetPropertiesResponseResource)(nil)).Elem()
}

func (o MongoDBDatabaseGetPropertiesResponseResourceOutput) ToMongoDBDatabaseGetPropertiesResponseResourceOutput() MongoDBDatabaseGetPropertiesResponseResourceOutput {
	return o
}

func (o MongoDBDatabaseGetPropertiesResponseResourceOutput) ToMongoDBDatabaseGetPropertiesResponseResourceOutputWithContext(ctx context.Context) MongoDBDatabaseGetPropertiesResponseResourceOutput {
	return o
}

func (o MongoDBDatabaseGetPropertiesResponseResourceOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v MongoDBDatabaseGetPropertiesResponseResource) string { return v.Etag }).(pulumi.StringOutput)
}

func (o MongoDBDatabaseGetPropertiesResponseResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v MongoDBDatabaseGetPropertiesResponseResource) string { return v.Id }).(pulumi.StringOutput)
}

func (o MongoDBDatabaseGetPropertiesResponseResourceOutput) Rid() pulumi.StringOutput {
	return o.ApplyT(func(v MongoDBDatabaseGetPropertiesResponseResource) string { return v.Rid }).(pulumi.StringOutput)
}

func (o MongoDBDatabaseGetPropertiesResponseResourceOutput) Ts() pulumi.Float64Output {
	return o.ApplyT(func(v MongoDBDatabaseGetPropertiesResponseResource) float64 { return v.Ts }).(pulumi.Float64Output)
}

type MongoDBDatabaseGetPropertiesResponseResourcePtrOutput struct{ *pulumi.OutputState }

func (MongoDBDatabaseGetPropertiesResponseResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoDBDatabaseGetPropertiesResponseResource)(nil)).Elem()
}

func (o MongoDBDatabaseGetPropertiesResponseResourcePtrOutput) ToMongoDBDatabaseGetPropertiesResponseResourcePtrOutput() MongoDBDatabaseGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o MongoDBDatabaseGetPropertiesResponseResourcePtrOutput) ToMongoDBDatabaseGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) MongoDBDatabaseGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o MongoDBDatabaseGetPropertiesResponseResourcePtrOutput) Elem() MongoDBDatabaseGetPropertiesResponseResourceOutput {
	return o.ApplyT(func(v *MongoDBDatabaseGetPropertiesResponseResource) MongoDBDatabaseGetPropertiesResponseResource {
		if v != nil {
			return *v
		}
		var ret MongoDBDatabaseGetPropertiesResponseResource
		return ret
	}).(MongoDBDatabaseGetPropertiesResponseResourceOutput)
}

func (o MongoDBDatabaseGetPropertiesResponseResourcePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MongoDBDatabaseGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o MongoDBDatabaseGetPropertiesResponseResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MongoDBDatabaseGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o MongoDBDatabaseGetPropertiesResponseResourcePtrOutput) Rid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MongoDBDatabaseGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Rid
	}).(pulumi.StringPtrOutput)
}

func (o MongoDBDatabaseGetPropertiesResponseResourcePtrOutput) Ts() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *MongoDBDatabaseGetPropertiesResponseResource) *float64 {
		if v == nil {
			return nil
		}
		return &v.Ts
	}).(pulumi.Float64PtrOutput)
}

type MongoDBDatabaseResource struct {
	Id string `pulumi:"id"`
}





type MongoDBDatabaseResourceInput interface {
	pulumi.Input

	ToMongoDBDatabaseResourceOutput() MongoDBDatabaseResourceOutput
	ToMongoDBDatabaseResourceOutputWithContext(context.Context) MongoDBDatabaseResourceOutput
}

type MongoDBDatabaseResourceArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (MongoDBDatabaseResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoDBDatabaseResource)(nil)).Elem()
}

func (i MongoDBDatabaseResourceArgs) ToMongoDBDatabaseResourceOutput() MongoDBDatabaseResourceOutput {
	return i.ToMongoDBDatabaseResourceOutputWithContext(context.Background())
}

func (i MongoDBDatabaseResourceArgs) ToMongoDBDatabaseResourceOutputWithContext(ctx context.Context) MongoDBDatabaseResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoDBDatabaseResourceOutput)
}

type MongoDBDatabaseResourceOutput struct{ *pulumi.OutputState }

func (MongoDBDatabaseResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoDBDatabaseResource)(nil)).Elem()
}

func (o MongoDBDatabaseResourceOutput) ToMongoDBDatabaseResourceOutput() MongoDBDatabaseResourceOutput {
	return o
}

func (o MongoDBDatabaseResourceOutput) ToMongoDBDatabaseResourceOutputWithContext(ctx context.Context) MongoDBDatabaseResourceOutput {
	return o
}

func (o MongoDBDatabaseResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v MongoDBDatabaseResource) string { return v.Id }).(pulumi.StringOutput)
}

type MongoIndex struct {
	Key     *MongoIndexKeys    `pulumi:"key"`
	Options *MongoIndexOptions `pulumi:"options"`
}





type MongoIndexInput interface {
	pulumi.Input

	ToMongoIndexOutput() MongoIndexOutput
	ToMongoIndexOutputWithContext(context.Context) MongoIndexOutput
}

type MongoIndexArgs struct {
	Key     MongoIndexKeysPtrInput    `pulumi:"key"`
	Options MongoIndexOptionsPtrInput `pulumi:"options"`
}

func (MongoIndexArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoIndex)(nil)).Elem()
}

func (i MongoIndexArgs) ToMongoIndexOutput() MongoIndexOutput {
	return i.ToMongoIndexOutputWithContext(context.Background())
}

func (i MongoIndexArgs) ToMongoIndexOutputWithContext(ctx context.Context) MongoIndexOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoIndexOutput)
}





type MongoIndexArrayInput interface {
	pulumi.Input

	ToMongoIndexArrayOutput() MongoIndexArrayOutput
	ToMongoIndexArrayOutputWithContext(context.Context) MongoIndexArrayOutput
}

type MongoIndexArray []MongoIndexInput

func (MongoIndexArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MongoIndex)(nil)).Elem()
}

func (i MongoIndexArray) ToMongoIndexArrayOutput() MongoIndexArrayOutput {
	return i.ToMongoIndexArrayOutputWithContext(context.Background())
}

func (i MongoIndexArray) ToMongoIndexArrayOutputWithContext(ctx context.Context) MongoIndexArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoIndexArrayOutput)
}

type MongoIndexOutput struct{ *pulumi.OutputState }

func (MongoIndexOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoIndex)(nil)).Elem()
}

func (o MongoIndexOutput) ToMongoIndexOutput() MongoIndexOutput {
	return o
}

func (o MongoIndexOutput) ToMongoIndexOutputWithContext(ctx context.Context) MongoIndexOutput {
	return o
}

func (o MongoIndexOutput) Key() MongoIndexKeysPtrOutput {
	return o.ApplyT(func(v MongoIndex) *MongoIndexKeys { return v.Key }).(MongoIndexKeysPtrOutput)
}

func (o MongoIndexOutput) Options() MongoIndexOptionsPtrOutput {
	return o.ApplyT(func(v MongoIndex) *MongoIndexOptions { return v.Options }).(MongoIndexOptionsPtrOutput)
}

type MongoIndexArrayOutput struct{ *pulumi.OutputState }

func (MongoIndexArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MongoIndex)(nil)).Elem()
}

func (o MongoIndexArrayOutput) ToMongoIndexArrayOutput() MongoIndexArrayOutput {
	return o
}

func (o MongoIndexArrayOutput) ToMongoIndexArrayOutputWithContext(ctx context.Context) MongoIndexArrayOutput {
	return o
}

func (o MongoIndexArrayOutput) Index(i pulumi.IntInput) MongoIndexOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MongoIndex {
		return vs[0].([]MongoIndex)[vs[1].(int)]
	}).(MongoIndexOutput)
}

type MongoIndexKeys struct {
	Keys []string `pulumi:"keys"`
}





type MongoIndexKeysInput interface {
	pulumi.Input

	ToMongoIndexKeysOutput() MongoIndexKeysOutput
	ToMongoIndexKeysOutputWithContext(context.Context) MongoIndexKeysOutput
}

type MongoIndexKeysArgs struct {
	Keys pulumi.StringArrayInput `pulumi:"keys"`
}

func (MongoIndexKeysArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoIndexKeys)(nil)).Elem()
}

func (i MongoIndexKeysArgs) ToMongoIndexKeysOutput() MongoIndexKeysOutput {
	return i.ToMongoIndexKeysOutputWithContext(context.Background())
}

func (i MongoIndexKeysArgs) ToMongoIndexKeysOutputWithContext(ctx context.Context) MongoIndexKeysOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoIndexKeysOutput)
}

func (i MongoIndexKeysArgs) ToMongoIndexKeysPtrOutput() MongoIndexKeysPtrOutput {
	return i.ToMongoIndexKeysPtrOutputWithContext(context.Background())
}

func (i MongoIndexKeysArgs) ToMongoIndexKeysPtrOutputWithContext(ctx context.Context) MongoIndexKeysPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoIndexKeysOutput).ToMongoIndexKeysPtrOutputWithContext(ctx)
}









type MongoIndexKeysPtrInput interface {
	pulumi.Input

	ToMongoIndexKeysPtrOutput() MongoIndexKeysPtrOutput
	ToMongoIndexKeysPtrOutputWithContext(context.Context) MongoIndexKeysPtrOutput
}

type mongoIndexKeysPtrType MongoIndexKeysArgs

func MongoIndexKeysPtr(v *MongoIndexKeysArgs) MongoIndexKeysPtrInput {
	return (*mongoIndexKeysPtrType)(v)
}

func (*mongoIndexKeysPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoIndexKeys)(nil)).Elem()
}

func (i *mongoIndexKeysPtrType) ToMongoIndexKeysPtrOutput() MongoIndexKeysPtrOutput {
	return i.ToMongoIndexKeysPtrOutputWithContext(context.Background())
}

func (i *mongoIndexKeysPtrType) ToMongoIndexKeysPtrOutputWithContext(ctx context.Context) MongoIndexKeysPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoIndexKeysPtrOutput)
}

type MongoIndexKeysOutput struct{ *pulumi.OutputState }

func (MongoIndexKeysOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoIndexKeys)(nil)).Elem()
}

func (o MongoIndexKeysOutput) ToMongoIndexKeysOutput() MongoIndexKeysOutput {
	return o
}

func (o MongoIndexKeysOutput) ToMongoIndexKeysOutputWithContext(ctx context.Context) MongoIndexKeysOutput {
	return o
}

func (o MongoIndexKeysOutput) ToMongoIndexKeysPtrOutput() MongoIndexKeysPtrOutput {
	return o.ToMongoIndexKeysPtrOutputWithContext(context.Background())
}

func (o MongoIndexKeysOutput) ToMongoIndexKeysPtrOutputWithContext(ctx context.Context) MongoIndexKeysPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MongoIndexKeys) *MongoIndexKeys {
		return &v
	}).(MongoIndexKeysPtrOutput)
}

func (o MongoIndexKeysOutput) Keys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v MongoIndexKeys) []string { return v.Keys }).(pulumi.StringArrayOutput)
}

type MongoIndexKeysPtrOutput struct{ *pulumi.OutputState }

func (MongoIndexKeysPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoIndexKeys)(nil)).Elem()
}

func (o MongoIndexKeysPtrOutput) ToMongoIndexKeysPtrOutput() MongoIndexKeysPtrOutput {
	return o
}

func (o MongoIndexKeysPtrOutput) ToMongoIndexKeysPtrOutputWithContext(ctx context.Context) MongoIndexKeysPtrOutput {
	return o
}

func (o MongoIndexKeysPtrOutput) Elem() MongoIndexKeysOutput {
	return o.ApplyT(func(v *MongoIndexKeys) MongoIndexKeys {
		if v != nil {
			return *v
		}
		var ret MongoIndexKeys
		return ret
	}).(MongoIndexKeysOutput)
}

func (o MongoIndexKeysPtrOutput) Keys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MongoIndexKeys) []string {
		if v == nil {
			return nil
		}
		return v.Keys
	}).(pulumi.StringArrayOutput)
}

type MongoIndexKeysResponse struct {
	Keys []string `pulumi:"keys"`
}

type MongoIndexKeysResponseOutput struct{ *pulumi.OutputState }

func (MongoIndexKeysResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoIndexKeysResponse)(nil)).Elem()
}

func (o MongoIndexKeysResponseOutput) ToMongoIndexKeysResponseOutput() MongoIndexKeysResponseOutput {
	return o
}

func (o MongoIndexKeysResponseOutput) ToMongoIndexKeysResponseOutputWithContext(ctx context.Context) MongoIndexKeysResponseOutput {
	return o
}

func (o MongoIndexKeysResponseOutput) Keys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v MongoIndexKeysResponse) []string { return v.Keys }).(pulumi.StringArrayOutput)
}

type MongoIndexKeysResponsePtrOutput struct{ *pulumi.OutputState }

func (MongoIndexKeysResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoIndexKeysResponse)(nil)).Elem()
}

func (o MongoIndexKeysResponsePtrOutput) ToMongoIndexKeysResponsePtrOutput() MongoIndexKeysResponsePtrOutput {
	return o
}

func (o MongoIndexKeysResponsePtrOutput) ToMongoIndexKeysResponsePtrOutputWithContext(ctx context.Context) MongoIndexKeysResponsePtrOutput {
	return o
}

func (o MongoIndexKeysResponsePtrOutput) Elem() MongoIndexKeysResponseOutput {
	return o.ApplyT(func(v *MongoIndexKeysResponse) MongoIndexKeysResponse {
		if v != nil {
			return *v
		}
		var ret MongoIndexKeysResponse
		return ret
	}).(MongoIndexKeysResponseOutput)
}

func (o MongoIndexKeysResponsePtrOutput) Keys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MongoIndexKeysResponse) []string {
		if v == nil {
			return nil
		}
		return v.Keys
	}).(pulumi.StringArrayOutput)
}

type MongoIndexOptions struct {
	ExpireAfterSeconds *int  `pulumi:"expireAfterSeconds"`
	Unique             *bool `pulumi:"unique"`
}





type MongoIndexOptionsInput interface {
	pulumi.Input

	ToMongoIndexOptionsOutput() MongoIndexOptionsOutput
	ToMongoIndexOptionsOutputWithContext(context.Context) MongoIndexOptionsOutput
}

type MongoIndexOptionsArgs struct {
	ExpireAfterSeconds pulumi.IntPtrInput  `pulumi:"expireAfterSeconds"`
	Unique             pulumi.BoolPtrInput `pulumi:"unique"`
}

func (MongoIndexOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoIndexOptions)(nil)).Elem()
}

func (i MongoIndexOptionsArgs) ToMongoIndexOptionsOutput() MongoIndexOptionsOutput {
	return i.ToMongoIndexOptionsOutputWithContext(context.Background())
}

func (i MongoIndexOptionsArgs) ToMongoIndexOptionsOutputWithContext(ctx context.Context) MongoIndexOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoIndexOptionsOutput)
}

func (i MongoIndexOptionsArgs) ToMongoIndexOptionsPtrOutput() MongoIndexOptionsPtrOutput {
	return i.ToMongoIndexOptionsPtrOutputWithContext(context.Background())
}

func (i MongoIndexOptionsArgs) ToMongoIndexOptionsPtrOutputWithContext(ctx context.Context) MongoIndexOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoIndexOptionsOutput).ToMongoIndexOptionsPtrOutputWithContext(ctx)
}









type MongoIndexOptionsPtrInput interface {
	pulumi.Input

	ToMongoIndexOptionsPtrOutput() MongoIndexOptionsPtrOutput
	ToMongoIndexOptionsPtrOutputWithContext(context.Context) MongoIndexOptionsPtrOutput
}

type mongoIndexOptionsPtrType MongoIndexOptionsArgs

func MongoIndexOptionsPtr(v *MongoIndexOptionsArgs) MongoIndexOptionsPtrInput {
	return (*mongoIndexOptionsPtrType)(v)
}

func (*mongoIndexOptionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoIndexOptions)(nil)).Elem()
}

func (i *mongoIndexOptionsPtrType) ToMongoIndexOptionsPtrOutput() MongoIndexOptionsPtrOutput {
	return i.ToMongoIndexOptionsPtrOutputWithContext(context.Background())
}

func (i *mongoIndexOptionsPtrType) ToMongoIndexOptionsPtrOutputWithContext(ctx context.Context) MongoIndexOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoIndexOptionsPtrOutput)
}

type MongoIndexOptionsOutput struct{ *pulumi.OutputState }

func (MongoIndexOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoIndexOptions)(nil)).Elem()
}

func (o MongoIndexOptionsOutput) ToMongoIndexOptionsOutput() MongoIndexOptionsOutput {
	return o
}

func (o MongoIndexOptionsOutput) ToMongoIndexOptionsOutputWithContext(ctx context.Context) MongoIndexOptionsOutput {
	return o
}

func (o MongoIndexOptionsOutput) ToMongoIndexOptionsPtrOutput() MongoIndexOptionsPtrOutput {
	return o.ToMongoIndexOptionsPtrOutputWithContext(context.Background())
}

func (o MongoIndexOptionsOutput) ToMongoIndexOptionsPtrOutputWithContext(ctx context.Context) MongoIndexOptionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MongoIndexOptions) *MongoIndexOptions {
		return &v
	}).(MongoIndexOptionsPtrOutput)
}

func (o MongoIndexOptionsOutput) ExpireAfterSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MongoIndexOptions) *int { return v.ExpireAfterSeconds }).(pulumi.IntPtrOutput)
}

func (o MongoIndexOptionsOutput) Unique() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MongoIndexOptions) *bool { return v.Unique }).(pulumi.BoolPtrOutput)
}

type MongoIndexOptionsPtrOutput struct{ *pulumi.OutputState }

func (MongoIndexOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoIndexOptions)(nil)).Elem()
}

func (o MongoIndexOptionsPtrOutput) ToMongoIndexOptionsPtrOutput() MongoIndexOptionsPtrOutput {
	return o
}

func (o MongoIndexOptionsPtrOutput) ToMongoIndexOptionsPtrOutputWithContext(ctx context.Context) MongoIndexOptionsPtrOutput {
	return o
}

func (o MongoIndexOptionsPtrOutput) Elem() MongoIndexOptionsOutput {
	return o.ApplyT(func(v *MongoIndexOptions) MongoIndexOptions {
		if v != nil {
			return *v
		}
		var ret MongoIndexOptions
		return ret
	}).(MongoIndexOptionsOutput)
}

func (o MongoIndexOptionsPtrOutput) ExpireAfterSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MongoIndexOptions) *int {
		if v == nil {
			return nil
		}
		return v.ExpireAfterSeconds
	}).(pulumi.IntPtrOutput)
}

func (o MongoIndexOptionsPtrOutput) Unique() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MongoIndexOptions) *bool {
		if v == nil {
			return nil
		}
		return v.Unique
	}).(pulumi.BoolPtrOutput)
}

type MongoIndexOptionsResponse struct {
	ExpireAfterSeconds *int  `pulumi:"expireAfterSeconds"`
	Unique             *bool `pulumi:"unique"`
}

type MongoIndexOptionsResponseOutput struct{ *pulumi.OutputState }

func (MongoIndexOptionsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoIndexOptionsResponse)(nil)).Elem()
}

func (o MongoIndexOptionsResponseOutput) ToMongoIndexOptionsResponseOutput() MongoIndexOptionsResponseOutput {
	return o
}

func (o MongoIndexOptionsResponseOutput) ToMongoIndexOptionsResponseOutputWithContext(ctx context.Context) MongoIndexOptionsResponseOutput {
	return o
}

func (o MongoIndexOptionsResponseOutput) ExpireAfterSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MongoIndexOptionsResponse) *int { return v.ExpireAfterSeconds }).(pulumi.IntPtrOutput)
}

func (o MongoIndexOptionsResponseOutput) Unique() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MongoIndexOptionsResponse) *bool { return v.Unique }).(pulumi.BoolPtrOutput)
}

type MongoIndexOptionsResponsePtrOutput struct{ *pulumi.OutputState }

func (MongoIndexOptionsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoIndexOptionsResponse)(nil)).Elem()
}

func (o MongoIndexOptionsResponsePtrOutput) ToMongoIndexOptionsResponsePtrOutput() MongoIndexOptionsResponsePtrOutput {
	return o
}

func (o MongoIndexOptionsResponsePtrOutput) ToMongoIndexOptionsResponsePtrOutputWithContext(ctx context.Context) MongoIndexOptionsResponsePtrOutput {
	return o
}

func (o MongoIndexOptionsResponsePtrOutput) Elem() MongoIndexOptionsResponseOutput {
	return o.ApplyT(func(v *MongoIndexOptionsResponse) MongoIndexOptionsResponse {
		if v != nil {
			return *v
		}
		var ret MongoIndexOptionsResponse
		return ret
	}).(MongoIndexOptionsResponseOutput)
}

func (o MongoIndexOptionsResponsePtrOutput) ExpireAfterSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MongoIndexOptionsResponse) *int {
		if v == nil {
			return nil
		}
		return v.ExpireAfterSeconds
	}).(pulumi.IntPtrOutput)
}

func (o MongoIndexOptionsResponsePtrOutput) Unique() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MongoIndexOptionsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Unique
	}).(pulumi.BoolPtrOutput)
}

type MongoIndexResponse struct {
	Key     *MongoIndexKeysResponse    `pulumi:"key"`
	Options *MongoIndexOptionsResponse `pulumi:"options"`
}

type MongoIndexResponseOutput struct{ *pulumi.OutputState }

func (MongoIndexResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoIndexResponse)(nil)).Elem()
}

func (o MongoIndexResponseOutput) ToMongoIndexResponseOutput() MongoIndexResponseOutput {
	return o
}

func (o MongoIndexResponseOutput) ToMongoIndexResponseOutputWithContext(ctx context.Context) MongoIndexResponseOutput {
	return o
}

func (o MongoIndexResponseOutput) Key() MongoIndexKeysResponsePtrOutput {
	return o.ApplyT(func(v MongoIndexResponse) *MongoIndexKeysResponse { return v.Key }).(MongoIndexKeysResponsePtrOutput)
}

func (o MongoIndexResponseOutput) Options() MongoIndexOptionsResponsePtrOutput {
	return o.ApplyT(func(v MongoIndexResponse) *MongoIndexOptionsResponse { return v.Options }).(MongoIndexOptionsResponsePtrOutput)
}

type MongoIndexResponseArrayOutput struct{ *pulumi.OutputState }

func (MongoIndexResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MongoIndexResponse)(nil)).Elem()
}

func (o MongoIndexResponseArrayOutput) ToMongoIndexResponseArrayOutput() MongoIndexResponseArrayOutput {
	return o
}

func (o MongoIndexResponseArrayOutput) ToMongoIndexResponseArrayOutputWithContext(ctx context.Context) MongoIndexResponseArrayOutput {
	return o
}

func (o MongoIndexResponseArrayOutput) Index(i pulumi.IntInput) MongoIndexResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MongoIndexResponse {
		return vs[0].([]MongoIndexResponse)[vs[1].(int)]
	}).(MongoIndexResponseOutput)
}

type PeriodicModeBackupPolicy struct {
	PeriodicModeProperties *PeriodicModeProperties `pulumi:"periodicModeProperties"`
	Type                   string                  `pulumi:"type"`
}

type PeriodicModeBackupPolicyResponse struct {
	PeriodicModeProperties *PeriodicModePropertiesResponse `pulumi:"periodicModeProperties"`
	Type                   string                          `pulumi:"type"`
}

type PeriodicModeProperties struct {
	BackupIntervalInMinutes        *int `pulumi:"backupIntervalInMinutes"`
	BackupRetentionIntervalInHours *int `pulumi:"backupRetentionIntervalInHours"`
}

type PeriodicModePropertiesResponse struct {
	BackupIntervalInMinutes        *int `pulumi:"backupIntervalInMinutes"`
	BackupRetentionIntervalInHours *int `pulumi:"backupRetentionIntervalInHours"`
}

type PrivateEndpointConnectionResponse struct {
	Id                                string                                             `pulumi:"id"`
	Name                              string                                             `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointPropertyResponse                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStatePropertyResponse `pulumi:"privateLinkServiceConnectionState"`
	Type                              string                                             `pulumi:"type"`
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

func (o PrivateEndpointConnectionResponseOutput) PrivateEndpoint() PrivateEndpointPropertyResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateEndpointPropertyResponse { return v.PrivateEndpoint }).(PrivateEndpointPropertyResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateLinkServiceConnectionStatePropertyResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStatePropertyResponsePtrOutput)
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

type PrivateEndpointPropertyResponse struct {
	Id *string `pulumi:"id"`
}

type PrivateEndpointPropertyResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointPropertyResponse)(nil)).Elem()
}

func (o PrivateEndpointPropertyResponseOutput) ToPrivateEndpointPropertyResponseOutput() PrivateEndpointPropertyResponseOutput {
	return o
}

func (o PrivateEndpointPropertyResponseOutput) ToPrivateEndpointPropertyResponseOutputWithContext(ctx context.Context) PrivateEndpointPropertyResponseOutput {
	return o
}

func (o PrivateEndpointPropertyResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointPropertyResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type PrivateEndpointPropertyResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPropertyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointPropertyResponse)(nil)).Elem()
}

func (o PrivateEndpointPropertyResponsePtrOutput) ToPrivateEndpointPropertyResponsePtrOutput() PrivateEndpointPropertyResponsePtrOutput {
	return o
}

func (o PrivateEndpointPropertyResponsePtrOutput) ToPrivateEndpointPropertyResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyResponsePtrOutput {
	return o
}

func (o PrivateEndpointPropertyResponsePtrOutput) Elem() PrivateEndpointPropertyResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointPropertyResponse) PrivateEndpointPropertyResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointPropertyResponse
		return ret
	}).(PrivateEndpointPropertyResponseOutput)
}

func (o PrivateEndpointPropertyResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointPropertyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStatePropertyResponse struct {
	ActionsRequired string  `pulumi:"actionsRequired"`
	Status          *string `pulumi:"status"`
}

type PrivateLinkServiceConnectionStatePropertyResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStatePropertyResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) ToPrivateLinkServiceConnectionStatePropertyResponseOutput() PrivateLinkServiceConnectionStatePropertyResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) ToPrivateLinkServiceConnectionStatePropertyResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) ActionsRequired() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStatePropertyResponse) string { return v.ActionsRequired }).(pulumi.StringOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStatePropertyResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStatePropertyResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStatePropertyResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutput() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) Elem() PrivateLinkServiceConnectionStatePropertyResponseOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStatePropertyResponse) PrivateLinkServiceConnectionStatePropertyResponse {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionStatePropertyResponse
		return ret
	}).(PrivateLinkServiceConnectionStatePropertyResponseOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStatePropertyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStatePropertyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type SpatialSpec struct {
	Path  *string  `pulumi:"path"`
	Types []string `pulumi:"types"`
}





type SpatialSpecInput interface {
	pulumi.Input

	ToSpatialSpecOutput() SpatialSpecOutput
	ToSpatialSpecOutputWithContext(context.Context) SpatialSpecOutput
}

type SpatialSpecArgs struct {
	Path  pulumi.StringPtrInput   `pulumi:"path"`
	Types pulumi.StringArrayInput `pulumi:"types"`
}

func (SpatialSpecArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SpatialSpec)(nil)).Elem()
}

func (i SpatialSpecArgs) ToSpatialSpecOutput() SpatialSpecOutput {
	return i.ToSpatialSpecOutputWithContext(context.Background())
}

func (i SpatialSpecArgs) ToSpatialSpecOutputWithContext(ctx context.Context) SpatialSpecOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpatialSpecOutput)
}





type SpatialSpecArrayInput interface {
	pulumi.Input

	ToSpatialSpecArrayOutput() SpatialSpecArrayOutput
	ToSpatialSpecArrayOutputWithContext(context.Context) SpatialSpecArrayOutput
}

type SpatialSpecArray []SpatialSpecInput

func (SpatialSpecArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SpatialSpec)(nil)).Elem()
}

func (i SpatialSpecArray) ToSpatialSpecArrayOutput() SpatialSpecArrayOutput {
	return i.ToSpatialSpecArrayOutputWithContext(context.Background())
}

func (i SpatialSpecArray) ToSpatialSpecArrayOutputWithContext(ctx context.Context) SpatialSpecArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpatialSpecArrayOutput)
}

type SpatialSpecOutput struct{ *pulumi.OutputState }

func (SpatialSpecOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SpatialSpec)(nil)).Elem()
}

func (o SpatialSpecOutput) ToSpatialSpecOutput() SpatialSpecOutput {
	return o
}

func (o SpatialSpecOutput) ToSpatialSpecOutputWithContext(ctx context.Context) SpatialSpecOutput {
	return o
}

func (o SpatialSpecOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SpatialSpec) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o SpatialSpecOutput) Types() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SpatialSpec) []string { return v.Types }).(pulumi.StringArrayOutput)
}

type SpatialSpecArrayOutput struct{ *pulumi.OutputState }

func (SpatialSpecArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SpatialSpec)(nil)).Elem()
}

func (o SpatialSpecArrayOutput) ToSpatialSpecArrayOutput() SpatialSpecArrayOutput {
	return o
}

func (o SpatialSpecArrayOutput) ToSpatialSpecArrayOutputWithContext(ctx context.Context) SpatialSpecArrayOutput {
	return o
}

func (o SpatialSpecArrayOutput) Index(i pulumi.IntInput) SpatialSpecOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SpatialSpec {
		return vs[0].([]SpatialSpec)[vs[1].(int)]
	}).(SpatialSpecOutput)
}

type SpatialSpecResponse struct {
	Path  *string  `pulumi:"path"`
	Types []string `pulumi:"types"`
}

type SpatialSpecResponseOutput struct{ *pulumi.OutputState }

func (SpatialSpecResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SpatialSpecResponse)(nil)).Elem()
}

func (o SpatialSpecResponseOutput) ToSpatialSpecResponseOutput() SpatialSpecResponseOutput {
	return o
}

func (o SpatialSpecResponseOutput) ToSpatialSpecResponseOutputWithContext(ctx context.Context) SpatialSpecResponseOutput {
	return o
}

func (o SpatialSpecResponseOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SpatialSpecResponse) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o SpatialSpecResponseOutput) Types() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SpatialSpecResponse) []string { return v.Types }).(pulumi.StringArrayOutput)
}

type SpatialSpecResponseArrayOutput struct{ *pulumi.OutputState }

func (SpatialSpecResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SpatialSpecResponse)(nil)).Elem()
}

func (o SpatialSpecResponseArrayOutput) ToSpatialSpecResponseArrayOutput() SpatialSpecResponseArrayOutput {
	return o
}

func (o SpatialSpecResponseArrayOutput) ToSpatialSpecResponseArrayOutputWithContext(ctx context.Context) SpatialSpecResponseArrayOutput {
	return o
}

func (o SpatialSpecResponseArrayOutput) Index(i pulumi.IntInput) SpatialSpecResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SpatialSpecResponse {
		return vs[0].([]SpatialSpecResponse)[vs[1].(int)]
	}).(SpatialSpecResponseOutput)
}

type SqlContainerGetPropertiesResponseOptions struct {
	AutoscaleSettings *AutoscaleSettingsResponse `pulumi:"autoscaleSettings"`
	Throughput        *int                       `pulumi:"throughput"`
}

type SqlContainerGetPropertiesResponseOptionsOutput struct{ *pulumi.OutputState }

func (SqlContainerGetPropertiesResponseOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlContainerGetPropertiesResponseOptions)(nil)).Elem()
}

func (o SqlContainerGetPropertiesResponseOptionsOutput) ToSqlContainerGetPropertiesResponseOptionsOutput() SqlContainerGetPropertiesResponseOptionsOutput {
	return o
}

func (o SqlContainerGetPropertiesResponseOptionsOutput) ToSqlContainerGetPropertiesResponseOptionsOutputWithContext(ctx context.Context) SqlContainerGetPropertiesResponseOptionsOutput {
	return o
}

func (o SqlContainerGetPropertiesResponseOptionsOutput) AutoscaleSettings() AutoscaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v SqlContainerGetPropertiesResponseOptions) *AutoscaleSettingsResponse {
		return v.AutoscaleSettings
	}).(AutoscaleSettingsResponsePtrOutput)
}

func (o SqlContainerGetPropertiesResponseOptionsOutput) Throughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SqlContainerGetPropertiesResponseOptions) *int { return v.Throughput }).(pulumi.IntPtrOutput)
}

type SqlContainerGetPropertiesResponseOptionsPtrOutput struct{ *pulumi.OutputState }

func (SqlContainerGetPropertiesResponseOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlContainerGetPropertiesResponseOptions)(nil)).Elem()
}

func (o SqlContainerGetPropertiesResponseOptionsPtrOutput) ToSqlContainerGetPropertiesResponseOptionsPtrOutput() SqlContainerGetPropertiesResponseOptionsPtrOutput {
	return o
}

func (o SqlContainerGetPropertiesResponseOptionsPtrOutput) ToSqlContainerGetPropertiesResponseOptionsPtrOutputWithContext(ctx context.Context) SqlContainerGetPropertiesResponseOptionsPtrOutput {
	return o
}

func (o SqlContainerGetPropertiesResponseOptionsPtrOutput) Elem() SqlContainerGetPropertiesResponseOptionsOutput {
	return o.ApplyT(func(v *SqlContainerGetPropertiesResponseOptions) SqlContainerGetPropertiesResponseOptions {
		if v != nil {
			return *v
		}
		var ret SqlContainerGetPropertiesResponseOptions
		return ret
	}).(SqlContainerGetPropertiesResponseOptionsOutput)
}

func (o SqlContainerGetPropertiesResponseOptionsPtrOutput) AutoscaleSettings() AutoscaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v *SqlContainerGetPropertiesResponseOptions) *AutoscaleSettingsResponse {
		if v == nil {
			return nil
		}
		return v.AutoscaleSettings
	}).(AutoscaleSettingsResponsePtrOutput)
}

func (o SqlContainerGetPropertiesResponseOptionsPtrOutput) Throughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SqlContainerGetPropertiesResponseOptions) *int {
		if v == nil {
			return nil
		}
		return v.Throughput
	}).(pulumi.IntPtrOutput)
}

type SqlContainerGetPropertiesResponseResource struct {
	AnalyticalStorageTtl     *float64                          `pulumi:"analyticalStorageTtl"`
	ConflictResolutionPolicy *ConflictResolutionPolicyResponse `pulumi:"conflictResolutionPolicy"`
	DefaultTtl               *int                              `pulumi:"defaultTtl"`
	Etag                     string                            `pulumi:"etag"`
	Id                       string                            `pulumi:"id"`
	IndexingPolicy           *IndexingPolicyResponse           `pulumi:"indexingPolicy"`
	PartitionKey             *ContainerPartitionKeyResponse    `pulumi:"partitionKey"`
	Rid                      string                            `pulumi:"rid"`
	Ts                       float64                           `pulumi:"ts"`
	UniqueKeyPolicy          *UniqueKeyPolicyResponse          `pulumi:"uniqueKeyPolicy"`
}


func (val *SqlContainerGetPropertiesResponseResource) Defaults() *SqlContainerGetPropertiesResponseResource {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ConflictResolutionPolicy = tmp.ConflictResolutionPolicy.Defaults()

	tmp.IndexingPolicy = tmp.IndexingPolicy.Defaults()

	tmp.PartitionKey = tmp.PartitionKey.Defaults()

	return &tmp
}

type SqlContainerGetPropertiesResponseResourceOutput struct{ *pulumi.OutputState }

func (SqlContainerGetPropertiesResponseResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlContainerGetPropertiesResponseResource)(nil)).Elem()
}

func (o SqlContainerGetPropertiesResponseResourceOutput) ToSqlContainerGetPropertiesResponseResourceOutput() SqlContainerGetPropertiesResponseResourceOutput {
	return o
}

func (o SqlContainerGetPropertiesResponseResourceOutput) ToSqlContainerGetPropertiesResponseResourceOutputWithContext(ctx context.Context) SqlContainerGetPropertiesResponseResourceOutput {
	return o
}

func (o SqlContainerGetPropertiesResponseResourceOutput) AnalyticalStorageTtl() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SqlContainerGetPropertiesResponseResource) *float64 { return v.AnalyticalStorageTtl }).(pulumi.Float64PtrOutput)
}

func (o SqlContainerGetPropertiesResponseResourceOutput) ConflictResolutionPolicy() ConflictResolutionPolicyResponsePtrOutput {
	return o.ApplyT(func(v SqlContainerGetPropertiesResponseResource) *ConflictResolutionPolicyResponse {
		return v.ConflictResolutionPolicy
	}).(ConflictResolutionPolicyResponsePtrOutput)
}

func (o SqlContainerGetPropertiesResponseResourceOutput) DefaultTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SqlContainerGetPropertiesResponseResource) *int { return v.DefaultTtl }).(pulumi.IntPtrOutput)
}

func (o SqlContainerGetPropertiesResponseResourceOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v SqlContainerGetPropertiesResponseResource) string { return v.Etag }).(pulumi.StringOutput)
}

func (o SqlContainerGetPropertiesResponseResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SqlContainerGetPropertiesResponseResource) string { return v.Id }).(pulumi.StringOutput)
}

func (o SqlContainerGetPropertiesResponseResourceOutput) IndexingPolicy() IndexingPolicyResponsePtrOutput {
	return o.ApplyT(func(v SqlContainerGetPropertiesResponseResource) *IndexingPolicyResponse { return v.IndexingPolicy }).(IndexingPolicyResponsePtrOutput)
}

func (o SqlContainerGetPropertiesResponseResourceOutput) PartitionKey() ContainerPartitionKeyResponsePtrOutput {
	return o.ApplyT(func(v SqlContainerGetPropertiesResponseResource) *ContainerPartitionKeyResponse {
		return v.PartitionKey
	}).(ContainerPartitionKeyResponsePtrOutput)
}

func (o SqlContainerGetPropertiesResponseResourceOutput) Rid() pulumi.StringOutput {
	return o.ApplyT(func(v SqlContainerGetPropertiesResponseResource) string { return v.Rid }).(pulumi.StringOutput)
}

func (o SqlContainerGetPropertiesResponseResourceOutput) Ts() pulumi.Float64Output {
	return o.ApplyT(func(v SqlContainerGetPropertiesResponseResource) float64 { return v.Ts }).(pulumi.Float64Output)
}

func (o SqlContainerGetPropertiesResponseResourceOutput) UniqueKeyPolicy() UniqueKeyPolicyResponsePtrOutput {
	return o.ApplyT(func(v SqlContainerGetPropertiesResponseResource) *UniqueKeyPolicyResponse { return v.UniqueKeyPolicy }).(UniqueKeyPolicyResponsePtrOutput)
}

type SqlContainerGetPropertiesResponseResourcePtrOutput struct{ *pulumi.OutputState }

func (SqlContainerGetPropertiesResponseResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlContainerGetPropertiesResponseResource)(nil)).Elem()
}

func (o SqlContainerGetPropertiesResponseResourcePtrOutput) ToSqlContainerGetPropertiesResponseResourcePtrOutput() SqlContainerGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o SqlContainerGetPropertiesResponseResourcePtrOutput) ToSqlContainerGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) SqlContainerGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o SqlContainerGetPropertiesResponseResourcePtrOutput) Elem() SqlContainerGetPropertiesResponseResourceOutput {
	return o.ApplyT(func(v *SqlContainerGetPropertiesResponseResource) SqlContainerGetPropertiesResponseResource {
		if v != nil {
			return *v
		}
		var ret SqlContainerGetPropertiesResponseResource
		return ret
	}).(SqlContainerGetPropertiesResponseResourceOutput)
}

func (o SqlContainerGetPropertiesResponseResourcePtrOutput) AnalyticalStorageTtl() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SqlContainerGetPropertiesResponseResource) *float64 {
		if v == nil {
			return nil
		}
		return v.AnalyticalStorageTtl
	}).(pulumi.Float64PtrOutput)
}

func (o SqlContainerGetPropertiesResponseResourcePtrOutput) ConflictResolutionPolicy() ConflictResolutionPolicyResponsePtrOutput {
	return o.ApplyT(func(v *SqlContainerGetPropertiesResponseResource) *ConflictResolutionPolicyResponse {
		if v == nil {
			return nil
		}
		return v.ConflictResolutionPolicy
	}).(ConflictResolutionPolicyResponsePtrOutput)
}

func (o SqlContainerGetPropertiesResponseResourcePtrOutput) DefaultTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SqlContainerGetPropertiesResponseResource) *int {
		if v == nil {
			return nil
		}
		return v.DefaultTtl
	}).(pulumi.IntPtrOutput)
}

func (o SqlContainerGetPropertiesResponseResourcePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlContainerGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o SqlContainerGetPropertiesResponseResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlContainerGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o SqlContainerGetPropertiesResponseResourcePtrOutput) IndexingPolicy() IndexingPolicyResponsePtrOutput {
	return o.ApplyT(func(v *SqlContainerGetPropertiesResponseResource) *IndexingPolicyResponse {
		if v == nil {
			return nil
		}
		return v.IndexingPolicy
	}).(IndexingPolicyResponsePtrOutput)
}

func (o SqlContainerGetPropertiesResponseResourcePtrOutput) PartitionKey() ContainerPartitionKeyResponsePtrOutput {
	return o.ApplyT(func(v *SqlContainerGetPropertiesResponseResource) *ContainerPartitionKeyResponse {
		if v == nil {
			return nil
		}
		return v.PartitionKey
	}).(ContainerPartitionKeyResponsePtrOutput)
}

func (o SqlContainerGetPropertiesResponseResourcePtrOutput) Rid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlContainerGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Rid
	}).(pulumi.StringPtrOutput)
}

func (o SqlContainerGetPropertiesResponseResourcePtrOutput) Ts() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SqlContainerGetPropertiesResponseResource) *float64 {
		if v == nil {
			return nil
		}
		return &v.Ts
	}).(pulumi.Float64PtrOutput)
}

func (o SqlContainerGetPropertiesResponseResourcePtrOutput) UniqueKeyPolicy() UniqueKeyPolicyResponsePtrOutput {
	return o.ApplyT(func(v *SqlContainerGetPropertiesResponseResource) *UniqueKeyPolicyResponse {
		if v == nil {
			return nil
		}
		return v.UniqueKeyPolicy
	}).(UniqueKeyPolicyResponsePtrOutput)
}

type SqlContainerResource struct {
	AnalyticalStorageTtl     *float64                  `pulumi:"analyticalStorageTtl"`
	ConflictResolutionPolicy *ConflictResolutionPolicy `pulumi:"conflictResolutionPolicy"`
	DefaultTtl               *int                      `pulumi:"defaultTtl"`
	Id                       string                    `pulumi:"id"`
	IndexingPolicy           *IndexingPolicy           `pulumi:"indexingPolicy"`
	PartitionKey             *ContainerPartitionKey    `pulumi:"partitionKey"`
	UniqueKeyPolicy          *UniqueKeyPolicy          `pulumi:"uniqueKeyPolicy"`
}


func (val *SqlContainerResource) Defaults() *SqlContainerResource {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ConflictResolutionPolicy = tmp.ConflictResolutionPolicy.Defaults()

	tmp.IndexingPolicy = tmp.IndexingPolicy.Defaults()

	tmp.PartitionKey = tmp.PartitionKey.Defaults()

	return &tmp
}





type SqlContainerResourceInput interface {
	pulumi.Input

	ToSqlContainerResourceOutput() SqlContainerResourceOutput
	ToSqlContainerResourceOutputWithContext(context.Context) SqlContainerResourceOutput
}

type SqlContainerResourceArgs struct {
	AnalyticalStorageTtl     pulumi.Float64PtrInput           `pulumi:"analyticalStorageTtl"`
	ConflictResolutionPolicy ConflictResolutionPolicyPtrInput `pulumi:"conflictResolutionPolicy"`
	DefaultTtl               pulumi.IntPtrInput               `pulumi:"defaultTtl"`
	Id                       pulumi.StringInput               `pulumi:"id"`
	IndexingPolicy           IndexingPolicyPtrInput           `pulumi:"indexingPolicy"`
	PartitionKey             ContainerPartitionKeyPtrInput    `pulumi:"partitionKey"`
	UniqueKeyPolicy          UniqueKeyPolicyPtrInput          `pulumi:"uniqueKeyPolicy"`
}

func (SqlContainerResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlContainerResource)(nil)).Elem()
}

func (i SqlContainerResourceArgs) ToSqlContainerResourceOutput() SqlContainerResourceOutput {
	return i.ToSqlContainerResourceOutputWithContext(context.Background())
}

func (i SqlContainerResourceArgs) ToSqlContainerResourceOutputWithContext(ctx context.Context) SqlContainerResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlContainerResourceOutput)
}

type SqlContainerResourceOutput struct{ *pulumi.OutputState }

func (SqlContainerResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlContainerResource)(nil)).Elem()
}

func (o SqlContainerResourceOutput) ToSqlContainerResourceOutput() SqlContainerResourceOutput {
	return o
}

func (o SqlContainerResourceOutput) ToSqlContainerResourceOutputWithContext(ctx context.Context) SqlContainerResourceOutput {
	return o
}

func (o SqlContainerResourceOutput) AnalyticalStorageTtl() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SqlContainerResource) *float64 { return v.AnalyticalStorageTtl }).(pulumi.Float64PtrOutput)
}

func (o SqlContainerResourceOutput) ConflictResolutionPolicy() ConflictResolutionPolicyPtrOutput {
	return o.ApplyT(func(v SqlContainerResource) *ConflictResolutionPolicy { return v.ConflictResolutionPolicy }).(ConflictResolutionPolicyPtrOutput)
}

func (o SqlContainerResourceOutput) DefaultTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SqlContainerResource) *int { return v.DefaultTtl }).(pulumi.IntPtrOutput)
}

func (o SqlContainerResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SqlContainerResource) string { return v.Id }).(pulumi.StringOutput)
}

func (o SqlContainerResourceOutput) IndexingPolicy() IndexingPolicyPtrOutput {
	return o.ApplyT(func(v SqlContainerResource) *IndexingPolicy { return v.IndexingPolicy }).(IndexingPolicyPtrOutput)
}

func (o SqlContainerResourceOutput) PartitionKey() ContainerPartitionKeyPtrOutput {
	return o.ApplyT(func(v SqlContainerResource) *ContainerPartitionKey { return v.PartitionKey }).(ContainerPartitionKeyPtrOutput)
}

func (o SqlContainerResourceOutput) UniqueKeyPolicy() UniqueKeyPolicyPtrOutput {
	return o.ApplyT(func(v SqlContainerResource) *UniqueKeyPolicy { return v.UniqueKeyPolicy }).(UniqueKeyPolicyPtrOutput)
}

type SqlDatabaseGetPropertiesResponseOptions struct {
	AutoscaleSettings *AutoscaleSettingsResponse `pulumi:"autoscaleSettings"`
	Throughput        *int                       `pulumi:"throughput"`
}

type SqlDatabaseGetPropertiesResponseOptionsOutput struct{ *pulumi.OutputState }

func (SqlDatabaseGetPropertiesResponseOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlDatabaseGetPropertiesResponseOptions)(nil)).Elem()
}

func (o SqlDatabaseGetPropertiesResponseOptionsOutput) ToSqlDatabaseGetPropertiesResponseOptionsOutput() SqlDatabaseGetPropertiesResponseOptionsOutput {
	return o
}

func (o SqlDatabaseGetPropertiesResponseOptionsOutput) ToSqlDatabaseGetPropertiesResponseOptionsOutputWithContext(ctx context.Context) SqlDatabaseGetPropertiesResponseOptionsOutput {
	return o
}

func (o SqlDatabaseGetPropertiesResponseOptionsOutput) AutoscaleSettings() AutoscaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v SqlDatabaseGetPropertiesResponseOptions) *AutoscaleSettingsResponse { return v.AutoscaleSettings }).(AutoscaleSettingsResponsePtrOutput)
}

func (o SqlDatabaseGetPropertiesResponseOptionsOutput) Throughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SqlDatabaseGetPropertiesResponseOptions) *int { return v.Throughput }).(pulumi.IntPtrOutput)
}

type SqlDatabaseGetPropertiesResponseOptionsPtrOutput struct{ *pulumi.OutputState }

func (SqlDatabaseGetPropertiesResponseOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlDatabaseGetPropertiesResponseOptions)(nil)).Elem()
}

func (o SqlDatabaseGetPropertiesResponseOptionsPtrOutput) ToSqlDatabaseGetPropertiesResponseOptionsPtrOutput() SqlDatabaseGetPropertiesResponseOptionsPtrOutput {
	return o
}

func (o SqlDatabaseGetPropertiesResponseOptionsPtrOutput) ToSqlDatabaseGetPropertiesResponseOptionsPtrOutputWithContext(ctx context.Context) SqlDatabaseGetPropertiesResponseOptionsPtrOutput {
	return o
}

func (o SqlDatabaseGetPropertiesResponseOptionsPtrOutput) Elem() SqlDatabaseGetPropertiesResponseOptionsOutput {
	return o.ApplyT(func(v *SqlDatabaseGetPropertiesResponseOptions) SqlDatabaseGetPropertiesResponseOptions {
		if v != nil {
			return *v
		}
		var ret SqlDatabaseGetPropertiesResponseOptions
		return ret
	}).(SqlDatabaseGetPropertiesResponseOptionsOutput)
}

func (o SqlDatabaseGetPropertiesResponseOptionsPtrOutput) AutoscaleSettings() AutoscaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v *SqlDatabaseGetPropertiesResponseOptions) *AutoscaleSettingsResponse {
		if v == nil {
			return nil
		}
		return v.AutoscaleSettings
	}).(AutoscaleSettingsResponsePtrOutput)
}

func (o SqlDatabaseGetPropertiesResponseOptionsPtrOutput) Throughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SqlDatabaseGetPropertiesResponseOptions) *int {
		if v == nil {
			return nil
		}
		return v.Throughput
	}).(pulumi.IntPtrOutput)
}

type SqlDatabaseGetPropertiesResponseResource struct {
	Colls *string `pulumi:"colls"`
	Etag  string  `pulumi:"etag"`
	Id    string  `pulumi:"id"`
	Rid   string  `pulumi:"rid"`
	Ts    float64 `pulumi:"ts"`
	Users *string `pulumi:"users"`
}

type SqlDatabaseGetPropertiesResponseResourceOutput struct{ *pulumi.OutputState }

func (SqlDatabaseGetPropertiesResponseResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlDatabaseGetPropertiesResponseResource)(nil)).Elem()
}

func (o SqlDatabaseGetPropertiesResponseResourceOutput) ToSqlDatabaseGetPropertiesResponseResourceOutput() SqlDatabaseGetPropertiesResponseResourceOutput {
	return o
}

func (o SqlDatabaseGetPropertiesResponseResourceOutput) ToSqlDatabaseGetPropertiesResponseResourceOutputWithContext(ctx context.Context) SqlDatabaseGetPropertiesResponseResourceOutput {
	return o
}

func (o SqlDatabaseGetPropertiesResponseResourceOutput) Colls() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlDatabaseGetPropertiesResponseResource) *string { return v.Colls }).(pulumi.StringPtrOutput)
}

func (o SqlDatabaseGetPropertiesResponseResourceOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v SqlDatabaseGetPropertiesResponseResource) string { return v.Etag }).(pulumi.StringOutput)
}

func (o SqlDatabaseGetPropertiesResponseResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SqlDatabaseGetPropertiesResponseResource) string { return v.Id }).(pulumi.StringOutput)
}

func (o SqlDatabaseGetPropertiesResponseResourceOutput) Rid() pulumi.StringOutput {
	return o.ApplyT(func(v SqlDatabaseGetPropertiesResponseResource) string { return v.Rid }).(pulumi.StringOutput)
}

func (o SqlDatabaseGetPropertiesResponseResourceOutput) Ts() pulumi.Float64Output {
	return o.ApplyT(func(v SqlDatabaseGetPropertiesResponseResource) float64 { return v.Ts }).(pulumi.Float64Output)
}

func (o SqlDatabaseGetPropertiesResponseResourceOutput) Users() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlDatabaseGetPropertiesResponseResource) *string { return v.Users }).(pulumi.StringPtrOutput)
}

type SqlDatabaseGetPropertiesResponseResourcePtrOutput struct{ *pulumi.OutputState }

func (SqlDatabaseGetPropertiesResponseResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlDatabaseGetPropertiesResponseResource)(nil)).Elem()
}

func (o SqlDatabaseGetPropertiesResponseResourcePtrOutput) ToSqlDatabaseGetPropertiesResponseResourcePtrOutput() SqlDatabaseGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o SqlDatabaseGetPropertiesResponseResourcePtrOutput) ToSqlDatabaseGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) SqlDatabaseGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o SqlDatabaseGetPropertiesResponseResourcePtrOutput) Elem() SqlDatabaseGetPropertiesResponseResourceOutput {
	return o.ApplyT(func(v *SqlDatabaseGetPropertiesResponseResource) SqlDatabaseGetPropertiesResponseResource {
		if v != nil {
			return *v
		}
		var ret SqlDatabaseGetPropertiesResponseResource
		return ret
	}).(SqlDatabaseGetPropertiesResponseResourceOutput)
}

func (o SqlDatabaseGetPropertiesResponseResourcePtrOutput) Colls() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlDatabaseGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return v.Colls
	}).(pulumi.StringPtrOutput)
}

func (o SqlDatabaseGetPropertiesResponseResourcePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlDatabaseGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o SqlDatabaseGetPropertiesResponseResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlDatabaseGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o SqlDatabaseGetPropertiesResponseResourcePtrOutput) Rid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlDatabaseGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Rid
	}).(pulumi.StringPtrOutput)
}

func (o SqlDatabaseGetPropertiesResponseResourcePtrOutput) Ts() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SqlDatabaseGetPropertiesResponseResource) *float64 {
		if v == nil {
			return nil
		}
		return &v.Ts
	}).(pulumi.Float64PtrOutput)
}

func (o SqlDatabaseGetPropertiesResponseResourcePtrOutput) Users() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlDatabaseGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return v.Users
	}).(pulumi.StringPtrOutput)
}

type SqlDatabaseResource struct {
	Id string `pulumi:"id"`
}





type SqlDatabaseResourceInput interface {
	pulumi.Input

	ToSqlDatabaseResourceOutput() SqlDatabaseResourceOutput
	ToSqlDatabaseResourceOutputWithContext(context.Context) SqlDatabaseResourceOutput
}

type SqlDatabaseResourceArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (SqlDatabaseResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlDatabaseResource)(nil)).Elem()
}

func (i SqlDatabaseResourceArgs) ToSqlDatabaseResourceOutput() SqlDatabaseResourceOutput {
	return i.ToSqlDatabaseResourceOutputWithContext(context.Background())
}

func (i SqlDatabaseResourceArgs) ToSqlDatabaseResourceOutputWithContext(ctx context.Context) SqlDatabaseResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlDatabaseResourceOutput)
}

type SqlDatabaseResourceOutput struct{ *pulumi.OutputState }

func (SqlDatabaseResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlDatabaseResource)(nil)).Elem()
}

func (o SqlDatabaseResourceOutput) ToSqlDatabaseResourceOutput() SqlDatabaseResourceOutput {
	return o
}

func (o SqlDatabaseResourceOutput) ToSqlDatabaseResourceOutputWithContext(ctx context.Context) SqlDatabaseResourceOutput {
	return o
}

func (o SqlDatabaseResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SqlDatabaseResource) string { return v.Id }).(pulumi.StringOutput)
}

type SqlStoredProcedureGetPropertiesResponseResource struct {
	Body *string `pulumi:"body"`
	Etag string  `pulumi:"etag"`
	Id   string  `pulumi:"id"`
	Rid  string  `pulumi:"rid"`
	Ts   float64 `pulumi:"ts"`
}

type SqlStoredProcedureGetPropertiesResponseResourceOutput struct{ *pulumi.OutputState }

func (SqlStoredProcedureGetPropertiesResponseResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlStoredProcedureGetPropertiesResponseResource)(nil)).Elem()
}

func (o SqlStoredProcedureGetPropertiesResponseResourceOutput) ToSqlStoredProcedureGetPropertiesResponseResourceOutput() SqlStoredProcedureGetPropertiesResponseResourceOutput {
	return o
}

func (o SqlStoredProcedureGetPropertiesResponseResourceOutput) ToSqlStoredProcedureGetPropertiesResponseResourceOutputWithContext(ctx context.Context) SqlStoredProcedureGetPropertiesResponseResourceOutput {
	return o
}

func (o SqlStoredProcedureGetPropertiesResponseResourceOutput) Body() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlStoredProcedureGetPropertiesResponseResource) *string { return v.Body }).(pulumi.StringPtrOutput)
}

func (o SqlStoredProcedureGetPropertiesResponseResourceOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v SqlStoredProcedureGetPropertiesResponseResource) string { return v.Etag }).(pulumi.StringOutput)
}

func (o SqlStoredProcedureGetPropertiesResponseResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SqlStoredProcedureGetPropertiesResponseResource) string { return v.Id }).(pulumi.StringOutput)
}

func (o SqlStoredProcedureGetPropertiesResponseResourceOutput) Rid() pulumi.StringOutput {
	return o.ApplyT(func(v SqlStoredProcedureGetPropertiesResponseResource) string { return v.Rid }).(pulumi.StringOutput)
}

func (o SqlStoredProcedureGetPropertiesResponseResourceOutput) Ts() pulumi.Float64Output {
	return o.ApplyT(func(v SqlStoredProcedureGetPropertiesResponseResource) float64 { return v.Ts }).(pulumi.Float64Output)
}

type SqlStoredProcedureGetPropertiesResponseResourcePtrOutput struct{ *pulumi.OutputState }

func (SqlStoredProcedureGetPropertiesResponseResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlStoredProcedureGetPropertiesResponseResource)(nil)).Elem()
}

func (o SqlStoredProcedureGetPropertiesResponseResourcePtrOutput) ToSqlStoredProcedureGetPropertiesResponseResourcePtrOutput() SqlStoredProcedureGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o SqlStoredProcedureGetPropertiesResponseResourcePtrOutput) ToSqlStoredProcedureGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) SqlStoredProcedureGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o SqlStoredProcedureGetPropertiesResponseResourcePtrOutput) Elem() SqlStoredProcedureGetPropertiesResponseResourceOutput {
	return o.ApplyT(func(v *SqlStoredProcedureGetPropertiesResponseResource) SqlStoredProcedureGetPropertiesResponseResource {
		if v != nil {
			return *v
		}
		var ret SqlStoredProcedureGetPropertiesResponseResource
		return ret
	}).(SqlStoredProcedureGetPropertiesResponseResourceOutput)
}

func (o SqlStoredProcedureGetPropertiesResponseResourcePtrOutput) Body() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlStoredProcedureGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return v.Body
	}).(pulumi.StringPtrOutput)
}

func (o SqlStoredProcedureGetPropertiesResponseResourcePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlStoredProcedureGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o SqlStoredProcedureGetPropertiesResponseResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlStoredProcedureGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o SqlStoredProcedureGetPropertiesResponseResourcePtrOutput) Rid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlStoredProcedureGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Rid
	}).(pulumi.StringPtrOutput)
}

func (o SqlStoredProcedureGetPropertiesResponseResourcePtrOutput) Ts() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SqlStoredProcedureGetPropertiesResponseResource) *float64 {
		if v == nil {
			return nil
		}
		return &v.Ts
	}).(pulumi.Float64PtrOutput)
}

type SqlStoredProcedureResource struct {
	Body *string `pulumi:"body"`
	Id   string  `pulumi:"id"`
}





type SqlStoredProcedureResourceInput interface {
	pulumi.Input

	ToSqlStoredProcedureResourceOutput() SqlStoredProcedureResourceOutput
	ToSqlStoredProcedureResourceOutputWithContext(context.Context) SqlStoredProcedureResourceOutput
}

type SqlStoredProcedureResourceArgs struct {
	Body pulumi.StringPtrInput `pulumi:"body"`
	Id   pulumi.StringInput    `pulumi:"id"`
}

func (SqlStoredProcedureResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlStoredProcedureResource)(nil)).Elem()
}

func (i SqlStoredProcedureResourceArgs) ToSqlStoredProcedureResourceOutput() SqlStoredProcedureResourceOutput {
	return i.ToSqlStoredProcedureResourceOutputWithContext(context.Background())
}

func (i SqlStoredProcedureResourceArgs) ToSqlStoredProcedureResourceOutputWithContext(ctx context.Context) SqlStoredProcedureResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlStoredProcedureResourceOutput)
}

type SqlStoredProcedureResourceOutput struct{ *pulumi.OutputState }

func (SqlStoredProcedureResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlStoredProcedureResource)(nil)).Elem()
}

func (o SqlStoredProcedureResourceOutput) ToSqlStoredProcedureResourceOutput() SqlStoredProcedureResourceOutput {
	return o
}

func (o SqlStoredProcedureResourceOutput) ToSqlStoredProcedureResourceOutputWithContext(ctx context.Context) SqlStoredProcedureResourceOutput {
	return o
}

func (o SqlStoredProcedureResourceOutput) Body() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlStoredProcedureResource) *string { return v.Body }).(pulumi.StringPtrOutput)
}

func (o SqlStoredProcedureResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SqlStoredProcedureResource) string { return v.Id }).(pulumi.StringOutput)
}

type SqlTriggerGetPropertiesResponseResource struct {
	Body             *string `pulumi:"body"`
	Etag             string  `pulumi:"etag"`
	Id               string  `pulumi:"id"`
	Rid              string  `pulumi:"rid"`
	TriggerOperation *string `pulumi:"triggerOperation"`
	TriggerType      *string `pulumi:"triggerType"`
	Ts               float64 `pulumi:"ts"`
}

type SqlTriggerGetPropertiesResponseResourceOutput struct{ *pulumi.OutputState }

func (SqlTriggerGetPropertiesResponseResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlTriggerGetPropertiesResponseResource)(nil)).Elem()
}

func (o SqlTriggerGetPropertiesResponseResourceOutput) ToSqlTriggerGetPropertiesResponseResourceOutput() SqlTriggerGetPropertiesResponseResourceOutput {
	return o
}

func (o SqlTriggerGetPropertiesResponseResourceOutput) ToSqlTriggerGetPropertiesResponseResourceOutputWithContext(ctx context.Context) SqlTriggerGetPropertiesResponseResourceOutput {
	return o
}

func (o SqlTriggerGetPropertiesResponseResourceOutput) Body() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlTriggerGetPropertiesResponseResource) *string { return v.Body }).(pulumi.StringPtrOutput)
}

func (o SqlTriggerGetPropertiesResponseResourceOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v SqlTriggerGetPropertiesResponseResource) string { return v.Etag }).(pulumi.StringOutput)
}

func (o SqlTriggerGetPropertiesResponseResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SqlTriggerGetPropertiesResponseResource) string { return v.Id }).(pulumi.StringOutput)
}

func (o SqlTriggerGetPropertiesResponseResourceOutput) Rid() pulumi.StringOutput {
	return o.ApplyT(func(v SqlTriggerGetPropertiesResponseResource) string { return v.Rid }).(pulumi.StringOutput)
}

func (o SqlTriggerGetPropertiesResponseResourceOutput) TriggerOperation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlTriggerGetPropertiesResponseResource) *string { return v.TriggerOperation }).(pulumi.StringPtrOutput)
}

func (o SqlTriggerGetPropertiesResponseResourceOutput) TriggerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlTriggerGetPropertiesResponseResource) *string { return v.TriggerType }).(pulumi.StringPtrOutput)
}

func (o SqlTriggerGetPropertiesResponseResourceOutput) Ts() pulumi.Float64Output {
	return o.ApplyT(func(v SqlTriggerGetPropertiesResponseResource) float64 { return v.Ts }).(pulumi.Float64Output)
}

type SqlTriggerGetPropertiesResponseResourcePtrOutput struct{ *pulumi.OutputState }

func (SqlTriggerGetPropertiesResponseResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlTriggerGetPropertiesResponseResource)(nil)).Elem()
}

func (o SqlTriggerGetPropertiesResponseResourcePtrOutput) ToSqlTriggerGetPropertiesResponseResourcePtrOutput() SqlTriggerGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o SqlTriggerGetPropertiesResponseResourcePtrOutput) ToSqlTriggerGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) SqlTriggerGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o SqlTriggerGetPropertiesResponseResourcePtrOutput) Elem() SqlTriggerGetPropertiesResponseResourceOutput {
	return o.ApplyT(func(v *SqlTriggerGetPropertiesResponseResource) SqlTriggerGetPropertiesResponseResource {
		if v != nil {
			return *v
		}
		var ret SqlTriggerGetPropertiesResponseResource
		return ret
	}).(SqlTriggerGetPropertiesResponseResourceOutput)
}

func (o SqlTriggerGetPropertiesResponseResourcePtrOutput) Body() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlTriggerGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return v.Body
	}).(pulumi.StringPtrOutput)
}

func (o SqlTriggerGetPropertiesResponseResourcePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlTriggerGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o SqlTriggerGetPropertiesResponseResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlTriggerGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o SqlTriggerGetPropertiesResponseResourcePtrOutput) Rid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlTriggerGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Rid
	}).(pulumi.StringPtrOutput)
}

func (o SqlTriggerGetPropertiesResponseResourcePtrOutput) TriggerOperation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlTriggerGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return v.TriggerOperation
	}).(pulumi.StringPtrOutput)
}

func (o SqlTriggerGetPropertiesResponseResourcePtrOutput) TriggerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlTriggerGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return v.TriggerType
	}).(pulumi.StringPtrOutput)
}

func (o SqlTriggerGetPropertiesResponseResourcePtrOutput) Ts() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SqlTriggerGetPropertiesResponseResource) *float64 {
		if v == nil {
			return nil
		}
		return &v.Ts
	}).(pulumi.Float64PtrOutput)
}

type SqlTriggerResource struct {
	Body             *string `pulumi:"body"`
	Id               string  `pulumi:"id"`
	TriggerOperation *string `pulumi:"triggerOperation"`
	TriggerType      *string `pulumi:"triggerType"`
}





type SqlTriggerResourceInput interface {
	pulumi.Input

	ToSqlTriggerResourceOutput() SqlTriggerResourceOutput
	ToSqlTriggerResourceOutputWithContext(context.Context) SqlTriggerResourceOutput
}

type SqlTriggerResourceArgs struct {
	Body             pulumi.StringPtrInput `pulumi:"body"`
	Id               pulumi.StringInput    `pulumi:"id"`
	TriggerOperation pulumi.StringPtrInput `pulumi:"triggerOperation"`
	TriggerType      pulumi.StringPtrInput `pulumi:"triggerType"`
}

func (SqlTriggerResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlTriggerResource)(nil)).Elem()
}

func (i SqlTriggerResourceArgs) ToSqlTriggerResourceOutput() SqlTriggerResourceOutput {
	return i.ToSqlTriggerResourceOutputWithContext(context.Background())
}

func (i SqlTriggerResourceArgs) ToSqlTriggerResourceOutputWithContext(ctx context.Context) SqlTriggerResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlTriggerResourceOutput)
}

type SqlTriggerResourceOutput struct{ *pulumi.OutputState }

func (SqlTriggerResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlTriggerResource)(nil)).Elem()
}

func (o SqlTriggerResourceOutput) ToSqlTriggerResourceOutput() SqlTriggerResourceOutput {
	return o
}

func (o SqlTriggerResourceOutput) ToSqlTriggerResourceOutputWithContext(ctx context.Context) SqlTriggerResourceOutput {
	return o
}

func (o SqlTriggerResourceOutput) Body() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlTriggerResource) *string { return v.Body }).(pulumi.StringPtrOutput)
}

func (o SqlTriggerResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SqlTriggerResource) string { return v.Id }).(pulumi.StringOutput)
}

func (o SqlTriggerResourceOutput) TriggerOperation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlTriggerResource) *string { return v.TriggerOperation }).(pulumi.StringPtrOutput)
}

func (o SqlTriggerResourceOutput) TriggerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlTriggerResource) *string { return v.TriggerType }).(pulumi.StringPtrOutput)
}

type SqlUserDefinedFunctionGetPropertiesResponseResource struct {
	Body *string `pulumi:"body"`
	Etag string  `pulumi:"etag"`
	Id   string  `pulumi:"id"`
	Rid  string  `pulumi:"rid"`
	Ts   float64 `pulumi:"ts"`
}

type SqlUserDefinedFunctionGetPropertiesResponseResourceOutput struct{ *pulumi.OutputState }

func (SqlUserDefinedFunctionGetPropertiesResponseResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlUserDefinedFunctionGetPropertiesResponseResource)(nil)).Elem()
}

func (o SqlUserDefinedFunctionGetPropertiesResponseResourceOutput) ToSqlUserDefinedFunctionGetPropertiesResponseResourceOutput() SqlUserDefinedFunctionGetPropertiesResponseResourceOutput {
	return o
}

func (o SqlUserDefinedFunctionGetPropertiesResponseResourceOutput) ToSqlUserDefinedFunctionGetPropertiesResponseResourceOutputWithContext(ctx context.Context) SqlUserDefinedFunctionGetPropertiesResponseResourceOutput {
	return o
}

func (o SqlUserDefinedFunctionGetPropertiesResponseResourceOutput) Body() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlUserDefinedFunctionGetPropertiesResponseResource) *string { return v.Body }).(pulumi.StringPtrOutput)
}

func (o SqlUserDefinedFunctionGetPropertiesResponseResourceOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v SqlUserDefinedFunctionGetPropertiesResponseResource) string { return v.Etag }).(pulumi.StringOutput)
}

func (o SqlUserDefinedFunctionGetPropertiesResponseResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SqlUserDefinedFunctionGetPropertiesResponseResource) string { return v.Id }).(pulumi.StringOutput)
}

func (o SqlUserDefinedFunctionGetPropertiesResponseResourceOutput) Rid() pulumi.StringOutput {
	return o.ApplyT(func(v SqlUserDefinedFunctionGetPropertiesResponseResource) string { return v.Rid }).(pulumi.StringOutput)
}

func (o SqlUserDefinedFunctionGetPropertiesResponseResourceOutput) Ts() pulumi.Float64Output {
	return o.ApplyT(func(v SqlUserDefinedFunctionGetPropertiesResponseResource) float64 { return v.Ts }).(pulumi.Float64Output)
}

type SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput struct{ *pulumi.OutputState }

func (SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlUserDefinedFunctionGetPropertiesResponseResource)(nil)).Elem()
}

func (o SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput) ToSqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput() SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput) ToSqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput) Elem() SqlUserDefinedFunctionGetPropertiesResponseResourceOutput {
	return o.ApplyT(func(v *SqlUserDefinedFunctionGetPropertiesResponseResource) SqlUserDefinedFunctionGetPropertiesResponseResource {
		if v != nil {
			return *v
		}
		var ret SqlUserDefinedFunctionGetPropertiesResponseResource
		return ret
	}).(SqlUserDefinedFunctionGetPropertiesResponseResourceOutput)
}

func (o SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput) Body() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlUserDefinedFunctionGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return v.Body
	}).(pulumi.StringPtrOutput)
}

func (o SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlUserDefinedFunctionGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlUserDefinedFunctionGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput) Rid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlUserDefinedFunctionGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Rid
	}).(pulumi.StringPtrOutput)
}

func (o SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput) Ts() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SqlUserDefinedFunctionGetPropertiesResponseResource) *float64 {
		if v == nil {
			return nil
		}
		return &v.Ts
	}).(pulumi.Float64PtrOutput)
}

type SqlUserDefinedFunctionResource struct {
	Body *string `pulumi:"body"`
	Id   string  `pulumi:"id"`
}





type SqlUserDefinedFunctionResourceInput interface {
	pulumi.Input

	ToSqlUserDefinedFunctionResourceOutput() SqlUserDefinedFunctionResourceOutput
	ToSqlUserDefinedFunctionResourceOutputWithContext(context.Context) SqlUserDefinedFunctionResourceOutput
}

type SqlUserDefinedFunctionResourceArgs struct {
	Body pulumi.StringPtrInput `pulumi:"body"`
	Id   pulumi.StringInput    `pulumi:"id"`
}

func (SqlUserDefinedFunctionResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlUserDefinedFunctionResource)(nil)).Elem()
}

func (i SqlUserDefinedFunctionResourceArgs) ToSqlUserDefinedFunctionResourceOutput() SqlUserDefinedFunctionResourceOutput {
	return i.ToSqlUserDefinedFunctionResourceOutputWithContext(context.Background())
}

func (i SqlUserDefinedFunctionResourceArgs) ToSqlUserDefinedFunctionResourceOutputWithContext(ctx context.Context) SqlUserDefinedFunctionResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlUserDefinedFunctionResourceOutput)
}

type SqlUserDefinedFunctionResourceOutput struct{ *pulumi.OutputState }

func (SqlUserDefinedFunctionResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlUserDefinedFunctionResource)(nil)).Elem()
}

func (o SqlUserDefinedFunctionResourceOutput) ToSqlUserDefinedFunctionResourceOutput() SqlUserDefinedFunctionResourceOutput {
	return o
}

func (o SqlUserDefinedFunctionResourceOutput) ToSqlUserDefinedFunctionResourceOutputWithContext(ctx context.Context) SqlUserDefinedFunctionResourceOutput {
	return o
}

func (o SqlUserDefinedFunctionResourceOutput) Body() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlUserDefinedFunctionResource) *string { return v.Body }).(pulumi.StringPtrOutput)
}

func (o SqlUserDefinedFunctionResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SqlUserDefinedFunctionResource) string { return v.Id }).(pulumi.StringOutput)
}

type TableGetPropertiesResponseOptions struct {
	AutoscaleSettings *AutoscaleSettingsResponse `pulumi:"autoscaleSettings"`
	Throughput        *int                       `pulumi:"throughput"`
}

type TableGetPropertiesResponseOptionsOutput struct{ *pulumi.OutputState }

func (TableGetPropertiesResponseOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TableGetPropertiesResponseOptions)(nil)).Elem()
}

func (o TableGetPropertiesResponseOptionsOutput) ToTableGetPropertiesResponseOptionsOutput() TableGetPropertiesResponseOptionsOutput {
	return o
}

func (o TableGetPropertiesResponseOptionsOutput) ToTableGetPropertiesResponseOptionsOutputWithContext(ctx context.Context) TableGetPropertiesResponseOptionsOutput {
	return o
}

func (o TableGetPropertiesResponseOptionsOutput) AutoscaleSettings() AutoscaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v TableGetPropertiesResponseOptions) *AutoscaleSettingsResponse { return v.AutoscaleSettings }).(AutoscaleSettingsResponsePtrOutput)
}

func (o TableGetPropertiesResponseOptionsOutput) Throughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v TableGetPropertiesResponseOptions) *int { return v.Throughput }).(pulumi.IntPtrOutput)
}

type TableGetPropertiesResponseOptionsPtrOutput struct{ *pulumi.OutputState }

func (TableGetPropertiesResponseOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TableGetPropertiesResponseOptions)(nil)).Elem()
}

func (o TableGetPropertiesResponseOptionsPtrOutput) ToTableGetPropertiesResponseOptionsPtrOutput() TableGetPropertiesResponseOptionsPtrOutput {
	return o
}

func (o TableGetPropertiesResponseOptionsPtrOutput) ToTableGetPropertiesResponseOptionsPtrOutputWithContext(ctx context.Context) TableGetPropertiesResponseOptionsPtrOutput {
	return o
}

func (o TableGetPropertiesResponseOptionsPtrOutput) Elem() TableGetPropertiesResponseOptionsOutput {
	return o.ApplyT(func(v *TableGetPropertiesResponseOptions) TableGetPropertiesResponseOptions {
		if v != nil {
			return *v
		}
		var ret TableGetPropertiesResponseOptions
		return ret
	}).(TableGetPropertiesResponseOptionsOutput)
}

func (o TableGetPropertiesResponseOptionsPtrOutput) AutoscaleSettings() AutoscaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v *TableGetPropertiesResponseOptions) *AutoscaleSettingsResponse {
		if v == nil {
			return nil
		}
		return v.AutoscaleSettings
	}).(AutoscaleSettingsResponsePtrOutput)
}

func (o TableGetPropertiesResponseOptionsPtrOutput) Throughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TableGetPropertiesResponseOptions) *int {
		if v == nil {
			return nil
		}
		return v.Throughput
	}).(pulumi.IntPtrOutput)
}

type TableGetPropertiesResponseResource struct {
	Etag string  `pulumi:"etag"`
	Id   string  `pulumi:"id"`
	Rid  string  `pulumi:"rid"`
	Ts   float64 `pulumi:"ts"`
}

type TableGetPropertiesResponseResourceOutput struct{ *pulumi.OutputState }

func (TableGetPropertiesResponseResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TableGetPropertiesResponseResource)(nil)).Elem()
}

func (o TableGetPropertiesResponseResourceOutput) ToTableGetPropertiesResponseResourceOutput() TableGetPropertiesResponseResourceOutput {
	return o
}

func (o TableGetPropertiesResponseResourceOutput) ToTableGetPropertiesResponseResourceOutputWithContext(ctx context.Context) TableGetPropertiesResponseResourceOutput {
	return o
}

func (o TableGetPropertiesResponseResourceOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v TableGetPropertiesResponseResource) string { return v.Etag }).(pulumi.StringOutput)
}

func (o TableGetPropertiesResponseResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v TableGetPropertiesResponseResource) string { return v.Id }).(pulumi.StringOutput)
}

func (o TableGetPropertiesResponseResourceOutput) Rid() pulumi.StringOutput {
	return o.ApplyT(func(v TableGetPropertiesResponseResource) string { return v.Rid }).(pulumi.StringOutput)
}

func (o TableGetPropertiesResponseResourceOutput) Ts() pulumi.Float64Output {
	return o.ApplyT(func(v TableGetPropertiesResponseResource) float64 { return v.Ts }).(pulumi.Float64Output)
}

type TableGetPropertiesResponseResourcePtrOutput struct{ *pulumi.OutputState }

func (TableGetPropertiesResponseResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TableGetPropertiesResponseResource)(nil)).Elem()
}

func (o TableGetPropertiesResponseResourcePtrOutput) ToTableGetPropertiesResponseResourcePtrOutput() TableGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o TableGetPropertiesResponseResourcePtrOutput) ToTableGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) TableGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o TableGetPropertiesResponseResourcePtrOutput) Elem() TableGetPropertiesResponseResourceOutput {
	return o.ApplyT(func(v *TableGetPropertiesResponseResource) TableGetPropertiesResponseResource {
		if v != nil {
			return *v
		}
		var ret TableGetPropertiesResponseResource
		return ret
	}).(TableGetPropertiesResponseResourceOutput)
}

func (o TableGetPropertiesResponseResourcePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TableGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o TableGetPropertiesResponseResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TableGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o TableGetPropertiesResponseResourcePtrOutput) Rid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TableGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Rid
	}).(pulumi.StringPtrOutput)
}

func (o TableGetPropertiesResponseResourcePtrOutput) Ts() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *TableGetPropertiesResponseResource) *float64 {
		if v == nil {
			return nil
		}
		return &v.Ts
	}).(pulumi.Float64PtrOutput)
}

type TableResource struct {
	Id string `pulumi:"id"`
}





type TableResourceInput interface {
	pulumi.Input

	ToTableResourceOutput() TableResourceOutput
	ToTableResourceOutputWithContext(context.Context) TableResourceOutput
}

type TableResourceArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (TableResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TableResource)(nil)).Elem()
}

func (i TableResourceArgs) ToTableResourceOutput() TableResourceOutput {
	return i.ToTableResourceOutputWithContext(context.Background())
}

func (i TableResourceArgs) ToTableResourceOutputWithContext(ctx context.Context) TableResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableResourceOutput)
}

type TableResourceOutput struct{ *pulumi.OutputState }

func (TableResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TableResource)(nil)).Elem()
}

func (o TableResourceOutput) ToTableResourceOutput() TableResourceOutput {
	return o
}

func (o TableResourceOutput) ToTableResourceOutputWithContext(ctx context.Context) TableResourceOutput {
	return o
}

func (o TableResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v TableResource) string { return v.Id }).(pulumi.StringOutput)
}

type UniqueKey struct {
	Paths []string `pulumi:"paths"`
}





type UniqueKeyInput interface {
	pulumi.Input

	ToUniqueKeyOutput() UniqueKeyOutput
	ToUniqueKeyOutputWithContext(context.Context) UniqueKeyOutput
}

type UniqueKeyArgs struct {
	Paths pulumi.StringArrayInput `pulumi:"paths"`
}

func (UniqueKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UniqueKey)(nil)).Elem()
}

func (i UniqueKeyArgs) ToUniqueKeyOutput() UniqueKeyOutput {
	return i.ToUniqueKeyOutputWithContext(context.Background())
}

func (i UniqueKeyArgs) ToUniqueKeyOutputWithContext(ctx context.Context) UniqueKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UniqueKeyOutput)
}





type UniqueKeyArrayInput interface {
	pulumi.Input

	ToUniqueKeyArrayOutput() UniqueKeyArrayOutput
	ToUniqueKeyArrayOutputWithContext(context.Context) UniqueKeyArrayOutput
}

type UniqueKeyArray []UniqueKeyInput

func (UniqueKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UniqueKey)(nil)).Elem()
}

func (i UniqueKeyArray) ToUniqueKeyArrayOutput() UniqueKeyArrayOutput {
	return i.ToUniqueKeyArrayOutputWithContext(context.Background())
}

func (i UniqueKeyArray) ToUniqueKeyArrayOutputWithContext(ctx context.Context) UniqueKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UniqueKeyArrayOutput)
}

type UniqueKeyOutput struct{ *pulumi.OutputState }

func (UniqueKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UniqueKey)(nil)).Elem()
}

func (o UniqueKeyOutput) ToUniqueKeyOutput() UniqueKeyOutput {
	return o
}

func (o UniqueKeyOutput) ToUniqueKeyOutputWithContext(ctx context.Context) UniqueKeyOutput {
	return o
}

func (o UniqueKeyOutput) Paths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v UniqueKey) []string { return v.Paths }).(pulumi.StringArrayOutput)
}

type UniqueKeyArrayOutput struct{ *pulumi.OutputState }

func (UniqueKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UniqueKey)(nil)).Elem()
}

func (o UniqueKeyArrayOutput) ToUniqueKeyArrayOutput() UniqueKeyArrayOutput {
	return o
}

func (o UniqueKeyArrayOutput) ToUniqueKeyArrayOutputWithContext(ctx context.Context) UniqueKeyArrayOutput {
	return o
}

func (o UniqueKeyArrayOutput) Index(i pulumi.IntInput) UniqueKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UniqueKey {
		return vs[0].([]UniqueKey)[vs[1].(int)]
	}).(UniqueKeyOutput)
}

type UniqueKeyPolicy struct {
	UniqueKeys []UniqueKey `pulumi:"uniqueKeys"`
}





type UniqueKeyPolicyInput interface {
	pulumi.Input

	ToUniqueKeyPolicyOutput() UniqueKeyPolicyOutput
	ToUniqueKeyPolicyOutputWithContext(context.Context) UniqueKeyPolicyOutput
}

type UniqueKeyPolicyArgs struct {
	UniqueKeys UniqueKeyArrayInput `pulumi:"uniqueKeys"`
}

func (UniqueKeyPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UniqueKeyPolicy)(nil)).Elem()
}

func (i UniqueKeyPolicyArgs) ToUniqueKeyPolicyOutput() UniqueKeyPolicyOutput {
	return i.ToUniqueKeyPolicyOutputWithContext(context.Background())
}

func (i UniqueKeyPolicyArgs) ToUniqueKeyPolicyOutputWithContext(ctx context.Context) UniqueKeyPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UniqueKeyPolicyOutput)
}

func (i UniqueKeyPolicyArgs) ToUniqueKeyPolicyPtrOutput() UniqueKeyPolicyPtrOutput {
	return i.ToUniqueKeyPolicyPtrOutputWithContext(context.Background())
}

func (i UniqueKeyPolicyArgs) ToUniqueKeyPolicyPtrOutputWithContext(ctx context.Context) UniqueKeyPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UniqueKeyPolicyOutput).ToUniqueKeyPolicyPtrOutputWithContext(ctx)
}









type UniqueKeyPolicyPtrInput interface {
	pulumi.Input

	ToUniqueKeyPolicyPtrOutput() UniqueKeyPolicyPtrOutput
	ToUniqueKeyPolicyPtrOutputWithContext(context.Context) UniqueKeyPolicyPtrOutput
}

type uniqueKeyPolicyPtrType UniqueKeyPolicyArgs

func UniqueKeyPolicyPtr(v *UniqueKeyPolicyArgs) UniqueKeyPolicyPtrInput {
	return (*uniqueKeyPolicyPtrType)(v)
}

func (*uniqueKeyPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UniqueKeyPolicy)(nil)).Elem()
}

func (i *uniqueKeyPolicyPtrType) ToUniqueKeyPolicyPtrOutput() UniqueKeyPolicyPtrOutput {
	return i.ToUniqueKeyPolicyPtrOutputWithContext(context.Background())
}

func (i *uniqueKeyPolicyPtrType) ToUniqueKeyPolicyPtrOutputWithContext(ctx context.Context) UniqueKeyPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UniqueKeyPolicyPtrOutput)
}

type UniqueKeyPolicyOutput struct{ *pulumi.OutputState }

func (UniqueKeyPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UniqueKeyPolicy)(nil)).Elem()
}

func (o UniqueKeyPolicyOutput) ToUniqueKeyPolicyOutput() UniqueKeyPolicyOutput {
	return o
}

func (o UniqueKeyPolicyOutput) ToUniqueKeyPolicyOutputWithContext(ctx context.Context) UniqueKeyPolicyOutput {
	return o
}

func (o UniqueKeyPolicyOutput) ToUniqueKeyPolicyPtrOutput() UniqueKeyPolicyPtrOutput {
	return o.ToUniqueKeyPolicyPtrOutputWithContext(context.Background())
}

func (o UniqueKeyPolicyOutput) ToUniqueKeyPolicyPtrOutputWithContext(ctx context.Context) UniqueKeyPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UniqueKeyPolicy) *UniqueKeyPolicy {
		return &v
	}).(UniqueKeyPolicyPtrOutput)
}

func (o UniqueKeyPolicyOutput) UniqueKeys() UniqueKeyArrayOutput {
	return o.ApplyT(func(v UniqueKeyPolicy) []UniqueKey { return v.UniqueKeys }).(UniqueKeyArrayOutput)
}

type UniqueKeyPolicyPtrOutput struct{ *pulumi.OutputState }

func (UniqueKeyPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UniqueKeyPolicy)(nil)).Elem()
}

func (o UniqueKeyPolicyPtrOutput) ToUniqueKeyPolicyPtrOutput() UniqueKeyPolicyPtrOutput {
	return o
}

func (o UniqueKeyPolicyPtrOutput) ToUniqueKeyPolicyPtrOutputWithContext(ctx context.Context) UniqueKeyPolicyPtrOutput {
	return o
}

func (o UniqueKeyPolicyPtrOutput) Elem() UniqueKeyPolicyOutput {
	return o.ApplyT(func(v *UniqueKeyPolicy) UniqueKeyPolicy {
		if v != nil {
			return *v
		}
		var ret UniqueKeyPolicy
		return ret
	}).(UniqueKeyPolicyOutput)
}

func (o UniqueKeyPolicyPtrOutput) UniqueKeys() UniqueKeyArrayOutput {
	return o.ApplyT(func(v *UniqueKeyPolicy) []UniqueKey {
		if v == nil {
			return nil
		}
		return v.UniqueKeys
	}).(UniqueKeyArrayOutput)
}

type UniqueKeyPolicyResponse struct {
	UniqueKeys []UniqueKeyResponse `pulumi:"uniqueKeys"`
}

type UniqueKeyPolicyResponseOutput struct{ *pulumi.OutputState }

func (UniqueKeyPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UniqueKeyPolicyResponse)(nil)).Elem()
}

func (o UniqueKeyPolicyResponseOutput) ToUniqueKeyPolicyResponseOutput() UniqueKeyPolicyResponseOutput {
	return o
}

func (o UniqueKeyPolicyResponseOutput) ToUniqueKeyPolicyResponseOutputWithContext(ctx context.Context) UniqueKeyPolicyResponseOutput {
	return o
}

func (o UniqueKeyPolicyResponseOutput) UniqueKeys() UniqueKeyResponseArrayOutput {
	return o.ApplyT(func(v UniqueKeyPolicyResponse) []UniqueKeyResponse { return v.UniqueKeys }).(UniqueKeyResponseArrayOutput)
}

type UniqueKeyPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (UniqueKeyPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UniqueKeyPolicyResponse)(nil)).Elem()
}

func (o UniqueKeyPolicyResponsePtrOutput) ToUniqueKeyPolicyResponsePtrOutput() UniqueKeyPolicyResponsePtrOutput {
	return o
}

func (o UniqueKeyPolicyResponsePtrOutput) ToUniqueKeyPolicyResponsePtrOutputWithContext(ctx context.Context) UniqueKeyPolicyResponsePtrOutput {
	return o
}

func (o UniqueKeyPolicyResponsePtrOutput) Elem() UniqueKeyPolicyResponseOutput {
	return o.ApplyT(func(v *UniqueKeyPolicyResponse) UniqueKeyPolicyResponse {
		if v != nil {
			return *v
		}
		var ret UniqueKeyPolicyResponse
		return ret
	}).(UniqueKeyPolicyResponseOutput)
}

func (o UniqueKeyPolicyResponsePtrOutput) UniqueKeys() UniqueKeyResponseArrayOutput {
	return o.ApplyT(func(v *UniqueKeyPolicyResponse) []UniqueKeyResponse {
		if v == nil {
			return nil
		}
		return v.UniqueKeys
	}).(UniqueKeyResponseArrayOutput)
}

type UniqueKeyResponse struct {
	Paths []string `pulumi:"paths"`
}

type UniqueKeyResponseOutput struct{ *pulumi.OutputState }

func (UniqueKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UniqueKeyResponse)(nil)).Elem()
}

func (o UniqueKeyResponseOutput) ToUniqueKeyResponseOutput() UniqueKeyResponseOutput {
	return o
}

func (o UniqueKeyResponseOutput) ToUniqueKeyResponseOutputWithContext(ctx context.Context) UniqueKeyResponseOutput {
	return o
}

func (o UniqueKeyResponseOutput) Paths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v UniqueKeyResponse) []string { return v.Paths }).(pulumi.StringArrayOutput)
}

type UniqueKeyResponseArrayOutput struct{ *pulumi.OutputState }

func (UniqueKeyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UniqueKeyResponse)(nil)).Elem()
}

func (o UniqueKeyResponseArrayOutput) ToUniqueKeyResponseArrayOutput() UniqueKeyResponseArrayOutput {
	return o
}

func (o UniqueKeyResponseArrayOutput) ToUniqueKeyResponseArrayOutputWithContext(ctx context.Context) UniqueKeyResponseArrayOutput {
	return o
}

func (o UniqueKeyResponseArrayOutput) Index(i pulumi.IntInput) UniqueKeyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UniqueKeyResponse {
		return vs[0].([]UniqueKeyResponse)[vs[1].(int)]
	}).(UniqueKeyResponseOutput)
}

type VirtualNetworkRule struct {
	Id                               *string `pulumi:"id"`
	IgnoreMissingVNetServiceEndpoint *bool   `pulumi:"ignoreMissingVNetServiceEndpoint"`
}





type VirtualNetworkRuleInput interface {
	pulumi.Input

	ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput
	ToVirtualNetworkRuleOutputWithContext(context.Context) VirtualNetworkRuleOutput
}

type VirtualNetworkRuleArgs struct {
	Id                               pulumi.StringPtrInput `pulumi:"id"`
	IgnoreMissingVNetServiceEndpoint pulumi.BoolPtrInput   `pulumi:"ignoreMissingVNetServiceEndpoint"`
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

func (o VirtualNetworkRuleOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkRule) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkRuleOutput) IgnoreMissingVNetServiceEndpoint() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkRule) *bool { return v.IgnoreMissingVNetServiceEndpoint }).(pulumi.BoolPtrOutput)
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
	Id                               *string `pulumi:"id"`
	IgnoreMissingVNetServiceEndpoint *bool   `pulumi:"ignoreMissingVNetServiceEndpoint"`
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

func (o VirtualNetworkRuleResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkRuleResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkRuleResponseOutput) IgnoreMissingVNetServiceEndpoint() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkRuleResponse) *bool { return v.IgnoreMissingVNetServiceEndpoint }).(pulumi.BoolPtrOutput)
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

type CompositePathArrayArray []CompositePathArrayInput

func (CompositePathArrayArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[][]CompositePath)(nil)).Elem()
}

func (i CompositePathArrayArray) ToCompositePathArrayArrayOutput() CompositePathArrayArrayOutput {
	return i.ToCompositePathArrayArrayOutputWithContext(context.Background())
}

func (i CompositePathArrayArray) ToCompositePathArrayArrayOutputWithContext(ctx context.Context) CompositePathArrayArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CompositePathArrayArrayOutput)
}





type CompositePathArrayArrayInput interface {
	pulumi.Input

	ToCompositePathArrayArrayOutput() CompositePathArrayArrayOutput
	ToCompositePathArrayArrayOutputWithContext(context.Context) CompositePathArrayArrayOutput
}

type CompositePathArrayArrayOutput struct{ *pulumi.OutputState }

func (CompositePathArrayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[][]CompositePath)(nil)).Elem()
}

func (o CompositePathArrayArrayOutput) ToCompositePathArrayArrayOutput() CompositePathArrayArrayOutput {
	return o
}

func (o CompositePathArrayArrayOutput) ToCompositePathArrayArrayOutputWithContext(ctx context.Context) CompositePathArrayArrayOutput {
	return o
}

func (o CompositePathArrayArrayOutput) Index(i pulumi.IntInput) CompositePathArrayOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) []CompositePath {
		return vs[0].([][]CompositePath)[vs[1].(int)]
	}).(CompositePathArrayOutput)
}

type CompositePathResponseArrayArrayOutput struct{ *pulumi.OutputState }

func (CompositePathResponseArrayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[][]CompositePathResponse)(nil)).Elem()
}

func (o CompositePathResponseArrayArrayOutput) ToCompositePathResponseArrayArrayOutput() CompositePathResponseArrayArrayOutput {
	return o
}

func (o CompositePathResponseArrayArrayOutput) ToCompositePathResponseArrayArrayOutputWithContext(ctx context.Context) CompositePathResponseArrayArrayOutput {
	return o
}

func (o CompositePathResponseArrayArrayOutput) Index(i pulumi.IntInput) CompositePathResponseArrayOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) []CompositePathResponse {
		return vs[0].([][]CompositePathResponse)[vs[1].(int)]
	}).(CompositePathResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ApiPropertiesOutput{})
	pulumi.RegisterOutputType(ApiPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ApiPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ApiPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(AutoscaleSettingsOutput{})
	pulumi.RegisterOutputType(AutoscaleSettingsPtrOutput{})
	pulumi.RegisterOutputType(AutoscaleSettingsResponseOutput{})
	pulumi.RegisterOutputType(AutoscaleSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(CapabilityOutput{})
	pulumi.RegisterOutputType(CapabilityArrayOutput{})
	pulumi.RegisterOutputType(CapabilityResponseOutput{})
	pulumi.RegisterOutputType(CapabilityResponseArrayOutput{})
	pulumi.RegisterOutputType(CassandraKeyspaceGetPropertiesResponseOptionsOutput{})
	pulumi.RegisterOutputType(CassandraKeyspaceGetPropertiesResponseOptionsPtrOutput{})
	pulumi.RegisterOutputType(CassandraKeyspaceGetPropertiesResponseResourceOutput{})
	pulumi.RegisterOutputType(CassandraKeyspaceGetPropertiesResponseResourcePtrOutput{})
	pulumi.RegisterOutputType(CassandraKeyspaceResourceOutput{})
	pulumi.RegisterOutputType(CassandraPartitionKeyOutput{})
	pulumi.RegisterOutputType(CassandraPartitionKeyArrayOutput{})
	pulumi.RegisterOutputType(CassandraPartitionKeyResponseOutput{})
	pulumi.RegisterOutputType(CassandraPartitionKeyResponseArrayOutput{})
	pulumi.RegisterOutputType(CassandraSchemaOutput{})
	pulumi.RegisterOutputType(CassandraSchemaPtrOutput{})
	pulumi.RegisterOutputType(CassandraSchemaResponseOutput{})
	pulumi.RegisterOutputType(CassandraSchemaResponsePtrOutput{})
	pulumi.RegisterOutputType(CassandraTableGetPropertiesResponseOptionsOutput{})
	pulumi.RegisterOutputType(CassandraTableGetPropertiesResponseOptionsPtrOutput{})
	pulumi.RegisterOutputType(CassandraTableGetPropertiesResponseResourceOutput{})
	pulumi.RegisterOutputType(CassandraTableGetPropertiesResponseResourcePtrOutput{})
	pulumi.RegisterOutputType(CassandraTableResourceOutput{})
	pulumi.RegisterOutputType(ClusterKeyOutput{})
	pulumi.RegisterOutputType(ClusterKeyArrayOutput{})
	pulumi.RegisterOutputType(ClusterKeyResponseOutput{})
	pulumi.RegisterOutputType(ClusterKeyResponseArrayOutput{})
	pulumi.RegisterOutputType(ColumnOutput{})
	pulumi.RegisterOutputType(ColumnArrayOutput{})
	pulumi.RegisterOutputType(ColumnResponseOutput{})
	pulumi.RegisterOutputType(ColumnResponseArrayOutput{})
	pulumi.RegisterOutputType(CompositePathOutput{})
	pulumi.RegisterOutputType(CompositePathArrayOutput{})
	pulumi.RegisterOutputType(CompositePathResponseOutput{})
	pulumi.RegisterOutputType(CompositePathResponseArrayOutput{})
	pulumi.RegisterOutputType(ConflictResolutionPolicyOutput{})
	pulumi.RegisterOutputType(ConflictResolutionPolicyPtrOutput{})
	pulumi.RegisterOutputType(ConflictResolutionPolicyResponseOutput{})
	pulumi.RegisterOutputType(ConflictResolutionPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(ConsistencyPolicyOutput{})
	pulumi.RegisterOutputType(ConsistencyPolicyPtrOutput{})
	pulumi.RegisterOutputType(ConsistencyPolicyResponseOutput{})
	pulumi.RegisterOutputType(ConsistencyPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerPartitionKeyOutput{})
	pulumi.RegisterOutputType(ContainerPartitionKeyPtrOutput{})
	pulumi.RegisterOutputType(ContainerPartitionKeyResponseOutput{})
	pulumi.RegisterOutputType(ContainerPartitionKeyResponsePtrOutput{})
	pulumi.RegisterOutputType(CorsPolicyOutput{})
	pulumi.RegisterOutputType(CorsPolicyArrayOutput{})
	pulumi.RegisterOutputType(CorsPolicyResponseOutput{})
	pulumi.RegisterOutputType(CorsPolicyResponseArrayOutput{})
	pulumi.RegisterOutputType(CreateUpdateOptionsOutput{})
	pulumi.RegisterOutputType(CreateUpdateOptionsPtrOutput{})
	pulumi.RegisterOutputType(ExcludedPathOutput{})
	pulumi.RegisterOutputType(ExcludedPathArrayOutput{})
	pulumi.RegisterOutputType(ExcludedPathResponseOutput{})
	pulumi.RegisterOutputType(ExcludedPathResponseArrayOutput{})
	pulumi.RegisterOutputType(FailoverPolicyResponseOutput{})
	pulumi.RegisterOutputType(FailoverPolicyResponseArrayOutput{})
	pulumi.RegisterOutputType(GremlinDatabaseGetPropertiesResponseOptionsOutput{})
	pulumi.RegisterOutputType(GremlinDatabaseGetPropertiesResponseOptionsPtrOutput{})
	pulumi.RegisterOutputType(GremlinDatabaseGetPropertiesResponseResourceOutput{})
	pulumi.RegisterOutputType(GremlinDatabaseGetPropertiesResponseResourcePtrOutput{})
	pulumi.RegisterOutputType(GremlinDatabaseResourceOutput{})
	pulumi.RegisterOutputType(GremlinGraphGetPropertiesResponseOptionsOutput{})
	pulumi.RegisterOutputType(GremlinGraphGetPropertiesResponseOptionsPtrOutput{})
	pulumi.RegisterOutputType(GremlinGraphGetPropertiesResponseResourceOutput{})
	pulumi.RegisterOutputType(GremlinGraphGetPropertiesResponseResourcePtrOutput{})
	pulumi.RegisterOutputType(GremlinGraphResourceOutput{})
	pulumi.RegisterOutputType(IncludedPathOutput{})
	pulumi.RegisterOutputType(IncludedPathArrayOutput{})
	pulumi.RegisterOutputType(IncludedPathResponseOutput{})
	pulumi.RegisterOutputType(IncludedPathResponseArrayOutput{})
	pulumi.RegisterOutputType(IndexesOutput{})
	pulumi.RegisterOutputType(IndexesArrayOutput{})
	pulumi.RegisterOutputType(IndexesResponseOutput{})
	pulumi.RegisterOutputType(IndexesResponseArrayOutput{})
	pulumi.RegisterOutputType(IndexingPolicyOutput{})
	pulumi.RegisterOutputType(IndexingPolicyPtrOutput{})
	pulumi.RegisterOutputType(IndexingPolicyResponseOutput{})
	pulumi.RegisterOutputType(IndexingPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(IpAddressOrRangeOutput{})
	pulumi.RegisterOutputType(IpAddressOrRangeArrayOutput{})
	pulumi.RegisterOutputType(IpAddressOrRangeResponseOutput{})
	pulumi.RegisterOutputType(IpAddressOrRangeResponseArrayOutput{})
	pulumi.RegisterOutputType(LocationOutput{})
	pulumi.RegisterOutputType(LocationArrayOutput{})
	pulumi.RegisterOutputType(LocationResponseOutput{})
	pulumi.RegisterOutputType(LocationResponseArrayOutput{})
	pulumi.RegisterOutputType(MongoDBCollectionGetPropertiesResponseOptionsOutput{})
	pulumi.RegisterOutputType(MongoDBCollectionGetPropertiesResponseOptionsPtrOutput{})
	pulumi.RegisterOutputType(MongoDBCollectionGetPropertiesResponseResourceOutput{})
	pulumi.RegisterOutputType(MongoDBCollectionGetPropertiesResponseResourcePtrOutput{})
	pulumi.RegisterOutputType(MongoDBCollectionResourceOutput{})
	pulumi.RegisterOutputType(MongoDBDatabaseGetPropertiesResponseOptionsOutput{})
	pulumi.RegisterOutputType(MongoDBDatabaseGetPropertiesResponseOptionsPtrOutput{})
	pulumi.RegisterOutputType(MongoDBDatabaseGetPropertiesResponseResourceOutput{})
	pulumi.RegisterOutputType(MongoDBDatabaseGetPropertiesResponseResourcePtrOutput{})
	pulumi.RegisterOutputType(MongoDBDatabaseResourceOutput{})
	pulumi.RegisterOutputType(MongoIndexOutput{})
	pulumi.RegisterOutputType(MongoIndexArrayOutput{})
	pulumi.RegisterOutputType(MongoIndexKeysOutput{})
	pulumi.RegisterOutputType(MongoIndexKeysPtrOutput{})
	pulumi.RegisterOutputType(MongoIndexKeysResponseOutput{})
	pulumi.RegisterOutputType(MongoIndexKeysResponsePtrOutput{})
	pulumi.RegisterOutputType(MongoIndexOptionsOutput{})
	pulumi.RegisterOutputType(MongoIndexOptionsPtrOutput{})
	pulumi.RegisterOutputType(MongoIndexOptionsResponseOutput{})
	pulumi.RegisterOutputType(MongoIndexOptionsResponsePtrOutput{})
	pulumi.RegisterOutputType(MongoIndexResponseOutput{})
	pulumi.RegisterOutputType(MongoIndexResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPropertyResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePropertyResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(SpatialSpecOutput{})
	pulumi.RegisterOutputType(SpatialSpecArrayOutput{})
	pulumi.RegisterOutputType(SpatialSpecResponseOutput{})
	pulumi.RegisterOutputType(SpatialSpecResponseArrayOutput{})
	pulumi.RegisterOutputType(SqlContainerGetPropertiesResponseOptionsOutput{})
	pulumi.RegisterOutputType(SqlContainerGetPropertiesResponseOptionsPtrOutput{})
	pulumi.RegisterOutputType(SqlContainerGetPropertiesResponseResourceOutput{})
	pulumi.RegisterOutputType(SqlContainerGetPropertiesResponseResourcePtrOutput{})
	pulumi.RegisterOutputType(SqlContainerResourceOutput{})
	pulumi.RegisterOutputType(SqlDatabaseGetPropertiesResponseOptionsOutput{})
	pulumi.RegisterOutputType(SqlDatabaseGetPropertiesResponseOptionsPtrOutput{})
	pulumi.RegisterOutputType(SqlDatabaseGetPropertiesResponseResourceOutput{})
	pulumi.RegisterOutputType(SqlDatabaseGetPropertiesResponseResourcePtrOutput{})
	pulumi.RegisterOutputType(SqlDatabaseResourceOutput{})
	pulumi.RegisterOutputType(SqlStoredProcedureGetPropertiesResponseResourceOutput{})
	pulumi.RegisterOutputType(SqlStoredProcedureGetPropertiesResponseResourcePtrOutput{})
	pulumi.RegisterOutputType(SqlStoredProcedureResourceOutput{})
	pulumi.RegisterOutputType(SqlTriggerGetPropertiesResponseResourceOutput{})
	pulumi.RegisterOutputType(SqlTriggerGetPropertiesResponseResourcePtrOutput{})
	pulumi.RegisterOutputType(SqlTriggerResourceOutput{})
	pulumi.RegisterOutputType(SqlUserDefinedFunctionGetPropertiesResponseResourceOutput{})
	pulumi.RegisterOutputType(SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput{})
	pulumi.RegisterOutputType(SqlUserDefinedFunctionResourceOutput{})
	pulumi.RegisterOutputType(TableGetPropertiesResponseOptionsOutput{})
	pulumi.RegisterOutputType(TableGetPropertiesResponseOptionsPtrOutput{})
	pulumi.RegisterOutputType(TableGetPropertiesResponseResourceOutput{})
	pulumi.RegisterOutputType(TableGetPropertiesResponseResourcePtrOutput{})
	pulumi.RegisterOutputType(TableResourceOutput{})
	pulumi.RegisterOutputType(UniqueKeyOutput{})
	pulumi.RegisterOutputType(UniqueKeyArrayOutput{})
	pulumi.RegisterOutputType(UniqueKeyPolicyOutput{})
	pulumi.RegisterOutputType(UniqueKeyPolicyPtrOutput{})
	pulumi.RegisterOutputType(UniqueKeyPolicyResponseOutput{})
	pulumi.RegisterOutputType(UniqueKeyPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(UniqueKeyResponseOutput{})
	pulumi.RegisterOutputType(UniqueKeyResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(CompositePathArrayArrayOutput{})
	pulumi.RegisterOutputType(CompositePathResponseArrayArrayOutput{})
}
