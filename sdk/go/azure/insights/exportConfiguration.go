


package insights

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ExportConfiguration struct {
	pulumi.CustomResourceState

	ApplicationName                  pulumi.StringOutput    `pulumi:"applicationName"`
	ContainerName                    pulumi.StringOutput    `pulumi:"containerName"`
	DestinationAccountId             pulumi.StringOutput    `pulumi:"destinationAccountId"`
	DestinationStorageLocationId     pulumi.StringOutput    `pulumi:"destinationStorageLocationId"`
	DestinationStorageSubscriptionId pulumi.StringOutput    `pulumi:"destinationStorageSubscriptionId"`
	DestinationType                  pulumi.StringOutput    `pulumi:"destinationType"`
	ExportId                         pulumi.StringOutput    `pulumi:"exportId"`
	ExportStatus                     pulumi.StringOutput    `pulumi:"exportStatus"`
	InstrumentationKey               pulumi.StringOutput    `pulumi:"instrumentationKey"`
	IsUserEnabled                    pulumi.StringOutput    `pulumi:"isUserEnabled"`
	LastGapTime                      pulumi.StringOutput    `pulumi:"lastGapTime"`
	LastSuccessTime                  pulumi.StringOutput    `pulumi:"lastSuccessTime"`
	LastUserUpdate                   pulumi.StringOutput    `pulumi:"lastUserUpdate"`
	NotificationQueueEnabled         pulumi.StringPtrOutput `pulumi:"notificationQueueEnabled"`
	PermanentErrorReason             pulumi.StringOutput    `pulumi:"permanentErrorReason"`
	RecordTypes                      pulumi.StringPtrOutput `pulumi:"recordTypes"`
	ResourceGroup                    pulumi.StringOutput    `pulumi:"resourceGroup"`
	StorageName                      pulumi.StringOutput    `pulumi:"storageName"`
	SubscriptionId                   pulumi.StringOutput    `pulumi:"subscriptionId"`
}


func NewExportConfiguration(ctx *pulumi.Context,
	name string, args *ExportConfigurationArgs, opts ...pulumi.ResourceOption) (*ExportConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:insights/v20150501:ExportConfiguration"),
		},
	})
	opts = append(opts, aliases)
	var resource ExportConfiguration
	err := ctx.RegisterResource("azure-native:insights:ExportConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetExportConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExportConfigurationState, opts ...pulumi.ResourceOption) (*ExportConfiguration, error) {
	var resource ExportConfiguration
	err := ctx.ReadResource("azure-native:insights:ExportConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type exportConfigurationState struct {
}

type ExportConfigurationState struct {
}

func (ExportConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*exportConfigurationState)(nil)).Elem()
}

type exportConfigurationArgs struct {
	DestinationAccountId             *string `pulumi:"destinationAccountId"`
	DestinationAddress               *string `pulumi:"destinationAddress"`
	DestinationStorageLocationId     *string `pulumi:"destinationStorageLocationId"`
	DestinationStorageSubscriptionId *string `pulumi:"destinationStorageSubscriptionId"`
	DestinationType                  *string `pulumi:"destinationType"`
	ExportId                         *string `pulumi:"exportId"`
	IsEnabled                        *string `pulumi:"isEnabled"`
	NotificationQueueEnabled         *string `pulumi:"notificationQueueEnabled"`
	NotificationQueueUri             *string `pulumi:"notificationQueueUri"`
	RecordTypes                      *string `pulumi:"recordTypes"`
	ResourceGroupName                string  `pulumi:"resourceGroupName"`
	ResourceName                     string  `pulumi:"resourceName"`
}


type ExportConfigurationArgs struct {
	DestinationAccountId             pulumi.StringPtrInput
	DestinationAddress               pulumi.StringPtrInput
	DestinationStorageLocationId     pulumi.StringPtrInput
	DestinationStorageSubscriptionId pulumi.StringPtrInput
	DestinationType                  pulumi.StringPtrInput
	ExportId                         pulumi.StringPtrInput
	IsEnabled                        pulumi.StringPtrInput
	NotificationQueueEnabled         pulumi.StringPtrInput
	NotificationQueueUri             pulumi.StringPtrInput
	RecordTypes                      pulumi.StringPtrInput
	ResourceGroupName                pulumi.StringInput
	ResourceName                     pulumi.StringInput
}

func (ExportConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*exportConfigurationArgs)(nil)).Elem()
}

type ExportConfigurationInput interface {
	pulumi.Input

	ToExportConfigurationOutput() ExportConfigurationOutput
	ToExportConfigurationOutputWithContext(ctx context.Context) ExportConfigurationOutput
}

func (*ExportConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportConfiguration)(nil)).Elem()
}

