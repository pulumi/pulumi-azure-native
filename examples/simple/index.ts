import * as pulumi from "@pulumi/pulumi";

class MyRes extends pulumi.CustomResource {
    constructor(name: string, opts?: pulumi.CustomResourceOptions) {
        super("azurerm:foo:MyRes", name, {}, opts);
    }
}

const x = new MyRes("abc");