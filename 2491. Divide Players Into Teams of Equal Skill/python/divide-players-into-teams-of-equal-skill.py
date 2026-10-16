import pytest


class Solution:
    def dividePlayers(self, skill: list[int]) -> int:
        skill.sort()

        l = 0
        r = len(skill) - 1
        skill_needed = skill[l] + skill[r]
        chemistry_sum = 0

        while l < r:
            if skill[l] + skill[r] != skill_needed:
                return -1
            chemistry_sum += skill[l] * skill[r]
            l += 1
            r -= 1

        return chemistry_sum


@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    (([3,2,5,1,3,4],), 22),
])
def test_solution(sol, args, expected):
    assert sol.dividePlayers(*args) == expected
