


package v20180901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupArtifactSource(ctx *pulumi.Context, args *LookupArtifactSourceArgs, opts ...pulumi.InvokeOption) (*LookupArtifactSourceResult, error) {
	var rv LookupArtifactSourceResult
	err := ctx.Invoke("azure-native:deploymentmanager/v20180901preview:getArtifactSource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupArtifactSourceArgs struct {
	ArtifactSourceName string `pulumi:"artifactSourceName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupArtifactSourceResult struct {
	ArtifactRoot   *string                   `pulumi:"artifactRoot"`
	Authentication SasAuthenticationResponse `pulumi:"authentication"`
	Id             string                    `pulumi:"id"`
	Location       string                    `pulumi:"location"`
	Name           string                    `pulumi:"name"`
	SourceType     string                    `pulumi:"sourceType"`
	Tags           map[string]string         `pulumi:"tags"`
	Type           string                    `pulumi:"type"`
}

func LookupArtifactSourceOutput(ctx *pulumi.Context, args LookupArtifactSourceOutputArgs, opts ...pulumi.InvokeOption) LookupArtifactSourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupArtifactSourceResult, error) {
			args := v.(LookupArtifactSourceArgs)
			r, err := LookupArtifactSource(ctx, &args, opts...)
			var s LookupArtifactSourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupArtifactSourceResultOutput)
}

type LookupArtifactSourceOutputArgs struct {
	ArtifactSourceName pulumi.StringInput `pulumi:"artifactSourceName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupArtifactSourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupArtifactSourceArgs)(nil)).Elem()
}


type LookupArtifactSourceResultOutput struct{ *pulumi.OutputState }

func (LookupArtifactSourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupArtifactSourceResult)(nil)).Elem()
}

func (o LookupArtifactSourceResultOutput) ToLookupArtifactSourceResultOutput() LookupArtifactSourceResultOutput {
	return o
}

func (o LookupArtifactSourceResultOutput) ToLookupArtifactSourceResultOutputWithContext(ctx context.Context) LookupArtifactSourceResultOutput {
	return o
}

func (o LookupArtifactSourceResultOutput) ArtifactRoot() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArtifactSourceResult) *string { return v.ArtifactRoot }).(pulumi.StringPtrOutput)
}

func (o LookupArtifactSourceResultOutput) Authentication() SasAuthenticationResponseOutput {
	return o.ApplyT(func(v LookupArtifactSourceResult) SasAuthenticationResponse { return v.Authentication }).(SasAuthenticationResponseOutput)
}

func (o LookupArtifactSourceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArtifactSourceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupArtifactSourceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArtifactSourceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupArtifactSourceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArtifactSourceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupArtifactSourceResultOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArtifactSourceResult) string { return v.SourceType }).(pulumi.StringOutput)
}

func (o LookupArtifactSourceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupArtifactSourceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupArtifactSourceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArtifactSourceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupArtifactSourceResultOutput{})
}
