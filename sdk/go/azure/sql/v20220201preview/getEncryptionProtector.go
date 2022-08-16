


package v20220201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEncryptionProtector(ctx *pulumi.Context, args *LookupEncryptionProtectorArgs, opts ...pulumi.InvokeOption) (*LookupEncryptionProtectorResult, error) {
	var rv LookupEncryptionProtectorResult
	err := ctx.Invoke("azure-native:sql/v20220201preview:getEncryptionProtector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEncryptionProtectorArgs struct {
	EncryptionProtectorName string `pulumi:"encryptionProtectorName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
	ServerName              string `pulumi:"serverName"`
}


type LookupEncryptionProtectorResult struct {
	AutoRotationEnabled *bool   `pulumi:"autoRotationEnabled"`
	Id                  string  `pulumi:"id"`
	Kind                string  `pulumi:"kind"`
	Location            string  `pulumi:"location"`
	Name                string  `pulumi:"name"`
	ServerKeyName       *string `pulumi:"serverKeyName"`
	ServerKeyType       string  `pulumi:"serverKeyType"`
	Subregion           string  `pulumi:"subregion"`
	Thumbprint          string  `pulumi:"thumbprint"`
	Type                string  `pulumi:"type"`
	Uri                 string  `pulumi:"uri"`
}

func LookupEncryptionProtectorOutput(ctx *pulumi.Context, args LookupEncryptionProtectorOutputArgs, opts ...pulumi.InvokeOption) LookupEncryptionProtectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEncryptionProtectorResult, error) {
			args := v.(LookupEncryptionProtectorArgs)
			r, err := LookupEncryptionProtector(ctx, &args, opts...)
			var s LookupEncryptionProtectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEncryptionProtectorResultOutput)
}

type LookupEncryptionProtectorOutputArgs struct {
	EncryptionProtectorName pulumi.StringInput `pulumi:"encryptionProtectorName"`
	ResourceGroupName       pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerName              pulumi.StringInput `pulumi:"serverName"`
}

func (LookupEncryptionProtectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEncryptionProtectorArgs)(nil)).Elem()
}


type LookupEncryptionProtectorResultOutput struct{ *pulumi.OutputState }

func (LookupEncryptionProtectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEncryptionProtectorResult)(nil)).Elem()
}

func (o LookupEncryptionProtectorResultOutput) ToLookupEncryptionProtectorResultOutput() LookupEncryptionProtectorResultOutput {
	return o
}

func (o LookupEncryptionProtectorResultOutput) ToLookupEncryptionProtectorResultOutputWithContext(ctx context.Context) LookupEncryptionProtectorResultOutput {
	return o
}

func (o LookupEncryptionProtectorResultOutput) AutoRotationEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupEncryptionProtectorResult) *bool { return v.AutoRotationEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupEncryptionProtectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEncryptionProtectorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupEncryptionProtectorResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEncryptionProtectorResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupEncryptionProtectorResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEncryptionProtectorResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupEncryptionProtectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEncryptionProtectorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupEncryptionProtectorResultOutput) ServerKeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEncryptionProtectorResult) *string { return v.ServerKeyName }).(pulumi.StringPtrOutput)
}

func (o LookupEncryptionProtectorResultOutput) ServerKeyType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEncryptionProtectorResult) string { return v.ServerKeyType }).(pulumi.StringOutput)
}

func (o LookupEncryptionProtectorResultOutput) Subregion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEncryptionProtectorResult) string { return v.Subregion }).(pulumi.StringOutput)
}

func (o LookupEncryptionProtectorResultOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEncryptionProtectorResult) string { return v.Thumbprint }).(pulumi.StringOutput)
}

func (o LookupEncryptionProtectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEncryptionProtectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupEncryptionProtectorResultOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEncryptionProtectorResult) string { return v.Uri }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEncryptionProtectorResultOutput{})
}
