


package datamigration

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

type ProjectSourcePlatform string

const (
	ProjectSourcePlatformSQL     = ProjectSourcePlatform("SQL")
	ProjectSourcePlatformUnknown = ProjectSourcePlatform("Unknown")
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
	ProjectTargetPlatformSQLDB   = ProjectTargetPlatform("SQLDB")
	ProjectTargetPlatformUnknown = ProjectTargetPlatform("Unknown")
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

func init() {
	pulumi.RegisterOutputType(AuthenticationTypeOutput{})
	pulumi.RegisterOutputType(AuthenticationTypePtrOutput{})
	pulumi.RegisterOutputType(BackupModeOutput{})
	pulumi.RegisterOutputType(BackupModePtrOutput{})
	pulumi.RegisterOutputType(ProjectSourcePlatformOutput{})
	pulumi.RegisterOutputType(ProjectSourcePlatformPtrOutput{})
	pulumi.RegisterOutputType(ProjectTargetPlatformOutput{})
	pulumi.RegisterOutputType(ProjectTargetPlatformPtrOutput{})
	pulumi.RegisterOutputType(ServerLevelPermissionsGroupOutput{})
	pulumi.RegisterOutputType(ServerLevelPermissionsGroupPtrOutput{})
	pulumi.RegisterOutputType(SqlSourcePlatformOutput{})
	pulumi.RegisterOutputType(SqlSourcePlatformPtrOutput{})
}
