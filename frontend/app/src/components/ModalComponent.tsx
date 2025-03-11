// TODO: Once we have the tenant info from the backend, make sure to populate the fields in the edit tenant modal so that the user can edit the tenant info easily
import { useState } from 'react';
import { Button, Divider, Form, Input, Modal } from 'antd';
import { PlusOutlined } from '@ant-design/icons';
import ButtonComponent from './reusableComponents/ButtonComponent';
import FormItemLabel from 'antd/es/form/FormItemLabel';

interface ModalComponentProps {
    buttonTitle: string;
    buttonType: "default" | "primary" | "secondary" | "accent" | "info" | "success" | "warning" | "danger";
    content: string;
    type: "default" | "Smart Locker" | "Guest Parking" | "Add Tenant" | "Edit Tenant";
    handleOkay: () => void;
    modalTitle?: string
}

const ModalComponent = (props: ModalComponentProps) => {
    const [isModalOpen, setIsModalOpen] = useState(false);

    const showModal = () => {
        setIsModalOpen(true);
    };

    const handleCancel = () => {
        setIsModalOpen(false);
    };

    const titles = {
        "default": "Default Modal",
        "Smart Locker": "Smart Locker Modal",
        "Guest Parking": "Register someone in Guest Parking",
        "Add Tenant": "Add Tenant"
    }

    return (
        <>
            {props.type === "default" && (
                <>
                    <ButtonComponent title={props.buttonTitle} type={props.buttonType} onClick={showModal} />
                    <Modal
                        title={<h3>{props.modalTitle}</h3>}
                        open={isModalOpen}
                        onOk={props.handleOkay}
                        onCancel={handleCancel}
                        okButtonProps={{ hidden: true, disabled: true }}
                        cancelButtonProps={{ hidden: true, disabled: true }}
                    >
                        <Divider />
                        <p>{props.content}</p>
                        <Divider />
                        <div className="flex justify-content-end gap-2">
                            <Button type="default" onClick={handleCancel}>
                                Cancel
                            </Button>
                            {/* Change to ButtonComponent.tsx once JJ's merge is completed with the Config Provider and Components */}
                            <Button type="primary" onClick={props.handleOkay}>
                                Confirm
                            </Button>
                        </div>
                    </Modal>
                </>
            )
            }
            {
                props.type === "Smart Locker" && (
                    <>
                        <Button type="primary" onClick={showModal}>
                            {props.buttonTitle}
                        </Button>
                        <Modal
                            className='p-3 flex-wrap-row'
                            title={<h3>{titles[props.type]}</h3>}
                            open={isModalOpen}
                            onOk={props.handleOkay}
                            onCancel={handleCancel}
                            okButtonProps={{ hidden: true, disabled: true }}
                            cancelButtonProps={{ hidden: true, disabled: true }}
                        >
                            <Divider />
                            <Form>
                                <Form.Item name="search">
                                    <Input placeholder='Search for a Tenant' />
                                </Form.Item>
                                <Form.Item name="locker-number">
                                    <Input placeholder='Locker Number' type='number' />
                                </Form.Item>
                                <Divider />
                                <div className="flex justify-content-end gap-2">
                                    {/* Cancel button */}
                                    <Form.Item name="cancel">
                                        <Button type="default" onClick={() => {
                                            setIsModalOpen(false)
                                        }}>
                                            Cancel
                                        </Button>
                                    </Form.Item>
                                    <Form.Item name="submit">
                                        <Button type="primary" htmlType="submit">
                                            Submit
                                        </Button>
                                    </Form.Item>
                                </div>
                            </Form>
                        </Modal>
                    </>
                )
            }
            {
                props.type === "Guest Parking" && (
                    <>
                        <Button type="primary" onClick={showModal}>
                            {props.buttonTitle}
                        </Button>
                        <Modal
                            className='p-3 flex-wrap-row'
                            title={<h3>{titles[props.type]}</h3>}
                            open={isModalOpen}
                            onOk={props.handleOkay}
                            onCancel={handleCancel}
                            okButtonProps={{ hidden: true, disabled: true }}
                            cancelButtonProps={{ hidden: true, disabled: true }}
                        >
                            <Divider />
                            <Form>
                                <Form.Item name="tenant-name">
                                    <Input placeholder='Tenant Name' />
                                </Form.Item>
                                <Form.Item name="license-plate-number">
                                    <Input placeholder='License Plate Number' />
                                </Form.Item>
                                <Form.Item name="car-color">
                                    <Input placeholder='Car Color' type='number' />
                                </Form.Item>
                                <Form.Item name="car-make">
                                    <Input placeholder='Car Make' />
                                </Form.Item>
                                <Form.Item name="duration-of-stay">
                                    <Input placeholder='Duration of Stay' type='number' />
                                </Form.Item>
                                <Divider />
                                <div className="flex justify-content-end gap-2">
                                    {/* Cancel button */}
                                    <Form.Item name="cancel">
                                        <Button type="default" onClick={() => {
                                            setIsModalOpen(false)
                                        }}>
                                            Cancel
                                        </Button>
                                    </Form.Item>
                                    <Form.Item name="submit">
                                        <Button type="primary" htmlType="submit">
                                            Submit
                                        </Button>
                                    </Form.Item>
                                </div>
                            </Form>
                        </Modal>
                    </>
                )
            }
            {props.type === "Add Tenant" && (
                <>
                    <Button type="primary" onClick={showModal}>
                        <PlusOutlined />
                        {props.buttonTitle}
                    </Button>
                    <Modal
                        className='p-3 flex-wrap-row'
                        title={<h3>{titles[props.type]}</h3>}
                        open={isModalOpen}
                        onOk={props.handleOkay}
                        onCancel={handleCancel}
                        okButtonProps={{ hidden: true, disabled: true }}
                        cancelButtonProps={{ hidden: true, disabled: true }}
                    >
                        <Divider />
                        <Form>
                            <Form.Item name="tenant-name">
                                <Input placeholder='Tenant Name' />
                            </Form.Item>
                            <Form.Item name="tenant-email">
                                <Input placeholder='Tenant Email' />
                            </Form.Item>
                            <Form.Item name="tenant-phone">
                                <Input placeholder='Tenant Phone' type='number' />
                            </Form.Item>
                            <Form.Item name="unit-number">
                                <Input placeholder='Unit Number' type='number' />
                            </Form.Item>
                            <Divider />
                            <div className="flex justify-content-end gap-2">
                                {/* Cancel button */}
                                <Form.Item name="cancel">
                                    <Button type="default" onClick={() => {
                                        setIsModalOpen(false)
                                    }}>
                                        Cancel
                                    </Button>
                                </Form.Item>
                                <Form.Item name="submit">
                                    <Button type="primary" htmlType="submit">
                                        Submit
                                    </Button>
                                </Form.Item>
                            </div>
                        </Form>
                    </Modal>
                </>
            )}
            {props.type === "Edit Tenant" && (
                <>
                    <Button type="primary" onClick={showModal}>
                        {props.buttonTitle}
                    </Button>
                    <Modal
                        className='p-3 flex-wrap-row'
                        title={<h3>{props.modalTitle}</h3>}
                        open={isModalOpen}
                        onOk={props.handleOkay}
                        onCancel={handleCancel}
                        okButtonProps={{ hidden: true, disabled: true }}
                        cancelButtonProps={{ hidden: true, disabled: true }}

                    >
                        <Divider />
                        <Form>
                            <Form.Item name="tenant-name">
                                <Input placeholder='Tenant Name' />
                            </Form.Item>
                            <Form.Item name="tenant-email">
                                <Input placeholder='Tenant Email' />
                            </Form.Item>
                            <Form.Item name="tenant-phone">
                                <Input placeholder='Tenant Phone' />
                            </Form.Item>
                            <Form.Item name="unit-number">
                                <Input placeholder='Unit Number' />
                            </Form.Item>
                            <Form.Item name="lease-status">
                                <Input placeholder='Lease Status' />
                            </Form.Item>
                            {/* <Form.Item name="lease-start" label="Lease Start">
                                <Input placeholder='Lease Start' type='date' />
                            </Form.Item> */}
                            <Form.Item name="lease-end" label="Lease End">
                                <Input placeholder='Lease End' type='date' />
                            </Form.Item>
                            <Divider />
                            <div className="flex justify-content-end gap-2">
                                {/* Cancel button */}
                                <Form.Item name="cancel">
                                    <Button type="default" onClick={() => {
                                        setIsModalOpen(false)
                                    }}>
                                        Cancel
                                    </Button>
                                </Form.Item>
                                <Form.Item name="submit">
                                    <Button type="primary" htmlType="submit">
                                        Submit
                                    </Button>
                                </Form.Item>
                            </div>
                        </Form>
                    </Modal>
                </>
            )
            }
        </>
    );
};
export default ModalComponent;