


package v20180701

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type ServerEndpoint struct {
	pulumi.CustomResourceState

	CloudTiering           pulumi.StringPtrOutput `pulumi:"cloudTiering"`
	FriendlyName           pulumi.StringPtrOutput `pulumi:"friendlyName"`
	LastOperationName      pulumi.StringPtrOutput `pulumi:"lastOperationName"`
	LastWorkflowId         pulumi.StringPtrOutput `pulumi:"lastWorkflowId"`
	Name                   pulumi.StringOutput    `pulumi:"name"`
	ProvisioningState      pulumi.StringPtrOutput `pulumi:"provisioningState"`
	ServerLocalPath        pulumi.StringPtrOutput `pulumi:"serverLocalPath"`
	ServerResourceId       pulumi.StringPtrOutput `pulumi:"serverResourceId"`
	SyncStatus             pulumi.AnyOutput       `pulumi:"syncStatus"`
	TierFilesOlderThanDays pulumi.IntPtrOutput    `pulumi:"tierFilesOlderThanDays"`
	Type                   pulumi.StringOutput    `pulumi:"type"`
	VolumeFreeSpacePercent pulumi.IntPtrOutput    `pulumi:"volumeFreeSpacePercent"`
}


func NewServerEndpoint(ctx *pulumi.Context,
	name string, args *ServerEndpointArgs, opts ...pulumi.ResourceOption) (*ServerEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StorageSyncServiceName == nil {
		return nil, errors.New("invalid value for required argument 'StorageSyncServiceName'")
	}
	if args.SyncGroupName == nil {
		return nil, errors.New("invalid value for required argument 'SyncGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storagesync:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20170605preview:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20180402:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20181001:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20190201:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20190301:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20190601:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20191001:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20200301:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20200901:ServerEndpoint"),
		},
	})
	opts = append(opts, aliases)
	var resource ServerEndpoint
	err := ctx.RegisterResource("azure-native:storagesync/v20180701:ServerEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetServerEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerEndpointState, opts ...pulumi.ResourceOption) (*ServerEndpoint, error) {
	var resource ServerEndpoint
	err := ctx.ReadResource("azure-native:storagesync/v20180701:ServerEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type serverEndpointState struct {
}

type ServerEndpointState struct {
}

func (ServerEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverEndpointState)(nil)).Elem()
}

type serverEndpointArgs struct {
	CloudTiering           *string `pulumi:"cloudTiering"`
	FriendlyName           *string `pulumi:"friendlyName"`
	ResourceGroupName      string  `pulumi:"resourceGroupName"`
	ServerEndpointName     *string `pulumi:"serverEndpointName"`
	ServerLocalPath        *string `pulumi:"serverLocalPath"`
	ServerResourceId       *string `pulumi:"serverResourceId"`
	StorageSyncServiceName string  `pulumi:"storageSyncServiceName"`
	SyncGroupName          string  `pulumi:"syncGroupName"`
	TierFilesOlderThanDays *int    `pulumi:"tierFilesOlderThanDays"`
	VolumeFreeSpacePercent *int    `pulumi:"volumeFreeSpacePercent"`
}


type ServerEndpointArgs struct {
	CloudTiering           pulumi.StringPtrInput
	FriendlyName           pulumi.StringPtrInput
	ResourceGroupName      pulumi.StringInput
	ServerEndpointName     pulumi.StringPtrInput
	ServerLocalPath        pulumi.StringPtrInput
	ServerResourceId       pulumi.StringPtrInput
	StorageSyncServiceName pulumi.StringInput
	SyncGroupName          pulumi.StringInput
	TierFilesOlderThanDays pulumi.IntPtrInput
	VolumeFreeSpacePercent pulumi.IntPtrInput
}

func (ServerEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverEndpointArgs)(nil)).Elem()
}

type ServerEndpointInput interface {
	pulumi.Input

	ToServerEndpointOutput() ServerEndpointOutput
	ToServerEndpointOutputWithContext(ctx context.Context) ServerEndpointOutput
}

func (*ServerEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerEndpoint)(nil)).Elem()
}

func (i *ServerEndpoint) ToServerEndpointOutput() ServerEndpointOutput {
	return i.ToServerEndpointOutputWithContext(context.Background())
}

func (i *ServerEndpoint) ToServerEndpointOutputWithContext(ctx context.Context) ServerEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerEndpointOutput)
}

type ServerEndpointOutput struct{ *pulumi.OutputState }

func (ServerEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerEndpoint)(nil)).Elem()
}

func (o ServerEndpointOutput) ToServerEndpointOutput() ServerEndpointOutput {
	return o
}

func (o ServerEndpointOutput) ToServerEndpointOutputWithContext(ctx context.Context) ServerEndpointOutput {
	return o
}

func (o ServerEndpointOutput) CloudTiering() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpoint) pulumi.StringPtrOutput { return v.CloudTiering }).(pulumi.StringPtrOutput)
}

func (o ServerEndpointOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpoint) pulumi.StringPtrOutput { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o ServerEndpointOutput) LastOperationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpoint) pulumi.StringPtrOutput { return v.LastOperationName }).(pulumi.StringPtrOutput)
}

func (o ServerEndpointOutput) LastWorkflowId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpoint) pulumi.StringPtrOutput { return v.LastWorkflowId }).(pulumi.StringPtrOutput)
}

func (o ServerEndpointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerEndpoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ServerEndpointOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpoint) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ServerEndpointOutput) ServerLocalPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpoint) pulumi.StringPtrOutput { return v.ServerLocalPath }).(pulumi.StringPtrOutput)
}

func (o ServerEndpointOutput) ServerResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpoint) pulumi.StringPtrOutput { return v.ServerResourceId }).(pulumi.StringPtrOutput)
}

func (o ServerEndpointOutput) SyncStatus() pulumi.AnyOutput {
	return o.ApplyT(func(v *ServerEndpoint) pulumi.AnyOutput { return v.SyncStatus }).(pulumi.AnyOutput)
}

func (o ServerEndpointOutput) TierFilesOlderThanDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServerEndpoint) pulumi.IntPtrOutput { return v.TierFilesOlderThanDays }).(pulumi.IntPtrOutput)
}

func (o ServerEndpointOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerEndpoint) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ServerEndpointOutput) VolumeFreeSpacePercent() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServerEndpoint) pulumi.IntPtrOutput { return v.VolumeFreeSpacePercent }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ServerEndpointOutput{})
}
