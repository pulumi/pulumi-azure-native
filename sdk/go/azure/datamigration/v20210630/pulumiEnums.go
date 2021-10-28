


package v20210630

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AuthenticationType string

const (
	AuthenticationTypeNone                      = AuthenticationType("None")
	AuthenticationTypeWindowsAuthentication     = AuthenticationType("WindowsAuthentication")
	AuthenticationTypeSqlAuthentication         = AuthenticationType("SqlAuthentication")
	AuthenticationTypeActiveDirectoryIntegrated = AuthenticationType("ActiveDirectoryIntegrated")
	AuthenticationTypeActiveDirectoryPassword   = AuthenticationType("ActiveDirectoryPassword")
)

func (AuthenticationType) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthenticationType)(nil)).Elem()
}

func (e AuthenticationType) ToAuthenticationTypeOutput() AuthenticationTypeOutput {
	return pulumi.ToOutput(e).(AuthenticationTypeOutput)
}

func (e AuthenticationType) ToAuthenticationTypeOutputWithContext(ctx context.Context) AuthenticationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AuthenticationTypeOutput)
}

func (e AuthenticationType) ToAuthenticationTypePtrOutput() AuthenticationTypePtrOutput {
	return e.ToAuthenticationTypePtrOutputWithContext(context.Background())
}

func (e AuthenticationType) ToAuthenticationTypePtrOutputWithContext(ctx context.Context) AuthenticationTypePtrOutput {
	return AuthenticationType(e).ToAuthenticationTypeOutputWithContext(ctx).ToAuthenticationTypePtrOutputWithContext(ctx)
}

func (e AuthenticationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AuthenticationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AuthenticationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AuthenticationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AuthenticationTypeOutput struct{ *pulumi.OutputState }

func (AuthenticationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthenticationType)(nil)).Elem()
}

func (o AuthenticationTypeOutput) ToAuthenticationTypeOutput() AuthenticationTypeOutput {
	return o
}

func (o AuthenticationTypeOutput) ToAuthenticationTypeOutputWithContext(ctx context.Context) AuthenticationTypeOutput {
	return o
}

func (o AuthenticationTypeOutput) ToAuthenticationTypePtrOutput() AuthenticationTypePtrOutput {
	return o.ToAuthenticationTypePtrOutputWithContext(context.Background())
}

func (o AuthenticationTypeOutput) ToAuthenticationTypePtrOutputWithContext(ctx context.Context) AuthenticationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AuthenticationType) *AuthenticationType {
		return &v
	}).(AuthenticationTypePtrOutput)
}

func (o AuthenticationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AuthenticationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AuthenticationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AuthenticationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AuthenticationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AuthenticationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AuthenticationTypePtrOutput struct{ *pulumi.OutputState }

func (AuthenticationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthenticationType)(nil)).Elem()
}

func (o AuthenticationTypePtrOutput) ToAuthenticationTypePtrOutput() AuthenticationTypePtrOutput {
	return o
}

func (o AuthenticationTypePtrOutput) ToAuthenticationTypePtrOutputWithContext(ctx context.Context) AuthenticationTypePtrOutput {
	return o
}

func (o AuthenticationTypePtrOutput) Elem() AuthenticationTypeOutput {
	return o.ApplyT(func(v *AuthenticationType) AuthenticationType {
		if v != nil {
			return *v
		}
		var ret AuthenticationType
		return ret
	}).(AuthenticationTypeOutput)
}

func (o AuthenticationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AuthenticationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AuthenticationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AuthenticationTypeInput interface {
	pulumi.Input

	ToAuthenticationTypeOutput() AuthenticationTypeOutput
	ToAuthenticationTypeOutputWithContext(context.Context) AuthenticationTypeOutput
}

var authenticationTypePtrType = reflect.TypeOf((**AuthenticationType)(nil)).Elem()

type AuthenticationTypePtrInput interface {
	pulumi.Input

	ToAuthenticationTypePtrOutput() AuthenticationTypePtrOutput
	ToAuthenticationTypePtrOutputWithContext(context.Context) AuthenticationTypePtrOutput
}

type authenticationTypePtr string

func AuthenticationTypePtr(v string) AuthenticationTypePtrInput {
	return (*authenticationTypePtr)(&v)
}

func (*authenticationTypePtr) ElementType() reflect.Type {
	return authenticationTypePtrType
}

func (in *authenticationTypePtr) ToAuthenticationTypePtrOutput() AuthenticationTypePtrOutput {
	return pulumi.ToOutput(in).(AuthenticationTypePtrOutput)
}

func (in *authenticationTypePtr) ToAuthenticationTypePtrOutputWithContext(ctx context.Context) AuthenticationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AuthenticationTypePtrOutput)
}

type BackupMode string

const (
	BackupModeCreateBackup   = BackupMode("CreateBackup")
	BackupModeExistingBackup = BackupMode("ExistingBackup")
)

func (BackupMode) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupMode)(nil)).Elem()
}

func (e BackupMode) ToBackupModeOutput() BackupModeOutput {
	return pulumi.ToOutput(e).(BackupModeOutput)
}

func (e BackupMode) ToBackupModeOutputWithContext(ctx context.Context) BackupModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(BackupModeOutput)
}

func (e BackupMode) ToBackupModePtrOutput() BackupModePtrOutput {
	return e.ToBackupModePtrOutputWithContext(context.Background())
}

func (e BackupMode) ToBackupModePtrOutputWithContext(ctx context.Context) BackupModePtrOutput {
	return BackupMode(e).ToBackupModeOutputWithContext(ctx).ToBackupModePtrOutputWithContext(ctx)
}

func (e BackupMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e BackupMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e BackupMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e BackupMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type BackupModeOutput struct{ *pulumi.OutputState }

func (BackupModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupMode)(nil)).Elem()
}

func (o BackupModeOutput) ToBackupModeOutput() BackupModeOutput {
	return o
}

func (o BackupModeOutput) ToBackupModeOutputWithContext(ctx context.Context) BackupModeOutput {
	return o
}

func (o BackupModeOutput) ToBackupModePtrOutput() BackupModePtrOutput {
	return o.ToBackupModePtrOutputWithContext(context.Background())
}

func (o BackupModeOutput) ToBackupModePtrOutputWithContext(ctx context.Context) BackupModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BackupMode) *BackupMode {
		return &v
	}).(BackupModePtrOutput)
}

func (o BackupModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o BackupModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BackupMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o BackupModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BackupModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BackupMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type BackupModePtrOutput struct{ *pulumi.OutputState }

func (BackupModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupMode)(nil)).Elem()
}

func (o BackupModePtrOutput) ToBackupModePtrOutput() BackupModePtrOutput {
	return o
}

func (o BackupModePtrOutput) ToBackupModePtrOutputWithContext(ctx context.Context) BackupModePtrOutput {
	return o
}

func (o BackupModePtrOutput) Elem() BackupModeOutput {
	return o.ApplyT(func(v *BackupMode) BackupMode {
		if v != nil {
			return *v
		}
		var ret BackupMode
		return ret
	}).(BackupModeOutput)
}

