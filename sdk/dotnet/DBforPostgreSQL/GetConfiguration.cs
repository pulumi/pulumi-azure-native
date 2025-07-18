// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DBforPostgreSQL
{
    public static class GetConfiguration
    {
        /// <summary>
        /// Gets information about a specific server parameter of a flexible server.
        /// 
        /// Uses Azure REST API version 2024-08-01.
        /// 
        /// Other available API versions: 2022-12-01, 2023-03-01-preview, 2023-06-01-preview, 2023-12-01-preview, 2024-03-01-preview, 2024-11-01-preview, 2025-01-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native dbforpostgresql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetConfigurationResult> InvokeAsync(GetConfigurationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetConfigurationResult>("azure-native:dbforpostgresql:getConfiguration", args ?? new GetConfigurationArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a specific server parameter of a flexible server.
        /// 
        /// Uses Azure REST API version 2024-08-01.
        /// 
        /// Other available API versions: 2022-12-01, 2023-03-01-preview, 2023-06-01-preview, 2023-12-01-preview, 2024-03-01-preview, 2024-11-01-preview, 2025-01-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native dbforpostgresql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetConfigurationResult> Invoke(GetConfigurationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetConfigurationResult>("azure-native:dbforpostgresql:getConfiguration", args ?? new GetConfigurationInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a specific server parameter of a flexible server.
        /// 
        /// Uses Azure REST API version 2024-08-01.
        /// 
        /// Other available API versions: 2022-12-01, 2023-03-01-preview, 2023-06-01-preview, 2023-12-01-preview, 2024-03-01-preview, 2024-11-01-preview, 2025-01-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native dbforpostgresql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetConfigurationResult> Invoke(GetConfigurationInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetConfigurationResult>("azure-native:dbforpostgresql:getConfiguration", args ?? new GetConfigurationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetConfigurationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the server parameter.
        /// </summary>
        [Input("configurationName", required: true)]
        public string ConfigurationName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the server.
        /// </summary>
        [Input("serverName", required: true)]
        public string ServerName { get; set; } = null!;

        public GetConfigurationArgs()
        {
        }
        public static new GetConfigurationArgs Empty => new GetConfigurationArgs();
    }

    public sealed class GetConfigurationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the server parameter.
        /// </summary>
        [Input("configurationName", required: true)]
        public Input<string> ConfigurationName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the server.
        /// </summary>
        [Input("serverName", required: true)]
        public Input<string> ServerName { get; set; } = null!;

        public GetConfigurationInvokeArgs()
        {
        }
        public static new GetConfigurationInvokeArgs Empty => new GetConfigurationInvokeArgs();
    }


    [OutputType]
    public sealed class GetConfigurationResult
    {
        /// <summary>
        /// Allowed values of the server parameter.
        /// </summary>
        public readonly string AllowedValues;
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Data type of the server parameter.
        /// </summary>
        public readonly string DataType;
        /// <summary>
        /// Value assigned by default to the server parameter.
        /// </summary>
        public readonly string DefaultValue;
        /// <summary>
        /// Description of the server parameter.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Link pointing to the documentation of the server parameter.
        /// </summary>
        public readonly string DocumentationLink;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Indicates if the value assigned to the server parameter is pending a server restart for the value to take effect.
        /// </summary>
        public readonly bool IsConfigPendingRestart;
        /// <summary>
        /// Indicates if it's a dynamic (true) or static (false) server parameter. Static server parameters require a server restart after changing the value assigned to it, for the change to take effect. Dynamic server parameters do not require a server restart after changing the value assigned to it, for the change to take effect.
        /// </summary>
        public readonly bool IsDynamicConfig;
        /// <summary>
        /// Indicates if it's a read-only (true) or modifiable (false) server parameter.
        /// </summary>
        public readonly bool IsReadOnly;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Source of the value assigned to the server parameter. Required to update the value assigned to a specific modifiable server parameter.
        /// </summary>
        public readonly string? Source;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Units in which the server parameter value is expressed.
        /// </summary>
        public readonly string Unit;
        /// <summary>
        /// Value of the server parameter (also known as configuration). Required to update the value assigned to a specific modifiable server parameter.
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private GetConfigurationResult(
            string allowedValues,

            string azureApiVersion,

            string dataType,

            string defaultValue,

            string description,

            string documentationLink,

            string id,

            bool isConfigPendingRestart,

            bool isDynamicConfig,

            bool isReadOnly,

            string name,

            string? source,

            Outputs.SystemDataResponse systemData,

            string type,

            string unit,

            string? value)
        {
            AllowedValues = allowedValues;
            AzureApiVersion = azureApiVersion;
            DataType = dataType;
            DefaultValue = defaultValue;
            Description = description;
            DocumentationLink = documentationLink;
            Id = id;
            IsConfigPendingRestart = isConfigPendingRestart;
            IsDynamicConfig = isDynamicConfig;
            IsReadOnly = isReadOnly;
            Name = name;
            Source = source;
            SystemData = systemData;
            Type = type;
            Unit = unit;
            Value = value;
        }
    }
}
