


package v20150521preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ArtifactDeploymentStatusProperties struct {
	ArtifactsApplied *int    `pulumi:"artifactsApplied"`
	DeploymentStatus *string `pulumi:"deploymentStatus"`
	TotalArtifacts   *int    `pulumi:"totalArtifacts"`
}





type ArtifactDeploymentStatusPropertiesInput interface {
	pulumi.Input

	ToArtifactDeploymentStatusPropertiesOutput() ArtifactDeploymentStatusPropertiesOutput
	ToArtifactDeploymentStatusPropertiesOutputWithContext(context.Context) ArtifactDeploymentStatusPropertiesOutput
}

type ArtifactDeploymentStatusPropertiesArgs struct {
	ArtifactsApplied pulumi.IntPtrInput    `pulumi:"artifactsApplied"`
	DeploymentStatus pulumi.StringPtrInput `pulumi:"deploymentStatus"`
	TotalArtifacts   pulumi.IntPtrInput    `pulumi:"totalArtifacts"`
}

func (ArtifactDeploymentStatusPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ArtifactDeploymentStatusProperties)(nil)).Elem()
}

func (i ArtifactDeploymentStatusPropertiesArgs) ToArtifactDeploymentStatusPropertiesOutput() ArtifactDeploymentStatusPropertiesOutput {
	return i.ToArtifactDeploymentStatusPropertiesOutputWithContext(context.Background())
}

func (i ArtifactDeploymentStatusPropertiesArgs) ToArtifactDeploymentStatusPropertiesOutputWithContext(ctx context.Context) ArtifactDeploymentStatusPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactDeploymentStatusPropertiesOutput)
}

func (i ArtifactDeploymentStatusPropertiesArgs) ToArtifactDeploymentStatusPropertiesPtrOutput() ArtifactDeploymentStatusPropertiesPtrOutput {
	return i.ToArtifactDeploymentStatusPropertiesPtrOutputWithContext(context.Background())
}

func (i ArtifactDeploymentStatusPropertiesArgs) ToArtifactDeploymentStatusPropertiesPtrOutputWithContext(ctx context.Context) ArtifactDeploymentStatusPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactDeploymentStatusPropertiesOutput).ToArtifactDeploymentStatusPropertiesPtrOutputWithContext(ctx)
}









type ArtifactDeploymentStatusPropertiesPtrInput interface {
	pulumi.Input

	ToArtifactDeploymentStatusPropertiesPtrOutput() ArtifactDeploymentStatusPropertiesPtrOutput
	ToArtifactDeploymentStatusPropertiesPtrOutputWithContext(context.Context) ArtifactDeploymentStatusPropertiesPtrOutput
}

type artifactDeploymentStatusPropertiesPtrType ArtifactDeploymentStatusPropertiesArgs

func ArtifactDeploymentStatusPropertiesPtr(v *ArtifactDeploymentStatusPropertiesArgs) ArtifactDeploymentStatusPropertiesPtrInput {
	return (*artifactDeploymentStatusPropertiesPtrType)(v)
}

func (*artifactDeploymentStatusPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ArtifactDeploymentStatusProperties)(nil)).Elem()
}

func (i *artifactDeploymentStatusPropertiesPtrType) ToArtifactDeploymentStatusPropertiesPtrOutput() ArtifactDeploymentStatusPropertiesPtrOutput {
	return i.ToArtifactDeploymentStatusPropertiesPtrOutputWithContext(context.Background())
}

func (i *artifactDeploymentStatusPropertiesPtrType) ToArtifactDeploymentStatusPropertiesPtrOutputWithContext(ctx context.Context) ArtifactDeploymentStatusPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactDeploymentStatusPropertiesPtrOutput)
}

type ArtifactDeploymentStatusPropertiesOutput struct{ *pulumi.OutputState }

func (ArtifactDeploymentStatusPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArtifactDeploymentStatusProperties)(nil)).Elem()
}

func (o ArtifactDeploymentStatusPropertiesOutput) ToArtifactDeploymentStatusPropertiesOutput() ArtifactDeploymentStatusPropertiesOutput {
	return o
}

func (o ArtifactDeploymentStatusPropertiesOutput) ToArtifactDeploymentStatusPropertiesOutputWithContext(ctx context.Context) ArtifactDeploymentStatusPropertiesOutput {
	return o
}

func (o ArtifactDeploymentStatusPropertiesOutput) ToArtifactDeploymentStatusPropertiesPtrOutput() ArtifactDeploymentStatusPropertiesPtrOutput {
	return o.ToArtifactDeploymentStatusPropertiesPtrOutputWithContext(context.Background())
}

func (o ArtifactDeploymentStatusPropertiesOutput) ToArtifactDeploymentStatusPropertiesPtrOutputWithContext(ctx context.Context) ArtifactDeploymentStatusPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ArtifactDeploymentStatusProperties) *ArtifactDeploymentStatusProperties {
		return &v
	}).(ArtifactDeploymentStatusPropertiesPtrOutput)
}

func (o ArtifactDeploymentStatusPropertiesOutput) ArtifactsApplied() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ArtifactDeploymentStatusProperties) *int { return v.ArtifactsApplied }).(pulumi.IntPtrOutput)
}

func (o ArtifactDeploymentStatusPropertiesOutput) DeploymentStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArtifactDeploymentStatusProperties) *string { return v.DeploymentStatus }).(pulumi.StringPtrOutput)
}

func (o ArtifactDeploymentStatusPropertiesOutput) TotalArtifacts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ArtifactDeploymentStatusProperties) *int { return v.TotalArtifacts }).(pulumi.IntPtrOutput)
}

type ArtifactDeploymentStatusPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ArtifactDeploymentStatusPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArtifactDeploymentStatusProperties)(nil)).Elem()
}

func (o ArtifactDeploymentStatusPropertiesPtrOutput) ToArtifactDeploymentStatusPropertiesPtrOutput() ArtifactDeploymentStatusPropertiesPtrOutput {
	return o
}

func (o ArtifactDeploymentStatusPropertiesPtrOutput) ToArtifactDeploymentStatusPropertiesPtrOutputWithContext(ctx context.Context) ArtifactDeploymentStatusPropertiesPtrOutput {
	return o
}

func (o ArtifactDeploymentStatusPropertiesPtrOutput) Elem() ArtifactDeploymentStatusPropertiesOutput {
	return o.ApplyT(func(v *ArtifactDeploymentStatusProperties) ArtifactDeploymentStatusProperties {
		if v != nil {
			return *v
		}
		var ret ArtifactDeploymentStatusProperties
		return ret
	}).(ArtifactDeploymentStatusPropertiesOutput)
}

func (o ArtifactDeploymentStatusPropertiesPtrOutput) ArtifactsApplied() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ArtifactDeploymentStatusProperties) *int {
		if v == nil {
			return nil
		}
		return v.ArtifactsApplied
	}).(pulumi.IntPtrOutput)
}

func (o ArtifactDeploymentStatusPropertiesPtrOutput) DeploymentStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArtifactDeploymentStatusProperties) *string {
		if v == nil {
			return nil
		}
		return v.DeploymentStatus
	}).(pulumi.StringPtrOutput)
}

func (o ArtifactDeploymentStatusPropertiesPtrOutput) TotalArtifacts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ArtifactDeploymentStatusProperties) *int {
		if v == nil {
			return nil
		}
		return v.TotalArtifacts
	}).(pulumi.IntPtrOutput)
}

type ArtifactDeploymentStatusPropertiesResponse struct {
	ArtifactsApplied *int    `pulumi:"artifactsApplied"`
	DeploymentStatus *string `pulumi:"deploymentStatus"`
	TotalArtifacts   *int    `pulumi:"totalArtifacts"`
}





type ArtifactDeploymentStatusPropertiesResponseInput interface {
	pulumi.Input

	ToArtifactDeploymentStatusPropertiesResponseOutput() ArtifactDeploymentStatusPropertiesResponseOutput
	ToArtifactDeploymentStatusPropertiesResponseOutputWithContext(context.Context) ArtifactDeploymentStatusPropertiesResponseOutput
}

type ArtifactDeploymentStatusPropertiesResponseArgs struct {
	ArtifactsApplied pulumi.IntPtrInput    `pulumi:"artifactsApplied"`
	DeploymentStatus pulumi.StringPtrInput `pulumi:"deploymentStatus"`
	TotalArtifacts   pulumi.IntPtrInput    `pulumi:"totalArtifacts"`
}

func (ArtifactDeploymentStatusPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ArtifactDeploymentStatusPropertiesResponse)(nil)).Elem()
}

func (i ArtifactDeploymentStatusPropertiesResponseArgs) ToArtifactDeploymentStatusPropertiesResponseOutput() ArtifactDeploymentStatusPropertiesResponseOutput {
	return i.ToArtifactDeploymentStatusPropertiesResponseOutputWithContext(context.Background())
}

func (i ArtifactDeploymentStatusPropertiesResponseArgs) ToArtifactDeploymentStatusPropertiesResponseOutputWithContext(ctx context.Context) ArtifactDeploymentStatusPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactDeploymentStatusPropertiesResponseOutput)
}

func (i ArtifactDeploymentStatusPropertiesResponseArgs) ToArtifactDeploymentStatusPropertiesResponsePtrOutput() ArtifactDeploymentStatusPropertiesResponsePtrOutput {
	return i.ToArtifactDeploymentStatusPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ArtifactDeploymentStatusPropertiesResponseArgs) ToArtifactDeploymentStatusPropertiesResponsePtrOutputWithContext(ctx context.Context) ArtifactDeploymentStatusPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactDeploymentStatusPropertiesResponseOutput).ToArtifactDeploymentStatusPropertiesResponsePtrOutputWithContext(ctx)
}









type ArtifactDeploymentStatusPropertiesResponsePtrInput interface {
	pulumi.Input

	ToArtifactDeploymentStatusPropertiesResponsePtrOutput() ArtifactDeploymentStatusPropertiesResponsePtrOutput
	ToArtifactDeploymentStatusPropertiesResponsePtrOutputWithContext(context.Context) ArtifactDeploymentStatusPropertiesResponsePtrOutput
}

type artifactDeploymentStatusPropertiesResponsePtrType ArtifactDeploymentStatusPropertiesResponseArgs

func ArtifactDeploymentStatusPropertiesResponsePtr(v *ArtifactDeploymentStatusPropertiesResponseArgs) ArtifactDeploymentStatusPropertiesResponsePtrInput {
	return (*artifactDeploymentStatusPropertiesResponsePtrType)(v)
}

func (*artifactDeploymentStatusPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ArtifactDeploymentStatusPropertiesResponse)(nil)).Elem()
}

func (i *artifactDeploymentStatusPropertiesResponsePtrType) ToArtifactDeploymentStatusPropertiesResponsePtrOutput() ArtifactDeploymentStatusPropertiesResponsePtrOutput {
	return i.ToArtifactDeploymentStatusPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *artifactDeploymentStatusPropertiesResponsePtrType) ToArtifactDeploymentStatusPropertiesResponsePtrOutputWithContext(ctx context.Context) ArtifactDeploymentStatusPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactDeploymentStatusPropertiesResponsePtrOutput)
}

type ArtifactDeploymentStatusPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ArtifactDeploymentStatusPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArtifactDeploymentStatusPropertiesResponse)(nil)).Elem()
}

func (o ArtifactDeploymentStatusPropertiesResponseOutput) ToArtifactDeploymentStatusPropertiesResponseOutput() ArtifactDeploymentStatusPropertiesResponseOutput {
	return o
}

func (o ArtifactDeploymentStatusPropertiesResponseOutput) ToArtifactDeploymentStatusPropertiesResponseOutputWithContext(ctx context.Context) ArtifactDeploymentStatusPropertiesResponseOutput {
	return o
}

func (o ArtifactDeploymentStatusPropertiesResponseOutput) ToArtifactDeploymentStatusPropertiesResponsePtrOutput() ArtifactDeploymentStatusPropertiesResponsePtrOutput {
	return o.ToArtifactDeploymentStatusPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ArtifactDeploymentStatusPropertiesResponseOutput) ToArtifactDeploymentStatusPropertiesResponsePtrOutputWithContext(ctx context.Context) ArtifactDeploymentStatusPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ArtifactDeploymentStatusPropertiesResponse) *ArtifactDeploymentStatusPropertiesResponse {
		return &v
	}).(ArtifactDeploymentStatusPropertiesResponsePtrOutput)
}

func (o ArtifactDeploymentStatusPropertiesResponseOutput) ArtifactsApplied() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ArtifactDeploymentStatusPropertiesResponse) *int { return v.ArtifactsApplied }).(pulumi.IntPtrOutput)
}

func (o ArtifactDeploymentStatusPropertiesResponseOutput) DeploymentStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArtifactDeploymentStatusPropertiesResponse) *string { return v.DeploymentStatus }).(pulumi.StringPtrOutput)
}

func (o ArtifactDeploymentStatusPropertiesResponseOutput) TotalArtifacts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ArtifactDeploymentStatusPropertiesResponse) *int { return v.TotalArtifacts }).(pulumi.IntPtrOutput)
}

type ArtifactDeploymentStatusPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ArtifactDeploymentStatusPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArtifactDeploymentStatusPropertiesResponse)(nil)).Elem()
}

func (o ArtifactDeploymentStatusPropertiesResponsePtrOutput) ToArtifactDeploymentStatusPropertiesResponsePtrOutput() ArtifactDeploymentStatusPropertiesResponsePtrOutput {
	return o
}

func (o ArtifactDeploymentStatusPropertiesResponsePtrOutput) ToArtifactDeploymentStatusPropertiesResponsePtrOutputWithContext(ctx context.Context) ArtifactDeploymentStatusPropertiesResponsePtrOutput {
	return o
}

func (o ArtifactDeploymentStatusPropertiesResponsePtrOutput) Elem() ArtifactDeploymentStatusPropertiesResponseOutput {
	return o.ApplyT(func(v *ArtifactDeploymentStatusPropertiesResponse) ArtifactDeploymentStatusPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ArtifactDeploymentStatusPropertiesResponse
		return ret
	}).(ArtifactDeploymentStatusPropertiesResponseOutput)
}

func (o ArtifactDeploymentStatusPropertiesResponsePtrOutput) ArtifactsApplied() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ArtifactDeploymentStatusPropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return v.ArtifactsApplied
	}).(pulumi.IntPtrOutput)
}

func (o ArtifactDeploymentStatusPropertiesResponsePtrOutput) DeploymentStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArtifactDeploymentStatusPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.DeploymentStatus
	}).(pulumi.StringPtrOutput)
}

func (o ArtifactDeploymentStatusPropertiesResponsePtrOutput) TotalArtifacts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ArtifactDeploymentStatusPropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return v.TotalArtifacts
	}).(pulumi.IntPtrOutput)
}

type ArtifactInstallProperties struct {
	ArtifactId *string                       `pulumi:"artifactId"`
	Parameters []ArtifactParameterProperties `pulumi:"parameters"`
}





type ArtifactInstallPropertiesInput interface {
	pulumi.Input

	ToArtifactInstallPropertiesOutput() ArtifactInstallPropertiesOutput
	ToArtifactInstallPropertiesOutputWithContext(context.Context) ArtifactInstallPropertiesOutput
}

type ArtifactInstallPropertiesArgs struct {
	ArtifactId pulumi.StringPtrInput                 `pulumi:"artifactId"`
	Parameters ArtifactParameterPropertiesArrayInput `pulumi:"parameters"`
}

func (ArtifactInstallPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ArtifactInstallProperties)(nil)).Elem()
}

func (i ArtifactInstallPropertiesArgs) ToArtifactInstallPropertiesOutput() ArtifactInstallPropertiesOutput {
	return i.ToArtifactInstallPropertiesOutputWithContext(context.Background())
}

func (i ArtifactInstallPropertiesArgs) ToArtifactInstallPropertiesOutputWithContext(ctx context.Context) ArtifactInstallPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactInstallPropertiesOutput)
}





type ArtifactInstallPropertiesArrayInput interface {
	pulumi.Input

	ToArtifactInstallPropertiesArrayOutput() ArtifactInstallPropertiesArrayOutput
	ToArtifactInstallPropertiesArrayOutputWithContext(context.Context) ArtifactInstallPropertiesArrayOutput
}

type ArtifactInstallPropertiesArray []ArtifactInstallPropertiesInput

func (ArtifactInstallPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ArtifactInstallProperties)(nil)).Elem()
}

func (i ArtifactInstallPropertiesArray) ToArtifactInstallPropertiesArrayOutput() ArtifactInstallPropertiesArrayOutput {
	return i.ToArtifactInstallPropertiesArrayOutputWithContext(context.Background())
}

func (i ArtifactInstallPropertiesArray) ToArtifactInstallPropertiesArrayOutputWithContext(ctx context.Context) ArtifactInstallPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactInstallPropertiesArrayOutput)
}

type ArtifactInstallPropertiesOutput struct{ *pulumi.OutputState }

func (ArtifactInstallPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArtifactInstallProperties)(nil)).Elem()
}

func (o ArtifactInstallPropertiesOutput) ToArtifactInstallPropertiesOutput() ArtifactInstallPropertiesOutput {
	return o
}

func (o ArtifactInstallPropertiesOutput) ToArtifactInstallPropertiesOutputWithContext(ctx context.Context) ArtifactInstallPropertiesOutput {
	return o
}

func (o ArtifactInstallPropertiesOutput) ArtifactId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArtifactInstallProperties) *string { return v.ArtifactId }).(pulumi.StringPtrOutput)
}

func (o ArtifactInstallPropertiesOutput) Parameters() ArtifactParameterPropertiesArrayOutput {
	return o.ApplyT(func(v ArtifactInstallProperties) []ArtifactParameterProperties { return v.Parameters }).(ArtifactParameterPropertiesArrayOutput)
}

type ArtifactInstallPropertiesArrayOutput struct{ *pulumi.OutputState }

func (ArtifactInstallPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ArtifactInstallProperties)(nil)).Elem()
}

func (o ArtifactInstallPropertiesArrayOutput) ToArtifactInstallPropertiesArrayOutput() ArtifactInstallPropertiesArrayOutput {
	return o
}

func (o ArtifactInstallPropertiesArrayOutput) ToArtifactInstallPropertiesArrayOutputWithContext(ctx context.Context) ArtifactInstallPropertiesArrayOutput {
	return o
}

func (o ArtifactInstallPropertiesArrayOutput) Index(i pulumi.IntInput) ArtifactInstallPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ArtifactInstallProperties {
		return vs[0].([]ArtifactInstallProperties)[vs[1].(int)]
	}).(ArtifactInstallPropertiesOutput)
}

type ArtifactInstallPropertiesResponse struct {
	ArtifactId *string                               `pulumi:"artifactId"`
	Parameters []ArtifactParameterPropertiesResponse `pulumi:"parameters"`
}





type ArtifactInstallPropertiesResponseInput interface {
	pulumi.Input

	ToArtifactInstallPropertiesResponseOutput() ArtifactInstallPropertiesResponseOutput
	ToArtifactInstallPropertiesResponseOutputWithContext(context.Context) ArtifactInstallPropertiesResponseOutput
}

type ArtifactInstallPropertiesResponseArgs struct {
	ArtifactId pulumi.StringPtrInput                         `pulumi:"artifactId"`
	Parameters ArtifactParameterPropertiesResponseArrayInput `pulumi:"parameters"`
}

func (ArtifactInstallPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ArtifactInstallPropertiesResponse)(nil)).Elem()
}

func (i ArtifactInstallPropertiesResponseArgs) ToArtifactInstallPropertiesResponseOutput() ArtifactInstallPropertiesResponseOutput {
	return i.ToArtifactInstallPropertiesResponseOutputWithContext(context.Background())
}

func (i ArtifactInstallPropertiesResponseArgs) ToArtifactInstallPropertiesResponseOutputWithContext(ctx context.Context) ArtifactInstallPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactInstallPropertiesResponseOutput)
}





type ArtifactInstallPropertiesResponseArrayInput interface {
	pulumi.Input

	ToArtifactInstallPropertiesResponseArrayOutput() ArtifactInstallPropertiesResponseArrayOutput
	ToArtifactInstallPropertiesResponseArrayOutputWithContext(context.Context) ArtifactInstallPropertiesResponseArrayOutput
}

type ArtifactInstallPropertiesResponseArray []ArtifactInstallPropertiesResponseInput

func (ArtifactInstallPropertiesResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ArtifactInstallPropertiesResponse)(nil)).Elem()
}

func (i ArtifactInstallPropertiesResponseArray) ToArtifactInstallPropertiesResponseArrayOutput() ArtifactInstallPropertiesResponseArrayOutput {
	return i.ToArtifactInstallPropertiesResponseArrayOutputWithContext(context.Background())
}

func (i ArtifactInstallPropertiesResponseArray) ToArtifactInstallPropertiesResponseArrayOutputWithContext(ctx context.Context) ArtifactInstallPropertiesResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactInstallPropertiesResponseArrayOutput)
}

type ArtifactInstallPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ArtifactInstallPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArtifactInstallPropertiesResponse)(nil)).Elem()
}

func (o ArtifactInstallPropertiesResponseOutput) ToArtifactInstallPropertiesResponseOutput() ArtifactInstallPropertiesResponseOutput {
	return o
}

