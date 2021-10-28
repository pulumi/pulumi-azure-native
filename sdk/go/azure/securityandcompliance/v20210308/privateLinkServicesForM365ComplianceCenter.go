


package v20210308

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
			Type: pulumi.String("azure-nextgen:securityandcompliance/v20210308:privateLinkServicesForM365ComplianceCenter"),
		},
		{
			Type: pulumi.String("azure-native:securityandcompliance:privateLinkServicesForM365ComplianceCenter"),
		},
		{
			Type: pulumi.String("azure-nextgen:securityandcompliance:privateLinkServicesForM365ComplianceCenter"),
		},
		{
			Type: pulumi.String("azure-native:securityandcompliance/v20210111:privateLinkServicesForM365ComplianceCenter"),
		},
		{
			Type: pulumi.String("azure-nextgen:securityandcompliance/v20210111:privateLinkServicesForM365ComplianceCenter"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateLinkServicesForM365ComplianceCenter
	err := ctx.RegisterResource("azure-native:securityandcompliance/v20210308:privateLinkServicesForM365ComplianceCenter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateLinkServicesForM365ComplianceCenter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateLinkServicesForM365ComplianceCenterState, opts ...pulumi.ResourceOption) (*PrivateLinkServicesForM365ComplianceCenter, error) {
	var resource PrivateLinkServicesForM365ComplianceCenter
	err := ctx.ReadResource("azure-native:securityandcompliance/v20210308:privateLinkServicesForM365ComplianceCenter", name, id, state, &resource, opts...)
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
	Etag              *string                   `pulumi:"etag"`
	Identity          *ServicesResourceIdentity `pulumi:"identity"`
	Kind              Kind                      `pulumi:"kind"`
	Location          *string                   `pulumi:"location"`
	Properties        *ServicesProperties       `pulumi:"properties"`
	ResourceGroupName string                    `pulumi:"resourceGroupName"`
	ResourceName      *string                   `pulumi:"resourceName"`
	Tags              map[string]string         `pulumi:"tags"`
}


type PrivateLinkServicesForM365ComplianceCenterArgs struct {
	Etag              pulumi.StringPtrInput
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
	return reflect.TypeOf((*PrivateLinkServicesForM365ComplianceCenter)(nil))
}

func (i *PrivateLinkServicesForM365ComplianceCenter) ToPrivateLinkServicesForM365ComplianceCenterOutput() PrivateLinkServicesForM365ComplianceCenterOutput {
	return i.ToPrivateLinkServicesForM365ComplianceCenterOutputWithContext(context.Background())
}

func (i *PrivateLinkServicesForM365ComplianceCenter) ToPrivateLinkServicesForM365ComplianceCenterOutputWithContext(ctx context.Context) PrivateLinkServicesForM365ComplianceCenterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServicesForM365ComplianceCenterOutput)
}

type PrivateLinkServicesForM365ComplianceCenterOutput struct{ *pulumi.OutputState }

func (PrivateLinkServicesForM365ComplianceCenterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServicesForM365ComplianceCenter)(nil))
}

func (o PrivateLinkServicesForM365ComplianceCenterOutput) ToPrivateLinkServicesForM365ComplianceCenterOutput() PrivateLinkServicesForM365ComplianceCenterOutput {
	return o
}

func (o PrivateLinkServicesForM365ComplianceCenterOutput) ToPrivateLinkServicesForM365ComplianceCenterOutputWithContext(ctx context.Context) PrivateLinkServicesForM365ComplianceCenterOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServicesForM365ComplianceCenterInput)(nil)).Elem(), &PrivateLinkServicesForM365ComplianceCenter{})
	pulumi.RegisterOutputType(PrivateLinkServicesForM365ComplianceCenterOutput{})
}
