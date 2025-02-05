// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.NetworkAnalytics
{
    public static class GetDataType
    {
        /// <summary>
        /// Retrieve data type resource.
        /// Azure REST API version: 2023-11-15.
        /// </summary>
        public static Task<GetDataTypeResult> InvokeAsync(GetDataTypeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDataTypeResult>("azure-native:networkanalytics:getDataType", args ?? new GetDataTypeArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieve data type resource.
        /// Azure REST API version: 2023-11-15.
        /// </summary>
        public static Output<GetDataTypeResult> Invoke(GetDataTypeInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDataTypeResult>("azure-native:networkanalytics:getDataType", args ?? new GetDataTypeInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieve data type resource.
        /// Azure REST API version: 2023-11-15.
        /// </summary>
        public static Output<GetDataTypeResult> Invoke(GetDataTypeInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDataTypeResult>("azure-native:networkanalytics:getDataType", args ?? new GetDataTypeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDataTypeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The data product resource name
        /// </summary>
        [Input("dataProductName", required: true)]
        public string DataProductName { get; set; } = null!;

        /// <summary>
        /// The data type name.
        /// </summary>
        [Input("dataTypeName", required: true)]
        public string DataTypeName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetDataTypeArgs()
        {
        }
        public static new GetDataTypeArgs Empty => new GetDataTypeArgs();
    }

    public sealed class GetDataTypeInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The data product resource name
        /// </summary>
        [Input("dataProductName", required: true)]
        public Input<string> DataProductName { get; set; } = null!;

        /// <summary>
        /// The data type name.
        /// </summary>
        [Input("dataTypeName", required: true)]
        public Input<string> DataTypeName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetDataTypeInvokeArgs()
        {
        }
        public static new GetDataTypeInvokeArgs Empty => new GetDataTypeInvokeArgs();
    }


    [OutputType]
    public sealed class GetDataTypeResult
    {
        /// <summary>
        /// Field for database cache retention in days.
        /// </summary>
        public readonly int? DatabaseCacheRetention;
        /// <summary>
        /// Field for database data retention in days.
        /// </summary>
        public readonly int? DatabaseRetention;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Latest provisioning state  of data product.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// State of data type.
        /// </summary>
        public readonly string? State;
        /// <summary>
        /// Reason for the state of data type.
        /// </summary>
        public readonly string StateReason;
        /// <summary>
        /// Field for storage output retention in days.
        /// </summary>
        public readonly int? StorageOutputRetention;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Url for data visualization.
        /// </summary>
        public readonly string VisualizationUrl;

        [OutputConstructor]
        private GetDataTypeResult(
            int? databaseCacheRetention,

            int? databaseRetention,

            string id,

            string name,

            string provisioningState,

            string? state,

            string stateReason,

            int? storageOutputRetention,

            Outputs.SystemDataResponse systemData,

            string type,

            string visualizationUrl)
        {
            DatabaseCacheRetention = databaseCacheRetention;
            DatabaseRetention = databaseRetention;
            Id = id;
            Name = name;
            ProvisioningState = provisioningState;
            State = state;
            StateReason = stateReason;
            StorageOutputRetention = storageOutputRetention;
            SystemData = systemData;
            Type = type;
            VisualizationUrl = visualizationUrl;
        }
    }
}
