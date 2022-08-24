


package v20180201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBuildTask(ctx *pulumi.Context, args *LookupBuildTaskArgs, opts ...pulumi.InvokeOption) (*LookupBuildTaskResult, error) {
	var rv LookupBuildTaskResult
	err := ctx.Invoke("azure-native:containerregistry/v20180201preview:getBuildTask", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupBuildTaskArgs struct {
	BuildTaskName     string `pulumi:"buildTaskName"`
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupBuildTaskResult struct {
	Alias             string                             `pulumi:"alias"`
	CreationDate      string                             `pulumi:"creationDate"`
	Id                string                             `pulumi:"id"`
	Location          string                             `pulumi:"location"`
	Name              string                             `pulumi:"name"`
	Platform          PlatformPropertiesResponse         `pulumi:"platform"`
	ProvisioningState string                             `pulumi:"provisioningState"`
	SourceRepository  SourceRepositoryPropertiesResponse `pulumi:"sourceRepository"`
	Status            *string                            `pulumi:"status"`
	Tags              map[string]string                  `pulumi:"tags"`
	Timeout           *int                               `pulumi:"timeout"`
	Type              string                             `pulumi:"type"`
}


func (val *LookupBuildTaskResult) Defaults() *LookupBuildTaskResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.SourceRepository = *tmp.SourceRepository.Defaults()

	if isZero(tmp.Timeout) {
		timeout_ := 3600
		tmp.Timeout = &timeout_
	}
	return &tmp
}

func LookupBuildTaskOutput(ctx *pulumi.Context, args LookupBuildTaskOutputArgs, opts ...pulumi.InvokeOption) LookupBuildTaskResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBuildTaskResult, error) {
			args := v.(LookupBuildTaskArgs)
			r, err := LookupBuildTask(ctx, &args, opts...)
			var s LookupBuildTaskResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBuildTaskResultOutput)
}

type LookupBuildTaskOutputArgs struct {
	BuildTaskName     pulumi.StringInput `pulumi:"buildTaskName"`
	RegistryName      pulumi.StringInput `pulumi:"registryName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupBuildTaskOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBuildTaskArgs)(nil)).Elem()
}


type LookupBuildTaskResultOutput struct{ *pulumi.OutputState }

func (LookupBuildTaskResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBuildTaskResult)(nil)).Elem()
}

func (o LookupBuildTaskResultOutput) ToLookupBuildTaskResultOutput() LookupBuildTaskResultOutput {
	return o
}

func (o LookupBuildTaskResultOutput) ToLookupBuildTaskResultOutputWithContext(ctx context.Context) LookupBuildTaskResultOutput {
	return o
}

func (o LookupBuildTaskResultOutput) Alias() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBuildTaskResult) string { return v.Alias }).(pulumi.StringOutput)
}

func (o LookupBuildTaskResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBuildTaskResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

func (o LookupBuildTaskResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBuildTaskResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBuildTaskResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBuildTaskResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupBuildTaskResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBuildTaskResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBuildTaskResultOutput) Platform() PlatformPropertiesResponseOutput {
	return o.ApplyT(func(v LookupBuildTaskResult) PlatformPropertiesResponse { return v.Platform }).(PlatformPropertiesResponseOutput)
}

func (o LookupBuildTaskResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBuildTaskResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupBuildTaskResultOutput) SourceRepository() SourceRepositoryPropertiesResponseOutput {
	return o.ApplyT(func(v LookupBuildTaskResult) SourceRepositoryPropertiesResponse { return v.SourceRepository }).(SourceRepositoryPropertiesResponseOutput)
}

func (o LookupBuildTaskResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBuildTaskResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o LookupBuildTaskResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupBuildTaskResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupBuildTaskResultOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupBuildTaskResult) *int { return v.Timeout }).(pulumi.IntPtrOutput)
}

func (o LookupBuildTaskResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBuildTaskResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBuildTaskResultOutput{})
}
