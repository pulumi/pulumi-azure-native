


package customerinsights

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRelationshipLink(ctx *pulumi.Context, args *LookupRelationshipLinkArgs, opts ...pulumi.InvokeOption) (*LookupRelationshipLinkResult, error) {
	var rv LookupRelationshipLinkResult
	err := ctx.Invoke("azure-native:customerinsights:getRelationshipLink", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRelationshipLinkArgs struct {
	HubName              string `pulumi:"hubName"`
	RelationshipLinkName string `pulumi:"relationshipLinkName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupRelationshipLinkResult struct {
	Description                      map[string]string                             `pulumi:"description"`
	DisplayName                      map[string]string                             `pulumi:"displayName"`
	Id                               string                                        `pulumi:"id"`
	InteractionType                  string                                        `pulumi:"interactionType"`
	LinkName                         string                                        `pulumi:"linkName"`
	Mappings                         []RelationshipLinkFieldMappingResponse        `pulumi:"mappings"`
	Name                             string                                        `pulumi:"name"`
	ProfilePropertyReferences        []ParticipantProfilePropertyReferenceResponse `pulumi:"profilePropertyReferences"`
	ProvisioningState                string                                        `pulumi:"provisioningState"`
	RelatedProfilePropertyReferences []ParticipantProfilePropertyReferenceResponse `pulumi:"relatedProfilePropertyReferences"`
	RelationshipGuidId               string                                        `pulumi:"relationshipGuidId"`
	RelationshipName                 string                                        `pulumi:"relationshipName"`
	TenantId                         string                                        `pulumi:"tenantId"`
	Type                             string                                        `pulumi:"type"`
}
