// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CosmosDB
{
    /// <summary>
    /// An Azure Cosmos DB Fleetspace Account
    /// 
    /// Uses Azure REST API version 2025-05-01-preview.
    /// </summary>
    [AzureNativeResourceType("azure-native:cosmosdb:FleetspaceAccount")]
    public partial class FleetspaceAccount : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Configuration for fleetspace Account in the fleetspace.
        /// </summary>
        [Output("globalDatabaseAccountProperties")]
        public Output<Outputs.FleetspaceAccountPropertiesResponseGlobalDatabaseAccountProperties?> GlobalDatabaseAccountProperties { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A provisioning state of the Fleetspace Account.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a FleetspaceAccount resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FleetspaceAccount(string name, FleetspaceAccountArgs args, CustomResourceOptions? options = null)
            : base("azure-native:cosmosdb:FleetspaceAccount", name, args ?? new FleetspaceAccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FleetspaceAccount(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:cosmosdb:FleetspaceAccount", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:cosmosdb/v20250501preview:FleetspaceAccount" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing FleetspaceAccount resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FleetspaceAccount Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new FleetspaceAccount(name, id, options);
        }
    }

    public sealed class FleetspaceAccountArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cosmos DB fleet name. Needs to be unique under a subscription.
        /// </summary>
        [Input("fleetName", required: true)]
        public Input<string> FleetName { get; set; } = null!;

        /// <summary>
        /// Cosmos DB fleetspace account name.
        /// </summary>
        [Input("fleetspaceAccountName")]
        public Input<string>? FleetspaceAccountName { get; set; }

        /// <summary>
        /// Cosmos DB fleetspace name. Needs to be unique under a fleet.
        /// </summary>
        [Input("fleetspaceName", required: true)]
        public Input<string> FleetspaceName { get; set; } = null!;

        /// <summary>
        /// Configuration for fleetspace Account in the fleetspace.
        /// </summary>
        [Input("globalDatabaseAccountProperties")]
        public Input<Inputs.FleetspaceAccountPropertiesGlobalDatabaseAccountPropertiesArgs>? GlobalDatabaseAccountProperties { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public FleetspaceAccountArgs()
        {
        }
        public static new FleetspaceAccountArgs Empty => new FleetspaceAccountArgs();
    }
}
