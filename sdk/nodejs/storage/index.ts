import * as pulumi from "@pulumi/pulumi";

export class StorageAccount extends pulumi.CustomResource {
    public readonly properties: pulumi.Output<StorageAccountProperties>;
    public readonly accountName: pulumi.Output<string>;
    constructor(name: string, args: any, opts?: pulumi.CustomResourceOptions) {
        super("azurerm:storage:StorageAccount", name, args, opts);
    }
}

export class StorageAccountProperties {
    public readonly primaryEndpoints: pulumi.Output<StorageAccountPropertiesPrimaryEndpoints>;
}

export class StorageAccountPropertiesPrimaryEndpoints {
    public readonly web: pulumi.Output<string>;
}

export class StorageAccountBlobServiceContainer extends pulumi.CustomResource {
    constructor(name: string, args: any, opts?: pulumi.CustomResourceOptions) {
        super("azurerm:storage:StorageAccountBlobServiceContainer", name, args, opts);
    }
}
