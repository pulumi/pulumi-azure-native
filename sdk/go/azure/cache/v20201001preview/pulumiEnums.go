


package v20201001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ClusteringPolicy string

const (
	ClusteringPolicyEnterpriseCluster = ClusteringPolicy("EnterpriseCluster")
	ClusteringPolicyOSSCluster        = ClusteringPolicy("OSSCluster")
)

func (ClusteringPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusteringPolicy)(nil)).Elem()
}

func (e ClusteringPolicy) ToClusteringPolicyOutput() ClusteringPolicyOutput {
	return pulumi.ToOutput(e).(ClusteringPolicyOutput)
}

func (e ClusteringPolicy) ToClusteringPolicyOutputWithContext(ctx context.Context) ClusteringPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ClusteringPolicyOutput)
}

func (e ClusteringPolicy) ToClusteringPolicyPtrOutput() ClusteringPolicyPtrOutput {
	return e.ToClusteringPolicyPtrOutputWithContext(context.Background())
}

func (e ClusteringPolicy) ToClusteringPolicyPtrOutputWithContext(ctx context.Context) ClusteringPolicyPtrOutput {
	return ClusteringPolicy(e).ToClusteringPolicyOutputWithContext(ctx).ToClusteringPolicyPtrOutputWithContext(ctx)
}

func (e ClusteringPolicy) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ClusteringPolicy) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ClusteringPolicy) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ClusteringPolicy) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ClusteringPolicyOutput struct{ *pulumi.OutputState }

func (ClusteringPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusteringPolicy)(nil)).Elem()
}

func (o ClusteringPolicyOutput) ToClusteringPolicyOutput() ClusteringPolicyOutput {
	return o
}

func (o ClusteringPolicyOutput) ToClusteringPolicyOutputWithContext(ctx context.Context) ClusteringPolicyOutput {
	return o
}

func (o ClusteringPolicyOutput) ToClusteringPolicyPtrOutput() ClusteringPolicyPtrOutput {
	return o.ToClusteringPolicyPtrOutputWithContext(context.Background())
}

func (o ClusteringPolicyOutput) ToClusteringPolicyPtrOutputWithContext(ctx context.Context) ClusteringPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusteringPolicy) *ClusteringPolicy {
		return &v
	}).(ClusteringPolicyPtrOutput)
}

func (o ClusteringPolicyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ClusteringPolicyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ClusteringPolicy) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ClusteringPolicyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ClusteringPolicyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ClusteringPolicy) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ClusteringPolicyPtrOutput struct{ *pulumi.OutputState }

func (ClusteringPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusteringPolicy)(nil)).Elem()
}

func (o ClusteringPolicyPtrOutput) ToClusteringPolicyPtrOutput() ClusteringPolicyPtrOutput {
	return o
}

func (o ClusteringPolicyPtrOutput) ToClusteringPolicyPtrOutputWithContext(ctx context.Context) ClusteringPolicyPtrOutput {
	return o
}

func (o ClusteringPolicyPtrOutput) Elem() ClusteringPolicyOutput {
	return o.ApplyT(func(v *ClusteringPolicy) ClusteringPolicy {
		if v != nil {
			return *v
		}
		var ret ClusteringPolicy
		return ret
	}).(ClusteringPolicyOutput)
}

func (o ClusteringPolicyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ClusteringPolicyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ClusteringPolicy) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ClusteringPolicyInput interface {
	pulumi.Input

	ToClusteringPolicyOutput() ClusteringPolicyOutput
	ToClusteringPolicyOutputWithContext(context.Context) ClusteringPolicyOutput
}

var clusteringPolicyPtrType = reflect.TypeOf((**ClusteringPolicy)(nil)).Elem()

type ClusteringPolicyPtrInput interface {
	pulumi.Input

	ToClusteringPolicyPtrOutput() ClusteringPolicyPtrOutput
	ToClusteringPolicyPtrOutputWithContext(context.Context) ClusteringPolicyPtrOutput
}

type clusteringPolicyPtr string

