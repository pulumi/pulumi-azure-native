


package v20220201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SecurityUserConfiguration struct {
	pulumi.CustomResourceState

	DeleteExistingNSGs pulumi.StringPtrOutput   `pulumi:"deleteExistingNSGs"`
	Description        pulumi.StringPtrOutput   `pulumi:"description"`
	Etag               pulumi.StringOutput      `pulumi:"etag"`
	Name               pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState  pulumi.StringOutput      `pulumi:"provisioningState"`
	SystemData         SystemDataResponseOutput `pulumi:"systemData"`
	Type               pulumi.StringOutput      `pulumi:"type"`
}


func NewSecurityUserConfiguration(ctx *pulumi.Context,
	name string, args *SecurityUserConfigurationArgs, opts ...pulumi.ResourceOption) (*SecurityUserConfiguration, error) {
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
			Type: pulumi.String("azure-native:network:SecurityUserConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201preview:SecurityUserConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501preview:SecurityUserConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220401preview:SecurityUserConfiguration"),
		},
	})
	opts = append(opts, aliases)
	var resource SecurityUserConfiguration
	err := ctx.RegisterResource("azure-native:network/v20220201preview:SecurityUserConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSecurityUserConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityUserConfigurationState, opts ...pulumi.ResourceOption) (*SecurityUserConfiguration, error) {
	var resource SecurityUserConfiguration
	err := ctx.ReadResource("azure-native:network/v20220201preview:SecurityUserConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type securityUserConfigurationState struct {
}

type SecurityUserConfigurationState struct {
}

func (SecurityUserConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityUserConfigurationState)(nil)).Elem()
}

type securityUserConfigurationArgs struct {
	ConfigurationName  *string `pulumi:"configurationName"`
	DeleteExistingNSGs *string `pulumi:"deleteExistingNSGs"`
	Description        *string `pulumi:"description"`
	NetworkManagerName string  `pulumi:"networkManagerName"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
}


type SecurityUserConfigurationArgs struct {
	ConfigurationName  pulumi.StringPtrInput
	DeleteExistingNSGs pulumi.StringPtrInput
	Description        pulumi.StringPtrInput
	NetworkManagerName pulumi.StringInput
	ResourceGroupName  pulumi.StringInput
}

func (SecurityUserConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityUserConfigurationArgs)(nil)).Elem()
}

type SecurityUserConfigurationInput interface {
	pulumi.Input

	ToSecurityUserConfigurationOutput() SecurityUserConfigurationOutput
	ToSecurityUserConfigurationOutputWithContext(ctx context.Context) SecurityUserConfigurationOutput
}

func (*SecurityUserConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityUserConfiguration)(nil)).Elem()
}

func (i *SecurityUserConfiguration) ToSecurityUserConfigurationOutput() SecurityUserConfigurationOutput {
	return i.ToSecurityUserConfigurationOutputWithContext(context.Background())
}

func (i *SecurityUserConfiguration) ToSecurityUserConfigurationOutputWithContext(ctx context.Context) SecurityUserConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityUserConfigurationOutput)
}

type SecurityUserConfigurationOutput struct{ *pulumi.OutputState }

func (SecurityUserConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityUserConfiguration)(nil)).Elem()
}

func (o SecurityUserConfigurationOutput) ToSecurityUserConfigurationOutput() SecurityUserConfigurationOutput {
	return o
}

func (o SecurityUserConfigurationOutput) ToSecurityUserConfigurationOutputWithContext(ctx context.Context) SecurityUserConfigurationOutput {
	return o
}

func (o SecurityUserConfigurationOutput) DeleteExistingNSGs() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityUserConfiguration) pulumi.StringPtrOutput { return v.DeleteExistingNSGs }).(pulumi.StringPtrOutput)
}

func (o SecurityUserConfigurationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityUserConfiguration) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SecurityUserConfigurationOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityUserConfiguration) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o SecurityUserConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityUserConfiguration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SecurityUserConfigurationOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityUserConfiguration) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o SecurityUserConfigurationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SecurityUserConfiguration) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SecurityUserConfigurationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityUserConfiguration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SecurityUserConfigurationOutput{})
}
