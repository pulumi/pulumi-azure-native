


package v20211101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVideo(ctx *pulumi.Context, args *LookupVideoArgs, opts ...pulumi.InvokeOption) (*LookupVideoResult, error) {
	var rv LookupVideoResult
	err := ctx.Invoke("azure-native:videoanalyzer/v20211101preview:getVideo", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVideoArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VideoName         string `pulumi:"videoName"`
}


type LookupVideoResult struct {
	Archival    *VideoArchivalResponse   `pulumi:"archival"`
	ContentUrls VideoContentUrlsResponse `pulumi:"contentUrls"`
	Description *string                  `pulumi:"description"`
	Flags       VideoFlagsResponse       `pulumi:"flags"`
	Id          string                   `pulumi:"id"`
	MediaInfo   *VideoMediaInfoResponse  `pulumi:"mediaInfo"`
	Name        string                   `pulumi:"name"`
	SystemData  SystemDataResponse       `pulumi:"systemData"`
	Title       *string                  `pulumi:"title"`
	Type        string                   `pulumi:"type"`
}

func LookupVideoOutput(ctx *pulumi.Context, args LookupVideoOutputArgs, opts ...pulumi.InvokeOption) LookupVideoResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVideoResult, error) {
			args := v.(LookupVideoArgs)
			r, err := LookupVideo(ctx, &args, opts...)
			var s LookupVideoResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVideoResultOutput)
}

type LookupVideoOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	VideoName         pulumi.StringInput `pulumi:"videoName"`
}

func (LookupVideoOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVideoArgs)(nil)).Elem()
}


type LookupVideoResultOutput struct{ *pulumi.OutputState }

func (LookupVideoResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVideoResult)(nil)).Elem()
}

func (o LookupVideoResultOutput) ToLookupVideoResultOutput() LookupVideoResultOutput {
	return o
}

func (o LookupVideoResultOutput) ToLookupVideoResultOutputWithContext(ctx context.Context) LookupVideoResultOutput {
	return o
}

func (o LookupVideoResultOutput) Archival() VideoArchivalResponsePtrOutput {
	return o.ApplyT(func(v LookupVideoResult) *VideoArchivalResponse { return v.Archival }).(VideoArchivalResponsePtrOutput)
}

func (o LookupVideoResultOutput) ContentUrls() VideoContentUrlsResponseOutput {
	return o.ApplyT(func(v LookupVideoResult) VideoContentUrlsResponse { return v.ContentUrls }).(VideoContentUrlsResponseOutput)
}

func (o LookupVideoResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVideoResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupVideoResultOutput) Flags() VideoFlagsResponseOutput {
	return o.ApplyT(func(v LookupVideoResult) VideoFlagsResponse { return v.Flags }).(VideoFlagsResponseOutput)
}

func (o LookupVideoResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVideoResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVideoResultOutput) MediaInfo() VideoMediaInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupVideoResult) *VideoMediaInfoResponse { return v.MediaInfo }).(VideoMediaInfoResponsePtrOutput)
}

func (o LookupVideoResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVideoResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVideoResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupVideoResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupVideoResultOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVideoResult) *string { return v.Title }).(pulumi.StringPtrOutput)
}

func (o LookupVideoResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVideoResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVideoResultOutput{})
}
