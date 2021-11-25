


package v20190801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CacheHealthResponse struct {
	State             *string `pulumi:"state"`
	StatusDescription *string `pulumi:"statusDescription"`
}





type CacheHealthResponseInput interface {
	pulumi.Input

	ToCacheHealthResponseOutput() CacheHealthResponseOutput
	ToCacheHealthResponseOutputWithContext(context.Context) CacheHealthResponseOutput
}

type CacheHealthResponseArgs struct {
	State             pulumi.StringPtrInput `pulumi:"state"`
	StatusDescription pulumi.StringPtrInput `pulumi:"statusDescription"`
}

func (CacheHealthResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheHealthResponse)(nil)).Elem()
}

func (i CacheHealthResponseArgs) ToCacheHealthResponseOutput() CacheHealthResponseOutput {
	return i.ToCacheHealthResponseOutputWithContext(context.Background())
}

func (i CacheHealthResponseArgs) ToCacheHealthResponseOutputWithContext(ctx context.Context) CacheHealthResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheHealthResponseOutput)
}

func (i CacheHealthResponseArgs) ToCacheHealthResponsePtrOutput() CacheHealthResponsePtrOutput {
	return i.ToCacheHealthResponsePtrOutputWithContext(context.Background())
}

func (i CacheHealthResponseArgs) ToCacheHealthResponsePtrOutputWithContext(ctx context.Context) CacheHealthResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheHealthResponseOutput).ToCacheHealthResponsePtrOutputWithContext(ctx)
}









type CacheHealthResponsePtrInput interface {
	pulumi.Input

	ToCacheHealthResponsePtrOutput() CacheHealthResponsePtrOutput
	ToCacheHealthResponsePtrOutputWithContext(context.Context) CacheHealthResponsePtrOutput
}

type cacheHealthResponsePtrType CacheHealthResponseArgs

func CacheHealthResponsePtr(v *CacheHealthResponseArgs) CacheHealthResponsePtrInput {
	return (*cacheHealthResponsePtrType)(v)
}

func (*cacheHealthResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheHealthResponse)(nil)).Elem()
}

func (i *cacheHealthResponsePtrType) ToCacheHealthResponsePtrOutput() CacheHealthResponsePtrOutput {
	return i.ToCacheHealthResponsePtrOutputWithContext(context.Background())
}

func (i *cacheHealthResponsePtrType) ToCacheHealthResponsePtrOutputWithContext(ctx context.Context) CacheHealthResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheHealthResponsePtrOutput)
}

type CacheHealthResponseOutput struct{ *pulumi.OutputState }

func (CacheHealthResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheHealthResponse)(nil)).Elem()
}

func (o CacheHealthResponseOutput) ToCacheHealthResponseOutput() CacheHealthResponseOutput {
	return o
}

func (o CacheHealthResponseOutput) ToCacheHealthResponseOutputWithContext(ctx context.Context) CacheHealthResponseOutput {
	return o
}

func (o CacheHealthResponseOutput) ToCacheHealthResponsePtrOutput() CacheHealthResponsePtrOutput {
	return o.ToCacheHealthResponsePtrOutputWithContext(context.Background())
}

func (o CacheHealthResponseOutput) ToCacheHealthResponsePtrOutputWithContext(ctx context.Context) CacheHealthResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CacheHealthResponse) *CacheHealthResponse {
		return &v
	}).(CacheHealthResponsePtrOutput)
}

func (o CacheHealthResponseOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheHealthResponse) *string { return v.State }).(pulumi.StringPtrOutput)
}

func (o CacheHealthResponseOutput) StatusDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheHealthResponse) *string { return v.StatusDescription }).(pulumi.StringPtrOutput)
}

type CacheHealthResponsePtrOutput struct{ *pulumi.OutputState }

func (CacheHealthResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheHealthResponse)(nil)).Elem()
}

func (o CacheHealthResponsePtrOutput) ToCacheHealthResponsePtrOutput() CacheHealthResponsePtrOutput {
	return o
}

func (o CacheHealthResponsePtrOutput) ToCacheHealthResponsePtrOutputWithContext(ctx context.Context) CacheHealthResponsePtrOutput {
	return o
}

func (o CacheHealthResponsePtrOutput) Elem() CacheHealthResponseOutput {
	return o.ApplyT(func(v *CacheHealthResponse) CacheHealthResponse {
		if v != nil {
			return *v
		}
		var ret CacheHealthResponse
		return ret
	}).(CacheHealthResponseOutput)
}

func (o CacheHealthResponsePtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheHealthResponse) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

func (o CacheHealthResponsePtrOutput) StatusDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheHealthResponse) *string {
		if v == nil {
			return nil
		}
		return v.StatusDescription
	}).(pulumi.StringPtrOutput)
}

type CacheResponseSku struct {
	Name *string `pulumi:"name"`
}





type CacheResponseSkuInput interface {
	pulumi.Input

	ToCacheResponseSkuOutput() CacheResponseSkuOutput
	ToCacheResponseSkuOutputWithContext(context.Context) CacheResponseSkuOutput
}

type CacheResponseSkuArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (CacheResponseSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheResponseSku)(nil)).Elem()
}

func (i CacheResponseSkuArgs) ToCacheResponseSkuOutput() CacheResponseSkuOutput {
	return i.ToCacheResponseSkuOutputWithContext(context.Background())
}

func (i CacheResponseSkuArgs) ToCacheResponseSkuOutputWithContext(ctx context.Context) CacheResponseSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheResponseSkuOutput)
}

func (i CacheResponseSkuArgs) ToCacheResponseSkuPtrOutput() CacheResponseSkuPtrOutput {
	return i.ToCacheResponseSkuPtrOutputWithContext(context.Background())
}

func (i CacheResponseSkuArgs) ToCacheResponseSkuPtrOutputWithContext(ctx context.Context) CacheResponseSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheResponseSkuOutput).ToCacheResponseSkuPtrOutputWithContext(ctx)
}









type CacheResponseSkuPtrInput interface {
	pulumi.Input

	ToCacheResponseSkuPtrOutput() CacheResponseSkuPtrOutput
	ToCacheResponseSkuPtrOutputWithContext(context.Context) CacheResponseSkuPtrOutput
}

type cacheResponseSkuPtrType CacheResponseSkuArgs

func CacheResponseSkuPtr(v *CacheResponseSkuArgs) CacheResponseSkuPtrInput {
	return (*cacheResponseSkuPtrType)(v)
}

func (*cacheResponseSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheResponseSku)(nil)).Elem()
}

func (i *cacheResponseSkuPtrType) ToCacheResponseSkuPtrOutput() CacheResponseSkuPtrOutput {
	return i.ToCacheResponseSkuPtrOutputWithContext(context.Background())
}

func (i *cacheResponseSkuPtrType) ToCacheResponseSkuPtrOutputWithContext(ctx context.Context) CacheResponseSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheResponseSkuPtrOutput)
}

type CacheResponseSkuOutput struct{ *pulumi.OutputState }

func (CacheResponseSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheResponseSku)(nil)).Elem()
}

