package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storage/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := storage.NewBlobInventoryPolicy(ctx, "blobInventoryPolicy", &storage.BlobInventoryPolicyArgs{
AccountName: pulumi.String("sto9699"),
BlobInventoryPolicyName: pulumi.String("default"),
Policy: storage.BlobInventoryPolicySchemaResponse{
Enabled: pulumi.Bool(true),
Rules: storage.BlobInventoryPolicyRuleArray{
interface{}{
Definition: interface{}{
Filters: &storage.BlobInventoryPolicyFilterArgs{
BlobTypes: pulumi.StringArray{
pulumi.String("blockBlob"),
pulumi.String("appendBlob"),
pulumi.String("pageBlob"),
},
ExcludePrefix: pulumi.StringArray{
pulumi.String("excludeprefix1"),
pulumi.String("excludeprefix2"),
},
IncludeBlobVersions: pulumi.Bool(true),
IncludeDeleted: pulumi.Bool(true),
IncludeSnapshots: pulumi.Bool(true),
PrefixMatch: pulumi.StringArray{
pulumi.String("inventoryprefix1"),
pulumi.String("inventoryprefix2"),
},
},
Format: pulumi.String("Csv"),
ObjectType: pulumi.String("Blob"),
Schedule: pulumi.String("Daily"),
SchemaFields: pulumi.StringArray{
pulumi.String("Name"),
pulumi.String("Creation-Time"),
pulumi.String("Last-Modified"),
pulumi.String("Content-Length"),
pulumi.String("Content-MD5"),
pulumi.String("BlobType"),
pulumi.String("AccessTier"),
pulumi.String("AccessTierChangeTime"),
pulumi.String("Snapshot"),
pulumi.String("VersionId"),
pulumi.String("IsCurrentVersion"),
pulumi.String("ContentType"),
pulumi.String("ContentEncoding"),
pulumi.String("ContentLanguage"),
pulumi.String("ContentCRC64"),
pulumi.String("CacheControl"),
pulumi.String("Metadata"),
pulumi.String("DeletionId"),
pulumi.String("Deleted"),
pulumi.String("DeletedTime"),
pulumi.String("RemainingRetentionDays"),
},
},
Destination: pulumi.String("container1"),
Enabled: pulumi.Bool(true),
Name: pulumi.String("inventoryPolicyRule1"),
},
interface{}{
Definition: &storage.BlobInventoryPolicyDefinitionArgs{
Format: pulumi.String("Parquet"),
ObjectType: pulumi.String("Container"),
Schedule: pulumi.String("Weekly"),
SchemaFields: pulumi.StringArray{
pulumi.String("Name"),
pulumi.String("Last-Modified"),
pulumi.String("Metadata"),
pulumi.String("LeaseStatus"),
pulumi.String("LeaseState"),
pulumi.String("LeaseDuration"),
pulumi.String("PublicAccess"),
pulumi.String("HasImmutabilityPolicy"),
pulumi.String("HasLegalHold"),
pulumi.String("Etag"),
pulumi.String("DefaultEncryptionScope"),
pulumi.String("DenyEncryptionScopeOverride"),
pulumi.String("ImmutableStorageWithVersioningEnabled"),
pulumi.String("Deleted"),
pulumi.String("Version"),
pulumi.String("DeletedTime"),
pulumi.String("RemainingRetentionDays"),
},
},
Destination: pulumi.String("container2"),
Enabled: pulumi.Bool(true),
Name: pulumi.String("inventoryPolicyRule2"),
},
},
Type: pulumi.String("Inventory"),
},
ResourceGroupName: pulumi.String("res7687"),
})
if err != nil {
return err
}
return nil
})
}
