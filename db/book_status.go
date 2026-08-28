package db

type BookStatus int

const (
	Available BookStatus = iota // 0
	Swapped                     // 1
)

func (o BookStatus) String() string {
	return [...]string{"AVAILABLE", "SWAPPED"}[o]
}

// func (o BookStatus) String() string {
//     if int(o) < 0 || int(o) >= len(statusStrings) {
//         return "UNKNOWN"
//     }
//     return statusStrings[o]
// }

// func (o BookStatus) String() string {
//     switch o {
//     case AVAILABLE:
//         return "AVAILABLE"
//     case SWAPPED:
//         return "SWAPPED"
//     default:
//         return fmt.Sprintf("UNKNOWN(%d)", o)
//     }
// }