func (o CacheResponseSkuOutput) ToCacheResponseSkuOutput() CacheResponseSkuOutput {
	return o
}

func (o CacheResponseSkuOutput) ToCacheResponseSkuOutputWithContext(ctx context.Context) CacheResponseSkuOutput {
	return o
}

func (o CacheResponseSkuOutput) ToCacheResponseSkuPtrOutput() CacheResponseSkuPtrOutput {
	return o.ToCacheResponseSkuPtrOutputWithContext(context.Background())
}

func (o CacheResponseSkuOutput) ToCacheResponseSkuPtrOutputWithContext(ctx context.Context) CacheResponseSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CacheResponseSku) *CacheResponseSku {
		return &v
	}).(CacheResponseSkuPtrOutput)
}

func (o CacheResponseSkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheResponseSku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type CacheResponseSkuPtrOutput struct{ *pulumi.OutputState }

func (CacheResponseSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheResponseSku)(nil)).Elem()
}

func (o CacheResponseSkuPtrOutput) ToCacheResponseSkuPtrOutput() CacheResponseSkuPtrOutput {
	return o
}

func (o CacheResponseSkuPtrOutput) ToCacheResponseSkuPtrOutputWithContext(ctx context.Context) CacheResponseSkuPtrOutput {
	return o
}

func (o CacheResponseSkuPtrOutput) Elem() CacheResponseSkuOutput {
	return o.ApplyT(func(v *CacheResponseSku) CacheResponseSku {
		if v != nil {
			return *v
		}
		var ret CacheResponseSku
		return ret
	}).(CacheResponseSkuOutput)
}

func (o CacheResponseSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheResponseSku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type CacheSku struct {
	Name *string `pulumi:"name"`
}





type CacheSkuInput interface {
	pulumi.Input

	ToCacheSkuOutput() CacheSkuOutput
	ToCacheSkuOutputWithContext(context.Context) CacheSkuOutput
}

type CacheSkuArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (CacheSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheSku)(nil)).Elem()
}

func (i CacheSkuArgs) ToCacheSkuOutput() CacheSkuOutput {
	return i.ToCacheSkuOutputWithContext(context.Background())
}

func (i CacheSkuArgs) ToCacheSkuOutputWithContext(ctx context.Context) CacheSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheSkuOutput)
}

func (i CacheSkuArgs) ToCacheSkuPtrOutput() CacheSkuPtrOutput {
	return i.ToCacheSkuPtrOutputWithContext(context.Background())
}

func (i CacheSkuArgs) ToCacheSkuPtrOutputWithContext(ctx context.Context) CacheSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheSkuOutput).ToCacheSkuPtrOutputWithContext(ctx)
}









type CacheSkuPtrInput interface {
	pulumi.Input

	ToCacheSkuPtrOutput() CacheSkuPtrOutput
	ToCacheSkuPtrOutputWithContext(context.Context) CacheSkuPtrOutput
}

type cacheSkuPtrType CacheSkuArgs

func CacheSkuPtr(v *CacheSkuArgs) CacheSkuPtrInput {
	return (*cacheSkuPtrType)(v)
}

func (*cacheSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheSku)(nil)).Elem()
}

func (i *cacheSkuPtrType) ToCacheSkuPtrOutput() CacheSkuPtrOutput {
	return i.ToCacheSkuPtrOutputWithContext(context.Background())
}

func (i *cacheSkuPtrType) ToCacheSkuPtrOutputWithContext(ctx context.Context) CacheSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheSkuPtrOutput)
}

type CacheSkuOutput struct{ *pulumi.OutputState }

func (CacheSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheSku)(nil)).Elem()
}

func (o CacheSkuOutput) ToCacheSkuOutput() CacheSkuOutput {
	return o
}

func (o CacheSkuOutput) ToCacheSkuOutputWithContext(ctx context.Context) CacheSkuOutput {
	return o
}

func (o CacheSkuOutput) ToCacheSkuPtrOutput() CacheSkuPtrOutput {
	return o.ToCacheSkuPtrOutputWithContext(context.Background())
}

func (o CacheSkuOutput) ToCacheSkuPtrOutputWithContext(ctx context.Context) CacheSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CacheSku) *CacheSku {
		return &v
	}).(CacheSkuPtrOutput)
}

func (o CacheSkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CacheSku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type CacheSkuPtrOutput struct{ *pulumi.OutputState }

func (CacheSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheSku)(nil)).Elem()
}

func (o CacheSkuPtrOutput) ToCacheSkuPtrOutput() CacheSkuPtrOutput {
	return o
}

func (o CacheSkuPtrOutput) ToCacheSkuPtrOutputWithContext(ctx context.Context) CacheSkuPtrOutput {
	return o
}

func (o CacheSkuPtrOutput) Elem() CacheSkuOutput {
	return o.ApplyT(func(v *CacheSku) CacheSku {
		if v != nil {
			return *v
		}
		var ret CacheSku
		return ret
	}).(CacheSkuOutput)
}

func (o CacheSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheSku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type CacheUpgradeStatusResponse struct {
	CurrentFirmwareVersion string `pulumi:"currentFirmwareVersion"`
	FirmwareUpdateDeadline string `pulumi:"firmwareUpdateDeadline"`
	FirmwareUpdateStatus   string `pulumi:"firmwareUpdateStatus"`
	LastFirmwareUpdate     string `pulumi:"lastFirmwareUpdate"`
	PendingFirmwareVersion string `pulumi:"pendingFirmwareVersion"`
}





type CacheUpgradeStatusResponseInput interface {
	pulumi.Input

	ToCacheUpgradeStatusResponseOutput() CacheUpgradeStatusResponseOutput
	ToCacheUpgradeStatusResponseOutputWithContext(context.Context) CacheUpgradeStatusResponseOutput
}

type CacheUpgradeStatusResponseArgs struct {
	CurrentFirmwareVersion pulumi.StringInput `pulumi:"currentFirmwareVersion"`
	FirmwareUpdateDeadline pulumi.StringInput `pulumi:"firmwareUpdateDeadline"`
	FirmwareUpdateStatus   pulumi.StringInput `pulumi:"firmwareUpdateStatus"`
	LastFirmwareUpdate     pulumi.StringInput `pulumi:"lastFirmwareUpdate"`
	PendingFirmwareVersion pulumi.StringInput `pulumi:"pendingFirmwareVersion"`
}

func (CacheUpgradeStatusResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheUpgradeStatusResponse)(nil)).Elem()
}

func (i CacheUpgradeStatusResponseArgs) ToCacheUpgradeStatusResponseOutput() CacheUpgradeStatusResponseOutput {
	return i.ToCacheUpgradeStatusResponseOutputWithContext(context.Background())
}

func (i CacheUpgradeStatusResponseArgs) ToCacheUpgradeStatusResponseOutputWithContext(ctx context.Context) CacheUpgradeStatusResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheUpgradeStatusResponseOutput)
}

func (i CacheUpgradeStatusResponseArgs) ToCacheUpgradeStatusResponsePtrOutput() CacheUpgradeStatusResponsePtrOutput {
	return i.ToCacheUpgradeStatusResponsePtrOutputWithContext(context.Background())
}

