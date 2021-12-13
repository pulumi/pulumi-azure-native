


package containerregistry

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConnectedRegistry(ctx *pulumi.Context, args *LookupConnectedRegistryArgs, opts ...pulumi.InvokeOption) (*LookupConnectedRegistryResult, error) {
	var rv LookupConnectedRegistryResult
	err := ctx.Invoke("azure-native:containerregistry:getConnectedRegistry", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupConnectedRegistryArgs struct {
	ConnectedRegistryName string `pulumi:"connectedRegistryName"`
	RegistryName          string `pulumi:"registryName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupConnectedRegistryResult struct {
	Activation        ActivationPropertiesResponse     `pulumi:"activation"`
	ClientTokenIds    []string                         `pulumi:"clientTokenIds"`
	ConnectionState   string                           `pulumi:"connectionState"`
	Id                string                           `pulumi:"id"`
	LastActivityTime  string                           `pulumi:"lastActivityTime"`
	Logging           *LoggingPropertiesResponse       `pulumi:"logging"`
	LoginServer       *LoginServerPropertiesResponse   `pulumi:"loginServer"`
	Mode              string                           `pulumi:"mode"`
	Name              string                           `pulumi:"name"`
	Parent            ParentPropertiesResponse         `pulumi:"parent"`
	ProvisioningState string                           `pulumi:"provisioningState"`
	StatusDetails     []StatusDetailPropertiesResponse `pulumi:"statusDetails"`
	SystemData        SystemDataResponse               `pulumi:"systemData"`
	Type              string                           `pulumi:"type"`
	Version           string                           `pulumi:"version"`
}


func (val *LookupConnectedRegistryResult) Defaults() *LookupConnectedRegistryResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Logging = tmp.Logging.Defaults()

	return &tmp
}
