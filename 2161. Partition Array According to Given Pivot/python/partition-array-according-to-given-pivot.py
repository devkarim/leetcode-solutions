import pytest


class Solution:
    def pivotArray(self, nums: list[int], pivot: int) -> list[int]:
        ans = [0] * len(nums)
        less_idx = 0
        greater_idx = len(nums) - 1

        i = 0
        j = len(nums) - 1

        while i < len(nums):
            if nums[i] < pivot:
                ans[less_idx] = nums[i]
                less_idx += 1
            if nums[j] > pivot:
                ans[greater_idx] = nums[j]
                greater_idx -= 1
            i += 1
            j -= 1

        while greater_idx >= less_idx:
            ans[less_idx] = pivot
            less_idx += 1

        return ans
@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    (([9,12,5,10,14,3,10],10), [9,5,3,10,10,12,14]),
])
def test_solution(sol, args, expected):
    assert sol.pivotArray(*args) == expected
