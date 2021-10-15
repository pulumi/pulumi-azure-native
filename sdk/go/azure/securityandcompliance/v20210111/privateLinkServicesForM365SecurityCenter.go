


package v20210111

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
			Type: pulumi.String("azure-nextgen:securityandcompliance/v20210111:privateLinkServicesForM365SecurityCenter"),
		},
		{
			Type: pulumi.String("azure-native:securityandcompliance:privateLinkServicesForM365SecurityCenter"),
		},
		{
			Type: pulumi.String("azure-nextgen:securityandcompliance:privateLinkServicesForM365SecurityCenter"),
		},
		{
			Type: pulumi.String("azure-native:securityandcompliance/v20210308:privateLinkServicesForM365SecurityCenter"),
		},
		{
			Type: pulumi.String("azure-nextgen:securityandcompliance/v20210308:privateLinkServicesForM365SecurityCenter"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateLinkServicesForM365SecurityCenter
	err := ctx.RegisterResource("azure-native:securityandcompliance/v20210111:privateLinkServicesForM365SecurityCenter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateLinkServicesForM365SecurityCenter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateLinkServicesForM365SecurityCenterState, opts ...pulumi.ResourceOption) (*PrivateLinkServicesForM365SecurityCenter, error) {
	var resource PrivateLinkServicesForM365SecurityCenter
	err := ctx.ReadResource("azure-native:securityandcompliance/v20210111:privateLinkServicesForM365SecurityCenter", name, id, state, &resource, opts...)
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
	Etag              *string                   `pulumi:"etag"`
	Identity          *ServicesResourceIdentity `pulumi:"identity"`
	Kind              Kind                      `pulumi:"kind"`
	Location          *string                   `pulumi:"location"`
	Properties        *ServicesProperties       `pulumi:"properties"`
	ResourceGroupName string                    `pulumi:"resourceGroupName"`
	ResourceName      *string                   `pulumi:"resourceName"`
	Tags              map[string]string         `pulumi:"tags"`
}


type PrivateLinkServicesForM365SecurityCenterArgs struct {
	Etag              pulumi.StringPtrInput
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
	return reflect.TypeOf((*PrivateLinkServicesForM365SecurityCenter)(nil))
}

func (i *PrivateLinkServicesForM365SecurityCenter) ToPrivateLinkServicesForM365SecurityCenterOutput() PrivateLinkServicesForM365SecurityCenterOutput {
	return i.ToPrivateLinkServicesForM365SecurityCenterOutputWithContext(context.Background())
}

func (i *PrivateLinkServicesForM365SecurityCenter) ToPrivateLinkServicesForM365SecurityCenterOutputWithContext(ctx context.Context) PrivateLinkServicesForM365SecurityCenterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServicesForM365SecurityCenterOutput)
}

type PrivateLinkServicesForM365SecurityCenterOutput struct{ *pulumi.OutputState }

func (PrivateLinkServicesForM365SecurityCenterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServicesForM365SecurityCenter)(nil))
}

func (o PrivateLinkServicesForM365SecurityCenterOutput) ToPrivateLinkServicesForM365SecurityCenterOutput() PrivateLinkServicesForM365SecurityCenterOutput {
	return o
}

func (o PrivateLinkServicesForM365SecurityCenterOutput) ToPrivateLinkServicesForM365SecurityCenterOutputWithContext(ctx context.Context) PrivateLinkServicesForM365SecurityCenterOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PrivateLinkServicesForM365SecurityCenterOutput{})
}
