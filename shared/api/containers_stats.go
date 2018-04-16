package api

import (
	"time"
)

//Struct that gets all the containers stats
type ContainersStats struct {
	StatsTime time.Time
}

type ContainerStats struct {
	//Time that stats were created
	CreatedAt	time.Time	`json:"created_at" yaml:"created_at"`
	//ContainerStats types (CPU, Memory, Network, ...)
	Type		string		`json:"type" yaml:"type"`
}

type ContanerCPUStats struct {
	//Time that stats were gathered
	CreatedAt	time.Time		`json:"created_at" yaml:"created_at"`
	user		int64			`json:"user" yaml:"user"`
	system		int64			`json:"system" yaml:"system"`
}
