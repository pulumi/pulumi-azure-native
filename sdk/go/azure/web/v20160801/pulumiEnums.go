


package v20160801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AutoHealActionType string

const (
	AutoHealActionTypeRecycle      = AutoHealActionType("Recycle")
	AutoHealActionTypeLogEvent     = AutoHealActionType("LogEvent")
	AutoHealActionTypeCustomAction = AutoHealActionType("CustomAction")
)

func (AutoHealActionType) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoHealActionType)(nil)).Elem()
}

func (e AutoHealActionType) ToAutoHealActionTypeOutput() AutoHealActionTypeOutput {
	return pulumi.ToOutput(e).(AutoHealActionTypeOutput)
}

func (e AutoHealActionType) ToAutoHealActionTypeOutputWithContext(ctx context.Context) AutoHealActionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AutoHealActionTypeOutput)
}

func (e AutoHealActionType) ToAutoHealActionTypePtrOutput() AutoHealActionTypePtrOutput {
	return e.ToAutoHealActionTypePtrOutputWithContext(context.Background())
}

func (e AutoHealActionType) ToAutoHealActionTypePtrOutputWithContext(ctx context.Context) AutoHealActionTypePtrOutput {
	return AutoHealActionType(e).ToAutoHealActionTypeOutputWithContext(ctx).ToAutoHealActionTypePtrOutputWithContext(ctx)
}

func (e AutoHealActionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutoHealActionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutoHealActionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AutoHealActionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AutoHealActionTypeOutput struct{ *pulumi.OutputState }

func (AutoHealActionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoHealActionType)(nil)).Elem()
}

func (o AutoHealActionTypeOutput) ToAutoHealActionTypeOutput() AutoHealActionTypeOutput {
	return o
}

func (o AutoHealActionTypeOutput) ToAutoHealActionTypeOutputWithContext(ctx context.Context) AutoHealActionTypeOutput {
	return o
}

func (o AutoHealActionTypeOutput) ToAutoHealActionTypePtrOutput() AutoHealActionTypePtrOutput {
	return o.ToAutoHealActionTypePtrOutputWithContext(context.Background())
}

func (o AutoHealActionTypeOutput) ToAutoHealActionTypePtrOutputWithContext(ctx context.Context) AutoHealActionTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoHealActionType) *AutoHealActionType {
		return &v
	}).(AutoHealActionTypePtrOutput)
}

func (o AutoHealActionTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AutoHealActionTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AutoHealActionType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AutoHealActionTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AutoHealActionTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AutoHealActionType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AutoHealActionTypePtrOutput struct{ *pulumi.OutputState }

func (AutoHealActionTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoHealActionType)(nil)).Elem()
}

func (o AutoHealActionTypePtrOutput) ToAutoHealActionTypePtrOutput() AutoHealActionTypePtrOutput {
	return o
}

func (o AutoHealActionTypePtrOutput) ToAutoHealActionTypePtrOutputWithContext(ctx context.Context) AutoHealActionTypePtrOutput {
	return o
}

func (o AutoHealActionTypePtrOutput) Elem() AutoHealActionTypeOutput {
	return o.ApplyT(func(v *AutoHealActionType) AutoHealActionType {
		if v != nil {
			return *v
		}
		var ret AutoHealActionType
		return ret
	}).(AutoHealActionTypeOutput)
}

func (o AutoHealActionTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AutoHealActionTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AutoHealActionType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AutoHealActionTypeInput interface {
	pulumi.Input

	ToAutoHealActionTypeOutput() AutoHealActionTypeOutput
	ToAutoHealActionTypeOutputWithContext(context.Context) AutoHealActionTypeOutput
}

var autoHealActionTypePtrType = reflect.TypeOf((**AutoHealActionType)(nil)).Elem()

type AutoHealActionTypePtrInput interface {
	pulumi.Input

	ToAutoHealActionTypePtrOutput() AutoHealActionTypePtrOutput
	ToAutoHealActionTypePtrOutputWithContext(context.Context) AutoHealActionTypePtrOutput
}

type autoHealActionTypePtr string

func AutoHealActionTypePtr(v string) AutoHealActionTypePtrInput {
	return (*autoHealActionTypePtr)(&v)
}

func (*autoHealActionTypePtr) ElementType() reflect.Type {
	return autoHealActionTypePtrType
}

func (in *autoHealActionTypePtr) ToAutoHealActionTypePtrOutput() AutoHealActionTypePtrOutput {
	return pulumi.ToOutput(in).(AutoHealActionTypePtrOutput)
}

func (in *autoHealActionTypePtr) ToAutoHealActionTypePtrOutputWithContext(ctx context.Context) AutoHealActionTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AutoHealActionTypePtrOutput)
}

type AzureResourceType string

const (
	AzureResourceTypeWebsite        = AzureResourceType("Website")
	AzureResourceTypeTrafficManager = AzureResourceType("TrafficManager")
)

func (AzureResourceType) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureResourceType)(nil)).Elem()
}

func (e AzureResourceType) ToAzureResourceTypeOutput() AzureResourceTypeOutput {
	return pulumi.ToOutput(e).(AzureResourceTypeOutput)
}

func (e AzureResourceType) ToAzureResourceTypeOutputWithContext(ctx context.Context) AzureResourceTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AzureResourceTypeOutput)
}

func (e AzureResourceType) ToAzureResourceTypePtrOutput() AzureResourceTypePtrOutput {
	return e.ToAzureResourceTypePtrOutputWithContext(context.Background())
}

func (e AzureResourceType) ToAzureResourceTypePtrOutputWithContext(ctx context.Context) AzureResourceTypePtrOutput {
	return AzureResourceType(e).ToAzureResourceTypeOutputWithContext(ctx).ToAzureResourceTypePtrOutputWithContext(ctx)
}

func (e AzureResourceType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureResourceType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureResourceType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AzureResourceType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AzureResourceTypeOutput struct{ *pulumi.OutputState }

func (AzureResourceTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureResourceType)(nil)).Elem()
}

func (o AzureResourceTypeOutput) ToAzureResourceTypeOutput() AzureResourceTypeOutput {
	return o
}

func (o AzureResourceTypeOutput) ToAzureResourceTypeOutputWithContext(ctx context.Context) AzureResourceTypeOutput {
	return o
}

func (o AzureResourceTypeOutput) ToAzureResourceTypePtrOutput() AzureResourceTypePtrOutput {
	return o.ToAzureResourceTypePtrOutputWithContext(context.Background())
}

func (o AzureResourceTypeOutput) ToAzureResourceTypePtrOutputWithContext(ctx context.Context) AzureResourceTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureResourceType) *AzureResourceType {
		return &v
	}).(AzureResourceTypePtrOutput)
}

func (o AzureResourceTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AzureResourceTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureResourceType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AzureResourceTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureResourceTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureResourceType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AzureResourceTypePtrOutput struct{ *pulumi.OutputState }

func (AzureResourceTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureResourceType)(nil)).Elem()
}

func (o AzureResourceTypePtrOutput) ToAzureResourceTypePtrOutput() AzureResourceTypePtrOutput {
	return o
}

func (o AzureResourceTypePtrOutput) ToAzureResourceTypePtrOutputWithContext(ctx context.Context) AzureResourceTypePtrOutput {
	return o
}

func (o AzureResourceTypePtrOutput) Elem() AzureResourceTypeOutput {
	return o.ApplyT(func(v *AzureResourceType) AzureResourceType {
		if v != nil {
			return *v
		}
		var ret AzureResourceType
		return ret
	}).(AzureResourceTypeOutput)
}

func (o AzureResourceTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureResourceTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AzureResourceType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AzureResourceTypeInput interface {
	pulumi.Input

	ToAzureResourceTypeOutput() AzureResourceTypeOutput
	ToAzureResourceTypeOutputWithContext(context.Context) AzureResourceTypeOutput
}

var azureResourceTypePtrType = reflect.TypeOf((**AzureResourceType)(nil)).Elem()

type AzureResourceTypePtrInput interface {
	pulumi.Input

	ToAzureResourceTypePtrOutput() AzureResourceTypePtrOutput
	ToAzureResourceTypePtrOutputWithContext(context.Context) AzureResourceTypePtrOutput
}

type azureResourceTypePtr string

func AzureResourceTypePtr(v string) AzureResourceTypePtrInput {
	return (*azureResourceTypePtr)(&v)
}

func (*azureResourceTypePtr) ElementType() reflect.Type {
	return azureResourceTypePtrType
}

func (in *azureResourceTypePtr) ToAzureResourceTypePtrOutput() AzureResourceTypePtrOutput {
	return pulumi.ToOutput(in).(AzureResourceTypePtrOutput)
}

func (in *azureResourceTypePtr) ToAzureResourceTypePtrOutputWithContext(ctx context.Context) AzureResourceTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AzureResourceTypePtrOutput)
}

type BackupRestoreOperationType string

const (
	BackupRestoreOperationTypeDefault    = BackupRestoreOperationType("Default")
	BackupRestoreOperationTypeClone      = BackupRestoreOperationType("Clone")
	BackupRestoreOperationTypeRelocation = BackupRestoreOperationType("Relocation")
	BackupRestoreOperationTypeSnapshot   = BackupRestoreOperationType("Snapshot")
)

func (BackupRestoreOperationType) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupRestoreOperationType)(nil)).Elem()
}

func (e BackupRestoreOperationType) ToBackupRestoreOperationTypeOutput() BackupRestoreOperationTypeOutput {
	return pulumi.ToOutput(e).(BackupRestoreOperationTypeOutput)
}

func (e BackupRestoreOperationType) ToBackupRestoreOperationTypeOutputWithContext(ctx context.Context) BackupRestoreOperationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(BackupRestoreOperationTypeOutput)
}

func (e BackupRestoreOperationType) ToBackupRestoreOperationTypePtrOutput() BackupRestoreOperationTypePtrOutput {
	return e.ToBackupRestoreOperationTypePtrOutputWithContext(context.Background())
}

func (e BackupRestoreOperationType) ToBackupRestoreOperationTypePtrOutputWithContext(ctx context.Context) BackupRestoreOperationTypePtrOutput {
	return BackupRestoreOperationType(e).ToBackupRestoreOperationTypeOutputWithContext(ctx).ToBackupRestoreOperationTypePtrOutputWithContext(ctx)
}

func (e BackupRestoreOperationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e BackupRestoreOperationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e BackupRestoreOperationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e BackupRestoreOperationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type BackupRestoreOperationTypeOutput struct{ *pulumi.OutputState }

func (BackupRestoreOperationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupRestoreOperationType)(nil)).Elem()
}

func (o BackupRestoreOperationTypeOutput) ToBackupRestoreOperationTypeOutput() BackupRestoreOperationTypeOutput {
	return o
}

func (o BackupRestoreOperationTypeOutput) ToBackupRestoreOperationTypeOutputWithContext(ctx context.Context) BackupRestoreOperationTypeOutput {
	return o
}

func (o BackupRestoreOperationTypeOutput) ToBackupRestoreOperationTypePtrOutput() BackupRestoreOperationTypePtrOutput {
	return o.ToBackupRestoreOperationTypePtrOutputWithContext(context.Background())
}

