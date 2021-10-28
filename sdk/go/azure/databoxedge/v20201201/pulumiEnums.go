


package v20201201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AccountType string

const (
	AccountTypeGeneralPurposeStorage = AccountType("GeneralPurposeStorage")
	AccountTypeBlobStorage           = AccountType("BlobStorage")
)

func (AccountType) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountType)(nil)).Elem()
}

func (e AccountType) ToAccountTypeOutput() AccountTypeOutput {
	return pulumi.ToOutput(e).(AccountTypeOutput)
}

func (e AccountType) ToAccountTypeOutputWithContext(ctx context.Context) AccountTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AccountTypeOutput)
}

func (e AccountType) ToAccountTypePtrOutput() AccountTypePtrOutput {
	return e.ToAccountTypePtrOutputWithContext(context.Background())
}

func (e AccountType) ToAccountTypePtrOutputWithContext(ctx context.Context) AccountTypePtrOutput {
	return AccountType(e).ToAccountTypeOutputWithContext(ctx).ToAccountTypePtrOutputWithContext(ctx)
}

func (e AccountType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccountType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccountType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AccountType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AccountTypeOutput struct{ *pulumi.OutputState }

func (AccountTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountType)(nil)).Elem()
}

func (o AccountTypeOutput) ToAccountTypeOutput() AccountTypeOutput {
	return o
}

func (o AccountTypeOutput) ToAccountTypeOutputWithContext(ctx context.Context) AccountTypeOutput {
	return o
}

func (o AccountTypeOutput) ToAccountTypePtrOutput() AccountTypePtrOutput {
	return o.ToAccountTypePtrOutputWithContext(context.Background())
}

func (o AccountTypeOutput) ToAccountTypePtrOutputWithContext(ctx context.Context) AccountTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AccountType) *AccountType {
		return &v
	}).(AccountTypePtrOutput)
}

func (o AccountTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AccountTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccountType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AccountTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccountTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccountType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AccountTypePtrOutput struct{ *pulumi.OutputState }

func (AccountTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountType)(nil)).Elem()
}

func (o AccountTypePtrOutput) ToAccountTypePtrOutput() AccountTypePtrOutput {
	return o
}

func (o AccountTypePtrOutput) ToAccountTypePtrOutputWithContext(ctx context.Context) AccountTypePtrOutput {
	return o
}

func (o AccountTypePtrOutput) Elem() AccountTypeOutput {
	return o.ApplyT(func(v *AccountType) AccountType {
		if v != nil {
			return *v
		}
		var ret AccountType
		return ret
	}).(AccountTypeOutput)
}

func (o AccountTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccountTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AccountType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AccountTypeInput interface {
	pulumi.Input

	ToAccountTypeOutput() AccountTypeOutput
	ToAccountTypeOutputWithContext(context.Context) AccountTypeOutput
}

var accountTypePtrType = reflect.TypeOf((**AccountType)(nil)).Elem()

type AccountTypePtrInput interface {
	pulumi.Input

	ToAccountTypePtrOutput() AccountTypePtrOutput
	ToAccountTypePtrOutputWithContext(context.Context) AccountTypePtrOutput
}

type accountTypePtr string

func AccountTypePtr(v string) AccountTypePtrInput {
	return (*accountTypePtr)(&v)
}

func (*accountTypePtr) ElementType() reflect.Type {
	return accountTypePtrType
}

func (in *accountTypePtr) ToAccountTypePtrOutput() AccountTypePtrOutput {
	return pulumi.ToOutput(in).(AccountTypePtrOutput)
}

func (in *accountTypePtr) ToAccountTypePtrOutputWithContext(ctx context.Context) AccountTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AccountTypePtrOutput)
}

type AddonType string

const (
	AddonTypeIotEdge          = AddonType("IotEdge")
	AddonTypeArcForKubernetes = AddonType("ArcForKubernetes")
)

func (AddonType) ElementType() reflect.Type {
	return reflect.TypeOf((*AddonType)(nil)).Elem()
}

func (e AddonType) ToAddonTypeOutput() AddonTypeOutput {
	return pulumi.ToOutput(e).(AddonTypeOutput)
}

func (e AddonType) ToAddonTypeOutputWithContext(ctx context.Context) AddonTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AddonTypeOutput)
}

func (e AddonType) ToAddonTypePtrOutput() AddonTypePtrOutput {
	return e.ToAddonTypePtrOutputWithContext(context.Background())
}

func (e AddonType) ToAddonTypePtrOutputWithContext(ctx context.Context) AddonTypePtrOutput {
	return AddonType(e).ToAddonTypeOutputWithContext(ctx).ToAddonTypePtrOutputWithContext(ctx)
}

func (e AddonType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AddonType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AddonType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AddonType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AddonTypeOutput struct{ *pulumi.OutputState }

func (AddonTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AddonType)(nil)).Elem()
}

func (o AddonTypeOutput) ToAddonTypeOutput() AddonTypeOutput {
	return o
}

func (o AddonTypeOutput) ToAddonTypeOutputWithContext(ctx context.Context) AddonTypeOutput {
	return o
}

func (o AddonTypeOutput) ToAddonTypePtrOutput() AddonTypePtrOutput {
	return o.ToAddonTypePtrOutputWithContext(context.Background())
}

func (o AddonTypeOutput) ToAddonTypePtrOutputWithContext(ctx context.Context) AddonTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AddonType) *AddonType {
		return &v
	}).(AddonTypePtrOutput)
}

func (o AddonTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AddonTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AddonType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AddonTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AddonTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AddonType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AddonTypePtrOutput struct{ *pulumi.OutputState }

func (AddonTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AddonType)(nil)).Elem()
}

func (o AddonTypePtrOutput) ToAddonTypePtrOutput() AddonTypePtrOutput {
	return o
}

func (o AddonTypePtrOutput) ToAddonTypePtrOutputWithContext(ctx context.Context) AddonTypePtrOutput {
	return o
}

func (o AddonTypePtrOutput) Elem() AddonTypeOutput {
	return o.ApplyT(func(v *AddonType) AddonType {
		if v != nil {
			return *v
		}
		var ret AddonType
		return ret
	}).(AddonTypeOutput)
}

func (o AddonTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AddonTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AddonType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AddonTypeInput interface {
	pulumi.Input

	ToAddonTypeOutput() AddonTypeOutput
	ToAddonTypeOutputWithContext(context.Context) AddonTypeOutput
}

var addonTypePtrType = reflect.TypeOf((**AddonType)(nil)).Elem()

type AddonTypePtrInput interface {
	pulumi.Input

	ToAddonTypePtrOutput() AddonTypePtrOutput
	ToAddonTypePtrOutputWithContext(context.Context) AddonTypePtrOutput
}

type addonTypePtr string

func AddonTypePtr(v string) AddonTypePtrInput {
	return (*addonTypePtr)(&v)
}

func (*addonTypePtr) ElementType() reflect.Type {
	return addonTypePtrType
}

func (in *addonTypePtr) ToAddonTypePtrOutput() AddonTypePtrOutput {
	return pulumi.ToOutput(in).(AddonTypePtrOutput)
}

func (in *addonTypePtr) ToAddonTypePtrOutputWithContext(ctx context.Context) AddonTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AddonTypePtrOutput)
}

type AzureContainerDataFormat string

const (
	AzureContainerDataFormatBlockBlob = AzureContainerDataFormat("BlockBlob")
	AzureContainerDataFormatPageBlob  = AzureContainerDataFormat("PageBlob")
	AzureContainerDataFormatAzureFile = AzureContainerDataFormat("AzureFile")
)

func (AzureContainerDataFormat) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureContainerDataFormat)(nil)).Elem()
}

func (e AzureContainerDataFormat) ToAzureContainerDataFormatOutput() AzureContainerDataFormatOutput {
	return pulumi.ToOutput(e).(AzureContainerDataFormatOutput)
}

func (e AzureContainerDataFormat) ToAzureContainerDataFormatOutputWithContext(ctx context.Context) AzureContainerDataFormatOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AzureContainerDataFormatOutput)
}

func (e AzureContainerDataFormat) ToAzureContainerDataFormatPtrOutput() AzureContainerDataFormatPtrOutput {
	return e.ToAzureContainerDataFormatPtrOutputWithContext(context.Background())
}

func (e AzureContainerDataFormat) ToAzureContainerDataFormatPtrOutputWithContext(ctx context.Context) AzureContainerDataFormatPtrOutput {
	return AzureContainerDataFormat(e).ToAzureContainerDataFormatOutputWithContext(ctx).ToAzureContainerDataFormatPtrOutputWithContext(ctx)
}

func (e AzureContainerDataFormat) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureContainerDataFormat) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureContainerDataFormat) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AzureContainerDataFormat) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AzureContainerDataFormatOutput struct{ *pulumi.OutputState }

func (AzureContainerDataFormatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureContainerDataFormat)(nil)).Elem()
}

func (o AzureContainerDataFormatOutput) ToAzureContainerDataFormatOutput() AzureContainerDataFormatOutput {
	return o
}

func (o AzureContainerDataFormatOutput) ToAzureContainerDataFormatOutputWithContext(ctx context.Context) AzureContainerDataFormatOutput {
	return o
}

func (o AzureContainerDataFormatOutput) ToAzureContainerDataFormatPtrOutput() AzureContainerDataFormatPtrOutput {
	return o.ToAzureContainerDataFormatPtrOutputWithContext(context.Background())
}

func (o AzureContainerDataFormatOutput) ToAzureContainerDataFormatPtrOutputWithContext(ctx context.Context) AzureContainerDataFormatPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureContainerDataFormat) *AzureContainerDataFormat {
		return &v
	}).(AzureContainerDataFormatPtrOutput)
}

func (o AzureContainerDataFormatOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AzureContainerDataFormatOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureContainerDataFormat) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AzureContainerDataFormatOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureContainerDataFormatOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureContainerDataFormat) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AzureContainerDataFormatPtrOutput struct{ *pulumi.OutputState }

func (AzureContainerDataFormatPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureContainerDataFormat)(nil)).Elem()
}

func (o AzureContainerDataFormatPtrOutput) ToAzureContainerDataFormatPtrOutput() AzureContainerDataFormatPtrOutput {
	return o
}

func (o AzureContainerDataFormatPtrOutput) ToAzureContainerDataFormatPtrOutputWithContext(ctx context.Context) AzureContainerDataFormatPtrOutput {
	return o
}

func (o AzureContainerDataFormatPtrOutput) Elem() AzureContainerDataFormatOutput {
	return o.ApplyT(func(v *AzureContainerDataFormat) AzureContainerDataFormat {
		if v != nil {
			return *v
		}
		var ret AzureContainerDataFormat
		return ret
	}).(AzureContainerDataFormatOutput)
}

func (o AzureContainerDataFormatPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureContainerDataFormatPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AzureContainerDataFormat) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AzureContainerDataFormatInput interface {
	pulumi.Input

	ToAzureContainerDataFormatOutput() AzureContainerDataFormatOutput
	ToAzureContainerDataFormatOutputWithContext(context.Context) AzureContainerDataFormatOutput
}

var azureContainerDataFormatPtrType = reflect.TypeOf((**AzureContainerDataFormat)(nil)).Elem()

type AzureContainerDataFormatPtrInput interface {
	pulumi.Input

	ToAzureContainerDataFormatPtrOutput() AzureContainerDataFormatPtrOutput
	ToAzureContainerDataFormatPtrOutputWithContext(context.Context) AzureContainerDataFormatPtrOutput
}

type azureContainerDataFormatPtr string

func AzureContainerDataFormatPtr(v string) AzureContainerDataFormatPtrInput {
	return (*azureContainerDataFormatPtr)(&v)
}

func (*azureContainerDataFormatPtr) ElementType() reflect.Type {
	return azureContainerDataFormatPtrType
}

func (in *azureContainerDataFormatPtr) ToAzureContainerDataFormatPtrOutput() AzureContainerDataFormatPtrOutput {
	return pulumi.ToOutput(in).(AzureContainerDataFormatPtrOutput)
}

func (in *azureContainerDataFormatPtr) ToAzureContainerDataFormatPtrOutputWithContext(ctx context.Context) AzureContainerDataFormatPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AzureContainerDataFormatPtrOutput)
}

