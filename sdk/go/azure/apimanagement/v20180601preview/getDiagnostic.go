


package v20180601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDiagnostic(ctx *pulumi.Context, args *LookupDiagnosticArgs, opts ...pulumi.InvokeOption) (*LookupDiagnosticResult, error) {
	var rv LookupDiagnosticResult
	err := ctx.Invoke("azure-native:apimanagement/v20180601preview:getDiagnostic", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDiagnosticArgs struct {
	DiagnosticId      string `pulumi:"diagnosticId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupDiagnosticResult struct {
	AlwaysLog                    *string                             `pulumi:"alwaysLog"`
	Backend                      *PipelineDiagnosticSettingsResponse `pulumi:"backend"`
	EnableHttpCorrelationHeaders *bool                               `pulumi:"enableHttpCorrelationHeaders"`
	Frontend                     *PipelineDiagnosticSettingsResponse `pulumi:"frontend"`
	Id                           string                              `pulumi:"id"`
	LoggerId                     string                              `pulumi:"loggerId"`
	Name                         string                              `pulumi:"name"`
	Sampling                     *SamplingSettingsResponse           `pulumi:"sampling"`
	Type                         string                              `pulumi:"type"`
}
