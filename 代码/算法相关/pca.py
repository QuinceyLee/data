import datetime

import numpy as np
import pandas as pd
import joblib
from sklearn.decomposition import PCA


def get_file(mode):
    if mode:
        data0 = pd.read_csv("../data/csv/d00.csv", header=None)
        data1 = pd.read_csv("../data/csv/d01.csv", header=None)
        data2 = pd.read_csv("../data/csv/d02.csv", header=None)
        data3 = pd.read_csv("../data/csv/d03.csv", header=None)
        data4 = pd.read_csv("../data/csv/d04.csv", header=None)
        data5 = pd.read_csv("../data/csv/d05.csv", header=None)
        data6 = pd.read_csv("../data/csv/d06.csv", header=None)
        data7 = pd.read_csv("../data/csv/d07.csv", header=None)
        data8 = pd.read_csv("../data/csv/d08.csv", header=None)
        data9 = pd.read_csv("../data/csv/d09.csv", header=None)
        data10 = pd.read_csv("../data/csv/d10.csv", header=None)
        data11 = pd.read_csv("../data/csv/d11.csv", header=None)
        data12 = pd.read_csv("../data/csv/d12.csv", header=None)
        data13 = pd.read_csv("../data/csv/d13.csv", header=None)
        data14 = pd.read_csv("../data/csv/d14.csv", header=None)
        data15 = pd.read_csv("../data/csv/d15.csv", header=None)
        data16 = pd.read_csv("../data/csv/d16.csv", header=None)
        data17 = pd.read_csv("../data/csv/d17.csv", header=None)
        data18 = pd.read_csv("../data/csv/d18.csv", header=None)
        data19 = pd.read_csv("../data/csv/d19.csv", header=None)
        data20 = pd.read_csv("../data/csv/d20.csv", header=None)
        data21 = pd.read_csv("../data/csv/d21.csv", header=None)
    else:
        data0 = pd.read_csv("../data/csv/d00_te.csv", header=None)
        data1 = pd.read_csv("../data/csv/d01_te.csv", header=None)
        data2 = pd.read_csv("../data/csv/d02_te.csv", header=None)
        data3 = pd.read_csv("../data/csv/d03_te.csv", header=None)
        data4 = pd.read_csv("../data/csv/d04_te.csv", header=None)
        data5 = pd.read_csv("../data/csv/d05_te.csv", header=None)
        data6 = pd.read_csv("../data/csv/d06_te.csv", header=None)
        data7 = pd.read_csv("../data/csv/d07_te.csv", header=None)
        data8 = pd.read_csv("../data/csv/d08_te.csv", header=None)
        data9 = pd.read_csv("../data/csv/d09_te.csv", header=None)
        data10 = pd.read_csv("../data/csv/d10_te.csv", header=None)
        data11 = pd.read_csv("../data/csv/d11_te.csv", header=None)
        data12 = pd.read_csv("../data/csv/d12_te.csv", header=None)
        data13 = pd.read_csv("../data/csv/d13_te.csv", header=None)
        data14 = pd.read_csv("../data/csv/d14_te.csv", header=None)
        data15 = pd.read_csv("../data/csv/d15_te.csv", header=None)
        data16 = pd.read_csv("../data/csv/d16_te.csv", header=None)
        data17 = pd.read_csv("../data/csv/d17_te.csv", header=None)
        data18 = pd.read_csv("../data/csv/d18_te.csv", header=None)
        data19 = pd.read_csv("../data/csv/d19_te.csv", header=None)
        data20 = pd.read_csv("../data/csv/d20_te.csv", header=None)
        data21 = pd.read_csv("../data/csv/d21_te.csv", header=None)
    file = [data0, data1, data2, data3, data4, data5, data6, data7,
            data8, data9, data10, data11, data12, data13, data14,
            data15, data16, data17, data18, data19, data20, data21]
    return file


def get_train():
    pca = joblib.load("../data/model/pca.joblib")
    file = get_file(True)
    count = 0
    for i, element in enumerate(file):
        element = pca.transform(element)
        element = pd.DataFrame(element)
        element.insert(element.shape[1], 'out', count)
        file[i] = element
        count = count + 1
    train = pd.concat(file)
    normal_dic = {
        2: 1, 3: 1, 4: 1, 5: 1, 6: 1, 7: 1,
        8: 2, 9: 2, 10: 2, 11: 2, 12: 2,
        13: 3,
        14: 4, 15: 4,
        16: 5, 17: 5, 18: 5, 19: 5, 20: 5,
        21: 6
    }
    new = train.replace({'out': normal_dic})
    new.to_csv('pca.csv', header=None, index=None)


def get_test():
    pca = joblib.load("../data/model/pca.joblib")
    file = get_file(False)
    count = 0
    for i, element in enumerate(file):
        element = pca.transform(element)
        element = pd.DataFrame(element)
        if i != 0:
            normal = np.zeros((1, 160))
            abnormal = np.full((1, 800), count)
            out = np.hstack((normal, abnormal))
            element.insert(element.shape[1], 'out', out.T)
        else:
            element.insert(element.shape[1], 'out', count)
        file[i] = element
        count = count + 1
    train = pd.concat(file)
    normal_dic = {
        2: 1, 3: 1, 4: 1, 5: 1, 6: 1, 7: 1,
        8: 2, 9: 2, 10: 2, 11: 2, 12: 2,
        13: 3,
        14: 4, 15: 4,
        16: 5, 17: 5, 18: 5, 19: 5, 20: 5,
        21: 6
    }
    new = train.replace({'out': normal_dic})
    new.to_csv('pca_test.csv', header=None, index=None)

get_train()
