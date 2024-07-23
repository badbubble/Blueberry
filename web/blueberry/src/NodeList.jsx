import { useEffect, useState } from 'react';
import axios from 'axios';
import './NodeList.css';

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
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Error: {error.message}</div>;
  }

  return (
    <div className="node-list">
      <h1>K8S Node List</h1>
      <table>
        <thead>
          <tr>
            <th>Name</th>
            <th>Status</th>
            <th>Age</th>
            <th>Internal IP</th>
            <th>External IP</th>
            <th>Version</th>
            <th>OS Image</th>
            <th>Kernel Version</th>
            <th>Container Runtime</th>
          </tr>
        </thead>
        <tbody>
          {nodes.map((node) => (
            <tr key={node.name}>
              <td>{node.name}</td>
              <td>{node.status}</td>
              <td>{node.age}</td>
              <td>{node.internalIp}</td>
              <td>{node.externalIp}</td>
              <td>{node.version}</td>
              <td>{node.osImage}</td>
              <td>{node.kernelVersion}</td>
              <td>{node.containerRuntime}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default NodeList;
