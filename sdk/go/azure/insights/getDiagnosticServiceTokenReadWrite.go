


package insights

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDiagnosticServiceTokenReadWrite(ctx *pulumi.Context, args *GetDiagnosticServiceTokenReadWriteArgs, opts ...pulumi.InvokeOption) (*GetDiagnosticServiceTokenReadWriteResult, error) {
	var rv GetDiagnosticServiceTokenReadWriteResult
	err := ctx.Invoke("azure-native:insights:getDiagnosticServiceTokenReadWrite", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetDiagnosticServiceTokenReadWriteArgs struct {
	ResourceUri string `pulumi:"resourceUri"`
}


type GetDiagnosticServiceTokenReadWriteResult struct {
	Token *string `pulumi:"token"`
}