type ClientPermissionType string

const (
	ClientPermissionTypeNoAccess  = ClientPermissionType("NoAccess")
	ClientPermissionTypeReadOnly  = ClientPermissionType("ReadOnly")
	ClientPermissionTypeReadWrite = ClientPermissionType("ReadWrite")
)

func (ClientPermissionType) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientPermissionType)(nil)).Elem()
}

func (e ClientPermissionType) ToClientPermissionTypeOutput() ClientPermissionTypeOutput {
	return pulumi.ToOutput(e).(ClientPermissionTypeOutput)
}

func (e ClientPermissionType) ToClientPermissionTypeOutputWithContext(ctx context.Context) ClientPermissionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ClientPermissionTypeOutput)
}

func (e ClientPermissionType) ToClientPermissionTypePtrOutput() ClientPermissionTypePtrOutput {
	return e.ToClientPermissionTypePtrOutputWithContext(context.Background())
}

func (e ClientPermissionType) ToClientPermissionTypePtrOutputWithContext(ctx context.Context) ClientPermissionTypePtrOutput {
	return ClientPermissionType(e).ToClientPermissionTypeOutputWithContext(ctx).ToClientPermissionTypePtrOutputWithContext(ctx)
}

func (e ClientPermissionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ClientPermissionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ClientPermissionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ClientPermissionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ClientPermissionTypeOutput struct{ *pulumi.OutputState }

func (ClientPermissionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientPermissionType)(nil)).Elem()
}

func (o ClientPermissionTypeOutput) ToClientPermissionTypeOutput() ClientPermissionTypeOutput {
	return o
}

func (o ClientPermissionTypeOutput) ToClientPermissionTypeOutputWithContext(ctx context.Context) ClientPermissionTypeOutput {
	return o
}

func (o ClientPermissionTypeOutput) ToClientPermissionTypePtrOutput() ClientPermissionTypePtrOutput {
	return o.ToClientPermissionTypePtrOutputWithContext(context.Background())
}

func (o ClientPermissionTypeOutput) ToClientPermissionTypePtrOutputWithContext(ctx context.Context) ClientPermissionTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClientPermissionType) *ClientPermissionType {
		return &v
	}).(ClientPermissionTypePtrOutput)
}

func (o ClientPermissionTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ClientPermissionTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ClientPermissionType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ClientPermissionTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ClientPermissionTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ClientPermissionType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ClientPermissionTypePtrOutput struct{ *pulumi.OutputState }

func (ClientPermissionTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientPermissionType)(nil)).Elem()
}

func (o ClientPermissionTypePtrOutput) ToClientPermissionTypePtrOutput() ClientPermissionTypePtrOutput {
	return o
}

func (o ClientPermissionTypePtrOutput) ToClientPermissionTypePtrOutputWithContext(ctx context.Context) ClientPermissionTypePtrOutput {
	return o
}

func (o ClientPermissionTypePtrOutput) Elem() ClientPermissionTypeOutput {
	return o.ApplyT(func(v *ClientPermissionType) ClientPermissionType {
		if v != nil {
			return *v
		}
		var ret ClientPermissionType
		return ret
	}).(ClientPermissionTypeOutput)
}

func (o ClientPermissionTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ClientPermissionTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ClientPermissionType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ClientPermissionTypeInput interface {
	pulumi.Input

	ToClientPermissionTypeOutput() ClientPermissionTypeOutput
	ToClientPermissionTypeOutputWithContext(context.Context) ClientPermissionTypeOutput
}

var clientPermissionTypePtrType = reflect.TypeOf((**ClientPermissionType)(nil)).Elem()

type ClientPermissionTypePtrInput interface {
	pulumi.Input

	ToClientPermissionTypePtrOutput() ClientPermissionTypePtrOutput
	ToClientPermissionTypePtrOutputWithContext(context.Context) ClientPermissionTypePtrOutput
}

type clientPermissionTypePtr string

func ClientPermissionTypePtr(v string) ClientPermissionTypePtrInput {
	return (*clientPermissionTypePtr)(&v)
}

func (*clientPermissionTypePtr) ElementType() reflect.Type {
	return clientPermissionTypePtrType
}

func (in *clientPermissionTypePtr) ToClientPermissionTypePtrOutput() ClientPermissionTypePtrOutput {
	return pulumi.ToOutput(in).(ClientPermissionTypePtrOutput)
}

func (in *clientPermissionTypePtr) ToClientPermissionTypePtrOutputWithContext(ctx context.Context) ClientPermissionTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ClientPermissionTypePtrOutput)
}

type DataBoxEdgeDeviceStatus string

const (
	DataBoxEdgeDeviceStatusReadyToSetup          = DataBoxEdgeDeviceStatus("ReadyToSetup")
	DataBoxEdgeDeviceStatusOnline                = DataBoxEdgeDeviceStatus("Online")
	DataBoxEdgeDeviceStatusOffline               = DataBoxEdgeDeviceStatus("Offline")
	DataBoxEdgeDeviceStatusNeedsAttention        = DataBoxEdgeDeviceStatus("NeedsAttention")
	DataBoxEdgeDeviceStatusDisconnected          = DataBoxEdgeDeviceStatus("Disconnected")
	DataBoxEdgeDeviceStatusPartiallyDisconnected = DataBoxEdgeDeviceStatus("PartiallyDisconnected")
	DataBoxEdgeDeviceStatusMaintenance           = DataBoxEdgeDeviceStatus("Maintenance")
)

func (DataBoxEdgeDeviceStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxEdgeDeviceStatus)(nil)).Elem()
}

func (e DataBoxEdgeDeviceStatus) ToDataBoxEdgeDeviceStatusOutput() DataBoxEdgeDeviceStatusOutput {
	return pulumi.ToOutput(e).(DataBoxEdgeDeviceStatusOutput)
}

func (e DataBoxEdgeDeviceStatus) ToDataBoxEdgeDeviceStatusOutputWithContext(ctx context.Context) DataBoxEdgeDeviceStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DataBoxEdgeDeviceStatusOutput)
}

func (e DataBoxEdgeDeviceStatus) ToDataBoxEdgeDeviceStatusPtrOutput() DataBoxEdgeDeviceStatusPtrOutput {
	return e.ToDataBoxEdgeDeviceStatusPtrOutputWithContext(context.Background())
}

func (e DataBoxEdgeDeviceStatus) ToDataBoxEdgeDeviceStatusPtrOutputWithContext(ctx context.Context) DataBoxEdgeDeviceStatusPtrOutput {
	return DataBoxEdgeDeviceStatus(e).ToDataBoxEdgeDeviceStatusOutputWithContext(ctx).ToDataBoxEdgeDeviceStatusPtrOutputWithContext(ctx)
}

func (e DataBoxEdgeDeviceStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataBoxEdgeDeviceStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataBoxEdgeDeviceStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DataBoxEdgeDeviceStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DataBoxEdgeDeviceStatusOutput struct{ *pulumi.OutputState }

func (DataBoxEdgeDeviceStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxEdgeDeviceStatus)(nil)).Elem()
}

func (o DataBoxEdgeDeviceStatusOutput) ToDataBoxEdgeDeviceStatusOutput() DataBoxEdgeDeviceStatusOutput {
	return o
}

func (o DataBoxEdgeDeviceStatusOutput) ToDataBoxEdgeDeviceStatusOutputWithContext(ctx context.Context) DataBoxEdgeDeviceStatusOutput {
	return o
}

func (o DataBoxEdgeDeviceStatusOutput) ToDataBoxEdgeDeviceStatusPtrOutput() DataBoxEdgeDeviceStatusPtrOutput {
	return o.ToDataBoxEdgeDeviceStatusPtrOutputWithContext(context.Background())
}

func (o DataBoxEdgeDeviceStatusOutput) ToDataBoxEdgeDeviceStatusPtrOutputWithContext(ctx context.Context) DataBoxEdgeDeviceStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataBoxEdgeDeviceStatus) *DataBoxEdgeDeviceStatus {
		return &v
	}).(DataBoxEdgeDeviceStatusPtrOutput)
}

func (o DataBoxEdgeDeviceStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DataBoxEdgeDeviceStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DataBoxEdgeDeviceStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DataBoxEdgeDeviceStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DataBoxEdgeDeviceStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DataBoxEdgeDeviceStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DataBoxEdgeDeviceStatusPtrOutput struct{ *pulumi.OutputState }

func (DataBoxEdgeDeviceStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataBoxEdgeDeviceStatus)(nil)).Elem()
}

func (o DataBoxEdgeDeviceStatusPtrOutput) ToDataBoxEdgeDeviceStatusPtrOutput() DataBoxEdgeDeviceStatusPtrOutput {
	return o
}

func (o DataBoxEdgeDeviceStatusPtrOutput) ToDataBoxEdgeDeviceStatusPtrOutputWithContext(ctx context.Context) DataBoxEdgeDeviceStatusPtrOutput {
	return o
}

func (o DataBoxEdgeDeviceStatusPtrOutput) Elem() DataBoxEdgeDeviceStatusOutput {
	return o.ApplyT(func(v *DataBoxEdgeDeviceStatus) DataBoxEdgeDeviceStatus {
		if v != nil {
			return *v
		}
		var ret DataBoxEdgeDeviceStatus
		return ret
	}).(DataBoxEdgeDeviceStatusOutput)
}

func (o DataBoxEdgeDeviceStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DataBoxEdgeDeviceStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DataBoxEdgeDeviceStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DataBoxEdgeDeviceStatusInput interface {
	pulumi.Input

	ToDataBoxEdgeDeviceStatusOutput() DataBoxEdgeDeviceStatusOutput
	ToDataBoxEdgeDeviceStatusOutputWithContext(context.Context) DataBoxEdgeDeviceStatusOutput
}

var dataBoxEdgeDeviceStatusPtrType = reflect.TypeOf((**DataBoxEdgeDeviceStatus)(nil)).Elem()

type DataBoxEdgeDeviceStatusPtrInput interface {
	pulumi.Input

	ToDataBoxEdgeDeviceStatusPtrOutput() DataBoxEdgeDeviceStatusPtrOutput
	ToDataBoxEdgeDeviceStatusPtrOutputWithContext(context.Context) DataBoxEdgeDeviceStatusPtrOutput
}

type dataBoxEdgeDeviceStatusPtr string

func DataBoxEdgeDeviceStatusPtr(v string) DataBoxEdgeDeviceStatusPtrInput {
	return (*dataBoxEdgeDeviceStatusPtr)(&v)
}

func (*dataBoxEdgeDeviceStatusPtr) ElementType() reflect.Type {
	return dataBoxEdgeDeviceStatusPtrType
}

func (in *dataBoxEdgeDeviceStatusPtr) ToDataBoxEdgeDeviceStatusPtrOutput() DataBoxEdgeDeviceStatusPtrOutput {
	return pulumi.ToOutput(in).(DataBoxEdgeDeviceStatusPtrOutput)
}

func (in *dataBoxEdgeDeviceStatusPtr) ToDataBoxEdgeDeviceStatusPtrOutputWithContext(ctx context.Context) DataBoxEdgeDeviceStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DataBoxEdgeDeviceStatusPtrOutput)
}

type DataPolicy string

const (
	DataPolicyCloud = DataPolicy("Cloud")
	DataPolicyLocal = DataPolicy("Local")
)

func (DataPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*DataPolicy)(nil)).Elem()
}

func (e DataPolicy) ToDataPolicyOutput() DataPolicyOutput {
	return pulumi.ToOutput(e).(DataPolicyOutput)
}

func (e DataPolicy) ToDataPolicyOutputWithContext(ctx context.Context) DataPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DataPolicyOutput)
}

func (e DataPolicy) ToDataPolicyPtrOutput() DataPolicyPtrOutput {
	return e.ToDataPolicyPtrOutputWithContext(context.Background())
}

func (e DataPolicy) ToDataPolicyPtrOutputWithContext(ctx context.Context) DataPolicyPtrOutput {
	return DataPolicy(e).ToDataPolicyOutputWithContext(ctx).ToDataPolicyPtrOutputWithContext(ctx)
}

func (e DataPolicy) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataPolicy) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataPolicy) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DataPolicy) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DataPolicyOutput struct{ *pulumi.OutputState }

