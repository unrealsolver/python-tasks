type Task struct {
	Id     string `bson:"_id" json:"id"`
	Secret string `bson:"secret" json:"secret"`

	// time information for task
	// create - the creation time for the task, derived from the commit time or the patch creation time.
	// dispatch - the time the task runner starts up the agent on the host
	// scheduled - the time the commit is scheduled
	// start - the time the agent starts the task on the host after spinning it up
	// finish - the time the task was completed on the remote host
	CreateTime    time.Time `bson:"create_time" json:"create_time"`
	IngestTime    time.Time `bson:"injest_time" json:"ingest_time"`
	DispatchTime  time.Time `bson:"dispatch_time" json:"dispatch_time"`
	ScheduledTime time.Time `bson:"scheduled_time" json:"scheduled_time"`
	StartTime     time.Time `bson:"start_time" json:"start_time"`
	FinishTime    time.Time `bson:"finish_time" json:"finish_time"`
	ActivatedTime time.Time `bson:"activated_time" json:"activated_time"`

	Version           string              `bson:"version" json:"version,omitempty"`
	Project           string              `bson:"branch" json:"branch,omitempty"`
	Revision          string              `bson:"gitspec" json:"gitspec"`
	Priority          int64               `bson:"priority" json:"priority"`
	TaskGroup         string              `bson:"task_group" json:"task_group"`
	TaskGroupMaxHosts int                 `bson:"task_group_max_hosts,omitempty" json:"task_group_max_hosts,omitempty"`
	TaskGroupOrder    int                 `bson:"task_group_order,omitempty" json:"task_group_order,omitempty"`
	Logs              *apimodels.TaskLogs `bson:"logs,omitempty" json:"logs,omitempty"`

	// only relevant if the task is runnin.  the time of the last heartbeat
	// sent back by the agent
	LastHeartbeat time.Time `bson:"last_heartbeat" json:"last_heartbeat"`

	// used to indicate whether task should be scheduled to run
	Activated            bool         `bson:"activated" json:"activated"`
	ActivatedBy          string       `bson:"activated_by" json:"activated_by"`
	BuildId              string       `bson:"build_id" json:"build_id"`
	DistroId             string       `bson:"distro" json:"distro"`
	BuildVariant         string       `bson:"build_variant" json:"build_variant"`
	DependsOn            []Dependency `bson:"depends_on" json:"depends_on"`
	NumDependents        int          `bson:"num_dependents,omitempty" json:"num_dependents,omitempty"`
	OverrideDependencies bool         `bson:"override_dependencies,omitempty" json:"override_dependencies,omitempty"`

	DistroAliases []string `bson:"distro_aliases,omitempty" json:"distro_aliases,omitempty"`

	// Human-readable name
	DisplayName string `bson:"display_name" json:"display_name"`

	// Tags that describe the task
	Tags []string `bson:"tags,omitempty" json:"tags,omitempty"`

	// The host the task was run on. This value is empty for display
	// tasks
	HostId string `bson:"host_id" json:"host_id"`

	// the number of times this task has been restarted
	Restarts            int    `bson:"restarts" json:"restarts,omitempty"`
	Execution           int    `bson:"execution" json:"execution"`
	OldTaskId           string `bson:"old_task_id,omitempty" json:"old_task_id,omitempty"`
	Archived            bool   `bson:"archived,omitempty" json:"archived,omitempty"`
	RevisionOrderNumber int    `bson:"order,omitempty" json:"order,omitempty"`

	// task requester - this is used to help tell the
	// reason this task was created. e.g. it could be
	// because the repotracker requested it (via tracking the
	// repository) or it was triggered by a developer
	// patch request
	Requester string `bson:"r" json:"r"`

	// Status represents the various stages the task could be in
	Status  string                  `bson:"status" json:"status"`
	Details apimodels.TaskEndDetail `bson:"details" json:"task_end_details"`
	Aborted bool                    `bson:"abort,omitempty" json:"abort"`

	// TimeTaken is how long the task took to execute.  meaningless if the task is not finished
	TimeTaken time.Duration `bson:"time_taken" json:"time_taken"`

	// how long we expect the task to take from start to
	// finish. expected duration is the legacy value, but the UI
	// probably depends on it, so we maintain both values.
	ExpectedDuration       time.Duration            `bson:"expected_duration,omitempty" json:"expected_duration,omitempty"`
	ExpectedDurationStdDev time.Duration            `bson:"expected_duration_std_dev,omitempty" json:"expected_duration_std_dev,omitempty"`
	DurationPrediction     util.CachedDurationValue `bson:"duration_prediction,omitempty" json:"-"`

	// an estimate of what the task cost to run, hidden from JSON views for now
	Cost float64 `bson:"cost,omitempty" json:"-"`
	// total estimated cost of hosts this task spawned
	SpawnedHostCost float64 `bson:"spawned_host_cost,omitempty" json:"spawned_host_cost,omitempty"`

	// test results embedded from the testresults collection
	LocalTestResults []TestResult `bson:"-" json:"test_results"`

	// display task fields
	DisplayOnly       bool     `bson:"display_only,omitempty" json:"display_only,omitempty"`
	ExecutionTasks    []string `bson:"execution_tasks,omitempty" json:"execution_tasks,omitempty"`
	ResetWhenFinished bool     `bson:"reset_when_finished,omitempty" json:"reset_when_finished,omitempty"`
	DisplayTask       *Task    `bson:"-" json:"-"` // this is a local pointer from an exec to display task

	// GenerateTask indicates that the task generates other tasks, which the
	// scheduler will use to prioritize this task.
	GenerateTask bool `bson:"generate_task,omitempty" json:"generate_task,omitempty"`
	// GeneratedTasks indicates that the task has already generated other tasks. This fields
	// allows us to noop future requests, since a task should only generate others once.
	GeneratedTasks bool `bson:"generated_tasks,omitempty" json:"generated_tasks,omitempty"`
	// GeneratedBy, if present, is the ID of the task that generated this task.
	GeneratedBy string `bson:"generated_by,omitempty" json:"generated_by,omitempty"`
	// GeneratedJSON is the the configuration information to create new tasks from.
	GeneratedJSON []json.RawMessage `bson:"generate_json,omitempty" json:"generate_json,omitempty"`
	// GenerateTasksError any encountered while generating tasks.
	GenerateTasksError string `bson:"generate_error,omitempty" json:"generate_error,omitempty"`

	// Fields set if triggered by an upstream build
	TriggerID    string `bson:"trigger_id,omitempty" json:"trigger_id,omitempty"`
	TriggerType  string `bson:"trigger_type,omitempty" json:"trigger_type,omitempty"`
	TriggerEvent string `bson:"trigger_event,omitempty" json:"trigger_event,omitempty"`

	CommitQueueMerge bool `bson:"commit_queue_merge,omitempty" json:"commit_queue_merge,omitempty"`
}