func (o ArtifactInstallPropertiesResponseOutput) ToArtifactInstallPropertiesResponseOutputWithContext(ctx context.Context) ArtifactInstallPropertiesResponseOutput {
	return o
}

func (o ArtifactInstallPropertiesResponseOutput) ArtifactId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArtifactInstallPropertiesResponse) *string { return v.ArtifactId }).(pulumi.StringPtrOutput)
}

func (o ArtifactInstallPropertiesResponseOutput) Parameters() ArtifactParameterPropertiesResponseArrayOutput {
	return o.ApplyT(func(v ArtifactInstallPropertiesResponse) []ArtifactParameterPropertiesResponse { return v.Parameters }).(ArtifactParameterPropertiesResponseArrayOutput)
}

type ArtifactInstallPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (ArtifactInstallPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ArtifactInstallPropertiesResponse)(nil)).Elem()
}

func (o ArtifactInstallPropertiesResponseArrayOutput) ToArtifactInstallPropertiesResponseArrayOutput() ArtifactInstallPropertiesResponseArrayOutput {
	return o
}

func (o ArtifactInstallPropertiesResponseArrayOutput) ToArtifactInstallPropertiesResponseArrayOutputWithContext(ctx context.Context) ArtifactInstallPropertiesResponseArrayOutput {
	return o
}

func (o ArtifactInstallPropertiesResponseArrayOutput) Index(i pulumi.IntInput) ArtifactInstallPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ArtifactInstallPropertiesResponse {
		return vs[0].([]ArtifactInstallPropertiesResponse)[vs[1].(int)]
	}).(ArtifactInstallPropertiesResponseOutput)
}

type ArtifactParameterProperties struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}





type ArtifactParameterPropertiesInput interface {
	pulumi.Input

	ToArtifactParameterPropertiesOutput() ArtifactParameterPropertiesOutput
	ToArtifactParameterPropertiesOutputWithContext(context.Context) ArtifactParameterPropertiesOutput
}

type ArtifactParameterPropertiesArgs struct {
	Name  pulumi.StringPtrInput `pulumi:"name"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (ArtifactParameterPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ArtifactParameterProperties)(nil)).Elem()
}

func (i ArtifactParameterPropertiesArgs) ToArtifactParameterPropertiesOutput() ArtifactParameterPropertiesOutput {
	return i.ToArtifactParameterPropertiesOutputWithContext(context.Background())
}

func (i ArtifactParameterPropertiesArgs) ToArtifactParameterPropertiesOutputWithContext(ctx context.Context) ArtifactParameterPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactParameterPropertiesOutput)
}





type ArtifactParameterPropertiesArrayInput interface {
	pulumi.Input

	ToArtifactParameterPropertiesArrayOutput() ArtifactParameterPropertiesArrayOutput
	ToArtifactParameterPropertiesArrayOutputWithContext(context.Context) ArtifactParameterPropertiesArrayOutput
}

type ArtifactParameterPropertiesArray []ArtifactParameterPropertiesInput

func (ArtifactParameterPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ArtifactParameterProperties)(nil)).Elem()
}

func (i ArtifactParameterPropertiesArray) ToArtifactParameterPropertiesArrayOutput() ArtifactParameterPropertiesArrayOutput {
	return i.ToArtifactParameterPropertiesArrayOutputWithContext(context.Background())
}

func (i ArtifactParameterPropertiesArray) ToArtifactParameterPropertiesArrayOutputWithContext(ctx context.Context) ArtifactParameterPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactParameterPropertiesArrayOutput)
}

type ArtifactParameterPropertiesOutput struct{ *pulumi.OutputState }

func (ArtifactParameterPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArtifactParameterProperties)(nil)).Elem()
}

func (o ArtifactParameterPropertiesOutput) ToArtifactParameterPropertiesOutput() ArtifactParameterPropertiesOutput {
	return o
}

func (o ArtifactParameterPropertiesOutput) ToArtifactParameterPropertiesOutputWithContext(ctx context.Context) ArtifactParameterPropertiesOutput {
	return o
}

func (o ArtifactParameterPropertiesOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArtifactParameterProperties) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ArtifactParameterPropertiesOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArtifactParameterProperties) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type ArtifactParameterPropertiesArrayOutput struct{ *pulumi.OutputState }

func (ArtifactParameterPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ArtifactParameterProperties)(nil)).Elem()
}

func (o ArtifactParameterPropertiesArrayOutput) ToArtifactParameterPropertiesArrayOutput() ArtifactParameterPropertiesArrayOutput {
	return o
}

func (o ArtifactParameterPropertiesArrayOutput) ToArtifactParameterPropertiesArrayOutputWithContext(ctx context.Context) ArtifactParameterPropertiesArrayOutput {
	return o
}

func (o ArtifactParameterPropertiesArrayOutput) Index(i pulumi.IntInput) ArtifactParameterPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ArtifactParameterProperties {
		return vs[0].([]ArtifactParameterProperties)[vs[1].(int)]
	}).(ArtifactParameterPropertiesOutput)
}

type ArtifactParameterPropertiesResponse struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}





type ArtifactParameterPropertiesResponseInput interface {
	pulumi.Input

	ToArtifactParameterPropertiesResponseOutput() ArtifactParameterPropertiesResponseOutput
	ToArtifactParameterPropertiesResponseOutputWithContext(context.Context) ArtifactParameterPropertiesResponseOutput
}

type ArtifactParameterPropertiesResponseArgs struct {
	Name  pulumi.StringPtrInput `pulumi:"name"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (ArtifactParameterPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ArtifactParameterPropertiesResponse)(nil)).Elem()
}

func (i ArtifactParameterPropertiesResponseArgs) ToArtifactParameterPropertiesResponseOutput() ArtifactParameterPropertiesResponseOutput {
	return i.ToArtifactParameterPropertiesResponseOutputWithContext(context.Background())
}

func (i ArtifactParameterPropertiesResponseArgs) ToArtifactParameterPropertiesResponseOutputWithContext(ctx context.Context) ArtifactParameterPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactParameterPropertiesResponseOutput)
}





type ArtifactParameterPropertiesResponseArrayInput interface {
	pulumi.Input

	ToArtifactParameterPropertiesResponseArrayOutput() ArtifactParameterPropertiesResponseArrayOutput
	ToArtifactParameterPropertiesResponseArrayOutputWithContext(context.Context) ArtifactParameterPropertiesResponseArrayOutput
}

type ArtifactParameterPropertiesResponseArray []ArtifactParameterPropertiesResponseInput

func (ArtifactParameterPropertiesResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ArtifactParameterPropertiesResponse)(nil)).Elem()
}

func (i ArtifactParameterPropertiesResponseArray) ToArtifactParameterPropertiesResponseArrayOutput() ArtifactParameterPropertiesResponseArrayOutput {
	return i.ToArtifactParameterPropertiesResponseArrayOutputWithContext(context.Background())
}

func (i ArtifactParameterPropertiesResponseArray) ToArtifactParameterPropertiesResponseArrayOutputWithContext(ctx context.Context) ArtifactParameterPropertiesResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactParameterPropertiesResponseArrayOutput)
}

type ArtifactParameterPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ArtifactParameterPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArtifactParameterPropertiesResponse)(nil)).Elem()
}

func (o ArtifactParameterPropertiesResponseOutput) ToArtifactParameterPropertiesResponseOutput() ArtifactParameterPropertiesResponseOutput {
	return o
}

func (o ArtifactParameterPropertiesResponseOutput) ToArtifactParameterPropertiesResponseOutputWithContext(ctx context.Context) ArtifactParameterPropertiesResponseOutput {
	return o
}

func (o ArtifactParameterPropertiesResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArtifactParameterPropertiesResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ArtifactParameterPropertiesResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArtifactParameterPropertiesResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type ArtifactParameterPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (ArtifactParameterPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ArtifactParameterPropertiesResponse)(nil)).Elem()
}

func (o ArtifactParameterPropertiesResponseArrayOutput) ToArtifactParameterPropertiesResponseArrayOutput() ArtifactParameterPropertiesResponseArrayOutput {
	return o
}

func (o ArtifactParameterPropertiesResponseArrayOutput) ToArtifactParameterPropertiesResponseArrayOutputWithContext(ctx context.Context) ArtifactParameterPropertiesResponseArrayOutput {
	return o
}

func (o ArtifactParameterPropertiesResponseArrayOutput) Index(i pulumi.IntInput) ArtifactParameterPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ArtifactParameterPropertiesResponse {
		return vs[0].([]ArtifactParameterPropertiesResponse)[vs[1].(int)]
	}).(ArtifactParameterPropertiesResponseOutput)
}

type CustomImagePropertiesCustom struct {
	ImageName *string `pulumi:"imageName"`
	SysPrep   *bool   `pulumi:"sysPrep"`
}





type CustomImagePropertiesCustomInput interface {
	pulumi.Input

	ToCustomImagePropertiesCustomOutput() CustomImagePropertiesCustomOutput
	ToCustomImagePropertiesCustomOutputWithContext(context.Context) CustomImagePropertiesCustomOutput
}

type CustomImagePropertiesCustomArgs struct {
	ImageName pulumi.StringPtrInput `pulumi:"imageName"`
	SysPrep   pulumi.BoolPtrInput   `pulumi:"sysPrep"`
}

func (CustomImagePropertiesCustomArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomImagePropertiesCustom)(nil)).Elem()
}

func (i CustomImagePropertiesCustomArgs) ToCustomImagePropertiesCustomOutput() CustomImagePropertiesCustomOutput {
	return i.ToCustomImagePropertiesCustomOutputWithContext(context.Background())
}

func (i CustomImagePropertiesCustomArgs) ToCustomImagePropertiesCustomOutputWithContext(ctx context.Context) CustomImagePropertiesCustomOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomImagePropertiesCustomOutput)
}

func (i CustomImagePropertiesCustomArgs) ToCustomImagePropertiesCustomPtrOutput() CustomImagePropertiesCustomPtrOutput {
	return i.ToCustomImagePropertiesCustomPtrOutputWithContext(context.Background())
}

func (i CustomImagePropertiesCustomArgs) ToCustomImagePropertiesCustomPtrOutputWithContext(ctx context.Context) CustomImagePropertiesCustomPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomImagePropertiesCustomOutput).ToCustomImagePropertiesCustomPtrOutputWithContext(ctx)
}









type CustomImagePropertiesCustomPtrInput interface {
	pulumi.Input

	ToCustomImagePropertiesCustomPtrOutput() CustomImagePropertiesCustomPtrOutput
	ToCustomImagePropertiesCustomPtrOutputWithContext(context.Context) CustomImagePropertiesCustomPtrOutput
}

type customImagePropertiesCustomPtrType CustomImagePropertiesCustomArgs

func CustomImagePropertiesCustomPtr(v *CustomImagePropertiesCustomArgs) CustomImagePropertiesCustomPtrInput {
	return (*customImagePropertiesCustomPtrType)(v)
}

func (*customImagePropertiesCustomPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomImagePropertiesCustom)(nil)).Elem()
}

func (i *customImagePropertiesCustomPtrType) ToCustomImagePropertiesCustomPtrOutput() CustomImagePropertiesCustomPtrOutput {
	return i.ToCustomImagePropertiesCustomPtrOutputWithContext(context.Background())
}

func (i *customImagePropertiesCustomPtrType) ToCustomImagePropertiesCustomPtrOutputWithContext(ctx context.Context) CustomImagePropertiesCustomPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomImagePropertiesCustomPtrOutput)
}

type CustomImagePropertiesCustomOutput struct{ *pulumi.OutputState }

func (CustomImagePropertiesCustomOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomImagePropertiesCustom)(nil)).Elem()
}

func (o CustomImagePropertiesCustomOutput) ToCustomImagePropertiesCustomOutput() CustomImagePropertiesCustomOutput {
	return o
}

func (o CustomImagePropertiesCustomOutput) ToCustomImagePropertiesCustomOutputWithContext(ctx context.Context) CustomImagePropertiesCustomOutput {
	return o
}

func (o CustomImagePropertiesCustomOutput) ToCustomImagePropertiesCustomPtrOutput() CustomImagePropertiesCustomPtrOutput {
	return o.ToCustomImagePropertiesCustomPtrOutputWithContext(context.Background())
}

func (o CustomImagePropertiesCustomOutput) ToCustomImagePropertiesCustomPtrOutputWithContext(ctx context.Context) CustomImagePropertiesCustomPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomImagePropertiesCustom) *CustomImagePropertiesCustom {
		return &v
	}).(CustomImagePropertiesCustomPtrOutput)
}

func (o CustomImagePropertiesCustomOutput) ImageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomImagePropertiesCustom) *string { return v.ImageName }).(pulumi.StringPtrOutput)
}

func (o CustomImagePropertiesCustomOutput) SysPrep() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CustomImagePropertiesCustom) *bool { return v.SysPrep }).(pulumi.BoolPtrOutput)
}

type CustomImagePropertiesCustomPtrOutput struct{ *pulumi.OutputState }

func (CustomImagePropertiesCustomPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomImagePropertiesCustom)(nil)).Elem()
}

func (o CustomImagePropertiesCustomPtrOutput) ToCustomImagePropertiesCustomPtrOutput() CustomImagePropertiesCustomPtrOutput {
	return o
}

func (o CustomImagePropertiesCustomPtrOutput) ToCustomImagePropertiesCustomPtrOutputWithContext(ctx context.Context) CustomImagePropertiesCustomPtrOutput {
	return o
}

func (o CustomImagePropertiesCustomPtrOutput) Elem() CustomImagePropertiesCustomOutput {
	return o.ApplyT(func(v *CustomImagePropertiesCustom) CustomImagePropertiesCustom {
		if v != nil {
			return *v
		}
		var ret CustomImagePropertiesCustom
		return ret
	}).(CustomImagePropertiesCustomOutput)
}

func (o CustomImagePropertiesCustomPtrOutput) ImageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomImagePropertiesCustom) *string {
		if v == nil {
			return nil
		}
		return v.ImageName
	}).(pulumi.StringPtrOutput)
}

func (o CustomImagePropertiesCustomPtrOutput) SysPrep() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CustomImagePropertiesCustom) *bool {
		if v == nil {
			return nil
		}
		return v.SysPrep
	}).(pulumi.BoolPtrOutput)
}

type CustomImagePropertiesCustomResponse struct {
	ImageName *string `pulumi:"imageName"`
	SysPrep   *bool   `pulumi:"sysPrep"`
}





type CustomImagePropertiesCustomResponseInput interface {
	pulumi.Input

	ToCustomImagePropertiesCustomResponseOutput() CustomImagePropertiesCustomResponseOutput
	ToCustomImagePropertiesCustomResponseOutputWithContext(context.Context) CustomImagePropertiesCustomResponseOutput
}

type CustomImagePropertiesCustomResponseArgs struct {
	ImageName pulumi.StringPtrInput `pulumi:"imageName"`
	SysPrep   pulumi.BoolPtrInput   `pulumi:"sysPrep"`
}

func (CustomImagePropertiesCustomResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomImagePropertiesCustomResponse)(nil)).Elem()
}

func (i CustomImagePropertiesCustomResponseArgs) ToCustomImagePropertiesCustomResponseOutput() CustomImagePropertiesCustomResponseOutput {
	return i.ToCustomImagePropertiesCustomResponseOutputWithContext(context.Background())
}

func (i CustomImagePropertiesCustomResponseArgs) ToCustomImagePropertiesCustomResponseOutputWithContext(ctx context.Context) CustomImagePropertiesCustomResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomImagePropertiesCustomResponseOutput)
}

func (i CustomImagePropertiesCustomResponseArgs) ToCustomImagePropertiesCustomResponsePtrOutput() CustomImagePropertiesCustomResponsePtrOutput {
	return i.ToCustomImagePropertiesCustomResponsePtrOutputWithContext(context.Background())
}

func (i CustomImagePropertiesCustomResponseArgs) ToCustomImagePropertiesCustomResponsePtrOutputWithContext(ctx context.Context) CustomImagePropertiesCustomResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomImagePropertiesCustomResponseOutput).ToCustomImagePropertiesCustomResponsePtrOutputWithContext(ctx)
}









type CustomImagePropertiesCustomResponsePtrInput interface {
	pulumi.Input

	ToCustomImagePropertiesCustomResponsePtrOutput() CustomImagePropertiesCustomResponsePtrOutput
	ToCustomImagePropertiesCustomResponsePtrOutputWithContext(context.Context) CustomImagePropertiesCustomResponsePtrOutput
}

type customImagePropertiesCustomResponsePtrType CustomImagePropertiesCustomResponseArgs

func CustomImagePropertiesCustomResponsePtr(v *CustomImagePropertiesCustomResponseArgs) CustomImagePropertiesCustomResponsePtrInput {
	return (*customImagePropertiesCustomResponsePtrType)(v)
}

func (*customImagePropertiesCustomResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomImagePropertiesCustomResponse)(nil)).Elem()
}

func (i *customImagePropertiesCustomResponsePtrType) ToCustomImagePropertiesCustomResponsePtrOutput() CustomImagePropertiesCustomResponsePtrOutput {
	return i.ToCustomImagePropertiesCustomResponsePtrOutputWithContext(context.Background())
}

func (i *customImagePropertiesCustomResponsePtrType) ToCustomImagePropertiesCustomResponsePtrOutputWithContext(ctx context.Context) CustomImagePropertiesCustomResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomImagePropertiesCustomResponsePtrOutput)
}

type CustomImagePropertiesCustomResponseOutput struct{ *pulumi.OutputState }

func (CustomImagePropertiesCustomResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomImagePropertiesCustomResponse)(nil)).Elem()
}

func (o CustomImagePropertiesCustomResponseOutput) ToCustomImagePropertiesCustomResponseOutput() CustomImagePropertiesCustomResponseOutput {
	return o
}

func (o CustomImagePropertiesCustomResponseOutput) ToCustomImagePropertiesCustomResponseOutputWithContext(ctx context.Context) CustomImagePropertiesCustomResponseOutput {
	return o
}

func (o CustomImagePropertiesCustomResponseOutput) ToCustomImagePropertiesCustomResponsePtrOutput() CustomImagePropertiesCustomResponsePtrOutput {
	return o.ToCustomImagePropertiesCustomResponsePtrOutputWithContext(context.Background())
}

func (o CustomImagePropertiesCustomResponseOutput) ToCustomImagePropertiesCustomResponsePtrOutputWithContext(ctx context.Context) CustomImagePropertiesCustomResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomImagePropertiesCustomResponse) *CustomImagePropertiesCustomResponse {
		return &v
	}).(CustomImagePropertiesCustomResponsePtrOutput)
}

func (o CustomImagePropertiesCustomResponseOutput) ImageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomImagePropertiesCustomResponse) *string { return v.ImageName }).(pulumi.StringPtrOutput)
}

func (o CustomImagePropertiesCustomResponseOutput) SysPrep() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CustomImagePropertiesCustomResponse) *bool { return v.SysPrep }).(pulumi.BoolPtrOutput)
}

type CustomImagePropertiesCustomResponsePtrOutput struct{ *pulumi.OutputState }

func (CustomImagePropertiesCustomResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomImagePropertiesCustomResponse)(nil)).Elem()
}

func (o CustomImagePropertiesCustomResponsePtrOutput) ToCustomImagePropertiesCustomResponsePtrOutput() CustomImagePropertiesCustomResponsePtrOutput {
	return o
}

func (o CustomImagePropertiesCustomResponsePtrOutput) ToCustomImagePropertiesCustomResponsePtrOutputWithContext(ctx context.Context) CustomImagePropertiesCustomResponsePtrOutput {
	return o
}

func (o CustomImagePropertiesCustomResponsePtrOutput) Elem() CustomImagePropertiesCustomResponseOutput {
	return o.ApplyT(func(v *CustomImagePropertiesCustomResponse) CustomImagePropertiesCustomResponse {
		if v != nil {
			return *v
		}
		var ret CustomImagePropertiesCustomResponse
		return ret
	}).(CustomImagePropertiesCustomResponseOutput)
}

func (o CustomImagePropertiesCustomResponsePtrOutput) ImageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomImagePropertiesCustomResponse) *string {
		if v == nil {
			return nil
		}
		return v.ImageName
	}).(pulumi.StringPtrOutput)
}

func (o CustomImagePropertiesCustomResponsePtrOutput) SysPrep() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CustomImagePropertiesCustomResponse) *bool {
		if v == nil {
			return nil
		}
		return v.SysPrep
	}).(pulumi.BoolPtrOutput)
}

type CustomImagePropertiesFromVm struct {
	LinuxOsInfo   *LinuxOsInfo   `pulumi:"linuxOsInfo"`
	SourceVmId    *string        `pulumi:"sourceVmId"`
	SysPrep       *bool          `pulumi:"sysPrep"`
	WindowsOsInfo *WindowsOsInfo `pulumi:"windowsOsInfo"`
}





type CustomImagePropertiesFromVmInput interface {
	pulumi.Input

	ToCustomImagePropertiesFromVmOutput() CustomImagePropertiesFromVmOutput
	ToCustomImagePropertiesFromVmOutputWithContext(context.Context) CustomImagePropertiesFromVmOutput
}

type CustomImagePropertiesFromVmArgs struct {
	LinuxOsInfo   LinuxOsInfoPtrInput   `pulumi:"linuxOsInfo"`
	SourceVmId    pulumi.StringPtrInput `pulumi:"sourceVmId"`
	SysPrep       pulumi.BoolPtrInput   `pulumi:"sysPrep"`
	WindowsOsInfo WindowsOsInfoPtrInput `pulumi:"windowsOsInfo"`
}

func (CustomImagePropertiesFromVmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomImagePropertiesFromVm)(nil)).Elem()
}

func (i CustomImagePropertiesFromVmArgs) ToCustomImagePropertiesFromVmOutput() CustomImagePropertiesFromVmOutput {
	return i.ToCustomImagePropertiesFromVmOutputWithContext(context.Background())
}

func (i CustomImagePropertiesFromVmArgs) ToCustomImagePropertiesFromVmOutputWithContext(ctx context.Context) CustomImagePropertiesFromVmOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomImagePropertiesFromVmOutput)
}

func (i CustomImagePropertiesFromVmArgs) ToCustomImagePropertiesFromVmPtrOutput() CustomImagePropertiesFromVmPtrOutput {
	return i.ToCustomImagePropertiesFromVmPtrOutputWithContext(context.Background())
}

func (i CustomImagePropertiesFromVmArgs) ToCustomImagePropertiesFromVmPtrOutputWithContext(ctx context.Context) CustomImagePropertiesFromVmPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomImagePropertiesFromVmOutput).ToCustomImagePropertiesFromVmPtrOutputWithContext(ctx)
}









type CustomImagePropertiesFromVmPtrInput interface {
	pulumi.Input

	ToCustomImagePropertiesFromVmPtrOutput() CustomImagePropertiesFromVmPtrOutput
	ToCustomImagePropertiesFromVmPtrOutputWithContext(context.Context) CustomImagePropertiesFromVmPtrOutput
}

type customImagePropertiesFromVmPtrType CustomImagePropertiesFromVmArgs

func CustomImagePropertiesFromVmPtr(v *CustomImagePropertiesFromVmArgs) CustomImagePropertiesFromVmPtrInput {
	return (*customImagePropertiesFromVmPtrType)(v)
}

func (*customImagePropertiesFromVmPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomImagePropertiesFromVm)(nil)).Elem()
}

func (i *customImagePropertiesFromVmPtrType) ToCustomImagePropertiesFromVmPtrOutput() CustomImagePropertiesFromVmPtrOutput {
	return i.ToCustomImagePropertiesFromVmPtrOutputWithContext(context.Background())
}

func (i *customImagePropertiesFromVmPtrType) ToCustomImagePropertiesFromVmPtrOutputWithContext(ctx context.Context) CustomImagePropertiesFromVmPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomImagePropertiesFromVmPtrOutput)
}

type CustomImagePropertiesFromVmOutput struct{ *pulumi.OutputState }

func (CustomImagePropertiesFromVmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomImagePropertiesFromVm)(nil)).Elem()
}

func (o CustomImagePropertiesFromVmOutput) ToCustomImagePropertiesFromVmOutput() CustomImagePropertiesFromVmOutput {
	return o
}

func (o CustomImagePropertiesFromVmOutput) ToCustomImagePropertiesFromVmOutputWithContext(ctx context.Context) CustomImagePropertiesFromVmOutput {
	return o
}

func (o CustomImagePropertiesFromVmOutput) ToCustomImagePropertiesFromVmPtrOutput() CustomImagePropertiesFromVmPtrOutput {
	return o.ToCustomImagePropertiesFromVmPtrOutputWithContext(context.Background())
}

func (o CustomImagePropertiesFromVmOutput) ToCustomImagePropertiesFromVmPtrOutputWithContext(ctx context.Context) CustomImagePropertiesFromVmPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomImagePropertiesFromVm) *CustomImagePropertiesFromVm {
		return &v
	}).(CustomImagePropertiesFromVmPtrOutput)
}

func (o CustomImagePropertiesFromVmOutput) LinuxOsInfo() LinuxOsInfoPtrOutput {
	return o.ApplyT(func(v CustomImagePropertiesFromVm) *LinuxOsInfo { return v.LinuxOsInfo }).(LinuxOsInfoPtrOutput)
}

func (o CustomImagePropertiesFromVmOutput) SourceVmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomImagePropertiesFromVm) *string { return v.SourceVmId }).(pulumi.StringPtrOutput)
}

func (o CustomImagePropertiesFromVmOutput) SysPrep() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CustomImagePropertiesFromVm) *bool { return v.SysPrep }).(pulumi.BoolPtrOutput)
}

func (o CustomImagePropertiesFromVmOutput) WindowsOsInfo() WindowsOsInfoPtrOutput {
	return o.ApplyT(func(v CustomImagePropertiesFromVm) *WindowsOsInfo { return v.WindowsOsInfo }).(WindowsOsInfoPtrOutput)
}

type CustomImagePropertiesFromVmPtrOutput struct{ *pulumi.OutputState }

func (CustomImagePropertiesFromVmPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomImagePropertiesFromVm)(nil)).Elem()
}

func (o CustomImagePropertiesFromVmPtrOutput) ToCustomImagePropertiesFromVmPtrOutput() CustomImagePropertiesFromVmPtrOutput {
	return o
}

func (o CustomImagePropertiesFromVmPtrOutput) ToCustomImagePropertiesFromVmPtrOutputWithContext(ctx context.Context) CustomImagePropertiesFromVmPtrOutput {
	return o
}

func (o CustomImagePropertiesFromVmPtrOutput) Elem() CustomImagePropertiesFromVmOutput {
	return o.ApplyT(func(v *CustomImagePropertiesFromVm) CustomImagePropertiesFromVm {
		if v != nil {
			return *v
		}
		var ret CustomImagePropertiesFromVm
		return ret
	}).(CustomImagePropertiesFromVmOutput)
}

func (o CustomImagePropertiesFromVmPtrOutput) LinuxOsInfo() LinuxOsInfoPtrOutput {
	return o.ApplyT(func(v *CustomImagePropertiesFromVm) *LinuxOsInfo {
		if v == nil {
			return nil
		}
		return v.LinuxOsInfo
	}).(LinuxOsInfoPtrOutput)
}

func (o CustomImagePropertiesFromVmPtrOutput) SourceVmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomImagePropertiesFromVm) *string {
		if v == nil {
			return nil
		}
		return v.SourceVmId
	}).(pulumi.StringPtrOutput)
}

func (o CustomImagePropertiesFromVmPtrOutput) SysPrep() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CustomImagePropertiesFromVm) *bool {
		if v == nil {
			return nil
		}
		return v.SysPrep
	}).(pulumi.BoolPtrOutput)
}

func (o CustomImagePropertiesFromVmPtrOutput) WindowsOsInfo() WindowsOsInfoPtrOutput {
	return o.ApplyT(func(v *CustomImagePropertiesFromVm) *WindowsOsInfo {
		if v == nil {
			return nil
		}
		return v.WindowsOsInfo
	}).(WindowsOsInfoPtrOutput)
}

type CustomImagePropertiesFromVmResponse struct {
	LinuxOsInfo   *LinuxOsInfoResponse   `pulumi:"linuxOsInfo"`
	SourceVmId    *string                `pulumi:"sourceVmId"`
	SysPrep       *bool                  `pulumi:"sysPrep"`
	WindowsOsInfo *WindowsOsInfoResponse `pulumi:"windowsOsInfo"`
}





type CustomImagePropertiesFromVmResponseInput interface {
	pulumi.Input

	ToCustomImagePropertiesFromVmResponseOutput() CustomImagePropertiesFromVmResponseOutput
	ToCustomImagePropertiesFromVmResponseOutputWithContext(context.Context) CustomImagePropertiesFromVmResponseOutput
}

type CustomImagePropertiesFromVmResponseArgs struct {
	LinuxOsInfo   LinuxOsInfoResponsePtrInput   `pulumi:"linuxOsInfo"`
	SourceVmId    pulumi.StringPtrInput         `pulumi:"sourceVmId"`
	SysPrep       pulumi.BoolPtrInput           `pulumi:"sysPrep"`
	WindowsOsInfo WindowsOsInfoResponsePtrInput `pulumi:"windowsOsInfo"`
}

func (CustomImagePropertiesFromVmResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomImagePropertiesFromVmResponse)(nil)).Elem()
}

func (i CustomImagePropertiesFromVmResponseArgs) ToCustomImagePropertiesFromVmResponseOutput() CustomImagePropertiesFromVmResponseOutput {
	return i.ToCustomImagePropertiesFromVmResponseOutputWithContext(context.Background())
}

func (i CustomImagePropertiesFromVmResponseArgs) ToCustomImagePropertiesFromVmResponseOutputWithContext(ctx context.Context) CustomImagePropertiesFromVmResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomImagePropertiesFromVmResponseOutput)
}

func (i CustomImagePropertiesFromVmResponseArgs) ToCustomImagePropertiesFromVmResponsePtrOutput() CustomImagePropertiesFromVmResponsePtrOutput {
	return i.ToCustomImagePropertiesFromVmResponsePtrOutputWithContext(context.Background())
}

func (i CustomImagePropertiesFromVmResponseArgs) ToCustomImagePropertiesFromVmResponsePtrOutputWithContext(ctx context.Context) CustomImagePropertiesFromVmResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomImagePropertiesFromVmResponseOutput).ToCustomImagePropertiesFromVmResponsePtrOutputWithContext(ctx)
}









type CustomImagePropertiesFromVmResponsePtrInput interface {
	pulumi.Input

	ToCustomImagePropertiesFromVmResponsePtrOutput() CustomImagePropertiesFromVmResponsePtrOutput
	ToCustomImagePropertiesFromVmResponsePtrOutputWithContext(context.Context) CustomImagePropertiesFromVmResponsePtrOutput
}

type customImagePropertiesFromVmResponsePtrType CustomImagePropertiesFromVmResponseArgs

func CustomImagePropertiesFromVmResponsePtr(v *CustomImagePropertiesFromVmResponseArgs) CustomImagePropertiesFromVmResponsePtrInput {
	return (*customImagePropertiesFromVmResponsePtrType)(v)
}

func (*customImagePropertiesFromVmResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomImagePropertiesFromVmResponse)(nil)).Elem()
}

func (i *customImagePropertiesFromVmResponsePtrType) ToCustomImagePropertiesFromVmResponsePtrOutput() CustomImagePropertiesFromVmResponsePtrOutput {
	return i.ToCustomImagePropertiesFromVmResponsePtrOutputWithContext(context.Background())
}

func (i *customImagePropertiesFromVmResponsePtrType) ToCustomImagePropertiesFromVmResponsePtrOutputWithContext(ctx context.Context) CustomImagePropertiesFromVmResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomImagePropertiesFromVmResponsePtrOutput)
}

type CustomImagePropertiesFromVmResponseOutput struct{ *pulumi.OutputState }

func (CustomImagePropertiesFromVmResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomImagePropertiesFromVmResponse)(nil)).Elem()
}

func (o CustomImagePropertiesFromVmResponseOutput) ToCustomImagePropertiesFromVmResponseOutput() CustomImagePropertiesFromVmResponseOutput {
	return o
}

func (o CustomImagePropertiesFromVmResponseOutput) ToCustomImagePropertiesFromVmResponseOutputWithContext(ctx context.Context) CustomImagePropertiesFromVmResponseOutput {
	return o
}

func (o CustomImagePropertiesFromVmResponseOutput) ToCustomImagePropertiesFromVmResponsePtrOutput() CustomImagePropertiesFromVmResponsePtrOutput {
	return o.ToCustomImagePropertiesFromVmResponsePtrOutputWithContext(context.Background())
}

func (o CustomImagePropertiesFromVmResponseOutput) ToCustomImagePropertiesFromVmResponsePtrOutputWithContext(ctx context.Context) CustomImagePropertiesFromVmResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomImagePropertiesFromVmResponse) *CustomImagePropertiesFromVmResponse {
		return &v
	}).(CustomImagePropertiesFromVmResponsePtrOutput)
}

func (o CustomImagePropertiesFromVmResponseOutput) LinuxOsInfo() LinuxOsInfoResponsePtrOutput {
	return o.ApplyT(func(v CustomImagePropertiesFromVmResponse) *LinuxOsInfoResponse { return v.LinuxOsInfo }).(LinuxOsInfoResponsePtrOutput)
}

func (o CustomImagePropertiesFromVmResponseOutput) SourceVmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomImagePropertiesFromVmResponse) *string { return v.SourceVmId }).(pulumi.StringPtrOutput)
}

func (o CustomImagePropertiesFromVmResponseOutput) SysPrep() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CustomImagePropertiesFromVmResponse) *bool { return v.SysPrep }).(pulumi.BoolPtrOutput)
}

func (o CustomImagePropertiesFromVmResponseOutput) WindowsOsInfo() WindowsOsInfoResponsePtrOutput {
	return o.ApplyT(func(v CustomImagePropertiesFromVmResponse) *WindowsOsInfoResponse { return v.WindowsOsInfo }).(WindowsOsInfoResponsePtrOutput)
}

type CustomImagePropertiesFromVmResponsePtrOutput struct{ *pulumi.OutputState }

func (CustomImagePropertiesFromVmResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomImagePropertiesFromVmResponse)(nil)).Elem()
}

func (o CustomImagePropertiesFromVmResponsePtrOutput) ToCustomImagePropertiesFromVmResponsePtrOutput() CustomImagePropertiesFromVmResponsePtrOutput {
	return o
}

func (o CustomImagePropertiesFromVmResponsePtrOutput) ToCustomImagePropertiesFromVmResponsePtrOutputWithContext(ctx context.Context) CustomImagePropertiesFromVmResponsePtrOutput {
	return o
}

func (o CustomImagePropertiesFromVmResponsePtrOutput) Elem() CustomImagePropertiesFromVmResponseOutput {
	return o.ApplyT(func(v *CustomImagePropertiesFromVmResponse) CustomImagePropertiesFromVmResponse {
		if v != nil {
			return *v
		}
		var ret CustomImagePropertiesFromVmResponse
		return ret
	}).(CustomImagePropertiesFromVmResponseOutput)
}

func (o CustomImagePropertiesFromVmResponsePtrOutput) LinuxOsInfo() LinuxOsInfoResponsePtrOutput {
	return o.ApplyT(func(v *CustomImagePropertiesFromVmResponse) *LinuxOsInfoResponse {
		if v == nil {
			return nil
		}
		return v.LinuxOsInfo
	}).(LinuxOsInfoResponsePtrOutput)
}

func (o CustomImagePropertiesFromVmResponsePtrOutput) SourceVmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomImagePropertiesFromVmResponse) *string {
		if v == nil {
			return nil
		}
		return v.SourceVmId
	}).(pulumi.StringPtrOutput)
}

func (o CustomImagePropertiesFromVmResponsePtrOutput) SysPrep() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CustomImagePropertiesFromVmResponse) *bool {
		if v == nil {
			return nil
		}
		return v.SysPrep
	}).(pulumi.BoolPtrOutput)
}

func (o CustomImagePropertiesFromVmResponsePtrOutput) WindowsOsInfo() WindowsOsInfoResponsePtrOutput {
	return o.ApplyT(func(v *CustomImagePropertiesFromVmResponse) *WindowsOsInfoResponse {
		if v == nil {
			return nil
		}
		return v.WindowsOsInfo
	}).(WindowsOsInfoResponsePtrOutput)
}

type DayDetails struct {
	Time *string `pulumi:"time"`
}





type DayDetailsInput interface {
	pulumi.Input

	ToDayDetailsOutput() DayDetailsOutput
	ToDayDetailsOutputWithContext(context.Context) DayDetailsOutput
}

type DayDetailsArgs struct {
	Time pulumi.StringPtrInput `pulumi:"time"`
}

func (DayDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DayDetails)(nil)).Elem()
}

func (i DayDetailsArgs) ToDayDetailsOutput() DayDetailsOutput {
	return i.ToDayDetailsOutputWithContext(context.Background())
}

func (i DayDetailsArgs) ToDayDetailsOutputWithContext(ctx context.Context) DayDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DayDetailsOutput)
}

func (i DayDetailsArgs) ToDayDetailsPtrOutput() DayDetailsPtrOutput {
	return i.ToDayDetailsPtrOutputWithContext(context.Background())
}

func (i DayDetailsArgs) ToDayDetailsPtrOutputWithContext(ctx context.Context) DayDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DayDetailsOutput).ToDayDetailsPtrOutputWithContext(ctx)
}









type DayDetailsPtrInput interface {
	pulumi.Input

	ToDayDetailsPtrOutput() DayDetailsPtrOutput
	ToDayDetailsPtrOutputWithContext(context.Context) DayDetailsPtrOutput
}

type dayDetailsPtrType DayDetailsArgs

func DayDetailsPtr(v *DayDetailsArgs) DayDetailsPtrInput {
	return (*dayDetailsPtrType)(v)
}

func (*dayDetailsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DayDetails)(nil)).Elem()
}

func (i *dayDetailsPtrType) ToDayDetailsPtrOutput() DayDetailsPtrOutput {
	return i.ToDayDetailsPtrOutputWithContext(context.Background())
}

func (i *dayDetailsPtrType) ToDayDetailsPtrOutputWithContext(ctx context.Context) DayDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DayDetailsPtrOutput)
}

type DayDetailsOutput struct{ *pulumi.OutputState }

func (DayDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DayDetails)(nil)).Elem()
}

func (o DayDetailsOutput) ToDayDetailsOutput() DayDetailsOutput {
	return o
}

func (o DayDetailsOutput) ToDayDetailsOutputWithContext(ctx context.Context) DayDetailsOutput {
	return o
}

func (o DayDetailsOutput) ToDayDetailsPtrOutput() DayDetailsPtrOutput {
	return o.ToDayDetailsPtrOutputWithContext(context.Background())
}

func (o DayDetailsOutput) ToDayDetailsPtrOutputWithContext(ctx context.Context) DayDetailsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DayDetails) *DayDetails {
		return &v
	}).(DayDetailsPtrOutput)
}

func (o DayDetailsOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DayDetails) *string { return v.Time }).(pulumi.StringPtrOutput)
}

type DayDetailsPtrOutput struct{ *pulumi.OutputState }

func (DayDetailsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DayDetails)(nil)).Elem()
}

func (o DayDetailsPtrOutput) ToDayDetailsPtrOutput() DayDetailsPtrOutput {
	return o
}

func (o DayDetailsPtrOutput) ToDayDetailsPtrOutputWithContext(ctx context.Context) DayDetailsPtrOutput {
	return o
}

func (o DayDetailsPtrOutput) Elem() DayDetailsOutput {
	return o.ApplyT(func(v *DayDetails) DayDetails {
		if v != nil {
			return *v
		}
		var ret DayDetails
		return ret
	}).(DayDetailsOutput)
}

func (o DayDetailsPtrOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DayDetails) *string {
		if v == nil {
			return nil
		}
		return v.Time
	}).(pulumi.StringPtrOutput)
}

type DayDetailsResponse struct {
	Time *string `pulumi:"time"`
}





type DayDetailsResponseInput interface {
	pulumi.Input

	ToDayDetailsResponseOutput() DayDetailsResponseOutput
	ToDayDetailsResponseOutputWithContext(context.Context) DayDetailsResponseOutput
}

type DayDetailsResponseArgs struct {
	Time pulumi.StringPtrInput `pulumi:"time"`
}

func (DayDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DayDetailsResponse)(nil)).Elem()
}

func (i DayDetailsResponseArgs) ToDayDetailsResponseOutput() DayDetailsResponseOutput {
	return i.ToDayDetailsResponseOutputWithContext(context.Background())
}

func (i DayDetailsResponseArgs) ToDayDetailsResponseOutputWithContext(ctx context.Context) DayDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DayDetailsResponseOutput)
}

func (i DayDetailsResponseArgs) ToDayDetailsResponsePtrOutput() DayDetailsResponsePtrOutput {
	return i.ToDayDetailsResponsePtrOutputWithContext(context.Background())
}

func (i DayDetailsResponseArgs) ToDayDetailsResponsePtrOutputWithContext(ctx context.Context) DayDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DayDetailsResponseOutput).ToDayDetailsResponsePtrOutputWithContext(ctx)
}









type DayDetailsResponsePtrInput interface {
	pulumi.Input

	ToDayDetailsResponsePtrOutput() DayDetailsResponsePtrOutput
	ToDayDetailsResponsePtrOutputWithContext(context.Context) DayDetailsResponsePtrOutput
}

type dayDetailsResponsePtrType DayDetailsResponseArgs

func DayDetailsResponsePtr(v *DayDetailsResponseArgs) DayDetailsResponsePtrInput {
	return (*dayDetailsResponsePtrType)(v)
}

func (*dayDetailsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DayDetailsResponse)(nil)).Elem()
}

func (i *dayDetailsResponsePtrType) ToDayDetailsResponsePtrOutput() DayDetailsResponsePtrOutput {
	return i.ToDayDetailsResponsePtrOutputWithContext(context.Background())
}

func (i *dayDetailsResponsePtrType) ToDayDetailsResponsePtrOutputWithContext(ctx context.Context) DayDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DayDetailsResponsePtrOutput)
}

type DayDetailsResponseOutput struct{ *pulumi.OutputState }

func (DayDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DayDetailsResponse)(nil)).Elem()
}

func (o DayDetailsResponseOutput) ToDayDetailsResponseOutput() DayDetailsResponseOutput {
	return o
}

func (o DayDetailsResponseOutput) ToDayDetailsResponseOutputWithContext(ctx context.Context) DayDetailsResponseOutput {
	return o
}

