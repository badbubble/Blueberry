// src/App.jsx
import React from 'react';
import { BrowserRouter as Router, Route, Routes, Link } from 'react-router-dom';
import { Layout, Menu } from 'antd';
import {
  DesktopOutlined,
  ContainerOutlined,
} from '@ant-design/icons';
import NodeList from './NodeList';
import PodList from './PodList';
import CreatePod from './CreatePod';
import { NamespaceProvider } from './NamespaceContext';
import NamespaceDropdown from './NamespaceDropdown';

const { Header, Content, Footer, Sider } = Layout;
const { SubMenu } = Menu;

const App = () => {
  return (
    <NamespaceProvider>
      <Router>
        <Layout style={{ minHeight: '100vh' }}>
          <Sider collapsible>
            <div>Blueberry</div>
            <div className="logo" style={{ height: '32px', margin: '16px', background: 'darkred', display: 'flex', alignItems: 'center', justifyContent: 'center', color: '#fff', fontWeight: 'bold' }}>
              Blueberry
            </div>            <Menu theme="dark" defaultSelectedKeys={['1']} mode="inline">
              <SubMenu key="sub1" icon={<DesktopOutlined />} title="Node">
                <Menu.Item key="1">
                  <Link to="/nodes/list">List</Link>
                </Menu.Item>
              </SubMenu>
              <SubMenu key="sub2" icon={<ContainerOutlined />} title="Pod">
                <Menu.Item key="2">
                  <Link to="/pods/list">List</Link>
                </Menu.Item>
                <Menu.Item key="3">
                  <Link to="/pods/create">Create Pod</Link>
                </Menu.Item>
              </SubMenu>
            </Menu>
          </Sider>
          <Layout className="site-layout">
            <Header className="site-layout-background" style={{ padding: 0, display: 'flex', justifyContent: 'flex-end' }}>
              <NamespaceDropdown />
            </Header>
            <Content style={{ margin: '0 16px' }}>
              <Routes>
                <Route path="/nodes/list" element={<NodeList />} />
                <Route path="/pods/list" element={<PodList />} />
                <Route path="/pods/create" element={<CreatePod />} />
              </Routes>
            </Content>
            <Footer style={{ textAlign: 'center' }}>Blueberry Dashboard ©2024</Footer>
          </Layout>
        </Layout>
      </Router>
    </NamespaceProvider>
  );
};

export default App;
