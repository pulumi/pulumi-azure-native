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

__all__ = ['RelationshipArgs', 'Relationship']

@pulumi.input_type
class RelationshipArgs:
    def __init__(__self__, *,
                 hub_name: pulumi.Input[builtins.str],
                 profile_type: pulumi.Input[builtins.str],
                 related_profile_type: pulumi.Input[builtins.str],
                 resource_group_name: pulumi.Input[builtins.str],
                 cardinality: Optional[pulumi.Input['CardinalityTypes']] = None,
                 description: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 display_name: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 expiry_date_time_utc: Optional[pulumi.Input[builtins.str]] = None,
                 fields: Optional[pulumi.Input[Sequence[pulumi.Input['PropertyDefinitionArgs']]]] = None,
                 lookup_mappings: Optional[pulumi.Input[Sequence[pulumi.Input['RelationshipTypeMappingArgs']]]] = None,
                 relationship_name: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a Relationship resource.
        :param pulumi.Input[builtins.str] hub_name: The name of the hub.
        :param pulumi.Input[builtins.str] profile_type: Profile type.
        :param pulumi.Input[builtins.str] related_profile_type: Related profile being referenced.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group.
        :param pulumi.Input['CardinalityTypes'] cardinality: The Relationship Cardinality.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] description: Localized descriptions for the Relationship.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] display_name: Localized display name for the Relationship.
        :param pulumi.Input[builtins.str] expiry_date_time_utc: The expiry date time in UTC.
        :param pulumi.Input[Sequence[pulumi.Input['PropertyDefinitionArgs']]] fields: The properties of the Relationship.
        :param pulumi.Input[Sequence[pulumi.Input['RelationshipTypeMappingArgs']]] lookup_mappings: Optional property to be used to map fields in profile to their strong ids in related profile.
        :param pulumi.Input[builtins.str] relationship_name: The name of the Relationship.
        """
        pulumi.set(__self__, "hub_name", hub_name)
        pulumi.set(__self__, "profile_type", profile_type)
        pulumi.set(__self__, "related_profile_type", related_profile_type)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if cardinality is not None:
            pulumi.set(__self__, "cardinality", cardinality)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if expiry_date_time_utc is not None:
            pulumi.set(__self__, "expiry_date_time_utc", expiry_date_time_utc)
        if fields is not None:
            pulumi.set(__self__, "fields", fields)
        if lookup_mappings is not None:
            pulumi.set(__self__, "lookup_mappings", lookup_mappings)
        if relationship_name is not None:
            pulumi.set(__self__, "relationship_name", relationship_name)

    @property
    @pulumi.getter(name="hubName")
    def hub_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the hub.
        """
        return pulumi.get(self, "hub_name")

    @hub_name.setter
    def hub_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "hub_name", value)

    @property
    @pulumi.getter(name="profileType")
    def profile_type(self) -> pulumi.Input[builtins.str]:
        """
        Profile type.
        """
        return pulumi.get(self, "profile_type")

    @profile_type.setter
    def profile_type(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "profile_type", value)

    @property
    @pulumi.getter(name="relatedProfileType")
    def related_profile_type(self) -> pulumi.Input[builtins.str]:
        """
        Related profile being referenced.
        """
        return pulumi.get(self, "related_profile_type")

    @related_profile_type.setter
    def related_profile_type(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "related_profile_type", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the resource group.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def cardinality(self) -> Optional[pulumi.Input['CardinalityTypes']]:
        """
        The Relationship Cardinality.
        """
        return pulumi.get(self, "cardinality")

    @cardinality.setter
    def cardinality(self, value: Optional[pulumi.Input['CardinalityTypes']]):
        pulumi.set(self, "cardinality", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        Localized descriptions for the Relationship.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        Localized display name for the Relationship.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="expiryDateTimeUtc")
    def expiry_date_time_utc(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The expiry date time in UTC.
        """
        return pulumi.get(self, "expiry_date_time_utc")

    @expiry_date_time_utc.setter
    def expiry_date_time_utc(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "expiry_date_time_utc", value)

    @property
    @pulumi.getter
    def fields(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PropertyDefinitionArgs']]]]:
        """
        The properties of the Relationship.
        """
        return pulumi.get(self, "fields")

    @fields.setter
    def fields(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PropertyDefinitionArgs']]]]):
        pulumi.set(self, "fields", value)

    @property
    @pulumi.getter(name="lookupMappings")
    def lookup_mappings(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RelationshipTypeMappingArgs']]]]:
        """
        Optional property to be used to map fields in profile to their strong ids in related profile.
        """
        return pulumi.get(self, "lookup_mappings")

    @lookup_mappings.setter
    def lookup_mappings(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RelationshipTypeMappingArgs']]]]):
        pulumi.set(self, "lookup_mappings", value)

    @property
    @pulumi.getter(name="relationshipName")
    def relationship_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the Relationship.
        """
        return pulumi.get(self, "relationship_name")

    @relationship_name.setter
    def relationship_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "relationship_name", value)


@pulumi.type_token("azure-native:customerinsights:Relationship")
class Relationship(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cardinality: Optional[pulumi.Input['CardinalityTypes']] = None,
                 description: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 display_name: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 expiry_date_time_utc: Optional[pulumi.Input[builtins.str]] = None,
                 fields: Optional[pulumi.Input[Sequence[pulumi.Input[Union['PropertyDefinitionArgs', 'PropertyDefinitionArgsDict']]]]] = None,
                 hub_name: Optional[pulumi.Input[builtins.str]] = None,
                 lookup_mappings: Optional[pulumi.Input[Sequence[pulumi.Input[Union['RelationshipTypeMappingArgs', 'RelationshipTypeMappingArgsDict']]]]] = None,
                 profile_type: Optional[pulumi.Input[builtins.str]] = None,
                 related_profile_type: Optional[pulumi.Input[builtins.str]] = None,
                 relationship_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        The relationship resource format.

        Uses Azure REST API version 2017-04-26. In version 2.x of the Azure Native provider, it used API version 2017-04-26.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input['CardinalityTypes'] cardinality: The Relationship Cardinality.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] description: Localized descriptions for the Relationship.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] display_name: Localized display name for the Relationship.
        :param pulumi.Input[builtins.str] expiry_date_time_utc: The expiry date time in UTC.
        :param pulumi.Input[Sequence[pulumi.Input[Union['PropertyDefinitionArgs', 'PropertyDefinitionArgsDict']]]] fields: The properties of the Relationship.
        :param pulumi.Input[builtins.str] hub_name: The name of the hub.
        :param pulumi.Input[Sequence[pulumi.Input[Union['RelationshipTypeMappingArgs', 'RelationshipTypeMappingArgsDict']]]] lookup_mappings: Optional property to be used to map fields in profile to their strong ids in related profile.
        :param pulumi.Input[builtins.str] profile_type: Profile type.
        :param pulumi.Input[builtins.str] related_profile_type: Related profile being referenced.
        :param pulumi.Input[builtins.str] relationship_name: The name of the Relationship.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RelationshipArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The relationship resource format.

        Uses Azure REST API version 2017-04-26. In version 2.x of the Azure Native provider, it used API version 2017-04-26.

        :param str resource_name: The name of the resource.
        :param RelationshipArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RelationshipArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cardinality: Optional[pulumi.Input['CardinalityTypes']] = None,
                 description: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 display_name: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 expiry_date_time_utc: Optional[pulumi.Input[builtins.str]] = None,
                 fields: Optional[pulumi.Input[Sequence[pulumi.Input[Union['PropertyDefinitionArgs', 'PropertyDefinitionArgsDict']]]]] = None,
                 hub_name: Optional[pulumi.Input[builtins.str]] = None,
                 lookup_mappings: Optional[pulumi.Input[Sequence[pulumi.Input[Union['RelationshipTypeMappingArgs', 'RelationshipTypeMappingArgsDict']]]]] = None,
                 profile_type: Optional[pulumi.Input[builtins.str]] = None,
                 related_profile_type: Optional[pulumi.Input[builtins.str]] = None,
                 relationship_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RelationshipArgs.__new__(RelationshipArgs)

            __props__.__dict__["cardinality"] = cardinality
            __props__.__dict__["description"] = description
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["expiry_date_time_utc"] = expiry_date_time_utc
            __props__.__dict__["fields"] = fields
            if hub_name is None and not opts.urn:
                raise TypeError("Missing required property 'hub_name'")
            __props__.__dict__["hub_name"] = hub_name
            __props__.__dict__["lookup_mappings"] = lookup_mappings
            if profile_type is None and not opts.urn:
                raise TypeError("Missing required property 'profile_type'")
            __props__.__dict__["profile_type"] = profile_type
            if related_profile_type is None and not opts.urn:
                raise TypeError("Missing required property 'related_profile_type'")
            __props__.__dict__["related_profile_type"] = related_profile_type
            __props__.__dict__["relationship_name"] = relationship_name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["provisioning_state"] = None
            __props__.__dict__["relationship_guid_id"] = None
            __props__.__dict__["tenant_id"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:customerinsights/v20170101:Relationship"), pulumi.Alias(type_="azure-native:customerinsights/v20170426:Relationship")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Relationship, __self__).__init__(
            'azure-native:customerinsights:Relationship',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Relationship':
        """
        Get an existing Relationship resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = RelationshipArgs.__new__(RelationshipArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["cardinality"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["expiry_date_time_utc"] = None
        __props__.__dict__["fields"] = None
        __props__.__dict__["lookup_mappings"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["profile_type"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["related_profile_type"] = None
        __props__.__dict__["relationship_guid_id"] = None
        __props__.__dict__["relationship_name"] = None
        __props__.__dict__["tenant_id"] = None
        __props__.__dict__["type"] = None
        return Relationship(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter
    def cardinality(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The Relationship Cardinality.
        """
        return pulumi.get(self, "cardinality")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[Mapping[str, builtins.str]]]:
        """
        Localized descriptions for the Relationship.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[Mapping[str, builtins.str]]]:
        """
        Localized display name for the Relationship.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="expiryDateTimeUtc")
    def expiry_date_time_utc(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The expiry date time in UTC.
        """
        return pulumi.get(self, "expiry_date_time_utc")

    @property
    @pulumi.getter
    def fields(self) -> pulumi.Output[Optional[Sequence['outputs.PropertyDefinitionResponse']]]:
        """
        The properties of the Relationship.
        """
        return pulumi.get(self, "fields")

    @property
    @pulumi.getter(name="lookupMappings")
    def lookup_mappings(self) -> pulumi.Output[Optional[Sequence['outputs.RelationshipTypeMappingResponse']]]:
        """
        Optional property to be used to map fields in profile to their strong ids in related profile.
        """
        return pulumi.get(self, "lookup_mappings")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="profileType")
    def profile_type(self) -> pulumi.Output[builtins.str]:
        """
        Profile type.
        """
        return pulumi.get(self, "profile_type")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[builtins.str]:
        """
        Provisioning state.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="relatedProfileType")
    def related_profile_type(self) -> pulumi.Output[builtins.str]:
        """
        Related profile being referenced.
        """
        return pulumi.get(self, "related_profile_type")

    @property
    @pulumi.getter(name="relationshipGuidId")
    def relationship_guid_id(self) -> pulumi.Output[builtins.str]:
        """
        The relationship guid id.
        """
        return pulumi.get(self, "relationship_guid_id")

    @property
    @pulumi.getter(name="relationshipName")
    def relationship_name(self) -> pulumi.Output[builtins.str]:
        """
        The Relationship name.
        """
        return pulumi.get(self, "relationship_name")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[builtins.str]:
        """
        The hub name.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

