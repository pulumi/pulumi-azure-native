


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Workflow struct {
	pulumi.CustomResourceState

	Acr                  ACRResponsePtrOutput                                  `pulumi:"acr"`
	AksResourceId        pulumi.StringPtrOutput                                `pulumi:"aksResourceId"`
	AuthStatus           pulumi.StringOutput                                   `pulumi:"authStatus"`
	BranchName           pulumi.StringPtrOutput                                `pulumi:"branchName"`
	DeploymentProperties DeploymentPropertiesResponsePtrOutput                 `pulumi:"deploymentProperties"`
	DockerBuildContext   pulumi.StringPtrOutput                                `pulumi:"dockerBuildContext"`
	Dockerfile           pulumi.StringPtrOutput                                `pulumi:"dockerfile"`
	LastWorkflowRun      WorkflowRunResponsePtrOutput                          `pulumi:"lastWorkflowRun"`
	Location             pulumi.StringOutput                                   `pulumi:"location"`
	Name                 pulumi.StringOutput                                   `pulumi:"name"`
	Namespace            pulumi.StringPtrOutput                                `pulumi:"namespace"`
	OidcCredentials      GitHubWorkflowProfileResponseOidcCredentialsPtrOutput `pulumi:"oidcCredentials"`
	PrStatus             pulumi.StringOutput                                   `pulumi:"prStatus"`
	PrURL                pulumi.StringOutput                                   `pulumi:"prURL"`
	PullNumber           pulumi.IntOutput                                      `pulumi:"pullNumber"`
	RepositoryName       pulumi.StringPtrOutput                                `pulumi:"repositoryName"`
	RepositoryOwner      pulumi.StringPtrOutput                                `pulumi:"repositoryOwner"`
	SystemData           SystemDataResponseOutput                              `pulumi:"systemData"`
	Tags                 pulumi.StringMapOutput                                `pulumi:"tags"`
	Type                 pulumi.StringOutput                                   `pulumi:"type"`
}


func NewWorkflow(ctx *pulumi.Context,
	name string, args *WorkflowArgs, opts ...pulumi.ResourceOption) (*Workflow, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devhub:Workflow"),
		},
	})
	opts = append(opts, aliases)
	var resource Workflow
	err := ctx.RegisterResource("azure-native:devhub/v20220401preview:Workflow", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWorkflow(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkflowState, opts ...pulumi.ResourceOption) (*Workflow, error) {
	var resource Workflow
	err := ctx.ReadResource("azure-native:devhub/v20220401preview:Workflow", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type workflowState struct {
}

type WorkflowState struct {
}

func (WorkflowState) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowState)(nil)).Elem()
}

type workflowArgs struct {
	Acr                  *ACR                                  `pulumi:"acr"`
	AksResourceId        *string                               `pulumi:"aksResourceId"`
	BranchName           *string                               `pulumi:"branchName"`
	DeploymentProperties *DeploymentProperties                 `pulumi:"deploymentProperties"`
	DockerBuildContext   *string                               `pulumi:"dockerBuildContext"`
	Dockerfile           *string                               `pulumi:"dockerfile"`
	Location             *string                               `pulumi:"location"`
	Namespace            *string                               `pulumi:"namespace"`
	OidcCredentials      *GitHubWorkflowProfileOidcCredentials `pulumi:"oidcCredentials"`
	RepositoryName       *string                               `pulumi:"repositoryName"`
	RepositoryOwner      *string                               `pulumi:"repositoryOwner"`
	ResourceGroupName    string                                `pulumi:"resourceGroupName"`
	Tags                 map[string]string                     `pulumi:"tags"`
	WorkflowName         *string                               `pulumi:"workflowName"`
}


