package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/desktopvirtualization/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := desktopvirtualization.NewMSIXPackage(ctx, "msixPackage", &desktopvirtualization.MSIXPackageArgs{
			DisplayName:           pulumi.String("displayname"),
			HostPoolName:          pulumi.String("hostpool1"),
			ImagePath:             pulumi.String("imagepath"),
			IsActive:              pulumi.Bool(false),
			IsRegularRegistration: pulumi.Bool(false),
			LastUpdated:           pulumi.String("2008-09-22T14:01:54.9571247Z"),
			MsixPackageFullName:   pulumi.String("msixpackagefullname"),
			PackageApplications: desktopvirtualization.MsixPackageApplicationsArray{
				&desktopvirtualization.MsixPackageApplicationsArgs{
					AppId:          pulumi.String("ApplicationId"),
					AppUserModelID: pulumi.String("AppUserModelId"),
					Description:    pulumi.String("application-desc"),
					FriendlyName:   pulumi.String("friendlyname"),
					IconImageName:  pulumi.String("Apptile"),
					RawIcon:        pulumi.String("VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo"),
					RawPng:         pulumi.String("VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo"),
				},
			},
			PackageDependencies: desktopvirtualization.MsixPackageDependenciesArray{
				&desktopvirtualization.MsixPackageDependenciesArgs{
					DependencyName: pulumi.String("MsixTest_Dependency_Name"),
					MinVersion:     pulumi.String("version"),
					Publisher:      pulumi.String("PublishedName"),
				},
			},
			PackageFamilyName:   pulumi.String("MsixPackage_FamilyName"),
			PackageName:         pulumi.String("MsixPackage_name"),
			PackageRelativePath: pulumi.String("packagerelativepath"),
			ResourceGroupName:   pulumi.String("resourceGroup1"),
			Version:             pulumi.String("version"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