func (o BackupModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BackupModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *BackupMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type BackupModeInput interface {
	pulumi.Input

	ToBackupModeOutput() BackupModeOutput
	ToBackupModeOutputWithContext(context.Context) BackupModeOutput
}

var backupModePtrType = reflect.TypeOf((**BackupMode)(nil)).Elem()

type BackupModePtrInput interface {
	pulumi.Input

	ToBackupModePtrOutput() BackupModePtrOutput
	ToBackupModePtrOutputWithContext(context.Context) BackupModePtrOutput
}

type backupModePtr string

func BackupModePtr(v string) BackupModePtrInput {
	return (*backupModePtr)(&v)
}

func (*backupModePtr) ElementType() reflect.Type {
	return backupModePtrType
}

func (in *backupModePtr) ToBackupModePtrOutput() BackupModePtrOutput {
	return pulumi.ToOutput(in).(BackupModePtrOutput)
}

func (in *backupModePtr) ToBackupModePtrOutputWithContext(ctx context.Context) BackupModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(BackupModePtrOutput)
}

type MongoDbReplication string

const (
	MongoDbReplicationDisabled   = MongoDbReplication("Disabled")
	MongoDbReplicationOneTime    = MongoDbReplication("OneTime")
	MongoDbReplicationContinuous = MongoDbReplication("Continuous")
)

func (MongoDbReplication) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoDbReplication)(nil)).Elem()
}

func (e MongoDbReplication) ToMongoDbReplicationOutput() MongoDbReplicationOutput {
	return pulumi.ToOutput(e).(MongoDbReplicationOutput)
}

func (e MongoDbReplication) ToMongoDbReplicationOutputWithContext(ctx context.Context) MongoDbReplicationOutput {
	return pulumi.ToOutputWithContext(ctx, e).(MongoDbReplicationOutput)
}

func (e MongoDbReplication) ToMongoDbReplicationPtrOutput() MongoDbReplicationPtrOutput {
	return e.ToMongoDbReplicationPtrOutputWithContext(context.Background())
}

func (e MongoDbReplication) ToMongoDbReplicationPtrOutputWithContext(ctx context.Context) MongoDbReplicationPtrOutput {
	return MongoDbReplication(e).ToMongoDbReplicationOutputWithContext(ctx).ToMongoDbReplicationPtrOutputWithContext(ctx)
}

func (e MongoDbReplication) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e MongoDbReplication) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e MongoDbReplication) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e MongoDbReplication) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type MongoDbReplicationOutput struct{ *pulumi.OutputState }

func (MongoDbReplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoDbReplication)(nil)).Elem()
}

func (o MongoDbReplicationOutput) ToMongoDbReplicationOutput() MongoDbReplicationOutput {
	return o
}

func (o MongoDbReplicationOutput) ToMongoDbReplicationOutputWithContext(ctx context.Context) MongoDbReplicationOutput {
	return o
}

func (o MongoDbReplicationOutput) ToMongoDbReplicationPtrOutput() MongoDbReplicationPtrOutput {
	return o.ToMongoDbReplicationPtrOutputWithContext(context.Background())
}

func (o MongoDbReplicationOutput) ToMongoDbReplicationPtrOutputWithContext(ctx context.Context) MongoDbReplicationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MongoDbReplication) *MongoDbReplication {
		return &v
	}).(MongoDbReplicationPtrOutput)
}

func (o MongoDbReplicationOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o MongoDbReplicationOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MongoDbReplication) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o MongoDbReplicationOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MongoDbReplicationOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MongoDbReplication) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type MongoDbReplicationPtrOutput struct{ *pulumi.OutputState }

func (MongoDbReplicationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoDbReplication)(nil)).Elem()
}

func (o MongoDbReplicationPtrOutput) ToMongoDbReplicationPtrOutput() MongoDbReplicationPtrOutput {
	return o
}

func (o MongoDbReplicationPtrOutput) ToMongoDbReplicationPtrOutputWithContext(ctx context.Context) MongoDbReplicationPtrOutput {
	return o
}

func (o MongoDbReplicationPtrOutput) Elem() MongoDbReplicationOutput {
	return o.ApplyT(func(v *MongoDbReplication) MongoDbReplication {
		if v != nil {
			return *v
		}
		var ret MongoDbReplication
		return ret
	}).(MongoDbReplicationOutput)
}

func (o MongoDbReplicationPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MongoDbReplicationPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *MongoDbReplication) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type MongoDbReplicationInput interface {
	pulumi.Input

	ToMongoDbReplicationOutput() MongoDbReplicationOutput
	ToMongoDbReplicationOutputWithContext(context.Context) MongoDbReplicationOutput
}

var mongoDbReplicationPtrType = reflect.TypeOf((**MongoDbReplication)(nil)).Elem()

type MongoDbReplicationPtrInput interface {
	pulumi.Input

	ToMongoDbReplicationPtrOutput() MongoDbReplicationPtrOutput
	ToMongoDbReplicationPtrOutputWithContext(context.Context) MongoDbReplicationPtrOutput
}

type mongoDbReplicationPtr string

func MongoDbReplicationPtr(v string) MongoDbReplicationPtrInput {
	return (*mongoDbReplicationPtr)(&v)
}

func (*mongoDbReplicationPtr) ElementType() reflect.Type {
	return mongoDbReplicationPtrType
}

func (in *mongoDbReplicationPtr) ToMongoDbReplicationPtrOutput() MongoDbReplicationPtrOutput {
	return pulumi.ToOutput(in).(MongoDbReplicationPtrOutput)
}

func (in *mongoDbReplicationPtr) ToMongoDbReplicationPtrOutputWithContext(ctx context.Context) MongoDbReplicationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(MongoDbReplicationPtrOutput)
}

type MongoDbShardKeyOrder string

const (
	MongoDbShardKeyOrderForward = MongoDbShardKeyOrder("Forward")
	MongoDbShardKeyOrderReverse = MongoDbShardKeyOrder("Reverse")
	MongoDbShardKeyOrderHashed  = MongoDbShardKeyOrder("Hashed")
)

func (MongoDbShardKeyOrder) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoDbShardKeyOrder)(nil)).Elem()
}

func (e MongoDbShardKeyOrder) ToMongoDbShardKeyOrderOutput() MongoDbShardKeyOrderOutput {
	return pulumi.ToOutput(e).(MongoDbShardKeyOrderOutput)
}

func (e MongoDbShardKeyOrder) ToMongoDbShardKeyOrderOutputWithContext(ctx context.Context) MongoDbShardKeyOrderOutput {
	return pulumi.ToOutputWithContext(ctx, e).(MongoDbShardKeyOrderOutput)
}

func (e MongoDbShardKeyOrder) ToMongoDbShardKeyOrderPtrOutput() MongoDbShardKeyOrderPtrOutput {
	return e.ToMongoDbShardKeyOrderPtrOutputWithContext(context.Background())
}

func (e MongoDbShardKeyOrder) ToMongoDbShardKeyOrderPtrOutputWithContext(ctx context.Context) MongoDbShardKeyOrderPtrOutput {
	return MongoDbShardKeyOrder(e).ToMongoDbShardKeyOrderOutputWithContext(ctx).ToMongoDbShardKeyOrderPtrOutputWithContext(ctx)
}

func (e MongoDbShardKeyOrder) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e MongoDbShardKeyOrder) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e MongoDbShardKeyOrder) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e MongoDbShardKeyOrder) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type MongoDbShardKeyOrderOutput struct{ *pulumi.OutputState }

func (MongoDbShardKeyOrderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoDbShardKeyOrder)(nil)).Elem()
}

func (o MongoDbShardKeyOrderOutput) ToMongoDbShardKeyOrderOutput() MongoDbShardKeyOrderOutput {
	return o
}

func (o MongoDbShardKeyOrderOutput) ToMongoDbShardKeyOrderOutputWithContext(ctx context.Context) MongoDbShardKeyOrderOutput {
	return o
}

func (o MongoDbShardKeyOrderOutput) ToMongoDbShardKeyOrderPtrOutput() MongoDbShardKeyOrderPtrOutput {
	return o.ToMongoDbShardKeyOrderPtrOutputWithContext(context.Background())
}

func (o MongoDbShardKeyOrderOutput) ToMongoDbShardKeyOrderPtrOutputWithContext(ctx context.Context) MongoDbShardKeyOrderPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MongoDbShardKeyOrder) *MongoDbShardKeyOrder {
		return &v
	}).(MongoDbShardKeyOrderPtrOutput)
}

func (o MongoDbShardKeyOrderOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o MongoDbShardKeyOrderOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MongoDbShardKeyOrder) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o MongoDbShardKeyOrderOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MongoDbShardKeyOrderOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MongoDbShardKeyOrder) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type MongoDbShardKeyOrderPtrOutput struct{ *pulumi.OutputState }

func (MongoDbShardKeyOrderPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoDbShardKeyOrder)(nil)).Elem()
}

func (o MongoDbShardKeyOrderPtrOutput) ToMongoDbShardKeyOrderPtrOutput() MongoDbShardKeyOrderPtrOutput {
	return o
}

func (o MongoDbShardKeyOrderPtrOutput) ToMongoDbShardKeyOrderPtrOutputWithContext(ctx context.Context) MongoDbShardKeyOrderPtrOutput {
	return o
}

func (o MongoDbShardKeyOrderPtrOutput) Elem() MongoDbShardKeyOrderOutput {
	return o.ApplyT(func(v *MongoDbShardKeyOrder) MongoDbShardKeyOrder {
		if v != nil {
			return *v
		}
		var ret MongoDbShardKeyOrder
		return ret
	}).(MongoDbShardKeyOrderOutput)
}

func (o MongoDbShardKeyOrderPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MongoDbShardKeyOrderPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *MongoDbShardKeyOrder) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type MongoDbShardKeyOrderInput interface {
	pulumi.Input

	ToMongoDbShardKeyOrderOutput() MongoDbShardKeyOrderOutput
	ToMongoDbShardKeyOrderOutputWithContext(context.Context) MongoDbShardKeyOrderOutput
}

var mongoDbShardKeyOrderPtrType = reflect.TypeOf((**MongoDbShardKeyOrder)(nil)).Elem()

type MongoDbShardKeyOrderPtrInput interface {
	pulumi.Input

	ToMongoDbShardKeyOrderPtrOutput() MongoDbShardKeyOrderPtrOutput
	ToMongoDbShardKeyOrderPtrOutputWithContext(context.Context) MongoDbShardKeyOrderPtrOutput
}

type mongoDbShardKeyOrderPtr string

func MongoDbShardKeyOrderPtr(v string) MongoDbShardKeyOrderPtrInput {
	return (*mongoDbShardKeyOrderPtr)(&v)
}

func (*mongoDbShardKeyOrderPtr) ElementType() reflect.Type {
	return mongoDbShardKeyOrderPtrType
}

func (in *mongoDbShardKeyOrderPtr) ToMongoDbShardKeyOrderPtrOutput() MongoDbShardKeyOrderPtrOutput {
	return pulumi.ToOutput(in).(MongoDbShardKeyOrderPtrOutput)
}

func (in *mongoDbShardKeyOrderPtr) ToMongoDbShardKeyOrderPtrOutputWithContext(ctx context.Context) MongoDbShardKeyOrderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(MongoDbShardKeyOrderPtrOutput)
}

type MySqlTargetPlatformType string

const (
	MySqlTargetPlatformTypeSqlServer       = MySqlTargetPlatformType("SqlServer")
	MySqlTargetPlatformTypeAzureDbForMySQL = MySqlTargetPlatformType("AzureDbForMySQL")
)

func (MySqlTargetPlatformType) ElementType() reflect.Type {
	return reflect.TypeOf((*MySqlTargetPlatformType)(nil)).Elem()
}

func (e MySqlTargetPlatformType) ToMySqlTargetPlatformTypeOutput() MySqlTargetPlatformTypeOutput {
	return pulumi.ToOutput(e).(MySqlTargetPlatformTypeOutput)
}

func (e MySqlTargetPlatformType) ToMySqlTargetPlatformTypeOutputWithContext(ctx context.Context) MySqlTargetPlatformTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(MySqlTargetPlatformTypeOutput)
}

func (e MySqlTargetPlatformType) ToMySqlTargetPlatformTypePtrOutput() MySqlTargetPlatformTypePtrOutput {
	return e.ToMySqlTargetPlatformTypePtrOutputWithContext(context.Background())
}

func (e MySqlTargetPlatformType) ToMySqlTargetPlatformTypePtrOutputWithContext(ctx context.Context) MySqlTargetPlatformTypePtrOutput {
	return MySqlTargetPlatformType(e).ToMySqlTargetPlatformTypeOutputWithContext(ctx).ToMySqlTargetPlatformTypePtrOutputWithContext(ctx)
}

func (e MySqlTargetPlatformType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e MySqlTargetPlatformType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e MySqlTargetPlatformType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e MySqlTargetPlatformType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type MySqlTargetPlatformTypeOutput struct{ *pulumi.OutputState }

func (MySqlTargetPlatformTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MySqlTargetPlatformType)(nil)).Elem()
}

func (o MySqlTargetPlatformTypeOutput) ToMySqlTargetPlatformTypeOutput() MySqlTargetPlatformTypeOutput {
	return o
}

func (o MySqlTargetPlatformTypeOutput) ToMySqlTargetPlatformTypeOutputWithContext(ctx context.Context) MySqlTargetPlatformTypeOutput {
	return o
}

func (o MySqlTargetPlatformTypeOutput) ToMySqlTargetPlatformTypePtrOutput() MySqlTargetPlatformTypePtrOutput {
	return o.ToMySqlTargetPlatformTypePtrOutputWithContext(context.Background())
}

func (o MySqlTargetPlatformTypeOutput) ToMySqlTargetPlatformTypePtrOutputWithContext(ctx context.Context) MySqlTargetPlatformTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MySqlTargetPlatformType) *MySqlTargetPlatformType {
		return &v
	}).(MySqlTargetPlatformTypePtrOutput)
}

func (o MySqlTargetPlatformTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o MySqlTargetPlatformTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MySqlTargetPlatformType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o MySqlTargetPlatformTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MySqlTargetPlatformTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MySqlTargetPlatformType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type MySqlTargetPlatformTypePtrOutput struct{ *pulumi.OutputState }

func (MySqlTargetPlatformTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MySqlTargetPlatformType)(nil)).Elem()
}

