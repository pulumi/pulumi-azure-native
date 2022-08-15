


package v20170605preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type ServerEndpoint struct {
	pulumi.CustomResourceState

	ByteProgress            pulumi.IntPtrOutput    `pulumi:"byteProgress"`
	ByteTotal               pulumi.IntPtrOutput    `pulumi:"byteTotal"`
	CloudTiering            pulumi.StringPtrOutput `pulumi:"cloudTiering"`
	CurrentProgressType     pulumi.StringPtrOutput `pulumi:"currentProgressType"`
	FriendlyName            pulumi.StringPtrOutput `pulumi:"friendlyName"`
	ItemDownloadErrorCount  pulumi.IntPtrOutput    `pulumi:"itemDownloadErrorCount"`
	ItemProgressCount       pulumi.IntPtrOutput    `pulumi:"itemProgressCount"`
	ItemTotalCount          pulumi.IntPtrOutput    `pulumi:"itemTotalCount"`
	ItemUploadErrorCount    pulumi.IntPtrOutput    `pulumi:"itemUploadErrorCount"`
	LastSyncSuccess         pulumi.StringPtrOutput `pulumi:"lastSyncSuccess"`
	LastWorkflowId          pulumi.StringPtrOutput `pulumi:"lastWorkflowId"`
	Name                    pulumi.StringOutput    `pulumi:"name"`
	ProvisioningState       pulumi.StringPtrOutput `pulumi:"provisioningState"`
	ServerLocalPath         pulumi.StringPtrOutput `pulumi:"serverLocalPath"`
	ServerResourceId        pulumi.StringPtrOutput `pulumi:"serverResourceId"`
	SyncErrorContext        pulumi.StringPtrOutput `pulumi:"syncErrorContext"`
	SyncErrorDirection      pulumi.StringPtrOutput `pulumi:"syncErrorDirection"`
	SyncErrorState          pulumi.StringPtrOutput `pulumi:"syncErrorState"`
	SyncErrorStateTimestamp pulumi.StringPtrOutput `pulumi:"syncErrorStateTimestamp"`
	TotalProgress           pulumi.IntPtrOutput    `pulumi:"totalProgress"`
	Type                    pulumi.StringOutput    `pulumi:"type"`
	VolumeFreeSpacePercent  pulumi.IntPtrOutput    `pulumi:"volumeFreeSpacePercent"`
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
			Type: pulumi.String("azure-native:storagesync/v20180402:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20180701:ServerEndpoint"),
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
	err := ctx.RegisterResource("azure-native:storagesync/v20170605preview:ServerEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetServerEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerEndpointState, opts ...pulumi.ResourceOption) (*ServerEndpoint, error) {
	var resource ServerEndpoint
	err := ctx.ReadResource("azure-native:storagesync/v20170605preview:ServerEndpoint", name, id, state, &resource, opts...)
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
	ByteProgress            *int    `pulumi:"byteProgress"`
	ByteTotal               *int    `pulumi:"byteTotal"`
	CloudTiering            *string `pulumi:"cloudTiering"`
	CurrentProgressType     *string `pulumi:"currentProgressType"`
	FriendlyName            *string `pulumi:"friendlyName"`
	ItemDownloadErrorCount  *int    `pulumi:"itemDownloadErrorCount"`
	ItemProgressCount       *int    `pulumi:"itemProgressCount"`
	ItemTotalCount          *int    `pulumi:"itemTotalCount"`
	ItemUploadErrorCount    *int    `pulumi:"itemUploadErrorCount"`
	LastSyncSuccess         *string `pulumi:"lastSyncSuccess"`
	LastWorkflowId          *string `pulumi:"lastWorkflowId"`
	ProvisioningState       *string `pulumi:"provisioningState"`
	ResourceGroupName       string  `pulumi:"resourceGroupName"`
	ServerEndpointName      *string `pulumi:"serverEndpointName"`
	ServerLocalPath         *string `pulumi:"serverLocalPath"`
	ServerResourceId        *string `pulumi:"serverResourceId"`
	StorageSyncServiceName  string  `pulumi:"storageSyncServiceName"`
	SyncErrorContext        *string `pulumi:"syncErrorContext"`
	SyncErrorDirection      *string `pulumi:"syncErrorDirection"`
	SyncErrorState          *string `pulumi:"syncErrorState"`
	SyncErrorStateTimestamp *string `pulumi:"syncErrorStateTimestamp"`
	SyncGroupName           string  `pulumi:"syncGroupName"`
	TotalProgress           *int    `pulumi:"totalProgress"`
	VolumeFreeSpacePercent  *int    `pulumi:"volumeFreeSpacePercent"`
}


type ServerEndpointArgs struct {
	ByteProgress            pulumi.IntPtrInput
	ByteTotal               pulumi.IntPtrInput
	CloudTiering            pulumi.StringPtrInput
	CurrentProgressType     pulumi.StringPtrInput
	FriendlyName            pulumi.StringPtrInput
	ItemDownloadErrorCount  pulumi.IntPtrInput
	ItemProgressCount       pulumi.IntPtrInput
	ItemTotalCount          pulumi.IntPtrInput
	ItemUploadErrorCount    pulumi.IntPtrInput
	LastSyncSuccess         pulumi.StringPtrInput
	LastWorkflowId          pulumi.StringPtrInput
	ProvisioningState       pulumi.StringPtrInput
	ResourceGroupName       pulumi.StringInput
	ServerEndpointName      pulumi.StringPtrInput
	ServerLocalPath         pulumi.StringPtrInput
	ServerResourceId        pulumi.StringPtrInput
	StorageSyncServiceName  pulumi.StringInput
	SyncErrorContext        pulumi.StringPtrInput
	SyncErrorDirection      pulumi.StringPtrInput
	SyncErrorState          pulumi.StringPtrInput
	SyncErrorStateTimestamp pulumi.StringPtrInput
	SyncGroupName           pulumi.StringInput
	TotalProgress           pulumi.IntPtrInput
	VolumeFreeSpacePercent  pulumi.IntPtrInput
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

func (o ServerEndpointOutput) ByteProgress() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServerEndpoint) pulumi.IntPtrOutput { return v.ByteProgress }).(pulumi.IntPtrOutput)
}

