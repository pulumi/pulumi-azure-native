import * as pulumi from "@pulumi/pulumi";

class ContainerGroup extends pulumi.CustomResource {
    constructor(name: string, args: any, opts?: pulumi.CustomResourceOptions) {
        super("azurerm:containerinstance:ContainerGroup", name, args, opts);
    }
}

const x = new ContainerGroup("abc", {
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