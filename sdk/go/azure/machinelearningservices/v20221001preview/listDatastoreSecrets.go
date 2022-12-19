


package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListDatastoreSecrets(ctx *pulumi.Context, args *ListDatastoreSecretsArgs, opts ...pulumi.InvokeOption) (*ListDatastoreSecretsResult, error) {
	var rv ListDatastoreSecretsResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20221001preview:listDatastoreSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDatastoreSecretsArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type ListDatastoreSecretsResult struct {
	SecretsType string `pulumi:"secretsType"`
}

func ListDatastoreSecretsOutput(ctx *pulumi.Context, args ListDatastoreSecretsOutputArgs, opts ...pulumi.InvokeOption) ListDatastoreSecretsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListDatastoreSecretsResult, error) {
			args := v.(ListDatastoreSecretsArgs)
			r, err := ListDatastoreSecrets(ctx, &args, opts...)
			var s ListDatastoreSecretsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListDatastoreSecretsResultOutput)
}

type ListDatastoreSecretsOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (ListDatastoreSecretsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDatastoreSecretsArgs)(nil)).Elem()
}


type ListDatastoreSecretsResultOutput struct{ *pulumi.OutputState }

func (ListDatastoreSecretsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDatastoreSecretsResult)(nil)).Elem()
}

func (o ListDatastoreSecretsResultOutput) ToListDatastoreSecretsResultOutput() ListDatastoreSecretsResultOutput {
	return o
}

func (o ListDatastoreSecretsResultOutput) ToListDatastoreSecretsResultOutputWithContext(ctx context.Context) ListDatastoreSecretsResultOutput {
	return o
}

func (o ListDatastoreSecretsResultOutput) SecretsType() pulumi.StringOutput {
	return o.ApplyT(func(v ListDatastoreSecretsResult) string { return v.SecretsType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDatastoreSecretsResultOutput{})
}
