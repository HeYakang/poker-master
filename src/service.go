package src

import (
	. "poker/define"
)
//初步简单处理得到的数据类型
//手牌信息
type Hand struct {
	HandStr  string    // 记录原始手牌字符串
	GhostNum uint64    // 鬼牌数量
	Suits    [4]uint64 // 记录手牌中出现过得所有牌的花色
	Faces    [4]uint64 // 记录手牌中出现过得所有牌的出现的次数（数组下标加1即为出现次数，bit位记录手牌牌面）
}


//程序处理后得到的结果数据结构
type MaxHand struct {
	MaxCase uint64  //记录最大牌型（StraightFlush, FourOfAKind, FullHouse...）
	MaxHand uint64  //五张牌的排序 （bit 为牌型顺序，因为重要的牌放高位，所有int值可以直接比较得出胜利方）
	FlushFlag bool  //同花标志位
	FlushSuit int   //如果有同花，记录同花的花色编号
}

// 获取获胜者编号
func getWinnerCompare(a, b uint64) int {
	return CaseWhen(a == b, 0, a > b, 1, a < b, 2).(int)
}

//获取胜利者编号第二种方法
func getWinner(strA string, strB string) int{
	handA := analysisHandStr(strA)
	handB := analysisHandStr(strB)
	maxHandA := MaxHand{}
	maxHandB := MaxHand{}

	var result int
	if result = util(maxHandA.isStraightFlush(handA),maxHandB.isStraightFlush(handB)); result>=0 {
	}else if result =util(maxHandA.isFourOfAKind(handA),maxHandB.isFourOfAKind(handB)); result>=0 {
	}else if result =util(maxHandA.isFullHouse(handA),maxHandB.isFullHouse(handB)); result>=0 {
	}else if result =util(maxHandA.isFlush(handA),maxHandB.isFlush(handB)); result>=0 {
	}else if result =util(maxHandA.isStraight(handA),maxHandB.isStraight(handB)); result>=0 {
	}else if result =util(maxHandA.isThreeOfAKind(handA),maxHandB.isThreeOfAKind(handB)); result>=0 {
	}else if result =util(handA.GhostNum ==0 &&maxHandA.isTwoPair(handA),handB.GhostNum ==0 &&maxHandB.isTwoPair(handB)); result>=0 {
	}else if result =util(maxHandA.isOnePair(handA),maxHandB.isOnePair(handB)); result>=0 {
	}else if result =util(maxHandA.isHighCard(handA),maxHandB.isHighCard(handB)); result>=0 {
	}
	if result==0 {
		if maxHandA.MaxCase !=StraightFlush &&maxHandA.MaxCase !=Straight {
			maxHandA.getmaxhand(handA)
			maxHandB.getmaxhand(handB)
		}else {
			//比较牌型
			// 顺子&同花顺存在“A2345”这一特殊情况，此时为最小顺子，需要手动标记（权值score设为0）
			scoreA := If(maxHandA.MaxHand == A2345, uint64(0), maxHandA.MaxHand).(uint64)
			scoreB := If(maxHandB.MaxHand == A2345, uint64(0), maxHandB.MaxHand).(uint64)
			return getWinnerCompare(scoreA, scoreB)
		}

		return getWinnerCompare(maxHandA.MaxHand, maxHandB.MaxHand)
	}
	return result
}

