


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Ambr struct {
	Downlink string `pulumi:"downlink"`
	Uplink   string `pulumi:"uplink"`
}





type AmbrInput interface {
	pulumi.Input

	ToAmbrOutput() AmbrOutput
	ToAmbrOutputWithContext(context.Context) AmbrOutput
}

type AmbrArgs struct {
	Downlink pulumi.StringInput `pulumi:"downlink"`
	Uplink   pulumi.StringInput `pulumi:"uplink"`
}

func (AmbrArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Ambr)(nil)).Elem()
}

func (i AmbrArgs) ToAmbrOutput() AmbrOutput {
	return i.ToAmbrOutputWithContext(context.Background())
}

func (i AmbrArgs) ToAmbrOutputWithContext(ctx context.Context) AmbrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AmbrOutput)
}

func (i AmbrArgs) ToAmbrPtrOutput() AmbrPtrOutput {
	return i.ToAmbrPtrOutputWithContext(context.Background())
}

func (i AmbrArgs) ToAmbrPtrOutputWithContext(ctx context.Context) AmbrPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AmbrOutput).ToAmbrPtrOutputWithContext(ctx)
}









type AmbrPtrInput interface {
	pulumi.Input

	ToAmbrPtrOutput() AmbrPtrOutput
	ToAmbrPtrOutputWithContext(context.Context) AmbrPtrOutput
}

type ambrPtrType AmbrArgs

func AmbrPtr(v *AmbrArgs) AmbrPtrInput {
	return (*ambrPtrType)(v)
}

func (*ambrPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Ambr)(nil)).Elem()
}

func (i *ambrPtrType) ToAmbrPtrOutput() AmbrPtrOutput {
	return i.ToAmbrPtrOutputWithContext(context.Background())
}

func (i *ambrPtrType) ToAmbrPtrOutputWithContext(ctx context.Context) AmbrPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AmbrPtrOutput)
}

type AmbrOutput struct{ *pulumi.OutputState }

func (AmbrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Ambr)(nil)).Elem()
}

func (o AmbrOutput) ToAmbrOutput() AmbrOutput {
	return o
}

func (o AmbrOutput) ToAmbrOutputWithContext(ctx context.Context) AmbrOutput {
	return o
}

func (o AmbrOutput) ToAmbrPtrOutput() AmbrPtrOutput {
	return o.ToAmbrPtrOutputWithContext(context.Background())
}

func (o AmbrOutput) ToAmbrPtrOutputWithContext(ctx context.Context) AmbrPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Ambr) *Ambr {
		return &v
	}).(AmbrPtrOutput)
}

func (o AmbrOutput) Downlink() pulumi.StringOutput {
	return o.ApplyT(func(v Ambr) string { return v.Downlink }).(pulumi.StringOutput)
}

func (o AmbrOutput) Uplink() pulumi.StringOutput {
	return o.ApplyT(func(v Ambr) string { return v.Uplink }).(pulumi.StringOutput)
}

type AmbrPtrOutput struct{ *pulumi.OutputState }

func (AmbrPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ambr)(nil)).Elem()
}

func (o AmbrPtrOutput) ToAmbrPtrOutput() AmbrPtrOutput {
	return o
}

func (o AmbrPtrOutput) ToAmbrPtrOutputWithContext(ctx context.Context) AmbrPtrOutput {
	return o
}

func (o AmbrPtrOutput) Elem() AmbrOutput {
	return o.ApplyT(func(v *Ambr) Ambr {
		if v != nil {
			return *v
		}
		var ret Ambr
		return ret
	}).(AmbrOutput)
}

func (o AmbrPtrOutput) Downlink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ambr) *string {
		if v == nil {
			return nil
		}
		return &v.Downlink
	}).(pulumi.StringPtrOutput)
}

func (o AmbrPtrOutput) Uplink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ambr) *string {
		if v == nil {
			return nil
		}
		return &v.Uplink
	}).(pulumi.StringPtrOutput)
}

type AmbrResponse struct {
	Downlink string `pulumi:"downlink"`
	Uplink   string `pulumi:"uplink"`
}

type AmbrResponseOutput struct{ *pulumi.OutputState }

func (AmbrResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AmbrResponse)(nil)).Elem()
}

func (o AmbrResponseOutput) ToAmbrResponseOutput() AmbrResponseOutput {
	return o
}

func (o AmbrResponseOutput) ToAmbrResponseOutputWithContext(ctx context.Context) AmbrResponseOutput {
	return o
}

func (o AmbrResponseOutput) Downlink() pulumi.StringOutput {
	return o.ApplyT(func(v AmbrResponse) string { return v.Downlink }).(pulumi.StringOutput)
}

func (o AmbrResponseOutput) Uplink() pulumi.StringOutput {
	return o.ApplyT(func(v AmbrResponse) string { return v.Uplink }).(pulumi.StringOutput)
}

type AmbrResponsePtrOutput struct{ *pulumi.OutputState }

func (AmbrResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AmbrResponse)(nil)).Elem()
}

func (o AmbrResponsePtrOutput) ToAmbrResponsePtrOutput() AmbrResponsePtrOutput {
	return o
}

func (o AmbrResponsePtrOutput) ToAmbrResponsePtrOutputWithContext(ctx context.Context) AmbrResponsePtrOutput {
	return o
}

func (o AmbrResponsePtrOutput) Elem() AmbrResponseOutput {
	return o.ApplyT(func(v *AmbrResponse) AmbrResponse {
		if v != nil {
			return *v
		}
		var ret AmbrResponse
		return ret
	}).(AmbrResponseOutput)
}

func (o AmbrResponsePtrOutput) Downlink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AmbrResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Downlink
	}).(pulumi.StringPtrOutput)
}

func (o AmbrResponsePtrOutput) Uplink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AmbrResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Uplink
	}).(pulumi.StringPtrOutput)
}

type AttachedDataNetworkResourceId struct {
	Id string `pulumi:"id"`
}





type AttachedDataNetworkResourceIdInput interface {
	pulumi.Input

	ToAttachedDataNetworkResourceIdOutput() AttachedDataNetworkResourceIdOutput
	ToAttachedDataNetworkResourceIdOutputWithContext(context.Context) AttachedDataNetworkResourceIdOutput
}

type AttachedDataNetworkResourceIdArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (AttachedDataNetworkResourceIdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AttachedDataNetworkResourceId)(nil)).Elem()
}

func (i AttachedDataNetworkResourceIdArgs) ToAttachedDataNetworkResourceIdOutput() AttachedDataNetworkResourceIdOutput {
	return i.ToAttachedDataNetworkResourceIdOutputWithContext(context.Background())
}

func (i AttachedDataNetworkResourceIdArgs) ToAttachedDataNetworkResourceIdOutputWithContext(ctx context.Context) AttachedDataNetworkResourceIdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttachedDataNetworkResourceIdOutput)
}

func (i AttachedDataNetworkResourceIdArgs) ToAttachedDataNetworkResourceIdPtrOutput() AttachedDataNetworkResourceIdPtrOutput {
	return i.ToAttachedDataNetworkResourceIdPtrOutputWithContext(context.Background())
}

func (i AttachedDataNetworkResourceIdArgs) ToAttachedDataNetworkResourceIdPtrOutputWithContext(ctx context.Context) AttachedDataNetworkResourceIdPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttachedDataNetworkResourceIdOutput).ToAttachedDataNetworkResourceIdPtrOutputWithContext(ctx)
}









type AttachedDataNetworkResourceIdPtrInput interface {
	pulumi.Input

	ToAttachedDataNetworkResourceIdPtrOutput() AttachedDataNetworkResourceIdPtrOutput
	ToAttachedDataNetworkResourceIdPtrOutputWithContext(context.Context) AttachedDataNetworkResourceIdPtrOutput
}

type attachedDataNetworkResourceIdPtrType AttachedDataNetworkResourceIdArgs

func AttachedDataNetworkResourceIdPtr(v *AttachedDataNetworkResourceIdArgs) AttachedDataNetworkResourceIdPtrInput {
	return (*attachedDataNetworkResourceIdPtrType)(v)
}

func (*attachedDataNetworkResourceIdPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AttachedDataNetworkResourceId)(nil)).Elem()
}

func (i *attachedDataNetworkResourceIdPtrType) ToAttachedDataNetworkResourceIdPtrOutput() AttachedDataNetworkResourceIdPtrOutput {
	return i.ToAttachedDataNetworkResourceIdPtrOutputWithContext(context.Background())
}

func (i *attachedDataNetworkResourceIdPtrType) ToAttachedDataNetworkResourceIdPtrOutputWithContext(ctx context.Context) AttachedDataNetworkResourceIdPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttachedDataNetworkResourceIdPtrOutput)
}

type AttachedDataNetworkResourceIdOutput struct{ *pulumi.OutputState }

func (AttachedDataNetworkResourceIdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AttachedDataNetworkResourceId)(nil)).Elem()
}

func (o AttachedDataNetworkResourceIdOutput) ToAttachedDataNetworkResourceIdOutput() AttachedDataNetworkResourceIdOutput {
	return o
}

func (o AttachedDataNetworkResourceIdOutput) ToAttachedDataNetworkResourceIdOutputWithContext(ctx context.Context) AttachedDataNetworkResourceIdOutput {
	return o
}

func (o AttachedDataNetworkResourceIdOutput) ToAttachedDataNetworkResourceIdPtrOutput() AttachedDataNetworkResourceIdPtrOutput {
	return o.ToAttachedDataNetworkResourceIdPtrOutputWithContext(context.Background())
}

func (o AttachedDataNetworkResourceIdOutput) ToAttachedDataNetworkResourceIdPtrOutputWithContext(ctx context.Context) AttachedDataNetworkResourceIdPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AttachedDataNetworkResourceId) *AttachedDataNetworkResourceId {
		return &v
	}).(AttachedDataNetworkResourceIdPtrOutput)
}

func (o AttachedDataNetworkResourceIdOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v AttachedDataNetworkResourceId) string { return v.Id }).(pulumi.StringOutput)
}

type AttachedDataNetworkResourceIdPtrOutput struct{ *pulumi.OutputState }

func (AttachedDataNetworkResourceIdPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AttachedDataNetworkResourceId)(nil)).Elem()
}

func (o AttachedDataNetworkResourceIdPtrOutput) ToAttachedDataNetworkResourceIdPtrOutput() AttachedDataNetworkResourceIdPtrOutput {
	return o
}

func (o AttachedDataNetworkResourceIdPtrOutput) ToAttachedDataNetworkResourceIdPtrOutputWithContext(ctx context.Context) AttachedDataNetworkResourceIdPtrOutput {
	return o
}

func (o AttachedDataNetworkResourceIdPtrOutput) Elem() AttachedDataNetworkResourceIdOutput {
	return o.ApplyT(func(v *AttachedDataNetworkResourceId) AttachedDataNetworkResourceId {
		if v != nil {
			return *v
		}
		var ret AttachedDataNetworkResourceId
		return ret
	}).(AttachedDataNetworkResourceIdOutput)
}

func (o AttachedDataNetworkResourceIdPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttachedDataNetworkResourceId) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type AttachedDataNetworkResourceIdResponse struct {
	Id string `pulumi:"id"`
}

type AttachedDataNetworkResourceIdResponseOutput struct{ *pulumi.OutputState }

func (AttachedDataNetworkResourceIdResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AttachedDataNetworkResourceIdResponse)(nil)).Elem()
}

func (o AttachedDataNetworkResourceIdResponseOutput) ToAttachedDataNetworkResourceIdResponseOutput() AttachedDataNetworkResourceIdResponseOutput {
	return o
}

func (o AttachedDataNetworkResourceIdResponseOutput) ToAttachedDataNetworkResourceIdResponseOutputWithContext(ctx context.Context) AttachedDataNetworkResourceIdResponseOutput {
	return o
}

func (o AttachedDataNetworkResourceIdResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v AttachedDataNetworkResourceIdResponse) string { return v.Id }).(pulumi.StringOutput)
}

type AttachedDataNetworkResourceIdResponsePtrOutput struct{ *pulumi.OutputState }

func (AttachedDataNetworkResourceIdResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AttachedDataNetworkResourceIdResponse)(nil)).Elem()
}

func (o AttachedDataNetworkResourceIdResponsePtrOutput) ToAttachedDataNetworkResourceIdResponsePtrOutput() AttachedDataNetworkResourceIdResponsePtrOutput {
	return o
}

func (o AttachedDataNetworkResourceIdResponsePtrOutput) ToAttachedDataNetworkResourceIdResponsePtrOutputWithContext(ctx context.Context) AttachedDataNetworkResourceIdResponsePtrOutput {
	return o
}

func (o AttachedDataNetworkResourceIdResponsePtrOutput) Elem() AttachedDataNetworkResourceIdResponseOutput {
	return o.ApplyT(func(v *AttachedDataNetworkResourceIdResponse) AttachedDataNetworkResourceIdResponse {
		if v != nil {
			return *v
		}
		var ret AttachedDataNetworkResourceIdResponse
		return ret
	}).(AttachedDataNetworkResourceIdResponseOutput)
}

func (o AttachedDataNetworkResourceIdResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttachedDataNetworkResourceIdResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type AzureStackEdgeDeviceResourceId struct {
	Id string `pulumi:"id"`
}





type AzureStackEdgeDeviceResourceIdInput interface {
	pulumi.Input

	ToAzureStackEdgeDeviceResourceIdOutput() AzureStackEdgeDeviceResourceIdOutput
	ToAzureStackEdgeDeviceResourceIdOutputWithContext(context.Context) AzureStackEdgeDeviceResourceIdOutput
}

type AzureStackEdgeDeviceResourceIdArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (AzureStackEdgeDeviceResourceIdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureStackEdgeDeviceResourceId)(nil)).Elem()
}

func (i AzureStackEdgeDeviceResourceIdArgs) ToAzureStackEdgeDeviceResourceIdOutput() AzureStackEdgeDeviceResourceIdOutput {
	return i.ToAzureStackEdgeDeviceResourceIdOutputWithContext(context.Background())
}

func (i AzureStackEdgeDeviceResourceIdArgs) ToAzureStackEdgeDeviceResourceIdOutputWithContext(ctx context.Context) AzureStackEdgeDeviceResourceIdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureStackEdgeDeviceResourceIdOutput)
}

func (i AzureStackEdgeDeviceResourceIdArgs) ToAzureStackEdgeDeviceResourceIdPtrOutput() AzureStackEdgeDeviceResourceIdPtrOutput {
	return i.ToAzureStackEdgeDeviceResourceIdPtrOutputWithContext(context.Background())
}

func (i AzureStackEdgeDeviceResourceIdArgs) ToAzureStackEdgeDeviceResourceIdPtrOutputWithContext(ctx context.Context) AzureStackEdgeDeviceResourceIdPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureStackEdgeDeviceResourceIdOutput).ToAzureStackEdgeDeviceResourceIdPtrOutputWithContext(ctx)
}









type AzureStackEdgeDeviceResourceIdPtrInput interface {
	pulumi.Input

	ToAzureStackEdgeDeviceResourceIdPtrOutput() AzureStackEdgeDeviceResourceIdPtrOutput
	ToAzureStackEdgeDeviceResourceIdPtrOutputWithContext(context.Context) AzureStackEdgeDeviceResourceIdPtrOutput
}

type azureStackEdgeDeviceResourceIdPtrType AzureStackEdgeDeviceResourceIdArgs

func AzureStackEdgeDeviceResourceIdPtr(v *AzureStackEdgeDeviceResourceIdArgs) AzureStackEdgeDeviceResourceIdPtrInput {
	return (*azureStackEdgeDeviceResourceIdPtrType)(v)
}

func (*azureStackEdgeDeviceResourceIdPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureStackEdgeDeviceResourceId)(nil)).Elem()
}

func (i *azureStackEdgeDeviceResourceIdPtrType) ToAzureStackEdgeDeviceResourceIdPtrOutput() AzureStackEdgeDeviceResourceIdPtrOutput {
	return i.ToAzureStackEdgeDeviceResourceIdPtrOutputWithContext(context.Background())
}

func (i *azureStackEdgeDeviceResourceIdPtrType) ToAzureStackEdgeDeviceResourceIdPtrOutputWithContext(ctx context.Context) AzureStackEdgeDeviceResourceIdPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureStackEdgeDeviceResourceIdPtrOutput)
}

type AzureStackEdgeDeviceResourceIdOutput struct{ *pulumi.OutputState }

func (AzureStackEdgeDeviceResourceIdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureStackEdgeDeviceResourceId)(nil)).Elem()
}

func (o AzureStackEdgeDeviceResourceIdOutput) ToAzureStackEdgeDeviceResourceIdOutput() AzureStackEdgeDeviceResourceIdOutput {
	return o
}

func (o AzureStackEdgeDeviceResourceIdOutput) ToAzureStackEdgeDeviceResourceIdOutputWithContext(ctx context.Context) AzureStackEdgeDeviceResourceIdOutput {
	return o
}

func (o AzureStackEdgeDeviceResourceIdOutput) ToAzureStackEdgeDeviceResourceIdPtrOutput() AzureStackEdgeDeviceResourceIdPtrOutput {
	return o.ToAzureStackEdgeDeviceResourceIdPtrOutputWithContext(context.Background())
}

func (o AzureStackEdgeDeviceResourceIdOutput) ToAzureStackEdgeDeviceResourceIdPtrOutputWithContext(ctx context.Context) AzureStackEdgeDeviceResourceIdPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureStackEdgeDeviceResourceId) *AzureStackEdgeDeviceResourceId {
		return &v
	}).(AzureStackEdgeDeviceResourceIdPtrOutput)
}

func (o AzureStackEdgeDeviceResourceIdOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v AzureStackEdgeDeviceResourceId) string { return v.Id }).(pulumi.StringOutput)
}

type AzureStackEdgeDeviceResourceIdPtrOutput struct{ *pulumi.OutputState }

func (AzureStackEdgeDeviceResourceIdPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureStackEdgeDeviceResourceId)(nil)).Elem()
}

func (o AzureStackEdgeDeviceResourceIdPtrOutput) ToAzureStackEdgeDeviceResourceIdPtrOutput() AzureStackEdgeDeviceResourceIdPtrOutput {
	return o
}

func (o AzureStackEdgeDeviceResourceIdPtrOutput) ToAzureStackEdgeDeviceResourceIdPtrOutputWithContext(ctx context.Context) AzureStackEdgeDeviceResourceIdPtrOutput {
	return o
}

func (o AzureStackEdgeDeviceResourceIdPtrOutput) Elem() AzureStackEdgeDeviceResourceIdOutput {
	return o.ApplyT(func(v *AzureStackEdgeDeviceResourceId) AzureStackEdgeDeviceResourceId {
		if v != nil {
			return *v
		}
		var ret AzureStackEdgeDeviceResourceId
		return ret
	}).(AzureStackEdgeDeviceResourceIdOutput)
}

func (o AzureStackEdgeDeviceResourceIdPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureStackEdgeDeviceResourceId) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type AzureStackEdgeDeviceResourceIdResponse struct {
	Id string `pulumi:"id"`
}

type AzureStackEdgeDeviceResourceIdResponseOutput struct{ *pulumi.OutputState }

func (AzureStackEdgeDeviceResourceIdResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureStackEdgeDeviceResourceIdResponse)(nil)).Elem()
}

func (o AzureStackEdgeDeviceResourceIdResponseOutput) ToAzureStackEdgeDeviceResourceIdResponseOutput() AzureStackEdgeDeviceResourceIdResponseOutput {
	return o
}

func (o AzureStackEdgeDeviceResourceIdResponseOutput) ToAzureStackEdgeDeviceResourceIdResponseOutputWithContext(ctx context.Context) AzureStackEdgeDeviceResourceIdResponseOutput {
	return o
}

func (o AzureStackEdgeDeviceResourceIdResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v AzureStackEdgeDeviceResourceIdResponse) string { return v.Id }).(pulumi.StringOutput)
}

type AzureStackEdgeDeviceResourceIdResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureStackEdgeDeviceResourceIdResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureStackEdgeDeviceResourceIdResponse)(nil)).Elem()
}

func (o AzureStackEdgeDeviceResourceIdResponsePtrOutput) ToAzureStackEdgeDeviceResourceIdResponsePtrOutput() AzureStackEdgeDeviceResourceIdResponsePtrOutput {
	return o
}

func (o AzureStackEdgeDeviceResourceIdResponsePtrOutput) ToAzureStackEdgeDeviceResourceIdResponsePtrOutputWithContext(ctx context.Context) AzureStackEdgeDeviceResourceIdResponsePtrOutput {
	return o
}

func (o AzureStackEdgeDeviceResourceIdResponsePtrOutput) Elem() AzureStackEdgeDeviceResourceIdResponseOutput {
	return o.ApplyT(func(v *AzureStackEdgeDeviceResourceIdResponse) AzureStackEdgeDeviceResourceIdResponse {
		if v != nil {
			return *v
		}
		var ret AzureStackEdgeDeviceResourceIdResponse
		return ret
	}).(AzureStackEdgeDeviceResourceIdResponseOutput)
}

func (o AzureStackEdgeDeviceResourceIdResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureStackEdgeDeviceResourceIdResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type ConnectedClusterResourceId struct {
	Id string `pulumi:"id"`
}





type ConnectedClusterResourceIdInput interface {
	pulumi.Input

	ToConnectedClusterResourceIdOutput() ConnectedClusterResourceIdOutput
	ToConnectedClusterResourceIdOutputWithContext(context.Context) ConnectedClusterResourceIdOutput
}

type ConnectedClusterResourceIdArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (ConnectedClusterResourceIdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectedClusterResourceId)(nil)).Elem()
}

func (i ConnectedClusterResourceIdArgs) ToConnectedClusterResourceIdOutput() ConnectedClusterResourceIdOutput {
	return i.ToConnectedClusterResourceIdOutputWithContext(context.Background())
}

func (i ConnectedClusterResourceIdArgs) ToConnectedClusterResourceIdOutputWithContext(ctx context.Context) ConnectedClusterResourceIdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectedClusterResourceIdOutput)
}

func (i ConnectedClusterResourceIdArgs) ToConnectedClusterResourceIdPtrOutput() ConnectedClusterResourceIdPtrOutput {
	return i.ToConnectedClusterResourceIdPtrOutputWithContext(context.Background())
}

func (i ConnectedClusterResourceIdArgs) ToConnectedClusterResourceIdPtrOutputWithContext(ctx context.Context) ConnectedClusterResourceIdPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectedClusterResourceIdOutput).ToConnectedClusterResourceIdPtrOutputWithContext(ctx)
}









type ConnectedClusterResourceIdPtrInput interface {
	pulumi.Input

	ToConnectedClusterResourceIdPtrOutput() ConnectedClusterResourceIdPtrOutput
	ToConnectedClusterResourceIdPtrOutputWithContext(context.Context) ConnectedClusterResourceIdPtrOutput
}

type connectedClusterResourceIdPtrType ConnectedClusterResourceIdArgs

func ConnectedClusterResourceIdPtr(v *ConnectedClusterResourceIdArgs) ConnectedClusterResourceIdPtrInput {
	return (*connectedClusterResourceIdPtrType)(v)
}

func (*connectedClusterResourceIdPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectedClusterResourceId)(nil)).Elem()
}

func (i *connectedClusterResourceIdPtrType) ToConnectedClusterResourceIdPtrOutput() ConnectedClusterResourceIdPtrOutput {
	return i.ToConnectedClusterResourceIdPtrOutputWithContext(context.Background())
}

func (i *connectedClusterResourceIdPtrType) ToConnectedClusterResourceIdPtrOutputWithContext(ctx context.Context) ConnectedClusterResourceIdPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectedClusterResourceIdPtrOutput)
}

type ConnectedClusterResourceIdOutput struct{ *pulumi.OutputState }

func (ConnectedClusterResourceIdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectedClusterResourceId)(nil)).Elem()
}

func (o ConnectedClusterResourceIdOutput) ToConnectedClusterResourceIdOutput() ConnectedClusterResourceIdOutput {
	return o
}

func (o ConnectedClusterResourceIdOutput) ToConnectedClusterResourceIdOutputWithContext(ctx context.Context) ConnectedClusterResourceIdOutput {
	return o
}

func (o ConnectedClusterResourceIdOutput) ToConnectedClusterResourceIdPtrOutput() ConnectedClusterResourceIdPtrOutput {
	return o.ToConnectedClusterResourceIdPtrOutputWithContext(context.Background())
}

func (o ConnectedClusterResourceIdOutput) ToConnectedClusterResourceIdPtrOutputWithContext(ctx context.Context) ConnectedClusterResourceIdPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectedClusterResourceId) *ConnectedClusterResourceId {
		return &v
	}).(ConnectedClusterResourceIdPtrOutput)
}

func (o ConnectedClusterResourceIdOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectedClusterResourceId) string { return v.Id }).(pulumi.StringOutput)
}

type ConnectedClusterResourceIdPtrOutput struct{ *pulumi.OutputState }

func (ConnectedClusterResourceIdPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectedClusterResourceId)(nil)).Elem()
}

func (o ConnectedClusterResourceIdPtrOutput) ToConnectedClusterResourceIdPtrOutput() ConnectedClusterResourceIdPtrOutput {
	return o
}

func (o ConnectedClusterResourceIdPtrOutput) ToConnectedClusterResourceIdPtrOutputWithContext(ctx context.Context) ConnectedClusterResourceIdPtrOutput {
	return o
}

func (o ConnectedClusterResourceIdPtrOutput) Elem() ConnectedClusterResourceIdOutput {
	return o.ApplyT(func(v *ConnectedClusterResourceId) ConnectedClusterResourceId {
		if v != nil {
			return *v
		}
		var ret ConnectedClusterResourceId
		return ret
	}).(ConnectedClusterResourceIdOutput)
}

func (o ConnectedClusterResourceIdPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectedClusterResourceId) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type ConnectedClusterResourceIdResponse struct {
	Id string `pulumi:"id"`
}

type ConnectedClusterResourceIdResponseOutput struct{ *pulumi.OutputState }

func (ConnectedClusterResourceIdResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectedClusterResourceIdResponse)(nil)).Elem()
}

func (o ConnectedClusterResourceIdResponseOutput) ToConnectedClusterResourceIdResponseOutput() ConnectedClusterResourceIdResponseOutput {
	return o
}

func (o ConnectedClusterResourceIdResponseOutput) ToConnectedClusterResourceIdResponseOutputWithContext(ctx context.Context) ConnectedClusterResourceIdResponseOutput {
	return o
}

func (o ConnectedClusterResourceIdResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectedClusterResourceIdResponse) string { return v.Id }).(pulumi.StringOutput)
}

type ConnectedClusterResourceIdResponsePtrOutput struct{ *pulumi.OutputState }

func (ConnectedClusterResourceIdResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectedClusterResourceIdResponse)(nil)).Elem()
}

func (o ConnectedClusterResourceIdResponsePtrOutput) ToConnectedClusterResourceIdResponsePtrOutput() ConnectedClusterResourceIdResponsePtrOutput {
	return o
}

func (o ConnectedClusterResourceIdResponsePtrOutput) ToConnectedClusterResourceIdResponsePtrOutputWithContext(ctx context.Context) ConnectedClusterResourceIdResponsePtrOutput {
	return o
}

func (o ConnectedClusterResourceIdResponsePtrOutput) Elem() ConnectedClusterResourceIdResponseOutput {
	return o.ApplyT(func(v *ConnectedClusterResourceIdResponse) ConnectedClusterResourceIdResponse {
		if v != nil {
			return *v
		}
		var ret ConnectedClusterResourceIdResponse
		return ret
	}).(ConnectedClusterResourceIdResponseOutput)
}

func (o ConnectedClusterResourceIdResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectedClusterResourceIdResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type CustomLocationResourceId struct {
	Id string `pulumi:"id"`
}





type CustomLocationResourceIdInput interface {
	pulumi.Input

	ToCustomLocationResourceIdOutput() CustomLocationResourceIdOutput
	ToCustomLocationResourceIdOutputWithContext(context.Context) CustomLocationResourceIdOutput
}

type CustomLocationResourceIdArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (CustomLocationResourceIdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomLocationResourceId)(nil)).Elem()
}

func (i CustomLocationResourceIdArgs) ToCustomLocationResourceIdOutput() CustomLocationResourceIdOutput {
	return i.ToCustomLocationResourceIdOutputWithContext(context.Background())
}

func (i CustomLocationResourceIdArgs) ToCustomLocationResourceIdOutputWithContext(ctx context.Context) CustomLocationResourceIdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomLocationResourceIdOutput)
}

func (i CustomLocationResourceIdArgs) ToCustomLocationResourceIdPtrOutput() CustomLocationResourceIdPtrOutput {
	return i.ToCustomLocationResourceIdPtrOutputWithContext(context.Background())
}

func (i CustomLocationResourceIdArgs) ToCustomLocationResourceIdPtrOutputWithContext(ctx context.Context) CustomLocationResourceIdPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomLocationResourceIdOutput).ToCustomLocationResourceIdPtrOutputWithContext(ctx)
}