func (DataPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataPolicy)(nil)).Elem()
}

func (o DataPolicyOutput) ToDataPolicyOutput() DataPolicyOutput {
	return o
}

func (o DataPolicyOutput) ToDataPolicyOutputWithContext(ctx context.Context) DataPolicyOutput {
	return o
}

func (o DataPolicyOutput) ToDataPolicyPtrOutput() DataPolicyPtrOutput {
	return o.ToDataPolicyPtrOutputWithContext(context.Background())
}

func (o DataPolicyOutput) ToDataPolicyPtrOutputWithContext(ctx context.Context) DataPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataPolicy) *DataPolicy {
		return &v
	}).(DataPolicyPtrOutput)
}

func (o DataPolicyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DataPolicyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DataPolicy) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DataPolicyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DataPolicyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DataPolicy) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DataPolicyPtrOutput struct{ *pulumi.OutputState }

func (DataPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataPolicy)(nil)).Elem()
}

func (o DataPolicyPtrOutput) ToDataPolicyPtrOutput() DataPolicyPtrOutput {
	return o
}

func (o DataPolicyPtrOutput) ToDataPolicyPtrOutputWithContext(ctx context.Context) DataPolicyPtrOutput {
	return o
}

func (o DataPolicyPtrOutput) Elem() DataPolicyOutput {
	return o.ApplyT(func(v *DataPolicy) DataPolicy {
		if v != nil {
			return *v
		}
		var ret DataPolicy
		return ret
	}).(DataPolicyOutput)
}

func (o DataPolicyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DataPolicyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DataPolicy) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DataPolicyInput interface {
	pulumi.Input

	ToDataPolicyOutput() DataPolicyOutput
	ToDataPolicyOutputWithContext(context.Context) DataPolicyOutput
}

var dataPolicyPtrType = reflect.TypeOf((**DataPolicy)(nil)).Elem()

type DataPolicyPtrInput interface {
	pulumi.Input

	ToDataPolicyPtrOutput() DataPolicyPtrOutput
	ToDataPolicyPtrOutputWithContext(context.Context) DataPolicyPtrOutput
}

type dataPolicyPtr string

func DataPolicyPtr(v string) DataPolicyPtrInput {
	return (*dataPolicyPtr)(&v)
}

func (*dataPolicyPtr) ElementType() reflect.Type {
	return dataPolicyPtrType
}

func (in *dataPolicyPtr) ToDataPolicyPtrOutput() DataPolicyPtrOutput {
	return pulumi.ToOutput(in).(DataPolicyPtrOutput)
}

func (in *dataPolicyPtr) ToDataPolicyPtrOutputWithContext(ctx context.Context) DataPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DataPolicyPtrOutput)
}

type DayOfWeek string

const (
	DayOfWeekSunday    = DayOfWeek("Sunday")
	DayOfWeekMonday    = DayOfWeek("Monday")
	DayOfWeekTuesday   = DayOfWeek("Tuesday")
	DayOfWeekWednesday = DayOfWeek("Wednesday")
	DayOfWeekThursday  = DayOfWeek("Thursday")
	DayOfWeekFriday    = DayOfWeek("Friday")
	DayOfWeekSaturday  = DayOfWeek("Saturday")
)

func (DayOfWeek) ElementType() reflect.Type {
	return reflect.TypeOf((*DayOfWeek)(nil)).Elem()
}

func (e DayOfWeek) ToDayOfWeekOutput() DayOfWeekOutput {
	return pulumi.ToOutput(e).(DayOfWeekOutput)
}

func (e DayOfWeek) ToDayOfWeekOutputWithContext(ctx context.Context) DayOfWeekOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DayOfWeekOutput)
}

func (e DayOfWeek) ToDayOfWeekPtrOutput() DayOfWeekPtrOutput {
	return e.ToDayOfWeekPtrOutputWithContext(context.Background())
}

func (e DayOfWeek) ToDayOfWeekPtrOutputWithContext(ctx context.Context) DayOfWeekPtrOutput {
	return DayOfWeek(e).ToDayOfWeekOutputWithContext(ctx).ToDayOfWeekPtrOutputWithContext(ctx)
}

func (e DayOfWeek) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DayOfWeek) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DayOfWeek) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DayOfWeek) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DayOfWeekOutput struct{ *pulumi.OutputState }

func (DayOfWeekOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DayOfWeek)(nil)).Elem()
}

func (o DayOfWeekOutput) ToDayOfWeekOutput() DayOfWeekOutput {
	return o
}

func (o DayOfWeekOutput) ToDayOfWeekOutputWithContext(ctx context.Context) DayOfWeekOutput {
	return o
}

func (o DayOfWeekOutput) ToDayOfWeekPtrOutput() DayOfWeekPtrOutput {
	return o.ToDayOfWeekPtrOutputWithContext(context.Background())
}

func (o DayOfWeekOutput) ToDayOfWeekPtrOutputWithContext(ctx context.Context) DayOfWeekPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DayOfWeek) *DayOfWeek {
		return &v
	}).(DayOfWeekPtrOutput)
}

func (o DayOfWeekOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DayOfWeekOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DayOfWeek) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DayOfWeekOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DayOfWeekOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DayOfWeek) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DayOfWeekPtrOutput struct{ *pulumi.OutputState }

func (DayOfWeekPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DayOfWeek)(nil)).Elem()
}

func (o DayOfWeekPtrOutput) ToDayOfWeekPtrOutput() DayOfWeekPtrOutput {
	return o
}

func (o DayOfWeekPtrOutput) ToDayOfWeekPtrOutputWithContext(ctx context.Context) DayOfWeekPtrOutput {
	return o
}

func (o DayOfWeekPtrOutput) Elem() DayOfWeekOutput {
	return o.ApplyT(func(v *DayOfWeek) DayOfWeek {
		if v != nil {
			return *v
		}
		var ret DayOfWeek
		return ret
	}).(DayOfWeekOutput)
}

func (o DayOfWeekPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DayOfWeekPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DayOfWeek) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DayOfWeekInput interface {
	pulumi.Input

	ToDayOfWeekOutput() DayOfWeekOutput
	ToDayOfWeekOutputWithContext(context.Context) DayOfWeekOutput
}

var dayOfWeekPtrType = reflect.TypeOf((**DayOfWeek)(nil)).Elem()

type DayOfWeekPtrInput interface {
	pulumi.Input

	ToDayOfWeekPtrOutput() DayOfWeekPtrOutput
	ToDayOfWeekPtrOutputWithContext(context.Context) DayOfWeekPtrOutput
}

type dayOfWeekPtr string

func DayOfWeekPtr(v string) DayOfWeekPtrInput {
	return (*dayOfWeekPtr)(&v)
}

func (*dayOfWeekPtr) ElementType() reflect.Type {
	return dayOfWeekPtrType
}

func (in *dayOfWeekPtr) ToDayOfWeekPtrOutput() DayOfWeekPtrOutput {
	return pulumi.ToOutput(in).(DayOfWeekPtrOutput)
}

func (in *dayOfWeekPtr) ToDayOfWeekPtrOutputWithContext(ctx context.Context) DayOfWeekPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DayOfWeekPtrOutput)
}

type EncryptionAlgorithm string

const (
	EncryptionAlgorithmNone               = EncryptionAlgorithm("None")
	EncryptionAlgorithmAES256             = EncryptionAlgorithm("AES256")
	EncryptionAlgorithm_RSAES_PKCS1_v_1_5 = EncryptionAlgorithm("RSAES_PKCS1_v_1_5")
)

func (EncryptionAlgorithm) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionAlgorithm)(nil)).Elem()
}

func (e EncryptionAlgorithm) ToEncryptionAlgorithmOutput() EncryptionAlgorithmOutput {
	return pulumi.ToOutput(e).(EncryptionAlgorithmOutput)
}

func (e EncryptionAlgorithm) ToEncryptionAlgorithmOutputWithContext(ctx context.Context) EncryptionAlgorithmOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EncryptionAlgorithmOutput)
}

func (e EncryptionAlgorithm) ToEncryptionAlgorithmPtrOutput() EncryptionAlgorithmPtrOutput {
	return e.ToEncryptionAlgorithmPtrOutputWithContext(context.Background())
}

func (e EncryptionAlgorithm) ToEncryptionAlgorithmPtrOutputWithContext(ctx context.Context) EncryptionAlgorithmPtrOutput {
	return EncryptionAlgorithm(e).ToEncryptionAlgorithmOutputWithContext(ctx).ToEncryptionAlgorithmPtrOutputWithContext(ctx)
}

func (e EncryptionAlgorithm) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EncryptionAlgorithm) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EncryptionAlgorithm) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EncryptionAlgorithm) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EncryptionAlgorithmOutput struct{ *pulumi.OutputState }

func (EncryptionAlgorithmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionAlgorithm)(nil)).Elem()
}

func (o EncryptionAlgorithmOutput) ToEncryptionAlgorithmOutput() EncryptionAlgorithmOutput {
	return o
}

func (o EncryptionAlgorithmOutput) ToEncryptionAlgorithmOutputWithContext(ctx context.Context) EncryptionAlgorithmOutput {
	return o
}

func (o EncryptionAlgorithmOutput) ToEncryptionAlgorithmPtrOutput() EncryptionAlgorithmPtrOutput {
	return o.ToEncryptionAlgorithmPtrOutputWithContext(context.Background())
}

func (o EncryptionAlgorithmOutput) ToEncryptionAlgorithmPtrOutputWithContext(ctx context.Context) EncryptionAlgorithmPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionAlgorithm) *EncryptionAlgorithm {
		return &v
	}).(EncryptionAlgorithmPtrOutput)
}

func (o EncryptionAlgorithmOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EncryptionAlgorithmOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EncryptionAlgorithm) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EncryptionAlgorithmOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EncryptionAlgorithmOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EncryptionAlgorithm) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EncryptionAlgorithmPtrOutput struct{ *pulumi.OutputState }

func (EncryptionAlgorithmPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionAlgorithm)(nil)).Elem()
}

func (o EncryptionAlgorithmPtrOutput) ToEncryptionAlgorithmPtrOutput() EncryptionAlgorithmPtrOutput {
	return o
}

func (o EncryptionAlgorithmPtrOutput) ToEncryptionAlgorithmPtrOutputWithContext(ctx context.Context) EncryptionAlgorithmPtrOutput {
	return o
}

func (o EncryptionAlgorithmPtrOutput) Elem() EncryptionAlgorithmOutput {
	return o.ApplyT(func(v *EncryptionAlgorithm) EncryptionAlgorithm {
		if v != nil {
			return *v
		}
		var ret EncryptionAlgorithm
		return ret
	}).(EncryptionAlgorithmOutput)
}

func (o EncryptionAlgorithmPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EncryptionAlgorithmPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EncryptionAlgorithm) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EncryptionAlgorithmInput interface {
	pulumi.Input

	ToEncryptionAlgorithmOutput() EncryptionAlgorithmOutput
	ToEncryptionAlgorithmOutputWithContext(context.Context) EncryptionAlgorithmOutput
}

var encryptionAlgorithmPtrType = reflect.TypeOf((**EncryptionAlgorithm)(nil)).Elem()

type EncryptionAlgorithmPtrInput interface {
	pulumi.Input

	ToEncryptionAlgorithmPtrOutput() EncryptionAlgorithmPtrOutput
	ToEncryptionAlgorithmPtrOutputWithContext(context.Context) EncryptionAlgorithmPtrOutput
}

type encryptionAlgorithmPtr string

func EncryptionAlgorithmPtr(v string) EncryptionAlgorithmPtrInput {
	return (*encryptionAlgorithmPtr)(&v)
}

func (*encryptionAlgorithmPtr) ElementType() reflect.Type {
	return encryptionAlgorithmPtrType
}

func (in *encryptionAlgorithmPtr) ToEncryptionAlgorithmPtrOutput() EncryptionAlgorithmPtrOutput {
	return pulumi.ToOutput(in).(EncryptionAlgorithmPtrOutput)
}

func (in *encryptionAlgorithmPtr) ToEncryptionAlgorithmPtrOutputWithContext(ctx context.Context) EncryptionAlgorithmPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EncryptionAlgorithmPtrOutput)
}

type MonitoringStatus string

