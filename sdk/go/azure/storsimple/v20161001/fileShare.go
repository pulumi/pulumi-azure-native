


package v20161001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type FileShare struct {
	pulumi.CustomResourceState

	AdminUser                  pulumi.StringOutput    `pulumi:"adminUser"`
	DataPolicy                 pulumi.StringOutput    `pulumi:"dataPolicy"`
	Description                pulumi.StringPtrOutput `pulumi:"description"`
	LocalUsedCapacityInBytes   pulumi.Float64Output   `pulumi:"localUsedCapacityInBytes"`
	MonitoringStatus           pulumi.StringOutput    `pulumi:"monitoringStatus"`
	Name                       pulumi.StringOutput    `pulumi:"name"`
	ProvisionedCapacityInBytes pulumi.Float64Output   `pulumi:"provisionedCapacityInBytes"`
	ShareStatus                pulumi.StringOutput    `pulumi:"shareStatus"`
	Type                       pulumi.StringOutput    `pulumi:"type"`
	UsedCapacityInBytes        pulumi.Float64Output   `pulumi:"usedCapacityInBytes"`
}


func NewFileShare(ctx *pulumi.Context,
	name string, args *FileShareArgs, opts ...pulumi.ResourceOption) (*FileShare, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AdminUser == nil {
		return nil, errors.New("invalid value for required argument 'AdminUser'")
	}
	if args.DataPolicy == nil {
		return nil, errors.New("invalid value for required argument 'DataPolicy'")
	}
	if args.DeviceName == nil {
		return nil, errors.New("invalid value for required argument 'DeviceName'")
	}
	if args.FileServerName == nil {
		return nil, errors.New("invalid value for required argument 'FileServerName'")
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
	if args.ShareStatus == nil {
		return nil, errors.New("invalid value for required argument 'ShareStatus'")
	}
	var resource FileShare
	err := ctx.RegisterResource("azure-native:storsimple/v20161001:FileShare", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetFileShare(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FileShareState, opts ...pulumi.ResourceOption) (*FileShare, error) {
	var resource FileShare
	err := ctx.ReadResource("azure-native:storsimple/v20161001:FileShare", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type fileShareState struct {
}

type FileShareState struct {
}

func (FileShareState) ElementType() reflect.Type {
	return reflect.TypeOf((*fileShareState)(nil)).Elem()
}

type fileShareArgs struct {
	AdminUser                  string           `pulumi:"adminUser"`
	DataPolicy                 DataPolicy       `pulumi:"dataPolicy"`
	Description                *string          `pulumi:"description"`
	DeviceName                 string           `pulumi:"deviceName"`
	FileServerName             string           `pulumi:"fileServerName"`
	ManagerName                string           `pulumi:"managerName"`
	MonitoringStatus           MonitoringStatus `pulumi:"monitoringStatus"`
	ProvisionedCapacityInBytes float64          `pulumi:"provisionedCapacityInBytes"`
	ResourceGroupName          string           `pulumi:"resourceGroupName"`
	ShareName                  *string          `pulumi:"shareName"`
	ShareStatus                ShareStatus      `pulumi:"shareStatus"`
}


type FileShareArgs struct {
	AdminUser                  pulumi.StringInput
	DataPolicy                 DataPolicyInput
	Description                pulumi.StringPtrInput
	DeviceName                 pulumi.StringInput
	FileServerName             pulumi.StringInput
	ManagerName                pulumi.StringInput
	MonitoringStatus           MonitoringStatusInput
	ProvisionedCapacityInBytes pulumi.Float64Input
	ResourceGroupName          pulumi.StringInput
	ShareName                  pulumi.StringPtrInput
	ShareStatus                ShareStatusInput
}

func (FileShareArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fileShareArgs)(nil)).Elem()
}

type FileShareInput interface {
	pulumi.Input

	ToFileShareOutput() FileShareOutput
	ToFileShareOutputWithContext(ctx context.Context) FileShareOutput
}

func (*FileShare) ElementType() reflect.Type {
	return reflect.TypeOf((**FileShare)(nil)).Elem()
}

func (i *FileShare) ToFileShareOutput() FileShareOutput {
	return i.ToFileShareOutputWithContext(context.Background())
}

func (i *FileShare) ToFileShareOutputWithContext(ctx context.Context) FileShareOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileShareOutput)
}

type FileShareOutput struct{ *pulumi.OutputState }

func (FileShareOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FileShare)(nil)).Elem()
}

func (o FileShareOutput) ToFileShareOutput() FileShareOutput {
	return o
}

func (o FileShareOutput) ToFileShareOutputWithContext(ctx context.Context) FileShareOutput {
	return o
}

func (o FileShareOutput) AdminUser() pulumi.StringOutput {
	return o.ApplyT(func(v *FileShare) pulumi.StringOutput { return v.AdminUser }).(pulumi.StringOutput)
}

func (o FileShareOutput) DataPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *FileShare) pulumi.StringOutput { return v.DataPolicy }).(pulumi.StringOutput)
}

func (o FileShareOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FileShare) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o FileShareOutput) LocalUsedCapacityInBytes() pulumi.Float64Output {
	return o.ApplyT(func(v *FileShare) pulumi.Float64Output { return v.LocalUsedCapacityInBytes }).(pulumi.Float64Output)
}

func (o FileShareOutput) MonitoringStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *FileShare) pulumi.StringOutput { return v.MonitoringStatus }).(pulumi.StringOutput)
}

func (o FileShareOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FileShare) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FileShareOutput) ProvisionedCapacityInBytes() pulumi.Float64Output {
	return o.ApplyT(func(v *FileShare) pulumi.Float64Output { return v.ProvisionedCapacityInBytes }).(pulumi.Float64Output)
}

func (o FileShareOutput) ShareStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *FileShare) pulumi.StringOutput { return v.ShareStatus }).(pulumi.StringOutput)
}

func (o FileShareOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FileShare) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o FileShareOutput) UsedCapacityInBytes() pulumi.Float64Output {
	return o.ApplyT(func(v *FileShare) pulumi.Float64Output { return v.UsedCapacityInBytes }).(pulumi.Float64Output)
}

func init() {
	pulumi.RegisterOutputType(FileShareOutput{})
}
