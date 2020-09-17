package arm2pulumi

import (
	"encoding/json"
	"fmt"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/pulumi/pulumi-azure-nextgen/provider/pkg/provider"
	"github.com/pulumi/pulumi/pkg/v2/codegen/dotnet"
	gogen "github.com/pulumi/pulumi/pkg/v2/codegen/go"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/nodejs"
	"github.com/pulumi/pulumi/pkg/v2/codegen/python"
	"github.com/sourcegraph/jsonx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRenderTemplate(t *testing.T) {
	var metadata provider.AzureAPIMetadata
	f, err := os.Open("../../cmd/pulumi-resource-azure-nextgen/metadata.json")
	os.Setenv("PATH", "../../../bin")
	require.NoError(t, err)
	require.NoError(t, json.NewDecoder(f).Decode(&metadata))
	f.Close()

	var pkgSpec schema.PackageSpec
	f, err = os.Open("../../cmd/pulumi-resource-azure-nextgen/schema-full.json")
	require.NoError(t, err)
	require.NoError(t, json.NewDecoder(f).Decode(&pkgSpec))
	f.Close()

	for _, test := range []testcase{
		aksCluster,
	} {
		t.Run(test.name, func(t *testing.T) {
			node, err := parseJsonxTree(test.template)
			require.NoError(t, err)
			body, _, err := renderTemplate(map[string]*jsonx.Node{
				"example.json": node,
			}, &metadata, &pkgSpec)
			if test.err == nil {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
				assert.EqualError(t, err, test.err.Error())
			}
			programBody := fmt.Sprintf("%v", body)
			fmt.Println(programBody)

			parser := syntax.NewParser()
			require.NoError(t, parser.ParseFile(strings.NewReader(programBody), "program.pp"))

			if parser.Diagnostics.HasErrors() {
				fmt.Print(programBody)
				parser.NewDiagnosticWriter(os.Stderr, 0, true).WriteDiagnostics(parser.Diagnostics)
				t.Fail()
			}

			program, diags, err := hcl2.BindProgram(parser.Files, hcl2.Cache(hcl2.NewPackageCache()))
			if err != nil {
				log.Fatalf("failed to bind program: %v", err)
			}
			if diags.HasErrors() {
				log.Printf(programBody)
				program.NewDiagnosticWriter(os.Stderr, 0, true).WriteDiagnostics(diags)
				t.Fail()
			}

			languages := []string{"nodejs"}
			languageExample := map[string]string{}
			for _, lang := range languages {
				var files map[string][]byte

				switch lang {
				case "dotnet":
					files, diags, err = dotnet.GenerateProgram(program)
				case "go":
					files, diags, err = gogen.GenerateProgram(program)
				case "nodejs":
					files, diags, err = nodejs.GenerateProgram(program)
				case "python":
					files, diags, err = python.GenerateProgram(program)
				case "schema":
					continue
				}
				if err != nil {
					log.Fatalf("failed to generate program: %v", err)
				}
				if diags.HasErrors() {
					log.Printf(programBody)
					program.NewDiagnosticWriter(os.Stderr, 0, true).WriteDiagnostics(diags)
					os.Exit(-1)
				}

				buf := strings.Builder{}
				for _, f := range files {
					_, err := buf.Write(f)
					require.NoError(t, err)
				}
				languageExample[lang] = buf.String()
				fmt.Printf("%s\n", buf.String())
				assert.Equal(t, test.expected, buf.String())
			}
		})

	}

}

type testcase struct {
	name     string
	template string
	expected string
	err      error
}