const (
	MonitoringStatusEnabled  = MonitoringStatus("Enabled")
	MonitoringStatusDisabled = MonitoringStatus("Disabled")
)

func (MonitoringStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitoringStatus)(nil)).Elem()
}

func (e MonitoringStatus) ToMonitoringStatusOutput() MonitoringStatusOutput {
	return pulumi.ToOutput(e).(MonitoringStatusOutput)
}

func (e MonitoringStatus) ToMonitoringStatusOutputWithContext(ctx context.Context) MonitoringStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(MonitoringStatusOutput)
}

func (e MonitoringStatus) ToMonitoringStatusPtrOutput() MonitoringStatusPtrOutput {
	return e.ToMonitoringStatusPtrOutputWithContext(context.Background())
}

func (e MonitoringStatus) ToMonitoringStatusPtrOutputWithContext(ctx context.Context) MonitoringStatusPtrOutput {
	return MonitoringStatus(e).ToMonitoringStatusOutputWithContext(ctx).ToMonitoringStatusPtrOutputWithContext(ctx)
}

func (e MonitoringStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e MonitoringStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e MonitoringStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e MonitoringStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type MonitoringStatusOutput struct{ *pulumi.OutputState }

func (MonitoringStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitoringStatus)(nil)).Elem()
}

func (o MonitoringStatusOutput) ToMonitoringStatusOutput() MonitoringStatusOutput {
	return o
}

func (o MonitoringStatusOutput) ToMonitoringStatusOutputWithContext(ctx context.Context) MonitoringStatusOutput {
	return o
}

func (o MonitoringStatusOutput) ToMonitoringStatusPtrOutput() MonitoringStatusPtrOutput {
	return o.ToMonitoringStatusPtrOutputWithContext(context.Background())
}

func (o MonitoringStatusOutput) ToMonitoringStatusPtrOutputWithContext(ctx context.Context) MonitoringStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MonitoringStatus) *MonitoringStatus {
		return &v
	}).(MonitoringStatusPtrOutput)
}

func (o MonitoringStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o MonitoringStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MonitoringStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o MonitoringStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MonitoringStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MonitoringStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type MonitoringStatusPtrOutput struct{ *pulumi.OutputState }

func (MonitoringStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitoringStatus)(nil)).Elem()
}

func (o MonitoringStatusPtrOutput) ToMonitoringStatusPtrOutput() MonitoringStatusPtrOutput {
	return o
}

func (o MonitoringStatusPtrOutput) ToMonitoringStatusPtrOutputWithContext(ctx context.Context) MonitoringStatusPtrOutput {
	return o
}

func (o MonitoringStatusPtrOutput) Elem() MonitoringStatusOutput {
	return o.ApplyT(func(v *MonitoringStatus) MonitoringStatus {
		if v != nil {
			return *v
		}
		var ret MonitoringStatus
		return ret
	}).(MonitoringStatusOutput)
}

func (o MonitoringStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MonitoringStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *MonitoringStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type MonitoringStatusInput interface {
	pulumi.Input

	ToMonitoringStatusOutput() MonitoringStatusOutput
	ToMonitoringStatusOutputWithContext(context.Context) MonitoringStatusOutput
}

var monitoringStatusPtrType = reflect.TypeOf((**MonitoringStatus)(nil)).Elem()

type MonitoringStatusPtrInput interface {
	pulumi.Input

	ToMonitoringStatusPtrOutput() MonitoringStatusPtrOutput
	ToMonitoringStatusPtrOutputWithContext(context.Context) MonitoringStatusPtrOutput
}

type monitoringStatusPtr string

func MonitoringStatusPtr(v string) MonitoringStatusPtrInput {
	return (*monitoringStatusPtr)(&v)
}

func (*monitoringStatusPtr) ElementType() reflect.Type {
	return monitoringStatusPtrType
}

func (in *monitoringStatusPtr) ToMonitoringStatusPtrOutput() MonitoringStatusPtrOutput {
	return pulumi.ToOutput(in).(MonitoringStatusPtrOutput)
}

func (in *monitoringStatusPtr) ToMonitoringStatusPtrOutputWithContext(ctx context.Context) MonitoringStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(MonitoringStatusPtrOutput)
}

type MsiIdentityType string

const (
	MsiIdentityTypeNone           = MsiIdentityType("None")
	MsiIdentityTypeSystemAssigned = MsiIdentityType("SystemAssigned")
	MsiIdentityTypeUserAssigned   = MsiIdentityType("UserAssigned")
)

func (MsiIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*MsiIdentityType)(nil)).Elem()
}

func (e MsiIdentityType) ToMsiIdentityTypeOutput() MsiIdentityTypeOutput {
	return pulumi.ToOutput(e).(MsiIdentityTypeOutput)
}

func (e MsiIdentityType) ToMsiIdentityTypeOutputWithContext(ctx context.Context) MsiIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(MsiIdentityTypeOutput)
}

func (e MsiIdentityType) ToMsiIdentityTypePtrOutput() MsiIdentityTypePtrOutput {
	return e.ToMsiIdentityTypePtrOutputWithContext(context.Background())
}

func (e MsiIdentityType) ToMsiIdentityTypePtrOutputWithContext(ctx context.Context) MsiIdentityTypePtrOutput {
	return MsiIdentityType(e).ToMsiIdentityTypeOutputWithContext(ctx).ToMsiIdentityTypePtrOutputWithContext(ctx)
}

func (e MsiIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e MsiIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e MsiIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e MsiIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type MsiIdentityTypeOutput struct{ *pulumi.OutputState }

func (MsiIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MsiIdentityType)(nil)).Elem()
}

func (o MsiIdentityTypeOutput) ToMsiIdentityTypeOutput() MsiIdentityTypeOutput {
	return o
}

func (o MsiIdentityTypeOutput) ToMsiIdentityTypeOutputWithContext(ctx context.Context) MsiIdentityTypeOutput {
	return o
}

func (o MsiIdentityTypeOutput) ToMsiIdentityTypePtrOutput() MsiIdentityTypePtrOutput {
	return o.ToMsiIdentityTypePtrOutputWithContext(context.Background())
}

func (o MsiIdentityTypeOutput) ToMsiIdentityTypePtrOutputWithContext(ctx context.Context) MsiIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MsiIdentityType) *MsiIdentityType {
		return &v
	}).(MsiIdentityTypePtrOutput)
}

func (o MsiIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o MsiIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MsiIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o MsiIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MsiIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MsiIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type MsiIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (MsiIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MsiIdentityType)(nil)).Elem()
}

func (o MsiIdentityTypePtrOutput) ToMsiIdentityTypePtrOutput() MsiIdentityTypePtrOutput {
	return o
}

func (o MsiIdentityTypePtrOutput) ToMsiIdentityTypePtrOutputWithContext(ctx context.Context) MsiIdentityTypePtrOutput {
	return o
}

func (o MsiIdentityTypePtrOutput) Elem() MsiIdentityTypeOutput {
	return o.ApplyT(func(v *MsiIdentityType) MsiIdentityType {
		if v != nil {
			return *v
		}
		var ret MsiIdentityType
		return ret
	}).(MsiIdentityTypeOutput)
}

func (o MsiIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MsiIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *MsiIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type MsiIdentityTypeInput interface {
	pulumi.Input

	ToMsiIdentityTypeOutput() MsiIdentityTypeOutput
	ToMsiIdentityTypeOutputWithContext(context.Context) MsiIdentityTypeOutput
}

var msiIdentityTypePtrType = reflect.TypeOf((**MsiIdentityType)(nil)).Elem()

type MsiIdentityTypePtrInput interface {
	pulumi.Input

	ToMsiIdentityTypePtrOutput() MsiIdentityTypePtrOutput
	ToMsiIdentityTypePtrOutputWithContext(context.Context) MsiIdentityTypePtrOutput
}

type msiIdentityTypePtr string

func MsiIdentityTypePtr(v string) MsiIdentityTypePtrInput {
	return (*msiIdentityTypePtr)(&v)
}

func (*msiIdentityTypePtr) ElementType() reflect.Type {
	return msiIdentityTypePtrType
}

func (in *msiIdentityTypePtr) ToMsiIdentityTypePtrOutput() MsiIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(MsiIdentityTypePtrOutput)
}

func (in *msiIdentityTypePtr) ToMsiIdentityTypePtrOutputWithContext(ctx context.Context) MsiIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(MsiIdentityTypePtrOutput)
}

type PlatformType string

const (
	PlatformTypeWindows = PlatformType("Windows")
	PlatformTypeLinux   = PlatformType("Linux")
)

func (PlatformType) ElementType() reflect.Type {
	return reflect.TypeOf((*PlatformType)(nil)).Elem()
}

func (e PlatformType) ToPlatformTypeOutput() PlatformTypeOutput {
	return pulumi.ToOutput(e).(PlatformTypeOutput)
}

func (e PlatformType) ToPlatformTypeOutputWithContext(ctx context.Context) PlatformTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PlatformTypeOutput)
}

func (e PlatformType) ToPlatformTypePtrOutput() PlatformTypePtrOutput {
	return e.ToPlatformTypePtrOutputWithContext(context.Background())
}

func (e PlatformType) ToPlatformTypePtrOutputWithContext(ctx context.Context) PlatformTypePtrOutput {
	return PlatformType(e).ToPlatformTypeOutputWithContext(ctx).ToPlatformTypePtrOutputWithContext(ctx)
}

func (e PlatformType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PlatformType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PlatformType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PlatformType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PlatformTypeOutput struct{ *pulumi.OutputState }

func (PlatformTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PlatformType)(nil)).Elem()
}

func (o PlatformTypeOutput) ToPlatformTypeOutput() PlatformTypeOutput {
	return o
}

func (o PlatformTypeOutput) ToPlatformTypeOutputWithContext(ctx context.Context) PlatformTypeOutput {
	return o
}

func (o PlatformTypeOutput) ToPlatformTypePtrOutput() PlatformTypePtrOutput {
	return o.ToPlatformTypePtrOutputWithContext(context.Background())
}

func (o PlatformTypeOutput) ToPlatformTypePtrOutputWithContext(ctx context.Context) PlatformTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PlatformType) *PlatformType {
		return &v
	}).(PlatformTypePtrOutput)
}

func (o PlatformTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PlatformTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PlatformType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PlatformTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PlatformTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PlatformType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PlatformTypePtrOutput struct{ *pulumi.OutputState }

func (PlatformTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PlatformType)(nil)).Elem()
}

func (o PlatformTypePtrOutput) ToPlatformTypePtrOutput() PlatformTypePtrOutput {
	return o
}

func (o PlatformTypePtrOutput) ToPlatformTypePtrOutputWithContext(ctx context.Context) PlatformTypePtrOutput {
	return o
}

func (o PlatformTypePtrOutput) Elem() PlatformTypeOutput {
	return o.ApplyT(func(v *PlatformType) PlatformType {
		if v != nil {
			return *v
		}
		var ret PlatformType
		return ret
	}).(PlatformTypeOutput)
}

func (o PlatformTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PlatformTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PlatformType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PlatformTypeInput interface {
	pulumi.Input

	ToPlatformTypeOutput() PlatformTypeOutput
	ToPlatformTypeOutputWithContext(context.Context) PlatformTypeOutput
}

var platformTypePtrType = reflect.TypeOf((**PlatformType)(nil)).Elem()

type PlatformTypePtrInput interface {
	pulumi.Input

	ToPlatformTypePtrOutput() PlatformTypePtrOutput
	ToPlatformTypePtrOutputWithContext(context.Context) PlatformTypePtrOutput
}

type platformTypePtr string

func PlatformTypePtr(v string) PlatformTypePtrInput {
	return (*platformTypePtr)(&v)
}

func (*platformTypePtr) ElementType() reflect.Type {
	return platformTypePtrType
}

func (in *platformTypePtr) ToPlatformTypePtrOutput() PlatformTypePtrOutput {
	return pulumi.ToOutput(in).(PlatformTypePtrOutput)
}

func (in *platformTypePtr) ToPlatformTypePtrOutputWithContext(ctx context.Context) PlatformTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PlatformTypePtrOutput)
}

type RoleStatus string

const (
	RoleStatusEnabled  = RoleStatus("Enabled")
	RoleStatusDisabled = RoleStatus("Disabled")
)