func (o MySqlTargetPlatformTypePtrOutput) ToMySqlTargetPlatformTypePtrOutput() MySqlTargetPlatformTypePtrOutput {
	return o
}

func (o MySqlTargetPlatformTypePtrOutput) ToMySqlTargetPlatformTypePtrOutputWithContext(ctx context.Context) MySqlTargetPlatformTypePtrOutput {
	return o
}

func (o MySqlTargetPlatformTypePtrOutput) Elem() MySqlTargetPlatformTypeOutput {
	return o.ApplyT(func(v *MySqlTargetPlatformType) MySqlTargetPlatformType {
		if v != nil {
			return *v
		}
		var ret MySqlTargetPlatformType
		return ret
	}).(MySqlTargetPlatformTypeOutput)
}

func (o MySqlTargetPlatformTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MySqlTargetPlatformTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *MySqlTargetPlatformType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type MySqlTargetPlatformTypeInput interface {
	pulumi.Input

	ToMySqlTargetPlatformTypeOutput() MySqlTargetPlatformTypeOutput
	ToMySqlTargetPlatformTypeOutputWithContext(context.Context) MySqlTargetPlatformTypeOutput
}

var mySqlTargetPlatformTypePtrType = reflect.TypeOf((**MySqlTargetPlatformType)(nil)).Elem()

type MySqlTargetPlatformTypePtrInput interface {
	pulumi.Input

	ToMySqlTargetPlatformTypePtrOutput() MySqlTargetPlatformTypePtrOutput
	ToMySqlTargetPlatformTypePtrOutputWithContext(context.Context) MySqlTargetPlatformTypePtrOutput
}

type mySqlTargetPlatformTypePtr string

func MySqlTargetPlatformTypePtr(v string) MySqlTargetPlatformTypePtrInput {
	return (*mySqlTargetPlatformTypePtr)(&v)
}

func (*mySqlTargetPlatformTypePtr) ElementType() reflect.Type {
	return mySqlTargetPlatformTypePtrType
}

func (in *mySqlTargetPlatformTypePtr) ToMySqlTargetPlatformTypePtrOutput() MySqlTargetPlatformTypePtrOutput {
	return pulumi.ToOutput(in).(MySqlTargetPlatformTypePtrOutput)
}

func (in *mySqlTargetPlatformTypePtr) ToMySqlTargetPlatformTypePtrOutputWithContext(ctx context.Context) MySqlTargetPlatformTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(MySqlTargetPlatformTypePtrOutput)
}

type ProjectSourcePlatform string

const (
	ProjectSourcePlatformSQL        = ProjectSourcePlatform("SQL")
	ProjectSourcePlatformMySQL      = ProjectSourcePlatform("MySQL")
	ProjectSourcePlatformPostgreSql = ProjectSourcePlatform("PostgreSql")
	ProjectSourcePlatformMongoDb    = ProjectSourcePlatform("MongoDb")
	ProjectSourcePlatformUnknown    = ProjectSourcePlatform("Unknown")
)

func (ProjectSourcePlatform) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectSourcePlatform)(nil)).Elem()
}

func (e ProjectSourcePlatform) ToProjectSourcePlatformOutput() ProjectSourcePlatformOutput {
	return pulumi.ToOutput(e).(ProjectSourcePlatformOutput)
}

func (e ProjectSourcePlatform) ToProjectSourcePlatformOutputWithContext(ctx context.Context) ProjectSourcePlatformOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ProjectSourcePlatformOutput)
}

func (e ProjectSourcePlatform) ToProjectSourcePlatformPtrOutput() ProjectSourcePlatformPtrOutput {
	return e.ToProjectSourcePlatformPtrOutputWithContext(context.Background())
}

func (e ProjectSourcePlatform) ToProjectSourcePlatformPtrOutputWithContext(ctx context.Context) ProjectSourcePlatformPtrOutput {
	return ProjectSourcePlatform(e).ToProjectSourcePlatformOutputWithContext(ctx).ToProjectSourcePlatformPtrOutputWithContext(ctx)
}

func (e ProjectSourcePlatform) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProjectSourcePlatform) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProjectSourcePlatform) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ProjectSourcePlatform) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ProjectSourcePlatformOutput struct{ *pulumi.OutputState }

func (ProjectSourcePlatformOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectSourcePlatform)(nil)).Elem()
}

func (o ProjectSourcePlatformOutput) ToProjectSourcePlatformOutput() ProjectSourcePlatformOutput {
	return o
}

func (o ProjectSourcePlatformOutput) ToProjectSourcePlatformOutputWithContext(ctx context.Context) ProjectSourcePlatformOutput {
	return o
}

func (o ProjectSourcePlatformOutput) ToProjectSourcePlatformPtrOutput() ProjectSourcePlatformPtrOutput {
	return o.ToProjectSourcePlatformPtrOutputWithContext(context.Background())
}

func (o ProjectSourcePlatformOutput) ToProjectSourcePlatformPtrOutputWithContext(ctx context.Context) ProjectSourcePlatformPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProjectSourcePlatform) *ProjectSourcePlatform {
		return &v
	}).(ProjectSourcePlatformPtrOutput)
}

func (o ProjectSourcePlatformOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ProjectSourcePlatformOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProjectSourcePlatform) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ProjectSourcePlatformOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProjectSourcePlatformOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProjectSourcePlatform) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ProjectSourcePlatformPtrOutput struct{ *pulumi.OutputState }

func (ProjectSourcePlatformPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectSourcePlatform)(nil)).Elem()
}

func (o ProjectSourcePlatformPtrOutput) ToProjectSourcePlatformPtrOutput() ProjectSourcePlatformPtrOutput {
	return o
}

func (o ProjectSourcePlatformPtrOutput) ToProjectSourcePlatformPtrOutputWithContext(ctx context.Context) ProjectSourcePlatformPtrOutput {
	return o
}

func (o ProjectSourcePlatformPtrOutput) Elem() ProjectSourcePlatformOutput {
	return o.ApplyT(func(v *ProjectSourcePlatform) ProjectSourcePlatform {
		if v != nil {
			return *v
		}
		var ret ProjectSourcePlatform
		return ret
	}).(ProjectSourcePlatformOutput)
}

func (o ProjectSourcePlatformPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProjectSourcePlatformPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ProjectSourcePlatform) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ProjectSourcePlatformInput interface {
	pulumi.Input

	ToProjectSourcePlatformOutput() ProjectSourcePlatformOutput
	ToProjectSourcePlatformOutputWithContext(context.Context) ProjectSourcePlatformOutput
}

var projectSourcePlatformPtrType = reflect.TypeOf((**ProjectSourcePlatform)(nil)).Elem()

type ProjectSourcePlatformPtrInput interface {
	pulumi.Input

	ToProjectSourcePlatformPtrOutput() ProjectSourcePlatformPtrOutput
	ToProjectSourcePlatformPtrOutputWithContext(context.Context) ProjectSourcePlatformPtrOutput
}

type projectSourcePlatformPtr string

func ProjectSourcePlatformPtr(v string) ProjectSourcePlatformPtrInput {
	return (*projectSourcePlatformPtr)(&v)
}

func (*projectSourcePlatformPtr) ElementType() reflect.Type {
	return projectSourcePlatformPtrType
}

