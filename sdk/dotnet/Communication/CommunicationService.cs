// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Communication
{
    /// <summary>
    /// A class representing a CommunicationService resource.
    /// 
    /// Uses Azure REST API version 2023-06-01-preview. In version 2.x of the Azure Native provider, it used API version 2023-03-31.
    /// 
    /// Other available API versions: 2023-03-31, 2023-04-01, 2023-04-01-preview, 2024-09-01-preview, 2025-05-01, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native communication [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:communication:CommunicationService")]
    public partial class CommunicationService : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// The location where the communication service stores its data at rest.
        /// </summary>
        [Output("dataLocation")]
        public Output<string> DataLocation { get; private set; } = null!;

        /// <summary>
        /// FQDN of the CommunicationService instance.
        /// </summary>
        [Output("hostName")]
        public Output<string> HostName { get; private set; } = null!;

        /// <summary>
        /// Managed service identity (system assigned and/or user assigned identities)
        /// </summary>
        [Output("identity")]
        public Output<Outputs.ManagedServiceIdentityResponse?> Identity { get; private set; } = null!;

        /// <summary>
        /// The immutable resource Id of the communication service.
        /// </summary>
        [Output("immutableResourceId")]
        public Output<string> ImmutableResourceId { get; private set; } = null!;

        /// <summary>
        /// List of email Domain resource Ids.
        /// </summary>
        [Output("linkedDomains")]
        public Output<ImmutableArray<string>> LinkedDomains { get; private set; } = null!;

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
        /// Resource ID of an Azure Notification Hub linked to this resource.
        /// </summary>
        [Output("notificationHubId")]
        public Output<string> NotificationHubId { get; private set; } = null!;

        /// <summary>
        /// Provisioning state of the resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

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
        /// Version of the CommunicationService resource. Probably you need the same or higher version of client SDKs.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a CommunicationService resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CommunicationService(string name, CommunicationServiceArgs args, CustomResourceOptions? options = null)
            : base("azure-native:communication:CommunicationService", name, args ?? new CommunicationServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CommunicationService(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:communication:CommunicationService", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:communication/v20200820:CommunicationService" },
                    new global::Pulumi.Alias { Type = "azure-native:communication/v20200820preview:CommunicationService" },
                    new global::Pulumi.Alias { Type = "azure-native:communication/v20211001preview:CommunicationService" },
                    new global::Pulumi.Alias { Type = "azure-native:communication/v20220701preview:CommunicationService" },
                    new global::Pulumi.Alias { Type = "azure-native:communication/v20230301preview:CommunicationService" },
                    new global::Pulumi.Alias { Type = "azure-native:communication/v20230331:CommunicationService" },
                    new global::Pulumi.Alias { Type = "azure-native:communication/v20230401:CommunicationService" },
                    new global::Pulumi.Alias { Type = "azure-native:communication/v20230401preview:CommunicationService" },
                    new global::Pulumi.Alias { Type = "azure-native:communication/v20230601preview:CommunicationService" },
                    new global::Pulumi.Alias { Type = "azure-native:communication/v20240901preview:CommunicationService" },
                    new global::Pulumi.Alias { Type = "azure-native:communication/v20250501:CommunicationService" },
                    new global::Pulumi.Alias { Type = "azure-native:communication/v20250501preview:CommunicationService" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CommunicationService resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CommunicationService Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new CommunicationService(name, id, options);
        }
    }

    public sealed class CommunicationServiceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the CommunicationService resource.
        /// </summary>
        [Input("communicationServiceName")]
        public Input<string>? CommunicationServiceName { get; set; }

        /// <summary>
        /// The location where the communication service stores its data at rest.
        /// </summary>
        [Input("dataLocation", required: true)]
        public Input<string> DataLocation { get; set; } = null!;

        /// <summary>
        /// Managed service identity (system assigned and/or user assigned identities)
        /// </summary>
        [Input("identity")]
        public Input<Inputs.ManagedServiceIdentityArgs>? Identity { get; set; }

        [Input("linkedDomains")]
        private InputList<string>? _linkedDomains;

        /// <summary>
        /// List of email Domain resource Ids.
        /// </summary>
        public InputList<string> LinkedDomains
        {
            get => _linkedDomains ?? (_linkedDomains = new InputList<string>());
            set => _linkedDomains = value;
        }

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

        public CommunicationServiceArgs()
        {
        }
        public static new CommunicationServiceArgs Empty => new CommunicationServiceArgs();
    }
}
