import pytest
import math


class Solution:
    def kthGrammar(self, n: int, k: int) -> int:
        def recursion(k, flip: bool = False):
            if k == 1:
                return 1 if flip else 0
            res = recursion(math.ceil(k/2), k % 2 == 0)
            if flip:
                res ^= 1 # 0 xor 1 = 1 | 1 xor 1 = 0
            return res
        return recursion(k)


@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    ((2,1), 0),
])
def test_solution(sol, args, expected):
    assert sol.kthGrammar(*args) == expected
