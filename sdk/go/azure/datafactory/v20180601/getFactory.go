


package v20180601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFactory(ctx *pulumi.Context, args *LookupFactoryArgs, opts ...pulumi.InvokeOption) (*LookupFactoryResult, error) {
	var rv LookupFactoryResult
	err := ctx.Invoke("azure-native:datafactory/v20180601:getFactory", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFactoryArgs struct {
	FactoryName       string `pulumi:"factoryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupFactoryResult struct {
	CreateTime          string                                          `pulumi:"createTime"`
	ETag                string                                          `pulumi:"eTag"`
	Encryption          *EncryptionConfigurationResponse                `pulumi:"encryption"`
	GlobalParameters    map[string]GlobalParameterSpecificationResponse `pulumi:"globalParameters"`
	Id                  string                                          `pulumi:"id"`
	Identity            *FactoryIdentityResponse                        `pulumi:"identity"`
	Location            *string                                         `pulumi:"location"`
	Name                string                                          `pulumi:"name"`
	ProvisioningState   string                                          `pulumi:"provisioningState"`
	PublicNetworkAccess *string                                         `pulumi:"publicNetworkAccess"`
	RepoConfiguration   interface{}                                     `pulumi:"repoConfiguration"`
	Tags                map[string]string                               `pulumi:"tags"`
	Type                string                                          `pulumi:"type"`
	Version             string                                          `pulumi:"version"`
}
