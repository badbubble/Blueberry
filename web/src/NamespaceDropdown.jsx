// src/NamespaceDropdown.jsx
import  { useContext } from 'react';
import { Menu, Dropdown, Button } from 'antd';
import { DownOutlined } from '@ant-design/icons';
import { NamespaceContext } from './NamespaceContext';

const NamespaceDropdown = () => {
  const { namespaces, selectedNamespace, setSelectedNamespace } = useContext(NamespaceContext);

  const handleMenuClick = (e) => {
    setSelectedNamespace(e.key);
  };

  const menu = (
    <Menu onClick={handleMenuClick}>
      {namespaces.map((ns) => (
        <Menu.Item key={ns.name}>{ns.name}</Menu.Item>
      ))}
    </Menu>
  );

  return (
    <Dropdown overlay={menu}>
      <Button>
        {selectedNamespace} <DownOutlined />
      </Button>
    </Dropdown>
  );
};

export default NamespaceDropdown;
