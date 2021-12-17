


package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AuthenticationMethod string

const (
	AuthenticationMethodNone      = AuthenticationMethod("None")
	AuthenticationMethodCassandra = AuthenticationMethod("Cassandra")
)

type BackupPolicyType string

const (
	BackupPolicyTypePeriodic   = BackupPolicyType("Periodic")
	BackupPolicyTypeContinuous = BackupPolicyType("Continuous")
)

type BackupStorageRedundancy string

const (
	BackupStorageRedundancyGeo   = BackupStorageRedundancy("Geo")
	BackupStorageRedundancyLocal = BackupStorageRedundancy("Local")
	BackupStorageRedundancyZone  = BackupStorageRedundancy("Zone")
)

type CompositePathSortOrder string

const (
	CompositePathSortOrderAscending  = CompositePathSortOrder("ascending")
	CompositePathSortOrderDescending = CompositePathSortOrder("descending")
)

type ConflictResolutionMode string

const (
	ConflictResolutionModeLastWriterWins = ConflictResolutionMode("LastWriterWins")
	ConflictResolutionModeCustom         = ConflictResolutionMode("Custom")
)

type ConnectorOffer string

const (
	ConnectorOfferSmall = ConnectorOffer("Small")
)

type CreateMode string

const (
	CreateModeDefault = CreateMode("Default")
	CreateModeRestore = CreateMode("Restore")
)

type DataType string

const (
	DataTypeString       = DataType("String")
	DataTypeNumber       = DataType("Number")
	DataTypePoint        = DataType("Point")
	DataTypePolygon      = DataType("Polygon")
	DataTypeLineString   = DataType("LineString")
	DataTypeMultiPolygon = DataType("MultiPolygon")
)

type DatabaseAccountKind string

const (
	DatabaseAccountKindGlobalDocumentDB = DatabaseAccountKind("GlobalDocumentDB")
	DatabaseAccountKindMongoDB          = DatabaseAccountKind("MongoDB")
	DatabaseAccountKindParse            = DatabaseAccountKind("Parse")
)

type DatabaseAccountOfferType string

const (
	DatabaseAccountOfferTypeStandard = DatabaseAccountOfferType("Standard")
)

type DefaultConsistencyLevel string

const (
	DefaultConsistencyLevelEventual         = DefaultConsistencyLevel("Eventual")
	DefaultConsistencyLevelSession          = DefaultConsistencyLevel("Session")
	DefaultConsistencyLevelBoundedStaleness = DefaultConsistencyLevel("BoundedStaleness")
	DefaultConsistencyLevelStrong           = DefaultConsistencyLevel("Strong")
	DefaultConsistencyLevelConsistentPrefix = DefaultConsistencyLevel("ConsistentPrefix")
)

type IndexKind string

const (
	IndexKindHash    = IndexKind("Hash")
	IndexKindRange   = IndexKind("Range")
	IndexKindSpatial = IndexKind("Spatial")
)

type IndexingMode string

const (
	IndexingModeConsistent = IndexingMode("consistent")
	IndexingModeLazy       = IndexingMode("lazy")
	IndexingModeNone       = IndexingMode("none")
)

type ManagedCassandraProvisioningState string

const (
	ManagedCassandraProvisioningStateCreating  = ManagedCassandraProvisioningState("Creating")
	ManagedCassandraProvisioningStateUpdating  = ManagedCassandraProvisioningState("Updating")
	ManagedCassandraProvisioningStateDeleting  = ManagedCassandraProvisioningState("Deleting")
	ManagedCassandraProvisioningStateSucceeded = ManagedCassandraProvisioningState("Succeeded")
	ManagedCassandraProvisioningStateFailed    = ManagedCassandraProvisioningState("Failed")
	ManagedCassandraProvisioningStateCanceled  = ManagedCassandraProvisioningState("Canceled")
)

type NetworkAclBypass string

const (
	NetworkAclBypassNone          = NetworkAclBypass("None")
	NetworkAclBypassAzureServices = NetworkAclBypass("AzureServices")
)

type PartitionKind string

const (
	PartitionKindHash      = PartitionKind("Hash")
	PartitionKindRange     = PartitionKind("Range")
	PartitionKindMultiHash = PartitionKind("MultiHash")
)

