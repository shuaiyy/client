package sdk

/*
router defined in `internal/server/router/router.go`
*/
const (
	PathJobServiceGetJobByID    = "/clusters/%d/namespaces/%d/jobs/%d" // /clusters/:cid/namespaces/:nid/jobs/:jid
	PathJobServiceCreateJob     = "/clusters/%d/namespaces/%d/jobs"    // /clusters/:cid/namespaces/:nid/jobs
	PathJobServiceListJobs      = "/clusters/%d/namespaces/%d/jobs"    // /clusters/:cid/namespaces/:nid/jobs
	PathJobServiceDeleteJobByID = "/clusters/%d/namespaces/%d/jobs/%d" // /clusters/:cid/namespaces/:nid/jobs/:jid
	PathJobServiceUpdateJobByID = "/clusters/%d/namespaces/%d/jobs/%d" // /clusters/:cid/namespaces/:nid/jobs/:jid

)

const (
	// HTTPHeaderToken for auth
	HTTPHeaderToken = "X-Authorization-Token"

	// HTTPHeaderRequestID get request ID
	HTTPHeaderRequestID = "X-bf-request-id"

	// HTTPHeaderInvocationType stores the invocation type.
	HTTPHeaderInvocationType = "X-Fc-Invocation-Type"

	// HTTPHeaderStatefulAsyncInvocationID get invocation ID
	HTTPHeaderStatefulAsyncInvocationID = "X-Fc-Stateful-Async-Invocation-Id"

	// HTTPHeaderAccountID stores the account ID
	HTTPHeaderAccountID = "X-Fc-Account-Id"

	// HTTPHeaderFCErrorType get the error type when invoke function
	HTTPHeaderFCErrorType = "X-Fc-Error-Type"

	// HTTPHeaderSecurityToken is the header key for STS security token
	HTTPHeaderSecurityToken = "X-Fc-Security-Token"

	// HTTPHeaderInvocationLogType is the header key for invoke function log type
	HTTPHeaderInvocationLogType = "X-Fc-Log-Type"

	// HTTPHeaderInvocationLogResult is the header key for invoke function log result
	HTTPHeaderInvocationLogResult = "X-Fc-Log-Result"

	// HTTPHeaderEtag get the etag of the resource
	HTTPHeaderEtag = "Etag"

	// HTTPHeaderPrefix :Prefix string in headers
	HTTPHeaderPrefix = "x-fc-"

	// HTTPHeaderContentMD5 :Key in request headers
	HTTPHeaderContentMD5 = "Content-MD5"

	// HTTPHeaderContentType :Key in request headers
	HTTPHeaderContentType = "Content-Type"

	// HTTPHeaderUserAgent : Key in request headers
	HTTPHeaderUserAgent = "User-Agent"

	// HTTPHeaderDate :Key in request headers
	HTTPHeaderDate = "Date"

	// CustomContainerConfigAccelerationTypeDefault : default acceleration type for custom-container runtime
	CustomContainerConfigAccelerationTypeDefault = "Default"

	// CustomContainerConfigAccelerationTypeNone : disable custom-container runtime acceleration
	CustomContainerConfigAccelerationTypeNone = "None"
)

// Supported api versions
const (
	APIVersionV1 = "2016-08-15"
)

// Supported tracing types
const (
	TracingTypeJaeger = "Jaeger"
)
