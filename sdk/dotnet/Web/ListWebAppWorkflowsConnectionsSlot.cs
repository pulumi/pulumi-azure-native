// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web
{
    public static class ListWebAppWorkflowsConnectionsSlot
    {
        /// <summary>
        /// Workflow properties definition.
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// 
        /// Other available API versions: 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<ListWebAppWorkflowsConnectionsSlotResult> InvokeAsync(ListWebAppWorkflowsConnectionsSlotArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListWebAppWorkflowsConnectionsSlotResult>("azure-native:web:listWebAppWorkflowsConnectionsSlot", args ?? new ListWebAppWorkflowsConnectionsSlotArgs(), options.WithDefaults());

        /// <summary>
        /// Workflow properties definition.
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// 
        /// Other available API versions: 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListWebAppWorkflowsConnectionsSlotResult> Invoke(ListWebAppWorkflowsConnectionsSlotInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListWebAppWorkflowsConnectionsSlotResult>("azure-native:web:listWebAppWorkflowsConnectionsSlot", args ?? new ListWebAppWorkflowsConnectionsSlotInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Workflow properties definition.
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// 
        /// Other available API versions: 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListWebAppWorkflowsConnectionsSlotResult> Invoke(ListWebAppWorkflowsConnectionsSlotInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListWebAppWorkflowsConnectionsSlotResult>("azure-native:web:listWebAppWorkflowsConnectionsSlot", args ?? new ListWebAppWorkflowsConnectionsSlotInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListWebAppWorkflowsConnectionsSlotArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Site name.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of the deployment slot.
        /// </summary>
        [Input("slot", required: true)]
        public string Slot { get; set; } = null!;

        public ListWebAppWorkflowsConnectionsSlotArgs()
        {
        }
        public static new ListWebAppWorkflowsConnectionsSlotArgs Empty => new ListWebAppWorkflowsConnectionsSlotArgs();
    }

    public sealed class ListWebAppWorkflowsConnectionsSlotInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Site name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of the deployment slot.
        /// </summary>
        [Input("slot", required: true)]
        public Input<string> Slot { get; set; } = null!;

        public ListWebAppWorkflowsConnectionsSlotInvokeArgs()
        {
        }
        public static new ListWebAppWorkflowsConnectionsSlotInvokeArgs Empty => new ListWebAppWorkflowsConnectionsSlotInvokeArgs();
    }


    [OutputType]
    public sealed class ListWebAppWorkflowsConnectionsSlotResult
    {
        /// <summary>
        /// The resource id.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The resource kind.
        /// </summary>
        public readonly string? Kind;
        /// <summary>
        /// The resource location.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Gets the resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Additional workflow properties.
        /// </summary>
        public readonly Outputs.WorkflowEnvelopeResponseProperties Properties;
        /// <summary>
        /// Gets the resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ListWebAppWorkflowsConnectionsSlotResult(
            string id,

            string? kind,

            string? location,

            string name,

            Outputs.WorkflowEnvelopeResponseProperties properties,

            string type)
        {
            Id = id;
            Kind = kind;
            Location = location;
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}
