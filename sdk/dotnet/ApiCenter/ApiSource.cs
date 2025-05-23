// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ApiCenter
{
    /// <summary>
    /// API source entity.
    /// 
    /// Uses Azure REST API version 2024-06-01-preview. In version 2.x of the Azure Native provider, it used API version 2024-06-01-preview.
    /// </summary>
    [AzureNativeResourceType("azure-native:apicenter:ApiSource")]
    public partial class ApiSource : global::Pulumi.CustomResource
    {
        /// <summary>
        /// API source configuration for Azure API Management.
        /// </summary>
        [Output("azureApiManagementSource")]
        public Output<Outputs.AzureApiManagementSourceResponse?> AzureApiManagementSource { get; private set; } = null!;

        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Indicates if the specification should be imported along with metadata.
        /// </summary>
        [Output("importSpecification")]
        public Output<string?> ImportSpecification { get; private set; } = null!;

        /// <summary>
        /// The state of the API source link
        /// </summary>
        [Output("linkState")]
        public Output<Outputs.LinkStateResponse> LinkState { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// The target environment resource ID.
        /// </summary>
        [Output("targetEnvironmentId")]
        public Output<string?> TargetEnvironmentId { get; private set; } = null!;

        /// <summary>
        /// The target lifecycle stage.
        /// </summary>
        [Output("targetLifecycleStage")]
        public Output<string?> TargetLifecycleStage { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ApiSource resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApiSource(string name, ApiSourceArgs args, CustomResourceOptions? options = null)
            : base("azure-native:apicenter:ApiSource", name, args ?? new ApiSourceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApiSource(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:apicenter:ApiSource", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:apicenter/v20240601preview:ApiSource" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ApiSource resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApiSource Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ApiSource(name, id, options);
        }
    }

    public sealed class ApiSourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the API.
        /// </summary>
        [Input("apiSourceName")]
        public Input<string>? ApiSourceName { get; set; }

        /// <summary>
        /// API source configuration for Azure API Management.
        /// </summary>
        [Input("azureApiManagementSource")]
        public Input<Inputs.AzureApiManagementSourceArgs>? AzureApiManagementSource { get; set; }

        /// <summary>
        /// Indicates if the specification should be imported along with metadata.
        /// </summary>
        [Input("importSpecification")]
        public InputUnion<string, Pulumi.AzureNative.ApiCenter.ImportSpecificationOptions>? ImportSpecification { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of Azure API Center service.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        /// <summary>
        /// The target environment resource ID.
        /// </summary>
        [Input("targetEnvironmentId")]
        public Input<string>? TargetEnvironmentId { get; set; }

        /// <summary>
        /// The target lifecycle stage.
        /// </summary>
        [Input("targetLifecycleStage")]
        public InputUnion<string, Pulumi.AzureNative.ApiCenter.LifecycleStage>? TargetLifecycleStage { get; set; }

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public ApiSourceArgs()
        {
            ImportSpecification = "ondemand";
        }
        public static new ApiSourceArgs Empty => new ApiSourceArgs();
    }
}
