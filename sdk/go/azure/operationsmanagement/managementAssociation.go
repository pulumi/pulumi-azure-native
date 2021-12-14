


package operationsmanagement

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ManagementAssociation struct {
	pulumi.CustomResourceState

	Location   pulumi.StringPtrOutput                        `pulumi:"location"`
	Name       pulumi.StringOutput                           `pulumi:"name"`
	Properties ManagementAssociationPropertiesResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput                           `pulumi:"type"`
}


func NewManagementAssociation(ctx *pulumi.Context,
	name string, args *ManagementAssociationArgs, opts ...pulumi.ResourceOption) (*ManagementAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProviderName == nil {
		return nil, errors.New("invalid value for required argument 'ProviderName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	if args.ResourceType == nil {
		return nil, errors.New("invalid value for required argument 'ResourceType'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:operationsmanagement/v20151101preview:ManagementAssociation"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagementAssociation
	err := ctx.RegisterResource("azure-native:operationsmanagement:ManagementAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetManagementAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagementAssociationState, opts ...pulumi.ResourceOption) (*ManagementAssociation, error) {
	var resource ManagementAssociation
	err := ctx.ReadResource("azure-native:operationsmanagement:ManagementAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type managementAssociationState struct {
}

type ManagementAssociationState struct {
}

func (ManagementAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*managementAssociationState)(nil)).Elem()
}

type managementAssociationArgs struct {
	Location                  *string                          `pulumi:"location"`
	ManagementAssociationName *string                          `pulumi:"managementAssociationName"`
	Properties                *ManagementAssociationProperties `pulumi:"properties"`
	ProviderName              string                           `pulumi:"providerName"`
	ResourceGroupName         string                           `pulumi:"resourceGroupName"`
	ResourceName              string                           `pulumi:"resourceName"`
	ResourceType              string                           `pulumi:"resourceType"`
}


type ManagementAssociationArgs struct {
	Location                  pulumi.StringPtrInput
	ManagementAssociationName pulumi.StringPtrInput
	Properties                ManagementAssociationPropertiesPtrInput
	ProviderName              pulumi.StringInput
	ResourceGroupName         pulumi.StringInput
	ResourceName              pulumi.StringInput
	ResourceType              pulumi.StringInput
}

func (ManagementAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managementAssociationArgs)(nil)).Elem()
}

type ManagementAssociationInput interface {
	pulumi.Input

	ToManagementAssociationOutput() ManagementAssociationOutput
	ToManagementAssociationOutputWithContext(ctx context.Context) ManagementAssociationOutput
}

func (*ManagementAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementAssociation)(nil)).Elem()
}

func (i *ManagementAssociation) ToManagementAssociationOutput() ManagementAssociationOutput {
	return i.ToManagementAssociationOutputWithContext(context.Background())
}

func (i *ManagementAssociation) ToManagementAssociationOutputWithContext(ctx context.Context) ManagementAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementAssociationOutput)
}

type ManagementAssociationOutput struct{ *pulumi.OutputState }

func (ManagementAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementAssociation)(nil)).Elem()
}

func (o ManagementAssociationOutput) ToManagementAssociationOutput() ManagementAssociationOutput {
	return o
}

func (o ManagementAssociationOutput) ToManagementAssociationOutputWithContext(ctx context.Context) ManagementAssociationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ManagementAssociationOutput{})
}
