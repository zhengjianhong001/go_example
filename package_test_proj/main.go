package main

import (
	"package_test_proj/mypkg"
	"package_test_proj2/mypkg2"
)

func main() {
	mypkg.Hello()
	mypkg.GetName()
	mypkg.GetNum()
	mypkg2.GetNum()
}
