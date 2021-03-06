package utils

const (
	VERSION                      = "0.9.1rc6"
	POSTGRES                     = "postgres"
	MYSQL                        = "mysql"
	MONGO                        = "mongo"
	REDIS                        = "redis"
	LOCALHOST                    = "127.0.0.1"
	FSCDR_FILE_CSV               = "freeswitch_file_csv"
	FSCDR_HTTP_JSON              = "freeswitch_http_json"
	NOT_IMPLEMENTED              = "not implemented"
	PREPAID                      = "prepaid"
	META_PREPAID                 = "*prepaid"
	POSTPAID                     = "postpaid"
	META_POSTPAID                = "*postpaid"
	PSEUDOPREPAID                = "pseudoprepaid"
	META_PSEUDOPREPAID           = "*pseudoprepaid"
	RATED                        = "rated"
	META_RATED                   = "*rated"
	META_NONE                    = "*none"
	META_NOW                     = "*now"
	ERR_NOT_IMPLEMENTED          = "NOT_IMPLEMENTED"
	ERR_SERVER_ERROR             = "SERVER_ERROR"
	ERR_NOT_FOUND                = "NOT_FOUND"
	ERR_MANDATORY_IE_MISSING     = "MANDATORY_IE_MISSING"
	ERR_EXISTS                   = "EXISTS"
	ERR_BROKEN_REFERENCE         = "BROKEN_REFERENCE"
	ERR_PARSER_ERROR             = "PARSER_ERROR"
	ERR_INVALID_PATH             = "INVALID_PATH"
	TBL_TP_TIMINGS               = "tp_timings"
	TBL_TP_DESTINATIONS          = "tp_destinations"
	TBL_TP_RATES                 = "tp_rates"
	TBL_TP_DESTINATION_RATES     = "tp_destination_rates"
	TBL_TP_RATING_PLANS          = "tp_rating_plans"
	TBL_TP_RATE_PROFILES         = "tp_rating_profiles"
	TBL_TP_SHARED_GROUPS         = "tp_shared_groups"
	TBL_TP_CDR_STATS             = "tp_cdr_stats"
	TBL_TP_LCRS                  = "tp_lcr_rules"
	TBL_TP_ACTIONS               = "tp_actions"
	TBL_TP_ACTION_PLANS          = "tp_action_plans"
	TBL_TP_ACTION_TRIGGERS       = "tp_action_triggers"
	TBL_TP_ACCOUNT_ACTIONS       = "tp_account_actions"
	TBL_TP_DERIVED_CHARGERS      = "tp_derived_chargers"
	TBL_CDRS_PRIMARY             = "cdrs_primary"
	TBL_CDRS_EXTRA               = "cdrs_extra"
	TBL_COST_DETAILS             = "cost_details"
	TBL_RATED_CDRS               = "rated_cdrs"
	TIMINGS_CSV                  = "Timings.csv"
	DESTINATIONS_CSV             = "Destinations.csv"
	RATES_CSV                    = "Rates.csv"
	DESTINATION_RATES_CSV        = "DestinationRates.csv"
	RATING_PLANS_CSV             = "RatingPlans.csv"
	RATING_PROFILES_CSV          = "RatingProfiles.csv"
	SHARED_GROUPS_CSV            = "SharedGroups.csv"
	LCRS_CSV                     = "LcrRules.csv"
	ACTIONS_CSV                  = "Actions.csv"
	ACTION_PLANS_CSV             = "ActionPlans.csv"
	ACTION_TRIGGERS_CSV          = "ActionTriggers.csv"
	ACCOUNT_ACTIONS_CSV          = "AccountActions.csv"
	DERIVED_CHARGERS_CSV         = "DerivedChargers.csv"
	CDR_STATS_CSV                = "CdrStats.csv"
	TIMINGS_NRCOLS               = 6
	DESTINATIONS_NRCOLS          = 2
	RATES_NRCOLS                 = 6
	DESTINATION_RATES_NRCOLS     = 7
	DESTRATE_TIMINGS_NRCOLS      = 4
	RATE_PROFILES_NRCOLS         = 8
	SHARED_GROUPS_NRCOLS         = 4
	LCRS_NRCOLS                  = 11
	ACTIONS_NRCOLS               = 15
	ACTION_PLANS_NRCOLS          = 4
	ACTION_TRIGGERS_NRCOLS       = 19
	ACCOUNT_ACTIONS_NRCOLS       = 5
	DERIVED_CHARGERS_NRCOLS      = 19
	CDR_STATS_NRCOLS             = 23
	ROUNDING_UP                  = "*up"
	ROUNDING_MIDDLE              = "*middle"
	ROUNDING_DOWN                = "*down"
	ANY                          = "*any"
	COMMENT_CHAR                 = '#'
	CSV_SEP                      = ','
	FALLBACK_SEP                 = ';'
	INFIELD_SEP                  = ";"
	FIELDS_SEP                   = ","
	STATIC_HDRVAL_SEP            = "::"
	REGEXP_PREFIX                = "~"
	FILTER_VAL_START             = "("
	FILTER_VAL_END               = ")"
	JSON                         = "json"
	GOB                          = "gob"
	MSGPACK                      = "msgpack"
	CSV_LOAD                     = "CSVLOAD"
	CGRID                        = "cgrid"
	ORDERID                      = "orderid"
	ACCID                        = "accid"
	CDRHOST                      = "cdrhost"
	CDRSOURCE                    = "cdrsource"
	REQTYPE                      = "reqtype"
	DIRECTION                    = "direction"
	TENANT                       = "tenant"
	CATEGORY                     = "category"
	ACCOUNT                      = "account"
	SUBJECT                      = "subject"
	DESTINATION                  = "destination"
	SETUP_TIME                   = "setup_time"
	ANSWER_TIME                  = "answer_time"
	USAGE                        = "usage"
	SUPPLIER                     = "supplier"
	MEDI_RUNID                   = "mediation_runid"
	RATED_ACCOUNT                = "rated_account"
	RATED_SUBJECT                = "rated_subject"
	COST                         = "cost"
	COST_DETAILS                 = "cost_details"
	DEFAULT_RUNID                = "*default"
	META_DEFAULT                 = "*default"
	STATIC_VALUE_PREFIX          = "^"
	CSV                          = "csv"
	DRYRUN                       = "dry_run"
	COMBIMED                     = "combimed"
	INTERNAL                     = "internal"
	ZERO_RATING_SUBJECT_PREFIX   = "*zero"
	OK                           = "OK"
	CDRE_FIXED_WIDTH             = "fwv"
	XML_PROFILE_PREFIX           = "*xml:"
	CDRE                         = "cdre"
	CDRC                         = "cdrc"
	MASK_CHAR                    = "*"
	CONCATENATED_KEY_SEP         = ":"
	FORKED_CDR                   = "forked_cdr"
	UNIT_TEST                    = "UNIT_TEST"
	HDR_VAL_SEP                  = "/"
	MONETARY                     = "*monetary"
	SMS                          = "*sms"
	DATA                         = "*data"
	VOICE                        = "*voice"
	MAX_COST_FREE                = "*free"
	MAX_COST_DISCONNECT          = "*disconnect"
	TOR                          = "tor"
	HOURS                        = "hours"
	MINUTES                      = "minutes"
	NANOSECONDS                  = "nanoseconds"
	SECONDS                      = "seconds"
	OUT                          = "*out"
	CDR_IMPORT                   = "cdr_import"
	CDR_EXPORT                   = "cdr_export"
	CDRFIELD                     = "cdrfield"
	ASR                          = "ASR"
	ACD                          = "ACD"
	FILTER_REGEXP_TPL            = "$1$2$3$4$5"
	ACTION_TIMING_PREFIX         = "apl_"
	RATING_PLAN_PREFIX           = "rpl_"
	RATING_PROFILE_PREFIX        = "rpf_"
	RP_ALIAS_PREFIX              = "ral_"
	ACC_ALIAS_PREFIX             = "aal_"
	ACTION_PREFIX                = "act_"
	SHARED_GROUP_PREFIX          = "shg_"
	ACCOUNT_PREFIX               = "ubl_"
	DESTINATION_PREFIX           = "dst_"
	LCR_PREFIX                   = "lcr_"
	DERIVEDCHARGERS_PREFIX       = "dcs_"
	TEMP_DESTINATION_PREFIX      = "tmp_"
	LOG_CALL_COST_PREFIX         = "cco_"
	LOG_ACTION_TIMMING_PREFIX    = "ltm_"
	LOG_ACTION_TRIGGER_PREFIX    = "ltr_"
	LOG_ERR                      = "ler_"
	LOG_CDR                      = "cdr_"
	LOG_MEDIATED_CDR             = "mcd_"
	SESSION_MANAGER_SOURCE       = "SMR"
	MEDIATOR_SOURCE              = "MED"
	CDRS_SOURCE                  = "CDRS"
	SCHED_SOURCE                 = "SCH"
	RATER_SOURCE                 = "RAT"
	CREATE_CDRS_TABLES_SQL       = "create_cdrs_tables.sql"
	CREATE_TARIFFPLAN_TABLES_SQL = "create_tariffplan_tables.sql"
	TEST_SQL                     = "TEST_SQL"
	CONSTANT                     = "constant"
	FILLER                       = "filler"
	METATAG                      = "metatag"
	HTTP_POST                    = "http_post"
	META_HTTP_POST               = "*http_post"
	META_HTTP_JSONRPC            = "*http_jsonrpc"
	NANO_MULTIPLIER              = 1000000000
	CGR_AUTHORIZE                = "CGR_AUTHORIZE"
	CONFIG_DIR                   = "/etc/cgrates/"
	CGR_SUPPLIER                 = "cgr_supplier"
	DISCONNECT_CAUSE             = "disconnect_cause"
	CGR_DISCONNECT_CAUSE         = "cgr_disconnect_cause"
)

var (
	CdreCdrFormats   = []string{CSV, DRYRUN, CDRE_FIXED_WIDTH}
	PrimaryCdrFields = []string{TOR, ACCID, CDRHOST, CDRSOURCE, REQTYPE, DIRECTION, TENANT, CATEGORY, ACCOUNT, SUBJECT, DESTINATION, SETUP_TIME, ANSWER_TIME, USAGE, SUPPLIER}
)
