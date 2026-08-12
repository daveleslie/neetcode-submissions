class Solution:
    def dailyTemperatures(self, temperatures: list[int]) -> list[int]:
        result = [0] * len(temperatures)
        stack = []
        for i in range(0, len(temperatures), 1):
            while len(stack) > 0 and temperatures[i] > temperatures[stack[-1]]:
                stack_index = stack.pop()
                result[stack_index] = i - stack_index
            stack.append(i)
        return result
