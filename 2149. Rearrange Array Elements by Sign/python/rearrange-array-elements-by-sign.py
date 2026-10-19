import pytest


class Solution:
    def rearrangeArray(self, nums: list[int]) -> list[int]:
        res = []
        pos_nums = []
        neg_nums = []

        for n in nums:
            if n > 0:
                pos_nums.append(n)
            else:
                neg_nums.append(n)

        for i in range(0, len(pos_nums)):
            res.append(pos_nums[i])
            res.append(neg_nums[i])

        return res


@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    (([3,1,-2,-5,2,-4],), [3,-2,1,-5,2,-4]),
])
def test_solution(sol, args, expected):
    assert sol.rearrangeArray(*args) == expected
