import pytest


class Solution:
    def threeSum(self, nums: list[int]) -> list[list[int]]:
        nums.sort()
        res = list()

        for idx, num1 in enumerate(nums):
            pass

        return res


@pytest.fixture
def sol():
    return Solution()

def test_solution(sol: Solution):
    nums = [-1,0,1,2,-1,-4] # [-4, -1, -1, 0, 1, 2]
    assert sol.threeSum(nums) == [[-1, -1, 2], [-1, 0, 1]]

