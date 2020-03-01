class Solution(object):
    def eval(self, expr):
        expr = expr.replace("True", "T")
        expr = expr.replace("False", "F")
        expr = expr.replace("AND", "&")
        expr = expr.replace("XOR", "x")
        expr = expr.replace("OR", "|")
        expr = expr.replace("NOT", "!")
        return self.mini_eval(expr)

    def mini_eval(self, expr):
        if expr == "T":
            return True
        if expr == "F":
            return False

        if expr[0] == "(":
            return self.eval(expr[1:-1])

        op = "&"
        op = "|" if expr[-1] == "|" else op
        op = "!" if expr[-1] == "!" else op
        op = "x" if expr[-1] == "x" else op
        res = op == "&"

        start = 0
        level = 0
        xor_count = 0
        for i in range(0, len(expr) - len(op)):
            if level == 0 and expr[i] in [",", ")"]:
                val = self.mini_eval(expr[start:i])
                start = i + 1
                if op == "&":
                    res &= val
                if op == "|":
                    res |= val
                if op == "!":
                    res = not val
                if op == "x":
                    if xor_count == 0:
                        res = val
                    else:
                        res = res != val
            if expr[i] == "(":
                level = level + 1
            elif expr[i] == ")":
                level = level - 1
        return res


print(Solution().eval("(True,False,AND)"))