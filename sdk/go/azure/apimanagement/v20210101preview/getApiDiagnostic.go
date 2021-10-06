


package v20210101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApiDiagnostic(ctx *pulumi.Context, args *LookupApiDiagnosticArgs, opts ...pulumi.InvokeOption) (*LookupApiDiagnosticResult, error) {
	var rv LookupApiDiagnosticResult
	err := ctx.Invoke("azure-native:apimanagement/v20210101preview:getApiDiagnostic", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApiDiagnosticArgs struct {
	ApiId             string `pulumi:"apiId"`
	DiagnosticId      string `pulumi:"diagnosticId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupApiDiagnosticResult struct {
	AlwaysLog               *string                             `pulumi:"alwaysLog"`
	Backend                 *PipelineDiagnosticSettingsResponse `pulumi:"backend"`
	Frontend                *PipelineDiagnosticSettingsResponse `pulumi:"frontend"`
	HttpCorrelationProtocol *string                             `pulumi:"httpCorrelationProtocol"`
	Id                      string                              `pulumi:"id"`
	LogClientIp             *bool                               `pulumi:"logClientIp"`
	LoggerId                string                              `pulumi:"loggerId"`
	Name                    string                              `pulumi:"name"`
	OperationNameFormat     *string                             `pulumi:"operationNameFormat"`
	Sampling                *SamplingSettingsResponse           `pulumi:"sampling"`
	Type                    string                              `pulumi:"type"`
	Verbosity               *string                             `pulumi:"verbosity"`
}
