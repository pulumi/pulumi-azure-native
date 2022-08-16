


package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RestorePoint struct {
	pulumi.CustomResourceState

	ConsistencyMode   pulumi.StringOutput                      `pulumi:"consistencyMode"`
	ExcludeDisks      ApiEntityReferenceResponseArrayOutput    `pulumi:"excludeDisks"`
	Name              pulumi.StringOutput                      `pulumi:"name"`
	ProvisioningState pulumi.StringOutput                      `pulumi:"provisioningState"`
	SourceMetadata    RestorePointSourceMetadataResponseOutput `pulumi:"sourceMetadata"`
	TimeCreated       pulumi.StringPtrOutput                   `pulumi:"timeCreated"`
	Type              pulumi.StringOutput                      `pulumi:"type"`
}


func NewRestorePoint(ctx *pulumi.Context,
	name string, args *RestorePointArgs, opts ...pulumi.ResourceOption) (*RestorePoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.RestorePointCollectionName == nil {
		return nil, errors.New("invalid value for required argument 'RestorePointCollectionName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute/v20210301:RestorePoint"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:RestorePoint"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210701:RestorePoint"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20211101:RestorePoint"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220301:RestorePoint"),
		},
	})
	opts = append(opts, aliases)
	var resource RestorePoint
	err := ctx.RegisterResource("azure-native:compute:RestorePoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRestorePoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RestorePointState, opts ...pulumi.ResourceOption) (*RestorePoint, error) {
	var resource RestorePoint
	err := ctx.ReadResource("azure-native:compute:RestorePoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type restorePointState struct {
}

type RestorePointState struct {
}

func (RestorePointState) ElementType() reflect.Type {
	return reflect.TypeOf((*restorePointState)(nil)).Elem()
}

type restorePointArgs struct {
	ExcludeDisks               []ApiEntityReference `pulumi:"excludeDisks"`
	ResourceGroupName          string               `pulumi:"resourceGroupName"`
	RestorePointCollectionName string               `pulumi:"restorePointCollectionName"`
	RestorePointName           *string              `pulumi:"restorePointName"`
	TimeCreated                *string              `pulumi:"timeCreated"`
}


type RestorePointArgs struct {
	ExcludeDisks               ApiEntityReferenceArrayInput
	ResourceGroupName          pulumi.StringInput
	RestorePointCollectionName pulumi.StringInput
	RestorePointName           pulumi.StringPtrInput
	TimeCreated                pulumi.StringPtrInput
}

func (RestorePointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*restorePointArgs)(nil)).Elem()
}

type RestorePointInput interface {
	pulumi.Input

	ToRestorePointOutput() RestorePointOutput
	ToRestorePointOutputWithContext(ctx context.Context) RestorePointOutput
}

func (*RestorePoint) ElementType() reflect.Type {
	return reflect.TypeOf((**RestorePoint)(nil)).Elem()
}

func (i *RestorePoint) ToRestorePointOutput() RestorePointOutput {
	return i.ToRestorePointOutputWithContext(context.Background())
}

func (i *RestorePoint) ToRestorePointOutputWithContext(ctx context.Context) RestorePointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestorePointOutput)
}

type RestorePointOutput struct{ *pulumi.OutputState }

func (RestorePointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RestorePoint)(nil)).Elem()
}

func (o RestorePointOutput) ToRestorePointOutput() RestorePointOutput {
	return o
}

func (o RestorePointOutput) ToRestorePointOutputWithContext(ctx context.Context) RestorePointOutput {
	return o
}

func (o RestorePointOutput) ConsistencyMode() pulumi.StringOutput {
	return o.ApplyT(func(v *RestorePoint) pulumi.StringOutput { return v.ConsistencyMode }).(pulumi.StringOutput)
}

func (o RestorePointOutput) ExcludeDisks() ApiEntityReferenceResponseArrayOutput {
	return o.ApplyT(func(v *RestorePoint) ApiEntityReferenceResponseArrayOutput { return v.ExcludeDisks }).(ApiEntityReferenceResponseArrayOutput)
}

func (o RestorePointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RestorePoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RestorePointOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *RestorePoint) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o RestorePointOutput) SourceMetadata() RestorePointSourceMetadataResponseOutput {
	return o.ApplyT(func(v *RestorePoint) RestorePointSourceMetadataResponseOutput { return v.SourceMetadata }).(RestorePointSourceMetadataResponseOutput)
}

func (o RestorePointOutput) TimeCreated() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RestorePoint) pulumi.StringPtrOutput { return v.TimeCreated }).(pulumi.StringPtrOutput)
}

func (o RestorePointOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RestorePoint) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RestorePointOutput{})
}
