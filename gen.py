m = "".join(map(chr, [4, 0, 0, 8, 0, 0, 0, 0]))
b = m * 1024


with open("of.hex", "w+") as f:
    f.write(b)
