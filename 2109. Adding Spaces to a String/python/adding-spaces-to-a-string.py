import pytest


class Solution:
    def addSpaces(self, s: str, spaces: list[int]) -> str:
        space_idx, res_idx = 0, 0
        res = [None] * (len(s) + len(spaces))
        for c_idx, c in enumerate(s):
            if space_idx < len(spaces) and c_idx == spaces[space_idx]:
                res[res_idx] = " "
                space_idx += 1
                res_idx += 1
            res[res_idx] = c
            res_idx += 1

        return "".join(res)


@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    (("icodeinpython", [1,5,7,9]), "i code in py thon"),
])
def test_solution(sol, args, expected):
    assert sol.addSpaces(*args) == expected
