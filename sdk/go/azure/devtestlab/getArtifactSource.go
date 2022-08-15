


package devtestlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupArtifactSource(ctx *pulumi.Context, args *LookupArtifactSourceArgs, opts ...pulumi.InvokeOption) (*LookupArtifactSourceResult, error) {
	var rv LookupArtifactSourceResult
	err := ctx.Invoke("azure-native:devtestlab:getArtifactSource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupArtifactSourceArgs struct {
	Expand            *string `pulumi:"expand"`
	LabName           string  `pulumi:"labName"`
	Name              string  `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupArtifactSourceResult struct {
	ArmTemplateFolderPath *string           `pulumi:"armTemplateFolderPath"`
	BranchRef             *string           `pulumi:"branchRef"`
	CreatedDate           string            `pulumi:"createdDate"`
	DisplayName           *string           `pulumi:"displayName"`
	FolderPath            *string           `pulumi:"folderPath"`
	Id                    string            `pulumi:"id"`
	Location              *string           `pulumi:"location"`
	Name                  string            `pulumi:"name"`
	ProvisioningState     string            `pulumi:"provisioningState"`
	SecurityToken         *string           `pulumi:"securityToken"`
	SourceType            *string           `pulumi:"sourceType"`
	Status                *string           `pulumi:"status"`
	Tags                  map[string]string `pulumi:"tags"`
	Type                  string            `pulumi:"type"`
	UniqueIdentifier      string            `pulumi:"uniqueIdentifier"`
	Uri                   *string           `pulumi:"uri"`
}

func LookupArtifactSourceOutput(ctx *pulumi.Context, args LookupArtifactSourceOutputArgs, opts ...pulumi.InvokeOption) LookupArtifactSourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupArtifactSourceResult, error) {
			args := v.(LookupArtifactSourceArgs)
			r, err := LookupArtifactSource(ctx, &args, opts...)
			var s LookupArtifactSourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupArtifactSourceResultOutput)
}

type LookupArtifactSourceOutputArgs struct {
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	LabName           pulumi.StringInput    `pulumi:"labName"`
	Name              pulumi.StringInput    `pulumi:"name"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupArtifactSourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupArtifactSourceArgs)(nil)).Elem()
}


type LookupArtifactSourceResultOutput struct{ *pulumi.OutputState }

func (LookupArtifactSourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupArtifactSourceResult)(nil)).Elem()
}

func (o LookupArtifactSourceResultOutput) ToLookupArtifactSourceResultOutput() LookupArtifactSourceResultOutput {
	return o
}

func (o LookupArtifactSourceResultOutput) ToLookupArtifactSourceResultOutputWithContext(ctx context.Context) LookupArtifactSourceResultOutput {
	return o
}

func (o LookupArtifactSourceResultOutput) ArmTemplateFolderPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArtifactSourceResult) *string { return v.ArmTemplateFolderPath }).(pulumi.StringPtrOutput)
}

func (o LookupArtifactSourceResultOutput) BranchRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArtifactSourceResult) *string { return v.BranchRef }).(pulumi.StringPtrOutput)
}

func (o LookupArtifactSourceResultOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArtifactSourceResult) string { return v.CreatedDate }).(pulumi.StringOutput)
}

func (o LookupArtifactSourceResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArtifactSourceResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupArtifactSourceResultOutput) FolderPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArtifactSourceResult) *string { return v.FolderPath }).(pulumi.StringPtrOutput)
}

func (o LookupArtifactSourceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArtifactSourceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupArtifactSourceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArtifactSourceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupArtifactSourceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArtifactSourceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupArtifactSourceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArtifactSourceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupArtifactSourceResultOutput) SecurityToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArtifactSourceResult) *string { return v.SecurityToken }).(pulumi.StringPtrOutput)
}

func (o LookupArtifactSourceResultOutput) SourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArtifactSourceResult) *string { return v.SourceType }).(pulumi.StringPtrOutput)
}

func (o LookupArtifactSourceResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArtifactSourceResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o LookupArtifactSourceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupArtifactSourceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupArtifactSourceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArtifactSourceResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupArtifactSourceResultOutput) UniqueIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArtifactSourceResult) string { return v.UniqueIdentifier }).(pulumi.StringOutput)
}

func (o LookupArtifactSourceResultOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArtifactSourceResult) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupArtifactSourceResultOutput{})
}
