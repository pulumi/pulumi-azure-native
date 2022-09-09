


package attestation

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AttestationProvider struct {
	pulumi.CustomResourceState

	AttestUri                  pulumi.StringPtrOutput                       `pulumi:"attestUri"`
	Location                   pulumi.StringOutput                          `pulumi:"location"`
	Name                       pulumi.StringOutput                          `pulumi:"name"`
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	Status                     pulumi.StringPtrOutput                       `pulumi:"status"`
	SystemData                 SystemDataResponseOutput                     `pulumi:"systemData"`
	Tags                       pulumi.StringMapOutput                       `pulumi:"tags"`
	TrustModel                 pulumi.StringPtrOutput                       `pulumi:"trustModel"`
	Type                       pulumi.StringOutput                          `pulumi:"type"`
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
			Type: pulumi.String("azure-native:attestation/v20180901preview:AttestationProvider"),
		},
		{
			Type: pulumi.String("azure-native:attestation/v20201001:AttestationProvider"),
		},
		{
			Type: pulumi.String("azure-native:attestation/v20210601preview:AttestationProvider"),
		},
	})
	opts = append(opts, aliases)
	var resource AttestationProvider
	err := ctx.RegisterResource("azure-native:attestation:AttestationProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAttestationProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AttestationProviderState, opts ...pulumi.ResourceOption) (*AttestationProvider, error) {
	var resource AttestationProvider
	err := ctx.ReadResource("azure-native:attestation:AttestationProvider", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((**AttestationProvider)(nil)).Elem()
}

func (i *AttestationProvider) ToAttestationProviderOutput() AttestationProviderOutput {
	return i.ToAttestationProviderOutputWithContext(context.Background())
}

func (i *AttestationProvider) ToAttestationProviderOutputWithContext(ctx context.Context) AttestationProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttestationProviderOutput)
}

type AttestationProviderOutput struct{ *pulumi.OutputState }

func (AttestationProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AttestationProvider)(nil)).Elem()
}

func (o AttestationProviderOutput) ToAttestationProviderOutput() AttestationProviderOutput {
	return o
}

func (o AttestationProviderOutput) ToAttestationProviderOutputWithContext(ctx context.Context) AttestationProviderOutput {
	return o
}

func (o AttestationProviderOutput) AttestUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttestationProvider) pulumi.StringPtrOutput { return v.AttestUri }).(pulumi.StringPtrOutput)
}

func (o AttestationProviderOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AttestationProvider) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o AttestationProviderOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AttestationProvider) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AttestationProviderOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *AttestationProvider) PrivateEndpointConnectionResponseArrayOutput {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o AttestationProviderOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttestationProvider) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

func (o AttestationProviderOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AttestationProvider) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o AttestationProviderOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AttestationProvider) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o AttestationProviderOutput) TrustModel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttestationProvider) pulumi.StringPtrOutput { return v.TrustModel }).(pulumi.StringPtrOutput)
}

func (o AttestationProviderOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AttestationProvider) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AttestationProviderOutput{})
}
