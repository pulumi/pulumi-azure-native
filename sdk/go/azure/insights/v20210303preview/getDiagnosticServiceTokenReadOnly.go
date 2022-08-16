


package v20210303preview

import (
	"context"
	"reflect"

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

func GetDiagnosticServiceTokenReadOnlyOutput(ctx *pulumi.Context, args GetDiagnosticServiceTokenReadOnlyOutputArgs, opts ...pulumi.InvokeOption) GetDiagnosticServiceTokenReadOnlyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDiagnosticServiceTokenReadOnlyResult, error) {
			args := v.(GetDiagnosticServiceTokenReadOnlyArgs)
			r, err := GetDiagnosticServiceTokenReadOnly(ctx, &args, opts...)
			var s GetDiagnosticServiceTokenReadOnlyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDiagnosticServiceTokenReadOnlyResultOutput)
}

type GetDiagnosticServiceTokenReadOnlyOutputArgs struct {
	ResourceUri pulumi.StringInput `pulumi:"resourceUri"`
}

func (GetDiagnosticServiceTokenReadOnlyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDiagnosticServiceTokenReadOnlyArgs)(nil)).Elem()
}


type GetDiagnosticServiceTokenReadOnlyResultOutput struct{ *pulumi.OutputState }

func (GetDiagnosticServiceTokenReadOnlyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDiagnosticServiceTokenReadOnlyResult)(nil)).Elem()
}

func (o GetDiagnosticServiceTokenReadOnlyResultOutput) ToGetDiagnosticServiceTokenReadOnlyResultOutput() GetDiagnosticServiceTokenReadOnlyResultOutput {
	return o
}

func (o GetDiagnosticServiceTokenReadOnlyResultOutput) ToGetDiagnosticServiceTokenReadOnlyResultOutputWithContext(ctx context.Context) GetDiagnosticServiceTokenReadOnlyResultOutput {
	return o
}

func (o GetDiagnosticServiceTokenReadOnlyResultOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDiagnosticServiceTokenReadOnlyResult) *string { return v.Token }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDiagnosticServiceTokenReadOnlyResultOutput{})
}
