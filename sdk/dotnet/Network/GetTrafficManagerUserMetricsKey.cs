// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network
{
    public static class GetTrafficManagerUserMetricsKey
    {
        /// <summary>
        /// Get the subscription-level key used for Real User Metrics collection.
        /// 
        /// Uses Azure REST API version 2022-04-01.
        /// 
        /// Other available API versions: 2022-04-01-preview.
        /// </summary>
        public static Task<GetTrafficManagerUserMetricsKeyResult> InvokeAsync(GetTrafficManagerUserMetricsKeyArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTrafficManagerUserMetricsKeyResult>("azure-native:network:getTrafficManagerUserMetricsKey", args ?? new GetTrafficManagerUserMetricsKeyArgs(), options.WithDefaults());

        /// <summary>
        /// Get the subscription-level key used for Real User Metrics collection.
        /// 
        /// Uses Azure REST API version 2022-04-01.
        /// 
        /// Other available API versions: 2022-04-01-preview.
        /// </summary>
        public static Output<GetTrafficManagerUserMetricsKeyResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTrafficManagerUserMetricsKeyResult>("azure-native:network:getTrafficManagerUserMetricsKey", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// Get the subscription-level key used for Real User Metrics collection.
        /// 
        /// Uses Azure REST API version 2022-04-01.
        /// 
        /// Other available API versions: 2022-04-01-preview.
        /// </summary>
        public static Output<GetTrafficManagerUserMetricsKeyResult> Invoke(InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetTrafficManagerUserMetricsKeyResult>("azure-native:network:getTrafficManagerUserMetricsKey", InvokeArgs.Empty, options.WithDefaults());
    }


    public sealed class GetTrafficManagerUserMetricsKeyArgs : global::Pulumi.InvokeArgs
    {
        public GetTrafficManagerUserMetricsKeyArgs()
        {
        }
        public static new GetTrafficManagerUserMetricsKeyArgs Empty => new GetTrafficManagerUserMetricsKeyArgs();
    }


    [OutputType]
    public sealed class GetTrafficManagerUserMetricsKeyResult
    {
        /// <summary>
        /// Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficManagerProfiles/{resourceName}
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The key returned by the User Metrics operation.
        /// </summary>
        public readonly string? Key;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private GetTrafficManagerUserMetricsKeyResult(
            string? id,

            string? key,

            string? name,

            string? type)
        {
            Id = id;
            Key = key;
            Name = name;
            Type = type;
        }
    }
}
