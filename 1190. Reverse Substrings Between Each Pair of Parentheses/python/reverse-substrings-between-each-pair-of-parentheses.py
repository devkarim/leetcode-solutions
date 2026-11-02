import pytest


class Solution:
    def reverseParentheses(self, s: str) -> str:
        pair = {}
        stack = []
        for i, c in enumerate(s):
            if c == "(":
                stack.append(i)
            elif c == ")":
                j = stack.pop()
                pair[i] = j
                pair[j] = i

        i, direction = 0, 1
        res = []

        while i < len(s):
            if s[i] == "(" or s[i] == ")":
                direction *= -1
                i = pair[i]
            else:
                res.append(s[i])
            i += direction

        return "".join(res)

@pytest.fixture
def sol():
    return Solution()


@pytest.mark.parametrize(
    "args, expected",
    [
        (("(abcd)",), "dcba"),
        (("(u(love)i)",), "iloveu"),
        (("(ed(et(oc))el)",), "leetcode"),
    ],
)
def test_solution(sol, args, expected):
    assert sol.reverseParentheses(*args) == expected
