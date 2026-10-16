import pytest

from collections import Counter


class Solution:
    def dividePlayers(self, skill: list[int]) -> int:
        target = (2 * sum(skill)) / len(skill)
        # if required skill has decimals, ints cannot add up and contain decimals
        if target % 1 != 0:
            return -1
        target = int(target)

        skill_map = Counter(skill)

        chemistry_sum = 0
        for n in skill:
            if skill_map[n] == 0:
                continue
            diff = target - n
            if not skill_map[diff]:
                return -1

            skill_map[n] -= 1
            skill_map[diff] -= 1

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
