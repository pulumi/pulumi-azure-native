


package machinelearningservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEnvironmentSpecificationVersion(ctx *pulumi.Context, args *LookupEnvironmentSpecificationVersionArgs, opts ...pulumi.InvokeOption) (*LookupEnvironmentSpecificationVersionResult, error) {
	var rv LookupEnvironmentSpecificationVersionResult
	err := ctx.Invoke("azure-native:machinelearningservices:getEnvironmentSpecificationVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEnvironmentSpecificationVersionArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Version           string `pulumi:"version"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupEnvironmentSpecificationVersionResult struct {
	Id         string                                  `pulumi:"id"`
	Name       string                                  `pulumi:"name"`
	Properties EnvironmentSpecificationVersionResponse `pulumi:"properties"`
	SystemData SystemDataResponse                      `pulumi:"systemData"`
	Type       string                                  `pulumi:"type"`
}

func LookupEnvironmentSpecificationVersionOutput(ctx *pulumi.Context, args LookupEnvironmentSpecificationVersionOutputArgs, opts ...pulumi.InvokeOption) LookupEnvironmentSpecificationVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEnvironmentSpecificationVersionResult, error) {
			args := v.(LookupEnvironmentSpecificationVersionArgs)
			r, err := LookupEnvironmentSpecificationVersion(ctx, &args, opts...)
			var s LookupEnvironmentSpecificationVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEnvironmentSpecificationVersionResultOutput)
}

type LookupEnvironmentSpecificationVersionOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Version           pulumi.StringInput `pulumi:"version"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupEnvironmentSpecificationVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentSpecificationVersionArgs)(nil)).Elem()
}


type LookupEnvironmentSpecificationVersionResultOutput struct{ *pulumi.OutputState }

func (LookupEnvironmentSpecificationVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentSpecificationVersionResult)(nil)).Elem()
}

func (o LookupEnvironmentSpecificationVersionResultOutput) ToLookupEnvironmentSpecificationVersionResultOutput() LookupEnvironmentSpecificationVersionResultOutput {
	return o
}

func (o LookupEnvironmentSpecificationVersionResultOutput) ToLookupEnvironmentSpecificationVersionResultOutputWithContext(ctx context.Context) LookupEnvironmentSpecificationVersionResultOutput {
	return o
}

func (o LookupEnvironmentSpecificationVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentSpecificationVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupEnvironmentSpecificationVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentSpecificationVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupEnvironmentSpecificationVersionResultOutput) Properties() EnvironmentSpecificationVersionResponseOutput {
	return o.ApplyT(func(v LookupEnvironmentSpecificationVersionResult) EnvironmentSpecificationVersionResponse {
		return v.Properties
	}).(EnvironmentSpecificationVersionResponseOutput)
}

func (o LookupEnvironmentSpecificationVersionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupEnvironmentSpecificationVersionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupEnvironmentSpecificationVersionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentSpecificationVersionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEnvironmentSpecificationVersionResultOutput{})
}
