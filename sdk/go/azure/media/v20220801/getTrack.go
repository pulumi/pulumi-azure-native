


package v20220801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTrack(ctx *pulumi.Context, args *LookupTrackArgs, opts ...pulumi.InvokeOption) (*LookupTrackResult, error) {
	var rv LookupTrackResult
	err := ctx.Invoke("azure-native:media/v20220801:getTrack", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTrackArgs struct {
	AccountName       string `pulumi:"accountName"`
	AssetName         string `pulumi:"assetName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	TrackName         string `pulumi:"trackName"`
}


type LookupTrackResult struct {
	Id                string      `pulumi:"id"`
	Name              string      `pulumi:"name"`
	ProvisioningState string      `pulumi:"provisioningState"`
	Track             interface{} `pulumi:"track"`
	Type              string      `pulumi:"type"`
}

func LookupTrackOutput(ctx *pulumi.Context, args LookupTrackOutputArgs, opts ...pulumi.InvokeOption) LookupTrackResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTrackResult, error) {
			args := v.(LookupTrackArgs)
			r, err := LookupTrack(ctx, &args, opts...)
			var s LookupTrackResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTrackResultOutput)
}

type LookupTrackOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	AssetName         pulumi.StringInput `pulumi:"assetName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	TrackName         pulumi.StringInput `pulumi:"trackName"`
}

func (LookupTrackOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTrackArgs)(nil)).Elem()
}


type LookupTrackResultOutput struct{ *pulumi.OutputState }

func (LookupTrackResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTrackResult)(nil)).Elem()
}

func (o LookupTrackResultOutput) ToLookupTrackResultOutput() LookupTrackResultOutput {
	return o
}

func (o LookupTrackResultOutput) ToLookupTrackResultOutputWithContext(ctx context.Context) LookupTrackResultOutput {
	return o
}

func (o LookupTrackResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTrackResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTrackResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTrackResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTrackResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTrackResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupTrackResultOutput) Track() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupTrackResult) interface{} { return v.Track }).(pulumi.AnyOutput)
}

func (o LookupTrackResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTrackResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTrackResultOutput{})
}
