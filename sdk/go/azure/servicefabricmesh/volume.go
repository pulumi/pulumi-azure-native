


package servicefabricmesh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Volume struct {
	pulumi.CustomResourceState

	AzureFileParameters VolumeProviderParametersAzureFileResponsePtrOutput `pulumi:"azureFileParameters"`
	Description         pulumi.StringPtrOutput                             `pulumi:"description"`
	Location            pulumi.StringOutput                                `pulumi:"location"`
	Name                pulumi.StringOutput                                `pulumi:"name"`
	Provider            pulumi.StringOutput                                `pulumi:"provider"`
	ProvisioningState   pulumi.StringOutput                                `pulumi:"provisioningState"`
	Status              pulumi.StringOutput                                `pulumi:"status"`
	StatusDetails       pulumi.StringOutput                                `pulumi:"statusDetails"`
	Tags                pulumi.StringMapOutput                             `pulumi:"tags"`
	Type                pulumi.StringOutput                                `pulumi:"type"`
}


func NewVolume(ctx *pulumi.Context,
	name string, args *VolumeArgs, opts ...pulumi.ResourceOption) (*Volume, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Provider == nil {
		return nil, errors.New("invalid value for required argument 'Provider'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicefabricmesh/v20180701preview:Volume"),
		},
		{
			Type: pulumi.String("azure-native:servicefabricmesh/v20180901preview:Volume"),
		},
	})
	opts = append(opts, aliases)
	var resource Volume
	err := ctx.RegisterResource("azure-native:servicefabricmesh:Volume", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVolume(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VolumeState, opts ...pulumi.ResourceOption) (*Volume, error) {
	var resource Volume
	err := ctx.ReadResource("azure-native:servicefabricmesh:Volume", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type volumeState struct {
}

type VolumeState struct {
}

func (VolumeState) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeState)(nil)).Elem()
}

type volumeArgs struct {
	AzureFileParameters *VolumeProviderParametersAzureFile `pulumi:"azureFileParameters"`
	Description         *string                            `pulumi:"description"`
	Location            *string                            `pulumi:"location"`
	Provider            string                             `pulumi:"provider"`
	ResourceGroupName   string                             `pulumi:"resourceGroupName"`
	Tags                map[string]string                  `pulumi:"tags"`
	VolumeResourceName  *string                            `pulumi:"volumeResourceName"`
}


type VolumeArgs struct {
	AzureFileParameters VolumeProviderParametersAzureFilePtrInput
	Description         pulumi.StringPtrInput
	Location            pulumi.StringPtrInput
	Provider            pulumi.StringInput
	ResourceGroupName   pulumi.StringInput
	Tags                pulumi.StringMapInput
	VolumeResourceName  pulumi.StringPtrInput
}

func (VolumeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeArgs)(nil)).Elem()
}

type VolumeInput interface {
	pulumi.Input

	ToVolumeOutput() VolumeOutput
	ToVolumeOutputWithContext(ctx context.Context) VolumeOutput
}

func (*Volume) ElementType() reflect.Type {
	return reflect.TypeOf((*Volume)(nil))
}

func (i *Volume) ToVolumeOutput() VolumeOutput {
	return i.ToVolumeOutputWithContext(context.Background())
}

func (i *Volume) ToVolumeOutputWithContext(ctx context.Context) VolumeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeOutput)
}

type VolumeOutput struct{ *pulumi.OutputState }

func (VolumeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Volume)(nil))
}

func (o VolumeOutput) ToVolumeOutput() VolumeOutput {
	return o
}

func (o VolumeOutput) ToVolumeOutputWithContext(ctx context.Context) VolumeOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VolumeOutput{})
}
