package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/media/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := media.NewMediaGraph(ctx, "mediaGraph", &media.MediaGraphArgs{
			AccountName:       pulumi.String("contosomedia"),
			Description:       pulumi.String("updated description"),
			MediaGraphName:    pulumi.String("SampleMediaGraph"),
			ResourceGroupName: pulumi.String("contoso"),
			Sinks: []media.MediaGraphAssetSinkArgs{
				{
					AssetName: pulumi.String("SampleAsset"),
					Inputs: pulumi.StringArray{
						pulumi.String("rtspSource"),
					},
					Name:      pulumi.String("AssetSink"),
					OdataType: pulumi.String("#Microsoft.Media.MediaGraphAssetSink"),
				},
			},
			Sources: []media.MediaGraphRtspSourceArgs{
				{
					Endpoint: {
						Credentials: {
							OdataType: "#Microsoft.Media.MediaGraphUsernamePasswordCredentials",
							Password:  "examplepassword",
							Username:  "exampleusername",
						},
						OdataType: "#Microsoft.Media.MediaGraphTlsEndpoint",
						TrustedCertificates: {
							Certificates: []string{
								`-----BEGIN CERTIFICATE-----
MIIDhTCCAm2gAwIBAgIUajvPKmoO+8qaO89/ZGATl7ZYnTswDQYJKoZIhvcNAQEL
BQAwUTESMBAGA1UECgwJTWljcm9zb2Z0MRQwEgYDVQQLDAtBenVyZSBNZWRpYTEl
MCMGA1UEAwwcKFVudHJ1c3RlZCkgVGVzdCBDZXJ0aWZpY2F0ZTAgFw0yMDAyMDYy
MTI5MTlaGA8zMDE5MDYwOTIxMjkxOVowUTESMBAGA1UECgwJTWljcm9zb2Z0MRQw
EgYDVQQLDAtBenVyZSBNZWRpYTElMCMGA1UEAwwcKFVudHJ1c3RlZCkgVGVzdCBD
ZXJ0aWZpY2F0ZTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAK2lg5ff
7xXPaBZXHl/zrTukdiBtu7BNIOchHba51eloruPRzpvQx7Pedk3CVTut7LYinijf
uol0EwkQ2FLt2i2jOqiva9nXR95ujIZHcKsEeMC4RSNSP4++k6SpP8FgyYVdv5ru
f8GC+HyYQ4j0TqpR/cJs53l/LGRSldaFZ6fcDde1jeyca4VivAbAH1/WDIOvmjzo
9XIGxZ10VSS5l5+DIgdkJZ+mDMLJIuVZ0YVF16ZGEB3beq1trk5lItvmSjQLTllH
qMFm9UGY8jKZSo/BY8ewHEtnGSAFQK0TVuRx1HhUWwu6C9jk+2zmRS2090BNpQWa
JMKFJrSPzFDPRX8CAwEAAaNTMFEwHQYDVR0OBBYEFIumbhu0lYk0EFDThEg0yyIn
/wZZMB8GA1UdIwQYMBaAFIumbhu0lYk0EFDThEg0yyIn/wZZMA8GA1UdEwEB/wQF
MAMBAf8wDQYJKoZIhvcNAQELBQADggEBADUNw+/NGNVtigq9tMJKqlk39MTpDn1s
Z1BVIAuAWSQjlevYZJeDIPUiWNWFhRe+xN7oOLnn2+NIXEKKeMSyuPoZYbN0mBkB
99oS3XVipSANpmDvIepNdCrOnjfqDFIifRF1Dqjtb6i1hb6v/qYKVPLQvcrgGur7
PKKkAu9p4YRZ3RBdwwaUuMgojrj/l6DGbeJY6IRVnVMY39rryMnZjA5xUlhCu55n
oB3t/jsJLwnQN+JbAjLAeuqgOWtgARsEFzvpt+VvDsaj0YLOJPhyJwTvHgaa/slB
nECzd3TuyFKYeGssSni/QQ1e7yZcLapQqz66g5otdriw0IRdOfDxm5M=
-----END CERTIFICATE-----`,
							},
							OdataType: "#Microsoft.Media.MediaGraphPemCertificateList",
						},
						Url: "rtsps://contoso.com:443/stream1",
						ValidationOptions: {
							IgnoreHostname:  true,
							IgnoreSignature: false,
						},
					},
					Name:      pulumi.String("rtspSource"),
					OdataType: pulumi.String("#Microsoft.Media.MediaGraphRtspSource"),
					Transport: pulumi.String("Http"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
