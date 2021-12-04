


package v20161001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IscsiServer struct {
	pulumi.CustomResourceState

	BackupScheduleGroupId pulumi.StringOutput    `pulumi:"backupScheduleGroupId"`
	ChapId                pulumi.StringPtrOutput `pulumi:"chapId"`
	Description           pulumi.StringPtrOutput `pulumi:"description"`
	Name                  pulumi.StringOutput    `pulumi:"name"`
	ReverseChapId         pulumi.StringPtrOutput `pulumi:"reverseChapId"`
	StorageDomainId       pulumi.StringOutput    `pulumi:"storageDomainId"`
	Type                  pulumi.StringOutput    `pulumi:"type"`
}


func NewIscsiServer(ctx *pulumi.Context,
	name string, args *IscsiServerArgs, opts ...pulumi.ResourceOption) (*IscsiServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BackupScheduleGroupId == nil {
		return nil, errors.New("invalid value for required argument 'BackupScheduleGroupId'")
	}
	if args.DeviceName == nil {
		return nil, errors.New("invalid value for required argument 'DeviceName'")
	}
	if args.ManagerName == nil {
		return nil, errors.New("invalid value for required argument 'ManagerName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StorageDomainId == nil {
		return nil, errors.New("invalid value for required argument 'StorageDomainId'")
	}
	var resource IscsiServer
	err := ctx.RegisterResource("azure-native:storsimple/v20161001:IscsiServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIscsiServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IscsiServerState, opts ...pulumi.ResourceOption) (*IscsiServer, error) {
	var resource IscsiServer
	err := ctx.ReadResource("azure-native:storsimple/v20161001:IscsiServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type iscsiServerState struct {
}

type IscsiServerState struct {
}

func (IscsiServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*iscsiServerState)(nil)).Elem()
}

type iscsiServerArgs struct {
	BackupScheduleGroupId string  `pulumi:"backupScheduleGroupId"`
	ChapId                *string `pulumi:"chapId"`
	Description           *string `pulumi:"description"`
	DeviceName            string  `pulumi:"deviceName"`
	IscsiServerName       *string `pulumi:"iscsiServerName"`
	ManagerName           string  `pulumi:"managerName"`
	ResourceGroupName     string  `pulumi:"resourceGroupName"`
	ReverseChapId         *string `pulumi:"reverseChapId"`
	StorageDomainId       string  `pulumi:"storageDomainId"`
}


type IscsiServerArgs struct {
	BackupScheduleGroupId pulumi.StringInput
	ChapId                pulumi.StringPtrInput
	Description           pulumi.StringPtrInput
	DeviceName            pulumi.StringInput
	IscsiServerName       pulumi.StringPtrInput
	ManagerName           pulumi.StringInput
	ResourceGroupName     pulumi.StringInput
	ReverseChapId         pulumi.StringPtrInput
	StorageDomainId       pulumi.StringInput
}

func (IscsiServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iscsiServerArgs)(nil)).Elem()
}

type IscsiServerInput interface {
	pulumi.Input

	ToIscsiServerOutput() IscsiServerOutput
	ToIscsiServerOutputWithContext(ctx context.Context) IscsiServerOutput
}

func (*IscsiServer) ElementType() reflect.Type {
	return reflect.TypeOf((*IscsiServer)(nil))
}

func (i *IscsiServer) ToIscsiServerOutput() IscsiServerOutput {
	return i.ToIscsiServerOutputWithContext(context.Background())
}

func (i *IscsiServer) ToIscsiServerOutputWithContext(ctx context.Context) IscsiServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IscsiServerOutput)
}

type IscsiServerOutput struct{ *pulumi.OutputState }

func (IscsiServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IscsiServer)(nil))
}

func (o IscsiServerOutput) ToIscsiServerOutput() IscsiServerOutput {
	return o
}

func (o IscsiServerOutput) ToIscsiServerOutputWithContext(ctx context.Context) IscsiServerOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IscsiServerOutput{})
}
