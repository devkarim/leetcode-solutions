import pytest


class Solution:
    def minimumLength(self, s: str) -> int:
        res = n = len(s)
        l, r = 0, n - 1

        while l < r and s[l] == s[r]:
            res -= 2
            while (l + 1) < n and (l + 1) != r and s[l + 1] == s[r]:
                res -= 1
                l += 1
            while (r - 1) >= 0 and l != (r - 1) and s[l] == s[r - 1]:
                res -= 1
                r -= 1
            l += 1
            r -= 1

        return res


@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    (("cabaabac",), 0),
])
def test_solution(sol, args, expected):
    assert sol.minimumLength(*args) == expected
