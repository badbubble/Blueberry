import { useEffect, useState } from 'react';
import axios from 'axios';
import { Table, Spin, Alert } from 'antd';

const NodeList = () => {
  const [nodes, setNodes] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchNodes = async () => {
      try {
        const response = await axios.get('http://127.0.0.1:8080/api/v1/k8s/node');
        setNodes(response.data.data);
        setLoading(false);
      } catch (error) {
        setError(error);
        setLoading(false);
      }
    };

    fetchNodes();
  }, []);

  if (loading) {
    return <Spin tip="Loading..." />;
  }

  if (error) {
    return <Alert message="Error" description={error.message} type="error" showIcon />;
  }

  const columns = [
    {
      title: 'Name',
      dataIndex: 'name',
      key: 'name',
    },
    {
      title: 'Status',
      dataIndex: 'status',
      key: 'status',
    },
    {
      title: 'Age',
      dataIndex: 'age',
      key: 'age',
    },
    {
      title: 'Internal IP',
      dataIndex: 'internalIp',
      key: 'internalIp',
    },
    {
      title: 'External IP',
      dataIndex: 'externalIp',
      key: 'externalIp',
    },
    {
      title: 'Version',
      dataIndex: 'version',
      key: 'version',
    },
    {
      title: 'OS Image',
      dataIndex: 'osImage',
      key: 'osImage',
    },
    {
      title: 'Kernel Version',
      dataIndex: 'kernelVersion',
      key: 'kernelVersion',
    },
    {
      title: 'Container Runtime',
      dataIndex: 'containerRuntime',
      key: 'containerRuntime',
    },
  ];

  return (
    <div style={{ padding: '20px' }}>
      <h1 style={{ textAlign: 'center' }}>K8S Node List</h1>
      <Table dataSource={nodes} columns={columns} rowKey="name" />
    </div>
  );
};

export default NodeList;