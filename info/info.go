package info

import (
	"time"
)

type GetInfoConf struct {
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
	AccessKeyId      string //`json:"AccessKeyId"`
	Action           string //`json:"Action"`
	DBInstanceId     string //`json:"DBInstanceId"`
	EndTime          string //`json:"EndTime"`
	Format           string //`json:"Format"`
	Key              string //`json:"Key"`
	SignatureMethod  string //`json:"SignatureMethod"`
	SignatureNonce   string //`json:"SignatureNonce"`
	SignatureVersion string //`json:"SignatureVersion"`
	StartTime        string //`json:"StartTime"`
	Timestamp        string //`json:"Timestamp"`
	Version          string //`json:"Version"`
}
