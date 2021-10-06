


package security

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListIngestionSettingConnectionStrings(ctx *pulumi.Context, args *ListIngestionSettingConnectionStringsArgs, opts ...pulumi.InvokeOption) (*ListIngestionSettingConnectionStringsResult, error) {
	var rv ListIngestionSettingConnectionStringsResult
	err := ctx.Invoke("azure-native:security:listIngestionSettingConnectionStrings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListIngestionSettingConnectionStringsArgs struct {
	IngestionSettingName string `pulumi:"ingestionSettingName"`
}


type ListIngestionSettingConnectionStringsResult struct {
	Value []IngestionConnectionStringResponse `pulumi:"value"`
}
