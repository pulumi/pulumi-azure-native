


package v20201015preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PartnerNamespace struct {
	pulumi.CustomResourceState

	Endpoint                            pulumi.StringOutput      `pulumi:"endpoint"`
	Location                            pulumi.StringOutput      `pulumi:"location"`
	Name                                pulumi.StringOutput      `pulumi:"name"`
	PartnerRegistrationFullyQualifiedId pulumi.StringPtrOutput   `pulumi:"partnerRegistrationFullyQualifiedId"`
	ProvisioningState                   pulumi.StringOutput      `pulumi:"provisioningState"`
	SystemData                          SystemDataResponseOutput `pulumi:"systemData"`
	Tags                                pulumi.StringMapOutput   `pulumi:"tags"`
	Type                                pulumi.StringOutput      `pulumi:"type"`
}


func NewPartnerNamespace(ctx *pulumi.Context,
	name string, args *PartnerNamespaceArgs, opts ...pulumi.ResourceOption) (*PartnerNamespace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:eventgrid:PartnerNamespace"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20200401preview:PartnerNamespace"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20210601preview:PartnerNamespace"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20211015preview:PartnerNamespace"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20220615:PartnerNamespace"),
		},
	})
	opts = append(opts, aliases)
	var resource PartnerNamespace
	err := ctx.RegisterResource("azure-native:eventgrid/v20201015preview:PartnerNamespace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPartnerNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PartnerNamespaceState, opts ...pulumi.ResourceOption) (*PartnerNamespace, error) {
	var resource PartnerNamespace
	err := ctx.ReadResource("azure-native:eventgrid/v20201015preview:PartnerNamespace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type partnerNamespaceState struct {
}

type PartnerNamespaceState struct {
}

func (PartnerNamespaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*partnerNamespaceState)(nil)).Elem()
}

type partnerNamespaceArgs struct {
	Location                            *string           `pulumi:"location"`
	PartnerNamespaceName                *string           `pulumi:"partnerNamespaceName"`
	PartnerRegistrationFullyQualifiedId *string           `pulumi:"partnerRegistrationFullyQualifiedId"`
	ResourceGroupName                   string            `pulumi:"resourceGroupName"`
	Tags                                map[string]string `pulumi:"tags"`
}


type PartnerNamespaceArgs struct {
	Location                            pulumi.StringPtrInput
	PartnerNamespaceName                pulumi.StringPtrInput
	PartnerRegistrationFullyQualifiedId pulumi.StringPtrInput
	ResourceGroupName                   pulumi.StringInput
	Tags                                pulumi.StringMapInput
}

func (PartnerNamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*partnerNamespaceArgs)(nil)).Elem()
}

type PartnerNamespaceInput interface {
	pulumi.Input

	ToPartnerNamespaceOutput() PartnerNamespaceOutput
	ToPartnerNamespaceOutputWithContext(ctx context.Context) PartnerNamespaceOutput
}

func (*PartnerNamespace) ElementType() reflect.Type {
	return reflect.TypeOf((**PartnerNamespace)(nil)).Elem()
}

func (i *PartnerNamespace) ToPartnerNamespaceOutput() PartnerNamespaceOutput {
	return i.ToPartnerNamespaceOutputWithContext(context.Background())
}

func (i *PartnerNamespace) ToPartnerNamespaceOutputWithContext(ctx context.Context) PartnerNamespaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartnerNamespaceOutput)
}

type PartnerNamespaceOutput struct{ *pulumi.OutputState }

func (PartnerNamespaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PartnerNamespace)(nil)).Elem()
}

func (o PartnerNamespaceOutput) ToPartnerNamespaceOutput() PartnerNamespaceOutput {
	return o
}

func (o PartnerNamespaceOutput) ToPartnerNamespaceOutputWithContext(ctx context.Context) PartnerNamespaceOutput {
	return o
}

func (o PartnerNamespaceOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *PartnerNamespace) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

func (o PartnerNamespaceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *PartnerNamespace) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o PartnerNamespaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PartnerNamespace) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PartnerNamespaceOutput) PartnerRegistrationFullyQualifiedId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerNamespace) pulumi.StringPtrOutput { return v.PartnerRegistrationFullyQualifiedId }).(pulumi.StringPtrOutput)
}

func (o PartnerNamespaceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PartnerNamespace) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PartnerNamespaceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PartnerNamespace) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o PartnerNamespaceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PartnerNamespace) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o PartnerNamespaceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PartnerNamespace) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PartnerNamespaceOutput{})
}
