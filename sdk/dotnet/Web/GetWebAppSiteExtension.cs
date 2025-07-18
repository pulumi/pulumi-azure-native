// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web
{
    public static class GetWebAppSiteExtension
    {
        /// <summary>
        /// Description for Get site extension information by its ID for a web site, or a deployment slot.
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// 
        /// Other available API versions: 2016-08-01, 2018-02-01, 2018-11-01, 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetWebAppSiteExtensionResult> InvokeAsync(GetWebAppSiteExtensionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetWebAppSiteExtensionResult>("azure-native:web:getWebAppSiteExtension", args ?? new GetWebAppSiteExtensionArgs(), options.WithDefaults());

        /// <summary>
        /// Description for Get site extension information by its ID for a web site, or a deployment slot.
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// 
        /// Other available API versions: 2016-08-01, 2018-02-01, 2018-11-01, 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetWebAppSiteExtensionResult> Invoke(GetWebAppSiteExtensionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetWebAppSiteExtensionResult>("azure-native:web:getWebAppSiteExtension", args ?? new GetWebAppSiteExtensionInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Description for Get site extension information by its ID for a web site, or a deployment slot.
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// 
        /// Other available API versions: 2016-08-01, 2018-02-01, 2018-11-01, 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetWebAppSiteExtensionResult> Invoke(GetWebAppSiteExtensionInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetWebAppSiteExtensionResult>("azure-native:web:getWebAppSiteExtension", args ?? new GetWebAppSiteExtensionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWebAppSiteExtensionArgs : global::Pulumi.InvokeArgs
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
        /// Site extension name.
        /// </summary>
        [Input("siteExtensionId", required: true)]
        public string SiteExtensionId { get; set; } = null!;

        public GetWebAppSiteExtensionArgs()
        {
        }
        public static new GetWebAppSiteExtensionArgs Empty => new GetWebAppSiteExtensionArgs();
    }

    public sealed class GetWebAppSiteExtensionInvokeArgs : global::Pulumi.InvokeArgs
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
        /// Site extension name.
        /// </summary>
        [Input("siteExtensionId", required: true)]
        public Input<string> SiteExtensionId { get; set; } = null!;

        public GetWebAppSiteExtensionInvokeArgs()
        {
        }
        public static new GetWebAppSiteExtensionInvokeArgs Empty => new GetWebAppSiteExtensionInvokeArgs();
    }


    [OutputType]
    public sealed class GetWebAppSiteExtensionResult
    {
        /// <summary>
        /// List of authors.
        /// </summary>
        public readonly ImmutableArray<string> Authors;
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Site Extension comment.
        /// </summary>
        public readonly string? Comment;
        /// <summary>
        /// Detailed description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Count of downloads.
        /// </summary>
        public readonly int? DownloadCount;
        /// <summary>
        /// Site extension ID.
        /// </summary>
        public readonly string? ExtensionId;
        /// <summary>
        /// Site extension type.
        /// </summary>
        public readonly string? ExtensionType;
        /// <summary>
        /// Extension URL.
        /// </summary>
        public readonly string? ExtensionUrl;
        /// <summary>
        /// Feed URL.
        /// </summary>
        public readonly string? FeedUrl;
        /// <summary>
        /// Icon URL.
        /// </summary>
        public readonly string? IconUrl;
        /// <summary>
        /// Resource Id.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Installed timestamp.
        /// </summary>
        public readonly string? InstalledDateTime;
        /// <summary>
        /// Installer command line parameters.
        /// </summary>
        public readonly string? InstallerCommandLineParams;
        /// <summary>
        /// Kind of resource.
        /// </summary>
        public readonly string? Kind;
        /// <summary>
        /// License URL.
        /// </summary>
        public readonly string? LicenseUrl;
        /// <summary>
        /// &lt;code&gt;true&lt;/code&gt; if the local version is the latest version; &lt;code&gt;false&lt;/code&gt; otherwise.
        /// </summary>
        public readonly bool? LocalIsLatestVersion;
        /// <summary>
        /// Local path.
        /// </summary>
        public readonly string? LocalPath;
        /// <summary>
        /// Resource Name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Project URL.
        /// </summary>
        public readonly string? ProjectUrl;
        /// <summary>
        /// Provisioning state.
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// Published timestamp.
        /// </summary>
        public readonly string? PublishedDateTime;
        /// <summary>
        /// Summary description.
        /// </summary>
        public readonly string? Summary;
        public readonly string? Title;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Version information.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private GetWebAppSiteExtensionResult(
            ImmutableArray<string> authors,

            string azureApiVersion,

            string? comment,

            string? description,

            int? downloadCount,

            string? extensionId,

            string? extensionType,

            string? extensionUrl,

            string? feedUrl,

            string? iconUrl,

            string id,

            string? installedDateTime,

            string? installerCommandLineParams,

            string? kind,

            string? licenseUrl,

            bool? localIsLatestVersion,

            string? localPath,

            string name,

            string? projectUrl,

            string? provisioningState,

            string? publishedDateTime,

            string? summary,

            string? title,

            string type,

            string? version)
        {
            Authors = authors;
            AzureApiVersion = azureApiVersion;
            Comment = comment;
            Description = description;
            DownloadCount = downloadCount;
            ExtensionId = extensionId;
            ExtensionType = extensionType;
            ExtensionUrl = extensionUrl;
            FeedUrl = feedUrl;
            IconUrl = iconUrl;
            Id = id;
            InstalledDateTime = installedDateTime;
            InstallerCommandLineParams = installerCommandLineParams;
            Kind = kind;
            LicenseUrl = licenseUrl;
            LocalIsLatestVersion = localIsLatestVersion;
            LocalPath = localPath;
            Name = name;
            ProjectUrl = projectUrl;
            ProvisioningState = provisioningState;
            PublishedDateTime = publishedDateTime;
            Summary = summary;
            Title = title;
            Type = type;
            Version = version;
        }
    }
}
