package operations

import (
	"fmt"

	"github.com/josephniel/go-api/app/schema"
)

// GetUser is the sample operations method used by the sample controller
func GetUser(sample *schema.Sample) (result string, err error) {
	result = fmt.Sprintf("Name is %s and ID is %d", sample.Name, sample.ID)
	return
}
