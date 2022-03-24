package model

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/prometheus/prometheus/promql"
	"github.com/prometheus/prometheus/util/stats"
	"k8s.io/apimachinery/pkg/labels"
)

type ApiError struct {
	Typ ErrorType
	Err error
}
type ErrorType string

const (
	ErrorNone           ErrorType = ""
	ErrorTimeout        ErrorType = "timeout"
	ErrorCanceled       ErrorType = "canceled"
	ErrorExec           ErrorType = "execution"
	ErrorBadData        ErrorType = "bad_data"
	ErrorInternal       ErrorType = "internal"
	ErrorUnavailable    ErrorType = "unavailable"
	ErrorNotFound       ErrorType = "not_found"
	ErrorNotImplemented ErrorType = "not_implemented"
)

type QueryData struct {
	ResultType promql.ValueType  `json:"resultType"`
	Result     promql.Value      `json:"result"`
	Stats      *stats.QueryStats `json:"stats,omitempty"`
}

type RuleResponseItem struct {
	Id        int       `json:"id" db:"id"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Data      string    `json:"data" db:"data"`
}

type ChannelItem struct {
	Id        int       `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Name      string    `json:"name" db:"name"`
	Type      string    `json:"type" db:"type"`
	Data      string    `json:"data" db:"data"`
}

// Receiver configuration provides configuration on how to contact a receiver.
type Receiver struct {
	// A unique identifier for this receiver.
	Name string `yaml:"name" json:"name"`

	EmailConfigs     interface{} `yaml:"email_configs,omitempty" json:"email_configs,omitempty"`
	PagerdutyConfigs interface{} `yaml:"pagerduty_configs,omitempty" json:"pagerduty_configs,omitempty"`
	SlackConfigs     interface{} `yaml:"slack_configs,omitempty" json:"slack_configs,omitempty"`
	WebhookConfigs   interface{} `yaml:"webhook_configs,omitempty" json:"webhook_configs,omitempty"`
	OpsGenieConfigs  interface{} `yaml:"opsgenie_configs,omitempty" json:"opsgenie_configs,omitempty"`
	WechatConfigs    interface{} `yaml:"wechat_configs,omitempty" json:"wechat_configs,omitempty"`
	PushoverConfigs  interface{} `yaml:"pushover_configs,omitempty" json:"pushover_configs,omitempty"`
	VictorOpsConfigs interface{} `yaml:"victorops_configs,omitempty" json:"victorops_configs,omitempty"`
	SNSConfigs       interface{} `yaml:"sns_configs,omitempty" json:"sns_configs,omitempty"`
}

type ReceiverResponse struct {
	Status string   `json:"status"`
	Data   Receiver `json:"data"`
}

// AlertDiscovery has info for all active alerts.
type AlertDiscovery struct {
	Alerts []*AlertingRuleResponse `json:"rules"`
}

// Alert has info for an alert.
type AlertingRuleResponse struct {
	Labels      labels.Labels `json:"labels"`
	Annotations labels.Labels `json:"annotations"`
	State       string        `json:"state"`
	Name        string        `json:"name"`
	Id          int           `json:"id"`
	// ActiveAt    *time.Time    `json:"activeAt,omitempty"`
	// Value       float64       `json:"value"`
}

type ServiceItem struct {
	ServiceName  string  `json:"serviceName" ch:"serviceName"`
	Percentile99 float64 `json:"p99" ch:"p99"`
	AvgDuration  float64 `json:"avgDuration" ch:"avgDuration"`
	NumCalls     uint64  `json:"numCalls" ch:"numCalls"`
	CallRate     float64 `json:"callRate" ch:"callRate"`
	NumErrors    uint64  `json:"numErrors" ch:"numErrors"`
	ErrorRate    float64 `json:"errorRate" ch:"errorRate"`
	Num4XX       uint64  `json:"num4XX" ch:"num4xx"`
	FourXXRate   float64 `json:"fourXXRate" ch:"fourXXRate"`
}
type ServiceErrorItem struct {
	Time      time.Time `json:"time" ch:"time"`
	Timestamp int64     `json:"timestamp" ch:"timestamp"`
	NumErrors uint64    `json:"numErrors" ch:"numErrors"`
}
type ServiceOverviewItem struct {
	Time         time.Time `json:"time" ch:"time"`
	Timestamp    int64     `json:"timestamp" ch:"timestamp"`
	Percentile50 float64   `json:"p50" ch:"p50"`
	Percentile95 float64   `json:"p95" ch:"p95"`
	Percentile99 float64   `json:"p99" ch:"p99"`
	NumCalls     uint64    `json:"numCalls" ch:"numCalls"`
	CallRate     float64   `json:"callRate" ch:"callRate"`
	NumErrors    uint64    `json:"numErrors" ch:"numErrors"`
	ErrorRate    float64   `json:"errorRate" ch:"errorRate"`
}

