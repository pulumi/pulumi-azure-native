


package v20220101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NetworkManager struct {
	pulumi.CustomResourceState

	Description                 pulumi.StringPtrOutput                                     `pulumi:"description"`
	Etag                        pulumi.StringOutput                                        `pulumi:"etag"`
	Location                    pulumi.StringPtrOutput                                     `pulumi:"location"`
	Name                        pulumi.StringOutput                                        `pulumi:"name"`
	NetworkManagerScopeAccesses pulumi.StringArrayOutput                                   `pulumi:"networkManagerScopeAccesses"`
	NetworkManagerScopes        NetworkManagerPropertiesResponseNetworkManagerScopesOutput `pulumi:"networkManagerScopes"`
	ProvisioningState           pulumi.StringOutput                                        `pulumi:"provisioningState"`
	SystemData                  SystemDataResponseOutput                                   `pulumi:"systemData"`
	Tags                        pulumi.StringMapOutput                                     `pulumi:"tags"`
	Type                        pulumi.StringOutput                                        `pulumi:"type"`
}


func NewNetworkManager(ctx *pulumi.Context,
	name string, args *NetworkManagerArgs, opts ...pulumi.ResourceOption) (*NetworkManager, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkManagerScopeAccesses == nil {
		return nil, errors.New("invalid value for required argument 'NetworkManagerScopeAccesses'")
	}
	if args.NetworkManagerScopes == nil {
		return nil, errors.New("invalid value for required argument 'NetworkManagerScopes'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:NetworkManager"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201preview:NetworkManager"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501preview:NetworkManager"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220201preview:NetworkManager"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220401preview:NetworkManager"),
		},
	})
	opts = append(opts, aliases)
	var resource NetworkManager
	err := ctx.RegisterResource("azure-native:network/v20220101:NetworkManager", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNetworkManager(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkManagerState, opts ...pulumi.ResourceOption) (*NetworkManager, error) {
	var resource NetworkManager
	err := ctx.ReadResource("azure-native:network/v20220101:NetworkManager", name, id, state, &resource, opts...)
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
	Description                 *string                                      `pulumi:"description"`
	Id                          *string                                      `pulumi:"id"`
	Location                    *string                                      `pulumi:"location"`
	NetworkManagerName          *string                                      `pulumi:"networkManagerName"`
	NetworkManagerScopeAccesses []string                                     `pulumi:"networkManagerScopeAccesses"`
	NetworkManagerScopes        NetworkManagerPropertiesNetworkManagerScopes `pulumi:"networkManagerScopes"`
	ResourceGroupName           string                                       `pulumi:"resourceGroupName"`
	Tags                        map[string]string                            `pulumi:"tags"`
}


type NetworkManagerArgs struct {
	Description                 pulumi.StringPtrInput
	Id                          pulumi.StringPtrInput
	Location                    pulumi.StringPtrInput
	NetworkManagerName          pulumi.StringPtrInput
	NetworkManagerScopeAccesses pulumi.StringArrayInput
	NetworkManagerScopes        NetworkManagerPropertiesNetworkManagerScopesInput
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

func (o NetworkManagerOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkManager) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o NetworkManagerOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkManager) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o NetworkManagerOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkManager) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o NetworkManagerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkManager) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NetworkManagerOutput) NetworkManagerScopeAccesses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkManager) pulumi.StringArrayOutput { return v.NetworkManagerScopeAccesses }).(pulumi.StringArrayOutput)
}

func (o NetworkManagerOutput) NetworkManagerScopes() NetworkManagerPropertiesResponseNetworkManagerScopesOutput {
	return o.ApplyT(func(v *NetworkManager) NetworkManagerPropertiesResponseNetworkManagerScopesOutput {
		return v.NetworkManagerScopes
	}).(NetworkManagerPropertiesResponseNetworkManagerScopesOutput)
}

func (o NetworkManagerOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkManager) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o NetworkManagerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *NetworkManager) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o NetworkManagerOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkManager) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o NetworkManagerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkManager) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkManagerOutput{})
}