func (o ServerEndpointOutput) ByteTotal() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServerEndpoint) pulumi.IntPtrOutput { return v.ByteTotal }).(pulumi.IntPtrOutput)
}

func (o ServerEndpointOutput) CloudTiering() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpoint) pulumi.StringPtrOutput { return v.CloudTiering }).(pulumi.StringPtrOutput)
}

func (o ServerEndpointOutput) CurrentProgressType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpoint) pulumi.StringPtrOutput { return v.CurrentProgressType }).(pulumi.StringPtrOutput)
}

func (o ServerEndpointOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpoint) pulumi.StringPtrOutput { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o ServerEndpointOutput) ItemDownloadErrorCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServerEndpoint) pulumi.IntPtrOutput { return v.ItemDownloadErrorCount }).(pulumi.IntPtrOutput)
}

func (o ServerEndpointOutput) ItemProgressCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServerEndpoint) pulumi.IntPtrOutput { return v.ItemProgressCount }).(pulumi.IntPtrOutput)
}

func (o ServerEndpointOutput) ItemTotalCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServerEndpoint) pulumi.IntPtrOutput { return v.ItemTotalCount }).(pulumi.IntPtrOutput)
}

func (o ServerEndpointOutput) ItemUploadErrorCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServerEndpoint) pulumi.IntPtrOutput { return v.ItemUploadErrorCount }).(pulumi.IntPtrOutput)
}

func (o ServerEndpointOutput) LastSyncSuccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpoint) pulumi.StringPtrOutput { return v.LastSyncSuccess }).(pulumi.StringPtrOutput)
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

func (o ServerEndpointOutput) SyncErrorContext() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpoint) pulumi.StringPtrOutput { return v.SyncErrorContext }).(pulumi.StringPtrOutput)
}

func (o ServerEndpointOutput) SyncErrorDirection() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpoint) pulumi.StringPtrOutput { return v.SyncErrorDirection }).(pulumi.StringPtrOutput)
}

func (o ServerEndpointOutput) SyncErrorState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpoint) pulumi.StringPtrOutput { return v.SyncErrorState }).(pulumi.StringPtrOutput)
}

func (o ServerEndpointOutput) SyncErrorStateTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpoint) pulumi.StringPtrOutput { return v.SyncErrorStateTimestamp }).(pulumi.StringPtrOutput)
}

func (o ServerEndpointOutput) TotalProgress() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServerEndpoint) pulumi.IntPtrOutput { return v.TotalProgress }).(pulumi.IntPtrOutput)
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
