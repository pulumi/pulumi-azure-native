


package v20220901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGitHubConnector(ctx *pulumi.Context, args *LookupGitHubConnectorArgs, opts ...pulumi.InvokeOption) (*LookupGitHubConnectorResult, error) {
	var rv LookupGitHubConnectorResult
	err := ctx.Invoke("azure-native:securitydevops/v20220901preview:getGitHubConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGitHubConnectorArgs struct {
	GitHubConnectorName string `pulumi:"gitHubConnectorName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}


type LookupGitHubConnectorResult struct {
	Id         string                            `pulumi:"id"`
	Location   string                            `pulumi:"location"`
	Name       string                            `pulumi:"name"`
	Properties GitHubConnectorPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse                `pulumi:"systemData"`
	Tags       map[string]string                 `pulumi:"tags"`
	Type       string                            `pulumi:"type"`
}

func LookupGitHubConnectorOutput(ctx *pulumi.Context, args LookupGitHubConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupGitHubConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGitHubConnectorResult, error) {
			args := v.(LookupGitHubConnectorArgs)
			r, err := LookupGitHubConnector(ctx, &args, opts...)
			var s LookupGitHubConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGitHubConnectorResultOutput)
}

type LookupGitHubConnectorOutputArgs struct {
	GitHubConnectorName pulumi.StringInput `pulumi:"gitHubConnectorName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupGitHubConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGitHubConnectorArgs)(nil)).Elem()
}


type LookupGitHubConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupGitHubConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGitHubConnectorResult)(nil)).Elem()
}

func (o LookupGitHubConnectorResultOutput) ToLookupGitHubConnectorResultOutput() LookupGitHubConnectorResultOutput {
	return o
}

func (o LookupGitHubConnectorResultOutput) ToLookupGitHubConnectorResultOutputWithContext(ctx context.Context) LookupGitHubConnectorResultOutput {
	return o
}

func (o LookupGitHubConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGitHubConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGitHubConnectorResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGitHubConnectorResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupGitHubConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGitHubConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupGitHubConnectorResultOutput) Properties() GitHubConnectorPropertiesResponseOutput {
	return o.ApplyT(func(v LookupGitHubConnectorResult) GitHubConnectorPropertiesResponse { return v.Properties }).(GitHubConnectorPropertiesResponseOutput)
}

func (o LookupGitHubConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupGitHubConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupGitHubConnectorResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupGitHubConnectorResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupGitHubConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGitHubConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGitHubConnectorResultOutput{})
}