type SearchSpansResult struct {
	Columns []string        `json:"columns"`
	Events  [][]interface{} `json:"events"`
}

type GetFilterSpansResponseItem struct {
	Time         time.Time `ch:"timestamp"`
	Timestamp    string    `json:"timestamp"`
	SpanID       string    `ch:"spanID" json:"spanID"`
	TraceID      string    `ch:"traceID" json:"traceID"`
	ServiceName  string    `ch:"serviceName" json:"serviceName"`
	Operation    string    `ch:"name" json:"operation"`
	DurationNano uint64    `ch:"durationNano" json:"durationNano"`
	HttpCode     string    `ch:"httpCode" json:"httpCode"`
	HttpMethod   string    `ch:"httpMethod" json:"httpMethod"`
}

type GetFilterSpansResponse struct {
	Spans      []GetFilterSpansResponseItem `json:"spans"`
	TotalSpans uint64                       `json:"totalSpans"`
}

type SearchSpanDBReponseItem struct {
	Timestamp time.Time `ch:"timestamp"`
	TraceID   string    `ch:"traceID"`
	Model     string    `ch:"model"`
}

type Event struct {
	Name         string                 `json:"name,omitempty"`
	TimeUnixNano uint64                 `json:"timeUnixNano,omitempty"`
	AttributeMap map[string]interface{} `json:"attributeMap,omitempty"`
	IsError      bool                   `json:"isError,omitempty"`
}

type SearchSpanReponseItem struct {
	TimeUnixNano uint64            `json:"timestamp"`
	SpanID       string            `json:"spanID"`
	TraceID      string            `json:"traceID"`
	ServiceName  string            `json:"serviceName"`
	Name         string            `json:"name"`
	Kind         int32             `json:"kind"`
	References   []OtelSpanRef     `json:"references,omitempty"`
	DurationNano int64             `json:"durationNano"`
	TagMap       map[string]string `json:"tagMap"`
	Events       []string          `json:"event"`
	HasError     bool              `json:"hasError"`
}

type OtelSpanRef struct {
	TraceId string `json:"traceId,omitempty"`
	SpanId  string `json:"spanId,omitempty"`
	RefType string `json:"refType,omitempty"`
}

func (ref *OtelSpanRef) toString() string {

	retString := fmt.Sprintf(`{TraceId=%s, SpanId=%s, RefType=%s}`, ref.TraceId, ref.SpanId, ref.RefType)

	return retString
}

func (item *SearchSpanReponseItem) GetValues() []interface{} {

	references := []OtelSpanRef{}
	jsonbody, _ := json.Marshal(item.References)
	json.Unmarshal(jsonbody, &references)

	referencesStringArray := []string{}
	for _, item := range references {
		referencesStringArray = append(referencesStringArray, item.toString())
	}

	if item.Events == nil {
		item.Events = []string{}
	}
	keys := make([]string, 0, len(item.TagMap))
	values := make([]string, 0, len(item.TagMap))

	for k, v := range item.TagMap {
		keys = append(keys, k)
		values = append(values, v)
	}
	returnArray := []interface{}{item.TimeUnixNano, item.SpanID, item.TraceID, item.ServiceName, item.Name, strconv.Itoa(int(item.Kind)), strconv.FormatInt(item.DurationNano, 10), keys, values, referencesStringArray, item.Events, item.HasError}

	return returnArray
}

type ServiceMapDependencyItem struct {
	SpanId       string `json:"spanId,omitempty" ch:"spanID"`
	ParentSpanId string `json:"parentSpanId,omitempty" ch:"parentSpanID"`
	ServiceName  string `json:"serviceName,omitempty" ch:"serviceName"`
}

type UsageItem struct {
	Time      time.Time `json:"time,omitempty" ch:"time"`
	Timestamp uint64    `json:"timestamp" ch:"timestamp"`
	Count     uint64    `json:"count" ch:"count"`
}

type TopEndpointsItem struct {
	Percentile50 float64 `json:"p50" ch:"p50"`
	Percentile95 float64 `json:"p95" ch:"p95"`
	Percentile99 float64 `json:"p99" ch:"p99"`
	NumCalls     uint64  `json:"numCalls" ch:"numCalls"`
	Name         string  `json:"name" ch:"name"`
}

type TagFilters struct {
	TagKeys string `json:"tagKeys" ch:"tagKeys"`
}

type TagValues struct {
	TagValues string `json:"tagValues" ch:"tagValues"`
}
type ServiceMapDependencyResponseItem struct {
	Parent    string `json:"parent,omitempty" ch:"parent"`
	Child     string `json:"child,omitempty" ch:"child"`
	CallCount int    `json:"callCount,omitempty" ch:"callCount"`
}

