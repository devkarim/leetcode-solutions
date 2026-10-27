import pytest


class Solution:
    def asteroidCollision(self, asteroids: list[int]) -> list[int]:
        st = []
        res = []
        for ast in asteroids:
            res.append(ast)
            if ast > 0:
                st.append(ast)
            else:
                while st:
                    last_pos = st.pop()
                    if last_pos == abs(ast):
                        res.pop()
                        res.pop()
                        break
                    elif last_pos > abs(ast):
                        st.append(last_pos)
                        res.pop()
                        break
                    else:
                        res.pop()
                        res.pop()
                        res.append(ast)
        return res

@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    (([5,7,4,-6,10,-11],), [-11]),
])
def test_solution(sol, args, expected):
    assert sol.asteroidCollision(*args) == expected
