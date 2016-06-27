package info

import (
	"time"
)

type YunFanConf struct {
	Ak           string
	Sk           string
	Action       string
	DBInstanceId string
	Key          string
	Token        string
}
type ALiYunReturnMC struct {
	DBInstanceID    string          `json:"DBInstanceId"`
	RequestID       string          `json:"RequestId"`
	PerformanceKeys PerformanceKeys `json:"PerformanceKeys"`
	EndTime         string          `json:"EndTime"`
	StartTime       string          `json:"StartTime"`
	Engine          string          `json:"Engine"`
}
type PerformanceKeys struct {
	PerformanceKey []PerformanceKey `json:"PerformanceKey"`
}

type PerformanceKey struct {
	Values      Values `json:"Values"`
	Key         string `json:"Key"`
	Unit        string `json:"Unit"`
	ValueFormat string `json:"ValueFormat"`
}
type Values struct {
	PerformanceValue []PerformanceValue `json:"PerformanceValue"`
}
type PerformanceValue struct {
	Value string    `json:"Value"`
	Date  time.Time `json:"Date"`
}

type SendALiYunINFO struct {
	Counter string
	Value   float64
}

type Signature struct {
	AccessKeyId      string
	Action           string
	DBInstanceId     string
	EndTime          string
	Format           string
	Key              string
	SignatureMethod  string
	SignatureNonce   string
	SignatureVersion string
	StartTime        string
	Timestamp        string
	Version          string
}
