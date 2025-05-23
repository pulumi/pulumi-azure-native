# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['KubernetesClusterFeatureArgs', 'KubernetesClusterFeature']

@pulumi.input_type
class KubernetesClusterFeatureArgs:
    def __init__(__self__, *,
                 kubernetes_cluster_name: pulumi.Input[builtins.str],
                 resource_group_name: pulumi.Input[builtins.str],
                 feature_name: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 options: Optional[pulumi.Input[Sequence[pulumi.Input['StringKeyValuePairArgs']]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None):
        """
        The set of arguments for constructing a KubernetesClusterFeature resource.
        :param pulumi.Input[builtins.str] kubernetes_cluster_name: The name of the Kubernetes cluster.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] feature_name: The name of the feature.
        :param pulumi.Input[builtins.str] location: The geo-location where the resource lives
        :param pulumi.Input[Sequence[pulumi.Input['StringKeyValuePairArgs']]] options: The configured options for the feature.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        """
        pulumi.set(__self__, "kubernetes_cluster_name", kubernetes_cluster_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if feature_name is not None:
            pulumi.set(__self__, "feature_name", feature_name)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if options is not None:
            pulumi.set(__self__, "options", options)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="kubernetesClusterName")
    def kubernetes_cluster_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the Kubernetes cluster.
        """
        return pulumi.get(self, "kubernetes_cluster_name")

    @kubernetes_cluster_name.setter
    def kubernetes_cluster_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "kubernetes_cluster_name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the resource group. The name is case insensitive.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="featureName")
    def feature_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the feature.
        """
        return pulumi.get(self, "feature_name")

    @feature_name.setter
    def feature_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "feature_name", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def options(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['StringKeyValuePairArgs']]]]:
        """
        The configured options for the feature.
        """
        return pulumi.get(self, "options")

    @options.setter
    def options(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['StringKeyValuePairArgs']]]]):
        pulumi.set(self, "options", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.type_token("azure-native:networkcloud:KubernetesClusterFeature")
class KubernetesClusterFeature(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 feature_name: Optional[pulumi.Input[builtins.str]] = None,
                 kubernetes_cluster_name: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 options: Optional[pulumi.Input[Sequence[pulumi.Input[Union['StringKeyValuePairArgs', 'StringKeyValuePairArgsDict']]]]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        """
        Uses Azure REST API version 2025-02-01. In version 2.x of the Azure Native provider, it used API version 2024-06-01-preview.

        Other available API versions: 2024-07-01, 2024-10-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native networkcloud [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] feature_name: The name of the feature.
        :param pulumi.Input[builtins.str] kubernetes_cluster_name: The name of the Kubernetes cluster.
        :param pulumi.Input[builtins.str] location: The geo-location where the resource lives
        :param pulumi.Input[Sequence[pulumi.Input[Union['StringKeyValuePairArgs', 'StringKeyValuePairArgsDict']]]] options: The configured options for the feature.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: KubernetesClusterFeatureArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Uses Azure REST API version 2025-02-01. In version 2.x of the Azure Native provider, it used API version 2024-06-01-preview.

        Other available API versions: 2024-07-01, 2024-10-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native networkcloud [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param KubernetesClusterFeatureArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(KubernetesClusterFeatureArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 feature_name: Optional[pulumi.Input[builtins.str]] = None,
                 kubernetes_cluster_name: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 options: Optional[pulumi.Input[Sequence[pulumi.Input[Union['StringKeyValuePairArgs', 'StringKeyValuePairArgsDict']]]]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = KubernetesClusterFeatureArgs.__new__(KubernetesClusterFeatureArgs)

            __props__.__dict__["feature_name"] = feature_name
            if kubernetes_cluster_name is None and not opts.urn:
                raise TypeError("Missing required property 'kubernetes_cluster_name'")
            __props__.__dict__["kubernetes_cluster_name"] = kubernetes_cluster_name
            __props__.__dict__["location"] = location
            __props__.__dict__["options"] = options
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["availability_lifecycle"] = None
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["detailed_status"] = None
            __props__.__dict__["detailed_status_message"] = None
            __props__.__dict__["etag"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["provisioning_state"] = None
            __props__.__dict__["required"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
            __props__.__dict__["version"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:networkcloud/v20240601preview:KubernetesClusterFeature"), pulumi.Alias(type_="azure-native:networkcloud/v20240701:KubernetesClusterFeature"), pulumi.Alias(type_="azure-native:networkcloud/v20241001preview:KubernetesClusterFeature"), pulumi.Alias(type_="azure-native:networkcloud/v20250201:KubernetesClusterFeature")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(KubernetesClusterFeature, __self__).__init__(
            'azure-native:networkcloud:KubernetesClusterFeature',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'KubernetesClusterFeature':
        """
        Get an existing KubernetesClusterFeature resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = KubernetesClusterFeatureArgs.__new__(KubernetesClusterFeatureArgs)

        __props__.__dict__["availability_lifecycle"] = None
        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["detailed_status"] = None
        __props__.__dict__["detailed_status_message"] = None
        __props__.__dict__["etag"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["options"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["required"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["version"] = None
        return KubernetesClusterFeature(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="availabilityLifecycle")
    def availability_lifecycle(self) -> pulumi.Output[builtins.str]:
        """
        The lifecycle indicator of the feature.
        """
        return pulumi.get(self, "availability_lifecycle")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="detailedStatus")
    def detailed_status(self) -> pulumi.Output[builtins.str]:
        """
        The detailed status of the feature.
        """
        return pulumi.get(self, "detailed_status")

    @property
    @pulumi.getter(name="detailedStatusMessage")
    def detailed_status_message(self) -> pulumi.Output[builtins.str]:
        """
        The descriptive message for the detailed status of the feature.
        """
        return pulumi.get(self, "detailed_status_message")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[builtins.str]:
        """
        Resource ETag.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[builtins.str]:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def options(self) -> pulumi.Output[Optional[Sequence['outputs.StringKeyValuePairResponse']]]:
        """
        The configured options for the feature.
        """
        return pulumi.get(self, "options")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[builtins.str]:
        """
        The provisioning state of the Kubernetes cluster feature.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def required(self) -> pulumi.Output[builtins.str]:
        """
        The indicator of if the feature is required or optional. Optional features may be deleted by the user, while required features are managed with the kubernetes cluster lifecycle.
        """
        return pulumi.get(self, "required")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> pulumi.Output['outputs.SystemDataResponse']:
        """
        Azure Resource Manager metadata containing createdBy and modifiedBy information.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, builtins.str]]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[builtins.str]:
        """
        The version of the feature.
        """
        return pulumi.get(self, "version")

