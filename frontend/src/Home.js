import { ArrowUpOutlined, LogoutOutlined, UserOutlined, VerticalAlignTopOutlined } from "@ant-design/icons";
import { BackTop, Dropdown, Image, Layout, Menu } from "antd";
import Avatar from "antd/lib/avatar/avatar";
import React, { useEffect, useState } from 'react';
import { Link, Router } from "react-router-dom";
import api from "./api";
import BookList from "./BookList";
import ModalBookDetails from "./ModalBookDetails";
import MyBookList from "./MyBookList";
import UserDetailsForm from "./UserDetailsForm";
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

function Home() {
    console.log(api.userPhoto())
    const { Header, Footer, Content } = Layout;
    const [book, setBook] = useState({});
    const [showModal, setShowModal] = useState(false);
    const [showUser, setShowUser] = useState(false);
    const [showUserDetails, setShowUserDetails] = useState(false);
    const location = document.location.href.substring(document.location.href.lastIndexOf('/'))

    const onBookClick = (book) => {
        setBook(book);
        setShowModal(true);
    }

    const onOkModal = () => {
        setShowModal(false)
        setBook({});
    }


    useEffect(() => {
        setShowUser(api.loggedIn())
    }, [])

    
    function reloadBook() {
        api.get(`/books/${book.id}`).then(r => {
            const e = r.data;
            setBook({ id: e.ID, author: e.Author, title: e.Title, cover: e.Cover, genre: e.Genre, summary: e.Summary, averageRate: e.Average, comments: e.Qtd })
        })
    }

    const subMenu = (
        <Menu>
          <Menu.Item key="1" onClick={() => {setShowUserDetails(true)}} icon={<UserOutlined />}>
            Meus Dados
          </Menu.Item>
          <Menu.Item onClick={() => {
              toast.success("Logout realizado!")
              setShowUser(false); 
              api.logout()}} key="2" icon={<LogoutOutlined />}>
            Sair
          </Menu.Item>
        </Menu>
      );

    return (

        <Layout>
            <Header style={{ position: 'fixed', zIndex: 1, width: '100%' }}>
                <div className="logo" style={{
                    width: 200,
                    height: 40,
                    margin: '0px 24px 16px 0',
                    float: 'left'
                }}>
                    <Link to="/">
                        <img height={40} src="https://dad-react-static.s3.us-east-2.amazonaws.com/untitled.png" /></Link></div>
                <Menu selectedKeys={location === '/interests' ? ['2'] : ['1']} theme="dark" mode="horizontal">
                    <Menu.Item key="1"><Link to="/">Todos os Livros</Link></Menu.Item>
                    {showUser ? <Menu.Item key="2"><Link to="/interests">Meus Livros</Link></Menu.Item> : null}
                    {
                        showUser ?

                            <div style={{ float: "right", marginRight: 40 }} key="4">
                                <Dropdown overlay={subMenu} trigger={['click']} placement="bottomLeft">
                                    <a onClick={(e) => {e.preventDefault()}}>
                                {
                                     
                                    !!api.userPhoto() ?
                                       <Avatar size={48} src={api.userPhoto()} /> :
                                        <Avatar size={48} style={{ background: '#f0f2f5', color: 'teal', fontWeight: "bold" }}>{api.getUsername()[0]}</Avatar>
                                }
                                </a>
                                </Dropdown>
                            </div> : null
                    }

                </Menu>
            </Header>
            <Layout>
                <Content style={{ padding: '0 50px', marginTop: 64 }}>
                    <div style={{
                        background: '#f0f2f5',
                        padding: '24px',
                        minHeight: '600px'
                    }}>
                        {
                            location === '/interests' && api.loggedIn() ?
                                <MyBookList onBookClick={onBookClick} />
                                : <BookList onBookClick={onBookClick} />
                        }

                    </div>
                    <BackTop>
                        <div style={{
                            height: 50,
                            width: 50,
                            lineHeight: '40px',
                            borderRadius: 4,
                            backgroundColor: 'teal',
                            color: '#fff',
                            textAlign: 'center',
                            verticalAlign: 'middle',
                            fontSize: 20
                        }}><ArrowUpOutlined /></div>
                    </BackTop>
                </Content>

            </Layout>
            <Footer style={{ textAlign: 'center' }}>Footer</Footer>
            <ModalBookDetails setShowUser={setShowUser} showModal={showModal} book={book} reloadBook={reloadBook} onOk={onOkModal} />
            <UserDetailsForm onSubmit={(u) => {api.setUser(u)}} handleCloseDrawer={() => setShowUserDetails(false)} showDrawer={showUserDetails} userDetails={api.userJSON()} />
            <ToastContainer />
        </Layout>
    );
}

export default Home;