type WorkflowArgs struct {
	Acr                  ACRPtrInput
	AksResourceId        pulumi.StringPtrInput
	BranchName           pulumi.StringPtrInput
	DeploymentProperties DeploymentPropertiesPtrInput
	DockerBuildContext   pulumi.StringPtrInput
	Dockerfile           pulumi.StringPtrInput
	Location             pulumi.StringPtrInput
	Namespace            pulumi.StringPtrInput
	OidcCredentials      GitHubWorkflowProfileOidcCredentialsPtrInput
	RepositoryName       pulumi.StringPtrInput
	RepositoryOwner      pulumi.StringPtrInput
	ResourceGroupName    pulumi.StringInput
	Tags                 pulumi.StringMapInput
	WorkflowName         pulumi.StringPtrInput
}

func (WorkflowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowArgs)(nil)).Elem()
}

type WorkflowInput interface {
	pulumi.Input

	ToWorkflowOutput() WorkflowOutput
	ToWorkflowOutputWithContext(ctx context.Context) WorkflowOutput
}

func (*Workflow) ElementType() reflect.Type {
	return reflect.TypeOf((**Workflow)(nil)).Elem()
}

func (i *Workflow) ToWorkflowOutput() WorkflowOutput {
	return i.ToWorkflowOutputWithContext(context.Background())
}

func (i *Workflow) ToWorkflowOutputWithContext(ctx context.Context) WorkflowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowOutput)
}

type WorkflowOutput struct{ *pulumi.OutputState }

func (WorkflowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Workflow)(nil)).Elem()
}

func (o WorkflowOutput) ToWorkflowOutput() WorkflowOutput {
	return o
}

func (o WorkflowOutput) ToWorkflowOutputWithContext(ctx context.Context) WorkflowOutput {
	return o
}

func (o WorkflowOutput) Acr() ACRResponsePtrOutput {
	return o.ApplyT(func(v *Workflow) ACRResponsePtrOutput { return v.Acr }).(ACRResponsePtrOutput)
}

func (o WorkflowOutput) AksResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringPtrOutput { return v.AksResourceId }).(pulumi.StringPtrOutput)
}

func (o WorkflowOutput) AuthStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.AuthStatus }).(pulumi.StringOutput)
}

func (o WorkflowOutput) BranchName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringPtrOutput { return v.BranchName }).(pulumi.StringPtrOutput)
}

func (o WorkflowOutput) DeploymentProperties() DeploymentPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *Workflow) DeploymentPropertiesResponsePtrOutput { return v.DeploymentProperties }).(DeploymentPropertiesResponsePtrOutput)
}

func (o WorkflowOutput) DockerBuildContext() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringPtrOutput { return v.DockerBuildContext }).(pulumi.StringPtrOutput)
}

func (o WorkflowOutput) Dockerfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringPtrOutput { return v.Dockerfile }).(pulumi.StringPtrOutput)
}

func (o WorkflowOutput) LastWorkflowRun() WorkflowRunResponsePtrOutput {
	return o.ApplyT(func(v *Workflow) WorkflowRunResponsePtrOutput { return v.LastWorkflowRun }).(WorkflowRunResponsePtrOutput)
}

func (o WorkflowOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o WorkflowOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WorkflowOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

func (o WorkflowOutput) OidcCredentials() GitHubWorkflowProfileResponseOidcCredentialsPtrOutput {
	return o.ApplyT(func(v *Workflow) GitHubWorkflowProfileResponseOidcCredentialsPtrOutput { return v.OidcCredentials }).(GitHubWorkflowProfileResponseOidcCredentialsPtrOutput)
}

func (o WorkflowOutput) PrStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.PrStatus }).(pulumi.StringOutput)
}

func (o WorkflowOutput) PrURL() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.PrURL }).(pulumi.StringOutput)
}

func (o WorkflowOutput) PullNumber() pulumi.IntOutput {
	return o.ApplyT(func(v *Workflow) pulumi.IntOutput { return v.PullNumber }).(pulumi.IntOutput)
}

func (o WorkflowOutput) RepositoryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringPtrOutput { return v.RepositoryName }).(pulumi.StringPtrOutput)
}

func (o WorkflowOutput) RepositoryOwner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringPtrOutput { return v.RepositoryOwner }).(pulumi.StringPtrOutput)
}

func (o WorkflowOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Workflow) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o WorkflowOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o WorkflowOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkflowOutput{})
}
