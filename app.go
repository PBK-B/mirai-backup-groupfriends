package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strconv"

	"github.com/Logiase/MiraiGo-Template/bot"
	"github.com/Logiase/MiraiGo-Template/config"
	"github.com/Logiase/MiraiGo-Template/utils"
	// _ "github.com/Logiase/MiraiGo-Template/modules/getids"
	// _ "github.com/Logiase/MiraiGo-Template/modules/logging"
)

func init() {
	utils.WriteLogToFS(utils.LogInfoLevel, utils.LogWithStack)
	config.Init()
}

func main() {
	// 快速初始化
	bot.Init()

	// 初始化 Modules
	bot.StartService()

	// 使用协议
	// 不同协议可能会有部分功能无法使用
	// 在登陆前切换协议
	bot.UseProtocol(bot.AndroidPhone)

	// 登录
	// bot.Login()
	bot.QrcodeLogin()

	// 刷新好友列表，群列表
	bot.RefreshList()

	bot.SaveToken()

	fmt.Println("【启动服务】")

	// ginfo, err := bot.Instance.GetGroupInfo(949605396)

	// 获取群列表遍历备份
	glist, err := bot.Instance.GetGroupList()
	if err != nil {
		fmt.Println("【异常】群列表加载异常", err)
	}
	for _, ginfo := range glist {
		BackGroupUsers(ginfo.Code)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	bot.Stop()
}

// 备份群友联系方式列表
func BackGroupUsers(groupCode int64) {
	var txt string = ""
	ginfo, err := bot.Instance.GetGroupInfo(groupCode)
	if err != nil {
		fmt.Println("【异常】", err)
	}
	ids, err := bot.Instance.QQClient.GetGroupMembers(ginfo)
	if err != nil {
		fmt.Println("【异常】", err)
	}
	fmt.Println("【加载群】", ginfo.Code, ginfo.MemberCount, len(ginfo.Members))
	for _, uinfo := range ids {
		// fmt.Println(uinfo.Nickname, uinfo.Uin)
		txt = txt + uinfo.Nickname + " " + strconv.FormatInt(uinfo.Uin, 10) + "\n"
	}

	// 创建文件夹
	dirPath := "./group_users/"
	if _, err := os.Stat(dirPath); err != nil {
		os.Mkdir(dirPath, os.ModePerm)
	}

	// 写入文本
	WriterTxt(dirPath+strconv.FormatInt(ginfo.Code, 10)+"_"+ginfo.Name+".txt", txt)
	// fmt.Println(txt)
}

func WriterTxt(filePath string, str string) {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("【文件打开失败】", err)
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	write.WriteString(str)

	//Flush将缓存的文件真正写入到文件中
	write.Flush()
}
