


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SecurityAdminConfiguration struct {
	pulumi.CustomResourceState

	DeleteExistingNSGs pulumi.StringPtrOutput   `pulumi:"deleteExistingNSGs"`
	Description        pulumi.StringPtrOutput   `pulumi:"description"`
	DisplayName        pulumi.StringPtrOutput   `pulumi:"displayName"`
	Etag               pulumi.StringOutput      `pulumi:"etag"`
	Name               pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState  pulumi.StringOutput      `pulumi:"provisioningState"`
	SecurityType       pulumi.StringPtrOutput   `pulumi:"securityType"`
	SystemData         SystemDataResponseOutput `pulumi:"systemData"`
	Type               pulumi.StringOutput      `pulumi:"type"`
}


func NewSecurityAdminConfiguration(ctx *pulumi.Context,
	name string, args *SecurityAdminConfigurationArgs, opts ...pulumi.ResourceOption) (*SecurityAdminConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkManagerName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkManagerName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:SecurityAdminConfiguration"),
		},
	})
	opts = append(opts, aliases)
	var resource SecurityAdminConfiguration
	err := ctx.RegisterResource("azure-native:network/v20210201preview:SecurityAdminConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSecurityAdminConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityAdminConfigurationState, opts ...pulumi.ResourceOption) (*SecurityAdminConfiguration, error) {
	var resource SecurityAdminConfiguration
	err := ctx.ReadResource("azure-native:network/v20210201preview:SecurityAdminConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type securityAdminConfigurationState struct {
}

type SecurityAdminConfigurationState struct {
}

func (SecurityAdminConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityAdminConfigurationState)(nil)).Elem()
}

type securityAdminConfigurationArgs struct {
	ConfigurationName  *string `pulumi:"configurationName"`
	DeleteExistingNSGs *string `pulumi:"deleteExistingNSGs"`
	Description        *string `pulumi:"description"`
	DisplayName        *string `pulumi:"displayName"`
	NetworkManagerName string  `pulumi:"networkManagerName"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
	SecurityType       *string `pulumi:"securityType"`
}


type SecurityAdminConfigurationArgs struct {
	ConfigurationName  pulumi.StringPtrInput
	DeleteExistingNSGs pulumi.StringPtrInput
	Description        pulumi.StringPtrInput
	DisplayName        pulumi.StringPtrInput
	NetworkManagerName pulumi.StringInput
	ResourceGroupName  pulumi.StringInput
	SecurityType       pulumi.StringPtrInput
}

func (SecurityAdminConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityAdminConfigurationArgs)(nil)).Elem()
}

type SecurityAdminConfigurationInput interface {
	pulumi.Input

	ToSecurityAdminConfigurationOutput() SecurityAdminConfigurationOutput
	ToSecurityAdminConfigurationOutputWithContext(ctx context.Context) SecurityAdminConfigurationOutput
}

func (*SecurityAdminConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityAdminConfiguration)(nil)).Elem()
}

func (i *SecurityAdminConfiguration) ToSecurityAdminConfigurationOutput() SecurityAdminConfigurationOutput {
	return i.ToSecurityAdminConfigurationOutputWithContext(context.Background())
}

func (i *SecurityAdminConfiguration) ToSecurityAdminConfigurationOutputWithContext(ctx context.Context) SecurityAdminConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityAdminConfigurationOutput)
}

type SecurityAdminConfigurationOutput struct{ *pulumi.OutputState }

func (SecurityAdminConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityAdminConfiguration)(nil)).Elem()
}

func (o SecurityAdminConfigurationOutput) ToSecurityAdminConfigurationOutput() SecurityAdminConfigurationOutput {
	return o
}

func (o SecurityAdminConfigurationOutput) ToSecurityAdminConfigurationOutputWithContext(ctx context.Context) SecurityAdminConfigurationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SecurityAdminConfigurationOutput{})
}
