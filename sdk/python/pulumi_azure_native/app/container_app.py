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

__all__ = ['ContainerAppArgs', 'ContainerApp']

@pulumi.input_type
class ContainerAppArgs:
    def __init__(__self__, *,
                 resource_group_name: pulumi.Input[builtins.str],
                 configuration: Optional[pulumi.Input['ConfigurationArgs']] = None,
                 container_app_name: Optional[pulumi.Input[builtins.str]] = None,
                 environment_id: Optional[pulumi.Input[builtins.str]] = None,
                 extended_location: Optional[pulumi.Input['ExtendedLocationArgs']] = None,
                 identity: Optional[pulumi.Input['ManagedServiceIdentityArgs']] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 managed_by: Optional[pulumi.Input[builtins.str]] = None,
                 managed_environment_id: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 template: Optional[pulumi.Input['TemplateArgs']] = None,
                 workload_profile_name: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a ContainerApp resource.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input['ConfigurationArgs'] configuration: Non versioned Container App configuration properties.
        :param pulumi.Input[builtins.str] container_app_name: Name of the Container App.
        :param pulumi.Input[builtins.str] environment_id: Resource ID of environment.
        :param pulumi.Input['ExtendedLocationArgs'] extended_location: The complex type of the extended location.
        :param pulumi.Input['ManagedServiceIdentityArgs'] identity: managed identities for the Container App to interact with other Azure services without maintaining any secrets or credentials in code.
        :param pulumi.Input[builtins.str] location: The geo-location where the resource lives
        :param pulumi.Input[builtins.str] managed_by: The fully qualified resource ID of the resource that manages this resource. Indicates if this resource is managed by another Azure resource. If this is present, complete mode deployment will not delete the resource if it is removed from the template since it is managed by another resource.
        :param pulumi.Input[builtins.str] managed_environment_id: Deprecated. Resource ID of the Container App's environment.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        :param pulumi.Input['TemplateArgs'] template: Container App versioned application definition.
        :param pulumi.Input[builtins.str] workload_profile_name: Workload profile name to pin for container app execution.
        """
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if configuration is not None:
            pulumi.set(__self__, "configuration", configuration)
        if container_app_name is not None:
            pulumi.set(__self__, "container_app_name", container_app_name)
        if environment_id is not None:
            pulumi.set(__self__, "environment_id", environment_id)
        if extended_location is not None:
            pulumi.set(__self__, "extended_location", extended_location)
        if identity is not None:
            pulumi.set(__self__, "identity", identity)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if managed_by is not None:
            pulumi.set(__self__, "managed_by", managed_by)
        if managed_environment_id is not None:
            pulumi.set(__self__, "managed_environment_id", managed_environment_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if template is not None:
            pulumi.set(__self__, "template", template)
        if workload_profile_name is not None:
            pulumi.set(__self__, "workload_profile_name", workload_profile_name)

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
    @pulumi.getter
    def configuration(self) -> Optional[pulumi.Input['ConfigurationArgs']]:
        """
        Non versioned Container App configuration properties.
        """
        return pulumi.get(self, "configuration")

    @configuration.setter
    def configuration(self, value: Optional[pulumi.Input['ConfigurationArgs']]):
        pulumi.set(self, "configuration", value)

    @property
    @pulumi.getter(name="containerAppName")
    def container_app_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of the Container App.
        """
        return pulumi.get(self, "container_app_name")

    @container_app_name.setter
    def container_app_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "container_app_name", value)

    @property
    @pulumi.getter(name="environmentId")
    def environment_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Resource ID of environment.
        """
        return pulumi.get(self, "environment_id")

    @environment_id.setter
    def environment_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "environment_id", value)

    @property
    @pulumi.getter(name="extendedLocation")
    def extended_location(self) -> Optional[pulumi.Input['ExtendedLocationArgs']]:
        """
        The complex type of the extended location.
        """
        return pulumi.get(self, "extended_location")

    @extended_location.setter
    def extended_location(self, value: Optional[pulumi.Input['ExtendedLocationArgs']]):
        pulumi.set(self, "extended_location", value)

    @property
    @pulumi.getter
    def identity(self) -> Optional[pulumi.Input['ManagedServiceIdentityArgs']]:
        """
        managed identities for the Container App to interact with other Azure services without maintaining any secrets or credentials in code.
        """
        return pulumi.get(self, "identity")

    @identity.setter
    def identity(self, value: Optional[pulumi.Input['ManagedServiceIdentityArgs']]):
        pulumi.set(self, "identity", value)

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
    @pulumi.getter(name="managedBy")
    def managed_by(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The fully qualified resource ID of the resource that manages this resource. Indicates if this resource is managed by another Azure resource. If this is present, complete mode deployment will not delete the resource if it is removed from the template since it is managed by another resource.
        """
        return pulumi.get(self, "managed_by")

    @managed_by.setter
    def managed_by(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "managed_by", value)

    @property
    @pulumi.getter(name="managedEnvironmentId")
    def managed_environment_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Deprecated. Resource ID of the Container App's environment.
        """
        return pulumi.get(self, "managed_environment_id")

    @managed_environment_id.setter
    def managed_environment_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "managed_environment_id", value)

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

    @property
    @pulumi.getter
    def template(self) -> Optional[pulumi.Input['TemplateArgs']]:
        """
        Container App versioned application definition.
        """
        return pulumi.get(self, "template")

    @template.setter
    def template(self, value: Optional[pulumi.Input['TemplateArgs']]):
        pulumi.set(self, "template", value)

    @property
    @pulumi.getter(name="workloadProfileName")
    def workload_profile_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Workload profile name to pin for container app execution.
        """
        return pulumi.get(self, "workload_profile_name")

    @workload_profile_name.setter
    def workload_profile_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "workload_profile_name", value)


