


package v20210101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkspaceConnection(ctx *pulumi.Context, args *LookupWorkspaceConnectionArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceConnectionResult, error) {
	var rv LookupWorkspaceConnectionResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20210101:getWorkspaceConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspaceConnectionArgs struct {
	ConnectionName    string `pulumi:"connectionName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupWorkspaceConnectionResult struct {
	AuthType    *string `pulumi:"authType"`
	Category    *string `pulumi:"category"`
	Id          string  `pulumi:"id"`
	Name        string  `pulumi:"name"`
	Target      *string `pulumi:"target"`
	Type        string  `pulumi:"type"`
	Value       *string `pulumi:"value"`
	ValueFormat *string `pulumi:"valueFormat"`
}

func LookupWorkspaceConnectionOutput(ctx *pulumi.Context, args LookupWorkspaceConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupWorkspaceConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkspaceConnectionResult, error) {
			args := v.(LookupWorkspaceConnectionArgs)
			r, err := LookupWorkspaceConnection(ctx, &args, opts...)
			var s LookupWorkspaceConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkspaceConnectionResultOutput)
}

type LookupWorkspaceConnectionOutputArgs struct {
	ConnectionName    pulumi.StringInput `pulumi:"connectionName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupWorkspaceConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceConnectionArgs)(nil)).Elem()
}


type LookupWorkspaceConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupWorkspaceConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceConnectionResult)(nil)).Elem()
}

func (o LookupWorkspaceConnectionResultOutput) ToLookupWorkspaceConnectionResultOutput() LookupWorkspaceConnectionResultOutput {
	return o
}

func (o LookupWorkspaceConnectionResultOutput) ToLookupWorkspaceConnectionResultOutputWithContext(ctx context.Context) LookupWorkspaceConnectionResultOutput {
	return o
}

func (o LookupWorkspaceConnectionResultOutput) AuthType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceConnectionResult) *string { return v.AuthType }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceConnectionResultOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceConnectionResult) *string { return v.Category }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWorkspaceConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWorkspaceConnectionResultOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceConnectionResult) *string { return v.Target }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupWorkspaceConnectionResultOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceConnectionResult) *string { return v.Value }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceConnectionResultOutput) ValueFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceConnectionResult) *string { return v.ValueFormat }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspaceConnectionResultOutput{})
}
