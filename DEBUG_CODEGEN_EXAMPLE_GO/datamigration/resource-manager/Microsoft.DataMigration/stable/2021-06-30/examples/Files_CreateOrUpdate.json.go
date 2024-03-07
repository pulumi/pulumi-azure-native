package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/datamigration/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datamigration.NewFile(ctx, "file", &datamigration.FileArgs{
			FileName:    pulumi.String("x114d023d8"),
			GroupName:   pulumi.String("DmsSdkRg"),
			ProjectName: pulumi.String("DmsSdkProject"),
			Properties: &datamigration.ProjectFilePropertiesArgs{
				FilePath: pulumi.String("DmsSdkFilePath/DmsSdkFile.sql"),
			},
			ServiceName: pulumi.String("DmsSdkService"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
