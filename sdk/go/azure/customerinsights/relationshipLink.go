


package customerinsights

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RelationshipLink struct {
	pulumi.CustomResourceState

	Description                      pulumi.StringMapOutput                                 `pulumi:"description"`
	DisplayName                      pulumi.StringMapOutput                                 `pulumi:"displayName"`
	InteractionType                  pulumi.StringOutput                                    `pulumi:"interactionType"`
	LinkName                         pulumi.StringOutput                                    `pulumi:"linkName"`
	Mappings                         RelationshipLinkFieldMappingResponseArrayOutput        `pulumi:"mappings"`
	Name                             pulumi.StringOutput                                    `pulumi:"name"`
	ProfilePropertyReferences        ParticipantProfilePropertyReferenceResponseArrayOutput `pulumi:"profilePropertyReferences"`
	ProvisioningState                pulumi.StringOutput                                    `pulumi:"provisioningState"`
	RelatedProfilePropertyReferences ParticipantProfilePropertyReferenceResponseArrayOutput `pulumi:"relatedProfilePropertyReferences"`
	RelationshipGuidId               pulumi.StringOutput                                    `pulumi:"relationshipGuidId"`
	RelationshipName                 pulumi.StringOutput                                    `pulumi:"relationshipName"`
	TenantId                         pulumi.StringOutput                                    `pulumi:"tenantId"`
	Type                             pulumi.StringOutput                                    `pulumi:"type"`
}


func NewRelationshipLink(ctx *pulumi.Context,
	name string, args *RelationshipLinkArgs, opts ...pulumi.ResourceOption) (*RelationshipLink, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HubName == nil {
		return nil, errors.New("invalid value for required argument 'HubName'")
	}
	if args.InteractionType == nil {
		return nil, errors.New("invalid value for required argument 'InteractionType'")
	}
	if args.ProfilePropertyReferences == nil {
		return nil, errors.New("invalid value for required argument 'ProfilePropertyReferences'")
	}
	if args.RelatedProfilePropertyReferences == nil {
		return nil, errors.New("invalid value for required argument 'RelatedProfilePropertyReferences'")
	}
	if args.RelationshipName == nil {
		return nil, errors.New("invalid value for required argument 'RelationshipName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:customerinsights:RelationshipLink"),
		},
		{
			Type: pulumi.String("azure-native:customerinsights/v20170101:RelationshipLink"),
		},
		{
			Type: pulumi.String("azure-nextgen:customerinsights/v20170101:RelationshipLink"),
		},
		{
			Type: pulumi.String("azure-native:customerinsights/v20170426:RelationshipLink"),
		},
		{
			Type: pulumi.String("azure-nextgen:customerinsights/v20170426:RelationshipLink"),
		},
	})
	opts = append(opts, aliases)
	var resource RelationshipLink
	err := ctx.RegisterResource("azure-native:customerinsights:RelationshipLink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRelationshipLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RelationshipLinkState, opts ...pulumi.ResourceOption) (*RelationshipLink, error) {
	var resource RelationshipLink
	err := ctx.ReadResource("azure-native:customerinsights:RelationshipLink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type relationshipLinkState struct {
}

type RelationshipLinkState struct {
}

func (RelationshipLinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*relationshipLinkState)(nil)).Elem()
}

type relationshipLinkArgs struct {
	Description                      map[string]string                     `pulumi:"description"`
	DisplayName                      map[string]string                     `pulumi:"displayName"`
	HubName                          string                                `pulumi:"hubName"`
	InteractionType                  string                                `pulumi:"interactionType"`
	Mappings                         []RelationshipLinkFieldMapping        `pulumi:"mappings"`
	ProfilePropertyReferences        []ParticipantProfilePropertyReference `pulumi:"profilePropertyReferences"`
	RelatedProfilePropertyReferences []ParticipantProfilePropertyReference `pulumi:"relatedProfilePropertyReferences"`
	RelationshipLinkName             *string                               `pulumi:"relationshipLinkName"`
	RelationshipName                 string                                `pulumi:"relationshipName"`
	ResourceGroupName                string                                `pulumi:"resourceGroupName"`
}


type RelationshipLinkArgs struct {
	Description                      pulumi.StringMapInput
	DisplayName                      pulumi.StringMapInput
	HubName                          pulumi.StringInput
	InteractionType                  pulumi.StringInput
	Mappings                         RelationshipLinkFieldMappingArrayInput
	ProfilePropertyReferences        ParticipantProfilePropertyReferenceArrayInput
	RelatedProfilePropertyReferences ParticipantProfilePropertyReferenceArrayInput
	RelationshipLinkName             pulumi.StringPtrInput
	RelationshipName                 pulumi.StringInput
	ResourceGroupName                pulumi.StringInput
}

func (RelationshipLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*relationshipLinkArgs)(nil)).Elem()
}

type RelationshipLinkInput interface {
	pulumi.Input

	ToRelationshipLinkOutput() RelationshipLinkOutput
	ToRelationshipLinkOutputWithContext(ctx context.Context) RelationshipLinkOutput
}

func (*RelationshipLink) ElementType() reflect.Type {
	return reflect.TypeOf((*RelationshipLink)(nil))
}

func (i *RelationshipLink) ToRelationshipLinkOutput() RelationshipLinkOutput {
	return i.ToRelationshipLinkOutputWithContext(context.Background())
}

func (i *RelationshipLink) ToRelationshipLinkOutputWithContext(ctx context.Context) RelationshipLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RelationshipLinkOutput)
}

type RelationshipLinkOutput struct{ *pulumi.OutputState }

func (RelationshipLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RelationshipLink)(nil))
}

func (o RelationshipLinkOutput) ToRelationshipLinkOutput() RelationshipLinkOutput {
	return o
}

func (o RelationshipLinkOutput) ToRelationshipLinkOutputWithContext(ctx context.Context) RelationshipLinkOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RelationshipLinkOutput{})
}
