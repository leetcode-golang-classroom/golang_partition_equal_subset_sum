# golang_partition_equal_subset_sum

Given a **non-empty** array `nums` containing **only positive integers**, find if the array can be partitioned into two subsets such that the sum of elements in both subsets is equal.

## Examples

**Example 1:**

```
Input: nums = [1,5,11,5]
Output: true
Explanation: The array can be partitioned as [1, 5, 5] and [11].

```

**Example 2:**

```
Input: nums = [1,2,3,5]
Output: false
Explanation: The array cannot be partitioned into equal sum subsets.

```

**Constraints:**

- `1 <= nums.length <= 200`
- `1 <= nums[i] <= 100`

## 解析

給定一個正整數陣列 nums, 

要求寫一個演算法判斷是否有辦法把 nums 分成兩個 陣列 set1, set2 使得 sum(set1) = sum(set2)

要分成兩個陣列使得兩個陣列合相等

因為所有元素都是 > 0

代表兩個陣列合 = sum(nums)/2

因為可以分成兩個陣列代表 sum(nums) 一定要是偶數

其實只要找到其中一個陣列其中和 = sum(nums)/2 則另一個一定也是相同

判斷是否有辦法把 nums 分成兩個 陣列 set1, set2 使得 sum(set1) = sum(set2) 

可以簡化為 

判斷 sum(nums) 是否為偶數且存在一個子集合其和 = sum(nums)/2

舉例來觀察問題：

假設有一個陣列 nums: [1,5,11,5]

透過窮舉法 可以先畫出以下決策樹

對每個元素都有選擇或是不選兩種選則

從 index: 0 開始

![decision_dfs.png](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/41f21602-b1ec-49f9-bcb0-d687363b6539/decision_dfs.png)

會發現透過 總共會有 $2^n$ 個結點

所以 最差的情況需要探訪所有結點 O($2^n)$而

換個角度針對和找出所有可能性

假設定義 dp[i, target_sum] 代表在前i個 可能夠組成和為 target_sum 可能性

dp[i, target_sum] = dp[i-i, target_sum] || dp[i-1,target_sum-num[i]]

前i 個能夠組成 target_sum 代表以下兩種其中一種可能 

1. 前i-1 可以組成 target_sum 
2. 前i-1 可以可以組成 target_sum - nums[i]

這樣需要對每個 end = i 去找 dp[i, 0] 到 dp[i, target_sum]

所以是時間複雜度  O(n*target_sum)

因為對 dp[i, target_sum] 最後都可以被下一個覆蓋

所以只需要 dp[0]~ dp[target_sum] 空間複雜度 O(target_sum)

![](https://i.imgur.com/Oe5dSyV.png)

而想要在優化不想要每個 sum 都去檢查

可以用一個 hash set 來紀錄每次考慮進去 element 所可能形成的 sum

初始話 dp[0]=struct{}{}

我們可以思考從最後一個開始考慮

每次從 dp 取出每個 sum + nums[i] 檢查如果等於 target_sum 則回傳 

否則把新的 sum+nums[i] 加入 hashset db 中

當把所有陣列元素都考慮進去之後

最後檢查是否 dp[target_sum] 是否存在 

如果是代表可以分成兩個相等的陣列，否則代表不行

![](https://i.imgur.com/8ymSwz2.png)

對每個 元素需要考慮 len(hashSet) 個檢查

所以大概是 O(n*len(hashSet))

而 hashSet 最多 target_sum 個

所以時間度複雜度 O(n*target_sum)

空間複雜度 O(target_sum) 

所以大致上跟上一個解法差不多 只是可以避免掉檢查那些無法合成的 sum

## 程式碼
```go
package sol

func canPartition(nums []int) bool {
	target_sum := 0
	for _, num := range nums {
		target_sum += num
	}
	if target_sum%2 != 0 {
		return false
	}
	target_sum /= 2
	dp := make(map[int]struct{})
	dp[0] = struct{}{}
	nLen := len(nums)
	for start := nLen - 1; start >= 0; start-- {
		nextDp := make(map[int]struct{})
		for sum := range dp {
			if sum+nums[start] == target_sum {
				return true
			}
			nextDp[target_sum+nums[start]] = struct{}{}
			nextDp[sum] = struct{}{}
		}
		dp = nextDp
	}
	_, exists := dp[target_sum]
	return exists
}
```
## 困難點

1. 要看出這個題目的遞迴子關係

## Solve Point

- [x]  建立了dp HashSet 來檢查已經可以組成的 sum
- [x]  每次檢查當下 nums[i] 與目前組成的 sum 能不能組成目標
- [x]  最後檢查 dp[target_sum] 有沒有在組成可能目標內