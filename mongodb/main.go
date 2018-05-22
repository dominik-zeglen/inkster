package mongodb

import (
	"github.com/dominik-zeglen/ecoknow/core"
)

// Adapter is an abstraction over mongodb session
type Adapter struct {
	core.Adapter

	ConnectionURI string
	DBName        string
}
