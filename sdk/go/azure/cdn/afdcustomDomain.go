


package cdn

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
			Type: pulumi.String("azure-nextgen:cdn:AFDCustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200901:AFDCustomDomain"),
		},
		{
			Type: pulumi.String("azure-nextgen:cdn/v20200901:AFDCustomDomain"),
		},
	})
	opts = append(opts, aliases)
	var resource AFDCustomDomain
	err := ctx.RegisterResource("azure-native:cdn:AFDCustomDomain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAFDCustomDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AFDCustomDomainState, opts ...pulumi.ResourceOption) (*AFDCustomDomain, error) {
	var resource AFDCustomDomain
	err := ctx.ReadResource("azure-native:cdn:AFDCustomDomain", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((*AFDCustomDomain)(nil))
}

func (i *AFDCustomDomain) ToAFDCustomDomainOutput() AFDCustomDomainOutput {
	return i.ToAFDCustomDomainOutputWithContext(context.Background())
}

func (i *AFDCustomDomain) ToAFDCustomDomainOutputWithContext(ctx context.Context) AFDCustomDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AFDCustomDomainOutput)
}

type AFDCustomDomainOutput struct{ *pulumi.OutputState }

func (AFDCustomDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AFDCustomDomain)(nil))
}

func (o AFDCustomDomainOutput) ToAFDCustomDomainOutput() AFDCustomDomainOutput {
	return o
}

func (o AFDCustomDomainOutput) ToAFDCustomDomainOutputWithContext(ctx context.Context) AFDCustomDomainOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AFDCustomDomainInput)(nil)).Elem(), &AFDCustomDomain{})
	pulumi.RegisterOutputType(AFDCustomDomainOutput{})
}
