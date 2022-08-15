


package v20210301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RestorePointCollection struct {
	pulumi.CustomResourceState

	Location                 pulumi.StringOutput                                     `pulumi:"location"`
	Name                     pulumi.StringOutput                                     `pulumi:"name"`
	ProvisioningState        pulumi.StringOutput                                     `pulumi:"provisioningState"`
	RestorePointCollectionId pulumi.StringOutput                                     `pulumi:"restorePointCollectionId"`
	RestorePoints            RestorePointResponseArrayOutput                         `pulumi:"restorePoints"`
	Source                   RestorePointCollectionSourcePropertiesResponsePtrOutput `pulumi:"source"`
	Tags                     pulumi.StringMapOutput                                  `pulumi:"tags"`
	Type                     pulumi.StringOutput                                     `pulumi:"type"`
}


func NewRestorePointCollection(ctx *pulumi.Context,
	name string, args *RestorePointCollectionArgs, opts ...pulumi.ResourceOption) (*RestorePointCollection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute:RestorePointCollection"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:RestorePointCollection"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210701:RestorePointCollection"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20211101:RestorePointCollection"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220301:RestorePointCollection"),
		},
	})
	opts = append(opts, aliases)
	var resource RestorePointCollection
	err := ctx.RegisterResource("azure-native:compute/v20210301:RestorePointCollection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRestorePointCollection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RestorePointCollectionState, opts ...pulumi.ResourceOption) (*RestorePointCollection, error) {
	var resource RestorePointCollection
	err := ctx.ReadResource("azure-native:compute/v20210301:RestorePointCollection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type restorePointCollectionState struct {
}

type RestorePointCollectionState struct {
}

func (RestorePointCollectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*restorePointCollectionState)(nil)).Elem()
}

type restorePointCollectionArgs struct {
	Location                   *string                                 `pulumi:"location"`
	ResourceGroupName          string                                  `pulumi:"resourceGroupName"`
	RestorePointCollectionName *string                                 `pulumi:"restorePointCollectionName"`
	Source                     *RestorePointCollectionSourceProperties `pulumi:"source"`
	Tags                       map[string]string                       `pulumi:"tags"`
}


type RestorePointCollectionArgs struct {
	Location                   pulumi.StringPtrInput
	ResourceGroupName          pulumi.StringInput
	RestorePointCollectionName pulumi.StringPtrInput
	Source                     RestorePointCollectionSourcePropertiesPtrInput
	Tags                       pulumi.StringMapInput
}

func (RestorePointCollectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*restorePointCollectionArgs)(nil)).Elem()
}

type RestorePointCollectionInput interface {
	pulumi.Input

	ToRestorePointCollectionOutput() RestorePointCollectionOutput
	ToRestorePointCollectionOutputWithContext(ctx context.Context) RestorePointCollectionOutput
}

func (*RestorePointCollection) ElementType() reflect.Type {
	return reflect.TypeOf((**RestorePointCollection)(nil)).Elem()
}

func (i *RestorePointCollection) ToRestorePointCollectionOutput() RestorePointCollectionOutput {
	return i.ToRestorePointCollectionOutputWithContext(context.Background())
}

func (i *RestorePointCollection) ToRestorePointCollectionOutputWithContext(ctx context.Context) RestorePointCollectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestorePointCollectionOutput)
}

type RestorePointCollectionOutput struct{ *pulumi.OutputState }

func (RestorePointCollectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RestorePointCollection)(nil)).Elem()
}

func (o RestorePointCollectionOutput) ToRestorePointCollectionOutput() RestorePointCollectionOutput {
	return o
}

func (o RestorePointCollectionOutput) ToRestorePointCollectionOutputWithContext(ctx context.Context) RestorePointCollectionOutput {
	return o
}

func (o RestorePointCollectionOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *RestorePointCollection) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o RestorePointCollectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RestorePointCollection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RestorePointCollectionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *RestorePointCollection) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o RestorePointCollectionOutput) RestorePointCollectionId() pulumi.StringOutput {
	return o.ApplyT(func(v *RestorePointCollection) pulumi.StringOutput { return v.RestorePointCollectionId }).(pulumi.StringOutput)
}

func (o RestorePointCollectionOutput) RestorePoints() RestorePointResponseArrayOutput {
	return o.ApplyT(func(v *RestorePointCollection) RestorePointResponseArrayOutput { return v.RestorePoints }).(RestorePointResponseArrayOutput)
}

func (o RestorePointCollectionOutput) Source() RestorePointCollectionSourcePropertiesResponsePtrOutput {
	return o.ApplyT(func(v *RestorePointCollection) RestorePointCollectionSourcePropertiesResponsePtrOutput {
		return v.Source
	}).(RestorePointCollectionSourcePropertiesResponsePtrOutput)
}

func (o RestorePointCollectionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *RestorePointCollection) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o RestorePointCollectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RestorePointCollection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RestorePointCollectionOutput{})
}
