


package powerplatform

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEnterprisePolicy(ctx *pulumi.Context, args *LookupEnterprisePolicyArgs, opts ...pulumi.InvokeOption) (*LookupEnterprisePolicyResult, error) {
	var rv LookupEnterprisePolicyResult
	err := ctx.Invoke("azure-native:powerplatform:getEnterprisePolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEnterprisePolicyArgs struct {
	EnterprisePolicyName string `pulumi:"enterprisePolicyName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupEnterprisePolicyResult struct {
	Encryption       *PropertiesResponseEncryption       `pulumi:"encryption"`
	Id               string                              `pulumi:"id"`
	Identity         *EnterprisePolicyIdentityResponse   `pulumi:"identity"`
	Kind             string                              `pulumi:"kind"`
	Location         string                              `pulumi:"location"`
	Lockbox          *PropertiesResponseLockbox          `pulumi:"lockbox"`
	Name             string                              `pulumi:"name"`
	NetworkInjection *PropertiesResponseNetworkInjection `pulumi:"networkInjection"`
	SystemData       SystemDataResponse                  `pulumi:"systemData"`
	Tags             map[string]string                   `pulumi:"tags"`
	Type             string                              `pulumi:"type"`
}
