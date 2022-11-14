


package v20181101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppSourceControl(ctx *pulumi.Context, args *LookupWebAppSourceControlArgs, opts ...pulumi.InvokeOption) (*LookupWebAppSourceControlResult, error) {
	var rv LookupWebAppSourceControlResult
	err := ctx.Invoke("azure-native:web/v20181101:getWebAppSourceControl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppSourceControlArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupWebAppSourceControlResult struct {
	Branch                    *string `pulumi:"branch"`
	DeploymentRollbackEnabled *bool   `pulumi:"deploymentRollbackEnabled"`
	Id                        string  `pulumi:"id"`
	IsManualIntegration       *bool   `pulumi:"isManualIntegration"`
	IsMercurial               *bool   `pulumi:"isMercurial"`
	Kind                      *string `pulumi:"kind"`
	Name                      string  `pulumi:"name"`
	RepoUrl                   *string `pulumi:"repoUrl"`
	Type                      string  `pulumi:"type"`
}

func LookupWebAppSourceControlOutput(ctx *pulumi.Context, args LookupWebAppSourceControlOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppSourceControlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppSourceControlResult, error) {
			args := v.(LookupWebAppSourceControlArgs)
			r, err := LookupWebAppSourceControl(ctx, &args, opts...)
			var s LookupWebAppSourceControlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppSourceControlResultOutput)
}

type LookupWebAppSourceControlOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupWebAppSourceControlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppSourceControlArgs)(nil)).Elem()
}


type LookupWebAppSourceControlResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppSourceControlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppSourceControlResult)(nil)).Elem()
}

func (o LookupWebAppSourceControlResultOutput) ToLookupWebAppSourceControlResultOutput() LookupWebAppSourceControlResultOutput {
	return o
}

func (o LookupWebAppSourceControlResultOutput) ToLookupWebAppSourceControlResultOutputWithContext(ctx context.Context) LookupWebAppSourceControlResultOutput {
	return o
}

func (o LookupWebAppSourceControlResultOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSourceControlResult) *string { return v.Branch }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSourceControlResultOutput) DeploymentRollbackEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppSourceControlResult) *bool { return v.DeploymentRollbackEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupWebAppSourceControlResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSourceControlResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWebAppSourceControlResultOutput) IsManualIntegration() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppSourceControlResult) *bool { return v.IsManualIntegration }).(pulumi.BoolPtrOutput)
}

func (o LookupWebAppSourceControlResultOutput) IsMercurial() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppSourceControlResult) *bool { return v.IsMercurial }).(pulumi.BoolPtrOutput)
}

func (o LookupWebAppSourceControlResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSourceControlResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSourceControlResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSourceControlResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWebAppSourceControlResultOutput) RepoUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSourceControlResult) *string { return v.RepoUrl }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSourceControlResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSourceControlResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppSourceControlResultOutput{})
}