//func compare(a,b *Hand) int{
//	maxHandA := MaxHand{}
//	maxHandB := MaxHand{}
//
//	var result int
//	if result = util(maxHandA.isStraightFlush(a),maxHandB.isStraightFlush(b)); result>=0 {
//	}else if result =util(maxHandA.isFourOfAKind(a),maxHandB.isFourOfAKind(b)); result>=0 {
//	}else if result =util(maxHandA.isFullHouse(a),maxHandB.isFullHouse(b)); result>=0 {
//	}else if result =util(maxHandA.isFlush(a),maxHandB.isFlush(b)); result>=0 {
//	}else if result =util(maxHandA.isStraight(a),maxHandB.isStraight(b)); result>=0 {
//	}else if result =util(maxHandA.isThreeOfAKind(a),maxHandB.isThreeOfAKind(b)); result>=0 {
//	}else if result =util(a.GhostNum ==0 &&maxHandA.isTwoPair(a),b.GhostNum ==0 &&maxHandB.isTwoPair(b)); result>=0 {
//	}else if result =util(maxHandA.isOnePair(a),maxHandB.isOnePair(b)); result>=0 {
//	}else if result =util(maxHandA.isHighCard(a),maxHandB.isHighCard(b)); result>=0 {
//	}
//	if result==0 {
//		if maxHandA.MaxCase !=StraightFlush &&maxHandA.MaxCase !=Straight {
//			maxHandA.getmaxhand(a)
//			maxHandB.getmaxhand(b)
//		}else {
//			//比较牌型
//			// 顺子&同花顺存在“A2345”这一特殊情况，此时为最小顺子，需要手动标记（权值score设为0）
//			scoreA := If(maxHandA.MaxHand == A2345, uint64(0), maxHandA.MaxHand).(uint64)
//			scoreB := If(maxHandB.MaxHand == A2345, uint64(0), maxHandB.MaxHand).(uint64)
//			return getWinnerCompare(scoreA, scoreB)
//		}
//
//		return getWinnerCompare(maxHandA.MaxHand, maxHandB.MaxHand)
//	}
//	return result
//}

func util(a,b bool) int{
	if !a && !b {
		return -1
	}else if a==true && b==false  {
		return 1
	}else if a==false && b==true  {
		return 2
	}else  {
		return 0
	}
}



//处理解析后的手牌
func (hand *Hand) getMaxHand() *MaxHand{
	maxHand := MaxHand{}

	//按照 同花顺 四条这顺序判断最大牌型（皇家同花顺作为同花顺特殊牌型，不另外做讨论）
	if maxHand.isStraightFlush(hand) {
	} else if maxHand.isFourOfAKind(hand) {
	} else if maxHand.isFullHouse(hand) {
	} else if maxHand.isFlush(hand) {
	} else if maxHand.isStraight(hand) {
	} else if maxHand.isThreeOfAKind(hand) {
	} else if hand.GhostNum ==0 && maxHand.isTwoPair(hand) {
	} else if maxHand.isOnePair(hand) {
	} else if maxHand.isHighCard(hand) {
	}
	return  &maxHand
}


//解析手牌字符串输出牌型
func analysisHandStr(handStr  string) *Hand{
	hand := Hand{HandStr: handStr}
	var faceValue uint64  //牌的数值
	for i :=0; i<len(handStr); i++{
		//遇到X为Ghost牌，跳过并且记录
		if handStr[i] =='X'{
			hand.GhostNum++
			i++
			continue
		}

		if i%2 == 0 {
			//根据出现的字符去查找值 并且记录下来
			faceValue = Faces[handStr[i]]
			//这个牌面出现4次
			hand.Faces[3] |= hand.Faces[2] & faceValue
			//这个牌面出现3次
			hand.Faces[2] |= hand.Faces[1] & faceValue
			//这个牌面出现2次
			hand.Faces[1] |= hand.Faces[0] & faceValue
			//这个牌面出现1次
			hand.Faces[0] |=  faceValue
		}else {//记录花色出现的情况
			hand.Suits[Suits[handStr[i]]] |= faceValue
		}
	}
	return &hand
}


/*****************************判断牌型*******************/

