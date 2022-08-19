


package v20200901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AFDCustomDomain struct {
	pulumi.CustomResourceState

	AzureDnsZone          ResourceReferenceResponsePtrOutput        `pulumi:"azureDnsZone"`
	DeploymentStatus      pulumi.StringOutput                       `pulumi:"deploymentStatus"`
	DomainValidationState pulumi.StringOutput                       `pulumi:"domainValidationState"`
	HostName              pulumi.StringOutput                       `pulumi:"hostName"`
	Name                  pulumi.StringOutput                       `pulumi:"name"`
	ProvisioningState     pulumi.StringOutput                       `pulumi:"provisioningState"`
	SystemData            SystemDataResponseOutput                  `pulumi:"systemData"`
	TlsSettings           AFDDomainHttpsParametersResponsePtrOutput `pulumi:"tlsSettings"`
	Type                  pulumi.StringOutput                       `pulumi:"type"`
	ValidationProperties  DomainValidationPropertiesResponseOutput  `pulumi:"validationProperties"`
}


func NewAFDCustomDomain(ctx *pulumi.Context,
	name string, args *AFDCustomDomainArgs, opts ...pulumi.ResourceOption) (*AFDCustomDomain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HostName == nil {
		return nil, errors.New("invalid value for required argument 'HostName'")
	}
	if args.ProfileName == nil {
		return nil, errors.New("invalid value for required argument 'ProfileName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cdn:AFDCustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20210601:AFDCustomDomain"),
		},
	})
	opts = append(opts, aliases)
	var resource AFDCustomDomain
	err := ctx.RegisterResource("azure-native:cdn/v20200901:AFDCustomDomain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAFDCustomDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AFDCustomDomainState, opts ...pulumi.ResourceOption) (*AFDCustomDomain, error) {
	var resource AFDCustomDomain
	err := ctx.ReadResource("azure-native:cdn/v20200901:AFDCustomDomain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type afdcustomDomainState struct {
}

type AFDCustomDomainState struct {
}

func (AFDCustomDomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*afdcustomDomainState)(nil)).Elem()
}

type afdcustomDomainArgs struct {
	AzureDnsZone      *ResourceReference        `pulumi:"azureDnsZone"`
	CustomDomainName  *string                   `pulumi:"customDomainName"`
	HostName          string                    `pulumi:"hostName"`
	ProfileName       string                    `pulumi:"profileName"`
	ResourceGroupName string                    `pulumi:"resourceGroupName"`
	TlsSettings       *AFDDomainHttpsParameters `pulumi:"tlsSettings"`
}


type AFDCustomDomainArgs struct {
	AzureDnsZone      ResourceReferencePtrInput
	CustomDomainName  pulumi.StringPtrInput
	HostName          pulumi.StringInput
	ProfileName       pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	TlsSettings       AFDDomainHttpsParametersPtrInput
}

func (AFDCustomDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*afdcustomDomainArgs)(nil)).Elem()
}

type AFDCustomDomainInput interface {
	pulumi.Input

	ToAFDCustomDomainOutput() AFDCustomDomainOutput
	ToAFDCustomDomainOutputWithContext(ctx context.Context) AFDCustomDomainOutput
}

func (*AFDCustomDomain) ElementType() reflect.Type {
	return reflect.TypeOf((**AFDCustomDomain)(nil)).Elem()
}

func (i *AFDCustomDomain) ToAFDCustomDomainOutput() AFDCustomDomainOutput {
	return i.ToAFDCustomDomainOutputWithContext(context.Background())
}

func (i *AFDCustomDomain) ToAFDCustomDomainOutputWithContext(ctx context.Context) AFDCustomDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AFDCustomDomainOutput)
}

type AFDCustomDomainOutput struct{ *pulumi.OutputState }

func (AFDCustomDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AFDCustomDomain)(nil)).Elem()
}

func (o AFDCustomDomainOutput) ToAFDCustomDomainOutput() AFDCustomDomainOutput {
	return o
}

func (o AFDCustomDomainOutput) ToAFDCustomDomainOutputWithContext(ctx context.Context) AFDCustomDomainOutput {
	return o
}

func (o AFDCustomDomainOutput) AzureDnsZone() ResourceReferenceResponsePtrOutput {
	return o.ApplyT(func(v *AFDCustomDomain) ResourceReferenceResponsePtrOutput { return v.AzureDnsZone }).(ResourceReferenceResponsePtrOutput)
}

func (o AFDCustomDomainOutput) DeploymentStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *AFDCustomDomain) pulumi.StringOutput { return v.DeploymentStatus }).(pulumi.StringOutput)
}

func (o AFDCustomDomainOutput) DomainValidationState() pulumi.StringOutput {
	return o.ApplyT(func(v *AFDCustomDomain) pulumi.StringOutput { return v.DomainValidationState }).(pulumi.StringOutput)
}

func (o AFDCustomDomainOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v *AFDCustomDomain) pulumi.StringOutput { return v.HostName }).(pulumi.StringOutput)
}

func (o AFDCustomDomainOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AFDCustomDomain) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AFDCustomDomainOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AFDCustomDomain) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o AFDCustomDomainOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AFDCustomDomain) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o AFDCustomDomainOutput) TlsSettings() AFDDomainHttpsParametersResponsePtrOutput {
	return o.ApplyT(func(v *AFDCustomDomain) AFDDomainHttpsParametersResponsePtrOutput { return v.TlsSettings }).(AFDDomainHttpsParametersResponsePtrOutput)
}

func (o AFDCustomDomainOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AFDCustomDomain) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o AFDCustomDomainOutput) ValidationProperties() DomainValidationPropertiesResponseOutput {
	return o.ApplyT(func(v *AFDCustomDomain) DomainValidationPropertiesResponseOutput { return v.ValidationProperties }).(DomainValidationPropertiesResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(AFDCustomDomainOutput{})
}
