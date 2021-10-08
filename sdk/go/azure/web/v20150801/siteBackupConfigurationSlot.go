


package v20150801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SiteBackupConfigurationSlot struct {
	pulumi.CustomResourceState

	BackupSchedule    BackupScheduleResponsePtrOutput          `pulumi:"backupSchedule"`
	Databases         DatabaseBackupSettingResponseArrayOutput `pulumi:"databases"`
	Enabled           pulumi.BoolPtrOutput                     `pulumi:"enabled"`
	Kind              pulumi.StringPtrOutput                   `pulumi:"kind"`
	Location          pulumi.StringOutput                      `pulumi:"location"`
	Name              pulumi.StringPtrOutput                   `pulumi:"name"`
	StorageAccountUrl pulumi.StringPtrOutput                   `pulumi:"storageAccountUrl"`
	Tags              pulumi.StringMapOutput                   `pulumi:"tags"`
	Type              pulumi.StringOutput                      `pulumi:"type"`
}


func NewSiteBackupConfigurationSlot(ctx *pulumi.Context,
	name string, args *SiteBackupConfigurationSlotArgs, opts ...pulumi.ResourceOption) (*SiteBackupConfigurationSlot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Slot == nil {
		return nil, errors.New("invalid value for required argument 'Slot'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:web/v20150801:SiteBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web:SiteBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web:SiteBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:SiteBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20160801:SiteBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:SiteBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20180201:SiteBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:SiteBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20181101:SiteBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:SiteBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20190801:SiteBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:SiteBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200601:SiteBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:SiteBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200901:SiteBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:SiteBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20201001:SiteBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:SiteBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20201201:SiteBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:SiteBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210101:SiteBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:SiteBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210115:SiteBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:SiteBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210201:SiteBackupConfigurationSlot"),
		},
	})
	opts = append(opts, aliases)
	var resource SiteBackupConfigurationSlot
	err := ctx.RegisterResource("azure-native:web/v20150801:SiteBackupConfigurationSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSiteBackupConfigurationSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteBackupConfigurationSlotState, opts ...pulumi.ResourceOption) (*SiteBackupConfigurationSlot, error) {
	var resource SiteBackupConfigurationSlot
	err := ctx.ReadResource("azure-native:web/v20150801:SiteBackupConfigurationSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type siteBackupConfigurationSlotState struct {
}

type SiteBackupConfigurationSlotState struct {
}

func (SiteBackupConfigurationSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*siteBackupConfigurationSlotState)(nil)).Elem()
}

type siteBackupConfigurationSlotArgs struct {
	BackupSchedule    *BackupSchedule         `pulumi:"backupSchedule"`
	Databases         []DatabaseBackupSetting `pulumi:"databases"`
	Enabled           *bool                   `pulumi:"enabled"`
	Id                *string                 `pulumi:"id"`
	Kind              *string                 `pulumi:"kind"`
	Location          *string                 `pulumi:"location"`
	Name              string                  `pulumi:"name"`
	ResourceGroupName string                  `pulumi:"resourceGroupName"`
	Slot              string                  `pulumi:"slot"`
	StorageAccountUrl *string                 `pulumi:"storageAccountUrl"`
	Tags              map[string]string       `pulumi:"tags"`
	Type              string                  `pulumi:"type"`
}


type SiteBackupConfigurationSlotArgs struct {
	BackupSchedule    BackupSchedulePtrInput
	Databases         DatabaseBackupSettingArrayInput
	Enabled           pulumi.BoolPtrInput
	Id                pulumi.StringPtrInput
	Kind              pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Name              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	Slot              pulumi.StringInput
	StorageAccountUrl pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
	Type              pulumi.StringInput
}

func (SiteBackupConfigurationSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*siteBackupConfigurationSlotArgs)(nil)).Elem()
}

type SiteBackupConfigurationSlotInput interface {
	pulumi.Input

	ToSiteBackupConfigurationSlotOutput() SiteBackupConfigurationSlotOutput
	ToSiteBackupConfigurationSlotOutputWithContext(ctx context.Context) SiteBackupConfigurationSlotOutput
}

func (*SiteBackupConfigurationSlot) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteBackupConfigurationSlot)(nil))
}

func (i *SiteBackupConfigurationSlot) ToSiteBackupConfigurationSlotOutput() SiteBackupConfigurationSlotOutput {
	return i.ToSiteBackupConfigurationSlotOutputWithContext(context.Background())
}

func (i *SiteBackupConfigurationSlot) ToSiteBackupConfigurationSlotOutputWithContext(ctx context.Context) SiteBackupConfigurationSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteBackupConfigurationSlotOutput)
}

type SiteBackupConfigurationSlotOutput struct{ *pulumi.OutputState }

func (SiteBackupConfigurationSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteBackupConfigurationSlot)(nil))
}

func (o SiteBackupConfigurationSlotOutput) ToSiteBackupConfigurationSlotOutput() SiteBackupConfigurationSlotOutput {
	return o
}

func (o SiteBackupConfigurationSlotOutput) ToSiteBackupConfigurationSlotOutputWithContext(ctx context.Context) SiteBackupConfigurationSlotOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SiteBackupConfigurationSlotOutput{})
}
