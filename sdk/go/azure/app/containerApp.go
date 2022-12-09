


package app

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ContainerApp struct {
	pulumi.CustomResourceState

	Configuration              ConfigurationResponsePtrOutput          `pulumi:"configuration"`
	CustomDomainVerificationId pulumi.StringOutput                     `pulumi:"customDomainVerificationId"`
	Identity                   ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	LatestRevisionFqdn         pulumi.StringOutput                     `pulumi:"latestRevisionFqdn"`
	LatestRevisionName         pulumi.StringOutput                     `pulumi:"latestRevisionName"`
	Location                   pulumi.StringOutput                     `pulumi:"location"`
	ManagedEnvironmentId       pulumi.StringPtrOutput                  `pulumi:"managedEnvironmentId"`
	Name                       pulumi.StringOutput                     `pulumi:"name"`
	OutboundIpAddresses        pulumi.StringArrayOutput                `pulumi:"outboundIpAddresses"`
	ProvisioningState          pulumi.StringOutput                     `pulumi:"provisioningState"`
	SystemData                 SystemDataResponseOutput                `pulumi:"systemData"`
	Tags                       pulumi.StringMapOutput                  `pulumi:"tags"`
	Template                   TemplateResponsePtrOutput               `pulumi:"template"`
	Type                       pulumi.StringOutput                     `pulumi:"type"`
}


func NewContainerApp(ctx *pulumi.Context,
	name string, args *ContainerAppArgs, opts ...pulumi.ResourceOption) (*ContainerApp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Configuration != nil {
		args.Configuration = args.Configuration.ToConfigurationPtrOutput().ApplyT(func(v *Configuration) *Configuration { return v.Defaults() }).(ConfigurationPtrOutput)
	}
	if args.Template != nil {
		args.Template = args.Template.ToTemplatePtrOutput().ApplyT(func(v *Template) *Template { return v.Defaults() }).(TemplatePtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:app/v20220101preview:ContainerApp"),
		},
		{
			Type: pulumi.String("azure-native:app/v20220301:ContainerApp"),
		},
		{
			Type: pulumi.String("azure-native:app/v20220601preview:ContainerApp"),
		},
		{
			Type: pulumi.String("azure-native:app/v20221001:ContainerApp"),
		},
	})
	opts = append(opts, aliases)
	var resource ContainerApp
	err := ctx.RegisterResource("azure-native:app:ContainerApp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetContainerApp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerAppState, opts ...pulumi.ResourceOption) (*ContainerApp, error) {
	var resource ContainerApp
	err := ctx.ReadResource("azure-native:app:ContainerApp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type containerAppState struct {
}

type ContainerAppState struct {
}

func (ContainerAppState) ElementType() reflect.Type {
	return reflect.TypeOf((*containerAppState)(nil)).Elem()
}

type containerAppArgs struct {
	Configuration        *Configuration          `pulumi:"configuration"`
	ContainerAppName     *string                 `pulumi:"containerAppName"`
	Identity             *ManagedServiceIdentity `pulumi:"identity"`
	Location             *string                 `pulumi:"location"`
	ManagedEnvironmentId *string                 `pulumi:"managedEnvironmentId"`
	ResourceGroupName    string                  `pulumi:"resourceGroupName"`
	Tags                 map[string]string       `pulumi:"tags"`
	Template             *Template               `pulumi:"template"`
}


type ContainerAppArgs struct {
	Configuration        ConfigurationPtrInput
	ContainerAppName     pulumi.StringPtrInput
	Identity             ManagedServiceIdentityPtrInput
	Location             pulumi.StringPtrInput
	ManagedEnvironmentId pulumi.StringPtrInput
	ResourceGroupName    pulumi.StringInput
	Tags                 pulumi.StringMapInput
	Template             TemplatePtrInput
}

func (ContainerAppArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerAppArgs)(nil)).Elem()
}

type ContainerAppInput interface {
	pulumi.Input

	ToContainerAppOutput() ContainerAppOutput
	ToContainerAppOutputWithContext(ctx context.Context) ContainerAppOutput
}

func (*ContainerApp) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerApp)(nil)).Elem()
}

func (i *ContainerApp) ToContainerAppOutput() ContainerAppOutput {
	return i.ToContainerAppOutputWithContext(context.Background())
}

func (i *ContainerApp) ToContainerAppOutputWithContext(ctx context.Context) ContainerAppOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerAppOutput)
}

type ContainerAppOutput struct{ *pulumi.OutputState }

func (ContainerAppOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerApp)(nil)).Elem()
}

func (o ContainerAppOutput) ToContainerAppOutput() ContainerAppOutput {
	return o
}

func (o ContainerAppOutput) ToContainerAppOutputWithContext(ctx context.Context) ContainerAppOutput {
	return o
}

func (o ContainerAppOutput) Configuration() ConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *ContainerApp) ConfigurationResponsePtrOutput { return v.Configuration }).(ConfigurationResponsePtrOutput)
}

func (o ContainerAppOutput) CustomDomainVerificationId() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerApp) pulumi.StringOutput { return v.CustomDomainVerificationId }).(pulumi.StringOutput)
}

func (o ContainerAppOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *ContainerApp) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o ContainerAppOutput) LatestRevisionFqdn() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerApp) pulumi.StringOutput { return v.LatestRevisionFqdn }).(pulumi.StringOutput)
}

func (o ContainerAppOutput) LatestRevisionName() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerApp) pulumi.StringOutput { return v.LatestRevisionName }).(pulumi.StringOutput)
}

func (o ContainerAppOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerApp) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ContainerAppOutput) ManagedEnvironmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerApp) pulumi.StringPtrOutput { return v.ManagedEnvironmentId }).(pulumi.StringPtrOutput)
}

func (o ContainerAppOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerApp) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ContainerAppOutput) OutboundIpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ContainerApp) pulumi.StringArrayOutput { return v.OutboundIpAddresses }).(pulumi.StringArrayOutput)
}

func (o ContainerAppOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerApp) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ContainerAppOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ContainerApp) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ContainerAppOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ContainerApp) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ContainerAppOutput) Template() TemplateResponsePtrOutput {
	return o.ApplyT(func(v *ContainerApp) TemplateResponsePtrOutput { return v.Template }).(TemplateResponsePtrOutput)
}

func (o ContainerAppOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerApp) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ContainerAppOutput{})
}
