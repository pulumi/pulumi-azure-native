


package v20180901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DatabaseProjectSummaryResponse struct {
	ExtendedSummary          map[string]string `pulumi:"extendedSummary"`
	InstanceType             string            `pulumi:"instanceType"`
	LastSummaryRefreshedTime *string           `pulumi:"lastSummaryRefreshedTime"`
	RefreshSummaryState      *string           `pulumi:"refreshSummaryState"`
}

type DatabasesSolutionSummaryResponse struct {
	DatabaseInstancesAssessedCount *int   `pulumi:"databaseInstancesAssessedCount"`
	DatabasesAssessedCount         *int   `pulumi:"databasesAssessedCount"`
	InstanceType                   string `pulumi:"instanceType"`
	MigrationReadyCount            *int   `pulumi:"migrationReadyCount"`
}

type MigrateProjectProperties struct {
	ProvisioningState *string  `pulumi:"provisioningState"`
	RegisteredTools   []string `pulumi:"registeredTools"`
}





type MigrateProjectPropertiesInput interface {
	pulumi.Input

	ToMigrateProjectPropertiesOutput() MigrateProjectPropertiesOutput
	ToMigrateProjectPropertiesOutputWithContext(context.Context) MigrateProjectPropertiesOutput
}

type MigrateProjectPropertiesArgs struct {
	ProvisioningState pulumi.StringPtrInput   `pulumi:"provisioningState"`
	RegisteredTools   pulumi.StringArrayInput `pulumi:"registeredTools"`
}

func (MigrateProjectPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MigrateProjectProperties)(nil)).Elem()
}

func (i MigrateProjectPropertiesArgs) ToMigrateProjectPropertiesOutput() MigrateProjectPropertiesOutput {
	return i.ToMigrateProjectPropertiesOutputWithContext(context.Background())
}

func (i MigrateProjectPropertiesArgs) ToMigrateProjectPropertiesOutputWithContext(ctx context.Context) MigrateProjectPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrateProjectPropertiesOutput)
}

func (i MigrateProjectPropertiesArgs) ToMigrateProjectPropertiesPtrOutput() MigrateProjectPropertiesPtrOutput {
	return i.ToMigrateProjectPropertiesPtrOutputWithContext(context.Background())
}

func (i MigrateProjectPropertiesArgs) ToMigrateProjectPropertiesPtrOutputWithContext(ctx context.Context) MigrateProjectPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrateProjectPropertiesOutput).ToMigrateProjectPropertiesPtrOutputWithContext(ctx)
}









type MigrateProjectPropertiesPtrInput interface {
	pulumi.Input

	ToMigrateProjectPropertiesPtrOutput() MigrateProjectPropertiesPtrOutput
	ToMigrateProjectPropertiesPtrOutputWithContext(context.Context) MigrateProjectPropertiesPtrOutput
}

type migrateProjectPropertiesPtrType MigrateProjectPropertiesArgs

func MigrateProjectPropertiesPtr(v *MigrateProjectPropertiesArgs) MigrateProjectPropertiesPtrInput {
	return (*migrateProjectPropertiesPtrType)(v)
}

func (*migrateProjectPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrateProjectProperties)(nil)).Elem()
}

func (i *migrateProjectPropertiesPtrType) ToMigrateProjectPropertiesPtrOutput() MigrateProjectPropertiesPtrOutput {
	return i.ToMigrateProjectPropertiesPtrOutputWithContext(context.Background())
}

func (i *migrateProjectPropertiesPtrType) ToMigrateProjectPropertiesPtrOutputWithContext(ctx context.Context) MigrateProjectPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrateProjectPropertiesPtrOutput)
}

type MigrateProjectPropertiesOutput struct{ *pulumi.OutputState }

func (MigrateProjectPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MigrateProjectProperties)(nil)).Elem()
}

func (o MigrateProjectPropertiesOutput) ToMigrateProjectPropertiesOutput() MigrateProjectPropertiesOutput {
	return o
}

func (o MigrateProjectPropertiesOutput) ToMigrateProjectPropertiesOutputWithContext(ctx context.Context) MigrateProjectPropertiesOutput {
	return o
}

func (o MigrateProjectPropertiesOutput) ToMigrateProjectPropertiesPtrOutput() MigrateProjectPropertiesPtrOutput {
	return o.ToMigrateProjectPropertiesPtrOutputWithContext(context.Background())
}

