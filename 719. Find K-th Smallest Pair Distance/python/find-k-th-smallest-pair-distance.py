import pytest


class Solution:
    def smallestDistancePair(self, nums: list[int], k: int) -> int:
        nums.sort()
        n = len(nums)

        def count_pairs(max_t):
            count = 0
            left = 0

            for right in range(1, n):
                while nums[right] - nums[left] > max_t:
                    left += 1

                count += right - left

            return count

        l, r = 0, nums[-1] - nums[0]

        while l <= r:
            t = (l + r) // 2
            if count_pairs(t) >= k:
                r = t - 1
            else:
                l = t + 1

        return l


@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    (([1,3,1], 1), 0),
])
def test_solution(sol, args, expected):
    assert sol.smallestDistancePair(*args) == expected
