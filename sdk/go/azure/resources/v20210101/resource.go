


package v20210101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Resource struct {
	pulumi.CustomResourceState

	ExtendedLocation ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	Identity         IdentityResponsePtrOutput         `pulumi:"identity"`
	Kind             pulumi.StringPtrOutput            `pulumi:"kind"`
	Location         pulumi.StringPtrOutput            `pulumi:"location"`
	ManagedBy        pulumi.StringPtrOutput            `pulumi:"managedBy"`
	Name             pulumi.StringOutput               `pulumi:"name"`
	Plan             PlanResponsePtrOutput             `pulumi:"plan"`
	Properties       pulumi.AnyOutput                  `pulumi:"properties"`
	Sku              SkuResponsePtrOutput              `pulumi:"sku"`
	Tags             pulumi.StringMapOutput            `pulumi:"tags"`
	Type             pulumi.StringOutput               `pulumi:"type"`
}


func NewResource(ctx *pulumi.Context,
	name string, args *ResourceArgs, opts ...pulumi.ResourceOption) (*Resource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ParentResourcePath == nil {
		return nil, errors.New("invalid value for required argument 'ParentResourcePath'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceProviderNamespace == nil {
		return nil, errors.New("invalid value for required argument 'ResourceProviderNamespace'")
	}
	if args.ResourceType == nil {
		return nil, errors.New("invalid value for required argument 'ResourceType'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:resources:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20151101:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20160201:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20160701:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20160901:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20170510:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20180201:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20180501:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190301:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190501:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190510:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190701:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190801:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20191001:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20200601:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20200801:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20201001:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20210401:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20220901:Resource"),
		},
	})
	opts = append(opts, aliases)
	var resource Resource
	err := ctx.RegisterResource("azure-native:resources/v20210101:Resource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceState, opts ...pulumi.ResourceOption) (*Resource, error) {
	var resource Resource
	err := ctx.ReadResource("azure-native:resources/v20210101:Resource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type resourceState struct {
}

type ResourceState struct {
}

func (ResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceState)(nil)).Elem()
}

type resourceArgs struct {
	ExtendedLocation          *ExtendedLocation `pulumi:"extendedLocation"`
	Identity                  *Identity         `pulumi:"identity"`
	Kind                      *string           `pulumi:"kind"`
	Location                  *string           `pulumi:"location"`
	ManagedBy                 *string           `pulumi:"managedBy"`
	ParentResourcePath        string            `pulumi:"parentResourcePath"`
	Plan                      *Plan             `pulumi:"plan"`
	Properties                interface{}       `pulumi:"properties"`
	ResourceGroupName         string            `pulumi:"resourceGroupName"`
	ResourceName              *string           `pulumi:"resourceName"`
	ResourceProviderNamespace string            `pulumi:"resourceProviderNamespace"`
	ResourceType              string            `pulumi:"resourceType"`
	Sku                       *Sku              `pulumi:"sku"`
	Tags                      map[string]string `pulumi:"tags"`
}


type ResourceArgs struct {
	ExtendedLocation          ExtendedLocationPtrInput
	Identity                  IdentityPtrInput
	Kind                      pulumi.StringPtrInput
	Location                  pulumi.StringPtrInput
	ManagedBy                 pulumi.StringPtrInput
	ParentResourcePath        pulumi.StringInput
	Plan                      PlanPtrInput
	Properties                pulumi.Input
	ResourceGroupName         pulumi.StringInput
	ResourceName              pulumi.StringPtrInput
	ResourceProviderNamespace pulumi.StringInput
	ResourceType              pulumi.StringInput
	Sku                       SkuPtrInput
	Tags                      pulumi.StringMapInput
}

func (ResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceArgs)(nil)).Elem()
}

type ResourceInput interface {
	pulumi.Input

	ToResourceOutput() ResourceOutput
	ToResourceOutputWithContext(ctx context.Context) ResourceOutput
}

func (*Resource) ElementType() reflect.Type {
	return reflect.TypeOf((**Resource)(nil)).Elem()
}

func (i *Resource) ToResourceOutput() ResourceOutput {
	return i.ToResourceOutputWithContext(context.Background())
}

func (i *Resource) ToResourceOutputWithContext(ctx context.Context) ResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceOutput)
}

type ResourceOutput struct{ *pulumi.OutputState }

func (ResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Resource)(nil)).Elem()
}

func (o ResourceOutput) ToResourceOutput() ResourceOutput {
	return o
}

func (o ResourceOutput) ToResourceOutputWithContext(ctx context.Context) ResourceOutput {
	return o
}

func (o ResourceOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *Resource) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o ResourceOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *Resource) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o ResourceOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Resource) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ResourceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Resource) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ResourceOutput) ManagedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Resource) pulumi.StringPtrOutput { return v.ManagedBy }).(pulumi.StringPtrOutput)
}

func (o ResourceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Resource) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ResourceOutput) Plan() PlanResponsePtrOutput {
	return o.ApplyT(func(v *Resource) PlanResponsePtrOutput { return v.Plan }).(PlanResponsePtrOutput)
}

func (o ResourceOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v *Resource) pulumi.AnyOutput { return v.Properties }).(pulumi.AnyOutput)
}

func (o ResourceOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *Resource) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o ResourceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Resource) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ResourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Resource) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ResourceOutput{})
}
