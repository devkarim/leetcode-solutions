import pytest


class Solution:
    def minimizedMaximum(self, n: int, quantities: list[int]) -> int:
        left = 1
        right = max(quantities)

        while left <= right:
            mid = (left + right) // 2
            n_needed = 0
            for q in quantities:
                n_needed += -(-q // mid)
            if n_needed > n:
                left = mid + 1
            else:
                right = mid - 1

        return left


@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    ((6, [11,6]), 3),
])
def test_solution(sol, args, expected):
    assert sol.minimizedMaximum(*args) == expected