func ClusteringPolicyPtr(v string) ClusteringPolicyPtrInput {
	return (*clusteringPolicyPtr)(&v)
}

func (*clusteringPolicyPtr) ElementType() reflect.Type {
	return clusteringPolicyPtrType
}

func (in *clusteringPolicyPtr) ToClusteringPolicyPtrOutput() ClusteringPolicyPtrOutput {
	return pulumi.ToOutput(in).(ClusteringPolicyPtrOutput)
}

func (in *clusteringPolicyPtr) ToClusteringPolicyPtrOutputWithContext(ctx context.Context) ClusteringPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ClusteringPolicyPtrOutput)
}

type EvictionPolicy string

const (
	EvictionPolicyAllKeysLFU     = EvictionPolicy("AllKeysLFU")
	EvictionPolicyAllKeysLRU     = EvictionPolicy("AllKeysLRU")
	EvictionPolicyAllKeysRandom  = EvictionPolicy("AllKeysRandom")
	EvictionPolicyVolatileLRU    = EvictionPolicy("VolatileLRU")
	EvictionPolicyVolatileLFU    = EvictionPolicy("VolatileLFU")
	EvictionPolicyVolatileTTL    = EvictionPolicy("VolatileTTL")
	EvictionPolicyVolatileRandom = EvictionPolicy("VolatileRandom")
	EvictionPolicyNoEviction     = EvictionPolicy("NoEviction")
)

func (EvictionPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*EvictionPolicy)(nil)).Elem()
}

func (e EvictionPolicy) ToEvictionPolicyOutput() EvictionPolicyOutput {
	return pulumi.ToOutput(e).(EvictionPolicyOutput)
}

func (e EvictionPolicy) ToEvictionPolicyOutputWithContext(ctx context.Context) EvictionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EvictionPolicyOutput)
}

func (e EvictionPolicy) ToEvictionPolicyPtrOutput() EvictionPolicyPtrOutput {
	return e.ToEvictionPolicyPtrOutputWithContext(context.Background())
}

func (e EvictionPolicy) ToEvictionPolicyPtrOutputWithContext(ctx context.Context) EvictionPolicyPtrOutput {
	return EvictionPolicy(e).ToEvictionPolicyOutputWithContext(ctx).ToEvictionPolicyPtrOutputWithContext(ctx)
}

func (e EvictionPolicy) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EvictionPolicy) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EvictionPolicy) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EvictionPolicy) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EvictionPolicyOutput struct{ *pulumi.OutputState }

func (EvictionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EvictionPolicy)(nil)).Elem()
}

func (o EvictionPolicyOutput) ToEvictionPolicyOutput() EvictionPolicyOutput {
	return o
}

func (o EvictionPolicyOutput) ToEvictionPolicyOutputWithContext(ctx context.Context) EvictionPolicyOutput {
	return o
}

func (o EvictionPolicyOutput) ToEvictionPolicyPtrOutput() EvictionPolicyPtrOutput {
	return o.ToEvictionPolicyPtrOutputWithContext(context.Background())
}

func (o EvictionPolicyOutput) ToEvictionPolicyPtrOutputWithContext(ctx context.Context) EvictionPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EvictionPolicy) *EvictionPolicy {
		return &v
	}).(EvictionPolicyPtrOutput)
}

func (o EvictionPolicyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EvictionPolicyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EvictionPolicy) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EvictionPolicyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EvictionPolicyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EvictionPolicy) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EvictionPolicyPtrOutput struct{ *pulumi.OutputState }

func (EvictionPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EvictionPolicy)(nil)).Elem()
}

func (o EvictionPolicyPtrOutput) ToEvictionPolicyPtrOutput() EvictionPolicyPtrOutput {
	return o
}

func (o EvictionPolicyPtrOutput) ToEvictionPolicyPtrOutputWithContext(ctx context.Context) EvictionPolicyPtrOutput {
	return o
}

