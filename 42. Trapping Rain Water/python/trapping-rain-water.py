import pytest


class Solution:
    def trap(self, height: list[int]) -> int:
        n = len(height)
        i = res = 0

        while i < n:
            j = i + 1
            st = []
            while j < n and height[i] > height[j]:
                st.append(height[j])
                j += 1
            if j < n and height[i] <= height[j]:
                while st:
                    res += (height[i] - st.pop())
            while st:
                tmp = st.pop()
                if st and tmp > st[-1]:
                    res += (tmp - st.pop())
                    st.append(tmp)
            i = j
        return res


@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    (([0,1,0,2,1,0,1,3,2,1,2,1],), 6),
])
def test_solution(sol, args, expected):
    assert sol.trap(*args) == expected
