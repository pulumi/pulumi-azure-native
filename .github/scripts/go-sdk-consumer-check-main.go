// Used by `make verify_go_sdk_consumable` to check that the generated Go SDK can be tidied and built
// by a downstream consumer. See the target's comment in the Makefile and
// https://github.com/pulumi/pulumi-azure-native/issues/4763 for why this exists.
package main

import (
	resources "github.com/pulumi/pulumi-azure-native-sdk/resources/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := resources.NewResourceGroup(ctx, "rg", nil)
		return err
	})
}
