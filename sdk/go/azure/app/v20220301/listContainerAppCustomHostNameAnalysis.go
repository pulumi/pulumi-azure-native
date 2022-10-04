


package v20220301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListContainerAppCustomHostNameAnalysis(ctx *pulumi.Context, args *ListContainerAppCustomHostNameAnalysisArgs, opts ...pulumi.InvokeOption) (*ListContainerAppCustomHostNameAnalysisResult, error) {
	var rv ListContainerAppCustomHostNameAnalysisResult
	err := ctx.Invoke("azure-native:app/v20220301:listContainerAppCustomHostNameAnalysis", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListContainerAppCustomHostNameAnalysisArgs struct {
	ContainerAppName  string  `pulumi:"containerAppName"`
	CustomHostname    *string `pulumi:"customHostname"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type ListContainerAppCustomHostNameAnalysisResult struct {
	ARecords                            []string                                                                `pulumi:"aRecords"`
	AlternateCNameRecords               []string                                                                `pulumi:"alternateCNameRecords"`
	AlternateTxtRecords                 []string                                                                `pulumi:"alternateTxtRecords"`
	CNameRecords                        []string                                                                `pulumi:"cNameRecords"`
	ConflictWithEnvironmentCustomDomain bool                                                                    `pulumi:"conflictWithEnvironmentCustomDomain"`
	ConflictingContainerAppResourceId   string                                                                  `pulumi:"conflictingContainerAppResourceId"`
	CustomDomainVerificationFailureInfo CustomHostnameAnalysisResultResponseCustomDomainVerificationFailureInfo `pulumi:"customDomainVerificationFailureInfo"`
	CustomDomainVerificationTest        string                                                                  `pulumi:"customDomainVerificationTest"`
	HasConflictOnManagedEnvironment     bool                                                                    `pulumi:"hasConflictOnManagedEnvironment"`
	HostName                            string                                                                  `pulumi:"hostName"`
	IsHostnameAlreadyVerified           bool                                                                    `pulumi:"isHostnameAlreadyVerified"`
	TxtRecords                          []string                                                                `pulumi:"txtRecords"`
}

func ListContainerAppCustomHostNameAnalysisOutput(ctx *pulumi.Context, args ListContainerAppCustomHostNameAnalysisOutputArgs, opts ...pulumi.InvokeOption) ListContainerAppCustomHostNameAnalysisResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListContainerAppCustomHostNameAnalysisResult, error) {
			args := v.(ListContainerAppCustomHostNameAnalysisArgs)
			r, err := ListContainerAppCustomHostNameAnalysis(ctx, &args, opts...)
			var s ListContainerAppCustomHostNameAnalysisResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListContainerAppCustomHostNameAnalysisResultOutput)
}

type ListContainerAppCustomHostNameAnalysisOutputArgs struct {
	ContainerAppName  pulumi.StringInput    `pulumi:"containerAppName"`
	CustomHostname    pulumi.StringPtrInput `pulumi:"customHostname"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (ListContainerAppCustomHostNameAnalysisOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListContainerAppCustomHostNameAnalysisArgs)(nil)).Elem()
}


type ListContainerAppCustomHostNameAnalysisResultOutput struct{ *pulumi.OutputState }

func (ListContainerAppCustomHostNameAnalysisResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListContainerAppCustomHostNameAnalysisResult)(nil)).Elem()
}

func (o ListContainerAppCustomHostNameAnalysisResultOutput) ToListContainerAppCustomHostNameAnalysisResultOutput() ListContainerAppCustomHostNameAnalysisResultOutput {
	return o
}

func (o ListContainerAppCustomHostNameAnalysisResultOutput) ToListContainerAppCustomHostNameAnalysisResultOutputWithContext(ctx context.Context) ListContainerAppCustomHostNameAnalysisResultOutput {
	return o
}

func (o ListContainerAppCustomHostNameAnalysisResultOutput) ARecords() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListContainerAppCustomHostNameAnalysisResult) []string { return v.ARecords }).(pulumi.StringArrayOutput)
}

func (o ListContainerAppCustomHostNameAnalysisResultOutput) AlternateCNameRecords() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListContainerAppCustomHostNameAnalysisResult) []string { return v.AlternateCNameRecords }).(pulumi.StringArrayOutput)
}

func (o ListContainerAppCustomHostNameAnalysisResultOutput) AlternateTxtRecords() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListContainerAppCustomHostNameAnalysisResult) []string { return v.AlternateTxtRecords }).(pulumi.StringArrayOutput)
}

func (o ListContainerAppCustomHostNameAnalysisResultOutput) CNameRecords() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListContainerAppCustomHostNameAnalysisResult) []string { return v.CNameRecords }).(pulumi.StringArrayOutput)
}

func (o ListContainerAppCustomHostNameAnalysisResultOutput) ConflictWithEnvironmentCustomDomain() pulumi.BoolOutput {
	return o.ApplyT(func(v ListContainerAppCustomHostNameAnalysisResult) bool {
		return v.ConflictWithEnvironmentCustomDomain
	}).(pulumi.BoolOutput)
}

func (o ListContainerAppCustomHostNameAnalysisResultOutput) ConflictingContainerAppResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v ListContainerAppCustomHostNameAnalysisResult) string {
		return v.ConflictingContainerAppResourceId
	}).(pulumi.StringOutput)
}

func (o ListContainerAppCustomHostNameAnalysisResultOutput) CustomDomainVerificationFailureInfo() CustomHostnameAnalysisResultResponseCustomDomainVerificationFailureInfoOutput {
	return o.ApplyT(func(v ListContainerAppCustomHostNameAnalysisResult) CustomHostnameAnalysisResultResponseCustomDomainVerificationFailureInfo {
		return v.CustomDomainVerificationFailureInfo
	}).(CustomHostnameAnalysisResultResponseCustomDomainVerificationFailureInfoOutput)
}

func (o ListContainerAppCustomHostNameAnalysisResultOutput) CustomDomainVerificationTest() pulumi.StringOutput {
	return o.ApplyT(func(v ListContainerAppCustomHostNameAnalysisResult) string { return v.CustomDomainVerificationTest }).(pulumi.StringOutput)
}

func (o ListContainerAppCustomHostNameAnalysisResultOutput) HasConflictOnManagedEnvironment() pulumi.BoolOutput {
	return o.ApplyT(func(v ListContainerAppCustomHostNameAnalysisResult) bool { return v.HasConflictOnManagedEnvironment }).(pulumi.BoolOutput)
}

func (o ListContainerAppCustomHostNameAnalysisResultOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v ListContainerAppCustomHostNameAnalysisResult) string { return v.HostName }).(pulumi.StringOutput)
}

func (o ListContainerAppCustomHostNameAnalysisResultOutput) IsHostnameAlreadyVerified() pulumi.BoolOutput {
	return o.ApplyT(func(v ListContainerAppCustomHostNameAnalysisResult) bool { return v.IsHostnameAlreadyVerified }).(pulumi.BoolOutput)
}

func (o ListContainerAppCustomHostNameAnalysisResultOutput) TxtRecords() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListContainerAppCustomHostNameAnalysisResult) []string { return v.TxtRecords }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListContainerAppCustomHostNameAnalysisResultOutput{})
}
