


package v20160515

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type Lab struct {
	pulumi.CustomResourceState

	ArtifactsStorageAccount       pulumi.StringOutput    `pulumi:"artifactsStorageAccount"`
	CreatedDate                   pulumi.StringOutput    `pulumi:"createdDate"`
	DefaultPremiumStorageAccount  pulumi.StringOutput    `pulumi:"defaultPremiumStorageAccount"`
	DefaultStorageAccount         pulumi.StringOutput    `pulumi:"defaultStorageAccount"`
	LabStorageType                pulumi.StringPtrOutput `pulumi:"labStorageType"`
	Location                      pulumi.StringPtrOutput `pulumi:"location"`
	Name                          pulumi.StringOutput    `pulumi:"name"`
	PremiumDataDiskStorageAccount pulumi.StringOutput    `pulumi:"premiumDataDiskStorageAccount"`
	PremiumDataDisks              pulumi.StringPtrOutput `pulumi:"premiumDataDisks"`
	ProvisioningState             pulumi.StringPtrOutput `pulumi:"provisioningState"`
	Tags                          pulumi.StringMapOutput `pulumi:"tags"`
	Type                          pulumi.StringOutput    `pulumi:"type"`
	UniqueIdentifier              pulumi.StringPtrOutput `pulumi:"uniqueIdentifier"`
	VaultName                     pulumi.StringOutput    `pulumi:"vaultName"`
}


func NewLab(ctx *pulumi.Context,
	name string, args *LabArgs, opts ...pulumi.ResourceOption) (*Lab, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devtestlab:Lab"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20150521preview:Lab"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20180915:Lab"),
		},
	})
	opts = append(opts, aliases)
	var resource Lab
	err := ctx.RegisterResource("azure-native:devtestlab/v20160515:Lab", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetLab(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LabState, opts ...pulumi.ResourceOption) (*Lab, error) {
	var resource Lab
	err := ctx.ReadResource("azure-native:devtestlab/v20160515:Lab", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type labState struct {
}

type LabState struct {
}

func (LabState) ElementType() reflect.Type {
	return reflect.TypeOf((*labState)(nil)).Elem()
}

type labArgs struct {
	LabStorageType    *string           `pulumi:"labStorageType"`
	Location          *string           `pulumi:"location"`
	Name              *string           `pulumi:"name"`
	PremiumDataDisks  *string           `pulumi:"premiumDataDisks"`
	ProvisioningState *string           `pulumi:"provisioningState"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
	UniqueIdentifier  *string           `pulumi:"uniqueIdentifier"`
}


type LabArgs struct {
	LabStorageType    pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	PremiumDataDisks  pulumi.StringPtrInput
	ProvisioningState pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	UniqueIdentifier  pulumi.StringPtrInput
}

func (LabArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*labArgs)(nil)).Elem()
}

type LabInput interface {
	pulumi.Input

	ToLabOutput() LabOutput
	ToLabOutputWithContext(ctx context.Context) LabOutput
}

func (*Lab) ElementType() reflect.Type {
	return reflect.TypeOf((**Lab)(nil)).Elem()
}

func (i *Lab) ToLabOutput() LabOutput {
	return i.ToLabOutputWithContext(context.Background())
}

func (i *Lab) ToLabOutputWithContext(ctx context.Context) LabOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabOutput)
}

type LabOutput struct{ *pulumi.OutputState }

func (LabOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Lab)(nil)).Elem()
}

func (o LabOutput) ToLabOutput() LabOutput {
	return o
}

func (o LabOutput) ToLabOutputWithContext(ctx context.Context) LabOutput {
	return o
}

func (o LabOutput) ArtifactsStorageAccount() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.ArtifactsStorageAccount }).(pulumi.StringOutput)
}

func (o LabOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.CreatedDate }).(pulumi.StringOutput)
}

func (o LabOutput) DefaultPremiumStorageAccount() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.DefaultPremiumStorageAccount }).(pulumi.StringOutput)
}

func (o LabOutput) DefaultStorageAccount() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.DefaultStorageAccount }).(pulumi.StringOutput)
}

func (o LabOutput) LabStorageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringPtrOutput { return v.LabStorageType }).(pulumi.StringPtrOutput)
}

func (o LabOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LabOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o LabOutput) PremiumDataDiskStorageAccount() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.PremiumDataDiskStorageAccount }).(pulumi.StringOutput)
}

func (o LabOutput) PremiumDataDisks() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringPtrOutput { return v.PremiumDataDisks }).(pulumi.StringPtrOutput)
}

func (o LabOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LabOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LabOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o LabOutput) UniqueIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringPtrOutput { return v.UniqueIdentifier }).(pulumi.StringPtrOutput)
}

func (o LabOutput) VaultName() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.VaultName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LabOutput{})
}
