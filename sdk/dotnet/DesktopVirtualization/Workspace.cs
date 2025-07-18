// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DesktopVirtualization
{
    /// <summary>
    /// Represents a Workspace definition.
    /// 
    /// Uses Azure REST API version 2024-04-03. In version 2.x of the Azure Native provider, it used API version 2022-09-09.
    /// 
    /// Other available API versions: 2022-09-09, 2022-10-14-preview, 2023-09-05, 2023-10-04-preview, 2023-11-01-preview, 2024-01-16-preview, 2024-03-06-preview, 2024-04-08-preview, 2024-08-08-preview, 2024-11-01-preview, 2025-03-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native desktopvirtualization [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:desktopvirtualization:Workspace")]
    public partial class Workspace : global::Pulumi.CustomResource
    {
        /// <summary>
        /// List of applicationGroup resource Ids.
        /// </summary>
        [Output("applicationGroupReferences")]
        public Output<ImmutableArray<string>> ApplicationGroupReferences { get; private set; } = null!;

        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Is cloud pc resource.
        /// </summary>
        [Output("cloudPcResource")]
        public Output<bool> CloudPcResource { get; private set; } = null!;

        /// <summary>
        /// Description of Workspace.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields. 
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Friendly name of Workspace.
        /// </summary>
        [Output("friendlyName")]
        public Output<string?> FriendlyName { get; private set; } = null!;

        [Output("identity")]
        public Output<Outputs.ResourceModelWithAllowedPropertySetResponseIdentity?> Identity { get; private set; } = null!;

        /// <summary>
        /// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type. E.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
        /// </summary>
        [Output("kind")]
        public Output<string?> Kind { get; private set; } = null!;

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The fully qualified resource ID of the resource that manages this resource. Indicates if this resource is managed by another Azure resource. If this is present, complete mode deployment will not delete the resource if it is removed from the template since it is managed by another resource.
        /// </summary>
        [Output("managedBy")]
        public Output<string?> ManagedBy { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// ObjectId of Workspace. (internal use)
        /// </summary>
        [Output("objectId")]
        public Output<string> ObjectId { get; private set; } = null!;

        [Output("plan")]
        public Output<Outputs.ResourceModelWithAllowedPropertySetResponsePlan?> Plan { get; private set; } = null!;

        /// <summary>
        /// List of private endpoint connection associated with the specified resource
        /// </summary>
        [Output("privateEndpointConnections")]
        public Output<ImmutableArray<Outputs.PrivateEndpointConnectionResponse>> PrivateEndpointConnections { get; private set; } = null!;

        /// <summary>
        /// Enabled allows this resource to be accessed from both public and private networks, Disabled allows this resource to only be accessed via private endpoints
        /// </summary>
        [Output("publicNetworkAccess")]
        public Output<string?> PublicNetworkAccess { get; private set; } = null!;

        [Output("sku")]
        public Output<Outputs.ResourceModelWithAllowedPropertySetResponseSku?> Sku { get; private set; } = null!;

        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Workspace resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Workspace(string name, WorkspaceArgs args, CustomResourceOptions? options = null)
            : base("azure-native:desktopvirtualization:Workspace", name, args ?? new WorkspaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Workspace(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:desktopvirtualization:Workspace", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:desktopvirtualization/v20190123preview:Workspace" },
                    new global::Pulumi.Alias { Type = "azure-native:desktopvirtualization/v20190924preview:Workspace" },
                    new global::Pulumi.Alias { Type = "azure-native:desktopvirtualization/v20191210preview:Workspace" },
                    new global::Pulumi.Alias { Type = "azure-native:desktopvirtualization/v20200921preview:Workspace" },
                    new global::Pulumi.Alias { Type = "azure-native:desktopvirtualization/v20201019preview:Workspace" },
                    new global::Pulumi.Alias { Type = "azure-native:desktopvirtualization/v20201102preview:Workspace" },
                    new global::Pulumi.Alias { Type = "azure-native:desktopvirtualization/v20201110preview:Workspace" },
                    new global::Pulumi.Alias { Type = "azure-native:desktopvirtualization/v20210114preview:Workspace" },
                    new global::Pulumi.Alias { Type = "azure-native:desktopvirtualization/v20210201preview:Workspace" },
                    new global::Pulumi.Alias { Type = "azure-native:desktopvirtualization/v20210309preview:Workspace" },
                    new global::Pulumi.Alias { Type = "azure-native:desktopvirtualization/v20210401preview:Workspace" },
                    new global::Pulumi.Alias { Type = "azure-native:desktopvirtualization/v20210712:Workspace" },
                    new global::Pulumi.Alias { Type = "azure-native:desktopvirtualization/v20210903preview:Workspace" },
                    new global::Pulumi.Alias { Type = "azure-native:desktopvirtualization/v20220210preview:Workspace" },
                    new global::Pulumi.Alias { Type = "azure-native:desktopvirtualization/v20220401preview:Workspace" },
                    new global::Pulumi.Alias { Type = "azure-native:desktopvirtualization/v20220909:Workspace" },
                    new global::Pulumi.Alias { Type = "azure-native:desktopvirtualization/v20221014preview:Workspace" },
                    new global::Pulumi.Alias { Type = "azure-native:desktopvirtualization/v20230707preview:Workspace" },
                    new global::Pulumi.Alias { Type = "azure-native:desktopvirtualization/v20230905:Workspace" },
                    new global::Pulumi.Alias { Type = "azure-native:desktopvirtualization/v20231004preview:Workspace" },
                    new global::Pulumi.Alias { Type = "azure-native:desktopvirtualization/v20231101preview:Workspace" },
                    new global::Pulumi.Alias { Type = "azure-native:desktopvirtualization/v20240116preview:Workspace" },
                    new global::Pulumi.Alias { Type = "azure-native:desktopvirtualization/v20240306preview:Workspace" },
                    new global::Pulumi.Alias { Type = "azure-native:desktopvirtualization/v20240403:Workspace" },
                    new global::Pulumi.Alias { Type = "azure-native:desktopvirtualization/v20240408preview:Workspace" },
                    new global::Pulumi.Alias { Type = "azure-native:desktopvirtualization/v20240808preview:Workspace" },
                    new global::Pulumi.Alias { Type = "azure-native:desktopvirtualization/v20241101preview:Workspace" },
                    new global::Pulumi.Alias { Type = "azure-native:desktopvirtualization/v20250301preview:Workspace" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Workspace resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Workspace Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Workspace(name, id, options);
        }
    }

    public sealed class WorkspaceArgs : global::Pulumi.ResourceArgs
    {
        [Input("applicationGroupReferences")]
        private InputList<string>? _applicationGroupReferences;

        /// <summary>
        /// List of applicationGroup resource Ids.
        /// </summary>
        public InputList<string> ApplicationGroupReferences
        {
            get => _applicationGroupReferences ?? (_applicationGroupReferences = new InputList<string>());
            set => _applicationGroupReferences = value;
        }

        /// <summary>
        /// Description of Workspace.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Friendly name of Workspace.
        /// </summary>
        [Input("friendlyName")]
        public Input<string>? FriendlyName { get; set; }

        [Input("identity")]
        public Input<Inputs.ResourceModelWithAllowedPropertySetIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type. E.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The fully qualified resource ID of the resource that manages this resource. Indicates if this resource is managed by another Azure resource. If this is present, complete mode deployment will not delete the resource if it is removed from the template since it is managed by another resource.
        /// </summary>
        [Input("managedBy")]
        public Input<string>? ManagedBy { get; set; }

        [Input("plan")]
        public Input<Inputs.ResourceModelWithAllowedPropertySetPlanArgs>? Plan { get; set; }

        /// <summary>
        /// Enabled allows this resource to be accessed from both public and private networks, Disabled allows this resource to only be accessed via private endpoints
        /// </summary>
        [Input("publicNetworkAccess")]
        public InputUnion<string, Pulumi.AzureNative.DesktopVirtualization.PublicNetworkAccess>? PublicNetworkAccess { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("sku")]
        public Input<Inputs.ResourceModelWithAllowedPropertySetSkuArgs>? Sku { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The name of the workspace
        /// </summary>
        [Input("workspaceName")]
        public Input<string>? WorkspaceName { get; set; }

        public WorkspaceArgs()
        {
        }
        public static new WorkspaceArgs Empty => new WorkspaceArgs();
    }
}
