// Package models assemble all struct, interface e.g ...
package models

import "time"

// Create struct
type Create struct {
	Workload     string `form:"workload" json:"workload" binding:"required,max=100"`
	Platform     string `form:"platform" json:"platform" binding:"required,max=100"`
	Environment  string `form:"environment" json:"environment" binding:"required,max=100"`
	Version      string `form:"version" json:"version" binding:"required,max=100"`
	ChangelogURL string `form:"changelogURL,default=N/A" json:"changelogURL" binding:"-"`
	Raw          string `form:"raw,default=N/A" json:"raw" binding:"-"`
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

// ReadEnvironmentLatest struct
type ReadEnvironmentLatest struct {
	Workload    string `form:"workload" json:"workload" binding:"required,max=100"`
	Platform    string `form:"platform" json:"platform" binding:"required,max=100"`
	Environment string `form:"environment" json:"environment" binding:"required,max=100"`
	Whatever    bool
}

// DBReadCommon
type DBReadCommon struct {
	Versions_id   int       `json:"versions_id"`
	Workload      string    `json:"workload"`
	Platform      string    `json:"platform"`
	Environment   string    `json:"environment"`
	Version       string    `json:"version"`
	Changelog_url string    `json:"changelog_url"`
	Raw           string    `json:"raw"`
	Status        string    `json:"status"`
	Date          time.Time `json:"date"`
	Total         int64     `json:"total"`
}

// DBCommons
type DBCommons struct {
	Versions_id   int       `json:"versions_id"`
	Workload      string    `json:"workload"`
	Platform      string    `json:"platform"`
	Environment   string    `json:"environment"`
	Version       string    `json:"version"`
	Changelog_url string    `json:"changelog_url"`
	Raw           string    `json:"raw"`
	Status        string    `json:"status"`
	Date          time.Time `json:"date"`
}

// DBReadDistinctWorkloads
type DBReadDistinctWorkloads struct {
	Workload    string `json:"workload"`
	Platform    string `json:"platform"`
	Environment string `json:"environment"`
}

// DBRaw
type DBRaw struct {
	Raw string `json:"raw"`
}

// DBReadForUnitTesting
type DBReadForUnitTesting struct {
	Versions_id int    `json:"versions_id"`
	Workload    string `json:"workload"`
	Platform    string `json:"platform"`
	Environment string `json:"environment"`
}

// DbVersion struct
type DbVersion struct {
	Version string `json:"version"`
}

// DBGetLastXDaysDeployments
type DBGetLastXDaysDeployments struct {
	Total       int64     `json:"total"`
	Workload    string    `json:"workload"`
	Platform    string    `json:"platform"`
	Environment string    `json:"environment"`
	Status      string    `json:"status"`
	Date        time.Time `json:"date"` // using string because time.Time doesn't work with mysql DATE_FORMAT
}
