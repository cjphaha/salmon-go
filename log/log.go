package log

import (
	"io"
	"log"
	"os"
	"time"
)

//NewLoggerWidthFile 保存到文件中的日志
func NewLoggerWidthFile(path string) (logger *log.Logger){
	logFile,err:=os.OpenFile(path,os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	if err!=nil {
		log.Fatalln("打开日志文件失败：", err)
	}
	logger = log.New(io.MultiWriter(os.Stderr,logFile),"Log:",log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("系统启动，开始记录日志",time.Now())
	return
}
