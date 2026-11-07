import pytest


class Solution:
    def calculate(self, s: str) -> int:
        stack = []
        n = 0
        op = "+"
        s += "+"

        for c in s:
            if c == " ":
                continue

            if c.isdigit():
                n = n * 10 + int(c)
                continue

            if op == "+":
                stack.append(n)
            elif op == "-":
                stack.append(-n)
            elif op == "*":
                stack.append(stack.pop() * n)
            elif op == "/":
                stack.append(int(stack.pop() / n))

            op = c
            n = 0

        return sum(stack)

@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    (("3+2*2",), 7),
    (("0-2147483647",), -2147483647),
])
def test_solution(sol, args, expected):
    assert sol.calculate(*args) == expected