func (in *projectSourcePlatformPtr) ToProjectSourcePlatformPtrOutput() ProjectSourcePlatformPtrOutput {
	return pulumi.ToOutput(in).(ProjectSourcePlatformPtrOutput)
}

func (in *projectSourcePlatformPtr) ToProjectSourcePlatformPtrOutputWithContext(ctx context.Context) ProjectSourcePlatformPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ProjectSourcePlatformPtrOutput)
}

type ProjectTargetPlatform string

const (
	ProjectTargetPlatformSQLDB                = ProjectTargetPlatform("SQLDB")
	ProjectTargetPlatformSQLMI                = ProjectTargetPlatform("SQLMI")
	ProjectTargetPlatformAzureDbForMySql      = ProjectTargetPlatform("AzureDbForMySql")
	ProjectTargetPlatformAzureDbForPostgreSql = ProjectTargetPlatform("AzureDbForPostgreSql")
	ProjectTargetPlatformMongoDb              = ProjectTargetPlatform("MongoDb")
	ProjectTargetPlatformUnknown              = ProjectTargetPlatform("Unknown")
)

func (ProjectTargetPlatform) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectTargetPlatform)(nil)).Elem()
}

func (e ProjectTargetPlatform) ToProjectTargetPlatformOutput() ProjectTargetPlatformOutput {
	return pulumi.ToOutput(e).(ProjectTargetPlatformOutput)
}

func (e ProjectTargetPlatform) ToProjectTargetPlatformOutputWithContext(ctx context.Context) ProjectTargetPlatformOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ProjectTargetPlatformOutput)
}

func (e ProjectTargetPlatform) ToProjectTargetPlatformPtrOutput() ProjectTargetPlatformPtrOutput {
	return e.ToProjectTargetPlatformPtrOutputWithContext(context.Background())
}

func (e ProjectTargetPlatform) ToProjectTargetPlatformPtrOutputWithContext(ctx context.Context) ProjectTargetPlatformPtrOutput {
	return ProjectTargetPlatform(e).ToProjectTargetPlatformOutputWithContext(ctx).ToProjectTargetPlatformPtrOutputWithContext(ctx)
}

func (e ProjectTargetPlatform) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProjectTargetPlatform) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProjectTargetPlatform) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ProjectTargetPlatform) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ProjectTargetPlatformOutput struct{ *pulumi.OutputState }

func (ProjectTargetPlatformOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectTargetPlatform)(nil)).Elem()
}

func (o ProjectTargetPlatformOutput) ToProjectTargetPlatformOutput() ProjectTargetPlatformOutput {
	return o
}

func (o ProjectTargetPlatformOutput) ToProjectTargetPlatformOutputWithContext(ctx context.Context) ProjectTargetPlatformOutput {
	return o
}

func (o ProjectTargetPlatformOutput) ToProjectTargetPlatformPtrOutput() ProjectTargetPlatformPtrOutput {
	return o.ToProjectTargetPlatformPtrOutputWithContext(context.Background())
}

func (o ProjectTargetPlatformOutput) ToProjectTargetPlatformPtrOutputWithContext(ctx context.Context) ProjectTargetPlatformPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProjectTargetPlatform) *ProjectTargetPlatform {
		return &v
	}).(ProjectTargetPlatformPtrOutput)
}

func (o ProjectTargetPlatformOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ProjectTargetPlatformOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProjectTargetPlatform) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ProjectTargetPlatformOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProjectTargetPlatformOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProjectTargetPlatform) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ProjectTargetPlatformPtrOutput struct{ *pulumi.OutputState }

func (ProjectTargetPlatformPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectTargetPlatform)(nil)).Elem()
}

func (o ProjectTargetPlatformPtrOutput) ToProjectTargetPlatformPtrOutput() ProjectTargetPlatformPtrOutput {
	return o
}

func (o ProjectTargetPlatformPtrOutput) ToProjectTargetPlatformPtrOutputWithContext(ctx context.Context) ProjectTargetPlatformPtrOutput {
	return o
}

func (o ProjectTargetPlatformPtrOutput) Elem() ProjectTargetPlatformOutput {
	return o.ApplyT(func(v *ProjectTargetPlatform) ProjectTargetPlatform {
		if v != nil {
			return *v
		}
		var ret ProjectTargetPlatform
		return ret
	}).(ProjectTargetPlatformOutput)
}

func (o ProjectTargetPlatformPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProjectTargetPlatformPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ProjectTargetPlatform) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ProjectTargetPlatformInput interface {
	pulumi.Input

	ToProjectTargetPlatformOutput() ProjectTargetPlatformOutput
	ToProjectTargetPlatformOutputWithContext(context.Context) ProjectTargetPlatformOutput
}

var projectTargetPlatformPtrType = reflect.TypeOf((**ProjectTargetPlatform)(nil)).Elem()

type ProjectTargetPlatformPtrInput interface {
	pulumi.Input

	ToProjectTargetPlatformPtrOutput() ProjectTargetPlatformPtrOutput
	ToProjectTargetPlatformPtrOutputWithContext(context.Context) ProjectTargetPlatformPtrOutput
}

type projectTargetPlatformPtr string

func ProjectTargetPlatformPtr(v string) ProjectTargetPlatformPtrInput {
	return (*projectTargetPlatformPtr)(&v)
}

func (*projectTargetPlatformPtr) ElementType() reflect.Type {
	return projectTargetPlatformPtrType
}

func (in *projectTargetPlatformPtr) ToProjectTargetPlatformPtrOutput() ProjectTargetPlatformPtrOutput {
	return pulumi.ToOutput(in).(ProjectTargetPlatformPtrOutput)
}

func (in *projectTargetPlatformPtr) ToProjectTargetPlatformPtrOutputWithContext(ctx context.Context) ProjectTargetPlatformPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ProjectTargetPlatformPtrOutput)
}

type ServerLevelPermissionsGroup string

const (
	ServerLevelPermissionsGroupDefault                             = ServerLevelPermissionsGroup("Default")
	ServerLevelPermissionsGroupMigrationFromSqlServerToAzureDB     = ServerLevelPermissionsGroup("MigrationFromSqlServerToAzureDB")
	ServerLevelPermissionsGroupMigrationFromSqlServerToAzureMI     = ServerLevelPermissionsGroup("MigrationFromSqlServerToAzureMI")
	ServerLevelPermissionsGroupMigrationFromMySQLToAzureDBForMySQL = ServerLevelPermissionsGroup("MigrationFromMySQLToAzureDBForMySQL")
)

func (ServerLevelPermissionsGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerLevelPermissionsGroup)(nil)).Elem()
}

func (e ServerLevelPermissionsGroup) ToServerLevelPermissionsGroupOutput() ServerLevelPermissionsGroupOutput {
	return pulumi.ToOutput(e).(ServerLevelPermissionsGroupOutput)
}

func (e ServerLevelPermissionsGroup) ToServerLevelPermissionsGroupOutputWithContext(ctx context.Context) ServerLevelPermissionsGroupOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ServerLevelPermissionsGroupOutput)
}

