


package cdn

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAFDCustomDomain(ctx *pulumi.Context, args *LookupAFDCustomDomainArgs, opts ...pulumi.InvokeOption) (*LookupAFDCustomDomainResult, error) {
	var rv LookupAFDCustomDomainResult
	err := ctx.Invoke("azure-native:cdn:getAFDCustomDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAFDCustomDomainArgs struct {
	CustomDomainName  string `pulumi:"customDomainName"`
	ProfileName       string `pulumi:"profileName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAFDCustomDomainResult struct {
	AzureDnsZone          *ResourceReferenceResponse         `pulumi:"azureDnsZone"`
	DeploymentStatus      string                             `pulumi:"deploymentStatus"`
	DomainValidationState string                             `pulumi:"domainValidationState"`
	HostName              string                             `pulumi:"hostName"`
	Id                    string                             `pulumi:"id"`
	Name                  string                             `pulumi:"name"`
	ProvisioningState     string                             `pulumi:"provisioningState"`
	SystemData            SystemDataResponse                 `pulumi:"systemData"`
	TlsSettings           *AFDDomainHttpsParametersResponse  `pulumi:"tlsSettings"`
	Type                  string                             `pulumi:"type"`
	ValidationProperties  DomainValidationPropertiesResponse `pulumi:"validationProperties"`
}
