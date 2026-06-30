class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False
        sNum = []
        tNum = []
        for i in range(0, len(s)):
            sNum.append(ord(s[i]))
            tNum.append(ord(t[i]))
        sNum.sort()
        tNum.sort()
        for i in range(0, len(sNum)):
            if sNum[i] != tNum[i]:
                return False
        return True
        
            