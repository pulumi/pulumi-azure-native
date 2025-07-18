// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SignalRService
{
    /// <summary>
    /// Describes a Shared Private Link Resource
    /// 
    /// Uses Azure REST API version 2024-03-01. In version 2.x of the Azure Native provider, it used API version 2023-02-01.
    /// 
    /// Other available API versions: 2023-02-01, 2023-03-01-preview, 2023-06-01-preview, 2023-08-01-preview, 2024-01-01-preview, 2024-04-01-preview, 2024-08-01-preview, 2024-10-01-preview, 2025-01-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native signalrservice [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:signalrservice:SignalRSharedPrivateLinkResource")]
    public partial class SignalRSharedPrivateLinkResource : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// The group id from the provider of resource the shared private link resource is for
        /// </summary>
        [Output("groupId")]
        public Output<string> GroupId { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The resource id of the resource the shared private link resource is for
        /// </summary>
        [Output("privateLinkResourceId")]
        public Output<string> PrivateLinkResourceId { get; private set; } = null!;

        /// <summary>
        /// Provisioning state of the resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The request message for requesting approval of the shared private link resource
        /// </summary>
        [Output("requestMessage")]
        public Output<string?> RequestMessage { get; private set; } = null!;

        /// <summary>
        /// Status of the shared private link resource
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

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
        /// Create a SignalRSharedPrivateLinkResource resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SignalRSharedPrivateLinkResource(string name, SignalRSharedPrivateLinkResourceArgs args, CustomResourceOptions? options = null)
            : base("azure-native:signalrservice:SignalRSharedPrivateLinkResource", name, args ?? new SignalRSharedPrivateLinkResourceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SignalRSharedPrivateLinkResource(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:signalrservice:SignalRSharedPrivateLinkResource", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:signalrservice/v20210401preview:SignalRSharedPrivateLinkResource" },
                    new global::Pulumi.Alias { Type = "azure-native:signalrservice/v20210601preview:SignalRSharedPrivateLinkResource" },
                    new global::Pulumi.Alias { Type = "azure-native:signalrservice/v20210901preview:SignalRSharedPrivateLinkResource" },
                    new global::Pulumi.Alias { Type = "azure-native:signalrservice/v20211001:SignalRSharedPrivateLinkResource" },
                    new global::Pulumi.Alias { Type = "azure-native:signalrservice/v20220201:SignalRSharedPrivateLinkResource" },
                    new global::Pulumi.Alias { Type = "azure-native:signalrservice/v20220801preview:SignalRSharedPrivateLinkResource" },
                    new global::Pulumi.Alias { Type = "azure-native:signalrservice/v20230201:SignalRSharedPrivateLinkResource" },
                    new global::Pulumi.Alias { Type = "azure-native:signalrservice/v20230301preview:SignalRSharedPrivateLinkResource" },
                    new global::Pulumi.Alias { Type = "azure-native:signalrservice/v20230601preview:SignalRSharedPrivateLinkResource" },
                    new global::Pulumi.Alias { Type = "azure-native:signalrservice/v20230801preview:SignalRSharedPrivateLinkResource" },
                    new global::Pulumi.Alias { Type = "azure-native:signalrservice/v20240101preview:SignalRSharedPrivateLinkResource" },
                    new global::Pulumi.Alias { Type = "azure-native:signalrservice/v20240301:SignalRSharedPrivateLinkResource" },
                    new global::Pulumi.Alias { Type = "azure-native:signalrservice/v20240401preview:SignalRSharedPrivateLinkResource" },
                    new global::Pulumi.Alias { Type = "azure-native:signalrservice/v20240801preview:SignalRSharedPrivateLinkResource" },
                    new global::Pulumi.Alias { Type = "azure-native:signalrservice/v20241001preview:SignalRSharedPrivateLinkResource" },
                    new global::Pulumi.Alias { Type = "azure-native:signalrservice/v20250101preview:SignalRSharedPrivateLinkResource" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SignalRSharedPrivateLinkResource resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SignalRSharedPrivateLinkResource Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new SignalRSharedPrivateLinkResource(name, id, options);
        }
    }

    public sealed class SignalRSharedPrivateLinkResourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The group id from the provider of resource the shared private link resource is for
        /// </summary>
        [Input("groupId", required: true)]
        public Input<string> GroupId { get; set; } = null!;

        /// <summary>
        /// The resource id of the resource the shared private link resource is for
        /// </summary>
        [Input("privateLinkResourceId", required: true)]
        public Input<string> PrivateLinkResourceId { get; set; } = null!;

        /// <summary>
        /// The request message for requesting approval of the shared private link resource
        /// </summary>
        [Input("requestMessage")]
        public Input<string>? RequestMessage { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Input("resourceName", required: true)]
        public Input<string> ResourceName { get; set; } = null!;

        /// <summary>
        /// The name of the shared private link resource.
        /// </summary>
        [Input("sharedPrivateLinkResourceName")]
        public Input<string>? SharedPrivateLinkResourceName { get; set; }

        public SignalRSharedPrivateLinkResourceArgs()
        {
        }
        public static new SignalRSharedPrivateLinkResourceArgs Empty => new SignalRSharedPrivateLinkResourceArgs();
    }
}
