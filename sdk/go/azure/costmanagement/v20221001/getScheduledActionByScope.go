


package v20221001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupScheduledActionByScope(ctx *pulumi.Context, args *LookupScheduledActionByScopeArgs, opts ...pulumi.InvokeOption) (*LookupScheduledActionByScopeResult, error) {
	var rv LookupScheduledActionByScopeResult
	err := ctx.Invoke("azure-native:costmanagement/v20221001:getScheduledActionByScope", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScheduledActionByScopeArgs struct {
	Name  string `pulumi:"name"`
	Scope string `pulumi:"scope"`
}


type LookupScheduledActionByScopeResult struct {
	DisplayName       string                         `pulumi:"displayName"`
	ETag              string                         `pulumi:"eTag"`
	FileDestination   *FileDestinationResponse       `pulumi:"fileDestination"`
	Id                string                         `pulumi:"id"`
	Kind              *string                        `pulumi:"kind"`
	Name              string                         `pulumi:"name"`
	Notification      NotificationPropertiesResponse `pulumi:"notification"`
	NotificationEmail *string                        `pulumi:"notificationEmail"`
	Schedule          SchedulePropertiesResponse     `pulumi:"schedule"`
	Scope             *string                        `pulumi:"scope"`
	Status            string                         `pulumi:"status"`
	SystemData        SystemDataResponse             `pulumi:"systemData"`
	Type              string                         `pulumi:"type"`
	ViewId            string                         `pulumi:"viewId"`
}

func LookupScheduledActionByScopeOutput(ctx *pulumi.Context, args LookupScheduledActionByScopeOutputArgs, opts ...pulumi.InvokeOption) LookupScheduledActionByScopeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupScheduledActionByScopeResult, error) {
			args := v.(LookupScheduledActionByScopeArgs)
			r, err := LookupScheduledActionByScope(ctx, &args, opts...)
			var s LookupScheduledActionByScopeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupScheduledActionByScopeResultOutput)
}

type LookupScheduledActionByScopeOutputArgs struct {
	Name  pulumi.StringInput `pulumi:"name"`
	Scope pulumi.StringInput `pulumi:"scope"`
}

func (LookupScheduledActionByScopeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScheduledActionByScopeArgs)(nil)).Elem()
}


type LookupScheduledActionByScopeResultOutput struct{ *pulumi.OutputState }

func (LookupScheduledActionByScopeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScheduledActionByScopeResult)(nil)).Elem()
}

func (o LookupScheduledActionByScopeResultOutput) ToLookupScheduledActionByScopeResultOutput() LookupScheduledActionByScopeResultOutput {
	return o
}

func (o LookupScheduledActionByScopeResultOutput) ToLookupScheduledActionByScopeResultOutputWithContext(ctx context.Context) LookupScheduledActionByScopeResultOutput {
	return o
}

func (o LookupScheduledActionByScopeResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledActionByScopeResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupScheduledActionByScopeResultOutput) ETag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledActionByScopeResult) string { return v.ETag }).(pulumi.StringOutput)
}

func (o LookupScheduledActionByScopeResultOutput) FileDestination() FileDestinationResponsePtrOutput {
	return o.ApplyT(func(v LookupScheduledActionByScopeResult) *FileDestinationResponse { return v.FileDestination }).(FileDestinationResponsePtrOutput)
}

func (o LookupScheduledActionByScopeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledActionByScopeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupScheduledActionByScopeResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduledActionByScopeResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupScheduledActionByScopeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledActionByScopeResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupScheduledActionByScopeResultOutput) Notification() NotificationPropertiesResponseOutput {
	return o.ApplyT(func(v LookupScheduledActionByScopeResult) NotificationPropertiesResponse { return v.Notification }).(NotificationPropertiesResponseOutput)
}

func (o LookupScheduledActionByScopeResultOutput) NotificationEmail() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduledActionByScopeResult) *string { return v.NotificationEmail }).(pulumi.StringPtrOutput)
}

func (o LookupScheduledActionByScopeResultOutput) Schedule() SchedulePropertiesResponseOutput {
	return o.ApplyT(func(v LookupScheduledActionByScopeResult) SchedulePropertiesResponse { return v.Schedule }).(SchedulePropertiesResponseOutput)
}

func (o LookupScheduledActionByScopeResultOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduledActionByScopeResult) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

func (o LookupScheduledActionByScopeResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledActionByScopeResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupScheduledActionByScopeResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupScheduledActionByScopeResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupScheduledActionByScopeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledActionByScopeResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupScheduledActionByScopeResultOutput) ViewId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledActionByScopeResult) string { return v.ViewId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupScheduledActionByScopeResultOutput{})
}
