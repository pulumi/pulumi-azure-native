


package v20220707

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupScript(ctx *pulumi.Context, args *LookupScriptArgs, opts ...pulumi.InvokeOption) (*LookupScriptResult, error) {
	var rv LookupScriptResult
	err := ctx.Invoke("azure-native:kusto/v20220707:getScript", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupScriptArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	DatabaseName      string `pulumi:"databaseName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ScriptName        string `pulumi:"scriptName"`
}


type LookupScriptResult struct {
	ContinueOnErrors  *bool              `pulumi:"continueOnErrors"`
	ForceUpdateTag    *string            `pulumi:"forceUpdateTag"`
	Id                string             `pulumi:"id"`
	Name              string             `pulumi:"name"`
	ProvisioningState string             `pulumi:"provisioningState"`
	ScriptUrl         *string            `pulumi:"scriptUrl"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	Type              string             `pulumi:"type"`
}


func (val *LookupScriptResult) Defaults() *LookupScriptResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ContinueOnErrors) {
		continueOnErrors_ := false
		tmp.ContinueOnErrors = &continueOnErrors_
	}
	return &tmp
}

func LookupScriptOutput(ctx *pulumi.Context, args LookupScriptOutputArgs, opts ...pulumi.InvokeOption) LookupScriptResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupScriptResult, error) {
			args := v.(LookupScriptArgs)
			r, err := LookupScript(ctx, &args, opts...)
			var s LookupScriptResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupScriptResultOutput)
}

type LookupScriptOutputArgs struct {
	ClusterName       pulumi.StringInput `pulumi:"clusterName"`
	DatabaseName      pulumi.StringInput `pulumi:"databaseName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ScriptName        pulumi.StringInput `pulumi:"scriptName"`
}

func (LookupScriptOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScriptArgs)(nil)).Elem()
}


type LookupScriptResultOutput struct{ *pulumi.OutputState }

func (LookupScriptResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScriptResult)(nil)).Elem()
}

func (o LookupScriptResultOutput) ToLookupScriptResultOutput() LookupScriptResultOutput {
	return o
}

func (o LookupScriptResultOutput) ToLookupScriptResultOutputWithContext(ctx context.Context) LookupScriptResultOutput {
	return o
}

func (o LookupScriptResultOutput) ContinueOnErrors() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupScriptResult) *bool { return v.ContinueOnErrors }).(pulumi.BoolPtrOutput)
}

func (o LookupScriptResultOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScriptResult) *string { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

func (o LookupScriptResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScriptResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupScriptResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScriptResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupScriptResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScriptResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupScriptResultOutput) ScriptUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScriptResult) *string { return v.ScriptUrl }).(pulumi.StringPtrOutput)
}

func (o LookupScriptResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupScriptResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupScriptResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScriptResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupScriptResultOutput{})
}