func (o BackupRestoreOperationTypeOutput) ToBackupRestoreOperationTypePtrOutputWithContext(ctx context.Context) BackupRestoreOperationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BackupRestoreOperationType) *BackupRestoreOperationType {
		return &v
	}).(BackupRestoreOperationTypePtrOutput)
}

func (o BackupRestoreOperationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o BackupRestoreOperationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BackupRestoreOperationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o BackupRestoreOperationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BackupRestoreOperationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BackupRestoreOperationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type BackupRestoreOperationTypePtrOutput struct{ *pulumi.OutputState }

func (BackupRestoreOperationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupRestoreOperationType)(nil)).Elem()
}

func (o BackupRestoreOperationTypePtrOutput) ToBackupRestoreOperationTypePtrOutput() BackupRestoreOperationTypePtrOutput {
	return o
}

func (o BackupRestoreOperationTypePtrOutput) ToBackupRestoreOperationTypePtrOutputWithContext(ctx context.Context) BackupRestoreOperationTypePtrOutput {
	return o
}

func (o BackupRestoreOperationTypePtrOutput) Elem() BackupRestoreOperationTypeOutput {
	return o.ApplyT(func(v *BackupRestoreOperationType) BackupRestoreOperationType {
		if v != nil {
			return *v
		}
		var ret BackupRestoreOperationType
		return ret
	}).(BackupRestoreOperationTypeOutput)
}

func (o BackupRestoreOperationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BackupRestoreOperationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *BackupRestoreOperationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type BackupRestoreOperationTypeInput interface {
	pulumi.Input

	ToBackupRestoreOperationTypeOutput() BackupRestoreOperationTypeOutput
	ToBackupRestoreOperationTypeOutputWithContext(context.Context) BackupRestoreOperationTypeOutput
}

var backupRestoreOperationTypePtrType = reflect.TypeOf((**BackupRestoreOperationType)(nil)).Elem()

type BackupRestoreOperationTypePtrInput interface {
	pulumi.Input

	ToBackupRestoreOperationTypePtrOutput() BackupRestoreOperationTypePtrOutput
	ToBackupRestoreOperationTypePtrOutputWithContext(context.Context) BackupRestoreOperationTypePtrOutput
}

type backupRestoreOperationTypePtr string

func BackupRestoreOperationTypePtr(v string) BackupRestoreOperationTypePtrInput {
	return (*backupRestoreOperationTypePtr)(&v)
}

func (*backupRestoreOperationTypePtr) ElementType() reflect.Type {
	return backupRestoreOperationTypePtrType
}

func (in *backupRestoreOperationTypePtr) ToBackupRestoreOperationTypePtrOutput() BackupRestoreOperationTypePtrOutput {
	return pulumi.ToOutput(in).(BackupRestoreOperationTypePtrOutput)
}

func (in *backupRestoreOperationTypePtr) ToBackupRestoreOperationTypePtrOutputWithContext(ctx context.Context) BackupRestoreOperationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(BackupRestoreOperationTypePtrOutput)
}

type BuiltInAuthenticationProvider string

const (
	BuiltInAuthenticationProviderAzureActiveDirectory = BuiltInAuthenticationProvider("AzureActiveDirectory")
	BuiltInAuthenticationProviderFacebook             = BuiltInAuthenticationProvider("Facebook")
	BuiltInAuthenticationProviderGoogle               = BuiltInAuthenticationProvider("Google")
	BuiltInAuthenticationProviderMicrosoftAccount     = BuiltInAuthenticationProvider("MicrosoftAccount")
	BuiltInAuthenticationProviderTwitter              = BuiltInAuthenticationProvider("Twitter")
)

func (BuiltInAuthenticationProvider) ElementType() reflect.Type {
	return reflect.TypeOf((*BuiltInAuthenticationProvider)(nil)).Elem()
}

func (e BuiltInAuthenticationProvider) ToBuiltInAuthenticationProviderOutput() BuiltInAuthenticationProviderOutput {
	return pulumi.ToOutput(e).(BuiltInAuthenticationProviderOutput)
}

func (e BuiltInAuthenticationProvider) ToBuiltInAuthenticationProviderOutputWithContext(ctx context.Context) BuiltInAuthenticationProviderOutput {
	return pulumi.ToOutputWithContext(ctx, e).(BuiltInAuthenticationProviderOutput)
}

func (e BuiltInAuthenticationProvider) ToBuiltInAuthenticationProviderPtrOutput() BuiltInAuthenticationProviderPtrOutput {
	return e.ToBuiltInAuthenticationProviderPtrOutputWithContext(context.Background())
}

func (e BuiltInAuthenticationProvider) ToBuiltInAuthenticationProviderPtrOutputWithContext(ctx context.Context) BuiltInAuthenticationProviderPtrOutput {
	return BuiltInAuthenticationProvider(e).ToBuiltInAuthenticationProviderOutputWithContext(ctx).ToBuiltInAuthenticationProviderPtrOutputWithContext(ctx)
}

func (e BuiltInAuthenticationProvider) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e BuiltInAuthenticationProvider) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e BuiltInAuthenticationProvider) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e BuiltInAuthenticationProvider) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type BuiltInAuthenticationProviderOutput struct{ *pulumi.OutputState }

func (BuiltInAuthenticationProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BuiltInAuthenticationProvider)(nil)).Elem()
}

func (o BuiltInAuthenticationProviderOutput) ToBuiltInAuthenticationProviderOutput() BuiltInAuthenticationProviderOutput {
	return o
}

func (o BuiltInAuthenticationProviderOutput) ToBuiltInAuthenticationProviderOutputWithContext(ctx context.Context) BuiltInAuthenticationProviderOutput {
	return o
}

func (o BuiltInAuthenticationProviderOutput) ToBuiltInAuthenticationProviderPtrOutput() BuiltInAuthenticationProviderPtrOutput {
	return o.ToBuiltInAuthenticationProviderPtrOutputWithContext(context.Background())
}

func (o BuiltInAuthenticationProviderOutput) ToBuiltInAuthenticationProviderPtrOutputWithContext(ctx context.Context) BuiltInAuthenticationProviderPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BuiltInAuthenticationProvider) *BuiltInAuthenticationProvider {
		return &v
	}).(BuiltInAuthenticationProviderPtrOutput)
}

func (o BuiltInAuthenticationProviderOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o BuiltInAuthenticationProviderOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BuiltInAuthenticationProvider) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o BuiltInAuthenticationProviderOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BuiltInAuthenticationProviderOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BuiltInAuthenticationProvider) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type BuiltInAuthenticationProviderPtrOutput struct{ *pulumi.OutputState }

func (BuiltInAuthenticationProviderPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BuiltInAuthenticationProvider)(nil)).Elem()
}

func (o BuiltInAuthenticationProviderPtrOutput) ToBuiltInAuthenticationProviderPtrOutput() BuiltInAuthenticationProviderPtrOutput {
	return o
}

func (o BuiltInAuthenticationProviderPtrOutput) ToBuiltInAuthenticationProviderPtrOutputWithContext(ctx context.Context) BuiltInAuthenticationProviderPtrOutput {
	return o
}

func (o BuiltInAuthenticationProviderPtrOutput) Elem() BuiltInAuthenticationProviderOutput {
	return o.ApplyT(func(v *BuiltInAuthenticationProvider) BuiltInAuthenticationProvider {
		if v != nil {
			return *v
		}
		var ret BuiltInAuthenticationProvider
		return ret
	}).(BuiltInAuthenticationProviderOutput)
}

func (o BuiltInAuthenticationProviderPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BuiltInAuthenticationProviderPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *BuiltInAuthenticationProvider) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type BuiltInAuthenticationProviderInput interface {
	pulumi.Input

	ToBuiltInAuthenticationProviderOutput() BuiltInAuthenticationProviderOutput
	ToBuiltInAuthenticationProviderOutputWithContext(context.Context) BuiltInAuthenticationProviderOutput
}

var builtInAuthenticationProviderPtrType = reflect.TypeOf((**BuiltInAuthenticationProvider)(nil)).Elem()

type BuiltInAuthenticationProviderPtrInput interface {
	pulumi.Input

	ToBuiltInAuthenticationProviderPtrOutput() BuiltInAuthenticationProviderPtrOutput
	ToBuiltInAuthenticationProviderPtrOutputWithContext(context.Context) BuiltInAuthenticationProviderPtrOutput
}

type builtInAuthenticationProviderPtr string

func BuiltInAuthenticationProviderPtr(v string) BuiltInAuthenticationProviderPtrInput {
	return (*builtInAuthenticationProviderPtr)(&v)
}

func (*builtInAuthenticationProviderPtr) ElementType() reflect.Type {
	return builtInAuthenticationProviderPtrType
}

func (in *builtInAuthenticationProviderPtr) ToBuiltInAuthenticationProviderPtrOutput() BuiltInAuthenticationProviderPtrOutput {
	return pulumi.ToOutput(in).(BuiltInAuthenticationProviderPtrOutput)
}

func (in *builtInAuthenticationProviderPtr) ToBuiltInAuthenticationProviderPtrOutputWithContext(ctx context.Context) BuiltInAuthenticationProviderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(BuiltInAuthenticationProviderPtrOutput)
}

type ConnectionStringType string

const (
	ConnectionStringTypeMySql           = ConnectionStringType("MySql")
	ConnectionStringTypeSQLServer       = ConnectionStringType("SQLServer")
	ConnectionStringTypeSQLAzure        = ConnectionStringType("SQLAzure")
	ConnectionStringTypeCustom          = ConnectionStringType("Custom")
	ConnectionStringTypeNotificationHub = ConnectionStringType("NotificationHub")
	ConnectionStringTypeServiceBus      = ConnectionStringType("ServiceBus")
	ConnectionStringTypeEventHub        = ConnectionStringType("EventHub")
	ConnectionStringTypeApiHub          = ConnectionStringType("ApiHub")
	ConnectionStringTypeDocDb           = ConnectionStringType("DocDb")
	ConnectionStringTypeRedisCache      = ConnectionStringType("RedisCache")
	ConnectionStringTypePostgreSQL      = ConnectionStringType("PostgreSQL")
)

func (ConnectionStringType) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionStringType)(nil)).Elem()
}

func (e ConnectionStringType) ToConnectionStringTypeOutput() ConnectionStringTypeOutput {
	return pulumi.ToOutput(e).(ConnectionStringTypeOutput)
}

func (e ConnectionStringType) ToConnectionStringTypeOutputWithContext(ctx context.Context) ConnectionStringTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ConnectionStringTypeOutput)
}

func (e ConnectionStringType) ToConnectionStringTypePtrOutput() ConnectionStringTypePtrOutput {
	return e.ToConnectionStringTypePtrOutputWithContext(context.Background())
}

func (e ConnectionStringType) ToConnectionStringTypePtrOutputWithContext(ctx context.Context) ConnectionStringTypePtrOutput {
	return ConnectionStringType(e).ToConnectionStringTypeOutputWithContext(ctx).ToConnectionStringTypePtrOutputWithContext(ctx)
}

func (e ConnectionStringType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConnectionStringType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConnectionStringType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ConnectionStringType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ConnectionStringTypeOutput struct{ *pulumi.OutputState }

func (ConnectionStringTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionStringType)(nil)).Elem()
}

func (o ConnectionStringTypeOutput) ToConnectionStringTypeOutput() ConnectionStringTypeOutput {
	return o
}

