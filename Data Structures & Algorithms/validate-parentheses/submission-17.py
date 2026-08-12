class Solution:
	def isValid(self, s: str) -> bool:
		stack = []
		closeToOpen = {
			"}":"{",
			"]":"[",
			")":"("
		}

		for c in s:
			# if c is a closing bracket
			if c in closeToOpen:
				if len(stack) == 0 or stack[len(stack) - 1] != closeToOpen[c]:
					return False
				else:
					stack.pop()

			else:
				# c is an opening bracket, append to stack
				stack.append(c)
		
		return len(stack) == 0