@pulumi.type_token("azure-native:app:ContainerApp")
class ContainerApp(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 configuration: Optional[pulumi.Input[Union['ConfigurationArgs', 'ConfigurationArgsDict']]] = None,
                 container_app_name: Optional[pulumi.Input[builtins.str]] = None,
                 environment_id: Optional[pulumi.Input[builtins.str]] = None,
                 extended_location: Optional[pulumi.Input[Union['ExtendedLocationArgs', 'ExtendedLocationArgsDict']]] = None,
                 identity: Optional[pulumi.Input[Union['ManagedServiceIdentityArgs', 'ManagedServiceIdentityArgsDict']]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 managed_by: Optional[pulumi.Input[builtins.str]] = None,
                 managed_environment_id: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 template: Optional[pulumi.Input[Union['TemplateArgs', 'TemplateArgsDict']]] = None,
                 workload_profile_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Container App.

        Uses Azure REST API version 2024-03-01. In version 2.x of the Azure Native provider, it used API version 2022-10-01.

        Other available API versions: 2022-10-01, 2022-11-01-preview, 2023-04-01-preview, 2023-05-01, 2023-05-02-preview, 2023-08-01-preview, 2023-11-02-preview, 2024-02-02-preview, 2024-08-02-preview, 2024-10-02-preview, 2025-01-01, 2025-02-02-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native app [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['ConfigurationArgs', 'ConfigurationArgsDict']] configuration: Non versioned Container App configuration properties.
        :param pulumi.Input[builtins.str] container_app_name: Name of the Container App.
        :param pulumi.Input[builtins.str] environment_id: Resource ID of environment.
        :param pulumi.Input[Union['ExtendedLocationArgs', 'ExtendedLocationArgsDict']] extended_location: The complex type of the extended location.
        :param pulumi.Input[Union['ManagedServiceIdentityArgs', 'ManagedServiceIdentityArgsDict']] identity: managed identities for the Container App to interact with other Azure services without maintaining any secrets or credentials in code.
        :param pulumi.Input[builtins.str] location: The geo-location where the resource lives
        :param pulumi.Input[builtins.str] managed_by: The fully qualified resource ID of the resource that manages this resource. Indicates if this resource is managed by another Azure resource. If this is present, complete mode deployment will not delete the resource if it is removed from the template since it is managed by another resource.
        :param pulumi.Input[builtins.str] managed_environment_id: Deprecated. Resource ID of the Container App's environment.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        :param pulumi.Input[Union['TemplateArgs', 'TemplateArgsDict']] template: Container App versioned application definition.
        :param pulumi.Input[builtins.str] workload_profile_name: Workload profile name to pin for container app execution.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ContainerAppArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Container App.

        Uses Azure REST API version 2024-03-01. In version 2.x of the Azure Native provider, it used API version 2022-10-01.

        Other available API versions: 2022-10-01, 2022-11-01-preview, 2023-04-01-preview, 2023-05-01, 2023-05-02-preview, 2023-08-01-preview, 2023-11-02-preview, 2024-02-02-preview, 2024-08-02-preview, 2024-10-02-preview, 2025-01-01, 2025-02-02-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native app [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param ContainerAppArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ContainerAppArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 configuration: Optional[pulumi.Input[Union['ConfigurationArgs', 'ConfigurationArgsDict']]] = None,
                 container_app_name: Optional[pulumi.Input[builtins.str]] = None,
                 environment_id: Optional[pulumi.Input[builtins.str]] = None,
                 extended_location: Optional[pulumi.Input[Union['ExtendedLocationArgs', 'ExtendedLocationArgsDict']]] = None,
                 identity: Optional[pulumi.Input[Union['ManagedServiceIdentityArgs', 'ManagedServiceIdentityArgsDict']]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 managed_by: Optional[pulumi.Input[builtins.str]] = None,
                 managed_environment_id: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 template: Optional[pulumi.Input[Union['TemplateArgs', 'TemplateArgsDict']]] = None,
                 workload_profile_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ContainerAppArgs.__new__(ContainerAppArgs)

            __props__.__dict__["configuration"] = configuration
            __props__.__dict__["container_app_name"] = container_app_name
            __props__.__dict__["environment_id"] = environment_id
            __props__.__dict__["extended_location"] = extended_location
            __props__.__dict__["identity"] = identity
            __props__.__dict__["location"] = location
            __props__.__dict__["managed_by"] = managed_by
            __props__.__dict__["managed_environment_id"] = managed_environment_id
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["template"] = template
            __props__.__dict__["workload_profile_name"] = workload_profile_name
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["custom_domain_verification_id"] = None
            __props__.__dict__["event_stream_endpoint"] = None
            __props__.__dict__["latest_ready_revision_name"] = None
            __props__.__dict__["latest_revision_fqdn"] = None
            __props__.__dict__["latest_revision_name"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["outbound_ip_addresses"] = None
            __props__.__dict__["provisioning_state"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:app/v20220101preview:ContainerApp"), pulumi.Alias(type_="azure-native:app/v20220301:ContainerApp"), pulumi.Alias(type_="azure-native:app/v20220601preview:ContainerApp"), pulumi.Alias(type_="azure-native:app/v20221001:ContainerApp"), pulumi.Alias(type_="azure-native:app/v20221101preview:ContainerApp"), pulumi.Alias(type_="azure-native:app/v20230401preview:ContainerApp"), pulumi.Alias(type_="azure-native:app/v20230501:ContainerApp"), pulumi.Alias(type_="azure-native:app/v20230502preview:ContainerApp"), pulumi.Alias(type_="azure-native:app/v20230801preview:ContainerApp"), pulumi.Alias(type_="azure-native:app/v20231102preview:ContainerApp"), pulumi.Alias(type_="azure-native:app/v20240202preview:ContainerApp"), pulumi.Alias(type_="azure-native:app/v20240301:ContainerApp"), pulumi.Alias(type_="azure-native:app/v20240802preview:ContainerApp"), pulumi.Alias(type_="azure-native:app/v20241002preview:ContainerApp"), pulumi.Alias(type_="azure-native:app/v20250101:ContainerApp"), pulumi.Alias(type_="azure-native:app/v20250202preview:ContainerApp")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(ContainerApp, __self__).__init__(
            'azure-native:app:ContainerApp',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ContainerApp':
        """
        Get an existing ContainerApp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ContainerAppArgs.__new__(ContainerAppArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["configuration"] = None
        __props__.__dict__["custom_domain_verification_id"] = None
        __props__.__dict__["environment_id"] = None
        __props__.__dict__["event_stream_endpoint"] = None
        __props__.__dict__["extended_location"] = None
        __props__.__dict__["identity"] = None
        __props__.__dict__["latest_ready_revision_name"] = None
        __props__.__dict__["latest_revision_fqdn"] = None
        __props__.__dict__["latest_revision_name"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["managed_by"] = None
        __props__.__dict__["managed_environment_id"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["outbound_ip_addresses"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["template"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["workload_profile_name"] = None
        return ContainerApp(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter
    def configuration(self) -> pulumi.Output[Optional['outputs.ConfigurationResponse']]:
        """
        Non versioned Container App configuration properties.
        """
        return pulumi.get(self, "configuration")

    @property
    @pulumi.getter(name="customDomainVerificationId")
    def custom_domain_verification_id(self) -> pulumi.Output[builtins.str]:
        """
        Id used to verify domain name ownership
        """
        return pulumi.get(self, "custom_domain_verification_id")

    @property
    @pulumi.getter(name="environmentId")
    def environment_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Resource ID of environment.
        """
        return pulumi.get(self, "environment_id")

    @property
    @pulumi.getter(name="eventStreamEndpoint")
    def event_stream_endpoint(self) -> pulumi.Output[builtins.str]:
        """
        The endpoint of the eventstream of the container app.
        """
        return pulumi.get(self, "event_stream_endpoint")

    @property
    @pulumi.getter(name="extendedLocation")
    def extended_location(self) -> pulumi.Output[Optional['outputs.ExtendedLocationResponse']]:
        """
        The complex type of the extended location.
        """
        return pulumi.get(self, "extended_location")

    @property
    @pulumi.getter
    def identity(self) -> pulumi.Output[Optional['outputs.ManagedServiceIdentityResponse']]:
        """
        managed identities for the Container App to interact with other Azure services without maintaining any secrets or credentials in code.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter(name="latestReadyRevisionName")
    def latest_ready_revision_name(self) -> pulumi.Output[builtins.str]:
        """
        Name of the latest ready revision of the Container App.
        """
        return pulumi.get(self, "latest_ready_revision_name")

    @property
    @pulumi.getter(name="latestRevisionFqdn")
    def latest_revision_fqdn(self) -> pulumi.Output[builtins.str]:
        """
        Fully Qualified Domain Name of the latest revision of the Container App.
        """
        return pulumi.get(self, "latest_revision_fqdn")

    @property
    @pulumi.getter(name="latestRevisionName")
    def latest_revision_name(self) -> pulumi.Output[builtins.str]:
        """
        Name of the latest revision of the Container App.
        """
        return pulumi.get(self, "latest_revision_name")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[builtins.str]:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="managedBy")
    def managed_by(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The fully qualified resource ID of the resource that manages this resource. Indicates if this resource is managed by another Azure resource. If this is present, complete mode deployment will not delete the resource if it is removed from the template since it is managed by another resource.
        """
        return pulumi.get(self, "managed_by")

    @property
    @pulumi.getter(name="managedEnvironmentId")
    def managed_environment_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Deprecated. Resource ID of the Container App's environment.
        """
        return pulumi.get(self, "managed_environment_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="outboundIpAddresses")
    def outbound_ip_addresses(self) -> pulumi.Output[Sequence[builtins.str]]:
        """
        Outbound IP Addresses for container app.
        """
        return pulumi.get(self, "outbound_ip_addresses")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[builtins.str]:
        """
        Provisioning state of the Container App.
        """
        return pulumi.get(self, "provisioning_state")

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
    def template(self) -> pulumi.Output[Optional['outputs.TemplateResponse']]:
        """
        Container App versioned application definition.
        """
        return pulumi.get(self, "template")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="workloadProfileName")
    def workload_profile_name(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Workload profile name to pin for container app execution.
        """
        return pulumi.get(self, "workload_profile_name")

