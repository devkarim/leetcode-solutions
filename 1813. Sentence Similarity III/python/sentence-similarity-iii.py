import pytest


class Solution:
    def areSentencesSimilar(self, sentence1: str, sentence2: str) -> bool:
        words1 = sentence1.split()
        words2 = sentence2.split()
        # make words1 always larger array
        if len(words1) < len(words2):
            words1, words2 = words2, words1
        n, m = len(words1), len(words2)

        i, common_prefix = 0, 0
        while i < m and words1[i] == words2[i]:
            common_prefix += 1
            i += 1

        i, common_suffix = -1, 0
        while i >= -m and words1[i] == words2[i]:
            common_suffix += 1
            i -= 1

        if common_suffix + common_prefix >= m:
            return True

        return False


@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    (("My name is Haley", "My Haley",), True),
])
def test_solution(sol, args, expected):
    assert sol.areSentencesSimilar(*args) == expected
