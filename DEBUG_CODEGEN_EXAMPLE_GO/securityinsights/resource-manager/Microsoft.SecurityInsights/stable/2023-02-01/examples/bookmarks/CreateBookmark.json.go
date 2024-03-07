package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/securityinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := securityinsights.NewBookmark(ctx, "bookmark", &securityinsights.BookmarkArgs{
			BookmarkId: pulumi.String("73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
			Created:    pulumi.String("2019-01-01T13:15:30Z"),
			CreatedBy: &securityinsights.UserInfoArgs{
				ObjectId: pulumi.String("2046feea-040d-4a46-9e2b-91c2941bfa70"),
			},
			DisplayName: pulumi.String("My bookmark"),
			Labels: pulumi.StringArray{
				pulumi.String("Tag1"),
				pulumi.String("Tag2"),
			},
			Notes:             pulumi.String("Found a suspicious activity"),
			Query:             pulumi.String("SecurityEvent | where TimeGenerated > ago(1d) and TimeGenerated < ago(2d)"),
			QueryResult:       pulumi.String("Security Event query result"),
			ResourceGroupName: pulumi.String("myRg"),
			Updated:           pulumi.String("2019-01-01T13:15:30Z"),
			UpdatedBy: &securityinsights.UserInfoArgs{
				ObjectId: pulumi.String("2046feea-040d-4a46-9e2b-91c2941bfa70"),
			},
			WorkspaceName: pulumi.String("myWorkspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
