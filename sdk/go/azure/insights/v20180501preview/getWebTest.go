


package v20180501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebTest(ctx *pulumi.Context, args *LookupWebTestArgs, opts ...pulumi.InvokeOption) (*LookupWebTestResult, error) {
	var rv LookupWebTestResult
	err := ctx.Invoke("azure-native:insights/v20180501preview:getWebTest", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupWebTestArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WebTestName       string `pulumi:"webTestName"`
}


type LookupWebTestResult struct {
	Configuration      *WebTestPropertiesResponseConfiguration   `pulumi:"configuration"`
	Description        *string                                   `pulumi:"description"`
	Enabled            *bool                                     `pulumi:"enabled"`
	Frequency          *int                                      `pulumi:"frequency"`
	Id                 string                                    `pulumi:"id"`
	Kind               *string                                   `pulumi:"kind"`
	Location           string                                    `pulumi:"location"`
	Locations          []WebTestGeolocationResponse              `pulumi:"locations"`
	Name               string                                    `pulumi:"name"`
	ProvisioningState  string                                    `pulumi:"provisioningState"`
	Request            *WebTestPropertiesResponseRequest         `pulumi:"request"`
	RetryEnabled       *bool                                     `pulumi:"retryEnabled"`
	SyntheticMonitorId string                                    `pulumi:"syntheticMonitorId"`
	Tags               map[string]string                         `pulumi:"tags"`
	Timeout            *int                                      `pulumi:"timeout"`
	Type               string                                    `pulumi:"type"`
	ValidationRules    *WebTestPropertiesResponseValidationRules `pulumi:"validationRules"`
	WebTestKind        string                                    `pulumi:"webTestKind"`
	WebTestName        string                                    `pulumi:"webTestName"`
}


func (val *LookupWebTestResult) Defaults() *LookupWebTestResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Frequency) {
		frequency_ := 300
		tmp.Frequency = &frequency_
	}
	if isZero(tmp.Kind) {
		kind_ := "ping"
		tmp.Kind = &kind_
	}
	if isZero(tmp.Timeout) {
		timeout_ := 30
		tmp.Timeout = &timeout_
	}
	if isZero(tmp.WebTestKind) {
		tmp.WebTestKind = "ping"
	}
	return &tmp
}

func LookupWebTestOutput(ctx *pulumi.Context, args LookupWebTestOutputArgs, opts ...pulumi.InvokeOption) LookupWebTestResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebTestResult, error) {
			args := v.(LookupWebTestArgs)
			r, err := LookupWebTest(ctx, &args, opts...)
			var s LookupWebTestResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebTestResultOutput)
}

type LookupWebTestOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WebTestName       pulumi.StringInput `pulumi:"webTestName"`
}

func (LookupWebTestOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebTestArgs)(nil)).Elem()
}


type LookupWebTestResultOutput struct{ *pulumi.OutputState }

func (LookupWebTestResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebTestResult)(nil)).Elem()
}

func (o LookupWebTestResultOutput) ToLookupWebTestResultOutput() LookupWebTestResultOutput {
	return o
}

func (o LookupWebTestResultOutput) ToLookupWebTestResultOutputWithContext(ctx context.Context) LookupWebTestResultOutput {
	return o
}

func (o LookupWebTestResultOutput) Configuration() WebTestPropertiesResponseConfigurationPtrOutput {
	return o.ApplyT(func(v LookupWebTestResult) *WebTestPropertiesResponseConfiguration { return v.Configuration }).(WebTestPropertiesResponseConfigurationPtrOutput)
}

func (o LookupWebTestResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebTestResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupWebTestResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebTestResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o LookupWebTestResultOutput) Frequency() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupWebTestResult) *int { return v.Frequency }).(pulumi.IntPtrOutput)
}

func (o LookupWebTestResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebTestResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWebTestResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebTestResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupWebTestResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebTestResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupWebTestResultOutput) Locations() WebTestGeolocationResponseArrayOutput {
	return o.ApplyT(func(v LookupWebTestResult) []WebTestGeolocationResponse { return v.Locations }).(WebTestGeolocationResponseArrayOutput)
}

func (o LookupWebTestResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebTestResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWebTestResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebTestResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupWebTestResultOutput) Request() WebTestPropertiesResponseRequestPtrOutput {
	return o.ApplyT(func(v LookupWebTestResult) *WebTestPropertiesResponseRequest { return v.Request }).(WebTestPropertiesResponseRequestPtrOutput)
}

func (o LookupWebTestResultOutput) RetryEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebTestResult) *bool { return v.RetryEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupWebTestResultOutput) SyntheticMonitorId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebTestResult) string { return v.SyntheticMonitorId }).(pulumi.StringOutput)
}

func (o LookupWebTestResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWebTestResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupWebTestResultOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupWebTestResult) *int { return v.Timeout }).(pulumi.IntPtrOutput)
}

func (o LookupWebTestResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebTestResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupWebTestResultOutput) ValidationRules() WebTestPropertiesResponseValidationRulesPtrOutput {
	return o.ApplyT(func(v LookupWebTestResult) *WebTestPropertiesResponseValidationRules { return v.ValidationRules }).(WebTestPropertiesResponseValidationRulesPtrOutput)
}

func (o LookupWebTestResultOutput) WebTestKind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebTestResult) string { return v.WebTestKind }).(pulumi.StringOutput)
}

func (o LookupWebTestResultOutput) WebTestName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebTestResult) string { return v.WebTestName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebTestResultOutput{})
}
