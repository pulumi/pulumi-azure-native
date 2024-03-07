package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storage/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := storage.NewManagementPolicy(ctx, "managementPolicy", &storage.ManagementPolicyArgs{
AccountName: pulumi.String("sto9699"),
ManagementPolicyName: pulumi.String("default"),
Policy: storage.ManagementPolicySchemaResponse{
Rules: storage.ManagementPolicyRuleArray{
interface{}{
Definition: interface{}{
Actions: interface{}{
BaseBlob: interface{}{
Delete: &storage.DateAfterModificationArgs{
DaysAfterModificationGreaterThan: pulumi.Float64(1000),
},
TierToArchive: &storage.DateAfterModificationArgs{
DaysAfterLastTierChangeGreaterThan: pulumi.Float64(120),
DaysAfterModificationGreaterThan: pulumi.Float64(90),
},
TierToCool: &storage.DateAfterModificationArgs{
DaysAfterModificationGreaterThan: pulumi.Float64(30),
},
},
Snapshot: interface{}{
TierToArchive: &storage.DateAfterCreationArgs{
DaysAfterCreationGreaterThan: pulumi.Float64(30),
DaysAfterLastTierChangeGreaterThan: pulumi.Float64(90),
},
},
Version: interface{}{
TierToArchive: &storage.DateAfterCreationArgs{
DaysAfterCreationGreaterThan: pulumi.Float64(30),
DaysAfterLastTierChangeGreaterThan: pulumi.Float64(90),
},
},
},
Filters: &storage.ManagementPolicyFilterArgs{
BlobTypes: pulumi.StringArray{
pulumi.String("blockBlob"),
},
PrefixMatch: pulumi.StringArray{
pulumi.String("olcmtestcontainer"),
},
},
},
Enabled: pulumi.Bool(true),
Name: pulumi.String("olcmtest"),
Type: pulumi.String("Lifecycle"),
},
},
},
ResourceGroupName: pulumi.String("res7687"),
})
if err != nil {
return err
}
return nil
})
}
