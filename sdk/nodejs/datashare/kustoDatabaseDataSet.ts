// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * A kusto database data set.
 *
 * Uses Azure REST API version 2021-08-01. In version 2.x of the Azure Native provider, it used API version 2021-08-01.
 */
export class KustoDatabaseDataSet extends pulumi.CustomResource {
    /**
     * Get an existing KustoDatabaseDataSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): KustoDatabaseDataSet {
        return new KustoDatabaseDataSet(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:datashare:KustoDatabaseDataSet';

    /**
     * Returns true if the given object is an instance of KustoDatabaseDataSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is KustoDatabaseDataSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === KustoDatabaseDataSet.__pulumiType;
    }

    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * Unique id for identifying a data set resource
     */
    public /*out*/ readonly dataSetId!: pulumi.Output<string>;
    /**
     * Kind of data set.
     * Expected value is 'KustoDatabase'.
     */
    public readonly kind!: pulumi.Output<"KustoDatabase">;
    /**
     * Resource id of the kusto database.
     */
    public readonly kustoDatabaseResourceId!: pulumi.Output<string>;
    /**
     * Location of the kusto cluster.
     */
    public /*out*/ readonly location!: pulumi.Output<string>;
    /**
     * Name of the azure resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Provisioning state of the kusto database data set.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * System Data of the Azure resource.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.datashare.SystemDataResponse>;
    /**
     * Type of the azure resource
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a KustoDatabaseDataSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KustoDatabaseDataSetArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.accountName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountName'");
            }
            if ((!args || args.kind === undefined) && !opts.urn) {
                throw new Error("Missing required property 'kind'");
            }
            if ((!args || args.kustoDatabaseResourceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'kustoDatabaseResourceId'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.shareName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'shareName'");
            }
            resourceInputs["accountName"] = args ? args.accountName : undefined;
            resourceInputs["dataSetName"] = args ? args.dataSetName : undefined;
            resourceInputs["kind"] = "KustoDatabase";
            resourceInputs["kustoDatabaseResourceId"] = args ? args.kustoDatabaseResourceId : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["shareName"] = args ? args.shareName : undefined;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["dataSetId"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["dataSetId"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["kustoDatabaseResourceId"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:datashare/v20181101preview:KustoDatabaseDataSet" }, { type: "azure-native:datashare/v20191101:KustoDatabaseDataSet" }, { type: "azure-native:datashare/v20200901:KustoDatabaseDataSet" }, { type: "azure-native:datashare/v20201001preview:ADLSGen2StorageAccountDataSet" }, { type: "azure-native:datashare/v20201001preview:BlobStorageAccountDataSet" }, { type: "azure-native:datashare/v20201001preview:KustoDatabaseDataSet" }, { type: "azure-native:datashare/v20210801:ADLSGen1FileDataSet" }, { type: "azure-native:datashare/v20210801:ADLSGen1FolderDataSet" }, { type: "azure-native:datashare/v20210801:ADLSGen2FileDataSet" }, { type: "azure-native:datashare/v20210801:ADLSGen2FileSystemDataSet" }, { type: "azure-native:datashare/v20210801:ADLSGen2FolderDataSet" }, { type: "azure-native:datashare/v20210801:BlobContainerDataSet" }, { type: "azure-native:datashare/v20210801:BlobDataSet" }, { type: "azure-native:datashare/v20210801:BlobFolderDataSet" }, { type: "azure-native:datashare/v20210801:KustoClusterDataSet" }, { type: "azure-native:datashare/v20210801:KustoDatabaseDataSet" }, { type: "azure-native:datashare/v20210801:KustoTableDataSet" }, { type: "azure-native:datashare/v20210801:SqlDBTableDataSet" }, { type: "azure-native:datashare/v20210801:SqlDWTableDataSet" }, { type: "azure-native:datashare/v20210801:SynapseWorkspaceSqlPoolTableDataSet" }, { type: "azure-native:datashare:ADLSGen1FileDataSet" }, { type: "azure-native:datashare:ADLSGen1FolderDataSet" }, { type: "azure-native:datashare:ADLSGen2FileDataSet" }, { type: "azure-native:datashare:ADLSGen2FileSystemDataSet" }, { type: "azure-native:datashare:ADLSGen2FolderDataSet" }, { type: "azure-native:datashare:BlobContainerDataSet" }, { type: "azure-native:datashare:BlobDataSet" }, { type: "azure-native:datashare:BlobFolderDataSet" }, { type: "azure-native:datashare:KustoClusterDataSet" }, { type: "azure-native:datashare:KustoTableDataSet" }, { type: "azure-native:datashare:SqlDBTableDataSet" }, { type: "azure-native:datashare:SqlDWTableDataSet" }, { type: "azure-native:datashare:SynapseWorkspaceSqlPoolTableDataSet" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(KustoDatabaseDataSet.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a KustoDatabaseDataSet resource.
 */
export interface KustoDatabaseDataSetArgs {
    /**
     * The name of the share account.
     */
    accountName: pulumi.Input<string>;
    /**
     * The name of the dataSet.
     */
    dataSetName?: pulumi.Input<string>;
    /**
     * Kind of data set.
     * Expected value is 'KustoDatabase'.
     */
    kind: pulumi.Input<"KustoDatabase">;
    /**
     * Resource id of the kusto database.
     */
    kustoDatabaseResourceId: pulumi.Input<string>;
    /**
     * The resource group name.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the share to add the data set to.
     */
    shareName: pulumi.Input<string>;
}