func (o EvictionPolicyPtrOutput) Elem() EvictionPolicyOutput {
	return o.ApplyT(func(v *EvictionPolicy) EvictionPolicy {
		if v != nil {
			return *v
		}
		var ret EvictionPolicy
		return ret
	}).(EvictionPolicyOutput)
}

func (o EvictionPolicyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EvictionPolicyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EvictionPolicy) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EvictionPolicyInput interface {
	pulumi.Input

	ToEvictionPolicyOutput() EvictionPolicyOutput
	ToEvictionPolicyOutputWithContext(context.Context) EvictionPolicyOutput
}

var evictionPolicyPtrType = reflect.TypeOf((**EvictionPolicy)(nil)).Elem()

type EvictionPolicyPtrInput interface {
	pulumi.Input

	ToEvictionPolicyPtrOutput() EvictionPolicyPtrOutput
	ToEvictionPolicyPtrOutputWithContext(context.Context) EvictionPolicyPtrOutput
}

type evictionPolicyPtr string

func EvictionPolicyPtr(v string) EvictionPolicyPtrInput {
	return (*evictionPolicyPtr)(&v)
}

func (*evictionPolicyPtr) ElementType() reflect.Type {
	return evictionPolicyPtrType
}

func (in *evictionPolicyPtr) ToEvictionPolicyPtrOutput() EvictionPolicyPtrOutput {
	return pulumi.ToOutput(in).(EvictionPolicyPtrOutput)
}

func (in *evictionPolicyPtr) ToEvictionPolicyPtrOutputWithContext(ctx context.Context) EvictionPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EvictionPolicyPtrOutput)
}

type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusPending  = PrivateEndpointServiceConnectionStatus("Pending")
	PrivateEndpointServiceConnectionStatusApproved = PrivateEndpointServiceConnectionStatus("Approved")
	PrivateEndpointServiceConnectionStatusRejected = PrivateEndpointServiceConnectionStatus("Rejected")
)

func (PrivateEndpointServiceConnectionStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointServiceConnectionStatus)(nil)).Elem()
}

func (e PrivateEndpointServiceConnectionStatus) ToPrivateEndpointServiceConnectionStatusOutput() PrivateEndpointServiceConnectionStatusOutput {
	return pulumi.ToOutput(e).(PrivateEndpointServiceConnectionStatusOutput)
}

func (e PrivateEndpointServiceConnectionStatus) ToPrivateEndpointServiceConnectionStatusOutputWithContext(ctx context.Context) PrivateEndpointServiceConnectionStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PrivateEndpointServiceConnectionStatusOutput)
}

func (e PrivateEndpointServiceConnectionStatus) ToPrivateEndpointServiceConnectionStatusPtrOutput() PrivateEndpointServiceConnectionStatusPtrOutput {
	return e.ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(context.Background())
}

func (e PrivateEndpointServiceConnectionStatus) ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateEndpointServiceConnectionStatusPtrOutput {
	return PrivateEndpointServiceConnectionStatus(e).ToPrivateEndpointServiceConnectionStatusOutputWithContext(ctx).ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(ctx)
}

func (e PrivateEndpointServiceConnectionStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrivateEndpointServiceConnectionStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrivateEndpointServiceConnectionStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PrivateEndpointServiceConnectionStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PrivateEndpointServiceConnectionStatusOutput struct{ *pulumi.OutputState }

func (PrivateEndpointServiceConnectionStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointServiceConnectionStatus)(nil)).Elem()
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToPrivateEndpointServiceConnectionStatusOutput() PrivateEndpointServiceConnectionStatusOutput {
	return o
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToPrivateEndpointServiceConnectionStatusOutputWithContext(ctx context.Context) PrivateEndpointServiceConnectionStatusOutput {
	return o
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToPrivateEndpointServiceConnectionStatusPtrOutput() PrivateEndpointServiceConnectionStatusPtrOutput {
	return o.ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateEndpointServiceConnectionStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointServiceConnectionStatus) *PrivateEndpointServiceConnectionStatus {
		return &v
	}).(PrivateEndpointServiceConnectionStatusPtrOutput)
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrivateEndpointServiceConnectionStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrivateEndpointServiceConnectionStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointServiceConnectionStatusPtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointServiceConnectionStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointServiceConnectionStatus)(nil)).Elem()
}

