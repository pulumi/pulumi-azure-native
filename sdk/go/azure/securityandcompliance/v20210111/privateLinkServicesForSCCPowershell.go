


package v20210111

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateLinkServicesForSCCPowershell struct {
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


func NewPrivateLinkServicesForSCCPowershell(ctx *pulumi.Context,
	name string, args *PrivateLinkServicesForSCCPowershellArgs, opts ...pulumi.ResourceOption) (*PrivateLinkServicesForSCCPowershell, error) {
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
			Type: pulumi.String("azure-native:securityandcompliance:privateLinkServicesForSCCPowershell"),
		},
		{
			Type: pulumi.String("azure-native:securityandcompliance/v20210308:privateLinkServicesForSCCPowershell"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateLinkServicesForSCCPowershell
	err := ctx.RegisterResource("azure-native:securityandcompliance/v20210111:privateLinkServicesForSCCPowershell", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateLinkServicesForSCCPowershell(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateLinkServicesForSCCPowershellState, opts ...pulumi.ResourceOption) (*PrivateLinkServicesForSCCPowershell, error) {
	var resource PrivateLinkServicesForSCCPowershell
	err := ctx.ReadResource("azure-native:securityandcompliance/v20210111:privateLinkServicesForSCCPowershell", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type privateLinkServicesForSCCPowershellState struct {
}

type PrivateLinkServicesForSCCPowershellState struct {
}

func (PrivateLinkServicesForSCCPowershellState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkServicesForSCCPowershellState)(nil)).Elem()
}

type privateLinkServicesForSCCPowershellArgs struct {
	Etag              *string                   `pulumi:"etag"`
	Identity          *ServicesResourceIdentity `pulumi:"identity"`
	Kind              Kind                      `pulumi:"kind"`
	Location          *string                   `pulumi:"location"`
	Properties        *ServicesProperties       `pulumi:"properties"`
	ResourceGroupName string                    `pulumi:"resourceGroupName"`
	ResourceName      *string                   `pulumi:"resourceName"`
	Tags              map[string]string         `pulumi:"tags"`
}


type PrivateLinkServicesForSCCPowershellArgs struct {
	Etag              pulumi.StringPtrInput
	Identity          ServicesResourceIdentityPtrInput
	Kind              KindInput
	Location          pulumi.StringPtrInput
	Properties        ServicesPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
}

func (PrivateLinkServicesForSCCPowershellArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkServicesForSCCPowershellArgs)(nil)).Elem()
}

type PrivateLinkServicesForSCCPowershellInput interface {
	pulumi.Input

	ToPrivateLinkServicesForSCCPowershellOutput() PrivateLinkServicesForSCCPowershellOutput
	ToPrivateLinkServicesForSCCPowershellOutputWithContext(ctx context.Context) PrivateLinkServicesForSCCPowershellOutput
}

func (*PrivateLinkServicesForSCCPowershell) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServicesForSCCPowershell)(nil)).Elem()
}

func (i *PrivateLinkServicesForSCCPowershell) ToPrivateLinkServicesForSCCPowershellOutput() PrivateLinkServicesForSCCPowershellOutput {
	return i.ToPrivateLinkServicesForSCCPowershellOutputWithContext(context.Background())
}

func (i *PrivateLinkServicesForSCCPowershell) ToPrivateLinkServicesForSCCPowershellOutputWithContext(ctx context.Context) PrivateLinkServicesForSCCPowershellOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServicesForSCCPowershellOutput)
}

type PrivateLinkServicesForSCCPowershellOutput struct{ *pulumi.OutputState }

func (PrivateLinkServicesForSCCPowershellOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServicesForSCCPowershell)(nil)).Elem()
}

func (o PrivateLinkServicesForSCCPowershellOutput) ToPrivateLinkServicesForSCCPowershellOutput() PrivateLinkServicesForSCCPowershellOutput {
	return o
}

func (o PrivateLinkServicesForSCCPowershellOutput) ToPrivateLinkServicesForSCCPowershellOutputWithContext(ctx context.Context) PrivateLinkServicesForSCCPowershellOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PrivateLinkServicesForSCCPowershellOutput{})
}
