


package v20221001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConnectedEnvironmentsStorage(ctx *pulumi.Context, args *LookupConnectedEnvironmentsStorageArgs, opts ...pulumi.InvokeOption) (*LookupConnectedEnvironmentsStorageResult, error) {
	var rv LookupConnectedEnvironmentsStorageResult
	err := ctx.Invoke("azure-native:app/v20221001:getConnectedEnvironmentsStorage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectedEnvironmentsStorageArgs struct {
	ConnectedEnvironmentName string `pulumi:"connectedEnvironmentName"`
	ResourceGroupName        string `pulumi:"resourceGroupName"`
	StorageName              string `pulumi:"storageName"`
}


type LookupConnectedEnvironmentsStorageResult struct {
	Id         string                                        `pulumi:"id"`
	Name       string                                        `pulumi:"name"`
	Properties ConnectedEnvironmentStorageResponseProperties `pulumi:"properties"`
	SystemData SystemDataResponse                            `pulumi:"systemData"`
	Type       string                                        `pulumi:"type"`
}

func LookupConnectedEnvironmentsStorageOutput(ctx *pulumi.Context, args LookupConnectedEnvironmentsStorageOutputArgs, opts ...pulumi.InvokeOption) LookupConnectedEnvironmentsStorageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConnectedEnvironmentsStorageResult, error) {
			args := v.(LookupConnectedEnvironmentsStorageArgs)
			r, err := LookupConnectedEnvironmentsStorage(ctx, &args, opts...)
			var s LookupConnectedEnvironmentsStorageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConnectedEnvironmentsStorageResultOutput)
}

type LookupConnectedEnvironmentsStorageOutputArgs struct {
	ConnectedEnvironmentName pulumi.StringInput `pulumi:"connectedEnvironmentName"`
	ResourceGroupName        pulumi.StringInput `pulumi:"resourceGroupName"`
	StorageName              pulumi.StringInput `pulumi:"storageName"`
}

func (LookupConnectedEnvironmentsStorageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectedEnvironmentsStorageArgs)(nil)).Elem()
}


type LookupConnectedEnvironmentsStorageResultOutput struct{ *pulumi.OutputState }

func (LookupConnectedEnvironmentsStorageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectedEnvironmentsStorageResult)(nil)).Elem()
}

func (o LookupConnectedEnvironmentsStorageResultOutput) ToLookupConnectedEnvironmentsStorageResultOutput() LookupConnectedEnvironmentsStorageResultOutput {
	return o
}

func (o LookupConnectedEnvironmentsStorageResultOutput) ToLookupConnectedEnvironmentsStorageResultOutputWithContext(ctx context.Context) LookupConnectedEnvironmentsStorageResultOutput {
	return o
}

func (o LookupConnectedEnvironmentsStorageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsStorageResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupConnectedEnvironmentsStorageResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsStorageResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupConnectedEnvironmentsStorageResultOutput) Properties() ConnectedEnvironmentStorageResponsePropertiesOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsStorageResult) ConnectedEnvironmentStorageResponseProperties {
		return v.Properties
	}).(ConnectedEnvironmentStorageResponsePropertiesOutput)
}

func (o LookupConnectedEnvironmentsStorageResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsStorageResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupConnectedEnvironmentsStorageResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsStorageResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectedEnvironmentsStorageResultOutput{})
}
