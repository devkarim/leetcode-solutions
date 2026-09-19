import pytest


class Solution:
    def minCapability(self, nums: List[int], k: int) -> int:
        nums_min = min(nums)
        if k == 1:
            return nums_min

        left = nums_min
        right = max(nums)
        while left <= right:
            mid = (left + right) // 2
            k_needed = 0
            last_idx = -2
            for curr_idx, n in enumerate(nums):
                if n <= mid and curr_idx - last_idx > 1:
                    k_needed += 1
                    last_idx = curr_idx

            if k_needed < k:
                left = mid + 1
            else:
                right = mid - 1

        return left


@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    (([2,3,5,9], 2), 5),
])
def test_solution(sol, args, expected):
    assert sol.minCapability(*args) == expected
