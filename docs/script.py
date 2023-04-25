from create_statistics import create_statistics
import numpy as np
import matplotlib.pyplot as plt
import statistics as st

times = create_statistics("is/is_A_NOBUCKETS.txt")["time"]

plt.xticks(np.arange(min(times), max(times)+1, 0.02))
plt.hist(times, bins=50, density=False)
plt.show()