type PublicNetworkAccess string

const (
	PublicNetworkAccessEnabled  = PublicNetworkAccess("Enabled")
	PublicNetworkAccessDisabled = PublicNetworkAccess("Disabled")
)

type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned               = ResourceIdentityType("SystemAssigned")
	ResourceIdentityTypeUserAssigned                 = ResourceIdentityType("UserAssigned")
	ResourceIdentityType_SystemAssigned_UserAssigned = ResourceIdentityType("SystemAssigned,UserAssigned")
	ResourceIdentityTypeNone                         = ResourceIdentityType("None")
)

func (ResourceIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return pulumi.ToOutput(e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return e.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return ResourceIdentityType(e).ToResourceIdentityTypeOutputWithContext(ctx).ToResourceIdentityTypePtrOutputWithContext(ctx)
}

func (e ResourceIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ResourceIdentityTypeOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentityType) *ResourceIdentityType {
		return &v
	}).(ResourceIdentityTypePtrOutput)
}

func (o ResourceIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ResourceIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) Elem() ResourceIdentityTypeOutput {
	return o.ApplyT(func(v *ResourceIdentityType) ResourceIdentityType {
		if v != nil {
			return *v
		}
		var ret ResourceIdentityType
		return ret
	}).(ResourceIdentityTypeOutput)
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ResourceIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ResourceIdentityTypeInput interface {
	pulumi.Input

	ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput
	ToResourceIdentityTypeOutputWithContext(context.Context) ResourceIdentityTypeOutput
}

var resourceIdentityTypePtrType = reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()

type ResourceIdentityTypePtrInput interface {
	pulumi.Input

	ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput
	ToResourceIdentityTypePtrOutputWithContext(context.Context) ResourceIdentityTypePtrOutput
}

type resourceIdentityTypePtr string

func ResourceIdentityTypePtr(v string) ResourceIdentityTypePtrInput {
	return (*resourceIdentityTypePtr)(&v)
}

func (*resourceIdentityTypePtr) ElementType() reflect.Type {
	return resourceIdentityTypePtrType
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(ResourceIdentityTypePtrOutput)
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ResourceIdentityTypePtrOutput)
}

type RestoreMode string

const (
	RestoreModePointInTime = RestoreMode("PointInTime")
)

type RoleDefinitionType string

const (
	RoleDefinitionTypeBuiltInRole = RoleDefinitionType("BuiltInRole")
	RoleDefinitionTypeCustomRole  = RoleDefinitionType("CustomRole")
)

func (RoleDefinitionType) ElementType() reflect.Type {
	return reflect.TypeOf((*RoleDefinitionType)(nil)).Elem()
}

func (e RoleDefinitionType) ToRoleDefinitionTypeOutput() RoleDefinitionTypeOutput {
	return pulumi.ToOutput(e).(RoleDefinitionTypeOutput)
}

func (e RoleDefinitionType) ToRoleDefinitionTypeOutputWithContext(ctx context.Context) RoleDefinitionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RoleDefinitionTypeOutput)
}

func (e RoleDefinitionType) ToRoleDefinitionTypePtrOutput() RoleDefinitionTypePtrOutput {
	return e.ToRoleDefinitionTypePtrOutputWithContext(context.Background())
}

func (e RoleDefinitionType) ToRoleDefinitionTypePtrOutputWithContext(ctx context.Context) RoleDefinitionTypePtrOutput {
	return RoleDefinitionType(e).ToRoleDefinitionTypeOutputWithContext(ctx).ToRoleDefinitionTypePtrOutputWithContext(ctx)
}

func (e RoleDefinitionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RoleDefinitionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RoleDefinitionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RoleDefinitionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RoleDefinitionTypeOutput struct{ *pulumi.OutputState }

func (RoleDefinitionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoleDefinitionType)(nil)).Elem()
}

func (o RoleDefinitionTypeOutput) ToRoleDefinitionTypeOutput() RoleDefinitionTypeOutput {
	return o
}

func (o RoleDefinitionTypeOutput) ToRoleDefinitionTypeOutputWithContext(ctx context.Context) RoleDefinitionTypeOutput {
	return o
}

