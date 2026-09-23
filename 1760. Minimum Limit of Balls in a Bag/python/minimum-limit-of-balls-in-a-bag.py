import pytest


class Solution:
    def minimumSize(self, nums: list[int], maxOperations: int) -> int:
        left = 1
        right = max(nums)

        while left <= right:
            mid = (left + right) // 2
            ops_needed = 0
            for n in nums:
                ops_needed += (-(-n // mid)) - 1
            if ops_needed <= maxOperations:
                right = mid - 1
            else:
                left = mid + 1
        return left


@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    (([9], 2), 3),
])
def test_solution(sol, args, expected):
    assert sol.minimumSize(*args) == expected
