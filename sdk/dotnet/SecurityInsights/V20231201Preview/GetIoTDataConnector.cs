// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecurityInsights.V20231201Preview
{
    public static class GetIoTDataConnector
    {
        /// <summary>
        /// Gets a data connector.
        /// </summary>
        public static Task<GetIoTDataConnectorResult> InvokeAsync(GetIoTDataConnectorArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIoTDataConnectorResult>("azure-native:securityinsights/v20231201preview:getIoTDataConnector", args ?? new GetIoTDataConnectorArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a data connector.
        /// </summary>
        public static Output<GetIoTDataConnectorResult> Invoke(GetIoTDataConnectorInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIoTDataConnectorResult>("azure-native:securityinsights/v20231201preview:getIoTDataConnector", args ?? new GetIoTDataConnectorInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a data connector.
        /// </summary>
        public static Output<GetIoTDataConnectorResult> Invoke(GetIoTDataConnectorInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetIoTDataConnectorResult>("azure-native:securityinsights/v20231201preview:getIoTDataConnector", args ?? new GetIoTDataConnectorInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIoTDataConnectorArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Connector ID
        /// </summary>
        [Input("dataConnectorId", required: true)]
        public string DataConnectorId { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public string WorkspaceName { get; set; } = null!;

        public GetIoTDataConnectorArgs()
        {
        }
        public static new GetIoTDataConnectorArgs Empty => new GetIoTDataConnectorArgs();
    }

    public sealed class GetIoTDataConnectorInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Connector ID
        /// </summary>
        [Input("dataConnectorId", required: true)]
        public Input<string> DataConnectorId { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public GetIoTDataConnectorInvokeArgs()
        {
        }
        public static new GetIoTDataConnectorInvokeArgs Empty => new GetIoTDataConnectorInvokeArgs();
    }


    [OutputType]
    public sealed class GetIoTDataConnectorResult
    {
        /// <summary>
        /// The available data types for the connector.
        /// </summary>
        public readonly Outputs.AlertsDataTypeOfDataConnectorResponse? DataTypes;
        /// <summary>
        /// Etag of the azure resource
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The kind of the data connector
        /// Expected value is 'IOT'.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The subscription id to connect to, and get the data from.
        /// </summary>
        public readonly string? SubscriptionId;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetIoTDataConnectorResult(
            Outputs.AlertsDataTypeOfDataConnectorResponse? dataTypes,

            string? etag,

            string id,

            string kind,

            string name,

            string? subscriptionId,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            DataTypes = dataTypes;
            Etag = etag;
            Id = id;
            Kind = kind;
            Name = name;
            SubscriptionId = subscriptionId;
            SystemData = systemData;
            Type = type;
        }
    }
}
