// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.OperationalInsights
{
    /// <summary>
    /// The top level storage insight resource container.
    /// 
    /// Uses Azure REST API version 2023-09-01. In version 2.x of the Azure Native provider, it used API version 2020-08-01.
    /// 
    /// Other available API versions: 2020-03-01-preview, 2020-08-01, 2025-02-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native operationalinsights [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:operationalinsights:StorageInsightConfig")]
    public partial class StorageInsightConfig : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// The names of the blob containers that the workspace should read
        /// </summary>
        [Output("containers")]
        public Output<ImmutableArray<string>> Containers { get; private set; } = null!;

        /// <summary>
        /// The ETag of the storage insight.
        /// </summary>
        [Output("eTag")]
        public Output<string?> ETag { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The status of the storage insight
        /// </summary>
        [Output("status")]
        public Output<Outputs.StorageInsightStatusResponse> Status { get; private set; } = null!;

        /// <summary>
        /// The storage account connection details
        /// </summary>
        [Output("storageAccount")]
        public Output<Outputs.StorageAccountResponse> StorageAccount { get; private set; } = null!;

        /// <summary>
        /// The names of the Azure tables that the workspace should read
        /// </summary>
        [Output("tables")]
        public Output<ImmutableArray<string>> Tables { get; private set; } = null!;

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
        /// Create a StorageInsightConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StorageInsightConfig(string name, StorageInsightConfigArgs args, CustomResourceOptions? options = null)
            : base("azure-native:operationalinsights:StorageInsightConfig", name, args ?? new StorageInsightConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StorageInsightConfig(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:operationalinsights:StorageInsightConfig", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:operationalinsights/v20150320:StorageInsightConfig" },
                    new global::Pulumi.Alias { Type = "azure-native:operationalinsights/v20200301preview:StorageInsightConfig" },
                    new global::Pulumi.Alias { Type = "azure-native:operationalinsights/v20200801:StorageInsightConfig" },
                    new global::Pulumi.Alias { Type = "azure-native:operationalinsights/v20230901:StorageInsightConfig" },
                    new global::Pulumi.Alias { Type = "azure-native:operationalinsights/v20250201:StorageInsightConfig" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing StorageInsightConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StorageInsightConfig Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new StorageInsightConfig(name, id, options);
        }
    }

    public sealed class StorageInsightConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("containers")]
        private InputList<string>? _containers;

        /// <summary>
        /// The names of the blob containers that the workspace should read
        /// </summary>
        public InputList<string> Containers
        {
            get => _containers ?? (_containers = new InputList<string>());
            set => _containers = value;
        }

        /// <summary>
        /// The ETag of the storage insight.
        /// </summary>
        [Input("eTag")]
        public Input<string>? ETag { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The storage account connection details
        /// </summary>
        [Input("storageAccount", required: true)]
        public Input<Inputs.StorageAccountArgs> StorageAccount { get; set; } = null!;

        /// <summary>
        /// Name of the storageInsightsConfigs resource
        /// </summary>
        [Input("storageInsightName")]
        public Input<string>? StorageInsightName { get; set; }

        [Input("tables")]
        private InputList<string>? _tables;

        /// <summary>
        /// The names of the Azure tables that the workspace should read
        /// </summary>
        public InputList<string> Tables
        {
            get => _tables ?? (_tables = new InputList<string>());
            set => _tables = value;
        }

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
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public StorageInsightConfigArgs()
        {
        }
        public static new StorageInsightConfigArgs Empty => new StorageInsightConfigArgs();
    }
}
