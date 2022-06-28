const in_types = {
  NATS_SERVER: [
    {
      type: 'el-input-text',
      name: 'host',
      label: '服务地址',
      value: '127.0.0.1',
      required: true,
      placeholder: '127.0.0.1'
    },
    {
      type: 'el-input-number',
      name: 'port',
      label: '服务端口',
      value: 4222,
      required: true,
      placeholder: '4222'
    },
    {
      type: 'el-input-text',
      name: 'user',
      label: '用户名称',
      value: 'rulex',
      required: true,
      placeholder: 'rulex'
    },
    {
      type: 'el-input-text',
      name: 'password',
      label: '用户密码',
      value: '',
      required: true,
      placeholder: ''
    },
    {
      type: 'el-input-text',
      name: 'topic',
      label: '订阅主题',
      value: 'rulex',
      required: true,
      placeholder: 'rulex'
    },
  ],
  // 西门子S7
  SIEMENS_S7: [
    {
      type: 'el-input-text',
      name: 'host',
      label: '服务地址',
      value: '127.0.0.1:1502',
      required: true,
      placeholder: '127.0.0.1:1502'
    },
    {
      type: 'el-input-number',
      name: 'rack',
      label: '机架编号',
      value: 0,
      required: true,
      placeholder: '0'
    },
    {
      type: 'el-input-number',
      name: 'slot',
      label: '插槽编号',
      value: 1,
      required: true,
      placeholder: '1'
    },
    {
      type: 'el-input-text',
      name: 'model',
      label: '型号选择',
      value: 'S7-200',
      required: true,
      placeholder: 'S7-200'
    },
    {
      type: 'el-input-number',
      name: 'timeout',
      label: '超时时间',
      value: 10,
      required: true,
      placeholder: '10'
    },
    {
      type: 'el-input-number',
      name: 'idleTimeout',
      label: '心跳时间',
      value: 60,
      required: true,
      placeholder: '60'
    },
    {
      type: 'el-input-number',
      name: 'frequency',
      label: '采集频率',
      value: 5,
      required: true,
      placeholder: '5'
    },
    {
      type: 'dynamicParams',
      label: 'DB配置',
      name: 'dbs',
      required: true,
      children: [
        {
          type: 'el-input-text',
          name: 'tag',
          label: '数据标签',
          value: 'A',
          required: true,
          placeholder: 'A'
        },
        {
          type: 'el-input-number',
          name: 'address',
          label: '地址编号',
          value: 0,
          required: true,
          placeholder: '0'
        },
        {
          type: 'el-input-number',
          name: 'start',
          label: '起始地址',
          value: 1,
          required: true,
          placeholder: '1'
        },
        {
          type: 'el-input-number',
          name: 'size',
          label: '读取数量',
          value: 1,
          required: true,
          placeholder: '1'
        }
      ]
    },
    {
      type: 'flag',
      name: 'isDynamic',
      label: '是否动态',
      value: '1',
      required: false,
      placeholder: '1'
    }
  ],

  MQTT: [
    {
      type: 'el-input-text',
      name: 'host',
      label: '主机地址',
      value: '127.0.0.1',
      required: true,
      placeholder: '127.0.0.1'
    },
    {
      type: 'el-input-number',
      name: 'port',
      label: '服务端口',
      value: 1883,
      required: true,
      placeholder: '1883'
    },
    {
      type: 'el-input-text',
      name: 'topic',
      label: '订阅主题',
      value: 'rulex-data-1',
      required: true,
      placeholder: 'rulex-data-1'
    },
    {
      type: 'el-input-text',
      name: 'clientId',
      label: '客户端ID',
      value: 'rulex-1',
      required: true,
      placeholder: 'rulex-1'
    },
    {
      type: 'el-input-text',
      name: 'username',
      value: 'rulex-1',
      label: '用  户',
      required: true,
      placeholder: 'rulex-1'
    },
    {
      type: 'el-input-text',
      name: 'password',
      label: '密  码',
      value: '',
      required: true,
      placeholder: 'rulex-1'
    },
    {
      type: 'flag',
      name: 'isDynamic',
      label: '是否动态',
      value: '0',
      required: false,
      placeholder: '0'
    }
    // label: "MQTT桥接",
    // host: "127.0.0.1",
    // port: 1883,
    // topic: "rulex-client-topic-1",
    // clientId: "rulex1",
    // username: "rulex1",
    // password: "******"
  ],

  HTTP: [
    {
      type: 'el-input-number',
      name: 'port',
      label: '服务端口',
      value: 2581,
      required: true,
      placeholder: '2581'
    },
    {
      type: 'flag',
      name: 'isDynamic',
      label: '是否动态',
      value: '0',
      required: false,
      placeholder: '0'
    }
    // label: 'HTTP桥接',
    // port: 2581
  ],

  RULEX_UDP: [
    {
      type: 'el-input-text',
      name: 'host',
      label: '服务地址',
      value: '127.0.0.1',
      required: true,
      placeholder: '127.0.0.1'
    },
    {
      type: 'el-input-number',
      name: 'port',
      label: '服务端口',
      value: 2582,
      required: true,
      placeholder: '2582'
    },
    {
      type: 'el-input-number',
      name: 'maxDataLength',
      label: '最大包长',
      value: 1024,
      required: true,
      placeholder: '1024'
    },
    {
      type: 'flag',
      name: 'isDynamic',
      label: '是否动态',
      value: '0',
      required: false,
      placeholder: '0'
    }
    // label: 'UDP桥接',
    // port: 2582
    // maxDataLength: 1024
  ],

  COAP: [
    {
      type: 'el-input-number',
      name: 'port',
      label: '服务端口',
      placeholder: '2583',
      required: true,
      value: 2583
    },
    {
      type: 'flag',
      name: 'isDynamic',
      label: '是否动态',
      value: '0',
      required: false,
      placeholder: '0'
    }
    // label: 'COAP桥接',
    // port: 2583
  ],

  GRPC: [
    {
      type: 'el-input-number',
      name: 'port',
      label: '服务端口',
      value: 2584,
      required: true,
      placeholder: '2584'
    },
    {
      type: 'flag',
      name: 'isDynamic',
      label: '是否动态',
      value: '0',
      required: false,
      placeholder: '0'
    }
    // label: 'GRPC桥接',
    // port: 2584
  ],

  UART_MODULE: [
    {
      type: 'el-input-text',
      name: 'address',
      label: '串口名称',
      value: '/dev/ttyS1',
      required: true,
      placeholder: '/dev/ttyS1'
    },
    {
      type: 'el-input-number',
      name: 'baudRate',
      label: '波特率',
      value: 115200,
      required: true,
      placeholder: '115200'
    },
    {
      type: 'el-input-number',
      name: 'timeout',
      label: '超时时长',
      value: 5,
      required: true,
      placeholder: '5'
    },
    {
      type: 'el-input-number',
      name: 'dataBits',
      label: '数据位',
      value: 8,
      required: true,
      placeholder: '8'
    },
    {
      type: 'el-select',
      name: 'parity',
      label: '奇偶校验',
      value: 'N',
      required: true,
      placeholder: '请选择',
      selectOptions: [
        { value: 'N', label: '无校验' },
        { value: 'E', label: '偶校验' },
        { value: 'O', label: '奇校验' }
      ]
    },
    {
      type: 'el-input-number',
      name: 'stopBits',
      label: '停止位',
      value: 1,
      required: true,
      placeholder: '1'
    },
    {
      type: 'flag',
      name: 'isDynamic',
      label: '是否动态',
      value: '0',
      required: false,
      placeholder: '0'
    }
    // label: '通用串口',
    // address: 'uart1',
    // baudRate: 115200,
    // timeout: 5,
    // dataBits: 8,
    // parity: 'N',
    // stopBits: 1
  ],

  MODBUS_MASTER: [
    {
      type: 'el-select',
      name: 'mode',
      label: '工作模式',
      value: '',
      required: true,
      placeholder: '请选择',
      selectOptions: [
        { value: 'TCP', label: 'TCP' },
        { value: 'RTU', label: 'RTU' }
      ]
    },
    {
      type: 'flag',
      name: 'isDynamic',
      label: '是否动态',
      value: '0',
      required: false,
      placeholder: '0'
    }
  ],

  SNMP_SERVER: [
    {
      type: 'el-input-number',
      name: 'frequency',
      label: '采集频率',
      value: 5,
      required: true,
      placeholder: '5'
    },
    {
      type: 'el-input-number',
      name: 'timeout',
      label: '超时时长',
      value: 5,
      required: true,
      placeholder: '5'
    },
    {
      type: 'el-inline',
      label: '目标配置',
      name: 'targets',
      children: [
        {
          type: 'el-input-text',
          name: 'target',
          label: '主机地址',
          value: '127.0.0.1',
          required: true,
          placeholder: '127.0.0.1'
        },
        {
          type: 'el-input-number',
          name: 'port',
          label: '主机端口',
          value: 1161,
          required: true,
          placeholder: '1161'
        },
        {
          type: 'el-select',
          name: 'transport',
          label: '传输模式',
          value: 'udp',
          required: true,
          placeholder: 'udp',
          selectOptions: [{ value: 'udp', label: 'UDP' }]
        },
        {
          type: 'el-input-text',
          name: 'community',
          label: '社区名称',
          value: 'public',
          required: true,
          placeholder: 'public'
        },
        {
          type: 'el-input-number',
          name: 'version',
          label: '协议版本',
          value: 3,
          required: true,
          placeholder: '3'
        },
        {
          type: 'el-input-text',
          name: 'dataModels',
          label: '数据模型',
          value: '',
          required: true,
          placeholder: ''
        }
      ]
    },
    {
      type: 'flag',
      name: 'isDynamic',
      label: '是否动态',
      value: '0',
      required: false,
      placeholder: '0'
    }
    // label: "SNMP协议",
    // frequency: 5,
    // timeout: 3,
    // targets: [
    //   {
    //     target: "127.0.0.1",
    //     port: 161,
    //     transport: "udp",
    //     community: "public",
    //     version: 3,
    //     dataModels: [],
    //   },
    // ],
  ],

  TCP: [
    {
      type: 'el-select',
      name: 'mode',
      label: '工作模式',
      value: 'TCP',
      required: true,
      placeholder: '请选择',
      selectOptions: [
        { value: 'TCP', label: 'TCP' },
        { value: 'RTU', label: 'RTU' }
      ]
    },
    {
      type: 'el-input-number',
      name: 'timeout',
      label: '超时时长',
      value: 5,
      required: true,
      placeholder: '5'
    },
    {
      type: 'el-input-number',
      name: 'slaverId',
      label: '从机地址',
      value: 1,
      required: true,
      placeholder: '1'
    },
    {
      type: 'el-input-number',
      name: 'frequency',
      label: '采集频率',
      value: 5,
      required: true,
      placeholder: '5'
    },
    {
      type: 'el-inline',
      label: 'TCP配置',
      name: 'config',
      children: [
        {
          type: 'el-input-text',
          name: 'ip',
          label: '采集地址',
          value: '127.0.0.1',
          required: true,
          placeholder: '127.0.0.1'
        },
        {
          type: 'el-input-number',
          name: 'port',
          label: '采集端口',
          value: 1254,
          required: true,
          placeholder: '1254'
        }
      ]
    },
    {
      type: 'dynamicParams',
      label: '寄存器配置',
      name: 'registerParams',
      required: false,
      children: [
        {
          type: 'el-input-text',
          name: 'tag',
          label: '数据标签',
          value: 'A',
          required: true,
          placeholder: 'A'
        },
        {
          type: 'el-input-number',
          name: 'function',
          label: '功能代码',
          value: 1,
          required: true,
          placeholder: '1'
        },
        {
          type: 'el-input-number',
          name: 'address',
          label: '起始地址',
          value: 0,
          required: true,
          placeholder: '0'
        },
        {
          type: 'el-input-number',
          name: 'quantity',
          label: '读取数量',
          value: 10,
          required: true,
          placeholder: '10'
        }
      ]
    },
    {
      type: 'flag',
      name: 'isDynamic',
      label: '是否动态',
      value: '1',
      required: false,
      placeholder: '1'
    }
  ],

  RTU: [
    {
      type: 'el-select',
      name: 'mode',
      label: '工作模式',
      value: 'RTU',
      required: true,
      placeholder: '请选择',
      selectOptions: [
        { value: 'TCP', label: 'TCP' },
        { value: 'RTU', label: 'RTU' }
      ]
    },
    {
      type: 'el-input-number',
      name: 'timeout',
      label: '超时时长',
      value: 5,
      required: true,
      placeholder: '5'
    },
    {
      type: 'el-input-number',
      name: 'slaverId',
      label: '从机地址',
      value: 1,
      required: true,
      placeholder: '1'
    },
    {
      type: 'el-input-number',
      name: 'frequency',
      label: '采集频率',
      value: 5,
      required: true,
      placeholder: '5'
    },
    {
      type: 'el-inline',
      label: '配置',
      name: 'config',
      children: [
        {
          type: 'el-input-text',
          name: 'uart',
          label: '串口名称',
          value: 'com1',
          required: true,
          placeholder: 'com1'
        },
        {
          type: 'el-input-number',
          name: 'baudRate',
          label: '波特率',
          value: 115200,
          required: true,
          placeholder: '115200'
        },
        {
          type: 'el-input-number',
          name: 'dataBits',
          label: '数据位',
          value: 8,
          required: true,
          placeholder: '8'
        },
        {
          type: 'el-select',
          name: 'parity',
          label: '奇偶校验',
          value: 'N',
          required: true,
          placeholder: '请选择',
          selectOptions: [
            { value: 'N', label: '无校验' },
            { value: 'E', label: '偶校验' },
            { value: 'O', label: '奇校验' }
          ]
        },
        {
          type: 'el-input-number',
          name: 'stopBits',
          label: '停止位',
          value: 1,
          required: true,
          placeholder: '1'
        }
      ]
    },
    {
      type: 'dynamicParams',
      label: '寄存器配置',
      required: false,
      name: 'registerParams',
      children: [
        {
          type: 'el-input-text',
          name: 'tag',
          label: '数据标签',
          value: 'A',
          required: true,
          placeholder: 'A'
        },
        {
          type: 'el-input-number',
          name: 'function',
          label: 'Modbus 功能',
          value: 1,
          required: true,
          placeholder: '1'
        },
        {
          type: 'el-input-number',
          name: 'address',
          label: '起始地址',
          value: 0,
          required: true,
          placeholder: '0'
        },
        {
          type: 'el-input-number',
          name: 'quantity',
          label: '读取数量',
          value: 10,
          required: true,
          placeholder: '10'
        }
      ]
    },
    {
      type: 'flag',
      name: 'isDynamic',
      label: '是否动态',
      value: '1',
      required: false,
      placeholder: '1'
    }
  ]
}

export default in_types
