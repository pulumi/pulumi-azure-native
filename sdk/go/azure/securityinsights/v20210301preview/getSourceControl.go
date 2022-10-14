


package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSourceControl(ctx *pulumi.Context, args *LookupSourceControlArgs, opts ...pulumi.InvokeOption) (*LookupSourceControlResult, error) {
	var rv LookupSourceControlResult
	err := ctx.Invoke("azure-native:securityinsights/v20210301preview:getSourceControl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSourceControlArgs struct {
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string `pulumi:"resourceGroupName"`
	SourceControlId                     string `pulumi:"sourceControlId"`
	WorkspaceName                       string `pulumi:"workspaceName"`
}


type LookupSourceControlResult struct {
	ContentTypes       []string           `pulumi:"contentTypes"`
	CreatedAt          *string            `pulumi:"createdAt"`
	CreatedBy          *string            `pulumi:"createdBy"`
	CreatedByType      *string            `pulumi:"createdByType"`
	Description        *string            `pulumi:"description"`
	DisplayName        string             `pulumi:"displayName"`
	Etag               *string            `pulumi:"etag"`
	Id                 string             `pulumi:"id"`
	LastModifiedAt     *string            `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string            `pulumi:"lastModifiedBy"`
	LastModifiedByType *string            `pulumi:"lastModifiedByType"`
	Name               string             `pulumi:"name"`
	RepoType           string             `pulumi:"repoType"`
	Repository         RepositoryResponse `pulumi:"repository"`
	SystemData         SystemDataResponse `pulumi:"systemData"`
	Type               string             `pulumi:"type"`
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
	OperationalInsightsResourceProvider pulumi.StringInput `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   pulumi.StringInput `pulumi:"resourceGroupName"`
	SourceControlId                     pulumi.StringInput `pulumi:"sourceControlId"`
	WorkspaceName                       pulumi.StringInput `pulumi:"workspaceName"`
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

func (o LookupSourceControlResultOutput) ContentTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSourceControlResult) []string { return v.ContentTypes }).(pulumi.StringArrayOutput)
}

func (o LookupSourceControlResultOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceControlResult) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o LookupSourceControlResultOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceControlResult) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o LookupSourceControlResultOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceControlResult) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o LookupSourceControlResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceControlResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupSourceControlResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceControlResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupSourceControlResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceControlResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupSourceControlResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceControlResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceControlResultOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceControlResult) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o LookupSourceControlResultOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceControlResult) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o LookupSourceControlResultOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceControlResult) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func (o LookupSourceControlResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceControlResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceControlResultOutput) RepoType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceControlResult) string { return v.RepoType }).(pulumi.StringOutput)
}

func (o LookupSourceControlResultOutput) Repository() RepositoryResponseOutput {
	return o.ApplyT(func(v LookupSourceControlResult) RepositoryResponse { return v.Repository }).(RepositoryResponseOutput)
}

func (o LookupSourceControlResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSourceControlResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSourceControlResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceControlResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceControlResultOutput{})
}
