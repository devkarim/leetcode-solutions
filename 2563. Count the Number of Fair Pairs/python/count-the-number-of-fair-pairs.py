import pytest


class Solution:
    def countFairPairs(self, nums: List[int], lower: int, upper: int) -> int:
        def count_pairs_below(val):
            l, r, res = 0, len(nums) - 1, 0
            while l <= r:
                total = nums[l] + nums[r]
                if total < val:
                    res += r - l
                    l += 1
                else:
                    r -= 1

            return res

        nums.sort()
        return count_pairs_below(upper + 1) - count_pairs_below(lower)

@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    (([0,1,7,4,4,5], 3, 6), 6),
])
def test_solution(sol, args, expected):
    assert sol.countFairPairs(*args) == expected
