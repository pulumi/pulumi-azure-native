package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewJobCredential(ctx, "jobCredential", &sql.JobCredentialArgs{
			CredentialName:    pulumi.String("cred1"),
			JobAgentName:      pulumi.String("agent1"),
			Password:          pulumi.String("<password>"),
			ResourceGroupName: pulumi.String("group1"),
			ServerName:        pulumi.String("server1"),
			Username:          pulumi.String("myuser"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
