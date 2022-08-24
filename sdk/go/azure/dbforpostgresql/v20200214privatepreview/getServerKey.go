


package v20200214privatepreview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServerKey(ctx *pulumi.Context, args *LookupServerKeyArgs, opts ...pulumi.InvokeOption) (*LookupServerKeyResult, error) {
	var rv LookupServerKeyResult
	err := ctx.Invoke("azure-native:dbforpostgresql/v20200214privatepreview:getServerKey", args, &rv, opts...)
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
	CreationDate  string  `pulumi:"creationDate"`
	Id            string  `pulumi:"id"`
	Kind          string  `pulumi:"kind"`
	Name          string  `pulumi:"name"`
	ServerKeyType string  `pulumi:"serverKeyType"`
	Type          string  `pulumi:"type"`
	Uri           *string `pulumi:"uri"`
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

func (o LookupServerKeyResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerKeyResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

func (o LookupServerKeyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerKeyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupServerKeyResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerKeyResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupServerKeyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerKeyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupServerKeyResultOutput) ServerKeyType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerKeyResult) string { return v.ServerKeyType }).(pulumi.StringOutput)
}

func (o LookupServerKeyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerKeyResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupServerKeyResultOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerKeyResult) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerKeyResultOutput{})
}
