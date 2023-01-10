


package v20220801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebPubSubHub(ctx *pulumi.Context, args *LookupWebPubSubHubArgs, opts ...pulumi.InvokeOption) (*LookupWebPubSubHubResult, error) {
	var rv LookupWebPubSubHubResult
	err := ctx.Invoke("azure-native:webpubsub/v20220801preview:getWebPubSubHub", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupWebPubSubHubArgs struct {
	HubName           string `pulumi:"hubName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupWebPubSubHubResult struct {
	Id         string                         `pulumi:"id"`
	Name       string                         `pulumi:"name"`
	Properties WebPubSubHubPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse             `pulumi:"systemData"`
	Type       string                         `pulumi:"type"`
}


func (val *LookupWebPubSubHubResult) Defaults() *LookupWebPubSubHubResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}

func LookupWebPubSubHubOutput(ctx *pulumi.Context, args LookupWebPubSubHubOutputArgs, opts ...pulumi.InvokeOption) LookupWebPubSubHubResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebPubSubHubResult, error) {
			args := v.(LookupWebPubSubHubArgs)
			r, err := LookupWebPubSubHub(ctx, &args, opts...)
			var s LookupWebPubSubHubResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebPubSubHubResultOutput)
}

type LookupWebPubSubHubOutputArgs struct {
	HubName           pulumi.StringInput `pulumi:"hubName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupWebPubSubHubOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebPubSubHubArgs)(nil)).Elem()
}


type LookupWebPubSubHubResultOutput struct{ *pulumi.OutputState }

func (LookupWebPubSubHubResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebPubSubHubResult)(nil)).Elem()
}

func (o LookupWebPubSubHubResultOutput) ToLookupWebPubSubHubResultOutput() LookupWebPubSubHubResultOutput {
	return o
}

func (o LookupWebPubSubHubResultOutput) ToLookupWebPubSubHubResultOutputWithContext(ctx context.Context) LookupWebPubSubHubResultOutput {
	return o
}

func (o LookupWebPubSubHubResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubHubResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWebPubSubHubResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubHubResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWebPubSubHubResultOutput) Properties() WebPubSubHubPropertiesResponseOutput {
	return o.ApplyT(func(v LookupWebPubSubHubResult) WebPubSubHubPropertiesResponse { return v.Properties }).(WebPubSubHubPropertiesResponseOutput)
}

func (o LookupWebPubSubHubResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupWebPubSubHubResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupWebPubSubHubResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubHubResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebPubSubHubResultOutput{})
}
