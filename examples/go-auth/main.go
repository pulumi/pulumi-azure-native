package main

import (
	"fmt"
	"regexp"

	"github.com/pulumi/pulumi-azure-native/sdk/go/azure/authorization"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tokenResult, err := authorization.GetClientToken(ctx, &authorization.GetClientTokenArgs{}, nil)
		if err != nil {
			return err
		}

		// Based on the ABNF from RFC 6750:
		//    1*( ALPHA / DIGIT / "-" / "." / "_" / "~" / "+" / "/" ) *"="
		tokenLooksValid := regexp.MustCompile(`^[._~+/a-zA-Z0-9-]+=*$`).MatchString(tokenResult.Token)
		if !tokenLooksValid {
			token := tokenResult.Token
			if len(token) > 40 {
				token = token[:40] + "..."
			}
			return fmt.Errorf("Token '%s' does not seem valid", token)
		}

		return nil
	})
}