type CustomLocationResourceIdPtrInput interface {
	pulumi.Input

	ToCustomLocationResourceIdPtrOutput() CustomLocationResourceIdPtrOutput
	ToCustomLocationResourceIdPtrOutputWithContext(context.Context) CustomLocationResourceIdPtrOutput
}

type customLocationResourceIdPtrType CustomLocationResourceIdArgs

func CustomLocationResourceIdPtr(v *CustomLocationResourceIdArgs) CustomLocationResourceIdPtrInput {
	return (*customLocationResourceIdPtrType)(v)
}

func (*customLocationResourceIdPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomLocationResourceId)(nil)).Elem()
}

func (i *customLocationResourceIdPtrType) ToCustomLocationResourceIdPtrOutput() CustomLocationResourceIdPtrOutput {
	return i.ToCustomLocationResourceIdPtrOutputWithContext(context.Background())
}

func (i *customLocationResourceIdPtrType) ToCustomLocationResourceIdPtrOutputWithContext(ctx context.Context) CustomLocationResourceIdPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomLocationResourceIdPtrOutput)
}

type CustomLocationResourceIdOutput struct{ *pulumi.OutputState }

func (CustomLocationResourceIdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomLocationResourceId)(nil)).Elem()
}

func (o CustomLocationResourceIdOutput) ToCustomLocationResourceIdOutput() CustomLocationResourceIdOutput {
	return o
}

func (o CustomLocationResourceIdOutput) ToCustomLocationResourceIdOutputWithContext(ctx context.Context) CustomLocationResourceIdOutput {
	return o
}

func (o CustomLocationResourceIdOutput) ToCustomLocationResourceIdPtrOutput() CustomLocationResourceIdPtrOutput {
	return o.ToCustomLocationResourceIdPtrOutputWithContext(context.Background())
}

func (o CustomLocationResourceIdOutput) ToCustomLocationResourceIdPtrOutputWithContext(ctx context.Context) CustomLocationResourceIdPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomLocationResourceId) *CustomLocationResourceId {
		return &v
	}).(CustomLocationResourceIdPtrOutput)
}

func (o CustomLocationResourceIdOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v CustomLocationResourceId) string { return v.Id }).(pulumi.StringOutput)
}

type CustomLocationResourceIdPtrOutput struct{ *pulumi.OutputState }

func (CustomLocationResourceIdPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomLocationResourceId)(nil)).Elem()
}

func (o CustomLocationResourceIdPtrOutput) ToCustomLocationResourceIdPtrOutput() CustomLocationResourceIdPtrOutput {
	return o
}

func (o CustomLocationResourceIdPtrOutput) ToCustomLocationResourceIdPtrOutputWithContext(ctx context.Context) CustomLocationResourceIdPtrOutput {
	return o
}

func (o CustomLocationResourceIdPtrOutput) Elem() CustomLocationResourceIdOutput {
	return o.ApplyT(func(v *CustomLocationResourceId) CustomLocationResourceId {
		if v != nil {
			return *v
		}
		var ret CustomLocationResourceId
		return ret
	}).(CustomLocationResourceIdOutput)
}

func (o CustomLocationResourceIdPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomLocationResourceId) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type CustomLocationResourceIdResponse struct {
	Id string `pulumi:"id"`
}

type CustomLocationResourceIdResponseOutput struct{ *pulumi.OutputState }

func (CustomLocationResourceIdResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomLocationResourceIdResponse)(nil)).Elem()
}

func (o CustomLocationResourceIdResponseOutput) ToCustomLocationResourceIdResponseOutput() CustomLocationResourceIdResponseOutput {
	return o
}

func (o CustomLocationResourceIdResponseOutput) ToCustomLocationResourceIdResponseOutputWithContext(ctx context.Context) CustomLocationResourceIdResponseOutput {
	return o
}

func (o CustomLocationResourceIdResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v CustomLocationResourceIdResponse) string { return v.Id }).(pulumi.StringOutput)
}

type CustomLocationResourceIdResponsePtrOutput struct{ *pulumi.OutputState }

func (CustomLocationResourceIdResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomLocationResourceIdResponse)(nil)).Elem()
}

func (o CustomLocationResourceIdResponsePtrOutput) ToCustomLocationResourceIdResponsePtrOutput() CustomLocationResourceIdResponsePtrOutput {
	return o
}

func (o CustomLocationResourceIdResponsePtrOutput) ToCustomLocationResourceIdResponsePtrOutputWithContext(ctx context.Context) CustomLocationResourceIdResponsePtrOutput {
	return o
}

func (o CustomLocationResourceIdResponsePtrOutput) Elem() CustomLocationResourceIdResponseOutput {
	return o.ApplyT(func(v *CustomLocationResourceIdResponse) CustomLocationResourceIdResponse {
		if v != nil {
			return *v
		}
		var ret CustomLocationResourceIdResponse
		return ret
	}).(CustomLocationResourceIdResponseOutput)
}

func (o CustomLocationResourceIdResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomLocationResourceIdResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type DataNetworkConfiguration struct {
	AdditionalAllowedSessionTypes       []string              `pulumi:"additionalAllowedSessionTypes"`
	AllocationAndRetentionPriorityLevel *int                  `pulumi:"allocationAndRetentionPriorityLevel"`
	AllowedServices                     []ServiceResourceId   `pulumi:"allowedServices"`
	DataNetwork                         DataNetworkResourceId `pulumi:"dataNetwork"`
	DefaultSessionType                  *string               `pulumi:"defaultSessionType"`
	FiveQi                              *int                  `pulumi:"fiveQi"`
	PreemptionCapability                *string               `pulumi:"preemptionCapability"`
	PreemptionVulnerability             *string               `pulumi:"preemptionVulnerability"`
	SessionAmbr                         Ambr                  `pulumi:"sessionAmbr"`
}


func (val *DataNetworkConfiguration) Defaults() *DataNetworkConfiguration {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AllocationAndRetentionPriorityLevel) {
		allocationAndRetentionPriorityLevel_ := 9
		tmp.AllocationAndRetentionPriorityLevel = &allocationAndRetentionPriorityLevel_
	}
	if isZero(tmp.DefaultSessionType) {
		defaultSessionType_ := "IPv4"
		tmp.DefaultSessionType = &defaultSessionType_
	}
	if isZero(tmp.FiveQi) {
		fiveQi_ := 9
		tmp.FiveQi = &fiveQi_
	}
	if isZero(tmp.PreemptionCapability) {
		preemptionCapability_ := "NotPreempt"
		tmp.PreemptionCapability = &preemptionCapability_
	}
	if isZero(tmp.PreemptionVulnerability) {
		preemptionVulnerability_ := "Preemptable"
		tmp.PreemptionVulnerability = &preemptionVulnerability_
	}
	return &tmp
}





type DataNetworkConfigurationInput interface {
	pulumi.Input

	ToDataNetworkConfigurationOutput() DataNetworkConfigurationOutput
	ToDataNetworkConfigurationOutputWithContext(context.Context) DataNetworkConfigurationOutput
}

type DataNetworkConfigurationArgs struct {
	AdditionalAllowedSessionTypes       pulumi.StringArrayInput     `pulumi:"additionalAllowedSessionTypes"`
	AllocationAndRetentionPriorityLevel pulumi.IntPtrInput          `pulumi:"allocationAndRetentionPriorityLevel"`
	AllowedServices                     ServiceResourceIdArrayInput `pulumi:"allowedServices"`
	DataNetwork                         DataNetworkResourceIdInput  `pulumi:"dataNetwork"`
	DefaultSessionType                  pulumi.StringPtrInput       `pulumi:"defaultSessionType"`
	FiveQi                              pulumi.IntPtrInput          `pulumi:"fiveQi"`
	PreemptionCapability                pulumi.StringPtrInput       `pulumi:"preemptionCapability"`
	PreemptionVulnerability             pulumi.StringPtrInput       `pulumi:"preemptionVulnerability"`
	SessionAmbr                         AmbrInput                   `pulumi:"sessionAmbr"`
}


func (val *DataNetworkConfigurationArgs) Defaults() *DataNetworkConfigurationArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AllocationAndRetentionPriorityLevel) {
		tmp.AllocationAndRetentionPriorityLevel = pulumi.IntPtr(9)
	}
	if isZero(tmp.DefaultSessionType) {
		tmp.DefaultSessionType = pulumi.StringPtr("IPv4")
	}
	if isZero(tmp.FiveQi) {
		tmp.FiveQi = pulumi.IntPtr(9)
	}
	if isZero(tmp.PreemptionCapability) {
		tmp.PreemptionCapability = pulumi.StringPtr("NotPreempt")
	}
	if isZero(tmp.PreemptionVulnerability) {
		tmp.PreemptionVulnerability = pulumi.StringPtr("Preemptable")
	}
	return &tmp
}
func (DataNetworkConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataNetworkConfiguration)(nil)).Elem()
}

func (i DataNetworkConfigurationArgs) ToDataNetworkConfigurationOutput() DataNetworkConfigurationOutput {
	return i.ToDataNetworkConfigurationOutputWithContext(context.Background())
}

func (i DataNetworkConfigurationArgs) ToDataNetworkConfigurationOutputWithContext(ctx context.Context) DataNetworkConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataNetworkConfigurationOutput)
}





type DataNetworkConfigurationArrayInput interface {
	pulumi.Input

	ToDataNetworkConfigurationArrayOutput() DataNetworkConfigurationArrayOutput
	ToDataNetworkConfigurationArrayOutputWithContext(context.Context) DataNetworkConfigurationArrayOutput
}

type DataNetworkConfigurationArray []DataNetworkConfigurationInput

func (DataNetworkConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataNetworkConfiguration)(nil)).Elem()
}

func (i DataNetworkConfigurationArray) ToDataNetworkConfigurationArrayOutput() DataNetworkConfigurationArrayOutput {
	return i.ToDataNetworkConfigurationArrayOutputWithContext(context.Background())
}

func (i DataNetworkConfigurationArray) ToDataNetworkConfigurationArrayOutputWithContext(ctx context.Context) DataNetworkConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataNetworkConfigurationArrayOutput)
}

type DataNetworkConfigurationOutput struct{ *pulumi.OutputState }

func (DataNetworkConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataNetworkConfiguration)(nil)).Elem()
}

func (o DataNetworkConfigurationOutput) ToDataNetworkConfigurationOutput() DataNetworkConfigurationOutput {
	return o
}

func (o DataNetworkConfigurationOutput) ToDataNetworkConfigurationOutputWithContext(ctx context.Context) DataNetworkConfigurationOutput {
	return o
}

func (o DataNetworkConfigurationOutput) AdditionalAllowedSessionTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DataNetworkConfiguration) []string { return v.AdditionalAllowedSessionTypes }).(pulumi.StringArrayOutput)
}

func (o DataNetworkConfigurationOutput) AllocationAndRetentionPriorityLevel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DataNetworkConfiguration) *int { return v.AllocationAndRetentionPriorityLevel }).(pulumi.IntPtrOutput)
}

func (o DataNetworkConfigurationOutput) AllowedServices() ServiceResourceIdArrayOutput {
	return o.ApplyT(func(v DataNetworkConfiguration) []ServiceResourceId { return v.AllowedServices }).(ServiceResourceIdArrayOutput)
}

func (o DataNetworkConfigurationOutput) DataNetwork() DataNetworkResourceIdOutput {
	return o.ApplyT(func(v DataNetworkConfiguration) DataNetworkResourceId { return v.DataNetwork }).(DataNetworkResourceIdOutput)
}

func (o DataNetworkConfigurationOutput) DefaultSessionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataNetworkConfiguration) *string { return v.DefaultSessionType }).(pulumi.StringPtrOutput)
}

func (o DataNetworkConfigurationOutput) FiveQi() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DataNetworkConfiguration) *int { return v.FiveQi }).(pulumi.IntPtrOutput)
}

func (o DataNetworkConfigurationOutput) PreemptionCapability() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataNetworkConfiguration) *string { return v.PreemptionCapability }).(pulumi.StringPtrOutput)
}

func (o DataNetworkConfigurationOutput) PreemptionVulnerability() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataNetworkConfiguration) *string { return v.PreemptionVulnerability }).(pulumi.StringPtrOutput)
}

func (o DataNetworkConfigurationOutput) SessionAmbr() AmbrOutput {
	return o.ApplyT(func(v DataNetworkConfiguration) Ambr { return v.SessionAmbr }).(AmbrOutput)
}

type DataNetworkConfigurationArrayOutput struct{ *pulumi.OutputState }

func (DataNetworkConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataNetworkConfiguration)(nil)).Elem()
}

func (o DataNetworkConfigurationArrayOutput) ToDataNetworkConfigurationArrayOutput() DataNetworkConfigurationArrayOutput {
	return o
}

func (o DataNetworkConfigurationArrayOutput) ToDataNetworkConfigurationArrayOutputWithContext(ctx context.Context) DataNetworkConfigurationArrayOutput {
	return o
}

func (o DataNetworkConfigurationArrayOutput) Index(i pulumi.IntInput) DataNetworkConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataNetworkConfiguration {
		return vs[0].([]DataNetworkConfiguration)[vs[1].(int)]
	}).(DataNetworkConfigurationOutput)
}

type DataNetworkConfigurationResponse struct {
	AdditionalAllowedSessionTypes       []string                      `pulumi:"additionalAllowedSessionTypes"`
	AllocationAndRetentionPriorityLevel *int                          `pulumi:"allocationAndRetentionPriorityLevel"`
	AllowedServices                     []ServiceResourceIdResponse   `pulumi:"allowedServices"`
	DataNetwork                         DataNetworkResourceIdResponse `pulumi:"dataNetwork"`
	DefaultSessionType                  *string                       `pulumi:"defaultSessionType"`
	FiveQi                              *int                          `pulumi:"fiveQi"`
	PreemptionCapability                *string                       `pulumi:"preemptionCapability"`
	PreemptionVulnerability             *string                       `pulumi:"preemptionVulnerability"`
	SessionAmbr                         AmbrResponse                  `pulumi:"sessionAmbr"`
}


func (val *DataNetworkConfigurationResponse) Defaults() *DataNetworkConfigurationResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AllocationAndRetentionPriorityLevel) {
		allocationAndRetentionPriorityLevel_ := 9
		tmp.AllocationAndRetentionPriorityLevel = &allocationAndRetentionPriorityLevel_
	}
	if isZero(tmp.DefaultSessionType) {
		defaultSessionType_ := "IPv4"
		tmp.DefaultSessionType = &defaultSessionType_
	}
	if isZero(tmp.FiveQi) {
		fiveQi_ := 9
		tmp.FiveQi = &fiveQi_
	}
	if isZero(tmp.PreemptionCapability) {
		preemptionCapability_ := "NotPreempt"
		tmp.PreemptionCapability = &preemptionCapability_
	}
	if isZero(tmp.PreemptionVulnerability) {
		preemptionVulnerability_ := "Preemptable"
		tmp.PreemptionVulnerability = &preemptionVulnerability_
	}
	return &tmp
}

type DataNetworkConfigurationResponseOutput struct{ *pulumi.OutputState }

func (DataNetworkConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataNetworkConfigurationResponse)(nil)).Elem()
}

func (o DataNetworkConfigurationResponseOutput) ToDataNetworkConfigurationResponseOutput() DataNetworkConfigurationResponseOutput {
	return o
}

func (o DataNetworkConfigurationResponseOutput) ToDataNetworkConfigurationResponseOutputWithContext(ctx context.Context) DataNetworkConfigurationResponseOutput {
	return o
}

func (o DataNetworkConfigurationResponseOutput) AdditionalAllowedSessionTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DataNetworkConfigurationResponse) []string { return v.AdditionalAllowedSessionTypes }).(pulumi.StringArrayOutput)
}

func (o DataNetworkConfigurationResponseOutput) AllocationAndRetentionPriorityLevel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DataNetworkConfigurationResponse) *int { return v.AllocationAndRetentionPriorityLevel }).(pulumi.IntPtrOutput)
}

func (o DataNetworkConfigurationResponseOutput) AllowedServices() ServiceResourceIdResponseArrayOutput {
	return o.ApplyT(func(v DataNetworkConfigurationResponse) []ServiceResourceIdResponse { return v.AllowedServices }).(ServiceResourceIdResponseArrayOutput)
}

func (o DataNetworkConfigurationResponseOutput) DataNetwork() DataNetworkResourceIdResponseOutput {
	return o.ApplyT(func(v DataNetworkConfigurationResponse) DataNetworkResourceIdResponse { return v.DataNetwork }).(DataNetworkResourceIdResponseOutput)
}

func (o DataNetworkConfigurationResponseOutput) DefaultSessionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataNetworkConfigurationResponse) *string { return v.DefaultSessionType }).(pulumi.StringPtrOutput)
}

func (o DataNetworkConfigurationResponseOutput) FiveQi() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DataNetworkConfigurationResponse) *int { return v.FiveQi }).(pulumi.IntPtrOutput)
}

func (o DataNetworkConfigurationResponseOutput) PreemptionCapability() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataNetworkConfigurationResponse) *string { return v.PreemptionCapability }).(pulumi.StringPtrOutput)
}

func (o DataNetworkConfigurationResponseOutput) PreemptionVulnerability() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataNetworkConfigurationResponse) *string { return v.PreemptionVulnerability }).(pulumi.StringPtrOutput)
}

func (o DataNetworkConfigurationResponseOutput) SessionAmbr() AmbrResponseOutput {
	return o.ApplyT(func(v DataNetworkConfigurationResponse) AmbrResponse { return v.SessionAmbr }).(AmbrResponseOutput)
}

type DataNetworkConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (DataNetworkConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataNetworkConfigurationResponse)(nil)).Elem()
}

func (o DataNetworkConfigurationResponseArrayOutput) ToDataNetworkConfigurationResponseArrayOutput() DataNetworkConfigurationResponseArrayOutput {
	return o
}

func (o DataNetworkConfigurationResponseArrayOutput) ToDataNetworkConfigurationResponseArrayOutputWithContext(ctx context.Context) DataNetworkConfigurationResponseArrayOutput {
	return o
}

func (o DataNetworkConfigurationResponseArrayOutput) Index(i pulumi.IntInput) DataNetworkConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataNetworkConfigurationResponse {
		return vs[0].([]DataNetworkConfigurationResponse)[vs[1].(int)]
	}).(DataNetworkConfigurationResponseOutput)
}

type DataNetworkResourceId struct {
	Id string `pulumi:"id"`
}





type DataNetworkResourceIdInput interface {
	pulumi.Input

	ToDataNetworkResourceIdOutput() DataNetworkResourceIdOutput
	ToDataNetworkResourceIdOutputWithContext(context.Context) DataNetworkResourceIdOutput
}

type DataNetworkResourceIdArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (DataNetworkResourceIdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataNetworkResourceId)(nil)).Elem()
}

func (i DataNetworkResourceIdArgs) ToDataNetworkResourceIdOutput() DataNetworkResourceIdOutput {
	return i.ToDataNetworkResourceIdOutputWithContext(context.Background())
}

func (i DataNetworkResourceIdArgs) ToDataNetworkResourceIdOutputWithContext(ctx context.Context) DataNetworkResourceIdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataNetworkResourceIdOutput)
}

type DataNetworkResourceIdOutput struct{ *pulumi.OutputState }

func (DataNetworkResourceIdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataNetworkResourceId)(nil)).Elem()
}

func (o DataNetworkResourceIdOutput) ToDataNetworkResourceIdOutput() DataNetworkResourceIdOutput {
	return o
}

func (o DataNetworkResourceIdOutput) ToDataNetworkResourceIdOutputWithContext(ctx context.Context) DataNetworkResourceIdOutput {
	return o
}

func (o DataNetworkResourceIdOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v DataNetworkResourceId) string { return v.Id }).(pulumi.StringOutput)
}

type DataNetworkResourceIdResponse struct {
	Id string `pulumi:"id"`
}

type DataNetworkResourceIdResponseOutput struct{ *pulumi.OutputState }

func (DataNetworkResourceIdResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataNetworkResourceIdResponse)(nil)).Elem()
}

func (o DataNetworkResourceIdResponseOutput) ToDataNetworkResourceIdResponseOutput() DataNetworkResourceIdResponseOutput {
	return o
}

func (o DataNetworkResourceIdResponseOutput) ToDataNetworkResourceIdResponseOutputWithContext(ctx context.Context) DataNetworkResourceIdResponseOutput {
	return o
}

func (o DataNetworkResourceIdResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v DataNetworkResourceIdResponse) string { return v.Id }).(pulumi.StringOutput)
}

type InterfaceProperties struct {
	Ipv4Address *string `pulumi:"ipv4Address"`
	Ipv4Gateway *string `pulumi:"ipv4Gateway"`
	Ipv4Subnet  *string `pulumi:"ipv4Subnet"`
	Name        *string `pulumi:"name"`
}





type InterfacePropertiesInput interface {
	pulumi.Input

	ToInterfacePropertiesOutput() InterfacePropertiesOutput
	ToInterfacePropertiesOutputWithContext(context.Context) InterfacePropertiesOutput
}

type InterfacePropertiesArgs struct {
	Ipv4Address pulumi.StringPtrInput `pulumi:"ipv4Address"`
	Ipv4Gateway pulumi.StringPtrInput `pulumi:"ipv4Gateway"`
	Ipv4Subnet  pulumi.StringPtrInput `pulumi:"ipv4Subnet"`
	Name        pulumi.StringPtrInput `pulumi:"name"`
}

func (InterfacePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InterfaceProperties)(nil)).Elem()
}

func (i InterfacePropertiesArgs) ToInterfacePropertiesOutput() InterfacePropertiesOutput {
	return i.ToInterfacePropertiesOutputWithContext(context.Background())
}

func (i InterfacePropertiesArgs) ToInterfacePropertiesOutputWithContext(ctx context.Context) InterfacePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InterfacePropertiesOutput)
}

type InterfacePropertiesOutput struct{ *pulumi.OutputState }

func (InterfacePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InterfaceProperties)(nil)).Elem()
}

func (o InterfacePropertiesOutput) ToInterfacePropertiesOutput() InterfacePropertiesOutput {
	return o
}

func (o InterfacePropertiesOutput) ToInterfacePropertiesOutputWithContext(ctx context.Context) InterfacePropertiesOutput {
	return o
}

func (o InterfacePropertiesOutput) Ipv4Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InterfaceProperties) *string { return v.Ipv4Address }).(pulumi.StringPtrOutput)
}

func (o InterfacePropertiesOutput) Ipv4Gateway() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InterfaceProperties) *string { return v.Ipv4Gateway }).(pulumi.StringPtrOutput)
}

func (o InterfacePropertiesOutput) Ipv4Subnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InterfaceProperties) *string { return v.Ipv4Subnet }).(pulumi.StringPtrOutput)
}

func (o InterfacePropertiesOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InterfaceProperties) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type InterfacePropertiesResponse struct {
	Ipv4Address *string `pulumi:"ipv4Address"`
	Ipv4Gateway *string `pulumi:"ipv4Gateway"`
	Ipv4Subnet  *string `pulumi:"ipv4Subnet"`
	Name        *string `pulumi:"name"`
}

type InterfacePropertiesResponseOutput struct{ *pulumi.OutputState }

func (InterfacePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InterfacePropertiesResponse)(nil)).Elem()
}

func (o InterfacePropertiesResponseOutput) ToInterfacePropertiesResponseOutput() InterfacePropertiesResponseOutput {
	return o
}

func (o InterfacePropertiesResponseOutput) ToInterfacePropertiesResponseOutputWithContext(ctx context.Context) InterfacePropertiesResponseOutput {
	return o
}

func (o InterfacePropertiesResponseOutput) Ipv4Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InterfacePropertiesResponse) *string { return v.Ipv4Address }).(pulumi.StringPtrOutput)
}

func (o InterfacePropertiesResponseOutput) Ipv4Gateway() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InterfacePropertiesResponse) *string { return v.Ipv4Gateway }).(pulumi.StringPtrOutput)
}

func (o InterfacePropertiesResponseOutput) Ipv4Subnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InterfacePropertiesResponse) *string { return v.Ipv4Subnet }).(pulumi.StringPtrOutput)
}

func (o InterfacePropertiesResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InterfacePropertiesResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type KeyVaultCertificate struct {
	CertificateUrl *string `pulumi:"certificateUrl"`
}





type KeyVaultCertificateInput interface {
	pulumi.Input

	ToKeyVaultCertificateOutput() KeyVaultCertificateOutput
	ToKeyVaultCertificateOutputWithContext(context.Context) KeyVaultCertificateOutput
}

type KeyVaultCertificateArgs struct {
	CertificateUrl pulumi.StringPtrInput `pulumi:"certificateUrl"`
}

func (KeyVaultCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultCertificate)(nil)).Elem()
}

func (i KeyVaultCertificateArgs) ToKeyVaultCertificateOutput() KeyVaultCertificateOutput {
	return i.ToKeyVaultCertificateOutputWithContext(context.Background())
}

func (i KeyVaultCertificateArgs) ToKeyVaultCertificateOutputWithContext(ctx context.Context) KeyVaultCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultCertificateOutput)
}

func (i KeyVaultCertificateArgs) ToKeyVaultCertificatePtrOutput() KeyVaultCertificatePtrOutput {
	return i.ToKeyVaultCertificatePtrOutputWithContext(context.Background())
}

func (i KeyVaultCertificateArgs) ToKeyVaultCertificatePtrOutputWithContext(ctx context.Context) KeyVaultCertificatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultCertificateOutput).ToKeyVaultCertificatePtrOutputWithContext(ctx)
}









type KeyVaultCertificatePtrInput interface {
	pulumi.Input

	ToKeyVaultCertificatePtrOutput() KeyVaultCertificatePtrOutput
	ToKeyVaultCertificatePtrOutputWithContext(context.Context) KeyVaultCertificatePtrOutput
}

type keyVaultCertificatePtrType KeyVaultCertificateArgs

func KeyVaultCertificatePtr(v *KeyVaultCertificateArgs) KeyVaultCertificatePtrInput {
	return (*keyVaultCertificatePtrType)(v)
}

func (*keyVaultCertificatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultCertificate)(nil)).Elem()
}

func (i *keyVaultCertificatePtrType) ToKeyVaultCertificatePtrOutput() KeyVaultCertificatePtrOutput {
	return i.ToKeyVaultCertificatePtrOutputWithContext(context.Background())
}

func (i *keyVaultCertificatePtrType) ToKeyVaultCertificatePtrOutputWithContext(ctx context.Context) KeyVaultCertificatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultCertificatePtrOutput)
}

type KeyVaultCertificateOutput struct{ *pulumi.OutputState }

func (KeyVaultCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultCertificate)(nil)).Elem()
}

func (o KeyVaultCertificateOutput) ToKeyVaultCertificateOutput() KeyVaultCertificateOutput {
	return o
}

func (o KeyVaultCertificateOutput) ToKeyVaultCertificateOutputWithContext(ctx context.Context) KeyVaultCertificateOutput {
	return o
}

func (o KeyVaultCertificateOutput) ToKeyVaultCertificatePtrOutput() KeyVaultCertificatePtrOutput {
	return o.ToKeyVaultCertificatePtrOutputWithContext(context.Background())
}

func (o KeyVaultCertificateOutput) ToKeyVaultCertificatePtrOutputWithContext(ctx context.Context) KeyVaultCertificatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultCertificate) *KeyVaultCertificate {
		return &v
	}).(KeyVaultCertificatePtrOutput)
}

func (o KeyVaultCertificateOutput) CertificateUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultCertificate) *string { return v.CertificateUrl }).(pulumi.StringPtrOutput)
}

type KeyVaultCertificatePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultCertificatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultCertificate)(nil)).Elem()
}

func (o KeyVaultCertificatePtrOutput) ToKeyVaultCertificatePtrOutput() KeyVaultCertificatePtrOutput {
	return o
}

func (o KeyVaultCertificatePtrOutput) ToKeyVaultCertificatePtrOutputWithContext(ctx context.Context) KeyVaultCertificatePtrOutput {
	return o
}

func (o KeyVaultCertificatePtrOutput) Elem() KeyVaultCertificateOutput {
	return o.ApplyT(func(v *KeyVaultCertificate) KeyVaultCertificate {
		if v != nil {
			return *v
		}
		var ret KeyVaultCertificate
		return ret
	}).(KeyVaultCertificateOutput)
}

func (o KeyVaultCertificatePtrOutput) CertificateUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultCertificate) *string {
		if v == nil {
			return nil
		}
		return v.CertificateUrl
	}).(pulumi.StringPtrOutput)
}

type KeyVaultCertificateResponse struct {
	CertificateUrl *string `pulumi:"certificateUrl"`
}

type KeyVaultCertificateResponseOutput struct{ *pulumi.OutputState }

func (KeyVaultCertificateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultCertificateResponse)(nil)).Elem()
}

