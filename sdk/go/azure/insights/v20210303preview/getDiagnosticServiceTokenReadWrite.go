


package v20210303preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDiagnosticServiceTokenReadWrite(ctx *pulumi.Context, args *GetDiagnosticServiceTokenReadWriteArgs, opts ...pulumi.InvokeOption) (*GetDiagnosticServiceTokenReadWriteResult, error) {
	var rv GetDiagnosticServiceTokenReadWriteResult
	err := ctx.Invoke("azure-native:insights/v20210303preview:getDiagnosticServiceTokenReadWrite", args, &rv, opts...)
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

func GetDiagnosticServiceTokenReadWriteOutput(ctx *pulumi.Context, args GetDiagnosticServiceTokenReadWriteOutputArgs, opts ...pulumi.InvokeOption) GetDiagnosticServiceTokenReadWriteResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDiagnosticServiceTokenReadWriteResult, error) {
			args := v.(GetDiagnosticServiceTokenReadWriteArgs)
			r, err := GetDiagnosticServiceTokenReadWrite(ctx, &args, opts...)
			var s GetDiagnosticServiceTokenReadWriteResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDiagnosticServiceTokenReadWriteResultOutput)
}

type GetDiagnosticServiceTokenReadWriteOutputArgs struct {
	ResourceUri pulumi.StringInput `pulumi:"resourceUri"`
}

func (GetDiagnosticServiceTokenReadWriteOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDiagnosticServiceTokenReadWriteArgs)(nil)).Elem()
}


type GetDiagnosticServiceTokenReadWriteResultOutput struct{ *pulumi.OutputState }

func (GetDiagnosticServiceTokenReadWriteResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDiagnosticServiceTokenReadWriteResult)(nil)).Elem()
}

func (o GetDiagnosticServiceTokenReadWriteResultOutput) ToGetDiagnosticServiceTokenReadWriteResultOutput() GetDiagnosticServiceTokenReadWriteResultOutput {
	return o
}

func (o GetDiagnosticServiceTokenReadWriteResultOutput) ToGetDiagnosticServiceTokenReadWriteResultOutputWithContext(ctx context.Context) GetDiagnosticServiceTokenReadWriteResultOutput {
	return o
}

func (o GetDiagnosticServiceTokenReadWriteResultOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDiagnosticServiceTokenReadWriteResult) *string { return v.Token }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDiagnosticServiceTokenReadWriteResultOutput{})
}
