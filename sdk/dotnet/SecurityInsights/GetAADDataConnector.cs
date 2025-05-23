// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecurityInsights
{
    public static class GetAADDataConnector
    {
        /// <summary>
        /// Gets a data connector.
        /// 
        /// Uses Azure REST API version 2024-09-01.
        /// </summary>
        public static Task<GetAADDataConnectorResult> InvokeAsync(GetAADDataConnectorArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAADDataConnectorResult>("azure-native:securityinsights:getAADDataConnector", args ?? new GetAADDataConnectorArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a data connector.
        /// 
        /// Uses Azure REST API version 2024-09-01.
        /// </summary>
        public static Output<GetAADDataConnectorResult> Invoke(GetAADDataConnectorInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAADDataConnectorResult>("azure-native:securityinsights:getAADDataConnector", args ?? new GetAADDataConnectorInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a data connector.
        /// 
        /// Uses Azure REST API version 2024-09-01.
        /// </summary>
        public static Output<GetAADDataConnectorResult> Invoke(GetAADDataConnectorInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAADDataConnectorResult>("azure-native:securityinsights:getAADDataConnector", args ?? new GetAADDataConnectorInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAADDataConnectorArgs : global::Pulumi.InvokeArgs
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

        public GetAADDataConnectorArgs()
        {
        }
        public static new GetAADDataConnectorArgs Empty => new GetAADDataConnectorArgs();
    }

    public sealed class GetAADDataConnectorInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetAADDataConnectorInvokeArgs()
        {
        }
        public static new GetAADDataConnectorInvokeArgs Empty => new GetAADDataConnectorInvokeArgs();
    }


    [OutputType]
    public sealed class GetAADDataConnectorResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
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
        /// Expected value is 'AzureActiveDirectory'.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The tenant id to connect to, and get the data from.
        /// </summary>
        public readonly string? TenantId;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetAADDataConnectorResult(
            string azureApiVersion,

            Outputs.AlertsDataTypeOfDataConnectorResponse? dataTypes,

            string? etag,

            string id,

            string kind,

            string name,

            Outputs.SystemDataResponse systemData,

            string? tenantId,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            DataTypes = dataTypes;
            Etag = etag;
            Id = id;
            Kind = kind;
            Name = name;
            SystemData = systemData;
            TenantId = tenantId;
            Type = type;
        }
    }
}