type GetFilteredSpansAggregatesResponse struct {
	Items map[int64]SpanAggregatesResponseItem `json:"items"`
}
type SpanAggregatesResponseItem struct {
	Timestamp int64              `json:"timestamp,omitempty" `
	Value     float32            `json:"value,omitempty"`
	GroupBy   map[string]float32 `json:"groupBy,omitempty"`
}
type SpanAggregatesDBResponseItem struct {
	Timestamp  int64     `ch:"timestamp" `
	Time       time.Time `ch:"time"`
	Value      uint64    `ch:"value"`
	FloatValue float32   `ch:"floatValue"`
	GroupBy    string    `ch:"groupBy"`
}

type SetTTLResponseItem struct {
	Message string `json:"message"`
}

type DBResponseTTL struct {
	EngineFull string `ch:"engine_full"`
}

type GetTTLResponseItem struct {
	MetricsTime int `json:"metrics_ttl_duration_hrs"`
	TracesTime  int `json:"traces_ttl_duration_hrs"`
}

type DBResponseMinMaxDuration struct {
	MinDuration uint64 `ch:"min(durationNano)"`
	MaxDuration uint64 `ch:"max(durationNano)"`
}

type DBResponseServiceName struct {
	ServiceName string `ch:"serviceName"`
	Count       uint64 `ch:"count"`
}

type DBResponseHttpCode struct {
	HttpCode string `ch:"httpCode"`
	Count    uint64 `ch:"count"`
}

type DBResponseHttpRoute struct {
	HttpRoute string `ch:"httpRoute"`
	Count     uint64 `ch:"count"`
}

type DBResponseHttpUrl struct {
	HttpUrl string `ch:"httpUrl"`
	Count   uint64 `ch:"count"`
}

type DBResponseHttpMethod struct {
	HttpMethod string `ch:"httpMethod"`
	Count      uint64 `ch:"count"`
}

type DBResponseHttpHost struct {
	HttpHost string `ch:"httpHost"`
	Count    uint64 `ch:"count"`
}

type DBResponseOperation struct {
	Operation string `ch:"name"`
	Count     uint64 `ch:"count"`
}

type DBResponseComponent struct {
	Component string `ch:"component"`
	Count     uint64 `ch:"count"`
}

type DBResponseErrors struct {
	NumErrors uint64 `ch:"numErrors"`
}

type DBResponseTotal struct {
	NumTotal uint64 `ch:"numTotal"`
}

type SpanFiltersResponse struct {
	ServiceName map[string]uint64 `json:"serviceName"`
	Status      map[string]uint64 `json:"status"`
	Duration    map[string]uint64 `json:"duration"`
	Operation   map[string]uint64 `json:"operation"`
	HttpCode    map[string]uint64 `json:"httpCode"`
	HttpUrl     map[string]uint64 `json:"httpUrl"`
	HttpMethod  map[string]uint64 `json:"httpMethod"`
	HttpRoute   map[string]uint64 `json:"httpRoute"`
	HttpHost    map[string]uint64 `json:"httpHost"`
	Component   map[string]uint64 `json:"component"`
}
type Error struct {
	ExceptionType  string    `json:"exceptionType" ch:"exceptionType"`
	ExceptionMsg   string    `json:"exceptionMessage" ch:"exceptionMessage"`
	ExceptionCount int64     `json:"exceptionCount" ch:"exceptionCount"`
	LastSeen       time.Time `json:"lastSeen" ch:"lastSeen"`
	FirstSeen      time.Time `json:"firstSeen" ch:"firstSeen"`
	ServiceName    string    `json:"serviceName" ch:"serviceName"`
}

type ErrorWithSpan struct {
	ErrorID            string    `json:"errorId" ch:"errorID"`
	ExceptionType      string    `json:"exceptionType" ch:"exceptionType"`
	ExcepionStacktrace string    `json:"excepionStacktrace" ch:"excepionStacktrace"`
	ExceptionEscaped   string    `json:"exceptionEscaped" ch:"exceptionEscaped"`
	ExceptionMsg       string    `json:"exceptionMessage" ch:"exceptionMessage"`
	Timestamp          time.Time `json:"timestamp" ch:"timestamp"`
	SpanID             string    `json:"spanID" ch:"spanID"`
	TraceID            string    `json:"traceID" ch:"traceID"`
	ServiceName        string    `json:"serviceName" ch:"serviceName"`
	NewerErrorID       string    `json:"newerErrorId" ch:"newerErrorId"`
	OlderErrorID       string    `json:"olderErrorId" ch:"olderErrorId"`
}
