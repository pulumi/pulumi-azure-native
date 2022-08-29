


package v20150501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupExportConfiguration(ctx *pulumi.Context, args *LookupExportConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupExportConfigurationResult, error) {
	var rv LookupExportConfigurationResult
	err := ctx.Invoke("azure-native:insights/v20150501:getExportConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExportConfigurationArgs struct {
	ExportId          string `pulumi:"exportId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupExportConfigurationResult struct {
	ApplicationName                  string  `pulumi:"applicationName"`
	ContainerName                    string  `pulumi:"containerName"`
	DestinationAccountId             string  `pulumi:"destinationAccountId"`
	DestinationStorageLocationId     string  `pulumi:"destinationStorageLocationId"`
	DestinationStorageSubscriptionId string  `pulumi:"destinationStorageSubscriptionId"`
	DestinationType                  string  `pulumi:"destinationType"`
	ExportId                         string  `pulumi:"exportId"`
	ExportStatus                     string  `pulumi:"exportStatus"`
	InstrumentationKey               string  `pulumi:"instrumentationKey"`
	IsUserEnabled                    string  `pulumi:"isUserEnabled"`
	LastGapTime                      string  `pulumi:"lastGapTime"`
	LastSuccessTime                  string  `pulumi:"lastSuccessTime"`
	LastUserUpdate                   string  `pulumi:"lastUserUpdate"`
	NotificationQueueEnabled         *string `pulumi:"notificationQueueEnabled"`
	PermanentErrorReason             string  `pulumi:"permanentErrorReason"`
	RecordTypes                      *string `pulumi:"recordTypes"`
	ResourceGroup                    string  `pulumi:"resourceGroup"`
	StorageName                      string  `pulumi:"storageName"`
	SubscriptionId                   string  `pulumi:"subscriptionId"`
}

func LookupExportConfigurationOutput(ctx *pulumi.Context, args LookupExportConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupExportConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupExportConfigurationResult, error) {
			args := v.(LookupExportConfigurationArgs)
			r, err := LookupExportConfiguration(ctx, &args, opts...)
			var s LookupExportConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupExportConfigurationResultOutput)
}

type LookupExportConfigurationOutputArgs struct {
	ExportId          pulumi.StringInput `pulumi:"exportId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupExportConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExportConfigurationArgs)(nil)).Elem()
}


type LookupExportConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupExportConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExportConfigurationResult)(nil)).Elem()
}

func (o LookupExportConfigurationResultOutput) ToLookupExportConfigurationResultOutput() LookupExportConfigurationResultOutput {
	return o
}

func (o LookupExportConfigurationResultOutput) ToLookupExportConfigurationResultOutputWithContext(ctx context.Context) LookupExportConfigurationResultOutput {
	return o
}

func (o LookupExportConfigurationResultOutput) ApplicationName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportConfigurationResult) string { return v.ApplicationName }).(pulumi.StringOutput)
}

func (o LookupExportConfigurationResultOutput) ContainerName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportConfigurationResult) string { return v.ContainerName }).(pulumi.StringOutput)
}

func (o LookupExportConfigurationResultOutput) DestinationAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportConfigurationResult) string { return v.DestinationAccountId }).(pulumi.StringOutput)
}

func (o LookupExportConfigurationResultOutput) DestinationStorageLocationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportConfigurationResult) string { return v.DestinationStorageLocationId }).(pulumi.StringOutput)
}

func (o LookupExportConfigurationResultOutput) DestinationStorageSubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportConfigurationResult) string { return v.DestinationStorageSubscriptionId }).(pulumi.StringOutput)
}

func (o LookupExportConfigurationResultOutput) DestinationType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportConfigurationResult) string { return v.DestinationType }).(pulumi.StringOutput)
}

func (o LookupExportConfigurationResultOutput) ExportId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportConfigurationResult) string { return v.ExportId }).(pulumi.StringOutput)
}

func (o LookupExportConfigurationResultOutput) ExportStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportConfigurationResult) string { return v.ExportStatus }).(pulumi.StringOutput)
}

func (o LookupExportConfigurationResultOutput) InstrumentationKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportConfigurationResult) string { return v.InstrumentationKey }).(pulumi.StringOutput)
}

func (o LookupExportConfigurationResultOutput) IsUserEnabled() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportConfigurationResult) string { return v.IsUserEnabled }).(pulumi.StringOutput)
}

func (o LookupExportConfigurationResultOutput) LastGapTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportConfigurationResult) string { return v.LastGapTime }).(pulumi.StringOutput)
}

func (o LookupExportConfigurationResultOutput) LastSuccessTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportConfigurationResult) string { return v.LastSuccessTime }).(pulumi.StringOutput)
}

func (o LookupExportConfigurationResultOutput) LastUserUpdate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportConfigurationResult) string { return v.LastUserUpdate }).(pulumi.StringOutput)
}

func (o LookupExportConfigurationResultOutput) NotificationQueueEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExportConfigurationResult) *string { return v.NotificationQueueEnabled }).(pulumi.StringPtrOutput)
}

func (o LookupExportConfigurationResultOutput) PermanentErrorReason() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportConfigurationResult) string { return v.PermanentErrorReason }).(pulumi.StringOutput)
}

func (o LookupExportConfigurationResultOutput) RecordTypes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExportConfigurationResult) *string { return v.RecordTypes }).(pulumi.StringPtrOutput)
}

func (o LookupExportConfigurationResultOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportConfigurationResult) string { return v.ResourceGroup }).(pulumi.StringOutput)
}

func (o LookupExportConfigurationResultOutput) StorageName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportConfigurationResult) string { return v.StorageName }).(pulumi.StringOutput)
}

func (o LookupExportConfigurationResultOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportConfigurationResult) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupExportConfigurationResultOutput{})
}
