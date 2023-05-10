package mock

import "github.com/soulteary/sparrow/internal/datatypes"

func Deactivate() datatypes.Deactivate {
	data := datatypes.Deactivate{
		Status: "queued",
	}
	return data
}

func DataExport() datatypes.Deactivate {
	data := datatypes.Deactivate{
		Status: "queued",
	}
	return data
}
