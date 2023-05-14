package mock

import (
	"github.com/soulteary/sparrow/internal/datatypes"
)

func Deactivate() datatypes.Deactivate {
	return datatypes.Deactivate{Status: "queued"}
}

func DataExport() datatypes.Deactivate {
	return datatypes.Deactivate{Status: "queued"}
}
