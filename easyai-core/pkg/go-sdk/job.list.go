package sdk

import (
	"fmt"
	"net/url"
	"path/filepath"
	"strings"

	"platgit.mihoyo.com/easyai/easyai-core/pkg/api/schema"
	"platgit.mihoyo.com/easyai/easyai-core/pkg/util/json"
)

// JobListInput  获取
type JobListInput struct {
	jobID          uint32
	owner          string
	cluster        string
	namespace      string
	name           string
	description    string
	entrypoint     string
	entrypointType string
	status         string
	framework      string
	Limit          int
	Offset         int
}

// NewJobListInput new
func NewJobListInput() *JobListInput {
	return &JobListInput{}
}

// WithJobID set args
func (jl *JobListInput) WithJobID(id uint32) *JobListInput {
	jl.jobID = id
	return jl
}

// WithNamespace args
func (jl *JobListInput) WithNamespace(ns string) *JobListInput {
	jl.namespace = ns
	return jl
}

// WithCluster set args
func (jl *JobListInput) WithCluster(c string) *JobListInput {
	jl.cluster = c
	return jl
}

// WithName set args
func (jl *JobListInput) WithName(n string) *JobListInput {
	jl.name = n
	return jl
}

// WithDescription set args
func (jl *JobListInput) WithDescription(d string) *JobListInput {
	jl.description = d
	return jl
}

// WithEntrypoint set args
func (jl *JobListInput) WithEntrypoint(e string) *JobListInput {
	jl.entrypoint = e
	return jl
}

// WithEntrypointType set args
func (jl *JobListInput) WithEntrypointType(et string) *JobListInput {
	jl.entrypointType = et
	return jl
}

// WithStatus set args
func (jl *JobListInput) WithStatus(s string) *JobListInput {
	jl.status = s
	return jl
}

// WithOwner set args
func (jl *JobListInput) WithOwner(o string) *JobListInput {
	jl.owner = o
	return jl
}

// WithFramework set args
func (jl *JobListInput) WithFramework(f string) *JobListInput {
	jl.framework = f
	return jl
}

// WithOffset set args
func (jl *JobListInput) WithOffset(of int) *JobListInput {
	jl.Offset = of
	return jl
}

// WithLimit set args
func (jl *JobListInput) WithLimit(li int) *JobListInput {
	jl.Limit = li
	return jl
}

// GetQueryParams get args
func (jl *JobListInput) GetQueryParams() url.Values {
	out := url.Values{}
	query := []string{}
	if jl.jobID > 0 {
		query = append(query, "id:"+fmt.Sprint(jl.jobID))
	} else {
		if jl.owner != "" {
			query = append(query, "owner:"+jl.owner)
		}
		if jl.cluster != "" {
			query = append(query, "cluster:"+jl.cluster)
		}
		if jl.namespace != "" {
			query = append(query, "namespace:"+jl.namespace)
		}
		if jl.name != "" {
			query = append(query, "name__contains:"+jl.name)
		}
		if jl.description != "" {
			query = append(query, "description__contains:"+jl.description)
		}
		if jl.entrypoint != "" {
			query = append(query, "entrypoint__contains:"+jl.entrypoint)
		}
		if jl.entrypointType != "" {
			query = append(query, "entrypoint_type:"+jl.entrypointType)
		}
		if jl.status != "" {
			query = append(query, "status:"+jl.status)
		}
		if jl.framework != "" {
			query = append(query, "framework:"+jl.framework)
		}
	}
	if len(query) > 0 {
		out.Set("query", strings.Join(query, ","))
	}
	out.Set("limit", fmt.Sprint(jl.Limit))
	out.Set("offset", fmt.Sprint(jl.Offset))
	out.Set("sortby", "id")
	out.Set("order", "desc")
	return out
}

// GetPath return api path
func (jl *JobListInput) GetPath() string {
	return "/jobs"
}

// GetHeaders return http header
func (jl *JobListInput) GetHeaders() Header {
	return map[string]string{}
}

// GetPayload return http body
func (jl *JobListInput) GetPayload() interface{} {
	return nil
}

// Validate validate
func (jl *JobListInput) Validate() error {
	if jl.jobID > 0 {
		jl.name = ""
		jl.namespace = ""
		jl.cluster = ""
		jl.description = ""
		jl.entrypoint = ""
		jl.entrypointType = ""
		jl.status = ""
		jl.framework = ""
		jl.Limit = 1
		jl.Offset = 0
		return nil
	}
	if jl.framework != "" {
		switch schema.JobFramework(jl.framework) {
		case schema.FrameworkMxnet, schema.FrameworkAny, schema.FrameworkHVD, schema.FrameworkMPI, schema.FrameworkNotebook,
			schema.FrameworkPytorch, schema.FrameworkTensorflow, schema.FrameworkTensorboard, schema.FrameworkPaddle, schema.FrameworkXgboost,
			schema.FrameworkVolcano:
			// do nothing
		default:
			return fmt.Errorf("invalid framework: %s", jl.framework)
		}
	}
	if jl.status != "" {
		switch schema.JobStatus(jl.status) {
		case schema.JobStatusRunning, schema.JobStatusFailed, schema.JobStatusPending, schema.JobStatusQueued, schema.JobStatusRestarting,
			schema.JobStatusSucceeded, schema.JobStatusStopped, schema.JobStatusUnknown:
		// do nothing
		default:
			return fmt.Errorf("invalid job status: %s", jl.status)
		}
	}

	if jl.Offset < 0 {
		jl.Offset = 0
	}
	if jl.Limit < 0 {
		jl.Limit = 10
	}
	if jl.Limit > 40 {
		jl.Limit = 40
	}
	return nil
}

// ListJob query jobs
func (clt *Client) ListJob(input *JobListInput) ([]*schema.Job, *schema.PaginationResult, error) {
	if err := input.Validate(); err != nil {
		return nil, nil, err
	}
	path := filepath.Join(clt.Config.APIPrefix, input.GetPath())
	headers := input.GetHeaders()
	headers[HTTPHeaderToken] = clt.Config.Token
	response, err := clt.Connect.SendRequest(path, "GET", input.GetPayload(), headers, input.GetQueryParams())
	if err != nil {
		return nil, nil, err
	}
	var resp = struct {
		Code    int64  `json:"code"`
		Message string `json:"message"`
		Data    struct {
			List       []*schema.Job            `json:"list"`
			Pagination *schema.PaginationResult `json:"pagination"`
		} `json:"data"`
	}{}
	if err := json.Unmarshal(response.Body(), &resp); err != nil {
		return nil, nil, err
	}
	if resp.Code != 200 {
		return nil, nil, fmt.Errorf("got error code: %d, message: %s", resp.Code, resp.Message)
	}
	return resp.Data.List, resp.Data.Pagination, nil
}
