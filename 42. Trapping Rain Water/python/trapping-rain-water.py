import pytest


class Solution:
    def trap(self, height: list[int]) -> int:
        l, r = 0, len(height) - 1
        left_max, right_max = height[l], height[r]

        res = 0
        while l < r:
            if left_max < right_max:
                l += 1
                left_max = max(height[l], left_max)
                res += left_max - height[l]
            else:
                r -= 1
                right_max = max(height[r], right_max)
                res += right_max - height[r]
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
