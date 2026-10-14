import pytest


class Solution:
    def fourSum(self, nums: list[int], target: int) -> list[list[int]]:
        nums = sorted(nums)
        res, tuples = [], []

        def kSum(k, start, t):
            # base case
            if k == 2:
                l, r = start, len(nums) - 1
                while l < r:
                    total = nums[l] + nums[r]
                    if total == t:
                        res.append(tuples + [nums[l], nums[r]])
                        l += 1
                        r -= 1
                        while l < r and nums[l] == nums[l - 1]:
                            l += 1
                    elif total < t:
                        l += 1
                    else:
                        r -= 1
                return
            for i in range(start, len(nums) - k + 1):
                if i > start and nums[i] == nums[i - 1]:
                    continue
                tuples.append(nums[i])
                kSum(k - 1, i + 1, t - nums[i])
                tuples.pop()

        kSum(4, 0, target)

        return res

@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    (([1, 0, -1, 0, -2, 2], 0), [[-2, -1, 1, 2], [-2, 0, 0, 2], [-1, 0, 0, 1]]),
])
def test_solution(sol, args, expected):
    assert sol.fourSum(*args) == expected
