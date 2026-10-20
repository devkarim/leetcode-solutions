import pytest


class Solution:
    def rearrangeArray(self, nums: list[int]) -> list[int]:
        pos_idx, neg_idx = 0, 1
        res = [0] * len(nums)
        for n in nums:
            if n > 0:
                res[pos_idx] = n
                pos_idx += 2
            else:
                res[neg_idx] = n
                neg_idx += 2
        return res

@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    (([3,1,-2,-5,2,-4],), [3,-2,1,-5,2,-4]),
])
def test_solution(sol, args, expected):
    assert sol.rearrangeArray(*args) == expected
