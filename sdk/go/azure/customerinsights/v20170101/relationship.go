


package v20170101

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
			Type: pulumi.String("azure-native:customerinsights:Relationship"),
		},
		{
			Type: pulumi.String("azure-native:customerinsights/v20170426:Relationship"),
		},
	})
	opts = append(opts, aliases)
	var resource Relationship
	err := ctx.RegisterResource("azure-native:customerinsights/v20170101:Relationship", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRelationship(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RelationshipState, opts ...pulumi.ResourceOption) (*Relationship, error) {
	var resource Relationship
	err := ctx.ReadResource("azure-native:customerinsights/v20170101:Relationship", name, id, state, &resource, opts...)
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

func (o RelationshipOutput) Cardinality() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Relationship) pulumi.StringPtrOutput { return v.Cardinality }).(pulumi.StringPtrOutput)
}

func (o RelationshipOutput) Description() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Relationship) pulumi.StringMapOutput { return v.Description }).(pulumi.StringMapOutput)
}

func (o RelationshipOutput) DisplayName() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Relationship) pulumi.StringMapOutput { return v.DisplayName }).(pulumi.StringMapOutput)
}

func (o RelationshipOutput) ExpiryDateTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Relationship) pulumi.StringPtrOutput { return v.ExpiryDateTimeUtc }).(pulumi.StringPtrOutput)
}

func (o RelationshipOutput) Fields() PropertyDefinitionResponseArrayOutput {
	return o.ApplyT(func(v *Relationship) PropertyDefinitionResponseArrayOutput { return v.Fields }).(PropertyDefinitionResponseArrayOutput)
}

func (o RelationshipOutput) LookupMappings() RelationshipTypeMappingResponseArrayOutput {
	return o.ApplyT(func(v *Relationship) RelationshipTypeMappingResponseArrayOutput { return v.LookupMappings }).(RelationshipTypeMappingResponseArrayOutput)
}

func (o RelationshipOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Relationship) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RelationshipOutput) ProfileType() pulumi.StringOutput {
	return o.ApplyT(func(v *Relationship) pulumi.StringOutput { return v.ProfileType }).(pulumi.StringOutput)
}

func (o RelationshipOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Relationship) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o RelationshipOutput) RelatedProfileType() pulumi.StringOutput {
	return o.ApplyT(func(v *Relationship) pulumi.StringOutput { return v.RelatedProfileType }).(pulumi.StringOutput)
}

func (o RelationshipOutput) RelationshipGuidId() pulumi.StringOutput {
	return o.ApplyT(func(v *Relationship) pulumi.StringOutput { return v.RelationshipGuidId }).(pulumi.StringOutput)
}

func (o RelationshipOutput) RelationshipName() pulumi.StringOutput {
	return o.ApplyT(func(v *Relationship) pulumi.StringOutput { return v.RelationshipName }).(pulumi.StringOutput)
}

func (o RelationshipOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *Relationship) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

func (o RelationshipOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Relationship) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RelationshipOutput{})
}
