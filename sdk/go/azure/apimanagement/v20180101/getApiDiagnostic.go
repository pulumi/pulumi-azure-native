


package v20180101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApiDiagnostic(ctx *pulumi.Context, args *LookupApiDiagnosticArgs, opts ...pulumi.InvokeOption) (*LookupApiDiagnosticResult, error) {
	var rv LookupApiDiagnosticResult
	err := ctx.Invoke("azure-native:apimanagement/v20180101:getApiDiagnostic", args, &rv, opts...)
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
	Enabled bool   `pulumi:"enabled"`
	Id      string `pulumi:"id"`
	Name    string `pulumi:"name"`
	Type    string `pulumi:"type"`
}
