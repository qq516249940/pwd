package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

var prompt = "Enter a radius and an angle (in degrees), e.g., 12.5 90, " +
	"or %s to quit."

func init() {
	if runtime.GOOS == "window" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else {
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
}

func uploadPath() {

	sep := os.PathSeparator

	// 查看当前平台的系统路径分隔符,windows平台是\
	fmt.Println(string(sep)) // \

	// 将分割符转为/
	fmt.Println(filepath.ToSlash(`C:\python37\python.exe`)) // C:/python37/python.exe

	// 注意：该函数不在意路径是否存在，只是当成普通的字符串进行处理
	// 比如我输入一个不存在的路径也是可以的
	fmt.Println(filepath.ToSlash(`C:\python37\python.exe\python.exe`))

	fmt.Println(filepath.FromSlash(`C:/python37/python.exe`)) // C:\python37\python.exe
	fmt.Println(filepath.FromSlash(`C:\python37/python.exe`)) // C:\python37\python.exe
	fmt.Println(filepath.FromSlash(`C:\python37\python.exe`)) // C:\python37\python.exe\\\\

}

func osInfo() {
	fileInfo, err := os.Stat("file")
	if err != nil {
		ftm.Printf(err)
	} else {
		is_dir := fileInfo.IsDir()
	        fmt.Printf("os.FileInfo.IsDir(): %v\n", is_dir)
	}

}

func main() {
	fmt.Println(prompt)
	fmt.Printf("runtime.GOROOT(): %v\n", runtime.GOROOT())
	uploadPath()
	osInfo()

}
