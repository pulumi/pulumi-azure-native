


package v20210601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRoute(ctx *pulumi.Context, args *LookupRouteArgs, opts ...pulumi.InvokeOption) (*LookupRouteResult, error) {
	var rv LookupRouteResult
	err := ctx.Invoke("azure-native:cdn/v20210601:getRoute", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRouteArgs struct {
	EndpointName      string `pulumi:"endpointName"`
	ProfileName       string `pulumi:"profileName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RouteName         string `pulumi:"routeName"`
}


type LookupRouteResult struct {
	CacheConfiguration  *AfdRouteCacheConfigurationResponse  `pulumi:"cacheConfiguration"`
	CustomDomains       []ActivatedResourceReferenceResponse `pulumi:"customDomains"`
	DeploymentStatus    string                               `pulumi:"deploymentStatus"`
	EnabledState        *string                              `pulumi:"enabledState"`
	EndpointName        string                               `pulumi:"endpointName"`
	ForwardingProtocol  *string                              `pulumi:"forwardingProtocol"`
	HttpsRedirect       *string                              `pulumi:"httpsRedirect"`
	Id                  string                               `pulumi:"id"`
	LinkToDefaultDomain *string                              `pulumi:"linkToDefaultDomain"`
	Name                string                               `pulumi:"name"`
	OriginGroup         ResourceReferenceResponse            `pulumi:"originGroup"`
	OriginPath          *string                              `pulumi:"originPath"`
	PatternsToMatch     []string                             `pulumi:"patternsToMatch"`
	ProvisioningState   string                               `pulumi:"provisioningState"`
	RuleSets            []ResourceReferenceResponse          `pulumi:"ruleSets"`
	SupportedProtocols  []string                             `pulumi:"supportedProtocols"`
	SystemData          SystemDataResponse                   `pulumi:"systemData"`
	Type                string                               `pulumi:"type"`
}

func LookupRouteOutput(ctx *pulumi.Context, args LookupRouteOutputArgs, opts ...pulumi.InvokeOption) LookupRouteResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRouteResult, error) {
			args := v.(LookupRouteArgs)
			r, err := LookupRoute(ctx, &args, opts...)
			var s LookupRouteResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRouteResultOutput)
}

type LookupRouteOutputArgs struct {
	EndpointName      pulumi.StringInput `pulumi:"endpointName"`
	ProfileName       pulumi.StringInput `pulumi:"profileName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	RouteName         pulumi.StringInput `pulumi:"routeName"`
}

func (LookupRouteOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteArgs)(nil)).Elem()
}


type LookupRouteResultOutput struct{ *pulumi.OutputState }

func (LookupRouteResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteResult)(nil)).Elem()
}

func (o LookupRouteResultOutput) ToLookupRouteResultOutput() LookupRouteResultOutput {
	return o
}

func (o LookupRouteResultOutput) ToLookupRouteResultOutputWithContext(ctx context.Context) LookupRouteResultOutput {
	return o
}

func (o LookupRouteResultOutput) CacheConfiguration() AfdRouteCacheConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *AfdRouteCacheConfigurationResponse { return v.CacheConfiguration }).(AfdRouteCacheConfigurationResponsePtrOutput)
}

func (o LookupRouteResultOutput) CustomDomains() ActivatedResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v LookupRouteResult) []ActivatedResourceReferenceResponse { return v.CustomDomains }).(ActivatedResourceReferenceResponseArrayOutput)
}

func (o LookupRouteResultOutput) DeploymentStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.DeploymentStatus }).(pulumi.StringOutput)
}

func (o LookupRouteResultOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *string { return v.EnabledState }).(pulumi.StringPtrOutput)
}

func (o LookupRouteResultOutput) EndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.EndpointName }).(pulumi.StringOutput)
}

func (o LookupRouteResultOutput) ForwardingProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *string { return v.ForwardingProtocol }).(pulumi.StringPtrOutput)
}

func (o LookupRouteResultOutput) HttpsRedirect() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *string { return v.HttpsRedirect }).(pulumi.StringPtrOutput)
}

func (o LookupRouteResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRouteResultOutput) LinkToDefaultDomain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *string { return v.LinkToDefaultDomain }).(pulumi.StringPtrOutput)
}

func (o LookupRouteResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRouteResultOutput) OriginGroup() ResourceReferenceResponseOutput {
	return o.ApplyT(func(v LookupRouteResult) ResourceReferenceResponse { return v.OriginGroup }).(ResourceReferenceResponseOutput)
}

func (o LookupRouteResultOutput) OriginPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *string { return v.OriginPath }).(pulumi.StringPtrOutput)
}

func (o LookupRouteResultOutput) PatternsToMatch() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRouteResult) []string { return v.PatternsToMatch }).(pulumi.StringArrayOutput)
}

func (o LookupRouteResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupRouteResultOutput) RuleSets() ResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v LookupRouteResult) []ResourceReferenceResponse { return v.RuleSets }).(ResourceReferenceResponseArrayOutput)
}

func (o LookupRouteResultOutput) SupportedProtocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRouteResult) []string { return v.SupportedProtocols }).(pulumi.StringArrayOutput)
}

func (o LookupRouteResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupRouteResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupRouteResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouteResultOutput{})
}
