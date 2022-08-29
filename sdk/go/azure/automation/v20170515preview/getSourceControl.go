


package v20170515preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSourceControl(ctx *pulumi.Context, args *LookupSourceControlArgs, opts ...pulumi.InvokeOption) (*LookupSourceControlResult, error) {
	var rv LookupSourceControlResult
	err := ctx.Invoke("azure-native:automation/v20170515preview:getSourceControl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSourceControlArgs struct {
	AutomationAccountName string `pulumi:"automationAccountName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	SourceControlName     string `pulumi:"sourceControlName"`
}


type LookupSourceControlResult struct {
	AutoSync         *bool   `pulumi:"autoSync"`
	Branch           *string `pulumi:"branch"`
	CreationTime     *string `pulumi:"creationTime"`
	Description      *string `pulumi:"description"`
	FolderPath       *string `pulumi:"folderPath"`
	Id               string  `pulumi:"id"`
	LastModifiedTime *string `pulumi:"lastModifiedTime"`
	Name             string  `pulumi:"name"`
	PublishRunbook   *bool   `pulumi:"publishRunbook"`
	RepoUrl          *string `pulumi:"repoUrl"`
	SourceType       *string `pulumi:"sourceType"`
	Type             string  `pulumi:"type"`
}

func LookupSourceControlOutput(ctx *pulumi.Context, args LookupSourceControlOutputArgs, opts ...pulumi.InvokeOption) LookupSourceControlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceControlResult, error) {
			args := v.(LookupSourceControlArgs)
			r, err := LookupSourceControl(ctx, &args, opts...)
			var s LookupSourceControlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceControlResultOutput)
}

type LookupSourceControlOutputArgs struct {
	AutomationAccountName pulumi.StringInput `pulumi:"automationAccountName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	SourceControlName     pulumi.StringInput `pulumi:"sourceControlName"`
}

func (LookupSourceControlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceControlArgs)(nil)).Elem()
}


type LookupSourceControlResultOutput struct{ *pulumi.OutputState }

func (LookupSourceControlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceControlResult)(nil)).Elem()
}

func (o LookupSourceControlResultOutput) ToLookupSourceControlResultOutput() LookupSourceControlResultOutput {
	return o
}

func (o LookupSourceControlResultOutput) ToLookupSourceControlResultOutputWithContext(ctx context.Context) LookupSourceControlResultOutput {
	return o
}

func (o LookupSourceControlResultOutput) AutoSync() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSourceControlResult) *bool { return v.AutoSync }).(pulumi.BoolPtrOutput)
}

func (o LookupSourceControlResultOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceControlResult) *string { return v.Branch }).(pulumi.StringPtrOutput)
}

func (o LookupSourceControlResultOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceControlResult) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

func (o LookupSourceControlResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceControlResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupSourceControlResultOutput) FolderPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceControlResult) *string { return v.FolderPath }).(pulumi.StringPtrOutput)
}

func (o LookupSourceControlResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceControlResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceControlResultOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceControlResult) *string { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

func (o LookupSourceControlResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceControlResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceControlResultOutput) PublishRunbook() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSourceControlResult) *bool { return v.PublishRunbook }).(pulumi.BoolPtrOutput)
}

func (o LookupSourceControlResultOutput) RepoUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceControlResult) *string { return v.RepoUrl }).(pulumi.StringPtrOutput)
}

func (o LookupSourceControlResultOutput) SourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceControlResult) *string { return v.SourceType }).(pulumi.StringPtrOutput)
}

func (o LookupSourceControlResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceControlResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceControlResultOutput{})
}
