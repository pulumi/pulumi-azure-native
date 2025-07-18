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
from ._enums import *
from ._inputs import *

__all__ = ['BatchDeploymentInitArgs', 'BatchDeployment']

@pulumi.input_type
class BatchDeploymentInitArgs:
    def __init__(__self__, *,
                 batch_deployment_properties: pulumi.Input['BatchDeploymentArgs'],
                 endpoint_name: pulumi.Input[builtins.str],
                 resource_group_name: pulumi.Input[builtins.str],
                 workspace_name: pulumi.Input[builtins.str],
                 deployment_name: Optional[pulumi.Input[builtins.str]] = None,
                 identity: Optional[pulumi.Input['ManagedServiceIdentityArgs']] = None,
                 kind: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 sku: Optional[pulumi.Input['SkuArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None):
        """
        The set of arguments for constructing a BatchDeployment resource.
        :param pulumi.Input['BatchDeploymentArgs'] batch_deployment_properties: [Required] Additional attributes of the entity.
        :param pulumi.Input[builtins.str] endpoint_name: Inference endpoint name
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] workspace_name: Name of Azure Machine Learning workspace.
        :param pulumi.Input[builtins.str] deployment_name: The identifier for the Batch inference deployment.
        :param pulumi.Input['ManagedServiceIdentityArgs'] identity: Managed service identity (system assigned and/or user assigned identities)
        :param pulumi.Input[builtins.str] kind: Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type.
        :param pulumi.Input[builtins.str] location: The geo-location where the resource lives
        :param pulumi.Input['SkuArgs'] sku: Sku details required for ARM contract for Autoscaling.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        """
        pulumi.set(__self__, "batch_deployment_properties", batch_deployment_properties)
        pulumi.set(__self__, "endpoint_name", endpoint_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "workspace_name", workspace_name)
        if deployment_name is not None:
            pulumi.set(__self__, "deployment_name", deployment_name)
        if identity is not None:
            pulumi.set(__self__, "identity", identity)
        if kind is not None:
            pulumi.set(__self__, "kind", kind)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if sku is not None:
            pulumi.set(__self__, "sku", sku)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="batchDeploymentProperties")
    def batch_deployment_properties(self) -> pulumi.Input['BatchDeploymentArgs']:
        """
        [Required] Additional attributes of the entity.
        """
        return pulumi.get(self, "batch_deployment_properties")

    @batch_deployment_properties.setter
    def batch_deployment_properties(self, value: pulumi.Input['BatchDeploymentArgs']):
        pulumi.set(self, "batch_deployment_properties", value)

    @property
    @pulumi.getter(name="endpointName")
    def endpoint_name(self) -> pulumi.Input[builtins.str]:
        """
        Inference endpoint name
        """
        return pulumi.get(self, "endpoint_name")

    @endpoint_name.setter
    def endpoint_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "endpoint_name", value)

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
    @pulumi.getter(name="workspaceName")
    def workspace_name(self) -> pulumi.Input[builtins.str]:
        """
        Name of Azure Machine Learning workspace.
        """
        return pulumi.get(self, "workspace_name")

    @workspace_name.setter
    def workspace_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "workspace_name", value)

    @property
    @pulumi.getter(name="deploymentName")
    def deployment_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The identifier for the Batch inference deployment.
        """
        return pulumi.get(self, "deployment_name")

    @deployment_name.setter
    def deployment_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "deployment_name", value)

    @property
    @pulumi.getter
    def identity(self) -> Optional[pulumi.Input['ManagedServiceIdentityArgs']]:
        """
        Managed service identity (system assigned and/or user assigned identities)
        """
        return pulumi.get(self, "identity")

    @identity.setter
    def identity(self, value: Optional[pulumi.Input['ManagedServiceIdentityArgs']]):
        pulumi.set(self, "identity", value)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type.
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "kind", value)

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
    def sku(self) -> Optional[pulumi.Input['SkuArgs']]:
        """
        Sku details required for ARM contract for Autoscaling.
        """
        return pulumi.get(self, "sku")

    @sku.setter
    def sku(self, value: Optional[pulumi.Input['SkuArgs']]):
        pulumi.set(self, "sku", value)

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


@pulumi.type_token("azure-native:machinelearningservices:BatchDeployment")
class BatchDeployment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 batch_deployment_properties: Optional[pulumi.Input[Union['BatchDeploymentArgs', 'BatchDeploymentArgsDict']]] = None,
                 deployment_name: Optional[pulumi.Input[builtins.str]] = None,
                 endpoint_name: Optional[pulumi.Input[builtins.str]] = None,
                 identity: Optional[pulumi.Input[Union['ManagedServiceIdentityArgs', 'ManagedServiceIdentityArgsDict']]] = None,
                 kind: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 sku: Optional[pulumi.Input[Union['SkuArgs', 'SkuArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 workspace_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Uses Azure REST API version 2024-10-01. In version 2.x of the Azure Native provider, it used API version 2023-04-01.

        Other available API versions: 2021-03-01-preview, 2022-02-01-preview, 2022-05-01, 2022-06-01-preview, 2022-10-01, 2022-10-01-preview, 2022-12-01-preview, 2023-02-01-preview, 2023-04-01, 2023-04-01-preview, 2023-06-01-preview, 2023-08-01-preview, 2023-10-01, 2024-01-01-preview, 2024-04-01, 2024-07-01-preview, 2024-10-01-preview, 2025-01-01-preview, 2025-04-01, 2025-04-01-preview, 2025-06-01, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native machinelearningservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['BatchDeploymentArgs', 'BatchDeploymentArgsDict']] batch_deployment_properties: [Required] Additional attributes of the entity.
        :param pulumi.Input[builtins.str] deployment_name: The identifier for the Batch inference deployment.
        :param pulumi.Input[builtins.str] endpoint_name: Inference endpoint name
        :param pulumi.Input[Union['ManagedServiceIdentityArgs', 'ManagedServiceIdentityArgsDict']] identity: Managed service identity (system assigned and/or user assigned identities)
        :param pulumi.Input[builtins.str] kind: Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type.
        :param pulumi.Input[builtins.str] location: The geo-location where the resource lives
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[Union['SkuArgs', 'SkuArgsDict']] sku: Sku details required for ARM contract for Autoscaling.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        :param pulumi.Input[builtins.str] workspace_name: Name of Azure Machine Learning workspace.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BatchDeploymentInitArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Uses Azure REST API version 2024-10-01. In version 2.x of the Azure Native provider, it used API version 2023-04-01.

        Other available API versions: 2021-03-01-preview, 2022-02-01-preview, 2022-05-01, 2022-06-01-preview, 2022-10-01, 2022-10-01-preview, 2022-12-01-preview, 2023-02-01-preview, 2023-04-01, 2023-04-01-preview, 2023-06-01-preview, 2023-08-01-preview, 2023-10-01, 2024-01-01-preview, 2024-04-01, 2024-07-01-preview, 2024-10-01-preview, 2025-01-01-preview, 2025-04-01, 2025-04-01-preview, 2025-06-01, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native machinelearningservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param BatchDeploymentInitArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BatchDeploymentInitArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 batch_deployment_properties: Optional[pulumi.Input[Union['BatchDeploymentArgs', 'BatchDeploymentArgsDict']]] = None,
                 deployment_name: Optional[pulumi.Input[builtins.str]] = None,
                 endpoint_name: Optional[pulumi.Input[builtins.str]] = None,
                 identity: Optional[pulumi.Input[Union['ManagedServiceIdentityArgs', 'ManagedServiceIdentityArgsDict']]] = None,
                 kind: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 sku: Optional[pulumi.Input[Union['SkuArgs', 'SkuArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 workspace_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BatchDeploymentInitArgs.__new__(BatchDeploymentInitArgs)

            if batch_deployment_properties is None and not opts.urn:
                raise TypeError("Missing required property 'batch_deployment_properties'")
            __props__.__dict__["batch_deployment_properties"] = batch_deployment_properties
            __props__.__dict__["deployment_name"] = deployment_name
            if endpoint_name is None and not opts.urn:
                raise TypeError("Missing required property 'endpoint_name'")
            __props__.__dict__["endpoint_name"] = endpoint_name
            __props__.__dict__["identity"] = identity
            __props__.__dict__["kind"] = kind
            __props__.__dict__["location"] = location
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["sku"] = sku
            __props__.__dict__["tags"] = tags
            if workspace_name is None and not opts.urn:
                raise TypeError("Missing required property 'workspace_name'")
            __props__.__dict__["workspace_name"] = workspace_name
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:machinelearningservices/v20210301preview:BatchDeployment"), pulumi.Alias(type_="azure-native:machinelearningservices/v20220201preview:BatchDeployment"), pulumi.Alias(type_="azure-native:machinelearningservices/v20220501:BatchDeployment"), pulumi.Alias(type_="azure-native:machinelearningservices/v20220601preview:BatchDeployment"), pulumi.Alias(type_="azure-native:machinelearningservices/v20221001:BatchDeployment"), pulumi.Alias(type_="azure-native:machinelearningservices/v20221001preview:BatchDeployment"), pulumi.Alias(type_="azure-native:machinelearningservices/v20221201preview:BatchDeployment"), pulumi.Alias(type_="azure-native:machinelearningservices/v20230201preview:BatchDeployment"), pulumi.Alias(type_="azure-native:machinelearningservices/v20230401:BatchDeployment"), pulumi.Alias(type_="azure-native:machinelearningservices/v20230401preview:BatchDeployment"), pulumi.Alias(type_="azure-native:machinelearningservices/v20230601preview:BatchDeployment"), pulumi.Alias(type_="azure-native:machinelearningservices/v20230801preview:BatchDeployment"), pulumi.Alias(type_="azure-native:machinelearningservices/v20231001:BatchDeployment"), pulumi.Alias(type_="azure-native:machinelearningservices/v20240101preview:BatchDeployment"), pulumi.Alias(type_="azure-native:machinelearningservices/v20240401:BatchDeployment"), pulumi.Alias(type_="azure-native:machinelearningservices/v20240401preview:BatchDeployment"), pulumi.Alias(type_="azure-native:machinelearningservices/v20240701preview:BatchDeployment"), pulumi.Alias(type_="azure-native:machinelearningservices/v20241001:BatchDeployment"), pulumi.Alias(type_="azure-native:machinelearningservices/v20241001preview:BatchDeployment"), pulumi.Alias(type_="azure-native:machinelearningservices/v20250101preview:BatchDeployment"), pulumi.Alias(type_="azure-native:machinelearningservices/v20250401:BatchDeployment"), pulumi.Alias(type_="azure-native:machinelearningservices/v20250401preview:BatchDeployment"), pulumi.Alias(type_="azure-native:machinelearningservices/v20250601:BatchDeployment"), pulumi.Alias(type_="azure-native:machinelearningservices/v20250701preview:BatchDeployment")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(BatchDeployment, __self__).__init__(
            'azure-native:machinelearningservices:BatchDeployment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'BatchDeployment':
        """
        Get an existing BatchDeployment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = BatchDeploymentInitArgs.__new__(BatchDeploymentInitArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["batch_deployment_properties"] = None
        __props__.__dict__["identity"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["sku"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["type"] = None
        return BatchDeployment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="batchDeploymentProperties")
    def batch_deployment_properties(self) -> pulumi.Output['outputs.BatchDeploymentResponse']:
        """
        [Required] Additional attributes of the entity.
        """
        return pulumi.get(self, "batch_deployment_properties")

    @property
    @pulumi.getter
    def identity(self) -> pulumi.Output[Optional['outputs.ManagedServiceIdentityResponse']]:
        """
        Managed service identity (system assigned and/or user assigned identities)
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type.
        """
        return pulumi.get(self, "kind")

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
    def sku(self) -> pulumi.Output[Optional['outputs.SkuResponse']]:
        """
        Sku details required for ARM contract for Autoscaling.
        """
        return pulumi.get(self, "sku")

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

