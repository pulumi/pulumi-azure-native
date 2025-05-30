// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Automation
{
    /// <summary>
    /// Definition of the module type.
    /// 
    /// Uses Azure REST API version 2023-11-01. In version 2.x of the Azure Native provider, it used API version 2022-08-08.
    /// 
    /// Other available API versions: 2018-06-30, 2019-06-01, 2020-01-13-preview, 2022-08-08, 2023-05-15-preview, 2024-10-23. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native automation [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:automation:Python2Package")]
    public partial class Python2Package : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Gets the activity count of the module.
        /// </summary>
        [Output("activityCount")]
        public Output<int?> ActivityCount { get; private set; } = null!;

        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Gets the creation time.
        /// </summary>
        [Output("creationTime")]
        public Output<string?> CreationTime { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Gets the error info of the module.
        /// </summary>
        [Output("error")]
        public Output<Outputs.ModuleErrorInfoResponse?> Error { get; private set; } = null!;

        /// <summary>
        /// Gets the etag of the resource.
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

        /// <summary>
        /// Gets type of module, if its composite or not.
        /// </summary>
        [Output("isComposite")]
        public Output<bool?> IsComposite { get; private set; } = null!;

        /// <summary>
        /// Gets the isGlobal flag of the module.
        /// </summary>
        [Output("isGlobal")]
        public Output<bool?> IsGlobal { get; private set; } = null!;

        /// <summary>
        /// Gets the last modified time.
        /// </summary>
        [Output("lastModifiedTime")]
        public Output<string?> LastModifiedTime { get; private set; } = null!;

        /// <summary>
        /// The Azure Region where the resource lives
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Gets the provisioning state of the module.
        /// </summary>
        [Output("provisioningState")]
        public Output<string?> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Gets the size in bytes of the module.
        /// </summary>
        [Output("sizeInBytes")]
        public Output<double?> SizeInBytes { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Gets the version of the module.
        /// </summary>
        [Output("version")]
        public Output<string?> Version { get; private set; } = null!;


        /// <summary>
        /// Create a Python2Package resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Python2Package(string name, Python2PackageArgs args, CustomResourceOptions? options = null)
            : base("azure-native:automation:Python2Package", name, args ?? new Python2PackageArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Python2Package(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:automation:Python2Package", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:automation/v20180630:Python2Package" },
                    new global::Pulumi.Alias { Type = "azure-native:automation/v20190601:Python2Package" },
                    new global::Pulumi.Alias { Type = "azure-native:automation/v20200113preview:Python2Package" },
                    new global::Pulumi.Alias { Type = "azure-native:automation/v20220808:Python2Package" },
                    new global::Pulumi.Alias { Type = "azure-native:automation/v20230515preview:Python2Package" },
                    new global::Pulumi.Alias { Type = "azure-native:automation/v20231101:Python2Package" },
                    new global::Pulumi.Alias { Type = "azure-native:automation/v20241023:Python2Package" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Python2Package resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Python2Package Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Python2Package(name, id, options);
        }
    }

    public sealed class Python2PackageArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the automation account.
        /// </summary>
        [Input("automationAccountName", required: true)]
        public Input<string> AutomationAccountName { get; set; } = null!;

        /// <summary>
        /// Gets or sets the module content link.
        /// </summary>
        [Input("contentLink", required: true)]
        public Input<Inputs.ContentLinkArgs> ContentLink { get; set; } = null!;

        /// <summary>
        /// The name of python package.
        /// </summary>
        [Input("packageName")]
        public Input<string>? PackageName { get; set; }

        /// <summary>
        /// Name of an Azure Resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Gets or sets the tags attached to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public Python2PackageArgs()
        {
        }
        public static new Python2PackageArgs Empty => new Python2PackageArgs();
    }
}
