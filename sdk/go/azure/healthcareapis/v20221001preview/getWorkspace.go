


package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkspace(ctx *pulumi.Context, args *LookupWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceResult, error) {
	var rv LookupWorkspaceResult
	err := ctx.Invoke("azure-native:healthcareapis/v20221001preview:getWorkspace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspaceArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupWorkspaceResult struct {
	Etag       *string                     `pulumi:"etag"`
	Id         string                      `pulumi:"id"`
	Location   *string                     `pulumi:"location"`
	Name       string                      `pulumi:"name"`
	Properties WorkspaceResponseProperties `pulumi:"properties"`
	SystemData SystemDataResponse          `pulumi:"systemData"`
	Tags       map[string]string           `pulumi:"tags"`
	Type       string                      `pulumi:"type"`
}

func LookupWorkspaceOutput(ctx *pulumi.Context, args LookupWorkspaceOutputArgs, opts ...pulumi.InvokeOption) LookupWorkspaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkspaceResult, error) {
			args := v.(LookupWorkspaceArgs)
			r, err := LookupWorkspace(ctx, &args, opts...)
			var s LookupWorkspaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkspaceResultOutput)
}

type LookupWorkspaceOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupWorkspaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceArgs)(nil)).Elem()
}


type LookupWorkspaceResultOutput struct{ *pulumi.OutputState }

func (LookupWorkspaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceResult)(nil)).Elem()
}

func (o LookupWorkspaceResultOutput) ToLookupWorkspaceResultOutput() LookupWorkspaceResultOutput {
	return o
}

func (o LookupWorkspaceResultOutput) ToLookupWorkspaceResultOutputWithContext(ctx context.Context) LookupWorkspaceResultOutput {
	return o
}

func (o LookupWorkspaceResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) Properties() WorkspaceResponsePropertiesOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) WorkspaceResponseProperties { return v.Properties }).(WorkspaceResponsePropertiesOutput)
}

func (o LookupWorkspaceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupWorkspaceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupWorkspaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspaceResultOutput{})
}
