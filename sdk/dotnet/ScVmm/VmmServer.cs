// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ScVmm
{
    /// <summary>
    /// The VmmServers resource definition.
    /// 
    /// Uses Azure REST API version 2023-04-01-preview. In version 2.x of the Azure Native provider, it used API version 2022-05-21-preview.
    /// 
    /// Other available API versions: 2022-05-21-preview, 2023-10-07, 2024-06-01, 2025-03-13. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native scvmm [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:scvmm:VmmServer")]
    public partial class VmmServer : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the connection status to the vmmServer.
        /// </summary>
        [Output("connectionStatus")]
        public Output<string> ConnectionStatus { get; private set; } = null!;

        /// <summary>
        /// Credentials to connect to VMMServer.
        /// </summary>
        [Output("credentials")]
        public Output<Outputs.VMMServerPropertiesResponseCredentials?> Credentials { get; private set; } = null!;

        /// <summary>
        /// Gets or sets any error message if connection to vmmServer is having any issue.
        /// </summary>
        [Output("errorMessage")]
        public Output<string> ErrorMessage { get; private set; } = null!;

        /// <summary>
        /// The extended location.
        /// </summary>
        [Output("extendedLocation")]
        public Output<Outputs.ExtendedLocationResponse> ExtendedLocation { get; private set; } = null!;

        /// <summary>
        /// Fqdn is the hostname/ip of the vmmServer.
        /// </summary>
        [Output("fqdn")]
        public Output<string> Fqdn { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the location.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Resource Name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Port is the port on which the vmmServer is listening.
        /// </summary>
        [Output("port")]
        public Output<int?> Port { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the provisioning state.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The system data.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// Resource tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Resource Type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Unique ID of vmmServer.
        /// </summary>
        [Output("uuid")]
        public Output<string> Uuid { get; private set; } = null!;

        /// <summary>
        /// Version is the version of the vmmSever.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a VmmServer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VmmServer(string name, VmmServerArgs args, CustomResourceOptions? options = null)
            : base("azure-native:scvmm:VmmServer", name, args ?? new VmmServerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VmmServer(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:scvmm:VmmServer", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:scvmm/v20200605preview:VmmServer" },
                    new global::Pulumi.Alias { Type = "azure-native:scvmm/v20220521preview:VmmServer" },
                    new global::Pulumi.Alias { Type = "azure-native:scvmm/v20230401preview:VmmServer" },
                    new global::Pulumi.Alias { Type = "azure-native:scvmm/v20231007:VmmServer" },
                    new global::Pulumi.Alias { Type = "azure-native:scvmm/v20240601:VmmServer" },
                    new global::Pulumi.Alias { Type = "azure-native:scvmm/v20250313:VmmServer" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VmmServer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VmmServer Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new VmmServer(name, id, options);
        }
    }

    public sealed class VmmServerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Credentials to connect to VMMServer.
        /// </summary>
        [Input("credentials")]
        public Input<Inputs.VMMServerPropertiesCredentialsArgs>? Credentials { get; set; }

        /// <summary>
        /// The extended location.
        /// </summary>
        [Input("extendedLocation", required: true)]
        public Input<Inputs.ExtendedLocationArgs> ExtendedLocation { get; set; } = null!;

        /// <summary>
        /// Fqdn is the hostname/ip of the vmmServer.
        /// </summary>
        [Input("fqdn", required: true)]
        public Input<string> Fqdn { get; set; } = null!;

        /// <summary>
        /// Gets or sets the location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Port is the port on which the vmmServer is listening.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Name of the VMMServer.
        /// </summary>
        [Input("vmmServerName")]
        public Input<string>? VmmServerName { get; set; }

        public VmmServerArgs()
        {
        }
        public static new VmmServerArgs Empty => new VmmServerArgs();
    }
}
