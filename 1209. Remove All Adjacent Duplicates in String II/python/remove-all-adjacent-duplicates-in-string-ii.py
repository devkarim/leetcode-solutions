import pytest


class Solution:
    def removeDuplicates(self, s: str, k: int) -> str:
        stack = [] # (char, cnt)
        for c in s:
            cnt = ((stack and stack[-1][0] == c and stack[-1][1]) or 0) + 1
            if cnt == k:
                stack.pop()
                cnt = 0
            if cnt != 0:
                if cnt == 1:
                    stack.append((c, cnt))
                else:
                    stack[-1] = (c, cnt)
        return "".join(pair[0] * pair[1] for pair in stack)


@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    (("deeedbbcccbdaa", 3), "aa"),
])
def test_solution(sol, args, expected):
    assert sol.removeDuplicates(*args) == expected
