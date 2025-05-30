// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ConnectedVMwarevSphere
{
    /// <summary>
    /// Defines the GuestAgent.
    /// 
    /// Uses Azure REST API version 2023-03-01-preview. In version 2.x of the Azure Native provider, it used API version 2022-07-15-preview.
    /// 
    /// Other available API versions: 2022-07-15-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native connectedvmwarevsphere [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:connectedvmwarevsphere:GuestAgent")]
    public partial class GuestAgent : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Username / Password Credentials to provision guest agent.
        /// </summary>
        [Output("credentials")]
        public Output<Outputs.GuestCredentialResponse?> Credentials { get; private set; } = null!;

        /// <summary>
        /// Gets the name of the corresponding resource in Kubernetes.
        /// </summary>
        [Output("customResourceName")]
        public Output<string> CustomResourceName { get; private set; } = null!;

        /// <summary>
        /// HTTP Proxy configuration for the VM.
        /// </summary>
        [Output("httpProxyConfig")]
        public Output<Outputs.HttpProxyConfigurationResponse?> HttpProxyConfig { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The resource id of the private link scope this machine is assigned to, if any.
        /// </summary>
        [Output("privateLinkScopeResourceId")]
        public Output<string?> PrivateLinkScopeResourceId { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the guest agent provisioning action.
        /// </summary>
        [Output("provisioningAction")]
        public Output<string?> ProvisioningAction { get; private set; } = null!;

        /// <summary>
        /// Gets the provisioning state.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the guest agent status.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The resource status information.
        /// </summary>
        [Output("statuses")]
        public Output<ImmutableArray<Outputs.ResourceStatusResponse>> Statuses { get; private set; } = null!;

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
        /// Gets or sets a unique identifier for this resource.
        /// </summary>
        [Output("uuid")]
        public Output<string> Uuid { get; private set; } = null!;


        /// <summary>
        /// Create a GuestAgent resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GuestAgent(string name, GuestAgentArgs args, CustomResourceOptions? options = null)
            : base("azure-native:connectedvmwarevsphere:GuestAgent", name, args ?? new GuestAgentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GuestAgent(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:connectedvmwarevsphere:GuestAgent", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:connectedvmwarevsphere/v20201001preview:GuestAgent" },
                    new global::Pulumi.Alias { Type = "azure-native:connectedvmwarevsphere/v20220110preview:GuestAgent" },
                    new global::Pulumi.Alias { Type = "azure-native:connectedvmwarevsphere/v20220715preview:GuestAgent" },
                    new global::Pulumi.Alias { Type = "azure-native:connectedvmwarevsphere/v20230301preview:GuestAgent" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing GuestAgent resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GuestAgent Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new GuestAgent(name, id, options);
        }
    }

    public sealed class GuestAgentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Username / Password Credentials to provision guest agent.
        /// </summary>
        [Input("credentials")]
        public Input<Inputs.GuestCredentialArgs>? Credentials { get; set; }

        /// <summary>
        /// HTTP Proxy configuration for the VM.
        /// </summary>
        [Input("httpProxyConfig")]
        public Input<Inputs.HttpProxyConfigurationArgs>? HttpProxyConfig { get; set; }

        /// <summary>
        /// Name of the guestAgents.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The resource id of the private link scope this machine is assigned to, if any.
        /// </summary>
        [Input("privateLinkScopeResourceId")]
        public Input<string>? PrivateLinkScopeResourceId { get; set; }

        /// <summary>
        /// Gets or sets the guest agent provisioning action.
        /// </summary>
        [Input("provisioningAction")]
        public InputUnion<string, Pulumi.AzureNative.ConnectedVMwarevSphere.ProvisioningAction>? ProvisioningAction { get; set; }

        /// <summary>
        /// The Resource Group Name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of the vm.
        /// </summary>
        [Input("virtualMachineName", required: true)]
        public Input<string> VirtualMachineName { get; set; } = null!;

        public GuestAgentArgs()
        {
        }
        public static new GuestAgentArgs Empty => new GuestAgentArgs();
    }
}
