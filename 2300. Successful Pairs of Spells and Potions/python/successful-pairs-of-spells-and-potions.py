import pytest


class Solution:
    def successfulPairs(self, spells: List[int], potions: List[int], success: int) -> List[int]:
        res = [0] * len(spells)

        potions = sorted(potions)

        i = -1
        for spell in spells:
            start = 0
            end = len(potions) - 1

            while start <= end:
                mid = (start + end) // 2
                if potions[mid]*spell >= success:
                    end = mid-1
                else:
                    start = mid+1

            i += 1
            res[i] = len(potions) - start
        return res


@pytest.fixture
def sol():
    return Solution()

def test_solution(sol):
    spells = [5, 1, 3]
    potions = [1, 2, 3, 4, 5]
    success = 7

    expected = [4, 0, 3]
    assert sol.successfulPairs(spells, potions, success) == expected