func (RoleStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*RoleStatus)(nil)).Elem()
}

func (e RoleStatus) ToRoleStatusOutput() RoleStatusOutput {
	return pulumi.ToOutput(e).(RoleStatusOutput)
}

func (e RoleStatus) ToRoleStatusOutputWithContext(ctx context.Context) RoleStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RoleStatusOutput)
}

func (e RoleStatus) ToRoleStatusPtrOutput() RoleStatusPtrOutput {
	return e.ToRoleStatusPtrOutputWithContext(context.Background())
}

func (e RoleStatus) ToRoleStatusPtrOutputWithContext(ctx context.Context) RoleStatusPtrOutput {
	return RoleStatus(e).ToRoleStatusOutputWithContext(ctx).ToRoleStatusPtrOutputWithContext(ctx)
}

func (e RoleStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RoleStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RoleStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RoleStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RoleStatusOutput struct{ *pulumi.OutputState }

func (RoleStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoleStatus)(nil)).Elem()
}

func (o RoleStatusOutput) ToRoleStatusOutput() RoleStatusOutput {
	return o
}

func (o RoleStatusOutput) ToRoleStatusOutputWithContext(ctx context.Context) RoleStatusOutput {
	return o
}

func (o RoleStatusOutput) ToRoleStatusPtrOutput() RoleStatusPtrOutput {
	return o.ToRoleStatusPtrOutputWithContext(context.Background())
}

func (o RoleStatusOutput) ToRoleStatusPtrOutputWithContext(ctx context.Context) RoleStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RoleStatus) *RoleStatus {
		return &v
	}).(RoleStatusPtrOutput)
}

func (o RoleStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RoleStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RoleStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RoleStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RoleStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RoleStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RoleStatusPtrOutput struct{ *pulumi.OutputState }

func (RoleStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleStatus)(nil)).Elem()
}

func (o RoleStatusPtrOutput) ToRoleStatusPtrOutput() RoleStatusPtrOutput {
	return o
}

func (o RoleStatusPtrOutput) ToRoleStatusPtrOutputWithContext(ctx context.Context) RoleStatusPtrOutput {
	return o
}

func (o RoleStatusPtrOutput) Elem() RoleStatusOutput {
	return o.ApplyT(func(v *RoleStatus) RoleStatus {
		if v != nil {
			return *v
		}
		var ret RoleStatus
		return ret
	}).(RoleStatusOutput)
}

func (o RoleStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RoleStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RoleStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type RoleStatusInput interface {
	pulumi.Input

	ToRoleStatusOutput() RoleStatusOutput
	ToRoleStatusOutputWithContext(context.Context) RoleStatusOutput
}

var roleStatusPtrType = reflect.TypeOf((**RoleStatus)(nil)).Elem()

type RoleStatusPtrInput interface {
	pulumi.Input

	ToRoleStatusPtrOutput() RoleStatusPtrOutput
	ToRoleStatusPtrOutputWithContext(context.Context) RoleStatusPtrOutput
}

type roleStatusPtr string

func RoleStatusPtr(v string) RoleStatusPtrInput {
	return (*roleStatusPtr)(&v)
}

func (*roleStatusPtr) ElementType() reflect.Type {
	return roleStatusPtrType
}

func (in *roleStatusPtr) ToRoleStatusPtrOutput() RoleStatusPtrOutput {
	return pulumi.ToOutput(in).(RoleStatusPtrOutput)
}

func (in *roleStatusPtr) ToRoleStatusPtrOutputWithContext(ctx context.Context) RoleStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RoleStatusPtrOutput)
}

type RoleTypes string

const (
	RoleTypesIOT                 = RoleTypes("IOT")
	RoleTypesASA                 = RoleTypes("ASA")
	RoleTypesFunctions           = RoleTypes("Functions")
	RoleTypesCognitive           = RoleTypes("Cognitive")
	RoleTypesMEC                 = RoleTypes("MEC")
	RoleTypesCloudEdgeManagement = RoleTypes("CloudEdgeManagement")
	RoleTypesKubernetes          = RoleTypes("Kubernetes")
)

func (RoleTypes) ElementType() reflect.Type {
	return reflect.TypeOf((*RoleTypes)(nil)).Elem()
}

func (e RoleTypes) ToRoleTypesOutput() RoleTypesOutput {
	return pulumi.ToOutput(e).(RoleTypesOutput)
}

func (e RoleTypes) ToRoleTypesOutputWithContext(ctx context.Context) RoleTypesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RoleTypesOutput)
}

func (e RoleTypes) ToRoleTypesPtrOutput() RoleTypesPtrOutput {
	return e.ToRoleTypesPtrOutputWithContext(context.Background())
}

func (e RoleTypes) ToRoleTypesPtrOutputWithContext(ctx context.Context) RoleTypesPtrOutput {
	return RoleTypes(e).ToRoleTypesOutputWithContext(ctx).ToRoleTypesPtrOutputWithContext(ctx)
}

func (e RoleTypes) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RoleTypes) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RoleTypes) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RoleTypes) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RoleTypesOutput struct{ *pulumi.OutputState }

func (RoleTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoleTypes)(nil)).Elem()
}

func (o RoleTypesOutput) ToRoleTypesOutput() RoleTypesOutput {
	return o
}

func (o RoleTypesOutput) ToRoleTypesOutputWithContext(ctx context.Context) RoleTypesOutput {
	return o
}

func (o RoleTypesOutput) ToRoleTypesPtrOutput() RoleTypesPtrOutput {
	return o.ToRoleTypesPtrOutputWithContext(context.Background())
}

func (o RoleTypesOutput) ToRoleTypesPtrOutputWithContext(ctx context.Context) RoleTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RoleTypes) *RoleTypes {
		return &v
	}).(RoleTypesPtrOutput)
}

func (o RoleTypesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RoleTypesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RoleTypes) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RoleTypesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RoleTypesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RoleTypes) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RoleTypesPtrOutput struct{ *pulumi.OutputState }

func (RoleTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleTypes)(nil)).Elem()
}

func (o RoleTypesPtrOutput) ToRoleTypesPtrOutput() RoleTypesPtrOutput {
	return o
}

func (o RoleTypesPtrOutput) ToRoleTypesPtrOutputWithContext(ctx context.Context) RoleTypesPtrOutput {
	return o
}

func (o RoleTypesPtrOutput) Elem() RoleTypesOutput {
	return o.ApplyT(func(v *RoleTypes) RoleTypes {
		if v != nil {
			return *v
		}
		var ret RoleTypes
		return ret
	}).(RoleTypesOutput)
}

func (o RoleTypesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RoleTypesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RoleTypes) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type RoleTypesInput interface {
	pulumi.Input

	ToRoleTypesOutput() RoleTypesOutput
	ToRoleTypesOutputWithContext(context.Context) RoleTypesOutput
}

var roleTypesPtrType = reflect.TypeOf((**RoleTypes)(nil)).Elem()

type RoleTypesPtrInput interface {
	pulumi.Input

	ToRoleTypesPtrOutput() RoleTypesPtrOutput
	ToRoleTypesPtrOutputWithContext(context.Context) RoleTypesPtrOutput
}

type roleTypesPtr string

func RoleTypesPtr(v string) RoleTypesPtrInput {
	return (*roleTypesPtr)(&v)
}

func (*roleTypesPtr) ElementType() reflect.Type {
	return roleTypesPtrType
}

func (in *roleTypesPtr) ToRoleTypesPtrOutput() RoleTypesPtrOutput {
	return pulumi.ToOutput(in).(RoleTypesPtrOutput)
}

func (in *roleTypesPtr) ToRoleTypesPtrOutputWithContext(ctx context.Context) RoleTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RoleTypesPtrOutput)
}

type SSLStatus string

const (
	SSLStatusEnabled  = SSLStatus("Enabled")
	SSLStatusDisabled = SSLStatus("Disabled")
)

func (SSLStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*SSLStatus)(nil)).Elem()
}

func (e SSLStatus) ToSSLStatusOutput() SSLStatusOutput {
	return pulumi.ToOutput(e).(SSLStatusOutput)
}

func (e SSLStatus) ToSSLStatusOutputWithContext(ctx context.Context) SSLStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SSLStatusOutput)
}

func (e SSLStatus) ToSSLStatusPtrOutput() SSLStatusPtrOutput {
	return e.ToSSLStatusPtrOutputWithContext(context.Background())
}

func (e SSLStatus) ToSSLStatusPtrOutputWithContext(ctx context.Context) SSLStatusPtrOutput {
	return SSLStatus(e).ToSSLStatusOutputWithContext(ctx).ToSSLStatusPtrOutputWithContext(ctx)
}

func (e SSLStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SSLStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SSLStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SSLStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SSLStatusOutput struct{ *pulumi.OutputState }

func (SSLStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SSLStatus)(nil)).Elem()
}

func (o SSLStatusOutput) ToSSLStatusOutput() SSLStatusOutput {
	return o
}

func (o SSLStatusOutput) ToSSLStatusOutputWithContext(ctx context.Context) SSLStatusOutput {
	return o
}

func (o SSLStatusOutput) ToSSLStatusPtrOutput() SSLStatusPtrOutput {
	return o.ToSSLStatusPtrOutputWithContext(context.Background())
}

func (o SSLStatusOutput) ToSSLStatusPtrOutputWithContext(ctx context.Context) SSLStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SSLStatus) *SSLStatus {
		return &v
	}).(SSLStatusPtrOutput)
}

func (o SSLStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SSLStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SSLStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SSLStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SSLStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SSLStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SSLStatusPtrOutput struct{ *pulumi.OutputState }

func (SSLStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SSLStatus)(nil)).Elem()
}

func (o SSLStatusPtrOutput) ToSSLStatusPtrOutput() SSLStatusPtrOutput {
	return o
}

func (o SSLStatusPtrOutput) ToSSLStatusPtrOutputWithContext(ctx context.Context) SSLStatusPtrOutput {
	return o
}

func (o SSLStatusPtrOutput) Elem() SSLStatusOutput {
	return o.ApplyT(func(v *SSLStatus) SSLStatus {
		if v != nil {
			return *v
		}
		var ret SSLStatus
		return ret
	}).(SSLStatusOutput)
}

func (o SSLStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SSLStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SSLStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SSLStatusInput interface {
	pulumi.Input

	ToSSLStatusOutput() SSLStatusOutput
	ToSSLStatusOutputWithContext(context.Context) SSLStatusOutput
}

var sslstatusPtrType = reflect.TypeOf((**SSLStatus)(nil)).Elem()

type SSLStatusPtrInput interface {
	pulumi.Input

	ToSSLStatusPtrOutput() SSLStatusPtrOutput
	ToSSLStatusPtrOutputWithContext(context.Context) SSLStatusPtrOutput
}

type sslstatusPtr string

func SSLStatusPtr(v string) SSLStatusPtrInput {
	return (*sslstatusPtr)(&v)
}

func (*sslstatusPtr) ElementType() reflect.Type {
	return sslstatusPtrType
}

func (in *sslstatusPtr) ToSSLStatusPtrOutput() SSLStatusPtrOutput {
	return pulumi.ToOutput(in).(SSLStatusPtrOutput)
}

func (in *sslstatusPtr) ToSSLStatusPtrOutputWithContext(ctx context.Context) SSLStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SSLStatusPtrOutput)
}

type ShareAccessProtocol string

const (
	ShareAccessProtocolSMB = ShareAccessProtocol("SMB")
	ShareAccessProtocolNFS = ShareAccessProtocol("NFS")
)

func (ShareAccessProtocol) ElementType() reflect.Type {
	return reflect.TypeOf((*ShareAccessProtocol)(nil)).Elem()
}

func (e ShareAccessProtocol) ToShareAccessProtocolOutput() ShareAccessProtocolOutput {
	return pulumi.ToOutput(e).(ShareAccessProtocolOutput)
}

func (e ShareAccessProtocol) ToShareAccessProtocolOutputWithContext(ctx context.Context) ShareAccessProtocolOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ShareAccessProtocolOutput)
}

