package test

import (
	"fmt"
	"testing"

	"github.com/shirou/gopsutil/host"
)

/* func TestFuncVal(t *testing.T) {

}

func TestGetInfo(t *testing.T) {
	ar := runtime.GOARCH
	info := runtime.GOOS

	fmt.Printf("ar: %v\n", ar)
	fmt.Printf("info: %v\n", info)
} */

func TestGetSysInof(t *testing.T) {
	platform, family, versiosn, _ := host.PlatformInformation()

	fmt.Printf("platform: %v\n", platform)
	fmt.Printf("family: %v\n", family)
	fmt.Printf("versiosn: %v\n", versiosn)

}
