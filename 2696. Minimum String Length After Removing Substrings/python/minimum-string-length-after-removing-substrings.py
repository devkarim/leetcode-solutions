import pytest


class Solution:
    def minLength(self, s: str) -> int:
        stack = []
        for c in s:
            if stack and stack[-1] == "C" and c == "D":
                stack.pop()
            elif stack and stack[-1] == "A" and c == "B":
                stack.pop()
            else:
                stack.append(c)
        return len(stack)


@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    (("ABFCACDB",), 2),
])
def test_solution(sol, args, expected):
    assert sol.minLength(*args) == expected
