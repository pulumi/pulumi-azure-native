


package v20210101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppSourceControlSlot(ctx *pulumi.Context, args *LookupWebAppSourceControlSlotArgs, opts ...pulumi.InvokeOption) (*LookupWebAppSourceControlSlotResult, error) {
	var rv LookupWebAppSourceControlSlotResult
	err := ctx.Invoke("azure-native:web/v20210101:getWebAppSourceControlSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppSourceControlSlotArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type LookupWebAppSourceControlSlotResult struct {
	Branch                    *string                            `pulumi:"branch"`
	DeploymentRollbackEnabled *bool                              `pulumi:"deploymentRollbackEnabled"`
	GitHubActionConfiguration *GitHubActionConfigurationResponse `pulumi:"gitHubActionConfiguration"`
	Id                        string                             `pulumi:"id"`
	IsGitHubAction            *bool                              `pulumi:"isGitHubAction"`
	IsManualIntegration       *bool                              `pulumi:"isManualIntegration"`
	IsMercurial               *bool                              `pulumi:"isMercurial"`
	Kind                      *string                            `pulumi:"kind"`
	Name                      string                             `pulumi:"name"`
	RepoUrl                   *string                            `pulumi:"repoUrl"`
	Type                      string                             `pulumi:"type"`
}

func LookupWebAppSourceControlSlotOutput(ctx *pulumi.Context, args LookupWebAppSourceControlSlotOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppSourceControlSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppSourceControlSlotResult, error) {
			args := v.(LookupWebAppSourceControlSlotArgs)
			r, err := LookupWebAppSourceControlSlot(ctx, &args, opts...)
			var s LookupWebAppSourceControlSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppSourceControlSlotResultOutput)
}

type LookupWebAppSourceControlSlotOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Slot              pulumi.StringInput `pulumi:"slot"`
}

func (LookupWebAppSourceControlSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppSourceControlSlotArgs)(nil)).Elem()
}


type LookupWebAppSourceControlSlotResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppSourceControlSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppSourceControlSlotResult)(nil)).Elem()
}

func (o LookupWebAppSourceControlSlotResultOutput) ToLookupWebAppSourceControlSlotResultOutput() LookupWebAppSourceControlSlotResultOutput {
	return o
}

func (o LookupWebAppSourceControlSlotResultOutput) ToLookupWebAppSourceControlSlotResultOutputWithContext(ctx context.Context) LookupWebAppSourceControlSlotResultOutput {
	return o
}

func (o LookupWebAppSourceControlSlotResultOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSourceControlSlotResult) *string { return v.Branch }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSourceControlSlotResultOutput) DeploymentRollbackEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppSourceControlSlotResult) *bool { return v.DeploymentRollbackEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupWebAppSourceControlSlotResultOutput) GitHubActionConfiguration() GitHubActionConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupWebAppSourceControlSlotResult) *GitHubActionConfigurationResponse {
		return v.GitHubActionConfiguration
	}).(GitHubActionConfigurationResponsePtrOutput)
}

func (o LookupWebAppSourceControlSlotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSourceControlSlotResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWebAppSourceControlSlotResultOutput) IsGitHubAction() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppSourceControlSlotResult) *bool { return v.IsGitHubAction }).(pulumi.BoolPtrOutput)
}

func (o LookupWebAppSourceControlSlotResultOutput) IsManualIntegration() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppSourceControlSlotResult) *bool { return v.IsManualIntegration }).(pulumi.BoolPtrOutput)
}

func (o LookupWebAppSourceControlSlotResultOutput) IsMercurial() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppSourceControlSlotResult) *bool { return v.IsMercurial }).(pulumi.BoolPtrOutput)
}

func (o LookupWebAppSourceControlSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSourceControlSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSourceControlSlotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSourceControlSlotResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWebAppSourceControlSlotResultOutput) RepoUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSourceControlSlotResult) *string { return v.RepoUrl }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSourceControlSlotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSourceControlSlotResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppSourceControlSlotResultOutput{})
}
