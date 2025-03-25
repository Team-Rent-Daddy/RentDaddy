import Icon, { ToolOutlined, WarningOutlined, InboxOutlined, CalendarOutlined, UserOutlined, CarOutlined } from "@ant-design/icons";
import { Tag, Modal, Button } from "antd";
import React, { useState, useEffect } from "react";
import { Link } from "react-router";
import ModalComponent from "../components/ModalComponent";
import AlertComponent from "../components/reusableComponents/AlertComponent";
import ButtonComponent from "../components/reusableComponents/ButtonComponent";
import { CardComponent } from "../components/reusableComponents/CardComponent";
import PageTitleComponent from "../components/reusableComponents/PageTitleComponent";
import MyChatBot from "../components/ChatBot";
import { useAuth } from "@clerk/react-router";
import { useMutation, useQuery } from "@tanstack/react-query";

export const TenantDashBoard = () => {
    const handleOpenLocker = () => {
        console.log("handle open locker");
        // Add your logic for getting a package here
    };

    const [isSigningModalVisible, setSigningModalVisible] = useState(false);
    const { userId } = useAuth();

    const [latestComplaint, setLatestComplaint] = useState<any[]>([]);
    const [latestWorkOrder, setLatestWorkOrder] = useState<any[]>([]);

    const { mutate: getComplaints } = useMutation({
        mutationFn: async () => {
            const res = await fetch(`http://localhost:8080/tenants/complaints`, {
                method: "GET",
                headers: { "Content-Type": "application/json" },
            });
            const data = await res.json();
            setLatestComplaint([data?.[0]]);
            return res;
        },
        onSuccess: () => {},
        onError: (e: any) => {
            console.log("error ", e);
        },
    });

    const { mutate: getWorkOrders } = useMutation({
        mutationFn: async () => {
            const res = await fetch(`http://localhost:8080/tenants/work_orders`, {
                method: "GET",
                headers: { "Content-Type": "application/json" },
            });
            const data = await res.json();
            setLatestWorkOrder([data?.[0]]);
            return res;
        },
        onSuccess: () => {},
        onError: (e: any) => {
            console.log("error ", e);
        },
    });

    const { mutate: postParking } = useMutation({
        mutationFn: async (parking) => {
            console.log(parking);
            const res = await fetch(`http://localhost:8080/tenants/parking`, {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify(parking),
            });
            return res;
        },
    });

    // Simulate fetching lease status using TanStack Query
    const {
        data: leaseStatus,
        isLoading,
        isError,
    } = useQuery({
        queryKey: ["leaseStatus", userId], // Unique key for the query
        queryFn: async () => {
            // Simulate a delay to mimic network request and give dummy data
            await new Promise((resolve) => setTimeout(resolve, 500));
            const leaseData = {
                // userId: userId,
                userId: "notme",
                lease_status: "pending_approval",
            };
            // const response = await fetch(`/api/leases?tenantId=${userId}`);
            // if (!response.ok) {
            //     throw new Error("Failed to fetch lease status");
            // }
            // const data = await response.json();

            // Return dummy data if the userId matches
            if (userId === leaseData.userId) {
                console.log(leaseData.lease_status);
                return leaseData.lease_status;
            } else {
                return "active";
            }
        },
        enabled: !!userId,
    });

    // This is the recommended approach in newer versions of TanStack Query. `onSuccess` is deprecated
    useEffect(() => {
        if (leaseStatus) {
            console.log("Lease status updated:", leaseStatus);
            if (["pending_approval", "terminated", "expired"].includes(leaseStatus)) {
                console.log("Setting modal visible based on lease status");
                setSigningModalVisible(true);
            }
        }
        getComplaints();
        getWorkOrders();
    }, [leaseStatus, latestComplaint]);

    const handleOk = () => {
        // Redirect to the lease signing page (THIS ISNT IT AT ALL, NEEDS documenso uri. TMP for now)
        window.location.href = "/tenant/sign-lease";
    };

    if (isLoading) {
        return <div>Loading...</div>;
    }

    if (isError) {
        return <div>Error fetching tenant's lease status. Please try again later.</div>;
    }

    return (
        <div className="container">
            {/* <h1 className="my-4">Tenant Dashboard</h1> */}
            <PageTitleComponent title="Tenant Dashboard" />
            {/* <div className="alert-container"> */}
            <AlertComponent
                message="Welcome to the Tenant Dashboard"
                description="Sign Yo Lease. Pay Daddy Rent"
                type="warning"
            />
            {/* </div> */}

            {/* Dashboard Statistics Cards */}
            <h2 className="my-3 p-3 text-center">Quick Actions</h2>
            <div className="flex-container my-3">
                <CardComponent
                    title="Open Complaint form"
                    value={10}
                    description="Something not working right or disturbing you? Let us know."
                    hoverable={true}
                    icon={<ToolOutlined className="icon" />}
                    button={
                        <Link to="/tenant/tenant-work-orders-and-complaints">
                            <ButtonComponent
                                title="View All"
                                type="primary"
                                onClick={() => {}}
                            />
                        </Link>
                    }
                />
                <CardComponent
                    title="Package info"
                    value={10}
                    description="You have a package. Click the button at your locker to open it"
                    hoverable={true}
                    icon={<InboxOutlined className="icon" />}
                    button={
                        <ModalComponent
                            type="Smart Locker"
                            userRole="tenant"
                            buttonTitle="Open Locker"
                            content="Open Locker"
                            buttonType="primary"
                            handleOkay={() => handleOpenLocker()}
                        />
                    }
                />
                <CardComponent
                    title="Guest Parking"
                    value={10}
                    description="Got a guest coming to visit? Make sure they have spots to park"
                    hoverable={true}
                    icon={<CarOutlined className="icon" />}
                    button={
                        <ModalComponent
                            type="Guest Parking"
                            buttonTitle="Add Guest"
                            content="Add guest to be able to park in the complex"
                            buttonType="primary"
                            handleOkay={(data: any) => postParking(data)}
                        />
                    }
                />
            </div>

            {/* Quick Access Documents Section */}
            <h2 className="my-3 p-3 text-center">Quick Access Documents Section</h2>
            <div className="flex-container mb-3">
                <CardComponent
                    title="Lease"
                    description="View or Resign your lease"
                    hoverable={true}
                    button={
                        <ModalComponent
                            type="default"
                            buttonTitle="View Lease"
                            content="Lease should go here"
                            buttonType="primary"
                            handleOkay={() => {}}
                        />
                    }
                />
                {latestWorkOrder.length ? (
                    <CardComponent
                        title="Work Order"
                        description={<Tag color="orange">Current In Progress</Tag>}
                        hoverable={true}
                        button={
                            <ModalComponent
                                type="default"
                                buttonTitle="View all work orders"
                                content="Work orders should go here"
                                buttonType="primary"
                                handleOkay={() => {}}
                            />
                        }
                    />
                ) : undefined}

                {latestComplaint.length ? (
                    <CardComponent
                        title="Complaint Received"
                        description={`Our office received your complaint and will investigate immediately. "${latestComplaint}"`}
                        hoverable={true}
                        button={
                            <Link to="/tenant/tenant-work-orders-and-complaints">
                                <ButtonComponent
                                    title="View all complaints"
                                    type="primary"
                                    onClick={() => {}}
                                />
                            </Link>
                        }
                    />
                ) : undefined}

                <MyChatBot />
            </div>

            {/* Inescapable Modal for lease signing */}
            <Modal
                title="Action Required: Lease Signing"
                open={isSigningModalVisible}
                onOk={handleOk}
                onCancel={() => {}} // Empty function prevents closing
                maskClosable={false} // Prevents closing when clicking outside
                keyboard={false} // Prevents closing with ESC key
                closable={false} // Removes the X button
                footer={[
                    <Button
                        key="submit"
                        type="primary"
                        onClick={handleOk}>
                        Sign Lease Now
                    </Button>,
                ]}>
                <div style={{ textAlign: "center" }}>
                    <WarningOutlined style={{ fontSize: "4rem", color: "#faad14", marginBottom: "1rem" }} />
                    <h3 style={{ marginBottom: "1rem" }}>Your Lease Requires Attention</h3>
                    <p>
                        Your lease status is <strong>{leaseStatus === "pending_approval" ? "Pending Approval" : leaseStatus}</strong>.
                    </p>
                    <p>You must sign your lease to continue using the tenant portal.</p>
                    <p style={{ marginTop: "1rem", fontStyle: "italic" }}>This action is required and cannot be dismissed.</p>
                </div>
            </Modal>
        </div>
    );
};
