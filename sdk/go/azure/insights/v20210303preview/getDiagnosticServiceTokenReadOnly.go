


package v20210303preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDiagnosticServiceTokenReadOnly(ctx *pulumi.Context, args *GetDiagnosticServiceTokenReadOnlyArgs, opts ...pulumi.InvokeOption) (*GetDiagnosticServiceTokenReadOnlyResult, error) {
	var rv GetDiagnosticServiceTokenReadOnlyResult
	err := ctx.Invoke("azure-native:insights/v20210303preview:getDiagnosticServiceTokenReadOnly", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetDiagnosticServiceTokenReadOnlyArgs struct {
	ResourceUri string `pulumi:"resourceUri"`
}


type GetDiagnosticServiceTokenReadOnlyResult struct {
	Token *string `pulumi:"token"`
}
