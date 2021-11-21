


package v20211201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateCloud struct {
	pulumi.CustomResourceState

	Availability                 AvailabilityPropertiesResponsePtrOutput `pulumi:"availability"`
	Circuit                      CircuitResponsePtrOutput                `pulumi:"circuit"`
	Encryption                   EncryptionResponsePtrOutput             `pulumi:"encryption"`
	Endpoints                    EndpointsResponseOutput                 `pulumi:"endpoints"`
	ExternalCloudLinks           pulumi.StringArrayOutput                `pulumi:"externalCloudLinks"`
	Identity                     PrivateCloudIdentityResponsePtrOutput   `pulumi:"identity"`
	IdentitySources              IdentitySourceResponseArrayOutput       `pulumi:"identitySources"`
	Internet                     pulumi.StringPtrOutput                  `pulumi:"internet"`
	Location                     pulumi.StringOutput                     `pulumi:"location"`
	ManagementCluster            ManagementClusterResponseOutput         `pulumi:"managementCluster"`
	ManagementNetwork            pulumi.StringOutput                     `pulumi:"managementNetwork"`
	Name                         pulumi.StringOutput                     `pulumi:"name"`
	NetworkBlock                 pulumi.StringOutput                     `pulumi:"networkBlock"`
	NsxtCertificateThumbprint    pulumi.StringOutput                     `pulumi:"nsxtCertificateThumbprint"`
	NsxtPassword                 pulumi.StringPtrOutput                  `pulumi:"nsxtPassword"`
	ProvisioningNetwork          pulumi.StringOutput                     `pulumi:"provisioningNetwork"`
	ProvisioningState            pulumi.StringOutput                     `pulumi:"provisioningState"`
	SecondaryCircuit             CircuitResponsePtrOutput                `pulumi:"secondaryCircuit"`
	Sku                          SkuResponseOutput                       `pulumi:"sku"`
	Tags                         pulumi.StringMapOutput                  `pulumi:"tags"`
	Type                         pulumi.StringOutput                     `pulumi:"type"`
	VcenterCertificateThumbprint pulumi.StringOutput                     `pulumi:"vcenterCertificateThumbprint"`
	VcenterPassword              pulumi.StringPtrOutput                  `pulumi:"vcenterPassword"`
	VmotionNetwork               pulumi.StringOutput                     `pulumi:"vmotionNetwork"`
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
	if args.Internet == nil {
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
			Type: pulumi.String("azure-native:avs/v20210601:PrivateCloud"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateCloud
	err := ctx.RegisterResource("azure-native:avs/v20211201:PrivateCloud", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateCloud(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateCloudState, opts ...pulumi.ResourceOption) (*PrivateCloud, error) {
	var resource PrivateCloud
	err := ctx.ReadResource("azure-native:avs/v20211201:PrivateCloud", name, id, state, &resource, opts...)
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
	Availability      *AvailabilityProperties `pulumi:"availability"`
	Encryption        *Encryption             `pulumi:"encryption"`
	Identity          *PrivateCloudIdentity   `pulumi:"identity"`
	IdentitySources   []IdentitySource        `pulumi:"identitySources"`
	Internet          *string                 `pulumi:"internet"`
	Location          *string                 `pulumi:"location"`
	ManagementCluster ManagementCluster       `pulumi:"managementCluster"`
	NetworkBlock      string                  `pulumi:"networkBlock"`
	NsxtPassword      *string                 `pulumi:"nsxtPassword"`
	PrivateCloudName  *string                 `pulumi:"privateCloudName"`
	ResourceGroupName string                  `pulumi:"resourceGroupName"`
	Sku               Sku                     `pulumi:"sku"`
	Tags              map[string]string       `pulumi:"tags"`
	VcenterPassword   *string                 `pulumi:"vcenterPassword"`
}


type PrivateCloudArgs struct {
	Availability      AvailabilityPropertiesPtrInput
	Encryption        EncryptionPtrInput
	Identity          PrivateCloudIdentityPtrInput
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
	return reflect.TypeOf((*PrivateCloud)(nil))
}

func (i *PrivateCloud) ToPrivateCloudOutput() PrivateCloudOutput {
	return i.ToPrivateCloudOutputWithContext(context.Background())
}

func (i *PrivateCloud) ToPrivateCloudOutputWithContext(ctx context.Context) PrivateCloudOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateCloudOutput)
}

type PrivateCloudOutput struct{ *pulumi.OutputState }

func (PrivateCloudOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateCloud)(nil))
}

func (o PrivateCloudOutput) ToPrivateCloudOutput() PrivateCloudOutput {
	return o
}

func (o PrivateCloudOutput) ToPrivateCloudOutputWithContext(ctx context.Context) PrivateCloudOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PrivateCloudOutput{})
}
