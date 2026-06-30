class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        for lIdx, num1 in enumerate(nums):
            for rIdx, num2 in enumerate(nums[lIdx+1:]):
                if num1 + num2 == target:
                    return [lIdx, rIdx + lIdx + 1]