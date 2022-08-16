


package v20220601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEnvironmentVersion(ctx *pulumi.Context, args *LookupEnvironmentVersionArgs, opts ...pulumi.InvokeOption) (*LookupEnvironmentVersionResult, error) {
	var rv LookupEnvironmentVersionResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20220601preview:getEnvironmentVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupEnvironmentVersionArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Version           string `pulumi:"version"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupEnvironmentVersionResult struct {
	EnvironmentVersionProperties EnvironmentVersionResponse `pulumi:"environmentVersionProperties"`
	Id                           string                     `pulumi:"id"`
	Name                         string                     `pulumi:"name"`
	SystemData                   SystemDataResponse         `pulumi:"systemData"`
	Type                         string                     `pulumi:"type"`
}


func (val *LookupEnvironmentVersionResult) Defaults() *LookupEnvironmentVersionResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.EnvironmentVersionProperties = *tmp.EnvironmentVersionProperties.Defaults()

	return &tmp
}

func LookupEnvironmentVersionOutput(ctx *pulumi.Context, args LookupEnvironmentVersionOutputArgs, opts ...pulumi.InvokeOption) LookupEnvironmentVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEnvironmentVersionResult, error) {
			args := v.(LookupEnvironmentVersionArgs)
			r, err := LookupEnvironmentVersion(ctx, &args, opts...)
			var s LookupEnvironmentVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEnvironmentVersionResultOutput)
}

type LookupEnvironmentVersionOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Version           pulumi.StringInput `pulumi:"version"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupEnvironmentVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentVersionArgs)(nil)).Elem()
}


type LookupEnvironmentVersionResultOutput struct{ *pulumi.OutputState }

func (LookupEnvironmentVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentVersionResult)(nil)).Elem()
}

func (o LookupEnvironmentVersionResultOutput) ToLookupEnvironmentVersionResultOutput() LookupEnvironmentVersionResultOutput {
	return o
}

func (o LookupEnvironmentVersionResultOutput) ToLookupEnvironmentVersionResultOutputWithContext(ctx context.Context) LookupEnvironmentVersionResultOutput {
	return o
}

func (o LookupEnvironmentVersionResultOutput) EnvironmentVersionProperties() EnvironmentVersionResponseOutput {
	return o.ApplyT(func(v LookupEnvironmentVersionResult) EnvironmentVersionResponse {
		return v.EnvironmentVersionProperties
	}).(EnvironmentVersionResponseOutput)
}

func (o LookupEnvironmentVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupEnvironmentVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupEnvironmentVersionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupEnvironmentVersionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupEnvironmentVersionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentVersionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEnvironmentVersionResultOutput{})
}
