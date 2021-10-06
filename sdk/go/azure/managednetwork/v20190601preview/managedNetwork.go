


package v20190601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ManagedNetwork struct {
	pulumi.CustomResourceState

	Connectivity      ConnectivityCollectionResponseOutput `pulumi:"connectivity"`
	Etag              pulumi.StringOutput                  `pulumi:"etag"`
	Location          pulumi.StringOutput                  `pulumi:"location"`
	Name              pulumi.StringOutput                  `pulumi:"name"`
	ProvisioningState pulumi.StringOutput                  `pulumi:"provisioningState"`
	Scope             ScopeResponsePtrOutput               `pulumi:"scope"`
	Tags              pulumi.StringMapOutput               `pulumi:"tags"`
	Type              pulumi.StringOutput                  `pulumi:"type"`
}


func NewManagedNetwork(ctx *pulumi.Context,
	name string, args *ManagedNetworkArgs, opts ...pulumi.ResourceOption) (*ManagedNetwork, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:managednetwork/v20190601preview:ManagedNetwork"),
		},
		{
			Type: pulumi.String("azure-native:managednetwork:ManagedNetwork"),
		},
		{
			Type: pulumi.String("azure-nextgen:managednetwork:ManagedNetwork"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagedNetwork
	err := ctx.RegisterResource("azure-native:managednetwork/v20190601preview:ManagedNetwork", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetManagedNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedNetworkState, opts ...pulumi.ResourceOption) (*ManagedNetwork, error) {
	var resource ManagedNetwork
	err := ctx.ReadResource("azure-native:managednetwork/v20190601preview:ManagedNetwork", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type managedNetworkState struct {
}

type ManagedNetworkState struct {
}

func (ManagedNetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedNetworkState)(nil)).Elem()
}

type managedNetworkArgs struct {
	Location           *string           `pulumi:"location"`
	ManagedNetworkName *string           `pulumi:"managedNetworkName"`
	ResourceGroupName  string            `pulumi:"resourceGroupName"`
	Scope              *Scope            `pulumi:"scope"`
	Tags               map[string]string `pulumi:"tags"`
}


type ManagedNetworkArgs struct {
	Location           pulumi.StringPtrInput
	ManagedNetworkName pulumi.StringPtrInput
	ResourceGroupName  pulumi.StringInput
	Scope              ScopePtrInput
	Tags               pulumi.StringMapInput
}

func (ManagedNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedNetworkArgs)(nil)).Elem()
}

type ManagedNetworkInput interface {
	pulumi.Input

	ToManagedNetworkOutput() ManagedNetworkOutput
	ToManagedNetworkOutputWithContext(ctx context.Context) ManagedNetworkOutput
}

func (*ManagedNetwork) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedNetwork)(nil))
}

func (i *ManagedNetwork) ToManagedNetworkOutput() ManagedNetworkOutput {
	return i.ToManagedNetworkOutputWithContext(context.Background())
}

func (i *ManagedNetwork) ToManagedNetworkOutputWithContext(ctx context.Context) ManagedNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedNetworkOutput)
}

type ManagedNetworkOutput struct{ *pulumi.OutputState }

func (ManagedNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedNetwork)(nil))
}

func (o ManagedNetworkOutput) ToManagedNetworkOutput() ManagedNetworkOutput {
	return o
}

func (o ManagedNetworkOutput) ToManagedNetworkOutputWithContext(ctx context.Context) ManagedNetworkOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ManagedNetworkOutput{})
}
