package main

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	fmt.Println("请输入要修改的文件夹绝对地址:")
	var dir string
	fmt.Scanf("%s", &dir)

	files, _ := ioutil.ReadDir(dir)

	for _, i := range files {
		read(i,dir)

	}
}

func read(f fs.FileInfo,dir string) {
	if !f.IsDir() {
		// 读取文件名
		name := f.Name()
		// 反转数组，判断是否为java文件
		r := []rune(name)
		for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
		reverseName := string(r)
		if len(reverseName) > 5 && reverseName[:5] == "avaj." {
			fmt.Printf("文件名为:%s\n", name)

			// 读取文件，并写入
			fi, err := os.Open(dir+"/"+f.Name())
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				return
			}
			defer fi.Close()

			br := bufio.NewReader(fi)
			for {
				a, _, c := br.ReadLine()
				if c == io.EOF {
					break
				}
				tracefile(string(a))
				fmt.Println(string(a))

			}

		}

		return
	}

	newDir := dir + "/" + f.Name()
	files, err := ioutil.ReadDir(newDir)
	if err != nil {
		fmt.Println(err)
	}
	for _, i := range files {
		read(i, newDir)
	}
}


func tracefile(str_content string)  {
	fd,_:=os.OpenFile("a.txt",os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	fd_content:=strings.Join([]string{str_content,"\n"},"")
	buf:=[]byte(fd_content)
	fd.Write(buf)
	fd.Close()
}