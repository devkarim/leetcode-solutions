import pytest


class Solution:
    def rotate(self, nums: list[int], k: int) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        k = k % len(nums)
        if k == 0:
            return
        nums[:k], nums[k:] = nums[-k:], nums[:-k]

@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("nums, k, expected", [
    ([1, 2, 3, 4, 5, 6, 7], 3, [5, 6, 7, 1, 2, 3, 4]),
])
def test_solution(sol, nums, k, expected):
    sol.rotate(nums, k)
    assert nums == expected
