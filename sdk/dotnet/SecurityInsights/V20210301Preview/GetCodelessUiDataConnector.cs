// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecurityInsights.V20210301Preview
{
    public static class GetCodelessUiDataConnector
    {
        /// <summary>
        /// Gets a data connector.
        /// </summary>
        public static Task<GetCodelessUiDataConnectorResult> InvokeAsync(GetCodelessUiDataConnectorArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCodelessUiDataConnectorResult>("azure-native:securityinsights/v20210301preview:getCodelessUiDataConnector", args ?? new GetCodelessUiDataConnectorArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a data connector.
        /// </summary>
        public static Output<GetCodelessUiDataConnectorResult> Invoke(GetCodelessUiDataConnectorInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCodelessUiDataConnectorResult>("azure-native:securityinsights/v20210301preview:getCodelessUiDataConnector", args ?? new GetCodelessUiDataConnectorInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a data connector.
        /// </summary>
        public static Output<GetCodelessUiDataConnectorResult> Invoke(GetCodelessUiDataConnectorInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetCodelessUiDataConnectorResult>("azure-native:securityinsights/v20210301preview:getCodelessUiDataConnector", args ?? new GetCodelessUiDataConnectorInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCodelessUiDataConnectorArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Connector ID
        /// </summary>
        [Input("dataConnectorId", required: true)]
        public string DataConnectorId { get; set; } = null!;

        /// <summary>
        /// The namespace of workspaces resource provider- Microsoft.OperationalInsights.
        /// </summary>
        [Input("operationalInsightsResourceProvider", required: true)]
        public string OperationalInsightsResourceProvider { get; set; } = null!;

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

        public GetCodelessUiDataConnectorArgs()
        {
        }
        public static new GetCodelessUiDataConnectorArgs Empty => new GetCodelessUiDataConnectorArgs();
    }

    public sealed class GetCodelessUiDataConnectorInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Connector ID
        /// </summary>
        [Input("dataConnectorId", required: true)]
        public Input<string> DataConnectorId { get; set; } = null!;

        /// <summary>
        /// The namespace of workspaces resource provider- Microsoft.OperationalInsights.
        /// </summary>
        [Input("operationalInsightsResourceProvider", required: true)]
        public Input<string> OperationalInsightsResourceProvider { get; set; } = null!;

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

        public GetCodelessUiDataConnectorInvokeArgs()
        {
        }
        public static new GetCodelessUiDataConnectorInvokeArgs Empty => new GetCodelessUiDataConnectorInvokeArgs();
    }


    [OutputType]
    public sealed class GetCodelessUiDataConnectorResult
    {
        /// <summary>
        /// Config to describe the instructions blade
        /// </summary>
        public readonly Outputs.CodelessUiConnectorConfigPropertiesResponse? ConnectorUiConfig;
        /// <summary>
        /// Etag of the azure resource
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// Azure resource Id
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The kind of the data connector
        /// Expected value is 'GenericUI'.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Azure resource name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Azure resource type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetCodelessUiDataConnectorResult(
            Outputs.CodelessUiConnectorConfigPropertiesResponse? connectorUiConfig,

            string? etag,

            string id,

            string kind,

            string name,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            ConnectorUiConfig = connectorUiConfig;
            Etag = etag;
            Id = id;
            Kind = kind;
            Name = name;
            SystemData = systemData;
            Type = type;
        }
    }
}
