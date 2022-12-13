package run

import (
	"fmt"

	"platgit.mihoyo.com/easyai/easyai-core/pkg/api/schema"
	"platgit.mihoyo.com/easyai/easyai-core/pkg/go-sdk"
	"platgit.mihoyo.com/easyai/easyai-core/pkg/util/json"

	"seelie/internal/cli/client"
	"seelie/internal/cli/utils/pretty"
)

// GetJobByID ...
func GetJobByID(jid uint32) error {
	clt, err := client.NewClient()
	if err != nil {
		return err
	}
	job, err := clt.GetJob(sdk.NewJobGetInput().WithJobID(jid))
	if err != nil {
		return err
	}
	// todo beautify the output
	fmt.Println(pretty.Blue("job id"), job.ID)
	fmt.Println(pretty.Blue("job name"), job.Name)
	fmt.Println(pretty.Blue("job description"), job.Description)
	fmt.Println(pretty.Blue("job owner"), job.Owner)
	fmt.Println(pretty.Blue("framework"), job.Framework)
	fmt.Println(pretty.Blue("status"), job.Status)
	if job.Status == schema.JobStatusRunning && job.State == schema.JobStateStop {
		fmt.Println(pretty.Red("job is stopping"), "the job is marked Stop, will be Stopped soon")
	}
	fmt.Println(pretty.Blue("message"), job.Message)
	fmt.Println(pretty.Blue("reason"), job.Reason)
	if len(job.Tasks) > 0 {
		fmt.Println("job task list:")
		table := [][]interface{}{
			// header
			[]interface{}{pretty.Green("pod name"), pretty.Green("role"), pretty.Green("index"),
				pretty.Green("gpu"), pretty.Green("node_ip"), pretty.Green("status"), pretty.Green("start time")},
		}
		for _, t := range job.Tasks {
			table = append(table, []interface{}{t.PodName, t.Role, t.Index, t.GPU, t.NodeIP, t.Status, t.StartTime})
		}
		pretty.PrintTable(table)
	}

	return nil
}

// DeleteJobByID 删除Job
func DeleteJobByID(jid uint32) error {
	clt, err := client.NewClient()
	if err != nil {
		return err
	}
	message, err := clt.DeleteJob(sdk.NewJobDeleteInput().WithJobID(jid))
	if err != nil {
		return err
	}
	fmt.Println("message:", pretty.Blue(message))
	return nil
}

// StopJobByID 删除Job
func StopJobByID(jid uint32) error {
	clt, err := client.NewClient()
	if err != nil {
		return err
	}
	message, err := clt.StopJob(sdk.NewJobStopInput().WithJobID(jid))
	if err != nil {
		return err
	}
	fmt.Println("message:", pretty.Blue(message))
	return nil
}

// ListJobs jobs
func ListJobs(input *sdk.JobListInput) error {
	clt, err := client.NewClient()
	if err != nil {
		return err
	}
	jobs, pages, err := clt.ListJob(input)
	if err != nil {
		return err
	}

	fmt.Println(pretty.Blue("result total count:"), pages.Total, pretty.Blue("offset:"), pages.Offset, pretty.Blue("limit:"), pages.Limit)
	table := [][]interface{}{
		// header
		[]interface{}{pretty.Green("id"), pretty.Green("name"), pretty.Green("owner"),
			pretty.Green("status"), pretty.Green("start time"), pretty.Green("duration(min)")},
	}
	for _, job := range jobs {
		table = append(table, []interface{}{job.ID, job.Name, job.Owner, job.Status, job.StartTime, job.Duration})
	}
	pretty.PrintTable(table)
	return nil
}

// SubmitTFJob 提交tf job
func SubmitTFJob(common schema.CommonConfig, tfJob schema.TFJobConfig) error {
	var job schema.Job
	job.CommonConfig = common
	job.Framework = schema.FrameworkTensorflow
	job.FrameworkConfig = json.MarshalToString(tfJob)
	input := sdk.NewJobSubmitInput().WithJob(job) // todo sdk 需要提供更友好的payload构建方式
	clt, err := client.NewClient()
	if err != nil {
		return err
	}
	id, err := clt.SubmitJob(input)
	if err != nil {
		return err
	}
	fmt.Println(pretty.Green("submit success, id:"), id)
	fmt.Println(pretty.Blue(fmt.Sprintf("job_detail: https://ml.ssr.mihoyo.com/job_detail/%d", id)))
	fmt.Println("通过seelie-cli查看job:", pretty.Green(fmt.Sprintf("seelie job get --job_id %d", id)))
	return nil
}