func (o KeyVaultCertificateResponseOutput) ToKeyVaultCertificateResponseOutput() KeyVaultCertificateResponseOutput {
	return o
}

func (o KeyVaultCertificateResponseOutput) ToKeyVaultCertificateResponseOutputWithContext(ctx context.Context) KeyVaultCertificateResponseOutput {
	return o
}

func (o KeyVaultCertificateResponseOutput) CertificateUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultCertificateResponse) *string { return v.CertificateUrl }).(pulumi.StringPtrOutput)
}

type KeyVaultCertificateResponsePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultCertificateResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultCertificateResponse)(nil)).Elem()
}

func (o KeyVaultCertificateResponsePtrOutput) ToKeyVaultCertificateResponsePtrOutput() KeyVaultCertificateResponsePtrOutput {
	return o
}

func (o KeyVaultCertificateResponsePtrOutput) ToKeyVaultCertificateResponsePtrOutputWithContext(ctx context.Context) KeyVaultCertificateResponsePtrOutput {
	return o
}

func (o KeyVaultCertificateResponsePtrOutput) Elem() KeyVaultCertificateResponseOutput {
	return o.ApplyT(func(v *KeyVaultCertificateResponse) KeyVaultCertificateResponse {
		if v != nil {
			return *v
		}
		var ret KeyVaultCertificateResponse
		return ret
	}).(KeyVaultCertificateResponseOutput)
}

func (o KeyVaultCertificateResponsePtrOutput) CertificateUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultCertificateResponse) *string {
		if v == nil {
			return nil
		}
		return v.CertificateUrl
	}).(pulumi.StringPtrOutput)
}

type KeyVaultKey struct {
	KeyUrl *string `pulumi:"keyUrl"`
}





type KeyVaultKeyInput interface {
	pulumi.Input

	ToKeyVaultKeyOutput() KeyVaultKeyOutput
	ToKeyVaultKeyOutputWithContext(context.Context) KeyVaultKeyOutput
}

type KeyVaultKeyArgs struct {
	KeyUrl pulumi.StringPtrInput `pulumi:"keyUrl"`
}

func (KeyVaultKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultKey)(nil)).Elem()
}

func (i KeyVaultKeyArgs) ToKeyVaultKeyOutput() KeyVaultKeyOutput {
	return i.ToKeyVaultKeyOutputWithContext(context.Background())
}

func (i KeyVaultKeyArgs) ToKeyVaultKeyOutputWithContext(ctx context.Context) KeyVaultKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultKeyOutput)
}

func (i KeyVaultKeyArgs) ToKeyVaultKeyPtrOutput() KeyVaultKeyPtrOutput {
	return i.ToKeyVaultKeyPtrOutputWithContext(context.Background())
}

func (i KeyVaultKeyArgs) ToKeyVaultKeyPtrOutputWithContext(ctx context.Context) KeyVaultKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultKeyOutput).ToKeyVaultKeyPtrOutputWithContext(ctx)
}









type KeyVaultKeyPtrInput interface {
	pulumi.Input

	ToKeyVaultKeyPtrOutput() KeyVaultKeyPtrOutput
	ToKeyVaultKeyPtrOutputWithContext(context.Context) KeyVaultKeyPtrOutput
}

type keyVaultKeyPtrType KeyVaultKeyArgs

func KeyVaultKeyPtr(v *KeyVaultKeyArgs) KeyVaultKeyPtrInput {
	return (*keyVaultKeyPtrType)(v)
}

func (*keyVaultKeyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultKey)(nil)).Elem()
}

func (i *keyVaultKeyPtrType) ToKeyVaultKeyPtrOutput() KeyVaultKeyPtrOutput {
	return i.ToKeyVaultKeyPtrOutputWithContext(context.Background())
}

func (i *keyVaultKeyPtrType) ToKeyVaultKeyPtrOutputWithContext(ctx context.Context) KeyVaultKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultKeyPtrOutput)
}

type KeyVaultKeyOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultKey)(nil)).Elem()
}

func (o KeyVaultKeyOutput) ToKeyVaultKeyOutput() KeyVaultKeyOutput {
	return o
}

func (o KeyVaultKeyOutput) ToKeyVaultKeyOutputWithContext(ctx context.Context) KeyVaultKeyOutput {
	return o
}

func (o KeyVaultKeyOutput) ToKeyVaultKeyPtrOutput() KeyVaultKeyPtrOutput {
	return o.ToKeyVaultKeyPtrOutputWithContext(context.Background())
}

func (o KeyVaultKeyOutput) ToKeyVaultKeyPtrOutputWithContext(ctx context.Context) KeyVaultKeyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultKey) *KeyVaultKey {
		return &v
	}).(KeyVaultKeyPtrOutput)
}

func (o KeyVaultKeyOutput) KeyUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultKey) *string { return v.KeyUrl }).(pulumi.StringPtrOutput)
}

type KeyVaultKeyPtrOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultKey)(nil)).Elem()
}

func (o KeyVaultKeyPtrOutput) ToKeyVaultKeyPtrOutput() KeyVaultKeyPtrOutput {
	return o
}

func (o KeyVaultKeyPtrOutput) ToKeyVaultKeyPtrOutputWithContext(ctx context.Context) KeyVaultKeyPtrOutput {
	return o
}

func (o KeyVaultKeyPtrOutput) Elem() KeyVaultKeyOutput {
	return o.ApplyT(func(v *KeyVaultKey) KeyVaultKey {
		if v != nil {
			return *v
		}
		var ret KeyVaultKey
		return ret
	}).(KeyVaultKeyOutput)
}

func (o KeyVaultKeyPtrOutput) KeyUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultKey) *string {
		if v == nil {
			return nil
		}
		return v.KeyUrl
	}).(pulumi.StringPtrOutput)
}

type KeyVaultKeyResponse struct {
	KeyUrl *string `pulumi:"keyUrl"`
}

type KeyVaultKeyResponseOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultKeyResponse)(nil)).Elem()
}

func (o KeyVaultKeyResponseOutput) ToKeyVaultKeyResponseOutput() KeyVaultKeyResponseOutput {
	return o
}

func (o KeyVaultKeyResponseOutput) ToKeyVaultKeyResponseOutputWithContext(ctx context.Context) KeyVaultKeyResponseOutput {
	return o
}

func (o KeyVaultKeyResponseOutput) KeyUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultKeyResponse) *string { return v.KeyUrl }).(pulumi.StringPtrOutput)
}

type KeyVaultKeyResponsePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultKeyResponse)(nil)).Elem()
}

func (o KeyVaultKeyResponsePtrOutput) ToKeyVaultKeyResponsePtrOutput() KeyVaultKeyResponsePtrOutput {
	return o
}

func (o KeyVaultKeyResponsePtrOutput) ToKeyVaultKeyResponsePtrOutputWithContext(ctx context.Context) KeyVaultKeyResponsePtrOutput {
	return o
}

func (o KeyVaultKeyResponsePtrOutput) Elem() KeyVaultKeyResponseOutput {
	return o.ApplyT(func(v *KeyVaultKeyResponse) KeyVaultKeyResponse {
		if v != nil {
			return *v
		}
		var ret KeyVaultKeyResponse
		return ret
	}).(KeyVaultKeyResponseOutput)
}

func (o KeyVaultKeyResponsePtrOutput) KeyUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultKeyResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyUrl
	}).(pulumi.StringPtrOutput)
}

type LocalDiagnosticsAccessConfiguration struct {
	HttpsServerCertificate *KeyVaultCertificate `pulumi:"httpsServerCertificate"`
}





type LocalDiagnosticsAccessConfigurationInput interface {
	pulumi.Input

	ToLocalDiagnosticsAccessConfigurationOutput() LocalDiagnosticsAccessConfigurationOutput
	ToLocalDiagnosticsAccessConfigurationOutputWithContext(context.Context) LocalDiagnosticsAccessConfigurationOutput
}

type LocalDiagnosticsAccessConfigurationArgs struct {
	HttpsServerCertificate KeyVaultCertificatePtrInput `pulumi:"httpsServerCertificate"`
}

func (LocalDiagnosticsAccessConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LocalDiagnosticsAccessConfiguration)(nil)).Elem()
}

func (i LocalDiagnosticsAccessConfigurationArgs) ToLocalDiagnosticsAccessConfigurationOutput() LocalDiagnosticsAccessConfigurationOutput {
	return i.ToLocalDiagnosticsAccessConfigurationOutputWithContext(context.Background())
}

func (i LocalDiagnosticsAccessConfigurationArgs) ToLocalDiagnosticsAccessConfigurationOutputWithContext(ctx context.Context) LocalDiagnosticsAccessConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalDiagnosticsAccessConfigurationOutput)
}

func (i LocalDiagnosticsAccessConfigurationArgs) ToLocalDiagnosticsAccessConfigurationPtrOutput() LocalDiagnosticsAccessConfigurationPtrOutput {
	return i.ToLocalDiagnosticsAccessConfigurationPtrOutputWithContext(context.Background())
}

func (i LocalDiagnosticsAccessConfigurationArgs) ToLocalDiagnosticsAccessConfigurationPtrOutputWithContext(ctx context.Context) LocalDiagnosticsAccessConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalDiagnosticsAccessConfigurationOutput).ToLocalDiagnosticsAccessConfigurationPtrOutputWithContext(ctx)
}









type LocalDiagnosticsAccessConfigurationPtrInput interface {
	pulumi.Input

	ToLocalDiagnosticsAccessConfigurationPtrOutput() LocalDiagnosticsAccessConfigurationPtrOutput
	ToLocalDiagnosticsAccessConfigurationPtrOutputWithContext(context.Context) LocalDiagnosticsAccessConfigurationPtrOutput
}

type localDiagnosticsAccessConfigurationPtrType LocalDiagnosticsAccessConfigurationArgs

func LocalDiagnosticsAccessConfigurationPtr(v *LocalDiagnosticsAccessConfigurationArgs) LocalDiagnosticsAccessConfigurationPtrInput {
	return (*localDiagnosticsAccessConfigurationPtrType)(v)
}

func (*localDiagnosticsAccessConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalDiagnosticsAccessConfiguration)(nil)).Elem()
}

func (i *localDiagnosticsAccessConfigurationPtrType) ToLocalDiagnosticsAccessConfigurationPtrOutput() LocalDiagnosticsAccessConfigurationPtrOutput {
	return i.ToLocalDiagnosticsAccessConfigurationPtrOutputWithContext(context.Background())
}

func (i *localDiagnosticsAccessConfigurationPtrType) ToLocalDiagnosticsAccessConfigurationPtrOutputWithContext(ctx context.Context) LocalDiagnosticsAccessConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalDiagnosticsAccessConfigurationPtrOutput)
}

type LocalDiagnosticsAccessConfigurationOutput struct{ *pulumi.OutputState }

func (LocalDiagnosticsAccessConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LocalDiagnosticsAccessConfiguration)(nil)).Elem()
}

func (o LocalDiagnosticsAccessConfigurationOutput) ToLocalDiagnosticsAccessConfigurationOutput() LocalDiagnosticsAccessConfigurationOutput {
	return o
}

func (o LocalDiagnosticsAccessConfigurationOutput) ToLocalDiagnosticsAccessConfigurationOutputWithContext(ctx context.Context) LocalDiagnosticsAccessConfigurationOutput {
	return o
}

func (o LocalDiagnosticsAccessConfigurationOutput) ToLocalDiagnosticsAccessConfigurationPtrOutput() LocalDiagnosticsAccessConfigurationPtrOutput {
	return o.ToLocalDiagnosticsAccessConfigurationPtrOutputWithContext(context.Background())
}

func (o LocalDiagnosticsAccessConfigurationOutput) ToLocalDiagnosticsAccessConfigurationPtrOutputWithContext(ctx context.Context) LocalDiagnosticsAccessConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LocalDiagnosticsAccessConfiguration) *LocalDiagnosticsAccessConfiguration {
		return &v
	}).(LocalDiagnosticsAccessConfigurationPtrOutput)
}

func (o LocalDiagnosticsAccessConfigurationOutput) HttpsServerCertificate() KeyVaultCertificatePtrOutput {
	return o.ApplyT(func(v LocalDiagnosticsAccessConfiguration) *KeyVaultCertificate { return v.HttpsServerCertificate }).(KeyVaultCertificatePtrOutput)
}

type LocalDiagnosticsAccessConfigurationPtrOutput struct{ *pulumi.OutputState }

func (LocalDiagnosticsAccessConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalDiagnosticsAccessConfiguration)(nil)).Elem()
}

func (o LocalDiagnosticsAccessConfigurationPtrOutput) ToLocalDiagnosticsAccessConfigurationPtrOutput() LocalDiagnosticsAccessConfigurationPtrOutput {
	return o
}

func (o LocalDiagnosticsAccessConfigurationPtrOutput) ToLocalDiagnosticsAccessConfigurationPtrOutputWithContext(ctx context.Context) LocalDiagnosticsAccessConfigurationPtrOutput {
	return o
}

func (o LocalDiagnosticsAccessConfigurationPtrOutput) Elem() LocalDiagnosticsAccessConfigurationOutput {
	return o.ApplyT(func(v *LocalDiagnosticsAccessConfiguration) LocalDiagnosticsAccessConfiguration {
		if v != nil {
			return *v
		}
		var ret LocalDiagnosticsAccessConfiguration
		return ret
	}).(LocalDiagnosticsAccessConfigurationOutput)
}

func (o LocalDiagnosticsAccessConfigurationPtrOutput) HttpsServerCertificate() KeyVaultCertificatePtrOutput {
	return o.ApplyT(func(v *LocalDiagnosticsAccessConfiguration) *KeyVaultCertificate {
		if v == nil {
			return nil
		}
		return v.HttpsServerCertificate
	}).(KeyVaultCertificatePtrOutput)
}

type LocalDiagnosticsAccessConfigurationResponse struct {
	HttpsServerCertificate *KeyVaultCertificateResponse `pulumi:"httpsServerCertificate"`
}

type LocalDiagnosticsAccessConfigurationResponseOutput struct{ *pulumi.OutputState }

func (LocalDiagnosticsAccessConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LocalDiagnosticsAccessConfigurationResponse)(nil)).Elem()
}

func (o LocalDiagnosticsAccessConfigurationResponseOutput) ToLocalDiagnosticsAccessConfigurationResponseOutput() LocalDiagnosticsAccessConfigurationResponseOutput {
	return o
}

func (o LocalDiagnosticsAccessConfigurationResponseOutput) ToLocalDiagnosticsAccessConfigurationResponseOutputWithContext(ctx context.Context) LocalDiagnosticsAccessConfigurationResponseOutput {
	return o
}

func (o LocalDiagnosticsAccessConfigurationResponseOutput) HttpsServerCertificate() KeyVaultCertificateResponsePtrOutput {
	return o.ApplyT(func(v LocalDiagnosticsAccessConfigurationResponse) *KeyVaultCertificateResponse {
		return v.HttpsServerCertificate
	}).(KeyVaultCertificateResponsePtrOutput)
}

type LocalDiagnosticsAccessConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (LocalDiagnosticsAccessConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalDiagnosticsAccessConfigurationResponse)(nil)).Elem()
}

func (o LocalDiagnosticsAccessConfigurationResponsePtrOutput) ToLocalDiagnosticsAccessConfigurationResponsePtrOutput() LocalDiagnosticsAccessConfigurationResponsePtrOutput {
	return o
}

func (o LocalDiagnosticsAccessConfigurationResponsePtrOutput) ToLocalDiagnosticsAccessConfigurationResponsePtrOutputWithContext(ctx context.Context) LocalDiagnosticsAccessConfigurationResponsePtrOutput {
	return o
}

func (o LocalDiagnosticsAccessConfigurationResponsePtrOutput) Elem() LocalDiagnosticsAccessConfigurationResponseOutput {
	return o.ApplyT(func(v *LocalDiagnosticsAccessConfigurationResponse) LocalDiagnosticsAccessConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret LocalDiagnosticsAccessConfigurationResponse
		return ret
	}).(LocalDiagnosticsAccessConfigurationResponseOutput)
}

func (o LocalDiagnosticsAccessConfigurationResponsePtrOutput) HttpsServerCertificate() KeyVaultCertificateResponsePtrOutput {
	return o.ApplyT(func(v *LocalDiagnosticsAccessConfigurationResponse) *KeyVaultCertificateResponse {
		if v == nil {
			return nil
		}
		return v.HttpsServerCertificate
	}).(KeyVaultCertificateResponsePtrOutput)
}

