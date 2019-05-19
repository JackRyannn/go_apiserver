package service

import (
	"strconv"
)

var Step = 0
var Flag = false

func Foo(i string) string {
	var ret string
	ret = "null"

	switch i {
	case "-1":
		ret = "忽然天昏地暗，我晕了过去……\n" +
			"我醒了，又是凌晨5点。 我，或者说我的灵魂，深深叹了一口气，然后任由着躯体起床，穿衣，走出卧室。\n" +
			"“新郎官兴奋地睡不着啦。”客厅里几个通宵打牌的兄弟一起起哄。\n" +
			"同样没睡的还有我爹，他见我起来，拿起一张纸走了过来：”你表哥搞的几辆宝马过会就到，你也别睡了，再最后检查下流程安排吧。“\n" +
			"A.我有点累，先放那吧\n" +
			"B.好的我看一下"
	//	A
	case "0":
		ret = "音乐响起，司仪清了清嗓子：\n" +
			"下面，你可以亲吻自己美丽的新娘了。在这之前，你有什么话想对她说吗？（三思而后行）\n" +
			"A.真情表白\n" +
			"B.冷酷沉默\n" +
			"C.仓皇逃跑"
	//	A
	case "1":
		ret = "“亲爱的，这是我一生中最美好的时刻，我真希望时间能永远停留在这一天。”\n" +
			"“我也是。”\n" +
			"“新郎，你愿意和面前这位女士成为夫妻，共度此生吗？”\n" +
			"A.“我愿意”\n" +
			"B.“我……我能再想一想吗”\n" +
			"C.“我不愿意”"
	// B
	case "2":
		ret = "我醒了，2019年5月20日凌晨5点。 我，或者说我的灵魂，深深叹了一口气，然后任由着躯体起床，穿衣，走出卧室。\n" +
			"“新郎官兴奋地睡不着啦。”客厅里几个通宵打牌的兄弟一起起哄。\n" +
			"同样没睡的还有我爹，他见我起来，拿起一张纸走了过来：“你表哥搞的几辆宝马过会就到，你也别睡了，再最后检查下流程安排吧。”\n" +
			"A.我有点累，先放那吧\n" +
			"B.好的我看一下"
	// B
	case "3":
		ret = "6:00 车队集合\n" +
			"6:58 出发接亲\n" +
			"7:30 抵达新娘家\n" +
			"…………\n" +
			"这已经是我第1351次看这张流程单了，接下来的24小时内，我还要再看4次。\n" +
			"没错，我被困在婚礼这一天了。今天，是我第271次结婚。\n" +
			"A.我受够了，我今天不想再受安排\n" +
			"B.出发吧，人生本就是无止境的轮回"
	//B
	case "4":
		ret = "我边走边想着今天的梦。做梦几乎是我最近一段时间唯一的乐趣，因为只有在梦境里，我才能过上每天不一样的生活。\n" +
			"是的，梦境里才能叫生活。而每天重复的那场婚礼，更像是一部电影，而我被牢牢困在座位上，撑开双眼，动弹不得，有如《发条橙》里的场景。" +
			"车到了，我要上车了\n" +
			"A.打开奔驰车门\n" +
			"B.打开宝马车门\n" +
			"C.打开加长林肯的车门"
	//	A
	case "5":
		ret = "我坐上了花车，想到每天都要重复这句话我就很难受，“亲爱的，这是我一生中最美好的时刻，我真希望时间能永远停留在这一天。”\n" +
			"如果事情全是因为这句话而起，那等我跳出时间漩涡之后，第一时间就会去宰了司仪。困了多少天，我就会捅他多少刀，每捅一刀，我都会问他一句：\n" +
			"A.直接亲不好吗！\n" +
			"B.还有272刀\n" +
			"C.司仪你是不是给我带了绿帽子"
	//	B
	case "6":
		ret = "其实在事情刚发生的时候，我还感觉很有趣，毕竟看过不少时间循环的电影。何况还是被困在自己的婚礼而不是葬礼上，总是让人开心的。\n" +
			"但很快我就厌烦了。我说的每句话，做的每件事，就像那张我看了上千次的流程单一样，按部就班，一成不变。我开始讨厌每一个环节，讨厌每一个人。\n" +
			"他们的笑容看上去越来越虚伪敷衍，就像现在在我面前的，那些抢到红包欢呼雀跃的娘家亲戚，和堵着门不让我进的伴娘团。除此之外我还发现了一个事实\n" +
			"A.伴郎小瑶是一位非常优秀的男生\n" +
			"B.伴娘小珺的胸非常的大"
	//	C
	case "7":
		ret = "当我发现没办法挣脱时间漩涡的时候，我开始给自己找乐子。我给这场婚礼定名为《复愁者》，然后在重复重复再重复的忧愁中苦中作乐。无法控制身体的我，竟然从每天重复的视觉信息中得到了很多有趣的彩蛋。比如司仪总是喜欢在没人注意他的时候挖鼻屎。比如摄像师总是偷拍小霞的胸。\n" +
			"这时候伴娘小李举着杯子就过来了，他是我学生时代最欣赏的女生，那时我们有好多好多共同话题，我们也一同憧憬着未来，可是后来我们因为去了不同的城市，渐渐失去了交集。如果，我是说如果，哎算了不想了\n" +
			"“新郎官，敬我们的过去，你们的未来”\n" +
			"A.谢谢，你知道吗？我很想你\n" +
			"B.谢谢，过去就过去了，敬未来就够了\n" +
			"C.谢谢，为了我们这份友谊，我干了你随意\n"
	//	A
	case "8":
		ret = "不过这都是些琐碎的事情，又过了三个月，我已经背下了那天我看到的所有车的车牌，酒席所有宾客的位置，甚至连他们衣服的颜色也都烂熟于胸。\n" +
			"再后来我就开始思考人生，我问自己，“为什么要结婚呢？”，苦思冥想了几个月我放弃了，我想不想结婚重要吗？反正我每天都必须结婚，我真的快崩溃了。\n" +
			"“先生，你可以亲吻自己美丽的新娘了。在这之前，你有什么话想对她说吗？”\n" +
			"A.真情表白\n" +
			"B.冷酷沉默\n" +
			"C.仓皇逃跑"
	//	A
	case "9":
		ret = "司仪的话把我从神游中拉了回来，我的身体，第271次出现在婚礼舞台的正中间。\n" +
			"“亲爱的，这是我一生中最美好的时刻，我真希望时间能永远停留在这一天。”\n" +
			"“我也是。”\n" +
			"等等，这个声音为什么听上去不大对劲？她的眼里噙着泪水，但为什么既看不出感动也看不出高兴，相反，却透着一丝绝望,我的脑袋顿时炸了，难道？\n" +
			"A.她跟我一样被困在这一天？\n" +
			"B.她该喝热水了吗？"
	//	B
	case "10":
		ret = "一阵天旋地转，我晕倒了。\n" +
			"醒来后几秒钟之后我便记起了一切。婚礼没过多久，我跟她发生了激烈的争吵，闹到了离婚的地步。为了挽救这段婚姻，我和她选择了一款时间漩涡的情感分析系统。我们被清除了接受治疗的记忆，送进婚礼24小时的梦境里循环。想要跳出很简单，就是发现对方状态和自己一样，也是被困在其中的。但一个人发现是没有用的，必须两个人都发现的那一刻，才能同时跳出。\n" +
			"我挣扎着起身，问医生：\n" +
			"A.大夫，小珺还好吗？\n" +
			"B.大夫，我的妻子呢？"
	//	C
	case "11":
		ret = "他用手一指，我的妻子，正靠在门边，面色冰冷地看着我。\n" +
			"“我第三天就发现你不对了。”她说。\n" +
			"“那我再问你一次，接下来，你要跪在美丽的新娘面前。在这之前，你有什么话想对她说吗？”\n" +
			"A.老婆你可真牛逼，是真爱没错了\n" +
			"B.既然你都看到了，那我也没什么好说的了\n" +
			"C.520，以后家务都归我干了"
	case "12":
		ret = "“最好说话算话哦，我可不能白白等你这么久。”\n" +
			"“现在几点了呀老婆”\n" +
			"“下午五点多了。" +
			"“不对，现在是我们幸福的起点”\n" +
			"“你别恶心我”\n" +
			"--游戏结束--"

	}

	return ret
}
func Receive(c string) string {
	if c == "start" {
		Flag = true
		return Foo("0")
	}
	if c == "exit" {
		Flag = false
		Step = 0
		return "游戏结束"
	}
	answer := "AABBBABCAABC"
	if Step == len(answer)+1 || (c != "A" && c != "B" && c != "C") || !Flag {
		Step = 0
		return "开始游戏请输入start"
	}
	if string(answer[Step]) == c {
		Step += 1
		r := Foo(strconv.Itoa(Step))
		return r

	} else {
		Step = 2
		return Foo("-1")
	}

}
