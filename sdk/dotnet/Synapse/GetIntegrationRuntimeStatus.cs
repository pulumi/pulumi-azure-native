// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Synapse
{
    public static class GetIntegrationRuntimeStatus
    {
        /// <summary>
        /// Get the integration runtime status
        /// 
        /// Uses Azure REST API version 2021-06-01.
        /// 
        /// Other available API versions: 2021-04-01-preview, 2021-05-01, 2021-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native synapse [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetIntegrationRuntimeStatusResult> InvokeAsync(GetIntegrationRuntimeStatusArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIntegrationRuntimeStatusResult>("azure-native:synapse:getIntegrationRuntimeStatus", args ?? new GetIntegrationRuntimeStatusArgs(), options.WithDefaults());

        /// <summary>
        /// Get the integration runtime status
        /// 
        /// Uses Azure REST API version 2021-06-01.
        /// 
        /// Other available API versions: 2021-04-01-preview, 2021-05-01, 2021-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native synapse [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetIntegrationRuntimeStatusResult> Invoke(GetIntegrationRuntimeStatusInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIntegrationRuntimeStatusResult>("azure-native:synapse:getIntegrationRuntimeStatus", args ?? new GetIntegrationRuntimeStatusInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get the integration runtime status
        /// 
        /// Uses Azure REST API version 2021-06-01.
        /// 
        /// Other available API versions: 2021-04-01-preview, 2021-05-01, 2021-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native synapse [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetIntegrationRuntimeStatusResult> Invoke(GetIntegrationRuntimeStatusInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetIntegrationRuntimeStatusResult>("azure-native:synapse:getIntegrationRuntimeStatus", args ?? new GetIntegrationRuntimeStatusInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIntegrationRuntimeStatusArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Integration runtime name
        /// </summary>
        [Input("integrationRuntimeName", required: true)]
        public string IntegrationRuntimeName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public string WorkspaceName { get; set; } = null!;

        public GetIntegrationRuntimeStatusArgs()
        {
        }
        public static new GetIntegrationRuntimeStatusArgs Empty => new GetIntegrationRuntimeStatusArgs();
    }

    public sealed class GetIntegrationRuntimeStatusInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Integration runtime name
        /// </summary>
        [Input("integrationRuntimeName", required: true)]
        public Input<string> IntegrationRuntimeName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public GetIntegrationRuntimeStatusInvokeArgs()
        {
        }
        public static new GetIntegrationRuntimeStatusInvokeArgs Empty => new GetIntegrationRuntimeStatusInvokeArgs();
    }


    [OutputType]
    public sealed class GetIntegrationRuntimeStatusResult
    {
        /// <summary>
        /// The integration runtime name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Integration runtime properties.
        /// </summary>
        public readonly Union<Outputs.ManagedIntegrationRuntimeStatusResponse, Outputs.SelfHostedIntegrationRuntimeStatusResponse> Properties;

        [OutputConstructor]
        private GetIntegrationRuntimeStatusResult(
            string name,

            Union<Outputs.ManagedIntegrationRuntimeStatusResponse, Outputs.SelfHostedIntegrationRuntimeStatusResponse> properties)
        {
            Name = name;
            Properties = properties;
        }
    }
}
