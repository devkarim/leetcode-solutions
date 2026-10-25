import pytest


class Solution:
    def removeStars(self, s: str) -> str:
        st = []
        for c in s:
            if c == "*":
                st.pop()
            else:
                st.append(c)
        return "".join(st)


@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    (("leet**cod*e",), "lecoe"),
])
def test_solution(sol, args, expected):
    assert sol.removeStars(*args) == expected