func (e ServerLevelPermissionsGroup) ToServerLevelPermissionsGroupPtrOutput() ServerLevelPermissionsGroupPtrOutput {
	return e.ToServerLevelPermissionsGroupPtrOutputWithContext(context.Background())
}

func (e ServerLevelPermissionsGroup) ToServerLevelPermissionsGroupPtrOutputWithContext(ctx context.Context) ServerLevelPermissionsGroupPtrOutput {
	return ServerLevelPermissionsGroup(e).ToServerLevelPermissionsGroupOutputWithContext(ctx).ToServerLevelPermissionsGroupPtrOutputWithContext(ctx)
}

func (e ServerLevelPermissionsGroup) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServerLevelPermissionsGroup) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServerLevelPermissionsGroup) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ServerLevelPermissionsGroup) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ServerLevelPermissionsGroupOutput struct{ *pulumi.OutputState }

func (ServerLevelPermissionsGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerLevelPermissionsGroup)(nil)).Elem()
}

func (o ServerLevelPermissionsGroupOutput) ToServerLevelPermissionsGroupOutput() ServerLevelPermissionsGroupOutput {
	return o
}

func (o ServerLevelPermissionsGroupOutput) ToServerLevelPermissionsGroupOutputWithContext(ctx context.Context) ServerLevelPermissionsGroupOutput {
	return o
}

func (o ServerLevelPermissionsGroupOutput) ToServerLevelPermissionsGroupPtrOutput() ServerLevelPermissionsGroupPtrOutput {
	return o.ToServerLevelPermissionsGroupPtrOutputWithContext(context.Background())
}

func (o ServerLevelPermissionsGroupOutput) ToServerLevelPermissionsGroupPtrOutputWithContext(ctx context.Context) ServerLevelPermissionsGroupPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServerLevelPermissionsGroup) *ServerLevelPermissionsGroup {
		return &v
	}).(ServerLevelPermissionsGroupPtrOutput)
}

func (o ServerLevelPermissionsGroupOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ServerLevelPermissionsGroupOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServerLevelPermissionsGroup) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ServerLevelPermissionsGroupOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServerLevelPermissionsGroupOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServerLevelPermissionsGroup) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ServerLevelPermissionsGroupPtrOutput struct{ *pulumi.OutputState }

func (ServerLevelPermissionsGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerLevelPermissionsGroup)(nil)).Elem()
}

func (o ServerLevelPermissionsGroupPtrOutput) ToServerLevelPermissionsGroupPtrOutput() ServerLevelPermissionsGroupPtrOutput {
	return o
}

func (o ServerLevelPermissionsGroupPtrOutput) ToServerLevelPermissionsGroupPtrOutputWithContext(ctx context.Context) ServerLevelPermissionsGroupPtrOutput {
	return o
}

func (o ServerLevelPermissionsGroupPtrOutput) Elem() ServerLevelPermissionsGroupOutput {
	return o.ApplyT(func(v *ServerLevelPermissionsGroup) ServerLevelPermissionsGroup {
		if v != nil {
			return *v
		}
		var ret ServerLevelPermissionsGroup
		return ret
	}).(ServerLevelPermissionsGroupOutput)
}

func (o ServerLevelPermissionsGroupPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServerLevelPermissionsGroupPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ServerLevelPermissionsGroup) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ServerLevelPermissionsGroupInput interface {
	pulumi.Input

	ToServerLevelPermissionsGroupOutput() ServerLevelPermissionsGroupOutput
	ToServerLevelPermissionsGroupOutputWithContext(context.Context) ServerLevelPermissionsGroupOutput
}

var serverLevelPermissionsGroupPtrType = reflect.TypeOf((**ServerLevelPermissionsGroup)(nil)).Elem()

type ServerLevelPermissionsGroupPtrInput interface {
	pulumi.Input

	ToServerLevelPermissionsGroupPtrOutput() ServerLevelPermissionsGroupPtrOutput
	ToServerLevelPermissionsGroupPtrOutputWithContext(context.Context) ServerLevelPermissionsGroupPtrOutput
}

type serverLevelPermissionsGroupPtr string

func ServerLevelPermissionsGroupPtr(v string) ServerLevelPermissionsGroupPtrInput {
	return (*serverLevelPermissionsGroupPtr)(&v)
}

func (*serverLevelPermissionsGroupPtr) ElementType() reflect.Type {
	return serverLevelPermissionsGroupPtrType
}

func (in *serverLevelPermissionsGroupPtr) ToServerLevelPermissionsGroupPtrOutput() ServerLevelPermissionsGroupPtrOutput {
	return pulumi.ToOutput(in).(ServerLevelPermissionsGroupPtrOutput)
}

func (in *serverLevelPermissionsGroupPtr) ToServerLevelPermissionsGroupPtrOutputWithContext(ctx context.Context) ServerLevelPermissionsGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ServerLevelPermissionsGroupPtrOutput)
}

type SqlSourcePlatform string

const (
	SqlSourcePlatformSqlOnPrem = SqlSourcePlatform("SqlOnPrem")
)

func (SqlSourcePlatform) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlSourcePlatform)(nil)).Elem()
}

func (e SqlSourcePlatform) ToSqlSourcePlatformOutput() SqlSourcePlatformOutput {
	return pulumi.ToOutput(e).(SqlSourcePlatformOutput)
}

func (e SqlSourcePlatform) ToSqlSourcePlatformOutputWithContext(ctx context.Context) SqlSourcePlatformOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SqlSourcePlatformOutput)
}

func (e SqlSourcePlatform) ToSqlSourcePlatformPtrOutput() SqlSourcePlatformPtrOutput {
	return e.ToSqlSourcePlatformPtrOutputWithContext(context.Background())
}

func (e SqlSourcePlatform) ToSqlSourcePlatformPtrOutputWithContext(ctx context.Context) SqlSourcePlatformPtrOutput {
	return SqlSourcePlatform(e).ToSqlSourcePlatformOutputWithContext(ctx).ToSqlSourcePlatformPtrOutputWithContext(ctx)
}

func (e SqlSourcePlatform) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SqlSourcePlatform) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SqlSourcePlatform) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SqlSourcePlatform) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SqlSourcePlatformOutput struct{ *pulumi.OutputState }

func (SqlSourcePlatformOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlSourcePlatform)(nil)).Elem()
}

func (o SqlSourcePlatformOutput) ToSqlSourcePlatformOutput() SqlSourcePlatformOutput {
	return o
}

func (o SqlSourcePlatformOutput) ToSqlSourcePlatformOutputWithContext(ctx context.Context) SqlSourcePlatformOutput {
	return o
}

func (o SqlSourcePlatformOutput) ToSqlSourcePlatformPtrOutput() SqlSourcePlatformPtrOutput {
	return o.ToSqlSourcePlatformPtrOutputWithContext(context.Background())
}

func (o SqlSourcePlatformOutput) ToSqlSourcePlatformPtrOutputWithContext(ctx context.Context) SqlSourcePlatformPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SqlSourcePlatform) *SqlSourcePlatform {
		return &v
	}).(SqlSourcePlatformPtrOutput)
}

func (o SqlSourcePlatformOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SqlSourcePlatformOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SqlSourcePlatform) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SqlSourcePlatformOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SqlSourcePlatformOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SqlSourcePlatform) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SqlSourcePlatformPtrOutput struct{ *pulumi.OutputState }

