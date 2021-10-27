


package v20170101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Link struct {
	pulumi.CustomResourceState

	Description                   pulumi.StringMapOutput                          `pulumi:"description"`
	DisplayName                   pulumi.StringMapOutput                          `pulumi:"displayName"`
	LinkName                      pulumi.StringOutput                             `pulumi:"linkName"`
	Mappings                      TypePropertiesMappingResponseArrayOutput        `pulumi:"mappings"`
	Name                          pulumi.StringOutput                             `pulumi:"name"`
	OperationType                 pulumi.StringPtrOutput                          `pulumi:"operationType"`
	ParticipantPropertyReferences ParticipantPropertyReferenceResponseArrayOutput `pulumi:"participantPropertyReferences"`
	ProvisioningState             pulumi.StringOutput                             `pulumi:"provisioningState"`
	ReferenceOnly                 pulumi.BoolPtrOutput                            `pulumi:"referenceOnly"`
	SourceInteractionType         pulumi.StringOutput                             `pulumi:"sourceInteractionType"`
	TargetProfileType             pulumi.StringOutput                             `pulumi:"targetProfileType"`
	TenantId                      pulumi.StringOutput                             `pulumi:"tenantId"`
	Type                          pulumi.StringOutput                             `pulumi:"type"`
}


func NewLink(ctx *pulumi.Context,
	name string, args *LinkArgs, opts ...pulumi.ResourceOption) (*Link, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HubName == nil {
		return nil, errors.New("invalid value for required argument 'HubName'")
	}
	if args.ParticipantPropertyReferences == nil {
		return nil, errors.New("invalid value for required argument 'ParticipantPropertyReferences'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SourceInteractionType == nil {
		return nil, errors.New("invalid value for required argument 'SourceInteractionType'")
	}
	if args.TargetProfileType == nil {
		return nil, errors.New("invalid value for required argument 'TargetProfileType'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:customerinsights/v20170101:Link"),
		},
		{
			Type: pulumi.String("azure-native:customerinsights:Link"),
		},
		{
			Type: pulumi.String("azure-nextgen:customerinsights:Link"),
		},
		{
			Type: pulumi.String("azure-native:customerinsights/v20170426:Link"),
		},
		{
			Type: pulumi.String("azure-nextgen:customerinsights/v20170426:Link"),
		},
	})
	opts = append(opts, aliases)
	var resource Link
	err := ctx.RegisterResource("azure-native:customerinsights/v20170101:Link", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkState, opts ...pulumi.ResourceOption) (*Link, error) {
	var resource Link
	err := ctx.ReadResource("azure-native:customerinsights/v20170101:Link", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type linkState struct {
}

type LinkState struct {
}

func (LinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkState)(nil)).Elem()
}

type linkArgs struct {
	Description                   map[string]string              `pulumi:"description"`
	DisplayName                   map[string]string              `pulumi:"displayName"`
	HubName                       string                         `pulumi:"hubName"`
	LinkName                      *string                        `pulumi:"linkName"`
	Mappings                      []TypePropertiesMapping        `pulumi:"mappings"`
	OperationType                 *InstanceOperationType         `pulumi:"operationType"`
	ParticipantPropertyReferences []ParticipantPropertyReference `pulumi:"participantPropertyReferences"`
	ReferenceOnly                 *bool                          `pulumi:"referenceOnly"`
	ResourceGroupName             string                         `pulumi:"resourceGroupName"`
	SourceInteractionType         string                         `pulumi:"sourceInteractionType"`
	TargetProfileType             string                         `pulumi:"targetProfileType"`
}


type LinkArgs struct {
	Description                   pulumi.StringMapInput
	DisplayName                   pulumi.StringMapInput
	HubName                       pulumi.StringInput
	LinkName                      pulumi.StringPtrInput
	Mappings                      TypePropertiesMappingArrayInput
	OperationType                 InstanceOperationTypePtrInput
	ParticipantPropertyReferences ParticipantPropertyReferenceArrayInput
	ReferenceOnly                 pulumi.BoolPtrInput
	ResourceGroupName             pulumi.StringInput
	SourceInteractionType         pulumi.StringInput
	TargetProfileType             pulumi.StringInput
}

func (LinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkArgs)(nil)).Elem()
}

type LinkInput interface {
	pulumi.Input

	ToLinkOutput() LinkOutput
	ToLinkOutputWithContext(ctx context.Context) LinkOutput
}

func (*Link) ElementType() reflect.Type {
	return reflect.TypeOf((*Link)(nil))
}

func (i *Link) ToLinkOutput() LinkOutput {
	return i.ToLinkOutputWithContext(context.Background())
}

func (i *Link) ToLinkOutputWithContext(ctx context.Context) LinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkOutput)
}

type LinkOutput struct{ *pulumi.OutputState }

func (LinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Link)(nil))
}

func (o LinkOutput) ToLinkOutput() LinkOutput {
	return o
}

func (o LinkOutput) ToLinkOutputWithContext(ctx context.Context) LinkOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LinkInput)(nil)).Elem(), &Link{})
	pulumi.RegisterOutputType(LinkOutput{})
}
