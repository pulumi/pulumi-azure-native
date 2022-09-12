


package v20220801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBookmark(ctx *pulumi.Context, args *LookupBookmarkArgs, opts ...pulumi.InvokeOption) (*LookupBookmarkResult, error) {
	var rv LookupBookmarkResult
	err := ctx.Invoke("azure-native:securityinsights/v20220801:getBookmark", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBookmarkArgs struct {
	BookmarkId        string `pulumi:"bookmarkId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupBookmarkResult struct {
	Created        *string               `pulumi:"created"`
	CreatedBy      *UserInfoResponse     `pulumi:"createdBy"`
	DisplayName    string                `pulumi:"displayName"`
	Etag           *string               `pulumi:"etag"`
	EventTime      *string               `pulumi:"eventTime"`
	Id             string                `pulumi:"id"`
	IncidentInfo   *IncidentInfoResponse `pulumi:"incidentInfo"`
	Labels         []string              `pulumi:"labels"`
	Name           string                `pulumi:"name"`
	Notes          *string               `pulumi:"notes"`
	Query          string                `pulumi:"query"`
	QueryEndTime   *string               `pulumi:"queryEndTime"`
	QueryResult    *string               `pulumi:"queryResult"`
	QueryStartTime *string               `pulumi:"queryStartTime"`
	SystemData     SystemDataResponse    `pulumi:"systemData"`
	Type           string                `pulumi:"type"`
	Updated        *string               `pulumi:"updated"`
	UpdatedBy      *UserInfoResponse     `pulumi:"updatedBy"`
}

func LookupBookmarkOutput(ctx *pulumi.Context, args LookupBookmarkOutputArgs, opts ...pulumi.InvokeOption) LookupBookmarkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBookmarkResult, error) {
			args := v.(LookupBookmarkArgs)
			r, err := LookupBookmark(ctx, &args, opts...)
			var s LookupBookmarkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBookmarkResultOutput)
}

type LookupBookmarkOutputArgs struct {
	BookmarkId        pulumi.StringInput `pulumi:"bookmarkId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupBookmarkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBookmarkArgs)(nil)).Elem()
}


type LookupBookmarkResultOutput struct{ *pulumi.OutputState }

func (LookupBookmarkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBookmarkResult)(nil)).Elem()
}

func (o LookupBookmarkResultOutput) ToLookupBookmarkResultOutput() LookupBookmarkResultOutput {
	return o
}

func (o LookupBookmarkResultOutput) ToLookupBookmarkResultOutputWithContext(ctx context.Context) LookupBookmarkResultOutput {
	return o
}

func (o LookupBookmarkResultOutput) Created() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBookmarkResult) *string { return v.Created }).(pulumi.StringPtrOutput)
}

func (o LookupBookmarkResultOutput) CreatedBy() UserInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupBookmarkResult) *UserInfoResponse { return v.CreatedBy }).(UserInfoResponsePtrOutput)
}

func (o LookupBookmarkResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBookmarkResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupBookmarkResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBookmarkResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupBookmarkResultOutput) EventTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBookmarkResult) *string { return v.EventTime }).(pulumi.StringPtrOutput)
}

func (o LookupBookmarkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBookmarkResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBookmarkResultOutput) IncidentInfo() IncidentInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupBookmarkResult) *IncidentInfoResponse { return v.IncidentInfo }).(IncidentInfoResponsePtrOutput)
}

func (o LookupBookmarkResultOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupBookmarkResult) []string { return v.Labels }).(pulumi.StringArrayOutput)
}

func (o LookupBookmarkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBookmarkResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBookmarkResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBookmarkResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupBookmarkResultOutput) Query() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBookmarkResult) string { return v.Query }).(pulumi.StringOutput)
}

func (o LookupBookmarkResultOutput) QueryEndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBookmarkResult) *string { return v.QueryEndTime }).(pulumi.StringPtrOutput)
}

func (o LookupBookmarkResultOutput) QueryResult() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBookmarkResult) *string { return v.QueryResult }).(pulumi.StringPtrOutput)
}

func (o LookupBookmarkResultOutput) QueryStartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBookmarkResult) *string { return v.QueryStartTime }).(pulumi.StringPtrOutput)
}

func (o LookupBookmarkResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupBookmarkResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupBookmarkResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBookmarkResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupBookmarkResultOutput) Updated() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBookmarkResult) *string { return v.Updated }).(pulumi.StringPtrOutput)
}

func (o LookupBookmarkResultOutput) UpdatedBy() UserInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupBookmarkResult) *UserInfoResponse { return v.UpdatedBy }).(UserInfoResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBookmarkResultOutput{})
}
