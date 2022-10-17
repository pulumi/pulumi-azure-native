


package v20190701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupLinkedServer(ctx *pulumi.Context, args *LookupLinkedServerArgs, opts ...pulumi.InvokeOption) (*LookupLinkedServerResult, error) {
	var rv LookupLinkedServerResult
	err := ctx.Invoke("azure-native:cache/v20190701:getLinkedServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLinkedServerArgs struct {
	LinkedServerName  string `pulumi:"linkedServerName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupLinkedServerResult struct {
	Id                       string `pulumi:"id"`
	LinkedRedisCacheId       string `pulumi:"linkedRedisCacheId"`
	LinkedRedisCacheLocation string `pulumi:"linkedRedisCacheLocation"`
	Name                     string `pulumi:"name"`
	ProvisioningState        string `pulumi:"provisioningState"`
	ServerRole               string `pulumi:"serverRole"`
	Type                     string `pulumi:"type"`
}

func LookupLinkedServerOutput(ctx *pulumi.Context, args LookupLinkedServerOutputArgs, opts ...pulumi.InvokeOption) LookupLinkedServerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLinkedServerResult, error) {
			args := v.(LookupLinkedServerArgs)
			r, err := LookupLinkedServer(ctx, &args, opts...)
			var s LookupLinkedServerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLinkedServerResultOutput)
}

type LookupLinkedServerOutputArgs struct {
	LinkedServerName  pulumi.StringInput `pulumi:"linkedServerName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupLinkedServerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLinkedServerArgs)(nil)).Elem()
}


type LookupLinkedServerResultOutput struct{ *pulumi.OutputState }

func (LookupLinkedServerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLinkedServerResult)(nil)).Elem()
}

func (o LookupLinkedServerResultOutput) ToLookupLinkedServerResultOutput() LookupLinkedServerResultOutput {
	return o
}

func (o LookupLinkedServerResultOutput) ToLookupLinkedServerResultOutputWithContext(ctx context.Context) LookupLinkedServerResultOutput {
	return o
}

func (o LookupLinkedServerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedServerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLinkedServerResultOutput) LinkedRedisCacheId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedServerResult) string { return v.LinkedRedisCacheId }).(pulumi.StringOutput)
}

func (o LookupLinkedServerResultOutput) LinkedRedisCacheLocation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedServerResult) string { return v.LinkedRedisCacheLocation }).(pulumi.StringOutput)
}

func (o LookupLinkedServerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedServerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupLinkedServerResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedServerResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupLinkedServerResultOutput) ServerRole() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedServerResult) string { return v.ServerRole }).(pulumi.StringOutput)
}

func (o LookupLinkedServerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedServerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLinkedServerResultOutput{})
}
