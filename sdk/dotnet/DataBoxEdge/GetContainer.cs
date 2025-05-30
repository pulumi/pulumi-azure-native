// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataBoxEdge
{
    public static class GetContainer
    {
        /// <summary>
        /// Represents a container on the  Data Box Edge/Gateway device.
        /// 
        /// Uses Azure REST API version 2023-07-01.
        /// 
        /// Other available API versions: 2022-03-01, 2022-04-01-preview, 2022-12-01-preview, 2023-01-01-preview, 2023-12-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native databoxedge [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetContainerResult> InvokeAsync(GetContainerArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetContainerResult>("azure-native:databoxedge:getContainer", args ?? new GetContainerArgs(), options.WithDefaults());

        /// <summary>
        /// Represents a container on the  Data Box Edge/Gateway device.
        /// 
        /// Uses Azure REST API version 2023-07-01.
        /// 
        /// Other available API versions: 2022-03-01, 2022-04-01-preview, 2022-12-01-preview, 2023-01-01-preview, 2023-12-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native databoxedge [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetContainerResult> Invoke(GetContainerInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetContainerResult>("azure-native:databoxedge:getContainer", args ?? new GetContainerInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Represents a container on the  Data Box Edge/Gateway device.
        /// 
        /// Uses Azure REST API version 2023-07-01.
        /// 
        /// Other available API versions: 2022-03-01, 2022-04-01-preview, 2022-12-01-preview, 2023-01-01-preview, 2023-12-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native databoxedge [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetContainerResult> Invoke(GetContainerInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetContainerResult>("azure-native:databoxedge:getContainer", args ?? new GetContainerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetContainerArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The container Name
        /// </summary>
        [Input("containerName", required: true)]
        public string ContainerName { get; set; } = null!;

        /// <summary>
        /// The device name.
        /// </summary>
        [Input("deviceName", required: true)]
        public string DeviceName { get; set; } = null!;

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The Storage Account Name
        /// </summary>
        [Input("storageAccountName", required: true)]
        public string StorageAccountName { get; set; } = null!;

        public GetContainerArgs()
        {
        }
        public static new GetContainerArgs Empty => new GetContainerArgs();
    }

    public sealed class GetContainerInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The container Name
        /// </summary>
        [Input("containerName", required: true)]
        public Input<string> ContainerName { get; set; } = null!;

        /// <summary>
        /// The device name.
        /// </summary>
        [Input("deviceName", required: true)]
        public Input<string> DeviceName { get; set; } = null!;

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The Storage Account Name
        /// </summary>
        [Input("storageAccountName", required: true)]
        public Input<string> StorageAccountName { get; set; } = null!;

        public GetContainerInvokeArgs()
        {
        }
        public static new GetContainerInvokeArgs Empty => new GetContainerInvokeArgs();
    }


    [OutputType]
    public sealed class GetContainerResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Current status of the container.
        /// </summary>
        public readonly string ContainerStatus;
        /// <summary>
        /// The UTC time when container got created.
        /// </summary>
        public readonly string CreatedDateTime;
        /// <summary>
        /// DataFormat for Container
        /// </summary>
        public readonly string DataFormat;
        /// <summary>
        /// The path ID that uniquely identifies the object.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The object name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Details of the refresh job on this container.
        /// </summary>
        public readonly Outputs.RefreshDetailsResponse RefreshDetails;
        /// <summary>
        /// Metadata pertaining to creation and last modification of Container
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The hierarchical type of the object.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetContainerResult(
            string azureApiVersion,

            string containerStatus,

            string createdDateTime,

            string dataFormat,

            string id,

            string name,

            Outputs.RefreshDetailsResponse refreshDetails,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            ContainerStatus = containerStatus;
            CreatedDateTime = createdDateTime;
            DataFormat = dataFormat;
            Id = id;
            Name = name;
            RefreshDetails = refreshDetails;
            SystemData = systemData;
            Type = type;
        }
    }
}
