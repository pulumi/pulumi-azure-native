


package v20170201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupRedisLinkedServer(ctx *pulumi.Context, args *LookupRedisLinkedServerArgs, opts ...pulumi.InvokeOption) (*LookupRedisLinkedServerResult, error) {
	var rv LookupRedisLinkedServerResult
	err := ctx.Invoke("azure-native:cache/v20170201:getRedisLinkedServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRedisLinkedServerArgs struct {
	LinkedServerName  string `pulumi:"linkedServerName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupRedisLinkedServerResult struct {
	Id                       string `pulumi:"id"`
	LinkedRedisCacheId       string `pulumi:"linkedRedisCacheId"`
	LinkedRedisCacheLocation string `pulumi:"linkedRedisCacheLocation"`
	Name                     string `pulumi:"name"`
	ProvisioningState        string `pulumi:"provisioningState"`
	ServerRole               string `pulumi:"serverRole"`
	Type                     string `pulumi:"type"`
}

func LookupRedisLinkedServerOutput(ctx *pulumi.Context, args LookupRedisLinkedServerOutputArgs, opts ...pulumi.InvokeOption) LookupRedisLinkedServerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRedisLinkedServerResult, error) {
			args := v.(LookupRedisLinkedServerArgs)
			r, err := LookupRedisLinkedServer(ctx, &args, opts...)
			var s LookupRedisLinkedServerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRedisLinkedServerResultOutput)
}

type LookupRedisLinkedServerOutputArgs struct {
	LinkedServerName  pulumi.StringInput `pulumi:"linkedServerName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupRedisLinkedServerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRedisLinkedServerArgs)(nil)).Elem()
}


type LookupRedisLinkedServerResultOutput struct{ *pulumi.OutputState }

func (LookupRedisLinkedServerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRedisLinkedServerResult)(nil)).Elem()
}

func (o LookupRedisLinkedServerResultOutput) ToLookupRedisLinkedServerResultOutput() LookupRedisLinkedServerResultOutput {
	return o
}

func (o LookupRedisLinkedServerResultOutput) ToLookupRedisLinkedServerResultOutputWithContext(ctx context.Context) LookupRedisLinkedServerResultOutput {
	return o
}

func (o LookupRedisLinkedServerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisLinkedServerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRedisLinkedServerResultOutput) LinkedRedisCacheId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisLinkedServerResult) string { return v.LinkedRedisCacheId }).(pulumi.StringOutput)
}

func (o LookupRedisLinkedServerResultOutput) LinkedRedisCacheLocation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisLinkedServerResult) string { return v.LinkedRedisCacheLocation }).(pulumi.StringOutput)
}

func (o LookupRedisLinkedServerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisLinkedServerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRedisLinkedServerResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisLinkedServerResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupRedisLinkedServerResultOutput) ServerRole() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisLinkedServerResult) string { return v.ServerRole }).(pulumi.StringOutput)
}

func (o LookupRedisLinkedServerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisLinkedServerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRedisLinkedServerResultOutput{})
}
