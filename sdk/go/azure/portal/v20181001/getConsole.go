


package v20181001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConsole(ctx *pulumi.Context, args *LookupConsoleArgs, opts ...pulumi.InvokeOption) (*LookupConsoleResult, error) {
	var rv LookupConsoleResult
	err := ctx.Invoke("azure-native:portal/v20181001:getConsole", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConsoleArgs struct {
	ConsoleName string `pulumi:"consoleName"`
}


type LookupConsoleResult struct {
	Properties ConsolePropertiesResponse `pulumi:"properties"`
}
