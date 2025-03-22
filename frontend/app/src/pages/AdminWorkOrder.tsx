import "../styles/styles.scss";

import { Tag } from "antd";
import dayjs from "dayjs";
import { Input, Select } from "antd";
import { SearchOutlined } from "@ant-design/icons";
import ModalComponent from "../components/ModalComponent";
import TableComponent from "../components/reusableComponents/TableComponent";
import ButtonComponent from "../components/reusableComponents/ButtonComponent";
import type { ColumnsType, ColumnType } from "antd/es/table/interface";
import AlertComponent from "../components/reusableComponents/AlertComponent";
import { WorkOrderData, ComplaintsData } from "../types/types";
import type { TablePaginationConfig } from "antd";
import { useState } from "react";

const getWorkOrderColumnSearchProps = (dataIndex: keyof WorkOrderData, title: string): ColumnType<WorkOrderData> => ({
    filterDropdown: (filterDropdownProps) => (
        <div style={{ padding: 8 }}>
            <Input
                placeholder={`Search ${title}`}
                value={filterDropdownProps.selectedKeys[0]}
                onChange={(e) => filterDropdownProps.setSelectedKeys(e.target.value ? [e.target.value] : [])}
                onPressEnter={() => filterDropdownProps.confirm()}
            />
        </div>
    ),
    filterIcon: (filtered) => (
        <SearchOutlined style={{ color: filtered ? "#1890ff" : undefined }} />
    ),
    onFilter: (value, record) => {
        const val = record[dataIndex];
        return val?.toString().toLowerCase().includes((value as string).toLowerCase()) ?? false;
    },
});

const getComplaintColumnSearchProps = (dataIndex: keyof ComplaintsData, title: string): ColumnType<ComplaintsData> => ({
    filterDropdown: (filterDropdownProps) => (
        <div style={{ padding: 8 }}>
            <Input
                placeholder={`Search ${title}`}
                value={filterDropdownProps.selectedKeys[0]}
                onChange={(e) => filterDropdownProps.setSelectedKeys(e.target.value ? [e.target.value] : [])}
                onPressEnter={() => filterDropdownProps.confirm()}
            />
        </div>
    ),
    filterIcon: (filtered) => (
        <SearchOutlined style={{ color: filtered ? "#1890ff" : undefined }} />
    ),
    onFilter: (value, record) => {
        const val = record[dataIndex];
        return val?.toString().toLowerCase().includes((value as string).toLowerCase()) ?? false;
    },
});

const shortenInput = (input: string, maxLength: number = 30) => {
    if (input.length > maxLength) {
        return input.substring(0, maxLength - 3) + "...";
    } else {
        return input;
    }
};

