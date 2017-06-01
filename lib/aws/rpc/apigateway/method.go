// *** WARNING: this file was generated by the Lumi IDL Compiler (LUMIDL). ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
    "errors"

    pbempty "github.com/golang/protobuf/ptypes/empty"
    pbstruct "github.com/golang/protobuf/ptypes/struct"
    "golang.org/x/net/context"

    "github.com/pulumi/lumi/pkg/resource"
    "github.com/pulumi/lumi/pkg/tokens"
    "github.com/pulumi/lumi/pkg/util/contract"
    "github.com/pulumi/lumi/pkg/util/mapper"
    "github.com/pulumi/lumi/sdk/go/pkg/lumirpc"
)

/* Marshalable Integration structure(s) */

// Integration is a marshalable representation of its corresponding IDL type.
type Integration struct {
    Type IntegrationType `json:"type"`
    CacheKeyParameters *[]string `json:"cacheKeyParameters,omitempty"`
    CacheNamespace *string `json:"cacheNamespace,omitempty"`
    Credentials *string `json:"credentials,omitempty"`
    IntegrationHTTPMethod *string `json:"integrationHTTPMethod,omitempty"`
    IntegrationResponse *[]IntegrationResponse `json:"integrationResponse,omitempty"`
    PassthroughBehavior *PassthroughBehavior `json:"passthroughBehavior,omitempty"`
    RequestParameters *map[string]string `json:"requestParameters,omitempty"`
    RequestTemplates *map[string]string `json:"requestTemplates,omitempty"`
    URI *string `json:"uri,omitempty"`
}

// Integration's properties have constants to make dealing with diffs and property bags easier.
const (
    Integration_Type = "type"
    Integration_CacheKeyParameters = "cacheKeyParameters"
    Integration_CacheNamespace = "cacheNamespace"
    Integration_Credentials = "credentials"
    Integration_IntegrationHTTPMethod = "integrationHTTPMethod"
    Integration_IntegrationResponse = "integrationResponse"
    Integration_PassthroughBehavior = "passthroughBehavior"
    Integration_RequestParameters = "requestParameters"
    Integration_RequestTemplates = "requestTemplates"
    Integration_URI = "uri"
)

/* Marshalable IntegrationResponse structure(s) */

// IntegrationResponse is a marshalable representation of its corresponding IDL type.
type IntegrationResponse struct {
    ResponseParameters *map[string]string `json:"responseParameters,omitempty"`
    ResponseTemplates *map[string]string `json:"responseTemplates,omitempty"`
    SelectionPattern *string `json:"selectionPattern,omitempty"`
    StatusCode *string `json:"statusCode,omitempty"`
}

// IntegrationResponse's properties have constants to make dealing with diffs and property bags easier.
const (
    IntegrationResponse_ResponseParameters = "responseParameters"
    IntegrationResponse_ResponseTemplates = "responseTemplates"
    IntegrationResponse_SelectionPattern = "selectionPattern"
    IntegrationResponse_StatusCode = "statusCode"
)

/* RPC stubs for Method resource provider */

// MethodToken is the type token corresponding to the Method package type.
const MethodToken = tokens.Type("aws:apigateway/method:Method")

// MethodProviderOps is a pluggable interface for Method-related management functionality.
type MethodProviderOps interface {
    Check(ctx context.Context, obj *Method) ([]mapper.FieldError, error)
    Create(ctx context.Context, obj *Method) (resource.ID, error)
    Get(ctx context.Context, id resource.ID) (*Method, error)
    InspectChange(ctx context.Context,
        id resource.ID, old *Method, new *Method, diff *resource.ObjectDiff) ([]string, error)
    Update(ctx context.Context,
        id resource.ID, old *Method, new *Method, diff *resource.ObjectDiff) error
    Delete(ctx context.Context, id resource.ID) error
}

// MethodProvider is a dynamic gRPC-based plugin for managing Method resources.
type MethodProvider struct {
    ops MethodProviderOps
}

// NewMethodProvider allocates a resource provider that delegates to a ops instance.
func NewMethodProvider(ops MethodProviderOps) lumirpc.ResourceProviderServer {
    contract.Assert(ops != nil)
    return &MethodProvider{ops: ops}
}

