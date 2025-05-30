// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Chaos
{
    public static class GetTarget
    {
        /// <summary>
        /// Get a Target resource that extends a tracked regional resource.
        /// 
        /// Uses Azure REST API version 2024-03-22-preview.
        /// 
        /// Other available API versions: 2023-04-15-preview, 2023-09-01-preview, 2023-10-27-preview, 2023-11-01, 2024-01-01, 2024-11-01-preview, 2025-01-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native chaos [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetTargetResult> InvokeAsync(GetTargetArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTargetResult>("azure-native:chaos:getTarget", args ?? new GetTargetArgs(), options.WithDefaults());

        /// <summary>
        /// Get a Target resource that extends a tracked regional resource.
        /// 
        /// Uses Azure REST API version 2024-03-22-preview.
        /// 
        /// Other available API versions: 2023-04-15-preview, 2023-09-01-preview, 2023-10-27-preview, 2023-11-01, 2024-01-01, 2024-11-01-preview, 2025-01-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native chaos [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetTargetResult> Invoke(GetTargetInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTargetResult>("azure-native:chaos:getTarget", args ?? new GetTargetInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get a Target resource that extends a tracked regional resource.
        /// 
        /// Uses Azure REST API version 2024-03-22-preview.
        /// 
        /// Other available API versions: 2023-04-15-preview, 2023-09-01-preview, 2023-10-27-preview, 2023-11-01, 2024-01-01, 2024-11-01-preview, 2025-01-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native chaos [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetTargetResult> Invoke(GetTargetInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetTargetResult>("azure-native:chaos:getTarget", args ?? new GetTargetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTargetArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// String that represents a resource provider namespace.
        /// </summary>
        [Input("parentProviderNamespace", required: true)]
        public string ParentProviderNamespace { get; set; } = null!;

        /// <summary>
        /// String that represents a resource name.
        /// </summary>
        [Input("parentResourceName", required: true)]
        public string ParentResourceName { get; set; } = null!;

        /// <summary>
        /// String that represents a resource type.
        /// </summary>
        [Input("parentResourceType", required: true)]
        public string ParentResourceType { get; set; } = null!;

        /// <summary>
        /// String that represents an Azure resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// String that represents a Target resource name.
        /// </summary>
        [Input("targetName", required: true)]
        public string TargetName { get; set; } = null!;

        public GetTargetArgs()
        {
        }
        public static new GetTargetArgs Empty => new GetTargetArgs();
    }

    public sealed class GetTargetInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// String that represents a resource provider namespace.
        /// </summary>
        [Input("parentProviderNamespace", required: true)]
        public Input<string> ParentProviderNamespace { get; set; } = null!;

        /// <summary>
        /// String that represents a resource name.
        /// </summary>
        [Input("parentResourceName", required: true)]
        public Input<string> ParentResourceName { get; set; } = null!;

        /// <summary>
        /// String that represents a resource type.
        /// </summary>
        [Input("parentResourceType", required: true)]
        public Input<string> ParentResourceType { get; set; } = null!;

        /// <summary>
        /// String that represents an Azure resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// String that represents a Target resource name.
        /// </summary>
        [Input("targetName", required: true)]
        public Input<string> TargetName { get; set; } = null!;

        public GetTargetInvokeArgs()
        {
        }
        public static new GetTargetInvokeArgs Empty => new GetTargetInvokeArgs();
    }


    [OutputType]
    public sealed class GetTargetResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Location of the target resource.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The properties of the target resource.
        /// </summary>
        public readonly object Properties;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetTargetResult(
            string azureApiVersion,

            string id,

            string? location,

            string name,

            object properties,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Id = id;
            Location = location;
            Name = name;
            Properties = properties;
            SystemData = systemData;
            Type = type;
        }
    }
}
