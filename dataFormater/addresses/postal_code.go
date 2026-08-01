package addresses

import (
	"fmt"

	"github.com/ViitoJooj/go-sdk/internal"
)

func PostalCode(postal string) string {
	digits := internal.StripNonDigits(postal)
	if len(digits) == 8 {
		return fmt.Sprintf("%s-%s", digits[0:5], digits[5:8])
	}
	return postal
}
