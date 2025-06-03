package go_test

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestPath(t *testing.T) {
	//获取当前源文件的路径
	_, file, _, _ := runtime.Caller(0)
	fmt.Println(file)

	//获取当前源文件的路径
	abs, _ := filepath.Abs(os.Args[0])
	fmt.Println("Abs :\n" + abs)

	dir := filepath.Dir(abs)
	fmt.Println("Bin Dir: \n" + dir)

	volumeName := filepath.VolumeName(abs)
	fmt.Println("volumeName : \n" + volumeName)

	base := filepath.Base(os.Args[0])
	fmt.Println("Program Name: \n" + base)

	dir, file = filepath.Split(abs)
	fmt.Println(dir)
	fmt.Println(file)

	confFilePath := filepath.Join(dir, "conf.yaml")
	fmt.Println("confFilePath: \n" + confFilePath)
}
