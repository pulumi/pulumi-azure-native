


package v20210501

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupOuContainer(ctx *pulumi.Context, args *LookupOuContainerArgs, opts ...pulumi.InvokeOption) (*LookupOuContainerResult, error) {
	var rv LookupOuContainerResult
	err := ctx.Invoke("azure-native:aad/v20210501:getOuContainer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOuContainerArgs struct {
	DomainServiceName string `pulumi:"domainServiceName"`
	OuContainerName   string `pulumi:"ouContainerName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupOuContainerResult struct {
	Accounts          []ContainerAccountResponse `pulumi:"accounts"`
	ContainerId       string                     `pulumi:"containerId"`
	DeploymentId      string                     `pulumi:"deploymentId"`
	DistinguishedName string                     `pulumi:"distinguishedName"`
	DomainName        string                     `pulumi:"domainName"`
	Etag              *string                    `pulumi:"etag"`
	Id                string                     `pulumi:"id"`
	Location          *string                    `pulumi:"location"`
	Name              string                     `pulumi:"name"`
	ProvisioningState string                     `pulumi:"provisioningState"`
	ServiceStatus     string                     `pulumi:"serviceStatus"`
	SystemData        SystemDataResponse         `pulumi:"systemData"`
	Tags              map[string]string          `pulumi:"tags"`
	TenantId          string                     `pulumi:"tenantId"`
	Type              string                     `pulumi:"type"`
}