func (o ConnectionStringTypeOutput) ToConnectionStringTypeOutputWithContext(ctx context.Context) ConnectionStringTypeOutput {
	return o
}

func (o ConnectionStringTypeOutput) ToConnectionStringTypePtrOutput() ConnectionStringTypePtrOutput {
	return o.ToConnectionStringTypePtrOutputWithContext(context.Background())
}

func (o ConnectionStringTypeOutput) ToConnectionStringTypePtrOutputWithContext(ctx context.Context) ConnectionStringTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectionStringType) *ConnectionStringType {
		return &v
	}).(ConnectionStringTypePtrOutput)
}

func (o ConnectionStringTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ConnectionStringTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConnectionStringType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ConnectionStringTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConnectionStringTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConnectionStringType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ConnectionStringTypePtrOutput struct{ *pulumi.OutputState }

func (ConnectionStringTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionStringType)(nil)).Elem()
}

func (o ConnectionStringTypePtrOutput) ToConnectionStringTypePtrOutput() ConnectionStringTypePtrOutput {
	return o
}

func (o ConnectionStringTypePtrOutput) ToConnectionStringTypePtrOutputWithContext(ctx context.Context) ConnectionStringTypePtrOutput {
	return o
}

func (o ConnectionStringTypePtrOutput) Elem() ConnectionStringTypeOutput {
	return o.ApplyT(func(v *ConnectionStringType) ConnectionStringType {
		if v != nil {
			return *v
		}
		var ret ConnectionStringType
		return ret
	}).(ConnectionStringTypeOutput)
}

func (o ConnectionStringTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConnectionStringTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ConnectionStringType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ConnectionStringTypeInput interface {
	pulumi.Input

	ToConnectionStringTypeOutput() ConnectionStringTypeOutput
	ToConnectionStringTypeOutputWithContext(context.Context) ConnectionStringTypeOutput
}

var connectionStringTypePtrType = reflect.TypeOf((**ConnectionStringType)(nil)).Elem()

type ConnectionStringTypePtrInput interface {
	pulumi.Input

	ToConnectionStringTypePtrOutput() ConnectionStringTypePtrOutput
	ToConnectionStringTypePtrOutputWithContext(context.Context) ConnectionStringTypePtrOutput
}

type connectionStringTypePtr string

func ConnectionStringTypePtr(v string) ConnectionStringTypePtrInput {
	return (*connectionStringTypePtr)(&v)
}

func (*connectionStringTypePtr) ElementType() reflect.Type {
	return connectionStringTypePtrType
}

func (in *connectionStringTypePtr) ToConnectionStringTypePtrOutput() ConnectionStringTypePtrOutput {
	return pulumi.ToOutput(in).(ConnectionStringTypePtrOutput)
}

func (in *connectionStringTypePtr) ToConnectionStringTypePtrOutputWithContext(ctx context.Context) ConnectionStringTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ConnectionStringTypePtrOutput)
}

type CustomHostNameDnsRecordType string

const (
	CustomHostNameDnsRecordTypeCName = CustomHostNameDnsRecordType("CName")
	CustomHostNameDnsRecordTypeA     = CustomHostNameDnsRecordType("A")
)

func (CustomHostNameDnsRecordType) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomHostNameDnsRecordType)(nil)).Elem()
}

func (e CustomHostNameDnsRecordType) ToCustomHostNameDnsRecordTypeOutput() CustomHostNameDnsRecordTypeOutput {
	return pulumi.ToOutput(e).(CustomHostNameDnsRecordTypeOutput)
}

func (e CustomHostNameDnsRecordType) ToCustomHostNameDnsRecordTypeOutputWithContext(ctx context.Context) CustomHostNameDnsRecordTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CustomHostNameDnsRecordTypeOutput)
}

func (e CustomHostNameDnsRecordType) ToCustomHostNameDnsRecordTypePtrOutput() CustomHostNameDnsRecordTypePtrOutput {
	return e.ToCustomHostNameDnsRecordTypePtrOutputWithContext(context.Background())
}

func (e CustomHostNameDnsRecordType) ToCustomHostNameDnsRecordTypePtrOutputWithContext(ctx context.Context) CustomHostNameDnsRecordTypePtrOutput {
	return CustomHostNameDnsRecordType(e).ToCustomHostNameDnsRecordTypeOutputWithContext(ctx).ToCustomHostNameDnsRecordTypePtrOutputWithContext(ctx)
}

func (e CustomHostNameDnsRecordType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CustomHostNameDnsRecordType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CustomHostNameDnsRecordType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CustomHostNameDnsRecordType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CustomHostNameDnsRecordTypeOutput struct{ *pulumi.OutputState }

func (CustomHostNameDnsRecordTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomHostNameDnsRecordType)(nil)).Elem()
}

func (o CustomHostNameDnsRecordTypeOutput) ToCustomHostNameDnsRecordTypeOutput() CustomHostNameDnsRecordTypeOutput {
	return o
}

func (o CustomHostNameDnsRecordTypeOutput) ToCustomHostNameDnsRecordTypeOutputWithContext(ctx context.Context) CustomHostNameDnsRecordTypeOutput {
	return o
}

func (o CustomHostNameDnsRecordTypeOutput) ToCustomHostNameDnsRecordTypePtrOutput() CustomHostNameDnsRecordTypePtrOutput {
	return o.ToCustomHostNameDnsRecordTypePtrOutputWithContext(context.Background())
}

func (o CustomHostNameDnsRecordTypeOutput) ToCustomHostNameDnsRecordTypePtrOutputWithContext(ctx context.Context) CustomHostNameDnsRecordTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomHostNameDnsRecordType) *CustomHostNameDnsRecordType {
		return &v
	}).(CustomHostNameDnsRecordTypePtrOutput)
}

func (o CustomHostNameDnsRecordTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CustomHostNameDnsRecordTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CustomHostNameDnsRecordType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CustomHostNameDnsRecordTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CustomHostNameDnsRecordTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CustomHostNameDnsRecordType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CustomHostNameDnsRecordTypePtrOutput struct{ *pulumi.OutputState }

func (CustomHostNameDnsRecordTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomHostNameDnsRecordType)(nil)).Elem()
}

func (o CustomHostNameDnsRecordTypePtrOutput) ToCustomHostNameDnsRecordTypePtrOutput() CustomHostNameDnsRecordTypePtrOutput {
	return o
}

func (o CustomHostNameDnsRecordTypePtrOutput) ToCustomHostNameDnsRecordTypePtrOutputWithContext(ctx context.Context) CustomHostNameDnsRecordTypePtrOutput {
	return o
}

func (o CustomHostNameDnsRecordTypePtrOutput) Elem() CustomHostNameDnsRecordTypeOutput {
	return o.ApplyT(func(v *CustomHostNameDnsRecordType) CustomHostNameDnsRecordType {
		if v != nil {
			return *v
		}
		var ret CustomHostNameDnsRecordType
		return ret
	}).(CustomHostNameDnsRecordTypeOutput)
}

func (o CustomHostNameDnsRecordTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CustomHostNameDnsRecordTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CustomHostNameDnsRecordType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type CustomHostNameDnsRecordTypeInput interface {
	pulumi.Input

	ToCustomHostNameDnsRecordTypeOutput() CustomHostNameDnsRecordTypeOutput
	ToCustomHostNameDnsRecordTypeOutputWithContext(context.Context) CustomHostNameDnsRecordTypeOutput
}

var customHostNameDnsRecordTypePtrType = reflect.TypeOf((**CustomHostNameDnsRecordType)(nil)).Elem()

type CustomHostNameDnsRecordTypePtrInput interface {
	pulumi.Input

	ToCustomHostNameDnsRecordTypePtrOutput() CustomHostNameDnsRecordTypePtrOutput
	ToCustomHostNameDnsRecordTypePtrOutputWithContext(context.Context) CustomHostNameDnsRecordTypePtrOutput
}

type customHostNameDnsRecordTypePtr string

func CustomHostNameDnsRecordTypePtr(v string) CustomHostNameDnsRecordTypePtrInput {
	return (*customHostNameDnsRecordTypePtr)(&v)
}

func (*customHostNameDnsRecordTypePtr) ElementType() reflect.Type {
	return customHostNameDnsRecordTypePtrType
}

func (in *customHostNameDnsRecordTypePtr) ToCustomHostNameDnsRecordTypePtrOutput() CustomHostNameDnsRecordTypePtrOutput {
	return pulumi.ToOutput(in).(CustomHostNameDnsRecordTypePtrOutput)
}

func (in *customHostNameDnsRecordTypePtr) ToCustomHostNameDnsRecordTypePtrOutputWithContext(ctx context.Context) CustomHostNameDnsRecordTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CustomHostNameDnsRecordTypePtrOutput)
}

type DatabaseType string

const (
	DatabaseTypeSqlAzure   = DatabaseType("SqlAzure")
	DatabaseTypeMySql      = DatabaseType("MySql")
	DatabaseTypeLocalMySql = DatabaseType("LocalMySql")
	DatabaseTypePostgreSql = DatabaseType("PostgreSql")
)

func (DatabaseType) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseType)(nil)).Elem()
}

func (e DatabaseType) ToDatabaseTypeOutput() DatabaseTypeOutput {
	return pulumi.ToOutput(e).(DatabaseTypeOutput)
}

func (e DatabaseType) ToDatabaseTypeOutputWithContext(ctx context.Context) DatabaseTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DatabaseTypeOutput)
}

func (e DatabaseType) ToDatabaseTypePtrOutput() DatabaseTypePtrOutput {
	return e.ToDatabaseTypePtrOutputWithContext(context.Background())
}

func (e DatabaseType) ToDatabaseTypePtrOutputWithContext(ctx context.Context) DatabaseTypePtrOutput {
	return DatabaseType(e).ToDatabaseTypeOutputWithContext(ctx).ToDatabaseTypePtrOutputWithContext(ctx)
}

func (e DatabaseType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DatabaseType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DatabaseType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DatabaseType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DatabaseTypeOutput struct{ *pulumi.OutputState }

func (DatabaseTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseType)(nil)).Elem()
}

func (o DatabaseTypeOutput) ToDatabaseTypeOutput() DatabaseTypeOutput {
	return o
}

func (o DatabaseTypeOutput) ToDatabaseTypeOutputWithContext(ctx context.Context) DatabaseTypeOutput {
	return o
}

func (o DatabaseTypeOutput) ToDatabaseTypePtrOutput() DatabaseTypePtrOutput {
	return o.ToDatabaseTypePtrOutputWithContext(context.Background())
}

func (o DatabaseTypeOutput) ToDatabaseTypePtrOutputWithContext(ctx context.Context) DatabaseTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DatabaseType) *DatabaseType {
		return &v
	}).(DatabaseTypePtrOutput)
}

func (o DatabaseTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DatabaseTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DatabaseType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DatabaseTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DatabaseTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DatabaseType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DatabaseTypePtrOutput struct{ *pulumi.OutputState }

func (DatabaseTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseType)(nil)).Elem()
}

func (o DatabaseTypePtrOutput) ToDatabaseTypePtrOutput() DatabaseTypePtrOutput {
	return o
}

func (o DatabaseTypePtrOutput) ToDatabaseTypePtrOutputWithContext(ctx context.Context) DatabaseTypePtrOutput {
	return o
}

