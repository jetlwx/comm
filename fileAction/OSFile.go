package fileAction

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/jetlwx/comm"
)

type FileORDirPath struct {
	Path string
}

//得到指定目录下所有文件或目录列表(不递归)),
func (fdp FileORDirPath) GetFileList(isDir bool) (list []string) {

	files, _ := ioutil.ReadDir(fdp.Path)
	for _, f := range files {
		if f.IsDir() && isDir {
			list = append(list, f.Name())
			continue
		}

		if !f.IsDir() && !isDir {
			list = append(list, f.Name())
			continue
		}
	}
	return list
}

//得到指定目录下所有文件或目录列表(不递归)),
func GetFileList(filepath string, isDir bool) (list []string) {

	files, _ := ioutil.ReadDir(filepath)
	for _, f := range files {
		if f.IsDir() && isDir {
			list = append(list, f.Name())
			continue
		}

		if !f.IsDir() && !isDir {
			list = append(list, f.Name())
			continue
		}
	}
	comm.JetLog("D", list)
	return list
}

//得到指定目录文件列表(含子目录)，同时指定后续，若不指定，则为全部文件,Ingnor:是否忽略后缀匹配的大小写
func GetAllFiles(dirPth string, stuffix string, IgnorCase bool) (files []string, err error) {
	var dirs []string
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)
	if IgnorCase {
		stuffix = strings.ToUpper(stuffix) //忽略后缀匹配的大小写
	}

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			GetAllFiles(dirPth+PthSep+fi.Name(), stuffix, IgnorCase)
		} else {
			// 过滤指定格式
			var ff string
			if IgnorCase {
				ff = strings.ToUpper(fi.Name())
			} else {
				ff = fi.Name()
			}
			ok := strings.HasSuffix(ff, stuffix)
			if ok {
				if dirPth+PthSep+fi.Name() != "" {
					files = append(files, dirPth+PthSep+fi.Name())
				}
			}
		}
	}

	// 读取子目录下文件
	for _, table := range dirs {
		temp, _ := GetAllFiles(table, stuffix, IgnorCase)
		for _, temp1 := range temp {
			if temp1 != "" {
				files = append(files, temp1)
			}
		}
	}

	return files, nil
}

//文件是否存在,返回nil说明文件存在
func (fdp FileORDirPath) PathExists() error {
	_, err := os.Stat(fdp.Path)
	if err == nil {
		return nil
	}
	return err
}

func PathExists(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	return err
}

//文件类型是否文本
func (fdp FileORDirPath) FileType() bool {
	cmd := "/usr/bin/file " + " " + fdp.Path
	if PathExists("/usr/bin/file") != nil {
		log.Println("[ E ]", "/usr/bin/file 文件 不存在，请先安装")
		return false
	}
	okstr, _ := comm.ExecOSCmd(cmd, false, 0)
	if strings.Contains(okstr, "text") {
		return true
	}
	return false
}

//文件类型是否文本
func FileType(filepath string) bool {
	cmd := "/usr/bin/file " + " " + filepath
	if PathExists("/usr/bin/file") != nil {
		log.Println("[ E ]", "/usr/bin/file 文件 不存在，请先安装")
		return false
	}
	okstr, _ := comm.ExecOSCmd(cmd, false, 0)
	if strings.Contains(okstr, "text") {
		return true
	}
	return false
}

//从路径中获取文件名
func (fdp FileORDirPath) FileNameInpath() (filedir, filename string) {
	return filepath.Split(fdp.Path)

}

//从路径中获取文件名
func FileNameInpath(filedirandname string) (filedir, filename string) {
	return filepath.Split(filedirandname)

}

//AppendToFileEnd  追加内容到文件末尾 fileName:文件名字(带全路径) content: 写入的内容
func (fdp FileORDirPath) AppendToFileEnd(content string) error {
	// 以只写的模式，打开文件
	f, err := os.OpenFile(fdp.Path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	// 查找文件末尾的偏移量
	n, _ := f.Seek(0, os.SEEK_END)
	// 从末尾的偏移量开始写入内容
	_, err = f.WriteAt([]byte(content), n)

	return err
}

//AppendToFileEnd  追加内容到文件末尾 fileName:文件名字(带全路径) content: 写入的内容
func AppendToFileEnd(filename, content string) error {
	// 以只写的模式，打开文件
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	// 查找文件末尾的偏移量
	n, _ := f.Seek(0, os.SEEK_END)
	// 从末尾的偏移量开始写入内容
	_, err = f.WriteAt([]byte(content), n)

	return err
}
