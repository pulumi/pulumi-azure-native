


package customerinsights

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLink(ctx *pulumi.Context, args *LookupLinkArgs, opts ...pulumi.InvokeOption) (*LookupLinkResult, error) {
	var rv LookupLinkResult
	err := ctx.Invoke("azure-native:customerinsights:getLink", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLinkArgs struct {
	HubName           string `pulumi:"hubName"`
	LinkName          string `pulumi:"linkName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupLinkResult struct {
	Description                   map[string]string                      `pulumi:"description"`
	DisplayName                   map[string]string                      `pulumi:"displayName"`
	Id                            string                                 `pulumi:"id"`
	LinkName                      string                                 `pulumi:"linkName"`
	Mappings                      []TypePropertiesMappingResponse        `pulumi:"mappings"`
	Name                          string                                 `pulumi:"name"`
	OperationType                 *string                                `pulumi:"operationType"`
	ParticipantPropertyReferences []ParticipantPropertyReferenceResponse `pulumi:"participantPropertyReferences"`
	ProvisioningState             string                                 `pulumi:"provisioningState"`
	ReferenceOnly                 *bool                                  `pulumi:"referenceOnly"`
	SourceEntityType              string                                 `pulumi:"sourceEntityType"`
	SourceEntityTypeName          string                                 `pulumi:"sourceEntityTypeName"`
	TargetEntityType              string                                 `pulumi:"targetEntityType"`
	TargetEntityTypeName          string                                 `pulumi:"targetEntityTypeName"`
	TenantId                      string                                 `pulumi:"tenantId"`
	Type                          string                                 `pulumi:"type"`
}