// DUMMY DATA THIS WILL BE DELETED :D
const workOrderDataRaw: WorkOrderData[] = [
    {
        key: 1,
        workOrderNumber: 10001,
        creatingBy: 3,
        category: "plumbing",
        title: "Leaking Kitchen Sink",
        description: "Water is slowly leaking from under the kitchen sink and forming a puddle on the floor.",
        apartmentNumber: "C466",
        status: "open",
        createdAt: new Date("2025-02-15T09:30:00"),
        updatedAt: new Date("2025-02-15T09:30:00"),
    },
    {
        key: 2,
        workOrderNumber: 10002,
        creatingBy: 1,
        category: "electrical",
        title: "Bathroom Light Flickering",
        description: "The bathroom light has been flickering for two days and sometimes goes out completely.",
        apartmentNumber: "B218",
        status: "in_progress",
        createdAt: new Date("2025-02-10T14:45:00"),
        updatedAt: new Date("2025-02-12T11:20:00"),
    },
    {
        key: 3,
        workOrderNumber: 10003,
        creatingBy: 10,
        category: "hvac",
        title: "AC Not Cooling",
        description: "Air conditioner is running but not cooling the apartment. Temperature is getting uncomfortable.",
        apartmentNumber: "A101",
        status: "awaiting_parts",
        createdAt: new Date("2025-01-30T16:20:00"),
        updatedAt: new Date("2025-02-02T09:15:00"),
    },
    {
        key: 4,
        workOrderNumber: 10004,
        creatingBy: 5,
        category: "carpentry",
        title: "Broken Cabinet Door",
        description: "Kitchen cabinet door hinge is broken and the door won't stay closed.",
        apartmentNumber: "C378",
        status: "completed",
        createdAt: new Date("2025-03-17T00:00:00"),
        updatedAt: new Date("2025-03-17T00:00:00"),
    },
    {
        key: 5,
        workOrderNumber: 10005,
        creatingBy: 8,
        category: "plumbing",
        title: "Clogged Toilet",
        description: "Toilet is clogged and won't flush properly. Plunger hasn't helped.",
        apartmentNumber: "C299",
        status: "open",
        createdAt: new Date("2025-02-18T08:10:00"),
        updatedAt: new Date("2025-02-18T08:10:00"),
    },
    {
        key: 6,
        workOrderNumber: 10006,
        creatingBy: 2,
        category: "electrical",
        title: "No Power in Bedroom",
        description: "Electrical outlets in the bedroom aren't working. Breaker hasn't tripped.",
        apartmentNumber: "A212",
        status: "in_progress",
        createdAt: new Date("2025-02-14T12:30:00"),
        updatedAt: new Date("2025-02-14T16:45:00"),
    },
    {
        key: 7,
        workOrderNumber: 10007,
        creatingBy: 4,
        category: "other",
        title: "Stuck Window",
        description: "Living room window is stuck and won't open. Frame seems to be warped.",
        apartmentNumber: "B179",
        status: "open",
        createdAt: new Date("2025-02-17T11:25:00"),
        updatedAt: new Date("2025-02-17T11:25:00"),
    },
    {
        key: 8,
        workOrderNumber: 10008,
        creatingBy: 6,
        category: "hvac",
        title: "Noisy Heater",
        description: "Heating system is making loud banging noises when it starts up.",
        apartmentNumber: "A333",
        status: "awaiting_parts",
        createdAt: new Date("2025-03-14T09:50:00"),
        updatedAt: new Date("2025-01-29T14:20:00"),
    },
    {
        key: 9,
        workOrderNumber: 10009,
        creatingBy: 9,
        category: "plumbing",
        title: "Low Water Pressure",
        description: "Water pressure in the shower is very low. All other faucets seem normal.",
        apartmentNumber: "B155",
        status: "completed",
        createdAt: new Date("2025-01-20T13:15:00"),
        updatedAt: new Date("2025-01-23T10:40:00"),
    },
    {
        key: 10,
        workOrderNumber: 10010,
        creatingBy: 7,
        category: "carpentry",
        title: "Damaged Baseboards",
        description: "Baseboards in the living room are damaged and coming away from the wall in several places.",
        apartmentNumber: "D401",
        status: "in_progress",
        createdAt: new Date("2025-02-12T15:00:00"),
        updatedAt: new Date("2025-02-13T11:30:00"),
    },
];

