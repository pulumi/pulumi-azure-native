


package connectedvmwarevsphere

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type HybridIdentityMetadatum struct {
	pulumi.CustomResourceState

	Identity          IdentityResponseOutput   `pulumi:"identity"`
	Name              pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState pulumi.StringOutput      `pulumi:"provisioningState"`
	PublicKey         pulumi.StringPtrOutput   `pulumi:"publicKey"`
	SystemData        SystemDataResponseOutput `pulumi:"systemData"`
	Type              pulumi.StringOutput      `pulumi:"type"`
	VmId              pulumi.StringPtrOutput   `pulumi:"vmId"`
}


func NewHybridIdentityMetadatum(ctx *pulumi.Context,
	name string, args *HybridIdentityMetadatumArgs, opts ...pulumi.ResourceOption) (*HybridIdentityMetadatum, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VirtualMachineName == nil {
		return nil, errors.New("invalid value for required argument 'VirtualMachineName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20201001preview:HybridIdentityMetadatum"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20220110preview:HybridIdentityMetadatum"),
		},
	})
	opts = append(opts, aliases)
	var resource HybridIdentityMetadatum
	err := ctx.RegisterResource("azure-native:connectedvmwarevsphere:HybridIdentityMetadatum", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetHybridIdentityMetadatum(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HybridIdentityMetadatumState, opts ...pulumi.ResourceOption) (*HybridIdentityMetadatum, error) {
	var resource HybridIdentityMetadatum
	err := ctx.ReadResource("azure-native:connectedvmwarevsphere:HybridIdentityMetadatum", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type hybridIdentityMetadatumState struct {
}

type HybridIdentityMetadatumState struct {
}

func (HybridIdentityMetadatumState) ElementType() reflect.Type {
	return reflect.TypeOf((*hybridIdentityMetadatumState)(nil)).Elem()
}

type hybridIdentityMetadatumArgs struct {
	MetadataName       *string `pulumi:"metadataName"`
	PublicKey          *string `pulumi:"publicKey"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
	VirtualMachineName string  `pulumi:"virtualMachineName"`
	VmId               *string `pulumi:"vmId"`
}


type HybridIdentityMetadatumArgs struct {
	MetadataName       pulumi.StringPtrInput
	PublicKey          pulumi.StringPtrInput
	ResourceGroupName  pulumi.StringInput
	VirtualMachineName pulumi.StringInput
	VmId               pulumi.StringPtrInput
}

func (HybridIdentityMetadatumArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hybridIdentityMetadatumArgs)(nil)).Elem()
}

type HybridIdentityMetadatumInput interface {
	pulumi.Input

	ToHybridIdentityMetadatumOutput() HybridIdentityMetadatumOutput
	ToHybridIdentityMetadatumOutputWithContext(ctx context.Context) HybridIdentityMetadatumOutput
}

func (*HybridIdentityMetadatum) ElementType() reflect.Type {
	return reflect.TypeOf((**HybridIdentityMetadatum)(nil)).Elem()
}

func (i *HybridIdentityMetadatum) ToHybridIdentityMetadatumOutput() HybridIdentityMetadatumOutput {
	return i.ToHybridIdentityMetadatumOutputWithContext(context.Background())
}

func (i *HybridIdentityMetadatum) ToHybridIdentityMetadatumOutputWithContext(ctx context.Context) HybridIdentityMetadatumOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HybridIdentityMetadatumOutput)
}

type HybridIdentityMetadatumOutput struct{ *pulumi.OutputState }

func (HybridIdentityMetadatumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HybridIdentityMetadatum)(nil)).Elem()
}

func (o HybridIdentityMetadatumOutput) ToHybridIdentityMetadatumOutput() HybridIdentityMetadatumOutput {
	return o
}

func (o HybridIdentityMetadatumOutput) ToHybridIdentityMetadatumOutputWithContext(ctx context.Context) HybridIdentityMetadatumOutput {
	return o
}

func (o HybridIdentityMetadatumOutput) Identity() IdentityResponseOutput {
	return o.ApplyT(func(v *HybridIdentityMetadatum) IdentityResponseOutput { return v.Identity }).(IdentityResponseOutput)
}

func (o HybridIdentityMetadatumOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridIdentityMetadatum) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o HybridIdentityMetadatumOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridIdentityMetadatum) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o HybridIdentityMetadatumOutput) PublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HybridIdentityMetadatum) pulumi.StringPtrOutput { return v.PublicKey }).(pulumi.StringPtrOutput)
}

func (o HybridIdentityMetadatumOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *HybridIdentityMetadatum) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o HybridIdentityMetadatumOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridIdentityMetadatum) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o HybridIdentityMetadatumOutput) VmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HybridIdentityMetadatum) pulumi.StringPtrOutput { return v.VmId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(HybridIdentityMetadatumOutput{})
}
