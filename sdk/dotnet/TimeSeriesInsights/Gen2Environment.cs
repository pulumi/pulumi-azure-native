// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.TimeSeriesInsights
{
    /// <summary>
    /// An environment is a set of time-series data available for query, and is the top level Azure Time Series Insights resource. Gen2 environments do not have set data retention limits.
    /// 
    /// Uses Azure REST API version 2020-05-15. In version 2.x of the Azure Native provider, it used API version 2020-05-15.
    /// </summary>
    [AzureNativeResourceType("azure-native:timeseriesinsights:Gen2Environment")]
    public partial class Gen2Environment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// The time the resource was created.
        /// </summary>
        [Output("creationTime")]
        public Output<string> CreationTime { get; private set; } = null!;

        /// <summary>
        /// The fully qualified domain name used to access the environment data, e.g. to query the environment's events or upload reference data for the environment.
        /// </summary>
        [Output("dataAccessFqdn")]
        public Output<string> DataAccessFqdn { get; private set; } = null!;

        /// <summary>
        /// An id used to access the environment data, e.g. to query the environment's events or upload reference data for the environment.
        /// </summary>
        [Output("dataAccessId")]
        public Output<string> DataAccessId { get; private set; } = null!;

        /// <summary>
        /// The kind of the environment.
        /// Expected value is 'Gen2'.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Resource location
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Provisioning state of the resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The sku determines the type of environment, either Gen1 (S1 or S2) or Gen2 (L1). For Gen1 environments the sku determines the capacity of the environment, the ingress rate, and the billing rate.
        /// </summary>
        [Output("sku")]
        public Output<Outputs.SkuResponse> Sku { get; private set; } = null!;

        /// <summary>
        /// An object that represents the status of the environment, and its internal state in the Time Series Insights service.
        /// </summary>
        [Output("status")]
        public Output<Outputs.EnvironmentStatusResponse> Status { get; private set; } = null!;

        /// <summary>
        /// The storage configuration provides the connection details that allows the Time Series Insights service to connect to the customer storage account that is used to store the environment's data.
        /// </summary>
        [Output("storageConfiguration")]
        public Output<Outputs.Gen2StorageConfigurationOutputResponse> StorageConfiguration { get; private set; } = null!;

        /// <summary>
        /// Resource tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The list of event properties which will be used to define the environment's time series id.
        /// </summary>
        [Output("timeSeriesIdProperties")]
        public Output<ImmutableArray<Outputs.TimeSeriesIdPropertyResponse>> TimeSeriesIdProperties { get; private set; } = null!;

        /// <summary>
        /// Resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The warm store configuration provides the details to create a warm store cache that will retain a copy of the environment's data available for faster query.
        /// </summary>
        [Output("warmStoreConfiguration")]
        public Output<Outputs.WarmStoreConfigurationPropertiesResponse?> WarmStoreConfiguration { get; private set; } = null!;


        /// <summary>
        /// Create a Gen2Environment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Gen2Environment(string name, Gen2EnvironmentArgs args, CustomResourceOptions? options = null)
            : base("azure-native:timeseriesinsights:Gen2Environment", name, MakeArgs(args), MakeResourceOptions(options, ""))
        {
        }

        private Gen2Environment(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:timeseriesinsights:Gen2Environment", name, null, MakeResourceOptions(options, id))
        {
        }

        private static Gen2EnvironmentArgs MakeArgs(Gen2EnvironmentArgs args)
        {
            args ??= new Gen2EnvironmentArgs();
            args.Kind = "Gen2";
            return args;
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:timeseriesinsights/v20170228preview:Gen2Environment" },
                    new global::Pulumi.Alias { Type = "azure-native:timeseriesinsights/v20171115:Gen2Environment" },
                    new global::Pulumi.Alias { Type = "azure-native:timeseriesinsights/v20180815preview:Gen2Environment" },
                    new global::Pulumi.Alias { Type = "azure-native:timeseriesinsights/v20200515:Gen2Environment" },
                    new global::Pulumi.Alias { Type = "azure-native:timeseriesinsights/v20210331preview:Gen2Environment" },
                    new global::Pulumi.Alias { Type = "azure-native:timeseriesinsights/v20210630preview:Gen1Environment" },
                    new global::Pulumi.Alias { Type = "azure-native:timeseriesinsights/v20210630preview:Gen2Environment" },
                    new global::Pulumi.Alias { Type = "azure-native:timeseriesinsights:Gen1Environment" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Gen2Environment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Gen2Environment Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Gen2Environment(name, id, options);
        }
    }

    public sealed class Gen2EnvironmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the environment
        /// </summary>
        [Input("environmentName")]
        public Input<string>? EnvironmentName { get; set; }

        /// <summary>
        /// The kind of the environment.
        /// Expected value is 'Gen2'.
        /// </summary>
        [Input("kind", required: true)]
        public Input<string> Kind { get; set; } = null!;

        /// <summary>
        /// The location of the resource.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Name of an Azure Resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The sku determines the type of environment, either Gen1 (S1 or S2) or Gen2 (L1). For Gen1 environments the sku determines the capacity of the environment, the ingress rate, and the billing rate.
        /// </summary>
        [Input("sku", required: true)]
        public Input<Inputs.SkuArgs> Sku { get; set; } = null!;

        /// <summary>
        /// The storage configuration provides the connection details that allows the Time Series Insights service to connect to the customer storage account that is used to store the environment's data.
        /// </summary>
        [Input("storageConfiguration", required: true)]
        public Input<Inputs.Gen2StorageConfigurationInputArgs> StorageConfiguration { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value pairs of additional properties for the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("timeSeriesIdProperties", required: true)]
        private InputList<Inputs.TimeSeriesIdPropertyArgs>? _timeSeriesIdProperties;

        /// <summary>
        /// The list of event properties which will be used to define the environment's time series id.
        /// </summary>
        public InputList<Inputs.TimeSeriesIdPropertyArgs> TimeSeriesIdProperties
        {
            get => _timeSeriesIdProperties ?? (_timeSeriesIdProperties = new InputList<Inputs.TimeSeriesIdPropertyArgs>());
            set => _timeSeriesIdProperties = value;
        }

        /// <summary>
        /// The warm store configuration provides the details to create a warm store cache that will retain a copy of the environment's data available for faster query.
        /// </summary>
        [Input("warmStoreConfiguration")]
        public Input<Inputs.WarmStoreConfigurationPropertiesArgs>? WarmStoreConfiguration { get; set; }

        public Gen2EnvironmentArgs()
        {
        }
        public static new Gen2EnvironmentArgs Empty => new Gen2EnvironmentArgs();
    }
}
