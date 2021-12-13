


package v20150521preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LabResource struct {
	pulumi.CustomResourceState

	ArtifactsStorageAccount pulumi.StringPtrOutput   `pulumi:"artifactsStorageAccount"`
	CreatedDate             pulumi.StringPtrOutput   `pulumi:"createdDate"`
	DefaultStorageAccount   pulumi.StringPtrOutput   `pulumi:"defaultStorageAccount"`
	DefaultVirtualNetworkId pulumi.StringPtrOutput   `pulumi:"defaultVirtualNetworkId"`
	LabStorageType          pulumi.StringPtrOutput   `pulumi:"labStorageType"`
	Location                pulumi.StringPtrOutput   `pulumi:"location"`
	Name                    pulumi.StringPtrOutput   `pulumi:"name"`
	ProvisioningState       pulumi.StringPtrOutput   `pulumi:"provisioningState"`
	StorageAccounts         pulumi.StringArrayOutput `pulumi:"storageAccounts"`
	Tags                    pulumi.StringMapOutput   `pulumi:"tags"`
	Type                    pulumi.StringPtrOutput   `pulumi:"type"`
	VaultName               pulumi.StringPtrOutput   `pulumi:"vaultName"`
}


func NewLabResource(ctx *pulumi.Context,
	name string, args *LabResourceArgs, opts ...pulumi.ResourceOption) (*LabResource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devtestlab:LabResource"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20160515:LabResource"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20180915:LabResource"),
		},
	})
	opts = append(opts, aliases)
	var resource LabResource
	err := ctx.RegisterResource("azure-native:devtestlab/v20150521preview:LabResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetLabResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LabResourceState, opts ...pulumi.ResourceOption) (*LabResource, error) {
	var resource LabResource
	err := ctx.ReadResource("azure-native:devtestlab/v20150521preview:LabResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type labResourceState struct {
}

type LabResourceState struct {
}

func (LabResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*labResourceState)(nil)).Elem()
}

type labResourceArgs struct {
	ArtifactsStorageAccount *string           `pulumi:"artifactsStorageAccount"`
	CreatedDate             *string           `pulumi:"createdDate"`
	DefaultStorageAccount   *string           `pulumi:"defaultStorageAccount"`
	DefaultVirtualNetworkId *string           `pulumi:"defaultVirtualNetworkId"`
	Id                      *string           `pulumi:"id"`
	LabStorageType          *string           `pulumi:"labStorageType"`
	Location                *string           `pulumi:"location"`
	Name                    *string           `pulumi:"name"`
	ProvisioningState       *string           `pulumi:"provisioningState"`
	ResourceGroupName       string            `pulumi:"resourceGroupName"`
	StorageAccounts         []string          `pulumi:"storageAccounts"`
	Tags                    map[string]string `pulumi:"tags"`
	Type                    *string           `pulumi:"type"`
	VaultName               *string           `pulumi:"vaultName"`
}


type LabResourceArgs struct {
	ArtifactsStorageAccount pulumi.StringPtrInput
	CreatedDate             pulumi.StringPtrInput
	DefaultStorageAccount   pulumi.StringPtrInput
	DefaultVirtualNetworkId pulumi.StringPtrInput
	Id                      pulumi.StringPtrInput
	LabStorageType          pulumi.StringPtrInput
	Location                pulumi.StringPtrInput
	Name                    pulumi.StringPtrInput
	ProvisioningState       pulumi.StringPtrInput
	ResourceGroupName       pulumi.StringInput
	StorageAccounts         pulumi.StringArrayInput
	Tags                    pulumi.StringMapInput
	Type                    pulumi.StringPtrInput
	VaultName               pulumi.StringPtrInput
}

func (LabResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*labResourceArgs)(nil)).Elem()
}

type LabResourceInput interface {
	pulumi.Input

	ToLabResourceOutput() LabResourceOutput
	ToLabResourceOutputWithContext(ctx context.Context) LabResourceOutput
}

func (*LabResource) ElementType() reflect.Type {
	return reflect.TypeOf((**LabResource)(nil)).Elem()
}

func (i *LabResource) ToLabResourceOutput() LabResourceOutput {
	return i.ToLabResourceOutputWithContext(context.Background())
}

func (i *LabResource) ToLabResourceOutputWithContext(ctx context.Context) LabResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabResourceOutput)
}

type LabResourceOutput struct{ *pulumi.OutputState }

func (LabResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LabResource)(nil)).Elem()
}

func (o LabResourceOutput) ToLabResourceOutput() LabResourceOutput {
	return o
}

func (o LabResourceOutput) ToLabResourceOutputWithContext(ctx context.Context) LabResourceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(LabResourceOutput{})
}