func (SqlSourcePlatformPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlSourcePlatform)(nil)).Elem()
}

func (o SqlSourcePlatformPtrOutput) ToSqlSourcePlatformPtrOutput() SqlSourcePlatformPtrOutput {
	return o
}

func (o SqlSourcePlatformPtrOutput) ToSqlSourcePlatformPtrOutputWithContext(ctx context.Context) SqlSourcePlatformPtrOutput {
	return o
}

func (o SqlSourcePlatformPtrOutput) Elem() SqlSourcePlatformOutput {
	return o.ApplyT(func(v *SqlSourcePlatform) SqlSourcePlatform {
		if v != nil {
			return *v
		}
		var ret SqlSourcePlatform
		return ret
	}).(SqlSourcePlatformOutput)
}

func (o SqlSourcePlatformPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SqlSourcePlatformPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SqlSourcePlatform) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SqlSourcePlatformInput interface {
	pulumi.Input

	ToSqlSourcePlatformOutput() SqlSourcePlatformOutput
	ToSqlSourcePlatformOutputWithContext(context.Context) SqlSourcePlatformOutput
}

var sqlSourcePlatformPtrType = reflect.TypeOf((**SqlSourcePlatform)(nil)).Elem()

type SqlSourcePlatformPtrInput interface {
	pulumi.Input

	ToSqlSourcePlatformPtrOutput() SqlSourcePlatformPtrOutput
	ToSqlSourcePlatformPtrOutputWithContext(context.Context) SqlSourcePlatformPtrOutput
}

type sqlSourcePlatformPtr string

func SqlSourcePlatformPtr(v string) SqlSourcePlatformPtrInput {
	return (*sqlSourcePlatformPtr)(&v)
}

func (*sqlSourcePlatformPtr) ElementType() reflect.Type {
	return sqlSourcePlatformPtrType
}

func (in *sqlSourcePlatformPtr) ToSqlSourcePlatformPtrOutput() SqlSourcePlatformPtrOutput {
	return pulumi.ToOutput(in).(SqlSourcePlatformPtrOutput)
}

func (in *sqlSourcePlatformPtr) ToSqlSourcePlatformPtrOutputWithContext(ctx context.Context) SqlSourcePlatformPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SqlSourcePlatformPtrOutput)
}

type SsisMigrationOverwriteOption string

const (
	SsisMigrationOverwriteOptionIgnore    = SsisMigrationOverwriteOption("Ignore")
	SsisMigrationOverwriteOptionOverwrite = SsisMigrationOverwriteOption("Overwrite")
)

func (SsisMigrationOverwriteOption) ElementType() reflect.Type {
	return reflect.TypeOf((*SsisMigrationOverwriteOption)(nil)).Elem()
}

func (e SsisMigrationOverwriteOption) ToSsisMigrationOverwriteOptionOutput() SsisMigrationOverwriteOptionOutput {
	return pulumi.ToOutput(e).(SsisMigrationOverwriteOptionOutput)
}

func (e SsisMigrationOverwriteOption) ToSsisMigrationOverwriteOptionOutputWithContext(ctx context.Context) SsisMigrationOverwriteOptionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SsisMigrationOverwriteOptionOutput)
}

func (e SsisMigrationOverwriteOption) ToSsisMigrationOverwriteOptionPtrOutput() SsisMigrationOverwriteOptionPtrOutput {
	return e.ToSsisMigrationOverwriteOptionPtrOutputWithContext(context.Background())
}

func (e SsisMigrationOverwriteOption) ToSsisMigrationOverwriteOptionPtrOutputWithContext(ctx context.Context) SsisMigrationOverwriteOptionPtrOutput {
	return SsisMigrationOverwriteOption(e).ToSsisMigrationOverwriteOptionOutputWithContext(ctx).ToSsisMigrationOverwriteOptionPtrOutputWithContext(ctx)
}

func (e SsisMigrationOverwriteOption) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SsisMigrationOverwriteOption) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SsisMigrationOverwriteOption) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SsisMigrationOverwriteOption) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SsisMigrationOverwriteOptionOutput struct{ *pulumi.OutputState }

func (SsisMigrationOverwriteOptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SsisMigrationOverwriteOption)(nil)).Elem()
}

func (o SsisMigrationOverwriteOptionOutput) ToSsisMigrationOverwriteOptionOutput() SsisMigrationOverwriteOptionOutput {
	return o
}

func (o SsisMigrationOverwriteOptionOutput) ToSsisMigrationOverwriteOptionOutputWithContext(ctx context.Context) SsisMigrationOverwriteOptionOutput {
	return o
}

func (o SsisMigrationOverwriteOptionOutput) ToSsisMigrationOverwriteOptionPtrOutput() SsisMigrationOverwriteOptionPtrOutput {
	return o.ToSsisMigrationOverwriteOptionPtrOutputWithContext(context.Background())
}

func (o SsisMigrationOverwriteOptionOutput) ToSsisMigrationOverwriteOptionPtrOutputWithContext(ctx context.Context) SsisMigrationOverwriteOptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SsisMigrationOverwriteOption) *SsisMigrationOverwriteOption {
		return &v
	}).(SsisMigrationOverwriteOptionPtrOutput)
}

func (o SsisMigrationOverwriteOptionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SsisMigrationOverwriteOptionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SsisMigrationOverwriteOption) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SsisMigrationOverwriteOptionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SsisMigrationOverwriteOptionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SsisMigrationOverwriteOption) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SsisMigrationOverwriteOptionPtrOutput struct{ *pulumi.OutputState }

func (SsisMigrationOverwriteOptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SsisMigrationOverwriteOption)(nil)).Elem()
}

func (o SsisMigrationOverwriteOptionPtrOutput) ToSsisMigrationOverwriteOptionPtrOutput() SsisMigrationOverwriteOptionPtrOutput {
	return o
}

func (o SsisMigrationOverwriteOptionPtrOutput) ToSsisMigrationOverwriteOptionPtrOutputWithContext(ctx context.Context) SsisMigrationOverwriteOptionPtrOutput {
	return o
}

func (o SsisMigrationOverwriteOptionPtrOutput) Elem() SsisMigrationOverwriteOptionOutput {
	return o.ApplyT(func(v *SsisMigrationOverwriteOption) SsisMigrationOverwriteOption {
		if v != nil {
			return *v
		}
		var ret SsisMigrationOverwriteOption
		return ret
	}).(SsisMigrationOverwriteOptionOutput)
}

func (o SsisMigrationOverwriteOptionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SsisMigrationOverwriteOptionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SsisMigrationOverwriteOption) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SsisMigrationOverwriteOptionInput interface {
	pulumi.Input

	ToSsisMigrationOverwriteOptionOutput() SsisMigrationOverwriteOptionOutput
	ToSsisMigrationOverwriteOptionOutputWithContext(context.Context) SsisMigrationOverwriteOptionOutput
}

var ssisMigrationOverwriteOptionPtrType = reflect.TypeOf((**SsisMigrationOverwriteOption)(nil)).Elem()

