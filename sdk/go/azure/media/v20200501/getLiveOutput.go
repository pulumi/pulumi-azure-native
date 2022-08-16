


package v20200501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLiveOutput(ctx *pulumi.Context, args *LookupLiveOutputArgs, opts ...pulumi.InvokeOption) (*LookupLiveOutputResult, error) {
	var rv LookupLiveOutputResult
	err := ctx.Invoke("azure-native:media/v20200501:getLiveOutput", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLiveOutputArgs struct {
	AccountName       string `pulumi:"accountName"`
	LiveEventName     string `pulumi:"liveEventName"`
	LiveOutputName    string `pulumi:"liveOutputName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupLiveOutputResult struct {
	ArchiveWindowLength string       `pulumi:"archiveWindowLength"`
	AssetName           string       `pulumi:"assetName"`
	Created             string       `pulumi:"created"`
	Description         *string      `pulumi:"description"`
	Hls                 *HlsResponse `pulumi:"hls"`
	Id                  string       `pulumi:"id"`
	LastModified        string       `pulumi:"lastModified"`
	ManifestName        *string      `pulumi:"manifestName"`
	Name                string       `pulumi:"name"`
	OutputSnapTime      *float64     `pulumi:"outputSnapTime"`
	ProvisioningState   string       `pulumi:"provisioningState"`
	ResourceState       string       `pulumi:"resourceState"`
	Type                string       `pulumi:"type"`
}

func LookupLiveOutputOutput(ctx *pulumi.Context, args LookupLiveOutputOutputArgs, opts ...pulumi.InvokeOption) LookupLiveOutputResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLiveOutputResult, error) {
			args := v.(LookupLiveOutputArgs)
			r, err := LookupLiveOutput(ctx, &args, opts...)
			var s LookupLiveOutputResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLiveOutputResultOutput)
}

type LookupLiveOutputOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	LiveEventName     pulumi.StringInput `pulumi:"liveEventName"`
	LiveOutputName    pulumi.StringInput `pulumi:"liveOutputName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupLiveOutputOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLiveOutputArgs)(nil)).Elem()
}


type LookupLiveOutputResultOutput struct{ *pulumi.OutputState }

func (LookupLiveOutputResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLiveOutputResult)(nil)).Elem()
}

func (o LookupLiveOutputResultOutput) ToLookupLiveOutputResultOutput() LookupLiveOutputResultOutput {
	return o
}

func (o LookupLiveOutputResultOutput) ToLookupLiveOutputResultOutputWithContext(ctx context.Context) LookupLiveOutputResultOutput {
	return o
}

func (o LookupLiveOutputResultOutput) ArchiveWindowLength() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLiveOutputResult) string { return v.ArchiveWindowLength }).(pulumi.StringOutput)
}

func (o LookupLiveOutputResultOutput) AssetName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLiveOutputResult) string { return v.AssetName }).(pulumi.StringOutput)
}

func (o LookupLiveOutputResultOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLiveOutputResult) string { return v.Created }).(pulumi.StringOutput)
}

func (o LookupLiveOutputResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLiveOutputResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupLiveOutputResultOutput) Hls() HlsResponsePtrOutput {
	return o.ApplyT(func(v LookupLiveOutputResult) *HlsResponse { return v.Hls }).(HlsResponsePtrOutput)
}

func (o LookupLiveOutputResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLiveOutputResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLiveOutputResultOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLiveOutputResult) string { return v.LastModified }).(pulumi.StringOutput)
}

func (o LookupLiveOutputResultOutput) ManifestName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLiveOutputResult) *string { return v.ManifestName }).(pulumi.StringPtrOutput)
}

func (o LookupLiveOutputResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLiveOutputResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupLiveOutputResultOutput) OutputSnapTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupLiveOutputResult) *float64 { return v.OutputSnapTime }).(pulumi.Float64PtrOutput)
}

func (o LookupLiveOutputResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLiveOutputResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupLiveOutputResultOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLiveOutputResult) string { return v.ResourceState }).(pulumi.StringOutput)
}

func (o LookupLiveOutputResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLiveOutputResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLiveOutputResultOutput{})
}