func (i *ExportConfiguration) ToExportConfigurationOutput() ExportConfigurationOutput {
	return i.ToExportConfigurationOutputWithContext(context.Background())
}

func (i *ExportConfiguration) ToExportConfigurationOutputWithContext(ctx context.Context) ExportConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportConfigurationOutput)
}

type ExportConfigurationOutput struct{ *pulumi.OutputState }

func (ExportConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportConfiguration)(nil)).Elem()
}

func (o ExportConfigurationOutput) ToExportConfigurationOutput() ExportConfigurationOutput {
	return o
}

func (o ExportConfigurationOutput) ToExportConfigurationOutputWithContext(ctx context.Context) ExportConfigurationOutput {
	return o
}

func (o ExportConfigurationOutput) ApplicationName() pulumi.StringOutput {
	return o.ApplyT(func(v *ExportConfiguration) pulumi.StringOutput { return v.ApplicationName }).(pulumi.StringOutput)
}

func (o ExportConfigurationOutput) ContainerName() pulumi.StringOutput {
	return o.ApplyT(func(v *ExportConfiguration) pulumi.StringOutput { return v.ContainerName }).(pulumi.StringOutput)
}

func (o ExportConfigurationOutput) DestinationAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *ExportConfiguration) pulumi.StringOutput { return v.DestinationAccountId }).(pulumi.StringOutput)
}

func (o ExportConfigurationOutput) DestinationStorageLocationId() pulumi.StringOutput {
	return o.ApplyT(func(v *ExportConfiguration) pulumi.StringOutput { return v.DestinationStorageLocationId }).(pulumi.StringOutput)
}

func (o ExportConfigurationOutput) DestinationStorageSubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v *ExportConfiguration) pulumi.StringOutput { return v.DestinationStorageSubscriptionId }).(pulumi.StringOutput)
}

func (o ExportConfigurationOutput) DestinationType() pulumi.StringOutput {
	return o.ApplyT(func(v *ExportConfiguration) pulumi.StringOutput { return v.DestinationType }).(pulumi.StringOutput)
}

func (o ExportConfigurationOutput) ExportId() pulumi.StringOutput {
	return o.ApplyT(func(v *ExportConfiguration) pulumi.StringOutput { return v.ExportId }).(pulumi.StringOutput)
}

func (o ExportConfigurationOutput) ExportStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *ExportConfiguration) pulumi.StringOutput { return v.ExportStatus }).(pulumi.StringOutput)
}

func (o ExportConfigurationOutput) InstrumentationKey() pulumi.StringOutput {
	return o.ApplyT(func(v *ExportConfiguration) pulumi.StringOutput { return v.InstrumentationKey }).(pulumi.StringOutput)
}

func (o ExportConfigurationOutput) IsUserEnabled() pulumi.StringOutput {
	return o.ApplyT(func(v *ExportConfiguration) pulumi.StringOutput { return v.IsUserEnabled }).(pulumi.StringOutput)
}

func (o ExportConfigurationOutput) LastGapTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ExportConfiguration) pulumi.StringOutput { return v.LastGapTime }).(pulumi.StringOutput)
}

func (o ExportConfigurationOutput) LastSuccessTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ExportConfiguration) pulumi.StringOutput { return v.LastSuccessTime }).(pulumi.StringOutput)
}

func (o ExportConfigurationOutput) LastUserUpdate() pulumi.StringOutput {
	return o.ApplyT(func(v *ExportConfiguration) pulumi.StringOutput { return v.LastUserUpdate }).(pulumi.StringOutput)
}

func (o ExportConfigurationOutput) NotificationQueueEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportConfiguration) pulumi.StringPtrOutput { return v.NotificationQueueEnabled }).(pulumi.StringPtrOutput)
}

func (o ExportConfigurationOutput) PermanentErrorReason() pulumi.StringOutput {
	return o.ApplyT(func(v *ExportConfiguration) pulumi.StringOutput { return v.PermanentErrorReason }).(pulumi.StringOutput)
}

func (o ExportConfigurationOutput) RecordTypes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportConfiguration) pulumi.StringPtrOutput { return v.RecordTypes }).(pulumi.StringPtrOutput)
}

func (o ExportConfigurationOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *ExportConfiguration) pulumi.StringOutput { return v.ResourceGroup }).(pulumi.StringOutput)
}

func (o ExportConfigurationOutput) StorageName() pulumi.StringOutput {
	return o.ApplyT(func(v *ExportConfiguration) pulumi.StringOutput { return v.StorageName }).(pulumi.StringOutput)
}

func (o ExportConfigurationOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v *ExportConfiguration) pulumi.StringOutput { return v.SubscriptionId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ExportConfigurationOutput{})
}
