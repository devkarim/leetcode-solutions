import pytest


class Solution:
    def decodeString(self, s: str) -> str:
        stack = []

        for i in range(len(s)):
            if s[i] != "]":
                stack.append(s[i])
            else:
                # get the substr
                substr = ""
                while stack and stack[-1] != "[":
                    substr = stack.pop() + substr
                stack.pop() # remove the last [

                # get the repeat digit
                repeat = ""
                while stack and stack[-1].isdigit():
                    repeat = stack.pop() + repeat

                # combine
                stack.append(int(repeat) * substr)

        return "".join(stack)
@pytest.fixture
def sol():
    return Solution()

@pytest.mark.parametrize("args, expected", [
    (("3[a]2[bc]",), "aaabcbc"),
    (("3[a2[c]]",), "accaccacc"),
    (("2[abc]3[cd]ef",), "abcabccdcdcdef"),
    (("3[z]2[2[y]pq4[2[jk]e1[f]]]ef",), "zzzyypqjkjkefjkjkefjkjkefjkjkefyypqjkjkefjkjkefjkjkefjkjkefef"),
])
def test_solution(sol, args, expected):
    assert sol.decodeString(*args) == expected
