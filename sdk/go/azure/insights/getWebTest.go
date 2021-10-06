


package insights

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebTest(ctx *pulumi.Context, args *LookupWebTestArgs, opts ...pulumi.InvokeOption) (*LookupWebTestResult, error) {
	var rv LookupWebTestResult
	err := ctx.Invoke("azure-native:insights:getWebTest", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebTestArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WebTestName       string `pulumi:"webTestName"`
}


type LookupWebTestResult struct {
	Configuration      *WebTestPropertiesResponseConfiguration `pulumi:"configuration"`
	Description        *string                                 `pulumi:"description"`
	Enabled            *bool                                   `pulumi:"enabled"`
	Frequency          *int                                    `pulumi:"frequency"`
	Id                 string                                  `pulumi:"id"`
	Kind               *string                                 `pulumi:"kind"`
	Location           string                                  `pulumi:"location"`
	Locations          []WebTestGeolocationResponse            `pulumi:"locations"`
	Name               string                                  `pulumi:"name"`
	ProvisioningState  string                                  `pulumi:"provisioningState"`
	RetryEnabled       *bool                                   `pulumi:"retryEnabled"`
	SyntheticMonitorId string                                  `pulumi:"syntheticMonitorId"`
	Tags               map[string]string                       `pulumi:"tags"`
	Timeout            *int                                    `pulumi:"timeout"`
	Type               string                                  `pulumi:"type"`
	WebTestKind        string                                  `pulumi:"webTestKind"`
	WebTestName        string                                  `pulumi:"webTestName"`
}
