


package v20200601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGroup(ctx *pulumi.Context, args *LookupGroupArgs, opts ...pulumi.InvokeOption) (*LookupGroupResult, error) {
	var rv LookupGroupResult
	err := ctx.Invoke("azure-native:apimanagement/v20200601preview:getGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGroupArgs struct {
	GroupId           string `pulumi:"groupId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupGroupResult struct {
	BuiltIn     bool    `pulumi:"builtIn"`
	Description *string `pulumi:"description"`
	DisplayName string  `pulumi:"displayName"`
	ExternalId  *string `pulumi:"externalId"`
	Id          string  `pulumi:"id"`
	Name        string  `pulumi:"name"`
	Type        string  `pulumi:"type"`
}

func LookupGroupOutput(ctx *pulumi.Context, args LookupGroupOutputArgs, opts ...pulumi.InvokeOption) LookupGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGroupResult, error) {
			args := v.(LookupGroupArgs)
			r, err := LookupGroup(ctx, &args, opts...)
			var s LookupGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGroupResultOutput)
}

type LookupGroupOutputArgs struct {
	GroupId           pulumi.StringInput `pulumi:"groupId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupArgs)(nil)).Elem()
}


type LookupGroupResultOutput struct{ *pulumi.OutputState }

func (LookupGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupResult)(nil)).Elem()
}

func (o LookupGroupResultOutput) ToLookupGroupResultOutput() LookupGroupResultOutput {
	return o
}

func (o LookupGroupResultOutput) ToLookupGroupResultOutputWithContext(ctx context.Context) LookupGroupResultOutput {
	return o
}

func (o LookupGroupResultOutput) BuiltIn() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupResult) bool { return v.BuiltIn }).(pulumi.BoolOutput)
}

func (o LookupGroupResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGroupResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupGroupResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupGroupResultOutput) ExternalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGroupResult) *string { return v.ExternalId }).(pulumi.StringPtrOutput)
}

func (o LookupGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGroupResultOutput{})
}
