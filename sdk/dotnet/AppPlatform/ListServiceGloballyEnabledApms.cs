// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AppPlatform
{
    public static class ListServiceGloballyEnabledApms
    {
        /// <summary>
        /// List globally enabled APMs for a Service.
        /// 
        /// Uses Azure REST API version 2024-01-01-preview.
        /// 
        /// Other available API versions: 2023-05-01-preview, 2023-07-01-preview, 2023-09-01-preview, 2023-11-01-preview, 2023-12-01, 2024-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native appplatform [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<ListServiceGloballyEnabledApmsResult> InvokeAsync(ListServiceGloballyEnabledApmsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListServiceGloballyEnabledApmsResult>("azure-native:appplatform:listServiceGloballyEnabledApms", args ?? new ListServiceGloballyEnabledApmsArgs(), options.WithDefaults());

        /// <summary>
        /// List globally enabled APMs for a Service.
        /// 
        /// Uses Azure REST API version 2024-01-01-preview.
        /// 
        /// Other available API versions: 2023-05-01-preview, 2023-07-01-preview, 2023-09-01-preview, 2023-11-01-preview, 2023-12-01, 2024-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native appplatform [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListServiceGloballyEnabledApmsResult> Invoke(ListServiceGloballyEnabledApmsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListServiceGloballyEnabledApmsResult>("azure-native:appplatform:listServiceGloballyEnabledApms", args ?? new ListServiceGloballyEnabledApmsInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// List globally enabled APMs for a Service.
        /// 
        /// Uses Azure REST API version 2024-01-01-preview.
        /// 
        /// Other available API versions: 2023-05-01-preview, 2023-07-01-preview, 2023-09-01-preview, 2023-11-01-preview, 2023-12-01, 2024-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native appplatform [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListServiceGloballyEnabledApmsResult> Invoke(ListServiceGloballyEnabledApmsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListServiceGloballyEnabledApmsResult>("azure-native:appplatform:listServiceGloballyEnabledApms", args ?? new ListServiceGloballyEnabledApmsInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListServiceGloballyEnabledApmsArgs : global::Pulumi.InvokeArgs
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

        public ListServiceGloballyEnabledApmsArgs()
        {
        }
        public static new ListServiceGloballyEnabledApmsArgs Empty => new ListServiceGloballyEnabledApmsArgs();
    }

    public sealed class ListServiceGloballyEnabledApmsInvokeArgs : global::Pulumi.InvokeArgs
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

        public ListServiceGloballyEnabledApmsInvokeArgs()
        {
        }
        public static new ListServiceGloballyEnabledApmsInvokeArgs Empty => new ListServiceGloballyEnabledApmsInvokeArgs();
    }


    [OutputType]
    public sealed class ListServiceGloballyEnabledApmsResult
    {
        /// <summary>
        /// Collection of the globally enabled APMs
        /// </summary>
        public readonly ImmutableArray<string> Value;

        [OutputConstructor]
        private ListServiceGloballyEnabledApmsResult(ImmutableArray<string> value)
        {
            Value = value;
        }
    }
}
