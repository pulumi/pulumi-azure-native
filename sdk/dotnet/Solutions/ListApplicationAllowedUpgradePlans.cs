// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Solutions
{
    public static class ListApplicationAllowedUpgradePlans
    {
        /// <summary>
        /// List allowed upgrade plans for application.
        /// 
        /// Uses Azure REST API version 2021-07-01.
        /// 
        /// Other available API versions: 2023-12-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native solutions [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<ListApplicationAllowedUpgradePlansResult> InvokeAsync(ListApplicationAllowedUpgradePlansArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListApplicationAllowedUpgradePlansResult>("azure-native:solutions:listApplicationAllowedUpgradePlans", args ?? new ListApplicationAllowedUpgradePlansArgs(), options.WithDefaults());

        /// <summary>
        /// List allowed upgrade plans for application.
        /// 
        /// Uses Azure REST API version 2021-07-01.
        /// 
        /// Other available API versions: 2023-12-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native solutions [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListApplicationAllowedUpgradePlansResult> Invoke(ListApplicationAllowedUpgradePlansInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListApplicationAllowedUpgradePlansResult>("azure-native:solutions:listApplicationAllowedUpgradePlans", args ?? new ListApplicationAllowedUpgradePlansInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// List allowed upgrade plans for application.
        /// 
        /// Uses Azure REST API version 2021-07-01.
        /// 
        /// Other available API versions: 2023-12-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native solutions [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListApplicationAllowedUpgradePlansResult> Invoke(ListApplicationAllowedUpgradePlansInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListApplicationAllowedUpgradePlansResult>("azure-native:solutions:listApplicationAllowedUpgradePlans", args ?? new ListApplicationAllowedUpgradePlansInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListApplicationAllowedUpgradePlansArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the managed application.
        /// </summary>
        [Input("applicationName", required: true)]
        public string ApplicationName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public ListApplicationAllowedUpgradePlansArgs()
        {
        }
        public static new ListApplicationAllowedUpgradePlansArgs Empty => new ListApplicationAllowedUpgradePlansArgs();
    }

    public sealed class ListApplicationAllowedUpgradePlansInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the managed application.
        /// </summary>
        [Input("applicationName", required: true)]
        public Input<string> ApplicationName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public ListApplicationAllowedUpgradePlansInvokeArgs()
        {
        }
        public static new ListApplicationAllowedUpgradePlansInvokeArgs Empty => new ListApplicationAllowedUpgradePlansInvokeArgs();
    }


    [OutputType]
    public sealed class ListApplicationAllowedUpgradePlansResult
    {
        /// <summary>
        /// The array of plans.
        /// </summary>
        public readonly ImmutableArray<Outputs.PlanResponse> Value;

        [OutputConstructor]
        private ListApplicationAllowedUpgradePlansResult(ImmutableArray<Outputs.PlanResponse> value)
        {
            Value = value;
        }
    }
}
