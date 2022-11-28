


package signalrservice

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SignalRCustomDomain struct {
	pulumi.CustomResourceState

	CustomCertificate ResourceReferenceResponseOutput `pulumi:"customCertificate"`
	DomainName        pulumi.StringOutput             `pulumi:"domainName"`
	Name              pulumi.StringOutput             `pulumi:"name"`
	ProvisioningState pulumi.StringOutput             `pulumi:"provisioningState"`
	SystemData        SystemDataResponseOutput        `pulumi:"systemData"`
	Type              pulumi.StringOutput             `pulumi:"type"`
}


func NewSignalRCustomDomain(ctx *pulumi.Context,
	name string, args *SignalRCustomDomainArgs, opts ...pulumi.ResourceOption) (*SignalRCustomDomain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CustomCertificate == nil {
		return nil, errors.New("invalid value for required argument 'CustomCertificate'")
	}
	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:signalrservice/v20220201:SignalRCustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20220801preview:SignalRCustomDomain"),
		},
	})
	opts = append(opts, aliases)
	var resource SignalRCustomDomain
	err := ctx.RegisterResource("azure-native:signalrservice:SignalRCustomDomain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSignalRCustomDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SignalRCustomDomainState, opts ...pulumi.ResourceOption) (*SignalRCustomDomain, error) {
	var resource SignalRCustomDomain
	err := ctx.ReadResource("azure-native:signalrservice:SignalRCustomDomain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type signalRCustomDomainState struct {
}

type SignalRCustomDomainState struct {
}

func (SignalRCustomDomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*signalRCustomDomainState)(nil)).Elem()
}

type signalRCustomDomainArgs struct {
	CustomCertificate ResourceReference `pulumi:"customCertificate"`
	DomainName        string            `pulumi:"domainName"`
	Name              *string           `pulumi:"name"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	ResourceName      string            `pulumi:"resourceName"`
}


type SignalRCustomDomainArgs struct {
	CustomCertificate ResourceReferenceInput
	DomainName        pulumi.StringInput
	Name              pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringInput
}

func (SignalRCustomDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*signalRCustomDomainArgs)(nil)).Elem()
}

type SignalRCustomDomainInput interface {
	pulumi.Input

	ToSignalRCustomDomainOutput() SignalRCustomDomainOutput
	ToSignalRCustomDomainOutputWithContext(ctx context.Context) SignalRCustomDomainOutput
}

func (*SignalRCustomDomain) ElementType() reflect.Type {
	return reflect.TypeOf((**SignalRCustomDomain)(nil)).Elem()
}

func (i *SignalRCustomDomain) ToSignalRCustomDomainOutput() SignalRCustomDomainOutput {
	return i.ToSignalRCustomDomainOutputWithContext(context.Background())
}

func (i *SignalRCustomDomain) ToSignalRCustomDomainOutputWithContext(ctx context.Context) SignalRCustomDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SignalRCustomDomainOutput)
}

type SignalRCustomDomainOutput struct{ *pulumi.OutputState }

func (SignalRCustomDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SignalRCustomDomain)(nil)).Elem()
}

func (o SignalRCustomDomainOutput) ToSignalRCustomDomainOutput() SignalRCustomDomainOutput {
	return o
}

func (o SignalRCustomDomainOutput) ToSignalRCustomDomainOutputWithContext(ctx context.Context) SignalRCustomDomainOutput {
	return o
}

func (o SignalRCustomDomainOutput) CustomCertificate() ResourceReferenceResponseOutput {
	return o.ApplyT(func(v *SignalRCustomDomain) ResourceReferenceResponseOutput { return v.CustomCertificate }).(ResourceReferenceResponseOutput)
}

func (o SignalRCustomDomainOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalRCustomDomain) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

func (o SignalRCustomDomainOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalRCustomDomain) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SignalRCustomDomainOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalRCustomDomain) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o SignalRCustomDomainOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SignalRCustomDomain) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SignalRCustomDomainOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalRCustomDomain) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SignalRCustomDomainOutput{})
}
