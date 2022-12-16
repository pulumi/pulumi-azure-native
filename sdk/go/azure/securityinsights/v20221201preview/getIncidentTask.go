


package v20221201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIncidentTask(ctx *pulumi.Context, args *LookupIncidentTaskArgs, opts ...pulumi.InvokeOption) (*LookupIncidentTaskResult, error) {
	var rv LookupIncidentTaskResult
	err := ctx.Invoke("azure-native:securityinsights/v20221201preview:getIncidentTask", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIncidentTaskArgs struct {
	IncidentId        string `pulumi:"incidentId"`
	IncidentTaskId    string `pulumi:"incidentTaskId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}

type LookupIncidentTaskResult struct {
	CreatedBy           *ClientInfoResponse `pulumi:"createdBy"`
	CreatedTimeUtc      string              `pulumi:"createdTimeUtc"`
	Description         *string             `pulumi:"description"`
	Etag                *string             `pulumi:"etag"`
	Id                  string              `pulumi:"id"`
	LastModifiedBy      *ClientInfoResponse `pulumi:"lastModifiedBy"`
	LastModifiedTimeUtc string              `pulumi:"lastModifiedTimeUtc"`
	Name                string              `pulumi:"name"`
	Status              string              `pulumi:"status"`
	SystemData          SystemDataResponse  `pulumi:"systemData"`
	Title               string              `pulumi:"title"`
	Type                string              `pulumi:"type"`
}

func LookupIncidentTaskOutput(ctx *pulumi.Context, args LookupIncidentTaskOutputArgs, opts ...pulumi.InvokeOption) LookupIncidentTaskResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIncidentTaskResult, error) {
			args := v.(LookupIncidentTaskArgs)
			r, err := LookupIncidentTask(ctx, &args, opts...)
			var s LookupIncidentTaskResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIncidentTaskResultOutput)
}

type LookupIncidentTaskOutputArgs struct {
	IncidentId        pulumi.StringInput `pulumi:"incidentId"`
	IncidentTaskId    pulumi.StringInput `pulumi:"incidentTaskId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupIncidentTaskOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIncidentTaskArgs)(nil)).Elem()
}

type LookupIncidentTaskResultOutput struct{ *pulumi.OutputState }

func (LookupIncidentTaskResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIncidentTaskResult)(nil)).Elem()
}

func (o LookupIncidentTaskResultOutput) ToLookupIncidentTaskResultOutput() LookupIncidentTaskResultOutput {
	return o
}

func (o LookupIncidentTaskResultOutput) ToLookupIncidentTaskResultOutputWithContext(ctx context.Context) LookupIncidentTaskResultOutput {
	return o
}

func (o LookupIncidentTaskResultOutput) CreatedBy() ClientInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupIncidentTaskResult) *ClientInfoResponse { return v.CreatedBy }).(ClientInfoResponsePtrOutput)
}

func (o LookupIncidentTaskResultOutput) CreatedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentTaskResult) string { return v.CreatedTimeUtc }).(pulumi.StringOutput)
}

func (o LookupIncidentTaskResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIncidentTaskResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupIncidentTaskResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIncidentTaskResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupIncidentTaskResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentTaskResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIncidentTaskResultOutput) LastModifiedBy() ClientInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupIncidentTaskResult) *ClientInfoResponse { return v.LastModifiedBy }).(ClientInfoResponsePtrOutput)
}

func (o LookupIncidentTaskResultOutput) LastModifiedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentTaskResult) string { return v.LastModifiedTimeUtc }).(pulumi.StringOutput)
}

func (o LookupIncidentTaskResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentTaskResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIncidentTaskResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentTaskResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupIncidentTaskResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupIncidentTaskResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupIncidentTaskResultOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentTaskResult) string { return v.Title }).(pulumi.StringOutput)
}

func (o LookupIncidentTaskResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentTaskResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIncidentTaskResultOutput{})
}
