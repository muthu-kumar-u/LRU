import { Table, Button, Flex, Col, Row } from 'antd';
import axios from 'axios';
import React, { useState } from "react";

const ListView = (props) => {
  const [get, setGet] = useState(true);

  const handlePost = async (e) => {
    e.preventDefault();

    const formData = new FormData(e.target);
    const data = {};

    formData.forEach((value, key) => {
      data[key] = key === 'expire' ? parseInt(value, 10) : value;
    });

    try {
      const response = await axios.put(`/PostKey`, data);
      if (response.status === 200) {
        props.trigger();
      }
    } catch (error) {
      console.error('Error fetching cache data:', error.message);
    }
  };

  const handleGet = async (e) => {
    e.preventDefault();

    const formData = new FormData(e.target);

    let key;
    for (const [inputKey, value] of formData.entries()) {
      if (inputKey === 'key') {
        key = parseInt(value, 10);
        break; 
      }
    }
    
    try {
      const response = await axios.get(`/GetValue?key=${key}`);
      if (response.status === 200) {
        confirm(`${response.data.Value ? `The value for the key is "${response.data.Value}"` : response.data.message}`);
      }
    } catch (error) {
      console.error('Error fetching cache data:', error.message);
    }
  };

  return (
    <Row>
      <Col lg={18} sm={24} className="m-auto p-5">
        <div className="d-flex flex-column" style={{ gap: "70px" }}>
          <h1>Cache List</h1>
          <Flex gap="small" wrap>
            <Button type={get ? "primary" : ""} onClick={() => setGet(true)}>Get</Button>
            <Button type={!get ? "primary" : ""} onClick={() => setGet(false)}>Post</Button>
          </Flex>
          {!get ? (
            <form className="d-flex flex-column" encType='multipart/json' style={{ gap: "30px" }} onSubmit={handlePost}>
              <div className="form-group">
                <input
                  type="text"
                  className="form-control"
                  name='key'
                  placeholder="Enter the key.."
                  required
                />
              </div>
              <div className="form-group">
                <input
                  type="text"
                  className="form-control"
                  name='value'
                  placeholder="Enter the value.."
                  required
                />
              </div>
              <div className="form-group">
                <input
                  type="number"
                  className="form-control"
                  name='expire'
                  placeholder="Enter the expire time (in seconds).."
                  required
                />
              </div>
              <button type="submit" className="btn btn-primary">
                Submit
              </button>
            </form>
          ) : (
            <form className="d-flex flex-column" style={{ gap: "30px" }} encType='multipart/json' onSubmit={handleGet}>
              <div className="form-group">
                <input
                  type="text"
                  className="form-control mb-4"
                  name='key'
                  placeholder="Enter the key.."
                />
                <small>if the key exists in the cache then it will return the value, or else it will show the message</small>
              </div>
              <button type="submit" className="btn btn-primary">
                Submit
              </button>
            </form>
          )}
          <div style={{ maxHeight: "300px", overflowY: "scroll" }}>
            <table className="table table-dark">
              <thead className="thead-light">
                <tr>
                  <th>Key</th>
                  <th>Value</th>
                </tr>
              </thead>
              <tbody>
                {props.cache.length > 0 ? 
                  props.cache.map((item, index) => (
                    <tr key={index}>
                      <td>{Object.keys(item)[0]}</td>
                      <td>{Object.values(item)[0]}</td>
                    </tr>
                  )) : 
                  <tr>
                    <td>No Data</td>
                    <td>No Data</td>
                  </tr>
                }
              </tbody>
            </table>
            <Flex>
              <Button onClick={() => props.trigger()}>Reload</Button>
            </Flex>
          </div>
        </div>
      </Col>
    </Row>
  );
};

export default ListView;