func (o MigrateProjectPropertiesOutput) ToMigrateProjectPropertiesPtrOutputWithContext(ctx context.Context) MigrateProjectPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MigrateProjectProperties) *MigrateProjectProperties {
		return &v
	}).(MigrateProjectPropertiesPtrOutput)
}

func (o MigrateProjectPropertiesOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MigrateProjectProperties) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o MigrateProjectPropertiesOutput) RegisteredTools() pulumi.StringArrayOutput {
	return o.ApplyT(func(v MigrateProjectProperties) []string { return v.RegisteredTools }).(pulumi.StringArrayOutput)
}

type MigrateProjectPropertiesPtrOutput struct{ *pulumi.OutputState }

func (MigrateProjectPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrateProjectProperties)(nil)).Elem()
}

func (o MigrateProjectPropertiesPtrOutput) ToMigrateProjectPropertiesPtrOutput() MigrateProjectPropertiesPtrOutput {
	return o
}

func (o MigrateProjectPropertiesPtrOutput) ToMigrateProjectPropertiesPtrOutputWithContext(ctx context.Context) MigrateProjectPropertiesPtrOutput {
	return o
}

func (o MigrateProjectPropertiesPtrOutput) Elem() MigrateProjectPropertiesOutput {
	return o.ApplyT(func(v *MigrateProjectProperties) MigrateProjectProperties {
		if v != nil {
			return *v
		}
		var ret MigrateProjectProperties
		return ret
	}).(MigrateProjectPropertiesOutput)
}

func (o MigrateProjectPropertiesPtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MigrateProjectProperties) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o MigrateProjectPropertiesPtrOutput) RegisteredTools() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MigrateProjectProperties) []string {
		if v == nil {
			return nil
		}
		return v.RegisteredTools
	}).(pulumi.StringArrayOutput)
}

type MigrateProjectPropertiesResponse struct {
	LastSummaryRefreshedTime string                 `pulumi:"lastSummaryRefreshedTime"`
	ProvisioningState        *string                `pulumi:"provisioningState"`
	RefreshSummaryState      string                 `pulumi:"refreshSummaryState"`
	RegisteredTools          []string               `pulumi:"registeredTools"`
	Summary                  map[string]interface{} `pulumi:"summary"`
}

type MigrateProjectPropertiesResponseOutput struct{ *pulumi.OutputState }

func (MigrateProjectPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MigrateProjectPropertiesResponse)(nil)).Elem()
}

func (o MigrateProjectPropertiesResponseOutput) ToMigrateProjectPropertiesResponseOutput() MigrateProjectPropertiesResponseOutput {
	return o
}

func (o MigrateProjectPropertiesResponseOutput) ToMigrateProjectPropertiesResponseOutputWithContext(ctx context.Context) MigrateProjectPropertiesResponseOutput {
	return o
}

func (o MigrateProjectPropertiesResponseOutput) LastSummaryRefreshedTime() pulumi.StringOutput {
	return o.ApplyT(func(v MigrateProjectPropertiesResponse) string { return v.LastSummaryRefreshedTime }).(pulumi.StringOutput)
}

func (o MigrateProjectPropertiesResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MigrateProjectPropertiesResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o MigrateProjectPropertiesResponseOutput) RefreshSummaryState() pulumi.StringOutput {
	return o.ApplyT(func(v MigrateProjectPropertiesResponse) string { return v.RefreshSummaryState }).(pulumi.StringOutput)
}

func (o MigrateProjectPropertiesResponseOutput) RegisteredTools() pulumi.StringArrayOutput {
	return o.ApplyT(func(v MigrateProjectPropertiesResponse) []string { return v.RegisteredTools }).(pulumi.StringArrayOutput)
}

func (o MigrateProjectPropertiesResponseOutput) Summary() pulumi.MapOutput {
	return o.ApplyT(func(v MigrateProjectPropertiesResponse) map[string]interface{} { return v.Summary }).(pulumi.MapOutput)
}

type MigrateProjectResponseTags struct {
	AdditionalProperties *string `pulumi:"additionalProperties"`
}

type MigrateProjectResponseTagsOutput struct{ *pulumi.OutputState }

func (MigrateProjectResponseTagsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MigrateProjectResponseTags)(nil)).Elem()
}

func (o MigrateProjectResponseTagsOutput) ToMigrateProjectResponseTagsOutput() MigrateProjectResponseTagsOutput {
	return o
}

func (o MigrateProjectResponseTagsOutput) ToMigrateProjectResponseTagsOutputWithContext(ctx context.Context) MigrateProjectResponseTagsOutput {
	return o
}

func (o MigrateProjectResponseTagsOutput) AdditionalProperties() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MigrateProjectResponseTags) *string { return v.AdditionalProperties }).(pulumi.StringPtrOutput)
}

type MigrateProjectResponseTagsPtrOutput struct{ *pulumi.OutputState }

func (MigrateProjectResponseTagsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrateProjectResponseTags)(nil)).Elem()
}

func (o MigrateProjectResponseTagsPtrOutput) ToMigrateProjectResponseTagsPtrOutput() MigrateProjectResponseTagsPtrOutput {
	return o
}

func (o MigrateProjectResponseTagsPtrOutput) ToMigrateProjectResponseTagsPtrOutputWithContext(ctx context.Context) MigrateProjectResponseTagsPtrOutput {
	return o
}

func (o MigrateProjectResponseTagsPtrOutput) Elem() MigrateProjectResponseTagsOutput {
	return o.ApplyT(func(v *MigrateProjectResponseTags) MigrateProjectResponseTags {
		if v != nil {
			return *v
		}
		var ret MigrateProjectResponseTags
		return ret
	}).(MigrateProjectResponseTagsOutput)
}

func (o MigrateProjectResponseTagsPtrOutput) AdditionalProperties() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MigrateProjectResponseTags) *string {
		if v == nil {
			return nil
		}
		return v.AdditionalProperties
	}).(pulumi.StringPtrOutput)
}

type MigrateProjectTags struct {
	AdditionalProperties *string `pulumi:"additionalProperties"`
}





type MigrateProjectTagsInput interface {
	pulumi.Input

	ToMigrateProjectTagsOutput() MigrateProjectTagsOutput
	ToMigrateProjectTagsOutputWithContext(context.Context) MigrateProjectTagsOutput
}

type MigrateProjectTagsArgs struct {
	AdditionalProperties pulumi.StringPtrInput `pulumi:"additionalProperties"`
}

func (MigrateProjectTagsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MigrateProjectTags)(nil)).Elem()
}

func (i MigrateProjectTagsArgs) ToMigrateProjectTagsOutput() MigrateProjectTagsOutput {
	return i.ToMigrateProjectTagsOutputWithContext(context.Background())
}

func (i MigrateProjectTagsArgs) ToMigrateProjectTagsOutputWithContext(ctx context.Context) MigrateProjectTagsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrateProjectTagsOutput)
}

func (i MigrateProjectTagsArgs) ToMigrateProjectTagsPtrOutput() MigrateProjectTagsPtrOutput {
	return i.ToMigrateProjectTagsPtrOutputWithContext(context.Background())
}

func (i MigrateProjectTagsArgs) ToMigrateProjectTagsPtrOutputWithContext(ctx context.Context) MigrateProjectTagsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrateProjectTagsOutput).ToMigrateProjectTagsPtrOutputWithContext(ctx)
}









type MigrateProjectTagsPtrInput interface {
	pulumi.Input

	ToMigrateProjectTagsPtrOutput() MigrateProjectTagsPtrOutput
	ToMigrateProjectTagsPtrOutputWithContext(context.Context) MigrateProjectTagsPtrOutput
}

type migrateProjectTagsPtrType MigrateProjectTagsArgs

func MigrateProjectTagsPtr(v *MigrateProjectTagsArgs) MigrateProjectTagsPtrInput {
	return (*migrateProjectTagsPtrType)(v)
}

func (*migrateProjectTagsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrateProjectTags)(nil)).Elem()
}

func (i *migrateProjectTagsPtrType) ToMigrateProjectTagsPtrOutput() MigrateProjectTagsPtrOutput {
	return i.ToMigrateProjectTagsPtrOutputWithContext(context.Background())
}

func (i *migrateProjectTagsPtrType) ToMigrateProjectTagsPtrOutputWithContext(ctx context.Context) MigrateProjectTagsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrateProjectTagsPtrOutput)
}

