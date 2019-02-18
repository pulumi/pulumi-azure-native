import * as pulumi from "@pulumi/pulumi";

class MyRes extends pulumi.CustomResource {
    constructor(name: string, args: any, opts?: pulumi.CustomResourceOptions) {
        super("azurerm:foo:MyRes", name, args, opts);
    }
}

const x = new MyRes("abc", {
    location: "westus2",
    properties: {
        osType: "Linux",
        containers: [{
            name: "foo",
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
});