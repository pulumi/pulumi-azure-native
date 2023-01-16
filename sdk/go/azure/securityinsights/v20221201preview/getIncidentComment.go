


package v20221201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIncidentComment(ctx *pulumi.Context, args *LookupIncidentCommentArgs, opts ...pulumi.InvokeOption) (*LookupIncidentCommentResult, error) {
	var rv LookupIncidentCommentResult
	err := ctx.Invoke("azure-native:securityinsights/v20221201preview:getIncidentComment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIncidentCommentArgs struct {
	IncidentCommentId string `pulumi:"incidentCommentId"`
	IncidentId        string `pulumi:"incidentId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupIncidentCommentResult struct {
	Author              ClientInfoResponse `pulumi:"author"`
	CreatedTimeUtc      string             `pulumi:"createdTimeUtc"`
	Etag                *string            `pulumi:"etag"`
	Id                  string             `pulumi:"id"`
	LastModifiedTimeUtc string             `pulumi:"lastModifiedTimeUtc"`
	Message             string             `pulumi:"message"`
	Name                string             `pulumi:"name"`
	SystemData          SystemDataResponse `pulumi:"systemData"`
	Type                string             `pulumi:"type"`
}

func LookupIncidentCommentOutput(ctx *pulumi.Context, args LookupIncidentCommentOutputArgs, opts ...pulumi.InvokeOption) LookupIncidentCommentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIncidentCommentResult, error) {
			args := v.(LookupIncidentCommentArgs)
			r, err := LookupIncidentComment(ctx, &args, opts...)
			var s LookupIncidentCommentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIncidentCommentResultOutput)
}

type LookupIncidentCommentOutputArgs struct {
	IncidentCommentId pulumi.StringInput `pulumi:"incidentCommentId"`
	IncidentId        pulumi.StringInput `pulumi:"incidentId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupIncidentCommentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIncidentCommentArgs)(nil)).Elem()
}


type LookupIncidentCommentResultOutput struct{ *pulumi.OutputState }

func (LookupIncidentCommentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIncidentCommentResult)(nil)).Elem()
}

func (o LookupIncidentCommentResultOutput) ToLookupIncidentCommentResultOutput() LookupIncidentCommentResultOutput {
	return o
}

func (o LookupIncidentCommentResultOutput) ToLookupIncidentCommentResultOutputWithContext(ctx context.Context) LookupIncidentCommentResultOutput {
	return o
}

func (o LookupIncidentCommentResultOutput) Author() ClientInfoResponseOutput {
	return o.ApplyT(func(v LookupIncidentCommentResult) ClientInfoResponse { return v.Author }).(ClientInfoResponseOutput)
}

func (o LookupIncidentCommentResultOutput) CreatedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentCommentResult) string { return v.CreatedTimeUtc }).(pulumi.StringOutput)
}

func (o LookupIncidentCommentResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIncidentCommentResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupIncidentCommentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentCommentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIncidentCommentResultOutput) LastModifiedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentCommentResult) string { return v.LastModifiedTimeUtc }).(pulumi.StringOutput)
}

func (o LookupIncidentCommentResultOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentCommentResult) string { return v.Message }).(pulumi.StringOutput)
}

func (o LookupIncidentCommentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentCommentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIncidentCommentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupIncidentCommentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupIncidentCommentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentCommentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIncidentCommentResultOutput{})
}
