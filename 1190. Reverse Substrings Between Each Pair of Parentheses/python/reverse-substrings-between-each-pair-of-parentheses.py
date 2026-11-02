import pytest


class Solution:
    def reverseParentheses(self, s: str) -> str:
        stack = []
        for c in s:
            if c != ")":
                stack.append(c)
            else:
                reversed_str = []
                while stack and stack[-1] != "(":
                    reversed_str.append(stack.pop())
                # remove last "("
                stack.pop()
                # add reversed_str to stack
                stack.extend(reversed_str)
        return "".join(stack)

@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    (("(abcd)",), "dcba"),
    (("(u(love)i)",), "iloveu"),
    (("(ed(et(oc))el)",), "leetcode"),
])
def test_solution(sol, args, expected):
    assert sol.reverseParentheses(*args) == expected
