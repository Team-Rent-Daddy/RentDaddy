import {
  Avatar,
  Button,
  Card,
  Col,
  ConfigProvider,
  Divider,
  Input,
  Row,
} from "antd"
import AntDesignTableComponent from "../components/AntDesignTableComponent"
import { SettingOutlined, UserOutlined } from "@ant-design/icons"
import TextArea from "antd/es/input/TextArea"
import { useState } from "react"
import UniversalSidebar from "../components/UniversalSidebar"
import { Link } from "react-router"
import RegistrationFormExample from "../components/FormExample"
import TimeRelatedFormExample from "../components/TimeRelatedFormExamples"
import AlertComponent from "../components/AlertComponent"

const ReusableComponents = () => {
  const [value, setValue] = useState("")

  return (
    <>
      {/* Had to comment this out to merge the layout, we need to use this higher up in the parent child component proccesses (maybe main.tsx?) */}
      {/* JJ Revisit this later when ready */}
      {/* <ConfigProvider
                theme={{
                    cssVar: true,
                    hashed: false,
                    token: {
                        // can put styles in here so they are applied to our components like table
                        colorPrimary: "#00674f",
                        colorBgContainer: "hsl(166, 100%, 20%, 5%)",
                        colorFillSecondary: "#7789f4",
                        colorFillTertiary: "#d86364",
                    },
                }}
            > */}
      <div className="reusable-components">
        <div className="title text-center fs-1">ReusableComponents</div>
        <h4 className="text-center mt-3 fst-italic fw-bold">
          ***There is not much styling on this yet, kinda stock, <br /> just so
          we can get the components up and running.***
        </h4>

        {/* Back to Home Button */}
        <Link to="/" className="d-flex justify-content-center">
          <Button type="primary" className="mt-3 btn btn-primary">
            Back to Home
          </Button>
        </Link>

        <Divider />

        {/* User Avatar */}
        <div className="user-avatar m-5 flex flex-column align-items-center">
          <h2 className="fs-2">User Avatar</h2>
          <Avatar size={64} icon={<UserOutlined />} />
        </div>

        <Divider />

        {/* alerts */}
        <AlertComponent
          title={"Success Example"}
          description={"Success Description"}
          type={"success"}
        />
        <AlertComponent
          title={"Error Example"}
          description={"Error Description"}
          type={"error"}
        />
        <AlertComponent
          title={"Warning Example"}
          description={"Warning Description"}
          type={"warning"}
        />

        {/* Sider (Sidebar) */}
        <div className="sider m-5">
          <h2 className="fs-2 text-center">Layout: Sider (Sidebar) & Footer</h2>
          <p className="text-center">
            The sidebar falls under the Layout Component of Ant Design.
          </p>
          <h1 className="text-center">
            MAKE SURE WE TALK ABOUT THIS IN THE NEXT MEETING OR IF SOMEONE KNOWS
            WHAT THE BEST WAY OF SETTING UP THE LAYOUT / SIDEBAR FEEL FREE TO
            CHANGE IT
          </h1>
          <UniversalSidebar />
        </div>

        <Divider />

        {/* Tables */}
        <div className="table m-5">
          <h2>Table</h2>
          <AntDesignTableComponent />
        </div>

        <Divider />

        {/* Cards */}
        <div className="cards m-5">
          <h2 className="fs-2 text-center">Cards</h2>
          <Row gutter={16}>
            <Col span={8}>
              <Card
                title="Card Primary"
                variant="borderless"
                className="card card-primary"
              >
                Lorem ipsum dolor sit, amet consectetur adipisicing elit. Esse
                quam blanditiis, ratione cumque, repudiandae mollitia ex tempora
                natus rem sint sapiente? Enim recusandae, similique voluptatibus
                facilis voluptatem non reprehenderit harum, unde fuga quos omnis
                molestias voluptate ducimus! Cupiditate eligendi distinctio
                eaque nulla soluta ab commodi nihil. Itaque cumque voluptatem a.
              </Card>
            </Col>
            <Col span={8}>
              <Card
                title="Card title"
                variant="borderless"
                className="card card-secondary"
              >
                Lorem ipsum dolor sit amet consectetur adipisicing elit. Ipsum
                eligendi dolores, quam necessitatibus voluptatum rem vero
                cupiditate laborum excepturi error totam corporis vitae, eaque,
                cumque doloremque id reiciendis consectetur magni aspernatur
                labore corrupti. Nobis omnis est eligendi error voluptatum a
                possimus minima dolorem quisquam, odio commodi ea adipisci alias
                reiciendis.
              </Card>
            </Col>
            <Col span={8}>
              <Card
                title="Card title"
                variant="borderless"
                className="card card-accent"
              >
                Lorem ipsum dolor sit, amet consectetur adipisicing elit. Esse
                quam blanditiis, ratione cumque, repudiandae mollitia ex tempora
                natus rem sint sapiente? Enim recusandae, similique voluptatibus
                facilis voluptatem non reprehenderit harum, unde fuga quos omnis
                molestias voluptate ducimus! Cupiditate eligendi distinctio
                eaque nulla soluta ab commodi nihil. Itaque cumque voluptatem a.
              </Card>
            </Col>
          </Row>
        </div>

        <Divider />

        {/* Buttons */}
        <div className="buttons-examples m-5">
          <h2 className="fs-2 text-center">Buttons</h2>
          <div className="buttons-container flex flex-column align-items-center gap-2">
            {/* Primary Button */}
            <div className="my-2">
              <h2>Primary Button</h2>
              <Button type="primary" className="flex mx-auto btn btn-primary">
                Click me
              </Button>
            </div>

            {/* Confirm Button */}
            <div className="my-2">
              <h2>Confirm Button</h2>
              <Button className="bg-success flex mx-auto"> Click me</Button>
            </div>

            {/* Decline Button */}
            <div className="my-2">
              <h2>Decline Button</h2>
              <Button className="bg-danger flex mx-auto">Click me</Button>
            </div>

            {/* Loading Button */}
            <div className="my-2">
              <h2>Loading Button</h2>
              <Button className="bg-primary flex mx-auto" loading>
                Click me
              </Button>
            </div>

            {/* Disabled Button */}
            <div className="my-2">
              <h2>Disabled Button</h2>
              <Button className="bg-secondary flex mx-auto" disabled>
                Click me
              </Button>
            </div>

            {/* Any other buttons? */}
          </div>
        </div>

        <Divider />

        {/* Text Area Examples*/}
        <div className="text-area-examples m-5">
          <h2 className="fs-2 text-center">Text Area Examples</h2>

          {/* Basic Input */}
          <div className="my-2">
            <h2>Basic Input</h2>
            <Input className="my-2" placeholder="Basic usage" />
          </div>

          {/* Search Input */}
          <div className="my-2">
            <h2>Search Input</h2>
            <Input
              className="my-2"
              addonAfter={<SettingOutlined />}
              defaultValue="mysite"
            />
            <Input
              className="my-2"
              addonBefore="http://"
              suffix=".com"
              defaultValue="mysite"
            />
          </div>

          {/* Controlled Expanding Text Area */}
          <div className="my-2">
            <h2>Controlled Expanding Text Area</h2>
            <TextArea
              className="my-2"
              value={value}
              onChange={(e) => setValue(e.target.value)}
              placeholder="Controlled autosize"
              autoSize={{ minRows: 3, maxRows: 5 }}
            />
          </div>

          <Divider />

          <div className="form-examples m-5 p-4 border rounded shadow-sm bg-light">
            <h2 className="fs-2 text-center mb-4 text-primary">
              Form Examples
            </h2>
            <div className="bg-white p-3 rounded">
              <h2 className="fs-3 text-center mb-4 text-primary">
                Registration Form
              </h2>
              <RegistrationFormExample />
            </div>
            <Divider />
            <div className="bg-white p-3 rounded">
              <h2 className="fs-3 text-center mb-4 text-primary">
                Time Related Form Examples
              </h2>
              <TimeRelatedFormExample />
            </div>
          </div>

          {/* Any other text area examples? */}
        </div>

        <Divider />
      </div>
      {/* </ConfigProvider> */}
    </>
  )
}

export default ReusableComponents
