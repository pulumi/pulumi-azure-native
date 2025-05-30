// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AppPlatform
{
    public static class GetApplicationAccelerator
    {
        /// <summary>
        /// Get the application accelerator.
        /// 
        /// Uses Azure REST API version 2024-01-01-preview.
        /// 
        /// Other available API versions: 2023-05-01-preview, 2023-07-01-preview, 2023-09-01-preview, 2023-11-01-preview, 2023-12-01, 2024-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native appplatform [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetApplicationAcceleratorResult> InvokeAsync(GetApplicationAcceleratorArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetApplicationAcceleratorResult>("azure-native:appplatform:getApplicationAccelerator", args ?? new GetApplicationAcceleratorArgs(), options.WithDefaults());

        /// <summary>
        /// Get the application accelerator.
        /// 
        /// Uses Azure REST API version 2024-01-01-preview.
        /// 
        /// Other available API versions: 2023-05-01-preview, 2023-07-01-preview, 2023-09-01-preview, 2023-11-01-preview, 2023-12-01, 2024-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native appplatform [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetApplicationAcceleratorResult> Invoke(GetApplicationAcceleratorInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetApplicationAcceleratorResult>("azure-native:appplatform:getApplicationAccelerator", args ?? new GetApplicationAcceleratorInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get the application accelerator.
        /// 
        /// Uses Azure REST API version 2024-01-01-preview.
        /// 
        /// Other available API versions: 2023-05-01-preview, 2023-07-01-preview, 2023-09-01-preview, 2023-11-01-preview, 2023-12-01, 2024-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native appplatform [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetApplicationAcceleratorResult> Invoke(GetApplicationAcceleratorInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetApplicationAcceleratorResult>("azure-native:appplatform:getApplicationAccelerator", args ?? new GetApplicationAcceleratorInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetApplicationAcceleratorArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the application accelerator.
        /// </summary>
        [Input("applicationAcceleratorName", required: true)]
        public string ApplicationAcceleratorName { get; set; } = null!;

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

        public GetApplicationAcceleratorArgs()
        {
        }
        public static new GetApplicationAcceleratorArgs Empty => new GetApplicationAcceleratorArgs();
    }

    public sealed class GetApplicationAcceleratorInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the application accelerator.
        /// </summary>
        [Input("applicationAcceleratorName", required: true)]
        public Input<string> ApplicationAcceleratorName { get; set; } = null!;

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

        public GetApplicationAcceleratorInvokeArgs()
        {
        }
        public static new GetApplicationAcceleratorInvokeArgs Empty => new GetApplicationAcceleratorInvokeArgs();
    }


    [OutputType]
    public sealed class GetApplicationAcceleratorResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Fully qualified resource Id for the resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Application accelerator properties payload
        /// </summary>
        public readonly Outputs.ApplicationAcceleratorPropertiesResponse Properties;
        /// <summary>
        /// Sku of the application accelerator resource
        /// </summary>
        public readonly Outputs.SkuResponse? Sku;
        /// <summary>
        /// Metadata pertaining to creation and last modification of the resource.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetApplicationAcceleratorResult(
            string azureApiVersion,

            string id,

            string name,

            Outputs.ApplicationAcceleratorPropertiesResponse properties,

            Outputs.SkuResponse? sku,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Id = id;
            Name = name;
            Properties = properties;
            Sku = sku;
            SystemData = systemData;
            Type = type;
        }
    }
}
