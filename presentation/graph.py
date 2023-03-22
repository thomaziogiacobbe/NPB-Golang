import matplotlib.pyplot as plt
import statistics as st

f_a = open("ep.a.txt", "r+")
f_b = open("ep.b.txt", "r+")
f_s = open("ep.s.txt", "r+")

times_a = [float(x) for x in f_a.readlines()]
times_b = [float(x) for x in f_b.readlines()]
times_s = [float(x) for x in f_s.readlines()]

statistics_a = dict(mean = st.mean(times_a), stdev = st.stdev(times_a))
statistics_b = dict(mean = st.mean(times_b), stdev = st.stdev(times_b))
statistics_s = dict(mean = st.mean(times_s), stdev = st.stdev(times_s))

#plt.scatter(range(len(times)), times)
plt.scatter(range(3), [statistics_a["mean"], statistics_b["mean"], statistics_s["mean"]])
plt.show()