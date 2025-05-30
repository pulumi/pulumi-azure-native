// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.App
{
    public static class GetAppResiliency
    {
        /// <summary>
        /// Get container app resiliency policy.
        /// 
        /// Uses Azure REST API version 2024-10-02-preview.
        /// 
        /// Other available API versions: 2023-08-01-preview, 2023-11-02-preview, 2024-02-02-preview, 2024-08-02-preview, 2025-02-02-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native app [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetAppResiliencyResult> InvokeAsync(GetAppResiliencyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAppResiliencyResult>("azure-native:app:getAppResiliency", args ?? new GetAppResiliencyArgs(), options.WithDefaults());

        /// <summary>
        /// Get container app resiliency policy.
        /// 
        /// Uses Azure REST API version 2024-10-02-preview.
        /// 
        /// Other available API versions: 2023-08-01-preview, 2023-11-02-preview, 2024-02-02-preview, 2024-08-02-preview, 2025-02-02-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native app [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetAppResiliencyResult> Invoke(GetAppResiliencyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAppResiliencyResult>("azure-native:app:getAppResiliency", args ?? new GetAppResiliencyInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get container app resiliency policy.
        /// 
        /// Uses Azure REST API version 2024-10-02-preview.
        /// 
        /// Other available API versions: 2023-08-01-preview, 2023-11-02-preview, 2024-02-02-preview, 2024-08-02-preview, 2025-02-02-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native app [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetAppResiliencyResult> Invoke(GetAppResiliencyInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAppResiliencyResult>("azure-native:app:getAppResiliency", args ?? new GetAppResiliencyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAppResiliencyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the Container App.
        /// </summary>
        [Input("appName", required: true)]
        public string AppName { get; set; } = null!;

        /// <summary>
        /// Name of the resiliency policy.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetAppResiliencyArgs()
        {
        }
        public static new GetAppResiliencyArgs Empty => new GetAppResiliencyArgs();
    }

    public sealed class GetAppResiliencyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the Container App.
        /// </summary>
        [Input("appName", required: true)]
        public Input<string> AppName { get; set; } = null!;

        /// <summary>
        /// Name of the resiliency policy.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetAppResiliencyInvokeArgs()
        {
        }
        public static new GetAppResiliencyInvokeArgs Empty => new GetAppResiliencyInvokeArgs();
    }


    [OutputType]
    public sealed class GetAppResiliencyResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Policy that defines circuit breaker conditions
        /// </summary>
        public readonly Outputs.CircuitBreakerPolicyResponse? CircuitBreakerPolicy;
        /// <summary>
        /// Defines parameters for http connection pooling
        /// </summary>
        public readonly Outputs.HttpConnectionPoolResponse? HttpConnectionPool;
        /// <summary>
        /// Policy that defines http request retry conditions
        /// </summary>
        public readonly Outputs.HttpRetryPolicyResponse? HttpRetryPolicy;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Defines parameters for tcp connection pooling
        /// </summary>
        public readonly Outputs.TcpConnectionPoolResponse? TcpConnectionPool;
        /// <summary>
        /// Policy that defines tcp request retry conditions
        /// </summary>
        public readonly Outputs.TcpRetryPolicyResponse? TcpRetryPolicy;
        /// <summary>
        /// Policy to set request timeouts
        /// </summary>
        public readonly Outputs.TimeoutPolicyResponse? TimeoutPolicy;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetAppResiliencyResult(
            string azureApiVersion,

            Outputs.CircuitBreakerPolicyResponse? circuitBreakerPolicy,

            Outputs.HttpConnectionPoolResponse? httpConnectionPool,

            Outputs.HttpRetryPolicyResponse? httpRetryPolicy,

            string id,

            string name,

            Outputs.SystemDataResponse systemData,

            Outputs.TcpConnectionPoolResponse? tcpConnectionPool,

            Outputs.TcpRetryPolicyResponse? tcpRetryPolicy,

            Outputs.TimeoutPolicyResponse? timeoutPolicy,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            CircuitBreakerPolicy = circuitBreakerPolicy;
            HttpConnectionPool = httpConnectionPool;
            HttpRetryPolicy = httpRetryPolicy;
            Id = id;
            Name = name;
            SystemData = systemData;
            TcpConnectionPool = tcpConnectionPool;
            TcpRetryPolicy = tcpRetryPolicy;
            TimeoutPolicy = timeoutPolicy;
            Type = type;
        }
    }
}
