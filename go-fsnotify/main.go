package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/CodyGuo/dingtalk"
	"github.com/fsnotify/fsnotify"
)

var (
	title = "Canal实例异常"
	base  = "/data/canal/logs"
)

func InitWebhookAndSecret() *dingtalk.DingTalk {
	var dingtalkwebHook string
	dingtalkwebHook = "https://oapi.dingtalk.com/robot/send?access_token=ae97a3e3a630bf40d7cbdfbc576bf8b253856489d556c4120be19576370dc0e7"
	dt := dingtalk.New(dingtalkwebHook)
	return dt
}

func sendDingtalkMsg(dt *dingtalk.DingTalk, ip, hostname, instance string) {
	markdownTitle := "markdown"
	markdownText := fmt.Sprintf("### %s\n"+
		"- IP: %s\n"+
		"- 主机名: %s\n"+
		"- canal实例: %s", title, ip, hostname, instance)

	if err := dt.RobotSendMarkdown(markdownTitle, markdownText); err != nil {
		log.Println(err)
	}
}

func getHostnameCmd() string {
	getHostnameCmd := fmt.Sprintf("echo %s", "$HOSTNAME")
	return execCmd(getHostnameCmd)
}

func getIpCmd() string {
	getIpCmd := fmt.Sprintf("ip a|grep global|awk  '{print $2}'|awk -F/ '{print $1}'")
	return execCmd(getIpCmd)
}

func execCmd(cmd string) string {
	c := exec.Command("bash", "-c", cmd)
	result, err := c.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	return string(result)
}

func watchFiles(watcher *fsnotify.Watcher, instance string) {
	hostname := getHostnameCmd()
	ip := getIpCmd()
	dt := InitWebhookAndSecret()

	fullPathInstance := base + "/" + instance + "/" + instance + ".log"
	fmt.Printf("instance: %v\n", fullPathInstance)
	err := watcher.Add(fullPathInstance)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Write == fsnotify.Write {
				sendDingtalkMsg(dt, ip, hostname, instance)
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("error:", err)
		}
	}
}

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	done := make(chan bool)

	for _, instance := range os.Args[1:] {
		go watchFiles(watcher, instance)
	}

	<-done
}
