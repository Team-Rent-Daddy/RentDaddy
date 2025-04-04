import "../styles/styles.scss";

import { Tag } from "antd";
import dayjs from "dayjs";
import { Input, Select } from "antd";
import { SearchOutlined } from "@ant-design/icons";
import ModalComponent from "../components/ModalComponent";
import TableComponent from "../components/reusableComponents/TableComponent";
import type { ColumnsType, ColumnType } from "antd/es/table/interface";
import { WorkOrderData, ComplaintsData } from "../types/types";
import type { TablePaginationConfig } from "antd";
import { useState } from "react";
import PageTitleComponent from "../components/reusableComponents/PageTitleComponent";
import { useQuery, useQueryClient } from "@tanstack/react-query";
import { useAuth } from "@clerk/clerk-react";

const serverUrl = import.meta.env.VITE_SERVER_URL;
const absoluteServerUrl = `${serverUrl}`;

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
    filterIcon: (filtered) => <SearchOutlined style={{ color: filtered ? "#1890ff" : undefined }} />,
    onFilter: (value, record) => {
        const val = record[dataIndex];
        return (
            val
                ?.toString()
                .toLowerCase()
                .includes((value as string).toLowerCase()) ?? false
        );
    },
});

const shortenInput = (input: string, maxLength: number = 30) => {
    if (input.length > maxLength) {
        return input.substring(0, maxLength - 3) + "...";
    } else {
        return input;
    }
};