var (
	aksCluster = testcase{
		name: "aksCluster",
		template: `{
    "$schema": "https://schema.management.azure.com/schemas/2019-04-01/deploymentTemplate.json#",
    "contentVersion": "1.0.0.1",
    "parameters": {
        "clusterName": {
            "type": "string",
            "defaultValue":"aks101cluster",
            "metadata": {
                "description": "The name of the Managed Cluster resource."
            }
        },
        /* Comments!!! */
        "location": {
            "type": "string", /* More comments */
            "defaultValue": "[resourceGroup().location]", // other comments
            "metadata": {
                "description": "The location of the Managed Cluster resource."
            }
        },
        "dnsPrefix": {
            "type": "string",
            "metadata": {
                "description": "Optional DNS prefix to use with hosted Kubernetes API server FQDN."
            }
        },
        "osDiskSizeGB": {
            "type": "int",
            "defaultValue": 0,
            "metadata": {
                "description": "Disk size (in GB) to provision for each of the agent pool nodes. This value ranges from 0 to 1023. Specifying 0 will apply the default disk size for that agentVMSize."
            },
            "minValue": 0,
            "maxValue": 1023
        },
        "agentCount": {
            "type": "int",
            "defaultValue": 3,
            "metadata": {
                "description": "The number of nodes for the cluster."
            },
            "minValue": 1,
            "maxValue": 50
        },
        "agentVMSize": {
            "type": "string",
            "defaultValue": "Standard_DS2_v2",
            "metadata": {
                "description": "The size of the Virtual Machine."
            }
        },
        "linuxAdminUsername": {
            "type": "string",
            "metadata": {
                "description": "User name for the Linux Virtual Machines."
            }
        },
        "sshRSAPublicKey": {
            "type": "string",
            "metadata": {
                "description": "Configure all linux machines with the SSH RSA public key string. Your key should include three parts, for example 'ssh-rsa AAAAB...snip...UcyupgH azureuser@linuxvm'"
            }
        },
        "servicePrincipalClientId": {
            "metadata": {
                "description": "Client ID (used by cloudprovider)"
            },
            "type": "securestring"
        },
        "servicePrincipalClientSecret": {
            "metadata": {
                "description": "The Service Principal Client Secret."
            },
            "type": "securestring"
        },
        "osType": {
            "type": "string",
            "defaultValue": "Linux",
            "allowedValues": [
                "Linux"
            ],
            "metadata": {
                "description": "The type of operating system."
            }
        }        
    },
    "resources": [
        {
            "apiVersion": "2020-03-01",
            "type": "Microsoft.ContainerService/managedClusters",
            "location": "[parameters('location')]",
            "name": "[parameters('clusterName')]",
            "properties": {
                "dnsPrefix": "[parameters('dnsPrefix')]",
                "agentPoolProfiles": [
                    {
                        "name": "agentpool",
                        "osDiskSizeGB": "[parameters('osDiskSizeGB')]",
                        "count": "[parameters('agentCount')]",
                        "vmSize": "[parameters('agentVMSize')]",
                        "osType": "[parameters('osType')]",
                        "storageProfile": "ManagedDisks"
                    }
                ],
                "linuxProfile": {
                    "adminUsername": "[parameters('linuxAdminUsername')]",
                    "ssh": {
                        "publicKeys": [
                            {
                                "keyData": "[parameters('sshRSAPublicKey')]"
                            }
                        ]
                    }
                },
                "servicePrincipalProfile": {
                    "clientId": "[parameters('servicePrincipalClientId')]",
                    "Secret": "[parameters('servicePrincipalClientSecret')]"
                }
            }
        }
    ],
    "outputs": {
        "controlPlaneFQDN": {
            "type": "string",
            "value": "[reference(parameters('clusterName')).fqdn]"
        }
    }
}`,
		expected: `import * as pulumi from "@pulumi/pulumi";
import * as azure_nextgen from "@pulumi/azure-nextgen";

const config = new pulumi.Config();
const agentCountParam = config.getNumber("agentCountParam") || "3";
const agentVMSizeParam = config.get("agentVMSizeParam") || "Standard_DS2_v2";
const clusterNameParam = config.get("clusterNameParam") || "aks101cluster";
const dnsPrefixParam = config.require("dnsPrefixParam");
const linuxAdminUsernameParam = config.require("linuxAdminUsernameParam");
const resourceGroupNameParam = config.require("resourceGroupNameParam");
const resourceGroupVar = azure_nextgen.resources.latest.getResourceGroup({
    resourceGroupName: resourceGroupNameParam,
});
const locationParam = config.get("locationParam") || resourceGroupVar.then(resourceGroupVar => resourceGroupVar.location);
const osDiskSizeGBParam = config.getNumber("osDiskSizeGBParam") || "0";
const osTypeParam = config.get("osTypeParam") || "Linux";
const servicePrincipalClientIdParam = config.require("servicePrincipalClientIdParam");
const servicePrincipalClientSecretParam = config.require("servicePrincipalClientSecretParam");
const sshRSAPublicKeyParam = config.require("sshRSAPublicKeyParam");
const managedClusterResource = new azure_nextgen.containerservice.v20200301.ManagedCluster("managedClusterResource", {
    agentPoolProfiles: [{
        count: agentCountParam,
        name: "agentpool",
        osDiskSizeGB: osDiskSizeGBParam,
        osType: osTypeParam,
        vmSize: agentVMSizeParam,
    }],
    dnsPrefix: dnsPrefixParam,
    linuxProfile: {
        adminUsername: linuxAdminUsernameParam,
        ssh: {
            publicKeys: [{
                keyData: sshRSAPublicKeyParam,
            }],
        },
    },
    location: locationParam,
    resourceName: clusterNameParam,
    servicePrincipalProfile: {
        clientId: servicePrincipalClientIdParam,
    },
});
export const controlPlaneFQDNOut = managedClusterResource.fqdn;
`,
	}
)
