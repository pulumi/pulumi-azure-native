


package v20210101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DomainOwnershipIdentifier struct {
	pulumi.CustomResourceState

	Kind        pulumi.StringPtrOutput `pulumi:"kind"`
	Name        pulumi.StringOutput    `pulumi:"name"`
	OwnershipId pulumi.StringPtrOutput `pulumi:"ownershipId"`
	Type        pulumi.StringOutput    `pulumi:"type"`
}


func NewDomainOwnershipIdentifier(ctx *pulumi.Context,
	name string, args *DomainOwnershipIdentifierArgs, opts ...pulumi.ResourceOption) (*DomainOwnershipIdentifier, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:domainregistration:DomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-native:domainregistration/v20150401:DomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-native:domainregistration/v20180201:DomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-native:domainregistration/v20190801:DomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-native:domainregistration/v20200601:DomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-native:domainregistration/v20200901:DomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-native:domainregistration/v20201001:DomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-native:domainregistration/v20201201:DomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-native:domainregistration/v20210115:DomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-native:domainregistration/v20210201:DomainOwnershipIdentifier"),
		},
	})
	opts = append(opts, aliases)
	var resource DomainOwnershipIdentifier
	err := ctx.RegisterResource("azure-native:domainregistration/v20210101:DomainOwnershipIdentifier", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDomainOwnershipIdentifier(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainOwnershipIdentifierState, opts ...pulumi.ResourceOption) (*DomainOwnershipIdentifier, error) {
	var resource DomainOwnershipIdentifier
	err := ctx.ReadResource("azure-native:domainregistration/v20210101:DomainOwnershipIdentifier", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type domainOwnershipIdentifierState struct {
}

type DomainOwnershipIdentifierState struct {
}

func (DomainOwnershipIdentifierState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainOwnershipIdentifierState)(nil)).Elem()
}

type domainOwnershipIdentifierArgs struct {
	DomainName        string  `pulumi:"domainName"`
	Kind              *string `pulumi:"kind"`
	Name              *string `pulumi:"name"`
	OwnershipId       *string `pulumi:"ownershipId"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type DomainOwnershipIdentifierArgs struct {
	DomainName        pulumi.StringInput
	Kind              pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	OwnershipId       pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
}

func (DomainOwnershipIdentifierArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainOwnershipIdentifierArgs)(nil)).Elem()
}

type DomainOwnershipIdentifierInput interface {
	pulumi.Input

	ToDomainOwnershipIdentifierOutput() DomainOwnershipIdentifierOutput
	ToDomainOwnershipIdentifierOutputWithContext(ctx context.Context) DomainOwnershipIdentifierOutput
}

func (*DomainOwnershipIdentifier) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainOwnershipIdentifier)(nil))
}

func (i *DomainOwnershipIdentifier) ToDomainOwnershipIdentifierOutput() DomainOwnershipIdentifierOutput {
	return i.ToDomainOwnershipIdentifierOutputWithContext(context.Background())
}

func (i *DomainOwnershipIdentifier) ToDomainOwnershipIdentifierOutputWithContext(ctx context.Context) DomainOwnershipIdentifierOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainOwnershipIdentifierOutput)
}

type DomainOwnershipIdentifierOutput struct{ *pulumi.OutputState }

func (DomainOwnershipIdentifierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainOwnershipIdentifier)(nil))
}

func (o DomainOwnershipIdentifierOutput) ToDomainOwnershipIdentifierOutput() DomainOwnershipIdentifierOutput {
	return o
}

func (o DomainOwnershipIdentifierOutput) ToDomainOwnershipIdentifierOutputWithContext(ctx context.Context) DomainOwnershipIdentifierOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DomainOwnershipIdentifierOutput{})
}
