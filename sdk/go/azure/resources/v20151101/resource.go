


package v20151101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Resource struct {
	pulumi.CustomResourceState

	Location   pulumi.StringOutput    `pulumi:"location"`
	Name       pulumi.StringOutput    `pulumi:"name"`
	Plan       PlanResponsePtrOutput  `pulumi:"plan"`
	Properties pulumi.AnyOutput       `pulumi:"properties"`
	Tags       pulumi.StringMapOutput `pulumi:"tags"`
	Type       pulumi.StringOutput    `pulumi:"type"`
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
			Type: pulumi.String("azure-nextgen:resources/v20151101:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources:Resource"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20160201:Resource"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources/v20160201:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20160701:Resource"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources/v20160701:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20160901:Resource"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources/v20160901:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20170510:Resource"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources/v20170510:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20180201:Resource"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources/v20180201:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20180501:Resource"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources/v20180501:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190301:Resource"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources/v20190301:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190501:Resource"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources/v20190501:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190510:Resource"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources/v20190510:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190701:Resource"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources/v20190701:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190801:Resource"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources/v20190801:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20191001:Resource"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources/v20191001:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20200601:Resource"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources/v20200601:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20200801:Resource"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources/v20200801:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20201001:Resource"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources/v20201001:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20210101:Resource"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources/v20210101:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20210401:Resource"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources/v20210401:Resource"),
		},
	})
	opts = append(opts, aliases)
	var resource Resource
	err := ctx.RegisterResource("azure-native:resources/v20151101:Resource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceState, opts ...pulumi.ResourceOption) (*Resource, error) {
	var resource Resource
	err := ctx.ReadResource("azure-native:resources/v20151101:Resource", name, id, state, &resource, opts...)
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
	Location                  *string           `pulumi:"location"`
	ParentResourcePath        string            `pulumi:"parentResourcePath"`
	Plan                      *Plan             `pulumi:"plan"`
	Properties                interface{}       `pulumi:"properties"`
	ResourceGroupName         string            `pulumi:"resourceGroupName"`
	ResourceName              *string           `pulumi:"resourceName"`
	ResourceProviderNamespace string            `pulumi:"resourceProviderNamespace"`
	ResourceType              string            `pulumi:"resourceType"`
	Tags                      map[string]string `pulumi:"tags"`
}


type ResourceArgs struct {
	Location                  pulumi.StringPtrInput
	ParentResourcePath        pulumi.StringInput
	Plan                      PlanPtrInput
	Properties                pulumi.Input
	ResourceGroupName         pulumi.StringInput
	ResourceName              pulumi.StringPtrInput
	ResourceProviderNamespace pulumi.StringInput
	ResourceType              pulumi.StringInput
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
	return reflect.TypeOf((*Resource)(nil))
}

func (i *Resource) ToResourceOutput() ResourceOutput {
	return i.ToResourceOutputWithContext(context.Background())
}

func (i *Resource) ToResourceOutputWithContext(ctx context.Context) ResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceOutput)
}

type ResourceOutput struct{ *pulumi.OutputState }

func (ResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Resource)(nil))
}

func (o ResourceOutput) ToResourceOutput() ResourceOutput {
	return o
}

func (o ResourceOutput) ToResourceOutputWithContext(ctx context.Context) ResourceOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceInput)(nil)).Elem(), &Resource{})
	pulumi.RegisterOutputType(ResourceOutput{})
}
