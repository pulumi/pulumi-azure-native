package main

import (
	"log"

	"github.com/pulumi/pulumi-azure-native/sdk/go/azure/authorization"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	q "github.com/ryboe/q"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tokenResult, err := authorization.GetClientToken(ctx, &authorization.GetClientTokenArgs{}, nil)
		if err != nil {
			return err
		}

		log.Println("Got OAuth token: ", tokenResult.Token[:20])
		q.Q(tokenResult.Token)

		return nil
	})
}
