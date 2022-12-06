


package v20220901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type OuContainer struct {
	pulumi.CustomResourceState

	Accounts          ContainerAccountResponseArrayOutput `pulumi:"accounts"`
	ContainerId       pulumi.StringOutput                 `pulumi:"containerId"`
	DeploymentId      pulumi.StringOutput                 `pulumi:"deploymentId"`
	DistinguishedName pulumi.StringOutput                 `pulumi:"distinguishedName"`
	DomainName        pulumi.StringOutput                 `pulumi:"domainName"`
	Etag              pulumi.StringPtrOutput              `pulumi:"etag"`
	Location          pulumi.StringPtrOutput              `pulumi:"location"`
	Name              pulumi.StringOutput                 `pulumi:"name"`
	ProvisioningState pulumi.StringOutput                 `pulumi:"provisioningState"`
	ServiceStatus     pulumi.StringOutput                 `pulumi:"serviceStatus"`
	SystemData        SystemDataResponseOutput            `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput              `pulumi:"tags"`
	TenantId          pulumi.StringOutput                 `pulumi:"tenantId"`
	Type              pulumi.StringOutput                 `pulumi:"type"`
}


func NewOuContainer(ctx *pulumi.Context,
	name string, args *OuContainerArgs, opts ...pulumi.ResourceOption) (*OuContainer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainServiceName == nil {
		return nil, errors.New("invalid value for required argument 'DomainServiceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:aad:OuContainer"),
		},
		{
			Type: pulumi.String("azure-native:aad/v20170601:OuContainer"),
		},
		{
			Type: pulumi.String("azure-native:aad/v20200101:OuContainer"),
		},
		{
			Type: pulumi.String("azure-native:aad/v20210301:OuContainer"),
		},
		{
			Type: pulumi.String("azure-native:aad/v20210501:OuContainer"),
		},
		{
			Type: pulumi.String("azure-native:aad/v20221201:OuContainer"),
		},
	})
	opts = append(opts, aliases)
	var resource OuContainer
	err := ctx.RegisterResource("azure-native:aad/v20220901:OuContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetOuContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OuContainerState, opts ...pulumi.ResourceOption) (*OuContainer, error) {
	var resource OuContainer
	err := ctx.ReadResource("azure-native:aad/v20220901:OuContainer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type ouContainerState struct {
}

type OuContainerState struct {
}

func (OuContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*ouContainerState)(nil)).Elem()
}

type ouContainerArgs struct {
	AccountName       *string `pulumi:"accountName"`
	DomainServiceName string  `pulumi:"domainServiceName"`
	OuContainerName   *string `pulumi:"ouContainerName"`
	Password          *string `pulumi:"password"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	Spn               *string `pulumi:"spn"`
}


type OuContainerArgs struct {
	AccountName       pulumi.StringPtrInput
	DomainServiceName pulumi.StringInput
	OuContainerName   pulumi.StringPtrInput
	Password          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Spn               pulumi.StringPtrInput
}

func (OuContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ouContainerArgs)(nil)).Elem()
}

type OuContainerInput interface {
	pulumi.Input

	ToOuContainerOutput() OuContainerOutput
	ToOuContainerOutputWithContext(ctx context.Context) OuContainerOutput
}

func (*OuContainer) ElementType() reflect.Type {
	return reflect.TypeOf((**OuContainer)(nil)).Elem()
}

func (i *OuContainer) ToOuContainerOutput() OuContainerOutput {
	return i.ToOuContainerOutputWithContext(context.Background())
}

func (i *OuContainer) ToOuContainerOutputWithContext(ctx context.Context) OuContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OuContainerOutput)
}

type OuContainerOutput struct{ *pulumi.OutputState }

func (OuContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OuContainer)(nil)).Elem()
}

func (o OuContainerOutput) ToOuContainerOutput() OuContainerOutput {
	return o
}

func (o OuContainerOutput) ToOuContainerOutputWithContext(ctx context.Context) OuContainerOutput {
	return o
}

func (o OuContainerOutput) Accounts() ContainerAccountResponseArrayOutput {
	return o.ApplyT(func(v *OuContainer) ContainerAccountResponseArrayOutput { return v.Accounts }).(ContainerAccountResponseArrayOutput)
}

func (o OuContainerOutput) ContainerId() pulumi.StringOutput {
	return o.ApplyT(func(v *OuContainer) pulumi.StringOutput { return v.ContainerId }).(pulumi.StringOutput)
}

func (o OuContainerOutput) DeploymentId() pulumi.StringOutput {
	return o.ApplyT(func(v *OuContainer) pulumi.StringOutput { return v.DeploymentId }).(pulumi.StringOutput)
}

func (o OuContainerOutput) DistinguishedName() pulumi.StringOutput {
	return o.ApplyT(func(v *OuContainer) pulumi.StringOutput { return v.DistinguishedName }).(pulumi.StringOutput)
}

func (o OuContainerOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *OuContainer) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

func (o OuContainerOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OuContainer) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o OuContainerOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OuContainer) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o OuContainerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OuContainer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o OuContainerOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *OuContainer) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o OuContainerOutput) ServiceStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *OuContainer) pulumi.StringOutput { return v.ServiceStatus }).(pulumi.StringOutput)
}

func (o OuContainerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *OuContainer) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o OuContainerOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *OuContainer) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o OuContainerOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *OuContainer) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

func (o OuContainerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *OuContainer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(OuContainerOutput{})
}
