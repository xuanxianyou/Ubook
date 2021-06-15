package database

import (
	"GoProject/WebProject/MicroService/UserBook/book/model"
	"testing"
)

func TestMongoDB_CreateCollection(t *testing.T) {
	mongo:=ConnectMongo()
	defer mongo.DisconnectMongo()
	book:=model.Book{
		Name:         "算法导论（原书第3版）",
		Cover:        "./cover/算法导论.jpg",
		Author:       "Thomas H.Cormen / Charles E.Leiserson / Ronald L.Rivest / Clifford Stein",
		Press:        "机械工业出版社",
		Publication:  "华章IT",
		Page:         780,
		Price:        128.00,
		Series:       "计算机科学丛书",
		ISBN:         "9787111407010",
		Score:        9.2,
		Tags:         []string{"算法","算法导论","计算机","计算机科学","编程","经典","程序设计","数学"},
		ContentBrief: `在有关算法的书中，有一些叙述非常严谨，但不够全面；另一些涉及了大量的题材，但又缺乏严谨性。本书将严谨性和全面性融为一体，深入讨论各类算法，并着力使这些算法的设计和分析能为各个层次的读者接受。全书各章自成体系，可以作为独立的学习单元；算法以英语和伪代码的形式描述，具备初步程序设计经验的人就能看懂；说明和解释力求浅显易懂，不失深度和数学严谨性。

全书选材经典、内容丰富、结构合理、逻辑清晰，对本科生的数据结构课程和研究生的算法课程都是非常实用的教材，在IT专业人员的职业生涯中，本书也是一本案头必备的参考书或工程实践手册。

第3版的主要变化：

新增了van Emde Boas树和多线程算法，并且将矩阵基础移至附录。

修订了递归式（现在称为“分治策略”）那一章的内容，更广泛地覆盖分治法。

移除两章很少讲授的内容：二项堆和排序网络。

修订了动态规划和贪心算法相关内容。

流网络相关材料现在基于边上的全部流。

由于关于矩阵基础和Strassen算法的材料移到了其他章，矩阵运算这一章的内容所占篇幅更小。

修改了对Knuth-Morris-Pratt字符串匹配算法的讨论。

新增100道练习和28道思考题，还更新并补充了参考文献。

`,
		AuthorBrief:  `Thomas H. Cormen （托马斯•科尔曼） 达特茅斯学院计算机科学系教授、系主任。目前的研究兴趣包括：算法工程、并行计算、具有高延迟的加速计算。他分别于1993年、1986年获得麻省理工学院电子工程和计算机科学博士、硕士学位，师从Charles E. Leiserson教授。由于他在计算机教育领域的突出贡献，Cormen教授荣获2009年ACM杰出教员奖。

Charles E. Leiserson（查尔斯•雷瑟尔森）麻省理工学院计算机科学与电气工程系教授，Margaret MacVicar Faculty Fellow。他目前主持MIT超级计算技术研究组，并是MIT计算机科学和人工智能实验室计算理论研究组的成员。他的研究兴趣集中在并行和分布式计算的理论原理，尤其是与工程现实相关的技术研究。Leiserson教授拥有卡内基•梅隆大学计算机科学博士学位，还是ACM、IEEE和SIAM的会士。

Ronald L. Rivest （罗纳德•李维斯特）现任麻省理工学院电子工程和计算机科学系安德鲁与厄纳•维特尔比（Andrew and Erna Viterbi）教授。他是MIT计算机科学和人工智能实验室的成员，并领导着其中的信息安全和隐私中心。他1977年从斯坦福大学获得计算机博士学位，主要从事密码安全、计算机安全算法的研究。他和Adi Shamir和Len Adleman一起发明了RSA公钥算法，这个算法在信息安全中获得最大的突破，这一成果也使他和Shamir、Adleman一起得到2002年ACM图灵奖。他现在担任国家密码学会的负责人。

Clifford Stein（克利福德•斯坦）哥伦比亚大学计算机科学系和工业工程与运筹学系教授，他还是工业工程与运筹学系的系主任。在加入哥伦比亚大学大学之前，他在达特茅斯学院计算机科学系任教9年。Stein教授拥有MIT硕士和博士学位。他的研究兴趣包括：算法的设计与分析，组合优化、运筹学、网络算法、调度、算法工程和生物计算。`,
		Catalogue:    `出版者的话
译者序
前言
第一部分 基础知识
第1章 算法在计算中的作用 3
1.1 算法 3
1.2 作为一种技术的算法 6
思考题 8
本章注记 8
第2章 算法基础 9
2.1 插入排序 9
2.2 分析算法 13
2.3 设计算法 16
2.3.1 分治法 16
2.3.2 分析分治算法 20
思考题 22
本章注记 24
第3章 函数的增长 25
3.1 渐近记号 25
3.2 标准记号与常用函数 30
思考题 35
本章注记 36
第4章 分治策略 37
4.1 最大子数组问题 38
4.2 矩阵乘法的Strassen算法 43
4.3 用代入法求解递归式 47
4.4 用递归树方法求解递归式 50
4.5 用主方法求解递归式 53
4.6 证明主定理 55
4.6.1 对b的幂证明主定理 56
4.6.2 向下取整和向上取整 58
思考题 60
本章注记 62
第5章 概率分析和随机算法 65
5.1 雇用问题 65
5.2 指示器随机变量 67
5.3 随机算法 69
5.4 概率分析和指示器随机变量的进一步使用 73
5.4.1 生日悖论 73
5.4.2 球与箱子 75
5.4.3 特征序列 76
5.4.4 在线雇用问题 78
思考题 79
本章注记 80
第二部分 排序和顺序统计量
第6章 堆排序 84
6.1 堆 84
6.2 维护堆的性质 85
6.3 建堆 87
6.4 堆排序算法 89
6.5 优先队列 90
思考题 93
本章注记 94
第7章 快速排序 95
7.1 快速排序的描述 95
7.2 快速排序的性能 97
7.3 快速排序的随机化版本 100
7.4 快速排序分析 101
7.4.1 最坏情况分析 101
7.4.2 期望运行时间 101
思考题 103
本章注记 106
第8章 线性时间排序 107
8.1 排序算法的下界 107
8.2 计数排序 108
8.3 基数排序 110
8.4 桶排序 112
思考题 114
本章注记 118
第9章 中位数和顺序统计量 119
9.1 最小值和最大值 119
9.2 期望为线性时间的选择算法 120
9.3 最坏情况为线性时间的选择算法 123
思考题 125
本章注记 126
第三部分 数据结构
第10章 基本数据结构 129
10.1 栈和队列 129
10.2 链表 131
10.3 指针和对象的实现 134
10.4 有根树的表示 137
思考题 139
本章注记 141
第11章 散列表 142
11.1 直接寻址表 142
11.2 散列表 143
11.3 散列函数 147
11.3.1 除法散列法 147
11.3.2 乘法散列法 148
11.3.3 全域散列法 148
11.4 开放寻址法 151
11.5 完全散列 156
思考题 158
本章注记 160
第12章 二叉搜索树 161
12.1 什么是二叉搜索树 161
12.2 查询二叉搜索树 163
12.3 插入和删除 165
12.4 随机构建二叉搜索树 169
思考题 171
本章注记 173
第13章 红黑树 174
13.1 红黑树的性质 174
13.2 旋转 176
13.3 插入 178
13.4 删除 183
思考题 187
本章注记 191
第14章 数据结构的扩张 193
14.1 动态顺序统计 193
14.2 如何扩张数据结构 196
14.3 区间树 198
思考题 202
本章注记 202
第四部分 高级设计和分析技术
第15章 动态规划 204
15.1 钢条切割 204
15.2 矩阵链乘法 210
15.3 动态规划原理 215
15.4 最长公共子序列 222
15.5 最优二叉搜索树 226
思考题 231
本章注记 236
第16章 贪心算法 237
16.1 活动选择问题 237
16.2 贪心算法原理 242
16.3 赫夫曼编码 245
16.4 拟阵和贪心算法 250
16.5 用拟阵求解任务调度问题 253
思考题 255
本章注记 257
第17章 摊还分析 258
17.1 聚合分析 258
17.2 核算法 261
17.3 势能法 262
17.4 动态表 264
17.4.1 表扩张 265
17.4.2 表扩张和收缩 267
思考题 270
本章注记 273
第五部分 高级数据结构
第18章 B树 277
18.1 B树的定义 279
18.2 B树上的基本操作 281
18.3 从B树中删除关键字 286
思考题 288
本章注记 289
第19章 斐波那契堆 290
19.1 斐波那契堆结构 291
19.2 可合并堆操作 292
19.3 关键字减值和删除一个结点 298
19.4 最大度数的界 300
思考题 302
本章注记 305
第20章 van Emde Boas树 306
20.1 基本方法 306
20.2 递归结构 308
20.2.1 原型van Emde Boas结构 310
20.2.2 原型van Emde Boas结构上的操作 311
20.3 van Emde Boas树及其操作 314
20.3.1 van Emde Boas树 315
20.3.2 van Emde Boas树的操作 317
思考题 322
本章注记 323
第21章 用于不相交集合的数据结构 324
21.1 不相交集合的操作 324
21.2 不相交集合的链表表示 326
21.3 不相交集合森林 328
21.4 带路径压缩的按秩合并的分析 331
思考题 336
本章注记 337
第六部分 图算法
第22章 基本的图算法 341
22.1 图的表示 341
22.2 广度优先搜索 343
22.3 深度优先搜索 349
22.4 拓扑排序 355
22.5 强连通分量 357
思考题 360
本章注记 361
第23章 最小生成树 362
23.1 最小生成树的形成 362
23.2 Kruskal算法和Prim算法 366
思考题 370
本章注记 373
第24章 单源最短路径 374
24.1 Bellman-Ford算法 379
24.2 有向无环图中的单源最短路径问题 381
24.3 Dijkstra算法 383
24.4 差分约束和最短路径 387
24.5 最短路径性质的证明 391
思考题 395
本章注记 398
第25章 所有结点对的最短路径问题 399
25.1 最短路径和矩阵乘法 400
25.2 Floyd-Warshall算法 404
25.3 用于稀疏图的Johnson算法 409
思考题 412
本章注记 412
第26章 最大流 414
26.1 流网络 414
26.2 Ford\Fulkerson方法 418
26.3 最大二分匹配 428
26.4 推送重贴标签算法 431
26.5 前置重贴标签算法 438
思考题 446
本章注记 449
第七部分 算法问题选编
第27章 多线程算法 453
27.1 动态多线程基础 454
27.2 多线程矩阵乘法 465
27.3 多线程归并排序 468
思考题 472
本章注记 476
第28章 矩阵运算 478
28.1 求解线性方程组 478
28.2 矩阵求逆 486
28.3 对称正定矩阵和最小二乘逼近 489
思考题 493
本章注记 494
第29章 线性规划 495
29.1 标准型和松弛型 499
29.2 将问题表达为线性规划 504
29.3 单纯形算法 507
29.4 对偶性 516
29.5 初始基本可行解 520
思考题 525
本章注记 526
第30章 多项式与快速傅里叶变换 527
30.1 多项式的表示 528
30.2 DFT与FFT 531
30.3 高效FFT实现 536
思考题 539
本章注记 541
第31章 数论算法 543
31.1 基础数论概念 543
31.2 最大公约数 547
31.3 模运算 550
31.4 求解模线性方程 554
31.5 中国余数定理 556
31.6 元素的幂 558
31.7 RSA公钥加密系统 561
31.8 素数的测试 565
31.9 整数的因子分解 571
思考题 574
本章注记 576
第32章 字符串匹配 577
32.1 朴素字符串匹配算法 578
32.2 Rabin\Karp算法 580
32.3 利用有限自动机进行字符串匹配 583
32.4 Knuth-Morris-Pratt算法 588
思考题 594
本章注记 594
第33章 计算几何学 595
33.1 线段的性质 595
33.2 确定任意一对线段是否相交 599
33.3 寻找凸包 604
33.4 寻找最近点对 610
思考题 613
本章注记 615
第34章 NP完全性 616
34.1 多项式时间 619
34.2 多项式时间的验证 623
34.3 NP完全性与可归约性 626
34.4 NP完全性的证明 633
34.5 NP完全问题 638
34.5.1 团问题 638
34.5.2 顶点覆盖问题 640
34.5.3 哈密顿回路问题 641
34.5.4 旅行商问题 644
34.5.5 子集和问题 645
思考题 647
本章注记 649
第35章 近似算法 651
35.1 顶点覆盖问题 652
35.2 旅行商问题 654
35.2.1 满足三角不等式的旅行商问题 654
35.2.2 一般旅行商问题 656
35.3 集合覆盖问题 658
35.4 随机化和线性规划 661
35.5 子集和问题 663
思考题 667
本章注记 669
第八部分 附录：数学基础知识
附录A 求和 672
A.1 求和公式及其性质 672
A.2 确定求和时间的界 674
思考题 678
附录注记 678
附录B 集合等离散数学内容 679
B.1 集合 679
B.2 关系 682
B.3 函数 683
B.4 图 685
B.5 树 687
B.5.1 自由树 688
B.5.2 有根树和有序树 689
B.5.3 二叉树和位置树 690
思考题 691
附录注记 692
附录C 计数与概率 693
C.1 计数 693
C.2 概率 696
C.3 离散随机变量 700
C.4 几何分布与二项分布 702
C.5 二项分布的尾部 705
思考题 708
附录注记 708
附录D 矩阵 709
D.1 矩阵与矩阵运算 709
D.2 矩阵基本性质 712
思考题 714
附录注记 715
参考文献 716
索引 732`,
	}
	mongo.CreateCollection(book)
}
