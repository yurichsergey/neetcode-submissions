class Solution {
    /**
     * @param {number[]} nums
     * @param {number} target
     * @return {number}
     */
    search(nums, target) {
		let l = 0
		let r = nums.length
		while (l <= r) {
			const n = ((l+r)/2)|0
			if (nums[n] === target) {
				return n
			}
			if (nums[n] > target) {
				r = n - 1
			} else {
				l = n + 1
			}
		}
		return -1
	}
}
