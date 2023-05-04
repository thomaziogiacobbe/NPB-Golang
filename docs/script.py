from create_statistics import create_statistics
from filenames import *
import matplotlib.pyplot as plt

def create_barchart_for_means(filenames, title, output):
	authors = list(filenames.keys())
	dataset = []
	for filename in filenames.values():
		dataset.append(create_statistics(filename))

	means = [d["mean"] for d in dataset]

	fig, ax = plt.subplots(figsize = (4.8,6.4))

	width = 0.25
	colors = ['blue', 'red', 'orange'] if len(authors) == 3 else ['blue', 'orange']

	author_count = 0
	for i in range(len(means)):
		rects = ax.bar(
			i/len(means),
			means[i],
			width=width,
			color=colors[author_count],
			label=authors[author_count])
		for rect in rects:
			height = rect.get_height()
			ax.annotate(f'{height:.4f}', xy=(rect.get_x() + rect.get_width() / 2, height), xytext=(0, 3), textcoords="offset points", ha='center', va='bottom')
		author_count = (author_count + 1) % 3

	ax.set_title(title)
	ax.set_xticks([i/len(authors) for i in range(len(authors))])
	ax.set_xticklabels(authors)
	ax.set_ylabel('Average time (s)')

	ax.legend()

	plt.savefig(output)

create_barchart_for_means(ep_s_graph, "Kernel EP Class S", "ep/ep_S.png")
create_barchart_for_means(ep_a_graph, "Kernel EP Class A", "ep/ep_A.png")
create_barchart_for_means(ep_b_graph, "Kernel EP Class B", "ep/ep_B.png")
create_barchart_for_means(is_s_graph, "Kernel IS (with buckets) Class S", "is/is_S.png")
create_barchart_for_means(is_w_graph, "Kernel IS (with buckets) Class W", "is/is_W.png")
create_barchart_for_means(is_a_graph, "Kernel IS (with buckets) Class A", "is/is_A.png")
create_barchart_for_means(is_b_graph, "Kernel IS (with buckets) Class B", "is/is_B.png")
create_barchart_for_means(is_nobuckets_s_graph, "Kernel IS (no buckets) Class S", "is/is_nobuckets_S.png")
create_barchart_for_means(is_nobuckets_w_graph, "Kernel IS (no buckets) Class W", "is/is_nobuckets_W.png")
create_barchart_for_means(is_nobuckets_a_graph, "Kernel IS (no buckets) Class A", "is/is_nobuckets_A.png")
create_barchart_for_means(is_nobuckets_b_graph, "Kernel IS (no buckets) Class B", "is/is_nobuckets_B.png")