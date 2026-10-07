import pytest


class Solution:
    def findMedianSortedArrays(self, nums1: list[int], nums2: list[int]) -> float:
        A, B = nums1, nums2
        if len(A) > len(B):
            A, B = B, A

        total = len(A) + len(B)

        half = total // 2
        l, r = 0, len(A) - 1

        while True:
            i = (l + r) // 2
            j = half - i - 2

            A_left = A[i] if i >= 0 else float('-inf')
            A_right = A[i + 1] if (i + 1) < len(A) else float('inf')

            B_left = B[j] if j >= 0 else float('-inf')
            B_right = B[j + 1] if (j + 1) < len(B) else float("inf")

            if A_right >= B_left and B_right >= A_left:
                if total % 2 != 0:
                    return min(A_right, B_right)
                return (max(A_left, B_left) + min(A_right, B_right)) / 2

            if A_right < B_left:
                l = i + 1
            elif A_left > B_right:
                r = i - 1


@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    (([1,3],[2]), 2.0),
])
def test_solution(sol, args, expected):
    assert sol.findMedianSortedArrays(*args) == expected