func (o RoleDefinitionTypeOutput) ToRoleDefinitionTypePtrOutput() RoleDefinitionTypePtrOutput {
	return o.ToRoleDefinitionTypePtrOutputWithContext(context.Background())
}

func (o RoleDefinitionTypeOutput) ToRoleDefinitionTypePtrOutputWithContext(ctx context.Context) RoleDefinitionTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RoleDefinitionType) *RoleDefinitionType {
		return &v
	}).(RoleDefinitionTypePtrOutput)
}

func (o RoleDefinitionTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RoleDefinitionTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RoleDefinitionType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RoleDefinitionTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RoleDefinitionTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RoleDefinitionType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RoleDefinitionTypePtrOutput struct{ *pulumi.OutputState }

func (RoleDefinitionTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleDefinitionType)(nil)).Elem()
}

func (o RoleDefinitionTypePtrOutput) ToRoleDefinitionTypePtrOutput() RoleDefinitionTypePtrOutput {
	return o
}

func (o RoleDefinitionTypePtrOutput) ToRoleDefinitionTypePtrOutputWithContext(ctx context.Context) RoleDefinitionTypePtrOutput {
	return o
}

func (o RoleDefinitionTypePtrOutput) Elem() RoleDefinitionTypeOutput {
	return o.ApplyT(func(v *RoleDefinitionType) RoleDefinitionType {
		if v != nil {
			return *v
		}
		var ret RoleDefinitionType
		return ret
	}).(RoleDefinitionTypeOutput)
}

func (o RoleDefinitionTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RoleDefinitionTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RoleDefinitionType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type RoleDefinitionTypeInput interface {
	pulumi.Input

	ToRoleDefinitionTypeOutput() RoleDefinitionTypeOutput
	ToRoleDefinitionTypeOutputWithContext(context.Context) RoleDefinitionTypeOutput
}

var roleDefinitionTypePtrType = reflect.TypeOf((**RoleDefinitionType)(nil)).Elem()

type RoleDefinitionTypePtrInput interface {
	pulumi.Input

	ToRoleDefinitionTypePtrOutput() RoleDefinitionTypePtrOutput
	ToRoleDefinitionTypePtrOutputWithContext(context.Context) RoleDefinitionTypePtrOutput
}

type roleDefinitionTypePtr string

func RoleDefinitionTypePtr(v string) RoleDefinitionTypePtrInput {
	return (*roleDefinitionTypePtr)(&v)
}

func (*roleDefinitionTypePtr) ElementType() reflect.Type {
	return roleDefinitionTypePtrType
}

func (in *roleDefinitionTypePtr) ToRoleDefinitionTypePtrOutput() RoleDefinitionTypePtrOutput {
	return pulumi.ToOutput(in).(RoleDefinitionTypePtrOutput)
}

func (in *roleDefinitionTypePtr) ToRoleDefinitionTypePtrOutputWithContext(ctx context.Context) RoleDefinitionTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RoleDefinitionTypePtrOutput)
}

type ServerVersion string

const (
	ServerVersion_3_2 = ServerVersion("3.2")
	ServerVersion_3_6 = ServerVersion("3.6")
	ServerVersion_4_0 = ServerVersion("4.0")
)

type SpatialType string

const (
	SpatialTypePoint        = SpatialType("Point")
	SpatialTypeLineString   = SpatialType("LineString")
	SpatialTypePolygon      = SpatialType("Polygon")
	SpatialTypeMultiPolygon = SpatialType("MultiPolygon")
)

type TriggerOperation string

const (
	TriggerOperationAll     = TriggerOperation("All")
	TriggerOperationCreate  = TriggerOperation("Create")
	TriggerOperationUpdate  = TriggerOperation("Update")
	TriggerOperationDelete  = TriggerOperation("Delete")
	TriggerOperationReplace = TriggerOperation("Replace")
)

type TriggerType string

const (
	TriggerTypePre  = TriggerType("Pre")
	TriggerTypePost = TriggerType("Post")
)

func init() {
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
	pulumi.RegisterOutputType(RoleDefinitionTypeOutput{})
	pulumi.RegisterOutputType(RoleDefinitionTypePtrOutput{})
}
