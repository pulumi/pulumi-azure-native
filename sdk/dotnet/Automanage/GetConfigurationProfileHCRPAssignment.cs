// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Automanage
{
    public static class GetConfigurationProfileHCRPAssignment
    {
        /// <summary>
        /// Get information about a configuration profile assignment
        /// 
        /// Uses Azure REST API version 2022-05-04.
        /// 
        /// Other available API versions: 2021-04-30-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native automanage [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetConfigurationProfileHCRPAssignmentResult> InvokeAsync(GetConfigurationProfileHCRPAssignmentArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetConfigurationProfileHCRPAssignmentResult>("azure-native:automanage:getConfigurationProfileHCRPAssignment", args ?? new GetConfigurationProfileHCRPAssignmentArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about a configuration profile assignment
        /// 
        /// Uses Azure REST API version 2022-05-04.
        /// 
        /// Other available API versions: 2021-04-30-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native automanage [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetConfigurationProfileHCRPAssignmentResult> Invoke(GetConfigurationProfileHCRPAssignmentInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetConfigurationProfileHCRPAssignmentResult>("azure-native:automanage:getConfigurationProfileHCRPAssignment", args ?? new GetConfigurationProfileHCRPAssignmentInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about a configuration profile assignment
        /// 
        /// Uses Azure REST API version 2022-05-04.
        /// 
        /// Other available API versions: 2021-04-30-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native automanage [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetConfigurationProfileHCRPAssignmentResult> Invoke(GetConfigurationProfileHCRPAssignmentInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetConfigurationProfileHCRPAssignmentResult>("azure-native:automanage:getConfigurationProfileHCRPAssignment", args ?? new GetConfigurationProfileHCRPAssignmentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetConfigurationProfileHCRPAssignmentArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The configuration profile assignment name.
        /// </summary>
        [Input("configurationProfileAssignmentName", required: true)]
        public string ConfigurationProfileAssignmentName { get; set; } = null!;

        /// <summary>
        /// The name of the Arc machine.
        /// </summary>
        [Input("machineName", required: true)]
        public string MachineName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetConfigurationProfileHCRPAssignmentArgs()
        {
        }
        public static new GetConfigurationProfileHCRPAssignmentArgs Empty => new GetConfigurationProfileHCRPAssignmentArgs();
    }

    public sealed class GetConfigurationProfileHCRPAssignmentInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The configuration profile assignment name.
        /// </summary>
        [Input("configurationProfileAssignmentName", required: true)]
        public Input<string> ConfigurationProfileAssignmentName { get; set; } = null!;

        /// <summary>
        /// The name of the Arc machine.
        /// </summary>
        [Input("machineName", required: true)]
        public Input<string> MachineName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetConfigurationProfileHCRPAssignmentInvokeArgs()
        {
        }
        public static new GetConfigurationProfileHCRPAssignmentInvokeArgs Empty => new GetConfigurationProfileHCRPAssignmentInvokeArgs();
    }


    [OutputType]
    public sealed class GetConfigurationProfileHCRPAssignmentResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Azure resource id. Indicates if this resource is managed by another Azure resource.
        /// </summary>
        public readonly string ManagedBy;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Properties of the configuration profile assignment.
        /// </summary>
        public readonly Outputs.ConfigurationProfileAssignmentPropertiesResponse Properties;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetConfigurationProfileHCRPAssignmentResult(
            string azureApiVersion,

            string id,

            string managedBy,

            string name,

            Outputs.ConfigurationProfileAssignmentPropertiesResponse properties,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Id = id;
            ManagedBy = managedBy;
            Name = name;
            Properties = properties;
            SystemData = systemData;
            Type = type;
        }
    }
}
