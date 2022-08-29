


package v20210601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateCloud struct {
	pulumi.CustomResourceState

	Circuit                      CircuitResponsePtrOutput          `pulumi:"circuit"`
	Endpoints                    EndpointsResponseOutput           `pulumi:"endpoints"`
	ExternalCloudLinks           pulumi.StringArrayOutput          `pulumi:"externalCloudLinks"`
	IdentitySources              IdentitySourceResponseArrayOutput `pulumi:"identitySources"`
	Internet                     pulumi.StringPtrOutput            `pulumi:"internet"`
	Location                     pulumi.StringOutput               `pulumi:"location"`
	ManagementCluster            ManagementClusterResponseOutput   `pulumi:"managementCluster"`
	ManagementNetwork            pulumi.StringOutput               `pulumi:"managementNetwork"`
	Name                         pulumi.StringOutput               `pulumi:"name"`
	NetworkBlock                 pulumi.StringOutput               `pulumi:"networkBlock"`
	NsxtCertificateThumbprint    pulumi.StringOutput               `pulumi:"nsxtCertificateThumbprint"`
	NsxtPassword                 pulumi.StringPtrOutput            `pulumi:"nsxtPassword"`
	ProvisioningNetwork          pulumi.StringOutput               `pulumi:"provisioningNetwork"`
	ProvisioningState            pulumi.StringOutput               `pulumi:"provisioningState"`
	Sku                          SkuResponseOutput                 `pulumi:"sku"`
	Tags                         pulumi.StringMapOutput            `pulumi:"tags"`
	Type                         pulumi.StringOutput               `pulumi:"type"`
	VcenterCertificateThumbprint pulumi.StringOutput               `pulumi:"vcenterCertificateThumbprint"`
	VcenterPassword              pulumi.StringPtrOutput            `pulumi:"vcenterPassword"`
	VmotionNetwork               pulumi.StringOutput               `pulumi:"vmotionNetwork"`
}


