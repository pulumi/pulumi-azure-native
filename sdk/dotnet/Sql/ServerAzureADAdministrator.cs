// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Sql
{
    /// <summary>
    /// Azure Active Directory administrator.
    /// 
    /// Uses Azure REST API version 2023-08-01. In version 2.x of the Azure Native provider, it used API version 2021-11-01.
    /// 
    /// Other available API versions: 2014-04-01, 2018-06-01-preview, 2019-06-01-preview, 2020-02-02-preview, 2020-08-01-preview, 2020-11-01-preview, 2021-02-01-preview, 2021-05-01-preview, 2021-08-01-preview, 2021-11-01, 2021-11-01-preview, 2022-02-01-preview, 2022-05-01-preview, 2022-08-01-preview, 2022-11-01-preview, 2023-02-01-preview, 2023-05-01-preview, 2023-08-01-preview, 2024-05-01-preview, 2024-11-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:sql:ServerAzureADAdministrator")]
    public partial class ServerAzureADAdministrator : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Type of the sever administrator.
        /// </summary>
        [Output("administratorType")]
        public Output<string?> AdministratorType { get; private set; } = null!;

        /// <summary>
        /// Azure Active Directory only Authentication enabled.
        /// </summary>
        [Output("azureADOnlyAuthentication")]
        public Output<bool> AzureADOnlyAuthentication { get; private set; } = null!;

        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Login name of the server administrator.
        /// </summary>
        [Output("login")]
        public Output<string> Login { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// SID (object ID) of the server administrator.
        /// </summary>
        [Output("sid")]
        public Output<string> Sid { get; private set; } = null!;

        /// <summary>
        /// Tenant ID of the administrator.
        /// </summary>
        [Output("tenantId")]
        public Output<string?> TenantId { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ServerAzureADAdministrator resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServerAzureADAdministrator(string name, ServerAzureADAdministratorArgs args, CustomResourceOptions? options = null)
            : base("azure-native:sql:ServerAzureADAdministrator", name, args ?? new ServerAzureADAdministratorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServerAzureADAdministrator(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:sql:ServerAzureADAdministrator", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20140401:ServerAzureADAdministrator" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20180601preview:ServerAzureADAdministrator" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20190601preview:ServerAzureADAdministrator" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20200202preview:ServerAzureADAdministrator" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20200801preview:ServerAzureADAdministrator" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20201101preview:ServerAzureADAdministrator" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20210201preview:ServerAzureADAdministrator" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20210501preview:ServerAzureADAdministrator" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20210801preview:ServerAzureADAdministrator" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20211101:ServerAzureADAdministrator" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20211101preview:ServerAzureADAdministrator" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20220201preview:ServerAzureADAdministrator" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20220501preview:ServerAzureADAdministrator" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20220801preview:ServerAzureADAdministrator" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20221101preview:ServerAzureADAdministrator" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20230201preview:ServerAzureADAdministrator" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20230501preview:ServerAzureADAdministrator" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20230801:ServerAzureADAdministrator" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20230801preview:ServerAzureADAdministrator" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20240501preview:ServerAzureADAdministrator" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20241101preview:ServerAzureADAdministrator" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ServerAzureADAdministrator resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServerAzureADAdministrator Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ServerAzureADAdministrator(name, id, options);
        }
    }

    public sealed class ServerAzureADAdministratorArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of server active directory administrator.
        /// </summary>
        [Input("administratorName")]
        public Input<string>? AdministratorName { get; set; }

        /// <summary>
        /// Type of the sever administrator.
        /// </summary>
        [Input("administratorType")]
        public InputUnion<string, Pulumi.AzureNative.Sql.AdministratorType>? AdministratorType { get; set; }

        /// <summary>
        /// Login name of the server administrator.
        /// </summary>
        [Input("login", required: true)]
        public Input<string> Login { get; set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the server.
        /// </summary>
        [Input("serverName", required: true)]
        public Input<string> ServerName { get; set; } = null!;

        /// <summary>
        /// SID (object ID) of the server administrator.
        /// </summary>
        [Input("sid", required: true)]
        public Input<string> Sid { get; set; } = null!;

        /// <summary>
        /// Tenant ID of the administrator.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public ServerAzureADAdministratorArgs()
        {
        }
        public static new ServerAzureADAdministratorArgs Empty => new ServerAzureADAdministratorArgs();
    }
}