func (e ShareAccessProtocol) ToShareAccessProtocolPtrOutput() ShareAccessProtocolPtrOutput {
	return e.ToShareAccessProtocolPtrOutputWithContext(context.Background())
}

func (e ShareAccessProtocol) ToShareAccessProtocolPtrOutputWithContext(ctx context.Context) ShareAccessProtocolPtrOutput {
	return ShareAccessProtocol(e).ToShareAccessProtocolOutputWithContext(ctx).ToShareAccessProtocolPtrOutputWithContext(ctx)
}

func (e ShareAccessProtocol) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ShareAccessProtocol) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ShareAccessProtocol) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ShareAccessProtocol) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ShareAccessProtocolOutput struct{ *pulumi.OutputState }

func (ShareAccessProtocolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ShareAccessProtocol)(nil)).Elem()
}

func (o ShareAccessProtocolOutput) ToShareAccessProtocolOutput() ShareAccessProtocolOutput {
	return o
}

func (o ShareAccessProtocolOutput) ToShareAccessProtocolOutputWithContext(ctx context.Context) ShareAccessProtocolOutput {
	return o
}

func (o ShareAccessProtocolOutput) ToShareAccessProtocolPtrOutput() ShareAccessProtocolPtrOutput {
	return o.ToShareAccessProtocolPtrOutputWithContext(context.Background())
}

func (o ShareAccessProtocolOutput) ToShareAccessProtocolPtrOutputWithContext(ctx context.Context) ShareAccessProtocolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ShareAccessProtocol) *ShareAccessProtocol {
		return &v
	}).(ShareAccessProtocolPtrOutput)
}

func (o ShareAccessProtocolOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ShareAccessProtocolOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ShareAccessProtocol) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ShareAccessProtocolOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ShareAccessProtocolOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ShareAccessProtocol) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ShareAccessProtocolPtrOutput struct{ *pulumi.OutputState }

func (ShareAccessProtocolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ShareAccessProtocol)(nil)).Elem()
}

func (o ShareAccessProtocolPtrOutput) ToShareAccessProtocolPtrOutput() ShareAccessProtocolPtrOutput {
	return o
}

func (o ShareAccessProtocolPtrOutput) ToShareAccessProtocolPtrOutputWithContext(ctx context.Context) ShareAccessProtocolPtrOutput {
	return o
}

func (o ShareAccessProtocolPtrOutput) Elem() ShareAccessProtocolOutput {
	return o.ApplyT(func(v *ShareAccessProtocol) ShareAccessProtocol {
		if v != nil {
			return *v
		}
		var ret ShareAccessProtocol
		return ret
	}).(ShareAccessProtocolOutput)
}

func (o ShareAccessProtocolPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ShareAccessProtocolPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ShareAccessProtocol) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ShareAccessProtocolInput interface {
	pulumi.Input

	ToShareAccessProtocolOutput() ShareAccessProtocolOutput
	ToShareAccessProtocolOutputWithContext(context.Context) ShareAccessProtocolOutput
}

var shareAccessProtocolPtrType = reflect.TypeOf((**ShareAccessProtocol)(nil)).Elem()

type ShareAccessProtocolPtrInput interface {
	pulumi.Input

	ToShareAccessProtocolPtrOutput() ShareAccessProtocolPtrOutput
	ToShareAccessProtocolPtrOutputWithContext(context.Context) ShareAccessProtocolPtrOutput
}

type shareAccessProtocolPtr string

func ShareAccessProtocolPtr(v string) ShareAccessProtocolPtrInput {
	return (*shareAccessProtocolPtr)(&v)
}

func (*shareAccessProtocolPtr) ElementType() reflect.Type {
	return shareAccessProtocolPtrType
}

func (in *shareAccessProtocolPtr) ToShareAccessProtocolPtrOutput() ShareAccessProtocolPtrOutput {
	return pulumi.ToOutput(in).(ShareAccessProtocolPtrOutput)
}

func (in *shareAccessProtocolPtr) ToShareAccessProtocolPtrOutputWithContext(ctx context.Context) ShareAccessProtocolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ShareAccessProtocolPtrOutput)
}

type ShareAccessType string

const (
	ShareAccessTypeChange = ShareAccessType("Change")
	ShareAccessTypeRead   = ShareAccessType("Read")
	ShareAccessTypeCustom = ShareAccessType("Custom")
)

func (ShareAccessType) ElementType() reflect.Type {
	return reflect.TypeOf((*ShareAccessType)(nil)).Elem()
}

func (e ShareAccessType) ToShareAccessTypeOutput() ShareAccessTypeOutput {
	return pulumi.ToOutput(e).(ShareAccessTypeOutput)
}

func (e ShareAccessType) ToShareAccessTypeOutputWithContext(ctx context.Context) ShareAccessTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ShareAccessTypeOutput)
}

func (e ShareAccessType) ToShareAccessTypePtrOutput() ShareAccessTypePtrOutput {
	return e.ToShareAccessTypePtrOutputWithContext(context.Background())
}

func (e ShareAccessType) ToShareAccessTypePtrOutputWithContext(ctx context.Context) ShareAccessTypePtrOutput {
	return ShareAccessType(e).ToShareAccessTypeOutputWithContext(ctx).ToShareAccessTypePtrOutputWithContext(ctx)
}

func (e ShareAccessType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ShareAccessType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ShareAccessType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ShareAccessType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ShareAccessTypeOutput struct{ *pulumi.OutputState }

func (ShareAccessTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ShareAccessType)(nil)).Elem()
}

func (o ShareAccessTypeOutput) ToShareAccessTypeOutput() ShareAccessTypeOutput {
	return o
}

func (o ShareAccessTypeOutput) ToShareAccessTypeOutputWithContext(ctx context.Context) ShareAccessTypeOutput {
	return o
}

func (o ShareAccessTypeOutput) ToShareAccessTypePtrOutput() ShareAccessTypePtrOutput {
	return o.ToShareAccessTypePtrOutputWithContext(context.Background())
}

func (o ShareAccessTypeOutput) ToShareAccessTypePtrOutputWithContext(ctx context.Context) ShareAccessTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ShareAccessType) *ShareAccessType {
		return &v
	}).(ShareAccessTypePtrOutput)
}

func (o ShareAccessTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ShareAccessTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ShareAccessType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ShareAccessTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ShareAccessTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ShareAccessType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ShareAccessTypePtrOutput struct{ *pulumi.OutputState }

func (ShareAccessTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ShareAccessType)(nil)).Elem()
}

func (o ShareAccessTypePtrOutput) ToShareAccessTypePtrOutput() ShareAccessTypePtrOutput {
	return o
}

func (o ShareAccessTypePtrOutput) ToShareAccessTypePtrOutputWithContext(ctx context.Context) ShareAccessTypePtrOutput {
	return o
}

func (o ShareAccessTypePtrOutput) Elem() ShareAccessTypeOutput {
	return o.ApplyT(func(v *ShareAccessType) ShareAccessType {
		if v != nil {
			return *v
		}
		var ret ShareAccessType
		return ret
	}).(ShareAccessTypeOutput)
}

func (o ShareAccessTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ShareAccessTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ShareAccessType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ShareAccessTypeInput interface {
	pulumi.Input

	ToShareAccessTypeOutput() ShareAccessTypeOutput
	ToShareAccessTypeOutputWithContext(context.Context) ShareAccessTypeOutput
}

var shareAccessTypePtrType = reflect.TypeOf((**ShareAccessType)(nil)).Elem()

type ShareAccessTypePtrInput interface {
	pulumi.Input

	ToShareAccessTypePtrOutput() ShareAccessTypePtrOutput
	ToShareAccessTypePtrOutputWithContext(context.Context) ShareAccessTypePtrOutput
}

type shareAccessTypePtr string

func ShareAccessTypePtr(v string) ShareAccessTypePtrInput {
	return (*shareAccessTypePtr)(&v)
}

func (*shareAccessTypePtr) ElementType() reflect.Type {
	return shareAccessTypePtrType
}

func (in *shareAccessTypePtr) ToShareAccessTypePtrOutput() ShareAccessTypePtrOutput {
	return pulumi.ToOutput(in).(ShareAccessTypePtrOutput)
}

func (in *shareAccessTypePtr) ToShareAccessTypePtrOutputWithContext(ctx context.Context) ShareAccessTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ShareAccessTypePtrOutput)
}

type ShareStatus string

const (
	ShareStatusOffline        = ShareStatus("Offline")
	ShareStatusUnknown        = ShareStatus("Unknown")
	ShareStatusOK             = ShareStatus("OK")
	ShareStatusUpdating       = ShareStatus("Updating")
	ShareStatusNeedsAttention = ShareStatus("NeedsAttention")
)

func (ShareStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*ShareStatus)(nil)).Elem()
}

func (e ShareStatus) ToShareStatusOutput() ShareStatusOutput {
	return pulumi.ToOutput(e).(ShareStatusOutput)
}

func (e ShareStatus) ToShareStatusOutputWithContext(ctx context.Context) ShareStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ShareStatusOutput)
}

func (e ShareStatus) ToShareStatusPtrOutput() ShareStatusPtrOutput {
	return e.ToShareStatusPtrOutputWithContext(context.Background())
}

func (e ShareStatus) ToShareStatusPtrOutputWithContext(ctx context.Context) ShareStatusPtrOutput {
	return ShareStatus(e).ToShareStatusOutputWithContext(ctx).ToShareStatusPtrOutputWithContext(ctx)
}

func (e ShareStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ShareStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ShareStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ShareStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ShareStatusOutput struct{ *pulumi.OutputState }

func (ShareStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ShareStatus)(nil)).Elem()
}

func (o ShareStatusOutput) ToShareStatusOutput() ShareStatusOutput {
	return o
}

func (o ShareStatusOutput) ToShareStatusOutputWithContext(ctx context.Context) ShareStatusOutput {
	return o
}

func (o ShareStatusOutput) ToShareStatusPtrOutput() ShareStatusPtrOutput {
	return o.ToShareStatusPtrOutputWithContext(context.Background())
}

func (o ShareStatusOutput) ToShareStatusPtrOutputWithContext(ctx context.Context) ShareStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ShareStatus) *ShareStatus {
		return &v
	}).(ShareStatusPtrOutput)
}

func (o ShareStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ShareStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ShareStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ShareStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ShareStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ShareStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ShareStatusPtrOutput struct{ *pulumi.OutputState }

func (ShareStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ShareStatus)(nil)).Elem()
}

func (o ShareStatusPtrOutput) ToShareStatusPtrOutput() ShareStatusPtrOutput {
	return o
}

func (o ShareStatusPtrOutput) ToShareStatusPtrOutputWithContext(ctx context.Context) ShareStatusPtrOutput {
	return o
}

func (o ShareStatusPtrOutput) Elem() ShareStatusOutput {
	return o.ApplyT(func(v *ShareStatus) ShareStatus {
		if v != nil {
			return *v
		}
		var ret ShareStatus
		return ret
	}).(ShareStatusOutput)
}

func (o ShareStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ShareStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ShareStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ShareStatusInput interface {
	pulumi.Input

	ToShareStatusOutput() ShareStatusOutput
	ToShareStatusOutputWithContext(context.Context) ShareStatusOutput
}

var shareStatusPtrType = reflect.TypeOf((**ShareStatus)(nil)).Elem()

type ShareStatusPtrInput interface {
	pulumi.Input

	ToShareStatusPtrOutput() ShareStatusPtrOutput
	ToShareStatusPtrOutputWithContext(context.Context) ShareStatusPtrOutput
}

type shareStatusPtr string

func ShareStatusPtr(v string) ShareStatusPtrInput {
	return (*shareStatusPtr)(&v)
}

func (*shareStatusPtr) ElementType() reflect.Type {
	return shareStatusPtrType
}

func (in *shareStatusPtr) ToShareStatusPtrOutput() ShareStatusPtrOutput {
	return pulumi.ToOutput(in).(ShareStatusPtrOutput)
}

func (in *shareStatusPtr) ToShareStatusPtrOutputWithContext(ctx context.Context) ShareStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ShareStatusPtrOutput)
}

type ShipmentType string

const (
	ShipmentTypeNotApplicable     = ShipmentType("NotApplicable")
	ShipmentTypeShippedToCustomer = ShipmentType("ShippedToCustomer")
	ShipmentTypeSelfPickup        = ShipmentType("SelfPickup")
)

