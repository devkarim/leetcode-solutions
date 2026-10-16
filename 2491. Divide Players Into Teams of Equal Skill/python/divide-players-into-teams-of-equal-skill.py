import pytest

from collections import Counter


class Solution:
    def dividePlayers(self, skill: list[int]) -> int:
        if (2 * sum(skill)) % len(skill) != 0:
            return -1

        target = (2 * sum(skill)) // len(skill)
        count = Counter(skill)
        chemistry_sum = 0

        for n in skill:
            if count[n] == 0:
                continue

            count[n] -= 1
            diff = target - n
            if not count[diff]:
                return -1

            count[diff] -= 1
            chemistry_sum += (n * diff)

        return chemistry_sum

@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    (([3,2,5,1,3,4],), 22),
])
def test_solution(sol, args, expected):
    assert sol.dividePlayers(*args) == expected