type MigrateProjectTagsOutput struct{ *pulumi.OutputState }

func (MigrateProjectTagsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MigrateProjectTags)(nil)).Elem()
}

func (o MigrateProjectTagsOutput) ToMigrateProjectTagsOutput() MigrateProjectTagsOutput {
	return o
}

func (o MigrateProjectTagsOutput) ToMigrateProjectTagsOutputWithContext(ctx context.Context) MigrateProjectTagsOutput {
	return o
}

func (o MigrateProjectTagsOutput) ToMigrateProjectTagsPtrOutput() MigrateProjectTagsPtrOutput {
	return o.ToMigrateProjectTagsPtrOutputWithContext(context.Background())
}

func (o MigrateProjectTagsOutput) ToMigrateProjectTagsPtrOutputWithContext(ctx context.Context) MigrateProjectTagsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MigrateProjectTags) *MigrateProjectTags {
		return &v
	}).(MigrateProjectTagsPtrOutput)
}

func (o MigrateProjectTagsOutput) AdditionalProperties() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MigrateProjectTags) *string { return v.AdditionalProperties }).(pulumi.StringPtrOutput)
}

type MigrateProjectTagsPtrOutput struct{ *pulumi.OutputState }

func (MigrateProjectTagsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrateProjectTags)(nil)).Elem()
}

func (o MigrateProjectTagsPtrOutput) ToMigrateProjectTagsPtrOutput() MigrateProjectTagsPtrOutput {
	return o
}

func (o MigrateProjectTagsPtrOutput) ToMigrateProjectTagsPtrOutputWithContext(ctx context.Context) MigrateProjectTagsPtrOutput {
	return o
}

func (o MigrateProjectTagsPtrOutput) Elem() MigrateProjectTagsOutput {
	return o.ApplyT(func(v *MigrateProjectTags) MigrateProjectTags {
		if v != nil {
			return *v
		}
		var ret MigrateProjectTags
		return ret
	}).(MigrateProjectTagsOutput)
}

func (o MigrateProjectTagsPtrOutput) AdditionalProperties() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MigrateProjectTags) *string {
		if v == nil {
			return nil
		}
		return v.AdditionalProperties
	}).(pulumi.StringPtrOutput)
}

type ServersProjectSummaryResponse struct {
	AssessedCount            *int              `pulumi:"assessedCount"`
	DiscoveredCount          *int              `pulumi:"discoveredCount"`
	ExtendedSummary          map[string]string `pulumi:"extendedSummary"`
	InstanceType             string            `pulumi:"instanceType"`
	LastSummaryRefreshedTime *string           `pulumi:"lastSummaryRefreshedTime"`
	MigratedCount            *int              `pulumi:"migratedCount"`
	RefreshSummaryState      *string           `pulumi:"refreshSummaryState"`
	ReplicatingCount         *int              `pulumi:"replicatingCount"`
	TestMigratedCount        *int              `pulumi:"testMigratedCount"`
}

type ServersSolutionSummaryResponse struct {
	AssessedCount     *int   `pulumi:"assessedCount"`
	DiscoveredCount   *int   `pulumi:"discoveredCount"`
	InstanceType      string `pulumi:"instanceType"`
	MigratedCount     *int   `pulumi:"migratedCount"`
	ReplicatingCount  *int   `pulumi:"replicatingCount"`
	TestMigratedCount *int   `pulumi:"testMigratedCount"`
}

type SolutionDetails struct {
	AssessmentCount *int              `pulumi:"assessmentCount"`
	ExtendedDetails map[string]string `pulumi:"extendedDetails"`
	GroupCount      *int              `pulumi:"groupCount"`
}





type SolutionDetailsInput interface {
	pulumi.Input

	ToSolutionDetailsOutput() SolutionDetailsOutput
	ToSolutionDetailsOutputWithContext(context.Context) SolutionDetailsOutput
}

type SolutionDetailsArgs struct {
	AssessmentCount pulumi.IntPtrInput    `pulumi:"assessmentCount"`
	ExtendedDetails pulumi.StringMapInput `pulumi:"extendedDetails"`
	GroupCount      pulumi.IntPtrInput    `pulumi:"groupCount"`
}

func (SolutionDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SolutionDetails)(nil)).Elem()
}

