// src/NamespaceContext.jsx
import { createContext, useState, useEffect } from 'react';
import axios from 'axios';

export const NamespaceContext = createContext();

export const NamespaceProvider = ({ children }) => {
  const [namespaces, setNamespaces] = useState([]);
  const [selectedNamespace, setSelectedNamespace] = useState('default');

  useEffect(() => {
    const fetchNamespaces = async () => {
      try {
        const response = await axios.get('http://127.0.0.1:8080/api/v1/k8s/namespace');
        setNamespaces(response.data.data);
      } catch (error) {
        console.error('Failed to fetch namespaces:', error);
      }
    };

    fetchNamespaces();
  }, []);

  return (
    <NamespaceContext.Provider value={{ namespaces, selectedNamespace, setSelectedNamespace }}>
      {children}
    </NamespaceContext.Provider>
  );
};
