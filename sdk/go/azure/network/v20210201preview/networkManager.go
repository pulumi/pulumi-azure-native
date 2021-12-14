


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NetworkManager struct {
	pulumi.CustomResourceState

	Description                 pulumi.StringPtrOutput                                        `pulumi:"description"`
	DisplayName                 pulumi.StringPtrOutput                                        `pulumi:"displayName"`
	Etag                        pulumi.StringOutput                                           `pulumi:"etag"`
	Location                    pulumi.StringPtrOutput                                        `pulumi:"location"`
	Name                        pulumi.StringOutput                                           `pulumi:"name"`
	NetworkManagerScopeAccesses pulumi.StringArrayOutput                                      `pulumi:"networkManagerScopeAccesses"`
	NetworkManagerScopes        NetworkManagerPropertiesResponseNetworkManagerScopesPtrOutput `pulumi:"networkManagerScopes"`
	ProvisioningState           pulumi.StringOutput                                           `pulumi:"provisioningState"`
	SystemData                  SystemDataResponseOutput                                      `pulumi:"systemData"`
	Tags                        pulumi.StringMapOutput                                        `pulumi:"tags"`
	Type                        pulumi.StringOutput                                           `pulumi:"type"`
}


func NewNetworkManager(ctx *pulumi.Context,
	name string, args *NetworkManagerArgs, opts ...pulumi.ResourceOption) (*NetworkManager, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:NetworkManager"),
		},
	})
	opts = append(opts, aliases)
	var resource NetworkManager
	err := ctx.RegisterResource("azure-native:network/v20210201preview:NetworkManager", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNetworkManager(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkManagerState, opts ...pulumi.ResourceOption) (*NetworkManager, error) {
	var resource NetworkManager
	err := ctx.ReadResource("azure-native:network/v20210201preview:NetworkManager", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type networkManagerState struct {
}

type NetworkManagerState struct {
}

func (NetworkManagerState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkManagerState)(nil)).Elem()
}

type networkManagerArgs struct {
	Description                 *string                                       `pulumi:"description"`
	DisplayName                 *string                                       `pulumi:"displayName"`
	Id                          *string                                       `pulumi:"id"`
	Location                    *string                                       `pulumi:"location"`
	NetworkManagerName          *string                                       `pulumi:"networkManagerName"`
	NetworkManagerScopeAccesses []string                                      `pulumi:"networkManagerScopeAccesses"`
	NetworkManagerScopes        *NetworkManagerPropertiesNetworkManagerScopes `pulumi:"networkManagerScopes"`
	ResourceGroupName           string                                        `pulumi:"resourceGroupName"`
	Tags                        map[string]string                             `pulumi:"tags"`
}


type NetworkManagerArgs struct {
	Description                 pulumi.StringPtrInput
	DisplayName                 pulumi.StringPtrInput
	Id                          pulumi.StringPtrInput
	Location                    pulumi.StringPtrInput
	NetworkManagerName          pulumi.StringPtrInput
	NetworkManagerScopeAccesses pulumi.StringArrayInput
	NetworkManagerScopes        NetworkManagerPropertiesNetworkManagerScopesPtrInput
	ResourceGroupName           pulumi.StringInput
	Tags                        pulumi.StringMapInput
}

func (NetworkManagerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkManagerArgs)(nil)).Elem()
}

type NetworkManagerInput interface {
	pulumi.Input

	ToNetworkManagerOutput() NetworkManagerOutput
	ToNetworkManagerOutputWithContext(ctx context.Context) NetworkManagerOutput
}

func (*NetworkManager) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkManager)(nil)).Elem()
}

func (i *NetworkManager) ToNetworkManagerOutput() NetworkManagerOutput {
	return i.ToNetworkManagerOutputWithContext(context.Background())
}

func (i *NetworkManager) ToNetworkManagerOutputWithContext(ctx context.Context) NetworkManagerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkManagerOutput)
}

type NetworkManagerOutput struct{ *pulumi.OutputState }

func (NetworkManagerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkManager)(nil)).Elem()
}

func (o NetworkManagerOutput) ToNetworkManagerOutput() NetworkManagerOutput {
	return o
}

func (o NetworkManagerOutput) ToNetworkManagerOutputWithContext(ctx context.Context) NetworkManagerOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NetworkManagerOutput{})
}
