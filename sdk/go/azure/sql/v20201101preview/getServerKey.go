


package v20201101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServerKey(ctx *pulumi.Context, args *LookupServerKeyArgs, opts ...pulumi.InvokeOption) (*LookupServerKeyResult, error) {
	var rv LookupServerKeyResult
	err := ctx.Invoke("azure-native:sql/v20201101preview:getServerKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerKeyArgs struct {
	KeyName           string `pulumi:"keyName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type LookupServerKeyResult struct {
	AutoRotationEnabled bool   `pulumi:"autoRotationEnabled"`
	CreationDate        string `pulumi:"creationDate"`
	Id                  string `pulumi:"id"`
	Kind                string `pulumi:"kind"`
	Location            string `pulumi:"location"`
	Name                string `pulumi:"name"`
	Subregion           string `pulumi:"subregion"`
	Thumbprint          string `pulumi:"thumbprint"`
	Type                string `pulumi:"type"`
}

func LookupServerKeyOutput(ctx *pulumi.Context, args LookupServerKeyOutputArgs, opts ...pulumi.InvokeOption) LookupServerKeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServerKeyResult, error) {
			args := v.(LookupServerKeyArgs)
			r, err := LookupServerKey(ctx, &args, opts...)
			var s LookupServerKeyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServerKeyResultOutput)
}

type LookupServerKeyOutputArgs struct {
	KeyName           pulumi.StringInput `pulumi:"keyName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerName        pulumi.StringInput `pulumi:"serverName"`
}

func (LookupServerKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerKeyArgs)(nil)).Elem()
}


type LookupServerKeyResultOutput struct{ *pulumi.OutputState }

func (LookupServerKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerKeyResult)(nil)).Elem()
}

func (o LookupServerKeyResultOutput) ToLookupServerKeyResultOutput() LookupServerKeyResultOutput {
	return o
}

func (o LookupServerKeyResultOutput) ToLookupServerKeyResultOutputWithContext(ctx context.Context) LookupServerKeyResultOutput {
	return o
}

func (o LookupServerKeyResultOutput) AutoRotationEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupServerKeyResult) bool { return v.AutoRotationEnabled }).(pulumi.BoolOutput)
}

func (o LookupServerKeyResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerKeyResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

func (o LookupServerKeyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerKeyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupServerKeyResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerKeyResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupServerKeyResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerKeyResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupServerKeyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerKeyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupServerKeyResultOutput) Subregion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerKeyResult) string { return v.Subregion }).(pulumi.StringOutput)
}

func (o LookupServerKeyResultOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerKeyResult) string { return v.Thumbprint }).(pulumi.StringOutput)
}

func (o LookupServerKeyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerKeyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerKeyResultOutput{})
}
