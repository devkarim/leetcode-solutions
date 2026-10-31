import pytest


class Solution:
    def removeKdigits(self, num: str, k: int) -> str:
        stack = []
        for n in num:
            while k != 0 and stack and int(stack[-1]) > int(n):
                stack.pop()
                k -= 1
            if n != "0" or stack:
                stack.append(n)
        while stack and k > 0:
            stack.pop()
            k -= 1
        if not stack:
            return "0"
        return "".join(stack)


@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    (("1432219", 3), "1219"),
    (("1432239", 3), "1223"),
    (("10200", 1), "200"),
    (("10", 2), "0"),
    (("9", 1), "0"),
])
def test_solution(sol, args, expected):
    assert sol.removeKdigits(*args) == expected
