


package v20180501preview

import (
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
