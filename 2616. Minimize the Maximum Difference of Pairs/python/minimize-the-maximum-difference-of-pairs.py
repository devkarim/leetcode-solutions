import pytest


class Solution:
    def minimizeMax(self, nums: List[int], p: int) -> int:
        nums_sorted = sorted(nums)
        nums_len = len(nums)

        left = 0
        right = nums_sorted[-1] - nums_sorted[0]

        while left <= right:
            mid = (left + right) // 2
            pairs_needed = 0
            i = 0
            while i+1 < nums_len:
                if nums_sorted[i+1] - nums_sorted[i] <= mid:
                    pairs_needed += 1
                    i += 2
                else:
                    i += 1
            if pairs_needed >= p:
                right = mid - 1
            else:
                left = mid + 1

        return left


@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    (([1,5,8,9],2), 4),
])
def test_solution(sol, args, expected):
    assert sol.minimizeMax(*args) == expected
