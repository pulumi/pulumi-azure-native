


package v20201101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NetworkWatcher struct {
	pulumi.CustomResourceState

	Etag              pulumi.StringOutput    `pulumi:"etag"`
	Location          pulumi.StringPtrOutput `pulumi:"location"`
	Name              pulumi.StringOutput    `pulumi:"name"`
	ProvisioningState pulumi.StringOutput    `pulumi:"provisioningState"`
	Tags              pulumi.StringMapOutput `pulumi:"tags"`
	Type              pulumi.StringOutput    `pulumi:"type"`
}


func NewNetworkWatcher(ctx *pulumi.Context,
	name string, args *NetworkWatcherArgs, opts ...pulumi.ResourceOption) (*NetworkWatcher, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160901:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20161201:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170601:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170801:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:NetworkWatcher"),
		},
	})
	opts = append(opts, aliases)
	var resource NetworkWatcher
	err := ctx.RegisterResource("azure-native:network/v20201101:NetworkWatcher", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNetworkWatcher(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkWatcherState, opts ...pulumi.ResourceOption) (*NetworkWatcher, error) {
	var resource NetworkWatcher
	err := ctx.ReadResource("azure-native:network/v20201101:NetworkWatcher", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type networkWatcherState struct {
}

type NetworkWatcherState struct {
}

func (NetworkWatcherState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkWatcherState)(nil)).Elem()
}

type networkWatcherArgs struct {
	Id                 *string           `pulumi:"id"`
	Location           *string           `pulumi:"location"`
	NetworkWatcherName *string           `pulumi:"networkWatcherName"`
	ResourceGroupName  string            `pulumi:"resourceGroupName"`
	Tags               map[string]string `pulumi:"tags"`
}


type NetworkWatcherArgs struct {
	Id                 pulumi.StringPtrInput
	Location           pulumi.StringPtrInput
	NetworkWatcherName pulumi.StringPtrInput
	ResourceGroupName  pulumi.StringInput
	Tags               pulumi.StringMapInput
}

func (NetworkWatcherArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkWatcherArgs)(nil)).Elem()
}

type NetworkWatcherInput interface {
	pulumi.Input

	ToNetworkWatcherOutput() NetworkWatcherOutput
	ToNetworkWatcherOutputWithContext(ctx context.Context) NetworkWatcherOutput
}

func (*NetworkWatcher) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkWatcher)(nil)).Elem()
}

func (i *NetworkWatcher) ToNetworkWatcherOutput() NetworkWatcherOutput {
	return i.ToNetworkWatcherOutputWithContext(context.Background())
}

func (i *NetworkWatcher) ToNetworkWatcherOutputWithContext(ctx context.Context) NetworkWatcherOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkWatcherOutput)
}

type NetworkWatcherOutput struct{ *pulumi.OutputState }

func (NetworkWatcherOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkWatcher)(nil)).Elem()
}

func (o NetworkWatcherOutput) ToNetworkWatcherOutput() NetworkWatcherOutput {
	return o
}

func (o NetworkWatcherOutput) ToNetworkWatcherOutputWithContext(ctx context.Context) NetworkWatcherOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NetworkWatcherOutput{})
}
