import matplotlib.pyplot as plt
import statistics as st

fig, ax = plt.subplots()

def create_statistics(filename):
	f = open(filename, "r+")
	times = [float(x) for x in f.readlines()]
	times.sort()
	ax.scatter(range(len(times)), times, label = filename)
	return dict(filename = filename, mean = st.mean(times), stdev = st.stdev(times), time = times)

a = create_statistics("ep.a.txt")
b = create_statistics("ep.b.txt")
s = create_statistics("ep.s.txt")
serial_a = create_statistics("ep_serial.a.txt")
serial_s = create_statistics("ep_serial.s.txt")
dalvan_serial_a = create_statistics("ep_serial.A.cpp.txt")
dalvan_serial_s = create_statistics("ep_serial.S.cpp.txt")

#mean
#points = [a["mean"], b["mean"], s["mean"], serial_a["mean"], serial_s["mean"], dalvan_serial_a["mean"], dalvan_serial_s["mean"]]
#plt.scatter(range(len(points)), points)

ax.legend()
plt.show()
