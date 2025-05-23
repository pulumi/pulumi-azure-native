// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.IoTFirmwareDefense
{
    public static class GetFirmware
    {
        /// <summary>
        /// Get firmware.
        /// 
        /// Uses Azure REST API version 2024-01-10.
        /// 
        /// Other available API versions: 2023-02-08-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native iotfirmwaredefense [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetFirmwareResult> InvokeAsync(GetFirmwareArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFirmwareResult>("azure-native:iotfirmwaredefense:getFirmware", args ?? new GetFirmwareArgs(), options.WithDefaults());

        /// <summary>
        /// Get firmware.
        /// 
        /// Uses Azure REST API version 2024-01-10.
        /// 
        /// Other available API versions: 2023-02-08-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native iotfirmwaredefense [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetFirmwareResult> Invoke(GetFirmwareInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFirmwareResult>("azure-native:iotfirmwaredefense:getFirmware", args ?? new GetFirmwareInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get firmware.
        /// 
        /// Uses Azure REST API version 2024-01-10.
        /// 
        /// Other available API versions: 2023-02-08-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native iotfirmwaredefense [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetFirmwareResult> Invoke(GetFirmwareInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetFirmwareResult>("azure-native:iotfirmwaredefense:getFirmware", args ?? new GetFirmwareInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFirmwareArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the firmware.
        /// </summary>
        [Input("firmwareId", required: true)]
        public string FirmwareId { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the firmware analysis workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public string WorkspaceName { get; set; } = null!;

        public GetFirmwareArgs()
        {
        }
        public static new GetFirmwareArgs Empty => new GetFirmwareArgs();
    }

    public sealed class GetFirmwareInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the firmware.
        /// </summary>
        [Input("firmwareId", required: true)]
        public Input<string> FirmwareId { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the firmware analysis workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public GetFirmwareInvokeArgs()
        {
        }
        public static new GetFirmwareInvokeArgs Empty => new GetFirmwareInvokeArgs();
    }


    [OutputType]
    public sealed class GetFirmwareResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// User-specified description of the firmware.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// File name for a firmware that user uploaded.
        /// </summary>
        public readonly string? FileName;
        /// <summary>
        /// File size of the uploaded firmware image.
        /// </summary>
        public readonly double? FileSize;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Firmware model.
        /// </summary>
        public readonly string? Model;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Provisioning state of the resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The status of firmware scan.
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// A list of errors or other messages generated during firmware analysis
        /// </summary>
        public readonly ImmutableArray<Outputs.StatusMessageResponse> StatusMessages;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Firmware vendor.
        /// </summary>
        public readonly string? Vendor;
        /// <summary>
        /// Firmware version.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private GetFirmwareResult(
            string azureApiVersion,

            string? description,

            string? fileName,

            double? fileSize,

            string id,

            string? model,

            string name,

            string provisioningState,

            string? status,

            ImmutableArray<Outputs.StatusMessageResponse> statusMessages,

            Outputs.SystemDataResponse systemData,

            string type,

            string? vendor,

            string? version)
        {
            AzureApiVersion = azureApiVersion;
            Description = description;
            FileName = fileName;
            FileSize = fileSize;
            Id = id;
            Model = model;
            Name = name;
            ProvisioningState = provisioningState;
            Status = status;
            StatusMessages = statusMessages;
            SystemData = systemData;
            Type = type;
            Vendor = vendor;
            Version = version;
        }
    }
}
