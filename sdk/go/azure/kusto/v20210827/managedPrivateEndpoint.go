


package v20210827

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ManagedPrivateEndpoint struct {
	pulumi.CustomResourceState

	GroupId                   pulumi.StringOutput      `pulumi:"groupId"`
	Name                      pulumi.StringOutput      `pulumi:"name"`
	PrivateLinkResourceId     pulumi.StringOutput      `pulumi:"privateLinkResourceId"`
	PrivateLinkResourceRegion pulumi.StringPtrOutput   `pulumi:"privateLinkResourceRegion"`
	ProvisioningState         pulumi.StringOutput      `pulumi:"provisioningState"`
	RequestMessage            pulumi.StringPtrOutput   `pulumi:"requestMessage"`
	SystemData                SystemDataResponseOutput `pulumi:"systemData"`
	Type                      pulumi.StringOutput      `pulumi:"type"`
}


func NewManagedPrivateEndpoint(ctx *pulumi.Context,
	name string, args *ManagedPrivateEndpointArgs, opts ...pulumi.ResourceOption) (*ManagedPrivateEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
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
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:kusto:ManagedPrivateEndpoint"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagedPrivateEndpoint
	err := ctx.RegisterResource("azure-native:kusto/v20210827:ManagedPrivateEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetManagedPrivateEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedPrivateEndpointState, opts ...pulumi.ResourceOption) (*ManagedPrivateEndpoint, error) {
	var resource ManagedPrivateEndpoint
	err := ctx.ReadResource("azure-native:kusto/v20210827:ManagedPrivateEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type managedPrivateEndpointState struct {
}

type ManagedPrivateEndpointState struct {
}

func (ManagedPrivateEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedPrivateEndpointState)(nil)).Elem()
}

type managedPrivateEndpointArgs struct {
	ClusterName                string  `pulumi:"clusterName"`
	GroupId                    string  `pulumi:"groupId"`
	ManagedPrivateEndpointName *string `pulumi:"managedPrivateEndpointName"`
	PrivateLinkResourceId      string  `pulumi:"privateLinkResourceId"`
	PrivateLinkResourceRegion  *string `pulumi:"privateLinkResourceRegion"`
	RequestMessage             *string `pulumi:"requestMessage"`
	ResourceGroupName          string  `pulumi:"resourceGroupName"`
}


type ManagedPrivateEndpointArgs struct {
	ClusterName                pulumi.StringInput
	GroupId                    pulumi.StringInput
	ManagedPrivateEndpointName pulumi.StringPtrInput
	PrivateLinkResourceId      pulumi.StringInput
	PrivateLinkResourceRegion  pulumi.StringPtrInput
	RequestMessage             pulumi.StringPtrInput
	ResourceGroupName          pulumi.StringInput
}

func (ManagedPrivateEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedPrivateEndpointArgs)(nil)).Elem()
}

type ManagedPrivateEndpointInput interface {
	pulumi.Input

	ToManagedPrivateEndpointOutput() ManagedPrivateEndpointOutput
	ToManagedPrivateEndpointOutputWithContext(ctx context.Context) ManagedPrivateEndpointOutput
}

func (*ManagedPrivateEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedPrivateEndpoint)(nil))
}

func (i *ManagedPrivateEndpoint) ToManagedPrivateEndpointOutput() ManagedPrivateEndpointOutput {
	return i.ToManagedPrivateEndpointOutputWithContext(context.Background())
}

func (i *ManagedPrivateEndpoint) ToManagedPrivateEndpointOutputWithContext(ctx context.Context) ManagedPrivateEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedPrivateEndpointOutput)
}

type ManagedPrivateEndpointOutput struct{ *pulumi.OutputState }

func (ManagedPrivateEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedPrivateEndpoint)(nil))
}

func (o ManagedPrivateEndpointOutput) ToManagedPrivateEndpointOutput() ManagedPrivateEndpointOutput {
	return o
}

func (o ManagedPrivateEndpointOutput) ToManagedPrivateEndpointOutputWithContext(ctx context.Context) ManagedPrivateEndpointOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ManagedPrivateEndpointOutput{})
}
