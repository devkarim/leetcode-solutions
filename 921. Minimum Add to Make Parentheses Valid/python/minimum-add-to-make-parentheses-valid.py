import pytest


class Solution:
    def minAddToMakeValid(self, s: str) -> int:
        open_brackets = closed_brackets = 0
        for c in s:
            if c == "(":
                open_brackets += 1
            elif c == ")":
                closed_brackets += 1
                if open_brackets:
                    closed_brackets -= 1
                    open_brackets -= 1
        return open_brackets + closed_brackets


@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    (("(((",), 3),
])
def test_solution(sol, args, expected):
    assert sol.minAddToMakeValid(*args) == expected
