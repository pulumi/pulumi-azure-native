// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Solutions
{
    public static class GetJitRequest
    {
        /// <summary>
        /// Gets the JIT request.
        /// 
        /// Uses Azure REST API version 2021-07-01.
        /// 
        /// Other available API versions: 2023-12-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native solutions [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetJitRequestResult> InvokeAsync(GetJitRequestArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetJitRequestResult>("azure-native:solutions:getJitRequest", args ?? new GetJitRequestArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the JIT request.
        /// 
        /// Uses Azure REST API version 2021-07-01.
        /// 
        /// Other available API versions: 2023-12-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native solutions [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetJitRequestResult> Invoke(GetJitRequestInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetJitRequestResult>("azure-native:solutions:getJitRequest", args ?? new GetJitRequestInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the JIT request.
        /// 
        /// Uses Azure REST API version 2021-07-01.
        /// 
        /// Other available API versions: 2023-12-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native solutions [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetJitRequestResult> Invoke(GetJitRequestInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetJitRequestResult>("azure-native:solutions:getJitRequest", args ?? new GetJitRequestInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetJitRequestArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the JIT request.
        /// </summary>
        [Input("jitRequestName", required: true)]
        public string JitRequestName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetJitRequestArgs()
        {
        }
        public static new GetJitRequestArgs Empty => new GetJitRequestArgs();
    }

    public sealed class GetJitRequestInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the JIT request.
        /// </summary>
        [Input("jitRequestName", required: true)]
        public Input<string> JitRequestName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetJitRequestInvokeArgs()
        {
        }
        public static new GetJitRequestInvokeArgs Empty => new GetJitRequestInvokeArgs();
    }


    [OutputType]
    public sealed class GetJitRequestResult
    {
        /// <summary>
        /// The parent application id.
        /// </summary>
        public readonly string ApplicationResourceId;
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// The client entity that created the JIT request.
        /// </summary>
        public readonly Outputs.ApplicationClientDetailsResponse CreatedBy;
        /// <summary>
        /// Resource ID
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The JIT authorization policies.
        /// </summary>
        public readonly ImmutableArray<Outputs.JitAuthorizationPoliciesResponse> JitAuthorizationPolicies;
        /// <summary>
        /// The JIT request state.
        /// </summary>
        public readonly string JitRequestState;
        /// <summary>
        /// The JIT request properties.
        /// </summary>
        public readonly Outputs.JitSchedulingPolicyResponse JitSchedulingPolicy;
        /// <summary>
        /// Resource location
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Resource name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The JIT request provisioning state.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The publisher tenant id.
        /// </summary>
        public readonly string PublisherTenantId;
        /// <summary>
        /// Metadata pertaining to creation and last modification of the resource.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Resource tags
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The client entity that last updated the JIT request.
        /// </summary>
        public readonly Outputs.ApplicationClientDetailsResponse UpdatedBy;

        [OutputConstructor]
        private GetJitRequestResult(
            string applicationResourceId,

            string azureApiVersion,

            Outputs.ApplicationClientDetailsResponse createdBy,

            string id,

            ImmutableArray<Outputs.JitAuthorizationPoliciesResponse> jitAuthorizationPolicies,

            string jitRequestState,

            Outputs.JitSchedulingPolicyResponse jitSchedulingPolicy,

            string? location,

            string name,

            string provisioningState,

            string publisherTenantId,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type,

            Outputs.ApplicationClientDetailsResponse updatedBy)
        {
            ApplicationResourceId = applicationResourceId;
            AzureApiVersion = azureApiVersion;
            CreatedBy = createdBy;
            Id = id;
            JitAuthorizationPolicies = jitAuthorizationPolicies;
            JitRequestState = jitRequestState;
            JitSchedulingPolicy = jitSchedulingPolicy;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            PublisherTenantId = publisherTenantId;
            SystemData = systemData;
            Tags = tags;
            Type = type;
            UpdatedBy = updatedBy;
        }
    }
}
