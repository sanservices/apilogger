package apilogger

// LogCat holds data regarding a logging category.
// Log categories can and should be used for all
// logging levels (INFO, WARN, ERROR)
type LogCat struct {
	Code string
	Type string
}

// All log categories. Any additions or removals to
// be registered in the `allLogCats` slice below.
var (
	// LogCatStartUp usage: service startup logs
	LogCatStartUp = LogCat{Code: "STT01", Type: "service_startup"}

	// LogCatHealth usage: health check logs
	LogCatHealth = LogCat{Code: "HTH01", Type: "health_check"}

	// LogCatRouterInit usage: api router setup logs
	LogCatRouterInit = LogCat{Code: "RTR01", Type: "router_initialization"}

	// LogCatRepoInit usage: repository setup logs
	LogCatRepoInit = LogCat{Code: "RPO01", Type: "repo_initialization"}

	// LogCatRepoOutput usage: logs relating to
	// the results of repository layer functions/methods
	LogCatRepoOutput = LogCat{Code: "RPO02", Type: "repository_output"}

	// LogCatReadConfig usage: configuration reading/init logs
	LogCatReadConfig = LogCat{Code: "CNF01", Type: "read_configuration"}

	// LogCatDatastoreConnect usage: datastore connect
	// logs (more specific than LogCatRepoInit)
	LogCatDatastoreConnect = LogCat{Code: "DTA01", Type: "datastore_connect"}

	// LogCatDatastoreClose usage: close datastore connection logs
	LogCatDatastoreClose = LogCat{Code: "DTA02", Type: "datastore_close"}

	// LogCatDatabase usage: datastore interaction
	// logs, e.g. query exec or record read errors
	LogCatDatabase = LogCat{Code: "DTA03", Type: "datastore_interaction"}

	// LogCatMarshallJSON usage: JSON marshalling logs,
	// i.e. logs for events when converting data or data structures to JSON
	LogCatMarshallJSON = LogCat{Code: "JSN01", Type: "marshall_json"}

	// LogCatUnmarshalReq usage: request decoding logs,
	// i.e. events relating to the deserialization of incoming requests
	LogCatUnmarshalReq = LogCat{Code: "REQ01", Type: "unmarshal_request_payload"}

	// LogCatAPIKey usage: logs regarding the api-key of an incoming request
	LogCatAPIKey = LogCat{Code: "REQ02", Type: "request_apikey"}

	// LogCatReqPath usage: logs of the path of an incoming request
	LogCatReqPath = LogCat{Code: "REQ03", Type: "request_path"}

	// LogCatReqValid usage: logs regarding validating an incoming request
	LogCatReqValid = LogCat{Code: "REQ04", Type: "request_validation"}

	// LogCatDebug usage: debug-related logs
	LogCatDebug = LogCat{Code: "DBG01", Type: "debug"}

	// LogCatTypeConv usage: logs regarding the conversion
	// of data from one type to another e.g. string to integer conversion errors
	LogCatTypeConv = LogCat{Code: "CNV01", Type: "type_conversion"}

	// LogCatDateTimeParse usage: date/time string parsing logs
	LogCatDateTimeParse = LogCat{Code: "PRS01", Type: "datetime_parse"}

	// LogCatServiceOutput usage: logs relating to the
	// results of service layer functions/methods
	LogCatServiceOutput = LogCat{Code: "SRV01", Type: "service_layer_output"}

	// LogCatInputValidation usage: logs relating to validating
	// the input parameters of a method/function
	LogCatInputValidation = LogCat{Code: "VAL01", Type: "method_input_validation"}

	// LogCatTemplateExec usage: logs relating to "executing" on a golang template
	LogCatTemplateExec = LogCat{Code: "TMP01", Type: "template_execution"}

	// LogCatCacheInit usage: logs relating to cache initialization
	LogCatCacheInit = LogCat{Code: "CCH01", Type: "cache_initialize"}

	// LogCatCacheRead usage: logs relating to reading from a cache
	LogCatCacheRead = LogCat{Code: "CCH02", Type: "cache_read"}

	// LogCatCacheWrite usage: logs relating to writing to a cache
	LogCatCacheWrite = LogCat{Code: "CCH03", Type: "cache_write"}

	// LogCatImplStatus usage: logs relating to the implementation progress of any particular
	// service functionality. E.g. an endpoint exists and is hit, but the intended functionality has
	// not yet been implemented
	LogCatImplStatus = LogCat{Code: "STS01", Type: "implementation_status"}

	// LogCatUncategorized usage: for temporary use in development if there has not yet
	// been an adequate category added. If this is the case, please
	// create a pull request to the logging repo with an appropriate new
	// logging category
	LogCatUncategorized = LogCat{Code: "UNCAT", Type: "uncategorized"}
)

// allLogCats is a register of all LogCats
// that can be used to test uniqueness etc.
// Any new LogCats should be registered here.
var allLogCats = []LogCat{
	LogCatStartUp,
	LogCatHealth,
	LogCatRouterInit,
	LogCatRepoInit,
	LogCatRepoOutput,
	LogCatReadConfig,
	LogCatDatastoreConnect,
	LogCatDatastoreClose,
	LogCatDatabase,
	LogCatMarshallJSON,
	LogCatUnmarshalReq,
	LogCatAPIKey,
	LogCatReqPath,
	LogCatReqValid,
	LogCatDebug,
	LogCatTypeConv,
	LogCatDateTimeParse,
	LogCatServiceOutput,
	LogCatInputValidation,
	LogCatTemplateExec,
	LogCatCacheInit,
	LogCatCacheRead,
	LogCatCacheWrite,
	LogCatImplStatus,
	LogCatUncategorized}