func (i SolutionDetailsArgs) ToSolutionDetailsOutput() SolutionDetailsOutput {
	return i.ToSolutionDetailsOutputWithContext(context.Background())
}

func (i SolutionDetailsArgs) ToSolutionDetailsOutputWithContext(ctx context.Context) SolutionDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SolutionDetailsOutput)
}

func (i SolutionDetailsArgs) ToSolutionDetailsPtrOutput() SolutionDetailsPtrOutput {
	return i.ToSolutionDetailsPtrOutputWithContext(context.Background())
}

func (i SolutionDetailsArgs) ToSolutionDetailsPtrOutputWithContext(ctx context.Context) SolutionDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SolutionDetailsOutput).ToSolutionDetailsPtrOutputWithContext(ctx)
}









type SolutionDetailsPtrInput interface {
	pulumi.Input

	ToSolutionDetailsPtrOutput() SolutionDetailsPtrOutput
	ToSolutionDetailsPtrOutputWithContext(context.Context) SolutionDetailsPtrOutput
}

type solutionDetailsPtrType SolutionDetailsArgs

func SolutionDetailsPtr(v *SolutionDetailsArgs) SolutionDetailsPtrInput {
	return (*solutionDetailsPtrType)(v)
}

func (*solutionDetailsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SolutionDetails)(nil)).Elem()
}

func (i *solutionDetailsPtrType) ToSolutionDetailsPtrOutput() SolutionDetailsPtrOutput {
	return i.ToSolutionDetailsPtrOutputWithContext(context.Background())
}

func (i *solutionDetailsPtrType) ToSolutionDetailsPtrOutputWithContext(ctx context.Context) SolutionDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SolutionDetailsPtrOutput)
}

type SolutionDetailsOutput struct{ *pulumi.OutputState }

func (SolutionDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SolutionDetails)(nil)).Elem()
}

func (o SolutionDetailsOutput) ToSolutionDetailsOutput() SolutionDetailsOutput {
	return o
}

func (o SolutionDetailsOutput) ToSolutionDetailsOutputWithContext(ctx context.Context) SolutionDetailsOutput {
	return o
}

func (o SolutionDetailsOutput) ToSolutionDetailsPtrOutput() SolutionDetailsPtrOutput {
	return o.ToSolutionDetailsPtrOutputWithContext(context.Background())
}

func (o SolutionDetailsOutput) ToSolutionDetailsPtrOutputWithContext(ctx context.Context) SolutionDetailsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SolutionDetails) *SolutionDetails {
		return &v
	}).(SolutionDetailsPtrOutput)
}

func (o SolutionDetailsOutput) AssessmentCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SolutionDetails) *int { return v.AssessmentCount }).(pulumi.IntPtrOutput)
}

func (o SolutionDetailsOutput) ExtendedDetails() pulumi.StringMapOutput {
	return o.ApplyT(func(v SolutionDetails) map[string]string { return v.ExtendedDetails }).(pulumi.StringMapOutput)
}

func (o SolutionDetailsOutput) GroupCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SolutionDetails) *int { return v.GroupCount }).(pulumi.IntPtrOutput)
}

type SolutionDetailsPtrOutput struct{ *pulumi.OutputState }

func (SolutionDetailsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SolutionDetails)(nil)).Elem()
}

func (o SolutionDetailsPtrOutput) ToSolutionDetailsPtrOutput() SolutionDetailsPtrOutput {
	return o
}

func (o SolutionDetailsPtrOutput) ToSolutionDetailsPtrOutputWithContext(ctx context.Context) SolutionDetailsPtrOutput {
	return o
}

func (o SolutionDetailsPtrOutput) Elem() SolutionDetailsOutput {
	return o.ApplyT(func(v *SolutionDetails) SolutionDetails {
		if v != nil {
			return *v
		}
		var ret SolutionDetails
		return ret
	}).(SolutionDetailsOutput)
}

func (o SolutionDetailsPtrOutput) AssessmentCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SolutionDetails) *int {
		if v == nil {
			return nil
		}
		return v.AssessmentCount
	}).(pulumi.IntPtrOutput)
}

func (o SolutionDetailsPtrOutput) ExtendedDetails() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SolutionDetails) map[string]string {
		if v == nil {
			return nil
		}
		return v.ExtendedDetails
	}).(pulumi.StringMapOutput)
}

