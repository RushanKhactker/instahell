package herror

import "fmt"
// во имя Сантаны!!!!!!1111
type HError struct {
	Hell string
	Err error
}

func (x *HError) Error() string {
	return fmt.Sprintf(
		"хеггог: х %s, еггог: %v",
		x.Hell,
		x.Err,
	)
}