const workOrderColumns: ColumnsType<WorkOrderData> = [
    {
        title: "Work Order #",
        dataIndex: "workOrderNumber",
        key: "workOrderNumber",
        ...getWorkOrderColumnSearchProps("workOrderNumber", "Work Order #"),
        sorter: (a, b) => a.workOrderNumber - b.workOrderNumber,
    },
    {
        title: "Category",
        dataIndex: "category",
        key: "category",
        ...getWorkOrderColumnSearchProps("category", "Category"),
        render: (category) => {
            let color = "";
            let text = "";

            switch (category) {
                case "plumbing":
                    text = "Plumbing 🛀";
                    color = "blue";
                    break;
                case "electrical":
                    text = "Electrical ⚡";
                    color = "yellow";
                    break;
                case "carpentry":
                    text = "Carpentry 🪚";
                    color = "brown";
                    break;
                case "hvac":
                    text = "HVAC 🌡️";
                    color = "grey";
                    break;
                default:
                    text = "Other";
            }

            return <Tag color={color}>{text}</Tag>;
        },
        className: "text-center",
    },
    {
        title: "Title",
        dataIndex: "title",
        key: "title",
        sorter: (a, b) => a.title.localeCompare(b.title),
        ...getWorkOrderColumnSearchProps("title", "Inquiry"),
        render: (title: string) => shortenInput(title, 25),
    },
    {
        title: "Description",
        dataIndex: "description",
        key: "description",
        ...getWorkOrderColumnSearchProps("description", "Description"),
        render: (description: string) => shortenInput(description),
    },
    {
        title: "Unit",
        dataIndex: "apartmentNumber",
        key: "apartmentNumber",
        ...getWorkOrderColumnSearchProps("apartmentNumber", "Unit"),
    },
    {
        title: "Status",
        dataIndex: "status",
        key: "status",
        ...getWorkOrderColumnSearchProps("status", "Status"),
        render: (status: string) => {
            let color = "default";
            switch (status) {
                case "open":
                    color = "red";
                    break;
                case "in_progress":
                    color = "blue";
                    break;
                case "awaiting_parts":
                    color = "orange";
                    break;
                case "completed":
                    color = "green";
                    break;
            }
            return <Tag color={color}>{status.replace("_", " ").toUpperCase()}</Tag>;
        },
        className: "text-center",
    },
    {
        title: "Created",
        dataIndex: "createdAt",
        key: "createdAt",
        ...getWorkOrderColumnSearchProps("createdAt", "Created"),
        sorter: (a, b) => dayjs(a.createdAt).unix() - dayjs(b.createdAt).unix(),
        render: (date) => dayjs(date).format("MMM D, YYYY h:mm A"),
    },
    {
        title: "Updated",
        dataIndex: "updatedAt",
        key: "updatedAt",
        ...getWorkOrderColumnSearchProps("updatedAt", "Updated"),
        sorter: (a, b) => dayjs(a.updatedAt).unix() - dayjs(b.updatedAt).unix(),
        render: (date) => dayjs(date).format("MMM D, YYYY h:mm A"),
    },
];

const complaintsDataRaw: ComplaintsData[] = [
    {
        key: 1,
        complaintNumber: 20001,
        createdBy: 4,
        category: "noise",
        title: "Loud Music at Night",
        description: "Neighbor plays loud music past midnight.",
        unitNumber: "A312",
        status: "open",
        createdAt: new Date("2025-03-10T22:15:00"),
        updatedAt: new Date("2025-03-11T08:00:00"),
    },
    {
        key: 2,
        complaintNumber: 20002,
        createdBy: 7,
        category: "parking",
        title: "Unauthorized Vehicle in My Spot",
        description: "A car is parked in my designated space.",
        unitNumber: "B210",
        status: "in_progress",
        createdAt: new Date("2025-02-28T18:30:00"),
        updatedAt: new Date("2025-03-01T09:45:00"),
    },
    {
        key: 3,
        complaintNumber: 20003,
        createdBy: 2,
        category: "maintenance",
        title: "Leaking Roof",
        description: "Water leaking from ceiling during rainstorms.",
        unitNumber: "C405",
        status: "resolved",
        createdAt: new Date("2025-02-20T14:00:00"),
        updatedAt: new Date("2025-02-22T16:00:00"),
    },
    {
        key: 4,
        complaintNumber: 20004,
        createdBy: 10,
        category: "security",
        title: "Suspicious Person Near Entrance",
        description: "Unfamiliar person lingering around entrance at night.",
        unitNumber: "E102",
        status: "closed",
        createdAt: new Date("2025-03-02T20:00:00"),
        updatedAt: new Date("2025-03-03T12:00:00"),
    },
];

