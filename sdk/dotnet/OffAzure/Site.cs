// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.OffAzure
{
    /// <summary>
    /// Site REST Resource.
    /// 
    /// Uses Azure REST API version 2020-07-07. In version 2.x of the Azure Native provider, it used API version 2020-07-07.
    /// </summary>
    [AzureNativeResourceType("azure-native:offazure:Site")]
    public partial class Site : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// eTag for concurrency control.
        /// </summary>
        [Output("eTag")]
        public Output<string?> ETag { get; private set; } = null!;

        /// <summary>
        /// Azure location in which Sites is created.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Name of the VMware site.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// Nested properties of VMWare site.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.SitePropertiesResponse> Properties { get; private set; } = null!;

        /// <summary>
        /// Metadata pertaining to creation and last modification of the resource.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Type of resource. Type = Microsoft.OffAzure/VMWareSites.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Site resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Site(string name, SiteArgs args, CustomResourceOptions? options = null)
            : base("azure-native:offazure:Site", name, args ?? new SiteArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Site(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:offazure:Site", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:offazure/v20200101:Site" },
                    new global::Pulumi.Alias { Type = "azure-native:offazure/v20200707:Site" },
                    new global::Pulumi.Alias { Type = "azure-native:offazure/v20230606:Site" },
                    new global::Pulumi.Alias { Type = "azure-native:offazure/v20230606:SitesController" },
                    new global::Pulumi.Alias { Type = "azure-native:offazure/v20231001preview:Site" },
                    new global::Pulumi.Alias { Type = "azure-native:offazure/v20231001preview:SitesController" },
                    new global::Pulumi.Alias { Type = "azure-native:offazure/v20240501preview:Site" },
                    new global::Pulumi.Alias { Type = "azure-native:offazure/v20240501preview:SitesController" },
                    new global::Pulumi.Alias { Type = "azure-native:offazure:SitesController" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Site resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Site Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Site(name, id, options);
        }
    }

    public sealed class SiteArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// eTag for concurrency control.
        /// </summary>
        [Input("eTag")]
        public Input<string>? ETag { get; set; }

        /// <summary>
        /// Azure location in which Sites is created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Name of the VMware site.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Nested properties of VMWare site.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.SitePropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Site name.
        /// </summary>
        [Input("siteName")]
        public Input<string>? SiteName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public SiteArgs()
        {
        }
        public static new SiteArgs Empty => new SiteArgs();
    }
}
