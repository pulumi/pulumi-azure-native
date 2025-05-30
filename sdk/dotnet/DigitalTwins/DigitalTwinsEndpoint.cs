// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DigitalTwins
{
    /// <summary>
    /// DigitalTwinsInstance endpoint resource.
    /// 
    /// Uses Azure REST API version 2023-01-31. In version 2.x of the Azure Native provider, it used API version 2023-01-31.
    /// </summary>
    [AzureNativeResourceType("azure-native:digitaltwins:DigitalTwinsEndpoint")]
    public partial class DigitalTwinsEndpoint : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Extension resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// DigitalTwinsInstance endpoint resource properties.
        /// </summary>
        [Output("properties")]
        public Output<object> Properties { get; private set; } = null!;

        /// <summary>
        /// Metadata pertaining to creation and last modification of the resource.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// The resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a DigitalTwinsEndpoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DigitalTwinsEndpoint(string name, DigitalTwinsEndpointArgs args, CustomResourceOptions? options = null)
            : base("azure-native:digitaltwins:DigitalTwinsEndpoint", name, args ?? new DigitalTwinsEndpointArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DigitalTwinsEndpoint(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:digitaltwins:DigitalTwinsEndpoint", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:digitaltwins/v20200301preview:DigitalTwinsEndpoint" },
                    new global::Pulumi.Alias { Type = "azure-native:digitaltwins/v20201031:DigitalTwinsEndpoint" },
                    new global::Pulumi.Alias { Type = "azure-native:digitaltwins/v20201201:DigitalTwinsEndpoint" },
                    new global::Pulumi.Alias { Type = "azure-native:digitaltwins/v20210630preview:DigitalTwinsEndpoint" },
                    new global::Pulumi.Alias { Type = "azure-native:digitaltwins/v20220531:DigitalTwinsEndpoint" },
                    new global::Pulumi.Alias { Type = "azure-native:digitaltwins/v20221031:DigitalTwinsEndpoint" },
                    new global::Pulumi.Alias { Type = "azure-native:digitaltwins/v20230131:DigitalTwinsEndpoint" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DigitalTwinsEndpoint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DigitalTwinsEndpoint Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DigitalTwinsEndpoint(name, id, options);
        }
    }

    public sealed class DigitalTwinsEndpointArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of Endpoint Resource.
        /// </summary>
        [Input("endpointName")]
        public Input<string>? EndpointName { get; set; }

        /// <summary>
        /// DigitalTwinsInstance endpoint resource properties.
        /// </summary>
        [Input("properties", required: true)]
        public object Properties { get; set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the DigitalTwinsInstance.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the DigitalTwinsInstance.
        /// </summary>
        [Input("resourceName", required: true)]
        public Input<string> ResourceName { get; set; } = null!;

        public DigitalTwinsEndpointArgs()
        {
        }
        public static new DigitalTwinsEndpointArgs Empty => new DigitalTwinsEndpointArgs();
    }
}
