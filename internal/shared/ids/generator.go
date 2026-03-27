package ids

import (
	"strings"

	"github.com/google/uuid"
)

func GenerateShortID() string {
	u := uuid.New().String()
	u = strings.ReplaceAll(u, "-", "")
	return u[:12]
}
