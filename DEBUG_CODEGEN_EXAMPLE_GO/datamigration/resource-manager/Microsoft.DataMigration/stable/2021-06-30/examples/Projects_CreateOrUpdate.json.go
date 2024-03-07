package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/datamigration/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datamigration.NewProject(ctx, "project", &datamigration.ProjectArgs{
			GroupName:      pulumi.String("DmsSdkRg"),
			Location:       pulumi.String("southcentralus"),
			ProjectName:    pulumi.String("DmsSdkProject"),
			ServiceName:    pulumi.String("DmsSdkService"),
			SourcePlatform: pulumi.String("SQL"),
			TargetPlatform: pulumi.String("SQLDB"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
