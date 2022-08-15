


package v20200501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateEndpointConnectionControllerPrivateEndpointConnection(ctx *pulumi.Context, args *LookupPrivateEndpointConnectionControllerPrivateEndpointConnectionArgs, opts ...pulumi.InvokeOption) (*LookupPrivateEndpointConnectionControllerPrivateEndpointConnectionResult, error) {
	var rv LookupPrivateEndpointConnectionControllerPrivateEndpointConnectionResult
	err := ctx.Invoke("azure-native:migrate/v20200501:getPrivateEndpointConnectionControllerPrivateEndpointConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateEndpointConnectionControllerPrivateEndpointConnectionArgs struct {
	MigrateProjectName string `pulumi:"migrateProjectName"`
	PeConnectionName   string `pulumi:"peConnectionName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupPrivateEndpointConnectionControllerPrivateEndpointConnectionResult struct {
	ETag       string                                      `pulumi:"eTag"`
	Id         string                                      `pulumi:"id"`
	Name       string                                      `pulumi:"name"`
	Properties PrivateEndpointConnectionPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse                          `pulumi:"systemData"`
	Type       string                                      `pulumi:"type"`
}

func LookupPrivateEndpointConnectionControllerPrivateEndpointConnectionOutput(ctx *pulumi.Context, args LookupPrivateEndpointConnectionControllerPrivateEndpointConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateEndpointConnectionControllerPrivateEndpointConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateEndpointConnectionControllerPrivateEndpointConnectionResult, error) {
			args := v.(LookupPrivateEndpointConnectionControllerPrivateEndpointConnectionArgs)
			r, err := LookupPrivateEndpointConnectionControllerPrivateEndpointConnection(ctx, &args, opts...)
			var s LookupPrivateEndpointConnectionControllerPrivateEndpointConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateEndpointConnectionControllerPrivateEndpointConnectionResultOutput)
}

type LookupPrivateEndpointConnectionControllerPrivateEndpointConnectionOutputArgs struct {
	MigrateProjectName pulumi.StringInput `pulumi:"migrateProjectName"`
	PeConnectionName   pulumi.StringInput `pulumi:"peConnectionName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPrivateEndpointConnectionControllerPrivateEndpointConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointConnectionControllerPrivateEndpointConnectionArgs)(nil)).Elem()
}


type LookupPrivateEndpointConnectionControllerPrivateEndpointConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateEndpointConnectionControllerPrivateEndpointConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointConnectionControllerPrivateEndpointConnectionResult)(nil)).Elem()
}

func (o LookupPrivateEndpointConnectionControllerPrivateEndpointConnectionResultOutput) ToLookupPrivateEndpointConnectionControllerPrivateEndpointConnectionResultOutput() LookupPrivateEndpointConnectionControllerPrivateEndpointConnectionResultOutput {
	return o
}

func (o LookupPrivateEndpointConnectionControllerPrivateEndpointConnectionResultOutput) ToLookupPrivateEndpointConnectionControllerPrivateEndpointConnectionResultOutputWithContext(ctx context.Context) LookupPrivateEndpointConnectionControllerPrivateEndpointConnectionResultOutput {
	return o
}

func (o LookupPrivateEndpointConnectionControllerPrivateEndpointConnectionResultOutput) ETag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionControllerPrivateEndpointConnectionResult) string { return v.ETag }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointConnectionControllerPrivateEndpointConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionControllerPrivateEndpointConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointConnectionControllerPrivateEndpointConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionControllerPrivateEndpointConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointConnectionControllerPrivateEndpointConnectionResultOutput) Properties() PrivateEndpointConnectionPropertiesResponseOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionControllerPrivateEndpointConnectionResult) PrivateEndpointConnectionPropertiesResponse {
		return v.Properties
	}).(PrivateEndpointConnectionPropertiesResponseOutput)
}

func (o LookupPrivateEndpointConnectionControllerPrivateEndpointConnectionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionControllerPrivateEndpointConnectionResult) SystemDataResponse {
		return v.SystemData
	}).(SystemDataResponseOutput)
}

func (o LookupPrivateEndpointConnectionControllerPrivateEndpointConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionControllerPrivateEndpointConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateEndpointConnectionControllerPrivateEndpointConnectionResultOutput{})
}
