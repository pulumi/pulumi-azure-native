


package servicelinker

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListLinkerConfigurations(ctx *pulumi.Context, args *ListLinkerConfigurationsArgs, opts ...pulumi.InvokeOption) (*ListLinkerConfigurationsResult, error) {
	var rv ListLinkerConfigurationsResult
	err := ctx.Invoke("azure-native:servicelinker:listLinkerConfigurations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListLinkerConfigurationsArgs struct {
	LinkerName  string `pulumi:"linkerName"`
	ResourceUri string `pulumi:"resourceUri"`
}


type ListLinkerConfigurationsResult struct {
	Configurations []SourceConfigurationResponse `pulumi:"configurations"`
}