func (o DatabaseTypePtrOutput) Elem() DatabaseTypeOutput {
	return o.ApplyT(func(v *DatabaseType) DatabaseType {
		if v != nil {
			return *v
		}
		var ret DatabaseType
		return ret
	}).(DatabaseTypeOutput)
}

func (o DatabaseTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DatabaseTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DatabaseType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DatabaseTypeInput interface {
	pulumi.Input

	ToDatabaseTypeOutput() DatabaseTypeOutput
	ToDatabaseTypeOutputWithContext(context.Context) DatabaseTypeOutput
}

var databaseTypePtrType = reflect.TypeOf((**DatabaseType)(nil)).Elem()

type DatabaseTypePtrInput interface {
	pulumi.Input

	ToDatabaseTypePtrOutput() DatabaseTypePtrOutput
	ToDatabaseTypePtrOutputWithContext(context.Context) DatabaseTypePtrOutput
}

type databaseTypePtr string

func DatabaseTypePtr(v string) DatabaseTypePtrInput {
	return (*databaseTypePtr)(&v)
}

func (*databaseTypePtr) ElementType() reflect.Type {
	return databaseTypePtrType
}

func (in *databaseTypePtr) ToDatabaseTypePtrOutput() DatabaseTypePtrOutput {
	return pulumi.ToOutput(in).(DatabaseTypePtrOutput)
}

func (in *databaseTypePtr) ToDatabaseTypePtrOutputWithContext(ctx context.Context) DatabaseTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DatabaseTypePtrOutput)
}

type FrequencyUnit string

const (
	FrequencyUnitDay  = FrequencyUnit("Day")
	FrequencyUnitHour = FrequencyUnit("Hour")
)

func (FrequencyUnit) ElementType() reflect.Type {
	return reflect.TypeOf((*FrequencyUnit)(nil)).Elem()
}

func (e FrequencyUnit) ToFrequencyUnitOutput() FrequencyUnitOutput {
	return pulumi.ToOutput(e).(FrequencyUnitOutput)
}

func (e FrequencyUnit) ToFrequencyUnitOutputWithContext(ctx context.Context) FrequencyUnitOutput {
	return pulumi.ToOutputWithContext(ctx, e).(FrequencyUnitOutput)
}

func (e FrequencyUnit) ToFrequencyUnitPtrOutput() FrequencyUnitPtrOutput {
	return e.ToFrequencyUnitPtrOutputWithContext(context.Background())
}

func (e FrequencyUnit) ToFrequencyUnitPtrOutputWithContext(ctx context.Context) FrequencyUnitPtrOutput {
	return FrequencyUnit(e).ToFrequencyUnitOutputWithContext(ctx).ToFrequencyUnitPtrOutputWithContext(ctx)
}

func (e FrequencyUnit) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e FrequencyUnit) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e FrequencyUnit) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e FrequencyUnit) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type FrequencyUnitOutput struct{ *pulumi.OutputState }

func (FrequencyUnitOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrequencyUnit)(nil)).Elem()
}

func (o FrequencyUnitOutput) ToFrequencyUnitOutput() FrequencyUnitOutput {
	return o
}

func (o FrequencyUnitOutput) ToFrequencyUnitOutputWithContext(ctx context.Context) FrequencyUnitOutput {
	return o
}

func (o FrequencyUnitOutput) ToFrequencyUnitPtrOutput() FrequencyUnitPtrOutput {
	return o.ToFrequencyUnitPtrOutputWithContext(context.Background())
}

func (o FrequencyUnitOutput) ToFrequencyUnitPtrOutputWithContext(ctx context.Context) FrequencyUnitPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FrequencyUnit) *FrequencyUnit {
		return &v
	}).(FrequencyUnitPtrOutput)
}

func (o FrequencyUnitOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o FrequencyUnitOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FrequencyUnit) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o FrequencyUnitOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FrequencyUnitOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FrequencyUnit) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type FrequencyUnitPtrOutput struct{ *pulumi.OutputState }

func (FrequencyUnitPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FrequencyUnit)(nil)).Elem()
}

func (o FrequencyUnitPtrOutput) ToFrequencyUnitPtrOutput() FrequencyUnitPtrOutput {
	return o
}

func (o FrequencyUnitPtrOutput) ToFrequencyUnitPtrOutputWithContext(ctx context.Context) FrequencyUnitPtrOutput {
	return o
}

func (o FrequencyUnitPtrOutput) Elem() FrequencyUnitOutput {
	return o.ApplyT(func(v *FrequencyUnit) FrequencyUnit {
		if v != nil {
			return *v
		}
		var ret FrequencyUnit
		return ret
	}).(FrequencyUnitOutput)
}

func (o FrequencyUnitPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FrequencyUnitPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *FrequencyUnit) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type FrequencyUnitInput interface {
	pulumi.Input

	ToFrequencyUnitOutput() FrequencyUnitOutput
	ToFrequencyUnitOutputWithContext(context.Context) FrequencyUnitOutput
}

var frequencyUnitPtrType = reflect.TypeOf((**FrequencyUnit)(nil)).Elem()

type FrequencyUnitPtrInput interface {
	pulumi.Input

	ToFrequencyUnitPtrOutput() FrequencyUnitPtrOutput
	ToFrequencyUnitPtrOutputWithContext(context.Context) FrequencyUnitPtrOutput
}

type frequencyUnitPtr string

func FrequencyUnitPtr(v string) FrequencyUnitPtrInput {
	return (*frequencyUnitPtr)(&v)
}

func (*frequencyUnitPtr) ElementType() reflect.Type {
	return frequencyUnitPtrType
}

func (in *frequencyUnitPtr) ToFrequencyUnitPtrOutput() FrequencyUnitPtrOutput {
	return pulumi.ToOutput(in).(FrequencyUnitPtrOutput)
}

func (in *frequencyUnitPtr) ToFrequencyUnitPtrOutputWithContext(ctx context.Context) FrequencyUnitPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(FrequencyUnitPtrOutput)
}

type HostNameType string

const (
	HostNameTypeVerified = HostNameType("Verified")
	HostNameTypeManaged  = HostNameType("Managed")
)

func (HostNameType) ElementType() reflect.Type {
	return reflect.TypeOf((*HostNameType)(nil)).Elem()
}

func (e HostNameType) ToHostNameTypeOutput() HostNameTypeOutput {
	return pulumi.ToOutput(e).(HostNameTypeOutput)
}

func (e HostNameType) ToHostNameTypeOutputWithContext(ctx context.Context) HostNameTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(HostNameTypeOutput)
}

func (e HostNameType) ToHostNameTypePtrOutput() HostNameTypePtrOutput {
	return e.ToHostNameTypePtrOutputWithContext(context.Background())
}

func (e HostNameType) ToHostNameTypePtrOutputWithContext(ctx context.Context) HostNameTypePtrOutput {
	return HostNameType(e).ToHostNameTypeOutputWithContext(ctx).ToHostNameTypePtrOutputWithContext(ctx)
}

func (e HostNameType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e HostNameType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e HostNameType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e HostNameType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type HostNameTypeOutput struct{ *pulumi.OutputState }

func (HostNameTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HostNameType)(nil)).Elem()
}

func (o HostNameTypeOutput) ToHostNameTypeOutput() HostNameTypeOutput {
	return o
}

func (o HostNameTypeOutput) ToHostNameTypeOutputWithContext(ctx context.Context) HostNameTypeOutput {
	return o
}

func (o HostNameTypeOutput) ToHostNameTypePtrOutput() HostNameTypePtrOutput {
	return o.ToHostNameTypePtrOutputWithContext(context.Background())
}

func (o HostNameTypeOutput) ToHostNameTypePtrOutputWithContext(ctx context.Context) HostNameTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HostNameType) *HostNameType {
		return &v
	}).(HostNameTypePtrOutput)
}

func (o HostNameTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o HostNameTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HostNameType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o HostNameTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HostNameTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HostNameType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type HostNameTypePtrOutput struct{ *pulumi.OutputState }

func (HostNameTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HostNameType)(nil)).Elem()
}

func (o HostNameTypePtrOutput) ToHostNameTypePtrOutput() HostNameTypePtrOutput {
	return o
}

func (o HostNameTypePtrOutput) ToHostNameTypePtrOutputWithContext(ctx context.Context) HostNameTypePtrOutput {
	return o
}

func (o HostNameTypePtrOutput) Elem() HostNameTypeOutput {
	return o.ApplyT(func(v *HostNameType) HostNameType {
		if v != nil {
			return *v
		}
		var ret HostNameType
		return ret
	}).(HostNameTypeOutput)
}

func (o HostNameTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HostNameTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *HostNameType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type HostNameTypeInput interface {
	pulumi.Input

	ToHostNameTypeOutput() HostNameTypeOutput
	ToHostNameTypeOutputWithContext(context.Context) HostNameTypeOutput
}

var hostNameTypePtrType = reflect.TypeOf((**HostNameType)(nil)).Elem()

type HostNameTypePtrInput interface {
	pulumi.Input

	ToHostNameTypePtrOutput() HostNameTypePtrOutput
	ToHostNameTypePtrOutputWithContext(context.Context) HostNameTypePtrOutput
}

type hostNameTypePtr string

func HostNameTypePtr(v string) HostNameTypePtrInput {
	return (*hostNameTypePtr)(&v)
}

func (*hostNameTypePtr) ElementType() reflect.Type {
	return hostNameTypePtrType
}

func (in *hostNameTypePtr) ToHostNameTypePtrOutput() HostNameTypePtrOutput {
	return pulumi.ToOutput(in).(HostNameTypePtrOutput)
}

func (in *hostNameTypePtr) ToHostNameTypePtrOutputWithContext(ctx context.Context) HostNameTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(HostNameTypePtrOutput)
}

type HostType string

const (
	HostTypeStandard   = HostType("Standard")
	HostTypeRepository = HostType("Repository")
)

func (HostType) ElementType() reflect.Type {
	return reflect.TypeOf((*HostType)(nil)).Elem()
}

func (e HostType) ToHostTypeOutput() HostTypeOutput {
	return pulumi.ToOutput(e).(HostTypeOutput)
}

func (e HostType) ToHostTypeOutputWithContext(ctx context.Context) HostTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(HostTypeOutput)
}

func (e HostType) ToHostTypePtrOutput() HostTypePtrOutput {
	return e.ToHostTypePtrOutputWithContext(context.Background())
}

func (e HostType) ToHostTypePtrOutputWithContext(ctx context.Context) HostTypePtrOutput {
	return HostType(e).ToHostTypeOutputWithContext(ctx).ToHostTypePtrOutputWithContext(ctx)
}

func (e HostType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e HostType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e HostType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e HostType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type HostTypeOutput struct{ *pulumi.OutputState }

func (HostTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HostType)(nil)).Elem()
}

func (o HostTypeOutput) ToHostTypeOutput() HostTypeOutput {
	return o
}

func (o HostTypeOutput) ToHostTypeOutputWithContext(ctx context.Context) HostTypeOutput {
	return o
}

func (o HostTypeOutput) ToHostTypePtrOutput() HostTypePtrOutput {
	return o.ToHostTypePtrOutputWithContext(context.Background())
}

func (o HostTypeOutput) ToHostTypePtrOutputWithContext(ctx context.Context) HostTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HostType) *HostType {
		return &v
	}).(HostTypePtrOutput)
}

