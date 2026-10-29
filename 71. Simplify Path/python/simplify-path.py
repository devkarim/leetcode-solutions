import pytest


class Solution:
    def simplifyPath(self, path: str) -> str:
        curr, stack = [], []
        for c in path + "/":
            if c == "/":
                name = "".join(curr)
                if name == "..":
                    if stack:
                        stack.pop()
                elif name != "." and name != "":
                    stack.append(name)
                curr = []
            else:
                curr.append(c)

        return "/" + "/".join(stack)


@pytest.fixture
def sol():
    return Solution()


@pytest.mark.parametrize(
    "args, expected",
    [
        (("/home/user/Documents/../Pictures",), "/home/user/Pictures"),
        (("/home//foo/",), "/home/foo"),
    ],
)
def test_solution(sol, args, expected):
    assert sol.simplifyPath(*args) == expected
