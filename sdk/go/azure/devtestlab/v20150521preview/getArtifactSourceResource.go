


package v20150521preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupArtifactSourceResource(ctx *pulumi.Context, args *LookupArtifactSourceResourceArgs, opts ...pulumi.InvokeOption) (*LookupArtifactSourceResourceResult, error) {
	var rv LookupArtifactSourceResourceResult
	err := ctx.Invoke("azure-native:devtestlab/v20150521preview:getArtifactSourceResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupArtifactSourceResourceArgs struct {
	LabName           string `pulumi:"labName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupArtifactSourceResourceResult struct {
	BranchRef         *string           `pulumi:"branchRef"`
	DisplayName       *string           `pulumi:"displayName"`
	FolderPath        *string           `pulumi:"folderPath"`
	Id                *string           `pulumi:"id"`
	Location          *string           `pulumi:"location"`
	Name              *string           `pulumi:"name"`
	ProvisioningState *string           `pulumi:"provisioningState"`
	SecurityToken     *string           `pulumi:"securityToken"`
	SourceType        *string           `pulumi:"sourceType"`
	Status            *string           `pulumi:"status"`
	Tags              map[string]string `pulumi:"tags"`
	Type              *string           `pulumi:"type"`
	Uri               *string           `pulumi:"uri"`
}

func LookupArtifactSourceResourceOutput(ctx *pulumi.Context, args LookupArtifactSourceResourceOutputArgs, opts ...pulumi.InvokeOption) LookupArtifactSourceResourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupArtifactSourceResourceResult, error) {
			args := v.(LookupArtifactSourceResourceArgs)
			r, err := LookupArtifactSourceResource(ctx, &args, opts...)
			var s LookupArtifactSourceResourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupArtifactSourceResourceResultOutput)
}

type LookupArtifactSourceResourceOutputArgs struct {
	LabName           pulumi.StringInput `pulumi:"labName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupArtifactSourceResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupArtifactSourceResourceArgs)(nil)).Elem()
}


type LookupArtifactSourceResourceResultOutput struct{ *pulumi.OutputState }

func (LookupArtifactSourceResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupArtifactSourceResourceResult)(nil)).Elem()
}

func (o LookupArtifactSourceResourceResultOutput) ToLookupArtifactSourceResourceResultOutput() LookupArtifactSourceResourceResultOutput {
	return o
}

func (o LookupArtifactSourceResourceResultOutput) ToLookupArtifactSourceResourceResultOutputWithContext(ctx context.Context) LookupArtifactSourceResourceResultOutput {
	return o
}

func (o LookupArtifactSourceResourceResultOutput) BranchRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArtifactSourceResourceResult) *string { return v.BranchRef }).(pulumi.StringPtrOutput)
}

func (o LookupArtifactSourceResourceResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArtifactSourceResourceResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupArtifactSourceResourceResultOutput) FolderPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArtifactSourceResourceResult) *string { return v.FolderPath }).(pulumi.StringPtrOutput)
}

func (o LookupArtifactSourceResourceResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArtifactSourceResourceResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupArtifactSourceResourceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArtifactSourceResourceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupArtifactSourceResourceResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArtifactSourceResourceResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupArtifactSourceResourceResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArtifactSourceResourceResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupArtifactSourceResourceResultOutput) SecurityToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArtifactSourceResourceResult) *string { return v.SecurityToken }).(pulumi.StringPtrOutput)
}

func (o LookupArtifactSourceResourceResultOutput) SourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArtifactSourceResourceResult) *string { return v.SourceType }).(pulumi.StringPtrOutput)
}

func (o LookupArtifactSourceResourceResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArtifactSourceResourceResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o LookupArtifactSourceResourceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupArtifactSourceResourceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupArtifactSourceResourceResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArtifactSourceResourceResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o LookupArtifactSourceResourceResultOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArtifactSourceResourceResult) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupArtifactSourceResourceResultOutput{})
}