func (o HostTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o HostTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HostType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o HostTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HostTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HostType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type HostTypePtrOutput struct{ *pulumi.OutputState }

func (HostTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HostType)(nil)).Elem()
}

func (o HostTypePtrOutput) ToHostTypePtrOutput() HostTypePtrOutput {
	return o
}

func (o HostTypePtrOutput) ToHostTypePtrOutputWithContext(ctx context.Context) HostTypePtrOutput {
	return o
}

func (o HostTypePtrOutput) Elem() HostTypeOutput {
	return o.ApplyT(func(v *HostType) HostType {
		if v != nil {
			return *v
		}
		var ret HostType
		return ret
	}).(HostTypeOutput)
}

func (o HostTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HostTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *HostType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type HostTypeInput interface {
	pulumi.Input

	ToHostTypeOutput() HostTypeOutput
	ToHostTypeOutputWithContext(context.Context) HostTypeOutput
}

var hostTypePtrType = reflect.TypeOf((**HostType)(nil)).Elem()

type HostTypePtrInput interface {
	pulumi.Input

	ToHostTypePtrOutput() HostTypePtrOutput
	ToHostTypePtrOutputWithContext(context.Context) HostTypePtrOutput
}

type hostTypePtr string

func HostTypePtr(v string) HostTypePtrInput {
	return (*hostTypePtr)(&v)
}

func (*hostTypePtr) ElementType() reflect.Type {
	return hostTypePtrType
}

func (in *hostTypePtr) ToHostTypePtrOutput() HostTypePtrOutput {
	return pulumi.ToOutput(in).(HostTypePtrOutput)
}

func (in *hostTypePtr) ToHostTypePtrOutputWithContext(ctx context.Context) HostTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(HostTypePtrOutput)
}

type LogLevel string

const (
	LogLevelOff         = LogLevel("Off")
	LogLevelVerbose     = LogLevel("Verbose")
	LogLevelInformation = LogLevel("Information")
	LogLevelWarning     = LogLevel("Warning")
	LogLevelError       = LogLevel("Error")
)

func (LogLevel) ElementType() reflect.Type {
	return reflect.TypeOf((*LogLevel)(nil)).Elem()
}

func (e LogLevel) ToLogLevelOutput() LogLevelOutput {
	return pulumi.ToOutput(e).(LogLevelOutput)
}

func (e LogLevel) ToLogLevelOutputWithContext(ctx context.Context) LogLevelOutput {
	return pulumi.ToOutputWithContext(ctx, e).(LogLevelOutput)
}

func (e LogLevel) ToLogLevelPtrOutput() LogLevelPtrOutput {
	return e.ToLogLevelPtrOutputWithContext(context.Background())
}

func (e LogLevel) ToLogLevelPtrOutputWithContext(ctx context.Context) LogLevelPtrOutput {
	return LogLevel(e).ToLogLevelOutputWithContext(ctx).ToLogLevelPtrOutputWithContext(ctx)
}

func (e LogLevel) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e LogLevel) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e LogLevel) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e LogLevel) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type LogLevelOutput struct{ *pulumi.OutputState }

func (LogLevelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogLevel)(nil)).Elem()
}

func (o LogLevelOutput) ToLogLevelOutput() LogLevelOutput {
	return o
}

func (o LogLevelOutput) ToLogLevelOutputWithContext(ctx context.Context) LogLevelOutput {
	return o
}

func (o LogLevelOutput) ToLogLevelPtrOutput() LogLevelPtrOutput {
	return o.ToLogLevelPtrOutputWithContext(context.Background())
}

func (o LogLevelOutput) ToLogLevelPtrOutputWithContext(ctx context.Context) LogLevelPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LogLevel) *LogLevel {
		return &v
	}).(LogLevelPtrOutput)
}

func (o LogLevelOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o LogLevelOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LogLevel) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o LogLevelOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LogLevelOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LogLevel) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type LogLevelPtrOutput struct{ *pulumi.OutputState }

func (LogLevelPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogLevel)(nil)).Elem()
}

func (o LogLevelPtrOutput) ToLogLevelPtrOutput() LogLevelPtrOutput {
	return o
}

func (o LogLevelPtrOutput) ToLogLevelPtrOutputWithContext(ctx context.Context) LogLevelPtrOutput {
	return o
}

func (o LogLevelPtrOutput) Elem() LogLevelOutput {
	return o.ApplyT(func(v *LogLevel) LogLevel {
		if v != nil {
			return *v
		}
		var ret LogLevel
		return ret
	}).(LogLevelOutput)
}

func (o LogLevelPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LogLevelPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *LogLevel) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type LogLevelInput interface {
	pulumi.Input

	ToLogLevelOutput() LogLevelOutput
	ToLogLevelOutputWithContext(context.Context) LogLevelOutput
}

var logLevelPtrType = reflect.TypeOf((**LogLevel)(nil)).Elem()

type LogLevelPtrInput interface {
	pulumi.Input

	ToLogLevelPtrOutput() LogLevelPtrOutput
	ToLogLevelPtrOutputWithContext(context.Context) LogLevelPtrOutput
}

type logLevelPtr string

func LogLevelPtr(v string) LogLevelPtrInput {
	return (*logLevelPtr)(&v)
}

func (*logLevelPtr) ElementType() reflect.Type {
	return logLevelPtrType
}

func (in *logLevelPtr) ToLogLevelPtrOutput() LogLevelPtrOutput {
	return pulumi.ToOutput(in).(LogLevelPtrOutput)
}

func (in *logLevelPtr) ToLogLevelPtrOutputWithContext(ctx context.Context) LogLevelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(LogLevelPtrOutput)
}

type ManagedPipelineMode string

const (
	ManagedPipelineModeIntegrated = ManagedPipelineMode("Integrated")
	ManagedPipelineModeClassic    = ManagedPipelineMode("Classic")
)

func (ManagedPipelineMode) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedPipelineMode)(nil)).Elem()
}

func (e ManagedPipelineMode) ToManagedPipelineModeOutput() ManagedPipelineModeOutput {
	return pulumi.ToOutput(e).(ManagedPipelineModeOutput)
}

func (e ManagedPipelineMode) ToManagedPipelineModeOutputWithContext(ctx context.Context) ManagedPipelineModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ManagedPipelineModeOutput)
}

func (e ManagedPipelineMode) ToManagedPipelineModePtrOutput() ManagedPipelineModePtrOutput {
	return e.ToManagedPipelineModePtrOutputWithContext(context.Background())
}

func (e ManagedPipelineMode) ToManagedPipelineModePtrOutputWithContext(ctx context.Context) ManagedPipelineModePtrOutput {
	return ManagedPipelineMode(e).ToManagedPipelineModeOutputWithContext(ctx).ToManagedPipelineModePtrOutputWithContext(ctx)
}

func (e ManagedPipelineMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedPipelineMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedPipelineMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ManagedPipelineMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ManagedPipelineModeOutput struct{ *pulumi.OutputState }

func (ManagedPipelineModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedPipelineMode)(nil)).Elem()
}

func (o ManagedPipelineModeOutput) ToManagedPipelineModeOutput() ManagedPipelineModeOutput {
	return o
}

func (o ManagedPipelineModeOutput) ToManagedPipelineModeOutputWithContext(ctx context.Context) ManagedPipelineModeOutput {
	return o
}

func (o ManagedPipelineModeOutput) ToManagedPipelineModePtrOutput() ManagedPipelineModePtrOutput {
	return o.ToManagedPipelineModePtrOutputWithContext(context.Background())
}

func (o ManagedPipelineModeOutput) ToManagedPipelineModePtrOutputWithContext(ctx context.Context) ManagedPipelineModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedPipelineMode) *ManagedPipelineMode {
		return &v
	}).(ManagedPipelineModePtrOutput)
}

func (o ManagedPipelineModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ManagedPipelineModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedPipelineMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ManagedPipelineModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedPipelineModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedPipelineMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ManagedPipelineModePtrOutput struct{ *pulumi.OutputState }

func (ManagedPipelineModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedPipelineMode)(nil)).Elem()
}

func (o ManagedPipelineModePtrOutput) ToManagedPipelineModePtrOutput() ManagedPipelineModePtrOutput {
	return o
}

func (o ManagedPipelineModePtrOutput) ToManagedPipelineModePtrOutputWithContext(ctx context.Context) ManagedPipelineModePtrOutput {
	return o
}

func (o ManagedPipelineModePtrOutput) Elem() ManagedPipelineModeOutput {
	return o.ApplyT(func(v *ManagedPipelineMode) ManagedPipelineMode {
		if v != nil {
			return *v
		}
		var ret ManagedPipelineMode
		return ret
	}).(ManagedPipelineModeOutput)
}

func (o ManagedPipelineModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedPipelineModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ManagedPipelineMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ManagedPipelineModeInput interface {
	pulumi.Input

	ToManagedPipelineModeOutput() ManagedPipelineModeOutput
	ToManagedPipelineModeOutputWithContext(context.Context) ManagedPipelineModeOutput
}

var managedPipelineModePtrType = reflect.TypeOf((**ManagedPipelineMode)(nil)).Elem()

type ManagedPipelineModePtrInput interface {
	pulumi.Input

	ToManagedPipelineModePtrOutput() ManagedPipelineModePtrOutput
	ToManagedPipelineModePtrOutputWithContext(context.Context) ManagedPipelineModePtrOutput
}

type managedPipelineModePtr string

func ManagedPipelineModePtr(v string) ManagedPipelineModePtrInput {
	return (*managedPipelineModePtr)(&v)
}

func (*managedPipelineModePtr) ElementType() reflect.Type {
	return managedPipelineModePtrType
}

func (in *managedPipelineModePtr) ToManagedPipelineModePtrOutput() ManagedPipelineModePtrOutput {
	return pulumi.ToOutput(in).(ManagedPipelineModePtrOutput)
}

func (in *managedPipelineModePtr) ToManagedPipelineModePtrOutputWithContext(ctx context.Context) ManagedPipelineModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ManagedPipelineModePtrOutput)
}

type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeSystemAssigned = ManagedServiceIdentityType("SystemAssigned")
)

func (ManagedServiceIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentityType)(nil)).Elem()
}

func (e ManagedServiceIdentityType) ToManagedServiceIdentityTypeOutput() ManagedServiceIdentityTypeOutput {
	return pulumi.ToOutput(e).(ManagedServiceIdentityTypeOutput)
}

func (e ManagedServiceIdentityType) ToManagedServiceIdentityTypeOutputWithContext(ctx context.Context) ManagedServiceIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ManagedServiceIdentityTypeOutput)
}

func (e ManagedServiceIdentityType) ToManagedServiceIdentityTypePtrOutput() ManagedServiceIdentityTypePtrOutput {
	return e.ToManagedServiceIdentityTypePtrOutputWithContext(context.Background())
}

func (e ManagedServiceIdentityType) ToManagedServiceIdentityTypePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityTypePtrOutput {
	return ManagedServiceIdentityType(e).ToManagedServiceIdentityTypeOutputWithContext(ctx).ToManagedServiceIdentityTypePtrOutputWithContext(ctx)
}

func (e ManagedServiceIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedServiceIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedServiceIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ManagedServiceIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ManagedServiceIdentityTypeOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentityType)(nil)).Elem()
}

