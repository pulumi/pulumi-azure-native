import * as pulumi from "@pulumi/pulumi";
import { networkInterfaces } from "os";

class ContainerGroup extends pulumi.CustomResource {
    constructor(name: string, args: any, opts?: pulumi.CustomResourceOptions) {
        super("azurerm:containerinstance:ContainerGroup", name, args, opts);
    }
}

class VirtualMachine extends pulumi.CustomResource {
    constructor(name: string, args: any, opts?: pulumi.CustomResourceOptions) {
        super("azurerm:compute:VirtualMachine", name, args, opts);
    }
}

class NetworkInterface extends pulumi.CustomResource {
    id: pulumi.Output<string>;
    constructor(name: string, args: any, opts?: pulumi.CustomResourceOptions) {
        super("azurerm:network:NetworkInterface", name, args, opts);
    }
}


class VirtualNetwork extends pulumi.CustomResource {
    id: pulumi.Output<string>;
    constructor(name: string, args: any, opts?: pulumi.CustomResourceOptions) {
        super("azurerm:network:VirtualNetwork", name, args, opts);
    }
}


const containerinstance = new ContainerGroup("abc", {
    resourceGroupName: "azuretest",
    // should be autonamed?
    containerGroupName: "abc-1234", 
    // should be inlined via use of `"in": "body"`?
    containerGroup: { 
        location: "westus2", 
        // should be inlined via 'x-ms-client-flatten'
        properties: { 
            osType: "Linux",
            containers: [{
                name: "foo",
                // should be inlined via 'x-ms-client-flatten'
                properties: { 
                    image: "nginx",
                    resources: {
                        requests: {
                            memoryInGB: "1",
                            cpu: "1",
                        },
                    },
                },
            }],
        },
    },
});

const vnet = new VirtualNetwork("vnet", {
    resourceGroupName: "azuretest",
    virtualNetworkName: "vnet-1234",
    parameters: {
        location: "westus2",
        properties: {
            addressSpace: {
                addressPrefixes: ["10.1.0.0/24"],
            },
            subnets: [{
                name: "default",
                properties: {
                    addressPrefix: "10.1.0.0/24",
                },
            }],
        },
    },
});

const networkInterface = new NetworkInterface("nic", {
    resourceGroupName: "azuretest",
    networkInterfaceName: "nic-1234",
    parameters: {
        location: "westus2",
        properties: {
            ipConfigurations: [{
                name: "ipconfig1",
                properties: {
                    subnet: {
                        id: pulumi.interpolate`${vnet.id}/subnets/default`,
                    },
                    privateIPAllocationMethod: "Dynamic",
                },
            }],
        },
    },
});

const virtualmachine  = new VirtualMachine("vm", {
    resourceGroupName: "azuretest",
    vmName: "abc-1234", 
    parameters: {
        location: "westus2",
        properties: {  
            hardwareProfile: {
                vmSize: "Standard_A0",
            },
            networkProfile: {
                networkInterfaces: [{
                    id: networkInterface.id,
                }],
            },
            storageProfile: {
                imageReference: {
                    publisher: "Canonical",
                    offer: "UbuntuServer",
                    sku: "18.04-LTS",
                    version: "latest"
                }
            },
            osProfile: {
                computerName: "foo",
                adminUsername: "lukehoban",
                linuxConfiguration: {
                    disablePasswordAuthentication: true,
                    ssh: {
                        publicKeys: [{
                            path: "/home/lukehoban/.ssh/authorized_keys",
                            keyData: "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCaMWLUeL02T3k5E+tr+0BwmHyzkvEWkOuC+Y8c4o5FY+WuCZqzK4pyHrhWqtVOQEnsMOjIRzPVjX3URcJxuCZnEC8Kblru+tIzf1xkpexxzYUPF3BgLLSCFp0hm28BVkmktkhaPzAWuCBsYIJY6t2SHxT2BbGrsXlmKItUP78ViZaMdKWAAToNPuvJnSV1XAKI0tet6bzMN2ZOWDByrXWi1AjMHuJwHWDAOWTKRcO1MgqNbIPI1mPGIbIwJ0bHyaJLIGqCJLWd+g9VOA4D/T3qahzO3y1xkJB315J/QyCVS3Cdbt4kxQIsCmYsAaxB4/uM8lNGShP6H2p72n4CY9Dov7Fh2je4jBGVT/1873f3xEE24KtIiulebusIIOUH1T+TgUo2mPU+wlO1jQaEbp8bHK/216/dXIzr67+4nmCasvIiI0ZbspQ0Yz/sSUoDDV3pF9WnG8Y11thjfu62TEM66iNaS0NgTHJHUtHs0/jezTIAW/yJ9a7i113xZAtnq3KHi6AzG08HlQyGLqsF1ny80+aT/KYcipXMrsBIO6p1zL4JGSnFyEQuc2SdQ/DTOKF2Tz55KXdrkQ5jeQNWcNSW5BfoWhVfa1/cIfxNBEi2ZqScmDwkienkd3CZG1ci9UJQyrtfhkO1c8mN8BeVhD2rljrk+y5cO3Tw9H1hTYkFWQ== luke@pulumi.com"
                        }],
                    },
                },
            },
        },
    },
});