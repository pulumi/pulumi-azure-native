


package v20180915

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNotificationChannel(ctx *pulumi.Context, args *LookupNotificationChannelArgs, opts ...pulumi.InvokeOption) (*LookupNotificationChannelResult, error) {
	var rv LookupNotificationChannelResult
	err := ctx.Invoke("azure-native:devtestlab/v20180915:getNotificationChannel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNotificationChannelArgs struct {
	Expand            *string `pulumi:"expand"`
	LabName           string  `pulumi:"labName"`
	Name              string  `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupNotificationChannelResult struct {
	CreatedDate        string            `pulumi:"createdDate"`
	Description        *string           `pulumi:"description"`
	EmailRecipient     *string           `pulumi:"emailRecipient"`
	Events             []EventResponse   `pulumi:"events"`
	Id                 string            `pulumi:"id"`
	Location           *string           `pulumi:"location"`
	Name               string            `pulumi:"name"`
	NotificationLocale *string           `pulumi:"notificationLocale"`
	ProvisioningState  string            `pulumi:"provisioningState"`
	Tags               map[string]string `pulumi:"tags"`
	Type               string            `pulumi:"type"`
	UniqueIdentifier   string            `pulumi:"uniqueIdentifier"`
	WebHookUrl         *string           `pulumi:"webHookUrl"`
}

func LookupNotificationChannelOutput(ctx *pulumi.Context, args LookupNotificationChannelOutputArgs, opts ...pulumi.InvokeOption) LookupNotificationChannelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNotificationChannelResult, error) {
			args := v.(LookupNotificationChannelArgs)
			r, err := LookupNotificationChannel(ctx, &args, opts...)
			var s LookupNotificationChannelResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNotificationChannelResultOutput)
}

type LookupNotificationChannelOutputArgs struct {
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	LabName           pulumi.StringInput    `pulumi:"labName"`
	Name              pulumi.StringInput    `pulumi:"name"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupNotificationChannelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNotificationChannelArgs)(nil)).Elem()
}


type LookupNotificationChannelResultOutput struct{ *pulumi.OutputState }

func (LookupNotificationChannelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNotificationChannelResult)(nil)).Elem()
}

func (o LookupNotificationChannelResultOutput) ToLookupNotificationChannelResultOutput() LookupNotificationChannelResultOutput {
	return o
}

func (o LookupNotificationChannelResultOutput) ToLookupNotificationChannelResultOutputWithContext(ctx context.Context) LookupNotificationChannelResultOutput {
	return o
}

func (o LookupNotificationChannelResultOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationChannelResult) string { return v.CreatedDate }).(pulumi.StringOutput)
}

func (o LookupNotificationChannelResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNotificationChannelResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupNotificationChannelResultOutput) EmailRecipient() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNotificationChannelResult) *string { return v.EmailRecipient }).(pulumi.StringPtrOutput)
}

func (o LookupNotificationChannelResultOutput) Events() EventResponseArrayOutput {
	return o.ApplyT(func(v LookupNotificationChannelResult) []EventResponse { return v.Events }).(EventResponseArrayOutput)
}

func (o LookupNotificationChannelResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationChannelResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNotificationChannelResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNotificationChannelResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupNotificationChannelResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationChannelResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNotificationChannelResultOutput) NotificationLocale() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNotificationChannelResult) *string { return v.NotificationLocale }).(pulumi.StringPtrOutput)
}

func (o LookupNotificationChannelResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationChannelResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupNotificationChannelResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNotificationChannelResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupNotificationChannelResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationChannelResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupNotificationChannelResultOutput) UniqueIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationChannelResult) string { return v.UniqueIdentifier }).(pulumi.StringOutput)
}

func (o LookupNotificationChannelResultOutput) WebHookUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNotificationChannelResult) *string { return v.WebHookUrl }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNotificationChannelResultOutput{})
}
