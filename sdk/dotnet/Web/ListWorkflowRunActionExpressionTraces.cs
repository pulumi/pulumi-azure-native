// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web
{
    public static class ListWorkflowRunActionExpressionTraces
    {
        /// <summary>
        /// Lists a workflow run expression trace.
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// 
        /// Other available API versions: 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<ListWorkflowRunActionExpressionTracesResult> InvokeAsync(ListWorkflowRunActionExpressionTracesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListWorkflowRunActionExpressionTracesResult>("azure-native:web:listWorkflowRunActionExpressionTraces", args ?? new ListWorkflowRunActionExpressionTracesArgs(), options.WithDefaults());

        /// <summary>
        /// Lists a workflow run expression trace.
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// 
        /// Other available API versions: 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListWorkflowRunActionExpressionTracesResult> Invoke(ListWorkflowRunActionExpressionTracesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListWorkflowRunActionExpressionTracesResult>("azure-native:web:listWorkflowRunActionExpressionTraces", args ?? new ListWorkflowRunActionExpressionTracesInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Lists a workflow run expression trace.
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// 
        /// Other available API versions: 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListWorkflowRunActionExpressionTracesResult> Invoke(ListWorkflowRunActionExpressionTracesInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListWorkflowRunActionExpressionTracesResult>("azure-native:web:listWorkflowRunActionExpressionTraces", args ?? new ListWorkflowRunActionExpressionTracesInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListWorkflowRunActionExpressionTracesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The workflow action name.
        /// </summary>
        [Input("actionName", required: true)]
        public string ActionName { get; set; } = null!;

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
        /// The workflow run name.
        /// </summary>
        [Input("runName", required: true)]
        public string RunName { get; set; } = null!;

        /// <summary>
        /// The workflow name.
        /// </summary>
        [Input("workflowName", required: true)]
        public string WorkflowName { get; set; } = null!;

        public ListWorkflowRunActionExpressionTracesArgs()
        {
        }
        public static new ListWorkflowRunActionExpressionTracesArgs Empty => new ListWorkflowRunActionExpressionTracesArgs();
    }

    public sealed class ListWorkflowRunActionExpressionTracesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The workflow action name.
        /// </summary>
        [Input("actionName", required: true)]
        public Input<string> ActionName { get; set; } = null!;

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
        /// The workflow run name.
        /// </summary>
        [Input("runName", required: true)]
        public Input<string> RunName { get; set; } = null!;

        /// <summary>
        /// The workflow name.
        /// </summary>
        [Input("workflowName", required: true)]
        public Input<string> WorkflowName { get; set; } = null!;

        public ListWorkflowRunActionExpressionTracesInvokeArgs()
        {
        }
        public static new ListWorkflowRunActionExpressionTracesInvokeArgs Empty => new ListWorkflowRunActionExpressionTracesInvokeArgs();
    }


    [OutputType]
    public sealed class ListWorkflowRunActionExpressionTracesResult
    {
        public readonly ImmutableArray<Outputs.ExpressionRootResponse> Inputs;
        /// <summary>
        /// The link used to get the next page of recommendations.
        /// </summary>
        public readonly string? NextLink;
        public readonly object? Value;

        [OutputConstructor]
        private ListWorkflowRunActionExpressionTracesResult(
            ImmutableArray<Outputs.ExpressionRootResponse> inputs,

            string? nextLink,

            object? value)
        {
            Inputs = inputs;
            NextLink = nextLink;
            Value = value;
        }
    }
}
