import * as pulumi from "@pulumi/pulumi";
import * as fs from "fs";
import * as azurerm from "../../sdk/nodejs";

interface TemplateResourceArgs {
    resourceGroupName: string;
    [key: string]: any;
}

class TemplateResource extends pulumi.CustomResource {
    constructor(typeName: string, args: TemplateResourceArgs, opts?: pulumi.CustomResourceOptions) {
        super(typeName, args.name, args, opts);
    }
}

interface TemplateArgs {
    resourceGroupName: string;
    fileName: string;
}

class Template extends pulumi.ComponentResource {
    constructor(name: string, args: TemplateArgs, opts?: pulumi.ComponentResourceOptions) {
        super("example:example:Template", name, {}, opts);

        const templateContent = fs.readFileSync(`${__dirname}/${args.fileName}`).toString();
        const template = JSON.parse(templateContent);

        const dependsOn: pulumi.Resource[] = [];

        for (let resource of template.resources) {
            const resourceArgs = {
                resourceGroupName: args.resourceGroupName,
                ...resource
            };
            const childOpts = { ...opts, parent: this, dependsOn: [...dependsOn] };
            const version = "v" + resource.apiVersion.replace("-", "").replace("-", "");

            let typeName;
            if (resource.type === "Microsoft.Compute/virtualMachines") {
                typeName = `azurerm:compute/${version}:VirtualMachine`;
            } else if (resource.type === "Microsoft.Network/networkInterfaces") {
                typeName = `azurerm:network/${version}:NetworkInterface`;
            } else if (resource.type === "Microsoft.Network/networkSecurityGroups") {
                typeName = `azurerm:network/${version}:NetworkSecurityGroup`;
            } else if (resource.type === "Microsoft.Network/publicIPAddresses") {
                typeName = `azurerm:network/${version}:PublicIPAddress`;
            } else if (resource.type === "Microsoft.Network/virtualNetworks") {
                typeName = `azurerm:network/${version}:VirtualNetwork`;
            } else if (resource.type === "Microsoft.Storage/storageAccounts") {
                typeName = `azurerm:storage/${version}:StorageAccount`;
            } else if (resource.type === "Microsoft.Web/serverfarms") {
                typeName = `azurerm:web/${version}:ServerFarm`;
            } else if (resource.type === "Microsoft.Web/sites") {
                typeName = `azurerm:web/${version}:Site`;
            } else if (resource.type === "Microsoft.Cache/Redis") {
                typeName = `azurerm:cache/${version}:Redis`;
            } else if (resource.type === "Microsoft.ContainerInstance/containerGroups") {
                typeName = `azurerm:containerinstance/${version}:ContainerGroup`;
            } else {
                throw new Error(`Unknown type ${resource.type}`);
            }

            dependsOn.push(new TemplateResource(typeName, resourceArgs, childOpts));
        }
    }
}

const resourceGroupName = "azurermtemplates";

const resourceGroup = new azurerm.resources.v20200601.ResourceGroup("azurerm", {
    name: resourceGroupName,
    location: "westus2",
    tags: {
        Owner: "mikhailshilkov",
    },
});

new Template("storageaccount", {
    resourceGroupName,
    fileName: "storageaccount.json",
}, { dependsOn: [resourceGroup] });

new Template("windowsvm", {
    resourceGroupName,
    fileName: "windowsvm.json",
}, { dependsOn: [resourceGroup] });

// Takes ~20 minutes to deploy Redis and may fail with timeout https://github.com/pulumi/pulumi-azurerm/issues/17
new Template("webapp", {
    resourceGroupName,
    fileName: "webapp.json",
}, { dependsOn: [resourceGroup] });

// Requires 2 passes to deploy:
// 1. Remove container groups to deploy just the storage account.
// 2. Copy the account key to container's env variables and run again.
// A proper fix would involve setting containers' inputs to SA's output.
new Template("aci", {
    resourceGroupName,
    fileName: "aci.json",
}, { dependsOn: [resourceGroup] });
