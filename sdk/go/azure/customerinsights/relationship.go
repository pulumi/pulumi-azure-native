


package customerinsights

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Relationship struct {
	pulumi.CustomResourceState

	Cardinality        pulumi.StringPtrOutput                     `pulumi:"cardinality"`
	Description        pulumi.StringMapOutput                     `pulumi:"description"`
	DisplayName        pulumi.StringMapOutput                     `pulumi:"displayName"`
	ExpiryDateTimeUtc  pulumi.StringPtrOutput                     `pulumi:"expiryDateTimeUtc"`
	Fields             PropertyDefinitionResponseArrayOutput      `pulumi:"fields"`
	LookupMappings     RelationshipTypeMappingResponseArrayOutput `pulumi:"lookupMappings"`
	Name               pulumi.StringOutput                        `pulumi:"name"`
	ProfileType        pulumi.StringOutput                        `pulumi:"profileType"`
	ProvisioningState  pulumi.StringOutput                        `pulumi:"provisioningState"`
	RelatedProfileType pulumi.StringOutput                        `pulumi:"relatedProfileType"`
	RelationshipGuidId pulumi.StringOutput                        `pulumi:"relationshipGuidId"`
	RelationshipName   pulumi.StringOutput                        `pulumi:"relationshipName"`
	TenantId           pulumi.StringOutput                        `pulumi:"tenantId"`
	Type               pulumi.StringOutput                        `pulumi:"type"`
}


func NewRelationship(ctx *pulumi.Context,
	name string, args *RelationshipArgs, opts ...pulumi.ResourceOption) (*Relationship, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HubName == nil {
		return nil, errors.New("invalid value for required argument 'HubName'")
	}
	if args.ProfileType == nil {
		return nil, errors.New("invalid value for required argument 'ProfileType'")
	}
	if args.RelatedProfileType == nil {
		return nil, errors.New("invalid value for required argument 'RelatedProfileType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:customerinsights/v20170101:Relationship"),
		},
		{
			Type: pulumi.String("azure-native:customerinsights/v20170426:Relationship"),
		},
	})
	opts = append(opts, aliases)
	var resource Relationship
	err := ctx.RegisterResource("azure-native:customerinsights:Relationship", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRelationship(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RelationshipState, opts ...pulumi.ResourceOption) (*Relationship, error) {
	var resource Relationship
	err := ctx.ReadResource("azure-native:customerinsights:Relationship", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type relationshipState struct {
}

type RelationshipState struct {
}

func (RelationshipState) ElementType() reflect.Type {
	return reflect.TypeOf((*relationshipState)(nil)).Elem()
}

type relationshipArgs struct {
	Cardinality        *CardinalityTypes         `pulumi:"cardinality"`
	Description        map[string]string         `pulumi:"description"`
	DisplayName        map[string]string         `pulumi:"displayName"`
	ExpiryDateTimeUtc  *string                   `pulumi:"expiryDateTimeUtc"`
	Fields             []PropertyDefinition      `pulumi:"fields"`
	HubName            string                    `pulumi:"hubName"`
	LookupMappings     []RelationshipTypeMapping `pulumi:"lookupMappings"`
	ProfileType        string                    `pulumi:"profileType"`
	RelatedProfileType string                    `pulumi:"relatedProfileType"`
	RelationshipName   *string                   `pulumi:"relationshipName"`
	ResourceGroupName  string                    `pulumi:"resourceGroupName"`
}


type RelationshipArgs struct {
	Cardinality        CardinalityTypesPtrInput
	Description        pulumi.StringMapInput
	DisplayName        pulumi.StringMapInput
	ExpiryDateTimeUtc  pulumi.StringPtrInput
	Fields             PropertyDefinitionArrayInput
	HubName            pulumi.StringInput
	LookupMappings     RelationshipTypeMappingArrayInput
	ProfileType        pulumi.StringInput
	RelatedProfileType pulumi.StringInput
	RelationshipName   pulumi.StringPtrInput
	ResourceGroupName  pulumi.StringInput
}

func (RelationshipArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*relationshipArgs)(nil)).Elem()
}

type RelationshipInput interface {
	pulumi.Input

	ToRelationshipOutput() RelationshipOutput
	ToRelationshipOutputWithContext(ctx context.Context) RelationshipOutput
}

func (*Relationship) ElementType() reflect.Type {
	return reflect.TypeOf((**Relationship)(nil)).Elem()
}

func (i *Relationship) ToRelationshipOutput() RelationshipOutput {
	return i.ToRelationshipOutputWithContext(context.Background())
}

func (i *Relationship) ToRelationshipOutputWithContext(ctx context.Context) RelationshipOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RelationshipOutput)
}

type RelationshipOutput struct{ *pulumi.OutputState }

func (RelationshipOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Relationship)(nil)).Elem()
}

func (o RelationshipOutput) ToRelationshipOutput() RelationshipOutput {
	return o
}

func (o RelationshipOutput) ToRelationshipOutputWithContext(ctx context.Context) RelationshipOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RelationshipOutput{})
}
