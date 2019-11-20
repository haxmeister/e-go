package sync

import "fmt"

// Meta function only syncs the meta repo and not the kits repo
func Meta() {
	fmt.Println("syncing meta repo only")
}

// Kits function only syncs kits and not the meta repo
func Kits() {
	fmt.Println("syncing kits only")
}
