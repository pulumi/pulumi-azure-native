// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Maintenance.V20230901Preview
{
    public static class GetConfigurationAssignment
    {
        /// <summary>
        /// Get configuration assignment for resource..
        /// </summary>
        public static Task<GetConfigurationAssignmentResult> InvokeAsync(GetConfigurationAssignmentArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetConfigurationAssignmentResult>("azure-native:maintenance/v20230901preview:getConfigurationAssignment", args ?? new GetConfigurationAssignmentArgs(), options.WithDefaults());

        /// <summary>
        /// Get configuration assignment for resource..
        /// </summary>
        public static Output<GetConfigurationAssignmentResult> Invoke(GetConfigurationAssignmentInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetConfigurationAssignmentResult>("azure-native:maintenance/v20230901preview:getConfigurationAssignment", args ?? new GetConfigurationAssignmentInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get configuration assignment for resource..
        /// </summary>
        public static Output<GetConfigurationAssignmentResult> Invoke(GetConfigurationAssignmentInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetConfigurationAssignmentResult>("azure-native:maintenance/v20230901preview:getConfigurationAssignment", args ?? new GetConfigurationAssignmentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetConfigurationAssignmentArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Configuration assignment name
        /// </summary>
        [Input("configurationAssignmentName", required: true)]
        public string ConfigurationAssignmentName { get; set; } = null!;

        /// <summary>
        /// Resource provider name
        /// </summary>
        [Input("providerName", required: true)]
        public string ProviderName { get; set; } = null!;

        /// <summary>
        /// Resource group name
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Resource identifier
        /// </summary>
        [Input("resourceName", required: true)]
        public string ResourceName { get; set; } = null!;

        /// <summary>
        /// Resource type
        /// </summary>
        [Input("resourceType", required: true)]
        public string ResourceType { get; set; } = null!;

        public GetConfigurationAssignmentArgs()
        {
        }
        public static new GetConfigurationAssignmentArgs Empty => new GetConfigurationAssignmentArgs();
    }

    public sealed class GetConfigurationAssignmentInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Configuration assignment name
        /// </summary>
        [Input("configurationAssignmentName", required: true)]
        public Input<string> ConfigurationAssignmentName { get; set; } = null!;

        /// <summary>
        /// Resource provider name
        /// </summary>
        [Input("providerName", required: true)]
        public Input<string> ProviderName { get; set; } = null!;

        /// <summary>
        /// Resource group name
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Resource identifier
        /// </summary>
        [Input("resourceName", required: true)]
        public Input<string> ResourceName { get; set; } = null!;

        /// <summary>
        /// Resource type
        /// </summary>
        [Input("resourceType", required: true)]
        public Input<string> ResourceType { get; set; } = null!;

        public GetConfigurationAssignmentInvokeArgs()
        {
        }
        public static new GetConfigurationAssignmentInvokeArgs Empty => new GetConfigurationAssignmentInvokeArgs();
    }


    [OutputType]
    public sealed class GetConfigurationAssignmentResult
    {
        /// <summary>
        /// Properties of the configuration assignment
        /// </summary>
        public readonly Outputs.ConfigurationAssignmentFilterPropertiesResponse? Filter;
        /// <summary>
        /// Fully qualified identifier of the resource
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Location of the resource
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// The maintenance configuration Id
        /// </summary>
        public readonly string? MaintenanceConfigurationId;
        /// <summary>
        /// Name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The unique resourceId
        /// </summary>
        public readonly string? ResourceId;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Type of the resource
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetConfigurationAssignmentResult(
            Outputs.ConfigurationAssignmentFilterPropertiesResponse? filter,

            string id,

            string? location,

            string? maintenanceConfigurationId,

            string name,

            string? resourceId,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            Filter = filter;
            Id = id;
            Location = location;
            MaintenanceConfigurationId = maintenanceConfigurationId;
            Name = name;
            ResourceId = resourceId;
            SystemData = systemData;
            Type = type;
        }
    }
}
