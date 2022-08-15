


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupContainerAppsSourceControl(ctx *pulumi.Context, args *LookupContainerAppsSourceControlArgs, opts ...pulumi.InvokeOption) (*LookupContainerAppsSourceControlResult, error) {
	var rv LookupContainerAppsSourceControlResult
	err := ctx.Invoke("azure-native:app/v20220101preview:getContainerAppsSourceControl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupContainerAppsSourceControlArgs struct {
	ContainerAppName  string `pulumi:"containerAppName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupContainerAppsSourceControlResult struct {
	Branch                    *string                            `pulumi:"branch"`
	GithubActionConfiguration *GithubActionConfigurationResponse `pulumi:"githubActionConfiguration"`
	Id                        string                             `pulumi:"id"`
	Name                      string                             `pulumi:"name"`
	OperationState            string                             `pulumi:"operationState"`
	RepoUrl                   *string                            `pulumi:"repoUrl"`
	SystemData                SystemDataResponse                 `pulumi:"systemData"`
	Type                      string                             `pulumi:"type"`
}

func LookupContainerAppsSourceControlOutput(ctx *pulumi.Context, args LookupContainerAppsSourceControlOutputArgs, opts ...pulumi.InvokeOption) LookupContainerAppsSourceControlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupContainerAppsSourceControlResult, error) {
			args := v.(LookupContainerAppsSourceControlArgs)
			r, err := LookupContainerAppsSourceControl(ctx, &args, opts...)
			var s LookupContainerAppsSourceControlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupContainerAppsSourceControlResultOutput)
}

type LookupContainerAppsSourceControlOutputArgs struct {
	ContainerAppName  pulumi.StringInput `pulumi:"containerAppName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupContainerAppsSourceControlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerAppsSourceControlArgs)(nil)).Elem()
}


type LookupContainerAppsSourceControlResultOutput struct{ *pulumi.OutputState }

func (LookupContainerAppsSourceControlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerAppsSourceControlResult)(nil)).Elem()
}

func (o LookupContainerAppsSourceControlResultOutput) ToLookupContainerAppsSourceControlResultOutput() LookupContainerAppsSourceControlResultOutput {
	return o
}

func (o LookupContainerAppsSourceControlResultOutput) ToLookupContainerAppsSourceControlResultOutputWithContext(ctx context.Context) LookupContainerAppsSourceControlResultOutput {
	return o
}

func (o LookupContainerAppsSourceControlResultOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerAppsSourceControlResult) *string { return v.Branch }).(pulumi.StringPtrOutput)
}

func (o LookupContainerAppsSourceControlResultOutput) GithubActionConfiguration() GithubActionConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupContainerAppsSourceControlResult) *GithubActionConfigurationResponse {
		return v.GithubActionConfiguration
	}).(GithubActionConfigurationResponsePtrOutput)
}

func (o LookupContainerAppsSourceControlResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerAppsSourceControlResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupContainerAppsSourceControlResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerAppsSourceControlResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupContainerAppsSourceControlResultOutput) OperationState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerAppsSourceControlResult) string { return v.OperationState }).(pulumi.StringOutput)
}

func (o LookupContainerAppsSourceControlResultOutput) RepoUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerAppsSourceControlResult) *string { return v.RepoUrl }).(pulumi.StringPtrOutput)
}

func (o LookupContainerAppsSourceControlResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupContainerAppsSourceControlResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupContainerAppsSourceControlResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerAppsSourceControlResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupContainerAppsSourceControlResultOutput{})
}
