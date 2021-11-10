


package v20210115

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type StaticSiteCustomDomain struct {
	pulumi.CustomResourceState

	CreatedOn       pulumi.StringOutput    `pulumi:"createdOn"`
	DomainName      pulumi.StringOutput    `pulumi:"domainName"`
	ErrorMessage    pulumi.StringOutput    `pulumi:"errorMessage"`
	Kind            pulumi.StringPtrOutput `pulumi:"kind"`
	Name            pulumi.StringOutput    `pulumi:"name"`
	Status          pulumi.StringOutput    `pulumi:"status"`
	Type            pulumi.StringOutput    `pulumi:"type"`
	ValidationToken pulumi.StringOutput    `pulumi:"validationToken"`
}


func NewStaticSiteCustomDomain(ctx *pulumi.Context,
	name string, args *StaticSiteCustomDomainArgs, opts ...pulumi.ResourceOption) (*StaticSiteCustomDomain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ValidationMethod == nil {
		args.ValidationMethod = pulumi.StringPtr("cname-delegation")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:StaticSiteCustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:StaticSiteCustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:StaticSiteCustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:StaticSiteCustomDomain"),
		},
	})
	opts = append(opts, aliases)
	var resource StaticSiteCustomDomain
	err := ctx.RegisterResource("azure-native:web/v20210115:StaticSiteCustomDomain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetStaticSiteCustomDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StaticSiteCustomDomainState, opts ...pulumi.ResourceOption) (*StaticSiteCustomDomain, error) {
	var resource StaticSiteCustomDomain
	err := ctx.ReadResource("azure-native:web/v20210115:StaticSiteCustomDomain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type staticSiteCustomDomainState struct {
}

type StaticSiteCustomDomainState struct {
}

func (StaticSiteCustomDomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*staticSiteCustomDomainState)(nil)).Elem()
}

type staticSiteCustomDomainArgs struct {
	DomainName        *string `pulumi:"domainName"`
	Kind              *string `pulumi:"kind"`
	Name              string  `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ValidationMethod  *string `pulumi:"validationMethod"`
}


type StaticSiteCustomDomainArgs struct {
	DomainName        pulumi.StringPtrInput
	Kind              pulumi.StringPtrInput
	Name              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	ValidationMethod  pulumi.StringPtrInput
}

func (StaticSiteCustomDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*staticSiteCustomDomainArgs)(nil)).Elem()
}

type StaticSiteCustomDomainInput interface {
	pulumi.Input

	ToStaticSiteCustomDomainOutput() StaticSiteCustomDomainOutput
	ToStaticSiteCustomDomainOutputWithContext(ctx context.Context) StaticSiteCustomDomainOutput
}

func (*StaticSiteCustomDomain) ElementType() reflect.Type {
	return reflect.TypeOf((*StaticSiteCustomDomain)(nil))
}

func (i *StaticSiteCustomDomain) ToStaticSiteCustomDomainOutput() StaticSiteCustomDomainOutput {
	return i.ToStaticSiteCustomDomainOutputWithContext(context.Background())
}

func (i *StaticSiteCustomDomain) ToStaticSiteCustomDomainOutputWithContext(ctx context.Context) StaticSiteCustomDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticSiteCustomDomainOutput)
}

type StaticSiteCustomDomainOutput struct{ *pulumi.OutputState }

func (StaticSiteCustomDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StaticSiteCustomDomain)(nil))
}

func (o StaticSiteCustomDomainOutput) ToStaticSiteCustomDomainOutput() StaticSiteCustomDomainOutput {
	return o
}

func (o StaticSiteCustomDomainOutput) ToStaticSiteCustomDomainOutputWithContext(ctx context.Context) StaticSiteCustomDomainOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(StaticSiteCustomDomainOutput{})
}
