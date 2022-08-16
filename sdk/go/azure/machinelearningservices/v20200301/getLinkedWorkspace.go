


package v20200301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLinkedWorkspace(ctx *pulumi.Context, args *LookupLinkedWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupLinkedWorkspaceResult, error) {
	var rv LookupLinkedWorkspaceResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20200301:getLinkedWorkspace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLinkedWorkspaceArgs struct {
	LinkName          string `pulumi:"linkName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupLinkedWorkspaceResult struct {
	Id         string                       `pulumi:"id"`
	Name       string                       `pulumi:"name"`
	Properties LinkedWorkspacePropsResponse `pulumi:"properties"`
	Type       string                       `pulumi:"type"`
}

func LookupLinkedWorkspaceOutput(ctx *pulumi.Context, args LookupLinkedWorkspaceOutputArgs, opts ...pulumi.InvokeOption) LookupLinkedWorkspaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLinkedWorkspaceResult, error) {
			args := v.(LookupLinkedWorkspaceArgs)
			r, err := LookupLinkedWorkspace(ctx, &args, opts...)
			var s LookupLinkedWorkspaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLinkedWorkspaceResultOutput)
}

type LookupLinkedWorkspaceOutputArgs struct {
	LinkName          pulumi.StringInput `pulumi:"linkName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupLinkedWorkspaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLinkedWorkspaceArgs)(nil)).Elem()
}


type LookupLinkedWorkspaceResultOutput struct{ *pulumi.OutputState }

func (LookupLinkedWorkspaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLinkedWorkspaceResult)(nil)).Elem()
}

func (o LookupLinkedWorkspaceResultOutput) ToLookupLinkedWorkspaceResultOutput() LookupLinkedWorkspaceResultOutput {
	return o
}

func (o LookupLinkedWorkspaceResultOutput) ToLookupLinkedWorkspaceResultOutputWithContext(ctx context.Context) LookupLinkedWorkspaceResultOutput {
	return o
}

func (o LookupLinkedWorkspaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedWorkspaceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLinkedWorkspaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedWorkspaceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupLinkedWorkspaceResultOutput) Properties() LinkedWorkspacePropsResponseOutput {
	return o.ApplyT(func(v LookupLinkedWorkspaceResult) LinkedWorkspacePropsResponse { return v.Properties }).(LinkedWorkspacePropsResponseOutput)
}

func (o LookupLinkedWorkspaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedWorkspaceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLinkedWorkspaceResultOutput{})
}
