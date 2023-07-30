import React from "react";
import { Container, Dropdown, Table } from "react-bootstrap";
import { useQuery } from "react-query";
import { API } from "../../config/Api";

function AdminListTransaction() {
  let { data: transaction } = useQuery("transactionCache", async () => {
    const response = await API.get("/transactions");
    return response.data.data.transaction;
  });

  console.log(transaction);

  return (
    <div>
      <Container fluid style={{ background: "#161616", height: "100vh" }}>
        <Container style={{ background: "#161616", height: "100vh" }}>
          <div style={{ height: "100px" }}></div>
          <div className="text-white">
            <h1>Incoming Transaciton</h1>
            <Table>
              <thead>
                <tr>
                  <th>No</th>
                  <th>Users</th>
                  <th>Bukti Transfer</th>
                  <th>Remaining Active</th>
                  <th>Status User</th>
                  <th>Status Payment</th>
                  <th>Action</th>
                </tr>
              </thead>
              <tbody>
                {transaction?.map((data, index) => (
                  <tr>
                    <td>{index + 1}</td>
                    <td>{data.user.fullname}</td>
                    <td>BCA</td>
                    <td>{data.due_date}</td>
                    <td>{!data.user.listAs ? "not active" : "active"}</td>
                    <td>{data.status}</td>
                    <td>
                      <Dropdown>
                        <Dropdown.Toggle variant="success" id="dropdown-basic">
                        </Dropdown.Toggle>

                        <Dropdown.Menu>
                          <Dropdown.Item href="#/action-1">
                            Action
                          </Dropdown.Item>
                          <Dropdown.Item href="#/action-2">
                            Another action
                          </Dropdown.Item>
                          <Dropdown.Item href="#/action-3">
                            Something else
                          </Dropdown.Item>
                        </Dropdown.Menu>
                      </Dropdown>
                    </td>
                  </tr>
                ))}
              </tbody>
            </Table>
          </div>
        </Container>
      </Container>
    </div>
  );
}

export default AdminListTransaction;
