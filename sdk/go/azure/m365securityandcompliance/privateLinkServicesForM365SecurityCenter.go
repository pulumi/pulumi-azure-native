


package m365securityandcompliance

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateLinkServicesForM365SecurityCenter struct {
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


func NewPrivateLinkServicesForM365SecurityCenter(ctx *pulumi.Context,
	name string, args *PrivateLinkServicesForM365SecurityCenterArgs, opts ...pulumi.ResourceOption) (*PrivateLinkServicesForM365SecurityCenter, error) {
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
			Type: pulumi.String("azure-native:m365securityandcompliance/v20210325preview:privateLinkServicesForM365SecurityCenter"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateLinkServicesForM365SecurityCenter
	err := ctx.RegisterResource("azure-native:m365securityandcompliance:privateLinkServicesForM365SecurityCenter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateLinkServicesForM365SecurityCenter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateLinkServicesForM365SecurityCenterState, opts ...pulumi.ResourceOption) (*PrivateLinkServicesForM365SecurityCenter, error) {
	var resource PrivateLinkServicesForM365SecurityCenter
	err := ctx.ReadResource("azure-native:m365securityandcompliance:privateLinkServicesForM365SecurityCenter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type privateLinkServicesForM365SecurityCenterState struct {
}

type PrivateLinkServicesForM365SecurityCenterState struct {
}

func (PrivateLinkServicesForM365SecurityCenterState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkServicesForM365SecurityCenterState)(nil)).Elem()
}

type privateLinkServicesForM365SecurityCenterArgs struct {
	Identity          *ServicesResourceIdentity `pulumi:"identity"`
	Kind              Kind                      `pulumi:"kind"`
	Location          *string                   `pulumi:"location"`
	Properties        *ServicesProperties       `pulumi:"properties"`
	ResourceGroupName string                    `pulumi:"resourceGroupName"`
	ResourceName      *string                   `pulumi:"resourceName"`
	Tags              map[string]string         `pulumi:"tags"`
}


type PrivateLinkServicesForM365SecurityCenterArgs struct {
	Identity          ServicesResourceIdentityPtrInput
	Kind              KindInput
	Location          pulumi.StringPtrInput
	Properties        ServicesPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
}

func (PrivateLinkServicesForM365SecurityCenterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkServicesForM365SecurityCenterArgs)(nil)).Elem()
}

type PrivateLinkServicesForM365SecurityCenterInput interface {
	pulumi.Input

	ToPrivateLinkServicesForM365SecurityCenterOutput() PrivateLinkServicesForM365SecurityCenterOutput
	ToPrivateLinkServicesForM365SecurityCenterOutputWithContext(ctx context.Context) PrivateLinkServicesForM365SecurityCenterOutput
}

func (*PrivateLinkServicesForM365SecurityCenter) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServicesForM365SecurityCenter)(nil)).Elem()
}

func (i *PrivateLinkServicesForM365SecurityCenter) ToPrivateLinkServicesForM365SecurityCenterOutput() PrivateLinkServicesForM365SecurityCenterOutput {
	return i.ToPrivateLinkServicesForM365SecurityCenterOutputWithContext(context.Background())
}

func (i *PrivateLinkServicesForM365SecurityCenter) ToPrivateLinkServicesForM365SecurityCenterOutputWithContext(ctx context.Context) PrivateLinkServicesForM365SecurityCenterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServicesForM365SecurityCenterOutput)
}

type PrivateLinkServicesForM365SecurityCenterOutput struct{ *pulumi.OutputState }

func (PrivateLinkServicesForM365SecurityCenterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServicesForM365SecurityCenter)(nil)).Elem()
}

func (o PrivateLinkServicesForM365SecurityCenterOutput) ToPrivateLinkServicesForM365SecurityCenterOutput() PrivateLinkServicesForM365SecurityCenterOutput {
	return o
}

func (o PrivateLinkServicesForM365SecurityCenterOutput) ToPrivateLinkServicesForM365SecurityCenterOutputWithContext(ctx context.Context) PrivateLinkServicesForM365SecurityCenterOutput {
	return o
}

func (o PrivateLinkServicesForM365SecurityCenterOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForM365SecurityCenter) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServicesForM365SecurityCenterOutput) Identity() ServicesResourceResponseIdentityPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForM365SecurityCenter) ServicesResourceResponseIdentityPtrOutput {
		return v.Identity
	}).(ServicesResourceResponseIdentityPtrOutput)
}

func (o PrivateLinkServicesForM365SecurityCenterOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForM365SecurityCenter) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o PrivateLinkServicesForM365SecurityCenterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForM365SecurityCenter) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o PrivateLinkServicesForM365SecurityCenterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForM365SecurityCenter) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateLinkServicesForM365SecurityCenterOutput) Properties() ServicesPropertiesResponseOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForM365SecurityCenter) ServicesPropertiesResponseOutput {
		return v.Properties
	}).(ServicesPropertiesResponseOutput)
}

func (o PrivateLinkServicesForM365SecurityCenterOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForM365SecurityCenter) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o PrivateLinkServicesForM365SecurityCenterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForM365SecurityCenter) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o PrivateLinkServicesForM365SecurityCenterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForM365SecurityCenter) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateLinkServicesForM365SecurityCenterOutput{})
}
