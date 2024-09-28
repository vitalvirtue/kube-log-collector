package types

import "time"

// CollectorOptions, log toplayıcı için yapılandırma seçeneklerini tanımlar
type CollectorOptions struct {
	Namespace  string
	PodLabel   string
	OutputFile string
	Since      time.Duration
	Tail       int64
}

// LogEntry, bir log girdisini temsil eder
type LogEntry struct {
	Timestamp time.Time
	Namespace string
	PodName   string
	Container string
	Message   string
}

// PodInfo, bir pod hakkında temel bilgileri içerir
type PodInfo struct {
	Name      string
	Namespace string
	Labels    map[string]string
}

// CollectionResult, log toplama işleminin sonucunu temsil eder
type CollectionResult struct {
	TotalLogs   int
	FailedPods  []string
	ElapsedTime time.Duration
}

// LogFilter, log filtreleme için kullanılan kriterleri tanımlar
type LogFilter struct {
	StartTime time.Time
	EndTime   time.Time
	Keywords  []string
	LogLevel  string
}

// CollectorStats, log toplayıcının istatistiklerini tutar
type CollectorStats struct {
	StartTime        time.Time
	EndTime          time.Time
	ProcessedPods    int
	TotalLogs        int
	TotalBytesRead   int64
	AverageLogSize   float64
	LargestLogSize   int64
	LargestLogPod    string
	FailedCollections int
}

// OutputFormat, çıktı formatı için bir enum tipi
type OutputFormat int

const (
	FormatText OutputFormat = iota
	FormatJSON
	FormatCSV
)

// String, OutputFormat için string temsili döndürür
func (f OutputFormat) String() string {
	return [...]string{"Text", "JSON", "CSV"}[f]
}