func (o ManagedServiceIdentityTypeOutput) ToManagedServiceIdentityTypeOutput() ManagedServiceIdentityTypeOutput {
	return o
}

func (o ManagedServiceIdentityTypeOutput) ToManagedServiceIdentityTypeOutputWithContext(ctx context.Context) ManagedServiceIdentityTypeOutput {
	return o
}

func (o ManagedServiceIdentityTypeOutput) ToManagedServiceIdentityTypePtrOutput() ManagedServiceIdentityTypePtrOutput {
	return o.ToManagedServiceIdentityTypePtrOutputWithContext(context.Background())
}

func (o ManagedServiceIdentityTypeOutput) ToManagedServiceIdentityTypePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedServiceIdentityType) *ManagedServiceIdentityType {
		return &v
	}).(ManagedServiceIdentityTypePtrOutput)
}

func (o ManagedServiceIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ManagedServiceIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedServiceIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedServiceIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedServiceIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ManagedServiceIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentityType)(nil)).Elem()
}

func (o ManagedServiceIdentityTypePtrOutput) ToManagedServiceIdentityTypePtrOutput() ManagedServiceIdentityTypePtrOutput {
	return o
}

func (o ManagedServiceIdentityTypePtrOutput) ToManagedServiceIdentityTypePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityTypePtrOutput {
	return o
}

func (o ManagedServiceIdentityTypePtrOutput) Elem() ManagedServiceIdentityTypeOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityType) ManagedServiceIdentityType {
		if v != nil {
			return *v
		}
		var ret ManagedServiceIdentityType
		return ret
	}).(ManagedServiceIdentityTypeOutput)
}

func (o ManagedServiceIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedServiceIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ManagedServiceIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ManagedServiceIdentityTypeInput interface {
	pulumi.Input

	ToManagedServiceIdentityTypeOutput() ManagedServiceIdentityTypeOutput
	ToManagedServiceIdentityTypeOutputWithContext(context.Context) ManagedServiceIdentityTypeOutput
}

var managedServiceIdentityTypePtrType = reflect.TypeOf((**ManagedServiceIdentityType)(nil)).Elem()

type ManagedServiceIdentityTypePtrInput interface {
	pulumi.Input

	ToManagedServiceIdentityTypePtrOutput() ManagedServiceIdentityTypePtrOutput
	ToManagedServiceIdentityTypePtrOutputWithContext(context.Context) ManagedServiceIdentityTypePtrOutput
}

type managedServiceIdentityTypePtr string

func ManagedServiceIdentityTypePtr(v string) ManagedServiceIdentityTypePtrInput {
	return (*managedServiceIdentityTypePtr)(&v)
}

func (*managedServiceIdentityTypePtr) ElementType() reflect.Type {
	return managedServiceIdentityTypePtrType
}

func (in *managedServiceIdentityTypePtr) ToManagedServiceIdentityTypePtrOutput() ManagedServiceIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(ManagedServiceIdentityTypePtrOutput)
}

func (in *managedServiceIdentityTypePtr) ToManagedServiceIdentityTypePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ManagedServiceIdentityTypePtrOutput)
}

type PublicCertificateLocation string

const (
	PublicCertificateLocationCurrentUserMy  = PublicCertificateLocation("CurrentUserMy")
	PublicCertificateLocationLocalMachineMy = PublicCertificateLocation("LocalMachineMy")
	PublicCertificateLocationUnknown        = PublicCertificateLocation("Unknown")
)

func (PublicCertificateLocation) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicCertificateLocation)(nil)).Elem()
}

func (e PublicCertificateLocation) ToPublicCertificateLocationOutput() PublicCertificateLocationOutput {
	return pulumi.ToOutput(e).(PublicCertificateLocationOutput)
}

func (e PublicCertificateLocation) ToPublicCertificateLocationOutputWithContext(ctx context.Context) PublicCertificateLocationOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PublicCertificateLocationOutput)
}

func (e PublicCertificateLocation) ToPublicCertificateLocationPtrOutput() PublicCertificateLocationPtrOutput {
	return e.ToPublicCertificateLocationPtrOutputWithContext(context.Background())
}

func (e PublicCertificateLocation) ToPublicCertificateLocationPtrOutputWithContext(ctx context.Context) PublicCertificateLocationPtrOutput {
	return PublicCertificateLocation(e).ToPublicCertificateLocationOutputWithContext(ctx).ToPublicCertificateLocationPtrOutputWithContext(ctx)
}

func (e PublicCertificateLocation) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PublicCertificateLocation) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PublicCertificateLocation) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PublicCertificateLocation) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PublicCertificateLocationOutput struct{ *pulumi.OutputState }

func (PublicCertificateLocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicCertificateLocation)(nil)).Elem()
}

func (o PublicCertificateLocationOutput) ToPublicCertificateLocationOutput() PublicCertificateLocationOutput {
	return o
}

func (o PublicCertificateLocationOutput) ToPublicCertificateLocationOutputWithContext(ctx context.Context) PublicCertificateLocationOutput {
	return o
}

func (o PublicCertificateLocationOutput) ToPublicCertificateLocationPtrOutput() PublicCertificateLocationPtrOutput {
	return o.ToPublicCertificateLocationPtrOutputWithContext(context.Background())
}

func (o PublicCertificateLocationOutput) ToPublicCertificateLocationPtrOutputWithContext(ctx context.Context) PublicCertificateLocationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PublicCertificateLocation) *PublicCertificateLocation {
		return &v
	}).(PublicCertificateLocationPtrOutput)
}

func (o PublicCertificateLocationOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PublicCertificateLocationOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PublicCertificateLocation) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PublicCertificateLocationOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PublicCertificateLocationOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PublicCertificateLocation) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PublicCertificateLocationPtrOutput struct{ *pulumi.OutputState }

func (PublicCertificateLocationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicCertificateLocation)(nil)).Elem()
}

func (o PublicCertificateLocationPtrOutput) ToPublicCertificateLocationPtrOutput() PublicCertificateLocationPtrOutput {
	return o
}

func (o PublicCertificateLocationPtrOutput) ToPublicCertificateLocationPtrOutputWithContext(ctx context.Context) PublicCertificateLocationPtrOutput {
	return o
}

func (o PublicCertificateLocationPtrOutput) Elem() PublicCertificateLocationOutput {
	return o.ApplyT(func(v *PublicCertificateLocation) PublicCertificateLocation {
		if v != nil {
			return *v
		}
		var ret PublicCertificateLocation
		return ret
	}).(PublicCertificateLocationOutput)
}

func (o PublicCertificateLocationPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PublicCertificateLocationPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PublicCertificateLocation) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PublicCertificateLocationInput interface {
	pulumi.Input

	ToPublicCertificateLocationOutput() PublicCertificateLocationOutput
	ToPublicCertificateLocationOutputWithContext(context.Context) PublicCertificateLocationOutput
}

var publicCertificateLocationPtrType = reflect.TypeOf((**PublicCertificateLocation)(nil)).Elem()

type PublicCertificateLocationPtrInput interface {
	pulumi.Input

	ToPublicCertificateLocationPtrOutput() PublicCertificateLocationPtrOutput
	ToPublicCertificateLocationPtrOutputWithContext(context.Context) PublicCertificateLocationPtrOutput
}

type publicCertificateLocationPtr string

func PublicCertificateLocationPtr(v string) PublicCertificateLocationPtrInput {
	return (*publicCertificateLocationPtr)(&v)
}

func (*publicCertificateLocationPtr) ElementType() reflect.Type {
	return publicCertificateLocationPtrType
}

func (in *publicCertificateLocationPtr) ToPublicCertificateLocationPtrOutput() PublicCertificateLocationPtrOutput {
	return pulumi.ToOutput(in).(PublicCertificateLocationPtrOutput)
}

func (in *publicCertificateLocationPtr) ToPublicCertificateLocationPtrOutputWithContext(ctx context.Context) PublicCertificateLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PublicCertificateLocationPtrOutput)
}

type ScmType string

const (
	ScmTypeNone         = ScmType("None")
	ScmTypeDropbox      = ScmType("Dropbox")
	ScmTypeTfs          = ScmType("Tfs")
	ScmTypeLocalGit     = ScmType("LocalGit")
	ScmTypeGitHub       = ScmType("GitHub")
	ScmTypeCodePlexGit  = ScmType("CodePlexGit")
	ScmTypeCodePlexHg   = ScmType("CodePlexHg")
	ScmTypeBitbucketGit = ScmType("BitbucketGit")
	ScmTypeBitbucketHg  = ScmType("BitbucketHg")
	ScmTypeExternalGit  = ScmType("ExternalGit")
	ScmTypeExternalHg   = ScmType("ExternalHg")
	ScmTypeOneDrive     = ScmType("OneDrive")
	ScmTypeVSO          = ScmType("VSO")
)

func (ScmType) ElementType() reflect.Type {
	return reflect.TypeOf((*ScmType)(nil)).Elem()
}

func (e ScmType) ToScmTypeOutput() ScmTypeOutput {
	return pulumi.ToOutput(e).(ScmTypeOutput)
}

func (e ScmType) ToScmTypeOutputWithContext(ctx context.Context) ScmTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ScmTypeOutput)
}

func (e ScmType) ToScmTypePtrOutput() ScmTypePtrOutput {
	return e.ToScmTypePtrOutputWithContext(context.Background())
}

func (e ScmType) ToScmTypePtrOutputWithContext(ctx context.Context) ScmTypePtrOutput {
	return ScmType(e).ToScmTypeOutputWithContext(ctx).ToScmTypePtrOutputWithContext(ctx)
}

func (e ScmType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScmType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScmType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ScmType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ScmTypeOutput struct{ *pulumi.OutputState }

func (ScmTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScmType)(nil)).Elem()
}

func (o ScmTypeOutput) ToScmTypeOutput() ScmTypeOutput {
	return o
}

func (o ScmTypeOutput) ToScmTypeOutputWithContext(ctx context.Context) ScmTypeOutput {
	return o
}

func (o ScmTypeOutput) ToScmTypePtrOutput() ScmTypePtrOutput {
	return o.ToScmTypePtrOutputWithContext(context.Background())
}

func (o ScmTypeOutput) ToScmTypePtrOutputWithContext(ctx context.Context) ScmTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScmType) *ScmType {
		return &v
	}).(ScmTypePtrOutput)
}

func (o ScmTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ScmTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScmType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ScmTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScmTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScmType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ScmTypePtrOutput struct{ *pulumi.OutputState }

func (ScmTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScmType)(nil)).Elem()
}

func (o ScmTypePtrOutput) ToScmTypePtrOutput() ScmTypePtrOutput {
	return o
}

func (o ScmTypePtrOutput) ToScmTypePtrOutputWithContext(ctx context.Context) ScmTypePtrOutput {
	return o
}

func (o ScmTypePtrOutput) Elem() ScmTypeOutput {
	return o.ApplyT(func(v *ScmType) ScmType {
		if v != nil {
			return *v
		}
		var ret ScmType
		return ret
	}).(ScmTypeOutput)
}

