import * as pulumi from "@pulumi/pulumi";
import * as random from "@pulumi/random";
import * as fs from "fs";
import * as resources from "@pulumi/azure-native/resources";

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
    resourceGroupName: pulumi.Output<string>;
    fileName: string;
}

class Template extends pulumi.ComponentResource {
    constructor(name: string, args: TemplateArgs, opts?: pulumi.ComponentResourceOptions) {
        super("example:example:Template", name, {}, opts);

        const templateContent = fs.readFileSync(`${__dirname}/${args.fileName}`).toString();
        const template = JSON.parse(templateContent);

        const dependsOn: pulumi.Resource[] = [];

        // This is a very rough approximation of flattening. The real converter would have to drive off the schema.
        function flatten(args: any): any {
            if (Array.isArray(args) && typeof args !== "string") {
                return args.map(flatten);
            }  else if (typeof args === "object") {
                const result: {[key:string]: any} = {};
                for (const [key, value] of Object.entries(args)) {
                    const props = flatten(value);
                    if (key === "properties") {
                        for (const [key, value] of Object.entries(props)) {
                            result[key] = value;
                        }
                    } else {
                        result[key] = props;
                    }
                }
                return result;
            } else {
                return args;
            }
        }

        for (let resource of template.resources) {
            const resourceArgs = {
                resourceGroupName: args.resourceGroupName,
                ...flatten(resource)
            };
            const childOpts = { ...opts, parent: this, dependsOn: [...dependsOn] };
            const version = "v" + resource.apiVersion.replace("-", "").replace("-", "");

            let typeName;
            if (resource.type === "Microsoft.Compute/virtualMachines") {
                typeName = `azure-native:compute/${version}:VirtualMachine`;
                resourceArgs.vmName = resource.name;
            } else if (resource.type === "Microsoft.Network/networkInterfaces") {
                typeName = `azure-native:network/${version}:NetworkInterface`;
                resourceArgs.networkInterfaceName = resource.name;
            } else if (resource.type === "Microsoft.Network/networkSecurityGroups") {
                typeName = `azure-native:network/${version}:NetworkSecurityGroup`;
                resourceArgs.networkSecurityGroupName = resource.name;
            } else if (resource.type === "Microsoft.Network/publicIPAddresses") {
                typeName = `azure-native:network/${version}:PublicIPAddress`;
                resourceArgs.publicIpAddressName = resource.name;
            } else if (resource.type === "Microsoft.Network/virtualNetworks") {
                typeName = `azure-native:network/${version}:VirtualNetwork`;
                resourceArgs.virtualNetworkName = resource.name;
            } else if (resource.type === "Microsoft.Storage/storageAccounts") {
                typeName = `azure-native:storage/${version}:StorageAccount`;
                resourceArgs.accountName = resource.name;
            } else if (resource.type === "Microsoft.Web/serverfarms") {
                typeName = `azure-native:web/${version}:ServerFarm`;
                resourceArgs.name = resource.name;
            } else if (resource.type === "Microsoft.Web/sites") {
                typeName = `azure-native:web/${version}:Site`;
                resourceArgs.name = resource.name;
            } else if (resource.type === "Microsoft.Cache/Redis") {
                typeName = `azure-native:cache/${version}:Redis`;
                resourceArgs.name = resource.name;                
            } else if (resource.type === "Microsoft.ContainerInstance/containerGroups") {
                typeName = `azure-native:containerinstance/${version}:ContainerGroup`;
                resourceArgs.containerGroupName = resource.name;                
            } else {
                throw new Error(`Unknown type ${resource.type}`);
            }

            dependsOn.push(new TemplateResource(typeName, resourceArgs, childOpts));
        }
    }
}

const randomString = new random.RandomString("random", {
    length: 12,
    special: false,
    upper: false,
});
const resourceGroupName = randomString.result;

const resourceGroup = new resources.ResourceGroup("azure-native", {
    resourceGroupName,
});

new Template("storageaccount", {
    resourceGroupName,
    fileName: "storageaccount.json",
}, { dependsOn: [resourceGroup] });

new Template("windowsvm", {
    resourceGroupName,
    fileName: "windowsvm.json",
}, { dependsOn: [resourceGroup] });

// Takes ~20 minutes to deploy Redis and may fail with timeout https://github.com/pulumi/pulumi-azure-native-provider/issues/17
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
