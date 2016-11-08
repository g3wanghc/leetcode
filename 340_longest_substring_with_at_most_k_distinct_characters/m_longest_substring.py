class Solution(object):
    def lengthOfLongestSubstringKDistinct(self, s, M):
        used_chars = dict()
        used_chars_count = 0
        head_index = 0
        longest_subarray_len = 0
        
        uniq_chars_s = len(set(s))
        if uniq_chars_s < M:
            return len(s)
        elif uniq_chars_s == M:
            return len(s)
        
        for i, c in enumerate(s):
            if c in used_chars:
                used_chars[c] += 1
            else:
                used_chars[c] = 1
                used_chars_count += 1
            if used_chars_count > M:
                expiring_char = s[head_index]
                if used_chars[expiring_char] == 1:
                    used_chars.pop(expiring_char)
                    used_chars_count -= 1
                else:
                    used_chars[expiring_char] -= 1
                head_index += 1
            else:
                longest_subarray_len += 1
                
        return longest_subarray_len
        
