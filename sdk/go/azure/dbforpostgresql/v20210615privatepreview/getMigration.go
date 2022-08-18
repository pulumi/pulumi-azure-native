


package v20210615privatepreview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMigration(ctx *pulumi.Context, args *LookupMigrationArgs, opts ...pulumi.InvokeOption) (*LookupMigrationResult, error) {
	var rv LookupMigrationResult
	err := ctx.Invoke("azure-native:dbforpostgresql/v20210615privatepreview:getMigration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMigrationArgs struct {
	MigrationName                   string `pulumi:"migrationName"`
	TargetDBServerName              string `pulumi:"targetDBServerName"`
	TargetDBServerResourceGroupName string `pulumi:"targetDBServerResourceGroupName"`
	TargetDBServerSubscriptionId    string `pulumi:"targetDBServerSubscriptionId"`
}


type LookupMigrationResult struct {
	CurrentStatus                             MigrationStatusResponse            `pulumi:"currentStatus"`
	DBsToMigrate                              []string                           `pulumi:"dBsToMigrate"`
	Id                                        string                             `pulumi:"id"`
	Location                                  string                             `pulumi:"location"`
	MigrationDetailsLevel                     string                             `pulumi:"migrationDetailsLevel"`
	MigrationId                               string                             `pulumi:"migrationId"`
	MigrationName                             string                             `pulumi:"migrationName"`
	MigrationResourceGroup                    *MigrationResourceGroupResponse    `pulumi:"migrationResourceGroup"`
	MigrationWindowStartTimeInUtc             *string                            `pulumi:"migrationWindowStartTimeInUtc"`
	Name                                      string                             `pulumi:"name"`
	OverwriteDBsInTarget                      *bool                              `pulumi:"overwriteDBsInTarget"`
	SecretParameters                          *MigrationSecretParametersResponse `pulumi:"secretParameters"`
	SetupLogicalReplicationOnSourceDBIfNeeded *bool                              `pulumi:"setupLogicalReplicationOnSourceDBIfNeeded"`
	SourceDBServerMetadata                    DBServerMetadataResponse           `pulumi:"sourceDBServerMetadata"`
	SourceDBServerResourceId                  *string                            `pulumi:"sourceDBServerResourceId"`
	StartDataMigration                        *bool                              `pulumi:"startDataMigration"`
	SystemData                                SystemDataResponse                 `pulumi:"systemData"`
	Tags                                      map[string]string                  `pulumi:"tags"`
	TargetDBServerMetadata                    DBServerMetadataResponse           `pulumi:"targetDBServerMetadata"`
	TargetDBServerResourceId                  string                             `pulumi:"targetDBServerResourceId"`
	TriggerCutover                            *bool                              `pulumi:"triggerCutover"`
	Type                                      string                             `pulumi:"type"`
	UserAssignedIdentityResourceId            *string                            `pulumi:"userAssignedIdentityResourceId"`
}

func LookupMigrationOutput(ctx *pulumi.Context, args LookupMigrationOutputArgs, opts ...pulumi.InvokeOption) LookupMigrationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMigrationResult, error) {
			args := v.(LookupMigrationArgs)
			r, err := LookupMigration(ctx, &args, opts...)
			var s LookupMigrationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMigrationResultOutput)
}

type LookupMigrationOutputArgs struct {
	MigrationName                   pulumi.StringInput `pulumi:"migrationName"`
	TargetDBServerName              pulumi.StringInput `pulumi:"targetDBServerName"`
	TargetDBServerResourceGroupName pulumi.StringInput `pulumi:"targetDBServerResourceGroupName"`
	TargetDBServerSubscriptionId    pulumi.StringInput `pulumi:"targetDBServerSubscriptionId"`
}

func (LookupMigrationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMigrationArgs)(nil)).Elem()
}


type LookupMigrationResultOutput struct{ *pulumi.OutputState }

func (LookupMigrationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMigrationResult)(nil)).Elem()
}