func (o SolutionDetailsPtrOutput) GroupCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SolutionDetails) *int {
		if v == nil {
			return nil
		}
		return v.GroupCount
	}).(pulumi.IntPtrOutput)
}

type SolutionDetailsResponse struct {
	AssessmentCount *int              `pulumi:"assessmentCount"`
	ExtendedDetails map[string]string `pulumi:"extendedDetails"`
	GroupCount      *int              `pulumi:"groupCount"`
}

type SolutionDetailsResponseOutput struct{ *pulumi.OutputState }

func (SolutionDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SolutionDetailsResponse)(nil)).Elem()
}

func (o SolutionDetailsResponseOutput) ToSolutionDetailsResponseOutput() SolutionDetailsResponseOutput {
	return o
}

func (o SolutionDetailsResponseOutput) ToSolutionDetailsResponseOutputWithContext(ctx context.Context) SolutionDetailsResponseOutput {
	return o
}

func (o SolutionDetailsResponseOutput) AssessmentCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SolutionDetailsResponse) *int { return v.AssessmentCount }).(pulumi.IntPtrOutput)
}

func (o SolutionDetailsResponseOutput) ExtendedDetails() pulumi.StringMapOutput {
	return o.ApplyT(func(v SolutionDetailsResponse) map[string]string { return v.ExtendedDetails }).(pulumi.StringMapOutput)
}

func (o SolutionDetailsResponseOutput) GroupCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SolutionDetailsResponse) *int { return v.GroupCount }).(pulumi.IntPtrOutput)
}

type SolutionDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (SolutionDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SolutionDetailsResponse)(nil)).Elem()
}

func (o SolutionDetailsResponsePtrOutput) ToSolutionDetailsResponsePtrOutput() SolutionDetailsResponsePtrOutput {
	return o
}

func (o SolutionDetailsResponsePtrOutput) ToSolutionDetailsResponsePtrOutputWithContext(ctx context.Context) SolutionDetailsResponsePtrOutput {
	return o
}

func (o SolutionDetailsResponsePtrOutput) Elem() SolutionDetailsResponseOutput {
	return o.ApplyT(func(v *SolutionDetailsResponse) SolutionDetailsResponse {
		if v != nil {
			return *v
		}
		var ret SolutionDetailsResponse
		return ret
	}).(SolutionDetailsResponseOutput)
}

func (o SolutionDetailsResponsePtrOutput) AssessmentCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SolutionDetailsResponse) *int {
		if v == nil {
			return nil
		}
		return v.AssessmentCount
	}).(pulumi.IntPtrOutput)
}

func (o SolutionDetailsResponsePtrOutput) ExtendedDetails() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SolutionDetailsResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.ExtendedDetails
	}).(pulumi.StringMapOutput)
}

func (o SolutionDetailsResponsePtrOutput) GroupCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SolutionDetailsResponse) *int {
		if v == nil {
			return nil
		}
		return v.GroupCount
	}).(pulumi.IntPtrOutput)
}

type SolutionProperties struct {
	CleanupState *string          `pulumi:"cleanupState"`
	Details      *SolutionDetails `pulumi:"details"`
	Goal         *string          `pulumi:"goal"`
	Purpose      *string          `pulumi:"purpose"`
	Status       *string          `pulumi:"status"`
	Tool         *string          `pulumi:"tool"`
}





type SolutionPropertiesInput interface {
	pulumi.Input

	ToSolutionPropertiesOutput() SolutionPropertiesOutput
	ToSolutionPropertiesOutputWithContext(context.Context) SolutionPropertiesOutput
}

type SolutionPropertiesArgs struct {
	CleanupState pulumi.StringPtrInput   `pulumi:"cleanupState"`
	Details      SolutionDetailsPtrInput `pulumi:"details"`
	Goal         pulumi.StringPtrInput   `pulumi:"goal"`
	Purpose      pulumi.StringPtrInput   `pulumi:"purpose"`
	Status       pulumi.StringPtrInput   `pulumi:"status"`
	Tool         pulumi.StringPtrInput   `pulumi:"tool"`
}

func (SolutionPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SolutionProperties)(nil)).Elem()
}

