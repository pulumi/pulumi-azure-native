


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkflow(ctx *pulumi.Context, args *LookupWorkflowArgs, opts ...pulumi.InvokeOption) (*LookupWorkflowResult, error) {
	var rv LookupWorkflowResult
	err := ctx.Invoke("azure-native:devhub/v20220401preview:getWorkflow", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkflowArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkflowName      string `pulumi:"workflowName"`
}


type LookupWorkflowResult struct {
	Acr                  *ACRResponse                                  `pulumi:"acr"`
	AksResourceId        *string                                       `pulumi:"aksResourceId"`
	AuthStatus           string                                        `pulumi:"authStatus"`
	BranchName           *string                                       `pulumi:"branchName"`
	DeploymentProperties *DeploymentPropertiesResponse                 `pulumi:"deploymentProperties"`
	DockerBuildContext   *string                                       `pulumi:"dockerBuildContext"`
	Dockerfile           *string                                       `pulumi:"dockerfile"`
	Id                   string                                        `pulumi:"id"`
	LastWorkflowRun      *WorkflowRunResponse                          `pulumi:"lastWorkflowRun"`
	Location             string                                        `pulumi:"location"`
	Name                 string                                        `pulumi:"name"`
	Namespace            *string                                       `pulumi:"namespace"`
	OidcCredentials      *GitHubWorkflowProfileResponseOidcCredentials `pulumi:"oidcCredentials"`
	PrStatus             string                                        `pulumi:"prStatus"`
	PrURL                string                                        `pulumi:"prURL"`
	PullNumber           int                                           `pulumi:"pullNumber"`
	RepositoryName       *string                                       `pulumi:"repositoryName"`
	RepositoryOwner      *string                                       `pulumi:"repositoryOwner"`
	SystemData           SystemDataResponse                            `pulumi:"systemData"`
	Tags                 map[string]string                             `pulumi:"tags"`
	Type                 string                                        `pulumi:"type"`
}

func LookupWorkflowOutput(ctx *pulumi.Context, args LookupWorkflowOutputArgs, opts ...pulumi.InvokeOption) LookupWorkflowResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkflowResult, error) {
			args := v.(LookupWorkflowArgs)
			r, err := LookupWorkflow(ctx, &args, opts...)
			var s LookupWorkflowResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkflowResultOutput)
}

type LookupWorkflowOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkflowName      pulumi.StringInput `pulumi:"workflowName"`
}

func (LookupWorkflowOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkflowArgs)(nil)).Elem()
}


type LookupWorkflowResultOutput struct{ *pulumi.OutputState }

func (LookupWorkflowResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkflowResult)(nil)).Elem()
}

func (o LookupWorkflowResultOutput) ToLookupWorkflowResultOutput() LookupWorkflowResultOutput {
	return o
}

func (o LookupWorkflowResultOutput) ToLookupWorkflowResultOutputWithContext(ctx context.Context) LookupWorkflowResultOutput {
	return o
}

func (o LookupWorkflowResultOutput) Acr() ACRResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *ACRResponse { return v.Acr }).(ACRResponsePtrOutput)
}

func (o LookupWorkflowResultOutput) AksResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *string { return v.AksResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupWorkflowResultOutput) AuthStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkflowResult) string { return v.AuthStatus }).(pulumi.StringOutput)
}

func (o LookupWorkflowResultOutput) BranchName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *string { return v.BranchName }).(pulumi.StringPtrOutput)
}

func (o LookupWorkflowResultOutput) DeploymentProperties() DeploymentPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *DeploymentPropertiesResponse { return v.DeploymentProperties }).(DeploymentPropertiesResponsePtrOutput)
}

func (o LookupWorkflowResultOutput) DockerBuildContext() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *string { return v.DockerBuildContext }).(pulumi.StringPtrOutput)
}

func (o LookupWorkflowResultOutput) Dockerfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *string { return v.Dockerfile }).(pulumi.StringPtrOutput)
}

func (o LookupWorkflowResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkflowResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWorkflowResultOutput) LastWorkflowRun() WorkflowRunResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *WorkflowRunResponse { return v.LastWorkflowRun }).(WorkflowRunResponsePtrOutput)
}

func (o LookupWorkflowResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkflowResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupWorkflowResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkflowResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWorkflowResultOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

func (o LookupWorkflowResultOutput) OidcCredentials() GitHubWorkflowProfileResponseOidcCredentialsPtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *GitHubWorkflowProfileResponseOidcCredentials { return v.OidcCredentials }).(GitHubWorkflowProfileResponseOidcCredentialsPtrOutput)
}

func (o LookupWorkflowResultOutput) PrStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkflowResult) string { return v.PrStatus }).(pulumi.StringOutput)
}

func (o LookupWorkflowResultOutput) PrURL() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkflowResult) string { return v.PrURL }).(pulumi.StringOutput)
}

func (o LookupWorkflowResultOutput) PullNumber() pulumi.IntOutput {
	return o.ApplyT(func(v LookupWorkflowResult) int { return v.PullNumber }).(pulumi.IntOutput)
}

func (o LookupWorkflowResultOutput) RepositoryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *string { return v.RepositoryName }).(pulumi.StringPtrOutput)
}

func (o LookupWorkflowResultOutput) RepositoryOwner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *string { return v.RepositoryOwner }).(pulumi.StringPtrOutput)
}

func (o LookupWorkflowResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupWorkflowResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupWorkflowResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWorkflowResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupWorkflowResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkflowResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkflowResultOutput{})
}
