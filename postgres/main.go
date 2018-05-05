package postgres

import (
	"github.com/dominik-zeglen/ecoknow/core"
	"github.com/go-pg/pg"
)

// Adapter is an abstraction over pg session
type Adapter struct {
	core.Adapter

	ConnectionOptions pg.Options
}
