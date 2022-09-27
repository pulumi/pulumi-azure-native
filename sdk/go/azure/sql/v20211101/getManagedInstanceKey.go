


package v20211101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagedInstanceKey(ctx *pulumi.Context, args *LookupManagedInstanceKeyArgs, opts ...pulumi.InvokeOption) (*LookupManagedInstanceKeyResult, error) {
	var rv LookupManagedInstanceKeyResult
	err := ctx.Invoke("azure-native:sql/v20211101:getManagedInstanceKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedInstanceKeyArgs struct {
	KeyName             string `pulumi:"keyName"`
	ManagedInstanceName string `pulumi:"managedInstanceName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}


type LookupManagedInstanceKeyResult struct {
	AutoRotationEnabled bool   `pulumi:"autoRotationEnabled"`
	CreationDate        string `pulumi:"creationDate"`
	Id                  string `pulumi:"id"`
	Kind                string `pulumi:"kind"`
	Name                string `pulumi:"name"`
	Thumbprint          string `pulumi:"thumbprint"`
	Type                string `pulumi:"type"`
}

func LookupManagedInstanceKeyOutput(ctx *pulumi.Context, args LookupManagedInstanceKeyOutputArgs, opts ...pulumi.InvokeOption) LookupManagedInstanceKeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagedInstanceKeyResult, error) {
			args := v.(LookupManagedInstanceKeyArgs)
			r, err := LookupManagedInstanceKey(ctx, &args, opts...)
			var s LookupManagedInstanceKeyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagedInstanceKeyResultOutput)
}

type LookupManagedInstanceKeyOutputArgs struct {
	KeyName             pulumi.StringInput `pulumi:"keyName"`
	ManagedInstanceName pulumi.StringInput `pulumi:"managedInstanceName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupManagedInstanceKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedInstanceKeyArgs)(nil)).Elem()
}


type LookupManagedInstanceKeyResultOutput struct{ *pulumi.OutputState }

func (LookupManagedInstanceKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedInstanceKeyResult)(nil)).Elem()
}

func (o LookupManagedInstanceKeyResultOutput) ToLookupManagedInstanceKeyResultOutput() LookupManagedInstanceKeyResultOutput {
	return o
}

func (o LookupManagedInstanceKeyResultOutput) ToLookupManagedInstanceKeyResultOutputWithContext(ctx context.Context) LookupManagedInstanceKeyResultOutput {
	return o
}

func (o LookupManagedInstanceKeyResultOutput) AutoRotationEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupManagedInstanceKeyResult) bool { return v.AutoRotationEnabled }).(pulumi.BoolOutput)
}

func (o LookupManagedInstanceKeyResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstanceKeyResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

func (o LookupManagedInstanceKeyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstanceKeyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupManagedInstanceKeyResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstanceKeyResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupManagedInstanceKeyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstanceKeyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupManagedInstanceKeyResultOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstanceKeyResult) string { return v.Thumbprint }).(pulumi.StringOutput)
}

func (o LookupManagedInstanceKeyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstanceKeyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedInstanceKeyResultOutput{})
}
