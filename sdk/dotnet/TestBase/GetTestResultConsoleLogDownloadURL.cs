// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.TestBase
{
    public static class GetTestResultConsoleLogDownloadURL
    {
        /// <summary>
        /// Gets the download URL of the test execution console log file.
        /// 
        /// Uses Azure REST API version 2023-11-01-preview.
        /// 
        /// Other available API versions: 2022-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native testbase [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetTestResultConsoleLogDownloadURLResult> InvokeAsync(GetTestResultConsoleLogDownloadURLArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTestResultConsoleLogDownloadURLResult>("azure-native:testbase:getTestResultConsoleLogDownloadURL", args ?? new GetTestResultConsoleLogDownloadURLArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the download URL of the test execution console log file.
        /// 
        /// Uses Azure REST API version 2023-11-01-preview.
        /// 
        /// Other available API versions: 2022-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native testbase [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetTestResultConsoleLogDownloadURLResult> Invoke(GetTestResultConsoleLogDownloadURLInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTestResultConsoleLogDownloadURLResult>("azure-native:testbase:getTestResultConsoleLogDownloadURL", args ?? new GetTestResultConsoleLogDownloadURLInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the download URL of the test execution console log file.
        /// 
        /// Uses Azure REST API version 2023-11-01-preview.
        /// 
        /// Other available API versions: 2022-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native testbase [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetTestResultConsoleLogDownloadURLResult> Invoke(GetTestResultConsoleLogDownloadURLInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetTestResultConsoleLogDownloadURLResult>("azure-native:testbase:getTestResultConsoleLogDownloadURL", args ?? new GetTestResultConsoleLogDownloadURLInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTestResultConsoleLogDownloadURLArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The log file name corresponding to the download URL.
        /// </summary>
        [Input("logFileName", required: true)]
        public string LogFileName { get; set; } = null!;

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

        /// <summary>
        /// The Test Result Name. It equals to TestResult-{TestResultId} string.
        /// </summary>
        [Input("testResultName", required: true)]
        public string TestResultName { get; set; } = null!;

        public GetTestResultConsoleLogDownloadURLArgs()
        {
        }
        public static new GetTestResultConsoleLogDownloadURLArgs Empty => new GetTestResultConsoleLogDownloadURLArgs();
    }

    public sealed class GetTestResultConsoleLogDownloadURLInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The log file name corresponding to the download URL.
        /// </summary>
        [Input("logFileName", required: true)]
        public Input<string> LogFileName { get; set; } = null!;

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

        /// <summary>
        /// The Test Result Name. It equals to TestResult-{TestResultId} string.
        /// </summary>
        [Input("testResultName", required: true)]
        public Input<string> TestResultName { get; set; } = null!;

        public GetTestResultConsoleLogDownloadURLInvokeArgs()
        {
        }
        public static new GetTestResultConsoleLogDownloadURLInvokeArgs Empty => new GetTestResultConsoleLogDownloadURLInvokeArgs();
    }


    [OutputType]
    public sealed class GetTestResultConsoleLogDownloadURLResult
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
        private GetTestResultConsoleLogDownloadURLResult(
            string downloadUrl,

            string expirationTime)
        {
            DownloadUrl = downloadUrl;
            ExpirationTime = expirationTime;
        }
    }
}
