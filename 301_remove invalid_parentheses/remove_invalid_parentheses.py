class Solution(object):
    def removeRightInvalidParentheses(self, s):
        cs = [[]]
        sb = []
        b = 0
        
        for c in s:
            if c == '(':
                b += 1
            if c == ')':
                b -= 1
            sb.append(c)
            
            if b < 0:
                cc = set()
                for ca in cs:
                    c_ca = ca + sb
                    rbi = [i for i, c in enumerate(c_ca) if c == ")"]
                    for i in rbi:
                        rc = c_ca[:]
                        del rc[i]
                        cc.add("".join(rc))
                    cs = [list(s) for s in cc]
                b = 0
                sb = []
    
        t = "".join(sb)
        return ["".join(s) + t for s in cs]    

    def swapBracket(self, c):
        if c == '(':
            return ')'
        if c == ')':
            return '('
        return c

    def removeInvalidParentheses(self, s):
        fs = []
        cs = self.removeRightInvalidParentheses(s)
        for i, s in enumerate(cs):
            fs += ["".join([self.swapBracket(c2) for c2 in s[::-1]]) 
                    for s in self.removeRightInvalidParentheses("".join([self.swapBracket(c1) for c1 in s[::-1]]))]
        return fs