//判断是否是为同花顺，如果是则记录最好牌序
func (maxHand *MaxHand) isStraightFlush(hand *Hand) bool{
	for i :=0;i < len(hand.Faces);i++{
		//先判断同花色牌是否达到要求，再判断这些牌可以组成同花顺吗
		if countOne(hand.Suits[i]) >= 5-hand.GhostNum{
			//标记是同花
			maxHand.FlushFlag = true
			//标记同花表的位置
			maxHand.FlushSuit = i

			if tempValue := findStraight(hand.Suits[i],hand.GhostNum);tempValue > 0{
				//与之前的牌面做一个对比
				if maxHand.MaxHand == 0 {
					maxHand.MaxHand = tempValue
				} else {
					maxHand.MaxHand = If(tempValue > maxHand.MaxHand && tempValue != A2345, tempValue, maxHand.MaxHand).(uint64)
				}
				maxHand.MaxCase = StraightFlush
				//这边不能直接return 因为还有其它花色需要判断，又可能出现更好的牌面
			}
		}
	}
	return maxHand.MaxCase == StraightFlush
}

// 筛选四条 赖子最多三个，超过三个必为同花顺
//判断是否是为四条，如果是则记录最好牌序
func (maxHand *MaxHand) isFourOfAKind(hand *Hand) bool{
	//四条只要判断Face[]表里是否存在
	if hand.Faces[3-hand.GhostNum] > 0{
		maxHand.MaxCase = FourOfAKind
		////判断赖子情况 考虑有赖子（1、2、3张）和没用赖子的情况
		//if hand.GhostNum == 0{
		//	//没有赖子就直接输出最大牌面 hand.Faces[3]
		//	firstOne := getFirstOne(hand.Faces[3])
		//	maxHand.MaxHand = leftMoveAndAdd(firstOne,4) | getFirstOne(firstOne^hand.Faces[0])
		//}else {
		//	//获得最大牌
		//	firstOne := getFirstOne(hand.Faces[3-hand.GhostNum])
		//	//如果有赖子多余，则补充A,否则选择最大一张牌
		//	maxHand.MaxHand = leftMoveAndAdd(firstOne,4) |
		//		If(firstOne&hand.Faces[3-hand.GhostNum+1] > 0,A,getFirstOne(firstOne^hand.Faces[0])).(uint64)
        //    //返回的结果记得转换成（uint64）
		//}
		return true
	}
	return false
}

//葫芦最多有几张赖子（筛选葫芦 赖子最多一个，超过一个必大于等于四条）
//判断是否是为葫芦，如果是则记录最好牌序
func (maxHand *MaxHand) isFullHouse(hand *Hand) bool{
	//这边选择两张牌必须>=2否则条件不成立
	if hand.Faces[2-hand.GhostNum] > 0 && countOne(hand.Faces[1]) >= 2  {
		maxHand.MaxCase = FullHouse
		//if hand.GhostNum == 0{
		//	//没有赖子就直接输出最大牌面 hand.Faces[2]
		//	firstOne := getFirstOne(hand.Faces[2])
		//	maxHand.MaxHand = leftMoveAndAdd(firstOne,3) | leftMoveAndAdd(getFirstOne(firstOne^hand.Faces[1]),2)
		//}else if hand.GhostNum == 1{
		//	//获得最大牌
		//	firstOne := getFirstOne(hand.Faces[1])
		//	maxHand.MaxHand = leftMoveAndAdd(firstOne,3) | leftMoveAndAdd(getFirstOne(firstOne^hand.Faces[1]),2)
		//}
		return true
	}
	return false
}

//判断是否是为同花，如果是则记录最好牌序  赖子[0,2]
func (maxHand *MaxHand) isFlush(hand *Hand) bool{
	if maxHand.FlushFlag{
		maxHand.MaxCase = Flush
		////最高五个位置必定出3个位置可以放赖子 最多2个赖子
		////先进行与操作留下最高五个位置，在进行亦或操作得到可以放的位置
		//tempValue := (hand.Suits[maxHand.FlushSuit] & AKQJT)^AKQJT
		////计算出可以填充的位置
		//tempValue = deleteLastOne(tempValue,int(countOne(tempValue)-hand.GhostNum))
		////开始填充
		//tempValue = hand.Suits[maxHand.FlushSuit] | tempValue
		////删除没用的位数，留下5位
		//maxHand.MaxHand = deleteLastOne(tempValue,int(countOne(tempValue)-5))
		return true
	}
	return false
}