func (o LookupMigrationResultOutput) ToLookupMigrationResultOutput() LookupMigrationResultOutput {
	return o
}

func (o LookupMigrationResultOutput) ToLookupMigrationResultOutputWithContext(ctx context.Context) LookupMigrationResultOutput {
	return o
}

func (o LookupMigrationResultOutput) CurrentStatus() MigrationStatusResponseOutput {
	return o.ApplyT(func(v LookupMigrationResult) MigrationStatusResponse { return v.CurrentStatus }).(MigrationStatusResponseOutput)
}

func (o LookupMigrationResultOutput) DBsToMigrate() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupMigrationResult) []string { return v.DBsToMigrate }).(pulumi.StringArrayOutput)
}

func (o LookupMigrationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMigrationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMigrationResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMigrationResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupMigrationResultOutput) MigrationDetailsLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMigrationResult) string { return v.MigrationDetailsLevel }).(pulumi.StringOutput)
}

func (o LookupMigrationResultOutput) MigrationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMigrationResult) string { return v.MigrationId }).(pulumi.StringOutput)
}

func (o LookupMigrationResultOutput) MigrationName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMigrationResult) string { return v.MigrationName }).(pulumi.StringOutput)
}

func (o LookupMigrationResultOutput) MigrationResourceGroup() MigrationResourceGroupResponsePtrOutput {
	return o.ApplyT(func(v LookupMigrationResult) *MigrationResourceGroupResponse { return v.MigrationResourceGroup }).(MigrationResourceGroupResponsePtrOutput)
}

func (o LookupMigrationResultOutput) MigrationWindowStartTimeInUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMigrationResult) *string { return v.MigrationWindowStartTimeInUtc }).(pulumi.StringPtrOutput)
}

func (o LookupMigrationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMigrationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMigrationResultOutput) OverwriteDBsInTarget() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupMigrationResult) *bool { return v.OverwriteDBsInTarget }).(pulumi.BoolPtrOutput)
}

func (o LookupMigrationResultOutput) SecretParameters() MigrationSecretParametersResponsePtrOutput {
	return o.ApplyT(func(v LookupMigrationResult) *MigrationSecretParametersResponse { return v.SecretParameters }).(MigrationSecretParametersResponsePtrOutput)
}

func (o LookupMigrationResultOutput) SetupLogicalReplicationOnSourceDBIfNeeded() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupMigrationResult) *bool { return v.SetupLogicalReplicationOnSourceDBIfNeeded }).(pulumi.BoolPtrOutput)
}

func (o LookupMigrationResultOutput) SourceDBServerMetadata() DBServerMetadataResponseOutput {
	return o.ApplyT(func(v LookupMigrationResult) DBServerMetadataResponse { return v.SourceDBServerMetadata }).(DBServerMetadataResponseOutput)
}

func (o LookupMigrationResultOutput) SourceDBServerResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMigrationResult) *string { return v.SourceDBServerResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupMigrationResultOutput) StartDataMigration() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupMigrationResult) *bool { return v.StartDataMigration }).(pulumi.BoolPtrOutput)
}

func (o LookupMigrationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMigrationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupMigrationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMigrationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupMigrationResultOutput) TargetDBServerMetadata() DBServerMetadataResponseOutput {
	return o.ApplyT(func(v LookupMigrationResult) DBServerMetadataResponse { return v.TargetDBServerMetadata }).(DBServerMetadataResponseOutput)
}

func (o LookupMigrationResultOutput) TargetDBServerResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMigrationResult) string { return v.TargetDBServerResourceId }).(pulumi.StringOutput)
}

func (o LookupMigrationResultOutput) TriggerCutover() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupMigrationResult) *bool { return v.TriggerCutover }).(pulumi.BoolPtrOutput)
}

func (o LookupMigrationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMigrationResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupMigrationResultOutput) UserAssignedIdentityResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMigrationResult) *string { return v.UserAssignedIdentityResourceId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMigrationResultOutput{})
}