func (o ScmTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScmTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ScmType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ScmTypeInput interface {
	pulumi.Input

	ToScmTypeOutput() ScmTypeOutput
	ToScmTypeOutputWithContext(context.Context) ScmTypeOutput
}

var scmTypePtrType = reflect.TypeOf((**ScmType)(nil)).Elem()

type ScmTypePtrInput interface {
	pulumi.Input

	ToScmTypePtrOutput() ScmTypePtrOutput
	ToScmTypePtrOutputWithContext(context.Context) ScmTypePtrOutput
}

type scmTypePtr string

func ScmTypePtr(v string) ScmTypePtrInput {
	return (*scmTypePtr)(&v)
}

func (*scmTypePtr) ElementType() reflect.Type {
	return scmTypePtrType
}

func (in *scmTypePtr) ToScmTypePtrOutput() ScmTypePtrOutput {
	return pulumi.ToOutput(in).(ScmTypePtrOutput)
}

func (in *scmTypePtr) ToScmTypePtrOutputWithContext(ctx context.Context) ScmTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ScmTypePtrOutput)
}

type SiteLoadBalancing string

const (
	SiteLoadBalancingWeightedRoundRobin   = SiteLoadBalancing("WeightedRoundRobin")
	SiteLoadBalancingLeastRequests        = SiteLoadBalancing("LeastRequests")
	SiteLoadBalancingLeastResponseTime    = SiteLoadBalancing("LeastResponseTime")
	SiteLoadBalancingWeightedTotalTraffic = SiteLoadBalancing("WeightedTotalTraffic")
	SiteLoadBalancingRequestHash          = SiteLoadBalancing("RequestHash")
)

func (SiteLoadBalancing) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteLoadBalancing)(nil)).Elem()
}

func (e SiteLoadBalancing) ToSiteLoadBalancingOutput() SiteLoadBalancingOutput {
	return pulumi.ToOutput(e).(SiteLoadBalancingOutput)
}

func (e SiteLoadBalancing) ToSiteLoadBalancingOutputWithContext(ctx context.Context) SiteLoadBalancingOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SiteLoadBalancingOutput)
}

func (e SiteLoadBalancing) ToSiteLoadBalancingPtrOutput() SiteLoadBalancingPtrOutput {
	return e.ToSiteLoadBalancingPtrOutputWithContext(context.Background())
}

func (e SiteLoadBalancing) ToSiteLoadBalancingPtrOutputWithContext(ctx context.Context) SiteLoadBalancingPtrOutput {
	return SiteLoadBalancing(e).ToSiteLoadBalancingOutputWithContext(ctx).ToSiteLoadBalancingPtrOutputWithContext(ctx)
}

func (e SiteLoadBalancing) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SiteLoadBalancing) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SiteLoadBalancing) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SiteLoadBalancing) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SiteLoadBalancingOutput struct{ *pulumi.OutputState }

func (SiteLoadBalancingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteLoadBalancing)(nil)).Elem()
}

func (o SiteLoadBalancingOutput) ToSiteLoadBalancingOutput() SiteLoadBalancingOutput {
	return o
}

func (o SiteLoadBalancingOutput) ToSiteLoadBalancingOutputWithContext(ctx context.Context) SiteLoadBalancingOutput {
	return o
}

func (o SiteLoadBalancingOutput) ToSiteLoadBalancingPtrOutput() SiteLoadBalancingPtrOutput {
	return o.ToSiteLoadBalancingPtrOutputWithContext(context.Background())
}

func (o SiteLoadBalancingOutput) ToSiteLoadBalancingPtrOutputWithContext(ctx context.Context) SiteLoadBalancingPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SiteLoadBalancing) *SiteLoadBalancing {
		return &v
	}).(SiteLoadBalancingPtrOutput)
}

func (o SiteLoadBalancingOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SiteLoadBalancingOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SiteLoadBalancing) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SiteLoadBalancingOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SiteLoadBalancingOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SiteLoadBalancing) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SiteLoadBalancingPtrOutput struct{ *pulumi.OutputState }

func (SiteLoadBalancingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteLoadBalancing)(nil)).Elem()
}

func (o SiteLoadBalancingPtrOutput) ToSiteLoadBalancingPtrOutput() SiteLoadBalancingPtrOutput {
	return o
}

func (o SiteLoadBalancingPtrOutput) ToSiteLoadBalancingPtrOutputWithContext(ctx context.Context) SiteLoadBalancingPtrOutput {
	return o
}

func (o SiteLoadBalancingPtrOutput) Elem() SiteLoadBalancingOutput {
	return o.ApplyT(func(v *SiteLoadBalancing) SiteLoadBalancing {
		if v != nil {
			return *v
		}
		var ret SiteLoadBalancing
		return ret
	}).(SiteLoadBalancingOutput)
}

func (o SiteLoadBalancingPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SiteLoadBalancingPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SiteLoadBalancing) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SiteLoadBalancingInput interface {
	pulumi.Input

	ToSiteLoadBalancingOutput() SiteLoadBalancingOutput
	ToSiteLoadBalancingOutputWithContext(context.Context) SiteLoadBalancingOutput
}

var siteLoadBalancingPtrType = reflect.TypeOf((**SiteLoadBalancing)(nil)).Elem()

type SiteLoadBalancingPtrInput interface {
	pulumi.Input

	ToSiteLoadBalancingPtrOutput() SiteLoadBalancingPtrOutput
	ToSiteLoadBalancingPtrOutputWithContext(context.Context) SiteLoadBalancingPtrOutput
}

type siteLoadBalancingPtr string

func SiteLoadBalancingPtr(v string) SiteLoadBalancingPtrInput {
	return (*siteLoadBalancingPtr)(&v)
}

func (*siteLoadBalancingPtr) ElementType() reflect.Type {
	return siteLoadBalancingPtrType
}

func (in *siteLoadBalancingPtr) ToSiteLoadBalancingPtrOutput() SiteLoadBalancingPtrOutput {
	return pulumi.ToOutput(in).(SiteLoadBalancingPtrOutput)
}

func (in *siteLoadBalancingPtr) ToSiteLoadBalancingPtrOutputWithContext(ctx context.Context) SiteLoadBalancingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SiteLoadBalancingPtrOutput)
}

type SslState string

const (
	SslStateDisabled       = SslState("Disabled")
	SslStateSniEnabled     = SslState("SniEnabled")
	SslStateIpBasedEnabled = SslState("IpBasedEnabled")
)

func (SslState) ElementType() reflect.Type {
	return reflect.TypeOf((*SslState)(nil)).Elem()
}

func (e SslState) ToSslStateOutput() SslStateOutput {
	return pulumi.ToOutput(e).(SslStateOutput)
}

func (e SslState) ToSslStateOutputWithContext(ctx context.Context) SslStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SslStateOutput)
}

func (e SslState) ToSslStatePtrOutput() SslStatePtrOutput {
	return e.ToSslStatePtrOutputWithContext(context.Background())
}

func (e SslState) ToSslStatePtrOutputWithContext(ctx context.Context) SslStatePtrOutput {
	return SslState(e).ToSslStateOutputWithContext(ctx).ToSslStatePtrOutputWithContext(ctx)
}

func (e SslState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SslState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SslState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SslState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SslStateOutput struct{ *pulumi.OutputState }

func (SslStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SslState)(nil)).Elem()
}

func (o SslStateOutput) ToSslStateOutput() SslStateOutput {
	return o
}

func (o SslStateOutput) ToSslStateOutputWithContext(ctx context.Context) SslStateOutput {
	return o
}

func (o SslStateOutput) ToSslStatePtrOutput() SslStatePtrOutput {
	return o.ToSslStatePtrOutputWithContext(context.Background())
}

func (o SslStateOutput) ToSslStatePtrOutputWithContext(ctx context.Context) SslStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SslState) *SslState {
		return &v
	}).(SslStatePtrOutput)
}

func (o SslStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SslStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SslState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SslStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SslStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SslState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SslStatePtrOutput struct{ *pulumi.OutputState }

func (SslStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SslState)(nil)).Elem()
}

func (o SslStatePtrOutput) ToSslStatePtrOutput() SslStatePtrOutput {
	return o
}

func (o SslStatePtrOutput) ToSslStatePtrOutputWithContext(ctx context.Context) SslStatePtrOutput {
	return o
}

func (o SslStatePtrOutput) Elem() SslStateOutput {
	return o.ApplyT(func(v *SslState) SslState {
		if v != nil {
			return *v
		}
		var ret SslState
		return ret
	}).(SslStateOutput)
}

func (o SslStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SslStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SslState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SslStateInput interface {
	pulumi.Input

	ToSslStateOutput() SslStateOutput
	ToSslStateOutputWithContext(context.Context) SslStateOutput
}

var sslStatePtrType = reflect.TypeOf((**SslState)(nil)).Elem()

type SslStatePtrInput interface {
	pulumi.Input

	ToSslStatePtrOutput() SslStatePtrOutput
	ToSslStatePtrOutputWithContext(context.Context) SslStatePtrOutput
}

type sslStatePtr string

func SslStatePtr(v string) SslStatePtrInput {
	return (*sslStatePtr)(&v)
}

func (*sslStatePtr) ElementType() reflect.Type {
	return sslStatePtrType
}

func (in *sslStatePtr) ToSslStatePtrOutput() SslStatePtrOutput {
	return pulumi.ToOutput(in).(SslStatePtrOutput)
}

func (in *sslStatePtr) ToSslStatePtrOutputWithContext(ctx context.Context) SslStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SslStatePtrOutput)
}

type SupportedTlsVersions string

const (
	SupportedTlsVersions_1_0 = SupportedTlsVersions("1.0")
	SupportedTlsVersions_1_1 = SupportedTlsVersions("1.1")
	SupportedTlsVersions_1_2 = SupportedTlsVersions("1.2")
)

func (SupportedTlsVersions) ElementType() reflect.Type {
	return reflect.TypeOf((*SupportedTlsVersions)(nil)).Elem()
}

func (e SupportedTlsVersions) ToSupportedTlsVersionsOutput() SupportedTlsVersionsOutput {
	return pulumi.ToOutput(e).(SupportedTlsVersionsOutput)
}

func (e SupportedTlsVersions) ToSupportedTlsVersionsOutputWithContext(ctx context.Context) SupportedTlsVersionsOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SupportedTlsVersionsOutput)
}

func (e SupportedTlsVersions) ToSupportedTlsVersionsPtrOutput() SupportedTlsVersionsPtrOutput {
	return e.ToSupportedTlsVersionsPtrOutputWithContext(context.Background())
}

func (e SupportedTlsVersions) ToSupportedTlsVersionsPtrOutputWithContext(ctx context.Context) SupportedTlsVersionsPtrOutput {
	return SupportedTlsVersions(e).ToSupportedTlsVersionsOutputWithContext(ctx).ToSupportedTlsVersionsPtrOutputWithContext(ctx)
}

func (e SupportedTlsVersions) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SupportedTlsVersions) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SupportedTlsVersions) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SupportedTlsVersions) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SupportedTlsVersionsOutput struct{ *pulumi.OutputState }

func (SupportedTlsVersionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SupportedTlsVersions)(nil)).Elem()
}

func (o SupportedTlsVersionsOutput) ToSupportedTlsVersionsOutput() SupportedTlsVersionsOutput {
	return o
}

func (o SupportedTlsVersionsOutput) ToSupportedTlsVersionsOutputWithContext(ctx context.Context) SupportedTlsVersionsOutput {
	return o
}