//判断是否是为顺子，如果是则记录最好牌序 赖子[0,2]
func (maxHand *MaxHand) isStraight(hand *Hand) bool{
    if maxHand.MaxHand = findStraight(hand.Faces[0],hand.GhostNum);maxHand.MaxHand!=0{
		maxHand.MaxCase = Straight
		return true
	}
	return false
}


//判断是否是为三条，如果是则记录最好牌序   赖子[0,2]
func (maxHand *MaxHand) isThreeOfAKind(hand *Hand) bool{

	if hand.Faces[2-hand.GhostNum] > 0{
		maxHand.MaxCase = ThreeOfAKind
		////赖子只能在三对上，不用考虑其它因素
		////获取三条
		//firstOne := getFirstOne(hand.Faces[2-hand.GhostNum])
		////去掉三条选取单根
		//tempValue := hand.Faces[0]^firstOne
		////留下最高的两个单根
		//maxHand.MaxHand = leftMoveAndAdd(firstOne,3) | deleteLastOne(tempValue,int(countOne(tempValue)-2))
		return true
	}
	return false
}

//两对不存在赖子，否则必然三条 赖子[0]
//判断是否是为两对，如果是则记录最好牌序
func (maxHand *MaxHand) isTwoPair(hand *Hand) bool{
	if countOne(hand.Faces[1]) >= 2{
		maxHand.MaxCase = TwoPair
		//tempValue := deleteLastOne(hand.Faces[1],int(countOne(hand.Faces[1])-2))
		//maxHand.MaxHand = leftMoveAndAdd(tempValue,2) |  getFirstOne(hand.Faces[0]^tempValue)
		////获取最大的对子
		//firstOne :=getFirstOne(hand.Faces[1])
		////获取第二大的对子
		//secondOne :=getFirstOne(hand.Faces[1]^firstOne)
		//maxHand.MaxHand = leftMoveAndAdd(firstOne,2) | leftMoveAndAdd(secondOne,2) | getFirstOne(hand.Faces[0]^(firstOne | secondOne))
		return true
	}
	return false
}

//对子最多一个赖子 赖子[0,1]
//判断是否是为对子，如果是则记录最好牌序
func (maxHand *MaxHand) isOnePair(hand *Hand) bool{
	if hand.Faces[1-hand.GhostNum] > 0{
		//如果有赖子，唯一的一个赖子一定在对子上
		maxHand.MaxCase = OnePair
		//firstOne :=getFirstOne(hand.Faces[1-hand.GhostNum])
		//tempValue := hand.Faces[0]^firstOne
		////对子左边移动，并且或上剩下牌面最高的三张
		//maxHand.MaxHand = leftMoveAndAdd(firstOne,2) | deleteLastOne(tempValue,int(countOne(tempValue)-3))
		return true
	}
	return false
}

//高牌不存在赖子 赖子[0]
//判断是否是为高牌，如果是则记录最好牌序
func (maxHand *MaxHand) isHighCard(hand *Hand) bool{
	//不是其它牌，则一定是高牌，直接去最高五张牌为排序
	maxHand.MaxCase = HighCard
	//maxHand.MaxHand = deleteLastOne(hand.Faces[0], int(countOne(hand.Faces[0])-5))
	return true
}



