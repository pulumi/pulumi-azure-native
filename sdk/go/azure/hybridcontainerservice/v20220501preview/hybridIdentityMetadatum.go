


package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type HybridIdentityMetadatum struct {
	pulumi.CustomResourceState

	Identity          ProvisionedClusterIdentityResponsePtrOutput `pulumi:"identity"`
	Name              pulumi.StringOutput                         `pulumi:"name"`
	ProvisioningState pulumi.StringOutput                         `pulumi:"provisioningState"`
	PublicKey         pulumi.StringPtrOutput                      `pulumi:"publicKey"`
	ResourceUid       pulumi.StringPtrOutput                      `pulumi:"resourceUid"`
	SystemData        SystemDataResponseOutput                    `pulumi:"systemData"`
	Type              pulumi.StringOutput                         `pulumi:"type"`
}


func NewHybridIdentityMetadatum(ctx *pulumi.Context,
	name string, args *HybridIdentityMetadatumArgs, opts ...pulumi.ResourceOption) (*HybridIdentityMetadatum, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProvisionedClustersName == nil {
		return nil, errors.New("invalid value for required argument 'ProvisionedClustersName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource HybridIdentityMetadatum
	err := ctx.RegisterResource("azure-native:hybridcontainerservice/v20220501preview:HybridIdentityMetadatum", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetHybridIdentityMetadatum(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HybridIdentityMetadatumState, opts ...pulumi.ResourceOption) (*HybridIdentityMetadatum, error) {
	var resource HybridIdentityMetadatum
	err := ctx.ReadResource("azure-native:hybridcontainerservice/v20220501preview:HybridIdentityMetadatum", name, id, state, &resource, opts...)
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
	HybridIdentityMetadataResourceName *string                     `pulumi:"hybridIdentityMetadataResourceName"`
	Identity                           *ProvisionedClusterIdentity `pulumi:"identity"`
	ProvisionedClustersName            string                      `pulumi:"provisionedClustersName"`
	PublicKey                          *string                     `pulumi:"publicKey"`
	ResourceGroupName                  string                      `pulumi:"resourceGroupName"`
	ResourceUid                        *string                     `pulumi:"resourceUid"`
}


type HybridIdentityMetadatumArgs struct {
	HybridIdentityMetadataResourceName pulumi.StringPtrInput
	Identity                           ProvisionedClusterIdentityPtrInput
	ProvisionedClustersName            pulumi.StringInput
	PublicKey                          pulumi.StringPtrInput
	ResourceGroupName                  pulumi.StringInput
	ResourceUid                        pulumi.StringPtrInput
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

func (o HybridIdentityMetadatumOutput) Identity() ProvisionedClusterIdentityResponsePtrOutput {
	return o.ApplyT(func(v *HybridIdentityMetadatum) ProvisionedClusterIdentityResponsePtrOutput { return v.Identity }).(ProvisionedClusterIdentityResponsePtrOutput)
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

func (o HybridIdentityMetadatumOutput) ResourceUid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HybridIdentityMetadatum) pulumi.StringPtrOutput { return v.ResourceUid }).(pulumi.StringPtrOutput)
}

func (o HybridIdentityMetadatumOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *HybridIdentityMetadatum) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o HybridIdentityMetadatumOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridIdentityMetadatum) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(HybridIdentityMetadatumOutput{})
}
