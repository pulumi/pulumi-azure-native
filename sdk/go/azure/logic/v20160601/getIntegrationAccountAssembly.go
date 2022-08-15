


package v20160601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIntegrationAccountAssembly(ctx *pulumi.Context, args *LookupIntegrationAccountAssemblyArgs, opts ...pulumi.InvokeOption) (*LookupIntegrationAccountAssemblyResult, error) {
	var rv LookupIntegrationAccountAssemblyResult
	err := ctx.Invoke("azure-native:logic/v20160601:getIntegrationAccountAssembly", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIntegrationAccountAssemblyArgs struct {
	AssemblyArtifactName   string `pulumi:"assemblyArtifactName"`
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
}


type LookupIntegrationAccountAssemblyResult struct {
	Id         string                     `pulumi:"id"`
	Location   *string                    `pulumi:"location"`
	Name       string                     `pulumi:"name"`
	Properties AssemblyPropertiesResponse `pulumi:"properties"`
	Tags       map[string]string          `pulumi:"tags"`
	Type       string                     `pulumi:"type"`
}

func LookupIntegrationAccountAssemblyOutput(ctx *pulumi.Context, args LookupIntegrationAccountAssemblyOutputArgs, opts ...pulumi.InvokeOption) LookupIntegrationAccountAssemblyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIntegrationAccountAssemblyResult, error) {
			args := v.(LookupIntegrationAccountAssemblyArgs)
			r, err := LookupIntegrationAccountAssembly(ctx, &args, opts...)
			var s LookupIntegrationAccountAssemblyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIntegrationAccountAssemblyResultOutput)
}

type LookupIntegrationAccountAssemblyOutputArgs struct {
	AssemblyArtifactName   pulumi.StringInput `pulumi:"assemblyArtifactName"`
	IntegrationAccountName pulumi.StringInput `pulumi:"integrationAccountName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupIntegrationAccountAssemblyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntegrationAccountAssemblyArgs)(nil)).Elem()
}


type LookupIntegrationAccountAssemblyResultOutput struct{ *pulumi.OutputState }

func (LookupIntegrationAccountAssemblyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntegrationAccountAssemblyResult)(nil)).Elem()
}

func (o LookupIntegrationAccountAssemblyResultOutput) ToLookupIntegrationAccountAssemblyResultOutput() LookupIntegrationAccountAssemblyResultOutput {
	return o
}

func (o LookupIntegrationAccountAssemblyResultOutput) ToLookupIntegrationAccountAssemblyResultOutputWithContext(ctx context.Context) LookupIntegrationAccountAssemblyResultOutput {
	return o
}

func (o LookupIntegrationAccountAssemblyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountAssemblyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIntegrationAccountAssemblyResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationAccountAssemblyResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupIntegrationAccountAssemblyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountAssemblyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIntegrationAccountAssemblyResultOutput) Properties() AssemblyPropertiesResponseOutput {
	return o.ApplyT(func(v LookupIntegrationAccountAssemblyResult) AssemblyPropertiesResponse { return v.Properties }).(AssemblyPropertiesResponseOutput)
}

func (o LookupIntegrationAccountAssemblyResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupIntegrationAccountAssemblyResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupIntegrationAccountAssemblyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountAssemblyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIntegrationAccountAssemblyResultOutput{})
}
