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

__all__ = ['ProfileArgs', 'Profile']

@pulumi.input_type
class ProfileArgs:
    def __init__(__self__, *,
                 resource_group_name: pulumi.Input[builtins.str],
                 sku: pulumi.Input['SkuArgs'],
                 identity: Optional[pulumi.Input['ManagedServiceIdentityArgs']] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 log_scrubbing: Optional[pulumi.Input['ProfileLogScrubbingArgs']] = None,
                 origin_response_timeout_seconds: Optional[pulumi.Input[builtins.int]] = None,
                 profile_name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None):
        """
        The set of arguments for constructing a Profile resource.
        :param pulumi.Input[builtins.str] resource_group_name: Name of the Resource group within the Azure subscription.
        :param pulumi.Input['SkuArgs'] sku: The pricing tier (defines Azure Front Door Standard or Premium or a CDN provider, feature list and rate) of the profile.
        :param pulumi.Input['ManagedServiceIdentityArgs'] identity: Managed service identity (system assigned and/or user assigned identities).
        :param pulumi.Input[builtins.str] location: Resource location.
        :param pulumi.Input['ProfileLogScrubbingArgs'] log_scrubbing: Defines rules that scrub sensitive fields in the Azure Front Door profile logs.
        :param pulumi.Input[builtins.int] origin_response_timeout_seconds: Send and receive timeout on forwarding request to the origin. When timeout is reached, the request fails and returns.
        :param pulumi.Input[builtins.str] profile_name: Name of the Azure Front Door Standard or Azure Front Door Premium or CDN profile which is unique within the resource group.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        """
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "sku", sku)
        if identity is not None:
            pulumi.set(__self__, "identity", identity)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if log_scrubbing is not None:
            pulumi.set(__self__, "log_scrubbing", log_scrubbing)
        if origin_response_timeout_seconds is not None:
            pulumi.set(__self__, "origin_response_timeout_seconds", origin_response_timeout_seconds)
        if profile_name is not None:
            pulumi.set(__self__, "profile_name", profile_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[builtins.str]:
        """
        Name of the Resource group within the Azure subscription.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def sku(self) -> pulumi.Input['SkuArgs']:
        """
        The pricing tier (defines Azure Front Door Standard or Premium or a CDN provider, feature list and rate) of the profile.
        """
        return pulumi.get(self, "sku")

    @sku.setter
    def sku(self, value: pulumi.Input['SkuArgs']):
        pulumi.set(self, "sku", value)

    @property
    @pulumi.getter
    def identity(self) -> Optional[pulumi.Input['ManagedServiceIdentityArgs']]:
        """
        Managed service identity (system assigned and/or user assigned identities).
        """
        return pulumi.get(self, "identity")

    @identity.setter
    def identity(self, value: Optional[pulumi.Input['ManagedServiceIdentityArgs']]):
        pulumi.set(self, "identity", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Resource location.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="logScrubbing")
    def log_scrubbing(self) -> Optional[pulumi.Input['ProfileLogScrubbingArgs']]:
        """
        Defines rules that scrub sensitive fields in the Azure Front Door profile logs.
        """
        return pulumi.get(self, "log_scrubbing")

    @log_scrubbing.setter
    def log_scrubbing(self, value: Optional[pulumi.Input['ProfileLogScrubbingArgs']]):
        pulumi.set(self, "log_scrubbing", value)

    @property
    @pulumi.getter(name="originResponseTimeoutSeconds")
    def origin_response_timeout_seconds(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Send and receive timeout on forwarding request to the origin. When timeout is reached, the request fails and returns.
        """
        return pulumi.get(self, "origin_response_timeout_seconds")

    @origin_response_timeout_seconds.setter
    def origin_response_timeout_seconds(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "origin_response_timeout_seconds", value)

    @property
    @pulumi.getter(name="profileName")
    def profile_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of the Azure Front Door Standard or Azure Front Door Premium or CDN profile which is unique within the resource group.
        """
        return pulumi.get(self, "profile_name")

    @profile_name.setter
    def profile_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "profile_name", value)

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


@pulumi.type_token("azure-native:cdn:Profile")
class Profile(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 identity: Optional[pulumi.Input[Union['ManagedServiceIdentityArgs', 'ManagedServiceIdentityArgsDict']]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 log_scrubbing: Optional[pulumi.Input[Union['ProfileLogScrubbingArgs', 'ProfileLogScrubbingArgsDict']]] = None,
                 origin_response_timeout_seconds: Optional[pulumi.Input[builtins.int]] = None,
                 profile_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 sku: Optional[pulumi.Input[Union['SkuArgs', 'SkuArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        """
        A profile is a logical grouping of endpoints that share the same settings.

        Uses Azure REST API version 2024-09-01. In version 2.x of the Azure Native provider, it used API version 2023-05-01.

        Other available API versions: 2023-05-01, 2023-07-01-preview, 2024-02-01, 2024-05-01-preview, 2024-06-01-preview, 2025-01-01-preview, 2025-04-15, 2025-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cdn [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['ManagedServiceIdentityArgs', 'ManagedServiceIdentityArgsDict']] identity: Managed service identity (system assigned and/or user assigned identities).
        :param pulumi.Input[builtins.str] location: Resource location.
        :param pulumi.Input[Union['ProfileLogScrubbingArgs', 'ProfileLogScrubbingArgsDict']] log_scrubbing: Defines rules that scrub sensitive fields in the Azure Front Door profile logs.
        :param pulumi.Input[builtins.int] origin_response_timeout_seconds: Send and receive timeout on forwarding request to the origin. When timeout is reached, the request fails and returns.
        :param pulumi.Input[builtins.str] profile_name: Name of the Azure Front Door Standard or Azure Front Door Premium or CDN profile which is unique within the resource group.
        :param pulumi.Input[builtins.str] resource_group_name: Name of the Resource group within the Azure subscription.
        :param pulumi.Input[Union['SkuArgs', 'SkuArgsDict']] sku: The pricing tier (defines Azure Front Door Standard or Premium or a CDN provider, feature list and rate) of the profile.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProfileArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A profile is a logical grouping of endpoints that share the same settings.

        Uses Azure REST API version 2024-09-01. In version 2.x of the Azure Native provider, it used API version 2023-05-01.

        Other available API versions: 2023-05-01, 2023-07-01-preview, 2024-02-01, 2024-05-01-preview, 2024-06-01-preview, 2025-01-01-preview, 2025-04-15, 2025-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cdn [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param ProfileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProfileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 identity: Optional[pulumi.Input[Union['ManagedServiceIdentityArgs', 'ManagedServiceIdentityArgsDict']]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 log_scrubbing: Optional[pulumi.Input[Union['ProfileLogScrubbingArgs', 'ProfileLogScrubbingArgsDict']]] = None,
                 origin_response_timeout_seconds: Optional[pulumi.Input[builtins.int]] = None,
                 profile_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 sku: Optional[pulumi.Input[Union['SkuArgs', 'SkuArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProfileArgs.__new__(ProfileArgs)

            __props__.__dict__["identity"] = identity
            __props__.__dict__["location"] = location
            __props__.__dict__["log_scrubbing"] = log_scrubbing
            __props__.__dict__["origin_response_timeout_seconds"] = origin_response_timeout_seconds
            __props__.__dict__["profile_name"] = profile_name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            if sku is None and not opts.urn:
                raise TypeError("Missing required property 'sku'")
            __props__.__dict__["sku"] = sku
            __props__.__dict__["tags"] = tags
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["extended_properties"] = None
            __props__.__dict__["front_door_id"] = None
            __props__.__dict__["kind"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["provisioning_state"] = None
            __props__.__dict__["resource_state"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:cdn/v20150601:Profile"), pulumi.Alias(type_="azure-native:cdn/v20160402:Profile"), pulumi.Alias(type_="azure-native:cdn/v20161002:Profile"), pulumi.Alias(type_="azure-native:cdn/v20170402:Profile"), pulumi.Alias(type_="azure-native:cdn/v20171012:Profile"), pulumi.Alias(type_="azure-native:cdn/v20190415:Profile"), pulumi.Alias(type_="azure-native:cdn/v20190615:Profile"), pulumi.Alias(type_="azure-native:cdn/v20190615preview:Profile"), pulumi.Alias(type_="azure-native:cdn/v20191231:Profile"), pulumi.Alias(type_="azure-native:cdn/v20200331:Profile"), pulumi.Alias(type_="azure-native:cdn/v20200415:Profile"), pulumi.Alias(type_="azure-native:cdn/v20200901:Profile"), pulumi.Alias(type_="azure-native:cdn/v20210601:Profile"), pulumi.Alias(type_="azure-native:cdn/v20220501preview:Profile"), pulumi.Alias(type_="azure-native:cdn/v20221101preview:Profile"), pulumi.Alias(type_="azure-native:cdn/v20230501:Profile"), pulumi.Alias(type_="azure-native:cdn/v20230701preview:Profile"), pulumi.Alias(type_="azure-native:cdn/v20240201:Profile"), pulumi.Alias(type_="azure-native:cdn/v20240501preview:Profile"), pulumi.Alias(type_="azure-native:cdn/v20240601preview:Profile"), pulumi.Alias(type_="azure-native:cdn/v20240901:Profile"), pulumi.Alias(type_="azure-native:cdn/v20250101preview:Profile"), pulumi.Alias(type_="azure-native:cdn/v20250415:Profile"), pulumi.Alias(type_="azure-native:cdn/v20250601:Profile")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Profile, __self__).__init__(
            'azure-native:cdn:Profile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Profile':
        """
        Get an existing Profile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ProfileArgs.__new__(ProfileArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["extended_properties"] = None
        __props__.__dict__["front_door_id"] = None
        __props__.__dict__["identity"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["log_scrubbing"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["origin_response_timeout_seconds"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["resource_state"] = None
        __props__.__dict__["sku"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["type"] = None
        return Profile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="extendedProperties")
    def extended_properties(self) -> pulumi.Output[Mapping[str, builtins.str]]:
        """
        Key-Value pair representing additional properties for profiles.
        """
        return pulumi.get(self, "extended_properties")

    @property
    @pulumi.getter(name="frontDoorId")
    def front_door_id(self) -> pulumi.Output[builtins.str]:
        """
        The Id of the frontdoor.
        """
        return pulumi.get(self, "front_door_id")

    @property
    @pulumi.getter
    def identity(self) -> pulumi.Output[Optional['outputs.ManagedServiceIdentityResponse']]:
        """
        Managed service identity (system assigned and/or user assigned identities).
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[builtins.str]:
        """
        Kind of the profile. Used by portal to differentiate traditional CDN profile and new AFD profile.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[builtins.str]:
        """
        Resource location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="logScrubbing")
    def log_scrubbing(self) -> pulumi.Output[Optional['outputs.ProfileLogScrubbingResponse']]:
        """
        Defines rules that scrub sensitive fields in the Azure Front Door profile logs.
        """
        return pulumi.get(self, "log_scrubbing")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="originResponseTimeoutSeconds")
    def origin_response_timeout_seconds(self) -> pulumi.Output[Optional[builtins.int]]:
        """
        Send and receive timeout on forwarding request to the origin. When timeout is reached, the request fails and returns.
        """
        return pulumi.get(self, "origin_response_timeout_seconds")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[builtins.str]:
        """
        Provisioning status of the profile.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="resourceState")
    def resource_state(self) -> pulumi.Output[builtins.str]:
        """
        Resource status of the profile.
        """
        return pulumi.get(self, "resource_state")

    @property
    @pulumi.getter
    def sku(self) -> pulumi.Output['outputs.SkuResponse']:
        """
        The pricing tier (defines Azure Front Door Standard or Premium or a CDN provider, feature list and rate) of the profile.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> pulumi.Output['outputs.SystemDataResponse']:
        """
        Read only system data
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
        Resource type.
        """
        return pulumi.get(self, "type")

