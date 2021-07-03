# LeetCode-GO

个人的LeetCode算法练习代码

## 心得体会

### 利用分治策略降低时间复杂度

对于某一个较大规模的问题，通过将其划分为多个小规模的问题（但是与之前的解法不变）。其每个小规模的问题的时间复杂度累积起来之后，原复杂度中的某个$n$项会变为$lg n$项。比如某算法复杂度为$O(n^2)$，分治后变为$O(n lg n)$。

一个很鲜明的例子是：[合并K个已排序链表](/leetcode/0023.Merge-K-Sorted-Lists/0023.Merge-K-Sorted-Lists.go)。

在[最初的版本](https://github.com/SalHe/LeetCode-GO/blob/c7d037aa6dc9b31c5fee2f39c4efe40a695606fe/leetcode/0023.Merge-K-Sorted-Lists/0023.Merge-K-Sorted-Lists.go)中，是普通的版本。而在[改进版本](https://github.com/SalHe/LeetCode-GO/blob/0e9fee9d4c0de7d5d077d7fbc040e180ae325b7b/leetcode/0023.Merge-K-Sorted-Lists/0023.Merge-K-Sorted-Lists.go)中，仅仅只是把算法改进成了运用分治策略的算法，速度就显著提升了N倍（力扣测试时间由312ms改进到了12ms）。
