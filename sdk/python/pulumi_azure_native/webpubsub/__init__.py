# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from .. import _utilities
import typing
# Export this package's modules as members:
from ._enums import *
from .get_web_pub_sub import *
from .get_web_pub_sub_custom_certificate import *
from .get_web_pub_sub_custom_domain import *
from .get_web_pub_sub_hub import *
from .get_web_pub_sub_private_endpoint_connection import *
from .get_web_pub_sub_replica import *
from .get_web_pub_sub_shared_private_link_resource import *
from .list_web_pub_sub_keys import *
from .web_pub_sub import *
from .web_pub_sub_custom_certificate import *
from .web_pub_sub_custom_domain import *
from .web_pub_sub_hub import *
from .web_pub_sub_private_endpoint_connection import *
from .web_pub_sub_replica import *
from .web_pub_sub_shared_private_link_resource import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_azure_native.webpubsub.v20210401preview as __v20210401preview
    v20210401preview = __v20210401preview
    import pulumi_azure_native.webpubsub.v20210601preview as __v20210601preview
    v20210601preview = __v20210601preview
    import pulumi_azure_native.webpubsub.v20210901preview as __v20210901preview
    v20210901preview = __v20210901preview
    import pulumi_azure_native.webpubsub.v20230201 as __v20230201
    v20230201 = __v20230201
    import pulumi_azure_native.webpubsub.v20230301preview as __v20230301preview
    v20230301preview = __v20230301preview
    import pulumi_azure_native.webpubsub.v20230601preview as __v20230601preview
    v20230601preview = __v20230601preview
    import pulumi_azure_native.webpubsub.v20230801preview as __v20230801preview
    v20230801preview = __v20230801preview
    import pulumi_azure_native.webpubsub.v20240101preview as __v20240101preview
    v20240101preview = __v20240101preview
    import pulumi_azure_native.webpubsub.v20240301 as __v20240301
    v20240301 = __v20240301
    import pulumi_azure_native.webpubsub.v20240401preview as __v20240401preview
    v20240401preview = __v20240401preview
    import pulumi_azure_native.webpubsub.v20240801preview as __v20240801preview
    v20240801preview = __v20240801preview
    import pulumi_azure_native.webpubsub.v20241001preview as __v20241001preview
    v20241001preview = __v20241001preview
else:
    v20210401preview = _utilities.lazy_import('pulumi_azure_native.webpubsub.v20210401preview')
    v20210601preview = _utilities.lazy_import('pulumi_azure_native.webpubsub.v20210601preview')
    v20210901preview = _utilities.lazy_import('pulumi_azure_native.webpubsub.v20210901preview')
    v20230201 = _utilities.lazy_import('pulumi_azure_native.webpubsub.v20230201')
    v20230301preview = _utilities.lazy_import('pulumi_azure_native.webpubsub.v20230301preview')
    v20230601preview = _utilities.lazy_import('pulumi_azure_native.webpubsub.v20230601preview')
    v20230801preview = _utilities.lazy_import('pulumi_azure_native.webpubsub.v20230801preview')
    v20240101preview = _utilities.lazy_import('pulumi_azure_native.webpubsub.v20240101preview')
    v20240301 = _utilities.lazy_import('pulumi_azure_native.webpubsub.v20240301')
    v20240401preview = _utilities.lazy_import('pulumi_azure_native.webpubsub.v20240401preview')
    v20240801preview = _utilities.lazy_import('pulumi_azure_native.webpubsub.v20240801preview')
    v20241001preview = _utilities.lazy_import('pulumi_azure_native.webpubsub.v20241001preview')

