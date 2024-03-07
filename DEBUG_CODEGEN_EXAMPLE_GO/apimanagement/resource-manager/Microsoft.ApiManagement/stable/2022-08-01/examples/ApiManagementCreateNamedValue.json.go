package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewNamedValue(ctx, "namedValue", &apimanagement.NamedValueArgs{
			DisplayName:       pulumi.String("prop3name"),
			NamedValueId:      pulumi.String("testprop2"),
			ResourceGroupName: pulumi.String("rg1"),
			Secret:            pulumi.Bool(false),
			ServiceName:       pulumi.String("apimService1"),
			Tags: pulumi.StringArray{
				pulumi.String("foo"),
				pulumi.String("bar"),
			},
			Value: pulumi.String("propValue"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
