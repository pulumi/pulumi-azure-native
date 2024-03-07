package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storage/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storage.NewTable(ctx, "table", &storage.TableArgs{
			AccountName:       pulumi.String("sto328"),
			ResourceGroupName: pulumi.String("res3376"),
			SignedIdentifiers: storage.TableSignedIdentifierArray{
				&storage.TableSignedIdentifierArgs{
					AccessPolicy: &storage.TableAccessPolicyArgs{
						ExpiryTime: pulumi.String("2022-03-20T08:49:37.0000000Z"),
						Permission: pulumi.String("raud"),
						StartTime:  pulumi.String("2022-03-17T08:49:37.0000000Z"),
					},
					Id: pulumi.String("MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI"),
				},
				&storage.TableSignedIdentifierArgs{
					AccessPolicy: &storage.TableAccessPolicyArgs{
						ExpiryTime: pulumi.String("2022-03-20T08:49:37.0000000Z"),
						Permission: pulumi.String("rad"),
						StartTime:  pulumi.String("2022-03-17T08:49:37.0000000Z"),
					},
					Id: pulumi.String("PTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODklMTI"),
				},
			},
			TableName: pulumi.String("table6185"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
