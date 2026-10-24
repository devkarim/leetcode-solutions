import pytest


class Solution:
    def trap(self, height: list[int]) -> int:
        n = len(height)
        max_left = [0] * n
        max_right = [0] * n

        max_left[0] = 0
        for i in range(1, n):
            max_left[i] = max(max_left[i - 1], height[i - 1])

        max_right[-1] = 0
        for i in range(n - 2, -1, -1):
            max_right[i] = max(max_right[i + 1], height[i + 1])

        res = 0

        for i, h in enumerate(height):
            water_blocks = min(max_right[i], max_left[i]) - h
            if water_blocks > 0:
                res += water_blocks

        return res


@pytest.fixture
def sol():
    return Solution()


@pytest.mark.parametrize(
    "args, expected",
    [
        (([0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1],), 6),
    ],
)
def test_solution(sol, args, expected):
    assert sol.trap(*args) == expected
