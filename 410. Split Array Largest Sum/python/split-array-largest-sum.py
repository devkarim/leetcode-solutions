import pytest


class Solution:
    def splitArray(self, nums: list[int], k: int) -> int:
        l, r = max(nums), sum(nums)
        while l <= r:
            m = (l + r) // 2

            subarrays = 1
            curr_sum = 0

            for n in nums:
                if curr_sum + n > m:
                    subarrays += 1
                    curr_sum = 0
                curr_sum += n

            if k >= subarrays:
                r = m - 1
            else:
                l = m + 1

        return l


@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    (([7,2,5,10,8],2), 18),
])
def test_solution(sol, args, expected):
    assert sol.splitArray(*args) == expected
