import pytest


class Solution:
    def clearDigits(self, s: str) -> str:
        def isdigit(c):
            return ord("0") <= ord(c) <= ord("9")

        stack = []
        for c in s:
            if isdigit(c):
                stack.pop()
            else:
                stack.append(c)
        return "".join(stack)


@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    (("abc34",), "a"),
])
def test_solution(sol, args, expected):
    assert sol.clearDigits(*args) == expected