func (i CacheUpgradeStatusResponseArgs) ToCacheUpgradeStatusResponsePtrOutputWithContext(ctx context.Context) CacheUpgradeStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheUpgradeStatusResponseOutput).ToCacheUpgradeStatusResponsePtrOutputWithContext(ctx)
}









type CacheUpgradeStatusResponsePtrInput interface {
	pulumi.Input

	ToCacheUpgradeStatusResponsePtrOutput() CacheUpgradeStatusResponsePtrOutput
	ToCacheUpgradeStatusResponsePtrOutputWithContext(context.Context) CacheUpgradeStatusResponsePtrOutput
}

type cacheUpgradeStatusResponsePtrType CacheUpgradeStatusResponseArgs

func CacheUpgradeStatusResponsePtr(v *CacheUpgradeStatusResponseArgs) CacheUpgradeStatusResponsePtrInput {
	return (*cacheUpgradeStatusResponsePtrType)(v)
}

func (*cacheUpgradeStatusResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheUpgradeStatusResponse)(nil)).Elem()
}

func (i *cacheUpgradeStatusResponsePtrType) ToCacheUpgradeStatusResponsePtrOutput() CacheUpgradeStatusResponsePtrOutput {
	return i.ToCacheUpgradeStatusResponsePtrOutputWithContext(context.Background())
}

func (i *cacheUpgradeStatusResponsePtrType) ToCacheUpgradeStatusResponsePtrOutputWithContext(ctx context.Context) CacheUpgradeStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheUpgradeStatusResponsePtrOutput)
}

type CacheUpgradeStatusResponseOutput struct{ *pulumi.OutputState }

func (CacheUpgradeStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheUpgradeStatusResponse)(nil)).Elem()
}

func (o CacheUpgradeStatusResponseOutput) ToCacheUpgradeStatusResponseOutput() CacheUpgradeStatusResponseOutput {
	return o
}

func (o CacheUpgradeStatusResponseOutput) ToCacheUpgradeStatusResponseOutputWithContext(ctx context.Context) CacheUpgradeStatusResponseOutput {
	return o
}

func (o CacheUpgradeStatusResponseOutput) ToCacheUpgradeStatusResponsePtrOutput() CacheUpgradeStatusResponsePtrOutput {
	return o.ToCacheUpgradeStatusResponsePtrOutputWithContext(context.Background())
}

func (o CacheUpgradeStatusResponseOutput) ToCacheUpgradeStatusResponsePtrOutputWithContext(ctx context.Context) CacheUpgradeStatusResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CacheUpgradeStatusResponse) *CacheUpgradeStatusResponse {
		return &v
	}).(CacheUpgradeStatusResponsePtrOutput)
}

func (o CacheUpgradeStatusResponseOutput) CurrentFirmwareVersion() pulumi.StringOutput {
	return o.ApplyT(func(v CacheUpgradeStatusResponse) string { return v.CurrentFirmwareVersion }).(pulumi.StringOutput)
}

func (o CacheUpgradeStatusResponseOutput) FirmwareUpdateDeadline() pulumi.StringOutput {
	return o.ApplyT(func(v CacheUpgradeStatusResponse) string { return v.FirmwareUpdateDeadline }).(pulumi.StringOutput)
}

func (o CacheUpgradeStatusResponseOutput) FirmwareUpdateStatus() pulumi.StringOutput {
	return o.ApplyT(func(v CacheUpgradeStatusResponse) string { return v.FirmwareUpdateStatus }).(pulumi.StringOutput)
}

func (o CacheUpgradeStatusResponseOutput) LastFirmwareUpdate() pulumi.StringOutput {
	return o.ApplyT(func(v CacheUpgradeStatusResponse) string { return v.LastFirmwareUpdate }).(pulumi.StringOutput)
}

func (o CacheUpgradeStatusResponseOutput) PendingFirmwareVersion() pulumi.StringOutput {
	return o.ApplyT(func(v CacheUpgradeStatusResponse) string { return v.PendingFirmwareVersion }).(pulumi.StringOutput)
}

type CacheUpgradeStatusResponsePtrOutput struct{ *pulumi.OutputState }

func (CacheUpgradeStatusResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheUpgradeStatusResponse)(nil)).Elem()
}

func (o CacheUpgradeStatusResponsePtrOutput) ToCacheUpgradeStatusResponsePtrOutput() CacheUpgradeStatusResponsePtrOutput {
	return o
}

func (o CacheUpgradeStatusResponsePtrOutput) ToCacheUpgradeStatusResponsePtrOutputWithContext(ctx context.Context) CacheUpgradeStatusResponsePtrOutput {
	return o
}

func (o CacheUpgradeStatusResponsePtrOutput) Elem() CacheUpgradeStatusResponseOutput {
	return o.ApplyT(func(v *CacheUpgradeStatusResponse) CacheUpgradeStatusResponse {
		if v != nil {
			return *v
		}
		var ret CacheUpgradeStatusResponse
		return ret
	}).(CacheUpgradeStatusResponseOutput)
}

func (o CacheUpgradeStatusResponsePtrOutput) CurrentFirmwareVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheUpgradeStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CurrentFirmwareVersion
	}).(pulumi.StringPtrOutput)
}

func (o CacheUpgradeStatusResponsePtrOutput) FirmwareUpdateDeadline() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheUpgradeStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.FirmwareUpdateDeadline
	}).(pulumi.StringPtrOutput)
}

func (o CacheUpgradeStatusResponsePtrOutput) FirmwareUpdateStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheUpgradeStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.FirmwareUpdateStatus
	}).(pulumi.StringPtrOutput)
}

func (o CacheUpgradeStatusResponsePtrOutput) LastFirmwareUpdate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheUpgradeStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastFirmwareUpdate
	}).(pulumi.StringPtrOutput)
}

func (o CacheUpgradeStatusResponsePtrOutput) PendingFirmwareVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheUpgradeStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PendingFirmwareVersion
	}).(pulumi.StringPtrOutput)
}

type ClfsTarget struct {
	Target *string `pulumi:"target"`
}





type ClfsTargetInput interface {
	pulumi.Input

	ToClfsTargetOutput() ClfsTargetOutput
	ToClfsTargetOutputWithContext(context.Context) ClfsTargetOutput
}

type ClfsTargetArgs struct {
	Target pulumi.StringPtrInput `pulumi:"target"`
}

func (ClfsTargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClfsTarget)(nil)).Elem()
}

func (i ClfsTargetArgs) ToClfsTargetOutput() ClfsTargetOutput {
	return i.ToClfsTargetOutputWithContext(context.Background())
}

func (i ClfsTargetArgs) ToClfsTargetOutputWithContext(ctx context.Context) ClfsTargetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClfsTargetOutput)
}

func (i ClfsTargetArgs) ToClfsTargetPtrOutput() ClfsTargetPtrOutput {
	return i.ToClfsTargetPtrOutputWithContext(context.Background())
}

func (i ClfsTargetArgs) ToClfsTargetPtrOutputWithContext(ctx context.Context) ClfsTargetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClfsTargetOutput).ToClfsTargetPtrOutputWithContext(ctx)
}









