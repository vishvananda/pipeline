#!/usr/bin/env python

import collections
import os
import os.path

import tensorflow as tf

from tensorflow.contrib.keras import backend as K
from tensorflow.contrib.keras import models
from tensorflow.python.framework import graph_io
from tensorflow.python.framework import graph_util
from tensorflow.python.keras.applications.vgg16 import VGG16


def write_graph(graph, fname):
    d, f = os.path.split(os.path.abspath(fname))
    graph_io.write_graph(graph, d, f, as_text=False)


def constantize(fname):
    K.clear_session()
    tf.reset_default_graph()
    K.set_learning_phase(False)
    mod = models.load_model(fname)
    outputs = mod.output
    if not isinstance(outputs, collections.Sequence):
        outputs = [outputs]
    output_names = []
    for output in outputs:
        output_names.append(output.name.split(':')[0])
    sess = K.get_session()
    cg = graph_util.convert_variables_to_constants(
        sess, sess.graph.as_graph_def(add_shapes=True), output_names)
    K.clear_session()
    return cg


def h5_to_pb(h5, pb):
    write_graph(constantize(h5), pb)

h5 = "vgg.h5"
pb = "vgg.pb"
model = VGG16(weights="imagenet")
model.save(h5)
h5_to_pb(h5, pb)