func (ShipmentType) ElementType() reflect.Type {
	return reflect.TypeOf((*ShipmentType)(nil)).Elem()
}

func (e ShipmentType) ToShipmentTypeOutput() ShipmentTypeOutput {
	return pulumi.ToOutput(e).(ShipmentTypeOutput)
}

func (e ShipmentType) ToShipmentTypeOutputWithContext(ctx context.Context) ShipmentTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ShipmentTypeOutput)
}

func (e ShipmentType) ToShipmentTypePtrOutput() ShipmentTypePtrOutput {
	return e.ToShipmentTypePtrOutputWithContext(context.Background())
}

func (e ShipmentType) ToShipmentTypePtrOutputWithContext(ctx context.Context) ShipmentTypePtrOutput {
	return ShipmentType(e).ToShipmentTypeOutputWithContext(ctx).ToShipmentTypePtrOutputWithContext(ctx)
}

func (e ShipmentType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ShipmentType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ShipmentType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ShipmentType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ShipmentTypeOutput struct{ *pulumi.OutputState }

func (ShipmentTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ShipmentType)(nil)).Elem()
}

func (o ShipmentTypeOutput) ToShipmentTypeOutput() ShipmentTypeOutput {
	return o
}

func (o ShipmentTypeOutput) ToShipmentTypeOutputWithContext(ctx context.Context) ShipmentTypeOutput {
	return o
}

func (o ShipmentTypeOutput) ToShipmentTypePtrOutput() ShipmentTypePtrOutput {
	return o.ToShipmentTypePtrOutputWithContext(context.Background())
}

func (o ShipmentTypeOutput) ToShipmentTypePtrOutputWithContext(ctx context.Context) ShipmentTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ShipmentType) *ShipmentType {
		return &v
	}).(ShipmentTypePtrOutput)
}

func (o ShipmentTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ShipmentTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ShipmentType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ShipmentTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ShipmentTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ShipmentType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ShipmentTypePtrOutput struct{ *pulumi.OutputState }

func (ShipmentTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ShipmentType)(nil)).Elem()
}

func (o ShipmentTypePtrOutput) ToShipmentTypePtrOutput() ShipmentTypePtrOutput {
	return o
}

func (o ShipmentTypePtrOutput) ToShipmentTypePtrOutputWithContext(ctx context.Context) ShipmentTypePtrOutput {
	return o
}

func (o ShipmentTypePtrOutput) Elem() ShipmentTypeOutput {
	return o.ApplyT(func(v *ShipmentType) ShipmentType {
		if v != nil {
			return *v
		}
		var ret ShipmentType
		return ret
	}).(ShipmentTypeOutput)
}

func (o ShipmentTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ShipmentTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ShipmentType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ShipmentTypeInput interface {
	pulumi.Input

	ToShipmentTypeOutput() ShipmentTypeOutput
	ToShipmentTypeOutputWithContext(context.Context) ShipmentTypeOutput
}

var shipmentTypePtrType = reflect.TypeOf((**ShipmentType)(nil)).Elem()

type ShipmentTypePtrInput interface {
	pulumi.Input

	ToShipmentTypePtrOutput() ShipmentTypePtrOutput
	ToShipmentTypePtrOutputWithContext(context.Context) ShipmentTypePtrOutput
}

type shipmentTypePtr string

func ShipmentTypePtr(v string) ShipmentTypePtrInput {
	return (*shipmentTypePtr)(&v)
}

func (*shipmentTypePtr) ElementType() reflect.Type {
	return shipmentTypePtrType
}

func (in *shipmentTypePtr) ToShipmentTypePtrOutput() ShipmentTypePtrOutput {
	return pulumi.ToOutput(in).(ShipmentTypePtrOutput)
}

func (in *shipmentTypePtr) ToShipmentTypePtrOutputWithContext(ctx context.Context) ShipmentTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ShipmentTypePtrOutput)
}

type SkuName string

const (
	SkuNameGateway               = SkuName("Gateway")
	SkuNameEdge                  = SkuName("Edge")
	SkuName_TEA_1Node            = SkuName("TEA_1Node")
	SkuName_TEA_1Node_UPS        = SkuName("TEA_1Node_UPS")
	SkuName_TEA_1Node_Heater     = SkuName("TEA_1Node_Heater")
	SkuName_TEA_1Node_UPS_Heater = SkuName("TEA_1Node_UPS_Heater")
	SkuName_TEA_4Node_Heater     = SkuName("TEA_4Node_Heater")
	SkuName_TEA_4Node_UPS_Heater = SkuName("TEA_4Node_UPS_Heater")
	SkuNameTMA                   = SkuName("TMA")
	SkuNameTDC                   = SkuName("TDC")
	SkuName_TCA_Small            = SkuName("TCA_Small")
	SkuNameGPU                   = SkuName("GPU")
	SkuName_TCA_Large            = SkuName("TCA_Large")
	SkuName_EdgeP_Base           = SkuName("EdgeP_Base")
	SkuName_EdgeP_High           = SkuName("EdgeP_High")
	SkuName_EdgePR_Base          = SkuName("EdgePR_Base")
	SkuName_EdgePR_Base_UPS      = SkuName("EdgePR_Base_UPS")
	SkuName_EdgeMR_Mini          = SkuName("EdgeMR_Mini")
	SkuName_RCA_Small            = SkuName("RCA_Small")
	SkuName_RCA_Large            = SkuName("RCA_Large")
	SkuNameRDC                   = SkuName("RDC")
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

type SkuTier string

const (
	SkuTierStandard = SkuTier("Standard")
)

func (SkuTier) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuTier)(nil)).Elem()
}

func (e SkuTier) ToSkuTierOutput() SkuTierOutput {
	return pulumi.ToOutput(e).(SkuTierOutput)
}

func (e SkuTier) ToSkuTierOutputWithContext(ctx context.Context) SkuTierOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SkuTierOutput)
}

func (e SkuTier) ToSkuTierPtrOutput() SkuTierPtrOutput {
	return e.ToSkuTierPtrOutputWithContext(context.Background())
}

func (e SkuTier) ToSkuTierPtrOutputWithContext(ctx context.Context) SkuTierPtrOutput {
	return SkuTier(e).ToSkuTierOutputWithContext(ctx).ToSkuTierPtrOutputWithContext(ctx)
}

func (e SkuTier) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuTier) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuTier) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SkuTier) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SkuTierOutput struct{ *pulumi.OutputState }

func (SkuTierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuTier)(nil)).Elem()
}

func (o SkuTierOutput) ToSkuTierOutput() SkuTierOutput {
	return o
}

func (o SkuTierOutput) ToSkuTierOutputWithContext(ctx context.Context) SkuTierOutput {
	return o
}

func (o SkuTierOutput) ToSkuTierPtrOutput() SkuTierPtrOutput {
	return o.ToSkuTierPtrOutputWithContext(context.Background())
}

func (o SkuTierOutput) ToSkuTierPtrOutputWithContext(ctx context.Context) SkuTierPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuTier) *SkuTier {
		return &v
	}).(SkuTierPtrOutput)
}

func (o SkuTierOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SkuTierOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuTier) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SkuTierOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuTierOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuTier) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SkuTierPtrOutput struct{ *pulumi.OutputState }

func (SkuTierPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuTier)(nil)).Elem()
}

func (o SkuTierPtrOutput) ToSkuTierPtrOutput() SkuTierPtrOutput {
	return o
}

func (o SkuTierPtrOutput) ToSkuTierPtrOutputWithContext(ctx context.Context) SkuTierPtrOutput {
	return o
}

func (o SkuTierPtrOutput) Elem() SkuTierOutput {
	return o.ApplyT(func(v *SkuTier) SkuTier {
		if v != nil {
			return *v
		}
		var ret SkuTier
		return ret
	}).(SkuTierOutput)
}

func (o SkuTierPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuTierPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SkuTier) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SkuTierInput interface {
	pulumi.Input

	ToSkuTierOutput() SkuTierOutput
	ToSkuTierOutputWithContext(context.Context) SkuTierOutput
}

var skuTierPtrType = reflect.TypeOf((**SkuTier)(nil)).Elem()

type SkuTierPtrInput interface {
	pulumi.Input

	ToSkuTierPtrOutput() SkuTierPtrOutput
	ToSkuTierPtrOutputWithContext(context.Context) SkuTierPtrOutput
}

type skuTierPtr string

func SkuTierPtr(v string) SkuTierPtrInput {
	return (*skuTierPtr)(&v)
}

func (*skuTierPtr) ElementType() reflect.Type {
	return skuTierPtrType
}

func (in *skuTierPtr) ToSkuTierPtrOutput() SkuTierPtrOutput {
	return pulumi.ToOutput(in).(SkuTierPtrOutput)
}

func (in *skuTierPtr) ToSkuTierPtrOutputWithContext(ctx context.Context) SkuTierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SkuTierPtrOutput)
}

type StorageAccountStatus string

const (
	StorageAccountStatusOK             = StorageAccountStatus("OK")
	StorageAccountStatusOffline        = StorageAccountStatus("Offline")
	StorageAccountStatusUnknown        = StorageAccountStatus("Unknown")
	StorageAccountStatusUpdating       = StorageAccountStatus("Updating")
	StorageAccountStatusNeedsAttention = StorageAccountStatus("NeedsAttention")
)

func (StorageAccountStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountStatus)(nil)).Elem()
}

func (e StorageAccountStatus) ToStorageAccountStatusOutput() StorageAccountStatusOutput {
	return pulumi.ToOutput(e).(StorageAccountStatusOutput)
}

func (e StorageAccountStatus) ToStorageAccountStatusOutputWithContext(ctx context.Context) StorageAccountStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(StorageAccountStatusOutput)
}

func (e StorageAccountStatus) ToStorageAccountStatusPtrOutput() StorageAccountStatusPtrOutput {
	return e.ToStorageAccountStatusPtrOutputWithContext(context.Background())
}

func (e StorageAccountStatus) ToStorageAccountStatusPtrOutputWithContext(ctx context.Context) StorageAccountStatusPtrOutput {
	return StorageAccountStatus(e).ToStorageAccountStatusOutputWithContext(ctx).ToStorageAccountStatusPtrOutputWithContext(ctx)
}

func (e StorageAccountStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e StorageAccountStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e StorageAccountStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e StorageAccountStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type StorageAccountStatusOutput struct{ *pulumi.OutputState }

func (StorageAccountStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountStatus)(nil)).Elem()
}

func (o StorageAccountStatusOutput) ToStorageAccountStatusOutput() StorageAccountStatusOutput {
	return o
}

func (o StorageAccountStatusOutput) ToStorageAccountStatusOutputWithContext(ctx context.Context) StorageAccountStatusOutput {
	return o
}

func (o StorageAccountStatusOutput) ToStorageAccountStatusPtrOutput() StorageAccountStatusPtrOutput {
	return o.ToStorageAccountStatusPtrOutputWithContext(context.Background())
}

func (o StorageAccountStatusOutput) ToStorageAccountStatusPtrOutputWithContext(ctx context.Context) StorageAccountStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageAccountStatus) *StorageAccountStatus {
		return &v
	}).(StorageAccountStatusPtrOutput)
}

func (o StorageAccountStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o StorageAccountStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StorageAccountStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o StorageAccountStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StorageAccountStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StorageAccountStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type StorageAccountStatusPtrOutput struct{ *pulumi.OutputState }

func (StorageAccountStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccountStatus)(nil)).Elem()
}

func (o StorageAccountStatusPtrOutput) ToStorageAccountStatusPtrOutput() StorageAccountStatusPtrOutput {
	return o
}

func (o StorageAccountStatusPtrOutput) ToStorageAccountStatusPtrOutputWithContext(ctx context.Context) StorageAccountStatusPtrOutput {
	return o
}

func (o StorageAccountStatusPtrOutput) Elem() StorageAccountStatusOutput {
	return o.ApplyT(func(v *StorageAccountStatus) StorageAccountStatus {
		if v != nil {
			return *v
		}
		var ret StorageAccountStatus
		return ret
	}).(StorageAccountStatusOutput)
}