type ClfsTargetPtrInput interface {
	pulumi.Input

	ToClfsTargetPtrOutput() ClfsTargetPtrOutput
	ToClfsTargetPtrOutputWithContext(context.Context) ClfsTargetPtrOutput
}

type clfsTargetPtrType ClfsTargetArgs

func ClfsTargetPtr(v *ClfsTargetArgs) ClfsTargetPtrInput {
	return (*clfsTargetPtrType)(v)
}

func (*clfsTargetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClfsTarget)(nil)).Elem()
}

func (i *clfsTargetPtrType) ToClfsTargetPtrOutput() ClfsTargetPtrOutput {
	return i.ToClfsTargetPtrOutputWithContext(context.Background())
}

func (i *clfsTargetPtrType) ToClfsTargetPtrOutputWithContext(ctx context.Context) ClfsTargetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClfsTargetPtrOutput)
}

type ClfsTargetOutput struct{ *pulumi.OutputState }

func (ClfsTargetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClfsTarget)(nil)).Elem()
}

func (o ClfsTargetOutput) ToClfsTargetOutput() ClfsTargetOutput {
	return o
}

func (o ClfsTargetOutput) ToClfsTargetOutputWithContext(ctx context.Context) ClfsTargetOutput {
	return o
}

func (o ClfsTargetOutput) ToClfsTargetPtrOutput() ClfsTargetPtrOutput {
	return o.ToClfsTargetPtrOutputWithContext(context.Background())
}

func (o ClfsTargetOutput) ToClfsTargetPtrOutputWithContext(ctx context.Context) ClfsTargetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClfsTarget) *ClfsTarget {
		return &v
	}).(ClfsTargetPtrOutput)
}

func (o ClfsTargetOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClfsTarget) *string { return v.Target }).(pulumi.StringPtrOutput)
}

type ClfsTargetPtrOutput struct{ *pulumi.OutputState }

func (ClfsTargetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClfsTarget)(nil)).Elem()
}

func (o ClfsTargetPtrOutput) ToClfsTargetPtrOutput() ClfsTargetPtrOutput {
	return o
}

func (o ClfsTargetPtrOutput) ToClfsTargetPtrOutputWithContext(ctx context.Context) ClfsTargetPtrOutput {
	return o
}

func (o ClfsTargetPtrOutput) Elem() ClfsTargetOutput {
	return o.ApplyT(func(v *ClfsTarget) ClfsTarget {
		if v != nil {
			return *v
		}
		var ret ClfsTarget
		return ret
	}).(ClfsTargetOutput)
}

func (o ClfsTargetPtrOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClfsTarget) *string {
		if v == nil {
			return nil
		}
		return v.Target
	}).(pulumi.StringPtrOutput)
}

type ClfsTargetResponse struct {
	Target *string `pulumi:"target"`
}





type ClfsTargetResponseInput interface {
	pulumi.Input

	ToClfsTargetResponseOutput() ClfsTargetResponseOutput
	ToClfsTargetResponseOutputWithContext(context.Context) ClfsTargetResponseOutput
}

type ClfsTargetResponseArgs struct {
	Target pulumi.StringPtrInput `pulumi:"target"`
}

func (ClfsTargetResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClfsTargetResponse)(nil)).Elem()
}

func (i ClfsTargetResponseArgs) ToClfsTargetResponseOutput() ClfsTargetResponseOutput {
	return i.ToClfsTargetResponseOutputWithContext(context.Background())
}

func (i ClfsTargetResponseArgs) ToClfsTargetResponseOutputWithContext(ctx context.Context) ClfsTargetResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClfsTargetResponseOutput)
}

func (i ClfsTargetResponseArgs) ToClfsTargetResponsePtrOutput() ClfsTargetResponsePtrOutput {
	return i.ToClfsTargetResponsePtrOutputWithContext(context.Background())
}

func (i ClfsTargetResponseArgs) ToClfsTargetResponsePtrOutputWithContext(ctx context.Context) ClfsTargetResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClfsTargetResponseOutput).ToClfsTargetResponsePtrOutputWithContext(ctx)
}









type ClfsTargetResponsePtrInput interface {
	pulumi.Input

	ToClfsTargetResponsePtrOutput() ClfsTargetResponsePtrOutput
	ToClfsTargetResponsePtrOutputWithContext(context.Context) ClfsTargetResponsePtrOutput
}

type clfsTargetResponsePtrType ClfsTargetResponseArgs

func ClfsTargetResponsePtr(v *ClfsTargetResponseArgs) ClfsTargetResponsePtrInput {
	return (*clfsTargetResponsePtrType)(v)
}

func (*clfsTargetResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClfsTargetResponse)(nil)).Elem()
}

func (i *clfsTargetResponsePtrType) ToClfsTargetResponsePtrOutput() ClfsTargetResponsePtrOutput {
	return i.ToClfsTargetResponsePtrOutputWithContext(context.Background())
}

func (i *clfsTargetResponsePtrType) ToClfsTargetResponsePtrOutputWithContext(ctx context.Context) ClfsTargetResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClfsTargetResponsePtrOutput)
}

type ClfsTargetResponseOutput struct{ *pulumi.OutputState }

func (ClfsTargetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClfsTargetResponse)(nil)).Elem()
}

func (o ClfsTargetResponseOutput) ToClfsTargetResponseOutput() ClfsTargetResponseOutput {
	return o
}

func (o ClfsTargetResponseOutput) ToClfsTargetResponseOutputWithContext(ctx context.Context) ClfsTargetResponseOutput {
	return o
}

func (o ClfsTargetResponseOutput) ToClfsTargetResponsePtrOutput() ClfsTargetResponsePtrOutput {
	return o.ToClfsTargetResponsePtrOutputWithContext(context.Background())
}

func (o ClfsTargetResponseOutput) ToClfsTargetResponsePtrOutputWithContext(ctx context.Context) ClfsTargetResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClfsTargetResponse) *ClfsTargetResponse {
		return &v
	}).(ClfsTargetResponsePtrOutput)
}

func (o ClfsTargetResponseOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClfsTargetResponse) *string { return v.Target }).(pulumi.StringPtrOutput)
}

type ClfsTargetResponsePtrOutput struct{ *pulumi.OutputState }

func (ClfsTargetResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClfsTargetResponse)(nil)).Elem()
}

func (o ClfsTargetResponsePtrOutput) ToClfsTargetResponsePtrOutput() ClfsTargetResponsePtrOutput {
	return o
}

func (o ClfsTargetResponsePtrOutput) ToClfsTargetResponsePtrOutputWithContext(ctx context.Context) ClfsTargetResponsePtrOutput {
	return o
}

func (o ClfsTargetResponsePtrOutput) Elem() ClfsTargetResponseOutput {
	return o.ApplyT(func(v *ClfsTargetResponse) ClfsTargetResponse {
		if v != nil {
			return *v
		}
		var ret ClfsTargetResponse
		return ret
	}).(ClfsTargetResponseOutput)
}

func (o ClfsTargetResponsePtrOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClfsTargetResponse) *string {
		if v == nil {
			return nil
		}
		return v.Target
	}).(pulumi.StringPtrOutput)
}

type NamespaceJunction struct {
	NamespacePath *string `pulumi:"namespacePath"`
	NfsExport     *string `pulumi:"nfsExport"`
	TargetPath    *string `pulumi:"targetPath"`
}





type NamespaceJunctionInput interface {
	pulumi.Input

	ToNamespaceJunctionOutput() NamespaceJunctionOutput
	ToNamespaceJunctionOutputWithContext(context.Context) NamespaceJunctionOutput
}

type NamespaceJunctionArgs struct {
	NamespacePath pulumi.StringPtrInput `pulumi:"namespacePath"`
	NfsExport     pulumi.StringPtrInput `pulumi:"nfsExport"`
	TargetPath    pulumi.StringPtrInput `pulumi:"targetPath"`
}

func (NamespaceJunctionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NamespaceJunction)(nil)).Elem()
}

func (i NamespaceJunctionArgs) ToNamespaceJunctionOutput() NamespaceJunctionOutput {
	return i.ToNamespaceJunctionOutputWithContext(context.Background())
}

func (i NamespaceJunctionArgs) ToNamespaceJunctionOutputWithContext(ctx context.Context) NamespaceJunctionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceJunctionOutput)
}





type NamespaceJunctionArrayInput interface {
	pulumi.Input

	ToNamespaceJunctionArrayOutput() NamespaceJunctionArrayOutput
	ToNamespaceJunctionArrayOutputWithContext(context.Context) NamespaceJunctionArrayOutput
}

type NamespaceJunctionArray []NamespaceJunctionInput

func (NamespaceJunctionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NamespaceJunction)(nil)).Elem()
}

func (i NamespaceJunctionArray) ToNamespaceJunctionArrayOutput() NamespaceJunctionArrayOutput {
	return i.ToNamespaceJunctionArrayOutputWithContext(context.Background())
}

func (i NamespaceJunctionArray) ToNamespaceJunctionArrayOutputWithContext(ctx context.Context) NamespaceJunctionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceJunctionArrayOutput)
}

type NamespaceJunctionOutput struct{ *pulumi.OutputState }

func (NamespaceJunctionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NamespaceJunction)(nil)).Elem()
}

func (o NamespaceJunctionOutput) ToNamespaceJunctionOutput() NamespaceJunctionOutput {
	return o
}

func (o NamespaceJunctionOutput) ToNamespaceJunctionOutputWithContext(ctx context.Context) NamespaceJunctionOutput {
	return o
}

func (o NamespaceJunctionOutput) NamespacePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespaceJunction) *string { return v.NamespacePath }).(pulumi.StringPtrOutput)
}

func (o NamespaceJunctionOutput) NfsExport() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespaceJunction) *string { return v.NfsExport }).(pulumi.StringPtrOutput)
}

func (o NamespaceJunctionOutput) TargetPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespaceJunction) *string { return v.TargetPath }).(pulumi.StringPtrOutput)
}

type NamespaceJunctionArrayOutput struct{ *pulumi.OutputState }

func (NamespaceJunctionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NamespaceJunction)(nil)).Elem()
}

func (o NamespaceJunctionArrayOutput) ToNamespaceJunctionArrayOutput() NamespaceJunctionArrayOutput {
	return o
}

func (o NamespaceJunctionArrayOutput) ToNamespaceJunctionArrayOutputWithContext(ctx context.Context) NamespaceJunctionArrayOutput {
	return o
}

func (o NamespaceJunctionArrayOutput) Index(i pulumi.IntInput) NamespaceJunctionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NamespaceJunction {
		return vs[0].([]NamespaceJunction)[vs[1].(int)]
	}).(NamespaceJunctionOutput)
}

type NamespaceJunctionResponse struct {
	NamespacePath *string `pulumi:"namespacePath"`
	NfsExport     *string `pulumi:"nfsExport"`
	TargetPath    *string `pulumi:"targetPath"`
}





type NamespaceJunctionResponseInput interface {
	pulumi.Input

	ToNamespaceJunctionResponseOutput() NamespaceJunctionResponseOutput
	ToNamespaceJunctionResponseOutputWithContext(context.Context) NamespaceJunctionResponseOutput
}

type NamespaceJunctionResponseArgs struct {
	NamespacePath pulumi.StringPtrInput `pulumi:"namespacePath"`
	NfsExport     pulumi.StringPtrInput `pulumi:"nfsExport"`
	TargetPath    pulumi.StringPtrInput `pulumi:"targetPath"`
}

func (NamespaceJunctionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NamespaceJunctionResponse)(nil)).Elem()
}

func (i NamespaceJunctionResponseArgs) ToNamespaceJunctionResponseOutput() NamespaceJunctionResponseOutput {
	return i.ToNamespaceJunctionResponseOutputWithContext(context.Background())
}

func (i NamespaceJunctionResponseArgs) ToNamespaceJunctionResponseOutputWithContext(ctx context.Context) NamespaceJunctionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceJunctionResponseOutput)
}





type NamespaceJunctionResponseArrayInput interface {
	pulumi.Input

	ToNamespaceJunctionResponseArrayOutput() NamespaceJunctionResponseArrayOutput
	ToNamespaceJunctionResponseArrayOutputWithContext(context.Context) NamespaceJunctionResponseArrayOutput
}

type NamespaceJunctionResponseArray []NamespaceJunctionResponseInput

func (NamespaceJunctionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NamespaceJunctionResponse)(nil)).Elem()
}

func (i NamespaceJunctionResponseArray) ToNamespaceJunctionResponseArrayOutput() NamespaceJunctionResponseArrayOutput {
	return i.ToNamespaceJunctionResponseArrayOutputWithContext(context.Background())
}

func (i NamespaceJunctionResponseArray) ToNamespaceJunctionResponseArrayOutputWithContext(ctx context.Context) NamespaceJunctionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceJunctionResponseArrayOutput)
}

type NamespaceJunctionResponseOutput struct{ *pulumi.OutputState }

func (NamespaceJunctionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NamespaceJunctionResponse)(nil)).Elem()
}

func (o NamespaceJunctionResponseOutput) ToNamespaceJunctionResponseOutput() NamespaceJunctionResponseOutput {
	return o
}

func (o NamespaceJunctionResponseOutput) ToNamespaceJunctionResponseOutputWithContext(ctx context.Context) NamespaceJunctionResponseOutput {
	return o
}

func (o NamespaceJunctionResponseOutput) NamespacePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespaceJunctionResponse) *string { return v.NamespacePath }).(pulumi.StringPtrOutput)
}

func (o NamespaceJunctionResponseOutput) NfsExport() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespaceJunctionResponse) *string { return v.NfsExport }).(pulumi.StringPtrOutput)
}

func (o NamespaceJunctionResponseOutput) TargetPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespaceJunctionResponse) *string { return v.TargetPath }).(pulumi.StringPtrOutput)
}

type NamespaceJunctionResponseArrayOutput struct{ *pulumi.OutputState }

