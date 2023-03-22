import matplotlib.pyplot as plt
import statistics as st

f_a = open("ep.a.txt", "r+")
f_b = open("ep.b.txt", "r+")
f_s = open("ep.s.txt", "r+")
f_sa = open("ep_serial.a.txt", "r+")
f_ss = open("ep_serial.s.txt", "r+")

times_a = [float(x) for x in f_a.readlines()]
times_b = [float(x) for x in f_b.readlines()]
times_s = [float(x) for x in f_s.readlines()]
times_serial_a = [float(x) for x in f_sa.readlines()]
times_serial_s = [float(x) for x in f_ss.readlines()]

statistics_a = dict(mean = st.mean(times_a), stdev = st.stdev(times_a), time = times_a)
statistics_b = dict(mean = st.mean(times_b), stdev = st.stdev(times_b), time = times_b)
statistics_s = dict(mean = st.mean(times_s), stdev = st.stdev(times_s), time = times_s)
statistics_serial_a = dict(mean = st.mean(times_serial_a), stdev = st.stdev(times_serial_a), time = times_serial_a)
statistics_serial_s = dict(mean = st.mean(times_serial_s), stdev = st.stdev(times_serial_s), time = times_serial_s)

#plt.scatter(range(len(times)), times)
plt.scatter(range(len(statistics_a["time"])), statistics_a["time"])
plt.scatter(range(len(statistics_b["time"])), statistics_b["time"])
plt.scatter(range(len(statistics_s["time"])), statistics_s["time"])
plt.scatter(range(len(statistics_serial_a["time"])), statistics_serial_a["time"])
plt.scatter(range(len(statistics_serial_s["time"])), statistics_serial_s["time"])
plt.show()