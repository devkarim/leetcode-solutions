import pytest


class Solution:
    def searchRange(self, nums: list[int], target: int) -> list[int]:
        def get_bound(find_start: bool, initial_l: int = 0):
            l = initial_l
            r = len(nums) - 1

            while l <= r:
                m = (l + r) // 2
                if (find_start and nums[m] < target) or (
                    not find_start and nums[m] <= target
                ):
                    l = m + 1
                else:
                    r = m - 1

            return l if find_start else r

        start = get_bound(True)
        end = get_bound(False, start)

        if start >= len(nums) or nums[start] != target:
            return [-1, -1]

        return [start, end]


@pytest.fixture
def sol():
    return Solution()


@pytest.mark.parametrize(
    "args, expected",
    [
        (([5, 7, 7, 8, 8, 10], 8), [3, 4]),
    ],
)
def test_solution(sol, args, expected):
    assert sol.searchRange(*args) == expected
