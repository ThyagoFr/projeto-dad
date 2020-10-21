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
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';


function LoginForm({ handleCloseDrawer, showDrawer, onSubmit }) {

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

  const login = () => {
    setShowUsuarioInvalido(false)
    setLoadingForm(true)
    form.validateFields(['email', 'password']).then((values) => {
      const userDetails = values

      api.post(`/login`, { email: userDetails.email, password: userDetails.password }).then(response => {
        toast.success("Login efetuado com successo!")
        console.log(response)
        const user = response.data.reader;
        api.login({ access_token: response.data.token, 
          user: { email: user.Email, age: user.Age, username: user.Name, id: user.ID, photo: user.Photo }})
        setLoadingForm(false)
        handleCloseDrawer()
        onSubmit()

      }).catch(error => {
        if (!!error && !!error.response && error.response.status == 401) {
          setLoadingForm(false)
          setShowUsuarioInvalido(true)
        }
      });

    })



  };

  const recuperarSenha = () => {
    setLoadingForm(true)
    form.validateFields().then((values) => {
      console.log(values)
      const userDetails = values
      api.post(`/sendemail`, { email: userDetails.email}).then(response => {
        console.log(response)
        setLoadingForm(false)
      }).catch((error) => {
        setLoadingForm(false)
      })
    })
  }

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

  function saveReader() {
    const formData = new FormData();
    setLoadingForm(true)

    form.validateFields().then((values) => {
      const { username, ...userDetails } = values

      Object.entries(userDetails).map(([k, v], i) => {
        formData.append(k, v)
      })

      formData.append("name", username)
      formData.append("profile", file)
      console.log(file)

      api.post(`/register`, formData, {
        headers: {
          "Content-Type": "multipart/form-data"
        }
      }).then(r => {
        toast.success("Registro Criado! Por favor, faça login para continuar.")
        setLoadingForm(false)
        setTab("1")
        setShowUsuarioInvalido(false)
        form.resetFields()
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
    <Form form={form} layout={"vertical"} style={{ width: 500, marginTop: 30, marginLeft: 80 }}>
      <Tabs activeKey={tab} onChange={(k) => {
          setTab(k)
          form.resetFields()
        }}>
        {isRecoverPass ? <TabPane tab={"Recuperar Senha"} key={"3"}>
          <Row gutter={24}>
            <Col md={24}>
              <Form.Item name={"email"} label={"E-mail"} rules={[{ required: true, type: 'email', message: 'Por favor, informe um e-mail válido!' }]}>
                <Input />
              </Form.Item>
            </Col>
            <Col md={12}>
              <Form.Item>
                <Button type="default" onClick={() => { 
                  form.resetFields()
                  setIsRecoverPass(false) 
                  setTab("1")
                  }} style={{ width: "100%" }}>Voltar</Button>
              </Form.Item>
            </Col>
            <Col md={12}>
              <Form.Item>
                <Button type="primary" loading={loadingForm} onClick={() => { recuperarSenha() }} style={{ width: "100%" }}>Recuperar</Button>

              </Form.Item>
            </Col>


          </Row>
        </TabPane> :
          <>
            <TabPane tab={"Entrar"} key={1}>
              <Form.Item
                name="email"
                rules={[{ required: true, message: 'Por favor, informe o e-mail!' }]}
              >
                <Input prefix={<UserOutlined className="site-form-item-icon" />} placeholder="E-mail" />
              </Form.Item>
              <Form.Item
                name="password"

                rules={[{ required: true, message: 'Por favor, informe a senha!' }]}
              >
                <Input.Password
                  prefix={<LockOutlined className="site-form-item-icon" />}
                  placeholder="Senha"
                  iconRender={visible => (visible ? <EyeTwoTone /> : <EyeInvisibleOutlined />)} />
              </Form.Item>
              {showUsuarioInvalido ? <Alert message="E-mail não cadastrado ou senha inválida!" type="error" /> : null}
              <Form.Item>

                <a onClick={(e) => { e.preventDefault(); setIsRecoverPass(true); setTab("3") }} style={{ float: "right" }} href="">
                  Esqueci minha senha
            </a>
              </Form.Item>

              <Form.Item>
                <Button loading={loadingForm} onClick={login} type="primary" style={{ width: "100%" }}>Entrar</Button>

              </Form.Item>
            </TabPane>
            <TabPane tab={"Registre-se"} key={2}>
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
                  <Form.Item name={"email"} label={"E-mail"} rules={[{ required: true, type: 'email', message: 'Por favor, informe um e-mail válido!' }]}>
                    <Input />
                  </Form.Item>
                </Col>

              </Row>
              <Row gutter={24}>
                <Col md={24}>
                  <Form.Item
                    name="password"
                    label={"Senha"}
                    rules={[{ required: true, message: 'Por favor, informe a senha!' }]}
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
                    <Button loading={loadingForm} onClick={saveReader} type="primary" style={{ width: "100%", float: "right" }}>Salvar</Button>

                  </Form.Item>
                </Col>
              </Row>
              <ToastContainer />
            </TabPane>
          </>}
      </Tabs>
    </Form>
  </Drawer>;
}

export default LoginForm;