func (o DayDetailsResponseOutput) ToDayDetailsResponsePtrOutput() DayDetailsResponsePtrOutput {
	return o.ToDayDetailsResponsePtrOutputWithContext(context.Background())
}

func (o DayDetailsResponseOutput) ToDayDetailsResponsePtrOutputWithContext(ctx context.Context) DayDetailsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DayDetailsResponse) *DayDetailsResponse {
		return &v
	}).(DayDetailsResponsePtrOutput)
}

func (o DayDetailsResponseOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DayDetailsResponse) *string { return v.Time }).(pulumi.StringPtrOutput)
}

type DayDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (DayDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DayDetailsResponse)(nil)).Elem()
}

func (o DayDetailsResponsePtrOutput) ToDayDetailsResponsePtrOutput() DayDetailsResponsePtrOutput {
	return o
}

func (o DayDetailsResponsePtrOutput) ToDayDetailsResponsePtrOutputWithContext(ctx context.Context) DayDetailsResponsePtrOutput {
	return o
}

func (o DayDetailsResponsePtrOutput) Elem() DayDetailsResponseOutput {
	return o.ApplyT(func(v *DayDetailsResponse) DayDetailsResponse {
		if v != nil {
			return *v
		}
		var ret DayDetailsResponse
		return ret
	}).(DayDetailsResponseOutput)
}

func (o DayDetailsResponsePtrOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DayDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Time
	}).(pulumi.StringPtrOutput)
}

type FormulaPropertiesFromVm struct {
	LabVmId *string `pulumi:"labVmId"`
}





type FormulaPropertiesFromVmInput interface {
	pulumi.Input

	ToFormulaPropertiesFromVmOutput() FormulaPropertiesFromVmOutput
	ToFormulaPropertiesFromVmOutputWithContext(context.Context) FormulaPropertiesFromVmOutput
}

type FormulaPropertiesFromVmArgs struct {
	LabVmId pulumi.StringPtrInput `pulumi:"labVmId"`
}

func (FormulaPropertiesFromVmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FormulaPropertiesFromVm)(nil)).Elem()
}

func (i FormulaPropertiesFromVmArgs) ToFormulaPropertiesFromVmOutput() FormulaPropertiesFromVmOutput {
	return i.ToFormulaPropertiesFromVmOutputWithContext(context.Background())
}

func (i FormulaPropertiesFromVmArgs) ToFormulaPropertiesFromVmOutputWithContext(ctx context.Context) FormulaPropertiesFromVmOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FormulaPropertiesFromVmOutput)
}

func (i FormulaPropertiesFromVmArgs) ToFormulaPropertiesFromVmPtrOutput() FormulaPropertiesFromVmPtrOutput {
	return i.ToFormulaPropertiesFromVmPtrOutputWithContext(context.Background())
}

func (i FormulaPropertiesFromVmArgs) ToFormulaPropertiesFromVmPtrOutputWithContext(ctx context.Context) FormulaPropertiesFromVmPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FormulaPropertiesFromVmOutput).ToFormulaPropertiesFromVmPtrOutputWithContext(ctx)
}









type FormulaPropertiesFromVmPtrInput interface {
	pulumi.Input

	ToFormulaPropertiesFromVmPtrOutput() FormulaPropertiesFromVmPtrOutput
	ToFormulaPropertiesFromVmPtrOutputWithContext(context.Context) FormulaPropertiesFromVmPtrOutput
}

type formulaPropertiesFromVmPtrType FormulaPropertiesFromVmArgs

func FormulaPropertiesFromVmPtr(v *FormulaPropertiesFromVmArgs) FormulaPropertiesFromVmPtrInput {
	return (*formulaPropertiesFromVmPtrType)(v)
}

func (*formulaPropertiesFromVmPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FormulaPropertiesFromVm)(nil)).Elem()
}

func (i *formulaPropertiesFromVmPtrType) ToFormulaPropertiesFromVmPtrOutput() FormulaPropertiesFromVmPtrOutput {
	return i.ToFormulaPropertiesFromVmPtrOutputWithContext(context.Background())
}

func (i *formulaPropertiesFromVmPtrType) ToFormulaPropertiesFromVmPtrOutputWithContext(ctx context.Context) FormulaPropertiesFromVmPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FormulaPropertiesFromVmPtrOutput)
}

type FormulaPropertiesFromVmOutput struct{ *pulumi.OutputState }

func (FormulaPropertiesFromVmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FormulaPropertiesFromVm)(nil)).Elem()
}

func (o FormulaPropertiesFromVmOutput) ToFormulaPropertiesFromVmOutput() FormulaPropertiesFromVmOutput {
	return o
}

func (o FormulaPropertiesFromVmOutput) ToFormulaPropertiesFromVmOutputWithContext(ctx context.Context) FormulaPropertiesFromVmOutput {
	return o
}

func (o FormulaPropertiesFromVmOutput) ToFormulaPropertiesFromVmPtrOutput() FormulaPropertiesFromVmPtrOutput {
	return o.ToFormulaPropertiesFromVmPtrOutputWithContext(context.Background())
}

func (o FormulaPropertiesFromVmOutput) ToFormulaPropertiesFromVmPtrOutputWithContext(ctx context.Context) FormulaPropertiesFromVmPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FormulaPropertiesFromVm) *FormulaPropertiesFromVm {
		return &v
	}).(FormulaPropertiesFromVmPtrOutput)
}

func (o FormulaPropertiesFromVmOutput) LabVmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FormulaPropertiesFromVm) *string { return v.LabVmId }).(pulumi.StringPtrOutput)
}

type FormulaPropertiesFromVmPtrOutput struct{ *pulumi.OutputState }

func (FormulaPropertiesFromVmPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FormulaPropertiesFromVm)(nil)).Elem()
}

func (o FormulaPropertiesFromVmPtrOutput) ToFormulaPropertiesFromVmPtrOutput() FormulaPropertiesFromVmPtrOutput {
	return o
}

func (o FormulaPropertiesFromVmPtrOutput) ToFormulaPropertiesFromVmPtrOutputWithContext(ctx context.Context) FormulaPropertiesFromVmPtrOutput {
	return o
}

func (o FormulaPropertiesFromVmPtrOutput) Elem() FormulaPropertiesFromVmOutput {
	return o.ApplyT(func(v *FormulaPropertiesFromVm) FormulaPropertiesFromVm {
		if v != nil {
			return *v
		}
		var ret FormulaPropertiesFromVm
		return ret
	}).(FormulaPropertiesFromVmOutput)
}

func (o FormulaPropertiesFromVmPtrOutput) LabVmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FormulaPropertiesFromVm) *string {
		if v == nil {
			return nil
		}
		return v.LabVmId
	}).(pulumi.StringPtrOutput)
}

type FormulaPropertiesFromVmResponse struct {
	LabVmId *string `pulumi:"labVmId"`
}





type FormulaPropertiesFromVmResponseInput interface {
	pulumi.Input

	ToFormulaPropertiesFromVmResponseOutput() FormulaPropertiesFromVmResponseOutput
	ToFormulaPropertiesFromVmResponseOutputWithContext(context.Context) FormulaPropertiesFromVmResponseOutput
}

type FormulaPropertiesFromVmResponseArgs struct {
	LabVmId pulumi.StringPtrInput `pulumi:"labVmId"`
}

func (FormulaPropertiesFromVmResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FormulaPropertiesFromVmResponse)(nil)).Elem()
}

func (i FormulaPropertiesFromVmResponseArgs) ToFormulaPropertiesFromVmResponseOutput() FormulaPropertiesFromVmResponseOutput {
	return i.ToFormulaPropertiesFromVmResponseOutputWithContext(context.Background())
}

func (i FormulaPropertiesFromVmResponseArgs) ToFormulaPropertiesFromVmResponseOutputWithContext(ctx context.Context) FormulaPropertiesFromVmResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FormulaPropertiesFromVmResponseOutput)
}

func (i FormulaPropertiesFromVmResponseArgs) ToFormulaPropertiesFromVmResponsePtrOutput() FormulaPropertiesFromVmResponsePtrOutput {
	return i.ToFormulaPropertiesFromVmResponsePtrOutputWithContext(context.Background())
}

func (i FormulaPropertiesFromVmResponseArgs) ToFormulaPropertiesFromVmResponsePtrOutputWithContext(ctx context.Context) FormulaPropertiesFromVmResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FormulaPropertiesFromVmResponseOutput).ToFormulaPropertiesFromVmResponsePtrOutputWithContext(ctx)
}









type FormulaPropertiesFromVmResponsePtrInput interface {
	pulumi.Input

	ToFormulaPropertiesFromVmResponsePtrOutput() FormulaPropertiesFromVmResponsePtrOutput
	ToFormulaPropertiesFromVmResponsePtrOutputWithContext(context.Context) FormulaPropertiesFromVmResponsePtrOutput
}

type formulaPropertiesFromVmResponsePtrType FormulaPropertiesFromVmResponseArgs

func FormulaPropertiesFromVmResponsePtr(v *FormulaPropertiesFromVmResponseArgs) FormulaPropertiesFromVmResponsePtrInput {
	return (*formulaPropertiesFromVmResponsePtrType)(v)
}

func (*formulaPropertiesFromVmResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FormulaPropertiesFromVmResponse)(nil)).Elem()
}

func (i *formulaPropertiesFromVmResponsePtrType) ToFormulaPropertiesFromVmResponsePtrOutput() FormulaPropertiesFromVmResponsePtrOutput {
	return i.ToFormulaPropertiesFromVmResponsePtrOutputWithContext(context.Background())
}

func (i *formulaPropertiesFromVmResponsePtrType) ToFormulaPropertiesFromVmResponsePtrOutputWithContext(ctx context.Context) FormulaPropertiesFromVmResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FormulaPropertiesFromVmResponsePtrOutput)
}

type FormulaPropertiesFromVmResponseOutput struct{ *pulumi.OutputState }

func (FormulaPropertiesFromVmResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FormulaPropertiesFromVmResponse)(nil)).Elem()
}

func (o FormulaPropertiesFromVmResponseOutput) ToFormulaPropertiesFromVmResponseOutput() FormulaPropertiesFromVmResponseOutput {
	return o
}

func (o FormulaPropertiesFromVmResponseOutput) ToFormulaPropertiesFromVmResponseOutputWithContext(ctx context.Context) FormulaPropertiesFromVmResponseOutput {
	return o
}

func (o FormulaPropertiesFromVmResponseOutput) ToFormulaPropertiesFromVmResponsePtrOutput() FormulaPropertiesFromVmResponsePtrOutput {
	return o.ToFormulaPropertiesFromVmResponsePtrOutputWithContext(context.Background())
}

func (o FormulaPropertiesFromVmResponseOutput) ToFormulaPropertiesFromVmResponsePtrOutputWithContext(ctx context.Context) FormulaPropertiesFromVmResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FormulaPropertiesFromVmResponse) *FormulaPropertiesFromVmResponse {
		return &v
	}).(FormulaPropertiesFromVmResponsePtrOutput)
}

func (o FormulaPropertiesFromVmResponseOutput) LabVmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FormulaPropertiesFromVmResponse) *string { return v.LabVmId }).(pulumi.StringPtrOutput)
}

type FormulaPropertiesFromVmResponsePtrOutput struct{ *pulumi.OutputState }

func (FormulaPropertiesFromVmResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FormulaPropertiesFromVmResponse)(nil)).Elem()
}

func (o FormulaPropertiesFromVmResponsePtrOutput) ToFormulaPropertiesFromVmResponsePtrOutput() FormulaPropertiesFromVmResponsePtrOutput {
	return o
}

func (o FormulaPropertiesFromVmResponsePtrOutput) ToFormulaPropertiesFromVmResponsePtrOutputWithContext(ctx context.Context) FormulaPropertiesFromVmResponsePtrOutput {
	return o
}

func (o FormulaPropertiesFromVmResponsePtrOutput) Elem() FormulaPropertiesFromVmResponseOutput {
	return o.ApplyT(func(v *FormulaPropertiesFromVmResponse) FormulaPropertiesFromVmResponse {
		if v != nil {
			return *v
		}
		var ret FormulaPropertiesFromVmResponse
		return ret
	}).(FormulaPropertiesFromVmResponseOutput)
}

func (o FormulaPropertiesFromVmResponsePtrOutput) LabVmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FormulaPropertiesFromVmResponse) *string {
		if v == nil {
			return nil
		}
		return v.LabVmId
	}).(pulumi.StringPtrOutput)
}

type GalleryImageReference struct {
	Offer     *string `pulumi:"offer"`
	OsType    *string `pulumi:"osType"`
	Publisher *string `pulumi:"publisher"`
	Sku       *string `pulumi:"sku"`
	Version   *string `pulumi:"version"`
}





type GalleryImageReferenceInput interface {
	pulumi.Input

	ToGalleryImageReferenceOutput() GalleryImageReferenceOutput
	ToGalleryImageReferenceOutputWithContext(context.Context) GalleryImageReferenceOutput
}

type GalleryImageReferenceArgs struct {
	Offer     pulumi.StringPtrInput `pulumi:"offer"`
	OsType    pulumi.StringPtrInput `pulumi:"osType"`
	Publisher pulumi.StringPtrInput `pulumi:"publisher"`
	Sku       pulumi.StringPtrInput `pulumi:"sku"`
	Version   pulumi.StringPtrInput `pulumi:"version"`
}

func (GalleryImageReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryImageReference)(nil)).Elem()
}

func (i GalleryImageReferenceArgs) ToGalleryImageReferenceOutput() GalleryImageReferenceOutput {
	return i.ToGalleryImageReferenceOutputWithContext(context.Background())
}

func (i GalleryImageReferenceArgs) ToGalleryImageReferenceOutputWithContext(ctx context.Context) GalleryImageReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryImageReferenceOutput)
}

func (i GalleryImageReferenceArgs) ToGalleryImageReferencePtrOutput() GalleryImageReferencePtrOutput {
	return i.ToGalleryImageReferencePtrOutputWithContext(context.Background())
}

func (i GalleryImageReferenceArgs) ToGalleryImageReferencePtrOutputWithContext(ctx context.Context) GalleryImageReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryImageReferenceOutput).ToGalleryImageReferencePtrOutputWithContext(ctx)
}









type GalleryImageReferencePtrInput interface {
	pulumi.Input

	ToGalleryImageReferencePtrOutput() GalleryImageReferencePtrOutput
	ToGalleryImageReferencePtrOutputWithContext(context.Context) GalleryImageReferencePtrOutput
}

type galleryImageReferencePtrType GalleryImageReferenceArgs

func GalleryImageReferencePtr(v *GalleryImageReferenceArgs) GalleryImageReferencePtrInput {
	return (*galleryImageReferencePtrType)(v)
}

func (*galleryImageReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryImageReference)(nil)).Elem()
}

func (i *galleryImageReferencePtrType) ToGalleryImageReferencePtrOutput() GalleryImageReferencePtrOutput {
	return i.ToGalleryImageReferencePtrOutputWithContext(context.Background())
}

func (i *galleryImageReferencePtrType) ToGalleryImageReferencePtrOutputWithContext(ctx context.Context) GalleryImageReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryImageReferencePtrOutput)
}

type GalleryImageReferenceOutput struct{ *pulumi.OutputState }

func (GalleryImageReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryImageReference)(nil)).Elem()
}

func (o GalleryImageReferenceOutput) ToGalleryImageReferenceOutput() GalleryImageReferenceOutput {
	return o
}

func (o GalleryImageReferenceOutput) ToGalleryImageReferenceOutputWithContext(ctx context.Context) GalleryImageReferenceOutput {
	return o
}

func (o GalleryImageReferenceOutput) ToGalleryImageReferencePtrOutput() GalleryImageReferencePtrOutput {
	return o.ToGalleryImageReferencePtrOutputWithContext(context.Background())
}

func (o GalleryImageReferenceOutput) ToGalleryImageReferencePtrOutputWithContext(ctx context.Context) GalleryImageReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GalleryImageReference) *GalleryImageReference {
		return &v
	}).(GalleryImageReferencePtrOutput)
}

func (o GalleryImageReferenceOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageReference) *string { return v.Offer }).(pulumi.StringPtrOutput)
}

func (o GalleryImageReferenceOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageReference) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o GalleryImageReferenceOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageReference) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o GalleryImageReferenceOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageReference) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

func (o GalleryImageReferenceOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageReference) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type GalleryImageReferencePtrOutput struct{ *pulumi.OutputState }

func (GalleryImageReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryImageReference)(nil)).Elem()
}

func (o GalleryImageReferencePtrOutput) ToGalleryImageReferencePtrOutput() GalleryImageReferencePtrOutput {
	return o
}

func (o GalleryImageReferencePtrOutput) ToGalleryImageReferencePtrOutputWithContext(ctx context.Context) GalleryImageReferencePtrOutput {
	return o
}

func (o GalleryImageReferencePtrOutput) Elem() GalleryImageReferenceOutput {
	return o.ApplyT(func(v *GalleryImageReference) GalleryImageReference {
		if v != nil {
			return *v
		}
		var ret GalleryImageReference
		return ret
	}).(GalleryImageReferenceOutput)
}

func (o GalleryImageReferencePtrOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImageReference) *string {
		if v == nil {
			return nil
		}
		return v.Offer
	}).(pulumi.StringPtrOutput)
}

func (o GalleryImageReferencePtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImageReference) *string {
		if v == nil {
			return nil
		}
		return v.OsType
	}).(pulumi.StringPtrOutput)
}

func (o GalleryImageReferencePtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImageReference) *string {
		if v == nil {
			return nil
		}
		return v.Publisher
	}).(pulumi.StringPtrOutput)
}

func (o GalleryImageReferencePtrOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImageReference) *string {
		if v == nil {
			return nil
		}
		return v.Sku
	}).(pulumi.StringPtrOutput)
}

func (o GalleryImageReferencePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImageReference) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type GalleryImageReferenceResponse struct {
	Offer     *string `pulumi:"offer"`
	OsType    *string `pulumi:"osType"`
	Publisher *string `pulumi:"publisher"`
	Sku       *string `pulumi:"sku"`
	Version   *string `pulumi:"version"`
}





type GalleryImageReferenceResponseInput interface {
	pulumi.Input

	ToGalleryImageReferenceResponseOutput() GalleryImageReferenceResponseOutput
	ToGalleryImageReferenceResponseOutputWithContext(context.Context) GalleryImageReferenceResponseOutput
}

type GalleryImageReferenceResponseArgs struct {
	Offer     pulumi.StringPtrInput `pulumi:"offer"`
	OsType    pulumi.StringPtrInput `pulumi:"osType"`
	Publisher pulumi.StringPtrInput `pulumi:"publisher"`
	Sku       pulumi.StringPtrInput `pulumi:"sku"`
	Version   pulumi.StringPtrInput `pulumi:"version"`
}

func (GalleryImageReferenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryImageReferenceResponse)(nil)).Elem()
}

func (i GalleryImageReferenceResponseArgs) ToGalleryImageReferenceResponseOutput() GalleryImageReferenceResponseOutput {
	return i.ToGalleryImageReferenceResponseOutputWithContext(context.Background())
}

func (i GalleryImageReferenceResponseArgs) ToGalleryImageReferenceResponseOutputWithContext(ctx context.Context) GalleryImageReferenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryImageReferenceResponseOutput)
}

func (i GalleryImageReferenceResponseArgs) ToGalleryImageReferenceResponsePtrOutput() GalleryImageReferenceResponsePtrOutput {
	return i.ToGalleryImageReferenceResponsePtrOutputWithContext(context.Background())
}

func (i GalleryImageReferenceResponseArgs) ToGalleryImageReferenceResponsePtrOutputWithContext(ctx context.Context) GalleryImageReferenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryImageReferenceResponseOutput).ToGalleryImageReferenceResponsePtrOutputWithContext(ctx)
}









type GalleryImageReferenceResponsePtrInput interface {
	pulumi.Input

	ToGalleryImageReferenceResponsePtrOutput() GalleryImageReferenceResponsePtrOutput
	ToGalleryImageReferenceResponsePtrOutputWithContext(context.Context) GalleryImageReferenceResponsePtrOutput
}

type galleryImageReferenceResponsePtrType GalleryImageReferenceResponseArgs

func GalleryImageReferenceResponsePtr(v *GalleryImageReferenceResponseArgs) GalleryImageReferenceResponsePtrInput {
	return (*galleryImageReferenceResponsePtrType)(v)
}

func (*galleryImageReferenceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryImageReferenceResponse)(nil)).Elem()
}

func (i *galleryImageReferenceResponsePtrType) ToGalleryImageReferenceResponsePtrOutput() GalleryImageReferenceResponsePtrOutput {
	return i.ToGalleryImageReferenceResponsePtrOutputWithContext(context.Background())
}

func (i *galleryImageReferenceResponsePtrType) ToGalleryImageReferenceResponsePtrOutputWithContext(ctx context.Context) GalleryImageReferenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryImageReferenceResponsePtrOutput)
}

type GalleryImageReferenceResponseOutput struct{ *pulumi.OutputState }

func (GalleryImageReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryImageReferenceResponse)(nil)).Elem()
}

func (o GalleryImageReferenceResponseOutput) ToGalleryImageReferenceResponseOutput() GalleryImageReferenceResponseOutput {
	return o
}

