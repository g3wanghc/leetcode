class Solution(object):
    def isPowerOfFour(self, num):
        """
        :type num: int
        :rtype: bool
        """
        num = float(num)
        while(num >= 4):
            num /= 4
            if not num.is_integer():
                return False
        return num == 1
