


package containerregistry

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupScopeMap(ctx *pulumi.Context, args *LookupScopeMapArgs, opts ...pulumi.InvokeOption) (*LookupScopeMapResult, error) {
	var rv LookupScopeMapResult
	err := ctx.Invoke("azure-native:containerregistry:getScopeMap", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScopeMapArgs struct {
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ScopeMapName      string `pulumi:"scopeMapName"`
}


type LookupScopeMapResult struct {
	Actions           []string           `pulumi:"actions"`
	CreationDate      string             `pulumi:"creationDate"`
	Description       *string            `pulumi:"description"`
	Id                string             `pulumi:"id"`
	Name              string             `pulumi:"name"`
	ProvisioningState string             `pulumi:"provisioningState"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	Type              string             `pulumi:"type"`
}
