import pytest


class Solution:
    def removeDuplicates(self, s: str, k: int) -> str:
        stack = [] # [char, cnt]
        for c in s:
            if stack and stack[-1][0] == c:
                stack[-1][1] += 1
            else:
                stack.append([c, 1])

            if stack[-1][1] == k:
                stack.pop()
        return "".join(pair[0] * pair[1] for pair in stack)

@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    (("deeedbbcccbdaa", 3), "aa"),
])
def test_solution(sol, args, expected):
    assert sol.removeDuplicates(*args) == expected
