


package v20220808

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRunbook(ctx *pulumi.Context, args *LookupRunbookArgs, opts ...pulumi.InvokeOption) (*LookupRunbookResult, error) {
	var rv LookupRunbookResult
	err := ctx.Invoke("azure-native:automation/v20220808:getRunbook", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRunbookArgs struct {
	AutomationAccountName string `pulumi:"automationAccountName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	RunbookName           string `pulumi:"runbookName"`
}


type LookupRunbookResult struct {
	CreationTime       *string                             `pulumi:"creationTime"`
	Description        *string                             `pulumi:"description"`
	Draft              *RunbookDraftResponse               `pulumi:"draft"`
	Etag               *string                             `pulumi:"etag"`
	Id                 string                              `pulumi:"id"`
	JobCount           *int                                `pulumi:"jobCount"`
	LastModifiedBy     *string                             `pulumi:"lastModifiedBy"`
	LastModifiedTime   *string                             `pulumi:"lastModifiedTime"`
	Location           *string                             `pulumi:"location"`
	LogActivityTrace   *int                                `pulumi:"logActivityTrace"`
	LogProgress        *bool                               `pulumi:"logProgress"`
	LogVerbose         *bool                               `pulumi:"logVerbose"`
	Name               string                              `pulumi:"name"`
	OutputTypes        []string                            `pulumi:"outputTypes"`
	Parameters         map[string]RunbookParameterResponse `pulumi:"parameters"`
	ProvisioningState  *string                             `pulumi:"provisioningState"`
	PublishContentLink *ContentLinkResponse                `pulumi:"publishContentLink"`
	RunbookType        *string                             `pulumi:"runbookType"`
	State              *string                             `pulumi:"state"`
	Tags               map[string]string                   `pulumi:"tags"`
	Type               string                              `pulumi:"type"`
}

func LookupRunbookOutput(ctx *pulumi.Context, args LookupRunbookOutputArgs, opts ...pulumi.InvokeOption) LookupRunbookResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRunbookResult, error) {
			args := v.(LookupRunbookArgs)
			r, err := LookupRunbook(ctx, &args, opts...)
			var s LookupRunbookResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRunbookResultOutput)
}

type LookupRunbookOutputArgs struct {
	AutomationAccountName pulumi.StringInput `pulumi:"automationAccountName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	RunbookName           pulumi.StringInput `pulumi:"runbookName"`
}

func (LookupRunbookOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRunbookArgs)(nil)).Elem()
}


type LookupRunbookResultOutput struct{ *pulumi.OutputState }

func (LookupRunbookResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRunbookResult)(nil)).Elem()
}

func (o LookupRunbookResultOutput) ToLookupRunbookResultOutput() LookupRunbookResultOutput {
	return o
}

func (o LookupRunbookResultOutput) ToLookupRunbookResultOutputWithContext(ctx context.Context) LookupRunbookResultOutput {
	return o
}

func (o LookupRunbookResultOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRunbookResult) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

func (o LookupRunbookResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRunbookResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupRunbookResultOutput) Draft() RunbookDraftResponsePtrOutput {
	return o.ApplyT(func(v LookupRunbookResult) *RunbookDraftResponse { return v.Draft }).(RunbookDraftResponsePtrOutput)
}

func (o LookupRunbookResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRunbookResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupRunbookResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRunbookResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRunbookResultOutput) JobCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRunbookResult) *int { return v.JobCount }).(pulumi.IntPtrOutput)
}

func (o LookupRunbookResultOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRunbookResult) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o LookupRunbookResultOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRunbookResult) *string { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

func (o LookupRunbookResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRunbookResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupRunbookResultOutput) LogActivityTrace() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRunbookResult) *int { return v.LogActivityTrace }).(pulumi.IntPtrOutput)
}

func (o LookupRunbookResultOutput) LogProgress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRunbookResult) *bool { return v.LogProgress }).(pulumi.BoolPtrOutput)
}

func (o LookupRunbookResultOutput) LogVerbose() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRunbookResult) *bool { return v.LogVerbose }).(pulumi.BoolPtrOutput)
}

func (o LookupRunbookResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRunbookResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRunbookResultOutput) OutputTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRunbookResult) []string { return v.OutputTypes }).(pulumi.StringArrayOutput)
}

func (o LookupRunbookResultOutput) Parameters() RunbookParameterResponseMapOutput {
	return o.ApplyT(func(v LookupRunbookResult) map[string]RunbookParameterResponse { return v.Parameters }).(RunbookParameterResponseMapOutput)
}

func (o LookupRunbookResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRunbookResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupRunbookResultOutput) PublishContentLink() ContentLinkResponsePtrOutput {
	return o.ApplyT(func(v LookupRunbookResult) *ContentLinkResponse { return v.PublishContentLink }).(ContentLinkResponsePtrOutput)
}

func (o LookupRunbookResultOutput) RunbookType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRunbookResult) *string { return v.RunbookType }).(pulumi.StringPtrOutput)
}

func (o LookupRunbookResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRunbookResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

func (o LookupRunbookResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRunbookResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupRunbookResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRunbookResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRunbookResultOutput{})
}
