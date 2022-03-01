'''
xxx
'''

import numpy
import operator


def create_dataset():
    '''create_dataset...
    '''
    group = numpy.array([
        [1.0, 1.1],
        [1.0, 1.0],
        [0, 0],
        [0, 0.1]
    ])
    labels = ['A', 'A', 'B', 'B']
    return group, labels


def classify0(inX, dataSet, labels, k):
    '''classify0...
    '''
    dataset_size = dataSet.shape[0]
    diff_mat = numpy.tile(inX, (dataset_size, 1)) - dataSet
    sq_diff_mat = diff_mat ** 2
    sq_distances = sq_diff_mat.sum(axis=1)
    distances = sq_distances ** 0.5
    sorted_dist_indicies = distances.argsort()
    class_count = {}
    for i in range(k):
        vote_i_label = labels[sorted_dist_indicies[i]]
        class_count[vote_i_label] = class_count.get(vote_i_label, 0) + 1
    sorted_class_count = sorted(class_count.items(), key=operator.itemgetter(1), reverse=True)
    return sorted_class_count[0][0]


group, labels = create_dataset()
print(classify0([0, 0], group, labels, 3))
