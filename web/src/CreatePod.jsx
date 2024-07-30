// src/CreatePod.jsx
import { useContext } from 'react';
import { Form, Input, Button, Space, Select, Checkbox } from 'antd';
import axios from 'axios';
import { NamespaceContext } from './NamespaceContext';

const { Option } = Select;

const CreatePod = () => {
  const { selectedNamespace } = useContext(NamespaceContext);
  const [form] = Form.useForm();

  const onFinish = async (values) => {
    const podData = {
      base: {
        name: values.name,
        namespace: selectedNamespace,
        labels: values.labels,
        restartPolicy: values.restartPolicy,
      },
      volumes: values.volumes,
      netWorking: values.netWorking,
      initContainers: values.initContainers,
      containers: values.containers,
    };

    try {
      const response = await axios.post('http://127.0.0.1:8080/api/v1/k8s/pod', podData);
      console.log('Pod created successfully:', response.data);
    } catch (error) {
      console.error('Error creating pod:', error);
    }
  };

  return (
    <Form form={form} onFinish={onFinish} layout="vertical">
      <Form.Item name="name" label="Pod Name" rules={[{ required: true }]}>
        <Input />
      </Form.Item>
      <Form.Item name="labels" label="Labels">
        <Form.List name="labels">
          {(fields, { add, remove }) => (
            <>
              {fields.map((field) => (
                <Space key={field.key} align="baseline">
                  <Form.Item
                    {...field}
                    name={[field.name, 'key']}
                    fieldKey={[field.fieldKey, 'key']}
                    rules={[{ required: true, message: 'Missing key' }]}
                  >
                    <Input placeholder="Key" />
                  </Form.Item>
                  <Form.Item
                    {...field}
                    name={[field.name, 'value']}
                    fieldKey={[field.fieldKey, 'value']}
                    rules={[{ required: true, message: 'Missing value' }]}
                  >
                    <Input placeholder="Value" />
                  </Form.Item>
                  <Button onClick={() => remove(field.name)}>Remove</Button>
                </Space>
              ))}
              <Button type="dashed" onClick={() => add()} block>
                Add Label
              </Button>
            </>
          )}
        </Form.List>
      </Form.Item>
      <Form.Item name="restartPolicy" label="Restart Policy" rules={[{ required: true }]}>
        <Select>
          <Option value="Always">Always</Option>
          <Option value="OnFailure">OnFailure</Option>
          <Option value="Never">Never</Option>
        </Select>
      </Form.Item>
      <Form.Item name="volumes" label="Volumes">
        <Form.List name="volumes">
          {(fields, { add, remove }) => (
            <>
              {fields.map((field) => (
                <Space key={field.key} align="baseline">
                  <Form.Item
                    {...field}
                    name={[field.name, 'name']}
                    fieldKey={[field.fieldKey, 'name']}
                    rules={[{ required: true, message: 'Missing volume name' }]}
                  >
                    <Input placeholder="Volume Name" />
                  </Form.Item>
                  <Form.Item
                    {...field}
                    name={[field.name, 'type']}
                    fieldKey={[field.fieldKey, 'type']}
                    rules={[{ required: true, message: 'Missing volume type' }]}
                  >
                    <Input placeholder="Volume Type" />
                  </Form.Item>
                  <Button onClick={() => remove(field.name)}>Remove</Button>
                </Space>
              ))}
              <Button type="dashed" onClick={() => add()} block>
                Add Volume
              </Button>
            </>
          )}
        </Form.List>
      </Form.Item>
      <Form.Item name="netWorking" label="Networking">
        <Form.Item name={['netWorking', 'hostNetwork']} valuePropName="checked">
          <Checkbox>Host Network</Checkbox>
        </Form.Item>
        <Form.Item name={['netWorking', 'hostName']} label="Host Name">
          <Input />
        </Form.Item>
        <Form.Item name={['netWorking', 'dnsPolicy']} label="DNS Policy">
          <Select>
            <Option value="Default">Default</Option>
            <Option value="ClusterFirst">ClusterFirst</Option>
            <Option value="ClusterFirstWithHostNet">ClusterFirstWithHostNet</Option>
            <Option value="None">None</Option>
          </Select>
        </Form.Item>
        <Form.Item name={['netWorking', 'dnsConfig']} label="DNS Config">
          <Form.List name={['netWorking', 'dnsConfig', 'nameservers']}>
            {(fields, { add, remove }) => (
              <>
                {fields.map((field) => (
                  <Space key={field.key} align="baseline">
                    <Form.Item
                      {...field}
                      name={[field.name]}
                      fieldKey={[field.fieldKey]}
                      rules={[{ required: true, message: 'Missing nameserver' }]}
                    >
                      <Input placeholder="Nameserver" />
                    </Form.Item>
                    <Button onClick={() => remove(field.name)}>Remove</Button>
                  </Space>
                ))}
                <Button type="dashed" onClick={() => add()} block>
                  Add Nameserver
                </Button>
              </>
            )}
          </Form.List>
        </Form.Item>
        <Form.Item name={['netWorking', 'hostAliases']} label="Host Aliases">
          <Form.List name={['netWorking', 'hostAliases']}>
            {(fields, { add, remove }) => (
              <>
                {fields.map((field) => (
                  <Space key={field.key} align="baseline">
                    <Form.Item
                      {...field}
                      name={[field.name, 'key']}
                      fieldKey={[field.fieldKey, 'key']}
                      rules={[{ required: true, message: 'Missing key' }]}
                    >
                      <Input placeholder="Key" />
                    </Form.Item>
                    <Form.Item
                      {...field}
                      name={[field.name, 'value']}
                      fieldKey={[field.fieldKey, 'value']}
                      rules={[{ required: true, message: 'Missing value' }]}
                    >
                      <Input placeholder="Value" />
                    </Form.Item>
                    <Button onClick={() => remove(field.name)}>Remove</Button>
                  </Space>
                ))}
                <Button type="dashed" onClick={() => add()} block>
                  Add Host Alias
                </Button>
              </>
            )}
          </Form.List>
        </Form.Item>
      </Form.Item>
      <Form.Item name="initContainers" label="Init Containers">
        <Form.List name="initContainers">
          {(fields, { add, remove }) => (
            <>
              {fields.map((field) => (
                <Space key={field.key} align="baseline">
                  <Form.Item
                    {...field}
                    name={[field.name, 'name']}
                    fieldKey={[field.fieldKey, 'name']}
                    rules={[{ required: true, message: 'Missing container name' }]}
                  >
                    <Input placeholder="Container Name" />
                  </Form.Item>
                  <Form.Item
                    {...field}
                    name={[field.name, 'image']}
                    fieldKey={[field.fieldKey, 'image']}
                    rules={[{ required: true, message: 'Missing image' }]}
                  >
                    <Input placeholder="Image" />
                  </Form.Item>
                  <Form.Item
                    {...field}
                    name={[field.name, 'imagePullPolicy']}
                    fieldKey={[field.fieldKey, 'imagePullPolicy']}
                  >
                    <Select placeholder="Image Pull Policy">
                      <Option value="Always">Always</Option>
                      <Option value="IfNotPresent">IfNotPresent</Option>
                      <Option value="Never">Never</Option>
                    </Select>
                  </Form.Item>
                  <Form.Item
                    {...field}
                    name={[field.name, 'command']}
                    fieldKey={[field.fieldKey, 'command']}
                  >
                    <Input placeholder="Command" />
                  </Form.Item>
                  <Form.Item
                    {...field}
                    name={[field.name, 'args']}
                    fieldKey={[field.fieldKey, 'args']}
                  >
                    <Input placeholder="Arguments" />
                  </Form.Item>
                  <Button onClick={() => remove(field.name)}>Remove</Button>
                </Space>
              ))}
              <Button type="dashed" onClick={() => add()} block>
                Add Init Container
              </Button>
            </>
          )}
        </Form.List>
      </Form.Item>
      <Form.Item name="containers" label="Containers">
        <Form.List name="containers">
          {(fields, { add, remove }) => (
            <>
              {fields.map((field) => (
                <Space key={field.key} align="baseline">
                  <Form.Item
                    {...field}
                    name={[field.name, 'name']}
                    fieldKey={[field.fieldKey, 'name']}
                    rules={[{ required: true, message: 'Missing container name' }]}
                  >
                    <Input placeholder="Container Name" />
                  </Form.Item>
                  <Form.Item
                    {...field}
                    name={[field.name, 'image']}
                    fieldKey={[field.fieldKey, 'image']}
                    rules={[{ required: true, message: 'Missing image' }]}
                  >
                    <Input placeholder="Image" />
                  </Form.Item>
                  <Form.Item
                    {...field}
                    name={[field.name, 'imagePullPolicy']}
                    fieldKey={[field.fieldKey, 'imagePullPolicy']}
                  >
                    <Select placeholder="Image Pull Policy">
                      <Option value="Always">Always</Option>
                      <Option value="IfNotPresent">IfNotPresent</Option>
                      <Option value="Never">Never</Option>
                    </Select>
                  </Form.Item>
                  <Form.Item
                    {...field}
                    name={[field.name, 'command']}
                    fieldKey={[field.fieldKey, 'command']}
                  >
                    <Input placeholder="Command" />
                  </Form.Item>
                  <Form.Item
                    {...field}
                    name={[field.name, 'args']}
                    fieldKey={[field.fieldKey, 'args']}
                  >
                    <Input placeholder="Arguments" />
                  </Form.Item>
                  <Button onClick={() => remove(field.name)}>Remove</Button>
                </Space>
              ))}
              <Button type="dashed" onClick={() => add()} block>
                Add Container
              </Button>
            </>
          )}
        </Form.List>
      </Form.Item>
      <Form.Item>
        <Button type="primary" htmlType="submit">
          Create Pod
        </Button>
      </Form.Item>
    </Form>
  );
};

export default CreatePod;
