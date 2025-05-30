// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.StorSimple
{
    /// <summary>
    /// The StorSimple Manager.
    /// 
    /// Uses Azure REST API version 2017-06-01. In version 2.x of the Azure Native provider, it used API version 2017-06-01.
    /// </summary>
    [AzureNativeResourceType("azure-native:storsimple:Manager")]
    public partial class Manager : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Represents the type of StorSimple Manager.
        /// </summary>
        [Output("cisIntrinsicSettings")]
        public Output<Outputs.ManagerIntrinsicSettingsResponse?> CisIntrinsicSettings { get; private set; } = null!;

        /// <summary>
        /// The etag of the manager.
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

        /// <summary>
        /// The geo location of the resource.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the state of the resource as it is getting provisioned. Value of "Succeeded" means the Manager was successfully created.
        /// </summary>
        [Output("provisioningState")]
        public Output<string?> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Specifies the Sku.
        /// </summary>
        [Output("sku")]
        public Output<Outputs.ManagerSkuResponse?> Sku { get; private set; } = null!;

        /// <summary>
        /// The tags attached to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Manager resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Manager(string name, ManagerArgs args, CustomResourceOptions? options = null)
            : base("azure-native:storsimple:Manager", name, args ?? new ManagerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Manager(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:storsimple:Manager", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:storsimple/v20161001:Manager" },
                    new global::Pulumi.Alias { Type = "azure-native:storsimple/v20170601:Manager" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Manager resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Manager Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Manager(name, id, options);
        }
    }

    public sealed class ManagerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Represents the type of StorSimple Manager.
        /// </summary>
        [Input("cisIntrinsicSettings")]
        public Input<Inputs.ManagerIntrinsicSettingsArgs>? CisIntrinsicSettings { get; set; }

        /// <summary>
        /// The geo location of the resource.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The manager name
        /// </summary>
        [Input("managerName")]
        public Input<string>? ManagerName { get; set; }

        /// <summary>
        /// Specifies the state of the resource as it is getting provisioned. Value of "Succeeded" means the Manager was successfully created.
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        /// <summary>
        /// The resource group name
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Specifies the Sku.
        /// </summary>
        [Input("sku")]
        public Input<Inputs.ManagerSkuArgs>? Sku { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The tags attached to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ManagerArgs()
        {
        }
        public static new ManagerArgs Empty => new ManagerArgs();
    }
}
