


package media

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLiveEvent(ctx *pulumi.Context, args *LookupLiveEventArgs, opts ...pulumi.InvokeOption) (*LookupLiveEventResult, error) {
	var rv LookupLiveEventResult
	err := ctx.Invoke("azure-native:media:getLiveEvent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLiveEventArgs struct {
	AccountName       string `pulumi:"accountName"`
	LiveEventName     string `pulumi:"liveEventName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupLiveEventResult struct {
	Created                 string                           `pulumi:"created"`
	CrossSiteAccessPolicies *CrossSiteAccessPoliciesResponse `pulumi:"crossSiteAccessPolicies"`
	Description             *string                          `pulumi:"description"`
	Encoding                *LiveEventEncodingResponse       `pulumi:"encoding"`
	HostnamePrefix          *string                          `pulumi:"hostnamePrefix"`
	Id                      string                           `pulumi:"id"`
	Input                   LiveEventInputResponse           `pulumi:"input"`
	LastModified            string                           `pulumi:"lastModified"`
	Location                string                           `pulumi:"location"`
	Name                    string                           `pulumi:"name"`
	Preview                 *LiveEventPreviewResponse        `pulumi:"preview"`
	ProvisioningState       string                           `pulumi:"provisioningState"`
	ResourceState           string                           `pulumi:"resourceState"`
	StreamOptions           []string                         `pulumi:"streamOptions"`
	SystemData              SystemDataResponse               `pulumi:"systemData"`
	Tags                    map[string]string                `pulumi:"tags"`
	Transcriptions          []LiveEventTranscriptionResponse `pulumi:"transcriptions"`
	Type                    string                           `pulumi:"type"`
	UseStaticHostname       *bool                            `pulumi:"useStaticHostname"`
}

func LookupLiveEventOutput(ctx *pulumi.Context, args LookupLiveEventOutputArgs, opts ...pulumi.InvokeOption) LookupLiveEventResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLiveEventResult, error) {
			args := v.(LookupLiveEventArgs)
			r, err := LookupLiveEvent(ctx, &args, opts...)
			var s LookupLiveEventResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLiveEventResultOutput)
}

type LookupLiveEventOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	LiveEventName     pulumi.StringInput `pulumi:"liveEventName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupLiveEventOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLiveEventArgs)(nil)).Elem()
}


type LookupLiveEventResultOutput struct{ *pulumi.OutputState }

func (LookupLiveEventResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLiveEventResult)(nil)).Elem()
}

func (o LookupLiveEventResultOutput) ToLookupLiveEventResultOutput() LookupLiveEventResultOutput {
	return o
}

func (o LookupLiveEventResultOutput) ToLookupLiveEventResultOutputWithContext(ctx context.Context) LookupLiveEventResultOutput {
	return o
}

func (o LookupLiveEventResultOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLiveEventResult) string { return v.Created }).(pulumi.StringOutput)
}

func (o LookupLiveEventResultOutput) CrossSiteAccessPolicies() CrossSiteAccessPoliciesResponsePtrOutput {
	return o.ApplyT(func(v LookupLiveEventResult) *CrossSiteAccessPoliciesResponse { return v.CrossSiteAccessPolicies }).(CrossSiteAccessPoliciesResponsePtrOutput)
}

func (o LookupLiveEventResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLiveEventResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupLiveEventResultOutput) Encoding() LiveEventEncodingResponsePtrOutput {
	return o.ApplyT(func(v LookupLiveEventResult) *LiveEventEncodingResponse { return v.Encoding }).(LiveEventEncodingResponsePtrOutput)
}

func (o LookupLiveEventResultOutput) HostnamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLiveEventResult) *string { return v.HostnamePrefix }).(pulumi.StringPtrOutput)
}

func (o LookupLiveEventResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLiveEventResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLiveEventResultOutput) Input() LiveEventInputResponseOutput {
	return o.ApplyT(func(v LookupLiveEventResult) LiveEventInputResponse { return v.Input }).(LiveEventInputResponseOutput)
}

func (o LookupLiveEventResultOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLiveEventResult) string { return v.LastModified }).(pulumi.StringOutput)
}

func (o LookupLiveEventResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLiveEventResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupLiveEventResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLiveEventResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupLiveEventResultOutput) Preview() LiveEventPreviewResponsePtrOutput {
	return o.ApplyT(func(v LookupLiveEventResult) *LiveEventPreviewResponse { return v.Preview }).(LiveEventPreviewResponsePtrOutput)
}

func (o LookupLiveEventResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLiveEventResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupLiveEventResultOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLiveEventResult) string { return v.ResourceState }).(pulumi.StringOutput)
}

func (o LookupLiveEventResultOutput) StreamOptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLiveEventResult) []string { return v.StreamOptions }).(pulumi.StringArrayOutput)
}

func (o LookupLiveEventResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupLiveEventResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupLiveEventResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupLiveEventResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupLiveEventResultOutput) Transcriptions() LiveEventTranscriptionResponseArrayOutput {
	return o.ApplyT(func(v LookupLiveEventResult) []LiveEventTranscriptionResponse { return v.Transcriptions }).(LiveEventTranscriptionResponseArrayOutput)
}

func (o LookupLiveEventResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLiveEventResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupLiveEventResultOutput) UseStaticHostname() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLiveEventResult) *bool { return v.UseStaticHostname }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLiveEventResultOutput{})
}
