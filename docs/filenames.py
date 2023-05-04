titles = [
    "Dalvan (C++)",
    "NAS (Fortran)",
    "NPB-Golang"
]

titles_nobuckets = [
    "NAS (Fortran)",
    "NPB-Golang"
]

ep_s = [
    'ep/ep_S_dalvan.txt',
    'ep/ep_S_nasa.txt',
    'ep/ep.s.txt'
]

ep_a = [
    'ep/ep_A_dalvan.txt',
    'ep/ep_A_nasa.txt',
    'ep/ep.a.txt'
]

ep_b = [
    'ep/ep_B_dalvan.txt',
    'ep/ep_B_nasa.txt',
    'ep/ep.b.txt'
]

is_s = [
    'is/dalvan_is_S.txt',
    'is/nas_is_S.txt',
    'is/IS_S.txt'
]

is_w = [
    'is/dalvan_is_W.txt',
    'is/nas_is_W.txt',
    'is/IS_W.txt'
]

is_a = [
    'is/dalvan_is_A.txt',
    'is/nas_is_A.txt',
    'is/IS_A.txt'
]

is_b = [
    'is/dalvan_is_B.txt',
    'is/nas_is_B.txt',
    'is/IS_B.txt'
]

is_nobuckets_s = [
    'is/is_S_NAS_NOBUCKETS.txt',
    'is/IS_S_NOBUCKETS.txt'
]

is_nobuckets_w = [
    'is/is_W_NAS_NOBUCKETS.txt',
    'is/IS_W_NOBUCKETS.txt'
]

is_nobuckets_a = [
    'is/is_A_NAS_NOBUCKETS.txt',
    'is/IS_A_NOBUCKETS.txt'
]

is_nobuckets_b = [
    'is/is_B_NAS_NOBUCKETS.txt',
    'is/IS_B_NOBUCKETS.txt'
]

ep_s_graph = dict(zip(titles, ep_s))
ep_a_graph = dict(zip(titles, ep_a))
ep_b_graph = dict(zip(titles, ep_b))
is_s_graph = dict(zip(titles, is_s))
is_w_graph = dict(zip(titles, is_w))
is_a_graph = dict(zip(titles, is_a))
is_b_graph = dict(zip(titles, is_b))
is_nobuckets_s_graph = dict(zip(titles_nobuckets, is_nobuckets_s))
is_nobuckets_w_graph = dict(zip(titles_nobuckets, is_nobuckets_w))
is_nobuckets_a_graph = dict(zip(titles_nobuckets, is_nobuckets_a))
is_nobuckets_b_graph = dict(zip(titles_nobuckets, is_nobuckets_b))

ep_graph = [ep_s_graph, ep_a_graph, ep_b_graph]
is_graph = [is_s_graph, is_w_graph, is_a_graph, is_b_graph]
is_nobuckets_graph = [is_nobuckets_s_graph, is_nobuckets_w_graph, is_nobuckets_a_graph, is_nobuckets_b_graph]