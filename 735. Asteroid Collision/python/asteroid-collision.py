import pytest


class Solution:
    def asteroidCollision(self, asteroids: list[int]) -> list[int]:
        stack = []
        for ast in asteroids:
            while stack and ast < 0 and stack[-1] > 0:
                if abs(ast) > stack[-1]:
                    stack.pop()
                elif abs(ast) < stack[-1]:
                    ast = 0
                else:
                    ast = 0
                    stack.pop()
            if ast:
                stack.append(ast)
        return stack


@pytest.fixture
def sol():
    return Solution()


@pytest.mark.parametrize(
    "args, expected",
    [
        (([5, 7, 4, -6, 10, -11],), [-11]),
    ],
)
def test_solution(sol, args, expected):
    assert sol.asteroidCollision(*args) == expected
