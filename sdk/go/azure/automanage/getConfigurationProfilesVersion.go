


package automanage

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConfigurationProfilesVersion(ctx *pulumi.Context, args *LookupConfigurationProfilesVersionArgs, opts ...pulumi.InvokeOption) (*LookupConfigurationProfilesVersionResult, error) {
	var rv LookupConfigurationProfilesVersionResult
	err := ctx.Invoke("azure-native:automanage:getConfigurationProfilesVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConfigurationProfilesVersionArgs struct {
	ConfigurationProfileName string `pulumi:"configurationProfileName"`
	ResourceGroupName        string `pulumi:"resourceGroupName"`
	VersionName              string `pulumi:"versionName"`
}


type LookupConfigurationProfilesVersionResult struct {
	Id         string                                 `pulumi:"id"`
	Location   string                                 `pulumi:"location"`
	Name       string                                 `pulumi:"name"`
	Properties ConfigurationProfilePropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse                     `pulumi:"systemData"`
	Tags       map[string]string                      `pulumi:"tags"`
	Type       string                                 `pulumi:"type"`
}
