


package v20201001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupDeploymentScript(ctx *pulumi.Context, args *LookupDeploymentScriptArgs, opts ...pulumi.InvokeOption) (*LookupDeploymentScriptResult, error) {
	var rv LookupDeploymentScriptResult
	err := ctx.Invoke("azure-native:resources/v20201001:getDeploymentScript", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDeploymentScriptArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ScriptName        string `pulumi:"scriptName"`
}


type LookupDeploymentScriptResult struct {
	Id         string                          `pulumi:"id"`
	Identity   *ManagedServiceIdentityResponse `pulumi:"identity"`
	Kind       string                          `pulumi:"kind"`
	Location   string                          `pulumi:"location"`
	Name       string                          `pulumi:"name"`
	SystemData SystemDataResponse              `pulumi:"systemData"`
	Tags       map[string]string               `pulumi:"tags"`
	Type       string                          `pulumi:"type"`
}

func LookupDeploymentScriptOutput(ctx *pulumi.Context, args LookupDeploymentScriptOutputArgs, opts ...pulumi.InvokeOption) LookupDeploymentScriptResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDeploymentScriptResult, error) {
			args := v.(LookupDeploymentScriptArgs)
			r, err := LookupDeploymentScript(ctx, &args, opts...)
			var s LookupDeploymentScriptResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDeploymentScriptResultOutput)
}

type LookupDeploymentScriptOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ScriptName        pulumi.StringInput `pulumi:"scriptName"`
}

func (LookupDeploymentScriptOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeploymentScriptArgs)(nil)).Elem()
}


type LookupDeploymentScriptResultOutput struct{ *pulumi.OutputState }

func (LookupDeploymentScriptResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeploymentScriptResult)(nil)).Elem()
}

func (o LookupDeploymentScriptResultOutput) ToLookupDeploymentScriptResultOutput() LookupDeploymentScriptResultOutput {
	return o
}

func (o LookupDeploymentScriptResultOutput) ToLookupDeploymentScriptResultOutputWithContext(ctx context.Context) LookupDeploymentScriptResultOutput {
	return o
}

func (o LookupDeploymentScriptResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentScriptResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDeploymentScriptResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupDeploymentScriptResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o LookupDeploymentScriptResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentScriptResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupDeploymentScriptResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentScriptResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupDeploymentScriptResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentScriptResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDeploymentScriptResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDeploymentScriptResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupDeploymentScriptResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDeploymentScriptResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDeploymentScriptResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentScriptResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDeploymentScriptResultOutput{})
}
