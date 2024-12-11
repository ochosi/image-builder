package v1

import (
	"github.com/osbuild/image-builder/internal/db"
)

var (
	BuildOSTreeOptions     = buildOSTreeOptions
	ValidateComposeRequest = validateComposeRequest
)

func (s *Server) GetDB() db.DB {
	return s.db
}