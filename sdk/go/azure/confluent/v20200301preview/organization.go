


package v20200301preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Organization struct {
	pulumi.CustomResourceState

	CreatedTime       pulumi.StringOutput                                        `pulumi:"createdTime"`
	Location          pulumi.StringPtrOutput                                     `pulumi:"location"`
	Name              pulumi.StringOutput                                        `pulumi:"name"`
	OfferDetail       OrganizationResourcePropertiesResponseOfferDetailPtrOutput `pulumi:"offerDetail"`
	OrganizationId    pulumi.StringOutput                                        `pulumi:"organizationId"`
	ProvisioningState pulumi.StringOutput                                        `pulumi:"provisioningState"`
	SsoUrl            pulumi.StringOutput                                        `pulumi:"ssoUrl"`
	Tags              pulumi.StringMapOutput                                     `pulumi:"tags"`
	Type              pulumi.StringOutput                                        `pulumi:"type"`
	UserDetail        OrganizationResourcePropertiesResponseUserDetailPtrOutput  `pulumi:"userDetail"`
}


func NewOrganization(ctx *pulumi.Context,
	name string, args *OrganizationArgs, opts ...pulumi.ResourceOption) (*Organization, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:confluent:Organization"),
		},
		{
			Type: pulumi.String("azure-native:confluent/v20200301:Organization"),
		},
		{
			Type: pulumi.String("azure-native:confluent/v20210301preview:Organization"),
		},
		{
			Type: pulumi.String("azure-native:confluent/v20210901preview:Organization"),
		},
		{
			Type: pulumi.String("azure-native:confluent/v20211201:Organization"),
		},
	})
	opts = append(opts, aliases)
	var resource Organization
	err := ctx.RegisterResource("azure-native:confluent/v20200301preview:Organization", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetOrganization(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationState, opts ...pulumi.ResourceOption) (*Organization, error) {
	var resource Organization
	err := ctx.ReadResource("azure-native:confluent/v20200301preview:Organization", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type organizationState struct {
}

type OrganizationState struct {
}

func (OrganizationState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationState)(nil)).Elem()
}

type organizationArgs struct {
	Location          *string                                    `pulumi:"location"`
	OfferDetail       *OrganizationResourcePropertiesOfferDetail `pulumi:"offerDetail"`
	OrganizationName  *string                                    `pulumi:"organizationName"`
	ResourceGroupName string                                     `pulumi:"resourceGroupName"`
	Tags              map[string]string                          `pulumi:"tags"`
	UserDetail        *OrganizationResourcePropertiesUserDetail  `pulumi:"userDetail"`
}


type OrganizationArgs struct {
	Location          pulumi.StringPtrInput
	OfferDetail       OrganizationResourcePropertiesOfferDetailPtrInput
	OrganizationName  pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	UserDetail        OrganizationResourcePropertiesUserDetailPtrInput
}

func (OrganizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationArgs)(nil)).Elem()
}

type OrganizationInput interface {
	pulumi.Input

	ToOrganizationOutput() OrganizationOutput
	ToOrganizationOutputWithContext(ctx context.Context) OrganizationOutput
}

func (*Organization) ElementType() reflect.Type {
	return reflect.TypeOf((**Organization)(nil)).Elem()
}

func (i *Organization) ToOrganizationOutput() OrganizationOutput {
	return i.ToOrganizationOutputWithContext(context.Background())
}

func (i *Organization) ToOrganizationOutputWithContext(ctx context.Context) OrganizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationOutput)
}

type OrganizationOutput struct{ *pulumi.OutputState }

func (OrganizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Organization)(nil)).Elem()
}

func (o OrganizationOutput) ToOrganizationOutput() OrganizationOutput {
	return o
}

func (o OrganizationOutput) ToOrganizationOutputWithContext(ctx context.Context) OrganizationOutput {
	return o
}

func (o OrganizationOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Organization) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o OrganizationOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Organization) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o OrganizationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Organization) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o OrganizationOutput) OfferDetail() OrganizationResourcePropertiesResponseOfferDetailPtrOutput {
	return o.ApplyT(func(v *Organization) OrganizationResourcePropertiesResponseOfferDetailPtrOutput { return v.OfferDetail }).(OrganizationResourcePropertiesResponseOfferDetailPtrOutput)
}

func (o OrganizationOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Organization) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o OrganizationOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Organization) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o OrganizationOutput) SsoUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Organization) pulumi.StringOutput { return v.SsoUrl }).(pulumi.StringOutput)
}

func (o OrganizationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Organization) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o OrganizationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Organization) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o OrganizationOutput) UserDetail() OrganizationResourcePropertiesResponseUserDetailPtrOutput {
	return o.ApplyT(func(v *Organization) OrganizationResourcePropertiesResponseUserDetailPtrOutput { return v.UserDetail }).(OrganizationResourcePropertiesResponseUserDetailPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(OrganizationOutput{})
}
