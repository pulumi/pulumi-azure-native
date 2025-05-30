// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.M365SecurityAndCompliance
{
    /// <summary>
    /// The description of the service.
    /// 
    /// Uses Azure REST API version 2021-03-25-preview.
    /// </summary>
    [AzureNativeResourceType("azure-native:m365securityandcompliance:PrivateLinkServicesForEDMUpload")]
    public partial class PrivateLinkServicesForEDMUpload : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// An etag associated with the resource, used for optimistic concurrency when editing it.
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

        /// <summary>
        /// Setting indicating whether the service has a managed identity associated with it.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.ServicesResourceResponseIdentity?> Identity { get; private set; } = null!;

        /// <summary>
        /// The kind of the service.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// The resource location.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The common properties of a service.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.ServicesPropertiesResponse> Properties { get; private set; } = null!;

        /// <summary>
        /// Required property for system data
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// The resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a PrivateLinkServicesForEDMUpload resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PrivateLinkServicesForEDMUpload(string name, PrivateLinkServicesForEDMUploadArgs args, CustomResourceOptions? options = null)
            : base("azure-native:m365securityandcompliance:PrivateLinkServicesForEDMUpload", name, args ?? new PrivateLinkServicesForEDMUploadArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PrivateLinkServicesForEDMUpload(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:m365securityandcompliance:PrivateLinkServicesForEDMUpload", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:m365securityandcompliance/v20210325preview:PrivateLinkServicesForEDMUpload" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PrivateLinkServicesForEDMUpload resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PrivateLinkServicesForEDMUpload Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PrivateLinkServicesForEDMUpload(name, id, options);
        }
    }

    public sealed class PrivateLinkServicesForEDMUploadArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Setting indicating whether the service has a managed identity associated with it.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.ServicesResourceIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// The kind of the service.
        /// </summary>
        [Input("kind", required: true)]
        public Input<Pulumi.AzureNative.M365SecurityAndCompliance.Kind> Kind { get; set; } = null!;

        /// <summary>
        /// The resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The common properties of a service.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.ServicesPropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// The name of the resource group that contains the service instance.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the service instance.
        /// </summary>
        [Input("resourceName")]
        public Input<string>? ResourceName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public PrivateLinkServicesForEDMUploadArgs()
        {
        }
        public static new PrivateLinkServicesForEDMUploadArgs Empty => new PrivateLinkServicesForEDMUploadArgs();
    }
}
