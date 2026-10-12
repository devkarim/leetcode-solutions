import pytest


class Solution:
    def removeDuplicates(self, nums: list[int]) -> int:
        i, res = 0, 0
        while i < len(nums):
            group_length = 1
            while (i + group_length) < len(nums) and nums[i] == nums[i + group_length]:
                group_length += 1

            nums[res] = nums[i]
            res += 1
            if group_length >= 2:
                nums[res] = nums[i]
                res += 1
            i += group_length

        return res


@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    (([1,1,1,2,2,3],), 5),
])
def test_solution(sol, args, expected):
    assert sol.removeDuplicates(*args) == expected
