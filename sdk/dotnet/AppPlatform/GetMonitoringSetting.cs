// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AppPlatform
{
    public static class GetMonitoringSetting
    {
        /// <summary>
        /// Get the Monitoring Setting and its properties.
        /// 
        /// Uses Azure REST API version 2023-05-01-preview.
        /// 
        /// Other available API versions: 2023-07-01-preview, 2023-09-01-preview, 2023-11-01-preview, 2023-12-01, 2024-01-01-preview, 2024-05-01-preview.
        /// </summary>
        public static Task<GetMonitoringSettingResult> InvokeAsync(GetMonitoringSettingArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMonitoringSettingResult>("azure-native:appplatform:getMonitoringSetting", args ?? new GetMonitoringSettingArgs(), options.WithDefaults());

        /// <summary>
        /// Get the Monitoring Setting and its properties.
        /// 
        /// Uses Azure REST API version 2023-05-01-preview.
        /// 
        /// Other available API versions: 2023-07-01-preview, 2023-09-01-preview, 2023-11-01-preview, 2023-12-01, 2024-01-01-preview, 2024-05-01-preview.
        /// </summary>
        public static Output<GetMonitoringSettingResult> Invoke(GetMonitoringSettingInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMonitoringSettingResult>("azure-native:appplatform:getMonitoringSetting", args ?? new GetMonitoringSettingInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get the Monitoring Setting and its properties.
        /// 
        /// Uses Azure REST API version 2023-05-01-preview.
        /// 
        /// Other available API versions: 2023-07-01-preview, 2023-09-01-preview, 2023-11-01-preview, 2023-12-01, 2024-01-01-preview, 2024-05-01-preview.
        /// </summary>
        public static Output<GetMonitoringSettingResult> Invoke(GetMonitoringSettingInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetMonitoringSettingResult>("azure-native:appplatform:getMonitoringSetting", args ?? new GetMonitoringSettingInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMonitoringSettingArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the Service resource.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetMonitoringSettingArgs()
        {
        }
        public static new GetMonitoringSettingArgs Empty => new GetMonitoringSettingArgs();
    }

    public sealed class GetMonitoringSettingInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the Service resource.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetMonitoringSettingInvokeArgs()
        {
        }
        public static new GetMonitoringSettingInvokeArgs Empty => new GetMonitoringSettingInvokeArgs();
    }


    [OutputType]
    public sealed class GetMonitoringSettingResult
    {
        /// <summary>
        /// Fully qualified resource Id for the resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Properties of the Monitoring Setting resource
        /// </summary>
        public readonly Outputs.MonitoringSettingPropertiesResponse Properties;
        /// <summary>
        /// Metadata pertaining to creation and last modification of the resource.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetMonitoringSettingResult(
            string id,

            string name,

            Outputs.MonitoringSettingPropertiesResponse properties,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            Id = id;
            Name = name;
            Properties = properties;
            SystemData = systemData;
            Type = type;
        }
    }
}
