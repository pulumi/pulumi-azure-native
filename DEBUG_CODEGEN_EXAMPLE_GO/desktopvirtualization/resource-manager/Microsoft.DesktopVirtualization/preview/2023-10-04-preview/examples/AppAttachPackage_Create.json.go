package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/desktopvirtualization/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := desktopvirtualization.NewAppAttachPackage(ctx, "appAttachPackage", &desktopvirtualization.AppAttachPackageArgs{
			AppAttachPackageName: pulumi.String("msixpackagefullname"),
			Location:             pulumi.String("southcentralus"),
			Properties: &desktopvirtualization.AppAttachPackagePropertiesArgs{
				FailHealthCheckOnStagingFailure: pulumi.String("NeedsAssistance"),
				HostPoolReferences:              pulumi.StringArray{},
				Image: &desktopvirtualization.AppAttachPackageInfoPropertiesArgs{
					CertificateExpiry:     pulumi.String("2023-01-02T17:18:19.1234567Z"),
					CertificateName:       pulumi.String("certName"),
					DisplayName:           pulumi.String("displayname"),
					ImagePath:             pulumi.String("imagepath"),
					IsActive:              pulumi.Bool(false),
					IsRegularRegistration: pulumi.Bool(false),
					LastUpdated:           pulumi.String("2008-09-22T14:01:54.9571247Z"),
					PackageAlias:          pulumi.String("msixpackagealias"),
					PackageApplications: desktopvirtualization.MsixPackageApplicationsArray{
						&desktopvirtualization.MsixPackageApplicationsArgs{
							AppId:          pulumi.String("AppId"),
							AppUserModelID: pulumi.String("AppUserModelId"),
							Description:    pulumi.String("PackageApplicationDescription"),
							FriendlyName:   pulumi.String("FriendlyName"),
							IconImageName:  pulumi.String("Iconimagename"),
							RawIcon:        pulumi.String("VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo"),
							RawPng:         pulumi.String("VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo"),
						},
					},
					PackageDependencies: desktopvirtualization.MsixPackageDependenciesArray{
						&desktopvirtualization.MsixPackageDependenciesArgs{
							DependencyName: pulumi.String("MsixPackage_Dependency_Name"),
							MinVersion:     pulumi.String("packageDep_version"),
							Publisher:      pulumi.String("MsixPackage_Dependency_Publisher"),
						},
					},
					PackageFamilyName:   pulumi.String("MsixPackage_FamilyName"),
					PackageFullName:     pulumi.String("MsixPackage_FullName"),
					PackageName:         pulumi.String("MsixPackageName"),
					PackageRelativePath: pulumi.String("packagerelativepath"),
					Version:             pulumi.String("packageversion"),
				},
				KeyVaultURL: pulumi.String(""),
			},
			ResourceGroupName: pulumi.String("resourceGroup1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