const complaintsColumns: ColumnsType<ComplaintsData> = [
    {
        title: "Complaint #",
        dataIndex: "complaintNumber",
        key: "complaintNumber",
        ...getComplaintColumnSearchProps("complaintNumber", "Complaint #"),
    },
    {
        title: "Category",
        dataIndex: "category",
        key: "category",
        sorter: (a, b) => a.category.localeCompare(b.category),
        filters: [
            { text: "Maintenance", value: "maintenance" },
            { text: "Noise", value: "noise" },
            { text: "Security", value: "security" },
            { text: "Parking", value: "parking" },
            { text: "Neighbor", value: "neighbor" },
            { text: "Trash", value: "trash" },
            { text: "Internet", value: "internet" },
            { text: "Lease", value: "lease" },
            { text: "Natural Disaster", value: "natural_disaster" },
            { text: "Other", value: "other" },
        ],
        onFilter: (value, record) => record.category === (value as ComplaintsData["category"]),
        render: (category) => {
            let color = "";
            let text = "";

            switch (category) {
                case "maintenance":
                    text = "Maintenance 🔧";
                    color = "blue";
                    break;
                case "noise":
                    text = "Noise 🔊";
                    color = "orange";
                    break;
                case "security":
                    text = "Security 🔒";
                    color = "red";
                    break;
                case "parking":
                    text = "Parking 🚗";
                    color = "purple";
                    break;
                case "neighbor":
                    text = "Neighbor 🏘️";
                    color = "green";
                    break;
                case "trash":
                    text = "Trash 🗑️";
                    color = "brown";
                    break;
                case "internet":
                    text = "Internet 🌐";
                    color = "cyan";
                    break;
                case "lease":
                    text = "Lease 📝";
                    color = "gold";
                    break;
                case "natural_disaster":
                    text = "Disaster 🌪️";
                    color = "grey";
                    break;
                default:
                    text = "Other";
                    color = "default";
            }

            return <Tag color={color}>{text}</Tag>;
        },
        className: "text-center",
    },
    {
        title: "Title",
        dataIndex: "title",
        key: "title",
        ...getComplaintColumnSearchProps("title", "Title"),
        render: (title: string) => shortenInput(title, 25),
    },
    {
        title: "Description",
        dataIndex: "description",
        key: "description",
        ...getComplaintColumnSearchProps("description", "Description"),
        render: (description: string) => shortenInput(description),
    },
    {
        title: "Unit",
        dataIndex: "unitNumber",
        key: "unitNumber",
        ...getComplaintColumnSearchProps("unitNumber", "Unit"),
    },
    {
        title: "Status",
        dataIndex: "status",
        key: "status",
        ...getComplaintColumnSearchProps("status", "Status"),
        render: (status: string) => {
            let color = "default";
            switch (status) {
                case "in_progress":
                    color = "blue";
                    break;
                case "resolved":
                    color = "green";
                    break;
                case "closed":
                    color = "gray";
                    break;
                case "open":
                    color = "red";
                    break;
            }
            return <Tag color={color}>{status.replace("_", " ").toUpperCase()}</Tag>;
        },
        className: "text-center",
    },
    {
        title: "Created",
        dataIndex: "createdAt",
        key: "createdAt",
        sorter: (a, b) => dayjs(a.createdAt).unix() - dayjs(b.createdAt).unix(),
        render: (date) => dayjs(date).format("MMM D, YYYY h:mm A"),
    },
    {
        title: "Updated",
        dataIndex: "updatedAt",
        key: "updatedAt",
        sorter: (a, b) => dayjs(a.updatedAt).unix() - dayjs(b.updatedAt).unix(),
        render: (date) => dayjs(date).format("MMM D, YYYY h:mm A"),
    },
];

const paginationConfig: TablePaginationConfig = {
    pageSize: 5,
    showSizeChanger: false,
};

