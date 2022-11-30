


package v20221120preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type APICollection struct {
	pulumi.CustomResourceState

	AdditionalData pulumi.StringMapOutput `pulumi:"additionalData"`
	DisplayName    pulumi.StringPtrOutput `pulumi:"displayName"`
	Name           pulumi.StringOutput    `pulumi:"name"`
	Type           pulumi.StringOutput    `pulumi:"type"`
}


func NewAPICollection(ctx *pulumi.Context,
	name string, args *APICollectionArgs, opts ...pulumi.ResourceOption) (*APICollection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	var resource APICollection
	err := ctx.RegisterResource("azure-native:security/v20221120preview:APICollection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAPICollection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *APICollectionState, opts ...pulumi.ResourceOption) (*APICollection, error) {
	var resource APICollection
	err := ctx.ReadResource("azure-native:security/v20221120preview:APICollection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type apicollectionState struct {
}

type APICollectionState struct {
}

func (APICollectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*apicollectionState)(nil)).Elem()
}

type apicollectionArgs struct {
	ApiCollectionId   *string `pulumi:"apiCollectionId"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServiceName       string  `pulumi:"serviceName"`
}


type APICollectionArgs struct {
	ApiCollectionId   pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
}

func (APICollectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apicollectionArgs)(nil)).Elem()
}

type APICollectionInput interface {
	pulumi.Input

	ToAPICollectionOutput() APICollectionOutput
	ToAPICollectionOutputWithContext(ctx context.Context) APICollectionOutput
}

func (*APICollection) ElementType() reflect.Type {
	return reflect.TypeOf((**APICollection)(nil)).Elem()
}

func (i *APICollection) ToAPICollectionOutput() APICollectionOutput {
	return i.ToAPICollectionOutputWithContext(context.Background())
}

func (i *APICollection) ToAPICollectionOutputWithContext(ctx context.Context) APICollectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(APICollectionOutput)
}

type APICollectionOutput struct{ *pulumi.OutputState }

func (APICollectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**APICollection)(nil)).Elem()
}

func (o APICollectionOutput) ToAPICollectionOutput() APICollectionOutput {
	return o
}

func (o APICollectionOutput) ToAPICollectionOutputWithContext(ctx context.Context) APICollectionOutput {
	return o
}

func (o APICollectionOutput) AdditionalData() pulumi.StringMapOutput {
	return o.ApplyT(func(v *APICollection) pulumi.StringMapOutput { return v.AdditionalData }).(pulumi.StringMapOutput)
}

func (o APICollectionOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *APICollection) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o APICollectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *APICollection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o APICollectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *APICollection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(APICollectionOutput{})
}
