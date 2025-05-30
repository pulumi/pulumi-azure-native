// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Migrate
{
    /// <summary>
    /// ModernizeProject model.
    /// 
    /// Uses Azure REST API version 2022-05-01-preview. In version 2.x of the Azure Native provider, it used API version 2022-05-01-preview.
    /// </summary>
    [AzureNativeResourceType("azure-native:migrate:ModernizeProject")]
    public partial class ModernizeProject : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        [Output("identity")]
        public Output<Outputs.ResourceIdentityResponse?> Identity { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the location of the modernizeProject.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the name of the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// ModernizeProject properties.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.ModernizeProjectModelPropertiesResponse> Properties { get; private set; } = null!;

        [Output("systemData")]
        public Output<Outputs.ModernizeProjectModelResponseSystemData> SystemData { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the type of the resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ModernizeProject resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ModernizeProject(string name, ModernizeProjectArgs args, CustomResourceOptions? options = null)
            : base("azure-native:migrate:ModernizeProject", name, args ?? new ModernizeProjectArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ModernizeProject(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:migrate:ModernizeProject", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:migrate/v20220501preview:ModernizeProject" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ModernizeProject resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ModernizeProject Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ModernizeProject(name, id, options);
        }
    }

    public sealed class ModernizeProjectArgs : global::Pulumi.ResourceArgs
    {
        [Input("identity")]
        public Input<Inputs.ResourceIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// Gets or sets the location of the modernizeProject.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// ModernizeProject Name.
        /// </summary>
        [Input("modernizeProjectName")]
        public Input<string>? ModernizeProjectName { get; set; }

        /// <summary>
        /// ModernizeProject properties.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.ModernizeProjectModelPropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// Name of the Azure Resource Group that project is part of.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Azure Subscription Id in which project was created.
        /// </summary>
        [Input("subscriptionId")]
        public Input<string>? SubscriptionId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Gets or sets the resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ModernizeProjectArgs()
        {
        }
        public static new ModernizeProjectArgs Empty => new ModernizeProjectArgs();
    }
}
