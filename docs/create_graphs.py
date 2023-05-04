from create_statistics import create_statistics
import matplotlib.pyplot as plt

def create_histogram_for_times(filename):
	times = create_statistics(filename)["time"]

	plt.hist(times, bins=10)

	plt.title('Histogram of Data')
	plt.xlabel('Value')
	plt.ylabel('Frequency')

	plt.show()

def create_boxplot_for_times(filenames):

	timeslist = []
	for filename in filenames.values():
		timeslist.append(create_statistics(filename)["time"])

	fig, ax = plt.subplots()
	ax.boxplot(timeslist)

	ax.set_title('Experimentos')
	ax.set_xticklabels(list(filenames.keys()))
	ax.set_xlabel('Versões')
	ax.set_ylabel('Tempo (s)')

	plt.show()

def create_barplot_for_means(filenames):
	meanslist = []
	for filename in filenames.values():
		meanslist.append(create_statistics(filename)["mean"])

	fig, ax = plt.subplots()

	ax.bar(list(filenames.keys()), meanslist, color = ["red", "blue", "yellow"])

	ax.set_title('Experimentos')
	ax.set_xlabel('Versao')
	ax.set_ylabel('Tempo medio (s)')

	plt.show()