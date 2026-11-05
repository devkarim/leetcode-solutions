import pytest


class Solution:
    def maxWidthRamp(self, nums: list[int]) -> int:
        max_right = [0] * len(nums)
        max_right[-1] = nums[-1]
        for r in range(len(nums) - 2, -1, -1):
            max_right[r] = max(max_right[r+1], nums[r])

        l, r = 0, 1
        res = 0
        while l < len(nums) and r < len(nums):
            if nums[l] <= max_right[r]:
                res = max(res, r - l)
                r += 1
            else:
                l += 1
                if l >= r:
                    r = l + 1
        return res


@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    (([9,8,1,0,1,9,4,0,4,1],), 7),
])
def test_solution(sol, args, expected):
    assert sol.maxWidthRamp(*args) == expected