func (NamespaceJunctionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NamespaceJunctionResponse)(nil)).Elem()
}

func (o NamespaceJunctionResponseArrayOutput) ToNamespaceJunctionResponseArrayOutput() NamespaceJunctionResponseArrayOutput {
	return o
}

func (o NamespaceJunctionResponseArrayOutput) ToNamespaceJunctionResponseArrayOutputWithContext(ctx context.Context) NamespaceJunctionResponseArrayOutput {
	return o
}

func (o NamespaceJunctionResponseArrayOutput) Index(i pulumi.IntInput) NamespaceJunctionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NamespaceJunctionResponse {
		return vs[0].([]NamespaceJunctionResponse)[vs[1].(int)]
	}).(NamespaceJunctionResponseOutput)
}

type Nfs3Target struct {
	Target     *string `pulumi:"target"`
	UsageModel *string `pulumi:"usageModel"`
}





type Nfs3TargetInput interface {
	pulumi.Input

	ToNfs3TargetOutput() Nfs3TargetOutput
	ToNfs3TargetOutputWithContext(context.Context) Nfs3TargetOutput
}

type Nfs3TargetArgs struct {
	Target     pulumi.StringPtrInput `pulumi:"target"`
	UsageModel pulumi.StringPtrInput `pulumi:"usageModel"`
}

func (Nfs3TargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Nfs3Target)(nil)).Elem()
}

func (i Nfs3TargetArgs) ToNfs3TargetOutput() Nfs3TargetOutput {
	return i.ToNfs3TargetOutputWithContext(context.Background())
}

func (i Nfs3TargetArgs) ToNfs3TargetOutputWithContext(ctx context.Context) Nfs3TargetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Nfs3TargetOutput)
}

func (i Nfs3TargetArgs) ToNfs3TargetPtrOutput() Nfs3TargetPtrOutput {
	return i.ToNfs3TargetPtrOutputWithContext(context.Background())
}

func (i Nfs3TargetArgs) ToNfs3TargetPtrOutputWithContext(ctx context.Context) Nfs3TargetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Nfs3TargetOutput).ToNfs3TargetPtrOutputWithContext(ctx)
}









type Nfs3TargetPtrInput interface {
	pulumi.Input

	ToNfs3TargetPtrOutput() Nfs3TargetPtrOutput
	ToNfs3TargetPtrOutputWithContext(context.Context) Nfs3TargetPtrOutput
}

type nfs3TargetPtrType Nfs3TargetArgs

func Nfs3TargetPtr(v *Nfs3TargetArgs) Nfs3TargetPtrInput {
	return (*nfs3TargetPtrType)(v)
}

func (*nfs3TargetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Nfs3Target)(nil)).Elem()
}

func (i *nfs3TargetPtrType) ToNfs3TargetPtrOutput() Nfs3TargetPtrOutput {
	return i.ToNfs3TargetPtrOutputWithContext(context.Background())
}

func (i *nfs3TargetPtrType) ToNfs3TargetPtrOutputWithContext(ctx context.Context) Nfs3TargetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Nfs3TargetPtrOutput)
}

type Nfs3TargetOutput struct{ *pulumi.OutputState }

func (Nfs3TargetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Nfs3Target)(nil)).Elem()
}

func (o Nfs3TargetOutput) ToNfs3TargetOutput() Nfs3TargetOutput {
	return o
}

func (o Nfs3TargetOutput) ToNfs3TargetOutputWithContext(ctx context.Context) Nfs3TargetOutput {
	return o
}

func (o Nfs3TargetOutput) ToNfs3TargetPtrOutput() Nfs3TargetPtrOutput {
	return o.ToNfs3TargetPtrOutputWithContext(context.Background())
}

func (o Nfs3TargetOutput) ToNfs3TargetPtrOutputWithContext(ctx context.Context) Nfs3TargetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Nfs3Target) *Nfs3Target {
		return &v
	}).(Nfs3TargetPtrOutput)
}

func (o Nfs3TargetOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Nfs3Target) *string { return v.Target }).(pulumi.StringPtrOutput)
}

func (o Nfs3TargetOutput) UsageModel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Nfs3Target) *string { return v.UsageModel }).(pulumi.StringPtrOutput)
}

type Nfs3TargetPtrOutput struct{ *pulumi.OutputState }

func (Nfs3TargetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Nfs3Target)(nil)).Elem()
}

func (o Nfs3TargetPtrOutput) ToNfs3TargetPtrOutput() Nfs3TargetPtrOutput {
	return o
}

func (o Nfs3TargetPtrOutput) ToNfs3TargetPtrOutputWithContext(ctx context.Context) Nfs3TargetPtrOutput {
	return o
}

func (o Nfs3TargetPtrOutput) Elem() Nfs3TargetOutput {
	return o.ApplyT(func(v *Nfs3Target) Nfs3Target {
		if v != nil {
			return *v
		}
		var ret Nfs3Target
		return ret
	}).(Nfs3TargetOutput)
}

func (o Nfs3TargetPtrOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Nfs3Target) *string {
		if v == nil {
			return nil
		}
		return v.Target
	}).(pulumi.StringPtrOutput)
}

func (o Nfs3TargetPtrOutput) UsageModel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Nfs3Target) *string {
		if v == nil {
			return nil
		}
		return v.UsageModel
	}).(pulumi.StringPtrOutput)
}

type Nfs3TargetResponse struct {
	Target     *string `pulumi:"target"`
	UsageModel *string `pulumi:"usageModel"`
}





type Nfs3TargetResponseInput interface {
	pulumi.Input

	ToNfs3TargetResponseOutput() Nfs3TargetResponseOutput
	ToNfs3TargetResponseOutputWithContext(context.Context) Nfs3TargetResponseOutput
}

type Nfs3TargetResponseArgs struct {
	Target     pulumi.StringPtrInput `pulumi:"target"`
	UsageModel pulumi.StringPtrInput `pulumi:"usageModel"`
}

func (Nfs3TargetResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Nfs3TargetResponse)(nil)).Elem()
}

func (i Nfs3TargetResponseArgs) ToNfs3TargetResponseOutput() Nfs3TargetResponseOutput {
	return i.ToNfs3TargetResponseOutputWithContext(context.Background())
}

func (i Nfs3TargetResponseArgs) ToNfs3TargetResponseOutputWithContext(ctx context.Context) Nfs3TargetResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Nfs3TargetResponseOutput)
}

func (i Nfs3TargetResponseArgs) ToNfs3TargetResponsePtrOutput() Nfs3TargetResponsePtrOutput {
	return i.ToNfs3TargetResponsePtrOutputWithContext(context.Background())
}

func (i Nfs3TargetResponseArgs) ToNfs3TargetResponsePtrOutputWithContext(ctx context.Context) Nfs3TargetResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Nfs3TargetResponseOutput).ToNfs3TargetResponsePtrOutputWithContext(ctx)
}









type Nfs3TargetResponsePtrInput interface {
	pulumi.Input

	ToNfs3TargetResponsePtrOutput() Nfs3TargetResponsePtrOutput
	ToNfs3TargetResponsePtrOutputWithContext(context.Context) Nfs3TargetResponsePtrOutput
}

