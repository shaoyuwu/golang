package main

import (
	"io/ioutil"

	"github.com/CodyGuo/dingtalk"
	"github.com/CodyGuo/dingtalk/pkg/robot"
	"github.com/CodyGuo/glog"
)

func main() {
	glog.SetFlags(glog.LglogFlags)
	webHook := "https://oapi.dingtalk.com/robot/send?access_token=xxxxxxxxxxxx"
	// **，机器人安全设置页面，加签一栏勾选之后下面显示的SEC开头的字符串
	secret := "xxxxxxxxxx"
	dt := dingtalk.New(webHook, dingtalk.WithSecret(secret))
	// dt := dingtalk.New(webHook)
	// fmt.Printf("webHook: %v\n", webHook)

	// text类型
	textContent := "我就是我是不一样的烟火@176XXXXXXXX"
	atMobiles := robot.SendWithAtMobiles([]string{"@15071244227"})
	// atMobiles := robot.SendWithAtMobiles([]string{"@176XXXXXXXX", "178xxxxxx28"})
	if err := dt.RobotSendText(textContent, atMobiles); err != nil {
		glog.Fatal(err)
	}
	printResult(dt)

	// link类型
	linkTitle := "时代的火车向前开"
	linkText := `这个即将发布的新版本，创始人xx称它为“红树林”。` +
		`而在此之前，每当面临重大升级，产品经理们都会取一个应景的代号，` +
		`这一次，为什么是“红树林”？`
	linkMessageURL := "https://www.dingtalk.com/s?__biz=MzA4NjMwMTA2Ng==&mid=2650316842&idx=1&sn=60da3ea2b29f1dcc43a7c8e4a7c97a16&scene=2&srcid=09189AnRJEdIiWVaKltFzNTw&from=timeline&isappinstalled=0&key=&ascene=2&uin=&devicetype=android-23&version=26031933&nettype=WIFI"
	linkPicURL := "https://cdn.pixabay.com/photo/2020/05/05/08/05/butterfly-5131967_960_720.jpg"
	if err := dt.RobotSendLink(linkTitle, linkText, linkMessageURL, linkPicURL); err != nil {
		glog.Fatal(err)
	}
	printResult(dt)

	// markdown类型
	markdownTitle := "markdown"
	markdownText := "#### 杭州天气 @176XXXXXXXX\n" +
		"> 9度，西北风1级，空气良89，相对温度73%\n" +
		"> ![screenshot](https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png)\n" +
		"> ###### 10点20分发布 [天气](https://www.dingtalk.com)\n"
	if err := dt.RobotSendMarkdown(markdownTitle, markdownText); err != nil {
		glog.Fatal(err)
	}
	printResult(dt)

	// 整体跳转ActionCard类型
	actionCardTitle := "乔布斯 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身"
	actionCardText := "![screenshot](@lADOpwk3K80C0M0FoA)\n" +
		"### 乔布斯 20 年前想打造的苹果咖啡厅\n" +
		"Apple Store 的设计正从原来满满的科技感走向生活化，" +
		"而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划"
	actionCardSingleTitle := "阅读全文"
	actionCardSingleURL := "https://www.dingtalk.com/"
	actionCardBtnOrientation := "0"
	if err := dt.RobotSendEntiretyActionCard(actionCardTitle,
		actionCardText,
		actionCardSingleTitle,
		actionCardSingleURL,
		actionCardBtnOrientation); err != nil {
		glog.Fatal(err)
	}
	printResult(dt)

	// 独立跳转ActionCard类型
	btns := map[string]string{
		"内容不错": actionCardSingleURL,
		"不感兴趣": actionCardSingleURL,
	}
	if err := dt.RobotSendIndependentActionCard(actionCardTitle,
		actionCardText,
		actionCardBtnOrientation,
		btns); err != nil {
		glog.Fatal(err)
	}
	printResult(dt)

	// FeedCard类型
	link1 := robot.FeedCardLink{
		Title:      linkTitle,
		MessageURL: linkMessageURL,
		PicURL:     linkPicURL,
	}
	link2 := robot.FeedCardLink{
		Title:      linkTitle + "2",
		MessageURL: linkMessageURL,
		PicURL:     linkPicURL,
	}
	links := []robot.FeedCardLink{link1, link2}
	if err := dt.RobotSendFeedCard(links); err != nil {
		glog.Fatal(err)
	}
	printResult(dt)
}

func printResult(dt *dingtalk.DingTalk) {
	response, err := dt.GetResponse()
	if err != nil {
		glog.Fatal(err)
	}
	reqBody, err := response.Request.GetBody()
	if err != nil {
		glog.Fatal(err)
	}
	reqData, err := ioutil.ReadAll(reqBody)
	if err != nil {
		glog.Fatal(err)
	}
	glog.Infof("发送消息成功, message: %s", reqData)
}
