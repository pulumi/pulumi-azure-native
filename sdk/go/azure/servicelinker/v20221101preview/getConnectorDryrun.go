


package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConnectorDryrun(ctx *pulumi.Context, args *LookupConnectorDryrunArgs, opts ...pulumi.InvokeOption) (*LookupConnectorDryrunResult, error) {
	var rv LookupConnectorDryrunResult
	err := ctx.Invoke("azure-native:servicelinker/v20221101preview:getConnectorDryrun", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectorDryrunArgs struct {
	DryrunName        string `pulumi:"dryrunName"`
	Location          string `pulumi:"location"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupConnectorDryrunResult struct {
	Id                  string                                  `pulumi:"id"`
	Name                string                                  `pulumi:"name"`
	OperationPreviews   []DryrunOperationPreviewResponse        `pulumi:"operationPreviews"`
	Parameters          *CreateOrUpdateDryrunParametersResponse `pulumi:"parameters"`
	PrerequisiteResults []interface{}                           `pulumi:"prerequisiteResults"`
	ProvisioningState   string                                  `pulumi:"provisioningState"`
	SystemData          SystemDataResponse                      `pulumi:"systemData"`
	Type                string                                  `pulumi:"type"`
}

func LookupConnectorDryrunOutput(ctx *pulumi.Context, args LookupConnectorDryrunOutputArgs, opts ...pulumi.InvokeOption) LookupConnectorDryrunResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConnectorDryrunResult, error) {
			args := v.(LookupConnectorDryrunArgs)
			r, err := LookupConnectorDryrun(ctx, &args, opts...)
			var s LookupConnectorDryrunResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConnectorDryrunResultOutput)
}

type LookupConnectorDryrunOutputArgs struct {
	DryrunName        pulumi.StringInput `pulumi:"dryrunName"`
	Location          pulumi.StringInput `pulumi:"location"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupConnectorDryrunOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectorDryrunArgs)(nil)).Elem()
}


type LookupConnectorDryrunResultOutput struct{ *pulumi.OutputState }

func (LookupConnectorDryrunResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectorDryrunResult)(nil)).Elem()
}

func (o LookupConnectorDryrunResultOutput) ToLookupConnectorDryrunResultOutput() LookupConnectorDryrunResultOutput {
	return o
}

func (o LookupConnectorDryrunResultOutput) ToLookupConnectorDryrunResultOutputWithContext(ctx context.Context) LookupConnectorDryrunResultOutput {
	return o
}

func (o LookupConnectorDryrunResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorDryrunResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupConnectorDryrunResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorDryrunResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupConnectorDryrunResultOutput) OperationPreviews() DryrunOperationPreviewResponseArrayOutput {
	return o.ApplyT(func(v LookupConnectorDryrunResult) []DryrunOperationPreviewResponse { return v.OperationPreviews }).(DryrunOperationPreviewResponseArrayOutput)
}

func (o LookupConnectorDryrunResultOutput) Parameters() CreateOrUpdateDryrunParametersResponsePtrOutput {
	return o.ApplyT(func(v LookupConnectorDryrunResult) *CreateOrUpdateDryrunParametersResponse { return v.Parameters }).(CreateOrUpdateDryrunParametersResponsePtrOutput)
}

func (o LookupConnectorDryrunResultOutput) PrerequisiteResults() pulumi.ArrayOutput {
	return o.ApplyT(func(v LookupConnectorDryrunResult) []interface{} { return v.PrerequisiteResults }).(pulumi.ArrayOutput)
}

func (o LookupConnectorDryrunResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorDryrunResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupConnectorDryrunResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupConnectorDryrunResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupConnectorDryrunResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorDryrunResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectorDryrunResultOutput{})
}