func (o PrivateEndpointServiceConnectionStatusPtrOutput) ToPrivateEndpointServiceConnectionStatusPtrOutput() PrivateEndpointServiceConnectionStatusPtrOutput {
	return o
}

func (o PrivateEndpointServiceConnectionStatusPtrOutput) ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateEndpointServiceConnectionStatusPtrOutput {
	return o
}

func (o PrivateEndpointServiceConnectionStatusPtrOutput) Elem() PrivateEndpointServiceConnectionStatusOutput {
	return o.ApplyT(func(v *PrivateEndpointServiceConnectionStatus) PrivateEndpointServiceConnectionStatus {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointServiceConnectionStatus
		return ret
	}).(PrivateEndpointServiceConnectionStatusOutput)
}

func (o PrivateEndpointServiceConnectionStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointServiceConnectionStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PrivateEndpointServiceConnectionStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PrivateEndpointServiceConnectionStatusInput interface {
	pulumi.Input

	ToPrivateEndpointServiceConnectionStatusOutput() PrivateEndpointServiceConnectionStatusOutput
	ToPrivateEndpointServiceConnectionStatusOutputWithContext(context.Context) PrivateEndpointServiceConnectionStatusOutput
}

var privateEndpointServiceConnectionStatusPtrType = reflect.TypeOf((**PrivateEndpointServiceConnectionStatus)(nil)).Elem()

type PrivateEndpointServiceConnectionStatusPtrInput interface {
	pulumi.Input

	ToPrivateEndpointServiceConnectionStatusPtrOutput() PrivateEndpointServiceConnectionStatusPtrOutput
	ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(context.Context) PrivateEndpointServiceConnectionStatusPtrOutput
}

type privateEndpointServiceConnectionStatusPtr string

func PrivateEndpointServiceConnectionStatusPtr(v string) PrivateEndpointServiceConnectionStatusPtrInput {
	return (*privateEndpointServiceConnectionStatusPtr)(&v)
}

func (*privateEndpointServiceConnectionStatusPtr) ElementType() reflect.Type {
	return privateEndpointServiceConnectionStatusPtrType
}

func (in *privateEndpointServiceConnectionStatusPtr) ToPrivateEndpointServiceConnectionStatusPtrOutput() PrivateEndpointServiceConnectionStatusPtrOutput {
	return pulumi.ToOutput(in).(PrivateEndpointServiceConnectionStatusPtrOutput)
}

func (in *privateEndpointServiceConnectionStatusPtr) ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateEndpointServiceConnectionStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PrivateEndpointServiceConnectionStatusPtrOutput)
}

type Protocol string

const (
	ProtocolEncrypted = Protocol("Encrypted")
	ProtocolPlaintext = Protocol("Plaintext")
)

func (Protocol) ElementType() reflect.Type {
	return reflect.TypeOf((*Protocol)(nil)).Elem()
}

func (e Protocol) ToProtocolOutput() ProtocolOutput {
	return pulumi.ToOutput(e).(ProtocolOutput)
}

func (e Protocol) ToProtocolOutputWithContext(ctx context.Context) ProtocolOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ProtocolOutput)
}

func (e Protocol) ToProtocolPtrOutput() ProtocolPtrOutput {
	return e.ToProtocolPtrOutputWithContext(context.Background())
}

func (e Protocol) ToProtocolPtrOutputWithContext(ctx context.Context) ProtocolPtrOutput {
	return Protocol(e).ToProtocolOutputWithContext(ctx).ToProtocolPtrOutputWithContext(ctx)
}

func (e Protocol) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Protocol) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Protocol) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Protocol) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ProtocolOutput struct{ *pulumi.OutputState }