const workOrderColumns: ColumnsType<WorkOrderData> = [
    {
        title: "Work Order ID",
        dataIndex: "id",
        key: "id",
        ...getWorkOrderColumnSearchProps("id", "Work Order ID"),
        sorter: (a, b) => a.id - b.id,
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
        dataIndex: "unitNumber",
        key: "unitNumber",
        ...getWorkOrderColumnSearchProps("unitNumber", "Unit"),
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
                    color = "orange";
                    break;
                case "resolved":
                    color = "blue";
                    break;
                case "closed":
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

const paginationConfig: TablePaginationConfig = {
    pageSize: 5,
    showSizeChanger: false,
};

const AdminWorkOrder = () => {
    // const [workOrderData, setWorkOrderData] = useState<WorkOrderData[]>(workOrderDataRaw);
    // const [complaintsData, setComplaintsData] = useState<ComplaintsData[]>(complaintsDataRaw);
    const [selectedItem, setSelectedItem] = useState<WorkOrderData | ComplaintsData | null>(null);
    const [isModalVisible, setIsModalVisible] = useState(false);
    const [itemType, setItemType] = useState<"workOrder" | "complaint">("workOrder");
    const [currentStatus, setCurrentStatus] = useState<string>("");

    const { getToken } = useAuth();
    const queryClient = useQueryClient();

    const {
        data: workOrderData,
        isLoading: isWorkOrdersLoading,
        error: workOrdersError,
    } = useQuery({
        queryKey: ["workOrders"],
        queryFn: async () => {
            const token = await getToken();
            const response = await fetch(`${absoluteServerUrl}/admin/work_orders`, {
                method: "GET",
                headers: {
                    Authorization: `Bearer ${token}`,
                    "Content-Type": "application/json",
                },
            });
            if (!response.ok) {
                throw new Error("Failed to fetch work orders");
            }
            const data = (await response.json()) as WorkOrderData[];
            if (!Array.isArray(data)) {
                throw new Error("No work orders");
            }

            return data;
        },
    });

    const handleStatusChange = (newStatus: string) => {
        setCurrentStatus(newStatus);
    };

    const handleConfirm = async () => {
        if (selectedItem && currentStatus) {
            try {
                const token = await getToken();
                const response = await fetch(`${absoluteServerUrl}/admin/complaints/${selectedItem.id}/status`, {
                    method: "PATCH",
                    headers: {
                        "Content-Type": "application/json",
                        Authorization: `Bearer ${token}`,
                    },
                    body: JSON.stringify({
                        status: currentStatus,
                    }),
                });

                if (!response.ok) {
                    throw new Error("Failed to update complaint");
                }

                queryClient.setQueryData(["complaints"], (oldData: ComplaintsData[] | undefined) => {
                    if (!oldData) return oldData;
                    return oldData.map((item) => (item.id === selectedItem.id ? { ...item, status: currentStatus, updatedAt: new Date() } : item));
                });
                setIsModalVisible(false);
            } catch (error) {
                console.error("Error updating status:", error);
            }
        }
    };

    const handleRowClick = (record: WorkOrderData | ComplaintsData, type: "workOrder") => {
        setSelectedItem(record);
        setItemType(type);
        setCurrentStatus(record.status);
        setIsModalVisible(true);
    };

    const hoursUntilOverdue: number = 48;
    const overdueServiceCount: number = workOrderData
        ? workOrderData.filter(({ createdAt, status }) => {
              const hoursSinceCreation = dayjs().diff(dayjs(createdAt), "hour");
              return status === "open" && hoursSinceCreation >= hoursUntilOverdue;
          }).length
        : 0;

    const hoursSinceRecentlyCreated: number = 24;
    const recentlyCreatedServiceCount: number = workOrderData
        ? workOrderData.filter(({ createdAt }) => {
              const hoursSinceCreation = dayjs().diff(dayjs(createdAt), "hour");
              return hoursSinceCreation <= hoursSinceRecentlyCreated;
          }).length
        : 0;

    const alerts: string[] = [];
    if (isWorkOrdersLoading) {
        alerts.push("Loading data...");
    } else if (workOrdersError) {
        alerts.push("Error loading data");
    } else {
        if (workOrderData?.length === 0) {
            alerts.push("No work orders found");
        }
        if (workOrderData && workOrderData.length > 0) {
            if (overdueServiceCount > 0) {
                alerts.push(`${overdueServiceCount} services open for >${hoursUntilOverdue} hours.`);
            } else if (recentlyCreatedServiceCount > 0) {
                alerts.push(`${recentlyCreatedServiceCount} services created recently.`);
            }
        }
    }

    const modalContent = selectedItem && (
        <div>
            <div className="mb-4">
                <strong>Title:</strong> {selectedItem.title}
            </div>
            <div className="mb-4">
                <strong>Description:</strong> {selectedItem.description}
            </div>
            <div className="mb-4">
                <strong>Unit Number:</strong> {selectedItem.unitNumber}
            </div>
            <div>
                <strong>Status:</strong>
                <Select
                    value={currentStatus}
                    style={{ width: 200, marginLeft: 10 }}
                    onChange={handleStatusChange}>
                    {itemType === "workOrder" ? (
                        <>
                            <Select.Option value="open">Open</Select.Option>
                            <Select.Option value="in_progress">In Progress</Select.Option>
                            <Select.Option value="resolved">Resolved</Select.Option>
                            <Select.Option value="closed">Closed</Select.Option>
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
            {/* PageTitleComponent header */}
            <PageTitleComponent title="Work Orders" />
            {/* Work Order Table */}
            <div className="mb-5">
                <TableComponent<WorkOrderData>
                    columns={workOrderColumns}
                    dataSource={workOrderData || []}
                    style=".lease-table-container"
                    loading={isWorkOrdersLoading}
                    pagination={paginationConfig}
                    onChange={(pagination, filters, sorter, extra) => {
                        console.log("Table changed:", pagination, filters, sorter, extra);
                    }}
                    onRow={(record: WorkOrderData) => ({
                        onClick: () => handleRowClick(record, "workOrder"),
                        style: {
                            cursor: "pointer",
                        },
                        className: "hoverable-row",
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
                    apartmentBuildingSetEditBuildingState={() => {}}
                />
            )}
        </div>
    );
};

export default AdminWorkOrder;
