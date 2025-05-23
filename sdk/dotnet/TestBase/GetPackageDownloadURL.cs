// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.TestBase
{
    public static class GetPackageDownloadURL
    {
        /// <summary>
        /// Gets the download URL of a package.
        /// 
        /// Uses Azure REST API version 2023-11-01-preview.
        /// 
        /// Other available API versions: 2022-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native testbase [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetPackageDownloadURLResult> InvokeAsync(GetPackageDownloadURLArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPackageDownloadURLResult>("azure-native:testbase:getPackageDownloadURL", args ?? new GetPackageDownloadURLArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the download URL of a package.
        /// 
        /// Uses Azure REST API version 2023-11-01-preview.
        /// 
        /// Other available API versions: 2022-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native testbase [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetPackageDownloadURLResult> Invoke(GetPackageDownloadURLInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPackageDownloadURLResult>("azure-native:testbase:getPackageDownloadURL", args ?? new GetPackageDownloadURLInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the download URL of a package.
        /// 
        /// Uses Azure REST API version 2023-11-01-preview.
        /// 
        /// Other available API versions: 2022-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native testbase [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetPackageDownloadURLResult> Invoke(GetPackageDownloadURLInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetPackageDownloadURLResult>("azure-native:testbase:getPackageDownloadURL", args ?? new GetPackageDownloadURLInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPackageDownloadURLArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The resource name of the Test Base Package.
        /// </summary>
        [Input("packageName", required: true)]
        public string PackageName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The resource name of the Test Base Account.
        /// </summary>
        [Input("testBaseAccountName", required: true)]
        public string TestBaseAccountName { get; set; } = null!;

        public GetPackageDownloadURLArgs()
        {
        }
        public static new GetPackageDownloadURLArgs Empty => new GetPackageDownloadURLArgs();
    }

    public sealed class GetPackageDownloadURLInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The resource name of the Test Base Package.
        /// </summary>
        [Input("packageName", required: true)]
        public Input<string> PackageName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The resource name of the Test Base Account.
        /// </summary>
        [Input("testBaseAccountName", required: true)]
        public Input<string> TestBaseAccountName { get; set; } = null!;

        public GetPackageDownloadURLInvokeArgs()
        {
        }
        public static new GetPackageDownloadURLInvokeArgs Empty => new GetPackageDownloadURLInvokeArgs();
    }


    [OutputType]
    public sealed class GetPackageDownloadURLResult
    {
        /// <summary>
        /// The download URL.
        /// </summary>
        public readonly string DownloadUrl;
        /// <summary>
        /// Expiry date of the download URL.
        /// </summary>
        public readonly string ExpirationTime;

        [OutputConstructor]
        private GetPackageDownloadURLResult(
            string downloadUrl,

            string expirationTime)
        {
            DownloadUrl = downloadUrl;
            ExpirationTime = expirationTime;
        }
    }
}
