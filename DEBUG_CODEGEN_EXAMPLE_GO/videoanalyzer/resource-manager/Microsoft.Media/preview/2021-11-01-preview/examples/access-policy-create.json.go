package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/videoanalyzer/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := videoanalyzer.NewAccessPolicy(ctx, "accessPolicy", &videoanalyzer.AccessPolicyArgs{
AccessPolicyName: pulumi.String("accessPolicyName1"),
AccountName: pulumi.String("testaccount2"),
Authentication: interface{}{
Audiences: pulumi.StringArray{
pulumi.String("audience1"),
},
Claims: videoanalyzer.TokenClaimArray{
&videoanalyzer.TokenClaimArgs{
Name: pulumi.String("claimname1"),
Value: pulumi.String("claimvalue1"),
},
&videoanalyzer.TokenClaimArgs{
Name: pulumi.String("claimname2"),
Value: pulumi.String("claimvalue2"),
},
},
Issuers: pulumi.StringArray{
pulumi.String("issuer1"),
pulumi.String("issuer2"),
},
Keys: pulumi.Array{
videoanalyzer.RsaTokenKey{
Alg: "RS256",
E: "ZLFzZTY0IQ==",
Kid: "123",
N: "YmFzZTY0IQ==",
Type: "#Microsoft.VideoAnalyzer.RsaTokenKey",
},
videoanalyzer.EccTokenKey{
Alg: "ES256",
Kid: "124",
Type: "#Microsoft.VideoAnalyzer.EccTokenKey",
X: "XX==",
Y: "YY==",
},
},
Type: pulumi.String("#Microsoft.VideoAnalyzer.JwtAuthentication"),
},
ResourceGroupName: pulumi.String("testrg"),
})
if err != nil {
return err
}
return nil
})
}