func (o StorageAccountStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StorageAccountStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *StorageAccountStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type StorageAccountStatusInput interface {
	pulumi.Input

	ToStorageAccountStatusOutput() StorageAccountStatusOutput
	ToStorageAccountStatusOutputWithContext(context.Context) StorageAccountStatusOutput
}

var storageAccountStatusPtrType = reflect.TypeOf((**StorageAccountStatus)(nil)).Elem()

type StorageAccountStatusPtrInput interface {
	pulumi.Input

	ToStorageAccountStatusPtrOutput() StorageAccountStatusPtrOutput
	ToStorageAccountStatusPtrOutputWithContext(context.Context) StorageAccountStatusPtrOutput
}

type storageAccountStatusPtr string

func StorageAccountStatusPtr(v string) StorageAccountStatusPtrInput {
	return (*storageAccountStatusPtr)(&v)
}

func (*storageAccountStatusPtr) ElementType() reflect.Type {
	return storageAccountStatusPtrType
}

func (in *storageAccountStatusPtr) ToStorageAccountStatusPtrOutput() StorageAccountStatusPtrOutput {
	return pulumi.ToOutput(in).(StorageAccountStatusPtrOutput)
}

func (in *storageAccountStatusPtr) ToStorageAccountStatusPtrOutputWithContext(ctx context.Context) StorageAccountStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(StorageAccountStatusPtrOutput)
}

type TriggerEventType string

const (
	TriggerEventTypeFileEvent          = TriggerEventType("FileEvent")
	TriggerEventTypePeriodicTimerEvent = TriggerEventType("PeriodicTimerEvent")
)

func (TriggerEventType) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerEventType)(nil)).Elem()
}

func (e TriggerEventType) ToTriggerEventTypeOutput() TriggerEventTypeOutput {
	return pulumi.ToOutput(e).(TriggerEventTypeOutput)
}

func (e TriggerEventType) ToTriggerEventTypeOutputWithContext(ctx context.Context) TriggerEventTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TriggerEventTypeOutput)
}

func (e TriggerEventType) ToTriggerEventTypePtrOutput() TriggerEventTypePtrOutput {
	return e.ToTriggerEventTypePtrOutputWithContext(context.Background())
}

func (e TriggerEventType) ToTriggerEventTypePtrOutputWithContext(ctx context.Context) TriggerEventTypePtrOutput {
	return TriggerEventType(e).ToTriggerEventTypeOutputWithContext(ctx).ToTriggerEventTypePtrOutputWithContext(ctx)
}

func (e TriggerEventType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TriggerEventType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TriggerEventType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TriggerEventType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TriggerEventTypeOutput struct{ *pulumi.OutputState }

func (TriggerEventTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerEventType)(nil)).Elem()
}

func (o TriggerEventTypeOutput) ToTriggerEventTypeOutput() TriggerEventTypeOutput {
	return o
}

func (o TriggerEventTypeOutput) ToTriggerEventTypeOutputWithContext(ctx context.Context) TriggerEventTypeOutput {
	return o
}

func (o TriggerEventTypeOutput) ToTriggerEventTypePtrOutput() TriggerEventTypePtrOutput {
	return o.ToTriggerEventTypePtrOutputWithContext(context.Background())
}

func (o TriggerEventTypeOutput) ToTriggerEventTypePtrOutputWithContext(ctx context.Context) TriggerEventTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TriggerEventType) *TriggerEventType {
		return &v
	}).(TriggerEventTypePtrOutput)
}

func (o TriggerEventTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TriggerEventTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TriggerEventType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TriggerEventTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TriggerEventTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TriggerEventType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TriggerEventTypePtrOutput struct{ *pulumi.OutputState }

func (TriggerEventTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TriggerEventType)(nil)).Elem()
}

func (o TriggerEventTypePtrOutput) ToTriggerEventTypePtrOutput() TriggerEventTypePtrOutput {
	return o
}

func (o TriggerEventTypePtrOutput) ToTriggerEventTypePtrOutputWithContext(ctx context.Context) TriggerEventTypePtrOutput {
	return o
}

func (o TriggerEventTypePtrOutput) Elem() TriggerEventTypeOutput {
	return o.ApplyT(func(v *TriggerEventType) TriggerEventType {
		if v != nil {
			return *v
		}
		var ret TriggerEventType
		return ret
	}).(TriggerEventTypeOutput)
}

func (o TriggerEventTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TriggerEventTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TriggerEventType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type TriggerEventTypeInput interface {
	pulumi.Input

	ToTriggerEventTypeOutput() TriggerEventTypeOutput
	ToTriggerEventTypeOutputWithContext(context.Context) TriggerEventTypeOutput
}

var triggerEventTypePtrType = reflect.TypeOf((**TriggerEventType)(nil)).Elem()

type TriggerEventTypePtrInput interface {
	pulumi.Input

	ToTriggerEventTypePtrOutput() TriggerEventTypePtrOutput
	ToTriggerEventTypePtrOutputWithContext(context.Context) TriggerEventTypePtrOutput
}

type triggerEventTypePtr string

func TriggerEventTypePtr(v string) TriggerEventTypePtrInput {
	return (*triggerEventTypePtr)(&v)
}

func (*triggerEventTypePtr) ElementType() reflect.Type {
	return triggerEventTypePtrType
}

func (in *triggerEventTypePtr) ToTriggerEventTypePtrOutput() TriggerEventTypePtrOutput {
	return pulumi.ToOutput(in).(TriggerEventTypePtrOutput)
}

func (in *triggerEventTypePtr) ToTriggerEventTypePtrOutputWithContext(ctx context.Context) TriggerEventTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TriggerEventTypePtrOutput)
}

type UserType string

const (
	UserTypeShare           = UserType("Share")
	UserTypeLocalManagement = UserType("LocalManagement")
	UserTypeARM             = UserType("ARM")
)

func (UserType) ElementType() reflect.Type {
	return reflect.TypeOf((*UserType)(nil)).Elem()
}

func (e UserType) ToUserTypeOutput() UserTypeOutput {
	return pulumi.ToOutput(e).(UserTypeOutput)
}

func (e UserType) ToUserTypeOutputWithContext(ctx context.Context) UserTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(UserTypeOutput)
}

func (e UserType) ToUserTypePtrOutput() UserTypePtrOutput {
	return e.ToUserTypePtrOutputWithContext(context.Background())
}

func (e UserType) ToUserTypePtrOutputWithContext(ctx context.Context) UserTypePtrOutput {
	return UserType(e).ToUserTypeOutputWithContext(ctx).ToUserTypePtrOutputWithContext(ctx)
}

func (e UserType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e UserType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e UserType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e UserType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type UserTypeOutput struct{ *pulumi.OutputState }

func (UserTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserType)(nil)).Elem()
}

func (o UserTypeOutput) ToUserTypeOutput() UserTypeOutput {
	return o
}

func (o UserTypeOutput) ToUserTypeOutputWithContext(ctx context.Context) UserTypeOutput {
	return o
}

func (o UserTypeOutput) ToUserTypePtrOutput() UserTypePtrOutput {
	return o.ToUserTypePtrOutputWithContext(context.Background())
}

func (o UserTypeOutput) ToUserTypePtrOutputWithContext(ctx context.Context) UserTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserType) *UserType {
		return &v
	}).(UserTypePtrOutput)
}

func (o UserTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o UserTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e UserType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o UserTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o UserTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e UserType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type UserTypePtrOutput struct{ *pulumi.OutputState }

func (UserTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserType)(nil)).Elem()
}

func (o UserTypePtrOutput) ToUserTypePtrOutput() UserTypePtrOutput {
	return o
}

func (o UserTypePtrOutput) ToUserTypePtrOutputWithContext(ctx context.Context) UserTypePtrOutput {
	return o
}

func (o UserTypePtrOutput) Elem() UserTypeOutput {
	return o.ApplyT(func(v *UserType) UserType {
		if v != nil {
			return *v
		}
		var ret UserType
		return ret
	}).(UserTypeOutput)
}

func (o UserTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o UserTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *UserType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type UserTypeInput interface {
	pulumi.Input

	ToUserTypeOutput() UserTypeOutput
	ToUserTypeOutputWithContext(context.Context) UserTypeOutput
}

var userTypePtrType = reflect.TypeOf((**UserType)(nil)).Elem()

type UserTypePtrInput interface {
	pulumi.Input

	ToUserTypePtrOutput() UserTypePtrOutput
	ToUserTypePtrOutputWithContext(context.Context) UserTypePtrOutput
}

type userTypePtr string

func UserTypePtr(v string) UserTypePtrInput {
	return (*userTypePtr)(&v)
}

func (*userTypePtr) ElementType() reflect.Type {
	return userTypePtrType
}

func (in *userTypePtr) ToUserTypePtrOutput() UserTypePtrOutput {
	return pulumi.ToOutput(in).(UserTypePtrOutput)
}

func (in *userTypePtr) ToUserTypePtrOutputWithContext(ctx context.Context) UserTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(UserTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AccountTypeOutput{})
	pulumi.RegisterOutputType(AccountTypePtrOutput{})
	pulumi.RegisterOutputType(AddonTypeOutput{})
	pulumi.RegisterOutputType(AddonTypePtrOutput{})
	pulumi.RegisterOutputType(AzureContainerDataFormatOutput{})
	pulumi.RegisterOutputType(AzureContainerDataFormatPtrOutput{})
	pulumi.RegisterOutputType(ClientPermissionTypeOutput{})
	pulumi.RegisterOutputType(ClientPermissionTypePtrOutput{})
	pulumi.RegisterOutputType(DataBoxEdgeDeviceStatusOutput{})
	pulumi.RegisterOutputType(DataBoxEdgeDeviceStatusPtrOutput{})
	pulumi.RegisterOutputType(DataPolicyOutput{})
	pulumi.RegisterOutputType(DataPolicyPtrOutput{})
	pulumi.RegisterOutputType(DayOfWeekOutput{})
	pulumi.RegisterOutputType(DayOfWeekPtrOutput{})
	pulumi.RegisterOutputType(EncryptionAlgorithmOutput{})
	pulumi.RegisterOutputType(EncryptionAlgorithmPtrOutput{})
	pulumi.RegisterOutputType(MonitoringStatusOutput{})
	pulumi.RegisterOutputType(MonitoringStatusPtrOutput{})
	pulumi.RegisterOutputType(MsiIdentityTypeOutput{})
	pulumi.RegisterOutputType(MsiIdentityTypePtrOutput{})
	pulumi.RegisterOutputType(PlatformTypeOutput{})
	pulumi.RegisterOutputType(PlatformTypePtrOutput{})
	pulumi.RegisterOutputType(RoleStatusOutput{})
	pulumi.RegisterOutputType(RoleStatusPtrOutput{})
	pulumi.RegisterOutputType(RoleTypesOutput{})
	pulumi.RegisterOutputType(RoleTypesPtrOutput{})
	pulumi.RegisterOutputType(SSLStatusOutput{})
	pulumi.RegisterOutputType(SSLStatusPtrOutput{})
	pulumi.RegisterOutputType(ShareAccessProtocolOutput{})
	pulumi.RegisterOutputType(ShareAccessProtocolPtrOutput{})
	pulumi.RegisterOutputType(ShareAccessTypeOutput{})
	pulumi.RegisterOutputType(ShareAccessTypePtrOutput{})
	pulumi.RegisterOutputType(ShareStatusOutput{})
	pulumi.RegisterOutputType(ShareStatusPtrOutput{})
	pulumi.RegisterOutputType(ShipmentTypeOutput{})
	pulumi.RegisterOutputType(ShipmentTypePtrOutput{})
	pulumi.RegisterOutputType(SkuNameOutput{})
	pulumi.RegisterOutputType(SkuNamePtrOutput{})
	pulumi.RegisterOutputType(SkuTierOutput{})
	pulumi.RegisterOutputType(SkuTierPtrOutput{})
	pulumi.RegisterOutputType(StorageAccountStatusOutput{})
	pulumi.RegisterOutputType(StorageAccountStatusPtrOutput{})
	pulumi.RegisterOutputType(TriggerEventTypeOutput{})
	pulumi.RegisterOutputType(TriggerEventTypePtrOutput{})
	pulumi.RegisterOutputType(UserTypeOutput{})
	pulumi.RegisterOutputType(UserTypePtrOutput{})
}
