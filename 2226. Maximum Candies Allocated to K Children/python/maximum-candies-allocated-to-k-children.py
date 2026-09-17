import pytest


class Solution:
    def maximumCandies(self, candies: list[int], k: int) -> int:
        piles_sum = sum(candies)
        if piles_sum < k:
            return 0

        left = 1
        right = max(candies)

        while left <= right:
            max_candies = (left + right) // 2
            k_needed = 0
            for n in candies:
                k_needed += (n // max_candies)
            if k_needed >= k:
                left = max_candies + 1
            else:
                right = max_candies - 1

        return right


@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    (([5,8,6], 3), 5),
])
def test_solution(sol, args, expected):
    assert sol.maximumCandies(*args) == expected