func (o SupportedTlsVersionsOutput) ToSupportedTlsVersionsPtrOutput() SupportedTlsVersionsPtrOutput {
	return o.ToSupportedTlsVersionsPtrOutputWithContext(context.Background())
}

func (o SupportedTlsVersionsOutput) ToSupportedTlsVersionsPtrOutputWithContext(ctx context.Context) SupportedTlsVersionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SupportedTlsVersions) *SupportedTlsVersions {
		return &v
	}).(SupportedTlsVersionsPtrOutput)
}

func (o SupportedTlsVersionsOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SupportedTlsVersionsOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SupportedTlsVersions) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SupportedTlsVersionsOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SupportedTlsVersionsOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SupportedTlsVersions) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SupportedTlsVersionsPtrOutput struct{ *pulumi.OutputState }

func (SupportedTlsVersionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SupportedTlsVersions)(nil)).Elem()
}

func (o SupportedTlsVersionsPtrOutput) ToSupportedTlsVersionsPtrOutput() SupportedTlsVersionsPtrOutput {
	return o
}

func (o SupportedTlsVersionsPtrOutput) ToSupportedTlsVersionsPtrOutputWithContext(ctx context.Context) SupportedTlsVersionsPtrOutput {
	return o
}

func (o SupportedTlsVersionsPtrOutput) Elem() SupportedTlsVersionsOutput {
	return o.ApplyT(func(v *SupportedTlsVersions) SupportedTlsVersions {
		if v != nil {
			return *v
		}
		var ret SupportedTlsVersions
		return ret
	}).(SupportedTlsVersionsOutput)
}

func (o SupportedTlsVersionsPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SupportedTlsVersionsPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SupportedTlsVersions) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SupportedTlsVersionsInput interface {
	pulumi.Input

	ToSupportedTlsVersionsOutput() SupportedTlsVersionsOutput
	ToSupportedTlsVersionsOutputWithContext(context.Context) SupportedTlsVersionsOutput
}

var supportedTlsVersionsPtrType = reflect.TypeOf((**SupportedTlsVersions)(nil)).Elem()

type SupportedTlsVersionsPtrInput interface {
	pulumi.Input

	ToSupportedTlsVersionsPtrOutput() SupportedTlsVersionsPtrOutput
	ToSupportedTlsVersionsPtrOutputWithContext(context.Context) SupportedTlsVersionsPtrOutput
}

type supportedTlsVersionsPtr string

func SupportedTlsVersionsPtr(v string) SupportedTlsVersionsPtrInput {
	return (*supportedTlsVersionsPtr)(&v)
}

func (*supportedTlsVersionsPtr) ElementType() reflect.Type {
	return supportedTlsVersionsPtrType
}

func (in *supportedTlsVersionsPtr) ToSupportedTlsVersionsPtrOutput() SupportedTlsVersionsPtrOutput {
	return pulumi.ToOutput(in).(SupportedTlsVersionsPtrOutput)
}

func (in *supportedTlsVersionsPtr) ToSupportedTlsVersionsPtrOutputWithContext(ctx context.Context) SupportedTlsVersionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SupportedTlsVersionsPtrOutput)
}

type UnauthenticatedClientAction string

const (
	UnauthenticatedClientActionRedirectToLoginPage = UnauthenticatedClientAction("RedirectToLoginPage")
	UnauthenticatedClientActionAllowAnonymous      = UnauthenticatedClientAction("AllowAnonymous")
)

func (UnauthenticatedClientAction) ElementType() reflect.Type {
	return reflect.TypeOf((*UnauthenticatedClientAction)(nil)).Elem()
}

func (e UnauthenticatedClientAction) ToUnauthenticatedClientActionOutput() UnauthenticatedClientActionOutput {
	return pulumi.ToOutput(e).(UnauthenticatedClientActionOutput)
}

func (e UnauthenticatedClientAction) ToUnauthenticatedClientActionOutputWithContext(ctx context.Context) UnauthenticatedClientActionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(UnauthenticatedClientActionOutput)
}

func (e UnauthenticatedClientAction) ToUnauthenticatedClientActionPtrOutput() UnauthenticatedClientActionPtrOutput {
	return e.ToUnauthenticatedClientActionPtrOutputWithContext(context.Background())
}

func (e UnauthenticatedClientAction) ToUnauthenticatedClientActionPtrOutputWithContext(ctx context.Context) UnauthenticatedClientActionPtrOutput {
	return UnauthenticatedClientAction(e).ToUnauthenticatedClientActionOutputWithContext(ctx).ToUnauthenticatedClientActionPtrOutputWithContext(ctx)
}

func (e UnauthenticatedClientAction) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e UnauthenticatedClientAction) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e UnauthenticatedClientAction) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e UnauthenticatedClientAction) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type UnauthenticatedClientActionOutput struct{ *pulumi.OutputState }

func (UnauthenticatedClientActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UnauthenticatedClientAction)(nil)).Elem()
}

func (o UnauthenticatedClientActionOutput) ToUnauthenticatedClientActionOutput() UnauthenticatedClientActionOutput {
	return o
}

func (o UnauthenticatedClientActionOutput) ToUnauthenticatedClientActionOutputWithContext(ctx context.Context) UnauthenticatedClientActionOutput {
	return o
}

func (o UnauthenticatedClientActionOutput) ToUnauthenticatedClientActionPtrOutput() UnauthenticatedClientActionPtrOutput {
	return o.ToUnauthenticatedClientActionPtrOutputWithContext(context.Background())
}

func (o UnauthenticatedClientActionOutput) ToUnauthenticatedClientActionPtrOutputWithContext(ctx context.Context) UnauthenticatedClientActionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UnauthenticatedClientAction) *UnauthenticatedClientAction {
		return &v
	}).(UnauthenticatedClientActionPtrOutput)
}

func (o UnauthenticatedClientActionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o UnauthenticatedClientActionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e UnauthenticatedClientAction) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o UnauthenticatedClientActionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o UnauthenticatedClientActionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e UnauthenticatedClientAction) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type UnauthenticatedClientActionPtrOutput struct{ *pulumi.OutputState }

func (UnauthenticatedClientActionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UnauthenticatedClientAction)(nil)).Elem()
}

func (o UnauthenticatedClientActionPtrOutput) ToUnauthenticatedClientActionPtrOutput() UnauthenticatedClientActionPtrOutput {
	return o
}

func (o UnauthenticatedClientActionPtrOutput) ToUnauthenticatedClientActionPtrOutputWithContext(ctx context.Context) UnauthenticatedClientActionPtrOutput {
	return o
}

func (o UnauthenticatedClientActionPtrOutput) Elem() UnauthenticatedClientActionOutput {
	return o.ApplyT(func(v *UnauthenticatedClientAction) UnauthenticatedClientAction {
		if v != nil {
			return *v
		}
		var ret UnauthenticatedClientAction
		return ret
	}).(UnauthenticatedClientActionOutput)
}

func (o UnauthenticatedClientActionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o UnauthenticatedClientActionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *UnauthenticatedClientAction) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type UnauthenticatedClientActionInput interface {
	pulumi.Input

	ToUnauthenticatedClientActionOutput() UnauthenticatedClientActionOutput
	ToUnauthenticatedClientActionOutputWithContext(context.Context) UnauthenticatedClientActionOutput
}

var unauthenticatedClientActionPtrType = reflect.TypeOf((**UnauthenticatedClientAction)(nil)).Elem()

type UnauthenticatedClientActionPtrInput interface {
	pulumi.Input

	ToUnauthenticatedClientActionPtrOutput() UnauthenticatedClientActionPtrOutput
	ToUnauthenticatedClientActionPtrOutputWithContext(context.Context) UnauthenticatedClientActionPtrOutput
}

type unauthenticatedClientActionPtr string

func UnauthenticatedClientActionPtr(v string) UnauthenticatedClientActionPtrInput {
	return (*unauthenticatedClientActionPtr)(&v)
}

func (*unauthenticatedClientActionPtr) ElementType() reflect.Type {
	return unauthenticatedClientActionPtrType
}

func (in *unauthenticatedClientActionPtr) ToUnauthenticatedClientActionPtrOutput() UnauthenticatedClientActionPtrOutput {
	return pulumi.ToOutput(in).(UnauthenticatedClientActionPtrOutput)
}

func (in *unauthenticatedClientActionPtr) ToUnauthenticatedClientActionPtrOutputWithContext(ctx context.Context) UnauthenticatedClientActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(UnauthenticatedClientActionPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AutoHealActionTypeOutput{})
	pulumi.RegisterOutputType(AutoHealActionTypePtrOutput{})
	pulumi.RegisterOutputType(AzureResourceTypeOutput{})
	pulumi.RegisterOutputType(AzureResourceTypePtrOutput{})
	pulumi.RegisterOutputType(BackupRestoreOperationTypeOutput{})
	pulumi.RegisterOutputType(BackupRestoreOperationTypePtrOutput{})
	pulumi.RegisterOutputType(BuiltInAuthenticationProviderOutput{})
	pulumi.RegisterOutputType(BuiltInAuthenticationProviderPtrOutput{})
	pulumi.RegisterOutputType(ConnectionStringTypeOutput{})
	pulumi.RegisterOutputType(ConnectionStringTypePtrOutput{})
	pulumi.RegisterOutputType(CustomHostNameDnsRecordTypeOutput{})
	pulumi.RegisterOutputType(CustomHostNameDnsRecordTypePtrOutput{})
	pulumi.RegisterOutputType(DatabaseTypeOutput{})
	pulumi.RegisterOutputType(DatabaseTypePtrOutput{})
	pulumi.RegisterOutputType(FrequencyUnitOutput{})
	pulumi.RegisterOutputType(FrequencyUnitPtrOutput{})
	pulumi.RegisterOutputType(HostNameTypeOutput{})
	pulumi.RegisterOutputType(HostNameTypePtrOutput{})
	pulumi.RegisterOutputType(HostTypeOutput{})
	pulumi.RegisterOutputType(HostTypePtrOutput{})
	pulumi.RegisterOutputType(LogLevelOutput{})
	pulumi.RegisterOutputType(LogLevelPtrOutput{})
	pulumi.RegisterOutputType(ManagedPipelineModeOutput{})
	pulumi.RegisterOutputType(ManagedPipelineModePtrOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityTypePtrOutput{})
	pulumi.RegisterOutputType(PublicCertificateLocationOutput{})
	pulumi.RegisterOutputType(PublicCertificateLocationPtrOutput{})
	pulumi.RegisterOutputType(ScmTypeOutput{})
	pulumi.RegisterOutputType(ScmTypePtrOutput{})
	pulumi.RegisterOutputType(SiteLoadBalancingOutput{})
	pulumi.RegisterOutputType(SiteLoadBalancingPtrOutput{})
	pulumi.RegisterOutputType(SslStateOutput{})
	pulumi.RegisterOutputType(SslStatePtrOutput{})
	pulumi.RegisterOutputType(SupportedTlsVersionsOutput{})
	pulumi.RegisterOutputType(SupportedTlsVersionsPtrOutput{})
	pulumi.RegisterOutputType(UnauthenticatedClientActionOutput{})
	pulumi.RegisterOutputType(UnauthenticatedClientActionPtrOutput{})
}