func (ProtocolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Protocol)(nil)).Elem()
}

func (o ProtocolOutput) ToProtocolOutput() ProtocolOutput {
	return o
}

func (o ProtocolOutput) ToProtocolOutputWithContext(ctx context.Context) ProtocolOutput {
	return o
}

func (o ProtocolOutput) ToProtocolPtrOutput() ProtocolPtrOutput {
	return o.ToProtocolPtrOutputWithContext(context.Background())
}

func (o ProtocolOutput) ToProtocolPtrOutputWithContext(ctx context.Context) ProtocolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Protocol) *Protocol {
		return &v
	}).(ProtocolPtrOutput)
}

func (o ProtocolOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ProtocolOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Protocol) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ProtocolOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProtocolOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Protocol) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ProtocolPtrOutput struct{ *pulumi.OutputState }

func (ProtocolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Protocol)(nil)).Elem()
}

func (o ProtocolPtrOutput) ToProtocolPtrOutput() ProtocolPtrOutput {
	return o
}

func (o ProtocolPtrOutput) ToProtocolPtrOutputWithContext(ctx context.Context) ProtocolPtrOutput {
	return o
}

func (o ProtocolPtrOutput) Elem() ProtocolOutput {
	return o.ApplyT(func(v *Protocol) Protocol {
		if v != nil {
			return *v
		}
		var ret Protocol
		return ret
	}).(ProtocolOutput)
}

func (o ProtocolPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProtocolPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Protocol) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ProtocolInput interface {
	pulumi.Input

	ToProtocolOutput() ProtocolOutput
	ToProtocolOutputWithContext(context.Context) ProtocolOutput
}

var protocolPtrType = reflect.TypeOf((**Protocol)(nil)).Elem()

type ProtocolPtrInput interface {
	pulumi.Input

	ToProtocolPtrOutput() ProtocolPtrOutput
	ToProtocolPtrOutputWithContext(context.Context) ProtocolPtrOutput
}

type protocolPtr string

func ProtocolPtr(v string) ProtocolPtrInput {
	return (*protocolPtr)(&v)
}

func (*protocolPtr) ElementType() reflect.Type {
	return protocolPtrType
}

func (in *protocolPtr) ToProtocolPtrOutput() ProtocolPtrOutput {
	return pulumi.ToOutput(in).(ProtocolPtrOutput)
}

func (in *protocolPtr) ToProtocolPtrOutputWithContext(ctx context.Context) ProtocolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ProtocolPtrOutput)
}

type SkuName string

const (
	SkuName_Enterprise_E10        = SkuName("Enterprise_E10")
	SkuName_Enterprise_E20        = SkuName("Enterprise_E20")
	SkuName_Enterprise_E50        = SkuName("Enterprise_E50")
	SkuName_Enterprise_E100       = SkuName("Enterprise_E100")
	SkuName_EnterpriseFlash_F300  = SkuName("EnterpriseFlash_F300")
	SkuName_EnterpriseFlash_F700  = SkuName("EnterpriseFlash_F700")
	SkuName_EnterpriseFlash_F1500 = SkuName("EnterpriseFlash_F1500")
)

func (SkuName) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuName)(nil)).Elem()
}

func (e SkuName) ToSkuNameOutput() SkuNameOutput {
	return pulumi.ToOutput(e).(SkuNameOutput)
}

func (e SkuName) ToSkuNameOutputWithContext(ctx context.Context) SkuNameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SkuNameOutput)
}

func (e SkuName) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return e.ToSkuNamePtrOutputWithContext(context.Background())
}

func (e SkuName) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return SkuName(e).ToSkuNameOutputWithContext(ctx).ToSkuNamePtrOutputWithContext(ctx)
}

func (e SkuName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SkuName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SkuNameOutput struct{ *pulumi.OutputState }

func (SkuNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuName)(nil)).Elem()
}

func (o SkuNameOutput) ToSkuNameOutput() SkuNameOutput {
	return o
}

