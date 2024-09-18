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
	// LogCatRunningTask usage: Schehuled taks is running
	LogCatRunningTask = LogCat{Code: "CJ001", Type: "running-task"}

	// LogCatStartUp usage: service startup logs
	LogCatStartUp = LogCat{Code: "STT001", Type: "service_startup"}

	// LogCatHealth usage: health check logs
	LogCatHealth = LogCat{Code: "HTH001", Type: "health_check"}

	// LogCatRouterInit usage: api router setup logs
	LogCatRouterInit = LogCat{Code: "RTR001", Type: "router_initialization"}

	// LogCatRepoInit usage: repository setup logs
	LogCatRepoInit = LogCat{Code: "RPO001", Type: "repo_initialization"}

	// LogCatRepoOutput usage: logs relating to
	// the results of repository layer functions/methods
	LogCatRepoOutput = LogCat{Code: "RPO002", Type: "repository_output"}

	// LogCatReadConfig usage: configuration reading/init logs
	LogCatReadConfig = LogCat{Code: "CNF001", Type: "read_configuration"}

	// LogCatDatastoreConnect usage: datastore connect
	// logs (more specific than LogCatRepoInit)
	LogCatDatastoreConnect = LogCat{Code: "DTA001", Type: "datastore_connect"}

	// LogCatDatastoreClose usage: close datastore connection logs
	LogCatDatastoreClose = LogCat{Code: "DTA002", Type: "datastore_close"}

	// LogCatDatabase usage: datastore interaction
	// logs, e.g. query exec or record read errors
	LogCatDatabase = LogCat{Code: "DTA003", Type: "datastore_interaction"}

	// LogCatMarshallJSON usage: JSON marshalling logs,
	// i.e. logs for events when converting data or data structures to JSON
	LogCatMarshallJSON = LogCat{Code: "JSN001", Type: "marshall_json"}

	// LogCatUnmarshalReq usage: request decoding logs,
	// i.e. events relating to the deserialization of incoming requests
	LogCatUnmarshalReq = LogCat{Code: "REQ001", Type: "unmarshal_request_payload"}

	// LogCatFileRead usage: logs relating to file reads
	LogCatFileRead = LogCat{Code: "FIL001", Type: "file_read"}

	// LogCatFileWrite usage: logs relating to file writes
	LogCatFileWrite = LogCat{Code: "FIL002", Type: "file_write"}

	// LogCatTLSLoadCerts usage: logs relating to reading in tls certificates
	LogCatTLSLoadCerts = LogCat{Code: "TLS001", Type: "tls_load_certs"}

	// LogCatAPIKey usage: logs regarding the api-key of an incoming request
	LogCatAPIKey = LogCat{Code: "REQ002", Type: "request_apikey"}

	// LogCatReqPath usage: logs of the path of an incoming request
	LogCatReqPath = LogCat{Code: "REQ003", Type: "request_path"}

	// LogCatReqValid usage: logs regarding validating an incoming request
	LogCatReqValid = LogCat{Code: "REQ004", Type: "request_validation"}

	// LogCatDebug usage: debug-related logs
	LogCatDebug = LogCat{Code: "DBG001", Type: "debug"}

	// LogCatTypeConv usage: logs regarding the conversion
	// of data from one type to another e.g. string to integer conversion errors
	LogCatTypeConv = LogCat{Code: "CNV001", Type: "type_conversion"}

	// LogCatDateTimeParse usage: date/time string parsing logs
	LogCatDateTimeParse = LogCat{Code: "PRS001", Type: "datetime_parse"}

	// LogCatServiceOutput usage: logs relating to the
	// results of service layer functions/methods
	LogCatServiceOutput = LogCat{Code: "SRV001", Type: "service_layer_output"}

	// LogCatInputValidation usage: logs relating to validating
	// the input parameters of a method/function
	LogCatInputValidation = LogCat{Code: "VAL001", Type: "method_input_validation"}

	// LogCatInvalidType can be used for logs concerning type validation or casting
	LogCatInvalidType = LogCat{Code: "VAL002", Type: "type_validation"}

	// LogCatTemplateExec usage: logs relating to "executing" on a golang template
	LogCatTemplateExec = LogCat{Code: "TMP001", Type: "template_execution"}

	// LogCatCacheInit usage: logs relating to cache initialization
	LogCatCacheInit = LogCat{Code: "CCH001", Type: "cache_initialize"}

	// LogCatCacheRead usage: logs relating to reading from a cache
	LogCatCacheRead = LogCat{Code: "CCH002", Type: "cache_read"}

	// LogCatCacheWrite usage: logs relating to writing to a cache
	LogCatCacheWrite = LogCat{Code: "CCH003", Type: "cache_write"}

	// LogCatImplStatus usage: logs relating to the implementation progress of any particular
	// service functionality. E.g. an endpoint exists and is hit, but the intended functionality has
	// not yet been implemented
	LogCatImplStatus = LogCat{Code: "STS001", Type: "implementation_status"}

	// LogCatKafkaSchemaReg usage: logs relating to initializing or
	// interating with a kafka schema registry
	LogCatKafkaSchemaReg = LogCat{Code: "KFK001", Type: "kafka_schemareg_init"}

	// LogCatKafkaDecode usage: logs relating to decoding kafka messages
	LogCatKafkaDecode = LogCat{Code: "KFK002", Type: "kafka_message_decode"}

	// LogCatKafkaEncode usage: logs relating to encoding kafka messages
	LogCatKafkaEncode = LogCat{Code: "KFK003", Type: "kafka_message_encode"}

	// LogCatKafkaConsumerInit usage: logs relating to initializing a kafka consumer
	LogCatKafkaConsumerInit = LogCat{Code: "KFK004", Type: "kafka_consumer_init"}

	// LogCatKafkaProducerInit usage: logs relating to initializing a kafka producer
	LogCatKafkaProducerInit = LogCat{Code: "KFK005", Type: "kafka_producer_init"}

	// LogCatKafkaConsumerClose usage: logs relating to attempting to close a kafka consumer
	LogCatKafkaConsumerClose = LogCat{Code: "KFK006", Type: "kafka_consumer_close"}

	// LogCatKafkaConsume usage: logs relating to the consumption of kafka messages
	LogCatKafkaConsume = LogCat{Code: "KFK007", Type: "kafka_consume"}

	// LogCatKafkaProduce usage: logs relating to the production of kafka messages
	LogCatKafkaProduce = LogCat{Code: "KFK008", Type: "kafka_produce"}

	// LogCatKafkaConfig usage: logs relating to configuring a kafka integration
	LogCatKafkaConfig = LogCat{Code: "KFK009", Type: "kafka_configure"}

	// LogCatKafkaProcessMessage usage: logs relating to the processing of kafka messages
	LogCatKafkaProcessMessage = LogCat{Code: "KFK010", Type: "kafka_process_message"}

	// LogCatKafkaCommitOffset usage: logs relating to committing offsets after reading
	// and successfully processing a kafka message and its contents
	LogCatKafkaCommitOffset = LogCat{Code: "KFK011", Type: "kafka_commit_offset"}

	// LogCatExternal usage: miscellaneous external library operation logs
	LogCatExternal = LogCat{Code: "EXT001", Type: "external_lib_op"}

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
	LogCatRunningTask,
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
	LogCatFileRead,
	LogCatFileWrite,
	LogCatTLSLoadCerts,
	LogCatAPIKey,
	LogCatReqPath,
	LogCatReqValid,
	LogCatDebug,
	LogCatTypeConv,
	LogCatDateTimeParse,
	LogCatServiceOutput,
	LogCatInputValidation,
	LogCatInvalidType,
	LogCatTemplateExec,
	LogCatCacheInit,
	LogCatCacheRead,
	LogCatCacheWrite,
	LogCatImplStatus,
	LogCatKafkaSchemaReg,
	LogCatKafkaDecode,
	LogCatKafkaEncode,
	LogCatKafkaConsumerInit,
	LogCatKafkaProducerInit,
	LogCatKafkaConsumerClose,
	LogCatKafkaConsume,
	LogCatKafkaProduce,
	LogCatKafkaConfig,
	LogCatKafkaProcessMessage,
	LogCatKafkaCommitOffset,
	LogCatExternal,
	LogCatUncategorized,
}
