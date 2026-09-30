import pytest


class Solution:
    def maximumBeauty(self, items: list[list[int]], queries: list[int]) -> list[int]:
        items.sort()
        queries = sorted([(q, i) for i, q in enumerate(queries)])

        # [[1,2],[2,4],[3,2],[3,5],[5,6]] -> j
        # [1,2,3,4,5,6] -> i

        j = 0
        max_beauty = 0
        ans = [0] * len(queries)
        for q, i in queries:
            while j < len(items) and items[j][0] <= q:
                max_beauty = max(max_beauty, items[j][1])
                j += 1
            ans[i] = max_beauty

        return ans

@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    (([[1,2],[3,2],[2,4],[5,6],[3,5]], [1,2,3,4,5,6]), [2,4,5,5,6,6]),
])
def test_solution(sol, args, expected):
    assert sol.maximumBeauty(*args) == expected
