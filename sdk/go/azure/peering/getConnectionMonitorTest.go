


package peering

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConnectionMonitorTest(ctx *pulumi.Context, args *LookupConnectionMonitorTestArgs, opts ...pulumi.InvokeOption) (*LookupConnectionMonitorTestResult, error) {
	var rv LookupConnectionMonitorTestResult
	err := ctx.Invoke("azure-native:peering:getConnectionMonitorTest", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectionMonitorTestArgs struct {
	ConnectionMonitorTestName string `pulumi:"connectionMonitorTestName"`
	PeeringServiceName        string `pulumi:"peeringServiceName"`
	ResourceGroupName         string `pulumi:"resourceGroupName"`
}


type LookupConnectionMonitorTestResult struct {
	Destination        *string  `pulumi:"destination"`
	DestinationPort    *int     `pulumi:"destinationPort"`
	Id                 string   `pulumi:"id"`
	IsTestSuccessful   bool     `pulumi:"isTestSuccessful"`
	Name               string   `pulumi:"name"`
	Path               []string `pulumi:"path"`
	ProvisioningState  string   `pulumi:"provisioningState"`
	SourceAgent        *string  `pulumi:"sourceAgent"`
	TestFrequencyInSec *int     `pulumi:"testFrequencyInSec"`
	Type               string   `pulumi:"type"`
}
