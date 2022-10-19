


package v20210501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVideo(ctx *pulumi.Context, args *LookupVideoArgs, opts ...pulumi.InvokeOption) (*LookupVideoResult, error) {
	var rv LookupVideoResult
	err := ctx.Invoke("azure-native:videoanalyzer/v20210501preview:getVideo", args, &rv, opts...)
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
	Description *string                `pulumi:"description"`
	Flags       VideoFlagsResponse     `pulumi:"flags"`
	Id          string                 `pulumi:"id"`
	MediaInfo   VideoMediaInfoResponse `pulumi:"mediaInfo"`
	Name        string                 `pulumi:"name"`
	Streaming   VideoStreamingResponse `pulumi:"streaming"`
	SystemData  SystemDataResponse     `pulumi:"systemData"`
	Title       *string                `pulumi:"title"`
	Type        string                 `pulumi:"type"`
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

func (o LookupVideoResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVideoResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupVideoResultOutput) Flags() VideoFlagsResponseOutput {
	return o.ApplyT(func(v LookupVideoResult) VideoFlagsResponse { return v.Flags }).(VideoFlagsResponseOutput)
}

func (o LookupVideoResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVideoResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVideoResultOutput) MediaInfo() VideoMediaInfoResponseOutput {
	return o.ApplyT(func(v LookupVideoResult) VideoMediaInfoResponse { return v.MediaInfo }).(VideoMediaInfoResponseOutput)
}

func (o LookupVideoResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVideoResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVideoResultOutput) Streaming() VideoStreamingResponseOutput {
	return o.ApplyT(func(v LookupVideoResult) VideoStreamingResponse { return v.Streaming }).(VideoStreamingResponseOutput)
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
