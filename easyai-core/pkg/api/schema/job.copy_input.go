package schema

import (
	"strings"

	"platgit.mihoyo.com/easyai/easyai-core/pkg/util/json"
	"platgit.mihoyo.com/easyai/easyai-core/pkg/util/structure"
	"platgit.mihoyo.com/easyai/easyai-core/pkg/util/yaml"
)

var _removeKeys = []string{"id", "uk_uuid", "status", "state", "message", "reason", "result", "is_deleted", "duration",
	"created_at", "modified_at", "deleted_at", "tasks", "fe_tasks", "fe_resource", "hardware"}

// SubmitConfig ...
func (j *Job) SubmitConfig() map[string]interface{} {
	var job Job
	_ = structure.Copy(j, &job)
	// clear job
	job.Tasks = nil
	for _, key := range job.MagicFlags {
		if !strings.HasPrefix(key, "m_") {
			delete(job.MagicFlags, key)
		}
	}
	// clear SSH_PUBLIC_KEY ?

	// trans to map
	var res = map[string]interface{}{}
	jsonStr := json.MarshalToString(job)
	_ = json.Unmarshal([]byte(jsonStr), &res)
	// clear field
	for _, key := range _removeKeys {
		delete(res, key)
	}
	return res
}

// SubmitYAMLConfig to yaml string
func (j *Job) SubmitYAMLConfig() string {
	return yaml.MarshalToString(j.SubmitConfig())
}

// SubmitJSONConfig to json string
func (j *Job) SubmitJSONConfig() string {
	return json.MarshalToString(j.SubmitConfig())
}
