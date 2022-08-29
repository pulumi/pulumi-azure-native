


package eventgrid

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PartnerRegistration struct {
	pulumi.CustomResourceState

	AuthorizedAzureSubscriptionIds  pulumi.StringArrayOutput `pulumi:"authorizedAzureSubscriptionIds"`
	CustomerServiceUri              pulumi.StringPtrOutput   `pulumi:"customerServiceUri"`
	Location                        pulumi.StringOutput      `pulumi:"location"`
	LogoUri                         pulumi.StringPtrOutput   `pulumi:"logoUri"`
	LongDescription                 pulumi.StringPtrOutput   `pulumi:"longDescription"`
	Name                            pulumi.StringOutput      `pulumi:"name"`
	PartnerCustomerServiceExtension pulumi.StringPtrOutput   `pulumi:"partnerCustomerServiceExtension"`
	PartnerCustomerServiceNumber    pulumi.StringPtrOutput   `pulumi:"partnerCustomerServiceNumber"`
	PartnerName                     pulumi.StringPtrOutput   `pulumi:"partnerName"`
	PartnerResourceTypeDescription  pulumi.StringPtrOutput   `pulumi:"partnerResourceTypeDescription"`
	PartnerResourceTypeDisplayName  pulumi.StringPtrOutput   `pulumi:"partnerResourceTypeDisplayName"`
	PartnerResourceTypeName         pulumi.StringPtrOutput   `pulumi:"partnerResourceTypeName"`
	ProvisioningState               pulumi.StringOutput      `pulumi:"provisioningState"`
	SetupUri                        pulumi.StringPtrOutput   `pulumi:"setupUri"`
	SystemData                      SystemDataResponseOutput `pulumi:"systemData"`
	Tags                            pulumi.StringMapOutput   `pulumi:"tags"`
	Type                            pulumi.StringOutput      `pulumi:"type"`
	VisibilityState                 pulumi.StringPtrOutput   `pulumi:"visibilityState"`
}


func NewPartnerRegistration(ctx *pulumi.Context,
	name string, args *PartnerRegistrationArgs, opts ...pulumi.ResourceOption) (*PartnerRegistration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:eventgrid/v20200401preview:PartnerRegistration"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20201015preview:PartnerRegistration"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20210601preview:PartnerRegistration"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20211015preview:PartnerRegistration"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20220615:PartnerRegistration"),
		},
	})
	opts = append(opts, aliases)
	var resource PartnerRegistration
	err := ctx.RegisterResource("azure-native:eventgrid:PartnerRegistration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPartnerRegistration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PartnerRegistrationState, opts ...pulumi.ResourceOption) (*PartnerRegistration, error) {
	var resource PartnerRegistration
	err := ctx.ReadResource("azure-native:eventgrid:PartnerRegistration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type partnerRegistrationState struct {
}

type PartnerRegistrationState struct {
}

func (PartnerRegistrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*partnerRegistrationState)(nil)).Elem()
}

type partnerRegistrationArgs struct {
	AuthorizedAzureSubscriptionIds  []string          `pulumi:"authorizedAzureSubscriptionIds"`
	CustomerServiceUri              *string           `pulumi:"customerServiceUri"`
	Location                        *string           `pulumi:"location"`
	LogoUri                         *string           `pulumi:"logoUri"`
	LongDescription                 *string           `pulumi:"longDescription"`
	PartnerCustomerServiceExtension *string           `pulumi:"partnerCustomerServiceExtension"`
	PartnerCustomerServiceNumber    *string           `pulumi:"partnerCustomerServiceNumber"`
	PartnerName                     *string           `pulumi:"partnerName"`
	PartnerRegistrationName         *string           `pulumi:"partnerRegistrationName"`
	PartnerResourceTypeDescription  *string           `pulumi:"partnerResourceTypeDescription"`
	PartnerResourceTypeDisplayName  *string           `pulumi:"partnerResourceTypeDisplayName"`
	PartnerResourceTypeName         *string           `pulumi:"partnerResourceTypeName"`
	ResourceGroupName               string            `pulumi:"resourceGroupName"`
	SetupUri                        *string           `pulumi:"setupUri"`
	Tags                            map[string]string `pulumi:"tags"`
	VisibilityState                 *string           `pulumi:"visibilityState"`
}


