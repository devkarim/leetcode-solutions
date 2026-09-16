import pytest


class Solution:
    def shipWithinDays(self, weights: list[int], days: int) -> int:
        w_len = len(weights)

        maxWeights = max(weights)
        if w_len == days:
            return maxWeights

        sumOfWeights = sum(weights)
        if days == 1:
            return sumOfWeights

        left = maxWeights
        right = sumOfWeights

        while left <= right:
            min_w = (left + right) // 2
            days_needed = 1
            currDay = 0
            for w in weights:
                if currDay + w <= min_w:
                    currDay += w
                else:
                    days_needed += 1
                    currDay = w

            if days_needed <= days:
                right = min_w - 1
            else:
                left = min_w + 1

        return left


@pytest.fixture
def sol():
    return Solution()


@pytest.mark.parametrize(
    "args, expected",
    [
        (([1, 2, 3, 1, 1], 4), 3),
    ],
)
def test_solution(sol, args, expected):
    assert sol.shipWithinDays(*args) == expected