func (i SolutionPropertiesArgs) ToSolutionPropertiesOutput() SolutionPropertiesOutput {
	return i.ToSolutionPropertiesOutputWithContext(context.Background())
}

func (i SolutionPropertiesArgs) ToSolutionPropertiesOutputWithContext(ctx context.Context) SolutionPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SolutionPropertiesOutput)
}

func (i SolutionPropertiesArgs) ToSolutionPropertiesPtrOutput() SolutionPropertiesPtrOutput {
	return i.ToSolutionPropertiesPtrOutputWithContext(context.Background())
}

func (i SolutionPropertiesArgs) ToSolutionPropertiesPtrOutputWithContext(ctx context.Context) SolutionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SolutionPropertiesOutput).ToSolutionPropertiesPtrOutputWithContext(ctx)
}









type SolutionPropertiesPtrInput interface {
	pulumi.Input

	ToSolutionPropertiesPtrOutput() SolutionPropertiesPtrOutput
	ToSolutionPropertiesPtrOutputWithContext(context.Context) SolutionPropertiesPtrOutput
}

type solutionPropertiesPtrType SolutionPropertiesArgs

func SolutionPropertiesPtr(v *SolutionPropertiesArgs) SolutionPropertiesPtrInput {
	return (*solutionPropertiesPtrType)(v)
}

func (*solutionPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SolutionProperties)(nil)).Elem()
}

func (i *solutionPropertiesPtrType) ToSolutionPropertiesPtrOutput() SolutionPropertiesPtrOutput {
	return i.ToSolutionPropertiesPtrOutputWithContext(context.Background())
}

func (i *solutionPropertiesPtrType) ToSolutionPropertiesPtrOutputWithContext(ctx context.Context) SolutionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SolutionPropertiesPtrOutput)
}

type SolutionPropertiesOutput struct{ *pulumi.OutputState }

func (SolutionPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SolutionProperties)(nil)).Elem()
}

func (o SolutionPropertiesOutput) ToSolutionPropertiesOutput() SolutionPropertiesOutput {
	return o
}

func (o SolutionPropertiesOutput) ToSolutionPropertiesOutputWithContext(ctx context.Context) SolutionPropertiesOutput {
	return o
}

func (o SolutionPropertiesOutput) ToSolutionPropertiesPtrOutput() SolutionPropertiesPtrOutput {
	return o.ToSolutionPropertiesPtrOutputWithContext(context.Background())
}

func (o SolutionPropertiesOutput) ToSolutionPropertiesPtrOutputWithContext(ctx context.Context) SolutionPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SolutionProperties) *SolutionProperties {
		return &v
	}).(SolutionPropertiesPtrOutput)
}

func (o SolutionPropertiesOutput) CleanupState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SolutionProperties) *string { return v.CleanupState }).(pulumi.StringPtrOutput)
}

func (o SolutionPropertiesOutput) Details() SolutionDetailsPtrOutput {
	return o.ApplyT(func(v SolutionProperties) *SolutionDetails { return v.Details }).(SolutionDetailsPtrOutput)
}

func (o SolutionPropertiesOutput) Goal() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SolutionProperties) *string { return v.Goal }).(pulumi.StringPtrOutput)
}

func (o SolutionPropertiesOutput) Purpose() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SolutionProperties) *string { return v.Purpose }).(pulumi.StringPtrOutput)
}

func (o SolutionPropertiesOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SolutionProperties) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o SolutionPropertiesOutput) Tool() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SolutionProperties) *string { return v.Tool }).(pulumi.StringPtrOutput)
}

type SolutionPropertiesPtrOutput struct{ *pulumi.OutputState }

func (SolutionPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SolutionProperties)(nil)).Elem()
}

func (o SolutionPropertiesPtrOutput) ToSolutionPropertiesPtrOutput() SolutionPropertiesPtrOutput {
	return o
}

func (o SolutionPropertiesPtrOutput) ToSolutionPropertiesPtrOutputWithContext(ctx context.Context) SolutionPropertiesPtrOutput {
	return o
}

func (o SolutionPropertiesPtrOutput) Elem() SolutionPropertiesOutput {
	return o.ApplyT(func(v *SolutionProperties) SolutionProperties {
		if v != nil {
			return *v
		}
		var ret SolutionProperties
		return ret
	}).(SolutionPropertiesOutput)
}

