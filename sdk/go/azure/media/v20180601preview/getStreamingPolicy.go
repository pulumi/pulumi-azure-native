


package v20180601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStreamingPolicy(ctx *pulumi.Context, args *LookupStreamingPolicyArgs, opts ...pulumi.InvokeOption) (*LookupStreamingPolicyResult, error) {
	var rv LookupStreamingPolicyResult
	err := ctx.Invoke("azure-native:media/v20180601preview:getStreamingPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStreamingPolicyArgs struct {
	AccountName         string `pulumi:"accountName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	StreamingPolicyName string `pulumi:"streamingPolicyName"`
}


type LookupStreamingPolicyResult struct {
	CommonEncryptionCbcs        *CommonEncryptionCbcsResponse `pulumi:"commonEncryptionCbcs"`
	CommonEncryptionCenc        *CommonEncryptionCencResponse `pulumi:"commonEncryptionCenc"`
	Created                     string                        `pulumi:"created"`
	DefaultContentKeyPolicyName *string                       `pulumi:"defaultContentKeyPolicyName"`
	EnvelopeEncryption          *EnvelopeEncryptionResponse   `pulumi:"envelopeEncryption"`
	Id                          string                        `pulumi:"id"`
	Name                        string                        `pulumi:"name"`
	NoEncryption                *NoEncryptionResponse         `pulumi:"noEncryption"`
	Type                        string                        `pulumi:"type"`
}
