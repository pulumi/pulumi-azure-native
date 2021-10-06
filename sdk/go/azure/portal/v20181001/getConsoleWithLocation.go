


package v20181001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConsoleWithLocation(ctx *pulumi.Context, args *LookupConsoleWithLocationArgs, opts ...pulumi.InvokeOption) (*LookupConsoleWithLocationResult, error) {
	var rv LookupConsoleWithLocationResult
	err := ctx.Invoke("azure-native:portal/v20181001:getConsoleWithLocation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConsoleWithLocationArgs struct {
	ConsoleName string `pulumi:"consoleName"`
	Location    string `pulumi:"location"`
}


type LookupConsoleWithLocationResult struct {
	Properties ConsolePropertiesResponse `pulumi:"properties"`
}
