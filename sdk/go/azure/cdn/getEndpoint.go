


package cdn

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEndpoint(ctx *pulumi.Context, args *LookupEndpointArgs, opts ...pulumi.InvokeOption) (*LookupEndpointResult, error) {
	var rv LookupEndpointResult
	err := ctx.Invoke("azure-native:cdn:getEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEndpointArgs struct {
	EndpointName      string `pulumi:"endpointName"`
	ProfileName       string `pulumi:"profileName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupEndpointResult struct {
	ContentTypesToCompress           []string                                                                    `pulumi:"contentTypesToCompress"`
	DefaultOriginGroup               *ResourceReferenceResponse                                                  `pulumi:"defaultOriginGroup"`
	DeliveryPolicy                   *EndpointPropertiesUpdateParametersResponseDeliveryPolicy                   `pulumi:"deliveryPolicy"`
	GeoFilters                       []GeoFilterResponse                                                         `pulumi:"geoFilters"`
	HostName                         string                                                                      `pulumi:"hostName"`
	Id                               string                                                                      `pulumi:"id"`
	IsCompressionEnabled             *bool                                                                       `pulumi:"isCompressionEnabled"`
	IsHttpAllowed                    *bool                                                                       `pulumi:"isHttpAllowed"`
	IsHttpsAllowed                   *bool                                                                       `pulumi:"isHttpsAllowed"`
	Location                         string                                                                      `pulumi:"location"`
	Name                             string                                                                      `pulumi:"name"`
	OptimizationType                 *string                                                                     `pulumi:"optimizationType"`
	OriginGroups                     []DeepCreatedOriginGroupResponse                                            `pulumi:"originGroups"`
	OriginHostHeader                 *string                                                                     `pulumi:"originHostHeader"`
	OriginPath                       *string                                                                     `pulumi:"originPath"`
	Origins                          []DeepCreatedOriginResponse                                                 `pulumi:"origins"`
	ProbePath                        *string                                                                     `pulumi:"probePath"`
	ProvisioningState                string                                                                      `pulumi:"provisioningState"`
	QueryStringCachingBehavior       *string                                                                     `pulumi:"queryStringCachingBehavior"`
	ResourceState                    string                                                                      `pulumi:"resourceState"`
	SystemData                       SystemDataResponse                                                          `pulumi:"systemData"`
	Tags                             map[string]string                                                           `pulumi:"tags"`
	Type                             string                                                                      `pulumi:"type"`
	UrlSigningKeys                   []UrlSigningKeyResponse                                                     `pulumi:"urlSigningKeys"`
	WebApplicationFirewallPolicyLink *EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLink `pulumi:"webApplicationFirewallPolicyLink"`
}

func LookupEndpointOutput(ctx *pulumi.Context, args LookupEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEndpointResult, error) {
			args := v.(LookupEndpointArgs)
			r, err := LookupEndpoint(ctx, &args, opts...)
			var s LookupEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEndpointResultOutput)
}

type LookupEndpointOutputArgs struct {
	EndpointName      pulumi.StringInput `pulumi:"endpointName"`
	ProfileName       pulumi.StringInput `pulumi:"profileName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEndpointArgs)(nil)).Elem()
}


type LookupEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEndpointResult)(nil)).Elem()
}

func (o LookupEndpointResultOutput) ToLookupEndpointResultOutput() LookupEndpointResultOutput {
	return o
}

func (o LookupEndpointResultOutput) ToLookupEndpointResultOutputWithContext(ctx context.Context) LookupEndpointResultOutput {
	return o
}

func (o LookupEndpointResultOutput) ContentTypesToCompress() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupEndpointResult) []string { return v.ContentTypesToCompress }).(pulumi.StringArrayOutput)
}

func (o LookupEndpointResultOutput) DefaultOriginGroup() ResourceReferenceResponsePtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *ResourceReferenceResponse { return v.DefaultOriginGroup }).(ResourceReferenceResponsePtrOutput)
}

func (o LookupEndpointResultOutput) DeliveryPolicy() EndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *EndpointPropertiesUpdateParametersResponseDeliveryPolicy {
		return v.DeliveryPolicy
	}).(EndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrOutput)
}

func (o LookupEndpointResultOutput) GeoFilters() GeoFilterResponseArrayOutput {
	return o.ApplyT(func(v LookupEndpointResult) []GeoFilterResponse { return v.GeoFilters }).(GeoFilterResponseArrayOutput)
}

func (o LookupEndpointResultOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.HostName }).(pulumi.StringOutput)
}

func (o LookupEndpointResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupEndpointResultOutput) IsCompressionEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *bool { return v.IsCompressionEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupEndpointResultOutput) IsHttpAllowed() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *bool { return v.IsHttpAllowed }).(pulumi.BoolPtrOutput)
}

func (o LookupEndpointResultOutput) IsHttpsAllowed() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *bool { return v.IsHttpsAllowed }).(pulumi.BoolPtrOutput)
}

func (o LookupEndpointResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupEndpointResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupEndpointResultOutput) OptimizationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *string { return v.OptimizationType }).(pulumi.StringPtrOutput)
}

func (o LookupEndpointResultOutput) OriginGroups() DeepCreatedOriginGroupResponseArrayOutput {
	return o.ApplyT(func(v LookupEndpointResult) []DeepCreatedOriginGroupResponse { return v.OriginGroups }).(DeepCreatedOriginGroupResponseArrayOutput)
}

func (o LookupEndpointResultOutput) OriginHostHeader() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *string { return v.OriginHostHeader }).(pulumi.StringPtrOutput)
}

func (o LookupEndpointResultOutput) OriginPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *string { return v.OriginPath }).(pulumi.StringPtrOutput)
}

func (o LookupEndpointResultOutput) Origins() DeepCreatedOriginResponseArrayOutput {
	return o.ApplyT(func(v LookupEndpointResult) []DeepCreatedOriginResponse { return v.Origins }).(DeepCreatedOriginResponseArrayOutput)
}

func (o LookupEndpointResultOutput) ProbePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *string { return v.ProbePath }).(pulumi.StringPtrOutput)
}

func (o LookupEndpointResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupEndpointResultOutput) QueryStringCachingBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *string { return v.QueryStringCachingBehavior }).(pulumi.StringPtrOutput)
}

func (o LookupEndpointResultOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.ResourceState }).(pulumi.StringOutput)
}

func (o LookupEndpointResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupEndpointResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupEndpointResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupEndpointResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupEndpointResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupEndpointResultOutput) UrlSigningKeys() UrlSigningKeyResponseArrayOutput {
	return o.ApplyT(func(v LookupEndpointResult) []UrlSigningKeyResponse { return v.UrlSigningKeys }).(UrlSigningKeyResponseArrayOutput)
}

func (o LookupEndpointResultOutput) WebApplicationFirewallPolicyLink() EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLink {
		return v.WebApplicationFirewallPolicyLink
	}).(EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEndpointResultOutput{})
}