type ManagedServiceIdentity struct {
	Type                   string                 `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type ManagedServiceIdentityInput interface {
	pulumi.Input

	ToManagedServiceIdentityOutput() ManagedServiceIdentityOutput
	ToManagedServiceIdentityOutputWithContext(context.Context) ManagedServiceIdentityOutput
}

type ManagedServiceIdentityArgs struct {
	Type                   pulumi.StringInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput    `pulumi:"userAssignedIdentities"`
}

func (ManagedServiceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentity)(nil)).Elem()
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityOutput() ManagedServiceIdentityOutput {
	return i.ToManagedServiceIdentityOutputWithContext(context.Background())
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityOutputWithContext(ctx context.Context) ManagedServiceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityOutput)
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return i.ToManagedServiceIdentityPtrOutputWithContext(context.Background())
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityOutput).ToManagedServiceIdentityPtrOutputWithContext(ctx)
}









type ManagedServiceIdentityPtrInput interface {
	pulumi.Input

	ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput
	ToManagedServiceIdentityPtrOutputWithContext(context.Context) ManagedServiceIdentityPtrOutput
}

type managedServiceIdentityPtrType ManagedServiceIdentityArgs

func ManagedServiceIdentityPtr(v *ManagedServiceIdentityArgs) ManagedServiceIdentityPtrInput {
	return (*managedServiceIdentityPtrType)(v)
}

func (*managedServiceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentity)(nil)).Elem()
}

func (i *managedServiceIdentityPtrType) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return i.ToManagedServiceIdentityPtrOutputWithContext(context.Background())
}

func (i *managedServiceIdentityPtrType) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityPtrOutput)
}

type ManagedServiceIdentityOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentity)(nil)).Elem()
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityOutput() ManagedServiceIdentityOutput {
	return o
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityOutputWithContext(ctx context.Context) ManagedServiceIdentityOutput {
	return o
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return o.ToManagedServiceIdentityPtrOutputWithContext(context.Background())
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedServiceIdentity) *ManagedServiceIdentity {
		return &v
	}).(ManagedServiceIdentityPtrOutput)
}

func (o ManagedServiceIdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentity) string { return v.Type }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v ManagedServiceIdentity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type ManagedServiceIdentityPtrOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentity)(nil)).Elem()
}

func (o ManagedServiceIdentityPtrOutput) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return o
}

func (o ManagedServiceIdentityPtrOutput) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return o
}

func (o ManagedServiceIdentityPtrOutput) Elem() ManagedServiceIdentityOutput {
	return o.ApplyT(func(v *ManagedServiceIdentity) ManagedServiceIdentity {
		if v != nil {
			return *v
		}
		var ret ManagedServiceIdentity
		return ret
	}).(ManagedServiceIdentityOutput)
}

func (o ManagedServiceIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *ManagedServiceIdentity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type ManagedServiceIdentityResponse struct {
	PrincipalId            string                                  `pulumi:"principalId"`
	TenantId               string                                  `pulumi:"tenantId"`
	Type                   string                                  `pulumi:"type"`
	UserAssignedIdentities map[string]UserAssignedIdentityResponse `pulumi:"userAssignedIdentities"`
}

type ManagedServiceIdentityResponseOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentityResponse)(nil)).Elem()
}

func (o ManagedServiceIdentityResponseOutput) ToManagedServiceIdentityResponseOutput() ManagedServiceIdentityResponseOutput {
	return o
}

func (o ManagedServiceIdentityResponseOutput) ToManagedServiceIdentityResponseOutputWithContext(ctx context.Context) ManagedServiceIdentityResponseOutput {
	return o
}

func (o ManagedServiceIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityResponseOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) map[string]UserAssignedIdentityResponse {
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
}

type ManagedServiceIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentityResponse)(nil)).Elem()
}

func (o ManagedServiceIdentityResponsePtrOutput) ToManagedServiceIdentityResponsePtrOutput() ManagedServiceIdentityResponsePtrOutput {
	return o
}

func (o ManagedServiceIdentityResponsePtrOutput) ToManagedServiceIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityResponsePtrOutput {
	return o
}

func (o ManagedServiceIdentityResponsePtrOutput) Elem() ManagedServiceIdentityResponseOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) ManagedServiceIdentityResponse {
		if v != nil {
			return *v
		}
		var ret ManagedServiceIdentityResponse
		return ret
	}).(ManagedServiceIdentityResponseOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) map[string]UserAssignedIdentityResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
}

type MobileNetworkResourceId struct {
	Id string `pulumi:"id"`
}





type MobileNetworkResourceIdInput interface {
	pulumi.Input

	ToMobileNetworkResourceIdOutput() MobileNetworkResourceIdOutput
	ToMobileNetworkResourceIdOutputWithContext(context.Context) MobileNetworkResourceIdOutput
}

type MobileNetworkResourceIdArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (MobileNetworkResourceIdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MobileNetworkResourceId)(nil)).Elem()
}

func (i MobileNetworkResourceIdArgs) ToMobileNetworkResourceIdOutput() MobileNetworkResourceIdOutput {
	return i.ToMobileNetworkResourceIdOutputWithContext(context.Background())
}

func (i MobileNetworkResourceIdArgs) ToMobileNetworkResourceIdOutputWithContext(ctx context.Context) MobileNetworkResourceIdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MobileNetworkResourceIdOutput)
}

func (i MobileNetworkResourceIdArgs) ToMobileNetworkResourceIdPtrOutput() MobileNetworkResourceIdPtrOutput {
	return i.ToMobileNetworkResourceIdPtrOutputWithContext(context.Background())
}

func (i MobileNetworkResourceIdArgs) ToMobileNetworkResourceIdPtrOutputWithContext(ctx context.Context) MobileNetworkResourceIdPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MobileNetworkResourceIdOutput).ToMobileNetworkResourceIdPtrOutputWithContext(ctx)
}









type MobileNetworkResourceIdPtrInput interface {
	pulumi.Input

	ToMobileNetworkResourceIdPtrOutput() MobileNetworkResourceIdPtrOutput
	ToMobileNetworkResourceIdPtrOutputWithContext(context.Context) MobileNetworkResourceIdPtrOutput
}

type mobileNetworkResourceIdPtrType MobileNetworkResourceIdArgs

func MobileNetworkResourceIdPtr(v *MobileNetworkResourceIdArgs) MobileNetworkResourceIdPtrInput {
	return (*mobileNetworkResourceIdPtrType)(v)
}

func (*mobileNetworkResourceIdPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MobileNetworkResourceId)(nil)).Elem()
}

func (i *mobileNetworkResourceIdPtrType) ToMobileNetworkResourceIdPtrOutput() MobileNetworkResourceIdPtrOutput {
	return i.ToMobileNetworkResourceIdPtrOutputWithContext(context.Background())
}

func (i *mobileNetworkResourceIdPtrType) ToMobileNetworkResourceIdPtrOutputWithContext(ctx context.Context) MobileNetworkResourceIdPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MobileNetworkResourceIdPtrOutput)
}

type MobileNetworkResourceIdOutput struct{ *pulumi.OutputState }

func (MobileNetworkResourceIdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MobileNetworkResourceId)(nil)).Elem()
}

func (o MobileNetworkResourceIdOutput) ToMobileNetworkResourceIdOutput() MobileNetworkResourceIdOutput {
	return o
}

func (o MobileNetworkResourceIdOutput) ToMobileNetworkResourceIdOutputWithContext(ctx context.Context) MobileNetworkResourceIdOutput {
	return o
}

func (o MobileNetworkResourceIdOutput) ToMobileNetworkResourceIdPtrOutput() MobileNetworkResourceIdPtrOutput {
	return o.ToMobileNetworkResourceIdPtrOutputWithContext(context.Background())
}

func (o MobileNetworkResourceIdOutput) ToMobileNetworkResourceIdPtrOutputWithContext(ctx context.Context) MobileNetworkResourceIdPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MobileNetworkResourceId) *MobileNetworkResourceId {
		return &v
	}).(MobileNetworkResourceIdPtrOutput)
}

func (o MobileNetworkResourceIdOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v MobileNetworkResourceId) string { return v.Id }).(pulumi.StringOutput)
}

type MobileNetworkResourceIdPtrOutput struct{ *pulumi.OutputState }

func (MobileNetworkResourceIdPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MobileNetworkResourceId)(nil)).Elem()
}

func (o MobileNetworkResourceIdPtrOutput) ToMobileNetworkResourceIdPtrOutput() MobileNetworkResourceIdPtrOutput {
	return o
}

func (o MobileNetworkResourceIdPtrOutput) ToMobileNetworkResourceIdPtrOutputWithContext(ctx context.Context) MobileNetworkResourceIdPtrOutput {
	return o
}

func (o MobileNetworkResourceIdPtrOutput) Elem() MobileNetworkResourceIdOutput {
	return o.ApplyT(func(v *MobileNetworkResourceId) MobileNetworkResourceId {
		if v != nil {
			return *v
		}
		var ret MobileNetworkResourceId
		return ret
	}).(MobileNetworkResourceIdOutput)
}

func (o MobileNetworkResourceIdPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MobileNetworkResourceId) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type MobileNetworkResourceIdResponse struct {
	Id string `pulumi:"id"`
}

type MobileNetworkResourceIdResponseOutput struct{ *pulumi.OutputState }

func (MobileNetworkResourceIdResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MobileNetworkResourceIdResponse)(nil)).Elem()
}

func (o MobileNetworkResourceIdResponseOutput) ToMobileNetworkResourceIdResponseOutput() MobileNetworkResourceIdResponseOutput {
	return o
}

func (o MobileNetworkResourceIdResponseOutput) ToMobileNetworkResourceIdResponseOutputWithContext(ctx context.Context) MobileNetworkResourceIdResponseOutput {
	return o
}

func (o MobileNetworkResourceIdResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v MobileNetworkResourceIdResponse) string { return v.Id }).(pulumi.StringOutput)
}

type MobileNetworkResourceIdResponsePtrOutput struct{ *pulumi.OutputState }

func (MobileNetworkResourceIdResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MobileNetworkResourceIdResponse)(nil)).Elem()
}

func (o MobileNetworkResourceIdResponsePtrOutput) ToMobileNetworkResourceIdResponsePtrOutput() MobileNetworkResourceIdResponsePtrOutput {
	return o
}

func (o MobileNetworkResourceIdResponsePtrOutput) ToMobileNetworkResourceIdResponsePtrOutputWithContext(ctx context.Context) MobileNetworkResourceIdResponsePtrOutput {
	return o
}

func (o MobileNetworkResourceIdResponsePtrOutput) Elem() MobileNetworkResourceIdResponseOutput {
	return o.ApplyT(func(v *MobileNetworkResourceIdResponse) MobileNetworkResourceIdResponse {
		if v != nil {
			return *v
		}
		var ret MobileNetworkResourceIdResponse
		return ret
	}).(MobileNetworkResourceIdResponseOutput)
}

func (o MobileNetworkResourceIdResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MobileNetworkResourceIdResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type NaptConfiguration struct {
	Enabled           *string             `pulumi:"enabled"`
	PinholeLimits     *int                `pulumi:"pinholeLimits"`
	PinholeTimeouts   *PinholeTimeouts    `pulumi:"pinholeTimeouts"`
	PortRange         *PortRange          `pulumi:"portRange"`
	PortReuseHoldTime *PortReuseHoldTimes `pulumi:"portReuseHoldTime"`
}


func (val *NaptConfiguration) Defaults() *NaptConfiguration {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.PinholeLimits) {
		pinholeLimits_ := 65536
		tmp.PinholeLimits = &pinholeLimits_
	}
	tmp.PinholeTimeouts = tmp.PinholeTimeouts.Defaults()

	tmp.PortRange = tmp.PortRange.Defaults()

	tmp.PortReuseHoldTime = tmp.PortReuseHoldTime.Defaults()

	return &tmp
}





type NaptConfigurationInput interface {
	pulumi.Input

	ToNaptConfigurationOutput() NaptConfigurationOutput
	ToNaptConfigurationOutputWithContext(context.Context) NaptConfigurationOutput
}

type NaptConfigurationArgs struct {
	Enabled           pulumi.StringPtrInput      `pulumi:"enabled"`
	PinholeLimits     pulumi.IntPtrInput         `pulumi:"pinholeLimits"`
	PinholeTimeouts   PinholeTimeoutsPtrInput    `pulumi:"pinholeTimeouts"`
	PortRange         PortRangePtrInput          `pulumi:"portRange"`
	PortReuseHoldTime PortReuseHoldTimesPtrInput `pulumi:"portReuseHoldTime"`
}


func (val *NaptConfigurationArgs) Defaults() *NaptConfigurationArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.PinholeLimits) {
		tmp.PinholeLimits = pulumi.IntPtr(65536)
	}

	return &tmp
}
func (NaptConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NaptConfiguration)(nil)).Elem()
}

func (i NaptConfigurationArgs) ToNaptConfigurationOutput() NaptConfigurationOutput {
	return i.ToNaptConfigurationOutputWithContext(context.Background())
}

func (i NaptConfigurationArgs) ToNaptConfigurationOutputWithContext(ctx context.Context) NaptConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NaptConfigurationOutput)
}

func (i NaptConfigurationArgs) ToNaptConfigurationPtrOutput() NaptConfigurationPtrOutput {
	return i.ToNaptConfigurationPtrOutputWithContext(context.Background())
}

func (i NaptConfigurationArgs) ToNaptConfigurationPtrOutputWithContext(ctx context.Context) NaptConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NaptConfigurationOutput).ToNaptConfigurationPtrOutputWithContext(ctx)
}









type NaptConfigurationPtrInput interface {
	pulumi.Input

	ToNaptConfigurationPtrOutput() NaptConfigurationPtrOutput
	ToNaptConfigurationPtrOutputWithContext(context.Context) NaptConfigurationPtrOutput
}

type naptConfigurationPtrType NaptConfigurationArgs

func NaptConfigurationPtr(v *NaptConfigurationArgs) NaptConfigurationPtrInput {
	return (*naptConfigurationPtrType)(v)
}

func (*naptConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NaptConfiguration)(nil)).Elem()
}

func (i *naptConfigurationPtrType) ToNaptConfigurationPtrOutput() NaptConfigurationPtrOutput {
	return i.ToNaptConfigurationPtrOutputWithContext(context.Background())
}

func (i *naptConfigurationPtrType) ToNaptConfigurationPtrOutputWithContext(ctx context.Context) NaptConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NaptConfigurationPtrOutput)
}

type NaptConfigurationOutput struct{ *pulumi.OutputState }

func (NaptConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NaptConfiguration)(nil)).Elem()
}

func (o NaptConfigurationOutput) ToNaptConfigurationOutput() NaptConfigurationOutput {
	return o
}

func (o NaptConfigurationOutput) ToNaptConfigurationOutputWithContext(ctx context.Context) NaptConfigurationOutput {
	return o
}

func (o NaptConfigurationOutput) ToNaptConfigurationPtrOutput() NaptConfigurationPtrOutput {
	return o.ToNaptConfigurationPtrOutputWithContext(context.Background())
}

func (o NaptConfigurationOutput) ToNaptConfigurationPtrOutputWithContext(ctx context.Context) NaptConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NaptConfiguration) *NaptConfiguration {
		return &v
	}).(NaptConfigurationPtrOutput)
}

func (o NaptConfigurationOutput) Enabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NaptConfiguration) *string { return v.Enabled }).(pulumi.StringPtrOutput)
}

func (o NaptConfigurationOutput) PinholeLimits() pulumi.IntPtrOutput {
	return o.ApplyT(func(v NaptConfiguration) *int { return v.PinholeLimits }).(pulumi.IntPtrOutput)
}

func (o NaptConfigurationOutput) PinholeTimeouts() PinholeTimeoutsPtrOutput {
	return o.ApplyT(func(v NaptConfiguration) *PinholeTimeouts { return v.PinholeTimeouts }).(PinholeTimeoutsPtrOutput)
}

func (o NaptConfigurationOutput) PortRange() PortRangePtrOutput {
	return o.ApplyT(func(v NaptConfiguration) *PortRange { return v.PortRange }).(PortRangePtrOutput)
}

func (o NaptConfigurationOutput) PortReuseHoldTime() PortReuseHoldTimesPtrOutput {
	return o.ApplyT(func(v NaptConfiguration) *PortReuseHoldTimes { return v.PortReuseHoldTime }).(PortReuseHoldTimesPtrOutput)
}

type NaptConfigurationPtrOutput struct{ *pulumi.OutputState }

func (NaptConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NaptConfiguration)(nil)).Elem()
}

func (o NaptConfigurationPtrOutput) ToNaptConfigurationPtrOutput() NaptConfigurationPtrOutput {
	return o
}

func (o NaptConfigurationPtrOutput) ToNaptConfigurationPtrOutputWithContext(ctx context.Context) NaptConfigurationPtrOutput {
	return o
}

func (o NaptConfigurationPtrOutput) Elem() NaptConfigurationOutput {
	return o.ApplyT(func(v *NaptConfiguration) NaptConfiguration {
		if v != nil {
			return *v
		}
		var ret NaptConfiguration
		return ret
	}).(NaptConfigurationOutput)
}

func (o NaptConfigurationPtrOutput) Enabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NaptConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.StringPtrOutput)
}

func (o NaptConfigurationPtrOutput) PinholeLimits() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NaptConfiguration) *int {
		if v == nil {
			return nil
		}
		return v.PinholeLimits
	}).(pulumi.IntPtrOutput)
}

func (o NaptConfigurationPtrOutput) PinholeTimeouts() PinholeTimeoutsPtrOutput {
	return o.ApplyT(func(v *NaptConfiguration) *PinholeTimeouts {
		if v == nil {
			return nil
		}
		return v.PinholeTimeouts
	}).(PinholeTimeoutsPtrOutput)
}

func (o NaptConfigurationPtrOutput) PortRange() PortRangePtrOutput {
	return o.ApplyT(func(v *NaptConfiguration) *PortRange {
		if v == nil {
			return nil
		}
		return v.PortRange
	}).(PortRangePtrOutput)
}

func (o NaptConfigurationPtrOutput) PortReuseHoldTime() PortReuseHoldTimesPtrOutput {
	return o.ApplyT(func(v *NaptConfiguration) *PortReuseHoldTimes {
		if v == nil {
			return nil
		}
		return v.PortReuseHoldTime
	}).(PortReuseHoldTimesPtrOutput)
}

type NaptConfigurationResponse struct {
	Enabled           *string                     `pulumi:"enabled"`
	PinholeLimits     *int                        `pulumi:"pinholeLimits"`
	PinholeTimeouts   *PinholeTimeoutsResponse    `pulumi:"pinholeTimeouts"`
	PortRange         *PortRangeResponse          `pulumi:"portRange"`
	PortReuseHoldTime *PortReuseHoldTimesResponse `pulumi:"portReuseHoldTime"`
}


func (val *NaptConfigurationResponse) Defaults() *NaptConfigurationResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.PinholeLimits) {
		pinholeLimits_ := 65536
		tmp.PinholeLimits = &pinholeLimits_
	}
	tmp.PinholeTimeouts = tmp.PinholeTimeouts.Defaults()

	tmp.PortRange = tmp.PortRange.Defaults()

	tmp.PortReuseHoldTime = tmp.PortReuseHoldTime.Defaults()

	return &tmp
}

type NaptConfigurationResponseOutput struct{ *pulumi.OutputState }

func (NaptConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NaptConfigurationResponse)(nil)).Elem()
}

func (o NaptConfigurationResponseOutput) ToNaptConfigurationResponseOutput() NaptConfigurationResponseOutput {
	return o
}

func (o NaptConfigurationResponseOutput) ToNaptConfigurationResponseOutputWithContext(ctx context.Context) NaptConfigurationResponseOutput {
	return o
}

func (o NaptConfigurationResponseOutput) Enabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NaptConfigurationResponse) *string { return v.Enabled }).(pulumi.StringPtrOutput)
}

func (o NaptConfigurationResponseOutput) PinholeLimits() pulumi.IntPtrOutput {
	return o.ApplyT(func(v NaptConfigurationResponse) *int { return v.PinholeLimits }).(pulumi.IntPtrOutput)
}

func (o NaptConfigurationResponseOutput) PinholeTimeouts() PinholeTimeoutsResponsePtrOutput {
	return o.ApplyT(func(v NaptConfigurationResponse) *PinholeTimeoutsResponse { return v.PinholeTimeouts }).(PinholeTimeoutsResponsePtrOutput)
}

func (o NaptConfigurationResponseOutput) PortRange() PortRangeResponsePtrOutput {
	return o.ApplyT(func(v NaptConfigurationResponse) *PortRangeResponse { return v.PortRange }).(PortRangeResponsePtrOutput)
}

func (o NaptConfigurationResponseOutput) PortReuseHoldTime() PortReuseHoldTimesResponsePtrOutput {
	return o.ApplyT(func(v NaptConfigurationResponse) *PortReuseHoldTimesResponse { return v.PortReuseHoldTime }).(PortReuseHoldTimesResponsePtrOutput)
}

type NaptConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (NaptConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NaptConfigurationResponse)(nil)).Elem()
}

func (o NaptConfigurationResponsePtrOutput) ToNaptConfigurationResponsePtrOutput() NaptConfigurationResponsePtrOutput {
	return o
}

func (o NaptConfigurationResponsePtrOutput) ToNaptConfigurationResponsePtrOutputWithContext(ctx context.Context) NaptConfigurationResponsePtrOutput {
	return o
}

func (o NaptConfigurationResponsePtrOutput) Elem() NaptConfigurationResponseOutput {
	return o.ApplyT(func(v *NaptConfigurationResponse) NaptConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret NaptConfigurationResponse
		return ret
	}).(NaptConfigurationResponseOutput)
}

func (o NaptConfigurationResponsePtrOutput) Enabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NaptConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.StringPtrOutput)
}

func (o NaptConfigurationResponsePtrOutput) PinholeLimits() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NaptConfigurationResponse) *int {
		if v == nil {
			return nil
		}
		return v.PinholeLimits
	}).(pulumi.IntPtrOutput)
}

func (o NaptConfigurationResponsePtrOutput) PinholeTimeouts() PinholeTimeoutsResponsePtrOutput {
	return o.ApplyT(func(v *NaptConfigurationResponse) *PinholeTimeoutsResponse {
		if v == nil {
			return nil
		}
		return v.PinholeTimeouts
	}).(PinholeTimeoutsResponsePtrOutput)
}

func (o NaptConfigurationResponsePtrOutput) PortRange() PortRangeResponsePtrOutput {
	return o.ApplyT(func(v *NaptConfigurationResponse) *PortRangeResponse {
		if v == nil {
			return nil
		}
		return v.PortRange
	}).(PortRangeResponsePtrOutput)
}

func (o NaptConfigurationResponsePtrOutput) PortReuseHoldTime() PortReuseHoldTimesResponsePtrOutput {
	return o.ApplyT(func(v *NaptConfigurationResponse) *PortReuseHoldTimesResponse {
		if v == nil {
			return nil
		}
		return v.PortReuseHoldTime
	}).(PortReuseHoldTimesResponsePtrOutput)
}

type PccRuleConfiguration struct {
	RuleName                 string                    `pulumi:"ruleName"`
	RulePrecedence           int                       `pulumi:"rulePrecedence"`
	RuleQosPolicy            *PccRuleQosPolicy         `pulumi:"ruleQosPolicy"`
	ServiceDataFlowTemplates []ServiceDataFlowTemplate `pulumi:"serviceDataFlowTemplates"`
	TrafficControl           *string                   `pulumi:"trafficControl"`
}


func (val *PccRuleConfiguration) Defaults() *PccRuleConfiguration {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.RuleQosPolicy = tmp.RuleQosPolicy.Defaults()

	if isZero(tmp.TrafficControl) {
		trafficControl_ := "Enabled"
		tmp.TrafficControl = &trafficControl_
	}
	return &tmp
}





type PccRuleConfigurationInput interface {
	pulumi.Input

	ToPccRuleConfigurationOutput() PccRuleConfigurationOutput
	ToPccRuleConfigurationOutputWithContext(context.Context) PccRuleConfigurationOutput
}

type PccRuleConfigurationArgs struct {
	RuleName                 pulumi.StringInput                `pulumi:"ruleName"`
	RulePrecedence           pulumi.IntInput                   `pulumi:"rulePrecedence"`
	RuleQosPolicy            PccRuleQosPolicyPtrInput          `pulumi:"ruleQosPolicy"`
	ServiceDataFlowTemplates ServiceDataFlowTemplateArrayInput `pulumi:"serviceDataFlowTemplates"`
	TrafficControl           pulumi.StringPtrInput             `pulumi:"trafficControl"`
}


func (val *PccRuleConfigurationArgs) Defaults() *PccRuleConfigurationArgs {
	if val == nil {
		return nil
	}
	tmp := *val

	if isZero(tmp.TrafficControl) {
		tmp.TrafficControl = pulumi.StringPtr("Enabled")
	}
	return &tmp
}
func (PccRuleConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PccRuleConfiguration)(nil)).Elem()
}

func (i PccRuleConfigurationArgs) ToPccRuleConfigurationOutput() PccRuleConfigurationOutput {
	return i.ToPccRuleConfigurationOutputWithContext(context.Background())
}

func (i PccRuleConfigurationArgs) ToPccRuleConfigurationOutputWithContext(ctx context.Context) PccRuleConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PccRuleConfigurationOutput)
}





type PccRuleConfigurationArrayInput interface {
	pulumi.Input

	ToPccRuleConfigurationArrayOutput() PccRuleConfigurationArrayOutput
	ToPccRuleConfigurationArrayOutputWithContext(context.Context) PccRuleConfigurationArrayOutput
}

type PccRuleConfigurationArray []PccRuleConfigurationInput

func (PccRuleConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PccRuleConfiguration)(nil)).Elem()
}

func (i PccRuleConfigurationArray) ToPccRuleConfigurationArrayOutput() PccRuleConfigurationArrayOutput {
	return i.ToPccRuleConfigurationArrayOutputWithContext(context.Background())
}

func (i PccRuleConfigurationArray) ToPccRuleConfigurationArrayOutputWithContext(ctx context.Context) PccRuleConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PccRuleConfigurationArrayOutput)
}

type PccRuleConfigurationOutput struct{ *pulumi.OutputState }

func (PccRuleConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PccRuleConfiguration)(nil)).Elem()
}

func (o PccRuleConfigurationOutput) ToPccRuleConfigurationOutput() PccRuleConfigurationOutput {
	return o
}

func (o PccRuleConfigurationOutput) ToPccRuleConfigurationOutputWithContext(ctx context.Context) PccRuleConfigurationOutput {
	return o
}

func (o PccRuleConfigurationOutput) RuleName() pulumi.StringOutput {
	return o.ApplyT(func(v PccRuleConfiguration) string { return v.RuleName }).(pulumi.StringOutput)
}

func (o PccRuleConfigurationOutput) RulePrecedence() pulumi.IntOutput {
	return o.ApplyT(func(v PccRuleConfiguration) int { return v.RulePrecedence }).(pulumi.IntOutput)
}

func (o PccRuleConfigurationOutput) RuleQosPolicy() PccRuleQosPolicyPtrOutput {
	return o.ApplyT(func(v PccRuleConfiguration) *PccRuleQosPolicy { return v.RuleQosPolicy }).(PccRuleQosPolicyPtrOutput)
}

func (o PccRuleConfigurationOutput) ServiceDataFlowTemplates() ServiceDataFlowTemplateArrayOutput {
	return o.ApplyT(func(v PccRuleConfiguration) []ServiceDataFlowTemplate { return v.ServiceDataFlowTemplates }).(ServiceDataFlowTemplateArrayOutput)
}

func (o PccRuleConfigurationOutput) TrafficControl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PccRuleConfiguration) *string { return v.TrafficControl }).(pulumi.StringPtrOutput)
}

type PccRuleConfigurationArrayOutput struct{ *pulumi.OutputState }

func (PccRuleConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PccRuleConfiguration)(nil)).Elem()
}

func (o PccRuleConfigurationArrayOutput) ToPccRuleConfigurationArrayOutput() PccRuleConfigurationArrayOutput {
	return o
}

func (o PccRuleConfigurationArrayOutput) ToPccRuleConfigurationArrayOutputWithContext(ctx context.Context) PccRuleConfigurationArrayOutput {
	return o
}

func (o PccRuleConfigurationArrayOutput) Index(i pulumi.IntInput) PccRuleConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PccRuleConfiguration {
		return vs[0].([]PccRuleConfiguration)[vs[1].(int)]
	}).(PccRuleConfigurationOutput)
}

type PccRuleConfigurationResponse struct {
	RuleName                 string                            `pulumi:"ruleName"`
	RulePrecedence           int                               `pulumi:"rulePrecedence"`
	RuleQosPolicy            *PccRuleQosPolicyResponse         `pulumi:"ruleQosPolicy"`
	ServiceDataFlowTemplates []ServiceDataFlowTemplateResponse `pulumi:"serviceDataFlowTemplates"`
	TrafficControl           *string                           `pulumi:"trafficControl"`
}


func (val *PccRuleConfigurationResponse) Defaults() *PccRuleConfigurationResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.RuleQosPolicy = tmp.RuleQosPolicy.Defaults()

	if isZero(tmp.TrafficControl) {
		trafficControl_ := "Enabled"
		tmp.TrafficControl = &trafficControl_
	}
	return &tmp
}

type PccRuleConfigurationResponseOutput struct{ *pulumi.OutputState }

func (PccRuleConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PccRuleConfigurationResponse)(nil)).Elem()
}

func (o PccRuleConfigurationResponseOutput) ToPccRuleConfigurationResponseOutput() PccRuleConfigurationResponseOutput {
	return o
}

func (o PccRuleConfigurationResponseOutput) ToPccRuleConfigurationResponseOutputWithContext(ctx context.Context) PccRuleConfigurationResponseOutput {
	return o
}

func (o PccRuleConfigurationResponseOutput) RuleName() pulumi.StringOutput {
	return o.ApplyT(func(v PccRuleConfigurationResponse) string { return v.RuleName }).(pulumi.StringOutput)
}

func (o PccRuleConfigurationResponseOutput) RulePrecedence() pulumi.IntOutput {
	return o.ApplyT(func(v PccRuleConfigurationResponse) int { return v.RulePrecedence }).(pulumi.IntOutput)
}

func (o PccRuleConfigurationResponseOutput) RuleQosPolicy() PccRuleQosPolicyResponsePtrOutput {
	return o.ApplyT(func(v PccRuleConfigurationResponse) *PccRuleQosPolicyResponse { return v.RuleQosPolicy }).(PccRuleQosPolicyResponsePtrOutput)
}

func (o PccRuleConfigurationResponseOutput) ServiceDataFlowTemplates() ServiceDataFlowTemplateResponseArrayOutput {
	return o.ApplyT(func(v PccRuleConfigurationResponse) []ServiceDataFlowTemplateResponse {
		return v.ServiceDataFlowTemplates
	}).(ServiceDataFlowTemplateResponseArrayOutput)
}

func (o PccRuleConfigurationResponseOutput) TrafficControl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PccRuleConfigurationResponse) *string { return v.TrafficControl }).(pulumi.StringPtrOutput)
}

type PccRuleConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (PccRuleConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PccRuleConfigurationResponse)(nil)).Elem()
}

func (o PccRuleConfigurationResponseArrayOutput) ToPccRuleConfigurationResponseArrayOutput() PccRuleConfigurationResponseArrayOutput {
	return o
}

func (o PccRuleConfigurationResponseArrayOutput) ToPccRuleConfigurationResponseArrayOutputWithContext(ctx context.Context) PccRuleConfigurationResponseArrayOutput {
	return o
}

func (o PccRuleConfigurationResponseArrayOutput) Index(i pulumi.IntInput) PccRuleConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PccRuleConfigurationResponse {
		return vs[0].([]PccRuleConfigurationResponse)[vs[1].(int)]
	}).(PccRuleConfigurationResponseOutput)
}

type PccRuleQosPolicy struct {
	AllocationAndRetentionPriorityLevel *int    `pulumi:"allocationAndRetentionPriorityLevel"`
	FiveQi                              *int    `pulumi:"fiveQi"`
	GuaranteedBitRate                   *Ambr   `pulumi:"guaranteedBitRate"`
	MaximumBitRate                      Ambr    `pulumi:"maximumBitRate"`
	PreemptionCapability                *string `pulumi:"preemptionCapability"`
	PreemptionVulnerability             *string `pulumi:"preemptionVulnerability"`
}


func (val *PccRuleQosPolicy) Defaults() *PccRuleQosPolicy {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AllocationAndRetentionPriorityLevel) {
		allocationAndRetentionPriorityLevel_ := 9
		tmp.AllocationAndRetentionPriorityLevel = &allocationAndRetentionPriorityLevel_
	}
	if isZero(tmp.FiveQi) {
		fiveQi_ := 9
		tmp.FiveQi = &fiveQi_
	}
	if isZero(tmp.PreemptionCapability) {
		preemptionCapability_ := "NotPreempt"
		tmp.PreemptionCapability = &preemptionCapability_
	}
	if isZero(tmp.PreemptionVulnerability) {
		preemptionVulnerability_ := "Preemptable"
		tmp.PreemptionVulnerability = &preemptionVulnerability_
	}
	return &tmp
}





type PccRuleQosPolicyInput interface {
	pulumi.Input

	ToPccRuleQosPolicyOutput() PccRuleQosPolicyOutput
	ToPccRuleQosPolicyOutputWithContext(context.Context) PccRuleQosPolicyOutput
}

type PccRuleQosPolicyArgs struct {
	AllocationAndRetentionPriorityLevel pulumi.IntPtrInput    `pulumi:"allocationAndRetentionPriorityLevel"`
	FiveQi                              pulumi.IntPtrInput    `pulumi:"fiveQi"`
	GuaranteedBitRate                   AmbrPtrInput          `pulumi:"guaranteedBitRate"`
	MaximumBitRate                      AmbrInput             `pulumi:"maximumBitRate"`
	PreemptionCapability                pulumi.StringPtrInput `pulumi:"preemptionCapability"`
	PreemptionVulnerability             pulumi.StringPtrInput `pulumi:"preemptionVulnerability"`
}


func (val *PccRuleQosPolicyArgs) Defaults() *PccRuleQosPolicyArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AllocationAndRetentionPriorityLevel) {
		tmp.AllocationAndRetentionPriorityLevel = pulumi.IntPtr(9)
	}
	if isZero(tmp.FiveQi) {
		tmp.FiveQi = pulumi.IntPtr(9)
	}
	if isZero(tmp.PreemptionCapability) {
		tmp.PreemptionCapability = pulumi.StringPtr("NotPreempt")
	}
	if isZero(tmp.PreemptionVulnerability) {
		tmp.PreemptionVulnerability = pulumi.StringPtr("Preemptable")
	}
	return &tmp
}
func (PccRuleQosPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PccRuleQosPolicy)(nil)).Elem()
}

func (i PccRuleQosPolicyArgs) ToPccRuleQosPolicyOutput() PccRuleQosPolicyOutput {
	return i.ToPccRuleQosPolicyOutputWithContext(context.Background())
}

func (i PccRuleQosPolicyArgs) ToPccRuleQosPolicyOutputWithContext(ctx context.Context) PccRuleQosPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PccRuleQosPolicyOutput)
}

func (i PccRuleQosPolicyArgs) ToPccRuleQosPolicyPtrOutput() PccRuleQosPolicyPtrOutput {
	return i.ToPccRuleQosPolicyPtrOutputWithContext(context.Background())
}

func (i PccRuleQosPolicyArgs) ToPccRuleQosPolicyPtrOutputWithContext(ctx context.Context) PccRuleQosPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PccRuleQosPolicyOutput).ToPccRuleQosPolicyPtrOutputWithContext(ctx)
}









type PccRuleQosPolicyPtrInput interface {
	pulumi.Input

	ToPccRuleQosPolicyPtrOutput() PccRuleQosPolicyPtrOutput
	ToPccRuleQosPolicyPtrOutputWithContext(context.Context) PccRuleQosPolicyPtrOutput
}

type pccRuleQosPolicyPtrType PccRuleQosPolicyArgs

func PccRuleQosPolicyPtr(v *PccRuleQosPolicyArgs) PccRuleQosPolicyPtrInput {
	return (*pccRuleQosPolicyPtrType)(v)
}

func (*pccRuleQosPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PccRuleQosPolicy)(nil)).Elem()
}

func (i *pccRuleQosPolicyPtrType) ToPccRuleQosPolicyPtrOutput() PccRuleQosPolicyPtrOutput {
	return i.ToPccRuleQosPolicyPtrOutputWithContext(context.Background())
}

func (i *pccRuleQosPolicyPtrType) ToPccRuleQosPolicyPtrOutputWithContext(ctx context.Context) PccRuleQosPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PccRuleQosPolicyPtrOutput)
}

type PccRuleQosPolicyOutput struct{ *pulumi.OutputState }

func (PccRuleQosPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PccRuleQosPolicy)(nil)).Elem()
}

func (o PccRuleQosPolicyOutput) ToPccRuleQosPolicyOutput() PccRuleQosPolicyOutput {
	return o
}

func (o PccRuleQosPolicyOutput) ToPccRuleQosPolicyOutputWithContext(ctx context.Context) PccRuleQosPolicyOutput {
	return o
}

func (o PccRuleQosPolicyOutput) ToPccRuleQosPolicyPtrOutput() PccRuleQosPolicyPtrOutput {
	return o.ToPccRuleQosPolicyPtrOutputWithContext(context.Background())
}

func (o PccRuleQosPolicyOutput) ToPccRuleQosPolicyPtrOutputWithContext(ctx context.Context) PccRuleQosPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PccRuleQosPolicy) *PccRuleQosPolicy {
		return &v
	}).(PccRuleQosPolicyPtrOutput)
}

func (o PccRuleQosPolicyOutput) AllocationAndRetentionPriorityLevel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PccRuleQosPolicy) *int { return v.AllocationAndRetentionPriorityLevel }).(pulumi.IntPtrOutput)
}

func (o PccRuleQosPolicyOutput) FiveQi() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PccRuleQosPolicy) *int { return v.FiveQi }).(pulumi.IntPtrOutput)
}

func (o PccRuleQosPolicyOutput) GuaranteedBitRate() AmbrPtrOutput {
	return o.ApplyT(func(v PccRuleQosPolicy) *Ambr { return v.GuaranteedBitRate }).(AmbrPtrOutput)
}

func (o PccRuleQosPolicyOutput) MaximumBitRate() AmbrOutput {
	return o.ApplyT(func(v PccRuleQosPolicy) Ambr { return v.MaximumBitRate }).(AmbrOutput)
}

func (o PccRuleQosPolicyOutput) PreemptionCapability() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PccRuleQosPolicy) *string { return v.PreemptionCapability }).(pulumi.StringPtrOutput)
}

func (o PccRuleQosPolicyOutput) PreemptionVulnerability() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PccRuleQosPolicy) *string { return v.PreemptionVulnerability }).(pulumi.StringPtrOutput)
}

type PccRuleQosPolicyPtrOutput struct{ *pulumi.OutputState }

func (PccRuleQosPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PccRuleQosPolicy)(nil)).Elem()
}

func (o PccRuleQosPolicyPtrOutput) ToPccRuleQosPolicyPtrOutput() PccRuleQosPolicyPtrOutput {
	return o
}

func (o PccRuleQosPolicyPtrOutput) ToPccRuleQosPolicyPtrOutputWithContext(ctx context.Context) PccRuleQosPolicyPtrOutput {
	return o
}

func (o PccRuleQosPolicyPtrOutput) Elem() PccRuleQosPolicyOutput {
	return o.ApplyT(func(v *PccRuleQosPolicy) PccRuleQosPolicy {
		if v != nil {
			return *v
		}
		var ret PccRuleQosPolicy
		return ret
	}).(PccRuleQosPolicyOutput)
}

func (o PccRuleQosPolicyPtrOutput) AllocationAndRetentionPriorityLevel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PccRuleQosPolicy) *int {
		if v == nil {
			return nil
		}
		return v.AllocationAndRetentionPriorityLevel
	}).(pulumi.IntPtrOutput)
}

func (o PccRuleQosPolicyPtrOutput) FiveQi() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PccRuleQosPolicy) *int {
		if v == nil {
			return nil
		}
		return v.FiveQi
	}).(pulumi.IntPtrOutput)
}

func (o PccRuleQosPolicyPtrOutput) GuaranteedBitRate() AmbrPtrOutput {
	return o.ApplyT(func(v *PccRuleQosPolicy) *Ambr {
		if v == nil {
			return nil
		}
		return v.GuaranteedBitRate
	}).(AmbrPtrOutput)
}

func (o PccRuleQosPolicyPtrOutput) MaximumBitRate() AmbrPtrOutput {
	return o.ApplyT(func(v *PccRuleQosPolicy) *Ambr {
		if v == nil {
			return nil
		}
		return &v.MaximumBitRate
	}).(AmbrPtrOutput)
}

func (o PccRuleQosPolicyPtrOutput) PreemptionCapability() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PccRuleQosPolicy) *string {
		if v == nil {
			return nil
		}
		return v.PreemptionCapability
	}).(pulumi.StringPtrOutput)
}

func (o PccRuleQosPolicyPtrOutput) PreemptionVulnerability() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PccRuleQosPolicy) *string {
		if v == nil {
			return nil
		}
		return v.PreemptionVulnerability
	}).(pulumi.StringPtrOutput)
}

type PccRuleQosPolicyResponse struct {
	AllocationAndRetentionPriorityLevel *int          `pulumi:"allocationAndRetentionPriorityLevel"`
	FiveQi                              *int          `pulumi:"fiveQi"`
	GuaranteedBitRate                   *AmbrResponse `pulumi:"guaranteedBitRate"`
	MaximumBitRate                      AmbrResponse  `pulumi:"maximumBitRate"`
	PreemptionCapability                *string       `pulumi:"preemptionCapability"`
	PreemptionVulnerability             *string       `pulumi:"preemptionVulnerability"`
}


func (val *PccRuleQosPolicyResponse) Defaults() *PccRuleQosPolicyResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AllocationAndRetentionPriorityLevel) {
		allocationAndRetentionPriorityLevel_ := 9
		tmp.AllocationAndRetentionPriorityLevel = &allocationAndRetentionPriorityLevel_
	}
	if isZero(tmp.FiveQi) {
		fiveQi_ := 9
		tmp.FiveQi = &fiveQi_
	}
	if isZero(tmp.PreemptionCapability) {
		preemptionCapability_ := "NotPreempt"
		tmp.PreemptionCapability = &preemptionCapability_
	}
	if isZero(tmp.PreemptionVulnerability) {
		preemptionVulnerability_ := "Preemptable"
		tmp.PreemptionVulnerability = &preemptionVulnerability_
	}
	return &tmp
}

type PccRuleQosPolicyResponseOutput struct{ *pulumi.OutputState }

func (PccRuleQosPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PccRuleQosPolicyResponse)(nil)).Elem()
}

func (o PccRuleQosPolicyResponseOutput) ToPccRuleQosPolicyResponseOutput() PccRuleQosPolicyResponseOutput {
	return o
}

func (o PccRuleQosPolicyResponseOutput) ToPccRuleQosPolicyResponseOutputWithContext(ctx context.Context) PccRuleQosPolicyResponseOutput {
	return o
}

func (o PccRuleQosPolicyResponseOutput) AllocationAndRetentionPriorityLevel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PccRuleQosPolicyResponse) *int { return v.AllocationAndRetentionPriorityLevel }).(pulumi.IntPtrOutput)
}

func (o PccRuleQosPolicyResponseOutput) FiveQi() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PccRuleQosPolicyResponse) *int { return v.FiveQi }).(pulumi.IntPtrOutput)
}

func (o PccRuleQosPolicyResponseOutput) GuaranteedBitRate() AmbrResponsePtrOutput {
	return o.ApplyT(func(v PccRuleQosPolicyResponse) *AmbrResponse { return v.GuaranteedBitRate }).(AmbrResponsePtrOutput)
}

func (o PccRuleQosPolicyResponseOutput) MaximumBitRate() AmbrResponseOutput {
	return o.ApplyT(func(v PccRuleQosPolicyResponse) AmbrResponse { return v.MaximumBitRate }).(AmbrResponseOutput)
}

func (o PccRuleQosPolicyResponseOutput) PreemptionCapability() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PccRuleQosPolicyResponse) *string { return v.PreemptionCapability }).(pulumi.StringPtrOutput)
}

func (o PccRuleQosPolicyResponseOutput) PreemptionVulnerability() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PccRuleQosPolicyResponse) *string { return v.PreemptionVulnerability }).(pulumi.StringPtrOutput)
}

type PccRuleQosPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (PccRuleQosPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PccRuleQosPolicyResponse)(nil)).Elem()
}

func (o PccRuleQosPolicyResponsePtrOutput) ToPccRuleQosPolicyResponsePtrOutput() PccRuleQosPolicyResponsePtrOutput {
	return o
}

func (o PccRuleQosPolicyResponsePtrOutput) ToPccRuleQosPolicyResponsePtrOutputWithContext(ctx context.Context) PccRuleQosPolicyResponsePtrOutput {
	return o
}

func (o PccRuleQosPolicyResponsePtrOutput) Elem() PccRuleQosPolicyResponseOutput {
	return o.ApplyT(func(v *PccRuleQosPolicyResponse) PccRuleQosPolicyResponse {
		if v != nil {
			return *v
		}
		var ret PccRuleQosPolicyResponse
		return ret
	}).(PccRuleQosPolicyResponseOutput)
}

func (o PccRuleQosPolicyResponsePtrOutput) AllocationAndRetentionPriorityLevel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PccRuleQosPolicyResponse) *int {
		if v == nil {
			return nil
		}
		return v.AllocationAndRetentionPriorityLevel
	}).(pulumi.IntPtrOutput)
}

func (o PccRuleQosPolicyResponsePtrOutput) FiveQi() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PccRuleQosPolicyResponse) *int {
		if v == nil {
			return nil
		}
		return v.FiveQi
	}).(pulumi.IntPtrOutput)
}

func (o PccRuleQosPolicyResponsePtrOutput) GuaranteedBitRate() AmbrResponsePtrOutput {
	return o.ApplyT(func(v *PccRuleQosPolicyResponse) *AmbrResponse {
		if v == nil {
			return nil
		}
		return v.GuaranteedBitRate
	}).(AmbrResponsePtrOutput)
}

func (o PccRuleQosPolicyResponsePtrOutput) MaximumBitRate() AmbrResponsePtrOutput {
	return o.ApplyT(func(v *PccRuleQosPolicyResponse) *AmbrResponse {
		if v == nil {
			return nil
		}
		return &v.MaximumBitRate
	}).(AmbrResponsePtrOutput)
}

func (o PccRuleQosPolicyResponsePtrOutput) PreemptionCapability() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PccRuleQosPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.PreemptionCapability
	}).(pulumi.StringPtrOutput)
}

func (o PccRuleQosPolicyResponsePtrOutput) PreemptionVulnerability() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PccRuleQosPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.PreemptionVulnerability
	}).(pulumi.StringPtrOutput)
}

type PinholeTimeouts struct {
	Icmp *int `pulumi:"icmp"`
	Tcp  *int `pulumi:"tcp"`
	Udp  *int `pulumi:"udp"`
}


func (val *PinholeTimeouts) Defaults() *PinholeTimeouts {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Icmp) {
		icmp_ := 30
		tmp.Icmp = &icmp_
	}
	if isZero(tmp.Tcp) {
		tcp_ := 180
		tmp.Tcp = &tcp_
	}
	if isZero(tmp.Udp) {
		udp_ := 30
		tmp.Udp = &udp_
	}
	return &tmp
}





type PinholeTimeoutsInput interface {
	pulumi.Input

	ToPinholeTimeoutsOutput() PinholeTimeoutsOutput
	ToPinholeTimeoutsOutputWithContext(context.Context) PinholeTimeoutsOutput
}

type PinholeTimeoutsArgs struct {
	Icmp pulumi.IntPtrInput `pulumi:"icmp"`
	Tcp  pulumi.IntPtrInput `pulumi:"tcp"`
	Udp  pulumi.IntPtrInput `pulumi:"udp"`
}


func (val *PinholeTimeoutsArgs) Defaults() *PinholeTimeoutsArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Icmp) {
		tmp.Icmp = pulumi.IntPtr(30)
	}
	if isZero(tmp.Tcp) {
		tmp.Tcp = pulumi.IntPtr(180)
	}
	if isZero(tmp.Udp) {
		tmp.Udp = pulumi.IntPtr(30)
	}
	return &tmp
}
func (PinholeTimeoutsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PinholeTimeouts)(nil)).Elem()
}

func (i PinholeTimeoutsArgs) ToPinholeTimeoutsOutput() PinholeTimeoutsOutput {
	return i.ToPinholeTimeoutsOutputWithContext(context.Background())
}

func (i PinholeTimeoutsArgs) ToPinholeTimeoutsOutputWithContext(ctx context.Context) PinholeTimeoutsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PinholeTimeoutsOutput)
}

func (i PinholeTimeoutsArgs) ToPinholeTimeoutsPtrOutput() PinholeTimeoutsPtrOutput {
	return i.ToPinholeTimeoutsPtrOutputWithContext(context.Background())
}

func (i PinholeTimeoutsArgs) ToPinholeTimeoutsPtrOutputWithContext(ctx context.Context) PinholeTimeoutsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PinholeTimeoutsOutput).ToPinholeTimeoutsPtrOutputWithContext(ctx)
}









type PinholeTimeoutsPtrInput interface {
	pulumi.Input

	ToPinholeTimeoutsPtrOutput() PinholeTimeoutsPtrOutput
	ToPinholeTimeoutsPtrOutputWithContext(context.Context) PinholeTimeoutsPtrOutput
}

type pinholeTimeoutsPtrType PinholeTimeoutsArgs

func PinholeTimeoutsPtr(v *PinholeTimeoutsArgs) PinholeTimeoutsPtrInput {
	return (*pinholeTimeoutsPtrType)(v)
}

func (*pinholeTimeoutsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PinholeTimeouts)(nil)).Elem()
}

func (i *pinholeTimeoutsPtrType) ToPinholeTimeoutsPtrOutput() PinholeTimeoutsPtrOutput {
	return i.ToPinholeTimeoutsPtrOutputWithContext(context.Background())
}

func (i *pinholeTimeoutsPtrType) ToPinholeTimeoutsPtrOutputWithContext(ctx context.Context) PinholeTimeoutsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PinholeTimeoutsPtrOutput)
}

type PinholeTimeoutsOutput struct{ *pulumi.OutputState }

func (PinholeTimeoutsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PinholeTimeouts)(nil)).Elem()
}

func (o PinholeTimeoutsOutput) ToPinholeTimeoutsOutput() PinholeTimeoutsOutput {
	return o
}

func (o PinholeTimeoutsOutput) ToPinholeTimeoutsOutputWithContext(ctx context.Context) PinholeTimeoutsOutput {
	return o
}

func (o PinholeTimeoutsOutput) ToPinholeTimeoutsPtrOutput() PinholeTimeoutsPtrOutput {
	return o.ToPinholeTimeoutsPtrOutputWithContext(context.Background())
}

func (o PinholeTimeoutsOutput) ToPinholeTimeoutsPtrOutputWithContext(ctx context.Context) PinholeTimeoutsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PinholeTimeouts) *PinholeTimeouts {
		return &v
	}).(PinholeTimeoutsPtrOutput)
}

func (o PinholeTimeoutsOutput) Icmp() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PinholeTimeouts) *int { return v.Icmp }).(pulumi.IntPtrOutput)
}

func (o PinholeTimeoutsOutput) Tcp() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PinholeTimeouts) *int { return v.Tcp }).(pulumi.IntPtrOutput)
}

func (o PinholeTimeoutsOutput) Udp() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PinholeTimeouts) *int { return v.Udp }).(pulumi.IntPtrOutput)
}

type PinholeTimeoutsPtrOutput struct{ *pulumi.OutputState }

func (PinholeTimeoutsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PinholeTimeouts)(nil)).Elem()
}

func (o PinholeTimeoutsPtrOutput) ToPinholeTimeoutsPtrOutput() PinholeTimeoutsPtrOutput {
	return o
}

func (o PinholeTimeoutsPtrOutput) ToPinholeTimeoutsPtrOutputWithContext(ctx context.Context) PinholeTimeoutsPtrOutput {
	return o
}

func (o PinholeTimeoutsPtrOutput) Elem() PinholeTimeoutsOutput {
	return o.ApplyT(func(v *PinholeTimeouts) PinholeTimeouts {
		if v != nil {
			return *v
		}
		var ret PinholeTimeouts
		return ret
	}).(PinholeTimeoutsOutput)
}

func (o PinholeTimeoutsPtrOutput) Icmp() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PinholeTimeouts) *int {
		if v == nil {
			return nil
		}
		return v.Icmp
	}).(pulumi.IntPtrOutput)
}

func (o PinholeTimeoutsPtrOutput) Tcp() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PinholeTimeouts) *int {
		if v == nil {
			return nil
		}
		return v.Tcp
	}).(pulumi.IntPtrOutput)
}

func (o PinholeTimeoutsPtrOutput) Udp() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PinholeTimeouts) *int {
		if v == nil {
			return nil
		}
		return v.Udp
	}).(pulumi.IntPtrOutput)
}

type PinholeTimeoutsResponse struct {
	Icmp *int `pulumi:"icmp"`
	Tcp  *int `pulumi:"tcp"`
	Udp  *int `pulumi:"udp"`
}


func (val *PinholeTimeoutsResponse) Defaults() *PinholeTimeoutsResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Icmp) {
		icmp_ := 30
		tmp.Icmp = &icmp_
	}
	if isZero(tmp.Tcp) {
		tcp_ := 180
		tmp.Tcp = &tcp_
	}
	if isZero(tmp.Udp) {
		udp_ := 30
		tmp.Udp = &udp_
	}
	return &tmp
}

type PinholeTimeoutsResponseOutput struct{ *pulumi.OutputState }

func (PinholeTimeoutsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PinholeTimeoutsResponse)(nil)).Elem()
}

func (o PinholeTimeoutsResponseOutput) ToPinholeTimeoutsResponseOutput() PinholeTimeoutsResponseOutput {
	return o
}

func (o PinholeTimeoutsResponseOutput) ToPinholeTimeoutsResponseOutputWithContext(ctx context.Context) PinholeTimeoutsResponseOutput {
	return o
}

func (o PinholeTimeoutsResponseOutput) Icmp() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PinholeTimeoutsResponse) *int { return v.Icmp }).(pulumi.IntPtrOutput)
}

func (o PinholeTimeoutsResponseOutput) Tcp() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PinholeTimeoutsResponse) *int { return v.Tcp }).(pulumi.IntPtrOutput)
}

func (o PinholeTimeoutsResponseOutput) Udp() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PinholeTimeoutsResponse) *int { return v.Udp }).(pulumi.IntPtrOutput)
}

type PinholeTimeoutsResponsePtrOutput struct{ *pulumi.OutputState }

func (PinholeTimeoutsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PinholeTimeoutsResponse)(nil)).Elem()
}

func (o PinholeTimeoutsResponsePtrOutput) ToPinholeTimeoutsResponsePtrOutput() PinholeTimeoutsResponsePtrOutput {
	return o
}

func (o PinholeTimeoutsResponsePtrOutput) ToPinholeTimeoutsResponsePtrOutputWithContext(ctx context.Context) PinholeTimeoutsResponsePtrOutput {
	return o
}

func (o PinholeTimeoutsResponsePtrOutput) Elem() PinholeTimeoutsResponseOutput {
	return o.ApplyT(func(v *PinholeTimeoutsResponse) PinholeTimeoutsResponse {
		if v != nil {
			return *v
		}
		var ret PinholeTimeoutsResponse
		return ret
	}).(PinholeTimeoutsResponseOutput)
}

func (o PinholeTimeoutsResponsePtrOutput) Icmp() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PinholeTimeoutsResponse) *int {
		if v == nil {
			return nil
		}
		return v.Icmp
	}).(pulumi.IntPtrOutput)
}

func (o PinholeTimeoutsResponsePtrOutput) Tcp() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PinholeTimeoutsResponse) *int {
		if v == nil {
			return nil
		}
		return v.Tcp
	}).(pulumi.IntPtrOutput)
}

func (o PinholeTimeoutsResponsePtrOutput) Udp() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PinholeTimeoutsResponse) *int {
		if v == nil {
			return nil
		}
		return v.Udp
	}).(pulumi.IntPtrOutput)
}

type PlatformConfiguration struct {
	AzureStackEdgeDevice *AzureStackEdgeDeviceResourceId `pulumi:"azureStackEdgeDevice"`
	ConnectedCluster     *ConnectedClusterResourceId     `pulumi:"connectedCluster"`
	CustomLocation       *CustomLocationResourceId       `pulumi:"customLocation"`
	Type                 string                          `pulumi:"type"`
}





type PlatformConfigurationInput interface {
	pulumi.Input

	ToPlatformConfigurationOutput() PlatformConfigurationOutput
	ToPlatformConfigurationOutputWithContext(context.Context) PlatformConfigurationOutput
}

type PlatformConfigurationArgs struct {
	AzureStackEdgeDevice AzureStackEdgeDeviceResourceIdPtrInput `pulumi:"azureStackEdgeDevice"`
	ConnectedCluster     ConnectedClusterResourceIdPtrInput     `pulumi:"connectedCluster"`
	CustomLocation       CustomLocationResourceIdPtrInput       `pulumi:"customLocation"`
	Type                 pulumi.StringInput                     `pulumi:"type"`
}

func (PlatformConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PlatformConfiguration)(nil)).Elem()
}

func (i PlatformConfigurationArgs) ToPlatformConfigurationOutput() PlatformConfigurationOutput {
	return i.ToPlatformConfigurationOutputWithContext(context.Background())
}

func (i PlatformConfigurationArgs) ToPlatformConfigurationOutputWithContext(ctx context.Context) PlatformConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlatformConfigurationOutput)
}

func (i PlatformConfigurationArgs) ToPlatformConfigurationPtrOutput() PlatformConfigurationPtrOutput {
	return i.ToPlatformConfigurationPtrOutputWithContext(context.Background())
}

func (i PlatformConfigurationArgs) ToPlatformConfigurationPtrOutputWithContext(ctx context.Context) PlatformConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlatformConfigurationOutput).ToPlatformConfigurationPtrOutputWithContext(ctx)
}









type PlatformConfigurationPtrInput interface {
	pulumi.Input

	ToPlatformConfigurationPtrOutput() PlatformConfigurationPtrOutput
	ToPlatformConfigurationPtrOutputWithContext(context.Context) PlatformConfigurationPtrOutput
}

type platformConfigurationPtrType PlatformConfigurationArgs

func PlatformConfigurationPtr(v *PlatformConfigurationArgs) PlatformConfigurationPtrInput {
	return (*platformConfigurationPtrType)(v)
}

func (*platformConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PlatformConfiguration)(nil)).Elem()
}

func (i *platformConfigurationPtrType) ToPlatformConfigurationPtrOutput() PlatformConfigurationPtrOutput {
	return i.ToPlatformConfigurationPtrOutputWithContext(context.Background())
}

func (i *platformConfigurationPtrType) ToPlatformConfigurationPtrOutputWithContext(ctx context.Context) PlatformConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlatformConfigurationPtrOutput)
}

type PlatformConfigurationOutput struct{ *pulumi.OutputState }

func (PlatformConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PlatformConfiguration)(nil)).Elem()
}

func (o PlatformConfigurationOutput) ToPlatformConfigurationOutput() PlatformConfigurationOutput {
	return o
}

func (o PlatformConfigurationOutput) ToPlatformConfigurationOutputWithContext(ctx context.Context) PlatformConfigurationOutput {
	return o
}

func (o PlatformConfigurationOutput) ToPlatformConfigurationPtrOutput() PlatformConfigurationPtrOutput {
	return o.ToPlatformConfigurationPtrOutputWithContext(context.Background())
}

func (o PlatformConfigurationOutput) ToPlatformConfigurationPtrOutputWithContext(ctx context.Context) PlatformConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PlatformConfiguration) *PlatformConfiguration {
		return &v
	}).(PlatformConfigurationPtrOutput)
}

func (o PlatformConfigurationOutput) AzureStackEdgeDevice() AzureStackEdgeDeviceResourceIdPtrOutput {
	return o.ApplyT(func(v PlatformConfiguration) *AzureStackEdgeDeviceResourceId { return v.AzureStackEdgeDevice }).(AzureStackEdgeDeviceResourceIdPtrOutput)
}

func (o PlatformConfigurationOutput) ConnectedCluster() ConnectedClusterResourceIdPtrOutput {
	return o.ApplyT(func(v PlatformConfiguration) *ConnectedClusterResourceId { return v.ConnectedCluster }).(ConnectedClusterResourceIdPtrOutput)
}

func (o PlatformConfigurationOutput) CustomLocation() CustomLocationResourceIdPtrOutput {
	return o.ApplyT(func(v PlatformConfiguration) *CustomLocationResourceId { return v.CustomLocation }).(CustomLocationResourceIdPtrOutput)
}

func (o PlatformConfigurationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PlatformConfiguration) string { return v.Type }).(pulumi.StringOutput)
}

type PlatformConfigurationPtrOutput struct{ *pulumi.OutputState }

func (PlatformConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PlatformConfiguration)(nil)).Elem()
}

func (o PlatformConfigurationPtrOutput) ToPlatformConfigurationPtrOutput() PlatformConfigurationPtrOutput {
	return o
}

func (o PlatformConfigurationPtrOutput) ToPlatformConfigurationPtrOutputWithContext(ctx context.Context) PlatformConfigurationPtrOutput {
	return o
}

func (o PlatformConfigurationPtrOutput) Elem() PlatformConfigurationOutput {
	return o.ApplyT(func(v *PlatformConfiguration) PlatformConfiguration {
		if v != nil {
			return *v
		}
		var ret PlatformConfiguration
		return ret
	}).(PlatformConfigurationOutput)
}

func (o PlatformConfigurationPtrOutput) AzureStackEdgeDevice() AzureStackEdgeDeviceResourceIdPtrOutput {
	return o.ApplyT(func(v *PlatformConfiguration) *AzureStackEdgeDeviceResourceId {
		if v == nil {
			return nil
		}
		return v.AzureStackEdgeDevice
	}).(AzureStackEdgeDeviceResourceIdPtrOutput)
}

func (o PlatformConfigurationPtrOutput) ConnectedCluster() ConnectedClusterResourceIdPtrOutput {
	return o.ApplyT(func(v *PlatformConfiguration) *ConnectedClusterResourceId {
		if v == nil {
			return nil
		}
		return v.ConnectedCluster
	}).(ConnectedClusterResourceIdPtrOutput)
}

func (o PlatformConfigurationPtrOutput) CustomLocation() CustomLocationResourceIdPtrOutput {
	return o.ApplyT(func(v *PlatformConfiguration) *CustomLocationResourceId {
		if v == nil {
			return nil
		}
		return v.CustomLocation
	}).(CustomLocationResourceIdPtrOutput)
}

func (o PlatformConfigurationPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlatformConfiguration) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type PlatformConfigurationResponse struct {
	AzureStackEdgeDevice *AzureStackEdgeDeviceResourceIdResponse `pulumi:"azureStackEdgeDevice"`
	ConnectedCluster     *ConnectedClusterResourceIdResponse     `pulumi:"connectedCluster"`
	CustomLocation       *CustomLocationResourceIdResponse       `pulumi:"customLocation"`
	Type                 string                                  `pulumi:"type"`
}

type PlatformConfigurationResponseOutput struct{ *pulumi.OutputState }

func (PlatformConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PlatformConfigurationResponse)(nil)).Elem()
}

func (o PlatformConfigurationResponseOutput) ToPlatformConfigurationResponseOutput() PlatformConfigurationResponseOutput {
	return o
}

func (o PlatformConfigurationResponseOutput) ToPlatformConfigurationResponseOutputWithContext(ctx context.Context) PlatformConfigurationResponseOutput {
	return o
}

func (o PlatformConfigurationResponseOutput) AzureStackEdgeDevice() AzureStackEdgeDeviceResourceIdResponsePtrOutput {
	return o.ApplyT(func(v PlatformConfigurationResponse) *AzureStackEdgeDeviceResourceIdResponse {
		return v.AzureStackEdgeDevice
	}).(AzureStackEdgeDeviceResourceIdResponsePtrOutput)
}

func (o PlatformConfigurationResponseOutput) ConnectedCluster() ConnectedClusterResourceIdResponsePtrOutput {
	return o.ApplyT(func(v PlatformConfigurationResponse) *ConnectedClusterResourceIdResponse { return v.ConnectedCluster }).(ConnectedClusterResourceIdResponsePtrOutput)
}

func (o PlatformConfigurationResponseOutput) CustomLocation() CustomLocationResourceIdResponsePtrOutput {
	return o.ApplyT(func(v PlatformConfigurationResponse) *CustomLocationResourceIdResponse { return v.CustomLocation }).(CustomLocationResourceIdResponsePtrOutput)
}

func (o PlatformConfigurationResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PlatformConfigurationResponse) string { return v.Type }).(pulumi.StringOutput)
}

type PlatformConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (PlatformConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PlatformConfigurationResponse)(nil)).Elem()
}

func (o PlatformConfigurationResponsePtrOutput) ToPlatformConfigurationResponsePtrOutput() PlatformConfigurationResponsePtrOutput {
	return o
}

func (o PlatformConfigurationResponsePtrOutput) ToPlatformConfigurationResponsePtrOutputWithContext(ctx context.Context) PlatformConfigurationResponsePtrOutput {
	return o
}

func (o PlatformConfigurationResponsePtrOutput) Elem() PlatformConfigurationResponseOutput {
	return o.ApplyT(func(v *PlatformConfigurationResponse) PlatformConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret PlatformConfigurationResponse
		return ret
	}).(PlatformConfigurationResponseOutput)
}

func (o PlatformConfigurationResponsePtrOutput) AzureStackEdgeDevice() AzureStackEdgeDeviceResourceIdResponsePtrOutput {
	return o.ApplyT(func(v *PlatformConfigurationResponse) *AzureStackEdgeDeviceResourceIdResponse {
		if v == nil {
			return nil
		}
		return v.AzureStackEdgeDevice
	}).(AzureStackEdgeDeviceResourceIdResponsePtrOutput)
}

func (o PlatformConfigurationResponsePtrOutput) ConnectedCluster() ConnectedClusterResourceIdResponsePtrOutput {
	return o.ApplyT(func(v *PlatformConfigurationResponse) *ConnectedClusterResourceIdResponse {
		if v == nil {
			return nil
		}
		return v.ConnectedCluster
	}).(ConnectedClusterResourceIdResponsePtrOutput)
}

func (o PlatformConfigurationResponsePtrOutput) CustomLocation() CustomLocationResourceIdResponsePtrOutput {
	return o.ApplyT(func(v *PlatformConfigurationResponse) *CustomLocationResourceIdResponse {
		if v == nil {
			return nil
		}
		return v.CustomLocation
	}).(CustomLocationResourceIdResponsePtrOutput)
}

func (o PlatformConfigurationResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlatformConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type PlmnId struct {
	Mcc string `pulumi:"mcc"`
	Mnc string `pulumi:"mnc"`
}





type PlmnIdInput interface {
	pulumi.Input

	ToPlmnIdOutput() PlmnIdOutput
	ToPlmnIdOutputWithContext(context.Context) PlmnIdOutput
}

type PlmnIdArgs struct {
	Mcc pulumi.StringInput `pulumi:"mcc"`
	Mnc pulumi.StringInput `pulumi:"mnc"`
}

func (PlmnIdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PlmnId)(nil)).Elem()
}

func (i PlmnIdArgs) ToPlmnIdOutput() PlmnIdOutput {
	return i.ToPlmnIdOutputWithContext(context.Background())
}

func (i PlmnIdArgs) ToPlmnIdOutputWithContext(ctx context.Context) PlmnIdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlmnIdOutput)
}

type PlmnIdOutput struct{ *pulumi.OutputState }

func (PlmnIdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PlmnId)(nil)).Elem()
}

func (o PlmnIdOutput) ToPlmnIdOutput() PlmnIdOutput {
	return o
}

func (o PlmnIdOutput) ToPlmnIdOutputWithContext(ctx context.Context) PlmnIdOutput {
	return o
}

func (o PlmnIdOutput) Mcc() pulumi.StringOutput {
	return o.ApplyT(func(v PlmnId) string { return v.Mcc }).(pulumi.StringOutput)
}

func (o PlmnIdOutput) Mnc() pulumi.StringOutput {
	return o.ApplyT(func(v PlmnId) string { return v.Mnc }).(pulumi.StringOutput)
}

type PlmnIdResponse struct {
	Mcc string `pulumi:"mcc"`
	Mnc string `pulumi:"mnc"`
}

type PlmnIdResponseOutput struct{ *pulumi.OutputState }

func (PlmnIdResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PlmnIdResponse)(nil)).Elem()
}

func (o PlmnIdResponseOutput) ToPlmnIdResponseOutput() PlmnIdResponseOutput {
	return o
}

func (o PlmnIdResponseOutput) ToPlmnIdResponseOutputWithContext(ctx context.Context) PlmnIdResponseOutput {
	return o
}

func (o PlmnIdResponseOutput) Mcc() pulumi.StringOutput {
	return o.ApplyT(func(v PlmnIdResponse) string { return v.Mcc }).(pulumi.StringOutput)
}

func (o PlmnIdResponseOutput) Mnc() pulumi.StringOutput {
	return o.ApplyT(func(v PlmnIdResponse) string { return v.Mnc }).(pulumi.StringOutput)
}

type PortRange struct {
	MaxPort *int `pulumi:"maxPort"`
	MinPort *int `pulumi:"minPort"`
}


func (val *PortRange) Defaults() *PortRange {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MaxPort) {
		maxPort_ := 49999
		tmp.MaxPort = &maxPort_
	}
	if isZero(tmp.MinPort) {
		minPort_ := 1024
		tmp.MinPort = &minPort_
	}
	return &tmp
}





type PortRangeInput interface {
	pulumi.Input

	ToPortRangeOutput() PortRangeOutput
	ToPortRangeOutputWithContext(context.Context) PortRangeOutput
}

type PortRangeArgs struct {
	MaxPort pulumi.IntPtrInput `pulumi:"maxPort"`
	MinPort pulumi.IntPtrInput `pulumi:"minPort"`
}


func (val *PortRangeArgs) Defaults() *PortRangeArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MaxPort) {
		tmp.MaxPort = pulumi.IntPtr(49999)
	}
	if isZero(tmp.MinPort) {
		tmp.MinPort = pulumi.IntPtr(1024)
	}
	return &tmp
}
func (PortRangeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PortRange)(nil)).Elem()
}

func (i PortRangeArgs) ToPortRangeOutput() PortRangeOutput {
	return i.ToPortRangeOutputWithContext(context.Background())
}

func (i PortRangeArgs) ToPortRangeOutputWithContext(ctx context.Context) PortRangeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortRangeOutput)
}

func (i PortRangeArgs) ToPortRangePtrOutput() PortRangePtrOutput {
	return i.ToPortRangePtrOutputWithContext(context.Background())
}

func (i PortRangeArgs) ToPortRangePtrOutputWithContext(ctx context.Context) PortRangePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortRangeOutput).ToPortRangePtrOutputWithContext(ctx)
}









type PortRangePtrInput interface {
	pulumi.Input

	ToPortRangePtrOutput() PortRangePtrOutput
	ToPortRangePtrOutputWithContext(context.Context) PortRangePtrOutput
}

type portRangePtrType PortRangeArgs

func PortRangePtr(v *PortRangeArgs) PortRangePtrInput {
	return (*portRangePtrType)(v)
}

func (*portRangePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PortRange)(nil)).Elem()
}

func (i *portRangePtrType) ToPortRangePtrOutput() PortRangePtrOutput {
	return i.ToPortRangePtrOutputWithContext(context.Background())
}

func (i *portRangePtrType) ToPortRangePtrOutputWithContext(ctx context.Context) PortRangePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortRangePtrOutput)
}

type PortRangeOutput struct{ *pulumi.OutputState }

func (PortRangeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PortRange)(nil)).Elem()
}

func (o PortRangeOutput) ToPortRangeOutput() PortRangeOutput {
	return o
}

func (o PortRangeOutput) ToPortRangeOutputWithContext(ctx context.Context) PortRangeOutput {
	return o
}

func (o PortRangeOutput) ToPortRangePtrOutput() PortRangePtrOutput {
	return o.ToPortRangePtrOutputWithContext(context.Background())
}

func (o PortRangeOutput) ToPortRangePtrOutputWithContext(ctx context.Context) PortRangePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PortRange) *PortRange {
		return &v
	}).(PortRangePtrOutput)
}

func (o PortRangeOutput) MaxPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PortRange) *int { return v.MaxPort }).(pulumi.IntPtrOutput)
}

func (o PortRangeOutput) MinPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PortRange) *int { return v.MinPort }).(pulumi.IntPtrOutput)
}

type PortRangePtrOutput struct{ *pulumi.OutputState }

func (PortRangePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PortRange)(nil)).Elem()
}

func (o PortRangePtrOutput) ToPortRangePtrOutput() PortRangePtrOutput {
	return o
}

func (o PortRangePtrOutput) ToPortRangePtrOutputWithContext(ctx context.Context) PortRangePtrOutput {
	return o
}

func (o PortRangePtrOutput) Elem() PortRangeOutput {
	return o.ApplyT(func(v *PortRange) PortRange {
		if v != nil {
			return *v
		}
		var ret PortRange
		return ret
	}).(PortRangeOutput)
}

func (o PortRangePtrOutput) MaxPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PortRange) *int {
		if v == nil {
			return nil
		}
		return v.MaxPort
	}).(pulumi.IntPtrOutput)
}

func (o PortRangePtrOutput) MinPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PortRange) *int {
		if v == nil {
			return nil
		}
		return v.MinPort
	}).(pulumi.IntPtrOutput)
}

type PortRangeResponse struct {
	MaxPort *int `pulumi:"maxPort"`
	MinPort *int `pulumi:"minPort"`
}


func (val *PortRangeResponse) Defaults() *PortRangeResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MaxPort) {
		maxPort_ := 49999
		tmp.MaxPort = &maxPort_
	}
	if isZero(tmp.MinPort) {
		minPort_ := 1024
		tmp.MinPort = &minPort_
	}
	return &tmp
}

type PortRangeResponseOutput struct{ *pulumi.OutputState }

func (PortRangeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PortRangeResponse)(nil)).Elem()
}

func (o PortRangeResponseOutput) ToPortRangeResponseOutput() PortRangeResponseOutput {
	return o
}

func (o PortRangeResponseOutput) ToPortRangeResponseOutputWithContext(ctx context.Context) PortRangeResponseOutput {
	return o
}

func (o PortRangeResponseOutput) MaxPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PortRangeResponse) *int { return v.MaxPort }).(pulumi.IntPtrOutput)
}

func (o PortRangeResponseOutput) MinPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PortRangeResponse) *int { return v.MinPort }).(pulumi.IntPtrOutput)
}

type PortRangeResponsePtrOutput struct{ *pulumi.OutputState }

func (PortRangeResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PortRangeResponse)(nil)).Elem()
}

func (o PortRangeResponsePtrOutput) ToPortRangeResponsePtrOutput() PortRangeResponsePtrOutput {
	return o
}

func (o PortRangeResponsePtrOutput) ToPortRangeResponsePtrOutputWithContext(ctx context.Context) PortRangeResponsePtrOutput {
	return o
}

func (o PortRangeResponsePtrOutput) Elem() PortRangeResponseOutput {
	return o.ApplyT(func(v *PortRangeResponse) PortRangeResponse {
		if v != nil {
			return *v
		}
		var ret PortRangeResponse
		return ret
	}).(PortRangeResponseOutput)
}

func (o PortRangeResponsePtrOutput) MaxPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PortRangeResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxPort
	}).(pulumi.IntPtrOutput)
}

func (o PortRangeResponsePtrOutput) MinPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PortRangeResponse) *int {
		if v == nil {
			return nil
		}
		return v.MinPort
	}).(pulumi.IntPtrOutput)
}

type PortReuseHoldTimes struct {
	Tcp *int `pulumi:"tcp"`
	Udp *int `pulumi:"udp"`
}


func (val *PortReuseHoldTimes) Defaults() *PortReuseHoldTimes {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Tcp) {
		tcp_ := 120
		tmp.Tcp = &tcp_
	}
	if isZero(tmp.Udp) {
		udp_ := 60
		tmp.Udp = &udp_
	}
	return &tmp
}





type PortReuseHoldTimesInput interface {
	pulumi.Input

	ToPortReuseHoldTimesOutput() PortReuseHoldTimesOutput
	ToPortReuseHoldTimesOutputWithContext(context.Context) PortReuseHoldTimesOutput
}

type PortReuseHoldTimesArgs struct {
	Tcp pulumi.IntPtrInput `pulumi:"tcp"`
	Udp pulumi.IntPtrInput `pulumi:"udp"`
}


func (val *PortReuseHoldTimesArgs) Defaults() *PortReuseHoldTimesArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Tcp) {
		tmp.Tcp = pulumi.IntPtr(120)
	}
	if isZero(tmp.Udp) {
		tmp.Udp = pulumi.IntPtr(60)
	}
	return &tmp
}
func (PortReuseHoldTimesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PortReuseHoldTimes)(nil)).Elem()
}

func (i PortReuseHoldTimesArgs) ToPortReuseHoldTimesOutput() PortReuseHoldTimesOutput {
	return i.ToPortReuseHoldTimesOutputWithContext(context.Background())
}

func (i PortReuseHoldTimesArgs) ToPortReuseHoldTimesOutputWithContext(ctx context.Context) PortReuseHoldTimesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortReuseHoldTimesOutput)
}

func (i PortReuseHoldTimesArgs) ToPortReuseHoldTimesPtrOutput() PortReuseHoldTimesPtrOutput {
	return i.ToPortReuseHoldTimesPtrOutputWithContext(context.Background())
}

func (i PortReuseHoldTimesArgs) ToPortReuseHoldTimesPtrOutputWithContext(ctx context.Context) PortReuseHoldTimesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortReuseHoldTimesOutput).ToPortReuseHoldTimesPtrOutputWithContext(ctx)
}









type PortReuseHoldTimesPtrInput interface {
	pulumi.Input

	ToPortReuseHoldTimesPtrOutput() PortReuseHoldTimesPtrOutput
	ToPortReuseHoldTimesPtrOutputWithContext(context.Context) PortReuseHoldTimesPtrOutput
}

type portReuseHoldTimesPtrType PortReuseHoldTimesArgs

func PortReuseHoldTimesPtr(v *PortReuseHoldTimesArgs) PortReuseHoldTimesPtrInput {
	return (*portReuseHoldTimesPtrType)(v)
}

func (*portReuseHoldTimesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PortReuseHoldTimes)(nil)).Elem()
}

func (i *portReuseHoldTimesPtrType) ToPortReuseHoldTimesPtrOutput() PortReuseHoldTimesPtrOutput {
	return i.ToPortReuseHoldTimesPtrOutputWithContext(context.Background())
}

func (i *portReuseHoldTimesPtrType) ToPortReuseHoldTimesPtrOutputWithContext(ctx context.Context) PortReuseHoldTimesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortReuseHoldTimesPtrOutput)
}

type PortReuseHoldTimesOutput struct{ *pulumi.OutputState }

func (PortReuseHoldTimesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PortReuseHoldTimes)(nil)).Elem()
}

func (o PortReuseHoldTimesOutput) ToPortReuseHoldTimesOutput() PortReuseHoldTimesOutput {
	return o
}

func (o PortReuseHoldTimesOutput) ToPortReuseHoldTimesOutputWithContext(ctx context.Context) PortReuseHoldTimesOutput {
	return o
}

func (o PortReuseHoldTimesOutput) ToPortReuseHoldTimesPtrOutput() PortReuseHoldTimesPtrOutput {
	return o.ToPortReuseHoldTimesPtrOutputWithContext(context.Background())
}

func (o PortReuseHoldTimesOutput) ToPortReuseHoldTimesPtrOutputWithContext(ctx context.Context) PortReuseHoldTimesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PortReuseHoldTimes) *PortReuseHoldTimes {
		return &v
	}).(PortReuseHoldTimesPtrOutput)
}

func (o PortReuseHoldTimesOutput) Tcp() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PortReuseHoldTimes) *int { return v.Tcp }).(pulumi.IntPtrOutput)
}

func (o PortReuseHoldTimesOutput) Udp() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PortReuseHoldTimes) *int { return v.Udp }).(pulumi.IntPtrOutput)
}

type PortReuseHoldTimesPtrOutput struct{ *pulumi.OutputState }

func (PortReuseHoldTimesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PortReuseHoldTimes)(nil)).Elem()
}

func (o PortReuseHoldTimesPtrOutput) ToPortReuseHoldTimesPtrOutput() PortReuseHoldTimesPtrOutput {
	return o
}

func (o PortReuseHoldTimesPtrOutput) ToPortReuseHoldTimesPtrOutputWithContext(ctx context.Context) PortReuseHoldTimesPtrOutput {
	return o
}

func (o PortReuseHoldTimesPtrOutput) Elem() PortReuseHoldTimesOutput {
	return o.ApplyT(func(v *PortReuseHoldTimes) PortReuseHoldTimes {
		if v != nil {
			return *v
		}
		var ret PortReuseHoldTimes
		return ret
	}).(PortReuseHoldTimesOutput)
}

func (o PortReuseHoldTimesPtrOutput) Tcp() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PortReuseHoldTimes) *int {
		if v == nil {
			return nil
		}
		return v.Tcp
	}).(pulumi.IntPtrOutput)
}

func (o PortReuseHoldTimesPtrOutput) Udp() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PortReuseHoldTimes) *int {
		if v == nil {
			return nil
		}
		return v.Udp
	}).(pulumi.IntPtrOutput)
}

type PortReuseHoldTimesResponse struct {
	Tcp *int `pulumi:"tcp"`
	Udp *int `pulumi:"udp"`
}


func (val *PortReuseHoldTimesResponse) Defaults() *PortReuseHoldTimesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Tcp) {
		tcp_ := 120
		tmp.Tcp = &tcp_
	}
	if isZero(tmp.Udp) {
		udp_ := 60
		tmp.Udp = &udp_
	}
	return &tmp
}

type PortReuseHoldTimesResponseOutput struct{ *pulumi.OutputState }

func (PortReuseHoldTimesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PortReuseHoldTimesResponse)(nil)).Elem()
}

func (o PortReuseHoldTimesResponseOutput) ToPortReuseHoldTimesResponseOutput() PortReuseHoldTimesResponseOutput {
	return o
}

func (o PortReuseHoldTimesResponseOutput) ToPortReuseHoldTimesResponseOutputWithContext(ctx context.Context) PortReuseHoldTimesResponseOutput {
	return o
}

func (o PortReuseHoldTimesResponseOutput) Tcp() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PortReuseHoldTimesResponse) *int { return v.Tcp }).(pulumi.IntPtrOutput)
}

func (o PortReuseHoldTimesResponseOutput) Udp() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PortReuseHoldTimesResponse) *int { return v.Udp }).(pulumi.IntPtrOutput)
}

type PortReuseHoldTimesResponsePtrOutput struct{ *pulumi.OutputState }

func (PortReuseHoldTimesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PortReuseHoldTimesResponse)(nil)).Elem()
}

func (o PortReuseHoldTimesResponsePtrOutput) ToPortReuseHoldTimesResponsePtrOutput() PortReuseHoldTimesResponsePtrOutput {
	return o
}

func (o PortReuseHoldTimesResponsePtrOutput) ToPortReuseHoldTimesResponsePtrOutputWithContext(ctx context.Context) PortReuseHoldTimesResponsePtrOutput {
	return o
}

func (o PortReuseHoldTimesResponsePtrOutput) Elem() PortReuseHoldTimesResponseOutput {
	return o.ApplyT(func(v *PortReuseHoldTimesResponse) PortReuseHoldTimesResponse {
		if v != nil {
			return *v
		}
		var ret PortReuseHoldTimesResponse
		return ret
	}).(PortReuseHoldTimesResponseOutput)
}

func (o PortReuseHoldTimesResponsePtrOutput) Tcp() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PortReuseHoldTimesResponse) *int {
		if v == nil {
			return nil
		}
		return v.Tcp
	}).(pulumi.IntPtrOutput)
}

func (o PortReuseHoldTimesResponsePtrOutput) Udp() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PortReuseHoldTimesResponse) *int {
		if v == nil {
			return nil
		}
		return v.Udp
	}).(pulumi.IntPtrOutput)
}

type QosPolicy struct {
	AllocationAndRetentionPriorityLevel *int    `pulumi:"allocationAndRetentionPriorityLevel"`
	FiveQi                              *int    `pulumi:"fiveQi"`
	MaximumBitRate                      Ambr    `pulumi:"maximumBitRate"`
	PreemptionCapability                *string `pulumi:"preemptionCapability"`
	PreemptionVulnerability             *string `pulumi:"preemptionVulnerability"`
}


func (val *QosPolicy) Defaults() *QosPolicy {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AllocationAndRetentionPriorityLevel) {
		allocationAndRetentionPriorityLevel_ := 9
		tmp.AllocationAndRetentionPriorityLevel = &allocationAndRetentionPriorityLevel_
	}
	if isZero(tmp.FiveQi) {
		fiveQi_ := 9
		tmp.FiveQi = &fiveQi_
	}
	if isZero(tmp.PreemptionCapability) {
		preemptionCapability_ := "NotPreempt"
		tmp.PreemptionCapability = &preemptionCapability_
	}
	if isZero(tmp.PreemptionVulnerability) {
		preemptionVulnerability_ := "Preemptable"
		tmp.PreemptionVulnerability = &preemptionVulnerability_
	}
	return &tmp
}





type QosPolicyInput interface {
	pulumi.Input

	ToQosPolicyOutput() QosPolicyOutput
	ToQosPolicyOutputWithContext(context.Context) QosPolicyOutput
}

type QosPolicyArgs struct {
	AllocationAndRetentionPriorityLevel pulumi.IntPtrInput    `pulumi:"allocationAndRetentionPriorityLevel"`
	FiveQi                              pulumi.IntPtrInput    `pulumi:"fiveQi"`
	MaximumBitRate                      AmbrInput             `pulumi:"maximumBitRate"`
	PreemptionCapability                pulumi.StringPtrInput `pulumi:"preemptionCapability"`
	PreemptionVulnerability             pulumi.StringPtrInput `pulumi:"preemptionVulnerability"`
}


func (val *QosPolicyArgs) Defaults() *QosPolicyArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AllocationAndRetentionPriorityLevel) {
		tmp.AllocationAndRetentionPriorityLevel = pulumi.IntPtr(9)
	}
	if isZero(tmp.FiveQi) {
		tmp.FiveQi = pulumi.IntPtr(9)
	}
	if isZero(tmp.PreemptionCapability) {
		tmp.PreemptionCapability = pulumi.StringPtr("NotPreempt")
	}
	if isZero(tmp.PreemptionVulnerability) {
		tmp.PreemptionVulnerability = pulumi.StringPtr("Preemptable")
	}
	return &tmp
}
func (QosPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QosPolicy)(nil)).Elem()
}

func (i QosPolicyArgs) ToQosPolicyOutput() QosPolicyOutput {
	return i.ToQosPolicyOutputWithContext(context.Background())
}

func (i QosPolicyArgs) ToQosPolicyOutputWithContext(ctx context.Context) QosPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QosPolicyOutput)
}

func (i QosPolicyArgs) ToQosPolicyPtrOutput() QosPolicyPtrOutput {
	return i.ToQosPolicyPtrOutputWithContext(context.Background())
}

func (i QosPolicyArgs) ToQosPolicyPtrOutputWithContext(ctx context.Context) QosPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QosPolicyOutput).ToQosPolicyPtrOutputWithContext(ctx)
}









type QosPolicyPtrInput interface {
	pulumi.Input

	ToQosPolicyPtrOutput() QosPolicyPtrOutput
	ToQosPolicyPtrOutputWithContext(context.Context) QosPolicyPtrOutput
}

type qosPolicyPtrType QosPolicyArgs

func QosPolicyPtr(v *QosPolicyArgs) QosPolicyPtrInput {
	return (*qosPolicyPtrType)(v)
}

func (*qosPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**QosPolicy)(nil)).Elem()
}

func (i *qosPolicyPtrType) ToQosPolicyPtrOutput() QosPolicyPtrOutput {
	return i.ToQosPolicyPtrOutputWithContext(context.Background())
}

func (i *qosPolicyPtrType) ToQosPolicyPtrOutputWithContext(ctx context.Context) QosPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QosPolicyPtrOutput)
}

type QosPolicyOutput struct{ *pulumi.OutputState }

func (QosPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QosPolicy)(nil)).Elem()
}

func (o QosPolicyOutput) ToQosPolicyOutput() QosPolicyOutput {
	return o
}

func (o QosPolicyOutput) ToQosPolicyOutputWithContext(ctx context.Context) QosPolicyOutput {
	return o
}

func (o QosPolicyOutput) ToQosPolicyPtrOutput() QosPolicyPtrOutput {
	return o.ToQosPolicyPtrOutputWithContext(context.Background())
}

func (o QosPolicyOutput) ToQosPolicyPtrOutputWithContext(ctx context.Context) QosPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v QosPolicy) *QosPolicy {
		return &v
	}).(QosPolicyPtrOutput)
}

func (o QosPolicyOutput) AllocationAndRetentionPriorityLevel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v QosPolicy) *int { return v.AllocationAndRetentionPriorityLevel }).(pulumi.IntPtrOutput)
}

func (o QosPolicyOutput) FiveQi() pulumi.IntPtrOutput {
	return o.ApplyT(func(v QosPolicy) *int { return v.FiveQi }).(pulumi.IntPtrOutput)
}

func (o QosPolicyOutput) MaximumBitRate() AmbrOutput {
	return o.ApplyT(func(v QosPolicy) Ambr { return v.MaximumBitRate }).(AmbrOutput)
}

func (o QosPolicyOutput) PreemptionCapability() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QosPolicy) *string { return v.PreemptionCapability }).(pulumi.StringPtrOutput)
}

func (o QosPolicyOutput) PreemptionVulnerability() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QosPolicy) *string { return v.PreemptionVulnerability }).(pulumi.StringPtrOutput)
}

type QosPolicyPtrOutput struct{ *pulumi.OutputState }

func (QosPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QosPolicy)(nil)).Elem()
}

func (o QosPolicyPtrOutput) ToQosPolicyPtrOutput() QosPolicyPtrOutput {
	return o
}

func (o QosPolicyPtrOutput) ToQosPolicyPtrOutputWithContext(ctx context.Context) QosPolicyPtrOutput {
	return o
}

func (o QosPolicyPtrOutput) Elem() QosPolicyOutput {
	return o.ApplyT(func(v *QosPolicy) QosPolicy {
		if v != nil {
			return *v
		}
		var ret QosPolicy
		return ret
	}).(QosPolicyOutput)
}

func (o QosPolicyPtrOutput) AllocationAndRetentionPriorityLevel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *QosPolicy) *int {
		if v == nil {
			return nil
		}
		return v.AllocationAndRetentionPriorityLevel
	}).(pulumi.IntPtrOutput)
}

func (o QosPolicyPtrOutput) FiveQi() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *QosPolicy) *int {
		if v == nil {
			return nil
		}
		return v.FiveQi
	}).(pulumi.IntPtrOutput)
}

func (o QosPolicyPtrOutput) MaximumBitRate() AmbrPtrOutput {
	return o.ApplyT(func(v *QosPolicy) *Ambr {
		if v == nil {
			return nil
		}
		return &v.MaximumBitRate
	}).(AmbrPtrOutput)
}

func (o QosPolicyPtrOutput) PreemptionCapability() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QosPolicy) *string {
		if v == nil {
			return nil
		}
		return v.PreemptionCapability
	}).(pulumi.StringPtrOutput)
}

func (o QosPolicyPtrOutput) PreemptionVulnerability() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QosPolicy) *string {
		if v == nil {
			return nil
		}
		return v.PreemptionVulnerability
	}).(pulumi.StringPtrOutput)
}

type QosPolicyResponse struct {
	AllocationAndRetentionPriorityLevel *int         `pulumi:"allocationAndRetentionPriorityLevel"`
	FiveQi                              *int         `pulumi:"fiveQi"`
	MaximumBitRate                      AmbrResponse `pulumi:"maximumBitRate"`
	PreemptionCapability                *string      `pulumi:"preemptionCapability"`
	PreemptionVulnerability             *string      `pulumi:"preemptionVulnerability"`
}


func (val *QosPolicyResponse) Defaults() *QosPolicyResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AllocationAndRetentionPriorityLevel) {
		allocationAndRetentionPriorityLevel_ := 9
		tmp.AllocationAndRetentionPriorityLevel = &allocationAndRetentionPriorityLevel_
	}
	if isZero(tmp.FiveQi) {
		fiveQi_ := 9
		tmp.FiveQi = &fiveQi_
	}
	if isZero(tmp.PreemptionCapability) {
		preemptionCapability_ := "NotPreempt"
		tmp.PreemptionCapability = &preemptionCapability_
	}
	if isZero(tmp.PreemptionVulnerability) {
		preemptionVulnerability_ := "Preemptable"
		tmp.PreemptionVulnerability = &preemptionVulnerability_
	}
	return &tmp
}

type QosPolicyResponseOutput struct{ *pulumi.OutputState }

func (QosPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QosPolicyResponse)(nil)).Elem()
}

func (o QosPolicyResponseOutput) ToQosPolicyResponseOutput() QosPolicyResponseOutput {
	return o
}

func (o QosPolicyResponseOutput) ToQosPolicyResponseOutputWithContext(ctx context.Context) QosPolicyResponseOutput {
	return o
}

func (o QosPolicyResponseOutput) AllocationAndRetentionPriorityLevel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v QosPolicyResponse) *int { return v.AllocationAndRetentionPriorityLevel }).(pulumi.IntPtrOutput)
}

func (o QosPolicyResponseOutput) FiveQi() pulumi.IntPtrOutput {
	return o.ApplyT(func(v QosPolicyResponse) *int { return v.FiveQi }).(pulumi.IntPtrOutput)
}

func (o QosPolicyResponseOutput) MaximumBitRate() AmbrResponseOutput {
	return o.ApplyT(func(v QosPolicyResponse) AmbrResponse { return v.MaximumBitRate }).(AmbrResponseOutput)
}

func (o QosPolicyResponseOutput) PreemptionCapability() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QosPolicyResponse) *string { return v.PreemptionCapability }).(pulumi.StringPtrOutput)
}

func (o QosPolicyResponseOutput) PreemptionVulnerability() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QosPolicyResponse) *string { return v.PreemptionVulnerability }).(pulumi.StringPtrOutput)
}

type QosPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (QosPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QosPolicyResponse)(nil)).Elem()
}

func (o QosPolicyResponsePtrOutput) ToQosPolicyResponsePtrOutput() QosPolicyResponsePtrOutput {
	return o
}

func (o QosPolicyResponsePtrOutput) ToQosPolicyResponsePtrOutputWithContext(ctx context.Context) QosPolicyResponsePtrOutput {
	return o
}

func (o QosPolicyResponsePtrOutput) Elem() QosPolicyResponseOutput {
	return o.ApplyT(func(v *QosPolicyResponse) QosPolicyResponse {
		if v != nil {
			return *v
		}
		var ret QosPolicyResponse
		return ret
	}).(QosPolicyResponseOutput)
}

func (o QosPolicyResponsePtrOutput) AllocationAndRetentionPriorityLevel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *QosPolicyResponse) *int {
		if v == nil {
			return nil
		}
		return v.AllocationAndRetentionPriorityLevel
	}).(pulumi.IntPtrOutput)
}

func (o QosPolicyResponsePtrOutput) FiveQi() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *QosPolicyResponse) *int {
		if v == nil {
			return nil
		}
		return v.FiveQi
	}).(pulumi.IntPtrOutput)
}

func (o QosPolicyResponsePtrOutput) MaximumBitRate() AmbrResponsePtrOutput {
	return o.ApplyT(func(v *QosPolicyResponse) *AmbrResponse {
		if v == nil {
			return nil
		}
		return &v.MaximumBitRate
	}).(AmbrResponsePtrOutput)
}

func (o QosPolicyResponsePtrOutput) PreemptionCapability() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QosPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.PreemptionCapability
	}).(pulumi.StringPtrOutput)
}

func (o QosPolicyResponsePtrOutput) PreemptionVulnerability() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QosPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.PreemptionVulnerability
	}).(pulumi.StringPtrOutput)
}

type ServiceDataFlowTemplate struct {
	Direction    string   `pulumi:"direction"`
	Ports        []string `pulumi:"ports"`
	Protocol     []string `pulumi:"protocol"`
	RemoteIpList []string `pulumi:"remoteIpList"`
	TemplateName string   `pulumi:"templateName"`
}





type ServiceDataFlowTemplateInput interface {
	pulumi.Input

	ToServiceDataFlowTemplateOutput() ServiceDataFlowTemplateOutput
	ToServiceDataFlowTemplateOutputWithContext(context.Context) ServiceDataFlowTemplateOutput
}

type ServiceDataFlowTemplateArgs struct {
	Direction    pulumi.StringInput      `pulumi:"direction"`
	Ports        pulumi.StringArrayInput `pulumi:"ports"`
	Protocol     pulumi.StringArrayInput `pulumi:"protocol"`
	RemoteIpList pulumi.StringArrayInput `pulumi:"remoteIpList"`
	TemplateName pulumi.StringInput      `pulumi:"templateName"`
}

func (ServiceDataFlowTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceDataFlowTemplate)(nil)).Elem()
}

func (i ServiceDataFlowTemplateArgs) ToServiceDataFlowTemplateOutput() ServiceDataFlowTemplateOutput {
	return i.ToServiceDataFlowTemplateOutputWithContext(context.Background())
}

func (i ServiceDataFlowTemplateArgs) ToServiceDataFlowTemplateOutputWithContext(ctx context.Context) ServiceDataFlowTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceDataFlowTemplateOutput)
}





type ServiceDataFlowTemplateArrayInput interface {
	pulumi.Input

	ToServiceDataFlowTemplateArrayOutput() ServiceDataFlowTemplateArrayOutput
	ToServiceDataFlowTemplateArrayOutputWithContext(context.Context) ServiceDataFlowTemplateArrayOutput
}

type ServiceDataFlowTemplateArray []ServiceDataFlowTemplateInput

func (ServiceDataFlowTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceDataFlowTemplate)(nil)).Elem()
}

func (i ServiceDataFlowTemplateArray) ToServiceDataFlowTemplateArrayOutput() ServiceDataFlowTemplateArrayOutput {
	return i.ToServiceDataFlowTemplateArrayOutputWithContext(context.Background())
}

func (i ServiceDataFlowTemplateArray) ToServiceDataFlowTemplateArrayOutputWithContext(ctx context.Context) ServiceDataFlowTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceDataFlowTemplateArrayOutput)
}

type ServiceDataFlowTemplateOutput struct{ *pulumi.OutputState }

func (ServiceDataFlowTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceDataFlowTemplate)(nil)).Elem()
}

func (o ServiceDataFlowTemplateOutput) ToServiceDataFlowTemplateOutput() ServiceDataFlowTemplateOutput {
	return o
}

func (o ServiceDataFlowTemplateOutput) ToServiceDataFlowTemplateOutputWithContext(ctx context.Context) ServiceDataFlowTemplateOutput {
	return o
}

func (o ServiceDataFlowTemplateOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceDataFlowTemplate) string { return v.Direction }).(pulumi.StringOutput)
}

func (o ServiceDataFlowTemplateOutput) Ports() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServiceDataFlowTemplate) []string { return v.Ports }).(pulumi.StringArrayOutput)
}

func (o ServiceDataFlowTemplateOutput) Protocol() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServiceDataFlowTemplate) []string { return v.Protocol }).(pulumi.StringArrayOutput)
}

func (o ServiceDataFlowTemplateOutput) RemoteIpList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServiceDataFlowTemplate) []string { return v.RemoteIpList }).(pulumi.StringArrayOutput)
}

func (o ServiceDataFlowTemplateOutput) TemplateName() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceDataFlowTemplate) string { return v.TemplateName }).(pulumi.StringOutput)
}

type ServiceDataFlowTemplateArrayOutput struct{ *pulumi.OutputState }

func (ServiceDataFlowTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceDataFlowTemplate)(nil)).Elem()
}

func (o ServiceDataFlowTemplateArrayOutput) ToServiceDataFlowTemplateArrayOutput() ServiceDataFlowTemplateArrayOutput {
	return o
}

func (o ServiceDataFlowTemplateArrayOutput) ToServiceDataFlowTemplateArrayOutputWithContext(ctx context.Context) ServiceDataFlowTemplateArrayOutput {
	return o
}

func (o ServiceDataFlowTemplateArrayOutput) Index(i pulumi.IntInput) ServiceDataFlowTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceDataFlowTemplate {
		return vs[0].([]ServiceDataFlowTemplate)[vs[1].(int)]
	}).(ServiceDataFlowTemplateOutput)
}

type ServiceDataFlowTemplateResponse struct {
	Direction    string   `pulumi:"direction"`
	Ports        []string `pulumi:"ports"`
	Protocol     []string `pulumi:"protocol"`
	RemoteIpList []string `pulumi:"remoteIpList"`
	TemplateName string   `pulumi:"templateName"`
}

type ServiceDataFlowTemplateResponseOutput struct{ *pulumi.OutputState }

func (ServiceDataFlowTemplateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceDataFlowTemplateResponse)(nil)).Elem()
}

func (o ServiceDataFlowTemplateResponseOutput) ToServiceDataFlowTemplateResponseOutput() ServiceDataFlowTemplateResponseOutput {
	return o
}

func (o ServiceDataFlowTemplateResponseOutput) ToServiceDataFlowTemplateResponseOutputWithContext(ctx context.Context) ServiceDataFlowTemplateResponseOutput {
	return o
}

func (o ServiceDataFlowTemplateResponseOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceDataFlowTemplateResponse) string { return v.Direction }).(pulumi.StringOutput)
}

func (o ServiceDataFlowTemplateResponseOutput) Ports() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServiceDataFlowTemplateResponse) []string { return v.Ports }).(pulumi.StringArrayOutput)
}

func (o ServiceDataFlowTemplateResponseOutput) Protocol() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServiceDataFlowTemplateResponse) []string { return v.Protocol }).(pulumi.StringArrayOutput)
}

func (o ServiceDataFlowTemplateResponseOutput) RemoteIpList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServiceDataFlowTemplateResponse) []string { return v.RemoteIpList }).(pulumi.StringArrayOutput)
}

func (o ServiceDataFlowTemplateResponseOutput) TemplateName() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceDataFlowTemplateResponse) string { return v.TemplateName }).(pulumi.StringOutput)
}

type ServiceDataFlowTemplateResponseArrayOutput struct{ *pulumi.OutputState }

func (ServiceDataFlowTemplateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceDataFlowTemplateResponse)(nil)).Elem()
}

func (o ServiceDataFlowTemplateResponseArrayOutput) ToServiceDataFlowTemplateResponseArrayOutput() ServiceDataFlowTemplateResponseArrayOutput {
	return o
}

func (o ServiceDataFlowTemplateResponseArrayOutput) ToServiceDataFlowTemplateResponseArrayOutputWithContext(ctx context.Context) ServiceDataFlowTemplateResponseArrayOutput {
	return o
}

func (o ServiceDataFlowTemplateResponseArrayOutput) Index(i pulumi.IntInput) ServiceDataFlowTemplateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceDataFlowTemplateResponse {
		return vs[0].([]ServiceDataFlowTemplateResponse)[vs[1].(int)]
	}).(ServiceDataFlowTemplateResponseOutput)
}

type ServiceResourceId struct {
	Id string `pulumi:"id"`
}





type ServiceResourceIdInput interface {
	pulumi.Input

	ToServiceResourceIdOutput() ServiceResourceIdOutput
	ToServiceResourceIdOutputWithContext(context.Context) ServiceResourceIdOutput
}

type ServiceResourceIdArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (ServiceResourceIdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceResourceId)(nil)).Elem()
}

func (i ServiceResourceIdArgs) ToServiceResourceIdOutput() ServiceResourceIdOutput {
	return i.ToServiceResourceIdOutputWithContext(context.Background())
}

func (i ServiceResourceIdArgs) ToServiceResourceIdOutputWithContext(ctx context.Context) ServiceResourceIdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceResourceIdOutput)
}





type ServiceResourceIdArrayInput interface {
	pulumi.Input

	ToServiceResourceIdArrayOutput() ServiceResourceIdArrayOutput
	ToServiceResourceIdArrayOutputWithContext(context.Context) ServiceResourceIdArrayOutput
}

type ServiceResourceIdArray []ServiceResourceIdInput

func (ServiceResourceIdArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceResourceId)(nil)).Elem()
}

func (i ServiceResourceIdArray) ToServiceResourceIdArrayOutput() ServiceResourceIdArrayOutput {
	return i.ToServiceResourceIdArrayOutputWithContext(context.Background())
}

func (i ServiceResourceIdArray) ToServiceResourceIdArrayOutputWithContext(ctx context.Context) ServiceResourceIdArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceResourceIdArrayOutput)
}

type ServiceResourceIdOutput struct{ *pulumi.OutputState }

func (ServiceResourceIdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceResourceId)(nil)).Elem()
}

func (o ServiceResourceIdOutput) ToServiceResourceIdOutput() ServiceResourceIdOutput {
	return o
}

func (o ServiceResourceIdOutput) ToServiceResourceIdOutputWithContext(ctx context.Context) ServiceResourceIdOutput {
	return o
}

func (o ServiceResourceIdOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceResourceId) string { return v.Id }).(pulumi.StringOutput)
}

type ServiceResourceIdArrayOutput struct{ *pulumi.OutputState }

func (ServiceResourceIdArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceResourceId)(nil)).Elem()
}

func (o ServiceResourceIdArrayOutput) ToServiceResourceIdArrayOutput() ServiceResourceIdArrayOutput {
	return o
}

func (o ServiceResourceIdArrayOutput) ToServiceResourceIdArrayOutputWithContext(ctx context.Context) ServiceResourceIdArrayOutput {
	return o
}

func (o ServiceResourceIdArrayOutput) Index(i pulumi.IntInput) ServiceResourceIdOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceResourceId {
		return vs[0].([]ServiceResourceId)[vs[1].(int)]
	}).(ServiceResourceIdOutput)
}

type ServiceResourceIdResponse struct {
	Id string `pulumi:"id"`
}

type ServiceResourceIdResponseOutput struct{ *pulumi.OutputState }

func (ServiceResourceIdResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceResourceIdResponse)(nil)).Elem()
}

func (o ServiceResourceIdResponseOutput) ToServiceResourceIdResponseOutput() ServiceResourceIdResponseOutput {
	return o
}

func (o ServiceResourceIdResponseOutput) ToServiceResourceIdResponseOutputWithContext(ctx context.Context) ServiceResourceIdResponseOutput {
	return o
}

func (o ServiceResourceIdResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceResourceIdResponse) string { return v.Id }).(pulumi.StringOutput)
}

type ServiceResourceIdResponseArrayOutput struct{ *pulumi.OutputState }

func (ServiceResourceIdResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceResourceIdResponse)(nil)).Elem()
}

func (o ServiceResourceIdResponseArrayOutput) ToServiceResourceIdResponseArrayOutput() ServiceResourceIdResponseArrayOutput {
	return o
}

func (o ServiceResourceIdResponseArrayOutput) ToServiceResourceIdResponseArrayOutputWithContext(ctx context.Context) ServiceResourceIdResponseArrayOutput {
	return o
}

func (o ServiceResourceIdResponseArrayOutput) Index(i pulumi.IntInput) ServiceResourceIdResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceResourceIdResponse {
		return vs[0].([]ServiceResourceIdResponse)[vs[1].(int)]
	}).(ServiceResourceIdResponseOutput)
}

type SimPolicyResourceId struct {
	Id string `pulumi:"id"`
}





type SimPolicyResourceIdInput interface {
	pulumi.Input

	ToSimPolicyResourceIdOutput() SimPolicyResourceIdOutput
	ToSimPolicyResourceIdOutputWithContext(context.Context) SimPolicyResourceIdOutput
}

type SimPolicyResourceIdArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (SimPolicyResourceIdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SimPolicyResourceId)(nil)).Elem()
}

func (i SimPolicyResourceIdArgs) ToSimPolicyResourceIdOutput() SimPolicyResourceIdOutput {
	return i.ToSimPolicyResourceIdOutputWithContext(context.Background())
}

func (i SimPolicyResourceIdArgs) ToSimPolicyResourceIdOutputWithContext(ctx context.Context) SimPolicyResourceIdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SimPolicyResourceIdOutput)
}

func (i SimPolicyResourceIdArgs) ToSimPolicyResourceIdPtrOutput() SimPolicyResourceIdPtrOutput {
	return i.ToSimPolicyResourceIdPtrOutputWithContext(context.Background())
}

func (i SimPolicyResourceIdArgs) ToSimPolicyResourceIdPtrOutputWithContext(ctx context.Context) SimPolicyResourceIdPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SimPolicyResourceIdOutput).ToSimPolicyResourceIdPtrOutputWithContext(ctx)
}









type SimPolicyResourceIdPtrInput interface {
	pulumi.Input

	ToSimPolicyResourceIdPtrOutput() SimPolicyResourceIdPtrOutput
	ToSimPolicyResourceIdPtrOutputWithContext(context.Context) SimPolicyResourceIdPtrOutput
}

type simPolicyResourceIdPtrType SimPolicyResourceIdArgs

func SimPolicyResourceIdPtr(v *SimPolicyResourceIdArgs) SimPolicyResourceIdPtrInput {
	return (*simPolicyResourceIdPtrType)(v)
}

func (*simPolicyResourceIdPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SimPolicyResourceId)(nil)).Elem()
}

func (i *simPolicyResourceIdPtrType) ToSimPolicyResourceIdPtrOutput() SimPolicyResourceIdPtrOutput {
	return i.ToSimPolicyResourceIdPtrOutputWithContext(context.Background())
}

func (i *simPolicyResourceIdPtrType) ToSimPolicyResourceIdPtrOutputWithContext(ctx context.Context) SimPolicyResourceIdPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SimPolicyResourceIdPtrOutput)
}

type SimPolicyResourceIdOutput struct{ *pulumi.OutputState }

func (SimPolicyResourceIdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SimPolicyResourceId)(nil)).Elem()
}

func (o SimPolicyResourceIdOutput) ToSimPolicyResourceIdOutput() SimPolicyResourceIdOutput {
	return o
}

func (o SimPolicyResourceIdOutput) ToSimPolicyResourceIdOutputWithContext(ctx context.Context) SimPolicyResourceIdOutput {
	return o
}

func (o SimPolicyResourceIdOutput) ToSimPolicyResourceIdPtrOutput() SimPolicyResourceIdPtrOutput {
	return o.ToSimPolicyResourceIdPtrOutputWithContext(context.Background())
}

func (o SimPolicyResourceIdOutput) ToSimPolicyResourceIdPtrOutputWithContext(ctx context.Context) SimPolicyResourceIdPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SimPolicyResourceId) *SimPolicyResourceId {
		return &v
	}).(SimPolicyResourceIdPtrOutput)
}

func (o SimPolicyResourceIdOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SimPolicyResourceId) string { return v.Id }).(pulumi.StringOutput)
}

type SimPolicyResourceIdPtrOutput struct{ *pulumi.OutputState }

func (SimPolicyResourceIdPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SimPolicyResourceId)(nil)).Elem()
}

func (o SimPolicyResourceIdPtrOutput) ToSimPolicyResourceIdPtrOutput() SimPolicyResourceIdPtrOutput {
	return o
}

func (o SimPolicyResourceIdPtrOutput) ToSimPolicyResourceIdPtrOutputWithContext(ctx context.Context) SimPolicyResourceIdPtrOutput {
	return o
}

func (o SimPolicyResourceIdPtrOutput) Elem() SimPolicyResourceIdOutput {
	return o.ApplyT(func(v *SimPolicyResourceId) SimPolicyResourceId {
		if v != nil {
			return *v
		}
		var ret SimPolicyResourceId
		return ret
	}).(SimPolicyResourceIdOutput)
}

func (o SimPolicyResourceIdPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SimPolicyResourceId) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type SimPolicyResourceIdResponse struct {
	Id string `pulumi:"id"`
}

type SimPolicyResourceIdResponseOutput struct{ *pulumi.OutputState }

func (SimPolicyResourceIdResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SimPolicyResourceIdResponse)(nil)).Elem()
}

func (o SimPolicyResourceIdResponseOutput) ToSimPolicyResourceIdResponseOutput() SimPolicyResourceIdResponseOutput {
	return o
}

func (o SimPolicyResourceIdResponseOutput) ToSimPolicyResourceIdResponseOutputWithContext(ctx context.Context) SimPolicyResourceIdResponseOutput {
	return o
}

func (o SimPolicyResourceIdResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SimPolicyResourceIdResponse) string { return v.Id }).(pulumi.StringOutput)
}

type SimPolicyResourceIdResponsePtrOutput struct{ *pulumi.OutputState }

func (SimPolicyResourceIdResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SimPolicyResourceIdResponse)(nil)).Elem()
}

func (o SimPolicyResourceIdResponsePtrOutput) ToSimPolicyResourceIdResponsePtrOutput() SimPolicyResourceIdResponsePtrOutput {
	return o
}

func (o SimPolicyResourceIdResponsePtrOutput) ToSimPolicyResourceIdResponsePtrOutputWithContext(ctx context.Context) SimPolicyResourceIdResponsePtrOutput {
	return o
}

func (o SimPolicyResourceIdResponsePtrOutput) Elem() SimPolicyResourceIdResponseOutput {
	return o.ApplyT(func(v *SimPolicyResourceIdResponse) SimPolicyResourceIdResponse {
		if v != nil {
			return *v
		}
		var ret SimPolicyResourceIdResponse
		return ret
	}).(SimPolicyResourceIdResponseOutput)
}

func (o SimPolicyResourceIdResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SimPolicyResourceIdResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type SimStaticIpProperties struct {
	AttachedDataNetwork *AttachedDataNetworkResourceId `pulumi:"attachedDataNetwork"`
	Slice               *SliceResourceId               `pulumi:"slice"`
	StaticIp            *SimStaticIpPropertiesStaticIp `pulumi:"staticIp"`
}





type SimStaticIpPropertiesInput interface {
	pulumi.Input

	ToSimStaticIpPropertiesOutput() SimStaticIpPropertiesOutput
	ToSimStaticIpPropertiesOutputWithContext(context.Context) SimStaticIpPropertiesOutput
}

type SimStaticIpPropertiesArgs struct {
	AttachedDataNetwork AttachedDataNetworkResourceIdPtrInput `pulumi:"attachedDataNetwork"`
	Slice               SliceResourceIdPtrInput               `pulumi:"slice"`
	StaticIp            SimStaticIpPropertiesStaticIpPtrInput `pulumi:"staticIp"`
}

func (SimStaticIpPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SimStaticIpProperties)(nil)).Elem()
}

func (i SimStaticIpPropertiesArgs) ToSimStaticIpPropertiesOutput() SimStaticIpPropertiesOutput {
	return i.ToSimStaticIpPropertiesOutputWithContext(context.Background())
}

func (i SimStaticIpPropertiesArgs) ToSimStaticIpPropertiesOutputWithContext(ctx context.Context) SimStaticIpPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SimStaticIpPropertiesOutput)
}





type SimStaticIpPropertiesArrayInput interface {
	pulumi.Input

	ToSimStaticIpPropertiesArrayOutput() SimStaticIpPropertiesArrayOutput
	ToSimStaticIpPropertiesArrayOutputWithContext(context.Context) SimStaticIpPropertiesArrayOutput
}

type SimStaticIpPropertiesArray []SimStaticIpPropertiesInput

func (SimStaticIpPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SimStaticIpProperties)(nil)).Elem()
}

func (i SimStaticIpPropertiesArray) ToSimStaticIpPropertiesArrayOutput() SimStaticIpPropertiesArrayOutput {
	return i.ToSimStaticIpPropertiesArrayOutputWithContext(context.Background())
}

func (i SimStaticIpPropertiesArray) ToSimStaticIpPropertiesArrayOutputWithContext(ctx context.Context) SimStaticIpPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SimStaticIpPropertiesArrayOutput)
}

type SimStaticIpPropertiesOutput struct{ *pulumi.OutputState }

func (SimStaticIpPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SimStaticIpProperties)(nil)).Elem()
}

func (o SimStaticIpPropertiesOutput) ToSimStaticIpPropertiesOutput() SimStaticIpPropertiesOutput {
	return o
}

func (o SimStaticIpPropertiesOutput) ToSimStaticIpPropertiesOutputWithContext(ctx context.Context) SimStaticIpPropertiesOutput {
	return o
}

func (o SimStaticIpPropertiesOutput) AttachedDataNetwork() AttachedDataNetworkResourceIdPtrOutput {
	return o.ApplyT(func(v SimStaticIpProperties) *AttachedDataNetworkResourceId { return v.AttachedDataNetwork }).(AttachedDataNetworkResourceIdPtrOutput)
}

func (o SimStaticIpPropertiesOutput) Slice() SliceResourceIdPtrOutput {
	return o.ApplyT(func(v SimStaticIpProperties) *SliceResourceId { return v.Slice }).(SliceResourceIdPtrOutput)
}

func (o SimStaticIpPropertiesOutput) StaticIp() SimStaticIpPropertiesStaticIpPtrOutput {
	return o.ApplyT(func(v SimStaticIpProperties) *SimStaticIpPropertiesStaticIp { return v.StaticIp }).(SimStaticIpPropertiesStaticIpPtrOutput)
}

type SimStaticIpPropertiesArrayOutput struct{ *pulumi.OutputState }

func (SimStaticIpPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SimStaticIpProperties)(nil)).Elem()
}

func (o SimStaticIpPropertiesArrayOutput) ToSimStaticIpPropertiesArrayOutput() SimStaticIpPropertiesArrayOutput {
	return o
}

func (o SimStaticIpPropertiesArrayOutput) ToSimStaticIpPropertiesArrayOutputWithContext(ctx context.Context) SimStaticIpPropertiesArrayOutput {
	return o
}

func (o SimStaticIpPropertiesArrayOutput) Index(i pulumi.IntInput) SimStaticIpPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SimStaticIpProperties {
		return vs[0].([]SimStaticIpProperties)[vs[1].(int)]
	}).(SimStaticIpPropertiesOutput)
}

type SimStaticIpPropertiesResponse struct {
	AttachedDataNetwork *AttachedDataNetworkResourceIdResponse `pulumi:"attachedDataNetwork"`
	Slice               *SliceResourceIdResponse               `pulumi:"slice"`
	StaticIp            *SimStaticIpPropertiesResponseStaticIp `pulumi:"staticIp"`
}

type SimStaticIpPropertiesResponseOutput struct{ *pulumi.OutputState }

func (SimStaticIpPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SimStaticIpPropertiesResponse)(nil)).Elem()
}

func (o SimStaticIpPropertiesResponseOutput) ToSimStaticIpPropertiesResponseOutput() SimStaticIpPropertiesResponseOutput {
	return o
}

func (o SimStaticIpPropertiesResponseOutput) ToSimStaticIpPropertiesResponseOutputWithContext(ctx context.Context) SimStaticIpPropertiesResponseOutput {
	return o
}

func (o SimStaticIpPropertiesResponseOutput) AttachedDataNetwork() AttachedDataNetworkResourceIdResponsePtrOutput {
	return o.ApplyT(func(v SimStaticIpPropertiesResponse) *AttachedDataNetworkResourceIdResponse {
		return v.AttachedDataNetwork
	}).(AttachedDataNetworkResourceIdResponsePtrOutput)
}

func (o SimStaticIpPropertiesResponseOutput) Slice() SliceResourceIdResponsePtrOutput {
	return o.ApplyT(func(v SimStaticIpPropertiesResponse) *SliceResourceIdResponse { return v.Slice }).(SliceResourceIdResponsePtrOutput)
}

func (o SimStaticIpPropertiesResponseOutput) StaticIp() SimStaticIpPropertiesResponseStaticIpPtrOutput {
	return o.ApplyT(func(v SimStaticIpPropertiesResponse) *SimStaticIpPropertiesResponseStaticIp { return v.StaticIp }).(SimStaticIpPropertiesResponseStaticIpPtrOutput)
}

type SimStaticIpPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (SimStaticIpPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SimStaticIpPropertiesResponse)(nil)).Elem()
}

func (o SimStaticIpPropertiesResponseArrayOutput) ToSimStaticIpPropertiesResponseArrayOutput() SimStaticIpPropertiesResponseArrayOutput {
	return o
}

func (o SimStaticIpPropertiesResponseArrayOutput) ToSimStaticIpPropertiesResponseArrayOutputWithContext(ctx context.Context) SimStaticIpPropertiesResponseArrayOutput {
	return o
}

func (o SimStaticIpPropertiesResponseArrayOutput) Index(i pulumi.IntInput) SimStaticIpPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SimStaticIpPropertiesResponse {
		return vs[0].([]SimStaticIpPropertiesResponse)[vs[1].(int)]
	}).(SimStaticIpPropertiesResponseOutput)
}

type SimStaticIpPropertiesResponseStaticIp struct {
	Ipv4Address *string `pulumi:"ipv4Address"`
}

type SimStaticIpPropertiesResponseStaticIpOutput struct{ *pulumi.OutputState }

func (SimStaticIpPropertiesResponseStaticIpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SimStaticIpPropertiesResponseStaticIp)(nil)).Elem()
}

func (o SimStaticIpPropertiesResponseStaticIpOutput) ToSimStaticIpPropertiesResponseStaticIpOutput() SimStaticIpPropertiesResponseStaticIpOutput {
	return o
}

func (o SimStaticIpPropertiesResponseStaticIpOutput) ToSimStaticIpPropertiesResponseStaticIpOutputWithContext(ctx context.Context) SimStaticIpPropertiesResponseStaticIpOutput {
	return o
}

func (o SimStaticIpPropertiesResponseStaticIpOutput) Ipv4Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SimStaticIpPropertiesResponseStaticIp) *string { return v.Ipv4Address }).(pulumi.StringPtrOutput)
}

type SimStaticIpPropertiesResponseStaticIpPtrOutput struct{ *pulumi.OutputState }

func (SimStaticIpPropertiesResponseStaticIpPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SimStaticIpPropertiesResponseStaticIp)(nil)).Elem()
}

func (o SimStaticIpPropertiesResponseStaticIpPtrOutput) ToSimStaticIpPropertiesResponseStaticIpPtrOutput() SimStaticIpPropertiesResponseStaticIpPtrOutput {
	return o
}

func (o SimStaticIpPropertiesResponseStaticIpPtrOutput) ToSimStaticIpPropertiesResponseStaticIpPtrOutputWithContext(ctx context.Context) SimStaticIpPropertiesResponseStaticIpPtrOutput {
	return o
}

func (o SimStaticIpPropertiesResponseStaticIpPtrOutput) Elem() SimStaticIpPropertiesResponseStaticIpOutput {
	return o.ApplyT(func(v *SimStaticIpPropertiesResponseStaticIp) SimStaticIpPropertiesResponseStaticIp {
		if v != nil {
			return *v
		}
		var ret SimStaticIpPropertiesResponseStaticIp
		return ret
	}).(SimStaticIpPropertiesResponseStaticIpOutput)
}

func (o SimStaticIpPropertiesResponseStaticIpPtrOutput) Ipv4Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SimStaticIpPropertiesResponseStaticIp) *string {
		if v == nil {
			return nil
		}
		return v.Ipv4Address
	}).(pulumi.StringPtrOutput)
}

type SimStaticIpPropertiesStaticIp struct {
	Ipv4Address *string `pulumi:"ipv4Address"`
}





type SimStaticIpPropertiesStaticIpInput interface {
	pulumi.Input

	ToSimStaticIpPropertiesStaticIpOutput() SimStaticIpPropertiesStaticIpOutput
	ToSimStaticIpPropertiesStaticIpOutputWithContext(context.Context) SimStaticIpPropertiesStaticIpOutput
}

type SimStaticIpPropertiesStaticIpArgs struct {
	Ipv4Address pulumi.StringPtrInput `pulumi:"ipv4Address"`
}

func (SimStaticIpPropertiesStaticIpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SimStaticIpPropertiesStaticIp)(nil)).Elem()
}

func (i SimStaticIpPropertiesStaticIpArgs) ToSimStaticIpPropertiesStaticIpOutput() SimStaticIpPropertiesStaticIpOutput {
	return i.ToSimStaticIpPropertiesStaticIpOutputWithContext(context.Background())
}

func (i SimStaticIpPropertiesStaticIpArgs) ToSimStaticIpPropertiesStaticIpOutputWithContext(ctx context.Context) SimStaticIpPropertiesStaticIpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SimStaticIpPropertiesStaticIpOutput)
}

func (i SimStaticIpPropertiesStaticIpArgs) ToSimStaticIpPropertiesStaticIpPtrOutput() SimStaticIpPropertiesStaticIpPtrOutput {
	return i.ToSimStaticIpPropertiesStaticIpPtrOutputWithContext(context.Background())
}

func (i SimStaticIpPropertiesStaticIpArgs) ToSimStaticIpPropertiesStaticIpPtrOutputWithContext(ctx context.Context) SimStaticIpPropertiesStaticIpPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SimStaticIpPropertiesStaticIpOutput).ToSimStaticIpPropertiesStaticIpPtrOutputWithContext(ctx)
}









type SimStaticIpPropertiesStaticIpPtrInput interface {
	pulumi.Input

	ToSimStaticIpPropertiesStaticIpPtrOutput() SimStaticIpPropertiesStaticIpPtrOutput
	ToSimStaticIpPropertiesStaticIpPtrOutputWithContext(context.Context) SimStaticIpPropertiesStaticIpPtrOutput
}

type simStaticIpPropertiesStaticIpPtrType SimStaticIpPropertiesStaticIpArgs

func SimStaticIpPropertiesStaticIpPtr(v *SimStaticIpPropertiesStaticIpArgs) SimStaticIpPropertiesStaticIpPtrInput {
	return (*simStaticIpPropertiesStaticIpPtrType)(v)
}

func (*simStaticIpPropertiesStaticIpPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SimStaticIpPropertiesStaticIp)(nil)).Elem()
}

func (i *simStaticIpPropertiesStaticIpPtrType) ToSimStaticIpPropertiesStaticIpPtrOutput() SimStaticIpPropertiesStaticIpPtrOutput {
	return i.ToSimStaticIpPropertiesStaticIpPtrOutputWithContext(context.Background())
}

func (i *simStaticIpPropertiesStaticIpPtrType) ToSimStaticIpPropertiesStaticIpPtrOutputWithContext(ctx context.Context) SimStaticIpPropertiesStaticIpPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SimStaticIpPropertiesStaticIpPtrOutput)
}

type SimStaticIpPropertiesStaticIpOutput struct{ *pulumi.OutputState }

func (SimStaticIpPropertiesStaticIpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SimStaticIpPropertiesStaticIp)(nil)).Elem()
}

func (o SimStaticIpPropertiesStaticIpOutput) ToSimStaticIpPropertiesStaticIpOutput() SimStaticIpPropertiesStaticIpOutput {
	return o
}

func (o SimStaticIpPropertiesStaticIpOutput) ToSimStaticIpPropertiesStaticIpOutputWithContext(ctx context.Context) SimStaticIpPropertiesStaticIpOutput {
	return o
}

func (o SimStaticIpPropertiesStaticIpOutput) ToSimStaticIpPropertiesStaticIpPtrOutput() SimStaticIpPropertiesStaticIpPtrOutput {
	return o.ToSimStaticIpPropertiesStaticIpPtrOutputWithContext(context.Background())
}

func (o SimStaticIpPropertiesStaticIpOutput) ToSimStaticIpPropertiesStaticIpPtrOutputWithContext(ctx context.Context) SimStaticIpPropertiesStaticIpPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SimStaticIpPropertiesStaticIp) *SimStaticIpPropertiesStaticIp {
		return &v
	}).(SimStaticIpPropertiesStaticIpPtrOutput)
}

func (o SimStaticIpPropertiesStaticIpOutput) Ipv4Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SimStaticIpPropertiesStaticIp) *string { return v.Ipv4Address }).(pulumi.StringPtrOutput)
}

type SimStaticIpPropertiesStaticIpPtrOutput struct{ *pulumi.OutputState }

func (SimStaticIpPropertiesStaticIpPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SimStaticIpPropertiesStaticIp)(nil)).Elem()
}

func (o SimStaticIpPropertiesStaticIpPtrOutput) ToSimStaticIpPropertiesStaticIpPtrOutput() SimStaticIpPropertiesStaticIpPtrOutput {
	return o
}

func (o SimStaticIpPropertiesStaticIpPtrOutput) ToSimStaticIpPropertiesStaticIpPtrOutputWithContext(ctx context.Context) SimStaticIpPropertiesStaticIpPtrOutput {
	return o
}

func (o SimStaticIpPropertiesStaticIpPtrOutput) Elem() SimStaticIpPropertiesStaticIpOutput {
	return o.ApplyT(func(v *SimStaticIpPropertiesStaticIp) SimStaticIpPropertiesStaticIp {
		if v != nil {
			return *v
		}
		var ret SimStaticIpPropertiesStaticIp
		return ret
	}).(SimStaticIpPropertiesStaticIpOutput)
}

func (o SimStaticIpPropertiesStaticIpPtrOutput) Ipv4Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SimStaticIpPropertiesStaticIp) *string {
		if v == nil {
			return nil
		}
		return v.Ipv4Address
	}).(pulumi.StringPtrOutput)
}

type SliceConfiguration struct {
	DataNetworkConfigurations []DataNetworkConfiguration `pulumi:"dataNetworkConfigurations"`
	DefaultDataNetwork        DataNetworkResourceId      `pulumi:"defaultDataNetwork"`
	Slice                     SliceResourceId            `pulumi:"slice"`
}





type SliceConfigurationInput interface {
	pulumi.Input

	ToSliceConfigurationOutput() SliceConfigurationOutput
	ToSliceConfigurationOutputWithContext(context.Context) SliceConfigurationOutput
}

type SliceConfigurationArgs struct {
	DataNetworkConfigurations DataNetworkConfigurationArrayInput `pulumi:"dataNetworkConfigurations"`
	DefaultDataNetwork        DataNetworkResourceIdInput         `pulumi:"defaultDataNetwork"`
	Slice                     SliceResourceIdInput               `pulumi:"slice"`
}

func (SliceConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SliceConfiguration)(nil)).Elem()
}

func (i SliceConfigurationArgs) ToSliceConfigurationOutput() SliceConfigurationOutput {
	return i.ToSliceConfigurationOutputWithContext(context.Background())
}

func (i SliceConfigurationArgs) ToSliceConfigurationOutputWithContext(ctx context.Context) SliceConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SliceConfigurationOutput)
}





type SliceConfigurationArrayInput interface {
	pulumi.Input

	ToSliceConfigurationArrayOutput() SliceConfigurationArrayOutput
	ToSliceConfigurationArrayOutputWithContext(context.Context) SliceConfigurationArrayOutput
}

type SliceConfigurationArray []SliceConfigurationInput

func (SliceConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SliceConfiguration)(nil)).Elem()
}

func (i SliceConfigurationArray) ToSliceConfigurationArrayOutput() SliceConfigurationArrayOutput {
	return i.ToSliceConfigurationArrayOutputWithContext(context.Background())
}

func (i SliceConfigurationArray) ToSliceConfigurationArrayOutputWithContext(ctx context.Context) SliceConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SliceConfigurationArrayOutput)
}

type SliceConfigurationOutput struct{ *pulumi.OutputState }

func (SliceConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SliceConfiguration)(nil)).Elem()
}

func (o SliceConfigurationOutput) ToSliceConfigurationOutput() SliceConfigurationOutput {
	return o
}

func (o SliceConfigurationOutput) ToSliceConfigurationOutputWithContext(ctx context.Context) SliceConfigurationOutput {
	return o
}

func (o SliceConfigurationOutput) DataNetworkConfigurations() DataNetworkConfigurationArrayOutput {
	return o.ApplyT(func(v SliceConfiguration) []DataNetworkConfiguration { return v.DataNetworkConfigurations }).(DataNetworkConfigurationArrayOutput)
}

func (o SliceConfigurationOutput) DefaultDataNetwork() DataNetworkResourceIdOutput {
	return o.ApplyT(func(v SliceConfiguration) DataNetworkResourceId { return v.DefaultDataNetwork }).(DataNetworkResourceIdOutput)
}

func (o SliceConfigurationOutput) Slice() SliceResourceIdOutput {
	return o.ApplyT(func(v SliceConfiguration) SliceResourceId { return v.Slice }).(SliceResourceIdOutput)
}

type SliceConfigurationArrayOutput struct{ *pulumi.OutputState }

func (SliceConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SliceConfiguration)(nil)).Elem()
}

func (o SliceConfigurationArrayOutput) ToSliceConfigurationArrayOutput() SliceConfigurationArrayOutput {
	return o
}

func (o SliceConfigurationArrayOutput) ToSliceConfigurationArrayOutputWithContext(ctx context.Context) SliceConfigurationArrayOutput {
	return o
}

func (o SliceConfigurationArrayOutput) Index(i pulumi.IntInput) SliceConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SliceConfiguration {
		return vs[0].([]SliceConfiguration)[vs[1].(int)]
	}).(SliceConfigurationOutput)
}

type SliceConfigurationResponse struct {
	DataNetworkConfigurations []DataNetworkConfigurationResponse `pulumi:"dataNetworkConfigurations"`
	DefaultDataNetwork        DataNetworkResourceIdResponse      `pulumi:"defaultDataNetwork"`
	Slice                     SliceResourceIdResponse            `pulumi:"slice"`
}

type SliceConfigurationResponseOutput struct{ *pulumi.OutputState }

func (SliceConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SliceConfigurationResponse)(nil)).Elem()
}

func (o SliceConfigurationResponseOutput) ToSliceConfigurationResponseOutput() SliceConfigurationResponseOutput {
	return o
}

func (o SliceConfigurationResponseOutput) ToSliceConfigurationResponseOutputWithContext(ctx context.Context) SliceConfigurationResponseOutput {
	return o
}

func (o SliceConfigurationResponseOutput) DataNetworkConfigurations() DataNetworkConfigurationResponseArrayOutput {
	return o.ApplyT(func(v SliceConfigurationResponse) []DataNetworkConfigurationResponse {
		return v.DataNetworkConfigurations
	}).(DataNetworkConfigurationResponseArrayOutput)
}

func (o SliceConfigurationResponseOutput) DefaultDataNetwork() DataNetworkResourceIdResponseOutput {
	return o.ApplyT(func(v SliceConfigurationResponse) DataNetworkResourceIdResponse { return v.DefaultDataNetwork }).(DataNetworkResourceIdResponseOutput)
}

func (o SliceConfigurationResponseOutput) Slice() SliceResourceIdResponseOutput {
	return o.ApplyT(func(v SliceConfigurationResponse) SliceResourceIdResponse { return v.Slice }).(SliceResourceIdResponseOutput)
}

type SliceConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (SliceConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SliceConfigurationResponse)(nil)).Elem()
}

func (o SliceConfigurationResponseArrayOutput) ToSliceConfigurationResponseArrayOutput() SliceConfigurationResponseArrayOutput {
	return o
}

func (o SliceConfigurationResponseArrayOutput) ToSliceConfigurationResponseArrayOutputWithContext(ctx context.Context) SliceConfigurationResponseArrayOutput {
	return o
}

func (o SliceConfigurationResponseArrayOutput) Index(i pulumi.IntInput) SliceConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SliceConfigurationResponse {
		return vs[0].([]SliceConfigurationResponse)[vs[1].(int)]
	}).(SliceConfigurationResponseOutput)
}

type SliceResourceId struct {
	Id string `pulumi:"id"`
}





type SliceResourceIdInput interface {
	pulumi.Input

	ToSliceResourceIdOutput() SliceResourceIdOutput
	ToSliceResourceIdOutputWithContext(context.Context) SliceResourceIdOutput
}

type SliceResourceIdArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (SliceResourceIdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SliceResourceId)(nil)).Elem()
}

func (i SliceResourceIdArgs) ToSliceResourceIdOutput() SliceResourceIdOutput {
	return i.ToSliceResourceIdOutputWithContext(context.Background())
}

func (i SliceResourceIdArgs) ToSliceResourceIdOutputWithContext(ctx context.Context) SliceResourceIdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SliceResourceIdOutput)
}

func (i SliceResourceIdArgs) ToSliceResourceIdPtrOutput() SliceResourceIdPtrOutput {
	return i.ToSliceResourceIdPtrOutputWithContext(context.Background())
}

func (i SliceResourceIdArgs) ToSliceResourceIdPtrOutputWithContext(ctx context.Context) SliceResourceIdPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SliceResourceIdOutput).ToSliceResourceIdPtrOutputWithContext(ctx)
}









type SliceResourceIdPtrInput interface {
	pulumi.Input

	ToSliceResourceIdPtrOutput() SliceResourceIdPtrOutput
	ToSliceResourceIdPtrOutputWithContext(context.Context) SliceResourceIdPtrOutput
}

type sliceResourceIdPtrType SliceResourceIdArgs

func SliceResourceIdPtr(v *SliceResourceIdArgs) SliceResourceIdPtrInput {
	return (*sliceResourceIdPtrType)(v)
}

func (*sliceResourceIdPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SliceResourceId)(nil)).Elem()
}

func (i *sliceResourceIdPtrType) ToSliceResourceIdPtrOutput() SliceResourceIdPtrOutput {
	return i.ToSliceResourceIdPtrOutputWithContext(context.Background())
}

func (i *sliceResourceIdPtrType) ToSliceResourceIdPtrOutputWithContext(ctx context.Context) SliceResourceIdPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SliceResourceIdPtrOutput)
}

type SliceResourceIdOutput struct{ *pulumi.OutputState }

func (SliceResourceIdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SliceResourceId)(nil)).Elem()
}

func (o SliceResourceIdOutput) ToSliceResourceIdOutput() SliceResourceIdOutput {
	return o
}

func (o SliceResourceIdOutput) ToSliceResourceIdOutputWithContext(ctx context.Context) SliceResourceIdOutput {
	return o
}

func (o SliceResourceIdOutput) ToSliceResourceIdPtrOutput() SliceResourceIdPtrOutput {
	return o.ToSliceResourceIdPtrOutputWithContext(context.Background())
}

func (o SliceResourceIdOutput) ToSliceResourceIdPtrOutputWithContext(ctx context.Context) SliceResourceIdPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SliceResourceId) *SliceResourceId {
		return &v
	}).(SliceResourceIdPtrOutput)
}

func (o SliceResourceIdOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SliceResourceId) string { return v.Id }).(pulumi.StringOutput)
}

type SliceResourceIdPtrOutput struct{ *pulumi.OutputState }

func (SliceResourceIdPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SliceResourceId)(nil)).Elem()
}

func (o SliceResourceIdPtrOutput) ToSliceResourceIdPtrOutput() SliceResourceIdPtrOutput {
	return o
}

func (o SliceResourceIdPtrOutput) ToSliceResourceIdPtrOutputWithContext(ctx context.Context) SliceResourceIdPtrOutput {
	return o
}

func (o SliceResourceIdPtrOutput) Elem() SliceResourceIdOutput {
	return o.ApplyT(func(v *SliceResourceId) SliceResourceId {
		if v != nil {
			return *v
		}
		var ret SliceResourceId
		return ret
	}).(SliceResourceIdOutput)
}

func (o SliceResourceIdPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SliceResourceId) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type SliceResourceIdResponse struct {
	Id string `pulumi:"id"`
}

type SliceResourceIdResponseOutput struct{ *pulumi.OutputState }

func (SliceResourceIdResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SliceResourceIdResponse)(nil)).Elem()
}

func (o SliceResourceIdResponseOutput) ToSliceResourceIdResponseOutput() SliceResourceIdResponseOutput {
	return o
}

func (o SliceResourceIdResponseOutput) ToSliceResourceIdResponseOutputWithContext(ctx context.Context) SliceResourceIdResponseOutput {
	return o
}

func (o SliceResourceIdResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SliceResourceIdResponse) string { return v.Id }).(pulumi.StringOutput)
}

type SliceResourceIdResponsePtrOutput struct{ *pulumi.OutputState }

func (SliceResourceIdResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SliceResourceIdResponse)(nil)).Elem()
}

func (o SliceResourceIdResponsePtrOutput) ToSliceResourceIdResponsePtrOutput() SliceResourceIdResponsePtrOutput {
	return o
}

func (o SliceResourceIdResponsePtrOutput) ToSliceResourceIdResponsePtrOutputWithContext(ctx context.Context) SliceResourceIdResponsePtrOutput {
	return o
}

func (o SliceResourceIdResponsePtrOutput) Elem() SliceResourceIdResponseOutput {
	return o.ApplyT(func(v *SliceResourceIdResponse) SliceResourceIdResponse {
		if v != nil {
			return *v
		}
		var ret SliceResourceIdResponse
		return ret
	}).(SliceResourceIdResponseOutput)
}

func (o SliceResourceIdResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SliceResourceIdResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type Snssai struct {
	Sd  *string `pulumi:"sd"`
	Sst int     `pulumi:"sst"`
}





type SnssaiInput interface {
	pulumi.Input

	ToSnssaiOutput() SnssaiOutput
	ToSnssaiOutputWithContext(context.Context) SnssaiOutput
}

type SnssaiArgs struct {
	Sd  pulumi.StringPtrInput `pulumi:"sd"`
	Sst pulumi.IntInput       `pulumi:"sst"`
}

func (SnssaiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Snssai)(nil)).Elem()
}

func (i SnssaiArgs) ToSnssaiOutput() SnssaiOutput {
	return i.ToSnssaiOutputWithContext(context.Background())
}

func (i SnssaiArgs) ToSnssaiOutputWithContext(ctx context.Context) SnssaiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnssaiOutput)
}

type SnssaiOutput struct{ *pulumi.OutputState }

func (SnssaiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Snssai)(nil)).Elem()
}

func (o SnssaiOutput) ToSnssaiOutput() SnssaiOutput {
	return o
}

func (o SnssaiOutput) ToSnssaiOutputWithContext(ctx context.Context) SnssaiOutput {
	return o
}

func (o SnssaiOutput) Sd() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Snssai) *string { return v.Sd }).(pulumi.StringPtrOutput)
}

func (o SnssaiOutput) Sst() pulumi.IntOutput {
	return o.ApplyT(func(v Snssai) int { return v.Sst }).(pulumi.IntOutput)
}

type SnssaiResponse struct {
	Sd  *string `pulumi:"sd"`
	Sst int     `pulumi:"sst"`
}

type SnssaiResponseOutput struct{ *pulumi.OutputState }

func (SnssaiResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SnssaiResponse)(nil)).Elem()
}

func (o SnssaiResponseOutput) ToSnssaiResponseOutput() SnssaiResponseOutput {
	return o
}

func (o SnssaiResponseOutput) ToSnssaiResponseOutputWithContext(ctx context.Context) SnssaiResponseOutput {
	return o
}

func (o SnssaiResponseOutput) Sd() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SnssaiResponse) *string { return v.Sd }).(pulumi.StringPtrOutput)
}

func (o SnssaiResponseOutput) Sst() pulumi.IntOutput {
	return o.ApplyT(func(v SnssaiResponse) int { return v.Sst }).(pulumi.IntOutput)
}

type SubResource struct {
	Id string `pulumi:"id"`
}





type SubResourceInput interface {
	pulumi.Input

	ToSubResourceOutput() SubResourceOutput
	ToSubResourceOutputWithContext(context.Context) SubResourceOutput
}

type SubResourceArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (SubResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResource)(nil)).Elem()
}

func (i SubResourceArgs) ToSubResourceOutput() SubResourceOutput {
	return i.ToSubResourceOutputWithContext(context.Background())
}

func (i SubResourceArgs) ToSubResourceOutputWithContext(ctx context.Context) SubResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceOutput)
}





type SubResourceArrayInput interface {
	pulumi.Input

	ToSubResourceArrayOutput() SubResourceArrayOutput
	ToSubResourceArrayOutputWithContext(context.Context) SubResourceArrayOutput
}

type SubResourceArray []SubResourceInput

func (SubResourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubResource)(nil)).Elem()
}

func (i SubResourceArray) ToSubResourceArrayOutput() SubResourceArrayOutput {
	return i.ToSubResourceArrayOutputWithContext(context.Background())
}

func (i SubResourceArray) ToSubResourceArrayOutputWithContext(ctx context.Context) SubResourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceArrayOutput)
}

type SubResourceOutput struct{ *pulumi.OutputState }

func (SubResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResource)(nil)).Elem()
}

func (o SubResourceOutput) ToSubResourceOutput() SubResourceOutput {
	return o
}

func (o SubResourceOutput) ToSubResourceOutputWithContext(ctx context.Context) SubResourceOutput {
	return o
}

func (o SubResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SubResource) string { return v.Id }).(pulumi.StringOutput)
}

type SubResourceArrayOutput struct{ *pulumi.OutputState }

func (SubResourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubResource)(nil)).Elem()
}

func (o SubResourceArrayOutput) ToSubResourceArrayOutput() SubResourceArrayOutput {
	return o
}

func (o SubResourceArrayOutput) ToSubResourceArrayOutputWithContext(ctx context.Context) SubResourceArrayOutput {
	return o
}

func (o SubResourceArrayOutput) Index(i pulumi.IntInput) SubResourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubResource {
		return vs[0].([]SubResource)[vs[1].(int)]
	}).(SubResourceOutput)
}

type SubResourceResponse struct {
	Id string `pulumi:"id"`
}

type SubResourceResponseOutput struct{ *pulumi.OutputState }

func (SubResourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResourceResponse)(nil)).Elem()
}

func (o SubResourceResponseOutput) ToSubResourceResponseOutput() SubResourceResponseOutput {
	return o
}

func (o SubResourceResponseOutput) ToSubResourceResponseOutputWithContext(ctx context.Context) SubResourceResponseOutput {
	return o
}

func (o SubResourceResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SubResourceResponse) string { return v.Id }).(pulumi.StringOutput)
}

type SubResourceResponseArrayOutput struct{ *pulumi.OutputState }

func (SubResourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubResourceResponse)(nil)).Elem()
}

func (o SubResourceResponseArrayOutput) ToSubResourceResponseArrayOutput() SubResourceResponseArrayOutput {
	return o
}

func (o SubResourceResponseArrayOutput) ToSubResourceResponseArrayOutputWithContext(ctx context.Context) SubResourceResponseArrayOutput {
	return o
}

func (o SubResourceResponseArrayOutput) Index(i pulumi.IntInput) SubResourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubResourceResponse {
		return vs[0].([]SubResourceResponse)[vs[1].(int)]
	}).(SubResourceResponseOutput)
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

type UserAssignedIdentityResponse struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}

type UserAssignedIdentityResponseOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedIdentityResponseOutput) ToUserAssignedIdentityResponseOutput() UserAssignedIdentityResponseOutput {
	return o
}

func (o UserAssignedIdentityResponseOutput) ToUserAssignedIdentityResponseOutputWithContext(ctx context.Context) UserAssignedIdentityResponseOutput {
	return o
}

func (o UserAssignedIdentityResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o UserAssignedIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type UserAssignedIdentityResponseMapOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedIdentityResponseMapOutput) ToUserAssignedIdentityResponseMapOutput() UserAssignedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedIdentityResponseMapOutput) ToUserAssignedIdentityResponseMapOutputWithContext(ctx context.Context) UserAssignedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedIdentityResponseMapOutput) MapIndex(k pulumi.StringInput) UserAssignedIdentityResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserAssignedIdentityResponse {
		return vs[0].(map[string]UserAssignedIdentityResponse)[vs[1].(string)]
	}).(UserAssignedIdentityResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(AmbrOutput{})
	pulumi.RegisterOutputType(AmbrPtrOutput{})
	pulumi.RegisterOutputType(AmbrResponseOutput{})
	pulumi.RegisterOutputType(AmbrResponsePtrOutput{})
	pulumi.RegisterOutputType(AttachedDataNetworkResourceIdOutput{})
	pulumi.RegisterOutputType(AttachedDataNetworkResourceIdPtrOutput{})
	pulumi.RegisterOutputType(AttachedDataNetworkResourceIdResponseOutput{})
	pulumi.RegisterOutputType(AttachedDataNetworkResourceIdResponsePtrOutput{})
	pulumi.RegisterOutputType(AzureStackEdgeDeviceResourceIdOutput{})
	pulumi.RegisterOutputType(AzureStackEdgeDeviceResourceIdPtrOutput{})
	pulumi.RegisterOutputType(AzureStackEdgeDeviceResourceIdResponseOutput{})
	pulumi.RegisterOutputType(AzureStackEdgeDeviceResourceIdResponsePtrOutput{})
	pulumi.RegisterOutputType(ConnectedClusterResourceIdOutput{})
	pulumi.RegisterOutputType(ConnectedClusterResourceIdPtrOutput{})
	pulumi.RegisterOutputType(ConnectedClusterResourceIdResponseOutput{})
	pulumi.RegisterOutputType(ConnectedClusterResourceIdResponsePtrOutput{})
	pulumi.RegisterOutputType(CustomLocationResourceIdOutput{})
	pulumi.RegisterOutputType(CustomLocationResourceIdPtrOutput{})
	pulumi.RegisterOutputType(CustomLocationResourceIdResponseOutput{})
	pulumi.RegisterOutputType(CustomLocationResourceIdResponsePtrOutput{})
	pulumi.RegisterOutputType(DataNetworkConfigurationOutput{})
	pulumi.RegisterOutputType(DataNetworkConfigurationArrayOutput{})
	pulumi.RegisterOutputType(DataNetworkConfigurationResponseOutput{})
	pulumi.RegisterOutputType(DataNetworkConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(DataNetworkResourceIdOutput{})
	pulumi.RegisterOutputType(DataNetworkResourceIdResponseOutput{})
	pulumi.RegisterOutputType(InterfacePropertiesOutput{})
	pulumi.RegisterOutputType(InterfacePropertiesResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultCertificateOutput{})
	pulumi.RegisterOutputType(KeyVaultCertificatePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultCertificateResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultCertificateResponsePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyPtrOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyResponsePtrOutput{})
	pulumi.RegisterOutputType(LocalDiagnosticsAccessConfigurationOutput{})
	pulumi.RegisterOutputType(LocalDiagnosticsAccessConfigurationPtrOutput{})
	pulumi.RegisterOutputType(LocalDiagnosticsAccessConfigurationResponseOutput{})
	pulumi.RegisterOutputType(LocalDiagnosticsAccessConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityPtrOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityResponseOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(MobileNetworkResourceIdOutput{})
	pulumi.RegisterOutputType(MobileNetworkResourceIdPtrOutput{})
	pulumi.RegisterOutputType(MobileNetworkResourceIdResponseOutput{})
	pulumi.RegisterOutputType(MobileNetworkResourceIdResponsePtrOutput{})
	pulumi.RegisterOutputType(NaptConfigurationOutput{})
	pulumi.RegisterOutputType(NaptConfigurationPtrOutput{})
	pulumi.RegisterOutputType(NaptConfigurationResponseOutput{})
	pulumi.RegisterOutputType(NaptConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(PccRuleConfigurationOutput{})
	pulumi.RegisterOutputType(PccRuleConfigurationArrayOutput{})
	pulumi.RegisterOutputType(PccRuleConfigurationResponseOutput{})
	pulumi.RegisterOutputType(PccRuleConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(PccRuleQosPolicyOutput{})
	pulumi.RegisterOutputType(PccRuleQosPolicyPtrOutput{})
	pulumi.RegisterOutputType(PccRuleQosPolicyResponseOutput{})
	pulumi.RegisterOutputType(PccRuleQosPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(PinholeTimeoutsOutput{})
	pulumi.RegisterOutputType(PinholeTimeoutsPtrOutput{})
	pulumi.RegisterOutputType(PinholeTimeoutsResponseOutput{})
	pulumi.RegisterOutputType(PinholeTimeoutsResponsePtrOutput{})
	pulumi.RegisterOutputType(PlatformConfigurationOutput{})
	pulumi.RegisterOutputType(PlatformConfigurationPtrOutput{})
	pulumi.RegisterOutputType(PlatformConfigurationResponseOutput{})
	pulumi.RegisterOutputType(PlatformConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(PlmnIdOutput{})
	pulumi.RegisterOutputType(PlmnIdResponseOutput{})
	pulumi.RegisterOutputType(PortRangeOutput{})
	pulumi.RegisterOutputType(PortRangePtrOutput{})
	pulumi.RegisterOutputType(PortRangeResponseOutput{})
	pulumi.RegisterOutputType(PortRangeResponsePtrOutput{})
	pulumi.RegisterOutputType(PortReuseHoldTimesOutput{})
	pulumi.RegisterOutputType(PortReuseHoldTimesPtrOutput{})
	pulumi.RegisterOutputType(PortReuseHoldTimesResponseOutput{})
	pulumi.RegisterOutputType(PortReuseHoldTimesResponsePtrOutput{})
	pulumi.RegisterOutputType(QosPolicyOutput{})
	pulumi.RegisterOutputType(QosPolicyPtrOutput{})
	pulumi.RegisterOutputType(QosPolicyResponseOutput{})
	pulumi.RegisterOutputType(QosPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(ServiceDataFlowTemplateOutput{})
	pulumi.RegisterOutputType(ServiceDataFlowTemplateArrayOutput{})
	pulumi.RegisterOutputType(ServiceDataFlowTemplateResponseOutput{})
	pulumi.RegisterOutputType(ServiceDataFlowTemplateResponseArrayOutput{})
	pulumi.RegisterOutputType(ServiceResourceIdOutput{})
	pulumi.RegisterOutputType(ServiceResourceIdArrayOutput{})
	pulumi.RegisterOutputType(ServiceResourceIdResponseOutput{})
	pulumi.RegisterOutputType(ServiceResourceIdResponseArrayOutput{})
	pulumi.RegisterOutputType(SimPolicyResourceIdOutput{})
	pulumi.RegisterOutputType(SimPolicyResourceIdPtrOutput{})
	pulumi.RegisterOutputType(SimPolicyResourceIdResponseOutput{})
	pulumi.RegisterOutputType(SimPolicyResourceIdResponsePtrOutput{})
	pulumi.RegisterOutputType(SimStaticIpPropertiesOutput{})
	pulumi.RegisterOutputType(SimStaticIpPropertiesArrayOutput{})
	pulumi.RegisterOutputType(SimStaticIpPropertiesResponseOutput{})
	pulumi.RegisterOutputType(SimStaticIpPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(SimStaticIpPropertiesResponseStaticIpOutput{})
	pulumi.RegisterOutputType(SimStaticIpPropertiesResponseStaticIpPtrOutput{})
	pulumi.RegisterOutputType(SimStaticIpPropertiesStaticIpOutput{})
	pulumi.RegisterOutputType(SimStaticIpPropertiesStaticIpPtrOutput{})
	pulumi.RegisterOutputType(SliceConfigurationOutput{})
	pulumi.RegisterOutputType(SliceConfigurationArrayOutput{})
	pulumi.RegisterOutputType(SliceConfigurationResponseOutput{})
	pulumi.RegisterOutputType(SliceConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(SliceResourceIdOutput{})
	pulumi.RegisterOutputType(SliceResourceIdPtrOutput{})
	pulumi.RegisterOutputType(SliceResourceIdResponseOutput{})
	pulumi.RegisterOutputType(SliceResourceIdResponsePtrOutput{})
	pulumi.RegisterOutputType(SnssaiOutput{})
	pulumi.RegisterOutputType(SnssaiResponseOutput{})
	pulumi.RegisterOutputType(SubResourceOutput{})
	pulumi.RegisterOutputType(SubResourceArrayOutput{})
	pulumi.RegisterOutputType(SubResourceResponseOutput{})
	pulumi.RegisterOutputType(SubResourceResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseMapOutput{})
}
