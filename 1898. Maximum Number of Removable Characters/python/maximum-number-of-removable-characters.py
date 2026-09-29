import pytest


class Solution:
    def maximumRemovals(self, s: str, p: str, removable: list[int]) -> int:
        def is_subsequence(k: int, exclude: set[int]) -> bool:
            p_idx = 0
            for c_idx, c in enumerate(s):
                if c_idx in exclude:
                    continue
                p_c = p[p_idx]
                if c == p_c:
                    p_idx += 1
                if p_idx >= len(p):
                    return True
            return False

        l = 0
        r = len(removable)

        while l <= r:
            k = (l + r) // 2
            exclude = set(removable[:k])

            if is_subsequence(k, exclude):
                l = k + 1
            else:
                r = k - 1

        return r



@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    (("abcacb", "ab", [3,1,0]), 2),
])
def test_solution(sol, args, expected):
    assert sol.maximumRemovals(*args) == expected
