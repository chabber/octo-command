package domain

const MAX_BED_TEMPERATURE float64 = 110.0
const MAX_TOOL_TEMPERATURE float64 = 285.0

const (
	SERVER_PROFILE_COLLECTION string = ".store/server_profiles"
	TEMP_PROFILE_COLLECTION   string = ".store/temp_profiles"
	CONFIG_COLLECTION         string = ".store"
	CONFIG_RESOURCE           string = "config"
)

type PrinterCommType string

const (
	OCTO_PRINT    PrinterCommType = "OCTO_PRINT"
	DIRECT_MARLIN PrinterCommType = "DIRECT_MARLIN"
)
