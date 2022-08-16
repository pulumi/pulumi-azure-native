


package insights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDataCollectionEndpoint(ctx *pulumi.Context, args *LookupDataCollectionEndpointArgs, opts ...pulumi.InvokeOption) (*LookupDataCollectionEndpointResult, error) {
	var rv LookupDataCollectionEndpointResult
	err := ctx.Invoke("azure-native:insights:getDataCollectionEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataCollectionEndpointArgs struct {
	DataCollectionEndpointName string `pulumi:"dataCollectionEndpointName"`
	ResourceGroupName          string `pulumi:"resourceGroupName"`
}


type LookupDataCollectionEndpointResult struct {
	ConfigurationAccess *DataCollectionEndpointResponseConfigurationAccess `pulumi:"configurationAccess"`
	Description         *string                                            `pulumi:"description"`
	Etag                string                                             `pulumi:"etag"`
	Id                  string                                             `pulumi:"id"`
	ImmutableId         *string                                            `pulumi:"immutableId"`
	Kind                *string                                            `pulumi:"kind"`
	Location            string                                             `pulumi:"location"`
	LogsIngestion       *DataCollectionEndpointResponseLogsIngestion       `pulumi:"logsIngestion"`
	Name                string                                             `pulumi:"name"`
	NetworkAcls         *DataCollectionEndpointResponseNetworkAcls         `pulumi:"networkAcls"`
	ProvisioningState   string                                             `pulumi:"provisioningState"`
	SystemData          DataCollectionEndpointResourceResponseSystemData   `pulumi:"systemData"`
	Tags                map[string]string                                  `pulumi:"tags"`
	Type                string                                             `pulumi:"type"`
}

func LookupDataCollectionEndpointOutput(ctx *pulumi.Context, args LookupDataCollectionEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupDataCollectionEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDataCollectionEndpointResult, error) {
			args := v.(LookupDataCollectionEndpointArgs)
			r, err := LookupDataCollectionEndpoint(ctx, &args, opts...)
			var s LookupDataCollectionEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDataCollectionEndpointResultOutput)
}

type LookupDataCollectionEndpointOutputArgs struct {
	DataCollectionEndpointName pulumi.StringInput `pulumi:"dataCollectionEndpointName"`
	ResourceGroupName          pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDataCollectionEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataCollectionEndpointArgs)(nil)).Elem()
}


type LookupDataCollectionEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupDataCollectionEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataCollectionEndpointResult)(nil)).Elem()
}

func (o LookupDataCollectionEndpointResultOutput) ToLookupDataCollectionEndpointResultOutput() LookupDataCollectionEndpointResultOutput {
	return o
}

func (o LookupDataCollectionEndpointResultOutput) ToLookupDataCollectionEndpointResultOutputWithContext(ctx context.Context) LookupDataCollectionEndpointResultOutput {
	return o
}

func (o LookupDataCollectionEndpointResultOutput) ConfigurationAccess() DataCollectionEndpointResponseConfigurationAccessPtrOutput {
	return o.ApplyT(func(v LookupDataCollectionEndpointResult) *DataCollectionEndpointResponseConfigurationAccess {
		return v.ConfigurationAccess
	}).(DataCollectionEndpointResponseConfigurationAccessPtrOutput)
}

func (o LookupDataCollectionEndpointResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataCollectionEndpointResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupDataCollectionEndpointResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataCollectionEndpointResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupDataCollectionEndpointResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataCollectionEndpointResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDataCollectionEndpointResultOutput) ImmutableId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataCollectionEndpointResult) *string { return v.ImmutableId }).(pulumi.StringPtrOutput)
}

func (o LookupDataCollectionEndpointResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataCollectionEndpointResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupDataCollectionEndpointResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataCollectionEndpointResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupDataCollectionEndpointResultOutput) LogsIngestion() DataCollectionEndpointResponseLogsIngestionPtrOutput {
	return o.ApplyT(func(v LookupDataCollectionEndpointResult) *DataCollectionEndpointResponseLogsIngestion {
		return v.LogsIngestion
	}).(DataCollectionEndpointResponseLogsIngestionPtrOutput)
}

func (o LookupDataCollectionEndpointResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataCollectionEndpointResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDataCollectionEndpointResultOutput) NetworkAcls() DataCollectionEndpointResponseNetworkAclsPtrOutput {
	return o.ApplyT(func(v LookupDataCollectionEndpointResult) *DataCollectionEndpointResponseNetworkAcls {
		return v.NetworkAcls
	}).(DataCollectionEndpointResponseNetworkAclsPtrOutput)
}

func (o LookupDataCollectionEndpointResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataCollectionEndpointResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupDataCollectionEndpointResultOutput) SystemData() DataCollectionEndpointResourceResponseSystemDataOutput {
	return o.ApplyT(func(v LookupDataCollectionEndpointResult) DataCollectionEndpointResourceResponseSystemData {
		return v.SystemData
	}).(DataCollectionEndpointResourceResponseSystemDataOutput)
}

func (o LookupDataCollectionEndpointResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDataCollectionEndpointResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDataCollectionEndpointResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataCollectionEndpointResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataCollectionEndpointResultOutput{})
}
