import { Button, Form, Col, Comment, Descriptions, Divider, Modal, Rate, Row, Statistic, List, Tooltip, Typography, Drawer, Select, Input, DatePicker, Tabs, InputNumber, Upload, message, Alert } from "antd";
import { CommentOutlined, EyeInvisibleOutlined, EyeTwoTone, HeartOutlined, InboxOutlined, LoadingOutlined, LockOutlined, PlusOutlined, UserOutlined } from '@ant-design/icons';
import React, { useEffect, useState } from "react";
import api from './api';
import { SERVER_URL } from "./config";
import Avatar from "antd/lib/avatar/avatar";
import axios from "axios";
import ImageUploader from 'react-images-upload';
import ImageUploading from 'react-images-uploading';
import ImgCrop from 'antd-img-crop';
import headers from "./headers";
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';


function UserDetailsForm({ handleCloseDrawer, showDrawer, onSubmit, userDetails }) {

  const [imageUrl, setImageUrl] = useState("");
  const [file, setFile] = useState(null);
  const [loadingFile, setLoadingFile] = useState(false);
  const [stateInfos, setStateInfos] = useState({});
  const [isRecoverPass, setIsRecoverPass] = useState(false);
  const [image, setImage] = useState("");
  const [loadingForm, setLoadingForm] = useState(false);
  const [form] = Form.useForm();
  const [showUsuarioInvalido, setShowUsuarioInvalido] = useState(false)
  const [tab, setTab] = useState('1')

  useEffect(() => {
    form.resetFields()
  }, [showDrawer])

  useEffect(() => {
    if(!!userDetails && !!userDetails.photo) {
      setImageUrl(userDetails.photo)
      var blob = null;
      var xhr = new XMLHttpRequest();
      xhr.open("GET", userDetails.photo);
      xhr.responseType = "blob";//force the HTTP response, response-type header to be blob
      xhr.onload = function()
      {
          blob = xhr.response;//xhr.response is now a blob object
      }
      xhr.send();


      setFile(new File([blob], userDetails.photo.substring(userDetails.photo.lastIndexOf("/"))))
    }
  }, [userDetails])

  const { TabPane } = Tabs;
  const handleChangeFile = info => {
    if (info.file.status === 'uploading') {
      setLoadingFile(true);
      return;
    }
    if (info.file.status === 'done') {
      getBase64(info.file.originFileObj, imageUrl => {
        setImageUrl(imageUrl);
        setLoadingFile(false)
      },
      );
    }
  };

  const normFile = e => {
    if (Array.isArray(e)) {
      return e;
    }
    return e && e.fileList;
  };

  function getBase64(img, callback) {
    const reader = new FileReader();
    reader.addEventListener('load', () => callback(reader.result));
    reader.readAsDataURL(img);
  }

  function beforeUpload(file) {
    
    const isJpgOrPng = file.type === 'image/jpeg' || file.type === 'image/png';
    if (!isJpgOrPng) {
      message.error('You can only upload JPG/PNG file!');
    }
    const isLt2M = file.size / 1024 / 1024 < 2;
    if (!isLt2M) {
      message.error('Image must smaller than 2MB!');
    }

    setImageUrl(URL.createObjectURL(file));
    setFile(file);
    return false;
  }

  function updateReader() {
    setLoadingForm(true)
    const formData = new FormData();


    form.validateFields().then((values) => {
      const { username, ...myUserDetails } = values

      Object.entries(myUserDetails).map(([k, v], i) => {
        formData.append(k, v)
      })

      formData.append("name", username)
      formData.append("profile", file)

      axios.patch(`${SERVER_URL}/api/v1/readers/${userDetails.id}`, formData, {
        headers: {
          ...headers(),
          "Content-Type": "multipart/form-data",
        }
      }).then(r => {
        toast.success("Dados atualizados com sucesso!")
        const newreader =  { email: myUserDetails.email, age: myUserDetails.age, username: username, id: r.data.ID, photo: r.data.Photo}
        setLoadingForm(false)
        form.resetFields()
        handleCloseDrawer()
        onSubmit(newreader)
      })
    })
  }

  const onPreview = async file => {
    let src = file.url;
    if (!src) {
      src = await new Promise(resolve => {
        const reader = new FileReader();
        reader.readAsDataURL(file.originFileObj);
        reader.onload = () => resolve(reader.result);
      });
    }
    const image = new Image();
    image.src = src;
    const imgWindow = window.open(src);
    imgWindow.document.write(image.outerHTML);
  };


  return <Drawer
    title="Login"
    width={720}
    onClose={() => { handleCloseDrawer(); }}
    visible={showDrawer}
    bodyStyle={{ paddingBottom: 80 }}
    footer={<div
      style={{
        textAlign: 'right',
      }}
    >
      <Button onClick={() => { handleCloseDrawer(); }} style={{ marginRight: 8 }}>
        Cancelar
      </Button>

    </div>}
  >
    <Form initialValues={userDetails} form={form} layout={"vertical"} style={{ width: 500, marginTop: 30, marginLeft: 80 }}>
      <Tabs activeKey={tab} onChange={(k) => {
          setTab(k)
          form.resetFields()
        }}>
        {
          <>
            <TabPane tab={"Atualize seus dados"} key={'1'}>
            <Row gutter={24}>
                <Col md={24}>
                  <Form.Item name={"email"} label={"E-mail"} rules={[{ required: true, type: 'email', message: 'Por favor, informe um e-mail válido!' }]}>
                    <Input disabled />
                  </Form.Item>
                </Col>

              </Row>
              <Row gutter={24}>
                <Col md={18}>
                  <Form.Item name="username" label={"Nome de usuário"}
                    rules={[{ required: true, message: 'Por favor, informe o nome de usuário!' }]}>
                    <Input />
                  </Form.Item>
                </Col>
                <Col md={6}>
                  <Form.Item name="age" label={"Idade"} rules={[{ type: 'number', min: 0, max: 99, message: "Por favor, informe um número entre 0 e 99!" }]}>
                    <InputNumber />
                  </Form.Item>
                </Col>
              </Row>
              <Row gutter={24}>
                <Col md={24}>
                  <Form.Item
                    name="password"
                    label={"Senha"}
                  >
                    <Input.Password
                      iconRender={visible => (visible ? <EyeTwoTone /> : <EyeInvisibleOutlined />)} />
                  </Form.Item>
                </Col>
              </Row>
              <Form.Item label="Foto">
                <Form.Item name="photo" valuePropName="fileList" getValueFromEvent={normFile} noStyle>
                <ImgCrop rotate>
                  <Upload.Dragger onPreview={onPreview} style={imageUrl ? { width: 100, height: 100 } : {}} action="https://www.mocky.io/v2/5cc8019d300000980a055e76" beforeUpload={beforeUpload} name="files" onChange={handleChangeFile}>
                    {imageUrl ? <Avatar size={64} src={imageUrl} /> :
                      <>
                        {loadingFile ? <LoadingOutlined /> :
                          <>
                            <p className="ant-upload-drag-icon">
                              <InboxOutlined />
                            </p>
                            <p className="ant-upload-text">Clique ou arraste o arquivo para carregar.</p>
                          </>}
                      </>}
                  </Upload.Dragger>
                  </ImgCrop>
                </Form.Item>
              </Form.Item>
              <Row>
                <Col md={24}>
                  <Form.Item>
                    <Button loading={loadingForm} onClick={updateReader} type="primary" style={{ width: "100%", float: "right" }}>Atualizar</Button>

                  </Form.Item>
                </Col>
              </Row>

            </TabPane>
          </>}
      </Tabs>
    </Form>
  </Drawer>;
}

export default UserDetailsForm;