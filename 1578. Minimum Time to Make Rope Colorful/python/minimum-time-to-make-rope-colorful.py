import pytest


class Solution:
    def minCost(self, colors: str, neededTime: list[int]) -> int:
        n = len(colors)
        left, cost = 0, 0
        while left < n:
            right = left + 1
            while right < n and colors[left] == colors[right]:
                cost += min(neededTime[left], neededTime[right])
                if neededTime[left] < neededTime[right]:
                    break
                right += 1
            left = right
        return cost


@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    (("abaac", [1,2,3,4,5]), 3),
])
def test_solution(sol, args, expected):
    assert sol.minCost(*args) == expected
