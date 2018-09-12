import sys

import matplotlib.pyplot as plt
import numpy as np


with open(sys.argv[1]) as f:
    rtts = map(float, f.readlines())

plt.hist(np.array(list(rtts)), bins=10)
plt.show()
