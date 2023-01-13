


package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLinkerDryrun(ctx *pulumi.Context, args *LookupLinkerDryrunArgs, opts ...pulumi.InvokeOption) (*LookupLinkerDryrunResult, error) {
	var rv LookupLinkerDryrunResult
	err := ctx.Invoke("azure-native:servicelinker/v20221101preview:getLinkerDryrun", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLinkerDryrunArgs struct {
	DryrunName  string `pulumi:"dryrunName"`
	ResourceUri string `pulumi:"resourceUri"`
}


type LookupLinkerDryrunResult struct {
	Id                  string                                  `pulumi:"id"`
	Name                string                                  `pulumi:"name"`
	OperationPreviews   []DryrunOperationPreviewResponse        `pulumi:"operationPreviews"`
	Parameters          *CreateOrUpdateDryrunParametersResponse `pulumi:"parameters"`
	PrerequisiteResults []interface{}                           `pulumi:"prerequisiteResults"`
	ProvisioningState   string                                  `pulumi:"provisioningState"`
	SystemData          SystemDataResponse                      `pulumi:"systemData"`
	Type                string                                  `pulumi:"type"`
}

func LookupLinkerDryrunOutput(ctx *pulumi.Context, args LookupLinkerDryrunOutputArgs, opts ...pulumi.InvokeOption) LookupLinkerDryrunResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLinkerDryrunResult, error) {
			args := v.(LookupLinkerDryrunArgs)
			r, err := LookupLinkerDryrun(ctx, &args, opts...)
			var s LookupLinkerDryrunResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLinkerDryrunResultOutput)
}

type LookupLinkerDryrunOutputArgs struct {
	DryrunName  pulumi.StringInput `pulumi:"dryrunName"`
	ResourceUri pulumi.StringInput `pulumi:"resourceUri"`
}

func (LookupLinkerDryrunOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLinkerDryrunArgs)(nil)).Elem()
}


type LookupLinkerDryrunResultOutput struct{ *pulumi.OutputState }

func (LookupLinkerDryrunResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLinkerDryrunResult)(nil)).Elem()
}

func (o LookupLinkerDryrunResultOutput) ToLookupLinkerDryrunResultOutput() LookupLinkerDryrunResultOutput {
	return o
}

func (o LookupLinkerDryrunResultOutput) ToLookupLinkerDryrunResultOutputWithContext(ctx context.Context) LookupLinkerDryrunResultOutput {
	return o
}

func (o LookupLinkerDryrunResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkerDryrunResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLinkerDryrunResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkerDryrunResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupLinkerDryrunResultOutput) OperationPreviews() DryrunOperationPreviewResponseArrayOutput {
	return o.ApplyT(func(v LookupLinkerDryrunResult) []DryrunOperationPreviewResponse { return v.OperationPreviews }).(DryrunOperationPreviewResponseArrayOutput)
}

func (o LookupLinkerDryrunResultOutput) Parameters() CreateOrUpdateDryrunParametersResponsePtrOutput {
	return o.ApplyT(func(v LookupLinkerDryrunResult) *CreateOrUpdateDryrunParametersResponse { return v.Parameters }).(CreateOrUpdateDryrunParametersResponsePtrOutput)
}

func (o LookupLinkerDryrunResultOutput) PrerequisiteResults() pulumi.ArrayOutput {
	return o.ApplyT(func(v LookupLinkerDryrunResult) []interface{} { return v.PrerequisiteResults }).(pulumi.ArrayOutput)
}

func (o LookupLinkerDryrunResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkerDryrunResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupLinkerDryrunResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupLinkerDryrunResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupLinkerDryrunResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkerDryrunResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLinkerDryrunResultOutput{})
}
