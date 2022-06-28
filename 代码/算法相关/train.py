import os

import numpy as np
import pandas as pd
import torch
from sklearn.metrics import accuracy_score, precision_score, recall_score, f1_score, classification_report
from sklearn.model_selection import train_test_split
from sklearn.preprocessing import StandardScaler
from torch import nn
import torch.utils.data as Data
from newporcess.model import MyModel
from plot import plot_confusion_matrix, get_confusion_matrix

epochs = 1000
batch_size = 200
learning_rate = 0.01

device = torch.device("cuda:0" if torch.cuda.is_available() else "cpu")
# 将训练数据装入Loader中
cnn = MyModel()
optimizer = torch.optim.Adam(cnn.parameters(), lr=learning_rate)  # optimize all cnn parameters
loss_func = nn.CrossEntropyLoss()
print("开始训练主循环")
dt = pd.read_csv('./input.csv', header=None)

train, test = train_test_split(dt, test_size=0.3)
test.to_csv('input_test.csv', header=None, index=None)
train.to_csv('input_train.csv', header=None, index=None)
print(test.iloc[:, -1].value_counts())
print(train.iloc[:, -1].value_counts())

ss = StandardScaler()
x_train = test.values[:, :-1].astype(float)
x_train = ss.fit_transform(x_train)
x_train = np.reshape(x_train, (len(x_train), 7, 7))
y_train = test.values[:, -1].astype(int)

x_train, y_train = torch.FloatTensor(x_train).to(device), torch.LongTensor(y_train).to(device)
x_train = torch.unsqueeze(x_train, dim=1)
train_dataset = torch.utils.data.TensorDataset(x_train, y_train)

train_loader = Data.DataLoader(dataset=train_dataset, batch_size=batch_size, num_workers=0, shuffle=True)
x_train, y_train = torch.FloatTensor(x_train).to(device), torch.LongTensor(y_train).to(device)

tot_loss = 0.0

cnn.train()
for number in range(epochs):
    train_pred = []
    train_trues = []
    for i, (train_data_batch, train_data_label) in enumerate(train_loader):
        # model.train()
        outputs = cnn(train_data_batch)
        # _, preds = torch.max(outputs.data, 1)
        loss = loss_func(outputs, train_data_label)
        # print(loss)
        # 反向传播优化网络参数
        loss.backward()
        optimizer.step()
        optimizer.zero_grad()
        # 累加每个step的损失
        tot_loss += loss.data
        train_outputs = outputs.argmax(dim=1)
        train_pred.extend(train_outputs.detach().cpu().numpy())
        train_trues.extend(train_data_label.detach().cpu().numpy())
        # tot_acc += (outputs.argmax(dim=1) == train_label_batch).sum().item()
    sklearn_accuracy = accuracy_score(train_trues, train_pred)
    sklearn_precision = precision_score(train_trues, train_pred, average='micro')
    sklearn_recall = recall_score(train_trues, train_pred, average='micro')
    sklearn_f1 = f1_score(train_trues, train_pred, average='micro')
    print(
        "[ sklearn_metrics for train ] "
        "Epoch:{} "
        "loss:{:.4f} "
        "accuracy:{:.4f} "
        "precision:{:.4f} "
        "recall:{:.4f} "
        "f1:{:.4f}".format(
            number,
            tot_loss,
            sklearn_accuracy,
            sklearn_precision,
            sklearn_recall,
            sklearn_f1))
    torch.save(cnn.state_dict(), '../data/drug/cnn' + str(number) + '.pkl')