func (o GalleryImageReferenceResponseOutput) ToGalleryImageReferenceResponseOutputWithContext(ctx context.Context) GalleryImageReferenceResponseOutput {
	return o
}

func (o GalleryImageReferenceResponseOutput) ToGalleryImageReferenceResponsePtrOutput() GalleryImageReferenceResponsePtrOutput {
	return o.ToGalleryImageReferenceResponsePtrOutputWithContext(context.Background())
}

func (o GalleryImageReferenceResponseOutput) ToGalleryImageReferenceResponsePtrOutputWithContext(ctx context.Context) GalleryImageReferenceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GalleryImageReferenceResponse) *GalleryImageReferenceResponse {
		return &v
	}).(GalleryImageReferenceResponsePtrOutput)
}

func (o GalleryImageReferenceResponseOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageReferenceResponse) *string { return v.Offer }).(pulumi.StringPtrOutput)
}

func (o GalleryImageReferenceResponseOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageReferenceResponse) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o GalleryImageReferenceResponseOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageReferenceResponse) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o GalleryImageReferenceResponseOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageReferenceResponse) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

func (o GalleryImageReferenceResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageReferenceResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type GalleryImageReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (GalleryImageReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryImageReferenceResponse)(nil)).Elem()
}

func (o GalleryImageReferenceResponsePtrOutput) ToGalleryImageReferenceResponsePtrOutput() GalleryImageReferenceResponsePtrOutput {
	return o
}

func (o GalleryImageReferenceResponsePtrOutput) ToGalleryImageReferenceResponsePtrOutputWithContext(ctx context.Context) GalleryImageReferenceResponsePtrOutput {
	return o
}

func (o GalleryImageReferenceResponsePtrOutput) Elem() GalleryImageReferenceResponseOutput {
	return o.ApplyT(func(v *GalleryImageReferenceResponse) GalleryImageReferenceResponse {
		if v != nil {
			return *v
		}
		var ret GalleryImageReferenceResponse
		return ret
	}).(GalleryImageReferenceResponseOutput)
}

func (o GalleryImageReferenceResponsePtrOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImageReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Offer
	}).(pulumi.StringPtrOutput)
}

func (o GalleryImageReferenceResponsePtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImageReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.OsType
	}).(pulumi.StringPtrOutput)
}

func (o GalleryImageReferenceResponsePtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImageReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Publisher
	}).(pulumi.StringPtrOutput)
}

func (o GalleryImageReferenceResponsePtrOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImageReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Sku
	}).(pulumi.StringPtrOutput)
}

func (o GalleryImageReferenceResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImageReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type HourDetails struct {
	Minute *int `pulumi:"minute"`
}





type HourDetailsInput interface {
	pulumi.Input

	ToHourDetailsOutput() HourDetailsOutput
	ToHourDetailsOutputWithContext(context.Context) HourDetailsOutput
}

type HourDetailsArgs struct {
	Minute pulumi.IntPtrInput `pulumi:"minute"`
}

func (HourDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HourDetails)(nil)).Elem()
}

func (i HourDetailsArgs) ToHourDetailsOutput() HourDetailsOutput {
	return i.ToHourDetailsOutputWithContext(context.Background())
}

func (i HourDetailsArgs) ToHourDetailsOutputWithContext(ctx context.Context) HourDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HourDetailsOutput)
}

func (i HourDetailsArgs) ToHourDetailsPtrOutput() HourDetailsPtrOutput {
	return i.ToHourDetailsPtrOutputWithContext(context.Background())
}

func (i HourDetailsArgs) ToHourDetailsPtrOutputWithContext(ctx context.Context) HourDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HourDetailsOutput).ToHourDetailsPtrOutputWithContext(ctx)
}









type HourDetailsPtrInput interface {
	pulumi.Input

	ToHourDetailsPtrOutput() HourDetailsPtrOutput
	ToHourDetailsPtrOutputWithContext(context.Context) HourDetailsPtrOutput
}

type hourDetailsPtrType HourDetailsArgs

func HourDetailsPtr(v *HourDetailsArgs) HourDetailsPtrInput {
	return (*hourDetailsPtrType)(v)
}

func (*hourDetailsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HourDetails)(nil)).Elem()
}

func (i *hourDetailsPtrType) ToHourDetailsPtrOutput() HourDetailsPtrOutput {
	return i.ToHourDetailsPtrOutputWithContext(context.Background())
}

func (i *hourDetailsPtrType) ToHourDetailsPtrOutputWithContext(ctx context.Context) HourDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HourDetailsPtrOutput)
}

type HourDetailsOutput struct{ *pulumi.OutputState }

func (HourDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HourDetails)(nil)).Elem()
}

func (o HourDetailsOutput) ToHourDetailsOutput() HourDetailsOutput {
	return o
}

func (o HourDetailsOutput) ToHourDetailsOutputWithContext(ctx context.Context) HourDetailsOutput {
	return o
}

func (o HourDetailsOutput) ToHourDetailsPtrOutput() HourDetailsPtrOutput {
	return o.ToHourDetailsPtrOutputWithContext(context.Background())
}

func (o HourDetailsOutput) ToHourDetailsPtrOutputWithContext(ctx context.Context) HourDetailsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HourDetails) *HourDetails {
		return &v
	}).(HourDetailsPtrOutput)
}

func (o HourDetailsOutput) Minute() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HourDetails) *int { return v.Minute }).(pulumi.IntPtrOutput)
}

type HourDetailsPtrOutput struct{ *pulumi.OutputState }

func (HourDetailsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HourDetails)(nil)).Elem()
}

func (o HourDetailsPtrOutput) ToHourDetailsPtrOutput() HourDetailsPtrOutput {
	return o
}

func (o HourDetailsPtrOutput) ToHourDetailsPtrOutputWithContext(ctx context.Context) HourDetailsPtrOutput {
	return o
}

func (o HourDetailsPtrOutput) Elem() HourDetailsOutput {
	return o.ApplyT(func(v *HourDetails) HourDetails {
		if v != nil {
			return *v
		}
		var ret HourDetails
		return ret
	}).(HourDetailsOutput)
}

func (o HourDetailsPtrOutput) Minute() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HourDetails) *int {
		if v == nil {
			return nil
		}
		return v.Minute
	}).(pulumi.IntPtrOutput)
}

type HourDetailsResponse struct {
	Minute *int `pulumi:"minute"`
}





type HourDetailsResponseInput interface {
	pulumi.Input

	ToHourDetailsResponseOutput() HourDetailsResponseOutput
	ToHourDetailsResponseOutputWithContext(context.Context) HourDetailsResponseOutput
}

type HourDetailsResponseArgs struct {
	Minute pulumi.IntPtrInput `pulumi:"minute"`
}

func (HourDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HourDetailsResponse)(nil)).Elem()
}

func (i HourDetailsResponseArgs) ToHourDetailsResponseOutput() HourDetailsResponseOutput {
	return i.ToHourDetailsResponseOutputWithContext(context.Background())
}

func (i HourDetailsResponseArgs) ToHourDetailsResponseOutputWithContext(ctx context.Context) HourDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HourDetailsResponseOutput)
}

func (i HourDetailsResponseArgs) ToHourDetailsResponsePtrOutput() HourDetailsResponsePtrOutput {
	return i.ToHourDetailsResponsePtrOutputWithContext(context.Background())
}

func (i HourDetailsResponseArgs) ToHourDetailsResponsePtrOutputWithContext(ctx context.Context) HourDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HourDetailsResponseOutput).ToHourDetailsResponsePtrOutputWithContext(ctx)
}









type HourDetailsResponsePtrInput interface {
	pulumi.Input

	ToHourDetailsResponsePtrOutput() HourDetailsResponsePtrOutput
	ToHourDetailsResponsePtrOutputWithContext(context.Context) HourDetailsResponsePtrOutput
}

type hourDetailsResponsePtrType HourDetailsResponseArgs

func HourDetailsResponsePtr(v *HourDetailsResponseArgs) HourDetailsResponsePtrInput {
	return (*hourDetailsResponsePtrType)(v)
}

func (*hourDetailsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HourDetailsResponse)(nil)).Elem()
}

func (i *hourDetailsResponsePtrType) ToHourDetailsResponsePtrOutput() HourDetailsResponsePtrOutput {
	return i.ToHourDetailsResponsePtrOutputWithContext(context.Background())
}

func (i *hourDetailsResponsePtrType) ToHourDetailsResponsePtrOutputWithContext(ctx context.Context) HourDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HourDetailsResponsePtrOutput)
}

type HourDetailsResponseOutput struct{ *pulumi.OutputState }

func (HourDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HourDetailsResponse)(nil)).Elem()
}

func (o HourDetailsResponseOutput) ToHourDetailsResponseOutput() HourDetailsResponseOutput {
	return o
}

func (o HourDetailsResponseOutput) ToHourDetailsResponseOutputWithContext(ctx context.Context) HourDetailsResponseOutput {
	return o
}

func (o HourDetailsResponseOutput) ToHourDetailsResponsePtrOutput() HourDetailsResponsePtrOutput {
	return o.ToHourDetailsResponsePtrOutputWithContext(context.Background())
}

func (o HourDetailsResponseOutput) ToHourDetailsResponsePtrOutputWithContext(ctx context.Context) HourDetailsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HourDetailsResponse) *HourDetailsResponse {
		return &v
	}).(HourDetailsResponsePtrOutput)
}

func (o HourDetailsResponseOutput) Minute() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HourDetailsResponse) *int { return v.Minute }).(pulumi.IntPtrOutput)
}

type HourDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (HourDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HourDetailsResponse)(nil)).Elem()
}

func (o HourDetailsResponsePtrOutput) ToHourDetailsResponsePtrOutput() HourDetailsResponsePtrOutput {
	return o
}

func (o HourDetailsResponsePtrOutput) ToHourDetailsResponsePtrOutputWithContext(ctx context.Context) HourDetailsResponsePtrOutput {
	return o
}

func (o HourDetailsResponsePtrOutput) Elem() HourDetailsResponseOutput {
	return o.ApplyT(func(v *HourDetailsResponse) HourDetailsResponse {
		if v != nil {
			return *v
		}
		var ret HourDetailsResponse
		return ret
	}).(HourDetailsResponseOutput)
}

func (o HourDetailsResponsePtrOutput) Minute() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HourDetailsResponse) *int {
		if v == nil {
			return nil
		}
		return v.Minute
	}).(pulumi.IntPtrOutput)
}

type LabVhdResponse struct {
	Id *string `pulumi:"id"`
}





type LabVhdResponseInput interface {
	pulumi.Input

	ToLabVhdResponseOutput() LabVhdResponseOutput
	ToLabVhdResponseOutputWithContext(context.Context) LabVhdResponseOutput
}

type LabVhdResponseArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (LabVhdResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LabVhdResponse)(nil)).Elem()
}

func (i LabVhdResponseArgs) ToLabVhdResponseOutput() LabVhdResponseOutput {
	return i.ToLabVhdResponseOutputWithContext(context.Background())
}

func (i LabVhdResponseArgs) ToLabVhdResponseOutputWithContext(ctx context.Context) LabVhdResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabVhdResponseOutput)
}





type LabVhdResponseArrayInput interface {
	pulumi.Input

	ToLabVhdResponseArrayOutput() LabVhdResponseArrayOutput
	ToLabVhdResponseArrayOutputWithContext(context.Context) LabVhdResponseArrayOutput
}

type LabVhdResponseArray []LabVhdResponseInput

func (LabVhdResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LabVhdResponse)(nil)).Elem()
}

func (i LabVhdResponseArray) ToLabVhdResponseArrayOutput() LabVhdResponseArrayOutput {
	return i.ToLabVhdResponseArrayOutputWithContext(context.Background())
}

func (i LabVhdResponseArray) ToLabVhdResponseArrayOutputWithContext(ctx context.Context) LabVhdResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabVhdResponseArrayOutput)
}

type LabVhdResponseOutput struct{ *pulumi.OutputState }

func (LabVhdResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LabVhdResponse)(nil)).Elem()
}

func (o LabVhdResponseOutput) ToLabVhdResponseOutput() LabVhdResponseOutput {
	return o
}

func (o LabVhdResponseOutput) ToLabVhdResponseOutputWithContext(ctx context.Context) LabVhdResponseOutput {
	return o
}

func (o LabVhdResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVhdResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type LabVhdResponseArrayOutput struct{ *pulumi.OutputState }

func (LabVhdResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LabVhdResponse)(nil)).Elem()
}

func (o LabVhdResponseArrayOutput) ToLabVhdResponseArrayOutput() LabVhdResponseArrayOutput {
	return o
}

func (o LabVhdResponseArrayOutput) ToLabVhdResponseArrayOutputWithContext(ctx context.Context) LabVhdResponseArrayOutput {
	return o
}

func (o LabVhdResponseArrayOutput) Index(i pulumi.IntInput) LabVhdResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LabVhdResponse {
		return vs[0].([]LabVhdResponse)[vs[1].(int)]
	}).(LabVhdResponseOutput)
}

type LabVirtualMachine struct {
	ArtifactDeploymentStatus   *ArtifactDeploymentStatusProperties `pulumi:"artifactDeploymentStatus"`
	Artifacts                  []ArtifactInstallProperties         `pulumi:"artifacts"`
	ComputeId                  *string                             `pulumi:"computeId"`
	CreatedByUser              *string                             `pulumi:"createdByUser"`
	CreatedByUserId            *string                             `pulumi:"createdByUserId"`
	CustomImageId              *string                             `pulumi:"customImageId"`
	DisallowPublicIpAddress    *bool                               `pulumi:"disallowPublicIpAddress"`
	Fqdn                       *string                             `pulumi:"fqdn"`
	GalleryImageReference      *GalleryImageReference              `pulumi:"galleryImageReference"`
	Id                         *string                             `pulumi:"id"`
	IsAuthenticationWithSshKey *bool                               `pulumi:"isAuthenticationWithSshKey"`
	LabSubnetName              *string                             `pulumi:"labSubnetName"`
	LabVirtualNetworkId        *string                             `pulumi:"labVirtualNetworkId"`
	Location                   *string                             `pulumi:"location"`
	Name                       *string                             `pulumi:"name"`
	Notes                      *string                             `pulumi:"notes"`
	OsType                     *string                             `pulumi:"osType"`
	OwnerObjectId              *string                             `pulumi:"ownerObjectId"`
	Password                   *string                             `pulumi:"password"`
	ProvisioningState          *string                             `pulumi:"provisioningState"`
	Size                       *string                             `pulumi:"size"`
	SshKey                     *string                             `pulumi:"sshKey"`
	Tags                       map[string]string                   `pulumi:"tags"`
	Type                       *string                             `pulumi:"type"`
	UserName                   *string                             `pulumi:"userName"`
}





type LabVirtualMachineInput interface {
	pulumi.Input

	ToLabVirtualMachineOutput() LabVirtualMachineOutput
	ToLabVirtualMachineOutputWithContext(context.Context) LabVirtualMachineOutput
}

type LabVirtualMachineArgs struct {
	ArtifactDeploymentStatus   ArtifactDeploymentStatusPropertiesPtrInput `pulumi:"artifactDeploymentStatus"`
	Artifacts                  ArtifactInstallPropertiesArrayInput        `pulumi:"artifacts"`
	ComputeId                  pulumi.StringPtrInput                      `pulumi:"computeId"`
	CreatedByUser              pulumi.StringPtrInput                      `pulumi:"createdByUser"`
	CreatedByUserId            pulumi.StringPtrInput                      `pulumi:"createdByUserId"`
	CustomImageId              pulumi.StringPtrInput                      `pulumi:"customImageId"`
	DisallowPublicIpAddress    pulumi.BoolPtrInput                        `pulumi:"disallowPublicIpAddress"`
	Fqdn                       pulumi.StringPtrInput                      `pulumi:"fqdn"`
	GalleryImageReference      GalleryImageReferencePtrInput              `pulumi:"galleryImageReference"`
	Id                         pulumi.StringPtrInput                      `pulumi:"id"`
	IsAuthenticationWithSshKey pulumi.BoolPtrInput                        `pulumi:"isAuthenticationWithSshKey"`
	LabSubnetName              pulumi.StringPtrInput                      `pulumi:"labSubnetName"`
	LabVirtualNetworkId        pulumi.StringPtrInput                      `pulumi:"labVirtualNetworkId"`
	Location                   pulumi.StringPtrInput                      `pulumi:"location"`
	Name                       pulumi.StringPtrInput                      `pulumi:"name"`
	Notes                      pulumi.StringPtrInput                      `pulumi:"notes"`
	OsType                     pulumi.StringPtrInput                      `pulumi:"osType"`
	OwnerObjectId              pulumi.StringPtrInput                      `pulumi:"ownerObjectId"`
	Password                   pulumi.StringPtrInput                      `pulumi:"password"`
	ProvisioningState          pulumi.StringPtrInput                      `pulumi:"provisioningState"`
	Size                       pulumi.StringPtrInput                      `pulumi:"size"`
	SshKey                     pulumi.StringPtrInput                      `pulumi:"sshKey"`
	Tags                       pulumi.StringMapInput                      `pulumi:"tags"`
	Type                       pulumi.StringPtrInput                      `pulumi:"type"`
	UserName                   pulumi.StringPtrInput                      `pulumi:"userName"`
}

func (LabVirtualMachineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LabVirtualMachine)(nil)).Elem()
}

func (i LabVirtualMachineArgs) ToLabVirtualMachineOutput() LabVirtualMachineOutput {
	return i.ToLabVirtualMachineOutputWithContext(context.Background())
}

func (i LabVirtualMachineArgs) ToLabVirtualMachineOutputWithContext(ctx context.Context) LabVirtualMachineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabVirtualMachineOutput)
}

func (i LabVirtualMachineArgs) ToLabVirtualMachinePtrOutput() LabVirtualMachinePtrOutput {
	return i.ToLabVirtualMachinePtrOutputWithContext(context.Background())
}

func (i LabVirtualMachineArgs) ToLabVirtualMachinePtrOutputWithContext(ctx context.Context) LabVirtualMachinePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabVirtualMachineOutput).ToLabVirtualMachinePtrOutputWithContext(ctx)
}









type LabVirtualMachinePtrInput interface {
	pulumi.Input

	ToLabVirtualMachinePtrOutput() LabVirtualMachinePtrOutput
	ToLabVirtualMachinePtrOutputWithContext(context.Context) LabVirtualMachinePtrOutput
}

type labVirtualMachinePtrType LabVirtualMachineArgs

func LabVirtualMachinePtr(v *LabVirtualMachineArgs) LabVirtualMachinePtrInput {
	return (*labVirtualMachinePtrType)(v)
}

func (*labVirtualMachinePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LabVirtualMachine)(nil)).Elem()
}

func (i *labVirtualMachinePtrType) ToLabVirtualMachinePtrOutput() LabVirtualMachinePtrOutput {
	return i.ToLabVirtualMachinePtrOutputWithContext(context.Background())
}

func (i *labVirtualMachinePtrType) ToLabVirtualMachinePtrOutputWithContext(ctx context.Context) LabVirtualMachinePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabVirtualMachinePtrOutput)
}

type LabVirtualMachineOutput struct{ *pulumi.OutputState }

func (LabVirtualMachineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LabVirtualMachine)(nil)).Elem()
}

func (o LabVirtualMachineOutput) ToLabVirtualMachineOutput() LabVirtualMachineOutput {
	return o
}

func (o LabVirtualMachineOutput) ToLabVirtualMachineOutputWithContext(ctx context.Context) LabVirtualMachineOutput {
	return o
}

func (o LabVirtualMachineOutput) ToLabVirtualMachinePtrOutput() LabVirtualMachinePtrOutput {
	return o.ToLabVirtualMachinePtrOutputWithContext(context.Background())
}

func (o LabVirtualMachineOutput) ToLabVirtualMachinePtrOutputWithContext(ctx context.Context) LabVirtualMachinePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LabVirtualMachine) *LabVirtualMachine {
		return &v
	}).(LabVirtualMachinePtrOutput)
}

func (o LabVirtualMachineOutput) ArtifactDeploymentStatus() ArtifactDeploymentStatusPropertiesPtrOutput {
	return o.ApplyT(func(v LabVirtualMachine) *ArtifactDeploymentStatusProperties { return v.ArtifactDeploymentStatus }).(ArtifactDeploymentStatusPropertiesPtrOutput)
}

func (o LabVirtualMachineOutput) Artifacts() ArtifactInstallPropertiesArrayOutput {
	return o.ApplyT(func(v LabVirtualMachine) []ArtifactInstallProperties { return v.Artifacts }).(ArtifactInstallPropertiesArrayOutput)
}

func (o LabVirtualMachineOutput) ComputeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachine) *string { return v.ComputeId }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineOutput) CreatedByUser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachine) *string { return v.CreatedByUser }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineOutput) CreatedByUserId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachine) *string { return v.CreatedByUserId }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineOutput) CustomImageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachine) *string { return v.CustomImageId }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineOutput) DisallowPublicIpAddress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LabVirtualMachine) *bool { return v.DisallowPublicIpAddress }).(pulumi.BoolPtrOutput)
}

func (o LabVirtualMachineOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachine) *string { return v.Fqdn }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineOutput) GalleryImageReference() GalleryImageReferencePtrOutput {
	return o.ApplyT(func(v LabVirtualMachine) *GalleryImageReference { return v.GalleryImageReference }).(GalleryImageReferencePtrOutput)
}

