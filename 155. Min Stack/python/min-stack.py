import pytest

class Node:
    def __init__(self, val: int, next_min_idx: int):
        self.val = val
        self.next_min_idx = next_min_idx

class MinStack:

    def __init__(self):
        self.st = []
        self.min_idx = -1


    def push(self, val: int) -> None:
        self.st.append(Node(val, self.min_idx))
        if self.min_idx == -1:
            self.min_idx = 0
            return
        if self.st[self.min_idx].val > val:
            self.min_idx = len(self.st) - 1


    def pop(self) -> None:
        node = self.st.pop()
        self.min_idx = node.next_min_idx


    def top(self) -> int:
        return self.st[-1].val


    def getMin(self) -> int:
        return self.st[self.min_idx].val



# Your MinStack object will be instantiated and called as such:
# obj = MinStack()
# obj.push(val)
# obj.pop()
# param_3 = obj.top()
# param_4 = obj.getMin()

@pytest.fixture
def minst():
    return MinStack()

def test_min_stack(minst):
    minst.push(-2);
    minst.push(0);
    minst.push(-3);
    assert minst.getMin() == -3;
    minst.pop();
    minst.top();
    assert minst.getMin() == -2;