type SsisMigrationOverwriteOptionPtrInput interface {
	pulumi.Input

	ToSsisMigrationOverwriteOptionPtrOutput() SsisMigrationOverwriteOptionPtrOutput
	ToSsisMigrationOverwriteOptionPtrOutputWithContext(context.Context) SsisMigrationOverwriteOptionPtrOutput
}

type ssisMigrationOverwriteOptionPtr string

func SsisMigrationOverwriteOptionPtr(v string) SsisMigrationOverwriteOptionPtrInput {
	return (*ssisMigrationOverwriteOptionPtr)(&v)
}

func (*ssisMigrationOverwriteOptionPtr) ElementType() reflect.Type {
	return ssisMigrationOverwriteOptionPtrType
}

func (in *ssisMigrationOverwriteOptionPtr) ToSsisMigrationOverwriteOptionPtrOutput() SsisMigrationOverwriteOptionPtrOutput {
	return pulumi.ToOutput(in).(SsisMigrationOverwriteOptionPtrOutput)
}

func (in *ssisMigrationOverwriteOptionPtr) ToSsisMigrationOverwriteOptionPtrOutputWithContext(ctx context.Context) SsisMigrationOverwriteOptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SsisMigrationOverwriteOptionPtrOutput)
}

type SsisStoreType string

const (
	SsisStoreTypeSsisCatalog = SsisStoreType("SsisCatalog")
)

func (SsisStoreType) ElementType() reflect.Type {
	return reflect.TypeOf((*SsisStoreType)(nil)).Elem()
}

func (e SsisStoreType) ToSsisStoreTypeOutput() SsisStoreTypeOutput {
	return pulumi.ToOutput(e).(SsisStoreTypeOutput)
}

func (e SsisStoreType) ToSsisStoreTypeOutputWithContext(ctx context.Context) SsisStoreTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SsisStoreTypeOutput)
}

func (e SsisStoreType) ToSsisStoreTypePtrOutput() SsisStoreTypePtrOutput {
	return e.ToSsisStoreTypePtrOutputWithContext(context.Background())
}

func (e SsisStoreType) ToSsisStoreTypePtrOutputWithContext(ctx context.Context) SsisStoreTypePtrOutput {
	return SsisStoreType(e).ToSsisStoreTypeOutputWithContext(ctx).ToSsisStoreTypePtrOutputWithContext(ctx)
}

func (e SsisStoreType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SsisStoreType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SsisStoreType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SsisStoreType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SsisStoreTypeOutput struct{ *pulumi.OutputState }

func (SsisStoreTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SsisStoreType)(nil)).Elem()
}

func (o SsisStoreTypeOutput) ToSsisStoreTypeOutput() SsisStoreTypeOutput {
	return o
}

func (o SsisStoreTypeOutput) ToSsisStoreTypeOutputWithContext(ctx context.Context) SsisStoreTypeOutput {
	return o
}

func (o SsisStoreTypeOutput) ToSsisStoreTypePtrOutput() SsisStoreTypePtrOutput {
	return o.ToSsisStoreTypePtrOutputWithContext(context.Background())
}

func (o SsisStoreTypeOutput) ToSsisStoreTypePtrOutputWithContext(ctx context.Context) SsisStoreTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SsisStoreType) *SsisStoreType {
		return &v
	}).(SsisStoreTypePtrOutput)
}

func (o SsisStoreTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SsisStoreTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SsisStoreType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SsisStoreTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SsisStoreTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SsisStoreType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SsisStoreTypePtrOutput struct{ *pulumi.OutputState }

func (SsisStoreTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SsisStoreType)(nil)).Elem()
}

func (o SsisStoreTypePtrOutput) ToSsisStoreTypePtrOutput() SsisStoreTypePtrOutput {
	return o
}

func (o SsisStoreTypePtrOutput) ToSsisStoreTypePtrOutputWithContext(ctx context.Context) SsisStoreTypePtrOutput {
	return o
}

func (o SsisStoreTypePtrOutput) Elem() SsisStoreTypeOutput {
	return o.ApplyT(func(v *SsisStoreType) SsisStoreType {
		if v != nil {
			return *v
		}
		var ret SsisStoreType
		return ret
	}).(SsisStoreTypeOutput)
}

func (o SsisStoreTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SsisStoreTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SsisStoreType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SsisStoreTypeInput interface {
	pulumi.Input

	ToSsisStoreTypeOutput() SsisStoreTypeOutput
	ToSsisStoreTypeOutputWithContext(context.Context) SsisStoreTypeOutput
}

var ssisStoreTypePtrType = reflect.TypeOf((**SsisStoreType)(nil)).Elem()

type SsisStoreTypePtrInput interface {
	pulumi.Input

	ToSsisStoreTypePtrOutput() SsisStoreTypePtrOutput
	ToSsisStoreTypePtrOutputWithContext(context.Context) SsisStoreTypePtrOutput
}

type ssisStoreTypePtr string

func SsisStoreTypePtr(v string) SsisStoreTypePtrInput {
	return (*ssisStoreTypePtr)(&v)
}

func (*ssisStoreTypePtr) ElementType() reflect.Type {
	return ssisStoreTypePtrType
}

func (in *ssisStoreTypePtr) ToSsisStoreTypePtrOutput() SsisStoreTypePtrOutput {
	return pulumi.ToOutput(in).(SsisStoreTypePtrOutput)
}

func (in *ssisStoreTypePtr) ToSsisStoreTypePtrOutputWithContext(ctx context.Context) SsisStoreTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SsisStoreTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AuthenticationTypeOutput{})
	pulumi.RegisterOutputType(AuthenticationTypePtrOutput{})
	pulumi.RegisterOutputType(BackupModeOutput{})
	pulumi.RegisterOutputType(BackupModePtrOutput{})
	pulumi.RegisterOutputType(MongoDbReplicationOutput{})
	pulumi.RegisterOutputType(MongoDbReplicationPtrOutput{})
	pulumi.RegisterOutputType(MongoDbShardKeyOrderOutput{})
	pulumi.RegisterOutputType(MongoDbShardKeyOrderPtrOutput{})
	pulumi.RegisterOutputType(MySqlTargetPlatformTypeOutput{})
	pulumi.RegisterOutputType(MySqlTargetPlatformTypePtrOutput{})
	pulumi.RegisterOutputType(ProjectSourcePlatformOutput{})
	pulumi.RegisterOutputType(ProjectSourcePlatformPtrOutput{})
	pulumi.RegisterOutputType(ProjectTargetPlatformOutput{})
	pulumi.RegisterOutputType(ProjectTargetPlatformPtrOutput{})
	pulumi.RegisterOutputType(ServerLevelPermissionsGroupOutput{})
	pulumi.RegisterOutputType(ServerLevelPermissionsGroupPtrOutput{})
	pulumi.RegisterOutputType(SqlSourcePlatformOutput{})
	pulumi.RegisterOutputType(SqlSourcePlatformPtrOutput{})
	pulumi.RegisterOutputType(SsisMigrationOverwriteOptionOutput{})
	pulumi.RegisterOutputType(SsisMigrationOverwriteOptionPtrOutput{})
	pulumi.RegisterOutputType(SsisStoreTypeOutput{})
	pulumi.RegisterOutputType(SsisStoreTypePtrOutput{})
}