func (o LabVirtualMachineOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachine) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineOutput) IsAuthenticationWithSshKey() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LabVirtualMachine) *bool { return v.IsAuthenticationWithSshKey }).(pulumi.BoolPtrOutput)
}

func (o LabVirtualMachineOutput) LabSubnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachine) *string { return v.LabSubnetName }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineOutput) LabVirtualNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachine) *string { return v.LabVirtualNetworkId }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachine) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachine) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachine) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachine) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineOutput) OwnerObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachine) *string { return v.OwnerObjectId }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachine) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachine) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachine) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineOutput) SshKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachine) *string { return v.SshKey }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LabVirtualMachine) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LabVirtualMachineOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachine) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachine) *string { return v.UserName }).(pulumi.StringPtrOutput)
}

type LabVirtualMachinePtrOutput struct{ *pulumi.OutputState }

func (LabVirtualMachinePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LabVirtualMachine)(nil)).Elem()
}

func (o LabVirtualMachinePtrOutput) ToLabVirtualMachinePtrOutput() LabVirtualMachinePtrOutput {
	return o
}

func (o LabVirtualMachinePtrOutput) ToLabVirtualMachinePtrOutputWithContext(ctx context.Context) LabVirtualMachinePtrOutput {
	return o
}

func (o LabVirtualMachinePtrOutput) Elem() LabVirtualMachineOutput {
	return o.ApplyT(func(v *LabVirtualMachine) LabVirtualMachine {
		if v != nil {
			return *v
		}
		var ret LabVirtualMachine
		return ret
	}).(LabVirtualMachineOutput)
}

func (o LabVirtualMachinePtrOutput) ArtifactDeploymentStatus() ArtifactDeploymentStatusPropertiesPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachine) *ArtifactDeploymentStatusProperties {
		if v == nil {
			return nil
		}
		return v.ArtifactDeploymentStatus
	}).(ArtifactDeploymentStatusPropertiesPtrOutput)
}

func (o LabVirtualMachinePtrOutput) Artifacts() ArtifactInstallPropertiesArrayOutput {
	return o.ApplyT(func(v *LabVirtualMachine) []ArtifactInstallProperties {
		if v == nil {
			return nil
		}
		return v.Artifacts
	}).(ArtifactInstallPropertiesArrayOutput)
}

func (o LabVirtualMachinePtrOutput) ComputeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachine) *string {
		if v == nil {
			return nil
		}
		return v.ComputeId
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachinePtrOutput) CreatedByUser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachine) *string {
		if v == nil {
			return nil
		}
		return v.CreatedByUser
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachinePtrOutput) CreatedByUserId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachine) *string {
		if v == nil {
			return nil
		}
		return v.CreatedByUserId
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachinePtrOutput) CustomImageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachine) *string {
		if v == nil {
			return nil
		}
		return v.CustomImageId
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachinePtrOutput) DisallowPublicIpAddress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachine) *bool {
		if v == nil {
			return nil
		}
		return v.DisallowPublicIpAddress
	}).(pulumi.BoolPtrOutput)
}

func (o LabVirtualMachinePtrOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachine) *string {
		if v == nil {
			return nil
		}
		return v.Fqdn
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachinePtrOutput) GalleryImageReference() GalleryImageReferencePtrOutput {
	return o.ApplyT(func(v *LabVirtualMachine) *GalleryImageReference {
		if v == nil {
			return nil
		}
		return v.GalleryImageReference
	}).(GalleryImageReferencePtrOutput)
}

func (o LabVirtualMachinePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachine) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachinePtrOutput) IsAuthenticationWithSshKey() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachine) *bool {
		if v == nil {
			return nil
		}
		return v.IsAuthenticationWithSshKey
	}).(pulumi.BoolPtrOutput)
}

func (o LabVirtualMachinePtrOutput) LabSubnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachine) *string {
		if v == nil {
			return nil
		}
		return v.LabSubnetName
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachinePtrOutput) LabVirtualNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachine) *string {
		if v == nil {
			return nil
		}
		return v.LabVirtualNetworkId
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachinePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachine) *string {
		if v == nil {
			return nil
		}
		return v.Location
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachinePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachine) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachinePtrOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachine) *string {
		if v == nil {
			return nil
		}
		return v.Notes
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachinePtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachine) *string {
		if v == nil {
			return nil
		}
		return v.OsType
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachinePtrOutput) OwnerObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachine) *string {
		if v == nil {
			return nil
		}
		return v.OwnerObjectId
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachinePtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachine) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachinePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachine) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachinePtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachine) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachinePtrOutput) SshKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachine) *string {
		if v == nil {
			return nil
		}
		return v.SshKey
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachinePtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LabVirtualMachine) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

func (o LabVirtualMachinePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachine) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachinePtrOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachine) *string {
		if v == nil {
			return nil
		}
		return v.UserName
	}).(pulumi.StringPtrOutput)
}

type LabVirtualMachineResponse struct {
	ArtifactDeploymentStatus   *ArtifactDeploymentStatusPropertiesResponse `pulumi:"artifactDeploymentStatus"`
	Artifacts                  []ArtifactInstallPropertiesResponse         `pulumi:"artifacts"`
	ComputeId                  *string                                     `pulumi:"computeId"`
	CreatedByUser              *string                                     `pulumi:"createdByUser"`
	CreatedByUserId            *string                                     `pulumi:"createdByUserId"`
	CustomImageId              *string                                     `pulumi:"customImageId"`
	DisallowPublicIpAddress    *bool                                       `pulumi:"disallowPublicIpAddress"`
	Fqdn                       *string                                     `pulumi:"fqdn"`
	GalleryImageReference      *GalleryImageReferenceResponse              `pulumi:"galleryImageReference"`
	Id                         *string                                     `pulumi:"id"`
	IsAuthenticationWithSshKey *bool                                       `pulumi:"isAuthenticationWithSshKey"`
	LabSubnetName              *string                                     `pulumi:"labSubnetName"`
	LabVirtualNetworkId        *string                                     `pulumi:"labVirtualNetworkId"`
	Location                   *string                                     `pulumi:"location"`
	Name                       *string                                     `pulumi:"name"`
	Notes                      *string                                     `pulumi:"notes"`
	OsType                     *string                                     `pulumi:"osType"`
	OwnerObjectId              *string                                     `pulumi:"ownerObjectId"`
	Password                   *string                                     `pulumi:"password"`
	ProvisioningState          *string                                     `pulumi:"provisioningState"`
	Size                       *string                                     `pulumi:"size"`
	SshKey                     *string                                     `pulumi:"sshKey"`
	Tags                       map[string]string                           `pulumi:"tags"`
	Type                       *string                                     `pulumi:"type"`
	UserName                   *string                                     `pulumi:"userName"`
}





type LabVirtualMachineResponseInput interface {
	pulumi.Input

	ToLabVirtualMachineResponseOutput() LabVirtualMachineResponseOutput
	ToLabVirtualMachineResponseOutputWithContext(context.Context) LabVirtualMachineResponseOutput
}

type LabVirtualMachineResponseArgs struct {
	ArtifactDeploymentStatus   ArtifactDeploymentStatusPropertiesResponsePtrInput `pulumi:"artifactDeploymentStatus"`
	Artifacts                  ArtifactInstallPropertiesResponseArrayInput        `pulumi:"artifacts"`
	ComputeId                  pulumi.StringPtrInput                              `pulumi:"computeId"`
	CreatedByUser              pulumi.StringPtrInput                              `pulumi:"createdByUser"`
	CreatedByUserId            pulumi.StringPtrInput                              `pulumi:"createdByUserId"`
	CustomImageId              pulumi.StringPtrInput                              `pulumi:"customImageId"`
	DisallowPublicIpAddress    pulumi.BoolPtrInput                                `pulumi:"disallowPublicIpAddress"`
	Fqdn                       pulumi.StringPtrInput                              `pulumi:"fqdn"`
	GalleryImageReference      GalleryImageReferenceResponsePtrInput              `pulumi:"galleryImageReference"`
	Id                         pulumi.StringPtrInput                              `pulumi:"id"`
	IsAuthenticationWithSshKey pulumi.BoolPtrInput                                `pulumi:"isAuthenticationWithSshKey"`
	LabSubnetName              pulumi.StringPtrInput                              `pulumi:"labSubnetName"`
	LabVirtualNetworkId        pulumi.StringPtrInput                              `pulumi:"labVirtualNetworkId"`
	Location                   pulumi.StringPtrInput                              `pulumi:"location"`
	Name                       pulumi.StringPtrInput                              `pulumi:"name"`
	Notes                      pulumi.StringPtrInput                              `pulumi:"notes"`
	OsType                     pulumi.StringPtrInput                              `pulumi:"osType"`
	OwnerObjectId              pulumi.StringPtrInput                              `pulumi:"ownerObjectId"`
	Password                   pulumi.StringPtrInput                              `pulumi:"password"`
	ProvisioningState          pulumi.StringPtrInput                              `pulumi:"provisioningState"`
	Size                       pulumi.StringPtrInput                              `pulumi:"size"`
	SshKey                     pulumi.StringPtrInput                              `pulumi:"sshKey"`
	Tags                       pulumi.StringMapInput                              `pulumi:"tags"`
	Type                       pulumi.StringPtrInput                              `pulumi:"type"`
	UserName                   pulumi.StringPtrInput                              `pulumi:"userName"`
}

func (LabVirtualMachineResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LabVirtualMachineResponse)(nil)).Elem()
}

func (i LabVirtualMachineResponseArgs) ToLabVirtualMachineResponseOutput() LabVirtualMachineResponseOutput {
	return i.ToLabVirtualMachineResponseOutputWithContext(context.Background())
}

func (i LabVirtualMachineResponseArgs) ToLabVirtualMachineResponseOutputWithContext(ctx context.Context) LabVirtualMachineResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabVirtualMachineResponseOutput)
}

func (i LabVirtualMachineResponseArgs) ToLabVirtualMachineResponsePtrOutput() LabVirtualMachineResponsePtrOutput {
	return i.ToLabVirtualMachineResponsePtrOutputWithContext(context.Background())
}

func (i LabVirtualMachineResponseArgs) ToLabVirtualMachineResponsePtrOutputWithContext(ctx context.Context) LabVirtualMachineResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabVirtualMachineResponseOutput).ToLabVirtualMachineResponsePtrOutputWithContext(ctx)
}









type LabVirtualMachineResponsePtrInput interface {
	pulumi.Input

	ToLabVirtualMachineResponsePtrOutput() LabVirtualMachineResponsePtrOutput
	ToLabVirtualMachineResponsePtrOutputWithContext(context.Context) LabVirtualMachineResponsePtrOutput
}

type labVirtualMachineResponsePtrType LabVirtualMachineResponseArgs

func LabVirtualMachineResponsePtr(v *LabVirtualMachineResponseArgs) LabVirtualMachineResponsePtrInput {
	return (*labVirtualMachineResponsePtrType)(v)
}

func (*labVirtualMachineResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LabVirtualMachineResponse)(nil)).Elem()
}

func (i *labVirtualMachineResponsePtrType) ToLabVirtualMachineResponsePtrOutput() LabVirtualMachineResponsePtrOutput {
	return i.ToLabVirtualMachineResponsePtrOutputWithContext(context.Background())
}

func (i *labVirtualMachineResponsePtrType) ToLabVirtualMachineResponsePtrOutputWithContext(ctx context.Context) LabVirtualMachineResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabVirtualMachineResponsePtrOutput)
}

type LabVirtualMachineResponseOutput struct{ *pulumi.OutputState }

func (LabVirtualMachineResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LabVirtualMachineResponse)(nil)).Elem()
}

func (o LabVirtualMachineResponseOutput) ToLabVirtualMachineResponseOutput() LabVirtualMachineResponseOutput {
	return o
}

func (o LabVirtualMachineResponseOutput) ToLabVirtualMachineResponseOutputWithContext(ctx context.Context) LabVirtualMachineResponseOutput {
	return o
}

func (o LabVirtualMachineResponseOutput) ToLabVirtualMachineResponsePtrOutput() LabVirtualMachineResponsePtrOutput {
	return o.ToLabVirtualMachineResponsePtrOutputWithContext(context.Background())
}

func (o LabVirtualMachineResponseOutput) ToLabVirtualMachineResponsePtrOutputWithContext(ctx context.Context) LabVirtualMachineResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LabVirtualMachineResponse) *LabVirtualMachineResponse {
		return &v
	}).(LabVirtualMachineResponsePtrOutput)
}

func (o LabVirtualMachineResponseOutput) ArtifactDeploymentStatus() ArtifactDeploymentStatusPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LabVirtualMachineResponse) *ArtifactDeploymentStatusPropertiesResponse {
		return v.ArtifactDeploymentStatus
	}).(ArtifactDeploymentStatusPropertiesResponsePtrOutput)
}

func (o LabVirtualMachineResponseOutput) Artifacts() ArtifactInstallPropertiesResponseArrayOutput {
	return o.ApplyT(func(v LabVirtualMachineResponse) []ArtifactInstallPropertiesResponse { return v.Artifacts }).(ArtifactInstallPropertiesResponseArrayOutput)
}

func (o LabVirtualMachineResponseOutput) ComputeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineResponse) *string { return v.ComputeId }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineResponseOutput) CreatedByUser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineResponse) *string { return v.CreatedByUser }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineResponseOutput) CreatedByUserId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineResponse) *string { return v.CreatedByUserId }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineResponseOutput) CustomImageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineResponse) *string { return v.CustomImageId }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineResponseOutput) DisallowPublicIpAddress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineResponse) *bool { return v.DisallowPublicIpAddress }).(pulumi.BoolPtrOutput)
}

func (o LabVirtualMachineResponseOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineResponse) *string { return v.Fqdn }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineResponseOutput) GalleryImageReference() GalleryImageReferenceResponsePtrOutput {
	return o.ApplyT(func(v LabVirtualMachineResponse) *GalleryImageReferenceResponse { return v.GalleryImageReference }).(GalleryImageReferenceResponsePtrOutput)
}

func (o LabVirtualMachineResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineResponseOutput) IsAuthenticationWithSshKey() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineResponse) *bool { return v.IsAuthenticationWithSshKey }).(pulumi.BoolPtrOutput)
}

func (o LabVirtualMachineResponseOutput) LabSubnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineResponse) *string { return v.LabSubnetName }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineResponseOutput) LabVirtualNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineResponse) *string { return v.LabVirtualNetworkId }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineResponseOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineResponse) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineResponseOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineResponse) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineResponseOutput) OwnerObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineResponse) *string { return v.OwnerObjectId }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineResponseOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineResponse) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineResponseOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineResponse) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineResponseOutput) SshKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineResponse) *string { return v.SshKey }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LabVirtualMachineResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LabVirtualMachineResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineResponseOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineResponse) *string { return v.UserName }).(pulumi.StringPtrOutput)
}

type LabVirtualMachineResponsePtrOutput struct{ *pulumi.OutputState }

func (LabVirtualMachineResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LabVirtualMachineResponse)(nil)).Elem()
}

func (o LabVirtualMachineResponsePtrOutput) ToLabVirtualMachineResponsePtrOutput() LabVirtualMachineResponsePtrOutput {
	return o
}

func (o LabVirtualMachineResponsePtrOutput) ToLabVirtualMachineResponsePtrOutputWithContext(ctx context.Context) LabVirtualMachineResponsePtrOutput {
	return o
}

func (o LabVirtualMachineResponsePtrOutput) Elem() LabVirtualMachineResponseOutput {
	return o.ApplyT(func(v *LabVirtualMachineResponse) LabVirtualMachineResponse {
		if v != nil {
			return *v
		}
		var ret LabVirtualMachineResponse
		return ret
	}).(LabVirtualMachineResponseOutput)
}

func (o LabVirtualMachineResponsePtrOutput) ArtifactDeploymentStatus() ArtifactDeploymentStatusPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineResponse) *ArtifactDeploymentStatusPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.ArtifactDeploymentStatus
	}).(ArtifactDeploymentStatusPropertiesResponsePtrOutput)
}

func (o LabVirtualMachineResponsePtrOutput) Artifacts() ArtifactInstallPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *LabVirtualMachineResponse) []ArtifactInstallPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.Artifacts
	}).(ArtifactInstallPropertiesResponseArrayOutput)
}

func (o LabVirtualMachineResponsePtrOutput) ComputeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineResponse) *string {
		if v == nil {
			return nil
		}
		return v.ComputeId
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineResponsePtrOutput) CreatedByUser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedByUser
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineResponsePtrOutput) CreatedByUserId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedByUserId
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineResponsePtrOutput) CustomImageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineResponse) *string {
		if v == nil {
			return nil
		}
		return v.CustomImageId
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineResponsePtrOutput) DisallowPublicIpAddress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineResponse) *bool {
		if v == nil {
			return nil
		}
		return v.DisallowPublicIpAddress
	}).(pulumi.BoolPtrOutput)
}

func (o LabVirtualMachineResponsePtrOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineResponse) *string {
		if v == nil {
			return nil
		}
		return v.Fqdn
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineResponsePtrOutput) GalleryImageReference() GalleryImageReferenceResponsePtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineResponse) *GalleryImageReferenceResponse {
		if v == nil {
			return nil
		}
		return v.GalleryImageReference
	}).(GalleryImageReferenceResponsePtrOutput)
}

func (o LabVirtualMachineResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineResponsePtrOutput) IsAuthenticationWithSshKey() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsAuthenticationWithSshKey
	}).(pulumi.BoolPtrOutput)
}

func (o LabVirtualMachineResponsePtrOutput) LabSubnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineResponse) *string {
		if v == nil {
			return nil
		}
		return v.LabSubnetName
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineResponsePtrOutput) LabVirtualNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineResponse) *string {
		if v == nil {
			return nil
		}
		return v.LabVirtualNetworkId
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineResponsePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineResponse) *string {
		if v == nil {
			return nil
		}
		return v.Location
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineResponsePtrOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineResponse) *string {
		if v == nil {
			return nil
		}
		return v.Notes
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineResponsePtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineResponse) *string {
		if v == nil {
			return nil
		}
		return v.OsType
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineResponsePtrOutput) OwnerObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineResponse) *string {
		if v == nil {
			return nil
		}
		return v.OwnerObjectId
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineResponsePtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineResponse) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineResponsePtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineResponse) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineResponsePtrOutput) SshKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineResponse) *string {
		if v == nil {
			return nil
		}
		return v.SshKey
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineResponsePtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LabVirtualMachineResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

func (o LabVirtualMachineResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineResponsePtrOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserName
	}).(pulumi.StringPtrOutput)
}

type LinuxOsInfo struct {
	LinuxOsState *string `pulumi:"linuxOsState"`
}





type LinuxOsInfoInput interface {
	pulumi.Input

	ToLinuxOsInfoOutput() LinuxOsInfoOutput
	ToLinuxOsInfoOutputWithContext(context.Context) LinuxOsInfoOutput
}

type LinuxOsInfoArgs struct {
	LinuxOsState pulumi.StringPtrInput `pulumi:"linuxOsState"`
}

func (LinuxOsInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxOsInfo)(nil)).Elem()
}

func (i LinuxOsInfoArgs) ToLinuxOsInfoOutput() LinuxOsInfoOutput {
	return i.ToLinuxOsInfoOutputWithContext(context.Background())
}

func (i LinuxOsInfoArgs) ToLinuxOsInfoOutputWithContext(ctx context.Context) LinuxOsInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxOsInfoOutput)
}

func (i LinuxOsInfoArgs) ToLinuxOsInfoPtrOutput() LinuxOsInfoPtrOutput {
	return i.ToLinuxOsInfoPtrOutputWithContext(context.Background())
}

func (i LinuxOsInfoArgs) ToLinuxOsInfoPtrOutputWithContext(ctx context.Context) LinuxOsInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxOsInfoOutput).ToLinuxOsInfoPtrOutputWithContext(ctx)
}









type LinuxOsInfoPtrInput interface {
	pulumi.Input

	ToLinuxOsInfoPtrOutput() LinuxOsInfoPtrOutput
	ToLinuxOsInfoPtrOutputWithContext(context.Context) LinuxOsInfoPtrOutput
}

type linuxOsInfoPtrType LinuxOsInfoArgs

func LinuxOsInfoPtr(v *LinuxOsInfoArgs) LinuxOsInfoPtrInput {
	return (*linuxOsInfoPtrType)(v)
}

func (*linuxOsInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxOsInfo)(nil)).Elem()
}

func (i *linuxOsInfoPtrType) ToLinuxOsInfoPtrOutput() LinuxOsInfoPtrOutput {
	return i.ToLinuxOsInfoPtrOutputWithContext(context.Background())
}

func (i *linuxOsInfoPtrType) ToLinuxOsInfoPtrOutputWithContext(ctx context.Context) LinuxOsInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxOsInfoPtrOutput)
}

type LinuxOsInfoOutput struct{ *pulumi.OutputState }

func (LinuxOsInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxOsInfo)(nil)).Elem()
}

func (o LinuxOsInfoOutput) ToLinuxOsInfoOutput() LinuxOsInfoOutput {
	return o
}

