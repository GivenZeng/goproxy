// 一个本地文件服务器
package main

type FileSrv struct {
	root string
}

// path为相对root的路径
func (fs *FileSrv) getAbsolutePath(path string) string {
	return ""
}

// 判断path是否为目录
// path为相对root的路径
func (fs *FileSrv) isDir(path string) bool {
	return false
}

// path为相对root的路径
// 获取目录的信息
func (fs *FileSrv) getDirMsg(path string) {

}

// filePath 是相对root的路径
// example：filPath=a/b，那么就会返回$root/a/b的内容
func (fs *FileSrv) getFileContent(filePath string) []byte {
	return nil
}

// 获取文件修改时间
func (fs *FileSrv) getFileModeTime(filePath string) []byte {
	return nil
}
