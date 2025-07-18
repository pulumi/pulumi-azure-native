// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Azure Resource Manager resource envelope.
 *
 * Uses Azure REST API version 2024-10-01. In version 2.x of the Azure Native provider, it used API version 2023-04-01.
 *
 * Other available API versions: 2021-03-01-preview, 2022-02-01-preview, 2022-05-01, 2022-06-01-preview, 2022-10-01, 2022-10-01-preview, 2022-12-01-preview, 2023-02-01-preview, 2023-04-01, 2023-04-01-preview, 2023-06-01-preview, 2023-08-01-preview, 2023-10-01, 2024-01-01-preview, 2024-04-01, 2024-07-01-preview, 2024-10-01-preview, 2025-01-01-preview, 2025-04-01, 2025-04-01-preview, 2025-06-01, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native machinelearningservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export class Datastore extends pulumi.CustomResource {
    /**
     * Get an existing Datastore resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Datastore {
        return new Datastore(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:machinelearningservices:Datastore';

    /**
     * Returns true if the given object is an instance of Datastore.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Datastore {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Datastore.__pulumiType;
    }

    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * [Required] Additional attributes of the entity.
     */
    public readonly datastoreProperties!: pulumi.Output<outputs.machinelearningservices.AzureBlobDatastoreResponse | outputs.machinelearningservices.AzureDataLakeGen1DatastoreResponse | outputs.machinelearningservices.AzureDataLakeGen2DatastoreResponse | outputs.machinelearningservices.AzureFileDatastoreResponse | outputs.machinelearningservices.OneLakeDatastoreResponse>;
    /**
     * The name of the resource
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.machinelearningservices.SystemDataResponse>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Datastore resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatastoreArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.datastoreProperties === undefined) && !opts.urn) {
                throw new Error("Missing required property 'datastoreProperties'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.workspaceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workspaceName'");
            }
            resourceInputs["datastoreProperties"] = args ? args.datastoreProperties : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["skipValidation"] = args ? args.skipValidation : undefined;
            resourceInputs["workspaceName"] = args ? args.workspaceName : undefined;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["datastoreProperties"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:machinelearningservices/v20200501preview:Datastore" }, { type: "azure-native:machinelearningservices/v20200501preview:MachineLearningDatastore" }, { type: "azure-native:machinelearningservices/v20210301preview:Datastore" }, { type: "azure-native:machinelearningservices/v20220201preview:Datastore" }, { type: "azure-native:machinelearningservices/v20220501:Datastore" }, { type: "azure-native:machinelearningservices/v20220601preview:Datastore" }, { type: "azure-native:machinelearningservices/v20221001:Datastore" }, { type: "azure-native:machinelearningservices/v20221001preview:Datastore" }, { type: "azure-native:machinelearningservices/v20221201preview:Datastore" }, { type: "azure-native:machinelearningservices/v20230201preview:Datastore" }, { type: "azure-native:machinelearningservices/v20230401:Datastore" }, { type: "azure-native:machinelearningservices/v20230401preview:Datastore" }, { type: "azure-native:machinelearningservices/v20230601preview:Datastore" }, { type: "azure-native:machinelearningservices/v20230801preview:Datastore" }, { type: "azure-native:machinelearningservices/v20231001:Datastore" }, { type: "azure-native:machinelearningservices/v20240101preview:Datastore" }, { type: "azure-native:machinelearningservices/v20240401:Datastore" }, { type: "azure-native:machinelearningservices/v20240401preview:Datastore" }, { type: "azure-native:machinelearningservices/v20240701preview:Datastore" }, { type: "azure-native:machinelearningservices/v20241001:Datastore" }, { type: "azure-native:machinelearningservices/v20241001preview:Datastore" }, { type: "azure-native:machinelearningservices/v20250101preview:Datastore" }, { type: "azure-native:machinelearningservices/v20250401:Datastore" }, { type: "azure-native:machinelearningservices/v20250401preview:Datastore" }, { type: "azure-native:machinelearningservices/v20250601:Datastore" }, { type: "azure-native:machinelearningservices/v20250701preview:Datastore" }, { type: "azure-native:machinelearningservices:MachineLearningDatastore" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(Datastore.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Datastore resource.
 */
export interface DatastoreArgs {
    /**
     * [Required] Additional attributes of the entity.
     */
    datastoreProperties: pulumi.Input<inputs.machinelearningservices.AzureBlobDatastoreArgs | inputs.machinelearningservices.AzureDataLakeGen1DatastoreArgs | inputs.machinelearningservices.AzureDataLakeGen2DatastoreArgs | inputs.machinelearningservices.AzureFileDatastoreArgs | inputs.machinelearningservices.OneLakeDatastoreArgs>;
    /**
     * Datastore name.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Flag to skip validation.
     */
    skipValidation?: pulumi.Input<boolean>;
    /**
     * Name of Azure Machine Learning workspace.
     */
    workspaceName: pulumi.Input<string>;
}
