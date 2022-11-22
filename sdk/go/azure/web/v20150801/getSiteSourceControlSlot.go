


package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupSiteSourceControlSlot(ctx *pulumi.Context, args *LookupSiteSourceControlSlotArgs, opts ...pulumi.InvokeOption) (*LookupSiteSourceControlSlotResult, error) {
	var rv LookupSiteSourceControlSlotResult
	err := ctx.Invoke("azure-native:web/v20150801:getSiteSourceControlSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSiteSourceControlSlotArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type LookupSiteSourceControlSlotResult struct {
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

func LookupSiteSourceControlSlotOutput(ctx *pulumi.Context, args LookupSiteSourceControlSlotOutputArgs, opts ...pulumi.InvokeOption) LookupSiteSourceControlSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSiteSourceControlSlotResult, error) {
			args := v.(LookupSiteSourceControlSlotArgs)
			r, err := LookupSiteSourceControlSlot(ctx, &args, opts...)
			var s LookupSiteSourceControlSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSiteSourceControlSlotResultOutput)
}

type LookupSiteSourceControlSlotOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Slot              pulumi.StringInput `pulumi:"slot"`
}

func (LookupSiteSourceControlSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSiteSourceControlSlotArgs)(nil)).Elem()
}


type LookupSiteSourceControlSlotResultOutput struct{ *pulumi.OutputState }

func (LookupSiteSourceControlSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSiteSourceControlSlotResult)(nil)).Elem()
}

func (o LookupSiteSourceControlSlotResultOutput) ToLookupSiteSourceControlSlotResultOutput() LookupSiteSourceControlSlotResultOutput {
	return o
}

func (o LookupSiteSourceControlSlotResultOutput) ToLookupSiteSourceControlSlotResultOutputWithContext(ctx context.Context) LookupSiteSourceControlSlotResultOutput {
	return o
}

func (o LookupSiteSourceControlSlotResultOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteSourceControlSlotResult) *string { return v.Branch }).(pulumi.StringPtrOutput)
}

func (o LookupSiteSourceControlSlotResultOutput) DeploymentRollbackEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSiteSourceControlSlotResult) *bool { return v.DeploymentRollbackEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupSiteSourceControlSlotResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteSourceControlSlotResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupSiteSourceControlSlotResultOutput) IsManualIntegration() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSiteSourceControlSlotResult) *bool { return v.IsManualIntegration }).(pulumi.BoolPtrOutput)
}

func (o LookupSiteSourceControlSlotResultOutput) IsMercurial() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSiteSourceControlSlotResult) *bool { return v.IsMercurial }).(pulumi.BoolPtrOutput)
}

func (o LookupSiteSourceControlSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteSourceControlSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupSiteSourceControlSlotResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteSourceControlSlotResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSiteSourceControlSlotResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteSourceControlSlotResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupSiteSourceControlSlotResultOutput) RepoUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteSourceControlSlotResult) *string { return v.RepoUrl }).(pulumi.StringPtrOutput)
}

func (o LookupSiteSourceControlSlotResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSiteSourceControlSlotResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSiteSourceControlSlotResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteSourceControlSlotResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSiteSourceControlSlotResultOutput{})
}