func (o SolutionPropertiesPtrOutput) CleanupState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SolutionProperties) *string {
		if v == nil {
			return nil
		}
		return v.CleanupState
	}).(pulumi.StringPtrOutput)
}

func (o SolutionPropertiesPtrOutput) Details() SolutionDetailsPtrOutput {
	return o.ApplyT(func(v *SolutionProperties) *SolutionDetails {
		if v == nil {
			return nil
		}
		return v.Details
	}).(SolutionDetailsPtrOutput)
}

func (o SolutionPropertiesPtrOutput) Goal() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SolutionProperties) *string {
		if v == nil {
			return nil
		}
		return v.Goal
	}).(pulumi.StringPtrOutput)
}

func (o SolutionPropertiesPtrOutput) Purpose() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SolutionProperties) *string {
		if v == nil {
			return nil
		}
		return v.Purpose
	}).(pulumi.StringPtrOutput)
}

func (o SolutionPropertiesPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SolutionProperties) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

func (o SolutionPropertiesPtrOutput) Tool() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SolutionProperties) *string {
		if v == nil {
			return nil
		}
		return v.Tool
	}).(pulumi.StringPtrOutput)
}

type SolutionPropertiesResponse struct {
	CleanupState *string                  `pulumi:"cleanupState"`
	Details      *SolutionDetailsResponse `pulumi:"details"`
	Goal         *string                  `pulumi:"goal"`
	Purpose      *string                  `pulumi:"purpose"`
	Status       *string                  `pulumi:"status"`
	Summary      interface{}              `pulumi:"summary"`
	Tool         *string                  `pulumi:"tool"`
}

type SolutionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (SolutionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SolutionPropertiesResponse)(nil)).Elem()
}

func (o SolutionPropertiesResponseOutput) ToSolutionPropertiesResponseOutput() SolutionPropertiesResponseOutput {
	return o
}

func (o SolutionPropertiesResponseOutput) ToSolutionPropertiesResponseOutputWithContext(ctx context.Context) SolutionPropertiesResponseOutput {
	return o
}

func (o SolutionPropertiesResponseOutput) CleanupState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SolutionPropertiesResponse) *string { return v.CleanupState }).(pulumi.StringPtrOutput)
}

func (o SolutionPropertiesResponseOutput) Details() SolutionDetailsResponsePtrOutput {
	return o.ApplyT(func(v SolutionPropertiesResponse) *SolutionDetailsResponse { return v.Details }).(SolutionDetailsResponsePtrOutput)
}

func (o SolutionPropertiesResponseOutput) Goal() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SolutionPropertiesResponse) *string { return v.Goal }).(pulumi.StringPtrOutput)
}

func (o SolutionPropertiesResponseOutput) Purpose() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SolutionPropertiesResponse) *string { return v.Purpose }).(pulumi.StringPtrOutput)
}

func (o SolutionPropertiesResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SolutionPropertiesResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o SolutionPropertiesResponseOutput) Summary() pulumi.AnyOutput {
	return o.ApplyT(func(v SolutionPropertiesResponse) interface{} { return v.Summary }).(pulumi.AnyOutput)
}

func (o SolutionPropertiesResponseOutput) Tool() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SolutionPropertiesResponse) *string { return v.Tool }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(MigrateProjectPropertiesOutput{})
	pulumi.RegisterOutputType(MigrateProjectPropertiesPtrOutput{})
	pulumi.RegisterOutputType(MigrateProjectPropertiesResponseOutput{})
	pulumi.RegisterOutputType(MigrateProjectResponseTagsOutput{})
	pulumi.RegisterOutputType(MigrateProjectResponseTagsPtrOutput{})
	pulumi.RegisterOutputType(MigrateProjectTagsOutput{})
	pulumi.RegisterOutputType(MigrateProjectTagsPtrOutput{})
	pulumi.RegisterOutputType(SolutionDetailsOutput{})
	pulumi.RegisterOutputType(SolutionDetailsPtrOutput{})
	pulumi.RegisterOutputType(SolutionDetailsResponseOutput{})
	pulumi.RegisterOutputType(SolutionDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(SolutionPropertiesOutput{})
	pulumi.RegisterOutputType(SolutionPropertiesPtrOutput{})
	pulumi.RegisterOutputType(SolutionPropertiesResponseOutput{})
}