func (o LinuxOsInfoOutput) ToLinuxOsInfoOutputWithContext(ctx context.Context) LinuxOsInfoOutput {
	return o
}

func (o LinuxOsInfoOutput) ToLinuxOsInfoPtrOutput() LinuxOsInfoPtrOutput {
	return o.ToLinuxOsInfoPtrOutputWithContext(context.Background())
}

func (o LinuxOsInfoOutput) ToLinuxOsInfoPtrOutputWithContext(ctx context.Context) LinuxOsInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LinuxOsInfo) *LinuxOsInfo {
		return &v
	}).(LinuxOsInfoPtrOutput)
}

func (o LinuxOsInfoOutput) LinuxOsState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinuxOsInfo) *string { return v.LinuxOsState }).(pulumi.StringPtrOutput)
}

type LinuxOsInfoPtrOutput struct{ *pulumi.OutputState }

func (LinuxOsInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxOsInfo)(nil)).Elem()
}

func (o LinuxOsInfoPtrOutput) ToLinuxOsInfoPtrOutput() LinuxOsInfoPtrOutput {
	return o
}

func (o LinuxOsInfoPtrOutput) ToLinuxOsInfoPtrOutputWithContext(ctx context.Context) LinuxOsInfoPtrOutput {
	return o
}

func (o LinuxOsInfoPtrOutput) Elem() LinuxOsInfoOutput {
	return o.ApplyT(func(v *LinuxOsInfo) LinuxOsInfo {
		if v != nil {
			return *v
		}
		var ret LinuxOsInfo
		return ret
	}).(LinuxOsInfoOutput)
}

func (o LinuxOsInfoPtrOutput) LinuxOsState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinuxOsInfo) *string {
		if v == nil {
			return nil
		}
		return v.LinuxOsState
	}).(pulumi.StringPtrOutput)
}

type LinuxOsInfoResponse struct {
	LinuxOsState *string `pulumi:"linuxOsState"`
}





type LinuxOsInfoResponseInput interface {
	pulumi.Input

	ToLinuxOsInfoResponseOutput() LinuxOsInfoResponseOutput
	ToLinuxOsInfoResponseOutputWithContext(context.Context) LinuxOsInfoResponseOutput
}

type LinuxOsInfoResponseArgs struct {
	LinuxOsState pulumi.StringPtrInput `pulumi:"linuxOsState"`
}

func (LinuxOsInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxOsInfoResponse)(nil)).Elem()
}

func (i LinuxOsInfoResponseArgs) ToLinuxOsInfoResponseOutput() LinuxOsInfoResponseOutput {
	return i.ToLinuxOsInfoResponseOutputWithContext(context.Background())
}

func (i LinuxOsInfoResponseArgs) ToLinuxOsInfoResponseOutputWithContext(ctx context.Context) LinuxOsInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxOsInfoResponseOutput)
}

func (i LinuxOsInfoResponseArgs) ToLinuxOsInfoResponsePtrOutput() LinuxOsInfoResponsePtrOutput {
	return i.ToLinuxOsInfoResponsePtrOutputWithContext(context.Background())
}

func (i LinuxOsInfoResponseArgs) ToLinuxOsInfoResponsePtrOutputWithContext(ctx context.Context) LinuxOsInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxOsInfoResponseOutput).ToLinuxOsInfoResponsePtrOutputWithContext(ctx)
}









type LinuxOsInfoResponsePtrInput interface {
	pulumi.Input

	ToLinuxOsInfoResponsePtrOutput() LinuxOsInfoResponsePtrOutput
	ToLinuxOsInfoResponsePtrOutputWithContext(context.Context) LinuxOsInfoResponsePtrOutput
}

type linuxOsInfoResponsePtrType LinuxOsInfoResponseArgs

func LinuxOsInfoResponsePtr(v *LinuxOsInfoResponseArgs) LinuxOsInfoResponsePtrInput {
	return (*linuxOsInfoResponsePtrType)(v)
}

func (*linuxOsInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxOsInfoResponse)(nil)).Elem()
}

func (i *linuxOsInfoResponsePtrType) ToLinuxOsInfoResponsePtrOutput() LinuxOsInfoResponsePtrOutput {
	return i.ToLinuxOsInfoResponsePtrOutputWithContext(context.Background())
}

func (i *linuxOsInfoResponsePtrType) ToLinuxOsInfoResponsePtrOutputWithContext(ctx context.Context) LinuxOsInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxOsInfoResponsePtrOutput)
}

type LinuxOsInfoResponseOutput struct{ *pulumi.OutputState }

func (LinuxOsInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxOsInfoResponse)(nil)).Elem()
}

func (o LinuxOsInfoResponseOutput) ToLinuxOsInfoResponseOutput() LinuxOsInfoResponseOutput {
	return o
}

func (o LinuxOsInfoResponseOutput) ToLinuxOsInfoResponseOutputWithContext(ctx context.Context) LinuxOsInfoResponseOutput {
	return o
}

func (o LinuxOsInfoResponseOutput) ToLinuxOsInfoResponsePtrOutput() LinuxOsInfoResponsePtrOutput {
	return o.ToLinuxOsInfoResponsePtrOutputWithContext(context.Background())
}

func (o LinuxOsInfoResponseOutput) ToLinuxOsInfoResponsePtrOutputWithContext(ctx context.Context) LinuxOsInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LinuxOsInfoResponse) *LinuxOsInfoResponse {
		return &v
	}).(LinuxOsInfoResponsePtrOutput)
}

func (o LinuxOsInfoResponseOutput) LinuxOsState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinuxOsInfoResponse) *string { return v.LinuxOsState }).(pulumi.StringPtrOutput)
}

type LinuxOsInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (LinuxOsInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxOsInfoResponse)(nil)).Elem()
}

func (o LinuxOsInfoResponsePtrOutput) ToLinuxOsInfoResponsePtrOutput() LinuxOsInfoResponsePtrOutput {
	return o
}

func (o LinuxOsInfoResponsePtrOutput) ToLinuxOsInfoResponsePtrOutputWithContext(ctx context.Context) LinuxOsInfoResponsePtrOutput {
	return o
}

func (o LinuxOsInfoResponsePtrOutput) Elem() LinuxOsInfoResponseOutput {
	return o.ApplyT(func(v *LinuxOsInfoResponse) LinuxOsInfoResponse {
		if v != nil {
			return *v
		}
		var ret LinuxOsInfoResponse
		return ret
	}).(LinuxOsInfoResponseOutput)
}

func (o LinuxOsInfoResponsePtrOutput) LinuxOsState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinuxOsInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.LinuxOsState
	}).(pulumi.StringPtrOutput)
}

type Subnet struct {
	AllowPublicIp *string `pulumi:"allowPublicIp"`
	LabSubnetName *string `pulumi:"labSubnetName"`
	ResourceId    *string `pulumi:"resourceId"`
}





type SubnetInput interface {
	pulumi.Input

	ToSubnetOutput() SubnetOutput
	ToSubnetOutputWithContext(context.Context) SubnetOutput
}

type SubnetArgs struct {
	AllowPublicIp pulumi.StringPtrInput `pulumi:"allowPublicIp"`
	LabSubnetName pulumi.StringPtrInput `pulumi:"labSubnetName"`
	ResourceId    pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (SubnetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Subnet)(nil)).Elem()
}

func (i SubnetArgs) ToSubnetOutput() SubnetOutput {
	return i.ToSubnetOutputWithContext(context.Background())
}

func (i SubnetArgs) ToSubnetOutputWithContext(ctx context.Context) SubnetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetOutput)
}





type SubnetArrayInput interface {
	pulumi.Input

	ToSubnetArrayOutput() SubnetArrayOutput
	ToSubnetArrayOutputWithContext(context.Context) SubnetArrayOutput
}

type SubnetArray []SubnetInput

func (SubnetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Subnet)(nil)).Elem()
}

func (i SubnetArray) ToSubnetArrayOutput() SubnetArrayOutput {
	return i.ToSubnetArrayOutputWithContext(context.Background())
}

func (i SubnetArray) ToSubnetArrayOutputWithContext(ctx context.Context) SubnetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetArrayOutput)
}

type SubnetOutput struct{ *pulumi.OutputState }

func (SubnetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Subnet)(nil)).Elem()
}

func (o SubnetOutput) ToSubnetOutput() SubnetOutput {
	return o
}

func (o SubnetOutput) ToSubnetOutputWithContext(ctx context.Context) SubnetOutput {
	return o
}

func (o SubnetOutput) AllowPublicIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Subnet) *string { return v.AllowPublicIp }).(pulumi.StringPtrOutput)
}

func (o SubnetOutput) LabSubnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Subnet) *string { return v.LabSubnetName }).(pulumi.StringPtrOutput)
}

func (o SubnetOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Subnet) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type SubnetArrayOutput struct{ *pulumi.OutputState }

func (SubnetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Subnet)(nil)).Elem()
}

func (o SubnetArrayOutput) ToSubnetArrayOutput() SubnetArrayOutput {
	return o
}

func (o SubnetArrayOutput) ToSubnetArrayOutputWithContext(ctx context.Context) SubnetArrayOutput {
	return o
}

func (o SubnetArrayOutput) Index(i pulumi.IntInput) SubnetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Subnet {
		return vs[0].([]Subnet)[vs[1].(int)]
	}).(SubnetOutput)
}

type SubnetOverride struct {
	LabSubnetName                *string `pulumi:"labSubnetName"`
	ResourceId                   *string `pulumi:"resourceId"`
	UseInVmCreationPermission    *string `pulumi:"useInVmCreationPermission"`
	UsePublicIpAddressPermission *string `pulumi:"usePublicIpAddressPermission"`
}





type SubnetOverrideInput interface {
	pulumi.Input

	ToSubnetOverrideOutput() SubnetOverrideOutput
	ToSubnetOverrideOutputWithContext(context.Context) SubnetOverrideOutput
}

type SubnetOverrideArgs struct {
	LabSubnetName                pulumi.StringPtrInput `pulumi:"labSubnetName"`
	ResourceId                   pulumi.StringPtrInput `pulumi:"resourceId"`
	UseInVmCreationPermission    pulumi.StringPtrInput `pulumi:"useInVmCreationPermission"`
	UsePublicIpAddressPermission pulumi.StringPtrInput `pulumi:"usePublicIpAddressPermission"`
}

func (SubnetOverrideArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetOverride)(nil)).Elem()
}

func (i SubnetOverrideArgs) ToSubnetOverrideOutput() SubnetOverrideOutput {
	return i.ToSubnetOverrideOutputWithContext(context.Background())
}

func (i SubnetOverrideArgs) ToSubnetOverrideOutputWithContext(ctx context.Context) SubnetOverrideOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetOverrideOutput)
}





type SubnetOverrideArrayInput interface {
	pulumi.Input

	ToSubnetOverrideArrayOutput() SubnetOverrideArrayOutput
	ToSubnetOverrideArrayOutputWithContext(context.Context) SubnetOverrideArrayOutput
}

type SubnetOverrideArray []SubnetOverrideInput

func (SubnetOverrideArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubnetOverride)(nil)).Elem()
}

func (i SubnetOverrideArray) ToSubnetOverrideArrayOutput() SubnetOverrideArrayOutput {
	return i.ToSubnetOverrideArrayOutputWithContext(context.Background())
}

func (i SubnetOverrideArray) ToSubnetOverrideArrayOutputWithContext(ctx context.Context) SubnetOverrideArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetOverrideArrayOutput)
}

type SubnetOverrideOutput struct{ *pulumi.OutputState }

func (SubnetOverrideOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetOverride)(nil)).Elem()
}

func (o SubnetOverrideOutput) ToSubnetOverrideOutput() SubnetOverrideOutput {
	return o
}

func (o SubnetOverrideOutput) ToSubnetOverrideOutputWithContext(ctx context.Context) SubnetOverrideOutput {
	return o
}

func (o SubnetOverrideOutput) LabSubnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetOverride) *string { return v.LabSubnetName }).(pulumi.StringPtrOutput)
}

func (o SubnetOverrideOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetOverride) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o SubnetOverrideOutput) UseInVmCreationPermission() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetOverride) *string { return v.UseInVmCreationPermission }).(pulumi.StringPtrOutput)
}

func (o SubnetOverrideOutput) UsePublicIpAddressPermission() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetOverride) *string { return v.UsePublicIpAddressPermission }).(pulumi.StringPtrOutput)
}

type SubnetOverrideArrayOutput struct{ *pulumi.OutputState }

func (SubnetOverrideArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubnetOverride)(nil)).Elem()
}

func (o SubnetOverrideArrayOutput) ToSubnetOverrideArrayOutput() SubnetOverrideArrayOutput {
	return o
}

func (o SubnetOverrideArrayOutput) ToSubnetOverrideArrayOutputWithContext(ctx context.Context) SubnetOverrideArrayOutput {
	return o
}

func (o SubnetOverrideArrayOutput) Index(i pulumi.IntInput) SubnetOverrideOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubnetOverride {
		return vs[0].([]SubnetOverride)[vs[1].(int)]
	}).(SubnetOverrideOutput)
}

type SubnetOverrideResponse struct {
	LabSubnetName                *string `pulumi:"labSubnetName"`
	ResourceId                   *string `pulumi:"resourceId"`
	UseInVmCreationPermission    *string `pulumi:"useInVmCreationPermission"`
	UsePublicIpAddressPermission *string `pulumi:"usePublicIpAddressPermission"`
}





type SubnetOverrideResponseInput interface {
	pulumi.Input

	ToSubnetOverrideResponseOutput() SubnetOverrideResponseOutput
	ToSubnetOverrideResponseOutputWithContext(context.Context) SubnetOverrideResponseOutput
}

type SubnetOverrideResponseArgs struct {
	LabSubnetName                pulumi.StringPtrInput `pulumi:"labSubnetName"`
	ResourceId                   pulumi.StringPtrInput `pulumi:"resourceId"`
	UseInVmCreationPermission    pulumi.StringPtrInput `pulumi:"useInVmCreationPermission"`
	UsePublicIpAddressPermission pulumi.StringPtrInput `pulumi:"usePublicIpAddressPermission"`
}

func (SubnetOverrideResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetOverrideResponse)(nil)).Elem()
}

func (i SubnetOverrideResponseArgs) ToSubnetOverrideResponseOutput() SubnetOverrideResponseOutput {
	return i.ToSubnetOverrideResponseOutputWithContext(context.Background())
}

func (i SubnetOverrideResponseArgs) ToSubnetOverrideResponseOutputWithContext(ctx context.Context) SubnetOverrideResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetOverrideResponseOutput)
}





type SubnetOverrideResponseArrayInput interface {
	pulumi.Input

	ToSubnetOverrideResponseArrayOutput() SubnetOverrideResponseArrayOutput
	ToSubnetOverrideResponseArrayOutputWithContext(context.Context) SubnetOverrideResponseArrayOutput
}

type SubnetOverrideResponseArray []SubnetOverrideResponseInput

func (SubnetOverrideResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubnetOverrideResponse)(nil)).Elem()
}

func (i SubnetOverrideResponseArray) ToSubnetOverrideResponseArrayOutput() SubnetOverrideResponseArrayOutput {
	return i.ToSubnetOverrideResponseArrayOutputWithContext(context.Background())
}

func (i SubnetOverrideResponseArray) ToSubnetOverrideResponseArrayOutputWithContext(ctx context.Context) SubnetOverrideResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetOverrideResponseArrayOutput)
}

type SubnetOverrideResponseOutput struct{ *pulumi.OutputState }

func (SubnetOverrideResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetOverrideResponse)(nil)).Elem()
}

func (o SubnetOverrideResponseOutput) ToSubnetOverrideResponseOutput() SubnetOverrideResponseOutput {
	return o
}

func (o SubnetOverrideResponseOutput) ToSubnetOverrideResponseOutputWithContext(ctx context.Context) SubnetOverrideResponseOutput {
	return o
}

func (o SubnetOverrideResponseOutput) LabSubnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetOverrideResponse) *string { return v.LabSubnetName }).(pulumi.StringPtrOutput)
}

func (o SubnetOverrideResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetOverrideResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o SubnetOverrideResponseOutput) UseInVmCreationPermission() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetOverrideResponse) *string { return v.UseInVmCreationPermission }).(pulumi.StringPtrOutput)
}

func (o SubnetOverrideResponseOutput) UsePublicIpAddressPermission() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetOverrideResponse) *string { return v.UsePublicIpAddressPermission }).(pulumi.StringPtrOutput)
}

type SubnetOverrideResponseArrayOutput struct{ *pulumi.OutputState }

func (SubnetOverrideResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubnetOverrideResponse)(nil)).Elem()
}

func (o SubnetOverrideResponseArrayOutput) ToSubnetOverrideResponseArrayOutput() SubnetOverrideResponseArrayOutput {
	return o
}

func (o SubnetOverrideResponseArrayOutput) ToSubnetOverrideResponseArrayOutputWithContext(ctx context.Context) SubnetOverrideResponseArrayOutput {
	return o
}

func (o SubnetOverrideResponseArrayOutput) Index(i pulumi.IntInput) SubnetOverrideResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubnetOverrideResponse {
		return vs[0].([]SubnetOverrideResponse)[vs[1].(int)]
	}).(SubnetOverrideResponseOutput)
}

type SubnetResponse struct {
	AllowPublicIp *string `pulumi:"allowPublicIp"`
	LabSubnetName *string `pulumi:"labSubnetName"`
	ResourceId    *string `pulumi:"resourceId"`
}





type SubnetResponseInput interface {
	pulumi.Input

	ToSubnetResponseOutput() SubnetResponseOutput
	ToSubnetResponseOutputWithContext(context.Context) SubnetResponseOutput
}

type SubnetResponseArgs struct {
	AllowPublicIp pulumi.StringPtrInput `pulumi:"allowPublicIp"`
	LabSubnetName pulumi.StringPtrInput `pulumi:"labSubnetName"`
	ResourceId    pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (SubnetResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetResponse)(nil)).Elem()
}

func (i SubnetResponseArgs) ToSubnetResponseOutput() SubnetResponseOutput {
	return i.ToSubnetResponseOutputWithContext(context.Background())
}

func (i SubnetResponseArgs) ToSubnetResponseOutputWithContext(ctx context.Context) SubnetResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetResponseOutput)
}





type SubnetResponseArrayInput interface {
	pulumi.Input

	ToSubnetResponseArrayOutput() SubnetResponseArrayOutput
	ToSubnetResponseArrayOutputWithContext(context.Context) SubnetResponseArrayOutput
}

type SubnetResponseArray []SubnetResponseInput

func (SubnetResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubnetResponse)(nil)).Elem()
}

func (i SubnetResponseArray) ToSubnetResponseArrayOutput() SubnetResponseArrayOutput {
	return i.ToSubnetResponseArrayOutputWithContext(context.Background())
}

func (i SubnetResponseArray) ToSubnetResponseArrayOutputWithContext(ctx context.Context) SubnetResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetResponseArrayOutput)
}

type SubnetResponseOutput struct{ *pulumi.OutputState }

func (SubnetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetResponse)(nil)).Elem()
}

func (o SubnetResponseOutput) ToSubnetResponseOutput() SubnetResponseOutput {
	return o
}

func (o SubnetResponseOutput) ToSubnetResponseOutputWithContext(ctx context.Context) SubnetResponseOutput {
	return o
}

func (o SubnetResponseOutput) AllowPublicIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetResponse) *string { return v.AllowPublicIp }).(pulumi.StringPtrOutput)
}

func (o SubnetResponseOutput) LabSubnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetResponse) *string { return v.LabSubnetName }).(pulumi.StringPtrOutput)
}

func (o SubnetResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type SubnetResponseArrayOutput struct{ *pulumi.OutputState }

func (SubnetResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubnetResponse)(nil)).Elem()
}

func (o SubnetResponseArrayOutput) ToSubnetResponseArrayOutput() SubnetResponseArrayOutput {
	return o
}

func (o SubnetResponseArrayOutput) ToSubnetResponseArrayOutputWithContext(ctx context.Context) SubnetResponseArrayOutput {
	return o
}

func (o SubnetResponseArrayOutput) Index(i pulumi.IntInput) SubnetResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubnetResponse {
		return vs[0].([]SubnetResponse)[vs[1].(int)]
	}).(SubnetResponseOutput)
}

type WeekDetails struct {
	Time     *string  `pulumi:"time"`
	Weekdays []string `pulumi:"weekdays"`
}





type WeekDetailsInput interface {
	pulumi.Input

	ToWeekDetailsOutput() WeekDetailsOutput
	ToWeekDetailsOutputWithContext(context.Context) WeekDetailsOutput
}

type WeekDetailsArgs struct {
	Time     pulumi.StringPtrInput   `pulumi:"time"`
	Weekdays pulumi.StringArrayInput `pulumi:"weekdays"`
}

func (WeekDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WeekDetails)(nil)).Elem()
}

func (i WeekDetailsArgs) ToWeekDetailsOutput() WeekDetailsOutput {
	return i.ToWeekDetailsOutputWithContext(context.Background())
}

func (i WeekDetailsArgs) ToWeekDetailsOutputWithContext(ctx context.Context) WeekDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WeekDetailsOutput)
}

func (i WeekDetailsArgs) ToWeekDetailsPtrOutput() WeekDetailsPtrOutput {
	return i.ToWeekDetailsPtrOutputWithContext(context.Background())
}

