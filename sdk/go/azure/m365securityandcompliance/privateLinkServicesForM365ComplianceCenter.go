


package m365securityandcompliance

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateLinkServicesForM365ComplianceCenter struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringPtrOutput                    `pulumi:"etag"`
	Identity   ServicesResourceResponseIdentityPtrOutput `pulumi:"identity"`
	Kind       pulumi.StringOutput                       `pulumi:"kind"`
	Location   pulumi.StringOutput                       `pulumi:"location"`
	Name       pulumi.StringOutput                       `pulumi:"name"`
	Properties ServicesPropertiesResponseOutput          `pulumi:"properties"`
	SystemData SystemDataResponseOutput                  `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput                    `pulumi:"tags"`
	Type       pulumi.StringOutput                       `pulumi:"type"`
}


func NewPrivateLinkServicesForM365ComplianceCenter(ctx *pulumi.Context,
	name string, args *PrivateLinkServicesForM365ComplianceCenterArgs, opts ...pulumi.ResourceOption) (*PrivateLinkServicesForM365ComplianceCenter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:m365securityandcompliance/v20210325preview:privateLinkServicesForM365ComplianceCenter"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateLinkServicesForM365ComplianceCenter
	err := ctx.RegisterResource("azure-native:m365securityandcompliance:privateLinkServicesForM365ComplianceCenter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateLinkServicesForM365ComplianceCenter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateLinkServicesForM365ComplianceCenterState, opts ...pulumi.ResourceOption) (*PrivateLinkServicesForM365ComplianceCenter, error) {
	var resource PrivateLinkServicesForM365ComplianceCenter
	err := ctx.ReadResource("azure-native:m365securityandcompliance:privateLinkServicesForM365ComplianceCenter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type privateLinkServicesForM365ComplianceCenterState struct {
}

type PrivateLinkServicesForM365ComplianceCenterState struct {
}

func (PrivateLinkServicesForM365ComplianceCenterState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkServicesForM365ComplianceCenterState)(nil)).Elem()
}

type privateLinkServicesForM365ComplianceCenterArgs struct {
	Identity          *ServicesResourceIdentity `pulumi:"identity"`
	Kind              Kind                      `pulumi:"kind"`
	Location          *string                   `pulumi:"location"`
	Properties        *ServicesProperties       `pulumi:"properties"`
	ResourceGroupName string                    `pulumi:"resourceGroupName"`
	ResourceName      *string                   `pulumi:"resourceName"`
	Tags              map[string]string         `pulumi:"tags"`
}


type PrivateLinkServicesForM365ComplianceCenterArgs struct {
	Identity          ServicesResourceIdentityPtrInput
	Kind              KindInput
	Location          pulumi.StringPtrInput
	Properties        ServicesPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
}

func (PrivateLinkServicesForM365ComplianceCenterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkServicesForM365ComplianceCenterArgs)(nil)).Elem()
}

type PrivateLinkServicesForM365ComplianceCenterInput interface {
	pulumi.Input

	ToPrivateLinkServicesForM365ComplianceCenterOutput() PrivateLinkServicesForM365ComplianceCenterOutput
	ToPrivateLinkServicesForM365ComplianceCenterOutputWithContext(ctx context.Context) PrivateLinkServicesForM365ComplianceCenterOutput
}

func (*PrivateLinkServicesForM365ComplianceCenter) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServicesForM365ComplianceCenter)(nil)).Elem()
}

func (i *PrivateLinkServicesForM365ComplianceCenter) ToPrivateLinkServicesForM365ComplianceCenterOutput() PrivateLinkServicesForM365ComplianceCenterOutput {
	return i.ToPrivateLinkServicesForM365ComplianceCenterOutputWithContext(context.Background())
}

func (i *PrivateLinkServicesForM365ComplianceCenter) ToPrivateLinkServicesForM365ComplianceCenterOutputWithContext(ctx context.Context) PrivateLinkServicesForM365ComplianceCenterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServicesForM365ComplianceCenterOutput)
}

type PrivateLinkServicesForM365ComplianceCenterOutput struct{ *pulumi.OutputState }

func (PrivateLinkServicesForM365ComplianceCenterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServicesForM365ComplianceCenter)(nil)).Elem()
}

func (o PrivateLinkServicesForM365ComplianceCenterOutput) ToPrivateLinkServicesForM365ComplianceCenterOutput() PrivateLinkServicesForM365ComplianceCenterOutput {
	return o
}

func (o PrivateLinkServicesForM365ComplianceCenterOutput) ToPrivateLinkServicesForM365ComplianceCenterOutputWithContext(ctx context.Context) PrivateLinkServicesForM365ComplianceCenterOutput {
	return o
}

func (o PrivateLinkServicesForM365ComplianceCenterOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForM365ComplianceCenter) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServicesForM365ComplianceCenterOutput) Identity() ServicesResourceResponseIdentityPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForM365ComplianceCenter) ServicesResourceResponseIdentityPtrOutput {
		return v.Identity
	}).(ServicesResourceResponseIdentityPtrOutput)
}

func (o PrivateLinkServicesForM365ComplianceCenterOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForM365ComplianceCenter) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o PrivateLinkServicesForM365ComplianceCenterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForM365ComplianceCenter) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o PrivateLinkServicesForM365ComplianceCenterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForM365ComplianceCenter) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateLinkServicesForM365ComplianceCenterOutput) Properties() ServicesPropertiesResponseOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForM365ComplianceCenter) ServicesPropertiesResponseOutput {
		return v.Properties
	}).(ServicesPropertiesResponseOutput)
}

func (o PrivateLinkServicesForM365ComplianceCenterOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForM365ComplianceCenter) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o PrivateLinkServicesForM365ComplianceCenterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForM365ComplianceCenter) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o PrivateLinkServicesForM365ComplianceCenterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForM365ComplianceCenter) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateLinkServicesForM365ComplianceCenterOutput{})
}
