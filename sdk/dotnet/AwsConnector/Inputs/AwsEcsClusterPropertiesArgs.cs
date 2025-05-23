// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector.Inputs
{

    /// <summary>
    /// Definition of awsEcsCluster
    /// </summary>
    public sealed class AwsEcsClusterPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Property arn
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        [Input("capacityProviders")]
        private InputList<string>? _capacityProviders;

        /// <summary>
        /// The short name of one or more capacity providers to associate with the cluster. A capacity provider must be associated with a cluster before it can be included as part of the default capacity provider strategy of the cluster or used in a capacity provider strategy when calling the [CreateService](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_CreateService.html) or [RunTask](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_RunTask.html) actions. If specifying a capacity provider that uses an Auto Scaling group, the capacity provider must be created but not associated with another cluster. New Auto Scaling group capacity providers can be created with the [CreateCapacityProvider](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_CreateCapacityProvider.html) API operation. To use a FARGATElong capacity provider, specify either the ``FARGATE`` or ``FARGATE_SPOT`` capacity providers. The FARGATElong capacity providers are available to all accounts and only need to be associated with a cluster to be used. The [PutCapacityProvider](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PutCapacityProvider.html) API operation is used to update the list of available capacity providers for a cluster after the cluster is created.
        /// </summary>
        public InputList<string> CapacityProviders
        {
            get => _capacityProviders ?? (_capacityProviders = new InputList<string>());
            set => _capacityProviders = value;
        }

        /// <summary>
        /// A user-generated string that you use to identify your cluster. If you don't specify a name, CFNlong generates a unique physical ID for the name.
        /// </summary>
        [Input("clusterName")]
        public Input<string>? ClusterName { get; set; }

        [Input("clusterSettings")]
        private InputList<Inputs.ClusterSettingsArgs>? _clusterSettings;

        /// <summary>
        /// The settings to use when creating a cluster. This parameter is used to turn on CloudWatch Container Insights for a cluster.
        /// </summary>
        public InputList<Inputs.ClusterSettingsArgs> ClusterSettings
        {
            get => _clusterSettings ?? (_clusterSettings = new InputList<Inputs.ClusterSettingsArgs>());
            set => _clusterSettings = value;
        }

        /// <summary>
        /// The execute command configuration for the cluster. The execute command configuration for the cluster.
        /// </summary>
        [Input("configuration")]
        public Input<Inputs.ClusterConfigurationArgs>? Configuration { get; set; }

        [Input("defaultCapacityProviderStrategy")]
        private InputList<Inputs.CapacityProviderStrategyItemArgs>? _defaultCapacityProviderStrategy;

        /// <summary>
        /// The default capacity provider strategy for the cluster. When services or tasks are run in the cluster with no launch type or capacity provider strategy specified, the default capacity provider strategy is used.
        /// </summary>
        public InputList<Inputs.CapacityProviderStrategyItemArgs> DefaultCapacityProviderStrategy
        {
            get => _defaultCapacityProviderStrategy ?? (_defaultCapacityProviderStrategy = new InputList<Inputs.CapacityProviderStrategyItemArgs>());
            set => _defaultCapacityProviderStrategy = value;
        }

        /// <summary>
        /// Use this parameter to set a default Service Connect namespace. After you set a default Service Connect namespace, any new services with Service Connect turned on that are created in the cluster are added as client services in the namespace. This setting only applies to new services that set the ``enabled`` parameter to ``true`` in the ``ServiceConnectConfiguration``. You can set the namespace of each service individually in the ``ServiceConnectConfiguration`` to override this default parameter. Tasks that run in a namespace can use short names to connect to services in the namespace. Tasks can connect to services across all of the clusters in the namespace. Tasks connect through a managed proxy container that collects logs and metrics for increased visibility. Only the tasks that Amazon ECS services create are supported with Service Connect. For more information, see [Service Connect](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-connect.html) in the *Amazon Elastic Container Service Developer Guide*. Use this parameter to set a default Service Connect namespace. After you set a default Service Connect namespace, any new services with Service Connect turned on that are created in the cluster are added as client services in the namespace. This setting only applies to new services that set the ``enabled`` parameter to ``true`` in the ``ServiceConnectConfiguration``. You can set the namespace of each service individually in the ``ServiceConnectConfiguration`` to override this default parameter. Tasks that run in a namespace can use short names to connect to services in the namespace. Tasks can connect to services across all of the clusters in the namespace. Tasks connect through a managed proxy container that collects logs and metrics for increased visibility. Only the tasks that Amazon ECS services create are supported with Service Connect. For more information, see [Service Connect](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-connect.html) in the *Amazon Elastic Container Service Developer Guide*.
        /// </summary>
        [Input("serviceConnectDefaults")]
        public Input<Inputs.ServiceConnectDefaultsArgs>? ServiceConnectDefaults { get; set; }

        [Input("tags")]
        private InputList<Inputs.TagArgs>? _tags;

        /// <summary>
        /// The metadata that you apply to the cluster to help you categorize and organize them. Each tag consists of a key and an optional value. You define both. The following basic restrictions apply to tags:  +  Maximum number of tags per resource - 50  +  For each resource, each tag key must be unique, and each tag key can have only one value.  +  Maximum key length - 128 Unicode characters in UTF-8  +  Maximum value length - 256 Unicode characters in UTF-8  +  If your tagging schema is used across multiple services and resources, remember that other services may have restrictions on allowed characters. Generally allowed characters are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.  +  Tag keys and values are case-sensitive.  +  Do not use ``aws:``, ``AWS:``, or any upper or lowercase combination of such as a prefix for either keys or values as it is reserved for AWS use. You cannot edit or delete tag keys or values with this prefix. Tags with this prefix do not count against your tags per resource limit.
        /// </summary>
        public InputList<Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.TagArgs>());
            set => _tags = value;
        }

        public AwsEcsClusterPropertiesArgs()
        {
        }
        public static new AwsEcsClusterPropertiesArgs Empty => new AwsEcsClusterPropertiesArgs();
    }
}
