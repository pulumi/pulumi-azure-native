// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataShare
{
    public static class GetSynapseWorkspaceSqlPoolTableDataSetMapping
    {
        /// <summary>
        /// Get a DataSetMapping in a shareSubscription
        /// 
        /// Uses Azure REST API version 2021-08-01.
        /// </summary>
        public static Task<GetSynapseWorkspaceSqlPoolTableDataSetMappingResult> InvokeAsync(GetSynapseWorkspaceSqlPoolTableDataSetMappingArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSynapseWorkspaceSqlPoolTableDataSetMappingResult>("azure-native:datashare:getSynapseWorkspaceSqlPoolTableDataSetMapping", args ?? new GetSynapseWorkspaceSqlPoolTableDataSetMappingArgs(), options.WithDefaults());

        /// <summary>
        /// Get a DataSetMapping in a shareSubscription
        /// 
        /// Uses Azure REST API version 2021-08-01.
        /// </summary>
        public static Output<GetSynapseWorkspaceSqlPoolTableDataSetMappingResult> Invoke(GetSynapseWorkspaceSqlPoolTableDataSetMappingInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSynapseWorkspaceSqlPoolTableDataSetMappingResult>("azure-native:datashare:getSynapseWorkspaceSqlPoolTableDataSetMapping", args ?? new GetSynapseWorkspaceSqlPoolTableDataSetMappingInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get a DataSetMapping in a shareSubscription
        /// 
        /// Uses Azure REST API version 2021-08-01.
        /// </summary>
        public static Output<GetSynapseWorkspaceSqlPoolTableDataSetMappingResult> Invoke(GetSynapseWorkspaceSqlPoolTableDataSetMappingInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetSynapseWorkspaceSqlPoolTableDataSetMappingResult>("azure-native:datashare:getSynapseWorkspaceSqlPoolTableDataSetMapping", args ?? new GetSynapseWorkspaceSqlPoolTableDataSetMappingInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSynapseWorkspaceSqlPoolTableDataSetMappingArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the share account.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the dataSetMapping.
        /// </summary>
        [Input("dataSetMappingName", required: true)]
        public string DataSetMappingName { get; set; } = null!;

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the shareSubscription.
        /// </summary>
        [Input("shareSubscriptionName", required: true)]
        public string ShareSubscriptionName { get; set; } = null!;

        public GetSynapseWorkspaceSqlPoolTableDataSetMappingArgs()
        {
        }
        public static new GetSynapseWorkspaceSqlPoolTableDataSetMappingArgs Empty => new GetSynapseWorkspaceSqlPoolTableDataSetMappingArgs();
    }

    public sealed class GetSynapseWorkspaceSqlPoolTableDataSetMappingInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the share account.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the dataSetMapping.
        /// </summary>
        [Input("dataSetMappingName", required: true)]
        public Input<string> DataSetMappingName { get; set; } = null!;

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the shareSubscription.
        /// </summary>
        [Input("shareSubscriptionName", required: true)]
        public Input<string> ShareSubscriptionName { get; set; } = null!;

        public GetSynapseWorkspaceSqlPoolTableDataSetMappingInvokeArgs()
        {
        }
        public static new GetSynapseWorkspaceSqlPoolTableDataSetMappingInvokeArgs Empty => new GetSynapseWorkspaceSqlPoolTableDataSetMappingInvokeArgs();
    }


    [OutputType]
    public sealed class GetSynapseWorkspaceSqlPoolTableDataSetMappingResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// The id of the source data set.
        /// </summary>
        public readonly string DataSetId;
        /// <summary>
        /// Gets the status of the data set mapping.
        /// </summary>
        public readonly string DataSetMappingStatus;
        /// <summary>
        /// The resource id of the azure resource
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Kind of data set mapping.
        /// Expected value is 'SynapseWorkspaceSqlPoolTable'.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Name of the azure resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Provisioning state of the data set mapping.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Resource id of the Synapse Workspace SQL Pool Table
        /// </summary>
        public readonly string SynapseWorkspaceSqlPoolTableResourceId;
        /// <summary>
        /// System Data of the Azure resource.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Type of the azure resource
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetSynapseWorkspaceSqlPoolTableDataSetMappingResult(
            string azureApiVersion,

            string dataSetId,

            string dataSetMappingStatus,

            string id,

            string kind,

            string name,

            string provisioningState,

            string synapseWorkspaceSqlPoolTableResourceId,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            DataSetId = dataSetId;
            DataSetMappingStatus = dataSetMappingStatus;
            Id = id;
            Kind = kind;
            Name = name;
            ProvisioningState = provisioningState;
            SynapseWorkspaceSqlPoolTableResourceId = synapseWorkspaceSqlPoolTableResourceId;
            SystemData = systemData;
            Type = type;
        }
    }
}