func NewPrivateCloud(ctx *pulumi.Context,
	name string, args *PrivateCloudArgs, opts ...pulumi.ResourceOption) (*PrivateCloud, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagementCluster == nil {
		return nil, errors.New("invalid value for required argument 'ManagementCluster'")
	}
	if args.NetworkBlock == nil {
		return nil, errors.New("invalid value for required argument 'NetworkBlock'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	if isZero(args.Internet) {
		args.Internet = pulumi.StringPtr("Disabled")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:avs:PrivateCloud"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20200320:PrivateCloud"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20200717preview:PrivateCloud"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210101preview:PrivateCloud"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20211201:PrivateCloud"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateCloud
	err := ctx.RegisterResource("azure-native:avs/v20210601:PrivateCloud", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateCloud(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateCloudState, opts ...pulumi.ResourceOption) (*PrivateCloud, error) {
	var resource PrivateCloud
	err := ctx.ReadResource("azure-native:avs/v20210601:PrivateCloud", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type privateCloudState struct {
}

type PrivateCloudState struct {
}

func (PrivateCloudState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateCloudState)(nil)).Elem()
}

type privateCloudArgs struct {
	IdentitySources   []IdentitySource  `pulumi:"identitySources"`
	Internet          *string           `pulumi:"internet"`
	Location          *string           `pulumi:"location"`
	ManagementCluster ManagementCluster `pulumi:"managementCluster"`
	NetworkBlock      string            `pulumi:"networkBlock"`
	NsxtPassword      *string           `pulumi:"nsxtPassword"`
	PrivateCloudName  *string           `pulumi:"privateCloudName"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Sku               Sku               `pulumi:"sku"`
	Tags              map[string]string `pulumi:"tags"`
	VcenterPassword   *string           `pulumi:"vcenterPassword"`
}


type PrivateCloudArgs struct {
	IdentitySources   IdentitySourceArrayInput
	Internet          pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	ManagementCluster ManagementClusterInput
	NetworkBlock      pulumi.StringInput
	NsxtPassword      pulumi.StringPtrInput
	PrivateCloudName  pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Sku               SkuInput
	Tags              pulumi.StringMapInput
	VcenterPassword   pulumi.StringPtrInput
}

func (PrivateCloudArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateCloudArgs)(nil)).Elem()
}

type PrivateCloudInput interface {
	pulumi.Input

	ToPrivateCloudOutput() PrivateCloudOutput
	ToPrivateCloudOutputWithContext(ctx context.Context) PrivateCloudOutput
}

func (*PrivateCloud) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateCloud)(nil)).Elem()
}

func (i *PrivateCloud) ToPrivateCloudOutput() PrivateCloudOutput {
	return i.ToPrivateCloudOutputWithContext(context.Background())
}

func (i *PrivateCloud) ToPrivateCloudOutputWithContext(ctx context.Context) PrivateCloudOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateCloudOutput)
}

type PrivateCloudOutput struct{ *pulumi.OutputState }

func (PrivateCloudOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateCloud)(nil)).Elem()
}

func (o PrivateCloudOutput) ToPrivateCloudOutput() PrivateCloudOutput {
	return o
}

func (o PrivateCloudOutput) ToPrivateCloudOutputWithContext(ctx context.Context) PrivateCloudOutput {
	return o
}

func (o PrivateCloudOutput) Circuit() CircuitResponsePtrOutput {
	return o.ApplyT(func(v *PrivateCloud) CircuitResponsePtrOutput { return v.Circuit }).(CircuitResponsePtrOutput)
}

func (o PrivateCloudOutput) Endpoints() EndpointsResponseOutput {
	return o.ApplyT(func(v *PrivateCloud) EndpointsResponseOutput { return v.Endpoints }).(EndpointsResponseOutput)
}

func (o PrivateCloudOutput) ExternalCloudLinks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PrivateCloud) pulumi.StringArrayOutput { return v.ExternalCloudLinks }).(pulumi.StringArrayOutput)
}

func (o PrivateCloudOutput) IdentitySources() IdentitySourceResponseArrayOutput {
	return o.ApplyT(func(v *PrivateCloud) IdentitySourceResponseArrayOutput { return v.IdentitySources }).(IdentitySourceResponseArrayOutput)
}

func (o PrivateCloudOutput) Internet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateCloud) pulumi.StringPtrOutput { return v.Internet }).(pulumi.StringPtrOutput)
}

func (o PrivateCloudOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateCloud) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o PrivateCloudOutput) ManagementCluster() ManagementClusterResponseOutput {
	return o.ApplyT(func(v *PrivateCloud) ManagementClusterResponseOutput { return v.ManagementCluster }).(ManagementClusterResponseOutput)
}

func (o PrivateCloudOutput) ManagementNetwork() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateCloud) pulumi.StringOutput { return v.ManagementNetwork }).(pulumi.StringOutput)
}

func (o PrivateCloudOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateCloud) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateCloudOutput) NetworkBlock() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateCloud) pulumi.StringOutput { return v.NetworkBlock }).(pulumi.StringOutput)
}

func (o PrivateCloudOutput) NsxtCertificateThumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateCloud) pulumi.StringOutput { return v.NsxtCertificateThumbprint }).(pulumi.StringOutput)
}

func (o PrivateCloudOutput) NsxtPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateCloud) pulumi.StringPtrOutput { return v.NsxtPassword }).(pulumi.StringPtrOutput)
}

func (o PrivateCloudOutput) ProvisioningNetwork() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateCloud) pulumi.StringOutput { return v.ProvisioningNetwork }).(pulumi.StringOutput)
}

func (o PrivateCloudOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateCloud) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PrivateCloudOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v *PrivateCloud) SkuResponseOutput { return v.Sku }).(SkuResponseOutput)
}

func (o PrivateCloudOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PrivateCloud) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o PrivateCloudOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateCloud) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o PrivateCloudOutput) VcenterCertificateThumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateCloud) pulumi.StringOutput { return v.VcenterCertificateThumbprint }).(pulumi.StringOutput)
}

func (o PrivateCloudOutput) VcenterPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateCloud) pulumi.StringPtrOutput { return v.VcenterPassword }).(pulumi.StringPtrOutput)
}

func (o PrivateCloudOutput) VmotionNetwork() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateCloud) pulumi.StringOutput { return v.VmotionNetwork }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateCloudOutput{})
}