func (o SkuNameOutput) ToSkuNameOutputWithContext(ctx context.Context) SkuNameOutput {
	return o
}

func (o SkuNameOutput) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return o.ToSkuNamePtrOutputWithContext(context.Background())
}

func (o SkuNameOutput) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuName) *SkuName {
		return &v
	}).(SkuNamePtrOutput)
}

func (o SkuNameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SkuNameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuName) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SkuNameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuNameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuName) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SkuNamePtrOutput struct{ *pulumi.OutputState }

func (SkuNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuName)(nil)).Elem()
}

func (o SkuNamePtrOutput) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return o
}

func (o SkuNamePtrOutput) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return o
}

func (o SkuNamePtrOutput) Elem() SkuNameOutput {
	return o.ApplyT(func(v *SkuName) SkuName {
		if v != nil {
			return *v
		}
		var ret SkuName
		return ret
	}).(SkuNameOutput)
}

func (o SkuNamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SkuName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SkuNameInput interface {
	pulumi.Input

	ToSkuNameOutput() SkuNameOutput
	ToSkuNameOutputWithContext(context.Context) SkuNameOutput
}

var skuNamePtrType = reflect.TypeOf((**SkuName)(nil)).Elem()

type SkuNamePtrInput interface {
	pulumi.Input

	ToSkuNamePtrOutput() SkuNamePtrOutput
	ToSkuNamePtrOutputWithContext(context.Context) SkuNamePtrOutput
}

type skuNamePtr string

func SkuNamePtr(v string) SkuNamePtrInput {
	return (*skuNamePtr)(&v)
}

func (*skuNamePtr) ElementType() reflect.Type {
	return skuNamePtrType
}

func (in *skuNamePtr) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return pulumi.ToOutput(in).(SkuNamePtrOutput)
}

func (in *skuNamePtr) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SkuNamePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusteringPolicyInput)(nil)).Elem(), ClusteringPolicy("EnterpriseCluster"))
	pulumi.RegisterInputType(reflect.TypeOf((*ClusteringPolicyPtrInput)(nil)).Elem(), ClusteringPolicy("EnterpriseCluster"))
	pulumi.RegisterInputType(reflect.TypeOf((*EvictionPolicyInput)(nil)).Elem(), EvictionPolicy("AllKeysLFU"))
	pulumi.RegisterInputType(reflect.TypeOf((*EvictionPolicyPtrInput)(nil)).Elem(), EvictionPolicy("AllKeysLFU"))
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointServiceConnectionStatusInput)(nil)).Elem(), PrivateEndpointServiceConnectionStatus("Pending"))
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointServiceConnectionStatusPtrInput)(nil)).Elem(), PrivateEndpointServiceConnectionStatus("Pending"))
	pulumi.RegisterInputType(reflect.TypeOf((*ProtocolInput)(nil)).Elem(), Protocol("Encrypted"))
	pulumi.RegisterInputType(reflect.TypeOf((*ProtocolPtrInput)(nil)).Elem(), Protocol("Encrypted"))
	pulumi.RegisterInputType(reflect.TypeOf((*SkuNameInput)(nil)).Elem(), SkuName("Enterprise_E10"))
	pulumi.RegisterInputType(reflect.TypeOf((*SkuNamePtrInput)(nil)).Elem(), SkuName("Enterprise_E10"))
	pulumi.RegisterOutputType(ClusteringPolicyOutput{})
	pulumi.RegisterOutputType(ClusteringPolicyPtrOutput{})
	pulumi.RegisterOutputType(EvictionPolicyOutput{})
	pulumi.RegisterOutputType(EvictionPolicyPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointServiceConnectionStatusOutput{})
	pulumi.RegisterOutputType(PrivateEndpointServiceConnectionStatusPtrOutput{})
	pulumi.RegisterOutputType(ProtocolOutput{})
	pulumi.RegisterOutputType(ProtocolPtrOutput{})
	pulumi.RegisterOutputType(SkuNameOutput{})
	pulumi.RegisterOutputType(SkuNamePtrOutput{})
}
