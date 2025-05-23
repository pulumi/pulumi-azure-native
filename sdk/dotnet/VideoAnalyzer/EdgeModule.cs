// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.VideoAnalyzer
{
    /// <summary>
    /// The representation of an edge module.
    /// 
    /// Uses Azure REST API version 2021-11-01-preview. In version 2.x of the Azure Native provider, it used API version 2021-11-01-preview.
    /// </summary>
    [AzureNativeResourceType("azure-native:videoanalyzer:EdgeModule")]
    public partial class EdgeModule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Internal ID generated for the instance of the Video Analyzer edge module.
        /// </summary>
        [Output("edgeModuleId")]
        public Output<string> EdgeModuleId { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

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
        /// Create a EdgeModule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EdgeModule(string name, EdgeModuleArgs args, CustomResourceOptions? options = null)
            : base("azure-native:videoanalyzer:EdgeModule", name, args ?? new EdgeModuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EdgeModule(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:videoanalyzer:EdgeModule", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:videoanalyzer/v20210501preview:EdgeModule" },
                    new global::Pulumi.Alias { Type = "azure-native:videoanalyzer/v20211101preview:EdgeModule" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing EdgeModule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EdgeModule Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new EdgeModule(name, id, options);
        }
    }

    public sealed class EdgeModuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Azure Video Analyzer account name.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The Edge Module name.
        /// </summary>
        [Input("edgeModuleName")]
        public Input<string>? EdgeModuleName { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public EdgeModuleArgs()
        {
        }
        public static new EdgeModuleArgs Empty => new EdgeModuleArgs();
    }
}
