package util

import (
	"github.com/google/uuid"
)

type Id struct{}

/**
* UUIDをString形式で返します。
 */
func (i *Id) Uuid() string {
	id := uuid.New()
	return id.String()
}
