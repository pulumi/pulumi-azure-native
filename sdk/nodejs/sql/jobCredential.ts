// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * A stored credential that can be used by a job to connect to target databases.
 *
 * Uses Azure REST API version 2023-08-01. In version 2.x of the Azure Native provider, it used API version 2021-11-01.
 *
 * Other available API versions: 2017-03-01-preview, 2020-02-02-preview, 2020-08-01-preview, 2020-11-01-preview, 2021-02-01-preview, 2021-05-01-preview, 2021-08-01-preview, 2021-11-01, 2021-11-01-preview, 2022-02-01-preview, 2022-05-01-preview, 2022-08-01-preview, 2022-11-01-preview, 2023-02-01-preview, 2023-05-01-preview, 2023-08-01-preview, 2024-05-01-preview, 2024-11-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export class JobCredential extends pulumi.CustomResource {
    /**
     * Get an existing JobCredential resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): JobCredential {
        return new JobCredential(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:sql:JobCredential';

    /**
     * Returns true if the given object is an instance of JobCredential.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is JobCredential {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === JobCredential.__pulumiType;
    }

    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The credential user name.
     */
    public readonly username!: pulumi.Output<string>;

    /**
     * Create a JobCredential resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: JobCredentialArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.jobAgentName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'jobAgentName'");
            }
            if ((!args || args.password === undefined) && !opts.urn) {
                throw new Error("Missing required property 'password'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.serverName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serverName'");
            }
            if ((!args || args.username === undefined) && !opts.urn) {
                throw new Error("Missing required property 'username'");
            }
            resourceInputs["credentialName"] = args ? args.credentialName : undefined;
            resourceInputs["jobAgentName"] = args ? args.jobAgentName : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["serverName"] = args ? args.serverName : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["username"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:sql/v20170301preview:JobCredential" }, { type: "azure-native:sql/v20200202preview:JobCredential" }, { type: "azure-native:sql/v20200801preview:JobCredential" }, { type: "azure-native:sql/v20201101preview:JobCredential" }, { type: "azure-native:sql/v20210201preview:JobCredential" }, { type: "azure-native:sql/v20210501preview:JobCredential" }, { type: "azure-native:sql/v20210801preview:JobCredential" }, { type: "azure-native:sql/v20211101:JobCredential" }, { type: "azure-native:sql/v20211101preview:JobCredential" }, { type: "azure-native:sql/v20220201preview:JobCredential" }, { type: "azure-native:sql/v20220501preview:JobCredential" }, { type: "azure-native:sql/v20220801preview:JobCredential" }, { type: "azure-native:sql/v20221101preview:JobCredential" }, { type: "azure-native:sql/v20230201preview:JobCredential" }, { type: "azure-native:sql/v20230501preview:JobCredential" }, { type: "azure-native:sql/v20230801:JobCredential" }, { type: "azure-native:sql/v20230801preview:JobCredential" }, { type: "azure-native:sql/v20240501preview:JobCredential" }, { type: "azure-native:sql/v20241101preview:JobCredential" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(JobCredential.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a JobCredential resource.
 */
export interface JobCredentialArgs {
    /**
     * The name of the credential.
     */
    credentialName?: pulumi.Input<string>;
    /**
     * The name of the job agent.
     */
    jobAgentName: pulumi.Input<string>;
    /**
     * The credential password.
     */
    password: pulumi.Input<string>;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the server.
     */
    serverName: pulumi.Input<string>;
    /**
     * The credential user name.
     */
    username: pulumi.Input<string>;
}
