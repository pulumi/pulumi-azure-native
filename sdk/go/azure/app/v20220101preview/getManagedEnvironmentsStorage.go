


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagedEnvironmentsStorage(ctx *pulumi.Context, args *LookupManagedEnvironmentsStorageArgs, opts ...pulumi.InvokeOption) (*LookupManagedEnvironmentsStorageResult, error) {
	var rv LookupManagedEnvironmentsStorageResult
	err := ctx.Invoke("azure-native:app/v20220101preview:getManagedEnvironmentsStorage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedEnvironmentsStorageArgs struct {
	EnvName           string `pulumi:"envName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupManagedEnvironmentsStorageResult struct {
	Id         string                                      `pulumi:"id"`
	Name       string                                      `pulumi:"name"`
	Properties ManagedEnvironmentStorageResponseProperties `pulumi:"properties"`
	SystemData SystemDataResponse                          `pulumi:"systemData"`
	Type       string                                      `pulumi:"type"`
}

func LookupManagedEnvironmentsStorageOutput(ctx *pulumi.Context, args LookupManagedEnvironmentsStorageOutputArgs, opts ...pulumi.InvokeOption) LookupManagedEnvironmentsStorageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagedEnvironmentsStorageResult, error) {
			args := v.(LookupManagedEnvironmentsStorageArgs)
			r, err := LookupManagedEnvironmentsStorage(ctx, &args, opts...)
			var s LookupManagedEnvironmentsStorageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagedEnvironmentsStorageResultOutput)
}

type LookupManagedEnvironmentsStorageOutputArgs struct {
	EnvName           pulumi.StringInput `pulumi:"envName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupManagedEnvironmentsStorageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedEnvironmentsStorageArgs)(nil)).Elem()
}


type LookupManagedEnvironmentsStorageResultOutput struct{ *pulumi.OutputState }

func (LookupManagedEnvironmentsStorageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedEnvironmentsStorageResult)(nil)).Elem()
}

func (o LookupManagedEnvironmentsStorageResultOutput) ToLookupManagedEnvironmentsStorageResultOutput() LookupManagedEnvironmentsStorageResultOutput {
	return o
}

func (o LookupManagedEnvironmentsStorageResultOutput) ToLookupManagedEnvironmentsStorageResultOutputWithContext(ctx context.Context) LookupManagedEnvironmentsStorageResultOutput {
	return o
}

func (o LookupManagedEnvironmentsStorageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentsStorageResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupManagedEnvironmentsStorageResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentsStorageResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupManagedEnvironmentsStorageResultOutput) Properties() ManagedEnvironmentStorageResponsePropertiesOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentsStorageResult) ManagedEnvironmentStorageResponseProperties {
		return v.Properties
	}).(ManagedEnvironmentStorageResponsePropertiesOutput)
}

func (o LookupManagedEnvironmentsStorageResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentsStorageResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupManagedEnvironmentsStorageResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentsStorageResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedEnvironmentsStorageResultOutput{})
}
