// src/PodList.jsx
import  { useEffect, useState, useContext } from 'react';
import axios from 'axios';
import {Table, Spin, Alert, Button, Space, message} from 'antd';
import { NamespaceContext } from './NamespaceContext';

const PodList = () => {
  const { selectedNamespace } = useContext(NamespaceContext);
  const [pods, setPods] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchPods = async () => {
      setLoading(true);
      try {
        const response = await axios.get(`http://127.0.0.1:8080/api/v1/k8s/pod?namespace=${selectedNamespace}`);
        setPods(response.data.data);
        setLoading(false);
      } catch (error) {
        setError(error);
        setLoading(false);
      }
    };

    fetchPods();
  }, [selectedNamespace]);

  const handleUpdate = (name) => {
    console.log(`Update ${name}`);
    // Add your update logic here
  };

  const handleStop = (name) => {
    console.log(`Stop ${name}`);
    // Add your stop logic here
  };

  const handleDelete = async (name) => {
    try {
      await axios.delete(`http://127.0.0.1:8080/api/v1/k8s/pod?namespace=${selectedNamespace}&name=${name}`);
      message.success(`Pod ${name} deleted successfully`);
      // Refresh the list after deletion
      setPods(pods.filter(pod => pod.name !== name));
    } catch (error) {
      message.error(`Failed to delete pod ${name}`);
      console.error('Error deleting pod:', error);
    }
  };

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
      title: 'Ready',
      dataIndex: 'ready',
      key: 'ready',
    },
    {
      title: 'Status',
      dataIndex: 'status',
      key: 'status',
    },
    {
      title: 'Restarts',
      dataIndex: 'restarts',
      key: 'restarts',
    },
    {
      title: 'Age',
      dataIndex: 'age',
      key: 'age',
    },
    {
      title: 'IP',
      dataIndex: 'IP',
      key: 'IP',
    },
    {
      title: 'Node',
      dataIndex: 'node',
      key: 'node',
    },
    {
      title: 'Action',
      key: 'action',
      render: (_, record) => (
        <Space size="middle">
          <Button
            type="primary"
            style={{ backgroundColor: 'green', borderColor: 'green' }}
            onClick={() => handleUpdate(record.name)}
          >
            Update
          </Button>
          <Button
            type="default"
            style={{ backgroundColor: 'yellow', borderColor: 'yellow', color: 'black' }}
            onClick={() => handleStop(record.name)}
          >
            Stop
          </Button>
          <Button
            type="danger"
            style={{ backgroundColor: 'red', borderColor: 'red' }}
            onClick={() => handleDelete(record.name)}
          >
            Delete
          </Button>
        </Space>
      ),
    },
  ];

  return (
    <div style={{ padding: '20px' }}>
      <h1 style={{ textAlign: 'center' }}>K8S Pod List</h1>
      <Table dataSource={pods} columns={columns} rowKey="name" />
    </div>
  );
};

export default PodList;
