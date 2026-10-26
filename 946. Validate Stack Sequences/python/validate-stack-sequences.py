import pytest


class Solution:
    def validateStackSequences(self, pushed: list[int], popped: list[int]) -> bool:
        st = []
        i = 0
        for p in pushed:
            st.append(p)
            while st and st[-1] == popped[i]:
                st.pop()
                i += 1
        return len(st) == 0


@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    (([1,2,3,4,5],[4,5,3,2,1]), True),
])
def test_solution(sol, args, expected):
    assert sol.validateStackSequences(*args) == expected
