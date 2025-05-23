// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataShare
{
    public static class GetADLSGen2FileSystemDataSet
    {
        /// <summary>
        /// Get a DataSet in a share
        /// 
        /// Uses Azure REST API version 2021-08-01.
        /// </summary>
        public static Task<GetADLSGen2FileSystemDataSetResult> InvokeAsync(GetADLSGen2FileSystemDataSetArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetADLSGen2FileSystemDataSetResult>("azure-native:datashare:getADLSGen2FileSystemDataSet", args ?? new GetADLSGen2FileSystemDataSetArgs(), options.WithDefaults());

        /// <summary>
        /// Get a DataSet in a share
        /// 
        /// Uses Azure REST API version 2021-08-01.
        /// </summary>
        public static Output<GetADLSGen2FileSystemDataSetResult> Invoke(GetADLSGen2FileSystemDataSetInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetADLSGen2FileSystemDataSetResult>("azure-native:datashare:getADLSGen2FileSystemDataSet", args ?? new GetADLSGen2FileSystemDataSetInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get a DataSet in a share
        /// 
        /// Uses Azure REST API version 2021-08-01.
        /// </summary>
        public static Output<GetADLSGen2FileSystemDataSetResult> Invoke(GetADLSGen2FileSystemDataSetInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetADLSGen2FileSystemDataSetResult>("azure-native:datashare:getADLSGen2FileSystemDataSet", args ?? new GetADLSGen2FileSystemDataSetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetADLSGen2FileSystemDataSetArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the share account.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the dataSet.
        /// </summary>
        [Input("dataSetName", required: true)]
        public string DataSetName { get; set; } = null!;

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the share.
        /// </summary>
        [Input("shareName", required: true)]
        public string ShareName { get; set; } = null!;

        public GetADLSGen2FileSystemDataSetArgs()
        {
        }
        public static new GetADLSGen2FileSystemDataSetArgs Empty => new GetADLSGen2FileSystemDataSetArgs();
    }

    public sealed class GetADLSGen2FileSystemDataSetInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the share account.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the dataSet.
        /// </summary>
        [Input("dataSetName", required: true)]
        public Input<string> DataSetName { get; set; } = null!;

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the share.
        /// </summary>
        [Input("shareName", required: true)]
        public Input<string> ShareName { get; set; } = null!;

        public GetADLSGen2FileSystemDataSetInvokeArgs()
        {
        }
        public static new GetADLSGen2FileSystemDataSetInvokeArgs Empty => new GetADLSGen2FileSystemDataSetInvokeArgs();
    }


    [OutputType]
    public sealed class GetADLSGen2FileSystemDataSetResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Unique id for identifying a data set resource
        /// </summary>
        public readonly string DataSetId;
        /// <summary>
        /// The file system name.
        /// </summary>
        public readonly string FileSystem;
        /// <summary>
        /// The resource id of the azure resource
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Kind of data set.
        /// Expected value is 'AdlsGen2FileSystem'.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Name of the azure resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Resource group of storage account
        /// </summary>
        public readonly string ResourceGroup;
        /// <summary>
        /// Storage account name of the source data set
        /// </summary>
        public readonly string StorageAccountName;
        /// <summary>
        /// Subscription id of storage account
        /// </summary>
        public readonly string SubscriptionId;
        /// <summary>
        /// System Data of the Azure resource.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Type of the azure resource
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetADLSGen2FileSystemDataSetResult(
            string azureApiVersion,

            string dataSetId,

            string fileSystem,

            string id,

            string kind,

            string name,

            string resourceGroup,

            string storageAccountName,

            string subscriptionId,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            DataSetId = dataSetId;
            FileSystem = fileSystem;
            Id = id;
            Kind = kind;
            Name = name;
            ResourceGroup = resourceGroup;
            StorageAccountName = storageAccountName;
            SubscriptionId = subscriptionId;
            SystemData = systemData;
            Type = type;
        }
    }
}
