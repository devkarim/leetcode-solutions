import pytest


class Solution:
    def maxDepth(self, s: str) -> int:
        nested_brackets = res = 0
        for c in s:
            if c == '(':
                nested_brackets += 1
                res = max(res, nested_brackets)
            if c == ')':
                nested_brackets -= 1
        return res


@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    (("(1+(2*3)+((8)/4))+1",), 3),
])
def test_solution(sol, args, expected):
    assert sol.maxDepth(*args) == expected
