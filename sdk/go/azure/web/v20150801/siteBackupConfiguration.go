


package v20150801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type SiteBackupConfiguration struct {
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


func NewSiteBackupConfiguration(ctx *pulumi.Context,
	name string, args *SiteBackupConfigurationArgs, opts ...pulumi.ResourceOption) (*SiteBackupConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:SiteBackupConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:SiteBackupConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:SiteBackupConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:SiteBackupConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:SiteBackupConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:SiteBackupConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:SiteBackupConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:SiteBackupConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:SiteBackupConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:SiteBackupConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:SiteBackupConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:SiteBackupConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:SiteBackupConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:SiteBackupConfiguration"),
		},
	})
	opts = append(opts, aliases)
	var resource SiteBackupConfiguration
	err := ctx.RegisterResource("azure-native:web/v20150801:SiteBackupConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSiteBackupConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteBackupConfigurationState, opts ...pulumi.ResourceOption) (*SiteBackupConfiguration, error) {
	var resource SiteBackupConfiguration
	err := ctx.ReadResource("azure-native:web/v20150801:SiteBackupConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type siteBackupConfigurationState struct {
}

type SiteBackupConfigurationState struct {
}

func (SiteBackupConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*siteBackupConfigurationState)(nil)).Elem()
}

type siteBackupConfigurationArgs struct {
	BackupSchedule    *BackupSchedule         `pulumi:"backupSchedule"`
	Databases         []DatabaseBackupSetting `pulumi:"databases"`
	Enabled           *bool                   `pulumi:"enabled"`
	Id                *string                 `pulumi:"id"`
	Kind              *string                 `pulumi:"kind"`
	Location          *string                 `pulumi:"location"`
	Name              string                  `pulumi:"name"`
	ResourceGroupName string                  `pulumi:"resourceGroupName"`
	StorageAccountUrl *string                 `pulumi:"storageAccountUrl"`
	Tags              map[string]string       `pulumi:"tags"`
	Type              string                  `pulumi:"type"`
}


type SiteBackupConfigurationArgs struct {
	BackupSchedule    BackupSchedulePtrInput
	Databases         DatabaseBackupSettingArrayInput
	Enabled           pulumi.BoolPtrInput
	Id                pulumi.StringPtrInput
	Kind              pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Name              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	StorageAccountUrl pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
	Type              pulumi.StringInput
}

func (SiteBackupConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*siteBackupConfigurationArgs)(nil)).Elem()
}

type SiteBackupConfigurationInput interface {
	pulumi.Input

	ToSiteBackupConfigurationOutput() SiteBackupConfigurationOutput
	ToSiteBackupConfigurationOutputWithContext(ctx context.Context) SiteBackupConfigurationOutput
}

func (*SiteBackupConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteBackupConfiguration)(nil)).Elem()
}

func (i *SiteBackupConfiguration) ToSiteBackupConfigurationOutput() SiteBackupConfigurationOutput {
	return i.ToSiteBackupConfigurationOutputWithContext(context.Background())
}

func (i *SiteBackupConfiguration) ToSiteBackupConfigurationOutputWithContext(ctx context.Context) SiteBackupConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteBackupConfigurationOutput)
}

type SiteBackupConfigurationOutput struct{ *pulumi.OutputState }

func (SiteBackupConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteBackupConfiguration)(nil)).Elem()
}

func (o SiteBackupConfigurationOutput) ToSiteBackupConfigurationOutput() SiteBackupConfigurationOutput {
	return o
}

func (o SiteBackupConfigurationOutput) ToSiteBackupConfigurationOutputWithContext(ctx context.Context) SiteBackupConfigurationOutput {
	return o
}

func (o SiteBackupConfigurationOutput) BackupSchedule() BackupScheduleResponsePtrOutput {
	return o.ApplyT(func(v *SiteBackupConfiguration) BackupScheduleResponsePtrOutput { return v.BackupSchedule }).(BackupScheduleResponsePtrOutput)
}

func (o SiteBackupConfigurationOutput) Databases() DatabaseBackupSettingResponseArrayOutput {
	return o.ApplyT(func(v *SiteBackupConfiguration) DatabaseBackupSettingResponseArrayOutput { return v.Databases }).(DatabaseBackupSettingResponseArrayOutput)
}

func (o SiteBackupConfigurationOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteBackupConfiguration) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o SiteBackupConfigurationOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteBackupConfiguration) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o SiteBackupConfigurationOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SiteBackupConfiguration) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o SiteBackupConfigurationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteBackupConfiguration) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SiteBackupConfigurationOutput) StorageAccountUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteBackupConfiguration) pulumi.StringPtrOutput { return v.StorageAccountUrl }).(pulumi.StringPtrOutput)
}

func (o SiteBackupConfigurationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SiteBackupConfiguration) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SiteBackupConfigurationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SiteBackupConfiguration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SiteBackupConfigurationOutput{})
}
