


package media

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStreamingEndpoint(ctx *pulumi.Context, args *LookupStreamingEndpointArgs, opts ...pulumi.InvokeOption) (*LookupStreamingEndpointResult, error) {
	var rv LookupStreamingEndpointResult
	err := ctx.Invoke("azure-native:media:getStreamingEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStreamingEndpointArgs struct {
	AccountName           string `pulumi:"accountName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	StreamingEndpointName string `pulumi:"streamingEndpointName"`
}


type LookupStreamingEndpointResult struct {
	AccessControl           *StreamingEndpointAccessControlResponse `pulumi:"accessControl"`
	AvailabilitySetName     *string                                 `pulumi:"availabilitySetName"`
	CdnEnabled              *bool                                   `pulumi:"cdnEnabled"`
	CdnProfile              *string                                 `pulumi:"cdnProfile"`
	CdnProvider             *string                                 `pulumi:"cdnProvider"`
	Created                 string                                  `pulumi:"created"`
	CrossSiteAccessPolicies *CrossSiteAccessPoliciesResponse        `pulumi:"crossSiteAccessPolicies"`
	CustomHostNames         []string                                `pulumi:"customHostNames"`
	Description             *string                                 `pulumi:"description"`
	FreeTrialEndTime        string                                  `pulumi:"freeTrialEndTime"`
	HostName                string                                  `pulumi:"hostName"`
	Id                      string                                  `pulumi:"id"`
	LastModified            string                                  `pulumi:"lastModified"`
	Location                string                                  `pulumi:"location"`
	MaxCacheAge             *float64                                `pulumi:"maxCacheAge"`
	Name                    string                                  `pulumi:"name"`
	ProvisioningState       string                                  `pulumi:"provisioningState"`
	ResourceState           string                                  `pulumi:"resourceState"`
	ScaleUnits              int                                     `pulumi:"scaleUnits"`
	SystemData              SystemDataResponse                      `pulumi:"systemData"`
	Tags                    map[string]string                       `pulumi:"tags"`
	Type                    string                                  `pulumi:"type"`
}

func LookupStreamingEndpointOutput(ctx *pulumi.Context, args LookupStreamingEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupStreamingEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStreamingEndpointResult, error) {
			args := v.(LookupStreamingEndpointArgs)
			r, err := LookupStreamingEndpoint(ctx, &args, opts...)
			var s LookupStreamingEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStreamingEndpointResultOutput)
}

type LookupStreamingEndpointOutputArgs struct {
	AccountName           pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	StreamingEndpointName pulumi.StringInput `pulumi:"streamingEndpointName"`
}

func (LookupStreamingEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStreamingEndpointArgs)(nil)).Elem()
}


type LookupStreamingEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupStreamingEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStreamingEndpointResult)(nil)).Elem()
}

func (o LookupStreamingEndpointResultOutput) ToLookupStreamingEndpointResultOutput() LookupStreamingEndpointResultOutput {
	return o
}

func (o LookupStreamingEndpointResultOutput) ToLookupStreamingEndpointResultOutputWithContext(ctx context.Context) LookupStreamingEndpointResultOutput {
	return o
}

func (o LookupStreamingEndpointResultOutput) AccessControl() StreamingEndpointAccessControlResponsePtrOutput {
	return o.ApplyT(func(v LookupStreamingEndpointResult) *StreamingEndpointAccessControlResponse { return v.AccessControl }).(StreamingEndpointAccessControlResponsePtrOutput)
}

func (o LookupStreamingEndpointResultOutput) AvailabilitySetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStreamingEndpointResult) *string { return v.AvailabilitySetName }).(pulumi.StringPtrOutput)
}

func (o LookupStreamingEndpointResultOutput) CdnEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupStreamingEndpointResult) *bool { return v.CdnEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupStreamingEndpointResultOutput) CdnProfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStreamingEndpointResult) *string { return v.CdnProfile }).(pulumi.StringPtrOutput)
}

func (o LookupStreamingEndpointResultOutput) CdnProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStreamingEndpointResult) *string { return v.CdnProvider }).(pulumi.StringPtrOutput)
}

func (o LookupStreamingEndpointResultOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamingEndpointResult) string { return v.Created }).(pulumi.StringOutput)
}

func (o LookupStreamingEndpointResultOutput) CrossSiteAccessPolicies() CrossSiteAccessPoliciesResponsePtrOutput {
	return o.ApplyT(func(v LookupStreamingEndpointResult) *CrossSiteAccessPoliciesResponse {
		return v.CrossSiteAccessPolicies
	}).(CrossSiteAccessPoliciesResponsePtrOutput)
}

func (o LookupStreamingEndpointResultOutput) CustomHostNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupStreamingEndpointResult) []string { return v.CustomHostNames }).(pulumi.StringArrayOutput)
}

func (o LookupStreamingEndpointResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStreamingEndpointResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupStreamingEndpointResultOutput) FreeTrialEndTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamingEndpointResult) string { return v.FreeTrialEndTime }).(pulumi.StringOutput)
}

func (o LookupStreamingEndpointResultOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamingEndpointResult) string { return v.HostName }).(pulumi.StringOutput)
}

func (o LookupStreamingEndpointResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamingEndpointResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupStreamingEndpointResultOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamingEndpointResult) string { return v.LastModified }).(pulumi.StringOutput)
}

func (o LookupStreamingEndpointResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamingEndpointResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupStreamingEndpointResultOutput) MaxCacheAge() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupStreamingEndpointResult) *float64 { return v.MaxCacheAge }).(pulumi.Float64PtrOutput)
}

func (o LookupStreamingEndpointResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamingEndpointResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupStreamingEndpointResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamingEndpointResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupStreamingEndpointResultOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamingEndpointResult) string { return v.ResourceState }).(pulumi.StringOutput)
}

func (o LookupStreamingEndpointResultOutput) ScaleUnits() pulumi.IntOutput {
	return o.ApplyT(func(v LookupStreamingEndpointResult) int { return v.ScaleUnits }).(pulumi.IntOutput)
}

func (o LookupStreamingEndpointResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupStreamingEndpointResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupStreamingEndpointResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupStreamingEndpointResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupStreamingEndpointResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamingEndpointResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStreamingEndpointResultOutput{})
}
