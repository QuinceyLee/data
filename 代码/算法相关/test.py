import numpy as np
import pandas as pd
import torch
from matplotlib import pyplot as plt
from pycm import ConfusionMatrix
from sklearn.metrics import accuracy_score, precision_score, recall_score, f1_score, classification_report
from sklearn.preprocessing import StandardScaler
import torch.utils.data as Data
from pretty_confusion_matrix import pp_matrix_from_data

from newporcess.confusion import plot_confusion_matrix
from newporcess.model import MyModel
from plot import get_confusion_matrix

model = MyModel()
model.load_state_dict(torch.load('../data/drug/cnn78.pkl'))
with torch.no_grad():
    test = pd.read_csv('./input_test.csv', header=None)
    batch_size = 100
    ss = StandardScaler()
    x_train = test.values[:, :-1].astype(float)
    x_train = ss.fit_transform(x_train)
    x_train = np.reshape(x_train, (len(x_train), 7, 7))
    print(test.iloc[:, -1].value_counts())
    y_train = test.values[:, -1].astype(int)

    x_train, y_train = torch.FloatTensor(x_train), torch.LongTensor(y_train)
    x_train = torch.unsqueeze(x_train, dim=1)
    train_dataset = torch.utils.data.TensorDataset(x_train, y_train)

    train_loader = Data.DataLoader(dataset=train_dataset, batch_size=batch_size, num_workers=0, shuffle=True)
    x_train, y_train = torch.FloatTensor(x_train), torch.LongTensor(y_train)

    test_pred = []
    test_trues = []
    for i, (test_data_batch, test_data_label) in enumerate(train_loader):
        outputs = model(test_data_batch)
        test_outputs = outputs.argmax(dim=1)
        test_pred.extend(test_outputs.detach().cpu().numpy())
        test_trues.extend(test_data_label.detach().cpu().numpy())
    sklearn_accuracy = accuracy_score(test_trues, test_pred)
    sklearn_precision = precision_score(test_trues, test_pred, average='micro')
    sklearn_recall = recall_score(test_trues, test_pred, average='micro')
    sklearn_f1 = f1_score(test_trues, test_pred, average='micro')
    print(classification_report(test_trues, test_pred))
    conf_matrix = get_confusion_matrix(test_trues, test_pred)
    print(conf_matrix)
    types = range(0, 4)
    plot_confusion_matrix(conf_matrix, types, "three channel CNN")
    plt.savefig('drug_TC-CNN.png', format='png', dpi=300)
    plt.show()
    tot_loss = 0
    print(
        "[ sklearn_metrics for test ] "
        "Epoch:{} "
        "loss:{:.4f} "
        "accuracy:{:.4f} "
        "precision:{:.4f} "
        "recall:{:.4f} "
        "f1:{:.4f}".format(
            i,
            tot_loss,
            sklearn_accuracy,
            sklearn_precision,
            sklearn_recall,
            sklearn_f1))