func (p *MethodProvider) Check(
    ctx context.Context, req *lumirpc.CheckRequest) (*lumirpc.CheckResponse, error) {
    contract.Assert(req.GetType() == string(MethodToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr == nil || len(decerr.Failures()) == 0 {
        failures, err := p.ops.Check(ctx, obj)
        if err != nil {
            return nil, err
        }
        if len(failures) > 0 {
            decerr = mapper.NewDecodeErr(failures)
        }
    }
    return resource.NewCheckResponse(decerr), nil
}

func (p *MethodProvider) Name(
    ctx context.Context, req *lumirpc.NameRequest) (*lumirpc.NameResponse, error) {
    contract.Assert(req.GetType() == string(MethodToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr != nil {
        return nil, decerr
    }
    if obj.Name == nil || *obj.Name == "" {
        if req.Unknowns[Method_Name] {
            return nil, errors.New("Name property cannot be computed from unknown outputs")
        }
        return nil, errors.New("Name property cannot be empty")
    }
    return &lumirpc.NameResponse{Name: *obj.Name}, nil
}

func (p *MethodProvider) Create(
    ctx context.Context, req *lumirpc.CreateRequest) (*lumirpc.CreateResponse, error) {
    contract.Assert(req.GetType() == string(MethodToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr != nil {
        return nil, decerr
    }
    id, err := p.ops.Create(ctx, obj)
    if err != nil {
        return nil, err
    }
    return &lumirpc.CreateResponse{Id: string(id)}, nil
}

func (p *MethodProvider) Get(
    ctx context.Context, req *lumirpc.GetRequest) (*lumirpc.GetResponse, error) {
    contract.Assert(req.GetType() == string(MethodToken))
    id := resource.ID(req.GetId())
    obj, err := p.ops.Get(ctx, id)
    if err != nil {
        return nil, err
    }
    return &lumirpc.GetResponse{
        Properties: resource.MarshalProperties(
            nil, resource.NewPropertyMap(obj), resource.MarshalOptions{}),
    }, nil
}

func (p *MethodProvider) InspectChange(
    ctx context.Context, req *lumirpc.InspectChangeRequest) (*lumirpc.InspectChangeResponse, error) {
    contract.Assert(req.GetType() == string(MethodToken))
    id := resource.ID(req.GetId())
    old, oldprops, decerr := p.Unmarshal(req.GetOlds())
    if decerr != nil {
        return nil, decerr
    }
    new, newprops, decerr := p.Unmarshal(req.GetNews())
    if decerr != nil {
        return nil, decerr
    }
    var replaces []string
    diff := oldprops.Diff(newprops)
    if diff != nil {
        if diff.Changed("name") {
            replaces = append(replaces, "name")
        }
    }
    more, err := p.ops.InspectChange(ctx, id, old, new, diff)
    if err != nil {
        return nil, err
    }
    return &lumirpc.InspectChangeResponse{
        Replaces: append(replaces, more...),
    }, err
}

func (p *MethodProvider) Update(
    ctx context.Context, req *lumirpc.UpdateRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(MethodToken))
    id := resource.ID(req.GetId())
    old, oldprops, err := p.Unmarshal(req.GetOlds())
    if err != nil {
        return nil, err
    }
    new, newprops, err := p.Unmarshal(req.GetNews())
    if err != nil {
        return nil, err
    }
    diff := oldprops.Diff(newprops)
    if err := p.ops.Update(ctx, id, old, new, diff); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *MethodProvider) Delete(
    ctx context.Context, req *lumirpc.DeleteRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(MethodToken))
    id := resource.ID(req.GetId())
    if err := p.ops.Delete(ctx, id); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *MethodProvider) Unmarshal(
    v *pbstruct.Struct) (*Method, resource.PropertyMap, mapper.DecodeError) {
    var obj Method
    props := resource.UnmarshalProperties(nil, v, resource.MarshalOptions{RawResources: true})
    result := mapper.MapIU(props.Mappable(), &obj)
    return &obj, props, result
}

/* Marshalable Method structure(s) */

// Method is a marshalable representation of its corresponding IDL type.
type Method struct {
    Name *string `json:"name,omitempty"`
    HTTPMethod string `json:"httpMethod"`
    APIResource resource.ID `json:"apiResource"`
    RestAPI resource.ID `json:"restAPI"`
    APIKeyRequired *bool `json:"apiKeyRequired,omitempty"`
    AuthorizationType *AuthorizationType `json:"authorizationType,omitempty"`
    Authorizer *resource.ID `json:"authorizer,omitempty"`
    Integration *Integration `json:"integration,omitempty"`
    MethodResponses *[]MethodResponse `json:"methodResponses,omitempty"`
    RequestModels *map[string]resource.ID `json:"requestModels,omitempty"`
    RequestParameters *map[string]bool `json:"requestParameters,omitempty"`
}

// Method's properties have constants to make dealing with diffs and property bags easier.
const (
    Method_Name = "name"
    Method_HTTPMethod = "httpMethod"
    Method_APIResource = "apiResource"
    Method_RestAPI = "restAPI"
    Method_APIKeyRequired = "apiKeyRequired"
    Method_AuthorizationType = "authorizationType"
    Method_Authorizer = "authorizer"
    Method_Integration = "integration"
    Method_MethodResponses = "methodResponses"
    Method_RequestModels = "requestModels"
    Method_RequestParameters = "requestParameters"
)

/* Marshalable MethodResponse structure(s) */

// MethodResponse is a marshalable representation of its corresponding IDL type.
type MethodResponse struct {
    StatusCode string `json:"statusCode"`
    ResponseModels *map[string]resource.ID `json:"responseModels,omitempty"`
    ResponseParameters *map[string]bool `json:"responseParameters,omitempty"`
}

// MethodResponse's properties have constants to make dealing with diffs and property bags easier.
const (
    MethodResponse_StatusCode = "statusCode"
    MethodResponse_ResponseModels = "responseModels"
    MethodResponse_ResponseParameters = "responseParameters"
)

/* Marshalable MethodSetting structure(s) */

// MethodSetting is a marshalable representation of its corresponding IDL type.
type MethodSetting struct {
    CacheDataEncrypted *bool `json:"cacheDataEncrypted,omitempty"`
    CacheTTLInSeconds *float64 `json:"cacheTTLInSeconds,omitempty"`
    CachingEnabled *bool `json:"cachingEnabled,omitempty"`
    DataTraceEnabled *bool `json:"dataTraceEnabled,omitempty"`
    HTTPMethod *string `json:"httpMethod,omitempty"`
    LoggingLevel *LoggingLevel `json:"loggingLevel,omitempty"`
    MetricsEnabled *bool `json:"metricsEnabled,omitempty"`
    ResourcePath *string `json:"resourcePath,omitempty"`
    ThrottlingBurstLimit *float64 `json:"throttlingBurstLimit,omitempty"`
    ThrottlingRateLimit *float64 `json:"throttlingRateLimit,omitempty"`
}

// MethodSetting's properties have constants to make dealing with diffs and property bags easier.
const (
    MethodSetting_CacheDataEncrypted = "cacheDataEncrypted"
    MethodSetting_CacheTTLInSeconds = "cacheTTLInSeconds"
    MethodSetting_CachingEnabled = "cachingEnabled"
    MethodSetting_DataTraceEnabled = "dataTraceEnabled"
    MethodSetting_HTTPMethod = "httpMethod"
    MethodSetting_LoggingLevel = "loggingLevel"
    MethodSetting_MetricsEnabled = "metricsEnabled"
    MethodSetting_ResourcePath = "resourcePath"
    MethodSetting_ThrottlingBurstLimit = "throttlingBurstLimit"
    MethodSetting_ThrottlingRateLimit = "throttlingRateLimit"
)

/* Typedefs */

type (
    AuthorizationType string
    IntegrationType string
    LoggingLevel string
    PassthroughBehavior string
)

/* Constants */

const (
    AWSIAMAuthorization AuthorizationType = "AWS_IAM"
    AWSIntegration IntegrationType = "AWS"
    AWSProxyIntegration IntegrationType = "AWS_PROXY"
    CognitoAuthorization AuthorizationType = "COGNITO_USER_POOLS"
    CustomAuthorization AuthorizationType = "CUSTOM"
    HTTPIntegration IntegrationType = "HTTP"
    HTTPProxyIntegration IntegrationType = "HTTP_PROXY"
    LoggingErrorLevel LoggingLevel = "ERROR"
    LoggingInfoLevel LoggingLevel = "INFO"
    LoggingOff LoggingLevel = "OFF"
    MockIntegration IntegrationType = "MOCK"
    NoAuthorization AuthorizationType = "NONE"
    PassthroughNever PassthroughBehavior = "NEVER"
    PassthroughWhenNoMatch PassthroughBehavior = "WHEN_NO_MATCH"
    PassthroughWhenNoTemplates PassthroughBehavior = "WHEN_NO_TEMPLATES"
)


