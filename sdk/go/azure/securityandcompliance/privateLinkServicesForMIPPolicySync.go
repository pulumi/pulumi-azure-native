


package securityandcompliance

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateLinkServicesForMIPPolicySync struct {
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


func NewPrivateLinkServicesForMIPPolicySync(ctx *pulumi.Context,
	name string, args *PrivateLinkServicesForMIPPolicySyncArgs, opts ...pulumi.ResourceOption) (*PrivateLinkServicesForMIPPolicySync, error) {
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
			Type: pulumi.String("azure-native:securityandcompliance/v20210308:privateLinkServicesForMIPPolicySync"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateLinkServicesForMIPPolicySync
	err := ctx.RegisterResource("azure-native:securityandcompliance:privateLinkServicesForMIPPolicySync", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateLinkServicesForMIPPolicySync(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateLinkServicesForMIPPolicySyncState, opts ...pulumi.ResourceOption) (*PrivateLinkServicesForMIPPolicySync, error) {
	var resource PrivateLinkServicesForMIPPolicySync
	err := ctx.ReadResource("azure-native:securityandcompliance:privateLinkServicesForMIPPolicySync", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type privateLinkServicesForMIPPolicySyncState struct {
}

type PrivateLinkServicesForMIPPolicySyncState struct {
}

func (PrivateLinkServicesForMIPPolicySyncState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkServicesForMIPPolicySyncState)(nil)).Elem()
}

type privateLinkServicesForMIPPolicySyncArgs struct {
	Identity          *ServicesResourceIdentity `pulumi:"identity"`
	Kind              Kind                      `pulumi:"kind"`
	Location          *string                   `pulumi:"location"`
	Properties        *ServicesProperties       `pulumi:"properties"`
	ResourceGroupName string                    `pulumi:"resourceGroupName"`
	ResourceName      *string                   `pulumi:"resourceName"`
	Tags              map[string]string         `pulumi:"tags"`
}


type PrivateLinkServicesForMIPPolicySyncArgs struct {
	Identity          ServicesResourceIdentityPtrInput
	Kind              KindInput
	Location          pulumi.StringPtrInput
	Properties        ServicesPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
}

func (PrivateLinkServicesForMIPPolicySyncArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkServicesForMIPPolicySyncArgs)(nil)).Elem()
}

type PrivateLinkServicesForMIPPolicySyncInput interface {
	pulumi.Input

	ToPrivateLinkServicesForMIPPolicySyncOutput() PrivateLinkServicesForMIPPolicySyncOutput
	ToPrivateLinkServicesForMIPPolicySyncOutputWithContext(ctx context.Context) PrivateLinkServicesForMIPPolicySyncOutput
}

func (*PrivateLinkServicesForMIPPolicySync) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServicesForMIPPolicySync)(nil)).Elem()
}

func (i *PrivateLinkServicesForMIPPolicySync) ToPrivateLinkServicesForMIPPolicySyncOutput() PrivateLinkServicesForMIPPolicySyncOutput {
	return i.ToPrivateLinkServicesForMIPPolicySyncOutputWithContext(context.Background())
}

func (i *PrivateLinkServicesForMIPPolicySync) ToPrivateLinkServicesForMIPPolicySyncOutputWithContext(ctx context.Context) PrivateLinkServicesForMIPPolicySyncOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServicesForMIPPolicySyncOutput)
}

type PrivateLinkServicesForMIPPolicySyncOutput struct{ *pulumi.OutputState }

func (PrivateLinkServicesForMIPPolicySyncOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServicesForMIPPolicySync)(nil)).Elem()
}

func (o PrivateLinkServicesForMIPPolicySyncOutput) ToPrivateLinkServicesForMIPPolicySyncOutput() PrivateLinkServicesForMIPPolicySyncOutput {
	return o
}

func (o PrivateLinkServicesForMIPPolicySyncOutput) ToPrivateLinkServicesForMIPPolicySyncOutputWithContext(ctx context.Context) PrivateLinkServicesForMIPPolicySyncOutput {
	return o
}

func (o PrivateLinkServicesForMIPPolicySyncOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForMIPPolicySync) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServicesForMIPPolicySyncOutput) Identity() ServicesResourceResponseIdentityPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForMIPPolicySync) ServicesResourceResponseIdentityPtrOutput {
		return v.Identity
	}).(ServicesResourceResponseIdentityPtrOutput)
}

func (o PrivateLinkServicesForMIPPolicySyncOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForMIPPolicySync) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o PrivateLinkServicesForMIPPolicySyncOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForMIPPolicySync) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o PrivateLinkServicesForMIPPolicySyncOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForMIPPolicySync) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateLinkServicesForMIPPolicySyncOutput) Properties() ServicesPropertiesResponseOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForMIPPolicySync) ServicesPropertiesResponseOutput { return v.Properties }).(ServicesPropertiesResponseOutput)
}

func (o PrivateLinkServicesForMIPPolicySyncOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForMIPPolicySync) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o PrivateLinkServicesForMIPPolicySyncOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForMIPPolicySync) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o PrivateLinkServicesForMIPPolicySyncOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForMIPPolicySync) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateLinkServicesForMIPPolicySyncOutput{})
}
