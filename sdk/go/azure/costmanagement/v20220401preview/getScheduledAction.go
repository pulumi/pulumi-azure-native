


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupScheduledAction(ctx *pulumi.Context, args *LookupScheduledActionArgs, opts ...pulumi.InvokeOption) (*LookupScheduledActionResult, error) {
	var rv LookupScheduledActionResult
	err := ctx.Invoke("azure-native:costmanagement/v20220401preview:getScheduledAction", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScheduledActionArgs struct {
	Name string `pulumi:"name"`
}


type LookupScheduledActionResult struct {
	DisplayName     string                         `pulumi:"displayName"`
	ETag            string                         `pulumi:"eTag"`
	FileDestination *FileDestinationResponse       `pulumi:"fileDestination"`
	Id              string                         `pulumi:"id"`
	Kind            *string                        `pulumi:"kind"`
	Name            string                         `pulumi:"name"`
	Notification    NotificationPropertiesResponse `pulumi:"notification"`
	Schedule        SchedulePropertiesResponse     `pulumi:"schedule"`
	Scope           *string                        `pulumi:"scope"`
	Status          string                         `pulumi:"status"`
	SystemData      SystemDataResponse             `pulumi:"systemData"`
	Type            string                         `pulumi:"type"`
	ViewId          string                         `pulumi:"viewId"`
}

func LookupScheduledActionOutput(ctx *pulumi.Context, args LookupScheduledActionOutputArgs, opts ...pulumi.InvokeOption) LookupScheduledActionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupScheduledActionResult, error) {
			args := v.(LookupScheduledActionArgs)
			r, err := LookupScheduledAction(ctx, &args, opts...)
			var s LookupScheduledActionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupScheduledActionResultOutput)
}

type LookupScheduledActionOutputArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupScheduledActionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScheduledActionArgs)(nil)).Elem()
}


type LookupScheduledActionResultOutput struct{ *pulumi.OutputState }

func (LookupScheduledActionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScheduledActionResult)(nil)).Elem()
}

func (o LookupScheduledActionResultOutput) ToLookupScheduledActionResultOutput() LookupScheduledActionResultOutput {
	return o
}

func (o LookupScheduledActionResultOutput) ToLookupScheduledActionResultOutputWithContext(ctx context.Context) LookupScheduledActionResultOutput {
	return o
}

func (o LookupScheduledActionResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledActionResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupScheduledActionResultOutput) ETag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledActionResult) string { return v.ETag }).(pulumi.StringOutput)
}

func (o LookupScheduledActionResultOutput) FileDestination() FileDestinationResponsePtrOutput {
	return o.ApplyT(func(v LookupScheduledActionResult) *FileDestinationResponse { return v.FileDestination }).(FileDestinationResponsePtrOutput)
}

func (o LookupScheduledActionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledActionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupScheduledActionResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduledActionResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupScheduledActionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledActionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupScheduledActionResultOutput) Notification() NotificationPropertiesResponseOutput {
	return o.ApplyT(func(v LookupScheduledActionResult) NotificationPropertiesResponse { return v.Notification }).(NotificationPropertiesResponseOutput)
}

func (o LookupScheduledActionResultOutput) Schedule() SchedulePropertiesResponseOutput {
	return o.ApplyT(func(v LookupScheduledActionResult) SchedulePropertiesResponse { return v.Schedule }).(SchedulePropertiesResponseOutput)
}

func (o LookupScheduledActionResultOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduledActionResult) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

func (o LookupScheduledActionResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledActionResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupScheduledActionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupScheduledActionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupScheduledActionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledActionResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupScheduledActionResultOutput) ViewId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledActionResult) string { return v.ViewId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupScheduledActionResultOutput{})
}
