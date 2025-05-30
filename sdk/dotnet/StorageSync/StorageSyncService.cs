// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.StorageSync
{
    /// <summary>
    /// Storage Sync Service object.
    /// 
    /// Uses Azure REST API version 2022-09-01. In version 2.x of the Azure Native provider, it used API version 2022-06-01.
    /// 
    /// Other available API versions: 2022-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native storagesync [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:storagesync:StorageSyncService")]
    public partial class StorageSyncService : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// managed identities for the Storage Sync service to interact with other Azure services without maintaining any secrets or credentials in code.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.ManagedServiceIdentityResponse?> Identity { get; private set; } = null!;

        /// <summary>
        /// Incoming Traffic Policy
        /// </summary>
        [Output("incomingTrafficPolicy")]
        public Output<string?> IncomingTrafficPolicy { get; private set; } = null!;

        /// <summary>
        /// Resource Last Operation Name
        /// </summary>
        [Output("lastOperationName")]
        public Output<string> LastOperationName { get; private set; } = null!;

        /// <summary>
        /// StorageSyncService lastWorkflowId
        /// </summary>
        [Output("lastWorkflowId")]
        public Output<string> LastWorkflowId { get; private set; } = null!;

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// List of private endpoint connection associated with the specified storage sync service
        /// </summary>
        [Output("privateEndpointConnections")]
        public Output<ImmutableArray<Outputs.PrivateEndpointConnectionResponse>> PrivateEndpointConnections { get; private set; } = null!;

        /// <summary>
        /// StorageSyncService Provisioning State
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Storage Sync service status.
        /// </summary>
        [Output("storageSyncServiceStatus")]
        public Output<int> StorageSyncServiceStatus { get; private set; } = null!;

        /// <summary>
        /// Storage Sync service Uid
        /// </summary>
        [Output("storageSyncServiceUid")]
        public Output<string> StorageSyncServiceUid { get; private set; } = null!;

        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Use Identity authorization when customer have finished setup RBAC permissions.
        /// </summary>
        [Output("useIdentity")]
        public Output<bool> UseIdentity { get; private set; } = null!;


        /// <summary>
        /// Create a StorageSyncService resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StorageSyncService(string name, StorageSyncServiceArgs args, CustomResourceOptions? options = null)
            : base("azure-native:storagesync:StorageSyncService", name, args ?? new StorageSyncServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StorageSyncService(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:storagesync:StorageSyncService", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:storagesync/v20170605preview:StorageSyncService" },
                    new global::Pulumi.Alias { Type = "azure-native:storagesync/v20180402:StorageSyncService" },
                    new global::Pulumi.Alias { Type = "azure-native:storagesync/v20180701:StorageSyncService" },
                    new global::Pulumi.Alias { Type = "azure-native:storagesync/v20181001:StorageSyncService" },
                    new global::Pulumi.Alias { Type = "azure-native:storagesync/v20190201:StorageSyncService" },
                    new global::Pulumi.Alias { Type = "azure-native:storagesync/v20190301:StorageSyncService" },
                    new global::Pulumi.Alias { Type = "azure-native:storagesync/v20190601:StorageSyncService" },
                    new global::Pulumi.Alias { Type = "azure-native:storagesync/v20191001:StorageSyncService" },
                    new global::Pulumi.Alias { Type = "azure-native:storagesync/v20200301:StorageSyncService" },
                    new global::Pulumi.Alias { Type = "azure-native:storagesync/v20200901:StorageSyncService" },
                    new global::Pulumi.Alias { Type = "azure-native:storagesync/v20220601:StorageSyncService" },
                    new global::Pulumi.Alias { Type = "azure-native:storagesync/v20220901:StorageSyncService" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing StorageSyncService resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StorageSyncService Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new StorageSyncService(name, id, options);
        }
    }

    public sealed class StorageSyncServiceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// managed identities for the Storage Sync to interact with other Azure services without maintaining any secrets or credentials in code.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.ManagedServiceIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// Incoming Traffic Policy
        /// </summary>
        [Input("incomingTrafficPolicy")]
        public InputUnion<string, Pulumi.AzureNative.StorageSync.IncomingTrafficPolicy>? IncomingTrafficPolicy { get; set; }

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of Storage Sync Service resource.
        /// </summary>
        [Input("storageSyncServiceName")]
        public Input<string>? StorageSyncServiceName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Use Identity authorization when customer have finished setup RBAC permissions.
        /// </summary>
        [Input("useIdentity")]
        public Input<bool>? UseIdentity { get; set; }

        public StorageSyncServiceArgs()
        {
        }
        public static new StorageSyncServiceArgs Empty => new StorageSyncServiceArgs();
    }
}
