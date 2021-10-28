


package authorization

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ResourceManagementPrivateLink struct {
	pulumi.CustomResourceState

	Location   pulumi.StringPtrOutput                                         `pulumi:"location"`
	Name       pulumi.StringOutput                                            `pulumi:"name"`
	Properties ResourceManagementPrivateLinkEndpointConnectionsResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput                                            `pulumi:"type"`
}


func NewResourceManagementPrivateLink(ctx *pulumi.Context,
	name string, args *ResourceManagementPrivateLinkArgs, opts ...pulumi.ResourceOption) (*ResourceManagementPrivateLink, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:authorization:ResourceManagementPrivateLink"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20200501:ResourceManagementPrivateLink"),
		},
		{
			Type: pulumi.String("azure-nextgen:authorization/v20200501:ResourceManagementPrivateLink"),
		},
	})
	opts = append(opts, aliases)
	var resource ResourceManagementPrivateLink
	err := ctx.RegisterResource("azure-native:authorization:ResourceManagementPrivateLink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetResourceManagementPrivateLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceManagementPrivateLinkState, opts ...pulumi.ResourceOption) (*ResourceManagementPrivateLink, error) {
	var resource ResourceManagementPrivateLink
	err := ctx.ReadResource("azure-native:authorization:ResourceManagementPrivateLink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type resourceManagementPrivateLinkState struct {
}

type ResourceManagementPrivateLinkState struct {
}

func (ResourceManagementPrivateLinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceManagementPrivateLinkState)(nil)).Elem()
}

type resourceManagementPrivateLinkArgs struct {
	Location          *string `pulumi:"location"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	RmplName          *string `pulumi:"rmplName"`
}


type ResourceManagementPrivateLinkArgs struct {
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	RmplName          pulumi.StringPtrInput
}

func (ResourceManagementPrivateLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceManagementPrivateLinkArgs)(nil)).Elem()
}

type ResourceManagementPrivateLinkInput interface {
	pulumi.Input

	ToResourceManagementPrivateLinkOutput() ResourceManagementPrivateLinkOutput
	ToResourceManagementPrivateLinkOutputWithContext(ctx context.Context) ResourceManagementPrivateLinkOutput
}

func (*ResourceManagementPrivateLink) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceManagementPrivateLink)(nil))
}

func (i *ResourceManagementPrivateLink) ToResourceManagementPrivateLinkOutput() ResourceManagementPrivateLinkOutput {
	return i.ToResourceManagementPrivateLinkOutputWithContext(context.Background())
}

func (i *ResourceManagementPrivateLink) ToResourceManagementPrivateLinkOutputWithContext(ctx context.Context) ResourceManagementPrivateLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceManagementPrivateLinkOutput)
}

type ResourceManagementPrivateLinkOutput struct{ *pulumi.OutputState }

func (ResourceManagementPrivateLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceManagementPrivateLink)(nil))
}

func (o ResourceManagementPrivateLinkOutput) ToResourceManagementPrivateLinkOutput() ResourceManagementPrivateLinkOutput {
	return o
}

func (o ResourceManagementPrivateLinkOutput) ToResourceManagementPrivateLinkOutputWithContext(ctx context.Context) ResourceManagementPrivateLinkOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceManagementPrivateLinkInput)(nil)).Elem(), &ResourceManagementPrivateLink{})
	pulumi.RegisterOutputType(ResourceManagementPrivateLinkOutput{})
}
