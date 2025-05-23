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
    /// Class representing a Kusto kusto pool.
    /// 
    /// Uses Azure REST API version 2021-06-01-preview. In version 2.x of the Azure Native provider, it used API version 2021-06-01-preview.
    /// 
    /// Other available API versions: 2021-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native synapse [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:synapse:KustoPool")]
    public partial class KustoPool : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// The Kusto Pool data ingestion URI.
        /// </summary>
        [Output("dataIngestionUri")]
        public Output<string> DataIngestionUri { get; private set; } = null!;

        /// <summary>
        /// A boolean value that indicates if the purge operations are enabled.
        /// </summary>
        [Output("enablePurge")]
        public Output<bool?> EnablePurge { get; private set; } = null!;

        /// <summary>
        /// A boolean value that indicates if the streaming ingest is enabled.
        /// </summary>
        [Output("enableStreamingIngest")]
        public Output<bool?> EnableStreamingIngest { get; private set; } = null!;

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// List of the Kusto Pool's language extensions.
        /// </summary>
        [Output("languageExtensions")]
        public Output<Outputs.LanguageExtensionsListResponse> LanguageExtensions { get; private set; } = null!;

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
        /// Optimized auto scale definition.
        /// </summary>
        [Output("optimizedAutoscale")]
        public Output<Outputs.OptimizedAutoscaleResponse?> OptimizedAutoscale { get; private set; } = null!;

        /// <summary>
        /// The provisioned state of the resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The SKU of the kusto pool.
        /// </summary>
        [Output("sku")]
        public Output<Outputs.AzureSkuResponse> Sku { get; private set; } = null!;

        /// <summary>
        /// The state of the resource.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The reason for the Kusto Pool's current state.
        /// </summary>
        [Output("stateReason")]
        public Output<string> StateReason { get; private set; } = null!;

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
        /// The Kusto Pool URI.
        /// </summary>
        [Output("uri")]
        public Output<string> Uri { get; private set; } = null!;

        /// <summary>
        /// The workspace unique identifier.
        /// </summary>
        [Output("workspaceUID")]
        public Output<string?> WorkspaceUID { get; private set; } = null!;


        /// <summary>
        /// Create a KustoPool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KustoPool(string name, KustoPoolArgs args, CustomResourceOptions? options = null)
            : base("azure-native:synapse:KustoPool", name, args ?? new KustoPoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private KustoPool(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:synapse:KustoPool", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:synapse/v20210401preview:KustoPool" },
                    new global::Pulumi.Alias { Type = "azure-native:synapse/v20210601preview:KustoPool" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing KustoPool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static KustoPool Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new KustoPool(name, id, options);
        }
    }

    public sealed class KustoPoolArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A boolean value that indicates if the purge operations are enabled.
        /// </summary>
        [Input("enablePurge")]
        public Input<bool>? EnablePurge { get; set; }

        /// <summary>
        /// A boolean value that indicates if the streaming ingest is enabled.
        /// </summary>
        [Input("enableStreamingIngest")]
        public Input<bool>? EnableStreamingIngest { get; set; }

        /// <summary>
        /// The name of the Kusto pool.
        /// </summary>
        [Input("kustoPoolName")]
        public Input<string>? KustoPoolName { get; set; }

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Optimized auto scale definition.
        /// </summary>
        [Input("optimizedAutoscale")]
        public Input<Inputs.OptimizedAutoscaleArgs>? OptimizedAutoscale { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The SKU of the kusto pool.
        /// </summary>
        [Input("sku", required: true)]
        public Input<Inputs.AzureSkuArgs> Sku { get; set; } = null!;

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

        /// <summary>
        /// The workspace unique identifier.
        /// </summary>
        [Input("workspaceUID")]
        public Input<string>? WorkspaceUID { get; set; }

        public KustoPoolArgs()
        {
            EnablePurge = false;
            EnableStreamingIngest = false;
        }
        public static new KustoPoolArgs Empty => new KustoPoolArgs();
    }
}
