


package blueprint

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupArtifact(ctx *pulumi.Context, args *LookupArtifactArgs, opts ...pulumi.InvokeOption) (*LookupArtifactResult, error) {
	var rv LookupArtifactResult
	err := ctx.Invoke("azure-native:blueprint:getArtifact", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupArtifactArgs struct {
	ArtifactName  string `pulumi:"artifactName"`
	BlueprintName string `pulumi:"blueprintName"`
	ResourceScope string `pulumi:"resourceScope"`
}


type LookupArtifactResult struct {
	Id   string `pulumi:"id"`
	Kind string `pulumi:"kind"`
	Name string `pulumi:"name"`
	Type string `pulumi:"type"`
}

func LookupArtifactOutput(ctx *pulumi.Context, args LookupArtifactOutputArgs, opts ...pulumi.InvokeOption) LookupArtifactResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupArtifactResult, error) {
			args := v.(LookupArtifactArgs)
			r, err := LookupArtifact(ctx, &args, opts...)
			var s LookupArtifactResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupArtifactResultOutput)
}

type LookupArtifactOutputArgs struct {
	ArtifactName  pulumi.StringInput `pulumi:"artifactName"`
	BlueprintName pulumi.StringInput `pulumi:"blueprintName"`
	ResourceScope pulumi.StringInput `pulumi:"resourceScope"`
}

func (LookupArtifactOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupArtifactArgs)(nil)).Elem()
}


type LookupArtifactResultOutput struct{ *pulumi.OutputState }

func (LookupArtifactResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupArtifactResult)(nil)).Elem()
}

func (o LookupArtifactResultOutput) ToLookupArtifactResultOutput() LookupArtifactResultOutput {
	return o
}

func (o LookupArtifactResultOutput) ToLookupArtifactResultOutputWithContext(ctx context.Context) LookupArtifactResultOutput {
	return o
}

func (o LookupArtifactResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArtifactResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupArtifactResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArtifactResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupArtifactResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArtifactResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupArtifactResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArtifactResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupArtifactResultOutput{})
}
