


package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupSiteSourceControl(ctx *pulumi.Context, args *LookupSiteSourceControlArgs, opts ...pulumi.InvokeOption) (*LookupSiteSourceControlResult, error) {
	var rv LookupSiteSourceControlResult
	err := ctx.Invoke("azure-native:web/v20150801:getSiteSourceControl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSiteSourceControlArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupSiteSourceControlResult struct {
	Branch                    *string           `pulumi:"branch"`
	DeploymentRollbackEnabled *bool             `pulumi:"deploymentRollbackEnabled"`
	Id                        *string           `pulumi:"id"`
	IsManualIntegration       *bool             `pulumi:"isManualIntegration"`
	IsMercurial               *bool             `pulumi:"isMercurial"`
	Kind                      *string           `pulumi:"kind"`
	Location                  string            `pulumi:"location"`
	Name                      *string           `pulumi:"name"`
	RepoUrl                   *string           `pulumi:"repoUrl"`
	Tags                      map[string]string `pulumi:"tags"`
	Type                      *string           `pulumi:"type"`
}

func LookupSiteSourceControlOutput(ctx *pulumi.Context, args LookupSiteSourceControlOutputArgs, opts ...pulumi.InvokeOption) LookupSiteSourceControlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSiteSourceControlResult, error) {
			args := v.(LookupSiteSourceControlArgs)
			r, err := LookupSiteSourceControl(ctx, &args, opts...)
			var s LookupSiteSourceControlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSiteSourceControlResultOutput)
}

type LookupSiteSourceControlOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupSiteSourceControlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSiteSourceControlArgs)(nil)).Elem()
}


type LookupSiteSourceControlResultOutput struct{ *pulumi.OutputState }

func (LookupSiteSourceControlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSiteSourceControlResult)(nil)).Elem()
}

func (o LookupSiteSourceControlResultOutput) ToLookupSiteSourceControlResultOutput() LookupSiteSourceControlResultOutput {
	return o
}

func (o LookupSiteSourceControlResultOutput) ToLookupSiteSourceControlResultOutputWithContext(ctx context.Context) LookupSiteSourceControlResultOutput {
	return o
}

func (o LookupSiteSourceControlResultOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteSourceControlResult) *string { return v.Branch }).(pulumi.StringPtrOutput)
}

func (o LookupSiteSourceControlResultOutput) DeploymentRollbackEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSiteSourceControlResult) *bool { return v.DeploymentRollbackEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupSiteSourceControlResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteSourceControlResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupSiteSourceControlResultOutput) IsManualIntegration() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSiteSourceControlResult) *bool { return v.IsManualIntegration }).(pulumi.BoolPtrOutput)
}

func (o LookupSiteSourceControlResultOutput) IsMercurial() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSiteSourceControlResult) *bool { return v.IsMercurial }).(pulumi.BoolPtrOutput)
}

func (o LookupSiteSourceControlResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteSourceControlResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupSiteSourceControlResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteSourceControlResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSiteSourceControlResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteSourceControlResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupSiteSourceControlResultOutput) RepoUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteSourceControlResult) *string { return v.RepoUrl }).(pulumi.StringPtrOutput)
}

func (o LookupSiteSourceControlResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSiteSourceControlResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSiteSourceControlResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteSourceControlResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSiteSourceControlResultOutput{})
}
