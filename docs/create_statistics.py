import statistics as st

def create_statistics(filename):
	f = open(filename, "r+")
	times = [float(x) for x in f.readlines()]
	times.sort()
	return dict(filename = filename, mean = st.mean(times), stdev = st.stdev(times), time = times)
