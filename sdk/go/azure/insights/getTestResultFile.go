


package insights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetTestResultFile(ctx *pulumi.Context, args *GetTestResultFileArgs, opts ...pulumi.InvokeOption) (*GetTestResultFileResult, error) {
	var rv GetTestResultFileResult
	err := ctx.Invoke("azure-native:insights:getTestResultFile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetTestResultFileArgs struct {
	ContinuationToken      *string `pulumi:"continuationToken"`
	DownloadAs             string  `pulumi:"downloadAs"`
	GeoLocationId          string  `pulumi:"geoLocationId"`
	ResourceGroupName      string  `pulumi:"resourceGroupName"`
	TestSuccessfulCriteria *bool   `pulumi:"testSuccessfulCriteria"`
	TimeStamp              int     `pulumi:"timeStamp"`
	WebTestName            string  `pulumi:"webTestName"`
}


type GetTestResultFileResult struct {
	Data     *string `pulumi:"data"`
	NextLink *string `pulumi:"nextLink"`
}

func GetTestResultFileOutput(ctx *pulumi.Context, args GetTestResultFileOutputArgs, opts ...pulumi.InvokeOption) GetTestResultFileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetTestResultFileResult, error) {
			args := v.(GetTestResultFileArgs)
			r, err := GetTestResultFile(ctx, &args, opts...)
			var s GetTestResultFileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetTestResultFileResultOutput)
}

type GetTestResultFileOutputArgs struct {
	ContinuationToken      pulumi.StringPtrInput `pulumi:"continuationToken"`
	DownloadAs             pulumi.StringInput    `pulumi:"downloadAs"`
	GeoLocationId          pulumi.StringInput    `pulumi:"geoLocationId"`
	ResourceGroupName      pulumi.StringInput    `pulumi:"resourceGroupName"`
	TestSuccessfulCriteria pulumi.BoolPtrInput   `pulumi:"testSuccessfulCriteria"`
	TimeStamp              pulumi.IntInput       `pulumi:"timeStamp"`
	WebTestName            pulumi.StringInput    `pulumi:"webTestName"`
}

func (GetTestResultFileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTestResultFileArgs)(nil)).Elem()
}


type GetTestResultFileResultOutput struct{ *pulumi.OutputState }

func (GetTestResultFileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTestResultFileResult)(nil)).Elem()
}

func (o GetTestResultFileResultOutput) ToGetTestResultFileResultOutput() GetTestResultFileResultOutput {
	return o
}

func (o GetTestResultFileResultOutput) ToGetTestResultFileResultOutputWithContext(ctx context.Context) GetTestResultFileResultOutput {
	return o
}

func (o GetTestResultFileResultOutput) Data() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTestResultFileResult) *string { return v.Data }).(pulumi.StringPtrOutput)
}

func (o GetTestResultFileResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTestResultFileResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTestResultFileResultOutput{})
}
