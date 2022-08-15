


package automation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupHybridRunbookWorkerGroup(ctx *pulumi.Context, args *LookupHybridRunbookWorkerGroupArgs, opts ...pulumi.InvokeOption) (*LookupHybridRunbookWorkerGroupResult, error) {
	var rv LookupHybridRunbookWorkerGroupResult
	err := ctx.Invoke("azure-native:automation:getHybridRunbookWorkerGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHybridRunbookWorkerGroupArgs struct {
	AutomationAccountName        string `pulumi:"automationAccountName"`
	HybridRunbookWorkerGroupName string `pulumi:"hybridRunbookWorkerGroupName"`
	ResourceGroupName            string `pulumi:"resourceGroupName"`
}


type LookupHybridRunbookWorkerGroupResult struct {
	Credential           *RunAsCredentialAssociationPropertyResponse `pulumi:"credential"`
	GroupType            *string                                     `pulumi:"groupType"`
	HybridRunbookWorkers []HybridRunbookWorkerLegacyResponse         `pulumi:"hybridRunbookWorkers"`
	Id                   *string                                     `pulumi:"id"`
	Name                 *string                                     `pulumi:"name"`
	SystemData           SystemDataResponse                          `pulumi:"systemData"`
	Type                 string                                      `pulumi:"type"`
}

func LookupHybridRunbookWorkerGroupOutput(ctx *pulumi.Context, args LookupHybridRunbookWorkerGroupOutputArgs, opts ...pulumi.InvokeOption) LookupHybridRunbookWorkerGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHybridRunbookWorkerGroupResult, error) {
			args := v.(LookupHybridRunbookWorkerGroupArgs)
			r, err := LookupHybridRunbookWorkerGroup(ctx, &args, opts...)
			var s LookupHybridRunbookWorkerGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupHybridRunbookWorkerGroupResultOutput)
}

type LookupHybridRunbookWorkerGroupOutputArgs struct {
	AutomationAccountName        pulumi.StringInput `pulumi:"automationAccountName"`
	HybridRunbookWorkerGroupName pulumi.StringInput `pulumi:"hybridRunbookWorkerGroupName"`
	ResourceGroupName            pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupHybridRunbookWorkerGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHybridRunbookWorkerGroupArgs)(nil)).Elem()
}


type LookupHybridRunbookWorkerGroupResultOutput struct{ *pulumi.OutputState }

func (LookupHybridRunbookWorkerGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHybridRunbookWorkerGroupResult)(nil)).Elem()
}

func (o LookupHybridRunbookWorkerGroupResultOutput) ToLookupHybridRunbookWorkerGroupResultOutput() LookupHybridRunbookWorkerGroupResultOutput {
	return o
}

func (o LookupHybridRunbookWorkerGroupResultOutput) ToLookupHybridRunbookWorkerGroupResultOutputWithContext(ctx context.Context) LookupHybridRunbookWorkerGroupResultOutput {
	return o
}

func (o LookupHybridRunbookWorkerGroupResultOutput) Credential() RunAsCredentialAssociationPropertyResponsePtrOutput {
	return o.ApplyT(func(v LookupHybridRunbookWorkerGroupResult) *RunAsCredentialAssociationPropertyResponse {
		return v.Credential
	}).(RunAsCredentialAssociationPropertyResponsePtrOutput)
}

func (o LookupHybridRunbookWorkerGroupResultOutput) GroupType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHybridRunbookWorkerGroupResult) *string { return v.GroupType }).(pulumi.StringPtrOutput)
}

func (o LookupHybridRunbookWorkerGroupResultOutput) HybridRunbookWorkers() HybridRunbookWorkerLegacyResponseArrayOutput {
	return o.ApplyT(func(v LookupHybridRunbookWorkerGroupResult) []HybridRunbookWorkerLegacyResponse {
		return v.HybridRunbookWorkers
	}).(HybridRunbookWorkerLegacyResponseArrayOutput)
}

func (o LookupHybridRunbookWorkerGroupResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHybridRunbookWorkerGroupResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupHybridRunbookWorkerGroupResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHybridRunbookWorkerGroupResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupHybridRunbookWorkerGroupResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupHybridRunbookWorkerGroupResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupHybridRunbookWorkerGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHybridRunbookWorkerGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHybridRunbookWorkerGroupResultOutput{})
}
