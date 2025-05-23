// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MobileNetwork
{
    public static class GetDiagnosticsPackage
    {
        /// <summary>
        /// Gets information about the specified diagnostics package.
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// 
        /// Other available API versions: 2023-06-01, 2023-09-01, 2024-02-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native mobilenetwork [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetDiagnosticsPackageResult> InvokeAsync(GetDiagnosticsPackageArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDiagnosticsPackageResult>("azure-native:mobilenetwork:getDiagnosticsPackage", args ?? new GetDiagnosticsPackageArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about the specified diagnostics package.
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// 
        /// Other available API versions: 2023-06-01, 2023-09-01, 2024-02-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native mobilenetwork [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetDiagnosticsPackageResult> Invoke(GetDiagnosticsPackageInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDiagnosticsPackageResult>("azure-native:mobilenetwork:getDiagnosticsPackage", args ?? new GetDiagnosticsPackageInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about the specified diagnostics package.
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// 
        /// Other available API versions: 2023-06-01, 2023-09-01, 2024-02-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native mobilenetwork [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetDiagnosticsPackageResult> Invoke(GetDiagnosticsPackageInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDiagnosticsPackageResult>("azure-native:mobilenetwork:getDiagnosticsPackage", args ?? new GetDiagnosticsPackageInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDiagnosticsPackageArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the diagnostics package.
        /// </summary>
        [Input("diagnosticsPackageName", required: true)]
        public string DiagnosticsPackageName { get; set; } = null!;

        /// <summary>
        /// The name of the packet core control plane.
        /// </summary>
        [Input("packetCoreControlPlaneName", required: true)]
        public string PacketCoreControlPlaneName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetDiagnosticsPackageArgs()
        {
        }
        public static new GetDiagnosticsPackageArgs Empty => new GetDiagnosticsPackageArgs();
    }

    public sealed class GetDiagnosticsPackageInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the diagnostics package.
        /// </summary>
        [Input("diagnosticsPackageName", required: true)]
        public Input<string> DiagnosticsPackageName { get; set; } = null!;

        /// <summary>
        /// The name of the packet core control plane.
        /// </summary>
        [Input("packetCoreControlPlaneName", required: true)]
        public Input<string> PacketCoreControlPlaneName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetDiagnosticsPackageInvokeArgs()
        {
        }
        public static new GetDiagnosticsPackageInvokeArgs Empty => new GetDiagnosticsPackageInvokeArgs();
    }


    [OutputType]
    public sealed class GetDiagnosticsPackageResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The provisioning state of the diagnostics package resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The reason for the current state of the diagnostics package collection.
        /// </summary>
        public readonly string Reason;
        /// <summary>
        /// The status of the diagnostics package collection.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetDiagnosticsPackageResult(
            string azureApiVersion,

            string id,

            string name,

            string provisioningState,

            string reason,

            string status,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Id = id;
            Name = name;
            ProvisioningState = provisioningState;
            Reason = reason;
            Status = status;
            SystemData = systemData;
            Type = type;
        }
    }
}
