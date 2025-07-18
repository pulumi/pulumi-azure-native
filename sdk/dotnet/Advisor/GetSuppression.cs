// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Advisor
{
    public static class GetSuppression
    {
        /// <summary>
        /// Obtains the details of a suppression.
        /// 
        /// Uses Azure REST API version 2023-09-01-preview.
        /// 
        /// Other available API versions: 2023-01-01, 2024-11-18-preview, 2025-01-01, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native advisor [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetSuppressionResult> InvokeAsync(GetSuppressionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSuppressionResult>("azure-native:advisor:getSuppression", args ?? new GetSuppressionArgs(), options.WithDefaults());

        /// <summary>
        /// Obtains the details of a suppression.
        /// 
        /// Uses Azure REST API version 2023-09-01-preview.
        /// 
        /// Other available API versions: 2023-01-01, 2024-11-18-preview, 2025-01-01, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native advisor [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetSuppressionResult> Invoke(GetSuppressionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSuppressionResult>("azure-native:advisor:getSuppression", args ?? new GetSuppressionInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Obtains the details of a suppression.
        /// 
        /// Uses Azure REST API version 2023-09-01-preview.
        /// 
        /// Other available API versions: 2023-01-01, 2024-11-18-preview, 2025-01-01, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native advisor [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetSuppressionResult> Invoke(GetSuppressionInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetSuppressionResult>("azure-native:advisor:getSuppression", args ?? new GetSuppressionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSuppressionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the suppression.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The recommendation ID.
        /// </summary>
        [Input("recommendationId", required: true)]
        public string RecommendationId { get; set; } = null!;

        /// <summary>
        /// The fully qualified Azure Resource Manager identifier of the resource to which the recommendation applies.
        /// </summary>
        [Input("resourceUri", required: true)]
        public string ResourceUri { get; set; } = null!;

        public GetSuppressionArgs()
        {
        }
        public static new GetSuppressionArgs Empty => new GetSuppressionArgs();
    }

    public sealed class GetSuppressionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the suppression.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The recommendation ID.
        /// </summary>
        [Input("recommendationId", required: true)]
        public Input<string> RecommendationId { get; set; } = null!;

        /// <summary>
        /// The fully qualified Azure Resource Manager identifier of the resource to which the recommendation applies.
        /// </summary>
        [Input("resourceUri", required: true)]
        public Input<string> ResourceUri { get; set; } = null!;

        public GetSuppressionInvokeArgs()
        {
        }
        public static new GetSuppressionInvokeArgs Empty => new GetSuppressionInvokeArgs();
    }


    [OutputType]
    public sealed class GetSuppressionResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Gets or sets the expiration time stamp.
        /// </summary>
        public readonly string ExpirationTimeStamp;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The GUID of the suppression.
        /// </summary>
        public readonly string? SuppressionId;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The duration for which the suppression is valid.
        /// </summary>
        public readonly string? Ttl;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetSuppressionResult(
            string azureApiVersion,

            string expirationTimeStamp,

            string id,

            string name,

            string? suppressionId,

            Outputs.SystemDataResponse systemData,

            string? ttl,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            ExpirationTimeStamp = expirationTimeStamp;
            Id = id;
            Name = name;
            SuppressionId = suppressionId;
            SystemData = systemData;
            Ttl = ttl;
            Type = type;
        }
    }
}
