// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AVS
{
    public static class GetAuthorization
    {
        /// <summary>
        /// ExpressRoute Circuit Authorization
        /// 
        /// Uses Azure REST API version 2022-05-01.
        /// 
        /// Other available API versions: 2023-03-01, 2023-09-01.
        /// </summary>
        public static Task<GetAuthorizationResult> InvokeAsync(GetAuthorizationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAuthorizationResult>("azure-native:avs:getAuthorization", args ?? new GetAuthorizationArgs(), options.WithDefaults());

        /// <summary>
        /// ExpressRoute Circuit Authorization
        /// 
        /// Uses Azure REST API version 2022-05-01.
        /// 
        /// Other available API versions: 2023-03-01, 2023-09-01.
        /// </summary>
        public static Output<GetAuthorizationResult> Invoke(GetAuthorizationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAuthorizationResult>("azure-native:avs:getAuthorization", args ?? new GetAuthorizationInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// ExpressRoute Circuit Authorization
        /// 
        /// Uses Azure REST API version 2022-05-01.
        /// 
        /// Other available API versions: 2023-03-01, 2023-09-01.
        /// </summary>
        public static Output<GetAuthorizationResult> Invoke(GetAuthorizationInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAuthorizationResult>("azure-native:avs:getAuthorization", args ?? new GetAuthorizationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAuthorizationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the ExpressRoute Circuit Authorization in the private cloud
        /// </summary>
        [Input("authorizationName", required: true)]
        public string AuthorizationName { get; set; } = null!;

        /// <summary>
        /// Name of the private cloud
        /// </summary>
        [Input("privateCloudName", required: true)]
        public string PrivateCloudName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetAuthorizationArgs()
        {
        }
        public static new GetAuthorizationArgs Empty => new GetAuthorizationArgs();
    }

    public sealed class GetAuthorizationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the ExpressRoute Circuit Authorization in the private cloud
        /// </summary>
        [Input("authorizationName", required: true)]
        public Input<string> AuthorizationName { get; set; } = null!;

        /// <summary>
        /// Name of the private cloud
        /// </summary>
        [Input("privateCloudName", required: true)]
        public Input<string> PrivateCloudName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetAuthorizationInvokeArgs()
        {
        }
        public static new GetAuthorizationInvokeArgs Empty => new GetAuthorizationInvokeArgs();
    }


    [OutputType]
    public sealed class GetAuthorizationResult
    {
        /// <summary>
        /// The ID of the ExpressRoute Circuit Authorization
        /// </summary>
        public readonly string ExpressRouteAuthorizationId;
        /// <summary>
        /// The key of the ExpressRoute Circuit Authorization
        /// </summary>
        public readonly string ExpressRouteAuthorizationKey;
        /// <summary>
        /// The ID of the ExpressRoute Circuit
        /// </summary>
        public readonly string? ExpressRouteId;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The state of the  ExpressRoute Circuit Authorization provisioning
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetAuthorizationResult(
            string expressRouteAuthorizationId,

            string expressRouteAuthorizationKey,

            string? expressRouteId,

            string id,

            string name,

            string provisioningState,

            string type)
        {
            ExpressRouteAuthorizationId = expressRouteAuthorizationId;
            ExpressRouteAuthorizationKey = expressRouteAuthorizationKey;
            ExpressRouteId = expressRouteId;
            Id = id;
            Name = name;
            ProvisioningState = provisioningState;
            Type = type;
        }
    }
}
