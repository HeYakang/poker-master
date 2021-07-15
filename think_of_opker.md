# Poker

## 题目

扑克牌52张，花色黑桃spades，红心hearts，方块diamonds，草花clubs各13张，2-10，J，Q，K，A

Face：即2-10，J，Q，K，A，其中10用T来表示。

Color：即S(spades)、H(hearts)、D(diamonds)、C(clubs)

用 Face字母+小写Color字母表示一张牌，比如As表示黑桃A，其中A为牌面，s为spades，即黑桃，Ah即红心A，以此类推。  现分别给定任意两手牌各有5张，例如：AsAhQsQhQc，vs KsKhKdKc2c，请按德州扑克的大小规则来判断双方大小。

代码要求有通用性，可以任意输入一手牌或几手牌来进行比较。

- 进阶1：

  给定两手各7张牌，分别找出最大的牌型（选取其中5张）

- 进阶2:

  如果有一张牌是赖子（即可以当任何牌），用Xn表示赖子。



## 评判标准

1、结果正确

2、代码可读性

3、代码行数

4、运行速度，比如可以接受1万手的输入来评估速度

5、可扩展性，是否支持7张？是否可以支持赖子?



## 附1：手牌

五张 - [match.json](http://jira:8100/download/attachments/7307298/match.json?version=1&modificationDate=1594199265000&api=v2)

五张-结果 [match_result.json](http://jira:8100/download/attachments/7307298/match_result.json?version=2&modificationDate=1597654158000&api=v2)

七张-结果 [seven_cards_with_ghost.json](http://jira:8100/download/attachments/7307298/seven_cards_with_ghost.json?version=2&modificationDate=1596422580000&api=v2)

七张带赖子-结果 [seven_cards_with_ghost.result.json](http://jira:8100/download/attachments/7307298/seven_cards_with_ghost.result.json?version=1&modificationDate=1596422019000&api=v2)

## 附2：德州扑克的大小规则

1. 皇家同花顺
   同花色的A, K, Q, J和10. 
   平手牌：在摊牌的时候有两副多副皇家同花顺时，平分筹码。
2. 同花顺
   五张同花色的连续牌。
   平手牌：如果摊牌时有两副或多副同花顺，连续牌的头张牌大的获得筹码。如果是两副或多副相同的连续牌，平分筹码。
3. 四条
   其中四张是相同点数但不同花的扑克牌，第五张是随意的一张牌 。
   平手牌：如果两组或者更多组摊牌，则四张牌中的最大者赢局，如果一组人持有的四张牌是一样的，那么第五张牌最大者赢局（起脚牌）。如果起脚牌也一样，平分彩池。
4. 满堂彩（葫芦，三带二）
   由三张相同点数及任何两张其他相同点数的扑克牌组成 。
   平手牌：如果两组或者更多组摊牌，那么三张相同点数中较大者赢局。如果三张牌都一样，则两张牌中点数较大者赢局，如果所有的牌都一样，则平分彩池。
5. 同花
   此牌由五张不按顺序但相同花的扑克牌组成 。
   平手牌：如果不止一人抓到此牌相，则牌点最高的人赢得该局，如果最大点相同，则由第二、第三、第四或者第五张牌来决定胜负，如果所有的牌都相同，平分彩池。
6. 顺子 
   此牌由五张顺序扑克牌组成。
   平手牌：如果不止一人抓到此牌，则五张牌中点数最大的赢得此局，如果所有牌点数都相同，平分彩池。
7. 三条
   由三张相同点数和两张不同点数的扑克组成 。
   平手牌：如果不止一人抓到此牌，则三张牌中最大点数者赢局，如果三张牌都相同，比较第四张牌，必要时比较第五张，点数大的人赢局。如果所有牌都相同，则平分彩池。
8. 两对
   两对点数相同但两两不同的扑克和随意的一张牌组成。
   平手牌：如果不止一人抓大此牌相，牌点比较大的人赢，如果比较大的牌点相同，那么较小牌点中的较大者赢，如果两对牌点相同，那么第五张牌点较大者赢（起脚牌）。如果起脚牌也相同，则平分彩池。
9. 一对
   由两张相同点数的扑克牌和另三张随意的牌组成。
   平手牌：如果不止一人抓到此牌，则两张牌中点数大的赢，如果对牌都一样，则比较另外三张牌中大的赢，如果另外三张牌中较大的也一样则比较第二大的和第三大的，如果所有的牌都一样，则平分彩池。
10. 单张大牌 
    既不是同一花色也不是同一点数的五张牌组成。 
    平手牌：如果不止一人抓到此牌，则比较点数最大者，如果点数最大的相同，则比较第二、第三、第四和第五大的，如果所有牌都相同，则平分彩池。



## 实现思路

1. Bit Map 位图来存储处理数据
2. 剪枝思想



## 项目安排计划

## 流程表	

### 提取数据

目的：从给定的json文件中提取牌信息，并且按照后续处理的需要进行存储。



### 循环比较结果

目的：将处理后的数据进行处理，得到最大的手牌类型与牌序



### 输出结果

目的：输出得到的结果与所耗费的时间



## 比较流程

**最大手牌类型比较过程**

同花？ → 同花顺？

↓

四条？→ 三带2（葫芦）？→ 顺子？→ 三条? → 两对？→ 一对？ → 高牌？

**同牌形型比较大小**

按照牌型从大到小牌排序进行比较

例如：两对(77669、77445) 则比较得到前者大

​            三条(55594、55576)则比较得到前者大