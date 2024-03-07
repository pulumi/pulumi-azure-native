package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/kusto/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := kusto.NewScript(ctx, "script", &kusto.ScriptArgs{
			ClusterName:       pulumi.String("kustoCluster"),
			ContinueOnErrors:  pulumi.Bool(true),
			DatabaseName:      pulumi.String("KustoDatabase8"),
			ForceUpdateTag:    pulumi.String("2bcf3c21-ffd1-4444-b9dd-e52e00ee53fe"),
			ResourceGroupName: pulumi.String("kustorptest"),
			ScriptName:        pulumi.String("kustoScript"),
			ScriptUrl:         pulumi.String("https://mysa.blob.core.windows.net/container/script.txt"),
			ScriptUrlSasToken: pulumi.String("?sv=2019-02-02&st=2019-04-29T22%3A18%3A26Z&se=2019-04-30T02%3A23%3A26Z&sr=b&sp=rw&sip=168.1.5.60-168.1.5.70&spr=https&sig=********************************"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
