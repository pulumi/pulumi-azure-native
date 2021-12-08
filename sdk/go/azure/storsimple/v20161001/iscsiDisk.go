


package v20161001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IscsiDisk struct {
	pulumi.CustomResourceState

	AccessControlRecords       pulumi.StringArrayOutput `pulumi:"accessControlRecords"`
	DataPolicy                 pulumi.StringOutput      `pulumi:"dataPolicy"`
	Description                pulumi.StringPtrOutput   `pulumi:"description"`
	DiskStatus                 pulumi.StringOutput      `pulumi:"diskStatus"`
	LocalUsedCapacityInBytes   pulumi.Float64Output     `pulumi:"localUsedCapacityInBytes"`
	MonitoringStatus           pulumi.StringOutput      `pulumi:"monitoringStatus"`
	Name                       pulumi.StringOutput      `pulumi:"name"`
	ProvisionedCapacityInBytes pulumi.Float64Output     `pulumi:"provisionedCapacityInBytes"`
	Type                       pulumi.StringOutput      `pulumi:"type"`
	UsedCapacityInBytes        pulumi.Float64Output     `pulumi:"usedCapacityInBytes"`
}


func NewIscsiDisk(ctx *pulumi.Context,
	name string, args *IscsiDiskArgs, opts ...pulumi.ResourceOption) (*IscsiDisk, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessControlRecords == nil {
		return nil, errors.New("invalid value for required argument 'AccessControlRecords'")
	}
	if args.DataPolicy == nil {
		return nil, errors.New("invalid value for required argument 'DataPolicy'")
	}
	if args.DeviceName == nil {
		return nil, errors.New("invalid value for required argument 'DeviceName'")
	}
	if args.DiskStatus == nil {
		return nil, errors.New("invalid value for required argument 'DiskStatus'")
	}
	if args.IscsiServerName == nil {
		return nil, errors.New("invalid value for required argument 'IscsiServerName'")
	}
	if args.ManagerName == nil {
		return nil, errors.New("invalid value for required argument 'ManagerName'")
	}
	if args.MonitoringStatus == nil {
		return nil, errors.New("invalid value for required argument 'MonitoringStatus'")
	}
	if args.ProvisionedCapacityInBytes == nil {
		return nil, errors.New("invalid value for required argument 'ProvisionedCapacityInBytes'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource IscsiDisk
	err := ctx.RegisterResource("azure-native:storsimple/v20161001:IscsiDisk", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIscsiDisk(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IscsiDiskState, opts ...pulumi.ResourceOption) (*IscsiDisk, error) {
	var resource IscsiDisk
	err := ctx.ReadResource("azure-native:storsimple/v20161001:IscsiDisk", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type iscsiDiskState struct {
}

type IscsiDiskState struct {
}

func (IscsiDiskState) ElementType() reflect.Type {
	return reflect.TypeOf((*iscsiDiskState)(nil)).Elem()
}

type iscsiDiskArgs struct {
	AccessControlRecords       []string         `pulumi:"accessControlRecords"`
	DataPolicy                 DataPolicy       `pulumi:"dataPolicy"`
	Description                *string          `pulumi:"description"`
	DeviceName                 string           `pulumi:"deviceName"`
	DiskName                   *string          `pulumi:"diskName"`
	DiskStatus                 DiskStatus       `pulumi:"diskStatus"`
	IscsiServerName            string           `pulumi:"iscsiServerName"`
	ManagerName                string           `pulumi:"managerName"`
	MonitoringStatus           MonitoringStatus `pulumi:"monitoringStatus"`
	ProvisionedCapacityInBytes float64          `pulumi:"provisionedCapacityInBytes"`
	ResourceGroupName          string           `pulumi:"resourceGroupName"`
}


type IscsiDiskArgs struct {
	AccessControlRecords       pulumi.StringArrayInput
	DataPolicy                 DataPolicyInput
	Description                pulumi.StringPtrInput
	DeviceName                 pulumi.StringInput
	DiskName                   pulumi.StringPtrInput
	DiskStatus                 DiskStatusInput
	IscsiServerName            pulumi.StringInput
	ManagerName                pulumi.StringInput
	MonitoringStatus           MonitoringStatusInput
	ProvisionedCapacityInBytes pulumi.Float64Input
	ResourceGroupName          pulumi.StringInput
}

func (IscsiDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iscsiDiskArgs)(nil)).Elem()
}

type IscsiDiskInput interface {
	pulumi.Input

	ToIscsiDiskOutput() IscsiDiskOutput
	ToIscsiDiskOutputWithContext(ctx context.Context) IscsiDiskOutput
}

func (*IscsiDisk) ElementType() reflect.Type {
	return reflect.TypeOf((*IscsiDisk)(nil))
}

func (i *IscsiDisk) ToIscsiDiskOutput() IscsiDiskOutput {
	return i.ToIscsiDiskOutputWithContext(context.Background())
}

func (i *IscsiDisk) ToIscsiDiskOutputWithContext(ctx context.Context) IscsiDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IscsiDiskOutput)
}

type IscsiDiskOutput struct{ *pulumi.OutputState }

func (IscsiDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IscsiDisk)(nil))
}

func (o IscsiDiskOutput) ToIscsiDiskOutput() IscsiDiskOutput {
	return o
}

func (o IscsiDiskOutput) ToIscsiDiskOutputWithContext(ctx context.Context) IscsiDiskOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IscsiDiskOutput{})
}
