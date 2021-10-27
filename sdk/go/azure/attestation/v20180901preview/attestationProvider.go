


package v20180901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AttestationProvider struct {
	pulumi.CustomResourceState

	AttestUri  pulumi.StringPtrOutput `pulumi:"attestUri"`
	Location   pulumi.StringOutput    `pulumi:"location"`
	Name       pulumi.StringOutput    `pulumi:"name"`
	Status     pulumi.StringPtrOutput `pulumi:"status"`
	Tags       pulumi.StringMapOutput `pulumi:"tags"`
	TrustModel pulumi.StringPtrOutput `pulumi:"trustModel"`
	Type       pulumi.StringOutput    `pulumi:"type"`
}


func NewAttestationProvider(ctx *pulumi.Context,
	name string, args *AttestationProviderArgs, opts ...pulumi.ResourceOption) (*AttestationProvider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:attestation/v20180901preview:AttestationProvider"),
		},
		{
			Type: pulumi.String("azure-native:attestation:AttestationProvider"),
		},
		{
			Type: pulumi.String("azure-nextgen:attestation:AttestationProvider"),
		},
		{
			Type: pulumi.String("azure-native:attestation/v20201001:AttestationProvider"),
		},
		{
			Type: pulumi.String("azure-nextgen:attestation/v20201001:AttestationProvider"),
		},
		{
			Type: pulumi.String("azure-native:attestation/v20210601preview:AttestationProvider"),
		},
		{
			Type: pulumi.String("azure-nextgen:attestation/v20210601preview:AttestationProvider"),
		},
	})
	opts = append(opts, aliases)
	var resource AttestationProvider
	err := ctx.RegisterResource("azure-native:attestation/v20180901preview:AttestationProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAttestationProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AttestationProviderState, opts ...pulumi.ResourceOption) (*AttestationProvider, error) {
	var resource AttestationProvider
	err := ctx.ReadResource("azure-native:attestation/v20180901preview:AttestationProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type attestationProviderState struct {
}

type AttestationProviderState struct {
}

func (AttestationProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*attestationProviderState)(nil)).Elem()
}

type attestationProviderArgs struct {
	Location          *string                                  `pulumi:"location"`
	Properties        AttestationServiceCreationSpecificParams `pulumi:"properties"`
	ProviderName      *string                                  `pulumi:"providerName"`
	ResourceGroupName string                                   `pulumi:"resourceGroupName"`
	Tags              map[string]string                        `pulumi:"tags"`
}


type AttestationProviderArgs struct {
	Location          pulumi.StringPtrInput
	Properties        AttestationServiceCreationSpecificParamsInput
	ProviderName      pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (AttestationProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*attestationProviderArgs)(nil)).Elem()
}

type AttestationProviderInput interface {
	pulumi.Input

	ToAttestationProviderOutput() AttestationProviderOutput
	ToAttestationProviderOutputWithContext(ctx context.Context) AttestationProviderOutput
}

func (*AttestationProvider) ElementType() reflect.Type {
	return reflect.TypeOf((*AttestationProvider)(nil))
}

func (i *AttestationProvider) ToAttestationProviderOutput() AttestationProviderOutput {
	return i.ToAttestationProviderOutputWithContext(context.Background())
}

func (i *AttestationProvider) ToAttestationProviderOutputWithContext(ctx context.Context) AttestationProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttestationProviderOutput)
}

type AttestationProviderOutput struct{ *pulumi.OutputState }

func (AttestationProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AttestationProvider)(nil))
}

func (o AttestationProviderOutput) ToAttestationProviderOutput() AttestationProviderOutput {
	return o
}

func (o AttestationProviderOutput) ToAttestationProviderOutputWithContext(ctx context.Context) AttestationProviderOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AttestationProviderInput)(nil)).Elem(), &AttestationProvider{})
	pulumi.RegisterOutputType(AttestationProviderOutput{})
}
