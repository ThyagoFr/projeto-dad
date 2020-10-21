import { Image, Row, Col, Card, Rate } from 'antd';
import axios from "axios";
import React from 'react';
import api from './api';
import { SERVER_URL } from './config';
import headers from './headers';

class BookList extends React.Component {
    state = {
        entities: [],
        totalCount: 0,
    }

    componentDidMount() {
        api.get("/books").then(r => {
            this.setState({
                entities: (r.data || []).map(e => ({
                    id: e.ID, author: e.Author, title: e.Title, cover: e.Cover, genre: e.Genre, summary: e.Summary, averageRate: e.Average, comments: e.Qtd
                }))
            })
        })

        if (api.loggedIn()) {

            axios.get(`${SERVER_URL}/api/v1/readers/${api.getBookId()}/interests`, { headers: headers() }).then(r => {
                if (!!r && !!r.data) {
                    api.setUserInterests({ interests: r.data.map(b => b.ID) })
                }
            })
        }
    }

    render() {
        const { Meta } = Card;
        return (
            <Row gutter={24}>
                {
                    this.state.entities.map((e, i) => {
                        return (
                            <Col key={e.id} span={4} style={{ padding: 20 }}>
                                <Card hoverable
                                actions={[<Rate style={{ marginLeft: 15 }} value={e.averageRate} allowHalf disabled />]}
                                    bordered={false}
                                    onClick={() => {
                                        this.props.onBookClick(e)
                                    }}
                                    bodyStyle={{ height: 120 }}
                                    cover={<img style={{ height: 200 }} alt={e.title} src={e.cover} />}>
                                    <Meta title={e.title} description={e.author}  />
                                </Card>
                            </Col>
                        )
                    })
                }
            </Row>
        )
    }
}

export default BookList;