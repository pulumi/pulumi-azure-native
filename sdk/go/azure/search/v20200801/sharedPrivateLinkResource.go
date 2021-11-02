


package v20200801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SharedPrivateLinkResource struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                               `pulumi:"name"`
	Properties SharedPrivateLinkResourcePropertiesResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput                               `pulumi:"type"`
}


func NewSharedPrivateLinkResource(ctx *pulumi.Context,
	name string, args *SharedPrivateLinkResourceArgs, opts ...pulumi.ResourceOption) (*SharedPrivateLinkResource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SearchServiceName == nil {
		return nil, errors.New("invalid value for required argument 'SearchServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:search/v20200801:SharedPrivateLinkResource"),
		},
		{
			Type: pulumi.String("azure-native:search:SharedPrivateLinkResource"),
		},
		{
			Type: pulumi.String("azure-nextgen:search:SharedPrivateLinkResource"),
		},
		{
			Type: pulumi.String("azure-native:search/v20200801preview:SharedPrivateLinkResource"),
		},
		{
			Type: pulumi.String("azure-nextgen:search/v20200801preview:SharedPrivateLinkResource"),
		},
		{
			Type: pulumi.String("azure-native:search/v20210401preview:SharedPrivateLinkResource"),
		},
		{
			Type: pulumi.String("azure-nextgen:search/v20210401preview:SharedPrivateLinkResource"),
		},
	})
	opts = append(opts, aliases)
	var resource SharedPrivateLinkResource
	err := ctx.RegisterResource("azure-native:search/v20200801:SharedPrivateLinkResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSharedPrivateLinkResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SharedPrivateLinkResourceState, opts ...pulumi.ResourceOption) (*SharedPrivateLinkResource, error) {
	var resource SharedPrivateLinkResource
	err := ctx.ReadResource("azure-native:search/v20200801:SharedPrivateLinkResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type sharedPrivateLinkResourceState struct {
}

type SharedPrivateLinkResourceState struct {
}

func (SharedPrivateLinkResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*sharedPrivateLinkResourceState)(nil)).Elem()
}

type sharedPrivateLinkResourceArgs struct {
	Properties                    *SharedPrivateLinkResourceProperties `pulumi:"properties"`
	ResourceGroupName             string                               `pulumi:"resourceGroupName"`
	SearchServiceName             string                               `pulumi:"searchServiceName"`
	SharedPrivateLinkResourceName *string                              `pulumi:"sharedPrivateLinkResourceName"`
}


type SharedPrivateLinkResourceArgs struct {
	Properties                    SharedPrivateLinkResourcePropertiesPtrInput
	ResourceGroupName             pulumi.StringInput
	SearchServiceName             pulumi.StringInput
	SharedPrivateLinkResourceName pulumi.StringPtrInput
}

func (SharedPrivateLinkResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sharedPrivateLinkResourceArgs)(nil)).Elem()
}

type SharedPrivateLinkResourceInput interface {
	pulumi.Input

	ToSharedPrivateLinkResourceOutput() SharedPrivateLinkResourceOutput
	ToSharedPrivateLinkResourceOutputWithContext(ctx context.Context) SharedPrivateLinkResourceOutput
}

func (*SharedPrivateLinkResource) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedPrivateLinkResource)(nil))
}

func (i *SharedPrivateLinkResource) ToSharedPrivateLinkResourceOutput() SharedPrivateLinkResourceOutput {
	return i.ToSharedPrivateLinkResourceOutputWithContext(context.Background())
}

func (i *SharedPrivateLinkResource) ToSharedPrivateLinkResourceOutputWithContext(ctx context.Context) SharedPrivateLinkResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedPrivateLinkResourceOutput)
}

type SharedPrivateLinkResourceOutput struct{ *pulumi.OutputState }

func (SharedPrivateLinkResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedPrivateLinkResource)(nil))
}

func (o SharedPrivateLinkResourceOutput) ToSharedPrivateLinkResourceOutput() SharedPrivateLinkResourceOutput {
	return o
}

func (o SharedPrivateLinkResourceOutput) ToSharedPrivateLinkResourceOutputWithContext(ctx context.Context) SharedPrivateLinkResourceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SharedPrivateLinkResourceOutput{})
}
