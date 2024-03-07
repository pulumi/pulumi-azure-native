package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/securityinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := securityinsights.NewWatchlist(ctx, "watchlist", &securityinsights.WatchlistArgs{
			ContentType:         pulumi.String("text/csv"),
			Description:         pulumi.String("Watchlist from CSV content"),
			DisplayName:         pulumi.String("High Value Assets Watchlist"),
			ItemsSearchKey:      pulumi.String("header1"),
			NumberOfLinesToSkip: pulumi.Int(1),
			Provider:            pulumi.String("Microsoft"),
			RawContent:          pulumi.String("This line will be skipped\nheader1,header2\nvalue1,value2"),
			ResourceGroupName:   pulumi.String("myRg"),
			Source:              pulumi.String("Local file"),
			WatchlistAlias:      pulumi.String("highValueAsset"),
			WorkspaceName:       pulumi.String("myWorkspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
