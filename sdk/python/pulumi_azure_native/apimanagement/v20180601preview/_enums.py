# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'ApiType',
    'BearerTokenSendingMethods',
    'ContentFormat',
    'PolicyContentFormat',
    'Protocol',
    'SoapApiType',
]


class ApiType(str, Enum):
    """
    Type of API.
    """
    HTTP = "http"
    SOAP = "soap"


class BearerTokenSendingMethods(str, Enum):
    """
    Form of an authorization grant, which the client uses to request the access token.
    """
    AUTHORIZATION_HEADER = "authorizationHeader"
    """
    Access token will be transmitted in the Authorization header using Bearer schema
    """
    QUERY = "query"
    """
    Access token will be transmitted as query parameters.
    """


class ContentFormat(str, Enum):
    """
    Format of the Content in which the API is getting imported.
    """
    WADL_XML = "wadl-xml"
    """
    The contents are inline and Content type is a WADL document.
    """
    WADL_LINK_JSON = "wadl-link-json"
    """
    The WADL document is hosted on a publicly accessible internet address.
    """
    SWAGGER_JSON = "swagger-json"
    """
    The contents are inline and Content Type is a OpenApi 2.0 Document.
    """
    SWAGGER_LINK_JSON = "swagger-link-json"
    """
    The Open Api 2.0 document is hosted on a publicly accessible internet address.
    """
    WSDL = "wsdl"
    """
    The contents are inline and the document is a WSDL/Soap document.
    """
    WSDL_LINK = "wsdl-link"
    """
    The WSDL document is hosted on a publicly accessible internet address.
    """
    OPENAPI = "openapi"
    """
    The contents are inline and Content Type is a OpenApi 3.0 Document in YAML format.
    """
    OPENAPI_JSON = "openapi+json"
    """
    The contents are inline and Content Type is a OpenApi 3.0 Document in JSON format.
    """
    OPENAPI_LINK = "openapi-link"
    """
    The Open Api 3.0 document is hosted on a publicly accessible internet address.
    """


class PolicyContentFormat(str, Enum):
    """
    Format of the policyContent.
    """
    XML = "xml"
    """
    The contents are inline and Content type is an XML document.
    """
    XML_LINK = "xml-link"
    """
    The policy XML document is hosted on a http endpoint accessible from the API Management service.
    """
    RAWXML = "rawxml"
    """
    The contents are inline and Content type is a non XML encoded policy document.
    """
    RAWXML_LINK = "rawxml-link"
    """
    The policy document is not Xml encoded and is hosted on a http endpoint accessible from the API Management service.
    """


class Protocol(str, Enum):
    HTTP = "http"
    HTTPS = "https"


class SoapApiType(str, Enum):
    """
    Type of Api to create. 
     * `http` creates a SOAP to REST API 
     * `soap` creates a SOAP pass-through API .
    """
    SOAP_TO_REST = "http"
    """
    Imports a SOAP API having a RESTful front end.
    """
    SOAP_PASS_THROUGH = "soap"
    """
    Imports the Soap API having a SOAP front end.
    """
