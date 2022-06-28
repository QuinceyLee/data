import torch
import torch.nn as nn


class CNN(nn.Module):
    def __init__(self):
        super(CNN, self).__init__()  # 继承__init__功能
        # 第一层卷积
        self.conv1 = nn.Sequential(
            # 输入[1,3，3]
            nn.Conv2d(
                in_channels=1,  # 输入图片的高度
                out_channels=16,
                kernel_size=(2, 2),
                stride=(1, 1),  # 卷积核在图上滑动，每隔一个扫一次
                padding=1,  # 给图外边补上0
            ),
            # 经过卷积层 输出[16,4,4] 传入池化层
            nn.LeakyReLU()
        )
        # 第二层卷积
        self.conv2 = nn.Sequential(
            nn.Conv2d(
                in_channels=16,
                out_channels=32,
                kernel_size=(2, 2),
                stride=(1, 1),
                padding=0
            ),
            # 经过卷积 输出[32, 3, 3]
            nn.LeakyReLU()
        )
        self.conv3 = nn.Sequential(
            nn.Conv2d(
                in_channels=32,  # 同上
                out_channels=16,
                kernel_size=(2, 2),
                stride=(1, 1),
                padding=0
            ),
            # 经过卷积 输出[64, 2, 2] 传入池化层
            nn.LeakyReLU()
        )
        self.fc1 = nn.Linear(16 * 2 * 2, 32)
        self.fc2 = nn.Linear(96, 64)
        self.leaky = nn.LeakyReLU()
        self.relu = nn.ReLU()

        # 输出层
        self.output = nn.Linear(in_features=64, out_features=7)
        self.softmax = nn.Softmax(dim=1)

    def forward(self, x, y, z):
        x = self.conv1(x)
        x = self.conv2(x)
        x = self.conv3(x)
        y = self.conv1(y)
        y = self.conv2(y)
        y = self.conv3(y)
        z = self.conv1(z)
        z = self.conv2(z)
        z = self.conv3(z)
        x = x.view(x.size(0), -1)  # 保留batch
        y = y.view(y.size(0), -1)
        z = z.view(z.size(0), -1)
        x = self.fc1(x)
        x = self.leaky(x)
        y = self.fc1(y)
        y = self.leaky(y)
        z = self.fc1(z)
        z = self.leaky(z)
        out = torch.cat((x, y, z), 1)
        out = self.fc2(out)
        output = self.output(out)  # 输出[7]
        return output
