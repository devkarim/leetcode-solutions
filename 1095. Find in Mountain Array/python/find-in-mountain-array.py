import pytest


class MountainArray:
    def __init__(self, arr: list[int]):
        self._arr = arr

    def get(self, index: int) -> int:
        return self._arr[index]

    def length(self) -> int:
        return len(self._arr)


class Solution:
    def findInMountainArray(self, target: int, mountainArr: 'MountainArray') -> int:
        cache = {}
        n = mountainArr.length()

        def binary_search(search_left: bool):
            res = float("inf")
            l, r = 0, n - 1
            while l <= r:
                m = (l + r) // 2
                curr_val = cache.get(m) if m in cache else mountainArr.get(m)
                if m not in cache:
                    cache[m] = curr_val
                if curr_val == target:
                    res = min(res, m)
                if (m - 1) < 0:
                    left_val = -1
                else:
                    left_val = cache.get(m - 1) if (m - 1) in cache else mountainArr.get(m - 1)
                    if (m - 1) not in cache:
                        cache[m - 1] = left_val
                if search_left:
                    # we are on the left side of the mountain
                    if curr_val > left_val:
                        if target < curr_val:
                            r = m - 1
                        else:
                            l = m + 1
                    else: # we are on the right side of the mountain
                        r = m - 1
                else:
                    # we are on the right side of the mountain
                    if curr_val < left_val:
                        if target > curr_val:
                            r = m - 1
                        else:
                            l = m + 1
                    else: # we are on the left side of the mountain
                        l = m + 1
            return res
        left_ans = binary_search(True)
        if left_ans != float("inf"):
            return left_ans
        right_ans = binary_search(False)
        return right_ans if right_ans != float("inf") else -1


@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("target, arr, expected", [
    (3, [1, 2, 3, 4, 5, 3, 1], 2),
])
def test_solution(sol, target, arr, expected):
    assert sol.findInMountainArray(target, MountainArray(arr)) == expected