type nfs3TargetResponsePtrType Nfs3TargetResponseArgs

func Nfs3TargetResponsePtr(v *Nfs3TargetResponseArgs) Nfs3TargetResponsePtrInput {
	return (*nfs3TargetResponsePtrType)(v)
}

func (*nfs3TargetResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Nfs3TargetResponse)(nil)).Elem()
}

func (i *nfs3TargetResponsePtrType) ToNfs3TargetResponsePtrOutput() Nfs3TargetResponsePtrOutput {
	return i.ToNfs3TargetResponsePtrOutputWithContext(context.Background())
}

func (i *nfs3TargetResponsePtrType) ToNfs3TargetResponsePtrOutputWithContext(ctx context.Context) Nfs3TargetResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Nfs3TargetResponsePtrOutput)
}

type Nfs3TargetResponseOutput struct{ *pulumi.OutputState }

func (Nfs3TargetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Nfs3TargetResponse)(nil)).Elem()
}

func (o Nfs3TargetResponseOutput) ToNfs3TargetResponseOutput() Nfs3TargetResponseOutput {
	return o
}

func (o Nfs3TargetResponseOutput) ToNfs3TargetResponseOutputWithContext(ctx context.Context) Nfs3TargetResponseOutput {
	return o
}

func (o Nfs3TargetResponseOutput) ToNfs3TargetResponsePtrOutput() Nfs3TargetResponsePtrOutput {
	return o.ToNfs3TargetResponsePtrOutputWithContext(context.Background())
}

func (o Nfs3TargetResponseOutput) ToNfs3TargetResponsePtrOutputWithContext(ctx context.Context) Nfs3TargetResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Nfs3TargetResponse) *Nfs3TargetResponse {
		return &v
	}).(Nfs3TargetResponsePtrOutput)
}

func (o Nfs3TargetResponseOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Nfs3TargetResponse) *string { return v.Target }).(pulumi.StringPtrOutput)
}

func (o Nfs3TargetResponseOutput) UsageModel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Nfs3TargetResponse) *string { return v.UsageModel }).(pulumi.StringPtrOutput)
}

type Nfs3TargetResponsePtrOutput struct{ *pulumi.OutputState }

func (Nfs3TargetResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Nfs3TargetResponse)(nil)).Elem()
}

func (o Nfs3TargetResponsePtrOutput) ToNfs3TargetResponsePtrOutput() Nfs3TargetResponsePtrOutput {
	return o
}

func (o Nfs3TargetResponsePtrOutput) ToNfs3TargetResponsePtrOutputWithContext(ctx context.Context) Nfs3TargetResponsePtrOutput {
	return o
}

func (o Nfs3TargetResponsePtrOutput) Elem() Nfs3TargetResponseOutput {
	return o.ApplyT(func(v *Nfs3TargetResponse) Nfs3TargetResponse {
		if v != nil {
			return *v
		}
		var ret Nfs3TargetResponse
		return ret
	}).(Nfs3TargetResponseOutput)
}

func (o Nfs3TargetResponsePtrOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Nfs3TargetResponse) *string {
		if v == nil {
			return nil
		}
		return v.Target
	}).(pulumi.StringPtrOutput)
}

func (o Nfs3TargetResponsePtrOutput) UsageModel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Nfs3TargetResponse) *string {
		if v == nil {
			return nil
		}
		return v.UsageModel
	}).(pulumi.StringPtrOutput)
}

type UnknownTarget struct {
	UnknownMap map[string]string `pulumi:"unknownMap"`
}





type UnknownTargetInput interface {
	pulumi.Input

	ToUnknownTargetOutput() UnknownTargetOutput
	ToUnknownTargetOutputWithContext(context.Context) UnknownTargetOutput
}

type UnknownTargetArgs struct {
	UnknownMap pulumi.StringMapInput `pulumi:"unknownMap"`
}

func (UnknownTargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UnknownTarget)(nil)).Elem()
}

func (i UnknownTargetArgs) ToUnknownTargetOutput() UnknownTargetOutput {
	return i.ToUnknownTargetOutputWithContext(context.Background())
}

func (i UnknownTargetArgs) ToUnknownTargetOutputWithContext(ctx context.Context) UnknownTargetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UnknownTargetOutput)
}

func (i UnknownTargetArgs) ToUnknownTargetPtrOutput() UnknownTargetPtrOutput {
	return i.ToUnknownTargetPtrOutputWithContext(context.Background())
}

func (i UnknownTargetArgs) ToUnknownTargetPtrOutputWithContext(ctx context.Context) UnknownTargetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UnknownTargetOutput).ToUnknownTargetPtrOutputWithContext(ctx)
}









type UnknownTargetPtrInput interface {
	pulumi.Input

	ToUnknownTargetPtrOutput() UnknownTargetPtrOutput
	ToUnknownTargetPtrOutputWithContext(context.Context) UnknownTargetPtrOutput
}

type unknownTargetPtrType UnknownTargetArgs

func UnknownTargetPtr(v *UnknownTargetArgs) UnknownTargetPtrInput {
	return (*unknownTargetPtrType)(v)
}

func (*unknownTargetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UnknownTarget)(nil)).Elem()
}

func (i *unknownTargetPtrType) ToUnknownTargetPtrOutput() UnknownTargetPtrOutput {
	return i.ToUnknownTargetPtrOutputWithContext(context.Background())
}

func (i *unknownTargetPtrType) ToUnknownTargetPtrOutputWithContext(ctx context.Context) UnknownTargetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UnknownTargetPtrOutput)
}

type UnknownTargetOutput struct{ *pulumi.OutputState }

func (UnknownTargetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UnknownTarget)(nil)).Elem()
}

func (o UnknownTargetOutput) ToUnknownTargetOutput() UnknownTargetOutput {
	return o
}

func (o UnknownTargetOutput) ToUnknownTargetOutputWithContext(ctx context.Context) UnknownTargetOutput {
	return o
}

func (o UnknownTargetOutput) ToUnknownTargetPtrOutput() UnknownTargetPtrOutput {
	return o.ToUnknownTargetPtrOutputWithContext(context.Background())
}

func (o UnknownTargetOutput) ToUnknownTargetPtrOutputWithContext(ctx context.Context) UnknownTargetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UnknownTarget) *UnknownTarget {
		return &v
	}).(UnknownTargetPtrOutput)
}

func (o UnknownTargetOutput) UnknownMap() pulumi.StringMapOutput {
	return o.ApplyT(func(v UnknownTarget) map[string]string { return v.UnknownMap }).(pulumi.StringMapOutput)
}

type UnknownTargetPtrOutput struct{ *pulumi.OutputState }

func (UnknownTargetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UnknownTarget)(nil)).Elem()
}

func (o UnknownTargetPtrOutput) ToUnknownTargetPtrOutput() UnknownTargetPtrOutput {
	return o
}

func (o UnknownTargetPtrOutput) ToUnknownTargetPtrOutputWithContext(ctx context.Context) UnknownTargetPtrOutput {
	return o
}

