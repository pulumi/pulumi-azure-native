// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Media
{
    /// <summary>
    /// A Streaming Locator resource
    /// 
    /// Uses Azure REST API version 2023-01-01. In version 2.x of the Azure Native provider, it used API version 2023-01-01.
    /// 
    /// Other available API versions: 2018-03-30-preview, 2018-06-01-preview, 2018-07-01, 2020-05-01, 2021-06-01, 2021-11-01, 2022-08-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native media [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:media:StreamingLocator")]
    public partial class StreamingLocator : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Alternative Media ID of this Streaming Locator
        /// </summary>
        [Output("alternativeMediaId")]
        public Output<string?> AlternativeMediaId { get; private set; } = null!;

        /// <summary>
        /// Asset Name
        /// </summary>
        [Output("assetName")]
        public Output<string> AssetName { get; private set; } = null!;

        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// The ContentKeys used by this Streaming Locator.
        /// </summary>
        [Output("contentKeys")]
        public Output<ImmutableArray<Outputs.StreamingLocatorContentKeyResponse>> ContentKeys { get; private set; } = null!;

        /// <summary>
        /// The creation time of the Streaming Locator.
        /// </summary>
        [Output("created")]
        public Output<string> Created { get; private set; } = null!;

        /// <summary>
        /// Name of the default ContentKeyPolicy used by this Streaming Locator.
        /// </summary>
        [Output("defaultContentKeyPolicyName")]
        public Output<string?> DefaultContentKeyPolicyName { get; private set; } = null!;

        /// <summary>
        /// The end time of the Streaming Locator.
        /// </summary>
        [Output("endTime")]
        public Output<string?> EndTime { get; private set; } = null!;

        /// <summary>
        /// A list of asset or account filters which apply to this streaming locator
        /// </summary>
        [Output("filters")]
        public Output<ImmutableArray<string>> Filters { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The start time of the Streaming Locator.
        /// </summary>
        [Output("startTime")]
        public Output<string?> StartTime { get; private set; } = null!;

        /// <summary>
        /// The StreamingLocatorId of the Streaming Locator.
        /// </summary>
        [Output("streamingLocatorId")]
        public Output<string?> StreamingLocatorId { get; private set; } = null!;

        /// <summary>
        /// Name of the Streaming Policy used by this Streaming Locator. Either specify the name of Streaming Policy you created or use one of the predefined Streaming Policies. The predefined Streaming Policies available are: 'Predefined_DownloadOnly', 'Predefined_ClearStreamingOnly', 'Predefined_DownloadAndClearStreaming', 'Predefined_ClearKey', 'Predefined_MultiDrmCencStreaming' and 'Predefined_MultiDrmStreaming'
        /// </summary>
        [Output("streamingPolicyName")]
        public Output<string> StreamingPolicyName { get; private set; } = null!;

        /// <summary>
        /// The system metadata relating to this resource.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a StreamingLocator resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StreamingLocator(string name, StreamingLocatorArgs args, CustomResourceOptions? options = null)
            : base("azure-native:media:StreamingLocator", name, args ?? new StreamingLocatorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StreamingLocator(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:media:StreamingLocator", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:media/v20180330preview:StreamingLocator" },
                    new global::Pulumi.Alias { Type = "azure-native:media/v20180601preview:StreamingLocator" },
                    new global::Pulumi.Alias { Type = "azure-native:media/v20180701:StreamingLocator" },
                    new global::Pulumi.Alias { Type = "azure-native:media/v20200501:StreamingLocator" },
                    new global::Pulumi.Alias { Type = "azure-native:media/v20210601:StreamingLocator" },
                    new global::Pulumi.Alias { Type = "azure-native:media/v20211101:StreamingLocator" },
                    new global::Pulumi.Alias { Type = "azure-native:media/v20220801:StreamingLocator" },
                    new global::Pulumi.Alias { Type = "azure-native:media/v20230101:StreamingLocator" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing StreamingLocator resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StreamingLocator Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new StreamingLocator(name, id, options);
        }
    }

    public sealed class StreamingLocatorArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Media Services account name.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// Alternative Media ID of this Streaming Locator
        /// </summary>
        [Input("alternativeMediaId")]
        public Input<string>? AlternativeMediaId { get; set; }

        /// <summary>
        /// Asset Name
        /// </summary>
        [Input("assetName", required: true)]
        public Input<string> AssetName { get; set; } = null!;

        [Input("contentKeys")]
        private InputList<Inputs.StreamingLocatorContentKeyArgs>? _contentKeys;

        /// <summary>
        /// The ContentKeys used by this Streaming Locator.
        /// </summary>
        public InputList<Inputs.StreamingLocatorContentKeyArgs> ContentKeys
        {
            get => _contentKeys ?? (_contentKeys = new InputList<Inputs.StreamingLocatorContentKeyArgs>());
            set => _contentKeys = value;
        }

        /// <summary>
        /// Name of the default ContentKeyPolicy used by this Streaming Locator.
        /// </summary>
        [Input("defaultContentKeyPolicyName")]
        public Input<string>? DefaultContentKeyPolicyName { get; set; }

        /// <summary>
        /// The end time of the Streaming Locator.
        /// </summary>
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        [Input("filters")]
        private InputList<string>? _filters;

        /// <summary>
        /// A list of asset or account filters which apply to this streaming locator
        /// </summary>
        public InputList<string> Filters
        {
            get => _filters ?? (_filters = new InputList<string>());
            set => _filters = value;
        }

        /// <summary>
        /// The name of the resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The start time of the Streaming Locator.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        /// <summary>
        /// The StreamingLocatorId of the Streaming Locator.
        /// </summary>
        [Input("streamingLocatorId")]
        public Input<string>? StreamingLocatorId { get; set; }

        /// <summary>
        /// The Streaming Locator name.
        /// </summary>
        [Input("streamingLocatorName")]
        public Input<string>? StreamingLocatorName { get; set; }

        /// <summary>
        /// Name of the Streaming Policy used by this Streaming Locator. Either specify the name of Streaming Policy you created or use one of the predefined Streaming Policies. The predefined Streaming Policies available are: 'Predefined_DownloadOnly', 'Predefined_ClearStreamingOnly', 'Predefined_DownloadAndClearStreaming', 'Predefined_ClearKey', 'Predefined_MultiDrmCencStreaming' and 'Predefined_MultiDrmStreaming'
        /// </summary>
        [Input("streamingPolicyName", required: true)]
        public Input<string> StreamingPolicyName { get; set; } = null!;

        public StreamingLocatorArgs()
        {
        }
        public static new StreamingLocatorArgs Empty => new StreamingLocatorArgs();
    }
}
