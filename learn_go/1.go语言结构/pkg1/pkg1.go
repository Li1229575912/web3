package pkg1

import (
	"fmt"

	_ "github.com/learn/init_order/pkg2"
)

const PkgName string = "package1"

var PkgNameVar string = getPkgName()

func init() {
	fmt.Println("package1 init method invoked.")
}

func getPkgName() string {
	fmt.Println("pkg1.PkgNameVar has been initialized.")
	return PkgName
}