func (o UnknownTargetPtrOutput) Elem() UnknownTargetOutput {
	return o.ApplyT(func(v *UnknownTarget) UnknownTarget {
		if v != nil {
			return *v
		}
		var ret UnknownTarget
		return ret
	}).(UnknownTargetOutput)
}

func (o UnknownTargetPtrOutput) UnknownMap() pulumi.StringMapOutput {
	return o.ApplyT(func(v *UnknownTarget) map[string]string {
		if v == nil {
			return nil
		}
		return v.UnknownMap
	}).(pulumi.StringMapOutput)
}

type UnknownTargetResponse struct {
	UnknownMap map[string]string `pulumi:"unknownMap"`
}





type UnknownTargetResponseInput interface {
	pulumi.Input

	ToUnknownTargetResponseOutput() UnknownTargetResponseOutput
	ToUnknownTargetResponseOutputWithContext(context.Context) UnknownTargetResponseOutput
}

type UnknownTargetResponseArgs struct {
	UnknownMap pulumi.StringMapInput `pulumi:"unknownMap"`
}

func (UnknownTargetResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UnknownTargetResponse)(nil)).Elem()
}

func (i UnknownTargetResponseArgs) ToUnknownTargetResponseOutput() UnknownTargetResponseOutput {
	return i.ToUnknownTargetResponseOutputWithContext(context.Background())
}

func (i UnknownTargetResponseArgs) ToUnknownTargetResponseOutputWithContext(ctx context.Context) UnknownTargetResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UnknownTargetResponseOutput)
}

func (i UnknownTargetResponseArgs) ToUnknownTargetResponsePtrOutput() UnknownTargetResponsePtrOutput {
	return i.ToUnknownTargetResponsePtrOutputWithContext(context.Background())
}

func (i UnknownTargetResponseArgs) ToUnknownTargetResponsePtrOutputWithContext(ctx context.Context) UnknownTargetResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UnknownTargetResponseOutput).ToUnknownTargetResponsePtrOutputWithContext(ctx)
}









type UnknownTargetResponsePtrInput interface {
	pulumi.Input

	ToUnknownTargetResponsePtrOutput() UnknownTargetResponsePtrOutput
	ToUnknownTargetResponsePtrOutputWithContext(context.Context) UnknownTargetResponsePtrOutput
}

type unknownTargetResponsePtrType UnknownTargetResponseArgs

func UnknownTargetResponsePtr(v *UnknownTargetResponseArgs) UnknownTargetResponsePtrInput {
	return (*unknownTargetResponsePtrType)(v)
}

func (*unknownTargetResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UnknownTargetResponse)(nil)).Elem()
}

func (i *unknownTargetResponsePtrType) ToUnknownTargetResponsePtrOutput() UnknownTargetResponsePtrOutput {
	return i.ToUnknownTargetResponsePtrOutputWithContext(context.Background())
}

func (i *unknownTargetResponsePtrType) ToUnknownTargetResponsePtrOutputWithContext(ctx context.Context) UnknownTargetResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UnknownTargetResponsePtrOutput)
}

type UnknownTargetResponseOutput struct{ *pulumi.OutputState }

func (UnknownTargetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UnknownTargetResponse)(nil)).Elem()
}

func (o UnknownTargetResponseOutput) ToUnknownTargetResponseOutput() UnknownTargetResponseOutput {
	return o
}

func (o UnknownTargetResponseOutput) ToUnknownTargetResponseOutputWithContext(ctx context.Context) UnknownTargetResponseOutput {
	return o
}

func (o UnknownTargetResponseOutput) ToUnknownTargetResponsePtrOutput() UnknownTargetResponsePtrOutput {
	return o.ToUnknownTargetResponsePtrOutputWithContext(context.Background())
}

func (o UnknownTargetResponseOutput) ToUnknownTargetResponsePtrOutputWithContext(ctx context.Context) UnknownTargetResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UnknownTargetResponse) *UnknownTargetResponse {
		return &v
	}).(UnknownTargetResponsePtrOutput)
}

func (o UnknownTargetResponseOutput) UnknownMap() pulumi.StringMapOutput {
	return o.ApplyT(func(v UnknownTargetResponse) map[string]string { return v.UnknownMap }).(pulumi.StringMapOutput)
}

type UnknownTargetResponsePtrOutput struct{ *pulumi.OutputState }

func (UnknownTargetResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UnknownTargetResponse)(nil)).Elem()
}

func (o UnknownTargetResponsePtrOutput) ToUnknownTargetResponsePtrOutput() UnknownTargetResponsePtrOutput {
	return o
}

func (o UnknownTargetResponsePtrOutput) ToUnknownTargetResponsePtrOutputWithContext(ctx context.Context) UnknownTargetResponsePtrOutput {
	return o
}

func (o UnknownTargetResponsePtrOutput) Elem() UnknownTargetResponseOutput {
	return o.ApplyT(func(v *UnknownTargetResponse) UnknownTargetResponse {
		if v != nil {
			return *v
		}
		var ret UnknownTargetResponse
		return ret
	}).(UnknownTargetResponseOutput)
}

func (o UnknownTargetResponsePtrOutput) UnknownMap() pulumi.StringMapOutput {
	return o.ApplyT(func(v *UnknownTargetResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.UnknownMap
	}).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(CacheHealthResponseOutput{})
	pulumi.RegisterOutputType(CacheHealthResponsePtrOutput{})
	pulumi.RegisterOutputType(CacheResponseSkuOutput{})
	pulumi.RegisterOutputType(CacheResponseSkuPtrOutput{})
	pulumi.RegisterOutputType(CacheSkuOutput{})
	pulumi.RegisterOutputType(CacheSkuPtrOutput{})
	pulumi.RegisterOutputType(CacheUpgradeStatusResponseOutput{})
	pulumi.RegisterOutputType(CacheUpgradeStatusResponsePtrOutput{})
	pulumi.RegisterOutputType(ClfsTargetOutput{})
	pulumi.RegisterOutputType(ClfsTargetPtrOutput{})
	pulumi.RegisterOutputType(ClfsTargetResponseOutput{})
	pulumi.RegisterOutputType(ClfsTargetResponsePtrOutput{})
	pulumi.RegisterOutputType(NamespaceJunctionOutput{})
	pulumi.RegisterOutputType(NamespaceJunctionArrayOutput{})
	pulumi.RegisterOutputType(NamespaceJunctionResponseOutput{})
	pulumi.RegisterOutputType(NamespaceJunctionResponseArrayOutput{})
	pulumi.RegisterOutputType(Nfs3TargetOutput{})
	pulumi.RegisterOutputType(Nfs3TargetPtrOutput{})
	pulumi.RegisterOutputType(Nfs3TargetResponseOutput{})
	pulumi.RegisterOutputType(Nfs3TargetResponsePtrOutput{})
	pulumi.RegisterOutputType(UnknownTargetOutput{})
	pulumi.RegisterOutputType(UnknownTargetPtrOutput{})
	pulumi.RegisterOutputType(UnknownTargetResponseOutput{})
	pulumi.RegisterOutputType(UnknownTargetResponsePtrOutput{})
}
