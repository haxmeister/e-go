package query

// Query document comment
import "fmt"

//func Query(subcmd, option string) {

//	switch subcmd {
//	case "versions":
//		fmt.Println("querying versions of", option)

//	case "origin":
//		fmt.Println("querying origins of", option)
//	}
//}

func Versions(option string) {
	fmt.Println("querying versions of", option)
}
func Origin(option string) {
	fmt.Println("querying origin of", option)
}
