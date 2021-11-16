


package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SignalRSharedPrivateLinkResource struct {
	pulumi.CustomResourceState

	GroupId               pulumi.StringOutput      `pulumi:"groupId"`
	Name                  pulumi.StringOutput      `pulumi:"name"`
	PrivateLinkResourceId pulumi.StringOutput      `pulumi:"privateLinkResourceId"`
	ProvisioningState     pulumi.StringOutput      `pulumi:"provisioningState"`
	RequestMessage        pulumi.StringPtrOutput   `pulumi:"requestMessage"`
	Status                pulumi.StringOutput      `pulumi:"status"`
	SystemData            SystemDataResponseOutput `pulumi:"systemData"`
	Type                  pulumi.StringOutput      `pulumi:"type"`
}


func NewSignalRSharedPrivateLinkResource(ctx *pulumi.Context,
	name string, args *SignalRSharedPrivateLinkResourceArgs, opts ...pulumi.ResourceOption) (*SignalRSharedPrivateLinkResource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	if args.PrivateLinkResourceId == nil {
		return nil, errors.New("invalid value for required argument 'PrivateLinkResourceId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:signalrservice:SignalRSharedPrivateLinkResource"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20210601preview:SignalRSharedPrivateLinkResource"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20210901preview:SignalRSharedPrivateLinkResource"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20211001:SignalRSharedPrivateLinkResource"),
		},
	})
	opts = append(opts, aliases)
	var resource SignalRSharedPrivateLinkResource
	err := ctx.RegisterResource("azure-native:signalrservice/v20210401preview:SignalRSharedPrivateLinkResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSignalRSharedPrivateLinkResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SignalRSharedPrivateLinkResourceState, opts ...pulumi.ResourceOption) (*SignalRSharedPrivateLinkResource, error) {
	var resource SignalRSharedPrivateLinkResource
	err := ctx.ReadResource("azure-native:signalrservice/v20210401preview:SignalRSharedPrivateLinkResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type signalRSharedPrivateLinkResourceState struct {
}

type SignalRSharedPrivateLinkResourceState struct {
}

func (SignalRSharedPrivateLinkResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*signalRSharedPrivateLinkResourceState)(nil)).Elem()
}

type signalRSharedPrivateLinkResourceArgs struct {
	GroupId                       string  `pulumi:"groupId"`
	PrivateLinkResourceId         string  `pulumi:"privateLinkResourceId"`
	RequestMessage                *string `pulumi:"requestMessage"`
	ResourceGroupName             string  `pulumi:"resourceGroupName"`
	ResourceName                  string  `pulumi:"resourceName"`
	SharedPrivateLinkResourceName *string `pulumi:"sharedPrivateLinkResourceName"`
}


type SignalRSharedPrivateLinkResourceArgs struct {
	GroupId                       pulumi.StringInput
	PrivateLinkResourceId         pulumi.StringInput
	RequestMessage                pulumi.StringPtrInput
	ResourceGroupName             pulumi.StringInput
	ResourceName                  pulumi.StringInput
	SharedPrivateLinkResourceName pulumi.StringPtrInput
}

func (SignalRSharedPrivateLinkResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*signalRSharedPrivateLinkResourceArgs)(nil)).Elem()
}

type SignalRSharedPrivateLinkResourceInput interface {
	pulumi.Input

	ToSignalRSharedPrivateLinkResourceOutput() SignalRSharedPrivateLinkResourceOutput
	ToSignalRSharedPrivateLinkResourceOutputWithContext(ctx context.Context) SignalRSharedPrivateLinkResourceOutput
}

func (*SignalRSharedPrivateLinkResource) ElementType() reflect.Type {
	return reflect.TypeOf((*SignalRSharedPrivateLinkResource)(nil))
}

func (i *SignalRSharedPrivateLinkResource) ToSignalRSharedPrivateLinkResourceOutput() SignalRSharedPrivateLinkResourceOutput {
	return i.ToSignalRSharedPrivateLinkResourceOutputWithContext(context.Background())
}

func (i *SignalRSharedPrivateLinkResource) ToSignalRSharedPrivateLinkResourceOutputWithContext(ctx context.Context) SignalRSharedPrivateLinkResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SignalRSharedPrivateLinkResourceOutput)
}

type SignalRSharedPrivateLinkResourceOutput struct{ *pulumi.OutputState }

func (SignalRSharedPrivateLinkResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SignalRSharedPrivateLinkResource)(nil))
}

func (o SignalRSharedPrivateLinkResourceOutput) ToSignalRSharedPrivateLinkResourceOutput() SignalRSharedPrivateLinkResourceOutput {
	return o
}

func (o SignalRSharedPrivateLinkResourceOutput) ToSignalRSharedPrivateLinkResourceOutputWithContext(ctx context.Context) SignalRSharedPrivateLinkResourceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SignalRSharedPrivateLinkResourceOutput{})
}
