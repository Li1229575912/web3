package pkg2

import "fmt"

const PkgName string = "package2"

var PkgNameVar string = getPkgName()

func init() {
	fmt.Println("package2 init method invoked.")
}

func getPkgName() string {
	fmt.Println("package2.PkgName has initialized.")
	return PkgName
}
