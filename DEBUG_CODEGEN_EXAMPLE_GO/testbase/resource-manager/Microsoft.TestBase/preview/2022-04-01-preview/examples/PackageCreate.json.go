package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/testbase/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := testbase.NewPackage(ctx, "package", &testbase.PackageArgs{
			ApplicationName:   pulumi.String("contoso-package2"),
			BlobPath:          pulumi.String("storageAccountPath/package.zip"),
			FlightingRing:     pulumi.String("Insider Beta Channel"),
			Location:          pulumi.String("westus"),
			PackageName:       pulumi.String("contoso-package2"),
			ResourceGroupName: pulumi.String("contoso-rg1"),
			Tags:              nil,
			TargetOSList: []testbase.TargetOSInfoArgs{
				{
					OsUpdateType: pulumi.String("Security updates"),
					TargetOSs: pulumi.StringArray{
						pulumi.String("Windows 10 2004"),
						pulumi.String("Windows 10 1903"),
					},
				},
			},
			TestBaseAccountName: pulumi.String("contoso-testBaseAccount1"),
			Tests: []testbase.TestArgs{
				{
					Commands: testbase.CommandArray{
						{
							Action:            pulumi.String("Install"),
							AlwaysRun:         pulumi.Bool(true),
							ApplyUpdateBefore: pulumi.Bool(false),
							Content:           pulumi.String("app/scripts/install/job.ps1"),
							ContentType:       pulumi.String("Path"),
							MaxRunTime:        pulumi.Int(1800),
							Name:              pulumi.String("Install"),
							RestartAfter:      pulumi.Bool(true),
							RunAsInteractive:  pulumi.Bool(true),
							RunElevated:       pulumi.Bool(true),
						},
						{
							Action:            pulumi.String("Launch"),
							AlwaysRun:         pulumi.Bool(false),
							ApplyUpdateBefore: pulumi.Bool(true),
							Content:           pulumi.String("app/scripts/launch/job.ps1"),
							ContentType:       pulumi.String("Path"),
							MaxRunTime:        pulumi.Int(1800),
							Name:              pulumi.String("Launch"),
							RestartAfter:      pulumi.Bool(false),
							RunAsInteractive:  pulumi.Bool(true),
							RunElevated:       pulumi.Bool(true),
						},
						{
							Action:            pulumi.String("Close"),
							AlwaysRun:         pulumi.Bool(false),
							ApplyUpdateBefore: pulumi.Bool(false),
							Content:           pulumi.String("app/scripts/close/job.ps1"),
							ContentType:       pulumi.String("Path"),
							MaxRunTime:        pulumi.Int(1800),
							Name:              pulumi.String("Close"),
							RestartAfter:      pulumi.Bool(false),
							RunAsInteractive:  pulumi.Bool(true),
							RunElevated:       pulumi.Bool(true),
						},
						{
							Action:            pulumi.String("Uninstall"),
							AlwaysRun:         pulumi.Bool(true),
							ApplyUpdateBefore: pulumi.Bool(false),
							Content:           pulumi.String("app/scripts/uninstall/job.ps1"),
							ContentType:       pulumi.String("Path"),
							MaxRunTime:        pulumi.Int(1800),
							Name:              pulumi.String("Uninstall"),
							RestartAfter:      pulumi.Bool(false),
							RunAsInteractive:  pulumi.Bool(true),
							RunElevated:       pulumi.Bool(true),
						},
					},
					IsActive: pulumi.Bool(true),
					TestType: pulumi.String("OutOfBoxTest"),
				},
			},
			Version: pulumi.String("1.0.0"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