type PartnerRegistrationArgs struct {
	AuthorizedAzureSubscriptionIds  pulumi.StringArrayInput
	CustomerServiceUri              pulumi.StringPtrInput
	Location                        pulumi.StringPtrInput
	LogoUri                         pulumi.StringPtrInput
	LongDescription                 pulumi.StringPtrInput
	PartnerCustomerServiceExtension pulumi.StringPtrInput
	PartnerCustomerServiceNumber    pulumi.StringPtrInput
	PartnerName                     pulumi.StringPtrInput
	PartnerRegistrationName         pulumi.StringPtrInput
	PartnerResourceTypeDescription  pulumi.StringPtrInput
	PartnerResourceTypeDisplayName  pulumi.StringPtrInput
	PartnerResourceTypeName         pulumi.StringPtrInput
	ResourceGroupName               pulumi.StringInput
	SetupUri                        pulumi.StringPtrInput
	Tags                            pulumi.StringMapInput
	VisibilityState                 pulumi.StringPtrInput
}

func (PartnerRegistrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*partnerRegistrationArgs)(nil)).Elem()
}

type PartnerRegistrationInput interface {
	pulumi.Input

	ToPartnerRegistrationOutput() PartnerRegistrationOutput
	ToPartnerRegistrationOutputWithContext(ctx context.Context) PartnerRegistrationOutput
}

func (*PartnerRegistration) ElementType() reflect.Type {
	return reflect.TypeOf((**PartnerRegistration)(nil)).Elem()
}

func (i *PartnerRegistration) ToPartnerRegistrationOutput() PartnerRegistrationOutput {
	return i.ToPartnerRegistrationOutputWithContext(context.Background())
}

func (i *PartnerRegistration) ToPartnerRegistrationOutputWithContext(ctx context.Context) PartnerRegistrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartnerRegistrationOutput)
}

type PartnerRegistrationOutput struct{ *pulumi.OutputState }

func (PartnerRegistrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PartnerRegistration)(nil)).Elem()
}

func (o PartnerRegistrationOutput) ToPartnerRegistrationOutput() PartnerRegistrationOutput {
	return o
}

func (o PartnerRegistrationOutput) ToPartnerRegistrationOutputWithContext(ctx context.Context) PartnerRegistrationOutput {
	return o
}

func (o PartnerRegistrationOutput) AuthorizedAzureSubscriptionIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PartnerRegistration) pulumi.StringArrayOutput { return v.AuthorizedAzureSubscriptionIds }).(pulumi.StringArrayOutput)
}

func (o PartnerRegistrationOutput) CustomerServiceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerRegistration) pulumi.StringPtrOutput { return v.CustomerServiceUri }).(pulumi.StringPtrOutput)
}

func (o PartnerRegistrationOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *PartnerRegistration) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o PartnerRegistrationOutput) LogoUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerRegistration) pulumi.StringPtrOutput { return v.LogoUri }).(pulumi.StringPtrOutput)
}

func (o PartnerRegistrationOutput) LongDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerRegistration) pulumi.StringPtrOutput { return v.LongDescription }).(pulumi.StringPtrOutput)
}

func (o PartnerRegistrationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PartnerRegistration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PartnerRegistrationOutput) PartnerCustomerServiceExtension() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerRegistration) pulumi.StringPtrOutput { return v.PartnerCustomerServiceExtension }).(pulumi.StringPtrOutput)
}

func (o PartnerRegistrationOutput) PartnerCustomerServiceNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerRegistration) pulumi.StringPtrOutput { return v.PartnerCustomerServiceNumber }).(pulumi.StringPtrOutput)
}

func (o PartnerRegistrationOutput) PartnerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerRegistration) pulumi.StringPtrOutput { return v.PartnerName }).(pulumi.StringPtrOutput)
}

func (o PartnerRegistrationOutput) PartnerResourceTypeDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerRegistration) pulumi.StringPtrOutput { return v.PartnerResourceTypeDescription }).(pulumi.StringPtrOutput)
}

func (o PartnerRegistrationOutput) PartnerResourceTypeDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerRegistration) pulumi.StringPtrOutput { return v.PartnerResourceTypeDisplayName }).(pulumi.StringPtrOutput)
}

func (o PartnerRegistrationOutput) PartnerResourceTypeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerRegistration) pulumi.StringPtrOutput { return v.PartnerResourceTypeName }).(pulumi.StringPtrOutput)
}

func (o PartnerRegistrationOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PartnerRegistration) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PartnerRegistrationOutput) SetupUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerRegistration) pulumi.StringPtrOutput { return v.SetupUri }).(pulumi.StringPtrOutput)
}

func (o PartnerRegistrationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PartnerRegistration) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o PartnerRegistrationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PartnerRegistration) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o PartnerRegistrationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PartnerRegistration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o PartnerRegistrationOutput) VisibilityState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerRegistration) pulumi.StringPtrOutput { return v.VisibilityState }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(PartnerRegistrationOutput{})
}
