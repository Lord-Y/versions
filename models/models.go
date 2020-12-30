// Package models assemble all struct, interface e.g ...
package models

// Create struct
type Create struct {
	Workload     string `form:"workload" json:"workload" binding:"required,max=100"`
	Platform     string `form:"platform" json:"platform" binding:"required,max=100"`
	Environment  string `form:"environment" json:"environment" binding:"required,max=100"`
	Version      string `form:"version" json:"version" binding:"required,max=100"`
	ChangelogURL string `form:"changelogURL,default=N/A" json:"changelogURL,default=N/A" binding:"-"`
	Raw          string `form:"raw,default=N/A" json:"raw,default=N/A" binding:"-"`
	Status       string `form:"status" json:"status" binding:"required,max=100"`
}

// UpdateStatus struct
type UpdateStatus struct {
	VersionId int    `form:"versionId" json:"versionId" binding:"required"`
	Status    string `form:"status" json:"status" binding:"required,max=100"`
}

// ReadEnvironment struct
type ReadEnvironment struct {
	Workload    string `form:"workload" json:"workload" binding:"required,max=100"`
	Platform    string `form:"platform" json:"platform" binding:"required,max=100"`
	Environment string `form:"environment" json:"environment" binding:"required,max=100"`
	Page        int    `form:"page,default=1" json:"page"`
	RangeLimit  int    `form:"rangeLimit,default=25" json:"rangeLimit"`
	StartLimit  int
	EndLimit    int
}

// ReadPlatform struct
type ReadPlatform struct {
	Workload   string `form:"workload" json:"workload" binding:"required,max=100"`
	Platform   string `form:"platform" json:"platform" binding:"required,max=100"`
	Page       int    `form:"page,default=1" json:"page"`
	RangeLimit int    `form:"rangeLimit,default=25" json:"rangeLimit"`
	StartLimit int
	EndLimit   int
}

// Raw struct
type Raw struct {
	Workload    string `form:"workload" json:"workload" binding:"required,max=100"`
	Platform    string `form:"platform" json:"platform" binding:"required,max=100"`
	Environment string `form:"environment" json:"environment" binding:"required,max=100"`
	Version     string `form:"version" json:"version" binding:"required,max=100"`
}

// RawById struct
type RawById struct {
	VersionID int `form:"versionId" json:"versionId" binding:"required"`
}
