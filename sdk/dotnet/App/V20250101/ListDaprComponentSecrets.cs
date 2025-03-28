// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.App.V20250101
{
    public static class ListDaprComponentSecrets
    {
        /// <summary>
        /// Dapr component Secrets Collection for ListSecrets Action.
        /// </summary>
        public static Task<ListDaprComponentSecretsResult> InvokeAsync(ListDaprComponentSecretsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListDaprComponentSecretsResult>("azure-native:app/v20250101:listDaprComponentSecrets", args ?? new ListDaprComponentSecretsArgs(), options.WithDefaults());

        /// <summary>
        /// Dapr component Secrets Collection for ListSecrets Action.
        /// </summary>
        public static Output<ListDaprComponentSecretsResult> Invoke(ListDaprComponentSecretsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListDaprComponentSecretsResult>("azure-native:app/v20250101:listDaprComponentSecrets", args ?? new ListDaprComponentSecretsInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Dapr component Secrets Collection for ListSecrets Action.
        /// </summary>
        public static Output<ListDaprComponentSecretsResult> Invoke(ListDaprComponentSecretsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListDaprComponentSecretsResult>("azure-native:app/v20250101:listDaprComponentSecrets", args ?? new ListDaprComponentSecretsInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListDaprComponentSecretsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the Dapr Component.
        /// </summary>
        [Input("componentName", required: true)]
        public string ComponentName { get; set; } = null!;

        /// <summary>
        /// Name of the Managed Environment.
        /// </summary>
        [Input("environmentName", required: true)]
        public string EnvironmentName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public ListDaprComponentSecretsArgs()
        {
        }
        public static new ListDaprComponentSecretsArgs Empty => new ListDaprComponentSecretsArgs();
    }

    public sealed class ListDaprComponentSecretsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the Dapr Component.
        /// </summary>
        [Input("componentName", required: true)]
        public Input<string> ComponentName { get; set; } = null!;

        /// <summary>
        /// Name of the Managed Environment.
        /// </summary>
        [Input("environmentName", required: true)]
        public Input<string> EnvironmentName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public ListDaprComponentSecretsInvokeArgs()
        {
        }
        public static new ListDaprComponentSecretsInvokeArgs Empty => new ListDaprComponentSecretsInvokeArgs();
    }


    [OutputType]
    public sealed class ListDaprComponentSecretsResult
    {
        /// <summary>
        /// Collection of secrets used by a Dapr component
        /// </summary>
        public readonly ImmutableArray<Outputs.DaprSecretResponse> Value;

        [OutputConstructor]
        private ListDaprComponentSecretsResult(ImmutableArray<Outputs.DaprSecretResponse> value)
        {
            Value = value;
        }
    }
}
