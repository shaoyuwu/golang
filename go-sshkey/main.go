package main

import (
	"fmt"
	"io/ioutil"
	"net"

	"github.com/mitchellh/go-homedir"
	"golang.org/x/crypto/ssh"
)

var keypath = "~/id_rsa"

//获取秘钥

func publicKey(path string) ssh.AuthMethod {
	keypath, err := homedir.Expand(path)
	if err != nil {
		fmt.Println("获取秘钥路径失败", err)
	}
	key, err1 := ioutil.ReadFile(keypath)
	if err1 != nil {
		fmt.Println("读取秘钥失败", err1)
	}
	signer, err2 := ssh.ParsePrivateKey(key)
	if err2 != nil {
		fmt.Println("ssh 秘钥签名失败", err2)
	}
	return ssh.PublicKeys(signer)
}

//获取ssh连接
func GetSSHConect(ip, user string, port int) *ssh.Client {
	con := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{publicKey(keypath)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	addr := fmt.Sprintf("%s:%d", ip, port)

	client, err := ssh.Dial("tcp", addr, con)

	if err != nil {

		fmt.Println("Dail failed: ", err)

		panic(err)

	}

	return client

}

// 远程执行脚本

func Exec_Task(ip, user, localpath, remotepath string) int {

	port := 22

	client := GetSSHConect(ip, user, port)

	// UploadFile(ip, user, localpath, remotepath, port)

	session, err := client.NewSession()

	if err != nil {

		fmt.Println("创建会话失败", err)

		panic(err)

	}

	defer session.Close()

	b, _ := session.CombinedOutput("df -h")
	fmt.Println(string(b))
	return 1

	// remoteFileName := path.Base(localpath)
	//
	// dstFile := path.Join(remotepath, remoteFileName)

	// err1 := session.Run(fmt.Sprintf("/usr/bin/sh %s", dstFile))

	// if err1 != nil {

	// 	fmt.Println("远程执行脚本失败", err1)

	// 	return 2

	// } else {

	// 	fmt.Println("远程执行脚本成功")

	// 	return 1

	// }
}

func main() {
	Exec_Task("192.168.158.160", "root", "/root/", "/root")
}
