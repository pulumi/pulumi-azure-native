


package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConnectedRegistry(ctx *pulumi.Context, args *LookupConnectedRegistryArgs, opts ...pulumi.InvokeOption) (*LookupConnectedRegistryResult, error) {
	var rv LookupConnectedRegistryResult
	err := ctx.Invoke("azure-native:containerregistry/v20210601preview:getConnectedRegistry", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupConnectedRegistryArgs struct {
	ConnectedRegistryName string `pulumi:"connectedRegistryName"`
	RegistryName          string `pulumi:"registryName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupConnectedRegistryResult struct {
	Activation        ActivationPropertiesResponse     `pulumi:"activation"`
	ClientTokenIds    []string                         `pulumi:"clientTokenIds"`
	ConnectionState   string                           `pulumi:"connectionState"`
	Id                string                           `pulumi:"id"`
	LastActivityTime  string                           `pulumi:"lastActivityTime"`
	Logging           *LoggingPropertiesResponse       `pulumi:"logging"`
	LoginServer       *LoginServerPropertiesResponse   `pulumi:"loginServer"`
	Mode              string                           `pulumi:"mode"`
	Name              string                           `pulumi:"name"`
	Parent            ParentPropertiesResponse         `pulumi:"parent"`
	ProvisioningState string                           `pulumi:"provisioningState"`
	StatusDetails     []StatusDetailPropertiesResponse `pulumi:"statusDetails"`
	SystemData        SystemDataResponse               `pulumi:"systemData"`
	Type              string                           `pulumi:"type"`
	Version           string                           `pulumi:"version"`
}


func (val *LookupConnectedRegistryResult) Defaults() *LookupConnectedRegistryResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Logging = tmp.Logging.Defaults()

	return &tmp
}

func LookupConnectedRegistryOutput(ctx *pulumi.Context, args LookupConnectedRegistryOutputArgs, opts ...pulumi.InvokeOption) LookupConnectedRegistryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConnectedRegistryResult, error) {
			args := v.(LookupConnectedRegistryArgs)
			r, err := LookupConnectedRegistry(ctx, &args, opts...)
			var s LookupConnectedRegistryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConnectedRegistryResultOutput)
}

type LookupConnectedRegistryOutputArgs struct {
	ConnectedRegistryName pulumi.StringInput `pulumi:"connectedRegistryName"`
	RegistryName          pulumi.StringInput `pulumi:"registryName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupConnectedRegistryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectedRegistryArgs)(nil)).Elem()
}


type LookupConnectedRegistryResultOutput struct{ *pulumi.OutputState }

func (LookupConnectedRegistryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectedRegistryResult)(nil)).Elem()
}

func (o LookupConnectedRegistryResultOutput) ToLookupConnectedRegistryResultOutput() LookupConnectedRegistryResultOutput {
	return o
}

func (o LookupConnectedRegistryResultOutput) ToLookupConnectedRegistryResultOutputWithContext(ctx context.Context) LookupConnectedRegistryResultOutput {
	return o
}

func (o LookupConnectedRegistryResultOutput) Activation() ActivationPropertiesResponseOutput {
	return o.ApplyT(func(v LookupConnectedRegistryResult) ActivationPropertiesResponse { return v.Activation }).(ActivationPropertiesResponseOutput)
}

func (o LookupConnectedRegistryResultOutput) ClientTokenIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupConnectedRegistryResult) []string { return v.ClientTokenIds }).(pulumi.StringArrayOutput)
}

func (o LookupConnectedRegistryResultOutput) ConnectionState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedRegistryResult) string { return v.ConnectionState }).(pulumi.StringOutput)
}

func (o LookupConnectedRegistryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedRegistryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupConnectedRegistryResultOutput) LastActivityTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedRegistryResult) string { return v.LastActivityTime }).(pulumi.StringOutput)
}

func (o LookupConnectedRegistryResultOutput) Logging() LoggingPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupConnectedRegistryResult) *LoggingPropertiesResponse { return v.Logging }).(LoggingPropertiesResponsePtrOutput)
}

func (o LookupConnectedRegistryResultOutput) LoginServer() LoginServerPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupConnectedRegistryResult) *LoginServerPropertiesResponse { return v.LoginServer }).(LoginServerPropertiesResponsePtrOutput)
}

func (o LookupConnectedRegistryResultOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedRegistryResult) string { return v.Mode }).(pulumi.StringOutput)
}

func (o LookupConnectedRegistryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedRegistryResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupConnectedRegistryResultOutput) Parent() ParentPropertiesResponseOutput {
	return o.ApplyT(func(v LookupConnectedRegistryResult) ParentPropertiesResponse { return v.Parent }).(ParentPropertiesResponseOutput)
}

func (o LookupConnectedRegistryResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedRegistryResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupConnectedRegistryResultOutput) StatusDetails() StatusDetailPropertiesResponseArrayOutput {
	return o.ApplyT(func(v LookupConnectedRegistryResult) []StatusDetailPropertiesResponse { return v.StatusDetails }).(StatusDetailPropertiesResponseArrayOutput)
}

func (o LookupConnectedRegistryResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupConnectedRegistryResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupConnectedRegistryResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedRegistryResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupConnectedRegistryResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedRegistryResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectedRegistryResultOutput{})
}
