// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DeviceUpdate
{
    public static class GetPrivateEndpointConnectionProxy
    {
        /// <summary>
        /// (INTERNAL - DO NOT USE) Get the specified private endpoint connection proxy associated with the device update account.
        /// 
        /// Uses Azure REST API version 2023-07-01.
        /// </summary>
        public static Task<GetPrivateEndpointConnectionProxyResult> InvokeAsync(GetPrivateEndpointConnectionProxyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPrivateEndpointConnectionProxyResult>("azure-native:deviceupdate:getPrivateEndpointConnectionProxy", args ?? new GetPrivateEndpointConnectionProxyArgs(), options.WithDefaults());

        /// <summary>
        /// (INTERNAL - DO NOT USE) Get the specified private endpoint connection proxy associated with the device update account.
        /// 
        /// Uses Azure REST API version 2023-07-01.
        /// </summary>
        public static Output<GetPrivateEndpointConnectionProxyResult> Invoke(GetPrivateEndpointConnectionProxyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPrivateEndpointConnectionProxyResult>("azure-native:deviceupdate:getPrivateEndpointConnectionProxy", args ?? new GetPrivateEndpointConnectionProxyInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// (INTERNAL - DO NOT USE) Get the specified private endpoint connection proxy associated with the device update account.
        /// 
        /// Uses Azure REST API version 2023-07-01.
        /// </summary>
        public static Output<GetPrivateEndpointConnectionProxyResult> Invoke(GetPrivateEndpointConnectionProxyInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetPrivateEndpointConnectionProxyResult>("azure-native:deviceupdate:getPrivateEndpointConnectionProxy", args ?? new GetPrivateEndpointConnectionProxyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPrivateEndpointConnectionProxyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Account name.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The ID of the private endpoint connection proxy object.
        /// </summary>
        [Input("privateEndpointConnectionProxyId", required: true)]
        public string PrivateEndpointConnectionProxyId { get; set; } = null!;

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetPrivateEndpointConnectionProxyArgs()
        {
        }
        public static new GetPrivateEndpointConnectionProxyArgs Empty => new GetPrivateEndpointConnectionProxyArgs();
    }

    public sealed class GetPrivateEndpointConnectionProxyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Account name.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The ID of the private endpoint connection proxy object.
        /// </summary>
        [Input("privateEndpointConnectionProxyId", required: true)]
        public Input<string> PrivateEndpointConnectionProxyId { get; set; } = null!;

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetPrivateEndpointConnectionProxyInvokeArgs()
        {
        }
        public static new GetPrivateEndpointConnectionProxyInvokeArgs Empty => new GetPrivateEndpointConnectionProxyInvokeArgs();
    }


    [OutputType]
    public sealed class GetPrivateEndpointConnectionProxyResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// ETag from NRP.
        /// </summary>
        public readonly string ETag;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The provisioning state of the private endpoint connection proxy resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Remote private endpoint details.
        /// </summary>
        public readonly Outputs.RemotePrivateEndpointResponse? RemotePrivateEndpoint;
        /// <summary>
        /// Operation status.
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetPrivateEndpointConnectionProxyResult(
            string azureApiVersion,

            string eTag,

            string id,

            string name,

            string provisioningState,

            Outputs.RemotePrivateEndpointResponse? remotePrivateEndpoint,

            string? status,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            ETag = eTag;
            Id = id;
            Name = name;
            ProvisioningState = provisioningState;
            RemotePrivateEndpoint = remotePrivateEndpoint;
            Status = status;
            SystemData = systemData;
            Type = type;
        }
    }
}
