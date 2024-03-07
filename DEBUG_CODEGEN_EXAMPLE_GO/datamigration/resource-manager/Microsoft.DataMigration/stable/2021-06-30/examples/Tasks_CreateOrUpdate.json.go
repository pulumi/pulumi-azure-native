package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/datamigration/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datamigration.NewTask(ctx, "task", &datamigration.TaskArgs{
			GroupName:   pulumi.String("DmsSdkRg"),
			ProjectName: pulumi.String("DmsSdkProject"),
			Properties: datamigration.ConnectToTargetSqlDbTaskProperties{
				Input: datamigration.ConnectToTargetSqlDbTaskInput{
					TargetConnectionInfo: datamigration.SqlConnectionInfo{
						Authentication:         "SqlAuthentication",
						DataSource:             "ssma-test-server.database.windows.net",
						EncryptConnection:      true,
						Password:               "testpassword",
						TrustServerCertificate: true,
						Type:                   "SqlConnectionInfo",
						UserName:               "testuser",
					},
				},
				TaskType: "ConnectToTarget.SqlDb",
			},
			ServiceName: pulumi.String("DmsSdkService"),
			TaskName:    pulumi.String("DmsSdkTask"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
