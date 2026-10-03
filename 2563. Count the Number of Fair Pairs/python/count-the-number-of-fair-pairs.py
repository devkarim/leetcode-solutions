import pytest


class Solution:
    def countFairPairs(self, nums: List[int], lower: int, upper: int) -> int:
        def lower_bound(n, l, r):
            while l <= r:
                m = (l + r) // 2
                if nums[m] >= n:
                    r = m - 1
                else:
                    l = m + 1
            return l

        nums.sort()
        res = 0

        for idx, n in enumerate(nums):
            l = lower_bound(lower - n, idx+1, len(nums) - 1)
            r = lower_bound(upper - n + 1, idx+1, len(nums) - 1)

            res += r - l

        return res


@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    (([0,1,7,4,4,5], 3, 6), 6),
])
def test_solution(sol, args, expected):
    assert sol.countFairPairs(*args) == expected
