class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        results = defaultdict(list)
        for s in strs:
            sortedS = ''.join(sorted(s))
            results[sortedS].append(s)
        return list(results.values())
        