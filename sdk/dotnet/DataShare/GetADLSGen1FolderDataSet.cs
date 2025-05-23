// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataShare
{
    public static class GetADLSGen1FolderDataSet
    {
        /// <summary>
        /// Get a DataSet in a share
        /// 
        /// Uses Azure REST API version 2021-08-01.
        /// </summary>
        public static Task<GetADLSGen1FolderDataSetResult> InvokeAsync(GetADLSGen1FolderDataSetArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetADLSGen1FolderDataSetResult>("azure-native:datashare:getADLSGen1FolderDataSet", args ?? new GetADLSGen1FolderDataSetArgs(), options.WithDefaults());

        /// <summary>
        /// Get a DataSet in a share
        /// 
        /// Uses Azure REST API version 2021-08-01.
        /// </summary>
        public static Output<GetADLSGen1FolderDataSetResult> Invoke(GetADLSGen1FolderDataSetInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetADLSGen1FolderDataSetResult>("azure-native:datashare:getADLSGen1FolderDataSet", args ?? new GetADLSGen1FolderDataSetInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get a DataSet in a share
        /// 
        /// Uses Azure REST API version 2021-08-01.
        /// </summary>
        public static Output<GetADLSGen1FolderDataSetResult> Invoke(GetADLSGen1FolderDataSetInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetADLSGen1FolderDataSetResult>("azure-native:datashare:getADLSGen1FolderDataSet", args ?? new GetADLSGen1FolderDataSetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetADLSGen1FolderDataSetArgs : global::Pulumi.InvokeArgs
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

        public GetADLSGen1FolderDataSetArgs()
        {
        }
        public static new GetADLSGen1FolderDataSetArgs Empty => new GetADLSGen1FolderDataSetArgs();
    }

    public sealed class GetADLSGen1FolderDataSetInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetADLSGen1FolderDataSetInvokeArgs()
        {
        }
        public static new GetADLSGen1FolderDataSetInvokeArgs Empty => new GetADLSGen1FolderDataSetInvokeArgs();
    }


    [OutputType]
    public sealed class GetADLSGen1FolderDataSetResult
    {
        /// <summary>
        /// The ADLS account name.
        /// </summary>
        public readonly string AccountName;
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Unique id for identifying a data set resource
        /// </summary>
        public readonly string DataSetId;
        /// <summary>
        /// The folder path within the ADLS account.
        /// </summary>
        public readonly string FolderPath;
        /// <summary>
        /// The resource id of the azure resource
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Kind of data set.
        /// Expected value is 'AdlsGen1Folder'.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Name of the azure resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Resource group of ADLS account.
        /// </summary>
        public readonly string ResourceGroup;
        /// <summary>
        /// Subscription id of ADLS account.
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
        private GetADLSGen1FolderDataSetResult(
            string accountName,

            string azureApiVersion,

            string dataSetId,

            string folderPath,

            string id,

            string kind,

            string name,

            string resourceGroup,

            string subscriptionId,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            AccountName = accountName;
            AzureApiVersion = azureApiVersion;
            DataSetId = dataSetId;
            FolderPath = folderPath;
            Id = id;
            Kind = kind;
            Name = name;
            ResourceGroup = resourceGroup;
            SubscriptionId = subscriptionId;
            SystemData = systemData;
            Type = type;
        }
    }
}