//获得牌型，因为在一些情况下不用获取牌型
func (maxHand *MaxHand)getmaxhand(hand *Hand)  {
	switch maxHand.MaxCase {
	case HighCard:
		maxHand.MaxHand = deleteLastOne(hand.Faces[0], int(countOne(hand.Faces[0])-5))
	case OnePair:
		firstOne :=getFirstOne(hand.Faces[1-hand.GhostNum])
		tempValue := hand.Faces[0]^firstOne
		//对子左边移动，并且或上剩下牌面最高的三张
		maxHand.MaxHand = leftMoveAndAdd(firstOne,2) | deleteLastOne(tempValue,int(countOne(tempValue)-3))
	case TwoPair:
		tempValue := deleteLastOne(hand.Faces[1],int(countOne(hand.Faces[1])-2))
		maxHand.MaxHand = leftMoveAndAdd(tempValue,2) |  getFirstOne(hand.Faces[0]^tempValue)
	case ThreeOfAKind:
		//赖子只能在三对上，不用考虑其它因素
		//获取三条
		firstOne := getFirstOne(hand.Faces[2-hand.GhostNum])
		//去掉三条选取单根
		tempValue := hand.Faces[0]^firstOne
		//留下最高的两个单根
		maxHand.MaxHand = leftMoveAndAdd(firstOne,3) | deleteLastOne(tempValue,int(countOne(tempValue)-2))
	case Flush:
		//最高五个位置必定出3个位置可以放赖子 最多2个赖子
		//先进行与操作留下最高五个位置，在进行亦或操作得到可以放的位置
		tempValue := (hand.Suits[maxHand.FlushSuit] & AKQJT)^AKQJT
		//计算出可以填充的位置
		tempValue = deleteLastOne(tempValue,int(countOne(tempValue)-hand.GhostNum))
		//开始填充
		tempValue = hand.Suits[maxHand.FlushSuit] | tempValue
		//删除没用的位数，留下5位
		maxHand.MaxHand = deleteLastOne(tempValue,int(countOne(tempValue)-5))
	case FullHouse:
		if hand.GhostNum == 0{
			//没有赖子就直接输出最大牌面 hand.Faces[2]
			firstOne := getFirstOne(hand.Faces[2])
			maxHand.MaxHand = leftMoveAndAdd(firstOne,3) | leftMoveAndAdd(getFirstOne(firstOne^hand.Faces[1]),2)
		}else if hand.GhostNum == 1{
			//获得最大牌
			firstOne := getFirstOne(hand.Faces[1])
			maxHand.MaxHand = leftMoveAndAdd(firstOne,3) | leftMoveAndAdd(getFirstOne(firstOne^hand.Faces[1]),2)
		}
	case FourOfAKind:
		//判断赖子情况 考虑有赖子（1、2、3张）和没用赖子的情况
		if hand.GhostNum == 0{
			//没有赖子就直接输出最大牌面 hand.Faces[3]
			firstOne := getFirstOne(hand.Faces[3])
			maxHand.MaxHand = leftMoveAndAdd(firstOne,4) | getFirstOne(firstOne^hand.Faces[0])
		}else {
			//获得最大牌
			firstOne := getFirstOne(hand.Faces[3-hand.GhostNum])
			//如果有赖子多余，则补充A,否则选择最大一张牌
			maxHand.MaxHand = leftMoveAndAdd(firstOne,4) |
				If(firstOne&hand.Faces[3-hand.GhostNum+1] > 0,A,getFirstOne(firstOne^hand.Faces[0])).(uint64)
			//返回的结果记得转换成（uint64）
		}
	}
}


// 比较两张手牌、支持任意数量手牌及任意数量赖子
func Compare(strA string, strB string) int {
	//解析牌型与得分
	playerA := analysisHandStr(strA).getMaxHand()
	playerB := analysisHandStr(strB).getMaxHand()

	// 比较最大牌型
	if winner := getWinnerCompare(playerA.MaxCase, playerB.MaxCase); winner != 0 {
		return winner
	}

	//比较牌型
	// 顺子&同花顺存在“A2345”这一特殊情况，此时为最小顺子，需要手动标记（权值score设为0）
	scoreA := If(playerA.MaxHand == A2345, uint64(0), playerA.MaxHand).(uint64)
	scoreB := If(playerB.MaxHand == A2345, uint64(0), playerB.MaxHand).(uint64)
	return getWinnerCompare(scoreA, scoreB)
}