// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * A Disk.
 *
 * Uses Azure REST API version 2018-09-15. In version 2.x of the Azure Native provider, it used API version 2018-09-15.
 */
export class Disk extends pulumi.CustomResource {
    /**
     * Get an existing Disk resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Disk {
        return new Disk(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:devtestlab:Disk';

    /**
     * Returns true if the given object is an instance of Disk.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Disk {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Disk.__pulumiType;
    }

    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * The creation date of the disk.
     */
    public /*out*/ readonly createdDate!: pulumi.Output<string>;
    /**
     * When backed by a blob, the name of the VHD blob without extension.
     */
    public readonly diskBlobName!: pulumi.Output<string | undefined>;
    /**
     * The size of the disk in Gibibytes.
     */
    public readonly diskSizeGiB!: pulumi.Output<number | undefined>;
    /**
     * The storage type for the disk (i.e. Standard, Premium).
     */
    public readonly diskType!: pulumi.Output<string | undefined>;
    /**
     * When backed by a blob, the URI of underlying blob.
     */
    public readonly diskUri!: pulumi.Output<string | undefined>;
    /**
     * The host caching policy of the disk (i.e. None, ReadOnly, ReadWrite).
     */
    public readonly hostCaching!: pulumi.Output<string | undefined>;
    /**
     * The resource ID of the VM to which this disk is leased.
     */
    public readonly leasedByLabVmId!: pulumi.Output<string | undefined>;
    /**
     * The location of the resource.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * When backed by managed disk, this is the ID of the compute disk resource.
     */
    public readonly managedDiskId!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The provisioning status of the resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * When backed by a blob, the storage account where the blob is.
     */
    public readonly storageAccountId!: pulumi.Output<string | undefined>;
    /**
     * The tags of the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The unique immutable identifier of a resource (Guid).
     */
    public /*out*/ readonly uniqueIdentifier!: pulumi.Output<string>;

    /**
     * Create a Disk resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DiskArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.labName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'labName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.userName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userName'");
            }
            resourceInputs["diskBlobName"] = args ? args.diskBlobName : undefined;
            resourceInputs["diskSizeGiB"] = args ? args.diskSizeGiB : undefined;
            resourceInputs["diskType"] = args ? args.diskType : undefined;
            resourceInputs["diskUri"] = args ? args.diskUri : undefined;
            resourceInputs["hostCaching"] = args ? args.hostCaching : undefined;
            resourceInputs["labName"] = args ? args.labName : undefined;
            resourceInputs["leasedByLabVmId"] = args ? args.leasedByLabVmId : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["managedDiskId"] = args ? args.managedDiskId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["storageAccountId"] = args ? args.storageAccountId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["userName"] = args ? args.userName : undefined;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["createdDate"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["uniqueIdentifier"] = undefined /*out*/;
        } else {
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["createdDate"] = undefined /*out*/;
            resourceInputs["diskBlobName"] = undefined /*out*/;
            resourceInputs["diskSizeGiB"] = undefined /*out*/;
            resourceInputs["diskType"] = undefined /*out*/;
            resourceInputs["diskUri"] = undefined /*out*/;
            resourceInputs["hostCaching"] = undefined /*out*/;
            resourceInputs["leasedByLabVmId"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["managedDiskId"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["storageAccountId"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["uniqueIdentifier"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:devtestlab/v20160515:Disk" }, { type: "azure-native:devtestlab/v20180915:Disk" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(Disk.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Disk resource.
 */
export interface DiskArgs {
    /**
     * When backed by a blob, the name of the VHD blob without extension.
     */
    diskBlobName?: pulumi.Input<string>;
    /**
     * The size of the disk in Gibibytes.
     */
    diskSizeGiB?: pulumi.Input<number>;
    /**
     * The storage type for the disk (i.e. Standard, Premium).
     */
    diskType?: pulumi.Input<string | enums.devtestlab.StorageType>;
    /**
     * When backed by a blob, the URI of underlying blob.
     */
    diskUri?: pulumi.Input<string>;
    /**
     * The host caching policy of the disk (i.e. None, ReadOnly, ReadWrite).
     */
    hostCaching?: pulumi.Input<string>;
    /**
     * The name of the lab.
     */
    labName: pulumi.Input<string>;
    /**
     * The resource ID of the VM to which this disk is leased.
     */
    leasedByLabVmId?: pulumi.Input<string>;
    /**
     * The location of the resource.
     */
    location?: pulumi.Input<string>;
    /**
     * When backed by managed disk, this is the ID of the compute disk resource.
     */
    managedDiskId?: pulumi.Input<string>;
    /**
     * The name of the Disk
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * When backed by a blob, the storage account where the blob is.
     */
    storageAccountId?: pulumi.Input<string>;
    /**
     * The tags of the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the user profile.
     */
    userName: pulumi.Input<string>;
}
