package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := compute.NewGalleryApplicationVersion(ctx, "galleryApplicationVersion", &compute.GalleryApplicationVersionArgs{
GalleryApplicationName: pulumi.String("myGalleryApplicationName"),
GalleryApplicationVersionName: pulumi.String("1.0.0"),
GalleryName: pulumi.String("myGalleryName"),
Location: pulumi.String("West US"),
PublishingProfile: compute.GalleryApplicationVersionPublishingProfileResponse{
CustomActions: compute.GalleryApplicationCustomActionArray{
interface{}{
Description: pulumi.String("This is the custom action description."),
Name: pulumi.String("myCustomAction"),
Parameters: compute.GalleryApplicationCustomActionParameterArray{
&compute.GalleryApplicationCustomActionParameterArgs{
DefaultValue: pulumi.String("default value of parameter."),
Description: pulumi.String("This is the description of the parameter"),
Name: pulumi.String("myCustomActionParameter"),
Required: pulumi.Bool(false),
Type: compute.GalleryApplicationCustomActionParameterTypeString,
},
},
Script: pulumi.String("myCustomActionScript"),
},
},
EndOfLifeDate: pulumi.String("2019-07-01T07:00:00Z"),
ManageActions: &compute.UserArtifactManageArgs{
Install: pulumi.String("powershell -command \"Expand-Archive -Path package.zip -DestinationPath C:\\package\""),
Remove: pulumi.String("del C:\\package "),
},
ReplicaCount: pulumi.Int(1),
Source: &compute.UserArtifactSourceArgs{
MediaLink: pulumi.String("https://mystorageaccount.blob.core.windows.net/mycontainer/package.zip?{sasKey}"),
},
StorageAccountType: pulumi.String("Standard_LRS"),
TargetRegions: compute.TargetRegionArray{
&compute.TargetRegionArgs{
ExcludeFromLatest: pulumi.Bool(false),
Name: pulumi.String("West US"),
RegionalReplicaCount: pulumi.Int(1),
StorageAccountType: pulumi.String("Standard_LRS"),
},
},
},
ResourceGroupName: pulumi.String("myResourceGroup"),
SafetyProfile: &compute.GalleryApplicationVersionSafetyProfileArgs{
AllowDeletionOfReplicatedLocations: pulumi.Bool(false),
},
})
if err != nil {
return err
}
return nil
})
}
