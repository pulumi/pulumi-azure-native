// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Synapse
{
    /// <summary>
    /// Class representing an attached database configuration.
    /// 
    /// Uses Azure REST API version 2021-06-01-preview. In version 2.x of the Azure Native provider, it used API version 2021-06-01-preview.
    /// </summary>
    [AzureNativeResourceType("azure-native:synapse:KustoPoolAttachedDatabaseConfiguration")]
    public partial class KustoPoolAttachedDatabaseConfiguration : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The list of databases from the clusterResourceId which are currently attached to the kusto pool.
        /// </summary>
        [Output("attachedDatabaseNames")]
        public Output<ImmutableArray<string>> AttachedDatabaseNames { get; private set; } = null!;

        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// The name of the database which you would like to attach, use * if you want to follow all current and future databases.
        /// </summary>
        [Output("databaseName")]
        public Output<string> DatabaseName { get; private set; } = null!;

        /// <summary>
        /// The default principals modification kind
        /// </summary>
        [Output("defaultPrincipalsModificationKind")]
        public Output<string> DefaultPrincipalsModificationKind { get; private set; } = null!;

        /// <summary>
        /// The resource id of the kusto pool where the databases you would like to attach reside.
        /// </summary>
        [Output("kustoPoolResourceId")]
        public Output<string> KustoPoolResourceId { get; private set; } = null!;

        /// <summary>
        /// Resource location.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The provisioned state of the resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// Table level sharing specifications
        /// </summary>
        [Output("tableLevelSharingProperties")]
        public Output<Outputs.TableLevelSharingPropertiesResponse?> TableLevelSharingProperties { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a KustoPoolAttachedDatabaseConfiguration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KustoPoolAttachedDatabaseConfiguration(string name, KustoPoolAttachedDatabaseConfigurationArgs args, CustomResourceOptions? options = null)
            : base("azure-native:synapse:KustoPoolAttachedDatabaseConfiguration", name, args ?? new KustoPoolAttachedDatabaseConfigurationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private KustoPoolAttachedDatabaseConfiguration(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:synapse:KustoPoolAttachedDatabaseConfiguration", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:synapse/v20210601preview:KustoPoolAttachedDatabaseConfiguration" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing KustoPoolAttachedDatabaseConfiguration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static KustoPoolAttachedDatabaseConfiguration Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new KustoPoolAttachedDatabaseConfiguration(name, id, options);
        }
    }

    public sealed class KustoPoolAttachedDatabaseConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the attached database configuration.
        /// </summary>
        [Input("attachedDatabaseConfigurationName")]
        public Input<string>? AttachedDatabaseConfigurationName { get; set; }

        /// <summary>
        /// The name of the database which you would like to attach, use * if you want to follow all current and future databases.
        /// </summary>
        [Input("databaseName", required: true)]
        public Input<string> DatabaseName { get; set; } = null!;

        /// <summary>
        /// The default principals modification kind
        /// </summary>
        [Input("defaultPrincipalsModificationKind", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Synapse.DefaultPrincipalsModificationKind> DefaultPrincipalsModificationKind { get; set; } = null!;

        /// <summary>
        /// The name of the Kusto pool.
        /// </summary>
        [Input("kustoPoolName", required: true)]
        public Input<string> KustoPoolName { get; set; } = null!;

        /// <summary>
        /// The resource id of the kusto pool where the databases you would like to attach reside.
        /// </summary>
        [Input("kustoPoolResourceId", required: true)]
        public Input<string> KustoPoolResourceId { get; set; } = null!;

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Table level sharing specifications
        /// </summary>
        [Input("tableLevelSharingProperties")]
        public Input<Inputs.TableLevelSharingPropertiesArgs>? TableLevelSharingProperties { get; set; }

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public KustoPoolAttachedDatabaseConfigurationArgs()
        {
        }
        public static new KustoPoolAttachedDatabaseConfigurationArgs Empty => new KustoPoolAttachedDatabaseConfigurationArgs();
    }
}
