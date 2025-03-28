// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Automation
{
    public static class GetConnection
    {
        /// <summary>
        /// Retrieve the connection identified by connection name.
        /// 
        /// Uses Azure REST API version 2022-08-08.
        /// 
        /// Other available API versions: 2023-05-15-preview, 2023-11-01, 2024-10-23.
        /// </summary>
        public static Task<GetConnectionResult> InvokeAsync(GetConnectionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetConnectionResult>("azure-native:automation:getConnection", args ?? new GetConnectionArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieve the connection identified by connection name.
        /// 
        /// Uses Azure REST API version 2022-08-08.
        /// 
        /// Other available API versions: 2023-05-15-preview, 2023-11-01, 2024-10-23.
        /// </summary>
        public static Output<GetConnectionResult> Invoke(GetConnectionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetConnectionResult>("azure-native:automation:getConnection", args ?? new GetConnectionInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieve the connection identified by connection name.
        /// 
        /// Uses Azure REST API version 2022-08-08.
        /// 
        /// Other available API versions: 2023-05-15-preview, 2023-11-01, 2024-10-23.
        /// </summary>
        public static Output<GetConnectionResult> Invoke(GetConnectionInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetConnectionResult>("azure-native:automation:getConnection", args ?? new GetConnectionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetConnectionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the automation account.
        /// </summary>
        [Input("automationAccountName", required: true)]
        public string AutomationAccountName { get; set; } = null!;

        /// <summary>
        /// The name of connection.
        /// </summary>
        [Input("connectionName", required: true)]
        public string ConnectionName { get; set; } = null!;

        /// <summary>
        /// Name of an Azure Resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetConnectionArgs()
        {
        }
        public static new GetConnectionArgs Empty => new GetConnectionArgs();
    }

    public sealed class GetConnectionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the automation account.
        /// </summary>
        [Input("automationAccountName", required: true)]
        public Input<string> AutomationAccountName { get; set; } = null!;

        /// <summary>
        /// The name of connection.
        /// </summary>
        [Input("connectionName", required: true)]
        public Input<string> ConnectionName { get; set; } = null!;

        /// <summary>
        /// Name of an Azure Resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetConnectionInvokeArgs()
        {
        }
        public static new GetConnectionInvokeArgs Empty => new GetConnectionInvokeArgs();
    }


    [OutputType]
    public sealed class GetConnectionResult
    {
        /// <summary>
        /// Gets or sets the connectionType of the connection.
        /// </summary>
        public readonly Outputs.ConnectionTypeAssociationPropertyResponse? ConnectionType;
        /// <summary>
        /// Gets the creation time.
        /// </summary>
        public readonly string CreationTime;
        /// <summary>
        /// Gets or sets the description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Gets the field definition values of the connection.
        /// </summary>
        public readonly ImmutableDictionary<string, string> FieldDefinitionValues;
        /// <summary>
        /// Fully qualified resource Id for the resource
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Gets the last modified time.
        /// </summary>
        public readonly string LastModifiedTime;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The type of the resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetConnectionResult(
            Outputs.ConnectionTypeAssociationPropertyResponse? connectionType,

            string creationTime,

            string? description,

            ImmutableDictionary<string, string> fieldDefinitionValues,

            string id,

            string lastModifiedTime,

            string name,

            string type)
        {
            ConnectionType = connectionType;
            CreationTime = creationTime;
            Description = description;
            FieldDefinitionValues = fieldDefinitionValues;
            Id = id;
            LastModifiedTime = lastModifiedTime;
            Name = name;
            Type = type;
        }
    }
}