func (i WeekDetailsArgs) ToWeekDetailsPtrOutputWithContext(ctx context.Context) WeekDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WeekDetailsOutput).ToWeekDetailsPtrOutputWithContext(ctx)
}









type WeekDetailsPtrInput interface {
	pulumi.Input

	ToWeekDetailsPtrOutput() WeekDetailsPtrOutput
	ToWeekDetailsPtrOutputWithContext(context.Context) WeekDetailsPtrOutput
}

type weekDetailsPtrType WeekDetailsArgs

func WeekDetailsPtr(v *WeekDetailsArgs) WeekDetailsPtrInput {
	return (*weekDetailsPtrType)(v)
}

func (*weekDetailsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WeekDetails)(nil)).Elem()
}

func (i *weekDetailsPtrType) ToWeekDetailsPtrOutput() WeekDetailsPtrOutput {
	return i.ToWeekDetailsPtrOutputWithContext(context.Background())
}

func (i *weekDetailsPtrType) ToWeekDetailsPtrOutputWithContext(ctx context.Context) WeekDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WeekDetailsPtrOutput)
}

type WeekDetailsOutput struct{ *pulumi.OutputState }

func (WeekDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WeekDetails)(nil)).Elem()
}

func (o WeekDetailsOutput) ToWeekDetailsOutput() WeekDetailsOutput {
	return o
}

func (o WeekDetailsOutput) ToWeekDetailsOutputWithContext(ctx context.Context) WeekDetailsOutput {
	return o
}

func (o WeekDetailsOutput) ToWeekDetailsPtrOutput() WeekDetailsPtrOutput {
	return o.ToWeekDetailsPtrOutputWithContext(context.Background())
}

func (o WeekDetailsOutput) ToWeekDetailsPtrOutputWithContext(ctx context.Context) WeekDetailsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WeekDetails) *WeekDetails {
		return &v
	}).(WeekDetailsPtrOutput)
}

func (o WeekDetailsOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WeekDetails) *string { return v.Time }).(pulumi.StringPtrOutput)
}

func (o WeekDetailsOutput) Weekdays() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WeekDetails) []string { return v.Weekdays }).(pulumi.StringArrayOutput)
}

type WeekDetailsPtrOutput struct{ *pulumi.OutputState }

func (WeekDetailsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WeekDetails)(nil)).Elem()
}

func (o WeekDetailsPtrOutput) ToWeekDetailsPtrOutput() WeekDetailsPtrOutput {
	return o
}

func (o WeekDetailsPtrOutput) ToWeekDetailsPtrOutputWithContext(ctx context.Context) WeekDetailsPtrOutput {
	return o
}

func (o WeekDetailsPtrOutput) Elem() WeekDetailsOutput {
	return o.ApplyT(func(v *WeekDetails) WeekDetails {
		if v != nil {
			return *v
		}
		var ret WeekDetails
		return ret
	}).(WeekDetailsOutput)
}

func (o WeekDetailsPtrOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WeekDetails) *string {
		if v == nil {
			return nil
		}
		return v.Time
	}).(pulumi.StringPtrOutput)
}

func (o WeekDetailsPtrOutput) Weekdays() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WeekDetails) []string {
		if v == nil {
			return nil
		}
		return v.Weekdays
	}).(pulumi.StringArrayOutput)
}

type WeekDetailsResponse struct {
	Time     *string  `pulumi:"time"`
	Weekdays []string `pulumi:"weekdays"`
}





type WeekDetailsResponseInput interface {
	pulumi.Input

	ToWeekDetailsResponseOutput() WeekDetailsResponseOutput
	ToWeekDetailsResponseOutputWithContext(context.Context) WeekDetailsResponseOutput
}

type WeekDetailsResponseArgs struct {
	Time     pulumi.StringPtrInput   `pulumi:"time"`
	Weekdays pulumi.StringArrayInput `pulumi:"weekdays"`
}

func (WeekDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WeekDetailsResponse)(nil)).Elem()
}

func (i WeekDetailsResponseArgs) ToWeekDetailsResponseOutput() WeekDetailsResponseOutput {
	return i.ToWeekDetailsResponseOutputWithContext(context.Background())
}

func (i WeekDetailsResponseArgs) ToWeekDetailsResponseOutputWithContext(ctx context.Context) WeekDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WeekDetailsResponseOutput)
}

func (i WeekDetailsResponseArgs) ToWeekDetailsResponsePtrOutput() WeekDetailsResponsePtrOutput {
	return i.ToWeekDetailsResponsePtrOutputWithContext(context.Background())
}

func (i WeekDetailsResponseArgs) ToWeekDetailsResponsePtrOutputWithContext(ctx context.Context) WeekDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WeekDetailsResponseOutput).ToWeekDetailsResponsePtrOutputWithContext(ctx)
}









type WeekDetailsResponsePtrInput interface {
	pulumi.Input

	ToWeekDetailsResponsePtrOutput() WeekDetailsResponsePtrOutput
	ToWeekDetailsResponsePtrOutputWithContext(context.Context) WeekDetailsResponsePtrOutput
}

type weekDetailsResponsePtrType WeekDetailsResponseArgs

func WeekDetailsResponsePtr(v *WeekDetailsResponseArgs) WeekDetailsResponsePtrInput {
	return (*weekDetailsResponsePtrType)(v)
}

func (*weekDetailsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WeekDetailsResponse)(nil)).Elem()
}

func (i *weekDetailsResponsePtrType) ToWeekDetailsResponsePtrOutput() WeekDetailsResponsePtrOutput {
	return i.ToWeekDetailsResponsePtrOutputWithContext(context.Background())
}

func (i *weekDetailsResponsePtrType) ToWeekDetailsResponsePtrOutputWithContext(ctx context.Context) WeekDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WeekDetailsResponsePtrOutput)
}

type WeekDetailsResponseOutput struct{ *pulumi.OutputState }

func (WeekDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WeekDetailsResponse)(nil)).Elem()
}

func (o WeekDetailsResponseOutput) ToWeekDetailsResponseOutput() WeekDetailsResponseOutput {
	return o
}

func (o WeekDetailsResponseOutput) ToWeekDetailsResponseOutputWithContext(ctx context.Context) WeekDetailsResponseOutput {
	return o
}

func (o WeekDetailsResponseOutput) ToWeekDetailsResponsePtrOutput() WeekDetailsResponsePtrOutput {
	return o.ToWeekDetailsResponsePtrOutputWithContext(context.Background())
}

func (o WeekDetailsResponseOutput) ToWeekDetailsResponsePtrOutputWithContext(ctx context.Context) WeekDetailsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WeekDetailsResponse) *WeekDetailsResponse {
		return &v
	}).(WeekDetailsResponsePtrOutput)
}

func (o WeekDetailsResponseOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WeekDetailsResponse) *string { return v.Time }).(pulumi.StringPtrOutput)
}

func (o WeekDetailsResponseOutput) Weekdays() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WeekDetailsResponse) []string { return v.Weekdays }).(pulumi.StringArrayOutput)
}

type WeekDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (WeekDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WeekDetailsResponse)(nil)).Elem()
}

func (o WeekDetailsResponsePtrOutput) ToWeekDetailsResponsePtrOutput() WeekDetailsResponsePtrOutput {
	return o
}

func (o WeekDetailsResponsePtrOutput) ToWeekDetailsResponsePtrOutputWithContext(ctx context.Context) WeekDetailsResponsePtrOutput {
	return o
}

func (o WeekDetailsResponsePtrOutput) Elem() WeekDetailsResponseOutput {
	return o.ApplyT(func(v *WeekDetailsResponse) WeekDetailsResponse {
		if v != nil {
			return *v
		}
		var ret WeekDetailsResponse
		return ret
	}).(WeekDetailsResponseOutput)
}

func (o WeekDetailsResponsePtrOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WeekDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Time
	}).(pulumi.StringPtrOutput)
}

func (o WeekDetailsResponsePtrOutput) Weekdays() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WeekDetailsResponse) []string {
		if v == nil {
			return nil
		}
		return v.Weekdays
	}).(pulumi.StringArrayOutput)
}

type WindowsOsInfo struct {
	WindowsOsState *string `pulumi:"windowsOsState"`
}





type WindowsOsInfoInput interface {
	pulumi.Input

	ToWindowsOsInfoOutput() WindowsOsInfoOutput
	ToWindowsOsInfoOutputWithContext(context.Context) WindowsOsInfoOutput
}

type WindowsOsInfoArgs struct {
	WindowsOsState pulumi.StringPtrInput `pulumi:"windowsOsState"`
}

func (WindowsOsInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WindowsOsInfo)(nil)).Elem()
}

func (i WindowsOsInfoArgs) ToWindowsOsInfoOutput() WindowsOsInfoOutput {
	return i.ToWindowsOsInfoOutputWithContext(context.Background())
}

func (i WindowsOsInfoArgs) ToWindowsOsInfoOutputWithContext(ctx context.Context) WindowsOsInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsOsInfoOutput)
}

func (i WindowsOsInfoArgs) ToWindowsOsInfoPtrOutput() WindowsOsInfoPtrOutput {
	return i.ToWindowsOsInfoPtrOutputWithContext(context.Background())
}

func (i WindowsOsInfoArgs) ToWindowsOsInfoPtrOutputWithContext(ctx context.Context) WindowsOsInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsOsInfoOutput).ToWindowsOsInfoPtrOutputWithContext(ctx)
}









type WindowsOsInfoPtrInput interface {
	pulumi.Input

	ToWindowsOsInfoPtrOutput() WindowsOsInfoPtrOutput
	ToWindowsOsInfoPtrOutputWithContext(context.Context) WindowsOsInfoPtrOutput
}

type windowsOsInfoPtrType WindowsOsInfoArgs

func WindowsOsInfoPtr(v *WindowsOsInfoArgs) WindowsOsInfoPtrInput {
	return (*windowsOsInfoPtrType)(v)
}

func (*windowsOsInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WindowsOsInfo)(nil)).Elem()
}

func (i *windowsOsInfoPtrType) ToWindowsOsInfoPtrOutput() WindowsOsInfoPtrOutput {
	return i.ToWindowsOsInfoPtrOutputWithContext(context.Background())
}

func (i *windowsOsInfoPtrType) ToWindowsOsInfoPtrOutputWithContext(ctx context.Context) WindowsOsInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsOsInfoPtrOutput)
}

type WindowsOsInfoOutput struct{ *pulumi.OutputState }

func (WindowsOsInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WindowsOsInfo)(nil)).Elem()
}

func (o WindowsOsInfoOutput) ToWindowsOsInfoOutput() WindowsOsInfoOutput {
	return o
}

func (o WindowsOsInfoOutput) ToWindowsOsInfoOutputWithContext(ctx context.Context) WindowsOsInfoOutput {
	return o
}

func (o WindowsOsInfoOutput) ToWindowsOsInfoPtrOutput() WindowsOsInfoPtrOutput {
	return o.ToWindowsOsInfoPtrOutputWithContext(context.Background())
}

func (o WindowsOsInfoOutput) ToWindowsOsInfoPtrOutputWithContext(ctx context.Context) WindowsOsInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WindowsOsInfo) *WindowsOsInfo {
		return &v
	}).(WindowsOsInfoPtrOutput)
}

func (o WindowsOsInfoOutput) WindowsOsState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WindowsOsInfo) *string { return v.WindowsOsState }).(pulumi.StringPtrOutput)
}

type WindowsOsInfoPtrOutput struct{ *pulumi.OutputState }

func (WindowsOsInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WindowsOsInfo)(nil)).Elem()
}

func (o WindowsOsInfoPtrOutput) ToWindowsOsInfoPtrOutput() WindowsOsInfoPtrOutput {
	return o
}

func (o WindowsOsInfoPtrOutput) ToWindowsOsInfoPtrOutputWithContext(ctx context.Context) WindowsOsInfoPtrOutput {
	return o
}

func (o WindowsOsInfoPtrOutput) Elem() WindowsOsInfoOutput {
	return o.ApplyT(func(v *WindowsOsInfo) WindowsOsInfo {
		if v != nil {
			return *v
		}
		var ret WindowsOsInfo
		return ret
	}).(WindowsOsInfoOutput)
}

func (o WindowsOsInfoPtrOutput) WindowsOsState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WindowsOsInfo) *string {
		if v == nil {
			return nil
		}
		return v.WindowsOsState
	}).(pulumi.StringPtrOutput)
}

type WindowsOsInfoResponse struct {
	WindowsOsState *string `pulumi:"windowsOsState"`
}





type WindowsOsInfoResponseInput interface {
	pulumi.Input

	ToWindowsOsInfoResponseOutput() WindowsOsInfoResponseOutput
	ToWindowsOsInfoResponseOutputWithContext(context.Context) WindowsOsInfoResponseOutput
}

type WindowsOsInfoResponseArgs struct {
	WindowsOsState pulumi.StringPtrInput `pulumi:"windowsOsState"`
}

func (WindowsOsInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WindowsOsInfoResponse)(nil)).Elem()
}

func (i WindowsOsInfoResponseArgs) ToWindowsOsInfoResponseOutput() WindowsOsInfoResponseOutput {
	return i.ToWindowsOsInfoResponseOutputWithContext(context.Background())
}

func (i WindowsOsInfoResponseArgs) ToWindowsOsInfoResponseOutputWithContext(ctx context.Context) WindowsOsInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsOsInfoResponseOutput)
}

func (i WindowsOsInfoResponseArgs) ToWindowsOsInfoResponsePtrOutput() WindowsOsInfoResponsePtrOutput {
	return i.ToWindowsOsInfoResponsePtrOutputWithContext(context.Background())
}

func (i WindowsOsInfoResponseArgs) ToWindowsOsInfoResponsePtrOutputWithContext(ctx context.Context) WindowsOsInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsOsInfoResponseOutput).ToWindowsOsInfoResponsePtrOutputWithContext(ctx)
}









type WindowsOsInfoResponsePtrInput interface {
	pulumi.Input

	ToWindowsOsInfoResponsePtrOutput() WindowsOsInfoResponsePtrOutput
	ToWindowsOsInfoResponsePtrOutputWithContext(context.Context) WindowsOsInfoResponsePtrOutput
}

type windowsOsInfoResponsePtrType WindowsOsInfoResponseArgs

func WindowsOsInfoResponsePtr(v *WindowsOsInfoResponseArgs) WindowsOsInfoResponsePtrInput {
	return (*windowsOsInfoResponsePtrType)(v)
}

func (*windowsOsInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WindowsOsInfoResponse)(nil)).Elem()
}

func (i *windowsOsInfoResponsePtrType) ToWindowsOsInfoResponsePtrOutput() WindowsOsInfoResponsePtrOutput {
	return i.ToWindowsOsInfoResponsePtrOutputWithContext(context.Background())
}

func (i *windowsOsInfoResponsePtrType) ToWindowsOsInfoResponsePtrOutputWithContext(ctx context.Context) WindowsOsInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsOsInfoResponsePtrOutput)
}

type WindowsOsInfoResponseOutput struct{ *pulumi.OutputState }

func (WindowsOsInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WindowsOsInfoResponse)(nil)).Elem()
}

func (o WindowsOsInfoResponseOutput) ToWindowsOsInfoResponseOutput() WindowsOsInfoResponseOutput {
	return o
}

func (o WindowsOsInfoResponseOutput) ToWindowsOsInfoResponseOutputWithContext(ctx context.Context) WindowsOsInfoResponseOutput {
	return o
}

func (o WindowsOsInfoResponseOutput) ToWindowsOsInfoResponsePtrOutput() WindowsOsInfoResponsePtrOutput {
	return o.ToWindowsOsInfoResponsePtrOutputWithContext(context.Background())
}

func (o WindowsOsInfoResponseOutput) ToWindowsOsInfoResponsePtrOutputWithContext(ctx context.Context) WindowsOsInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WindowsOsInfoResponse) *WindowsOsInfoResponse {
		return &v
	}).(WindowsOsInfoResponsePtrOutput)
}

func (o WindowsOsInfoResponseOutput) WindowsOsState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WindowsOsInfoResponse) *string { return v.WindowsOsState }).(pulumi.StringPtrOutput)
}

type WindowsOsInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (WindowsOsInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WindowsOsInfoResponse)(nil)).Elem()
}

func (o WindowsOsInfoResponsePtrOutput) ToWindowsOsInfoResponsePtrOutput() WindowsOsInfoResponsePtrOutput {
	return o
}

func (o WindowsOsInfoResponsePtrOutput) ToWindowsOsInfoResponsePtrOutputWithContext(ctx context.Context) WindowsOsInfoResponsePtrOutput {
	return o
}

func (o WindowsOsInfoResponsePtrOutput) Elem() WindowsOsInfoResponseOutput {
	return o.ApplyT(func(v *WindowsOsInfoResponse) WindowsOsInfoResponse {
		if v != nil {
			return *v
		}
		var ret WindowsOsInfoResponse
		return ret
	}).(WindowsOsInfoResponseOutput)
}

func (o WindowsOsInfoResponsePtrOutput) WindowsOsState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WindowsOsInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.WindowsOsState
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ArtifactDeploymentStatusPropertiesOutput{})
	pulumi.RegisterOutputType(ArtifactDeploymentStatusPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ArtifactDeploymentStatusPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ArtifactDeploymentStatusPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ArtifactInstallPropertiesOutput{})
	pulumi.RegisterOutputType(ArtifactInstallPropertiesArrayOutput{})
	pulumi.RegisterOutputType(ArtifactInstallPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ArtifactInstallPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(ArtifactParameterPropertiesOutput{})
	pulumi.RegisterOutputType(ArtifactParameterPropertiesArrayOutput{})
	pulumi.RegisterOutputType(ArtifactParameterPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ArtifactParameterPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(CustomImagePropertiesCustomOutput{})
	pulumi.RegisterOutputType(CustomImagePropertiesCustomPtrOutput{})
	pulumi.RegisterOutputType(CustomImagePropertiesCustomResponseOutput{})
	pulumi.RegisterOutputType(CustomImagePropertiesCustomResponsePtrOutput{})
	pulumi.RegisterOutputType(CustomImagePropertiesFromVmOutput{})
	pulumi.RegisterOutputType(CustomImagePropertiesFromVmPtrOutput{})
	pulumi.RegisterOutputType(CustomImagePropertiesFromVmResponseOutput{})
	pulumi.RegisterOutputType(CustomImagePropertiesFromVmResponsePtrOutput{})
	pulumi.RegisterOutputType(DayDetailsOutput{})
	pulumi.RegisterOutputType(DayDetailsPtrOutput{})
	pulumi.RegisterOutputType(DayDetailsResponseOutput{})
	pulumi.RegisterOutputType(DayDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(FormulaPropertiesFromVmOutput{})
	pulumi.RegisterOutputType(FormulaPropertiesFromVmPtrOutput{})
	pulumi.RegisterOutputType(FormulaPropertiesFromVmResponseOutput{})
	pulumi.RegisterOutputType(FormulaPropertiesFromVmResponsePtrOutput{})
	pulumi.RegisterOutputType(GalleryImageReferenceOutput{})
	pulumi.RegisterOutputType(GalleryImageReferencePtrOutput{})
	pulumi.RegisterOutputType(GalleryImageReferenceResponseOutput{})
	pulumi.RegisterOutputType(GalleryImageReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(HourDetailsOutput{})
	pulumi.RegisterOutputType(HourDetailsPtrOutput{})
	pulumi.RegisterOutputType(HourDetailsResponseOutput{})
	pulumi.RegisterOutputType(HourDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(LabVhdResponseOutput{})
	pulumi.RegisterOutputType(LabVhdResponseArrayOutput{})
	pulumi.RegisterOutputType(LabVirtualMachineOutput{})
	pulumi.RegisterOutputType(LabVirtualMachinePtrOutput{})
	pulumi.RegisterOutputType(LabVirtualMachineResponseOutput{})
	pulumi.RegisterOutputType(LabVirtualMachineResponsePtrOutput{})
	pulumi.RegisterOutputType(LinuxOsInfoOutput{})
	pulumi.RegisterOutputType(LinuxOsInfoPtrOutput{})
	pulumi.RegisterOutputType(LinuxOsInfoResponseOutput{})
	pulumi.RegisterOutputType(LinuxOsInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(SubnetOutput{})
	pulumi.RegisterOutputType(SubnetArrayOutput{})
	pulumi.RegisterOutputType(SubnetOverrideOutput{})
	pulumi.RegisterOutputType(SubnetOverrideArrayOutput{})
	pulumi.RegisterOutputType(SubnetOverrideResponseOutput{})
	pulumi.RegisterOutputType(SubnetOverrideResponseArrayOutput{})
	pulumi.RegisterOutputType(SubnetResponseOutput{})
	pulumi.RegisterOutputType(SubnetResponseArrayOutput{})
	pulumi.RegisterOutputType(WeekDetailsOutput{})
	pulumi.RegisterOutputType(WeekDetailsPtrOutput{})
	pulumi.RegisterOutputType(WeekDetailsResponseOutput{})
	pulumi.RegisterOutputType(WeekDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(WindowsOsInfoOutput{})
	pulumi.RegisterOutputType(WindowsOsInfoPtrOutput{})
	pulumi.RegisterOutputType(WindowsOsInfoResponseOutput{})
	pulumi.RegisterOutputType(WindowsOsInfoResponsePtrOutput{})
}