const AdminWorkOrder = () => {
    const [workOrderData, setWorkOrderData] = useState<WorkOrderData[]>(workOrderDataRaw);
    const [complaintsData, setComplaintsData] = useState<ComplaintsData[]>(complaintsDataRaw);
    const [selectedItem, setSelectedItem] = useState<WorkOrderData | ComplaintsData | null>(null);
    const [isModalVisible, setIsModalVisible] = useState(false);
    const [itemType, setItemType] = useState<"workOrder" | "complaint">("workOrder");
    const [currentStatus, setCurrentStatus] = useState<string>("");

    const handleStatusChange = (newStatus: string) => {
        setCurrentStatus(newStatus);
    };

    const handleConfirm = () => {
        if (selectedItem && currentStatus) {
            if (itemType === "workOrder") {
                const updatedWorkOrders = workOrderData.map(item => {
                    if (item.key === selectedItem.key) {
                        return {
                            ...item,
                            status: currentStatus,
                            updatedAt: new Date()
                        };
                    }
                    return item;
                });
                setWorkOrderData(updatedWorkOrders);
            } else {
                const updatedComplaints = complaintsData.map(item => {
                    if (item.key === selectedItem.key) {
                        return {
                            ...item,
                            status: currentStatus,
                            updatedAt: new Date(),
                        };
                    }
                    return item;
                });
                setComplaintsData(updatedComplaints);
            }
            setIsModalVisible(false);
        }
    };

    const getUnitNumber = (item: WorkOrderData | ComplaintsData): string => {
        // I HAVE NO CLUE WHAT THE NAMING CONVENTION IS NOW SO ADDING THIS SINCE I'VE SEEN BOTH
        // FOR WORK ORDER AND COMPLAINTS
        if ("apartmentNumber" in item) {
            return item.apartmentNumber;
        }
        return item.unitNumber;
    };

    const sortedWorkOrders = workOrderDataRaw.sort((a, b) => {
        const statusPriority = { open: 1, in_progress: 2, awaiting_parts: 3, completed: 4 };
        const priorityDiff = statusPriority[a.status] - statusPriority[b.status];
        if (priorityDiff !== 0) return priorityDiff;

        if (a.status !== "completed" && b.status !== "completed") {
            return new Date(a.createdAt).getTime() - new Date(b.createdAt).getTime();
        }

        return new Date(b.updatedAt).getTime() - new Date(a.updatedAt).getTime();
    });

    const sortedComplaints = complaintsDataRaw.sort((a, b) => {
        const statusPriority = { open: 1, in_progress: 2, resolved: 3, closed: 4 };
        const priorityDiff = statusPriority[a.status] - statusPriority[b.status];
        if (priorityDiff !== 0) { return priorityDiff; }

        if (!(a.status in ["resolved", "closed"]) && !(b.status in ["resolved", "closed"])) {
            return new Date(a.createdAt).getTime() - new Date(b.createdAt).getTime();
        }

        return new Date(b.updatedAt).getTime() - new Date(a.updatedAt).getTime();
    });

    const hoursUntilOverdue: number = 48;
    const overdueServiceCount: number = workOrderDataRaw.filter(({ createdAt, status }) => {
        const hoursSinceCreation = dayjs().diff(dayjs(createdAt), "hour");
        return status === "open" && hoursSinceCreation >= hoursUntilOverdue;
    }).length;
    const hoursSinceRecentlyCreated: number = 24;
    const recentlyCreatedServiceCount: number = workOrderDataRaw.filter(({ createdAt }) => {
        const hoursSinceCreation = dayjs().diff(dayjs(createdAt), "hour");
        return hoursSinceCreation <= hoursSinceRecentlyCreated;
    }).length;

    const hoursSinceRecentlyCompleted: number = 24;
    const recentlyCompletedServiceCount: number = workOrderDataRaw.filter(({ updatedAt, status }) => {
        const hoursSinceUpdate = dayjs().diff(dayjs(updatedAt), "hour");
        return status === "completed" && hoursSinceUpdate <= hoursSinceRecentlyCompleted;
    }).length;

    let alerts: string[] = [];
    if (overdueServiceCount > 0) {
        alerts.push(`${overdueServiceCount} services open for >${hoursUntilOverdue} hours.`);
    } else if (recentlyCreatedServiceCount > 0) {
        alerts.push(`${recentlyCreatedServiceCount} services created in past ${hoursSinceRecentlyCreated} hours.`)
    } else if (recentlyCompletedServiceCount > 0) {
        alerts.push(`${recentlyCompletedServiceCount} services completed in past ${hoursSinceRecentlyCompleted} hours.`)
    }
    const alertDescription: string = alerts.join(" ") ?? "";

    const modalContent = selectedItem && (
        <div>
            <div className="mb-4">
                <strong>Title:</strong> {selectedItem.title}
            </div>
            <div className="mb-4">
                <strong>Description:</strong> {selectedItem.description}
            </div>
            <div className="mb-4">
                <strong>Unit Number:</strong> {getUnitNumber(selectedItem)}
            </div>
            <div>
                <strong>Status:</strong>
                <Select
                    defaultValue={selectedItem.status}
                    style={{ width: 200, marginLeft: 10 }}
                    onChange={handleStatusChange}
                >
                    {itemType === "workOrder" ? (
                        <>
                            <Select.Option value="open">Open</Select.Option>
                            <Select.Option value="in_progress">In Progress</Select.Option>
                            <Select.Option value="awaiting_parts">Awaiting Parts</Select.Option>
                            <Select.Option value="completed">Completed</Select.Option>
                        </>
                    ) : (
                        <>
                            <Select.Option value="open">Open</Select.Option>
                            <Select.Option value="in_progress">In Progress</Select.Option>
                            <Select.Option value="resolved">Resolved</Select.Option>
                            <Select.Option value="closed">Closed</Select.Option>
                        </>
                    )}
                </Select>
            </div>
        </div>
    );

    return (
        <div className="container">
            <h1 className="mb-4">Work-Orders & Complaints</h1>

            {/* Alerts headers */}
            <div className="w-100 justify-content-between mb-4 left-text text-start">
                {alertDescription ? <AlertComponent description={alertDescription} /> : null}
            </div>

            {/* Work Order Table */}
            <div className="mb-5">
                <h4 className="mb-3">Work Orders</h4>
                <TableComponent<WorkOrderData>
                    columns={workOrderColumns}
                    dataSource={workOrderData}
                    style=".lease-table-container"
                    pagination={paginationConfig}
                    onChange={(pagination, filters, sorter, extra) => {
                        console.log("Table changed:", pagination, filters, sorter, extra);
                    }}
                    onRow={(record: WorkOrderData) => ({
                        onClick: () => {
                            setSelectedItem(record);
                            setItemType("workOrder");
                            setIsModalVisible(true);
                        },
                        style: {
                            cursor: 'pointer',
                        },
                        className: 'hoverable-row'
                    })}
                />
            </div>

            {/* Complaints Table */}
            <div className="mb-5">
                <h4 className="mb-3">Complaints</h4>

                <TableComponent<ComplaintsData>
                    columns={complaintsColumns}
                    dataSource={complaintsData}
                    style=".lease-table-container"
                    pagination={paginationConfig}
                    onChange={(pagination, filters, sorter, extra) => {
                        console.log("Table changed:", pagination, filters, sorter, extra);
                    }}
                    onRow={(record: ComplaintsData) => ({
                        onClick: () => {
                            setSelectedItem(record);
                            setItemType("complaint");
                            setIsModalVisible(true);
                        },
                        style: {
                            cursor: 'pointer',
                        },
                        className: 'hoverable-row'
                    })}
                />
            </div>

            {selectedItem && (
                <ModalComponent
                    buttonTitle=""
                    buttonType="default"
                    content={modalContent}
                    type="default"
                    handleOkay={handleConfirm}
                    modalTitle={`${itemType === "workOrder" ? "Work Order" : "Complaint"} Details`}
                    isModalOpen={isModalVisible}
                    onCancel={() => setIsModalVisible(false)}
                    apartmentBuildingSetEditBuildingState={() => { }}
                />
            )}
        </div>
    );
};

export default AdminWorkOrder;
