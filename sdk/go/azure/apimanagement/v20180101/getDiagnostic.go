


package v20180101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDiagnostic(ctx *pulumi.Context, args *LookupDiagnosticArgs, opts ...pulumi.InvokeOption) (*LookupDiagnosticResult, error) {
	var rv LookupDiagnosticResult
	err := ctx.Invoke("azure-native:apimanagement/v20180101:getDiagnostic", args, &rv, opts...)
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
	Enabled bool   `pulumi:"enabled"`
	Id      string `pulumi:"id"`
	Name    string `pulumi:"name"`
	Type    string `pulumi:"type"`
}
