import { Button, Form, Col, Comment, Descriptions, Divider, Modal, Rate, Row, Statistic, List, Tooltip, Typography, Drawer, Select, Input, DatePicker, Tabs, InputNumber, Upload, message } from "antd";
import { IconMap } from "antd/lib/result";
import axios from "axios";
import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { CheckOutlined, CommentOutlined, EyeInvisibleOutlined, EyeTwoTone, HeartOutlined, InboxOutlined, LoadingOutlined, LockOutlined, PlusOutlined, UserOutlined, VerifiedOutlined } from '@ant-design/icons';
import Avatar from "antd/lib/avatar/avatar";
import LoginForm from "./LoginForm";
import TextArea from "antd/lib/input/TextArea";
import api from "./api";
import headers from "./headers";
import { SERVER_URL } from "./config";
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';


function ModalBookDetails({ book, showModal, onOk, reloadBook, setShowUser }) {
  const [bookId, setBookId] = useState({});
  const [showComments, setShowComments] = useState(false);
  const [showDrawer, setShowDrawer] = useState(false);
  const { Paragraph } = Typography;
  const [userDetails, setUserDetails] = useState({});
  const [userComment, setUserComment] = useState("");
  const [userRate, setUserRate] = useState(5);
  const [submittingComment, setSubmittingComment] = useState(false);
  const [userAlreadyCommented, setUserAlreadyCommented] = useState(false)
  const [reviews, setReviews] = useState([]);
  const [isReaderInterest, setIsReaderInterest] = useState(false)
  const [form] = Form.useForm();

  function retrieveUserDetails() {
    setUserDetails(JSON.parse(localStorage.user))
  }

  useEffect(() => {
    setBookId(book.id)
    setShowComments(false)
    setShowDrawer(false)
    setUserComment("")
    setUserRate(5)
    setReviews([])
    setUserAlreadyCommented(false)
    let user = api.userJSON()
    form.resetFields()
    setIsReaderInterest(user && user.interests && user.interests.includes(book.id))
  }, [book])

  function getReviews() {
    if (api.loggedIn()) {
      api.get(`${SERVER_URL}/api/v1/books/${bookId}/comments`, { headers: headers() }).then(r => {
        const readers = r.data.map(u => ({
          reader: u.User, email: u.Email, comment: u.Comment, rate: u.Rate, picture: u.UserPhoto || "https://dad-react-static.s3.us-east-2.amazonaws.com/icon.png"
        }))

        retrieveUserDetails()
        setUserAlreadyCommented(readers.some(r => r.email === JSON.parse(localStorage.user).email))
        setReviews(readers)
        setShowComments(true)
      })
    } else {
      setShowDrawer(true)
    }
  }

  function onLogin() {
    setShowUser(true)
  }

  const Editor = ({ onChange, onSubmit, submitting, value }) => (
    <>
      <Form form={form}>
        <Form.Item name="userComment">
          <TextArea rows={4} />
        </Form.Item>
        <Row>
          <Form.Item>
            <Rate defaultValue={5} value={userRate} onChange={setUserRate} style={{ marginRight: 20 }} allowHalf />
          </Form.Item>
          <Form.Item>
            <Button htmlType="submit" loading={submitting} onClick={onSubmit} type="primary">
              Salvar Avaliação
          </Button>
          </Form.Item>
        </Row>
      </Form>
    </>
  );

  function verificaUsuarioLogado() {
    if (api.loggedIn()) {
      retrieveUserDetails()
      setShowComments(true)
    } else {
      setShowDrawer(true)
    }
  }

  function handleChangeComment(e) {
    setUserComment(e.target.value)
  }

  function handleSubmitComment() {
    setSubmittingComment(true);
    form.validateFields().then(values => {
      const comment = values.userComment;
      const rate = userRate;
      const readerId = api.getBookId();

      const body = { comment: comment, rate: rate, reader_id: readerId }

      axios.post(`${SERVER_URL}/api/v1/books/${bookId}/comments`, body, { headers: headers() }).then(r => {
        setSubmittingComment(false)
        setUserAlreadyCommented(true)
        setReviews([...reviews, { reader: userDetails.username, comment: comment, rate: rate, picture: userDetails.photo || "https://dad-react-static.s3.us-east-2.amazonaws.com/icon.png" }])
        reloadBook()
      })
    })
  }

  function removeBookFromInterests() {
    if (api.loggedIn() && book.interestId) {

      axios.delete(`${SERVER_URL}/api/v1/readers/${api.getBookId()}/interests/${book.interestId}`, { headers: headers() }).then(r => {
      })

      setIsReaderInterest(false)

      onOk()
    }
  }

  function addBookToInterests() {
    if (api.loggedIn()) {
      const body = { book_id: bookId }

      axios.post(`${SERVER_URL}/api/v1/readers/${api.getBookId()}/interests`, body, { headers: headers() }).then(r => {
      })

      setIsReaderInterest(true)

      axios.get(`${SERVER_URL}/api/v1/readers/${api.getBookId()}/interests`, { headers: headers() }).then(resp => {
        if (!!resp && !!resp.data) {
          api.setUserInterests({ interests: resp.data.map(b => b.ID) })
        }
      })
    } else {
      setShowDrawer(true)
    }
  }

  const UserCommentComponent = () => {
    if (!userAlreadyCommented) {
      return (
        <>
          <Comment
            avatar={
              <Avatar
                src={userDetails.photo || "https://dad-react-static.s3.us-east-2.amazonaws.com/icon.png"}
                alt={userDetails.username}
              />
            }
            content={
              <Editor
                onChange={handleChangeComment}
                onSubmit={handleSubmitComment}
                submitting={submittingComment}
                value={userComment}
              />
            }
          />
        </>
      );
    } else {
      return null
    }
  }

  return (
    <Modal width={900} onCancel={onOk} onOk={onOk} visible={showModal} >
      <Row gutter={24}>
        <img style={{ height: 300, marginRight: 20 }} width={"25%"} alt={book.title} src={book.cover} />
        <Descriptions title={"Info"} column={1} style={{ display: 'inline-block', width: "70%" }}>
          <Descriptions.Item label={"Título"}>{book.title}</Descriptions.Item>
          <Descriptions.Item label={"Autor"}>{book.author}</Descriptions.Item>
          <Descriptions.Item label={"Gênero"}>{book.genre}</Descriptions.Item>
          <Descriptions.Item label={"Resumo"}>
            <Paragraph ellipsis={{ rows: 5, expandable: true, symbol: 'mais' }}>
              {book.summary}
            </Paragraph>
          </Descriptions.Item>
        </Descriptions>
      </Row>
      <Row style={{ marginTop: 30 }} gutter={24}>
        <Rate style={{ marginLeft: 15 }} value={book.averageRate} allowHalf disabled />
        {
          book.comments > 0 ? <>
            <Statistic style={{ paddingLeft: 8 }} value={book.averageRate} suffix="/ 5" />
            <Col md={10}>

              <a onClick={getReviews}>
                <Statistic prefix={<CommentOutlined />} style={{ paddingLeft: 8 }} value={book.comments} suffix="avaliações" />
              </a>
            </Col></> : <Col md={11}>
              Este livro ainda não foi avaliado.
                  <Button onClick={verificaUsuarioLogado} style={{ marginTop: 4 }} type="text">
                <span style={{ fontWeight: "bold", color: "teal" }}>Deixe sua avaliação.</span>
              </Button>
            </Col>
        }

        <Col md={8}>
          {
            isReaderInterest ?
              <Tooltip title="Tenho interesse">
                <Button type="primary" style={{ float: "right", marginLeft: 15 }} shape="circle" size={"large"} icon={<CheckOutlined />} onClick={removeBookFromInterests} />
              </Tooltip>
              :
              <Tooltip title="Tenho interesse">
                <Button style={{ float: "right", marginLeft: 15 }} shape="circle" size={"large"} icon={<PlusOutlined />} onClick={addBookToInterests} />
              </Tooltip>
          }

        </Col>
      </Row>

      {
        showComments ? <>
          <Divider orientation="left">Comentários</Divider>
          <List split={false}
            className="comment-list"
            itemLayout="horizontal"
            dataSource={reviews}
            style={{ width: "100%" }}
            size="small"
            itemLayout="horizontal"
            renderItem={item => (
              <div>
                <List.Item extra={<Statistic style={{ paddingTop: 8 }} value={item.rate} suffix="/ 5" />}>
                  <Comment
                    author={item.reader}
                    avatar={<Avatar
                      src={item.picture}
                      alt={item.reader}
                    />}
                    content={item.comment}
                  />

                </List.Item>
              </div>
            )}
          />
          <UserCommentComponent />
        </>
          : null
      }

      <LoginForm onSubmit={onLogin} handleCloseDrawer={() => setShowDrawer(false)} showDrawer={showDrawer} />
    </Modal>
  );
}

export default ModalBookDetails;